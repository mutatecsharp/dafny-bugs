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
      return _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(!(true), true, !(true))));
    };
    static fm1(p0, p1, p2, globalState) {
      if (false) {
        return true;
      } else {
        return (!(!(!(!(!(!(!(false)))))))) && (true);
      }
    };
    static fm2(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, true, false)).length)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(463),false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(308), new BigNumber((_dafny.Set.fromElements(new BigNumber(930), new BigNumber(305))).length), new BigNumber((_dafny.Seq.UnicodeFromString("kie")).length), new BigNumber(941), new BigNumber((_dafny.MultiSet.fromElements(_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(true))).length), true, true))).cardinality()))).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-918),true)));
    };
    static fm3(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(true, true)).Difference(_dafny.MultiSet.fromElements(false))).Intersect(_dafny.MultiSet.fromElements(true));
    };
    static fm4(globalState) {
      if (false) {
        return new BigNumber(-885);
      } else {
        return new BigNumber((_dafny.Seq.of(true, !(false), true, false)).length);
      }
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return (((_module.D6.create_DC18(_dafny.Set.fromElements(false))).dtor_cf27).Union(_dafny.Set.fromElements(false, true))).Difference((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(!(true), !(!(false)))));
    };
    static fm6(p0, p1, globalState) {
      let _source0 = _module.D2.create_DC7(true, new BigNumber(462));
      if (_source0.is_DC7) {
        let _0___mcc_h0 = (_source0).cf11;
        let _1___mcc_h1 = (_source0).cf12;
        let _2_cf12 = _1___mcc_h1;
        let _3_cf11 = _0___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(new BigNumber(127), true, _3_cf11),_3_cf11);
      } else if (_source0.is_DC6) {
        let _4___mcc_h2 = (_source0).cf10;
        let _5_cf10 = _4___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wdogpmi")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_6_i0) {
    return new _dafny.CodePoint('e'.codePointAt(0));
  })).length))).Keys.Elements) {
    let _7_v0 = _compr_0;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wdogpmi")).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_6_i0) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    })).length))).contains(_7_v0)) {
      _coll0.push([_module.__default.safeModuloInt(_7_v0, new BigNumber(857)),true]);
    }
  }
  return _coll0;
}()).length)),new BigNumber(969))).length), false, !(false)),false);
      } else {
        let _8___mcc_h3 = (_source0).cf13;
        let _9_cf13 = _8___mcc_h3;
        return function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Set.fromElements(_module.D1.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("ogqm")).length), true, true), _module.D1.create_DC2(new BigNumber(-484), !(false), true))).Elements) {
            let _10_v1 = _compr_1;
            if ((_dafny.Set.fromElements(_module.D1.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("ogqm")).length), true, true), _module.D1.create_DC2(new BigNumber(-484), !(false), true))).contains(_10_v1)) {
              _coll1.push([_10_v1,true]);
            }
          }
          return _coll1;
        }();
      }
    };
    static fm7(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(589), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(645))).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-720)), function (_11_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-525)), function (_12_i1) {
        return new BigNumber(44);
      })), _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, !(!(!(true))))).cardinality()))), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mqtumvw")).length))));
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mrikrvux"), (_module.D4.create_DC13(_dafny.Seq.UnicodeFromString("ur"))).dtor_cf20), _dafny.Seq.UnicodeFromString("hytv"));
    };
    static fm9(p0, globalState) {
      return new _dafny.CodePoint('t'.codePointAt(0));
    };
    static fm10(p0, p1, p2, globalState) {
      return (_module.D9.create_DC24(_dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(878)))))).dtor_cf38;
    };
    static fm11(globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)));
    };
    static fm12(globalState) {
      return _module.D4.create_DC13(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("urylnfh"), _dafny.Seq.UnicodeFromString("mimbr")));
    };
    static fm13(p0, globalState) {
      return function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-208), new BigNumber(303))) {
          let _13_v0 = _compr_2;
          if (((new BigNumber(-208)).isLessThanOrEqualTo(_13_v0)) && ((_13_v0).isLessThan(new BigNumber(303)))) {
            _coll2.add(_module.__default.safeDivisionInt(_13_v0, new BigNumber(869)));
          }
        }
        return _coll2;
      }();
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _14_v0;
      _14_v0 = _dafny.MultiSet.fromElements(true);
      let _15_v1;
      _15_v1 = _dafny.Seq.of(_14_v0);
      if ((((_15_v1)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_15_v1).length))]).Union(_14_v0)).IsProperSubsetOf((_dafny.MultiSet.fromElements(p2, !(!(p2)), p2, false)).Intersect(_14_v0))) {
        if (!(_dafny.MultiSet.fromElements(p1)).contains(_module.__default.safeDivisionInt(p1, p1))) {
          let _16_v2;
          _16_v2 = _dafny.Set.fromElements(!(p0), p0);
          r1 = _module.__default.safeDivisionInt((p1).plus(p1), new BigNumber((_16_v2).length));
          let _17_v3;
          let _nw0 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _17_v3 = _nw0;
          let _index0 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_17_v3).length));
          (_17_v3)[_index0] = p1;
          let _index1 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_17_v3).length));
          (_17_v3)[_index1] = p1;
          let _index2 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_17_v3).length));
          (_17_v3)[_index2] = p1;
          let _18_v4;
          let _init0 = ((_19_p0) => function (_20_i0) {
            return _19_p0;
          })(p0);
          let _nw1 = Array((new BigNumber(10)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
            _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _18_v4 = _nw1;
          let _21_v5;
          let _nw2 = new _module.C0();
          _nw2.__ctor(_18_v4);
          _21_v5 = _nw2;
          r2 = p2;
        } else {
          let _22_v6;
          _22_v6 = _dafny.Seq.of(p0, p0);
          let _23_v7;
          _23_v7 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_22_v6)).cardinality()));
          let _24_v8;
          _24_v8 = _module.D3.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(835)), ((_25_p1) => function (_26_i1) {
  return _25_p1;
})(p1)), p2);
          let _27_v9;
          _27_v9 = _dafny.Seq.UnicodeFromString("dwnjkq");
          let _28_v10;
          _28_v10 = _dafny.Seq.of(_23_v7);
          let _29_v11;
          _29_v11 = new _dafny.CodePoint('r'.codePointAt(0));
          let _30_v12;
          let _nw3 = Array((new BigNumber(18)).toNumber());
          _nw3[(_dafny.ZERO).toNumber()] = _23_v7;
          _nw3[(_dafny.ONE).toNumber()] = (_24_v8).dtor_cf17;
          _nw3[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-640)), ((_31_p1) => function (_32_i2) {
            return _31_p1;
          })(p1)), _23_v7);
          _nw3[(new BigNumber(3)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(p1, p1, p1, p1, p1);
          _nw3[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(p1, p1, p1, p1);
          _nw3[(new BigNumber(6)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), ((_33_p1) => function (_34_i3) {
            return _33_p1;
          })(p1)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), ((_35_p1) => function (_36_i3) {
            return _35_p1;
          })(p1))).length)), p1);
          _nw3[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_23_v7, _module.__default.safeIndex(p1, new BigNumber((_23_v7).length)), new BigNumber((_27_v9).length));
          _nw3[(new BigNumber(9)).toNumber()] = (_28_v10)[_module.__default.safeIndex(p1, new BigNumber((_28_v10).length))];
          _nw3[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_23_v7, _module.__default.safeIndex(_module.__default.fm4(globalState), new BigNumber((_23_v7).length)), new BigNumber(-246));
          _nw3[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), function (_37_i4) {
            return new BigNumber(882);
          });
          _nw3[(new BigNumber(12)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(13)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(14)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(15)).toNumber()] = _23_v7;
          _nw3[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_23_v7, _23_v7);
          _nw3[(new BigNumber(17)).toNumber()] = _module.__default.fm7(p2, _29_v11, new BigNumber(349), globalState);
          _30_v12 = _nw3;
          let _index3 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_30_v12).length));
          (_30_v12)[_index3] = _23_v7;
          let _index4 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_30_v12).length));
          (_30_v12)[_index4] = _23_v7;
          let _38_v13;
          _38_v13 = _module.D3.create_DC9((_30_v12)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_30_v12).length))]);
          _38_v13 = _38_v13;
          let _39_v14;
          let _nw4 = Array((new BigNumber(8)).toNumber());
          _nw4[(_dafny.ZERO).toNumber()] = _module.__default.fm9(p2, globalState);
          _nw4[(_dafny.ONE).toNumber()] = _29_v11;
          _nw4[(new BigNumber(2)).toNumber()] = _29_v11;
          _nw4[(new BigNumber(3)).toNumber()] = ((p2) ? (_29_v11) : (new _dafny.CodePoint('v'.codePointAt(0))));
          _nw4[(new BigNumber(4)).toNumber()] = _29_v11;
          _nw4[(new BigNumber(5)).toNumber()] = _29_v11;
          _nw4[(new BigNumber(6)).toNumber()] = _29_v11;
          _nw4[(new BigNumber(7)).toNumber()] = _29_v11;
          _39_v14 = _nw4;
          let _index5 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_39_v14).length));
          (_39_v14)[_index5] = _29_v11;
          let _40_v15;
          let _init1 = ((_41_p1) => function (_42_i5) {
            return _module.__default.safeModuloInt(_42_i5, _41_p1);
          })(p1);
          let _nw5 = Array((new BigNumber(19)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
            _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _40_v15 = _nw5;
          let _index6 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_39_v14).length));
          let _rhs0 = p0;
          let _rhs1 = _29_v11;
          let _rhs2 = _40_v15;
          let _lhs0 = _39_v14;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_39_v14).length));
          r2 = _rhs0;
          _lhs0[_lhs1] = _rhs1;
          _40_v15 = _rhs2;
          let _43_v16;
          let _init2 = ((_44_v0) => function (_45_i6) {
            return _44_v0;
          })(_14_v0);
          let _nw6 = Array((new BigNumber(29)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _43_v16 = _nw6;
          let _index7 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_43_v16).length));
          (_43_v16)[_index7] = _14_v0;
          let _index8 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_43_v16).length));
          (_43_v16)[_index8] = _14_v0;
          let _46_v17;
          let _nw7 = Array((new BigNumber(16)).toNumber()).fill(false);
          _46_v17 = _nw7;
          let _47_v18;
          let _nw8 = new _module.C0();
          _nw8.__ctor(_46_v17);
          _47_v18 = _nw8;
        }
        let _48_v19;
        _48_v19 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),p0);
        let _49_v20;
        _49_v20 = _dafny.Seq.UnicodeFromString("fxdguiwrc");
        let _50_v21;
        _50_v21 = _module.D4.create_DC13(_49_v20);
        let _51_v22;
        _51_v22 = _dafny.Seq.of(p0, false, p0);
        let _52_v23;
        _52_v23 = _dafny.Map.Empty.slice().updateUnsafe(_51_v22,p0);
        let _53_v24;
        _53_v24 = _module.D3.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_54_p1) => function (_55_i7) {
  return _54_p1;
})(p1)), p0);
        let _56_v25;
        let _nw9 = Array((new BigNumber(29)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = false;
        _nw9[(_dafny.ONE).toNumber()] = p0;
        _nw9[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_48_v19, (_50_v21).dtor_cf20, _dafny.Seq.UnicodeFromString("lenrsoqb"), globalState);
        _nw9[(new BigNumber(3)).toNumber()] = p2;
        _nw9[(new BigNumber(4)).toNumber()] = p0;
        _nw9[(new BigNumber(5)).toNumber()] = p0;
        _nw9[(new BigNumber(6)).toNumber()] = p0;
        _nw9[(new BigNumber(7)).toNumber()] = p2;
        _nw9[(new BigNumber(8)).toNumber()] = p0;
        _nw9[(new BigNumber(9)).toNumber()] = p0;
        _nw9[(new BigNumber(10)).toNumber()] = p2;
        _nw9[(new BigNumber(11)).toNumber()] = !(!(p2));
        _nw9[(new BigNumber(12)).toNumber()] = true;
        _nw9[(new BigNumber(13)).toNumber()] = !(p0);
        _nw9[(new BigNumber(14)).toNumber()] = false;
        _nw9[(new BigNumber(15)).toNumber()] = p2;
        _nw9[(new BigNumber(16)).toNumber()] = p2;
        _nw9[(new BigNumber(17)).toNumber()] = p0;
        _nw9[(new BigNumber(18)).toNumber()] = (((_52_v23).contains(_51_v22)) ? ((_52_v23).get(_51_v22)) : (!(p0)));
        _nw9[(new BigNumber(19)).toNumber()] = true;
        _nw9[(new BigNumber(20)).toNumber()] = p2;
        _nw9[(new BigNumber(21)).toNumber()] = !(p0);
        _nw9[(new BigNumber(22)).toNumber()] = (_53_v24).dtor_cf18;
        _nw9[(new BigNumber(23)).toNumber()] = p0;
        _nw9[(new BigNumber(24)).toNumber()] = p2;
        _nw9[(new BigNumber(25)).toNumber()] = p0;
        _nw9[(new BigNumber(26)).toNumber()] = p2;
        _nw9[(new BigNumber(27)).toNumber()] = !(p2);
        _nw9[(new BigNumber(28)).toNumber()] = p0;
        _56_v25 = _nw9;
        let _57_v26;
        let _nw10 = new _module.C0();
        _nw10.__ctor(_56_v25);
        _57_v26 = _nw10;
        let _58_v27;
        let _nw11 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _58_v27 = _nw11;
        let _59_v28;
        let _init3 = ((_60_p1) => function (_61_i8) {
          return (_61_i8).plus(_60_p1);
        })(p1);
        let _nw12 = Array((new BigNumber(21)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw12.length); _i0_3++) {
          _nw12[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _59_v28 = _nw12;
        let _62_v29;
        _62_v29 = _dafny.Seq.of(_59_v28);
        let _index9 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_58_v27).length));
        (_58_v27)[_index9] = _62_v29;
        let _index10 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_58_v27).length));
        (_58_v27)[_index10] = _62_v29;
        let _63_v30;
        let _nw13 = new _module.C0();
        _nw13.__ctor((_57_v26).f15);
        _63_v30 = _nw13;
        (globalState).f0 = _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(p1));
      } else {
        if (!(p0)) {
          let _64_v31;
          let _nw14 = Array((new BigNumber(16)).toNumber()).fill(false);
          _64_v31 = _nw14;
          let _index11 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_64_v31).length));
          (_64_v31)[_index11] = p2;
          let _index12 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_64_v31).length));
          (_64_v31)[_index12] = p2;
          let _65_v32;
          let _nw15 = new _module.C0();
          _nw15.__ctor(_64_v31);
          _65_v32 = _nw15;
          _65_v32 = _65_v32;
          let _66_v33;
          let _nw16 = new _module.C0();
          _nw16.__ctor((_65_v32).f15);
          _66_v33 = _nw16;
          let _67_v34;
          _67_v34 = _dafny.Set.fromElements(p0, p0, p2, p2, (_64_v31)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_64_v31).length))]);
          let _68_v35;
          _68_v35 = _dafny.MultiSet.fromElements(_67_v34, _67_v34, _67_v34);
          let _69_v36;
          _69_v36 = _module.D5.create_DC15(_68_v35);
          let _index13 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_64_v31).length));
          (_64_v31)[_index13] = ((_69_v36).dtor_cf23).IsDisjointFrom(_dafny.MultiSet.fromElements(_67_v34));
          let _70_v37;
          let _nw17 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _70_v37 = _nw17;
          let _index14 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_70_v37).length));
          (_70_v37)[_index14] = _module.__default.fm11(globalState);
          let _71_v38;
          _71_v38 = _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)));
          let _72_v39;
          _72_v39 = new _dafny.CodePoint('v'.codePointAt(0));
          let _index15 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_70_v37).length));
          (_70_v37)[_index15] = ((_71_v38).Intersect(_dafny.Set.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), _72_v39, _72_v39))).Union(_71_v38);
        } else {
          let _73_v40;
          _73_v40 = _module.D5.create_DC17(p2);
          let _pat_let_tv0 = p0;
          _73_v40 = ((true) ? (function (_pat_let0_0) {
            return function (_74_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_75_dt__update_hcf26_h0) {
                  return _module.D5.create_DC17(_75_dt__update_hcf26_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_73_v40)) : (_73_v40));
          let _76_v41;
          let _nw18 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _76_v41 = _nw18;
          let _index16 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_76_v41).length));
          (_76_v41)[_index16] = p1;
          let _index17 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_76_v41).length));
          (_76_v41)[_index17] = ((_dafny.ZERO).minus(p1)).minus(new BigNumber(-577));
          let _77_v42;
          let _nw19 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _77_v42 = _nw19;
          let _78_v43;
          _78_v43 = _dafny.Seq.UnicodeFromString("ctnu");
          let _index18 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_77_v42).length));
          (_77_v42)[_index18] = _78_v43;
          let _index19 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_77_v42).length));
          (_77_v42)[_index19] = _78_v43;
          let _79_v44;
          let _init4 = ((_80_p0) => function (_81_i9) {
            return _80_p0;
          })(p0);
          let _nw20 = Array((new BigNumber(24)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw20.length); _i0_4++) {
            _nw20[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _79_v44 = _nw20;
          let _82_v45;
          let _nw21 = new _module.C0();
          _nw21.__ctor(_79_v44);
          _82_v45 = _nw21;
          let _83_v46;
          _83_v46 = new _dafny.CodePoint('r'.codePointAt(0));
          let _rhs3 = _module.__default.fm8(!(p0) || (p0), _83_v46, _83_v46, globalState);
          let _rhs4 = p1;
          let _lhs2 = globalState;
          let _lhs3 = globalState;
          _lhs2.f5 = _rhs3;
          _lhs3.f0 = _rhs4;
        }
        let _84_v47;
        _84_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _85_v48;
        let _nw22 = Array((new BigNumber(4)).toNumber());
        _nw22[(_dafny.ZERO).toNumber()] = p0;
        _nw22[(_dafny.ONE).toNumber()] = p2;
        _nw22[(new BigNumber(2)).toNumber()] = p2;
        _nw22[(new BigNumber(3)).toNumber()] = (((_84_v47).contains(p1)) ? ((_84_v47).get(p1)) : (p2));
        _85_v48 = _nw22;
        let _86_v49;
        let _nw23 = new _module.C0();
        _nw23.__ctor(_85_v48);
        _86_v49 = _nw23;
        let _87_v50;
        _87_v50 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_84_v47).Merge(_84_v47));
        let _88_v51;
        _88_v51 = _dafny.MultiSet.fromElements(_86_v49, _86_v49);
        _87_v50 = (_87_v50).update((new BigNumber((_88_v51).cardinality())).multipliedBy(p1), (((_87_v50).contains(new BigNumber(70))) ? ((_87_v50).get(new BigNumber(70))) : (_84_v47)));
        let _89_v52;
        let _nw24 = new _module.C0();
        _nw24.__ctor((_86_v49).f15);
        _89_v52 = _nw24;
        _89_v52 = _89_v52;
        let _index20 = _module.__default.safeIndex(new BigNumber(186), new BigNumber(((_89_v52).f15).length));
        ((_89_v52).f15)[_index20] = p2;
        let _90_v53;
        _90_v53 = _dafny.Seq.UnicodeFromString("uuolsr");
        let _91_v54;
        _91_v54 = _dafny.MultiSet.fromElements(_90_v53);
        let _index21 = _module.__default.safeIndex(new BigNumber(186), new BigNumber(((_89_v52).f15).length));
        ((_89_v52).f15)[_index21] = !(_91_v54).equals(_91_v54);
      }
      if (true) {
        r1 = _module.__default.fm4(globalState);
        let _source1 = _module.__default.fm12(globalState);
        if (_source1.is_DC14) {
          let _92___mcc_h0 = (_source1).cf21;
          let _93___mcc_h1 = (_source1).cf22;
          let _94_cf22 = _93___mcc_h1;
          let _95_cf21 = _92___mcc_h0;
          r2 = p2;
          r2 = p2;
          let _96_v55;
          _96_v55 = _module.D1.create_DC1((_dafny.ZERO).minus(p1));
          let _97_v56;
          _97_v56 = _module.D2.create_DC7(p2, (_96_v55).dtor_cf1);
          _97_v56 = _97_v56;
          let _98_v57;
          _98_v57 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _99_v58;
          _99_v58 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(521)));
          let _100_v59;
          _100_v59 = _dafny.Set.fromElements(p1, new BigNumber(-743));
          _98_v57 = (_98_v57).update(true, (_99_v58).IsDisjointFrom(_dafny.MultiSet.fromElements(_100_v59, _100_v59)));
        } else {
          let _101___mcc_h2 = (_source1).cf20;
          let _102_cf20 = _101___mcc_h2;
          let _103_v60;
          _103_v60 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(316),((_dafny.MultiSet.fromElements(false)).update(p0, _module.__default.abs(p1))).update(p0, _module.__default.abs(new BigNumber(884))))).length));
          let _104_v61;
          _104_v61 = new _dafny.CodePoint('d'.codePointAt(0));
          let _105_v62;
          _105_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jyww"),p0);
          let _106_v63;
          _106_v63 = _dafny.Seq.of(!(p0), p2, _module.__default.fm1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_103_v60).length),p0), _dafny.Seq.update(_102_cf20, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_102_cf20).length)), _104_v61), _dafny.Seq.UnicodeFromString("rswrhg"), globalState), (((_105_v62).contains(_dafny.Seq.update(_102_cf20, _module.__default.safeIndex(p1, new BigNumber((_102_cf20).length)), _104_v61))) ? ((_105_v62).get(_dafny.Seq.update(_102_cf20, _module.__default.safeIndex(p1, new BigNumber((_102_cf20).length)), _104_v61))) : (p0)));
          let _107_v64;
          _107_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,_106_v63);
          _107_v64 = _107_v64;
          let _108_v65;
          let _nw25 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _108_v65 = _nw25;
          let _index22 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_108_v65).length));
          (_108_v65)[_index22] = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _109_v66;
          _109_v66 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),((true) ? (p0) : (p2)));
          let _index23 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_108_v65).length));
          (_108_v65)[_index23] = _109_v66;
          let _110_v67;
          let _nw26 = Array((new BigNumber(23)).toNumber()).fill(false);
          _110_v67 = _nw26;
          let _index24 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_110_v67).length));
          (_110_v67)[_index24] = p0;
          let _index25 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_110_v67).length));
          (_110_v67)[_index25] = !(new BigNumber(989)).isEqualTo(p1);
          let _111_v68;
          _111_v68 = _dafny.Map.Empty.slice().updateUnsafe((_110_v67)[_module.__default.safeIndex(new BigNumber(603), new BigNumber((_110_v67).length))],p1);
          let _112_v69;
          _112_v69 = _module.D2.create_DC6(_111_v68);
          let _113_v70;
          _113_v70 = _dafny.Seq.of(_112_v69, _112_v69, _module.D2.create_DC6(_111_v68));
          let _index26 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_110_v67).length));
          (_110_v67)[_index26] = _dafny.Seq.contains(_113_v70, _112_v69);
        }
        let _114_v71;
        let _init5 = ((_115_p0) => function (_116_i10) {
          return _115_p0;
        })(p0);
        let _nw27 = Array((new BigNumber(13)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw27.length); _i0_5++) {
          _nw27[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _114_v71 = _nw27;
        let _117_v72;
        let _nw28 = new _module.C0();
        _nw28.__ctor(_114_v71);
        _117_v72 = _nw28;
        if (p0) {
          r0 = p1;
          let _118_v73;
          _118_v73 = _dafny.Set.fromElements(true, !(true));
          let _119_v74;
          _119_v74 = _module.D6.create_DC18(_118_v73);
          let _rhs5 = (_dafny.ZERO).minus(p1);
          let _rhs6 = (_119_v74).dtor_cf27;
          r1 = _rhs5;
          _118_v73 = _rhs6;
          let _120_v75;
          let _nw29 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _120_v75 = _nw29;
          let _121_v76;
          _121_v76 = _module.D7.create_DC20(_120_v75);
          let _122_v77;
          let _nw30 = Array((new BigNumber(18)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _120_v75;
          _nw30[(_dafny.ONE).toNumber()] = _120_v75;
          _nw30[(new BigNumber(2)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(3)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(4)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(5)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(6)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(7)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(8)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(9)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(10)).toNumber()] = (_121_v76).dtor_cf31;
          _nw30[(new BigNumber(11)).toNumber()] = (_121_v76).dtor_cf31;
          _nw30[(new BigNumber(12)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(13)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(14)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(15)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(16)).toNumber()] = _120_v75;
          _nw30[(new BigNumber(17)).toNumber()] = _120_v75;
          _122_v77 = _nw30;
          let _123_v78;
          _123_v78 = _dafny.Map.Empty.slice().updateUnsafe(p2,_122_v77);
          _123_v78 = (_123_v78).update(p0, _122_v77);
          (globalState).f0 = p1;
          let _124_v79;
          _124_v79 = _dafny.Seq.UnicodeFromString("qcyapl");
          r2 = !(((!_dafny.areEqual(_124_v79, _124_v79)) ? (!(p2)) : (false)));
        } else {
          let _125_v80;
          _125_v80 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _126_v81;
          _126_v81 = _dafny.Seq.of(new BigNumber((_125_v80).length), new BigNumber((_module.__default.fm0(p0, globalState)).length), p1);
          let _127_v82;
          _127_v82 = _module.D3.create_DC11(_126_v81, p2);
          _127_v82 = _127_v82;
          r1 = p1;
          let _128_v83;
          let _init6 = ((_129_p1) => function (_130_i11) {
            return (_130_i11).multipliedBy(_129_p1);
          })(p1);
          let _nw31 = Array((new BigNumber(6)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw31.length); _i0_6++) {
            _nw31[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _128_v83 = _nw31;
          let _index27 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_128_v83).length));
          (_128_v83)[_index27] = (_dafny.ZERO).minus(p1);
          let _index28 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_128_v83).length));
          (_128_v83)[_index28] = ((((_14_v0).contains(true)) ? ((_14_v0).get(true)) : (p1))).minus(new BigNumber((_dafny.Seq.UnicodeFromString("eskgngpo")).length));
          let _131_v84;
          let _nw32 = new _module.C0();
          _nw32.__ctor((_117_v72).f15);
          _131_v84 = _nw32;
          let _index29 = _module.__default.safeIndex(new BigNumber(112), new BigNumber(((_117_v72).f15).length));
          ((_117_v72).f15)[_index29] = p0;
          let _132_v85;
          _132_v85 = _dafny.Set.fromElements((_128_v83)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_128_v83).length))], p1);
          let _133_v86;
          _133_v86 = _dafny.Seq.of(p0);
          let _134_v87;
          _134_v87 = new _dafny.CodePoint('f'.codePointAt(0));
          let _135_v88;
          _135_v88 = _dafny.Map.Empty.slice().updateUnsafe((_133_v86)[_module.__default.safeIndex(new BigNumber((_126_v81).length), new BigNumber((_133_v86).length))],_134_v87);
          let _index30 = _module.__default.safeIndex(new BigNumber(112), new BigNumber(((_117_v72).f15).length));
          ((_117_v72).f15)[_index30] = ((_dafny.Set.fromElements(p1, new BigNumber((_dafny.Seq.of(_135_v88)).length))).Difference(_132_v85)).IsProperSubsetOf(_132_v85);
        }
        let _136_v89;
        _136_v89 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),(_117_v72).f15);
        _136_v89 = (_136_v89).update(p2, _114_v71);
      } else {
        let _137_v90;
        _137_v90 = _dafny.Seq.of((p1).multipliedBy(p1), p1);
        let _138_v91;
        let _nw33 = Array((new BigNumber(23)).toNumber()).fill(false);
        _138_v91 = _nw33;
        let _index31 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v91).length));
        (_138_v91)[_index31] = p0;
        let _139_v92;
        _139_v92 = _dafny.Seq.UnicodeFromString("j");
        let _140_v93;
        _140_v93 = new _dafny.CodePoint('p'.codePointAt(0));
        let _index32 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v91).length));
        let _rhs7 = _137_v90;
        let _rhs8 = ((p1).plus(new BigNumber((_dafny.Seq.update(_139_v92, _module.__default.safeIndex(p1, new BigNumber((_139_v92).length)), _140_v93)).length))).isLessThan((_137_v90)[_module.__default.safeIndex(p1, new BigNumber((_137_v90).length))]);
        let _lhs4 = _138_v91;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v91).length));
        _137_v90 = _rhs7;
        _lhs4[_lhs5] = _rhs8;
        let _141_v94;
        let _nw34 = new _module.C0();
        _nw34.__ctor(_138_v91);
        _141_v94 = _nw34;
        let _142_v95;
        _142_v95 = _dafny.Seq.of(_141_v94);
        let _143_v96;
        _143_v96 = _dafny.MultiSet.fromElements(_141_v94, _141_v94, (_142_v95)[_module.__default.safeIndex(new BigNumber(-526), new BigNumber((_142_v95).length))], _141_v94);
        _143_v96 = _143_v96;
        let _144_v97;
        let _nw35 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
        _144_v97 = _nw35;
        let _145_v98;
        _145_v98 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),p1);
        let _146_v99;
        _146_v99 = _dafny.MultiSet.fromElements(_145_v98);
        let _147_v100;
        _147_v100 = _146_v99;
        let _index33 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_144_v97).length));
        (_144_v97)[_index33] = (_147_v100);
        let _index34 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_144_v97).length));
        (_144_v97)[_index34] = _146_v99;
        let _148_v101;
        _148_v101 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _149_v102;
        _149_v102 = _dafny.Map.Empty.slice().updateUnsafe(_148_v101,(_138_v91)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_138_v91).length))]);
        _149_v102 = (_149_v102).update(_148_v101, (new BigNumber((_139_v92).length)).isEqualTo(p1));
        r0 = p1;
      }
      let _150_v103;
      _150_v103 = _dafny.MultiSet.fromElements(p1);
      let _151_v105;
      _151_v105 = new _dafny.CodePoint('a'.codePointAt(0));
      let _152_v106;
      _152_v106 = _dafny.Seq.of(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_151_v105,p2)).Keys.Elements) {
          let _153_v104 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_151_v105,p2)).contains(_153_v104)) {
            _coll3.push([_153_v104,new BigNumber((_dafny.Seq.UnicodeFromString("spyhcfqae")).length)]);
          }
        }
        return _coll3;
      }());
      let _154_v107;
      _154_v107 = _dafny.Map.Empty.slice().updateUnsafe(_151_v105,new BigNumber(844));
      let _155_v108;
      _155_v108 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus(p1), new BigNumber((_150_v103).cardinality()), new BigNumber((_dafny.Seq.update(_152_v106, _module.__default.safeIndex(_module.__default.fm4(globalState), new BigNumber((_152_v106).length)), _154_v107)).length), p1);
      let _156_v109;
      _156_v109 = _dafny.Seq.of(_module.__default.fm13(new BigNumber(503), globalState));
      let _157_v110;
      _157_v110 = _dafny.Set.fromElements(p0);
      let _158_v111;
      _158_v111 = _dafny.MultiSet.fromElements(_157_v110, _dafny.Set.fromElements(true, p2), _157_v110, _157_v110);
      let _159_v112;
      _159_v112 = _module.D5.create_DC15(_158_v111);
      let _160_v113;
      _160_v113 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_dafny.Set.fromElements(p1), _155_v108, (_156_v109)[_module.__default.safeIndex(p1, new BigNumber((_156_v109).length))], _155_v108, _155_v108),_159_v112);
      _160_v113 = (_160_v113).update(_156_v109, _159_v112);
      let _161_v114;
      let _init7 = ((_162_p2) => function (_163_i13) {
        return _162_p2;
      })(p2);
      let _nw36 = Array((new BigNumber(22)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw36.length); _i0_7++) {
        _nw36[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _161_v114 = _nw36;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_161_v114).length))) {
        let _164_i12 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_164_i12)) && ((_164_i12).isLessThan(new BigNumber((_161_v114).length))))) {
          (_161_v114)[(_164_i12)] = ((p0) ? (!(p2)) : ((new BigNumber((_dafny.Seq.UnicodeFromString("ldv")).length)).isLessThanOrEqualTo(p1)));
        }
      }
      let _165_i14;
      _165_i14 = _dafny.ZERO;
      L0: {
        while ((_14_v0).IsProperSubsetOf(((_14_v0).update(p2, _module.__default.abs((_dafny.ZERO).minus(p1)))).update(false, _module.__default.abs(p1)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_165_i14)) {
              break L0;
            }
            _165_i14 = (_165_i14).plus(_dafny.ONE);
            r2 = ((p2) ? (p0) : (!(((p0) ? (p2) : (p2)))));
            r2 = ((p2) ? (false) : (p2));
            let _index35 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            (_161_v114)[_index35] = !(p0);
            let _index36 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            (_161_v114)[_index36] = false;
            let _166_v115;
            _166_v115 = _module.D5.create_DC16((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p0)).length)).multipliedBy(p1), p1);
            let _167_v116;
            _167_v116 = _dafny.Seq.UnicodeFromString("gsrufvirc");
            let _index37 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            let _index38 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            let _rhs9 = function (_pat_let2_0) {
              return function (_168_dt__update__tmp_h1) {
                return function (_pat_let3_0) {
                  return function (_169_dt__update_hcf24_h0) {
                    return _module.D5.create_DC16(_169_dt__update_hcf24_h0, (_168_dt__update__tmp_h1).dtor_cf25);
                  }(_pat_let3_0);
                }(new BigNumber(590));
              }(_pat_let2_0);
            }(_module.D5.create_DC16(p1, p1));
            let _rhs10 = new BigNumber((_167_v116).length);
            let _rhs11 = !(p2);
            let _rhs12 = (_161_v114)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length))];
            let _lhs6 = globalState;
            let _lhs7 = _161_v114;
            let _lhs8 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            let _lhs9 = _161_v114;
            let _lhs10 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_161_v114).length));
            _166_v115 = _rhs9;
            _lhs6.f0 = _rhs10;
            _lhs7[_lhs8] = _rhs11;
            _lhs9[_lhs10] = _rhs12;
          }
        }
      }
      let _170_v117;
      let _nw37 = new _module.C0();
      _nw37.__ctor(_161_v114);
      _170_v117 = _nw37;
      r0 = p1;
      r1 = (_dafny.ZERO).minus((new BigNumber(937)).plus(_module.__default.fm4(globalState)));
      r2 = !(p0);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _171_v0;
      _171_v0 = true;
      let _172_v1;
      _172_v1 = _dafny.Seq.UnicodeFromString("fyhqxeqer");
      let _173_v2;
      let _nw38 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
      _173_v2 = _nw38;
      let _174_globalState;
      let _nw39 = new _module.GlobalState();
      _nw39.__ctor(new BigNumber(-982), _dafny.Seq.of(_171_v0), false, _dafny.Seq.Concat(_172_v1, _172_v1), new BigNumber(745), _172_v1, new BigNumber(292), new _dafny.CodePoint('x'.codePointAt(0)), new BigNumber(420), _dafny.Seq.Concat(_172_v1, _172_v1), new BigNumber(835), false, new BigNumber(-89), new BigNumber(488), _173_v2);
      _174_globalState = _nw39;
      let _175_v3;
      _175_v3 = _dafny.Seq.of(_171_v0, _171_v0);
      _175_v3 = _dafny.Seq.Concat(_175_v3, _module.__default.fm0(_171_v0, _174_globalState));
      let _176_v4;
      _176_v4 = _dafny.Seq.of(_175_v3);
      let _177_v5;
      _177_v5 = _dafny.Map.Empty.slice().updateUnsafe((true),_171_v0);
      let _178_v6;
      _178_v6 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual((_176_v4)[_module.__default.safeIndex(new BigNumber((_177_v5).length), new BigNumber((_176_v4).length))], _175_v3),_172_v1);
      _178_v6 = (_178_v6).update(_module.__default.fm1(_module.__default.fm2(_171_v0, _171_v0, _174_globalState), _dafny.Seq.UnicodeFromString("xqfstui"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), function (_179_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), _174_globalState), _dafny.Seq.Concat(_172_v1, _172_v1));
      let _180_i1;
      _180_i1 = _dafny.ZERO;
      L1: {
        while (((_171_v0) ? (_171_v0) : (_171_v0))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_180_i1)) {
              break L1;
            }
            _180_i1 = (_180_i1).plus(_dafny.ONE);
            let _181_v7;
            _181_v7 = new BigNumber(795);
            let _182_v8;
            let _183_v9;
            let _184_v10;
            let _out0;
            let _out1;
            let _out2;
            let _outcollector0 = _module.__default.m0(!(_171_v0), (_dafny.ZERO).minus(_181_v7), _171_v0, _174_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _out2 = _outcollector0[2];
            _182_v8 = _out0;
            _183_v9 = _out1;
            _184_v10 = _out2;
            let _185_v11;
            let _nw40 = Array((new BigNumber(29)).toNumber()).fill([]);
            _185_v11 = _nw40;
            let _186_v12;
            let _nw41 = Array((new BigNumber(20)).toNumber()).fill(false);
            _186_v12 = _nw41;
            let _nw42 = Array((new BigNumber(27)).toNumber());
            _nw42[(_dafny.ZERO).toNumber()] = _186_v12;
            _nw42[(_dafny.ONE).toNumber()] = _186_v12;
            _nw42[(new BigNumber(2)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(3)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(4)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(5)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(6)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(7)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(8)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(9)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(10)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(11)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(12)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(13)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(14)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(15)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(16)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(17)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(18)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(19)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(20)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(21)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(22)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(23)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(24)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(25)).toNumber()] = _186_v12;
            _nw42[(new BigNumber(26)).toNumber()] = _186_v12;
            _185_v11 = _nw42;
            let _187_v13;
            let _nw43 = Array((new BigNumber(15)).toNumber()).fill(false);
            _187_v13 = _nw43;
            let _188_v14;
            _188_v14 = _171_v0;
            let _index39 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_187_v13).length));
            (_187_v13)[_index39] = _188_v14;
            let _index40 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_187_v13).length));
            (_187_v13)[_index40] = _188_v14;
            let _rhs13 = _181_v7;
            let _rhs14 = _184_v10;
            _181_v7 = _rhs13;
            _184_v10 = _rhs14;
          }
        }
      }
      let _189_v15;
      _189_v15 = new BigNumber(261);
      (_174_globalState).f0 = (_189_v15).multipliedBy(_189_v15);
      let _190_v16;
      _190_v16 = !(_171_v0);
      _190_v16 = _190_v16;
      (_174_globalState).f0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(((_171_v0) ? (_189_v15) : (_189_v15)))), (_189_v15).minus(_189_v15));
      let _191_i2;
      _191_i2 = _dafny.ZERO;
      L2: {
        while ((_171_v0)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_191_i2)) {
              break L2;
            }
            _191_i2 = (_191_i2).plus(_dafny.ONE);
            let _192_v17;
            let _init8 = ((_193_v0) => function (_194_i3) {
              return _193_v0;
            })(_171_v0);
            let _nw44 = Array((new BigNumber(28)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw44.length); _i0_8++) {
              _nw44[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _192_v17 = _nw44;
            let _index41 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_192_v17).length));
            (_192_v17)[_index41] = _171_v0;
            let _195_v18;
            let _nw45 = Array((_dafny.ONE).toNumber());
            _nw45[(_dafny.ZERO).toNumber()] = _190_v16;
            _195_v18 = _nw45;
            let _196_v19;
            _196_v19 = _dafny.Map.Empty.slice().updateUnsafe(_195_v18,_189_v15);
            let _index42 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_192_v17).length));
            (_192_v17)[_index42] = (_196_v19).equals(_196_v19);
            let _197_v20;
            _197_v20 = new _dafny.CodePoint('b'.codePointAt(0));
            let _rhs15 = _189_v15;
            let _rhs16 = _module.__default.safeModuloInt(_189_v15, _189_v15);
            let _rhs17 = !((_192_v17)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_192_v17).length))]);
            let _rhs18 = _172_v1;
            let _rhs19 = new _dafny.CodePoint('c'.codePointAt(0));
            let _lhs11 = _174_globalState;
            let _lhs12 = _174_globalState;
            _lhs11.f0 = _rhs15;
            _189_v15 = _rhs16;
            _171_v0 = _rhs17;
            _lhs12.f5 = _rhs18;
            _197_v20 = _rhs19;
            let _198_v21;
            _198_v21 = _dafny.Set.fromElements(_dafny.Seq.of((_175_v3)[_module.__default.safeIndex(new BigNumber((_module.__default.fm3(_189_v15, _189_v15, _189_v15, _174_globalState)).cardinality()), new BigNumber((_175_v3).length))]), _175_v3);
            let _199_v22;
            _199_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_198_v21).length), new BigNumber(426)),_189_v15);
            _199_v22 = ((_199_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(_189_v15,_189_v15))).Merge(_199_v22);
            let _200_v23;
            _200_v23 = _dafny.MultiSet.fromElements(_189_v15);
            let _201_v24;
            _201_v24 = _dafny.Map.Empty.slice().updateUnsafe((_200_v23).update(_189_v15, _module.__default.abs(_189_v15)),_189_v15);
            _175_v3 = _dafny.Seq.Concat(_dafny.Seq.update(_175_v3, _module.__default.safeIndex(new BigNumber((_201_v24).length), new BigNumber((_175_v3).length)), _171_v0), _175_v3);
          }
        }
      }
      let _202_v25;
      let _nw46 = Array((new BigNumber(28)).toNumber()).fill(false);
      _202_v25 = _nw46;
      let _203_v26;
      _203_v26 = _dafny.Set.fromElements(_202_v25, _202_v25);
      let _204_v27;
      _204_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_174_globalState),_203_v26);
      let _205_v28;
      let _nw47 = Array((new BigNumber(25)).toNumber());
      _nw47[(_dafny.ZERO).toNumber()] = _203_v26;
      _nw47[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_202_v25, _202_v25);
      _nw47[(new BigNumber(2)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_202_v25);
      _nw47[(new BigNumber(4)).toNumber()] = (_203_v26).Union(_203_v26);
      _nw47[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_202_v25);
      _nw47[(new BigNumber(6)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(7)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_202_v25, _202_v25, _202_v25);
      _nw47[(new BigNumber(9)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(10)).toNumber()] = (_203_v26).Intersect(_203_v26);
      _nw47[(new BigNumber(11)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(12)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(13)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_202_v25);
      _nw47[(new BigNumber(15)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(16)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(17)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(18)).toNumber()] = (_dafny.Set.fromElements(_202_v25, _202_v25)).Union(_203_v26);
      _nw47[(new BigNumber(19)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(20)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(21)).toNumber()] = ((((_204_v27).contains(_189_v15)) ? ((_204_v27).get(_189_v15)) : (_203_v26))).Union(_203_v26);
      _nw47[(new BigNumber(22)).toNumber()] = _203_v26;
      _nw47[(new BigNumber(23)).toNumber()] = (((_175_v3)[_module.__default.safeIndex(_189_v15, new BigNumber((_175_v3).length))]) ? (_203_v26) : (_203_v26));
      _nw47[(new BigNumber(24)).toNumber()] = _203_v26;
      _205_v28 = _nw47;
      let _index43 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_205_v28).length));
      (_205_v28)[_index43] = _203_v26;
      let _index44 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_205_v28).length));
      (_205_v28)[_index44] = _203_v26;
      let _source2 = _190_v16;
      let _206___mcc_h0 = _source2;
      let _207_cf0 = _206___mcc_h0;
      _178_v6 = (_178_v6).update((!(_171_v0)) === (_171_v0), _dafny.Seq.Concat(_172_v1, _172_v1));
      let _208_v29;
      _208_v29 = _dafny.Set.fromElements(_207_cf0);
      let _209_v30;
      let _nw48 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _209_v30 = _nw48;
      let _index45 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length));
      (_209_v30)[_index45] = _189_v15;
      let _210_v31;
      _210_v31 = _dafny.Set.fromElements((_dafny.ZERO).minus(_189_v15));
      let _211_v32;
      _211_v32 = _dafny.MultiSet.fromElements(_172_v1);
      let _index46 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length));
      let _rhs20 = _module.__default.fm5(_189_v15, (new BigNumber((_210_v31).length)).isLessThan(_189_v15), _dafny.Seq.Concat(_172_v1, _172_v1), (_211_v32).Union(_211_v32), _174_globalState);
      let _rhs21 = _171_v0;
      let _rhs22 = (_189_v15).plus(_module.__default.fm4(_174_globalState));
      let _lhs13 = _209_v30;
      let _lhs14 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length));
      _208_v29 = _rhs20;
      _207_cf0 = _rhs21;
      _lhs13[_lhs14] = _rhs22;
      let _212_v33;
      _212_v33 = _dafny.Map.Empty.slice().updateUnsafe((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))],_171_v0);
      let _213_v34;
      _213_v34 = new _dafny.CodePoint('f'.codePointAt(0));
      let _index47 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length));
      let _rhs23 = new BigNumber(373);
      let _rhs24 = ((_dafny.Set.fromElements(_207_cf0)).Intersect(_208_v29)).IsProperSubsetOf(_208_v29);
      let _rhs25 = _module.__default.fm1(_212_v33, _dafny.Seq.UnicodeFromString("mtng"), _172_v1, _174_globalState);
      let _rhs26 = (_dafny.ZERO).minus(((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))]).plus(((_207_cf0) ? (new BigNumber((_dafny.Seq.update(_172_v1, _module.__default.safeIndex((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))], new BigNumber((_172_v1).length)), _213_v34)).length)) : ((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))]))));
      let _rhs27 = (new BigNumber((_172_v1).length)).isLessThanOrEqualTo(_189_v15);
      let _lhs15 = _174_globalState;
      let _lhs16 = _209_v30;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length));
      _lhs15.f0 = _rhs23;
      _207_cf0 = _rhs24;
      _171_v0 = _rhs25;
      _lhs16[_lhs17] = _rhs26;
      _207_cf0 = _rhs27;
      if (!(false)) {
        let _214_v35;
        let _nw49 = new _module.C0();
        _nw49.__ctor(_202_v25);
        _214_v35 = _nw49;
        _171_v0 = ((new BigNumber((_175_v3).length)).multipliedBy((_dafny.ZERO).minus((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))]))).isLessThanOrEqualTo(_module.__default.fm4(_174_globalState));
        let _215_v36;
        let _216_v37;
        let _217_v38;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0(_207_cf0, (_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))], _207_cf0, _174_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _215_v36 = _out3;
        _216_v37 = _out4;
        _217_v38 = _out5;
        _177_v5 = (_177_v5).Merge((_177_v5).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_207_cf0),_171_v0)));
        let _218_v39;
        let _219_v40;
        let _220_v41;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = _module.__default.m0((_215_v36).isLessThan(_module.__default.fm4(_174_globalState)), (_dafny.ZERO).minus((_216_v37).plus((_module.D1.create_DC1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), ((_221_v34) => function (_222_i4) {
  return _221_v34;
})(_213_v34))).length))).dtor_cf1)), _207_cf0, _174_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _218_v39 = _out6;
        _219_v40 = _out7;
        _220_v41 = _out8;
      } else {
        let _223_v42;
        _223_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_212_v33).length),(_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))]);
        let _224_v43;
        _224_v43 = _dafny.Map.Empty.slice().updateUnsafe(_207_cf0,new BigNumber((_223_v42).length));
        let _225_v44;
        _225_v44 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,_224_v43);
        let _226_v45;
        _226_v45 = _dafny.Seq.of((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))], _189_v15, (_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))], new BigNumber(((((_225_v44).contains(_189_v15)) ? ((_225_v44).get(_189_v15)) : ((_224_v43).update(false, _189_v15)))).length), new BigNumber((_210_v31).length));
        let _227_v46;
        let _228_v47;
        let _229_v48;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector3 = _module.__default.m0(_171_v0, ((_207_cf0) ? (new BigNumber((_226_v45).length)) : ((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))])), !(!(new BigNumber((_177_v5).length)).isEqualTo((_209_v30)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_209_v30).length))])), _174_globalState);
        _out9 = _outcollector3[0];
        _out10 = _outcollector3[1];
        _out11 = _outcollector3[2];
        _227_v46 = _out9;
        _228_v47 = _out10;
        _229_v48 = _out11;
        _178_v6 = (_178_v6).update(false, _dafny.Seq.UnicodeFromString("yhphth"));
        let _230_v49;
        let _231_v50;
        let _232_v51;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m0((true) && (_207_cf0), new BigNumber(825), _module.__default.fm1(_module.__default.fm2(_229_v48, _171_v0, _174_globalState), _172_v1, _172_v1, _174_globalState), _174_globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _230_v49 = _out12;
        _231_v50 = _out13;
        _232_v51 = _out14;
        (_174_globalState).f5 = _172_v1;
        (_174_globalState).f5 = _dafny.Seq.Concat(_172_v1, _dafny.Seq.Concat(_172_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), ((_233_v34) => function (_234_i5) {
          return _233_v34;
        })(_213_v34))));
      }
      let _235_v52;
      let _nw50 = new _module.C0();
      _nw50.__ctor(_202_v25);
      _235_v52 = _nw50;
      let _236_v53;
      _236_v53 = _dafny.Set.fromElements(_171_v0);
      let _237_v54;
      _237_v54 = _dafny.MultiSet.fromElements(_172_v1, _172_v1);
      let _238_v55;
      _238_v55 = _module.D1.create_DC4((_module.__default.fm5(new BigNumber(867), !(false), _172_v1, _237_v54, _174_globalState)).IsSubsetOf(_236_v53), _171_v0, _202_v25, _171_v0);
      let _source3 = _238_v55;
      if (_source3.is_DC2) {
        let _239___mcc_h1 = (_source3).cf2;
        let _240___mcc_h2 = (_source3).cf3;
        let _241___mcc_h3 = (_source3).cf4;
        let _242_cf4 = _241___mcc_h3;
        let _243_cf3 = _240___mcc_h2;
        let _244_cf2 = _239___mcc_h1;
        let _245_v56;
        _245_v56 = _module.D1.create_DC2(_244_cf2, _171_v0, _171_v0);
        let _246_v57;
        _246_v57 = _dafny.Map.Empty.slice().updateUnsafe(_245_v56,_243_cf3);
        let _247_v58;
        _247_v58 = _dafny.MultiSet.fromElements(_246_v57, _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(_244_cf2, false, _190_v16),_243_cf3));
        let _248_v59;
        _248_v59 = _dafny.Seq.of(new BigNumber(357), (((_247_v58).contains(_module.__default.fm6(_dafny.Seq.UnicodeFromString("yhshk"), new BigNumber((_175_v3).length), _174_globalState))) ? ((_247_v58).get(_module.__default.fm6(_dafny.Seq.UnicodeFromString("yhshk"), new BigNumber((_175_v3).length), _174_globalState))) : (_244_cf2)), _244_cf2, _244_cf2, _189_v15);
        let _249_v60;
        _249_v60 = new _dafny.CodePoint('d'.codePointAt(0));
        let _250_v61;
        let _nw51 = Array((new BigNumber(28)).toNumber());
        _250_v61 = _nw51;
        let _251_v62;
        _251_v62 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,_250_v61);
        let _252_v63;
        let _nw52 = Array((new BigNumber(19)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = _248_v59;
        _nw52[(_dafny.ONE).toNumber()] = _248_v59;
        _nw52[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_189_v15), _248_v59);
        _nw52[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), ((_253_v15) => function (_254_i6) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_253_v15,true)).length);
        })(_189_v15)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-511)), ((_255_v15) => function (_256_i7) {
          return _255_v15;
        })(_189_v15)));
        _nw52[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), ((_257_v15) => function (_258_i8) {
          return _257_v15;
        })(_189_v15));
        _nw52[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_248_v59, _248_v59);
        _nw52[(new BigNumber(6)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(7)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_248_v59, _248_v59);
        _nw52[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), ((_259_v15) => function (_260_i9) {
          return (_dafny.ZERO).minus(_259_v15);
        })(_189_v15));
        _nw52[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-277)), ((_261_cf2) => function (_262_i10) {
          return _261_cf2;
        })(_244_cf2));
        _nw52[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_248_v59, _248_v59);
        _nw52[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_248_v59, _248_v59);
        _nw52[(new BigNumber(13)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(((true) ? (_248_v59) : (_module.__default.fm7(_243_cf3, _249_v60, _244_cf2, _174_globalState))), _module.__default.safeIndex(new BigNumber(((_251_v62).update(_189_v15, _250_v61)).length), new BigNumber((((true) ? (_248_v59) : (_module.__default.fm7(_243_cf3, _249_v60, _244_cf2, _174_globalState)))).length)), _244_cf2);
        _nw52[(new BigNumber(15)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(16)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(17)).toNumber()] = _248_v59;
        _nw52[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), ((_263_v3) => function (_264_i11) {
          return new BigNumber((_263_v3).length);
        })(_175_v3));
        _252_v63 = _nw52;
        let _index48 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_252_v63).length));
        (_252_v63)[_index48] = _248_v59;
        let _index49 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_252_v63).length));
        (_252_v63)[_index49] = _dafny.Seq.update(_dafny.Seq.Concat(_248_v59, _248_v59), _module.__default.safeIndex((_244_cf2).multipliedBy(new BigNumber(25)), new BigNumber((_dafny.Seq.Concat(_248_v59, _248_v59)).length)), new BigNumber((_172_v1).length));
        let _rhs28 = !(!((_243_cf3)));
        let _rhs29 = _235_v52;
        let _rhs30 = _171_v0;
        let _rhs31 = _244_cf2;
        let _rhs32 = _189_v15;
        _243_cf3 = _rhs28;
        _235_v52 = _rhs29;
        _171_v0 = _rhs30;
        _189_v15 = _rhs31;
        _189_v15 = _rhs32;
        let _265_v64;
        let _266_v65;
        let _267_v66;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector5 = _module.__default.m0((_175_v3)[_module.__default.safeIndex((_dafny.ZERO).minus(_244_cf2), new BigNumber((_175_v3).length))], _189_v15, _171_v0, _174_globalState);
        _out15 = _outcollector5[0];
        _out16 = _outcollector5[1];
        _out17 = _outcollector5[2];
        _265_v64 = _out15;
        _266_v65 = _out16;
        _267_v66 = _out17;
        let _268_v67;
        _268_v67 = _dafny.Map.Empty.slice().updateUnsafe(_267_v66,_175_v3);
        _189_v15 = (new BigNumber(((_268_v67).Merge(_268_v67)).length)).multipliedBy(new BigNumber(214));
      } else if (_source3.is_DC3) {
        (_174_globalState).f0 = (_dafny.ZERO).minus(_189_v15);
        let _index50 = _module.__default.safeIndex(new BigNumber(441), new BigNumber(((_235_v52).f15).length));
        ((_235_v52).f15)[_index50] = _171_v0;
        let _index51 = _module.__default.safeIndex(new BigNumber(441), new BigNumber(((_235_v52).f15).length));
        ((_235_v52).f15)[_index51] = _171_v0;
        let _index52 = _module.__default.safeIndex(new BigNumber(441), new BigNumber(((_235_v52).f15).length));
        ((_235_v52).f15)[_index52] = _171_v0;
        let _269_v69;
        let _init9 = ((_270_v15, _271_v52) => function (_272_i12) {
          return _module.__default.safeModuloInt(_272_i12, new BigNumber((function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.Seq.of(_270_v15)).Elements) {
              let _273_v68 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(_270_v15), _273_v68)) {
                _coll4.push([_module.__default.safeDivisionInt(_273_v68, _270_v15),((_271_v52).f15)[_module.__default.safeIndex(new BigNumber(441), new BigNumber(((_271_v52).f15).length))]]);
              }
            }
            return _coll4;
          }()).length));
        })(_189_v15, _235_v52);
        let _nw53 = Array((new BigNumber(11)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw53.length); _i0_9++) {
          _nw53[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _269_v69 = _nw53;
        _269_v69 = _269_v69;
      } else if (_source3.is_DC4) {
        let _274___mcc_h4 = (_source3).cf5;
        let _275___mcc_h5 = (_source3).cf6;
        let _276___mcc_h6 = (_source3).cf7;
        let _277___mcc_h7 = (_source3).cf8;
        let _278_cf8 = _277___mcc_h7;
        let _279_cf7 = _276___mcc_h6;
        let _280_cf6 = _275___mcc_h5;
        let _281_cf5 = _274___mcc_h4;
        let _282_v70;
        let _nw54 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _282_v70 = _nw54;
        let _index53 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_282_v70).length));
        (_282_v70)[_index53] = _dafny.Seq.UnicodeFromString("cccpixh");
        let _index54 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_282_v70).length));
        (_282_v70)[_index54] = _172_v1;
        if (_278_cf8) {
          (_174_globalState).f0 = _189_v15;
          let _283_v71;
          let _nw55 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _283_v71 = _nw55;
          let _index55 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_283_v71).length));
          (_283_v71)[_index55] = _module.__default.fm4(_174_globalState);
          let _index56 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_283_v71).length));
          (_283_v71)[_index56] = _module.__default.fm4(_174_globalState);
          let _284_v72;
          _284_v72 = _dafny.MultiSet.fromElements(true, _281_cf5);
          let _rhs33 = !(!(((_283_v71)[_module.__default.safeIndex(new BigNumber(181), new BigNumber((_283_v71).length))]).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(550)), ((_285_v15) => function (_286_i13) {
            return _285_v15;
          })(_189_v15))).length))) || ((_238_v55).dtor_cf5));
          let _rhs34 = (_284_v72).IsProperSubsetOf(_284_v72);
          let _rhs35 = (_282_v70)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_282_v70).length))];
          _171_v0 = _rhs33;
          _280_cf6 = _rhs34;
          _172_v1 = _rhs35;
          let _287_v73;
          _287_v73 = _dafny.Map.Empty.slice().updateUnsafe((_283_v71)[_module.__default.safeIndex(new BigNumber(181), new BigNumber((_283_v71).length))],_280_cf6);
          let _288_v74;
          _288_v74 = new _dafny.CodePoint('b'.codePointAt(0));
          _278_cf8 = _module.__default.fm1(_287_v73, _172_v1, _module.__default.fm8(_278_cf8, _288_v74, _module.__default.fm9(_281_cf5, _174_globalState), _174_globalState), _174_globalState);
          let _289_v75;
          let _nw56 = Array((new BigNumber(15)).toNumber());
          _289_v75 = _nw56;
          let _290_v76;
          _290_v76 = _dafny.Seq.of(_289_v75, _289_v75, _289_v75, _289_v75, _289_v75);
          _290_v76 = _290_v76;
        } else {
          (_174_globalState).f0 = _189_v15;
          (_174_globalState).f0 = (_189_v15).plus(_189_v15);
          let _291_v77;
          let _292_v78;
          let _293_v79;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector6 = _module.__default.m0(!(_189_v15).isEqualTo(_189_v15), (_dafny.ZERO).minus(_189_v15), _278_cf8, _174_globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _291_v77 = _out18;
          _292_v78 = _out19;
          _293_v79 = _out20;
          let _294_v80;
          let _nw57 = new _module.C0();
          _nw57.__ctor((_235_v52).f15);
          _294_v80 = _nw57;
          let _index57 = _module.__default.safeIndex(new BigNumber(295), new BigNumber(((_235_v52).f15).length));
          ((_235_v52).f15)[_index57] = _171_v0;
          let _index58 = _module.__default.safeIndex(new BigNumber(295), new BigNumber(((_235_v52).f15).length));
          ((_235_v52).f15)[_index58] = _281_cf5;
        }
        let _295_v81;
        _295_v81 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,_278_cf8);
        _295_v81 = (_295_v81).update(_189_v15, _171_v0);
        (_174_globalState).f5 = _172_v1;
      } else if (_source3.is_DC1) {
        let _296___mcc_h8 = (_source3).cf1;
        let _297_cf1 = _296___mcc_h8;
        (_174_globalState).f0 = _module.__default.safeDivisionInt(_297_cf1, _297_cf1);
        let _298_v82;
        _298_v82 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,_171_v0);
        let _299_v83;
        _299_v83 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,new BigNumber((_298_v82).length));
        let _300_v84;
        _300_v84 = _module.D1.create_DC1(_297_cf1);
        (_174_globalState).f0 = ((true) ? (_297_cf1) : (new BigNumber(((_299_v83).update((_dafny.ZERO).minus((_300_v84).dtor_cf1), _297_cf1)).length)));
        _171_v0 = _171_v0;
        (_174_globalState).f0 = _189_v15;
      } else {
        let _301___mcc_h9 = (_source3).cf9;
        let _302_cf9 = _301___mcc_h9;
        let _303_v85;
        _303_v85 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,_189_v15);
        (_174_globalState).f0 = (_dafny.ZERO).minus((_module.__default.safeModuloInt(_189_v15, _189_v15)).minus((((_303_v85).contains(_189_v15)) ? ((_303_v85).get(_189_v15)) : (_189_v15))));
        let _304_v86;
        let _nw58 = new _module.C0();
        _nw58.__ctor(_202_v25);
        _304_v86 = _nw58;
        let _305_v87;
        _305_v87 = _dafny.Map.Empty.slice().updateUnsafe((((_177_v5).contains(_171_v0)) ? ((_177_v5).get(_171_v0)) : (_171_v0)),_189_v15);
        let _306_v88;
        _306_v88 = _module.D2.create_DC6(_305_v87);
        _305_v87 = (((_306_v88).dtor_cf10).Merge(_dafny.Map.Empty.slice().updateUnsafe(_171_v0,_189_v15))).Merge(_305_v87);
        _171_v0 = _171_v0;
      }
      let _307_v89;
      _307_v89 = _dafny.Map.Empty.slice().updateUnsafe(_189_v15,false);
      let _308_v90;
      _308_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_307_v89, _172_v1, _dafny.Seq.UnicodeFromString("vvidmhm"), _174_globalState),_module.__default.fm7(!(_171_v0), new _dafny.CodePoint('p'.codePointAt(0)), _189_v15, _174_globalState));
      let _309_v91;
      _309_v91 = _dafny.Seq.of(new BigNumber((_236_v53).length), _module.__default.fm4(_174_globalState), _189_v15);
      _308_v90 = (_308_v90).update(_171_v0, _309_v91);
      let _310_v92;
      _310_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xndmxbthy"),_172_v1);
      let _311_v93;
      _311_v93 = _dafny.Map.Empty.slice().updateUnsafe(_310_v92,_171_v0);
      _311_v93 = (_311_v93).update(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_237_v54).Elements) {
          let _312_v94 = _compr_5;
          if ((_237_v54).contains(_312_v94)) {
            _coll5.push([_312_v94,_172_v1]);
          }
        }
        return _coll5;
      }(), ((_171_v0) ? (_171_v0) : (_171_v0)));
      let _313_v95;
      _313_v95 = new _dafny.CodePoint('p'.codePointAt(0));
      let _314_v96;
      _314_v96 = _dafny.MultiSet.fromElements(_171_v0, _171_v0, _171_v0);
      let _315_v97;
      _315_v97 = _dafny.Seq.of(_314_v96);
      let _316_v98;
      _316_v98 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_313_v95, _171_v0, new BigNumber((_315_v97).length), _174_globalState),(_235_v52).f15);
      let _317_v99;
      _317_v99 = _dafny.Map.Empty.slice().updateUnsafe(_171_v0,_189_v15);
      let _318_v100;
      _318_v100 = _module.D2.create_DC6(_317_v99);
      let _319_v101;
      _319_v101 = _dafny.Seq.of(_318_v100, _318_v100);
      _316_v98 = (_316_v98).update(_319_v101, (_235_v52).f15);
      let _320_v102;
      let _321_v103;
      let _322_v104;
      let _out21;
      let _out22;
      let _out23;
      let _outcollector7 = _module.__default.m0(_171_v0, new BigNumber((_175_v3).length), ((_171_v0) ? (_171_v0) : (false)), _174_globalState);
      _out21 = _outcollector7[0];
      _out22 = _outcollector7[1];
      _out23 = _outcollector7[2];
      _320_v102 = _out21;
      _321_v103 = _out22;
      _322_v104 = _out23;
      _322_v104 = !(_171_v0);
      process.stdout.write(_dafny.toString(_171_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_172_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_174_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_174_globalState).f1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_174_globalState).f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_174_globalState.f5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_174_globalState).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_175_v3, _dafny.Seq.of(true, false, true, true, false, false, true, false, true, true, true, true, false, false, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_176_v4, _dafny.Seq.of(_dafny.Seq.of(true, true, true, true, false, false, true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("fyhqxeqerfyhqxeqer")).updateUnsafe(true,_dafny.Seq.UnicodeFromString("fyhqxeqerfyhqxeqer")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_180_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_189_v15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v16)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_191_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_202_v25)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_203_v26).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_204_v27).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[_dafny.ZERO]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[_dafny.ONE]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(2)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(3)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(4)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(5)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(6)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(7)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(8)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(9)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(10)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(11)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(12)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(13)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(14)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(15)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(16)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(17)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(18)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(19)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(20)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(21)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(22)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(23)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_205_v28)[new BigNumber(24)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_235_v52).f15)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v53).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v54).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fyhqxeqer"), _dafny.Seq.UnicodeFromString("fyhqxeqer")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v55).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v55).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_v55).dtor_cf7)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v55).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_307_v89).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_308_v90).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(589), new BigNumber(-1), new BigNumber(-720), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(44), new BigNumber(-2), new BigNumber(7))).updateUnsafe(false,_dafny.Seq.of(_dafny.ONE, new BigNumber(4), _dafny.ZERO)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_309_v91, _dafny.Seq.of(_dafny.ONE, new BigNumber(4), _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v92).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xndmxbthy"),_dafny.Seq.UnicodeFromString("fyhqxeqer")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_311_v93).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xndmxbthy"),_dafny.Seq.UnicodeFromString("fyhqxeqer")),false).updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fyhqxeqer"),_dafny.Seq.UnicodeFromString("fyhqxeqer")),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_313_v95));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_314_v96).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_315_v97, _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_316_v98).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v99).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_318_v100).dtor_cf10).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_319_v101, _dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO)), _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_320_v102));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_321_v103));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_322_v104));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3() {
      let $dt = new D1(1);
      return $dt;
    }
    static create_DC4(cf5, cf6, cf7, cf8) {
      let $dt = new D1(2);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(3);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D1(4);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC1() { return this.$tag === 3; }
    get is_DC5() { return this.$tag === 4; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 4) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, false, false);
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
    static create_DC7(cf11, cf12) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D2(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, _dafny.ZERO);
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
    static create_DC10(cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC11(cf17, cf18) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D3(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO, _dafny.Seq.of());
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
    static create_DC14(cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + this.cf20.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(null, null);
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
    static create_DC16(cf24, cf25) {
      let $dt = new D5(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC17(cf26) {
      let $dt = new D5(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC15(cf23) {
      let $dt = new D5(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf26 === other.cf26;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC19(cf28, cf29, cf30) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf27) {
      let $dt = new D6(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC19(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.Map.Empty);
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
    static create_DC21(cf32, cf33) {
      let $dt = new D7(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC22(cf34, cf35, cf36) {
      let $dt = new D7(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC20(cf31) {
      let $dt = new D7(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf31 === other.cf31;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21(_dafny.ZERO, false);
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
    static create_DC23(cf37) {
      let $dt = new D8(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf37, other.cf37);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf39) {
      let $dt = new D9(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC26(cf40) {
      let $dt = new D9(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC27(cf41) {
      let $dt = new D9(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC24(cf38) {
      let $dt = new D9(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf41 === other.cf41;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f5 = _dafny.Seq.UnicodeFromString("");
      this.f14 = [];
      this._f1 = _dafny.Seq.of();
      this._f2 = false;
      this._f3 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      return;
    }
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
    get f4() {
      let _this = this;
      return _this._f4;
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
    get f9() {
      let _this = this;
      return _this._f9;
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
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f15 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f15) {
      let _this = this;
      (_this)._f15 = f15;
      return;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
