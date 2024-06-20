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
      if (false) {
        return !_dafny.areEqual(_dafny.Seq.UnicodeFromString("wvefrkb"), _dafny.Seq.UnicodeFromString("jwssl"));
      } else {
        return false;
      }
    };
    static fm1(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("domj"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), function (_0_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("vg")));
    };
    static fm2(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, true, true, false), _dafny.Seq.of(!(false))), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(!(true))));
    };
    static fm3(p0, p1, globalState) {
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))).multipliedBy((new BigNumber((_dafny.Seq.of(!(false), !(false))).length)).minus(new BigNumber(598)));
    };
    static fm15(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), function (_1_i0) {
        return new BigNumber(519);
      }),true);
    };
    static fm16(p0, globalState) {
      return (((true) ? (_dafny.Set.fromElements(new BigNumber(892), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length)), new BigNumber((_dafny.Set.fromElements(new BigNumber(353), new BigNumber(-515))).length))) : (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(-214))).length))))).Union((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-357), new BigNumber(330))) {
          let _2_v0 = _compr_0;
          if (((new BigNumber(-357)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(330)))) {
            _coll0.add(_module.__default.safeDivisionInt(_2_v0, new BigNumber((_dafny.Seq.of(false)).length)));
          }
        }
        return _coll0;
      }()).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(799),true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(52),true)).length))));
    };
    static fm17(p0, p1, globalState) {
      let _source0 = _module.D1.create_DC4(new BigNumber((function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("prkc")).length))))).Elements) {
    let _3_v0 = _compr_1;
    if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("prkc")).length))))).contains(_3_v0)) {
      _coll1.add(_3_v0);
    }
  }
  return _coll1;
}()).length));
      if (_source0.is_DC5) {
        let _4___mcc_h0 = (_source0).cf5;
        let _5_cf5 = _4___mcc_h0;
        return _dafny.Set.fromElements(_5_cf5);
      } else {
        let _6___mcc_h1 = (_source0).cf4;
        let _7_cf4 = _6___mcc_h1;
        return _dafny.Set.fromElements(true);
      }
    };
    static fm18(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(946),new BigNumber((_dafny.Seq.of(new BigNumber(461))).length))).length))).cardinality()),new BigNumber((_dafny.Set.fromElements(true)).length)),true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("begfq")).length), new BigNumber((_dafny.Seq.UnicodeFromString("bcfe")).length), new BigNumber(-714))).length),false);
    };
    static fm19(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(false)));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D2.create_DC6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), function (_8_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}));
      if (_source1.is_DC7) {
        let _9___mcc_h0 = (_source1).cf7;
        let _10___mcc_h1 = (_source1).cf8;
        let _11_cf8 = _10___mcc_h1;
        let _12_cf7 = _9___mcc_h0;
        return new _dafny.CodePoint('y'.codePointAt(0));
      } else if (_source1.is_DC8) {
        let _13___mcc_h2 = (_source1).cf9;
        let _14_cf9 = _13___mcc_h2;
        return new _dafny.CodePoint('i'.codePointAt(0));
      } else if (_source1.is_DC6) {
        let _15___mcc_h3 = (_source1).cf6;
        let _16_cf6 = _15___mcc_h3;
        return new _dafny.CodePoint('o'.codePointAt(0));
      } else {
        let _17___mcc_h4 = (_source1).cf10;
        let _18_cf10 = _17___mcc_h4;
        if (true) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('m'.codePointAt(0));
        }
      }
    };
    static fm23(globalState) {
      return _module.D0.create_DC1();
    };
    static fm24(globalState) {
      if (_dafny.areEqual(_dafny.Seq.UnicodeFromString("shj"), _dafny.Seq.UnicodeFromString("wfuu"))) {
        return _module.D0.create_DC0(false);
      } else {
        return _module.D0.create_DC0(!(!(false)));
      }
    };
    static fm25(p0, p1, p2, p3, globalState) {
      if (!(_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).contains(false)) {
        return function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-970), new BigNumber(-460))) {
            let _19_v0 = _compr_2;
            if (((new BigNumber(-970)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(-460)))) {
              _coll2.push([(_19_v0).multipliedBy(new BigNumber(802)),_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("thvvevq")).length)), new BigNumber(-477), new BigNumber(803), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))]);
            }
          }
          return _coll2;
        }();
      } else {
        return function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(488), new BigNumber(518))) {
            let _20_v1 = _compr_3;
            if (((new BigNumber(488)).isLessThanOrEqualTo(_20_v1)) && ((_20_v1).isLessThan(new BigNumber(518)))) {
              _coll3.push([_module.__default.safeDivisionInt(_20_v1, new BigNumber(567)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-215)), function (_21_i0) {
                return new BigNumber(-734);
              })]);
            }
          }
          return _coll3;
        }();
      }
    };
    static fm26(globalState) {
      return _module.D1.create_DC5(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("ero"), new _dafny.CodePoint('f'.codePointAt(0))));
    };
    static fm27(p0, p1, p2, globalState) {
      return (function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(236)), function (_22_i0) {
          return new BigNumber(-977);
        })).Elements) {
          let _23_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(236)), function (_22_i0) {
            return new BigNumber(-977);
          }), _23_v0)) {
            _coll4.push([_module.__default.safeDivisionInt(_23_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tpiay")).length), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
              let _coll5 = new _dafny.Set();
              for (const _compr_5 of _dafny.IntegerRange(new BigNumber(14), new BigNumber(706))) {
                let _24_v1 = _compr_5;
                if (((new BigNumber(14)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(706)))) {
                  _coll5.add(_module.__default.safeModuloInt(_24_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), function (_25_i1) {
                    return new _dafny.CodePoint('x'.codePointAt(0));
                  })).length)));
                }
              }
              return _coll5;
            }()).length), new BigNumber(659))).length), new BigNumber(310))).length)),new BigNumber((_dafny.Set.fromElements(false, false, !(true))).length)]);
          }
        }
        return _coll4;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-232),new BigNumber(13)));
    };
    static fm28(p0, p1, p2, globalState) {
      let _source2 = _module.D21.create_DC46(true, false, new BigNumber(966));
      if (_source2.is_DC46) {
        let _26___mcc_h0 = (_source2).cf61;
        let _27___mcc_h1 = (_source2).cf62;
        let _28___mcc_h2 = (_source2).cf63;
        let _29_cf63 = _28___mcc_h2;
        let _30_cf62 = _27___mcc_h1;
        let _31_cf61 = _26___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(!(true),_29_cf63);
      } else if (_source2.is_DC47) {
        let _32___mcc_h3 = (_source2).cf64;
        let _33_cf64 = _32___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_33_cf64,new BigNumber((_dafny.Seq.UnicodeFromString("dgerly")).length));
      } else {
        let _34___mcc_h4 = (_source2).cf60;
        let _35_cf60 = _34___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(872));
      }
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("tfhxej")).length))).length)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(-270))).cardinality()))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-916), new BigNumber(191)), _dafny.Seq.of(new BigNumber(401), new BigNumber(4), new BigNumber((_dafny.Seq.UnicodeFromString("cwex")).length))));
    };
    static fm32(globalState) {
      if (false) {
        return (_dafny.MultiSet.fromElements(new BigNumber(-763))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("slnptiny"))).length), new BigNumber(872), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(467))).cardinality())));
      } else {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).cardinality()), new BigNumber(601));
      }
    };
    static fm33(globalState) {
      let _source3 = _module.D12.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(false,false));
      if (_source3.is_DC22) {
        let _36___mcc_h0 = (_source3).cf24;
        let _37___mcc_h1 = (_source3).cf25;
        let _38___mcc_h2 = (_source3).cf26;
        let _39___mcc_h3 = (_source3).cf27;
        let _40___mcc_h4 = (_source3).cf28;
        let _41_cf28 = _40___mcc_h4;
        let _42_cf27 = _39___mcc_h3;
        let _43_cf26 = _38___mcc_h2;
        let _44_cf25 = _37___mcc_h1;
        let _45_cf24 = _36___mcc_h0;
        return (function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(934))).length), new BigNumber((_dafny.Set.fromElements(_42_cf27)).length), new BigNumber(-580), new BigNumber((_41_cf28).length), (_dafny.ZERO).minus(new BigNumber(-570)))).length))).Elements) {
            let _46_v0 = _compr_6;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(934))).length), new BigNumber((_dafny.Set.fromElements(_42_cf27)).length), new BigNumber(-580), new BigNumber((_41_cf28).length), (_dafny.ZERO).minus(new BigNumber(-570)))).length))).contains(_46_v0)) {
              _coll6.push([_module.__default.safeDivisionInt(_46_v0, new BigNumber(-820)),_42_cf27]);
            }
          }
          return _coll6;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_41_cf28).length),false));
      } else if (_source3.is_DC23) {
        let _47___mcc_h5 = (_source3).cf29;
        let _48_cf29 = _47___mcc_h5;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-711),true);
      } else {
        let _49___mcc_h6 = (_source3).cf23;
        let _50_cf23 = _49___mcc_h6;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(34),true);
      }
    };
    static fm34(p0, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(false));
    };
    static fm35(p0, p1, p2, globalState) {
      return (function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(379), new BigNumber(-580))) {
          let _51_v0 = _compr_7;
          if (((new BigNumber(379)).isLessThanOrEqualTo(_51_v0)) && ((_51_v0).isLessThan(new BigNumber(-580)))) {
            _coll7.add(_module.__default.safeModuloInt(_51_v0, new BigNumber(624)));
          }
        }
        return _coll7;
      }()).Intersect(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-9), new BigNumber(42))) {
          let _52_v1 = _compr_8;
          if (((new BigNumber(-9)).isLessThanOrEqualTo(_52_v1)) && ((_52_v1).isLessThan(new BigNumber(42)))) {
            _coll8.add(_module.__default.safeModuloInt(_52_v1, new BigNumber(-429)));
          }
        }
        return _coll8;
      }());
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      if (false) {
        return _dafny.Seq.of(new BigNumber(-517), new BigNumber(898), new BigNumber((_dafny.Seq.of(!(true))).length));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(668)), _dafny.Seq.of(new BigNumber(452)));
      }
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-299)), _dafny.Seq.of(new BigNumber(-751)), _dafny.Seq.of(new BigNumber(854), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("iwgq")).length)),false)).length))), _dafny.Seq.of(new BigNumber(356))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(493), new BigNumber((_dafny.MultiSet.fromElements(!(!(false)))).cardinality()))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-902)))));
    };
    static fm39(p0, p1, p2, p3, globalState) {
      if (false) {
        return _dafny.Seq.of(false);
      } else {
        return _dafny.Seq.of(false, !(false));
      }
    };
    static fm40(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true, true, true),new BigNumber(649));
    };
    static fm41(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
    };
    static fm43(p0, p1, globalState) {
      return (_dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), function (_53_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length))))).Union((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(853), new BigNumber(3))) {
          let _54_v0 = _compr_9;
          if (((new BigNumber(853)).isLessThanOrEqualTo(_54_v0)) && ((_54_v0).isLessThan(new BigNumber(3)))) {
            _coll9.add(_module.__default.safeDivisionInt(_54_v0, new BigNumber((_dafny.Seq.UnicodeFromString("gsah")).length)));
          }
        }
        return _coll9;
      }()).Intersect(function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(-232), new BigNumber(274))) {
          let _55_v1 = _compr_10;
          if (((new BigNumber(-232)).isLessThanOrEqualTo(_55_v1)) && ((_55_v1).isLessThan(new BigNumber(274)))) {
            _coll10.add((_55_v1).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jynfus")).length)));
          }
        }
        return _coll10;
      }()));
    };
    static fm44(p0, globalState) {
      return _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)));
    };
    static fm45(globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(true), true, true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-220)), function (_56_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.of(new BigNumber(236))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(971), new BigNumber(-640), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(!(true), true))).cardinality()), new BigNumber((_dafny.Set.fromElements(false, true)).length))).cardinality())))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-804)), function (_57_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, true),new _dafny.CodePoint('x'.codePointAt(0)))).length);
      })))).Union(_dafny.MultiSet.fromElements(new BigNumber(473), new BigNumber((_dafny.Seq.UnicodeFromString("hmbwmks")).length)));
    };
    static fm46(p0, globalState) {
      if (!(((false) ? (false) : (!(false))))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true));
      }
    };
    static fm47(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(305)), function (_58_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length),_dafny.Map.Empty.slice().updateUnsafe(false,true));
    };
    static fm48(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(!(!(true)), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("f"), _dafny.Seq.UnicodeFromString("ju")), _dafny.Seq.contains(_dafny.Seq.of(false, !(true)), !(!(true))));
    };
    static fm49(globalState) {
      if (true) {
        return _dafny.Set.fromElements(true);
      } else if (!(true)) {
        return _dafny.Set.fromElements(true, !(false), true, true, true);
      } else {
        return _dafny.Set.fromElements(true, true);
      }
    };
    static fm50(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(247), new BigNumber(643))) {
          let _59_v0 = _compr_11;
          if (((new BigNumber(247)).isLessThanOrEqualTo(_59_v0)) && ((_59_v0).isLessThan(new BigNumber(643)))) {
            _coll11.push([(_59_v0).plus(new BigNumber(454)),!(true)]);
          }
        }
        return _coll11;
      }());
    };
    static fm51(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)));
    };
    static fm52(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(398)), function (_60_i0) {
        return new BigNumber(-816);
      }), (_module.D6.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_61_i1) {
  return new BigNumber(247);
}))).dtor_cf14, _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-873))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xmsk")).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_62_i2) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(262),new BigNumber((_dafny.MultiSet.fromElements(!(!(true)), true)).cardinality()))).length);
      }));
    };
    static fm53(p0, p1, globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm54(globalState) {
      return _module.D15.create_DC28(false, !(!(!((new BigNumber(48)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D19.create_DC41(_module.D19.create_DC38(function () {
  let _coll12 = new _dafny.Set();
  for (const _compr_12 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jwfty"))).Elements) {
    let _63_v0 = _compr_12;
    if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jwfty"))).contains(_63_v0)) {
      _coll12.add(_63_v0);
    }
  }
  return _coll12;
}())), _module.D19.create_DC41(_module.D19.create_DC41(_module.D19.create_DC40(true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wkltshb")).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(467),_dafny.Seq.of(new BigNumber(111), new BigNumber(871))))))),true)).length))))));
    };
    static fm55(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(479),_dafny.Seq.of(true))).Merge(function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Seq.of(new BigNumber((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true))).cardinality()))).Elements) {
            let _64_v1 = _compr_14;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true))).cardinality()))).contains(_64_v1)) {
              _coll14.push([_module.__default.safeModuloInt(_64_v1, new BigNumber(67)),true]);
            }
          }
          return _coll14;
        }()).length))).Elements) {
          let _65_v0 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true))).cardinality()))).Elements) {
              let _64_v1 = _compr_15;
              if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true))).cardinality()))).contains(_64_v1)) {
                _coll15.push([_module.__default.safeModuloInt(_64_v1, new BigNumber(67)),true]);
              }
            }
            return _coll15;
          }()).length)), _65_v0)) {
            _coll13.push([(_65_v0).minus(new BigNumber(520)),_dafny.Seq.of(false)]);
          }
        }
        return _coll13;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(979)), function (_66_i0) {
        return _module.D16.create_DC30(_dafny.Seq.of(_dafny.Set.fromElements(false, false)));
      })).length),_dafny.Seq.of(true, true, true)));
    };
    static fm56(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D12.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(false,!(false))))).length)), _dafny.Seq.of(new BigNumber((function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).Elements) {
          let _67_v0 = _compr_16;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).contains(_67_v0)) {
            _coll16.add(_67_v0);
          }
        }
        return _coll16;
      }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(373))).cardinality())))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(829)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(361)), function (_68_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("xibg")).length);
      })));
    };
    static fm57(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(350))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(81))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("eauvvtm"))).length))));
    };
    static fm58(globalState) {
      if ((_module.D26.create_DC57(true)).dtor_cf77) {
        return _module.D11.create_DC20(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ugle")).length),new BigNumber(317))).length));
      } else if (!(false)) {
        return _module.D11.create_DC20(false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))).length), new BigNumber(711))).length));
      } else {
        return _module.D11.create_DC20(true, new BigNumber(-365));
      }
    };
    static fm59(globalState) {
      let _source4 = _module.D0.create_DC0((_module.D26.create_DC57(!(true))).dtor_cf77);
      if (_source4.is_DC1) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),true)).Merge(function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).Elements) {
            let _69_v0 = _compr_17;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))), _69_v0)) {
              _coll17.push([_69_v0,!(true)]);
            }
          }
          return _coll17;
        }());
      } else if (_source4.is_DC2) {
        let _70___mcc_h0 = (_source4).cf1;
        let _71___mcc_h1 = (_source4).cf2;
        let _72_cf2 = _71___mcc_h1;
        let _73_cf1 = _70___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_72_cf2,false));
      } else if (_source4.is_DC0) {
        let _74___mcc_h2 = (_source4).cf0;
        let _75_cf0 = _74___mcc_h2;
        return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),_75_cf0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_75_cf0));
      } else {
        let _76___mcc_h3 = (_source4).cf3;
        let _77_cf3 = _76___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),true);
      }
    };
    static fm60(globalState) {
      return _module.D14.create_DC26((true) && (true), ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("oxkxihvnx")).length)))).plus(new BigNumber(113)), new _dafny.CodePoint('v'.codePointAt(0)));
    };
    static fm61(p0, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jemqeaufg"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), function (_78_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("oyqd")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("e"), _dafny.Seq.UnicodeFromString("yk")), _dafny.Seq.UnicodeFromString("acxl"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kqlctexte"), _dafny.Seq.UnicodeFromString("alohcbeea")));
    };
    static fm62(p0, p1, p2, p3, globalState) {
      let _source5 = _module.D11.create_DC19(new _dafny.CodePoint('g'.codePointAt(0)));
      if (_source5.is_DC20) {
        let _79___mcc_h0 = (_source5).cf21;
        let _80___mcc_h1 = (_source5).cf22;
        let _81_cf22 = _80___mcc_h1;
        let _82_cf21 = _79___mcc_h0;
        return _module.D2.create_DC8(_82_cf21);
      } else {
        let _83___mcc_h2 = (_source5).cf20;
        let _84_cf20 = _83___mcc_h2;
        return _module.D2.create_DC8(false);
      }
    };
    static fm63(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(736)), function (_85_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }));
    };
    static fm64(p0, globalState) {
      return _module.D6.create_DC14((true) === (false));
    };
    static fm65(p0, p1, p2, globalState) {
      return _module.D17.create_DC34(new _dafny.CodePoint('r'.codePointAt(0)));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _module.D0.Default();
      let r1 = false;
      let _86_v0;
      _86_v0 = new BigNumber(842);
      let _87_v1;
      _87_v1 = _dafny.Seq.of(_86_v0);
      let _88_v2;
      _88_v2 = _dafny.Map.Empty.slice().updateUnsafe(_86_v0,_87_v1);
      let _89_v3;
      _89_v3 = _module.D19.create_DC40(!((!(p2)) === (p2)), _86_v0, _88_v2);
      _89_v3 = _89_v3;
      (globalState).f1 = _module.__default.safeModuloInt(_86_v0, new BigNumber(239));
      let _90_i0;
      _90_i0 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm0(p2, globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_90_i0)) {
              break L0;
            }
            _90_i0 = (_90_i0).plus(_dafny.ONE);
            let _91_v4;
            let _nw0 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _91_v4 = _nw0;
            let _92_v5;
            _92_v5 = _dafny.Set.fromElements(_91_v4);
            let _93_v6;
            _93_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_92_v5).length),_module.D21.create_DC47(p2));
            let _94_v7;
            _94_v7 = _dafny.Map.Empty.slice().updateUnsafe(_93_v6,p2);
            let _95_v8;
            _95_v8 = _dafny.Map.Empty.slice().updateUnsafe(p3,_93_v6);
            r1 = (((_94_v7).contains((((_95_v8).contains(p3)) ? ((_95_v8).get(p3)) : (_93_v6)))) ? ((_94_v7).get((((_95_v8).contains(p3)) ? ((_95_v8).get(p3)) : (_93_v6)))) : (!(p2)));
            let _96_v9;
            _96_v9 = _dafny.Seq.UnicodeFromString("ioa");
            let _rhs0 = (_dafny.ZERO).minus(_module.__default.fm3(p2, p2, globalState));
            let _rhs1 = p2;
            let _rhs2 = (_87_v1)[_module.__default.safeIndex((_86_v0).plus(new BigNumber(-259)), new BigNumber((_87_v1).length))];
            let _rhs3 = p2;
            let _rhs4 = _96_v9;
            let _lhs0 = globalState;
            let _lhs1 = globalState;
            _lhs0.f1 = _rhs0;
            r1 = _rhs1;
            _lhs1.f1 = _rhs2;
            r1 = _rhs3;
            _96_v9 = _rhs4;
            (globalState).f1 = _module.__default.safeDivisionInt(_86_v0, _86_v0);
            r1 = p2;
          }
        }
      }
      (globalState).f1 = (_dafny.ZERO).minus((_86_v0).minus(_86_v0));
      let _97_v10;
      _97_v10 = _module.D21.create_DC46(p2, p2, _86_v0);
      (globalState).f1 = new BigNumber((function (_source6) {
        if (_source6.is_DC46) {
          let _98___mcc_h0 = (_source6).cf61;
          let _99___mcc_h1 = (_source6).cf62;
          let _100___mcc_h2 = (_source6).cf63;
          let _101_cf63 = _100___mcc_h2;
          let _102_cf62 = _99___mcc_h1;
          let _103_cf61 = _98___mcc_h0;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("irn"), _dafny.Seq.UnicodeFromString("vcyxtburc"));
        } else if (_source6.is_DC47) {
          let _104___mcc_h3 = (_source6).cf64;
          let _105_cf64 = _104___mcc_h3;
          return _dafny.Seq.UnicodeFromString("d");
        } else {
          let _106___mcc_h4 = (_source6).cf60;
          let _107_cf60 = _106___mcc_h4;
          return _dafny.Seq.UnicodeFromString("oesqkuq");
        }
      }(_97_v10)).length);
      let _108_i1;
      _108_i1 = _dafny.ZERO;
      L1: {
        while (!(_86_v0).isEqualTo(_86_v0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_108_i1)) {
              break L1;
            }
            _108_i1 = (_108_i1).plus(_dafny.ONE);
            let _109_v11;
            let _nw1 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
            _109_v11 = _nw1;
            let _110_v12;
            _110_v12 = _dafny.MultiSet.fromElements(_86_v0, _86_v0);
            let _index0 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length));
            (_109_v11)[_index0] = (_110_v12).Union((_110_v12).update(new BigNumber(145), _module.__default.abs(_86_v0)));
            let _index1 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length));
            (_109_v11)[_index1] = _110_v12;
            if (p2) {
              let _111_v13;
              _111_v13 = _dafny.Seq.of(true, p2);
              let _112_v14;
              _112_v14 = _dafny.Map.Empty.slice().updateUnsafe((_111_v13)[_module.__default.safeIndex(_86_v0, new BigNumber((_111_v13).length))],p2);
              r1 = (_111_v13)[_module.__default.safeIndex(new BigNumber((_112_v14).length), new BigNumber((_111_v13).length))];
              let _113_v15;
              _113_v15 = _module.D11.create_DC20(p2, new BigNumber(((_109_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length))]).cardinality()));
              let _114_v16;
              _114_v16 = _dafny.Map.Empty.slice().updateUnsafe((_113_v15).dtor_cf21,p0);
              let _115_v17;
              _115_v17 = _dafny.Set.fromElements(_86_v0, _86_v0);
              let _116_v18;
              _116_v18 = _dafny.Map.Empty.slice().updateUnsafe(_115_v17,p0);
              let _117_v19;
              let _nw2 = Array((new BigNumber(25)).toNumber());
              _nw2[(_dafny.ZERO).toNumber()] = p0;
              _nw2[(_dafny.ONE).toNumber()] = p0;
              _nw2[(new BigNumber(2)).toNumber()] = (((_114_v16).contains(!(p2))) ? ((_114_v16).get(!(p2))) : (p0));
              _nw2[(new BigNumber(3)).toNumber()] = p0;
              _nw2[(new BigNumber(4)).toNumber()] = p0;
              _nw2[(new BigNumber(5)).toNumber()] = p0;
              _nw2[(new BigNumber(6)).toNumber()] = p0;
              _nw2[(new BigNumber(7)).toNumber()] = p0;
              _nw2[(new BigNumber(8)).toNumber()] = p0;
              _nw2[(new BigNumber(9)).toNumber()] = p0;
              _nw2[(new BigNumber(10)).toNumber()] = p0;
              _nw2[(new BigNumber(11)).toNumber()] = p0;
              _nw2[(new BigNumber(12)).toNumber()] = p0;
              _nw2[(new BigNumber(13)).toNumber()] = p0;
              _nw2[(new BigNumber(14)).toNumber()] = p0;
              _nw2[(new BigNumber(15)).toNumber()] = p0;
              _nw2[(new BigNumber(16)).toNumber()] = p0;
              _nw2[(new BigNumber(17)).toNumber()] = (((_116_v18).contains(_dafny.Set.fromElements(new BigNumber(148), _86_v0))) ? ((_116_v18).get(_dafny.Set.fromElements(new BigNumber(148), _86_v0))) : (p0));
              _nw2[(new BigNumber(18)).toNumber()] = ((p2) ? (p0) : (p0));
              _nw2[(new BigNumber(19)).toNumber()] = p0;
              _nw2[(new BigNumber(20)).toNumber()] = p0;
              _nw2[(new BigNumber(21)).toNumber()] = p0;
              _nw2[(new BigNumber(22)).toNumber()] = p0;
              _nw2[(new BigNumber(23)).toNumber()] = p0;
              _nw2[(new BigNumber(24)).toNumber()] = p0;
              _117_v19 = _nw2;
              let _index2 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length));
              (_117_v19)[_index2] = p0;
              let _index3 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length));
              (_117_v19)[_index3] = p0;
              let _118_v20;
              let _init0 = ((_119_v0) => function (_120_i2) {
                return (_120_i2).plus(_119_v0);
              })(_86_v0);
              let _nw3 = Array((new BigNumber(28)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
                _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _118_v20 = _nw3;
              let _121_v21;
              let _nw4 = Array((new BigNumber(22)).toNumber());
              _nw4[(_dafny.ZERO).toNumber()] = _118_v20;
              _nw4[(_dafny.ONE).toNumber()] = _118_v20;
              _nw4[(new BigNumber(2)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(3)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(4)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(5)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(6)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(7)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(8)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(9)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(10)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(11)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(12)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(13)).toNumber()] = (_118_v20);
              _nw4[(new BigNumber(14)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(15)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(16)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(17)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(18)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(19)).toNumber()] = _118_v20;
              _nw4[(new BigNumber(20)).toNumber()] = ((p2) ? (_118_v20) : (_118_v20));
              _nw4[(new BigNumber(21)).toNumber()] = _118_v20;
              _121_v21 = _nw4;
              let _index4 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_121_v21).length));
              (_121_v21)[_index4] = _118_v20;
              let _index5 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_121_v21).length));
              (_121_v21)[_index5] = _118_v20;
              let _122_v22;
              _122_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
              _122_v22 = (_122_v22).update((_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))], p2);
              let _arr0 = (_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))];
              let _index6 = _module.__default.safeIndex(new BigNumber(943), new BigNumber(((_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))]).length));
              _arr0[_index6] = p2;
              let _arr1 = (_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))];
              let _index7 = _module.__default.safeIndex(new BigNumber(943), new BigNumber(((_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))]).length));
              let _rhs5 = (_dafny.ZERO).minus(_86_v0);
              let _rhs6 = (new BigNumber(-243)).minus(_86_v0);
              let _rhs7 = true;
              let _rhs8 = (_111_v13)[_module.__default.safeIndex((_dafny.ZERO).minus(_86_v0), new BigNumber((_111_v13).length))];
              let _lhs2 = globalState;
              let _lhs3 = (_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))];
              let _lhs4 = _module.__default.safeIndex(new BigNumber(943), new BigNumber(((_117_v19)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_117_v19).length))]).length));
              _lhs2.f1 = _rhs5;
              _86_v0 = _rhs6;
              _lhs3[_lhs4] = _rhs7;
              r1 = _rhs8;
            } else {
              (globalState).f1 = (_dafny.ZERO).minus(_86_v0);
              let _123_v23;
              let _init1 = ((_124_p2) => function (_125_i3) {
                return _dafny.Seq.Concat(_dafny.Seq.of(_124_p2, _124_p2), _dafny.Seq.of(true, _124_p2));
              })(p2);
              let _nw5 = Array((new BigNumber(19)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
                _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _123_v23 = _nw5;
              let _126_v24;
              _126_v24 = _dafny.Seq.of(p2, p2);
              let _127_v25;
              _127_v25 = _dafny.Seq.of((_126_v24)[_module.__default.safeIndex(new BigNumber((_87_v1).length), new BigNumber((_126_v24).length))], p2);
              let _index8 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_123_v23).length));
              (_123_v23)[_index8] = _126_v24;
              let _index9 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_123_v23).length));
              (_123_v23)[_index9] = _dafny.Seq.Concat(_127_v25, _126_v24);
              let _128_v26;
              _128_v26 = _module.D14.create_DC26(true, _86_v0, p3);
              let _129_v27;
              let _nw6 = new _module.C7();
              _nw6.__ctor((new BigNumber(112)).plus(_86_v0), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("oabcxl")).length), _86_v0), p3, (_128_v26).dtor_cf32);
              _129_v27 = _nw6;
              let _130_v28;
              _130_v28 = _dafny.Set.fromElements((_109_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length))], _110_v12, (_109_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length))]);
              r1 = ((_130_v28).Union(_130_v28)).equals((_130_v28).Union(_130_v28));
              let _131_v29;
              _131_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-444),_129_v27.f10);
              let _132_v30;
              _132_v30 = _dafny.Set.fromElements(p2, !(p2));
              let _133_v31;
              _133_v31 = _dafny.Seq.UnicodeFromString("ycbrppf");
              let _134_v32;
              _134_v32 = _dafny.MultiSet.fromElements(_133_v31);
              let _135_v33;
              _135_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_136_i4) {
                return new _dafny.CodePoint('j'.codePointAt(0));
              }));
              let _137_v34;
              let _nw7 = Array((new BigNumber(18)).toNumber());
              _nw7[(_dafny.ZERO).toNumber()] = _86_v0;
              _nw7[(_dafny.ONE).toNumber()] = _86_v0;
              _nw7[(new BigNumber(2)).toNumber()] = _129_v27.f10;
              _nw7[(new BigNumber(3)).toNumber()] = new BigNumber((_131_v29).length);
              _nw7[(new BigNumber(4)).toNumber()] = (_129_v27).f9;
              _nw7[(new BigNumber(5)).toNumber()] = _86_v0;
              _nw7[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.of(_129_v27.f10, new BigNumber((_132_v30).length), new BigNumber((_134_v32).cardinality()), new BigNumber((_135_v33).length))).length);
              _nw7[(new BigNumber(7)).toNumber()] = _129_v27.f10;
              _nw7[(new BigNumber(8)).toNumber()] = new BigNumber(552);
              _nw7[(new BigNumber(9)).toNumber()] = _86_v0;
              _nw7[(new BigNumber(10)).toNumber()] = (_129_v27).f9;
              _nw7[(new BigNumber(11)).toNumber()] = _86_v0;
              _nw7[(new BigNumber(12)).toNumber()] = _86_v0;
              _nw7[(new BigNumber(13)).toNumber()] = new BigNumber(60);
              _nw7[(new BigNumber(14)).toNumber()] = (_129_v27).f9;
              _nw7[(new BigNumber(15)).toNumber()] = new BigNumber(((_109_v11)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_109_v11).length))]).cardinality());
              _nw7[(new BigNumber(16)).toNumber()] = (_129_v27).f9;
              _nw7[(new BigNumber(17)).toNumber()] = _129_v27.f10;
              _137_v34 = _nw7;
              let _138_v35;
              _138_v35 = _dafny.Map.Empty.slice().updateUnsafe(_137_v34,(_129_v27).f9);
              let _139_v36;
              _139_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
              _138_v35 = (_138_v35).update(_137_v34, (_129_v27.f10).minus(new BigNumber((_139_v36).length)));
            }
            let _140_v37;
            _140_v37 = _dafny.Seq.of(p2);
            let _141_v38;
            _141_v38 = _dafny.Seq.of(p2, p2, (_140_v37)[_module.__default.safeIndex(_86_v0, new BigNumber((_140_v37).length))]);
            _86_v0 = ((((p2) ? (p2) : (p2))) ? (new BigNumber((_dafny.Seq.Concat(_140_v37, _141_v38)).length)) : (_86_v0));
            if (((_86_v0).minus(new BigNumber(-634))).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_87_v1)[_module.__default.safeIndex(_86_v0, new BigNumber((_87_v1).length))], _86_v0))) {
              let _142_v39;
              let _nw8 = Array((new BigNumber(6)).toNumber());
              _nw8[(_dafny.ZERO).toNumber()] = p1;
              _nw8[(_dafny.ONE).toNumber()] = p1;
              _nw8[(new BigNumber(2)).toNumber()] = p3;
              _nw8[(new BigNumber(3)).toNumber()] = p1;
              _nw8[(new BigNumber(4)).toNumber()] = p1;
              _nw8[(new BigNumber(5)).toNumber()] = p1;
              _142_v39 = _nw8;
              let _143_v40;
              _143_v40 = _dafny.MultiSet.fromElements(_142_v39, _142_v39);
              let _144_v41;
              let _nw9 = new _module.C7();
              _nw9.__ctor(_86_v0, _86_v0, p3, p2);
              _144_v41 = _nw9;
              let _145_v42;
              _145_v42 = _dafny.Map.Empty.slice().updateUnsafe(p2,_144_v41);
              let _146_v43;
              _146_v43 = _dafny.Map.Empty.slice().updateUnsafe(_143_v40,(_145_v42).Merge(_145_v42));
              let _147_v44;
              _147_v44 = _dafny.Seq.of(_145_v42, _145_v42, _145_v42);
              _146_v43 = (_146_v43).update((_143_v40).Union(_143_v40), (_147_v44)[_module.__default.safeIndex(_86_v0, new BigNumber((_147_v44).length))]);
              let _148_v45;
              let _nw10 = Array((new BigNumber(18)).toNumber()).fill([]);
              _148_v45 = _nw10;
              let _index10 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              (p0)[_index10] = !(new BigNumber(865)).isEqualTo(_86_v0);
              let _149_v46;
              _149_v46 = _module.D26.create_DC55(_148_v45);
              let _index11 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _rhs9 = (_149_v46).dtor_cf74;
              let _rhs10 = p2;
              let _lhs5 = p0;
              let _lhs6 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              _148_v45 = _rhs9;
              _lhs5[_lhs6] = _rhs10;
              let _150_v47;
              _150_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,_86_v0);
              let _151_v48;
              _151_v48 = _dafny.Set.fromElements(new BigNumber(642));
              let _152_v49;
              _152_v49 = _dafny.Set.fromElements(_151_v48, _151_v48);
              let _index12 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _index13 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _rhs11 = (_150_v47).update(p3, (_dafny.ZERO).minus(_86_v0));
              let _rhs12 = (_144_v41).f6;
              let _rhs13 = (_152_v49).IsProperSubsetOf(_dafny.Set.fromElements(_151_v48));
              let _lhs7 = p0;
              let _lhs8 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _lhs9 = p0;
              let _lhs10 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              _150_v47 = _rhs11;
              _lhs7[_lhs8] = _rhs12;
              _lhs9[_lhs10] = _rhs13;
              let _153_v50;
              _153_v50 = _dafny.Map.Empty.slice().updateUnsafe((_144_v41).f6,_87_v1);
              let _154_v51;
              _154_v51 = _dafny.MultiSet.fromElements(p2);
              let _155_v52;
              _155_v52 = _dafny.Seq.UnicodeFromString("swp");
              _153_v50 = (_153_v50).update((new BigNumber((_155_v52).length)).isLessThan(new BigNumber((_154_v51).cardinality())), _87_v1);
              let _156_v53;
              _156_v53 = _dafny.MultiSet.fromElements(_155_v52);
              let _157_v54;
              let _init2 = ((_158_v41, _159_v0) => function (_160_i5) {
                return (_160_i5).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_158_v41).f6,_159_v0)).length));
              })(_144_v41, _86_v0);
              let _nw11 = Array((new BigNumber(19)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
                _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _157_v54 = _nw11;
              let _161_v55;
              _161_v55 = _dafny.Map.Empty.slice().updateUnsafe(_157_v54,_156_v53);
              let _index14 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _index15 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _rhs14 = !(!((_144_v41).f6));
              let _rhs15 = !((((p0)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length))]) ? (_156_v53) : ((((_161_v55).contains(_157_v54)) ? ((_161_v55).get(_157_v54)) : (_dafny.MultiSet.fromElements(_155_v52)))))).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tyay"), _155_v52));
              let _rhs16 = _86_v0;
              let _rhs17 = ((_87_v1)[_module.__default.safeIndex(_86_v0, new BigNumber((_87_v1).length))]).minus(_86_v0);
              let _rhs18 = _module.__default.fm0(!(p2), globalState);
              let _lhs11 = p0;
              let _lhs12 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              let _lhs13 = globalState;
              let _lhs14 = globalState;
              let _lhs15 = p0;
              let _lhs16 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p0).length));
              r1 = _rhs14;
              _lhs11[_lhs12] = _rhs15;
              _lhs13.f1 = _rhs16;
              _lhs14.f1 = _rhs17;
              _lhs15[_lhs16] = _rhs18;
            } else {
              let _162_v56;
              _162_v56 = _dafny.Seq.UnicodeFromString("xarvj");
              _162_v56 = _162_v56;
              let _163_v57;
              let _nw12 = new _module.C4();
              _nw12.__ctor();
              _163_v57 = _nw12;
              let _index16 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
              (p0)[_index16] = p2;
              let _index17 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
              (p0)[_index17] = true;
              (globalState).f1 = (((p0)[_module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length))]) ? (_86_v0) : (_86_v0));
              _86_v0 = (_dafny.ZERO).minus(_86_v0);
            }
          }
        }
      }
      let _164_v58;
      _164_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
      let _165_v59;
      _165_v59 = _module.D0.create_DC2((_164_v58).update(p0, p3), p1);
      r0 = _165_v59;
      let _166_v60;
      _166_v60 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_87_v1));
      r1 = ((_166_v60).equals(_166_v60)) === (_module.__default.fm0(p2, globalState));
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _167_v0;
      _167_v0 = new BigNumber(647);
      let _168_v1;
      _168_v1 = _dafny.Seq.of(_167_v0, _167_v0);
      let _169_v2;
      _169_v2 = _dafny.Seq.of(_168_v1);
      let _170_v3;
      _170_v3 = true;
      let _171_v4;
      _171_v4 = _module.D0.create_DC0(_170_v3);
      let _172_v5;
      _172_v5 = _dafny.Seq.of(_170_v3, _170_v3, true, _170_v3, _170_v3);
      let _173_globalState;
      let _nw13 = new _module.GlobalState();
      _nw13.__ctor(new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(10), new BigNumber(430), _169_v2, _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_170_v3, (_171_v4).dtor_cf0), _172_v5)));
      _173_globalState = _nw13;
      let _174_i0;
      _174_i0 = _dafny.ZERO;
      L2: {
        while (_170_v3) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_174_i0)) {
              break L2;
            }
            _174_i0 = (_174_i0).plus(_dafny.ONE);
            let _175_v6;
            let _init3 = ((_176_v3) => function (_177_i1) {
              return _176_v3;
            })(_170_v3);
            let _nw14 = Array((new BigNumber(16)).toNumber());
            for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw14.length); _i0_3++) {
              _nw14[_i0_3] = _init3(new BigNumber(_i0_3));
            }
            _175_v6 = _nw14;
            let _178_v7;
            _178_v7 = new _dafny.CodePoint('w'.codePointAt(0));
            let _179_v8;
            let _180_v9;
            let _out0;
            let _out1;
            let _outcollector0 = _module.__default.m0(_175_v6, _178_v7, _170_v3, _178_v7, _173_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _179_v8 = _out0;
            _180_v9 = _out1;
            let _181_v10;
            let _nw15 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
            _181_v10 = _nw15;
            let _index18 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_181_v10).length));
            (_181_v10)[_index18] = _167_v0;
            let _index19 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_181_v10).length));
            (_181_v10)[_index19] = (_167_v0).plus(((_180_v9) ? (_167_v0) : (_167_v0)));
            let _182_v11;
            _182_v11 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_179_v8);
            let _183_v12;
            _183_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(325),_167_v0);
            let _184_v13;
            _184_v13 = _dafny.Map.Empty.slice().updateUnsafe(_183_v12,_182_v11);
            let _185_v14;
            _185_v14 = _dafny.Map.Empty.slice().updateUnsafe(((_181_v10)[_module.__default.safeIndex(new BigNumber(975), new BigNumber((_181_v10).length))]).multipliedBy(_167_v0),_dafny.Set.fromElements(_182_v11, (((_184_v13).contains(_183_v12)) ? ((_184_v13).get(_183_v12)) : (_182_v11)), _182_v11));
            let _186_v15;
            _186_v15 = _dafny.Set.fromElements(_182_v11);
            _185_v14 = (_185_v14).update(_module.__default.safeDivisionInt((_181_v10)[_module.__default.safeIndex(new BigNumber(975), new BigNumber((_181_v10).length))], (_181_v10)[_module.__default.safeIndex(new BigNumber(975), new BigNumber((_181_v10).length))]), _186_v15);
            let _187_v16;
            _187_v16 = _dafny.MultiSet.fromElements(_167_v0, new BigNumber(493));
            let _188_v17;
            _188_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_170_v3, _173_globalState),_module.__default.safeDivisionInt(new BigNumber(379), new BigNumber((_187_v16).cardinality())));
            _188_v17 = (_188_v17).update(_170_v3, _167_v0);
          }
        }
      }
      let _189_v18;
      _189_v18 = _module.D0.create_DC3(_module.D0.create_DC0((_172_v5)[_module.__default.safeIndex(_167_v0, new BigNumber((_172_v5).length))]));
      _189_v18 = _189_v18;
      let _190_v19;
      let _nw16 = Array((new BigNumber(7)).toNumber());
      _nw16[(_dafny.ZERO).toNumber()] = _170_v3;
      _nw16[(_dafny.ONE).toNumber()] = true;
      _nw16[(new BigNumber(2)).toNumber()] = _170_v3;
      _nw16[(new BigNumber(3)).toNumber()] = _170_v3;
      _nw16[(new BigNumber(4)).toNumber()] = _170_v3;
      _nw16[(new BigNumber(5)).toNumber()] = _170_v3;
      _nw16[(new BigNumber(6)).toNumber()] = _170_v3;
      _190_v19 = _nw16;
      let _191_v20;
      _191_v20 = new _dafny.CodePoint('j'.codePointAt(0));
      let _192_v21;
      let _193_v22;
      let _out2;
      let _out3;
      let _outcollector1 = _module.__default.m0(_190_v19, _191_v20, _170_v3, _191_v20, _173_globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _192_v21 = _out2;
      _193_v22 = _out3;
      let _194_v23;
      _194_v23 = _dafny.Map.Empty.slice().updateUnsafe(_191_v20,_167_v0);
      _194_v23 = (_194_v23).Merge((_194_v23).Merge(_194_v23));
      _191_v20 = _191_v20;
      let _195_v24;
      let _nw17 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _195_v24 = _nw17;
      let _index20 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
      (_195_v24)[_index20] = _dafny.Seq.UnicodeFromString("ebdgugv");
      let _196_v25;
      _196_v25 = _dafny.Seq.UnicodeFromString("ds");
      let _index21 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
      (_195_v24)[_index21] = _196_v25;
      let _197_v26;
      _197_v26 = _dafny.Map.Empty.slice().updateUnsafe(_190_v19,_191_v20);
      let _198_v27;
      let _nw18 = Array((new BigNumber(4)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = _192_v21;
      _nw18[(_dafny.ONE).toNumber()] = _module.D0.create_DC2(_197_v26, _191_v20);
      _nw18[(new BigNumber(2)).toNumber()] = _192_v21;
      _nw18[(new BigNumber(3)).toNumber()] = _192_v21;
      _198_v27 = _nw18;
      let _199_v28;
      _199_v28 = _dafny.Seq.of(_198_v27);
      _198_v27 = (_199_v28)[_module.__default.safeIndex(_167_v0, new BigNumber((_199_v28).length))];
      if (_193_v22) {
        let _index22 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
        (_190_v19)[_index22] = _170_v3;
        let _index23 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
        (_190_v19)[_index23] = !(_dafny.areEqual((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))], _module.__default.fm1(_167_v0, (_172_v5)[_module.__default.safeIndex(_167_v0, new BigNumber((_172_v5).length))], new BigNumber(606), _173_globalState)));
        let _200_v29;
        let _nw19 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _200_v29 = _nw19;
        _200_v29 = _200_v29;
        let _201_v31;
        _201_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_170_v3, _193_v22, _173_globalState),_168_v1);
        let _202_v32;
        _202_v32 = _dafny.Map.Empty.slice().updateUnsafe(_168_v1,(((_201_v31).contains(_dafny.Seq.of((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))], (_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]))) ? ((_201_v31).get(_dafny.Seq.of((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))], (_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]))) : ((_169_v2)[_module.__default.safeIndex(new BigNumber(((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]).length), new BigNumber((_169_v2).length))])));
        let _203_v33;
        _203_v33 = _dafny.MultiSet.fromElements(_191_v20, _191_v20);
        let _204_v34;
        _204_v34 = _dafny.Map.Empty.slice().updateUnsafe((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))],_203_v33);
        let _205_v35;
        _205_v35 = _module.D1.create_DC4(new BigNumber((_204_v34).length));
        let _206_v36;
        _206_v36 = _module.D1.create_DC5(!((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]));
        let _rhs19 = new BigNumber(-242);
        let _rhs20 = (new BigNumber((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of ((_202_v32).update(_168_v1, _168_v1)).Keys.Elements) {
            let _207_v30 = _compr_18;
            if (((_202_v32).update(_168_v1, _168_v1)).contains(_207_v30)) {
              _coll18.push([_207_v30,_193_v22]);
            }
          }
          return _coll18;
        }()).length)).plus((_205_v35).dtor_cf4);
        let _rhs21 = (_206_v36).dtor_cf5;
        let _lhs17 = _173_globalState;
        _lhs17.f1 = _rhs19;
        _167_v0 = _rhs20;
        _170_v3 = _rhs21;
        let _208_v37;
        _208_v37 = _dafny.Set.fromElements((_dafny.ZERO).minus((((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]) ? (new BigNumber(370)) : (_167_v0))), _167_v0, _167_v0, _167_v0);
        let _rhs22 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_169_v2)[_module.__default.safeIndex(_167_v0, new BigNumber((_169_v2).length))], _168_v1)).length))).plus(_module.__default.fm3((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))], _170_v3, _173_globalState));
        let _rhs23 = (function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of _dafny.IntegerRange(new BigNumber(270), new BigNumber(-534))) {
            let _209_v38 = _compr_19;
            if (((new BigNumber(270)).isLessThanOrEqualTo(_209_v38)) && ((_209_v38).isLessThan(new BigNumber(-534)))) {
              _coll19.add((_209_v38).multipliedBy(_167_v0));
            }
          }
          return _coll19;
        }()).Union(_208_v37);
        let _lhs18 = _173_globalState;
        _lhs18.f1 = _rhs22;
        _208_v37 = _rhs23;
        let _210_v39;
        _210_v39 = _module.D2.create_DC6((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]);
        if (!_dafny.areEqual(_dafny.Seq.Concat(_196_v25, (_210_v39).dtor_cf6), _196_v25)) {
          let _211_v40;
          _211_v40 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,(_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]);
          let _212_v41;
          _212_v41 = _dafny.Map.Empty.slice().updateUnsafe(_170_v3,(((_211_v40).contains(_167_v0)) ? ((_211_v40).get(_167_v0)) : (_170_v3)));
          let _213_v42;
          let _nw20 = new _module.C7();
          _nw20.__ctor(_167_v0, new BigNumber(((_212_v41).Merge(_212_v41)).length), _191_v20, (((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]) ? (_170_v3) : ((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))])));
          _213_v42 = _nw20;
          let _index24 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_200_v29).length));
          (_200_v29)[_index24] = (_213_v42).f9;
          let _index25 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_200_v29).length));
          (_200_v29)[_index25] = _167_v0;
          let _214_v43;
          let _nw21 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
          _214_v43 = _nw21;
          let _index26 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_214_v43).length));
          (_214_v43)[_index26] = _module.__default.fm2(_193_v22, !(_170_v3), _173_globalState);
          let _index27 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_214_v43).length));
          (_214_v43)[_index27] = _dafny.Seq.Concat(_dafny.Seq.of((_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))]), _172_v5);
          let _index28 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          let _index29 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_200_v29).length));
          let _index30 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          let _rhs24 = _193_v22;
          let _rhs25 = new BigNumber(-536);
          let _rhs26 = (_213_v42).fm8(!(((_213_v42).fm7(_193_v22, _173_globalState)).isLessThanOrEqualTo(_167_v0)), (_190_v19)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length))], _193_v22, _173_globalState);
          let _rhs27 = _170_v3;
          let _lhs19 = _190_v19;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          let _lhs21 = _200_v29;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_200_v29).length));
          let _lhs23 = _190_v19;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          _lhs19[_lhs20] = _rhs24;
          _lhs21[_lhs22] = _rhs25;
          _167_v0 = _rhs26;
          _lhs23[_lhs24] = _rhs27;
          let _index31 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          (_190_v19)[_index31] = _193_v22;
        } else {
          let _index32 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          (_190_v19)[_index32] = _170_v3;
          (_173_globalState).f1 = _module.__default.fm3(_193_v22, _193_v22, _173_globalState);
          let _215_v44;
          let _nw22 = new _module.C6();
          _nw22.__ctor();
          _215_v44 = _nw22;
          let _216_v45;
          let _nw23 = Array((new BigNumber(9)).toNumber()).fill(false);
          _216_v45 = _nw23;
          let _217_v46;
          _217_v46 = _dafny.Seq.of(_216_v45);
          let _218_v47;
          let _nw24 = Array((new BigNumber(7)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = ((_193_v22) ? (_216_v45) : (_216_v45));
          _nw24[(_dafny.ONE).toNumber()] = _216_v45;
          _nw24[(new BigNumber(2)).toNumber()] = (_217_v46)[_module.__default.safeIndex(_167_v0, new BigNumber((_217_v46).length))];
          _nw24[(new BigNumber(3)).toNumber()] = _216_v45;
          _nw24[(new BigNumber(4)).toNumber()] = _216_v45;
          _nw24[(new BigNumber(5)).toNumber()] = _216_v45;
          _nw24[(new BigNumber(6)).toNumber()] = _216_v45;
          _218_v47 = _nw24;
          let _index33 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_218_v47).length));
          (_218_v47)[_index33] = _216_v45;
          let _index34 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_218_v47).length));
          (_218_v47)[_index34] = (_217_v46)[_module.__default.safeIndex((_dafny.ZERO).minus(_167_v0), new BigNumber((_217_v46).length))];
          let _index35 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_190_v19).length));
          (_190_v19)[_index35] = !(!(_170_v3));
        }
      } else {
        let _219_v48;
        _219_v48 = _dafny.Set.fromElements(_167_v0, new BigNumber(((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]).length));
        _167_v0 = _module.__default.fm3((false) && (_193_v22), (_219_v48).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(306), _167_v0)), _173_globalState);
        if ((_170_v3) || (_module.__default.fm0(_193_v22, _173_globalState))) {
          let _220_v49;
          let _init4 = ((_221_v20, _222_v0) => function (_223_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_221_v20,_222_v0);
          })(_191_v20, _167_v0);
          let _nw25 = Array((new BigNumber(21)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw25.length); _i0_4++) {
            _nw25[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _220_v49 = _nw25;
          let _224_v50;
          _224_v50 = _dafny.Seq.of(_220_v49, _220_v49, _220_v49, _220_v49, _220_v49);
          let _225_v51;
          _225_v51 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,new BigNumber((_172_v5).length));
          let _226_v52;
          let _nw26 = Array((new BigNumber(5)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _220_v49;
          _nw26[(_dafny.ONE).toNumber()] = _220_v49;
          _nw26[(new BigNumber(2)).toNumber()] = (_224_v50)[_module.__default.safeIndex(new BigNumber((_225_v51).length), new BigNumber((_224_v50).length))];
          _nw26[(new BigNumber(3)).toNumber()] = _220_v49;
          _nw26[(new BigNumber(4)).toNumber()] = _220_v49;
          _226_v52 = _nw26;
          let _index36 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_226_v52).length));
          (_226_v52)[_index36] = _220_v49;
          let _index37 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_226_v52).length));
          let _rhs28 = _module.__default.safeModuloInt(((_dafny.ZERO).minus(_167_v0)).multipliedBy(_167_v0), _167_v0);
          let _rhs29 = _220_v49;
          let _lhs25 = _173_globalState;
          let _lhs26 = _226_v52;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_226_v52).length));
          _lhs25.f1 = _rhs28;
          _lhs26[_lhs27] = _rhs29;
          let _227_v53;
          let _nw27 = new _module.C6();
          _nw27.__ctor();
          _227_v53 = _nw27;
          let _228_v54;
          _228_v54 = _dafny.Map.Empty.slice().updateUnsafe(_227_v53,_193_v22);
          let _229_v55;
          _229_v55 = _dafny.Map.Empty.slice().updateUnsafe((_228_v54).Merge(_dafny.Map.Empty.slice().updateUnsafe(_227_v53,_170_v3)),new BigNumber(719));
          (_173_globalState).f1 = new BigNumber((_229_v55).length);
          _167_v0 = (_dafny.ZERO).minus(_167_v0);
          _170_v3 = !(_module.__default.fm0((_227_v53).fm6(_dafny.Seq.of(_167_v0), _173_globalState), _173_globalState)) || (_170_v3);
          let _230_v56;
          let _out4;
          _out4 = (_227_v53).m3(_173_globalState);
          _230_v56 = _out4;
        } else {
          _219_v48 = _219_v48;
          let _231_v57;
          _231_v57 = _dafny.Map.Empty.slice().updateUnsafe(_193_v22,_167_v0);
          let _232_v58;
          _232_v58 = _dafny.MultiSet.fromElements(_231_v57, _231_v57);
          let _233_v60;
          _233_v60 = _dafny.MultiSet.fromElements(new BigNumber((_231_v57).length));
          let _234_v61;
          _234_v61 = function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_233_v60).Elements) {
              let _235_v59 = _compr_20;
              if ((_233_v60).contains(_235_v59)) {
                _coll20.push([(_235_v59).plus(_167_v0),_193_v22]);
              }
            }
            return _coll20;
          }();
          let _236_v62;
          _236_v62 = _dafny.Map.Empty.slice().updateUnsafe(_234_v61,_170_v3);
          let _237_v63;
          _237_v63 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_170_v3,new BigNumber((_236_v62).length)), _231_v57);
          let _238_v64;
          _238_v64 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_237_v63));
          let _239_v65;
          _239_v65 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_232_v58);
          let _rhs30 = (_232_v58).Union(((_193_v22) ? ((_238_v64)[_module.__default.safeIndex((_dafny.ZERO).minus(_167_v0), new BigNumber((_238_v64).length))]) : ((((_239_v65).contains(_167_v0)) ? ((_239_v65).get(_167_v0)) : (_232_v58)))));
          let _rhs31 = (new BigNumber((_172_v5).length)).multipliedBy(_167_v0);
          _232_v58 = _rhs30;
          _167_v0 = _rhs31;
          let _240_v66;
          let _init5 = ((_241_v20) => function (_242_i3) {
            return _241_v20;
          })(_191_v20);
          let _nw28 = Array((new BigNumber(20)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw28.length); _i0_5++) {
            _nw28[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _240_v66 = _nw28;
          let _243_v67;
          _243_v67 = _dafny.Map.Empty.slice().updateUnsafe(_240_v66,_170_v3);
          _243_v67 = (_243_v67).update(_240_v66, (new BigNumber((_233_v60).cardinality())).isLessThanOrEqualTo(_167_v0));
          let _244_v68;
          let _nw29 = new _module.C6();
          _nw29.__ctor();
          _244_v68 = _nw29;
          let _245_v69;
          _245_v69 = _dafny.Map.Empty.slice().updateUnsafe(_170_v3,_170_v3);
          let _246_v70;
          _246_v70 = _dafny.MultiSet.fromElements((_245_v69).Merge(_245_v69), (_245_v69).update(_170_v3, _170_v3));
          let _index38 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_190_v19).length));
          (_190_v19)[_index38] = !(_245_v69).contains(_170_v3);
          let _247_v71;
          _247_v71 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,new BigNumber((_219_v48).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_190_v19).length));
          let _rhs32 = _dafny.MultiSet.fromElements(_245_v69, ((_dafny.Map.Empty.slice().updateUnsafe(_193_v22,_170_v3)).update(_170_v3, true)).update(_193_v22, _193_v22), _245_v69);
          let _rhs33 = _dafny.Seq.Concat(_168_v1, _168_v1);
          let _rhs34 = ((_193_v22) ? ((_170_v3) === (_170_v3)) : (_170_v3));
          let _rhs35 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_167_v0, (((_247_v71).contains(new BigNumber(929))) ? ((_247_v71).get(new BigNumber(929))) : ((_dafny.ZERO).minus(_167_v0)))));
          let _lhs28 = _190_v19;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_190_v19).length));
          let _lhs30 = _173_globalState;
          _246_v70 = _rhs32;
          _168_v1 = _rhs33;
          _lhs28[_lhs29] = _rhs34;
          _lhs30.f1 = _rhs35;
        }
        let _248_v72;
        _248_v72 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_170_v3);
        let _249_v73;
        _249_v73 = _dafny.MultiSet.fromElements(new BigNumber((_248_v72).length), new BigNumber((_168_v1).length));
        let _250_v74;
        let _251_v75;
        let _out5;
        let _out6;
        let _outcollector2 = _module.__default.m0(_190_v19, _191_v20, (_249_v73).IsSubsetOf(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]).length)), _module.__default.fm3(_193_v22, _170_v3, _173_globalState))), _191_v20, _173_globalState);
        _out5 = _outcollector2[0];
        _out6 = _outcollector2[1];
        _250_v74 = _out5;
        _251_v75 = _out6;
        _248_v72 = (_248_v72).update(new BigNumber(253), _251_v75);
        let _252_v76;
        _252_v76 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,new BigNumber(-98));
        let _253_v77;
        let _254_v78;
        let _out7;
        let _out8;
        let _outcollector3 = _module.__default.m0(_190_v19, _191_v20, (_170_v3) && (_193_v22), _module.__default.fm36(_167_v0, _167_v0, _167_v0, _252_v76, _173_globalState), _173_globalState);
        _out7 = _outcollector3[0];
        _out8 = _outcollector3[1];
        _253_v77 = _out7;
        _254_v78 = _out8;
      }
      _167_v0 = ((_170_v3) ? (_167_v0) : (_167_v0));
      let _255_v79;
      _255_v79 = _dafny.Map.Empty.slice().updateUnsafe(_170_v3,_189_v18);
      let _256_v80;
      _256_v80 = _dafny.Map.Empty.slice().updateUnsafe(!(_193_v22) || (_170_v3),_255_v79);
      let _257_v81;
      let _nw30 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _257_v81 = _nw30;
      let _258_v82;
      _258_v82 = _dafny.Seq.of(_257_v81, _257_v81, _257_v81, _257_v81);
      let _259_v83;
      let _nw31 = new _module.C8();
      _nw31.__ctor(_190_v19, _258_v82, _191_v20, _193_v22);
      _259_v83 = _nw31;
      let _260_v84;
      _260_v84 = _dafny.MultiSet.fromElements(_170_v3);
      let _261_v85;
      _261_v85 = _module.D22.create_DC48(_259_v83);
      let _rhs36 = ((_170_v3) ? (_256_v80) : (_256_v80));
      let _rhs37 = ((new BigNumber((_260_v84).cardinality())).minus(_167_v0)).plus(_167_v0);
      let _rhs38 = _dafny.Seq.UnicodeFromString("ibtmhviyp");
      let _rhs39 = (_261_v85).dtor_cf65;
      _256_v80 = _rhs36;
      _167_v0 = _rhs37;
      _196_v25 = _rhs38;
      _259_v83 = _rhs39;
      let _262_v86;
      _262_v86 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_167_v0),_dafny.Seq.of(new BigNumber(598), _167_v0, _167_v0, _167_v0));
      let _263_v87;
      _263_v87 = _module.D19.create_DC40(_170_v3, (_259_v83).fm8(_193_v22, _170_v3, _193_v22, _173_globalState), _262_v86);
      let _264_v88;
      _264_v88 = _dafny.Map.Empty.slice().updateUnsafe(_170_v3,_167_v0);
      let _265_v89;
      _265_v89 = _dafny.MultiSet.fromElements(_264_v88);
      if ((new BigNumber((_265_v89).cardinality())).isLessThanOrEqualTo((_263_v87).dtor_cf52)) {
        let _266_v90;
        _266_v90 = _dafny.Set.fromElements(_257_v81, _257_v81);
        let _267_v91;
        _267_v91 = _module.D22.create_DC49(_266_v90);
        let _source7 = _267_v91;
        if (_source7.is_DC49) {
          let _268___mcc_h0 = (_source7).cf66;
          let _269_cf66 = _268___mcc_h0;
          _170_v3 = _193_v22;
          let _270_v92;
          _270_v92 = _dafny.Set.fromElements(_170_v3);
          let _271_v93;
          _271_v93 = _dafny.Map.Empty.slice().updateUnsafe(_270_v92,new BigNumber(((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]).length));
          let _272_v94;
          let _273_v95;
          let _out9;
          let _out10;
          let _outcollector4 = (_259_v83).m2(_193_v22, (new BigNumber((_271_v93).length)).plus(_167_v0), _196_v25, _173_globalState);
          _out9 = _outcollector4[0];
          _out10 = _outcollector4[1];
          _272_v94 = _out9;
          _273_v95 = _out10;
          (_173_globalState).f1 = (_167_v0).plus(_167_v0);
          (_173_globalState).f1 = _167_v0;
        } else {
          let _274___mcc_h1 = (_source7).cf65;
          let _275_cf65 = _274___mcc_h1;
          let _276_v96;
          _276_v96 = _dafny.Set.fromElements(((_dafny.ZERO).minus(_167_v0)).minus(_167_v0));
          let _277_v97;
          _277_v97 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_167_v0)).length));
          _276_v96 = _dafny.Set.fromElements(((_170_v3) ? (new BigNumber((_277_v97).cardinality())) : ((((_260_v84).contains(_193_v22)) ? ((_260_v84).get(_193_v22)) : ((_259_v83).fm8(_170_v3, _170_v3, _193_v22, _173_globalState))))), _167_v0, _167_v0, (_275_cf65).fm8(_170_v3, _170_v3, _170_v3, _173_globalState));
          (_275_cf65).m5(_module.__default.fm1(_167_v0, _170_v3, _167_v0, _173_globalState), _167_v0, _173_globalState);
          (_173_globalState).f1 = (_module.__default.safeDivisionInt(_167_v0, (_dafny.ZERO).minus(_167_v0))).plus(new BigNumber((_dafny.Seq.Concat((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))], _196_v25)).length));
          let _278_v98;
          let _nw32 = new _module.C0();
          _nw32.__ctor(_193_v22, _dafny.Seq.Concat(_169_v2, _169_v2));
          _278_v98 = _nw32;
        }
        _168_v1 = _168_v1;
        let _index40 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length));
        (_257_v81)[_index40] = _167_v0;
        let _279_v99;
        _279_v99 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_193_v22);
        let _280_v100;
        _280_v100 = _dafny.Set.fromElements(_193_v22);
        let _281_v101;
        let _nw33 = new _module.C5();
        _nw33.__ctor(new BigNumber((_280_v100).length), _191_v20, _193_v22);
        _281_v101 = _nw33;
        let _282_v102;
        _282_v102 = _dafny.Seq.of(_281_v101, _281_v101);
        let _283_v103;
        _283_v103 = _dafny.Map.Empty.slice().updateUnsafe((((_279_v99).contains(_167_v0)) ? ((_279_v99).get(_167_v0)) : (!(_193_v22))),_dafny.Seq.update(_282_v102, _module.__default.safeIndex((_dafny.ZERO).minus(_167_v0), new BigNumber((_282_v102).length)), _281_v101));
        let _index41 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length));
        let _rhs40 = (_dafny.ZERO).minus(_167_v0);
        let _rhs41 = _dafny.Seq.Concat(_168_v1, _dafny.Seq.Concat(_168_v1, _168_v1));
        let _rhs42 = (_283_v103).Merge((_283_v103).Merge((_283_v103).update(_193_v22, _282_v102)));
        let _rhs43 = _167_v0;
        let _lhs31 = _257_v81;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length));
        let _lhs33 = _281_v101;
        _lhs31[_lhs32] = _rhs40;
        _168_v1 = _rhs41;
        _283_v103 = _rhs42;
        _lhs33.f15 = _rhs43;
        let _284_v104;
        let _nw34 = new _module.C0();
        _nw34.__ctor(((((_264_v88).contains(!(_193_v22))) ? ((_264_v88).get(!(_193_v22))) : (_281_v101.f15))).isLessThan(new BigNumber(642)), _dafny.Seq.update(_dafny.Seq.Concat(_169_v2, _169_v2), _module.__default.safeIndex(_281_v101.f15, new BigNumber((_dafny.Seq.Concat(_169_v2, _169_v2)).length)), _168_v1));
        _284_v104 = _nw34;
        let _index42 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
        let _rhs44 = _284_v104;
        let _rhs45 = (_dafny.ZERO).minus(_167_v0);
        let _rhs46 = _196_v25;
        let _lhs34 = _173_globalState;
        let _lhs35 = _195_v24;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
        _284_v104 = _rhs44;
        _lhs34.f1 = _rhs45;
        _lhs35[_lhs36] = _rhs46;
        let _285_v105;
        _285_v105 = _dafny.MultiSet.fromElements(_191_v20);
        let _286_v106;
        let _nw35 = Array((new BigNumber(29)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _281_v101.f15;
        _nw35[(_dafny.ONE).toNumber()] = (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))];
        _nw35[(new BigNumber(2)).toNumber()] = (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))];
        _nw35[(new BigNumber(3)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(4)).toNumber()] = _167_v0;
        _nw35[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_168_v1).length));
        _nw35[(new BigNumber(6)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(7)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(8)).toNumber()] = new BigNumber(((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]).length);
        _nw35[(new BigNumber(9)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(10)).toNumber()] = _167_v0;
        _nw35[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.of(false)).length);
        _nw35[(new BigNumber(12)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(13)).toNumber()] = (((_285_v105).contains(_191_v20)) ? ((_285_v105).get(_191_v20)) : (new BigNumber(73)));
        _nw35[(new BigNumber(14)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(15)).toNumber()] = (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))];
        _nw35[(new BigNumber(16)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(17)).toNumber()] = new BigNumber(864);
        _nw35[(new BigNumber(18)).toNumber()] = new BigNumber(933);
        _nw35[(new BigNumber(19)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(20)).toNumber()] = new BigNumber((_279_v99).length);
        _nw35[(new BigNumber(21)).toNumber()] = (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))];
        _nw35[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Set.fromElements(new BigNumber((_168_v1).length))).length);
        _nw35[(new BigNumber(23)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(24)).toNumber()] = new BigNumber((_168_v1).length);
        _nw35[(new BigNumber(25)).toNumber()] = (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))];
        _nw35[(new BigNumber(26)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(27)).toNumber()] = _281_v101.f15;
        _nw35[(new BigNumber(28)).toNumber()] = (_281_v101).fm8(_170_v3, _170_v3, !(_193_v22), _173_globalState);
        _286_v106 = _nw35;
        let _287_v107;
        let _288_v108;
        let _out11;
        let _out12;
        let _outcollector5 = (_259_v83).m4(_286_v106, _172_v5, (_284_v104).f11, (_257_v81)[_module.__default.safeIndex(new BigNumber(742), new BigNumber((_257_v81).length))], _173_globalState);
        _out11 = _outcollector5[0];
        _out12 = _outcollector5[1];
        _287_v107 = _out11;
        _288_v108 = _out12;
      } else {
        let _289_v109;
        _289_v109 = _dafny.Set.fromElements(_259_v83.f7);
        let _index43 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length));
        (_190_v19)[_index43] = !(_289_v109).contains(_259_v83.f7);
        let _index44 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length));
        (_190_v19)[_index44] = _170_v3;
        let _source8 = _192_v21;
        if (_source8.is_DC1) {
          (_173_globalState).f1 = (_167_v0).plus((_dafny.ZERO).minus((_167_v0).plus((_dafny.ZERO).minus(_167_v0))));
          let _290_v110;
          _290_v110 = _dafny.Map.Empty.slice().updateUnsafe(_193_v22,_191_v20);
          let _291_v111;
          _291_v111 = _dafny.Set.fromElements(_290_v110);
          (_259_v83).m5(_196_v25, new BigNumber((_291_v111).length), _173_globalState);
          let _292_v112;
          _292_v112 = _dafny.Map.Empty.slice().updateUnsafe(_170_v3,_193_v22);
          let _293_v113;
          _293_v113 = _dafny.MultiSet.fromElements(_292_v112);
          (_173_globalState).f1 = (new BigNumber((_293_v113).cardinality())).multipliedBy(_167_v0);
          let _index45 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length));
          (_190_v19)[_index45] = false;
        } else if (_source8.is_DC2) {
          let _294___mcc_h2 = (_source8).cf1;
          let _295___mcc_h3 = (_source8).cf2;
          let _296_cf2 = _295___mcc_h3;
          let _297_cf1 = _294___mcc_h2;
          let _298_v114;
          let _nw36 = new _module.C3();
          _nw36.__ctor(_dafny.Seq.IsPrefixOf(_196_v25, (_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]), true);
          _298_v114 = _nw36;
          let _299_v115;
          _299_v115 = _dafny.Set.fromElements(_170_v3);
          let _300_v116;
          _300_v116 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_299_v115);
          let _301_v117;
          _301_v117 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_193_v22),_170_v3);
          let _rhs47 = _module.__default.safeModuloInt(_167_v0, _module.__default.safeDivisionInt(_167_v0, _167_v0));
          let _rhs48 = _module.__default.fm22((_299_v115).Difference(((((_300_v116).contains(_167_v0)) ? ((_300_v116).get(_167_v0)) : (_299_v115)))), _module.__default.safeDivisionInt((_259_v83).fm7((_298_v114).f14, _173_globalState), _167_v0), (_259_v83).fm10(_260_v84, _301_v117, _167_v0, _173_globalState), (_module.__default.fm48(_167_v0, !(true), _173_globalState)).IsProperSubsetOf(_dafny.MultiSet.fromElements(true)), _173_globalState);
          let _rhs49 = _167_v0;
          let _lhs37 = _173_globalState;
          let _lhs38 = _173_globalState;
          _lhs37.f1 = _rhs47;
          _296_cf2 = _rhs48;
          _lhs38.f1 = _rhs49;
          let _rhs50 = _170_v3;
          let _rhs51 = (_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))];
          let _rhs52 = _167_v0;
          let _lhs39 = _173_globalState;
          _193_v22 = _rhs50;
          _196_v25 = _rhs51;
          _lhs39.f1 = _rhs52;
          let _302_v118;
          let _nw37 = Array((_dafny.ONE).toNumber()).fill([]);
          _302_v118 = _nw37;
          _302_v118 = _302_v118;
        } else if (_source8.is_DC0) {
          let _303___mcc_h4 = (_source8).cf0;
          let _304_cf0 = _303___mcc_h4;
          let _305_v119;
          let _nw38 = new _module.C4();
          _nw38.__ctor();
          _305_v119 = _nw38;
          let _306_v120;
          _306_v120 = _dafny.Seq.of(_305_v119);
          let _307_v121;
          _307_v121 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_172_v5)).cardinality()), _167_v0);
          let _308_v122;
          _308_v122 = _dafny.Set.fromElements(_167_v0, new BigNumber((_307_v121).cardinality()));
          let _309_v123;
          let _nw39 = Array((new BigNumber(21)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _305_v119;
          _nw39[(_dafny.ONE).toNumber()] = _305_v119;
          _nw39[(new BigNumber(2)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(3)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(4)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(5)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(6)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(7)).toNumber()] = (_306_v120)[_module.__default.safeIndex(new BigNumber((_308_v122).length), new BigNumber((_306_v120).length))];
          _nw39[(new BigNumber(8)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(9)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(10)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(11)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(12)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(13)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(14)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(15)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(16)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(17)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(18)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(19)).toNumber()] = _305_v119;
          _nw39[(new BigNumber(20)).toNumber()] = _305_v119;
          _309_v123 = _nw39;
          let _index46 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_309_v123).length));
          (_309_v123)[_index46] = ((_170_v3) ? (_305_v119) : (_305_v119));
          let _index47 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_309_v123).length));
          let _nw40 = new _module.C4();
          _nw40.__ctor();
          (_309_v123)[_index47] = _nw40;
          let _310_v124;
          let _311_v125;
          let _312_v126;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector6 = (_259_v83).m6(true, (_259_v83).fm10(_260_v84, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_193_v22),(_module.D1.create_DC5(_193_v22)).dtor_cf5), _167_v0, _173_globalState), _173_globalState);
          _out13 = _outcollector6[0];
          _out14 = _outcollector6[1];
          _out15 = _outcollector6[2];
          _310_v124 = _out13;
          _311_v125 = _out14;
          _312_v126 = _out15;
          let _313_v127;
          _313_v127 = _dafny.MultiSet.fromElements(_172_v5, _172_v5);
          let _index48 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length));
          (_190_v19)[_index48] = ((_313_v127).update(_172_v5, _module.__default.abs(_167_v0))).contains(_172_v5);
          let _index49 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_257_v81).length));
          (_257_v81)[_index49] = _167_v0;
          let _index50 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_257_v81).length));
          (_257_v81)[_index50] = _module.__default.safeModuloInt(((_312_v126) ? (new BigNumber(110)) : (new BigNumber(438))), _167_v0);
        } else {
          let _314___mcc_h5 = (_source8).cf3;
          let _315_cf3 = _314___mcc_h5;
          let _316_v128;
          _316_v128 = _dafny.Map.Empty.slice().updateUnsafe(_257_v81,_170_v3);
          let _317_v129;
          let _nw41 = new _module.C5();
          _nw41.__ctor(new BigNumber(((((_190_v19)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length))]) ? ((_316_v128).update(_257_v81, _193_v22)) : (_316_v128))).length), (_196_v25)[_module.__default.safeIndex((_dafny.ZERO).minus(_167_v0), new BigNumber((_196_v25).length))], !(_170_v3));
          _317_v129 = _nw41;
          let _318_v130;
          let _319_v131;
          let _out16;
          let _out17;
          let _outcollector7 = (_259_v83).m2(_193_v22, _317_v129.f15, _dafny.Seq.UnicodeFromString("ylqpcjtj"), _173_globalState);
          _out16 = _outcollector7[0];
          _out17 = _outcollector7[1];
          _318_v130 = _out16;
          _319_v131 = _out17;
          let _index51 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_190_v19).length));
          (_190_v19)[_index51] = _170_v3;
          let _index52 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_257_v81).length));
          (_257_v81)[_index52] = new BigNumber((_dafny.Seq.Concat(_172_v5, _172_v5)).length);
          let _index53 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_257_v81).length));
          (_257_v81)[_index53] = (_167_v0).multipliedBy((_317_v129.f15).multipliedBy(_167_v0));
        }
        _191_v20 = _191_v20;
        let _nw42 = new _module.C8();
        _nw42.__ctor(_259_v83.f7, _258_v82, _191_v20, _170_v3);
        _259_v83 = _nw42;
        _193_v22 = !_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("asmrl"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("pjcwgjki")).length), new BigNumber((_dafny.Seq.UnicodeFromString("asmrl")).length)), _191_v20), new _dafny.CodePoint('f'.codePointAt(0)));
      }
      let _hi0 = (_167_v0).minus(_167_v0);
      for (let _320_i4 = _167_v0; _320_i4.isLessThan(_hi0); _320_i4 = _320_i4.plus(_dafny.ONE)) {
        _167_v0 = _167_v0;
        let _index54 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_190_v19).length));
        (_190_v19)[_index54] = _170_v3;
        let _321_v132;
        _321_v132 = _dafny.MultiSet.fromElements((_320_i4).plus(new BigNumber((_260_v84).cardinality())));
        let _322_v133;
        let _nw43 = new _module.C5();
        _nw43.__ctor((_dafny.ZERO).minus(_167_v0), _191_v20, false);
        _322_v133 = _nw43;
        let _323_v134;
        _323_v134 = _dafny.Map.Empty.slice().updateUnsafe(_320_i4,_322_v133);
        let _324_v135;
        _324_v135 = _dafny.Set.fromElements(_322_v133, _322_v133, _322_v133, _322_v133, (((_323_v134).contains(new BigNumber(-341))) ? ((_323_v134).get(new BigNumber(-341))) : (_322_v133)));
        let _325_v136;
        _325_v136 = _dafny.Seq.of(_321_v132, _321_v132);
        let _index55 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_190_v19).length));
        let _rhs53 = _170_v3;
        let _rhs54 = ((_dafny.MultiSet.fromElements((_259_v83).fm4(_172_v5, _320_i4, _320_i4, _dafny.MultiSet.fromElements(_167_v0), _173_globalState), _320_i4, _167_v0, (_dafny.ZERO).minus(_167_v0), new BigNumber((_324_v135).length))).Union(_321_v132)).Union((_325_v136)[_module.__default.safeIndex(new BigNumber((_196_v25).length), new BigNumber((_325_v136).length))]);
        let _rhs55 = (((_322_v133).f6) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(983), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), ((_326_v133) => function (_327_i5) {
          return (_326_v133).f5;
        })(_322_v133))).length))))) : ((_322_v133).fm7(_193_v22, _173_globalState)));
        let _lhs40 = _190_v19;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_190_v19).length));
        let _lhs42 = _173_globalState;
        _lhs40[_lhs41] = _rhs53;
        _321_v132 = _rhs54;
        _lhs42.f1 = _rhs55;
        let _328_v137;
        let _nw44 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
        _328_v137 = _nw44;
        let _index56 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_328_v137).length));
        (_328_v137)[_index56] = _321_v132;
        let _index57 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_328_v137).length));
        let _rhs56 = (_321_v132).Difference(_321_v132);
        let _rhs57 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm1(new BigNumber(347), (_190_v19)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_190_v19).length))], _167_v0, _173_globalState), _196_v25)).length);
        let _rhs58 = (_259_v83).fm4(_172_v5, _320_i4, _167_v0, _321_v132, _173_globalState);
        let _lhs43 = _328_v137;
        let _lhs44 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_328_v137).length));
        let _lhs45 = _173_globalState;
        _lhs43[_lhs44] = _rhs56;
        _lhs45.f1 = _rhs57;
        _167_v0 = _rhs58;
        let _329_v138;
        _329_v138 = _dafny.Map.Empty.slice().updateUnsafe(_320_i4,_320_i4);
        let _330_v139;
        _330_v139 = _dafny.Map.Empty.slice().updateUnsafe(_329_v138,_172_v5);
        let _index58 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_257_v81).length));
        (_257_v81)[_index58] = new BigNumber(782);
        let _331_v140;
        let _nw45 = new _module.C3();
        _nw45.__ctor(_193_v22, true);
        _331_v140 = _nw45;
        let _332_v141;
        _332_v141 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,_331_v140);
        let _333_v142;
        _333_v142 = _330_v139;
        let _334_v143;
        _334_v143 = _module.D20.create_DC43(new BigNumber(692), (_190_v19)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_190_v19).length))], _167_v0);
        let _index59 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_257_v81).length));
        let _index60 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
        let _rhs59 = (_dafny.ZERO).minus(_320_i4);
        let _rhs60 = !(_167_v0).isEqualTo(new BigNumber((_332_v141).length));
        let _rhs61 = (_333_v142);
        let _rhs62 = (_334_v143).dtor_cf56;
        let _rhs63 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), ((_335_v133) => function (_336_i6) {
          return (_335_v133).f5;
        })(_322_v133));
        let _lhs46 = _173_globalState;
        let _lhs47 = _257_v81;
        let _lhs48 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_257_v81).length));
        let _lhs49 = _195_v24;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length));
        _lhs46.f1 = _rhs59;
        _193_v22 = _rhs60;
        _330_v139 = _rhs61;
        _lhs47[_lhs48] = _rhs62;
        _lhs49[_lhs50] = _rhs63;
      }
      if ((_167_v0).isLessThan(_167_v0)) {
        let _337_v144;
        let _nw46 = new _module.C2();
        _nw46.__ctor(_191_v20, _170_v3);
        _337_v144 = _nw46;
        _337_v144 = _337_v144;
        (_173_globalState).f1 = (new BigNumber((_196_v25).length)).minus(_167_v0);
        let _338_v145;
        _338_v145 = _dafny.Seq.of((_195_v24)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_195_v24).length))]);
        let _index61 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_257_v81).length));
        (_257_v81)[_index61] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("b"), (_338_v145)[_module.__default.safeIndex((_337_v144).fm7(_193_v22, _173_globalState), new BigNumber((_338_v145).length))])).length);
        let _index62 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_257_v81).length));
        let _rhs64 = (_dafny.ZERO).minus((_167_v0).minus(_167_v0));
        let _rhs65 = _167_v0;
        let _rhs66 = (_337_v144).fm9(_167_v0, _171_v4, false, _167_v0, _173_globalState);
        let _rhs67 = (_167_v0).minus((_259_v83).fm8(_193_v22, true, _170_v3, _173_globalState));
        let _lhs51 = _173_globalState;
        let _lhs52 = _257_v81;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_257_v81).length));
        _lhs51.f1 = _rhs64;
        _167_v0 = _rhs65;
        _193_v22 = _rhs66;
        _lhs52[_lhs53] = _rhs67;
        _167_v0 = new BigNumber(358);
        _193_v22 = true;
      } else {
        (_173_globalState).f1 = _167_v0;
        _170_v3 = _193_v22;
        let _339_v146;
        let _nw47 = new _module.C3();
        _nw47.__ctor(!(_170_v3), _dafny.areEqual(_172_v5, _172_v5));
        _339_v146 = _nw47;
        _193_v22 = !((_339_v146).f14);
        _195_v24 = ((((_339_v146).f13) && (_193_v22)) ? (_195_v24) : (_195_v24));
      }
      let _340_v147;
      _340_v147 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,_167_v0);
      _170_v3 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(151)), ((_341_v20) => function (_342_i7) {
        return _341_v20;
      })(_191_v20))).length), _167_v0)).isLessThanOrEqualTo((_dafny.ZERO).minus((_167_v0).minus(new BigNumber((_340_v147).length))));
      let _343_v148;
      _343_v148 = _dafny.Map.Empty.slice().updateUnsafe(_193_v22,_170_v3);
      let _344_v149;
      let _nw48 = new _module.C4();
      _nw48.__ctor();
      _344_v149 = _nw48;
      let _345_v150;
      _345_v150 = _dafny.Seq.of(_344_v149);
      let _346_v151;
      _346_v151 = _dafny.MultiSet.fromElements(new BigNumber(829), _167_v0, new BigNumber((_345_v150).length), _167_v0, _167_v0);
      let _347_i8;
      _347_i8 = _dafny.ZERO;
      L3: {
        while (((_259_v83).fm4(_172_v5, new BigNumber((_343_v148).length), (_dafny.ZERO).minus(_167_v0), _346_v151, _173_globalState)).isLessThan(_167_v0)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_347_i8)) {
              break L3;
            }
            _347_i8 = (_347_i8).plus(_dafny.ONE);
            let _348_v152;
            let _nw49 = Array((new BigNumber(19)).toNumber()).fill([]);
            _348_v152 = _nw49;
            let _index63 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_348_v152).length));
            (_348_v152)[_index63] = _190_v19;
            let _index64 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_348_v152).length));
            (_348_v152)[_index64] = _259_v83.f7;
            let _349_v153;
            _349_v153 = _dafny.Set.fromElements(_167_v0);
            let _350_v154;
            _350_v154 = _dafny.Map.Empty.slice().updateUnsafe(_349_v153,(((_340_v147).contains(_167_v0)) ? ((_340_v147).get(_167_v0)) : (_167_v0)));
            let _351_v155;
            _351_v155 = _dafny.Set.fromElements(_171_v4);
            (_173_globalState).f1 = (((_260_v84).contains(true)) ? ((_260_v84).get(true)) : ((_167_v0).minus((((_350_v154).contains(_dafny.Set.fromElements(_167_v0, _167_v0))) ? ((_350_v154).get(_dafny.Set.fromElements(_167_v0, _167_v0))) : (new BigNumber((_351_v155).length))))));
            let _352_v156;
            let _353_v157;
            let _out18;
            let _out19;
            let _outcollector8 = (_259_v83).m4(_257_v81, _dafny.Seq.of(_193_v22), _193_v22, (((_260_v84).contains(!(false))) ? ((_260_v84).get(!(false))) : ((_259_v83).fm4(_172_v5, _167_v0, new BigNumber(903), _346_v151, _173_globalState))), _173_globalState);
            _out18 = _outcollector8[0];
            _out19 = _outcollector8[1];
            _352_v156 = _out18;
            _353_v157 = _out19;
            let _rhs68 = ((_module.__default.fm0((_172_v5)[_module.__default.safeIndex(_167_v0, new BigNumber((_172_v5).length))], _173_globalState)) ? (_170_v3) : ((_346_v151).IsSubsetOf(_346_v151)));
            let _rhs69 = (_dafny.MultiSet.FromArray(_168_v1)).Intersect(_346_v151);
            _170_v3 = _rhs68;
            _346_v151 = _rhs69;
          }
        }
      }
      let _354_i9;
      _354_i9 = _dafny.ZERO;
      L4: {
        while (_193_v22) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_354_i9)) {
              break L4;
            }
            _354_i9 = (_354_i9).plus(_dafny.ONE);
            _344_v149 = _344_v149;
            (_173_globalState).f1 = new BigNumber((_346_v151).cardinality());
            (_173_globalState).f1 = (_167_v0).multipliedBy(_167_v0);
            let _index65 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_257_v81).length));
            (_257_v81)[_index65] = (_module.__default.fm3(false, _170_v3, _173_globalState)).minus((_dafny.ZERO).minus(_167_v0));
            let _355_v158;
            let _nw50 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
            _355_v158 = _nw50;
            let _356_v159;
            _356_v159 = _module.D16.create_DC31(_193_v22);
            let _357_v160;
            _357_v160 = _dafny.MultiSet.fromElements(_356_v159);
            let _index66 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_355_v158).length));
            (_355_v158)[_index66] = _357_v160;
            let _358_v161;
            _358_v161 = _module.D25.create_DC52(_dafny.MultiSet.fromElements(_356_v159));
            let _359_v162;
            _359_v162 = _dafny.Map.Empty.slice().updateUnsafe(_167_v0,(_358_v161).dtor_cf69);
            let _index67 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_257_v81).length));
            let _index68 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_355_v158).length));
            let _rhs70 = _167_v0;
            let _rhs71 = ((((_359_v162).contains(new BigNumber(491))) ? ((_359_v162).get(new BigNumber(491))) : (_357_v160))).Intersect((_357_v160).update(_356_v159, _module.__default.abs(_167_v0)));
            let _rhs72 = _190_v19;
            let _rhs73 = _167_v0;
            let _lhs54 = _257_v81;
            let _lhs55 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_257_v81).length));
            let _lhs56 = _355_v158;
            let _lhs57 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_355_v158).length));
            let _lhs58 = _259_v83;
            _lhs54[_lhs55] = _rhs70;
            _lhs56[_lhs57] = _rhs71;
            _lhs58.f7 = _rhs72;
            _167_v0 = _rhs73;
          }
        }
      }
      process.stdout.write(_dafny.toString(_167_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_168_v1, _dafny.Seq.of(new BigNumber(647), new BigNumber(647), new BigNumber(647), new BigNumber(647), new BigNumber(647), new BigNumber(647)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_169_v2, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(647), new BigNumber(647))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_170_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v4).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_172_v5, _dafny.Seq.of(true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_173_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_173_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_173_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_173_globalState).f3, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(647), new BigNumber(647))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_173_globalState.f4).equals(_dafny.MultiSet.fromElements(true, true, true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_174_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_189_v18).dtor_cf3).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v19)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_191_v20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_192_v21).dtor_cf1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v21).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_193_v22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_194_v23).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber(647)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_195_v24)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_196_v25).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_197_v26).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_198_v27)[_dafny.ZERO]).dtor_cf1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v27)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_198_v27)[_dafny.ONE]).dtor_cf1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v27)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_198_v27)[new BigNumber(2)]).dtor_cf1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v27)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_198_v27)[new BigNumber(3)]).dtor_cf1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v27)[new BigNumber(3)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_199_v28).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_255_v79).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_256_v80).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v81)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_258_v82).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v83.f7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_259_v83).f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v84).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65.f7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_261_v85).dtor_cf65).f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v85).dtor_cf65).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v86).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),_dafny.Seq.of(new BigNumber(598), _dafny.ONE, _dafny.ONE, _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v87).dtor_cf51));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v87).dtor_cf52));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_263_v87).dtor_cf53).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-1),_dafny.Seq.of(new BigNumber(598), _dafny.ONE, _dafny.ONE, _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_264_v88).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v89).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_340_v147).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_343_v148).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_345_v150).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_346_v151).equals(_dafny.MultiSet.fromElements(new BigNumber(829), _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_347_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_354_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC2(cf1, cf2) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D0(3);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC5(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false);
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
    static create_DC7(cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D2(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC9(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + this.cf6.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf9 === other.cf9;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC10(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ")";
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf12) {
      let $dt = new D4(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
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
    static create_DC12(cf13) {
      let $dt = new D5(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13;
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
    static create_DC14(cf15) {
      let $dt = new D6(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC13(cf14) {
      let $dt = new D6(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC14(false);
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
    static create_DC15(cf16) {
      let $dt = new D7(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf17) {
      let $dt = new D8(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC17(cf18) {
      let $dt = new D9(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC17" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18(cf19) {
      let $dt = new D10(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC18" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19;
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf21, cf22) {
      let $dt = new D11(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC19(cf20) {
      let $dt = new D11(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC20" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC19" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC20(false, _dafny.ZERO);
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
    static create_DC22(cf24, cf25, cf26, cf27, cf28) {
      let $dt = new D12(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC23(cf29) {
      let $dt = new D12(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf23) {
      let $dt = new D12(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC22" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + this.cf28.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC23" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC21" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf29 === other.cf29;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC22(_dafny.Map.Empty, _dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC24(cf30) {
      let $dt = new D13(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC24" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf32, cf33, cf34) {
      let $dt = new D14(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC25(cf31) {
      let $dt = new D14(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC26" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC25" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC26(false, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC28(cf36, cf37) {
      let $dt = new D15(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC27(cf35) {
      let $dt = new D15(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC29(cf38) {
      let $dt = new D15(2);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC28" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC27" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC29" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && this.cf37 === other.cf37;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC28(false, false);
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
    static create_DC31(cf40) {
      let $dt = new D16(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC32(cf41) {
      let $dt = new D16(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC30(cf39) {
      let $dt = new D16(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC31" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC32" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC30" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf41 === other.cf41;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC31(false);
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
    static create_DC34(cf43) {
      let $dt = new D17(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC33(cf42) {
      let $dt = new D17(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC35(cf44) {
      let $dt = new D17(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC34" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC33" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC35" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC34(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC37(cf46, cf47, cf48) {
      let $dt = new D18(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC36(cf45) {
      let $dt = new D18(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC37" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC36" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf45 === other.cf45;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC37(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC39(cf50) {
      let $dt = new D19(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC40(cf51, cf52, cf53) {
      let $dt = new D19(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC38(cf49) {
      let $dt = new D19(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC41(cf54) {
      let $dt = new D19(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC39" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC40" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC38" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC41" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC39(_dafny.Seq.of());
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
    static create_DC43(cf56, cf57, cf58) {
      let $dt = new D20(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC42(cf55) {
      let $dt = new D20(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC44(cf59) {
      let $dt = new D20(2);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC43" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC42" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC44" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf55 === other.cf55;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC43(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC46(cf61, cf62, cf63) {
      let $dt = new D21(0);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC47(cf64) {
      let $dt = new D21(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC45(cf60) {
      let $dt = new D21(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC46" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC47" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC45" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf61 === other.cf61 && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf64 === other.cf64;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC46(false, false, _dafny.ZERO);
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
    static create_DC49(cf66) {
      let $dt = new D22(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC48(cf65) {
      let $dt = new D22(1);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC49" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC48" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf65 === other.cf65;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC49(_dafny.Set.Empty);
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
    static create_DC50(cf67) {
      let $dt = new D23(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC50" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC51(cf68) {
      let $dt = new D24(0);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC51" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68);
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
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf70) {
      let $dt = new D25(0);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC54(cf71, cf72, cf73) {
      let $dt = new D25(1);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC52(cf69) {
      let $dt = new D25(2);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC53" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC54" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC52" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf71 === other.cf71 && this.cf72 === other.cf72 && this.cf73 === other.cf73;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf69, other.cf69);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC53(false);
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
    static create_DC56(cf75, cf76) {
      let $dt = new D26(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC57(cf77) {
      let $dt = new D26(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC58(cf78, cf79) {
      let $dt = new D26(2);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC55(cf74) {
      let $dt = new D26(3);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC59(cf80) {
      let $dt = new D26(4);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC58() { return this.$tag === 2; }
    get is_DC55() { return this.$tag === 3; }
    get is_DC59() { return this.$tag === 4; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC56" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC57" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC58" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 3) {
        return "D26.DC55" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 4) {
        return "D26.DC59" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75) && this.cf76 === other.cf76;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf77 === other.cf77;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf74 === other.cf74;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC56(_dafny.Seq.of(), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.T3 = class T3 {
  };

  $module.T4 = class T4 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this.f4 = _dafny.MultiSet.Empty;
      this._f0 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
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
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f12 = _dafny.Seq.of();
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f11, f12) {
      let _this = this;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-765),(_this).f11))).length);
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4(_module.__default.safeModuloInt(new BigNumber(31), new BigNumber(-303)));
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      if (!(!(!(false)))) {
        return (_dafny.MultiSet.fromElements(new BigNumber(-16), new BigNumber(436), new BigNumber(81), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(824)), function (_360_i0) {
          return new BigNumber(645);
        })).length)), new BigNumber(-807))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(484), new BigNumber(447)));
      } else {
        return (true) === (!(!(false)));
      }
    };
    fm7(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber(-399), new BigNumber((function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of (_dafny.Seq.of(_module.D1.create_DC5(true))).Elements) {
          let _361_v0 = _compr_21;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC5(true)), _361_v0)) {
            _coll21.add(_361_v0);
          }
        }
        return _coll21;
      }()).length));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.of(!((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(907),true)).Keys.Elements) {
          let _362_v0 = _compr_22;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(907),true)).contains(_362_v0)) {
            _coll22.add((_362_v0).minus(new BigNumber(85)));
          }
        }
        return _coll22;
      }()).length))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(46),!(true))).length), new BigNumber(873)))))).length);
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4(new BigNumber(-49));
    };
    fm14(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aam"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(143)), function (_363_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("hteems")))).length);
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      r0 = (_this).fm6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), function (_364_i0) {
        return new BigNumber(-615);
      }), globalState);
      let _365_v0;
      _365_v0 = _dafny.Seq.of(p0, p0, p0);
      _365_v0 = _dafny.Seq.Concat(_dafny.Seq.of(p0), _365_v0);
      let _366_v1;
      _366_v1 = _dafny.MultiSet.fromElements((_this).fm7(p0, globalState));
      let _rhs74 = p0;
      let _rhs75 = (_this).fm4(_dafny.Seq.of(p0), new BigNumber(325), _module.__default.safeModuloInt(p1, p1), _366_v1, globalState);
      let _lhs59 = globalState;
      r0 = _rhs74;
      _lhs59.f1 = _rhs75;
      let _367_v2;
      let _nw51 = Array((new BigNumber(9)).toNumber());
      _nw51[(_dafny.ZERO).toNumber()] = _module.__default.fm0(p0, globalState);
      _nw51[(_dafny.ONE).toNumber()] = !(p0);
      _nw51[(new BigNumber(2)).toNumber()] = !(!((_this).fm6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(578)), ((_368_p1) => function (_369_i1) {
        return _368_p1;
      })(p1)), globalState)));
      _nw51[(new BigNumber(3)).toNumber()] = p0;
      _nw51[(new BigNumber(4)).toNumber()] = p0;
      _nw51[(new BigNumber(5)).toNumber()] = p0;
      _nw51[(new BigNumber(6)).toNumber()] = p0;
      _nw51[(new BigNumber(7)).toNumber()] = p0;
      _nw51[(new BigNumber(8)).toNumber()] = false;
      _367_v2 = _nw51;
      let _370_v3;
      _370_v3 = _dafny.MultiSet.fromElements(_367_v2);
      let _371_v4;
      _371_v4 = _module.D0.create_DC0(p0);
      let _pat_let_tv0 = p2;
      let _pat_let_tv1 = p2;
      let _pat_let_tv2 = p2;
      let _pat_let_tv3 = p2;
      let _pat_let_tv4 = p2;
      let _pat_let_tv5 = p2;
      let _rhs76 = _dafny.MultiSet.fromElements(_367_v2, _367_v2, _367_v2);
      let _rhs77 = new BigNumber((function (_source9) {
        if (_source9.is_DC1) {
          return _pat_let_tv0;
        } else if (_source9.is_DC2) {
          let _372___mcc_h0 = (_source9).cf1;
          let _373___mcc_h1 = (_source9).cf2;
          let _374_cf2 = _373___mcc_h1;
          let _375_cf1 = _372___mcc_h0;
          return _pat_let_tv1;
        } else if (_source9.is_DC0) {
          let _376___mcc_h2 = (_source9).cf0;
          let _377_cf0 = _376___mcc_h2;
          return _dafny.Seq.Concat(_pat_let_tv2, _pat_let_tv3);
        } else {
          let _378___mcc_h3 = (_source9).cf3;
          let _379_cf3 = _378___mcc_h3;
          return _dafny.Seq.Concat(_pat_let_tv4, _pat_let_tv5);
        }
      }(_371_v4)).length);
      let _lhs60 = globalState;
      _370_v3 = _rhs76;
      _lhs60.f1 = _rhs77;
      let _380_v5;
      let _nw52 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
      _380_v5 = _nw52;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_380_v5).length))) {
        let _381_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_381_i2)) && ((_381_i2).isLessThan(new BigNumber((_380_v5).length))))) {
          (_380_v5)[(_381_i2)] = ((p0) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(179),p1)) : (_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Set.fromElements(p0, p0)).length))));
        }
      }
      let _382_v6;
      let _nw53 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _382_v6 = _nw53;
      let _383_v7;
      _383_v7 = _module.D2.create_DC7(p1, (_dafny.ZERO).minus((_this).fm14(globalState)));
      let _index69 = _module.__default.safeIndex(new BigNumber(308), new BigNumber((_382_v6).length));
      (_382_v6)[_index69] = (p1).minus((_383_v7).dtor_cf7);
      let _index70 = _module.__default.safeIndex(new BigNumber(308), new BigNumber((_382_v6).length));
      (_382_v6)[_index70] = new BigNumber((_370_v3).cardinality());
      r0 = p0;
      r1 = p2;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _384_v0;
      _384_v0 = true;
      let _385_v1;
      _385_v1 = _dafny.Seq.UnicodeFromString("mof");
      let _386_v2;
      _386_v2 = _dafny.Map.Empty.slice().updateUnsafe((!(_384_v0)) === (_384_v0),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _385_v1));
      _386_v2 = _386_v2;
      let _387_v3;
      _387_v3 = new BigNumber(-319);
      let _388_v4;
      _388_v4 = _dafny.Map.Empty.slice().updateUnsafe(_387_v3,_387_v3);
      let _389_v5;
      _389_v5 = _dafny.Map.Empty.slice().updateUnsafe(_384_v0,_388_v4);
      let _390_v6;
      _390_v6 = _dafny.Map.Empty.slice().updateUnsafe((_389_v5).Merge(_389_v5),true);
      if ((((_390_v6).contains(_dafny.Map.Empty.slice().updateUnsafe(_384_v0,_388_v4))) ? ((_390_v6).get(_dafny.Map.Empty.slice().updateUnsafe(_384_v0,_388_v4))) : (!(_384_v0)))) {
        let _391_v7;
        _391_v7 = _dafny.MultiSet.fromElements(_384_v0);
        _384_v0 = (((_391_v7).IsProperSubsetOf(_391_v7)) ? (!(_384_v0)) : (_384_v0));
        (globalState).f1 = _387_v3;
        let _392_v8;
        let _nw54 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _392_v8 = _nw54;
        let _rhs78 = _module.__default.safeModuloInt(_387_v3, _387_v3);
        let _rhs79 = _392_v8;
        _387_v3 = _rhs78;
        _392_v8 = _rhs79;
        let _393_v9;
        _393_v9 = _dafny.Seq.of(_384_v0);
        if (!_dafny.areEqual(_dafny.Seq.Concat(_393_v9, _393_v9), _393_v9)) {
          let _394_v10;
          _394_v10 = new _dafny.CodePoint('k'.codePointAt(0));
          _394_v10 = _394_v10;
          let _395_v11;
          let _nw55 = Array((new BigNumber(29)).toNumber()).fill(false);
          _395_v11 = _nw55;
          let _index71 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_395_v11).length));
          (_395_v11)[_index71] = _384_v0;
          let _index72 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_395_v11).length));
          (_395_v11)[_index72] = (_393_v9)[_module.__default.safeIndex((_dafny.ZERO).minus(_387_v3), new BigNumber((_393_v9).length))];
          _394_v10 = _394_v10;
          let _396_v12;
          _396_v12 = _dafny.Map.Empty.slice().updateUnsafe((_395_v11)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_395_v11).length))],_384_v0);
          let _index73 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_395_v11).length));
          (_395_v11)[_index73] = (((_396_v12).contains(false)) ? ((_396_v12).get(false)) : (!(!((_387_v3).isLessThan(_387_v3)))));
          let _index74 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_392_v8).length));
          (_392_v8)[_index74] = _387_v3;
          let _index75 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_392_v8).length));
          (_392_v8)[_index75] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_387_v3));
        } else {
          _384_v0 = !(_384_v0) || (_384_v0);
          let _397_v13;
          let _nw56 = Array((new BigNumber(27)).toNumber()).fill(_module.D2.Default());
          _397_v13 = _nw56;
          _397_v13 = _397_v13;
          _384_v0 = !(_384_v0);
          _384_v0 = !((_384_v0) && (_384_v0));
          let _398_v14;
          _398_v14 = new _dafny.CodePoint('k'.codePointAt(0));
          _398_v14 = _398_v14;
        }
        _384_v0 = true;
      } else {
        _389_v5 = _389_v5;
        (globalState).f1 = (_387_v3).multipliedBy(_387_v3);
        _385_v1 = _385_v1;
        let _399_v15;
        _399_v15 = _dafny.Map.Empty.slice().updateUnsafe(_384_v0,new BigNumber(-385));
        let _400_v16;
        _400_v16 = _dafny.Seq.of(new BigNumber(-967), new BigNumber((_399_v15).length));
        let _401_v17;
        _401_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_400_v16, _module.__default.safeIndex(_387_v3, new BigNumber((_400_v16).length)), _387_v3),_384_v0);
        let _402_v18;
        _402_v18 = _dafny.Set.fromElements(false, !(_384_v0));
        let _403_v19;
        _403_v19 = _dafny.Map.Empty.slice().updateUnsafe((_401_v17).Merge(_401_v17),!((_402_v18).IsProperSubsetOf(_402_v18)));
        _403_v19 = (_403_v19).update((_401_v17).Merge(_module.__default.fm15((_dafny.ZERO).minus(_387_v3), globalState)), _384_v0);
        let _404_v20;
        let _nw57 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _404_v20 = _nw57;
        let _405_v21;
        _405_v21 = _dafny.Seq.of(_400_v16, _400_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-573)), ((_406_v3) => function (_407_i0) {
          return _406_v3;
        })(_387_v3)));
        let _408_v22;
        let _nw58 = new _module.C0();
        _nw58.__ctor(_384_v0, _dafny.Seq.Concat(_405_v21, _405_v21));
        _408_v22 = _nw58;
        let _409_v23;
        _409_v23 = _dafny.Seq.of(_384_v0);
        let _rhs80 = _384_v0;
        let _rhs81 = _404_v20;
        let _rhs82 = (_409_v23)[_module.__default.safeIndex(_387_v3, new BigNumber((_409_v23).length))];
        let _rhs83 = _408_v22;
        _384_v0 = _rhs80;
        _404_v20 = _rhs81;
        _384_v0 = _rhs82;
        _408_v22 = _rhs83;
      }
      let _410_v24;
      _410_v24 = _dafny.Seq.of(_384_v0, _384_v0, !(_387_v3).isEqualTo(new BigNumber(-789)), (_387_v3).isLessThanOrEqualTo(_387_v3));
      if ((_410_v24)[_module.__default.safeIndex((_dafny.ZERO).minus(_387_v3), new BigNumber((_410_v24).length))]) {
        let _411_v25;
        _411_v25 = _dafny.Seq.of(new BigNumber((_388_v4).length), new BigNumber(547), new BigNumber(858), _387_v3, new BigNumber(288));
        let _412_v26;
        _412_v26 = _dafny.Seq.of(_411_v25, _411_v25);
        let _413_v27;
        let _nw59 = new _module.C0();
        _nw59.__ctor(_384_v0, _412_v26);
        _413_v27 = _nw59;
        _413_v27 = _413_v27;
        let _414_v28;
        _414_v28 = _dafny.MultiSet.fromElements(_384_v0, _384_v0, !(_384_v0));
        (globalState).f4 = _414_v28;
        let _415_v29;
        let _nw60 = Array((new BigNumber(28)).toNumber()).fill(false);
        _415_v29 = _nw60;
        let _index76 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_415_v29).length));
        (_415_v29)[_index76] = ((_384_v0) ? ((_413_v27).f11) : (_384_v0));
        let _416_v30;
        _416_v30 = _module.D1.create_DC5((_413_v27).f11);
        let _index77 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_415_v29).length));
        let _rhs84 = (_416_v30).dtor_cf5;
        let _rhs85 = _384_v0;
        let _rhs86 = (_413_v27).f11;
        let _lhs61 = _415_v29;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_415_v29).length));
        _lhs61[_lhs62] = _rhs84;
        _384_v0 = _rhs85;
        _384_v0 = _rhs86;
        let _index78 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_415_v29).length));
        (_415_v29)[_index78] = (_this).fm6((_413_v27.f12)[_module.__default.safeIndex(_387_v3, new BigNumber((_413_v27.f12).length))], globalState);
        _384_v0 = _384_v0;
      } else {
        let _417_v31;
        _417_v31 = _dafny.Seq.of(_410_v24);
        let _418_v32;
        _418_v32 = _dafny.Seq.of(_387_v3);
        _388_v4 = (_388_v4).update((_dafny.ZERO).minus(new BigNumber((_417_v31).length)), (new BigNumber((_dafny.Seq.of(_418_v32)).length)).plus((_dafny.ZERO).minus(_387_v3)));
        let _419_v33;
        let _init6 = ((_420_v0) => function (_421_i1) {
          return _420_v0;
        })(_384_v0);
        let _nw61 = Array((new BigNumber(7)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw61.length); _i0_6++) {
          _nw61[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _419_v33 = _nw61;
        let _index79 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length));
        (_419_v33)[_index79] = _384_v0;
        let _index80 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length));
        let _rhs87 = !(!(_384_v0)) || (_384_v0);
        let _rhs88 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_422_i2) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }), _385_v1);
        let _rhs89 = _385_v1;
        let _rhs90 = new BigNumber(535);
        let _lhs63 = _419_v33;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length));
        _lhs63[_lhs64] = _rhs87;
        _385_v1 = _rhs88;
        _385_v1 = _rhs89;
        _387_v3 = _rhs90;
        let _423_v34;
        let _nw62 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _423_v34 = _nw62;
        let _index81 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_423_v34).length));
        (_423_v34)[_index81] = (_387_v3).plus(_387_v3);
        let _424_v35;
        _424_v35 = _dafny.Map.Empty.slice().updateUnsafe((_419_v33)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length))],(_dafny.ZERO).minus(_387_v3));
        let _425_v36;
        _425_v36 = _dafny.MultiSet.fromElements(_388_v4, _388_v4, _dafny.Map.Empty.slice().updateUnsafe(_387_v3,new BigNumber(-861)));
        let _426_v37;
        _426_v37 = new _dafny.CodePoint('e'.codePointAt(0));
        let _427_v38;
        _427_v38 = _dafny.Map.Empty.slice().updateUnsafe(_426_v37,_387_v3);
        let _index82 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_423_v34).length));
        (_423_v34)[_index82] = (((_424_v35).contains((_425_v36).IsDisjointFrom(_425_v36))) ? ((_424_v35).get((_425_v36).IsDisjointFrom(_425_v36))) : (((((_427_v38).contains(_426_v37)) ? ((_427_v38).get(_426_v37)) : (_387_v3))).minus(_387_v3)));
        let _index83 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length));
        (_419_v33)[_index83] = (_419_v33)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length))];
        let _428_v39;
        _428_v39 = _module.D1.create_DC5((_419_v33)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_419_v33).length))]);
        let _429_v40;
        _429_v40 = _dafny.Map.Empty.slice().updateUnsafe(_387_v3,_428_v39);
        let _430_v41;
        _430_v41 = _dafny.Map.Empty.slice().updateUnsafe((_423_v34)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_423_v34).length))],_385_v1);
        _429_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_430_v41).length),_428_v39);
      }
      let _431_v42;
      _431_v42 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(299)), function (_432_i3) {
        return new BigNumber(-622);
      }));
      let _433_v43;
      let _nw63 = new _module.C0();
      _nw63.__ctor(_384_v0, _dafny.Seq.Concat(_431_v42, _431_v42));
      _433_v43 = _nw63;
      _384_v0 = false;
      let _434_i4;
      _434_i4 = _dafny.ZERO;
      L5: {
        while (!(_module.__default.fm0(!((_433_v43).f11), globalState))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_434_i4)) {
              break L5;
            }
            _434_i4 = (_434_i4).plus(_dafny.ONE);
            (globalState).f1 = _387_v3;
            _384_v0 = _384_v0;
            if ((_433_v43).f11) {
              let _435_v44;
              let _nw64 = Array((new BigNumber(14)).toNumber()).fill(false);
              _435_v44 = _nw64;
              let _436_v45;
              _436_v45 = _dafny.MultiSet.fromElements(_435_v44, _435_v44, _435_v44);
              _384_v0 = ((_436_v45).Union(_436_v45)).IsSubsetOf((_dafny.MultiSet.fromElements(_435_v44)).Difference(_436_v45));
              let _index84 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_435_v44).length));
              (_435_v44)[_index84] = false;
              let _437_v46;
              _437_v46 = _dafny.Set.fromElements((_433_v43).f11);
              let _index85 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_435_v44).length));
              (_435_v44)[_index85] = (((_433_v43).f11) ? ((_437_v46).IsDisjointFrom(_437_v46)) : (_384_v0));
              let _rhs91 = _385_v1;
              let _rhs92 = _435_v44;
              let _rhs93 = new BigNumber(-241);
              let _lhs65 = globalState;
              _385_v1 = _rhs91;
              _435_v44 = _rhs92;
              _lhs65.f1 = _rhs93;
              let _438_v47;
              let _nw65 = new _module.C0();
              _nw65.__ctor(_384_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(437)), ((_439_v3) => function (_440_i5) {
                return _dafny.Seq.of(_439_v3, _439_v3);
              })(_387_v3)));
              _438_v47 = _nw65;
              _387_v3 = _387_v3;
            } else {
              let _441_v48;
              _441_v48 = new _dafny.CodePoint('p'.codePointAt(0));
              let _442_v49;
              _442_v49 = _dafny.Map.Empty.slice().updateUnsafe(_441_v48,_387_v3);
              (globalState).f1 = (((_442_v49).contains(_441_v48)) ? ((_442_v49).get(_441_v48)) : (_387_v3));
              _441_v48 = _441_v48;
              r0 = (_this).fm7((_433_v43).f11, globalState);
              r0 = (_this).fm14(globalState);
              let _443_v50;
              let _nw66 = Array((new BigNumber(20)).toNumber()).fill(false);
              _443_v50 = _nw66;
              let _444_v51;
              _444_v51 = _dafny.Map.Empty.slice().updateUnsafe(_443_v50,new _dafny.CodePoint('w'.codePointAt(0)));
              let _445_v52;
              _445_v52 = _module.D0.create_DC2(_444_v51, _441_v48);
              _445_v52 = _445_v52;
            }
            _433_v43 = _433_v43;
          }
        }
      }
      r0 = _387_v3;
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _446_v0;
      _446_v0 = true;
      _446_v0 = _446_v0;
      let _447_v1;
      _447_v1 = _dafny.Seq.UnicodeFromString("pvlac");
      _447_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(444)), function (_448_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      });
      if (p2) {
        let _449_v2;
        _449_v2 = new BigNumber(188);
        let _450_v3;
        _450_v3 = _dafny.Set.fromElements(_449_v2);
        let _451_v4;
        _451_v4 = _dafny.Seq.of(p2, _446_v0, false);
        let _452_v5;
        _452_v5 = _dafny.Map.Empty.slice().updateUnsafe(true,_449_v2);
        let _453_v6;
        let _init7 = function (_454_i1) {
          return true;
        };
        let _nw67 = Array((new BigNumber(22)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw67.length); _i0_7++) {
          _nw67[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _453_v6 = _nw67;
        let _455_v7;
        _455_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,_453_v6);
        let _456_v8;
        _456_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_452_v5).length),_455_v7);
        let _457_v9;
        _457_v9 = _dafny.Map.Empty.slice().updateUnsafe(_449_v2,_449_v2);
        let _458_v10;
        _458_v10 = _dafny.Seq.of(_451_v4);
        let _459_v11;
        let _nw68 = Array((new BigNumber(22)).toNumber());
        _nw68[(_dafny.ZERO).toNumber()] = p0;
        _nw68[(_dafny.ONE).toNumber()] = false;
        _nw68[(new BigNumber(2)).toNumber()] = p0;
        _nw68[(new BigNumber(3)).toNumber()] = p0;
        _nw68[(new BigNumber(4)).toNumber()] = p0;
        _nw68[(new BigNumber(5)).toNumber()] = (_450_v3).equals(_dafny.Set.fromElements(_449_v2));
        _nw68[(new BigNumber(6)).toNumber()] = p0;
        _nw68[(new BigNumber(7)).toNumber()] = p2;
        _nw68[(new BigNumber(8)).toNumber()] = p0;
        _nw68[(new BigNumber(9)).toNumber()] = (_module.__default.fm16(new BigNumber(71), globalState)).IsDisjointFrom(_450_v3);
        _nw68[(new BigNumber(10)).toNumber()] = _446_v0;
        _nw68[(new BigNumber(11)).toNumber()] = !(((p2) ? (_446_v0) : ((_451_v4)[_module.__default.safeIndex(_449_v2, new BigNumber((_451_v4).length))])));
        _nw68[(new BigNumber(12)).toNumber()] = p0;
        _nw68[(new BigNumber(13)).toNumber()] = _446_v0;
        _nw68[(new BigNumber(14)).toNumber()] = (false) || (p2);
        _nw68[(new BigNumber(15)).toNumber()] = !((((_456_v8).contains((((_457_v9).contains(_449_v2)) ? ((_457_v9).get(_449_v2)) : (_449_v2)))) ? ((_456_v8).get((((_457_v9).contains(_449_v2)) ? ((_457_v9).get(_449_v2)) : (_449_v2)))) : (_dafny.Map.Empty.slice().updateUnsafe(p1,_453_v6)))).contains(p1);
        _nw68[(new BigNumber(16)).toNumber()] = !(p0);
        _nw68[(new BigNumber(17)).toNumber()] = _446_v0;
        _nw68[(new BigNumber(18)).toNumber()] = _module.__default.fm0(false, globalState);
        _nw68[(new BigNumber(19)).toNumber()] = _dafny.Seq.IsPrefixOf(_451_v4, (_458_v10)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_458_v10).length))]);
        _nw68[(new BigNumber(20)).toNumber()] = p0;
        _nw68[(new BigNumber(21)).toNumber()] = p0;
        _459_v11 = _nw68;
        let _index86 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_459_v11).length));
        (_459_v11)[_index86] = true;
        let _index87 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_459_v11).length));
        (_459_v11)[_index87] = _446_v0;
        let _460_v12;
        let _nw69 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _460_v12 = _nw69;
        let _index88 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length));
        (_460_v12)[_index88] = (new BigNumber((_dafny.Seq.UnicodeFromString("cf")).length)).multipliedBy(new BigNumber((_dafny.Seq.of(_449_v2, _449_v2, _449_v2)).length));
        let _index89 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length));
        (_460_v12)[_index89] = _449_v2;
        (globalState).f1 = (_460_v12)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length))];
        let _461_v13;
        _461_v13 = new _dafny.CodePoint('k'.codePointAt(0));
        _447_v1 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_447_v1, _447_v1), _447_v1), _module.__default.safeIndex(_449_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_447_v1, _447_v1), _447_v1)).length)), _461_v13);
        let _462_v14;
        _462_v14 = _module.D1.create_DC5(p2);
        let _463_v15;
        _463_v15 = _dafny.Map.Empty.slice().updateUnsafe(_462_v14,p2);
        let _464_v16;
        _464_v16 = _dafny.MultiSet.fromElements(_461_v13, _461_v13, _461_v13);
        let _index90 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length));
        let _index91 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_459_v11).length));
        let _rhs94 = (new BigNumber((_464_v16).cardinality())).multipliedBy((_460_v12)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length))]);
        let _rhs95 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_449_v2, ((_460_v12)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length))]).minus(new BigNumber((_457_v9).length))));
        let _rhs96 = (_463_v15).Merge(_dafny.Map.Empty.slice().updateUnsafe(_462_v14,p2));
        let _rhs97 = _446_v0;
        let _lhs66 = globalState;
        let _lhs67 = _460_v12;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_460_v12).length));
        let _lhs69 = _459_v11;
        let _lhs70 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_459_v11).length));
        _lhs66.f1 = _rhs94;
        _lhs67[_lhs68] = _rhs95;
        _463_v15 = _rhs96;
        _lhs69[_lhs70] = _rhs97;
      } else {
        let _465_v17;
        _465_v17 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-754)), function (_466_i2) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }));
        _446_v0 = (_465_v17).IsSubsetOf(_465_v17);
        let _467_v18;
        _467_v18 = _dafny.Set.fromElements(_446_v0);
        let _468_v19;
        let _nw70 = Array((new BigNumber(3)).toNumber()).fill(false);
        _468_v19 = _nw70;
        let _index92 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        (_468_v19)[_index92] = p0;
        let _469_v20;
        _469_v20 = new BigNumber(192);
        let _470_v21;
        _470_v21 = _module.D2.create_DC8(p0);
        let _index93 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        let _rhs98 = !(_446_v0);
        let _rhs99 = (_467_v18).Difference((_467_v18).Difference(_module.__default.fm17(_469_v20, _446_v0, globalState)));
        let _rhs100 = !(_446_v0) || ((_469_v20).isLessThan(_469_v20));
        let _rhs101 = ((_446_v0) ? ((((function (_pat_let0_0) {
          return function (_471_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_472_dt__update_hcf9_h0) {
                return _module.D2.create_DC8(_472_dt__update_hcf9_h0);
              }(_pat_let1_0);
            }(true);
          }(_pat_let0_0);
        }(_470_v21)).dtor_cf9) ? (p2) : (p0))) : (p2));
        let _rhs102 = (_dafny.ZERO).minus(_module.__default.fm3(_446_v0, p2, globalState));
        let _lhs71 = _468_v19;
        let _lhs72 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        let _lhs73 = globalState;
        _446_v0 = _rhs98;
        _467_v18 = _rhs99;
        _lhs71[_lhs72] = _rhs100;
        _446_v0 = _rhs101;
        _lhs73.f1 = _rhs102;
        let _473_v22;
        _473_v22 = _dafny.Map.Empty.slice().updateUnsafe(_469_v20,_469_v20);
        let _474_v23;
        _474_v23 = new _dafny.CodePoint('e'.codePointAt(0));
        let _475_v24;
        _475_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber((_447_v1).length), false, _469_v20, globalState),_dafny.Seq.update(_447_v1, _module.__default.safeIndex(_469_v20, new BigNumber((_447_v1).length)), _474_v23));
        _473_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_469_v20, new BigNumber(((((_475_v24).contains(_447_v1)) ? ((_475_v24).get(_447_v1)) : (_447_v1))).length)),_469_v20);
        let _476_v25;
        _476_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(211),p2);
        let _477_v26;
        _477_v26 = _476_v25;
        let _478_v27;
        let _nw71 = Array((new BigNumber(28)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = _477_v26;
        _nw71[(_dafny.ONE).toNumber()] = _476_v25;
        _nw71[(new BigNumber(2)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(3)).toNumber()] = _476_v25;
        _nw71[(new BigNumber(4)).toNumber()] = _476_v25;
        _nw71[(new BigNumber(5)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(6)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(7)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(8)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(9)).toNumber()] = _module.__default.fm18(_469_v20, globalState);
        _nw71[(new BigNumber(10)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(11)).toNumber()] = _module.__default.fm18(_469_v20, globalState);
        _nw71[(new BigNumber(12)).toNumber()] = _476_v25;
        _nw71[(new BigNumber(13)).toNumber()] = _476_v25;
        _nw71[(new BigNumber(14)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(15)).toNumber()] = _module.__default.fm18(new BigNumber((_447_v1).length), globalState);
        _nw71[(new BigNumber(16)).toNumber()] = _476_v25;
        _nw71[(new BigNumber(17)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(18)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(19)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(20)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(21)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(22)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(23)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(24)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(25)).toNumber()] = _477_v26;
        _nw71[(new BigNumber(26)).toNumber()] = ((_446_v0) ? (_477_v26) : (_477_v26));
        _nw71[(new BigNumber(27)).toNumber()] = _477_v26;
        _478_v27 = _nw71;
        _478_v27 = _478_v27;
        let _479_v28;
        _479_v28 = _dafny.Seq.of(_469_v20, _469_v20, (_469_v20).minus(_module.__default.fm3(!(p2), _446_v0, globalState)));
        let _480_v29;
        _480_v29 = _dafny.Seq.of(_479_v28, _479_v28);
        let _481_v30;
        let _nw72 = new _module.C0();
        _nw72.__ctor((_this).fm6(_479_v28, globalState), _480_v29);
        _481_v30 = _nw72;
        let _index94 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        let _index95 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        let _rhs103 = _446_v0;
        let _rhs104 = _dafny.Seq.update(_dafny.Seq.Concat(_479_v28, _479_v28), _module.__default.safeIndex(_469_v20, new BigNumber((_dafny.Seq.Concat(_479_v28, _479_v28)).length)), (_dafny.ZERO).minus(new BigNumber((_447_v1).length)));
        let _rhs105 = _481_v30;
        let _rhs106 = (p1).IsSubsetOf(p1);
        let _lhs74 = _468_v19;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        let _lhs76 = _468_v19;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_468_v19).length));
        _lhs74[_lhs75] = _rhs103;
        _479_v28 = _rhs104;
        _481_v30 = _rhs105;
        _lhs76[_lhs77] = _rhs106;
      }
      let _482_v31;
      _482_v31 = new BigNumber(174);
      let _483_v32;
      _483_v32 = _dafny.Seq.of(_482_v31, new BigNumber((p1).cardinality()));
      let _484_v33;
      _484_v33 = _dafny.Seq.of(_483_v32);
      let _485_v34;
      let _nw73 = new _module.C0();
      _nw73.__ctor(p0, _484_v33);
      _485_v34 = _nw73;
      let _486_v35;
      let _nw74 = Array((new BigNumber(12)).toNumber()).fill([]);
      _486_v35 = _nw74;
      let _487_v36;
      _487_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_446_v0),_447_v1);
      let _488_v37;
      let _nw75 = Array((new BigNumber(8)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = _447_v1;
      _nw75[(_dafny.ONE).toNumber()] = (((_487_v36).contains(p2)) ? ((_487_v36).get(p2)) : (_447_v1));
      _nw75[(new BigNumber(2)).toNumber()] = _447_v1;
      _nw75[(new BigNumber(3)).toNumber()] = _447_v1;
      _nw75[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("cucmgtuti");
      _nw75[(new BigNumber(5)).toNumber()] = _447_v1;
      _nw75[(new BigNumber(6)).toNumber()] = ((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(717)), function (_489_i3) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })) : (_dafny.Seq.UnicodeFromString("n")));
      _nw75[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_482_v31, p2, _482_v31, globalState);
      _488_v37 = _nw75;
      let _index96 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_488_v37).length));
      (_488_v37)[_index96] = _447_v1;
      let _490_v38;
      _490_v38 = _dafny.Map.Empty.slice().updateUnsafe(_446_v0,false);
      let _index97 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_488_v37).length));
      let _rhs107 = !(!(_490_v38).contains(!((_485_v34).f11) || (_446_v0)));
      let _rhs108 = _486_v35;
      let _rhs109 = _447_v1;
      let _rhs110 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_482_v31, _482_v31));
      let _lhs78 = _488_v37;
      let _lhs79 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_488_v37).length));
      _446_v0 = _rhs107;
      _486_v35 = _rhs108;
      _lhs78[_lhs79] = _rhs109;
      _482_v31 = _rhs110;
      let _491_v39;
      let _nw76 = Array((new BigNumber(26)).toNumber()).fill([]);
      _491_v39 = _nw76;
      let _492_v40;
      let _init8 = function (_493_i4) {
        return (_493_i4).minus(new BigNumber(410));
      };
      let _nw77 = Array((new BigNumber(3)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw77.length); _i0_8++) {
        _nw77[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _492_v40 = _nw77;
      let _index98 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_491_v39).length));
      (_491_v39)[_index98] = _492_v40;
      let _index99 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_491_v39).length));
      let _init9 = ((_494_v31) => function (_495_i5) {
        return (_495_i5).minus(_494_v31);
      })(_482_v31);
      let _nw78 = Array((new BigNumber(26)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
        _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      (_491_v39)[_index99] = _nw78;
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
    }
    _parentTraits() {
      return [_module.T4, _module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f5, f6) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), _dafny.Seq.UnicodeFromString("vlmbv"))).length)).isEqualTo((new BigNumber(619)).minus(new BigNumber((_dafny.Seq.of(new BigNumber(721))).length)));
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(292);
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm7(p0, globalState) {
      let _this = this;
      return (new BigNumber(588)).multipliedBy(new BigNumber(737));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((new BigNumber(212)).minus(new BigNumber(663)))).plus(new BigNumber(((_dafny.Set.fromElements((_this).f5)).Union(_dafny.Set.fromElements((_this).f5))).length));
    };
    fm5(p0, globalState) {
      let _this = this;
      let _source10 = ((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(63),new BigNumber((_dafny.Seq.of(false)).length))).length),(_this).f6)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(136),(_this).f6)));
      let _496___mcc_h0 = _source10;
      let _497_cf12 = _496___mcc_h0;
      return _module.D1.create_DC4(new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6)).length));
    };
    fm20(p0, p1, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f6)).Difference(_dafny.MultiSet.fromElements((_this).f6));
    };
    fm21(globalState) {
      let _this = this;
      return _dafny.Set.fromElements(false);
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D2.Default();
      let r1 = _dafny.MultiSet.Empty;
      let _498_i0;
      _498_i0 = _dafny.ZERO;
      L6: {
        while (!(p2)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_498_i0)) {
              break L6;
            }
            _498_i0 = (_498_i0).plus(_dafny.ONE);
            if ((_this).f6) {
              let _499_v0;
              _499_v0 = _dafny.Seq.UnicodeFromString("ljybtapti");
              let _500_v1;
              _500_v1 = _module.D2.create_DC8(p2);
              let _501_v2;
              _501_v2 = _dafny.Set.fromElements((_this).f6);
              let _502_v3;
              let _nw79 = Array((new BigNumber(9)).toNumber());
              _nw79[(_dafny.ZERO).toNumber()] = (_this).f6;
              _nw79[(_dafny.ONE).toNumber()] = !_dafny.Seq.contains(_499_v0, (_this).f5);
              _nw79[(new BigNumber(2)).toNumber()] = p2;
              _nw79[(new BigNumber(3)).toNumber()] = (_500_v1).dtor_cf9;
              _nw79[(new BigNumber(4)).toNumber()] = (_this).f6;
              _nw79[(new BigNumber(5)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber((_module.__default.fm2((_this).f6, (_this).f6, globalState)).length), new BigNumber((p1).length))];
              _nw79[(new BigNumber(6)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("fllshammk"), _module.__default.fm22(_501_v2, new BigNumber(102), (_this).f6, (_this).f6, globalState));
              _nw79[(new BigNumber(7)).toNumber()] = !((_this).f6);
              _nw79[(new BigNumber(8)).toNumber()] = p2;
              _502_v3 = _nw79;
              let _index100 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length));
              (_502_v3)[_index100] = p2;
              let _503_v4;
              _503_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
              let _504_v5;
              _504_v5 = _dafny.Seq.of(_503_v4);
              let _index101 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length));
              (_502_v3)[_index101] = !((_504_v5)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dcgthjfh")).length)), new BigNumber((_504_v5).length))]).contains((_this).f6);
              (globalState).f1 = p3;
              let _505_v6;
              _505_v6 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
              let _506_v7;
              _506_v7 = _dafny.Map.Empty.slice().updateUnsafe((((_505_v6).contains(p3)) ? ((_505_v6).get(p3)) : (p0)),(_this).f6);
              let _507_v8;
              _507_v8 = _dafny.Map.Empty.slice().updateUnsafe(_501_v2,(_502_v3)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length))]);
              let _508_v9;
              _508_v9 = _dafny.MultiSet.fromElements(p2, p2, (_this).f6, (_this).fm10(_dafny.MultiSet.FromArray(p1), (_507_v8).update(_501_v2, p2), p3, globalState), (_502_v3)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length))]);
              _506_v7 = (_506_v7).update(p0, (_this).fm10(_508_v9, _507_v8, p3, globalState));
              _499_v0 = _dafny.Seq.Concat(_499_v0, _module.__default.fm1(new BigNumber(228), (_502_v3)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length))], new BigNumber(343), globalState));
              let _509_v10;
              _509_v10 = _dafny.Seq.of(_502_v3, _502_v3);
              let _index102 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_502_v3).length));
              (_502_v3)[_index102] = _dafny.Seq.IsPrefixOf(_509_v10, _dafny.Seq.of(_502_v3, _502_v3));
            } else {
              let _510_v11;
              _510_v11 = _dafny.MultiSet.fromElements(p3, new BigNumber(53), p3);
              let _index103 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length));
              (p0)[_index103] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(p2),new BigNumber((_510_v11).cardinality())),(_this).f6)).length), p3));
              let _index104 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length));
              (p0)[_index104] = new BigNumber(325);
              let _index105 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length));
              (p0)[_index105] = (p0)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length))];
              let _511_v12;
              _511_v12 = _dafny.Seq.UnicodeFromString("pxu");
              let _512_v13;
              _512_v13 = _module.D0.create_DC1();
              let _513_v14;
              let _nw80 = Array((new BigNumber(23)).toNumber());
              _nw80[(_dafny.ZERO).toNumber()] = _512_v13;
              _nw80[(_dafny.ONE).toNumber()] = _512_v13;
              _nw80[(new BigNumber(2)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(3)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(4)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(5)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(6)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(7)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(8)).toNumber()] = _module.__default.fm23(globalState);
              _nw80[(new BigNumber(9)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(10)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(11)).toNumber()] = ((!((_this).f6)) ? (_512_v13) : (_512_v13));
              _nw80[(new BigNumber(12)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(13)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(14)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(15)).toNumber()] = _module.__default.fm23(globalState);
              _nw80[(new BigNumber(16)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(17)).toNumber()] = ((p2) ? (_512_v13) : (_512_v13));
              _nw80[(new BigNumber(18)).toNumber()] = _module.D0.create_DC1();
              _nw80[(new BigNumber(19)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(20)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(21)).toNumber()] = _512_v13;
              _nw80[(new BigNumber(22)).toNumber()] = _512_v13;
              _513_v14 = _nw80;
              let _index106 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_513_v14).length));
              (_513_v14)[_index106] = _512_v13;
              let _514_v15;
              _514_v15 = true;
              let _index107 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_513_v14).length));
              let _index108 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length));
              let _rhs111 = _511_v12;
              let _rhs112 = (p0)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length))];
              let _rhs113 = _512_v13;
              let _rhs114 = _514_v15;
              let _rhs115 = (p3).minus(p3);
              let _lhs80 = globalState;
              let _lhs81 = _513_v14;
              let _lhs82 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_513_v14).length));
              let _lhs83 = p0;
              let _lhs84 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((p0).length));
              _511_v12 = _rhs111;
              _lhs80.f1 = _rhs112;
              _lhs81[_lhs82] = _rhs113;
              _514_v15 = _rhs114;
              _lhs83[_lhs84] = _rhs115;
              let _515_v16;
              _515_v16 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ukntgl")).length));
              let _516_v17;
              _516_v17 = _dafny.Seq.of(_515_v16);
              let _517_v18;
              let _nw81 = new _module.C0();
              _nw81.__ctor(_514_v15, _516_v17);
              _517_v18 = _nw81;
              _514_v15 = (_517_v18).f11;
            }
            let _518_v19;
            _518_v19 = false;
            let _519_v20;
            _519_v20 = _dafny.Map.Empty.slice().updateUnsafe(_518_v19,!(_518_v19));
            let _520_v21;
            _520_v21 = _dafny.Map.Empty.slice().updateUnsafe(_518_v19,p3);
            let _521_v22;
            _521_v22 = _dafny.MultiSet.fromElements(p3);
            let _522_v23;
            _522_v23 = _dafny.Seq.of(_521_v22);
            let _523_v24;
            _523_v24 = _dafny.Seq.UnicodeFromString("pvirq");
            let _524_v25;
            _524_v25 = _module.D1.create_DC5(false);
            let _525_v26;
            _525_v26 = _dafny.Seq.of(p3, p3, p3, p3);
            let _526_v27;
            _526_v27 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), ((_527_p3) => function (_528_i1) {
              return _527_p3;
            })(p3)), _525_v26);
            let _529_v28;
            let _nw82 = new _module.C0();
            _nw82.__ctor((_this).f6, _526_v27);
            _529_v28 = _nw82;
            let _530_v29;
            _530_v29 = _dafny.Map.Empty.slice().updateUnsafe(p2,_529_v28);
            let _531_v30;
            _531_v30 = _dafny.MultiSet.fromElements(_530_v29);
            let _532_v31;
            let _nw83 = Array((new BigNumber(17)).toNumber()).fill(false);
            _532_v31 = _nw83;
            let _533_v32;
            _533_v32 = _dafny.MultiSet.fromElements(_532_v31, _532_v31);
            let _534_v33;
            let _nw84 = Array((new BigNumber(21)).toNumber());
            _nw84[(_dafny.ZERO).toNumber()] = _518_v19;
            _nw84[(_dafny.ONE).toNumber()] = (((_519_v20).contains(_518_v19)) ? ((_519_v20).get(_518_v19)) : (p2));
            _nw84[(new BigNumber(2)).toNumber()] = (((_519_v20).contains((_this).f6)) ? ((_519_v20).get((_this).f6)) : ((_this).f6));
            _nw84[(new BigNumber(3)).toNumber()] = !(_518_v19) || (_518_v19);
            _nw84[(new BigNumber(4)).toNumber()] = (new BigNumber(-89)).isLessThanOrEqualTo(new BigNumber((p1).length));
            _nw84[(new BigNumber(5)).toNumber()] = false;
            _nw84[(new BigNumber(6)).toNumber()] = !(_module.__default.fm0(p2, globalState));
            _nw84[(new BigNumber(7)).toNumber()] = (p3).isLessThan((_this).fm8(p2, (_this).f6, p2, globalState));
            _nw84[(new BigNumber(8)).toNumber()] = false;
            _nw84[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p2,p3)).equals(_520_v21);
            _nw84[(new BigNumber(10)).toNumber()] = p2;
            _nw84[(new BigNumber(11)).toNumber()] = (_521_v22).IsSubsetOf(((_522_v23)[_module.__default.safeIndex(new BigNumber((_523_v24).length), new BigNumber((_522_v23).length))]).update(p3, _module.__default.abs(new BigNumber(108))));
            _nw84[(new BigNumber(12)).toNumber()] = (_524_v25).dtor_cf5;
            _nw84[(new BigNumber(13)).toNumber()] = ((!(_518_v19)) ? (false) : ((_this).f6));
            _nw84[(new BigNumber(14)).toNumber()] = (p3).isLessThanOrEqualTo(p3);
            _nw84[(new BigNumber(15)).toNumber()] = !(_531_v30).contains(_530_v29);
            _nw84[(new BigNumber(16)).toNumber()] = _518_v19;
            _nw84[(new BigNumber(17)).toNumber()] = _518_v19;
            _nw84[(new BigNumber(18)).toNumber()] = (_529_v28).f11;
            _nw84[(new BigNumber(19)).toNumber()] = !((new BigNumber((_533_v32).cardinality())).isLessThan(p3));
            _nw84[(new BigNumber(20)).toNumber()] = ((_521_v22).update(p3, _module.__default.abs(p3))).IsDisjointFrom(_521_v22);
            _534_v33 = _nw84;
            let _index109 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length));
            (_534_v33)[_index109] = (_529_v28).f11;
            let _index110 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length));
            let _rhs116 = _518_v19;
            let _rhs117 = ((_518_v19) && ((_529_v28).f11)) || (_518_v19);
            let _rhs118 = p3;
            let _rhs119 = p2;
            let _rhs120 = (new BigNumber(((((_this).f6) ? (_523_v24) : (_523_v24))).length)).plus(p3);
            let _lhs85 = _534_v33;
            let _lhs86 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length));
            let _lhs87 = globalState;
            let _lhs88 = globalState;
            _518_v19 = _rhs116;
            _lhs85[_lhs86] = _rhs117;
            _lhs87.f1 = _rhs118;
            _518_v19 = _rhs119;
            _lhs88.f1 = _rhs120;
            let _index111 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length));
            (_534_v33)[_index111] = (_529_v28).f11;
            if ((_this).f6) {
              (globalState).f1 = p3;
              let _535_v34;
              let _nw85 = new _module.C1();
              _nw85.__ctor();
              _535_v34 = _nw85;
              (globalState).f1 = ((p3).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f6), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.of((_this).f6)).length)), (_this).f6)).length))).multipliedBy(p3);
              let _536_v35;
              let _init10 = ((_537_v24) => function (_538_i2) {
                return _537_v24;
              })(_523_v24);
              let _nw86 = Array((new BigNumber(14)).toNumber());
              for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw86.length); _i0_10++) {
                _nw86[_i0_10] = _init10(new BigNumber(_i0_10));
              }
              _536_v35 = _nw86;
              let _539_v36;
              _539_v36 = _module.D0.create_DC0(_518_v19);
              let _540_v37;
              _540_v37 = _dafny.Map.Empty.slice().updateUnsafe(_539_v36,_518_v19);
              let _index112 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_536_v35).length));
              (_536_v35)[_index112] = (((((_540_v37).contains(_539_v36)) ? ((_540_v37).get(_539_v36)) : ((_534_v33)[_module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length))]))) ? (_523_v24) : (_dafny.Seq.UnicodeFromString("btwvxfgd")));
              let _index113 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_536_v35).length));
              (_536_v35)[_index113] = _523_v24;
              let _541_v38;
              _541_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
              _541_v38 = _541_v38;
            } else {
              let _542_v39;
              _542_v39 = new _dafny.CodePoint('u'.codePointAt(0));
              _542_v39 = (_this).f5;
              let _543_v40;
              _543_v40 = _dafny.Set.fromElements((_534_v33)[_module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length))]);
              let _544_v41;
              let _545_v42;
              let _out20;
              let _out21;
              let _outcollector9 = _module.__default.m0(_534_v33, _module.__default.fm22(_543_v40, p3, p2, (_529_v28).f11, globalState), (_529_v28).f11, (_this).f5, globalState);
              _out20 = _outcollector9[0];
              _out21 = _outcollector9[1];
              _544_v41 = _out20;
              _545_v42 = _out21;
              let _546_v43;
              _546_v43 = _dafny.MultiSet.fromElements((_529_v28).f11, _545_v42);
              let _547_v44;
              _547_v44 = _dafny.MultiSet.fromElements(((_546_v43).update((_529_v28).f11, _module.__default.abs(p3))).update((_529_v28).f11, _module.__default.abs((_this).fm4(p1, new BigNumber((_546_v43).cardinality()), p3, _521_v22, globalState))), (_this).fm20(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-7)), ((_548_v39) => function (_549_i3) {
                return _548_v39;
              })(_542_v39)), false, globalState), _546_v43);
              _547_v44 = ((_547_v44).Difference(_547_v44)).Difference((_547_v44).Union(_547_v44));
              let _index114 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_534_v33).length));
              (_534_v33)[_index114] = true;
              let _550_v45;
              let _nw87 = new _module.C1();
              _nw87.__ctor();
              _550_v45 = _nw87;
            }
          }
        }
      }
      let _551_v46;
      let _nw88 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
      _551_v46 = _nw88;
      let _552_v47;
      let _nw89 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _552_v47 = _nw89;
      let _index115 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_551_v46).length));
      (_551_v46)[_index115] = _dafny.MultiSet.fromElements(_552_v47);
      let _553_v48;
      _553_v48 = _dafny.MultiSet.fromElements(_552_v47, _552_v47, _552_v47);
      let _index116 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_551_v46).length));
      (_551_v46)[_index116] = ((_dafny.MultiSet.fromElements(_552_v47)).update(_552_v47, _module.__default.abs(_module.__default.fm3(p2, p2, globalState)))).Difference(_553_v48);
      let _554_v49;
      let _nw90 = Array((new BigNumber(6)).toNumber()).fill(_module.D0.Default());
      _554_v49 = _nw90;
      _554_v49 = _554_v49;
      let _555_v50;
      _555_v50 = false;
      let _556_v51;
      _556_v51 = _dafny.Seq.UnicodeFromString("rlqrprnjt");
      let _557_v52;
      _557_v52 = _dafny.Map.Empty.slice().updateUnsafe(_555_v50,_555_v50);
      let _558_v53;
      _558_v53 = _dafny.Set.fromElements((_dafny.ZERO).minus(p3));
      let _rhs121 = !(_557_v52).equals(_557_v52);
      let _rhs122 = ((_558_v53).Intersect(_dafny.Set.fromElements(p3))).contains((_this).fm8((_this).f6, p2, p2, globalState));
      let _rhs123 = _dafny.Seq.Concat(_556_v51, _dafny.Seq.UnicodeFromString("hs"));
      let _rhs124 = ((((_this).f6) ? (p3) : (p3))).isLessThanOrEqualTo((_dafny.ZERO).minus((((_this).f6) ? (new BigNumber(900)) : (p3))));
      let _rhs125 = (_this).f6;
      _555_v50 = _rhs121;
      _555_v50 = _rhs122;
      _556_v51 = _rhs123;
      _555_v50 = _rhs124;
      _555_v50 = _rhs125;
      let _559_v54;
      _559_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm20(_556_v51, _555_v50, globalState),(_this).f6);
      let _560_v55;
      _560_v55 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, p2);
      let _561_v56;
      _561_v56 = _dafny.Seq.of(((_560_v55).update(_module.__default.fm0(p2, globalState), _module.__default.abs(p3))).update(_555_v50, _module.__default.abs(p3)), _560_v55);
      _559_v54 = (_559_v54).update((_561_v56)[_module.__default.safeIndex(p3, new BigNumber((_561_v56).length))], true);
      _555_v50 = (new BigNumber((_556_v51).length)).isLessThan(p3);
      r0 = _module.D2.create_DC8((_module.__default.fm24(globalState)).dtor_cf0);
      r1 = (_560_v55).Union(_560_v55);
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let _562_v0;
      _562_v0 = _dafny.Seq.of(p1);
      let _hi1 = (_562_v0)[_module.__default.safeIndex(p1, new BigNumber((_562_v0).length))];
      for (let _563_i0 = (_dafny.ZERO).minus(p1); _563_i0.isLessThan(_hi1); _563_i0 = _563_i0.plus(_dafny.ONE)) {
        if ((_this).f6) {
          let _564_v1;
          _564_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vqwmkiqrg"),_563_i0);
          _564_v1 = (_564_v1).update(p0, _563_i0);
          let _565_v2;
          let _init11 = function (_566_i1) {
            return (_this).f6;
          };
          let _nw91 = Array((new BigNumber(17)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw91.length); _i0_11++) {
            _nw91[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _565_v2 = _nw91;
          let _567_v3;
          let _568_v4;
          let _out22;
          let _out23;
          let _outcollector10 = _module.__default.m0(_565_v2, (_this).f5, (_this).f6, (_this).f5, globalState);
          _out22 = _outcollector10[0];
          _out23 = _outcollector10[1];
          _567_v3 = _out22;
          _568_v4 = _out23;
          let _569_v5;
          _569_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_563_i0);
          let _570_v6;
          _570_v6 = _module.D0.create_DC0(!(_568_v4));
          _568_v4 = (_this).fm9((_563_i0).multipliedBy(new BigNumber((_569_v5).length)), _570_v6, (_this).f6, p1, globalState);
          let _571_v7;
          let _572_v8;
          let _out24;
          let _out25;
          let _outcollector11 = _module.__default.m0(_565_v2, new _dafny.CodePoint('y'.codePointAt(0)), ((_this).fm20(p0, !((_this).f6), globalState)).contains(_568_v4), (_this).f5, globalState);
          _out24 = _outcollector11[0];
          _out25 = _outcollector11[1];
          _571_v7 = _out24;
          _572_v8 = _out25;
          let _573_v9;
          _573_v9 = _dafny.Map.Empty.slice().updateUnsafe(_563_i0,_dafny.Seq.UnicodeFromString("ycs"));
          let _574_v10;
          _574_v10 = _dafny.Seq.of((_this).f6, _572_v8, (_this).f6);
          let _575_v11;
          let _nw92 = Array((new BigNumber(29)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _574_v10;
          _nw92[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_572_v8), _574_v10);
          _nw92[(new BigNumber(2)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_574_v10, _module.__default.safeIndex(_563_i0, new BigNumber((_574_v10).length)), (_this).f6);
          _nw92[(new BigNumber(4)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(5)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_572_v8);
          _nw92[(new BigNumber(7)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(8)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(9)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(10)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(11)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(12)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(13)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(14)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(15)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(16)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(17)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_574_v10, _574_v10);
          _nw92[(new BigNumber(19)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_574_v10, _574_v10);
          _nw92[(new BigNumber(21)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(22)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_574_v10, _574_v10);
          _nw92[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_574_v10, _574_v10);
          _nw92[(new BigNumber(25)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(26)).toNumber()] = _dafny.Seq.of((_this).f6);
          _nw92[(new BigNumber(27)).toNumber()] = _574_v10;
          _nw92[(new BigNumber(28)).toNumber()] = _dafny.Seq.Concat(_574_v10, _module.__default.fm2(_568_v4, _568_v4, globalState));
          _575_v11 = _nw92;
          let _576_v12;
          _576_v12 = _dafny.Seq.of(_575_v11);
          let _rhs126 = _573_v9;
          let _rhs127 = (_576_v12)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_563_i0,p1)).length), new BigNumber((_576_v12).length))];
          let _rhs128 = p1;
          let _lhs89 = globalState;
          _573_v9 = _rhs126;
          _575_v11 = _rhs127;
          _lhs89.f1 = _rhs128;
        } else {
          let _577_v13;
          let _nw93 = Array((new BigNumber(12)).toNumber()).fill(false);
          _577_v13 = _nw93;
          let _578_v14;
          let _579_v15;
          let _out26;
          let _out27;
          let _outcollector12 = _module.__default.m0(_577_v13, (_this).f5, (_this).f6, new _dafny.CodePoint('x'.codePointAt(0)), globalState);
          _out26 = _outcollector12[0];
          _out27 = _outcollector12[1];
          _578_v14 = _out26;
          _579_v15 = _out27;
          let _580_v16;
          _580_v16 = _dafny.Seq.of(_577_v13, _577_v13);
          let _581_v17;
          _581_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,_577_v13);
          let _582_v18;
          _582_v18 = _dafny.MultiSet.fromElements((_580_v16)[_module.__default.safeIndex(new BigNumber((_562_v0).length), new BigNumber((_580_v16).length))], _577_v13, (((_581_v17).contains(_563_i0)) ? ((_581_v17).get(_563_i0)) : (_577_v13)));
          let _583_v19;
          _583_v19 = _dafny.Set.fromElements(false);
          _582_v18 = (_582_v18).update(_577_v13, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_583_v19).length))));
          _577_v13 = _577_v13;
          let _584_v20;
          _584_v20 = _dafny.Seq.of((_563_i0).isLessThan(_563_i0), (_this).f6, _579_v15, !((_this).f6), (_this).f6);
          _584_v20 = _dafny.Seq.Concat(_584_v20, _584_v20);
          let _585_v21;
          _585_v21 = _dafny.Map.Empty.slice().updateUnsafe((false) === (true),new BigNumber((p0).length));
          _585_v21 = _585_v21;
        }
        let _586_v22;
        _586_v22 = new _dafny.CodePoint('a'.codePointAt(0));
        _586_v22 = _586_v22;
        let _587_v23;
        _587_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f6);
        let _588_v24;
        _588_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_562_v0, _dafny.Seq.of(new BigNumber((_587_v23).length), new BigNumber((p0).length)))).length),p0);
        _588_v24 = _588_v24;
        let _589_v25;
        let _nw94 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _589_v25 = _nw94;
        let _index117 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_589_v25).length));
        (_589_v25)[_index117] = new BigNumber((p0).length);
        let _index118 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_589_v25).length));
        (_589_v25)[_index118] = (p1).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-825)), ((_590_v22) => function (_591_i2) {
          return _590_v22;
        })(_586_v22))).length)));
      }
      let _592_v26;
      _592_v26 = false;
      _592_v26 = _592_v26;
      let _593_v27;
      _593_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6) || ((_this).f6),(((_this).f6) ? (p1) : (p1)));
      let _594_v28;
      _594_v28 = _dafny.Set.fromElements(p0);
      let _595_v29;
      let _init12 = ((_596_v26) => function (_597_i3) {
        return (_597_i3).plus(new BigNumber((_dafny.Seq.of(_596_v26, _596_v26, (_this).f6)).length));
      })(_592_v26);
      let _nw95 = Array((new BigNumber(14)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw95.length); _i0_12++) {
        _nw95[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _595_v29 = _nw95;
      let _598_v30;
      _598_v30 = _595_v29;
      let _599_v31;
      _599_v31 = _dafny.Map.Empty.slice().updateUnsafe(_598_v30,p1);
      _593_v27 = (_593_v27).update((_594_v28).IsDisjointFrom(_594_v28), new BigNumber((_599_v31).length));
      let _600_v32;
      _600_v32 = _dafny.MultiSet.fromElements(_595_v29, _595_v29, _595_v29);
      let _601_i4;
      _601_i4 = _dafny.ZERO;
      L7: {
        while (((_600_v32).Difference(_600_v32)).IsSubsetOf(_600_v32)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_601_i4)) {
              break L7;
            }
            _601_i4 = (_601_i4).plus(_dafny.ONE);
            let _602_v33;
            _602_v33 = _dafny.Set.fromElements((_this).f6);
            let _603_v34;
            let _nw96 = Array((new BigNumber(5)).toNumber());
            _nw96[(_dafny.ZERO).toNumber()] = (_602_v33).Intersect(_602_v33);
            _nw96[(_dafny.ONE).toNumber()] = (_602_v33).Intersect(_602_v33);
            _nw96[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(false, _592_v26);
            _nw96[(new BigNumber(3)).toNumber()] = (_602_v33).Difference(_602_v33);
            _nw96[(new BigNumber(4)).toNumber()] = (_this).fm21(globalState);
            _603_v34 = _nw96;
            let _index119 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_603_v34).length));
            (_603_v34)[_index119] = (_602_v33).Difference(_602_v33);
            let _index120 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_603_v34).length));
            let _rhs129 = (_602_v33).Union(_602_v33);
            let _rhs130 = p1;
            let _lhs90 = _603_v34;
            let _lhs91 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_603_v34).length));
            let _lhs92 = globalState;
            _lhs90[_lhs91] = _rhs129;
            _lhs92.f1 = _rhs130;
            let _604_v35;
            _604_v35 = _module.D0.create_DC0(_592_v26);
            let _605_v36;
            _605_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
            let _606_v37;
            _606_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25((_this).f6, p1, _604_v35, _605_v36, globalState),_592_v26);
            let _607_v38;
            _607_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_562_v0);
            let _608_v39;
            _608_v39 = _dafny.MultiSet.fromElements(_592_v26, (_this).f6);
            _606_v37 = (_606_v37).update((_607_v38).Merge(_607_v38), (new BigNumber((_608_v39).cardinality())).isLessThanOrEqualTo(p1));
            let _source11 = _595_v29;
            let _609___mcc_h0 = _source11;
            let _610_cf13 = _609___mcc_h0;
            _592_v26 = _592_v26;
            let _611_v40;
            _611_v40 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_592_v26);
            (globalState).f1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_611_v40).Merge(_611_v40)).length)));
            _593_v27 = (_593_v27).update((_this).f6, _module.__default.safeModuloInt(_module.__default.fm3(_592_v26, _module.__default.fm0(_592_v26, globalState), globalState), p1));
            let _612_v43;
            _612_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),p1);
            _592_v26 = (_this).fm10(_608_v39, function () {
              let _coll23 = new _dafny.Map();
              for (const _compr_23 of (function () {
                let _coll24 = new _dafny.Map();
                for (const _compr_24 of (_612_v43).Keys.Elements) {
                  let _613_v42 = _compr_24;
                  if ((_612_v43).contains(_613_v42)) {
                    _coll24.push([_613_v42,(_this).f6]);
                  }
                }
                return _coll24;
              }()).Keys.Elements) {
                let _614_v41 = _compr_23;
                if ((function () {
                  let _coll25 = new _dafny.Map();
                  for (const _compr_25 of (_612_v43).Keys.Elements) {
                    let _613_v42 = _compr_25;
                    if ((_612_v43).contains(_613_v42)) {
                      _coll25.push([_613_v42,(_this).f6]);
                    }
                  }
                  return _coll25;
                }()).contains(_614_v41)) {
                  _coll23.push([_614_v41,(_this).f6]);
                }
              }
              return _coll23;
            }(), p1, globalState);
            let _615_v45;
            _615_v45 = _dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll26 = new _dafny.Set();
              for (const _compr_26 of _dafny.IntegerRange(new BigNumber(572), new BigNumber(521))) {
                let _616_v44 = _compr_26;
                if (((new BigNumber(572)).isLessThanOrEqualTo(_616_v44)) && ((_616_v44).isLessThan(new BigNumber(521)))) {
                  _coll26.add((_616_v44).plus(new BigNumber(611)));
                }
              }
              return _coll26;
            }(),p1);
            let _617_v46;
            _617_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_608_v39).cardinality()),(_this).f6);
            let _618_v47;
            _618_v47 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(976),_617_v46)).length));
            (globalState).f1 = (_dafny.ZERO).minus((((_615_v45).contains(_618_v47)) ? ((_615_v45).get(_618_v47)) : ((p1).plus(p1))));
          }
        }
      }
      let _source12 = _598_v30;
      let _619___mcc_h1 = _source12;
      let _620_cf13 = _619___mcc_h1;
      let _621_v48;
      _621_v48 = _dafny.Seq.of(_562_v0);
      let _622_v49;
      let _nw97 = new _module.C0();
      _nw97.__ctor(((_this).f6) && (_592_v26), _621_v48);
      _622_v49 = _nw97;
      let _623_v50;
      _623_v50 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f6);
      _623_v50 = _dafny.Map.Empty.slice().updateUnsafe(p1,_592_v26);
      let _624_v51;
      _624_v51 = _dafny.Set.fromElements((_this).f6, (_622_v49).f11);
      let _625_v52;
      let _nw98 = Array((new BigNumber(6)).toNumber());
      _nw98[(_dafny.ZERO).toNumber()] = _592_v26;
      _nw98[(_dafny.ONE).toNumber()] = !(_592_v26) || (!((_622_v49).f11));
      _nw98[(new BigNumber(2)).toNumber()] = (_622_v49).f11;
      _nw98[(new BigNumber(3)).toNumber()] = true;
      _nw98[(new BigNumber(4)).toNumber()] = !((_dafny.Set.fromElements((_622_v49).f11, _592_v26, (_622_v49).f11, (_622_v49).f11, (_622_v49).f11)).IsProperSubsetOf(_624_v51));
      _nw98[(new BigNumber(5)).toNumber()] = ((_622_v49).f11) || (false);
      _625_v52 = _nw98;
      let _index121 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_625_v52).length));
      (_625_v52)[_index121] = (_622_v49).f11;
      let _626_v53;
      let _nw99 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _626_v53 = _nw99;
      let _627_v54;
      _627_v54 = _dafny.Seq.of(_592_v26, _592_v26, (_this).f6);
      let _index122 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_625_v52).length));
      let _rhs131 = p1;
      let _rhs132 = (((_this).f6) ? (_592_v26) : ((new BigNumber((_627_v54).length)).isLessThanOrEqualTo(p1)));
      let _rhs133 = (p1).isLessThan((_dafny.ZERO).minus(p1));
      let _rhs134 = _module.__default.fm3(_module.__default.fm0(_592_v26, globalState), ((true) ? (_592_v26) : (_592_v26)), globalState);
      let _rhs135 = _626_v53;
      let _lhs93 = globalState;
      let _lhs94 = _625_v52;
      let _lhs95 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_625_v52).length));
      let _lhs96 = globalState;
      _lhs93.f1 = _rhs131;
      _592_v26 = _rhs132;
      _lhs94[_lhs95] = _rhs133;
      _lhs96.f1 = _rhs134;
      _626_v53 = _rhs135;
      let _628_v55;
      _628_v55 = _dafny.Map.Empty.slice().updateUnsafe(_592_v26,_595_v29);
      let _629_v56;
      let _nw100 = Array((new BigNumber(20)).toNumber());
      _nw100[(_dafny.ZERO).toNumber()] = _595_v29;
      _nw100[(_dafny.ONE).toNumber()] = (((_628_v55).contains((_this).f6)) ? ((_628_v55).get((_this).f6)) : (_620_cf13));
      _nw100[(new BigNumber(2)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(3)).toNumber()] = _595_v29;
      _nw100[(new BigNumber(4)).toNumber()] = (((_625_v52)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_625_v52).length))]) ? (_595_v29) : (_595_v29));
      _nw100[(new BigNumber(5)).toNumber()] = (_598_v30);
      _nw100[(new BigNumber(6)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(7)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(8)).toNumber()] = _595_v29;
      _nw100[(new BigNumber(9)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(10)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(11)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(12)).toNumber()] = _595_v29;
      _nw100[(new BigNumber(13)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(14)).toNumber()] = _595_v29;
      _nw100[(new BigNumber(15)).toNumber()] = _595_v29;
      _nw100[(new BigNumber(16)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(17)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(18)).toNumber()] = _620_cf13;
      _nw100[(new BigNumber(19)).toNumber()] = _620_cf13;
      _629_v56 = _nw100;
      _629_v56 = _629_v56;
      if (_592_v26) {
        _592_v26 = (((_this).f6) ? (_592_v26) : (!(!(true)) || ((_this).f6)));
        (globalState).f1 = p1;
        let _630_v57;
        _630_v57 = _dafny.Set.fromElements(p1, p1);
        _592_v26 = (_592_v26) || ((_630_v57).IsSubsetOf(_630_v57));
        if ((_module.__default.fm26(globalState)).dtor_cf5) {
          let _631_v58;
          _631_v58 = _module.D2.create_DC7(p1, p1);
          let _632_v59;
          _632_v59 = _dafny.Map.Empty.slice().updateUnsafe(_631_v58,p0);
          _632_v59 = (_632_v59).update(_631_v58, p0);
          let _633_v60;
          let _nw101 = new _module.C0();
          _nw101.__ctor((_this).f6, _dafny.Seq.of(_562_v0));
          _633_v60 = _nw101;
          _633_v60 = _633_v60;
          let _634_v61;
          let _init13 = ((_635_v26) => function (_636_i5) {
            return _635_v26;
          })(_592_v26);
          let _nw102 = Array((new BigNumber(21)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw102.length); _i0_13++) {
            _nw102[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _634_v61 = _nw102;
          let _nw103 = Array((new BigNumber(24)).toNumber()).fill(false);
          _634_v61 = _nw103;
          let _637_v62;
          _637_v62 = _dafny.Seq.of(!((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).contains(p1)), (_this).f6, _592_v26);
          let _rhs136 = _dafny.Seq.Concat(_637_v62, _dafny.Seq.of((_633_v60).f11, (_this).f6, _592_v26));
          let _rhs137 = p1;
          let _rhs138 = p1;
          let _rhs139 = _dafny.Seq.update(_dafny.Seq.Concat(_562_v0, _dafny.Seq.Concat(_562_v0, _562_v0)), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.Concat(_562_v0, _dafny.Seq.Concat(_562_v0, _562_v0))).length)), p1);
          let _lhs97 = globalState;
          let _lhs98 = globalState;
          _637_v62 = _rhs136;
          _lhs97.f1 = _rhs137;
          _lhs98.f1 = _rhs138;
          _562_v0 = _rhs139;
          let _638_v63;
          _638_v63 = _dafny.MultiSet.fromElements((p1).plus(new BigNumber(-917)));
          _638_v63 = _638_v63;
        } else {
          (globalState).f1 = (p1).plus(p1);
          let _639_v64;
          let _nw104 = Array((new BigNumber(27)).toNumber()).fill(false);
          _639_v64 = _nw104;
          let _640_v65;
          _640_v65 = _dafny.Seq.of(_593_v27);
          let _641_v66;
          _641_v66 = _dafny.MultiSet.fromElements(_592_v26);
          let _642_v67;
          _642_v67 = _dafny.Set.fromElements(_592_v26, (_this).f6);
          let _643_v68;
          _643_v68 = _dafny.Map.Empty.slice().updateUnsafe(_642_v67,_592_v26);
          let _rhs140 = _639_v64;
          let _rhs141 = (_module.__default.safeModuloInt(new BigNumber((p0).length), p1)).plus(new BigNumber(349));
          let _rhs142 = (p1).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_644_p1) => function (_645_i6) {
            return _dafny.Map.Empty.slice().updateUnsafe(true,_644_p1);
          })(p1)), _640_v65)).length));
          let _rhs143 = (_this).fm10(_641_v66, _643_v68, (p1).plus(p1), globalState);
          let _lhs99 = globalState;
          let _lhs100 = globalState;
          _639_v64 = _rhs140;
          _lhs99.f1 = _rhs141;
          _lhs100.f1 = _rhs142;
          _592_v26 = _rhs143;
          (globalState).f1 = (_dafny.ZERO).minus((_this).fm7(!(_dafny.Seq.IsProperPrefixOf(_562_v0, _562_v0)), globalState));
          _592_v26 = (_this).f6;
          let _646_v69;
          let _nw105 = new _module.C1();
          _nw105.__ctor();
          _646_v69 = _nw105;
        }
        let _index123 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_595_v29).length));
        (_595_v29)[_index123] = p1;
        let _index124 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_595_v29).length));
        (_595_v29)[_index124] = (p1).plus(p1);
      } else {
        let _647_v70;
        let _nw106 = Array((new BigNumber(6)).toNumber()).fill(false);
        _647_v70 = _nw106;
        let _648_v71;
        _648_v71 = _module.D0.create_DC0(false);
        let _649_v72;
        _649_v72 = _module.D0.create_DC3(_648_v71);
        let _rhs144 = _647_v70;
        let _rhs145 = ((_562_v0)[_module.__default.safeIndex(p1, new BigNumber((_562_v0).length))]).isLessThanOrEqualTo(p1);
        let _rhs146 = _module.D0.create_DC3(_648_v71);
        _647_v70 = _rhs144;
        _592_v26 = _rhs145;
        _649_v72 = _rhs146;
        let _650_v73;
        let _init14 = ((_651_v26, _652_p0) => function (_653_i7) {
          return ((_651_v26) ? (_652_p0) : (_dafny.Seq.UnicodeFromString("cqdfgekt")));
        })(_592_v26, p0);
        let _nw107 = Array((new BigNumber(2)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw107.length); _i0_14++) {
          _nw107[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _650_v73 = _nw107;
        let _rhs147 = (_dafny.ZERO).minus(p1);
        let _rhs148 = ((p1).multipliedBy(p1)).minus((new BigNumber((p0).length)).plus(p1));
        let _rhs149 = ((p1).multipliedBy(p1)).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(-35), p1));
        let _rhs150 = ((!((_this).f6)) ? (_650_v73) : ((((_this).f6) ? (_650_v73) : (_650_v73))));
        let _lhs101 = globalState;
        let _lhs102 = globalState;
        let _lhs103 = globalState;
        _lhs101.f1 = _rhs147;
        _lhs102.f1 = _rhs148;
        _lhs103.f1 = _rhs149;
        _650_v73 = _rhs150;
        let _654_v74;
        _654_v74 = _dafny.MultiSet.fromElements((_this).f6);
        let _655_v75;
        _655_v75 = _dafny.Set.fromElements(true, (_this).f6, _592_v26, _592_v26);
        let _656_v76;
        _656_v76 = _dafny.Map.Empty.slice().updateUnsafe(_655_v75,true);
        let _657_v77;
        _657_v77 = _dafny.Seq.of((false) || ((_this).fm10(_654_v74, _656_v76, p1, globalState)), true, _592_v26);
        if ((_657_v77)[_module.__default.safeIndex(_module.__default.safeDivisionInt(p1, new BigNumber((_562_v0).length)), new BigNumber((_657_v77).length))]) {
          let _658_v78;
          _658_v78 = _dafny.Seq.UnicodeFromString("oomt");
          _658_v78 = p0;
          let _659_v80;
          _659_v80 = _dafny.Set.fromElements(p1);
          _592_v26 = (_659_v80).IsSubsetOf(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-802)), function (_660_i8) {
              return new BigNumber(29);
            })).Elements) {
              let _661_v79 = _compr_27;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-802)), function (_660_i8) {
                return new BigNumber(29);
              }), _661_v79)) {
                _coll27.add((_661_v79).minus(new BigNumber(-659)));
              }
            }
            return _coll27;
          }());
          let _662_v81;
          _662_v81 = new _dafny.CodePoint('v'.codePointAt(0));
          let _663_v82;
          _663_v82 = _module.D1.create_DC4(new BigNumber(683));
          let _rhs151 = _647_v70;
          let _rhs152 = _647_v70;
          let _rhs153 = new _dafny.CodePoint('k'.codePointAt(0));
          let _rhs154 = (_663_v82).dtor_cf4;
          let _lhs104 = globalState;
          _647_v70 = _rhs151;
          _647_v70 = _rhs152;
          _662_v81 = _rhs153;
          _lhs104.f1 = _rhs154;
          let _index125 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_647_v70).length));
          (_647_v70)[_index125] = (_592_v26) === (_592_v26);
          let _index126 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_647_v70).length));
          (_647_v70)[_index126] = true;
          _592_v26 = ((_this).fm21(globalState)).IsDisjointFrom(_dafny.Set.fromElements((_this).f6, (_this).f6));
        } else {
          let _664_v83;
          let _nw108 = new _module.C1();
          _nw108.__ctor();
          _664_v83 = _nw108;
          let _665_v84;
          _665_v84 = _dafny.Map.Empty.slice().updateUnsafe(_592_v26,_664_v83);
          _665_v84 = _665_v84;
          let _666_v85;
          let _nw109 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _666_v85 = _nw109;
          _666_v85 = _666_v85;
          (globalState).f1 = new BigNumber(961);
          _592_v26 = (_this).f6;
          let _667_v86;
          let _nw110 = new _module.C1();
          _nw110.__ctor();
          _667_v86 = _nw110;
        }
        (globalState).f1 = p1;
        (globalState).f1 = p1;
      }
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      r0 = !_dafny.areEqual(p2, p2);
      let _668_v0;
      let _nw111 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _668_v0 = _nw111;
      let _index127 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length));
      (_668_v0)[_index127] = p1;
      let _index128 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length));
      (_668_v0)[_index128] = _module.__default.safeDivisionInt(p1, p1);
      let _669_v1;
      _669_v1 = _dafny.Seq.of(p0, (_this).f6);
      let _670_v2;
      _670_v2 = _module.D1.create_DC4(new BigNumber((_669_v1).length));
      let _pat_let_tv8 = _668_v0;
      let _pat_let_tv9 = _668_v0;
      let _hi2 = p1;
      for (let _671_i0 = (function (_pat_let4_0) {
        return function (_680_dt__update__tmp_h0) {
          return function (_pat_let5_0) {
            return function (_681_dt__update_hcf4_h0) {
              return _module.D1.create_DC4(_681_dt__update_hcf4_h0);
            }(_pat_let5_0);
          }((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_pat_let_tv8).length))]);
        }(_pat_let4_0);
      }(_670_v2)).dtor_cf4; _671_i0.isLessThan(_hi2); _671_i0 = _671_i0.plus(_dafny.ONE)) {
        let _672_v3;
        let _nw112 = Array((new BigNumber(16)).toNumber()).fill(false);
        _672_v3 = _nw112;
        let _index129 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        (_672_v3)[_index129] = (_this).f6;
        let _index130 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        (_672_v3)[_index130] = (_this).f6;
        r0 = ((p1).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_671_i0)))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ocyq"), p2)).length));
        let _673_v4;
        _673_v4 = _dafny.Seq.of((_668_v0)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length))]);
        let _674_v5;
        _674_v5 = _module.D6.create_DC13(_673_v4);
        r0 = !_dafny.areEqual((_674_v5).dtor_cf14, _673_v4);
        let _675_v6;
        let _nw113 = Array((new BigNumber(4)).toNumber()).fill(_module.D0.Default());
        _675_v6 = _nw113;
        let _676_v7;
        _676_v7 = _dafny.Map.Empty.slice().updateUnsafe(_672_v3,(_this).f5);
        let _677_v8;
        _677_v8 = _module.D0.create_DC2(_676_v7, (_this).f5);
        let _index131 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_675_v6).length));
        (_675_v6)[_index131] = ((p0) ? (_677_v8) : (_677_v8));
        let _pat_let_tv6 = _676_v7;
        let _pat_let_tv7 = _676_v7;
        let _index132 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length));
        let _index133 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        let _index134 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        let _index135 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_675_v6).length));
        let _rhs155 = (p1).plus((_668_v0)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length))]);
        let _rhs156 = !((((_this).f6) ? (!(!(_671_i0).isEqualTo(new BigNumber((_669_v1).length)))) : ((((_this).f6) ? ((_672_v3)[_module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length))]) : ((_this).f6)))));
        let _rhs157 = (_671_i0).isEqualTo(p1);
        let _rhs158 = function (_pat_let2_0) {
          return function (_678_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_679_dt__update_hcf1_h0) {
                return _module.D0.create_DC2(_679_dt__update_hcf1_h0, (_678_dt__update__tmp_h1).dtor_cf2);
              }(_pat_let3_0);
            }((_pat_let_tv6).Merge(_pat_let_tv7));
          }(_pat_let2_0);
        }(_677_v8);
        let _rhs159 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
        let _lhs105 = _668_v0;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length));
        let _lhs107 = _672_v3;
        let _lhs108 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        let _lhs109 = _672_v3;
        let _lhs110 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_672_v3).length));
        let _lhs111 = _675_v6;
        let _lhs112 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_675_v6).length));
        let _lhs113 = globalState;
        _lhs105[_lhs106] = _rhs155;
        _lhs107[_lhs108] = _rhs156;
        _lhs109[_lhs110] = _rhs157;
        _lhs111[_lhs112] = _rhs158;
        _lhs113.f1 = _rhs159;
      }
      let _682_v9;
      _682_v9 = _dafny.MultiSet.fromElements(p0);
      let _683_v10;
      _683_v10 = _dafny.Set.fromElements(p0, (_this).f6);
      let _684_v11;
      _684_v11 = _dafny.Map.Empty.slice().updateUnsafe(_683_v10,p0);
      let _685_v12;
      _685_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p1);
      let _686_v13;
      _686_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber((_685_v12).length));
      let _rhs160 = (_this).fm10(_682_v9, (_684_v11).Merge(_684_v11), ((true) ? (new BigNumber((_686_v13).length)) : ((_668_v0)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length))])), globalState);
      let _rhs161 = _668_v0;
      let _rhs162 = (_this).f6;
      r0 = _rhs160;
      _668_v0 = _rhs161;
      r0 = _rhs162;
      let _687_v14;
      _687_v14 = _dafny.Seq.of((_668_v0)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length))]);
      let _688_v15;
      _688_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_687_v14).length),p0);
      let _689_v16;
      _689_v16 = _dafny.Seq.of(new BigNumber((_688_v15).length), p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-294)), function (_690_i1) {
        return (_this).f5;
      })).length), p1);
      r0 = _dafny.areEqual(_689_v16, _687_v14);
      let _hi3 = p1;
      for (let _691_i2 = p1; _691_i2.isLessThan(_hi3); _691_i2 = _691_i2.plus(_dafny.ONE)) {
        r0 = (_685_v12).contains((_this).f6);
        let _692_v17;
        let _init15 = function (_693_i3) {
          return !((_this).f6);
        };
        let _nw114 = Array((new BigNumber(2)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw114.length); _i0_15++) {
          _nw114[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _692_v17 = _nw114;
        let _694_v18;
        _694_v18 = _module.D0.create_DC0((_this).f6);
        let _index136 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_692_v17).length));
        (_692_v17)[_index136] = (_this).fm9(_691_i2, _694_v18, (_this).f6, new BigNumber(236), globalState);
        let _index137 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_692_v17).length));
        (_692_v17)[_index137] = (_669_v1)[_module.__default.safeIndex((_668_v0)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_668_v0).length))], new BigNumber((_669_v1).length))];
        let _695_v19;
        _695_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_668_v0);
        _668_v0 = (((_695_v19).contains((_this).f5)) ? ((_695_v19).get((_this).f5)) : (((!(p0)) ? (_668_v0) : (_668_v0))));
        let _696_v20;
        _696_v20 = _668_v0;
        let _index138 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_692_v17).length));
        let _rhs163 = (_682_v9).IsDisjointFrom(_682_v9);
        let _rhs164 = _696_v20;
        let _lhs114 = _692_v17;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_692_v17).length));
        _lhs114[_lhs115] = _rhs163;
        _696_v20 = _rhs164;
      }
      r0 = (_dafny.Set.fromElements(p0, p0, p0)).IsSubsetOf((_683_v10).Intersect(_683_v10));
      r1 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), function (_697_i4) {
        return (_this).f5;
      }), _dafny.Seq.UnicodeFromString("jr"));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _698_v0;
      _698_v0 = false;
      let _699_v1;
      let _nw115 = Array((new BigNumber(10)).toNumber()).fill(_module.D2.Default());
      _699_v1 = _nw115;
      let _700_v2;
      _700_v2 = _dafny.Set.fromElements(_699_v1);
      _698_v0 = (_700_v2).IsProperSubsetOf(_700_v2);
      let _701_v3;
      _701_v3 = new BigNumber(-976);
      let _702_v4;
      _702_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_698_v0, _698_v0, (_this).f6, globalState),_701_v3);
      let _703_v5;
      _703_v5 = _702_v4;
      let _704_v6;
      _704_v6 = _dafny.Seq.of((_this).f5);
      _703_v5 = _module.__default.fm27(_704_v6, _698_v0, _698_v0, globalState);
      r0 = _701_v3;
      let _705_v8;
      _705_v8 = _module.D2.create_DC7(_701_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(488)), function (_706_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length));
      let _707_v9;
      _707_v9 = _dafny.Set.fromElements(_705_v8);
      let _708_v10;
      _708_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_707_v9).Elements) {
          let _709_v7 = _compr_28;
          if ((_707_v9).contains(_709_v7)) {
            _coll28.push([_709_v7,new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("tfqu"))).cardinality())]);
          }
        }
        return _coll28;
      }());
      let _710_v11;
      _710_v11 = _dafny.Map.Empty.slice().updateUnsafe(_705_v8,_701_v3);
      _708_v10 = (_708_v10).update((_this).f5, _710_v11);
      let _711_v12;
      let _nw116 = new _module.C1();
      _nw116.__ctor();
      _711_v12 = _nw116;
      let _712_v13;
      _712_v13 = _dafny.Set.fromElements(_711_v12);
      let _713_i1;
      _713_i1 = _dafny.ZERO;
      L8: {
        while ((_712_v13).IsSubsetOf(_712_v13)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_713_i1)) {
              break L8;
            }
            _713_i1 = (_713_i1).plus(_dafny.ONE);
            let _714_v14;
            _714_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), function (_715_i2) {
              return (_this).f5;
            }),(_this).f6);
            _714_v14 = _714_v14;
            r0 = _701_v3;
            _698_v0 = (true) || (_698_v0);
            let _716_v15;
            _716_v15 = new _dafny.CodePoint('w'.codePointAt(0));
            _716_v15 = _716_v15;
          }
        }
      }
      let _717_i3;
      _717_i3 = _dafny.ZERO;
      L9: {
        while (_698_v0) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_717_i3)) {
              break L9;
            }
            _717_i3 = (_717_i3).plus(_dafny.ONE);
            if ((_this).f6) {
              let _718_v16;
              let _init16 = ((_719_v0) => function (_720_i4) {
                return _719_v0;
              })(_698_v0);
              let _nw117 = Array((new BigNumber(15)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw117.length); _i0_16++) {
                _nw117[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _718_v16 = _nw117;
              let _721_v17;
              _721_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_718_v16);
              (globalState).f1 = new BigNumber((_721_v17).length);
              let _722_v18;
              let _nw118 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
              _722_v18 = _nw118;
              _722_v18 = _722_v18;
              let _723_v19;
              _723_v19 = _dafny.MultiSet.fromElements(_718_v16, _718_v16);
              _698_v0 = (new BigNumber((_704_v6).length)).isLessThanOrEqualTo(new BigNumber((_723_v19).cardinality()));
              let _724_v20;
              _724_v20 = _dafny.Seq.of(_698_v0, (_this).f6);
              r0 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_701_v3), _701_v3)).minus(_module.__default.fm3((_724_v20)[_module.__default.safeIndex(new BigNumber((_724_v20).length), new BigNumber((_724_v20).length))], _module.__default.fm0(false, globalState), globalState));
              let _725_v21;
              _725_v21 = _dafny.MultiSet.fromElements(_711_v12);
              let _726_v22;
              _726_v22 = _dafny.Seq.of(_701_v3);
              let _727_v23;
              _727_v23 = _dafny.MultiSet.fromElements(_701_v3);
              let _index139 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_722_v18).length));
              (_722_v18)[_index139] = (_dafny.ZERO).minus((((_725_v21).contains(_711_v12)) ? ((_725_v21).get(_711_v12)) : ((_711_v12).fm4(_dafny.Seq.update(_724_v20, _module.__default.safeIndex(new BigNumber((_724_v20).length), new BigNumber((_724_v20).length)), (_711_v12).fm6(_726_v22, globalState)), _701_v3, _701_v3, _727_v23, globalState))));
              let _728_v24;
              _728_v24 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_729_i5) {
                return (_this).f5;
              }), _704_v6, _dafny.Seq.UnicodeFromString("ea"), _704_v6, _704_v6);
              let _730_v25;
              _730_v25 = _dafny.Seq.of(_dafny.Seq.Concat(_728_v24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(93)), ((_731_v6) => function (_732_i6) {
                return _731_v6;
              })(_704_v6))), _728_v24);
              let _index140 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_722_v18).length));
              (_722_v18)[_index140] = new BigNumber(((_730_v25)[_module.__default.safeIndex(_701_v3, new BigNumber((_730_v25).length))]).length);
            } else {
              let _733_v26;
              let _init17 = ((_734_v6) => function (_735_i7) {
                return (_735_i7).minus(new BigNumber((_dafny.Set.fromElements(_734_v6, _734_v6)).length));
              })(_704_v6);
              let _nw119 = Array((new BigNumber(9)).toNumber());
              for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw119.length); _i0_17++) {
                _nw119[_i0_17] = _init17(new BigNumber(_i0_17));
              }
              _733_v26 = _nw119;
              let _index141 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_733_v26).length));
              (_733_v26)[_index141] = new BigNumber((_704_v6).length);
              let _index142 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_733_v26).length));
              (_733_v26)[_index142] = ((((_dafny.ZERO).minus(_701_v3)).isLessThanOrEqualTo(_701_v3)) ? (_701_v3) : (_module.__default.safeDivisionInt(new BigNumber((_704_v6).length), new BigNumber((_702_v4).length))));
              _733_v26 = _733_v26;
              let _736_v27;
              _736_v27 = _dafny.Map.Empty.slice().updateUnsafe(_698_v0,(_this).f6);
              let _737_v28;
              _737_v28 = _dafny.Seq.of((_this).fm7(false, globalState));
              _736_v27 = (_736_v27).update((_this).f6, ((_this).fm6(_737_v28, globalState)) || (_698_v0));
              let _738_v29;
              _738_v29 = _dafny.Seq.of(_698_v0);
              let _index143 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_733_v26).length));
              (_733_v26)[_index143] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_738_v29, _738_v29), _dafny.Seq.update(_738_v29, _module.__default.safeIndex(_701_v3, new BigNumber((_738_v29).length)), (_this).f6))).length);
              let _739_v30;
              let _nw120 = Array((new BigNumber(20)).toNumber()).fill(false);
              _739_v30 = _nw120;
              let _740_v31;
              _740_v31 = new _dafny.CodePoint('c'.codePointAt(0));
              let _index144 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_733_v26).length));
              let _rhs165 = (_737_v28)[_module.__default.safeIndex(_701_v3, new BigNumber((_737_v28).length))];
              let _rhs166 = _704_v6;
              let _rhs167 = _739_v30;
              let _rhs168 = (_this).f5;
              let _rhs169 = (_this).f5;
              let _lhs116 = _733_v26;
              let _lhs117 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_733_v26).length));
              _lhs116[_lhs117] = _rhs165;
              _704_v6 = _rhs166;
              _739_v30 = _rhs167;
              _740_v31 = _rhs168;
              _740_v31 = _rhs169;
            }
            let _741_v32;
            let _nw121 = Array((new BigNumber(8)).toNumber()).fill(false);
            _741_v32 = _nw121;
            let _index145 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_741_v32).length));
            (_741_v32)[_index145] = _698_v0;
            let _index146 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_741_v32).length));
            (_741_v32)[_index146] = _698_v0;
            (globalState).f1 = (_dafny.ZERO).minus(_701_v3);
            let _742_v33;
            _742_v33 = _dafny.Set.fromElements(_698_v0);
            let _743_v34;
            let _nw122 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _743_v34 = _nw122;
            let _index147 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_743_v34).length));
            (_743_v34)[_index147] = new BigNumber((_dafny.Seq.Concat(_704_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), function (_744_i8) {
              return (_this).f5;
            }))).length);
            let _745_v35;
            _745_v35 = _dafny.Seq.of((_this).f6);
            let _index148 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_743_v34).length));
            let _rhs170 = (_742_v33).Difference(_dafny.Set.fromElements((_741_v32)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_741_v32).length))], _698_v0));
            let _rhs171 = _module.__default.safeModuloInt((_701_v3).plus(new BigNumber((_745_v35).length)), _701_v3);
            let _lhs118 = _743_v34;
            let _lhs119 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_743_v34).length));
            _742_v33 = _rhs170;
            _lhs118[_lhs119] = _rhs171;
          }
        }
      }
      r0 = (_dafny.ZERO).minus(_701_v3);
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _746_v0;
      _746_v0 = _dafny.Seq.UnicodeFromString("i");
      _746_v0 = _dafny.Seq.UnicodeFromString("e");
      let _747_v1;
      _747_v1 = false;
      let _748_v2;
      _748_v2 = _dafny.Set.fromElements(p2, _747_v1);
      _747_v1 = (_748_v2).IsProperSubsetOf(_748_v2);
      let _749_v3;
      _749_v3 = new BigNumber(153);
      let _750_v4;
      _750_v4 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_749_v3);
      let _751_v5;
      let _nw123 = Array((new BigNumber(6)).toNumber());
      _nw123[(_dafny.ZERO).toNumber()] = _749_v3;
      _nw123[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_749_v3);
      _nw123[(new BigNumber(2)).toNumber()] = _749_v3;
      _nw123[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((((_750_v4).contains((_this).f5)) ? ((_750_v4).get((_this).f5)) : (_749_v3)), _749_v3);
      _nw123[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_749_v3);
      _nw123[(new BigNumber(5)).toNumber()] = new BigNumber(632);
      _751_v5 = _nw123;
      let _index149 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length));
      (_751_v5)[_index149] = (_dafny.ZERO).minus(_749_v3);
      let _752_v6;
      _752_v6 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-352)), function (_753_i0) {
        return (_this).f5;
      }), _746_v0, _746_v0, _746_v0, _dafny.Seq.Concat(_746_v0, _746_v0));
      let _index150 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length));
      let _rhs172 = (_747_v1) && (p0);
      let _rhs173 = new BigNumber((_752_v6).length);
      let _lhs120 = _751_v5;
      let _lhs121 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length));
      _747_v1 = _rhs172;
      _lhs120[_lhs121] = _rhs173;
      let _754_v7;
      let _nw124 = Array((new BigNumber(13)).toNumber()).fill([]);
      _754_v7 = _nw124;
      let _755_v8;
      _755_v8 = _751_v5;
      let _index151 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_754_v7).length));
      (_754_v7)[_index151] = _755_v8;
      let _index152 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_754_v7).length));
      (_754_v7)[_index152] = _751_v5;
      _749_v3 = _749_v3;
      let _756_v9;
      _756_v9 = _dafny.Seq.of(_747_v1);
      let _757_i1;
      _757_i1 = _dafny.ZERO;
      L10: {
        while (((_756_v9)[_module.__default.safeIndex((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))], new BigNumber((_756_v9).length))]) === (p0)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_757_i1)) {
              break L10;
            }
            _757_i1 = (_757_i1).plus(_dafny.ONE);
            let _758_v10;
            _758_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,!(p2));
            let _759_v11;
            let _init18 = ((_760_v3) => function (_761_i2) {
              return (_dafny.MultiSet.fromElements(_760_v3, new BigNumber(4))).Union(_dafny.MultiSet.fromElements(_760_v3));
            })(_749_v3);
            let _nw125 = Array((new BigNumber(5)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw125.length); _i0_18++) {
              _nw125[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _759_v11 = _nw125;
            let _762_v12;
            _762_v12 = _dafny.Seq.of(_749_v3);
            let _763_v13;
            _763_v13 = _dafny.Map.Empty.slice().updateUnsafe(_749_v3,new BigNumber((_746_v0).length));
            let _764_v14;
            _764_v14 = _dafny.Set.fromElements(new BigNumber((_763_v13).length));
            let _765_v15;
            _765_v15 = _dafny.MultiSet.fromElements((_762_v12)[_module.__default.safeIndex(new BigNumber((_764_v14).length), new BigNumber((_762_v12).length))], (_dafny.ZERO).minus(new BigNumber((_762_v12).length)));
            let _index153 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_759_v11).length));
            (_759_v11)[_index153] = _765_v15;
            let _index154 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_759_v11).length));
            let _rhs174 = _758_v10;
            let _rhs175 = _746_v0;
            let _rhs176 = (_dafny.ZERO).minus(_749_v3);
            let _rhs177 = _module.__default.safeModuloInt((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))], _749_v3);
            let _rhs178 = (_765_v15).update(_749_v3, _module.__default.abs(_749_v3));
            let _lhs122 = _759_v11;
            let _lhs123 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_759_v11).length));
            _758_v10 = _rhs174;
            _746_v0 = _rhs175;
            _749_v3 = _rhs176;
            _749_v3 = _rhs177;
            _lhs122[_lhs123] = _rhs178;
            (globalState).f1 = _749_v3;
            let _766_v16;
            _766_v16 = _dafny.Map.Empty.slice().updateUnsafe(_748_v2,(_this).f6);
            let _767_v17;
            _767_v17 = _dafny.Map.Empty.slice().updateUnsafe(_749_v3,p0);
            let _768_v18;
            _768_v18 = _dafny.Map.Empty.slice().updateUnsafe(_762_v12,new BigNumber(316));
            let _769_v19;
            _769_v19 = _dafny.Seq.of(_768_v18, (_768_v18).update(_762_v12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_747_v1,(_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))])).length)));
            let _770_v20;
            _770_v20 = _dafny.Seq.of((_769_v19)[_module.__default.safeIndex(_749_v3, new BigNumber((_769_v19).length))]);
            let _771_v21;
            _771_v21 = _module.D0.create_DC0((_this).fm10(_dafny.MultiSet.FromArray(_756_v9), _766_v16, (_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))], globalState));
            let _772_v22;
            let _nw126 = Array((new BigNumber(16)).toNumber());
            _nw126[(_dafny.ZERO).toNumber()] = (_this).fm6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), ((_773_v5) => function (_774_i3) {
              return (_dafny.ZERO).minus((_773_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_773_v5).length))]);
            })(_751_v5)), globalState);
            _nw126[(_dafny.ONE).toNumber()] = p0;
            _nw126[(new BigNumber(2)).toNumber()] = (_this).f6;
            _nw126[(new BigNumber(3)).toNumber()] = true;
            _nw126[(new BigNumber(4)).toNumber()] = !((_756_v9)[_module.__default.safeIndex(new BigNumber((_746_v0).length), new BigNumber((_756_v9).length))]);
            _nw126[(new BigNumber(5)).toNumber()] = ((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))]).isEqualTo(new BigNumber(468));
            _nw126[(new BigNumber(6)).toNumber()] = (_this).fm10(_dafny.MultiSet.fromElements(p2, _747_v1), _766_v16, _749_v3, globalState);
            _nw126[(new BigNumber(7)).toNumber()] = (_this).f6;
            _nw126[(new BigNumber(8)).toNumber()] = p0;
            _nw126[(new BigNumber(9)).toNumber()] = (_this).f6;
            _nw126[(new BigNumber(10)).toNumber()] = p2;
            _nw126[(new BigNumber(11)).toNumber()] = !((((_767_v17).contains((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))])) ? ((_767_v17).get((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))])) : (_747_v1))) || (false);
            _nw126[(new BigNumber(12)).toNumber()] = !(!(!((new BigNumber((_module.__default.fm28(_747_v1, p0, (_769_v19)[_module.__default.safeIndex((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))], new BigNumber((_769_v19).length))], globalState)).length)).isLessThanOrEqualTo((_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))]))));
            _nw126[(new BigNumber(13)).toNumber()] = (_771_v21).dtor_cf0;
            _nw126[(new BigNumber(14)).toNumber()] = ((_this).f6) || (!((_this).f6));
            _nw126[(new BigNumber(15)).toNumber()] = !(new BigNumber(686)).isEqualTo(new BigNumber(53));
            _772_v22 = _nw126;
            _772_v22 = _772_v22;
            let _775_v23;
            _775_v23 = _module.D0.create_DC0(p2);
            let _776_v24;
            _776_v24 = _module.D0.create_DC3(_775_v23);
            let _777_v25;
            _777_v25 = _module.D2.create_DC7(_749_v3, (_751_v5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_751_v5).length))]);
            let _778_v26;
            _778_v26 = _dafny.Map.Empty.slice().updateUnsafe(_776_v24,(_777_v25).dtor_cf8);
            let _pat_let_tv10 = _775_v23;
            _778_v26 = (_778_v26).update(function (_pat_let6_0) {
              return function (_779_dt__update__tmp_h1) {
                return function (_pat_let7_0) {
                  return function (_780_dt__update_hcf3_h0) {
                    return _module.D0.create_DC3(_780_dt__update_hcf3_h0);
                  }(_pat_let7_0);
                }(_pat_let_tv10);
              }(_pat_let6_0);
            }(_776_v24), (_762_v12)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_749_v3)).cardinality()), new BigNumber((_762_v12).length))]);
          }
        }
      }
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f13 = false;
      this._f14 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f13, f14) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(42), new BigNumber(((((_this).f14) ? (_dafny.Seq.of(function () {
        let _coll29 = new _dafny.Set();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(403), new BigNumber(547))) {
          let _781_v0 = _compr_29;
          if (((new BigNumber(403)).isLessThanOrEqualTo(_781_v0)) && ((_781_v0).isLessThan(new BigNumber(547)))) {
            _coll29.add((_781_v0).plus(new BigNumber(-565)));
          }
        }
        return _coll29;
      }(), _dafny.Set.fromElements(new BigNumber(46), new BigNumber(880)), function () {
        let _coll30 = new _dafny.Set();
        for (const _compr_30 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length),(_this).f13)).length),(_this).f14)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ueuscelno")).length))).Elements) {
          let _782_v1 = _compr_30;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length),(_this).f13)).length),(_this).f14)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ueuscelno")).length))).contains(_782_v1)) {
            _coll30.add((_782_v1).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), function (_783_i0) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            })).length)));
          }
        }
        return _coll30;
      }())) : (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f13)).cardinality())))))).length)));
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4((_module.D2.create_DC7(new BigNumber(135), new BigNumber((_dafny.Seq.of((_this).f13, (_this).f13)).length))).dtor_cf7);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _784_v0;
      let _init19 = ((_785_p1) => function (_786_i0) {
        return (_786_i0).plus(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-324)), new BigNumber((_785_p1).cardinality()))).length));
      })(p1);
      let _nw127 = Array((new BigNumber(17)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw127.length); _i0_19++) {
        _nw127[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _784_v0 = _nw127;
      let _787_v1;
      _787_v1 = new BigNumber(-263);
      let _788_v2;
      _788_v2 = _dafny.Seq.of(p0);
      let _789_v3;
      _789_v3 = _dafny.Set.fromElements(new BigNumber(114), _787_v1, new BigNumber((_788_v2).length));
      let _index155 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length));
      (_784_v0)[_index155] = new BigNumber((_789_v3).length);
      let _index156 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length));
      (_784_v0)[_index156] = _787_v1;
      let _790_v4;
      _790_v4 = false;
      _790_v4 = (_788_v2)[_module.__default.safeIndex((_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))], new BigNumber((_788_v2).length))];
      let _791_v5;
      _791_v5 = _dafny.Seq.of((_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))], _787_v1);
      let _792_v6;
      _792_v6 = _dafny.Seq.UnicodeFromString("t");
      let _793_v7;
      _793_v7 = _dafny.Seq.of(_792_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_794_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }));
      let _795_v8;
      _795_v8 = _module.D2.create_DC6(_module.__default.fm1(new BigNumber(897), true, (_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))], globalState));
      let _796_v9;
      _796_v9 = _module.D2.create_DC9(_795_v8);
      let _797_v10;
      _797_v10 = _module.D1.create_DC4(new BigNumber((_793_v7).length));
      let _798_v11;
      let _nw128 = Array((new BigNumber(28)).toNumber()).fill(false);
      _798_v11 = _nw128;
      let _799_v12;
      _799_v12 = _798_v11;
      let _800_v13;
      _800_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_790_v4, globalState),_799_v12);
      let _801_v14;
      _801_v14 = _dafny.Seq.of(_800_v13);
      let _802_v15;
      _802_v15 = _dafny.Map.Empty.slice().updateUnsafe(_787_v1,(_801_v14)[_module.__default.safeIndex(new BigNumber((_792_v6).length), new BigNumber((_801_v14).length))]);
      let _803_v16;
      _803_v16 = _dafny.Map.Empty.slice().updateUnsafe(_802_v15,!(p2));
      let _rhs179 = _module.__default.fm37(!((_this).f13), _module.__default.safeDivisionInt(new BigNumber((_791_v5).length), new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)), _797_v10, _787_v1, globalState);
      let _rhs180 = _793_v7;
      let _rhs181 = _796_v9;
      let _rhs182 = (((_803_v16).contains(_802_v15)) ? ((_803_v16).get(_802_v15)) : ((_this).f14));
      let _rhs183 = _787_v1;
      let _lhs124 = globalState;
      _791_v5 = _rhs179;
      _793_v7 = _rhs180;
      _796_v9 = _rhs181;
      _790_v4 = _rhs182;
      _lhs124.f1 = _rhs183;
      let _804_v17;
      _804_v17 = _dafny.Map.Empty.slice().updateUnsafe(_787_v1,(_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))]);
      let _805_v18;
      _805_v18 = _dafny.Set.fromElements(_module.__default.fm0((_this).f14, globalState), (_this).f13, (_this).f14);
      (globalState).f1 = _module.__default.safeModuloInt(new BigNumber((_804_v17).length), new BigNumber((_805_v18).length));
      if ((_this).f13) {
        let _806_v19;
        _806_v19 = _module.D0.create_DC0(_790_v4);
        _790_v4 = (_806_v19).dtor_cf0;
        let _807_v20;
        _807_v20 = new _dafny.CodePoint('a'.codePointAt(0));
        let _808_v21;
        _808_v21 = _dafny.Map.Empty.slice().updateUnsafe(_798_v11,_807_v20);
        let _809_v22;
        _809_v22 = _module.D0.create_DC2(_808_v21, _807_v20);
        let _pat_let_tv11 = _798_v11;
        let _pat_let_tv12 = _807_v20;
        _807_v20 = (function (_pat_let8_0) {
          return function (_810_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_811_dt__update_hcf1_h0) {
                return _module.D0.create_DC2(_811_dt__update_hcf1_h0, (_810_dt__update__tmp_h0).dtor_cf2);
              }(_pat_let9_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv11,_pat_let_tv12));
          }(_pat_let8_0);
        }(_809_v22)).dtor_cf2;
        _790_v4 = (new BigNumber(58)).isLessThan(_787_v1);
        _790_v4 = (p2) === ((new BigNumber(243)).isEqualTo(_787_v1));
        let _812_v23;
        let _nw129 = Array((new BigNumber(23)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = _807_v20;
        _nw129[(_dafny.ONE).toNumber()] = _807_v20;
        _nw129[(new BigNumber(2)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(3)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(4)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(5)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw129[(new BigNumber(7)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(8)).toNumber()] = ((p0) ? (_807_v20) : (_807_v20));
        _nw129[(new BigNumber(9)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(10)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(11)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(12)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(13)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(14)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(15)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(16)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(17)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(18)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(19)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(20)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(21)).toNumber()] = _807_v20;
        _nw129[(new BigNumber(22)).toNumber()] = (_792_v6)[_module.__default.safeIndex(new BigNumber((_792_v6).length), new BigNumber((_792_v6).length))];
        _812_v23 = _nw129;
        let _index157 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_812_v23).length));
        (_812_v23)[_index157] = new _dafny.CodePoint('d'.codePointAt(0));
        let _index158 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_812_v23).length));
        (_812_v23)[_index158] = _807_v20;
      } else {
        let _index159 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length));
        (_784_v0)[_index159] = (_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))];
        let _index160 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length));
        let _rhs184 = !((_788_v2)[_module.__default.safeIndex(_787_v1, new BigNumber((_788_v2).length))]);
        let _rhs185 = (_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))];
        let _rhs186 = _788_v2;
        let _lhs125 = _784_v0;
        let _lhs126 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length));
        _790_v4 = _rhs184;
        _lhs125[_lhs126] = _rhs185;
        _788_v2 = _rhs186;
        let _813_v24;
        _813_v24 = _dafny.MultiSet.fromElements((_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))]);
        (globalState).f1 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(872), (_this).fm4(_788_v2, _787_v1, new BigNumber(176), _813_v24, globalState)), (_dafny.ZERO).minus((new BigNumber(781)).minus(_787_v1)));
        let _814_v25;
        _814_v25 = new _dafny.CodePoint('p'.codePointAt(0));
        let _815_v26;
        let _816_v27;
        let _out28;
        let _out29;
        let _outcollector13 = _module.__default.m0(_798_v11, _814_v25, p0, _814_v25, globalState);
        _out28 = _outcollector13[0];
        _out29 = _outcollector13[1];
        _815_v26 = _out28;
        _816_v27 = _out29;
        _814_v25 = new _dafny.CodePoint('x'.codePointAt(0));
      }
      (globalState).f1 = (_dafny.ZERO).minus((_784_v0)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_784_v0).length))]);
      return;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).IsSubsetOf(_dafny.MultiSet.fromElements(true))) {
        return (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("lsyykqbc")).length)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ol")).length)));
      } else {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(338)),false)).length);
      }
    };
    fm5(p0, globalState) {
      let _this = this;
      if (true) {
        return _module.D1.create_DC4(new BigNumber((function () {
  let _coll31 = new _dafny.Map();
  for (const _compr_31 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qfa")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length))))).Elements) {
    let _817_v0 = _compr_31;
    if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qfa")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length))))).contains(_817_v0)) {
      _coll31.push([_module.__default.safeDivisionInt(_817_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_dafny.Seq.of(new BigNumber(-454)))).length)),new BigNumber(-177)]);
    }
  }
  return _coll31;
}()).length));
      } else {
        return _module.D1.create_DC4(new BigNumber(368));
      }
    };
    fm30(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-327)), function (_818_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })) : (_dafny.Seq.UnicodeFromString("wx")))).length);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _819_v0;
      _819_v0 = true;
      _819_v0 = false;
      let _820_v1;
      let _nw130 = new _module.C0();
      _nw130.__ctor(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-580)), function (_821_i0) {
        return _dafny.Seq.of(new BigNumber(822), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(645), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-537)))).cardinality()))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), function (_822_i1) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length), new BigNumber(830), new BigNumber(979));
      }));
      _820_v1 = _nw130;
      _819_v0 = _819_v0;
      let _823_v2;
      _823_v2 = _dafny.Seq.of(_819_v0);
      let _824_i2;
      _824_i2 = _dafny.ZERO;
      L11: {
        while (!((_823_v2)[_module.__default.safeIndex(new BigNumber(-736), new BigNumber((_823_v2).length))])) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_824_i2)) {
              break L11;
            }
            _824_i2 = (_824_i2).plus(_dafny.ONE);
            let _825_v3;
            _825_v3 = _dafny.Seq.UnicodeFromString("iygwbkys");
            let _826_v4;
            _826_v4 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("xkbpmdi"));
            _825_v3 = _dafny.Seq.Concat((_826_v4)[_module.__default.safeIndex(new BigNumber(-365), new BigNumber((_826_v4).length))], _825_v3);
            let _827_v5;
            _827_v5 = new BigNumber(723);
            let _828_v6;
            _828_v6 = _module.D2.create_DC6(_825_v3);
            let _829_v7;
            _829_v7 = _dafny.Set.fromElements((_820_v1).f11);
            let _830_v8;
            _830_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_829_v7);
            _819_v0 = (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm31(p0, _827_v5, globalState), _module.__default.safeIndex(new BigNumber(((_828_v6).dtor_cf6).length), new BigNumber((_module.__default.fm31(p0, _827_v5, globalState)).length)), new BigNumber(888)), _module.__default.safeIndex(new BigNumber((_830_v8).length), new BigNumber((_dafny.Seq.update(_module.__default.fm31(p0, _827_v5, globalState), _module.__default.safeIndex(new BigNumber(((_828_v6).dtor_cf6).length), new BigNumber((_module.__default.fm31(p0, _827_v5, globalState)).length)), new BigNumber(888))).length)), new BigNumber((_825_v3).length)))).IsDisjointFrom(_module.__default.fm32(globalState));
            _827_v5 = _827_v5;
            (globalState).f1 = _827_v5;
          }
        }
      }
      let _831_v9;
      _831_v9 = new BigNumber(140);
      let _832_v10;
      _832_v10 = _dafny.Map.Empty.slice().updateUnsafe(_831_v9,p2);
      let _833_v11;
      _833_v11 = _dafny.Map.Empty.slice().updateUnsafe((_831_v9).multipliedBy(_831_v9),(_module.__default.fm33(globalState)).Merge(_832_v10));
      _832_v10 = (((_833_v11).contains(_831_v9)) ? ((_833_v11).get(_831_v9)) : (_832_v10));
      let _834_i3;
      _834_i3 = _dafny.ZERO;
      L12: {
        while (((!(_819_v0) || (false)) ? (_module.__default.fm0(_819_v0, globalState)) : (_819_v0))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_834_i3)) {
              break L12;
            }
            _834_i3 = (_834_i3).plus(_dafny.ONE);
            let _835_v12;
            let _nw131 = new _module.C1();
            _nw131.__ctor();
            _835_v12 = _nw131;
            let _836_v13;
            let _nw132 = Array((new BigNumber(18)).toNumber()).fill(false);
            _836_v13 = _nw132;
            let _837_v14;
            _837_v14 = _836_v13;
            let _838_v15;
            _838_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_836_v13);
            let _839_v16;
            let _nw133 = Array((new BigNumber(13)).toNumber());
            _nw133[(_dafny.ZERO).toNumber()] = _836_v13;
            _nw133[(_dafny.ONE).toNumber()] = _836_v13;
            _nw133[(new BigNumber(2)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(3)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(4)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(5)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(6)).toNumber()] = (_837_v14);
            _nw133[(new BigNumber(7)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(8)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(9)).toNumber()] = (((_838_v15).contains((_820_v1).f11)) ? ((_838_v15).get((_820_v1).f11)) : (_836_v13));
            _nw133[(new BigNumber(10)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(11)).toNumber()] = _836_v13;
            _nw133[(new BigNumber(12)).toNumber()] = _836_v13;
            _839_v16 = _nw133;
            let _index161 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_839_v16).length));
            (_839_v16)[_index161] = _836_v13;
            let _840_v17;
            _840_v17 = _dafny.Seq.of(_836_v13);
            let _index162 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_839_v16).length));
            (_839_v16)[_index162] = (_840_v17)[_module.__default.safeIndex(_831_v9, new BigNumber((_840_v17).length))];
            let _841_v18;
            _841_v18 = _module.D2.create_DC8(true);
            _819_v0 = ((_819_v0) ? ((_841_v18).dtor_cf9) : ((_820_v1).f11));
            if ((_831_v9).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p0, false, _819_v0)).length),_831_v9)).length))) {
              _819_v0 = (_823_v2)[_module.__default.safeIndex((_831_v9).plus(_831_v9), new BigNumber((_823_v2).length))];
              let _842_v19;
              _842_v19 = _module.D2.create_DC6(_dafny.Seq.UnicodeFromString("s"));
              _842_v19 = _842_v19;
              let _843_v20;
              let _nw134 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
              _843_v20 = _nw134;
              let _844_v21;
              _844_v21 = _dafny.Map.Empty.slice().updateUnsafe(_836_v13,((_820_v1).fm5(_841_v18, globalState)).dtor_cf4);
              let _index163 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_843_v20).length));
              (_843_v20)[_index163] = _844_v21;
              let _845_v22;
              _845_v22 = _dafny.Seq.of(_844_v21, _844_v21, _844_v21, _844_v21, _844_v21);
              let _846_v23;
              _846_v23 = _dafny.Map.Empty.slice().updateUnsafe((_820_v1).f11,_module.__default.fm34((_820_v1).f11, globalState));
              let _index164 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_843_v20).length));
              let _rhs187 = (_845_v22)[_module.__default.safeIndex(_831_v9, new BigNumber((_845_v22).length))];
              let _rhs188 = (((((_846_v23).contains(p0)) ? ((_846_v23).get(p0)) : (p1))).equals(_dafny.MultiSet.fromElements(p0))) && ((_820_v1).f11);
              let _rhs189 = (_820_v1).f11;
              let _lhs127 = _843_v20;
              let _lhs128 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_843_v20).length));
              _lhs127[_lhs128] = _rhs187;
              _819_v0 = _rhs188;
              _819_v0 = _rhs189;
              _819_v0 = true;
              let _847_v24;
              _847_v24 = _dafny.MultiSet.fromElements(_831_v9);
              let _848_v25;
              _848_v25 = new _dafny.CodePoint('d'.codePointAt(0));
              let _849_v26;
              _849_v26 = _dafny.Map.Empty.slice().updateUnsafe(_848_v25,_831_v9);
              let _850_v27;
              _850_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_847_v24).cardinality())),new BigNumber((_849_v26).length));
              let _851_v28;
              _851_v28 = _dafny.MultiSet.fromElements(new BigNumber((_850_v27).length), _831_v9, _831_v9, (_dafny.ZERO).minus(_831_v9), new BigNumber(741));
              let _852_v29;
              _852_v29 = _dafny.Map.Empty.slice().updateUnsafe(_820_v1,_851_v28);
              let _853_v30;
              _853_v30 = _dafny.Map.Empty.slice().updateUnsafe(_831_v9,(_this).fm4(_823_v2, new BigNumber(574), _831_v9, (((_852_v29).contains(_820_v1)) ? ((_852_v29).get(_820_v1)) : (_851_v28)), globalState));
              let _854_v31;
              _854_v31 = _dafny.Set.fromElements(p0);
              _853_v30 = (_853_v30).update(_831_v9, new BigNumber((_module.__default.fm31(_819_v0, new BigNumber((_854_v31).length), globalState)).length));
            } else {
              let _855_v32;
              let _nw135 = new _module.C1();
              _nw135.__ctor();
              _855_v32 = _nw135;
              _835_v12 = _835_v12;
              (globalState).f1 = (_dafny.ZERO).minus(_831_v9);
              let _856_v34;
              _856_v34 = _dafny.Map.Empty.slice().updateUnsafe((_820_v1).f11,(new BigNumber((function () {
                let _coll32 = new _dafny.Map();
                for (const _compr_32 of _dafny.IntegerRange(new BigNumber(599), new BigNumber(961))) {
                  let _857_v33 = _compr_32;
                  if (((new BigNumber(599)).isLessThanOrEqualTo(_857_v33)) && ((_857_v33).isLessThan(new BigNumber(961)))) {
                    _coll32.push([(_857_v33).plus(_831_v9),new _dafny.CodePoint('q'.codePointAt(0))]);
                  }
                }
                return _coll32;
              }()).length)).plus(_831_v9));
              _856_v34 = (_856_v34).update(!(_819_v0) || ((_820_v1).f11), _831_v9);
              _819_v0 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(_831_v9), (_dafny.ZERO).minus(_831_v9))).isEqualTo(_831_v9);
            }
          }
        }
      }
      return;
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = undefined;
      let r1 = [];
      let r2 = [];
      if (false) {
        let _858_v0;
        _858_v0 = false;
        let _859_v1;
        _859_v1 = new BigNumber(479);
        let _860_v2;
        _860_v2 = _dafny.Map.Empty.slice().updateUnsafe(_858_v0,_859_v1);
        let _861_v3;
        _861_v3 = _dafny.Seq.of(_dafny.Seq.of(_859_v1), _module.__default.fm31(_858_v0, _859_v1, globalState), _dafny.Seq.of(new BigNumber((_860_v2).length), _859_v1, _859_v1));
        let _862_v4;
        let _nw136 = new _module.C0();
        _nw136.__ctor(!(_858_v0) || (_858_v0), _861_v3);
        _862_v4 = _nw136;
        let _863_v5;
        _863_v5 = new _dafny.CodePoint('d'.codePointAt(0));
        _863_v5 = _863_v5;
        let _864_v6;
        _864_v6 = _dafny.Seq.of((_862_v4).f11, _858_v0, false);
        let _865_v7;
        _865_v7 = _module.D1.create_DC4(_859_v1);
        let _866_v8;
        _866_v8 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_859_v1));
        (globalState).f1 = (_862_v4).fm4(_864_v6, _859_v1, (_865_v7).dtor_cf4, _866_v8, globalState);
        (globalState).f1 = (((_858_v0) ? (_859_v1) : (_859_v1))).plus(new BigNumber(460));
        let _index165 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((p0).length));
        (p0)[_index165] = (_862_v4).f11;
        let _867_v9;
        _867_v9 = _dafny.MultiSet.fromElements(_858_v0);
        let _868_v10;
        _868_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_867_v9).cardinality()),_858_v0);
        let _869_v11;
        _869_v11 = _dafny.Map.Empty.slice().updateUnsafe(_862_v4,(((_868_v10).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_859_v1,(_862_v4).f11)).length)))) ? ((_868_v10).get((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_859_v1,(_862_v4).f11)).length)))) : (_858_v0)));
        let _index166 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((p0).length));
        (p0)[_index166] = (_869_v11).contains(_862_v4);
        let _index167 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((p0).length));
        let _index168 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((p0).length));
        let _rhs190 = _module.__default.fm0((_862_v4).f11, globalState);
        let _rhs191 = _858_v0;
        let _rhs192 = ((_858_v0) ? (((_858_v0) ? (_858_v0) : ((_862_v4).f11))) : (_858_v0));
        let _lhs129 = p0;
        let _lhs130 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((p0).length));
        let _lhs131 = p0;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((p0).length));
        _lhs129[_lhs130] = _rhs190;
        _lhs131[_lhs132] = _rhs191;
        _858_v0 = _rhs192;
      } else {
        let _870_v12;
        _870_v12 = true;
        _870_v12 = _module.__default.fm0(_870_v12, globalState);
        let _871_v13;
        _871_v13 = _dafny.Seq.of(p0, ((_870_v12) ? (p0) : (p0)));
        let _872_v14;
        _872_v14 = new BigNumber(497);
        _871_v13 = _dafny.Seq.update(_871_v13, _module.__default.safeIndex(_872_v14, new BigNumber((_871_v13).length)), p0);
        let _873_v15;
        _873_v15 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_870_v12)).length))));
        _873_v15 = _dafny.Seq.Concat(_873_v15, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-694)), function (_874_i0) {
          return new BigNumber(389);
        }));
        let _875_v16;
        _875_v16 = _dafny.Seq.of(_873_v15, _873_v15);
        let _876_v17;
        _876_v17 = _dafny.Seq.of(_dafny.Seq.of(_873_v15), _875_v16);
        let _877_v18;
        let _nw137 = new _module.C0();
        _nw137.__ctor(((false) ? (_module.__default.fm0(_870_v12, globalState)) : (!(_870_v12))), _dafny.Seq.Concat(_875_v16, (_876_v17)[_module.__default.safeIndex(_872_v14, new BigNumber((_876_v17).length))]));
        _877_v18 = _nw137;
        (globalState).f1 = _872_v14;
      }
      let _878_v19;
      _878_v19 = new BigNumber(470);
      let _879_v20;
      _879_v20 = _dafny.Seq.of(_878_v19);
      let _hi4 = (_879_v20)[_module.__default.safeIndex(_878_v19, new BigNumber((_879_v20).length))];
      for (let _880_i1 = _module.__default.safeModuloInt(_878_v19, _878_v19); _880_i1.isLessThan(_hi4); _880_i1 = _880_i1.plus(_dafny.ONE)) {
        let _881_v21;
        _881_v21 = false;
        let _index169 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
        (p0)[_index169] = _881_v21;
        let _index170 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
        (p0)[_index170] = _881_v21;
        if ((_880_i1).isLessThanOrEqualTo(_880_i1)) {
          let _882_v22;
          _882_v22 = _dafny.Seq.of(_879_v20);
          let _883_v23;
          let _nw138 = new _module.C0();
          _nw138.__ctor(_module.__default.fm0((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], globalState), (((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))]) ? (_882_v22) : (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_884_i1) => function (_885_i2) {
            return (_dafny.ZERO).minus(_884_i1);
          })(_880_i1)), _879_v20))));
          _883_v23 = _nw138;
          let _index171 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          (p0)[_index171] = (_883_v23).f11;
          let _886_v24;
          _886_v24 = _dafny.Seq.of(!(_881_v21), _module.__default.fm0(!((_883_v23).f11), globalState));
          let _887_v25;
          _887_v25 = _dafny.Set.fromElements(new BigNumber((_886_v24).length), _878_v19, _878_v19);
          let _888_v27;
          _888_v27 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(_880_i1, _878_v19, _880_i1, _880_i1, _878_v19));
          let _889_v28;
          _889_v28 = _dafny.MultiSet.fromElements(_880_i1, new BigNumber((_dafny.Seq.UnicodeFromString("tkfxh")).length));
          let _890_v29;
          _890_v29 = _dafny.Map.Empty.slice().updateUnsafe((_883_v23).f11,(_883_v23).f11);
          let _891_v30;
          _891_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-696),_878_v19);
          let _892_v31;
          let _nw139 = Array((new BigNumber(11)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = _887_v25;
          _nw139[(_dafny.ONE).toNumber()] = _887_v25;
          _nw139[(new BigNumber(2)).toNumber()] = (_887_v25).Union(_887_v25);
          _nw139[(new BigNumber(3)).toNumber()] = (((_883_v23).f11) ? (_887_v25) : (_887_v25));
          _nw139[(new BigNumber(4)).toNumber()] = (_887_v25).Union(_887_v25);
          _nw139[(new BigNumber(5)).toNumber()] = _887_v25;
          _nw139[(new BigNumber(6)).toNumber()] = function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(347), new BigNumber(914))) {
              let _893_v26 = _compr_33;
              if (((new BigNumber(347)).isLessThanOrEqualTo(_893_v26)) && ((_893_v26).isLessThan(new BigNumber(914)))) {
                _coll33.add(_module.__default.safeModuloInt(_893_v26, new BigNumber(690)));
              }
            }
            return _coll33;
          }();
          _nw139[(new BigNumber(7)).toNumber()] = _887_v25;
          _nw139[(new BigNumber(8)).toNumber()] = _887_v25;
          _nw139[(new BigNumber(9)).toNumber()] = (((_888_v27).contains((_883_v23).f11)) ? ((_888_v27).get((_883_v23).f11)) : (_887_v25));
          _nw139[(new BigNumber(10)).toNumber()] = _module.__default.fm35(_module.__default.fm1(_880_i1, (p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], _880_i1, globalState), _889_v28, _module.__default.fm36(_880_i1, new BigNumber((_890_v29).length), _880_i1, _891_v30, globalState), globalState);
          _892_v31 = _nw139;
          _892_v31 = _892_v31;
          _879_v20 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_894_i1) => function (_895_i3) {
            return _894_i1;
          })(_880_i1)), _879_v20);
          let _896_v32;
          _896_v32 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))],_dafny.Seq.Concat(_886_v24, _886_v24));
          let _897_v33;
          _897_v33 = _dafny.Seq.UnicodeFromString("iwoqi");
          let _898_v34;
          _898_v34 = new _dafny.CodePoint('a'.codePointAt(0));
          let _index172 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          let _index173 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          let _index174 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          let _rhs193 = !((_883_v23).f11);
          let _rhs194 = _module.__default.fm0((_883_v23).f11, globalState);
          let _rhs195 = _878_v19;
          let _rhs196 = (((_896_v32).contains((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))])) ? ((_896_v32).get((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))])) : (_886_v24));
          let _rhs197 = _dafny.areEqual(_dafny.Seq.Concat(_897_v33, _897_v33), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(211)), function (_899_i4) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }), _module.__default.safeIndex(_880_i1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(211)), function (_900_i4) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })).length)), _898_v34));
          let _lhs133 = p0;
          let _lhs134 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          let _lhs135 = p0;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          let _lhs137 = globalState;
          let _lhs138 = p0;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          _lhs133[_lhs134] = _rhs193;
          _lhs135[_lhs136] = _rhs194;
          _lhs137.f1 = _rhs195;
          _886_v24 = _rhs196;
          _lhs138[_lhs139] = _rhs197;
        } else {
          let _901_v35;
          let _nw140 = Array((new BigNumber(8)).toNumber()).fill([]);
          _901_v35 = _nw140;
          _901_v35 = _901_v35;
          let _902_v36;
          let _nw141 = new _module.C3();
          _nw141.__ctor(true, _881_v21);
          _902_v36 = _nw141;
          let _903_v37;
          _903_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_881_v21, (p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], globalState),_878_v19);
          let _904_v38;
          _904_v38 = _dafny.Map.Empty.slice().updateUnsafe(_903_v37,_881_v21);
          let _905_v39;
          _905_v39 = _dafny.Seq.of(_881_v21, (p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))]);
          let _906_v40;
          _906_v40 = _dafny.MultiSet.fromElements(_905_v39, _905_v39, _905_v39);
          let _907_v41;
          _907_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm30(_904_v38, _906_v40, globalState),_902_v36);
          _902_v36 = (((_907_v41).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), function (_909_i5) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)))) ? ((_907_v41).get((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), function (_908_i5) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)))) : (_902_v36));
          let _910_v42;
          let _nw142 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _910_v42 = _nw142;
          let _index175 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_910_v42).length));
          (_910_v42)[_index175] = _878_v19;
          let _index176 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_910_v42).length));
          (_910_v42)[_index176] = _878_v19;
          let _911_v43;
          let _nw143 = new _module.C3();
          _nw143.__ctor(false, true);
          _911_v43 = _nw143;
          let _912_v44;
          let _nw144 = new _module.C2();
          _nw144.__ctor(new _dafny.CodePoint('m'.codePointAt(0)), _881_v21);
          _912_v44 = _nw144;
        }
        if (!(_module.__default.fm0(true, globalState))) {
          let _913_v45;
          let _nw145 = new _module.C1();
          _nw145.__ctor();
          _913_v45 = _nw145;
          let _init20 = ((_914_v21, _915_p0, _916_i1, _917_v19) => function (_918_i6) {
            return !(!(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_914_v21, (_915_p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_915_p0).length))])).length), _916_i1, new BigNumber(266), _917_v19)).equals(_dafny.MultiSet.fromElements(_916_i1)));
          })(_881_v21, p0, _880_i1, _878_v19);
          let _nw146 = Array((new BigNumber(26)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw146.length); _i0_20++) {
            _nw146[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          r2 = _nw146;
          let _919_v46;
          _919_v46 = _dafny.Seq.UnicodeFromString("xr");
          let _index177 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length));
          (p0)[_index177] = _dafny.Seq.IsProperPrefixOf(_919_v46, _919_v46);
          let _920_v47;
          _920_v47 = _module.D0.create_DC1();
          _920_v47 = _920_v47;
          let _921_v48;
          let _init21 = ((_922_p0) => function (_923_i7) {
            return (_922_p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_922_p0).length))];
          })(p0);
          let _nw147 = Array((new BigNumber(25)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw147.length); _i0_21++) {
            _nw147[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _921_v48 = _nw147;
          let _924_v49;
          _924_v49 = _dafny.Seq.of((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], false);
          let _925_v50;
          _925_v50 = _dafny.Seq.of(true, (p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], !((_924_v49)[_module.__default.safeIndex((_dafny.ZERO).minus(_878_v19), new BigNumber((_924_v49).length))]), !(false));
          let _926_v51;
          _926_v51 = _dafny.Map.Empty.slice().updateUnsafe(_921_v48,(_924_v49)[_module.__default.safeIndex(_880_i1, new BigNumber((_924_v49).length))]);
          _926_v51 = (_926_v51).update(_921_v48, _881_v21);
        } else {
          let _927_v52;
          _927_v52 = new _dafny.CodePoint('x'.codePointAt(0));
          let _928_v53;
          _928_v53 = _dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))]);
          let _929_v54;
          let _nw148 = new _module.C2();
          _nw148.__ctor(_927_v52, (_928_v53).IsSubsetOf(_928_v53));
          _929_v54 = _nw148;
          (globalState).f1 = _880_i1;
          let _930_v55;
          let _nw149 = new _module.C2();
          _nw149.__ctor(_927_v52, (_878_v19).isLessThan(_878_v19));
          _930_v55 = _nw149;
          let _931_v56;
          _931_v56 = _module.D2.create_DC7(_module.__default.safeDivisionInt(_880_i1, new BigNumber((_dafny.Seq.update(_879_v20, _module.__default.safeIndex(new BigNumber(646), new BigNumber((_879_v20).length)), (_dafny.ZERO).minus(_880_i1))).length)), _880_i1);
          let _932_v57;
          _932_v57 = _dafny.Seq.UnicodeFromString("tecdwogxu");
          let _rhs198 = _878_v19;
          let _rhs199 = _931_v56;
          let _rhs200 = _932_v57;
          let _rhs201 = _881_v21;
          let _rhs202 = ((_881_v21) ? (_927_v52) : (_927_v52));
          let _lhs140 = globalState;
          _lhs140.f1 = _rhs198;
          _931_v56 = _rhs199;
          _932_v57 = _rhs200;
          _881_v21 = _rhs201;
          _927_v52 = _rhs202;
          _927_v52 = _927_v52;
        }
        let _933_v58;
        let _init22 = ((_934_v21) => function (_935_i8) {
          return _934_v21;
        })(_881_v21);
        let _nw150 = Array((new BigNumber(6)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw150.length); _i0_22++) {
          _nw150[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _933_v58 = _nw150;
        let _936_v59;
        _936_v59 = _dafny.Map.Empty.slice().updateUnsafe(_933_v58,(p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))]);
        let _937_v60;
        let _nw151 = new _module.C3();
        _nw151.__ctor((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))], (((_936_v59).contains(_933_v58)) ? ((_936_v59).get(_933_v58)) : ((p0)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((p0).length))])));
        _937_v60 = _nw151;
      }
      let _938_v61;
      _938_v61 = false;
      let _939_v62;
      let _nw152 = new _module.C3();
      _nw152.__ctor(_938_v61, _938_v61);
      _939_v62 = _nw152;
      let _hi5 = _878_v19;
      for (let _940_i9 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_939_v62,new _dafny.CodePoint('l'.codePointAt(0)))).length); _940_i9.isLessThan(_hi5); _940_i9 = _940_i9.plus(_dafny.ONE)) {
        let _941_v63;
        _941_v63 = _dafny.Seq.UnicodeFromString("amn");
        let _942_v64;
        _942_v64 = _dafny.Map.Empty.slice().updateUnsafe(_940_i9,new BigNumber((_941_v63).length));
        let _943_v65;
        _943_v65 = _dafny.Map.Empty.slice().updateUnsafe(_941_v63,_940_i9);
        (globalState).f1 = (new BigNumber((_942_v64).length)).plus(_module.__default.safeModuloInt(_940_i9, new BigNumber((_943_v65).length)));
        _878_v19 = _878_v19;
        let _944_v66;
        let _init23 = function (_945_i10) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), function (_946_i11) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          });
        };
        let _nw153 = Array((new BigNumber(24)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw153.length); _i0_23++) {
          _nw153[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _944_v66 = _nw153;
        let _index178 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_944_v66).length));
        (_944_v66)[_index178] = _941_v63;
        let _947_v67;
        _947_v67 = _dafny.Seq.of(!(_938_v61));
        let _948_v68;
        _948_v68 = _dafny.MultiSet.fromElements(_878_v19);
        let _index179 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_944_v66).length));
        (_944_v66)[_index179] = _module.__default.fm1((new BigNumber((_947_v67).length)).plus((((_948_v68).contains(_940_i9)) ? ((_948_v68).get(_940_i9)) : (_940_i9))), _938_v61, _940_i9, globalState);
        let _949_v69;
        _949_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("buasua")).length),false);
        let _950_v70;
        _950_v70 = _dafny.Map.Empty.slice().updateUnsafe(_938_v61,true);
        let _951_v71;
        _951_v71 = new _dafny.CodePoint('t'.codePointAt(0));
        let _952_v72;
        _952_v72 = _dafny.Map.Empty.slice().updateUnsafe(_938_v61,_951_v71);
        let _953_v73;
        let _nw154 = Array((new BigNumber(19)).toNumber());
        _nw154[(_dafny.ZERO).toNumber()] = _940_i9;
        _nw154[(_dafny.ONE).toNumber()] = _940_i9;
        _nw154[(new BigNumber(2)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(3)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(4)).toNumber()] = new BigNumber((_879_v20).length);
        _nw154[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_940_i9);
        _nw154[(new BigNumber(6)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(7)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(8)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(9)).toNumber()] = (_939_v62).fm4(_dafny.Seq.of((((_949_v69).contains(new BigNumber((_950_v70).length))) ? ((_949_v69).get(new BigNumber((_950_v70).length))) : (_938_v61))), _940_i9, new BigNumber(301), _948_v68, globalState);
        _nw154[(new BigNumber(10)).toNumber()] = _940_i9;
        _nw154[(new BigNumber(11)).toNumber()] = new BigNumber((_952_v72).length);
        _nw154[(new BigNumber(12)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(13)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(14)).toNumber()] = _940_i9;
        _nw154[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_940_i9);
        _nw154[(new BigNumber(16)).toNumber()] = _878_v19;
        _nw154[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.of(true, _938_v61, true, _938_v61)).length);
        _nw154[(new BigNumber(18)).toNumber()] = _940_i9;
        _953_v73 = _nw154;
        let _954_v74;
        _954_v74 = _dafny.Seq.of(((!(_938_v61)) ? (_953_v73) : (_953_v73)), _953_v73, _953_v73, _953_v73);
        _954_v74 = _dafny.Seq.Concat(_dafny.Seq.Concat(_954_v74, _954_v74), _dafny.Seq.of(_953_v73, _953_v73, (_954_v74)[_module.__default.safeIndex(_878_v19, new BigNumber((_954_v74).length))], _953_v73, _953_v73));
      }
      (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_878_v19, _878_v19));
      if ((_938_v61) || (_938_v61)) {
        let _955_v75;
        _955_v75 = _dafny.Seq.UnicodeFromString("ig");
        _955_v75 = _955_v75;
        (globalState).f1 = _878_v19;
        _878_v19 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_878_v19, _878_v19), _878_v19);
        (globalState).f1 = _878_v19;
        let _956_v76;
        _956_v76 = _dafny.Set.fromElements(_938_v61, !((_dafny.ZERO).minus(_878_v19)).isEqualTo(_878_v19));
        _956_v76 = _956_v76;
      } else {
        _938_v61 = _938_v61;
        let _957_v77;
        _957_v77 = _dafny.Map.Empty.slice().updateUnsafe(_878_v19,_878_v19);
        let _958_v78;
        _958_v78 = _dafny.Map.Empty.slice().updateUnsafe(_957_v77,_938_v61);
        let _959_v79;
        _959_v79 = _dafny.Seq.of(false, _938_v61);
        let _960_v80;
        _960_v80 = _dafny.MultiSet.fromElements(_959_v79);
        _878_v19 = ((_878_v19).plus((_dafny.ZERO).minus(_878_v19))).multipliedBy((_this).fm30(_958_v78, _960_v80, globalState));
        let _961_v81;
        _961_v81 = _dafny.Seq.UnicodeFromString("xcfaexkop");
        let _962_v82;
        let _init24 = ((_963_v19) => function (_964_i12) {
          return (_964_i12).plus(_963_v19);
        })(_878_v19);
        let _nw155 = Array((new BigNumber(5)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw155.length); _i0_24++) {
          _nw155[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _962_v82 = _nw155;
        let _965_v83;
        _965_v83 = _962_v82;
        let _966_v84;
        _966_v84 = _dafny.Map.Empty.slice().updateUnsafe(_938_v61,_965_v83);
        let _967_v85;
        _967_v85 = _dafny.MultiSet.fromElements(new BigNumber((_966_v84).length));
        let _968_v86;
        _968_v86 = _module.D2.create_DC7((_this).fm4(_959_v79, _878_v19, new BigNumber((_959_v79).length), (_967_v85).update(_878_v19, _module.__default.abs(_878_v19)), globalState), _878_v19);
        let _969_v87;
        _969_v87 = _dafny.Map.Empty.slice().updateUnsafe(_961_v81,(_968_v86).dtor_cf7);
        _969_v87 = (_969_v87).update(_961_v81, _878_v19);
        let _970_v88;
        _970_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_961_v81, _module.__default.fm1(_878_v19, _938_v61, new BigNumber(726), globalState))).length),true);
        _970_v88 = (_970_v88).update(_878_v19, _938_v61);
        (globalState).f1 = new BigNumber(-63);
      }
      let _971_v89;
      let _nw156 = Array((new BigNumber(21)).toNumber());
      _971_v89 = _nw156;
      _971_v89 = _971_v89;
      let _972_v90;
      _972_v90 = new _dafny.CodePoint('a'.codePointAt(0));
      let _973_v91;
      let _nw157 = new _module.C0();
      _nw157.__ctor(_938_v61, ((true) ? (_dafny.Seq.of(_879_v20)) : (_module.__default.fm38(_module.__default.fm0(_938_v61, globalState), _938_v61, new BigNumber((_dafny.Seq.UnicodeFromString("bj")).length), _972_v90, globalState))));
      _973_v91 = _nw157;
      r0 = _973_v91;
      r1 = p0;
      r2 = p0;
      return [r0, r1, r2];
    }
    m11(p0, p1, p2, globalState) {
      let _this = this;
      let _974_v0;
      _974_v0 = _dafny.Seq.UnicodeFromString("latlvtxl");
      let _975_v1;
      _975_v1 = _module.D2.create_DC6(_974_v0);
      let _source13 = _975_v1;
      if (_source13.is_DC7) {
        let _976___mcc_h0 = (_source13).cf7;
        let _977___mcc_h1 = (_source13).cf8;
        let _978_cf8 = _977___mcc_h1;
        let _979_cf7 = _976___mcc_h0;
        let _980_v2;
        _980_v2 = _module.D2.create_DC7(p1, (_dafny.ZERO).minus(_978_cf8));
        (globalState).f1 = ((_980_v2).dtor_cf8).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(305), _979_cf7));
        let _981_v3;
        let _nw158 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _981_v3 = _nw158;
        _981_v3 = _981_v3;
        let _982_v4;
        _982_v4 = _module.D2.create_DC8(p2);
        let _pat_let_tv13 = _978_cf8;
        let _source14 = function (_pat_let10_0) {
          return function (_983_dt__update__tmp_h0) {
            return function (_pat_let11_0) {
              return function (_984_dt__update_hcf4_h0) {
                return _module.D1.create_DC4(_984_dt__update_hcf4_h0);
              }(_pat_let11_0);
            }(_pat_let_tv13);
          }(_pat_let10_0);
        }((_this).fm5(_982_v4, globalState));
        if (_source14.is_DC5) {
          let _985___mcc_h5 = (_source14).cf5;
          let _986_cf5 = _985___mcc_h5;
          (globalState).f1 = _module.__default.fm3(p2, p2, globalState);
          let _rhs203 = p2;
          let _rhs204 = _module.__default.safeDivisionInt(_979_cf7, _module.__default.safeModuloInt((_dafny.ZERO).minus(_979_cf7), (_dafny.ZERO).minus(new BigNumber((_974_v0).length))));
          let _lhs141 = globalState;
          _986_cf5 = _rhs203;
          _lhs141.f1 = _rhs204;
          let _987_v5;
          _987_v5 = new _dafny.CodePoint('c'.codePointAt(0));
          let _988_v6;
          _988_v6 = _dafny.Map.Empty.slice().updateUnsafe(_987_v5,_978_cf8);
          _988_v6 = _988_v6;
          (globalState).f1 = p1;
        } else {
          let _989___mcc_h6 = (_source14).cf4;
          let _990_cf4 = _989___mcc_h6;
          let _991_v7;
          _991_v7 = false;
          _991_v7 = p2;
          let _992_v8;
          _992_v8 = _module.D0.create_DC0(p2);
          let _993_v9;
          _993_v9 = _module.D0.create_DC3(_992_v8);
          let _994_v10;
          _994_v10 = _module.D0.create_DC3(_993_v9);
          let _995_v11;
          _995_v11 = _module.D0.create_DC3(_992_v8);
          let _996_v12;
          _996_v12 = _dafny.Seq.of(_991_v7, true, false);
          let _997_v13;
          _997_v13 = _dafny.Seq.of((_module.D0.create_DC3(_992_v8)).dtor_cf3);
          let _998_v14;
          let _init25 = function (_999_i0) {
            return _module.__default.safeDivisionInt(_999_i0, new BigNumber(864));
          };
          let _nw159 = Array((new BigNumber(10)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw159.length); _i0_25++) {
            _nw159[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _998_v14 = _nw159;
          let _1000_v15;
          _1000_v15 = _dafny.Map.Empty.slice().updateUnsafe(_998_v14,p1);
          let _pat_let_tv14 = _997_v13;
          let _pat_let_tv15 = _990_cf4;
          let _pat_let_tv16 = _997_v13;
          let _rhs205 = function (_pat_let12_0) {
            return function (_1001_dt__update__tmp_h1) {
              return function (_pat_let13_0) {
                return function (_1002_dt__update_hcf3_h0) {
                  return _module.D0.create_DC3(_1002_dt__update_hcf3_h0);
                }(_pat_let13_0);
              }((_pat_let_tv14)[_module.__default.safeIndex(_pat_let_tv15, new BigNumber((_pat_let_tv16).length))]);
            }(_pat_let12_0);
          }(_995_v11);
          let _rhs206 = ((_dafny.ZERO).minus(new BigNumber((_1000_v15).length))).plus(_990_cf4);
          let _rhs207 = _996_v12;
          _995_v11 = _rhs205;
          _978_cf8 = _rhs206;
          _996_v12 = _rhs207;
          _991_v7 = p2;
          _974_v0 = _974_v0;
        }
        let _1003_v16;
        _1003_v16 = _module.D0.create_DC1();
        let _1004_v17;
        let _init26 = ((_1005_v0) => function (_1006_i1) {
          return _dafny.Seq.Concat((_module.D2.create_DC6(_dafny.Seq.UnicodeFromString("vgtgm"))).dtor_cf6, _1005_v0);
        })(_974_v0);
        let _nw160 = Array((new BigNumber(14)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw160.length); _i0_26++) {
          _nw160[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1004_v17 = _nw160;
        let _1007_v18;
        _1007_v18 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index180 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1004_v17).length));
        (_1004_v17)[_index180] = _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm1(p1, p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-826)), function (_1008_i2) {
          return (_module.D11.create_DC19(new _dafny.CodePoint('q'.codePointAt(0)))).dtor_cf20;
        })).length), globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm1(p1, p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-826)), function (_1009_i2) {
          return (_module.D11.create_DC19(new _dafny.CodePoint('q'.codePointAt(0)))).dtor_cf20;
        })).length), globalState)).length)), _1007_v18), _dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_1010_v18) => function (_1011_i3) {
          return _1010_v18;
        })(_1007_v18)));
        let _1012_v19;
        _1012_v19 = _dafny.Map.Empty.slice().updateUnsafe(_978_cf8,p2);
        let _index181 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1004_v17).length));
        let _rhs208 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(p1, _979_cf7), (new BigNumber((_1012_v19).length)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("vfndmpit")).length))));
        let _rhs209 = _1003_v16;
        let _rhs210 = _974_v0;
        let _lhs142 = globalState;
        let _lhs143 = _1004_v17;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1004_v17).length));
        _lhs142.f1 = _rhs208;
        _1003_v16 = _rhs209;
        _lhs143[_lhs144] = _rhs210;
      } else if (_source13.is_DC8) {
        let _1013___mcc_h2 = (_source13).cf9;
        let _1014_cf9 = _1013___mcc_h2;
        let _1015_v20;
        _1015_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1014_cf9,p2);
        if ((((_1015_v20).contains(p2)) ? ((_1015_v20).get(p2)) : (p2))) {
          let _1016_v21;
          _1016_v21 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1017_v22;
          _1017_v22 = _dafny.Map.Empty.slice().updateUnsafe(_974_v0,_1016_v21);
          (globalState).f1 = new BigNumber((_1017_v22).length);
          let _1018_v23;
          _1018_v23 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          let _1019_v24;
          _1019_v24 = _module.D11.create_DC19(_1016_v21);
          let _rhs211 = (_1018_v23).Merge(_1018_v23);
          let _rhs212 = _module.D11.create_DC19(_1016_v21);
          let _rhs213 = _974_v0;
          let _rhs214 = _dafny.Seq.update(_dafny.Seq.Concat(_974_v0, (_module.D2.create_DC6(_974_v0)).dtor_cf6), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_974_v0, (_module.D2.create_DC6(_974_v0)).dtor_cf6)).length)), _1016_v21);
          _1018_v23 = _rhs211;
          _1019_v24 = _rhs212;
          _974_v0 = _rhs213;
          _974_v0 = _rhs214;
          let _1020_v25;
          let _init27 = ((_1021_p0) => function (_1022_i4) {
            return _1021_p0;
          })(p0);
          let _nw161 = Array((new BigNumber(8)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw161.length); _i0_27++) {
            _nw161[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1020_v25 = _nw161;
          let _1023_v26;
          _1023_v26 = _dafny.MultiSet.fromElements(false, p0, p0);
          let _index182 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1020_v25).length));
          (_1020_v25)[_index182] = (_1023_v26).IsProperSubsetOf(_module.__default.fm34(p2, globalState));
          let _1024_v27;
          _1024_v27 = _dafny.Map.Empty.slice().updateUnsafe(_974_v0,!(new BigNumber((_1015_v20).length)).isEqualTo(p1));
          let _index183 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1020_v25).length));
          (_1020_v25)[_index183] = (((_1024_v27).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_1027_v21) => function (_1028_i5) {
            return _1027_v21;
          })(_1016_v21)))) ? ((_1024_v27).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_1025_v21) => function (_1026_i5) {
            return _1025_v21;
          })(_1016_v21)))) : (_1014_cf9));
          (globalState).f1 = (p1).multipliedBy(p1);
          let _1029_v28;
          _1029_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1020_v25,_1016_v21);
          let _1030_v29;
          _1030_v29 = _module.D0.create_DC2(_1029_v28, _1016_v21);
          let _1031_v30;
          _1031_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1030_v29,_1014_cf9);
          _1031_v30 = _1031_v30;
        } else {
          (globalState).f1 = _module.__default.fm3(p0, _1014_cf9, globalState);
          let _1032_v31;
          let _nw162 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1032_v31 = _nw162;
          let _1033_v32;
          _1033_v32 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1034_v33;
          let _1035_v34;
          let _out30;
          let _out31;
          let _outcollector14 = _module.__default.m0(_1032_v31, _1033_v32, p0, _1033_v32, globalState);
          _out30 = _outcollector14[0];
          _out31 = _outcollector14[1];
          _1034_v33 = _out30;
          _1035_v34 = _out31;
          (globalState).f1 = (_dafny.ZERO).minus(p1);
          let _1036_v35;
          _1036_v35 = _dafny.Map.Empty.slice().updateUnsafe(_974_v0,p0);
          let _1037_v36;
          let _1038_v37;
          let _out32;
          let _out33;
          let _outcollector15 = _module.__default.m0(_1032_v31, _1033_v32, (((_1036_v35).contains(_974_v0)) ? ((_1036_v35).get(_974_v0)) : (_1035_v34)), _1033_v32, globalState);
          _out32 = _outcollector15[0];
          _out33 = _outcollector15[1];
          _1037_v36 = _out32;
          _1038_v37 = _out33;
          let _1039_v38;
          _1039_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(699));
          _1039_v38 = (_1039_v38).update(_1014_cf9, new BigNumber((_974_v0).length));
        }
        let _1040_v39;
        _1040_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _1041_v40;
        _1041_v40 = _1040_v39;
        let _1042_v41;
        _1042_v41 = _module.D6.create_DC14(false);
        let _1043_v42;
        _1043_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1041_v40),(_1042_v41).dtor_cf15);
        let _1044_v43;
        _1044_v43 = _dafny.Seq.of(p0);
        let _1045_v44;
        _1045_v44 = _dafny.MultiSet.fromElements(_1044_v43);
        let _1046_v45;
        _1046_v45 = _dafny.MultiSet.fromElements((_this).fm30(_1043_v42, _1045_v44, globalState), p1);
        let _1047_v46;
        let _nw163 = Array((new BigNumber(14)).toNumber());
        _nw163[(_dafny.ZERO).toNumber()] = p1;
        _nw163[(_dafny.ONE).toNumber()] = p1;
        _nw163[(new BigNumber(2)).toNumber()] = p1;
        _nw163[(new BigNumber(3)).toNumber()] = p1;
        _nw163[(new BigNumber(4)).toNumber()] = p1;
        _nw163[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("rodcjqa")).length);
        _nw163[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw163[(new BigNumber(7)).toNumber()] = new BigNumber(96);
        _nw163[(new BigNumber(8)).toNumber()] = p1;
        _nw163[(new BigNumber(9)).toNumber()] = p1;
        _nw163[(new BigNumber(10)).toNumber()] = p1;
        _nw163[(new BigNumber(11)).toNumber()] = new BigNumber((_974_v0).length);
        _nw163[(new BigNumber(12)).toNumber()] = p1;
        _nw163[(new BigNumber(13)).toNumber()] = p1;
        _1047_v46 = _nw163;
        let _1048_v47;
        _1048_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1047_v46,false);
        _1046_v45 = _dafny.MultiSet.fromElements(new BigNumber((_1048_v47).length));
        let _1049_v48;
        _1049_v48 = _dafny.Seq.of(_974_v0, _dafny.Seq.UnicodeFromString("i"), _dafny.Seq.Concat(_974_v0, _974_v0));
        _1049_v48 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-320)), ((_1050_v0, _1051_p1) => function (_1052_i6) {
          return _dafny.Seq.update(_1050_v0, _module.__default.safeIndex(_1051_p1, new BigNumber((_1050_v0).length)), new _dafny.CodePoint('o'.codePointAt(0)));
        })(_974_v0, p1));
        let _1053_v49;
        _1053_v49 = _dafny.Seq.of(p1);
        _1053_v49 = _1053_v49;
      } else if (_source13.is_DC6) {
        let _1054___mcc_h3 = (_source13).cf6;
        let _1055_cf6 = _1054___mcc_h3;
        let _1056_v50;
        _1056_v50 = true;
        _1056_v50 = _1056_v50;
        let _1057_v51;
        _1057_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1056_v50,p0);
        let _1058_v52;
        _1058_v52 = _module.D12.create_DC21(_1057_v51);
        let _1059_v53;
        _1059_v53 = _dafny.Set.fromElements(_1057_v51, (_1058_v52).dtor_cf23);
        (globalState).f1 = new BigNumber(((_1059_v53).Intersect(_1059_v53)).length);
        let _1060_v54;
        let _nw164 = new _module.C3();
        _nw164.__ctor(p2, true);
        _1060_v54 = _nw164;
        let _1061_v55;
        _1061_v55 = _dafny.Seq.of(p2, (_1060_v54).f13, (_1060_v54).f14);
        _1061_v55 = _dafny.Seq.Concat(_1061_v55, _1061_v55);
      } else {
        let _1062___mcc_h4 = (_source13).cf10;
        let _1063_cf10 = _1062___mcc_h4;
        let _1064_v56;
        _1064_v56 = _dafny.Seq.of(p0);
        _1064_v56 = _dafny.Seq.of(p2, p0, p0, p2);
        let _1065_v57;
        _1065_v57 = new _dafny.CodePoint('p'.codePointAt(0));
        let _1066_v58;
        let _nw165 = Array((new BigNumber(9)).toNumber());
        _nw165[(_dafny.ZERO).toNumber()] = true;
        _nw165[(_dafny.ONE).toNumber()] = p0;
        _nw165[(new BigNumber(2)).toNumber()] = p2;
        _nw165[(new BigNumber(3)).toNumber()] = (p1).isLessThan(p1);
        _nw165[(new BigNumber(4)).toNumber()] = p0;
        _nw165[(new BigNumber(5)).toNumber()] = !(true) || (p2);
        _nw165[(new BigNumber(6)).toNumber()] = !(!(!(_dafny.Seq.contains(_974_v0, _1065_v57))));
        _nw165[(new BigNumber(7)).toNumber()] = p2;
        _nw165[(new BigNumber(8)).toNumber()] = ((p0) ? (!(p0)) : (p0));
        _1066_v58 = _nw165;
        let _index184 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1066_v58).length));
        (_1066_v58)[_index184] = p2;
        let _index185 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1066_v58).length));
        let _rhs215 = p0;
        let _rhs216 = _1065_v57;
        let _lhs145 = _1066_v58;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1066_v58).length));
        _lhs145[_lhs146] = _rhs215;
        _1065_v57 = _rhs216;
        let _index186 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1066_v58).length));
        (_1066_v58)[_index186] = (_1064_v56)[_module.__default.safeIndex(p1, new BigNumber((_1064_v56).length))];
        let _1067_v59;
        _1067_v59 = _dafny.Seq.of(p1);
        let _1068_v60;
        _1068_v60 = _dafny.Seq.of(_1067_v59, _dafny.Seq.of(new BigNumber(743)));
        let _1069_v61;
        let _nw166 = new _module.C0();
        _nw166.__ctor((_1066_v58)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1066_v58).length))], _1068_v60);
        _1069_v61 = _nw166;
      }
      let _1070_i7;
      _1070_i7 = _dafny.ZERO;
      L13: {
        while ((p2) === (p0)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1070_i7)) {
              break L13;
            }
            _1070_i7 = (_1070_i7).plus(_dafny.ONE);
            let _1071_v62;
            _1071_v62 = false;
            _1071_v62 = p2;
            let _1072_v63;
            _1072_v63 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_1071_v62, p2, p0));
            let _1073_v64;
            _1073_v64 = _module.D11.create_DC20(_1071_v62, p1);
            let _1074_v65;
            _1074_v65 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus((_1073_v64).dtor_cf22));
            let _1075_v66;
            _1075_v66 = _dafny.Map.Empty.slice().updateUnsafe((_1072_v63).IsDisjointFrom(_1072_v63),new BigNumber((_1074_v65).length));
            _1075_v66 = (_1075_v66).update(_1071_v62, p1);
            _975_v1 = _module.D2.create_DC6(_974_v0);
            (globalState).f1 = p1;
          }
        }
      }
      if (!(p0) || (p2)) {
        let _1076_v67;
        _1076_v67 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1077_v68;
        let _nw167 = new _module.C2();
        _nw167.__ctor(_1076_v67, true);
        _1077_v68 = _nw167;
        let _1078_v69;
        _1078_v69 = _module.D0.create_DC0(true);
        let _1079_v70;
        _1079_v70 = _module.D0.create_DC3(_1078_v69);
        let _pat_let_tv17 = _1078_v69;
        let _source15 = function (_pat_let14_0) {
          return function (_1080_dt__update__tmp_h2) {
            return function (_pat_let15_0) {
              return function (_1081_dt__update_hcf3_h1) {
                return _module.D0.create_DC3(_1081_dt__update_hcf3_h1);
              }(_pat_let15_0);
            }(_pat_let_tv17);
          }(_pat_let14_0);
        }(_1079_v70);
        if (_source15.is_DC1) {
          let _1082_v71;
          _1082_v71 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1083_v72;
          _1083_v72 = _module.D12.create_DC21(_1082_v71);
          _1083_v72 = _1083_v72;
          let _1084_v73;
          _1084_v73 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm35(_974_v0, _dafny.MultiSet.fromElements(p1, (_dafny.ZERO).minus(new BigNumber(-114))), _1076_v67, globalState));
          let _1085_v74;
          _1085_v74 = _dafny.Set.fromElements(p1, p1);
          _1084_v73 = (_1084_v73).update(new BigNumber((_dafny.Seq.Concat(_974_v0, _974_v0)).length), _1085_v74);
          let _1086_v75;
          _1086_v75 = true;
          _1086_v75 = (((p2) ? (p2) : (p0))) || (p2);
          let _1087_v76;
          let _init28 = ((_1088_p1) => function (_1089_i8) {
            return _module.__default.safeModuloInt(_1089_i8, (_dafny.ZERO).minus(_1088_p1));
          })(p1);
          let _nw168 = Array((new BigNumber(18)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw168.length); _i0_28++) {
            _nw168[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1087_v76 = _nw168;
          let _index187 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1087_v76).length));
          (_1087_v76)[_index187] = (_dafny.ZERO).minus(p1);
          let _index188 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_1087_v76).length));
          (_1087_v76)[_index188] = p1;
        } else if (_source15.is_DC2) {
          let _1090___mcc_h7 = (_source15).cf1;
          let _1091___mcc_h8 = (_source15).cf2;
          let _1092_cf2 = _1091___mcc_h8;
          let _1093_cf1 = _1090___mcc_h7;
          let _1094_v77;
          let _nw169 = Array((new BigNumber(7)).toNumber());
          _nw169[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), ((_1095_cf2) => function (_1096_i9) {
            return _1095_cf2;
          })(_1092_cf2));
          _nw169[(_dafny.ONE).toNumber()] = _974_v0;
          _nw169[(new BigNumber(2)).toNumber()] = _974_v0;
          _nw169[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_974_v0, _module.__default.safeIndex(new BigNumber(-556), new BigNumber((_974_v0).length)), _1092_cf2);
          _nw169[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_974_v0, _974_v0);
          _nw169[(new BigNumber(5)).toNumber()] = _974_v0;
          _nw169[(new BigNumber(6)).toNumber()] = _974_v0;
          _1094_v77 = _nw169;
          let _index189 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1094_v77).length));
          (_1094_v77)[_index189] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-35)), ((_1097_v67) => function (_1098_i10) {
            return _1097_v67;
          })(_1076_v67)), _974_v0);
          let _index190 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_1094_v77).length));
          (_1094_v77)[_index190] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("remd"), _974_v0);
          let _1099_v78;
          _1099_v78 = false;
          let _1100_v79;
          _1100_v79 = _dafny.Seq.of(p0);
          let _1101_v80;
          _1101_v80 = _dafny.Seq.of(!(_module.__default.fm0(_1099_v78, globalState)), false, !(p0), _1099_v78);
          let _1102_v81;
          let _nw170 = Array((new BigNumber(13)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _1100_v79;
          _nw170[(_dafny.ONE).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(2)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(3)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(4)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(5)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(6)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(7)).toNumber()] = _1101_v80;
          _nw170[(new BigNumber(8)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(9)).toNumber()] = _module.__default.fm39(new BigNumber((_dafny.Seq.UnicodeFromString("xun")).length), p0, p1, p1, globalState);
          _nw170[(new BigNumber(10)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(11)).toNumber()] = _1100_v79;
          _nw170[(new BigNumber(12)).toNumber()] = _1100_v79;
          _1102_v81 = _nw170;
          let _1103_v82;
          _1103_v82 = _dafny.Set.fromElements(new BigNumber(((_1094_v77)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_1094_v77).length))]).length));
          let _rhs217 = (_1077_v68).fm8(true, _1099_v78, (p1).isEqualTo(new BigNumber((_1103_v82).length)), globalState);
          let _rhs218 = !(p1).isEqualTo((p1).plus(p1));
          let _rhs219 = _1102_v81;
          let _lhs147 = globalState;
          _lhs147.f1 = _rhs217;
          _1099_v78 = _rhs218;
          _1102_v81 = _rhs219;
          let _1104_v83;
          let _nw171 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1104_v83 = _nw171;
          let _index191 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length));
          (_1104_v83)[_index191] = p1;
          let _index192 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length));
          (_1104_v83)[_index192] = (new BigNumber(((_1094_v77)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_1094_v77).length))]).length)).multipliedBy(p1);
          let _1105_v84;
          let _nw172 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
          _1105_v84 = _nw172;
          let _index193 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1105_v84).length));
          (_1105_v84)[_index193] = _1103_v82;
          let _1106_v86;
          _1106_v86 = _dafny.Map.Empty.slice().updateUnsafe((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))],(_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))]);
          let _1107_v87;
          _1107_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1106_v86,_1099_v78);
          let _1108_v88;
          _1108_v88 = _dafny.Seq.of(_1101_v80);
          let _1109_v89;
          _1109_v89 = _dafny.MultiSet.fromElements((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))], p1, (_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))], p1, p1);
          let _1110_v90;
          _1110_v90 = _dafny.Seq.of((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))], p1);
          let _index194 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1105_v84).length));
          let _rhs220 = _module.__default.safeModuloInt(p1, ((_this).fm30(function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of (_1107_v87).Keys.Elements) {
              let _1111_v85 = _compr_34;
              if ((_1107_v87).contains(_1111_v85)) {
                _coll34.push([_1111_v85,_1099_v78]);
              }
            }
            return _coll34;
          }(), _dafny.MultiSet.FromArray(_1108_v88), globalState)).multipliedBy(new BigNumber(362)));
          let _rhs221 = _module.__default.fm35(_974_v0, _1109_v89, _1076_v67, globalState);
          let _rhs222 = (((_1109_v89).contains(((_1099_v78) ? ((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))]) : (new BigNumber((_1109_v89).cardinality()))))) ? ((_1109_v89).get(((_1099_v78) ? ((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))]) : (new BigNumber((_1109_v89).cardinality()))))) : (_module.__default.safeDivisionInt(new BigNumber(191), p1)));
          let _rhs223 = !(((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))]).isEqualTo((_1110_v90)[_module.__default.safeIndex((_1104_v83)[_module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length))], new BigNumber((_1110_v90).length))]));
          let _lhs148 = _1104_v83;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_1104_v83).length));
          let _lhs150 = _1105_v84;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1105_v84).length));
          let _lhs152 = globalState;
          _lhs148[_lhs149] = _rhs220;
          _lhs150[_lhs151] = _rhs221;
          _lhs152.f1 = _rhs222;
          _1099_v78 = _rhs223;
        } else if (_source15.is_DC0) {
          let _1112___mcc_h9 = (_source15).cf0;
          let _1113_cf0 = _1112___mcc_h9;
          let _1114_v91;
          let _nw173 = Array((new BigNumber(18)).toNumber()).fill([]);
          _1114_v91 = _nw173;
          let _1115_v92;
          let _init29 = ((_1116_p2) => function (_1117_i11) {
            return _1116_p2;
          })(p2);
          let _nw174 = Array((_dafny.ONE).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw174.length); _i0_29++) {
            _nw174[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1115_v92 = _nw174;
          let _index196 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1114_v91).length));
          (_1114_v91)[_index196] = _1115_v92;
          let _index197 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1114_v91).length));
          let _nw175 = Array((new BigNumber(5)).toNumber()).fill(false);
          (_1114_v91)[_index197] = _nw175;
          let _1118_v93;
          _1118_v93 = _module.D11.create_DC19(_1076_v67);
          let _1119_v94;
          _1119_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v93,_module.__default.safeModuloInt(p1, p1));
          let _1120_v95;
          _1120_v95 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lrxmykw")).length), p1);
          _1119_v94 = (_1119_v94).update(_1118_v93, (((_1120_v95).contains((_this).fm4(_dafny.Seq.of(_1113_cf0, p0), p1, p1, (_1120_v95).update(p1, _module.__default.abs(p1)), globalState))) ? ((_1120_v95).get((_this).fm4(_dafny.Seq.of(_1113_cf0, p0), p1, p1, (_1120_v95).update(p1, _module.__default.abs(p1)), globalState))) : (p1)));
          let _1121_v96;
          let _init30 = ((_1122_p1) => function (_1123_i12) {
            return _module.__default.safeModuloInt(_1123_i12, _1122_p1);
          })(p1);
          let _nw176 = Array((new BigNumber(29)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw176.length); _i0_30++) {
            _nw176[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1121_v96 = _nw176;
          let _index198 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1121_v96).length));
          (_1121_v96)[_index198] = p1;
          let _index199 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1121_v96).length));
          let _rhs224 = _module.__default.fm3(p2, p0, globalState);
          let _rhs225 = (new BigNumber(970)).isLessThan(p1);
          let _rhs226 = _1121_v96;
          let _lhs153 = _1121_v96;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1121_v96).length));
          _lhs153[_lhs154] = _rhs224;
          _1113_cf0 = _rhs225;
          _1121_v96 = _rhs226;
          (globalState).f1 = new BigNumber((_module.__default.fm40(globalState)).length);
        } else {
          let _1124___mcc_h10 = (_source15).cf3;
          let _1125_cf3 = _1124___mcc_h10;
          let _1126_v97;
          _1126_v97 = _dafny.Set.fromElements((_dafny.ZERO).minus(p1));
          let _1127_v98;
          let _nw177 = Array((new BigNumber(16)).toNumber());
          _nw177[(_dafny.ZERO).toNumber()] = new BigNumber((_974_v0).length);
          _nw177[(_dafny.ONE).toNumber()] = p1;
          _nw177[(new BigNumber(2)).toNumber()] = p1;
          _nw177[(new BigNumber(3)).toNumber()] = new BigNumber(((_1126_v97).Difference(_1126_v97)).length);
          _nw177[(new BigNumber(4)).toNumber()] = p1;
          _nw177[(new BigNumber(5)).toNumber()] = (p1).multipliedBy(p1);
          _nw177[(new BigNumber(6)).toNumber()] = p1;
          _nw177[(new BigNumber(7)).toNumber()] = p1;
          _nw177[(new BigNumber(8)).toNumber()] = p1;
          _nw177[(new BigNumber(9)).toNumber()] = p1;
          _nw177[(new BigNumber(10)).toNumber()] = p1;
          _nw177[(new BigNumber(11)).toNumber()] = p1;
          _nw177[(new BigNumber(12)).toNumber()] = (p1).minus(new BigNumber(-580));
          _nw177[(new BigNumber(13)).toNumber()] = p1;
          _nw177[(new BigNumber(14)).toNumber()] = p1;
          _nw177[(new BigNumber(15)).toNumber()] = p1;
          _1127_v98 = _nw177;
          let _index200 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1127_v98).length));
          (_1127_v98)[_index200] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((!(true)) ? ((_dafny.ZERO).minus(p1)) : (p1))));
          let _index201 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1127_v98).length));
          (_1127_v98)[_index201] = (p1).minus(_module.__default.safeDivisionInt(p1, p1));
          let _1128_v99;
          _1128_v99 = true;
          let _1129_v100;
          _1129_v100 = _dafny.Seq.of(_1128_v99, p2);
          _1128_v99 = (new BigNumber((_1129_v100).length)).isLessThan(new BigNumber(589));
          let _1130_v101;
          let _nw178 = new _module.C2();
          _nw178.__ctor((_974_v0)[_module.__default.safeIndex((_1127_v98)[_module.__default.safeIndex(new BigNumber(796), new BigNumber((_1127_v98).length))], new BigNumber((_974_v0).length))], !(_module.__default.fm0(p0, globalState)));
          _1130_v101 = _nw178;
          let _1131_v102;
          let _nw179 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _1131_v102 = _nw179;
          let _1132_v103;
          let _nw180 = Array((new BigNumber(8)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = p0;
          _nw180[(_dafny.ONE).toNumber()] = p2;
          _nw180[(new BigNumber(2)).toNumber()] = p2;
          _nw180[(new BigNumber(3)).toNumber()] = p2;
          _nw180[(new BigNumber(4)).toNumber()] = p2;
          _nw180[(new BigNumber(5)).toNumber()] = false;
          _nw180[(new BigNumber(6)).toNumber()] = !(p0);
          _nw180[(new BigNumber(7)).toNumber()] = _1128_v99;
          _1132_v103 = _nw180;
          let _1133_v104;
          _1133_v104 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1132_v103);
          let _index202 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1131_v102).length));
          (_1131_v102)[_index202] = (_1133_v104).Merge(_1133_v104);
          let _index203 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1131_v102).length));
          (_1131_v102)[_index203] = _1133_v104;
        }
        let _1134_v105;
        let _nw181 = new _module.C0();
        _nw181.__ctor(p2, _dafny.Seq.of(_dafny.Seq.of(p1)));
        _1134_v105 = _nw181;
        let _1135_v106;
        let _nw182 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1135_v106 = _nw182;
        let _1136_v107;
        _1136_v107 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1135_v106);
        let _1137_v108;
        _1137_v108 = _dafny.Seq.of(_1135_v106);
        let _1138_v109;
        _1138_v109 = (((_1136_v107).contains(p1)) ? ((_1136_v107).get(p1)) : ((_1137_v108)[_module.__default.safeIndex(p1, new BigNumber((_1137_v108).length))]));
        let _1139_v110;
        _1139_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1138_v109,_1135_v106);
        let _1140_v111;
        _1140_v111 = _1139_v110;
        _1139_v110 = (_1140_v111);
        let _1141_v112;
        let _init31 = ((_1142_p1) => function (_1143_i13) {
          return (_1143_i13).minus(new BigNumber((_dafny.Seq.of(_1142_p1)).length));
        })(p1);
        let _nw183 = Array((new BigNumber(20)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw183.length); _i0_31++) {
          _nw183[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1141_v112 = _nw183;
        let _1144_v113;
        _1144_v113 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1141_v112);
        _1141_v112 = (((_1144_v113).contains(!(p0))) ? ((_1144_v113).get(!(p0))) : (_1141_v112));
      } else {
        let _1145_v114;
        _1145_v114 = _dafny.MultiSet.fromElements(p2);
        (globalState).f4 = _1145_v114;
        let _1146_v115;
        _1146_v115 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(663)), function (_1147_i14) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
        let _1148_v116;
        _1148_v116 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1146_v115);
        _1148_v116 = (_1148_v116).update(_dafny.areEqual(_974_v0, _dafny.Seq.UnicodeFromString("bbc")), (_1146_v115).Union(_1146_v115));
        let _1149_v117;
        _1149_v117 = _module.D0.create_DC0(p0);
        let _1150_v118;
        _1150_v118 = _dafny.Seq.of(_1149_v117);
        let _1151_v119;
        _1151_v119 = _dafny.Seq.of(p0);
        let _1152_v120;
        _1152_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v118,_dafny.MultiSet.fromElements(_1151_v119));
        let _1153_v121;
        _1153_v121 = _dafny.MultiSet.fromElements(_module.__default.fm2(p2, p0, globalState));
        _1152_v120 = (_1152_v120).update(_1150_v118, (_dafny.MultiSet.fromElements(_1151_v119, _1151_v119)).Intersect(_1153_v121));
        let _1154_v122;
        let _init32 = ((_1155_p2) => function (_1156_i15) {
          return !(_1155_p2) || (_1155_p2);
        })(p2);
        let _nw184 = Array((new BigNumber(26)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw184.length); _i0_32++) {
          _nw184[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1154_v122 = _nw184;
        let _index204 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1154_v122).length));
        (_1154_v122)[_index204] = ((p0) ? (p2) : (p2));
        let _index205 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1154_v122).length));
        (_1154_v122)[_index205] = true;
        let _1157_v123;
        _1157_v123 = _dafny.Set.fromElements(_dafny.Seq.Concat(_974_v0, _974_v0));
        let _1158_v124;
        _1158_v124 = new _dafny.CodePoint('n'.codePointAt(0));
        let _index206 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1154_v122).length));
        let _rhs227 = _1157_v123;
        let _rhs228 = _module.__default.fm0((new BigNumber(65)).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))), globalState);
        let _rhs229 = _1158_v124;
        let _lhs155 = _1154_v122;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1154_v122).length));
        _1157_v123 = _rhs227;
        _lhs155[_lhs156] = _rhs228;
        _1158_v124 = _rhs229;
      }
      let _1159_v125;
      let _init33 = ((_1160_p1) => function (_1161_i16) {
        return (_1161_i16).minus((_dafny.ZERO).minus(_1160_p1));
      })(p1);
      let _nw185 = Array((new BigNumber(7)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw185.length); _i0_33++) {
        _nw185[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1159_v125 = _nw185;
      let _1162_v126;
      _1162_v126 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
      let _1163_v127;
      _1163_v127 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p2, globalState),new BigNumber((_1162_v126).length));
      let _index207 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length));
      (_1159_v125)[_index207] = (new BigNumber((_1163_v127).length)).minus(p1);
      let _1164_v128;
      _1164_v128 = _dafny.Map.Empty.slice().updateUnsafe(p1,_974_v0);
      let _1165_v129;
      _1165_v129 = _dafny.Seq.of(p2);
      let _1166_v130;
      _1166_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1165_v129,new BigNumber(396));
      let _index208 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length));
      (_1159_v125)[_index208] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_974_v0, (((_1164_v128).contains(new BigNumber((_1166_v130).length))) ? ((_1164_v128).get(new BigNumber((_1166_v130).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), function (_1167_i17) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mvaeqnmre"), _974_v0))).length);
      let _1168_v131;
      _1168_v131 = _1165_v129;
      let _source16 = _1168_v131;
      let _1169___mcc_h11 = _source16;
      let _1170_cf16 = _1169___mcc_h11;
      let _1171_v132;
      _1171_v132 = false;
      _1171_v132 = p2;
      let _1172_v133;
      _1172_v133 = _module.D11.create_DC20(p2, new BigNumber(730));
      let _index209 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length));
      (_1159_v125)[_index209] = _module.__default.safeModuloInt((_1172_v133).dtor_cf22, (_1159_v125)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length))]);
      let _index210 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length));
      (_1159_v125)[_index210] = (_1159_v125)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length))];
      let _1173_v134;
      _1173_v134 = _dafny.Seq.of(p1);
      let _1174_v135;
      _1174_v135 = _dafny.Seq.of(_1173_v134, _1173_v134, _1173_v134);
      let _1175_v136;
      let _nw186 = new _module.C0();
      _nw186.__ctor(p2, _1174_v135);
      _1175_v136 = _nw186;
      let _1176_v137;
      _1176_v137 = _dafny.Seq.of(p1, new BigNumber(709), (_1159_v125)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length))]);
      let _1177_v138;
      _1177_v138 = _dafny.Seq.of(_1176_v137, _1176_v137);
      let _1178_v139;
      let _nw187 = new _module.C0();
      _nw187.__ctor((new BigNumber((_1176_v137).length)).isLessThan((_1159_v125)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1159_v125).length))]), _dafny.Seq.Concat(_1177_v138, _1177_v138));
      _1178_v139 = _nw187;
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
      this.f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f15, f5, f6) {
      let _this = this;
      (_this).f15 = f15;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f15;
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm7(p0, globalState) {
      let _this = this;
      return new BigNumber(((((_this).f6) ? (_dafny.Seq.UnicodeFromString("mbq")) : (_dafny.Seq.UnicodeFromString("pt")))).length);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f15;
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4((_dafny.ZERO).minus(_this.f15));
    };
    fm42(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f15,function () {
        let _coll35 = new _dafny.Set();
        for (const _compr_35 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_1179_i0) {
          return (_this).f5;
        })).length),_this.f15),(_this).f5)).Keys.Elements) {
          let _1180_v0 = _compr_35;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_1179_i0) {
            return (_this).f5;
          })).length),_this.f15),(_this).f5)).contains(_1180_v0)) {
            _coll35.add(_1180_v0);
          }
        }
        return _coll35;
      }())).Merge(function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of (_dafny.Seq.of((_module.D1.create_DC4(_this.f15)).dtor_cf4)).Elements) {
          let _1181_v1 = _compr_36;
          if (_dafny.Seq.contains(_dafny.Seq.of((_module.D1.create_DC4(_this.f15)).dtor_cf4), _1181_v1)) {
            _coll36.push([_module.__default.safeDivisionInt(_1181_v1, _this.f15),_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15), _dafny.Map.Empty.slice().updateUnsafe(_this.f15,new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)))]);
          }
        }
        return _coll36;
      }())).Merge((function () {
        let _coll37 = new _dafny.Map();
        for (const _compr_37 of _dafny.IntegerRange(new BigNumber(844), new BigNumber(86))) {
          let _1182_v2 = _compr_37;
          if (((new BigNumber(844)).isLessThanOrEqualTo(_1182_v2)) && ((_1182_v2).isLessThan(new BigNumber(86)))) {
            _coll37.push([_module.__default.safeModuloInt(_1182_v2, new BigNumber((_dafny.Seq.UnicodeFromString("jpe")).length)),function () {
              let _coll38 = new _dafny.Set();
              for (const _compr_38 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),_this.f15), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f15)).length),_this.f15))).Elements) {
                let _1183_v3 = _compr_38;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),_this.f15), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f15)).length),_this.f15)), _1183_v3)) {
                  _coll38.add(_1183_v3);
                }
              }
              return _coll38;
            }()]);
          }
        }
        return _coll37;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f15,function () {
        let _coll39 = new _dafny.Set();
        for (const _compr_39 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15),_this.f15)).Keys.Elements) {
          let _1184_v4 = _compr_39;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15),_this.f15)).contains(_1184_v4)) {
            _coll39.add(_1184_v4);
          }
        }
        return _coll39;
      }())));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _1185_v0;
      _1185_v0 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f15,true)).length));
      let _hi6 = p1;
      for (let _1186_i0 = (new BigNumber(535)).multipliedBy(new BigNumber((_1185_v0).length)); _1186_i0.isLessThan(_hi6); _1186_i0 = _1186_i0.plus(_dafny.ONE)) {
        r1 = ((!((_1186_i0).isLessThanOrEqualTo(new BigNumber((p2).length)))) ? (p2) : (_dafny.Seq.UnicodeFromString("bpeabsjw")));
        let _1187_v1;
        let _nw188 = new _module.C1();
        _nw188.__ctor();
        _1187_v1 = _nw188;
        let _1188_v2;
        let _init34 = function (_1189_i1) {
          return (_this).f6;
        };
        let _nw189 = Array((new BigNumber(8)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw189.length); _i0_34++) {
          _nw189[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1188_v2 = _nw189;
        let _1190_v3;
        _1190_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1187_v1,_1188_v2);
        let _source17 = _1190_v3;
        let _1191___mcc_h0 = _source17;
        let _1192_cf18 = _1191___mcc_h0;
        r1 = _dafny.Seq.Concat(p2, p2);
        let _1193_v4;
        _1193_v4 = _dafny.MultiSet.fromElements(false);
        let _1194_v5;
        _1194_v5 = _dafny.Seq.of(_module.__default.fm0((_this).f6, globalState));
        (globalState).f1 = (new BigNumber((_1193_v4).cardinality())).multipliedBy(new BigNumber((_1194_v5).length));
        r0 = p0;
        let _1195_v6;
        _1195_v6 = _dafny.MultiSet.fromElements(_this.f15);
        let _1196_v7;
        _1196_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1186_i0);
        let _1197_v8;
        _1197_v8 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-800)),p1);
        let _1198_v9;
        _1198_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1197_v8,(_this).f6);
        let _1199_v10;
        _1199_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1193_v4).IsProperSubsetOf(_1193_v4),(_1195_v6).IsSubsetOf(((_1195_v6).update(new BigNumber(-251), _module.__default.abs(new BigNumber((_1196_v7).length)))).update(p1, _module.__default.abs(new BigNumber((_1198_v9).length)))));
        _1199_v10 = (_1199_v10).update(p0, p0);
        let _1200_v11;
        _1200_v11 = _module.D6.create_DC14((_1186_i0).isLessThan(new BigNumber(332)));
        let _1201_v12;
        let _nw190 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1201_v12 = _nw190;
        let _index211 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1201_v12).length));
        (_1201_v12)[_index211] = p2;
        let _1202_v13;
        let _nw191 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
        _1202_v13 = _nw191;
        let _index212 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1202_v13).length));
        (_1202_v13)[_index212] = _module.__default.fm43(!(p0), !((_this).f6), globalState);
        let _1203_v14;
        _1203_v14 = _dafny.Set.fromElements(p1, _this.f15, new BigNumber(439), new BigNumber((_1185_v0).length), p1);
        let _1204_v15;
        _1204_v15 = _dafny.Seq.of((_dafny.Set.fromElements(_this.f15)).Difference(_1203_v14));
        let _1205_v16;
        _1205_v16 = _dafny.MultiSet.fromElements((_this.f15).multipliedBy(_this.f15));
        let _1206_v17;
        _1206_v17 = _dafny.Map.Empty.slice().updateUnsafe((_1187_v1).fm7((_this).f6, globalState),!((_this).f6));
        let _index213 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1201_v12).length));
        let _index214 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1202_v13).length));
        let _rhs230 = _1200_v11;
        let _rhs231 = (((_1205_v16).contains((p1).multipliedBy(_this.f15))) ? ((_1205_v16).get((p1).multipliedBy(_this.f15))) : (_1186_i0));
        let _rhs232 = ((!(_1186_i0).isEqualTo(_1186_i0)) ? (p2) : (p2));
        let _rhs233 = (_1203_v14).Union((_1203_v14).Union(function () {
          let _coll40 = new _dafny.Set();
          for (const _compr_40 of (_1206_v17).Keys.Elements) {
            let _1207_v18 = _compr_40;
            if ((_1206_v17).contains(_1207_v18)) {
              _coll40.add(_module.__default.safeModuloInt(_1207_v18, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), function (_1208_i2) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              })).length)));
            }
          }
          return _coll40;
        }()));
        let _rhs234 = _1204_v15;
        let _lhs157 = globalState;
        let _lhs158 = _1201_v12;
        let _lhs159 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1201_v12).length));
        let _lhs160 = _1202_v13;
        let _lhs161 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1202_v13).length));
        _1200_v11 = _rhs230;
        _lhs157.f1 = _rhs231;
        _lhs158[_lhs159] = _rhs232;
        _lhs160[_lhs161] = _rhs233;
        _1204_v15 = _rhs234;
        let _index215 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1188_v2).length));
        (_1188_v2)[_index215] = p0;
        let _index216 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1188_v2).length));
        (_1188_v2)[_index216] = ((((_this).f6) || ((_this).f6)) ? (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-713)), function (_1209_i3) {
          return (_this).f5;
        }), (_1201_v12)[_module.__default.safeIndex(new BigNumber(153), new BigNumber((_1201_v12).length))])) : (!(!(p0) || (false))));
      }
      (_this).f15 = p1;
      let _1210_v19;
      let _nw192 = new _module.C3();
      _nw192.__ctor(p0, (_this).fm6(_1185_v0, globalState));
      _1210_v19 = _nw192;
      (_this).f15 = (_1185_v0)[_module.__default.safeIndex(_this.f15, new BigNumber((_1185_v0).length))];
      let _1211_v20;
      _1211_v20 = _dafny.MultiSet.fromElements((_1210_v19).f13, false, p0, (_1210_v19).f14, (_1210_v19).f14);
      r0 = (_dafny.MultiSet.fromElements((_1210_v19).f13)).IsProperSubsetOf((((_this).f6) ? (_1211_v20) : (_1211_v20)));
      let _1212_v21;
      _1212_v21 = _module.D0.create_DC1();
      let _1213_v22;
      _1213_v22 = _module.D0.create_DC3(_1212_v21);
      let _1214_v23;
      _1214_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1213_v22,(_dafny.ZERO).minus(p1));
      let _1215_v24;
      _1215_v24 = _dafny.Seq.of(_1214_v23, _1214_v23, _1214_v23, _1214_v23);
      let _1216_v25;
      let _nw193 = Array((new BigNumber(3)).toNumber());
      _nw193[(_dafny.ZERO).toNumber()] = _1215_v24;
      _nw193[(_dafny.ONE).toNumber()] = _1215_v24;
      _nw193[(new BigNumber(2)).toNumber()] = _1215_v24;
      _1216_v25 = _nw193;
      let _index217 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1216_v25).length));
      (_1216_v25)[_index217] = _dafny.Seq.of(_1214_v23);
      let _index218 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1216_v25).length));
      (_1216_v25)[_index218] = _dafny.Seq.Concat(_1215_v24, _1215_v24);
      r0 = (_this).f6;
      r1 = _dafny.Seq.Concat(_dafny.Seq.Concat(p2, _dafny.Seq.update(_dafny.Seq.update(p2, _module.__default.safeIndex(_this.f15, new BigNumber((p2).length)), (_this).f5), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(_this.f15, new BigNumber((p2).length)), (_this).f5)).length)), (_this).f5)), _dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), (_this).f5));
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _source18 = _module.D0.create_DC1();
      if (_source18.is_DC1) {
        let _1217_v0;
        _1217_v0 = false;
        let _1218_v1;
        _1218_v1 = _dafny.Seq.of((_this).f6, false);
        _1217_v0 = _dafny.Seq.IsPrefixOf(_1218_v1, _1218_v1);
        let _1219_v2;
        _1219_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this.f15).isLessThan(new BigNumber(455)));
        _1219_v2 = (_1219_v2).update((_this).f6, (_this).f6);
        let _1220_v3;
        let _nw194 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1220_v3 = _nw194;
        let _index219 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_1220_v3).length));
        (_1220_v3)[_index219] = (_1217_v0) || (!(false));
        let _1221_v4;
        _1221_v4 = _dafny.MultiSet.fromElements(_1217_v0, _1217_v0);
        let _1222_v5;
        _1222_v5 = _module.D15.create_DC27(_1221_v4);
        let _index220 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_1220_v3).length));
        let _rhs235 = ((_1221_v4).Union(_1221_v4)).Intersect((_1222_v5).dtor_cf35);
        let _rhs236 = !((_dafny.ZERO).minus(_this.f15)).isEqualTo(_this.f15);
        let _lhs162 = globalState;
        let _lhs163 = _1220_v3;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_1220_v3).length));
        _lhs162.f4 = _rhs235;
        _lhs163[_lhs164] = _rhs236;
        let _1223_v6;
        _1223_v6 = _dafny.MultiSet.fromElements(_this.f15, _this.f15);
        let _1224_v7;
        _1224_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wl")).length),_1223_v6);
        let _1225_v8;
        _1225_v8 = _dafny.Seq.of(new BigNumber(758));
        let _1226_v9;
        _1226_v9 = _dafny.Seq.of(_1223_v6);
        let _1227_v10;
        let _nw195 = Array((new BigNumber(14)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm44(!((_1220_v3)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_1220_v3).length))]), globalState)).cardinality()));
        _nw195[(_dafny.ONE).toNumber()] = _module.__default.fm45(globalState);
        _nw195[(new BigNumber(2)).toNumber()] = (((_1224_v7).contains(new BigNumber((_module.__default.fm46((_this).fm6(_1225_v8, globalState), globalState)).length))) ? ((_1224_v7).get(new BigNumber((_module.__default.fm46((_this).fm6(_1225_v8, globalState), globalState)).length))) : (_dafny.MultiSet.fromElements(_this.f15)));
        _nw195[(new BigNumber(3)).toNumber()] = (_1223_v6).Union((_1223_v6).update(_this.f15, _module.__default.abs(_this.f15)));
        _nw195[(new BigNumber(4)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(5)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(6)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_1225_v8).length));
        _nw195[(new BigNumber(8)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(9)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(10)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(11)).toNumber()] = _1223_v6;
        _nw195[(new BigNumber(12)).toNumber()] = (_1226_v9)[_module.__default.safeIndex(_this.f15, new BigNumber((_1226_v9).length))];
        _nw195[(new BigNumber(13)).toNumber()] = ((_1223_v6).update(_this.f15, _module.__default.abs(_this.f15))).Intersect(_1223_v6);
        _1227_v10 = _nw195;
        let _index221 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1227_v10).length));
        (_1227_v10)[_index221] = _1223_v6;
        let _index222 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1227_v10).length));
        (_1227_v10)[_index222] = _1223_v6;
      } else if (_source18.is_DC2) {
        let _1228___mcc_h0 = (_source18).cf1;
        let _1229___mcc_h1 = (_source18).cf2;
        let _1230_cf2 = _1229___mcc_h1;
        let _1231_cf1 = _1228___mcc_h0;
        let _1232_v11;
        _1232_v11 = false;
        _1232_v11 = (_this).f6;
        r0 = (((((_this).f6) ? ((_this).f6) : (false))) ? (_this.f15) : (_this.f15));
        let _1233_v12;
        _1233_v12 = _dafny.MultiSet.fromElements(_1232_v11);
        let _source19 = _module.D15.create_DC27(_1233_v12);
        if (_source19.is_DC28) {
          let _1234___mcc_h4 = (_source19).cf36;
          let _1235___mcc_h5 = (_source19).cf37;
          let _1236_cf37 = _1235___mcc_h5;
          let _1237_cf36 = _1234___mcc_h4;
          (globalState).f1 = (_this).fm7((_this).f6, globalState);
          let _1238_v13;
          let _nw196 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1238_v13 = _nw196;
          let _1239_v14;
          _1239_v14 = _dafny.Seq.UnicodeFromString("g");
          let _1240_v15;
          _1240_v15 = _dafny.MultiSet.fromElements(_1239_v14, _1239_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(969)), ((_1241_cf2) => function (_1242_i0) {
            return _1241_cf2;
          })(_1230_cf2)), _dafny.Seq.UnicodeFromString("bjrtxt"), _1239_v14);
          let _1243_v17;
          _1243_v17 = function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(717), new BigNumber(-196))) {
              let _1244_v16 = _compr_41;
              if (((new BigNumber(717)).isLessThanOrEqualTo(_1244_v16)) && ((_1244_v16).isLessThan(new BigNumber(-196)))) {
                _coll41.push([_module.__default.safeDivisionInt(_1244_v16, new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, (_this).f6))).cardinality())),new BigNumber(868)]);
              }
            }
            return _coll41;
          }();
          let _1245_v18;
          _1245_v18 = _module.D12.create_DC22(_1243_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), ((_1246_v11, _1247_cf37) => function (_1248_i1) {
  return _dafny.Seq.of(_1246_v11, _1247_cf37);
})(_1232_v11, _1236_cf37)), _1230_cf2, _1232_v11, _1239_v14);
          let _nw197 = Array((new BigNumber(19)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = (_this).f6;
          _nw197[(_dafny.ONE).toNumber()] = (_1240_v15).IsSubsetOf(_1240_v15);
          _nw197[(new BigNumber(2)).toNumber()] = _1236_cf37;
          _nw197[(new BigNumber(3)).toNumber()] = ((_1232_v11) ? (!(_1232_v11)) : (_1237_cf36));
          _nw197[(new BigNumber(4)).toNumber()] = _1236_cf37;
          _nw197[(new BigNumber(5)).toNumber()] = false;
          _nw197[(new BigNumber(6)).toNumber()] = _1237_cf36;
          _nw197[(new BigNumber(7)).toNumber()] = _1237_cf36;
          _nw197[(new BigNumber(8)).toNumber()] = _1236_cf37;
          _nw197[(new BigNumber(9)).toNumber()] = _1232_v11;
          _nw197[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1239_v14, _dafny.Seq.update(_1239_v14, _module.__default.safeIndex(_this.f15, new BigNumber((_1239_v14).length)), (_1245_v18).dtor_cf26));
          _nw197[(new BigNumber(11)).toNumber()] = (_this.f15).isLessThan(new BigNumber(440));
          _nw197[(new BigNumber(12)).toNumber()] = _1232_v11;
          _nw197[(new BigNumber(13)).toNumber()] = _1232_v11;
          _nw197[(new BigNumber(14)).toNumber()] = _1236_cf37;
          _nw197[(new BigNumber(15)).toNumber()] = _1237_cf36;
          _nw197[(new BigNumber(16)).toNumber()] = _1237_cf36;
          _nw197[(new BigNumber(17)).toNumber()] = !(!(_1236_cf37));
          _nw197[(new BigNumber(18)).toNumber()] = _1232_v11;
          _1238_v13 = _nw197;
          (globalState).f1 = (_this.f15).multipliedBy(_this.f15);
          let _1249_v19;
          _1249_v19 = _module.D11.create_DC20((((_this).f6) ? (_1236_cf37) : (_1232_v11)), (_this.f15).minus(_this.f15));
          _1249_v19 = _1249_v19;
        } else if (_source19.is_DC27) {
          let _1250___mcc_h6 = (_source19).cf35;
          let _1251_cf35 = _1250___mcc_h6;
          let _1252_v20;
          _1252_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
          let _1253_v21;
          _1253_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1252_v20);
          let _1254_v22;
          _1254_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1232_v11,_module.__default.fm47((_this).f6, globalState));
          (globalState).f1 = new BigNumber(((_1253_v21).Merge((((_1254_v22).contains((_this).f6)) ? ((_1254_v22).get((_this).f6)) : (_1253_v21)))).length);
          (_this).f15 = _this.f15;
          let _1255_v23;
          let _nw198 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1255_v23 = _nw198;
          let _1256_v24;
          _1256_v24 = _1255_v23;
          let _1257_v25;
          let _nw199 = Array((new BigNumber(28)).toNumber());
          _nw199[(_dafny.ZERO).toNumber()] = _1255_v23;
          _nw199[(_dafny.ONE).toNumber()] = (_1256_v24);
          _nw199[(new BigNumber(2)).toNumber()] = ((_1232_v11) ? (_1255_v23) : (_1255_v23));
          _nw199[(new BigNumber(3)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(4)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(5)).toNumber()] = (_1256_v24);
          _nw199[(new BigNumber(6)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(7)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(8)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(9)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(10)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(11)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(12)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(13)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(14)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(15)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(16)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(17)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(18)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(19)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(20)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(21)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(22)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(23)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(24)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(25)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(26)).toNumber()] = _1255_v23;
          _nw199[(new BigNumber(27)).toNumber()] = _1255_v23;
          _1257_v25 = _nw199;
          let _index223 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1257_v25).length));
          (_1257_v25)[_index223] = ((_1232_v11) ? (_1255_v23) : (_1255_v23));
          let _index224 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1257_v25).length));
          (_1257_v25)[_index224] = _1255_v23;
          (globalState).f1 = (((_this).f6) ? (_this.f15) : (_this.f15));
        } else {
          let _1258___mcc_h7 = (_source19).cf38;
          let _1259_cf38 = _1258___mcc_h7;
          _1232_v11 = (_this).f6;
          let _1260_v26;
          _1260_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_1233_v12).cardinality()));
          let _1261_v27;
          let _init35 = function (_1262_i2) {
            return (_1262_i2).plus(_this.f15);
          };
          let _nw200 = Array((new BigNumber(15)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw200.length); _i0_35++) {
            _nw200[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1261_v27 = _nw200;
          let _1263_v28;
          _1263_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_1232_v11,(_dafny.ZERO).minus(_this.f15))).Merge(_1260_v26),_1261_v27);
          let _1264_v29;
          _1264_v29 = _dafny.Set.fromElements(_1260_v26, _1260_v26, (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f15)).Merge(_1260_v26), _1260_v26, (_1260_v26).Merge(_1260_v26));
          let _index225 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1261_v27).length));
          (_1261_v27)[_index225] = _this.f15;
          let _index226 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1261_v27).length));
          let _rhs237 = _this.f15;
          let _rhs238 = (_dafny.Map.Empty.slice().updateUnsafe(_1260_v26,_1261_v27)).Merge(_1263_v28);
          let _rhs239 = _1264_v29;
          let _rhs240 = _this.f15;
          let _rhs241 = new BigNumber(459);
          let _lhs165 = globalState;
          let _lhs166 = _1261_v27;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1261_v27).length));
          let _lhs168 = globalState;
          _lhs165.f1 = _rhs237;
          _1263_v28 = _rhs238;
          _1264_v29 = _rhs239;
          _lhs166[_lhs167] = _rhs240;
          _lhs168.f1 = _rhs241;
          let _1265_v30;
          let _init36 = function (_1266_i3) {
            return (_this).f6;
          };
          let _nw201 = Array((new BigNumber(3)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw201.length); _i0_36++) {
            _nw201[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1265_v30 = _nw201;
          let _1267_v31;
          let _1268_v32;
          let _out34;
          let _out35;
          let _outcollector16 = _module.__default.m0(_1265_v30, (_this).f5, !((_this).f6), _1230_cf2, globalState);
          _out34 = _outcollector16[0];
          _out35 = _outcollector16[1];
          _1267_v31 = _out34;
          _1268_v32 = _out35;
          let _1269_v33;
          _1269_v33 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_module.__default.fm48((_1261_v27)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_1261_v27).length))], true, globalState)).cardinality())),_1268_v32);
          let _1270_v34;
          _1270_v34 = _1269_v33;
          let _1271_v35;
          _1271_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v32,_1270_v34);
          _1271_v35 = (_1271_v35).update(((_1268_v32) ? ((_this).f6) : (_1232_v11)), _1270_v34);
        }
        let _1272_v36;
        _1272_v36 = _dafny.Seq.of(_this.f15);
        let _1273_v37;
        _1273_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC26(_1232_v11, new BigNumber((_1272_v36).length), _1230_cf2),(_this).f5);
        _1232_v11 = !(_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1273_v37)).equals(function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of (_1272_v36).Elements) {
            let _1274_v38 = _compr_42;
            if (_dafny.Seq.contains(_1272_v36, _1274_v38)) {
              _coll42.push([(_1274_v38).plus(_this.f15),(function () {
                let _coll43 = new _dafny.Map();
                for (const _compr_43 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-187)), ((_1275_v11, _1276_cf2) => function (_1277_i4) {
                  return _module.D14.create_DC26(_1275_v11, _this.f15, _1276_cf2);
                })(_1232_v11, _1230_cf2))).Elements) {
                  let _1278_v39 = _compr_43;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-187)), ((_1279_v11, _1280_cf2) => function (_1277_i4) {
                    return _module.D14.create_DC26(_1279_v11, _this.f15, _1280_cf2);
                  })(_1232_v11, _1230_cf2)), _1278_v39)) {
                    _coll43.push([_1278_v39,new _dafny.CodePoint('q'.codePointAt(0))]);
                  }
                }
                return _coll43;
              }()).update(_module.D14.create_DC26(_1232_v11, _this.f15, (_this).f5), _1230_cf2)]);
            }
          }
          return _coll42;
        }());
      } else if (_source18.is_DC0) {
        let _1281___mcc_h2 = (_source18).cf0;
        let _1282_cf0 = _1281___mcc_h2;
        let _1283_v40;
        let _init37 = function (_1284_i5) {
          return (_1284_i5).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),(_module.D6.create_DC14((_this).f6)).dtor_cf15)).length));
        };
        let _nw202 = Array((new BigNumber(27)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw202.length); _i0_37++) {
          _nw202[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1283_v40 = _nw202;
        let _1285_v41;
        _1285_v41 = _dafny.Seq.UnicodeFromString("epvurt");
        let _1286_v42;
        _1286_v42 = _dafny.Seq.of(new BigNumber((_1285_v41).length));
        let _1287_v43;
        _1287_v43 = _dafny.Seq.of(new BigNumber((_1286_v42).length));
        let _1288_v44;
        let _nw203 = Array((new BigNumber(3)).toNumber());
        _nw203[(_dafny.ZERO).toNumber()] = _1282_cf0;
        _nw203[(_dafny.ONE).toNumber()] = (_this).fm6(_1287_v43, globalState);
        _nw203[(new BigNumber(2)).toNumber()] = _1282_cf0;
        _1288_v44 = _nw203;
        let _1289_v45;
        _1289_v45 = _dafny.Set.fromElements(new BigNumber(797), new BigNumber(526), new BigNumber((_module.__default.fm49(globalState)).length));
        let _index227 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1288_v44).length));
        (_1288_v44)[_index227] = (_1289_v45).contains(new BigNumber((_1289_v45).length));
        let _1290_v46;
        _1290_v46 = _dafny.Seq.of(_1282_cf0, (_this).f6);
        let _1291_v47;
        _1291_v47 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_1290_v46).length), new BigNumber(53), _this.f15, _this.f15));
        let _1292_v48;
        _1292_v48 = _module.D6.create_DC13(_1286_v42);
        let _1293_v49;
        _1293_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1289_v45).length),_1282_cf0);
        let _1294_v50;
        let _nw204 = Array((new BigNumber(22)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1286_v42, _module.__default.safeIndex(new BigNumber(-560), new BigNumber((_1286_v42).length)), new BigNumber(501)), _1286_v42);
        _nw204[(_dafny.ONE).toNumber()] = _1286_v42;
        _nw204[(new BigNumber(2)).toNumber()] = _1287_v43;
        _nw204[(new BigNumber(3)).toNumber()] = (_1291_v47)[_module.__default.safeIndex(_this.f15, new BigNumber((_1291_v47).length))];
        _nw204[(new BigNumber(4)).toNumber()] = _1287_v43;
        _nw204[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_1295_i6) {
          return new BigNumber(189);
        });
        _nw204[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_this.f15, new BigNumber((_1285_v41).length), _this.f15);
        _nw204[(new BigNumber(7)).toNumber()] = _1286_v42;
        _nw204[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1287_v43, _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1287_v43).length)), _this.f15);
        _nw204[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f15, _this.f15), (_1292_v48).dtor_cf14);
        _nw204[(new BigNumber(10)).toNumber()] = _1286_v42;
        _nw204[(new BigNumber(11)).toNumber()] = _1287_v43;
        _nw204[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1293_v49).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-2)), function (_1296_i7) {
          return new BigNumber(-85);
        }));
        _nw204[(new BigNumber(13)).toNumber()] = _1287_v43;
        _nw204[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1286_v42, _1286_v42);
        _nw204[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_1286_v42, _module.__default.safeIndex((_dafny.ZERO).minus(_this.f15), new BigNumber((_1286_v42).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1282_cf0,_this.f15)).length));
        _nw204[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), function (_1297_i8) {
          return _this.f15;
        }), _1286_v42), _module.__default.safeIndex((_dafny.ZERO).minus(_this.f15), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), function (_1298_i8) {
          return _this.f15;
        }), _1286_v42)).length)), _this.f15);
        _nw204[(new BigNumber(17)).toNumber()] = _1287_v43;
        _nw204[(new BigNumber(18)).toNumber()] = _1286_v42;
        _nw204[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1287_v43, _1287_v43);
        _nw204[(new BigNumber(20)).toNumber()] = _1286_v42;
        _nw204[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), function (_1299_i9) {
          return _this.f15;
        });
        _1294_v50 = _nw204;
        let _index228 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1294_v50).length));
        (_1294_v50)[_index228] = _dafny.Seq.of(_this.f15, _this.f15);
        let _1300_v51;
        _1300_v51 = _dafny.MultiSet.fromElements((_this).f6);
        let _1301_v52;
        _1301_v52 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1300_v51);
        let _1302_v53;
        _1302_v53 = _dafny.Seq.of(_1300_v51);
        let _1303_v54;
        let _nw205 = Array((new BigNumber(25)).toNumber());
        _nw205[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.FromArray(_1290_v46)).Union((_1300_v51).update(true, _module.__default.abs(_this.f15)));
        _nw205[(_dafny.ONE).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(2)).toNumber()] = (_1300_v51).Difference(_1300_v51);
        _nw205[(new BigNumber(3)).toNumber()] = (_1300_v51).Intersect(_1300_v51);
        _nw205[(new BigNumber(4)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(5)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(6)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(7)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(8)).toNumber()] = (_1300_v51).Difference(_1300_v51);
        _nw205[(new BigNumber(9)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(10)).toNumber()] = (_1300_v51).Intersect(_module.__default.fm48(_this.f15, (_this).f6, globalState));
        _nw205[(new BigNumber(11)).toNumber()] = (_1300_v51).Union(_1300_v51);
        _nw205[(new BigNumber(12)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(true))));
        _nw205[(new BigNumber(14)).toNumber()] = (_1300_v51).Union(_1300_v51);
        _nw205[(new BigNumber(15)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(16)).toNumber()] = (((_1301_v52).contains((_this).fm7(_1282_cf0, globalState))) ? ((_1301_v52).get((_this).fm7(_1282_cf0, globalState))) : (_1300_v51));
        _nw205[(new BigNumber(17)).toNumber()] = (_1302_v53)[_module.__default.safeIndex(_this.f15, new BigNumber((_1302_v53).length))];
        _nw205[(new BigNumber(18)).toNumber()] = (_1300_v51).Difference(_1300_v51);
        _nw205[(new BigNumber(19)).toNumber()] = (_dafny.MultiSet.FromArray(_1290_v46)).update((_1290_v46)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_1290_v46).length))], _module.__default.abs(_this.f15));
        _nw205[(new BigNumber(20)).toNumber()] = (_dafny.MultiSet.fromElements(_1282_cf0, _1282_cf0)).Union(_1300_v51);
        _nw205[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.fromElements(_1282_cf0);
        _nw205[(new BigNumber(22)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(23)).toNumber()] = _1300_v51;
        _nw205[(new BigNumber(24)).toNumber()] = (_1300_v51).Difference(_1300_v51);
        _1303_v54 = _nw205;
        let _1304_v55;
        let _nw206 = Array((new BigNumber(21)).toNumber()).fill([]);
        _1304_v55 = _nw206;
        let _1305_v57;
        _1305_v57 = _dafny.MultiSet.fromElements(new BigNumber((_1285_v41).length));
        let _1306_v59;
        let _nw207 = Array((new BigNumber(28)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = _1293_v49;
        _nw207[(_dafny.ONE).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(2)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(3)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(4)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(5)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(6)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(7)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(8)).toNumber()] = function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of (_1305_v57).Elements) {
            let _1307_v56 = _compr_44;
            if ((_1305_v57).contains(_1307_v56)) {
              _coll44.push([(_1307_v56).plus(new BigNumber((_1290_v46).length)),_1282_cf0]);
            }
          }
          return _coll44;
        }();
        _nw207[(new BigNumber(9)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(10)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(11)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(12)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(13)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(14)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(15)).toNumber()] = function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-991), new BigNumber(889))) {
            let _1308_v58 = _compr_45;
            if (((new BigNumber(-991)).isLessThanOrEqualTo(_1308_v58)) && ((_1308_v58).isLessThan(new BigNumber(889)))) {
              _coll45.push([(_1308_v58).plus(_this.f15),(_module.D14.create_DC26(false, _this.f15, (_this).f5)).dtor_cf32]);
            }
          }
          return _coll45;
        }();
        _nw207[(new BigNumber(16)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(17)).toNumber()] = (_1293_v49).update(_this.f15, (_this).f6);
        _nw207[(new BigNumber(18)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(19)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(20)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(21)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(22)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(23)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(24)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,(_this).f6);
        _nw207[(new BigNumber(26)).toNumber()] = _1293_v49;
        _nw207[(new BigNumber(27)).toNumber()] = _1293_v49;
        _1306_v59 = _nw207;
        let _index229 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_1304_v55).length));
        (_1304_v55)[_index229] = _1306_v59;
        let _index230 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1288_v44).length));
        let _index231 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1294_v50).length));
        let _index232 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_1304_v55).length));
        let _rhs242 = _1283_v40;
        let _rhs243 = !((_1289_v45).IsSubsetOf(((!(!(true))) ? (_1289_v45) : (_1289_v45))));
        let _rhs244 = _dafny.Seq.update(_1287_v43, _module.__default.safeIndex(_this.f15, new BigNumber((_1287_v43).length)), new BigNumber(777));
        let _rhs245 = _1303_v54;
        let _rhs246 = _1306_v59;
        let _lhs169 = _1288_v44;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1288_v44).length));
        let _lhs171 = _1294_v50;
        let _lhs172 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_1294_v50).length));
        let _lhs173 = _1304_v55;
        let _lhs174 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_1304_v55).length));
        _1283_v40 = _rhs242;
        _lhs169[_lhs170] = _rhs243;
        _lhs171[_lhs172] = _rhs244;
        _1303_v54 = _rhs245;
        _lhs173[_lhs174] = _rhs246;
        let _1309_v60;
        _1309_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("erhqhi"),_1285_v41);
        _1309_v60 = _1309_v60;
        let _1310_v61;
        _1310_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1288_v44)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_1288_v44).length))],_1283_v40);
        let _1311_v62;
        _1311_v62 = ((!(false)) ? (_1283_v40) : ((((_1310_v61).contains((_1290_v46)[_module.__default.safeIndex(_this.f15, new BigNumber((_1290_v46).length))])) ? ((_1310_v61).get((_1290_v46)[_module.__default.safeIndex(_this.f15, new BigNumber((_1290_v46).length))])) : (_1283_v40))));
        let _rhs247 = !(_1289_v45).equals(_1289_v45);
        let _rhs248 = _1289_v45;
        let _rhs249 = _1311_v62;
        _1282_cf0 = _rhs247;
        _1289_v45 = _rhs248;
        _1311_v62 = _rhs249;
        (_this).f15 = new BigNumber(75);
      } else {
        let _1312___mcc_h3 = (_source18).cf3;
        let _1313_cf3 = _1312___mcc_h3;
        if ((_this).f6) {
          let _1314_v63;
          _1314_v63 = true;
          let _1315_v64;
          _1315_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f15);
          let _1316_v65;
          _1316_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v63,(_this).f5);
          _1314_v63 = (_dafny.Map.Empty.slice().updateUnsafe(_this.f15,(((_1315_v64).contains((((_1316_v65).contains(_1314_v63)) ? ((_1316_v65).get(_1314_v63)) : ((_this).f5)))) ? ((_1315_v64).get((((_1316_v65).contains(_1314_v63)) ? ((_1316_v65).get(_1314_v63)) : ((_this).f5)))) : (_this.f15)))).equals(function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(938), new BigNumber(755))) {
              let _1317_v66 = _compr_46;
              if (((new BigNumber(938)).isLessThanOrEqualTo(_1317_v66)) && ((_1317_v66).isLessThan(new BigNumber(755)))) {
                _coll46.push([(_1317_v66).multipliedBy(new BigNumber(198)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(998),function () {
                  let _coll47 = new _dafny.Map();
                  for (const _compr_47 of (_dafny.Set.fromElements((_this).f5, (_this).f5, (_this).f5)).Elements) {
                    let _1318_v67 = _compr_47;
                    if ((_dafny.Set.fromElements((_this).f5, (_this).f5, (_this).f5)).contains(_1318_v67)) {
                      _coll47.push([_1318_v67,false]);
                    }
                  }
                  return _coll47;
                }())).length)]);
              }
            }
            return _coll46;
          }());
          _1314_v63 = _1314_v63;
          let _1319_v68;
          let _init38 = function (_1320_i10) {
            return _dafny.Seq.UnicodeFromString("tvpi");
          };
          let _nw208 = Array((new BigNumber(8)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw208.length); _i0_38++) {
            _nw208[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1319_v68 = _nw208;
          let _1321_v69;
          _1321_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v63,_dafny.Seq.UnicodeFromString("yfuflx"));
          let _1322_v70;
          _1322_v70 = _dafny.Seq.UnicodeFromString("bc");
          let _index233 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1319_v68).length));
          (_1319_v68)[_index233] = _dafny.Seq.Concat((((_1321_v69).contains((_this).f6)) ? ((_1321_v69).get((_this).f6)) : (_1322_v70)), _1322_v70);
          let _index234 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1319_v68).length));
          (_1319_v68)[_index234] = _1322_v70;
          let _1323_v71;
          let _nw209 = Array((new BigNumber(13)).toNumber()).fill(false);
          _1323_v71 = _nw209;
          let _index235 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1323_v71).length));
          (_1323_v71)[_index235] = !(_1314_v63);
          let _index236 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_1323_v71).length));
          (_1323_v71)[_index236] = _1314_v63;
          let _1324_v72;
          _1324_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),(_this).f6);
          let _1325_v73;
          _1325_v73 = _1324_v72;
          let _1326_v74;
          _1326_v74 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1325_v73),_this.f15);
          (globalState).f1 = (((_1326_v74).contains(_module.__default.fm50(_module.__default.fm0(_1314_v63, globalState), (_this).f6, function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of _dafny.IntegerRange(new BigNumber(477), new BigNumber(-150))) {
              let _1328_v75 = _compr_49;
              if (((new BigNumber(477)).isLessThanOrEqualTo(_1328_v75)) && ((_1328_v75).isLessThan(new BigNumber(-150)))) {
                _coll49.push([(_1328_v75).multipliedBy(new BigNumber(((_1319_v68)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1319_v68).length))]).length)),_this.f15]);
              }
            }
            return _coll49;
          }(), globalState))) ? ((_1326_v74).get(_module.__default.fm50(_module.__default.fm0(_1314_v63, globalState), (_this).f6, function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of _dafny.IntegerRange(new BigNumber(477), new BigNumber(-150))) {
              let _1327_v75 = _compr_48;
              if (((new BigNumber(477)).isLessThanOrEqualTo(_1327_v75)) && ((_1327_v75).isLessThan(new BigNumber(-150)))) {
                _coll48.push([(_1327_v75).multipliedBy(new BigNumber(((_1319_v68)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1319_v68).length))]).length)),_this.f15]);
              }
            }
            return _coll48;
          }(), globalState))) : (_this.f15));
        } else {
          let _1329_v76;
          _1329_v76 = false;
          let _1330_v77;
          _1330_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1329_v76);
          _1329_v76 = ((new BigNumber((_1330_v77).length)).minus(_this.f15)).isEqualTo(_this.f15);
          (globalState).f1 = (_module.__default.fm3((_this).f6, true, globalState)).plus(_module.__default.safeModuloInt(_this.f15, _this.f15));
          let _1331_v78;
          _1331_v78 = _dafny.Seq.UnicodeFromString("paf");
          let _1332_v79;
          _1332_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f15);
          let _1333_v80;
          _1333_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1332_v79).length),_this.f15);
          let _1334_v81;
          _1334_v81 = _dafny.Seq.of(_this.f15, _this.f15, new BigNumber((_1333_v80).length));
          let _1335_v82;
          let _nw210 = Array((new BigNumber(25)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_this.f15);
          _nw210[(_dafny.ONE).toNumber()] = _this.f15;
          _nw210[(new BigNumber(2)).toNumber()] = (_this.f15).plus(new BigNumber(462));
          _nw210[(new BigNumber(3)).toNumber()] = new BigNumber((_1331_v78).length);
          _nw210[(new BigNumber(4)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(5)).toNumber()] = (_this.f15).plus(new BigNumber((_1331_v78).length));
          _nw210[(new BigNumber(6)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(7)).toNumber()] = ((true) ? (new BigNumber(159)) : (new BigNumber((_1331_v78).length)));
          _nw210[(new BigNumber(8)).toNumber()] = new BigNumber(72);
          _nw210[(new BigNumber(9)).toNumber()] = (_1334_v81)[_module.__default.safeIndex(_this.f15, new BigNumber((_1334_v81).length))];
          _nw210[(new BigNumber(10)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(11)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(12)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(13)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(14)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(15)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(16)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(17)).toNumber()] = (_this.f15).multipliedBy(_this.f15);
          _nw210[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_this.f15);
          _nw210[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(_this.f15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_this.f15)).length));
          _nw210[(new BigNumber(20)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(21)).toNumber()] = (_1334_v81)[_module.__default.safeIndex(_this.f15, new BigNumber((_1334_v81).length))];
          _nw210[(new BigNumber(22)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(23)).toNumber()] = _this.f15;
          _nw210[(new BigNumber(24)).toNumber()] = _this.f15;
          _1335_v82 = _nw210;
          let _index237 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1335_v82).length));
          (_1335_v82)[_index237] = _module.__default.safeDivisionInt(_this.f15, (_this).fm8(true, (_this).f6, (_this).f6, globalState));
          let _index238 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1335_v82).length));
          (_1335_v82)[_index238] = new BigNumber(149);
          let _index239 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1335_v82).length));
          (_1335_v82)[_index239] = (_1335_v82)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1335_v82).length))];
          _1331_v78 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_1336_i11) {
            return (_this).f5;
          }), _dafny.Seq.update(_1331_v78, _module.__default.safeIndex((_1335_v82)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1335_v82).length))], new BigNumber((_1331_v78).length)), (_this).f5));
        }
        if ((_this).f6) {
          (globalState).f1 = (new BigNumber(355)).multipliedBy(_this.f15);
          let _1337_v83;
          _1337_v83 = false;
          _1337_v83 = _1337_v83;
          let _1338_v84;
          _1338_v84 = _dafny.Seq.of(_this.f15, _this.f15, new BigNumber(904), _this.f15);
          let _1339_v85;
          _1339_v85 = _dafny.Seq.of(_1338_v84);
          let _1340_v86;
          let _nw211 = new _module.C0();
          _nw211.__ctor(_1337_v83, _1339_v85);
          _1340_v86 = _nw211;
          let _1341_v87;
          _1341_v87 = _dafny.Seq.of((_module.D14.create_DC26(_1337_v83, _this.f15, (_this).f5)).dtor_cf32, (_this).f6, false);
          let _1342_v88;
          _1342_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_1340_v86).f11, _1337_v83),(_this).f6);
          let _1343_v89;
          _1343_v89 = _dafny.Set.fromElements((_this).fm6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_1344_i12) {
            return _this.f15;
          }), globalState));
          let _1345_v90;
          _1345_v90 = _dafny.Seq.UnicodeFromString("jdvhobhq");
          let _1346_v91;
          _1346_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8(_1337_v83, (_1340_v86).f11, (_1340_v86).f11, globalState),(_1340_v86).f11);
          let _1347_v92;
          _1347_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1338_v84).length),_1345_v90);
          let _1348_v93;
          let _nw212 = Array((new BigNumber(26)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = (_this).f6;
          _nw212[(_dafny.ONE).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(2)).toNumber()] = _1337_v83;
          _nw212[(new BigNumber(3)).toNumber()] = (_this.f15).isLessThanOrEqualTo(_this.f15);
          _nw212[(new BigNumber(4)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(5)).toNumber()] = (_module.__default.fm0(_1337_v83, globalState)) || ((_1340_v86).f11);
          _nw212[(new BigNumber(6)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(7)).toNumber()] = (_1341_v87)[_module.__default.safeIndex(_this.f15, new BigNumber((_1341_v87).length))];
          _nw212[(new BigNumber(8)).toNumber()] = !((_this).f6) || ((_1340_v86).f11);
          _nw212[(new BigNumber(9)).toNumber()] = false;
          _nw212[(new BigNumber(10)).toNumber()] = !((_1340_v86).f11);
          _nw212[(new BigNumber(11)).toNumber()] = (_this.f15).isLessThan(_this.f15);
          _nw212[(new BigNumber(12)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(13)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(14)).toNumber()] = (((_1342_v88).contains(_1343_v89)) ? ((_1342_v88).get(_1343_v89)) : (_1337_v83));
          _nw212[(new BigNumber(15)).toNumber()] = _dafny.Seq.contains(_1345_v90, (_this).f5);
          _nw212[(new BigNumber(16)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(17)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(18)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(19)).toNumber()] = (_this).f6;
          _nw212[(new BigNumber(20)).toNumber()] = (((_1346_v91).contains(new BigNumber((_1347_v92).length))) ? ((_1346_v91).get(new BigNumber((_1347_v92).length))) : (_1337_v83));
          _nw212[(new BigNumber(21)).toNumber()] = (_this).f6;
          _nw212[(new BigNumber(22)).toNumber()] = ((_1340_v86).f11) && (_1337_v83);
          _nw212[(new BigNumber(23)).toNumber()] = (_this).f6;
          _nw212[(new BigNumber(24)).toNumber()] = (_1340_v86).f11;
          _nw212[(new BigNumber(25)).toNumber()] = true;
          _1348_v93 = _nw212;
          let _1349_v94;
          _1349_v94 = _module.D2.create_DC8(_1337_v83);
          let _rhs250 = _this.f15;
          let _rhs251 = (((_dafny.MultiSet.fromElements(_1337_v83)).update((_1340_v86).f11, _module.__default.abs(_this.f15))).update((_1340_v86).f11, _module.__default.abs(_this.f15))).update(_1337_v83, _module.__default.abs(_this.f15));
          let _rhs252 = ((!((_this).f6)) ? (_1348_v93) : (_1348_v93));
          let _rhs253 = _1349_v94;
          let _rhs254 = _module.__default.safeModuloInt(new BigNumber(539), _module.__default.safeModuloInt(new BigNumber((_module.__default.fm46((_1340_v86).f11, globalState)).length), _this.f15));
          let _lhs175 = _this;
          let _lhs176 = globalState;
          _lhs175.f15 = _rhs250;
          _lhs176.f4 = _rhs251;
          _1348_v93 = _rhs252;
          _1349_v94 = _rhs253;
          r0 = _rhs254;
          let _1350_v95;
          let _nw213 = new _module.C3();
          _nw213.__ctor((_module.__default.fm51(_1337_v83, new BigNumber((_1346_v91).length), (_this).f6, globalState)).IsSubsetOf(_dafny.Set.fromElements((_this).f5)), (_this).f6);
          _1350_v95 = _nw213;
        } else {
          let _1351_v96;
          _1351_v96 = _dafny.MultiSet.fromElements((_this).f6);
          let _1352_v97;
          let _nw214 = new _module.C3();
          _nw214.__ctor((_1351_v96).IsSubsetOf(_1351_v96), !((_this).f6));
          _1352_v97 = _nw214;
          let _1353_v98;
          _1353_v98 = _dafny.Seq.of((_this).f5);
          let _1354_v99;
          _1354_v99 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of((_this).f5, (_this).f5, (_this).f5, (_this).f5), _1353_v98),_this.f15);
          _1354_v99 = (_1354_v99).update(_1353_v98, _this.f15);
          (_this).f15 = (_this.f15).plus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1353_v98).length),(_1352_v97).f14)).update(new BigNumber((_dafny.Seq.UnicodeFromString("aknri")).length), (_1352_v97).f13)).length));
          let _1355_v100;
          let _init39 = function (_1356_i13) {
            return _module.__default.safeDivisionInt(_1356_i13, _this.f15);
          };
          let _nw215 = Array((new BigNumber(4)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw215.length); _i0_39++) {
            _nw215[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1355_v100 = _nw215;
          let _index240 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1355_v100).length));
          (_1355_v100)[_index240] = _this.f15;
          let _index241 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1355_v100).length));
          (_1355_v100)[_index241] = _this.f15;
          let _1357_v101;
          _1357_v101 = _module.D15.create_DC27(_1351_v96);
          _1357_v101 = _1357_v101;
        }
        let _1358_v102;
        let _1359_v103;
        let _1360_v104;
        let _1361_v105;
        let _out36;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector17 = (_this).m12(_this.f15, (_this).f6, _this.f15, globalState);
        _out36 = _outcollector17[0];
        _out37 = _outcollector17[1];
        _out38 = _outcollector17[2];
        _out39 = _outcollector17[3];
        _1358_v102 = _out36;
        _1359_v103 = _out37;
        _1360_v104 = _out38;
        _1361_v105 = _out39;
        let _1362_v106;
        let _nw216 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1362_v106 = _nw216;
        let _index242 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1362_v106).length));
        (_1362_v106)[_index242] = _this.f15;
        let _1363_v107;
        _1363_v107 = _dafny.Seq.of((_this).f6, (_this).f6, _1361_v105, _1361_v105, (_this).f6);
        let _1364_v108;
        _1364_v108 = _dafny.MultiSet.fromElements(new BigNumber((_1359_v103).length), _this.f15);
        let _1365_v109;
        _1365_v109 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f15), (_this).fm8(false, (_this).f6, _1361_v105, globalState), (_this).fm4(_1363_v107, _this.f15, _this.f15, _1364_v108, globalState), _this.f15, _this.f15);
        let _1366_v110;
        _1366_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f15);
        let _1367_v111;
        _1367_v111 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f15, (((_1364_v108).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15)).length))) ? ((_1364_v108).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15)).length))) : (_this.f15))),_1366_v110));
        let _index243 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1362_v106).length));
        (_1362_v106)[_index243] = (_this.f15).plus(new BigNumber((((_1367_v111)[_module.__default.safeIndex(_this.f15, new BigNumber((_1367_v111).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1359_v103,_1366_v110))).length));
      }
      (_this).f15 = ((_this).fm7((_this).f6, globalState)).minus(_this.f15);
      let _1368_v112;
      _1368_v112 = _dafny.MultiSet.fromElements(_this.f15, new BigNumber(521));
      let _hi7 = _module.__default.safeModuloInt(_this.f15, _this.f15);
      for (let _1369_i14 = (((_1368_v112).contains(_this.f15)) ? ((_1368_v112).get(_this.f15)) : (_this.f15)); _1369_i14.isLessThan(_hi7); _1369_i14 = _1369_i14.plus(_dafny.ONE)) {
        let _1370_v113;
        let _init40 = function (_1371_i15) {
          return (_1371_i15).multipliedBy(_this.f15);
        };
        let _nw217 = Array((new BigNumber(11)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw217.length); _i0_40++) {
          _nw217[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1370_v113 = _nw217;
        let _index244 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1370_v113).length));
        (_1370_v113)[_index244] = new BigNumber((_dafny.Seq.UnicodeFromString("jwrbku")).length);
        let _1372_v114;
        _1372_v114 = _dafny.Set.fromElements(_1369_i14);
        let _1373_v115;
        _1373_v115 = _dafny.Seq.of((_this).f6, (_this).f6, !((_this).f6) || ((_this).f6), (new BigNumber((_1372_v114).length)).isLessThanOrEqualTo(_this.f15));
        let _1374_v116;
        _1374_v116 = _dafny.MultiSet.fromElements((_this).f6);
        let _1375_v117;
        _1375_v117 = _module.D15.create_DC27(_1374_v116);
        let _index245 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1370_v113).length));
        let _rhs255 = (_dafny.ZERO).minus(_this.f15);
        let _rhs256 = _1373_v115;
        let _rhs257 = new BigNumber(-813);
        let _rhs258 = _module.D15.create_DC27(_1374_v116);
        let _lhs177 = _1370_v113;
        let _lhs178 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1370_v113).length));
        let _lhs179 = _this;
        _lhs177[_lhs178] = _rhs255;
        _1373_v115 = _rhs256;
        _lhs179.f15 = _rhs257;
        _1375_v117 = _rhs258;
        let _index246 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1370_v113).length));
        (_1370_v113)[_index246] = new BigNumber(327);
        let _1376_v118;
        let _nw218 = Array((_dafny.ONE).toNumber()).fill(_module.D2.Default());
        _1376_v118 = _nw218;
        _1376_v118 = _1376_v118;
        let _1377_v119;
        _1377_v119 = false;
        let _1378_v120;
        _1378_v120 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), (_this).f5, (_this).f5, (_this).f5, new _dafny.CodePoint('c'.codePointAt(0)));
        _1377_v119 = ((_dafny.ZERO).minus(_1369_i14)).isLessThan(new BigNumber((_1378_v120).cardinality()));
      }
      let _1379_v121;
      _1379_v121 = false;
      _1379_v121 = (_1379_v121) && (_1379_v121);
      let _1380_v122;
      _1380_v122 = _dafny.Seq.UnicodeFromString("wqjcniels");
      _1380_v122 = _1380_v122;
      let _1381_v123;
      let _nw219 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _1381_v123 = _nw219;
      let _1382_v124;
      let _nw220 = new _module.C1();
      _nw220.__ctor();
      _1382_v124 = _nw220;
      let _1383_v125;
      let _nw221 = Array((new BigNumber(26)).toNumber()).fill(false);
      _1383_v125 = _nw221;
      let _1384_v126;
      _1384_v126 = _dafny.Map.Empty.slice().updateUnsafe(_1382_v124,_1383_v125);
      let _1385_v127;
      _1385_v127 = _1384_v126;
      let _index247 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1381_v123).length));
      (_1381_v123)[_index247] = new BigNumber((_dafny.Seq.of(_1385_v127)).length);
      let _1386_v128;
      _1386_v128 = _dafny.Seq.of(_this.f15, _this.f15);
      let _1387_v129;
      _1387_v129 = _dafny.Seq.of(_1386_v128);
      let _1388_v130;
      _1388_v130 = _dafny.Set.fromElements(_1379_v121);
      let _1389_v131;
      _1389_v131 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1387_v129, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1388_v130).length)), new BigNumber((_1387_v129).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(764)), function (_1390_i16) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("utb")).length);
      })),_1380_v122);
      let _index248 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1381_v123).length));
      let _rhs259 = _this.f15;
      let _rhs260 = (((_1389_v131).contains(_dafny.Seq.Concat(_1387_v129, _module.__default.fm52((_this).f6, false, globalState)))) ? ((_1389_v131).get(_dafny.Seq.Concat(_1387_v129, _module.__default.fm52((_this).f6, false, globalState)))) : (_dafny.Seq.UnicodeFromString("amowydn")));
      let _lhs180 = _1381_v123;
      let _lhs181 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1381_v123).length));
      _lhs180[_lhs181] = _rhs259;
      _1380_v122 = _rhs260;
      r0 = (_1386_v128)[_module.__default.safeIndex((_1381_v123)[_module.__default.safeIndex(new BigNumber(712), new BigNumber((_1381_v123).length))], new BigNumber((_1386_v128).length))];
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _1391_v0;
      _1391_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f15);
      let _hi8 = (_this.f15).multipliedBy((_dafny.ZERO).minus(_this.f15));
      for (let _1392_i0 = (((_1391_v0).contains(_module.__default.fm0(p2, globalState))) ? ((_1391_v0).get(_module.__default.fm0(p2, globalState))) : (_this.f15)); _1392_i0.isLessThan(_hi8); _1392_i0 = _1392_i0.plus(_dafny.ONE)) {
        let _1393_v1;
        _1393_v1 = _dafny.Seq.UnicodeFromString("petsff");
        let _1394_v2;
        _1394_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1393_v1);
        let _1395_v3;
        let _nw222 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _1395_v3 = _nw222;
        let _1396_v4;
        _1396_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1395_v3,_1392_i0);
        let _1397_v5;
        _1397_v5 = _dafny.Map.Empty.slice().updateUnsafe(false,!((_this).f6));
        let _1398_v6;
        let _nw223 = Array((new BigNumber(25)).toNumber());
        _nw223[(_dafny.ZERO).toNumber()] = _1392_i0;
        _nw223[(_dafny.ONE).toNumber()] = _this.f15;
        _nw223[(new BigNumber(2)).toNumber()] = ((_this).fm7(p2, globalState)).plus((_dafny.ZERO).minus(_1392_i0));
        _nw223[(new BigNumber(3)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(4)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(5)).toNumber()] = (_1392_i0).multipliedBy(_this.f15);
        _nw223[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((((_this).f6) ? (_1393_v1) : (_1393_v1))).length));
        _nw223[(new BigNumber(7)).toNumber()] = _1392_i0;
        _nw223[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1394_v2).length));
        _nw223[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_1392_i0);
        _nw223[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_1396_v4).contains(_1395_v3)) ? ((_1396_v4).get(_1395_v3)) : (_1392_i0))));
        _nw223[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1392_i0));
        _nw223[(new BigNumber(12)).toNumber()] = _1392_i0;
        _nw223[(new BigNumber(13)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(14)).toNumber()] = (new BigNumber(317)).multipliedBy(new BigNumber((_1393_v1).length));
        _nw223[(new BigNumber(15)).toNumber()] = new BigNumber((_1393_v1).length);
        _nw223[(new BigNumber(16)).toNumber()] = new BigNumber(362);
        _nw223[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f15, _1392_i0));
        _nw223[(new BigNumber(18)).toNumber()] = _1392_i0;
        _nw223[(new BigNumber(19)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(20)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(21)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_1392_i0);
        _nw223[(new BigNumber(23)).toNumber()] = _this.f15;
        _nw223[(new BigNumber(24)).toNumber()] = (_this.f15).multipliedBy(new BigNumber((_1397_v5).length));
        _1398_v6 = _nw223;
        _1395_v3 = _1398_v6;
        if (_module.__default.fm0(!((_this).f6) || (p2), globalState)) {
          (globalState).f1 = ((_dafny.ZERO).minus(_1392_i0)).plus(_this.f15);
          let _1399_v7;
          _1399_v7 = true;
          _1399_v7 = p2;
          let _1400_v8;
          let _nw224 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1400_v8 = _nw224;
          let _index249 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1400_v8).length));
          (_1400_v8)[_index249] = _1399_v7;
          let _1401_v9;
          _1401_v9 = _module.D11.create_DC20(_1399_v7, new BigNumber(77));
          let _index250 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1400_v8).length));
          (_1400_v8)[_index250] = (_1401_v9).dtor_cf21;
          let _1402_v10;
          _1402_v10 = _dafny.Seq.of(_1400_v8, _1400_v8, _1400_v8);
          let _1403_v11;
          let _nw225 = Array((new BigNumber(16)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1400_v8;
          _nw225[(_dafny.ONE).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(2)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(3)).toNumber()] = (((_this).f6) ? (_1400_v8) : (_1400_v8));
          _nw225[(new BigNumber(4)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(5)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(6)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(7)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(8)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(9)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(10)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(11)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(12)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(13)).toNumber()] = _1400_v8;
          _nw225[(new BigNumber(14)).toNumber()] = (_1402_v10)[_module.__default.safeIndex(_1392_i0, new BigNumber((_1402_v10).length))];
          _nw225[(new BigNumber(15)).toNumber()] = _1400_v8;
          _1403_v11 = _nw225;
          _1403_v11 = _1403_v11;
          let _1404_v12;
          _1404_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1392_i0,new BigNumber((_dafny.Seq.of((_1400_v8)[_module.__default.safeIndex(new BigNumber(256), new BigNumber((_1400_v8).length))])).length));
          let _1405_v13;
          _1405_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,new BigNumber((_1404_v12).length));
          let _1406_v14;
          _1406_v14 = _dafny.MultiSet.fromElements((_this).f5);
          _1399_v7 = (_dafny.MultiSet.fromElements((_this).f5, _module.__default.fm53(_1393_v1, (_dafny.ZERO).minus(new BigNumber((_1405_v13).length)), globalState))).IsDisjointFrom((_1406_v14).Union(_dafny.MultiSet.fromElements((_this).f5, (_this).f5)));
        } else {
          let _1407_v15;
          _1407_v15 = true;
          _1407_v15 = (_1392_i0).isLessThan(_1392_i0);
          let _1408_v16;
          let _init41 = ((_1409_p0) => function (_1410_i1) {
            return _1409_p0;
          })(p0);
          let _nw226 = Array((new BigNumber(27)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw226.length); _i0_41++) {
            _nw226[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1408_v16 = _nw226;
          let _1411_v17;
          _1411_v17 = new _dafny.CodePoint('b'.codePointAt(0));
          let _1412_v18;
          _1412_v18 = _dafny.Seq.of(new BigNumber((_1393_v1).length), _1392_i0);
          let _rhs261 = _this.f15;
          let _rhs262 = _1408_v16;
          let _rhs263 = _1411_v17;
          let _rhs264 = (_this).fm6(_1412_v18, globalState);
          let _lhs182 = globalState;
          _lhs182.f1 = _rhs261;
          _1408_v16 = _rhs262;
          _1411_v17 = _rhs263;
          _1407_v15 = _rhs264;
          (globalState).f1 = _this.f15;
          (globalState).f4 = (p1).Intersect(_dafny.MultiSet.fromElements(p0));
          _1411_v17 = _1411_v17;
        }
        let _1413_v19;
        _1413_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,p2);
        _1413_v19 = (_1413_v19).update(_module.__default.safeDivisionInt(_this.f15, _this.f15), (_this).f6);
        let _1414_v20;
        let _nw227 = new _module.C3();
        _nw227.__ctor((_this).f6, p2);
        _1414_v20 = _nw227;
      }
      let _1415_v21;
      _1415_v21 = _dafny.Seq.of((_this).f6);
      _1415_v21 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f6), _dafny.Seq.update(_1415_v21, _module.__default.safeIndex(_this.f15, new BigNumber((_1415_v21).length)), (_this).f6));
      let _1416_v22;
      _1416_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_this.f15);
      let _1417_v23;
      _1417_v23 = _1416_v22;
      let _1418_v24;
      _1418_v24 = _dafny.Seq.of(_1415_v21);
      let _1419_v25;
      _1419_v25 = _dafny.Seq.UnicodeFromString("vinlx");
      let _1420_v26;
      _1420_v26 = _module.D12.create_DC22(_1417_v23, _1418_v24, (_1419_v25)[_module.__default.safeIndex(_this.f15, new BigNumber((_1419_v25).length))], p0, _1419_v25);
      let _source20 = _1420_v26;
      if (_source20.is_DC22) {
        let _1421___mcc_h0 = (_source20).cf24;
        let _1422___mcc_h1 = (_source20).cf25;
        let _1423___mcc_h2 = (_source20).cf26;
        let _1424___mcc_h3 = (_source20).cf27;
        let _1425___mcc_h4 = (_source20).cf28;
        let _1426_cf28 = _1425___mcc_h4;
        let _1427_cf27 = _1424___mcc_h3;
        let _1428_cf26 = _1423___mcc_h2;
        let _1429_cf25 = _1422___mcc_h1;
        let _1430_cf24 = _1421___mcc_h0;
        _1419_v25 = _dafny.Seq.UnicodeFromString("pnlpvun");
        if ((_this).f6) {
          _1426_cf28 = _dafny.Seq.UnicodeFromString("stpkvw");
          let _1431_v27;
          let _nw228 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _1431_v27 = _nw228;
          let _1432_v28;
          let _nw229 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1432_v28 = _nw229;
          let _1433_v29;
          _1433_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1432_v28,_this.f15);
          let _index251 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1431_v27).length));
          (_1431_v27)[_index251] = _1433_v29;
          let _index252 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1431_v27).length));
          (_1431_v27)[_index252] = _1433_v29;
          let _1434_v30;
          _1434_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1428_cf26,_this.f15);
          _1434_v30 = (_1434_v30).update(_module.__default.fm53(_1419_v25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,!(_1427_cf27))).length), globalState), _this.f15);
          let _1435_v31;
          _1435_v31 = _dafny.MultiSet.fromElements(_this.f15);
          _1427_cf27 = (_1435_v31).IsDisjointFrom(_1435_v31);
          (_this).f15 = new BigNumber(64);
        } else {
          let _1436_v32;
          _1436_v32 = _dafny.Seq.of(_this.f15);
          let _1437_v33;
          _1437_v33 = _dafny.Seq.of(_1436_v32);
          let _1438_v34;
          let _nw230 = new _module.C0();
          _nw230.__ctor(!(!(_1427_cf27)), _dafny.Seq.Concat(_1437_v33, _1437_v33));
          _1438_v34 = _nw230;
          let _1439_v35;
          let _nw231 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1439_v35 = _nw231;
          let _1440_v36;
          _1440_v36 = _dafny.Set.fromElements(_this.f15);
          let _index253 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length));
          (_1439_v35)[_index253] = !(_1440_v36).contains(_this.f15);
          let _1441_v37;
          _1441_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1439_v35,_dafny.Set.fromElements(new BigNumber(563)));
          let _index254 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length));
          let _rhs265 = !((p0) === (!_dafny.areEqual(_1415_v21, _1415_v21)));
          let _rhs266 = ((p1).Intersect(_dafny.MultiSet.fromElements(p2, p2, (_this).f6))).Union(_dafny.MultiSet.fromElements(p0));
          let _rhs267 = (((_this).fm6(_1436_v32, globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe(_1439_v35,_1440_v36)) : (_1441_v37));
          let _lhs183 = _1439_v35;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length));
          let _lhs185 = globalState;
          _lhs183[_lhs184] = _rhs265;
          _lhs185.f4 = _rhs266;
          _1441_v37 = _rhs267;
          let _1442_v38;
          let _nw232 = new _module.C3();
          _nw232.__ctor(!(p0), (_1439_v35)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length))]);
          _1442_v38 = _nw232;
          let _index255 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length));
          let _rhs268 = (_1439_v35)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length))];
          let _rhs269 = (_1438_v34).f11;
          let _lhs186 = _1439_v35;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_1439_v35).length));
          _lhs186[_lhs187] = _rhs268;
          _1427_cf27 = _rhs269;
          let _1443_v39;
          _1443_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1442_v38).f14,p2);
          let _1444_v40;
          let _nw233 = new _module.C3();
          _nw233.__ctor(!((_1443_v39).equals((_1443_v39).update((_this).f6, p0))), p2);
          _1444_v40 = _nw233;
        }
        let _1445_v41;
        let _nw234 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Set.Empty);
        _1445_v41 = _nw234;
        let _1446_v42;
        _1446_v42 = _dafny.Seq.of(_this.f15);
        let _1447_v43;
        _1447_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).fm6(_1446_v42, globalState));
        let _1448_v44;
        _1448_v44 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1447_v43);
        let _1449_v45;
        _1449_v45 = _dafny.Set.fromElements((((_1448_v44).contains(_this.f15)) ? ((_1448_v44).get(_this.f15)) : (_1447_v43)));
        let _index256 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1445_v41).length));
        (_1445_v41)[_index256] = _1449_v45;
        let _1450_v46;
        _1450_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1449_v45);
        let _1451_v47;
        _1451_v47 = _dafny.Seq.of(_1447_v43);
        let _index257 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1445_v41).length));
        (_1445_v41)[_index257] = (((_1450_v46).contains((p0) && (p0))) ? ((_1450_v46).get((p0) && (p0))) : (function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of (_1451_v47).Elements) {
            let _1452_v48 = _compr_50;
            if (_dafny.Seq.contains(_1451_v47, _1452_v48)) {
              _coll50.add(_1452_v48);
            }
          }
          return _coll50;
        }()));
        let _1453_v49;
        let _nw235 = Array((new BigNumber(28)).toNumber()).fill(false);
        _1453_v49 = _nw235;
        let _1454_v50;
        _1454_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1453_v49,_module.__default.safeModuloInt(_this.f15, _this.f15));
        _1454_v50 = (_1454_v50).update(_1453_v49, _module.__default.safeDivisionInt(_this.f15, _this.f15));
      } else if (_source20.is_DC23) {
        let _1455___mcc_h5 = (_source20).cf29;
        let _1456_cf29 = _1455___mcc_h5;
        let _1457_v51;
        _1457_v51 = _module.D15.create_DC28((false) === (p0), ((_module.D1.create_DC4(_this.f15)).dtor_cf4).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), (_this).f5, new _dafny.CodePoint('g'.codePointAt(0)), (_this).f5, (_this).f5)).cardinality())));
        let _pat_let_tv18 = p0;
        _1457_v51 = function (_pat_let16_0) {
          return function (_1458_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_1459_dt__update_hcf37_h0) {
                return _module.D15.create_DC28((_1458_dt__update__tmp_h0).dtor_cf36, _1459_dt__update_hcf37_h0);
              }(_pat_let17_0);
            }(_pat_let_tv18);
          }(_pat_let16_0);
        }(_module.__default.fm54(globalState));
        (globalState).f1 = _this.f15;
        (globalState).f1 = _this.f15;
        let _1460_v52;
        _1460_v52 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,p2);
        let _1461_v53;
        let _1462_v54;
        let _1463_v55;
        let _1464_v56;
        let _out40;
        let _out41;
        let _out42;
        let _out43;
        let _outcollector18 = (_this).m12(new BigNumber((_module.__default.fm55(_this.f15, _this.f15, globalState)).length), (_this.f15).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_1460_v52).length))), _this.f15, globalState);
        _out40 = _outcollector18[0];
        _out41 = _outcollector18[1];
        _out42 = _outcollector18[2];
        _out43 = _outcollector18[3];
        _1461_v53 = _out40;
        _1462_v54 = _out41;
        _1463_v55 = _out42;
        _1464_v56 = _out43;
      } else {
        let _1465___mcc_h6 = (_source20).cf23;
        let _1466_cf23 = _1465___mcc_h6;
        let _1467_v57;
        let _init42 = function (_1468_i2) {
          return (_this).f5;
        };
        let _nw236 = Array((new BigNumber(24)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw236.length); _i0_42++) {
          _nw236[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1467_v57 = _nw236;
        _1467_v57 = _1467_v57;
        (globalState).f1 = new BigNumber(229);
        let _1469_v58;
        let _nw237 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1469_v58 = _nw237;
        let _index258 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1469_v58).length));
        (_1469_v58)[_index258] = (_dafny.MultiSet.fromElements((_this).f6)).Intersect(p1);
        let _index259 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1469_v58).length));
        (_1469_v58)[_index259] = p1;
        let _1470_v59;
        let _nw238 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1470_v59 = _nw238;
        let _1471_v60;
        _1471_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1470_v59,(((_this).f6) ? ((_this).f6) : (p0)));
        if ((((_1471_v60).contains(_1470_v59)) ? ((_1471_v60).get(_1470_v59)) : (true))) {
          (_this).f15 = new BigNumber((_dafny.Seq.UnicodeFromString("yl")).length);
          let _1472_v61;
          _1472_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_1470_v59);
          _1472_v61 = (_1472_v61).update(_this.f15, _1470_v59);
          let _1473_v62;
          _1473_v62 = false;
          _1473_v62 = p0;
          (globalState).f1 = (_dafny.ZERO).minus((_this.f15).plus(_this.f15));
          let _1474_v63;
          _1474_v63 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_dafny.ZERO).minus(new BigNumber((_1419_v25).length))));
          let _1475_v64;
          _1475_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1474_v63,new _dafny.CodePoint('p'.codePointAt(0)));
          _1475_v64 = (_1475_v64).update(_dafny.Seq.Concat(_1474_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), ((_1476_v0) => function (_1477_i3) {
            return _1476_v0;
          })(_1391_v0))), (_this).f5);
        } else {
          let _1478_v65;
          _1478_v65 = false;
          let _1479_v66;
          let _init43 = function (_1480_i4) {
            return _module.__default.safeDivisionInt(_1480_i4, _this.f15);
          };
          let _nw239 = Array((new BigNumber(22)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw239.length); _i0_43++) {
            _nw239[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1479_v66 = _nw239;
          let _1481_v67;
          _1481_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1479_v66,p0);
          let _1482_v68;
          _1482_v68 = _dafny.Seq.of(_1479_v66, _1479_v66);
          let _1483_v69;
          _1483_v69 = _module.D1.create_DC5(_1478_v65);
          _1478_v65 = !(!((((_1481_v67).contains((_1482_v68)[_module.__default.safeIndex(new BigNumber((_1415_v21).length), new BigNumber((_1482_v68).length))])) ? ((_1481_v67).get((_1482_v68)[_module.__default.safeIndex(new BigNumber((_1415_v21).length), new BigNumber((_1482_v68).length))])) : ((_1483_v69).dtor_cf5))));
          (globalState).f1 = _this.f15;
          let _1484_v70;
          _1484_v70 = _dafny.MultiSet.fromElements(_this.f15);
          let _1485_v71;
          _1485_v71 = _dafny.Seq.of(new BigNumber((_1419_v25).length));
          let _1486_v72;
          _1486_v72 = _dafny.Seq.of((((_1484_v70).contains((((p1).contains(_1478_v65)) ? ((p1).get(_1478_v65)) : (_this.f15)))) ? ((_1484_v70).get((((p1).contains(_1478_v65)) ? ((p1).get(_1478_v65)) : (_this.f15)))) : (_this.f15)), new BigNumber((_1485_v71).length), new BigNumber((_dafny.Seq.UnicodeFromString("sf")).length), _this.f15, _this.f15);
          let _1487_v73;
          _1487_v73 = _dafny.MultiSet.fromElements(new BigNumber((_1485_v71).length));
          let _1488_v74;
          _1488_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1484_v70,_this.f15);
          (_this).f15 = ((((p0) ? ((_this).f6) : (p0))) ? ((((_1488_v74).contains(_1487_v73)) ? ((_1488_v74).get(_1487_v73)) : (_this.f15))) : (_this.f15));
          let _1489_v75;
          _1489_v75 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1490_v76;
          _1490_v76 = _module.D1.create_DC4(_this.f15);
          let _rhs270 = !(p0);
          let _rhs271 = (p2) && (!((p1).IsProperSubsetOf((_1469_v58)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1469_v58).length))])));
          let _rhs272 = !((_this).f6) || (!(_module.__default.fm0((_this).f6, globalState)));
          let _rhs273 = _1489_v75;
          let _rhs274 = _1490_v76;
          _1478_v65 = _rhs270;
          _1478_v65 = _rhs271;
          _1478_v65 = _rhs272;
          _1489_v75 = _rhs273;
          _1490_v76 = _rhs274;
          _1478_v65 = (_1478_v65) === (p2);
        }
      }
      let _1491_v77;
      _1491_v77 = _module.D15.create_DC28(p2, true);
      let _1492_v78;
      _1492_v78 = _module.D15.create_DC29(_1491_v77);
      let _source21 = _1492_v78;
      if (_source21.is_DC28) {
        let _1493___mcc_h7 = (_source21).cf36;
        let _1494___mcc_h8 = (_source21).cf37;
        let _1495_cf37 = _1494___mcc_h8;
        let _1496_cf36 = _1493___mcc_h7;
        let _1497_v79;
        let _nw240 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _1497_v79 = _nw240;
        let _index260 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1497_v79).length));
        (_1497_v79)[_index260] = _this.f15;
        let _1498_v80;
        _1498_v80 = _dafny.Seq.of(new BigNumber(-106));
        let _1499_v81;
        _1499_v81 = _dafny.MultiSet.fromElements((_this.f15).multipliedBy(_this.f15), (_1498_v80)[_module.__default.safeIndex(new BigNumber(851), new BigNumber((_1498_v80).length))], _this.f15);
        let _1500_v82;
        _1500_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1496_cf36);
        let _1501_v83;
        _1501_v83 = _module.D12.create_DC21(_1500_v82);
        let _1502_v84;
        _1502_v84 = _module.D2.create_DC7(_this.f15, (_this).fm7(_module.__default.fm0(_1495_cf37, globalState), globalState));
        let _index261 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1497_v79).length));
        let _rhs275 = p0;
        let _rhs276 = (_dafny.ZERO).minus(_this.f15);
        let _rhs277 = ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_1419_v25, _dafny.Seq.update(_1419_v25, _module.__default.safeIndex(new BigNumber(746), new BigNumber((_1419_v25).length)), (_this).f5))).length), _this.f15, new BigNumber(597))).Union(_1499_v81)).Union(_dafny.MultiSet.fromElements(_this.f15, new BigNumber(-663), new BigNumber(804)));
        let _rhs278 = (_1502_v84).dtor_cf7;
        let _rhs279 = _1501_v83;
        let _lhs188 = _1497_v79;
        let _lhs189 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1497_v79).length));
        let _lhs190 = globalState;
        _1496_cf36 = _rhs275;
        _lhs188[_lhs189] = _rhs276;
        _1499_v81 = _rhs277;
        _lhs190.f1 = _rhs278;
        _1501_v83 = _rhs279;
        let _1503_v85;
        let _nw241 = new _module.C3();
        _nw241.__ctor(_1495_cf37, p2);
        _1503_v85 = _nw241;
        let _1504_v86;
        _1504_v86 = _module.D11.create_DC20((_this).f6, (_1497_v79)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1497_v79).length))]);
        let _1505_v87;
        let _nw242 = new _module.C2();
        _nw242.__ctor((_this).f5, (_1504_v86).dtor_cf21);
        _1505_v87 = _nw242;
        _1505_v87 = _1505_v87;
        let _1506_v88;
        let _nw243 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1506_v88 = _nw243;
        let _index262 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1506_v88).length));
        (_1506_v88)[_index262] = _1496_cf36;
        let _1507_v89;
        let _nw244 = Array((new BigNumber(28)).toNumber()).fill([]);
        _1507_v89 = _nw244;
        let _1508_v90;
        let _nw245 = Array((new BigNumber(18)).toNumber()).fill([]);
        _1508_v90 = _nw245;
        let _index263 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1507_v89).length));
        (_1507_v89)[_index263] = _1508_v90;
        let _index264 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1506_v88).length));
        let _index265 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1507_v89).length));
        let _rhs280 = !(_1496_cf36);
        let _rhs281 = _1508_v90;
        let _lhs191 = _1506_v88;
        let _lhs192 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_1506_v88).length));
        let _lhs193 = _1507_v89;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1507_v89).length));
        _lhs191[_lhs192] = _rhs280;
        _lhs193[_lhs194] = _rhs281;
      } else if (_source21.is_DC27) {
        let _1509___mcc_h9 = (_source21).cf35;
        let _1510_cf35 = _1509___mcc_h9;
        let _1511_v91;
        let _nw246 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1511_v91 = _nw246;
        let _1512_v92;
        _1512_v92 = _1511_v91;
        let _1513_v93;
        _1513_v93 = _dafny.Seq.of(_1512_v92);
        _1513_v93 = _dafny.Seq.Concat(_1513_v93, _dafny.Seq.Concat(_1513_v93, _1513_v93));
        let _1514_v94;
        let _nw247 = new _module.C1();
        _nw247.__ctor();
        _1514_v94 = _nw247;
        let _1515_v95;
        _1515_v95 = _module.D11.create_DC20(p0, _this.f15);
        let _1516_v96;
        _1516_v96 = _dafny.Seq.of(_this.f15);
        let _1517_v97;
        _1517_v97 = _dafny.MultiSet.fromElements((_this).fm7(p0, globalState), _this.f15);
        let _1518_v98;
        let _nw248 = Array((new BigNumber(21)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = _this.f15;
        _nw248[(_dafny.ONE).toNumber()] = _this.f15;
        _nw248[(new BigNumber(2)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(3)).toNumber()] = (_1515_v95).dtor_cf22;
        _nw248[(new BigNumber(4)).toNumber()] = (_1516_v96)[_module.__default.safeIndex(_this.f15, new BigNumber((_1516_v96).length))];
        _nw248[(new BigNumber(5)).toNumber()] = new BigNumber(632);
        _nw248[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_this.f15);
        _nw248[(new BigNumber(7)).toNumber()] = new BigNumber(725);
        _nw248[(new BigNumber(8)).toNumber()] = new BigNumber(223);
        _nw248[(new BigNumber(9)).toNumber()] = new BigNumber(898);
        _nw248[(new BigNumber(10)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(11)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(_this.f15, _this.f15);
        _nw248[(new BigNumber(13)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(14)).toNumber()] = (_1514_v94).fm14(globalState);
        _nw248[(new BigNumber(15)).toNumber()] = ((((_1517_v97).contains((_1514_v94).fm4(_1415_v21, new BigNumber(832), _this.f15, _1517_v97, globalState))) ? ((_1517_v97).get((_1514_v94).fm4(_1415_v21, new BigNumber(832), _this.f15, _1517_v97, globalState))) : ((_this).fm7(p2, globalState)))).minus(_this.f15);
        _nw248[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-5)), function (_1519_i5) {
          return (_this).f5;
        }), _1419_v25)).length);
        _nw248[(new BigNumber(17)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(18)).toNumber()] = (_this.f15).multipliedBy(_this.f15);
        _nw248[(new BigNumber(19)).toNumber()] = _this.f15;
        _nw248[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_1516_v96)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1516_v96).length))]);
        _1518_v98 = _nw248;
        let _index266 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        (_1518_v98)[_index266] = new BigNumber(-44);
        let _1520_v99;
        _1520_v99 = true;
        let _index267 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        let _rhs282 = ((_this).fm8((_this).f6, p2, p2, globalState)).minus(new BigNumber(-696));
        let _rhs283 = (_1514_v94).fm7(p2, globalState);
        let _rhs284 = !(_module.__default.fm3(p2, (_this).f6, globalState)).isEqualTo(_this.f15);
        let _rhs285 = (_dafny.ZERO).minus((_1516_v96)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1516_v96, _1516_v96), _module.__default.safeIndex(_this.f15, new BigNumber((_dafny.Seq.Concat(_1516_v96, _1516_v96)).length)), _this.f15)).length), new BigNumber((_1516_v96).length))]);
        let _lhs195 = globalState;
        let _lhs196 = _1518_v98;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        let _lhs198 = globalState;
        _lhs195.f1 = _rhs282;
        _lhs196[_lhs197] = _rhs283;
        _1520_v99 = _rhs284;
        _lhs198.f1 = _rhs285;
        let _1521_v100;
        _1521_v100 = _dafny.MultiSet.fromElements(_1511_v91, _1511_v91, _1511_v91);
        let _index268 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        let _index269 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        let _rhs286 = (((_1521_v100).contains((((_this).f6) ? (_1511_v91) : (_1511_v91)))) ? ((_1521_v100).get((((_this).f6) ? (_1511_v91) : (_1511_v91)))) : (_this.f15));
        let _rhs287 = (_1518_v98)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length))];
        let _lhs199 = _1518_v98;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        let _lhs201 = _1518_v98;
        let _lhs202 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1518_v98).length));
        _lhs199[_lhs200] = _rhs286;
        _lhs201[_lhs202] = _rhs287;
      } else {
        let _1522___mcc_h10 = (_source21).cf38;
        let _1523_cf38 = _1522___mcc_h10;
        let _1524_v101;
        _1524_v101 = true;
        let _1525_v102;
        _1525_v102 = _dafny.Seq.of(_this.f15, _this.f15);
        let _rhs288 = (_this).f6;
        let _rhs289 = _this.f15;
        let _rhs290 = (new BigNumber((_1525_v102).length)).multipliedBy(_module.__default.safeDivisionInt(_this.f15, _this.f15));
        let _rhs291 = new BigNumber(148);
        let _lhs203 = globalState;
        let _lhs204 = globalState;
        let _lhs205 = globalState;
        _1524_v101 = _rhs288;
        _lhs203.f1 = _rhs289;
        _lhs204.f1 = _rhs290;
        _lhs205.f1 = _rhs291;
        let _1526_v103;
        _1526_v103 = _dafny.Set.fromElements(p2);
        _1524_v101 = (_dafny.Set.fromElements(p0, _1524_v101)).IsProperSubsetOf(_1526_v103);
        _1524_v101 = (_1415_v21)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_this.f15))).length), new BigNumber((_1415_v21).length))];
        let _1527_v104;
        _1527_v104 = _dafny.MultiSet.fromElements(_this.f15, _this.f15);
        let _1528_v105;
        let _nw249 = Array((new BigNumber(20)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = _1527_v104;
        _nw249[(_dafny.ONE).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(2)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(3)).toNumber()] = (_1527_v104).Intersect(_1527_v104);
        _nw249[(new BigNumber(4)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1525_v102, _dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_1529_i6) {
          return _this.f15;
        })));
        _nw249[(new BigNumber(6)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(7)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(8)).toNumber()] = (_module.__default.fm45(globalState)).Union(_1527_v104);
        _nw249[(new BigNumber(9)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(10)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(11)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(12)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(13)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(14)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_this.f15, _this.f15)).Union(_1527_v104);
        _nw249[(new BigNumber(16)).toNumber()] = (_module.__default.fm45(globalState)).Intersect(_1527_v104);
        _nw249[(new BigNumber(17)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(18)).toNumber()] = _1527_v104;
        _nw249[(new BigNumber(19)).toNumber()] = _1527_v104;
        _1528_v105 = _nw249;
        let _rhs292 = false;
        let _rhs293 = new BigNumber(-677);
        let _rhs294 = _1528_v105;
        let _rhs295 = _1524_v101;
        let _lhs206 = globalState;
        _1524_v101 = _rhs292;
        _lhs206.f1 = _rhs293;
        _1528_v105 = _rhs294;
        _1524_v101 = _rhs295;
      }
      let _1530_v106;
      _1530_v106 = false;
      let _1531_v107;
      let _nw250 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _1531_v107 = _nw250;
      let _1532_v108;
      _1532_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1531_v107,(_this.f15).plus(_this.f15));
      let _1533_v109;
      _1533_v109 = _dafny.Seq.of(_this.f15);
      let _1534_v110;
      _1534_v110 = _module.D14.create_DC26((_this).fm6(_1533_v109, globalState), _this.f15, (_this).f5);
      let _rhs296 = (_1534_v110).dtor_cf33;
      let _rhs297 = !_dafny.areEqual(_dafny.Seq.Concat(_1419_v25, _1419_v25), _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1419_v25, _module.__default.safeIndex(_this.f15, new BigNumber((_1419_v25).length)), (_this).f5), _1419_v25), _module.__default.safeIndex(_this.f15, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1419_v25, _module.__default.safeIndex(_this.f15, new BigNumber((_1419_v25).length)), (_this).f5), _1419_v25)).length)), (_this).f5), _module.__default.safeIndex(_this.f15, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1419_v25, _module.__default.safeIndex(_this.f15, new BigNumber((_1419_v25).length)), (_this).f5), _1419_v25), _module.__default.safeIndex(_this.f15, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1419_v25, _module.__default.safeIndex(_this.f15, new BigNumber((_1419_v25).length)), (_this).f5), _1419_v25)).length)), (_this).f5)).length)), (_this).f5));
      let _rhs298 = (_1532_v108).Merge(_1532_v108);
      let _lhs207 = _this;
      _lhs207.f15 = _rhs296;
      _1530_v106 = _rhs297;
      _1532_v108 = _rhs298;
      let _1535_v111;
      _1535_v111 = _module.D6.create_DC13(_1533_v109);
      let _pat_let_tv19 = _1533_v109;
      _1535_v111 = function (_pat_let18_0) {
        return function (_1536_dt__update__tmp_h1) {
          return function (_pat_let19_0) {
            return function (_1537_dt__update_hcf14_h0) {
              return _module.D6.create_DC13(_1537_dt__update_hcf14_h0);
            }(_pat_let19_0);
          }(_pat_let_tv19);
        }(_pat_let18_0);
      }(_1535_v111);
      return;
    }
    m12(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _1538_v0;
      let _nw251 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1538_v0 = _nw251;
      let _1539_v1;
      _1539_v1 = _dafny.Seq.UnicodeFromString("d");
      let _1540_v2;
      _1540_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_1539_v1).length));
      let _1541_v3;
      _1541_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,(_this).f6);
      let _1542_v4;
      _1542_v4 = _module.D2.create_DC7(_this.f15, new BigNumber((_1541_v3).length));
      let _1543_v5;
      _1543_v5 = _dafny.Seq.of(_this.f15);
      let _1544_v6;
      _1544_v6 = _dafny.MultiSet.fromElements((_this).f6);
      let _1545_v7;
      _1545_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,p1);
      let _1546_v8;
      let _nw252 = Array((new BigNumber(24)).toNumber());
      _nw252[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1538_v0, _1538_v0)).length));
      _nw252[(_dafny.ONE).toNumber()] = p2;
      _nw252[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_this.f15, p2);
      _nw252[(new BigNumber(3)).toNumber()] = ((p1) ? ((((_1540_v2).contains((_this).f6)) ? ((_1540_v2).get((_this).f6)) : ((_1542_v4).dtor_cf8))) : (p2));
      _nw252[(new BigNumber(4)).toNumber()] = _module.__default.fm3((_this).f6, p1, globalState);
      _nw252[(new BigNumber(5)).toNumber()] = new BigNumber(-81);
      _nw252[(new BigNumber(6)).toNumber()] = p2;
      _nw252[(new BigNumber(7)).toNumber()] = (_1543_v5)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p1)).length), new BigNumber((_1543_v5).length))];
      _nw252[(new BigNumber(8)).toNumber()] = p2;
      _nw252[(new BigNumber(9)).toNumber()] = _this.f15;
      _nw252[(new BigNumber(10)).toNumber()] = ((((_1544_v6).contains((_this).f6)) ? ((_1544_v6).get((_this).f6)) : (_this.f15))).minus(p2);
      _nw252[(new BigNumber(11)).toNumber()] = (_1543_v5)[_module.__default.safeIndex(new BigNumber(-258), new BigNumber((_1543_v5).length))];
      _nw252[(new BigNumber(12)).toNumber()] = p2;
      _nw252[(new BigNumber(13)).toNumber()] = (_this.f15).multipliedBy(new BigNumber((_1545_v7).length));
      _nw252[(new BigNumber(14)).toNumber()] = (((_1544_v6).contains(false)) ? ((_1544_v6).get(false)) : (_this.f15));
      _nw252[(new BigNumber(15)).toNumber()] = _this.f15;
      _nw252[(new BigNumber(16)).toNumber()] = ((p1) ? (p0) : (p0));
      _nw252[(new BigNumber(17)).toNumber()] = new BigNumber(886);
      _nw252[(new BigNumber(18)).toNumber()] = new BigNumber((_module.__default.fm1(p0, (_this).f6, new BigNumber(-623), globalState)).length);
      _nw252[(new BigNumber(19)).toNumber()] = (p2).minus(p2);
      _nw252[(new BigNumber(20)).toNumber()] = new BigNumber((_1540_v2).length);
      _nw252[(new BigNumber(21)).toNumber()] = _this.f15;
      _nw252[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw252[(new BigNumber(23)).toNumber()] = p2;
      _1546_v8 = _nw252;
      let _index270 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length));
      (_1546_v8)[_index270] = _module.__default.safeModuloInt(p0, _this.f15);
      let _index271 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length));
      (_1546_v8)[_index271] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(364)), function (_1547_i0) {
        return (_this).f5;
      })).length)).minus(p2);
      let _1548_i1;
      _1548_i1 = _dafny.ZERO;
      L14: {
        while (p1) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1548_i1)) {
              break L14;
            }
            _1548_i1 = (_1548_i1).plus(_dafny.ONE);
            let _1549_v9;
            _1549_v9 = _dafny.Seq.of(_1543_v5);
            let _1550_v10;
            let _nw253 = new _module.C0();
            _nw253.__ctor((_this).f6, _1549_v9);
            _1550_v10 = _nw253;
            r3 = true;
            r3 = (_this).f6;
            let _1551_v11;
            let _nw254 = Array((new BigNumber(24)).toNumber()).fill(false);
            _1551_v11 = _nw254;
            let _index272 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1551_v11).length));
            (_1551_v11)[_index272] = p1;
            let _1552_v13;
            _1552_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(723),_1539_v1);
            let _index273 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1551_v11).length));
            (_1551_v11)[_index273] = (function () {
              let _coll51 = new _dafny.Map();
              for (const _compr_51 of (_module.__default.fm56(globalState)).Elements) {
                let _1553_v12 = _compr_51;
                if (_dafny.Seq.contains(_module.__default.fm56(globalState), _1553_v12)) {
                  _coll51.push([_module.__default.safeModuloInt(_1553_v12, p0),_1539_v1]);
                }
              }
              return _coll51;
            }()).equals(_1552_v13);
          }
        }
      }
      let _index274 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length));
      (_1546_v8)[_index274] = _module.__default.safeModuloInt(_this.f15, _module.__default.safeModuloInt(new BigNumber((function () {
        let _coll52 = new _dafny.Set();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(804), new BigNumber(781))) {
          let _1554_v14 = _compr_52;
          if (((new BigNumber(804)).isLessThanOrEqualTo(_1554_v14)) && ((_1554_v14).isLessThan(new BigNumber(781)))) {
            _coll52.add((_1554_v14).multipliedBy((_dafny.ZERO).minus(p2)));
          }
        }
        return _coll52;
      }()).length), p0));
      let _1555_v15;
      let _init44 = function (_1556_i2) {
        return (_this).f6;
      };
      let _nw255 = Array((new BigNumber(25)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw255.length); _i0_44++) {
        _nw255[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1555_v15 = _nw255;
      let _1557_v16;
      let _nw256 = Array((new BigNumber(17)).toNumber());
      _nw256[(_dafny.ZERO).toNumber()] = (_this).f5;
      _nw256[(_dafny.ONE).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(2)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
      _nw256[(new BigNumber(4)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(5)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(6)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(7)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(8)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(9)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(10)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(11)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(12)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(13)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(14)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(15)).toNumber()] = (_this).f5;
      _nw256[(new BigNumber(16)).toNumber()] = (function (_pat_let20_0) {
        return function (_1558_dt__update__tmp_h0) {
          return function (_pat_let21_0) {
            return function (_1559_dt__update_hcf2_h0) {
              return _module.D0.create_DC2((_1558_dt__update__tmp_h0).dtor_cf1, _1559_dt__update_hcf2_h0);
            }(_pat_let21_0);
          }((_this).f5);
        }(_pat_let20_0);
      }(_module.D0.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_1555_v15,(_this).f5), (_this).f5))).dtor_cf2;
      _1557_v16 = _nw256;
      let _index275 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1557_v16).length));
      (_1557_v16)[_index275] = (_this).f5;
      let _index276 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1557_v16).length));
      (_1557_v16)[_index276] = (_this).f5;
      let _1560_i3;
      _1560_i3 = _dafny.ZERO;
      L15: {
        while ((_this).f6) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1560_i3)) {
              break L15;
            }
            _1560_i3 = (_1560_i3).plus(_dafny.ONE);
            let _index277 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1557_v16).length));
            (_1557_v16)[_index277] = new _dafny.CodePoint('p'.codePointAt(0));
            let _index278 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1555_v15).length));
            (_1555_v15)[_index278] = ((true) ? ((_this).f6) : (p1));
            let _index279 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1555_v15).length));
            (_1555_v15)[_index279] = (_this).f6;
            r3 = (_1555_v15)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((_1555_v15).length))];
            let _1561_v17;
            _1561_v17 = _dafny.Set.fromElements((_1546_v8)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length))], _this.f15);
            let _1562_v18;
            _1562_v18 = _dafny.Set.fromElements((_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), function (_1563_i4) {
              return (_this).f5;
            })).length)).multipliedBy(new BigNumber((_1561_v17).length))), p0);
            _1561_v17 = _1562_v18;
          }
        }
      }
      let _1564_v19;
      let _init45 = ((_1565_v8) => function (_1566_i5) {
        return (_dafny.Set.fromElements((_1565_v8)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_1565_v8).length))])).Intersect(_dafny.Set.fromElements(_this.f15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)).length)));
      })(_1546_v8);
      let _nw257 = Array((new BigNumber(23)).toNumber());
      for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw257.length); _i0_45++) {
        _nw257[_i0_45] = _init45(new BigNumber(_i0_45));
      }
      _1564_v19 = _nw257;
      let _index280 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length));
      let _rhs299 = _1564_v19;
      let _rhs300 = p1;
      let _rhs301 = (_this).f6;
      let _rhs302 = _this.f15;
      let _rhs303 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-24)), ((_1567_v16) => function (_1568_i6) {
        return (_1567_v16)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_1567_v16).length))];
      })(_1557_v16)), _module.__default.safeIndex((_1546_v8)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-24)), ((_1569_v16) => function (_1570_i6) {
        return (_1569_v16)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_1569_v16).length))];
      })(_1557_v16))).length)), ((p1) ? ((_this).f5) : ((_this).f5)))).length);
      let _lhs208 = _1546_v8;
      let _lhs209 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length));
      let _lhs210 = _this;
      _1564_v19 = _rhs299;
      r3 = _rhs300;
      r3 = _rhs301;
      _lhs208[_lhs209] = _rhs302;
      _lhs210.f15 = _rhs303;
      let _1571_v20;
      _1571_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber((_1539_v1).length), (_this).fm6(_1543_v5, globalState), (_1546_v8)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length))], globalState),p0);
      r0 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vudfnvin"),p0)).Merge(_1571_v20);
      r1 = ((p1) ? (_dafny.Seq.Concat(_1543_v5, _1543_v5)) : (_1543_v5));
      let _1572_v21;
      _1572_v21 = _dafny.Set.fromElements((_this).f6);
      r2 = (_1572_v21).Difference(_module.__default.fm49(globalState));
      let _1573_v22;
      _1573_v22 = _dafny.Seq.of(_1538_v0);
      let _1574_v23;
      _1574_v23 = _dafny.Seq.of(_1539_v1);
      let _1575_v24;
      _1575_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm1(new BigNumber(138), (_this).f6, new BigNumber((_1573_v22).length), globalState), (_1574_v23)[_module.__default.safeIndex((_1546_v8)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_1546_v8).length))], new BigNumber((_1574_v23).length))]),(_this).f6);
      r3 = (((_1575_v24).contains(_dafny.Seq.UnicodeFromString("rdv"))) ? ((_1575_v24).get(_dafny.Seq.UnicodeFromString("rdv"))) : (false));
      return [r0, r1, r2, r3];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      return !(false) || (false);
    };
    fm7(p0, globalState) {
      let _this = this;
      return ((new BigNumber(982)).minus(new BigNumber(571))).multipliedBy(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of (_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).Elements) {
          let _1576_v0 = _compr_53;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).contains(_1576_v0)) {
            _coll53.push([_1576_v0,true]);
          }
        }
        return _coll53;
      }()).length)), new BigNumber(708)));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("nyudstnyx")).length), new BigNumber(832))).cardinality()),false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), function (_1577_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)))).plus(new BigNumber(111));
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4(_module.__default.safeDivisionInt(new BigNumber(276), new BigNumber(330)));
    };
    fm29(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).cardinality())).multipliedBy(new BigNumber(568)))).minus(new BigNumber(215));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      r0 = !(!(p0)) || (p0);
      let _1578_v0;
      _1578_v0 = _module.D0.create_DC1();
      let _1579_v1;
      _1579_v1 = _dafny.Set.fromElements(_1578_v0, _1578_v0, _1578_v0);
      let _1580_v2;
      _1580_v2 = _1579_v1;
      let _source22 = _1580_v2;
      let _1581___mcc_h0 = _source22;
      let _1582_cf17 = _1581___mcc_h0;
      (globalState).f1 = p1;
      r0 = (p1).isLessThanOrEqualTo(p1);
      let _1583_v3;
      let _nw258 = new _module.C1();
      _nw258.__ctor();
      _1583_v3 = _nw258;
      let _1584_v4;
      _1584_v4 = _dafny.Seq.of(p1);
      let _1585_v5;
      _1585_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1586_v6;
      let _nw259 = Array((new BigNumber(28)).toNumber());
      _nw259[(_dafny.ZERO).toNumber()] = p0;
      _nw259[(_dafny.ONE).toNumber()] = p0;
      _nw259[(new BigNumber(2)).toNumber()] = p0;
      _nw259[(new BigNumber(3)).toNumber()] = p0;
      _nw259[(new BigNumber(4)).toNumber()] = p0;
      _nw259[(new BigNumber(5)).toNumber()] = p0;
      _nw259[(new BigNumber(6)).toNumber()] = false;
      _nw259[(new BigNumber(7)).toNumber()] = true;
      _nw259[(new BigNumber(8)).toNumber()] = p0;
      _nw259[(new BigNumber(9)).toNumber()] = p0;
      _nw259[(new BigNumber(10)).toNumber()] = p0;
      _nw259[(new BigNumber(11)).toNumber()] = _module.__default.fm0(false, globalState);
      _nw259[(new BigNumber(12)).toNumber()] = p0;
      _nw259[(new BigNumber(13)).toNumber()] = p0;
      _nw259[(new BigNumber(14)).toNumber()] = p0;
      _nw259[(new BigNumber(15)).toNumber()] = (_this).fm6(_1584_v4, globalState);
      _nw259[(new BigNumber(16)).toNumber()] = p0;
      _nw259[(new BigNumber(17)).toNumber()] = (((_1585_v5).contains(p0)) ? ((_1585_v5).get(p0)) : (p0));
      _nw259[(new BigNumber(18)).toNumber()] = (_1583_v3).fm6(_1584_v4, globalState);
      _nw259[(new BigNumber(19)).toNumber()] = p0;
      _nw259[(new BigNumber(20)).toNumber()] = p0;
      _nw259[(new BigNumber(21)).toNumber()] = p0;
      _nw259[(new BigNumber(22)).toNumber()] = p0;
      _nw259[(new BigNumber(23)).toNumber()] = p0;
      _nw259[(new BigNumber(24)).toNumber()] = p0;
      _nw259[(new BigNumber(25)).toNumber()] = p0;
      _nw259[(new BigNumber(26)).toNumber()] = p0;
      _nw259[(new BigNumber(27)).toNumber()] = p0;
      _1586_v6 = _nw259;
      let _1587_v7;
      _1587_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1583_v3,_1586_v6);
      let _1588_v8;
      _1588_v8 = _1587_v7;
      let _source23 = _1588_v8;
      let _1589___mcc_h1 = _source23;
      let _1590_cf18 = _1589___mcc_h1;
      let _1591_v9;
      let _nw260 = new _module.C4();
      _nw260.__ctor();
      _1591_v9 = _nw260;
      let _1592_v10;
      _1592_v10 = _dafny.MultiSet.fromElements(_1591_v9);
      let _1593_v11;
      _1593_v11 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).Merge(_1585_v5),_1592_v10);
      let _1594_v12;
      _1594_v12 = _dafny.Seq.of(_1578_v0, _1578_v0);
      let _1595_v13;
      _1595_v13 = _dafny.Seq.of(_1592_v10);
      let _rhs304 = !(p0);
      let _rhs305 = (_dafny.Map.Empty.slice().updateUnsafe(_1585_v5,_1592_v10)).update(_1585_v5, (_1595_v13)[_module.__default.safeIndex(p1, new BigNumber((_1595_v13).length))]);
      let _rhs306 = _1594_v12;
      r0 = _rhs304;
      _1593_v11 = _rhs305;
      _1594_v12 = _rhs306;
      let _1596_v14;
      _1596_v14 = new _dafny.CodePoint('q'.codePointAt(0));
      let _1597_v15;
      let _nw261 = new _module.C3();
      _nw261.__ctor(_module.__default.fm0(p0, globalState), !_dafny.Seq.contains(p2, _1596_v14));
      _1597_v15 = _nw261;
      let _1598_v16;
      _1598_v16 = _dafny.Seq.of(!((_1597_v15).f14));
      let _1599_v17;
      let _nw262 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _1599_v17 = _nw262;
      let _1600_v18;
      _1600_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1598_v16,_1599_v17);
      let _1601_v19;
      _1601_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1600_v18,(_1597_v15).f14);
      _1601_v19 = (_1601_v19).update(_1600_v18, p0);
      let _1602_v20;
      _1602_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1598_v16,(_1597_v15).f13);
      let _1603_v21;
      let _nw263 = Array((new BigNumber(15)).toNumber());
      _nw263[(_dafny.ZERO).toNumber()] = (_1585_v5).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1597_v15).f14,(_1597_v15).f14));
      _nw263[(_dafny.ONE).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1597_v15).f14,p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(p0)));
      _nw263[(new BigNumber(3)).toNumber()] = (_1585_v5).Merge((_module.__default.fm41(p0, true, p0, globalState)).update((_1597_v15).f14, (_1597_v15).f14));
      _nw263[(new BigNumber(4)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(5)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(6)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(7)).toNumber()] = (_1585_v5).Merge(_1585_v5);
      _nw263[(new BigNumber(8)).toNumber()] = (_1585_v5).update((_1597_v15).f13, false);
      _nw263[(new BigNumber(9)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_1602_v20).contains(_1598_v16)) ? ((_1602_v20).get(_1598_v16)) : (p0)),(_1597_v15).f13);
      _nw263[(new BigNumber(11)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(12)).toNumber()] = _1585_v5;
      _nw263[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1597_v15).f13,(_1597_v15).f14)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1597_v15).f13,p0));
      _nw263[(new BigNumber(14)).toNumber()] = _1585_v5;
      _1603_v21 = _nw263;
      let _index281 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1603_v21).length));
      (_1603_v21)[_index281] = _1585_v5;
      let _1604_v22;
      _1604_v22 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus(p1));
      let _index282 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1603_v21).length));
      (_1603_v21)[_index282] = _dafny.Map.Empty.slice().updateUnsafe((_1597_v15).f14,(_1604_v22).IsSubsetOf(_1604_v22));
      let _1605_v23;
      _1605_v23 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1606_v24;
      _1606_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _1607_v25;
      let _nw264 = new _module.C4();
      _nw264.__ctor();
      _1607_v25 = _nw264;
      let _1608_v26;
      _1608_v26 = _module.D14.create_DC25(_1607_v25);
      let _1609_v27;
      _1609_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1608_v26).dtor_cf31);
      let _1610_v28;
      _1610_v28 = _dafny.Seq.of(p0);
      let _1611_v29;
      _1611_v29 = _dafny.MultiSet.fromElements(_1605_v23);
      let _1612_v30;
      let _nw265 = Array((new BigNumber(10)).toNumber());
      _nw265[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(p2, _dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), _1605_v23))).length));
      _nw265[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(339), p1);
      _nw265[(new BigNumber(2)).toNumber()] = p1;
      _nw265[(new BigNumber(3)).toNumber()] = p1;
      _nw265[(new BigNumber(4)).toNumber()] = (_1584_v4)[_module.__default.safeIndex((((_1606_v24).contains(p0)) ? ((_1606_v24).get(p0)) : (new BigNumber((_1609_v27).length))), new BigNumber((_1584_v4).length))];
      _nw265[(new BigNumber(5)).toNumber()] = p1;
      _nw265[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1585_v5).length), p1, (_dafny.ZERO).minus(new BigNumber((_1610_v28).length)))).cardinality())).multipliedBy(p1)));
      _nw265[(new BigNumber(7)).toNumber()] = p1;
      _nw265[(new BigNumber(8)).toNumber()] = (p1).multipliedBy(p1);
      _nw265[(new BigNumber(9)).toNumber()] = (((_1611_v29).contains(_1605_v23)) ? ((_1611_v29).get(_1605_v23)) : (p1));
      _1612_v30 = _nw265;
      let _index283 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1612_v30).length));
      (_1612_v30)[_index283] = p1;
      let _1613_v31;
      _1613_v31 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus(p1));
      let _1614_v32;
      _1614_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1613_v31).length));
      let _1615_v33;
      _1615_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1614_v32,p0);
      let _1616_v34;
      _1616_v34 = _dafny.Seq.of(_1610_v28);
      let _1617_v35;
      _1617_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1610_v28);
      let _index284 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1612_v30).length));
      (_1612_v30)[_index284] = new BigNumber((_module.__default.fm1(((_1607_v25).fm30(_1615_v33, _dafny.MultiSet.FromArray(_1616_v34), globalState)).multipliedBy(p1), p0, new BigNumber(((((_1617_v35).contains(p0)) ? ((_1617_v35).get(p0)) : (_1610_v28))).length), globalState)).length);
      let _1618_v36;
      let _nw266 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1618_v36 = _nw266;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1618_v36).length))) {
        let _1619_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1619_i0)) && ((_1619_i0).isLessThan(new BigNumber((_1618_v36).length))))) {
          (_1618_v36)[(_1619_i0)] = ((_dafny.MultiSet.fromElements(p0, p0)).Intersect(_dafny.MultiSet.fromElements(p0))).Union((_dafny.MultiSet.fromElements(p0)).Intersect(_dafny.MultiSet.fromElements(p0, p0, p0)));
        }
      }
      let _hi9 = new BigNumber((p2).length);
      for (let _1620_i1 = p1; _1620_i1.isLessThan(_hi9); _1620_i1 = _1620_i1.plus(_dafny.ONE)) {
        r0 = p0;
        let _1621_v37;
        let _init46 = ((_1622_p2) => function (_1623_i2) {
          return !_dafny.Seq.contains(_1622_p2, new _dafny.CodePoint('p'.codePointAt(0)));
        })(p2);
        let _nw267 = Array((new BigNumber(7)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw267.length); _i0_46++) {
          _nw267[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1621_v37 = _nw267;
        let _1624_v38;
        _1624_v38 = _dafny.Seq.of(false, p0, p0, true);
        let _1625_v39;
        _1625_v39 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length));
        let _1626_v40;
        _1626_v40 = _dafny.Set.fromElements(_1625_v39);
        let _1627_v41;
        let _nw268 = Array((new BigNumber(5)).toNumber());
        _nw268[(_dafny.ZERO).toNumber()] = p1;
        _nw268[(_dafny.ONE).toNumber()] = ((p0) ? (new BigNumber((_dafny.MultiSet.fromElements(_1624_v38)).cardinality())) : (p1));
        _nw268[(new BigNumber(2)).toNumber()] = (_1620_i1).multipliedBy(_1620_i1);
        _nw268[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1624_v38, _dafny.Seq.update(_1624_v38, _module.__default.safeIndex(p1, new BigNumber((_1624_v38).length)), p0))).length);
        _nw268[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1626_v40).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-942))));
        _1627_v41 = _nw268;
        let _index285 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length));
        (_1627_v41)[_index285] = p1;
        let _index286 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length));
        let _rhs307 = _1621_v37;
        let _rhs308 = p1;
        let _lhs211 = _1627_v41;
        let _lhs212 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length));
        _1621_v37 = _rhs307;
        _lhs211[_lhs212] = _rhs308;
        let _1628_v42;
        _1628_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1620_i1,new BigNumber((_1625_v39).length));
        r0 = !(!((_module.__default.fm3(true, p0, globalState)).isLessThan(p1)) || (((((_1628_v42).contains(_1620_i1)) ? ((_1628_v42).get(_1620_i1)) : (p1))).isLessThan((_1627_v41)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length))])));
        let _1629_v43;
        _1629_v43 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1630_v44;
        _1630_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),!(p0));
        let _1631_v45;
        _1631_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1620_i1,(((_1630_v44).contains((_1627_v41)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length))])) ? ((_1630_v44).get((_1627_v41)[_module.__default.safeIndex(new BigNumber(214), new BigNumber((_1627_v41).length))])) : (p0)));
        let _1632_v46;
        let _1633_v47;
        let _out44;
        let _out45;
        let _outcollector19 = _module.__default.m0(_1621_v37, _1629_v43, p0, _module.__default.fm36(_1620_i1, p1, new BigNumber((_1631_v45).length), _1628_v42, globalState), globalState);
        _out44 = _outcollector19[0];
        _out45 = _outcollector19[1];
        _1632_v46 = _out44;
        _1633_v47 = _out45;
      }
      let _1634_v48;
      _1634_v48 = _module.D2.create_DC6(p2);
      let _1635_v49;
      _1635_v49 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1636_v50;
      let _nw269 = Array((new BigNumber(22)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = p1;
      _nw269[(_dafny.ONE).toNumber()] = p1;
      _nw269[(new BigNumber(2)).toNumber()] = p1;
      _nw269[(new BigNumber(3)).toNumber()] = new BigNumber((p2).length);
      _nw269[(new BigNumber(4)).toNumber()] = p1;
      _nw269[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw269[(new BigNumber(6)).toNumber()] = p1;
      _nw269[(new BigNumber(7)).toNumber()] = p1;
      _nw269[(new BigNumber(8)).toNumber()] = p1;
      _nw269[(new BigNumber(9)).toNumber()] = p1;
      _nw269[(new BigNumber(10)).toNumber()] = p1;
      _nw269[(new BigNumber(11)).toNumber()] = p1;
      _nw269[(new BigNumber(12)).toNumber()] = p1;
      _nw269[(new BigNumber(13)).toNumber()] = new BigNumber((_1635_v49).length);
      _nw269[(new BigNumber(14)).toNumber()] = p1;
      _nw269[(new BigNumber(15)).toNumber()] = new BigNumber(-618);
      _nw269[(new BigNumber(16)).toNumber()] = p1;
      _nw269[(new BigNumber(17)).toNumber()] = p1;
      _nw269[(new BigNumber(18)).toNumber()] = p1;
      _nw269[(new BigNumber(19)).toNumber()] = p1;
      _nw269[(new BigNumber(20)).toNumber()] = p1;
      _nw269[(new BigNumber(21)).toNumber()] = p1;
      _1636_v50 = _nw269;
      let _1637_v51;
      _1637_v51 = _dafny.MultiSet.fromElements(_1636_v50, _1636_v50);
      let _1638_v52;
      _1638_v52 = _dafny.Seq.of((((_1637_v51).contains(_1636_v50)) ? ((_1637_v51).get(_1636_v50)) : (p1)));
      let _hi10 = (_1638_v52)[_module.__default.safeIndex(p1, new BigNumber((_1638_v52).length))];
      for (let _1639_i3 = _module.__default.safeDivisionInt((_this).fm29(p1, (_1634_v48).dtor_cf6, globalState), p1); _1639_i3.isLessThan(_hi10); _1639_i3 = _1639_i3.plus(_dafny.ONE)) {
        let _1640_v53;
        _1640_v53 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1641_v54;
        _1641_v54 = _dafny.Set.fromElements(_1640_v53, _1640_v53);
        let _1642_v55;
        let _nw270 = Array((new BigNumber(5)).toNumber());
        _1642_v55 = _nw270;
        let _1643_v56;
        let _nw271 = new _module.C3();
        _nw271.__ctor(p0, p0);
        _1643_v56 = _nw271;
        let _index287 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_1642_v55).length));
        (_1642_v55)[_index287] = ((p0) ? (_1643_v56) : (_1643_v56));
        let _index288 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1636_v50).length));
        (_1636_v50)[_index288] = p1;
        let _1644_v57;
        let _nw272 = Array((new BigNumber(13)).toNumber());
        _nw272[(_dafny.ZERO).toNumber()] = _1640_v53;
        _nw272[(_dafny.ONE).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(2)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(3)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(4)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(5)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(6)).toNumber()] = ((p0) ? (_1640_v53) : (_1640_v53));
        _nw272[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
        _nw272[(new BigNumber(8)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(9)).toNumber()] = (((_1643_v56).f13) ? (_1640_v53) : (_1640_v53));
        _nw272[(new BigNumber(10)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(11)).toNumber()] = _1640_v53;
        _nw272[(new BigNumber(12)).toNumber()] = _1640_v53;
        _1644_v57 = _nw272;
        let _index289 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1644_v57).length));
        (_1644_v57)[_index289] = _1640_v53;
        let _1645_v58;
        let _init47 = function (_1646_i4) {
          return false;
        };
        let _nw273 = Array((new BigNumber(8)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw273.length); _i0_47++) {
          _nw273[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1645_v58 = _nw273;
        let _1647_v59;
        _1647_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1645_v58);
        let _index290 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_1642_v55).length));
        let _index291 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1636_v50).length));
        let _index292 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1644_v57).length));
        let _rhs309 = _1641_v54;
        let _rhs310 = _1643_v56;
        let _rhs311 = _1639_i3;
        let _rhs312 = _1640_v53;
        let _rhs313 = _1647_v59;
        let _lhs213 = _1642_v55;
        let _lhs214 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_1642_v55).length));
        let _lhs215 = _1636_v50;
        let _lhs216 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1636_v50).length));
        let _lhs217 = _1644_v57;
        let _lhs218 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_1644_v57).length));
        _1641_v54 = _rhs309;
        _lhs213[_lhs214] = _rhs310;
        _lhs215[_lhs216] = _rhs311;
        _lhs217[_lhs218] = _rhs312;
        _1647_v59 = _rhs313;
        let _1648_v60;
        let _nw274 = new _module.C2();
        _nw274.__ctor(_1640_v53, false);
        _1648_v60 = _nw274;
        let _index293 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1636_v50).length));
        (_1636_v50)[_index293] = _module.__default.safeDivisionInt(_1639_i3, _module.__default.safeModuloInt(p1, new BigNumber((p2).length)));
        r0 = false;
      }
      let _1649_v61;
      let _nw275 = Array((new BigNumber(14)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1649_v61 = _nw275;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1649_v61).length))) {
        let _1650_i5 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1650_i5)) && ((_1650_i5).isLessThan(new BigNumber((_1649_v61).length))))) {
          (_1649_v61)[(_1650_i5)] = ((p0) ? ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(p0, p0)).length)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), ((_1651_p1) => function (_1652_i6) {
            return _1651_p1;
          })(p1))))) : ((_dafny.MultiSet.fromElements(p1)).Intersect(_dafny.MultiSet.fromElements(p1, new BigNumber((_1638_v52).length)))));
        }
      }
      r0 = p0;
      let _1653_v62;
      _1653_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(691),p0);
      let _1654_v63;
      _1654_v63 = new _dafny.CodePoint('x'.codePointAt(0));
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(p2, p2), _dafny.Seq.Concat(p2, p2)), _module.__default.safeIndex(new BigNumber(((_1653_v62).update(p1, p0)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(p2, p2), _dafny.Seq.Concat(p2, p2))).length)), _1654_v63);
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1655_v0;
      _1655_v0 = _dafny.Seq.UnicodeFromString("ojxcdwwr");
      _1655_v0 = _1655_v0;
      let _1656_v1;
      _1656_v1 = true;
      let _1657_i0;
      _1657_i0 = _dafny.ZERO;
      L16: {
        while (_1656_v1) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1657_i0)) {
              break L16;
            }
            _1657_i0 = (_1657_i0).plus(_dafny.ONE);
            let _1658_v2;
            _1658_v2 = _dafny.Seq.of(_1656_v1);
            let _1659_v3;
            _1659_v3 = new BigNumber(715);
            let _1660_v4;
            _1660_v4 = _module.D11.create_DC20(_1656_v1, (_dafny.ZERO).minus(_1659_v3));
            if ((_1658_v2)[_module.__default.safeIndex((((_1660_v4).dtor_cf21) ? (_1659_v3) : (_1659_v3)), new BigNumber((_1658_v2).length))]) {
              let _1661_v5;
              let _nw276 = new _module.C4();
              _nw276.__ctor();
              _1661_v5 = _nw276;
              let _1662_v6;
              _1662_v6 = new _dafny.CodePoint('o'.codePointAt(0));
              _1662_v6 = _1662_v6;
              let _1663_v7;
              let _nw277 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1663_v7 = _nw277;
              let _index294 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1663_v7).length));
              (_1663_v7)[_index294] = _dafny.Seq.Concat(_1655_v0, _1655_v0);
              let _index295 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1663_v7).length));
              (_1663_v7)[_index295] = ((!(!(_1656_v1))) ? (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-456)), ((_1664_v6) => function (_1665_i1) {
                return _1664_v6;
              })(_1662_v6)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-219)), ((_1666_v6) => function (_1667_i2) {
                return _1666_v6;
              })(_1662_v6)))) : (_1655_v0));
              let _1668_v8;
              let _nw278 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
              _1668_v8 = _nw278;
              let _1669_v9;
              _1669_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,(_dafny.ZERO).minus(_1659_v3));
              let _index296 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1668_v8).length));
              (_1668_v8)[_index296] = _1669_v9;
              let _index297 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1668_v8).length));
              let _rhs314 = _1659_v3;
              let _rhs315 = !(!((_dafny.ZERO).minus((_1659_v3).multipliedBy(_1659_v3))).isEqualTo(_1659_v3));
              let _rhs316 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qn"), _dafny.Seq.UnicodeFromString("syeqbdag")), _dafny.Seq.Concat(_1655_v0, (_1663_v7)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_1663_v7).length))]));
              let _rhs317 = _1656_v1;
              let _rhs318 = (_1669_v9).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1659_v3)).Merge(_1669_v9));
              let _lhs219 = _1668_v8;
              let _lhs220 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1668_v8).length));
              _1659_v3 = _rhs314;
              _1656_v1 = _rhs315;
              _1655_v0 = _rhs316;
              _1656_v1 = _rhs317;
              _lhs219[_lhs220] = _rhs318;
              _1661_v5 = _1661_v5;
            } else {
              let _1670_v10;
              let _nw279 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1670_v10 = _nw279;
              let _1671_v11;
              let _nw280 = Array((new BigNumber(13)).toNumber());
              _nw280[(_dafny.ZERO).toNumber()] = _1670_v10;
              _nw280[(_dafny.ONE).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(2)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(3)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(4)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(5)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(6)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(7)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(8)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(9)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(10)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(11)).toNumber()] = _1670_v10;
              _nw280[(new BigNumber(12)).toNumber()] = _1670_v10;
              _1671_v11 = _nw280;
              let _1672_v12;
              let _nw281 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
              _1672_v12 = _nw281;
              let _1673_v13;
              _1673_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v11,_1672_v12);
              _1673_v13 = (_1673_v13).update(_1671_v11, _1672_v12);
              let _1674_v14;
              _1674_v14 = new _dafny.CodePoint('r'.codePointAt(0));
              let _1675_v15;
              _1675_v15 = _module.D11.create_DC19(_1674_v14);
              _1675_v15 = _1675_v15;
              let _1676_v16;
              let _nw282 = new _module.C4();
              _nw282.__ctor();
              _1676_v16 = _nw282;
              let _1677_v17;
              _1677_v17 = _dafny.Set.fromElements(_1656_v1, _1656_v1);
              r0 = _module.__default.safeDivisionInt(_1659_v3, _module.__default.safeDivisionInt(new BigNumber((_1677_v17).length), _1659_v3));
              let _index298 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1672_v12).length));
              (_1672_v12)[_index298] = _1659_v3;
              let _1678_v18;
              _1678_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1659_v3,_1656_v1);
              let _1679_v19;
              _1679_v19 = _1678_v18;
              let _1680_v20;
              let _nw283 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
              _1680_v20 = _nw283;
              let _index299 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1680_v20).length));
              (_1680_v20)[_index299] = _1658_v2;
              let _1681_v21;
              _1681_v21 = _dafny.Seq.of(new BigNumber(298), _1659_v3, _1659_v3);
              let _1682_v23;
              _1682_v23 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yhq"), _1655_v0);
              let _index300 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1672_v12).length));
              let _index301 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1680_v20).length));
              let _rhs319 = ((!(((_1656_v1) ? (_1656_v1) : (false)))) ? (_1659_v3) : (new BigNumber((_1681_v21).length)));
              let _rhs320 = _1679_v19;
              let _rhs321 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_1658_v2, _module.__default.safeIndex(new BigNumber((_1678_v18).length), new BigNumber((_1658_v2).length)), false), _module.__default.safeIndex(new BigNumber((function () {
                let _coll54 = new _dafny.Map();
                for (const _compr_54 of (function () {
                  let _coll55 = new _dafny.Set();
                  for (const _compr_55 of (_1682_v23).Elements) {
                    let _1683_v24 = _compr_55;
                    if (_dafny.Seq.contains(_1682_v23, _1683_v24)) {
                      _coll55.add(_1683_v24);
                    }
                  }
                  return _coll55;
                }()).Elements) {
                  let _1684_v22 = _compr_54;
                  if ((function () {
                    let _coll56 = new _dafny.Set();
                    for (const _compr_56 of (_1682_v23).Elements) {
                      let _1685_v24 = _compr_56;
                      if (_dafny.Seq.contains(_1682_v23, _1685_v24)) {
                        _coll56.add(_1685_v24);
                      }
                    }
                    return _coll56;
                  }()).contains(_1684_v22)) {
                    _coll54.push([_1684_v22,_1656_v1]);
                  }
                }
                return _coll54;
              }()).length), new BigNumber((_dafny.Seq.update(_1658_v2, _module.__default.safeIndex(new BigNumber((_1678_v18).length), new BigNumber((_1658_v2).length)), false)).length)), _1656_v1), _module.__default.safeIndex(new BigNumber((_1658_v2).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1658_v2, _module.__default.safeIndex(new BigNumber((_1678_v18).length), new BigNumber((_1658_v2).length)), false), _module.__default.safeIndex(new BigNumber((function () {
                let _coll57 = new _dafny.Map();
                for (const _compr_57 of (function () {
                  let _coll58 = new _dafny.Set();
                  for (const _compr_58 of (_1682_v23).Elements) {
                    let _1686_v24 = _compr_58;
                    if (_dafny.Seq.contains(_1682_v23, _1686_v24)) {
                      _coll58.add(_1686_v24);
                    }
                  }
                  return _coll58;
                }()).Elements) {
                  let _1687_v22 = _compr_57;
                  if ((function () {
                    let _coll59 = new _dafny.Set();
                    for (const _compr_59 of (_1682_v23).Elements) {
                      let _1688_v24 = _compr_59;
                      if (_dafny.Seq.contains(_1682_v23, _1688_v24)) {
                        _coll59.add(_1688_v24);
                      }
                    }
                    return _coll59;
                  }()).contains(_1687_v22)) {
                    _coll57.push([_1687_v22,_1656_v1]);
                  }
                }
                return _coll57;
              }()).length), new BigNumber((_dafny.Seq.update(_1658_v2, _module.__default.safeIndex(new BigNumber((_1678_v18).length), new BigNumber((_1658_v2).length)), false)).length)), _1656_v1)).length)), _1656_v1), _1658_v2);
              let _lhs221 = _1672_v12;
              let _lhs222 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1672_v12).length));
              let _lhs223 = _1680_v20;
              let _lhs224 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1680_v20).length));
              _lhs221[_lhs222] = _rhs319;
              _1679_v19 = _rhs320;
              _lhs223[_lhs224] = _rhs321;
            }
            let _1689_v25;
            let _nw284 = new _module.C4();
            _nw284.__ctor();
            _1689_v25 = _nw284;
            let _1690_v26;
            _1690_v26 = _module.D14.create_DC25(_1689_v25);
            let _pat_let_tv20 = _1689_v25;
            let _1691_v27;
            let _nw285 = Array((new BigNumber(13)).toNumber());
            _nw285[(_dafny.ZERO).toNumber()] = _1690_v26;
            _nw285[(_dafny.ONE).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(2)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(3)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(4)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(5)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(6)).toNumber()] = _module.D14.create_DC25(_1689_v25);
            _nw285[(new BigNumber(7)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(8)).toNumber()] = _module.D14.create_DC25(_1689_v25);
            _nw285[(new BigNumber(9)).toNumber()] = _module.D14.create_DC25(_1689_v25);
            _nw285[(new BigNumber(10)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(11)).toNumber()] = _1690_v26;
            _nw285[(new BigNumber(12)).toNumber()] = function (_pat_let22_0) {
              return function (_1692_dt__update__tmp_h0) {
                return function (_pat_let23_0) {
                  return function (_1693_dt__update_hcf31_h0) {
                    return _module.D14.create_DC25(_1693_dt__update_hcf31_h0);
                  }(_pat_let23_0);
                }(_pat_let_tv20);
              }(_pat_let22_0);
            }(_1690_v26);
            _1691_v27 = _nw285;
            let _pat_let_tv21 = _1689_v25;
            let _index302 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1691_v27).length));
            (_1691_v27)[_index302] = function (_pat_let24_0) {
              return function (_1694_dt__update__tmp_h1) {
                return function (_pat_let25_0) {
                  return function (_1695_dt__update_hcf31_h1) {
                    return _module.D14.create_DC25(_1695_dt__update_hcf31_h1);
                  }(_pat_let25_0);
                }(_pat_let_tv21);
              }(_pat_let24_0);
            }(_1690_v26);
            let _1696_v28;
            let _nw286 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
            _1696_v28 = _nw286;
            let _index303 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1696_v28).length));
            (_1696_v28)[_index303] = _dafny.Seq.of(_1659_v3);
            let _1697_v29;
            let _init48 = ((_1698_v2, _1699_v3) => function (_1700_i3) {
              return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1698_v2)).cardinality()),_1699_v3);
            })(_1658_v2, _1659_v3);
            let _nw287 = Array((new BigNumber(10)).toNumber());
            for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw287.length); _i0_48++) {
              _nw287[_i0_48] = _init48(new BigNumber(_i0_48));
            }
            _1697_v29 = _nw287;
            let _1701_v30;
            _1701_v30 = new _dafny.CodePoint('b'.codePointAt(0));
            let _1702_v31;
            _1702_v31 = _module.D11.create_DC19(_1701_v30);
            let _1703_v32;
            _1703_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(!(_1656_v1),(_1702_v31).dtor_cf20)).update(_1656_v1, _1701_v30),new BigNumber(88));
            let _1704_v33;
            let _nw288 = Array((new BigNumber(6)).toNumber()).fill(false);
            _1704_v33 = _nw288;
            let _1705_v34;
            _1705_v34 = _1704_v33;
            let _1706_v35;
            _1706_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1705_v34,_1656_v1);
            let _1707_v36;
            _1707_v36 = _dafny.Seq.of(new BigNumber((_1703_v32).length), new BigNumber((((_module.__default.fm0(_module.__default.fm0(!(_1656_v1), globalState), globalState)) ? (_1706_v35) : (_1706_v35))).length), _1659_v3, _1659_v3);
            let _index304 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1691_v27).length));
            let _index305 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1696_v28).length));
            let _rhs322 = _1690_v26;
            let _rhs323 = _1707_v36;
            let _rhs324 = !(_1656_v1) || (_1656_v1);
            let _rhs325 = new BigNumber(342);
            let _rhs326 = _1697_v29;
            let _lhs225 = _1691_v27;
            let _lhs226 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1691_v27).length));
            let _lhs227 = _1696_v28;
            let _lhs228 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1696_v28).length));
            let _lhs229 = globalState;
            _lhs225[_lhs226] = _rhs322;
            _lhs227[_lhs228] = _rhs323;
            _1656_v1 = _rhs324;
            _lhs229.f1 = _rhs325;
            _1697_v29 = _rhs326;
            let _1708_v37;
            _1708_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1655_v0);
            let _1709_v38;
            _1709_v38 = _dafny.MultiSet.fromElements((((_1708_v37).contains(_1656_v1)) ? ((_1708_v37).get(_1656_v1)) : (_1655_v0)));
            let _1710_v39;
            _1710_v39 = _dafny.Map.Empty.slice().updateUnsafe(true,_1659_v3);
            let _1711_v40;
            _1711_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1710_v39);
            let _1712_v42;
            _1712_v42 = _dafny.MultiSet.fromElements(_1656_v1, _1656_v1, _1656_v1, _1656_v1);
            let _1713_v43;
            _1713_v43 = _dafny.MultiSet.fromElements(_1659_v3);
            let _1714_v44;
            let _nw289 = Array((new BigNumber(17)).toNumber());
            _nw289[(_dafny.ZERO).toNumber()] = ((_1656_v1) ? (new BigNumber(-450)) : (_1659_v3));
            _nw289[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1711_v40).length));
            _nw289[(new BigNumber(2)).toNumber()] = (_1659_v3).plus(new BigNumber((_1655_v0).length));
            _nw289[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm35(_1655_v0, _dafny.MultiSet.FromArray((_1696_v28)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_1696_v28).length))]), _1701_v30, globalState)).length);
            _nw289[(new BigNumber(4)).toNumber()] = new BigNumber((_1655_v0).length);
            _nw289[(new BigNumber(5)).toNumber()] = new BigNumber((function () {
              let _coll60 = new _dafny.Set();
              for (const _compr_60 of _dafny.IntegerRange(new BigNumber(973), new BigNumber(8))) {
                let _1715_v41 = _compr_60;
                if (((new BigNumber(973)).isLessThanOrEqualTo(_1715_v41)) && ((_1715_v41).isLessThan(new BigNumber(8)))) {
                  _coll60.add((_1715_v41).minus(_1659_v3));
                }
              }
              return _coll60;
            }()).length);
            _nw289[(new BigNumber(6)).toNumber()] = ((_1656_v1) ? ((_dafny.ZERO).minus(new BigNumber((_1655_v0).length))) : (_1659_v3));
            _nw289[(new BigNumber(7)).toNumber()] = _1659_v3;
            _nw289[(new BigNumber(8)).toNumber()] = _1659_v3;
            _nw289[(new BigNumber(9)).toNumber()] = new BigNumber(927);
            _nw289[(new BigNumber(10)).toNumber()] = _1659_v3;
            _nw289[(new BigNumber(11)).toNumber()] = _module.__default.fm3(_1656_v1, _1656_v1, globalState);
            _nw289[(new BigNumber(12)).toNumber()] = _1659_v3;
            _nw289[(new BigNumber(13)).toNumber()] = _1659_v3;
            _nw289[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1655_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), ((_1716_v30) => function (_1717_i4) {
              return _1716_v30;
            })(_1701_v30)))).length);
            _nw289[(new BigNumber(15)).toNumber()] = new BigNumber((_1712_v42).cardinality());
            _nw289[(new BigNumber(16)).toNumber()] = new BigNumber((_1713_v43).cardinality());
            _1714_v44 = _nw289;
            let _1718_v45;
            _1718_v45 = _1714_v44;
            let _1719_v46;
            _1719_v46 = _module.D12.create_DC23(_1714_v44);
            let _1720_v47;
            _1720_v47 = _dafny.Seq.of(_1719_v46, _1719_v46);
            let _1721_v48;
            _1721_v48 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(264)).multipliedBy(new BigNumber((_1655_v0).length)),_dafny.Seq.update(_1720_v47, _module.__default.safeIndex(new BigNumber((_1707_v36).length), new BigNumber((_1720_v47).length)), _1719_v46));
            let _1722_v49;
            _1722_v49 = _dafny.Seq.of(_1655_v0, _dafny.Seq.update(_dafny.Seq.update(_1655_v0, _module.__default.safeIndex(_1659_v3, new BigNumber((_1655_v0).length)), _1701_v30), _module.__default.safeIndex(_1659_v3, new BigNumber((_dafny.Seq.update(_1655_v0, _module.__default.safeIndex(_1659_v3, new BigNumber((_1655_v0).length)), _1701_v30)).length)), _1701_v30));
            let _1723_v50;
            _1723_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(292),_dafny.Map.Empty.slice().updateUnsafe(_1659_v3,_1720_v47));
            let _rhs327 = (_1659_v3).multipliedBy(_1659_v3);
            let _rhs328 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1722_v49, _1722_v49));
            let _rhs329 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-808), _1659_v3))).minus((_1659_v3).minus(new BigNumber(393)));
            let _rhs330 = (_1718_v45);
            let _rhs331 = (_1721_v48).Merge((((_1723_v50).contains(_1659_v3)) ? ((_1723_v50).get(_1659_v3)) : (_1721_v48)));
            let _lhs230 = globalState;
            _1659_v3 = _rhs327;
            _1709_v38 = _rhs328;
            _lhs230.f1 = _rhs329;
            _1714_v44 = _rhs330;
            _1721_v48 = _rhs331;
            (globalState).f1 = _1659_v3;
          }
        }
      }
      if (_1656_v1) {
        let _1724_v53;
        let _init49 = ((_1725_v0, _1726_v1) => function (_1727_i5) {
          return (_1725_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
            let _coll61 = new _dafny.Map();
            for (const _compr_61 of (function () {
              let _coll62 = new _dafny.Set();
              for (const _compr_62 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-862)), function (_1728_i6) {
                return new _dafny.CodePoint('p'.codePointAt(0));
              }), _1725_v0)).Elements) {
                let _1729_v52 = _compr_62;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-862)), function (_1728_i6) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                }), _1725_v0), _1729_v52)) {
                  _coll62.add(_1729_v52);
                }
              }
              return _coll62;
            }()).Elements) {
              let _1730_v51 = _compr_61;
              if ((function () {
                let _coll63 = new _dafny.Set();
                for (const _compr_63 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-862)), function (_1728_i6) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                }), _1725_v0)).Elements) {
                  let _1731_v52 = _compr_63;
                  if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-862)), function (_1728_i6) {
                    return new _dafny.CodePoint('p'.codePointAt(0));
                  }), _1725_v0), _1731_v52)) {
                    _coll63.add(_1731_v52);
                  }
                }
                return _coll63;
              }()).contains(_1730_v51)) {
                _coll61.push([_1730_v51,_dafny.Map.Empty.slice().updateUnsafe(!(!(_1726_v1)),new BigNumber(989))]);
              }
            }
            return _coll61;
          }()).length))).length), new BigNumber((_1725_v0).length))];
        })(_1655_v0, _1656_v1);
        let _nw290 = Array((new BigNumber(23)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw290.length); _i0_49++) {
          _nw290[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1724_v53 = _nw290;
        _1724_v53 = _1724_v53;
        if (!(_1656_v1)) {
          let _1732_v54;
          _1732_v54 = new BigNumber(487);
          (globalState).f1 = (_1732_v54).multipliedBy(_1732_v54);
          let _1733_v55;
          _1733_v55 = _dafny.MultiSet.fromElements(!(_1656_v1));
          _1732_v54 = (((_1733_v55).contains(_1656_v1)) ? ((_1733_v55).get(_1656_v1)) : (new BigNumber(94)));
          let _1734_v56;
          let _init50 = ((_1735_v1) => function (_1736_i7) {
            return _1735_v1;
          })(_1656_v1);
          let _nw291 = Array((new BigNumber(17)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw291.length); _i0_50++) {
            _nw291[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1734_v56 = _nw291;
          let _index306 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1734_v56).length));
          (_1734_v56)[_index306] = _1656_v1;
          let _1737_v57;
          _1737_v57 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1738_v58;
          _1738_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1734_v56,new _dafny.CodePoint('c'.codePointAt(0)));
          let _1739_v59;
          _1739_v59 = _module.D0.create_DC2(_1738_v58, _1737_v57);
          let _1740_v60;
          _1740_v60 = _dafny.Seq.of(_1739_v59, _module.D0.create_DC2(_1738_v58, _1737_v57));
          let _1741_v61;
          _1741_v61 = _dafny.Seq.of(_1656_v1, _1656_v1);
          let _1742_v62;
          _1742_v62 = _dafny.Set.fromElements(_1656_v1);
          let _1743_v63;
          _1743_v63 = _dafny.Seq.of(_1732_v54);
          let _1744_v64;
          _1744_v64 = _dafny.MultiSet.fromElements(new BigNumber((_1743_v63).length));
          let _1745_v65;
          _1745_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1732_v54,(_this).fm4(_dafny.Seq.update(_dafny.Seq.update(_1741_v61, _module.__default.safeIndex(new BigNumber((_1742_v62).length), new BigNumber((_1741_v61).length)), _1656_v1), _module.__default.safeIndex(_1732_v54, new BigNumber((_dafny.Seq.update(_1741_v61, _module.__default.safeIndex(new BigNumber((_1742_v62).length), new BigNumber((_1741_v61).length)), _1656_v1)).length)), !(_1656_v1)), _1732_v54, (_dafny.ZERO).minus(_1732_v54), _1744_v64, globalState));
          let _index307 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1734_v56).length));
          let _rhs332 = _1656_v1;
          let _rhs333 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_1739_v59), _1740_v60);
          let _rhs334 = _module.__default.fm36(_1732_v54, _1732_v54, (_1732_v54).minus(_1732_v54), _1745_v65, globalState);
          let _lhs231 = _1734_v56;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_1734_v56).length));
          _lhs231[_lhs232] = _rhs332;
          _1656_v1 = _rhs333;
          _1737_v57 = _rhs334;
          _1656_v1 = _1656_v1;
          _1732_v54 = _1732_v54;
        } else {
          let _1746_v66;
          _1746_v66 = _dafny.Seq.of(_module.__default.fm3(_1656_v1, _1656_v1, globalState));
          let _1747_v67;
          _1747_v67 = new BigNumber(-506);
          let _1748_v68;
          let _nw292 = Array((new BigNumber(14)).toNumber()).fill([]);
          _1748_v68 = _nw292;
          let _1749_v69;
          _1749_v69 = _dafny.Map.Empty.slice().updateUnsafe((_1746_v66)[_module.__default.safeIndex(_1747_v67, new BigNumber((_1746_v66).length))],_1748_v68);
          _1749_v69 = (_1749_v69).update(_1747_v67, _1748_v68);
          let _1750_v70;
          let _init51 = ((_1751_v67) => function (_1752_i8) {
            return (_dafny.MultiSet.fromElements(_1751_v67, _1751_v67)).Difference(_dafny.MultiSet.fromElements(_1751_v67, new BigNumber(459)));
          })(_1747_v67);
          let _nw293 = Array((new BigNumber(16)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw293.length); _i0_51++) {
            _nw293[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1750_v70 = _nw293;
          let _1753_v71;
          _1753_v71 = _dafny.MultiSet.fromElements(_1747_v67);
          let _index308 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1750_v70).length));
          (_1750_v70)[_index308] = (_1753_v71).update(_1747_v67, _module.__default.abs(_1747_v67));
          let _index309 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1750_v70).length));
          (_1750_v70)[_index309] = _1753_v71;
          let _1754_v72;
          _1754_v72 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1755_v73;
          let _nw294 = new _module.C5();
          _nw294.__ctor((_dafny.ZERO).minus(_1747_v67), _1754_v72, _1656_v1);
          _1755_v73 = _nw294;
          let _1756_v74;
          let _nw295 = Array((new BigNumber(2)).toNumber());
          _nw295[(_dafny.ZERO).toNumber()] = _1755_v73;
          _nw295[(_dafny.ONE).toNumber()] = _1755_v73;
          _1756_v74 = _nw295;
          let _index310 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1756_v74).length));
          (_1756_v74)[_index310] = _1755_v73;
          let _index311 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1756_v74).length));
          (_1756_v74)[_index311] = _1755_v73;
          _1747_v67 = (_dafny.ZERO).minus((new BigNumber(-624)).multipliedBy(_module.__default.safeDivisionInt(_1747_v67, _1747_v67)));
          _1656_v1 = (_1755_v73).f6;
        }
        if (_1656_v1) {
          let _1757_v75;
          _1757_v75 = _module.D2.create_DC6(_1655_v0);
          let _pat_let_tv22 = _1655_v0;
          _1757_v75 = ((_1656_v1) ? (((_1656_v1) ? (_1757_v75) : (function (_pat_let26_0) {
            return function (_1758_dt__update__tmp_h3) {
              return function (_pat_let27_0) {
                return function (_1759_dt__update_hcf6_h0) {
                  return _module.D2.create_DC6(_1759_dt__update_hcf6_h0);
                }(_pat_let27_0);
              }(_pat_let_tv22);
            }(_pat_let26_0);
          }(_1757_v75)))) : (_1757_v75));
          _1656_v1 = _1656_v1;
          let _1760_v76;
          let _nw296 = Array((new BigNumber(8)).toNumber());
          _nw296[(_dafny.ZERO).toNumber()] = _1656_v1;
          _nw296[(_dafny.ONE).toNumber()] = _1656_v1;
          _nw296[(new BigNumber(2)).toNumber()] = _1656_v1;
          _nw296[(new BigNumber(3)).toNumber()] = _1656_v1;
          _nw296[(new BigNumber(4)).toNumber()] = _1656_v1;
          _nw296[(new BigNumber(5)).toNumber()] = _1656_v1;
          _nw296[(new BigNumber(6)).toNumber()] = false;
          _nw296[(new BigNumber(7)).toNumber()] = _1656_v1;
          _1760_v76 = _nw296;
          let _1761_v77;
          _1761_v77 = _1760_v76;
          let _1762_v78;
          _1762_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1761_v77,_1760_v76);
          let _1763_v79;
          _1763_v79 = _1762_v78;
          let _1764_v80;
          let _nw297 = Array((new BigNumber(13)).toNumber());
          _nw297[(_dafny.ZERO).toNumber()] = _1763_v79;
          _nw297[(_dafny.ONE).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(2)).toNumber()] = _1762_v78;
          _nw297[(new BigNumber(3)).toNumber()] = _1762_v78;
          _nw297[(new BigNumber(4)).toNumber()] = ((_1656_v1) ? (_1763_v79) : (_1763_v79));
          _nw297[(new BigNumber(5)).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(6)).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(7)).toNumber()] = _1762_v78;
          _nw297[(new BigNumber(8)).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(9)).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(10)).toNumber()] = _1763_v79;
          _nw297[(new BigNumber(11)).toNumber()] = _1762_v78;
          _nw297[(new BigNumber(12)).toNumber()] = _1762_v78;
          _1764_v80 = _nw297;
          let _index312 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1764_v80).length));
          (_1764_v80)[_index312] = _1762_v78;
          let _index313 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1764_v80).length));
          (_1764_v80)[_index313] = _1763_v79;
          let _1765_v81;
          _1765_v81 = new BigNumber(958);
          let _1766_v82;
          _1766_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm29(_1765_v81, _1655_v0, globalState),_1760_v76);
          _1766_v82 = (_1766_v82).update((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((function () {
            let _coll64 = new _dafny.Set();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(342), new BigNumber(17))) {
              let _1767_v83 = _compr_64;
              if (((new BigNumber(342)).isLessThanOrEqualTo(_1767_v83)) && ((_1767_v83).isLessThan(new BigNumber(17)))) {
                _coll64.add(_module.__default.safeDivisionInt(_1767_v83, (_dafny.ZERO).minus(_1765_v81)));
              }
            }
            return _coll64;
          }()).length), _1765_v81)), _1760_v76);
          let _1768_v84;
          _1768_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1765_v81,_1656_v1);
          let _index314 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_1760_v76).length));
          (_1760_v76)[_index314] = (((_1768_v84).contains(_1765_v81)) ? ((_1768_v84).get(_1765_v81)) : (_1656_v1));
          let _index315 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_1760_v76).length));
          (_1760_v76)[_index315] = false;
        } else {
          let _1769_v85;
          _1769_v85 = new BigNumber(168);
          let _1770_v86;
          _1770_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(268),_1769_v85);
          let _1771_v87;
          _1771_v87 = _1770_v86;
          let _1772_v88;
          _1772_v88 = _dafny.Seq.of(_1656_v1);
          let _1773_v89;
          _1773_v89 = _dafny.Seq.of(_1772_v88);
          let _1774_v90;
          _1774_v90 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1775_v91;
          _1775_v91 = _module.D12.create_DC22(_1771_v87, _1773_v89, _1774_v90, _1656_v1, _1655_v0);
          let _1776_v92;
          _1776_v92 = _dafny.MultiSet.fromElements(_1775_v91);
          let _1777_v93;
          _1777_v93 = _dafny.Seq.of((new BigNumber((_1776_v92).cardinality())).isLessThanOrEqualTo(_1769_v85));
          _1656_v1 = (_1777_v93)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_1777_v93).length))];
          let _1778_v94;
          let _nw298 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _1778_v94 = _nw298;
          let _index316 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1778_v94).length));
          (_1778_v94)[_index316] = new BigNumber(824);
          let _index317 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1778_v94).length));
          (_1778_v94)[_index317] = new BigNumber((_1655_v0).length);
          let _1779_v95;
          _1779_v95 = _dafny.MultiSet.fromElements(_1656_v1);
          let _1780_v96;
          _1780_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1779_v95);
          let _1781_v97;
          let _nw299 = new _module.C2();
          _nw299.__ctor(_1774_v90, _1656_v1);
          _1781_v97 = _nw299;
          let _1782_v98;
          _1782_v98 = _dafny.Map.Empty.slice().updateUnsafe(((((_1780_v96).contains(_1656_v1)) ? ((_1780_v96).get(_1656_v1)) : (_1779_v95))).IsProperSubsetOf(_1779_v95),_1781_v97);
          let _1783_v99;
          _1783_v99 = _dafny.Seq.of(_1769_v85);
          _1782_v98 = (_1782_v98).update(!((_1769_v85).isLessThanOrEqualTo((_1783_v99)[_module.__default.safeIndex((_1778_v94)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1778_v94).length))], new BigNumber((_1783_v99).length))])), _1781_v97);
          _1656_v1 = (!(!(!(_1656_v1)) || (_1656_v1))) && (_1656_v1);
          let _1784_v100;
          let _nw300 = new _module.C2();
          _nw300.__ctor(_1774_v90, (_1656_v1) || (_1656_v1));
          _1784_v100 = _nw300;
        }
        let _1785_v101;
        _1785_v101 = new BigNumber(904);
        let _1786_v102;
        _1786_v102 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1785_v101),_1656_v1);
        let _1787_v103;
        _1787_v103 = _dafny.Seq.of(_1785_v101, new BigNumber((_1786_v102).length));
        let _1788_v104;
        _1788_v104 = _dafny.Seq.of(_1787_v103);
        let _1789_v105;
        let _nw301 = new _module.C0();
        _nw301.__ctor(true, _1788_v104);
        _1789_v105 = _nw301;
        let _1790_v106;
        _1790_v106 = _dafny.MultiSet.fromElements(_1789_v105, _1789_v105);
        let _rhs335 = _module.__default.safeDivisionInt((((_1790_v106).contains(_1789_v105)) ? ((_1790_v106).get(_1789_v105)) : ((_dafny.ZERO).minus(_1785_v101))), _1785_v101);
        let _rhs336 = (_module.__default.safeDivisionInt(_1785_v101, _1785_v101)).isEqualTo(_1785_v101);
        let _lhs233 = globalState;
        _lhs233.f1 = _rhs335;
        _1656_v1 = _rhs336;
        let _1791_v107;
        _1791_v107 = _dafny.Seq.of(_1656_v1);
        let _1792_v108;
        _1792_v108 = _dafny.MultiSet.fromElements(_1791_v107);
        let _1793_v109;
        _1793_v109 = _dafny.Seq.of(_1791_v107);
        _1792_v108 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1793_v109, _1793_v109));
      } else {
        if (_1656_v1) {
          let _1794_v110;
          let _nw302 = Array((new BigNumber(2)).toNumber());
          _nw302[(_dafny.ZERO).toNumber()] = _1656_v1;
          _nw302[(_dafny.ONE).toNumber()] = _1656_v1;
          _1794_v110 = _nw302;
          let _1795_v111;
          _1795_v111 = _dafny.Map.Empty.slice().updateUnsafe(_1794_v110,new BigNumber(-376));
          let _1796_v112;
          _1796_v112 = _dafny.Seq.of(_1794_v110);
          let _1797_v113;
          _1797_v113 = new BigNumber(688);
          let _1798_v114;
          let _nw303 = new _module.C2();
          _nw303.__ctor(new _dafny.CodePoint('l'.codePointAt(0)), _1656_v1);
          _1798_v114 = _nw303;
          let _1799_v115;
          _1799_v115 = _dafny.MultiSet.fromElements(_1798_v114, _1798_v114);
          let _1800_v116;
          _1800_v116 = _dafny.Seq.of(_1795_v111, _1795_v111);
          let _1801_v117;
          _1801_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1797_v113,_1794_v110);
          let _1802_v118;
          _1802_v118 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1803_v119;
          _1803_v119 = _dafny.Seq.of(_1797_v113);
          let _1804_v120;
          _1804_v120 = _dafny.Map.Empty.slice().updateUnsafe(false,_1803_v119);
          let _1805_v121;
          _1805_v121 = _dafny.Set.fromElements(false);
          let _1806_v122;
          let _nw304 = Array((new BigNumber(22)).toNumber());
          _nw304[(_dafny.ZERO).toNumber()] = _1797_v113;
          _nw304[(_dafny.ONE).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(2)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(3)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(4)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(5)).toNumber()] = new BigNumber((_1655_v0).length);
          _nw304[(new BigNumber(6)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(7)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(8)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(9)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(10)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(11)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(12)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(13)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(14)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(15)).toNumber()] = new BigNumber((_1804_v120).length);
          _nw304[(new BigNumber(16)).toNumber()] = _1797_v113;
          _nw304[(new BigNumber(17)).toNumber()] = new BigNumber(105);
          _nw304[(new BigNumber(18)).toNumber()] = new BigNumber((_1805_v121).length);
          _nw304[(new BigNumber(19)).toNumber()] = new BigNumber(153);
          _nw304[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("toyypgf")).length);
          _nw304[(new BigNumber(21)).toNumber()] = _1797_v113;
          _1806_v122 = _nw304;
          let _1807_v123;
          _1807_v123 = _dafny.Map.Empty.slice().updateUnsafe(_1802_v118,_1806_v122);
          let _1808_v124;
          let _nw305 = Array((new BigNumber(24)).toNumber());
          _nw305[(_dafny.ZERO).toNumber()] = (_1795_v111).Merge(_1795_v111);
          _nw305[(_dafny.ONE).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(2)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1796_v112)[_module.__default.safeIndex(_1797_v113, new BigNumber((_1796_v112).length))],_1797_v113)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1794_v110,new BigNumber((_1799_v115).cardinality())));
          _nw305[(new BigNumber(4)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(5)).toNumber()] = (_1795_v111).update(_1794_v110, new BigNumber(948));
          _nw305[(new BigNumber(6)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(7)).toNumber()] = (_1800_v116)[_module.__default.safeIndex(_1797_v113, new BigNumber((_1800_v116).length))];
          _nw305[(new BigNumber(8)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(9)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_1801_v117).contains(new BigNumber((_1655_v0).length))) ? ((_1801_v117).get(new BigNumber((_1655_v0).length))) : (_1794_v110)),_1797_v113);
          _nw305[(new BigNumber(11)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1794_v110,_1797_v113);
          _nw305[(new BigNumber(13)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(14)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1794_v110,new BigNumber((_1807_v123).length));
          _nw305[(new BigNumber(16)).toNumber()] = (_1795_v111).Merge(_1795_v111);
          _nw305[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1794_v110,_1797_v113);
          _nw305[(new BigNumber(18)).toNumber()] = (_1795_v111).Merge(_1795_v111);
          _nw305[(new BigNumber(19)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(20)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(21)).toNumber()] = (_1795_v111).Merge(_1795_v111);
          _nw305[(new BigNumber(22)).toNumber()] = _1795_v111;
          _nw305[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1794_v110,_1797_v113);
          _1808_v124 = _nw305;
          let _index318 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1808_v124).length));
          (_1808_v124)[_index318] = _1795_v111;
          let _1809_v125;
          _1809_v125 = _dafny.MultiSet.fromElements(_1797_v113, _1797_v113);
          let _index319 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_1808_v124).length));
          (_1808_v124)[_index319] = ((_1795_v111).update(_1794_v110, _1797_v113)).update(_1794_v110, _module.__default.safeModuloInt(_1797_v113, new BigNumber((_1809_v125).cardinality())));
          let _rhs337 = _1802_v118;
          let _rhs338 = ((((_1656_v1) ? (false) : (_1656_v1))) ? (_module.__default.safeDivisionInt(new BigNumber(811), _1797_v113)) : (_1797_v113));
          let _lhs234 = globalState;
          _1802_v118 = _rhs337;
          _lhs234.f1 = _rhs338;
          let _1810_v126;
          _1810_v126 = _module.D6.create_DC13(_1803_v119);
          _1810_v126 = _1810_v126;
          let _1811_v127;
          _1811_v127 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1797_v113);
          let _1812_v128;
          _1812_v128 = (_dafny.Map.Empty.slice().updateUnsafe(_1797_v113,new BigNumber((_1811_v127).length))).update(_1797_v113, new BigNumber(-464));
          let _1813_v129;
          _1813_v129 = _dafny.Seq.of(_1656_v1);
          let _1814_v130;
          _1814_v130 = _dafny.Seq.of(_1813_v129);
          let _1815_v131;
          _1815_v131 = _module.D12.create_DC22(_1812_v128, _1814_v130, _1802_v118, _1656_v1, _1655_v0);
          let _1816_v132;
          _1816_v132 = _dafny.Set.fromElements(_1794_v110);
          let _1817_v133;
          let _nw306 = new _module.C2();
          _nw306.__ctor(((_1656_v1) ? (_1802_v118) : ((_1815_v131).dtor_cf26)), (_1816_v132).contains(_1794_v110));
          _1817_v133 = _nw306;
          let _1818_v134;
          _1818_v134 = _dafny.Set.fromElements(_module.D11.create_DC19(new _dafny.CodePoint('l'.codePointAt(0))));
          let _index320 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1794_v110).length));
          (_1794_v110)[_index320] = (_1813_v129)[_module.__default.safeIndex(_1797_v113, new BigNumber((_1813_v129).length))];
          let _1819_v135;
          _1819_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1797_v113,(_1818_v134).Union(_1818_v134));
          let _1820_v136;
          _1820_v136 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(137),_1656_v1);
          let _1821_v137;
          _1821_v137 = _dafny.Map.Empty.slice().updateUnsafe(_1820_v136,_1797_v113);
          let _index321 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1794_v110).length));
          let _rhs339 = _1656_v1;
          let _rhs340 = (((_1819_v135).contains(_module.__default.safeDivisionInt((((_1821_v137).contains(_1820_v136)) ? ((_1821_v137).get(_1820_v136)) : (new BigNumber((_dafny.MultiSet.fromElements(_1794_v110)).cardinality()))), _1797_v113))) ? ((_1819_v135).get(_module.__default.safeDivisionInt((((_1821_v137).contains(_1820_v136)) ? ((_1821_v137).get(_1820_v136)) : (new BigNumber((_dafny.MultiSet.fromElements(_1794_v110)).cardinality()))), _1797_v113))) : (_1818_v134));
          let _rhs341 = _1656_v1;
          let _lhs235 = _1794_v110;
          let _lhs236 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_1794_v110).length));
          _1656_v1 = _rhs339;
          _1818_v134 = _rhs340;
          _lhs235[_lhs236] = _rhs341;
        } else {
          let _1822_v138;
          _1822_v138 = new BigNumber(-987);
          (globalState).f1 = (_dafny.ZERO).minus(_1822_v138);
          let _1823_v139;
          _1823_v139 = _dafny.Seq.of(_1822_v138, _1822_v138);
          _1656_v1 = (_this).fm6(_1823_v139, globalState);
          let _1824_v140;
          let _nw307 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1824_v140 = _nw307;
          let _1825_v141;
          let _nw308 = Array((new BigNumber(13)).toNumber());
          _nw308[(_dafny.ZERO).toNumber()] = _1824_v140;
          _nw308[(_dafny.ONE).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(2)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(3)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(4)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(5)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(6)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(7)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(8)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(9)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(10)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(11)).toNumber()] = _1824_v140;
          _nw308[(new BigNumber(12)).toNumber()] = _1824_v140;
          _1825_v141 = _nw308;
          let _rhs342 = _1825_v141;
          let _rhs343 = _1656_v1;
          _1825_v141 = _rhs342;
          _1656_v1 = _rhs343;
          (globalState).f1 = _1822_v138;
          (_this).m9(_dafny.Seq.UnicodeFromString("cdpjpt"), globalState);
        }
        let _1826_v142;
        _1826_v142 = new BigNumber(790);
        let _1827_v143;
        _1827_v143 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(763),_1826_v142);
        (globalState).f1 = (((_1827_v143).contains(_1826_v142)) ? ((_1827_v143).get(_1826_v142)) : ((_1826_v142).minus(_1826_v142)));
        let _1828_v144;
        _1828_v144 = _dafny.Seq.of(_1656_v1);
        let _1829_v145;
        _1829_v145 = _dafny.MultiSet.fromElements(_1828_v144);
        _1829_v145 = _dafny.MultiSet.fromElements(_1828_v144);
        let _1830_v146;
        _1830_v146 = _dafny.Set.fromElements(_1656_v1);
        if ((_1830_v146).IsSubsetOf(((_1656_v1) ? (_dafny.Set.fromElements(_1656_v1, _1656_v1, _1656_v1, _1656_v1)) : (_1830_v146)))) {
          let _1831_v147;
          let _nw309 = Array((new BigNumber(5)).toNumber()).fill(false);
          _1831_v147 = _nw309;
          let _1832_v148;
          _1832_v148 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1826_v142);
          let _1833_v149;
          _1833_v149 = _module.D15.create_DC28(_1656_v1, _1656_v1);
          let _pat_let_tv23 = _1656_v1;
          let _pat_let_tv24 = globalState;
          let _pat_let_tv25 = _1827_v143;
          let _rhs344 = (function (_pat_let28_0) {
            return function (_1834_dt__update__tmp_h4) {
              return function (_pat_let29_0) {
                return function (_1835_dt__update_hcf36_h0) {
                  return _module.D15.create_DC28(_1835_dt__update_hcf36_h0, (_1834_dt__update__tmp_h4).dtor_cf37);
                }(_pat_let29_0);
              }((_module.D11.create_DC20(_module.__default.fm0(_pat_let_tv23, _pat_let_tv24), new BigNumber((_pat_let_tv25).length))).dtor_cf21);
            }(_pat_let28_0);
          }(_1833_v149)).dtor_cf37;
          let _rhs345 = _1826_v142;
          let _rhs346 = (_1830_v146).Intersect(_1830_v146);
          let _rhs347 = _1831_v147;
          let _rhs348 = _1832_v148;
          let _lhs237 = globalState;
          _1656_v1 = _rhs344;
          _lhs237.f1 = _rhs345;
          _1830_v146 = _rhs346;
          _1831_v147 = _rhs347;
          _1832_v148 = _rhs348;
          let _1836_v150;
          _1836_v150 = _dafny.MultiSet.fromElements(_1656_v1);
          let _1837_v151;
          _1837_v151 = _dafny.Seq.of(_1826_v142);
          let _1838_v152;
          let _nw310 = Array((new BigNumber(19)).toNumber());
          _nw310[(_dafny.ZERO).toNumber()] = _1826_v142;
          _nw310[(_dafny.ONE).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(2)).toNumber()] = (_this).fm7(_1656_v1, globalState);
          _nw310[(new BigNumber(3)).toNumber()] = new BigNumber(413);
          _nw310[(new BigNumber(4)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(5)).toNumber()] = new BigNumber(342);
          _nw310[(new BigNumber(6)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(7)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(8)).toNumber()] = new BigNumber((_1836_v150).cardinality());
          _nw310[(new BigNumber(9)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(10)).toNumber()] = new BigNumber((_1837_v151).length);
          _nw310[(new BigNumber(11)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(12)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(13)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(14)).toNumber()] = new BigNumber(166);
          _nw310[(new BigNumber(15)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(16)).toNumber()] = (((_1832_v148).contains(true)) ? ((_1832_v148).get(true)) : (_1826_v142));
          _nw310[(new BigNumber(17)).toNumber()] = _1826_v142;
          _nw310[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_1826_v142);
          _1838_v152 = _nw310;
          let _1839_v153;
          _1839_v153 = _dafny.Seq.of(_1838_v152, _1838_v152, _1838_v152);
          let _1840_v154;
          _1840_v154 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1826_v142),(_1839_v153)[_module.__default.safeIndex(_1826_v142, new BigNumber((_1839_v153).length))]);
          let _rhs349 = ((_1656_v1) ? (((_1656_v1) ? (new BigNumber((_1655_v0).length)) : (_1826_v142))) : ((_1826_v142).minus(_1826_v142)));
          let _rhs350 = _1840_v154;
          let _lhs238 = globalState;
          _lhs238.f1 = _rhs349;
          _1840_v154 = _rhs350;
          let _1841_v155;
          _1841_v155 = _module.D0.create_DC1();
          let _1842_v156;
          _1842_v156 = _module.D0.create_DC3(_1841_v155);
          let _1843_v157;
          _1843_v157 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1844_v158;
          let _nw311 = new _module.C2();
          _nw311.__ctor(_1843_v157, _module.__default.fm0(_1656_v1, globalState));
          _1844_v158 = _nw311;
          let _1845_v159;
          _1845_v159 = _dafny.Map.Empty.slice().updateUnsafe(_1842_v156,_1844_v158);
          _1845_v159 = (_1845_v159).update(_1842_v156, _1844_v158);
          let _1846_v160;
          _1846_v160 = _module.D11.create_DC20((_1844_v158).fm6(_1837_v151, globalState), _1826_v142);
          _1832_v148 = (_1832_v148).update((_1846_v160).dtor_cf21, _1826_v142);
          _1838_v152 = _1838_v152;
        } else {
          _1656_v1 = !(_1656_v1) || (_1656_v1);
          let _1847_v161;
          _1847_v161 = _dafny.Map.Empty.slice().updateUnsafe(_1655_v0,_1656_v1);
          _1847_v161 = (_1847_v161).update(_dafny.Seq.Concat(_1655_v0, _1655_v0), !(_1656_v1));
          let _1848_v162;
          let _nw312 = new _module.C1();
          _nw312.__ctor();
          _1848_v162 = _nw312;
          let _1849_v163;
          _1849_v163 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1656_v1);
          let _1850_v164;
          _1850_v164 = _module.D12.create_DC21(_1849_v163);
          let _1851_v165;
          _1851_v165 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v1,_1850_v164);
          _1851_v165 = _1851_v165;
          _1656_v1 = _1656_v1;
        }
        (globalState).f1 = _module.__default.safeModuloInt(_1826_v142, _1826_v142);
      }
      let _1852_v166;
      _1852_v166 = new BigNumber(892);
      let _1853_v167;
      _1853_v167 = _module.D1.create_DC4(_1852_v166);
      let _1854_v168;
      _1854_v168 = _module.D2.create_DC8(_1656_v1);
      let _pat_let_tv26 = _1852_v166;
      _1853_v167 = function (_pat_let30_0) {
        return function (_1855_dt__update__tmp_h5) {
          return function (_pat_let31_0) {
            return function (_1856_dt__update_hcf4_h0) {
              return _module.D1.create_DC4(_1856_dt__update_hcf4_h0);
            }(_pat_let31_0);
          }(_pat_let_tv26);
        }(_pat_let30_0);
      }((_this).fm5(_1854_v168, globalState));
      let _1857_v169;
      _1857_v169 = _dafny.Seq.of(new BigNumber(-263));
      let _1858_v170;
      _1858_v170 = _dafny.Seq.of(_1857_v169, _1857_v169, _1857_v169);
      let _1859_v171;
      _1859_v171 = _dafny.Seq.of(_1857_v169, _1857_v169, _1857_v169, _1857_v169, (_1858_v170)[_module.__default.safeIndex(_1852_v166, new BigNumber((_1858_v170).length))]);
      let _1860_v172;
      let _nw313 = new _module.C0();
      _nw313.__ctor(((_1656_v1) ? (_1656_v1) : (!(_1656_v1))), _dafny.Seq.Concat(_1858_v170, _1858_v170));
      _1860_v172 = _nw313;
      let _1861_v173;
      _1861_v173 = _dafny.Map.Empty.slice().updateUnsafe(!((_1860_v172).f11),_1852_v166);
      let _1862_v174;
      _1862_v174 = _dafny.MultiSet.fromElements(_1656_v1);
      let _1863_v175;
      _1863_v175 = _dafny.Seq.of((_1854_v168).dtor_cf9, _1656_v1, (_1860_v172).f11);
      let _1864_v176;
      _1864_v176 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-512),new BigNumber((_dafny.Seq.of(_1656_v1, false)).length));
      let _1865_v177;
      _1865_v177 = _dafny.Set.fromElements((_1860_v172).f11, (_1860_v172).f11);
      let _1866_v178;
      let _nw314 = Array((new BigNumber(27)).toNumber());
      _nw314[(_dafny.ZERO).toNumber()] = _1852_v166;
      _nw314[(_dafny.ONE).toNumber()] = new BigNumber(((_1861_v173).Merge(_module.__default.fm57((_dafny.ZERO).minus(_1852_v166), new BigNumber((_1862_v174).cardinality()), globalState))).length);
      _nw314[(new BigNumber(2)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(3)).toNumber()] = (_1852_v166).plus(_1852_v166);
      _nw314[(new BigNumber(4)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_1852_v166);
      _nw314[(new BigNumber(6)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(7)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(8)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("pmtyrbao")).length), _1852_v166);
      _nw314[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_1857_v169).length)).minus(_1852_v166));
      _nw314[(new BigNumber(11)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(12)).toNumber()] = new BigNumber((_1863_v175).length);
      _nw314[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_1852_v166), _1852_v166);
      _nw314[(new BigNumber(14)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(15)).toNumber()] = (((_1864_v176).contains((((_1864_v176).contains(_1852_v166)) ? ((_1864_v176).get(_1852_v166)) : (_1852_v166)))) ? ((_1864_v176).get((((_1864_v176).contains(_1852_v166)) ? ((_1864_v176).get(_1852_v166)) : (_1852_v166)))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1852_v166,_1852_v166)).length)));
      _nw314[(new BigNumber(16)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(17)).toNumber()] = new BigNumber((_1857_v169).length);
      _nw314[(new BigNumber(18)).toNumber()] = ((false) ? (_1852_v166) : (new BigNumber((_1863_v175).length)));
      _nw314[(new BigNumber(19)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(20)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(21)).toNumber()] = new BigNumber((_1655_v0).length);
      _nw314[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(_1852_v166, _1852_v166);
      _nw314[(new BigNumber(23)).toNumber()] = new BigNumber((((false) ? (_1865_v177) : (_1865_v177))).length);
      _nw314[(new BigNumber(24)).toNumber()] = (_1852_v166).plus(_1852_v166);
      _nw314[(new BigNumber(25)).toNumber()] = _1852_v166;
      _nw314[(new BigNumber(26)).toNumber()] = _1852_v166;
      _1866_v178 = _nw314;
      let _rhs351 = _1866_v178;
      let _rhs352 = _1656_v1;
      _1866_v178 = _rhs351;
      _1656_v1 = _rhs352;
      r0 = _1852_v166;
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _1867_v0;
      _1867_v0 = new BigNumber(442);
      let _1868_v1;
      _1868_v1 = _dafny.Seq.of(_module.__default.safeModuloInt(_module.__default.fm3(p0, _module.__default.fm0(p0, globalState), globalState), _1867_v0));
      _1868_v1 = _1868_v1;
      let _1869_v2;
      _1869_v2 = _dafny.Seq.UnicodeFromString("yqnx");
      _1867_v0 = _module.__default.safeModuloInt(new BigNumber((_1869_v2).length), _1867_v0);
      let _1870_i0;
      _1870_i0 = _dafny.ZERO;
      L17: {
        while (p0) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1870_i0)) {
              break L17;
            }
            _1870_i0 = (_1870_i0).plus(_dafny.ONE);
            let _1871_v3;
            _1871_v3 = true;
            let _1872_v4;
            let _nw315 = Array((new BigNumber(5)).toNumber());
            _nw315[(_dafny.ZERO).toNumber()] = p0;
            _nw315[(_dafny.ONE).toNumber()] = (p0) && (p2);
            _nw315[(new BigNumber(2)).toNumber()] = _module.__default.fm0(p2, globalState);
            _nw315[(new BigNumber(3)).toNumber()] = !(_1871_v3);
            _nw315[(new BigNumber(4)).toNumber()] = p0;
            _1872_v4 = _nw315;
            let _1873_v5;
            _1873_v5 = _dafny.Seq.of(!(p0));
            let _index322 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1872_v4).length));
            (_1872_v4)[_index322] = (_1873_v5)[_module.__default.safeIndex(_1867_v0, new BigNumber((_1873_v5).length))];
            let _index323 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1872_v4).length));
            let _rhs353 = (_1873_v5)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_1867_v0, _1867_v0), new BigNumber((_1873_v5).length))];
            let _rhs354 = !(p0) || (p2);
            let _lhs239 = _1872_v4;
            let _lhs240 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1872_v4).length));
            _1871_v3 = _rhs353;
            _lhs239[_lhs240] = _rhs354;
            let _1874_v6;
            _1874_v6 = _dafny.Set.fromElements(_1871_v3, p2);
            let _1875_v7;
            _1875_v7 = _dafny.Seq.of(_1874_v6, _1874_v6, _1874_v6, _module.__default.fm49(globalState));
            let _1876_v8;
            _1876_v8 = _module.D16.create_DC30(_1875_v7);
            (globalState).f1 = new BigNumber(((_1876_v8).dtor_cf39).length);
            (globalState).f1 = (_module.__default.fm58(globalState)).dtor_cf22;
            let _1877_v9;
            _1877_v9 = _module.D0.create_DC1();
            let _1878_v10;
            _1878_v10 = _dafny.Set.fromElements(_1877_v9, _1877_v9, _1877_v9);
            let _1879_v11;
            _1879_v11 = _1878_v10;
            _1879_v11 = _1879_v11;
          }
        }
      }
      let _1880_v12;
      let _nw316 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1880_v12 = _nw316;
      let _index324 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1880_v12).length));
      (_1880_v12)[_index324] = _1869_v2;
      let _index325 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1880_v12).length));
      (_1880_v12)[_index325] = _dafny.Seq.Concat(_1869_v2, _dafny.Seq.UnicodeFromString("q"));
      let _1881_v13;
      _1881_v13 = false;
      let _1882_v14;
      _1882_v14 = _dafny.Seq.of((_1880_v12)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_1880_v12).length))], (_1880_v12)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_1880_v12).length))]);
      _1881_v13 = ((p2) ? (true) : (!_dafny.Seq.contains(_1882_v14, (_1880_v12)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_1880_v12).length))])));
      let _1883_v15;
      let _nw317 = Array((new BigNumber(3)).toNumber());
      _nw317[(_dafny.ZERO).toNumber()] = p0;
      _nw317[(_dafny.ONE).toNumber()] = _1881_v13;
      _nw317[(new BigNumber(2)).toNumber()] = p0;
      _1883_v15 = _nw317;
      let _index326 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1883_v15).length));
      (_1883_v15)[_index326] = _1881_v13;
      let _1884_v16;
      _1884_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1867_v0,_1867_v0);
      let _1885_v17;
      _1885_v17 = _1884_v16;
      let _pat_let_tv27 = p0;
      let _index327 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_1883_v15).length));
      (_1883_v15)[_index327] = !(function (_source24) {
        let _1886___mcc_h0 = _source24;
        let _1887_cf11 = _1886___mcc_h0;
        return _pat_let_tv27;
      }(_1884_v16));
      return;
    }
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = undefined;
      let _1888_v0;
      let _nw318 = Array((new BigNumber(10)).toNumber()).fill(false);
      _1888_v0 = _nw318;
      let _index328 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_1888_v0).length));
      (_1888_v0)[_index328] = _module.__default.fm0(p1, globalState);
      let _1889_v1;
      _1889_v1 = new BigNumber(-792);
      let _index329 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length));
      (p2)[_index329] = _1889_v1;
      let _1890_v2;
      _1890_v2 = _dafny.Seq.UnicodeFromString("oyvrsduxe");
      let _1891_v3;
      _1891_v3 = _dafny.Seq.of(p0, p1);
      let _1892_v4;
      _1892_v4 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("thordvy")).length));
      let _1893_v5;
      _1893_v5 = _module.D1.create_DC4(_1889_v1);
      let _index330 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_1888_v0).length));
      let _index331 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length));
      let _rhs355 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(941)), function (_1894_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length)).isEqualTo(new BigNumber(-462));
      let _rhs356 = new BigNumber((_dafny.Seq.Concat(_1890_v2, _1890_v2)).length);
      let _rhs357 = ((_dafny.ZERO).minus((_1889_v1).minus(_1889_v1))).multipliedBy(_module.__default.safeDivisionInt(_1889_v1, _1889_v1));
      let _rhs358 = ((_this).fm4(_1891_v3, _1889_v1, _1889_v1, _1892_v4, globalState)).minus((_1893_v5).dtor_cf4);
      let _lhs241 = _1888_v0;
      let _lhs242 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_1888_v0).length));
      let _lhs243 = p2;
      let _lhs244 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length));
      let _lhs245 = globalState;
      let _lhs246 = globalState;
      _lhs241[_lhs242] = _rhs355;
      _lhs243[_lhs244] = _rhs356;
      _lhs245.f1 = _rhs357;
      _lhs246.f1 = _rhs358;
      _1890_v2 = _1890_v2;
      let _1895_v6;
      _1895_v6 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length))],(p2)[_module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length))]);
      let _1896_v7;
      _1896_v7 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1897_v8;
      let _1898_v9;
      let _out46;
      let _out47;
      let _outcollector20 = _module.__default.m0(_1888_v0, _module.__default.fm36(_module.__default.fm3(p1, p1, globalState), (p2)[_module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length))], _1889_v1, _1895_v6, globalState), _dafny.areEqual(_1890_v2, _1890_v2), _1896_v7, globalState);
      _out46 = _outcollector20[0];
      _out47 = _outcollector20[1];
      _1897_v8 = _out46;
      _1898_v9 = _out47;
      _1896_v7 = _1896_v7;
      let _index332 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length));
      (p2)[_index332] = _module.__default.safeModuloInt((_this).fm7(p0, globalState), _1889_v1);
      let _1899_v10;
      _1899_v10 = p2;
      let _1900_v11;
      let _nw319 = Array((new BigNumber(9)).toNumber());
      _nw319[(_dafny.ZERO).toNumber()] = p2;
      _nw319[(_dafny.ONE).toNumber()] = p2;
      _nw319[(new BigNumber(2)).toNumber()] = p2;
      _nw319[(new BigNumber(3)).toNumber()] = p2;
      _nw319[(new BigNumber(4)).toNumber()] = p2;
      _nw319[(new BigNumber(5)).toNumber()] = p2;
      _nw319[(new BigNumber(6)).toNumber()] = (_1899_v10);
      _nw319[(new BigNumber(7)).toNumber()] = p2;
      _nw319[(new BigNumber(8)).toNumber()] = p2;
      _1900_v11 = _nw319;
      let _index333 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1900_v11).length));
      (_1900_v11)[_index333] = p2;
      let _1901_v12;
      _1901_v12 = _dafny.Set.fromElements(_1896_v7, _1896_v7);
      let _1902_v13;
      _1902_v13 = _dafny.Set.fromElements((_dafny.Set.fromElements(_1896_v7)).Union(_1901_v12));
      let _index334 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1900_v11).length));
      let _rhs359 = p2;
      let _rhs360 = (_1902_v13).Union(_1902_v13);
      let _lhs247 = _1900_v11;
      let _lhs248 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1900_v11).length));
      _lhs247[_lhs248] = _rhs359;
      _1902_v13 = _rhs360;
      r0 = (p2)[_module.__default.safeIndex(new BigNumber(112), new BigNumber((p2).length))];
      let _1903_v14;
      let _nw320 = new _module.C1();
      _nw320.__ctor();
      _1903_v14 = _nw320;
      r1 = _1903_v14;
      return [r0, r1];
    }
    m9(p0, globalState) {
      let _this = this;
      let _1904_v0;
      _1904_v0 = new BigNumber(-163);
      let _1905_v1;
      _1905_v1 = true;
      let _1906_v2;
      _1906_v2 = _dafny.Set.fromElements(p0);
      let _1907_v3;
      _1907_v3 = _dafny.Seq.of(true);
      let _1908_v4;
      _1908_v4 = _dafny.MultiSet.fromElements(_1904_v0);
      let _1909_v5;
      _1909_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1904_v0,_1904_v0);
      let _1910_v6;
      _1910_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1904_v0,_1905_v1);
      let _1911_v7;
      _1911_v7 = _module.D16.create_DC32((((_1910_v6).contains(new BigNumber(230))) ? ((_1910_v6).get(new BigNumber(230))) : (_1905_v1)));
      let _1912_v8;
      _1912_v8 = _dafny.Seq.of(_1911_v7, _module.D16.create_DC32(!(_1905_v1)), _1911_v7);
      let _1913_v9;
      _1913_v9 = _dafny.Seq.of(_1904_v0);
      let _1914_v10;
      let _nw321 = Array((new BigNumber(24)).toNumber());
      _nw321[(_dafny.ZERO).toNumber()] = new BigNumber(167);
      _nw321[(_dafny.ONE).toNumber()] = (_1904_v0).plus(_1904_v0);
      _nw321[(new BigNumber(2)).toNumber()] = ((_1905_v1) ? (_1904_v0) : (_1904_v0));
      _nw321[(new BigNumber(3)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(4)).toNumber()] = new BigNumber((_1906_v2).length);
      _nw321[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm59(globalState)).length));
      _nw321[(new BigNumber(6)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(7)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(8)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(9)).toNumber()] = new BigNumber(158);
      _nw321[(new BigNumber(10)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(11)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(12)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(13)).toNumber()] = _module.__default.fm3(_1905_v1, _1905_v1, globalState);
      _nw321[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1907_v3)).cardinality());
      _nw321[(new BigNumber(15)).toNumber()] = new BigNumber(-478);
      _nw321[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_module.__default.fm3(_1905_v1, _1905_v1, globalState), (((_1908_v4).contains(_1904_v0)) ? ((_1908_v4).get(_1904_v0)) : ((_dafny.ZERO).minus(_1904_v0))));
      _nw321[(new BigNumber(17)).toNumber()] = new BigNumber(((_1908_v4).Union((_1908_v4).update(_1904_v0, _module.__default.abs(_1904_v0)))).cardinality());
      _nw321[(new BigNumber(18)).toNumber()] = _1904_v0;
      _nw321[(new BigNumber(19)).toNumber()] = new BigNumber((_module.__default.fm49(globalState)).length);
      _nw321[(new BigNumber(20)).toNumber()] = new BigNumber((_module.__default.fm57(new BigNumber((_1909_v5).length), (_dafny.ZERO).minus(_1904_v0), globalState)).length);
      _nw321[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1912_v8).length));
      _nw321[(new BigNumber(22)).toNumber()] = (_1904_v0).minus(new BigNumber((_1913_v9).length));
      _nw321[(new BigNumber(23)).toNumber()] = _1904_v0;
      _1914_v10 = _nw321;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1914_v10).length))) {
        let _1915_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1915_i0)) && ((_1915_i0).isLessThan(new BigNumber((_1914_v10).length))))) {
          (_1914_v10)[(_1915_i0)] = _module.__default.safeModuloInt(_1915_i0, (((_1909_v5).contains(_1904_v0)) ? ((_1909_v5).get(_1904_v0)) : (new BigNumber((_dafny.Seq.of(_module.D15.create_DC29(_module.D15.create_DC27(_dafny.MultiSet.fromElements(_1905_v1, false))))).length))));
        }
      }
      _1905_v1 = (_1904_v0).isLessThanOrEqualTo(_1904_v0);
      let _1916_i1;
      _1916_i1 = _dafny.ZERO;
      L18: {
        while (!(_module.__default.fm0((_this).fm6(_1913_v9, globalState), globalState))) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1916_i1)) {
              break L18;
            }
            _1916_i1 = (_1916_i1).plus(_dafny.ONE);
            let _1917_v11;
            _1917_v11 = _1910_v6;
            let _source25 = _1917_v11;
            let _1918___mcc_h0 = _source25;
            let _1919_cf12 = _1918___mcc_h0;
            (globalState).f1 = _1904_v0;
            let _1920_v12;
            _1920_v12 = new _dafny.CodePoint('c'.codePointAt(0));
            let _1921_v13;
            _1921_v13 = _module.D14.create_DC26(true, _1904_v0, _1920_v12);
            let _1922_v14;
            _1922_v14 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1905_v1,new BigNumber(-334)));
            let _1923_v15;
            let _nw322 = Array((new BigNumber(28)).toNumber());
            _nw322[(_dafny.ZERO).toNumber()] = _module.D14.create_DC26(_1905_v1, _1904_v0, _1920_v12);
            _nw322[(_dafny.ONE).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(2)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(3)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(4)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(5)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(6)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(7)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(8)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(9)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(10)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(11)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(12)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(13)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(14)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(15)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(16)).toNumber()] = _module.D14.create_DC26(_1905_v1, new BigNumber(538), _1920_v12);
            _nw322[(new BigNumber(17)).toNumber()] = _module.D14.create_DC26((_1907_v3)[_module.__default.safeIndex(new BigNumber((_1922_v14).cardinality()), new BigNumber((_1907_v3).length))], new BigNumber(245), _1920_v12);
            _nw322[(new BigNumber(18)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(19)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(20)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(21)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(22)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(23)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(24)).toNumber()] = _1921_v13;
            _nw322[(new BigNumber(25)).toNumber()] = _module.__default.fm60(globalState);
            _nw322[(new BigNumber(26)).toNumber()] = ((_1905_v1) ? (_1921_v13) : (_1921_v13));
            _nw322[(new BigNumber(27)).toNumber()] = _module.__default.fm60(globalState);
            _1923_v15 = _nw322;
            _1923_v15 = _1923_v15;
            let _1924_v16;
            let _nw323 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
            _1924_v16 = _nw323;
            let _index335 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1924_v16).length));
            (_1924_v16)[_index335] = (_1909_v5).Merge(_1909_v5);
            let _index336 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1924_v16).length));
            (_1924_v16)[_index336] = _1909_v5;
            let _1925_v17;
            let _nw324 = new _module.C3();
            _nw324.__ctor(_1905_v1, _1905_v1);
            _1925_v17 = _nw324;
            let _1926_v18;
            _1926_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1925_v17,_1905_v1);
            _1905_v1 = (((_1926_v18).contains(_1925_v17)) ? ((_1926_v18).get(_1925_v17)) : (_1905_v1));
            _1905_v1 = _1905_v1;
            let _1927_v19;
            _1927_v19 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-648)), ((_1928_v0) => function (_1929_i2) {
              return _1928_v0;
            })(_1904_v0)), _1913_v9);
            let _1930_v20;
            let _nw325 = new _module.C0();
            _nw325.__ctor(_1905_v1, _1927_v19);
            _1930_v20 = _nw325;
            let _1931_v21;
            let _nw326 = new _module.C1();
            _nw326.__ctor();
            _1931_v21 = _nw326;
          }
        }
      }
      let _1932_i3;
      _1932_i3 = _dafny.ZERO;
      L19: {
        while ((_module.__default.safeModuloInt(_1904_v0, _1904_v0)).isLessThanOrEqualTo(_1904_v0)) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1932_i3)) {
              break L19;
            }
            _1932_i3 = (_1932_i3).plus(_dafny.ONE);
            if (_1905_v1) {
              let _1933_v22;
              _1933_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1905_v1,!(_1905_v1));
              _1933_v22 = (_1933_v22).update(_dafny.Seq.IsProperPrefixOf(_1913_v9, _1913_v9), _1905_v1);
              (globalState).f1 = (_1904_v0).minus((_1904_v0).multipliedBy(_1904_v0));
              let _1934_v23;
              _1934_v23 = new _dafny.CodePoint('p'.codePointAt(0));
              let _rhs361 = (_dafny.ZERO).minus(_1904_v0);
              let _rhs362 = _1905_v1;
              let _rhs363 = _1904_v0;
              let _rhs364 = _1934_v23;
              let _lhs249 = globalState;
              let _lhs250 = globalState;
              _lhs249.f1 = _rhs361;
              _1905_v1 = _rhs362;
              _lhs250.f1 = _rhs363;
              _1934_v23 = _rhs364;
              _1904_v0 = _1904_v0;
              let _1935_v24;
              let _nw327 = Array((_dafny.ONE).toNumber());
              _nw327[(_dafny.ZERO).toNumber()] = _1905_v1;
              _1935_v24 = _nw327;
              let _1936_v25;
              _1936_v25 = _dafny.Seq.of(_1935_v24);
              let _1937_v26;
              _1937_v26 = _dafny.Seq.of((_1936_v25)[_module.__default.safeIndex(_1904_v0, new BigNumber((_1936_v25).length))], _1935_v24, _1935_v24);
              _1935_v24 = (_1937_v26)[_module.__default.safeIndex(_1904_v0, new BigNumber((_1937_v26).length))];
            } else {
              (globalState).f1 = _1904_v0;
              let _1938_v27;
              _1938_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(p0, p0),(_this).fm29(_1904_v0, p0, globalState));
              _1938_v27 = (_1938_v27).update(_dafny.Seq.UnicodeFromString("eqjlgq"), (((_1909_v5).contains(_1904_v0)) ? ((_1909_v5).get(_1904_v0)) : (_1904_v0)));
              (globalState).f1 = _1904_v0;
              let _1939_v28;
              let _nw328 = new _module.C1();
              _nw328.__ctor();
              _1939_v28 = _nw328;
              _1908_v4 = _module.__default.fm32(globalState);
            }
            let _1940_v29;
            let _nw329 = Array((new BigNumber(3)).toNumber()).fill(false);
            _1940_v29 = _nw329;
            let _1941_v30;
            _1941_v30 = _1940_v29;
            let _1942_v31;
            _1942_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1941_v30,!(_1905_v1));
            _1942_v31 = (_1942_v31).update(_1940_v29, false);
            _1905_v1 = _1905_v1;
            let _1943_v32;
            let _nw330 = new _module.C3();
            _nw330.__ctor(_1905_v1, _1905_v1);
            _1943_v32 = _nw330;
          }
        }
      }
      let _1944_v33;
      _1944_v33 = new _dafny.CodePoint('v'.codePointAt(0));
      let _1945_v34;
      let _nw331 = new _module.C5();
      _nw331.__ctor(_1904_v0, _1944_v33, !((_1905_v1) || (_1905_v1)));
      _1945_v34 = _nw331;
      let _1946_v35;
      _1946_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1914_v10,_1907_v3);
      _1946_v35 = (_1946_v35).update(_1914_v10, _dafny.Seq.Concat(_1907_v3, _1907_v3));
      return;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
      this.f10 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f9, f10, f5, f6) {
      let _this = this;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f9, _this.f10);
    };
    fm6(p0, globalState) {
      let _this = this;
      return !((_this).f6);
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_this).f9;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f10;
    };
    fm5(p0, globalState) {
      let _this = this;
      if ((_this).f6) {
        return _module.D1.create_DC4((_this).f9);
      } else {
        return _module.D1.create_DC4(_this.f10);
      }
    };
    fm13(globalState) {
      let _this = this;
      return ((_this).f9).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("mhlxx")).length));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _1947_v0;
      let _init52 = ((_1948_p2) => function (_1949_i0) {
        return !_dafny.Seq.contains(_1948_p2, (_this).f5);
      })(p2);
      let _nw332 = Array((new BigNumber(29)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw332.length); _i0_52++) {
        _nw332[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _1947_v0 = _nw332;
      let _index337 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length));
      (_1947_v0)[_index337] = ((!((_this).f6)) ? (!(p0)) : ((_this).f6));
      let _index338 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length));
      (_1947_v0)[_index338] = p0;
      let _1950_v1;
      _1950_v1 = _dafny.MultiSet.fromElements(p0);
      (_this).f10 = (_this.f10).minus(_module.__default.safeModuloInt(new BigNumber((_1950_v1).cardinality()), _this.f10));
      let _1951_v2;
      _1951_v2 = _module.D2.create_DC6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), function (_1952_i1) {
  return (_this).f5;
}));
      let _source26 = _1951_v2;
      if (_source26.is_DC7) {
        let _1953___mcc_h0 = (_source26).cf7;
        let _1954___mcc_h1 = (_source26).cf8;
        let _1955_cf8 = _1954___mcc_h1;
        let _1956_cf7 = _1953___mcc_h0;
        let _1957_v3;
        _1957_v3 = new _dafny.CodePoint('l'.codePointAt(0));
        _1957_v3 = (_this).f5;
        let _1958_v4;
        _1958_v4 = _dafny.Seq.of(_this.f10, new BigNumber(929), _1956_cf7, (_dafny.ZERO).minus((_this).f9));
        (globalState).f1 = new BigNumber((_1958_v4).length);
        let _1959_v5;
        _1959_v5 = _dafny.Seq.of((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], (_this).f6);
        _1959_v5 = _module.__default.fm2((_1959_v5)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p1, _this.f10)).length))), new BigNumber((_1959_v5).length))], (_1959_v5)[_module.__default.safeIndex(p1, new BigNumber((_1959_v5).length))], globalState);
        r1 = p2;
      } else if (_source26.is_DC8) {
        let _1960___mcc_h2 = (_source26).cf9;
        let _1961_cf9 = _1960___mcc_h2;
        let _1962_v6;
        _1962_v6 = _dafny.Set.fromElements(true, (_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], p0, _1961_cf9);
        let _1963_v7;
        _1963_v7 = _dafny.Seq.of(false);
        let _1964_v8;
        _1964_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1963_v7,p1);
        let _1965_v9;
        _1965_v9 = _dafny.Seq.of(_1963_v7);
        let _1966_v10;
        _1966_v10 = _dafny.Map.Empty.slice().updateUnsafe(!(_1961_cf9),true);
        let _1967_v11;
        _1967_v11 = _dafny.Seq.of(p1, new BigNumber((p2).length), p1, p1);
        let _1968_v12;
        let _nw333 = Array((new BigNumber(12)).toNumber());
        _nw333[(_dafny.ZERO).toNumber()] = new BigNumber((_1962_v6).length);
        _nw333[(_dafny.ONE).toNumber()] = (_this.f10).plus((((_1964_v8).contains((_1965_v9)[_module.__default.safeIndex(_this.f10, new BigNumber((_1965_v9).length))])) ? ((_1964_v8).get((_1965_v9)[_module.__default.safeIndex(_this.f10, new BigNumber((_1965_v9).length))])) : (_this.f10)));
        _nw333[(new BigNumber(2)).toNumber()] = _this.f10;
        _nw333[(new BigNumber(3)).toNumber()] = _this.f10;
        _nw333[(new BigNumber(4)).toNumber()] = _this.f10;
        _nw333[(new BigNumber(5)).toNumber()] = _this.f10;
        _nw333[(new BigNumber(6)).toNumber()] = new BigNumber((p2).length);
        _nw333[(new BigNumber(7)).toNumber()] = new BigNumber(534);
        _nw333[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1966_v10).length), new BigNumber((_1967_v11).length));
        _nw333[(new BigNumber(9)).toNumber()] = p1;
        _nw333[(new BigNumber(10)).toNumber()] = new BigNumber(599);
        _nw333[(new BigNumber(11)).toNumber()] = (_this).fm7((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], globalState);
        _1968_v12 = _nw333;
        let _index339 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1968_v12).length));
        (_1968_v12)[_index339] = p1;
        let _1969_v13;
        _1969_v13 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f6) ? ((_1951_v2).dtor_cf6) : (p2)),!(_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f6)).contains(p1));
        let _index340 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1968_v12).length));
        let _rhs365 = (_this).f9;
        let _rhs366 = _1969_v13;
        let _rhs367 = (_module.__default.safeModuloInt((_this).f9, (_dafny.ZERO).minus((_this).f9))).multipliedBy((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(953), new BigNumber(-194))));
        let _lhs251 = _1968_v12;
        let _lhs252 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1968_v12).length));
        let _lhs253 = _this;
        _lhs251[_lhs252] = _rhs365;
        _1969_v13 = _rhs366;
        _lhs253.f10 = _rhs367;
        _1963_v7 = _1963_v7;
        let _1970_v14;
        let _nw334 = new _module.C0();
        _nw334.__ctor((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), ((_1971_v11) => function (_1972_i2) {
          return _1971_v11;
        })(_1967_v11)));
        _1970_v14 = _nw334;
        let _1973_v15;
        let _nw335 = new _module.C0();
        _nw335.__ctor(p0, _1970_v14.f12);
        _1973_v15 = _nw335;
      } else if (_source26.is_DC6) {
        let _1974___mcc_h3 = (_source26).cf6;
        let _1975_cf6 = _1974___mcc_h3;
        (globalState).f1 = _this.f10;
        (globalState).f1 = _this.f10;
        (_this).f10 = (_this).f9;
        let _index341 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length));
        (_1947_v0)[_index341] = !(((_this).f9).isLessThan(_module.__default.fm3((_this).f6, p0, globalState)));
      } else {
        let _1976___mcc_h4 = (_source26).cf10;
        let _1977_cf10 = _1976___mcc_h4;
        let _1978_v16;
        _1978_v16 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cofnqf"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), function (_1979_i3) {
          return (_this).f5;
        }), p2));
        _1978_v16 = _1978_v16;
        let _1980_v17;
        let _nw336 = new _module.C0();
        _nw336.__ctor((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), function (_1981_i4) {
          return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f6, (_this).f6)).length));
        }));
        _1980_v17 = _nw336;
        let _1982_v18;
        _1982_v18 = _dafny.Seq.of(p1);
        let _1983_v19;
        _1983_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1980_v17,_1982_v18);
        let _1984_v20;
        _1984_v20 = _module.D2.create_DC8(p0);
        let _index342 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length));
        let _rhs368 = ((_this.f10).multipliedBy(new BigNumber(611))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6)).length));
        let _rhs369 = ((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))]) || ((_1984_v20).dtor_cf9);
        let _rhs370 = ((_1983_v19).update(_1980_v17, _1982_v18)).Merge((_1983_v19).update(_1980_v17, _dafny.Seq.of(new BigNumber((_1950_v1).cardinality()))));
        let _lhs254 = _1947_v0;
        let _lhs255 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length));
        r0 = _rhs368;
        _lhs254[_lhs255] = _rhs369;
        _1983_v19 = _rhs370;
        let _1985_v21;
        _1985_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_this.f10);
        let _1986_v22;
        _1986_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1985_v21,((!(p0)) ? (p1) : (_this.f10)));
        _1986_v22 = (_1986_v22).update(_1985_v21, (_this.f10).plus(new BigNumber((_1982_v18).length)));
        let _1987_v23;
        _1987_v23 = _dafny.Seq.of(p0);
        let _1988_v24;
        let _nw337 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1988_v24 = _nw337;
        let _1989_v25;
        _1989_v25 = _1988_v24;
        let _rhs371 = _dafny.Seq.Concat(_1987_v23, _dafny.Seq.Concat(_dafny.Seq.of(p0), _1987_v23));
        let _rhs372 = ((false) ? (_1951_v2) : (_1951_v2));
        let _rhs373 = (((!(p0)) ? (_1989_v25) : (_1989_v25)));
        _1987_v23 = _rhs371;
        _1951_v2 = _rhs372;
        _1988_v24 = _rhs373;
      }
      let _1990_v26;
      _1990_v26 = _dafny.Seq.of(p0);
      let _1991_v27;
      _1991_v27 = _dafny.Seq.of(_module.__default.fm3(p0, (_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))], globalState));
      let _1992_v28;
      _1992_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1990_v26,(_this).fm6(_1991_v27, globalState));
      _1992_v28 = (_1992_v28).update(_1990_v26, p0);
      let _1993_v29;
      _1993_v29 = _dafny.Seq.of(p2, p2);
      (_this).f10 = _module.__default.safeDivisionInt(new BigNumber(((_1993_v29)[_module.__default.safeIndex(_this.f10, new BigNumber((_1993_v29).length))]).length), p1);
      let _1994_v30;
      let _nw338 = new _module.C1();
      _nw338.__ctor();
      _1994_v30 = _nw338;
      r0 = ((false) ? ((_this).f6) : ((_1947_v0)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1947_v0).length))]));
      r1 = p2;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1995_v0;
      _1995_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!((_this).f6));
      let _1996_v1;
      _1996_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(-162));
      let _1997_v2;
      _1997_v2 = _dafny.Seq.of(_1995_v0, _1995_v0, _1995_v0);
      let _1998_v3;
      let _nw339 = Array((new BigNumber(14)).toNumber());
      _nw339[(_dafny.ZERO).toNumber()] = _1995_v0;
      _nw339[(_dafny.ONE).toNumber()] = _1995_v0;
      _nw339[(new BigNumber(2)).toNumber()] = (_1995_v0).Merge(_1995_v0);
      _nw339[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      _nw339[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f6)),(_this).f6);
      _nw339[(new BigNumber(5)).toNumber()] = (_1995_v0).update((_this).f6, (_this).f6);
      _nw339[(new BigNumber(6)).toNumber()] = (_1995_v0).Merge(_1995_v0);
      _nw339[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      _nw339[(new BigNumber(8)).toNumber()] = _module.__default.fm19(_module.__default.fm0((_this).f6, globalState), (((_1996_v1).contains((_this).f9)) ? ((_1996_v1).get((_this).f9)) : (new BigNumber((_dafny.Seq.UnicodeFromString("nud")).length))), globalState);
      _nw339[(new BigNumber(9)).toNumber()] = _1995_v0;
      _nw339[(new BigNumber(10)).toNumber()] = _1995_v0;
      _nw339[(new BigNumber(11)).toNumber()] = (_1995_v0).Merge(_1995_v0);
      _nw339[(new BigNumber(12)).toNumber()] = _1995_v0;
      _nw339[(new BigNumber(13)).toNumber()] = (_1995_v0).Merge((_1997_v2)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1997_v2).length))]);
      _1998_v3 = _nw339;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1998_v3).length))) {
        let _1999_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1999_i0)) && ((_1999_i0).isLessThan(new BigNumber((_1998_v3).length))))) {
          (_1998_v3)[(_1999_i0)] = _1995_v0;
        }
      }
      let _2000_i1;
      _2000_i1 = _dafny.ZERO;
      L20: {
        while ((_this).f6) {
          C20: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2000_i1)) {
              break L20;
            }
            _2000_i1 = (_2000_i1).plus(_dafny.ONE);
            let _2001_v4;
            _2001_v4 = false;
            _2001_v4 = (((_1995_v0).contains((_this).f6)) ? ((_1995_v0).get((_this).f6)) : (((_this).f6) && (_2001_v4)));
            (globalState).f1 = (_this).f9;
            let _2002_v5;
            _2002_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f5);
            let _2003_v6;
            _2003_v6 = _dafny.Seq.of(_2001_v4);
            let _2004_v7;
            _2004_v7 = _dafny.Set.fromElements(new BigNumber((_2003_v6).length), new BigNumber((_dafny.MultiSet.fromElements(_2001_v4, (_this).f6)).cardinality()));
            let _2005_v8;
            _2005_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(133),_module.D1.create_DC4((_this).f9));
            let _2006_v9;
            _2006_v9 = _dafny.MultiSet.fromElements(_2004_v7, _dafny.Set.fromElements(new BigNumber((_2005_v8).length)));
            let _2007_v10;
            let _nw340 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
            _2007_v10 = _nw340;
            let _2008_v11;
            _2008_v11 = _dafny.MultiSet.fromElements(_2007_v10, _2007_v10, _2007_v10, _2007_v10, _2007_v10);
            _2002_v5 = (_2002_v5).update((new BigNumber((_2006_v9).cardinality())).isLessThanOrEqualTo(new BigNumber((_2008_v11).cardinality())), (_this).f5);
            let _2009_v12;
            _2009_v12 = _dafny.Seq.of(new BigNumber(445), _this.f10);
            let _2010_v13;
            _2010_v13 = _dafny.Seq.of(_2009_v12, _2009_v12);
            let _2011_v14;
            let _nw341 = new _module.C0();
            _nw341.__ctor(_module.__default.fm0((_this).f6, globalState), _dafny.Seq.Concat(_2010_v13, _2010_v13));
            _2011_v14 = _nw341;
          }
        }
      }
      if ((_this).f6) {
        (globalState).f1 = _this.f10;
        let _2012_v15;
        let _nw342 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _2012_v15 = _nw342;
        let _2013_v16;
        _2013_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2012_v15);
        let _index343 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
        (_2012_v15)[_index343] = ((_this).f9).multipliedBy((_this).f9);
        let _2014_v17;
        _2014_v17 = false;
        let _2015_v18;
        _2015_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,false);
        let _2016_v19;
        _2016_v19 = _2015_v18;
        let _2017_v20;
        _2017_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2016_v19,new _dafny.CodePoint('j'.codePointAt(0)));
        let _index344 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
        let _rhs374 = _2013_v16;
        let _rhs375 = _this.f10;
        let _rhs376 = false;
        let _rhs377 = (_this).fm8(!(_2017_v20).equals(_2017_v20), (_this).f6, false, globalState);
        let _lhs256 = _2012_v15;
        let _lhs257 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
        let _lhs258 = _this;
        _2013_v16 = _rhs374;
        _lhs256[_lhs257] = _rhs375;
        _2014_v17 = _rhs376;
        _lhs258.f10 = _rhs377;
        let _2018_v21;
        _2018_v21 = _dafny.Seq.of((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))], _this.f10);
        let _2019_v22;
        let _nw343 = new _module.C0();
        _nw343.__ctor(_2014_v17, _dafny.Seq.of(_2018_v21));
        _2019_v22 = _nw343;
        _2019_v22 = _2019_v22;
        if (((_this).f9).isLessThan(_module.__default.safeDivisionInt((_this).f9, (_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]))) {
          let _2020_v23;
          _2020_v23 = new _dafny.CodePoint('a'.codePointAt(0));
          _2020_v23 = _2020_v23;
          let _2021_v24;
          _2021_v24 = _dafny.Seq.UnicodeFromString("mcwnadjh");
          let _2022_v25;
          _2022_v25 = _dafny.Set.fromElements(_this.f10);
          let _2023_v26;
          _2023_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))],(_this).f5);
          let _index345 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
          let _index346 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
          let _rhs378 = (_2019_v22).f11;
          let _rhs379 = ((_2014_v17) ? (_dafny.Seq.Concat(_2021_v24, _2021_v24)) : (_2021_v24));
          let _rhs380 = ((_2022_v25).Difference(_2022_v25)).IsSubsetOf(_2022_v25);
          let _rhs381 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1996_v1).length), (_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]));
          let _rhs382 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2021_v24, _2021_v24), _module.__default.safeIndex((_this).f9, new BigNumber((_dafny.Seq.Concat(_2021_v24, _2021_v24)).length)), (((_2023_v26).contains((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))])) ? ((_2023_v26).get((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))])) : ((_this).f5)))).length)), (_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]);
          let _lhs259 = _2012_v15;
          let _lhs260 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
          let _lhs261 = _2012_v15;
          let _lhs262 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length));
          _2014_v17 = _rhs378;
          _2021_v24 = _rhs379;
          _2014_v17 = _rhs380;
          _lhs259[_lhs260] = _rhs381;
          _lhs261[_lhs262] = _rhs382;
          let _2024_v27;
          _2024_v27 = _module.D2.create_DC6(_2021_v24);
          _2024_v27 = _module.D2.create_DC6(_2021_v24);
          let _2025_v28;
          _2025_v28 = _dafny.Seq.of((_2019_v22).f11);
          let _2026_v29;
          _2026_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2025_v28, _2025_v28),_2014_v17);
          _2026_v29 = (_2026_v29).Merge(_2026_v29);
          let _2027_v30;
          _2027_v30 = _dafny.Set.fromElements((_this).f6);
          (globalState).f1 = (_this.f10).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_2019_v22).f11)).length)).minus(new BigNumber((_2027_v30).length)));
        } else {
          r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt((_module.D1.create_DC4((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))])).dtor_cf4, (_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]), new BigNumber((_2019_v22.f12).length))));
          (_this).f10 = new BigNumber(((((!((_this).f6)) ? (_1995_v0) : (_dafny.Map.Empty.slice().updateUnsafe((_2019_v22).f11,(_2019_v22).f11)))).Merge(_1995_v0)).length);
          let _2028_v31;
          _2028_v31 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))],!((_this).f6)));
          _2014_v17 = ((_this).f9).isEqualTo(new BigNumber((_2028_v31).length));
          let _2029_v32;
          let _nw344 = Array((new BigNumber(23)).toNumber()).fill(false);
          _2029_v32 = _nw344;
          let _2030_v33;
          _2030_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2029_v32);
          let _2031_v34;
          _2031_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_this.f10),(_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]);
          let _rhs383 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("latagnlxg"), _dafny.Seq.UnicodeFromString("avl"));
          let _rhs384 = (((_2030_v33).contains(new BigNumber((_2031_v34).length))) ? ((_2030_v33).get(new BigNumber((_2031_v34).length))) : (_2029_v32));
          let _rhs385 = (_dafny.ZERO).minus(_this.f10);
          _2014_v17 = _rhs383;
          _2029_v32 = _rhs384;
          r0 = _rhs385;
          let _2032_v35;
          _2032_v35 = _dafny.Seq.UnicodeFromString("ywr");
          _2032_v35 = _dafny.Seq.Concat(_2032_v35, (((_this).f6) ? (_2032_v35) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), function (_2033_i2) {
            return (_this).f5;
          }))));
        }
        let _2034_v36;
        _2034_v36 = _module.D0.create_DC3(_module.D0.create_DC1());
        let _2035_v37;
        let _init53 = function (_2036_i3) {
          return (_this).f6;
        };
        let _nw345 = Array((new BigNumber(6)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw345.length); _i0_53++) {
          _nw345[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2035_v37 = _nw345;
        let _2037_v38;
        _2037_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2035_v37,(_this).f5);
        let _2038_v39;
        _2038_v39 = _module.D0.create_DC2(_2037_v38, new _dafny.CodePoint('e'.codePointAt(0)));
        let _2039_v40;
        _2039_v40 = _module.D0.create_DC3(_2038_v39);
        let _pat_let_tv28 = _2038_v39;
        let _source27 = function (_pat_let32_0) {
          return function (_2040_dt__update__tmp_h0) {
            return function (_pat_let33_0) {
              return function (_2041_dt__update_hcf3_h0) {
                return _module.D0.create_DC3(_2041_dt__update_hcf3_h0);
              }(_pat_let33_0);
            }(_pat_let_tv28);
          }(_pat_let32_0);
        }(_2034_v36);
        if (_source27.is_DC1) {
          _2014_v17 = (_2019_v22).f11;
          let _2042_v41;
          _2042_v41 = _module.D1.create_DC4(new BigNumber(-663));
          let _2043_v42;
          let _nw346 = Array((new BigNumber(7)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = _2042_v41;
          _nw346[(_dafny.ONE).toNumber()] = _2042_v41;
          _nw346[(new BigNumber(2)).toNumber()] = _module.D1.create_DC4(new BigNumber((_dafny.Set.fromElements((_this).f6)).length));
          _nw346[(new BigNumber(3)).toNumber()] = _2042_v41;
          _nw346[(new BigNumber(4)).toNumber()] = _2042_v41;
          _nw346[(new BigNumber(5)).toNumber()] = _2042_v41;
          _nw346[(new BigNumber(6)).toNumber()] = _2042_v41;
          _2043_v42 = _nw346;
          let _index347 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2043_v42).length));
          (_2043_v42)[_index347] = _2042_v41;
          let _index348 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2043_v42).length));
          (_2043_v42)[_index348] = _2042_v41;
          let _2044_v43;
          _2044_v43 = _dafny.Seq.of(_2012_v15, _2012_v15);
          let _2045_v44;
          _2045_v44 = _dafny.Seq.UnicodeFromString("qucb");
          let _rhs386 = (_dafny.ZERO).minus((_2018_v21)[_module.__default.safeIndex((_this).f9, new BigNumber((_2018_v21).length))]);
          let _rhs387 = ((_this).f9).minus((_dafny.ZERO).minus((_this.f10).minus((_this).f9)));
          let _rhs388 = _dafny.Seq.of(_2012_v15);
          let _rhs389 = _dafny.Seq.UnicodeFromString("yhwkuxs");
          let _lhs263 = _this;
          let _lhs264 = globalState;
          _lhs263.f10 = _rhs386;
          _lhs264.f1 = _rhs387;
          _2044_v43 = _rhs388;
          _2045_v44 = _rhs389;
          let _2046_v45;
          _2046_v45 = _dafny.MultiSet.fromElements(_module.__default.fm0((_2019_v22).f11, globalState), (_2019_v22).f11);
          let _2047_v46;
          _2047_v46 = _dafny.Seq.of((_this).f6, !(_2014_v17));
          let _2048_v47;
          let _nw347 = Array((new BigNumber(18)).toNumber());
          _nw347[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements((_2019_v22).f11, (_this).f6);
          _nw347[(_dafny.ONE).toNumber()] = (_2046_v45).Union(_2046_v45);
          _nw347[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements((_2019_v22).f11, (_2019_v22).f11)).Difference(_2046_v45);
          _nw347[(new BigNumber(3)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(4)).toNumber()] = (_2046_v45).Difference(_dafny.MultiSet.FromArray(_2047_v46));
          _nw347[(new BigNumber(5)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements((_this).f6, false, (_2019_v22).f11, !((_this).f6), (_2019_v22).f11);
          _nw347[(new BigNumber(7)).toNumber()] = (_2046_v45).Union(_dafny.MultiSet.FromArray(_2047_v46));
          _nw347[(new BigNumber(8)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(9)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(10)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(11)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(12)).toNumber()] = (_2046_v45).Union(_2046_v45);
          _nw347[(new BigNumber(13)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(14)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(15)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(16)).toNumber()] = _2046_v45;
          _nw347[(new BigNumber(17)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f6)).Union((_dafny.MultiSet.fromElements((_this).f6, (_2019_v22).f11, (_this).f6)).update((_2019_v22).f11, _module.__default.abs(((_2043_v42)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_2043_v42).length))]).dtor_cf4)));
          _2048_v47 = _nw347;
          let _index349 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_2048_v47).length));
          (_2048_v47)[_index349] = (_2046_v45).Union(_2046_v45);
          let _index350 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_2048_v47).length));
          (_2048_v47)[_index350] = _2046_v45;
        } else if (_source27.is_DC2) {
          let _2049___mcc_h0 = (_source27).cf1;
          let _2050___mcc_h1 = (_source27).cf2;
          let _2051_cf2 = _2050___mcc_h1;
          let _2052_cf1 = _2049___mcc_h0;
          let _2053_v48;
          let _nw348 = new _module.C1();
          _nw348.__ctor();
          _2053_v48 = _nw348;
          _1995_v0 = (_1995_v0).update((_this.f10).isLessThan((_this).f9), (_2019_v22).f11);
          _2014_v17 = !((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]).isEqualTo(_module.__default.safeModuloInt((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))], (_this).f9));
          _2014_v17 = (_2019_v22).f11;
        } else if (_source27.is_DC0) {
          let _2054___mcc_h2 = (_source27).cf0;
          let _2055_cf0 = _2054___mcc_h2;
          let _2056_v49;
          _2056_v49 = _dafny.Seq.UnicodeFromString("psooqryel");
          _2055_cf0 = _dafny.Seq.IsPrefixOf(_2056_v49, _dafny.Seq.UnicodeFromString("udq"));
          let _2057_v50;
          _2057_v50 = _dafny.Seq.of(_2012_v15);
          let _rhs390 = _2056_v49;
          let _rhs391 = _module.__default.safeModuloInt(new BigNumber(-128), (_this).f9);
          let _rhs392 = true;
          let _rhs393 = _2055_cf0;
          let _rhs394 = _dafny.Seq.IsProperPrefixOf(_2057_v50, _2057_v50);
          let _lhs265 = globalState;
          _2056_v49 = _rhs390;
          _lhs265.f1 = _rhs391;
          _2014_v17 = _rhs392;
          _2014_v17 = _rhs393;
          _2055_cf0 = _rhs394;
          let _2058_v51;
          _2058_v51 = _dafny.Seq.of(_2014_v17);
          _2014_v17 = (true) || ((_2058_v51)[_module.__default.safeIndex((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))], new BigNumber((_2058_v51).length))]);
          _2014_v17 = !(((_2019_v22).f11) === (!(_dafny.MultiSet.FromArray(_2058_v51)).contains((_2019_v22).f11)));
        } else {
          let _2059___mcc_h3 = (_source27).cf3;
          let _2060_cf3 = _2059___mcc_h3;
          let _2061_v52;
          _2061_v52 = _2012_v15;
          let _2062_v53;
          _2062_v53 = _dafny.Set.fromElements(_2061_v52, _2061_v52, _2061_v52, _2061_v52);
          _2062_v53 = _2062_v53;
          let _2063_v54;
          _2063_v54 = _module.D0.create_DC0((_2019_v22).f11);
          let _2064_v55;
          _2064_v55 = _dafny.Set.fromElements(_2063_v54, _module.D0.create_DC0(_2014_v17));
          _2014_v17 = !(_2064_v55).equals(_2064_v55);
          let _2065_v56;
          _2065_v56 = _module.D0.create_DC2(_2037_v38, (_this).f5);
          let _2066_v57;
          _2066_v57 = _dafny.Seq.of(_2065_v56);
          let _2067_v58;
          let _nw349 = new _module.C2();
          _nw349.__ctor((_this).f5, (_this).f6);
          _2067_v58 = _nw349;
          let _2068_v59;
          _2068_v59 = _dafny.MultiSet.fromElements(_2067_v58);
          let _2069_v60;
          _2069_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2066_v57),_2068_v59);
          let _2070_v61;
          _2070_v61 = _dafny.Seq.of(_2066_v57, _2066_v57);
          _2069_v60 = (_2069_v60).update(_dafny.Seq.Concat(_2070_v61, _2070_v61), (_2068_v59).Intersect(_2068_v59));
          let _2071_v62;
          _2071_v62 = _dafny.Seq.of(true);
          _1995_v0 = (_1995_v0).update((_2071_v62)[_module.__default.safeIndex(_this.f10, new BigNumber((_2071_v62).length))], (_this.f10).isLessThanOrEqualTo((_2012_v15)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_2012_v15).length))]));
        }
      } else {
        let _2072_v63;
        _2072_v63 = true;
        let _2073_v64;
        _2073_v64 = _dafny.Set.fromElements(!(_2072_v63), _2072_v63);
        let _2074_v65;
        _2074_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_2073_v64).IsDisjointFrom(_2073_v64));
        _2072_v63 = (((_2074_v65).contains((_this).f9)) ? ((_2074_v65).get((_this).f9)) : ((_this).f6));
        _2072_v63 = _2072_v63;
        let _2075_v66;
        _2075_v66 = _dafny.Seq.of(false);
        let _2076_v67;
        _2076_v67 = _2075_v66;
        let _2077_v68;
        _2077_v68 = _dafny.Seq.UnicodeFromString("ekfgutbiw");
        let _2078_v69;
        _2078_v69 = _dafny.Seq.of(new BigNumber((_2077_v68).length));
        let _rhs395 = (_2076_v67);
        let _rhs396 = ((_2073_v64).IsSubsetOf(_2073_v64)) === ((_2075_v66)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f10), new BigNumber((_2075_v66).length))]);
        let _rhs397 = !_dafny.Seq.contains(_2078_v69, _module.__default.safeDivisionInt(new BigNumber(652), (_dafny.ZERO).minus((_this).f9)));
        let _rhs398 = (_this).f6;
        _2075_v66 = _rhs395;
        _2072_v63 = _rhs396;
        _2072_v63 = _rhs397;
        _2072_v63 = _rhs398;
        let _2079_v70;
        let _init54 = ((_2080_v65) => function (_2081_i4) {
          return (((_2080_v65).contains(new BigNumber((_dafny.Seq.of((_this).f5, (_this).f5, (_this).f5)).length))) ? ((_2080_v65).get(new BigNumber((_dafny.Seq.of((_this).f5, (_this).f5, (_this).f5)).length))) : (true));
        })(_2074_v65);
        let _nw350 = Array((new BigNumber(8)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw350.length); _i0_54++) {
          _nw350[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2079_v70 = _nw350;
        _2079_v70 = _2079_v70;
        let _2082_v71;
        let _nw351 = Array((new BigNumber(10)).toNumber()).fill([]);
        _2082_v71 = _nw351;
        let _index351 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2082_v71).length));
        (_2082_v71)[_index351] = _2079_v70;
        let _index352 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2082_v71).length));
        let _rhs399 = (_this.f10).isLessThanOrEqualTo(new BigNumber((_2077_v68).length));
        let _rhs400 = _this.f10;
        let _rhs401 = (_this).f9;
        let _rhs402 = _2079_v70;
        let _rhs403 = new BigNumber(848);
        let _lhs266 = _this;
        let _lhs267 = globalState;
        let _lhs268 = _2082_v71;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2082_v71).length));
        let _lhs270 = globalState;
        _2072_v63 = _rhs399;
        _lhs266.f10 = _rhs400;
        _lhs267.f1 = _rhs401;
        _lhs268[_lhs269] = _rhs402;
        _lhs270.f1 = _rhs403;
      }
      let _2083_v72;
      _2083_v72 = false;
      _2083_v72 = _2083_v72;
      let _hi11 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f10, (_this).f9));
      for (let _2084_i5 = new BigNumber(-364); _2084_i5.isLessThan(_hi11); _2084_i5 = _2084_i5.plus(_dafny.ONE)) {
        let _2085_v73;
        let _nw352 = Array((new BigNumber(7)).toNumber()).fill(false);
        _2085_v73 = _nw352;
        let _index353 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2085_v73).length));
        (_2085_v73)[_index353] = (_this).f6;
        let _index354 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2085_v73).length));
        (_2085_v73)[_index354] = !(_2083_v72);
        let _2086_v74;
        let _nw353 = new _module.C1();
        _nw353.__ctor();
        _2086_v74 = _nw353;
        let _2087_v75;
        _2087_v75 = _dafny.MultiSet.fromElements(new BigNumber(54));
        let _2088_v76;
        _2088_v76 = _dafny.Seq.of(true);
        let _2089_v77;
        _2089_v77 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_2087_v75);
        let _2090_v78;
        _2090_v78 = _dafny.Seq.of((_this).f9, _this.f10, (_this).f9);
        let _2091_v79;
        _2091_v79 = _dafny.Seq.of(_2090_v78, _2090_v78);
        let _2092_v80;
        let _nw354 = new _module.C0();
        _nw354.__ctor((_2086_v74).fm6(_dafny.Seq.update(_2090_v78, _module.__default.safeIndex(_this.f10, new BigNumber((_2090_v78).length)), _2084_i5), globalState), _2091_v79);
        _2092_v80 = _nw354;
        let _2093_v81;
        _2093_v81 = _dafny.MultiSet.fromElements(_2092_v80, _2092_v80);
        let _index355 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2085_v73).length));
        let _rhs404 = (((_2089_v77).contains(_2084_i5)) ? ((_2089_v77).get(_2084_i5)) : (_2087_v75));
        let _rhs405 = (((_2093_v81).contains(_2092_v80)) ? ((_2093_v81).get(_2092_v80)) : ((_this).f9));
        let _rhs406 = _2088_v76;
        let _rhs407 = (_this).fm6(_dafny.Seq.Concat(_2090_v78, _2090_v78), globalState);
        let _lhs271 = _2085_v73;
        let _lhs272 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2085_v73).length));
        _2087_v75 = _rhs404;
        r0 = _rhs405;
        _2088_v76 = _rhs406;
        _lhs271[_lhs272] = _rhs407;
        let _index356 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2085_v73).length));
        (_2085_v73)[_index356] = _2083_v72;
      }
      let _2094_i6;
      _2094_i6 = _dafny.ZERO;
      L21: {
        while (_module.__default.fm0(!(_2083_v72), globalState)) {
          C21: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2094_i6)) {
              break L21;
            }
            _2094_i6 = (_2094_i6).plus(_dafny.ONE);
            if ((_this).f6) {
              let _2095_v82;
              _2095_v82 = _dafny.Set.fromElements((_this).f5, (_this).f5, (_this).f5, (_this).f5);
              let _2096_v83;
              let _init55 = function (_2097_i7) {
                return (_2097_i7).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ukbkpwm")).length));
              };
              let _nw355 = Array((new BigNumber(10)).toNumber());
              for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw355.length); _i0_55++) {
                _nw355[_i0_55] = _init55(new BigNumber(_i0_55));
              }
              _2096_v83 = _nw355;
              let _index357 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length));
              (_2096_v83)[_index357] = _this.f10;
              let _2098_v84;
              _2098_v84 = _module.D2.create_DC8(_2083_v72);
              let _2099_v85;
              _2099_v85 = _dafny.MultiSet.fromElements((_this).f9);
              let _2100_v86;
              _2100_v86 = _dafny.MultiSet.fromElements((_this).f9, (_this).f9, _this.f10, new BigNumber((_2099_v85).cardinality()));
              let _2101_v87;
              _2101_v87 = _dafny.MultiSet.fromElements(_this.f10, new BigNumber((_2099_v85).cardinality()));
              let _2102_v88;
              _2102_v88 = _module.D6.create_DC14(_2083_v72);
              let _2103_v89;
              _2103_v89 = _dafny.Seq.UnicodeFromString("caicnmft");
              let _2104_v90;
              _2104_v90 = _module.D2.create_DC7(_this.f10, new BigNumber((_2103_v89).length));
              let _2105_v91;
              _2105_v91 = _dafny.Seq.of(_dafny.Seq.of((_2104_v90).dtor_cf7, (_this).f9));
              let _2106_v92;
              let _nw356 = new _module.C0();
              _nw356.__ctor((_2102_v88).dtor_cf15, _2105_v91);
              _2106_v92 = _nw356;
              let _2107_v93;
              _2107_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2106_v92,(_this).f9);
              let _2108_v94;
              _2108_v94 = _dafny.Seq.of(_2083_v72, (_this).f6, false);
              let _2109_v95;
              let _nw357 = Array((new BigNumber(15)).toNumber());
              _nw357[(_dafny.ZERO).toNumber()] = !((_this).f6) || (false);
              _nw357[(_dafny.ONE).toNumber()] = (_this).f6;
              _nw357[(new BigNumber(2)).toNumber()] = (_2098_v84).dtor_cf9;
              _nw357[(new BigNumber(3)).toNumber()] = (_2099_v85).IsProperSubsetOf(_2100_v86);
              _nw357[(new BigNumber(4)).toNumber()] = _2083_v72;
              _nw357[(new BigNumber(5)).toNumber()] = _2083_v72;
              _nw357[(new BigNumber(6)).toNumber()] = (_this).f6;
              _nw357[(new BigNumber(7)).toNumber()] = ((_this).f9).isLessThan(new BigNumber((_2107_v93).length));
              _nw357[(new BigNumber(8)).toNumber()] = (_2106_v92).f11;
              _nw357[(new BigNumber(9)).toNumber()] = (_this).f6;
              _nw357[(new BigNumber(10)).toNumber()] = !((false) && (false));
              _nw357[(new BigNumber(11)).toNumber()] = (_this).f6;
              _nw357[(new BigNumber(12)).toNumber()] = (_2106_v92).f11;
              _nw357[(new BigNumber(13)).toNumber()] = (_2108_v94)[_module.__default.safeIndex(_this.f10, new BigNumber((_2108_v94).length))];
              _nw357[(new BigNumber(14)).toNumber()] = false;
              _2109_v95 = _nw357;
              let _index358 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2109_v95).length));
              (_2109_v95)[_index358] = (_this).f6;
              let _2110_v96;
              _2110_v96 = _dafny.Seq.of((_this).f9, new BigNumber((_2103_v89).length), new BigNumber((_2108_v94).length));
              let _index359 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length));
              let _index360 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2109_v95).length));
              let _rhs408 = _2095_v82;
              let _rhs409 = (_this).f9;
              let _rhs410 = (_this.f10).isLessThanOrEqualTo((((_2106_v92).f11) ? (_this.f10) : (_this.f10)));
              let _rhs411 = _2083_v72;
              let _rhs412 = _dafny.Seq.Concat(_2110_v96, _dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), ((_2111_v89) => function (_2112_i8) {
                return new BigNumber((_2111_v89).length);
              })(_2103_v89)));
              let _lhs273 = _2096_v83;
              let _lhs274 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length));
              let _lhs275 = _2109_v95;
              let _lhs276 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2109_v95).length));
              _2095_v82 = _rhs408;
              _lhs273[_lhs274] = _rhs409;
              _2083_v72 = _rhs410;
              _lhs275[_lhs276] = _rhs411;
              _2110_v96 = _rhs412;
              (globalState).f1 = ((_this).f9).minus((_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))]);
              let _2113_v97;
              let _nw358 = Array((new BigNumber(2)).toNumber()).fill([]);
              _2113_v97 = _nw358;
              let _2114_v98;
              _2114_v98 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_2113_v97);
              _2114_v98 = (_2114_v98).update((_dafny.ZERO).minus((_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))]), _2113_v97);
              let _2115_v99;
              let _nw359 = new _module.C1();
              _nw359.__ctor();
              _2115_v99 = _nw359;
              let _2116_v100;
              _2116_v100 = _dafny.MultiSet.fromElements(_2109_v95, _2109_v95, _2109_v95, _2109_v95);
              let _2117_v101;
              _2117_v101 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_2116_v100);
              let _2118_v102;
              _2118_v102 = _dafny.Map.Empty.slice().updateUnsafe(_2083_v72,_2109_v95);
              let _2119_v103;
              _2119_v103 = _dafny.Map.Empty.slice().updateUnsafe((_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))],_2116_v100);
              let _2120_v104;
              _2120_v104 = _dafny.Seq.of(_2116_v100);
              let _2121_v105;
              let _nw360 = Array((new BigNumber(28)).toNumber());
              _nw360[(_dafny.ZERO).toNumber()] = _2116_v100;
              _nw360[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_2109_v95, _2109_v95)).Intersect(_2116_v100);
              _nw360[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_2109_v95);
              _nw360[(new BigNumber(3)).toNumber()] = (((_2117_v101).contains((_this).f5)) ? ((_2117_v101).get((_this).f5)) : (_2116_v100));
              _nw360[(new BigNumber(4)).toNumber()] = (_2116_v100).update(_2109_v95, _module.__default.abs((_this).f9));
              _nw360[(new BigNumber(5)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_2109_v95, _2109_v95, _2109_v95, _2109_v95, _2109_v95);
              _nw360[(new BigNumber(7)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(8)).toNumber()] = (_2116_v100).Intersect(_dafny.MultiSet.fromElements(_2109_v95, _2109_v95, (((_2118_v102).contains(!(!((_this).f6)))) ? ((_2118_v102).get(!(!((_this).f6)))) : (_2109_v95)), _2109_v95));
              _nw360[(new BigNumber(9)).toNumber()] = (((_2119_v103).contains((_this).f9)) ? ((_2119_v103).get((_this).f9)) : (_2116_v100));
              _nw360[(new BigNumber(10)).toNumber()] = (_2120_v104)[_module.__default.safeIndex((_2115_v99).fm4(_2108_v94, _this.f10, (_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))], _dafny.MultiSet.fromElements(new BigNumber(-337)), globalState), new BigNumber((_2120_v104).length))];
              _nw360[(new BigNumber(11)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(12)).toNumber()] = (_2116_v100).Difference(_2116_v100);
              _nw360[(new BigNumber(13)).toNumber()] = (_2116_v100).Difference(_2116_v100);
              _nw360[(new BigNumber(14)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(15)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(16)).toNumber()] = (((_2119_v103).contains((_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))])) ? ((_2119_v103).get((_2096_v83)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_2096_v83).length))])) : (_2116_v100));
              _nw360[(new BigNumber(17)).toNumber()] = ((((_2119_v103).contains((_this).f9)) ? ((_2119_v103).get((_this).f9)) : (_2116_v100))).Intersect(_dafny.MultiSet.fromElements(_2109_v95));
              _nw360[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(_2109_v95);
              _nw360[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_2109_v95);
              _nw360[(new BigNumber(20)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(21)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(22)).toNumber()] = (_2116_v100).Difference(_dafny.MultiSet.fromElements(_2109_v95, _2109_v95));
              _nw360[(new BigNumber(23)).toNumber()] = _dafny.MultiSet.fromElements(_2109_v95, _2109_v95, _2109_v95, _2109_v95);
              _nw360[(new BigNumber(24)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(25)).toNumber()] = _2116_v100;
              _nw360[(new BigNumber(26)).toNumber()] = (_2116_v100).Union(_2116_v100);
              _nw360[(new BigNumber(27)).toNumber()] = _2116_v100;
              _2121_v105 = _nw360;
              let _index361 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2121_v105).length));
              (_2121_v105)[_index361] = _2116_v100;
              let _index362 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2121_v105).length));
              (_2121_v105)[_index362] = (_2116_v100).update(_2109_v95, _module.__default.abs(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_2122_i9) {
                return (_this).f5;
              }), _2103_v89)).length)));
            } else {
              let _2123_v106;
              _2123_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f5);
              _2123_v106 = (_2123_v106).update(true, new _dafny.CodePoint('s'.codePointAt(0)));
              _2083_v72 = (_this).f6;
              let _2124_v107;
              _2124_v107 = new _dafny.CodePoint('e'.codePointAt(0));
              _2124_v107 = new _dafny.CodePoint('t'.codePointAt(0));
              (_this).f10 = (_this).f9;
              r0 = _this.f10;
            }
            (globalState).f1 = _this.f10;
            r0 = new BigNumber(-304);
            let _2125_v108;
            let _nw361 = Array((new BigNumber(26)).toNumber());
            _2125_v108 = _nw361;
            let _2126_v109;
            _2126_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(128),!(_2083_v72));
            let _2127_v110;
            let _nw362 = new _module.C2();
            _nw362.__ctor((_this).f5, !((((_2126_v109).contains((_this).f9)) ? ((_2126_v109).get((_this).f9)) : ((_this).f6))));
            _2127_v110 = _nw362;
            let _index363 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_2125_v108).length));
            (_2125_v108)[_index363] = _2127_v110;
            let _2128_v111;
            _2128_v111 = _dafny.Map.Empty.slice().updateUnsafe(((!(_2083_v72)) ? ((_this).f9) : (new BigNumber(-367))),_2127_v110);
            let _index364 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_2125_v108).length));
            (_2125_v108)[_index364] = (((_2128_v111).contains((_this).f9)) ? ((_2128_v111).get((_this).f9)) : (_2127_v110));
          }
        }
      }
      r0 = _this.f10;
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _2129_v0;
      let _nw363 = new _module.C1();
      _nw363.__ctor();
      _2129_v0 = _nw363;
      let _2130_v1;
      _2130_v1 = _dafny.Seq.of(_this.f10);
      let _2131_v2;
      let _nw364 = Array((new BigNumber(11)).toNumber());
      _nw364[(_dafny.ZERO).toNumber()] = _this.f10;
      _nw364[(_dafny.ONE).toNumber()] = ((_this).f9).plus(new BigNumber((_2130_v1).length));
      _nw364[(new BigNumber(2)).toNumber()] = (_this).f9;
      _nw364[(new BigNumber(3)).toNumber()] = (_this).f9;
      _nw364[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_this.f10);
      _nw364[(new BigNumber(5)).toNumber()] = _this.f10;
      _nw364[(new BigNumber(6)).toNumber()] = (_this.f10).plus((_this).f9);
      _nw364[(new BigNumber(7)).toNumber()] = new BigNumber(37);
      _nw364[(new BigNumber(8)).toNumber()] = _this.f10;
      _nw364[(new BigNumber(9)).toNumber()] = _this.f10;
      _nw364[(new BigNumber(10)).toNumber()] = (_this).f9;
      _2131_v2 = _nw364;
      let _index365 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
      (_2131_v2)[_index365] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((p1).cardinality()))), _this.f10);
      let _2132_v3;
      _2132_v3 = true;
      let _2133_v4;
      let _nw365 = Array((new BigNumber(15)).toNumber()).fill(false);
      _2133_v4 = _nw365;
      let _2134_v5;
      _2134_v5 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2135_v6;
      _2135_v6 = _dafny.Seq.of((_this).f6, (_this).f6);
      let _2136_v7;
      _2136_v7 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f10,new BigNumber((_2135_v6).length)));
      let _2137_v8;
      _2137_v8 = _dafny.Map.Empty.slice().updateUnsafe((_2129_v0).fm14(globalState),(_this).f5);
      let _2138_v9;
      _2138_v9 = _module.D1.create_DC4(_this.f10);
      let _index366 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
      let _rhs413 = new BigNumber((_dafny.Seq.Concat(_2136_v7, _2136_v7)).length);
      let _rhs414 = true;
      let _rhs415 = _2133_v4;
      let _rhs416 = (((_2137_v8).contains((function (_pat_let36_0) {
        return function (_2141_dt__update__tmp_h0) {
          return function (_pat_let37_0) {
            return function (_2142_dt__update_hcf4_h0) {
              return _module.D1.create_DC4(_2142_dt__update_hcf4_h0);
            }(_pat_let37_0);
          }((_this).f9);
        }(_pat_let36_0);
      }(_2138_v9)).dtor_cf4)) ? ((_2137_v8).get((function (_pat_let34_0) {
        return function (_2139_dt__update__tmp_h1) {
          return function (_pat_let35_0) {
            return function (_2140_dt__update_hcf4_h1) {
              return _module.D1.create_DC4(_2140_dt__update_hcf4_h1);
            }(_pat_let35_0);
          }((_this).f9);
        }(_pat_let34_0);
      }(_2138_v9)).dtor_cf4)) : (new _dafny.CodePoint('b'.codePointAt(0))));
      let _lhs277 = _2131_v2;
      let _lhs278 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
      _lhs277[_lhs278] = _rhs413;
      _2132_v3 = _rhs414;
      _2133_v4 = _rhs415;
      _2134_v5 = _rhs416;
      _2132_v3 = !(((_this).f9).multipliedBy((_this).f9)).isEqualTo(((_dafny.ZERO).minus((_this).f9)).multipliedBy(new BigNumber(580)));
      let _2143_v10;
      _2143_v10 = _dafny.Seq.UnicodeFromString("jrsms");
      if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_2143_v10, _dafny.Seq.UnicodeFromString("sqhcfc")), _module.__default.fm1((_this).f9, p2, new BigNumber(338), globalState))) {
        if (false) {
          let _2144_v11;
          let _init56 = function (_2145_i0) {
            return _dafny.Set.fromElements(_module.D0.create_DC1(), _module.D0.create_DC1(), _module.D0.create_DC1(), _module.D0.create_DC1());
          };
          let _nw366 = Array((new BigNumber(4)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw366.length); _i0_56++) {
            _nw366[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2144_v11 = _nw366;
          let _2146_v12;
          _2146_v12 = _module.D0.create_DC1();
          let _2147_v13;
          _2147_v13 = _dafny.Set.fromElements(_2146_v12, _2146_v12);
          let _index367 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_2144_v11).length));
          (_2144_v11)[_index367] = _2147_v13;
          let _index368 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_2144_v11).length));
          let _rhs417 = (_2147_v13);
          let _rhs418 = (_this).f9;
          let _rhs419 = (_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))];
          let _lhs279 = _2144_v11;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_2144_v11).length));
          let _lhs281 = globalState;
          let _lhs282 = globalState;
          _lhs279[_lhs280] = _rhs417;
          _lhs281.f1 = _rhs418;
          _lhs282.f1 = _rhs419;
          (_this).f10 = (_module.__default.fm3(false, (_this).f6, globalState)).multipliedBy(_this.f10);
          let _2148_v14;
          _2148_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm27(_2143_v10, !(!(p2)), true, globalState),(_this).f9);
          let _2149_v15;
          _2149_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))],new BigNumber((_dafny.Seq.of(_2131_v2)).length));
          let _2150_v16;
          _2150_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), ((_2151_v5) => function (_2152_i1) {
            return _2151_v5;
          })(_2134_v5))).length),_this.f10);
          let _2153_v17;
          _2153_v17 = _module.D2.create_DC8(_2132_v3);
          _2148_v14 = (_2148_v14).update(_2150_v16, new BigNumber((_dafny.Seq.Concat(_module.__default.fm1(new BigNumber((_dafny.Set.fromElements((_this).f6)).length), (_2153_v17).dtor_cf9, (_this).f9, globalState), _2143_v10)).length));
          let _2154_v18;
          let _nw367 = new _module.C2();
          _nw367.__ctor((_this).f5, _2132_v3);
          _2154_v18 = _nw367;
          let _2155_v19;
          let _nw368 = Array((new BigNumber(29)).toNumber());
          _nw368[(_dafny.ZERO).toNumber()] = _2154_v18;
          _nw368[(_dafny.ONE).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(2)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(3)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(4)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(5)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(6)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(7)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(8)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(9)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(10)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(11)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(12)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(13)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(14)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(15)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(16)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(17)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(18)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(19)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(20)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(21)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(22)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(23)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(24)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(25)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(26)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(27)).toNumber()] = _2154_v18;
          _nw368[(new BigNumber(28)).toNumber()] = _2154_v18;
          _2155_v19 = _nw368;
          let _2156_v20;
          _2156_v20 = _dafny.Seq.of(_2155_v19);
          let _2157_v21;
          _2157_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v4,(_2156_v20)[_module.__default.safeIndex((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], new BigNumber((_2156_v20).length))]);
          let _2158_v22;
          _2158_v22 = _dafny.Map.Empty.slice().updateUnsafe(p2,((_dafny.Map.Empty.slice().updateUnsafe(_2133_v4,_2155_v19)).update(_2133_v4, _2155_v19)).update(_2133_v4, _2155_v19));
          let _2159_v23;
          _2159_v23 = _2135_v6;
          let _2160_v24;
          _2160_v24 = _dafny.Set.fromElements(_2135_v6, _2135_v6, _2135_v6, _dafny.Seq.update((_2159_v23), _module.__default.safeIndex((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], new BigNumber(((_2159_v23)).length)), (_this).f6));
          _2157_v21 = (((_2158_v22).contains((_dafny.Set.fromElements(_2135_v6)).IsSubsetOf(_2160_v24))) ? ((_2158_v22).get((_dafny.Set.fromElements(_2135_v6)).IsSubsetOf(_2160_v24))) : (_2157_v21));
          let _2161_v25;
          _2161_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,!(_2132_v3));
          let _2162_v26;
          _2162_v26 = _dafny.Set.fromElements(_2161_v25, _2161_v25, _2161_v25, _2161_v25);
          let _2163_v28;
          _2163_v28 = _dafny.Seq.of(_2162_v26, _2162_v26);
          _2162_v26 = (_2162_v26).Union((function () {
            let _coll65 = new _dafny.Set();
            for (const _compr_65 of (_2162_v26).Elements) {
              let _2164_v27 = _compr_65;
              if ((_2162_v26).contains(_2164_v27)) {
                _coll65.add(_2164_v27);
              }
            }
            return _coll65;
          }()).Intersect((_2163_v28)[_module.__default.safeIndex(_this.f10, new BigNumber((_2163_v28).length))]));
        } else {
          let _index369 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
          (_2131_v2)[_index369] = new BigNumber((_2135_v6).length);
          _2132_v3 = p0;
          _2132_v3 = (_this).f6;
          (globalState).f1 = ((_dafny.ZERO).minus((_this).fm7(true, globalState))).minus(new BigNumber(759));
          let _2165_v29;
          _2165_v29 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), function (_2166_i2) {
            return (_this).f9;
          }));
          let _2167_v30;
          let _nw369 = new _module.C0();
          _nw369.__ctor(p0, _2165_v29);
          _2167_v30 = _nw369;
          let _2168_v31;
          _2168_v31 = _dafny.Seq.of(((p0) ? (_2167_v30) : (_2167_v30)), _2167_v30, _2167_v30, _2167_v30);
          _2168_v31 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2168_v31, _dafny.Seq.of(_2167_v30)), _2168_v31);
        }
        let _index370 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
        (_2131_v2)[_index370] = (_dafny.ZERO).minus(_this.f10);
        _2132_v3 = true;
        let _2169_v32;
        _2169_v32 = _dafny.Seq.of(_2130_v1);
        let _2170_v33;
        let _nw370 = new _module.C0();
        _nw370.__ctor(p2, _2169_v32);
        _2170_v33 = _nw370;
        _2170_v33 = _2170_v33;
        if (!((_this).f9).isEqualTo((_dafny.ZERO).minus((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))]))) {
          let _2171_v34;
          _2171_v34 = _dafny.Set.fromElements((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], _this.f10);
          let _index371 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
          (_2131_v2)[_index371] = (_this).fm8(p2, _2132_v3, (_2171_v34).IsDisjointFrom(_2171_v34), globalState);
          (globalState).f4 = p1;
          let _2172_v35;
          let _nw371 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2172_v35 = _nw371;
          let _index372 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_2172_v35).length));
          (_2172_v35)[_index372] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), function (_2173_i3) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }), _2143_v10);
          let _2174_v36;
          _2174_v36 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f9),_2135_v6);
          let _index373 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_2172_v35).length));
          (_2172_v35)[_index373] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), ((_2175_v5) => function (_2176_i4) {
            return _2175_v5;
          })(_2134_v5)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-26)), function (_2177_i5) {
            return (_this).f5;
          }), _dafny.Seq.update(_dafny.Seq.update(_2143_v10, _module.__default.safeIndex((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], new BigNumber((_2143_v10).length)), (_this).f5), _module.__default.safeIndex(new BigNumber((_2174_v36).length), new BigNumber((_dafny.Seq.update(_2143_v10, _module.__default.safeIndex((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], new BigNumber((_2143_v10).length)), (_this).f5)).length)), new _dafny.CodePoint('r'.codePointAt(0)))));
          let _rhs420 = _2133_v4;
          let _rhs421 = (_this).f6;
          _2133_v4 = _rhs420;
          _2132_v3 = _rhs421;
          let _2178_v37;
          _2178_v37 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-664)), ((_2179_v5) => function (_2180_i6) {
            return _2179_v5;
          })(_2134_v5)), _dafny.Seq.UnicodeFromString("l")), _2143_v10);
          _2178_v37 = _2178_v37;
        } else {
          let _2181_v38;
          _2181_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,(_this).f9);
          let _rhs422 = _2132_v3;
          let _rhs423 = p2;
          let _rhs424 = _2181_v38;
          let _rhs425 = !((_this.f10).isEqualTo(new BigNumber((_2143_v10).length))) || (_dafny.areEqual(_2143_v10, _dafny.Seq.update(_2143_v10, _module.__default.safeIndex(new BigNumber(574), new BigNumber((_2143_v10).length)), _2134_v5)));
          _2132_v3 = _rhs422;
          _2132_v3 = _rhs423;
          _2181_v38 = _rhs424;
          _2132_v3 = _rhs425;
          let _index374 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
          (_2131_v2)[_index374] = (_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))];
          let _2182_v39;
          _2182_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2129_v0,_2133_v4);
          let _2183_v40;
          _2183_v40 = _2182_v39;
          let _2184_v41;
          _2184_v41 = _dafny.Map.Empty.slice().updateUnsafe((_2183_v40),_2143_v10);
          _2184_v41 = (_2184_v41).update(((_2182_v39).update(_2129_v0, _2133_v4)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2129_v0,_2133_v4)), _dafny.Seq.Concat(_2143_v10, _2143_v10));
          let _2185_v42;
          let _nw372 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _2185_v42 = _nw372;
          let _2186_v43;
          let _nw373 = new _module.C3();
          _nw373.__ctor(p0, p2);
          _2186_v43 = _nw373;
          let _2187_v44;
          _2187_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2138_v9,_2186_v43);
          let _2188_v45;
          _2188_v45 = _dafny.Seq.of(_2187_v44);
          let _2189_v46;
          _2189_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()));
          let _2190_v47;
          _2190_v47 = _dafny.Set.fromElements(_2187_v44, (_2188_v45)[_module.__default.safeIndex((((_2189_v46).contains(new BigNumber((_2143_v10).length))) ? ((_2189_v46).get(new BigNumber((_2143_v10).length))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,(_2135_v6)[_module.__default.safeIndex((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], new BigNumber((_2135_v6).length))])).length))), new BigNumber((_2188_v45).length))]);
          let _index375 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2185_v42).length));
          (_2185_v42)[_index375] = _2190_v47;
          let _index376 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2185_v42).length));
          (_2185_v42)[_index376] = ((_2190_v47).Intersect(_2190_v47)).Union(_2190_v47);
          let _2191_v48;
          let _nw374 = new _module.C3();
          _nw374.__ctor((_this).f6, false);
          _2191_v48 = _nw374;
        }
      } else {
        let _2192_v49;
        let _nw375 = new _module.C6();
        _nw375.__ctor();
        _2192_v49 = _nw375;
        _2192_v49 = _2192_v49;
        (_this).f10 = (_this).f9;
        if (!(p0) || (p0)) {
          _2132_v3 = true;
          let _2193_v50;
          _2193_v50 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2131_v2);
          _2193_v50 = (_2193_v50).update(_2132_v3, _2131_v2);
          let _2194_v51;
          _2194_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2143_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-377)), ((_2195_v5) => function (_2196_i7) {
            return _2195_v5;
          })(_2134_v5))),(_this).f6);
          _2194_v51 = (_2194_v51).update(_2143_v10, false);
          _2132_v3 = !(true);
          let _2197_v52;
          _2197_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),_2132_v3);
          let _2198_v53;
          _2198_v53 = _dafny.Map.Empty.slice().updateUnsafe(!((_2135_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(127), (_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], (_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))])).cardinality())), new BigNumber((_2135_v6).length))]),(_2197_v52).update((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))], (_this).f6));
          _2198_v53 = ((_2198_v53).Merge(_2198_v53)).update(!(_dafny.Seq.IsPrefixOf(_2130_v1, _2130_v1)), (_2197_v52).Merge(_2197_v52));
        } else {
          let _index377 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
          let _rhs426 = ((_this).f6) && (false);
          let _rhs427 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.fm3((_this).f6, _2132_v3, globalState), (_this).f9));
          let _lhs283 = _2131_v2;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length));
          _2132_v3 = _rhs426;
          _lhs283[_lhs284] = _rhs427;
          let _2199_v54;
          let _init57 = function (_2200_i8) {
            return _module.D11.create_DC19(new _dafny.CodePoint('j'.codePointAt(0)));
          };
          let _nw376 = Array((new BigNumber(29)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw376.length); _i0_57++) {
            _nw376[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2199_v54 = _nw376;
          let _2201_v55;
          _2201_v55 = _module.D11.create_DC19(_2134_v5);
          let _index378 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2199_v54).length));
          (_2199_v54)[_index378] = _2201_v55;
          let _2202_v56;
          _2202_v56 = _dafny.Seq.of(_2130_v1, _2130_v1);
          let _2203_v57;
          let _nw377 = new _module.C0();
          _nw377.__ctor((_this).f6, _2202_v56);
          _2203_v57 = _nw377;
          let _2204_v58;
          _2204_v58 = _dafny.Seq.of(_2203_v57);
          let _2205_v59;
          _2205_v59 = _dafny.Set.fromElements(new BigNumber(102));
          let _2206_v60;
          _2206_v60 = _dafny.Seq.of(_2205_v59);
          let _index379 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2199_v54).length));
          let _rhs428 = _2132_v3;
          let _rhs429 = (_this).f9;
          let _rhs430 = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(((_2132_v3) ? (_2204_v58) : (_2204_v58)), _dafny.Seq.Concat(_2204_v58, _2204_v58)), _module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f9, new BigNumber((_2206_v60).length)), new BigNumber((_dafny.Seq.Concat(((_2132_v3) ? (_2204_v58) : (_2204_v58)), _dafny.Seq.Concat(_2204_v58, _2204_v58))).length)), _2203_v57), _module.__default.safeIndex(_this.f10, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(((_2132_v3) ? (_2204_v58) : (_2204_v58)), _dafny.Seq.Concat(_2204_v58, _2204_v58)), _module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f9, new BigNumber((_2206_v60).length)), new BigNumber((_dafny.Seq.Concat(((_2132_v3) ? (_2204_v58) : (_2204_v58)), _dafny.Seq.Concat(_2204_v58, _2204_v58))).length)), _2203_v57)).length)), _2203_v57))).cardinality());
          let _rhs431 = (_this).f6;
          let _rhs432 = _2201_v55;
          let _lhs285 = globalState;
          let _lhs286 = globalState;
          let _lhs287 = _2199_v54;
          let _lhs288 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2199_v54).length));
          _2132_v3 = _rhs428;
          _lhs285.f1 = _rhs429;
          _lhs286.f1 = _rhs430;
          _2132_v3 = _rhs431;
          _lhs287[_lhs288] = _rhs432;
          _2131_v2 = _2131_v2;
          let _2207_v61;
          let _init58 = ((_2208_v10) => function (_2209_i9) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tbknyswhp"), _2208_v10);
          })(_2143_v10);
          let _nw378 = Array((new BigNumber(19)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw378.length); _i0_58++) {
            _nw378[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2207_v61 = _nw378;
          let _index380 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_2207_v61).length));
          (_2207_v61)[_index380] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("penmrr"), _2143_v10);
          let _index381 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_2207_v61).length));
          (_2207_v61)[_index381] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kq"), _dafny.Seq.UnicodeFromString("onejtsx"));
          let _2210_v62;
          _2210_v62 = _module.D2.create_DC8(_2132_v3);
          let _rhs433 = (_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))];
          let _rhs434 = (((_this.f10).isLessThan((_2131_v2)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_2131_v2).length))])) ? (_2138_v9) : ((_this).fm5(_2210_v62, globalState)));
          let _lhs289 = globalState;
          _lhs289.f1 = _rhs433;
          _2138_v9 = _rhs434;
        }
        _2134_v5 = ((p0) ? (((p2) ? (_2134_v5) : ((_this).f5))) : ((_this).f5));
        let _2211_v63;
        _2211_v63 = _2131_v2;
        let _2212_v64;
        _2212_v64 = _module.D12.create_DC23(_2211_v63);
        let _pat_let_tv29 = _2131_v2;
        let _pat_let_tv30 = _2211_v63;
        let _2213_v65;
        let _nw379 = Array((new BigNumber(13)).toNumber());
        _nw379[(_dafny.ZERO).toNumber()] = _2212_v64;
        _nw379[(_dafny.ONE).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(2)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(3)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(4)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(5)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(6)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(7)).toNumber()] = function (_pat_let38_0) {
          return function (_2216_dt__update__tmp_h4) {
            return function (_pat_let41_0) {
              return function (_2217_dt__update_hcf29_h1) {
                return _module.D12.create_DC23(_2217_dt__update_hcf29_h1);
              }(_pat_let41_0);
            }(_pat_let_tv30);
          }(_pat_let38_0);
        }(function (_pat_let39_0) {
          return function (_2214_dt__update__tmp_h3) {
            return function (_pat_let40_0) {
              return function (_2215_dt__update_hcf29_h0) {
                return _module.D12.create_DC23(_2215_dt__update_hcf29_h0);
              }(_pat_let40_0);
            }(_pat_let_tv29);
          }(_pat_let39_0);
        }(_2212_v64));
        _nw379[(new BigNumber(8)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(9)).toNumber()] = _2212_v64;
        _nw379[(new BigNumber(10)).toNumber()] = _module.D12.create_DC23(_2211_v63);
        _nw379[(new BigNumber(11)).toNumber()] = _module.D12.create_DC23(_2211_v63);
        _nw379[(new BigNumber(12)).toNumber()] = _2212_v64;
        _2213_v65 = _nw379;
        let _2218_v66;
        let _nw380 = Array((new BigNumber(2)).toNumber());
        _nw380[(_dafny.ZERO).toNumber()] = _2213_v65;
        _nw380[(_dafny.ONE).toNumber()] = _2213_v65;
        _2218_v66 = _nw380;
        _2218_v66 = _2218_v66;
      }
      let _2219_v67;
      _2219_v67 = _dafny.Set.fromElements((_this).f9);
      let _2220_v68;
      _2220_v68 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.MultiSet.fromElements((_this).fm6(_2130_v1, globalState))).contains((_module.D16.create_DC32(p2)).dtor_cf41),_2219_v67);
      _2220_v68 = (_2220_v68).update(false, _2219_v67);
      let _index382 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2133_v4).length));
      (_2133_v4)[_index382] = p2;
      let _index383 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2133_v4).length));
      (_2133_v4)[_index383] = ((_2130_v1)[_module.__default.safeIndex(_this.f10, new BigNumber((_2130_v1).length))]).isLessThan(_this.f10);
      return;
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Seq.of();
      let _2221_v0;
      _2221_v0 = _module.D0.create_DC3(_module.D0.create_DC0((_this).f6));
      _2221_v0 = (((_this).f6) ? (_2221_v0) : (_2221_v0));
      if ((_this).f6) {
        let _2222_v1;
        _2222_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,false);
        _2222_v1 = (_2222_v1).Merge(_2222_v1);
        let _2223_v2;
        let _nw381 = new _module.C3();
        _nw381.__ctor((_this).f6, (_this).f6);
        _2223_v2 = _nw381;
        let _2224_v3;
        _2224_v3 = _dafny.Seq.of(p0, p0);
        let _2225_v4;
        _2225_v4 = _dafny.MultiSet.fromElements((_2224_v3)[_module.__default.safeIndex(new BigNumber((_2224_v3).length), new BigNumber((_2224_v3).length))], new BigNumber(-856));
        _2225_v4 = _2225_v4;
        let _2226_v5;
        _2226_v5 = new _dafny.CodePoint('n'.codePointAt(0));
        _2226_v5 = _2226_v5;
        (globalState).f1 = ((!((_2223_v2).f13)) ? (_this.f10) : (new BigNumber(-730)));
      } else {
        let _2227_v6;
        _2227_v6 = _dafny.Seq.UnicodeFromString("o");
        if (!_dafny.Seq.contains((_module.D2.create_DC6(_2227_v6)).dtor_cf6, (_this).f5)) {
          let _2228_v7;
          _2228_v7 = false;
          _2228_v7 = !(true);
          r0 = p0;
          let _2229_v8;
          let _nw382 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _2229_v8 = _nw382;
          let _2230_v9;
          let _nw383 = Array((new BigNumber(10)).toNumber()).fill(false);
          _2230_v9 = _nw383;
          let _2231_v10;
          _2231_v10 = _dafny.Seq.of(_2230_v9, _2230_v9, _2230_v9);
          let _2232_v11;
          _2232_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2229_v8,(_2231_v10)[_module.__default.safeIndex((_this).fm8(_2228_v7, _2228_v7, (_this).f6, globalState), new BigNumber((_2231_v10).length))]);
          let _2233_v12;
          _2233_v12 = _module.D17.create_DC33(_2232_v11);
          _2232_v11 = (_2233_v12).dtor_cf42;
          let _2234_v13;
          let _nw384 = new _module.C6();
          _nw384.__ctor();
          _2234_v13 = _nw384;
          let _2235_v14;
          let _nw385 = Array((new BigNumber(2)).toNumber()).fill([]);
          _2235_v14 = _nw385;
          let _index384 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_2235_v14).length));
          (_2235_v14)[_index384] = _2230_v9;
          let _2236_v15;
          _2236_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2227_v6);
          let _2237_v16;
          _2237_v16 = _dafny.Seq.of(_2227_v6, _2227_v6);
          let _2238_v17;
          _2238_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2237_v16)[_module.__default.safeIndex((_this).f9, new BigNumber((_2237_v16).length))]);
          let _2239_v18;
          _2239_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _2240_v19;
          _2240_v19 = _2239_v18;
          let _2241_v20;
          _2241_v20 = _dafny.Seq.of(_2228_v7);
          let _2242_v21;
          _2242_v21 = _dafny.Seq.of(_2241_v20, _module.__default.fm2((_this).f6, false, globalState), _2241_v20, _2241_v20);
          let _2243_v22;
          _2243_v22 = _module.D12.create_DC22(_2240_v19, _2242_v21, (_this).f5, _2228_v7, _module.__default.fm1(new BigNumber((_2227_v6).length), _2228_v7, p0, globalState));
          let _2244_v23;
          let _nw386 = Array((new BigNumber(6)).toNumber());
          _nw386[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((((_2236_v15).contains(p0)) ? ((_2236_v15).get(p0)) : (_dafny.Seq.UnicodeFromString("v"))), _dafny.Seq.UnicodeFromString("j"));
          _nw386[(_dafny.ONE).toNumber()] = (((_2238_v17).contains(_2228_v7)) ? ((_2238_v17).get(_2228_v7)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_2245_i0) {
            return (_this).f5;
          })));
          _nw386[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_2227_v6, _2227_v6);
          _nw386[(new BigNumber(3)).toNumber()] = _2227_v6;
          _nw386[(new BigNumber(4)).toNumber()] = _2227_v6;
          _nw386[(new BigNumber(5)).toNumber()] = (_2243_v22).dtor_cf28;
          _2244_v23 = _nw386;
          let _index385 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2244_v23).length));
          (_2244_v23)[_index385] = _2227_v6;
          let _index386 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_2235_v14).length));
          let _index387 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2244_v23).length));
          let _rhs435 = _2230_v9;
          let _rhs436 = _2227_v6;
          let _lhs290 = _2235_v14;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_2235_v14).length));
          let _lhs292 = _2244_v23;
          let _lhs293 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2244_v23).length));
          _lhs290[_lhs291] = _rhs435;
          _lhs292[_lhs293] = _rhs436;
        } else {
          let _2246_v24;
          let _nw387 = new _module.C2();
          _nw387.__ctor((_2227_v6)[_module.__default.safeIndex((_this).f9, new BigNumber((_2227_v6).length))], _module.__default.fm0((_this).f6, globalState));
          _2246_v24 = _nw387;
          let _2247_v25;
          _2247_v25 = true;
          _2247_v25 = _2247_v25;
          let _2248_v26;
          _2248_v26 = _dafny.MultiSet.fromElements((_this).f6);
          let _2249_v27;
          _2249_v27 = _dafny.Seq.of((_this).f6, _2247_v25, _2247_v25);
          let _2250_v28;
          let _nw388 = Array((new BigNumber(20)).toNumber());
          _nw388[(_dafny.ZERO).toNumber()] = _2248_v26;
          _nw388[(_dafny.ONE).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(2)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(3)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(4)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(5)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(6)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(7)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_2249_v27);
          _nw388[(new BigNumber(9)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(10)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(11)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(12)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(13)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(14)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(15)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(16)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(17)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(18)).toNumber()] = _2248_v26;
          _nw388[(new BigNumber(19)).toNumber()] = _2248_v26;
          _2250_v28 = _nw388;
          let _2251_v29;
          _2251_v29 = _module.D18.create_DC36(_2250_v28);
          let _2252_v30;
          let _nw389 = Array((new BigNumber(20)).toNumber());
          _nw389[(_dafny.ZERO).toNumber()] = (_2251_v29).dtor_cf45;
          _nw389[(_dafny.ONE).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(2)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(3)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(4)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(5)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(6)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(7)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(8)).toNumber()] = ((true) ? (_2250_v28) : (_2250_v28));
          _nw389[(new BigNumber(9)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(10)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(11)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(12)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(13)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(14)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(15)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(16)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(17)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(18)).toNumber()] = _2250_v28;
          _nw389[(new BigNumber(19)).toNumber()] = _2250_v28;
          _2252_v30 = _nw389;
          let _index388 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2252_v30).length));
          (_2252_v30)[_index388] = _2250_v28;
          let _index389 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2252_v30).length));
          (_2252_v30)[_index389] = _2250_v28;
          _2247_v25 = true;
          let _rhs437 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_2247_v25)).length), (_dafny.ZERO).minus(p0));
          let _rhs438 = new BigNumber(287);
          let _lhs294 = _this;
          r1 = _rhs437;
          _lhs294.f10 = _rhs438;
        }
        let _2253_v31;
        let _nw390 = Array((new BigNumber(7)).toNumber());
        _nw390[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw390[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f9);
        _nw390[(new BigNumber(2)).toNumber()] = new BigNumber(792);
        _nw390[(new BigNumber(3)).toNumber()] = new BigNumber((_2227_v6).length);
        _nw390[(new BigNumber(4)).toNumber()] = _this.f10;
        _nw390[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw390[(new BigNumber(6)).toNumber()] = _this.f10;
        _2253_v31 = _nw390;
        let _index390 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_2253_v31).length));
        (_2253_v31)[_index390] = (_this).fm7((_this).f6, globalState);
        let _index391 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_2253_v31).length));
        (_2253_v31)[_index391] = (new BigNumber(837)).multipliedBy((_this).f9);
        let _2254_v32;
        let _nw391 = new _module.C5();
        _nw391.__ctor(new BigNumber(109), (_this).f5, (_this).f6);
        _2254_v32 = _nw391;
        _2254_v32 = _2254_v32;
        let _2255_v33;
        let _nw392 = new _module.C6();
        _nw392.__ctor();
        _2255_v33 = _nw392;
        let _2256_v34;
        let _nw393 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _2256_v34 = _nw393;
        _2256_v34 = _2256_v34;
      }
      let _2257_v35;
      _2257_v35 = _dafny.Set.fromElements(_module.__default.fm0((_this).f6, globalState));
      _2257_v35 = _2257_v35;
      let _2258_v36;
      _2258_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_2259_i1) {
        return (_this).f5;
      }),(_this).f5);
      let _2260_v37;
      _2260_v37 = _dafny.Seq.UnicodeFromString("adh");
      let _2261_v38;
      _2261_v38 = _dafny.Seq.of(_2260_v37, _2260_v37);
      _2258_v36 = (_2258_v36).update((_2261_v38)[_module.__default.safeIndex(_this.f10, new BigNumber((_2261_v38).length))], new _dafny.CodePoint('k'.codePointAt(0)));
      r1 = p0;
      let _2262_v39;
      let _nw394 = Array((new BigNumber(16)).toNumber()).fill(false);
      _2262_v39 = _nw394;
      let _2263_v40;
      _2263_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _2264_v41;
      _2264_v41 = _module.D1.create_DC4(new BigNumber((_2263_v40).length));
      let _2265_v42;
      _2265_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2264_v41,_module.__default.fm0(false, globalState));
      let _2266_v43;
      let _2267_v44;
      let _out48;
      let _out49;
      let _outcollector21 = _module.__default.m0(_2262_v39, (_this).f5, (((_2265_v42).contains(_2264_v41)) ? ((_2265_v42).get(_2264_v41)) : ((_this).f6)), new _dafny.CodePoint('i'.codePointAt(0)), globalState);
      _out48 = _outcollector21[0];
      _out49 = _outcollector21[1];
      _2266_v43 = _out48;
      _2267_v44 = _out49;
      r0 = _this.f10;
      r1 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_2260_v37).length), _this.f10), _module.__default.fm3((_this).f6, _2267_v44, globalState));
      let _2268_v45;
      _2268_v45 = _module.D6.create_DC14(_2267_v44);
      let _pat_let_tv31 = p0;
      r2 = new BigNumber((function (_source28) {
        if (_source28.is_DC14) {
          let _2269___mcc_h0 = (_source28).cf15;
          let _2270_cf15 = _2269___mcc_h0;
          if (true) {
            return _dafny.Seq.of((_this).f9, _this.f10, (_this).f9, _pat_let_tv31);
          } else {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_2271_i2) {
              return _this.f10;
            });
          }
        } else {
          let _2272___mcc_h1 = (_source28).cf14;
          let _2273_cf14 = _2272___mcc_h1;
          return _2273_cf14;
        }
      }(_2268_v45)).length);
      let _2274_v46;
      _2274_v46 = _dafny.Seq.of(_2262_v39);
      r3 = _2274_v46;
      return [r0, r1, r2, r3];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
      this.f7 = [];
      this._f8 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T4, _module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f7, f8, f5, f6) {
      let _this = this;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((function () {
        let _coll66 = new _dafny.Set();
        for (const _compr_66 of (_dafny.Seq.of(_module.D0.create_DC0(false), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0(true), _module.D0.create_DC0((_this).f6))).Elements) {
          let _2275_v0 = _compr_66;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0(false), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0(true), _module.D0.create_DC0((_this).f6)), _2275_v0)) {
            _coll66.add(_2275_v0);
          }
        }
        return _coll66;
      }()).Union(_dafny.Set.fromElements(_module.D0.create_DC0(true), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0((_this).f6), _module.D0.create_DC0(!((_this).f6))))).IsDisjointFrom(function () {
        let _coll67 = new _dafny.Set();
        for (const _compr_67 of (function () {
          let _coll68 = new _dafny.Set();
          for (const _compr_68 of (_dafny.MultiSet.fromElements(_module.D0.create_DC0(!((_this).f6)))).Elements) {
            let _2276_v1 = _compr_68;
            if ((_dafny.MultiSet.fromElements(_module.D0.create_DC0(!((_this).f6)))).contains(_2276_v1)) {
              _coll68.add(_2276_v1);
            }
          }
          return _coll68;
        }()).Elements) {
          let _2277_v2 = _compr_67;
          if ((function () {
            let _coll69 = new _dafny.Set();
            for (const _compr_69 of (_dafny.MultiSet.fromElements(_module.D0.create_DC0(!((_this).f6)))).Elements) {
              let _2278_v1 = _compr_69;
              if ((_dafny.MultiSet.fromElements(_module.D0.create_DC0(!((_this).f6)))).contains(_2278_v1)) {
                _coll69.add(_2278_v1);
              }
            }
            return _coll69;
          }()).contains(_2277_v2)) {
            _coll67.add(_2277_v2);
          }
        }
        return _coll67;
      }());
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return !(_module.__default.safeModuloInt(new BigNumber(-966), new BigNumber((_dafny.Seq.of((_this).f6)).length))).isEqualTo(((_module.D2.create_DC7(new BigNumber((_dafny.Seq.of(new BigNumber(-661))).length), new BigNumber(-667))).dtor_cf8).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f6)),(_this).f6)).length)));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(434);
    };
    fm6(p0, globalState) {
      let _this = this;
      return !((_this).f6);
    };
    fm7(p0, globalState) {
      let _this = this;
      let _source29 = _module.D2.create_DC7(new BigNumber(899), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f6, (_this).f6))).length)));
      if (_source29.is_DC7) {
        let _2279___mcc_h0 = (_source29).cf7;
        let _2280___mcc_h1 = (_source29).cf8;
        let _2281_cf8 = _2280___mcc_h1;
        let _2282_cf7 = _2279___mcc_h0;
        return new BigNumber(((_dafny.Set.fromElements((_this).f6)).Union(_dafny.Set.fromElements(!((_this).f6)))).length);
      } else if (_source29.is_DC8) {
        let _2283___mcc_h2 = (_source29).cf9;
        let _2284_cf9 = _2283___mcc_h2;
        return new BigNumber((function () {
          let _coll70 = new _dafny.Map();
          for (const _compr_70 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ngprv")).length),false)).length))).Elements) {
            let _2285_v0 = _compr_70;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ngprv")).length),false)).length))).contains(_2285_v0)) {
              _coll70.push([_module.__default.safeDivisionInt(_2285_v0, new BigNumber(266)),_2284_cf9]);
            }
          }
          return _coll70;
        }()).length);
      } else if (_source29.is_DC6) {
        let _2286___mcc_h3 = (_source29).cf6;
        let _2287_cf6 = _2286___mcc_h3;
        return (new BigNumber(-427)).plus(new BigNumber(94));
      } else {
        let _2288___mcc_h4 = (_source29).cf10;
        let _2289_cf10 = _2288___mcc_h4;
        return _module.__default.safeDivisionInt(new BigNumber(409), new BigNumber(250));
      }
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!((_this).f6)), _dafny.Seq.of((_this).f6, (_this).f6)), _dafny.Seq.of(false))).length);
    };
    fm5(p0, globalState) {
      let _this = this;
      return _module.D1.create_DC4(_module.__default.safeDivisionInt(new BigNumber(113), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-300))).length))).length)));
    };
    fm11(p0, p1, p2, globalState) {
      let _this = this;
      if ((_dafny.Set.fromElements((_this).f6)).IsSubsetOf(_dafny.Set.fromElements((_this).f6, (_this).f6, (_this).f6))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(609),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-597)), function (_2290_i0) {
          return new BigNumber((_dafny.MultiSet.fromElements(true, (_this).f6)).cardinality());
        })).length));
      } else {
        return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll71 = new _dafny.Set();
          for (const _compr_71 of _dafny.IntegerRange(new BigNumber(19), new BigNumber(9))) {
            let _2291_v0 = _compr_71;
            if (((new BigNumber(19)).isLessThanOrEqualTo(_2291_v0)) && ((_2291_v0).isLessThan(new BigNumber(9)))) {
              _coll71.add((_2291_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(585))).cardinality())));
            }
          }
          return _coll71;
        }()).length),new BigNumber((_dafny.Set.fromElements((_this).f6, (_this).f6)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(168),new BigNumber(385)));
      }
    };
    fm12(globalState) {
      let _this = this;
      return (_this).f6;
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D2.Default();
      let r1 = _dafny.MultiSet.Empty;
      let _index392 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length));
      (p0)[_index392] = p3;
      let _index393 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length));
      (p0)[_index393] = _module.__default.fm3(!(p2), ((p2) ? (p2) : ((_this).f6)), globalState);
      let _2292_v0;
      _2292_v0 = true;
      let _2293_v1;
      _2293_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f5);
      let _2294_v2;
      _2294_v2 = _module.D0.create_DC2(_2293_v1, (_this).f5);
      let _2295_v3;
      _2295_v3 = _dafny.Seq.UnicodeFromString("ghkphwmku");
      let _2296_v4;
      _2296_v4 = _dafny.MultiSet.fromElements(p3, p3);
      let _2297_v5;
      _2297_v5 = _dafny.Seq.of(_2296_v4, _2296_v4);
      let _2298_v6;
      _2298_v6 = _dafny.MultiSet.fromElements((_this).f6, p2, (_2296_v4).IsSubsetOf((_2297_v5)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], new BigNumber((_2297_v5).length))]), (_2292_v0) && ((_this).f6));
      let _rhs439 = _2292_v0;
      let _rhs440 = _2294_v2;
      let _rhs441 = new BigNumber((_2298_v6).cardinality());
      let _rhs442 = _2295_v3;
      let _lhs295 = globalState;
      _2292_v0 = _rhs439;
      _2294_v2 = _rhs440;
      _lhs295.f1 = _rhs441;
      _2295_v3 = _rhs442;
      let _2299_v7;
      let _nw395 = new _module.C1();
      _nw395.__ctor();
      _2299_v7 = _nw395;
      let _2300_v8;
      _2300_v8 = _dafny.Seq.of(p3, (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]);
      let _2301_v9;
      _2301_v9 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))],_2300_v8);
      let _2302_v10;
      _2302_v10 = _module.D6.create_DC13((((_2301_v9).contains((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))])) ? ((_2301_v9).get((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))])) : (_2300_v8)));
      let _2303_v11;
      _2303_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2302_v10,p3);
      let _hi12 = _module.__default.safeModuloInt(p3, (_dafny.ZERO).minus((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]));
      for (let _2304_i0 = (((_2303_v11).contains(_2302_v10)) ? ((_2303_v11).get(_2302_v10)) : (new BigNumber(415))); _2304_i0.isLessThan(_hi12); _2304_i0 = _2304_i0.plus(_dafny.ONE)) {
        let _index394 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length));
        (p0)[_index394] = (_2300_v8)[_module.__default.safeIndex(new BigNumber(-641), new BigNumber((_2300_v8).length))];
        let _2305_v12;
        let _init59 = function (_2306_i1) {
          return (_this).f5;
        };
        let _nw396 = Array((new BigNumber(9)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw396.length); _i0_59++) {
          _nw396[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2305_v12 = _nw396;
        let _nw397 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2305_v12 = _nw397;
        _2300_v8 = _2300_v8;
        let _2307_v13;
        let _nw398 = new _module.C3();
        _nw398.__ctor((_this).f6, p2);
        _2307_v13 = _nw398;
      }
      if (_2292_v0) {
        if ((p2) && (_2292_v0)) {
          let _2308_v14;
          let _nw399 = new _module.C7();
          _nw399.__ctor((_dafny.ZERO).minus(p3), (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], (_this).f5, _2292_v0);
          _2308_v14 = _nw399;
          let _2309_v15;
          _2309_v15 = new _dafny.CodePoint('q'.codePointAt(0));
          _2309_v15 = _2309_v15;
          let _2310_v16;
          _2310_v16 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_2308_v14).f9);
          let _2311_v17;
          _2311_v17 = _dafny.MultiSet.fromElements(_2295_v3);
          _2292_v0 = (_2311_v17).IsSubsetOf(_module.__default.fm61(_2310_v16, globalState));
          let _index395 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length));
          (p0)[_index395] = new BigNumber(567);
          let _2312_v18;
          _2312_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm12(globalState),_2292_v0);
          _2312_v18 = (_2312_v18).update(_2292_v0, _2292_v0);
        } else {
          let _2313_v19;
          let _nw400 = Array((new BigNumber(20)).toNumber());
          _2313_v19 = _nw400;
          let _2314_v20;
          let _nw401 = new _module.C2();
          _nw401.__ctor(_module.__default.fm22(_dafny.Set.fromElements((_this).f6), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), p2, (p1)[_module.__default.safeIndex(p3, new BigNumber((p1).length))], globalState), !(_2292_v0));
          _2314_v20 = _nw401;
          let _index396 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2313_v19).length));
          (_2313_v19)[_index396] = _2314_v20;
          let _index397 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2313_v19).length));
          (_2313_v19)[_index397] = _2314_v20;
          let _2315_v21;
          let _nw402 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _2315_v21 = _nw402;
          let _2316_v22;
          _2316_v22 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))],_dafny.Seq.UnicodeFromString("lyknjnvny"));
          let _index398 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2315_v21).length));
          (_2315_v21)[_index398] = (_2316_v22).update(p3, _2295_v3);
          let _index399 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_2315_v21).length));
          (_2315_v21)[_index399] = (_2316_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,_2295_v3));
          let _2317_v23;
          _2317_v23 = _module.D11.create_DC20(_2292_v0, p3);
          let _rhs443 = p2;
          let _rhs444 = (_2317_v23).dtor_cf21;
          let _rhs445 = p3;
          let _lhs296 = globalState;
          _2292_v0 = _rhs443;
          _2292_v0 = _rhs444;
          _lhs296.f1 = _rhs445;
          _2292_v0 = ((_this).f6) === (_2292_v0);
          let _2318_v24;
          _2318_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2292_v0,(new BigNumber((_2295_v3).length)).plus(new BigNumber(-220)));
          _2318_v24 = (_2318_v24).update((_this).f6, (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]);
        }
        let _2319_v25;
        _2319_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], (_dafny.ZERO).minus(p3)),p3);
        _2319_v25 = (_2319_v25).update((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]);
        if ((_2296_v4).IsProperSubsetOf((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("smcp")).length), (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], new BigNumber((_2298_v6).cardinality()))).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f6)).length), _module.__default.abs(p3)))) {
          _2295_v3 = _2295_v3;
          _2292_v0 = false;
          let _rhs446 = (p3).multipliedBy(p3);
          let _rhs447 = (p3).plus(((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]).multipliedBy((_dafny.ZERO).minus(p3)));
          let _lhs297 = globalState;
          let _lhs298 = globalState;
          _lhs297.f1 = _rhs446;
          _lhs298.f1 = _rhs447;
          let _2320_v26;
          _2320_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f6);
          _2320_v26 = _2320_v26;
          (globalState).f1 = p3;
        } else {
          let _2321_v27;
          let _init60 = ((_2322_p3) => function (_2323_i2) {
            return (_2323_i2).minus(_2322_p3);
          })(p3);
          let _nw403 = Array((new BigNumber(19)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw403.length); _i0_60++) {
            _nw403[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2321_v27 = _nw403;
          let _rhs448 = (p3).multipliedBy(new BigNumber(((_module.D6.create_DC13(_2300_v8)).dtor_cf14).length));
          let _rhs449 = _this.f7;
          let _rhs450 = (_this).f6;
          let _rhs451 = p0;
          let _lhs299 = globalState;
          let _lhs300 = _this;
          _lhs299.f1 = _rhs448;
          _lhs300.f7 = _rhs449;
          _2292_v0 = _rhs450;
          _2321_v27 = _rhs451;
          _2292_v0 = (_this).f6;
          let _2324_v28;
          _2324_v28 = _dafny.Set.fromElements(_2292_v0);
          let _2325_v29;
          _2325_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2324_v28).Union(_2324_v28)).length),_2292_v0);
          _2325_v29 = (_2325_v29).update((new BigNumber(-999)).minus(p3), p2);
          let _2326_v30;
          _2326_v30 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm1((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], !(p2), (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], globalState)).length));
          _2326_v30 = _2326_v30;
          _2292_v0 = p2;
        }
        let _2327_v31;
        _2327_v31 = new _dafny.CodePoint('g'.codePointAt(0));
        _2327_v31 = (_this).f5;
        _2292_v0 = _2292_v0;
      } else {
        let _2328_v32;
        let _nw404 = new _module.C6();
        _nw404.__ctor();
        _2328_v32 = _nw404;
        let _2329_v33;
        _2329_v33 = _module.D14.create_DC26((_this).f6, p3, _module.__default.fm53(_2295_v3, (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], globalState));
        (globalState).f1 = (_2300_v8)[_module.__default.safeIndex((_2329_v33).dtor_cf33, new BigNumber((_2300_v8).length))];
        _2292_v0 = (_this).f6;
        let _2330_v34;
        _2330_v34 = _dafny.Set.fromElements(p2);
        let _2331_v35;
        _2331_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_2330_v34).IsDisjointFrom(_2330_v34));
        let _2332_v36;
        _2332_v36 = _dafny.Seq.of(_2331_v35, _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p2));
        let _rhs452 = ((_2331_v35).update(_this.f7, (_this).f6)).Merge((_2332_v36)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))], new BigNumber((_2332_v36).length))]);
        let _rhs453 = false;
        let _rhs454 = _2292_v0;
        let _rhs455 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,p2)).update(_module.__default.fm0((_this).f6, globalState), (p1)[_module.__default.safeIndex(p3, new BigNumber((p1).length))])).length);
        let _lhs301 = globalState;
        _2331_v35 = _rhs452;
        _2292_v0 = _rhs453;
        _2292_v0 = _rhs454;
        _lhs301.f1 = _rhs455;
        let _index400 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length));
        (p0)[_index400] = (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))];
      }
      let _2333_v37;
      _2333_v37 = _module.D11.create_DC20(!(true), (p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))]);
      let _2334_v38;
      _2334_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2300_v8,((_2333_v37).dtor_cf22).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("ymtkf")).length)));
      (globalState).f1 = new BigNumber((_2334_v38).length);
      let _2335_v39;
      _2335_v39 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
      let _2336_v40;
      _2336_v40 = _dafny.Map.Empty.slice().updateUnsafe((((_2335_v39).contains(!(p2))) ? ((_2335_v39).get(!(p2))) : ((p0)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((p0).length))])),_2292_v0);
      r0 = _module.__default.fm62(_module.D15.create_DC27(_dafny.MultiSet.fromElements(false, p2)), (((_2336_v40).contains(p3)) ? ((_2336_v40).get(p3)) : (_2292_v0)), _2292_v0, _dafny.Seq.Concat(_2300_v8, _2300_v8), globalState);
      r1 = _2298_v6;
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let _2337_v0;
      _2337_v0 = _dafny.Seq.UnicodeFromString("duo");
      _2337_v0 = p0;
      let _2338_v1;
      _2338_v1 = false;
      _2338_v1 = _2338_v1;
      let _2339_v2;
      _2339_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2338_v1);
      let _source30 = (_2339_v2).Merge(_2339_v2);
      let _2340___mcc_h0 = _source30;
      let _2341_cf12 = _2340___mcc_h0;
      _2338_v1 = _2338_v1;
      (globalState).f1 = p1;
      let _2342_v3;
      _2342_v3 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p1, new BigNumber((_dafny.Set.fromElements(_2338_v1)).length)),p1);
      _2342_v3 = (_2342_v3).update(new BigNumber(133), (_dafny.ZERO).minus(p1));
      let _2343_v4;
      _2343_v4 = _this.f7;
      let _2344_v5;
      _2344_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2343_v4,_this.f7);
      let _2345_v6;
      _2345_v6 = _2344_v5;
      let _source31 = _2345_v6;
      let _2346___mcc_h1 = _source31;
      let _2347_cf30 = _2346___mcc_h1;
      let _rhs456 = p0;
      let _rhs457 = (_2342_v3).Merge((_2342_v3).Merge((_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(-500))).update(p1, p1)));
      _2337_v0 = _rhs456;
      _2342_v3 = _rhs457;
      _2338_v1 = (_this).f6;
      let _2348_v7;
      let _init61 = ((_2349_p1) => function (_2350_i0) {
        return (_2350_i0).minus(_2349_p1);
      })(p1);
      let _nw405 = Array((new BigNumber(16)).toNumber());
      for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw405.length); _i0_61++) {
        _nw405[_i0_61] = _init61(new BigNumber(_i0_61));
      }
      _2348_v7 = _nw405;
      let _2351_v8;
      _2351_v8 = _dafny.Seq.of((_this).f6, !(_2338_v1), (_this).f6, (_this).f6, true);
      let _index401 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_2348_v7).length));
      (_2348_v7)[_index401] = (new BigNumber((_2351_v8).length)).multipliedBy(p1);
      let _arr2 = _this.f7;
      let _index402 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_this.f7).length));
      _arr2[_index402] = false;
      let _2352_v10;
      _2352_v10 = _dafny.Seq.of(p1);
      let _index403 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_2348_v7).length));
      let _arr3 = _this.f7;
      let _index404 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_this.f7).length));
      let _rhs458 = p1;
      let _rhs459 = _module.__default.safeModuloInt(p1, p1);
      let _rhs460 = (_module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll72 = new _dafny.Map();
        for (const _compr_72 of (_2352_v10).Elements) {
          let _2353_v9 = _compr_72;
          if (_dafny.Seq.contains(_2352_v10, _2353_v9)) {
            _coll72.push([_module.__default.safeDivisionInt(_2353_v9, p1),p1]);
          }
        }
        return _coll72;
      }()).length)))).plus(p1);
      let _rhs461 = (_this).f6;
      let _lhs302 = globalState;
      let _lhs303 = _2348_v7;
      let _lhs304 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_2348_v7).length));
      let _lhs305 = globalState;
      let _lhs306 = _this.f7;
      let _lhs307 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_this.f7).length));
      _lhs302.f1 = _rhs458;
      _lhs303[_lhs304] = _rhs459;
      _lhs305.f1 = _rhs460;
      _lhs306[_lhs307] = _rhs461;
      let _2354_v11;
      _2354_v11 = _2342_v3;
      let _2355_v12;
      _2355_v12 = _dafny.Seq.of(_2351_v8, _2351_v8);
      let _2356_v13;
      _2356_v13 = _module.D12.create_DC22(_2354_v11, _2355_v12, (_this).f5, _2338_v1, _2337_v0);
      let _2357_v14;
      _2357_v14 = _dafny.Seq.of((_2356_v13).dtor_cf28, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aldrxfvuv"), p0), _2337_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(461)), function (_2358_i1) {
        return (_this).f5;
      }), _dafny.Seq.UnicodeFromString("toj"));
      _2357_v14 = _2357_v14;
      _2338_v1 = !(_2338_v1);
      (globalState).f1 = new BigNumber(((_module.D19.create_DC38(_module.__default.fm63(p1, globalState))).dtor_cf49).length);
      let _2359_v15;
      let _init62 = function (_2360_i3) {
        return (_2360_i3).minus(new BigNumber((_dafny.Seq.of((_this).f6)).length));
      };
      let _nw406 = Array((new BigNumber(3)).toNumber());
      for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw406.length); _i0_62++) {
        _nw406[_i0_62] = _init62(new BigNumber(_i0_62));
      }
      _2359_v15 = _nw406;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2359_v15).length))) {
        let _2361_i2 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2361_i2)) && ((_2361_i2).isLessThan(new BigNumber((_2359_v15).length))))) {
          (_2359_v15)[(_2361_i2)] = (_2361_i2).plus((_module.D14.create_DC26((_this).f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),false)).length), (_this).f5)).dtor_cf33);
        }
      }
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _2362_v0;
      _2362_v0 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, (_this).f6);
      let _hi13 = p1;
      for (let _2363_i0 = new BigNumber((_2362_v0).cardinality()); _2363_i0.isLessThan(_hi13); _2363_i0 = _2363_i0.plus(_dafny.ONE)) {
        let _2364_v1;
        let _init63 = ((_2365_p1) => function (_2366_i1) {
          return (_2366_i1).minus(_2365_p1);
        })(p1);
        let _nw407 = Array((new BigNumber(5)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw407.length); _i0_63++) {
          _nw407[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2364_v1 = _nw407;
        let _index405 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length));
        (_2364_v1)[_index405] = _2363_i0;
        let _2367_v2;
        _2367_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(815));
        let _2368_v3;
        _2368_v3 = _module.D0.create_DC0((_this).f6);
        let _2369_v5;
        _2369_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_2363_i0)).length), p1));
        let _2370_v6;
        _2370_v6 = _dafny.Seq.of(p1, new BigNumber((p2).length));
        let _index406 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length));
        let _rhs462 = (p1).plus((((_2367_v2).contains(p0)) ? ((_2367_v2).get(p0)) : ((_dafny.ZERO).minus(p1))));
        let _rhs463 = _module.__default.fm0((_this).fm9(p1, _2368_v3, (_this).f6, p1, globalState), globalState);
        let _rhs464 = (_2363_i0).isLessThanOrEqualTo((((_2362_v0).contains(p0)) ? ((_2362_v0).get(p0)) : (_module.__default.fm3((_this).f6, (_this).f6, globalState))));
        let _rhs465 = _module.__default.safeDivisionInt((new BigNumber((function () {
          let _coll73 = new _dafny.Map();
          for (const _compr_73 of ((((_2369_v5).contains((_this).f6)) ? ((_2369_v5).get((_this).f6)) : (_2370_v6))).Elements) {
            let _2371_v4 = _compr_73;
            if (_dafny.Seq.contains((((_2369_v5).contains((_this).f6)) ? ((_2369_v5).get((_this).f6)) : (_2370_v6)), _2371_v4)) {
              _coll73.push([(_2371_v4).multipliedBy(_2363_i0),new BigNumber((_2367_v2).length)]);
            }
          }
          return _coll73;
        }()).length)).plus(p1), (_2363_i0).multipliedBy(p1));
        let _lhs308 = globalState;
        let _lhs309 = _2364_v1;
        let _lhs310 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length));
        _lhs308.f1 = _rhs462;
        r0 = _rhs463;
        r0 = _rhs464;
        _lhs309[_lhs310] = _rhs465;
        let _2372_v7;
        _2372_v7 = _dafny.Seq.of(p0);
        let _rhs466 = (_this).f6;
        let _rhs467 = _2362_v0;
        let _rhs468 = (!((_this).f6)) === ((_2372_v7)[_module.__default.safeIndex((_2364_v1)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length))], new BigNumber((_2372_v7).length))]);
        let _rhs469 = _this.f7;
        let _lhs311 = globalState;
        let _lhs312 = _this;
        r0 = _rhs466;
        _lhs311.f4 = _rhs467;
        r0 = _rhs468;
        _lhs312.f7 = _rhs469;
        r0 = (_this).f6;
        let _index407 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length));
        (_2364_v1)[_index407] = (_2364_v1)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_2364_v1).length))];
      }
      let _2373_i2;
      _2373_i2 = _dafny.ZERO;
      L22: {
        while ((_this).f6) {
          C22: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2373_i2)) {
              break L22;
            }
            _2373_i2 = (_2373_i2).plus(_dafny.ONE);
            (globalState).f1 = p1;
            let _2374_v8;
            _2374_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
            let _2375_v9;
            _2375_v9 = _module.D11.create_DC20((_this).f6, new BigNumber((_2374_v8).length));
            let _2376_v10;
            let _nw408 = Array((new BigNumber(12)).toNumber());
            _nw408[(_dafny.ZERO).toNumber()] = (_2375_v9).dtor_cf22;
            _nw408[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(p1);
            _nw408[(new BigNumber(2)).toNumber()] = p1;
            _nw408[(new BigNumber(3)).toNumber()] = new BigNumber((_2362_v0).cardinality());
            _nw408[(new BigNumber(4)).toNumber()] = new BigNumber(544);
            _nw408[(new BigNumber(5)).toNumber()] = new BigNumber((p2).length);
            _nw408[(new BigNumber(6)).toNumber()] = p1;
            _nw408[(new BigNumber(7)).toNumber()] = p1;
            _nw408[(new BigNumber(8)).toNumber()] = new BigNumber(669);
            _nw408[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p1);
            _nw408[(new BigNumber(10)).toNumber()] = p1;
            _nw408[(new BigNumber(11)).toNumber()] = p1;
            _2376_v10 = _nw408;
            let _2377_v11;
            let _nw409 = new _module.C4();
            _nw409.__ctor();
            _2377_v11 = _nw409;
            let _2378_v12;
            _2378_v12 = _dafny.MultiSet.fromElements(_2377_v11, _2377_v11, _2377_v11, _2377_v11);
            let _index408 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2376_v10).length));
            (_2376_v10)[_index408] = (((_2378_v12).contains(_2377_v11)) ? ((_2378_v12).get(_2377_v11)) : (new BigNumber((p2).length)));
            let _2379_v13;
            _2379_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p2);
            let _index409 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2376_v10).length));
            (_2376_v10)[_index409] = (new BigNumber((_2379_v13).length)).minus(p1);
            let _2380_v14;
            let _nw410 = Array((new BigNumber(5)).toNumber()).fill([]);
            _2380_v14 = _nw410;
            let _2381_v15;
            let _nw411 = Array((new BigNumber(29)).toNumber()).fill(_module.D2.Default());
            _2381_v15 = _nw411;
            let _2382_v16;
            let _nw412 = Array((new BigNumber(17)).toNumber());
            _nw412[(_dafny.ZERO).toNumber()] = _2381_v15;
            _nw412[(_dafny.ONE).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(2)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(3)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(4)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(5)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(6)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(7)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(8)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(9)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(10)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(11)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(12)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(13)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(14)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(15)).toNumber()] = _2381_v15;
            _nw412[(new BigNumber(16)).toNumber()] = _2381_v15;
            _2382_v16 = _nw412;
            let _index410 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2380_v14).length));
            (_2380_v14)[_index410] = _2382_v16;
            let _index411 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2380_v14).length));
            (_2380_v14)[_index411] = _2382_v16;
            let _2383_v17;
            let _nw413 = new _module.C6();
            _nw413.__ctor();
            _2383_v17 = _nw413;
            _2383_v17 = _2383_v17;
          }
        }
      }
      let _2384_v18;
      _2384_v18 = _dafny.Seq.of(p2, _dafny.Seq.UnicodeFromString("s"));
      let _2385_i3;
      _2385_i3 = _dafny.ZERO;
      L23: {
        while (_dafny.Seq.IsProperPrefixOf(_2384_v18, _dafny.Seq.of(p2))) {
          C23: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2385_i3)) {
              break L23;
            }
            _2385_i3 = (_2385_i3).plus(_dafny.ONE);
            let _2386_v19;
            let _nw414 = new _module.C6();
            _nw414.__ctor();
            _2386_v19 = _nw414;
            let _2387_v20;
            _2387_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2386_v19,p0);
            let _2388_v22;
            _2388_v22 = _module.D14.create_DC26((_this).f6, p1, _module.__default.fm36(p1, p1, new BigNumber((_2387_v20).length), function () {
  let _coll74 = new _dafny.Map();
  for (const _compr_74 of _dafny.IntegerRange(new BigNumber(137), new BigNumber(471))) {
    let _2389_v21 = _compr_74;
    if (((new BigNumber(137)).isLessThanOrEqualTo(_2389_v21)) && ((_2389_v21).isLessThan(new BigNumber(471)))) {
      _coll74.push([_module.__default.safeModuloInt(_2389_v21, p1),p1]);
    }
  }
  return _coll74;
}(), globalState));
            let _source32 = _2388_v22;
            if (_source32.is_DC26) {
              let _2390___mcc_h0 = (_source32).cf32;
              let _2391___mcc_h1 = (_source32).cf33;
              let _2392___mcc_h2 = (_source32).cf34;
              let _2393_cf34 = _2392___mcc_h2;
              let _2394_cf33 = _2391___mcc_h1;
              let _2395_cf32 = _2390___mcc_h0;
              let _2396_v23;
              _2396_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
              _2396_v23 = (_2396_v23).update(((((_2362_v0).contains(false)) ? ((_2362_v0).get(false)) : (p1))).isLessThanOrEqualTo(_2394_cf33), p1);
              (globalState).f1 = new BigNumber(209);
              let _2397_v24;
              _2397_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2395_cf32,p2);
              _2394_cf33 = ((_2394_cf33).multipliedBy(p1)).multipliedBy((new BigNumber((_2397_v24).length)).multipliedBy(p1));
              let _2398_v25;
              _2398_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _2399_v26;
              _2399_v26 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("inqarh")).length), new BigNumber((_2398_v25).length));
              let _2400_v27;
              _2400_v27 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_2399_v26, _2399_v26), _dafny.Seq.of(p1), _2399_v26, _2399_v26);
              _2400_v27 = (_2400_v27).Difference(_2400_v27);
            } else {
              let _2401___mcc_h3 = (_source32).cf31;
              let _2402_cf31 = _2401___mcc_h3;
              let _2403_v28;
              let _nw415 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2403_v28 = _nw415;
              let _index412 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2403_v28).length));
              (_2403_v28)[_index412] = new _dafny.CodePoint('q'.codePointAt(0));
              let _index413 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2403_v28).length));
              (_2403_v28)[_index413] = new _dafny.CodePoint('t'.codePointAt(0));
              let _2404_v29;
              let _init64 = ((_2405_p0) => function (_2406_i4) {
                return (_2406_i4).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_2405_p0, (_this).f6)).length), new BigNumber(218))).cardinality()));
              })(p0);
              let _nw416 = Array((new BigNumber(28)).toNumber());
              for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw416.length); _i0_64++) {
                _nw416[_i0_64] = _init64(new BigNumber(_i0_64));
              }
              _2404_v29 = _nw416;
              let _2407_v30;
              _2407_v30 = _dafny.Set.fromElements((_this).f6, (_this).f6);
              let _index414 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2404_v29).length));
              (_2404_v29)[_index414] = (new BigNumber((_2407_v30).length)).minus((_dafny.ZERO).minus(p1));
              let _2408_v31;
              _2408_v31 = _dafny.Seq.of((_this).f6);
              let _2409_v32;
              _2409_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
              let _2410_v33;
              _2410_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _2411_v34;
              _2411_v34 = _dafny.MultiSet.fromElements(p1, new BigNumber((_2410_v33).length));
              let _2412_v35;
              _2412_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),(_2386_v19).fm4(_2408_v31, new BigNumber(306), new BigNumber((_2409_v32).length), _2411_v34, globalState));
              let _index415 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2404_v29).length));
              (_2404_v29)[_index415] = ((((_2412_v35).contains(p1)) ? ((_2412_v35).get(p1)) : (p1))).minus(_module.__default.safeDivisionInt(new BigNumber(30), p1));
              let _2413_v36;
              let _nw417 = new _module.C2();
              _nw417.__ctor(((p0) ? (new _dafny.CodePoint('l'.codePointAt(0))) : (new _dafny.CodePoint('c'.codePointAt(0)))), (_this).f6);
              _2413_v36 = _nw417;
              let _2414_v37;
              _2414_v37 = _module.D20.create_DC42(_2386_v19);
              let _2415_v38;
              _2415_v38 = _dafny.Map.Empty.slice().updateUnsafe(((_2413_v36).fm8(true, (_this).f6, p0, globalState)).plus(p1),(_2414_v37).dtor_cf55);
              _2415_v38 = (_2415_v38).update((_2386_v19).fm4(_2408_v31, p1, p1, _2411_v34, globalState), _2386_v19);
            }
            let _2416_v39;
            _2416_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(!(p0), p0, (_this).f6, (_this).f6, (_this).f6)).cardinality()),p1);
            let _2417_v40;
            _2417_v40 = _dafny.Set.fromElements(p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p0)).length));
            let _2418_v41;
            _2418_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),_dafny.Seq.of(p1));
            let _2419_v42;
            _2419_v42 = _module.D19.create_DC40(p0, new BigNumber((_2417_v40).length), _2418_v41);
            _2416_v39 = (_2416_v39).update(p1, (_2419_v42).dtor_cf52);
            let _2420_v43;
            let _nw418 = Array((new BigNumber(13)).toNumber()).fill([]);
            _2420_v43 = _nw418;
            let _2421_v44;
            let _nw419 = Array((_dafny.ONE).toNumber());
            _nw419[(_dafny.ZERO).toNumber()] = p1;
            _2421_v44 = _nw419;
            let _index416 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_2420_v43).length));
            (_2420_v43)[_index416] = _2421_v44;
            let _index417 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_2421_v44).length));
            (_2421_v44)[_index417] = (new BigNumber(797)).multipliedBy(p1);
            let _2422_v45;
            _2422_v45 = _dafny.Seq.of(p0, (_this).f6, p0);
            let _index418 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_2420_v43).length));
            let _index419 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_2421_v44).length));
            let _rhs470 = _2421_v44;
            let _rhs471 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_2423_i5) {
              return (_this).f5;
            }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vhbetwjr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), function (_2424_i6) {
              return (_this).f5;
            })));
            let _rhs472 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1);
            let _rhs473 = _2422_v45;
            let _lhs313 = _2420_v43;
            let _lhs314 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_2420_v43).length));
            let _lhs315 = _2421_v44;
            let _lhs316 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_2421_v44).length));
            _lhs313[_lhs314] = _rhs470;
            r0 = _rhs471;
            _lhs315[_lhs316] = _rhs472;
            _2422_v45 = _rhs473;
            let _arr4 = _this.f7;
            let _index420 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_this.f7).length));
            _arr4[_index420] = ((p0) ? (p0) : ((_2422_v45)[_module.__default.safeIndex((_2421_v44)[_module.__default.safeIndex(new BigNumber(16), new BigNumber((_2421_v44).length))], new BigNumber((_2422_v45).length))]));
            let _arr5 = _this.f7;
            let _index421 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_this.f7).length));
            _arr5[_index421] = p0;
          }
        }
      }
      let _2425_v46;
      let _2426_v47;
      let _2427_v48;
      let _out50;
      let _out51;
      let _out52;
      let _outcollector22 = (_this).m6(((_this).f6) === ((_this).f6), (p1).isLessThan(p1), globalState);
      _out50 = _outcollector22[0];
      _out51 = _outcollector22[1];
      _out52 = _outcollector22[2];
      _2425_v46 = _out50;
      _2426_v47 = _out51;
      _2427_v48 = _out52;
      (globalState).f1 = p1;
      r0 = _2426_v47;
      r0 = _2427_v48;
      r1 = p2;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2428_v0;
      _2428_v0 = new BigNumber(210);
      let _2429_v1;
      _2429_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2428_v0,(_this).f6);
      let _2430_v2;
      _2430_v2 = _dafny.Seq.of(_module.__default.fm33(globalState), _2429_v1);
      let _2431_v3;
      _2431_v3 = (_2430_v2)[_module.__default.safeIndex(_2428_v0, new BigNumber((_2430_v2).length))];
      let _source33 = _2431_v3;
      let _2432___mcc_h0 = _source33;
      let _2433_cf12 = _2432___mcc_h0;
      let _2434_v4;
      let _nw420 = Array((new BigNumber(9)).toNumber());
      _nw420[(_dafny.ZERO).toNumber()] = _2428_v0;
      _nw420[(_dafny.ONE).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(2)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(3)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(4)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(5)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(6)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(7)).toNumber()] = _2428_v0;
      _nw420[(new BigNumber(8)).toNumber()] = _2428_v0;
      _2434_v4 = _nw420;
      let _2435_v5;
      _2435_v5 = _2434_v4;
      let _2436_v6;
      _2436_v6 = _module.D12.create_DC23(_2435_v5);
      let _pat_let_tv32 = _2435_v5;
      let _pat_let_tv33 = _2434_v4;
      let _source34 = function (_pat_let42_0) {
        return function (_2439_dt__update__tmp_h1) {
          return function (_pat_let45_0) {
            return function (_2440_dt__update_hcf29_h1) {
              return _module.D12.create_DC23(_2440_dt__update_hcf29_h1);
            }(_pat_let45_0);
          }(_pat_let_tv33);
        }(_pat_let42_0);
      }((((_this).f6) ? (function (_pat_let43_0) {
        return function (_2437_dt__update__tmp_h0) {
          return function (_pat_let44_0) {
            return function (_2438_dt__update_hcf29_h0) {
              return _module.D12.create_DC23(_2438_dt__update_hcf29_h0);
            }(_pat_let44_0);
          }(_pat_let_tv32);
        }(_pat_let43_0);
      }(_2436_v6)) : (_2436_v6)));
      if (_source34.is_DC22) {
        let _2441___mcc_h1 = (_source34).cf24;
        let _2442___mcc_h2 = (_source34).cf25;
        let _2443___mcc_h3 = (_source34).cf26;
        let _2444___mcc_h4 = (_source34).cf27;
        let _2445___mcc_h5 = (_source34).cf28;
        let _2446_cf28 = _2445___mcc_h5;
        let _2447_cf27 = _2444___mcc_h4;
        let _2448_cf26 = _2443___mcc_h3;
        let _2449_cf25 = _2442___mcc_h2;
        let _2450_cf24 = _2441___mcc_h1;
        let _2451_v7;
        _2451_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2428_v0),_2428_v0);
        let _2452_v8;
        _2452_v8 = _dafny.Seq.of(new BigNumber((_2451_v7).length));
        let _2453_v9;
        let _nw421 = new _module.C3();
        _nw421.__ctor(true, _2447_cf27);
        _2453_v9 = _nw421;
        let _2454_v10;
        _2454_v10 = _dafny.Set.fromElements(_2453_v9, _2453_v9, _2453_v9);
        let _2455_v11;
        _2455_v11 = _dafny.MultiSet.fromElements(_2452_v8, _2452_v8);
        let _2456_v12;
        _2456_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2455_v11,_2428_v0);
        let _rhs474 = (((_2454_v10).IsSubsetOf(_2454_v10)) ? (_dafny.Seq.update(_dafny.Seq.Concat(_2452_v8, _dafny.Seq.update(_2452_v8, _module.__default.safeIndex(_2428_v0, new BigNumber((_2452_v8).length)), _2428_v0)), _module.__default.safeIndex((((_2456_v12).contains(_2455_v11)) ? ((_2456_v12).get(_2455_v11)) : (_2428_v0)), new BigNumber((_dafny.Seq.Concat(_2452_v8, _dafny.Seq.update(_2452_v8, _module.__default.safeIndex(_2428_v0, new BigNumber((_2452_v8).length)), _2428_v0))).length)), new BigNumber(672))) : (_2452_v8));
        let _rhs475 = _2447_cf27;
        _2452_v8 = _rhs474;
        _2447_cf27 = _rhs475;
        let _2457_v13;
        let _nw422 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2457_v13 = _nw422;
        let _2458_v14;
        _2458_v14 = _module.D18.create_DC36(_2457_v13);
        let _2459_v15;
        _2459_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2458_v14,(_2428_v0).minus(new BigNumber(-10)));
        _2459_v15 = (_2459_v15).update(_2458_v14, _module.__default.safeDivisionInt(_2428_v0, _2428_v0));
        _2428_v0 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2434_v4,_2428_v0)).length)).plus(_2428_v0);
        let _index422 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2434_v4).length));
        (_2434_v4)[_index422] = _2428_v0;
        let _index423 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2434_v4).length));
        (_2434_v4)[_index423] = _module.__default.safeDivisionInt(_2428_v0, _module.__default.safeDivisionInt(new BigNumber((_2446_cf28).length), new BigNumber((_2452_v8).length)));
      } else if (_source34.is_DC23) {
        let _2460___mcc_h6 = (_source34).cf29;
        let _2461_cf29 = _2460___mcc_h6;
        let _2462_v16;
        _2462_v16 = new _dafny.CodePoint('e'.codePointAt(0));
        _2462_v16 = (_this).f5;
        let _2463_v17;
        _2463_v17 = _dafny.Set.fromElements((_this).f6);
        _2463_v17 = _2463_v17;
        let _2464_v18;
        _2464_v18 = _dafny.Seq.of(_2428_v0, _2428_v0);
        let _2465_v19;
        _2465_v19 = _dafny.MultiSet.fromElements((_2464_v18)[_module.__default.safeIndex(_2428_v0, new BigNumber((_2464_v18).length))]);
        _2465_v19 = _module.__default.fm32(globalState);
        let _2466_v20;
        _2466_v20 = true;
        let _2467_v21;
        _2467_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2428_v0,_2462_v16);
        let _2468_v22;
        _2468_v22 = _dafny.Seq.UnicodeFromString("ryhpyo");
        let _2469_v23;
        _2469_v23 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),(((_2467_v21).contains(new BigNumber((_2468_v22).length))) ? ((_2467_v21).get(new BigNumber((_2468_v22).length))) : ((_this).f5)));
        let _2470_v24;
        _2470_v24 = _dafny.Set.fromElements(_2469_v23);
        let _2471_v25;
        _2471_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2469_v23,(_this).f6);
        let _2472_v27;
        _2472_v27 = _dafny.MultiSet.fromElements(!((_this).f6), (_this).f6);
        let _rhs476 = (_this).f6;
        let _rhs477 = (function () {
          let _coll75 = new _dafny.Set();
          for (const _compr_75 of (_2471_v25).Keys.Elements) {
            let _2473_v26 = _compr_75;
            if ((_2471_v25).contains(_2473_v26)) {
              _coll75.add(_2473_v26);
            }
          }
          return _coll75;
        }()).Union(_2470_v24);
        let _rhs478 = (_2472_v27).Difference(_dafny.MultiSet.fromElements(!(_module.__default.fm0(_2466_v20, globalState))));
        let _lhs317 = globalState;
        _2466_v20 = _rhs476;
        _2470_v24 = _rhs477;
        _lhs317.f4 = _rhs478;
      } else {
        let _2474___mcc_h7 = (_source34).cf23;
        let _2475_cf23 = _2474___mcc_h7;
        (globalState).f1 = _2428_v0;
        let _2476_v28;
        let _nw423 = new _module.C5();
        _nw423.__ctor(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6))).length), (_this).f5, (_this).f6);
        _2476_v28 = _nw423;
        let _2477_v29;
        _2477_v29 = _module.D6.create_DC14((_this).f6);
        let _2478_v30;
        let _init65 = function (_2479_i0) {
          return _dafny.MultiSet.fromElements((_this).f6);
        };
        let _nw424 = Array((new BigNumber(11)).toNumber());
        for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw424.length); _i0_65++) {
          _nw424[_i0_65] = _init65(new BigNumber(_i0_65));
        }
        _2478_v30 = _nw424;
        let _2480_v31;
        _2480_v31 = _module.D18.create_DC36(_2478_v30);
        let _2481_v32;
        _2481_v32 = _dafny.MultiSet.fromElements(_2480_v31, _2480_v31, _module.D18.create_DC36(_2478_v30), _2480_v31, (((_this).f6) ? (_2480_v31) : (_2480_v31)));
        let _rhs479 = _module.__default.fm64((_this).f5, globalState);
        let _rhs480 = (_dafny.MultiSet.fromElements(_2480_v31)).Union(_2481_v32);
        _2477_v29 = _rhs479;
        _2481_v32 = _rhs480;
        let _2482_v33;
        _2482_v33 = _dafny.Set.fromElements(_2428_v0, _2476_v28.f15);
        let _2483_v34;
        _2483_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2428_v0,_2482_v33);
        let _2484_v35;
        _2484_v35 = _dafny.Seq.UnicodeFromString("wae");
        _2483_v34 = (_2483_v34).update(_2428_v0, (_2482_v33).Union(_dafny.Set.fromElements(new BigNumber((_2484_v35).length), _2476_v28.f15)));
      }
      let _2485_v36;
      _2485_v36 = true;
      let _2486_v37;
      _2486_v37 = _dafny.Seq.of(_2428_v0);
      let _2487_v38;
      _2487_v38 = _dafny.MultiSet.fromElements(true);
      _2485_v36 = (_this).fm6(_dafny.Seq.Concat(_2486_v37, _dafny.Seq.of(new BigNumber((_2487_v38).cardinality()), _2428_v0)), globalState);
      let _arr6 = _this.f7;
      let _index424 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_this.f7).length));
      _arr6[_index424] = _2485_v36;
      let _arr7 = _this.f7;
      let _index425 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_this.f7).length));
      _arr7[_index425] = (_this).f6;
      let _arr8 = _this.f7;
      let _index426 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_this.f7).length));
      _arr8[_index426] = (_this.f7)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_this.f7).length))];
      let _2488_v39;
      _2488_v39 = _dafny.MultiSet.fromElements((_this).f6);
      let _2489_v40;
      _2489_v40 = _dafny.Seq.of((_this).f6, (_this).f6, (_this).f6);
      let _2490_v41;
      let _nw425 = new _module.C6();
      _nw425.__ctor();
      _2490_v41 = _nw425;
      let _2491_v42;
      _2491_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2490_v41,(_this).f6);
      let _2492_v43;
      _2492_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2491_v42).length),_2428_v0);
      let _2493_v44;
      _2493_v44 = _dafny.Set.fromElements(_2488_v39, _dafny.MultiSet.FromArray(_2489_v40), _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of((_this).f6), _module.__default.safeIndex(_2428_v0, new BigNumber((_dafny.Seq.of((_this).f6)).length)), (_this).f6), _module.__default.safeIndex((((_2492_v43).contains(new BigNumber(-249))) ? ((_2492_v43).get(new BigNumber(-249))) : (_2428_v0)), new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f6), _module.__default.safeIndex(_2428_v0, new BigNumber((_dafny.Seq.of((_this).f6)).length)), (_this).f6)).length)), (_this).f6)));
      let _2494_v45;
      let _nw426 = new _module.C7();
      _nw426.__ctor(new BigNumber((_2493_v44).length), _module.__default.safeModuloInt(_2428_v0, _2428_v0), (_this).f5, (_this).f6);
      _2494_v45 = _nw426;
      if ((new BigNumber(800)).isLessThan(new BigNumber(946))) {
        let _2495_v46;
        _2495_v46 = true;
        _2495_v46 = ((_this).f6) && ((_this).f6);
        let _arr9 = _this.f7;
        let _index427 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length));
        _arr9[_index427] = (_this).f6;
        let _2496_v47;
        _2496_v47 = _dafny.Seq.UnicodeFromString("mkmwkypd");
        let _2497_v48;
        _2497_v48 = _dafny.Seq.of((_2494_v45).f9);
        let _2498_v49;
        _2498_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2428_v0,_2497_v48);
        let _2499_v50;
        _2499_v50 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_2428_v0, _2494_v45.f10)).length), (_2494_v45).f9);
        let _arr10 = _this.f7;
        let _index428 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length));
        let _rhs481 = ((true) ? (_2495_v46) : ((_2490_v41).fm6(_2497_v48, globalState)));
        let _rhs482 = new BigNumber(111);
        let _rhs483 = ((((!((_module.D19.create_DC40((_this).f6, (_2494_v45).f9, _2498_v49)).dtor_cf51)) ? ((_this).f6) : ((_this).f6))) ? ((_this).f6) : (!((_this).f6)));
        let _rhs484 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2496_v47, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_2500_i1) {
          return (_this).f5;
        }), _module.__default.safeIndex((_2494_v45).fm4(_2489_v40, _2428_v0, _2494_v45.f10, _2499_v50, globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_2501_i1) {
          return (_this).f5;
        })).length)), (_this).f5)), _2496_v47);
        let _lhs318 = _this.f7;
        let _lhs319 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length));
        _2495_v46 = _rhs481;
        _2428_v0 = _rhs482;
        _lhs318[_lhs319] = _rhs483;
        _2496_v47 = _rhs484;
        let _2502_v51;
        _2502_v51 = _dafny.Set.fromElements(_2428_v0, (_2494_v45).f9);
        _2495_v46 = (_2502_v51).equals(_2502_v51);
        let _2503_v52;
        _2503_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_2495_v46, globalState),false);
        let _2504_v53;
        _2504_v53 = _module.D21.create_DC45(_2499_v50);
        let _2505_v54;
        _2505_v54 = _dafny.Map.Empty.slice().updateUnsafe((_2488_v39).update((_this.f7)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length))], _module.__default.abs(new BigNumber((_2503_v52).length))),((_2504_v53).dtor_cf60).IsDisjointFrom(_2499_v50));
        _2505_v54 = (_2505_v54).update(_2488_v39, !(!((_this.f7)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length))])));
        let _2506_v55;
        let _2507_v56;
        let _2508_v57;
        let _out53;
        let _out54;
        let _out55;
        let _outcollector23 = (_this).m6((_this.f7)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length))], (_this.f7)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_this.f7).length))], globalState);
        _out53 = _outcollector23[0];
        _out54 = _outcollector23[1];
        _out55 = _outcollector23[2];
        _2506_v55 = _out53;
        _2507_v56 = _out54;
        _2508_v57 = _out55;
      } else {
        if (true) {
          let _2509_v58;
          _2509_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_2494_v45.f10, !((_this).f6), _2494_v45.f10, globalState),_2428_v0);
          let _2510_v59;
          _2510_v59 = _dafny.Seq.UnicodeFromString("deayq");
          let _2511_v60;
          _2511_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2494_v45.f10);
          _2509_v58 = (_2509_v58).update(_dafny.Seq.Concat(_2510_v59, _2510_v59), (((_2511_v60).contains((_this).f6)) ? ((_2511_v60).get((_this).f6)) : (_2428_v0)));
          let _2512_v61;
          let _nw427 = new _module.C1();
          _nw427.__ctor();
          _2512_v61 = _nw427;
          let _2513_v62;
          _2513_v62 = true;
          _2513_v62 = !((_this).f6) || (!((_this).f6));
          let _2514_v63;
          _2514_v63 = _dafny.Seq.of(new BigNumber(638));
          let _2515_v64;
          _2515_v64 = _dafny.Seq.of(_2514_v63);
          let _rhs485 = !_dafny.Seq.contains(_2515_v64, (_2515_v64)[_module.__default.safeIndex(new BigNumber((_2514_v63).length), new BigNumber((_2515_v64).length))]);
          let _rhs486 = _2428_v0;
          let _rhs487 = (_2494_v45).f9;
          let _lhs320 = _2494_v45;
          let _lhs321 = globalState;
          _2513_v62 = _rhs485;
          _lhs320.f10 = _rhs486;
          _lhs321.f1 = _rhs487;
          let _2516_v65;
          _2516_v65 = _dafny.Set.fromElements((_this).f6);
          let _2517_v66;
          _2517_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm6(_dafny.Seq.update(_2514_v63, _module.__default.safeIndex(_2494_v45.f10, new BigNumber((_2514_v63).length)), (_2494_v45).f9), globalState),_2516_v65);
          let _2518_v67;
          _2518_v67 = _dafny.Seq.of(_2516_v65, (((_2517_v66).contains(true)) ? ((_2517_v66).get(true)) : (_2516_v65)));
          let _2519_v68;
          _2519_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f7);
          let _2520_v69;
          _2520_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2518_v67,new BigNumber(((_2519_v68).Merge(_2519_v68)).length));
          _2520_v69 = (_2520_v69).update(_2518_v67, _2428_v0);
        } else {
          let _arr11 = _this.f7;
          let _index429 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f7).length));
          _arr11[_index429] = !(((_2490_v41).fm7((_this).f6, globalState)).isLessThanOrEqualTo((_2494_v45).f9));
          let _arr12 = _this.f7;
          let _index430 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f7).length));
          _arr12[_index430] = _module.__default.fm0((_this).f6, globalState);
          let _2521_v70;
          _2521_v70 = _dafny.Seq.UnicodeFromString("baymc");
          let _2522_v71;
          let _nw428 = Array((new BigNumber(7)).toNumber()).fill(false);
          _2522_v71 = _nw428;
          let _2523_v72;
          _2523_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2494_v45.f10,_2522_v71);
          let _2524_v73;
          _2524_v73 = _dafny.Seq.of(_2494_v45.f10, _2494_v45.f10, (_dafny.ZERO).minus(new BigNumber((_2521_v70).length)));
          let _2525_v74;
          _2525_v74 = _dafny.MultiSet.fromElements((_2494_v45).f9);
          let _2526_v75;
          _2526_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2494_v45.f10,_dafny.Seq.of((((_2525_v74).contains(_2494_v45.f10)) ? ((_2525_v74).get(_2494_v45.f10)) : ((_2494_v45).f9))));
          let _2527_v76;
          _2527_v76 = _dafny.MultiSet.fromElements(new BigNumber((_2526_v75).length), new BigNumber(-162));
          let _2528_v77;
          let _nw429 = Array((new BigNumber(24)).toNumber());
          _nw429[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((((_this).f6) ? ((_2494_v45).f9) : (_2494_v45.f10)));
          _nw429[(_dafny.ONE).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(2)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(3)).toNumber()] = (new BigNumber((_dafny.Seq.update(_2521_v70, _module.__default.safeIndex((_2494_v45).f9, new BigNumber((_2521_v70).length)), (_this).f5)).length)).multipliedBy((_2494_v45).f9);
          _nw429[(new BigNumber(4)).toNumber()] = new BigNumber((((_2523_v72).update(_2494_v45.f10, _2522_v71)).Merge(_2523_v72)).length);
          _nw429[(new BigNumber(5)).toNumber()] = ((_2494_v45).f9).minus(new BigNumber(-572));
          _nw429[(new BigNumber(6)).toNumber()] = new BigNumber((_2524_v73).length);
          _nw429[(new BigNumber(7)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(8)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(9)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(10)).toNumber()] = _2494_v45.f10;
          _nw429[(new BigNumber(11)).toNumber()] = _2494_v45.f10;
          _nw429[(new BigNumber(12)).toNumber()] = (_2494_v45).fm4(_2489_v40, (((_2488_v39).contains((_this).f6)) ? ((_2488_v39).get((_this).f6)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), ((_2529_v40) => function (_2530_i2) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), ((_2531_v40) => function (_2532_i3) {
              return _2531_v40;
            })(_2529_v40))).length);
          })(_2489_v40))).length))), _2494_v45.f10, _2527_v76, globalState);
          _nw429[(new BigNumber(13)).toNumber()] = (_2494_v45).fm7((_this.f7)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f7).length))], globalState);
          _nw429[(new BigNumber(14)).toNumber()] = new BigNumber((_2525_v74).cardinality());
          _nw429[(new BigNumber(15)).toNumber()] = new BigNumber(-739);
          _nw429[(new BigNumber(16)).toNumber()] = _module.__default.fm3((_this).f6, (_this).f6, globalState);
          _nw429[(new BigNumber(17)).toNumber()] = _2494_v45.f10;
          _nw429[(new BigNumber(18)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(19)).toNumber()] = new BigNumber(221);
          _nw429[(new BigNumber(20)).toNumber()] = ((_2494_v45).f9).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this.f7)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_this.f7).length))],_2494_v45.f10)).length));
          _nw429[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(738)), function (_2533_i4) {
            return (_this).f5;
          })).length);
          _nw429[(new BigNumber(22)).toNumber()] = (_2494_v45).f9;
          _nw429[(new BigNumber(23)).toNumber()] = new BigNumber((_2521_v70).length);
          _2528_v77 = _nw429;
          let _2534_v78;
          _2534_v78 = _dafny.Set.fromElements(_2494_v45.f10, _2494_v45.f10);
          let _2535_v79;
          _2535_v79 = _dafny.Seq.of(_2534_v78, _2534_v78);
          let _index431 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2528_v77).length));
          (_2528_v77)[_index431] = new BigNumber((_2535_v79).length);
          let _index432 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2528_v77).length));
          (_2528_v77)[_index432] = _module.__default.safeModuloInt((_2494_v45).f9, _2494_v45.f10);
          _2428_v0 = new BigNumber(135);
          let _2536_v80;
          _2536_v80 = _2528_v77;
          let _2537_v81;
          _2537_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2489_v40).length),_2536_v80);
          (_2494_v45).f10 = _module.__default.safeModuloInt(((_2528_v77)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_2528_v77).length))]).minus((_2528_v77)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_2528_v77).length))]), new BigNumber((_2537_v81).length));
          let _index433 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2528_v77).length));
          (_2528_v77)[_index433] = _2494_v45.f10;
        }
        let _2538_v82;
        _2538_v82 = _dafny.MultiSet.fromElements((_2494_v45).f9);
        let _2539_v83;
        _2539_v83 = _dafny.Map.Empty.slice().updateUnsafe(true,_2538_v82);
        (globalState).f1 = new BigNumber(((_2539_v83).Merge((_2539_v83).Merge(_2539_v83))).length);
        let _arr13 = _this.f7;
        let _index434 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_this.f7).length));
        _arr13[_index434] = ((_this).f6) && ((_2489_v40)[_module.__default.safeIndex(_2494_v45.f10, new BigNumber((_2489_v40).length))]);
        let _arr14 = _this.f7;
        let _index435 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_this.f7).length));
        _arr14[_index435] = false;
        let _2540_v84;
        _2540_v84 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2494_v45).f9);
        let _2541_v85;
        _2541_v85 = _module.D6.create_DC13(_dafny.Seq.of((_2494_v45).f9));
        let _2542_v86;
        _2542_v86 = _dafny.Seq.UnicodeFromString("uuafj");
        let _2543_v87;
        _2543_v87 = _dafny.Map.Empty.slice().updateUnsafe((_2541_v85).dtor_cf14,new BigNumber((_2542_v86).length));
        let _rhs488 = _module.__default.fm28((_this.f7)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_this.f7).length))], (_this.f7)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_this.f7).length))], _2543_v87, globalState);
        let _rhs489 = (_2494_v45).f9;
        let _lhs322 = globalState;
        _2540_v84 = _rhs488;
        _lhs322.f1 = _rhs489;
        let _2544_v88;
        let _init66 = ((_2545_v45) => function (_2546_i5) {
          return (_2546_i5).multipliedBy(_2545_v45.f10);
        })(_2494_v45);
        let _nw430 = Array((new BigNumber(5)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw430.length); _i0_66++) {
          _nw430[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2544_v88 = _nw430;
        let _init67 = ((_2547_v45) => function (_2548_i6) {
          return (_2548_i6).multipliedBy(_2547_v45.f10);
        })(_2494_v45);
        let _nw431 = Array((_dafny.ONE).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw431.length); _i0_67++) {
          _nw431[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2544_v88 = _nw431;
      }
      let _2549_v89;
      _2549_v89 = true;
      _2549_v89 = _2549_v89;
      _2549_v89 = _2549_v89;
      let _2550_v90;
      _2550_v90 = _dafny.Map.Empty.slice().updateUnsafe(_2549_v89,!(true));
      let _2551_v91;
      let _nw432 = new _module.C5();
      _nw432.__ctor(new BigNumber((_2488_v39).cardinality()), (_this).f5, !(!((_2492_v43).update((_2494_v45).f9, new BigNumber((_module.__default.fm1(new BigNumber((_2550_v90).length), (_this).f6, _2428_v0, globalState)).length))).contains(_2494_v45.f10)));
      _2551_v91 = _nw432;
      r0 = (_2428_v0).multipliedBy(_2551_v91.f15);
      return r0;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _2552_v0;
      _2552_v0 = new _dafny.CodePoint('l'.codePointAt(0));
      _2552_v0 = (_this).f5;
      let _2553_v1;
      _2553_v1 = _dafny.Seq.of(p0);
      let _2554_v2;
      _2554_v2 = _module.D15.create_DC28((_this).f6, (_2553_v1)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2553_v1).length))), new BigNumber((_2553_v1).length))]);
      let _source35 = _2554_v2;
      if (_source35.is_DC28) {
        let _2555___mcc_h0 = (_source35).cf36;
        let _2556___mcc_h1 = (_source35).cf37;
        let _2557_cf37 = _2556___mcc_h1;
        let _2558_cf36 = _2555___mcc_h0;
        let _2559_v3;
        _2559_v3 = new BigNumber(-712);
        let _2560_v4;
        _2560_v4 = _dafny.Seq.of(_2553_v1);
        let _rhs490 = (_2559_v3).isLessThanOrEqualTo((_2559_v3).plus(_2559_v3));
        let _rhs491 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2553_v1, _2553_v1), (_2560_v4)[_module.__default.safeIndex(_2559_v3, new BigNumber((_2560_v4).length))]);
        let _rhs492 = (_2558_cf36) && (_2557_cf37);
        _2558_cf36 = _rhs490;
        _2553_v1 = _rhs491;
        _2557_cf37 = _rhs492;
        _2558_cf36 = _2557_cf37;
        let _arr15 = _this.f7;
        let _index436 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_this.f7).length));
        _arr15[_index436] = p0;
        let _2561_v5;
        _2561_v5 = _module.D0.create_DC0(p2);
        let _2562_v6;
        _2562_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2559_v3,_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_this).fm9(_2559_v3, _2561_v5, _2558_cf36, _2559_v3, globalState))).length)));
        let _2563_v7;
        _2563_v7 = _dafny.MultiSet.fromElements(_2559_v3);
        let _2564_v8;
        _2564_v8 = _dafny.Seq.UnicodeFromString("rgibkeo");
        let _2565_v9;
        _2565_v9 = _module.D6.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), ((_2566_v3) => function (_2567_i0) {
  return _2566_v3;
})(_2559_v3)));
        let _2568_v10;
        _2568_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0);
        let _2569_v11;
        _2569_v11 = _module.D20.create_DC43(new BigNumber(((_2565_v9).dtor_cf14).length), (((_2568_v10).contains(p0)) ? ((_2568_v10).get(p0)) : ((_this).f6)), new BigNumber(88));
        let _2570_v12;
        _2570_v12 = _module.D20.create_DC44(_2569_v11);
        let _2571_v13;
        _2571_v13 = _dafny.Set.fromElements(_2570_v12);
        let _arr16 = _this.f7;
        let _index437 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_this.f7).length));
        _arr16[_index437] = ((((_2562_v6).contains((_dafny.ZERO).minus(_2559_v3))) ? ((_2562_v6).get((_dafny.ZERO).minus(_2559_v3))) : ((_dafny.MultiSet.fromElements(new BigNumber((_2564_v8).length))).update(new BigNumber(689), _module.__default.abs(new BigNumber((_2571_v13).length)))))).IsSubsetOf((((_2562_v6).contains(_2559_v3)) ? ((_2562_v6).get(_2559_v3)) : (_2563_v7)));
        (globalState).f1 = _2559_v3;
      } else if (_source35.is_DC27) {
        let _2572___mcc_h2 = (_source35).cf35;
        let _2573_cf35 = _2572___mcc_h2;
        if (p0) {
          let _2574_v14;
          _2574_v14 = _dafny.Seq.UnicodeFromString("giybh");
          let _2575_v15;
          _2575_v15 = new BigNumber(752);
          _2552_v0 = (function (_pat_let46_0) {
            return function (_2576_dt__update__tmp_h0) {
              return function (_pat_let47_0) {
                return function (_2577_dt__update_hcf43_h0) {
                  return _module.D17.create_DC34(_2577_dt__update_hcf43_h0);
                }(_pat_let47_0);
              }((_this).f5);
            }(_pat_let46_0);
          }(_module.__default.fm65((_this).f6, new BigNumber((_2574_v14).length), _2575_v15, globalState))).dtor_cf43;
          let _2578_v16;
          let _nw433 = new _module.C5();
          _nw433.__ctor(_2575_v15, (_this).f5, (new BigNumber(-195)).isLessThan((_dafny.ZERO).minus(_2575_v15)));
          _2578_v16 = _nw433;
          let _arr17 = _this.f7;
          let _index438 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_this.f7).length));
          _arr17[_index438] = !(p0) || (p0);
          let _arr18 = _this.f7;
          let _index439 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_this.f7).length));
          _arr18[_index439] = !(false) || (_dafny.Seq.IsProperPrefixOf(_2553_v1, _2553_v1));
          let _2579_v17;
          _2579_v17 = _dafny.Seq.of(_2578_v16.f15, new BigNumber((_2574_v14).length));
          let _2580_v18;
          _2580_v18 = _dafny.Seq.of(_2579_v17);
          let _2581_v19;
          let _nw434 = new _module.C0();
          _nw434.__ctor(p0, _2580_v18);
          _2581_v19 = _nw434;
          let _2582_v20;
          let _nw435 = new _module.C1();
          _nw435.__ctor();
          _2582_v20 = _nw435;
          let _2583_v21;
          _2583_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2574_v14).length),(_2581_v19).f11);
          let _2584_v22;
          _2584_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2553_v1,_2583_v21);
          let _2585_v23;
          _2585_v23 = _dafny.Map.Empty.slice().updateUnsafe((_2584_v22).update(_2553_v1, _2583_v21),_2581_v19);
          let _arr19 = _this.f7;
          let _index440 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_this.f7).length));
          let _rhs493 = (((_2585_v23).contains((_2584_v22).Merge(_2584_v22))) ? ((_2585_v23).get((_2584_v22).Merge(_2584_v22))) : (_2581_v19));
          let _rhs494 = (_2553_v1)[_module.__default.safeIndex(_2578_v16.f15, new BigNumber((_2553_v1).length))];
          let _rhs495 = _2582_v20;
          let _lhs323 = _this.f7;
          let _lhs324 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_this.f7).length));
          _2581_v19 = _rhs493;
          _lhs323[_lhs324] = _rhs494;
          _2582_v20 = _rhs495;
          let _2586_v24;
          let _nw436 = Array((new BigNumber(2)).toNumber()).fill([]);
          _2586_v24 = _nw436;
          let _2587_v25;
          let _nw437 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _2587_v25 = _nw437;
          let _index441 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_2586_v24).length));
          (_2586_v24)[_index441] = _2587_v25;
          let _2588_v26;
          _2588_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this.f7)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_this.f7).length))],new BigNumber(-212));
          let _2589_v27;
          _2589_v27 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_2588_v26);
          let _index442 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_2586_v24).length));
          let _rhs496 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_2578_v16.f15), _dafny.Seq.Concat(_dafny.Seq.of(_2575_v15), _2579_v17)), _module.__default.safeIndex(new BigNumber((_2589_v27).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_2578_v16.f15), _dafny.Seq.Concat(_dafny.Seq.of(_2575_v15), _2579_v17))).length)), (_2578_v16.f15).minus(new BigNumber((_2579_v17).length)));
          let _rhs497 = _2587_v25;
          let _lhs325 = _2586_v24;
          let _lhs326 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_2586_v24).length));
          _2579_v17 = _rhs496;
          _lhs325[_lhs326] = _rhs497;
        } else {
          let _2590_v28;
          _2590_v28 = _module.D21.create_DC47((_this).f6);
          let _2591_v29;
          _2591_v29 = false;
          let _2592_v30;
          let _nw438 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _2592_v30 = _nw438;
          let _2593_v31;
          _2593_v31 = new BigNumber(893);
          let _index443 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length));
          (_2592_v30)[_index443] = _module.__default.safeModuloInt(_2593_v31, _2593_v31);
          let _2594_v32;
          _2594_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm7(p0, globalState)),(_this).f6);
          let _2595_v33;
          _2595_v33 = _dafny.Seq.of(_this.f7);
          let _2596_v34;
          _2596_v34 = _dafny.Seq.UnicodeFromString("eu");
          let _pat_let_tv34 = _2593_v31;
          let _pat_let_tv35 = _2593_v31;
          let _index444 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length));
          let _rhs498 = ((!((((_2594_v32).contains(_2593_v31)) ? ((_2594_v32).get(_2593_v31)) : (_2591_v29)))) ? (_2593_v31) : (new BigNumber((p1).cardinality())));
          let _rhs499 = function (_pat_let48_0) {
            return function (_2597_dt__update__tmp_h1) {
              return function (_pat_let49_0) {
                return function (_2598_dt__update_hcf64_h0) {
                  return _module.D21.create_DC47(_2598_dt__update_hcf64_h0);
                }(_pat_let49_0);
              }((_pat_let_tv34).isLessThanOrEqualTo((_dafny.ZERO).minus(_pat_let_tv35)));
            }(_pat_let48_0);
          }(_2590_v28);
          let _rhs500 = _2591_v29;
          let _rhs501 = (_2595_v33)[_module.__default.safeIndex(_2593_v31, new BigNumber((_2595_v33).length))];
          let _rhs502 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_2596_v34, _2596_v34, _2596_v34, _2596_v34, _2596_v34)).length), (_this).fm8(p2, _2591_v29, !(true), globalState));
          let _lhs327 = globalState;
          let _lhs328 = _this;
          let _lhs329 = _2592_v30;
          let _lhs330 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length));
          _lhs327.f1 = _rhs498;
          _2590_v28 = _rhs499;
          _2591_v29 = _rhs500;
          _lhs328.f7 = _rhs501;
          _lhs329[_lhs330] = _rhs502;
          let _arr20 = _this.f7;
          let _index445 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_this.f7).length));
          _arr20[_index445] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(419)), function (_2599_i1) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }), _2596_v34);
          let _arr21 = _this.f7;
          let _index446 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_this.f7).length));
          _arr21[_index446] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2553_v1, _2553_v1), _2553_v1);
          _2593_v31 = new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ff"))).length);
          let _2600_v35;
          _2600_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2593_v31,_2596_v34);
          let _2601_v36;
          _2601_v36 = _dafny.Set.fromElements(_dafny.Seq.update(_2596_v34, _module.__default.safeIndex(_2593_v31, new BigNumber((_2596_v34).length)), new _dafny.CodePoint('f'.codePointAt(0))));
          let _2602_v37;
          _2602_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2592_v30,(_2553_v1)[_module.__default.safeIndex(_2593_v31, new BigNumber((_2553_v1).length))]);
          _2596_v34 = _dafny.Seq.Concat(_dafny.Seq.update((((_2600_v35).contains((_2592_v30)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length))])) ? ((_2600_v35).get((_2592_v30)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length))])) : (_2596_v34)), _module.__default.safeIndex(new BigNumber((_2601_v36).length), new BigNumber(((((_2600_v35).contains((_2592_v30)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length))])) ? ((_2600_v35).get((_2592_v30)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length))])) : (_2596_v34))).length)), _2552_v0), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), ((_2603_v0) => function (_2604_i2) {
            return _2603_v0;
          })(_2552_v0)), _module.__default.safeIndex(new BigNumber((_2602_v37).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), ((_2605_v0) => function (_2606_i2) {
            return _2605_v0;
          })(_2552_v0))).length)), (_module.__default.fm65((_this).f6, new BigNumber(266), (_2592_v30)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_2592_v30).length))], globalState)).dtor_cf43));
          let _arr22 = _this.f7;
          let _index447 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_this.f7).length));
          _arr22[_index447] = _2591_v29;
        }
        if (p2) {
          let _2607_v38;
          _2607_v38 = new BigNumber(788);
          (globalState).f1 = (_2607_v38).minus((_this).fm8(p2, p2, p2, globalState));
          let _2608_v39;
          _2608_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2607_v38);
          _2608_v39 = (_2608_v39).update(false, _2607_v38);
          let _2609_v40;
          _2609_v40 = _dafny.Seq.of(_2607_v38);
          (globalState).f1 = (_this).fm4(_2553_v1, new BigNumber((_2609_v40).length), (_dafny.ZERO).minus(_2607_v38), _dafny.MultiSet.fromElements(_2607_v38), globalState);
          let _2610_v41;
          let _nw439 = new _module.C3();
          _nw439.__ctor(p2, p0);
          _2610_v41 = _nw439;
          let _2611_v42;
          _2611_v42 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2607_v38));
          let _2612_v43;
          let _nw440 = new _module.C7();
          _nw440.__ctor(_2607_v38, new BigNumber((_2611_v42).cardinality()), (_this).f5, p2);
          _2612_v43 = _nw440;
        } else {
          _2552_v0 = (_this).f5;
          let _arr23 = _this.f7;
          let _index448 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_this.f7).length));
          _arr23[_index448] = (p0) === (false);
          let _2613_v44;
          _2613_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(globalState),p2);
          let _2614_v45;
          _2614_v45 = new BigNumber(120);
          let _2615_v46;
          _2615_v46 = _dafny.MultiSet.fromElements(_2614_v45);
          let _2616_v47;
          let _nw441 = new _module.C5();
          _nw441.__ctor((((_this).fm10(_2573_cf35, _2613_v44, new BigNumber(697), globalState)) ? (new BigNumber(345)) : (new BigNumber((_2615_v46).cardinality()))), (_this).f5, p2);
          _2616_v47 = _nw441;
          let _2617_v48;
          _2617_v48 = _dafny.Seq.of(new BigNumber(-920), _2614_v45, new BigNumber(128), _2614_v45);
          let _arr24 = _this.f7;
          let _index449 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_this.f7).length));
          let _rhs503 = p0;
          let _rhs504 = new BigNumber((_2553_v1).length);
          let _rhs505 = _module.__default.safeModuloInt((_2614_v45).minus((_2617_v48)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_2614_v45)), new BigNumber((_2617_v48).length))]), (((_2573_cf35).contains(p2)) ? ((_2573_cf35).get(p2)) : (_2614_v45)));
          let _rhs506 = _2616_v47;
          let _lhs331 = _this.f7;
          let _lhs332 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_this.f7).length));
          let _lhs333 = globalState;
          let _lhs334 = globalState;
          _lhs331[_lhs332] = _rhs503;
          _lhs333.f1 = _rhs504;
          _lhs334.f1 = _rhs505;
          _2616_v47 = _rhs506;
          let _2618_v49;
          _2618_v49 = _dafny.Seq.of(_dafny.Seq.Concat(_2617_v48, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-385)), function (_2619_i3) {
            return new BigNumber(205);
          })), _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm3((_2616_v47).f6, p2, globalState), new BigNumber(-206)), _2617_v48));
          let _2620_v50;
          _2620_v50 = _dafny.Seq.UnicodeFromString("fdtgl");
          _2618_v49 = _module.__default.fm38(!(new BigNumber((_2620_v50).length)).isEqualTo(_2614_v45), p0, new BigNumber((_2620_v50).length), _2552_v0, globalState);
          _2614_v45 = new BigNumber((_2617_v48).length);
          (globalState).f1 = ((_dafny.ZERO).minus(_2614_v45)).multipliedBy(_2614_v45);
        }
        let _2621_v51;
        _2621_v51 = new BigNumber(710);
        let _2622_v52;
        _2622_v52 = _dafny.MultiSet.fromElements(new BigNumber(422));
        let _2623_v53;
        _2623_v53 = _dafny.Seq.UnicodeFromString("i");
        let _2624_v54;
        _2624_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_dafny.Seq.of(false), _2621_v51, _2621_v51, _2622_v52, globalState),_2623_v53);
        (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(644), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_2624_v54).length), _2621_v51))));
        let _2625_v55;
        let _nw442 = Array((new BigNumber(29)).toNumber()).fill([]);
        _2625_v55 = _nw442;
        let _2626_v56;
        _2626_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2621_v51,(_this).f6);
        let _2627_v58;
        _2627_v58 = _dafny.Set.fromElements(_2621_v51, _2621_v51);
        let _2628_v59;
        let _nw443 = Array((new BigNumber(25)).toNumber());
        _nw443[(_dafny.ZERO).toNumber()] = _2626_v56;
        _nw443[(_dafny.ONE).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(2)).toNumber()] = function () {
          let _coll76 = new _dafny.Map();
          for (const _compr_76 of (_2627_v58).Elements) {
            let _2629_v57 = _compr_76;
            if ((_2627_v58).contains(_2629_v57)) {
              _coll76.push([(_2629_v57).plus((_dafny.ZERO).minus(_2621_v51)),p2]);
            }
          }
          return _coll76;
        }();
        _nw443[(new BigNumber(3)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_2621_v51)).length),p0);
        _nw443[(new BigNumber(5)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(6)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(7)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_2621_v51,(_this).f6))).update(_2621_v51, (_this).f6);
        _nw443[(new BigNumber(8)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(9)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(10)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(11)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(12)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(13)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2621_v51,p2);
        _nw443[(new BigNumber(15)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(16)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(17)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(18)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(19)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2621_v51,p2);
        _nw443[(new BigNumber(21)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(22)).toNumber()] = _2626_v56;
        _nw443[(new BigNumber(23)).toNumber()] = (_2626_v56).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2621_v51,p2)).length), p2);
        _nw443[(new BigNumber(24)).toNumber()] = _2626_v56;
        _2628_v59 = _nw443;
        let _index450 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_2625_v55).length));
        (_2625_v55)[_index450] = _2628_v59;
        let _index451 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_2625_v55).length));
        (_2625_v55)[_index451] = _2628_v59;
      } else {
        let _2630___mcc_h3 = (_source35).cf38;
        let _2631_cf38 = _2630___mcc_h3;
        let _2632_v60;
        _2632_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-135),_2553_v1);
        let _2633_v61;
        _2633_v61 = new BigNumber(202);
        (globalState).f1 = _module.__default.safeModuloInt(new BigNumber((_2632_v60).length), _2633_v61);
        let _2634_v62;
        let _nw444 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _2634_v62 = _nw444;
        let _index452 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length));
        (_2634_v62)[_index452] = _2633_v61;
        let _2635_v63;
        _2635_v63 = _module.D20.create_DC43(_2633_v61, p2, _2633_v61);
        let _index453 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length));
        let _rhs507 = _this.f7;
        let _rhs508 = _2633_v61;
        let _rhs509 = (_2635_v63).dtor_cf58;
        let _lhs335 = _this;
        let _lhs336 = _2634_v62;
        let _lhs337 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length));
        _lhs335.f7 = _rhs507;
        _lhs336[_lhs337] = _rhs508;
        _2633_v61 = _rhs509;
        _2633_v61 = ((p2) ? ((_dafny.ZERO).minus((_this).fm7(true, globalState))) : (new BigNumber(956)));
        let _2636_v64;
        _2636_v64 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2634_v62)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length))]),(_2553_v1)[_module.__default.safeIndex(_2633_v61, new BigNumber((_2553_v1).length))]);
        if ((p2) && ((((_2636_v64).contains(new BigNumber(176))) ? ((_2636_v64).get(new BigNumber(176))) : (p2)))) {
          (_this).f7 = _this.f7;
          _2553_v1 = _2553_v1;
          let _2637_v65;
          _2637_v65 = false;
          _2637_v65 = !(!(!((_this).f6)));
          let _2638_v66;
          _2638_v66 = _this.f7;
          _2638_v66 = _2638_v66;
          _2637_v65 = p2;
        } else {
          (globalState).f1 = _2633_v61;
          let _arr25 = _this.f7;
          let _index454 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_this.f7).length));
          _arr25[_index454] = p2;
          let _2639_v67;
          _2639_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(118));
          let _2640_v68;
          _2640_v68 = _dafny.MultiSet.fromElements((_this).f5, (_this).f5);
          let _arr26 = _this.f7;
          let _index455 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_this.f7).length));
          let _rhs510 = _dafny.Seq.update(_2553_v1, _module.__default.safeIndex(_2633_v61, new BigNumber((_2553_v1).length)), (_this).f6);
          let _rhs511 = _2552_v0;
          let _rhs512 = p2;
          let _rhs513 = _2639_v67;
          let _rhs514 = new BigNumber(((_this).fm11(_2633_v61, p0, new BigNumber((_2640_v68).cardinality()), globalState)).length);
          let _lhs338 = _this.f7;
          let _lhs339 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_this.f7).length));
          let _lhs340 = globalState;
          _2553_v1 = _rhs510;
          _2552_v0 = _rhs511;
          _lhs338[_lhs339] = _rhs512;
          _2639_v67 = _rhs513;
          _lhs340.f1 = _rhs514;
          (globalState).f1 = (_2634_v62)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length))];
          let _2641_v69;
          _2641_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this.f7)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_this.f7).length))],(_this.f7)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_this.f7).length))]);
          _2641_v69 = _2641_v69;
          (globalState).f1 = _module.__default.safeDivisionInt(((p0) ? (new BigNumber((_dafny.Seq.of((_2634_v62)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_2634_v62).length))])).length)) : (_2633_v61)), new BigNumber((p1).cardinality()));
        }
      }
      let _2642_v70;
      _2642_v70 = new BigNumber(712);
      (globalState).f1 = _2642_v70;
      let _2643_v71;
      _2643_v71 = _dafny.Seq.UnicodeFromString("fadnxqgcp");
      let _2644_v72;
      _2644_v72 = _dafny.Seq.of(_dafny.Seq.Concat(_2643_v71, _2643_v71), _2643_v71);
      _2644_v72 = _dafny.Seq.Concat(_2644_v72, _2644_v72);
      let _2645_v73;
      _2645_v73 = _dafny.MultiSet.fromElements(_2643_v71);
      let _2646_v74;
      _2646_v74 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6) || ((_2553_v1)[_module.__default.safeIndex(_2642_v70, new BigNumber((_2553_v1).length))]),_dafny.Seq.Create(_module.__default.abs(new BigNumber(261)), function (_2647_i4) {
        return (_this).f5;
      }));
      let _rhs515 = _2642_v70;
      let _rhs516 = (((_2646_v74).contains((_this).f6)) ? ((_2646_v74).get((_this).f6)) : (_2643_v71));
      let _rhs517 = (_2642_v70).minus(_2642_v70);
      let _rhs518 = (_2645_v73).Intersect(((_dafny.MultiSet.fromElements(_2643_v71, _2643_v71)).update(_dafny.Seq.UnicodeFromString("anuxkuyu"), _module.__default.abs(new BigNumber(-918)))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xbkjrf"))));
      let _lhs341 = globalState;
      _lhs341.f1 = _rhs515;
      _2643_v71 = _rhs516;
      _2642_v70 = _rhs517;
      _2645_v73 = _rhs518;
      let _2648_v75;
      _2648_v75 = _dafny.Set.fromElements(_this.f7, _this.f7);
      _2648_v75 = ((p0) ? (_2648_v75) : (_dafny.Set.fromElements(_this.f7)));
      return;
    }
    m6(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let r2 = false;
      let _2649_v0;
      _2649_v0 = _dafny.Seq.of(p0);
      r1 = !_dafny.areEqual(_2649_v0, _2649_v0);
      let _2650_v1;
      _2650_v1 = new BigNumber(556);
      let _2651_v2;
      _2651_v2 = _module.D21.create_DC46(p0, p1, _2650_v1);
      r2 = (_2651_v2).dtor_cf61;
      let _2652_v3;
      let _nw445 = Array((new BigNumber(2)).toNumber());
      _nw445[(_dafny.ZERO).toNumber()] = _2650_v1;
      _nw445[(_dafny.ONE).toNumber()] = _2650_v1;
      _2652_v3 = _nw445;
      let _index456 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length));
      (_2652_v3)[_index456] = _2650_v1;
      let _index457 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length));
      (_2652_v3)[_index457] = _module.__default.safeModuloInt((_2650_v1).plus(_2650_v1), new BigNumber(33));
      r2 = (_this).f6;
      let _2653_v4;
      _2653_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _2654_v5;
      _2654_v5 = _module.D15.create_DC28((((_2653_v4).contains(_module.__default.fm0(p1, globalState))) ? ((_2653_v4).get(_module.__default.fm0(p1, globalState))) : ((_this).f6)), (_this).f6);
      let _2655_i0;
      _2655_i0 = _dafny.ZERO;
      L24: {
        while ((_2654_v5).dtor_cf37) {
          C24: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2655_i0)) {
              break L24;
            }
            _2655_i0 = (_2655_i0).plus(_dafny.ONE);
            let _2656_v6;
            _2656_v6 = _module.D0.create_DC0(p1);
            let _2657_v7;
            let _nw446 = new _module.C3();
            _nw446.__ctor(p1, (_this).fm9((_dafny.ZERO).minus(_module.__default.fm3(p1, p0, globalState)), _2656_v6, _module.__default.fm0(p0, globalState), _2650_v1, globalState));
            _2657_v7 = _nw446;
            let _2658_v8;
            _2658_v8 = _dafny.Seq.of(_2657_v7);
            let _2659_v9;
            _2659_v9 = _dafny.Seq.of(((!(p0)) ? (_2657_v7) : (_2657_v7)), (_2658_v8)[_module.__default.safeIndex((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], new BigNumber((_2658_v8).length))], _2657_v7);
            let _2660_v10;
            _2660_v10 = _dafny.Seq.of((_dafny.ZERO).minus((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]), (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], _2650_v1, (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]);
            _2657_v7 = (_2659_v9)[_module.__default.safeIndex(new BigNumber((_2660_v10).length), new BigNumber((_2659_v9).length))];
            let _2661_v11;
            _2661_v11 = _dafny.MultiSet.fromElements(_2650_v1);
            let _2662_v12;
            _2662_v12 = _dafny.Set.fromElements((_this).fm4(_2649_v0, new BigNumber(914), _2650_v1, _2661_v11, globalState));
            _2660_v10 = _dafny.Seq.of(_2650_v1, _2650_v1, new BigNumber(((_2662_v12).Union(_2662_v12)).length), (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]);
            let _index458 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length));
            (_2652_v3)[_index458] = _module.__default.fm3(!((_2657_v7).f14), p0, globalState);
            r2 = (_2662_v12).IsDisjointFrom(_2662_v12);
          }
        }
      }
      let _2663_v13;
      _2663_v13 = _module.D6.create_DC14((_this).f6);
      let _2664_i1;
      _2664_i1 = _dafny.ZERO;
      L25: {
        while (((((_this).f6) ? (_module.D6.create_DC14(p1)) : (_2663_v13))).dtor_cf15) {
          C25: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2664_i1)) {
              break L25;
            }
            _2664_i1 = (_2664_i1).plus(_dafny.ONE);
            let _2665_v14;
            _2665_v14 = _module.D0.create_DC1();
            let _2666_v15;
            _2666_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2665_v14,new BigNumber(25));
            _2666_v15 = (_2666_v15).Merge(_2666_v15);
            let _arr27 = _this.f7;
            let _index459 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
            _arr27[_index459] = p1;
            let _arr28 = _this.f7;
            let _index460 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
            _arr28[_index460] = ((!((_this).f6)) ? ((_this).f6) : (p1));
            let _2667_v16;
            let _nw447 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
            _2667_v16 = _nw447;
            let _2668_v17;
            _2668_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2667_v16,!((_this).f6));
            _2668_v17 = (_2668_v17).update(_2667_v16, p0);
            let _2669_v18;
            _2669_v18 = _dafny.MultiSet.fromElements(true, p0);
            let _2670_v19;
            _2670_v19 = _module.D15.create_DC27((_2669_v18).Union(_dafny.MultiSet.FromArray(_2649_v0)));
            let _source36 = _2670_v19;
            if (_source36.is_DC28) {
              let _2671___mcc_h0 = (_source36).cf36;
              let _2672___mcc_h1 = (_source36).cf37;
              let _2673_cf37 = _2672___mcc_h1;
              let _2674_cf36 = _2671___mcc_h0;
              let _2675_v20;
              _2675_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2674_cf36,_dafny.MultiSet.fromElements(_2650_v1, _2650_v1, (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]));
              let _2676_v21;
              _2676_v21 = _dafny.MultiSet.fromElements(new BigNumber((_2675_v20).length));
              let _2677_v22;
              _2677_v22 = _dafny.MultiSet.fromElements(new BigNumber((_2676_v21).cardinality()));
              let _2678_v24;
              _2678_v24 = function () {
                let _coll77 = new _dafny.Map();
                for (const _compr_77 of _dafny.IntegerRange(new BigNumber(138), new BigNumber(86))) {
                  let _2679_v23 = _compr_77;
                  if (((new BigNumber(138)).isLessThanOrEqualTo(_2679_v23)) && ((_2679_v23).isLessThan(new BigNumber(86)))) {
                    _coll77.push([(_2679_v23).minus(new BigNumber(-264)),false]);
                  }
                }
                return _coll77;
              }();
              let _2680_v25;
              _2680_v25 = _dafny.Map.Empty.slice().updateUnsafe((_2676_v21).contains(_2650_v1),_2678_v24);
              _2680_v25 = (_2680_v25).update(!(true), ((p1) ? (_2678_v24) : (_2678_v24)));
              let _2681_v26;
              _2681_v26 = _dafny.Set.fromElements(_2673_cf37);
              let _2682_v27;
              _2682_v27 = _dafny.Seq.of(_2681_v26);
              let _2683_v28;
              _2683_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2673_cf37,new BigNumber((_2649_v0).length));
              let _2684_v29;
              let _nw448 = Array((new BigNumber(8)).toNumber());
              _nw448[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(true);
              _nw448[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(p1)).Difference(_2681_v26);
              _nw448[(new BigNumber(2)).toNumber()] = _2681_v26;
              _nw448[(new BigNumber(3)).toNumber()] = (_2682_v27)[_module.__default.safeIndex(new BigNumber((_2683_v28).length), new BigNumber((_2682_v27).length))];
              _nw448[(new BigNumber(4)).toNumber()] = _module.__default.fm17(_2650_v1, p1, globalState);
              _nw448[(new BigNumber(5)).toNumber()] = _module.__default.fm17(_2650_v1, (_this).f6, globalState);
              _nw448[(new BigNumber(6)).toNumber()] = (_2681_v26).Union(_2681_v26);
              _nw448[(new BigNumber(7)).toNumber()] = _module.__default.fm49(globalState);
              _2684_v29 = _nw448;
              let _2685_v30;
              let _nw449 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _2685_v30 = _nw449;
              let _index461 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_2685_v30).length));
              (_2685_v30)[_index461] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-409)), function (_2686_i2) {
                return (_this).f5;
              });
              let _2687_v31;
              _2687_v31 = _dafny.Seq.UnicodeFromString("ayjjit");
              let _index462 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_2685_v30).length));
              let _rhs519 = _2684_v29;
              let _rhs520 = _2687_v31;
              let _lhs342 = _2685_v30;
              let _lhs343 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_2685_v30).length));
              _2684_v29 = _rhs519;
              _lhs342[_lhs343] = _rhs520;
              (globalState).f1 = ((!((p0) === (_2673_cf37))) ? ((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]) : (_module.__default.safeDivisionInt((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], new BigNumber((_module.__default.fm1(new BigNumber(597), _2674_cf36, (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], globalState)).length))));
              let _index463 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_2685_v30).length));
              (_2685_v30)[_index463] = _2687_v31;
            } else if (_source36.is_DC27) {
              let _2688___mcc_h2 = (_source36).cf35;
              let _2689_cf35 = _2688___mcc_h2;
              r1 = false;
              (globalState).f1 = (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))];
              let _2690_v32;
              _2690_v32 = _dafny.Seq.of((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]);
              let _2691_v33;
              let _nw450 = Array((new BigNumber(13)).toNumber());
              _nw450[(_dafny.ZERO).toNumber()] = !((((_this).f6) ? (p1) : (p0)));
              _nw450[(_dafny.ONE).toNumber()] = p1;
              _nw450[(new BigNumber(2)).toNumber()] = (_2649_v0)[_module.__default.safeIndex(_2650_v1, new BigNumber((_2649_v0).length))];
              _nw450[(new BigNumber(3)).toNumber()] = p1;
              _nw450[(new BigNumber(4)).toNumber()] = p1;
              _nw450[(new BigNumber(5)).toNumber()] = (_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))];
              _nw450[(new BigNumber(6)).toNumber()] = _dafny.areEqual(_2690_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_2692_v1) => function (_2693_i3) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2692_v1,(_this).f5)).length);
              })(_2650_v1)));
              _nw450[(new BigNumber(7)).toNumber()] = (_this).f6;
              _nw450[(new BigNumber(8)).toNumber()] = !((_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))]) || (p0);
              _nw450[(new BigNumber(9)).toNumber()] = (_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))];
              _nw450[(new BigNumber(10)).toNumber()] = (_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))];
              _nw450[(new BigNumber(11)).toNumber()] = p0;
              _nw450[(new BigNumber(12)).toNumber()] = !_dafny.Seq.contains(_2690_v32, _2650_v1);
              _2691_v33 = _nw450;
              let _2694_v34;
              _2694_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]).plus(_2650_v1));
              let _index464 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length));
              let _arr29 = _this.f7;
              let _index465 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
              let _rhs521 = _2650_v1;
              let _rhs522 = _2691_v33;
              let _rhs523 = _2691_v33;
              let _rhs524 = _2694_v34;
              let _rhs525 = (_this).f6;
              let _lhs344 = _2652_v3;
              let _lhs345 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length));
              let _lhs346 = _this.f7;
              let _lhs347 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
              _lhs344[_lhs345] = _rhs521;
              _2691_v33 = _rhs522;
              _2691_v33 = _rhs523;
              _2694_v34 = _rhs524;
              _lhs346[_lhs347] = _rhs525;
              let _2695_v35;
              _2695_v35 = _dafny.Set.fromElements((_this).f6);
              let _2696_v36;
              _2696_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2695_v35,(_this).f6);
              let _arr30 = _this.f7;
              let _index466 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
              _arr30[_index466] = !((_this).fm10(_2669_v18, _2696_v36, ((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]).minus(_2650_v1), globalState));
            } else {
              let _2697___mcc_h3 = (_source36).cf38;
              let _2698_cf38 = _2697___mcc_h3;
              let _2699_v37;
              _2699_v37 = _dafny.Seq.of(_2650_v1);
              let _2700_v38;
              _2700_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(709),_dafny.Seq.update(_2699_v37, _module.__default.safeIndex((_this).fm8(true, (_this).f6, (_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))], globalState), new BigNumber((_2699_v37).length)), (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]));
              let _2701_v39;
              _2701_v39 = _dafny.Seq.of((((_2700_v38).contains((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))])) ? ((_2700_v38).get((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))])) : (_2699_v37)), _2699_v37);
              let _2702_v40;
              _2702_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2701_v39).length),true);
              _2702_v40 = (_2702_v40).update((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))], !((_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))]));
              let _2703_v41;
              _2703_v41 = _dafny.Set.fromElements(_2650_v1);
              let _2704_v42;
              _2704_v42 = _dafny.Seq.UnicodeFromString("gxqw");
              let _2705_v43;
              _2705_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2703_v41).length),(_this).fm8((_2649_v0)[_module.__default.safeIndex(new BigNumber((_2704_v42).length), new BigNumber((_2649_v0).length))], p1, (_this.f7)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length))], globalState));
              _2705_v43 = (_2705_v43).update(_2650_v1, (_this).fm8((_this).f6, false, true, globalState));
              let _arr31 = _this.f7;
              let _index467 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_this.f7).length));
              _arr31[_index467] = (_this).f6;
              r1 = false;
            }
          }
        }
      }
      let _2706_v44;
      _2706_v44 = _dafny.Seq.of((_dafny.ZERO).minus((_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]), _2650_v1);
      r0 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_2706_v44, _2706_v44), _dafny.Seq.of((_dafny.ZERO).minus(_2650_v1), (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))])), _module.__default.safeIndex((_2650_v1).multipliedBy(_2650_v1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2706_v44, _2706_v44), _dafny.Seq.of((_dafny.ZERO).minus(_2650_v1), (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]))).length)), _2650_v1);
      let _2707_v45;
      _2707_v45 = _module.D11.create_DC20(p0, (_2652_v3)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_2652_v3).length))]);
      r1 = (_2707_v45).dtor_cf21;
      r2 = !(p1);
      return [r0, r1, r2];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
