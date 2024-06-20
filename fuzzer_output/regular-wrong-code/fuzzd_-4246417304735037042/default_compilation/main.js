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
    static fm0(p0, p1, p2, globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(773), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(506))).length))).length))).Elements) {
          let _0_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(773), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(506))).length))).length)), _0_v0)) {
            _coll0.add((_0_v0).plus(new BigNumber(357)));
          }
        }
        return _coll0;
      }()).Difference(((true) ? (_dafny.Set.fromElements(new BigNumber(766))) : (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(412))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(98)), function (_1_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length)))));
    };
    static fm1(globalState) {
      return new BigNumber((((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D8.create_DC18(new BigNumber((_dafny.Seq.UnicodeFromString("khb")).length))))).Union(_dafny.MultiSet.fromElements(_module.D8.create_DC18(new BigNumber(-677))))).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D8.create_DC18(new BigNumber(-488))))).Intersect(_dafny.MultiSet.fromElements(_module.D8.create_DC18(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).length)))))).cardinality());
    };
    static fm2(p0, p1, globalState) {
      if (true) {
        return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("bhudtnhvn"), new _dafny.CodePoint('w'.codePointAt(0)));
      } else {
        return !(new BigNumber(-773)).isEqualTo(new BigNumber(-371));
      }
    };
    static fm4(p0, globalState) {
      return new BigNumber(4);
    };
    static fm8(p0, globalState) {
      let _source0 = _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),_dafny.Seq.UnicodeFromString("xtkckav")));
      if (_source0.is_DC16) {
        let _2___mcc_h0 = (_source0).cf26;
        let _3___mcc_h1 = (_source0).cf27;
        let _4___mcc_h2 = (_source0).cf28;
        let _5___mcc_h3 = (_source0).cf29;
        let _6_cf29 = _5___mcc_h3;
        let _7_cf28 = _4___mcc_h2;
        let _8_cf27 = _3___mcc_h1;
        let _9_cf26 = _2___mcc_h0;
        return _module.D3.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(26),_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_9_cf26))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wlkdfem")).length)), new BigNumber(610))));
      } else {
        let _10___mcc_h4 = (_source0).cf25;
        let _11_cf25 = _10___mcc_h4;
        if (false) {
          return _module.D3.create_DC9(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(974), new BigNumber(89))) {
    let _12_v0 = _compr_1;
    if (((new BigNumber(974)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(89)))) {
      _coll1.push([(_12_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_13_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length))),_dafny.Set.fromElements(new BigNumber(74), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber(799), new BigNumber((_dafny.Seq.UnicodeFromString("d")).length))).length))]);
    }
  }
  return _coll1;
}());
        } else {
          return _module.D3.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-232),_dafny.Set.fromElements(new BigNumber(822))));
        }
      }
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("njtrptc"), _dafny.Seq.UnicodeFromString("l"));
    };
    static fm10(p0, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("yn"), ((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_14_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })) : (_dafny.Seq.UnicodeFromString("uegdnut"))));
    };
    static fm11(p0, p1, p2, globalState) {
      if ((_dafny.MultiSet.fromElements(false)).IsDisjointFrom(_dafny.MultiSet.fromElements(false, true))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(false,true);
      }
    };
    static fm16(p0, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber(-719), new BigNumber(-755))).isEqualTo(new BigNumber((((false) ? (_dafny.Set.fromElements(!(false))) : (_dafny.Set.fromElements(true, !(false))))).length));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(122)), function (_15_i0) {
        return new BigNumber((_dafny.Set.fromElements(false)).length);
      }), _dafny.Seq.of(new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
              let _16_v2 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), _16_v2)) {
                _coll4.push([_16_v2,(_module.D9.create_DC23(true)).dtor_cf38]);
              }
            }
            return _coll4;
          }()).Keys.Elements) {
            let _17_v1 = _compr_3;
            if ((function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
                let _16_v2 = _compr_5;
                if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), _16_v2)) {
                  _coll5.push([_16_v2,(_module.D9.create_DC23(true)).dtor_cf38]);
                }
              }
              return _coll5;
            }()).contains(_17_v1)) {
              _coll3.push([_17_v1,new BigNumber((_dafny.Set.fromElements(new BigNumber(-244), new BigNumber(369))).length)]);
            }
          }
          return _coll3;
        }()).Keys.Elements) {
          let _18_v0 = _compr_2;
          if ((function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (function () {
              let _coll7 = new _dafny.Map();
              for (const _compr_7 of (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
                let _16_v2 = _compr_7;
                if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), _16_v2)) {
                  _coll7.push([_16_v2,(_module.D9.create_DC23(true)).dtor_cf38]);
                }
              }
              return _coll7;
            }()).Keys.Elements) {
              let _17_v1 = _compr_6;
              if ((function () {
                let _coll8 = new _dafny.Map();
                for (const _compr_8 of (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
                  let _16_v2 = _compr_8;
                  if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), _16_v2)) {
                    _coll8.push([_16_v2,(_module.D9.create_DC23(true)).dtor_cf38]);
                  }
                }
                return _coll8;
              }()).contains(_17_v1)) {
                _coll6.push([_17_v1,new BigNumber((_dafny.Set.fromElements(new BigNumber(-244), new BigNumber(369))).length)]);
              }
            }
            return _coll6;
          }()).contains(_18_v0)) {
            _coll2.push([_18_v0,new BigNumber(866)]);
          }
        }
        return _coll2;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(567),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cpapaci")).length)))).length)));
    };
    static fm18(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat((_module.D0.create_DC2(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(739), new BigNumber(-766)),_module.D1.create_DC7(new _dafny.CodePoint('q'.codePointAt(0)), _dafny.Seq.UnicodeFromString("p")))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-981)), function (_19_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}))).dtor_cf6, _dafny.Seq.UnicodeFromString("rvlldbjbx"));
    };
    static fm21(p0, p1, globalState) {
      return _module.D0.create_DC2((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-240)), function (_20_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length))).length))).Union(function () {
  let _coll9 = new _dafny.Set();
  for (const _compr_9 of (function () {
    let _coll10 = new _dafny.Set();
    for (const _compr_10 of _dafny.IntegerRange(new BigNumber(951), new BigNumber(-986))) {
      let _21_v0 = _compr_10;
      if (((new BigNumber(951)).isLessThanOrEqualTo(_21_v0)) && ((_21_v0).isLessThan(new BigNumber(-986)))) {
        _coll10.add(_module.__default.safeDivisionInt(_21_v0, new BigNumber(28)));
      }
    }
    return _coll10;
  }()).Elements) {
    let _22_v1 = _compr_9;
    if ((function () {
      let _coll11 = new _dafny.Set();
      for (const _compr_11 of _dafny.IntegerRange(new BigNumber(951), new BigNumber(-986))) {
        let _23_v0 = _compr_11;
        if (((new BigNumber(951)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(-986)))) {
          _coll11.add(_module.__default.safeDivisionInt(_23_v0, new BigNumber(28)));
        }
      }
      return _coll11;
    }()).contains(_22_v1)) {
      _coll9.add(_module.__default.safeDivisionInt(_22_v1, new BigNumber(-818)));
    }
  }
  return _coll9;
}()), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), function (_24_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("oiuswc")));
    };
    static fm22(p0, p1, globalState) {
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sxdj"),new BigNumber((_dafny.Seq.UnicodeFromString("ypeb")).length))).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))))).isEqualTo(((!(!(false))) ? (new BigNumber((_dafny.Seq.UnicodeFromString("cn")).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(800)), function (_25_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(708)), function (_26_i1) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length);
      })).length))));
    };
    static fm23(globalState) {
      return _module.D1.create_DC6(new BigNumber((_dafny.Set.fromElements(false)).length));
    };
    static fm24(globalState) {
      return ((_dafny.Set.fromElements(_module.D1.create_DC6(new BigNumber(-922)))).Union(_dafny.Set.fromElements(_module.D1.create_DC6(new BigNumber(-649)), _module.D1.create_DC6(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(922)), function (_27_i0) {
  return new _dafny.CodePoint('d'.codePointAt(0));
})).length)), _module.D1.create_DC6(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(656), new BigNumber((_dafny.Seq.UnicodeFromString("qjdg")).length), new BigNumber((_dafny.Seq.of(false, false)).length)))).cardinality()))))).Intersect(function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of (function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_dafny.Seq.of(_module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),false)).length)))).Elements) {
            let _28_v0 = _compr_13;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),false)).length))), _28_v0)) {
              _coll13.add(_28_v0);
            }
          }
          return _coll13;
        }()).Elements) {
          let _29_v1 = _compr_12;
          if ((function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_dafny.Seq.of(_module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),false)).length)))).Elements) {
              let _30_v0 = _compr_14;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),false)).length))), _30_v0)) {
                _coll14.add(_30_v0);
              }
            }
            return _coll14;
          }()).contains(_29_v1)) {
            _coll12.add(_29_v1);
          }
        }
        return _coll12;
      }());
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(true), true, true), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(!(true), false)));
    };
    static fm26(p0, p1, p2, globalState) {
      if (_dafny.areEqual(_dafny.Seq.of(false), _dafny.Seq.of(false, !(false)))) {
        return (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false, true, true, false));
      } else {
        return (_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(false, false));
      }
    };
    static fm27(p0, p1, p2, globalState) {
      return (function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of _dafny.IntegerRange(new BigNumber(456), new BigNumber(-349))) {
          let _31_v0 = _compr_15;
          if (((new BigNumber(456)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(-349)))) {
            _coll15.push([(_31_v0).plus(new BigNumber(131)),_module.D1.create_DC7(new _dafny.CodePoint('e'.codePointAt(0)), _dafny.Seq.UnicodeFromString("gstwlmab"))]);
          }
        }
        return _coll15;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false, !(!(false))))).cardinality()),_module.D1.create_DC7(new _dafny.CodePoint('a'.codePointAt(0)), _dafny.Seq.UnicodeFromString("leq"))));
    };
    static fm28(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-714)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-198)), function (_32_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length)),_dafny.Seq.UnicodeFromString("emwqfdps"));
    };
    static fm29(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm30(p0, p1, p2, globalState) {
      let _source1 = _module.D1.create_DC6(new BigNumber(135));
      if (_source1.is_DC6) {
        let _33___mcc_h0 = (_source1).cf12;
        let _34_cf12 = _33___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(271),true);
      } else if (_source1.is_DC7) {
        let _35___mcc_h1 = (_source1).cf13;
        let _36___mcc_h2 = (_source1).cf14;
        let _37_cf14 = _36___mcc_h2;
        let _38_cf13 = _35___mcc_h1;
        return (function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(556), new BigNumber(947))) {
            let _39_v0 = _compr_16;
            if (((new BigNumber(556)).isLessThanOrEqualTo(_39_v0)) && ((_39_v0).isLessThan(new BigNumber(947)))) {
              _coll16.push([(_39_v0).plus(new BigNumber(-346)),true]);
            }
          }
          return _coll16;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-892),true));
      } else {
        let _40___mcc_h3 = (_source1).cf11;
        let _41_cf11 = _40___mcc_h3;
        return (function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(-751), new BigNumber(772))) {
            let _42_v1 = _compr_17;
            if (((new BigNumber(-751)).isLessThanOrEqualTo(_42_v1)) && ((_42_v1).isLessThan(new BigNumber(772)))) {
              _coll17.push([(_42_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))),_dafny.Seq.UnicodeFromString("fwcp"))).length)),false]);
            }
          }
          return _coll17;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kkb")).length),false));
      }
    };
    static fm31(globalState) {
      return ((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(!(!(true)), true, true, false))).Union(((!(true)) ? (_dafny.Set.fromElements(true)) : (_dafny.Set.fromElements(true))));
    };
    static fm32(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("rxmfcp");
    };
    static fm33(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(840)), function (_43_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kyandbjqq")).length), new BigNumber((_dafny.Seq.of(true)).length));
      });
    };
    static fm36(globalState) {
      return new BigNumber(-592);
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!((_module.D0.create_DC3(new BigNumber(330), true, false)).dtor_cf8),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),false)));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true))).Difference(_dafny.Set.fromElements(!(false)));
    };
    static fm39(globalState) {
      let _source2 = _module.D11.create_DC28();
      if (_source2.is_DC28) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_44_i0) {
          return _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(995),_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), function (_45_i1) {
  return new _dafny.CodePoint('c'.codePointAt(0));
})));
        });
      } else {
        let _46___mcc_h0 = (_source2).cf42;
        let _47_cf42 = _46___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(560)), function (_48_i2) {
          return _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(881),_dafny.Seq.UnicodeFromString("rqaiipv")));
        }), _dafny.Seq.of(_module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),_dafny.Seq.UnicodeFromString("tnwkwu"))), _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-651)), function (_49_i3) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_50_i4) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}))), _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qdadesapq")).length)),_dafny.Seq.UnicodeFromString("mdilqo"))), _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, false, true)).length)),_dafny.Seq.UnicodeFromString("cmgqbmpj")))));
      }
    };
    static fm40(globalState) {
      return !(!_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(22)), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(((_module.D14.create_DC31(_dafny.Seq.of(true))).dtor_cf45).length)))).cardinality()))));
    };
    static fm41(p0, globalState) {
      return _module.D11.create_DC28();
    };
    static fm42(p0, globalState) {
      return new _dafny.CodePoint('q'.codePointAt(0));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)),false),false)).Keys.Elements) {
          let _51_v0 = _compr_18;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)),false),false)).contains(_51_v0)) {
            _coll18.add(_51_v0);
          }
        }
        return _coll18;
      }()).length), new BigNumber(-731))).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(87)), function (_52_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length))).Merge((_module.D15.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(923), new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber(496)))).dtor_cf48)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)))).length))),new BigNumber(417))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),new BigNumber(310))));
    };
    static fm44(globalState) {
      return _module.D0.create_DC1();
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))))).Union(((!(!(true))) ? ((_module.D16.create_DC36(_dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0))))).dtor_cf50) : (_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0))))));
    };
    static fm46(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("yvjprk")).length)),new BigNumber(800))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-433),new BigNumber(576)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-714)), function (_53_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))).length))).length)),new BigNumber((function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(270)), function (_54_i1) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-971)), function (_55_i2) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length));
        }))).Elements) {
          let _56_v0 = _compr_19;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(270)), function (_54_i1) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-971)), function (_55_i2) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            })).length));
          })), _56_v0)) {
            _coll19.push([_56_v0,new BigNumber(460)]);
          }
        }
        return _coll19;
      }()).length)));
    };
    static fm47(p0, p1, p2, globalState) {
      return _module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(315), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,!(false)))).length), new BigNumber(228), new BigNumber(809), new BigNumber((_dafny.Seq.UnicodeFromString("ywclxcu")).length))));
    };
    static fm48(p0, p1, p2, p3, globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-356)))).length)).isLessThan(new BigNumber(90))) {
        return _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length))).length)),_dafny.Seq.UnicodeFromString("khcfpyqr")));
      } else {
        return _module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D18.create_DC40(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(699)))).dtor_cf57).length),_dafny.Seq.UnicodeFromString("pqprsunx")));
      }
    };
    static fm49(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), function (_57_i0) {
        return _dafny.Seq.of(true, false, true, false);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), function (_58_i1) {
        return _dafny.Seq.of(false);
      })), _dafny.Seq.of(_dafny.Seq.of(false)));
    };
    static fm50(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(!(false))))),new BigNumber(615))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-21)));
    };
    static fm51(p0, globalState) {
      let _source3 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)));
      let _59___mcc_h0 = _source3;
      let _60_cf19 = _59___mcc_h0;
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('r'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('y'.codePointAt(0))));
    };
    static fm52(p0, p1, p2, globalState) {
      return ((!(false)) ? (_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(false))) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(true, false))));
    };
    static fm53(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, true, false)),!(true))).Merge(function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, !(true)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-591))).cardinality()))).Keys.Elements) {
          let _61_v0 = _compr_20;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, !(true)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-591))).cardinality()))).contains(_61_v0)) {
            _coll20.push([_61_v0,false]);
          }
        }
        return _coll20;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, false),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true)));
    };
    static fm54(p0, p1, p2, p3, globalState) {
      let _source4 = ((false) ? (_module.D7.create_DC16(true, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(330),!(!(true))), new _dafny.CodePoint('h'.codePointAt(0)), true)) : (_module.D7.create_DC16(false, function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (function () {
    let _coll22 = new _dafny.Map();
    for (const _compr_22 of _dafny.IntegerRange(new BigNumber(313), new BigNumber(-556))) {
      let _62_v1 = _compr_22;
      if (((new BigNumber(313)).isLessThanOrEqualTo(_62_v1)) && ((_62_v1).isLessThan(new BigNumber(-556)))) {
        _coll22.push([(_62_v1).multipliedBy(new BigNumber(-215)),false]);
      }
    }
    return _coll22;
  }()).Keys.Elements) {
    let _63_v0 = _compr_21;
    if ((function () {
      let _coll23 = new _dafny.Map();
      for (const _compr_23 of _dafny.IntegerRange(new BigNumber(313), new BigNumber(-556))) {
        let _62_v1 = _compr_23;
        if (((new BigNumber(313)).isLessThanOrEqualTo(_62_v1)) && ((_62_v1).isLessThan(new BigNumber(-556)))) {
          _coll23.push([(_62_v1).multipliedBy(new BigNumber(-215)),false]);
        }
      }
      return _coll23;
    }()).contains(_63_v0)) {
      _coll21.push([_module.__default.safeModuloInt(_63_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-921))).length)),false]);
    }
  }
  return _coll21;
}(), new _dafny.CodePoint('r'.codePointAt(0)), true)));
      if (_source4.is_DC16) {
        let _64___mcc_h0 = (_source4).cf26;
        let _65___mcc_h1 = (_source4).cf27;
        let _66___mcc_h2 = (_source4).cf28;
        let _67___mcc_h3 = (_source4).cf29;
        let _68_cf29 = _67___mcc_h3;
        let _69_cf28 = _66___mcc_h2;
        let _70_cf27 = _65___mcc_h1;
        let _71_cf26 = _64___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-156)),_module.D8.create_DC20(_module.D8.create_DC19())),_dafny.Set.fromElements(_69_cf28))).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of _dafny.IntegerRange(new BigNumber(259), new BigNumber(610))) {
            let _72_v2 = _compr_24;
            if (((new BigNumber(259)).isLessThanOrEqualTo(_72_v2)) && ((_72_v2).isLessThan(new BigNumber(610)))) {
              _coll24.push([(_72_v2).plus(new BigNumber(216)),_module.D8.create_DC20(_module.D8.create_DC18(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-865),_71_cf26)).length)))]);
            }
          }
          return _coll24;
        }(),_dafny.Set.fromElements(_69_cf28)));
      } else {
        let _73___mcc_h4 = (_source4).cf25;
        let _74_cf25 = _73___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(35),_module.D8.create_DC20(_module.D8.create_DC20(_module.D8.create_DC19()))),_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))));
      }
    };
    static fm55(p0, globalState) {
      let _source5 = _module.D16.create_DC36(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))));
      if (_source5.is_DC37) {
        let _75___mcc_h0 = (_source5).cf51;
        let _76___mcc_h1 = (_source5).cf52;
        let _77_cf52 = _76___mcc_h1;
        let _78_cf51 = _75___mcc_h0;
        return _dafny.Seq.of(new BigNumber(616));
      } else if (_source5.is_DC38) {
        let _79___mcc_h2 = (_source5).cf53;
        let _80___mcc_h3 = (_source5).cf54;
        let _81___mcc_h4 = (_source5).cf55;
        let _82_cf55 = _81___mcc_h4;
        let _83_cf54 = _80___mcc_h3;
        let _84_cf53 = _79___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_82_cf55), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-104)), ((_85_cf55) => function (_86_i0) {
          return _85_cf55;
        })(_82_cf55)));
      } else {
        let _87___mcc_h5 = (_source5).cf50;
        let _88_cf50 = _87___mcc_h5;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-5)), function (_89_i1) {
          return new BigNumber((_dafny.Seq.of(true)).length);
        });
      }
    };
    static fm56(p0, p1, globalState) {
      return function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, !(false)),true))).Keys.Elements) {
          let _90_v0 = _compr_25;
          if (((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, !(false)),true))).contains(_90_v0)) {
            _coll25.push([_90_v0,(_module.D3.create_DC10(true, false)).dtor_cf17]);
          }
        }
        return _coll25;
      }();
    };
    static m0(globalState) {
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let _91_v0;
      _91_v0 = new BigNumber(805);
      let _92_v1;
      _92_v1 = false;
      let _93_v2;
      _93_v2 = _dafny.MultiSet.fromElements(_92_v1, _92_v1);
      let _94_v3;
      _94_v3 = _module.D10.create_DC25(_93_v2);
      let _95_v4;
      _95_v4 = _dafny.Seq.of(_94_v3);
      let _rhs0 = _91_v0;
      let _rhs1 = _91_v0;
      let _rhs2 = (_91_v0).isEqualTo(new BigNumber((_95_v4).length));
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      r1 = _rhs0;
      _lhs0.f4 = _rhs1;
      _lhs1.f6 = _rhs2;
      let _96_v5;
      let _nw0 = new _module.C4();
      _nw0.__ctor();
      _96_v5 = _nw0;
      let _97_v6;
      _97_v6 = _dafny.Map.Empty.slice().updateUnsafe(_92_v1,true);
      let _98_v7;
      _98_v7 = _dafny.MultiSet.fromElements(_91_v0);
      let _99_v8;
      _99_v8 = _dafny.MultiSet.fromElements(new BigNumber((_98_v7).cardinality()));
      let _100_v9;
      _100_v9 = _dafny.Seq.of(new BigNumber((_99_v8).cardinality()), new BigNumber(32), _91_v0);
      let _101_v10;
      _101_v10 = _dafny.Seq.UnicodeFromString("sxtwt");
      let _102_v11;
      let _nw1 = Array((new BigNumber(6)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = new BigNumber((_100_v9).length);
      _nw1[(_dafny.ONE).toNumber()] = (((_98_v7).contains(_91_v0)) ? ((_98_v7).get(_91_v0)) : (_91_v0));
      _nw1[(new BigNumber(2)).toNumber()] = _91_v0;
      _nw1[(new BigNumber(3)).toNumber()] = new BigNumber((_101_v10).length);
      _nw1[(new BigNumber(4)).toNumber()] = _91_v0;
      _nw1[(new BigNumber(5)).toNumber()] = _91_v0;
      _102_v11 = _nw1;
      let _103_v12;
      let _init0 = ((_104_v1) => function (_105_i2) {
        return _104_v1;
      })(_92_v1);
      let _nw2 = Array((new BigNumber(5)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _103_v12 = _nw2;
      let _106_v13;
      let _nw3 = Array((new BigNumber(9)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = _103_v12;
      _nw3[(_dafny.ONE).toNumber()] = _103_v12;
      _nw3[(new BigNumber(2)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(3)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(4)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(5)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(6)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(7)).toNumber()] = _103_v12;
      _nw3[(new BigNumber(8)).toNumber()] = _103_v12;
      _106_v13 = _nw3;
      let _107_v14;
      _107_v14 = _module.D9.create_DC22(_102_v11, _92_v1, _106_v13, _101_v10);
      let _108_v15;
      _108_v15 = _dafny.Seq.of(_module.__default.fm16(_module.__default.fm11((_107_v14).dtor_cf35, _92_v1, _92_v1, globalState), globalState), !(!(true)), _92_v1);
      let _109_v16;
      let _nw4 = Array((new BigNumber(18)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = (_module.__default.fm16(_97_v6, globalState)) || (_92_v1);
      _nw4[(_dafny.ONE).toNumber()] = _92_v1;
      _nw4[(new BigNumber(2)).toNumber()] = !(_92_v1);
      _nw4[(new BigNumber(3)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(4)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(5)).toNumber()] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), ((_110_v0) => function (_111_i1) {
        return _110_v0;
      })(_91_v0)), _100_v9);
      _nw4[(new BigNumber(6)).toNumber()] = false;
      _nw4[(new BigNumber(7)).toNumber()] = (_108_v15)[_module.__default.safeIndex(_91_v0, new BigNumber((_108_v15).length))];
      _nw4[(new BigNumber(8)).toNumber()] = (new BigNumber(-904)).isLessThan(new BigNumber((_99_v8).cardinality()));
      _nw4[(new BigNumber(9)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(10)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(11)).toNumber()] = ((_99_v8).update(_91_v0, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(29)), ((_112_v0) => function (_113_i4) {
        return _112_v0;
      })(_91_v0))).length)))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), ((_114_v0) => function (_115_i3) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_114_v0,_114_v0)).length));
      })(_91_v0))));
      _nw4[(new BigNumber(12)).toNumber()] = true;
      _nw4[(new BigNumber(13)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(14)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(15)).toNumber()] = _92_v1;
      _nw4[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_108_v15, _108_v15);
      _nw4[(new BigNumber(17)).toNumber()] = ((_92_v1) ? (_92_v1) : (_module.__default.fm22(false, _91_v0, globalState)));
      _109_v16 = _nw4;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_109_v16).length))) {
        let _116_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_116_i0)) && ((_116_i0).isLessThan(new BigNumber((_109_v16).length))))) {
          (_109_v16)[(_116_i0)] = _92_v1;
        }
      }
      let _117_v17;
      let _nw5 = new _module.C7();
      _nw5.__ctor(_109_v16, _92_v1, _91_v0);
      _117_v17 = _nw5;
      let _118_v18;
      _118_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_101_v10).length),_117_v17);
      let _119_v19;
      _119_v19 = _dafny.MultiSet.fromElements(_118_v18);
      if (((_dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_117_v17.f29)).length),_117_v17)).update(_91_v0, _117_v17))).Difference(_119_v19)).IsSubsetOf(_dafny.MultiSet.fromElements(_118_v18, _118_v18))) {
        let _index0 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_102_v11).length));
        (_102_v11)[_index0] = _117_v17.f29;
        let _index1 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_102_v11).length));
        (_102_v11)[_index1] = _module.__default.safeModuloInt(new BigNumber(-587), _117_v17.f29);
        let _120_v20;
        _120_v20 = new _dafny.CodePoint('p'.codePointAt(0));
        (globalState).f25 = _120_v20;
        _91_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-565), (((_117_v17).f28) ? (_module.__default.fm1(globalState)) : ((_102_v11)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_102_v11).length))])))));
        (globalState).f6 = (_117_v17).f28;
        (globalState).f17 = _dafny.areEqual(_101_v10, _101_v10);
      } else {
        let _index2 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length));
        (_102_v11)[_index2] = _117_v17.f29;
        let _index3 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length));
        (_102_v11)[_index3] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_117_v17.f29, (_117_v17.f29).minus(_91_v0)));
        _91_v0 = ((_102_v11)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length))]).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_117_v17).f28,_97_v6)).length));
        let _121_v21;
        _121_v21 = new _dafny.CodePoint('u'.codePointAt(0));
        let _122_v22;
        let _nw6 = new _module.C8();
        _nw6.__ctor(_121_v21, _109_v16, (_117_v17).f28);
        _122_v22 = _nw6;
        let _123_v23;
        _123_v23 = _module.D1.create_DC6((_102_v11)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length))]);
        let _124_v24;
        _124_v24 = _dafny.MultiSet.fromElements(_123_v23, _123_v23, _123_v23, _123_v23);
        let _index4 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length));
        let _rhs3 = _91_v0;
        let _rhs4 = _122_v22;
        let _rhs5 = (_117_v17.f29).isLessThanOrEqualTo(_module.__default.fm4(_124_v24, globalState));
        let _lhs2 = _102_v11;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length));
        let _lhs4 = globalState;
        _lhs2[_lhs3] = _rhs3;
        _122_v22 = _rhs4;
        _lhs4.f17 = _rhs5;
        let _125_v25;
        let _nw7 = new _module.C0();
        _nw7.__ctor(_101_v10, (_102_v11)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length))], _109_v16, (_117_v17).f28);
        _125_v25 = _nw7;
        _125_v25 = _125_v25;
        let _126_v26;
        _126_v26 = _dafny.Set.fromElements((_122_v22).f32, _module.__default.fm42(new BigNumber((_108_v15).length), globalState));
        (globalState).f4 = ((_102_v11)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_102_v11).length))]).multipliedBy(_module.__default.safeModuloInt(_117_v17.f29, new BigNumber((_126_v26).length)));
      }
      (globalState).f6 = (_module.D5.create_DC13(true, _102_v11, false)).dtor_cf23;
      (globalState).f4 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), ((_127_v17) => function (_128_i5) {
        return _dafny.Set.fromElements(_127_v17.f29, _127_v17.f29);
      })(_117_v17))).length);
      let _129_v27;
      _129_v27 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_117_v17.f29),(_117_v17).f27);
      r0 = _129_v27;
      r1 = new BigNumber(820);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _130_v0;
      let _init1 = function (_131_i0) {
        return false;
      };
      let _nw8 = Array((new BigNumber(27)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
        _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _130_v0 = _nw8;
      let _132_v1;
      _132_v1 = new BigNumber(900);
      let _133_v2;
      _133_v2 = false;
      let _134_v3;
      _134_v3 = _dafny.MultiSet.fromElements(_133_v2);
      let _135_v4;
      _135_v4 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_132_v1),_134_v3);
      let _136_v5;
      let _nw9 = Array((new BigNumber(18)).toNumber()).fill([]);
      _136_v5 = _nw9;
      let _137_v6;
      _137_v6 = _dafny.Map.Empty.slice().updateUnsafe(_134_v3,_136_v5);
      let _138_v7;
      let _nw10 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _138_v7 = _nw10;
      let _139_v8;
      let _init2 = ((_140_v1, _141_v2) => function (_142_i1) {
        return _module.__default.safeModuloInt(_142_i1, new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_140_v1, _140_v1)).cardinality()),_141_v2))).length));
      })(_132_v1, _133_v2);
      let _nw11 = Array((new BigNumber(3)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
        _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _139_v8 = _nw11;
      let _143_v9;
      _143_v9 = _dafny.Seq.of(_139_v8, _139_v8);
      let _144_v10;
      _144_v10 = _dafny.Map.Empty.slice().updateUnsafe(_133_v2,_136_v5);
      let _145_v11;
      _145_v11 = _dafny.Seq.of(_133_v2, _133_v2);
      let _146_v12;
      _146_v12 = _dafny.Seq.of(_132_v1);
      let _147_v13;
      _147_v13 = _dafny.Seq.of(_136_v5, (((_144_v10).contains((_145_v11)[_module.__default.safeIndex(new BigNumber((_146_v12).length), new BigNumber((_145_v11).length))])) ? ((_144_v10).get((_145_v11)[_module.__default.safeIndex(new BigNumber((_146_v12).length), new BigNumber((_145_v11).length))])) : (_136_v5)));
      let _148_v14;
      _148_v14 = _dafny.Seq.UnicodeFromString("prvhys");
      let _149_v15;
      _149_v15 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("y"), _148_v14);
      let _150_v16;
      _150_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(76),_139_v8);
      let _151_globalState;
      let _nw12 = new _module.GlobalState();
      _nw12.__ctor(false, true, new BigNumber(888), true, new BigNumber(277), new BigNumber(131), true, new BigNumber(306), _130_v0, _135_v4, _dafny.MultiSet.fromElements(_132_v1), (((_137_v6).contains(_134_v3)) ? ((_137_v6).get(_134_v3)) : (_136_v5)), new BigNumber(321), new BigNumber(529), _138_v7, false, _dafny.Seq.Concat(_143_v9, _143_v9), false, (_147_v13)[_module.__default.safeIndex(_132_v1, new BigNumber((_147_v13).length))], _dafny.Seq.Concat(_149_v15, _149_v15), new BigNumber(778), new BigNumber(481), _150_v16, new BigNumber(726), new BigNumber(300), new _dafny.CodePoint('u'.codePointAt(0)), new BigNumber(792));
      _151_globalState = _nw12;
      let _152_v17;
      _152_v17 = _dafny.Map.Empty.slice().updateUnsafe(_133_v2,_132_v1);
      let _153_v18;
      _153_v18 = _dafny.MultiSet.fromElements(_152_v17);
      if ((_153_v18).IsSubsetOf(_dafny.MultiSet.fromElements(_152_v17))) {
        let _154_v19;
        let _155_v20;
        let _out0;
        let _out1;
        let _outcollector0 = _module.__default.m0(_151_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _154_v19 = _out0;
        _155_v20 = _out1;
        let _156_v21;
        _156_v21 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((((_134_v3).contains(_133_v2)) ? ((_134_v3).get(_133_v2)) : (_155_v20))), new BigNumber((_148_v14).length));
        let _index5 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_130_v0).length));
        (_130_v0)[_index5] = _133_v2;
        let _index6 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_130_v0).length));
        (_130_v0)[_index6] = _133_v2;
        let _157_v22;
        _157_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(_133_v2) || (_133_v2),_133_v2);
        let _158_v23;
        _158_v23 = _dafny.Set.fromElements(_132_v1, _132_v1, (_dafny.ZERO).minus(_132_v1), new BigNumber((_134_v3).cardinality()));
        let _159_v24;
        _159_v24 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_155_v20,_158_v23));
        let _160_v25;
        _160_v25 = _dafny.Map.Empty.slice().updateUnsafe(_130_v0,true);
        let _index7 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_130_v0).length));
        let _index8 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_130_v0).length));
        let _rhs6 = _156_v21;
        let _rhs7 = _133_v2;
        let _rhs8 = new BigNumber((_module.__default.fm0((_159_v24)[_module.__default.safeIndex(_155_v20, new BigNumber((_159_v24).length))], (((_160_v25).contains(_130_v0)) ? ((_160_v25).get(_130_v0)) : (_133_v2)), _133_v2, _151_globalState)).length);
        let _rhs9 = (new BigNumber((_dafny.Seq.Concat(_146_v12, _146_v12)).length)).isLessThanOrEqualTo((_155_v20).minus(new BigNumber((_148_v14).length)));
        let _rhs10 = _157_v22;
        let _lhs5 = _130_v0;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_130_v0).length));
        let _lhs7 = _130_v0;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_130_v0).length));
        _156_v21 = _rhs6;
        _lhs5[_lhs6] = _rhs7;
        _132_v1 = _rhs8;
        _lhs7[_lhs8] = _rhs9;
        _157_v22 = _rhs10;
        let _161_v26;
        let _nw13 = new _module.C4();
        _nw13.__ctor();
        _161_v26 = _nw13;
        _148_v14 = _dafny.Seq.Concat(_dafny.Seq.Concat(_148_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_162_i2) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })), _148_v14);
        _152_v17 = (_152_v17).update(_133_v2, _155_v20);
      } else {
        _148_v14 = _148_v14;
        let _163_v27;
        let _nw14 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _163_v27 = _nw14;
        _163_v27 = _163_v27;
        let _164_v28;
        _164_v28 = _dafny.Map.Empty.slice().updateUnsafe(_139_v8,_148_v14);
        _164_v28 = _164_v28;
        let _rhs11 = _134_v3;
        let _rhs12 = _133_v2;
        let _lhs9 = _151_globalState;
        _134_v3 = _rhs11;
        _lhs9.f6 = _rhs12;
        let _165_v29;
        _165_v29 = _dafny.Set.fromElements((((_134_v3).contains(_133_v2)) ? ((_134_v3).get(_133_v2)) : (_132_v1)));
        _165_v29 = ((_133_v2) ? ((function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(735))) {
            let _166_v30 = _compr_26;
            if (((new BigNumber(667)).isLessThanOrEqualTo(_166_v30)) && ((_166_v30).isLessThan(new BigNumber(735)))) {
              _coll26.add(_module.__default.safeModuloInt(_166_v30, _132_v1));
            }
          }
          return _coll26;
        }()).Union(_165_v29)) : ((((_145_v11)[_module.__default.safeIndex(_132_v1, new BigNumber((_145_v11).length))]) ? (_165_v29) : (_165_v29))));
      }
      let _167_v31;
      _167_v31 = _dafny.Set.fromElements(true, false);
      let _168_v32;
      _168_v32 = _dafny.Seq.of(_167_v31);
      let _169_v33;
      _169_v33 = _dafny.Seq.of(_module.D5.create_DC12(_168_v32));
      let _170_v34;
      _170_v34 = _dafny.Seq.of(_169_v33);
      let _171_v35;
      _171_v35 = _170_v34;
      let _172_v36;
      let _nw15 = new _module.C6();
      _nw15.__ctor(new BigNumber(609), ((_133_v2) ? (_130_v0) : (_130_v0)), _dafny.Seq.contains((_171_v35), _169_v33));
      _172_v36 = _nw15;
      if (_133_v2) {
        (_151_globalState).f4 = _132_v1;
        let _173_v37;
        _173_v37 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-185)));
        let _rhs13 = (_173_v37).Intersect(_173_v37);
        let _rhs14 = _130_v0;
        _173_v37 = _rhs13;
        _130_v0 = _rhs14;
        _152_v17 = (_152_v17).update(_133_v2, _132_v1);
        (_151_globalState).f13 = _132_v1;
        (_151_globalState).f5 = _132_v1;
      } else {
        let _174_v38;
        _174_v38 = _dafny.Set.fromElements(_132_v1);
        let _175_v39;
        _175_v39 = _module.D0.create_DC2(_174_v38, _148_v14);
        let _176_v40;
        _176_v40 = _dafny.Map.Empty.slice().updateUnsafe(_132_v1,_132_v1);
        let _177_v41;
        _177_v41 = _dafny.MultiSet.fromElements(_132_v1, new BigNumber(((_175_v39).dtor_cf5).length), ((_146_v12)[_module.__default.safeIndex(_132_v1, new BigNumber((_146_v12).length))]).minus(new BigNumber((_176_v40).length)), _132_v1);
        let _index9 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length));
        (_139_v8)[_index9] = (_132_v1).plus((_dafny.ZERO).minus(_132_v1));
        let _178_v42;
        _178_v42 = _dafny.Map.Empty.slice().updateUnsafe(_132_v1,_174_v38);
        let _179_v43;
        _179_v43 = _dafny.Map.Empty.slice().updateUnsafe(_132_v1,_145_v11);
        let _pat_let_tv0 = _132_v1;
        let _pat_let_tv1 = _132_v1;
        let _pat_let_tv2 = _132_v1;
        let _pat_let_tv3 = _178_v42;
        let _pat_let_tv4 = _133_v2;
        let _pat_let_tv5 = _133_v2;
        let _pat_let_tv6 = _151_globalState;
        let _index10 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length));
        let _rhs15 = function (_pat_let0_0) {
          return function (_182_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_183_dt__update_hcf5_h1) {
                return _module.D0.create_DC2(_183_dt__update_hcf5_h1, (_182_dt__update__tmp_h1).dtor_cf6);
              }(_pat_let3_0);
            }(_module.__default.fm0(_pat_let_tv3, _pat_let_tv4, _pat_let_tv5, _pat_let_tv6));
          }(_pat_let0_0);
        }(function (_pat_let1_0) {
          return function (_180_dt__update__tmp_h0) {
            return function (_pat_let2_0) {
              return function (_181_dt__update_hcf5_h0) {
                return _module.D0.create_DC2(_181_dt__update_hcf5_h0, (_180_dt__update__tmp_h0).dtor_cf6);
              }(_pat_let2_0);
            }(_dafny.Set.fromElements(_pat_let_tv0, (_dafny.ZERO).minus(_pat_let_tv1), _pat_let_tv2));
          }(_pat_let1_0);
        }(_175_v39));
        let _rhs16 = _177_v41;
        let _rhs17 = (_dafny.ZERO).minus((_module.__default.safeModuloInt(new BigNumber(((((_179_v43).contains(_132_v1)) ? ((_179_v43).get(_132_v1)) : (_dafny.Seq.of(_133_v2)))).length), new BigNumber((_145_v11).length))).multipliedBy((new BigNumber((_167_v31).length)).plus(_132_v1)));
        let _rhs18 = _133_v2;
        let _lhs10 = _139_v8;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length));
        let _lhs12 = _151_globalState;
        _175_v39 = _rhs15;
        _177_v41 = _rhs16;
        _lhs10[_lhs11] = _rhs17;
        _lhs12.f15 = _rhs18;
        let _184_v44;
        let _nw16 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _184_v44 = _nw16;
        _184_v44 = _184_v44;
        let _index11 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length));
        (_139_v8)[_index11] = _module.__default.safeModuloInt(_132_v1, (_139_v8)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length))]);
        let _185_v45;
        _185_v45 = _module.D8.create_DC19();
        let _186_v46;
        _186_v46 = new _dafny.CodePoint('a'.codePointAt(0));
        let _187_v47;
        _187_v47 = _dafny.Set.fromElements(_186_v46, _186_v46, new _dafny.CodePoint('p'.codePointAt(0)), _186_v46, _186_v46);
        let _188_v48;
        _188_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_132_v1,_module.D8.create_DC20(_185_v45)),((true) ? (_187_v47) : (_187_v47)));
        let _189_v49;
        _189_v49 = _dafny.Map.Empty.slice().updateUnsafe(_133_v2,_167_v31);
        let _rhs19 = (((_133_v2) ? (_189_v49) : (_189_v49))).contains(_133_v2);
        let _rhs20 = _module.__default.fm54((_132_v1).isLessThan((_139_v8)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_139_v8).length))]), false, (_133_v2) || ((_145_v11)[_module.__default.safeIndex(_132_v1, new BigNumber((_145_v11).length))]), _133_v2, _151_globalState);
        let _lhs13 = _151_globalState;
        _lhs13.f17 = _rhs19;
        _188_v48 = _rhs20;
        let _190_v50;
        _190_v50 = _module.D0.create_DC3(_module.__default.fm36(_151_globalState), _module.__default.fm2(_module.__default.fm42(_132_v1, _151_globalState), new BigNumber((_167_v31).length), _151_globalState), !((_133_v2) === (_133_v2)));
        let _pat_let_tv7 = _133_v2;
        let _pat_let_tv8 = _139_v8;
        let _pat_let_tv9 = _139_v8;
        let _pat_let_tv10 = _152_v17;
        _190_v50 = function (_pat_let4_0) {
          return function (_191_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_192_dt__update_hcf9_h0) {
                return function (_pat_let6_0) {
                  return function (_193_dt__update_hcf8_h0) {
                    return _module.D0.create_DC3((_191_dt__update__tmp_h2).dtor_cf7, _193_dt__update_hcf8_h0, _192_dt__update_hcf9_h0);
                  }(_pat_let6_0);
                }(((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_pat_let_tv8).length))]).isLessThan(new BigNumber((_pat_let_tv10).length)));
              }(_pat_let5_0);
            }(_pat_let_tv7);
          }(_pat_let4_0);
        }(_190_v50);
      }
      (_151_globalState).f6 = _133_v2;
      (_151_globalState).f6 = _133_v2;
      let _194_v51;
      _194_v51 = new _dafny.CodePoint('e'.codePointAt(0));
      (_151_globalState).f25 = _194_v51;
      let _195_v52;
      _195_v52 = _module.D8.create_DC18(_132_v1);
      let _196_v53;
      _196_v53 = _module.D8.create_DC20(_195_v52);
      let _197_v54;
      _197_v54 = _module.D8.create_DC20(_195_v52);
      let _198_v55;
      _198_v55 = _module.D8.create_DC20(_197_v54);
      _198_v55 = _module.D8.create_DC20(_196_v53);
      let _199_v56;
      _199_v56 = _module.D0.create_DC3(_132_v1, _133_v2, _133_v2);
      let _200_v57;
      _200_v57 = _dafny.Map.Empty.slice().updateUnsafe(_199_v56,new BigNumber((_148_v14).length));
      let _201_v58;
      let _nw17 = new _module.C3();
      _nw17.__ctor(new BigNumber((_148_v14).length), _200_v57, _132_v1, _130_v0, !(_132_v1).isEqualTo(_132_v1));
      _201_v58 = _nw17;
      let _202_v59;
      _202_v59 = _module.D10.create_DC25(_dafny.MultiSet.fromElements(_133_v2));
      let _203_v60;
      _203_v60 = _dafny.Set.fromElements((_202_v59).dtor_cf40);
      let _204_i3;
      _204_i3 = _dafny.ZERO;
      L0: {
        while ((_203_v60).contains(_134_v3)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_204_i3)) {
              break L0;
            }
            _204_i3 = (_204_i3).plus(_dafny.ONE);
            (_151_globalState).f25 = _194_v51;
            (_172_v36).m6(_133_v2, (_133_v2) && (_133_v2), _module.__default.safeDivisionInt((_dafny.ZERO).minus((_201_v58).f34), new BigNumber((_145_v11).length)), _151_globalState);
            let _205_v61;
            _205_v61 = _dafny.MultiSet.fromElements(_132_v1);
            let _206_v62;
            let _nw18 = new _module.C9();
            _nw18.__ctor(_133_v2, _205_v61, _132_v1, _130_v0, _133_v2);
            _206_v62 = _nw18;
            let _207_v63;
            _207_v63 = _dafny.Map.Empty.slice().updateUnsafe(_132_v1,_206_v62);
            let _208_v64;
            _208_v64 = _dafny.Seq.of(_152_v17);
            _207_v63 = (_207_v63).update(new BigNumber((_208_v64).length), _206_v62);
            let _209_v65;
            _209_v65 = _dafny.Map.Empty.slice().updateUnsafe(_132_v1,_133_v2);
            let _210_v67;
            _210_v67 = _dafny.Set.fromElements(_209_v65, function () {
              let _coll27 = new _dafny.Map();
              for (const _compr_27 of ((_206_v62).f31).Elements) {
                let _211_v66 = _compr_27;
                if (((_206_v62).f31).contains(_211_v66)) {
                  _coll27.push([(_211_v66).plus((_201_v58).f34),!(_133_v2)]);
                }
              }
              return _coll27;
            }());
            (_151_globalState).f6 = ((_210_v67).Intersect(_210_v67)).IsSubsetOf((_210_v67).Difference(_210_v67));
          }
        }
      }
      let _212_v68;
      _212_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(498)), ((_213_v51) => function (_214_i4) {
        return _213_v51;
      })(_194_v51))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_201_v58).f34,_133_v2)).length));
      let _215_v69;
      _215_v69 = _dafny.Seq.of(_172_v36, _172_v36);
      let _216_v70;
      let _217_v71;
      let _218_v72;
      let _219_v73;
      let _out2;
      let _out3;
      let _out4;
      let _out5;
      let _outcollector1 = (_201_v58).m10(((_201_v58).fm14((_201_v58).f34, _203_v60, _151_globalState)).isLessThan((((_212_v68).contains(_132_v1)) ? ((_212_v68).get(_132_v1)) : (new BigNumber((_215_v69).length)))), _151_globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _out4 = _outcollector1[2];
      _out5 = _outcollector1[3];
      _216_v70 = _out2;
      _217_v71 = _out3;
      _218_v72 = _out4;
      _219_v73 = _out5;
      let _220_v74;
      _220_v74 = _dafny.Set.fromElements(_130_v0, _130_v0, _130_v0);
      let _221_v75;
      _221_v75 = _dafny.Map.Empty.slice().updateUnsafe(_220_v74,(_167_v31).IsProperSubsetOf(_167_v31));
      let _222_v76;
      _222_v76 = _dafny.Set.fromElements(new BigNumber(-854));
      let _223_v77;
      _223_v77 = _dafny.Map.Empty.slice().updateUnsafe(_217_v71,_222_v76);
      let _224_v78;
      _224_v78 = _dafny.Map.Empty.slice().updateUnsafe(_219_v73,_222_v76);
      _221_v75 = (_221_v75).update((_220_v74).Intersect(_220_v74), (_module.__default.fm0(_224_v78, _216_v70, true, _151_globalState)).IsSubsetOf((((_223_v77).contains(_133_v2)) ? ((_223_v77).get(_133_v2)) : (_222_v76))));
      if ((_145_v11)[_module.__default.safeIndex(_219_v73, new BigNumber((_145_v11).length))]) {
        _132_v1 = _module.__default.safeDivisionInt((_201_v58).f34, (_dafny.ZERO).minus((_201_v58).f34));
        let _index12 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
        (_139_v8)[_index12] = _219_v73;
        let _index13 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
        (_139_v8)[_index13] = _module.__default.fm36(_151_globalState);
        if (false) {
          let _225_v79;
          _225_v79 = _dafny.Map.Empty.slice().updateUnsafe((_201_v58).f34,_216_v70);
          _225_v79 = ((!(new BigNumber(564)).isEqualTo(_219_v73)) ? ((_225_v79).Merge(_module.__default.fm30((((_225_v79).contains((_139_v8)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length))])) ? ((_225_v79).get((_139_v8)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length))])) : (_216_v70)), _148_v14, _133_v2, _151_globalState))) : (_225_v79));
          let _226_v80;
          _226_v80 = _dafny.Map.Empty.slice().updateUnsafe(_133_v2,_130_v0);
          let _227_v81;
          let _nw19 = new _module.C7();
          _nw19.__ctor(_130_v0, _217_v71, _132_v1);
          _227_v81 = _nw19;
          let _228_v82;
          _228_v82 = _dafny.Map.Empty.slice().updateUnsafe(_227_v81,(_227_v81).f27);
          let _229_v83;
          _229_v83 = _dafny.Seq.of((((_226_v80).contains(!(_216_v70))) ? ((_226_v80).get(!(_216_v70))) : ((((_228_v82).contains(_227_v81)) ? ((_228_v82).get(_227_v81)) : (_130_v0)))), _130_v0, (_227_v81).f27, (_227_v81).f27);
          _130_v0 = (_229_v83)[_module.__default.safeIndex(_219_v73, new BigNumber((_229_v83).length))];
          (_151_globalState).f15 = _216_v70;
          let _230_v84;
          let _nw20 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _230_v84 = _nw20;
          let _231_v85;
          _231_v85 = _dafny.Seq.of(_dafny.Seq.of(_230_v84));
          let _232_v86;
          let _nw21 = new _module.C5();
          _nw21.__ctor(_dafny.Seq.Concat((_231_v85)[_module.__default.safeIndex((_201_v58).fm14(new BigNumber(83), _203_v60, _151_globalState), new BigNumber((_231_v85).length))], _dafny.Seq.update(_143_v9, _module.__default.safeIndex((_201_v58).f34, new BigNumber((_143_v9).length)), _230_v84)), ((_216_v70) ? ((_201_v58).f34) : (new BigNumber(-87))), _130_v0, _216_v70);
          _232_v86 = _nw21;
          _219_v73 = _219_v73;
        } else {
          _218_v72 = _218_v72;
          _133_v2 = !(_133_v2);
          let _233_v87;
          let _nw22 = new _module.C4();
          _nw22.__ctor();
          _233_v87 = _nw22;
          let _234_v88;
          _234_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_146_v12, _146_v12),_233_v87);
          _234_v88 = (_234_v88).update(_module.__default.fm55(_216_v70, _151_globalState), _233_v87);
          (_151_globalState).f21 = (_201_v58).f34;
          let _index14 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_130_v0).length));
          (_130_v0)[_index14] = _217_v71;
          let _235_v89;
          let _nw23 = Array((new BigNumber(12)).toNumber()).fill([]);
          _235_v89 = _nw23;
          let _index15 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_235_v89).length));
          (_235_v89)[_index15] = _130_v0;
          let _236_v90;
          _236_v90 = _module.D1.create_DC6(new BigNumber(318));
          let _237_v91;
          _237_v91 = _dafny.MultiSet.fromElements(_module.D1.create_DC6(new BigNumber(((_212_v68).update((_201_v58).f34, new BigNumber(117))).length)), _236_v90, _module.D1.create_DC6(_132_v1));
          let _index16 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_130_v0).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
          let _index18 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_235_v89).length));
          let _rhs21 = ((_201_v58).f34).isEqualTo(new BigNumber(956));
          let _rhs22 = _146_v12;
          let _rhs23 = _module.__default.fm4(_237_v91, _151_globalState);
          let _rhs24 = (_201_v58).f34;
          let _rhs25 = _130_v0;
          let _lhs14 = _130_v0;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_130_v0).length));
          let _lhs16 = _139_v8;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
          let _lhs18 = _139_v8;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_139_v8).length));
          let _lhs20 = _235_v89;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_235_v89).length));
          _lhs14[_lhs15] = _rhs21;
          _146_v12 = _rhs22;
          _lhs16[_lhs17] = _rhs23;
          _lhs18[_lhs19] = _rhs24;
          _lhs20[_lhs21] = _rhs25;
        }
        _132_v1 = _132_v1;
        let _238_v92;
        _238_v92 = _dafny.MultiSet.fromElements(_132_v1);
        let _239_v93;
        _239_v93 = _dafny.MultiSet.fromElements(_238_v92, _238_v92);
        let _240_v94;
        let _nw24 = new _module.C2();
        _nw24.__ctor(_module.__default.fm42(new BigNumber(992), _151_globalState), _217_v71, _130_v0, _216_v70);
        _240_v94 = _nw24;
        let _241_v95;
        _241_v95 = _dafny.Seq.of(_240_v94, _240_v94);
        let _242_v96;
        _242_v96 = _dafny.Map.Empty.slice().updateUnsafe((_239_v93).update(_238_v92, _module.__default.abs(new BigNumber((_241_v95).length))),_218_v72);
        _242_v96 = (_242_v96).update(_239_v93, _194_v51);
      } else {
        let _243_v97;
        _243_v97 = _dafny.MultiSet.fromElements(_148_v14);
        (_151_globalState).f17 = (_243_v97).IsSubsetOf(_243_v97);
        _194_v51 = ((_133_v2) ? (_218_v72) : (_194_v51));
        let _244_v98;
        let _nw25 = Array((new BigNumber(8)).toNumber()).fill([]);
        _244_v98 = _nw25;
        let _245_v99;
        _245_v99 = _dafny.Map.Empty.slice().updateUnsafe(_217_v71,_244_v98);
        let _246_v100;
        _246_v100 = _module.D5.create_DC13(false, _139_v8, _133_v2);
        _244_v98 = (((_245_v99).contains((_246_v100).dtor_cf21)) ? ((_245_v99).get((_246_v100).dtor_cf21)) : (_244_v98));
        if (_module.__default.fm22(!((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-179)), (_201_v58).f34)).IsSubsetOf(_222_v76)), new BigNumber(((_149_v15)[_module.__default.safeIndex(_219_v73, new BigNumber((_149_v15).length))]).length), _151_globalState)) {
          let _247_v101;
          let _nw26 = new _module.C7();
          _nw26.__ctor(_130_v0, _module.__default.fm2(_218_v72, (_201_v58).f34, _151_globalState), _219_v73);
          _247_v101 = _nw26;
          let _248_v102;
          let _nw27 = new _module.C9();
          _nw27.__ctor(_217_v71, _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_219_v73)), _132_v1, _130_v0, _133_v2);
          _248_v102 = _nw27;
          let _249_v103;
          _249_v103 = _dafny.Map.Empty.slice().updateUnsafe(_248_v102,_247_v101);
          let _rhs26 = (((_249_v103).contains(_248_v102)) ? ((_249_v103).get(_248_v102)) : (_247_v101));
          let _rhs27 = _172_v36;
          _247_v101 = _rhs26;
          _172_v36 = _rhs27;
          let _250_v104;
          _250_v104 = _dafny.Map.Empty.slice().updateUnsafe(_216_v70,!(_217_v71));
          _250_v104 = (_250_v104).update((_247_v101).f28, _216_v70);
          let _251_v105;
          _251_v105 = _module.D8.create_DC19();
          let _252_v106;
          _252_v106 = _dafny.Map.Empty.slice().updateUnsafe(_251_v105,_module.__default.fm40(_151_globalState));
          let _253_v107;
          _253_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(860),_252_v106);
          let _rhs28 = _248_v102.f29;
          let _rhs29 = !(_217_v71);
          let _rhs30 = ((_252_v106).Merge(_252_v106)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_251_v105,_133_v2)).Merge(((((_253_v107).contains(_248_v102.f29)) ? ((_253_v107).get(_248_v102.f29)) : (_252_v106))).update(_251_v105, _216_v70)));
          let _rhs31 = _218_v72;
          let _lhs22 = _151_globalState;
          let _lhs23 = _151_globalState;
          let _lhs24 = _151_globalState;
          _lhs22.f4 = _rhs28;
          _lhs23.f15 = _rhs29;
          _252_v106 = _rhs30;
          _lhs24.f25 = _rhs31;
          _149_v15 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), ((_254_v14) => function (_255_i5) {
            return _254_v14;
          })(_148_v14)), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(808)), function (_256_i6) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          }), _148_v14, _dafny.Seq.UnicodeFromString("fqiqust"), _dafny.Seq.UnicodeFromString("uapyvflv")));
          (_151_globalState).f15 = _217_v71;
        } else {
          let _257_v108;
          let _nw28 = new _module.C6();
          _nw28.__ctor(new BigNumber((_146_v12).length), _130_v0, _217_v71);
          _257_v108 = _nw28;
          _148_v14 = _dafny.Seq.Concat(_dafny.Seq.Concat(_148_v14, _148_v14), _148_v14);
          (_151_globalState).f6 = _133_v2;
          let _258_v109;
          _258_v109 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_216_v70,(_201_v58).f34), _152_v17, _152_v17, _dafny.Map.Empty.slice().updateUnsafe(_216_v70,_132_v1), _module.__default.fm50(_133_v2, new BigNumber((_module.__default.fm31(_151_globalState)).length), _151_globalState));
          let _259_v110;
          _259_v110 = _dafny.Map.Empty.slice().updateUnsafe(_258_v109,_139_v8);
          _259_v110 = (_259_v110).update(_258_v109, _139_v8);
          let _260_v111;
          let _nw29 = new _module.C4();
          _nw29.__ctor();
          _260_v111 = _nw29;
        }
        _132_v1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_148_v14, (_149_v15)[_module.__default.safeIndex(_132_v1, new BigNumber((_149_v15).length))]), _dafny.Seq.Concat(_148_v14, _module.__default.fm32(new BigNumber((_243_v97).cardinality()), _216_v70, _151_globalState)))).length);
      }
      let _261_v112;
      _261_v112 = _dafny.Map.Empty.slice().updateUnsafe(_216_v70,_133_v2);
      _261_v112 = (_261_v112).update(!(!(_217_v71)) || (_216_v70), _217_v71);
      let _262_v113;
      let _nw30 = new _module.C0();
      _nw30.__ctor(_148_v14, _132_v1, _130_v0, _216_v70);
      _262_v113 = _nw30;
      (_151_globalState).f21 = (_module.__default.safeDivisionInt(_219_v73, _132_v1)).plus(_132_v1);
      if (!((_133_v2) || (_217_v71))) {
        let _263_v115;
        _263_v115 = _dafny.MultiSet.fromElements(_132_v1);
        let _264_v116;
        _264_v116 = _dafny.Seq.of(function () {
          let _coll28 = new _dafny.Map();
          for (const _compr_28 of (_263_v115).Elements) {
            let _265_v114 = _compr_28;
            if ((_263_v115).contains(_265_v114)) {
              _coll28.push([(_265_v114).minus((_201_v58).f34),(_145_v11)[_module.__default.safeIndex((_dafny.ZERO).minus((_201_v58).f34), new BigNumber((_145_v11).length))]]);
            }
          }
          return _coll28;
        }());
        _264_v116 = _264_v116;
        _212_v68 = (_212_v68).update(_219_v73, new BigNumber(((_262_v113).f36).length));
        let _index20 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_130_v0).length));
        (_130_v0)[_index20] = !(_133_v2);
        let _index21 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_130_v0).length));
        let _rhs32 = (_149_v15)[_module.__default.safeIndex((_201_v58).f34, new BigNumber((_149_v15).length))];
        let _rhs33 = _133_v2;
        let _rhs34 = ((new BigNumber((_263_v115).cardinality())).multipliedBy(new BigNumber((_145_v11).length))).isLessThanOrEqualTo(_132_v1);
        let _rhs35 = _dafny.Seq.contains((_262_v113).f36, new _dafny.CodePoint('e'.codePointAt(0)));
        let _lhs25 = _130_v0;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_130_v0).length));
        let _lhs27 = _151_globalState;
        _148_v14 = _rhs32;
        _lhs25[_lhs26] = _rhs33;
        _217_v71 = _rhs34;
        _lhs27.f15 = _rhs35;
        let _index22 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_139_v8).length));
        (_139_v8)[_index22] = _219_v73;
        let _266_v117;
        _266_v117 = _module.D1.create_DC6((_201_v58).f34);
        let _267_v118;
        _267_v118 = _dafny.MultiSet.fromElements(_module.D1.create_DC6(new BigNumber((_263_v115).cardinality())), _266_v117, _266_v117, _266_v117);
        let _index23 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_139_v8).length));
        (_139_v8)[_index23] = (_module.__default.fm4(_267_v118, _151_globalState)).plus(_module.__default.safeDivisionInt(new BigNumber((_145_v11).length), _219_v73));
        let _268_v119;
        let _nw31 = Array((new BigNumber(11)).toNumber()).fill(_module.D9.Default());
        _268_v119 = _nw31;
        let _index24 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_268_v119).length));
        (_268_v119)[_index24] = (((_130_v0)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_130_v0).length))]) ? (_module.D9.create_DC21(_138_v7)) : (_module.D9.create_DC21(_138_v7)));
        let _pat_let_tv11 = _138_v7;
        let _index25 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_268_v119).length));
        let _rhs36 = _133_v2;
        let _rhs37 = function (_pat_let7_0) {
          return function (_269_dt__update__tmp_h3) {
            return function (_pat_let8_0) {
              return function (_270_dt__update_hcf33_h0) {
                return _module.D9.create_DC21(_270_dt__update_hcf33_h0);
              }(_pat_let8_0);
            }(_pat_let_tv11);
          }(_pat_let7_0);
        }(_module.D9.create_DC21(_138_v7));
        let _rhs38 = _217_v71;
        let _lhs28 = _151_globalState;
        let _lhs29 = _268_v119;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_268_v119).length));
        _lhs28.f6 = _rhs36;
        _lhs29[_lhs30] = _rhs37;
        _133_v2 = _rhs38;
      } else {
        let _index26 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length));
        (_130_v0)[_index26] = _217_v71;
        let _index27 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length));
        (_130_v0)[_index27] = _217_v71;
        let _271_v120;
        _271_v120 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(892), (_201_v58).f34),_216_v70);
        _271_v120 = (_271_v120).Merge(_271_v120);
        (_151_globalState).f5 = (_201_v58).f34;
        let _272_v121;
        _272_v121 = _module.D0.create_DC2((_222_v76).Intersect(_dafny.Set.fromElements(_132_v1, (_201_v58).f34)), _148_v14);
        _272_v121 = _272_v121;
        let _273_v122;
        let _nw32 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _273_v122 = _nw32;
        let _274_v123;
        _274_v123 = _dafny.Map.Empty.slice().updateUnsafe(_167_v31,_133_v2);
        let _275_v125;
        _275_v125 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_217_v71), _167_v31);
        let _276_v129;
        _276_v129 = _dafny.Map.Empty.slice().updateUnsafe(_167_v31,(_201_v58).f34);
        let _277_v130;
        let _nw33 = Array((new BigNumber(21)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _274_v123;
        _nw33[(_dafny.ONE).toNumber()] = _274_v123;
        _nw33[(new BigNumber(2)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(3)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(4)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(5)).toNumber()] = (function () {
          let _coll29 = new _dafny.Map();
          for (const _compr_29 of (_275_v125).Elements) {
            let _278_v124 = _compr_29;
            if ((_275_v125).contains(_278_v124)) {
              _coll29.push([_278_v124,(_130_v0)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length))]]);
            }
          }
          return _coll29;
        }()).Merge(_274_v123);
        _nw33[(new BigNumber(6)).toNumber()] = _module.__default.fm56(_217_v71, _218_v72, _151_globalState);
        _nw33[(new BigNumber(7)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(8)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(9)).toNumber()] = (_274_v123).Merge(_274_v123);
        _nw33[(new BigNumber(10)).toNumber()] = function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_275_v125).Elements) {
            let _279_v126 = _compr_30;
            if ((_275_v125).contains(_279_v126)) {
              _coll30.push([_279_v126,_133_v2]);
            }
          }
          return _coll30;
        }();
        _nw33[(new BigNumber(11)).toNumber()] = function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of (_274_v123).Keys.Elements) {
            let _280_v127 = _compr_31;
            if ((_274_v123).contains(_280_v127)) {
              _coll31.push([_280_v127,(_145_v11)[_module.__default.safeIndex(_219_v73, new BigNumber((_145_v11).length))]]);
            }
          }
          return _coll31;
        }();
        _nw33[(new BigNumber(12)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(13)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(14)).toNumber()] = (_274_v123).Merge(function () {
          let _coll32 = new _dafny.Map();
          for (const _compr_32 of (_276_v129).Keys.Elements) {
            let _281_v128 = _compr_32;
            if ((_276_v129).contains(_281_v128)) {
              _coll32.push([_281_v128,(_130_v0)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length))]]);
            }
          }
          return _coll32;
        }());
        _nw33[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_167_v31,_133_v2)).update(_dafny.Set.fromElements((_130_v0)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length))]), true);
        _nw33[(new BigNumber(16)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(17)).toNumber()] = (_module.__default.fm56(_216_v70, ((_262_v113).f36)[_module.__default.safeIndex((_199_v56).dtor_cf7, new BigNumber(((_262_v113).f36).length))], _151_globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_130_v0)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length))], _216_v70, _216_v70),true));
        _nw33[(new BigNumber(18)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(19)).toNumber()] = _274_v123;
        _nw33[(new BigNumber(20)).toNumber()] = _274_v123;
        _277_v130 = _nw33;
        let _index28 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_277_v130).length));
        (_277_v130)[_index28] = _274_v123;
        let _282_v131;
        _282_v131 = _dafny.MultiSet.fromElements(_132_v1, new BigNumber((_dafny.Seq.Concat(_148_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), ((_283_v51) => function (_284_i7) {
          return _283_v51;
        })(_194_v51)))).length));
        let _285_v132;
        let _nw34 = Array((new BigNumber(16)).toNumber()).fill([]);
        _285_v132 = _nw34;
        let _286_v133;
        _286_v133 = _dafny.Map.Empty.slice().updateUnsafe(_285_v132,_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-663)), (_201_v58).f34));
        let _287_v134;
        _287_v134 = _module.D9.create_DC22(_139_v8, _133_v2, _285_v132, (_262_v113).f36);
        let _288_v135;
        _288_v135 = _dafny.Seq.of(_274_v123, _274_v123);
        let _index29 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_277_v130).length));
        let _rhs39 = ((_217_v71) ? (_273_v122) : (_273_v122));
        let _rhs40 = (((_286_v133).contains((_287_v134).dtor_cf36)) ? ((_286_v133).get((_287_v134).dtor_cf36)) : (_219_v73));
        let _rhs41 = _132_v1;
        let _rhs42 = ((_274_v123).Merge((_288_v135)[_module.__default.safeIndex(_132_v1, new BigNumber((_288_v135).length))])).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_216_v70),(_130_v0)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_130_v0).length))]));
        let _rhs43 = (_282_v131).Difference(_282_v131);
        let _lhs31 = _151_globalState;
        let _lhs32 = _277_v130;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_277_v130).length));
        _273_v122 = _rhs39;
        _lhs31.f13 = _rhs40;
        _132_v1 = _rhs41;
        _lhs32[_lhs33] = _rhs42;
        _282_v131 = _rhs43;
      }
      process.stdout.write(_dafny.toString((_130_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v0)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_132_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v3).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-900),_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_137_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_143_v9).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_144_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_145_v11, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_146_v12, _dafny.Seq.of(new BigNumber(900)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_147_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_148_v14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_149_v15, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("prvhys")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_150_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f8)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState.f9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-900),_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f10).equals(_dafny.MultiSet.fromElements(new BigNumber(900)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_151_globalState).f16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_151_globalState).f19, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("prvhys"), _dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("prvhys")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_151_globalState).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(820)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v18).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(900))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v31).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_168_v32, _dafny.Seq.of(_dafny.Set.fromElements(true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_169_v33, _dafny.Seq.of(_module.D5.create_DC12(_dafny.Seq.of(_dafny.Set.fromElements(true, false)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_170_v34, _dafny.Seq.of(_dafny.Seq.of(_module.D5.create_DC12(_dafny.Seq.of(_dafny.Set.fromElements(true, false))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_171_v35), _dafny.Seq.of(_dafny.Seq.of(_module.D5.create_DC12(_dafny.Seq.of(_dafny.Set.fromElements(true, false))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_194_v51));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v52).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_196_v53).dtor_cf32).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_197_v54).dtor_cf32).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_198_v55).dtor_cf32).dtor_cf32).dtor_cf31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_v56).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_v56).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_v56).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_200_v57).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(_dafny.ONE, false, false),new BigNumber(763)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_201_v58).f34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_201_v58).f35).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(_dafny.ONE, false, false),new BigNumber(763)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_202_v59).dtor_cf40).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_v60).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_204_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v68).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(498),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_215_v69).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_216_v70));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_217_v71));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_218_v72));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_219_v73));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_220_v74).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_221_v75).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v76).equals(_dafny.Set.fromElements(new BigNumber(-854)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v77).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(new BigNumber(-854))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v78).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(763),_dafny.Set.fromElements(new BigNumber(-854))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_261_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_262_v113).f36).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC1() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC2(cf5, cf6) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf7, cf8, cf9) {
      let $dt = new D0(3);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf10) {
      let $dt = new D0(4);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get is_DC4() { return this.$tag === 4; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + this.cf6.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 4) {
        return "D0.DC4" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.Map.Empty, false, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC6(cf12) {
      let $dt = new D1(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf13, cf14) {
      let $dt = new D1(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D1(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf13) + ", " + this.cf14.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC6(_dafny.ZERO);
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
    static create_DC8(cf15) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf17, cf18) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf16) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17 && this.cf18 === other.cf18;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, false);
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
    static create_DC11(cf19) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf21, cf22, cf23) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC12(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(false, [], false);
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
    static create_DC14(cf24) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24;
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf26, cf27, cf28, cf29) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16(false, _dafny.Map.Empty, new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC18(cf31) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC19() {
      let $dt = new D8(1);
      return $dt;
    }
    static create_DC17(cf30) {
      let $dt = new D8(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf32) {
      let $dt = new D8(3);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19";
      } else if (this.$tag === 2) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC18(_dafny.ZERO);
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
    static create_DC22(cf34, cf35, cf36, cf37) {
      let $dt = new D9(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC23(cf38) {
      let $dt = new D9(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D9(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC24(cf39) {
      let $dt = new D9(3);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + this.cf37.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && this.cf35 === other.cf35 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf33 === other.cf33;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC22([], false, [], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC26(cf41) {
      let $dt = new D10(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC25(cf40) {
      let $dt = new D10(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(false);
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
    static create_DC28() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC27(cf42) {
      let $dt = new D11(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf42) + ")";
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
        return other.$tag === 1 && this.cf42 === other.cf42;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28();
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
    static create_DC29(cf43) {
      let $dt = new D12(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43);
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf44) {
      let $dt = new D13(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32() {
      let $dt = new D14(0);
      return $dt;
    }
    static create_DC33(cf46, cf47) {
      let $dt = new D14(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC31(cf45) {
      let $dt = new D14(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC32";
      } else if (this.$tag === 1) {
        return "D14.DC33" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf45) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC32();
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
    static create_DC35(cf49) {
      let $dt = new D15(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC34(cf48) {
      let $dt = new D15(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC35(_dafny.ZERO);
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
    static create_DC37(cf51, cf52) {
      let $dt = new D16(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC38(cf53, cf54, cf55) {
      let $dt = new D16(1);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC36(cf50) {
      let $dt = new D16(2);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC37" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC38" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC36" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf51 === other.cf51 && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf53 === other.cf53 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC37(false, false);
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
    static create_DC39(cf56) {
      let $dt = new D17(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC39" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf58) {
      let $dt = new D18(0);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC40(cf57) {
      let $dt = new D18(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC42(cf59) {
      let $dt = new D18(2);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC41" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC40" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC42" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC41(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
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
      this.f4 = _dafny.ZERO;
      this.f5 = _dafny.ZERO;
      this.f6 = false;
      this.f9 = _dafny.Map.Empty;
      this.f13 = _dafny.ZERO;
      this.f15 = false;
      this.f17 = false;
      this.f18 = [];
      this.f21 = _dafny.ZERO;
      this.f25 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f0 = false;
      this._f1 = false;
      this._f2 = _dafny.ZERO;
      this._f3 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = [];
      this._f10 = _dafny.MultiSet.Empty;
      this._f11 = [];
      this._f12 = _dafny.ZERO;
      this._f14 = [];
      this._f16 = _dafny.Seq.of();
      this._f19 = _dafny.Seq.of();
      this._f20 = _dafny.ZERO;
      this._f22 = _dafny.Map.Empty;
      this._f23 = _dafny.ZERO;
      this._f24 = _dafny.ZERO;
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this).f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this).f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f10() {
      let _this = this;
      return _this._f10;
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
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
      this._f36 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f36, f29, f27, f28) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f29 = f29;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm15(globalState) {
      let _this = this;
      return (_this.f29).plus((_dafny.ZERO).minus(_this.f29));
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f27 = [];
      this._f28 = false;
      this._f39 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f39, f27, f28) {
      let _this = this;
      (_this)._f39 = f39;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm34(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("rhictu");
    };
    fm35(p0, p1, globalState) {
      let _this = this;
      return _module.D3.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(801),_dafny.Set.fromElements(new BigNumber(527), new BigNumber(65), new BigNumber(543))));
    };
    m14(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = _dafny.Map.Empty;
      let _289_v0;
      _289_v0 = _dafny.Seq.UnicodeFromString("usplmirjv");
      (globalState).f13 = new BigNumber((_289_v0).length);
      let _290_v1;
      _290_v1 = new BigNumber(511);
      let _291_v2;
      _291_v2 = _dafny.MultiSet.fromElements(_290_v1, _290_v1, _290_v1);
      let _292_v3;
      _292_v3 = _dafny.Seq.of(_290_v1);
      if (((_dafny.MultiSet.FromArray(_292_v3)).Difference(_291_v2)).IsSubsetOf((_291_v2).update(_290_v1, _module.__default.abs(_290_v1)))) {
        let _index30 = _module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index30] = (_this).f39;
        let _index31 = _module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index31] = p0;
        let _293_v4;
        _293_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_291_v2).cardinality())), new BigNumber((_289_v0).length)),_dafny.Seq.UnicodeFromString("jj"));
        _293_v4 = _293_v4;
        let _294_v5;
        _294_v5 = _dafny.Seq.of(_289_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), function (_295_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
        let _296_v6;
        let _nw35 = Array((new BigNumber(4)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_294_v5)[_module.__default.safeIndex(_290_v1, new BigNumber((_294_v5).length))], _289_v0);
        _nw35[(_dafny.ONE).toNumber()] = _289_v0;
        _nw35[(new BigNumber(2)).toNumber()] = _289_v0;
        _nw35[(new BigNumber(3)).toNumber()] = _289_v0;
        _296_v6 = _nw35;
        let _297_v7;
        _297_v7 = _module.D9.create_DC21(_296_v6);
        _296_v6 = (_297_v7).dtor_cf33;
        let _index32 = _module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index32] = (_290_v1).isLessThanOrEqualTo(_module.__default.fm36(globalState));
        let _298_v8;
        _298_v8 = _dafny.Seq.of(p0, ((_this).f27)[_module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f27).length))]);
        (globalState).f5 = _module.__default.safeDivisionInt((new BigNumber((_298_v8).length)).minus(new BigNumber((_292_v3).length)), new BigNumber((_dafny.Seq.Concat(_289_v0, _289_v0)).length));
      } else {
        let _299_v9;
        _299_v9 = _module.D8.create_DC18(_290_v1);
        let _300_v10;
        _300_v10 = _dafny.Set.fromElements(_299_v9, _299_v9, _299_v9, _module.D8.create_DC18(_290_v1), _299_v9);
        _300_v10 = (_300_v10).Intersect((_300_v10).Difference(_300_v10));
        let _301_v11;
        _301_v11 = _dafny.Set.fromElements(false);
        (globalState).f5 = (new BigNumber(783)).multipliedBy(new BigNumber((_301_v11).length));
        (globalState).f5 = (_290_v1).multipliedBy(_290_v1);
        (globalState).f13 = (_290_v1).multipliedBy((_290_v1).plus(_290_v1));
        (globalState).f6 = p0;
      }
      (globalState).f15 = (_this).f39;
      let _302_v12;
      let _nw36 = new _module.C0();
      _nw36.__ctor(_289_v0, _290_v1, (_this).f27, p0);
      _302_v12 = _nw36;
      (globalState).f15 = !(!(p0));
      let _303_v13;
      _303_v13 = _dafny.MultiSet.fromElements((_this).f39);
      let _304_v14;
      _304_v14 = _module.D10.create_DC25(_303_v13);
      let _hi0 = new BigNumber((_dafny.MultiSet.fromElements((((_303_v13).contains((_this).f39)) ? ((_303_v13).get((_this).f39)) : (_290_v1)))).cardinality());
      for (let _305_i1 = new BigNumber(((_304_v14).dtor_cf40).cardinality()); _305_i1.isLessThan(_hi0); _305_i1 = _305_i1.plus(_dafny.ONE)) {
        let _306_v16;
        let _init3 = ((_307_i1, _308_v1) => function (_309_i2) {
          return function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(-710), new BigNumber(610))) {
              let _310_v15 = _compr_33;
              if (((new BigNumber(-710)).isLessThanOrEqualTo(_310_v15)) && ((_310_v15).isLessThan(new BigNumber(610)))) {
                _coll33.push([(_310_v15).multipliedBy(new BigNumber((_dafny.Set.fromElements(_307_i1, new BigNumber(954))).length)),_dafny.Map.Empty.slice().updateUnsafe(_308_v1,new BigNumber(947))]);
              }
            }
            return _coll33;
          }();
        })(_305_i1, _290_v1);
        let _nw37 = Array((new BigNumber(15)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw37.length); _i0_3++) {
          _nw37[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _306_v16 = _nw37;
        let _311_v18;
        _311_v18 = _dafny.Map.Empty.slice().updateUnsafe(_305_i1,((function () {
          let _coll34 = new _dafny.Map();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-859), new BigNumber(455))) {
            let _312_v17 = _compr_34;
            if (((new BigNumber(-859)).isLessThanOrEqualTo(_312_v17)) && ((_312_v17).isLessThan(new BigNumber(455)))) {
              _coll34.push([_module.__default.safeModuloInt(_312_v17, _305_i1),_305_i1]);
            }
          }
          return _coll34;
        }()).update(_290_v1, _290_v1)).update((((_291_v2).contains(_305_i1)) ? ((_291_v2).get(_305_i1)) : (_290_v1)), _305_i1));
        let _index33 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_306_v16).length));
        (_306_v16)[_index33] = _311_v18;
        let _index34 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_306_v16).length));
        (_306_v16)[_index34] = _311_v18;
        (globalState).f5 = (((_this).f39) ? (new BigNumber((_303_v13).cardinality())) : (_290_v1));
        let _313_v19;
        let _init4 = ((_314_i1) => function (_315_i3) {
          return _module.__default.safeModuloInt(_315_i3, _314_i1);
        })(_305_i1);
        let _nw38 = Array((new BigNumber(21)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw38.length); _i0_4++) {
          _nw38[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _313_v19 = _nw38;
        let _index35 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_313_v19).length));
        (_313_v19)[_index35] = _290_v1;
        let _index36 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_313_v19).length));
        (_313_v19)[_index36] = (_module.__default.safeDivisionInt(new BigNumber((_292_v3).length), _290_v1)).plus(_290_v1);
        let _index37 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_313_v19).length));
        (_313_v19)[_index37] = (_313_v19)[_module.__default.safeIndex(new BigNumber(165), new BigNumber((_313_v19).length))];
      }
      r0 = _dafny.Seq.Concat((_302_v12).f36, _289_v0);
      let _316_v20;
      _316_v20 = _module.D9.create_DC23(p0);
      let _pat_let_tv12 = p0;
      let _pat_let_tv13 = p0;
      let _pat_let_tv14 = _290_v1;
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = p0;
      let _pat_let_tv17 = p0;
      let _pat_let_tv18 = p0;
      r1 = function (_source6) {
        if (_source6.is_DC22) {
          let _317___mcc_h0 = (_source6).cf34;
          let _318___mcc_h1 = (_source6).cf35;
          let _319___mcc_h2 = (_source6).cf36;
          let _320___mcc_h3 = (_source6).cf37;
          let _321_cf37 = _320___mcc_h3;
          let _322_cf36 = _319___mcc_h2;
          let _323_cf35 = _318___mcc_h1;
          let _324_cf34 = _317___mcc_h0;
          return _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv12, _pat_let_tv13), _dafny.Seq.of((_this).f28)), _module.__default.safeIndex(_pat_let_tv14, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv15, _pat_let_tv16), _dafny.Seq.of((_this).f28))).length)), (_this).f28);
        } else if (_source6.is_DC23) {
          let _325___mcc_h4 = (_source6).cf38;
          let _326_cf38 = _325___mcc_h4;
          return _dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv17, _326_cf38), _dafny.Seq.of(_326_cf38, (_this).f39));
        } else if (_source6.is_DC21) {
          let _327___mcc_h5 = (_source6).cf33;
          let _328_cf33 = _327___mcc_h5;
          return _dafny.Seq.Concat(_dafny.Seq.of((_this).f28), _dafny.Seq.of(false));
        } else {
          let _329___mcc_h6 = (_source6).cf39;
          let _330_cf39 = _329___mcc_h6;
          return _dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv18), _dafny.Seq.of((_this).f28));
        }
      }(_316_v20);
      let _331_v21;
      _331_v21 = _dafny.Seq.of((_302_v12).f36);
      r2 = (_331_v21)[_module.__default.safeIndex(_module.__default.safeModuloInt(_module.__default.fm36(globalState), _290_v1), new BigNumber((_331_v21).length))];
      let _332_v22;
      _332_v22 = _dafny.Seq.of((_this).f28);
      let _333_v23;
      _333_v23 = _dafny.Seq.of(_332_v22, _332_v22);
      r3 = _module.__default.fm37(p0, false, _290_v1, new BigNumber(((_333_v23)[_module.__default.safeIndex(_290_v1, new BigNumber((_333_v23).length))]).length), globalState);
      return [r0, r1, r2, r3];
    }
    m15(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      let _334_v0;
      let _nw39 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
      _334_v0 = _nw39;
      let _index38 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_334_v0).length));
      (_334_v0)[_index38] = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p1);
      let _335_v1;
      _335_v1 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ubsto"));
      let _336_v2;
      _336_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f39)),(p1).plus((_dafny.ZERO).minus(new BigNumber((_335_v1).length))));
      let _index39 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_334_v0).length));
      (_334_v0)[_index39] = _336_v2;
      let _337_v3;
      let _nw40 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
      _337_v3 = _nw40;
      let _338_v4;
      _338_v4 = _module.D11.create_DC27(_337_v3);
      let _pat_let_tv19 = _337_v3;
      let _339_v5;
      let _nw41 = Array((new BigNumber(11)).toNumber());
      _nw41[(_dafny.ZERO).toNumber()] = _337_v3;
      _nw41[(_dafny.ONE).toNumber()] = _337_v3;
      _nw41[(new BigNumber(2)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(3)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(4)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(5)).toNumber()] = (function (_pat_let9_0) {
        return function (_340_dt__update__tmp_h0) {
          return function (_pat_let10_0) {
            return function (_341_dt__update_hcf42_h0) {
              return _module.D11.create_DC27(_341_dt__update_hcf42_h0);
            }(_pat_let10_0);
          }(_pat_let_tv19);
        }(_pat_let9_0);
      }(_338_v4)).dtor_cf42;
      _nw41[(new BigNumber(6)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(7)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(8)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(9)).toNumber()] = _337_v3;
      _nw41[(new BigNumber(10)).toNumber()] = _337_v3;
      _339_v5 = _nw41;
      let _index40 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_339_v5).length));
      (_339_v5)[_index40] = _337_v3;
      let _index41 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_339_v5).length));
      (_339_v5)[_index41] = _337_v3;
      let _342_v6;
      _342_v6 = _dafny.Seq.of((_this).f27, (_this).f27, (_this).f27);
      let _343_v7;
      _343_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of((_this).f27));
      _342_v6 = _dafny.Seq.Concat((((_343_v7).contains(p1)) ? ((_343_v7).get(p1)) : (_342_v6)), _dafny.Seq.of((_this).f27));
      if ((_this).f28) {
        (globalState).f6 = (p1).isEqualTo(p1);
        let _344_v8;
        _344_v8 = _dafny.Seq.of(p1, p1);
        let _345_v9;
        _345_v9 = _dafny.Seq.Concat(_344_v8, _344_v8);
        _345_v9 = _345_v9;
        (globalState).f15 = ((p0).Intersect(p0)).IsProperSubsetOf((p0).Union(p0));
        let _346_v10;
        let _nw42 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _346_v10 = _nw42;
        let _347_v11;
        _347_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _index42 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_346_v10).length));
        (_346_v10)[_index42] = _347_v11;
        let _index43 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_346_v10).length));
        (_346_v10)[_index43] = (_347_v11).Merge((_347_v11).Merge(_347_v11));
        let _348_v12;
        let _nw43 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _348_v12 = _nw43;
        let _349_v13;
        _349_v13 = _dafny.Seq.of(_348_v12);
        let _350_v14;
        _350_v14 = _dafny.Seq.UnicodeFromString("xuh");
        let _351_v15;
        _351_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(p1)).multipliedBy(p1),(new BigNumber((_350_v14).length)).isLessThan(new BigNumber((_349_v13).length)));
        let _352_v16;
        _352_v16 = _dafny.Seq.of((_this).f28, (_this).f39);
        let _353_v17;
        _353_v17 = _dafny.MultiSet.fromElements(p0, _module.__default.fm38(p1, !((_this).f28), (_this).f28, (_this).f39, globalState));
        _351_v15 = (_351_v15).update(new BigNumber((_dafny.Seq.Concat(_352_v16, _352_v16)).length), (_353_v17).equals(_353_v17));
      } else {
        let _354_v18;
        let _nw44 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _354_v18 = _nw44;
        let _355_v19;
        _355_v19 = _dafny.MultiSet.fromElements(p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber(-418))).length));
        let _356_v20;
        _356_v20 = _dafny.Set.fromElements(_355_v19, _dafny.MultiSet.fromElements(p1, p1), _355_v19);
        let _357_v21;
        _357_v21 = _dafny.Map.Empty.slice().updateUnsafe(_354_v18,_356_v20);
        _357_v21 = (_357_v21).update(_354_v18, (_356_v20).Union(_356_v20));
        if ((_this).f28) {
          let _358_v22;
          let _nw45 = Array((new BigNumber(14)).toNumber()).fill([]);
          _358_v22 = _nw45;
          let _index44 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length));
          (_358_v22)[_index44] = _354_v18;
          let _index45 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length));
          (_358_v22)[_index45] = _354_v18;
          let _arr0 = (_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))];
          let _index46 = _module.__default.safeIndex(new BigNumber(756), new BigNumber(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))]).length));
          _arr0[_index46] = p1;
          let _arr1 = (_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))];
          let _index47 = _module.__default.safeIndex(new BigNumber(756), new BigNumber(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))]).length));
          _arr1[_index47] = (new BigNumber((_dafny.Set.fromElements(p1, new BigNumber((function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of _dafny.IntegerRange(new BigNumber(580), new BigNumber(746))) {
              let _359_v23 = _compr_35;
              if (((new BigNumber(580)).isLessThanOrEqualTo(_359_v23)) && ((_359_v23).isLessThan(new BigNumber(746)))) {
                _coll35.push([_module.__default.safeModuloInt(_359_v23, new BigNumber((_dafny.Seq.of(p1)).length)),(_this).f39]);
              }
            }
            return _coll35;
          }()).length), p1, (_dafny.ZERO).minus(p1))).length)).multipliedBy(p1);
          let _360_v24;
          _360_v24 = _dafny.Seq.UnicodeFromString("txp");
          let _361_v25;
          _361_v25 = _dafny.Seq.of((_this).f39);
          _360_v24 = (_this).fm34((false) && (!((_this).f28)), new BigNumber(222), _dafny.MultiSet.fromElements((_this).f39, (_361_v25)[_module.__default.safeIndex(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))])[_module.__default.safeIndex(new BigNumber(756), new BigNumber(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))]).length))], new BigNumber((_361_v25).length))], p2, (_this).f28, (_this).f28), (_dafny.ZERO).minus(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))])[_module.__default.safeIndex(new BigNumber(756), new BigNumber(((_358_v22)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_358_v22).length))]).length))]), globalState);
          _360_v24 = _dafny.Seq.Concat(_360_v24, _360_v24);
          let _362_v26;
          _362_v26 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f28);
          let _363_v27;
          _363_v27 = _dafny.Seq.of(_355_v19, (_355_v19).update(p1, _module.__default.abs(new BigNumber((_362_v26).length))), _355_v19);
          _363_v27 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-501)), ((_364_v19, _365_p1) => function (_366_i0) {
            return (_364_v19).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(159),_dafny.MultiSet.fromElements((_this).f39))).length), _module.__default.abs(_365_p1));
          })(_355_v19, p1)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), ((_367_v19) => function (_368_i1) {
            return _367_v19;
          })(_355_v19)));
        } else {
          let _369_v28;
          let _nw46 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _369_v28 = _nw46;
          let _index48 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length));
          (_354_v18)[_index48] = p1;
          let _370_v29;
          _370_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f39),_354_v18);
          let _371_v30;
          _371_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
          let _index49 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length));
          let _rhs44 = _module.__default.safeDivisionInt(p1, p1);
          let _rhs45 = _369_v28;
          let _rhs46 = (_dafny.ZERO).minus(new BigNumber((((p2) ? (_370_v29) : ((_370_v29).Merge(_370_v29)))).length));
          let _rhs47 = (((_371_v30).contains(!(p2) || ((_this).f39))) ? ((_371_v30).get(!(p2) || ((_this).f39))) : (p2));
          let _rhs48 = p2;
          let _lhs34 = globalState;
          let _lhs35 = _354_v18;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length));
          let _lhs37 = globalState;
          let _lhs38 = globalState;
          _lhs34.f5 = _rhs44;
          _369_v28 = _rhs45;
          _lhs35[_lhs36] = _rhs46;
          _lhs37.f15 = _rhs47;
          _lhs38.f17 = _rhs48;
          let _372_v31;
          _372_v31 = _dafny.MultiSet.fromElements((_this).f39);
          let _373_v32;
          _373_v32 = _dafny.Seq.UnicodeFromString("i");
          let _374_v33;
          _374_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_354_v18)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length))]);
          let _375_v34;
          _375_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_374_v33).contains(_dafny.Set.fromElements(p2))) ? ((_374_v33).get(_dafny.Set.fromElements(p2))) : (p1)),(_this).f39);
          let _376_v35;
          _376_v35 = new _dafny.CodePoint('i'.codePointAt(0));
          let _377_v36;
          let _nw47 = new _module.C0();
          _nw47.__ctor((_this).fm34((_this).f39, (_354_v18)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length))], _372_v31, p1, globalState), new BigNumber((_dafny.Seq.update(_373_v32, _module.__default.safeIndex(new BigNumber((_375_v34).length), new BigNumber((_373_v32).length)), _376_v35)).length), (_this).f27, !(p2));
          _377_v36 = _nw47;
          let _378_v37;
          let _nw48 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _378_v37 = _nw48;
          let _379_v38;
          _379_v38 = _dafny.Map.Empty.slice().updateUnsafe(_378_v37,p2);
          _379_v38 = (_379_v38).update(_378_v37, (_this).f28);
          let _index50 = _module.__default.safeIndex(new BigNumber(911), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index50] = (_this).f39;
          let _index51 = _module.__default.safeIndex(new BigNumber(911), new BigNumber(((_this).f27).length));
          let _index52 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length));
          let _rhs49 = (_373_v32)[_module.__default.safeIndex((_354_v18)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length))], new BigNumber((_373_v32).length))];
          let _rhs50 = (((_this).f28) ? (p2) : (!((_this).f28)));
          let _rhs51 = (_this).f28;
          let _rhs52 = (p1).minus((((_this).f28) ? ((_dafny.ZERO).minus(_module.__default.fm36(globalState))) : (p1)));
          let _lhs39 = (_this).f27;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(911), new BigNumber(((_this).f27).length));
          let _lhs41 = globalState;
          let _lhs42 = _354_v18;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length));
          _376_v35 = _rhs49;
          _lhs39[_lhs40] = _rhs50;
          _lhs41.f6 = _rhs51;
          _lhs42[_lhs43] = _rhs52;
          let _380_v39;
          let _nw49 = Array((_dafny.ONE).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = p1;
          _380_v39 = _nw49;
          let _381_v40;
          _381_v40 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_354_v18)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_354_v18).length))]),_376_v35);
          let _382_v41;
          _382_v41 = _dafny.Seq.of(p1, p1, new BigNumber((_381_v40).length));
          let _383_v42;
          _383_v42 = _dafny.Map.Empty.slice().updateUnsafe(_382_v41,_380_v39);
          let _384_v43;
          _384_v43 = _dafny.Seq.of(_380_v39, _380_v39, _380_v39);
          let _385_v44;
          _385_v44 = _dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), _376_v35);
          _380_v39 = (((_383_v42).contains(_382_v41)) ? ((_383_v42).get(_382_v41)) : ((_384_v43)[_module.__default.safeIndex(new BigNumber((_385_v44).length), new BigNumber((_384_v43).length))]));
        }
        (globalState).f21 = p1;
        let _386_v45;
        _386_v45 = _dafny.MultiSet.fromElements((_this).f39);
        if (!(p1).isEqualTo((((_this).f28) ? ((((_386_v45).contains(p2)) ? ((_386_v45).get(p2)) : (p1))) : (p1)))) {
          let _387_v46;
          _387_v46 = _dafny.Seq.UnicodeFromString("qouh");
          let _388_v47;
          _388_v47 = new _dafny.CodePoint('c'.codePointAt(0));
          let _389_v48;
          let _nw50 = new _module.C0();
          _nw50.__ctor(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nulmgbv"), _module.__default.safeIndex(new BigNumber((_387_v46).length), new BigNumber((_dafny.Seq.UnicodeFromString("nulmgbv")).length)), _388_v47), _387_v46), p1, (_this).f27, !((_this).f28) || ((_this).f28));
          _389_v48 = _nw50;
          let _390_v49;
          _390_v49 = _module.D0.create_DC1();
          let _391_v50;
          _391_v50 = _module.D0.create_DC4(_390_v49);
          _391_v50 = _391_v50;
          (globalState).f21 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-95), p1));
          let _392_v51;
          _392_v51 = _dafny.Seq.of(p2, (_this).f28);
          let _393_v52;
          _393_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_392_v51).length),(_389_v48).f36);
          let _394_v53;
          _394_v53 = _module.D7.create_DC15(_393_v52);
          let _395_v54;
          _395_v54 = _dafny.Seq.of(_394_v53);
          _395_v54 = _dafny.Seq.Concat(_395_v54, _dafny.Seq.Concat(_395_v54, _module.__default.fm39(globalState)));
          let _396_v55;
          _396_v55 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _397_v56;
          _397_v56 = _dafny.Set.fromElements(p1, new BigNumber(((_396_v55).Merge(_396_v55)).length), p1);
          let _index53 = _module.__default.safeIndex(new BigNumber(102), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index53] = true;
          let _398_v57;
          _398_v57 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_397_v56).Intersect(_dafny.Set.fromElements(new BigNumber((_387_v46).length))));
          let _index54 = _module.__default.safeIndex(new BigNumber(102), new BigNumber(((_this).f27).length));
          let _rhs53 = (((_398_v57).contains(p1)) ? ((_398_v57).get(p1)) : (_397_v56));
          let _rhs54 = (_this).f39;
          let _rhs55 = false;
          let _lhs44 = (_this).f27;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(102), new BigNumber(((_this).f27).length));
          _397_v56 = _rhs53;
          _lhs44[_lhs45] = _rhs54;
          r1 = _rhs55;
        } else {
          let _399_v58;
          _399_v58 = _dafny.Seq.UnicodeFromString("x");
          _399_v58 = _dafny.Seq.Concat(_399_v58, _399_v58);
          let _400_v59;
          _400_v59 = _module.D8.create_DC18(_module.__default.safeModuloInt(p1, p1));
          _400_v59 = _400_v59;
          (globalState).f17 = _module.__default.fm40(globalState);
          (globalState).f5 = p1;
          let _401_v60;
          _401_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(337),_module.__default.safeDivisionInt(new BigNumber((_355_v19).cardinality()), p1));
          let _402_v61;
          _402_v61 = _module.D11.create_DC28();
          let _403_v62;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_399_v58, p1, (_this).f27, p2);
          _403_v62 = _nw51;
          let _404_v63;
          _404_v63 = _dafny.Map.Empty.slice().updateUnsafe(_403_v62,_403_v62.f29);
          let _405_v64;
          _405_v64 = _dafny.Map.Empty.slice().updateUnsafe(_404_v63,p2);
          let _406_v65;
          _406_v65 = new _dafny.CodePoint('c'.codePointAt(0));
          let _407_v66;
          _407_v66 = _module.D0.create_DC0(_405_v64, (_this).f39, new BigNumber(-710), _406_v65, (_this).f28);
          let _index55 = _module.__default.safeIndex(new BigNumber(56), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index55] = (_407_v66).dtor_cf1;
          let _index56 = _module.__default.safeIndex(new BigNumber(56), new BigNumber(((_this).f27).length));
          let _rhs56 = _401_v60;
          let _rhs57 = (_this).f28;
          let _rhs58 = p1;
          let _rhs59 = _module.__default.fm41(_399_v58, globalState);
          let _rhs60 = (_this).f39;
          let _lhs46 = globalState;
          let _lhs47 = globalState;
          let _lhs48 = (_this).f27;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(56), new BigNumber(((_this).f27).length));
          _401_v60 = _rhs56;
          _lhs46.f17 = _rhs57;
          _lhs47.f4 = _rhs58;
          _402_v61 = _rhs59;
          _lhs48[_lhs49] = _rhs60;
        }
        _336_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p1);
      }
      let _408_v67;
      _408_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm42(p1, globalState),p1);
      let _409_v68;
      _409_v68 = new _dafny.CodePoint('r'.codePointAt(0));
      let _410_v69;
      _410_v69 = _dafny.Seq.UnicodeFromString("dpjwj");
      let _hi1 = (((_408_v67).contains(_409_v68)) ? ((_408_v67).get(_409_v68)) : (new BigNumber((_410_v69).length)));
      for (let _411_i2 = p1; _411_i2.isLessThan(_hi1); _411_i2 = _411_i2.plus(_dafny.ONE)) {
        let _412_v70;
        _412_v70 = _dafny.MultiSet.fromElements(_411_i2);
        let _413_v71;
        _413_v71 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oom"));
        let _414_v72;
        _414_v72 = _dafny.Map.Empty.slice().updateUnsafe(false,_409_v68);
        r0 = ((_412_v70).Difference(_dafny.MultiSet.fromElements(_411_i2, p1, _411_i2, new BigNumber(-612), p1))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_413_v71).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f39,(_dafny.ZERO).minus(p1))).length), new BigNumber((_dafny.Seq.UnicodeFromString("my")).length), _411_i2, new BigNumber((_414_v72).length)));
        let _415_v73;
        let _init5 = function (_416_i3) {
          return true;
        };
        let _nw52 = Array((new BigNumber(16)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw52.length); _i0_5++) {
          _nw52[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _415_v73 = _nw52;
        _415_v73 = (_this).f27;
        let _417_v74;
        _417_v74 = _dafny.Seq.of(_411_i2, new BigNumber(403));
        r1 = ((((_this).f39) ? (new BigNumber((p0).length)) : (_411_i2))).isLessThan(new BigNumber((_dafny.Seq.Concat(_417_v74, _417_v74)).length));
        (globalState).f21 = new BigNumber(935);
      }
      let _418_v75;
      _418_v75 = _module.D1.create_DC7(_409_v68, _410_v69);
      let _419_v76;
      let _nw53 = Array((new BigNumber(18)).toNumber());
      _nw53[(_dafny.ZERO).toNumber()] = _409_v68;
      _nw53[(_dafny.ONE).toNumber()] = (_418_v75).dtor_cf13;
      _nw53[(new BigNumber(2)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(3)).toNumber()] = _module.__default.fm42(p1, globalState);
      _nw53[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
      _nw53[(new BigNumber(5)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
      _nw53[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
      _nw53[(new BigNumber(8)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(9)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(10)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(11)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(12)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(13)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(14)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(15)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(16)).toNumber()] = _409_v68;
      _nw53[(new BigNumber(17)).toNumber()] = _409_v68;
      _419_v76 = _nw53;
      let _420_v77;
      _420_v77 = _dafny.Map.Empty.slice().updateUnsafe(_419_v76,(((_this).f39) ? (_410_v69) : (_410_v69)));
      _420_v77 = _420_v77;
      let _421_v78;
      _421_v78 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber(487), p1));
      r0 = _421_v78;
      r1 = p2;
      return [r0, r1];
    }
    get f39() {
      let _this = this;
      return _this._f39;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f27 = [];
      this._f28 = false;
      this.f38 = false;
      this._f37 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f37, f38, f27, f28) {
      let _this = this;
      (_this)._f37 = f37;
      (_this).f38 = f38;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm19(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_this.f38)).cardinality())))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements((_this).f28, _this.f38)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f38,new BigNumber((function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of (_dafny.Set.fromElements((_this).f37, (_this).f37)).Elements) {
          let _422_v0 = _compr_36;
          if ((_dafny.Set.fromElements((_this).f37, (_this).f37)).contains(_422_v0)) {
            _coll36.push([_422_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber((_dafny.Seq.of(_this.f38, _this.f38)).length))).length)]);
          }
        }
        return _coll36;
      }()).length))));
    };
    fm20(globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.fromElements(!(false) || (_this.f38), (_this).f28)).cardinality());
    };
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      if (_this.f38) {
        let _423_v0;
        _423_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f38);
        _423_v0 = (_423_v0).update((_dafny.ZERO).minus((p0).minus(p0)), (_this).f28);
        let _424_v1;
        let _nw54 = Array((new BigNumber(18)).toNumber()).fill([]);
        _424_v1 = _nw54;
        let _425_v2;
        _425_v2 = _dafny.Seq.UnicodeFromString("tnt");
        let _426_v3;
        let _nw55 = new _module.C0();
        _nw55.__ctor(_425_v2, p3, (_this).f27, !(_this.f38));
        _426_v3 = _nw55;
        let _427_v4;
        _427_v4 = _dafny.Seq.of(_426_v3);
        let _428_v5;
        _428_v5 = _dafny.Map.Empty.slice().updateUnsafe(_424_v1,_dafny.Seq.Concat(_427_v4, _427_v4));
        _428_v5 = (_428_v5).update(_424_v1, _427_v4);
        (globalState).f17 = (_426_v3).f28;
        (globalState).f5 = (new BigNumber((_425_v2).length)).plus(_426_v3.f29);
        _425_v2 = _425_v2;
      } else {
        let _429_v6;
        _429_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28);
        let _430_v7;
        _430_v7 = _dafny.Seq.UnicodeFromString("gntq");
        let _431_v8;
        _431_v8 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), p0);
        let _432_v9;
        let _nw56 = Array((new BigNumber(18)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = new BigNumber(((_429_v6).Merge(_429_v6)).length);
        _nw56[(_dafny.ONE).toNumber()] = new BigNumber(607);
        _nw56[(new BigNumber(2)).toNumber()] = p3;
        _nw56[(new BigNumber(3)).toNumber()] = (new BigNumber(-478)).multipliedBy(p0);
        _nw56[(new BigNumber(4)).toNumber()] = (p3).plus(p0);
        _nw56[(new BigNumber(5)).toNumber()] = p0;
        _nw56[(new BigNumber(6)).toNumber()] = new BigNumber(68);
        _nw56[(new BigNumber(7)).toNumber()] = new BigNumber(101);
        _nw56[(new BigNumber(8)).toNumber()] = new BigNumber(996);
        _nw56[(new BigNumber(9)).toNumber()] = new BigNumber((_430_v7).length);
        _nw56[(new BigNumber(10)).toNumber()] = p3;
        _nw56[(new BigNumber(11)).toNumber()] = ((_this.f38) ? (p0) : (p0));
        _nw56[(new BigNumber(12)).toNumber()] = new BigNumber(43);
        _nw56[(new BigNumber(13)).toNumber()] = p3;
        _nw56[(new BigNumber(14)).toNumber()] = (_this).fm20(globalState);
        _nw56[(new BigNumber(15)).toNumber()] = p3;
        _nw56[(new BigNumber(16)).toNumber()] = p3;
        _nw56[(new BigNumber(17)).toNumber()] = new BigNumber((_431_v8).length);
        _432_v9 = _nw56;
        _432_v9 = _432_v9;
        (globalState).f21 = p0;
        let _433_v10;
        _433_v10 = _module.D1.create_DC7(p1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fkuwjpj"), _430_v7));
        let _source7 = _433_v10;
        if (_source7.is_DC6) {
          let _434___mcc_h0 = (_source7).cf12;
          let _435_cf12 = _434___mcc_h0;
          let _436_v11;
          _436_v11 = _module.D0.create_DC1();
          let _437_v12;
          _437_v12 = _module.D0.create_DC4(_436_v11);
          _437_v12 = (((p3).isLessThanOrEqualTo(new BigNumber((_430_v7).length))) ? (_437_v12) : (_437_v12));
          let _438_v13;
          let _nw57 = new _module.C0();
          _nw57.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), function (_439_i0) {
            return (_this).f37;
          }), _435_cf12, (_this).f27, _this.f38);
          _438_v13 = _nw57;
          let _440_v14;
          _440_v14 = _dafny.Seq.of((_435_cf12).isLessThan(p0), (_this).f28, _dafny.Seq.IsProperPrefixOf(_431_v8, _431_v8), (_this).f28, _module.__default.fm22(true, new BigNumber(979), globalState));
          let _rhs61 = true;
          let _rhs62 = (_module.__default.fm21((_this).f28, p3, globalState)).dtor_cf6;
          let _rhs63 = true;
          let _rhs64 = (new BigNumber((_dafny.Seq.of((_this).f28)).length)).isEqualTo(p0);
          let _rhs65 = (_440_v14)[_module.__default.safeIndex(p0, new BigNumber((_440_v14).length))];
          let _lhs50 = globalState;
          let _lhs51 = globalState;
          let _lhs52 = globalState;
          _lhs50.f15 = _rhs61;
          _430_v7 = _rhs62;
          r0 = _rhs63;
          _lhs51.f6 = _rhs64;
          _lhs52.f17 = _rhs65;
          let _index57 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_432_v9).length));
          (_432_v9)[_index57] = ((_dafny.ZERO).minus(_435_cf12)).multipliedBy(_435_cf12);
          let _index58 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_432_v9).length));
          (_432_v9)[_index58] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_435_cf12), (_this).fm20(globalState));
        } else if (_source7.is_DC7) {
          let _441___mcc_h1 = (_source7).cf13;
          let _442___mcc_h2 = (_source7).cf14;
          let _443_cf14 = _442___mcc_h2;
          let _444_cf13 = _441___mcc_h1;
          (globalState).f25 = p1;
          (globalState).f21 = p3;
          let _index59 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_432_v9).length));
          (_432_v9)[_index59] = p0;
          let _index60 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_432_v9).length));
          (_432_v9)[_index60] = _module.__default.safeDivisionInt(p3, new BigNumber(887));
          let _445_v15;
          _445_v15 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.fm23(globalState));
          let _446_v16;
          _446_v16 = _dafny.Set.fromElements((_this).f28, (_this).f28);
          let _447_v17;
          _447_v17 = _module.D1.create_DC6(p3);
          _445_v15 = (_445_v15).update((_446_v16).IsProperSubsetOf(_dafny.Set.fromElements(_this.f38, _this.f38)), _447_v17);
        } else {
          let _448___mcc_h3 = (_source7).cf11;
          let _449_cf11 = _448___mcc_h3;
          (globalState).f13 = (_this).fm20(globalState);
          let _450_v18;
          let _nw58 = Array((_dafny.ONE).toNumber()).fill([]);
          _450_v18 = _nw58;
          let _index61 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_450_v18).length));
          (_450_v18)[_index61] = (_this).f27;
          let _index62 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_450_v18).length));
          (_450_v18)[_index62] = (_this).f27;
          let _451_v19;
          _451_v19 = _dafny.MultiSet.fromElements(p0, new BigNumber(338), p3);
          (globalState).f5 = ((_dafny.ZERO).minus((p0).minus(p0))).plus(new BigNumber((_451_v19).cardinality()));
          let _index63 = _module.__default.safeIndex(new BigNumber(691), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index63] = _this.f38;
          let _452_v20;
          _452_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p3);
          let _453_v21;
          _453_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_452_v20).length),p3);
          let _index64 = _module.__default.safeIndex(new BigNumber(691), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index64] = (((_this).f28) ? (_module.__default.fm22(true, new BigNumber((_dafny.Seq.of(_this.f38, _this.f38)).length), globalState)) : ((new BigNumber((_dafny.Set.fromElements(true)).length)).isLessThan(new BigNumber((_453_v21).length))));
        }
        let _454_v22;
        _454_v22 = _dafny.Seq.of(false, (_this).f28, (_this).f28, _this.f38);
        let _rhs66 = _this.f38;
        let _rhs67 = p0;
        let _rhs68 = (new BigNumber((_454_v22).length)).isLessThanOrEqualTo((_this).fm20(globalState));
        let _lhs53 = globalState;
        let _lhs54 = globalState;
        let _lhs55 = globalState;
        _lhs53.f6 = _rhs66;
        _lhs54.f21 = _rhs67;
        _lhs55.f17 = _rhs68;
        let _455_v23;
        _455_v23 = _dafny.MultiSet.fromElements(_432_v9);
        _455_v23 = _455_v23;
      }
      if ((_this).f28) {
        let _456_v24;
        _456_v24 = _dafny.Map.Empty.slice().updateUnsafe(((!(true)) ? ((_this).f28) : (_this.f38)),p0);
        let _457_v25;
        _457_v25 = _dafny.Set.fromElements(p1);
        _456_v24 = (_456_v24).update((_457_v25).IsProperSubsetOf(_457_v25), p3);
        let _458_v26;
        let _init6 = function (_459_i1) {
          return (_459_i1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("chc")).length));
        };
        let _nw59 = Array((new BigNumber(22)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw59.length); _i0_6++) {
          _nw59[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _458_v26 = _nw59;
        _458_v26 = _458_v26;
        let _460_v27;
        _460_v27 = _module.D0.create_DC1();
        let _source8 = _460_v27;
        if (_source8.is_DC0) {
          let _461___mcc_h4 = (_source8).cf0;
          let _462___mcc_h5 = (_source8).cf1;
          let _463___mcc_h6 = (_source8).cf2;
          let _464___mcc_h7 = (_source8).cf3;
          let _465___mcc_h8 = (_source8).cf4;
          let _466_cf4 = _465___mcc_h8;
          let _467_cf3 = _464___mcc_h7;
          let _468_cf2 = _463___mcc_h6;
          let _469_cf1 = _462___mcc_h5;
          let _470_cf0 = _461___mcc_h4;
          r0 = _469_cf1;
          let _471_v28;
          _471_v28 = _dafny.Set.fromElements((new BigNumber(-739)).minus(p3));
          let _472_v29;
          _472_v29 = _dafny.Map.Empty.slice().updateUnsafe(_468_cf2,new BigNumber(24));
          let _473_v30;
          _473_v30 = _dafny.Map.Empty.slice().updateUnsafe(_472_v29,_471_v28);
          let _474_v31;
          _474_v31 = _dafny.MultiSet.fromElements(_466_cf4);
          let _475_v32;
          _475_v32 = _dafny.Seq.UnicodeFromString("qqvo");
          let _476_v33;
          _476_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22(!(_466_cf4), p0, globalState),_475_v32);
          let _477_v34;
          _477_v34 = _module.D0.create_DC2(_471_v28, _475_v32);
          let _rhs69 = (_this).f28;
          let _rhs70 = _dafny.Set.fromElements((p3).minus(p0), p3, (_dafny.ZERO).minus(new BigNumber(((_474_v31).update(_466_cf4, _module.__default.abs(new BigNumber(((((_476_v33).contains(_466_cf4)) ? ((_476_v33).get(_466_cf4)) : ((_477_v34).dtor_cf6))).length)))).cardinality())));
          let _rhs71 = ((_473_v30).update(_472_v29, _471_v28)).update(_472_v29, function () {
            let _coll37 = new _dafny.Set();
            for (const _compr_37 of _dafny.IntegerRange(new BigNumber(904), new BigNumber(560))) {
              let _478_v35 = _compr_37;
              if (((new BigNumber(904)).isLessThanOrEqualTo(_478_v35)) && ((_478_v35).isLessThan(new BigNumber(560)))) {
                _coll37.add(_module.__default.safeModuloInt(_478_v35, _468_cf2));
              }
            }
            return _coll37;
          }());
          let _lhs56 = globalState;
          _lhs56.f15 = _rhs69;
          _471_v28 = _rhs70;
          _473_v30 = _rhs71;
          (globalState).f21 = _468_cf2;
          let _479_v36;
          _479_v36 = _module.D1.create_DC6(p0);
          (globalState).f6 = ((_module.__default.fm24(globalState)).Difference(_dafny.Set.fromElements(_479_v36, _479_v36, _479_v36, _479_v36))).equals(_dafny.Set.fromElements(_479_v36));
        } else if (_source8.is_DC1) {
          let _480_v37;
          _480_v37 = _dafny.Seq.of(p0);
          let _rhs72 = _this.f38;
          let _rhs73 = _module.__default.fm22((_this).f28, (_dafny.ZERO).minus((_480_v37)[_module.__default.safeIndex(p0, new BigNumber((_480_v37).length))]), globalState);
          let _lhs57 = globalState;
          let _lhs58 = globalState;
          _lhs57.f15 = _rhs72;
          _lhs58.f6 = _rhs73;
          let _481_v38;
          _481_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
          let _482_v39;
          _482_v39 = _module.D3.create_DC10(_this.f38, _this.f38);
          let _483_v40;
          _483_v40 = _dafny.Set.fromElements(_482_v39);
          let _484_v41;
          _484_v41 = _dafny.Set.fromElements(_483_v40, _483_v40);
          let _485_v42;
          _485_v42 = _dafny.Map.Empty.slice().updateUnsafe((((_481_v38).contains(_this.f38)) ? ((_481_v38).get(_this.f38)) : (_this.f38)),((_this.f38) ? (_484_v41) : (_dafny.Set.fromElements(_483_v40, _dafny.Set.fromElements(_module.D3.create_DC10((_this).f28, _this.f38)), _483_v40, _dafny.Set.fromElements(_482_v39, _482_v39, _482_v39, function (_pat_let11_0) {
            return function (_486_dt__update__tmp_h0) {
              return function (_pat_let12_0) {
                return function (_487_dt__update_hcf18_h0) {
                  return _module.D3.create_DC10((_486_dt__update__tmp_h0).dtor_cf17, _487_dt__update_hcf18_h0);
                }(_pat_let12_0);
              }((_this).f28);
            }(_pat_let11_0);
          }(_module.D3.create_DC10(!(_this.f38), (_this).f28)), _482_v39), _dafny.Set.fromElements(_482_v39)))));
          _485_v42 = (_485_v42).update((_this).f28, (_484_v41).Difference(_484_v41));
          (globalState).f5 = _module.__default.safeModuloInt(p3, p3);
          let _488_v43;
          _488_v43 = _dafny.Seq.of(false, _this.f38);
          _488_v43 = _module.__default.fm25(p0, new BigNumber((_module.__default.fm26((_this).f28, p0, new _dafny.CodePoint('n'.codePointAt(0)), globalState)).cardinality()), _module.__default.safeDivisionInt(p0, p3), globalState);
        } else if (_source8.is_DC2) {
          let _489___mcc_h9 = (_source8).cf5;
          let _490___mcc_h10 = (_source8).cf6;
          let _491_cf6 = _490___mcc_h10;
          let _492_cf5 = _489___mcc_h9;
          let _493_v44;
          _493_v44 = _dafny.Seq.of((_this).f28, (_this).f28);
          _493_v44 = ((true) ? (_493_v44) : (_module.__default.fm25(new BigNumber((_491_cf6).length), new BigNumber(689), p0, globalState)));
          let _index65 = _module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index65] = (_this).f28;
          let _index66 = _module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index66] = ((_this).f28) && (!(_dafny.Set.fromElements(new BigNumber(557))).equals(_492_cf5));
          let _494_v45;
          _494_v45 = _dafny.Set.fromElements(_this.f38);
          let _495_v46;
          _495_v46 = _dafny.Seq.of(_494_v45, _494_v45, (_494_v45).Union(_dafny.Set.fromElements(((_this).f27)[_module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f27).length))], !(((_this).f27)[_module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f27).length))]), _this.f38)));
          let _496_v47;
          _496_v47 = _module.D5.create_DC12(_495_v46);
          _495_v46 = (_496_v47).dtor_cf20;
          (globalState).f5 = (_dafny.ZERO).minus(p3);
        } else if (_source8.is_DC3) {
          let _497___mcc_h11 = (_source8).cf7;
          let _498___mcc_h12 = (_source8).cf8;
          let _499___mcc_h13 = (_source8).cf9;
          let _500_cf9 = _499___mcc_h13;
          let _501_cf8 = _498___mcc_h12;
          let _502_cf7 = _497___mcc_h11;
          let _503_v48;
          _503_v48 = _dafny.Set.fromElements(_502_cf7, p3, p3);
          let _504_v49;
          _504_v49 = _dafny.Map.Empty.slice().updateUnsafe(_502_cf7,_503_v48);
          let _rhs74 = _module.__default.fm22(_this.f38, (_dafny.ZERO).minus((p0).minus(p3)), globalState);
          let _rhs75 = _module.__default.fm22((_module.__default.fm0(_504_v49, false, false, globalState)).IsDisjointFrom(_503_v48), _module.__default.safeDivisionInt(p0, _502_cf7), globalState);
          let _rhs76 = _this.f38;
          let _lhs59 = globalState;
          let _lhs60 = globalState;
          let _lhs61 = _this;
          _lhs59.f17 = _rhs74;
          _lhs60.f17 = _rhs75;
          _lhs61.f38 = _rhs76;
          let _505_v50;
          _505_v50 = _dafny.Set.fromElements(_501_cf8, _500_cf9);
          let _506_v51;
          _506_v51 = _module.D0.create_DC1();
          let _507_v52;
          _507_v52 = _module.D0.create_DC4(_506_v51);
          let _508_v53;
          _508_v53 = _dafny.Map.Empty.slice().updateUnsafe(_505_v50,_507_v52);
          _508_v53 = (_508_v53).update(_505_v50, _507_v52);
          (globalState).f13 = ((_501_cf8) ? (new BigNumber(175)) : (p0));
          let _509_v54;
          _509_v54 = _dafny.Seq.UnicodeFromString("v");
          let _510_v55;
          let _nw60 = new _module.C0();
          _nw60.__ctor(_509_v54, new BigNumber(735), (_this).f27, (_this).f28);
          _510_v55 = _nw60;
          let _511_v56;
          _511_v56 = _dafny.Map.Empty.slice().updateUnsafe(_510_v55,_510_v55.f29);
          let _512_v57;
          _512_v57 = _dafny.Map.Empty.slice().updateUnsafe(_511_v56,!(_this.f38));
          let _513_v58;
          _513_v58 = _module.D0.create_DC0(_512_v57, _500_cf9, _510_v55.f29, (_this).f37, (_this).f28);
          (globalState).f17 = (_513_v58).dtor_cf1;
        } else {
          let _514___mcc_h14 = (_source8).cf10;
          let _515_cf10 = _514___mcc_h14;
          let _516_v59;
          _516_v59 = _module.D5.create_DC13((_this).f28, _458_v26, _this.f38);
          let _517_v60;
          _517_v60 = _dafny.Set.fromElements(_516_v59);
          let _518_v61;
          let _nw61 = Array((new BigNumber(10)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = _517_v60;
          _nw61[(_dafny.ONE).toNumber()] = _517_v60;
          _nw61[(new BigNumber(2)).toNumber()] = _517_v60;
          _nw61[(new BigNumber(3)).toNumber()] = (_517_v60).Union(_517_v60);
          _nw61[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_module.D5.create_DC13(_this.f38, _458_v26, (_this).f28));
          _nw61[(new BigNumber(5)).toNumber()] = _517_v60;
          _nw61[(new BigNumber(6)).toNumber()] = _517_v60;
          _nw61[(new BigNumber(7)).toNumber()] = _517_v60;
          _nw61[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_516_v59, _516_v59);
          _nw61[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_516_v59, _516_v59, _516_v59)).Union(_517_v60);
          _518_v61 = _nw61;
          let _index67 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_518_v61).length));
          (_518_v61)[_index67] = _517_v60;
          let _519_v62;
          _519_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.fromElements((_this).f27, (_this).f27));
          let _520_v63;
          _520_v63 = _dafny.MultiSet.fromElements((_this).f27);
          let _521_v64;
          _521_v64 = _dafny.Set.fromElements(p0, p3);
          let _index68 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_518_v61).length));
          let _rhs77 = ((((_519_v62).contains(p0)) ? ((_519_v62).get(p0)) : ((_520_v63).update((_this).f27, _module.__default.abs(new BigNumber((_521_v64).length)))))).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f27));
          let _rhs78 = _dafny.Set.fromElements(_516_v59);
          let _rhs79 = new BigNumber(19);
          let _rhs80 = (_this).f37;
          let _rhs81 = (_this).f28;
          let _lhs62 = globalState;
          let _lhs63 = _518_v61;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_518_v61).length));
          let _lhs65 = globalState;
          let _lhs66 = globalState;
          let _lhs67 = globalState;
          _lhs62.f17 = _rhs77;
          _lhs63[_lhs64] = _rhs78;
          _lhs65.f5 = _rhs79;
          _lhs66.f25 = _rhs80;
          _lhs67.f17 = _rhs81;
          let _522_v65;
          _522_v65 = _dafny.MultiSet.fromElements(p1);
          let _523_v66;
          _523_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_522_v65).cardinality()),_this.f38);
          let _524_v68;
          let _nw62 = Array((new BigNumber(5)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _523_v66;
          _nw62[(_dafny.ONE).toNumber()] = (function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(236), new BigNumber(663))) {
              let _525_v67 = _compr_38;
              if (((new BigNumber(236)).isLessThanOrEqualTo(_525_v67)) && ((_525_v67).isLessThan(new BigNumber(663)))) {
                _coll38.push([_module.__default.safeModuloInt(_525_v67, p3),_this.f38]);
              }
            }
            return _coll38;
          }()).Merge(_523_v66);
          _nw62[(new BigNumber(2)).toNumber()] = _523_v66;
          _nw62[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f28);
          _nw62[(new BigNumber(4)).toNumber()] = _523_v66;
          _524_v68 = _nw62;
          let _526_v69;
          _526_v69 = _dafny.Seq.UnicodeFromString("osp");
          let _527_v70;
          _527_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.D1.create_DC7((_this).f37, _526_v69));
          let _528_v71;
          _528_v71 = _dafny.Seq.of(_527_v70);
          let _rhs82 = function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(637), new BigNumber(773))) {
              let _529_v72 = _compr_39;
              if (((new BigNumber(637)).isLessThanOrEqualTo(_529_v72)) && ((_529_v72).isLessThan(new BigNumber(773)))) {
                _coll39.add((_529_v72).plus(p3));
              }
            }
            return _coll39;
          }();
          let _rhs83 = p0;
          let _rhs84 = (_this).f28;
          let _rhs85 = _524_v68;
          let _rhs86 = _dafny.Seq.of(_module.__default.fm27(_this.f38, p3, _this.f38, globalState));
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          _521_v64 = _rhs82;
          _lhs68.f21 = _rhs83;
          _lhs69.f15 = _rhs84;
          _524_v68 = _rhs85;
          _528_v71 = _rhs86;
          _526_v69 = _526_v69;
          let _530_v73;
          _530_v73 = _module.D0.create_DC2(_521_v64, _526_v69);
          let _531_v74;
          _531_v74 = _dafny.Set.fromElements(_this.f38, ((_this).f28) && ((_this).f28), (new BigNumber(((_530_v73).dtor_cf6).length)).isEqualTo(p0), (_this).f28, !(false));
          _531_v74 = _531_v74;
        }
        (globalState).f21 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
        let _532_v75;
        let _533_v76;
        let _out6;
        let _out7;
        let _outcollector2 = _module.__default.m0(globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _532_v75 = _out6;
        _533_v76 = _out7;
      } else {
        if ((_this).f28) {
          let _534_v77;
          _534_v77 = _dafny.Seq.of(_this.f38);
          let _535_v78;
          _535_v78 = _dafny.MultiSet.fromElements(((_dafny.ZERO).minus(p3)).isEqualTo(p3), (_534_v77)[_module.__default.safeIndex(p0, new BigNumber((_534_v77).length))], _this.f38);
          let _536_v79;
          _536_v79 = _dafny.Seq.UnicodeFromString("gio");
          let _537_v80;
          let _nw63 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _537_v80 = _nw63;
          let _538_v81;
          _538_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this.f38) || ((_534_v77)[_module.__default.safeIndex(p3, new BigNumber((_534_v77).length))]),_dafny.Seq.update(_536_v79, _module.__default.safeIndex(new BigNumber(474), new BigNumber((_536_v79).length)), (_this).f37));
          let _index69 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length));
          (_537_v80)[_index69] = p0;
          let _539_v82;
          _539_v82 = _dafny.Set.fromElements((_this).fm20(globalState));
          let _540_v83;
          _540_v83 = _dafny.Seq.of(_539_v82, _539_v82, _dafny.Set.fromElements(p0), _539_v82, _539_v82);
          let _index70 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length));
          let _rhs87 = (_535_v78).Difference((_535_v78).Intersect(_module.__default.fm26((_534_v77)[_module.__default.safeIndex(p3, new BigNumber((_534_v77).length))], p0, p1, globalState)));
          let _rhs88 = _536_v79;
          let _rhs89 = _537_v80;
          let _rhs90 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_536_v79);
          let _rhs91 = (new BigNumber((_dafny.Seq.Concat(_540_v83, _dafny.Seq.of(_dafny.Set.fromElements(p0)))).length)).plus(p0);
          let _lhs70 = _537_v80;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length));
          _535_v78 = _rhs87;
          _536_v79 = _rhs88;
          _537_v80 = _rhs89;
          _538_v81 = _rhs90;
          _lhs70[_lhs71] = _rhs91;
          let _541_v84;
          let _nw64 = new _module.C0();
          _nw64.__ctor(_536_v79, (_dafny.ZERO).minus(new BigNumber(-36)), (_this).f27, (_this).f28);
          _541_v84 = _nw64;
          (globalState).f21 = (_537_v80)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length))];
          let _542_v85;
          _542_v85 = (_this).f27;
          let _543_v86;
          let _nw65 = new _module.C0();
          _nw65.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fervo"), (_541_v84).f36), (_this).fm20(globalState), (_542_v85), false);
          _543_v86 = _nw65;
          let _index71 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index71] = (_this).f28;
          let _544_v87;
          _544_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("vo"));
          let _545_v89;
          _545_v89 = _dafny.Map.Empty.slice().updateUnsafe(p3,_544_v87);
          let _546_v90;
          _546_v90 = _dafny.Seq.of(_module.__default.fm28(p0, globalState));
          let _547_v92;
          let _nw66 = Array((new BigNumber(15)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _544_v87;
          _nw66[(_dafny.ONE).toNumber()] = _544_v87;
          _nw66[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_537_v80)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length))],_dafny.Seq.UnicodeFromString("je"));
          _nw66[(new BigNumber(3)).toNumber()] = _544_v87;
          _nw66[(new BigNumber(4)).toNumber()] = function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(547), new BigNumber(-19))) {
              let _548_v88 = _compr_40;
              if (((new BigNumber(547)).isLessThanOrEqualTo(_548_v88)) && ((_548_v88).isLessThan(new BigNumber(-19)))) {
                _coll40.push([_module.__default.safeModuloInt(_548_v88, new BigNumber(510)),_536_v79]);
              }
            }
            return _coll40;
          }();
          _nw66[(new BigNumber(5)).toNumber()] = (((_545_v89).contains(p3)) ? ((_545_v89).get(p3)) : ((_546_v90)[_module.__default.safeIndex((_537_v80)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length))], new BigNumber((_546_v90).length))]));
          _nw66[(new BigNumber(6)).toNumber()] = _544_v87;
          _nw66[(new BigNumber(7)).toNumber()] = (_544_v87).Merge((_module.D7.create_DC15(_544_v87)).dtor_cf25);
          _nw66[(new BigNumber(8)).toNumber()] = _module.__default.fm28(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f38,_this.f38)).length), p3, p3)).length), globalState);
          _nw66[(new BigNumber(9)).toNumber()] = (_544_v87).Merge(_544_v87);
          _nw66[(new BigNumber(10)).toNumber()] = function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(568), new BigNumber(385))) {
              let _549_v91 = _compr_41;
              if (((new BigNumber(568)).isLessThanOrEqualTo(_549_v91)) && ((_549_v91).isLessThan(new BigNumber(385)))) {
                _coll41.push([(_549_v91).plus(p0),_dafny.Seq.UnicodeFromString("fnogundpk")]);
              }
            }
            return _coll41;
          }();
          _nw66[(new BigNumber(11)).toNumber()] = _544_v87;
          _nw66[(new BigNumber(12)).toNumber()] = (_544_v87).update(p0, _536_v79);
          _nw66[(new BigNumber(13)).toNumber()] = (((_this).f28) ? (_544_v87) : (_544_v87));
          _nw66[(new BigNumber(14)).toNumber()] = _544_v87;
          _547_v92 = _nw66;
          let _index72 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_547_v92).length));
          (_547_v92)[_index72] = _544_v87;
          let _index73 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f27).length));
          let _index74 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_547_v92).length));
          let _rhs92 = _this.f38;
          let _rhs93 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), (_537_v80)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_537_v80).length))]);
          let _rhs94 = _544_v87;
          let _rhs95 = !((_this).f28) || (_this.f38);
          let _lhs72 = (_this).f27;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f27).length));
          let _lhs74 = globalState;
          let _lhs75 = _547_v92;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_547_v92).length));
          let _lhs77 = globalState;
          _lhs72[_lhs73] = _rhs92;
          _lhs74.f13 = _rhs93;
          _lhs75[_lhs76] = _rhs94;
          _lhs77.f15 = _rhs95;
        } else {
          let _550_v93;
          _550_v93 = _module.D0.create_DC3(p3, (_this).f28, (_this).f28);
          let _551_v94;
          _551_v94 = _module.D0.create_DC4(_550_v93);
          let _552_v95;
          _552_v95 = _dafny.MultiSet.fromElements(_551_v94);
          let _553_v96;
          _553_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p0);
          let _rhs96 = (((_this).f28) ? (new BigNumber((_553_v96).length)) : (p3));
          let _rhs97 = _this.f38;
          let _rhs98 = _dafny.MultiSet.fromElements(_551_v94, _551_v94);
          let _rhs99 = _this.f38;
          let _rhs100 = (_dafny.ZERO).minus(p3);
          let _lhs78 = globalState;
          let _lhs79 = globalState;
          let _lhs80 = globalState;
          let _lhs81 = globalState;
          _lhs78.f4 = _rhs96;
          _lhs79.f17 = _rhs97;
          _552_v95 = _rhs98;
          _lhs80.f15 = _rhs99;
          _lhs81.f13 = _rhs100;
          let _554_v97;
          let _nw67 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _554_v97 = _nw67;
          let _555_v98;
          _555_v98 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _556_v99;
          _556_v99 = _dafny.Seq.of(_555_v98);
          let _557_v100;
          _557_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_556_v99);
          let _index75 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_554_v97).length));
          (_554_v97)[_index75] = (((_557_v100).contains(p0)) ? ((_557_v100).get(p0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), ((_558_v98) => function (_559_i2) {
            return _558_v98;
          })(_555_v98))));
          let _index76 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_554_v97).length));
          (_554_v97)[_index76] = _dafny.Seq.Concat(_556_v99, _556_v99);
          let _560_v101;
          _560_v101 = _dafny.Seq.UnicodeFromString("jjcirftt");
          let _561_v102;
          let _out8;
          _out8 = (_this).m13(_this.f38, (_this).f28, new BigNumber((_module.__default.fm29(_this.f38, true, globalState)).length), _560_v101, globalState);
          _561_v102 = _out8;
          let _562_v103;
          let _563_v104;
          let _out9;
          let _out10;
          let _outcollector3 = _module.__default.m0(globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _562_v103 = _out9;
          _563_v104 = _out10;
          let _564_v105;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_dafny.Seq.Concat(_560_v101, _dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), function (_565_i3) {
            return (_this).f37;
          })), p3, ((!((_this).f28)) ? ((_this).f27) : ((_this).f27)), _module.__default.fm22(false, p0, globalState));
          _564_v105 = _nw68;
        }
        let _566_v106;
        _566_v106 = _dafny.Seq.UnicodeFromString("nke");
        let _567_v107;
        let _nw69 = new _module.C0();
        _nw69.__ctor(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("unybyg"), _566_v106), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("unybyg"), _566_v106)).length)), (_this).f37), new BigNumber((_dafny.Seq.of((_this).fm20(globalState))).length), (_this).f27, (_this).f28);
        _567_v107 = _nw69;
        r0 = _module.__default.fm22(_this.f38, p3, globalState);
        let _568_v108;
        _568_v108 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Set.fromElements(!((_this).f28))).length));
        (globalState).f4 = (_dafny.ZERO).minus((((_568_v108).contains(p3)) ? ((_568_v108).get(p3)) : (p0)));
        if (!(_module.__default.fm22(_this.f38, p3, globalState))) {
          (globalState).f4 = ((_this).fm20(globalState)).multipliedBy(p3);
          (globalState).f17 = (_this).f28;
          let _569_v109;
          _569_v109 = _module.D0.create_DC3((_this).fm20(globalState), _this.f38, (_this).f28);
          let _570_v110;
          _570_v110 = _dafny.Map.Empty.slice().updateUnsafe((_569_v109).dtor_cf7,(_this).f28);
          _570_v110 = (_570_v110).Merge(_module.__default.fm30(_this.f38, _dafny.Seq.UnicodeFromString("a"), _this.f38, globalState));
          let _571_v111;
          _571_v111 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,!(_this.f38));
          r0 = (((_571_v111).contains(_this.f38)) ? ((_571_v111).get(_this.f38)) : ((p3).isLessThanOrEqualTo(p3)));
          let _572_v112;
          _572_v112 = _dafny.Seq.of((!((_this).f28)) || (_this.f38), _this.f38, (!(true)) || (!(!((_this).f28))));
          let _573_v113;
          _573_v113 = _dafny.Seq.of(_572_v112, _module.__default.fm25(p0, (_this).fm20(globalState), p3, globalState));
          _572_v112 = _dafny.Seq.Concat((_573_v113)[_module.__default.safeIndex(p0, new BigNumber((_573_v113).length))], _dafny.Seq.of(_this.f38, (_572_v112)[_module.__default.safeIndex(p3, new BigNumber((_572_v112).length))], !(_this.f38)));
        } else {
          let _574_v114;
          _574_v114 = _dafny.Map.Empty.slice().updateUnsafe(p3,!((_this).f28));
          let _575_v115;
          _575_v115 = _dafny.MultiSet.fromElements(p0, p3);
          _574_v114 = (_574_v114).update(p0, ((_575_v115).update(p3, _module.__default.abs(p3))).IsDisjointFrom(_575_v115));
          let _576_v116;
          let _nw70 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _576_v116 = _nw70;
          let _577_v117;
          _577_v117 = _dafny.MultiSet.fromElements(_576_v116, _576_v116, _576_v116);
          let _index77 = _module.__default.safeIndex(new BigNumber(426), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index77] = !(_577_v117).contains(_576_v116);
          let _index78 = _module.__default.safeIndex(new BigNumber(426), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index78] = _this.f38;
          (globalState).f4 = (_module.__default.safeModuloInt(p0, p0)).minus(p3);
          let _578_v118;
          _578_v118 = _dafny.Seq.of(((_this).f27)[_module.__default.safeIndex(new BigNumber(426), new BigNumber(((_this).f27).length))], (_this).f28);
          let _579_v119;
          let _nw71 = Array((new BigNumber(28)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = p1;
          _nw71[(_dafny.ONE).toNumber()] = p1;
          _nw71[(new BigNumber(2)).toNumber()] = p1;
          _nw71[(new BigNumber(3)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(4)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(5)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(6)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _nw71[(new BigNumber(8)).toNumber()] = p1;
          _nw71[(new BigNumber(9)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(10)).toNumber()] = p1;
          _nw71[(new BigNumber(11)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(12)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(13)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(14)).toNumber()] = ((_this.f38) ? ((_this).f37) : ((_this).f37));
          _nw71[(new BigNumber(15)).toNumber()] = p1;
          _nw71[(new BigNumber(16)).toNumber()] = ((_567_v107).f36)[_module.__default.safeIndex(new BigNumber((_578_v118).length), new BigNumber(((_567_v107).f36).length))];
          _nw71[(new BigNumber(17)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(18)).toNumber()] = p1;
          _nw71[(new BigNumber(19)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw71[(new BigNumber(21)).toNumber()] = p1;
          _nw71[(new BigNumber(22)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(23)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(24)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(25)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(26)).toNumber()] = (_this).f37;
          _nw71[(new BigNumber(27)).toNumber()] = (_this).f37;
          _579_v119 = _nw71;
          _579_v119 = _579_v119;
          (globalState).f4 = (((_this).f28) ? ((_this).fm20(globalState)) : (p0));
        }
      }
      let _580_v120;
      _580_v120 = _module.D0.create_DC3(p3, !((_this).f28), (_this).f28);
      let _581_v121;
      _581_v121 = _dafny.MultiSet.fromElements((_580_v120).dtor_cf8, _this.f38);
      let _582_v122;
      _582_v122 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_581_v121).cardinality()),(_dafny.ZERO).minus(p0));
      _582_v122 = ((_582_v122).Merge(_582_v122)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm20(globalState),p0));
      if (_this.f38) {
        let _583_v123;
        _583_v123 = _dafny.Seq.UnicodeFromString("vkegrqm");
        (globalState).f21 = _module.__default.safeDivisionInt(new BigNumber(314), new BigNumber((_583_v123).length));
        let _584_v124;
        _584_v124 = _dafny.Seq.of((_this).fm20(globalState));
        (globalState).f13 = (p0).minus(new BigNumber((_584_v124).length));
        let _585_v125;
        let _nw72 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _585_v125 = _nw72;
        _585_v125 = p2;
        let _586_v126;
        _586_v126 = _dafny.Seq.of((_this).f28);
        let _587_v127;
        _587_v127 = _dafny.Seq.of(_586_v126);
        let _588_v128;
        _588_v128 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_587_v127, _587_v127),(p3).multipliedBy(new BigNumber((_583_v123).length)));
        let _589_v129;
        _589_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_587_v127);
        let _590_v130;
        let _init7 = function (_591_i4) {
          return (_591_i4).plus(new BigNumber(199));
        };
        let _nw73 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw73.length); _i0_7++) {
          _nw73[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _590_v130 = _nw73;
        let _592_v131;
        _592_v131 = _dafny.MultiSet.fromElements(_590_v130, _590_v130, _590_v130);
        let _593_v132;
        _593_v132 = _dafny.Seq.of(_592_v131);
        _588_v128 = (_588_v128).update(_dafny.Seq.Concat(_587_v127, (((_589_v129).contains(_this.f38)) ? ((_589_v129).get(_this.f38)) : (_dafny.Seq.of(_586_v126, _586_v126)))), new BigNumber(((_dafny.MultiSet.fromElements(_590_v130, _590_v130)).Union((_593_v132)[_module.__default.safeIndex(new BigNumber((_583_v123).length), new BigNumber((_593_v132).length))])).cardinality()));
        let _index79 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_590_v130).length));
        (_590_v130)[_index79] = p3;
        let _index80 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_590_v130).length));
        (_590_v130)[_index80] = _module.__default.safeDivisionInt(p3, new BigNumber(-240));
      } else {
        let _594_v133;
        let _init8 = ((_595_p0) => function (_596_i5) {
          return (_596_i5).multipliedBy(_595_p0);
        })(p0);
        let _nw74 = Array((new BigNumber(14)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw74.length); _i0_8++) {
          _nw74[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _594_v133 = _nw74;
        _594_v133 = _594_v133;
        let _597_v134;
        let _nw75 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _597_v134 = _nw75;
        _597_v134 = _597_v134;
        let _598_v135;
        let _nw76 = Array((new BigNumber(7)).toNumber()).fill([]);
        _598_v135 = _nw76;
        _598_v135 = _598_v135;
        (globalState).f4 = p0;
        let _599_v136;
        _599_v136 = _dafny.Seq.of(_this.f38);
        if (!(new BigNumber((_599_v136).length)).isEqualTo(p3)) {
          let _600_v137;
          _600_v137 = _dafny.Map.Empty.slice().updateUnsafe(true,!((_this).f28) || (_this.f38));
          _600_v137 = (_600_v137).update(false, (_this).f28);
          let _601_v138;
          let _nw77 = Array((new BigNumber(7)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = ((_module.__default.fm22((_this).f28, p0, globalState)) ? (_599_v136) : (_599_v136));
          _nw77[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_this.f38, (_this).f28);
          _nw77[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_599_v136, _599_v136);
          _nw77[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_this.f38, _this.f38);
          _nw77[(new BigNumber(4)).toNumber()] = _599_v136;
          _nw77[(new BigNumber(5)).toNumber()] = _599_v136;
          _nw77[(new BigNumber(6)).toNumber()] = _599_v136;
          _601_v138 = _nw77;
          let _602_v139;
          _602_v139 = _dafny.Seq.UnicodeFromString("kkbyosc");
          let _nw78 = Array((new BigNumber(15)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = _599_v136;
          _nw78[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_this.f38);
          _nw78[(new BigNumber(2)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(3)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(4)).toNumber()] = _module.__default.fm25((_dafny.ZERO).minus(p3), (_dafny.ZERO).minus(p0), p0, globalState);
          _nw78[(new BigNumber(5)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_599_v136, _599_v136), _module.__default.safeIndex(new BigNumber((_602_v139).length), new BigNumber((_dafny.Seq.Concat(_599_v136, _599_v136)).length)), (_this).f28);
          _nw78[(new BigNumber(7)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_599_v136, _599_v136);
          _nw78[(new BigNumber(9)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(10)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(11)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_599_v136, _599_v136);
          _nw78[(new BigNumber(13)).toNumber()] = _599_v136;
          _nw78[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(!(false), (_this).f28, (_this).f28), _dafny.Seq.of(_this.f38, false));
          _601_v138 = _nw78;
          let _603_v140;
          _603_v140 = _dafny.MultiSet.fromElements(p0, p3);
          let _index81 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_594_v133).length));
          (_594_v133)[_index81] = new BigNumber((_603_v140).cardinality());
          let _604_v141;
          let _nw79 = Array((new BigNumber(28)).toNumber()).fill([]);
          _604_v141 = _nw79;
          let _index82 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_604_v141).length));
          (_604_v141)[_index82] = _594_v133;
          let _index83 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_594_v133).length));
          let _index84 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_604_v141).length));
          let _rhs101 = p0;
          let _rhs102 = p3;
          let _rhs103 = p3;
          let _rhs104 = new BigNumber((((_module.__default.fm22((_this).f28, p0, globalState)) ? (_dafny.Seq.Concat(_dafny.Seq.of((_this).f28, (_599_v136)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_599_v136).length))], !(_module.__default.fm22((_this).f28, p0, globalState))), _599_v136)) : (_599_v136))).length);
          let _rhs105 = _594_v133;
          let _lhs82 = _594_v133;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_594_v133).length));
          let _lhs84 = globalState;
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = _604_v141;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_604_v141).length));
          _lhs82[_lhs83] = _rhs101;
          _lhs84.f13 = _rhs102;
          _lhs85.f21 = _rhs103;
          _lhs86.f5 = _rhs104;
          _lhs87[_lhs88] = _rhs105;
          (globalState).f5 = (_594_v133)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_594_v133).length))];
          let _605_v142;
          let _nw80 = new _module.C0();
          _nw80.__ctor(_602_v139, (_594_v133)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_594_v133).length))], (_this).f27, (_this).f28);
          _605_v142 = _nw80;
          let _606_v143;
          _606_v143 = _dafny.Map.Empty.slice().updateUnsafe(_605_v142,new BigNumber(542));
          let _607_v144;
          _607_v144 = _dafny.Map.Empty.slice().updateUnsafe(_606_v143,_this.f38);
          let _608_v145;
          _608_v145 = _dafny.Seq.of(p3);
          let _609_v146;
          _609_v146 = _module.D0.create_DC4(_module.D0.create_DC0(_607_v144, (_this).f28, (_608_v145)[_module.__default.safeIndex(new BigNumber(128), new BigNumber((_608_v145).length))], (_this).f37, (_605_v142).f28));
          let _rhs106 = p3;
          let _rhs107 = _609_v146;
          let _lhs89 = globalState;
          _lhs89.f5 = _rhs106;
          _609_v146 = _rhs107;
        } else {
          let _610_v147;
          _610_v147 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f28),(_this).f28);
          (globalState).f21 = _module.__default.safeModuloInt(new BigNumber((_610_v147).length), p0);
          let _611_v148;
          _611_v148 = _dafny.Seq.of((_this).f27, (_this).f27, (_this).f27, (_this).f27);
          let _612_v149;
          let _nw81 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _612_v149 = _nw81;
          let _rhs108 = _611_v148;
          let _rhs109 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("syuird")).length), p0)).length), p0)).minus(p3);
          let _rhs110 = (_dafny.ZERO).minus(p3);
          let _rhs111 = _612_v149;
          let _lhs90 = globalState;
          let _lhs91 = globalState;
          _611_v148 = _rhs108;
          _lhs90.f21 = _rhs109;
          _lhs91.f5 = _rhs110;
          _612_v149 = _rhs111;
          (globalState).f13 = p0;
          (_this).f38 = _this.f38;
          (globalState).f6 = _this.f38;
        }
      }
      let _613_v150;
      let _init9 = ((_614_p0, _615_p3) => function (_616_i7) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.of(_614_p0)))), _dafny.Seq.of(_module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), ((_617_p3) => function (_618_i8) {
  return _617_p3;
})(_615_p3)), _dafny.Seq.of(_615_p3, _615_p3))), _module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_619_p0) => function (_620_i9) {
  return _619_p0;
})(_614_p0)), _dafny.Seq.of(_614_p0, _614_p0, _615_p3, _615_p3, _615_p3))), _module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.of(_614_p0, _614_p0)))));
      })(p0, p3);
      let _nw82 = Array((new BigNumber(10)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw82.length); _i0_9++) {
        _nw82[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _613_v150 = _nw82;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_613_v150).length))) {
        let _621_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_621_i6)) && ((_621_i6).isLessThan(new BigNumber((_613_v150).length))))) {
          (_613_v150)[(_621_i6)] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), ((_622_p0, _623_p3) => function (_624_i10) {
            return _module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.of(_622_p0, _622_p0, _623_p3)));
          })(p0, p3));
        }
      }
      if (!((_this).f28)) {
        (globalState).f5 = p3;
        let _625_v151;
        _625_v151 = _dafny.Set.fromElements(p3);
        if ((function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(-827), new BigNumber(751))) {
            let _626_v153 = _compr_42;
            if (((new BigNumber(-827)).isLessThanOrEqualTo(_626_v153)) && ((_626_v153).isLessThan(new BigNumber(751)))) {
              _coll42.add(_module.__default.safeDivisionInt(_626_v153, p3));
            }
          }
          return _coll42;
        }()).IsProperSubsetOf(function () {
          let _coll43 = new _dafny.Set();
          for (const _compr_43 of (_625_v151).Elements) {
            let _627_v152 = _compr_43;
            if ((_625_v151).contains(_627_v152)) {
              _coll43.add((_627_v152).plus(new BigNumber(-414)));
            }
          }
          return _coll43;
        }())) {
          let _628_v154;
          let _init10 = function (_629_i11) {
            return (_629_i11).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_630_i12) {
              return (_this).f37;
            })).length));
          };
          let _nw83 = Array((new BigNumber(28)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw83.length); _i0_10++) {
            _nw83[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _628_v154 = _nw83;
          let _index85 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_628_v154).length));
          (_628_v154)[_index85] = new BigNumber(29);
          let _631_v155;
          _631_v155 = _dafny.Seq.UnicodeFromString("mrefi");
          let _632_v156;
          _632_v156 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_631_v155).length),_this.f38);
          let _index86 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_628_v154).length));
          (_628_v154)[_index86] = new BigNumber((_632_v156).length);
          let _633_v157;
          _633_v157 = (_this).f27;
          let _634_v158;
          _634_v158 = _dafny.Map.Empty.slice().updateUnsafe(_633_v157,_580_v120);
          let _635_v159;
          _635_v159 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_633_v157,_580_v120)).Merge(_634_v158),(p3).minus(p3));
          let _index87 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_628_v154).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_628_v154).length));
          let _rhs112 = p3;
          let _rhs113 = p0;
          let _rhs114 = p0;
          let _rhs115 = (_635_v159).update(_634_v158, p3);
          let _rhs116 = p0;
          let _lhs92 = globalState;
          let _lhs93 = _628_v154;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_628_v154).length));
          let _lhs95 = _628_v154;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_628_v154).length));
          let _lhs97 = globalState;
          _lhs92.f4 = _rhs112;
          _lhs93[_lhs94] = _rhs113;
          _lhs95[_lhs96] = _rhs114;
          _635_v159 = _rhs115;
          _lhs97.f4 = _rhs116;
          let _636_v160;
          let _637_v161;
          let _out11;
          let _out12;
          let _outcollector4 = _module.__default.m0(globalState);
          _out11 = _outcollector4[0];
          _out12 = _outcollector4[1];
          _636_v160 = _out11;
          _637_v161 = _out12;
          (globalState).f21 = (p3).multipliedBy((_628_v154)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_628_v154).length))]);
          let _638_v162;
          _638_v162 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC6(_637_v161),false);
          let _639_v163;
          _639_v163 = _module.D1.create_DC6(p0);
          let _640_v164;
          _640_v164 = _dafny.Set.fromElements(_this.f38);
          (globalState).f17 = (((_638_v162).contains(_639_v163)) ? ((_638_v162).get(_639_v163)) : ((_module.__default.fm31(globalState)).equals(_640_v164)));
          let _641_v165;
          let _init11 = function (_642_i13) {
            return (_this).f28;
          };
          let _nw84 = Array((_dafny.ONE).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw84.length); _i0_11++) {
            _nw84[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _641_v165 = _nw84;
          _641_v165 = (_this).f27;
        } else {
          let _index89 = _module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index89] = (_this).f28;
          let _index90 = _module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index90] = (_module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p0))).isLessThan((_dafny.ZERO).minus((((_582_v122).contains(p3)) ? ((_582_v122).get(p3)) : (p3))));
          let _643_v166;
          _643_v166 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)));
          let _644_v167;
          _644_v167 = _dafny.Seq.of(_643_v166, _643_v166);
          let _645_v168;
          let _nw85 = Array((new BigNumber(13)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _this.f38;
          _nw85[(_dafny.ONE).toNumber()] = _this.f38;
          _nw85[(new BigNumber(2)).toNumber()] = (_this).f28;
          _nw85[(new BigNumber(3)).toNumber()] = ((_this).f27)[_module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length))];
          _nw85[(new BigNumber(4)).toNumber()] = !(((_this).f27)[_module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length))]);
          _nw85[(new BigNumber(5)).toNumber()] = _this.f38;
          _nw85[(new BigNumber(6)).toNumber()] = _this.f38;
          _nw85[(new BigNumber(7)).toNumber()] = true;
          _nw85[(new BigNumber(8)).toNumber()] = !((_this).f28);
          _nw85[(new BigNumber(9)).toNumber()] = !((_this).f28);
          _nw85[(new BigNumber(10)).toNumber()] = ((_this).f27)[_module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length))];
          _nw85[(new BigNumber(11)).toNumber()] = (_this).f28;
          _nw85[(new BigNumber(12)).toNumber()] = _this.f38;
          _645_v168 = _nw85;
          let _646_v169;
          let _nw86 = new _module.C0();
          _nw86.__ctor(_module.__default.fm32(p0, ((_this).f27)[_module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length))], globalState), (_dafny.ZERO).minus((new BigNumber((_644_v167).length)).minus((_dafny.ZERO).minus(p0))), _645_v168, (_this).f28);
          _646_v169 = _nw86;
          (globalState).f13 = p0;
          let _index91 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p2).length));
          (p2)[_index91] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("k"), (_646_v169).f36), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("k"), (_646_v169).f36)).length)), p1);
          let _index92 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((p2).length));
          (p2)[_index92] = (_646_v169).f36;
          (globalState).f6 = ((((_this).f27)[_module.__default.safeIndex(new BigNumber(209), new BigNumber(((_this).f27).length))]) ? (_this.f38) : (_this.f38));
        }
        let _647_v170;
        _647_v170 = _dafny.Seq.UnicodeFromString("pteeog");
        let _648_v171;
        let _nw87 = new _module.C0();
        _nw87.__ctor(_647_v170, _module.__default.safeDivisionInt(p3, p3), (_this).f27, _this.f38);
        _648_v171 = _nw87;
        let _649_v172;
        _649_v172 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28);
        let _650_v173;
        _650_v173 = _dafny.Seq.of(_module.__default.fm22((((_649_v172).contains(p0)) ? ((_649_v172).get(p0)) : (_this.f38)), p0, globalState), (_this).f28);
        if (((new BigNumber(695)).isLessThan(new BigNumber((_650_v173).length))) || ((_this).f28)) {
          let _651_v174;
          let _nw88 = Array((new BigNumber(14)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = (_581_v121).Difference(_581_v121);
          _nw88[(_dafny.ONE).toNumber()] = _581_v121;
          _nw88[(new BigNumber(2)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(3)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(4)).toNumber()] = _module.__default.fm26((_this).f28, p3, p1, globalState);
          _nw88[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_650_v173);
          _nw88[(new BigNumber(6)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(7)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(8)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(9)).toNumber()] = (_dafny.MultiSet.FromArray(_650_v173)).Intersect(_581_v121);
          _nw88[(new BigNumber(10)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(11)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(12)).toNumber()] = _581_v121;
          _nw88[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_this.f38);
          _651_v174 = _nw88;
          let _index93 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_651_v174).length));
          (_651_v174)[_index93] = _dafny.MultiSet.fromElements(_this.f38);
          let _index94 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_651_v174).length));
          (_651_v174)[_index94] = ((_this.f38) ? (_581_v121) : ((((_650_v173)[_module.__default.safeIndex(p0, new BigNumber((_650_v173).length))]) ? (_dafny.MultiSet.FromArray(_module.__default.fm25(p0, p0, p0, globalState))) : (_dafny.MultiSet.fromElements(false, _this.f38, (_this).f28, (_this).f28, (_this).f28)))));
          let _652_v175;
          let _nw89 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _652_v175 = _nw89;
          let _653_v176;
          let _nw90 = Array((new BigNumber(2)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _652_v175;
          _nw90[(_dafny.ONE).toNumber()] = _652_v175;
          _653_v176 = _nw90;
          let _index95 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_653_v176).length));
          (_653_v176)[_index95] = _652_v175;
          let _index96 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_653_v176).length));
          (_653_v176)[_index96] = _652_v175;
          let _654_v177;
          _654_v177 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm22(_module.__default.fm22((_this).f28, new BigNumber((_625_v151).length), globalState), p3, globalState));
          let _655_v178;
          _655_v178 = _dafny.Set.fromElements(_654_v177, _654_v177, _654_v177);
          (globalState).f4 = new BigNumber(((_655_v178).Difference((_dafny.Set.fromElements(_654_v177, _654_v177)).Union(_655_v178))).length);
          (globalState).f17 = (p0).isLessThanOrEqualTo(p0);
          (globalState).f5 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(515)), ((_656_v121) => function (_657_i14) {
            return _656_v121;
          })(_581_v121))).length);
        } else {
          let _658_v179;
          let _nw91 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _658_v179 = _nw91;
          let _659_v180;
          _659_v180 = _dafny.Map.Empty.slice().updateUnsafe(p3,_658_v179);
          _659_v180 = (_659_v180).update(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p3)), _658_v179);
          let _660_v181;
          _660_v181 = _dafny.Seq.of(p3, (_dafny.ZERO).minus(p3), p3);
          let _661_v182;
          _661_v182 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_660_v181)).cardinality()), _module.__default.safeModuloInt(p3, p3));
          let _662_v183;
          let _nw92 = Array((new BigNumber(14)).toNumber()).fill(false);
          _662_v183 = _nw92;
          let _663_v184;
          _663_v184 = _dafny.Seq.of(_581_v121);
          let _664_v185;
          let _nw93 = Array((new BigNumber(18)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _581_v121;
          _nw93[(_dafny.ONE).toNumber()] = _581_v121;
          _nw93[(new BigNumber(2)).toNumber()] = (_581_v121).Difference(_581_v121);
          _nw93[(new BigNumber(3)).toNumber()] = (_581_v121).Union(_581_v121);
          _nw93[(new BigNumber(4)).toNumber()] = (_581_v121).Union(_581_v121);
          _nw93[(new BigNumber(5)).toNumber()] = (_663_v184)[_module.__default.safeIndex(new BigNumber((_647_v170).length), new BigNumber((_663_v184).length))];
          _nw93[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(_this.f38, (_this).f28, _module.__default.fm22((_this).f28, p0, globalState), (_this).f28)).Intersect(_dafny.MultiSet.FromArray(_650_v173));
          _nw93[(new BigNumber(7)).toNumber()] = ((!((_this).f28)) ? ((_dafny.MultiSet.FromArray(_650_v173)).update((_this).f28, _module.__default.abs(p0))) : (_581_v121));
          _nw93[(new BigNumber(8)).toNumber()] = ((_this.f38) ? (_581_v121) : (_581_v121));
          _nw93[(new BigNumber(9)).toNumber()] = _581_v121;
          _nw93[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_this.f38);
          _nw93[(new BigNumber(11)).toNumber()] = (_581_v121).Intersect(_581_v121);
          _nw93[(new BigNumber(12)).toNumber()] = (_581_v121).Union(_581_v121);
          _nw93[(new BigNumber(13)).toNumber()] = _581_v121;
          _nw93[(new BigNumber(14)).toNumber()] = (_581_v121).Difference(_581_v121);
          _nw93[(new BigNumber(15)).toNumber()] = _581_v121;
          _nw93[(new BigNumber(16)).toNumber()] = _581_v121;
          _nw93[(new BigNumber(17)).toNumber()] = (_581_v121).Difference(_581_v121);
          _664_v185 = _nw93;
          let _index97 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_664_v185).length));
          (_664_v185)[_index97] = _581_v121;
          let _index98 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_664_v185).length));
          let _rhs117 = ((_this).f28) && (_this.f38);
          let _rhs118 = _661_v182;
          let _rhs119 = (_648_v171).fm15(globalState);
          let _rhs120 = (_this).f27;
          let _rhs121 = _581_v121;
          let _lhs98 = _this;
          let _lhs99 = globalState;
          let _lhs100 = _664_v185;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_664_v185).length));
          _lhs98.f38 = _rhs117;
          _661_v182 = _rhs118;
          _lhs99.f4 = _rhs119;
          _662_v183 = _rhs120;
          _lhs100[_lhs101] = _rhs121;
          let _rhs122 = p3;
          let _rhs123 = _649_v172;
          let _lhs102 = globalState;
          _lhs102.f4 = _rhs122;
          _649_v172 = _rhs123;
          (globalState).f13 = _module.__default.safeModuloInt((p0).minus(p0), (_dafny.ZERO).minus(p3));
          let _665_v186;
          let _init12 = ((_666_p1) => function (_667_i15) {
            return _666_p1;
          })(p1);
          let _nw94 = Array((new BigNumber(9)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw94.length); _i0_12++) {
            _nw94[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _665_v186 = _nw94;
          let _668_v187;
          _668_v187 = _dafny.Seq.of(_665_v186);
          (globalState).f21 = new BigNumber((_668_v187).length);
        }
        let _669_v188;
        let _init13 = function (_670_i16) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_671_i17) {
            return _module.D5.create_DC12(_dafny.Seq.Create(_module.__default.abs(new BigNumber(458)), function (_672_i18) {
  return _dafny.Set.fromElements((_this).f28, (_this).f28, _this.f38);
}));
          });
        };
        let _nw95 = Array((new BigNumber(24)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw95.length); _i0_13++) {
          _nw95[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _669_v188 = _nw95;
        _669_v188 = _669_v188;
      } else {
        let _673_v189;
        _673_v189 = _dafny.Seq.of(_this.f38, (_this).f28);
        let _674_v190;
        _674_v190 = _dafny.MultiSet.fromElements(new BigNumber((_673_v189).length), p3);
        let _675_v191;
        _675_v191 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,p3);
        let _676_v192;
        _676_v192 = _dafny.Seq.of((_674_v190).update(new BigNumber((_675_v191).length), _module.__default.abs(p0)), _674_v190);
        (globalState).f6 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_674_v190, (_676_v192)[_module.__default.safeIndex(p3, new BigNumber((_676_v192).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), ((_677_v190) => function (_678_i19) {
          return _677_v190;
        })(_674_v190))), _676_v192);
        let _679_v193;
        _679_v193 = _dafny.Seq.UnicodeFromString("iiaprsa");
        let _680_v194;
        let _nw96 = new _module.C0();
        _nw96.__ctor(_dafny.Seq.Concat(_679_v193, _679_v193), p0, (_this).f27, _this.f38);
        _680_v194 = _nw96;
        let _681_v195;
        let _nw97 = Array((_dafny.ONE).toNumber()).fill([]);
        _681_v195 = _nw97;
        let _682_v196;
        _682_v196 = _dafny.Seq.of(_681_v195);
        let _683_v197;
        let _nw98 = Array((new BigNumber(20)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = _681_v195;
        _nw98[(_dafny.ONE).toNumber()] = _681_v195;
        _nw98[(new BigNumber(2)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(3)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(4)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(5)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(6)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(7)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(8)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(9)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(10)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(11)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(12)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(13)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(14)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(15)).toNumber()] = (_682_v196)[_module.__default.safeIndex(p3, new BigNumber((_682_v196).length))];
        _nw98[(new BigNumber(16)).toNumber()] = ((!(_this.f38)) ? ((_682_v196)[_module.__default.safeIndex(p0, new BigNumber((_682_v196).length))]) : (_681_v195));
        _nw98[(new BigNumber(17)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(18)).toNumber()] = _681_v195;
        _nw98[(new BigNumber(19)).toNumber()] = _681_v195;
        _683_v197 = _nw98;
        let _index99 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_683_v197).length));
        (_683_v197)[_index99] = _681_v195;
        let _index100 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_683_v197).length));
        let _rhs124 = (_this).f28;
        let _rhs125 = _681_v195;
        let _rhs126 = _this.f38;
        let _lhs103 = globalState;
        let _lhs104 = _683_v197;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_683_v197).length));
        let _lhs106 = globalState;
        _lhs103.f17 = _rhs124;
        _lhs104[_lhs105] = _rhs125;
        _lhs106.f6 = _rhs126;
        let _684_v198;
        _684_v198 = _dafny.Seq.of((_580_v120).dtor_cf7, p3, p0);
        let _685_v199;
        _685_v199 = _684_v198;
        let _686_v200;
        _686_v200 = ((_this.f38) ? (_684_v198) : ((_685_v199)));
        let _source9 = _685_v199;
        let _687___mcc_h15 = _source9;
        let _688_cf15 = _687___mcc_h15;
        let _689_v201;
        let _init14 = ((_690_p1) => function (_691_i20) {
          return _690_p1;
        })(p1);
        let _nw99 = Array((new BigNumber(15)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw99.length); _i0_14++) {
          _nw99[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _689_v201 = _nw99;
        let _692_v202;
        _692_v202 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (_689_v201) : (_689_v201)),_dafny.Seq.IsPrefixOf(_module.__default.fm32(p3, !(_this.f38), globalState), (_680_v194).f36));
        (globalState).f17 = (((_692_v202).contains(_689_v201)) ? ((_692_v202).get(_689_v201)) : ((((_this).f28) ? (false) : ((_this).f28))));
        let _693_v203;
        _693_v203 = _dafny.Set.fromElements((_this).f37);
        let _694_v204;
        _694_v204 = _dafny.Map.Empty.slice().updateUnsafe(_693_v203,(_this).f28);
        let _695_v205;
        _695_v205 = _dafny.Map.Empty.slice().updateUnsafe((p3).isLessThanOrEqualTo(p3),(_694_v204).Merge(_694_v204));
        _695_v205 = (_695_v205).update(((_dafny.ZERO).minus((_this).fm20(globalState))).isEqualTo(p3), _694_v204);
        let _696_v206;
        _696_v206 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sgpyl")).length),p1);
        let _rhs127 = _module.__default.safeModuloInt(p3, new BigNumber((_688_cf15).length));
        let _rhs128 = ((_this.f38) ? (p0) : (((_dafny.ZERO).minus((_this).fm20(globalState))).multipliedBy(new BigNumber((_696_v206).length))));
        let _lhs107 = globalState;
        let _lhs108 = globalState;
        _lhs107.f5 = _rhs127;
        _lhs108.f13 = _rhs128;
        (globalState).f21 = p0;
        let _rhs129 = p0;
        let _rhs130 = (_673_v189)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_673_v189).length))];
        let _lhs109 = globalState;
        _lhs109.f13 = _rhs129;
        r0 = _rhs130;
      }
      r0 = (p3).isEqualTo(p3);
      return r0;
    }
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let _697_i0;
      _697_i0 = _dafny.ZERO;
      L1: {
        while ((p1) || (!(p0))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_697_i0)) {
              break L1;
            }
            _697_i0 = (_697_i0).plus(_dafny.ONE);
            let _698_v0;
            let _nw100 = new _module.C0();
            _nw100.__ctor(p3, p2, (_this).f27, p1);
            _698_v0 = _nw100;
            let _699_v1;
            _699_v1 = _dafny.MultiSet.fromElements(new BigNumber(548));
            let _700_v2;
            _700_v2 = _dafny.Map.Empty.slice().updateUnsafe(_698_v0,(_dafny.ZERO).minus(new BigNumber((_699_v1).cardinality())));
            let _701_v3;
            _701_v3 = _dafny.Map.Empty.slice().updateUnsafe(_700_v2,p0);
            let _702_v4;
            _702_v4 = _module.D0.create_DC0(_701_v3, (_699_v1).IsProperSubsetOf(_699_v1), new BigNumber((_dafny.Set.fromElements(p2)).length), (_this).f37, (_this).f28);
            let _703_v5;
            _703_v5 = _dafny.Set.fromElements(new BigNumber((p3).length));
            let _704_v8;
            _704_v8 = _dafny.Map.Empty.slice().updateUnsafe(_698_v0.f29,function () {
              let _coll44 = new _dafny.Set();
              for (const _compr_44 of _dafny.IntegerRange(new BigNumber(524), new BigNumber(662))) {
                let _705_v7 = _compr_44;
                if (((new BigNumber(524)).isLessThanOrEqualTo(_705_v7)) && ((_705_v7).isLessThan(new BigNumber(662)))) {
                  _coll44.add((_705_v7).plus(p2));
                }
              }
              return _coll44;
            }());
            let _706_v9;
            let _nw101 = Array((new BigNumber(16)).toNumber());
            _nw101[(_dafny.ZERO).toNumber()] = _703_v5;
            _nw101[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(p2)).Union(_703_v5);
            _nw101[(new BigNumber(2)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(3)).toNumber()] = function () {
              let _coll45 = new _dafny.Set();
              for (const _compr_45 of _dafny.IntegerRange(new BigNumber(182), new BigNumber(243))) {
                let _707_v6 = _compr_45;
                if (((new BigNumber(182)).isLessThanOrEqualTo(_707_v6)) && ((_707_v6).isLessThan(new BigNumber(243)))) {
                  _coll45.add((_707_v6).minus(_698_v0.f29));
                }
              }
              return _coll45;
            }();
            _nw101[(new BigNumber(4)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(5)).toNumber()] = (_module.__default.fm0(_704_v8, p0, _this.f38, globalState)).Difference(_703_v5);
            _nw101[(new BigNumber(6)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(7)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(8)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(9)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(10)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(11)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(12)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(new BigNumber(-582));
            _nw101[(new BigNumber(14)).toNumber()] = _703_v5;
            _nw101[(new BigNumber(15)).toNumber()] = _703_v5;
            _706_v9 = _nw101;
            let _708_v10;
            _708_v10 = _dafny.Seq.of(_703_v5);
            let _index101 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_706_v9).length));
            (_706_v9)[_index101] = (_708_v10)[_module.__default.safeIndex(new BigNumber(256), new BigNumber((_708_v10).length))];
            let _709_v11;
            let _init15 = function (_710_i1) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), function (_711_i2) {
                return (_this).f37;
              });
            };
            let _nw102 = Array((new BigNumber(11)).toNumber());
            for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw102.length); _i0_15++) {
              _nw102[_i0_15] = _init15(new BigNumber(_i0_15));
            }
            _709_v11 = _nw102;
            let _index102 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_709_v11).length));
            (_709_v11)[_index102] = p3;
            let _712_v12;
            _712_v12 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(p2));
            let _713_v13;
            _713_v13 = _dafny.Seq.of(_module.__default.fm22(p0, (((_712_v12).contains((_this).f28)) ? ((_712_v12).get((_this).f28)) : (_698_v0.f29)), globalState), (_698_v0).f28);
            let _714_v14;
            let _init16 = ((_715_v0) => function (_716_i3) {
              return (_716_i3).multipliedBy(_715_v0.f29);
            })(_698_v0);
            let _nw103 = Array((new BigNumber(18)).toNumber());
            for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw103.length); _i0_16++) {
              _nw103[_i0_16] = _init16(new BigNumber(_i0_16));
            }
            _714_v14 = _nw103;
            let _717_v15;
            _717_v15 = _dafny.MultiSet.fromElements(_714_v14);
            let _index103 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_706_v9).length));
            let _index104 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_709_v11).length));
            let _rhs131 = _702_v4;
            let _rhs132 = (_703_v5).Difference(_703_v5);
            let _rhs133 = _dafny.Seq.update(p3, _module.__default.safeIndex(new BigNumber((_712_v12).length), new BigNumber((p3).length)), (_this).f37);
            let _rhs134 = (_713_v13)[_module.__default.safeIndex(_module.__default.safeDivisionInt(p2, new BigNumber((_717_v15).cardinality())), new BigNumber((_713_v13).length))];
            let _lhs110 = _706_v9;
            let _lhs111 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_706_v9).length));
            let _lhs112 = _709_v11;
            let _lhs113 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_709_v11).length));
            let _lhs114 = globalState;
            _702_v4 = _rhs131;
            _lhs110[_lhs111] = _rhs132;
            _lhs112[_lhs113] = _rhs133;
            _lhs114.f15 = _rhs134;
            _702_v4 = _702_v4;
            (globalState).f5 = (_dafny.ZERO).minus(p2);
            let _index105 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_709_v11).length));
            (_709_v11)[_index105] = p3;
          }
        }
      }
      if (p0) {
        if (p0) {
          let _718_v16;
          _718_v16 = _dafny.Seq.of(p2);
          let _719_v17;
          _719_v17 = _dafny.Seq.of(new BigNumber((_718_v16).length));
          let _720_v18;
          _720_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_719_v17)[_module.__default.safeIndex(p2, new BigNumber((_719_v17).length))]);
          (globalState).f15 = _module.__default.fm22(true, (_dafny.ZERO).minus((((_720_v18).contains(true)) ? ((_720_v18).get(true)) : (p2))), globalState);
          (globalState).f17 = !_dafny.areEqual(_718_v16, _718_v16);
          (globalState).f15 = (function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(-895), new BigNumber(596))) {
              let _721_v19 = _compr_46;
              if (((new BigNumber(-895)).isLessThanOrEqualTo(_721_v19)) && ((_721_v19).isLessThan(new BigNumber(596)))) {
                _coll46.push([(_721_v19).minus(p2),new BigNumber(647)]);
              }
            }
            return _coll46;
          }()).contains(p2);
          let _722_v20;
          let _nw104 = new _module.C0();
          _nw104.__ctor(p3, (p2).multipliedBy(p2), (_this).f27, false);
          _722_v20 = _nw104;
          (globalState).f15 = false;
        } else {
          let _723_v21;
          let _nw105 = new _module.C0();
          _nw105.__ctor(p3, new BigNumber(-948), (_this).f27, p1);
          _723_v21 = _nw105;
          let _724_v22;
          _724_v22 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _725_v23;
          _725_v23 = _dafny.Seq.of(_module.__default.fm22(true, p2, globalState), (p2).isEqualTo((((_724_v22).contains(p2)) ? ((_724_v22).get(p2)) : (p2))), true);
          _725_v23 = _dafny.Seq.update(_725_v23, _module.__default.safeIndex(p2, new BigNumber((_725_v23).length)), p1);
          let _726_v24;
          _726_v24 = _dafny.Seq.of(p2, p2);
          let _727_v25;
          _727_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_726_v24).length));
          let _728_v26;
          _728_v26 = _dafny.Map.Empty.slice().updateUnsafe(_727_v25,p2);
          let _729_v27;
          _729_v27 = _module.D8.create_DC17(_728_v26);
          let _730_v28;
          _730_v28 = _dafny.Seq.of(_728_v26, _728_v26);
          let _731_v29;
          _731_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,p0);
          let _732_v30;
          let _nw106 = Array((new BigNumber(25)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw106[(_dafny.ONE).toNumber()] = (_723_v21).fm15(globalState);
          _nw106[(new BigNumber(2)).toNumber()] = p2;
          _nw106[(new BigNumber(3)).toNumber()] = p2;
          _nw106[(new BigNumber(4)).toNumber()] = p2;
          _nw106[(new BigNumber(5)).toNumber()] = p2;
          _nw106[(new BigNumber(6)).toNumber()] = p2;
          _nw106[(new BigNumber(7)).toNumber()] = ((false) ? (p2) : (new BigNumber(144)));
          _nw106[(new BigNumber(8)).toNumber()] = new BigNumber((_module.__default.fm33((_dafny.ZERO).minus(p2), (_this).f28, globalState)).length);
          _nw106[(new BigNumber(9)).toNumber()] = (p2).plus(p2);
          _nw106[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("sdbbk")).length)).plus(p2));
          _nw106[(new BigNumber(11)).toNumber()] = p2;
          _nw106[(new BigNumber(12)).toNumber()] = p2;
          _nw106[(new BigNumber(13)).toNumber()] = new BigNumber((((_729_v27).dtor_cf30).Merge((_730_v28)[_module.__default.safeIndex((_723_v21).fm15(globalState), new BigNumber((_730_v28).length))])).length);
          _nw106[(new BigNumber(14)).toNumber()] = p2;
          _nw106[(new BigNumber(15)).toNumber()] = p2;
          _nw106[(new BigNumber(16)).toNumber()] = p2;
          _nw106[(new BigNumber(17)).toNumber()] = p2;
          _nw106[(new BigNumber(18)).toNumber()] = (_726_v24)[_module.__default.safeIndex(p2, new BigNumber((_726_v24).length))];
          _nw106[(new BigNumber(19)).toNumber()] = p2;
          _nw106[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_725_v23)[_module.__default.safeIndex(p2, new BigNumber((_725_v23).length))],(_this).f28)).length));
          _nw106[(new BigNumber(21)).toNumber()] = p2;
          _nw106[(new BigNumber(22)).toNumber()] = p2;
          _nw106[(new BigNumber(23)).toNumber()] = p2;
          _nw106[(new BigNumber(24)).toNumber()] = new BigNumber(((_731_v29).update(_this.f38, (_this).f28)).length);
          _732_v30 = _nw106;
          let _index106 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_732_v30).length));
          (_732_v30)[_index106] = p2;
          let _733_v31;
          _733_v31 = _dafny.MultiSet.fromElements(new BigNumber((_726_v24).length), p2);
          let _index107 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_732_v30).length));
          (_732_v30)[_index107] = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_dafny.ZERO).minus(p2), new BigNumber((_733_v31).cardinality())), _module.__default.safeModuloInt(p2, p2));
          (globalState).f4 = (_732_v30)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_732_v30).length))];
          let _734_v32;
          let _nw107 = new _module.C1();
          _nw107.__ctor(_this.f38, (_this).f27, p1);
          _734_v32 = _nw107;
          let _735_v33;
          _735_v33 = _dafny.Map.Empty.slice().updateUnsafe(_734_v32,(_731_v29).update(false, p0));
          let _736_v34;
          _736_v34 = _dafny.Set.fromElements((_732_v30)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_732_v30).length))], p2, new BigNumber((_735_v33).length), p2);
          (globalState).f15 = (_dafny.Set.fromElements(p2, p2, (_dafny.ZERO).minus((_732_v30)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_732_v30).length))]))).IsSubsetOf(_736_v34);
        }
        let _737_v35;
        _737_v35 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(p2), p2),p0);
        _737_v35 = (_737_v35).update(new BigNumber((_dafny.Seq.UnicodeFromString("iivinjgue")).length), (_this).f28);
        let _738_v36;
        let _nw108 = new _module.C0();
        _nw108.__ctor(p3, p2, (_this).f27, (_this).f28);
        _738_v36 = _nw108;
        _738_v36 = _738_v36;
        let _739_v37;
        let _nw109 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _739_v37 = _nw109;
        let _nw110 = Array((new BigNumber(9)).toNumber());
        _nw110[(_dafny.ZERO).toNumber()] = _739_v37;
        _nw110[(_dafny.ONE).toNumber()] = _739_v37;
        _nw110[(new BigNumber(2)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(3)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(4)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(5)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(6)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(7)).toNumber()] = _739_v37;
        _nw110[(new BigNumber(8)).toNumber()] = _739_v37;
        (globalState).f18 = _nw110;
        let _740_v38;
        _740_v38 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f28));
        let _741_v39;
        _741_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_740_v38);
        let _742_v40;
        _742_v40 = (_this).f27;
        let _pat_let_tv20 = _741_v39;
        let _pat_let_tv21 = _742_v40;
        let _pat_let_tv22 = _740_v38;
        let _pat_let_tv23 = _741_v39;
        let _pat_let_tv24 = _742_v40;
        let _source10 = function (_pat_let13_0) {
          return function (_743_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_744_dt__update_hcf20_h0) {
                return _module.D5.create_DC12(_744_dt__update_hcf20_h0);
              }(_pat_let14_0);
            }((((_pat_let_tv23).contains(_pat_let_tv24)) ? ((_pat_let_tv20).get(_pat_let_tv21)) : (_pat_let_tv22)));
          }(_pat_let13_0);
        }(_module.D5.create_DC12(_740_v38));
        if (_source10.is_DC13) {
          let _745___mcc_h0 = (_source10).cf21;
          let _746___mcc_h1 = (_source10).cf22;
          let _747___mcc_h2 = (_source10).cf23;
          let _748_cf23 = _747___mcc_h2;
          let _749_cf22 = _746___mcc_h1;
          let _750_cf21 = _745___mcc_h0;
          let _751_v41;
          _751_v41 = _dafny.Seq.of(p0);
          let _752_v42;
          _752_v42 = _dafny.Seq.of(_751_v41);
          let _rhs135 = _752_v42;
          let _rhs136 = p2;
          let _lhs115 = globalState;
          _752_v42 = _rhs135;
          _lhs115.f21 = _rhs136;
          let _753_v43;
          _753_v43 = _dafny.MultiSet.fromElements(_this.f38);
          let _754_v44;
          _754_v44 = _dafny.Seq.of(p2);
          let _rhs137 = p2;
          let _rhs138 = _module.__default.safeDivisionInt(new BigNumber((_754_v44).length), p2);
          let _rhs139 = _753_v43;
          let _lhs116 = globalState;
          let _lhs117 = globalState;
          _lhs116.f13 = _rhs137;
          _lhs117.f5 = _rhs138;
          _753_v43 = _rhs139;
          let _755_v45;
          _755_v45 = _dafny.Seq.of((_this).f27, (_this).f27, (_this).f27);
          let _756_v46;
          _756_v46 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Concat(_755_v45, _755_v45), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_755_v45, _755_v45)).length)), (_this).f27), _dafny.Seq.Concat(_755_v45, _755_v45), _755_v45, _dafny.Seq.Concat(_755_v45, _755_v45));
          _755_v45 = (_756_v46)[_module.__default.safeIndex(p2, new BigNumber((_756_v46).length))];
          (globalState).f6 = _748_cf23;
        } else {
          let _757___mcc_h3 = (_source10).cf20;
          let _758_cf20 = _757___mcc_h3;
          let _759_v47;
          let _nw111 = Array((new BigNumber(27)).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = _739_v37;
          _nw111[(_dafny.ONE).toNumber()] = _739_v37;
          _nw111[(new BigNumber(2)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(3)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(4)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(5)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(6)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(7)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(8)).toNumber()] = ((p0) ? (_739_v37) : (_739_v37));
          _nw111[(new BigNumber(9)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(10)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(11)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(12)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(13)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(14)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(15)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(16)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(17)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(18)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(19)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(20)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(21)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(22)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(23)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(24)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(25)).toNumber()] = _739_v37;
          _nw111[(new BigNumber(26)).toNumber()] = _739_v37;
          _759_v47 = _nw111;
          let _760_v48;
          _760_v48 = _dafny.Seq.of(_739_v37);
          let _index108 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_759_v47).length));
          (_759_v47)[_index108] = (_760_v48)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_760_v48).length))];
          let _index109 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_759_v47).length));
          (_759_v47)[_index109] = _739_v37;
          let _761_v49;
          let _nw112 = new _module.C1();
          _nw112.__ctor(p0, (_this).f27, p1);
          _761_v49 = _nw112;
          let _762_v50;
          _762_v50 = _dafny.MultiSet.fromElements(p1);
          let _763_v51;
          _763_v51 = _dafny.Set.fromElements((_761_v49).f39);
          let _764_v52;
          let _nw113 = new _module.C0();
          _nw113.__ctor((_761_v49).fm34((_761_v49).f39, p2, _762_v50, p2, globalState), (p2).multipliedBy(p2), (_this).f27, (_763_v51).contains(true));
          _764_v52 = _nw113;
          let _765_v53;
          let _nw114 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _765_v53 = _nw114;
          _765_v53 = _765_v53;
        }
      } else {
        (globalState).f21 = (p2).minus(((p1) ? (p2) : (p2)));
        let _766_v54;
        let _nw115 = new _module.C1();
        _nw115.__ctor(!(_this.f38), (_this).f27, (_this).f28);
        _766_v54 = _nw115;
        if ((_766_v54).f39) {
          let _767_v55;
          _767_v55 = _dafny.Set.fromElements((_766_v54).f39);
          let _rhs140 = (_dafny.MultiSet.fromElements((_this).f28, (_this).f28)).IsProperSubsetOf(_dafny.MultiSet.fromElements(p0, (_this).f28));
          let _rhs141 = p2;
          let _rhs142 = _module.__default.fm36(globalState);
          let _rhs143 = p2;
          let _rhs144 = !(!((_767_v55).IsSubsetOf(_767_v55)) || (p0));
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          let _lhs120 = globalState;
          let _lhs121 = globalState;
          let _lhs122 = globalState;
          _lhs118.f6 = _rhs140;
          _lhs119.f13 = _rhs141;
          _lhs120.f13 = _rhs142;
          _lhs121.f21 = _rhs143;
          _lhs122.f15 = _rhs144;
          let _768_v56;
          _768_v56 = _dafny.Set.fromElements(new BigNumber(840));
          let _769_v57;
          _769_v57 = _dafny.Seq.of(p2, _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("vxsf")).length), p2), new BigNumber((_768_v56).length), p2);
          let _770_v58;
          _770_v58 = _769_v57;
          _769_v57 = (_770_v58);
          let _771_v59;
          _771_v59 = _module.D8.create_DC18(p2);
          (globalState).f25 = _module.__default.fm42((_771_v59).dtor_cf31, globalState);
          (globalState).f15 = _this.f38;
          let _772_v60;
          _772_v60 = _dafny.Map.Empty.slice().updateUnsafe(!((_766_v54).f39),(_this).f27);
          _772_v60 = (_772_v60).update((_766_v54).f39, (_this).f27);
        } else {
          (_this).f38 = true;
          let _773_v61;
          _773_v61 = _dafny.Seq.of(new BigNumber(863));
          let _774_v62;
          _774_v62 = _773_v61;
          let _775_v63;
          _775_v63 = _dafny.Map.Empty.slice().updateUnsafe(_774_v62,(_773_v61)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_773_v61).length))]);
          let _776_v64;
          _776_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,_775_v63);
          let _777_v65;
          _777_v65 = _dafny.Seq.of((((_776_v64).contains(false)) ? ((_776_v64).get(false)) : (_775_v63)));
          let _778_v66;
          _778_v66 = _module.D0.create_DC3(p2, (_this).f28, p0);
          let _779_v68;
          _779_v68 = _dafny.MultiSet.fromElements(_774_v62);
          let _780_v69;
          _780_v69 = _dafny.Seq.of(p0, p0);
          let _781_v70;
          _781_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe((_766_v54).f39,(_778_v66).dtor_cf7)).update(true, p2))).length),_775_v63);
          let _782_v72;
          let _nw116 = Array((new BigNumber(25)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = (_777_v65)[_module.__default.safeIndex(p2, new BigNumber((_777_v65).length))];
          _nw116[(_dafny.ONE).toNumber()] = _775_v63;
          _nw116[(new BigNumber(2)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(3)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(4)).toNumber()] = (_775_v63).Merge((_775_v63).update(_773_v61, p2));
          _nw116[(new BigNumber(5)).toNumber()] = _module.__default.fm43(p2, _778_v66, p2, _this.f38, globalState);
          _nw116[(new BigNumber(6)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(7)).toNumber()] = (function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_779_v68).Elements) {
              let _783_v67 = _compr_47;
              if ((_779_v68).contains(_783_v67)) {
                _coll47.push([_783_v67,p2]);
              }
            }
            return _coll47;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_774_v62,p2));
          _nw116[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_774_v62,new BigNumber(592));
          _nw116[(new BigNumber(9)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(10)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(11)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_774_v62,p2)).Merge(_775_v63);
          _nw116[(new BigNumber(12)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(13)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(14)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(15)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((p3).length), p2, new BigNumber((_773_v61).length)),new BigNumber((_780_v69).length))).Merge(_775_v63);
          _nw116[(new BigNumber(17)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(18)).toNumber()] = (_775_v63).update(_774_v62, p2);
          _nw116[(new BigNumber(19)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(20)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p2, new BigNumber((p3).length)),p2)).Merge(_775_v63);
          _nw116[(new BigNumber(21)).toNumber()] = _775_v63;
          _nw116[(new BigNumber(22)).toNumber()] = (((_781_v70).contains(p2)) ? ((_781_v70).get(p2)) : (function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of (_779_v68).Elements) {
              let _784_v71 = _compr_48;
              if ((_779_v68).contains(_784_v71)) {
                _coll48.push([_784_v71,p2]);
              }
            }
            return _coll48;
          }()));
          _nw116[(new BigNumber(23)).toNumber()] = (_775_v63).Merge(_775_v63);
          _nw116[(new BigNumber(24)).toNumber()] = (_775_v63).Merge(_775_v63);
          _782_v72 = _nw116;
          let _index110 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_782_v72).length));
          (_782_v72)[_index110] = _775_v63;
          let _index111 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_782_v72).length));
          (_782_v72)[_index111] = (_775_v63).Merge(_775_v63);
          let _785_v73;
          _785_v73 = _dafny.MultiSet.fromElements(p0, true);
          let _786_v74;
          _786_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p2);
          let _787_v75;
          let _nw117 = new _module.C0();
          _nw117.__ctor(_dafny.Seq.UnicodeFromString("at"), new BigNumber((_786_v74).length), (_this).f27, _this.f38);
          _787_v75 = _nw117;
          let _788_v76;
          _788_v76 = _dafny.Map.Empty.slice().updateUnsafe(_787_v75,p2);
          let _789_v77;
          _789_v77 = _dafny.Map.Empty.slice().updateUnsafe(_788_v76,(_766_v54).f39);
          let _790_v78;
          _790_v78 = _dafny.Seq.of((_787_v75).f27);
          let _791_v79;
          _791_v79 = _dafny.Map.Empty.slice().updateUnsafe(_787_v75.f29,p2);
          let _792_v80;
          let _nw118 = Array((new BigNumber(19)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(p2)).plus(new BigNumber(922));
          _nw118[(_dafny.ONE).toNumber()] = (((_this).f28) ? (p2) : (p2));
          _nw118[(new BigNumber(2)).toNumber()] = p2;
          _nw118[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_773_v61).length), new BigNumber((_785_v73).cardinality()));
          _nw118[(new BigNumber(4)).toNumber()] = p2;
          _nw118[(new BigNumber(5)).toNumber()] = p2;
          _nw118[(new BigNumber(6)).toNumber()] = (_module.D0.create_DC0(_789_v77, (_this).f28, new BigNumber((_790_v78).length), (_this).f37, p1)).dtor_cf2;
          _nw118[(new BigNumber(7)).toNumber()] = _787_v75.f29;
          _nw118[(new BigNumber(8)).toNumber()] = (_787_v75.f29).plus(new BigNumber((_791_v79).length));
          _nw118[(new BigNumber(9)).toNumber()] = p2;
          _nw118[(new BigNumber(10)).toNumber()] = _787_v75.f29;
          _nw118[(new BigNumber(11)).toNumber()] = _787_v75.f29;
          _nw118[(new BigNumber(12)).toNumber()] = p2;
          _nw118[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
          _nw118[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((p3).length), (_dafny.ZERO).minus(new BigNumber((p3).length)));
          _nw118[(new BigNumber(15)).toNumber()] = p2;
          _nw118[(new BigNumber(16)).toNumber()] = (p2).minus(_787_v75.f29);
          _nw118[(new BigNumber(17)).toNumber()] = ((((_791_v79).contains((_dafny.ZERO).minus(_787_v75.f29))) ? ((_791_v79).get((_dafny.ZERO).minus(_787_v75.f29))) : (_787_v75.f29))).minus(new BigNumber((_786_v74).length));
          _nw118[(new BigNumber(18)).toNumber()] = new BigNumber((_790_v78).length);
          _792_v80 = _nw118;
          let _index112 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_792_v80).length));
          (_792_v80)[_index112] = p2;
          let _793_v81;
          let _nw119 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
          _793_v81 = _nw119;
          let _794_v82;
          _794_v82 = _dafny.Set.fromElements(_module.__default.fm40(globalState), p1, !(p1));
          let _index113 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_793_v81).length));
          (_793_v81)[_index113] = _794_v82;
          let _index114 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_792_v80).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_793_v81).length));
          let _rhs145 = (_766_v54).f39;
          let _rhs146 = p0;
          let _rhs147 = p2;
          let _rhs148 = _dafny.Set.fromElements(p1);
          let _lhs123 = globalState;
          let _lhs124 = globalState;
          let _lhs125 = _792_v80;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_792_v80).length));
          let _lhs127 = _793_v81;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_793_v81).length));
          _lhs123.f17 = _rhs145;
          _lhs124.f6 = _rhs146;
          _lhs125[_lhs126] = _rhs147;
          _lhs127[_lhs128] = _rhs148;
          let _795_v83;
          _795_v83 = _module.D3.create_DC10(true, _this.f38);
          _795_v83 = _795_v83;
          (globalState).f15 = !(p1);
        }
        let _796_v84;
        let _797_v85;
        let _798_v86;
        let _799_v87;
        let _out13;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector5 = (_766_v54).m14(_this.f38, globalState);
        _out13 = _outcollector5[0];
        _out14 = _outcollector5[1];
        _out15 = _outcollector5[2];
        _out16 = _outcollector5[3];
        _796_v84 = _out13;
        _797_v85 = _out14;
        _798_v86 = _out15;
        _799_v87 = _out16;
        let _800_v88;
        _800_v88 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        _800_v88 = (_800_v88).update(p2, (_766_v54).f39);
      }
      if (!(!(!(p1)))) {
        if ((_this).f28) {
          let _801_v89;
          let _init17 = function (_802_i4) {
            return _this.f38;
          };
          let _nw120 = Array((new BigNumber(20)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw120.length); _i0_17++) {
            _nw120[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _801_v89 = _nw120;
          _801_v89 = _801_v89;
          let _803_v90;
          let _nw121 = new _module.C0();
          _nw121.__ctor(p3, p2, _801_v89, _this.f38);
          _803_v90 = _nw121;
          let _804_v91;
          let _nw122 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _804_v91 = _nw122;
          let _index116 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_804_v91).length));
          (_804_v91)[_index116] = p2;
          let _index117 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_804_v91).length));
          (_804_v91)[_index117] = (p2).multipliedBy(p2);
          (globalState).f4 = (_804_v91)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_804_v91).length))];
          let _805_v92;
          let _806_v93;
          let _out17;
          let _out18;
          let _outcollector6 = _module.__default.m0(globalState);
          _out17 = _outcollector6[0];
          _out18 = _outcollector6[1];
          _805_v92 = _out17;
          _806_v93 = _out18;
        } else {
          let _807_v94;
          _807_v94 = _dafny.Set.fromElements(p1, (_this).f28);
          let _808_v95;
          let _nw123 = new _module.C0();
          _nw123.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(257)), function (_809_i5) {
            return (_this).f37;
          }), p2, (_this).f27, (new BigNumber((p3).length)).isLessThan(new BigNumber((_807_v94).length)));
          _808_v95 = _nw123;
          _808_v95 = _808_v95;
          let _810_v96;
          _810_v96 = _dafny.Seq.UnicodeFromString("cplehhjuc");
          let _811_v97;
          let _nw124 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _811_v97 = _nw124;
          let _index118 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_811_v97).length));
          (_811_v97)[_index118] = new BigNumber(855);
          let _812_v98;
          _812_v98 = _dafny.Map.Empty.slice().updateUnsafe((_807_v94).IsSubsetOf(_807_v94),(_808_v95).fm15(globalState));
          let _index119 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_811_v97).length));
          let _rhs149 = p2;
          let _rhs150 = _dafny.Seq.Concat(p3, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_813_i6) {
            return (_this).f37;
          }), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_814_i6) {
            return (_this).f37;
          })).length)), (_this).f37));
          let _rhs151 = (((_812_v98).contains(p1)) ? ((_812_v98).get(p1)) : (p2));
          let _rhs152 = p2;
          let _lhs129 = globalState;
          let _lhs130 = globalState;
          let _lhs131 = _811_v97;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_811_v97).length));
          _lhs129.f4 = _rhs149;
          _810_v96 = _rhs150;
          _lhs130.f13 = _rhs151;
          _lhs131[_lhs132] = _rhs152;
          let _815_v99;
          _815_v99 = _module.D9.create_DC23(_this.f38);
          let _816_v100;
          _816_v100 = _dafny.MultiSet.fromElements(_module.D9.create_DC23(p0), _815_v99, _815_v99);
          let _817_v101;
          _817_v101 = _dafny.Seq.of(_this.f38);
          let _818_v102;
          _818_v102 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,(_this).f28);
          let _819_v103;
          let _nw125 = new _module.C0();
          _nw125.__ctor((_808_v95).f36, (p2).minus(new BigNumber(((_dafny.MultiSet.fromElements(_816_v100, _816_v100, _dafny.MultiSet.fromElements(_815_v99, _module.D9.create_DC23((_817_v101)[_module.__default.safeIndex(p2, new BigNumber((_817_v101).length))])))).update(_816_v100, _module.__default.abs((_811_v97)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_811_v97).length))]))).cardinality())), (_this).f27, ((((_818_v102).contains(_this.f38)) ? ((_818_v102).get(_this.f38)) : ((_this).f28))) === ((_this).f28));
          _819_v103 = _nw125;
          (globalState).f21 = p2;
          let _820_v104;
          _820_v104 = _dafny.Seq.of((_this).f27, (_this).f27);
          (globalState).f17 = !_dafny.Seq.contains(_dafny.Seq.Concat(_820_v104, _820_v104), (_this).f27);
        }
        let _821_v105;
        _821_v105 = _module.D8.create_DC20(_module.D8.create_DC19());
        let _822_v106;
        _822_v106 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((p3).length)),p2);
        let _823_v107;
        _823_v107 = _module.D8.create_DC17(_822_v106);
        let _pat_let_tv25 = _823_v107;
        let _824_v108;
        _824_v108 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let15_0) {
          return function (_825_dt__update__tmp_h1) {
            return function (_pat_let16_0) {
              return function (_826_dt__update_hcf32_h0) {
                return _module.D8.create_DC20(_826_dt__update_hcf32_h0);
              }(_pat_let16_0);
            }(_pat_let_tv25);
          }(_pat_let15_0);
        }(_821_v105),_dafny.Seq.Concat(p3, _dafny.Seq.UnicodeFromString("dxkphh")));
        (globalState).f4 = new BigNumber((_824_v108).length);
        if (p1) {
          let _827_v109;
          let _828_v110;
          let _out19;
          let _out20;
          let _outcollector7 = _module.__default.m0(globalState);
          _out19 = _outcollector7[0];
          _out20 = _outcollector7[1];
          _827_v109 = _out19;
          _828_v110 = _out20;
          let _829_v111;
          _829_v111 = _dafny.Seq.of(_828_v110);
          let _830_v112;
          let _nw126 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _830_v112 = _nw126;
          let _831_v113;
          _831_v113 = _dafny.Seq.of(_830_v112, _830_v112);
          let _832_v114;
          _832_v114 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_829_v111)[_module.__default.safeIndex(p2, new BigNumber((_829_v111).length))]), (_dafny.ZERO).minus(_828_v110), new BigNumber(-504), (p2).multipliedBy(p2), new BigNumber(((((_this).f28) ? (_831_v113) : (_831_v113))).length));
          (globalState).f21 = (((_832_v114).contains(_828_v110)) ? ((_832_v114).get(_828_v110)) : (_828_v110));
          (globalState).f15 = p0;
          let _833_v115;
          _833_v115 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p3).length),(_this).f28);
          (globalState).f21 = (new BigNumber((_833_v115).length)).multipliedBy(_module.__default.fm36(globalState));
          let _834_v116;
          let _init18 = function (_835_i7) {
            return (_this).f37;
          };
          let _nw127 = Array((new BigNumber(26)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw127.length); _i0_18++) {
            _nw127[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _834_v116 = _nw127;
          let _index120 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_834_v116).length));
          (_834_v116)[_index120] = (_this).f37;
          let _index121 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_834_v116).length));
          (_834_v116)[_index121] = (_this).f37;
        } else {
          let _836_v117;
          _836_v117 = _dafny.Seq.of(false);
          let _837_v118;
          _837_v118 = _dafny.Map.Empty.slice().updateUnsafe(_836_v117,(_this).f28);
          let _838_v119;
          _838_v119 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,(_837_v118).Merge(_837_v118));
          _838_v119 = (_838_v119).update((_this).f28, _837_v118);
          (globalState).f4 = new BigNumber(884);
          (globalState).f15 = _this.f38;
          let _839_v120;
          let _nw128 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _839_v120 = _nw128;
          let _index122 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_839_v120).length));
          (_839_v120)[_index122] = (_module.__default.fm36(globalState)).minus(p2);
          let _index123 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_839_v120).length));
          (_839_v120)[_index123] = (new BigNumber((_836_v117).length)).plus(p2);
          (globalState).f13 = (_839_v120)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_839_v120).length))];
        }
        (globalState).f21 = (p2).plus(p2);
        let _840_v121;
        let _nw129 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _840_v121 = _nw129;
        let _841_v122;
        _841_v122 = _dafny.Seq.of(p2);
        let _842_v123;
        _842_v123 = _dafny.MultiSet.fromElements(_841_v122);
        let _843_v124;
        _843_v124 = _dafny.Seq.of(p0, p0, p1);
        let _844_v125;
        _844_v125 = _dafny.Set.fromElements((_843_v124)[_module.__default.safeIndex(p2, new BigNumber((_843_v124).length))]);
        let _index124 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_840_v121).length));
        (_840_v121)[_index124] = (new BigNumber((_842_v123).cardinality())).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_844_v125).length))).length));
        let _845_v126;
        _845_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,new BigNumber((p3).length));
        let _index125 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_840_v121).length));
        (_840_v121)[_index125] = (((_845_v126).contains((_this).f37)) ? ((_845_v126).get((_this).f37)) : ((_dafny.ZERO).minus(p2)));
      } else {
        let _846_v127;
        let _nw130 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
        _846_v127 = _nw130;
        let _847_v128;
        _847_v128 = _module.D0.create_DC3(p2, p1, p0);
        let _pat_let_tv26 = p1;
        let _848_v129;
        let _nw131 = Array((new BigNumber(26)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _847_v128;
        _nw131[(_dafny.ONE).toNumber()] = _847_v128;
        _nw131[(new BigNumber(2)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(3)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(4)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(5)).toNumber()] = function (_pat_let17_0) {
          return function (_849_dt__update__tmp_h2) {
            return function (_pat_let18_0) {
              return function (_850_dt__update_hcf8_h0) {
                return _module.D0.create_DC3((_849_dt__update__tmp_h2).dtor_cf7, _850_dt__update_hcf8_h0, (_849_dt__update__tmp_h2).dtor_cf9);
              }(_pat_let18_0);
            }(_pat_let_tv26);
          }(_pat_let17_0);
        }(_847_v128);
        _nw131[(new BigNumber(6)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(7)).toNumber()] = _module.D0.create_DC3(p2, p1, p1);
        _nw131[(new BigNumber(8)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(9)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(10)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(11)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(12)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(13)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(14)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(15)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(16)).toNumber()] = _module.D0.create_DC3(p2, p0, p1);
        _nw131[(new BigNumber(17)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(18)).toNumber()] = _module.D0.create_DC3(p2, p0, p0);
        _nw131[(new BigNumber(19)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(20)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(21)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(22)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(23)).toNumber()] = _847_v128;
        _nw131[(new BigNumber(24)).toNumber()] = _module.D0.create_DC3(p2, (_this).f28, p0);
        _nw131[(new BigNumber(25)).toNumber()] = _847_v128;
        _848_v129 = _nw131;
        let _851_v130;
        _851_v130 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)),p1),_848_v129);
        let _index126 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_846_v127).length));
        (_846_v127)[_index126] = _851_v130;
        let _index127 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_846_v127).length));
        (_846_v127)[_index127] = _851_v130;
        let _index128 = _module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index128] = (_this.f38) && ((_this).f28);
        let _index129 = _module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index129] = (p2).isEqualTo((_dafny.ZERO).minus((new BigNumber(829)).multipliedBy(p2)));
        let _852_v131;
        _852_v131 = _dafny.Seq.UnicodeFromString("cqp");
        let _853_v132;
        let _nw132 = Array((new BigNumber(8)).toNumber());
        _nw132[(_dafny.ZERO).toNumber()] = p1;
        _nw132[(_dafny.ONE).toNumber()] = (_this).f28;
        _nw132[(new BigNumber(2)).toNumber()] = (_this).f28;
        _nw132[(new BigNumber(3)).toNumber()] = p1;
        _nw132[(new BigNumber(4)).toNumber()] = p1;
        _nw132[(new BigNumber(5)).toNumber()] = ((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))];
        _nw132[(new BigNumber(6)).toNumber()] = ((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))];
        _nw132[(new BigNumber(7)).toNumber()] = p0;
        _853_v132 = _nw132;
        let _854_v133;
        let _nw133 = new _module.C0();
        _nw133.__ctor(_dafny.Seq.UnicodeFromString("p"), p2, _853_v132, ((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))]);
        _854_v133 = _nw133;
        let _855_v134;
        _855_v134 = _dafny.Map.Empty.slice().updateUnsafe(p1,_854_v133);
        let _rhs153 = ((p0) ? (p2) : (((true) ? (new BigNumber(-641)) : (new BigNumber((_855_v134).length)))));
        let _rhs154 = new BigNumber(187);
        let _rhs155 = p3;
        let _rhs156 = new BigNumber(997);
        let _lhs133 = globalState;
        let _lhs134 = globalState;
        let _lhs135 = globalState;
        _lhs133.f5 = _rhs153;
        _lhs134.f13 = _rhs154;
        _852_v131 = _rhs155;
        _lhs135.f21 = _rhs156;
        let _856_v135;
        let _init19 = ((_857_v133) => function (_858_i8) {
          return (_858_i8).minus((_dafny.ZERO).minus(_857_v133.f29));
        })(_854_v133);
        let _nw134 = Array((new BigNumber(16)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw134.length); _i0_19++) {
          _nw134[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _856_v135 = _nw134;
        _856_v135 = _856_v135;
        let _859_v136;
        _859_v136 = _module.D8.create_DC18(_854_v133.f29);
        let _source11 = _859_v136;
        if (_source11.is_DC18) {
          let _860___mcc_h4 = (_source11).cf31;
          let _861_cf31 = _860___mcc_h4;
          (globalState).f6 = (_847_v128).dtor_cf8;
          let _862_v137;
          _862_v137 = _dafny.MultiSet.fromElements((_854_v133).f28);
          let _863_v138;
          _863_v138 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_862_v137).cardinality()));
          let _864_v139;
          _864_v139 = _dafny.Set.fromElements(((_854_v133).f27), _853_v132);
          _863_v138 = (_863_v138).update(_854_v133.f29, new BigNumber((_864_v139).length));
          (globalState).f21 = _module.__default.fm36(globalState);
          let _865_v140;
          _865_v140 = _dafny.Seq.of((_854_v133).f28, !((_854_v133).f28), ((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))], true);
          let _866_v141;
          _866_v141 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.update(_865_v140, _module.__default.safeIndex(_854_v133.f29, new BigNumber((_865_v140).length)), p1));
          let _867_v142;
          _867_v142 = _dafny.MultiSet.fromElements((_854_v133).f27);
          let _rhs157 = (_867_v142).IsSubsetOf((_867_v142).Intersect(_867_v142));
          let _rhs158 = _866_v141;
          let _rhs159 = (((_dafny.ZERO).minus(_854_v133.f29)).minus(_854_v133.f29)).multipliedBy(_861_cf31);
          let _lhs136 = globalState;
          let _lhs137 = globalState;
          _lhs136.f17 = _rhs157;
          _866_v141 = _rhs158;
          _lhs137.f21 = _rhs159;
        } else if (_source11.is_DC19) {
          (globalState).f13 = p2;
          (_854_v133).f29 = p2;
          let _868_v143;
          let _nw135 = new _module.C1();
          _nw135.__ctor((_854_v133).f28, (_854_v133).f27, p0);
          _868_v143 = _nw135;
          (globalState).f4 = _854_v133.f29;
        } else if (_source11.is_DC17) {
          let _869___mcc_h5 = (_source11).cf30;
          let _870_cf30 = _869___mcc_h5;
          let _871_v144;
          let _nw136 = Array((new BigNumber(17)).toNumber()).fill([]);
          _871_v144 = _nw136;
          let _index130 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v144).length));
          (_871_v144)[_index130] = _856_v135;
          let _index131 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v144).length));
          (_871_v144)[_index131] = _856_v135;
          let _872_v145;
          _872_v145 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_871_v144)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v144).length))]);
          let _873_v146;
          _873_v146 = _dafny.Map.Empty.slice().updateUnsafe(_854_v133.f29,(((_872_v145).contains(((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))])) ? ((_872_v145).get(((_this).f27)[_module.__default.safeIndex(new BigNumber(183), new BigNumber(((_this).f27).length))])) : (_856_v135)));
          _873_v146 = (_873_v146).update((_this).fm20(globalState), (_871_v144)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v144).length))]);
          let _874_v147;
          _874_v147 = _dafny.Seq.of((_854_v133).f28);
          let _875_v148;
          _875_v148 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Concat(_874_v147, _874_v147)).length), _854_v133.f29);
          let _rhs160 = _875_v148;
          let _rhs161 = (new BigNumber(690)).minus(p2);
          let _rhs162 = (_852_v131)[_module.__default.safeIndex(p2, new BigNumber((_852_v131).length))];
          let _lhs138 = globalState;
          let _lhs139 = globalState;
          _875_v148 = _rhs160;
          _lhs138.f5 = _rhs161;
          _lhs139.f25 = _rhs162;
          (globalState).f13 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_854_v133,new BigNumber((p3).length))).length)).minus(p2);
        } else {
          let _876___mcc_h6 = (_source11).cf32;
          let _877_cf32 = _876___mcc_h6;
          let _878_v149;
          let _nw137 = new _module.C1();
          _nw137.__ctor(true, (_854_v133).f27, p1);
          _878_v149 = _nw137;
          let _879_v150;
          _879_v150 = _dafny.Map.Empty.slice().updateUnsafe(p1,_878_v149);
          let _880_v151;
          _880_v151 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p2, _854_v133.f29),(((_879_v150).contains(p1)) ? ((_879_v150).get(p1)) : (_878_v149)));
          _880_v151 = (_880_v151).update(_854_v133.f29, _878_v149);
          let _881_v152;
          let _nw138 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _881_v152 = _nw138;
          let _882_v153;
          _882_v153 = _dafny.Set.fromElements(_854_v133.f29);
          let _index132 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_881_v152).length));
          (_881_v152)[_index132] = _dafny.Map.Empty.slice().updateUnsafe(_882_v153,_854_v133.f29);
          let _883_v154;
          _883_v154 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f28);
          let _884_v155;
          _884_v155 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_883_v154).length),_854_v133.f29);
          let _885_v156;
          _885_v156 = _dafny.Map.Empty.slice().updateUnsafe(true,_854_v133.f29);
          let _886_v157;
          _886_v157 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber(826), p2, new BigNumber((_884_v155).length), p2, (_dafny.ZERO).minus(p2))).Union(_882_v153),_module.__default.safeModuloInt(new BigNumber((_885_v156).length), p2));
          let _index133 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_881_v152).length));
          (_881_v152)[_index133] = _886_v157;
          let _887_v158;
          let _nw139 = Array((new BigNumber(11)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = (_878_v149).f27;
          _nw139[(_dafny.ONE).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(2)).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(3)).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(4)).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(5)).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(6)).toNumber()] = (_878_v149).f27;
          _nw139[(new BigNumber(7)).toNumber()] = (_878_v149).f27;
          _nw139[(new BigNumber(8)).toNumber()] = (_854_v133).f27;
          _nw139[(new BigNumber(9)).toNumber()] = (_878_v149).f27;
          _nw139[(new BigNumber(10)).toNumber()] = (_854_v133).f27;
          _887_v158 = _nw139;
          let _888_v159;
          _888_v159 = _module.D9.create_DC22(_856_v135, false, _887_v158, _852_v131);
          _885_v156 = (_885_v156).update((_888_v159).dtor_cf35, (_854_v133.f29).multipliedBy(_854_v133.f29));
          let _889_v160;
          _889_v160 = _dafny.Map.Empty.slice().updateUnsafe((_878_v149).f28,p3);
          let _890_v161;
          _890_v161 = _dafny.Seq.of(_852_v131, (((_889_v160).contains((_854_v133).f28)) ? ((_889_v160).get((_854_v133).f28)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-306)), function (_891_i9) {
            return (_this).f37;
          }))));
          let _892_v162;
          _892_v162 = _dafny.Seq.of(p2);
          let _893_v163;
          _893_v163 = _module.D1.create_DC7((_this).f37, _852_v131);
          let _894_v164;
          let _nw140 = Array((new BigNumber(25)).toNumber());
          _nw140[(_dafny.ZERO).toNumber()] = _852_v131;
          _nw140[(_dafny.ONE).toNumber()] = _module.__default.fm32((_dafny.ZERO).minus(_854_v133.f29), (_854_v133).f28, globalState);
          _nw140[(new BigNumber(2)).toNumber()] = (_890_v161)[_module.__default.safeIndex(new BigNumber((_892_v162).length), new BigNumber((_890_v161).length))];
          _nw140[(new BigNumber(3)).toNumber()] = p3;
          _nw140[(new BigNumber(4)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(5)).toNumber()] = p3;
          _nw140[(new BigNumber(6)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(7)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(8)).toNumber()] = (_893_v163).dtor_cf14;
          _nw140[(new BigNumber(9)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(10)).toNumber()] = _module.__default.fm32(_854_v133.f29, p0, globalState);
          _nw140[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(p3, p3);
          _nw140[(new BigNumber(12)).toNumber()] = p3;
          _nw140[(new BigNumber(13)).toNumber()] = p3;
          _nw140[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm32(_854_v133.f29, (_this).f28, globalState), _852_v131);
          _nw140[(new BigNumber(15)).toNumber()] = p3;
          _nw140[(new BigNumber(16)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(17)).toNumber()] = p3;
          _nw140[(new BigNumber(18)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(19)).toNumber()] = p3;
          _nw140[(new BigNumber(20)).toNumber()] = (_890_v161)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_890_v161).length))];
          _nw140[(new BigNumber(21)).toNumber()] = _852_v131;
          _nw140[(new BigNumber(22)).toNumber()] = p3;
          _nw140[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_852_v131, p3);
          _nw140[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_852_v131, _852_v131);
          _894_v164 = _nw140;
          let _index134 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_894_v164).length));
          (_894_v164)[_index134] = _dafny.Seq.UnicodeFromString("obsvgeys");
          let _index135 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_894_v164).length));
          (_894_v164)[_index135] = _852_v131;
        }
      }
      let _895_v165;
      _895_v165 = _dafny.Seq.of(new BigNumber(8), p2, p2);
      let _896_v166;
      _896_v166 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(378)), ((_897_p2) => function (_898_i10) {
        return _897_p2;
      })(p2));
      _895_v165 = _dafny.Seq.Concat(_895_v165, ((!((_this).f28)) ? (_895_v165) : ((_896_v166))));
      (globalState).f6 = !(_this.f38) || ((p0) || ((_this).f28));
      (globalState).f4 = p2;
      r0 = _module.__default.fm44(globalState);
      return r0;
    }
    get f37() {
      let _this = this;
      return _this._f37;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
      this._f34 = _dafny.ZERO;
      this._f35 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f34, f35, f29, f27, f28) {
      let _this = this;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f29 = f29;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm14(p0, p1, globalState) {
      let _this = this;
      return ((_this.f29).multipliedBy((_this).f34)).minus(_this.f29);
    };
    m10(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = _dafny.ZERO;
      let _899_v0;
      _899_v0 = _dafny.Seq.of((_this).f27);
      (globalState).f17 = !(_dafny.Seq.IsPrefixOf(_899_v0, _dafny.Seq.Concat(_899_v0, _899_v0)));
      let _900_v1;
      _900_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
      let _901_v2;
      let _nw141 = new _module.C0();
      _nw141.__ctor(_dafny.Seq.UnicodeFromString("phmmcxi"), (_this).f34, (_this).f27, _module.__default.fm16(_900_v1, globalState));
      _901_v2 = _nw141;
      let _902_v3;
      _902_v3 = _dafny.Seq.of((_this).f28, (_this).f28, p0);
      let _903_v4;
      _903_v4 = _dafny.Seq.of((_902_v3)[_module.__default.safeIndex(_this.f29, new BigNumber((_902_v3).length))]);
      let _904_v5;
      _904_v5 = _dafny.Seq.of(_903_v4);
      let _905_v6;
      _905_v6 = _dafny.Seq.of((_this).f34);
      let _906_v7;
      _906_v7 = new _dafny.CodePoint('p'.codePointAt(0));
      let _907_v8;
      let _nw142 = new _module.C0();
      _nw142.__ctor((_901_v2).f36, _this.f29, (_this).f27, p0);
      _907_v8 = _nw142;
      let _908_v9;
      _908_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_901_v2).f36, _module.__default.safeIndex((_905_v6)[_module.__default.safeIndex(_this.f29, new BigNumber((_905_v6).length))], new BigNumber(((_901_v2).f36).length)), _906_v7),_907_v8);
      let _909_v10;
      _909_v10 = _dafny.Set.fromElements((((_908_v9).contains((_901_v2).f36)) ? ((_908_v9).get((_901_v2).f36)) : (_907_v8)));
      let _910_v11;
      _910_v11 = _dafny.Map.Empty.slice().updateUnsafe(_904_v5,_909_v10);
      _910_v11 = (_910_v11).update(_904_v5, _909_v10);
      let _911_v12;
      _911_v12 = _905_v6;
      let _912_v13;
      _912_v13 = _dafny.Seq.of(_905_v6, _905_v6, (_911_v12), _905_v6);
      let _913_v14;
      _913_v14 = _module.D1.create_DC5(_912_v13);
      let _914_v15;
      _914_v15 = _dafny.MultiSet.fromElements(_module.D1.create_DC5(_module.__default.fm17(p0, (_907_v8).f28, new BigNumber(((_901_v2).f36).length), p0, globalState)), _module.D1.create_DC5(_912_v13), _913_v14, _module.D1.create_DC5(_912_v13), _913_v14);
      let _915_v16;
      _915_v16 = _dafny.MultiSet.fromElements(_914_v15);
      _915_v16 = ((_915_v16).update(_914_v15, _module.__default.abs((_dafny.ZERO).minus(new BigNumber(((_901_v2).f36).length))))).Intersect(_915_v16);
      let _916_v17;
      _916_v17 = _module.D0.create_DC3(_907_v8.f29, !((_this).f28), true);
      let _917_v18;
      _917_v18 = _dafny.Seq.of(_916_v17);
      let _918_v19;
      let _init20 = ((_919_v4) => function (_920_i0) {
        return (_920_i0).plus(new BigNumber((_919_v4).length));
      })(_903_v4);
      let _nw143 = Array((new BigNumber(14)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw143.length); _i0_20++) {
        _nw143[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _918_v19 = _nw143;
      let _921_v20;
      _921_v20 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_917_v18).length)).multipliedBy((_this).f34),_918_v19);
      _921_v20 = _921_v20;
      r0 = !((_this).f28);
      let _922_v21;
      _922_v21 = _dafny.MultiSet.fromElements((_907_v8).f28, (_this).f28);
      let _923_v22;
      _923_v22 = _dafny.Seq.of(_922_v21);
      let _924_v23;
      _924_v23 = _923_v22;
      r0 = ((_this).fm14(new BigNumber((_905_v6).length), function () {
        let _coll49 = new _dafny.Set();
        for (const _compr_49 of ((_924_v23)).Elements) {
          let _925_v24 = _compr_49;
          if (_dafny.Seq.contains((_924_v23), _925_v24)) {
            _coll49.add(_925_v24);
          }
        }
        return _coll49;
      }(), globalState)).isLessThanOrEqualTo((_905_v6)[_module.__default.safeIndex((_this).f34, new BigNumber((_905_v6).length))]);
      let _926_v25;
      _926_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f34);
      let _927_v26;
      _927_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_926_v25).contains(p0)) ? ((_926_v25).get(p0)) : (_907_v8.f29)),(_this).f34);
      r1 = (new BigNumber((_dafny.Seq.Concat(_903_v4, _902_v3)).length)).isLessThanOrEqualTo((((_927_v26).contains(new BigNumber(159))) ? ((_927_v26).get(new BigNumber(159))) : (new BigNumber(241))));
      r2 = _906_v7;
      r3 = (_this).f34;
      return [r0, r1, r2, r3];
    }
    m11(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = false;
      let _928_v0;
      _928_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f28);
      let _929_v1;
      _929_v1 = _dafny.Map.Empty.slice().updateUnsafe((_928_v0).update((_this).f34, p0),(_this).f34);
      let _930_v2;
      _930_v2 = _dafny.MultiSet.fromElements(true);
      let _931_v3;
      _931_v3 = _dafny.Set.fromElements(_930_v2);
      (globalState).f13 = (_this).fm14((_this).fm14((((_929_v1).contains(_928_v0)) ? ((_929_v1).get(_928_v0)) : (_this.f29)), _931_v3, globalState), _931_v3, globalState);
      let _hi2 = (_this).f34;
      for (let _932_i0 = (_this).f34; _932_i0.isLessThan(_hi2); _932_i0 = _932_i0.plus(_dafny.ONE)) {
        let _933_v4;
        _933_v4 = _dafny.Seq.of(_932_i0, (_this).f34, new BigNumber(137), new BigNumber(-189), _932_i0);
        let _934_v5;
        _934_v5 = _933_v4;
        let _source12 = _934_v5;
        let _935___mcc_h0 = _source12;
        let _936_cf15 = _935___mcc_h0;
        let _937_v6;
        let _init21 = ((_938_v4) => function (_939_i1) {
          return (_939_i1).minus(new BigNumber((_938_v4).length));
        })(_933_v4);
        let _nw144 = Array((new BigNumber(25)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw144.length); _i0_21++) {
          _nw144[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _937_v6 = _nw144;
        let _index136 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_937_v6).length));
        (_937_v6)[_index136] = (_this).f34;
        let _index137 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_937_v6).length));
        (_937_v6)[_index137] = _this.f29;
        (globalState).f25 = new _dafny.CodePoint('h'.codePointAt(0));
        let _940_v7;
        _940_v7 = _module.D1.create_DC6(_932_i0);
        (globalState).f21 = (_dafny.ZERO).minus((_940_v7).dtor_cf12);
        let _941_v8;
        _941_v8 = _dafny.Map.Empty.slice().updateUnsafe(_940_v7,_932_i0);
        _941_v8 = (_941_v8).update(_940_v7, _this.f29);
        (globalState).f6 = (_this.f29).isLessThan(_932_i0);
        let _index138 = _module.__default.safeIndex(new BigNumber(874), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index138] = p0;
        let _index139 = _module.__default.safeIndex(new BigNumber(874), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index139] = (_this).f28;
        let _index140 = _module.__default.safeIndex(new BigNumber(874), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index140] = false;
      }
      let _942_v9;
      let _nw145 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _942_v9 = _nw145;
      let _943_v10;
      _943_v10 = _dafny.Map.Empty.slice().updateUnsafe(_942_v9,_dafny.Seq.contains(_dafny.Seq.of(false), (((_928_v0).contains(_this.f29)) ? ((_928_v0).get(_this.f29)) : ((_this).f28))));
      _943_v10 = (_943_v10).update(_942_v9, !((_this).f28));
      let _944_v11;
      _944_v11 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _945_i2;
      _945_i2 = _dafny.ZERO;
      L2: {
        while (!(((_module.__default.fm16(_944_v11, globalState)) ? (p0) : ((_this).f28)))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_945_i2)) {
              break L2;
            }
            _945_i2 = (_945_i2).plus(_dafny.ONE);
            if (!(p0)) {
              let _946_v12;
              _946_v12 = _module.D3.create_DC10((_this).f28, p0);
              _928_v0 = (_928_v0).update((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_946_v12).dtor_cf18,_dafny.MultiSet.fromElements((_this).f34))).length)), p0);
              let _947_v13;
              _947_v13 = new _dafny.CodePoint('e'.codePointAt(0));
              let _948_v14;
              let _nw146 = new _module.C0();
              _nw146.__ctor(_module.__default.fm18(p0, (_this).f34, _947_v13, globalState), (_dafny.ZERO).minus((_this).f34), (_this).f27, !((_this).f28) || (p0));
              _948_v14 = _nw146;
              let _949_v15;
              _949_v15 = _dafny.Seq.of(p0);
              let _950_v16;
              _950_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_949_v15).length))).cardinality()));
              let _951_v17;
              _951_v17 = _dafny.Set.fromElements(_950_v16, _950_v16, _950_v16, _950_v16);
              let _952_v18;
              _952_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(704),_951_v17);
              _951_v17 = ((((_952_v18).contains(_this.f29)) ? ((_952_v18).get(_this.f29)) : (_dafny.Set.fromElements(_950_v16)))).Union((_dafny.Set.fromElements(_950_v16, _950_v16, _950_v16)).Union(_951_v17));
              let _953_v19;
              _953_v19 = _dafny.Seq.of((_this).f27, (_this).f27);
              let _954_v20;
              let _nw147 = new _module.C1();
              _nw147.__ctor((_this).f28, (_953_v19)[_module.__default.safeIndex(_this.f29, new BigNumber((_953_v19).length))], p0);
              _954_v20 = _nw147;
              let _955_v21;
              _955_v21 = (_this).f27;
              let _956_v22;
              _956_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_955_v21);
              let _957_v23;
              _957_v23 = _dafny.Set.fromElements(_956_v22, _956_v22);
              let _rhs163 = _dafny.Seq.Concat(_dafny.Seq.Concat(_949_v15, _dafny.Seq.of(p0)), _dafny.Seq.of((_this).f28, p0));
              let _rhs164 = !(((_957_v23).Difference(_957_v23)).IsSubsetOf(_957_v23));
              let _rhs165 = new BigNumber(((_948_v14).f36).length);
              let _rhs166 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), function (_958_i3) {
                return new _dafny.CodePoint('o'.codePointAt(0));
              });
              let _rhs167 = ((!(p0) || ((_this).f28)) ? (_954_v20) : (_954_v20));
              let _lhs140 = _this;
              _949_v15 = _rhs163;
              r2 = _rhs164;
              _lhs140.f29 = _rhs165;
              r1 = _rhs166;
              _954_v20 = _rhs167;
              let _959_v24;
              _959_v24 = _dafny.MultiSet.fromElements((_this).f34);
              _959_v24 = (_959_v24).Intersect(_959_v24);
            } else {
              (globalState).f21 = ((_this).f34).plus((new BigNumber(783)).plus(_this.f29));
              (globalState).f6 = (_this).f28;
              let _960_v25;
              let _init22 = ((_961_v11, _962_p0) => function (_963_i4) {
                return (_dafny.Set.fromElements((((_961_v11).contains((_this).f28)) ? ((_961_v11).get((_this).f28)) : (_962_p0)), (((_961_v11).contains((_this).f28)) ? ((_961_v11).get((_this).f28)) : ((_this).f28)))).Union(_dafny.Set.fromElements((_this).f28));
              })(_944_v11, p0);
              let _nw148 = Array((new BigNumber(22)).toNumber());
              for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw148.length); _i0_22++) {
                _nw148[_i0_22] = _init22(new BigNumber(_i0_22));
              }
              _960_v25 = _nw148;
              let _964_v26;
              _964_v26 = _dafny.Set.fromElements((_this).f28);
              let _index141 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_960_v25).length));
              (_960_v25)[_index141] = _964_v26;
              let _index142 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_960_v25).length));
              (_960_v25)[_index142] = _964_v26;
              (globalState).f6 = p0;
              (globalState).f13 = _this.f29;
            }
            r2 = !((_this).f28);
            let _index143 = _module.__default.safeIndex(new BigNumber(809), new BigNumber(((_this).f27).length));
            ((_this).f27)[_index143] = ((p0) ? (p0) : ((_this).f28));
            let _index144 = _module.__default.safeIndex(new BigNumber(809), new BigNumber(((_this).f27).length));
            ((_this).f27)[_index144] = p0;
            let _965_v27;
            let _nw149 = Array((new BigNumber(14)).toNumber()).fill([]);
            _965_v27 = _nw149;
            let _966_v28;
            let _nw150 = Array((new BigNumber(13)).toNumber()).fill([]);
            _966_v28 = _nw150;
            let _index145 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_965_v27).length));
            (_965_v27)[_index145] = _966_v28;
            let _967_v29;
            _967_v29 = _dafny.Seq.UnicodeFromString("djlpjj");
            let _index146 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_965_v27).length));
            let _index147 = _module.__default.safeIndex(new BigNumber(809), new BigNumber(((_this).f27).length));
            let _rhs168 = (_dafny.ZERO).minus(new BigNumber((_967_v29).length));
            let _rhs169 = _966_v28;
            let _rhs170 = ((_this).f27)[_module.__default.safeIndex(new BigNumber(809), new BigNumber(((_this).f27).length))];
            let _lhs141 = globalState;
            let _lhs142 = _965_v27;
            let _lhs143 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_965_v27).length));
            let _lhs144 = (_this).f27;
            let _lhs145 = _module.__default.safeIndex(new BigNumber(809), new BigNumber(((_this).f27).length));
            _lhs141.f5 = _rhs168;
            _lhs142[_lhs143] = _rhs169;
            _lhs144[_lhs145] = _rhs170;
          }
        }
      }
      let _968_v30;
      _968_v30 = _dafny.Seq.of(p0);
      let _969_v31;
      let _nw151 = Array((_dafny.ONE).toNumber());
      _nw151[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f28, p0), _968_v30);
      _969_v31 = _nw151;
      let _index148 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_969_v31).length));
      (_969_v31)[_index148] = _968_v30;
      let _index149 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_969_v31).length));
      (_969_v31)[_index149] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_968_v30, _module.__default.safeIndex(_this.f29, new BigNumber((_968_v30).length)), (_this).f28), _968_v30), _968_v30);
      let _970_v32;
      _970_v32 = _dafny.Seq.UnicodeFromString("vn");
      let _971_v33;
      _971_v33 = _dafny.Seq.of(_this.f29);
      let _972_v34;
      _972_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_970_v32).length)),_971_v33);
      let _973_v35;
      _973_v35 = (((_972_v34).contains(new BigNumber(93))) ? ((_972_v34).get(new BigNumber(93))) : (_971_v33));
      let _source13 = function (_source14) {
        let _974___mcc_h4 = _source14;
        let _975_cf15 = _974___mcc_h4;
        return _module.D8.create_DC20(_module.D8.create_DC20(_module.D8.create_DC20(_module.D8.create_DC18((_this).f34))));
      }(_973_v35);
      if (_source13.is_DC18) {
        let _976___mcc_h1 = (_source13).cf31;
        let _977_cf31 = _976___mcc_h1;
        let _978_v36;
        _978_v36 = _module.D1.create_DC6((_this).f34);
        let _979_v37;
        _979_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f29);
        let _980_v38;
        _980_v38 = _dafny.MultiSet.fromElements(_977_cf31);
        let _981_v39;
        _981_v39 = _dafny.Set.fromElements(_977_cf31, (_dafny.ZERO).minus(new BigNumber((_970_v32).length)));
        let _982_v40;
        _982_v40 = _dafny.Map.Empty.slice().updateUnsafe(_981_v39,p0);
        let _983_v41;
        let _nw152 = new _module.C0();
        _nw152.__ctor(_970_v32, _977_cf31, (_this).f27, (_this).f28);
        _983_v41 = _nw152;
        let _984_v42;
        _984_v42 = _dafny.Seq.of(_983_v41, _983_v41);
        let _985_v43;
        _985_v43 = new _dafny.CodePoint('y'.codePointAt(0));
        let _986_v44;
        _986_v44 = _dafny.Map.Empty.slice().updateUnsafe(_985_v43,(_this).f28);
        let _987_v45;
        let _nw153 = Array((new BigNumber(25)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = (_this).f34;
        _nw153[(_dafny.ONE).toNumber()] = (_978_v36).dtor_cf12;
        _nw153[(new BigNumber(2)).toNumber()] = (_this).fm14(new BigNumber(((_979_v37).update(new BigNumber((_980_v38).cardinality()), _this.f29)).length), _931_v3, globalState);
        _nw153[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_970_v32).length), new BigNumber((_971_v33).length));
        _nw153[(new BigNumber(4)).toNumber()] = _977_cf31;
        _nw153[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_977_cf31);
        _nw153[(new BigNumber(6)).toNumber()] = _977_cf31;
        _nw153[(new BigNumber(7)).toNumber()] = (_this).f34;
        _nw153[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_928_v0).length));
        _nw153[(new BigNumber(9)).toNumber()] = _module.__default.fm36(globalState);
        _nw153[(new BigNumber(10)).toNumber()] = _977_cf31;
        _nw153[(new BigNumber(11)).toNumber()] = new BigNumber((_970_v32).length);
        _nw153[(new BigNumber(12)).toNumber()] = (_this).f34;
        _nw153[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_this).f34);
        _nw153[(new BigNumber(14)).toNumber()] = _module.__default.fm36(globalState);
        _nw153[(new BigNumber(15)).toNumber()] = new BigNumber(-535);
        _nw153[(new BigNumber(16)).toNumber()] = new BigNumber((_982_v40).length);
        _nw153[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_970_v32).length));
        _nw153[(new BigNumber(18)).toNumber()] = _977_cf31;
        _nw153[(new BigNumber(19)).toNumber()] = (_this).f34;
        _nw153[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt(_this.f29, new BigNumber((_944_v11).length));
        _nw153[(new BigNumber(21)).toNumber()] = _977_cf31;
        _nw153[(new BigNumber(22)).toNumber()] = (_this).f34;
        _nw153[(new BigNumber(23)).toNumber()] = new BigNumber((_984_v42).length);
        _nw153[(new BigNumber(24)).toNumber()] = new BigNumber((_module.__default.fm45(new BigNumber(877), p0, _this.f29, (((_986_v44).contains(_985_v43)) ? ((_986_v44).get(_985_v43)) : (!(p0))), globalState)).cardinality());
        _987_v45 = _nw153;
        let _index150 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_987_v45).length));
        (_987_v45)[_index150] = _977_cf31;
        let _index151 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_987_v45).length));
        (_987_v45)[_index151] = _this.f29;
        (globalState).f15 = (_this).f28;
        let _index152 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_987_v45).length));
        (_987_v45)[_index152] = (_this.f29).minus((_this).f34);
        (globalState).f15 = (_this).f28;
      } else if (_source13.is_DC19) {
        (globalState).f5 = _module.__default.fm36(globalState);
        (globalState).f15 = false;
        let _988_v46;
        _988_v46 = _dafny.MultiSet.fromElements(_970_v32, _970_v32);
        let _989_v47;
        _989_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_988_v46).cardinality()),(_this).f34);
        let _990_v48;
        _990_v48 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("uk"));
        let _991_v49;
        _991_v49 = new _dafny.CodePoint('l'.codePointAt(0));
        _989_v47 = _module.__default.fm46(_dafny.Seq.IsProperPrefixOf(_990_v48, _dafny.Seq.of(_970_v32, _module.__default.fm18(p0, new BigNumber(22), _991_v49, globalState))), globalState);
        let _index153 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_969_v31).length));
        (_969_v31)[_index153] = (_969_v31)[_module.__default.safeIndex(new BigNumber(427), new BigNumber((_969_v31).length))];
      } else if (_source13.is_DC17) {
        let _992___mcc_h2 = (_source13).cf30;
        let _993_cf30 = _992___mcc_h2;
        (globalState).f25 = _module.__default.fm42((_this.f29).minus(new BigNumber((_970_v32).length)), globalState);
        (globalState).f17 = _module.__default.fm40(globalState);
        if (!_dafny.areEqual(_970_v32, _970_v32)) {
          let _994_v50;
          _994_v50 = _dafny.Set.fromElements((_this).f28, p0, _module.__default.fm40(globalState));
          let _995_v51;
          _995_v51 = _dafny.Set.fromElements((_994_v50).IsDisjointFrom(_994_v50));
          _995_v51 = _995_v51;
          (globalState).f17 = _module.__default.fm16(_944_v11, globalState);
          (globalState).f15 = (_this).f28;
          (globalState).f5 = ((p0) ? (_this.f29) : (_this.f29));
          let _996_v52;
          let _nw154 = Array((new BigNumber(24)).toNumber()).fill(_module.D9.Default());
          _996_v52 = _nw154;
          let _997_v53;
          _997_v53 = new _dafny.CodePoint('a'.codePointAt(0));
          let _998_v54;
          _998_v54 = _dafny.Set.fromElements(_997_v53);
          let _999_v55;
          _999_v55 = _module.D1.create_DC6(new BigNumber((_998_v54).length));
          let _1000_v56;
          _1000_v56 = _dafny.Map.Empty.slice().updateUnsafe(_996_v52,_999_v55);
          let _1001_v57;
          _1001_v57 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(597)).multipliedBy((_dafny.ZERO).minus((_this).f34)),(_970_v32)[_module.__default.safeIndex(new BigNumber((_1000_v56).length), new BigNumber((_970_v32).length))]);
          _1001_v57 = (_1001_v57).update((_this).f34, (_970_v32)[_module.__default.safeIndex(_this.f29, new BigNumber((_970_v32).length))]);
        } else {
          let _1002_v58;
          _1002_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_970_v32);
          _970_v32 = _dafny.Seq.Concat(_970_v32, _dafny.Seq.Concat((((_1002_v58).contains(new BigNumber((_968_v30).length))) ? ((_1002_v58).get(new BigNumber((_968_v30).length))) : (_970_v32)), _970_v32));
          let _1003_v59;
          _1003_v59 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f34,false), _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f28));
          let _1004_v60;
          _1004_v60 = _dafny.Seq.of(_970_v32);
          let _1005_v61;
          _1005_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1003_v59)[_module.__default.safeIndex((_this).f34, new BigNumber((_1003_v59).length))],_1004_v60);
          _1005_v61 = (_1005_v61).update((_928_v0).Merge(_928_v0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-231)), ((_1006_v32) => function (_1007_i5) {
            return _1006_v32;
          })(_970_v32)));
          _930_v2 = _930_v2;
          (globalState).f13 = _this.f29;
          (globalState).f21 = (_971_v33)[_module.__default.safeIndex(_this.f29, new BigNumber((_971_v33).length))];
        }
        let _1008_v62;
        _1008_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f29);
        let _1009_v63;
        let _nw155 = new _module.C0();
        _nw155.__ctor(_970_v32, ((_this).f34).multipliedBy(new BigNumber((_1008_v62).length)), (_this).f27, p0);
        _1009_v63 = _nw155;
      } else {
        let _1010___mcc_h3 = (_source13).cf32;
        let _1011_cf32 = _1010___mcc_h3;
        let _1012_v64;
        let _nw156 = new _module.C1();
        _nw156.__ctor(p0, (_this).f27, (_this).f28);
        _1012_v64 = _nw156;
        let _1013_v65;
        _1013_v65 = _dafny.Seq.of(_1012_v64);
        (_this).f29 = ((!(_dafny.Seq.contains(_1013_v65, _1012_v64))) ? (_module.__default.safeModuloInt(_this.f29, _this.f29)) : (new BigNumber((_dafny.Seq.Concat((_969_v31)[_module.__default.safeIndex(new BigNumber(427), new BigNumber((_969_v31).length))], _968_v30)).length)));
        (globalState).f15 = (_1012_v64).f39;
        (globalState).f15 = p0;
        let _1014_v66;
        let _nw157 = new _module.C0();
        _nw157.__ctor(_970_v32, (_this).f34, (_this).f27, (_this).f28);
        _1014_v66 = _nw157;
        let _1015_v67;
        _1015_v67 = _dafny.Seq.of(_1014_v66);
        let _1016_v68;
        _1016_v68 = _dafny.Set.fromElements((_1015_v67)[_module.__default.safeIndex((_this).f34, new BigNumber((_1015_v67).length))], _1014_v66);
        (globalState).f15 = (_dafny.Set.fromElements(_1014_v66)).IsSubsetOf(_1016_v68);
      }
      let _1017_v69;
      _1017_v69 = _dafny.Set.fromElements(_this.f29);
      r0 = new BigNumber((_1017_v69).length);
      let _1018_v70;
      _1018_v70 = new _dafny.CodePoint('r'.codePointAt(0));
      let _1019_v71;
      _1019_v71 = _module.D1.create_DC7(_1018_v70, _970_v32);
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_970_v32, _dafny.Seq.UnicodeFromString("dyjwr")), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1019_v71,(_this).f27)).length), new BigNumber((_dafny.Seq.Concat(_970_v32, _dafny.Seq.UnicodeFromString("dyjwr"))).length)), _1018_v70);
      r2 = (_this).f28;
      return [r0, r1, r2];
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm12(p0, p1, p2, globalState) {
      let _this = this;
      let _source15 = _module.D1.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(187), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(55)), function (_1020_i0) {
  return new BigNumber(-962);
})).length))).cardinality()));
      if (_source15.is_DC6) {
        let _1021___mcc_h0 = (_source15).cf12;
        let _1022_cf12 = _1021___mcc_h0;
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)))).Intersect(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,!(false)), _dafny.Map.Empty.slice().updateUnsafe(true,false)));
      } else if (_source15.is_DC7) {
        let _1023___mcc_h1 = (_source15).cf13;
        let _1024___mcc_h2 = (_source15).cf14;
        let _1025_cf14 = _1024___mcc_h2;
        let _1026_cf13 = _1023___mcc_h1;
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(false,!(!(true))));
      } else {
        let _1027___mcc_h3 = (_source15).cf11;
        let _1028_cf11 = _1027___mcc_h3;
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,true));
      }
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.of(true);
    };
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Set.Empty;
      let r3 = _dafny.Seq.of();
      let _1029_v0;
      _1029_v0 = _dafny.Seq.UnicodeFromString("xxd");
      let _1030_v1;
      _1030_v1 = true;
      let _1031_v2;
      let _nw158 = Array((new BigNumber(9)).toNumber());
      _nw158[(_dafny.ZERO).toNumber()] = _1030_v1;
      _nw158[(_dafny.ONE).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(2)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(3)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(4)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(5)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(6)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(7)).toNumber()] = _1030_v1;
      _nw158[(new BigNumber(8)).toNumber()] = _1030_v1;
      _1031_v2 = _nw158;
      let _1032_v3;
      let _nw159 = new _module.C0();
      _nw159.__ctor(_1029_v0, p1, _1031_v2, _1030_v1);
      _1032_v3 = _nw159;
      let _1033_v4;
      _1033_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1032_v3,p1);
      let _1034_v5;
      _1034_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v4,(_1032_v3).f28);
      let _1035_v6;
      _1035_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1032_v3.f29,_module.__default.fm36(globalState));
      let _1036_v7;
      _1036_v7 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1037_v8;
      _1037_v8 = _module.D0.create_DC0(_1034_v5, false, new BigNumber((_1035_v6).length), _1036_v7, !(_1030_v1));
      r0 = (_1037_v8).dtor_cf2;
      let _1038_v9;
      let _nw160 = new _module.C1();
      _nw160.__ctor((((_1032_v3).f28) ? (_1030_v1) : (_1030_v1)), (_1032_v3).f27, _1030_v1);
      _1038_v9 = _nw160;
      let _hi3 = (((_1038_v9).f39) ? (_1032_v3.f29) : ((p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]));
      for (let _1039_i0 = new BigNumber((_dafny.Seq.of(new BigNumber((_1029_v0).length), p1, _1032_v3.f29)).length); _1039_i0.isLessThan(_hi3); _1039_i0 = _1039_i0.plus(_dafny.ONE)) {
        let _1040_v10;
        _1040_v10 = _module.D0.create_DC3(_module.__default.fm36(globalState), (_1038_v9).f39, (_1032_v3).f28);
        let _index154 = _module.__default.safeIndex(new BigNumber(799), new BigNumber(((_1032_v3).f27).length));
        ((_1032_v3).f27)[_index154] = (_1040_v10).dtor_cf8;
        let _index155 = _module.__default.safeIndex(new BigNumber(799), new BigNumber(((_1032_v3).f27).length));
        ((_1032_v3).f27)[_index155] = !(_1030_v1);
        let _1041_v11;
        _1041_v11 = _dafny.Set.fromElements((_1038_v9).f39);
        let _1042_v12;
        _1042_v12 = _dafny.MultiSet.fromElements(((((_1032_v3).f27)[_module.__default.safeIndex(new BigNumber(799), new BigNumber(((_1032_v3).f27).length))]) ? (_dafny.Set.fromElements(_1030_v1, ((_1032_v3).f27)[_module.__default.safeIndex(new BigNumber(799), new BigNumber(((_1032_v3).f27).length))])) : (_1041_v11)));
        _1042_v12 = (((_module.D0.create_DC3(_1032_v3.f29, true, false)).dtor_cf9) ? (_1042_v12) : (_1042_v12));
        (globalState).f25 = (_1029_v0)[_module.__default.safeIndex(_module.__default.fm36(globalState), new BigNumber((_1029_v0).length))];
        let _index156 = _module.__default.safeIndex(new BigNumber(799), new BigNumber(((_1032_v3).f27).length));
        ((_1032_v3).f27)[_index156] = ((_1041_v11).Intersect(_1041_v11)).IsProperSubsetOf(_1041_v11);
      }
      let _1043_i1;
      _1043_i1 = _dafny.ZERO;
      L3: {
        while ((_1032_v3).f28) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1043_i1)) {
              break L3;
            }
            _1043_i1 = (_1043_i1).plus(_dafny.ONE);
            let _1044_v13;
            _1044_v13 = _dafny.MultiSet.fromElements(!(!((_1032_v3).f28)), (_1038_v9).f39);
            let _1045_v14;
            _1045_v14 = _dafny.Seq.of(_1044_v13);
            let _1046_v15;
            _1046_v15 = _module.D9.create_DC23((_1038_v9).f39);
            let _1047_v16;
            _1047_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1038_v9).f39,_1046_v15);
            let _1048_v17;
            _1048_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1032_v3.f29,_1047_v16);
            let _1049_v18;
            _1049_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_1048_v17).contains((_dafny.ZERO).minus(p1))) ? ((_1048_v17).get((_dafny.ZERO).minus(p1))) : (_1047_v16)),_1030_v1);
            let _1050_v20;
            let _init23 = ((_1051_v3, _1052_p1, _1053_v1) => function (_1054_i2) {
              return (_dafny.Set.fromElements(_1051_v3.f29, _1052_p1, _1051_v3.f29)).Intersect(function () {
                let _coll50 = new _dafny.Set();
                for (const _compr_50 of _dafny.IntegerRange(new BigNumber(-97), new BigNumber(757))) {
                  let _1055_v19 = _compr_50;
                  if (((new BigNumber(-97)).isLessThanOrEqualTo(_1055_v19)) && ((_1055_v19).isLessThan(new BigNumber(757)))) {
                    _coll50.add((_1055_v19).plus(new BigNumber((_dafny.MultiSet.fromElements(_1053_v1)).cardinality())));
                  }
                }
                return _coll50;
              }());
            })(_1032_v3, p1, _1030_v1);
            let _nw161 = Array((new BigNumber(16)).toNumber());
            for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw161.length); _i0_23++) {
              _nw161[_i0_23] = _init23(new BigNumber(_i0_23));
            }
            _1050_v20 = _nw161;
            let _1056_v21;
            _1056_v21 = _dafny.Set.fromElements(new BigNumber((_1029_v0).length));
            let _index157 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1050_v20).length));
            (_1050_v20)[_index157] = (_1056_v21).Intersect(_1056_v21);
            let _1057_v22;
            _1057_v22 = _module.D0.create_DC3(_module.__default.fm36(globalState), (_1038_v9).f39, true);
            let _index158 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1050_v20).length));
            let _rhs171 = _1032_v3.f29;
            let _rhs172 = _1045_v14;
            let _rhs173 = _dafny.Map.Empty.slice().updateUnsafe(_1047_v16,!(_1032_v3.f29).isEqualTo(_1032_v3.f29));
            let _rhs174 = _1056_v21;
            let _rhs175 = (_1057_v22).dtor_cf8;
            let _lhs146 = globalState;
            let _lhs147 = _1050_v20;
            let _lhs148 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1050_v20).length));
            let _lhs149 = globalState;
            _lhs146.f13 = _rhs171;
            _1045_v14 = _rhs172;
            _1049_v18 = _rhs173;
            _lhs147[_lhs148] = _rhs174;
            _lhs149.f17 = _rhs175;
            let _1058_v23;
            let _nw162 = Array((new BigNumber(10)).toNumber()).fill(_module.D8.Default());
            _1058_v23 = _nw162;
            let _1059_v24;
            _1059_v24 = _dafny.MultiSet.fromElements(_1058_v23);
            _1059_v24 = _1059_v24;
            _1057_v22 = _module.D0.create_DC3(_1032_v3.f29, (_1038_v9).f39, (_1032_v3).f28);
            (globalState).f4 = (p1).plus((p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]);
          }
        }
      }
      (globalState).f5 = _1032_v3.f29;
      let _1060_v25;
      let _nw163 = Array((new BigNumber(6)).toNumber()).fill([]);
      _1060_v25 = _nw163;
      let _1061_v26;
      _1061_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(417),_1036_v7);
      let _1062_v27;
      _1062_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1032_v3).f28,(_1038_v9).f39);
      let _1063_v28;
      let _nw164 = Array((new BigNumber(16)).toNumber());
      _nw164[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_1036_v7);
      _nw164[(_dafny.ONE).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(2)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(3)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(4)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(5)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(6)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(7)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(8)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(9)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(10)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(11)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(12)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(13)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(14)).toNumber()] = _1061_v26;
      _nw164[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1062_v27).length),_1036_v7);
      _1063_v28 = _nw164;
      let _index159 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1060_v25).length));
      (_1060_v25)[_index159] = _1063_v28;
      let _index160 = _module.__default.safeIndex(new BigNumber(133), new BigNumber(((_1032_v3).f27).length));
      ((_1032_v3).f27)[_index160] = _1030_v1;
      let _1064_v29;
      _1064_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1030_v1,_1029_v0);
      let _index161 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1060_v25).length));
      let _index162 = _module.__default.safeIndex(new BigNumber(133), new BigNumber(((_1032_v3).f27).length));
      let _rhs176 = _1063_v28;
      let _rhs177 = (p0)[_module.__default.safeIndex(_module.__default.safeDivisionInt(p1, new BigNumber(((((_1064_v29).contains((_1032_v3).f28)) ? ((_1064_v29).get((_1032_v3).f28)) : (_1029_v0))).length)), new BigNumber((p0).length))];
      let _rhs178 = (_1032_v3).f28;
      let _lhs150 = _1060_v25;
      let _lhs151 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1060_v25).length));
      let _lhs152 = globalState;
      let _lhs153 = (_1032_v3).f27;
      let _lhs154 = _module.__default.safeIndex(new BigNumber(133), new BigNumber(((_1032_v3).f27).length));
      _lhs150[_lhs151] = _rhs176;
      _lhs152.f13 = _rhs177;
      _lhs153[_lhs154] = _rhs178;
      r0 = p1;
      r1 = (_1038_v9).f39;
      r2 = function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-443),(_1038_v9).f39)).update(new BigNumber((_1029_v0).length), ((_1032_v3).f27)[_module.__default.safeIndex(new BigNumber(133), new BigNumber(((_1032_v3).f27).length))])).Keys.Elements) {
          let _1065_v30 = _compr_51;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-443),(_1038_v9).f39)).update(new BigNumber((_1029_v0).length), ((_1032_v3).f27)[_module.__default.safeIndex(new BigNumber(133), new BigNumber(((_1032_v3).f27).length))])).contains(_1065_v30)) {
            _coll51.add((_1065_v30).multipliedBy((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, false, true, !(false), true)).length))).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0))))).cardinality())))));
          }
        }
        return _coll51;
      }();
      r3 = _dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(_1032_v3.f29, new BigNumber((p0).length)), p1), p0);
      return [r0, r1, r2, r3];
    }
    m9(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1066_v0;
      _1066_v0 = _dafny.Seq.UnicodeFromString("islesyilu");
      let _1067_v1;
      _1067_v1 = _module.D8.create_DC18(p3);
      let _1068_v2;
      let _init24 = ((_1069_p2) => function (_1070_i0) {
        return _1069_p2;
      })(p2);
      let _nw165 = Array((new BigNumber(6)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw165.length); _i0_24++) {
        _nw165[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _1068_v2 = _nw165;
      let _1071_v3;
      let _nw166 = new _module.C0();
      _nw166.__ctor(_1066_v0, (_1067_v1).dtor_cf31, _1068_v2, false);
      _1071_v3 = _nw166;
      let _1072_v5;
      _1072_v5 = _dafny.Set.fromElements(p0, p3);
      if (((_dafny.Seq.IsPrefixOf((_1071_v3).f36, _1066_v0)) ? ((_1072_v5).IsProperSubsetOf(function () {
        let _coll52 = new _dafny.Set();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(623), new BigNumber(207))) {
          let _1073_v4 = _compr_52;
          if (((new BigNumber(623)).isLessThanOrEqualTo(_1073_v4)) && ((_1073_v4).isLessThan(new BigNumber(207)))) {
            _coll52.add(_module.__default.safeDivisionInt(_1073_v4, p3));
          }
        }
        return _coll52;
      }())) : (!(p2)))) {
        let _1074_v6;
        _1074_v6 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1075_v7;
        _1075_v7 = _module.D1.create_DC7(_1074_v6, (_1071_v3).f36);
        let _source16 = ((p2) ? (_1075_v7) : (_1075_v7));
        if (_source16.is_DC6) {
          let _1076___mcc_h0 = (_source16).cf12;
          let _1077_cf12 = _1076___mcc_h0;
          let _1078_v8;
          let _nw167 = new _module.C1();
          _nw167.__ctor(p2, _1068_v2, p2);
          _1078_v8 = _nw167;
          let _1079_v9;
          _1079_v9 = _module.D1.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), ((_1080_cf12, _1081_p3, _1082_v3, _1083_p0) => function (_1084_i1) {
  return _dafny.Seq.of(_1080_cf12, _1080_cf12, _1081_p3, new BigNumber(((_1082_v3).f36).length), _1083_p0);
})(_1077_cf12, p3, _1071_v3, p0)));
          let _1085_v10;
          _1085_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(p3));
          let _1086_v11;
          _1086_v11 = _dafny.MultiSet.fromElements(p3);
          _1079_v9 = _module.__default.fm47(_1085_v10, _1086_v11, p3, globalState);
          (globalState).f15 = (_1078_v8).f39;
          (globalState).f25 = _1074_v6;
        } else if (_source16.is_DC7) {
          let _1087___mcc_h1 = (_source16).cf13;
          let _1088___mcc_h2 = (_source16).cf14;
          let _1089_cf14 = _1088___mcc_h2;
          let _1090_cf13 = _1087___mcc_h1;
          let _1091_v12;
          let _nw168 = new _module.C1();
          _nw168.__ctor(false, _1068_v2, p2);
          _1091_v12 = _nw168;
          let _1092_v13;
          _1092_v13 = _dafny.Set.fromElements(_1074_v6, _module.__default.fm42(p0, globalState));
          let _1093_v14;
          let _nw169 = new _module.C0();
          _nw169.__ctor(_1066_v0, (new BigNumber((_dafny.Seq.of(p2)).length)).minus(new BigNumber((_1092_v13).length)), _1068_v2, p2);
          _1093_v14 = _nw169;
          let _1094_v15;
          _1094_v15 = _dafny.Seq.of(p3);
          let _index163 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
          (p1)[_index163] = (_1094_v15)[_module.__default.safeIndex(p0, new BigNumber((_1094_v15).length))];
          let _1095_v16;
          _1095_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),p0);
          let _1096_v17;
          _1096_v17 = _dafny.Seq.of((_1095_v16).Merge(_module.__default.fm46(!(p2), globalState)), _1095_v16);
          let _index164 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((p1).length));
          (p1)[_index164] = p3;
          let _index165 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
          let _index166 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((p1).length));
          let _rhs179 = _module.__default.safeModuloInt(p0, p0);
          let _rhs180 = _dafny.Seq.Concat(_1096_v17, _dafny.Seq.Concat(_1096_v17, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,p3), _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-711)))));
          let _rhs181 = p3;
          let _lhs155 = p1;
          let _lhs156 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
          let _lhs157 = p1;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((p1).length));
          _lhs155[_lhs156] = _rhs179;
          _1096_v17 = _rhs180;
          _lhs157[_lhs158] = _rhs181;
          _1090_cf13 = _1074_v6;
        } else {
          let _1097___mcc_h3 = (_source16).cf11;
          let _1098_cf11 = _1097___mcc_h3;
          let _1099_v18;
          _1099_v18 = _dafny.MultiSet.fromElements(p2, p2, p2, p2, (p3).isLessThan(p3));
          _1099_v18 = ((_1099_v18).Difference(_1099_v18)).Difference(_1099_v18);
          (globalState).f15 = (_1099_v18).IsSubsetOf((_1099_v18).Intersect(_1099_v18));
          _1066_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat((_1071_v3).f36, (_1071_v3).f36), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ppri"), _dafny.Seq.UnicodeFromString("clfofo")));
          let _1100_v19;
          _1100_v19 = _dafny.Set.fromElements(p2, p2);
          let _1101_v20;
          _1101_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1100_v19).Intersect(_1100_v19),false);
          let _1102_v21;
          _1102_v21 = _dafny.Seq.of(p2);
          _1101_v20 = (_1101_v20).update((_1100_v19).Union(_dafny.Set.fromElements(p2)), !((_1102_v21)[_module.__default.safeIndex(p0, new BigNumber((_1102_v21).length))]));
        }
        let _1103_v22;
        _1103_v22 = _dafny.Seq.of(p0);
        let _1104_v23;
        _1104_v23 = _dafny.Set.fromElements(p2);
        let _1105_v24;
        _1105_v24 = _dafny.MultiSet.fromElements((_1103_v22)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1104_v23)).length), new BigNumber((_1103_v22).length))]);
        let _1106_v25;
        _1106_v25 = _dafny.Map.Empty.slice().updateUnsafe(((_1105_v24).update(p0, _module.__default.abs(p3))).update(p3, _module.__default.abs(p0)),_1074_v6);
        let _1107_v27;
        _1107_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1105_v24,new BigNumber((_dafny.Seq.UnicodeFromString("wydr")).length));
        _1106_v25 = (_1106_v25).Merge(function () {
          let _coll53 = new _dafny.Map();
          for (const _compr_53 of (_1107_v27).Keys.Elements) {
            let _1108_v26 = _compr_53;
            if ((_1107_v27).contains(_1108_v26)) {
              _coll53.push([_1108_v26,_1074_v6]);
            }
          }
          return _coll53;
        }());
        let _1109_v28;
        _1109_v28 = _dafny.Seq.of(p2);
        let _1110_v29;
        _1110_v29 = _module.D0.create_DC3(new BigNumber((_1109_v28).length), !(p2), p2);
        let _1111_v30;
        _1111_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1110_v29,p0);
        let _1112_v31;
        _1112_v31 = _1068_v2;
        let _1113_v32;
        let _nw170 = new _module.C3();
        _nw170.__ctor(p0, (_1111_v30).Merge(_1111_v30), (_dafny.ZERO).minus((p0).plus(p3)), (_1068_v2), p2);
        _1113_v32 = _nw170;
        let _index167 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((p1).length));
        (p1)[_index167] = (new BigNumber((_1104_v23).length)).minus(p3);
        let _index168 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((p1).length));
        (p1)[_index168] = (_1113_v32).f34;
        let _1114_v33;
        _1114_v33 = _module.D1.create_DC6(p0);
        let _index169 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((p1).length));
        (p1)[_index169] = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_1114_v33).dtor_cf12), _1103_v22)).length)).minus((_1113_v32).f34);
      } else {
        let _1115_v34;
        _1115_v34 = _module.D8.create_DC18((_dafny.ZERO).minus(p3));
        let _1116_v35;
        _1116_v35 = _module.D8.create_DC20(_1115_v34);
        let _source17 = _1116_v35;
        if (_source17.is_DC18) {
          let _1117___mcc_h4 = (_source17).cf31;
          let _1118_cf31 = _1117___mcc_h4;
          let _1119_v36;
          _1119_v36 = _dafny.Seq.of(p2, !(!(p2)), !(p2), p2, p2);
          let _1120_v37;
          _1120_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1071_v3).f36).length),p2);
          (globalState).f6 = ((p2) ? ((_1119_v36)[_module.__default.safeIndex((_dafny.ZERO).minus(p3), new BigNumber((_1119_v36).length))]) : (!((((_1120_v37).contains(new BigNumber(635))) ? ((_1120_v37).get(new BigNumber(635))) : (p2)))));
          let _1121_v38;
          _1121_v38 = _dafny.Seq.of(_1072_v5, _1072_v5, _1072_v5);
          let _1122_v39;
          _1122_v39 = _dafny.MultiSet.fromElements(p0, p0, p3);
          (globalState).f17 = ((_1121_v38)[_module.__default.safeIndex(new BigNumber((_1122_v39).cardinality()), new BigNumber((_1121_v38).length))]).IsSubsetOf(_dafny.Set.fromElements(_1118_cf31));
          let _1123_v40;
          let _1124_v41;
          let _out21;
          let _out22;
          let _outcollector8 = _module.__default.m0(globalState);
          _out21 = _outcollector8[0];
          _out22 = _outcollector8[1];
          _1123_v40 = _out21;
          _1124_v41 = _out22;
          let _1125_v42;
          _1125_v42 = _dafny.Seq.of(p0);
          let _1126_v43;
          _1126_v43 = _dafny.Seq.of(_1125_v42, _1125_v42, _1125_v42, _1125_v42, _1125_v42);
          let _1127_v44;
          _1127_v44 = (_1126_v43)[_module.__default.safeIndex(p3, new BigNumber((_1126_v43).length))];
          let _1128_v45;
          _1128_v45 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1129_v46;
          _1129_v46 = _dafny.Set.fromElements(_module.__default.fm26(p2, _1118_cf31, _1128_v45, globalState));
          let _1130_v47;
          _1130_v47 = _module.D0.create_DC3((_1071_v3).fm15(globalState), p2, p2);
          let _rhs182 = p0;
          let _rhs183 = p2;
          let _rhs184 = ((_1129_v46).Union(_1129_v46)).IsProperSubsetOf(_1129_v46);
          let _rhs185 = _1127_v44;
          let _rhs186 = (_1130_v47).dtor_cf9;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          let _lhs162 = globalState;
          _lhs159.f5 = _rhs182;
          _lhs160.f6 = _rhs183;
          _lhs161.f6 = _rhs184;
          _1127_v44 = _rhs185;
          _lhs162.f15 = _rhs186;
        } else if (_source17.is_DC19) {
          let _1131_v48;
          _1131_v48 = new _dafny.CodePoint('r'.codePointAt(0));
          _1066_v0 = _dafny.Seq.update((_1071_v3).f36, _module.__default.safeIndex(_module.__default.fm36(globalState), new BigNumber(((_1071_v3).f36).length)), _1131_v48);
          let _1132_v50;
          _1132_v50 = _dafny.Seq.of(_1072_v5);
          let _rhs187 = _module.__default.fm32(p0, p2, globalState);
          let _rhs188 = p2;
          let _rhs189 = !((function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of ((_1132_v50)[_module.__default.safeIndex(new BigNumber((_1072_v5).length), new BigNumber((_1132_v50).length))]).Elements) {
              let _1133_v49 = _compr_54;
              if (((_1132_v50)[_module.__default.safeIndex(new BigNumber((_1072_v5).length), new BigNumber((_1132_v50).length))]).contains(_1133_v49)) {
                _coll54.push([(_1133_v49).plus(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)),p0]);
              }
            }
            return _coll54;
          }()).equals((function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of _dafny.IntegerRange(new BigNumber(615), new BigNumber(805))) {
              let _1134_v51 = _compr_55;
              if (((new BigNumber(615)).isLessThanOrEqualTo(_1134_v51)) && ((_1134_v51).isLessThan(new BigNumber(805)))) {
                _coll55.push([_module.__default.safeDivisionInt(_1134_v51, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(37)), ((_1135_v48) => function (_1136_i2) {
                  return _1135_v48;
                })(_1131_v48))).length))),new BigNumber(-578)]);
              }
            }
            return _coll55;
          }()).Merge(function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of _dafny.IntegerRange(new BigNumber(787), new BigNumber(-73))) {
              let _1137_v52 = _compr_56;
              if (((new BigNumber(787)).isLessThanOrEqualTo(_1137_v52)) && ((_1137_v52).isLessThan(new BigNumber(-73)))) {
                _coll56.push([_module.__default.safeModuloInt(_1137_v52, (_dafny.ZERO).minus(p3)),p3]);
              }
            }
            return _coll56;
          }())));
          let _lhs163 = globalState;
          let _lhs164 = globalState;
          _1066_v0 = _rhs187;
          _lhs163.f17 = _rhs188;
          _lhs164.f15 = _rhs189;
          let _1138_v53;
          _1138_v53 = _dafny.MultiSet.fromElements(p2, p2);
          let _1139_v54;
          _1139_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1138_v53).cardinality()));
          let _1140_v56;
          let _nw171 = new _module.C1();
          _nw171.__ctor((_1072_v5).IsProperSubsetOf(function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_1139_v54).Keys.Elements) {
              let _1141_v55 = _compr_57;
              if ((_1139_v54).contains(_1141_v55)) {
                _coll57.add((_1141_v55).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), function (_1142_i3) {
                  return new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality());
                })).length)));
              }
            }
            return _coll57;
          }()), _1068_v2, p2);
          _1140_v56 = _nw171;
          let _rhs190 = _1140_v56;
          let _rhs191 = p2;
          let _lhs165 = globalState;
          _1140_v56 = _rhs190;
          _lhs165.f6 = _rhs191;
          (globalState).f5 = _module.__default.safeModuloInt(_module.__default.fm36(globalState), _module.__default.safeDivisionInt((_1071_v3).fm15(globalState), p0));
        } else if (_source17.is_DC17) {
          let _1143___mcc_h5 = (_source17).cf30;
          let _1144_cf30 = _1143___mcc_h5;
          let _1145_v57;
          _1145_v57 = _dafny.MultiSet.fromElements(_module.D7.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1072_v5).length),_dafny.Seq.UnicodeFromString("sgx"))));
          let _1146_v58;
          _1146_v58 = _module.D8.create_DC17(_1144_cf30);
          (globalState).f4 = (((_1145_v57).contains(_module.__default.fm48(p3, p2, p0, _1146_v58, globalState))) ? ((_1145_v57).get(_module.__default.fm48(p3, p2, p0, _1146_v58, globalState))) : (p0));
          (globalState).f4 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p3, p3), p0);
          (globalState).f5 = _module.__default.safeDivisionInt((_module.__default.fm36(globalState)).minus(p3), _module.__default.safeModuloInt(new BigNumber(625), new BigNumber(647)));
          let _1147_v59;
          let _init25 = function (_1148_i4) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          };
          let _nw172 = Array((new BigNumber(23)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw172.length); _i0_25++) {
            _nw172[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1147_v59 = _nw172;
          let _1149_v60;
          _1149_v60 = new _dafny.CodePoint('u'.codePointAt(0));
          let _index170 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1147_v59).length));
          (_1147_v59)[_index170] = _1149_v60;
          let _1150_v61;
          _1150_v61 = _dafny.Seq.of(new BigNumber(145));
          let _1151_v62;
          _1151_v62 = _dafny.MultiSet.fromElements(((_1150_v61)[_module.__default.safeIndex(p3, new BigNumber((_1150_v61).length))]).isLessThan(new BigNumber(((_1071_v3).f36).length)));
          let _1152_v63;
          _1152_v63 = _dafny.Seq.of(p2);
          let _1153_v64;
          _1153_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1154_v65;
          _1154_v65 = _dafny.Seq.of(_1150_v61);
          let _1155_v66;
          _1155_v66 = _module.D0.create_DC3(new BigNumber((_1150_v61).length), p2, p2);
          let _index171 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1147_v59).length));
          let _rhs192 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_1152_v63, _1152_v63)).length), p0);
          let _rhs193 = (((_1153_v64).contains(_module.__default.fm36(globalState))) ? ((_1153_v64).get(_module.__default.fm36(globalState))) : (new BigNumber((_dafny.Seq.Concat((_1154_v65)[_module.__default.safeIndex(p0, new BigNumber((_1154_v65).length))], _dafny.Seq.of(p0))).length)));
          let _rhs194 = ((p2) ? (_1149_v60) : (_1149_v60));
          let _rhs195 = (_1151_v62).Union(_dafny.MultiSet.fromElements(p2));
          let _rhs196 = _module.__default.safeModuloInt((new BigNumber((_module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe(p3,_1072_v5), p2, (_1155_v66).dtor_cf8, globalState)).length)).plus(p3), p0);
          let _lhs166 = globalState;
          let _lhs167 = globalState;
          let _lhs168 = _1147_v59;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1147_v59).length));
          let _lhs170 = globalState;
          _lhs166.f13 = _rhs192;
          _lhs167.f21 = _rhs193;
          _lhs168[_lhs169] = _rhs194;
          _1151_v62 = _rhs195;
          _lhs170.f5 = _rhs196;
        } else {
          let _1156___mcc_h6 = (_source17).cf32;
          let _1157_cf32 = _1156___mcc_h6;
          _1068_v2 = _1068_v2;
          (globalState).f6 = !(p2);
          let _1158_v67;
          _1158_v67 = _dafny.Seq.of(p1);
          let _1159_v68;
          _1159_v68 = _dafny.Set.fromElements((_1158_v67)[_module.__default.safeIndex(p0, new BigNumber((_1158_v67).length))], p1);
          let _1160_v69;
          _1160_v69 = _module.D0.create_DC3(p0, !(p2), p2);
          let _1161_v70;
          _1161_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1160_v69,p3);
          let _1162_v71;
          let _nw173 = new _module.C3();
          _nw173.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("vfeldy")).length), _1161_v70, p0, _1068_v2, p2);
          _1162_v71 = _nw173;
          let _1163_v72;
          _1163_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1162_v71,_1162_v71.f29);
          let _1164_v73;
          _1164_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_1165_p0) => function (_1166_i5) {
            return _1165_p0;
          })(p0)))).cardinality()),p2);
          let _1167_v74;
          _1167_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v72,(((_1164_v73).contains((_dafny.ZERO).minus(_1162_v71.f29))) ? ((_1164_v73).get((_dafny.ZERO).minus(_1162_v71.f29))) : ((_1162_v71).f28)));
          let _1168_v75;
          _1168_v75 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1169_v76;
          _1169_v76 = _module.D0.create_DC0(_1167_v74, p2, _1162_v71.f29, _1168_v75, !(p2));
          let _index172 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_1068_v2).length));
          (_1068_v2)[_index172] = !((_1169_v76).dtor_cf1);
          let _index173 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_1068_v2).length));
          (_1068_v2)[_index173] = (_1162_v71).f28;
          let _index174 = _module.__default.safeIndex(new BigNumber(523), new BigNumber(((_1162_v71).f27).length));
          ((_1162_v71).f27)[_index174] = !(p2);
          let _1170_v77;
          _1170_v77 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
          let _index175 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_1068_v2).length));
          let _index176 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_1068_v2).length));
          let _index177 = _module.__default.safeIndex(new BigNumber(523), new BigNumber(((_1162_v71).f27).length));
          let _rhs197 = new BigNumber(720);
          let _rhs198 = ((false) ? (_1159_v68) : (_1159_v68));
          let _rhs199 = (_1162_v71).f28;
          let _rhs200 = _module.__default.fm40(globalState);
          let _rhs201 = (_module.__default.safeModuloInt(_1162_v71.f29, p3)).isLessThanOrEqualTo(new BigNumber(((_1170_v77).Merge(_1170_v77)).length));
          let _lhs171 = globalState;
          let _lhs172 = _1068_v2;
          let _lhs173 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_1068_v2).length));
          let _lhs174 = _1068_v2;
          let _lhs175 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_1068_v2).length));
          let _lhs176 = (_1162_v71).f27;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(523), new BigNumber(((_1162_v71).f27).length));
          _lhs171.f5 = _rhs197;
          _1159_v68 = _rhs198;
          _lhs172[_lhs173] = _rhs199;
          _lhs174[_lhs175] = _rhs200;
          _lhs176[_lhs177] = _rhs201;
          let _1171_v78;
          _1171_v78 = _dafny.MultiSet.fromElements((_1068_v2)[_module.__default.safeIndex(new BigNumber(582), new BigNumber((_1068_v2).length))]);
          _1171_v78 = _1171_v78;
        }
        (globalState).f15 = true;
        (globalState).f6 = !(false);
        _1066_v0 = _dafny.Seq.Concat(_1066_v0, _dafny.Seq.Concat((_1071_v3).f36, _1066_v0));
        let _1172_v79;
        _1172_v79 = _dafny.Seq.of(p2);
        (globalState).f15 = _dafny.Seq.contains(_1172_v79, p2);
      }
      let _index178 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
      (_1068_v2)[_index178] = p2;
      let _index179 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
      let _rhs202 = (p0).multipliedBy(p0);
      let _rhs203 = p0;
      let _rhs204 = _1068_v2;
      let _rhs205 = false;
      let _lhs178 = globalState;
      let _lhs179 = globalState;
      let _lhs180 = _1068_v2;
      let _lhs181 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
      _lhs178.f5 = _rhs202;
      _lhs179.f4 = _rhs203;
      _1068_v2 = _rhs204;
      _lhs180[_lhs181] = _rhs205;
      let _1173_v80;
      _1173_v80 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))],false),p0);
      let _1174_v81;
      _1174_v81 = _dafny.Map.Empty.slice().updateUnsafe((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))],!((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))]));
      let _hi4 = (_dafny.ZERO).minus(new BigNumber(-130));
      for (let _1175_i6 = (((_1173_v80).contains(_1174_v81)) ? ((_1173_v80).get(_1174_v81)) : (new BigNumber(((_1071_v3).f36).length))); _1175_i6.isLessThan(_hi4); _1175_i6 = _1175_i6.plus(_dafny.ONE)) {
        let _1176_v82;
        _1176_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1072_v5);
        let _1177_v83;
        let _nw174 = new _module.C2();
        _nw174.__ctor(new _dafny.CodePoint('t'.codePointAt(0)), p2, _1068_v2, !((_1072_v5).IsDisjointFrom((((_1176_v82).contains((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])) ? ((_1176_v82).get((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])) : (_1072_v5)))));
        _1177_v83 = _nw174;
        let _1178_v84;
        _1178_v84 = _dafny.Map.Empty.slice().updateUnsafe(true,p3);
        let _1179_v85;
        _1179_v85 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(_1178_v84));
        let _1180_v86;
        _1180_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1072_v5).length),(((_1179_v85).contains((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])) ? ((_1179_v85).get((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])) : (_dafny.Seq.of(_1178_v84, _1178_v84))));
        let _1181_v87;
        _1181_v87 = _dafny.Seq.of(_1178_v84);
        _1180_v86 = (_1180_v86).update(_1175_i6, _1181_v87);
        let _1182_v88;
        _1182_v88 = _dafny.Seq.of(_1177_v83.f38, _1177_v83.f38, _1177_v83.f38);
        let _1183_v89;
        _1183_v89 = _dafny.MultiSet.fromElements(_1177_v83.f38);
        if ((_1183_v89).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_1182_v88, _1182_v88), _module.__default.safeIndex(new BigNumber(((_1071_v3).f36).length), new BigNumber((_dafny.Seq.Concat(_1182_v88, _1182_v88)).length)), (_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])))) {
          let _1184_v90;
          _1184_v90 = _dafny.Seq.of(_1068_v2);
          let _index180 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
          let _index181 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
          let _rhs206 = !_dafny.areEqual(_1184_v90, _1184_v90);
          let _rhs207 = (new BigNumber(-605)).isEqualTo(p0);
          let _rhs208 = new BigNumber((_module.__default.fm18(p2, p3, (_1177_v83).f37, globalState)).length);
          let _lhs182 = _1068_v2;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
          let _lhs184 = _1068_v2;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length));
          let _lhs186 = globalState;
          _lhs182[_lhs183] = _rhs206;
          _lhs184[_lhs185] = _rhs207;
          _lhs186.f5 = _rhs208;
          let _1185_v91;
          _1185_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1178_v84,_1175_i6);
          let _1186_v92;
          _1186_v92 = _module.D8.create_DC17(_1185_v91);
          let _1187_v93;
          _1187_v93 = _module.D8.create_DC20(_1186_v92);
          let _rhs209 = new BigNumber(894);
          let _rhs210 = _1187_v93;
          let _lhs187 = globalState;
          _lhs187.f13 = _rhs209;
          _1187_v93 = _rhs210;
          (globalState).f4 = _module.__default.safeModuloInt(p0, p3);
          _1066_v0 = (_1071_v3).f36;
          let _1188_v94;
          _1188_v94 = _dafny.Seq.of(_1066_v0);
          let _1189_v95;
          _1189_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1188_v94, _1188_v94)).length),_1068_v2);
          _1068_v2 = (((_1189_v95).contains(_1175_i6)) ? ((_1189_v95).get(_1175_i6)) : (_1068_v2));
        } else {
          let _1190_v96;
          _1190_v96 = _dafny.Set.fromElements(true);
          _1190_v96 = _1190_v96;
          let _1191_v97;
          let _nw175 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1191_v97 = _nw175;
          let _index182 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1191_v97).length));
          (_1191_v97)[_index182] = (_1177_v83).f37;
          let _index183 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_1191_v97).length));
          (_1191_v97)[_index183] = (_1177_v83).f37;
          let _1192_v98;
          _1192_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1175_i6,_1190_v96);
          _1192_v98 = (_1192_v98).update((((_1178_v84).contains(_1177_v83.f38)) ? ((_1178_v84).get(_1177_v83.f38)) : (_1175_i6)), (_1190_v96).Difference(_1190_v96));
          (globalState).f21 = (_module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(p3)))).minus((_1177_v83).fm20(globalState));
          let _1193_v99;
          _1193_v99 = _dafny.Map.Empty.slice().updateUnsafe(_1178_v84,p0);
          let _1194_v100;
          _1194_v100 = _module.D8.create_DC20(_module.D8.create_DC17(_1193_v99));
          _1194_v100 = _1194_v100;
        }
        let _1195_v101;
        let _nw176 = Array((new BigNumber(14)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = (_1071_v3).f36;
        _nw176[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("meyjff"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_1196_v83) => function (_1197_i7) {
          return (_1196_v83).f37;
        })(_1177_v83)));
        _nw176[(new BigNumber(2)).toNumber()] = _1066_v0;
        _nw176[(new BigNumber(3)).toNumber()] = (_1071_v3).f36;
        _nw176[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kqcgi"), (_1071_v3).f36);
        _nw176[(new BigNumber(5)).toNumber()] = (_1071_v3).f36;
        _nw176[(new BigNumber(6)).toNumber()] = _dafny.Seq.update((_1071_v3).f36, _module.__default.safeIndex(p3, new BigNumber(((_1071_v3).f36).length)), new _dafny.CodePoint('t'.codePointAt(0)));
        _nw176[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("krlvulblo");
        _nw176[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("b");
        _nw176[(new BigNumber(9)).toNumber()] = (_1071_v3).f36;
        _nw176[(new BigNumber(10)).toNumber()] = _1066_v0;
        _nw176[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_1198_v83) => function (_1199_i8) {
          return (_1198_v83).f37;
        })(_1177_v83));
        _nw176[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1066_v0, (_1071_v3).f36);
        _nw176[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_1200_v83) => function (_1201_i9) {
          return (_1200_v83).f37;
        })(_1177_v83));
        _1195_v101 = _nw176;
        let _index184 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_1195_v101).length));
        (_1195_v101)[_index184] = (_1071_v3).f36;
        let _1202_v102;
        _1202_v102 = _module.D9.create_DC21(_1195_v101);
        let _pat_let_tv27 = _1202_v102;
        let _1203_v103;
        _1203_v103 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let19_0) {
          return function (_1204_dt__update__tmp_h1) {
            return function (_pat_let20_0) {
              return function (_1205_dt__update_hcf39_h0) {
                return _module.D9.create_DC24(_1205_dt__update_hcf39_h0);
              }(_pat_let20_0);
            }(_pat_let_tv27);
          }(_pat_let19_0);
        }(_module.D9.create_DC24(_1202_v102)),_dafny.Seq.UnicodeFromString("wgbvpofn"));
        let _1206_v104;
        _1206_v104 = _module.D9.create_DC24(_1202_v102);
        let _index185 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_1195_v101).length));
        (_1195_v101)[_index185] = _dafny.Seq.Concat((_1071_v3).f36, _dafny.Seq.update((((_1203_v103).contains(_1206_v104)) ? ((_1203_v103).get(_1206_v104)) : (_1066_v0)), _module.__default.safeIndex(new BigNumber(525), new BigNumber(((((_1203_v103).contains(_1206_v104)) ? ((_1203_v103).get(_1206_v104)) : (_1066_v0))).length)), (_1177_v83).f37));
      }
      let _1207_v105;
      _1207_v105 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p3,p2));
      (globalState).f5 = (new BigNumber((_1207_v105).length)).minus((p0).plus(p0));
      let _index186 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((p1).length));
      (p1)[_index186] = p3;
      let _index187 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((p1).length));
      (p1)[_index187] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))],(_1068_v2)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_1068_v2).length))])).length)));
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
      this._f33 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f33, f29, f27, f28) {
      let _this = this;
      (_this)._f33 = f33;
      (_this)._f29 = f29;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f29;
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_this.f29).multipliedBy(_this.f29);
    };
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1208_v0;
      let _init26 = function (_1209_i0) {
        return (_1209_i0).multipliedBy(_this.f29);
      };
      let _nw177 = Array((new BigNumber(8)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw177.length); _i0_26++) {
        _nw177[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1208_v0 = _nw177;
      let _1210_v1;
      _1210_v1 = _dafny.MultiSet.fromElements(p0, true);
      let _index188 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length));
      (_1208_v0)[_index188] = new BigNumber((_1210_v1).cardinality());
      let _1211_v2;
      _1211_v2 = _module.D3.create_DC10(p0, (_this).f28);
      let _1212_v3;
      _1212_v3 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!((_this).f28))).length));
      let _1213_v4;
      _1213_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1211_v2,_1212_v3);
      let _index189 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length));
      (_1208_v0)[_index189] = new BigNumber((((p0) ? (_1213_v4) : ((_1213_v4).Merge(_1213_v4)))).length);
      let _hi5 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f29, (_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]));
      for (let _1214_i1 = (_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]; _1214_i1.isLessThan(_hi5); _1214_i1 = _1214_i1.plus(_dafny.ONE)) {
        (globalState).f17 = p0;
        let _1215_v8;
        _1215_v8 = _dafny.MultiSet.fromElements(function () {
          let _coll58 = new _dafny.Set();
          for (const _compr_58 of _dafny.IntegerRange(new BigNumber(-224), new BigNumber(926))) {
            let _1216_v5 = _compr_58;
            if (((new BigNumber(-224)).isLessThanOrEqualTo(_1216_v5)) && ((_1216_v5).isLessThan(new BigNumber(926)))) {
              _coll58.add((_1216_v5).multipliedBy(new BigNumber((function () {
                let _coll59 = new _dafny.Set();
                for (const _compr_59 of (function () {
                  let _coll60 = new _dafny.Map();
                  for (const _compr_60 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_1214_i1)).Keys.Elements) {
                    let _1217_v6 = _compr_60;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_1214_i1)).contains(_1217_v6)) {
                      _coll60.push([_1217_v6,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), function (_1218_i2) {
                        return new _dafny.CodePoint('l'.codePointAt(0));
                      })).length)]);
                    }
                  }
                  return _coll60;
                }()).Keys.Elements) {
                  let _1219_v7 = _compr_59;
                  if ((function () {
                    let _coll61 = new _dafny.Map();
                    for (const _compr_61 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_1214_i1)).Keys.Elements) {
                      let _1217_v6 = _compr_61;
                      if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_1214_i1)).contains(_1217_v6)) {
                        _coll61.push([_1217_v6,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), function (_1218_i2) {
                          return new _dafny.CodePoint('l'.codePointAt(0));
                        })).length)]);
                      }
                    }
                    return _coll61;
                  }()).contains(_1219_v7)) {
                    _coll59.add(_1219_v7);
                  }
                }
                return _coll59;
              }()).length)));
            }
          }
          return _coll58;
        }(), _1212_v3, (_1212_v3).Intersect(_1212_v3));
        let _1220_v9;
        _1220_v9 = _dafny.Seq.of(_1212_v3);
        _1215_v8 = ((_dafny.MultiSet.FromArray(_1220_v9)).Difference(_1215_v8)).update(_1212_v3, _module.__default.abs(((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]).multipliedBy(_this.f29)));
        let _1221_v10;
        _1221_v10 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1222_v11;
        _1222_v11 = _module.D1.create_DC7(_1221_v10, _module.__default.fm9((_this).f28, globalState));
        let _source18 = _module.__default.fm8(((true) ? (_1222_v11) : (_1222_v11)), globalState);
        if (_source18.is_DC10) {
          let _1223___mcc_h0 = (_source18).cf17;
          let _1224___mcc_h1 = (_source18).cf18;
          let _1225_cf18 = _1224___mcc_h1;
          let _1226_cf17 = _1223___mcc_h0;
          _1226_cf17 = (_this).f28;
          let _1227_v12;
          _1227_v12 = _dafny.Seq.UnicodeFromString("d");
          let _1228_v14;
          _1228_v14 = _dafny.MultiSet.fromElements(_1227_v12, _1227_v12);
          (globalState).f4 = new BigNumber((((true) ? (_dafny.Map.Empty.slice().updateUnsafe(_1227_v12,_1222_v11)) : (function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of (_1228_v14).Elements) {
              let _1229_v13 = _compr_62;
              if ((_1228_v14).contains(_1229_v13)) {
                _coll62.push([_1229_v13,_1222_v11]);
              }
            }
            return _coll62;
          }()))).length);
          let _index190 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length));
          (_1208_v0)[_index190] = _1214_i1;
          let _1230_v15;
          let _nw178 = Array((new BigNumber(3)).toNumber()).fill([]);
          _1230_v15 = _nw178;
          let _1231_v16;
          _1231_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1230_v15,new BigNumber(594));
          (globalState).f21 = (((_1231_v16).contains(_1230_v15)) ? ((_1231_v16).get(_1230_v15)) : (_this.f29));
        } else {
          let _1232___mcc_h2 = (_source18).cf16;
          let _1233_cf16 = _1232___mcc_h2;
          (globalState).f25 = _1221_v10;
          let _1234_v17;
          _1234_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(new _dafny.CodePoint('u'.codePointAt(0)), globalState),(_this).f28);
          let _1235_v18;
          _1235_v18 = _dafny.Seq.UnicodeFromString("jrlr");
          _1234_v17 = _module.__default.fm11((_1210_v1).IsProperSubsetOf(_1210_v1), !(_dafny.Seq.IsPrefixOf(_1235_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-321)), ((_1236_v10) => function (_1237_i3) {
            return _1236_v10;
          })(_1221_v10)))), p0, globalState);
          let _1238_v19;
          _1238_v19 = _module.D0.create_DC3(_this.f29, p0, p0);
          let _1239_v20;
          _1239_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1238_v19,_1214_i1);
          let _1240_v21;
          let _nw179 = new _module.C3();
          _nw179.__ctor((_this.f29).multipliedBy((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]), _1239_v20, _this.f29, (_this).f27, p0);
          _1240_v21 = _nw179;
          let _1241_v22;
          let _nw180 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1241_v22 = _nw180;
          let _index191 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1241_v22).length));
          (_1241_v22)[_index191] = _1235_v18;
          let _index192 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1241_v22).length));
          (_1241_v22)[_index192] = _dafny.Seq.Concat(_1235_v18, (((p2).contains(p0)) ? ((p2).get(p0)) : (_dafny.Seq.update(_1235_v18, _module.__default.safeIndex(_this.f29, new BigNumber((_1235_v18).length)), _1221_v10))));
        }
        (globalState).f4 = (_dafny.ZERO).minus(_this.f29);
      }
      let _1242_i4;
      _1242_i4 = _dafny.ZERO;
      L4: {
        while (p0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1242_i4)) {
              break L4;
            }
            _1242_i4 = (_1242_i4).plus(_dafny.ONE);
            let _1243_v23;
            _1243_v23 = _dafny.MultiSet.fromElements((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]);
            let _1244_v24;
            _1244_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))],_1208_v0);
            let _1245_v25;
            _1245_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1244_v24).length));
            let _1246_v26;
            _1246_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1245_v25);
            let _1247_v27;
            _1247_v27 = _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(!(p0),(_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))])).update(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1243_v23)).length)), _1245_v25, (((_1246_v26).contains((_this).f28)) ? ((_1246_v26).get((_this).f28)) : (_1245_v25)));
            (globalState).f6 = (_1247_v27).IsDisjointFrom((_1247_v27).update(_1245_v25, _module.__default.abs(new BigNumber(922))));
            let _1248_v28;
            _1248_v28 = _dafny.Seq.of((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))], (_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))], _this.f29, new BigNumber(441));
            let _1249_v29;
            _1249_v29 = _dafny.Seq.UnicodeFromString("mcumuyosc");
            let _1250_v30;
            _1250_v30 = _dafny.Seq.of(new BigNumber((_module.__default.fm25((_1248_v28)[_module.__default.safeIndex(new BigNumber((_1249_v29).length), new BigNumber((_1248_v28).length))], (_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))], _this.f29, globalState)).length));
            let _1251_v31;
            _1251_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28);
            let _1252_v32;
            _1252_v32 = new _dafny.CodePoint('w'.codePointAt(0));
            let _1253_v33;
            let _nw181 = new _module.C2();
            _nw181.__ctor(_1252_v32, p0, (_this).f27, true);
            _1253_v33 = _nw181;
            let _1254_v34;
            _1254_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1253_v33,_1252_v32);
            (globalState).f13 = (new BigNumber((_1248_v28).length)).minus(new BigNumber((((_module.__default.fm16(_1251_v31, globalState)) ? (_1254_v34) : (_1254_v34))).length));
            _1250_v30 = _1250_v30;
            let _1255_v35;
            let _nw182 = new _module.C4();
            _nw182.__ctor();
            _1255_v35 = _nw182;
          }
        }
      }
      let _1256_v36;
      _1256_v36 = _dafny.Seq.of((_dafny.ZERO).minus((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]));
      let _1257_v37;
      _1257_v37 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_1256_v36))).IsSubsetOf(_dafny.Set.fromElements(_1256_v36, _1256_v36)),!(new BigNumber(243)).isEqualTo((_1208_v0)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_1208_v0).length))]));
      _1257_v37 = _1257_v37;
      let _1258_v38;
      _1258_v38 = _dafny.Seq.UnicodeFromString("g");
      (globalState).f5 = new BigNumber((_1258_v38).length);
      (globalState).f15 = !(!(p0) || (p0));
      r0 = _this.f29;
      return r0;
    }
    get f33() {
      let _this = this;
      return _this._f33;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f29, f27, f28) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return _this.f29;
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f6 = !(p0);
      let _1259_v0;
      _1259_v0 = _dafny.Set.fromElements(p2, new BigNumber(-885), p2);
      let _1260_v1;
      _1260_v1 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1261_v2;
      _1261_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1260_v1);
      let _1262_v4;
      _1262_v4 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f29),function () {
        let _coll63 = new _dafny.Set();
        for (const _compr_63 of (_1261_v2).Keys.Elements) {
          let _1263_v3 = _compr_63;
          if ((_1261_v2).contains(_1263_v3)) {
            _coll63.add((_1263_v3).minus(new BigNumber(-91)));
          }
        }
        return _coll63;
      }());
      let _1264_v5;
      _1264_v5 = _module.D3.create_DC9(_1262_v4);
      _1259_v0 = _module.__default.fm0((_1264_v5).dtor_cf16, p0, !(((_this).f28) === (p1)), globalState);
      let _1265_v6;
      let _1266_v7;
      let _out23;
      let _out24;
      let _outcollector9 = _module.__default.m0(globalState);
      _out23 = _outcollector9[0];
      _out24 = _outcollector9[1];
      _1265_v6 = _out23;
      _1266_v7 = _out24;
      let _1267_v8;
      _1267_v8 = _dafny.Seq.of(_1259_v0);
      let _1268_v9;
      _1268_v9 = _dafny.Seq.UnicodeFromString("kwuoanirg");
      let _1269_v10;
      _1269_v10 = _module.D0.create_DC2((_1267_v8)[_module.__default.safeIndex(new BigNumber((_1268_v9).length), new BigNumber((_1267_v8).length))], _1268_v9);
      let _source19 = _1269_v10;
      if (_source19.is_DC0) {
        let _1270___mcc_h0 = (_source19).cf0;
        let _1271___mcc_h1 = (_source19).cf1;
        let _1272___mcc_h2 = (_source19).cf2;
        let _1273___mcc_h3 = (_source19).cf3;
        let _1274___mcc_h4 = (_source19).cf4;
        let _1275_cf4 = _1274___mcc_h4;
        let _1276_cf3 = _1273___mcc_h3;
        let _1277_cf2 = _1272___mcc_h2;
        let _1278_cf1 = _1271___mcc_h1;
        let _1279_cf0 = _1270___mcc_h0;
        let _1280_v11;
        let _nw183 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1280_v11 = _nw183;
        _1280_v11 = (_this).f27;
        let _1281_v12;
        let _init27 = ((_1282_p1, _1283_cf4) => function (_1284_i0) {
          return (_1284_i0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1282_p1,_1283_cf4)).length));
        })(p1, _1275_cf4);
        let _nw184 = Array((new BigNumber(21)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw184.length); _i0_27++) {
          _nw184[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1281_v12 = _nw184;
        let _index193 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1281_v12).length));
        (_1281_v12)[_index193] = _this.f29;
        let _index194 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1281_v12).length));
        (_1281_v12)[_index194] = _this.f29;
        let _1285_v13;
        _1285_v13 = _module.D0.create_DC1();
        _1285_v13 = _1285_v13;
        let _1286_v14;
        let _nw185 = new _module.C2();
        _nw185.__ctor(_1276_cf3, true, _1280_v11, false);
        _1286_v14 = _nw185;
      } else if (_source19.is_DC1) {
        let _1287_v15;
        _1287_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1266_v7,_this.f29);
        (globalState).f4 = ((!(((p1) ? (p1) : ((_this).f28)))) ? (_1266_v7) : ((_this.f29).minus(new BigNumber((_1287_v15).length))));
        (globalState).f5 = new BigNumber(-700);
        let _1288_v16;
        _1288_v16 = _module.D0.create_DC3(_this.f29, p0, !(p1));
        let _1289_v17;
        _1289_v17 = _dafny.Seq.of(p1, p1);
        let _1290_v18;
        _1290_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1288_v16,new BigNumber((_1289_v17).length));
        let _1291_v19;
        let _nw186 = new _module.C3();
        _nw186.__ctor(_this.f29, ((p1) ? (_1290_v18) : (_dafny.Map.Empty.slice().updateUnsafe(_1288_v16,(_dafny.ZERO).minus(_1266_v7)))), ((_dafny.ZERO).minus(p2)).plus(p2), (_this).f27, (_this).f28);
        _1291_v19 = _nw186;
        let _1292_v20;
        let _init28 = ((_1293_p0, _1294_p1) => function (_1295_i1) {
          return _module.__default.safeDivisionInt(_1295_i1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1293_p0,_1294_p1)).length));
        })(p0, p1);
        let _nw187 = Array((new BigNumber(16)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw187.length); _i0_28++) {
          _nw187[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1292_v20 = _nw187;
        let _index195 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1292_v20).length));
        (_1292_v20)[_index195] = (_this.f29).minus((_1291_v19).f34);
        let _1296_v21;
        _1296_v21 = _dafny.Seq.of(_1292_v20);
        let _1297_v22;
        _1297_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f27);
        let _1298_v23;
        let _nw188 = new _module.C5();
        _nw188.__ctor(_1296_v21, _this.f29, (((_1297_v22).contains(true)) ? ((_1297_v22).get(true)) : ((_this).f27)), (_this).f28);
        _1298_v23 = _nw188;
        let _1299_v24;
        _1299_v24 = _dafny.MultiSet.fromElements((_1296_v21)[_module.__default.safeIndex((_1291_v19).f34, new BigNumber((_1296_v21).length))]);
        let _1300_v25;
        _1300_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v23,new BigNumber((_1299_v24).cardinality()));
        let _index196 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1292_v20).length));
        (_1292_v20)[_index196] = (_module.__default.safeModuloInt((_dafny.ZERO).minus((((_1300_v25).contains(_1298_v23)) ? ((_1300_v25).get(_1298_v23)) : (_1266_v7))), (_1291_v19).f34)).minus(_1266_v7);
      } else if (_source19.is_DC2) {
        let _1301___mcc_h5 = (_source19).cf5;
        let _1302___mcc_h6 = (_source19).cf6;
        let _1303_cf6 = _1302___mcc_h6;
        let _1304_cf5 = _1301___mcc_h5;
        let _1305_v26;
        let _1306_v27;
        let _out25;
        let _out26;
        let _outcollector10 = _module.__default.m0(globalState);
        _out25 = _outcollector10[0];
        _out26 = _outcollector10[1];
        _1305_v26 = _out25;
        _1306_v27 = _out26;
        let _1307_v28;
        _1307_v28 = _dafny.MultiSet.fromElements(_module.__default.fm46(p0, globalState));
        let _rhs211 = _1268_v9;
        let _rhs212 = ((_1307_v28).Union(_1307_v28)).Difference(_1307_v28);
        let _rhs213 = p0;
        let _lhs188 = globalState;
        _1303_cf6 = _rhs211;
        _1307_v28 = _rhs212;
        _lhs188.f6 = _rhs213;
        let _1308_v29;
        let _nw189 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _1308_v29 = _nw189;
        let _1309_v30;
        _1309_v30 = _dafny.MultiSet.fromElements(_1266_v7);
        let _index197 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1308_v29).length));
        (_1308_v29)[_index197] = _module.__default.fm49(_1309_v30, globalState);
        let _1310_v31;
        _1310_v31 = _dafny.Seq.of((_this).f28);
        let _1311_v32;
        _1311_v32 = _dafny.Seq.of(_1310_v31);
        let _index198 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1308_v29).length));
        (_1308_v29)[_index198] = _dafny.Seq.Concat(_1311_v32, _dafny.Seq.Concat(_1311_v32, _1311_v32));
        (globalState).f6 = p0;
      } else if (_source19.is_DC3) {
        let _1312___mcc_h7 = (_source19).cf7;
        let _1313___mcc_h8 = (_source19).cf8;
        let _1314___mcc_h9 = (_source19).cf9;
        let _1315_cf9 = _1314___mcc_h9;
        let _1316_cf8 = _1313___mcc_h8;
        let _1317_cf7 = _1312___mcc_h7;
        (globalState).f13 = (_module.__default.fm36(globalState)).plus(_1266_v7);
        let _1318_v33;
        _1318_v33 = _module.D3.create_DC10(p1, _1315_cf9);
        let _1319_v34;
        _1319_v34 = _dafny.Set.fromElements(_1318_v33);
        let _1320_v35;
        _1320_v35 = _dafny.MultiSet.fromElements((_this).f28, !(_1316_cf8), p1);
        let _1321_v36;
        _1321_v36 = _dafny.Seq.of((((_1320_v35).contains(false)) ? ((_1320_v35).get(false)) : (_this.f29)), _1266_v7);
        let _rhs214 = _1266_v7;
        let _rhs215 = ((true) ? (_dafny.Set.fromElements(_1318_v33)) : (_dafny.Set.fromElements(_module.D3.create_DC10(_1315_cf9, (_this).f28))));
        let _rhs216 = _1321_v36;
        let _lhs189 = globalState;
        _lhs189.f21 = _rhs214;
        _1319_v34 = _rhs215;
        _1321_v36 = _rhs216;
        (_this).f29 = _module.__default.fm36(globalState);
        let _1322_v37;
        let _nw190 = new _module.C4();
        _nw190.__ctor();
        _1322_v37 = _nw190;
      } else {
        let _1323___mcc_h10 = (_source19).cf10;
        let _1324_cf10 = _1323___mcc_h10;
        let _1325_v38;
        _1325_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _1326_v39;
        let _nw191 = new _module.C0();
        _nw191.__ctor(_dafny.Seq.Concat(_1268_v9, _1268_v9), (_dafny.ZERO).minus(((p1) ? (new BigNumber((_1325_v38).length)) : (p2))), (_this).f27, (_this).f28);
        _1326_v39 = _nw191;
        let _1327_v40;
        _1327_v40 = _dafny.Set.fromElements(p1, _module.__default.fm40(globalState));
        let _1328_v41;
        _1328_v41 = _dafny.Seq.of((_this).f28, p0);
        let _1329_v42;
        let _nw192 = Array((new BigNumber(7)).toNumber());
        _nw192[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1327_v40).Intersect(_1327_v40)).length));
        _nw192[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1266_v7);
        _nw192[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw192[(new BigNumber(3)).toNumber()] = _this.f29;
        _nw192[(new BigNumber(4)).toNumber()] = _this.f29;
        _nw192[(new BigNumber(5)).toNumber()] = new BigNumber((_1328_v41).length);
        _nw192[(new BigNumber(6)).toNumber()] = p2;
        _1329_v42 = _nw192;
        let _1330_v43;
        _1330_v43 = _dafny.Seq.of(_1329_v42);
        let _1331_v44;
        _1331_v44 = _dafny.MultiSet.fromElements(_1329_v42, (_1330_v43)[_module.__default.safeIndex(_1266_v7, new BigNumber((_1330_v43).length))], _1329_v42);
        let _index199 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1329_v42).length));
        (_1329_v42)[_index199] = new BigNumber((_1331_v44).cardinality());
        let _index200 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1329_v42).length));
        (_1329_v42)[_index200] = (_this).fm5(globalState);
        let _1332_v45;
        _1332_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1260_v1,_1330_v43);
        let _1333_v46;
        let _nw193 = new _module.C5();
        _nw193.__ctor((((_1332_v45).contains(new _dafny.CodePoint('q'.codePointAt(0)))) ? ((_1332_v45).get(new _dafny.CodePoint('q'.codePointAt(0)))) : (_1330_v43)), _1266_v7, (_this).f27, (_1328_v41)[_module.__default.safeIndex((_1329_v42)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_1329_v42).length))], new BigNumber((_1328_v41).length))]);
        _1333_v46 = _nw193;
        let _1334_v47;
        _1334_v47 = _dafny.Seq.of(_1266_v7, (_1326_v39).fm15(globalState), new BigNumber(406), (_dafny.ZERO).minus(new BigNumber(-368)), _module.__default.fm36(globalState));
        let _rhs217 = ((_1334_v47)[_module.__default.safeIndex(p2, new BigNumber((_1334_v47).length))]).multipliedBy(_module.__default.safeModuloInt(_1266_v7, _module.__default.fm36(globalState)));
        let _rhs218 = _dafny.Seq.IsPrefixOf(_1328_v41, _1328_v41);
        let _rhs219 = (_1333_v46).f33;
        let _rhs220 = !(!(p0)) || (p1);
        let _lhs190 = globalState;
        let _lhs191 = globalState;
        let _lhs192 = globalState;
        _lhs190.f13 = _rhs217;
        _lhs191.f15 = _rhs218;
        _1330_v43 = _rhs219;
        _lhs192.f6 = _rhs220;
      }
      let _1335_v48;
      _1335_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this.f29).minus(_1266_v7),p0);
      _1335_v48 = (_1335_v48).update(new BigNumber((_1268_v9).length), (_this).f28);
      let _1336_v49;
      _1336_v49 = _dafny.Seq.of(p1);
      let _1337_v50;
      _1337_v50 = _dafny.MultiSet.fromElements(p0, !(p0), (_this).f28, false, true);
      let _1338_v51;
      _1338_v51 = _dafny.Seq.of(_1337_v50);
      let _index201 = _module.__default.safeIndex(new BigNumber(359), new BigNumber(((_this).f27).length));
      ((_this).f27)[_index201] = (_1336_v49)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1338_v51)).length), new BigNumber((_1336_v49).length))];
      let _index202 = _module.__default.safeIndex(new BigNumber(359), new BigNumber(((_this).f27).length));
      ((_this).f27)[_index202] = (p2).isEqualTo(_1266_v7);
      return;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f27, f28, f29) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _1339_v0;
      _1339_v0 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1340_v1;
      _1340_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_dafny.ZERO).minus(_this.f29));
      let _1341_v2;
      _1341_v2 = _dafny.Seq.of((_dafny.ZERO).minus((((_1340_v1).contains(true)) ? ((_1340_v1).get(true)) : (_this.f29))), _this.f29);
      let _rhs221 = _this.f29;
      let _rhs222 = _1339_v0;
      let _rhs223 = !(!((_this).f28));
      let _rhs224 = _module.__default.safeModuloInt(_this.f29, (_1341_v2)[_module.__default.safeIndex(_this.f29, new BigNumber((_1341_v2).length))]);
      let _lhs193 = globalState;
      let _lhs194 = globalState;
      r0 = _rhs221;
      _lhs193.f25 = _rhs222;
      _lhs194.f17 = _rhs223;
      r0 = _rhs224;
      let _1342_v3;
      _1342_v3 = _dafny.Seq.UnicodeFromString("yaenitp");
      (globalState).f17 = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("rhmgaits"), _dafny.Seq.Concat(_1342_v3, _1342_v3));
      (globalState).f6 = (((_this).f28) ? ((_this).f28) : ((((_this).f28) ? ((_this).f28) : ((_this).f28))));
      let _1343_v4;
      _1343_v4 = _1341_v2;
      let _source20 = _1343_v4;
      let _1344___mcc_h0 = _source20;
      let _1345_cf15 = _1344___mcc_h0;
      (globalState).f25 = (_1342_v3)[_module.__default.safeIndex(_this.f29, new BigNumber((_1342_v3).length))];
      _1339_v0 = _1339_v0;
      let _1346_v5;
      _1346_v5 = _module.D1.create_DC6(_this.f29);
      let _1347_v6;
      _1347_v6 = _dafny.MultiSet.fromElements(_1346_v5, _1346_v5);
      let _1348_v7;
      _1348_v7 = _dafny.Seq.of(_1347_v6);
      let _1349_v8;
      _1349_v8 = _dafny.Seq.of(true);
      let _1350_v9;
      _1350_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1339_v0);
      let _1351_v10;
      _1351_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1350_v9);
      let _1352_v11;
      _1352_v11 = _dafny.Seq.of((((_1351_v10).contains(false)) ? ((_1351_v10).get(false)) : (_1350_v9)));
      let _1353_v12;
      _1353_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_dafny.MultiSet.fromElements(_1346_v5, _1346_v5));
      let _1354_v13;
      _1354_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f28);
      let _1355_v14;
      let _nw194 = Array((new BigNumber(29)).toNumber());
      _nw194[(_dafny.ZERO).toNumber()] = _this.f29;
      _nw194[(_dafny.ONE).toNumber()] = _this.f29;
      _nw194[(new BigNumber(2)).toNumber()] = _module.__default.fm4((_1348_v7)[_module.__default.safeIndex(_this.f29, new BigNumber((_1348_v7).length))], globalState);
      _nw194[(new BigNumber(3)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(4)).toNumber()] = new BigNumber((_1342_v3).length);
      _nw194[(new BigNumber(5)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(6)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(7)).toNumber()] = new BigNumber((_1349_v8).length);
      _nw194[(new BigNumber(8)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(9)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(10)).toNumber()] = (((_1340_v1).contains(true)) ? ((_1340_v1).get(true)) : (_this.f29));
      _nw194[(new BigNumber(11)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(12)).toNumber()] = new BigNumber(220);
      _nw194[(new BigNumber(13)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(14)).toNumber()] = new BigNumber(((_1352_v11)[_module.__default.safeIndex(_this.f29, new BigNumber((_1352_v11).length))]).length);
      _nw194[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f29));
      _nw194[(new BigNumber(16)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(17)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(18)).toNumber()] = new BigNumber(824);
      _nw194[(new BigNumber(19)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(20)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(21)).toNumber()] = new BigNumber(20);
      _nw194[(new BigNumber(22)).toNumber()] = new BigNumber(360);
      _nw194[(new BigNumber(23)).toNumber()] = _module.__default.fm4((((_1353_v12).contains(_this.f29)) ? ((_1353_v12).get(_this.f29)) : (_1347_v6)), globalState);
      _nw194[(new BigNumber(24)).toNumber()] = _this.f29;
      _nw194[(new BigNumber(25)).toNumber()] = new BigNumber((_1345_cf15).length);
      _nw194[(new BigNumber(26)).toNumber()] = new BigNumber(660);
      _nw194[(new BigNumber(27)).toNumber()] = new BigNumber((_1342_v3).length);
      _nw194[(new BigNumber(28)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1354_v13).length));
      _1355_v14 = _nw194;
      let _1356_v15;
      _1356_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f28) ? (true) : ((_this).f28)),_1355_v14);
      _1356_v15 = (_1356_v15).update(!((_this).f28), _1355_v14);
      let _1357_v16;
      _1357_v16 = _dafny.Seq.of(_1355_v14);
      let _1358_v17;
      let _nw195 = new _module.C5();
      _nw195.__ctor(_1357_v16, _this.f29, (_this).f27, !((_this).f28));
      _1358_v17 = _nw195;
      let _1359_v18;
      _1359_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1358_v17,_this.f29);
      let _1360_v19;
      _1360_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1359_v18).update(_1358_v17, new BigNumber((_1342_v3).length)),(_this).f28);
      let _1361_v20;
      _1361_v20 = _module.D0.create_DC0(_1360_v19, (_this).f28, _1358_v17.f29, _1339_v0, (_1358_v17).f28);
      let _1362_v21;
      _1362_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_1358_v17).f28);
      let _1363_v22;
      let _nw196 = Array((new BigNumber(24)).toNumber());
      _nw196[(_dafny.ZERO).toNumber()] = (_module.__default.fm4(_1347_v6, globalState)).isLessThan(new BigNumber(902));
      _nw196[(_dafny.ONE).toNumber()] = !((_this).f28);
      _nw196[(new BigNumber(2)).toNumber()] = (_this).f28;
      _nw196[(new BigNumber(3)).toNumber()] = !((_this).f28);
      _nw196[(new BigNumber(4)).toNumber()] = !((_this).f28);
      _nw196[(new BigNumber(5)).toNumber()] = (_1361_v20).dtor_cf1;
      _nw196[(new BigNumber(6)).toNumber()] = (_this).f28;
      _nw196[(new BigNumber(7)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(8)).toNumber()] = !((_1358_v17).f28);
      _nw196[(new BigNumber(9)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(10)).toNumber()] = !(_module.__default.fm10(_1339_v0, globalState));
      _nw196[(new BigNumber(11)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(12)).toNumber()] = ((_this).f28) === ((((_1362_v21).contains((_this).f28)) ? ((_1362_v21).get((_this).f28)) : ((_1358_v17).f28)));
      _nw196[(new BigNumber(13)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(14)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(15)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(16)).toNumber()] = _module.__default.fm22((_1358_v17).f28, _1358_v17.f29, globalState);
      _nw196[(new BigNumber(17)).toNumber()] = ((true) ? ((_this).f28) : (true));
      _nw196[(new BigNumber(18)).toNumber()] = (_this).f28;
      _nw196[(new BigNumber(19)).toNumber()] = !((_this).f28);
      _nw196[(new BigNumber(20)).toNumber()] = (_this).f28;
      _nw196[(new BigNumber(21)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(22)).toNumber()] = (_1358_v17).f28;
      _nw196[(new BigNumber(23)).toNumber()] = (_this).f28;
      _1363_v22 = _nw196;
      _1363_v22 = (_1358_v17).f27;
      let _1364_v23;
      _1364_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC26(true),(_this).f28);
      let _1365_v24;
      let _nw197 = new _module.C1();
      _nw197.__ctor((((_1364_v23).contains(_module.D10.create_DC26((_this).f28))) ? ((_1364_v23).get(_module.D10.create_DC26((_this).f28))) : ((_this).f28)), (_this).f27, (_this.f29).isLessThan(new BigNumber(-978)));
      _1365_v24 = _nw197;
      let _hi6 = _this.f29;
      for (let _1366_i0 = _this.f29; _1366_i0.isLessThan(_hi6); _1366_i0 = _1366_i0.plus(_dafny.ONE)) {
        r0 = (_module.__default.safeModuloInt(new BigNumber(894), _this.f29)).minus(_module.__default.fm36(globalState));
        let _1367_v25;
        let _nw198 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _1367_v25 = _nw198;
        let _source21 = _module.D11.create_DC27(_1367_v25);
        if (_source21.is_DC28) {
          _1342_v3 = _1342_v3;
          _1342_v3 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bdswexrbt"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("am"), _1342_v3));
          let _1368_v26;
          _1368_v26 = _dafny.Set.fromElements((_1365_v24).f39, false, (_1365_v24).f39);
          let _1369_v27;
          _1369_v27 = _module.D8.create_DC18(new BigNumber((_1368_v26).length));
          _1369_v27 = _1369_v27;
          let _1370_v28;
          _1370_v28 = _dafny.MultiSet.fromElements((_this).f28);
          let _1371_v29;
          _1371_v29 = _dafny.MultiSet.fromElements((_1370_v28).Difference(_1370_v28), (((_1365_v24).f39) ? (_1370_v28) : (_1370_v28)), _1370_v28);
          let _rhs225 = (_dafny.ZERO).minus(_1366_i0);
          let _rhs226 = new BigNumber((_1371_v29).cardinality());
          let _lhs195 = globalState;
          let _lhs196 = globalState;
          _lhs195.f5 = _rhs225;
          _lhs196.f13 = _rhs226;
        } else {
          let _1372___mcc_h1 = (_source21).cf42;
          let _1373_cf42 = _1372___mcc_h1;
          (globalState).f25 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1374_v30;
          _1374_v30 = _dafny.MultiSet.fromElements(_this.f29);
          let _1375_v31;
          _1375_v31 = _dafny.MultiSet.fromElements(true);
          let _1376_v32;
          _1376_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1375_v31);
          let _1377_v33;
          _1377_v33 = _dafny.Seq.of((_1375_v31).update((_1365_v24).f39, _module.__default.abs(_this.f29)), (((_1376_v32).contains(_1366_i0)) ? ((_1376_v32).get(_1366_i0)) : (_1375_v31)), _1375_v31);
          (globalState).f6 = ((((_1374_v30).contains(_this.f29)) ? ((_1374_v30).get(_this.f29)) : (new BigNumber(((_1377_v33)[_module.__default.safeIndex(_1366_i0, new BigNumber((_1377_v33).length))]).cardinality())))).isLessThanOrEqualTo(_1366_i0);
          let _1378_v34;
          let _init29 = ((_1379_v3) => function (_1380_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1379_v3).length),(_this).f28);
          })(_1342_v3);
          let _nw199 = Array((new BigNumber(13)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw199.length); _i0_29++) {
            _nw199[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1378_v34 = _nw199;
          let _nw200 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _1378_v34 = _nw200;
          let _1381_v35;
          let _nw201 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1381_v35 = _nw201;
          let _1382_v36;
          _1382_v36 = _dafny.Seq.of(_1381_v35, _1381_v35, _1381_v35);
          let _1383_v37;
          _1383_v37 = _dafny.MultiSet.fromElements(_1381_v35, _1381_v35, (_1382_v36)[_module.__default.safeIndex(_this.f29, new BigNumber((_1382_v36).length))], _1381_v35, _1381_v35);
          (globalState).f17 = (_1383_v37).IsProperSubsetOf((_1383_v37).update(_1381_v35, _module.__default.abs(_this.f29)));
        }
        let _1384_v38;
        _1384_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm50((_this).f28, _this.f29, globalState),(_dafny.ZERO).minus(_1366_i0));
        _1384_v38 = (_1384_v38).update(_1340_v1, new BigNumber(370));
        (globalState).f13 = new BigNumber((_1341_v2).length);
      }
      r0 = (_dafny.ZERO).minus(_this.f29);
      r1 = (_this.f29).multipliedBy(_this.f29);
      return [r0, r1];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f27 = [];
      this._f28 = false;
      this._f32 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f32, f27, f28) {
      let _this = this;
      (_this)._f32 = f32;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((_module.D1.create_DC5(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), function (_1385_i0) {
  return new BigNumber(-103);
})))).dtor_cf11, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(323)), _dafny.Seq.of(new BigNumber(-823), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(164))).cardinality())), _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f28)).length)), (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-361))).length))))));
    };
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _1386_v0;
      let _nw202 = new _module.C0();
      _nw202.__ctor(_dafny.Seq.UnicodeFromString("yovobx"), new BigNumber(139), (_this).f27, (_this).f28);
      _1386_v0 = _nw202;
      let _1387_v1;
      _1387_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1386_v0,_1386_v0.f29);
      let _1388_v2;
      _1388_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1387_v1,(_1386_v0).f28);
      let _1389_v3;
      _1389_v3 = _module.D0.create_DC0(_1388_v2, (_1386_v0).f28, new BigNumber(130), (_this).f32, (_1386_v0).f28);
      let _1390_v4;
      _1390_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_1389_v3).dtor_cf1);
      let _1391_v5;
      _1391_v5 = _module.D1.create_DC6(_1386_v0.f29);
      let _1392_v6;
      _1392_v6 = _dafny.Set.fromElements(_module.__default.fm4(_dafny.MultiSet.fromElements(_1391_v5, _module.D1.create_DC6(p0), _1391_v5), globalState));
      let _1393_v7;
      _1393_v7 = _dafny.Seq.UnicodeFromString("upehw");
      let _1394_v8;
      _1394_v8 = _module.D0.create_DC2(_1392_v6, (((_1386_v0).f28) ? (_1393_v7) : (_dafny.Seq.UnicodeFromString("ojwsghveo"))));
      let _1395_v9;
      _1395_v9 = _dafny.Seq.of((_1386_v0).f28, (_this).f28, (_1386_v0).f28);
      let _rhs227 = _1390_v4;
      let _rhs228 = ((true) ? ((_this).f28) : ((_1395_v9)[_module.__default.safeIndex(p0, new BigNumber((_1395_v9).length))]));
      let _rhs229 = _1394_v8;
      _1390_v4 = _rhs227;
      r1 = _rhs228;
      _1394_v8 = _rhs229;
      _1390_v4 = (_1390_v4).update((_this).f28, (_1386_v0).f28);
      let _1396_i0;
      _1396_i0 = _dafny.ZERO;
      L5: {
        while (((_1386_v0).f28) || ((_this).f28)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1396_i0)) {
              break L5;
            }
            _1396_i0 = (_1396_i0).plus(_dafny.ONE);
            let _1397_v10;
            let _init30 = function (_1398_i1) {
              return (_1398_i1).minus(new BigNumber(124));
            };
            let _nw203 = Array((new BigNumber(14)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw203.length); _i0_30++) {
              _nw203[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _1397_v10 = _nw203;
            _1397_v10 = _1397_v10;
            (_1386_v0).f29 = _1386_v0.f29;
            let _1399_v11;
            let _nw204 = new _module.C4();
            _nw204.__ctor();
            _1399_v11 = _nw204;
            _1399_v11 = _1399_v11;
            let _1400_v12;
            _1400_v12 = _dafny.Seq.of(p0);
            let _1401_v13;
            _1401_v13 = _1400_v12;
            let _rhs230 = _1392_v6;
            let _rhs231 = _1401_v13;
            let _rhs232 = _1386_v0.f29;
            let _rhs233 = (_this).f28;
            let _lhs197 = globalState;
            _1392_v6 = _rhs230;
            _1401_v13 = _rhs231;
            _lhs197.f13 = _rhs232;
            r1 = _rhs233;
          }
        }
      }
      let _1402_v14;
      _1402_v14 = _dafny.MultiSet.fromElements(_1386_v0.f29);
      let _1403_v16;
      _1403_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1386_v0).f28,p0);
      let _1404_v17;
      _1404_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1403_v16,(_1386_v0).f28);
      let _1405_v18;
      _1405_v18 = _dafny.Set.fromElements(_module.__default.fm16(_1390_v4, globalState), (new BigNumber((_1402_v14).cardinality())).isLessThanOrEqualTo(new BigNumber((function () {
        let _coll64 = new _dafny.Map();
        for (const _compr_64 of (_1404_v17).Keys.Elements) {
          let _1406_v15 = _compr_64;
          if ((_1404_v17).contains(_1406_v15)) {
            _coll64.push([_1406_v15,_1403_v16]);
          }
        }
        return _coll64;
      }()).length)));
      (globalState).f13 = (_dafny.ZERO).minus(new BigNumber((_1405_v18).length));
      let _rhs234 = (((_1386_v0).f28) ? (!(!((_this).f28))) : (_module.__default.fm40(globalState)));
      let _rhs235 = (_this).f28;
      let _rhs236 = _1386_v0.f29;
      let _lhs198 = globalState;
      let _lhs199 = globalState;
      r1 = _rhs234;
      _lhs198.f17 = _rhs235;
      _lhs199.f21 = _rhs236;
      let _1407_i2;
      _1407_i2 = _dafny.ZERO;
      L6: {
        while ((_this).f28) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1407_i2)) {
              break L6;
            }
            _1407_i2 = (_1407_i2).plus(_dafny.ONE);
            let _1408_v19;
            let _nw205 = new _module.C1();
            _nw205.__ctor((_this).f28, (_1386_v0).f27, (_this).f28);
            _1408_v19 = _nw205;
            _1405_v18 = _1405_v18;
            let _1409_v20;
            _1409_v20 = _dafny.MultiSet.fromElements(_1391_v5);
            (_1386_v0).f29 = _module.__default.fm4(_1409_v20, globalState);
            if ((_1408_v19).f39) {
              let _1410_v21;
              _1410_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cdotqlbut")).length),(_1386_v0).f27);
              let _1411_v22;
              _1411_v22 = (_1386_v0).f27;
              let _1412_v23;
              let _nw206 = Array((new BigNumber(26)).toNumber());
              _nw206[(_dafny.ZERO).toNumber()] = (_1386_v0).f27;
              _nw206[(_dafny.ONE).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(2)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(3)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(4)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(5)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(6)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(7)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(8)).toNumber()] = (_this).f27;
              _nw206[(new BigNumber(9)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(10)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(11)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(12)).toNumber()] = (_this).f27;
              _nw206[(new BigNumber(13)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(14)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(15)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(16)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(17)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(18)).toNumber()] = (_this).f27;
              _nw206[(new BigNumber(19)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(20)).toNumber()] = (((_1410_v21).contains(_1386_v0.f29)) ? ((_1410_v21).get(_1386_v0.f29)) : ((_1386_v0).f27));
              _nw206[(new BigNumber(21)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(22)).toNumber()] = (_this).f27;
              _nw206[(new BigNumber(23)).toNumber()] = (_1386_v0).f27;
              _nw206[(new BigNumber(24)).toNumber()] = (_1411_v22);
              _nw206[(new BigNumber(25)).toNumber()] = (_this).f27;
              _1412_v23 = _nw206;
              let _index203 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1412_v23).length));
              (_1412_v23)[_index203] = (_1386_v0).f27;
              let _1413_v24;
              let _nw207 = new _module.C4();
              _nw207.__ctor();
              _1413_v24 = _nw207;
              let _index204 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1412_v23).length));
              let _rhs237 = (_1408_v19).f39;
              let _rhs238 = (_1386_v0).f28;
              let _rhs239 = (((((_1386_v0).f28) ? ((_1408_v19).f39) : ((_1386_v0).f28))) ? ((_1386_v0).f27) : ((_1386_v0).f27));
              let _rhs240 = _1413_v24;
              let _lhs200 = globalState;
              let _lhs201 = _1412_v23;
              let _lhs202 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1412_v23).length));
              r1 = _rhs237;
              _lhs200.f6 = _rhs238;
              _lhs201[_lhs202] = _rhs239;
              _1413_v24 = _rhs240;
              (globalState).f17 = (((_this).f28) ? ((_1395_v9)[_module.__default.safeIndex(p0, new BigNumber((_1395_v9).length))]) : ((((_1390_v4).contains((_1386_v0).f28)) ? ((_1390_v4).get((_1386_v0).f28)) : (!((_1408_v19).f39)))));
              _1393_v7 = _module.__default.fm18((((_1408_v19).f39) ? ((_1408_v19).f39) : ((_1408_v19).f39)), (_dafny.ZERO).minus(_1386_v0.f29), (_this).f32, globalState);
              let _1414_v25;
              let _nw208 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
              _1414_v25 = _nw208;
              _1414_v25 = _1414_v25;
              r1 = (_1386_v0).f28;
            } else {
              let _1415_v26;
              let _nw209 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
              _1415_v26 = _nw209;
              let _1416_v27;
              _1416_v27 = _dafny.Seq.of(_1415_v26, _1415_v26);
              let _1417_v28;
              _1417_v28 = _module.D8.create_DC18(new BigNumber(-681));
              let _1418_v29;
              _1418_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1417_v28,(_1386_v0).f28);
              let _1419_v30;
              let _nw210 = new _module.C5();
              _nw210.__ctor(_1416_v27, new BigNumber((_1418_v29).length), (_this).f27, (_this).f28);
              _1419_v30 = _nw210;
              _1402_v14 = (_1402_v14).Union((_1402_v14).Union(_dafny.MultiSet.fromElements(p0)));
              (globalState).f5 = _module.__default.safeModuloInt(_1386_v0.f29, (_1419_v30).fm7((_this).f32, globalState));
              (globalState).f15 = (_1386_v0).f28;
              let _1420_v31;
              let _nw211 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1420_v31 = _nw211;
              let _index205 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_1420_v31).length));
              (_1420_v31)[_index205] = (_this).f32;
              let _1421_v32;
              _1421_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1386_v0.f29,(_1386_v0).f28);
              let _1422_v33;
              _1422_v33 = _module.D3.create_DC10((_1408_v19).f39, (_this).f28);
              let _1423_v34;
              _1423_v34 = _module.D5.create_DC13((_1386_v0).f28, _1415_v26, (_module.D7.create_DC16((_1408_v19).f39, _1421_v32, (_this).f32, (_1422_v33).dtor_cf17)).dtor_cf29);
              let _pat_let_tv28 = _1386_v0;
              let _1424_v35;
              let _nw212 = Array((new BigNumber(2)).toNumber());
              _nw212[(_dafny.ZERO).toNumber()] = function (_pat_let21_0) {
                return function (_1425_dt__update__tmp_h0) {
                  return function (_pat_let22_0) {
                    return function (_1426_dt__update_hcf21_h0) {
                      return _module.D5.create_DC13(_1426_dt__update_hcf21_h0, (_1425_dt__update__tmp_h0).dtor_cf22, (_1425_dt__update__tmp_h0).dtor_cf23);
                    }(_pat_let22_0);
                  }((_pat_let_tv28).f28);
                }(_pat_let21_0);
              }(_1423_v34);
              _nw212[(_dafny.ONE).toNumber()] = _1423_v34;
              _1424_v35 = _nw212;
              let _index206 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_1424_v35).length));
              (_1424_v35)[_index206] = _1423_v34;
              let _1427_v36;
              _1427_v36 = _dafny.Seq.of(_1393_v7, _1393_v7, _1393_v7, _1393_v7);
              let _index207 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_1420_v31).length));
              let _index208 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_1424_v35).length));
              let _rhs241 = (_this).f32;
              let _rhs242 = _1423_v34;
              let _rhs243 = _dafny.Seq.Concat((_1427_v36)[_module.__default.safeIndex(new BigNumber((_1421_v32).length), new BigNumber((_1427_v36).length))], _1393_v7);
              let _rhs244 = (_1423_v34).dtor_cf22;
              let _rhs245 = (_1405_v18).IsDisjointFrom(_1405_v18);
              let _lhs203 = _1420_v31;
              let _lhs204 = _module.__default.safeIndex(new BigNumber(681), new BigNumber((_1420_v31).length));
              let _lhs205 = _1424_v35;
              let _lhs206 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_1424_v35).length));
              let _lhs207 = globalState;
              _lhs203[_lhs204] = _rhs241;
              _lhs205[_lhs206] = _rhs242;
              _1393_v7 = _rhs243;
              _1415_v26 = _rhs244;
              _lhs207.f15 = _rhs245;
            }
          }
        }
      }
      let _1428_v37;
      let _nw213 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _1428_v37 = _nw213;
      let _1429_v38;
      _1429_v38 = _dafny.Seq.of(_1428_v37);
      r0 = _1429_v38;
      r1 = (new BigNumber(225)).isEqualTo(_module.__default.safeDivisionInt(new BigNumber((_1393_v7).length), (_dafny.ZERO).minus(p0)));
      return [r0, r1];
    }
    m4(p0, globalState) {
      let _this = this;
      let _1430_v0;
      _1430_v0 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("krl")).length));
      let _1431_v1;
      _1431_v1 = _module.D0.create_DC2(_1430_v0, _module.__default.fm32(p0, (_this).f28, globalState));
      let _1432_v2;
      _1432_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f32);
      let _1433_v3;
      let _nw214 = Array((new BigNumber(6)).toNumber());
      _nw214[(_dafny.ZERO).toNumber()] = new BigNumber(((_1431_v1).dtor_cf6).length);
      _nw214[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw214[(new BigNumber(2)).toNumber()] = p0;
      _nw214[(new BigNumber(3)).toNumber()] = new BigNumber((_1432_v2).length);
      _nw214[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_1434_i0) {
        return (_this).f32;
      })).length);
      _nw214[(new BigNumber(5)).toNumber()] = p0;
      _1433_v3 = _nw214;
      let _1435_v4;
      _1435_v4 = _dafny.Seq.UnicodeFromString("wuevn");
      let _1436_v5;
      _1436_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1435_v4,_1433_v3);
      let _1437_v6;
      _1437_v6 = _dafny.Seq.of(_1433_v3, _1433_v3, _1433_v3, _1433_v3, _1433_v3);
      let _1438_v7;
      let _nw215 = Array((new BigNumber(16)).toNumber());
      _nw215[(_dafny.ZERO).toNumber()] = ((!((_this).f28)) ? (_1433_v3) : (_1433_v3));
      _nw215[(_dafny.ONE).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(2)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(3)).toNumber()] = (((_1436_v5).contains(_dafny.Seq.update(_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), _module.__default.fm42(p0, globalState)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), _module.__default.fm42(p0, globalState))).length)), (_this).f32))) ? ((_1436_v5).get(_dafny.Seq.update(_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), _module.__default.fm42(p0, globalState)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), _module.__default.fm42(p0, globalState))).length)), (_this).f32))) : (_1433_v3));
      _nw215[(new BigNumber(4)).toNumber()] = (_1437_v6)[_module.__default.safeIndex(p0, new BigNumber((_1437_v6).length))];
      _nw215[(new BigNumber(5)).toNumber()] = (((_this).f28) ? (_1433_v3) : (_1433_v3));
      _nw215[(new BigNumber(6)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(7)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(8)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(9)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(10)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(11)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(12)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(13)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(14)).toNumber()] = _1433_v3;
      _nw215[(new BigNumber(15)).toNumber()] = _1433_v3;
      _1438_v7 = _nw215;
      let _index209 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length));
      (_1438_v7)[_index209] = _1433_v3;
      let _index210 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length));
      (_1438_v7)[_index210] = _1433_v3;
      if ((_this).f28) {
        let _1439_v8;
        _1439_v8 = _dafny.Set.fromElements((_this).f28);
        let _1440_v9;
        _1440_v9 = _dafny.Seq.of(_1439_v8);
        let _1441_v10;
        _1441_v10 = _module.D5.create_DC12(_1440_v9);
        let _1442_v11;
        _1442_v11 = _dafny.MultiSet.fromElements(false, _module.__default.fm40(globalState));
        let _1443_v12;
        _1443_v12 = _dafny.Seq.of(_1442_v11);
        let _1444_v13;
        _1444_v13 = _1443_v12;
        let _1445_v14;
        _1445_v14 = _dafny.Seq.of((_this).f28);
        let _1446_v15;
        _1446_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of((_this).f28));
        let _rhs246 = _1441_v10;
        let _rhs247 = (_this).f28;
        let _rhs248 = _1444_v13;
        let _rhs249 = _dafny.Seq.Concat(_dafny.Seq.Concat((((_1446_v15).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), function (_1448_i1) {
          return (_this).f32;
        })).length))) ? ((_1446_v15).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), function (_1447_i1) {
          return (_this).f32;
        })).length))) : (_1445_v14)), _1445_v14), _dafny.Seq.Concat(_1445_v14, _1445_v14));
        let _lhs208 = globalState;
        _1441_v10 = _rhs246;
        _lhs208.f15 = _rhs247;
        _1444_v13 = _rhs248;
        _1445_v14 = _rhs249;
        let _1449_v16;
        let _nw216 = new _module.C0();
        _nw216.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_1450_i2) {
          return (_this).f32;
        }), _module.__default.safeDivisionInt(p0, p0), (_this).f27, (_this).f28);
        _1449_v16 = _nw216;
        (globalState).f17 = (_this).f28;
        (globalState).f21 = new BigNumber((_module.__default.fm51((((_this).f28) ? (p0) : (new BigNumber((_1446_v15).length))), globalState)).length);
        let _1451_v17;
        let _nw217 = new _module.C1();
        _nw217.__ctor((_this).f28, (_this).f27, (p0).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))));
        _1451_v17 = _nw217;
      } else {
        let _1452_v18;
        let _nw218 = new _module.C4();
        _nw218.__ctor();
        _1452_v18 = _nw218;
        let _arr2 = (_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))];
        let _index211 = _module.__default.safeIndex(new BigNumber(271), new BigNumber(((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))]).length));
        _arr2[_index211] = p0;
        let _arr3 = (_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))];
        let _index212 = _module.__default.safeIndex(new BigNumber(271), new BigNumber(((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))]).length));
        _arr3[_index212] = p0;
        let _1453_v19;
        _1453_v19 = _dafny.Seq.of(true);
        let _1454_v20;
        _1454_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1453_v19);
        let _1455_v21;
        _1455_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v20,new BigNumber(-922));
        _1455_v21 = (_1455_v21).update(_1454_v20, ((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))])[_module.__default.safeIndex(new BigNumber(271), new BigNumber(((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))]).length))]);
        let _1456_v22;
        _1456_v22 = _dafny.MultiSet.fromElements(true, !(!((_this).f28)));
        let _1457_v23;
        _1457_v23 = _dafny.Seq.of(_1456_v22, _1456_v22);
        let _1458_v24;
        _1458_v24 = _1457_v23;
        let _1459_v25;
        let _nw219 = Array((new BigNumber(25)).toNumber());
        _nw219[(_dafny.ZERO).toNumber()] = _1457_v23;
        _nw219[(_dafny.ONE).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(2)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(3)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(4)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(5)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(6)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(7)).toNumber()] = ((!((_this).f28)) ? (_1458_v24) : (_1458_v24));
        _nw219[(new BigNumber(8)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(9)).toNumber()] = _1457_v23;
        _nw219[(new BigNumber(10)).toNumber()] = ((true) ? (_1458_v24) : (_1458_v24));
        _nw219[(new BigNumber(11)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(12)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(13)).toNumber()] = _module.__default.fm52(p0, p0, ((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))])[_module.__default.safeIndex(new BigNumber(271), new BigNumber(((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))]).length))], globalState);
        _nw219[(new BigNumber(14)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(15)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(16)).toNumber()] = (((_this).f28) ? (_1458_v24) : (_1457_v23));
        _nw219[(new BigNumber(17)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(18)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(19)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(20)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(21)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(22)).toNumber()] = _1458_v24;
        _nw219[(new BigNumber(23)).toNumber()] = _1457_v23;
        _nw219[(new BigNumber(24)).toNumber()] = _1458_v24;
        _1459_v25 = _nw219;
        _1459_v25 = _1459_v25;
        (globalState).f4 = ((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))])[_module.__default.safeIndex(new BigNumber(271), new BigNumber(((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))]).length))];
      }
      let _1460_v26;
      _1460_v26 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f28),p0);
      let _1461_v27;
      _1461_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1460_v26,new _dafny.CodePoint('i'.codePointAt(0)));
      let _1462_v28;
      _1462_v28 = _dafny.Seq.of((_this).f28, (_this).f28);
      _1461_v27 = (_1461_v27).update(_dafny.Map.Empty.slice().updateUnsafe((_1462_v28)[_module.__default.safeIndex(p0, new BigNumber((_1462_v28).length))],p0), (_this).f32);
      let _hi7 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_1463_p0) => function (_1464_i4) {
        return _1463_p0;
      })(p0))).length);
      for (let _1465_i3 = (p0).multipliedBy(p0); _1465_i3.isLessThan(_hi7); _1465_i3 = _1465_i3.plus(_dafny.ONE)) {
        let _index213 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length));
        let _rhs250 = _1433_v3;
        let _rhs251 = ((!(true)) ? (_1465_i3) : ((p0).multipliedBy(p0)));
        let _lhs209 = _1438_v7;
        let _lhs210 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length));
        let _lhs211 = globalState;
        _lhs209[_lhs210] = _rhs250;
        _lhs211.f13 = _rhs251;
        let _1466_v29;
        _1466_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1465_i3,(_this).f28);
        let _1467_v31;
        _1467_v31 = _dafny.Set.fromElements(_1466_v29, _1466_v29, (function () {
          let _coll65 = new _dafny.Map();
          for (const _compr_65 of _dafny.IntegerRange(new BigNumber(515), new BigNumber(829))) {
            let _1468_v30 = _compr_65;
            if (((new BigNumber(515)).isLessThanOrEqualTo(_1468_v30)) && ((_1468_v30).isLessThan(new BigNumber(829)))) {
              _coll65.push([(_1468_v30).plus(p0),(_this).f28]);
            }
          }
          return _coll65;
        }()).update(_1465_i3, (_this).f28), _1466_v29, _1466_v29);
        _1460_v26 = (_1460_v26).update((_this).f28, new BigNumber((_1467_v31).length));
        let _1469_v32;
        _1469_v32 = _module.D1.create_DC7((_this).f32, _1435_v4);
        let _source22 = _1469_v32;
        if (_source22.is_DC6) {
          let _1470___mcc_h0 = (_source22).cf12;
          let _1471_cf12 = _1470___mcc_h0;
          let _1472_v33;
          _1472_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1460_v26).Merge(_1460_v26),(_1471_cf12).multipliedBy((_dafny.ZERO).minus(_1465_i3)));
          _1472_v33 = _1472_v33;
          let _1473_v34;
          _1473_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_1471_cf12);
          let _rhs252 = _dafny.Seq.update(_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), (((_this).f28) ? ((_1469_v32).dtor_cf13) : ((_this).f32))), _module.__default.safeIndex(_1471_cf12, new BigNumber((_dafny.Seq.update(_1435_v4, _module.__default.safeIndex(p0, new BigNumber((_1435_v4).length)), (((_this).f28) ? ((_1469_v32).dtor_cf13) : ((_this).f32)))).length)), (_this).f32);
          let _rhs253 = !((_1473_v34).equals(_1473_v34));
          let _lhs212 = globalState;
          _1435_v4 = _rhs252;
          _lhs212.f15 = _rhs253;
          let _1474_v35;
          let _1475_v36;
          let _out27;
          let _out28;
          let _outcollector11 = _module.__default.m0(globalState);
          _out27 = _outcollector11[0];
          _out28 = _outcollector11[1];
          _1474_v35 = _out27;
          _1475_v36 = _out28;
          let _1476_v37;
          let _nw220 = new _module.C4();
          _nw220.__ctor();
          _1476_v37 = _nw220;
          _1476_v37 = _1476_v37;
        } else if (_source22.is_DC7) {
          let _1477___mcc_h1 = (_source22).cf13;
          let _1478___mcc_h2 = (_source22).cf14;
          let _1479_cf14 = _1478___mcc_h2;
          let _1480_cf13 = _1477___mcc_h1;
          let _1481_v38;
          _1481_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1482_v39;
          let _nw221 = new _module.C5();
          _nw221.__ctor(_1437_v6, new BigNumber((_1481_v38).length), (_this).f27, (_this).f28);
          _1482_v39 = _nw221;
          let _1483_v40;
          _1483_v40 = _dafny.MultiSet.fromElements(_1482_v39, _1482_v39);
          let _1484_v41;
          _1484_v41 = _dafny.Seq.of(_1479_cf14, _1479_cf14);
          let _rhs254 = !((p0).isLessThan(new BigNumber((_1435_v4).length))) || ((_1483_v40).equals(_1483_v40));
          let _rhs255 = (_module.__default.safeDivisionInt(new BigNumber(796), new BigNumber((_1484_v41).length))).plus(p0);
          let _rhs256 = _1465_i3;
          let _rhs257 = (_dafny.ZERO).minus(_1465_i3);
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          let _lhs215 = globalState;
          let _lhs216 = globalState;
          _lhs213.f15 = _rhs254;
          _lhs214.f5 = _rhs255;
          _lhs215.f5 = _rhs256;
          _lhs216.f4 = _rhs257;
          let _1485_v42;
          _1485_v42 = _dafny.MultiSet.fromElements(_1435_v4);
          let _1486_v43;
          _1486_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm30((_this).f28, _dafny.Seq.UnicodeFromString("kcwjncj"), (_this).f28, globalState)).length),_1485_v42);
          let _1487_v44;
          _1487_v44 = _dafny.Seq.of(_1485_v42);
          _1486_v43 = (_1486_v43).update(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_1488_i5) {
            return (_this).f32;
          })).length), (_1487_v44)[_module.__default.safeIndex(_1465_i3, new BigNumber((_1487_v44).length))]);
          let _1489_v45;
          let _1490_v46;
          let _out29;
          let _out30;
          let _outcollector12 = (_this).m3((p0).multipliedBy((_dafny.ZERO).minus(p0)), globalState);
          _out29 = _outcollector12[0];
          _out30 = _outcollector12[1];
          _1489_v45 = _out29;
          _1490_v46 = _out30;
          let _1491_v47;
          let _nw222 = new _module.C1();
          _nw222.__ctor((_this).f28, (_this).f27, (new BigNumber(65)).isLessThanOrEqualTo(new BigNumber(998)));
          _1491_v47 = _nw222;
        } else {
          let _1492___mcc_h3 = (_source22).cf11;
          let _1493_cf11 = _1492___mcc_h3;
          (globalState).f15 = !((_this).f28);
          (globalState).f6 = (_this).f28;
          let _index214 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1433_v3).length));
          (_1433_v3)[_index214] = p0;
          let _index215 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1433_v3).length));
          (_1433_v3)[_index215] = ((_1465_i3).plus(_1465_i3)).minus(_1465_i3);
          let _1494_v48;
          let _nw223 = new _module.C4();
          _nw223.__ctor();
          _1494_v48 = _nw223;
        }
        let _1495_v49;
        _1495_v49 = _module.D7.create_DC16((_this).f28, _1466_v29, (_this).f32, (_this).f28);
        let _1496_v50;
        _1496_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,false);
        if (_module.__default.fm16((((_1495_v49).dtor_cf26) ? (_1496_v50) : (_1496_v50)), globalState)) {
          let _1497_v52;
          _1497_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f32);
          let _1498_v53;
          _1498_v53 = _dafny.MultiSet.fromElements((((_1497_v52).contains(_1465_i3)) ? ((_1497_v52).get(_1465_i3)) : ((_this).f32)));
          let _1499_v54;
          _1499_v54 = _dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll66 = new _dafny.Map();
            for (const _compr_66 of ((_1498_v53).update((_this).f32, _module.__default.abs(_1465_i3))).Elements) {
              let _1500_v51 = _compr_66;
              if (((_1498_v53).update((_this).f32, _module.__default.abs(_1465_i3))).contains(_1500_v51)) {
                _coll66.push([_1500_v51,_1465_i3]);
              }
            }
            return _coll66;
          }()).length), _1465_i3, _1465_i3, _1465_i3);
          let _1501_v55;
          _1501_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1432_v2,new BigNumber((_1499_v54).cardinality()));
          (globalState).f5 = (((_1501_v55).contains(_1432_v2)) ? ((_1501_v55).get(_1432_v2)) : ((new BigNumber((_1462_v28).length)).minus(_1465_i3)));
          let _1502_v56;
          _1502_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(_1465_i3));
          (globalState).f21 = _module.__default.safeDivisionInt((((_1502_v56).contains(_1465_i3)) ? ((_1502_v56).get(_1465_i3)) : (p0)), p0);
          let _1503_v57;
          _1503_v57 = _module.D9.create_DC23(!((_this).f28));
          let _1504_v58;
          _1504_v58 = _module.D9.create_DC24(_1503_v57);
          let _1505_v59;
          _1505_v59 = _module.D11.create_DC28();
          let _1506_v60;
          let _nw224 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
          _1506_v60 = _nw224;
          let _1507_v65;
          let _nw225 = Array((new BigNumber(25)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = (_1466_v29).Merge(_1466_v29);
          _nw225[(_dafny.ONE).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(2)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(3)).toNumber()] = (function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of (function () {
              let _coll68 = new _dafny.Set();
              for (const _compr_68 of _dafny.IntegerRange(new BigNumber(347), new BigNumber(714))) {
                let _1508_v62 = _compr_68;
                if (((new BigNumber(347)).isLessThanOrEqualTo(_1508_v62)) && ((_1508_v62).isLessThan(new BigNumber(714)))) {
                  _coll68.add((_1508_v62).minus(p0));
                }
              }
              return _coll68;
            }()).Elements) {
              let _1509_v61 = _compr_67;
              if ((function () {
                let _coll69 = new _dafny.Set();
                for (const _compr_69 of _dafny.IntegerRange(new BigNumber(347), new BigNumber(714))) {
                  let _1510_v62 = _compr_69;
                  if (((new BigNumber(347)).isLessThanOrEqualTo(_1510_v62)) && ((_1510_v62).isLessThan(new BigNumber(714)))) {
                    _coll69.add((_1510_v62).minus(p0));
                  }
                }
                return _coll69;
              }()).contains(_1509_v61)) {
                _coll67.push([(_1509_v61).minus(_1465_i3),(_this).f28]);
              }
            }
            return _coll67;
          }()).Merge(_1466_v29);
          _nw225[(new BigNumber(4)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1465_i3,(_this).f28);
          _nw225[(new BigNumber(6)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(7)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(8)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(9)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(10)).toNumber()] = (_1466_v29).Merge(_1466_v29);
          _nw225[(new BigNumber(11)).toNumber()] = (_1466_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28));
          _nw225[(new BigNumber(12)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1465_i3,(_this).f28)).Merge((_1466_v29).update(p0, (_this).f28));
          _nw225[(new BigNumber(14)).toNumber()] = _module.__default.fm30((_this).f28, _1435_v4, (_this).f28, globalState);
          _nw225[(new BigNumber(15)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1465_i3,(_this).f28)).Merge(_1466_v29);
          _nw225[(new BigNumber(17)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(18)).toNumber()] = _module.__default.fm30((_this).f28, _dafny.Seq.UnicodeFromString("vtw"), (_this).f28, globalState);
          _nw225[(new BigNumber(19)).toNumber()] = (_1495_v49).dtor_cf27;
          _nw225[(new BigNumber(20)).toNumber()] = (_1466_v29).Merge(function () {
            let _coll70 = new _dafny.Map();
            for (const _compr_70 of (_1466_v29).Keys.Elements) {
              let _1511_v63 = _compr_70;
              if ((_1466_v29).contains(_1511_v63)) {
                _coll70.push([(_1511_v63).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1465_i3)).length)),(_this).f28]);
              }
            }
            return _coll70;
          }());
          _nw225[(new BigNumber(21)).toNumber()] = (function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of _dafny.IntegerRange(new BigNumber(174), new BigNumber(816))) {
              let _1512_v64 = _compr_71;
              if (((new BigNumber(174)).isLessThanOrEqualTo(_1512_v64)) && ((_1512_v64).isLessThan(new BigNumber(816)))) {
                _coll71.push([(_1512_v64).minus((_module.D1.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f28, (_this).f28, (_this).f28),(_this).f28)).length))).dtor_cf12),(_module.D9.create_DC23(true)).dtor_cf38]);
              }
            }
            return _coll71;
          }()).Merge(_1466_v29);
          _nw225[(new BigNumber(22)).toNumber()] = ((_1466_v29).update(p0, (_this).f28)).Merge(_1466_v29);
          _nw225[(new BigNumber(23)).toNumber()] = _1466_v29;
          _nw225[(new BigNumber(24)).toNumber()] = (_1495_v49).dtor_cf27;
          _1507_v65 = _nw225;
          let _index216 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1507_v65).length));
          (_1507_v65)[_index216] = _1466_v29;
          let _1513_v66;
          _1513_v66 = (_this).f27;
          let _1514_v67;
          let _nw226 = Array((new BigNumber(28)).toNumber());
          _nw226[(_dafny.ZERO).toNumber()] = (_this).f27;
          _nw226[(_dafny.ONE).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(2)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(3)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(4)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(5)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(6)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(7)).toNumber()] = (_1513_v66);
          _nw226[(new BigNumber(8)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(9)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(10)).toNumber()] = (_1513_v66);
          _nw226[(new BigNumber(11)).toNumber()] = (_1513_v66);
          _nw226[(new BigNumber(12)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(13)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(14)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(15)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(16)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(17)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(18)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(19)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(20)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(21)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(22)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(23)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(24)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(25)).toNumber()] = (_1513_v66);
          _nw226[(new BigNumber(26)).toNumber()] = (_this).f27;
          _nw226[(new BigNumber(27)).toNumber()] = (_this).f27;
          _1514_v67 = _nw226;
          let _index217 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1507_v65).length));
          let _rhs258 = _module.D9.create_DC24(_module.D9.create_DC22((_1438_v7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1438_v7).length))], !(false), _1514_v67, _1435_v4));
          let _rhs259 = _1505_v59;
          let _rhs260 = _1506_v60;
          let _rhs261 = (_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28)).Merge((_1466_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28)));
          let _lhs217 = _1507_v65;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_1507_v65).length));
          _1504_v58 = _rhs258;
          _1505_v59 = _rhs259;
          _1506_v60 = _rhs260;
          _lhs217[_lhs218] = _rhs261;
          (globalState).f25 = _module.__default.fm42(p0, globalState);
          let _1515_v68;
          _1515_v68 = _module.D0.create_DC3(p0, false, (_this).f28);
          let _1516_v69;
          _1516_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1515_v68,p0);
          let _1517_v70;
          _1517_v70 = _dafny.MultiSet.fromElements((_this).f28, (_this).f28, (_this).f28);
          let _1518_v71;
          let _nw227 = new _module.C3();
          _nw227.__ctor(_1465_i3, _1516_v69, new BigNumber((_1517_v70).cardinality()), (_this).f27, true);
          _1518_v71 = _nw227;
          let _1519_v72;
          _1519_v72 = _dafny.Set.fromElements(_1518_v71);
          (globalState).f15 = ((_1519_v72).Union(_dafny.Set.fromElements(_1518_v71))).IsDisjointFrom(_1519_v72);
        } else {
          (globalState).f15 = !((_this).f28) || ((_this).f28);
          (globalState).f15 = false;
          let _1520_v73;
          let _nw228 = new _module.C1();
          _nw228.__ctor((_this).f28, (_this).f27, (_this).f28);
          _1520_v73 = _nw228;
          (globalState).f21 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0);
          let _1521_v74;
          _1521_v74 = _dafny.MultiSet.fromElements(_module.D1.create_DC6(_1465_i3));
          (globalState).f13 = _module.__default.fm4(_1521_v74, globalState);
        }
      }
      let _1522_v75;
      _1522_v75 = _dafny.Seq.of(_1430_v0);
      let _hi8 = p0;
      for (let _1523_i6 = new BigNumber((_1522_v75).length); _1523_i6.isLessThan(_hi8); _1523_i6 = _1523_i6.plus(_dafny.ONE)) {
        if ((_this).f28) {
          let _1524_v76;
          _1524_v76 = _module.D1.create_DC6(new BigNumber(60));
          let _1525_v77;
          _1525_v77 = _dafny.MultiSet.fromElements(_1524_v76, _1524_v76);
          (globalState).f21 = _module.__default.fm4(_1525_v77, globalState);
          let _1526_v78;
          _1526_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1460_v26,p0);
          let _1527_v79;
          _1527_v79 = _module.D8.create_DC17(_1526_v78);
          let _1528_v80;
          _1528_v80 = _dafny.Seq.of(_1527_v79);
          (globalState).f21 = new BigNumber((_1528_v80).length);
          let _1529_v81;
          let _nw229 = new _module.C2();
          _nw229.__ctor((_this).f32, (_this).f28, (_this).f27, (_this).f28);
          _1529_v81 = _nw229;
          let _rhs262 = _1529_v81;
          let _rhs263 = _1523_i6;
          let _lhs219 = globalState;
          _1529_v81 = _rhs262;
          _lhs219.f13 = _rhs263;
          let _1530_v82;
          let _nw230 = new _module.C5();
          _nw230.__ctor(_1437_v6, _1523_i6, (_this).f27, ((_this).f28) || (true));
          _1530_v82 = _nw230;
          let _1531_v83;
          _1531_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f28)),_1529_v81.f38);
          let _1532_v84;
          _1532_v84 = _dafny.MultiSet.fromElements(_1529_v81.f38);
          let _1533_v86;
          _1533_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1532_v84,p0);
          let _1534_v87;
          _1534_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1523_i6,(_this).f28);
          let _1535_v88;
          _1535_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(-288));
          let _1536_v89;
          let _nw231 = Array((new BigNumber(24)).toNumber());
          _nw231[(_dafny.ZERO).toNumber()] = _1531_v83;
          _nw231[(_dafny.ONE).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(2)).toNumber()] = (_1531_v83).Merge(_1531_v83);
          _nw231[(new BigNumber(3)).toNumber()] = (_1531_v83).Merge(_1531_v83);
          _nw231[(new BigNumber(4)).toNumber()] = ((_1531_v83).update(_1532_v84, false)).Merge(function () {
            let _coll72 = new _dafny.Map();
            for (const _compr_72 of (_1533_v86).Keys.Elements) {
              let _1537_v85 = _compr_72;
              if ((_1533_v86).contains(_1537_v85)) {
                _coll72.push([_1537_v85,_1529_v81.f38]);
              }
            }
            return _coll72;
          }());
          _nw231[(new BigNumber(5)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1532_v84,(_module.D7.create_DC16(_1529_v81.f38, _1534_v87, (_1529_v81).f37, _1529_v81.f38)).dtor_cf29)).update(_1532_v84, _module.__default.fm22(_1529_v81.f38, _1523_i6, globalState));
          _nw231[(new BigNumber(7)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(8)).toNumber()] = ((!(_1529_v81.f38)) ? (_1531_v83) : (_1531_v83));
          _nw231[(new BigNumber(9)).toNumber()] = (_1531_v83).update((_1532_v84).update((_this).f28, _module.__default.abs(p0)), _1529_v81.f38);
          _nw231[(new BigNumber(10)).toNumber()] = ((!((_this).f28)) ? ((_1531_v83).update(_dafny.MultiSet.fromElements((_this).f28), (_this).f28)) : (_1531_v83));
          _nw231[(new BigNumber(11)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1532_v84,!(_1529_v81.f38))).Merge((_1531_v83));
          _nw231[(new BigNumber(12)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(13)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(14)).toNumber()] = (_1531_v83).update(_1532_v84, (_this).f28);
          _nw231[(new BigNumber(15)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(16)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(17)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1532_v84,!(true));
          _nw231[(new BigNumber(19)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(20)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(21)).toNumber()] = _module.__default.fm53(p0, _1431_v1, globalState);
          _nw231[(new BigNumber(22)).toNumber()] = _1531_v83;
          _nw231[(new BigNumber(23)).toNumber()] = _module.__default.fm53(new BigNumber((_1535_v88).length), _1431_v1, globalState);
          _1536_v89 = _nw231;
          let _index218 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_1536_v89).length));
          (_1536_v89)[_index218] = _dafny.Map.Empty.slice().updateUnsafe((_1532_v84).update((_this).f28, _module.__default.abs(_1523_i6)),_1529_v81.f38);
          let _index219 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_1536_v89).length));
          (_1536_v89)[_index219] = (_1531_v83).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1532_v84,(_this).f28)).Merge(_1531_v83));
        } else {
          let _index220 = _module.__default.safeIndex(new BigNumber(53), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index220] = (_this).f28;
          let _1538_v90;
          _1538_v90 = _dafny.Set.fromElements(_1462_v28);
          let _index221 = _module.__default.safeIndex(new BigNumber(53), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index221] = !(_1538_v90).contains(_1462_v28);
          let _1539_v91;
          _1539_v91 = _module.D0.create_DC3(p0, ((_this).f27)[_module.__default.safeIndex(new BigNumber(53), new BigNumber(((_this).f27).length))], (_this).f28);
          _1539_v91 = _1539_v91;
          let _index222 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1433_v3).length));
          (_1433_v3)[_index222] = p0;
          let _index223 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1433_v3).length));
          (_1433_v3)[_index223] = _module.__default.safeModuloInt(new BigNumber(470), (_dafny.ZERO).minus((_dafny.ZERO).minus((_1523_i6).multipliedBy((_dafny.ZERO).minus(p0)))));
          let _index224 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1433_v3).length));
          (_1433_v3)[_index224] = (_1433_v3)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_1433_v3).length))];
          (globalState).f4 = (_dafny.ZERO).minus(p0);
        }
        (globalState).f6 = !(_module.__default.fm40(globalState));
        let _1540_v92;
        let _nw232 = new _module.C5();
        _nw232.__ctor(_1437_v6, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm30((_this).f28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_1541_i7) {
          return (_this).f32;
        }), (_this).f28, globalState)).length)), (((_this).f28) ? ((_this).f27) : ((_this).f27)), (_this).f28);
        _1540_v92 = _nw232;
        let _1542_v93;
        _1542_v93 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)),p0);
        let _1543_v94;
        _1543_v94 = _module.D8.create_DC18(_1523_i6);
        let _1544_v95;
        _1544_v95 = _dafny.Seq.of(p0);
        let _1545_v96;
        let _nw233 = new _module.C5();
        _nw233.__ctor(_1437_v6, (((_1542_v93).contains((_1543_v94).dtor_cf31)) ? ((_1542_v93).get((_1543_v94).dtor_cf31)) : ((_1544_v95)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_1544_v95).length))])), (_this).f27, (_this).f28);
        _1545_v96 = _nw233;
      }
      let _hi9 = p0;
      for (let _1546_i8 = _module.__default.safeDivisionInt(p0, p0); _1546_i8.isLessThan(_hi9); _1546_i8 = _1546_i8.plus(_dafny.ONE)) {
        (globalState).f4 = _1546_i8;
        (globalState).f15 = !((_this).f28);
        let _1547_v97;
        _1547_v97 = _module.D0.create_DC3(new BigNumber((_1430_v0).length), (_this).f28, (_this).f28);
        let _1548_v98;
        _1548_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1547_v97,p0);
        let _1549_v99;
        _1549_v99 = _module.D5.create_DC13(true, _1433_v3, (_this).f28);
        let _1550_v100;
        _1550_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1549_v99).dtor_cf23);
        let _1551_v101;
        _1551_v101 = _module.D3.create_DC10((_this).f28, (((_1550_v100).contains(p0)) ? ((_1550_v100).get(p0)) : ((_this).f28)));
        let _1552_v102;
        let _nw234 = new _module.C3();
        _nw234.__ctor((((_this).f28) ? (p0) : (p0)), (_1548_v98).Merge(_1548_v98), _module.__default.safeModuloInt(p0, _1546_i8), (_this).f27, (_1551_v101).dtor_cf17);
        _1552_v102 = _nw234;
        let _rhs264 = _1435_v4;
        let _rhs265 = (_module.__default.safeModuloInt(_1546_i8, p0)).plus(((_1552_v102).f34).multipliedBy(_1546_i8));
        let _lhs220 = globalState;
        _1435_v4 = _rhs264;
        _lhs220.f4 = _rhs265;
      }
      return;
    }
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f29 = _dafny.ZERO;
      this._f27 = [];
      this._f28 = false;
      this._f30 = false;
      this._f31 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f30, f31, f29, f27, f28) {
      let _this = this;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f29 = f29;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = false;
      let _1553_v0;
      let _nw235 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _1553_v0 = _nw235;
      let _1554_v1;
      _1554_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1553_v0,_1553_v0);
      _1554_v1 = (_1554_v1).update(_1553_v0, _1553_v0);
      let _1555_v2;
      _1555_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f30);
      let _1556_v3;
      _1556_v3 = _dafny.Set.fromElements(_module.__default.fm1(globalState));
      let _1557_v4;
      _1557_v4 = _dafny.Seq.UnicodeFromString("hulb");
      let _1558_v5;
      _1558_v5 = _module.D0.create_DC3(new BigNumber((_1555_v2).length), (_1556_v3).contains(_this.f29), _module.__default.fm2((_1557_v4)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f29)), new BigNumber((_1557_v4).length))], _this.f29, globalState));
      let _source23 = _1558_v5;
      if (_source23.is_DC0) {
        let _1559___mcc_h0 = (_source23).cf0;
        let _1560___mcc_h1 = (_source23).cf1;
        let _1561___mcc_h2 = (_source23).cf2;
        let _1562___mcc_h3 = (_source23).cf3;
        let _1563___mcc_h4 = (_source23).cf4;
        let _1564_cf4 = _1563___mcc_h4;
        let _1565_cf3 = _1562___mcc_h3;
        let _1566_cf2 = _1561___mcc_h2;
        let _1567_cf1 = _1560___mcc_h1;
        let _1568_cf0 = _1559___mcc_h0;
        (globalState).f5 = _this.f29;
        let _1569_v6;
        _1569_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1567_cf1,_1565_cf3);
        let _1570_v7;
        _1570_v7 = _dafny.MultiSet.fromElements(_1567_cf1, (_this).f28, (_this).f28);
        (globalState).f15 = _module.__default.fm2((((_1569_v6).contains(false)) ? ((_1569_v6).get(false)) : (_1565_cf3)), (((_1570_v7).contains((_this).f28)) ? ((_1570_v7).get((_this).f28)) : (_this.f29)), globalState);
        (globalState).f5 = _this.f29;
        let _1571_v8;
        let _nw236 = new _module.C2();
        _nw236.__ctor(_1565_cf3, !(_1564_cf4), (_this).f27, (_this).f30);
        _1571_v8 = _nw236;
        let _1572_v9;
        _1572_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v8,(_this).f28);
        (globalState).f13 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_this.f29, new BigNumber((_dafny.Seq.UnicodeFromString("efreopb")).length)), new BigNumber((_1572_v9).length));
      } else if (_source23.is_DC1) {
        let _1573_v10;
        _1573_v10 = _dafny.Set.fromElements(!((_this).f30), (_this).f30, _module.__default.fm40(globalState), (_this).f30, _module.__default.fm16(_1555_v2, globalState));
        let _index225 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length));
        (_1553_v0)[_index225] = new BigNumber((_1573_v10).length);
        let _1574_v11;
        _1574_v11 = _dafny.Seq.of(_module.__default.safeDivisionInt(_this.f29, _this.f29));
        let _1575_v12;
        let _nw237 = new _module.C0();
        _nw237.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(501)), function (_1576_i0) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }), _this.f29, (_this).f27, (_this).f30);
        _1575_v12 = _nw237;
        let _1577_v13;
        _1577_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1575_v12);
        let _1578_v14;
        _1578_v14 = _dafny.Set.fromElements(_1577_v13, _1577_v13);
        let _index226 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length));
        let _rhs266 = (((_1578_v14).equals(_1578_v14)) ? ((_this.f29).minus(_this.f29)) : (_this.f29));
        let _rhs267 = _dafny.Seq.update(_1574_v11, _module.__default.safeIndex(_this.f29, new BigNumber((_1574_v11).length)), _this.f29);
        let _lhs221 = _1553_v0;
        let _lhs222 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length));
        _lhs221[_lhs222] = _rhs266;
        _1574_v11 = _rhs267;
        let _1579_v15;
        _1579_v15 = _dafny.Seq.of(true);
        let _index227 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length));
        (_1553_v0)[_index227] = (_dafny.ZERO).minus((new BigNumber((_1579_v15).length)).minus(_module.__default.safeDivisionInt(new BigNumber(928), _this.f29)));
        let _1580_v16;
        _1580_v16 = _module.D1.create_DC6(new BigNumber((_1555_v2).length));
        (globalState).f13 = _module.__default.safeModuloInt(_this.f29, _module.__default.fm4(_dafny.MultiSet.fromElements(_1580_v16), globalState));
        let _1581_v17;
        _1581_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f30) ? (((((_this).f31).contains((_1553_v0)[_module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length))])) ? (((_this).f31).get((_1553_v0)[_module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length))])) : ((_1553_v0)[_module.__default.safeIndex(new BigNumber(137), new BigNumber((_1553_v0).length))]))) : (_this.f29)),_module.__default.fm1(globalState));
        (globalState).f13 = (((_1581_v17).contains(_this.f29)) ? ((_1581_v17).get(_this.f29)) : ((_dafny.ZERO).minus(_this.f29)));
      } else if (_source23.is_DC2) {
        let _1582___mcc_h5 = (_source23).cf5;
        let _1583___mcc_h6 = (_source23).cf6;
        let _1584_cf6 = _1583___mcc_h6;
        let _1585_cf5 = _1582___mcc_h5;
        let _1586_v18;
        let _nw238 = Array((new BigNumber(15)).toNumber()).fill(_module.D10.Default());
        _1586_v18 = _nw238;
        let _rhs268 = ((_this.f29).minus(_this.f29)).plus((_this.f29).minus(_this.f29));
        let _rhs269 = _1586_v18;
        let _lhs223 = globalState;
        _lhs223.f13 = _rhs268;
        _1586_v18 = _rhs269;
        let _1587_v19;
        _1587_v19 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1588_v20;
        _1588_v20 = _dafny.Seq.of((_this.f29).isLessThan(_this.f29), (_this).f28, (_this).f28, (new BigNumber((_1557_v4).length)).isLessThan(_this.f29), !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1557_v4, _module.__default.safeIndex(_this.f29, new BigNumber((_1557_v4).length)), _1587_v19), _1557_v4)));
        (globalState).f15 = (_1588_v20)[_module.__default.safeIndex(_this.f29, new BigNumber((_1588_v20).length))];
        let _1589_v21;
        _1589_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30);
        let _1590_v22;
        _1590_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1584_cf6,(_this).f27);
        _1589_v21 = (_1589_v21).update((_this.f29).multipliedBy(new BigNumber((_1590_v22).length)), (_this).f28);
        let _1591_v23;
        let _nw239 = Array((new BigNumber(21)).toNumber());
        _nw239[(_dafny.ZERO).toNumber()] = (_this).f27;
        _nw239[(_dafny.ONE).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(2)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(3)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(4)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(5)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(6)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(7)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(8)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(9)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(10)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(11)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(12)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(13)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(14)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(15)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(16)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(17)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(18)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(19)).toNumber()] = (_this).f27;
        _nw239[(new BigNumber(20)).toNumber()] = (_this).f27;
        _1591_v23 = _nw239;
        let _1592_v24;
        _1592_v24 = _module.D9.create_DC22(_1553_v0, (_this).f28, _1591_v23, _1584_cf6);
        _1557_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat((_1592_v24).dtor_cf37, _1584_cf6), _1557_v4);
      } else if (_source23.is_DC3) {
        let _1593___mcc_h7 = (_source23).cf7;
        let _1594___mcc_h8 = (_source23).cf8;
        let _1595___mcc_h9 = (_source23).cf9;
        let _1596_cf9 = _1595___mcc_h9;
        let _1597_cf8 = _1594___mcc_h8;
        let _1598_cf7 = _1593___mcc_h7;
        let _1599_v25;
        let _nw240 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _1599_v25 = _nw240;
        _1599_v25 = _1599_v25;
        (globalState).f17 = (((_this).f28) ? ((_this).f30) : (_1596_cf9));
        (globalState).f5 = _module.__default.fm36(globalState);
        let _1600_v26;
        let _init31 = function (_1601_i1) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-313)), function (_1602_i2) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          });
        };
        let _nw241 = Array((new BigNumber(2)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw241.length); _i0_31++) {
          _nw241[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1600_v26 = _nw241;
        let _index228 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1600_v26).length));
        (_1600_v26)[_index228] = _1557_v4;
        let _1603_v27;
        _1603_v27 = _dafny.MultiSet.fromElements(_1553_v0);
        let _1604_v28;
        _1604_v28 = new _dafny.CodePoint('a'.codePointAt(0));
        let _index229 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1600_v26).length));
        let _rhs270 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fs"), _1557_v4), _dafny.Seq.Concat(_1557_v4, _module.__default.fm18(_1596_cf9, _this.f29, _1604_v28, globalState)));
        let _rhs271 = _1598_cf7;
        let _rhs272 = (_this).f28;
        let _rhs273 = (_1603_v27).Difference(_1603_v27);
        let _lhs224 = _1600_v26;
        let _lhs225 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1600_v26).length));
        let _lhs226 = globalState;
        _lhs224[_lhs225] = _rhs270;
        _lhs226.f21 = _rhs271;
        r2 = _rhs272;
        _1603_v27 = _rhs273;
      } else {
        let _1605___mcc_h10 = (_source23).cf10;
        let _1606_cf10 = _1605___mcc_h10;
        let _1607_v29;
        _1607_v29 = _dafny.Seq.of((_this).f30, (_this).f30, (_this).f30);
        let _1608_v30;
        _1608_v30 = _dafny.Seq.of(_1553_v0);
        let _1609_v31;
        let _nw242 = new _module.C5();
        _nw242.__ctor(_1608_v30, _this.f29, (_this).f27, (_this).f30);
        _1609_v31 = _nw242;
        let _1610_v32;
        _1610_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1609_v31,(((_this).f28) ? (_dafny.Seq.of((_this).f30, (_this).f28)) : (_1607_v29)));
        _1607_v29 = (((_1610_v32).contains(_1609_v31)) ? ((_1610_v32).get(_1609_v31)) : (_1607_v29));
        (globalState).f15 = (_this).f28;
        let _1611_v33;
        let _nw243 = new _module.C7();
        _nw243.__ctor((_this).f27, (_this).f28, _this.f29);
        _1611_v33 = _nw243;
        let _1612_v34;
        _1612_v34 = _dafny.Set.fromElements(_1557_v4, _1557_v4);
        let _nw244 = new _module.C6();
        _nw244.__ctor((_dafny.ZERO).minus(_1611_v33.f29), (_this).f27, (_1612_v34).IsSubsetOf(_1612_v34));
        _1611_v33 = _nw244;
        (globalState).f13 = _1611_v33.f29;
      }
      let _hi10 = _this.f29;
      for (let _1613_i3 = new BigNumber((_dafny.Seq.Concat(_1557_v4, _1557_v4)).length); _1613_i3.isLessThan(_hi10); _1613_i3 = _1613_i3.plus(_dafny.ONE)) {
        let _1614_v35;
        _1614_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1613_i3,_this.f29);
        let _1615_v36;
        let _nw245 = new _module.C0();
        _nw245.__ctor(_1557_v4, new BigNumber((_1614_v35).length), (_this).f27, !((_this).f28));
        _1615_v36 = _nw245;
        let _1616_v37;
        let _init32 = ((_1617_v3) => function (_1618_i4) {
          return _1617_v3;
        })(_1556_v3);
        let _nw246 = Array((new BigNumber(13)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw246.length); _i0_32++) {
          _nw246[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1616_v37 = _nw246;
        let _index230 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length));
        (_1616_v37)[_index230] = (_1556_v3).Intersect(_1556_v3);
        let _1619_v38;
        _1619_v38 = _dafny.Seq.of(_this.f29);
        let _1620_v39;
        _1620_v39 = _dafny.Seq.of(_1619_v38);
        let _1621_v40;
        _1621_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_1620_v39),(_this).f28);
        let _1622_v41;
        _1622_v41 = _dafny.Seq.of((_this).f30, (_this).f30, (_this).f30);
        let _1623_v42;
        _1623_v42 = _dafny.Seq.of(new BigNumber((_1621_v40).length), _1613_i3, new BigNumber((_1622_v41).length));
        let _index231 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length));
        let _rhs274 = _1615_v36;
        let _rhs275 = _1556_v3;
        let _rhs276 = ((_1623_v42)[_module.__default.safeIndex(_1613_i3, new BigNumber((_1623_v42).length))]).minus(_this.f29);
        let _lhs227 = _1616_v37;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length));
        let _lhs229 = globalState;
        _1615_v36 = _rhs274;
        _lhs227[_lhs228] = _rhs275;
        _lhs229.f21 = _rhs276;
        _1623_v42 = _1619_v38;
        let _1624_v43;
        let _nw247 = Array((new BigNumber(7)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = _1557_v4;
        _nw247[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("uhkbfbg");
        _nw247[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1557_v4, _dafny.Seq.UnicodeFromString("nnmatc"));
        _nw247[(new BigNumber(3)).toNumber()] = (_1615_v36).f36;
        _nw247[(new BigNumber(4)).toNumber()] = _module.__default.fm9((_this).f28, globalState);
        _nw247[(new BigNumber(5)).toNumber()] = ((_module.__default.fm40(globalState)) ? (_1557_v4) : (_1557_v4));
        _nw247[(new BigNumber(6)).toNumber()] = _1557_v4;
        _1624_v43 = _nw247;
        _1624_v43 = _1624_v43;
        if ((_this).f30) {
          let _1625_v44;
          _1625_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1558_v5,new BigNumber(20));
          let _1626_v45;
          _1626_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1625_v44);
          let _1627_v46;
          let _nw248 = new _module.C3();
          _nw248.__ctor(new BigNumber((_1557_v4).length), (((_1626_v45).contains(!((_this).f28))) ? ((_1626_v45).get(!((_this).f28))) : (_1625_v44)), new BigNumber(568), (_this).f27, (_this).f28);
          _1627_v46 = _nw248;
          let _1628_v47;
          _1628_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1627_v46.f29,_1627_v46);
          let _1629_v48;
          _1629_v48 = _module.D1.create_DC6(_this.f29);
          let _1630_v49;
          _1630_v49 = _dafny.MultiSet.fromElements(_1629_v48);
          let _1631_v50;
          let _nw249 = Array((new BigNumber(16)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1627_v46;
          _nw249[(_dafny.ONE).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(2)).toNumber()] = (((_this).f30) ? (_1627_v46) : (_1627_v46));
          _nw249[(new BigNumber(3)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(4)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(5)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(6)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(7)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(8)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(9)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(10)).toNumber()] = (((_1628_v47).contains(_module.__default.fm4(_1630_v49, globalState))) ? ((_1628_v47).get(_module.__default.fm4(_1630_v49, globalState))) : (_1627_v46));
          _nw249[(new BigNumber(11)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(12)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(13)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(14)).toNumber()] = _1627_v46;
          _nw249[(new BigNumber(15)).toNumber()] = _1627_v46;
          _1631_v50 = _nw249;
          let _1632_v51;
          let _nw250 = Array((_dafny.ONE).toNumber()).fill(_module.D0.Default());
          _1632_v51 = _nw250;
          let _index232 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1632_v51).length));
          (_1632_v51)[_index232] = _1558_v5;
          let _index233 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1632_v51).length));
          let _rhs277 = _1631_v50;
          let _rhs278 = _1558_v5;
          let _lhs230 = _1632_v51;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1632_v51).length));
          _1631_v50 = _rhs277;
          _lhs230[_lhs231] = _rhs278;
          (globalState).f4 = ((_1613_i3).multipliedBy((_dafny.ZERO).minus(_1613_i3))).plus(_1627_v46.f29);
          let _1633_v52;
          _1633_v52 = _1619_v38;
          let _1634_v53;
          _1634_v53 = _dafny.MultiSet.fromElements((_this).f30);
          _1619_v38 = _dafny.Seq.Concat(_dafny.Seq.update((_1633_v52), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_1615_v36).f36).length)), new BigNumber(((_1633_v52)).length)), new BigNumber((_1634_v53).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_1635_v36) => function (_1636_i5) {
            return new BigNumber(((_1635_v36).f36).length);
          })(_1615_v36)));
          let _index234 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1553_v0).length));
          (_1553_v0)[_index234] = new BigNumber(310);
          let _index235 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1553_v0).length));
          (_1553_v0)[_index235] = _this.f29;
          let _index236 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1624_v43).length));
          (_1624_v43)[_index236] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), function (_1637_i6) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          let _index237 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1624_v43).length));
          (_1624_v43)[_index237] = (_1615_v36).f36;
        } else {
          let _index238 = _module.__default.safeIndex(new BigNumber(474), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index238] = (_this).f28;
          let _index239 = _module.__default.safeIndex(new BigNumber(474), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index239] = (_this).f28;
          (globalState).f4 = (new BigNumber(-30)).plus(new BigNumber(((_1616_v37)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length))]).length));
          let _1638_v54;
          let _init33 = function (_1639_i7) {
            return !(true);
          };
          let _nw251 = Array((new BigNumber(24)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw251.length); _i0_33++) {
            _nw251[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1638_v54 = _nw251;
          _1638_v54 = _1638_v54;
          r1 = !(((_this).f27)[_module.__default.safeIndex(new BigNumber(474), new BigNumber(((_this).f27).length))]);
          let _1640_v55;
          _1640_v55 = _dafny.Seq.of(_1556_v3, _1556_v3, (_1556_v3).Intersect((_1616_v37)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length))]));
          _1640_v55 = _dafny.Seq.of(_1556_v3, (_1616_v37)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1616_v37).length))]);
        }
      }
      let _1641_v56;
      _1641_v56 = _dafny.Set.fromElements((_this).f28, !((_this).f30), (_this).f30, (_this).f28, (_this).f28);
      _1641_v56 = _1641_v56;
      let _1642_v57;
      _1642_v57 = new _dafny.CodePoint('c'.codePointAt(0));
      let _1643_v58;
      _1643_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1642_v57);
      let _index240 = _module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length));
      ((_this).f27)[_index240] = (_this).f30;
      let _1644_v59;
      _1644_v59 = _dafny.Seq.of(!((_this).f30) || ((_this).f30), _module.__default.fm16(_1555_v2, globalState), ((_dafny.ZERO).minus(_this.f29)).isLessThan(_this.f29));
      let _index241 = _module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length));
      let _rhs279 = _this.f29;
      let _rhs280 = _1643_v58;
      let _rhs281 = (_1644_v59)[_module.__default.safeIndex(_this.f29, new BigNumber((_1644_v59).length))];
      let _lhs232 = globalState;
      let _lhs233 = (_this).f27;
      let _lhs234 = _module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length));
      _lhs232.f5 = _rhs279;
      _1643_v58 = _rhs280;
      _lhs233[_lhs234] = _rhs281;
      let _hi11 = (new BigNumber(585)).multipliedBy(_this.f29);
      for (let _1645_i8 = new BigNumber((_1557_v4).length); _1645_i8.isLessThan(_hi11); _1645_i8 = _1645_i8.plus(_dafny.ONE)) {
        let _index242 = _module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length));
        let _rhs282 = (_this).f28;
        let _rhs283 = _dafny.Seq.Concat(_1557_v4, _1557_v4);
        let _rhs284 = (_this.f29).plus(_module.__default.fm36(globalState));
        let _rhs285 = ((_this).f27)[_module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length))];
        let _lhs235 = (_this).f27;
        let _lhs236 = _module.__default.safeIndex(new BigNumber(639), new BigNumber(((_this).f27).length));
        let _lhs237 = globalState;
        let _lhs238 = globalState;
        _lhs235[_lhs236] = _rhs282;
        _1557_v4 = _rhs283;
        _lhs237.f5 = _rhs284;
        _lhs238.f17 = _rhs285;
        let _1646_v60;
        let _nw252 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1646_v60 = _nw252;
        let _1647_v61;
        let _nw253 = Array((new BigNumber(6)).toNumber());
        _nw253[(_dafny.ZERO).toNumber()] = _1646_v60;
        _nw253[(_dafny.ONE).toNumber()] = _1646_v60;
        _nw253[(new BigNumber(2)).toNumber()] = _1646_v60;
        _nw253[(new BigNumber(3)).toNumber()] = _1646_v60;
        _nw253[(new BigNumber(4)).toNumber()] = _1646_v60;
        _nw253[(new BigNumber(5)).toNumber()] = _1646_v60;
        _1647_v61 = _nw253;
        let _1648_v62;
        _1648_v62 = _module.D9.create_DC22(_1553_v0, !((_this).f30), _1647_v61, _1557_v4);
        r2 = (_1648_v62).dtor_cf35;
        r1 = (_this).f28;
        let _1649_v63;
        _1649_v63 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("odjwsypnq"), _1557_v4);
        let _1650_v64;
        _1650_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1649_v63);
        _1650_v64 = (_1650_v64).update((_this).f30, _1649_v63);
      }
      r0 = (_this.f29).minus(_this.f29);
      r1 = (_this).f30;
      r2 = ((_this).f30) === ((_this).f30);
      r3 = (_this).f30;
      return [r0, r1, r2, r3];
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1651_v0;
      _1651_v0 = _dafny.Seq.UnicodeFromString("wfqmokab");
      _1651_v0 = _dafny.Seq.Concat(_module.__default.fm32(_this.f29, (_this).f30, globalState), _1651_v0);
      (globalState).f13 = _this.f29;
      let _1652_v1;
      let _nw254 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1652_v1 = _nw254;
      let _1653_v2;
      let _nw255 = Array((new BigNumber(17)).toNumber()).fill([]);
      _1653_v2 = _nw255;
      let _1654_v3;
      _1654_v3 = _module.D9.create_DC22(_1652_v1, (_this).f30, _1653_v2, _1651_v0);
      let _1655_v4;
      _1655_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29);
      let _1656_v5;
      _1656_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_1655_v4);
      let _1657_v6;
      _1657_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1656_v5).length),(_this).f28);
      let _1658_v7;
      let _nw256 = Array((new BigNumber(17)).toNumber());
      _nw256[(_dafny.ZERO).toNumber()] = (_this).f30;
      _nw256[(_dafny.ONE).toNumber()] = (_this).f28;
      _nw256[(new BigNumber(2)).toNumber()] = (_this).f28;
      _nw256[(new BigNumber(3)).toNumber()] = (_1654_v3).dtor_cf35;
      _nw256[(new BigNumber(4)).toNumber()] = (_this).f28;
      _nw256[(new BigNumber(5)).toNumber()] = false;
      _nw256[(new BigNumber(6)).toNumber()] = true;
      _nw256[(new BigNumber(7)).toNumber()] = !(_this.f29).isEqualTo(new BigNumber(-922));
      _nw256[(new BigNumber(8)).toNumber()] = (_this).f30;
      _nw256[(new BigNumber(9)).toNumber()] = (_this).f30;
      _nw256[(new BigNumber(10)).toNumber()] = (_this).f28;
      _nw256[(new BigNumber(11)).toNumber()] = (_this).f30;
      _nw256[(new BigNumber(12)).toNumber()] = (_module.D9.create_DC22(_1652_v1, (_this).f28, _1653_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-72)), function (_1659_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}))).dtor_cf35;
      _nw256[(new BigNumber(13)).toNumber()] = (_this).f30;
      _nw256[(new BigNumber(14)).toNumber()] = (_this).f30;
      _nw256[(new BigNumber(15)).toNumber()] = (((_1657_v6).contains(_this.f29)) ? ((_1657_v6).get(_this.f29)) : (!(false)));
      _nw256[(new BigNumber(16)).toNumber()] = true;
      _1658_v7 = _nw256;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1658_v7).length))) {
        let _1660_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1660_i0)) && ((_1660_i0).isLessThan(new BigNumber((_1658_v7).length))))) {
          (_1658_v7)[(_1660_i0)] = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(function () {
            let _coll73 = new _dafny.Set();
            for (const _compr_73 of (_dafny.Set.fromElements(_module.D7.create_DC16((_this).f30, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(747), _this.f29, _this.f29)).cardinality()))).length),(_this).f28), new _dafny.CodePoint('q'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('k'.codePointAt(0)), !((_this).f28)), _module.D7.create_DC16((_this).f30, (_1657_v6).update(_this.f29, false), new _dafny.CodePoint('g'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('v'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('b'.codePointAt(0)), (_this).f28))).Elements) {
              let _1661_v8 = _compr_73;
              if ((_dafny.Set.fromElements(_module.D7.create_DC16((_this).f30, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(747), _this.f29, _this.f29)).cardinality()))).length),(_this).f28), new _dafny.CodePoint('q'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('k'.codePointAt(0)), !((_this).f28)), _module.D7.create_DC16((_this).f30, (_1657_v6).update(_this.f29, false), new _dafny.CodePoint('g'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('v'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('b'.codePointAt(0)), (_this).f28))).contains(_1661_v8)) {
                _coll73.add(_1661_v8);
              }
            }
            return _coll73;
          }()), _dafny.Seq.of(_dafny.Set.fromElements(_module.D7.create_DC16((_this).f30, (function () {
  let _coll74 = new _dafny.Map();
  for (const _compr_74 of ((_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_1662_i2) {
    return _this.f29;
  }))).Elements) {
    let _1663_v9 = _compr_74;
    if (_dafny.Seq.contains((_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_1662_i2) {
      return _this.f29;
    })), _1663_v9)) {
      _coll74.push([_module.__default.safeDivisionInt(_1663_v9, _this.f29),(_this).f30]);
    }
  }
  return _coll74;
}()).update(_this.f29, (_this).f30), new _dafny.CodePoint('k'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('j'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('h'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('o'.codePointAt(0)), (_this).f28)))), (_dafny.Set.fromElements(_module.D7.create_DC16((_this).f30, _1657_v6, new _dafny.CodePoint('b'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30), new _dafny.CodePoint('f'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('y'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f28, _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30), new _dafny.CodePoint('k'.codePointAt(0)), (_this).f30), _module.D7.create_DC16((_this).f28, _1657_v6, new _dafny.CodePoint('p'.codePointAt(0)), (_this).f28))).Difference(_dafny.Set.fromElements(_module.D7.create_DC16(true, (_1657_v6).update(new BigNumber((_dafny.MultiSet.fromElements((_this).f28, (_this).f28)).cardinality()), (_this).f28), new _dafny.CodePoint('d'.codePointAt(0)), (_this).f28), _module.D7.create_DC16((_this).f28, (_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f28)).update(new BigNumber(558), (_this).f28), new _dafny.CodePoint('m'.codePointAt(0)), true))));
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1652_v1).length))) {
        let _1664_i3 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1664_i3)) && ((_1664_i3).isLessThan(new BigNumber((_1652_v1).length))))) {
          (_1652_v1)[(_1664_i3)] = (_1664_i3).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1655_v4,(_this).f28)).length));
        }
      }
      r0 = (_this.f29).isEqualTo(_this.f29);
      let _1665_v10;
      _1665_v10 = _dafny.Set.fromElements(_this.f29);
      _1665_v10 = (_1665_v10).Union(function () {
        let _coll75 = new _dafny.Set();
        for (const _compr_75 of (_dafny.Seq.of(_this.f29, new BigNumber(250))).Elements) {
          let _1666_v11 = _compr_75;
          if (_dafny.Seq.contains(_dafny.Seq.of(_this.f29, new BigNumber(250)), _1666_v11)) {
            _coll75.add(_module.__default.safeDivisionInt(_1666_v11, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber(44))).length)));
          }
        }
        return _coll75;
      }());
      r0 = false;
      r1 = _this.f29;
      r2 = !((_this).f30) || (!(_this.f29).isEqualTo(_this.f29));
      let _1667_v12;
      _1667_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f30);
      let _1668_v13;
      _1668_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1667_v12,(_this).f28);
      let _1669_v14;
      _1669_v14 = _dafny.Seq.of(_this.f29);
      r3 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1668_v13).length)), _1669_v14)).length);
      return [r0, r1, r2, r3];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
