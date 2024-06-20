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
    static fm3(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!((_module.D5.create_DC13(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length),new BigNumber((_dafny.Seq.UnicodeFromString("q")).length))).length))).length), _dafny.Seq.of(_module.D3.create_DC9(), _module.D3.create_DC9(), _module.D3.create_DC9()))).dtor_cf23) || (true),(new BigNumber((_dafny.Seq.UnicodeFromString("enyymaag")).length)).isLessThan(new BigNumber(319)));
    };
    static fm4(p0, globalState) {
      return _module.__default.safeDivisionInt(new BigNumber(997), _module.__default.safeDivisionInt((_module.D11.create_DC26(true, new BigNumber(212))).dtor_cf49, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), function (_0_i0) {
        return new BigNumber(-486);
      })).length)));
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(942), new BigNumber(311))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(942)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(311)))) {
            _coll0.add((_1_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pay")).length))).cardinality()))));
          }
        }
        return _coll0;
      }()).length),new BigNumber(-559)),(_dafny.Set.fromElements(true)).IsProperSubsetOf(_dafny.Set.fromElements(true, true)));
    };
    static fm7(p0, p1, p2, globalState) {
      return ((new BigNumber((_dafny.Seq.of(!(false))).length)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(857), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-21),false)).length))).length)))).cardinality())))).minus(new BigNumber(791));
    };
    static fm8(p0, p1, p2, globalState) {
      if (_dafny.areEqual(_dafny.Seq.of(true), _dafny.Seq.of(true, true))) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-786),new BigNumber(-348)),_module.D0.create_DC0(false));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kgxwf")).length))).Elements) {
            let _2_v0 = _compr_1;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kgxwf")).length))).contains(_2_v0)) {
              _coll1.push([(_2_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D1.create_DC3(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-941)), function (_3_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))))).length)),new BigNumber(-24)]);
            }
          }
          return _coll1;
        }(),_module.D0.create_DC0(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(864), new BigNumber(366))) {
            let _4_v1 = _compr_2;
            if (((new BigNumber(864)).isLessThanOrEqualTo(_4_v1)) && ((_4_v1).isLessThan(new BigNumber(366)))) {
              _coll2.push([(_4_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),false)).length)),new BigNumber((_dafny.Seq.of(!(true), false, true)).length)]);
            }
          }
          return _coll2;
        }(),_module.D0.create_DC0(false)));
      }
    };
    static fm9(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-995)), function (_5_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length),(_dafny.Set.fromElements(true)).IsDisjointFrom(_dafny.Set.fromElements(false, true)));
    };
    static fm10(globalState) {
      return _dafny.Seq.UnicodeFromString("csqbv");
    };
    static fm11(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(939), new BigNumber(8))) {
          let _6_v0 = _compr_3;
          if (((new BigNumber(939)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(8)))) {
            _coll3.push([(_6_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),true)).length)),false]);
          }
        }
        return _coll3;
      }(), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ymwukergo")).length)),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("aax")).length),true), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ahojjj")).length)),false))).Union(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Seq.of(function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(516), new BigNumber(408))) {
            let _7_v1 = _compr_5;
            if (((new BigNumber(516)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(408)))) {
              _coll5.push([(_7_v1).plus(new BigNumber(531)),false]);
            }
          }
          return _coll5;
        }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, false)).length),true), function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_8_i0) {
            return new BigNumber(-995);
          })).Elements) {
            let _9_v2 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_8_i0) {
              return new BigNumber(-995);
            }), _9_v2)) {
              _coll6.push([_module.__default.safeModuloInt(_9_v2, new BigNumber((_dafny.Seq.UnicodeFromString("agsobcif")).length)),!(false)]);
            }
          }
          return _coll6;
        }())).Elements) {
          let _10_v3 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.of(function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(516), new BigNumber(408))) {
              let _7_v1 = _compr_7;
              if (((new BigNumber(516)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(408)))) {
                _coll7.push([(_7_v1).plus(new BigNumber(531)),false]);
              }
            }
            return _coll7;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, false)).length),true), function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_8_i0) {
              return new BigNumber(-995);
            })).Elements) {
              let _9_v2 = _compr_8;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_8_i0) {
                return new BigNumber(-995);
              }), _9_v2)) {
                _coll8.push([_module.__default.safeModuloInt(_9_v2, new BigNumber((_dafny.Seq.UnicodeFromString("agsobcif")).length)),!(false)]);
              }
            }
            return _coll8;
          }()), _10_v3)) {
            _coll4.add(_10_v3);
          }
        }
        return _coll4;
      }())).Union(function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(899),false))).Elements) {
          let _11_v4 = _compr_9;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(899),false))).contains(_11_v4)) {
            _coll9.add(_11_v4);
          }
        }
        return _coll9;
      }());
    };
    static fm12(p0, globalState) {
      return new _dafny.CodePoint('j'.codePointAt(0));
    };
    static fm13(globalState) {
      return !(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(189))).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_12_i0) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }))).Keys.Elements) {
          let _13_v0 = _compr_10;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(189))).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), function (_12_i0) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }))).contains(_13_v0)) {
            _coll10.push([_13_v0,false]);
          }
        }
        return _coll10;
      }()).equals(function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ogphrbwfb")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(407), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length)))).cardinality())))).Elements) {
          let _14_v1 = _compr_11;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ogphrbwfb")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(407), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length)))).cardinality()))), _14_v1)) {
            _coll11.push([_14_v1,!(true)]);
          }
        }
        return _coll11;
      }());
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_module.D3.create_DC9(), _module.D3.create_DC9(), _module.D3.create_DC9(), _module.D3.create_DC9(), _module.D3.create_DC9());
    };
    static fm15(p0, p1, p2, globalState) {
      return _module.D2.create_DC5(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true))));
    };
    static fm16(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(!(true))));
    };
    static fm17(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(!(true)), true), _dafny.Seq.of(false, true));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC10(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))), true, _dafny.Seq.UnicodeFromString("mlnbv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), function (_15_i0) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}), _module.__default.safeModuloInt(new BigNumber(771), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vojrh")).length))));
    };
    static fm21(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-52),!(true))).Merge((function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false), false)).length)))))).Elements) {
          let _16_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false), false)).length))))), _16_v0)) {
            _coll12.push([(_16_v0).multipliedBy(new BigNumber(492)),false]);
          }
        }
        return _coll12;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-583),!(false))));
    };
    static fm22(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wfd"), _dafny.Seq.UnicodeFromString("pkkiujvqb")), _dafny.Seq.UnicodeFromString("hhqapw"));
    };
    static fm23(p0, p1, p2, globalState) {
      return _module.D2.create_DC5((_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))));
    };
    static fm24(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-244)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(220)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(148)))).Elements) {
          let _17_v0 = _compr_13;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(220)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(148)))).contains(_17_v0)) {
            _coll13.add(_17_v0);
          }
        }
        return _coll13;
      }()).length)))).Difference((_dafny.MultiSet.fromElements(new BigNumber(((_dafny.Seq.of(false))).length), new BigNumber(207))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(872))).length))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_18_i0) {
        return new BigNumber(226);
      })).length))));
    };
    static fm29(globalState) {
      return (_dafny.MultiSet.fromElements(!(!(true)), true, true, !(true), !(true))).Difference((_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(false)));
    };
    static fm30(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(!(false), !(true), false))).length),new BigNumber((_dafny.Set.fromElements(new BigNumber(731))).length));
    };
    static fm31(globalState) {
      return _module.D3.create_DC9();
    };
    static fm32(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, !(true), true, false),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-370),new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false)),function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of _dafny.IntegerRange(new BigNumber(215), new BigNumber(835))) {
          let _19_v0 = _compr_14;
          if (((new BigNumber(215)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(835)))) {
            _coll14.push([(_19_v0).multipliedBy(new BigNumber((function () {
              let _coll15 = new _dafny.Set();
              for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-54), new BigNumber(615))) {
                let _21_v1 = _compr_15;
                if (((new BigNumber(-54)).isLessThanOrEqualTo(_21_v1)) && ((_21_v1).isLessThan(new BigNumber(615)))) {
                  _coll15.add((_21_v1).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vrmfs")).length))).length)));
                }
              }
              return _coll15;
            }()).length)),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), function (_20_i0) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("gavh")).length);
            })).length)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("hdtu")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),true)).length))).cardinality()))).length)]);
          }
        }
        return _coll14;
      }()));
    };
    static fm33(p0, p1, globalState) {
      if (!_dafny.areEqual(_dafny.Seq.of(false), _dafny.Seq.of(true, true))) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("r"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vvnpdwre"), _dafny.Seq.UnicodeFromString("uqmklfrag")));
    };
    static fm35(p0, p1, p2, globalState) {
      let _source0 = _module.D5.create_DC14(false, new BigNumber(49), new BigNumber((_dafny.Seq.of(new BigNumber(990))).length));
      if (_source0.is_DC13) {
        let _22___mcc_h0 = (_source0).cf23;
        let _23___mcc_h1 = (_source0).cf24;
        let _24___mcc_h2 = (_source0).cf25;
        let _25_cf25 = _24___mcc_h2;
        let _26_cf24 = _23___mcc_h1;
        let _27_cf23 = _22___mcc_h0;
        return _25_cf25;
      } else if (_source0.is_DC14) {
        let _28___mcc_h3 = (_source0).cf26;
        let _29___mcc_h4 = (_source0).cf27;
        let _30___mcc_h5 = (_source0).cf28;
        let _31_cf28 = _30___mcc_h5;
        let _32_cf27 = _29___mcc_h4;
        let _33_cf26 = _28___mcc_h3;
        return _dafny.Seq.of(_module.D3.create_DC9());
      } else {
        let _34___mcc_h6 = (_source0).cf22;
        let _35_cf22 = _34___mcc_h6;
        if (false) {
          return _dafny.Seq.of(_module.D3.create_DC9(), _module.D3.create_DC9());
        } else {
          return _dafny.Seq.of(_module.D3.create_DC9(), _module.D3.create_DC9(), _module.D3.create_DC9());
        }
      }
    };
    static fm36(globalState) {
      if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(673),_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("cwaejs")).length))).length)), new BigNumber(213), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-927)), function (_36_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false)).length))).length);
      })).length)))).length), new BigNumber(732), new BigNumber((_dafny.Set.fromElements(true)).length))).contains(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)))).length))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(515),_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_37_i1) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(false), true)).cardinality())),_dafny.Seq.UnicodeFromString("y"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false, false)).length),_dafny.Seq.UnicodeFromString("ay")));
      }
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(false)).Union((_dafny.Set.fromElements(false, false)).Difference(_dafny.Set.fromElements((_module.D5.create_DC13(!(!(false)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(546)), function (_38_i0) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})).length), _dafny.Seq.of(_module.D3.create_DC9()))).dtor_cf23, false)));
    };
    static fm38(p0, p1, globalState) {
      return (function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).cardinality()))).Elements) {
          let _39_v0 = _compr_16;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).cardinality()))).contains(_39_v0)) {
            _coll16.push([(_39_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))).cardinality())),!(!(!(false)))]);
          }
        }
        return _coll16;
      }()).Merge(((false) ? (function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(715), new BigNumber(203))) {
          let _40_v1 = _compr_17;
          if (((new BigNumber(715)).isLessThanOrEqualTo(_40_v1)) && ((_40_v1).isLessThan(new BigNumber(203)))) {
            _coll17.push([_module.__default.safeDivisionInt(_40_v1, new BigNumber(-514)),true]);
          }
        }
        return _coll17;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).cardinality()),true))));
    };
    static fm39(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(true), _dafny.Seq.of(false, false)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)))).length),!(!(true))));
    };
    static fm40(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_41_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dce")).length));
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("oeo")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(848)), function (_42_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length))).length), new BigNumber(157)));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return _module.D5.create_DC13(!(!(!(!(true))) || (false)), _module.__default.safeModuloInt(new BigNumber(556), new BigNumber(662)), _dafny.Seq.of(_module.D3.create_DC9()));
    };
    static fm42(p0, p1, p2, globalState) {
      return (function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(604), new BigNumber(271))) {
          let _43_v0 = _compr_18;
          if (((new BigNumber(604)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(271)))) {
            _coll18.add(_module.__default.safeModuloInt(_43_v0, new BigNumber((_dafny.Seq.UnicodeFromString("po")).length)));
          }
        }
        return _coll18;
      }()).Intersect((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(!(!(false)))).length))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_44_i0) {
        return new BigNumber((_dafny.Seq.of(true, true)).length);
      })).length),false)).length), new BigNumber((_dafny.Set.fromElements(true, true)).length))));
    };
    static fm43(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), function (_45_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(993)));
    };
    static fm44(p0, globalState) {
      return _module.D5.create_DC14((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("dwcbpk")).length)))), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), function (_46_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}))).length)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("evt")).length)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length),!(false))).length)).minus(new BigNumber(305)));
    };
    static fm45(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),!(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, !(true)))).cardinality()))).length)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("qgmkk")).length)));
    };
    static fm46(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Set.fromElements(new BigNumber(-257), new BigNumber(309), new BigNumber((_dafny.Set.fromElements(!(false))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of _dafny.IntegerRange(new BigNumber(-750), new BigNumber(898))) {
          let _47_v0 = _compr_19;
          if (((new BigNumber(-750)).isLessThanOrEqualTo(_47_v0)) && ((_47_v0).isLessThan(new BigNumber(898)))) {
            _coll19.add((_47_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D3.create_DC9()))).cardinality())));
          }
        }
        return _coll19;
      }()))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new BigNumber(385)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).cardinality())))));
    };
    static fm47(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC5(_dafny.MultiSet.fromElements(!(false), true, !(false), true, true)),((true) ? (!(true)) : (false)));
    };
    static Main(__noArgsParameter) {
      let _48_v0;
      let _nw0 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _48_v0 = _nw0;
      let _49_v1;
      let _nw1 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _49_v1 = _nw1;
      let _50_v2;
      _50_v2 = new BigNumber(548);
      let _51_v3;
      _51_v3 = false;
      let _52_v4;
      _52_v4 = _dafny.Seq.of(_51_v3, _51_v3);
      let _53_v5;
      _53_v5 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_50_v2), new BigNumber((_52_v4).length));
      let _54_v6;
      let _init0 = function (_55_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      };
      let _nw2 = Array((new BigNumber(9)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _54_v6 = _nw2;
      let _56_globalState;
      let _nw3 = new _module.GlobalState();
      _nw3.__ctor(true, true, false, true, new BigNumber(432), new BigNumber(-409), new BigNumber(541), _48_v0, new BigNumber(766), true, new _dafny.CodePoint('n'.codePointAt(0)), false, new BigNumber(-713), _49_v1, new BigNumber(804), _53_v5, _54_v6);
      _56_globalState = _nw3;
      let _57_v7;
      _57_v7 = _dafny.Set.fromElements(_50_v2);
      let _58_v8;
      _58_v8 = _dafny.Seq.of(_dafny.Set.fromElements(_50_v2), _57_v7, _57_v7);
      let _59_v9;
      _59_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_58_v8).length),new BigNumber((_57_v7).length));
      let _60_v10;
      _60_v10 = _dafny.Seq.UnicodeFromString("yvxdaq");
      let _61_v11;
      _61_v11 = new _dafny.CodePoint('v'.codePointAt(0));
      let _62_v12;
      _62_v12 = _dafny.Map.Empty.slice().updateUnsafe(_59_v9,_dafny.Seq.update(_60_v10, _module.__default.safeIndex(new BigNumber((_57_v7).length), new BigNumber((_60_v10).length)), _61_v11));
      let _index0 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length));
      (_49_v1)[_index0] = (((_62_v12).contains(_59_v9)) ? ((_62_v12).get(_59_v9)) : (_dafny.Seq.UnicodeFromString("angrdqnc")));
      let _index1 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length));
      let _rhs0 = _51_v3;
      let _rhs1 = _dafny.Seq.UnicodeFromString("wagfakwvy");
      let _rhs2 = _51_v3;
      let _lhs0 = _56_globalState;
      let _lhs1 = _49_v1;
      let _lhs2 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length));
      let _lhs3 = _56_globalState;
      _lhs0.f3 = _rhs0;
      _lhs1[_lhs2] = _rhs1;
      _lhs3.f9 = _rhs2;
      if (_51_v3) {
        let _63_v13;
        _63_v13 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_51_v3);
        let _64_v14;
        _64_v14 = _module.D1.create_DC4(false, _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_50_v2,_51_v3)), _50_v2);
        let _65_v15;
        let _nw4 = Array((new BigNumber(27)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _51_v3;
        _nw4[(_dafny.ONE).toNumber()] = _51_v3;
        _nw4[(new BigNumber(2)).toNumber()] = true;
        _nw4[(new BigNumber(3)).toNumber()] = false;
        _nw4[(new BigNumber(4)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(5)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(6)).toNumber()] = true;
        _nw4[(new BigNumber(7)).toNumber()] = (_64_v14).dtor_cf9;
        _nw4[(new BigNumber(8)).toNumber()] = !(_51_v3);
        _nw4[(new BigNumber(9)).toNumber()] = !(_51_v3);
        _nw4[(new BigNumber(10)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(11)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(12)).toNumber()] = true;
        _nw4[(new BigNumber(13)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(14)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(15)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(16)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(17)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(18)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(19)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(20)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(21)).toNumber()] = _module.__default.fm13(_56_globalState);
        _nw4[(new BigNumber(22)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(23)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(24)).toNumber()] = !(!(_51_v3));
        _nw4[(new BigNumber(25)).toNumber()] = _51_v3;
        _nw4[(new BigNumber(26)).toNumber()] = _51_v3;
        _65_v15 = _nw4;
        let _66_v16;
        _66_v16 = _dafny.Set.fromElements(_65_v15);
        let _67_v17;
        let _nw5 = new _module.C5();
        _nw5.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(446)), ((_68_v1, _69_v2, _70_v3) => function (_71_i1) {
          return ((_68_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_68_v1).length))])[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_69_v2,_70_v3)).length), new BigNumber(((_68_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_68_v1).length))]).length))];
        })(_49_v1, _50_v2, _51_v3)), _51_v3, _module.__default.safeModuloInt(new BigNumber(969), (_dafny.ZERO).minus(_50_v2)), _63_v13, _65_v15, (_66_v16).Difference(_66_v16), _51_v3);
        _67_v17 = _nw5;
        let _index2 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length));
        (_48_v0)[_index2] = new BigNumber(884);
        let _index3 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length));
        (_48_v0)[_index3] = new BigNumber(7);
        let _72_v18;
        _72_v18 = _module.D11.create_DC26(_51_v3, new BigNumber((_dafny.Set.fromElements((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))])).length));
        let _73_v19;
        _73_v19 = _module.D11.create_DC28(_72_v18);
        let _74_v20;
        _74_v20 = _module.D11.create_DC28(_73_v19);
        let _75_v21;
        _75_v21 = _module.D11.create_DC28(_74_v20);
        let _pat_let_tv0 = _73_v19;
        _75_v21 = function (_pat_let0_0) {
          return function (_76_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_77_dt__update_hcf53_h0) {
                return _module.D11.create_DC28(_77_dt__update_hcf53_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_75_v21);
        if (_51_v3) {
          let _78_v22;
          _78_v22 = _dafny.Set.fromElements((_67_v17).f30);
          let _index4 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length));
          (_48_v0)[_index4] = (_dafny.ZERO).minus(new BigNumber((((_78_v22).Difference(_78_v22)).Difference((_78_v22).Union(_78_v22))).length));
          _60_v10 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(803)), ((_79_v11) => function (_80_i2) {
            return _79_v11;
          })(_61_v11));
          let _81_v23;
          _81_v23 = _dafny.Map.Empty.slice().updateUnsafe((_57_v7).Difference(_57_v7),((_51_v3) ? (_52_v4) : (_52_v4)));
          _81_v23 = (_81_v23).update((_57_v7).Difference(_57_v7), _dafny.Seq.of((_67_v17).f30, (_67_v17).f30, _51_v3, !(_51_v3), !(!((_67_v17).f30))));
          let _82_v24;
          _82_v24 = _module.D5.create_DC14((((_67_v17).f30) ? (_51_v3) : ((_67_v17).f30)), _module.__default.safeDivisionInt(new BigNumber(467), _50_v2), (_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]);
          _82_v24 = _82_v24;
          let _83_v25;
          _83_v25 = _dafny.Seq.of((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]);
          let _84_v26;
          _84_v26 = _dafny.MultiSet.fromElements(_61_v11);
          (_56_globalState).f14 = (_83_v25)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_60_v10, (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]), _module.__default.safeIndex(new BigNumber((_84_v26).cardinality()), new BigNumber((_dafny.Seq.Concat(_60_v10, (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))])).length)), _61_v11)).length), new BigNumber((_83_v25).length))];
        } else {
          let _85_v27;
          let _nw6 = new _module.C9();
          _nw6.__ctor(new BigNumber((_53_v5).cardinality()), _50_v2, _dafny.Map.Empty.slice().updateUnsafe(_61_v11,(_67_v17).f30), _65_v15, _66_v16, (_67_v17).f30);
          _85_v27 = _nw6;
          let _86_v28;
          _86_v28 = _dafny.Map.Empty.slice().updateUnsafe(_85_v27,(_dafny.ZERO).minus(_85_v27.f22));
          let _87_v29;
          let _88_v30;
          let _89_v31;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_67_v17).m0(((_51_v3) ? ((((_86_v28).contains(_85_v27)) ? ((_86_v28).get(_85_v27)) : ((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]))) : (_50_v2)), _51_v3, _85_v27.f22, _56_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _87_v29 = _out0;
          _88_v30 = _out1;
          _89_v31 = _out2;
          let _90_v32;
          _90_v32 = _dafny.Map.Empty.slice().updateUnsafe((_67_v17).f30,_module.__default.fm13(_56_globalState));
          let _91_v33;
          _91_v33 = _dafny.Map.Empty.slice().updateUnsafe(_59_v9,_89_v31);
          let _92_v34;
          _92_v34 = _dafny.Set.fromElements(_88_v30);
          let _93_v35;
          _93_v35 = _module.D7.create_DC16(_92_v34);
          _90_v32 = (_90_v32).update((_85_v27).fm1(_91_v33, _56_globalState), ((_93_v35).dtor_cf30).IsProperSubsetOf(_92_v34));
          let _94_v36;
          let _nw7 = new _module.C2();
          _nw7.__ctor((_66_v16).Union(_66_v16), true);
          _94_v36 = _nw7;
          let _95_v37;
          _95_v37 = _dafny.Seq.of(new BigNumber((_92_v34).length));
          let _96_v38;
          _96_v38 = _dafny.Seq.of(_95_v37, _95_v37, _95_v37, _95_v37);
          let _97_v39;
          let _nw8 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _97_v39 = _nw8;
          let _98_v40;
          _98_v40 = _module.D0.create_DC2(_88_v30, _50_v2, _97_v39);
          let _99_v41;
          _99_v41 = _dafny.Map.Empty.slice().updateUnsafe(_87_v29,new _dafny.CodePoint('x'.codePointAt(0)));
          let _100_v42;
          let _nw9 = Array((new BigNumber(25)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = _95_v37;
          _nw9[(_dafny.ONE).toNumber()] = _95_v37;
          _nw9[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_85_v27.f22, _85_v27.f22);
          _nw9[(new BigNumber(3)).toNumber()] = _module.__default.fm40(!(_88_v30), _88_v30, _56_globalState);
          _nw9[(new BigNumber(4)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(5)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(6)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_95_v37, _95_v37);
          _nw9[(new BigNumber(8)).toNumber()] = (_96_v38)[_module.__default.safeIndex(_85_v27.f22, new BigNumber((_96_v38).length))];
          _nw9[(new BigNumber(9)).toNumber()] = ((_87_v29) ? (_95_v37) : (_95_v37));
          _nw9[(new BigNumber(10)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(11)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(12)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(13)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_50_v2, _50_v2, _50_v2);
          _nw9[(new BigNumber(15)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), ((_101_v31) => function (_102_i3) {
            return (_dafny.ZERO).minus(_101_v31);
          })(_89_v31)), _95_v37);
          _nw9[(new BigNumber(17)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(18)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-558)), ((_103_v37) => function (_104_i4) {
            return new BigNumber((_103_v37).length);
          })(_95_v37));
          _nw9[(new BigNumber(20)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_95_v37, _dafny.Seq.of((_98_v40).dtor_cf6));
          _nw9[(new BigNumber(22)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), ((_105_v27) => function (_106_i5) {
            return _105_v27.f22;
          })(_85_v27));
          _nw9[(new BigNumber(23)).toNumber()] = _95_v37;
          _nw9[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_50_v2, new BigNumber((_99_v41).length)), _dafny.Seq.of(_85_v27.f22));
          _100_v42 = _nw9;
          _100_v42 = _100_v42;
          let _107_v45;
          let _nw10 = Array((new BigNumber(26)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = _59_v9;
          _nw10[(_dafny.ONE).toNumber()] = function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-832), new BigNumber(895))) {
              let _108_v43 = _compr_20;
              if (((new BigNumber(-832)).isLessThanOrEqualTo(_108_v43)) && ((_108_v43).isLessThan(new BigNumber(895)))) {
                _coll20.push([_module.__default.safeDivisionInt(_108_v43, _89_v31),_50_v2]);
              }
            }
            return _coll20;
          }();
          _nw10[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(193),_85_v27.f22);
          _nw10[(new BigNumber(3)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(4)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(5)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(6)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(7)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(8)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(9)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(10)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(11)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(12)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(13)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(14)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(15)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(16)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(17)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(18)).toNumber()] = function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(765), new BigNumber(4))) {
              let _109_v44 = _compr_21;
              if (((new BigNumber(765)).isLessThanOrEqualTo(_109_v44)) && ((_109_v44).isLessThan(new BigNumber(4)))) {
                _coll21.push([_module.__default.safeModuloInt(_109_v44, _85_v27.f22),new BigNumber((_dafny.Seq.of((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))])).length)]);
              }
            }
            return _coll21;
          }();
          _nw10[(new BigNumber(19)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(20)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(21)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(22)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(23)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(24)).toNumber()] = _59_v9;
          _nw10[(new BigNumber(25)).toNumber()] = _59_v9;
          _107_v45 = _nw10;
          let _110_v46;
          _110_v46 = _module.D7.create_DC17(_107_v45, _87_v29, (_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]);
          let _111_v47;
          _111_v47 = _module.D5.create_DC14((_92_v34).IsDisjointFrom(_92_v34), (_110_v46).dtor_cf33, new BigNumber(453));
          let _112_v48;
          _112_v48 = _module.D3.create_DC9();
          let _113_v49;
          _113_v49 = _dafny.Seq.of(_112_v48, _112_v48, _112_v48);
          let _114_v50;
          _114_v50 = _113_v49;
          let _index5 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_65_v15).length));
          (_65_v15)[_index5] = _88_v30;
          let _index6 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_65_v15).length));
          let _rhs3 = false;
          let _rhs4 = _module.__default.fm44(_module.__default.safeDivisionInt(_50_v2, (_67_v17).fm0(new BigNumber(930), _61_v11, _56_globalState)), _56_globalState);
          let _rhs5 = _114_v50;
          let _rhs6 = !((_67_v17).f30);
          let _lhs4 = _56_globalState;
          let _lhs5 = _65_v15;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_65_v15).length));
          _lhs4.f9 = _rhs3;
          _111_v47 = _rhs4;
          _114_v50 = _rhs5;
          _lhs5[_lhs6] = _rhs6;
        }
        let _115_v51;
        _115_v51 = _dafny.Seq.of((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]);
        let _116_v52;
        let _nw11 = new _module.C1();
        _nw11.__ctor();
        _116_v52 = _nw11;
        let _117_v53;
        let _nw12 = Array((new BigNumber(28)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _116_v52;
        _nw12[(_dafny.ONE).toNumber()] = _116_v52;
        _nw12[(new BigNumber(2)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(3)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(4)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(5)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(6)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(7)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(8)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(9)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(10)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(11)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(12)).toNumber()] = ((_51_v3) ? (_116_v52) : (_116_v52));
        _nw12[(new BigNumber(13)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(14)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(15)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(16)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(17)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(18)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(19)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(20)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(21)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(22)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(23)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(24)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(25)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(26)).toNumber()] = _116_v52;
        _nw12[(new BigNumber(27)).toNumber()] = _116_v52;
        _117_v53 = _nw12;
        let _index7 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_117_v53).length));
        (_117_v53)[_index7] = _116_v52;
        let _118_v54;
        let _nw13 = new _module.C8();
        _nw13.__ctor(new BigNumber((_53_v5).cardinality()), _51_v3, _65_v15, (_66_v16).Union(_66_v16), _dafny.Seq.IsPrefixOf(_52_v4, _52_v4), ((_48_v0)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_48_v0).length))]).plus(_50_v2), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),(_67_v17).f30));
        _118_v54 = _nw13;
        let _119_v55;
        _119_v55 = _dafny.Seq.of(_60_v10, (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(589)), function (_120_i6) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), (_67_v17).f29, _dafny.Seq.UnicodeFromString("geplenmxg"));
        let _121_v56;
        _121_v56 = _dafny.Seq.of(_53_v5);
        let _index8 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_117_v53).length));
        let _rhs7 = (_67_v17).f30;
        let _rhs8 = (_118_v54).f20;
        let _rhs9 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xwcwqwn"), (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]), _module.__default.safeIndex(new BigNumber((_119_v55).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xwcwqwn"), (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))])).length)), ((_67_v17).f29)[_module.__default.safeIndex((((_59_v9).contains(new BigNumber((_121_v56).length))) ? ((_59_v9).get(new BigNumber((_121_v56).length))) : ((_118_v54).f20)), new BigNumber(((_67_v17).f29).length))])).length));
        let _rhs10 = _116_v52;
        let _rhs11 = _118_v54;
        let _lhs7 = _56_globalState;
        let _lhs8 = _56_globalState;
        let _lhs9 = _117_v53;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_117_v53).length));
        _lhs7.f3 = _rhs7;
        _lhs8.f14 = _rhs8;
        _115_v51 = _rhs9;
        _lhs9[_lhs10] = _rhs10;
        _118_v54 = _rhs11;
      } else {
        let _122_v57;
        let _init1 = ((_123_v3) => function (_124_i7) {
          return _123_v3;
        })(_51_v3);
        let _nw14 = Array((new BigNumber(28)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
          _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _122_v57 = _nw14;
        let _index9 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length));
        (_122_v57)[_index9] = (_51_v3) && (_51_v3);
        let _125_v58;
        _125_v58 = _dafny.Set.fromElements(_122_v57);
        let _126_v59;
        _126_v59 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,_51_v3);
        let _127_v60;
        let _nw15 = new _module.C8();
        _nw15.__ctor(_50_v2, _51_v3, _122_v57, _125_v58, _51_v3, _50_v2, _126_v59);
        _127_v60 = _nw15;
        let _128_v61;
        _128_v61 = _dafny.Seq.of(_127_v60);
        let _129_v62;
        _129_v62 = _dafny.Set.fromElements(_128_v61);
        let _index10 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length));
        (_122_v57)[_index10] = (_129_v62).IsProperSubsetOf(_129_v62);
        if ((_52_v4)[_module.__default.safeIndex(_127_v60.f23, new BigNumber((_52_v4).length))]) {
          let _130_v63;
          _130_v63 = _module.D10.create_DC22(_57_v7);
          let _131_v64;
          _131_v64 = _dafny.Map.Empty.slice().updateUnsafe((_122_v57)[_module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length))],_130_v63);
          let _pat_let_tv1 = _57_v7;
          let _132_v65;
          let _133_v66;
          let _134_v67;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = (_127_v60).m2(new BigNumber(((_131_v64).update(_51_v3, function (_pat_let2_0) {
            return function (_135_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_136_dt__update_hcf42_h0) {
                  return _module.D10.create_DC22(_136_dt__update_hcf42_h0);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let2_0);
          }(_130_v63))).length), _56_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _132_v65 = _out3;
          _133_v66 = _out4;
          _134_v67 = _out5;
          let _index11 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_54_v6).length));
          (_54_v6)[_index11] = new _dafny.CodePoint('o'.codePointAt(0));
          let _index12 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_54_v6).length));
          (_54_v6)[_index12] = _61_v11;
          let _137_v68;
          _137_v68 = _dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(false, _51_v3, _133_v66, _133_v66, (_127_v60).f24));
          let _138_v69;
          let _nw16 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _138_v69 = _nw16;
          let _139_v70;
          _139_v70 = _dafny.Map.Empty.slice().updateUnsafe(_127_v60.f23,_138_v69);
          let _140_v71;
          _140_v71 = _dafny.MultiSet.fromElements(_138_v69, _138_v69, (((_139_v70).contains(_127_v60.f23)) ? ((_139_v70).get(_127_v60.f23)) : (_138_v69)));
          let _rhs12 = _134_v67;
          let _rhs13 = _dafny.Seq.Concat(_137_v68, _137_v68);
          let _rhs14 = (_dafny.ZERO).minus(_134_v67);
          let _rhs15 = _122_v57;
          let _rhs16 = (((_140_v71).contains(_138_v69)) ? ((_140_v71).get(_138_v69)) : (_127_v60.f23));
          let _lhs11 = _56_globalState;
          let _lhs12 = _127_v60;
          _lhs11.f14 = _rhs12;
          _137_v68 = _rhs13;
          _134_v67 = _rhs14;
          _122_v57 = _rhs15;
          _lhs12.f23 = _rhs16;
          let _141_v72;
          _141_v72 = _dafny.Map.Empty.slice().updateUnsafe(_132_v65,_53_v5);
          _141_v72 = (_141_v72).Merge((_141_v72).Merge(_141_v72));
          _50_v2 = _127_v60.f23;
        } else {
          let _142_v73;
          let _nw17 = new _module.C9();
          _nw17.__ctor(_50_v2, _50_v2, _126_v59, _122_v57, _125_v58, (_127_v60).f24);
          _142_v73 = _nw17;
          let _143_v74;
          _143_v74 = _module.D10.create_DC23((_127_v60).f24, (_127_v60).f24, _142_v73);
          let _144_v75;
          _144_v75 = _module.D10.create_DC24(_143_v74);
          let _145_v76;
          let _nw18 = Array((new BigNumber(5)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _144_v75;
          _nw18[(_dafny.ONE).toNumber()] = _module.D10.create_DC24(_143_v74);
          _nw18[(new BigNumber(2)).toNumber()] = _144_v75;
          _nw18[(new BigNumber(3)).toNumber()] = _144_v75;
          _nw18[(new BigNumber(4)).toNumber()] = _module.D10.create_DC24(_143_v74);
          _145_v76 = _nw18;
          let _146_v77;
          _146_v77 = _dafny.Map.Empty.slice().updateUnsafe(_50_v2,_145_v76);
          let _147_v78;
          let _nw19 = Array((new BigNumber(5)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _145_v76;
          _nw19[(_dafny.ONE).toNumber()] = (((_146_v77).contains((_142_v73).f20)) ? ((_146_v77).get((_142_v73).f20)) : (_145_v76));
          _nw19[(new BigNumber(2)).toNumber()] = _145_v76;
          _nw19[(new BigNumber(3)).toNumber()] = _145_v76;
          _nw19[(new BigNumber(4)).toNumber()] = _145_v76;
          _147_v78 = _nw19;
          let _148_v79;
          _148_v79 = _module.D14.create_DC34(_54_v6);
          let _149_v80;
          _149_v80 = _dafny.MultiSet.fromElements(_148_v79);
          let _150_v81;
          _150_v81 = _dafny.Map.Empty.slice().updateUnsafe((_50_v2).isLessThan(new BigNumber(502)),(_149_v80).update(_148_v79, _module.__default.abs(_127_v60.f23)));
          let _rhs17 = _147_v78;
          let _rhs18 = (_127_v60).f24;
          let _rhs19 = _150_v81;
          let _rhs20 = new BigNumber(748);
          let _lhs13 = _56_globalState;
          let _lhs14 = _127_v60;
          _147_v78 = _rhs17;
          _lhs13.f9 = _rhs18;
          _150_v81 = _rhs19;
          _lhs14.f23 = _rhs20;
          let _151_v82;
          _151_v82 = _dafny.MultiSet.fromElements((_122_v57)[_module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length))]);
          let _152_v83;
          let _nw20 = Array((new BigNumber(7)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _59_v9;
          _nw20[(_dafny.ONE).toNumber()] = _59_v9;
          _nw20[(new BigNumber(2)).toNumber()] = (_59_v9).update(_127_v60.f23, new BigNumber((_dafny.Set.fromElements(_51_v3, (_127_v60).f24)).length));
          _nw20[(new BigNumber(3)).toNumber()] = _59_v9;
          _nw20[(new BigNumber(4)).toNumber()] = _module.__default.fm30((_127_v60).f24, _56_globalState);
          _nw20[(new BigNumber(5)).toNumber()] = _59_v9;
          _nw20[(new BigNumber(6)).toNumber()] = _59_v9;
          _152_v83 = _nw20;
          let _153_v84;
          _153_v84 = _module.D7.create_DC17(_152_v83, _51_v3, _127_v60.f23);
          let _154_v85;
          _154_v85 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_61_v11,(_127_v60).f24));
          let _155_v86;
          let _156_v87;
          let _157_v88;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = (_127_v60).m0((((_151_v82).contains((_153_v84).dtor_cf32)) ? ((_151_v82).get((_153_v84).dtor_cf32)) : (_50_v2)), _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(201)), ((_158_v73) => function (_159_i8) {
            return _158_v73.f21;
          })(_142_v73)), _154_v85), _127_v60.f23, _56_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _155_v86 = _out6;
          _156_v87 = _out7;
          _157_v88 = _out8;
          _156_v87 = (_142_v73).f18;
          let _160_v89;
          let _nw21 = new _module.C2();
          _nw21.__ctor(_dafny.Set.fromElements(_142_v73.f19, _142_v73.f19, _142_v73.f19, _122_v57, _122_v57), false);
          _160_v89 = _nw21;
          let _161_v90;
          let _nw22 = new _module.C2();
          _nw22.__ctor(_125_v58, (new BigNumber(562)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_60_v10).length),_160_v89)).length)));
          _161_v90 = _nw22;
          let _162_v91;
          _162_v91 = _dafny.Seq.of(_161_v90);
          _161_v90 = (_162_v91)[_module.__default.safeIndex((_50_v2).plus(new BigNumber((_60_v10).length)), new BigNumber((_162_v91).length))];
          let _163_v92;
          let _164_v93;
          let _165_v94;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = (_127_v60).m0(_module.__default.safeDivisionInt(new BigNumber(((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]).length), (_dafny.ZERO).minus((_142_v73).f20)), ((_142_v73).f18) && ((_127_v60).f24), _50_v2, _56_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _163_v92 = _out9;
          _164_v93 = _out10;
          _165_v94 = _out11;
        }
        let _166_v95;
        _166_v95 = _dafny.Map.Empty.slice().updateUnsafe(_50_v2,(_122_v57)[_module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length))]);
        let _167_v96;
        _167_v96 = _dafny.MultiSet.fromElements(_60_v10);
        let _168_v98;
        _168_v98 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_167_v96).Elements) {
            let _169_v97 = _compr_22;
            if ((_167_v96).contains(_169_v97)) {
              _coll22.add(_169_v97);
            }
          }
          return _coll22;
        }(),(_122_v57)[_module.__default.safeIndex(new BigNumber(154), new BigNumber((_122_v57).length))]);
        _166_v95 = (_166_v95).update(new BigNumber(-78), (((_168_v98).contains(_dafny.Set.fromElements((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]))) ? ((_168_v98).get(_dafny.Set.fromElements((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]))) : (_51_v3)));
        (_127_v60).f23 = _module.__default.safeModuloInt(_127_v60.f23, _module.__default.safeDivisionInt(_50_v2, new BigNumber(797)));
        let _170_v99;
        let _nw23 = new _module.C1();
        _nw23.__ctor();
        _170_v99 = _nw23;
      }
      let _hi0 = _50_v2;
      for (let _171_i9 = _50_v2; _171_i9.isLessThan(_hi0); _171_i9 = _171_i9.plus(_dafny.ONE)) {
        let _172_v100;
        let _init2 = ((_173_v3) => function (_174_i10) {
          return _173_v3;
        })(_51_v3);
        let _nw24 = Array((new BigNumber(27)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
          _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _172_v100 = _nw24;
        let _175_v101;
        _175_v101 = _dafny.Set.fromElements(_172_v100);
        let _176_v102;
        _176_v102 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,_51_v3);
        let _177_v103;
        let _nw25 = new _module.C8();
        _nw25.__ctor(_50_v2, _51_v3, _172_v100, _175_v101, _51_v3, _171_i9, _176_v102);
        _177_v103 = _nw25;
        let _178_v104;
        _178_v104 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_177_v103);
        let _179_v105;
        _179_v105 = _dafny.Seq.of(new BigNumber(16), _177_v103.f23);
        let _180_v107;
        _180_v107 = _dafny.Set.fromElements(_module.__default.fm33(new BigNumber((_dafny.Seq.UnicodeFromString("lx")).length), _51_v3, _56_globalState), _61_v11);
        let _181_v108;
        let _nw26 = Array((new BigNumber(7)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = (_178_v104).equals(_178_v104);
        _nw26[(_dafny.ONE).toNumber()] = (function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of (_179_v105).Elements) {
            let _182_v106 = _compr_23;
            if (_dafny.Seq.contains(_179_v105, _182_v106)) {
              _coll23.add((_182_v106).minus(new BigNumber(481)));
            }
          }
          return _coll23;
        }()).IsSubsetOf(_57_v7);
        _nw26[(new BigNumber(2)).toNumber()] = false;
        _nw26[(new BigNumber(3)).toNumber()] = false;
        _nw26[(new BigNumber(4)).toNumber()] = _51_v3;
        _nw26[(new BigNumber(5)).toNumber()] = (_177_v103).f24;
        _nw26[(new BigNumber(6)).toNumber()] = (_180_v107).IsProperSubsetOf(_180_v107);
        _181_v108 = _nw26;
        let _index13 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_181_v108).length));
        (_181_v108)[_index13] = (_177_v103).f24;
        let _183_v109;
        _183_v109 = _dafny.Map.Empty.slice().updateUnsafe(_171_i9,_61_v11);
        let _184_v110;
        let _nw27 = new _module.C4();
        _nw27.__ctor(new BigNumber((_183_v109).length), new BigNumber(-688), _175_v101, true);
        _184_v110 = _nw27;
        let _185_v111;
        _185_v111 = _dafny.Seq.of(_184_v110);
        let _index14 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_181_v108).length));
        (_181_v108)[_index14] = !(_dafny.Seq.IsProperPrefixOf(_185_v111, _dafny.Seq.Concat(_185_v111, _185_v111)));
        let _186_v112;
        _186_v112 = _module.D3.create_DC10((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))], _51_v3, (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), ((_187_v11) => function (_188_i11) {
  return _187_v11;
})(_61_v11)), (_184_v110).f25);
        let _189_v113;
        let _out12;
        _out12 = (_177_v103).m1(new BigNumber(-149), _dafny.Seq.contains((_186_v112).dtor_cf18, new _dafny.CodePoint('k'.codePointAt(0))), (_171_i9).multipliedBy(_171_i9), _51_v3, _56_globalState);
        _189_v113 = _out12;
        _51_v3 = _51_v3;
        (_56_globalState).f3 = (_177_v103).f24;
      }
      _52_v4 = _52_v4;
      let _190_i12;
      _190_i12 = _dafny.ZERO;
      L0: {
        while ((_50_v2).isEqualTo(_50_v2)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_190_i12)) {
              break L0;
            }
            _190_i12 = (_190_i12).plus(_dafny.ONE);
            let _191_v114;
            _191_v114 = _dafny.Map.Empty.slice().updateUnsafe(_52_v4,_59_v9);
            let _source1 = _module.D9.create_DC19(_191_v114);
            if (_source1.is_DC20) {
              let _192___mcc_h0 = (_source1).cf36;
              let _193___mcc_h1 = (_source1).cf37;
              let _194___mcc_h2 = (_source1).cf38;
              let _195___mcc_h3 = (_source1).cf39;
              let _196___mcc_h4 = (_source1).cf40;
              let _197_cf40 = _196___mcc_h4;
              let _198_cf39 = _195___mcc_h3;
              let _199_cf38 = _194___mcc_h2;
              let _200_cf37 = _193___mcc_h1;
              let _201_cf36 = _192___mcc_h0;
              let _202_v115;
              _202_v115 = _dafny.Map.Empty.slice().updateUnsafe(_199_cf38,(((_53_v5).contains(new BigNumber((_201_cf36).length))) ? ((_53_v5).get(new BigNumber((_201_cf36).length))) : (_200_cf37)));
              let _203_v116;
              _203_v116 = _module.D0.create_DC0(_199_cf38);
              let _204_v117;
              let _init3 = ((_205_v3) => function (_206_i13) {
                return _205_v3;
              })(_51_v3);
              let _nw28 = Array((new BigNumber(21)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw28.length); _i0_3++) {
                _nw28[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _204_v117 = _nw28;
              let _207_v118;
              _207_v118 = _dafny.Set.fromElements(_204_v117);
              let _208_v119;
              let _nw29 = new _module.C4();
              _nw29.__ctor(_200_cf37, (_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.fm7(_200_cf37, _202_v115, _203_v116, _56_globalState))).minus(_200_cf37)), (_207_v118).Union(_207_v118), _51_v3);
              _208_v119 = _nw29;
              let _209_v120;
              _209_v120 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_204_v117);
              _209_v120 = ((_209_v120).Merge(_209_v120)).Merge(_209_v120);
              _209_v120 = (_209_v120).update(_51_v3, _204_v117);
              let _210_v121;
              let _nw30 = new _module.C2();
              _nw30.__ctor(_207_v118, _51_v3);
              _210_v121 = _nw30;
              let _211_v122;
              _211_v122 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(848),_210_v121);
              let _212_v123;
              _212_v123 = _dafny.Map.Empty.slice().updateUnsafe(_200_cf37,_211_v122);
              _212_v123 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(577),_211_v122);
            } else if (_source1.is_DC19) {
              let _213___mcc_h5 = (_source1).cf35;
              let _214_cf35 = _213___mcc_h5;
              let _215_v124;
              let _nw31 = Array((new BigNumber(28)).toNumber()).fill(false);
              _215_v124 = _nw31;
              let _216_v125;
              _216_v125 = _dafny.Set.fromElements(_215_v124, _215_v124, _215_v124);
              let _217_v126;
              let _nw32 = new _module.C3();
              _nw32.__ctor(_215_v124, _216_v125, _51_v3);
              _217_v126 = _nw32;
              let _218_v127;
              _218_v127 = _dafny.Map.Empty.slice().updateUnsafe(_217_v126,_49_v1);
              let _219_v128;
              _219_v128 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_49_v1);
              let _220_v129;
              let _nw33 = Array((new BigNumber(11)).toNumber());
              _nw33[(_dafny.ZERO).toNumber()] = _49_v1;
              _nw33[(_dafny.ONE).toNumber()] = _49_v1;
              _nw33[(new BigNumber(2)).toNumber()] = (((_218_v127).contains(_217_v126)) ? ((_218_v127).get(_217_v126)) : (_49_v1));
              _nw33[(new BigNumber(3)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(4)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(5)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(6)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(7)).toNumber()] = (((_219_v128).contains(_51_v3)) ? ((_219_v128).get(_51_v3)) : (_49_v1));
              _nw33[(new BigNumber(8)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(9)).toNumber()] = _49_v1;
              _nw33[(new BigNumber(10)).toNumber()] = _49_v1;
              _220_v129 = _nw33;
              let _index15 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_220_v129).length));
              (_220_v129)[_index15] = _49_v1;
              let _221_v130;
              _221_v130 = _dafny.Map.Empty.slice().updateUnsafe(_52_v4,_dafny.Seq.of(_51_v3, _51_v3, _51_v3));
              let _index16 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_220_v129).length));
              let _rhs21 = _dafny.Seq.Concat(_52_v4, (((_221_v130).contains(_52_v4)) ? ((_221_v130).get(_52_v4)) : (_52_v4)));
              let _rhs22 = _49_v1;
              let _rhs23 = !(!(_51_v3));
              let _lhs15 = _220_v129;
              let _lhs16 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_220_v129).length));
              let _lhs17 = _56_globalState;
              _52_v4 = _rhs21;
              _lhs15[_lhs16] = _rhs22;
              _lhs17.f3 = _rhs23;
              (_56_globalState).f14 = _50_v2;
              let _222_v131;
              _222_v131 = _module.D9.create_DC19(_214_cf35);
              let _223_v132;
              _223_v132 = _module.D9.create_DC21(_222_v131);
              let _224_v133;
              _224_v133 = _module.D9.create_DC21(_222_v131);
              let _pat_let_tv2 = _223_v132;
              _224_v133 = function (_pat_let4_0) {
                return function (_225_dt__update__tmp_h2) {
                  return function (_pat_let5_0) {
                    return function (_226_dt__update_hcf41_h0) {
                      return _module.D9.create_DC21(_226_dt__update_hcf41_h0);
                    }(_pat_let5_0);
                  }(_pat_let_tv2);
                }(_pat_let4_0);
              }(_224_v133);
              _215_v124 = _215_v124;
            } else {
              let _227___mcc_h6 = (_source1).cf41;
              let _228_cf41 = _227___mcc_h6;
              let _229_v134;
              let _nw34 = Array((new BigNumber(6)).toNumber());
              _229_v134 = _nw34;
              let _230_v135;
              _230_v135 = _module.D3.create_DC9();
              let _231_v136;
              _231_v136 = _dafny.Seq.of(_230_v135);
              let _232_v137;
              _232_v137 = _module.D5.create_DC13(_51_v3, _50_v2, _231_v136);
              let _233_v138;
              _233_v138 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,_51_v3);
              let _234_v139;
              let _init4 = function (_235_i14) {
                return true;
              };
              let _nw35 = Array((new BigNumber(29)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw35.length); _i0_4++) {
                _nw35[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _234_v139 = _nw35;
              let _236_v140;
              _236_v140 = _dafny.Set.fromElements(_234_v139);
              let _237_v141;
              let _nw36 = new _module.C7();
              _nw36.__ctor(_50_v2, _233_v138, _234_v139, _236_v140, _51_v3);
              _237_v141 = _nw36;
              let _index17 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_229_v134).length));
              (_229_v134)[_index17] = (((_232_v137).dtor_cf23) ? (_237_v141) : (_237_v141));
              let _index18 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_229_v134).length));
              (_229_v134)[_index18] = _237_v141;
              let _index19 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length));
              (_49_v1)[_index19] = (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))];
              _59_v9 = (_59_v9).update(_50_v2, (_50_v2).plus(_50_v2));
              let _238_v142;
              _238_v142 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_50_v2);
              let _239_v143;
              _239_v143 = _module.D0.create_DC0(!(true));
              (_56_globalState).f14 = _module.__default.fm7(_50_v2, _238_v142, _239_v143, _56_globalState);
            }
            let _240_v144;
            _240_v144 = _dafny.MultiSet.fromElements(_51_v3);
            let _241_v145;
            _241_v145 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_50_v2, new BigNumber(((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]).length)),_module.D2.create_DC5(_240_v144));
            let _242_v146;
            _242_v146 = _module.D2.create_DC5(_240_v144);
            _241_v145 = (_241_v145).update((_dafny.ZERO).minus(_50_v2), _242_v146);
            let _243_v147;
            let _nw37 = Array((new BigNumber(25)).toNumber()).fill([]);
            _243_v147 = _nw37;
            let _rhs24 = _243_v147;
            let _rhs25 = _51_v3;
            let _rhs26 = _51_v3;
            let _rhs27 = _60_v10;
            let _lhs18 = _56_globalState;
            let _lhs19 = _56_globalState;
            _243_v147 = _rhs24;
            _lhs18.f3 = _rhs25;
            _lhs19.f9 = _rhs26;
            _60_v10 = _rhs27;
            let _244_v148;
            let _init5 = ((_245_v3) => function (_246_i15) {
              return _dafny.Set.fromElements(_245_v3);
            })(_51_v3);
            let _nw38 = Array((new BigNumber(29)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw38.length); _i0_5++) {
              _nw38[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _244_v148 = _nw38;
            let _247_v149;
            _247_v149 = _dafny.Set.fromElements(!(_51_v3));
            let _index20 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_244_v148).length));
            (_244_v148)[_index20] = _247_v149;
            let _index21 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_244_v148).length));
            (_244_v148)[_index21] = _247_v149;
          }
        }
      }
      if (_51_v3) {
        if (!((((_51_v3) ? (_50_v2) : (_50_v2))).isLessThanOrEqualTo(_50_v2))) {
          (_56_globalState).f14 = (((_53_v5).contains(_50_v2)) ? ((_53_v5).get(_50_v2)) : (_50_v2));
          let _248_v150;
          _248_v150 = _dafny.Seq.of(_60_v10);
          _248_v150 = _dafny.Seq.update(_248_v150, _module.__default.safeIndex(_50_v2, new BigNumber((_248_v150).length)), (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]);
          let _249_v151;
          let _nw39 = Array((new BigNumber(24)).toNumber()).fill([]);
          _249_v151 = _nw39;
          let _250_v152;
          _250_v152 = _dafny.Seq.of(_50_v2);
          let _251_v153;
          _251_v153 = _dafny.Seq.of(_250_v152, _250_v152);
          let _252_v154;
          let _nw40 = Array((new BigNumber(10)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = _250_v152;
          _nw40[(_dafny.ONE).toNumber()] = _250_v152;
          _nw40[(new BigNumber(2)).toNumber()] = (_251_v153)[_module.__default.safeIndex(new BigNumber((_52_v4).length), new BigNumber((_251_v153).length))];
          _nw40[(new BigNumber(3)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(4)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(5)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(6)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(7)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(8)).toNumber()] = _250_v152;
          _nw40[(new BigNumber(9)).toNumber()] = (_250_v152);
          _252_v154 = _nw40;
          let _index22 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_249_v151).length));
          (_249_v151)[_index22] = _252_v154;
          let _index23 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_249_v151).length));
          let _nw41 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          (_249_v151)[_index23] = _nw41;
          let _253_v155;
          let _init6 = ((_254_v3) => function (_255_i16) {
            return _254_v3;
          })(_51_v3);
          let _nw42 = Array((new BigNumber(28)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw42.length); _i0_6++) {
            _nw42[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _253_v155 = _nw42;
          let _256_v156;
          _256_v156 = _dafny.Set.fromElements(_253_v155);
          let _257_v157;
          let _nw43 = new _module.C4();
          _nw43.__ctor(_50_v2, new BigNumber(978), _256_v156, _51_v3);
          _257_v157 = _nw43;
          _257_v157 = _257_v157;
          let _258_v158;
          _258_v158 = _dafny.Map.Empty.slice().updateUnsafe((_257_v157).f18,(_257_v157).f18);
          let _259_v159;
          let _nw44 = new _module.C7();
          _nw44.__ctor((_250_v152)[_module.__default.safeIndex(new BigNumber(-841), new BigNumber((_250_v152).length))], _module.__default.fm45(_53_v5, new BigNumber((_258_v158).length), _51_v3, _56_globalState), _253_v155, _257_v157.f17, _module.__default.fm13(_56_globalState));
          _259_v159 = _nw44;
          _259_v159 = _259_v159;
        } else {
          let _260_v160;
          _260_v160 = _dafny.MultiSet.fromElements(_module.__default.fm30(true, _56_globalState));
          _260_v160 = _260_v160;
          (_56_globalState).f7 = _48_v0;
          let _261_v161;
          let _init7 = function (_262_i17) {
            return false;
          };
          let _nw45 = Array((_dafny.ONE).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw45.length); _i0_7++) {
            _nw45[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _261_v161 = _nw45;
          let _index24 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_261_v161).length));
          (_261_v161)[_index24] = _51_v3;
          let _263_v162;
          _263_v162 = _module.D0.create_DC0(_module.__default.fm13(_56_globalState));
          let _index25 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_261_v161).length));
          (_261_v161)[_index25] = (_module.__default.safeModuloInt(_50_v2, _50_v2)).isLessThan(((false) ? (new BigNumber((_dafny.Seq.update(_60_v10, _module.__default.safeIndex(_50_v2, new BigNumber((_60_v10).length)), _module.__default.fm12(_263_v162, _56_globalState))).length)) : ((_dafny.ZERO).minus(_50_v2))));
          let _264_v163;
          _264_v163 = _dafny.MultiSet.fromElements(_261_v161);
          let _index26 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_261_v161).length));
          (_261_v161)[_index26] = !(_264_v163).equals(_264_v163);
          (_56_globalState).f14 = _50_v2;
        }
        _60_v10 = ((true) ? (_dafny.Seq.Concat(_60_v10, _module.__default.fm10(_56_globalState))) : (_dafny.Seq.UnicodeFromString("udophjwub")));
        let _index27 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length));
        (_49_v1)[_index27] = (_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))];
        let _265_v164;
        _265_v164 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_51_v3);
        _51_v3 = ((_module.__default.fm13(_56_globalState)) ? (_51_v3) : ((_50_v2).isLessThan(new BigNumber((_265_v164).length))));
        _50_v2 = _50_v2;
      } else {
        let _266_v165;
        _266_v165 = _dafny.Map.Empty.slice().updateUnsafe(_50_v2,(new BigNumber(341)).isLessThanOrEqualTo(new BigNumber(((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]).length)));
        _266_v165 = (_266_v165).update(new BigNumber((_dafny.Seq.of(_51_v3)).length), !(_dafny.Seq.IsProperPrefixOf((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))], _dafny.Seq.UnicodeFromString("kpts"))));
        (_56_globalState).f14 = new BigNumber(((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]).length);
        (_56_globalState).f3 = (_51_v3) && (_51_v3);
        let _index28 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_54_v6).length));
        (_54_v6)[_index28] = _61_v11;
        let _index29 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_54_v6).length));
        (_54_v6)[_index29] = new _dafny.CodePoint('r'.codePointAt(0));
        let _267_v166;
        _267_v166 = _module.D2.create_DC6(new BigNumber(((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))]).length));
        let _268_v167;
        _268_v167 = _module.D2.create_DC7(_267_v166);
        let _source2 = _268_v167;
        if (_source2.is_DC6) {
          let _269___mcc_h7 = (_source2).cf13;
          let _270_cf13 = _269___mcc_h7;
          let _index30 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_48_v0).length));
          (_48_v0)[_index30] = _270_cf13;
          let _index31 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_48_v0).length));
          (_48_v0)[_index31] = _270_cf13;
          let _271_v168;
          _271_v168 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,_51_v3);
          let _272_v169;
          let _init8 = ((_273_v3) => function (_274_i18) {
            return _273_v3;
          })(_51_v3);
          let _nw46 = Array((new BigNumber(20)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw46.length); _i0_8++) {
            _nw46[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _272_v169 = _nw46;
          let _275_v170;
          _275_v170 = _dafny.Set.fromElements(_272_v169, _272_v169);
          let _276_v171;
          let _nw47 = new _module.C2();
          _nw47.__ctor(_275_v170, _51_v3);
          _276_v171 = _nw47;
          let _277_v172;
          _277_v172 = _dafny.Map.Empty.slice().updateUnsafe((((((_271_v168).contains(_51_v3)) ? ((_271_v168).get(_51_v3)) : (_51_v3))) ? (_276_v171) : (_276_v171)),_50_v2);
          _277_v172 = (_277_v172).update(_276_v171, (_48_v0)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_48_v0).length))]);
          let _index32 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_48_v0).length));
          (_48_v0)[_index32] = _270_cf13;
          let _278_v173;
          _278_v173 = _dafny.MultiSet.fromElements(!(!((((_266_v165).contains((_dafny.ZERO).minus(_270_cf13))) ? ((_266_v165).get((_dafny.ZERO).minus(_270_cf13))) : (_51_v3)))), true, (_276_v171).f18, (_276_v171).f18);
          _278_v173 = ((_278_v173).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_51_v3)))).Intersect(_278_v173);
        } else if (_source2.is_DC5) {
          let _279___mcc_h8 = (_source2).cf12;
          let _280_cf12 = _279___mcc_h8;
          _51_v3 = _51_v3;
          let _281_v174;
          _281_v174 = _module.D0.create_DC0(_51_v3);
          let _282_v175;
          let _init9 = ((_283_v3) => function (_284_i19) {
            return _283_v3;
          })(_51_v3);
          let _nw48 = Array((new BigNumber(12)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw48.length); _i0_9++) {
            _nw48[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _282_v175 = _nw48;
          let _285_v176;
          _285_v176 = _dafny.Set.fromElements(_282_v175, _282_v175);
          let _286_v177;
          _286_v177 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,_51_v3);
          let _287_v178;
          let _nw49 = new _module.C8();
          _nw49.__ctor(_module.__default.fm7(_50_v2, _dafny.Map.Empty.slice().updateUnsafe(!(_51_v3),new BigNumber(274)), function (_pat_let6_0) {
            return function (_288_dt__update__tmp_h3) {
              return function (_pat_let7_0) {
                return function (_289_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_289_dt__update_hcf0_h0);
                }(_pat_let7_0);
              }(true);
            }(_pat_let6_0);
          }(_281_v174), _56_globalState), _51_v3, _282_v175, _285_v176, !_dafny.Seq.contains(_60_v10, (_54_v6)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_54_v6).length))]), (((_59_v9).contains(_50_v2)) ? ((_59_v9).get(_50_v2)) : (_50_v2)), _286_v177);
          _287_v178 = _nw49;
          let _290_v179;
          let _nw50 = new _module.C5();
          _nw50.__ctor(_60_v10, (_287_v178).f24, _50_v2, _286_v177, _282_v175, _285_v176, _51_v3);
          _290_v179 = _nw50;
          _290_v179 = _290_v179;
          let _index33 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_282_v175).length));
          (_282_v175)[_index33] = (_287_v178).f24;
          let _index34 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_282_v175).length));
          (_282_v175)[_index34] = !((_290_v179).f30);
          let _291_v180;
          _291_v180 = _dafny.Map.Empty.slice().updateUnsafe((_290_v179).f30,(_287_v178).f24);
          let _292_v181;
          _292_v181 = _module.D0.create_DC1((_49_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_49_v1).length))], new BigNumber((_291_v180).length), _287_v178.f23, _61_v11);
          let _pat_let_tv3 = _54_v6;
          let _pat_let_tv4 = _54_v6;
          let _pat_let_tv5 = _50_v2;
          let _index35 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_282_v175).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_282_v175).length));
          let _rhs28 = (_287_v178).f24;
          let _rhs29 = (_50_v2).minus((function (_pat_let8_0) {
            return function (_293_dt__update__tmp_h4) {
              return function (_pat_let9_0) {
                return function (_294_dt__update_hcf4_h0) {
                  return function (_pat_let10_0) {
                    return function (_295_dt__update_hcf3_h0) {
                      return _module.D0.create_DC1((_293_dt__update__tmp_h4).dtor_cf1, (_293_dt__update__tmp_h4).dtor_cf2, _295_dt__update_hcf3_h0, _294_dt__update_hcf4_h0);
                    }(_pat_let10_0);
                  }(_pat_let_tv5);
                }(_pat_let9_0);
              }((_pat_let_tv4)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_pat_let_tv3).length))]);
            }(_pat_let8_0);
          }(_292_v181)).dtor_cf3);
          let _rhs30 = _287_v178.f23;
          let _rhs31 = !((_50_v2).isLessThan(_287_v178.f23));
          let _lhs20 = _282_v175;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_282_v175).length));
          let _lhs22 = _56_globalState;
          let _lhs23 = _282_v175;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_282_v175).length));
          _lhs20[_lhs21] = _rhs28;
          _50_v2 = _rhs29;
          _lhs22.f14 = _rhs30;
          _lhs23[_lhs24] = _rhs31;
        } else {
          let _296___mcc_h9 = (_source2).cf14;
          let _297_cf14 = _296___mcc_h9;
          let _298_v183;
          _298_v183 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,_50_v2);
          let _299_v184;
          let _init10 = ((_300_v10, _301_v3, _302_v1, _303_v2) => function (_304_i20) {
            return (_module.D3.create_DC10(_300_v10, _301_v3, (_302_v1)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_302_v1).length))], _dafny.Seq.UnicodeFromString("go"), _303_v2)).dtor_cf17;
          })(_60_v10, _51_v3, _49_v1, _50_v2);
          let _nw51 = Array((new BigNumber(20)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw51.length); _i0_10++) {
            _nw51[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _299_v184 = _nw51;
          let _305_v185;
          let _nw52 = new _module.C7();
          _nw52.__ctor(new BigNumber(900), function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of (_298_v183).Keys.Elements) {
              let _306_v182 = _compr_24;
              if ((_298_v183).contains(_306_v182)) {
                _coll24.push([_306_v182,false]);
              }
            }
            return _coll24;
          }(), _299_v184, _dafny.Set.fromElements(_299_v184, _299_v184, _299_v184, _299_v184, _299_v184), _51_v3);
          _305_v185 = _nw52;
          _305_v185 = _305_v185;
          let _307_v186;
          _307_v186 = _dafny.Map.Empty.slice().updateUnsafe((_52_v4)[_module.__default.safeIndex(_50_v2, new BigNumber((_52_v4).length))],_57_v7);
          let _308_v187;
          _308_v187 = _dafny.Map.Empty.slice().updateUnsafe((_54_v6)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_54_v6).length))],_61_v11);
          let _309_v188;
          _309_v188 = _dafny.Map.Empty.slice().updateUnsafe(_308_v187,_307_v186);
          let _310_v189;
          _310_v189 = _dafny.MultiSet.fromElements(_51_v3, _51_v3);
          let _311_v190;
          _311_v190 = _dafny.Set.fromElements(_299_v184);
          let _312_v191;
          let _init11 = ((_313_v9) => function (_314_i21) {
            return _313_v9;
          })(_59_v9);
          let _nw53 = Array((new BigNumber(10)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw53.length); _i0_11++) {
            _nw53[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _312_v191 = _nw53;
          let _315_v192;
          _315_v192 = _module.D7.create_DC17(_312_v191, _51_v3, _50_v2);
          let _316_v193;
          _316_v193 = _dafny.Seq.of(_315_v192);
          let _317_v194;
          _317_v194 = _module.D12.create_DC30(_310_v189, _51_v3, _module.D11.create_DC25(_311_v190), _316_v193, false);
          let _318_v195;
          let _nw54 = Array((new BigNumber(14)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _307_v186;
          _nw54[(_dafny.ONE).toNumber()] = (((_309_v188).contains(_308_v187)) ? ((_309_v188).get(_308_v187)) : (_307_v186));
          _nw54[(new BigNumber(2)).toNumber()] = ((_51_v3) ? (_307_v186) : (_module.__default.fm46(_56_globalState)));
          _nw54[(new BigNumber(3)).toNumber()] = (_307_v186).update((_317_v194).dtor_cf59, _57_v7);
          _nw54[(new BigNumber(4)).toNumber()] = (_307_v186).Merge(_307_v186);
          _nw54[(new BigNumber(5)).toNumber()] = _307_v186;
          _nw54[(new BigNumber(6)).toNumber()] = _307_v186;
          _nw54[(new BigNumber(7)).toNumber()] = _307_v186;
          _nw54[(new BigNumber(8)).toNumber()] = _307_v186;
          _nw54[(new BigNumber(9)).toNumber()] = (_307_v186).update(_51_v3, _57_v7);
          _nw54[(new BigNumber(10)).toNumber()] = _307_v186;
          _nw54[(new BigNumber(11)).toNumber()] = (_307_v186).Merge(_307_v186);
          _nw54[(new BigNumber(12)).toNumber()] = ((_51_v3) ? (_307_v186) : (_307_v186));
          _nw54[(new BigNumber(13)).toNumber()] = _307_v186;
          _318_v195 = _nw54;
          let _index37 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_318_v195).length));
          (_318_v195)[_index37] = _307_v186;
          let _index38 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_318_v195).length));
          (_318_v195)[_index38] = (_307_v186).Merge(_307_v186);
          let _index39 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_54_v6).length));
          (_54_v6)[_index39] = new _dafny.CodePoint('r'.codePointAt(0));
          let _319_v196;
          _319_v196 = _dafny.Seq.of(new BigNumber(796));
          let _320_v197;
          _320_v197 = _dafny.Seq.of(_319_v196, _dafny.Seq.of(_50_v2));
          (_56_globalState).f14 = (_dafny.ZERO).minus(_module.__default.fm7(_50_v2, _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(((_320_v197)[_module.__default.safeIndex(new BigNumber((_53_v5).cardinality()), new BigNumber((_320_v197).length))]).length)), _module.D0.create_DC0(!(_module.__default.fm13(_56_globalState))), _56_globalState));
        }
      }
      let _321_v198;
      let _init12 = ((_322_v3) => function (_323_i22) {
        return _322_v3;
      })(_51_v3);
      let _nw55 = Array((new BigNumber(12)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw55.length); _i0_12++) {
        _nw55[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _321_v198 = _nw55;
      let _324_v199;
      let _nw56 = new _module.C1();
      _nw56.__ctor();
      _324_v199 = _nw56;
      let _325_v200;
      _325_v200 = _dafny.Map.Empty.slice().updateUnsafe(_324_v199,_324_v199);
      let _326_v201;
      _326_v201 = _dafny.Seq.of(new BigNumber((_325_v200).length));
      let _327_v202;
      let _nw57 = new _module.C2();
      _nw57.__ctor(_dafny.Set.fromElements(_321_v198), !(!(_50_v2).isEqualTo(new BigNumber((_326_v201).length))));
      _327_v202 = _nw57;
      let _328_v203;
      _328_v203 = _module.D3.create_DC9();
      let _329_v204;
      _329_v204 = _dafny.Seq.of(_328_v203);
      let _330_v205;
      _330_v205 = _329_v204;
      let _331_i23;
      _331_i23 = _dafny.ZERO;
      L1: {
        let _pat_let_tv6 = _51_v3;
        while (function (_source3) {
          let _347___mcc_h10 = _source3;
          let _348_cf21 = _347___mcc_h10;
          return _pat_let_tv6;
        }(_330_v205)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_331_i23)) {
              break L1;
            }
            _331_i23 = (_331_i23).plus(_dafny.ONE);
            let _332_v206;
            _332_v206 = _dafny.Set.fromElements(_48_v0);
            let _333_v207;
            _333_v207 = _dafny.Map.Empty.slice().updateUnsafe(_51_v3,(_332_v206).IsProperSubsetOf(_332_v206));
            let _rhs32 = _321_v198;
            let _rhs33 = _50_v2;
            let _rhs34 = ((_333_v207).Merge(_333_v207)).update((_51_v3) && (false), _51_v3);
            let _lhs25 = _56_globalState;
            _321_v198 = _rhs32;
            _lhs25.f14 = _rhs33;
            _333_v207 = _rhs34;
            let _334_v208;
            let _335_v209;
            let _336_v210;
            let _337_v211;
            let _out13;
            let _out14;
            let _out15;
            let _out16;
            let _outcollector4 = (_324_v199).m5((_dafny.ZERO).minus(_50_v2), _56_globalState);
            _out13 = _outcollector4[0];
            _out14 = _outcollector4[1];
            _out15 = _outcollector4[2];
            _out16 = _outcollector4[3];
            _334_v208 = _out13;
            _335_v209 = _out14;
            _336_v210 = _out15;
            _337_v211 = _out16;
            let _338_v212;
            _338_v212 = _dafny.MultiSet.fromElements(_336_v210);
            _338_v212 = _338_v212;
            let _339_v215;
            _339_v215 = _module.D10.create_DC22(function () {
  let _coll25 = new _dafny.Set();
  for (const _compr_25 of (_326_v201).Elements) {
    let _340_v214 = _compr_25;
    if (_dafny.Seq.contains(_326_v201, _340_v214)) {
      _coll25.add(_module.__default.safeModuloInt(_340_v214, new BigNumber(238)));
    }
  }
  return _coll25;
}());
            let _341_v216;
            _341_v216 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_59_v9).length),_339_v215);
            let _342_v217;
            let _343_v218;
            let _344_v219;
            let _345_v220;
            let _out17;
            let _out18;
            let _out19;
            let _out20;
            let _outcollector5 = (_324_v199).m5(new BigNumber(((function () {
              let _coll26 = new _dafny.Map();
              for (const _compr_26 of _dafny.IntegerRange(new BigNumber(-608), new BigNumber(634))) {
                let _346_v213 = _compr_26;
                if (((new BigNumber(-608)).isLessThanOrEqualTo(_346_v213)) && ((_346_v213).isLessThan(new BigNumber(634)))) {
                  _coll26.push([_module.__default.safeModuloInt(_346_v213, _50_v2),_module.D10.create_DC22(_57_v7)]);
                }
              }
              return _coll26;
            }()).Merge(_341_v216)).length), _56_globalState);
            _out17 = _outcollector5[0];
            _out18 = _outcollector5[1];
            _out19 = _outcollector5[2];
            _out20 = _outcollector5[3];
            _342_v217 = _out17;
            _343_v218 = _out18;
            _344_v219 = _out19;
            _345_v220 = _out20;
          }
        }
      }
      let _349_v221;
      let _350_v222;
      let _351_v223;
      let _352_v224;
      let _out21;
      let _out22;
      let _out23;
      let _out24;
      let _outcollector6 = (_324_v199).m5((_50_v2).minus(_50_v2), _56_globalState);
      _out21 = _outcollector6[0];
      _out22 = _outcollector6[1];
      _out23 = _outcollector6[2];
      _out24 = _outcollector6[3];
      _349_v221 = _out21;
      _350_v222 = _out22;
      _351_v223 = _out23;
      _352_v224 = _out24;
      let _353_v225;
      let _354_v226;
      let _355_v227;
      let _356_v228;
      let _out25;
      let _out26;
      let _out27;
      let _out28;
      let _outcollector7 = (_324_v199).m5(new BigNumber(832), _56_globalState);
      _out25 = _outcollector7[0];
      _out26 = _outcollector7[1];
      _out27 = _outcollector7[2];
      _out28 = _outcollector7[3];
      _353_v225 = _out25;
      _354_v226 = _out26;
      _355_v227 = _out27;
      _356_v228 = _out28;
      _351_v223 = (((_50_v2).isLessThanOrEqualTo((((_59_v9).contains(_50_v2)) ? ((_59_v9).get(_50_v2)) : (_50_v2)))) ? (_352_v224) : (!(_356_v228)));
      (_56_globalState).f3 = _351_v223;
      let _357_v229;
      let _358_v230;
      let _359_v231;
      let _360_v232;
      let _out29;
      let _out30;
      let _out31;
      let _out32;
      let _outcollector8 = (_324_v199).m5((new BigNumber((_60_v10).length)).multipliedBy(_50_v2), _56_globalState);
      _out29 = _outcollector8[0];
      _out30 = _outcollector8[1];
      _out31 = _outcollector8[2];
      _out32 = _outcollector8[3];
      _357_v229 = _out29;
      _358_v230 = _out30;
      _359_v231 = _out31;
      _360_v232 = _out32;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_321_v198).length))) {
        let _361_i24 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_361_i24)) && ((_361_i24).isLessThan(new BigNumber((_321_v198).length))))) {
          (_321_v198)[(_361_i24)] = ((_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("vws"), _61_v11)) ? (!(!(_359_v231))) : (_359_v231));
        }
      }
      let _362_v233;
      _362_v233 = _dafny.MultiSet.fromElements(_359_v231, _51_v3);
      let _363_i25;
      _363_i25 = _dafny.ZERO;
      L2: {
        while (((((_362_v233).contains(_352_v224)) ? ((_362_v233).get(_352_v224)) : (new BigNumber((_52_v4).length)))).isEqualTo(_50_v2)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_363_i25)) {
              break L2;
            }
            _363_i25 = (_363_i25).plus(_dafny.ONE);
            let _364_v234;
            _364_v234 = _dafny.Set.fromElements(_321_v198);
            let _365_v235;
            _365_v235 = _dafny.Map.Empty.slice().updateUnsafe(_61_v11,false);
            let _366_v236;
            let _nw58 = new _module.C8();
            _nw58.__ctor(new BigNumber((_60_v10).length), _355_v227, _321_v198, _364_v234, _351_v223, _50_v2, _365_v235);
            _366_v236 = _nw58;
            let _367_v237;
            let _nw59 = Array((new BigNumber(3)).toNumber());
            _nw59[(_dafny.ZERO).toNumber()] = _366_v236;
            _nw59[(_dafny.ONE).toNumber()] = _366_v236;
            _nw59[(new BigNumber(2)).toNumber()] = _366_v236;
            _367_v237 = _nw59;
            let _index40 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_367_v237).length));
            (_367_v237)[_index40] = _366_v236;
            let _index41 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_321_v198).length));
            (_321_v198)[_index41] = _51_v3;
            let _index42 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_367_v237).length));
            let _index43 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_321_v198).length));
            let _rhs35 = _366_v236;
            let _rhs36 = (_366_v236).f20;
            let _rhs37 = _355_v227;
            let _rhs38 = !(_352_v224);
            let _rhs39 = _module.__default.safeModuloInt(new BigNumber(915), _50_v2);
            let _lhs26 = _367_v237;
            let _lhs27 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_367_v237).length));
            let _lhs28 = _56_globalState;
            let _lhs29 = _321_v198;
            let _lhs30 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_321_v198).length));
            let _lhs31 = _56_globalState;
            _lhs26[_lhs27] = _rhs35;
            _lhs28.f14 = _rhs36;
            _lhs29[_lhs30] = _rhs37;
            _lhs31.f3 = _rhs38;
            _50_v2 = _rhs39;
            let _368_v238;
            let _nw60 = new _module.C3();
            _nw60.__ctor(_366_v236.f19, (_364_v234).Union(_366_v236.f17), _351_v223);
            _368_v238 = _nw60;
            let _369_v239;
            let _nw61 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
            _369_v239 = _nw61;
            let _370_v240;
            _370_v240 = _dafny.Set.fromElements(_48_v0);
            let _index44 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_369_v239).length));
            (_369_v239)[_index44] = _370_v240;
            let _371_v241;
            let _nw62 = new _module.C3();
            _nw62.__ctor(_366_v236.f19, _366_v236.f17, (_366_v236).f18);
            _371_v241 = _nw62;
            let _372_v242;
            _372_v242 = _dafny.Map.Empty.slice().updateUnsafe(_50_v2,_371_v241);
            let _index45 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_369_v239).length));
            let _rhs40 = ((_370_v240).Intersect(_370_v240)).Difference(_370_v240);
            let _rhs41 = new BigNumber((_372_v242).length);
            let _lhs32 = _369_v239;
            let _lhs33 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_369_v239).length));
            _lhs32[_lhs33] = _rhs40;
            _50_v2 = _rhs41;
            let _373_v243;
            _373_v243 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC5(_362_v233),true);
            let _374_v244;
            _374_v244 = _module.D2.create_DC5(_362_v233);
            let _375_v246;
            _375_v246 = _dafny.Seq.of(_374_v244);
            let _376_v248;
            _376_v248 = _dafny.Seq.of(_373_v243);
            let _377_v252;
            _377_v252 = _dafny.MultiSet.fromElements(_374_v244);
            let _378_v253;
            let _nw63 = Array((new BigNumber(24)).toNumber());
            _nw63[(_dafny.ZERO).toNumber()] = _373_v243;
            _nw63[(_dafny.ONE).toNumber()] = _module.__default.fm47(_356_v228, _56_globalState);
            _nw63[(new BigNumber(2)).toNumber()] = (_module.__default.fm47(_51_v3, _56_globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_374_v244,_355_v227));
            _nw63[(new BigNumber(3)).toNumber()] = (((_366_v236).f18) ? (_373_v243) : (_dafny.Map.Empty.slice().updateUnsafe(_374_v244,false)));
            _nw63[(new BigNumber(4)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(5)).toNumber()] = (_373_v243).Merge(_373_v243);
            _nw63[(new BigNumber(6)).toNumber()] = function () {
              let _coll27 = new _dafny.Map();
              for (const _compr_27 of (_dafny.Seq.update(_375_v246, _module.__default.safeIndex((_366_v236).f20, new BigNumber((_375_v246).length)), _module.D2.create_DC5(_dafny.MultiSet.fromElements(!(_51_v3))))).Elements) {
                let _379_v245 = _compr_27;
                if (_dafny.Seq.contains(_dafny.Seq.update(_375_v246, _module.__default.safeIndex((_366_v236).f20, new BigNumber((_375_v246).length)), _module.D2.create_DC5(_dafny.MultiSet.fromElements(!(_51_v3)))), _379_v245)) {
                  _coll27.push([_379_v245,_352_v224]);
                }
              }
              return _coll27;
            }();
            _nw63[(new BigNumber(7)).toNumber()] = (function () {
              let _coll28 = new _dafny.Map();
              for (const _compr_28 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), ((_380_v244) => function (_381_i26) {
                return _380_v244;
              })(_374_v244))).Elements) {
                let _382_v247 = _compr_28;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), ((_383_v244) => function (_381_i26) {
                  return _383_v244;
                })(_374_v244)), _382_v247)) {
                  _coll28.push([_382_v247,(_366_v236).f18]);
                }
              }
              return _coll28;
            }()).update(_374_v244, _51_v3);
            _nw63[(new BigNumber(8)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(9)).toNumber()] = (_373_v243).Merge(_373_v243);
            _nw63[(new BigNumber(10)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_374_v244,_51_v3);
            _nw63[(new BigNumber(12)).toNumber()] = (_373_v243).Merge((_376_v248)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_376_v248).length))]);
            _nw63[(new BigNumber(13)).toNumber()] = function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of (_375_v246).Elements) {
                let _384_v249 = _compr_29;
                if (_dafny.Seq.contains(_375_v246, _384_v249)) {
                  _coll29.push([_384_v249,(_module.D5.create_DC14(false, new BigNumber((_dafny.Set.fromElements((_366_v236).f20, (_366_v236).f20)).length), (_366_v236).f20)).dtor_cf26]);
                }
              }
              return _coll29;
            }();
            _nw63[(new BigNumber(14)).toNumber()] = function () {
              let _coll30 = new _dafny.Map();
              for (const _compr_30 of (_dafny.Seq.update(_375_v246, _module.__default.safeIndex(_50_v2, new BigNumber((_375_v246).length)), _374_v244)).Elements) {
                let _385_v250 = _compr_30;
                if (_dafny.Seq.contains(_dafny.Seq.update(_375_v246, _module.__default.safeIndex(_50_v2, new BigNumber((_375_v246).length)), _374_v244), _385_v250)) {
                  _coll30.push([_385_v250,_360_v232]);
                }
              }
              return _coll30;
            }();
            _nw63[(new BigNumber(15)).toNumber()] = (_373_v243).Merge(_373_v243);
            _nw63[(new BigNumber(16)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(17)).toNumber()] = _module.__default.fm47(_51_v3, _56_globalState);
            _nw63[(new BigNumber(18)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(19)).toNumber()] = function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of (_377_v252).Elements) {
                let _386_v251 = _compr_31;
                if ((_377_v252).contains(_386_v251)) {
                  _coll31.push([_386_v251,(_366_v236).f18]);
                }
              }
              return _coll31;
            }();
            _nw63[(new BigNumber(20)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(21)).toNumber()] = _module.__default.fm47(_352_v224, _56_globalState);
            _nw63[(new BigNumber(22)).toNumber()] = _373_v243;
            _nw63[(new BigNumber(23)).toNumber()] = _373_v243;
            _378_v253 = _nw63;
            let _index46 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_378_v253).length));
            (_378_v253)[_index46] = function () {
              let _coll32 = new _dafny.Map();
              for (const _compr_32 of (_377_v252).Elements) {
                let _387_v254 = _compr_32;
                if ((_377_v252).contains(_387_v254)) {
                  _coll32.push([_387_v254,(_371_v241).f18]);
                }
              }
              return _coll32;
            }();
            let _388_v255;
            let _init13 = ((_389_v201) => function (_390_i27) {
              return _389_v201;
            })(_326_v201);
            let _nw64 = Array((new BigNumber(25)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw64.length); _i0_13++) {
              _nw64[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _388_v255 = _nw64;
            let _index47 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_388_v255).length));
            (_388_v255)[_index47] = _dafny.Seq.Concat(_326_v201, _326_v201);
            let _391_v256;
            _391_v256 = _dafny.Map.Empty.slice().updateUnsafe((_366_v236).f20,_359_v231);
            let _index48 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_378_v253).length));
            let _index49 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_388_v255).length));
            let _rhs42 = _349_v221;
            let _rhs43 = (_366_v236).f20;
            let _rhs44 = (((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_374_v244,(((_391_v256).contains(new BigNumber(-278))) ? ((_391_v256).get(new BigNumber(-278))) : (_355_v227)))) : (_module.__default.fm47(false, _56_globalState)))).Merge(_373_v243);
            let _rhs45 = _dafny.Seq.update(_326_v201, _module.__default.safeIndex((_366_v236).f20, new BigNumber((_326_v201).length)), (_dafny.ZERO).minus(_50_v2));
            let _lhs34 = _56_globalState;
            let _lhs35 = _378_v253;
            let _lhs36 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_378_v253).length));
            let _lhs37 = _388_v255;
            let _lhs38 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_388_v255).length));
            _353_v225 = _rhs42;
            _lhs34.f14 = _rhs43;
            _lhs35[_lhs36] = _rhs44;
            _lhs37[_lhs38] = _rhs45;
          }
        }
      }
      (_56_globalState).f3 = _360_v232;
      process.stdout.write(((_49_v1)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_50_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_51_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_52_v4, _dafny.Seq.of(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_v5).equals(_dafny.MultiSet.fromElements(new BigNumber(-548), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_56_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_56_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_56_globalState).f13)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_56_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f15).equals(_dafny.MultiSet.fromElements(_dafny.ZERO, _dafny.ZERO, new BigNumber(3288), new BigNumber(3288)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_globalState.f16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v7).equals(_dafny.Set.fromElements(new BigNumber(548)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_58_v8, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(548)), _dafny.Set.fromElements(new BigNumber(548)), _dafny.Set.fromElements(new BigNumber(548))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_60_v10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_61_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_62_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.ONE),_dafny.Seq.UnicodeFromString("yvxdaq")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_321_v198)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_325_v200).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_326_v201, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_329_v204, _dafny.Seq.of(_module.D3.create_DC9()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_330_v205), _dafny.Seq.of(_module.D3.create_DC9()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_331_i23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_349_v221).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_350_v222).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v222).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v222).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v222).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_351_v223));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_352_v224));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_353_v225).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_354_v226).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v226).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v226).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v226).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_355_v227));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_356_v228));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_357_v229).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_358_v230).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_358_v230).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_358_v230).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_358_v230).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_359_v231));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_360_v232));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_362_v233).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_363_i25));
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
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5, cf6, cf7) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + this.cf1.toVerbatimString(true) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(false);
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
    static create_DC4(cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC6(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC5(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf14) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO);
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
    static create_DC9() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC10(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D3(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + this.cf16.toVerbatimString(true) + ", " + _dafny.toString(this.cf17) + ", " + this.cf18.toVerbatimString(true) + ", " + this.cf19.toVerbatimString(true) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9();
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
    static create_DC11(cf21) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21);
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
    static create_DC13(cf23, cf24, cf25) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf26, cf27, cf28) {
      let $dt = new D5(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC12(cf22) {
      let $dt = new D5(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(false, _dafny.ZERO, _dafny.Seq.of());
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
    static create_DC15(cf29) {
      let $dt = new D6(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf31, cf32, cf33) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf30) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf31 === other.cf31 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17([], false, _dafny.ZERO);
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
    static create_DC18(cf34) {
      let $dt = new D8(0);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf36, cf37, cf38, cf39, cf40) {
      let $dt = new D9(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC19(cf35) {
      let $dt = new D9(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC21(cf41) {
      let $dt = new D9(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20(_dafny.Map.Empty, _dafny.ZERO, false, _dafny.Map.Empty, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC23(cf43, cf44, cf45) {
      let $dt = new D10(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC22(cf42) {
      let $dt = new D10(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24(cf46) {
      let $dt = new D10(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43 && this.cf44 === other.cf44 && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC23(false, false, null);
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
    static create_DC26(cf48, cf49) {
      let $dt = new D11(0);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC27(cf50, cf51, cf52) {
      let $dt = new D11(1);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC25(cf47) {
      let $dt = new D11(2);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC28(cf53) {
      let $dt = new D11(3);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC26(false, _dafny.ZERO);
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
    static create_DC30(cf55, cf56, cf57, cf58, cf59) {
      let $dt = new D12(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC29(cf54) {
      let $dt = new D12(1);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56 && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58) && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC30(_dafny.MultiSet.Empty, false, _module.D11.Default(), _dafny.Seq.of(), false);
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
    static create_DC32(cf61, cf62) {
      let $dt = new D13(0);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC31(cf60) {
      let $dt = new D13(1);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC33(cf63) {
      let $dt = new D13(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC32(false, _dafny.ZERO);
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
    static create_DC35() {
      let $dt = new D14(0);
      return $dt;
    }
    static create_DC34(cf64) {
      let $dt = new D14(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC36(cf65) {
      let $dt = new D14(2);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC35";
      } else if (this.$tag === 1) {
        return "D14.DC34" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf65) + ")";
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
        return other.$tag === 1 && this.cf64 === other.cf64;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC35();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = false;
      this.f7 = [];
      this.f9 = false;
      this.f14 = _dafny.ZERO;
      this.f15 = _dafny.MultiSet.Empty;
      this.f16 = [];
      this._f0 = false;
      this._f1 = false;
      this._f2 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f10 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f11 = false;
      this._f12 = _dafny.ZERO;
      this._f13 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
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
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(583)).isLessThan(new BigNumber(122));
    };
    m5(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D0.Default();
      let r2 = false;
      let r3 = false;
      let _392_v0;
      _392_v0 = new _dafny.CodePoint('n'.codePointAt(0));
      let _393_v2;
      _393_v2 = _dafny.Map.Empty.slice().updateUnsafe(_392_v0,p0);
      let _394_v4;
      _394_v4 = _dafny.Seq.of((function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_393_v2).Keys.Elements) {
          let _395_v1 = _compr_33;
          if ((_393_v2).contains(_395_v1)) {
            _coll33.push([_395_v1,new BigNumber((function () {
              let _coll34 = new _dafny.Map();
              for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-472), new BigNumber(765))) {
                let _396_v3 = _compr_34;
                if (((new BigNumber(-472)).isLessThanOrEqualTo(_396_v3)) && ((_396_v3).isLessThan(new BigNumber(765)))) {
                  _coll34.push([(_396_v3).minus(p0),true]);
                }
              }
              return _coll34;
            }()).length)]);
          }
        }
        return _coll33;
      }()).contains((_module.D0.create_DC1(_dafny.Seq.UnicodeFromString("ecidbhho"), p0, (_dafny.ZERO).minus(p0), _392_v0)).dtor_cf4));
      let _397_i0;
      _397_i0 = _dafny.ZERO;
      L3: {
        while ((_394_v4)[_module.__default.safeIndex(p0, new BigNumber((_394_v4).length))]) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_397_i0)) {
              break L3;
            }
            _397_i0 = (_397_i0).plus(_dafny.ONE);
            let _398_v5;
            let _nw65 = Array((new BigNumber(15)).toNumber()).fill(false);
            _398_v5 = _nw65;
            let _399_v6;
            _399_v6 = true;
            let _index50 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length));
            (_398_v5)[_index50] = _399_v6;
            let _index51 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length));
            (_398_v5)[_index51] = _399_v6;
            if ((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))]) {
              let _index52 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length));
              let _rhs46 = p0;
              let _rhs47 = (_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))];
              let _rhs48 = _399_v6;
              let _rhs49 = _398_v5;
              let _lhs39 = globalState;
              let _lhs40 = _398_v5;
              let _lhs41 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length));
              _lhs39.f14 = _rhs46;
              r2 = _rhs47;
              _lhs40[_lhs41] = _rhs48;
              _398_v5 = _rhs49;
              let _400_v7;
              _400_v7 = _dafny.Seq.of(p0);
              (globalState).f9 = (_this).fm5(p0, _dafny.Seq.Concat(_400_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_401_i1) {
                return new BigNumber(-71);
              })), globalState);
              let _402_v8;
              _402_v8 = _dafny.MultiSet.fromElements(p0);
              let _403_v9;
              _403_v9 = _module.D1.create_DC3(_402_v8);
              (globalState).f14 = new BigNumber(((_403_v9).dtor_cf8).cardinality());
              (globalState).f14 = p0;
              _399_v6 = _399_v6;
            } else {
              let _404_v10;
              _404_v10 = _dafny.Seq.UnicodeFromString("lylpu");
              let _405_v11;
              _405_v11 = _dafny.Map.Empty.slice().updateUnsafe(_404_v10,_404_v10);
              let _406_v12;
              _406_v12 = _dafny.MultiSet.fromElements(p0, p0);
              let _407_v13;
              let _nw66 = Array((new BigNumber(2)).toNumber());
              _nw66[(_dafny.ZERO).toNumber()] = (((_406_v12).contains(new BigNumber((_dafny.Seq.UnicodeFromString("kx")).length))) ? ((_406_v12).get(new BigNumber((_dafny.Seq.UnicodeFromString("kx")).length))) : (p0));
              _nw66[(_dafny.ONE).toNumber()] = p0;
              _407_v13 = _nw66;
              let _408_v14;
              _408_v14 = _module.D0.create_DC2((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))], new BigNumber((_405_v11).length), _407_v13);
              let _pat_let_tv7 = p0;
              let _409_v15;
              _409_v15 = _dafny.MultiSet.fromElements(function (_pat_let11_0) {
                return function (_410_dt__update__tmp_h0) {
                  return function (_pat_let12_0) {
                    return function (_411_dt__update_hcf6_h0) {
                      return _module.D0.create_DC2((_410_dt__update__tmp_h0).dtor_cf5, _411_dt__update_hcf6_h0, (_410_dt__update__tmp_h0).dtor_cf7);
                    }(_pat_let12_0);
                  }(_pat_let_tv7);
                }(_pat_let11_0);
              }(_408_v14));
              _409_v15 = (((_399_v6) || ((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))])) ? ((_dafny.MultiSet.fromElements(_408_v14, _module.D0.create_DC2((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))], p0, _407_v13), _408_v14)).Difference(_dafny.MultiSet.fromElements(_module.D0.create_DC2((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))], (_dafny.ZERO).minus(p0), _407_v13), _408_v14))) : (_409_v15));
              let _412_v16;
              let _nw67 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
              _412_v16 = _nw67;
              let _413_v17;
              _413_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(140),(_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))]);
              let _index53 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_412_v16).length));
              (_412_v16)[_index53] = _413_v17;
              let _414_v18;
              _414_v18 = _dafny.Seq.of(_404_v10);
              let _index54 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_412_v16).length));
              (_412_v16)[_index54] = (_413_v17).update(((_399_v6) ? (new BigNumber(169)) : (new BigNumber((_414_v18).length))), false);
              let _415_v19;
              _415_v19 = _dafny.Seq.of(p0, p0);
              let _416_v20;
              _416_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_415_v19).length),p0),(_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))]);
              (globalState).f14 = new BigNumber(((_module.__default.fm6((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))], (((_413_v17).contains(p0)) ? ((_413_v17).get(p0)) : (_399_v6)), p0, globalState)).Merge(_416_v20)).length);
              let _rhs50 = _404_v10;
              let _rhs51 = _399_v6;
              let _lhs42 = globalState;
              r0 = _rhs50;
              _lhs42.f3 = _rhs51;
              let _417_v21;
              _417_v21 = _dafny.Map.Empty.slice().updateUnsafe(false,(p0).isLessThanOrEqualTo(p0));
              _417_v21 = (_417_v21).update(_399_v6, (_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))]);
            }
            _398_v5 = (((_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))]) ? (_398_v5) : (_398_v5));
            let _418_v22;
            _418_v22 = _dafny.Set.fromElements(_399_v6);
            _418_v22 = (_418_v22).Union((_dafny.Set.fromElements(_399_v6, (_398_v5)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_398_v5).length))])).Intersect(_418_v22));
          }
        }
      }
      let _419_v23;
      _419_v23 = _dafny.Seq.UnicodeFromString("cmpkw");
      let _420_v24;
      _420_v24 = _dafny.Seq.of(p0);
      let _421_v25;
      _421_v25 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber((_419_v23).length), new BigNumber((_420_v24).length)), p0);
      let _422_v26;
      _422_v26 = false;
      let _423_v27;
      _423_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,_422_v26),new BigNumber((_421_v25).cardinality()));
      let _424_v28;
      _424_v28 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,_422_v26);
      let _425_v29;
      _425_v29 = _dafny.MultiSet.fromElements(_424_v28, _424_v28);
      (globalState).f14 = (((_421_v25).contains((new BigNumber((_423_v27).length)).plus((((_425_v29).contains(_424_v28)) ? ((_425_v29).get(_424_v28)) : (p0))))) ? ((_421_v25).get((new BigNumber((_423_v27).length)).plus((((_425_v29).contains(_424_v28)) ? ((_425_v29).get(_424_v28)) : (p0))))) : (new BigNumber(162)));
      let _hi1 = (_dafny.ZERO).minus(p0);
      for (let _426_i2 = p0; _426_i2.isLessThan(_hi1); _426_i2 = _426_i2.plus(_dafny.ONE)) {
        let _427_v30;
        _427_v30 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,_426_i2);
        let _428_v31;
        _428_v31 = _module.D0.create_DC0(_422_v26);
        let _429_v32;
        _429_v32 = _dafny.Set.fromElements(_422_v26, _422_v26);
        let _430_v33;
        _430_v33 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("fegnc"), _module.__default.fm7(_module.__default.fm7(_426_i2, _dafny.Map.Empty.slice().updateUnsafe(_422_v26,_426_i2), _428_v31, globalState), _427_v30, _428_v31, globalState), _426_i2, _392_v0);
        let _431_v34;
        _431_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,_392_v0);
        let _432_v35;
        let _init14 = ((_433_i2) => function (_434_i4) {
          return _module.__default.safeModuloInt(_434_i4, _433_i2);
        })(_426_i2);
        let _nw68 = Array((new BigNumber(28)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw68.length); _i0_14++) {
          _nw68[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _432_v35 = _nw68;
        let _435_v36;
        _435_v36 = _module.D0.create_DC2(_422_v26, _426_i2, _432_v35);
        let _pat_let_tv8 = _422_v26;
        let _436_v37;
        let _nw69 = Array((new BigNumber(25)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = (_426_i2).plus(_426_i2);
        _nw69[(_dafny.ONE).toNumber()] = _module.__default.fm7(_426_i2, _427_v30, function (_pat_let13_0) {
          return function (_437_dt__update__tmp_h1) {
            return function (_pat_let14_0) {
              return function (_438_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_438_dt__update_hcf0_h0);
              }(_pat_let14_0);
            }(_pat_let_tv8);
          }(_pat_let13_0);
        }(_428_v31), globalState);
        _nw69[(new BigNumber(2)).toNumber()] = p0;
        _nw69[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm5(new BigNumber((_429_v32).length), _420_v24, globalState),(_this).fm5(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-495)), ((_439_v23) => function (_440_i3) {
          return _dafny.MultiSet.fromElements(new BigNumber((_439_v23).length));
        })(_419_v23)))).cardinality()), _420_v24, globalState))).length), _426_i2);
        _nw69[(new BigNumber(4)).toNumber()] = new BigNumber((_424_v28).length);
        _nw69[(new BigNumber(5)).toNumber()] = new BigNumber((_427_v30).length);
        _nw69[(new BigNumber(6)).toNumber()] = (p0).minus(p0);
        _nw69[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(669), _426_i2);
        _nw69[(new BigNumber(8)).toNumber()] = p0;
        _nw69[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(((_422_v26) ? (p0) : (p0)));
        _nw69[(new BigNumber(10)).toNumber()] = (_430_v33).dtor_cf2;
        _nw69[(new BigNumber(11)).toNumber()] = (new BigNumber((_module.__default.fm8(_422_v26, _422_v26, (((_431_v34).contains(new BigNumber((_419_v23).length))) ? ((_431_v34).get(new BigNumber((_419_v23).length))) : (_392_v0)), globalState)).length)).multipliedBy(new BigNumber((_394_v4).length));
        _nw69[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_435_v36).dtor_cf6);
        _nw69[(new BigNumber(13)).toNumber()] = _426_i2;
        _nw69[(new BigNumber(14)).toNumber()] = new BigNumber(193);
        _nw69[(new BigNumber(15)).toNumber()] = _426_i2;
        _nw69[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw69[(new BigNumber(17)).toNumber()] = _426_i2;
        _nw69[(new BigNumber(18)).toNumber()] = (_426_i2).multipliedBy(p0);
        _nw69[(new BigNumber(19)).toNumber()] = new BigNumber(552);
        _nw69[(new BigNumber(20)).toNumber()] = p0;
        _nw69[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber((_419_v23).length));
        _nw69[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_394_v4, _394_v4)).length);
        _nw69[(new BigNumber(23)).toNumber()] = _426_i2;
        _nw69[(new BigNumber(24)).toNumber()] = p0;
        _436_v37 = _nw69;
        let _441_v38;
        _441_v38 = _dafny.MultiSet.fromElements(_422_v26);
        let _442_v39;
        _442_v39 = _module.D2.create_DC5(_441_v38);
        let _index55 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_436_v37).length));
        (_436_v37)[_index55] = _module.__default.fm7(_module.__default.fm7(new BigNumber((_427_v30).length), _427_v30, _428_v31, globalState), _dafny.Map.Empty.slice().updateUnsafe(_422_v26,new BigNumber(((_442_v39).dtor_cf12).cardinality())), _428_v31, globalState);
        let _443_v40;
        _443_v40 = _dafny.Set.fromElements(p0, p0);
        let _444_v41;
        _444_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_443_v40).length),new BigNumber((_module.__default.fm9(_392_v0, globalState)).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_436_v37).length));
        (_436_v37)[_index56] = (_dafny.ZERO).minus((_426_i2).minus((_dafny.ZERO).minus((((_444_v41).contains(new BigNumber(813))) ? ((_444_v41).get(new BigNumber(813))) : (new BigNumber(394))))));
        r3 = !((_this).fm5((_436_v37)[_module.__default.safeIndex(new BigNumber(750), new BigNumber((_436_v37).length))], _dafny.Seq.of(_426_i2), globalState));
        (globalState).f14 = (_dafny.ZERO).minus(p0);
        let _445_v42;
        let _nw70 = new _module.C0();
        _nw70.__ctor();
        _445_v42 = _nw70;
      }
      let _446_v43;
      _446_v43 = _dafny.Seq.of(_dafny.Seq.Concat(_module.__default.fm10(globalState), _419_v23));
      let _rhs52 = _422_v26;
      let _rhs53 = _module.__default.safeDivisionInt((_420_v24)[_module.__default.safeIndex(p0, new BigNumber((_420_v24).length))], p0);
      let _rhs54 = ((_421_v25).Intersect((_421_v25).update(new BigNumber((_dafny.Seq.UnicodeFromString("ug")).length), _module.__default.abs(p0)))).Union(_421_v25);
      let _rhs55 = _446_v43;
      let _lhs43 = globalState;
      let _lhs44 = globalState;
      let _lhs45 = globalState;
      _lhs43.f9 = _rhs52;
      _lhs44.f14 = _rhs53;
      _lhs45.f15 = _rhs54;
      _446_v43 = _rhs55;
      let _447_v44;
      let _nw71 = Array((new BigNumber(2)).toNumber()).fill(false);
      _447_v44 = _nw71;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_447_v44).length))) {
        let _448_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_448_i5)) && ((_448_i5).isLessThan(new BigNumber((_447_v44).length))))) {
          (_447_v44)[(_448_i5)] = ((p0).minus((_dafny.ZERO).minus(p0))).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(_422_v26, _422_v26, _422_v26, _422_v26)).length));
        }
      }
      if ((_394_v4)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_394_v4).length))]) {
        _422_v26 = (_421_v25).IsSubsetOf((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs(new BigNumber((_394_v4).length))));
        if (!(_dafny.Seq.contains(_dafny.Seq.Concat(_394_v4, _394_v4), _422_v26))) {
          let _449_v45;
          _449_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_420_v24).length),_422_v26);
          let _450_v46;
          _450_v46 = _module.D0.create_DC0((((_449_v45).contains(p0)) ? ((_449_v45).get(p0)) : ((_this).fm5(p0, _420_v24, globalState))));
          let _451_v47;
          _451_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm7(p0, _dafny.Map.Empty.slice().updateUnsafe(_422_v26,p0), _450_v46, globalState));
          _451_v47 = _451_v47;
          (globalState).f3 = _422_v26;
          let _452_v48;
          _452_v48 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,p0);
          (globalState).f14 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(p0, new BigNumber((_452_v48).length)), p0);
          let _453_v49;
          let _nw72 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _453_v49 = _nw72;
          let _index57 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_453_v49).length));
          (_453_v49)[_index57] = p0;
          let _index58 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_453_v49).length));
          (_453_v49)[_index58] = p0;
          let _index59 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_447_v44).length));
          (_447_v44)[_index59] = false;
          let _index60 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_447_v44).length));
          (_447_v44)[_index60] = _422_v26;
        } else {
          (globalState).f3 = _422_v26;
          let _454_v50;
          _454_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,!(_422_v26)),p0));
          let _455_v51;
          _455_v51 = _module.D0.create_DC0(_422_v26);
          let _456_v53;
          _456_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,_422_v26);
          let _457_v54;
          _457_v54 = _dafny.Seq.of(_456_v53);
          _423_v27 = ((((_454_v50).contains(p0)) ? ((_454_v50).get(p0)) : (_423_v27))).Merge((((_455_v51).dtor_cf0) ? (function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_dafny.Seq.update(_457_v54, _module.__default.safeIndex(new BigNumber(-468), new BigNumber((_457_v54).length)), _456_v53)).Elements) {
              let _458_v52 = _compr_35;
              if (_dafny.Seq.contains(_dafny.Seq.update(_457_v54, _module.__default.safeIndex(new BigNumber(-468), new BigNumber((_457_v54).length)), _456_v53), _458_v52)) {
                _coll35.push([_458_v52,p0]);
              }
            }
            return _coll35;
          }()) : (_423_v27)));
          let _459_v55;
          _459_v55 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,(_420_v24)[_module.__default.safeIndex(p0, new BigNumber((_420_v24).length))]);
          let _460_v56;
          let _nw73 = Array((new BigNumber(9)).toNumber());
          _nw73[(_dafny.ZERO).toNumber()] = new BigNumber(122);
          _nw73[(_dafny.ONE).toNumber()] = p0;
          _nw73[(new BigNumber(2)).toNumber()] = new BigNumber((_394_v4).length);
          _nw73[(new BigNumber(3)).toNumber()] = (new BigNumber(930)).minus(p0);
          _nw73[(new BigNumber(4)).toNumber()] = _module.__default.fm7(p0, _459_v55, _455_v51, globalState);
          _nw73[(new BigNumber(5)).toNumber()] = p0;
          _nw73[(new BigNumber(6)).toNumber()] = (p0).multipliedBy(p0);
          _nw73[(new BigNumber(7)).toNumber()] = new BigNumber(693);
          _nw73[(new BigNumber(8)).toNumber()] = new BigNumber(183);
          _460_v56 = _nw73;
          let _index61 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_460_v56).length));
          (_460_v56)[_index61] = p0;
          let _index62 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_460_v56).length));
          (_460_v56)[_index62] = _module.__default.safeModuloInt(p0, p0);
          let _461_v57;
          _461_v57 = _dafny.Set.fromElements((((_459_v55).contains(_422_v26)) ? ((_459_v55).get(_422_v26)) : (p0)));
          _461_v57 = _461_v57;
          let _462_v58;
          let _nw74 = new _module.C0();
          _nw74.__ctor();
          _462_v58 = _nw74;
        }
        let _463_v59;
        _463_v59 = _dafny.Map.Empty.slice().updateUnsafe(_447_v44,_419_v23);
        _463_v59 = ((_463_v59).Merge(_463_v59)).Merge(_463_v59);
        (globalState).f3 = (_this).fm5(new BigNumber(655), _420_v24, globalState);
        (globalState).f14 = p0;
      } else {
        let _464_v60;
        let _nw75 = new _module.C0();
        _nw75.__ctor();
        _464_v60 = _nw75;
        (globalState).f9 = _dafny.areEqual(_419_v23, _dafny.Seq.update(_419_v23, _module.__default.safeIndex(new BigNumber(735), new BigNumber((_419_v23).length)), _392_v0));
        let _465_v61;
        _465_v61 = _dafny.MultiSet.fromElements(_422_v26);
        let _466_v62;
        _466_v62 = _module.D2.create_DC5(_465_v61);
        let _467_v63;
        _467_v63 = _module.D1.create_DC4(_422_v26, _module.__default.fm11(p0, _422_v26, _466_v62, globalState), p0);
        let _468_v64;
        _468_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm5(new BigNumber(604), _420_v24, globalState));
        let _469_v65;
        _469_v65 = _dafny.Set.fromElements(_468_v64);
        let _470_v66;
        _470_v66 = _dafny.Set.fromElements(_467_v63, _467_v63, _467_v63, _module.D1.create_DC4(false, _469_v65, p0), _467_v63);
        let _471_v67;
        _471_v67 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,_470_v66);
        let _rhs56 = _471_v67;
        let _rhs57 = (p0).multipliedBy(p0);
        let _rhs58 = !(false);
        let _rhs59 = p0;
        let _lhs46 = globalState;
        let _lhs47 = globalState;
        _471_v67 = _rhs56;
        _lhs46.f14 = _rhs57;
        _422_v26 = _rhs58;
        _lhs47.f14 = _rhs59;
        (globalState).f14 = p0;
        let _nw76 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        (globalState).f7 = _nw76;
      }
      r0 = _dafny.Seq.Concat(_419_v23, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), ((_472_v0) => function (_473_i6) {
        return _472_v0;
      })(_392_v0)), _419_v23));
      let _474_v69;
      _474_v69 = _dafny.Map.Empty.slice().updateUnsafe(_422_v26,p0);
      let _475_v70;
      _475_v70 = _dafny.Seq.of(_474_v69, _474_v69);
      let _476_v71;
      _476_v71 = _dafny.MultiSet.fromElements(_422_v26);
      let _477_v72;
      _477_v72 = _module.D0.create_DC0(_422_v26);
      let _478_v73;
      _478_v73 = _module.D0.create_DC1(_419_v23, new BigNumber(((function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of (_475_v70).Elements) {
    let _479_v68 = _compr_36;
    if (_dafny.Seq.contains(_475_v70, _479_v68)) {
      _coll36.push([_479_v68,_422_v26]);
    }
  }
  return _coll36;
}()).update(_dafny.Map.Empty.slice().updateUnsafe(_422_v26,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length)), !(true))).length), new BigNumber((_476_v71).cardinality()), _module.__default.fm12(_477_v72, globalState));
      r1 = _478_v73;
      let _480_v74;
      _480_v74 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), function (_481_i8) {
        return _dafny.Set.fromElements(new BigNumber(-428));
      }));
      let _482_v75;
      _482_v75 = _dafny.Set.fromElements(new BigNumber(-46), p0);
      let _483_v76;
      _483_v76 = _dafny.Seq.of(_482_v75, _482_v75, _482_v75);
      r2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_484_p0, _485_v4) => function (_486_i7) {
        return _dafny.Set.fromElements(_484_p0, new BigNumber((_485_v4).length));
      })(p0, _394_v4)), (((_480_v74).contains(p0)) ? ((_480_v74).get(p0)) : (_483_v76)));
      r3 = _422_v26;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f17 = _dafny.Set.Empty;
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f17, f18) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm19(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements((_this).f18), _dafny.Set.fromElements((_this).f18, (_this).f18)), (((_this).f18) ? (_dafny.Seq.of(_dafny.Set.fromElements((_this).f18, (_this).f18, (_this).f18), _dafny.Set.fromElements((_this).f18))) : (_dafny.Seq.of(_dafny.Set.fromElements((_this).f18), _dafny.Set.fromElements((_this).f18), _dafny.Set.fromElements(!(false))))));
    };
    fm20(globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("lxunet");
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f19, f17, f18) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.UnicodeFromString("ooi")).length);
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _487_v0;
      let _nw77 = new _module.C1();
      _nw77.__ctor();
      _487_v0 = _nw77;
      (globalState).f9 = ((!(p1) || (p1)) ? (p1) : ((_this).f18));
      let _488_v1;
      _488_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
      let _489_v2;
      _489_v2 = _dafny.Set.fromElements(_488_v1);
      let _490_v3;
      _490_v3 = _module.D1.create_DC4(true, _489_v2, new BigNumber(164));
      let _491_v4;
      _491_v4 = _dafny.MultiSet.fromElements(_490_v3);
      let _492_v5;
      _492_v5 = _dafny.Seq.of(_module.D1.create_DC4(p1, _489_v2, p2), _490_v3);
      let _493_v6;
      _493_v6 = _dafny.Seq.of(p1, p1, (_this).f18);
      let _494_v7;
      _494_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.MultiSet.fromElements(p0, new BigNumber((_493_v6).length)));
      let _495_v8;
      _495_v8 = _dafny.MultiSet.fromElements(p2, p2);
      let _496_v9;
      _496_v9 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.FromArray(_492_v5)).IsSubsetOf(_491_v4),((((_494_v7).contains(p1)) ? ((_494_v7).get(p1)) : (_495_v8))).Union(_495_v8));
      _496_v9 = _496_v9;
      let _497_v10;
      _497_v10 = _dafny.Seq.UnicodeFromString("w");
      let _498_v11;
      let _nw78 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _498_v11 = _nw78;
      let _499_v12;
      _499_v12 = _module.D0.create_DC2(p1, new BigNumber((_497_v10).length), _498_v11);
      let _500_v13;
      _500_v13 = new _dafny.CodePoint('v'.codePointAt(0));
      let _501_v14;
      _501_v14 = _dafny.Seq.of((_this).fm0((_dafny.ZERO).minus(p2), _500_v13, globalState));
      let _502_v15;
      _502_v15 = _module.D3.create_DC9();
      let _503_v16;
      _503_v16 = _dafny.Seq.of(_502_v15);
      let _504_v17;
      _504_v17 = _503_v16;
      let _505_v18;
      _505_v18 = _module.D5.create_DC13(p1, (_501_v14)[_module.__default.safeIndex(p2, new BigNumber((_501_v14).length))], _504_v17);
      if ((((p1) ? (_499_v12) : (_module.D0.create_DC2(p1, (_505_v18).dtor_cf24, _498_v11)))).dtor_cf5) {
        r0 = (_this).f18;
        let _506_v19;
        let _nw79 = new _module.C0();
        _nw79.__ctor();
        _506_v19 = _nw79;
        (globalState).f14 = p0;
        let _507_v20;
        _507_v20 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18) || (false),(_dafny.ZERO).minus(p2));
        _507_v20 = _507_v20;
        if (!((_dafny.ZERO).minus(p0)).isEqualTo((p0).plus(new BigNumber((_497_v10).length)))) {
          (globalState).f3 = !(p1) || ((p1) || ((_this).f18));
          let _508_v21;
          _508_v21 = _dafny.Set.fromElements((_this).f18, false, p1, p1, (_this).f18);
          _508_v21 = (_module.D7.create_DC16(_508_v21)).dtor_cf30;
          let _509_v22;
          let _nw80 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _509_v22 = _nw80;
          _509_v22 = (((_this).f18) ? (_509_v22) : (_509_v22));
          let _510_v23;
          let _nw81 = Array((new BigNumber(15)).toNumber()).fill([]);
          _510_v23 = _nw81;
          _510_v23 = _510_v23;
          let _511_v24;
          _511_v24 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p0)).length)).isLessThanOrEqualTo(p2),_497_v10);
          _511_v24 = (_511_v24).update(false, _497_v10);
        } else {
          let _512_v25;
          let _init15 = ((_513_v6) => function (_514_i0) {
            return _513_v6;
          })(_493_v6);
          let _nw82 = Array((new BigNumber(15)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw82.length); _i0_15++) {
            _nw82[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _512_v25 = _nw82;
          let _index63 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_512_v25).length));
          (_512_v25)[_index63] = _dafny.Seq.update(_493_v6, _module.__default.safeIndex(p2, new BigNumber((_493_v6).length)), p1);
          let _index64 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_512_v25).length));
          (_512_v25)[_index64] = _493_v6;
          let _rhs60 = false;
          let _rhs61 = (_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0))).minus(p2);
          let _lhs48 = globalState;
          r0 = _rhs60;
          _lhs48.f14 = _rhs61;
          let _515_v26;
          let _nw83 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _515_v26 = _nw83;
          let _index65 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_515_v26).length));
          (_515_v26)[_index65] = _500_v13;
          let _516_v27;
          _516_v27 = _module.D3.create_DC10(_497_v10, p1, _497_v10, _dafny.Seq.Concat(_497_v10, _497_v10), p2);
          let _index66 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_515_v26).length));
          let _rhs62 = p0;
          let _rhs63 = _500_v13;
          let _rhs64 = _module.__default.safeDivisionInt((p2).minus(p0), p0);
          let _rhs65 = _dafny.MultiSet.fromElements(p2);
          let _rhs66 = _516_v27;
          let _lhs49 = globalState;
          let _lhs50 = _515_v26;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_515_v26).length));
          let _lhs52 = globalState;
          let _lhs53 = globalState;
          _lhs49.f14 = _rhs62;
          _lhs50[_lhs51] = _rhs63;
          _lhs52.f14 = _rhs64;
          _lhs53.f15 = _rhs65;
          _516_v27 = _rhs66;
          let _517_v28;
          _517_v28 = _dafny.MultiSet.fromElements(p1);
          let _518_v29;
          _518_v29 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, (_this).f18, (_this).f18)).length)),_517_v28);
          _518_v29 = (_518_v29).update(new BigNumber((_497_v10).length), (_517_v28).Difference(_517_v28));
          let _519_v30;
          _519_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(532));
          let _rhs67 = !((_519_v30).contains(((p1) ? (p0) : (p2))));
          let _rhs68 = _this.f19;
          let _lhs54 = globalState;
          let _lhs55 = _this;
          _lhs54.f9 = _rhs67;
          _lhs55.f19 = _rhs68;
        }
      } else {
        _497_v10 = _497_v10;
        (globalState).f14 = p0;
        _500_v13 = _500_v13;
        let _520_v31;
        _520_v31 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _521_v32;
        let _522_v33;
        let _523_v34;
        let _524_v35;
        let _out33;
        let _out34;
        let _out35;
        let _out36;
        let _outcollector9 = (_487_v0).m5((((_520_v31).contains(new BigNumber((_497_v10).length))) ? ((_520_v31).get(new BigNumber((_497_v10).length))) : (p0)), globalState);
        _out33 = _outcollector9[0];
        _out34 = _outcollector9[1];
        _out35 = _outcollector9[2];
        _out36 = _outcollector9[3];
        _521_v32 = _out33;
        _522_v33 = _out34;
        _523_v34 = _out35;
        _524_v35 = _out36;
        (globalState).f3 = !(p0).isEqualTo(p2);
      }
      if (false) {
        (globalState).f3 = p1;
        (globalState).f3 = (_this).f18;
        let _arr0 = _this.f19;
        let _index67 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length));
        _arr0[_index67] = p1;
        let _525_v36;
        _525_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.MultiSet.fromElements(_498_v11, _498_v11, _498_v11));
        let _526_v37;
        _526_v37 = _dafny.MultiSet.fromElements(_498_v11, _498_v11);
        let _arr1 = _this.f19;
        let _index68 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length));
        _arr1[_index68] = (((p1) || (false)) ? ((_this).f18) : (!((_dafny.ZERO).minus(new BigNumber(((((_525_v36).contains(p2)) ? ((_525_v36).get(p2)) : (_526_v37))).cardinality()))).isEqualTo(new BigNumber(558))));
        let _527_v38;
        _527_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_497_v10).length),(_dafny.ZERO).minus((_dafny.ZERO).minus(p0)));
        let _528_v39;
        _528_v39 = _dafny.MultiSet.fromElements(p1);
        let _529_v40;
        let _nw84 = Array((new BigNumber(22)).toNumber());
        _nw84[(_dafny.ZERO).toNumber()] = true;
        _nw84[(_dafny.ONE).toNumber()] = false;
        _nw84[(new BigNumber(2)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(3)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(4)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(5)).toNumber()] = p1;
        _nw84[(new BigNumber(6)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(7)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(8)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(9)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(10)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(11)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(12)).toNumber()] = !((_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))]);
        _nw84[(new BigNumber(13)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(14)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(15)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(16)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(17)).toNumber()] = p1;
        _nw84[(new BigNumber(18)).toNumber()] = (_this).f18;
        _nw84[(new BigNumber(19)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))];
        _nw84[(new BigNumber(20)).toNumber()] = p1;
        _nw84[(new BigNumber(21)).toNumber()] = (_this).f18;
        _529_v40 = _nw84;
        let _530_v41;
        _530_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(_module.D3.create_DC10(_497_v10, (_this).f18, _497_v10, _497_v10, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), function (_531_i1) {
  return new BigNumber(664);
})).length)), _module.__default.fm18(new _dafny.CodePoint('d'.codePointAt(0)), (_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))], p2, p1, globalState), _module.D3.create_DC10(_497_v10, p1, _497_v10, _497_v10, p0)));
        let _532_v42;
        _532_v42 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_493_v6).length));
        let _533_v44;
        _533_v44 = _dafny.Set.fromElements((((_532_v42).contains((_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))])) ? ((_532_v42).get((_this.f19)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_this.f19).length))])) : (new BigNumber((function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of _dafny.IntegerRange(new BigNumber(490), new BigNumber(49))) {
            let _534_v43 = _compr_37;
            if (((new BigNumber(490)).isLessThanOrEqualTo(_534_v43)) && ((_534_v43).isLessThan(new BigNumber(49)))) {
              _coll37.push([_module.__default.safeModuloInt(_534_v43, p0),(_490_v3).dtor_cf9]);
            }
          }
          return _coll37;
        }()).length))), p2);
        let _535_v45;
        let _nw85 = new _module.C2();
        _nw85.__ctor(_this.f17, (_this).f18);
        _535_v45 = _nw85;
        let _536_v46;
        _536_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_533_v44).length),_535_v45);
        let _537_v47;
        _537_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_487_v0).fm5(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-349)), ((_538_p2) => function (_539_i2) {
          return _538_p2;
        })(p2))).length), _501_v14, globalState), true),_536_v46);
        let _nw86 = Array((new BigNumber(27)).toNumber());
        _nw86[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((((_527_v38).contains(p0)) ? ((_527_v38).get(p0)) : (p0)));
        _nw86[(_dafny.ONE).toNumber()] = p2;
        _nw86[(new BigNumber(2)).toNumber()] = p0;
        _nw86[(new BigNumber(3)).toNumber()] = (p0).multipliedBy(p2);
        _nw86[(new BigNumber(4)).toNumber()] = p0;
        _nw86[(new BigNumber(5)).toNumber()] = p2;
        _nw86[(new BigNumber(6)).toNumber()] = p2;
        _nw86[(new BigNumber(7)).toNumber()] = p2;
        _nw86[(new BigNumber(8)).toNumber()] = p0;
        _nw86[(new BigNumber(9)).toNumber()] = (_this).fm0(new BigNumber((_dafny.Seq.update(_497_v10, _module.__default.safeIndex(p2, new BigNumber((_497_v10).length)), _500_v13)).length), _500_v13, globalState);
        _nw86[(new BigNumber(10)).toNumber()] = p2;
        _nw86[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_528_v39).cardinality()),_529_v40)).length);
        _nw86[(new BigNumber(12)).toNumber()] = (_this).fm0(p2, _500_v13, globalState);
        _nw86[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_493_v6).length), p0);
        _nw86[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_530_v41).length)).minus(p0));
        _nw86[(new BigNumber(15)).toNumber()] = (_this).fm0(p2, _500_v13, globalState);
        _nw86[(new BigNumber(16)).toNumber()] = p0;
        _nw86[(new BigNumber(17)).toNumber()] = p2;
        _nw86[(new BigNumber(18)).toNumber()] = new BigNumber((_537_v47).length);
        _nw86[(new BigNumber(19)).toNumber()] = p2;
        _nw86[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_497_v10).length), (_this).fm0(new BigNumber((_dafny.Seq.UnicodeFromString("pwulprhlu")).length), new _dafny.CodePoint('g'.codePointAt(0)), globalState));
        _nw86[(new BigNumber(21)).toNumber()] = new BigNumber(738);
        _nw86[(new BigNumber(22)).toNumber()] = p2;
        _nw86[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ytgmpt")).length);
        _nw86[(new BigNumber(24)).toNumber()] = p2;
        _nw86[(new BigNumber(25)).toNumber()] = p2;
        _nw86[(new BigNumber(26)).toNumber()] = _module.__default.safeModuloInt(p2, new BigNumber((_493_v6).length));
        (globalState).f7 = _nw86;
        let _540_v48;
        let _541_v49;
        let _542_v50;
        let _543_v51;
        let _out37;
        let _out38;
        let _out39;
        let _out40;
        let _outcollector10 = (_487_v0).m5(p2, globalState);
        _out37 = _outcollector10[0];
        _out38 = _outcollector10[1];
        _out39 = _outcollector10[2];
        _out40 = _outcollector10[3];
        _540_v48 = _out37;
        _541_v49 = _out38;
        _542_v50 = _out39;
        _543_v51 = _out40;
      } else {
        if ((_this).f18) {
          let _544_v52;
          let _nw87 = new _module.C2();
          _nw87.__ctor(_this.f17, (((_this).f18) ? (p1) : (p1)));
          _544_v52 = _nw87;
          let _545_v53;
          _545_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,p2);
          let _rhs69 = p0;
          let _rhs70 = _dafny.Seq.update((_544_v52).fm20(globalState), _module.__default.safeIndex((((_this).f18) ? ((((_545_v53).contains(p1)) ? ((_545_v53).get(p1)) : (p0))) : ((_dafny.ZERO).minus(p0))), new BigNumber(((_544_v52).fm20(globalState)).length)), _500_v13);
          let _lhs56 = globalState;
          _lhs56.f14 = _rhs69;
          _497_v10 = _rhs70;
          (globalState).f3 = p1;
          let _546_v54;
          _546_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
          let _547_v55;
          _547_v55 = _dafny.Map.Empty.slice().updateUnsafe((_487_v0).fm5(new BigNumber((_546_v54).length), _501_v14, globalState),false);
          _547_v55 = (_546_v54).Merge((_546_v54).update((_this).f18, p1));
          (globalState).f9 = p1;
        } else {
          r2 = p0;
          let _548_v56;
          _548_v56 = _module.D5.create_DC14(p1, new BigNumber(-226), new BigNumber((_493_v6).length));
          let _549_v57;
          _549_v57 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _550_v58;
          _550_v58 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_548_v56)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_548_v56)),_549_v57);
          _550_v58 = (_550_v58).Merge(_550_v58);
          let _551_v59;
          let _nw88 = new _module.C1();
          _nw88.__ctor();
          _551_v59 = _nw88;
          let _552_v60;
          let _nw89 = Array((new BigNumber(2)).toNumber()).fill([]);
          _552_v60 = _nw89;
          let _553_v61;
          _553_v61 = _dafny.Set.fromElements(_500_v13, _500_v13);
          let _554_v62;
          _554_v62 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Set.fromElements((_497_v10)[_module.__default.safeIndex(p0, new BigNumber((_497_v10).length))], _500_v13, new _dafny.CodePoint('h'.codePointAt(0))));
          let _555_v63;
          _555_v63 = _dafny.Seq.of(_553_v61, _dafny.Set.fromElements(_500_v13, _500_v13));
          let _556_v65;
          let _nw90 = Array((new BigNumber(15)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _553_v61;
          _nw90[(_dafny.ONE).toNumber()] = (((_554_v62).contains(true)) ? ((_554_v62).get(true)) : (_553_v61));
          _nw90[(new BigNumber(2)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(3)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(4)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(5)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(6)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(7)).toNumber()] = (_555_v63)[_module.__default.safeIndex((_this).fm0(p0, _500_v13, globalState), new BigNumber((_555_v63).length))];
          _nw90[(new BigNumber(8)).toNumber()] = function () {
            let _coll38 = new _dafny.Set();
            for (const _compr_38 of (_497_v10).Elements) {
              let _557_v64 = _compr_38;
              if (_dafny.Seq.contains(_497_v10, _557_v64)) {
                _coll38.add(_557_v64);
              }
            }
            return _coll38;
          }();
          _nw90[(new BigNumber(9)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(10)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(11)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(12)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(13)).toNumber()] = _553_v61;
          _nw90[(new BigNumber(14)).toNumber()] = _553_v61;
          _556_v65 = _nw90;
          let _index69 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_552_v60).length));
          (_552_v60)[_index69] = _556_v65;
          let _index70 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_552_v60).length));
          (_552_v60)[_index70] = _556_v65;
          let _558_v66;
          _558_v66 = _dafny.Set.fromElements(p2);
          (globalState).f9 = !(_558_v66).contains(p0);
        }
        let _index71 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length));
        (_498_v11)[_index71] = _module.__default.safeDivisionInt(new BigNumber(-741), p0);
        let _index72 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length));
        (_498_v11)[_index72] = p2;
        let _559_v67;
        let _init16 = ((_560_v10, _561_v11) => function (_562_i3) {
          return _dafny.Seq.update(_560_v10, _module.__default.safeIndex((_561_v11)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_561_v11).length))], new BigNumber((_560_v10).length)), new _dafny.CodePoint('q'.codePointAt(0)));
        })(_497_v10, _498_v11);
        let _nw91 = Array((new BigNumber(28)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw91.length); _i0_16++) {
          _nw91[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _559_v67 = _nw91;
        let _index73 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_559_v67).length));
        (_559_v67)[_index73] = _dafny.Seq.Concat(_497_v10, _497_v10);
        let _index74 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_559_v67).length));
        (_559_v67)[_index74] = _497_v10;
        let _index75 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_559_v67).length));
        let _index76 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length));
        let _index77 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_559_v67).length));
        let _rhs71 = _module.__default.fm21(_dafny.Seq.of((_498_v11)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length))], p0, p0, (_dafny.ZERO).minus(p0), p2), globalState);
        let _rhs72 = _dafny.Seq.UnicodeFromString("nosuekq");
        let _rhs73 = (_this).fm0((_498_v11)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length))], _500_v13, globalState);
        let _rhs74 = _module.__default.fm22(p0, globalState);
        let _lhs57 = _559_v67;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_559_v67).length));
        let _lhs59 = _498_v11;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_498_v11).length));
        let _lhs61 = _559_v67;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_559_v67).length));
        _488_v1 = _rhs71;
        _lhs57[_lhs58] = _rhs72;
        _lhs59[_lhs60] = _rhs73;
        _lhs61[_lhs62] = _rhs74;
        let _arr2 = _this.f19;
        let _index78 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_this.f19).length));
        _arr2[_index78] = (_this).f18;
        let _arr3 = _this.f19;
        let _index79 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_this.f19).length));
        _arr3[_index79] = !((_this).f18);
        let _563_v68;
        let _nw92 = new _module.C1();
        _nw92.__ctor();
        _563_v68 = _nw92;
      }
      let _564_v69;
      _564_v69 = _module.D2.create_DC7(_module.__default.fm23(p0, new BigNumber(942), (_this).f18, globalState));
      let _565_v70;
      _565_v70 = _module.D2.create_DC7(_564_v69);
      let _source4 = _565_v70;
      if (_source4.is_DC6) {
        let _566___mcc_h0 = (_source4).cf13;
        let _567_cf13 = _566___mcc_h0;
        let _568_v71;
        let _nw93 = Array((new BigNumber(2)).toNumber()).fill([]);
        _568_v71 = _nw93;
        let _569_v72;
        let _nw94 = Array((new BigNumber(3)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = _498_v11;
        _nw94[(_dafny.ONE).toNumber()] = _498_v11;
        _nw94[(new BigNumber(2)).toNumber()] = _498_v11;
        _569_v72 = _nw94;
        let _index80 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_568_v71).length));
        (_568_v71)[_index80] = _569_v72;
        let _index81 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_568_v71).length));
        (_568_v71)[_index81] = _569_v72;
        let _570_v73;
        let _nw95 = new _module.C2();
        _nw95.__ctor(_this.f17, (_this).f18);
        _570_v73 = _nw95;
        let _rhs75 = _567_cf13;
        let _rhs76 = _570_v73;
        let _lhs63 = globalState;
        _lhs63.f14 = _rhs75;
        _570_v73 = _rhs76;
        let _rhs77 = p1;
        let _rhs78 = p2;
        let _lhs64 = globalState;
        let _lhs65 = globalState;
        _lhs64.f3 = _rhs77;
        _lhs65.f14 = _rhs78;
        (globalState).f3 = (((_this).f18) && (p1)) || ((_this).f18);
      } else if (_source4.is_DC5) {
        let _571___mcc_h1 = (_source4).cf12;
        let _572_cf12 = _571___mcc_h1;
        _497_v10 = _dafny.Seq.Concat(_497_v10, _497_v10);
        let _573_v74;
        _573_v74 = _dafny.MultiSet.fromElements(_this.f19, _this.f19, _this.f19, _this.f19);
        let _574_v75;
        _574_v75 = _dafny.Map.Empty.slice().updateUnsafe((_501_v14)[_module.__default.safeIndex(p0, new BigNumber((_501_v14).length))],p0);
        (globalState).f14 = (new BigNumber((_573_v74).cardinality())).multipliedBy((((_574_v75).contains(p0)) ? ((_574_v75).get(p0)) : (p0)));
        if (!((_this).f18)) {
          (globalState).f9 = (_this).f18;
          let _575_v76;
          let _nw96 = new _module.C0();
          _nw96.__ctor();
          _575_v76 = _nw96;
          (globalState).f9 = (_this).f18;
          let _576_v77;
          let _nw97 = Array((new BigNumber(5)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _487_v0;
          _nw97[(_dafny.ONE).toNumber()] = ((p1) ? (_487_v0) : (_487_v0));
          _nw97[(new BigNumber(2)).toNumber()] = _487_v0;
          _nw97[(new BigNumber(3)).toNumber()] = _487_v0;
          _nw97[(new BigNumber(4)).toNumber()] = _487_v0;
          _576_v77 = _nw97;
          let _577_v78;
          let _init17 = function (_578_i4) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          };
          let _nw98 = Array((new BigNumber(14)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw98.length); _i0_17++) {
            _nw98[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _577_v78 = _nw98;
          let _index82 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_577_v78).length));
          (_577_v78)[_index82] = ((p1) ? (_500_v13) : (_500_v13));
          let _579_v79;
          _579_v79 = _dafny.Seq.of((((_487_v0).fm5(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), ((_580_p0) => function (_581_i5) {
            return _580_p0;
          })(p0)), globalState)) ? (_576_v77) : (_576_v77)));
          let _index83 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_577_v78).length));
          let _rhs79 = (_579_v79)[_module.__default.safeIndex(p2, new BigNumber((_579_v79).length))];
          let _rhs80 = _500_v13;
          let _lhs66 = _577_v78;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_577_v78).length));
          _576_v77 = _rhs79;
          _lhs66[_lhs67] = _rhs80;
          (globalState).f3 = (_this).f18;
        } else {
          let _index84 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length));
          (_498_v11)[_index84] = (p0).plus(new BigNumber(205));
          let _582_v80;
          _582_v80 = _dafny.Set.fromElements(_497_v10);
          let _583_v81;
          let _nw99 = Array((new BigNumber(26)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = _497_v10;
          _nw99[(_dafny.ONE).toNumber()] = _497_v10;
          _nw99[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_497_v10, _497_v10);
          _nw99[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), ((_584_v13) => function (_585_i6) {
            return _584_v13;
          })(_500_v13));
          _nw99[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_497_v10, _497_v10);
          _nw99[(new BigNumber(5)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(6)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_497_v10, _module.__default.safeIndex(new BigNumber((_582_v80).length), new BigNumber((_497_v10).length)), _500_v13);
          _nw99[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("gdic");
          _nw99[(new BigNumber(9)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm22(new BigNumber((_572_cf12).cardinality()), globalState), _dafny.Seq.UnicodeFromString("fe"));
          _nw99[(new BigNumber(11)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(12)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(13)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_586_v13) => function (_587_i7) {
            return _586_v13;
          })(_500_v13)), _497_v10);
          _nw99[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_497_v10, _dafny.Seq.UnicodeFromString("su"));
          _nw99[(new BigNumber(16)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), ((_588_v13) => function (_589_i8) {
            return _588_v13;
          })(_500_v13));
          _nw99[(new BigNumber(18)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(19)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(20)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("wfxv");
          _nw99[(new BigNumber(22)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(917)), ((_590_v13) => function (_591_i9) {
            return _590_v13;
          })(_500_v13));
          _nw99[(new BigNumber(23)).toNumber()] = _497_v10;
          _nw99[(new BigNumber(24)).toNumber()] = _module.__default.fm22(p2, globalState);
          _nw99[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_497_v10, _497_v10);
          _583_v81 = _nw99;
          let _index85 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length));
          (_583_v81)[_index85] = _497_v10;
          let _index86 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length));
          let _index87 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length));
          let _rhs81 = (p1) && ((_this).f18);
          let _rhs82 = p1;
          let _rhs83 = (p0).minus((((_this).f18) ? (new BigNumber(394)) : (p2)));
          let _rhs84 = _497_v10;
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          let _lhs70 = _498_v11;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length));
          let _lhs72 = _583_v81;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length));
          _lhs68.f3 = _rhs81;
          _lhs69.f3 = _rhs82;
          _lhs70[_lhs71] = _rhs83;
          _lhs72[_lhs73] = _rhs84;
          (globalState).f14 = new BigNumber((_dafny.Seq.Concat((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))], _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yh"), (_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))]))).length);
          let _592_v82;
          let _nw100 = new _module.C2();
          _nw100.__ctor(_this.f17, (_this).f18);
          _592_v82 = _nw100;
          let _593_v83;
          _593_v83 = _dafny.Map.Empty.slice().updateUnsafe(p1,_592_v82);
          let _594_v84;
          _594_v84 = _dafny.Map.Empty.slice().updateUnsafe((((_593_v83).contains(false)) ? ((_593_v83).get(false)) : (_592_v82)),(_this).f18);
          let _595_v85;
          _595_v85 = _dafny.Seq.of(_594_v84);
          let _index88 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length));
          (_498_v11)[_index88] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_595_v85, _dafny.Seq.of(_594_v84, _594_v84)), _595_v85)).length);
          let _596_v86;
          _596_v86 = _module.D0.create_DC1((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))], (_498_v11)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length))], (_498_v11)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length))], ((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))])[_module.__default.safeIndex(new BigNumber(((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))]).length), new BigNumber(((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))]).length))]);
          _596_v86 = (((_this).f18) ? (_596_v86) : (_596_v86));
          let _arr4 = _this.f19;
          let _index89 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_this.f19).length));
          _arr4[_index89] = true;
          let _597_v87;
          let _nw101 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _597_v87 = _nw101;
          let _598_v88;
          _598_v88 = _dafny.Seq.of(_597_v87, _597_v87);
          let _599_v89;
          _599_v89 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-552),_597_v87);
          let _arr5 = _this.f19;
          let _index90 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_this.f19).length));
          let _rhs85 = ((_dafny.Seq.IsProperPrefixOf((_583_v81)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_583_v81).length))], _dafny.Seq.UnicodeFromString("g"))) ? ((_598_v88)[_module.__default.safeIndex((_498_v11)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length))], new BigNumber((_598_v88).length))]) : ((((_599_v89).contains((_498_v11)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length))])) ? ((_599_v89).get((_498_v11)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_498_v11).length))])) : (_597_v87))));
          let _rhs86 = p1;
          let _rhs87 = (_this).f18;
          let _rhs88 = (((_this).f18) ? (p1) : ((_this).f18));
          let _lhs74 = globalState;
          let _lhs75 = globalState;
          let _lhs76 = _this.f19;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_this.f19).length));
          let _lhs78 = globalState;
          _lhs74.f7 = _rhs85;
          _lhs75.f3 = _rhs86;
          _lhs76[_lhs77] = _rhs87;
          _lhs78.f9 = _rhs88;
        }
        let _index91 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_498_v11).length));
        (_498_v11)[_index91] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_493_v6).length), new BigNumber((_501_v14).length))).length), p0);
        let _index92 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_498_v11).length));
        (_498_v11)[_index92] = p0;
      } else {
        let _600___mcc_h2 = (_source4).cf14;
        let _601_cf14 = _600___mcc_h2;
        let _602_v90;
        _602_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f19);
        _602_v90 = (_602_v90).update(((_this).f18) === ((_this).f18), _this.f19);
        r1 = (_this).f18;
        r0 = !((_this).f18);
        let _603_v91;
        _603_v91 = _dafny.Map.Empty.slice().updateUnsafe(_498_v11,_497_v10);
        _603_v91 = (_603_v91).update(_498_v11, _497_v10);
      }
      let _604_v92;
      _604_v92 = _dafny.Set.fromElements(p0);
      let _605_v93;
      _605_v93 = _dafny.Seq.of(_604_v92);
      r0 = (((function () {
        let _coll39 = new _dafny.Set();
        for (const _compr_39 of (_605_v93).Elements) {
          let _606_v94 = _compr_39;
          if (_dafny.Seq.contains(_605_v93, _606_v94)) {
            _coll39.add(_606_v94);
          }
        }
        return _coll39;
      }()).contains(_604_v92)) ? (!((((_this).f18) ? ((_this).f18) : ((_this).f18)))) : (p1));
      r1 = (_487_v0).fm5(_module.__default.safeDivisionInt(p2, p0), _dafny.Seq.Concat(_501_v14, _dafny.Seq.of(p0)), globalState);
      r2 = p2;
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f17 = _dafny.Set.Empty;
      this._f18 = false;
      this.f26 = _dafny.ZERO;
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f25, f26, f17, f18) {
      let _this = this;
      (_this)._f25 = f25;
      (_this).f26 = f26;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _607_v0;
      let _608_v1;
      let _609_v2;
      let _out41;
      let _out42;
      let _out43;
      let _outcollector11 = (_this).m4(globalState);
      _out41 = _outcollector11[0];
      _out42 = _outcollector11[1];
      _out43 = _outcollector11[2];
      _607_v0 = _out41;
      _608_v1 = _out42;
      _609_v2 = _out43;
      let _610_v3;
      let _init18 = ((_611_p0) => function (_612_i0) {
        return _611_p0;
      })(p0);
      let _nw102 = Array((new BigNumber(22)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw102.length); _i0_18++) {
        _nw102[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _610_v3 = _nw102;
      let _index93 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_610_v3).length));
      (_610_v3)[_index93] = p0;
      let _index94 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_610_v3).length));
      (_610_v3)[_index94] = p0;
      let _613_v4;
      let _init19 = ((_614_v2) => function (_615_i2) {
        return (_615_i2).plus(_614_v2);
      })(_609_v2);
      let _nw103 = Array((new BigNumber(20)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw103.length); _i0_19++) {
        _nw103[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _613_v4 = _nw103;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_613_v4).length))) {
        let _616_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_616_i1)) && ((_616_i1).isLessThan(new BigNumber((_613_v4).length))))) {
          (_613_v4)[(_616_i1)] = (_616_i1).multipliedBy(_this.f26);
        }
      }
      let _617_v5;
      _617_v5 = _dafny.Seq.UnicodeFromString("n");
      let _618_v6;
      _618_v6 = _dafny.Seq.of(_617_v5, _dafny.Seq.UnicodeFromString("vvpbsjve"));
      _618_v6 = _dafny.Seq.update(_618_v6, _module.__default.safeIndex(_609_v2, new BigNumber((_618_v6).length)), _dafny.Seq.Concat(_617_v5, _dafny.Seq.UnicodeFromString("aiiyumo")));
      (_this).f26 = ((_this).f25).multipliedBy(_609_v2);
      (_this).f26 = _this.f26;
      r0 = new BigNumber(636);
      return r0;
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      let _619_v0;
      let _init20 = function (_620_i0) {
        return (_620_i0).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f26, (_this).f25, (_this).f25)).length));
      };
      let _nw104 = Array((new BigNumber(6)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw104.length); _i0_20++) {
        _nw104[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _619_v0 = _nw104;
      let _pat_let_tv9 = _619_v0;
      let _source5 = function (_pat_let15_0) {
        return function (_621_dt__update__tmp_h0) {
          return function (_pat_let16_0) {
            return function (_622_dt__update_hcf6_h0) {
              return function (_pat_let17_0) {
                return function (_623_dt__update_hcf7_h0) {
                  return _module.D0.create_DC2((_621_dt__update__tmp_h0).dtor_cf5, _622_dt__update_hcf6_h0, _623_dt__update_hcf7_h0);
                }(_pat_let17_0);
              }(_pat_let_tv9);
            }(_pat_let16_0);
          }(_this.f26);
        }(_pat_let15_0);
      }(_module.D0.create_DC2((_this).f18, (_this).f25, _619_v0));
      if (_source5.is_DC0) {
        let _624___mcc_h0 = (_source5).cf0;
        let _625_cf0 = _624___mcc_h0;
        let _index95 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
        (_619_v0)[_index95] = _module.__default.safeDivisionInt((_this).f25, _this.f26);
        let _index96 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
        (_619_v0)[_index96] = new BigNumber(646);
        let _626_v1;
        _626_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this.f26).isLessThanOrEqualTo((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))]),false);
        let _627_v2;
        _627_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f18);
        let _628_v3;
        _628_v3 = _dafny.MultiSet.fromElements(_627_v2);
        _626_v1 = _module.__default.fm3(_this.f26, true, _628_v3, globalState);
        let _source6 = _module.D0.create_DC0((_this).f18);
        if (_source6.is_DC0) {
          let _629___mcc_h8 = (_source6).cf0;
          let _630_cf0 = _629___mcc_h8;
          let _631_v4;
          _631_v4 = _dafny.Seq.UnicodeFromString("qnhyyx");
          _631_v4 = _631_v4;
          let _632_v5;
          _632_v5 = new _dafny.CodePoint('w'.codePointAt(0));
          _632_v5 = _632_v5;
          let _index97 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
          (_619_v0)[_index97] = _module.__default.safeModuloInt(_module.__default.fm4((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))], globalState), _this.f26);
          let _633_v6;
          _633_v6 = _dafny.Set.fromElements(_630_cf0);
          let _634_v7;
          let _nw105 = Array((new BigNumber(5)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))];
          _nw105[(_dafny.ONE).toNumber()] = (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))];
          _nw105[(new BigNumber(2)).toNumber()] = (_this).f25;
          _nw105[(new BigNumber(3)).toNumber()] = _this.f26;
          _nw105[(new BigNumber(4)).toNumber()] = new BigNumber((_633_v6).length);
          _634_v7 = _nw105;
          let _635_v8;
          _635_v8 = _module.D0.create_DC2(_630_cf0, (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))], _634_v7);
          let _636_v9;
          _636_v9 = _dafny.MultiSet.fromElements((_635_v8).dtor_cf5);
          let _637_v10;
          _637_v10 = _dafny.Set.fromElements(new BigNumber((_636_v9).cardinality()));
          let _638_v11;
          _638_v11 = _dafny.Set.fromElements(_637_v10, _637_v10, _637_v10);
          let _639_v12;
          _639_v12 = _dafny.Map.Empty.slice().updateUnsafe(_637_v10,_dafny.Set.fromElements(_this.f26));
          (globalState).f9 = ((function () {
            let _coll40 = new _dafny.Set();
            for (const _compr_40 of (_639_v12).Keys.Elements) {
              let _640_v13 = _compr_40;
              if ((_639_v12).contains(_640_v13)) {
                _coll40.add(_640_v13);
              }
            }
            return _coll40;
          }()).Intersect(_638_v11)).IsProperSubsetOf((_638_v11).Union(_638_v11));
        } else if (_source6.is_DC1) {
          let _641___mcc_h9 = (_source6).cf1;
          let _642___mcc_h10 = (_source6).cf2;
          let _643___mcc_h11 = (_source6).cf3;
          let _644___mcc_h12 = (_source6).cf4;
          let _645_cf4 = _644___mcc_h12;
          let _646_cf3 = _643___mcc_h11;
          let _647_cf2 = _642___mcc_h10;
          let _648_cf1 = _641___mcc_h9;
          let _649_v14;
          _649_v14 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18, !(_625_cf0));
          _649_v14 = _649_v14;
          let _index98 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
          let _rhs89 = _this.f26;
          let _rhs90 = _this.f26;
          let _rhs91 = (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))];
          let _lhs79 = _619_v0;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
          _647_cf2 = _rhs89;
          _lhs79[_lhs80] = _rhs90;
          _646_cf3 = _rhs91;
          let _650_v15;
          let _nw106 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _650_v15 = _nw106;
          let _index99 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_650_v15).length));
          (_650_v15)[_index99] = _dafny.Seq.Concat(_648_cf1, _648_cf1);
          let _651_v16;
          _651_v16 = _dafny.Map.Empty.slice().updateUnsafe(!(_625_cf0),_645_cf4);
          let _index100 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_650_v15).length));
          (_650_v15)[_index100] = _dafny.Seq.update(_648_cf1, _module.__default.safeIndex((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))], new BigNumber((_648_cf1).length)), (((_651_v16).contains((_this).f18)) ? ((_651_v16).get((_this).f18)) : (new _dafny.CodePoint('p'.codePointAt(0)))));
          let _652_v17;
          let _nw107 = new _module.C0();
          _nw107.__ctor();
          _652_v17 = _nw107;
        } else {
          let _653___mcc_h13 = (_source6).cf5;
          let _654___mcc_h14 = (_source6).cf6;
          let _655___mcc_h15 = (_source6).cf7;
          let _656_cf7 = _655___mcc_h15;
          let _657_cf6 = _654___mcc_h14;
          let _658_cf5 = _653___mcc_h13;
          let _659_v18;
          let _init21 = ((_660_cf0) => function (_661_i1) {
            return ((true) ? (_660_cf0) : (_660_cf0));
          })(_625_cf0);
          let _nw108 = Array((new BigNumber(27)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw108.length); _i0_21++) {
            _nw108[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _659_v18 = _nw108;
          let _index101 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_659_v18).length));
          (_659_v18)[_index101] = false;
          let _index102 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_659_v18).length));
          (_659_v18)[_index102] = _658_cf5;
          let _662_v19;
          let _init22 = ((_663_v1) => function (_664_i2) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_663_v1), (_module.D3.create_DC8(_dafny.Seq.of(_663_v1))).dtor_cf15);
          })(_626_v1);
          let _nw109 = Array((new BigNumber(24)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw109.length); _i0_22++) {
            _nw109[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _662_v19 = _nw109;
          let _665_v20;
          _665_v20 = _dafny.Seq.of(_626_v1);
          let _index103 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_662_v19).length));
          (_662_v19)[_index103] = _dafny.Seq.Concat(_665_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-803)), ((_666_v1) => function (_667_i3) {
            return _666_v1;
          })(_626_v1)));
          let _index104 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_662_v19).length));
          (_662_v19)[_index104] = _665_v20;
          let _668_v21;
          _668_v21 = _module.D0.create_DC0(_658_cf5);
          let _669_v22;
          _669_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(globalState),_668_v21);
          _669_v22 = (_669_v22).update(_625_cf0, _668_v21);
          (globalState).f14 = _module.__default.safeModuloInt(_657_cf6, (_657_cf6).plus((_this).f25));
        }
        let _670_v23;
        _670_v23 = new _dafny.CodePoint('c'.codePointAt(0));
        let _671_v24;
        _671_v24 = _dafny.Seq.of(_670_v23, _670_v23);
        let _source7 = _module.D3.create_DC10(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(367)), function (_672_i4) {
  return new _dafny.CodePoint('q'.codePointAt(0));
}), _671_v24), _625_cf0, _dafny.Seq.Concat(_671_v24, _671_v24), _671_v24, _module.__default.fm4(new BigNumber((_626_v1).length), globalState));
        if (_source7.is_DC9) {
          (globalState).f14 = _module.__default.safeModuloInt((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))], (_this).f25);
          let _index105 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length));
          (_619_v0)[_index105] = ((_this).f25).minus(((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))]).multipliedBy((_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))]));
          let _673_v25;
          let _nw110 = Array((new BigNumber(15)).toNumber()).fill(false);
          _673_v25 = _nw110;
          let _674_v26;
          _674_v26 = _dafny.Map.Empty.slice().updateUnsafe(_673_v25,_670_v23);
          _674_v26 = (_674_v26).update(_673_v25, _670_v23);
          (globalState).f3 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-115)), ((_675_v23) => function (_676_i5) {
            return _675_v23;
          })(_670_v23)), _670_v23);
        } else if (_source7.is_DC10) {
          let _677___mcc_h16 = (_source7).cf16;
          let _678___mcc_h17 = (_source7).cf17;
          let _679___mcc_h18 = (_source7).cf18;
          let _680___mcc_h19 = (_source7).cf19;
          let _681___mcc_h20 = (_source7).cf20;
          let _682_cf20 = _681___mcc_h20;
          let _683_cf19 = _680___mcc_h19;
          let _684_cf18 = _679___mcc_h18;
          let _685_cf17 = _678___mcc_h17;
          let _686_cf16 = _677___mcc_h16;
          (globalState).f3 = _685_cf17;
          let _687_v27;
          let _init23 = ((_688_v0) => function (_689_i6) {
            return (_689_i6).multipliedBy((_688_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_688_v0).length))]);
          })(_619_v0);
          let _nw111 = Array((new BigNumber(6)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw111.length); _i0_23++) {
            _nw111[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _687_v27 = _nw111;
          let _690_v28;
          _690_v28 = _dafny.Set.fromElements(_687_v27, _687_v27, _687_v27);
          let _691_v29;
          _691_v29 = _module.D3.create_DC9();
          let _692_v30;
          _692_v30 = _dafny.Seq.of(_691_v29);
          let _693_v31;
          _693_v31 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(316)), ((_694_v29) => function (_695_i8) {
            return _694_v29;
          })(_691_v29));
          let _696_v32;
          _696_v32 = _dafny.Seq.of(_692_v30, _692_v30, _692_v30);
          let _697_v33;
          let _nw112 = Array((new BigNumber(26)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _692_v30;
          _nw112[(_dafny.ONE).toNumber()] = _692_v30;
          _nw112[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_691_v29);
          _nw112[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_692_v30, _692_v30);
          _nw112[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_692_v30, _692_v30);
          _nw112[(new BigNumber(5)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_692_v30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(54)), ((_698_v29) => function (_699_i7) {
            return _698_v29;
          })(_691_v29)));
          _nw112[(new BigNumber(7)).toNumber()] = (_693_v31);
          _nw112[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_692_v30, _dafny.Seq.of(_691_v29));
          _nw112[(new BigNumber(9)).toNumber()] = (_696_v32)[_module.__default.safeIndex((_this).f25, new BigNumber((_696_v32).length))];
          _nw112[(new BigNumber(10)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(11)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_691_v29);
          _nw112[(new BigNumber(13)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(14)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_692_v30, _692_v30);
          _nw112[(new BigNumber(16)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), ((_700_v29) => function (_701_i9) {
            return _700_v29;
          })(_691_v29)), _692_v30);
          _nw112[(new BigNumber(18)).toNumber()] = _module.__default.fm14(_685_cf17, (_this).f18, new BigNumber(-325), globalState);
          _nw112[(new BigNumber(19)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_691_v29);
          _nw112[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_692_v30, _692_v30);
          _nw112[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_692_v30, _dafny.Seq.of(_691_v29, _module.D3.create_DC9(), _691_v29));
          _nw112[(new BigNumber(23)).toNumber()] = _692_v30;
          _nw112[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(118)), ((_702_v29) => function (_703_i10) {
            return _702_v29;
          })(_691_v29));
          _nw112[(new BigNumber(25)).toNumber()] = _692_v30;
          _697_v33 = _nw112;
          let _index106 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_697_v33).length));
          (_697_v33)[_index106] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), ((_704_v29) => function (_705_i11) {
            return _704_v29;
          })(_691_v29));
          let _index107 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_697_v33).length));
          let _rhs92 = _module.__default.fm13(globalState);
          let _rhs93 = _690_v28;
          let _rhs94 = _dafny.Seq.update(_692_v30, _module.__default.safeIndex(_this.f26, new BigNumber((_692_v30).length)), _691_v29);
          let _lhs81 = _697_v33;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_697_v33).length));
          _625_cf0 = _rhs92;
          _690_v28 = _rhs93;
          _lhs81[_lhs82] = _rhs94;
          _685_cf17 = !((_685_cf17) && (((_625_cf0) ? (!(_685_cf17)) : (_685_cf17))));
          let _706_v34;
          _706_v34 = _dafny.Seq.of((_this).f18);
          let _707_v35;
          _707_v35 = _module.D3.create_DC10(_671_v24, _625_cf0, _686_cf16, _686_cf16, (_this).f25);
          let _708_v36;
          _708_v36 = _dafny.Map.Empty.slice().updateUnsafe((_706_v34)[_module.__default.safeIndex(_this.f26, new BigNumber((_706_v34).length))],_707_v35);
          let _709_v37;
          _709_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_708_v36).length),(_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))]);
          r2 = ((_this).f25).minus(new BigNumber((_709_v37).length));
        } else {
          let _710___mcc_h21 = (_source7).cf15;
          let _711_cf15 = _710___mcc_h21;
          let _712_v38;
          let _nw113 = Array((new BigNumber(7)).toNumber());
          _nw113[(_dafny.ZERO).toNumber()] = (_this).f25;
          _nw113[(_dafny.ONE).toNumber()] = (_this).f25;
          _nw113[(new BigNumber(2)).toNumber()] = (_this).f25;
          _nw113[(new BigNumber(3)).toNumber()] = (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))];
          _nw113[(new BigNumber(4)).toNumber()] = _this.f26;
          _nw113[(new BigNumber(5)).toNumber()] = _this.f26;
          _nw113[(new BigNumber(6)).toNumber()] = new BigNumber(244);
          _712_v38 = _nw113;
          let _713_v39;
          _713_v39 = _module.D0.create_DC2((_this).f18, (_619_v0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_619_v0).length))], _712_v38);
          (globalState).f14 = (_713_v39).dtor_cf6;
          (globalState).f3 = _625_cf0;
          let _714_v40;
          let _nw114 = new _module.C0();
          _nw114.__ctor();
          _714_v40 = _nw114;
          let _715_v41;
          let _nw115 = Array((new BigNumber(13)).toNumber()).fill(false);
          _715_v41 = _nw115;
          let _716_v42;
          _716_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_715_v41);
          let _717_v43;
          _717_v43 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,_716_v42);
          let _718_v44;
          _718_v44 = _module.D5.create_DC12(_716_v42);
          _717_v43 = (_717_v43).update((_this).f25, (_718_v44).dtor_cf22);
        }
      } else if (_source5.is_DC1) {
        let _719___mcc_h1 = (_source5).cf1;
        let _720___mcc_h2 = (_source5).cf2;
        let _721___mcc_h3 = (_source5).cf3;
        let _722___mcc_h4 = (_source5).cf4;
        let _723_cf4 = _722___mcc_h4;
        let _724_cf3 = _721___mcc_h3;
        let _725_cf2 = _720___mcc_h2;
        let _726_cf1 = _719___mcc_h1;
        _726_cf1 = _726_cf1;
        let _727_v45;
        let _init24 = function (_728_i12) {
          return (_this).f18;
        };
        let _nw116 = Array((new BigNumber(15)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw116.length); _i0_24++) {
          _nw116[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _727_v45 = _nw116;
        let _index108 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
        (_727_v45)[_index108] = (_this).f18;
        let _729_v46;
        _729_v46 = _dafny.MultiSet.fromElements(_619_v0, _619_v0);
        let _730_v47;
        _730_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_729_v46);
        let _731_v48;
        _731_v48 = _dafny.Seq.of(_729_v46, _729_v46);
        let _732_v49;
        let _nw117 = Array((new BigNumber(20)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = _729_v46;
        _nw117[(_dafny.ONE).toNumber()] = _729_v46;
        _nw117[(new BigNumber(2)).toNumber()] = (_729_v46).update(_619_v0, _module.__default.abs(_this.f26));
        _nw117[(new BigNumber(3)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(4)).toNumber()] = (_729_v46).Union(_dafny.MultiSet.fromElements(_619_v0, _619_v0, _619_v0));
        _nw117[(new BigNumber(5)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(6)).toNumber()] = (((_730_v47).contains((_this).f18)) ? ((_730_v47).get((_this).f18)) : (_729_v46));
        _nw117[(new BigNumber(7)).toNumber()] = (_729_v46).Difference(_729_v46);
        _nw117[(new BigNumber(8)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(9)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(10)).toNumber()] = (_729_v46).Intersect(_729_v46);
        _nw117[(new BigNumber(11)).toNumber()] = (_731_v48)[_module.__default.safeIndex(_724_cf3, new BigNumber((_731_v48).length))];
        _nw117[(new BigNumber(12)).toNumber()] = (_729_v46).update(_619_v0, _module.__default.abs(_module.__default.fm4(_this.f26, globalState)));
        _nw117[(new BigNumber(13)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_619_v0, _619_v0);
        _nw117[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_619_v0)).Union(_729_v46);
        _nw117[(new BigNumber(16)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(17)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(18)).toNumber()] = _729_v46;
        _nw117[(new BigNumber(19)).toNumber()] = _729_v46;
        _732_v49 = _nw117;
        let _index109 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_732_v49).length));
        (_732_v49)[_index109] = _729_v46;
        let _733_v50;
        _733_v50 = _module.D0.create_DC0((_this).f18);
        let _734_v51;
        _734_v51 = _dafny.Seq.of((_this).f18);
        let _index110 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
        let _index111 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_732_v49).length));
        let _rhs95 = _723_cf4;
        let _rhs96 = (_725_cf2).isLessThan(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f18), _module.__default.safeIndex(_724_cf3, new BigNumber((_dafny.Seq.of((_this).f18)).length)), (_733_v50).dtor_cf0)).length));
        let _rhs97 = _723_cf4;
        let _rhs98 = (_734_v51)[_module.__default.safeIndex((((_729_v46).contains(_619_v0)) ? ((_729_v46).get(_619_v0)) : (_724_cf3)), new BigNumber((_734_v51).length))];
        let _rhs99 = (_729_v46).Union(_dafny.MultiSet.fromElements(_619_v0));
        let _lhs83 = _727_v45;
        let _lhs84 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
        let _lhs85 = globalState;
        let _lhs86 = _732_v49;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_732_v49).length));
        _723_cf4 = _rhs95;
        _lhs83[_lhs84] = _rhs96;
        _723_cf4 = _rhs97;
        _lhs85.f3 = _rhs98;
        _lhs86[_lhs87] = _rhs99;
        if ((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]) {
          let _index112 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_619_v0).length));
          (_619_v0)[_index112] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_735_cf4) => function (_736_i13) {
            return _735_cf4;
          })(_723_cf4)), _726_cf1)).length);
          let _737_v52;
          _737_v52 = _dafny.Set.fromElements((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]);
          let _738_v53;
          _738_v53 = _dafny.Map.Empty.slice().updateUnsafe(_726_cf1,_724_cf3);
          let _739_v54;
          _739_v54 = _dafny.Seq.of(_725_cf2, _725_cf2);
          let _index113 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
          let _index114 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_619_v0).length));
          let _rhs100 = !(!((_this).f25).isEqualTo((new BigNumber((_737_v52).length)).multipliedBy((_this).f25)));
          let _rhs101 = false;
          let _rhs102 = !(new BigNumber((_738_v53).length)).isEqualTo((new BigNumber((_739_v54).length)).minus(new BigNumber(536)));
          let _rhs103 = ((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]) && (_module.__default.fm13(globalState));
          let _rhs104 = new BigNumber(647);
          let _lhs88 = globalState;
          let _lhs89 = _727_v45;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
          let _lhs91 = globalState;
          let _lhs92 = _727_v45;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length));
          let _lhs94 = _619_v0;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_619_v0).length));
          _lhs88.f3 = _rhs100;
          _lhs89[_lhs90] = _rhs101;
          _lhs91.f9 = _rhs102;
          _lhs92[_lhs93] = _rhs103;
          _lhs94[_lhs95] = _rhs104;
          let _740_v55;
          _740_v55 = _dafny.MultiSet.fromElements(false, (_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]);
          let _741_v56;
          _741_v56 = _module.D2.create_DC7(_module.D2.create_DC5(_740_v55));
          let _742_v57;
          _742_v57 = _module.D2.create_DC7(_741_v56);
          _742_v57 = _742_v57;
          let _743_v58;
          let _init25 = function (_744_i14) {
            return (_744_i14).minus((_this).f25);
          };
          let _nw118 = Array((new BigNumber(14)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw118.length); _i0_25++) {
            _nw118[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _743_v58 = _nw118;
          let _745_v59;
          _745_v59 = _module.D0.create_DC2((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))], _725_cf2, _743_v58);
          let _rhs105 = (new BigNumber(615)).isLessThanOrEqualTo((_745_v59).dtor_cf6);
          let _rhs106 = (_this).f18;
          let _lhs96 = globalState;
          let _lhs97 = globalState;
          _lhs96.f9 = _rhs105;
          _lhs97.f3 = _rhs106;
          (globalState).f3 = (_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))];
          let _746_v60;
          _746_v60 = _module.D2.create_DC5(_740_v55);
          let _747_v61;
          _747_v61 = _dafny.Set.fromElements(_746_v60, _module.__default.fm15((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))], _726_cf1, _726_cf1, globalState));
          (globalState).f3 = !((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]) || ((_747_v61).IsSubsetOf(_747_v61));
        } else {
          (globalState).f9 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(824)), ((_748_cf4) => function (_749_i15) {
            return _748_cf4;
          })(_723_cf4))).length)).isLessThan(_725_cf2);
          let _750_v62;
          let _init26 = ((_751_v51) => function (_752_i16) {
            return (_751_v51);
          })(_734_v51);
          let _nw119 = Array((new BigNumber(6)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw119.length); _i0_26++) {
            _nw119[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _750_v62 = _nw119;
          let _index116 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_750_v62).length));
          (_750_v62)[_index116] = _dafny.Seq.Concat(_dafny.Seq.update(_734_v51, _module.__default.safeIndex(_this.f26, new BigNumber((_734_v51).length)), (_this).f18), _module.__default.fm16(globalState));
          let _753_v63;
          _753_v63 = _dafny.Seq.of(_725_cf2);
          let _index117 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_750_v62).length));
          (_750_v62)[_index117] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_734_v51, _dafny.Seq.Concat(_734_v51, _dafny.Seq.update(_dafny.Seq.of(false), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.of(false)).length)), (_this).f18))), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_753_v63, _753_v63)).length), new BigNumber((_dafny.Seq.Concat(_734_v51, _dafny.Seq.Concat(_734_v51, _dafny.Seq.update(_dafny.Seq.of(false), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.of(false)).length)), (_this).f18)))).length)), false), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_734_v51, _dafny.Seq.Concat(_734_v51, _dafny.Seq.update(_dafny.Seq.of(false), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.of(false)).length)), (_this).f18))), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_753_v63, _753_v63)).length), new BigNumber((_dafny.Seq.Concat(_734_v51, _dafny.Seq.Concat(_734_v51, _dafny.Seq.update(_dafny.Seq.of(false), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.of(false)).length)), (_this).f18)))).length)), false)).length)), (_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]);
          let _754_v64;
          _754_v64 = _dafny.Set.fromElements(_module.__default.fm17(globalState));
          let _755_v65;
          _755_v65 = _dafny.Seq.of(!((_this).f18), (_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]);
          _754_v64 = (((_this).f18) ? (_754_v64) : (_dafny.Set.fromElements(_755_v65, _755_v65)));
          let _rhs107 = !(_module.__default.fm13(globalState)) || ((_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))]);
          let _rhs108 = _this.f26;
          let _rhs109 = (_725_cf2).minus(_724_cf3);
          let _lhs98 = globalState;
          let _lhs99 = globalState;
          let _lhs100 = globalState;
          _lhs98.f9 = _rhs107;
          _lhs99.f14 = _rhs108;
          _lhs100.f14 = _rhs109;
          (globalState).f14 = _724_cf3;
        }
        (globalState).f3 = (_727_v45)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_727_v45).length))];
      } else {
        let _756___mcc_h5 = (_source5).cf5;
        let _757___mcc_h6 = (_source5).cf6;
        let _758___mcc_h7 = (_source5).cf7;
        let _759_cf7 = _758___mcc_h7;
        let _760_cf6 = _757___mcc_h6;
        let _761_cf5 = _756___mcc_h5;
        let _762_v66;
        _762_v66 = _dafny.Set.fromElements((_this).f18);
        let _763_v67;
        _763_v67 = _dafny.Seq.of(new BigNumber((_762_v66).length), (_this).f25, _760_cf6, _760_cf6, _760_cf6);
        let _764_v68;
        _764_v68 = _dafny.Set.fromElements((_763_v67)[_module.__default.safeIndex(_760_cf6, new BigNumber((_763_v67).length))]);
        (globalState).f9 = (_764_v68).IsProperSubsetOf(_764_v68);
        let _765_v69;
        _765_v69 = new _dafny.CodePoint('q'.codePointAt(0));
        let _index118 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_759_cf7).length));
        (_759_cf7)[_index118] = ((_this).f25).multipliedBy((_dafny.ZERO).minus(_760_cf6));
        let _766_v70;
        _766_v70 = _dafny.MultiSet.fromElements((_this).f18);
        let _767_v71;
        _767_v71 = _module.D2.create_DC5(_766_v70);
        let _768_v72;
        _768_v72 = _module.D2.create_DC5((_767_v71).dtor_cf12);
        let _769_v73;
        _769_v73 = _dafny.Seq.of((_this).f18, _761_cf5);
        let _770_v74;
        _770_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_769_v73).length),(_this).f25);
        let _771_v75;
        _771_v75 = _dafny.Seq.UnicodeFromString("ayplhylwi");
        let _index119 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_759_cf7).length));
        let _rhs110 = new _dafny.CodePoint('c'.codePointAt(0));
        let _rhs111 = new BigNumber(-472);
        let _rhs112 = _768_v72;
        let _rhs113 = _module.__default.safeModuloInt(_this.f26, _760_cf6);
        let _rhs114 = (new BigNumber((_770_v74).length)).minus(new BigNumber((_dafny.Seq.Concat(_771_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_772_i17) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }))).length));
        let _lhs101 = _759_cf7;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_759_cf7).length));
        let _lhs103 = _this;
        _765_v69 = _rhs110;
        _lhs101[_lhs102] = _rhs111;
        _767_v71 = _rhs112;
        _lhs103.f26 = _rhs113;
        r2 = _rhs114;
        let _index120 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_759_cf7).length));
        (_759_cf7)[_index120] = _this.f26;
        let _773_v76;
        let _nw120 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _773_v76 = _nw120;
        let _774_v77;
        _774_v77 = _dafny.Seq.of(_759_cf7);
        let _index121 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_773_v76).length));
        (_773_v76)[_index121] = ((_761_cf5) ? (_774_v77) : (_774_v77));
        let _index122 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_773_v76).length));
        (_773_v76)[_index122] = _dafny.Seq.Concat(_774_v77, _dafny.Seq.update(_774_v77, _module.__default.safeIndex((_759_cf7)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_759_cf7).length))], new BigNumber((_774_v77).length)), _619_v0));
      }
      let _775_v78;
      _775_v78 = _dafny.Set.fromElements((_this).f18);
      let _776_v79;
      _776_v79 = _dafny.Seq.of(_this.f26, _this.f26, new BigNumber(224));
      if (_dafny.Seq.contains(_dafny.Seq.Concat(_776_v79, _776_v79), (new BigNumber((_775_v78).length)).multipliedBy(_this.f26))) {
        (globalState).f9 = (_this).f18;
        let _777_v80;
        _777_v80 = _dafny.Seq.of((_this).f18);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_777_v80, _777_v80), _777_v80)) {
          let _778_v81;
          let _init27 = function (_779_i18) {
            return (_this).f18;
          };
          let _nw121 = Array((new BigNumber(27)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw121.length); _i0_27++) {
            _nw121[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _778_v81 = _nw121;
          let _780_v82;
          let _nw122 = new _module.C3();
          _nw122.__ctor(_778_v81, _dafny.Set.fromElements(_778_v81, _778_v81), (_this).f18);
          _780_v82 = _nw122;
          (_this).f26 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_780_v82)).length))));
          let _781_v83;
          _781_v83 = new _dafny.CodePoint('y'.codePointAt(0));
          let _782_v84;
          _782_v84 = _dafny.Seq.of(_781_v83, _781_v83);
          (globalState).f14 = (new BigNumber(((_dafny.MultiSet.fromElements(_781_v83)).Union(_dafny.MultiSet.FromArray(_782_v84))).cardinality())).plus((_this).f25);
          let _783_v85;
          _783_v85 = _dafny.Set.fromElements(new BigNumber(854), _this.f26);
          _783_v85 = _783_v85;
          let _784_v86;
          _784_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,new BigNumber(-870));
          let _index123 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_778_v81).length));
          (_778_v81)[_index123] = !((((_784_v86).contains(new BigNumber((_dafny.Seq.UnicodeFromString("vcaginhv")).length))) ? ((_784_v86).get(new BigNumber((_dafny.Seq.UnicodeFromString("vcaginhv")).length))) : (new BigNumber(-392)))).isEqualTo((_dafny.ZERO).minus(new BigNumber((_776_v79).length)));
          let _index124 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_778_v81).length));
          (_778_v81)[_index124] = ((_this).f18) === ((_this).f18);
          let _785_v88;
          let _nw123 = new _module.C2();
          _nw123.__ctor(_780_v82.f17, !((_this).f18));
          _785_v88 = _nw123;
          let _786_v89;
          _786_v89 = _dafny.Map.Empty.slice().updateUnsafe((_778_v81)[_module.__default.safeIndex(new BigNumber(735), new BigNumber((_778_v81).length))],_785_v88);
          let _787_v90;
          let _788_v91;
          let _789_v92;
          let _out44;
          let _out45;
          let _out46;
          let _outcollector12 = (_780_v82).m0((_this).f25, (function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-119)), function (_790_i19) {
              return _this.f26;
            })).Elements) {
              let _791_v87 = _compr_41;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-119)), function (_790_i19) {
                return _this.f26;
              }), _791_v87)) {
                _coll41.add((_791_v87).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(33)), function (_792_i20) {
                  return _module.D3.create_DC9();
                }))).length))));
              }
            }
            return _coll41;
          }()).IsSubsetOf(_783_v85), new BigNumber(((((_778_v81)[_module.__default.safeIndex(new BigNumber(735), new BigNumber((_778_v81).length))]) ? (_786_v89) : (_786_v89))).length), globalState);
          _out44 = _outcollector12[0];
          _out45 = _outcollector12[1];
          _out46 = _outcollector12[2];
          _787_v90 = _out44;
          _788_v91 = _out45;
          _789_v92 = _out46;
        } else {
          let _793_v93;
          let _nw124 = new _module.C2();
          _nw124.__ctor(_this.f17, (_this).f18);
          _793_v93 = _nw124;
          let _794_v94;
          _794_v94 = _module.D3.create_DC9();
          let _795_v95;
          _795_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_794_v94);
          (globalState).f3 = !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_775_v78).length),_794_v94)).equals(_795_v95);
          let _796_v96;
          _796_v96 = _dafny.Seq.UnicodeFromString("ntmfjd");
          let _797_v97;
          _797_v97 = _dafny.Seq.of(_dafny.Seq.Concat(_796_v96, _796_v96));
          r2 = new BigNumber(((_797_v97)[_module.__default.safeIndex((_this).f25, new BigNumber((_797_v97).length))]).length);
          let _798_v98;
          let _init28 = function (_799_i21) {
            return (_this).f18;
          };
          let _nw125 = Array((new BigNumber(27)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw125.length); _i0_28++) {
            _nw125[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _798_v98 = _nw125;
          let _800_v99;
          let _nw126 = new _module.C3();
          _nw126.__ctor(_798_v98, _this.f17, (_this).f18);
          _800_v99 = _nw126;
          let _801_v100;
          _801_v100 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),_this.f26);
          _801_v100 = (_801_v100).update(_this.f26, (_this).f25);
        }
        let _802_v101;
        let _init29 = function (_803_i22) {
          return _dafny.Seq.UnicodeFromString("wllppav");
        };
        let _nw127 = Array((new BigNumber(2)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw127.length); _i0_29++) {
          _nw127[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _802_v101 = _nw127;
        let _804_v102;
        _804_v102 = _dafny.Seq.UnicodeFromString("gwvqrarpx");
        let _index125 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_802_v101).length));
        (_802_v101)[_index125] = _804_v102;
        let _index126 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_802_v101).length));
        (_802_v101)[_index126] = _804_v102;
        _776_v79 = (((_this).f18) ? (_776_v79) : (_776_v79));
        let _805_v103;
        _805_v103 = new _dafny.CodePoint('b'.codePointAt(0));
        _805_v103 = _805_v103;
      } else {
        (globalState).f3 = (_this).f18;
        let _806_v104;
        let _nw128 = new _module.C2();
        _nw128.__ctor(_this.f17, !((_this).f18));
        _806_v104 = _nw128;
        _806_v104 = _806_v104;
        let _807_v105;
        let _nw129 = Array((new BigNumber(17)).toNumber()).fill(false);
        _807_v105 = _nw129;
        let _808_v106;
        let _nw130 = new _module.C3();
        _nw130.__ctor(_807_v105, _this.f17, (_this).f18);
        _808_v106 = _nw130;
        (globalState).f3 = (_this).f18;
        (globalState).f9 = _module.__default.fm13(globalState);
      }
      let _809_v107;
      _809_v107 = _dafny.Set.fromElements((_this).f25);
      let _index127 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length));
      (_619_v0)[_index127] = (_776_v79)[_module.__default.safeIndex(new BigNumber((_809_v107).length), new BigNumber((_776_v79).length))];
      let _index128 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length));
      (_619_v0)[_index128] = _this.f26;
      let _810_i23;
      _810_i23 = _dafny.ZERO;
      L4: {
        while ((_this).f18) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_810_i23)) {
              break L4;
            }
            _810_i23 = (_810_i23).plus(_dafny.ONE);
            let _811_v108;
            _811_v108 = _dafny.Map.Empty.slice().updateUnsafe(_776_v79,_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("gsdlppl"), _dafny.Seq.UnicodeFromString("ghnl")));
            let _812_v109;
            _812_v109 = _dafny.Seq.of(!((_this).f18), (_this).f18);
            _811_v108 = (_811_v108).update(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-581)), _dafny.Seq.update(_776_v79, _module.__default.safeIndex((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))], new BigNumber((_776_v79).length)), new BigNumber((_812_v109).length))), (_this).f18);
            let _813_v110;
            _813_v110 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f25),(_this).f18);
            let _814_v111;
            _814_v111 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))]);
            _813_v110 = (_813_v110).update((_814_v111).Merge(_814_v111), (_this).f18);
            let _815_v112;
            _815_v112 = _dafny.Seq.UnicodeFromString("hn");
            (_this).f26 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_this).f25, _this.f26), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(972)), function (_816_i24) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            }), _815_v112)).length));
            let _817_v113;
            _817_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(10),_815_v112);
            let _818_v114;
            _818_v114 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_815_v112).length), (_this).f25));
            let _819_v115;
            _819_v115 = new _dafny.CodePoint('c'.codePointAt(0));
            _815_v112 = _dafny.Seq.update(_dafny.Seq.Concat(_815_v112, _dafny.Seq.Concat(_815_v112, (((_817_v113).contains((_776_v79)[_module.__default.safeIndex(new BigNumber((_818_v114).length), new BigNumber((_776_v79).length))])) ? ((_817_v113).get((_776_v79)[_module.__default.safeIndex(new BigNumber((_818_v114).length), new BigNumber((_776_v79).length))])) : (_dafny.Seq.UnicodeFromString("ptbvpe"))))), _module.__default.safeIndex((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))], new BigNumber((_dafny.Seq.Concat(_815_v112, _dafny.Seq.Concat(_815_v112, (((_817_v113).contains((_776_v79)[_module.__default.safeIndex(new BigNumber((_818_v114).length), new BigNumber((_776_v79).length))])) ? ((_817_v113).get((_776_v79)[_module.__default.safeIndex(new BigNumber((_818_v114).length), new BigNumber((_776_v79).length))])) : (_dafny.Seq.UnicodeFromString("ptbvpe")))))).length)), (((_this).f18) ? (new _dafny.CodePoint('p'.codePointAt(0))) : (_819_v115)));
          }
        }
      }
      let _820_i25;
      _820_i25 = _dafny.ZERO;
      L5: {
        while (((_this).f25).isLessThan(_this.f26)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_820_i25)) {
              break L5;
            }
            _820_i25 = (_820_i25).plus(_dafny.ONE);
            r2 = (_dafny.ZERO).minus(_module.__default.fm4(_this.f26, globalState));
            let _821_v116;
            let _nw131 = Array((new BigNumber(28)).toNumber()).fill([]);
            _821_v116 = _nw131;
            let _822_v117;
            let _init30 = function (_823_i26) {
              return _dafny.Seq.UnicodeFromString("ekeegfkgk");
            };
            let _nw132 = Array((new BigNumber(21)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw132.length); _i0_30++) {
              _nw132[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _822_v117 = _nw132;
            let _index129 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_821_v116).length));
            (_821_v116)[_index129] = _822_v117;
            let _index130 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_821_v116).length));
            (_821_v116)[_index130] = _822_v117;
            let _824_v118;
            let _nw133 = new _module.C0();
            _nw133.__ctor();
            _824_v118 = _nw133;
            let _index131 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length));
            (_619_v0)[_index131] = _this.f26;
          }
        }
      }
      let _hi2 = (_this).f25;
      for (let _825_i27 = _this.f26; _825_i27.isLessThan(_hi2); _825_i27 = _825_i27.plus(_dafny.ONE)) {
        let _826_v119;
        _826_v119 = new _dafny.CodePoint('d'.codePointAt(0));
        _826_v119 = _826_v119;
        let _827_v120;
        _827_v120 = _dafny.Seq.of(_826_v119);
        let _828_v121;
        _828_v121 = _dafny.Map.Empty.slice().updateUnsafe((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))],_this.f26);
        let _829_v122;
        _829_v122 = _module.D3.create_DC10(_827_v120, (_this).f18, _dafny.Seq.UnicodeFromString("yivqbw"), _827_v120, new BigNumber((_828_v121).length));
        (globalState).f9 = !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_829_v122,(_this).f18)).length)).isEqualTo(_825_i27);
        let _830_v123;
        _830_v123 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-912)), ((_831_v119) => function (_832_i28) {
          return _831_v119;
        })(_826_v119)));
        _830_v123 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.Create(_module.__default.abs(new BigNumber(103)), ((_833_v119) => function (_834_i29) {
          return _833_v119;
        })(_826_v119)));
        if ((_this).f18) {
          let _835_v124;
          _835_v124 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f18, (_this).f18), _module.__default.safeIndex(_this.f26, new BigNumber((_dafny.Seq.of((_this).f18, (_this).f18)).length)), false)).length)).isLessThan(_this.f26),(_this).f25);
          _835_v124 = (_835_v124).update(!((_this).f18), (_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))]);
          let _836_v125;
          let _nw134 = new _module.C0();
          _nw134.__ctor();
          _836_v125 = _nw134;
          let _837_v126;
          let _nw135 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
          _837_v126 = _nw135;
          let _index132 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_837_v126).length));
          (_837_v126)[_index132] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_838_i30) {
            return _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(262), new BigNumber((_dafny.Seq.of(true)).length), _this.f26, new BigNumber((_dafny.MultiSet.fromElements((_this).f18)).cardinality()))).cardinality()), new BigNumber(715));
          }));
          let _index133 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_837_v126).length));
          (_837_v126)[_index133] = (_dafny.MultiSet.fromElements(_776_v79)).Difference(_dafny.MultiSet.fromElements(_776_v79, _776_v79, _776_v79));
          let _839_v127;
          let _nw136 = new _module.C0();
          _nw136.__ctor();
          _839_v127 = _nw136;
          (globalState).f9 = (_module.__default.safeDivisionInt(_this.f26, (_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))])).isLessThanOrEqualTo((new BigNumber(-298)).minus(new BigNumber((_775_v78).length)));
        } else {
          (globalState).f9 = !((_this).f18);
          let _840_v128;
          let _nw137 = new _module.C1();
          _nw137.__ctor();
          _840_v128 = _nw137;
          let _841_v129;
          _841_v129 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_827_v120).length),(_this).f18);
          let _842_v130;
          _842_v130 = _dafny.Map.Empty.slice().updateUnsafe(_825_i27,_841_v129);
          let _843_v131;
          _843_v131 = _dafny.Map.Empty.slice().updateUnsafe(_842_v130,(_this).f25);
          (globalState).f14 = (((_843_v131).contains(_842_v130)) ? ((_843_v131).get(_842_v130)) : (new BigNumber(-449)));
          let _844_v132;
          let _nw138 = Array((new BigNumber(6)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _619_v0;
          _nw138[(_dafny.ONE).toNumber()] = _619_v0;
          _nw138[(new BigNumber(2)).toNumber()] = _619_v0;
          _nw138[(new BigNumber(3)).toNumber()] = _619_v0;
          _nw138[(new BigNumber(4)).toNumber()] = _619_v0;
          _nw138[(new BigNumber(5)).toNumber()] = _619_v0;
          _844_v132 = _nw138;
          let _845_v133;
          _845_v133 = _dafny.Map.Empty.slice().updateUnsafe(_844_v132,_this.f26);
          let _846_v134;
          _846_v134 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_844_v132);
          _845_v133 = (_845_v133).update((((_846_v134).contains((_840_v128).fm5(_this.f26, _776_v79, globalState))) ? ((_846_v134).get((_840_v128).fm5(_this.f26, _776_v79, globalState))) : (_844_v132)), (_dafny.ZERO).minus((((_this).f18) ? (_825_i27) : (_this.f26))));
          let _847_v135;
          _847_v135 = _module.D0.create_DC2((_this).f18, (_this).f25, _619_v0);
          let _848_v136;
          _848_v136 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), ((_849_v119) => function (_850_i31) {
            return _849_v119;
          })(_826_v119)));
          let _851_v137;
          _851_v137 = _dafny.MultiSet.fromElements(_825_i27);
          let _852_v138;
          let _nw139 = Array((new BigNumber(14)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = (((_this).f18) ? ((_this).f18) : ((_this).f18));
          _nw139[(_dafny.ONE).toNumber()] = !((_this).f18) || ((_this).f18);
          _nw139[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(3)).toNumber()] = ((_dafny.MultiSet.fromElements(_825_i27)).update((_this).f25, _module.__default.abs(_this.f26))).IsSubsetOf(_module.__default.fm24((_847_v135).dtor_cf5, (_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))], globalState));
          _nw139[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(6)).toNumber()] = (new BigNumber((_827_v120).length)).isLessThan((((_848_v136).contains(_827_v120)) ? ((_848_v136).get(_827_v120)) : ((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))])));
          _nw139[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(8)).toNumber()] = ((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))]).isLessThan((_this).f25);
          _nw139[(new BigNumber(9)).toNumber()] = (_840_v128).fm5(_this.f26, _776_v79, globalState);
          _nw139[(new BigNumber(10)).toNumber()] = !((_this).f18);
          _nw139[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(12)).toNumber()] = (_this).f18;
          _nw139[(new BigNumber(13)).toNumber()] = ((((_851_v137).contains(_this.f26)) ? ((_851_v137).get(_this.f26)) : (_825_i27))).isLessThanOrEqualTo(new BigNumber((_827_v120).length));
          _852_v138 = _nw139;
          let _index134 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_852_v138).length));
          (_852_v138)[_index134] = (_this).f18;
          let _index135 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_852_v138).length));
          let _rhs115 = (_this).f18;
          let _rhs116 = _826_v119;
          let _lhs104 = _852_v138;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_852_v138).length));
          _lhs104[_lhs105] = _rhs115;
          _826_v119 = _rhs116;
        }
      }
      r0 = ((_this).f25).isEqualTo(new BigNumber(730));
      let _853_v139;
      let _nw140 = Array((new BigNumber(11)).toNumber()).fill(false);
      _853_v139 = _nw140;
      let _854_v140;
      _854_v140 = _dafny.Map.Empty.slice().updateUnsafe(_853_v139,_this.f26);
      r1 = _854_v140;
      let _855_v141;
      _855_v141 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_809_v107);
      r2 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))], (_619_v0)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_619_v0).length))]), new BigNumber(((((_855_v141).contains(!((_this).f18))) ? ((_855_v141).get(!((_this).f18))) : (_809_v107))).length));
      return [r0, r1, r2];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f21 = _dafny.Map.Empty;
      this._f20 = _dafny.ZERO;
      this._f18 = false;
      this._f29 = _dafny.Seq.UnicodeFromString("");
      this._f30 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f29, f30, f20, f21, f19, f17, f18) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!((_this).f30), (_this).f30)).length), new BigNumber((_dafny.Seq.of((_this).f20, (_this).f20)).length), (_this).f20, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.UnicodeFromString("lbbbe"))).length), (_this).f20)).cardinality())).isEqualTo((((_this).f18) ? ((_this).f20) : ((_this).f20)));
    };
    fm0(p0, p1, globalState) {
      let _this = this;
      let _source8 = _module.D1.create_DC3(_dafny.MultiSet.fromElements((_this).f20));
      if (_source8.is_DC4) {
        let _856___mcc_h0 = (_source8).cf9;
        let _857___mcc_h1 = (_source8).cf10;
        let _858___mcc_h2 = (_source8).cf11;
        let _859_cf11 = _858___mcc_h2;
        let _860_cf10 = _857___mcc_h1;
        let _861_cf9 = _856___mcc_h0;
        return new BigNumber(902);
      } else {
        let _862___mcc_h3 = (_source8).cf8;
        let _863_cf8 = _862___mcc_h3;
        return (_dafny.ZERO).minus((_this).f20);
      }
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      r1 = (_this).f18;
      let _864_v0;
      let _nw141 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _864_v0 = _nw141;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_864_v0).length))) {
        let _865_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_865_i0)) && ((_865_i0).isLessThan(new BigNumber((_864_v0).length))))) {
          (_864_v0)[(_865_i0)] = (_865_i0).multipliedBy(((_this).f20).minus(p0));
        }
      }
      let _866_v1;
      _866_v1 = _dafny.Set.fromElements(true, p1);
      let _867_i1;
      _867_i1 = _dafny.ZERO;
      L6: {
        while ((new BigNumber((_866_v1).length)).isLessThan((_this).f20)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_867_i1)) {
              break L6;
            }
            _867_i1 = (_867_i1).plus(_dafny.ONE);
            let _arr6 = _this.f19;
            let _index136 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_this.f19).length));
            _arr6[_index136] = (_this).f18;
            let _868_v2;
            _868_v2 = new _dafny.CodePoint('v'.codePointAt(0));
            let _arr7 = _this.f19;
            let _index137 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_this.f19).length));
            _arr7[_index137] = !_dafny.Seq.contains(_dafny.Seq.of((_this).f18, p1, !(p1)), false);
            let _869_v3;
            _869_v3 = _dafny.Seq.of((_this).f18);
            let _arr8 = _this.f19;
            let _index138 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_this.f19).length));
            let _arr9 = _this.f19;
            let _index139 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_this.f19).length));
            let _rhs117 = new BigNumber(((_this).f29).length);
            let _rhs118 = ((false) ? ((_this).f30) : (true));
            let _rhs119 = _868_v2;
            let _rhs120 = !(_dafny.Seq.IsProperPrefixOf(_869_v3, _869_v3));
            let _lhs106 = globalState;
            let _lhs107 = _this.f19;
            let _lhs108 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_this.f19).length));
            let _lhs109 = _this.f19;
            let _lhs110 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_this.f19).length));
            _lhs106.f14 = _rhs117;
            _lhs107[_lhs108] = _rhs118;
            _868_v2 = _rhs119;
            _lhs109[_lhs110] = _rhs120;
            let _870_v4;
            let _nw142 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
            _870_v4 = _nw142;
            let _871_v5;
            _871_v5 = _dafny.Map.Empty.slice().updateUnsafe(_868_v2,_870_v4);
            let _872_v6;
            _872_v6 = _dafny.Seq.of(_870_v4, _870_v4);
            let _873_v7;
            _873_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_872_v6)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_872_v6).length))]);
            _871_v5 = (_871_v5).update(new _dafny.CodePoint('g'.codePointAt(0)), (((_873_v7).contains((_this).f29)) ? ((_873_v7).get((_this).f29)) : (_870_v4)));
            let _874_v8;
            _874_v8 = _dafny.Seq.of(p0);
            let _875_v9;
            _875_v9 = _module.D1.create_DC3(_dafny.MultiSet.FromArray(_874_v8));
            _875_v9 = _module.D1.create_DC3(_dafny.MultiSet.fromElements((_this).f20));
            let _876_v10;
            let _init31 = function (_877_i2) {
              return (_this).f18;
            };
            let _nw143 = Array((new BigNumber(28)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw143.length); _i0_31++) {
              _nw143[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _876_v10 = _nw143;
            _876_v10 = _876_v10;
          }
        }
      }
      let _878_v11;
      let _nw144 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
      _878_v11 = _nw144;
      let _879_v12;
      _879_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f18);
      let _880_v13;
      _880_v13 = _dafny.Set.fromElements(_879_v12, _879_v12);
      let _index140 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_878_v11).length));
      (_878_v11)[_index140] = _880_v13;
      let _881_v14;
      _881_v14 = _dafny.Set.fromElements((_this).f29);
      let _index141 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_878_v11).length));
      (_878_v11)[_index141] = (((_881_v14).IsProperSubsetOf(_881_v14)) ? (_880_v13) : (_880_v13));
      let _arr10 = _this.f19;
      let _index142 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length));
      _arr10[_index142] = (_this).f30;
      let _882_v15;
      _882_v15 = _dafny.Seq.of((_this).f30);
      let _883_v16;
      _883_v16 = _882_v15;
      let _arr11 = _this.f19;
      let _index143 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length));
      _arr11[_index143] = !_dafny.areEqual(_dafny.Seq.Concat((_883_v16), _882_v15), (((_this).f30) ? (_882_v15) : (_882_v15)));
      let _884_v17;
      _884_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), function (_885_i4) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length),(_dafny.ZERO).minus(new BigNumber(((_this).f29).length)));
      let _886_v18;
      _886_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f18);
      let _hi3 = (_this).f20;
      for (let _887_i3 = (((_884_v17).contains(new BigNumber((_886_v18).length))) ? ((_884_v17).get(new BigNumber((_886_v18).length))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f29,p1)).length))); _887_i3.isLessThan(_hi3); _887_i3 = _887_i3.plus(_dafny.ONE)) {
        let _888_v19;
        let _nw145 = Array((new BigNumber(26)).toNumber()).fill(_module.D5.Default());
        _888_v19 = _nw145;
        let _889_v20;
        _889_v20 = _module.D3.create_DC10((_this).f29, (_this.f19)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length))], (_this).f29, (_this).f29, new BigNumber((_884_v17).length));
        _888_v19 = ((!((_889_v20).dtor_cf20).isEqualTo(p0)) ? (_888_v19) : (((p1) ? (_888_v19) : (_888_v19))));
        let _890_v22;
        _890_v22 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(500), new BigNumber(409))) {
            let _891_v21 = _compr_42;
            if (((new BigNumber(500)).isLessThanOrEqualTo(_891_v21)) && ((_891_v21).isLessThan(new BigNumber(409)))) {
              _coll42.push([(_891_v21).minus(_887_i3),new BigNumber(147)]);
            }
          }
          return _coll42;
        }(),p2);
        let _892_v23;
        _892_v23 = _dafny.MultiSet.fromElements((_this).fm1(_890_v22, globalState));
        (globalState).f3 = !((_892_v23).contains(!_dafny.Seq.contains(_882_v15, (_this).f30)));
        let _arr12 = _this.f19;
        let _index144 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length));
        let _rhs121 = (_this).f30;
        let _rhs122 = new BigNumber(652);
        let _rhs123 = p0;
        let _lhs111 = _this.f19;
        let _lhs112 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length));
        let _lhs113 = globalState;
        let _lhs114 = globalState;
        _lhs111[_lhs112] = _rhs121;
        _lhs113.f14 = _rhs122;
        _lhs114.f14 = _rhs123;
        let _893_v24;
        _893_v24 = _dafny.Seq.of(_892_v23, _892_v23, _892_v23);
        _892_v23 = (_893_v24)[_module.__default.safeIndex((new BigNumber((_866_v1).length)).plus(new BigNumber((_866_v1).length)), new BigNumber((_893_v24).length))];
      }
      r0 = (_this).f30;
      let _894_v25;
      _894_v25 = _dafny.Set.fromElements((_this).f20);
      let _895_v26;
      let _nw146 = Array((new BigNumber(29)).toNumber()).fill(_module.D11.Default());
      _895_v26 = _nw146;
      let _896_v27;
      _896_v27 = _dafny.Map.Empty.slice().updateUnsafe(_895_v26,_894_v25);
      let _897_v28;
      _897_v28 = _dafny.MultiSet.fromElements((_this).f30, (_this).f30);
      let _898_v29;
      _898_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this.f19)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_this.f19).length))]);
      let _899_v30;
      _899_v30 = _dafny.Seq.of((_this).f20);
      r1 = ((!((_894_v25).IsDisjointFrom((((_896_v27).contains(_895_v26)) ? ((_896_v27).get(_895_v26)) : (_dafny.Set.fromElements(p0, p2, p2, (((_897_v28).contains((_this).f30)) ? ((_897_v28).get((_this).f30)) : (p2)), p0)))))) ? (_dafny.Seq.contains((_this).f29, _module.__default.fm33(p2, (((_898_v29).contains((_this).f30)) ? ((_898_v29).get((_this).f30)) : ((_this).f18)), globalState))) : (_dafny.Seq.contains(_899_v30, new BigNumber(409))));
      r2 = (_this).f20;
      return [r0, r1, r2];
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

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f21 = _dafny.Map.Empty;
      this._f20 = _dafny.ZERO;
      this._f18 = false;
      this._f28 = _dafny.ZERO;
      this._f27 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f27, f28, f20, f21, f19, f17, f18) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return !(true);
    };
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_this).f28;
    };
    fm27(globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm28(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f20;
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi4 = (_this).f28;
      for (let _900_i0 = (_dafny.ZERO).minus(((p1) ? (p2) : (p0))); _900_i0.isLessThan(_hi4); _900_i0 = _900_i0.plus(_dafny.ONE)) {
        let _901_v0;
        _901_v0 = _dafny.Seq.of(p2, p2, p2);
        let _902_v1;
        _902_v1 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f28).isEqualTo(_900_i0)),_dafny.Seq.Concat(_901_v0, _901_v0));
        _902_v1 = (_902_v1).Merge(_902_v1);
        let _903_v2;
        _903_v2 = _901_v0;
        let _904_v3;
        _904_v3 = _dafny.Seq.of((_903_v2), _901_v0);
        let _905_v4;
        _905_v4 = _dafny.Seq.of(p1);
        let _rhs124 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_900_i0, (_this).f28), (new BigNumber((_904_v3).length)).minus(new BigNumber((_905_v4).length)));
        let _rhs125 = new BigNumber((_901_v0).length);
        r2 = _rhs124;
        r2 = _rhs125;
        let _906_v5;
        _906_v5 = _module.D0.create_DC0((_this).f18);
        (globalState).f3 = (_906_v5).dtor_cf0;
        let _907_v6;
        _907_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _908_v7;
        _908_v7 = _dafny.Seq.of(_907_v6, _907_v6, (_907_v6).update(p1, p1));
        let _source9 = _module.D3.create_DC8(_908_v7);
        if (_source9.is_DC9) {
          let _909_v8;
          _909_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(globalState),((_this).f20).multipliedBy(new BigNumber(73)));
          _909_v8 = (_909_v8).update(_dafny.MultiSet.fromElements(true, p1), (_this).f20);
          let _910_v9;
          _910_v9 = _dafny.Map.Empty.slice().updateUnsafe(_905_v4,_module.__default.fm30((_this).f18, globalState));
          (globalState).f14 = (((_this).fm28((_this).f20, new BigNumber(784), _910_v9, globalState)).plus((_this).f20)).plus((_this).f20);
          let _911_v10;
          let _nw147 = new _module.C2();
          _nw147.__ctor(_this.f17, (_this).f18);
          _911_v10 = _nw147;
          let _912_v11;
          _912_v11 = _dafny.Seq.of(_911_v10);
          (globalState).f9 = (((_dafny.ZERO).minus(p2)).multipliedBy((_this).fm28((_this).fm28(p2, new BigNumber((_dafny.Seq.UnicodeFromString("yrjchpu")).length), _910_v9, globalState), (_this).f28, _910_v9, globalState))).isLessThanOrEqualTo(new BigNumber((_912_v11).length));
          let _913_v12;
          _913_v12 = new _dafny.CodePoint('i'.codePointAt(0));
          let _914_v13;
          _914_v13 = _dafny.Seq.UnicodeFromString("jovhhqucn");
          let _915_v14;
          _915_v14 = _dafny.Seq.of(_913_v12, _913_v12, (_914_v13)[_module.__default.safeIndex(new BigNumber(867), new BigNumber((_914_v13).length))]);
          _914_v13 = _914_v13;
        } else if (_source9.is_DC10) {
          let _916___mcc_h0 = (_source9).cf16;
          let _917___mcc_h1 = (_source9).cf17;
          let _918___mcc_h2 = (_source9).cf18;
          let _919___mcc_h3 = (_source9).cf19;
          let _920___mcc_h4 = (_source9).cf20;
          let _921_cf20 = _920___mcc_h4;
          let _922_cf19 = _919___mcc_h3;
          let _923_cf18 = _918___mcc_h2;
          let _924_cf17 = _917___mcc_h1;
          let _925_cf16 = _916___mcc_h0;
          let _926_v15;
          let _nw148 = new _module.C1();
          _nw148.__ctor();
          _926_v15 = _nw148;
          (globalState).f9 = true;
          let _927_v16;
          _927_v16 = _dafny.Set.fromElements((_this).f20);
          let _arr13 = _this.f19;
          let _index145 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_this.f19).length));
          _arr13[_index145] = (_927_v16).IsDisjointFrom(function () {
            let _coll43 = new _dafny.Set();
            for (const _compr_43 of _dafny.IntegerRange(new BigNumber(-267), new BigNumber(467))) {
              let _928_v17 = _compr_43;
              if (((new BigNumber(-267)).isLessThanOrEqualTo(_928_v17)) && ((_928_v17).isLessThan(new BigNumber(467)))) {
                _coll43.add(_module.__default.safeDivisionInt(_928_v17, (_this).f20));
              }
            }
            return _coll43;
          }());
          let _arr14 = _this.f19;
          let _index146 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_this.f19).length));
          _arr14[_index146] = _924_cf17;
          let _929_v18;
          _929_v18 = _dafny.Map.Empty.slice().updateUnsafe(_900_i0,p0);
          let _930_v19;
          _930_v19 = _dafny.Map.Empty.slice().updateUnsafe(_929_v18,p2);
          (globalState).f3 = (_this).fm1(_930_v19, globalState);
        } else {
          let _931___mcc_h5 = (_source9).cf15;
          let _932_cf15 = _931___mcc_h5;
          let _933_v20;
          _933_v20 = new _dafny.CodePoint('s'.codePointAt(0));
          let _934_v21;
          _934_v21 = _dafny.Set.fromElements(_933_v20);
          let _935_v22;
          _935_v22 = _dafny.Seq.of(_module.D3.create_DC9(), _module.__default.fm31(globalState));
          let _936_v23;
          _936_v23 = _dafny.MultiSet.fromElements((_module.D5.create_DC13((_this).f18, (_this).f28, _935_v22)).dtor_cf24);
          let _rhs126 = p1;
          let _rhs127 = (_934_v21).IsSubsetOf(_934_v21);
          let _rhs128 = !(p1);
          let _rhs129 = _module.__default.safeModuloInt((((_936_v23).contains(_900_i0)) ? ((_936_v23).get(_900_i0)) : ((_this).f20)), (_900_i0).multipliedBy(p0));
          let _lhs115 = globalState;
          let _lhs116 = globalState;
          let _lhs117 = globalState;
          _lhs115.f9 = _rhs126;
          r1 = _rhs127;
          _lhs116.f3 = _rhs128;
          _lhs117.f14 = _rhs129;
          let _937_v24;
          let _nw149 = new _module.C3();
          _nw149.__ctor(_this.f19, _this.f17, p1);
          _937_v24 = _nw149;
          let _938_v25;
          _938_v25 = _dafny.Set.fromElements(_937_v24, _937_v24);
          _938_v25 = _938_v25;
          let _939_v26;
          _939_v26 = _module.D9.create_DC19(_module.__default.fm32(globalState));
          r2 = (_this).fm28((_this).f20, ((_this).f28).minus(new BigNumber(144)), (_939_v26).dtor_cf35, globalState);
          let _940_v27;
          _940_v27 = _dafny.Set.fromElements(p0, new BigNumber(-324), p0);
          let _941_v28;
          _941_v28 = _module.D10.create_DC22(_940_v27);
          _940_v27 = (((_941_v28).dtor_cf42).Union(_940_v27)).Intersect(function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_940_v27).Elements) {
              let _942_v29 = _compr_44;
              if ((_940_v27).contains(_942_v29)) {
                _coll44.add((_942_v29).plus(new BigNumber((_dafny.Seq.of(!(true), !(false))).length)));
              }
            }
            return _coll44;
          }());
        }
      }
      let _943_i1;
      _943_i1 = _dafny.ZERO;
      L7: {
        while (p1) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_943_i1)) {
              break L7;
            }
            _943_i1 = (_943_i1).plus(_dafny.ONE);
            let _944_v30;
            let _nw150 = new _module.C1();
            _nw150.__ctor();
            _944_v30 = _nw150;
            let _945_v31;
            _945_v31 = _dafny.Seq.UnicodeFromString("ert");
            let _946_v32;
            _946_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(118),(_this).f28);
            let _947_v33;
            _947_v33 = new _dafny.CodePoint('c'.codePointAt(0));
            let _948_v34;
            let _nw151 = Array((new BigNumber(26)).toNumber());
            _nw151[(_dafny.ZERO).toNumber()] = p0;
            _nw151[(_dafny.ONE).toNumber()] = (_this).f20;
            _nw151[(new BigNumber(2)).toNumber()] = (_this).f28;
            _nw151[(new BigNumber(3)).toNumber()] = (_this).f20;
            _nw151[(new BigNumber(4)).toNumber()] = p2;
            _nw151[(new BigNumber(5)).toNumber()] = new BigNumber(887);
            _nw151[(new BigNumber(6)).toNumber()] = new BigNumber((_945_v31).length);
            _nw151[(new BigNumber(7)).toNumber()] = new BigNumber(((_this).f27).length);
            _nw151[(new BigNumber(8)).toNumber()] = p2;
            _nw151[(new BigNumber(9)).toNumber()] = (_this).f28;
            _nw151[(new BigNumber(10)).toNumber()] = new BigNumber(548);
            _nw151[(new BigNumber(11)).toNumber()] = p2;
            _nw151[(new BigNumber(12)).toNumber()] = p2;
            _nw151[(new BigNumber(13)).toNumber()] = ((((_this).f27).contains((_this).f18)) ? (((_this).f27).get((_this).f18)) : (new BigNumber((_dafny.Seq.UnicodeFromString("pkdr")).length)));
            _nw151[(new BigNumber(14)).toNumber()] = (_this).f28;
            _nw151[(new BigNumber(15)).toNumber()] = new BigNumber((_946_v32).length);
            _nw151[(new BigNumber(16)).toNumber()] = p0;
            _nw151[(new BigNumber(17)).toNumber()] = (_this).f28;
            _nw151[(new BigNumber(18)).toNumber()] = p2;
            _nw151[(new BigNumber(19)).toNumber()] = (_this).f20;
            _nw151[(new BigNumber(20)).toNumber()] = (_this).fm0((_this).f20, _947_v33, globalState);
            _nw151[(new BigNumber(21)).toNumber()] = (_this).f28;
            _nw151[(new BigNumber(22)).toNumber()] = (_this).f20;
            _nw151[(new BigNumber(23)).toNumber()] = (_this).f20;
            _nw151[(new BigNumber(24)).toNumber()] = p0;
            _nw151[(new BigNumber(25)).toNumber()] = (_this).f20;
            _948_v34 = _nw151;
            let _949_v35;
            let _nw152 = Array((new BigNumber(10)).toNumber());
            _nw152[(_dafny.ZERO).toNumber()] = _948_v34;
            _nw152[(_dafny.ONE).toNumber()] = _948_v34;
            _nw152[(new BigNumber(2)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(3)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(4)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(5)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(6)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(7)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(8)).toNumber()] = _948_v34;
            _nw152[(new BigNumber(9)).toNumber()] = ((false) ? (_948_v34) : (_948_v34));
            _949_v35 = _nw152;
            _949_v35 = _949_v35;
            let _950_v36;
            _950_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
            _950_v36 = (_950_v36).update(new BigNumber((_945_v31).length), true);
            let _951_v37;
            _951_v37 = _dafny.Seq.of(true, true);
            let _952_v38;
            _952_v38 = _dafny.Set.fromElements(((_this).fm0(new BigNumber((_951_v37).length), _947_v33, globalState)).minus(new BigNumber((_950_v36).length)));
            let _953_v39;
            _953_v39 = _module.D10.create_DC22(_952_v38);
            let _954_v40;
            _954_v40 = _dafny.Map.Empty.slice().updateUnsafe(false,(_953_v39).dtor_cf42);
            _952_v38 = (((_954_v40).contains((_this).f18)) ? ((_954_v40).get((_this).f18)) : (_952_v38));
          }
        }
      }
      let _955_v41;
      _955_v41 = _dafny.Seq.UnicodeFromString("ky");
      let _956_v42;
      _956_v42 = new _dafny.CodePoint('m'.codePointAt(0));
      let _957_v43;
      _957_v43 = _module.D11.create_DC25(_this.f17);
      let _958_v44;
      let _nw153 = new _module.C4();
      _nw153.__ctor(_module.__default.safeDivisionInt(new BigNumber((_955_v41).length), new BigNumber((_dafny.Seq.update(_955_v41, _module.__default.safeIndex(new BigNumber(-751), new BigNumber((_955_v41).length)), _956_v42)).length)), p2, (_957_v43).dtor_cf47, ((_this).f28).isEqualTo(new BigNumber(-569)));
      _958_v44 = _nw153;
      if (p1) {
        let _959_v45;
        let _nw154 = Array((new BigNumber(3)).toNumber()).fill([]);
        _959_v45 = _nw154;
        let _960_v46;
        let _nw155 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _960_v46 = _nw155;
        let _index147 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_959_v45).length));
        (_959_v45)[_index147] = _960_v46;
        let _961_v47;
        _961_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,_960_v46);
        let _index148 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_959_v45).length));
        (_959_v45)[_index148] = (((_961_v47).contains(p1)) ? ((_961_v47).get(p1)) : (_960_v46));
        let _962_v48;
        _962_v48 = _dafny.MultiSet.fromElements((_this).f20);
        (globalState).f15 = _962_v48;
        (_958_v44).f26 = (_dafny.ZERO).minus(p2);
        _960_v46 = _960_v46;
        let _963_v49;
        _963_v49 = _dafny.Map.Empty.slice().updateUnsafe(_958_v44.f26,_958_v44.f26);
        let _964_v50;
        _964_v50 = _dafny.Map.Empty.slice().updateUnsafe(_963_v49,(_this).f20);
        let _965_v51;
        _965_v51 = _dafny.MultiSet.fromElements((_this).f18);
        let _966_v52;
        let _nw156 = new _module.C5();
        _nw156.__ctor(_955_v41, (_this).fm1((_964_v50).update(_963_v49, (((_965_v51).contains((_this).f18)) ? ((_965_v51).get((_this).f18)) : (_958_v44.f26))), globalState), (_this).f28, _this.f21, ((p1) ? (_this.f19) : (_this.f19)), (_this.f17).Union(_this.f17), (_this).f18);
        _966_v52 = _nw156;
        let _967_v53;
        _967_v53 = _dafny.Map.Empty.slice().updateUnsafe((_966_v52).f18,p1);
        let _arr15 = _966_v52.f19;
        let _index149 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_966_v52.f19).length));
        _arr15[_index149] = (p0).isLessThan(new BigNumber((_967_v53).length));
        let _968_v54;
        _968_v54 = _dafny.Set.fromElements((_this).f18);
        let _969_v55;
        _969_v55 = _module.D10.create_DC23(!((_966_v52).f18), p1, _966_v52);
        let _arr16 = _966_v52.f19;
        let _index150 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_966_v52.f19).length));
        let _rhs130 = new BigNumber(130);
        let _rhs131 = new BigNumber(986);
        let _rhs132 = (_969_v55).dtor_cf45;
        let _rhs133 = !(((_966_v52).f18) || (((_dafny.ZERO).minus(p0)).isEqualTo(p0)));
        let _rhs134 = _dafny.Set.fromElements((_966_v52).f18);
        let _lhs118 = globalState;
        let _lhs119 = globalState;
        let _lhs120 = _966_v52.f19;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_966_v52.f19).length));
        _lhs118.f14 = _rhs130;
        _lhs119.f14 = _rhs131;
        _966_v52 = _rhs132;
        _lhs120[_lhs121] = _rhs133;
        _968_v54 = _rhs134;
      } else {
        _955_v41 = _955_v41;
        let _970_v56;
        let _init32 = function (_971_i2) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        };
        let _nw157 = Array((new BigNumber(19)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw157.length); _i0_32++) {
          _nw157[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _970_v56 = _nw157;
        (globalState).f16 = _970_v56;
        r2 = (_dafny.ZERO).minus(p0);
        (globalState).f9 = p1;
        let _972_v57;
        _972_v57 = _dafny.Seq.of((_this).f18, p1);
        let _973_v58;
        _973_v58 = _module.D3.create_DC9();
        let _974_v59;
        _974_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_972_v57, _972_v57),_module.__default.fm34(_973_v58, p1, _956_v42, (_this).f18, globalState));
        _955_v41 = (((_974_v59).contains(_dafny.Seq.Concat(_972_v57, _972_v57))) ? ((_974_v59).get(_dafny.Seq.Concat(_972_v57, _972_v57))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(243)), ((_975_v42) => function (_976_i3) {
          return _975_v42;
        })(_956_v42))));
      }
      let _977_v60;
      _977_v60 = _dafny.Seq.of((_958_v44).f25);
      let _978_v62;
      _978_v62 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_977_v60, _module.__default.safeIndex(new BigNumber((function () {
        let _coll45 = new _dafny.Map();
        for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-896), new BigNumber(31))) {
          let _979_v61 = _compr_45;
          if (((new BigNumber(-896)).isLessThanOrEqualTo(_979_v61)) && ((_979_v61).isLessThan(new BigNumber(31)))) {
            _coll45.push([_module.__default.safeDivisionInt(_979_v61, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(810), new BigNumber(797))).cardinality())),_977_v60]);
          }
        }
        return _coll45;
      }()).length), new BigNumber((_977_v60).length)), p2)).length), (_this).f20, p2);
      _978_v62 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-36)), ((_980_p0) => function (_981_i4) {
        return _980_p0;
      })(p0));
      if ((_this).f18) {
        r1 = (_this).f18;
        let _982_v63;
        _982_v63 = _dafny.Seq.of((_this).f18, p1, (_this).f18, (_this).f18, !((_this).f18));
        let _rhs135 = false;
        let _rhs136 = ((_this).f18) || (!(_dafny.areEqual(_982_v63, _982_v63)));
        let _rhs137 = !(true);
        let _lhs122 = globalState;
        let _lhs123 = globalState;
        r0 = _rhs135;
        _lhs122.f9 = _rhs136;
        _lhs123.f3 = _rhs137;
        let _983_v64;
        _983_v64 = _dafny.MultiSet.fromElements((_this).f20);
        let _984_v65;
        _984_v65 = _dafny.Set.fromElements(_958_v44.f26, (_958_v44).f25);
        let _nw158 = Array((new BigNumber(13)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = p1;
        _nw158[(_dafny.ONE).toNumber()] = (_this).f18;
        _nw158[(new BigNumber(2)).toNumber()] = (_983_v64).IsProperSubsetOf(_983_v64);
        _nw158[(new BigNumber(3)).toNumber()] = p1;
        _nw158[(new BigNumber(4)).toNumber()] = p1;
        _nw158[(new BigNumber(5)).toNumber()] = (_this).f18;
        _nw158[(new BigNumber(6)).toNumber()] = (_this).f18;
        _nw158[(new BigNumber(7)).toNumber()] = p1;
        _nw158[(new BigNumber(8)).toNumber()] = p1;
        _nw158[(new BigNumber(9)).toNumber()] = !((_this).f18);
        _nw158[(new BigNumber(10)).toNumber()] = (_this).f18;
        _nw158[(new BigNumber(11)).toNumber()] = !((_dafny.Set.fromElements((_this).f20)).IsDisjointFrom(_984_v65));
        _nw158[(new BigNumber(12)).toNumber()] = !(!((_this).f18)) || ((_this).f18);
        (_this).f19 = _nw158;
        let _985_v66;
        let _986_v67;
        let _987_v68;
        let _out47;
        let _out48;
        let _out49;
        let _outcollector13 = (_958_v44).m4(globalState);
        _out47 = _outcollector13[0];
        _out48 = _outcollector13[1];
        _out49 = _outcollector13[2];
        _985_v66 = _out47;
        _986_v67 = _out48;
        _987_v68 = _out49;
        let _988_v69;
        let _init33 = ((_989_v44) => function (_990_i5) {
          return (_990_i5).multipliedBy((_989_v44).f25);
        })(_958_v44);
        let _nw159 = Array((new BigNumber(17)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw159.length); _i0_33++) {
          _nw159[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _988_v69 = _nw159;
        let _index151 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_988_v69).length));
        (_988_v69)[_index151] = _958_v44.f26;
        let _991_v70;
        _991_v70 = _module.D0.create_DC2(_985_v66, (_958_v44).f25, _988_v69);
        let _index152 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_988_v69).length));
        let _rhs138 = _958_v44.f26;
        let _rhs139 = new BigNumber(832);
        let _rhs140 = (_977_v60)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f28, _958_v44.f26), new BigNumber((_977_v60).length))];
        let _rhs141 = (_991_v70).dtor_cf7;
        let _lhs124 = _988_v69;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_988_v69).length));
        let _lhs126 = _958_v44;
        _lhs124[_lhs125] = _rhs138;
        _lhs126.f26 = _rhs139;
        _987_v68 = _rhs140;
        _988_v69 = _rhs141;
      } else {
        let _992_v71;
        let _nw160 = Array((new BigNumber(2)).toNumber());
        _992_v71 = _nw160;
        let _993_v72;
        _993_v72 = _dafny.Map.Empty.slice().updateUnsafe(_992_v71,(_this).f18);
        _993_v72 = (_993_v72).Merge(_993_v72);
        let _994_v73;
        let _nw161 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _994_v73 = _nw161;
        let _index153 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_994_v73).length));
        (_994_v73)[_index153] = _dafny.Seq.Concat(_955_v41, _955_v41);
        let _index154 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_994_v73).length));
        (_994_v73)[_index154] = _955_v41;
        let _995_v74;
        let _init34 = ((_996_p2) => function (_997_i6) {
          return (_997_i6).minus(_996_p2);
        })(p2);
        let _nw162 = Array((new BigNumber(11)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw162.length); _i0_34++) {
          _nw162[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _995_v74 = _nw162;
        let _index155 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_995_v74).length));
        (_995_v74)[_index155] = (_958_v44).f25;
        let _index156 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_995_v74).length));
        (_995_v74)[_index156] = (_this).f20;
        (globalState).f9 = (_this).f18;
        let _index157 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_995_v74).length));
        (_995_v74)[_index157] = (_958_v44.f26).multipliedBy(_module.__default.safeDivisionInt((_this).f28, _958_v44.f26));
      }
      r0 = p1;
      r1 = (_this).f18;
      let _998_v75;
      _998_v75 = _dafny.Seq.of(true);
      r2 = _module.__default.safeDivisionInt((p0).plus(((((_this).f27).contains((_this).f18)) ? (((_this).f27).get((_this).f18)) : ((_this).fm0(new BigNumber((_998_v75).length), _956_v42, globalState)))), new BigNumber(117));
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Seq.of();
      let _999_i0;
      _999_i0 = _dafny.ZERO;
      L8: {
        while (((_this).f20).isLessThanOrEqualTo(new BigNumber(757))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_999_i0)) {
              break L8;
            }
            _999_i0 = (_999_i0).plus(_dafny.ONE);
            let _index158 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((p1).length));
            (p1)[_index158] = (_this).f28;
            let _1000_v0;
            let _nw163 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1000_v0 = _nw163;
            let _1001_v1;
            _1001_v1 = _dafny.Set.fromElements(_1000_v0, _1000_v0);
            let _index159 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((p1).length));
            (p1)[_index159] = _module.__default.safeModuloInt(new BigNumber(698), new BigNumber((_1001_v1).length));
            let _1002_v2;
            let _nw164 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
            _1002_v2 = _nw164;
            let _1003_v3;
            _1003_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1002_v2,_dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((p1).length))]));
            let _1004_v4;
            _1004_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((p1).length))]);
            _1003_v3 = (_1003_v3).update(_1002_v2, (_dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f20)).Merge(_1004_v4));
            let _1005_v5;
            _1005_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f18);
            r2 = _dafny.Seq.of((p0).isLessThanOrEqualTo(new BigNumber((_1005_v5).length)), (_this).f18);
            let _1006_v6;
            _1006_v6 = _dafny.Seq.of(p0);
            let _1007_v7;
            let _nw165 = new _module.C4();
            _nw165.__ctor(p3, _module.__default.safeModuloInt((p1)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((p1).length))], (_this).f28), _this.f17, !((_this).f18) || ((_module.D9.create_DC20(_1005_v5, new BigNumber((_1006_v6).length), (_this).f18, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), new _dafny.CodePoint('h'.codePointAt(0)))).dtor_cf38));
            _1007_v7 = _nw165;
          }
        }
      }
      let _1008_i1;
      _1008_i1 = _dafny.ZERO;
      L9: {
        while ((p0).isLessThan((_this).f28)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1008_i1)) {
              break L9;
            }
            _1008_i1 = (_1008_i1).plus(_dafny.ONE);
            let _1009_v8;
            _1009_v8 = _dafny.Seq.of((_this).f18, (_this).f18);
            let _1010_v9;
            _1010_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1009_v8,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1009_v8).length),(_this).f28));
            let _1011_v10;
            _1011_v10 = _dafny.Map.Empty.slice().updateUnsafe(((((_this).f27).contains(false)) ? (((_this).f27).get(false)) : (p0)),(_dafny.ZERO).minus((_this).fm28(p3, new BigNumber(426), _1010_v9, globalState)));
            let _1012_v11;
            _1012_v11 = _dafny.Seq.of(new BigNumber((_1011_v10).length), (_this).f28, p0);
            let _1013_v12;
            _1013_v12 = _dafny.Seq.UnicodeFromString("bimux");
            let _source10 = _module.__default.fm35(_module.__default.safeModuloInt(p0, (_this).f20), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), ((_1014_p0) => function (_1015_i2) {
              return _1014_p0;
            })(p0)), _1012_v11), new BigNumber((_1013_v12).length), globalState);
            let _1016___mcc_h0 = _source10;
            let _1017_cf21 = _1016___mcc_h0;
            let _1018_v13;
            _1018_v13 = _module.D2.create_DC6((_this).f28);
            (globalState).f14 = (p3).plus((_1018_v13).dtor_cf13);
            _1013_v12 = _1013_v12;
            let _1019_v14;
            _1019_v14 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),(_this).f18);
            (globalState).f14 = _module.__default.safeModuloInt((new BigNumber((_1019_v14).length)).plus((_this).f20), (new BigNumber(((_this).f27).length)).minus((_this).f28));
            let _arr17 = _this.f19;
            let _index160 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_this.f19).length));
            _arr17[_index160] = false;
            let _1020_v15;
            _1020_v15 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18, (_this).f18, !((_this).f18), (_this).f18);
            let _arr18 = _this.f19;
            let _index161 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_this.f19).length));
            _arr18[_index161] = (_1009_v8)[_module.__default.safeIndex(new BigNumber((_1020_v15).cardinality()), new BigNumber((_1009_v8).length))];
            let _1021_v16;
            let _nw166 = new _module.C2();
            _nw166.__ctor(_this.f17, (_this).f18);
            _1021_v16 = _nw166;
            let _1022_v17;
            _1022_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1013_v12,_1021_v16);
            (globalState).f14 = (((_this).f18) ? (p3) : (_module.__default.safeDivisionInt(new BigNumber((_1022_v17).length), (_this).f28)));
            let _1023_v18;
            _1023_v18 = new _dafny.CodePoint('x'.codePointAt(0));
            let _1024_v19;
            _1024_v19 = _module.D11.create_DC27(_dafny.Seq.of(true, (_1021_v16).f18), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).fm0((_this).f20, _1023_v18, globalState))), new BigNumber(334));
            let _pat_let_tv10 = _1009_v8;
            let _pat_let_tv11 = _1009_v8;
            let _1025_v20;
            _1025_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).fm27(globalState)),function (_pat_let18_0) {
              return function (_1026_dt__update__tmp_h0) {
                return function (_pat_let19_0) {
                  return function (_1027_dt__update_hcf52_h0) {
                    return function (_pat_let20_0) {
                      return function (_1028_dt__update_hcf50_h0) {
                        return _module.D11.create_DC27(_1028_dt__update_hcf50_h0, (_1026_dt__update__tmp_h0).dtor_cf51, _1027_dt__update_hcf52_h0);
                      }(_pat_let20_0);
                    }(_pat_let_tv11);
                  }(_pat_let19_0);
                }(new BigNumber((_pat_let_tv10).length));
              }(_pat_let18_0);
            }(_1024_v19));
            _1025_v20 = (_1025_v20).update(_1009_v8, _1024_v19);
            let _1029_v21;
            _1029_v21 = _dafny.MultiSet.fromElements((_this).f20);
            let _1030_v22;
            _1030_v22 = _dafny.Seq.of(_1029_v21);
            _1030_v22 = _1030_v22;
          }
        }
      }
      if ((_this).f18) {
        if ((_this).f18) {
          let _1031_v23;
          _1031_v23 = _dafny.Seq.of(p1);
          (globalState).f7 = (_1031_v23)[_module.__default.safeIndex(new BigNumber(137), new BigNumber((_1031_v23).length))];
          let _1032_v24;
          _1032_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_this.f19);
          (globalState).f14 = new BigNumber(((_1032_v24).Merge((_1032_v24).Merge(_1032_v24))).length);
          let _1033_v25;
          _1033_v25 = _dafny.Seq.of(new BigNumber(106), p3, (_dafny.ZERO).minus(p0));
          r1 = _dafny.Seq.Concat(_1033_v25, _1033_v25);
          let _1034_v26;
          _1034_v26 = _dafny.Seq.UnicodeFromString("kfekrwy");
          _1034_v26 = _1034_v26;
          let _index162 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((p1).length));
          (p1)[_index162] = p3;
          let _1035_v27;
          _1035_v27 = new _dafny.CodePoint('k'.codePointAt(0));
          let _index163 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((p1).length));
          (p1)[_index163] = (_this).fm0((_this).f20, _1035_v27, globalState);
        } else {
          let _1036_v28;
          _1036_v28 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1037_v29;
          _1037_v29 = _dafny.Seq.of(_1036_v28, _1036_v28, _1036_v28, new _dafny.CodePoint('m'.codePointAt(0)));
          let _1038_v30;
          let _nw167 = Array((new BigNumber(15)).toNumber());
          _nw167[(_dafny.ZERO).toNumber()] = _1036_v28;
          _nw167[(_dafny.ONE).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(2)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(3)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(4)).toNumber()] = _module.__default.fm33(p3, (_this).f18, globalState);
          _nw167[(new BigNumber(5)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(6)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(7)).toNumber()] = (_1037_v29)[_module.__default.safeIndex((_this).f28, new BigNumber((_1037_v29).length))];
          _nw167[(new BigNumber(8)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(9)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(10)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(11)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(12)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(13)).toNumber()] = _1036_v28;
          _nw167[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
          _1038_v30 = _nw167;
          (globalState).f16 = _1038_v30;
          let _1039_v31;
          _1039_v31 = _dafny.Seq.of(p0);
          let _1040_v32;
          _1040_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber((_1039_v31).length)),p3);
          _1036_v28 = _module.__default.fm33((p3).plus(p0), (_this).fm1(_1040_v32, globalState), globalState);
          (globalState).f14 = (_dafny.ZERO).minus(p0);
          let _1041_v33;
          _1041_v33 = _dafny.MultiSet.fromElements(new BigNumber(((_this).f27).length));
          let _1042_v34;
          _1042_v34 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.update(_1039_v31, _module.__default.safeIndex((_this).f28, new BigNumber((_1039_v31).length)), (_this).f20)), _dafny.MultiSet.fromElements(new BigNumber((_1037_v29).length)), _1041_v33, (_1041_v33).Difference(_1041_v33), _dafny.MultiSet.fromElements(new BigNumber(500)));
          _1042_v34 = _1042_v34;
          (globalState).f14 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_this).f20).minus(p3)).minus((_this).f20)));
        }
        let _1043_v35;
        _1043_v35 = _module.D3.create_DC9();
        let _1044_v36;
        _1044_v36 = _dafny.Seq.of(_1043_v35);
        let _1045_v37;
        _1045_v37 = _module.D5.create_DC13((_this).f18, new BigNumber((_dafny.Seq.UnicodeFromString("lhjby")).length), _1044_v36);
        (globalState).f9 = (_1045_v37).dtor_cf23;
        r0 = (_dafny.ZERO).minus((_this).f20);
        let _1046_v38;
        let _nw168 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _1046_v38 = _nw168;
        let _1047_v39;
        _1047_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ch"),p1);
        let _index164 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1046_v38).length));
        (_1046_v38)[_index164] = _1047_v39;
        let _1048_v40;
        _1048_v40 = _dafny.Seq.UnicodeFromString("uqjs");
        let _1049_v41;
        _1049_v41 = _module.D0.create_DC2((_this).f18, new BigNumber((_1048_v40).length), p1);
        let _index165 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1046_v38).length));
        (_1046_v38)[_index165] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1048_v40, _dafny.Seq.UnicodeFromString("dsssjt")),(_1049_v41).dtor_cf7);
        if ((_this).f18) {
          (globalState).f14 = p3;
          let _rhs142 = (_this).f28;
          let _rhs143 = p1;
          let _rhs144 = (_this).f18;
          let _rhs145 = (_this).f18;
          let _lhs127 = globalState;
          let _lhs128 = globalState;
          let _lhs129 = globalState;
          r0 = _rhs142;
          _lhs127.f7 = _rhs143;
          _lhs128.f9 = _rhs144;
          _lhs129.f9 = _rhs145;
          _1048_v40 = _dafny.Seq.UnicodeFromString("ydy");
          let _1050_v42;
          _1050_v42 = _dafny.Seq.of((_this).f18, (_this).f18);
          let _1051_v43;
          _1051_v43 = _dafny.Map.Empty.slice().updateUnsafe(_this.f19,_1050_v42);
          let _rhs146 = _1051_v43;
          let _rhs147 = (_this).f28;
          let _rhs148 = (_this).f18;
          let _lhs130 = globalState;
          let _lhs131 = globalState;
          _1051_v43 = _rhs146;
          _lhs130.f14 = _rhs147;
          _lhs131.f3 = _rhs148;
          let _1052_v44;
          let _nw169 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1052_v44 = _nw169;
          _1052_v44 = _1052_v44;
        } else {
          (globalState).f14 = p0;
          (globalState).f9 = (((_this).f18) ? ((_this).f18) : (true));
          let _1053_v45;
          _1053_v45 = new _dafny.CodePoint('s'.codePointAt(0));
          (_this).f21 = (_dafny.Map.Empty.slice().updateUnsafe(_1053_v45,(_this).f18)).Merge(_this.f21);
          let _index166 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((p1).length));
          (p1)[_index166] = (_this).f20;
          let _index167 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((p1).length));
          (p1)[_index167] = new BigNumber((_dafny.Seq.update(_1048_v40, _module.__default.safeIndex((_this).f20, new BigNumber((_1048_v40).length)), _1053_v45)).length);
          let _1054_v46;
          let _nw170 = Array((new BigNumber(21)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _1053_v45;
          _nw170[(_dafny.ONE).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(2)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(3)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(4)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(5)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(6)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(7)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(8)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw170[(new BigNumber(10)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(11)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(12)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(13)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw170[(new BigNumber(15)).toNumber()] = (_1048_v40)[_module.__default.safeIndex(new BigNumber((_1048_v40).length), new BigNumber((_1048_v40).length))];
          _nw170[(new BigNumber(16)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(17)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(18)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(19)).toNumber()] = _1053_v45;
          _nw170[(new BigNumber(20)).toNumber()] = _1053_v45;
          _1054_v46 = _nw170;
          let _1055_v47;
          _1055_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1054_v46);
          let _1056_v48;
          let _nw171 = Array((new BigNumber(9)).toNumber());
          _nw171[(_dafny.ZERO).toNumber()] = _1054_v46;
          _nw171[(_dafny.ONE).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(2)).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(3)).toNumber()] = (((_1055_v47).contains((_this).f18)) ? ((_1055_v47).get((_this).f18)) : (_1054_v46));
          _nw171[(new BigNumber(4)).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(5)).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(6)).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(7)).toNumber()] = _1054_v46;
          _nw171[(new BigNumber(8)).toNumber()] = _1054_v46;
          _1056_v48 = _nw171;
          let _1057_v49;
          _1057_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1056_v48,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1048_v40).length),(_this).f20));
          let _1058_v50;
          _1058_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
          _1057_v49 = (_1057_v49).update(_1056_v48, _1058_v50);
        }
      } else {
        (globalState).f14 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), function (_1059_i3) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        })).length));
        let _1060_v51;
        _1060_v51 = _dafny.Seq.of(true, !((_this).f18), !((_this).f18), (_this).f18, (_this).f18);
        let _1061_v52;
        _1061_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v51,false);
        let _1062_v53;
        _1062_v53 = _dafny.Set.fromElements(_1060_v51);
        (globalState).f14 = (_dafny.ZERO).minus(new BigNumber(((_1061_v52).update(_1060_v51, !(!(_1062_v53).equals(function () {
          let _coll46 = new _dafny.Set();
          for (const _compr_46 of (_1062_v53).Elements) {
            let _1063_v54 = _compr_46;
            if ((_1062_v53).contains(_1063_v54)) {
              _coll46.add(_1063_v54);
            }
          }
          return _coll46;
        }())))).length));
        let _1064_v55;
        _1064_v55 = _dafny.Set.fromElements((_this).f18, (_this).f18, (((_this).f18) ? ((_this).f18) : ((_this).f18)), !(false) || ((_this).f18));
        let _1065_v56;
        _1065_v56 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18, (_this).f18);
        let _1066_v57;
        _1066_v57 = _dafny.Map.Empty.slice().updateUnsafe((_1065_v56).Intersect(_1065_v56),(_this).f18);
        let _1067_v58;
        _1067_v58 = _dafny.Seq.UnicodeFromString("wupo");
        let _1068_v59;
        _1068_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1067_v58);
        let _rhs149 = !(!((_this).f18));
        let _rhs150 = p2;
        let _rhs151 = new BigNumber(((_1068_v59).Merge(_module.__default.fm36(globalState))).length);
        let _rhs152 = _1066_v57;
        let _lhs132 = globalState;
        _lhs132.f3 = _rhs149;
        _1064_v55 = _rhs150;
        r0 = _rhs151;
        _1066_v57 = _rhs152;
        let _1069_v60;
        let _nw172 = new _module.C5();
        _nw172.__ctor(_1067_v58, (_this).f18, new BigNumber((_1060_v51).length), _this.f21, _this.f19, _this.f17, (_this).f18);
        _1069_v60 = _nw172;
        let _1070_v61;
        _1070_v61 = _module.D10.create_DC23((_this).f18, (_this).f18, _1069_v60);
        let _1071_v62;
        _1071_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1064_v55).Union(_1064_v55),_1070_v61);
        _1071_v62 = (_1071_v62).update(_module.__default.fm37((_1069_v60).f18, new BigNumber((_dafny.Seq.UnicodeFromString("yf")).length), (_1069_v60).f18, (_1069_v60).f18, globalState), _1070_v61);
        (globalState).f14 = p0;
      }
      if ((_this).f18) {
        let _index168 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((p1).length));
        (p1)[_index168] = p0;
        let _1072_v63;
        _1072_v63 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((p2).length)).isLessThan(new BigNumber(-786)),new BigNumber(202));
        let _1073_v64;
        _1073_v64 = _dafny.Seq.of((_this).f18);
        let _1074_v65;
        _1074_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
        let _1075_v66;
        _1075_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1074_v65);
        let _1076_v67;
        _1076_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1073_v64,(((_1075_v66).contains((_this).f18)) ? ((_1075_v66).get((_this).f18)) : (_1074_v65)));
        let _index169 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((p1).length));
        let _rhs153 = ((_this).fm28(p3, (_this).f20, _1076_v67, globalState)).isLessThan(p3);
        let _rhs154 = (_this).f18;
        let _rhs155 = (_dafny.ZERO).minus(p3);
        let _rhs156 = _1072_v63;
        let _lhs133 = globalState;
        let _lhs134 = globalState;
        let _lhs135 = p1;
        let _lhs136 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((p1).length));
        _lhs133.f9 = _rhs153;
        _lhs134.f9 = _rhs154;
        _lhs135[_lhs136] = _rhs155;
        _1072_v63 = _rhs156;
        let _1077_v68;
        _1077_v68 = _dafny.Seq.UnicodeFromString("hifw");
        let _1078_v69;
        _1078_v69 = _dafny.MultiSet.fromElements(_module.D11.create_DC27(_1073_v64, (_dafny.ZERO).minus(p3), new BigNumber((_dafny.Set.fromElements(p3)).length)));
        let _1079_v70;
        _1079_v70 = _dafny.Seq.of((_1078_v69).Intersect(_1078_v69), (_1078_v69).Union(_1078_v69));
        let _1080_v71;
        _1080_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1073_v64);
        let _rhs157 = _1077_v68;
        let _rhs158 = (_1079_v70)[_module.__default.safeIndex(new BigNumber(((((_1080_v71).contains(!(!((_this).f18)))) ? ((_1080_v71).get(!(!((_this).f18)))) : (_1073_v64))).length), new BigNumber((_1079_v70).length))];
        let _rhs159 = (_this).f18;
        let _lhs137 = globalState;
        _1077_v68 = _rhs157;
        _1078_v69 = _rhs158;
        _lhs137.f3 = _rhs159;
        let _arr19 = _this.f19;
        let _index170 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_this.f19).length));
        _arr19[_index170] = (_this).f18;
        let _arr20 = _this.f19;
        let _index171 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_this.f19).length));
        _arr20[_index171] = !(((((!((_this).f18)) ? ((_this).f18) : ((_this).f18))) ? (!((_this).f18)) : ((_this).f18)));
        (globalState).f14 = (p0).minus((_this).f28);
        let _1081_v72;
        let _nw173 = Array((new BigNumber(10)).toNumber()).fill(false);
        _1081_v72 = _nw173;
        let _1082_v73;
        _1082_v73 = _dafny.Seq.of(_1081_v72, _1081_v72, _1081_v72);
        _1082_v73 = _dafny.Seq.Concat(_1082_v73, _1082_v73);
      } else {
        let _1083_v74;
        let _nw174 = new _module.C3();
        _nw174.__ctor(_this.f19, _this.f17, (_this).f18);
        _1083_v74 = _nw174;
        let _1084_v75;
        _1084_v75 = _dafny.Seq.UnicodeFromString("ikjvinf");
        let _1085_v76;
        _1085_v76 = _module.D3.create_DC9();
        let _1086_v77;
        _1086_v77 = new _dafny.CodePoint('a'.codePointAt(0));
        _1084_v75 = _dafny.Seq.Concat(_module.__default.fm34(_1085_v76, (_this).f18, _1086_v77, (_this).f18, globalState), _dafny.Seq.Concat(_1084_v75, _1084_v75));
        let _1087_v78;
        _1087_v78 = _dafny.Seq.of(true);
        let _1088_v79;
        _1088_v79 = _dafny.MultiSet.fromElements(_1087_v78, _1087_v78);
        let _1089_v80;
        _1089_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1085_v76,(new BigNumber((_1088_v79).cardinality())).isLessThanOrEqualTo((_1083_v74).fm0(p0, _1086_v77, globalState)));
        _1089_v80 = (_1089_v80).update(_1085_v76, (_this).f18);
        let _1090_v81;
        _1090_v81 = _dafny.Set.fromElements(p3, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-98), (_this).f28)), new BigNumber((_dafny.Seq.of(((_this).f27).update((_this).f18, (_this).f28))).length));
        _1090_v81 = (_1090_v81).Union(_1090_v81);
        if ((_1087_v78)[_module.__default.safeIndex((_this).f20, new BigNumber((_1087_v78).length))]) {
          let _1091_v82;
          _1091_v82 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f18));
          let _1092_v83;
          _1092_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f18);
          let _1093_v84;
          _1093_v84 = _dafny.Seq.of(_module.__default.fm38(new BigNumber((_1091_v82).cardinality()), (_this).f18, globalState), (_1092_v83).update(new BigNumber(273), true), _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f18), ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f27).length),true)).update((_this).f20, (_this).fm27(globalState))).update((_this).f28, (_this).f18), _1092_v83);
          let _1094_v86;
          _1094_v86 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_dafny.Seq.of(p0, (_this).f28)).Elements) {
              let _1095_v85 = _compr_47;
              if (_dafny.Seq.contains(_dafny.Seq.of(p0, (_this).f28), _1095_v85)) {
                _coll47.push([_module.__default.safeModuloInt(_1095_v85, (_this).f28),new BigNumber(218)]);
              }
            }
            return _coll47;
          }(),new BigNumber((_1084_v75).length));
          _1093_v84 = (((((_this).f18) ? ((_this).f18) : ((_this).fm1(_1094_v86, globalState)))) ? (_dafny.Seq.Concat(_1093_v84, _1093_v84)) : ((((_this).f18) ? (_1093_v84) : (_1093_v84))));
          let _arr21 = _this.f19;
          let _index172 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length));
          _arr21[_index172] = (_this).f18;
          let _arr22 = _this.f19;
          let _index173 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length));
          _arr22[_index173] = !(_module.__default.fm39(p3, globalState)).contains((_this).fm27(globalState));
          let _1096_v87;
          _1096_v87 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18, (_this).f18);
          let _arr23 = _this.f19;
          let _index174 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length));
          let _rhs160 = !(_1096_v87).equals(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _1087_v78)));
          let _rhs161 = (_this).fm27(globalState);
          let _rhs162 = p1;
          let _lhs138 = globalState;
          let _lhs139 = _this.f19;
          let _lhs140 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length));
          let _lhs141 = globalState;
          _lhs138.f9 = _rhs160;
          _lhs139[_lhs140] = _rhs161;
          _lhs141.f7 = _rhs162;
          let _1097_v88;
          _1097_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))],true);
          let _rhs163 = p1;
          let _rhs164 = !((((_1097_v88).contains((!((_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))])) || ((_this).f18))) ? ((_1097_v88).get((!((_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))])) || ((_this).f18))) : ((_this).f18)));
          let _rhs165 = _dafny.Seq.UnicodeFromString("mfq");
          let _rhs166 = _dafny.Seq.contains(_1084_v75, _1086_v77);
          let _rhs167 = (_this).f18;
          let _lhs142 = globalState;
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          _lhs142.f7 = _rhs163;
          _lhs143.f3 = _rhs164;
          _1084_v75 = _rhs165;
          _lhs144.f3 = _rhs166;
          _lhs145.f9 = _rhs167;
          let _1098_v90;
          _1098_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1087_v78,function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))])).Keys.Elements) {
              let _1099_v89 = _compr_48;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))])).contains(_1099_v89)) {
                _coll48.push([(_1099_v89).minus(p0),p3]);
              }
            }
            return _coll48;
          }());
          let _rhs168 = (_this).f28;
          let _rhs169 = (_this).f18;
          let _rhs170 = (_this).fm28((_this).f20, (p0).minus(p0), _1098_v90, globalState);
          let _rhs171 = (_this.f19)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_this.f19).length))];
          let _rhs172 = p3;
          let _lhs146 = globalState;
          let _lhs147 = globalState;
          let _lhs148 = globalState;
          let _lhs149 = globalState;
          let _lhs150 = globalState;
          _lhs146.f14 = _rhs168;
          _lhs147.f3 = _rhs169;
          _lhs148.f14 = _rhs170;
          _lhs149.f3 = _rhs171;
          _lhs150.f14 = _rhs172;
        } else {
          let _1100_v91;
          _1100_v91 = _dafny.Seq.of(_1084_v75);
          let _1101_v92;
          _1101_v92 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.IsProperPrefixOf(_1100_v91, _1100_v91));
          _1101_v92 = (_1101_v92).update((_this).f18, (((_this).f18) ? ((_this).f18) : ((_1087_v78)[_module.__default.safeIndex(p3, new BigNumber((_1087_v78).length))])));
          _1090_v81 = _1090_v81;
          let _1102_v93;
          _1102_v93 = _dafny.Seq.of(new BigNumber(((_dafny.Set.fromElements(false, (_this).f18, (_this).f18)).Union(p2)).length), (_this).f20);
          r0 = (_1102_v93)[_module.__default.safeIndex(p0, new BigNumber((_1102_v93).length))];
          (_this).f19 = _this.f19;
          (globalState).f9 = ((_dafny.ZERO).minus((p3).multipliedBy((_this).f28))).isEqualTo(p3);
        }
      }
      let _1103_v94;
      let _nw175 = new _module.C2();
      _nw175.__ctor((_this.f17).Difference(_this.f17), (_this).f18);
      _1103_v94 = _nw175;
      let _1104_v95;
      _1104_v95 = _dafny.Seq.of((_this).f18);
      let _1105_v96;
      _1105_v96 = _dafny.MultiSet.fromElements((_this).f18);
      let _arr24 = _this.f19;
      let _index175 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_this.f19).length));
      _arr24[_index175] = (_1104_v95)[_module.__default.safeIndex(new BigNumber((_1105_v96).cardinality()), new BigNumber((_1104_v95).length))];
      let _arr25 = _this.f19;
      let _index176 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_this.f19).length));
      _arr25[_index176] = (_this).f18;
      r0 = new BigNumber((_1105_v96).cardinality());
      let _1106_v97;
      _1106_v97 = _dafny.Seq.of(p0);
      let _1107_v98;
      _1107_v98 = _dafny.Set.fromElements((_this).f20);
      let _1108_v99;
      _1108_v99 = _dafny.Seq.UnicodeFromString("dsaa");
      let _1109_v100;
      _1109_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1104_v95).length),p3);
      let _1110_v101;
      _1110_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1104_v95,_1109_v100);
      let _1111_v102;
      _1111_v102 = _dafny.Seq.update(_1106_v97, _module.__default.safeIndex(new BigNumber((_1107_v98).length), new BigNumber((_1106_v97).length)), (_this).fm28(new BigNumber((_1108_v99).length), p3, _1110_v101, globalState));
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_1106_v97, _dafny.Seq.Concat(_1106_v97, (_1111_v102))), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1108_v99, _dafny.Seq.UnicodeFromString("o"))).length), new BigNumber((_dafny.Seq.Concat(_1106_v97, _dafny.Seq.Concat(_1106_v97, (_1111_v102)))).length)), (_this).f20);
      r2 = _1104_v95;
      return [r0, r1, r2];
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _1112_v0;
      _1112_v0 = _dafny.Seq.UnicodeFromString("beccgqvi");
      (globalState).f14 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_1113_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _1112_v0)).length));
      let _1114_v1;
      _1114_v1 = _dafny.Seq.of((_this).f18, (_this).f18, (_this).f18, (_this).f18, (_this).f18);
      let _1115_v2;
      _1115_v2 = _dafny.MultiSet.fromElements(_1114_v1);
      (globalState).f3 = (p0).isLessThanOrEqualTo(new BigNumber((_1115_v2).cardinality()));
      let _1116_v3;
      _1116_v3 = new _dafny.CodePoint('j'.codePointAt(0));
      _1116_v3 = _1116_v3;
      let _1117_v4;
      _1117_v4 = _dafny.Seq.of(new BigNumber((_1112_v0).length), p0, p0);
      let _1118_v5;
      _1118_v5 = _dafny.Set.fromElements((_this).f18);
      (globalState).f14 = (new BigNumber((_1117_v4).length)).multipliedBy(new BigNumber(((_dafny.Set.fromElements((_this).f18)).Union(_1118_v5)).length));
      (globalState).f14 = ((_this).f28).minus(new BigNumber((_1112_v0).length));
      let _1119_i1;
      _1119_i1 = _dafny.ZERO;
      L10: {
        while ((_this).f18) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1119_i1)) {
              break L10;
            }
            _1119_i1 = (_1119_i1).plus(_dafny.ONE);
            (globalState).f9 = (_this).f18;
            let _1120_v6;
            _1120_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_1112_v0, _1112_v0),(_this).f18);
            _1120_v6 = (_1120_v6).update((_this).f18, !((_this).f18));
            _1112_v0 = (((_this).f18) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), ((_1121_v3) => function (_1122_i2) {
              return _1121_v3;
            })(_1116_v3))) : (_1112_v0));
            let _1123_v7;
            let _nw176 = new _module.C1();
            _nw176.__ctor();
            _1123_v7 = _nw176;
          }
        }
      }
      let _1124_v8;
      let _nw177 = new _module.C4();
      _nw177.__ctor((_this).f28, new BigNumber((_1118_v5).length), _this.f17, (_this).f18);
      _1124_v8 = _nw177;
      let _1125_v9;
      _1125_v9 = _dafny.Seq.of(_1124_v8);
      let _1126_v10;
      _1126_v10 = _module.D12.create_DC29(_1125_v9);
      let _1127_v11;
      _1127_v11 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_1125_v9, _module.__default.safeIndex(((((_this).f27).contains((_1124_v8).f18)) ? (((_this).f27).get((_1124_v8).f18)) : ((_this).f28)), new BigNumber((_1125_v9).length)), _1124_v8), _1125_v9), _1125_v9, (_1126_v10).dtor_cf54);
      r0 = (_1127_v11)[_module.__default.safeIndex((_this).f28, new BigNumber((_1127_v11).length))];
      r1 = (((_this).f18) ? ((((_this).f18) ? ((_this).f28) : (p0))) : ((_this).f20));
      let _1128_v12;
      _1128_v12 = _module.D3.create_DC10(_1112_v0, (_1124_v8).f18, _dafny.Seq.UnicodeFromString("hym"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), ((_1129_v3) => function (_1130_i3) {
  return _1129_v3;
})(_1116_v3)), (_dafny.ZERO).minus(p0));
      let _1131_v13;
      _1131_v13 = _dafny.MultiSet.fromElements((_1124_v8).f18, (_this).f18, (_this).f18, (_1128_v12).dtor_cf17);
      r2 = (_1131_v13).IsSubsetOf(_1131_v13);
      r3 = _1117_v4;
      return [r0, r1, r2, r3];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f21 = _dafny.Map.Empty;
      this._f20 = _dafny.ZERO;
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f20, f21, f19, f17, f18) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_this).f20;
    };
    fm25(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.of((new BigNumber((_dafny.Seq.UnicodeFromString("cd")).length)).isLessThan((_dafny.ZERO).minus((_this).f20)), (_this).f18, (_this).f18, (_this).f18);
    };
    fm26(globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements((_this).f18, (_this).f18));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1132_v0;
      _1132_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,p1);
      let _1133_v1;
      _1133_v1 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(797),p1));
      let _1134_v2;
      _1134_v2 = _module.D1.create_DC4(((((_1132_v0).contains(p0)) ? ((_1132_v0).get(p0)) : (true))) && ((_this).f18), (_dafny.Set.fromElements(_1132_v0)).Intersect(_1133_v1), p0);
      let _source11 = _1134_v2;
      if (_source11.is_DC4) {
        let _1135___mcc_h0 = (_source11).cf9;
        let _1136___mcc_h1 = (_source11).cf10;
        let _1137___mcc_h2 = (_source11).cf11;
        let _1138_cf11 = _1137___mcc_h2;
        let _1139_cf10 = _1136___mcc_h1;
        let _1140_cf9 = _1135___mcc_h0;
        let _1141_v3;
        _1141_v3 = _dafny.Seq.UnicodeFromString("edctjguo");
        let _1142_v4;
        let _nw178 = new _module.C5();
        _nw178.__ctor(_1141_v3, p1, (_this).f20, _this.f21, _this.f19, _this.f17, _1140_cf9);
        _1142_v4 = _nw178;
        _1142_v4 = _1142_v4;
        _1138_cf11 = (p0).minus(_1138_cf11);
        let _1143_v5;
        _1143_v5 = _dafny.MultiSet.fromElements(_1134_v2);
        (globalState).f14 = (((_1143_v5).contains(_module.D1.create_DC4(_1140_cf9, _1133_v1, new BigNumber(885)))) ? ((_1143_v5).get(_module.D1.create_DC4(_1140_cf9, _1133_v1, new BigNumber(885)))) : ((_this).f20));
        let _1144_v6;
        _1144_v6 = _dafny.Map.Empty.slice().updateUnsafe((p2).isLessThan(_1138_cf11),_this.f19);
        let _1145_v7;
        _1145_v7 = _dafny.Seq.of((((_1132_v0).contains((_1142_v4).f20)) ? ((_1132_v0).get((_1142_v4).f20)) : (false)), (_1142_v4).f18, _1140_cf9, !(!(!((_1142_v4).f18))), (_1142_v4).f18);
        _1144_v6 = (_1144_v6).update((_1145_v7)[_module.__default.safeIndex(p2, new BigNumber((_1145_v7).length))], _this.f19);
      } else {
        let _1146___mcc_h3 = (_source11).cf8;
        let _1147_cf8 = _1146___mcc_h3;
        let _1148_v8;
        _1148_v8 = _dafny.Seq.UnicodeFromString("k");
        let _1149_v9;
        _1149_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.Concat(_1148_v8, _1148_v8));
        let _1150_v10;
        _1150_v10 = _dafny.Seq.of(_1149_v9);
        let _rhs173 = !(!(p1));
        let _rhs174 = (_1150_v10)[_module.__default.safeIndex(p0, new BigNumber((_1150_v10).length))];
        let _rhs175 = (_module.D2.create_DC6(new BigNumber(-752))).dtor_cf13;
        let _rhs176 = _this.f19;
        let _rhs177 = _1147_cf8;
        let _lhs151 = globalState;
        let _lhs152 = _this;
        _lhs151.f3 = _rhs173;
        _1149_v9 = _rhs174;
        r2 = _rhs175;
        _lhs152.f19 = _rhs176;
        _1147_cf8 = _rhs177;
        let _1151_v11;
        let _init35 = ((_1152_p2) => function (_1153_i0) {
          return _module.__default.safeModuloInt(_1153_i0, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_1152_p2, new BigNumber((_dafny.Set.fromElements(false)).length), (_this).f20)).length), _1152_p2)).length));
        })(p2);
        let _nw179 = Array((new BigNumber(2)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw179.length); _i0_35++) {
          _nw179[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1151_v11 = _nw179;
        let _index177 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_1151_v11).length));
        (_1151_v11)[_index177] = _module.__default.safeModuloInt(p0, new BigNumber(336));
        let _index178 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_1151_v11).length));
        (_1151_v11)[_index178] = (_this).f20;
        let _1154_v12;
        _1154_v12 = new _dafny.CodePoint('v'.codePointAt(0));
        let _arr26 = _this.f19;
        let _index179 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_this.f19).length));
        _arr26[_index179] = p1;
        let _1155_v13;
        _1155_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,p2);
        let _1156_v14;
        _1156_v14 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_1148_v8).length), p0));
        let _arr27 = _this.f19;
        let _index180 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_this.f19).length));
        let _rhs178 = ((p1) ? (_1154_v12) : (_1154_v12));
        let _rhs179 = ((((_1155_v13).contains(new BigNumber((_1156_v14).length))) ? ((_1155_v13).get(new BigNumber((_1156_v14).length))) : (new BigNumber(249)))).isLessThan(_module.__default.safeModuloInt((_1151_v11)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_1151_v11).length))], (_1151_v11)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_1151_v11).length))]));
        let _lhs153 = _this.f19;
        let _lhs154 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_this.f19).length));
        _1154_v12 = _rhs178;
        _lhs153[_lhs154] = _rhs179;
        let _1157_v15;
        _1157_v15 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("aw"), p0, new BigNumber(155), _1154_v12);
        let _1158_v16;
        _1158_v16 = _dafny.Seq.of(_1148_v8, _dafny.Seq.Concat((_1157_v15).dtor_cf1, _1148_v8), _dafny.Seq.Create(_module.__default.abs(new BigNumber(557)), function (_1159_i1) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _1148_v8);
        _1158_v16 = _dafny.Seq.Concat(_1158_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(386)), ((_1160_v8) => function (_1161_i2) {
          return _1160_v8;
        })(_1148_v8)));
      }
      let _1162_v17;
      _1162_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f19);
      let _1163_v18;
      _1163_v18 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1162_v17).length));
      if ((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(_1163_v18,p0), globalState)) {
        let _1164_v19;
        let _nw180 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1164_v19 = _nw180;
        let _1165_v20;
        _1165_v20 = _dafny.Seq.of(p0, p2, p2);
        let _1166_v21;
        _1166_v21 = new _dafny.CodePoint('s'.codePointAt(0));
        let _rhs180 = _1164_v19;
        let _rhs181 = (_this).fm0(new BigNumber((_dafny.Seq.Concat(_1165_v20, _1165_v20)).length), _1166_v21, globalState);
        let _rhs182 = false;
        let _rhs183 = !(true);
        let _rhs184 = !(!((_this).f18));
        let _lhs155 = globalState;
        let _lhs156 = globalState;
        let _lhs157 = globalState;
        let _lhs158 = globalState;
        let _lhs159 = globalState;
        _lhs155.f7 = _rhs180;
        _lhs156.f14 = _rhs181;
        _lhs157.f3 = _rhs182;
        _lhs158.f9 = _rhs183;
        _lhs159.f9 = _rhs184;
        let _1167_v22;
        _1167_v22 = _dafny.Seq.UnicodeFromString("fm");
        let _1168_v23;
        let _nw181 = new _module.C5();
        _nw181.__ctor(_1167_v22, !(p1), new BigNumber((_1167_v22).length), _this.f21, _this.f19, (_this.f17).Difference(_this.f17), (_this).f18);
        _1168_v23 = _nw181;
        let _1169_v24;
        _1169_v24 = _dafny.Map.Empty.slice().updateUnsafe(false,p2);
        _1169_v24 = _1169_v24;
        if (false) {
          _1164_v19 = _1164_v19;
          _1166_v21 = _1166_v21;
          r1 = ((_this).f20).isLessThanOrEqualTo(p2);
          let _1170_v25;
          let _nw182 = new _module.C2();
          _nw182.__ctor(_this.f17, (_1168_v23).f30);
          _1170_v25 = _nw182;
          let _index181 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1164_v19).length));
          (_1164_v19)[_index181] = (_this).f20;
          let _index182 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1164_v19).length));
          (_1164_v19)[_index182] = (_this).f20;
        } else {
          let _arr28 = _this.f19;
          let _index183 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_this.f19).length));
          _arr28[_index183] = !(false) || ((_this).f18);
          let _arr29 = _this.f19;
          let _index184 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_this.f19).length));
          _arr29[_index184] = (_1168_v23).f30;
          let _1171_v26;
          let _nw183 = Array((new BigNumber(22)).toNumber()).fill([]);
          _1171_v26 = _nw183;
          let _1172_v27;
          _1172_v27 = _dafny.Seq.of(_1171_v26);
          let _1173_v28;
          _1173_v28 = _dafny.Map.Empty.slice().updateUnsafe((((_1163_v18).contains(p0)) ? ((_1163_v18).get(p0)) : ((_this).f20)),_1171_v26);
          let _1174_v29;
          _1174_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(((_1173_v28).contains(p2)) ? ((_1173_v28).get(p2)) : (_1171_v26)));
          let _1175_v30;
          let _nw184 = Array((new BigNumber(18)).toNumber());
          _nw184[(_dafny.ZERO).toNumber()] = _1171_v26;
          _nw184[(_dafny.ONE).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(2)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(3)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(4)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(5)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(6)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(7)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(8)).toNumber()] = (_1172_v27)[_module.__default.safeIndex(p0, new BigNumber((_1172_v27).length))];
          _nw184[(new BigNumber(9)).toNumber()] = (((_1173_v28).contains(p0)) ? ((_1173_v28).get(p0)) : (_1171_v26));
          _nw184[(new BigNumber(10)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(11)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(12)).toNumber()] = (((_1174_v29).contains(p1)) ? ((_1174_v29).get(p1)) : (_1171_v26));
          _nw184[(new BigNumber(13)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(14)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(15)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(16)).toNumber()] = _1171_v26;
          _nw184[(new BigNumber(17)).toNumber()] = _1171_v26;
          _1175_v30 = _nw184;
          let _index185 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1175_v30).length));
          (_1175_v30)[_index185] = _1171_v26;
          let _index186 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1175_v30).length));
          (_1175_v30)[_index186] = _1171_v26;
          _1164_v19 = _1164_v19;
          let _1176_v31;
          _1176_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v18,p0);
          let _1177_v35;
          let _nw185 = Array((new BigNumber(14)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = (_1163_v18).Merge(_1163_v18);
          _nw185[(_dafny.ONE).toNumber()] = (_1163_v18).Merge(_1163_v18);
          _nw185[(new BigNumber(2)).toNumber()] = _1163_v18;
          _nw185[(new BigNumber(3)).toNumber()] = (((_1168_v23).fm1(_1176_v31, globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe(p2,p2)) : (function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of _dafny.IntegerRange(new BigNumber(982), new BigNumber(808))) {
              let _1178_v32 = _compr_49;
              if (((new BigNumber(982)).isLessThanOrEqualTo(_1178_v32)) && ((_1178_v32).isLessThan(new BigNumber(808)))) {
                _coll49.push([(_1178_v32).plus(p0),p2]);
              }
            }
            return _coll49;
          }()));
          _nw185[(new BigNumber(4)).toNumber()] = _1163_v18;
          _nw185[(new BigNumber(5)).toNumber()] = (_1163_v18).Merge(function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of (_dafny.MultiSet.fromElements((_this).f20, p0)).Elements) {
              let _1179_v33 = _compr_50;
              if ((_dafny.MultiSet.fromElements((_this).f20, p0)).contains(_1179_v33)) {
                _coll50.push([_module.__default.safeModuloInt(_1179_v33, p2),(_this).f20]);
              }
            }
            return _coll50;
          }());
          _nw185[(new BigNumber(6)).toNumber()] = _1163_v18;
          _nw185[(new BigNumber(7)).toNumber()] = _1163_v18;
          _nw185[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_dafny.ZERO).minus(p0));
          _nw185[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _nw185[(new BigNumber(10)).toNumber()] = (_1163_v18).Merge(_1163_v18);
          _nw185[(new BigNumber(11)).toNumber()] = _1163_v18;
          _nw185[(new BigNumber(12)).toNumber()] = (_1163_v18).Merge(function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(163), new BigNumber(967))) {
              let _1180_v34 = _compr_51;
              if (((new BigNumber(163)).isLessThanOrEqualTo(_1180_v34)) && ((_1180_v34).isLessThan(new BigNumber(967)))) {
                _coll51.push([(_1180_v34).plus((_this).f20),p2]);
              }
            }
            return _coll51;
          }());
          _nw185[(new BigNumber(13)).toNumber()] = _1163_v18;
          _1177_v35 = _nw185;
          let _index187 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1177_v35).length));
          (_1177_v35)[_index187] = _1163_v18;
          let _index188 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1177_v35).length));
          (_1177_v35)[_index188] = _1163_v18;
          _1132_v0 = (_1132_v0).update(new BigNumber(289), p1);
        }
        let _1181_v36;
        _1181_v36 = _module.D0.create_DC2((_this).f18, p0, _1164_v19);
        let _1182_v37;
        let _nw186 = new _module.C4();
        _nw186.__ctor(p0, p0, _this.f17, (p0).isLessThan((_dafny.ZERO).minus((_1181_v36).dtor_cf6)));
        _1182_v37 = _nw186;
      } else {
        let _1183_v38;
        _1183_v38 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1184_v39;
        _1184_v39 = _dafny.Seq.UnicodeFromString("n");
        let _1185_v40;
        _1185_v40 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qlsh"), _1184_v39, _1184_v39, _1184_v39, _1184_v39);
        let _1186_v41;
        _1186_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1183_v38,new BigNumber(((_1185_v40)[_module.__default.safeIndex((_this).fm0((_this).f20, _module.__default.fm33(new BigNumber((_1184_v39).length), (_this).f18, globalState), globalState), new BigNumber((_1185_v40).length))]).length));
        let _1187_v42;
        _1187_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v18,new BigNumber((_1186_v41).length));
        if (!((_this).fm1((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(p1, globalState),(_this).f20)).Merge(_1187_v42), globalState))) {
          let _arr30 = _this.f19;
          let _index189 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length));
          _arr30[_index189] = (((_this).f18) ? (true) : ((_this).f18));
          let _1188_v43;
          _1188_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_this.f19);
          let _1189_v44;
          _1189_v44 = _module.D5.create_DC12(_1188_v43);
          let _1190_v45;
          _1190_v45 = _dafny.MultiSet.fromElements(_1189_v44);
          let _1191_v46;
          _1191_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1190_v45);
          let _1192_v47;
          _1192_v47 = _dafny.MultiSet.fromElements(p2);
          let _1193_v48;
          _1193_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1191_v46,new BigNumber((_1192_v47).cardinality()));
          let _arr31 = _this.f19;
          let _index190 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length));
          _arr31[_index190] = (p0).isLessThanOrEqualTo((((_1193_v48).contains(_1191_v46)) ? ((_1193_v48).get(_1191_v46)) : (new BigNumber(748))));
          (globalState).f14 = (_this).f20;
          let _1194_v49;
          let _nw187 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1194_v49 = _nw187;
          let _index191 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1194_v49).length));
          (_1194_v49)[_index191] = _1183_v38;
          let _1195_v51;
          _1195_v51 = _dafny.Set.fromElements(_1183_v38, _1183_v38, _1183_v38, _1183_v38, _1183_v38);
          let _1196_v52;
          _1196_v52 = _dafny.MultiSet.fromElements(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of (_this.f21).Keys.Elements) {
              let _1197_v50 = _compr_52;
              if ((_this.f21).contains(_1197_v50)) {
                _coll52.add(_1197_v50);
              }
            }
            return _coll52;
          }(), _1195_v51);
          let _index192 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1194_v49).length));
          (_1194_v49)[_index192] = _module.__default.fm33((((_1196_v52).contains(_1195_v51)) ? ((_1196_v52).get(_1195_v51)) : (p2)), (_this.f19)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length))], globalState);
          let _arr32 = _this.f19;
          let _index193 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length));
          _arr32[_index193] = (_this.f19)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length))];
          let _1198_v53;
          _1198_v53 = _module.D11.create_DC26(p1, p2);
          let _1199_v54;
          let _nw188 = Array((new BigNumber(19)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = p1;
          _nw188[(_dafny.ONE).toNumber()] = true;
          _nw188[(new BigNumber(2)).toNumber()] = true;
          _nw188[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(4)).toNumber()] = true;
          _nw188[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(7)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length))];
          _nw188[(new BigNumber(8)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(9)).toNumber()] = true;
          _nw188[(new BigNumber(10)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length))];
          _nw188[(new BigNumber(11)).toNumber()] = p1;
          _nw188[(new BigNumber(12)).toNumber()] = !(p1);
          _nw188[(new BigNumber(13)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(14)).toNumber()] = (_1198_v53).dtor_cf48;
          _nw188[(new BigNumber(15)).toNumber()] = (_this).f18;
          _nw188[(new BigNumber(16)).toNumber()] = p1;
          _nw188[(new BigNumber(17)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_this.f19).length))];
          _nw188[(new BigNumber(18)).toNumber()] = p1;
          _1199_v54 = _nw188;
          let _1200_v55;
          let _nw189 = new _module.C3();
          _nw189.__ctor(_1199_v54, _this.f17, (_1198_v53).dtor_cf48);
          _1200_v55 = _nw189;
          _1200_v55 = (((p0).isLessThan(p0)) ? (_1200_v55) : (_1200_v55));
        } else {
          let _1201_v56;
          _1201_v56 = _dafny.MultiSet.fromElements(p2);
          let _1202_v57;
          _1202_v57 = _dafny.Seq.of((_dafny.ZERO).minus(p2), p0, (_dafny.ZERO).minus((p0).plus((_this).fm0(p0, _1183_v38, globalState))), new BigNumber(((_1201_v56).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_this).f18)).length), p0, p0, p0, p0))).cardinality()));
          let _1203_v58;
          _1203_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          let _1204_v59;
          let _nw190 = Array((new BigNumber(5)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = new BigNumber(708);
          _nw190[(_dafny.ONE).toNumber()] = (new BigNumber((_1203_v58).length)).plus((_this).f20);
          _nw190[(new BigNumber(2)).toNumber()] = (p0).plus(p0);
          _nw190[(new BigNumber(3)).toNumber()] = ((_this).f20).multipliedBy(p0);
          _nw190[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p2);
          _1204_v59 = _nw190;
          let _1205_v60;
          _1205_v60 = _dafny.Seq.of(_1203_v58, _1203_v58);
          let _index194 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length));
          (_1204_v59)[_index194] = (_dafny.ZERO).minus(new BigNumber((((_1205_v60)[_module.__default.safeIndex(p2, new BigNumber((_1205_v60).length))]).update((_this).f18, p1)).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length));
          let _rhs185 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_1206_v57) => function (_1207_i3) {
            return new BigNumber((_1206_v57).length);
          })(_1202_v57));
          let _rhs186 = p0;
          let _rhs187 = (p2).isLessThan((_this).f20);
          let _lhs160 = _1204_v59;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length));
          _1202_v57 = _rhs185;
          _lhs160[_lhs161] = _rhs186;
          r1 = _rhs187;
          let _1208_v61;
          _1208_v61 = _module.D11.create_DC25(_dafny.Set.fromElements(_this.f19));
          let _1209_v62;
          _1209_v62 = _dafny.Set.fromElements(_1208_v61, _1208_v61, _module.D11.create_DC25(_this.f17), _1208_v61);
          let _index196 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length));
          let _rhs188 = (_1209_v62).Difference(_1209_v62);
          let _rhs189 = _module.__default.safeModuloInt((_this).f20, (_1204_v59)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length))]);
          let _lhs162 = _1204_v59;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length));
          _1209_v62 = _rhs188;
          _lhs162[_lhs163] = _rhs189;
          let _arr33 = _this.f19;
          let _index197 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f19).length));
          _arr33[_index197] = (_this).f18;
          let _1210_v63;
          let _nw191 = new _module.C3();
          _nw191.__ctor(_this.f19, _this.f17, (_this).f18);
          _1210_v63 = _nw191;
          let _1211_v64;
          _1211_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1210_v63,(_this).f20);
          let _arr34 = _this.f19;
          let _index198 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f19).length));
          _arr34[_index198] = !((_1211_v64).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1210_v63,new BigNumber((_1201_v56).cardinality())))).contains(_1210_v63);
          let _1212_v65;
          _1212_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1204_v59,_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f18));
          let _rhs190 = (_this.f19)[_module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f19).length))];
          let _rhs191 = ((_1204_v59)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_1204_v59).length))]).multipliedBy((_this).f20);
          let _rhs192 = _1204_v59;
          let _rhs193 = _dafny.Seq.update(_dafny.Seq.update(_1184_v39, _module.__default.safeIndex((_1210_v63).fm0(new BigNumber((_module.__default.fm40((_this).f18, p1, globalState)).length), _1183_v38, globalState), new BigNumber((_1184_v39).length)), _1183_v38), _module.__default.safeIndex(new BigNumber((((((_1212_v65).contains(_1204_v59)) ? ((_1212_v65).get(_1204_v59)) : (_1203_v58))).Merge(_1203_v58)).length), new BigNumber((_dafny.Seq.update(_1184_v39, _module.__default.safeIndex((_1210_v63).fm0(new BigNumber((_module.__default.fm40((_this).f18, p1, globalState)).length), _1183_v38, globalState), new BigNumber((_1184_v39).length)), _1183_v38)).length)), _1183_v38);
          let _lhs164 = globalState;
          let _lhs165 = globalState;
          _lhs164.f3 = _rhs190;
          r2 = _rhs191;
          _lhs165.f7 = _rhs192;
          _1184_v39 = _rhs193;
          r2 = _module.__default.safeDivisionInt(p2, new BigNumber((_1202_v57).length));
        }
        let _1213_v66;
        _1213_v66 = _dafny.MultiSet.fromElements((_this).f20);
        (globalState).f15 = _1213_v66;
        let _1214_v67;
        let _nw192 = Array((new BigNumber(2)).toNumber()).fill(_module.D3.Default());
        _1214_v67 = _nw192;
        let _1215_v68;
        _1215_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1214_v67,_this.f17);
        let _1216_v69;
        _1216_v69 = _module.D11.create_DC25((((_1215_v68).contains(_1214_v67)) ? ((_1215_v68).get(_1214_v67)) : (_this.f17)));
        let _1217_v70;
        let _init36 = ((_1218_v0, _1219_p0, _1220_p1, _1221_v18, _1222_v38) => function (_1223_i4) {
          return _module.D9.create_DC20(_1218_v0, _1219_p0, _1220_p1, _1221_v18, _1222_v38);
        })(_1132_v0, p0, p1, _1163_v18, _1183_v38);
        let _nw193 = Array((new BigNumber(8)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw193.length); _i0_36++) {
          _nw193[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1217_v70 = _nw193;
        let _1224_v71;
        _1224_v71 = _module.D9.create_DC20(_1132_v0, p0, p1, (_1163_v18).update(p0, p0), _1183_v38);
        let _index199 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1217_v70).length));
        (_1217_v70)[_index199] = _1224_v71;
        let _1225_v72;
        _1225_v72 = _dafny.Set.fromElements((_this).f18, true, p1);
        let _index200 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1217_v70).length));
        let _rhs194 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1184_v39, _1184_v39), _1184_v39);
        let _rhs195 = _1216_v69;
        let _rhs196 = (((_1163_v18).contains(new BigNumber(551))) ? ((_1163_v18).get(new BigNumber(551))) : (p0));
        let _rhs197 = (_this).fm1(_1187_v42, globalState);
        let _rhs198 = _module.D9.create_DC20(_1132_v0, p2, (_1225_v72).IsSubsetOf(_dafny.Set.fromElements(p1)), _1163_v18, _1183_v38);
        let _lhs166 = _1217_v70;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1217_v70).length));
        _1184_v39 = _rhs194;
        _1216_v69 = _rhs195;
        r2 = _rhs196;
        r1 = _rhs197;
        _lhs166[_lhs167] = _rhs198;
        _1132_v0 = (_1132_v0).update(p2, (_1213_v66).IsSubsetOf(_1213_v66));
        (globalState).f9 = (p2).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("dca")).length));
      }
      r1 = p1;
      let _1226_v73;
      _1226_v73 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1227_v74;
      _1227_v74 = _dafny.Seq.UnicodeFromString("xiu");
      let _1228_v75;
      _1228_v75 = _module.D3.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), function (_1229_i5) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("xwn"), _1226_v73), _1227_v74, _1227_v74, new BigNumber((_1227_v74).length));
      let _source12 = _1228_v75;
      if (_source12.is_DC9) {
        if (true) {
          let _1230_v76;
          _1230_v76 = _dafny.Seq.of(p0, p0);
          let _1231_v77;
          let _nw194 = new _module.C5();
          _nw194.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_1232_i6) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }), (_this).f18, (_dafny.ZERO).minus((_1230_v76)[_module.__default.safeIndex((_this).f20, new BigNumber((_1230_v76).length))]), _dafny.Map.Empty.slice().updateUnsafe(_1226_v73,(_this).f18), _this.f19, _this.f17, p1);
          _1231_v77 = _nw194;
          let _arr35 = _this.f19;
          let _index201 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f19).length));
          _arr35[_index201] = (p2).isLessThanOrEqualTo(p0);
          let _1233_v78;
          _1233_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v18,p0);
          let _arr36 = _this.f19;
          let _index202 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f19).length));
          _arr36[_index202] = (_this).fm1(_1233_v78, globalState);
          let _1234_v79;
          let _init37 = ((_1235_v77, _1236_p0, _1237_v0, _1238_v76) => function (_1239_i7) {
            return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1235_v77).f30,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_1235_v77).f30, (((_1237_v0).contains(_1236_p0)) ? ((_1237_v0).get(_1236_p0)) : (true))))).cardinality()))).length)), _1238_v76);
          })(_1231_v77, p0, _1132_v0, _1230_v76);
          let _nw195 = Array((new BigNumber(13)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw195.length); _i0_37++) {
            _nw195[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1234_v79 = _nw195;
          let _index203 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1234_v79).length));
          (_1234_v79)[_index203] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), function (_1240_i8) {
            return new BigNumber(940);
          });
          let _1241_v80;
          _1241_v80 = _dafny.Set.fromElements((_this).f20);
          let _1242_v81;
          let _nw196 = Array((new BigNumber(9)).toNumber());
          _1242_v81 = _nw196;
          let _1243_v82;
          _1243_v82 = _dafny.Seq.of(!(false));
          let _index204 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1234_v79).length));
          let _rhs199 = _1230_v76;
          let _rhs200 = _1241_v80;
          let _rhs201 = _1242_v81;
          let _rhs202 = (p0).minus(_module.__default.safeModuloInt(p0, p0));
          let _rhs203 = (_1243_v82)[_module.__default.safeIndex(p2, new BigNumber((_1243_v82).length))];
          let _lhs168 = _1234_v79;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_1234_v79).length));
          let _lhs170 = globalState;
          let _lhs171 = globalState;
          _lhs168[_lhs169] = _rhs199;
          _1241_v80 = _rhs200;
          _1242_v81 = _rhs201;
          _lhs170.f14 = _rhs202;
          _lhs171.f9 = _rhs203;
          let _1244_v83;
          _1244_v83 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("mjn"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("mjn")).length)), _1226_v73), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("mjn"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("mjn")).length)), _1226_v73)).length)), _1226_v73), _1227_v74));
          let _rhs204 = ((_1244_v83).Merge(_1244_v83)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_1231_v77).f29));
          let _rhs205 = (p0).multipliedBy(p0);
          let _rhs206 = (_this).fm25((_this).f18, p1, p2, (p1) === ((_this).f18), globalState);
          let _rhs207 = (((_this.f19)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f19).length))]) ? ((_this).f18) : ((_1231_v77).fm1(_1233_v78, globalState)));
          let _lhs172 = globalState;
          let _lhs173 = globalState;
          _1244_v83 = _rhs204;
          _lhs172.f14 = _rhs205;
          _1243_v82 = _rhs206;
          _lhs173.f3 = _rhs207;
          let _1245_v84;
          _1245_v84 = _module.D3.create_DC9();
          let _1246_v85;
          _1246_v85 = _dafny.Seq.of(_1245_v84);
          let _1247_v86;
          _1247_v86 = _module.D5.create_DC13(!((_this).fm1(_1233_v78, globalState)), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f20)), _1246_v85);
          let _1248_v87;
          let _nw197 = Array((new BigNumber(8)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = _module.__default.fm41(p2, (_this).f18, (_1231_v77).f29, _1226_v73, globalState);
          _nw197[(_dafny.ONE).toNumber()] = _module.__default.fm41(new BigNumber(((_1231_v77).f29).length), (_this).f18, _dafny.Seq.UnicodeFromString("c"), new _dafny.CodePoint('g'.codePointAt(0)), globalState);
          _nw197[(new BigNumber(2)).toNumber()] = _1247_v86;
          _nw197[(new BigNumber(3)).toNumber()] = _1247_v86;
          _nw197[(new BigNumber(4)).toNumber()] = _module.__default.fm41((_this).f20, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-486)), ((_1249_v73) => function (_1250_i9) {
            return _1249_v73;
          })(_1226_v73)), new _dafny.CodePoint('x'.codePointAt(0)), globalState);
          _nw197[(new BigNumber(5)).toNumber()] = _1247_v86;
          _nw197[(new BigNumber(6)).toNumber()] = _1247_v86;
          _nw197[(new BigNumber(7)).toNumber()] = _1247_v86;
          _1248_v87 = _nw197;
          let _1251_v88;
          _1251_v88 = _1246_v85;
          let _index205 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1248_v87).length));
          (_1248_v87)[_index205] = _module.D5.create_DC13(p1, p2, _1251_v88);
          let _1252_v89;
          _1252_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1231_v77).f30);
          let _1253_v90;
          let _nw198 = Array((new BigNumber(12)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = p2;
          _nw198[(_dafny.ONE).toNumber()] = (_this).fm0(p0, new _dafny.CodePoint('g'.codePointAt(0)), globalState);
          _nw198[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1252_v89).length));
          _nw198[(new BigNumber(3)).toNumber()] = (_this).f20;
          _nw198[(new BigNumber(4)).toNumber()] = (_this).f20;
          _nw198[(new BigNumber(5)).toNumber()] = p2;
          _nw198[(new BigNumber(6)).toNumber()] = new BigNumber(193);
          _nw198[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.update((_1231_v77).f29, _module.__default.safeIndex((_this).f20, new BigNumber(((_1231_v77).f29).length)), _1226_v73)).length);
          _nw198[(new BigNumber(8)).toNumber()] = p0;
          _nw198[(new BigNumber(9)).toNumber()] = (_this).f20;
          _nw198[(new BigNumber(10)).toNumber()] = p2;
          _nw198[(new BigNumber(11)).toNumber()] = new BigNumber(18);
          _1253_v90 = _nw198;
          let _1254_v91;
          _1254_v91 = _dafny.MultiSet.fromElements(_1253_v90, _1253_v90);
          let _pat_let_tv12 = _1254_v91;
          let _pat_let_tv13 = _1253_v90;
          let _pat_let_tv14 = _1254_v91;
          let _pat_let_tv15 = _1253_v90;
          let _pat_let_tv16 = _1251_v88;
          let _index206 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_1248_v87).length));
          (_1248_v87)[_index206] = function (_pat_let21_0) {
            return function (_1255_dt__update__tmp_h0) {
              return function (_pat_let22_0) {
                return function (_1256_dt__update_hcf24_h0) {
                  return function (_pat_let23_0) {
                    return function (_1257_dt__update_hcf25_h0) {
                      return _module.D5.create_DC13((_1255_dt__update__tmp_h0).dtor_cf23, _1256_dt__update_hcf24_h0, _1257_dt__update_hcf25_h0);
                    }(_pat_let23_0);
                  }(_pat_let_tv16);
                }(_pat_let22_0);
              }(_module.__default.safeModuloInt((_this).f20, (((_pat_let_tv14).contains(_pat_let_tv15)) ? ((_pat_let_tv12).get(_pat_let_tv13)) : (new BigNumber(913)))));
            }(_pat_let21_0);
          }(_1247_v86);
        } else {
          _1227_v74 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), ((_1258_p2, _1259_v74, _1260_v73) => function (_1261_i10) {
            return (_module.D0.create_DC1(_dafny.Seq.UnicodeFromString("gf"), _1258_p2, new BigNumber(-503), (_module.D0.create_DC1(_1259_v74, (_this).f20, _1258_p2, _1260_v73)).dtor_cf4)).dtor_cf4;
          })(p2, _1227_v74, _1226_v73));
          let _arr37 = _this.f19;
          let _index207 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length));
          _arr37[_index207] = !(p0).isEqualTo((_this).f20);
          let _arr38 = _this.f19;
          let _index208 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length));
          _arr38[_index208] = p1;
          let _1262_v92;
          _1262_v92 = _dafny.Set.fromElements(p1);
          let _1263_v93;
          _1263_v93 = _dafny.Map.Empty.slice().updateUnsafe(false,!(p1));
          let _1264_v94;
          _1264_v94 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1263_v93).length));
          let _1265_v95;
          _1265_v95 = _dafny.Seq.of((_this).f20);
          let _1266_v96;
          _1266_v96 = _dafny.Seq.of(p0, new BigNumber(((_1264_v94).update((_this).f18, new BigNumber((_dafny.Seq.of(_dafny.Seq.of((_this).f18, (_this).f18))).length))).length), new BigNumber((_1265_v95).length));
          let _1267_v97;
          _1267_v97 = _dafny.Seq.of(new BigNumber((_1227_v74).length), (_1266_v96)[_module.__default.safeIndex((_this).fm0((_this).f20, _1226_v73, globalState), new BigNumber((_1266_v96).length))]);
          let _1268_v98;
          let _nw199 = new _module.C4();
          _nw199.__ctor(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(241))).cardinality()), (_dafny.ZERO).minus((_this).fm0(new BigNumber((_1163_v18).length), _1226_v73, globalState)), _this.f17, !((_this).f18));
          _1268_v98 = _nw199;
          let _1269_v99;
          _1269_v99 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1268_v98,_module.__default.fm37((_this).f18, new BigNumber(384), p1, false, globalState))).length));
          let _1270_v100;
          let _init38 = function (_1271_i11) {
            return (_1271_i11).plus((_this).f20);
          };
          let _nw200 = Array((new BigNumber(6)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw200.length); _i0_38++) {
            _nw200[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1270_v100 = _nw200;
          let _1272_v101;
          _1272_v101 = _module.D0.create_DC2(p1, (_1268_v98).f25, _1270_v100);
          let _1273_v102;
          let _nw201 = Array((new BigNumber(20)).toNumber());
          _nw201[(_dafny.ZERO).toNumber()] = p1;
          _nw201[(_dafny.ONE).toNumber()] = p1;
          _nw201[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw201[(new BigNumber(3)).toNumber()] = false;
          _nw201[(new BigNumber(4)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(5)).toNumber()] = p1;
          _nw201[(new BigNumber(6)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(7)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(8)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(9)).toNumber()] = true;
          _nw201[(new BigNumber(10)).toNumber()] = true;
          _nw201[(new BigNumber(11)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(12)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(13)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(14)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(15)).toNumber()] = !(false);
          _nw201[(new BigNumber(16)).toNumber()] = p1;
          _nw201[(new BigNumber(17)).toNumber()] = (_this).f18;
          _nw201[(new BigNumber(18)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))];
          _nw201[(new BigNumber(19)).toNumber()] = (_1272_v101).dtor_cf5;
          _1273_v102 = _nw201;
          let _1274_v103;
          _1274_v103 = _dafny.Seq.of(_1273_v102, _1273_v102);
          let _1275_v104;
          let _nw202 = Array((new BigNumber(6)).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = new BigNumber((_1269_v99).length);
          _nw202[(_dafny.ONE).toNumber()] = new BigNumber((_1266_v96).length);
          _nw202[(new BigNumber(2)).toNumber()] = (_1268_v98).f25;
          _nw202[(new BigNumber(3)).toNumber()] = (_this).f20;
          _nw202[(new BigNumber(4)).toNumber()] = new BigNumber((_1274_v103).length);
          _nw202[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("uosnx")).length);
          _1275_v104 = _nw202;
          let _1276_v105;
          _1276_v105 = _module.D0.create_DC2((_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))], new BigNumber((_1262_v92).length), _1275_v104);
          let _arr39 = _this.f19;
          let _index209 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length));
          let _rhs208 = p2;
          let _rhs209 = (_1272_v101).dtor_cf7;
          let _rhs210 = _1262_v92;
          let _rhs211 = _1265_v95;
          let _rhs212 = !((_this.f19)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length))]);
          let _lhs174 = globalState;
          let _lhs175 = _this.f19;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length));
          r2 = _rhs208;
          _lhs174.f7 = _rhs209;
          _1262_v92 = _rhs210;
          _1267_v97 = _rhs211;
          _lhs175[_lhs176] = _rhs212;
          let _arr40 = _this.f19;
          let _index210 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_this.f19).length));
          _arr40[_index210] = (_this).f18;
          let _1277_v106;
          let _nw203 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1277_v106 = _nw203;
          let _index211 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1277_v106).length));
          (_1277_v106)[_index211] = _1226_v73;
          let _index212 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1277_v106).length));
          (_1277_v106)[_index212] = _1226_v73;
        }
        (globalState).f9 = (_this).f18;
        let _1278_v107;
        let _nw204 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1278_v107 = _nw204;
        let _index213 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1278_v107).length));
        (_1278_v107)[_index213] = (p0).minus(p0);
        let _1279_v108;
        _1279_v108 = _dafny.MultiSet.fromElements(p1);
        let _index214 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1278_v107).length));
        (_1278_v107)[_index214] = new BigNumber((_1279_v108).cardinality());
        let _1280_v109;
        _1280_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_this.f19);
        let _1281_v110;
        _1281_v110 = _module.D5.create_DC12(_1280_v109);
        let _1282_v111;
        _1282_v111 = _dafny.Seq.of(_1281_v110, _module.D5.create_DC12(_1280_v109));
        let _index215 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1278_v107).length));
        let _rhs213 = p1;
        let _rhs214 = (p0).plus(new BigNumber((_dafny.Seq.Concat(_1282_v111, _dafny.Seq.update(_1282_v111, _module.__default.safeIndex((_this).f20, new BigNumber((_1282_v111).length)), _1281_v110))).length));
        let _rhs215 = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus((_this).f20));
        let _rhs216 = _1226_v73;
        let _rhs217 = _1226_v73;
        let _lhs177 = globalState;
        let _lhs178 = _1278_v107;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1278_v107).length));
        _lhs177.f3 = _rhs213;
        _lhs178[_lhs179] = _rhs214;
        r2 = _rhs215;
        _1226_v73 = _rhs216;
        _1226_v73 = _rhs217;
      } else if (_source12.is_DC10) {
        let _1283___mcc_h4 = (_source12).cf16;
        let _1284___mcc_h5 = (_source12).cf17;
        let _1285___mcc_h6 = (_source12).cf18;
        let _1286___mcc_h7 = (_source12).cf19;
        let _1287___mcc_h8 = (_source12).cf20;
        let _1288_cf20 = _1287___mcc_h8;
        let _1289_cf19 = _1286___mcc_h7;
        let _1290_cf18 = _1285___mcc_h6;
        let _1291_cf17 = _1284___mcc_h5;
        let _1292_cf16 = _1283___mcc_h4;
        (globalState).f14 = _module.__default.safeModuloInt(new BigNumber(617), p2);
        let _1293_v112;
        let _nw205 = new _module.C0();
        _nw205.__ctor();
        _1293_v112 = _nw205;
        let _1294_v113;
        _1294_v113 = _dafny.Seq.of(_module.__default.fm30(p1, globalState), _1163_v18, _module.__default.fm30((_this).f18, globalState));
        let _1295_v114;
        _1295_v114 = _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber(((_1294_v113)[_module.__default.safeIndex(p2, new BigNumber((_1294_v113).length))]).length), new BigNumber(944)), p0, (_this).fm0((_this).f20, _1226_v73, globalState));
        _1288_cf20 = new BigNumber((_1295_v114).length);
        let _1296_v115;
        _1296_v115 = _dafny.Set.fromElements(_1288_cf20);
        let _1297_v116;
        _1297_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v18,(_dafny.ZERO).minus((_this).f20));
        let _1298_v117;
        _1298_v117 = _dafny.Seq.of((_this).fm1(_1297_v116, globalState));
        _1296_v115 = (_module.__default.fm42(_dafny.MultiSet.fromElements(true), p1, _1298_v117, globalState)).Intersect(_1296_v115);
      } else {
        let _1299___mcc_h9 = (_source12).cf15;
        let _1300_cf15 = _1299___mcc_h9;
        (globalState).f14 = (_this).f20;
        let _1301_v118;
        let _init39 = function (_1302_i12) {
          return _dafny.Seq.Concat(_dafny.Seq.of((_this).f18), _dafny.Seq.of((_this).f18));
        };
        let _nw206 = Array((new BigNumber(27)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw206.length); _i0_39++) {
          _nw206[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1301_v118 = _nw206;
        let _1303_v119;
        _1303_v119 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _1304_v120;
        _1304_v120 = _module.D9.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1303_v119).length),false), new BigNumber(-81), (_this).f18, _1163_v18, _1226_v73);
        let _1305_v121;
        _1305_v121 = _dafny.Seq.of(!(p1));
        let _1306_v122;
        _1306_v122 = _dafny.Set.fromElements((_this).f20, p2, p0);
        let _1307_v123;
        _1307_v123 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of(p1));
        let _nw207 = Array((new BigNumber(17)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p1, p1, p1, (_1304_v120).dtor_cf38);
        _nw207[(_dafny.ONE).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f18), _1305_v121);
        _nw207[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1305_v121, _1305_v121);
        _nw207[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1305_v121, _1305_v121);
        _nw207[(new BigNumber(5)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(6)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p1), _1305_v121);
        _nw207[(new BigNumber(8)).toNumber()] = (_this).fm25(!((_this).f18), p1, new BigNumber((_1306_v122).length), (_this).f18, globalState);
        _nw207[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1305_v121, _dafny.Seq.of(p1, p1));
        _nw207[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((_this).f18);
        _nw207[(new BigNumber(11)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(12)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(13)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(14)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(15)).toNumber()] = _1305_v121;
        _nw207[(new BigNumber(16)).toNumber()] = (((_1307_v123).contains(false)) ? ((_1307_v123).get(false)) : (_1305_v121));
        _1301_v118 = _nw207;
        let _1308_v124;
        let _nw208 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1308_v124 = _nw208;
        let _index216 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1308_v124).length));
        (_1308_v124)[_index216] = _this.f19;
        let _index217 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1308_v124).length));
        (_1308_v124)[_index217] = _this.f19;
        let _1309_v125;
        let _init40 = function (_1310_i13) {
          return (_1310_i13).plus(new BigNumber(647));
        };
        let _nw209 = Array((new BigNumber(26)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw209.length); _i0_40++) {
          _nw209[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1309_v125 = _nw209;
        (globalState).f7 = _1309_v125;
      }
      let _1311_v126;
      _1311_v126 = _dafny.Set.fromElements((new BigNumber(((_this).fm25((_this).f18, p1, p0, (_this).f18, globalState)).length)).isLessThanOrEqualTo(p0));
      let _1312_v127;
      _1312_v127 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v18,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,p0)).length));
      let _1313_v128;
      _1313_v128 = _dafny.Seq.of(_1311_v126, _dafny.Set.fromElements(p1), _1311_v126, _module.__default.fm37((_this).f18, (_this).f20, p1, (_this).fm1(_1312_v127, globalState), globalState));
      _1311_v126 = (_1313_v128)[_module.__default.safeIndex(p0, new BigNumber((_1313_v128).length))];
      let _1314_v129;
      _1314_v129 = _dafny.Seq.of((_this).f18);
      let _1315_v130;
      let _nw210 = new _module.C1();
      _nw210.__ctor();
      _1315_v130 = _nw210;
      let _1316_v131;
      _1316_v131 = _dafny.Set.fromElements(_1315_v130, _1315_v130);
      let _1317_v132;
      _1317_v132 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1316_v131);
      let _1318_v133;
      _1318_v133 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1227_v74, _module.__default.safeIndex(new BigNumber((_1314_v129).length), new BigNumber((_1227_v74).length)), _1226_v73),(((_1317_v132).contains((_this).f18)) ? ((_1317_v132).get((_this).f18)) : (_1316_v131)));
      let _1319_i14;
      _1319_i14 = _dafny.ZERO;
      L11: {
        while (!(_1318_v133).contains(_dafny.Seq.Concat(_1227_v74, _1227_v74))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1319_i14)) {
              break L11;
            }
            _1319_i14 = (_1319_i14).plus(_dafny.ONE);
            (globalState).f14 = p2;
            let _1320_v134;
            let _nw211 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1320_v134 = _nw211;
            let _index218 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1320_v134).length));
            (_1320_v134)[_index218] = _1226_v73;
            let _index219 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1320_v134).length));
            (_1320_v134)[_index219] = _1226_v73;
            r2 = p2;
            let _1321_v135;
            _1321_v135 = _module.D3.create_DC9();
            let _1322_v136;
            _1322_v136 = _dafny.Seq.of(_module.D3.create_DC9(), _1321_v135);
            let _1323_v137;
            _1323_v137 = _1322_v136;
            let _source13 = _module.D5.create_DC13(p1, ((_this).f20).multipliedBy(p2), _1323_v137);
            if (_source13.is_DC13) {
              let _1324___mcc_h10 = (_source13).cf23;
              let _1325___mcc_h11 = (_source13).cf24;
              let _1326___mcc_h12 = (_source13).cf25;
              let _1327_cf25 = _1326___mcc_h12;
              let _1328_cf24 = _1325___mcc_h11;
              let _1329_cf23 = _1324___mcc_h10;
              (globalState).f14 = new BigNumber(-602);
              let _1330_v138;
              _1330_v138 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20, p0, p2);
              (globalState).f14 = (((_1330_v138).contains((_this).f20)) ? ((_1330_v138).get((_this).f20)) : (p0));
              let _1331_v139;
              let _nw212 = new _module.C0();
              _nw212.__ctor();
              _1331_v139 = _nw212;
              let _1332_v140;
              _1332_v140 = _dafny.Seq.of(new BigNumber((_1132_v0).length));
              let _1333_v141;
              _1333_v141 = _dafny.Map.Empty.slice().updateUnsafe((_1332_v140)[_module.__default.safeIndex(p0, new BigNumber((_1332_v140).length))],_1331_v139);
              let _1334_v142;
              let _nw213 = new _module.C2();
              _nw213.__ctor(_this.f17, _1329_cf23);
              _1334_v142 = _nw213;
              let _1335_v143;
              _1335_v143 = _dafny.Seq.of(_1334_v142);
              let _1336_v144;
              _1336_v144 = _dafny.Seq.of(_1334_v142, _1334_v142, _1334_v142, (_1335_v143)[_module.__default.safeIndex(p2, new BigNumber((_1335_v143).length))]);
              let _1337_v145;
              _1337_v145 = _dafny.Seq.of(_1331_v139, _1331_v139);
              let _1338_v146;
              let _nw214 = Array((new BigNumber(29)).toNumber());
              _nw214[(_dafny.ZERO).toNumber()] = _1331_v139;
              _nw214[(_dafny.ONE).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(2)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(3)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(4)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(5)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(6)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(7)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(8)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(9)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(10)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(11)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(12)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(13)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(14)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(15)).toNumber()] = (((_1333_v141).contains(new BigNumber((_dafny.MultiSet.FromArray(_1335_v143)).cardinality()))) ? ((_1333_v141).get(new BigNumber((_dafny.MultiSet.FromArray(_1335_v143)).cardinality()))) : (_1331_v139));
              _nw214[(new BigNumber(16)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(17)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(18)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(19)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(20)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(21)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(22)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(23)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(24)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(25)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(26)).toNumber()] = (((_this).f18) ? (_1331_v139) : (_1331_v139));
              _nw214[(new BigNumber(27)).toNumber()] = _1331_v139;
              _nw214[(new BigNumber(28)).toNumber()] = (_1337_v145)[_module.__default.safeIndex(new BigNumber(-362), new BigNumber((_1337_v145).length))];
              _1338_v146 = _nw214;
              _1338_v146 = _1338_v146;
              let _1339_v147;
              _1339_v147 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1227_v74);
              let _1340_v148;
              let _nw215 = Array((new BigNumber(15)).toNumber());
              _nw215[(_dafny.ZERO).toNumber()] = p2;
              _nw215[(_dafny.ONE).toNumber()] = (_this).f20;
              _nw215[(new BigNumber(2)).toNumber()] = p2;
              _nw215[(new BigNumber(3)).toNumber()] = (_this).f20;
              _nw215[(new BigNumber(4)).toNumber()] = (_this).f20;
              _nw215[(new BigNumber(5)).toNumber()] = p2;
              _nw215[(new BigNumber(6)).toNumber()] = new BigNumber((_1339_v147).length);
              _nw215[(new BigNumber(7)).toNumber()] = new BigNumber(((_1311_v126).Intersect(_1311_v126)).length);
              _nw215[(new BigNumber(8)).toNumber()] = _1328_cf24;
              _nw215[(new BigNumber(9)).toNumber()] = (_this).f20;
              _nw215[(new BigNumber(10)).toNumber()] = _1328_cf24;
              _nw215[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1328_cf24));
              _nw215[(new BigNumber(12)).toNumber()] = (_1332_v140)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-78)), ((_1341_v74) => function (_1342_i15) {
                return _1341_v74;
              })(_1227_v74))).length), new BigNumber((_1332_v140).length))];
              _nw215[(new BigNumber(13)).toNumber()] = (_this).f20;
              _nw215[(new BigNumber(14)).toNumber()] = (((_1334_v142).f18) ? (p2) : ((_dafny.ZERO).minus(p0)));
              _1340_v148 = _nw215;
              let _index220 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_1340_v148).length));
              (_1340_v148)[_index220] = new BigNumber((_1227_v74).length);
              let _1343_v149;
              _1343_v149 = _dafny.MultiSet.fromElements(_1227_v74, _1227_v74, _1227_v74);
              let _index221 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_1340_v148).length));
              (_1340_v148)[_index221] = ((_this).f20).multipliedBy((new BigNumber(187)).minus(new BigNumber((_1343_v149).cardinality())));
            } else if (_source13.is_DC14) {
              let _1344___mcc_h13 = (_source13).cf26;
              let _1345___mcc_h14 = (_source13).cf27;
              let _1346___mcc_h15 = (_source13).cf28;
              let _1347_cf28 = _1346___mcc_h15;
              let _1348_cf27 = _1345___mcc_h14;
              let _1349_cf26 = _1344___mcc_h13;
              (_this).f19 = _this.f19;
              let _arr41 = _this.f19;
              let _index222 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_this.f19).length));
              _arr41[_index222] = (_this).f18;
              let _arr42 = _this.f19;
              let _index223 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_this.f19).length));
              _arr42[_index223] = true;
              let _1350_v150;
              _1350_v150 = _dafny.Seq.of(p2, (_this).f20);
              let _1351_v151;
              _1351_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1227_v74,_1350_v150);
              _1351_v151 = (_1351_v151).update(_dafny.Seq.Concat(_1227_v74, _dafny.Seq.UnicodeFromString("kla")), _dafny.Seq.Concat(_1350_v150, _dafny.Seq.of((_this).f20, p2)));
              let _1352_v152;
              _1352_v152 = _dafny.MultiSet.fromElements(_1349_cf26);
              (globalState).f14 = (_this).fm0(new BigNumber((_1352_v152).cardinality()), (_1320_v134)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_1320_v134).length))], globalState);
            } else {
              let _1353___mcc_h16 = (_source13).cf22;
              let _1354_cf22 = _1353___mcc_h16;
              (globalState).f14 = (_dafny.ZERO).minus(p2);
              let _1355_v153;
              let _nw216 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
              _1355_v153 = _nw216;
              let _1356_v154;
              _1356_v154 = _module.D7.create_DC17(_1355_v153, (_this).f18, p2);
              let _1357_v155;
              _1357_v155 = _dafny.Map.Empty.slice().updateUnsafe(_1321_v135,_1356_v154);
              let _1358_v156;
              let _init41 = ((_1359_v73, _1360_v134) => function (_1361_i16) {
                return (_1361_i16).plus(new BigNumber((_dafny.Set.fromElements(_1359_v73, (_1360_v134)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_1360_v134).length))])).length));
              })(_1226_v73, _1320_v134);
              let _nw217 = Array((new BigNumber(23)).toNumber());
              for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw217.length); _i0_41++) {
                _nw217[_i0_41] = _init41(new BigNumber(_i0_41));
              }
              _1358_v156 = _nw217;
              let _index224 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1358_v156).length));
              (_1358_v156)[_index224] = p0;
              let _index225 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1358_v156).length));
              (_1358_v156)[_index225] = (_this).f20;
              let _1362_v157;
              _1362_v157 = _dafny.Seq.of(_1357_v155);
              let _index226 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1358_v156).length));
              let _index227 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1358_v156).length));
              let _rhs218 = ((_1362_v157)[_module.__default.safeIndex(new BigNumber((_1314_v129).length), new BigNumber((_1362_v157).length))]).Merge(_1357_v155);
              let _rhs219 = (_dafny.ZERO).minus(p2);
              let _rhs220 = _module.__default.safeModuloInt(p0, p0);
              let _rhs221 = p0;
              let _lhs180 = globalState;
              let _lhs181 = _1358_v156;
              let _lhs182 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1358_v156).length));
              let _lhs183 = _1358_v156;
              let _lhs184 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1358_v156).length));
              _1357_v155 = _rhs218;
              _lhs180.f14 = _rhs219;
              _lhs181[_lhs182] = _rhs220;
              _lhs183[_lhs184] = _rhs221;
              let _1363_v158;
              let _nw218 = Array((new BigNumber(14)).toNumber()).fill([]);
              _1363_v158 = _nw218;
              let _index228 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1363_v158).length));
              (_1363_v158)[_index228] = _this.f19;
              let _index229 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_1363_v158).length));
              (_1363_v158)[_index229] = _this.f19;
              let _1364_v160;
              _1364_v160 = _dafny.Map.Empty.slice().updateUnsafe(p1,(function () {
                let _coll53 = new _dafny.Map();
                for (const _compr_53 of (_1132_v0).Keys.Elements) {
                  let _1365_v159 = _compr_53;
                  if ((_1132_v0).contains(_1365_v159)) {
                    _coll53.push([(_1365_v159).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_1366_i17) {
                      return new _dafny.CodePoint('q'.codePointAt(0));
                    })).length)),p1]);
                  }
                }
                return _coll53;
              }()).Merge(_1132_v0));
              _1364_v160 = (_1364_v160).update(((_this).f18) || ((_this).f18), (_1132_v0).Merge(_1132_v0));
            }
          }
        }
      }
      r0 = p1;
      r1 = p1;
      r2 = (_dafny.ZERO).minus(((new BigNumber(595)).multipliedBy(p0)).minus(((_this).f20).plus(p2)));
      return [r0, r1, r2];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f21 = _dafny.Map.Empty;
      this._f20 = _dafny.ZERO;
      this._f18 = false;
      this.f23 = _dafny.ZERO;
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f23, f24, f19, f17, f18, f20, f21) {
      let _this = this;
      (_this).f23 = f23;
      (_this)._f24 = f24;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,new BigNumber((_dafny.Seq.UnicodeFromString("fmocyx")).length))).length)).minus(_this.f23)).multipliedBy(new BigNumber(482));
    };
    fm1(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1367_i0;
      _1367_i0 = _dafny.ZERO;
      L12: {
        while ((_this).f24) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1367_i0)) {
              break L12;
            }
            _1367_i0 = (_1367_i0).plus(_dafny.ONE);
            let _1368_v0;
            _1368_v0 = _dafny.Seq.UnicodeFromString("fgnisgrc");
            let _1369_v1;
            _1369_v1 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).update((_this).f20, (_this).f20),_this.f23);
            _1368_v0 = (((_this).fm1(_1369_v1, globalState)) ? (_1368_v0) : (_dafny.Seq.UnicodeFromString("bkcjnysf")));
            (globalState).f3 = !((_this).f18);
            let _1370_v2;
            let _nw219 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
            _1370_v2 = _nw219;
            (globalState).f7 = _1370_v2;
            let _1371_v3;
            let _nw220 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _1371_v3 = _nw220;
            _1371_v3 = (((_this).f18) ? (_1371_v3) : (_1371_v3));
          }
        }
      }
      let _1372_v4;
      _1372_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_this.f19);
      let _1373_v5;
      let _nw221 = new _module.C2();
      _nw221.__ctor((_this.f17).Union(_dafny.Set.fromElements(_this.f19, (((_1372_v4).contains(p1)) ? ((_1372_v4).get(p1)) : (_this.f19)), _this.f19)), !(p1));
      _1373_v5 = _nw221;
      _1373_v5 = _1373_v5;
      (globalState).f9 = _module.__default.fm13(globalState);
      let _1374_v6;
      _1374_v6 = _dafny.Set.fromElements((_1373_v5).f18, (_1373_v5).f18);
      let _1375_v7;
      let _nw222 = new _module.C4();
      _nw222.__ctor(new BigNumber((_1374_v6).length), (_this).f20, _1373_v5.f17, (_this).f24);
      _1375_v7 = _nw222;
      r1 = (_this).f24;
      let _rhs222 = (_1373_v5).f18;
      let _rhs223 = (_1373_v5).f18;
      let _lhs185 = globalState;
      let _lhs186 = globalState;
      _lhs185.f9 = _rhs222;
      _lhs186.f9 = _rhs223;
      r0 = p1;
      r1 = ((_1374_v6).Union(_dafny.Set.fromElements(p1))).IsDisjointFrom((_1374_v6).Intersect(_1374_v6));
      let _1376_v8;
      let _nw223 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1376_v8 = _nw223;
      let _1377_v9;
      _1377_v9 = _dafny.MultiSet.fromElements(_1376_v8);
      r2 = (new BigNumber(((_1377_v9).Intersect(_1377_v9)).cardinality())).minus((_this).f20);
      return [r0, r1, r2];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f14 = _module.__default.safeDivisionInt(p2, (_dafny.ZERO).minus(new BigNumber(-181)));
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_this.f19).length))) {
        let _1378_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1378_i0)) && ((_1378_i0).isLessThan(new BigNumber((_this.f19).length))))) {
          let _arr43 = _this.f19;
          _arr43[(_1378_i0)] = (_this).f24;
        }
      }
      let _1379_i1;
      _1379_i1 = _dafny.ZERO;
      L13: {
        while (!((_this).f24)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1379_i1)) {
              break L13;
            }
            _1379_i1 = (_1379_i1).plus(_dafny.ONE);
            let _1380_v0;
            _1380_v0 = _dafny.Seq.UnicodeFromString("t");
            _1380_v0 = _1380_v0;
            let _1381_v1;
            let _nw224 = new _module.C7();
            _nw224.__ctor((_this).f20, _this.f21, _this.f19, _this.f17, !((_this).f18));
            _1381_v1 = _nw224;
            let _1382_v2;
            _1382_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1381_v1);
            let _1383_v3;
            _1383_v3 = _dafny.Seq.of(_1382_v2, _1382_v2, _1382_v2, _1382_v2, _1382_v2);
            let _1384_v4;
            _1384_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1383_v3).length),_1381_v1.f19);
            let _1385_v5;
            _1385_v5 = _module.D13.create_DC31(_1384_v4);
            let _1386_v6;
            let _nw225 = Array((new BigNumber(10)).toNumber());
            _nw225[(_dafny.ZERO).toNumber()] = _1384_v4;
            _nw225[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1381_v1.f19);
            _nw225[(new BigNumber(2)).toNumber()] = _1384_v4;
            _nw225[(new BigNumber(3)).toNumber()] = _1384_v4;
            _nw225[(new BigNumber(4)).toNumber()] = (_1384_v4).Merge(_1384_v4);
            _nw225[(new BigNumber(5)).toNumber()] = (_1385_v5).dtor_cf60;
            _nw225[(new BigNumber(6)).toNumber()] = (_1384_v4).Merge(_1384_v4);
            _nw225[(new BigNumber(7)).toNumber()] = _1384_v4;
            _nw225[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1381_v1).f20,_1381_v1.f19);
            _nw225[(new BigNumber(9)).toNumber()] = (_1384_v4).Merge(_1384_v4);
            _1386_v6 = _nw225;
            let _index230 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1386_v6).length));
            (_1386_v6)[_index230] = _1384_v4;
            let _index231 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1386_v6).length));
            (_1386_v6)[_index231] = _1384_v4;
            let _1387_v7;
            _1387_v7 = _dafny.Set.fromElements((_this).f18);
            let _1388_v8;
            _1388_v8 = _dafny.Seq.of(_1387_v7, _1387_v7, _1387_v7, _1387_v7);
            let _1389_v9;
            _1389_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1381_v1).f18,p0);
            let _1390_v10;
            _1390_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1381_v1).f20,new BigNumber((_dafny.Seq.UnicodeFromString("vxjbxgjo")).length));
            let _1391_v11;
            _1391_v11 = _dafny.Seq.of((_this).f18, p1);
            let _1392_v12;
            _1392_v12 = _dafny.MultiSet.fromElements(_1389_v9);
            let _1393_v13;
            let _nw226 = Array((new BigNumber(15)).toNumber());
            _nw226[(_dafny.ZERO).toNumber()] = p0;
            _nw226[(_dafny.ONE).toNumber()] = new BigNumber((_1390_v10).length);
            _nw226[(new BigNumber(2)).toNumber()] = new BigNumber((_1391_v11).length);
            _nw226[(new BigNumber(3)).toNumber()] = new BigNumber(-285);
            _nw226[(new BigNumber(4)).toNumber()] = p0;
            _nw226[(new BigNumber(5)).toNumber()] = new BigNumber((_1392_v12).cardinality());
            _nw226[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f20);
            _nw226[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.of(p1, (_this).f24, (_this).f18)).length);
            _nw226[(new BigNumber(8)).toNumber()] = (_this).f20;
            _nw226[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bvsdc")).length)));
            _nw226[(new BigNumber(10)).toNumber()] = p0;
            _nw226[(new BigNumber(11)).toNumber()] = p0;
            _nw226[(new BigNumber(12)).toNumber()] = p2;
            _nw226[(new BigNumber(13)).toNumber()] = _this.f23;
            _nw226[(new BigNumber(14)).toNumber()] = (_this).f20;
            _1393_v13 = _nw226;
            let _1394_v14;
            _1394_v14 = _module.D0.create_DC2(p1, p0, _1393_v13);
            let _1395_v16;
            _1395_v16 = _dafny.MultiSet.fromElements((_1381_v1).f20);
            let _1396_v17;
            let _nw227 = Array((new BigNumber(8)).toNumber());
            _nw227[(_dafny.ZERO).toNumber()] = _module.__default.fm7(new BigNumber(916), _1389_v9, _module.D0.create_DC0((_1394_v14).dtor_cf5), globalState);
            _nw227[(_dafny.ONE).toNumber()] = (_this).f20;
            _nw227[(new BigNumber(2)).toNumber()] = (new BigNumber((_dafny.Seq.of(p1, p3)).length)).minus((_dafny.ZERO).minus(new BigNumber((_1389_v9).length)));
            _nw227[(new BigNumber(3)).toNumber()] = p0;
            _nw227[(new BigNumber(4)).toNumber()] = p2;
            _nw227[(new BigNumber(5)).toNumber()] = new BigNumber(((function () {
              let _coll54 = new _dafny.Set();
              for (const _compr_54 of _dafny.IntegerRange(new BigNumber(60), new BigNumber(957))) {
                let _1397_v15 = _compr_54;
                if (((new BigNumber(60)).isLessThanOrEqualTo(_1397_v15)) && ((_1397_v15).isLessThan(new BigNumber(957)))) {
                  _coll54.add((_1397_v15).multipliedBy(new BigNumber((_dafny.Set.fromElements(p2, (_this).f20)).length)));
                }
              }
              return _coll54;
            }()).Difference(_dafny.Set.fromElements(new BigNumber(((_1395_v16).update((_this).f20, _module.__default.abs(_this.f23))).cardinality())))).length);
            _nw227[(new BigNumber(6)).toNumber()] = (_this.f23).plus((_this).f20);
            _nw227[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p2);
            _1396_v17 = _nw227;
            let _rhs224 = ((((_this).f24) ? (_1387_v7) : ((_1388_v8)[_module.__default.safeIndex(_this.f23, new BigNumber((_1388_v8).length))]))).Union(_1387_v7);
            let _rhs225 = (_this).f18;
            let _rhs226 = _1396_v17;
            let _rhs227 = (_this).f18;
            let _lhs187 = globalState;
            let _lhs188 = globalState;
            let _lhs189 = globalState;
            _1387_v7 = _rhs224;
            _lhs187.f9 = _rhs225;
            _lhs188.f7 = _rhs226;
            _lhs189.f9 = _rhs227;
            let _1398_v18;
            let _nw228 = new _module.C0();
            _nw228.__ctor();
            _1398_v18 = _nw228;
          }
        }
      }
      let _arr44 = _this.f19;
      let _index232 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length));
      _arr44[_index232] = (_module.__default.fm29(globalState)).IsSubsetOf(_dafny.MultiSet.fromElements(p3));
      let _1399_v19;
      _1399_v19 = _module.D5.create_DC14(!(false), _this.f23, p2);
      let _1400_v20;
      _1400_v20 = _dafny.Seq.of(p0, p0);
      let _1401_v21;
      _1401_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,((_1399_v19).dtor_cf27).plus(new BigNumber((_1400_v20).length)));
      let _1402_v22;
      _1402_v22 = _dafny.Seq.of(false);
      let _1403_v23;
      _1403_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
      let _arr45 = _this.f19;
      let _index233 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length));
      let _rhs228 = ((_this).f20).isLessThan((new BigNumber((_1402_v22).length)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("dn")).length)));
      let _rhs229 = ((p3) ? ((_this.f23).isLessThanOrEqualTo(p2)) : ((_this).f24));
      let _rhs230 = (_dafny.Map.Empty.slice().updateUnsafe((((_1403_v23).contains(_this.f23)) ? ((_1403_v23).get(_this.f23)) : (p3)),_this.f23)).Merge(_1401_v21);
      let _lhs190 = _this.f19;
      let _lhs191 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length));
      let _lhs192 = globalState;
      _lhs190[_lhs191] = _rhs228;
      _lhs192.f9 = _rhs229;
      _1401_v21 = _rhs230;
      let _1404_i2;
      _1404_i2 = _dafny.ZERO;
      L14: {
        while ((_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))]) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1404_i2)) {
              break L14;
            }
            _1404_i2 = (_1404_i2).plus(_dafny.ONE);
            if ((_this).f18) {
              let _1405_v24;
              _1405_v24 = new _dafny.CodePoint('d'.codePointAt(0));
              _1405_v24 = _1405_v24;
              let _1406_v25;
              let _init42 = function (_1407_i3) {
                return (_this).f18;
              };
              let _nw229 = Array((new BigNumber(14)).toNumber());
              for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw229.length); _i0_42++) {
                _nw229[_i0_42] = _init42(new BigNumber(_i0_42));
              }
              _1406_v25 = _nw229;
              let _1408_v26;
              let _nw230 = new _module.C7();
              _nw230.__ctor((_this).f20, _this.f21, _1406_v25, _this.f17, (_this).f24);
              _1408_v26 = _nw230;
              _1408_v26 = _1408_v26;
              let _arr46 = _this.f19;
              let _index234 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length));
              _arr46[_index234] = p1;
              let _1409_v27;
              _1409_v27 = _dafny.Seq.UnicodeFromString("kme");
              let _1410_v28;
              let _nw231 = new _module.C4();
              _nw231.__ctor(new BigNumber((_1403_v23).length), p2, _this.f17, p1);
              _1410_v28 = _nw231;
              let _1411_v29;
              _1411_v29 = _dafny.Seq.of(_1410_v28, _1410_v28);
              let _1412_v30;
              let _nw232 = Array((new BigNumber(4)).toNumber());
              _nw232[(_dafny.ZERO).toNumber()] = (p2).minus(p2);
              _nw232[(_dafny.ONE).toNumber()] = (p2).minus(new BigNumber((_1409_v27).length));
              _nw232[(new BigNumber(2)).toNumber()] = (_1400_v20)[_module.__default.safeIndex(new BigNumber((_1411_v29).length), new BigNumber((_1400_v20).length))];
              _nw232[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).f20, (_this).f20);
              _1412_v30 = _nw232;
              let _index235 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1412_v30).length));
              (_1412_v30)[_index235] = p2;
              let _index236 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1412_v30).length));
              (_1412_v30)[_index236] = (_dafny.ZERO).minus(p0);
              _1403_v23 = (_1403_v23).update((_1412_v30)[_module.__default.safeIndex(new BigNumber(668), new BigNumber((_1412_v30).length))], (((_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))]) ? (p3) : (p3)));
            } else {
              let _1413_v31;
              _1413_v31 = new _dafny.CodePoint('w'.codePointAt(0));
              let _rhs231 = ((_this).fm0((_dafny.ZERO).minus(p0), _1413_v31, globalState)).multipliedBy(p0);
              let _rhs232 = p0;
              let _lhs193 = globalState;
              let _lhs194 = globalState;
              _lhs193.f14 = _rhs231;
              _lhs194.f14 = _rhs232;
              (globalState).f14 = ((_this).f20).plus(p2);
              let _1414_v32;
              let _init43 = ((_1415_p2) => function (_1416_i4) {
                return (_1416_i4).minus(_1415_p2);
              })(p2);
              let _nw233 = Array((new BigNumber(16)).toNumber());
              for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw233.length); _i0_43++) {
                _nw233[_i0_43] = _init43(new BigNumber(_i0_43));
              }
              _1414_v32 = _nw233;
              let _index237 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1414_v32).length));
              (_1414_v32)[_index237] = _this.f23;
              let _index238 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1414_v32).length));
              (_1414_v32)[_index238] = (((_this).f20).minus(p0)).multipliedBy(new BigNumber(586));
              let _index239 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1414_v32).length));
              (_1414_v32)[_index239] = p2;
              (globalState).f3 = !_dafny.Seq.contains(_module.__default.fm16(globalState), true);
            }
            r0 = _module.__default.safeModuloInt((_this).f20, ((_dafny.ZERO).minus(new BigNumber((_1403_v23).length))).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,(_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))])).length)));
            let _1417_v33;
            _1417_v33 = _dafny.MultiSet.fromElements(p1);
            let _1418_v34;
            _1418_v34 = _dafny.Seq.UnicodeFromString("pinsphqf");
            let _1419_v35;
            let _nw234 = Array((new BigNumber(4)).toNumber());
            _nw234[(_dafny.ZERO).toNumber()] = _dafny.Seq.contains(_1418_v34, new _dafny.CodePoint('g'.codePointAt(0)));
            _nw234[(_dafny.ONE).toNumber()] = !((_this).f24);
            _nw234[(new BigNumber(2)).toNumber()] = _module.__default.fm13(globalState);
            _nw234[(new BigNumber(3)).toNumber()] = p1;
            _1419_v35 = _nw234;
            let _1420_v36;
            _1420_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
            let _rhs233 = _module.__default.fm29(globalState);
            let _rhs234 = (_this).f20;
            let _rhs235 = !((((_1420_v36).contains(p1)) ? ((_1420_v36).get(p1)) : ((_this).f18)));
            let _rhs236 = _1419_v35;
            let _lhs195 = globalState;
            _1417_v33 = _rhs233;
            r0 = _rhs234;
            _lhs195.f3 = _rhs235;
            _1419_v35 = _rhs236;
            let _1421_v37;
            let _nw235 = new _module.C4();
            _nw235.__ctor(p2, new BigNumber(((_1401_v21).update(p1, p2)).length), _this.f17, p3);
            _1421_v37 = _nw235;
            let _1422_v38;
            _1422_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1421_v37);
            _1422_v38 = (_1422_v38).update((p0).multipliedBy((_this).f20), _1421_v37);
          }
        }
      }
      let _1423_v39;
      _1423_v39 = _dafny.MultiSet.fromElements(new BigNumber((_1401_v21).length));
      let _1424_v40;
      _1424_v40 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1425_v41;
      _1425_v41 = _dafny.Seq.of(_1424_v40);
      let _1426_v42;
      _1426_v42 = _module.D3.create_DC9();
      let _1427_v44;
      _1427_v44 = _dafny.Set.fromElements(_1403_v23, function () {
        let _coll55 = new _dafny.Map();
        for (const _compr_55 of _dafny.IntegerRange(new BigNumber(662), new BigNumber(960))) {
          let _1428_v43 = _compr_55;
          if (((new BigNumber(662)).isLessThanOrEqualTo(_1428_v43)) && ((_1428_v43).isLessThan(new BigNumber(960)))) {
            _coll55.push([(_1428_v43).minus(p0),(_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))]]);
          }
        }
        return _coll55;
      }());
      let _1429_v45;
      _1429_v45 = _module.D0.create_DC0(true);
      let _1430_v46;
      _1430_v46 = _module.D1.create_DC4(!(true), _1427_v44, _module.__default.fm7((_this).f20, (_1401_v21).update((_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))], p2), _1429_v45, globalState));
      let _1431_v47;
      _1431_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1402_v22)[_module.__default.safeIndex(_this.f23, new BigNumber((_1402_v22).length))]);
      let _1432_v48;
      _1432_v48 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1424_v40);
      let _1433_v49;
      _1433_v49 = _dafny.Set.fromElements(p3);
      let _1434_v50;
      _1434_v50 = _dafny.MultiSet.fromElements(_1424_v40, (((_1432_v48).contains(new BigNumber((_1433_v49).length))) ? ((_1432_v48).get(new BigNumber((_1433_v49).length))) : (_1424_v40)));
      let _1435_v51;
      _1435_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1402_v22);
      let _1436_v52;
      _1436_v52 = _dafny.Set.fromElements(new BigNumber(((((_1435_v51).contains(new BigNumber(123))) ? ((_1435_v51).get(new BigNumber(123))) : (_1402_v22))).length));
      let _pat_let_tv17 = p1;
      let _1437_v53;
      let _nw236 = Array((new BigNumber(24)).toNumber());
      _nw236[(_dafny.ZERO).toNumber()] = new BigNumber(((_1423_v39).Intersect(_1423_v39)).cardinality());
      _nw236[(_dafny.ONE).toNumber()] = new BigNumber((_1402_v22).length);
      _nw236[(new BigNumber(2)).toNumber()] = new BigNumber((_1425_v41).length);
      _nw236[(new BigNumber(3)).toNumber()] = _this.f23;
      _nw236[(new BigNumber(4)).toNumber()] = p0;
      _nw236[(new BigNumber(5)).toNumber()] = p0;
      _nw236[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm34(_1426_v42, (_this).f18, _1424_v40, p1, globalState)).length);
      _nw236[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus(p2)).multipliedBy((function (_pat_let24_0) {
        return function (_1438_dt__update__tmp_h0) {
          return function (_pat_let25_0) {
            return function (_1439_dt__update_hcf9_h0) {
              return _module.D1.create_DC4(_1439_dt__update_hcf9_h0, (_1438_dt__update__tmp_h0).dtor_cf10, (_1438_dt__update__tmp_h0).dtor_cf11);
            }(_pat_let25_0);
          }(_pat_let_tv17);
        }(_pat_let24_0);
      }(_1430_v46)).dtor_cf11);
      _nw236[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_this.f23);
      _nw236[(new BigNumber(9)).toNumber()] = p2;
      _nw236[(new BigNumber(10)).toNumber()] = new BigNumber((_1431_v47).length);
      _nw236[(new BigNumber(11)).toNumber()] = new BigNumber((_1434_v50).cardinality());
      _nw236[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,!((_this.f19)[_module.__default.safeIndex(new BigNumber(17), new BigNumber((_this.f19).length))]))).length);
      _nw236[(new BigNumber(13)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_1424_v40)).Merge(_1432_v48)).length);
      _nw236[(new BigNumber(14)).toNumber()] = _this.f23;
      _nw236[(new BigNumber(15)).toNumber()] = (_this).f20;
      _nw236[(new BigNumber(16)).toNumber()] = p2;
      _nw236[(new BigNumber(17)).toNumber()] = (_this).f20;
      _nw236[(new BigNumber(18)).toNumber()] = (_this).f20;
      _nw236[(new BigNumber(19)).toNumber()] = _this.f23;
      _nw236[(new BigNumber(20)).toNumber()] = new BigNumber((((!((_this).f18)) ? (_1400_v20) : (_1400_v20))).length);
      _nw236[(new BigNumber(21)).toNumber()] = new BigNumber((_1436_v52).length);
      _nw236[(new BigNumber(22)).toNumber()] = (p0).multipliedBy((_this).f20);
      _nw236[(new BigNumber(23)).toNumber()] = p0;
      _1437_v53 = _nw236;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1437_v53).length))) {
        let _1440_i5 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1440_i5)) && ((_1440_i5).isLessThan(new BigNumber((_1437_v53).length))))) {
          (_1437_v53)[(_1440_i5)] = (_1440_i5).multipliedBy(p2);
        }
      }
      r0 = p2;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1441_v0;
      _1441_v0 = _dafny.Set.fromElements(_this.f23);
      let _1442_v1;
      _1442_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).f24);
      let _1443_v2;
      let _nw237 = new _module.C3();
      _nw237.__ctor(_this.f19, _dafny.Set.fromElements(_this.f19, _this.f19), !((((_1442_v1).contains(p0)) ? ((_1442_v1).get(p0)) : (false))));
      _1443_v2 = _nw237;
      let _1444_v3;
      _1444_v3 = _dafny.MultiSet.fromElements(_1443_v2, _1443_v2);
      let _1445_v4;
      _1445_v4 = _dafny.Seq.of((_1443_v2).f18, _module.__default.fm13(globalState));
      let _1446_v5;
      _1446_v5 = _dafny.Seq.UnicodeFromString("bwffq");
      let _1447_v6;
      _1447_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,new BigNumber((_1446_v5).length));
      let _1448_v7;
      _1448_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1444_v3).cardinality()),(_1445_v4)[_module.__default.safeIndex(new BigNumber((_1447_v6).length), new BigNumber((_1445_v4).length))]);
      let _arr47 = _this.f19;
      let _index240 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length));
      _arr47[_index240] = !(_1441_v0).equals(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1442_v1).length))));
      let _arr48 = _this.f19;
      let _index241 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length));
      _arr48[_index241] = (_1443_v2).f18;
      let _1449_i0;
      _1449_i0 = _dafny.ZERO;
      L15: {
        while (true) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1449_i0)) {
              break L15;
            }
            _1449_i0 = (_1449_i0).plus(_dafny.ONE);
            let _init44 = function (_1450_i1) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            };
            let _nw238 = Array((new BigNumber(3)).toNumber());
            for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw238.length); _i0_44++) {
              _nw238[_i0_44] = _init44(new BigNumber(_i0_44));
            }
            (globalState).f16 = _nw238;
            if ((_1443_v2).f18) {
              let _1451_v8;
              let _nw239 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1451_v8 = _nw239;
              let _1452_v9;
              _1452_v9 = new _dafny.CodePoint('y'.codePointAt(0));
              let _index242 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1451_v8).length));
              (_1451_v8)[_index242] = _1452_v9;
              let _index243 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_1451_v8).length));
              (_1451_v8)[_index243] = _1452_v9;
              let _1453_v10;
              let _nw240 = new _module.C3();
              _nw240.__ctor(_1443_v2.f19, _this.f17, true);
              _1453_v10 = _nw240;
              let _1454_v11;
              _1454_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_1453_v10);
              let _1455_v12;
              _1455_v12 = _module.D3.create_DC9();
              let _1456_v13;
              _1456_v13 = _dafny.Seq.of(_1455_v12, _1455_v12, _1455_v12);
              let _1457_v14;
              _1457_v14 = _module.D5.create_DC13(!((_1443_v2).f18), new BigNumber((_1454_v11).length), _1456_v13);
              _1457_v14 = _1457_v14;
              let _1458_v15;
              let _nw241 = new _module.C2();
              _nw241.__ctor(_1443_v2.f17, true);
              _1458_v15 = _nw241;
              let _1459_v16;
              _1459_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this.f23).minus(new BigNumber((_dafny.Seq.UnicodeFromString("oed")).length)),_1458_v15);
              _1459_v16 = _1459_v16;
              let _1460_v17;
              let _nw242 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
              _1460_v17 = _nw242;
              let _index244 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1460_v17).length));
              (_1460_v17)[_index244] = _this.f23;
              let _index245 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1460_v17).length));
              (_1460_v17)[_index245] = (_dafny.ZERO).minus(p0);
              _1460_v17 = _1460_v17;
            } else {
              (globalState).f14 = _this.f23;
              let _1461_v18;
              let _nw243 = new _module.C2();
              _nw243.__ctor((_this.f17).Intersect(_1443_v2.f17), ((_this).f20).isEqualTo((_this).f20));
              _1461_v18 = _nw243;
              let _1462_v19;
              _1462_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))],(_this).f20);
              let _1463_v20;
              _1463_v20 = _dafny.Seq.of((_1462_v19).Merge(_1462_v19));
              let _1464_v21;
              _1464_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v6,(_this).f20);
              let _rhs237 = (_this).fm1(_1464_v21, globalState);
              let _rhs238 = _1463_v20;
              let _lhs196 = globalState;
              _lhs196.f9 = _rhs237;
              _1463_v20 = _rhs238;
              (_1443_v2).f19 = _1443_v2.f19;
              let _1465_v22;
              _1465_v22 = _dafny.MultiSet.fromElements((_this).f24);
              let _1466_v23;
              _1466_v23 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f18);
              let _1467_v24;
              _1467_v24 = _module.D3.create_DC10(_1446_v5, (_1443_v2).f18, _dafny.Seq.UnicodeFromString("te"), _1446_v5, new BigNumber((_1466_v23).length));
              let _1468_v25;
              _1468_v25 = _module.D11.create_DC26((_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))], new BigNumber((_1446_v5).length));
              let _1469_v26;
              let _nw244 = Array((new BigNumber(9)).toNumber());
              _nw244[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements((_1443_v2).f18);
              _nw244[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).Union(_1465_v22);
              _nw244[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements((_1467_v24).dtor_cf17, true, (_1443_v2).f18, (_1443_v2).f18);
              _nw244[(new BigNumber(3)).toNumber()] = _1465_v22;
              _nw244[(new BigNumber(4)).toNumber()] = _1465_v22;
              _nw244[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.FromArray(_1445_v4)).update((_1443_v2).f18, _module.__default.abs(_this.f23));
              _nw244[(new BigNumber(6)).toNumber()] = (_module.__default.fm29(globalState)).Intersect(_1465_v22);
              _nw244[(new BigNumber(7)).toNumber()] = (_1465_v22).Difference(_1465_v22);
              _nw244[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_1468_v25).dtor_cf48);
              _1469_v26 = _nw244;
              let _index246 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1469_v26).length));
              (_1469_v26)[_index246] = _1465_v22;
              let _1470_v27;
              _1470_v27 = _dafny.Seq.of(_module.__default.fm29(globalState), _module.__default.fm29(globalState));
              let _index247 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1469_v26).length));
              (_1469_v26)[_index247] = (_1470_v27)[_module.__default.safeIndex(_this.f23, new BigNumber((_1470_v27).length))];
            }
            _1442_v1 = (_1442_v1).update((_this).f20, (_1443_v2).f18);
            let _1471_v28;
            _1471_v28 = new _dafny.CodePoint('h'.codePointAt(0));
            let _1472_v29;
            _1472_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_dafny.Set.fromElements(new BigNumber((_1447_v6).length), _this.f23));
            let _1473_v30;
            _1473_v30 = _dafny.Set.fromElements((_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))], (_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))]);
            let _1474_v31;
            _1474_v31 = _dafny.Seq.of((_this).fm0(new BigNumber((_1447_v6).length), _1471_v28, globalState), _module.__default.safeDivisionInt((_this).f20, (_this).f20), _module.__default.safeModuloInt((_this).f20, new BigNumber(((((_1472_v29).contains(p0)) ? ((_1472_v29).get(p0)) : (_dafny.Set.fromElements(new BigNumber((_1473_v30).length))))).length)));
            let _1475_v32;
            _1475_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p0);
            let _1476_v33;
            _1476_v33 = _module.D0.create_DC0(_module.__default.fm13(globalState));
            let _1477_v34;
            _1477_v34 = _module.D3.create_DC9();
            let _1478_v35;
            _1478_v35 = _dafny.Seq.of(_1477_v34, _1477_v34);
            let _1479_v36;
            _1479_v36 = _1478_v35;
            let _1480_v38;
            let _nw245 = Array((new BigNumber(16)).toNumber());
            _nw245[(_dafny.ZERO).toNumber()] = _this.f23;
            _nw245[(_dafny.ONE).toNumber()] = _module.__default.fm7((_this).f20, _1475_v32, _1476_v33, globalState);
            _nw245[(new BigNumber(2)).toNumber()] = (_this).f20;
            _nw245[(new BigNumber(3)).toNumber()] = (p0).minus(new BigNumber(696));
            _nw245[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f20), (_this).f20);
            _nw245[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_this.f23, new BigNumber(440));
            _nw245[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw245[(new BigNumber(7)).toNumber()] = new BigNumber((_1473_v30).length);
            _nw245[(new BigNumber(8)).toNumber()] = (_module.D5.create_DC13((_this).f24, p0, _1479_v36)).dtor_cf24;
            _nw245[(new BigNumber(9)).toNumber()] = p0;
            _nw245[(new BigNumber(10)).toNumber()] = p0;
            _nw245[(new BigNumber(11)).toNumber()] = (((_1475_v32).contains((_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))])) ? ((_1475_v32).get((_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))])) : (p0));
            _nw245[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of (_1442_v1).Keys.Elements) {
                let _1481_v37 = _compr_56;
                if ((_1442_v1).contains(_1481_v37)) {
                  _coll56.push([(_1481_v37).minus((_dafny.ZERO).minus(p0)),_1471_v28]);
                }
              }
              return _coll56;
            }()).length), new BigNumber((_1445_v4).length));
            _nw245[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(p0, _this.f23);
            _nw245[(new BigNumber(14)).toNumber()] = (_this).f20;
            _nw245[(new BigNumber(15)).toNumber()] = (((_1475_v32).contains((_1443_v2).f18)) ? ((_1475_v32).get((_1443_v2).f18)) : ((_dafny.ZERO).minus(_this.f23)));
            _1480_v38 = _nw245;
            let _index248 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1480_v38).length));
            (_1480_v38)[_index248] = (_this.f23).plus(_this.f23);
            let _index249 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1480_v38).length));
            let _rhs239 = _dafny.Seq.Concat(_1474_v31, _1474_v31);
            let _rhs240 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f20), (_this).f20);
            let _rhs241 = p0;
            let _lhs197 = _1480_v38;
            let _lhs198 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_1480_v38).length));
            let _lhs199 = globalState;
            _1474_v31 = _rhs239;
            _lhs197[_lhs198] = _rhs240;
            _lhs199.f14 = _rhs241;
          }
        }
      }
      let _1482_v39;
      _1482_v39 = _module.D11.create_DC27(_dafny.Seq.update(_1445_v4, _module.__default.safeIndex(_this.f23, new BigNumber((_1445_v4).length)), (_this.f19)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_this.f19).length))]), p0, (_this).f20);
      let _1483_v40;
      _1483_v40 = _dafny.Seq.of(_1448_v7);
      let _1484_v41;
      _1484_v41 = _dafny.MultiSet.fromElements(p0);
      let _1485_v42;
      _1485_v42 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1486_v43;
      _1486_v43 = _dafny.Set.fromElements(_1442_v1);
      let _1487_v44;
      _1487_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,false);
      let _1488_v45;
      _1488_v45 = _module.D1.create_DC4((_this).f18, _1486_v43, new BigNumber((_1487_v44).length));
      let _pat_let_tv18 = _1486_v43;
      let _1489_v46;
      let _nw246 = Array((new BigNumber(16)).toNumber());
      _nw246[(_dafny.ZERO).toNumber()] = (p0).plus((_dafny.ZERO).minus(p0));
      _nw246[(_dafny.ONE).toNumber()] = p0;
      _nw246[(new BigNumber(2)).toNumber()] = _this.f23;
      _nw246[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((((_1447_v6).contains(_this.f23)) ? ((_1447_v6).get(_this.f23)) : ((_this).f20)));
      _nw246[(new BigNumber(4)).toNumber()] = (_1482_v39).dtor_cf52;
      _nw246[(new BigNumber(5)).toNumber()] = (_this).f20;
      _nw246[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw246[(new BigNumber(7)).toNumber()] = (p0).plus((_this).f20);
      _nw246[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(588), new BigNumber(-132));
      _nw246[(new BigNumber(9)).toNumber()] = new BigNumber(323);
      _nw246[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1483_v40)[_module.__default.safeIndex((_this).f20, new BigNumber((_1483_v40).length))],(_this).f20)).length);
      _nw246[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1484_v41).cardinality()), (_this).f20);
      _nw246[(new BigNumber(12)).toNumber()] = p0;
      _nw246[(new BigNumber(13)).toNumber()] = _this.f23;
      _nw246[(new BigNumber(14)).toNumber()] = (_this).fm0(_this.f23, _1485_v42, globalState);
      _nw246[(new BigNumber(15)).toNumber()] = (((function (_pat_let26_0) {
        return function (_1490_dt__update__tmp_h0) {
          return function (_pat_let27_0) {
            return function (_1491_dt__update_hcf10_h0) {
              return _module.D1.create_DC4((_1490_dt__update__tmp_h0).dtor_cf9, _1491_dt__update_hcf10_h0, (_1490_dt__update__tmp_h0).dtor_cf11);
            }(_pat_let27_0);
          }(_pat_let_tv18);
        }(_pat_let26_0);
      }(_1488_v45)).dtor_cf9) ? ((_this).f20) : (new BigNumber((_1441_v0).length)));
      _1489_v46 = _nw246;
      let _index250 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_1489_v46).length));
      (_1489_v46)[_index250] = new BigNumber(358);
      let _index251 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_1489_v46).length));
      (_1489_v46)[_index251] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1446_v5, _1446_v5)).length));
      let _rhs242 = (_this).f24;
      let _rhs243 = _1445_v4;
      let _lhs200 = globalState;
      _lhs200.f3 = _rhs242;
      _1445_v4 = _rhs243;
      let _1492_v47;
      _1492_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1485_v42,(_1489_v46)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_1489_v46).length))]);
      let _1493_v48;
      _1493_v48 = _dafny.Map.Empty.slice().updateUnsafe((_1492_v47).Merge(_1492_v47),_1487_v44);
      let _1494_v50;
      _1494_v50 = _dafny.Set.fromElements(_1485_v42);
      _1493_v48 = (_1493_v48).update(function () {
        let _coll57 = new _dafny.Map();
        for (const _compr_57 of (_1494_v50).Elements) {
          let _1495_v49 = _compr_57;
          if ((_1494_v50).contains(_1495_v49)) {
            _coll57.push([_1495_v49,(_module.D5.create_DC14((_1443_v2).f18, new BigNumber((_1446_v5).length), p0)).dtor_cf28]);
          }
        }
        return _coll57;
      }(), _1487_v44);
      let _1496_v51;
      let _nw247 = Array((new BigNumber(27)).toNumber()).fill([]);
      _1496_v51 = _nw247;
      let _index252 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1496_v51).length));
      (_1496_v51)[_index252] = _1489_v46;
      let _index253 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1496_v51).length));
      (_1496_v51)[_index253] = _1489_v46;
      r0 = _this.f23;
      r1 = !((_this).f18);
      r2 = _module.__default.fm4(((_1489_v46)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_1489_v46).length))]).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), ((_1497_v42) => function (_1498_i2) {
        return _1497_v42;
      })(_1485_v42))).length)), globalState);
      return [r0, r1, r2];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f17 = _dafny.Set.Empty;
      this._f19 = [];
      this._f21 = _dafny.Map.Empty;
      this._f20 = _dafny.ZERO;
      this._f18 = false;
      this.f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    set f19(value) {
      let _this = this;
      _this._f19 = value;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f22, f20, f21, f19, f17, f18) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18)).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f18)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18)))).length));
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber(770)).minus((_this).f20);
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1499_v0;
      _1499_v0 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1500_v1;
      _1500_v1 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("x"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("x")).length)), _1499_v0));
      _1500_v1 = _1500_v1;
      let _1501_v2;
      _1501_v2 = _dafny.Seq.of(!((_this).f18));
      let _1502_v3;
      _1502_v3 = _dafny.Seq.of(_1501_v2, _1501_v2);
      let _1503_v4;
      let _nw248 = Array((new BigNumber(6)).toNumber());
      _nw248[(_dafny.ZERO).toNumber()] = _1502_v3;
      _nw248[(_dafny.ONE).toNumber()] = _1502_v3;
      _nw248[(new BigNumber(2)).toNumber()] = _1502_v3;
      _nw248[(new BigNumber(3)).toNumber()] = _1502_v3;
      _nw248[(new BigNumber(4)).toNumber()] = _1502_v3;
      _nw248[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1502_v3, _1502_v3);
      _1503_v4 = _nw248;
      let _index254 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1503_v4).length));
      (_1503_v4)[_index254] = _1502_v3;
      let _arr49 = _this.f19;
      let _index255 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
      _arr49[_index255] = p1;
      let _1504_v5;
      _1504_v5 = _dafny.MultiSet.fromElements(!(p1));
      let _1505_v7;
      _1505_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,(_this).f18);
      let _1506_v8;
      _1506_v8 = _dafny.Seq.of(p2, _this.f22, p0);
      let _index256 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1503_v4).length));
      let _arr50 = _this.f19;
      let _index257 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
      let _rhs244 = !((((_this).f18) ? (p1) : (!((_this).f18))));
      let _rhs245 = !(((((_this).f18) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f18))) : (_1504_v5))).IsDisjointFrom(_1504_v5));
      let _rhs246 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-140)), ((_1507_p1) => function (_1508_i0) {
        return _dafny.Seq.of(_1507_p1, (_this).f18, _1507_p1, !(_1507_p1), (_this).f18);
      })(p1));
      let _rhs247 = (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll58 = new _dafny.Map();
        for (const _compr_58 of (_1505_v7).Keys.Elements) {
          let _1509_v6 = _compr_58;
          if ((_1505_v7).contains(_1509_v6)) {
            _coll58.push([(_1509_v6).multipliedBy((_dafny.ZERO).minus(_this.f22)),_this.f22]);
          }
        }
        return _coll58;
      }(),new BigNumber((_1506_v8).length)), globalState);
      let _lhs201 = _1503_v4;
      let _lhs202 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_1503_v4).length));
      let _lhs203 = _this.f19;
      let _lhs204 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
      r0 = _rhs244;
      r0 = _rhs245;
      _lhs201[_lhs202] = _rhs246;
      _lhs203[_lhs204] = _rhs247;
      let _1510_i1;
      _1510_i1 = _dafny.ZERO;
      L16: {
        while ((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1510_i1)) {
              break L16;
            }
            _1510_i1 = (_1510_i1).plus(_dafny.ONE);
            let _1511_v9;
            let _init45 = ((_1512_v0) => function (_1513_i2) {
              return _1512_v0;
            })(_1499_v0);
            let _nw249 = Array((new BigNumber(10)).toNumber());
            for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw249.length); _i0_45++) {
              _nw249[_i0_45] = _init45(new BigNumber(_i0_45));
            }
            _1511_v9 = _nw249;
            let _index258 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length));
            (_1511_v9)[_index258] = _1499_v0;
            let _1514_v10;
            _1514_v10 = _dafny.Seq.UnicodeFromString("bmfhaq");
            let _1515_v11;
            _1515_v11 = _dafny.MultiSet.fromElements(_1499_v0);
            let _1516_v12;
            _1516_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1515_v11).cardinality()),p2);
            let _1517_v13;
            _1517_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1516_v12,p0);
            let _index259 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length));
            let _rhs248 = new BigNumber((_dafny.Seq.Concat(_1514_v10, _1514_v10)).length);
            let _rhs249 = (_this).fm1(_1517_v13, globalState);
            let _rhs250 = _1499_v0;
            let _rhs251 = new BigNumber(795);
            let _lhs205 = _this;
            let _lhs206 = _1511_v9;
            let _lhs207 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length));
            _lhs205.f22 = _rhs248;
            r0 = _rhs249;
            _lhs206[_lhs207] = _rhs250;
            r2 = _rhs251;
            let _1518_v14;
            _1518_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))],true);
            if ((_1518_v14).contains(!((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]))) {
              (globalState).f3 = p1;
              let _arr51 = _this.f19;
              let _index260 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
              let _rhs252 = _this.f22;
              let _rhs253 = !_dafny.Seq.contains(_1514_v10, (_1511_v9)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length))]);
              let _rhs254 = (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))];
              let _rhs255 = (_this).fm0((_1506_v8)[_module.__default.safeIndex((_this).f20, new BigNumber((_1506_v8).length))], _1499_v0, globalState);
              let _rhs256 = (_1516_v12).contains((_this).f20);
              let _lhs208 = globalState;
              let _lhs209 = globalState;
              let _lhs210 = _this.f19;
              let _lhs211 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
              r2 = _rhs252;
              _lhs208.f3 = _rhs253;
              r1 = _rhs254;
              _lhs209.f14 = _rhs255;
              _lhs210[_lhs211] = _rhs256;
              _1514_v10 = _dafny.Seq.UnicodeFromString("jdrmxxf");
              r0 = (!(p1)) || (((p1) ? ((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]) : ((_1501_v2)[_module.__default.safeIndex((_this).f20, new BigNumber((_1501_v2).length))])));
              let _arr52 = _this.f19;
              let _index261 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
              _arr52[_index261] = !_dafny.Seq.contains(_dafny.Seq.Concat(_1514_v10, _1514_v10), _1499_v0);
            } else {
              let _1519_v15;
              let _init46 = function (_1520_i3) {
                return (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))];
              };
              let _nw250 = Array((new BigNumber(15)).toNumber());
              for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw250.length); _i0_46++) {
                _nw250[_i0_46] = _init46(new BigNumber(_i0_46));
              }
              _1519_v15 = _nw250;
              let _1521_v16;
              let _nw251 = new _module.C6();
              _nw251.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))],(_this).f20), (_this).f20, (_this).f20, _this.f21, _1519_v15, _this.f17, (_this).f18);
              _1521_v16 = _nw251;
              let _1522_v17;
              _1522_v17 = _dafny.Set.fromElements(_1521_v16, _1521_v16);
              (globalState).f14 = (((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]) ? (new BigNumber((_1514_v10).length)) : (((_this).f20).multipliedBy(new BigNumber((_1522_v17).length))));
              (globalState).f14 = _this.f22;
              let _1523_v18;
              let _nw252 = new _module.C3();
              _nw252.__ctor(_1519_v15, (_dafny.Set.fromElements(_1519_v15)).Union(_this.f17), !(_1505_v7).contains((_1521_v16).f20));
              _1523_v18 = _nw252;
              let _index262 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length));
              (_1511_v9)[_index262] = (_1511_v9)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length))];
              let _arr53 = _this.f19;
              let _index263 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length));
              _arr53[_index263] = !((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]);
            }
            if (p1) {
              let _1524_v19;
              let _nw253 = Array((new BigNumber(28)).toNumber()).fill(false);
              _1524_v19 = _nw253;
              let _1525_v20;
              _1525_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v19,_1524_v19);
              _1525_v20 = ((_1525_v20).Merge(_1525_v20)).Merge(_1525_v20);
              let _1526_v21;
              _1526_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(_dafny.Seq.of((_this).f20, new BigNumber(231)), globalState),p2);
              _1526_v21 = (_1526_v21).update(_1505_v7, p0);
              let _1527_v22;
              let _nw254 = new _module.C2();
              _nw254.__ctor(_this.f17, _module.__default.fm13(globalState));
              _1527_v22 = _nw254;
              let _1528_v23;
              _1528_v23 = _module.D11.create_DC27(_dafny.Seq.update(_1501_v2, _module.__default.safeIndex((_this).f20, new BigNumber((_1501_v2).length)), (_this).f18), p0, p2);
              let _1529_v24;
              _1529_v24 = _module.D0.create_DC1(_dafny.Seq.Concat(_1514_v10, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("hcpfqafh"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("hcpfqafh")).length)), (_1511_v9)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1511_v9).length))])), _module.__default.safeDivisionInt((_dafny.ZERO).minus((_1528_v23).dtor_cf51), p0), p0, _1499_v0);
              _1529_v24 = _1529_v24;
              (globalState).f9 = p1;
            } else {
              let _1530_v25;
              let _nw255 = Array((new BigNumber(11)).toNumber());
              _nw255[(_dafny.ZERO).toNumber()] = (_this).f18;
              _nw255[(_dafny.ONE).toNumber()] = (_this).f18;
              _nw255[(new BigNumber(2)).toNumber()] = (_this).f18;
              _nw255[(new BigNumber(3)).toNumber()] = _module.__default.fm13(globalState);
              _nw255[(new BigNumber(4)).toNumber()] = !((_this).fm1(_1517_v13, globalState));
              _nw255[(new BigNumber(5)).toNumber()] = (_this).f18;
              _nw255[(new BigNumber(6)).toNumber()] = false;
              _nw255[(new BigNumber(7)).toNumber()] = (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))];
              _nw255[(new BigNumber(8)).toNumber()] = (_this).f18;
              _nw255[(new BigNumber(9)).toNumber()] = p1;
              _nw255[(new BigNumber(10)).toNumber()] = (_this).f18;
              _1530_v25 = _nw255;
              let _1531_v27;
              let _nw256 = new _module.C8();
              _nw256.__ctor(p0, !((p1) === ((((_1505_v7).contains(_this.f22)) ? ((_1505_v7).get(_this.f22)) : ((_this).f18)))), _1530_v25, _this.f17, (_this).fm1(_1517_v13, globalState), p2, (function () {
                let _coll59 = new _dafny.Map();
                for (const _compr_59 of (_1515_v11).Elements) {
                  let _1532_v26 = _compr_59;
                  if ((_1515_v11).contains(_1532_v26)) {
                    _coll59.push([_1532_v26,p1]);
                  }
                }
                return _coll59;
              }()).Merge(_this.f21));
              _1531_v27 = _nw256;
              _1531_v27 = _1531_v27;
              _1530_v25 = _1530_v25;
              let _1533_v28;
              let _nw257 = new _module.C0();
              _nw257.__ctor();
              _1533_v28 = _nw257;
              _1533_v28 = _1533_v28;
              let _1534_v29;
              _1534_v29 = _dafny.Set.fromElements((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]);
              let _1535_v30;
              _1535_v30 = _1501_v2;
              let _1536_v31;
              _1536_v31 = _dafny.MultiSet.fromElements((_this).f20);
              let _1537_v32;
              let _nw258 = Array((new BigNumber(17)).toNumber());
              _nw258[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1501_v2, _dafny.Seq.update(_1501_v2, _module.__default.safeIndex(p0, new BigNumber((_1501_v2).length)), false));
              _nw258[(_dafny.ONE).toNumber()] = (((_1531_v27).f24) ? (_1501_v2) : (_1501_v2));
              _nw258[(new BigNumber(2)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1501_v2, _module.__default.safeIndex(new BigNumber(506), new BigNumber((_1501_v2).length)), (_1531_v27).f24);
              _nw258[(new BigNumber(4)).toNumber()] = (_1535_v30);
              _nw258[(new BigNumber(5)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(6)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_1501_v2, _module.__default.safeIndex(p0, new BigNumber((_1501_v2).length)), false);
              _nw258[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1501_v2, _module.__default.fm16(globalState));
              _nw258[(new BigNumber(9)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(10)).toNumber()] = _module.__default.fm16(globalState);
              _nw258[(new BigNumber(11)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f18, (_this).f18, (_this).f18, p1, p1), _1501_v2);
              _nw258[(new BigNumber(13)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(14)).toNumber()] = _module.__default.fm16(globalState);
              _nw258[(new BigNumber(15)).toNumber()] = _1501_v2;
              _nw258[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1501_v2, _module.__default.safeIndex(new BigNumber((_1536_v31).cardinality()), new BigNumber((_1501_v2).length)), (_this).f18), _1501_v2);
              _1537_v32 = _nw258;
              let _index264 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1537_v32).length));
              (_1537_v32)[_index264] = _1501_v2;
              let _1538_v33;
              _1538_v33 = _dafny.Seq.of(_module.__default.fm37((_this).f18, _this.f22, (_this).f18, (_1531_v27).f24, globalState));
              let _1539_v34;
              _1539_v34 = _dafny.MultiSet.fromElements(((_1538_v33)[_module.__default.safeIndex(new BigNumber((_1518_v14).length), new BigNumber((_1538_v33).length))]).Intersect(_1534_v29));
              let _index265 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1537_v32).length));
              let _rhs257 = _dafny.Set.fromElements((new BigNumber((_1514_v10).length)).isEqualTo(p2));
              let _rhs258 = (new BigNumber((_1506_v8).length)).multipliedBy(new BigNumber((_module.__default.fm40(p1, (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))], globalState)).length));
              let _rhs259 = _dafny.Seq.Concat(_1501_v2, _1501_v2);
              let _rhs260 = _1539_v34;
              let _lhs212 = _this;
              let _lhs213 = _1537_v32;
              let _lhs214 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1537_v32).length));
              _1534_v29 = _rhs257;
              _lhs212.f22 = _rhs258;
              _lhs213[_lhs214] = _rhs259;
              _1539_v34 = _rhs260;
              let _1540_v35;
              _1540_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v25,(_this).f20);
              let _rhs261 = _1540_v35;
              let _rhs262 = _this.f22;
              let _rhs263 = true;
              let _rhs264 = (_this).f20;
              let _rhs265 = (_dafny.ZERO).minus((_this).f20);
              let _lhs215 = globalState;
              let _lhs216 = globalState;
              let _lhs217 = globalState;
              _1540_v35 = _rhs261;
              _lhs215.f14 = _rhs262;
              _lhs216.f9 = _rhs263;
              _lhs217.f14 = _rhs264;
              r2 = _rhs265;
            }
            r1 = !((_this).fm1((_dafny.Map.Empty.slice().updateUnsafe(_1516_v12,p2)).Merge(_1517_v13), globalState));
          }
        }
      }
      let _1541_v36;
      _1541_v36 = _dafny.Seq.of(_1499_v0);
      let _1542_v37;
      let _nw259 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1542_v37 = _nw259;
      let _1543_v38;
      _1543_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v0,_1542_v37);
      let _1544_v39;
      let _nw260 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1544_v39 = _nw260;
      let _1545_v40;
      _1545_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1546_v41;
      _1546_v41 = _dafny.Set.fromElements(_1505_v7, _1505_v7);
      let _1547_v42;
      _1547_v42 = _module.D1.create_DC4((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))], _1546_v41, new BigNumber(-445));
      let _1548_v43;
      _1548_v43 = _dafny.Seq.of(_1547_v42);
      let _1549_v44;
      _1549_v44 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20, (((_1545_v40).contains(_this.f22)) ? ((_1545_v40).get(_this.f22)) : (p2)), new BigNumber((_1548_v43).length), p2);
      let _1550_v45;
      _1550_v45 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D14.create_DC34(_1544_v39)).dtor_cf64,_1499_v0)).length), (((_1549_v44).contains(new BigNumber(346))) ? ((_1549_v44).get(new BigNumber(346))) : ((_this).f20)));
      let _1551_v46;
      _1551_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v0,new BigNumber((_1550_v45).length));
      let _1552_v47;
      _1552_v47 = _dafny.Seq.of(_1551_v46);
      let _1553_v48;
      _1553_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1552_v47,_dafny.Seq.UnicodeFromString("ymwc"));
      let _rhs266 = _1541_v36;
      let _rhs267 = _1543_v38;
      let _rhs268 = p0;
      let _rhs269 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1541_v36, _module.__default.safeIndex(p0, new BigNumber((_1541_v36).length)), _1499_v0), _dafny.Seq.Concat((((_1553_v48).contains(_1552_v47)) ? ((_1553_v48).get(_1552_v47)) : (_1541_v36)), _1541_v36)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1501_v2, _dafny.Seq.update(_1501_v2, _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1501_v2).length)), (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]))).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1541_v36, _module.__default.safeIndex(p0, new BigNumber((_1541_v36).length)), _1499_v0), _dafny.Seq.Concat((((_1553_v48).contains(_1552_v47)) ? ((_1553_v48).get(_1552_v47)) : (_1541_v36)), _1541_v36))).length)), _1499_v0), _module.__default.safeIndex(new BigNumber((_1541_v36).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1541_v36, _module.__default.safeIndex(p0, new BigNumber((_1541_v36).length)), _1499_v0), _dafny.Seq.Concat((((_1553_v48).contains(_1552_v47)) ? ((_1553_v48).get(_1552_v47)) : (_1541_v36)), _1541_v36)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1501_v2, _dafny.Seq.update(_1501_v2, _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1501_v2).length)), (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]))).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1541_v36, _module.__default.safeIndex(p0, new BigNumber((_1541_v36).length)), _1499_v0), _dafny.Seq.Concat((((_1553_v48).contains(_1552_v47)) ? ((_1553_v48).get(_1552_v47)) : (_1541_v36)), _1541_v36))).length)), _1499_v0)).length)), _1499_v0);
      let _rhs270 = p0;
      let _lhs218 = globalState;
      _1541_v36 = _rhs266;
      _1543_v38 = _rhs267;
      _lhs218.f14 = _rhs268;
      _1541_v36 = _rhs269;
      r2 = _rhs270;
      let _1554_i4;
      _1554_i4 = _dafny.ZERO;
      L17: {
        while (!(p1) || ((_this).f18)) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1554_i4)) {
              break L17;
            }
            _1554_i4 = (_1554_i4).plus(_dafny.ONE);
            (_this).f22 = p2;
            _1551_v46 = (_1551_v46).update(_1499_v0, (_this).f20);
            let _1555_v49;
            _1555_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1501_v2,_1499_v0);
            let _1556_v50;
            let _nw261 = Array((new BigNumber(3)).toNumber());
            _nw261[(_dafny.ZERO).toNumber()] = _1555_v49;
            _nw261[(_dafny.ONE).toNumber()] = _1555_v49;
            _nw261[(new BigNumber(2)).toNumber()] = _1555_v49;
            _1556_v50 = _nw261;
            let _index266 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1556_v50).length));
            (_1556_v50)[_index266] = _1555_v49;
            let _index267 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1556_v50).length));
            let _rhs271 = (_this).f18;
            let _rhs272 = (_1550_v45).Difference(_1550_v45);
            let _rhs273 = _dafny.Map.Empty.slice().updateUnsafe(_1501_v2,_1499_v0);
            let _rhs274 = (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))];
            let _lhs219 = globalState;
            let _lhs220 = _1556_v50;
            let _lhs221 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1556_v50).length));
            let _lhs222 = globalState;
            _lhs219.f9 = _rhs271;
            _1550_v45 = _rhs272;
            _lhs220[_lhs221] = _rhs273;
            _lhs222.f3 = _rhs274;
            (globalState).f3 = (p1) && (p1);
          }
        }
      }
      let _1557_i5;
      _1557_i5 = _dafny.ZERO;
      L18: {
        while ((_this).f18) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1557_i5)) {
              break L18;
            }
            _1557_i5 = (_1557_i5).plus(_dafny.ONE);
            r0 = !((_this).f18);
            if (false) {
              let _1558_v51;
              _1558_v51 = _dafny.Map.Empty.slice().updateUnsafe(!((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]) || ((_this).f18),!((_1547_v42).dtor_cf9));
              let _1559_v52;
              let _nw262 = new _module.C2();
              _nw262.__ctor(_this.f17, p1);
              _1559_v52 = _nw262;
              let _1560_v53;
              _1560_v53 = _dafny.MultiSet.fromElements(_1559_v52);
              _1558_v51 = (_1558_v51).update((_1547_v42).dtor_cf9, (_dafny.MultiSet.fromElements(_1559_v52)).IsDisjointFrom(_1560_v53));
              _1499_v0 = _1499_v0;
              _1501_v2 = _dafny.Seq.Concat(_dafny.Seq.update(_1501_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1541_v36, _module.__default.safeIndex((_this).f20, new BigNumber((_1541_v36).length)), _1499_v0)).length), new BigNumber((_1501_v2).length)), (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]), _dafny.Seq.update(_dafny.Seq.of((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]), _module.__default.safeIndex(_this.f22, new BigNumber((_dafny.Seq.of((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))])).length)), (_1559_v52).f18));
              let _1561_v54;
              let _nw263 = new _module.C4();
              _nw263.__ctor(new BigNumber(((_1505_v7).Merge(_1505_v7)).length), p0, _this.f17, (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(_1545_v40,p0), globalState));
              _1561_v54 = _nw263;
              (_1561_v54).f26 = new BigNumber(((((new BigNumber((_1541_v36).length)).isLessThanOrEqualTo((_this).fm2((_1559_v52).f18, new BigNumber((_1541_v36).length), _module.__default.fm43(globalState), globalState))) ? (_dafny.Seq.Concat(_1501_v2, _module.__default.fm16(globalState))) : (_1501_v2))).length);
            } else {
              let _1562_v55;
              _1562_v55 = _module.D0.create_DC0((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]);
              _1499_v0 = _module.__default.fm33(_module.__default.fm7(p2, _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f22), function (_pat_let28_0) {
                return function (_1563_dt__update__tmp_h0) {
                  return function (_pat_let29_0) {
                    return function (_1564_dt__update_hcf0_h0) {
                      return _module.D0.create_DC0(_1564_dt__update_hcf0_h0);
                    }(_pat_let29_0);
                  }((_this).f18);
                }(_pat_let28_0);
              }(_1562_v55), globalState), (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))], globalState);
              let _1565_v56;
              let _nw264 = Array((new BigNumber(17)).toNumber()).fill(false);
              _1565_v56 = _nw264;
              let _1566_v57;
              let _nw265 = new _module.C7();
              _nw265.__ctor(_this.f22, _this.f21, _1565_v56, _this.f17, p1);
              _1566_v57 = _nw265;
              let _1567_v58;
              _1567_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1566_v57,(_this).f20);
              _1567_v58 = (_1567_v58).update(_1566_v57, new BigNumber((_1541_v36).length));
              (globalState).f3 = p1;
              let _1568_v59;
              let _nw266 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
              _1568_v59 = _nw266;
              let _1569_v60;
              _1569_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v0,_1565_v56);
              let _1570_v61;
              let _nw267 = new _module.C5();
              _nw267.__ctor(_1541_v36, (_this).f18, p0, _this.f21, (((_1569_v60).contains(_1499_v0)) ? ((_1569_v60).get(_1499_v0)) : (_1565_v56)), _dafny.Set.fromElements(_1565_v56), false);
              _1570_v61 = _nw267;
              let _1571_v62;
              _1571_v62 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))]);
              let _index268 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1568_v59).length));
              (_1568_v59)[_index268] = (_dafny.Map.Empty.slice().updateUnsafe(_1570_v61,_1571_v62)).update(_1570_v61, _1571_v62);
              let _1572_v63;
              _1572_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v5,_1570_v61);
              let _1573_v64;
              _1573_v64 = _dafny.Map.Empty.slice().updateUnsafe((((_1572_v63).contains(_1504_v5)) ? ((_1572_v63).get(_1504_v5)) : (_1570_v61)),_1571_v62);
              let _index269 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1568_v59).length));
              (_1568_v59)[_index269] = _1573_v64;
              let _1574_v67;
              _1574_v67 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll60 = new _dafny.Map();
                for (const _compr_60 of (_1506_v8).Elements) {
                  let _1575_v66 = _compr_60;
                  if (_dafny.Seq.contains(_1506_v8, _1575_v66)) {
                    _coll60.push([(_1575_v66).plus(p2),(_dafny.ZERO).minus((_this).f20)]);
                  }
                }
                return _coll60;
              }(),(_this).f20);
              (globalState).f3 = (_this).fm1(function () {
                let _coll61 = new _dafny.Map();
                for (const _compr_61 of (_1574_v67).Keys.Elements) {
                  let _1576_v65 = _compr_61;
                  if ((_1574_v67).contains(_1576_v65)) {
                    _coll61.push([_1576_v65,p0]);
                  }
                }
                return _coll61;
              }(), globalState);
            }
            let _1577_v68;
            let _init47 = function (_1578_i6) {
              return (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))];
            };
            let _nw268 = Array((new BigNumber(8)).toNumber());
            for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw268.length); _i0_47++) {
              _nw268[_i0_47] = _init47(new BigNumber(_i0_47));
            }
            _1577_v68 = _nw268;
            let _1579_v70;
            let _nw269 = new _module.C8();
            _nw269.__ctor(p0, (_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))], _1577_v68, _this.f17, (_this).f18, p2, function () {
              let _coll62 = new _dafny.Map();
              for (const _compr_62 of (_1551_v46).Keys.Elements) {
                let _1580_v69 = _compr_62;
                if ((_1551_v46).contains(_1580_v69)) {
                  _coll62.push([_1580_v69,p1]);
                }
              }
              return _coll62;
            }());
            _1579_v70 = _nw269;
            let _1581_v71;
            let _nw270 = Array((new BigNumber(10)).toNumber());
            _nw270[(_dafny.ZERO).toNumber()] = _1579_v70;
            _nw270[(_dafny.ONE).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(2)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(3)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(4)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(5)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(6)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(7)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(8)).toNumber()] = _1579_v70;
            _nw270[(new BigNumber(9)).toNumber()] = _1579_v70;
            _1581_v71 = _nw270;
            let _1582_v72;
            _1582_v72 = _dafny.Set.fromElements(_1581_v71);
            _1582_v72 = _1582_v72;
            (_this).f22 = p2;
          }
        }
      }
      let _1583_v73;
      _1583_v73 = _dafny.Set.fromElements((_this.f19)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_this.f19).length))], (_this).f18);
      r0 = (new BigNumber((_1583_v73).length)).isLessThan(new BigNumber((_1541_v36).length));
      r1 = !((((!((_this).f18)) ? (_dafny.MultiSet.fromElements(p0)) : (_1549_v44))).IsDisjointFrom(_1549_v44));
      r2 = (p2).multipliedBy(p0);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
