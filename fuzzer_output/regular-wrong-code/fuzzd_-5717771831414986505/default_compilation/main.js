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
    static fm5(globalState) {
      if (!_dafny.areEqual(_dafny.Seq.of(true), _dafny.Seq.of(false))) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(988)), function (_0_i0) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(811)), function (_1_i1) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          });
        });
      } else {
        return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("tuqlhd"), _dafny.Seq.UnicodeFromString("wcvbklxx"));
      }
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(297)), function (_2_i0) {
        return _module.__default.safeDivisionInt(new BigNumber(411), new BigNumber(749));
      });
    };
    static fm9(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),!(true))));
    };
    static fm11(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_3_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      });
    };
    static fm12(globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-509))).length), new BigNumber(214)));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gyflnn"),true)).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), function (_4_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }),!(true))).Keys.Elements) {
          let _5_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), function (_4_i0) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }),!(true))).contains(_5_v0)) {
            _coll0.push([_5_v0,false]);
          }
        }
        return _coll0;
      }());
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(false, !(true)))).Intersect(_dafny.MultiSet.fromElements(true, true, true));
    };
    static fm15(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(193), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0)))).length), new BigNumber(546))).Difference(_dafny.MultiSet.fromElements(new BigNumber(545), new BigNumber(965)))).Difference(_dafny.MultiSet.fromElements(new BigNumber(808)));
    };
    static fm16(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('k'.codePointAt(0));
    };
    static fm17(globalState) {
      return _dafny.Seq.UnicodeFromString("qrabdg");
    };
    static fm18(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), function (_6_i0) {
        return new BigNumber((_dafny.Seq.of(!(true), !(true))).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(-533), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)));
    };
    static fm19(p0, globalState) {
      return (((true) ? ((_module.D7.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).length)))).dtor_cf28) : (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Seq.UnicodeFromString("n")).length)));
    };
    static fm20(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, !(false)),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)));
    };
    static fm21(p0, globalState) {
      let _source0 = _module.D5.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), function (_7_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}));
      if (_source0.is_DC16) {
        let _8___mcc_h0 = (_source0).cf19;
        let _9___mcc_h1 = (_source0).cf20;
        let _10___mcc_h2 = (_source0).cf21;
        let _11___mcc_h3 = (_source0).cf22;
        let _12___mcc_h4 = (_source0).cf23;
        let _13_cf23 = _12___mcc_h4;
        let _14_cf22 = _11___mcc_h3;
        let _15_cf21 = _10___mcc_h2;
        let _16_cf20 = _9___mcc_h1;
        let _17_cf19 = _8___mcc_h0;
        return _module.D3.create_DC10(true, _16_cf20);
      } else {
        let _18___mcc_h5 = (_source0).cf18;
        let _19_cf18 = _18___mcc_h5;
        return _module.D3.create_DC10(true, new BigNumber(713));
      }
    };
    static fm22(globalState) {
      if (((true) ? (false) : (false))) {
        return _module.D1.create_DC3(_module.D0.create_DC1(), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false))).length));
      } else {
        return _module.D1.create_DC3(_module.D0.create_DC1(), new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.Seq.of(new BigNumber(470), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(334)), function (_20_i0) {
    return new _dafny.CodePoint('n'.codePointAt(0));
  })).length), new BigNumber(178))).Elements) {
    let _21_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(470), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(334)), function (_20_i0) {
      return new _dafny.CodePoint('n'.codePointAt(0));
    })).length), new BigNumber(178)), _21_v0)) {
      _coll1.push([(_21_v0).plus(new BigNumber(493)),new BigNumber(650)]);
    }
  }
  return _coll1;
}()).length));
      }
    };
    static fm23(p0, globalState) {
      return _dafny.Seq.of(false);
    };
    static fm24(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(670), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(900), new BigNumber(919), new BigNumber(804), new BigNumber(901))).cardinality()), new BigNumber(-743))).length)))).Difference((_dafny.Set.fromElements(new BigNumber(-722), new BigNumber(152), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber(-562))).length), new BigNumber(865), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), function (_22_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length), new BigNumber(132))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(800),new BigNumber((_dafny.Seq.of(new BigNumber(331))).length)),new BigNumber(-891))).length))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).cardinality()), new BigNumber(-476), new BigNumber(959), new BigNumber((_dafny.Seq.of(false, true)).length))));
    };
    static fm25(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(new BigNumber((_dafny.Seq.of(!(true), false, false)).length));
    };
    static fm26(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(212), new BigNumber(741))).length),new BigNumber(393)),true)).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-93), new BigNumber(105))) {
            let _23_v1 = _compr_3;
            if (((new BigNumber(-93)).isLessThanOrEqualTo(_23_v1)) && ((_23_v1).isLessThan(new BigNumber(105)))) {
              _coll3.push([(_23_v1).minus(new BigNumber(511)),new BigNumber(96)]);
            }
          }
          return _coll3;
        }())).Elements) {
          let _24_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-93), new BigNumber(105))) {
              let _23_v1 = _compr_4;
              if (((new BigNumber(-93)).isLessThanOrEqualTo(_23_v1)) && ((_23_v1).isLessThan(new BigNumber(105)))) {
                _coll4.push([(_23_v1).minus(new BigNumber(511)),new BigNumber(96)]);
              }
            }
            return _coll4;
          }()), _24_v0)) {
            _coll2.push([_24_v0,false]);
          }
        }
        return _coll2;
      }());
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(790),new BigNumber((((_module.D2.create_DC4(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-544),!(true))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-406)),!(true))).length))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("sqqfutbxm")).length))).cardinality())))).dtor_cf4).Difference(_dafny.Set.fromElements(new BigNumber(854)))).length));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      if (false) {
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-750),new BigNumber((_dafny.Set.fromElements(new BigNumber(-201), new BigNumber(521))).length)), function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-999), new BigNumber(240))) {
            let _25_v0 = _compr_5;
            if (((new BigNumber(-999)).isLessThanOrEqualTo(_25_v0)) && ((_25_v0).isLessThan(new BigNumber(240)))) {
              _coll5.push([(_25_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-219))).length)),new BigNumber(-167)]);
            }
          }
          return _coll5;
        }())).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(245),new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-699),new BigNumber((_dafny.Set.fromElements(true)).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(999),new BigNumber(629)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(52))).length),true)).length)), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gsppdqr")).length)),new BigNumber(272))));
      } else if (true) {
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(553), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), function (_26_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })).length))).length))).length),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D18.create_DC49(_dafny.Seq.of(true), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).cardinality()))))).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(-594))).length),new BigNumber(913)));
      } else {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(78),new BigNumber(765)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(665),new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-553)), function (_27_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(414), new BigNumber(809))) {
              let _28_v1 = _compr_6;
              if (((new BigNumber(414)).isLessThanOrEqualTo(_28_v1)) && ((_28_v1).isLessThan(new BigNumber(809)))) {
                _coll6.add((_28_v1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(232)), function (_29_i2) {
                  return new _dafny.CodePoint('y'.codePointAt(0));
                })).length)));
              }
            }
            return _coll6;
          }()).length),new BigNumber(906));
        })).length)))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(458),new BigNumber(240))));
      }
    };
    static fm29(p0, p1, p2, globalState) {
      return _module.D7.create_DC21((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).cardinality()))).length)),new BigNumber((_dafny.Seq.of(false)).length))).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qyyg")).length)),new BigNumber(29))).length)), new BigNumber((_dafny.Seq.of(false)).length));
    };
    static fm30(globalState) {
      return _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe(true,false));
    };
    static fm31(globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(700),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-206),true), function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(137), new BigNumber(-918))) {
          let _30_v0 = _compr_7;
          if (((new BigNumber(137)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(-918)))) {
            _coll7.push([(_30_v0).multipliedBy(new BigNumber(-89)),false]);
          }
        }
        return _coll7;
      }(), function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(432),new BigNumber(305))).Keys.Elements) {
          let _31_v1 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(432),new BigNumber(305))).contains(_31_v1)) {
            _coll8.push([_module.__default.safeModuloInt(_31_v1, new BigNumber(45)),true]);
          }
        }
        return _coll8;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(836),!(false)))).Intersect(_dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),!(true))), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-814))).length)),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("tlws")).length))).length),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length))).cardinality()),true), function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(561), new BigNumber(572))) {
          let _32_v2 = _compr_9;
          if (((new BigNumber(561)).isLessThanOrEqualTo(_32_v2)) && ((_32_v2).isLessThan(new BigNumber(572)))) {
            _coll9.push([_module.__default.safeDivisionInt(_32_v2, new BigNumber((_dafny.Seq.UnicodeFromString("gswko")).length)),false]);
          }
        }
        return _coll9;
      }()));
    };
    static fm32(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(false, false)).Difference(_dafny.Set.fromElements(true));
    };
    static fm33(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(((!(true)) ? (new _dafny.CodePoint('h'.codePointAt(0))) : (new _dafny.CodePoint('t'.codePointAt(0)))));
    };
    static fm34(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), function (_33_i0) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)));
      });
    };
    static fm35(p0, p1, globalState) {
      return !(function (_source1) {
        if (_source1.is_DC5) {
          let _34___mcc_h0 = (_source1).cf5;
          let _35_cf5 = _34___mcc_h0;
          return (_dafny.Set.fromElements(_35_cf5, _35_cf5, _35_cf5, _35_cf5)).IsProperSubsetOf(_dafny.Set.fromElements(_35_cf5));
        } else if (_source1.is_DC6) {
          return (false) === (!(!(false)));
        } else if (_source1.is_DC7) {
          let _36___mcc_h1 = (_source1).cf6;
          let _37___mcc_h2 = (_source1).cf7;
          let _38___mcc_h3 = (_source1).cf8;
          let _39_cf8 = _38___mcc_h3;
          let _40_cf7 = _37___mcc_h2;
          let _41_cf6 = _36___mcc_h1;
          return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("fucxg"), _dafny.Seq.UnicodeFromString("gwvrtg"));
        } else if (_source1.is_DC4) {
          let _42___mcc_h4 = (_source1).cf4;
          let _43_cf4 = _42___mcc_h4;
          return !(false);
        } else {
          let _44___mcc_h5 = (_source1).cf9;
          let _45_cf9 = _44___mcc_h5;
          return true;
        }
      }(_module.D2.create_DC8(_module.D2.create_DC8(_module.D2.create_DC4(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).length))).length), new BigNumber(870), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-713),false)).length)))))));
    };
    static fm36(p0, p1, p2, globalState) {
      return _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false)),new BigNumber(379)));
    };
    static fm37(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),true)).Merge(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_46_i0) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })).Elements) {
          let _47_v0 = _compr_10;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_46_i0) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }), _47_v0)) {
            _coll10.push([_47_v0,!(false)]);
          }
        }
        return _coll10;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),false));
    };
    static Main(__noArgsParameter) {
      let _48_v0;
      _48_v0 = _dafny.Seq.UnicodeFromString("rwdkj");
      let _49_v1;
      _49_v1 = new BigNumber(143);
      let _50_v2;
      _50_v2 = new _dafny.CodePoint('w'.codePointAt(0));
      let _51_v3;
      let _init0 = function (_52_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      };
      let _nw0 = Array((new BigNumber(8)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _51_v3 = _nw0;
      let _53_globalState;
      let _nw1 = new _module.GlobalState();
      _nw1.__ctor(new BigNumber(287), _dafny.Seq.Concat(_48_v0, _dafny.Seq.update(_48_v0, _module.__default.safeIndex((_dafny.ZERO).minus(_49_v1), new BigNumber((_48_v0).length)), _50_v2)), _51_v3);
      _53_globalState = _nw1;
      let _54_v4;
      _54_v4 = true;
      let _55_v5;
      _55_v5 = _dafny.Map.Empty.slice().updateUnsafe(_50_v2,_54_v4);
      let _56_v6;
      let _nw2 = Array((new BigNumber(16)).toNumber()).fill(false);
      _56_v6 = _nw2;
      let _57_v7;
      _57_v7 = _dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), _50_v2);
      let _58_v8;
      _58_v8 = _dafny.Seq.of(_48_v0, _dafny.Seq.update(_48_v0, _module.__default.safeIndex(new BigNumber((_57_v7).length), new BigNumber((_48_v0).length)), _50_v2));
      let _59_v9;
      _59_v9 = _dafny.MultiSet.fromElements(false);
      let _60_v10;
      let _nw3 = new _module.C5();
      _nw3.__ctor(_49_v1, (((_55_v5).contains(_50_v2)) ? ((_55_v5).get(_50_v2)) : (_54_v4)), _56_v6, ((_54_v4) ? (_54_v4) : (_54_v4)), ((_module.__default.fm35(_58_v8, _54_v4, _53_globalState)) ? (_54_v4) : (false)), _59_v9);
      _60_v10 = _nw3;
      _54_v4 = (_60_v10).f10;
      let _61_i1;
      _61_i1 = _dafny.ZERO;
      L0: {
        while ((_60_v10).f10) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_61_i1)) {
              break L0;
            }
            _61_i1 = (_61_i1).plus(_dafny.ONE);
            let _index0 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            (_56_v6)[_index0] = (_60_v10).f10;
            let _index1 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            (_56_v6)[_index1] = _54_v4;
            if ((_60_v10).f10) {
              let _62_v11;
              _62_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cfhadr"),(_60_v10).f10);
              let _index2 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
              let _rhs0 = ((_60_v10).fm4(_49_v1, new BigNumber(364), _62_v11, (_60_v10).f9, _53_globalState)).minus(new BigNumber(589));
              let _rhs1 = (((_60_v10).f10) ? (!(!(!((_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))])))) : (_54_v4));
              let _lhs0 = _53_globalState;
              let _lhs1 = _56_v6;
              let _lhs2 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
              _lhs0.f0 = _rhs0;
              _lhs1[_lhs2] = _rhs1;
              let _63_v12;
              _63_v12 = _dafny.Map.Empty.slice().updateUnsafe((_60_v10).f9,(_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))]);
              (_60_v10).m1((_60_v10).f9, (_60_v10).f10, (((_63_v12).contains((_60_v10).f9)) ? ((_63_v12).get((_60_v10).f9)) : ((_60_v10).f10)), _53_globalState);
              let _64_v13;
              _64_v13 = _dafny.Seq.of(_49_v1, (_dafny.ZERO).minus((_60_v10).f9));
              let _65_v14;
              _65_v14 = _dafny.Seq.of((_64_v13)[_module.__default.safeIndex(new BigNumber(91), new BigNumber((_64_v13).length))], new BigNumber((_48_v0).length));
              (_53_globalState).f0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_64_v13).length), (_60_v10).f9));
              (_53_globalState).f0 = ((true) ? ((_49_v1).minus((_60_v10).f9)) : (_49_v1));
              _54_v4 = _dafny.Seq.contains(_48_v0, _50_v2);
            } else {
              (_53_globalState).f0 = new BigNumber((_dafny.Seq.UnicodeFromString("qr")).length);
              let _66_v15;
              _66_v15 = _dafny.Set.fromElements(_49_v1);
              let _67_v16;
              _67_v16 = _dafny.Map.Empty.slice().updateUnsafe((_60_v10).f9,(_60_v10).f9);
              let _68_v17;
              _68_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm16((_60_v10).f10, _67_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(970)), ((_69_v2) => function (_70_i2) {
                return _69_v2;
              })(_50_v2)), _53_globalState));
              let _71_v18;
              _71_v18 = _module.D19.create_DC51(_68_v17);
              let _72_v19;
              _72_v19 = _dafny.Map.Empty.slice().updateUnsafe(_48_v0,(_60_v10).f10);
              let _73_v21;
              _73_v21 = _dafny.Seq.of((_60_v10).f9);
              let _74_v22;
              _74_v22 = _dafny.Seq.of(_54_v4);
              let _75_v23;
              let _nw4 = Array((new BigNumber(27)).toNumber());
              _nw4[(_dafny.ZERO).toNumber()] = (_60_v10).f9;
              _nw4[(_dafny.ONE).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(2)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(3)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(4)).toNumber()] = new BigNumber((_66_v15).length);
              _nw4[(new BigNumber(5)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(6)).toNumber()] = (_60_v10).fm4(new BigNumber(825), (_dafny.ZERO).minus(new BigNumber(((_71_v18).dtor_cf81).length)), _72_v19, (_60_v10).f9, _53_globalState);
              _nw4[(new BigNumber(7)).toNumber()] = _49_v1;
              _nw4[(new BigNumber(8)).toNumber()] = new BigNumber((_48_v0).length);
              _nw4[(new BigNumber(9)).toNumber()] = new BigNumber((function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of _dafny.IntegerRange(new BigNumber(473), new BigNumber(327))) {
                  let _76_v20 = _compr_11;
                  if (((new BigNumber(473)).isLessThanOrEqualTo(_76_v20)) && ((_76_v20).isLessThan(new BigNumber(327)))) {
                    _coll11.push([_module.__default.safeDivisionInt(_76_v20, (_60_v10).f9),_50_v2]);
                  }
                }
                return _coll11;
              }()).length);
              _nw4[(new BigNumber(10)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(11)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(12)).toNumber()] = _49_v1;
              _nw4[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-561)), ((_77_v2) => function (_78_i3) {
                return _77_v2;
              })(_50_v2))).length);
              _nw4[(new BigNumber(14)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(15)).toNumber()] = new BigNumber((_73_v21).length);
              _nw4[(new BigNumber(16)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(17)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(18)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(19)).toNumber()] = new BigNumber(307);
              _nw4[(new BigNumber(20)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(21)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(22)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(23)).toNumber()] = new BigNumber((_48_v0).length);
              _nw4[(new BigNumber(24)).toNumber()] = (_60_v10).fm4((_60_v10).f9, new BigNumber((_74_v22).length), _72_v19, (_60_v10).f9, _53_globalState);
              _nw4[(new BigNumber(25)).toNumber()] = (_60_v10).f9;
              _nw4[(new BigNumber(26)).toNumber()] = new BigNumber((_74_v22).length);
              _75_v23 = _nw4;
              let _79_v24;
              _79_v24 = _dafny.Seq.of(_75_v23);
              let _80_v25;
              _80_v25 = _dafny.Set.fromElements((_60_v10).f10);
              let _81_v26;
              _81_v26 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))],(_60_v10).f10)).length), (_dafny.ZERO).minus((_60_v10).f9), (_60_v10).f9);
              let _82_v28;
              _82_v28 = _dafny.Set.fromElements(_48_v0);
              let _83_v29;
              _83_v29 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)), _50_v2, _50_v2);
              let _84_v30;
              let _nw5 = Array((new BigNumber(21)).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_79_v24, _79_v24)).length);
              _nw5[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_80_v25).length), (_60_v10).f9);
              _nw5[(new BigNumber(2)).toNumber()] = (_60_v10).f9;
              _nw5[(new BigNumber(3)).toNumber()] = new BigNumber(733);
              _nw5[(new BigNumber(4)).toNumber()] = _49_v1;
              _nw5[(new BigNumber(5)).toNumber()] = new BigNumber((_81_v26).cardinality());
              _nw5[(new BigNumber(6)).toNumber()] = (_60_v10).f9;
              _nw5[(new BigNumber(7)).toNumber()] = _49_v1;
              _nw5[(new BigNumber(8)).toNumber()] = _49_v1;
              _nw5[(new BigNumber(9)).toNumber()] = ((_60_v10).f9).plus(new BigNumber((_66_v15).length));
              _nw5[(new BigNumber(10)).toNumber()] = (_60_v10).fm4((_60_v10).f9, new BigNumber(520), (function () {
                let _coll12 = new _dafny.Map();
                for (const _compr_12 of (_82_v28).Elements) {
                  let _85_v27 = _compr_12;
                  if ((_82_v28).contains(_85_v27)) {
                    _coll12.push([_85_v27,(_60_v10).f10]);
                  }
                }
                return _coll12;
              }()).update(_48_v0, _54_v4), new BigNumber((_83_v29).cardinality()), _53_globalState);
              _nw5[(new BigNumber(11)).toNumber()] = new BigNumber(588);
              _nw5[(new BigNumber(12)).toNumber()] = _49_v1;
              _nw5[(new BigNumber(13)).toNumber()] = (_60_v10).f9;
              _nw5[(new BigNumber(14)).toNumber()] = (((_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))]) ? ((_60_v10).f9) : ((_60_v10).fm4(new BigNumber((_68_v17).length), (_dafny.ZERO).minus((_60_v10).f9), _72_v19, _49_v1, _53_globalState)));
              _nw5[(new BigNumber(15)).toNumber()] = ((_60_v10).f9).plus((_60_v10).f9);
              _nw5[(new BigNumber(16)).toNumber()] = (_60_v10).f9;
              _nw5[(new BigNumber(17)).toNumber()] = (_60_v10).f9;
              _nw5[(new BigNumber(18)).toNumber()] = ((_60_v10).f9).plus(new BigNumber((_73_v21).length));
              _nw5[(new BigNumber(19)).toNumber()] = _49_v1;
              _nw5[(new BigNumber(20)).toNumber()] = new BigNumber((_48_v0).length);
              _84_v30 = _nw5;
              let _index3 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_84_v30).length));
              (_84_v30)[_index3] = (_60_v10).f9;
              let _86_v31;
              _86_v31 = _module.D19.create_DC52(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_60_v10).f10),_49_v1)).length));
              let _index4 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_84_v30).length));
              (_84_v30)[_index4] = _module.__default.safeDivisionInt((_86_v31).dtor_cf82, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_60_v10).f9,true)).Merge(function () {
                let _coll13 = new _dafny.Map();
                for (const _compr_13 of _dafny.IntegerRange(new BigNumber(779), new BigNumber(783))) {
                  let _87_v32 = _compr_13;
                  if (((new BigNumber(779)).isLessThanOrEqualTo(_87_v32)) && ((_87_v32).isLessThan(new BigNumber(783)))) {
                    _coll13.push([(_87_v32).multipliedBy(new BigNumber(742)),_54_v4]);
                  }
                }
                return _coll13;
              }())).length));
              let _index5 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_84_v30).length));
              (_84_v30)[_index5] = _49_v1;
              _54_v4 = (_60_v10).f10;
              let _88_v33;
              let _nw6 = Array((new BigNumber(3)).toNumber()).fill(false);
              _88_v33 = _nw6;
              let _89_v34;
              let _nw7 = new _module.C4();
              _nw7.__ctor(_88_v33, _54_v4, false, _59_v9);
              _89_v34 = _nw7;
              let _90_v35;
              _90_v35 = _dafny.Map.Empty.slice().updateUnsafe(_54_v4,_89_v34);
              (_53_globalState).f0 = (_60_v10).fm4(new BigNumber((_90_v35).length), (_60_v10).f9, _72_v19, new BigNumber((_48_v0).length), _53_globalState);
            }
            let _91_v36;
            _91_v36 = _dafny.Map.Empty.slice().updateUnsafe(_59_v9,_48_v0);
            let _92_v37;
            _92_v37 = _module.D16.create_DC44(_54_v4, (_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))], new BigNumber((_91_v36).length));
            let _93_v38;
            _93_v38 = _dafny.Map.Empty.slice().updateUnsafe(_92_v37,_49_v1);
            let _94_v39;
            _94_v39 = _dafny.Map.Empty.slice().updateUnsafe(_49_v1,(_60_v10).f10);
            let _95_v40;
            _95_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rwq"),_54_v4);
            let _96_v41;
            _96_v41 = _dafny.Seq.of((((_94_v39).contains((_60_v10).f9)) ? ((_94_v39).get((_60_v10).f9)) : ((_60_v10).f10)), (_60_v10).f10, (((_94_v39).contains((_60_v10).fm4((_60_v10).f9, (_60_v10).f9, _95_v40, (_60_v10).f9, _53_globalState))) ? ((_94_v39).get((_60_v10).fm4((_60_v10).f9, (_60_v10).f9, _95_v40, (_60_v10).f9, _53_globalState))) : (!(true))), (_60_v10).f10);
            (_60_v10).m1((((_93_v38).contains(_92_v37)) ? ((_93_v38).get(_92_v37)) : ((_60_v10).f9)), (_56_v6)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length))], (_96_v41)[_module.__default.safeIndex((_60_v10).f9, new BigNumber((_96_v41).length))], _53_globalState);
            let _97_v42;
            _97_v42 = _dafny.Seq.of((_60_v10).f9);
            let _98_v43;
            let _nw8 = new _module.C1();
            _nw8.__ctor((_60_v10).f10, _48_v0, _54_v4, _59_v9);
            _98_v43 = _nw8;
            let _99_v44;
            _99_v44 = _dafny.Map.Empty.slice().updateUnsafe(_98_v43,(_60_v10).f10);
            let _100_v45;
            _100_v45 = _dafny.Seq.of(_dafny.Seq.update(_96_v41, _module.__default.safeIndex(new BigNumber(-771), new BigNumber((_96_v41).length)), !((_98_v43).f6)));
            let _index6 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            let _index7 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            let _rhs2 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-143)), function (_101_i4) {
              return new _dafny.CodePoint('e'.codePointAt(0));
            }), _48_v0);
            let _rhs3 = _dafny.Seq.update(_dafny.Seq.Concat(_97_v42, _dafny.Seq.of((_60_v10).f9, new BigNumber((_99_v44).length), _49_v1, (_dafny.ZERO).minus((_60_v10).f9))), _module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_96_v41).length), (_60_v10).f9), new BigNumber((_dafny.Seq.Concat(_97_v42, _dafny.Seq.of((_60_v10).f9, new BigNumber((_99_v44).length), _49_v1, (_dafny.ZERO).minus((_60_v10).f9)))).length)), (_dafny.ZERO).minus(new BigNumber(-259)));
            let _rhs4 = (_54_v4) || ((_60_v10).f10);
            let _rhs5 = !_dafny.Seq.contains(_100_v45, _dafny.Seq.Concat(_96_v41, _96_v41));
            let _rhs6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_96_v41, _dafny.Seq.of((_98_v43).f6, (_60_v10).f10)), _dafny.Seq.Concat(_96_v41, _96_v41));
            let _lhs3 = _56_v6;
            let _lhs4 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            let _lhs5 = _56_v6;
            let _lhs6 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_56_v6).length));
            _48_v0 = _rhs2;
            _97_v42 = _rhs3;
            _lhs3[_lhs4] = _rhs4;
            _lhs5[_lhs6] = _rhs5;
            _96_v41 = _rhs6;
          }
        }
      }
      _49_v1 = _49_v1;
      let _102_v46;
      _102_v46 = _dafny.Map.Empty.slice().updateUnsafe(true,_48_v0);
      _102_v46 = _102_v46;
      let _103_v47;
      _103_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-434),(_60_v10).f9),_54_v4);
      let _104_v48;
      let _nw9 = new _module.C0();
      _nw9.__ctor(_103_v47);
      _104_v48 = _nw9;
      let _105_v49;
      let _init1 = ((_106_v1) => function (_107_i5) {
        return (_107_i5).plus(_106_v1);
      })(_49_v1);
      let _nw10 = Array((new BigNumber(12)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw10.length); _i0_1++) {
        _nw10[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _105_v49 = _nw10;
      let _index8 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      (_105_v49)[_index8] = (_60_v10).f9;
      let _108_v51;
      let _init2 = ((_109_v10, _110_v4) => function (_111_i6) {
        return function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.MultiSet.fromElements(_module.D8.create_DC24((_109_v10).f9, _110_v4, _dafny.Seq.of((_109_v10).f10), new BigNumber(959), (_109_v10).f9))).Elements) {
            let _112_v50 = _compr_14;
            if ((_dafny.MultiSet.fromElements(_module.D8.create_DC24((_109_v10).f9, _110_v4, _dafny.Seq.of((_109_v10).f10), new BigNumber(959), (_109_v10).f9))).contains(_112_v50)) {
              _coll14.push([_112_v50,_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))]);
            }
          }
          return _coll14;
        }();
      })(_60_v10, _54_v4);
      let _nw11 = Array((new BigNumber(20)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
        _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _108_v51 = _nw11;
      let _113_v52;
      _113_v52 = _dafny.Seq.of((_60_v10).f10, (_60_v10).f10, (_60_v10).fm1((_60_v10).f10, new BigNumber(-253), (_60_v10).f9, _53_globalState));
      let _114_v53;
      _114_v53 = _module.D8.create_DC24(new BigNumber((_dafny.MultiSet.FromArray(_113_v52)).cardinality()), (_60_v10).f10, _113_v52, (_60_v10).f9, _49_v1);
      let _115_v54;
      _115_v54 = _dafny.Map.Empty.slice().updateUnsafe(_49_v1,(_60_v10).f9);
      let _116_v55;
      _116_v55 = _dafny.Map.Empty.slice().updateUnsafe(_114_v53,_dafny.Seq.of(_50_v2, _50_v2, _module.__default.fm16(_54_v4, _115_v54, _48_v0, _53_globalState), _50_v2));
      let _index9 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_108_v51).length));
      (_108_v51)[_index9] = _116_v55;
      let _117_v57;
      _117_v57 = _dafny.Seq.of(_114_v53, _114_v53);
      let _118_v58;
      _118_v58 = _dafny.Seq.of(function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_117_v57).Elements) {
          let _119_v56 = _compr_15;
          if (_dafny.Seq.contains(_117_v57, _119_v56)) {
            _coll15.push([_119_v56,_48_v0]);
          }
        }
        return _coll15;
      }());
      let _index10 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      let _index11 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_108_v51).length));
      let _rhs7 = _104_v48;
      let _rhs8 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_60_v10).f9), (_60_v10).f9);
      let _rhs9 = ((_118_v58)[_module.__default.safeIndex(_49_v1, new BigNumber((_118_v58).length))]).Merge(_116_v55);
      let _rhs10 = (_60_v10).f9;
      let _rhs11 = (((_60_v10).f10) ? (_54_v4) : (_54_v4));
      let _lhs7 = _105_v49;
      let _lhs8 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      let _lhs9 = _108_v51;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_108_v51).length));
      let _lhs11 = _53_globalState;
      _104_v48 = _rhs7;
      _lhs7[_lhs8] = _rhs8;
      _lhs9[_lhs10] = _rhs9;
      _lhs11.f0 = _rhs10;
      _54_v4 = _rhs11;
      let _120_v60;
      _120_v60 = _dafny.Set.fromElements((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]);
      let _source2 = _module.__default.fm36((_115_v54).Merge(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_115_v54).Keys.Elements) {
          let _121_v59 = _compr_16;
          if ((_115_v54).contains(_121_v59)) {
            _coll16.push([(_121_v59).multipliedBy(new BigNumber(-411)),(_60_v10).f9]);
          }
        }
        return _coll16;
      }()), _54_v4, _120_v60, _53_globalState);
      if (_source2.is_DC24) {
        let _122___mcc_h0 = (_source2).cf34;
        let _123___mcc_h1 = (_source2).cf35;
        let _124___mcc_h2 = (_source2).cf36;
        let _125___mcc_h3 = (_source2).cf37;
        let _126___mcc_h4 = (_source2).cf38;
        let _127_cf38 = _126___mcc_h4;
        let _128_cf37 = _125___mcc_h3;
        let _129_cf36 = _124___mcc_h2;
        let _130_cf35 = _123___mcc_h1;
        let _131_cf34 = _122___mcc_h0;
        let _index12 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
        (_105_v49)[_index12] = ((_dafny.ZERO).minus((_127_cf38).multipliedBy(_128_cf37))).minus((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]);
        if ((_60_v10).f10) {
          let _index13 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          (_105_v49)[_index13] = new BigNumber(59);
          let _132_v61;
          _132_v61 = _dafny.Map.Empty.slice().updateUnsafe((_60_v10).f9,_130_cf35);
          _54_v4 = (((_132_v61).contains(_128_cf37)) ? ((_132_v61).get(_128_cf37)) : ((_60_v10).f10));
          _127_cf38 = (_60_v10).f9;
          let _index14 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_56_v6).length));
          (_56_v6)[_index14] = _54_v4;
          let _133_v62;
          _133_v62 = _dafny.Map.Empty.slice().updateUnsafe(_48_v0,!((_60_v10).f10));
          let _134_v63;
          _134_v63 = _dafny.Map.Empty.slice().updateUnsafe(_130_cf35,_54_v4);
          let _index15 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_56_v6).length));
          (_56_v6)[_index15] = (_60_v10).fm1(!((_60_v10).f10) || (_54_v4), _module.__default.safeDivisionInt((_60_v10).fm4((_60_v10).f9, (_60_v10).f9, _133_v62, (_60_v10).f9, _53_globalState), new BigNumber(((_134_v63).update((_60_v10).fm2((_60_v10).f10, _53_globalState), (_60_v10).f10)).length)), (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], _53_globalState);
          _48_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), ((_135_v2) => function (_136_i7) {
            return _135_v2;
          })(_50_v2));
        } else {
          (_60_v10).m2(_50_v2, _53_globalState);
          _54_v4 = !(_130_cf35);
          _48_v0 = _48_v0;
          let _137_v64;
          let _nw12 = Array((new BigNumber(15)).toNumber()).fill([]);
          _137_v64 = _nw12;
          let _138_v65;
          let _nw13 = new _module.C2();
          _nw13.__ctor((_59_v9));
          _138_v65 = _nw13;
          let _139_v66;
          _139_v66 = _module.D9.create_DC26(_138_v65);
          let _140_v67;
          _140_v67 = _dafny.MultiSet.fromElements((_60_v10).f9);
          let _141_v68;
          _141_v68 = _dafny.Set.fromElements((_60_v10).f10, (_60_v10).f10);
          let _rhs12 = _dafny.Seq.IsPrefixOf(_129_cf36, _module.__default.fm23((((_140_v67).contains((_dafny.ZERO).minus(new BigNumber((_141_v68).length)))) ? ((_140_v67).get((_dafny.ZERO).minus(new BigNumber((_141_v68).length)))) : (_131_cf34)), _53_globalState));
          let _rhs13 = _137_v64;
          let _rhs14 = _50_v2;
          let _rhs15 = _139_v66;
          _130_cf35 = _rhs12;
          _137_v64 = _rhs13;
          _50_v2 = _rhs14;
          _139_v66 = _rhs15;
          let _index16 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_56_v6).length));
          (_56_v6)[_index16] = _54_v4;
          let _index17 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_56_v6).length));
          (_56_v6)[_index17] = (_60_v10).f10;
        }
        let _142_v69;
        _142_v69 = _dafny.Map.Empty.slice().updateUnsafe(_54_v4,_dafny.Set.fromElements(_131_cf34));
        let _143_v70;
        _143_v70 = _dafny.Set.fromElements((((_142_v69).contains(_54_v4)) ? ((_142_v69).get(_54_v4)) : (_120_v60)), _module.__default.fm24(new BigNumber((_48_v0).length), true, _49_v1, _53_globalState));
        _143_v70 = function () {
          let _coll17 = new _dafny.Set();
          for (const _compr_17 of (_143_v70).Elements) {
            let _144_v71 = _compr_17;
            if ((_143_v70).contains(_144_v71)) {
              _coll17.add(_144_v71);
            }
          }
          return _coll17;
        }();
        let _145_v72;
        let _nw14 = Array((new BigNumber(24)).toNumber()).fill([]);
        _145_v72 = _nw14;
        let _index18 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_145_v72).length));
        (_145_v72)[_index18] = _51_v3;
        let _index19 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_145_v72).length));
        (_145_v72)[_index19] = _51_v3;
      } else if (_source2.is_DC25) {
        let _146___mcc_h5 = (_source2).cf39;
        let _147___mcc_h6 = (_source2).cf40;
        let _148_cf40 = _147___mcc_h6;
        let _149_cf39 = _146___mcc_h5;
        let _150_v73;
        _150_v73 = _dafny.Map.Empty.slice().updateUnsafe((_60_v10).f10,_105_v49);
        _150_v73 = _150_v73;
        let _rhs16 = (_60_v10).f9;
        let _rhs17 = _148_cf40;
        let _rhs18 = !((_60_v10).f10) || (((_dafny.Map.Empty.slice().updateUnsafe((_60_v10).f10,_105_v49)).update(false, _105_v49)).contains(_148_cf40));
        let _rhs19 = _48_v0;
        let _rhs20 = true;
        let _lhs12 = _53_globalState;
        _lhs12.f0 = _rhs16;
        _54_v4 = _rhs17;
        _149_cf39 = _rhs18;
        _48_v0 = _rhs19;
        _149_cf39 = _rhs20;
        _50_v2 = new _dafny.CodePoint('d'.codePointAt(0));
        (_53_globalState).f0 = _49_v1;
      } else {
        let _151___mcc_h7 = (_source2).cf33;
        let _152_cf33 = _151___mcc_h7;
        if ((_60_v10).f10) {
          (_53_globalState).f0 = (_60_v10).f9;
          let _index20 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_56_v6).length));
          (_56_v6)[_index20] = !((_60_v10).f10);
          let _index21 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_56_v6).length));
          (_56_v6)[_index21] = !((_60_v10).f10);
          let _153_v74;
          _153_v74 = _dafny.Map.Empty.slice().updateUnsafe(_54_v4,(_56_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_56_v6).length))]);
          let _154_v75;
          _154_v75 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_54_v4, (_60_v10).f10),_153_v74);
          let _155_v77;
          _155_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_54_v4, !(_54_v4), (_60_v10).fm1((_60_v10).f10, new BigNumber((_48_v0).length), new BigNumber((_113_v52).length), _53_globalState), (_60_v10).f10),new BigNumber((_153_v74).length));
          let _156_v79;
          _156_v79 = _dafny.Set.fromElements(_59_v9);
          _154_v75 = (_154_v75).Merge((function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_155_v77).Keys.Elements) {
              let _157_v76 = _compr_18;
              if ((_155_v77).contains(_157_v76)) {
                _coll18.push([_157_v76,_153_v74]);
              }
            }
            return _coll18;
          }()).Merge(function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_156_v79).Elements) {
              let _158_v78 = _compr_19;
              if ((_156_v79).contains(_158_v78)) {
                _coll19.push([_158_v78,_153_v74]);
              }
            }
            return _coll19;
          }()));
          (_53_globalState).f0 = (_60_v10).f9;
          let _159_v80;
          let _init3 = ((_160_v10) => function (_161_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe((_160_v10).f10,new BigNumber(-120));
          })(_60_v10);
          let _nw15 = Array((new BigNumber(19)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw15.length); _i0_3++) {
            _nw15[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _159_v80 = _nw15;
          let _162_v81;
          let _163_v82;
          let _164_v83;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_60_v10).m5(new BigNumber(-675), _59_v9, _159_v80, _54_v4, _53_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _162_v81 = _out0;
          _163_v82 = _out1;
          _164_v83 = _out2;
        } else {
          let _165_v84;
          _165_v84 = _dafny.Seq.of(_49_v1);
          (_60_v10).m4((_60_v10).fm2((_60_v10).f10, _53_globalState), _165_v84, _53_globalState);
          let _166_v85;
          let _nw16 = Array((new BigNumber(12)).toNumber()).fill([]);
          _166_v85 = _nw16;
          let _index22 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_166_v85).length));
          (_166_v85)[_index22] = _56_v6;
          let _index23 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_166_v85).length));
          (_166_v85)[_index23] = _56_v6;
          let _167_v86;
          _167_v86 = _dafny.Seq.of(_104_v48, _104_v48);
          let _168_v87;
          _168_v87 = _module.D2.create_DC7(_50_v2, _115_v54, (_60_v10).f9);
          let _169_v88;
          _169_v88 = _dafny.Map.Empty.slice().updateUnsafe(_48_v0,_54_v4);
          let _170_v89;
          let _nw17 = Array((new BigNumber(28)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = _104_v48;
          _nw17[(_dafny.ONE).toNumber()] = _104_v48;
          _nw17[(new BigNumber(2)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(3)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(4)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(5)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(6)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(7)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(8)).toNumber()] = (_167_v86)[_module.__default.safeIndex((_60_v10).fm4(_49_v1, (_168_v87).dtor_cf8, _169_v88, _49_v1, _53_globalState), new BigNumber((_167_v86).length))];
          _nw17[(new BigNumber(9)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(10)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(11)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(12)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(13)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(14)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(15)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(16)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(17)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(18)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(19)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(20)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(21)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(22)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(23)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(24)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(25)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(26)).toNumber()] = _104_v48;
          _nw17[(new BigNumber(27)).toNumber()] = _104_v48;
          _170_v89 = _nw17;
          let _index24 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_170_v89).length));
          (_170_v89)[_index24] = _104_v48;
          let _index25 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_170_v89).length));
          let _rhs21 = !((new BigNumber(381)).isLessThan((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]));
          let _rhs22 = (_60_v10).f9;
          let _rhs23 = _104_v48;
          let _lhs13 = _53_globalState;
          let _lhs14 = _170_v89;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_170_v89).length));
          _54_v4 = _rhs21;
          _lhs13.f0 = _rhs22;
          _lhs14[_lhs15] = _rhs23;
          let _index26 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          (_105_v49)[_index26] = (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))];
          let _171_v90;
          let _nw18 = new _module.C2();
          _nw18.__ctor(_59_v9);
          _171_v90 = _nw18;
          let _rhs24 = (_60_v10).f10;
          let _rhs25 = _171_v90;
          _54_v4 = _rhs24;
          _171_v90 = _rhs25;
        }
        let _172_v91;
        _172_v91 = _dafny.Seq.of(_49_v1);
        let _173_v92;
        _173_v92 = _module.D15.create_DC42(_dafny.MultiSet.FromArray(_172_v91));
        let _174_v93;
        _174_v93 = _dafny.Seq.of(_module.D15.create_DC42(_dafny.MultiSet.FromArray(_172_v91)), _173_v92);
        let _175_v94;
        _175_v94 = _module.D21.create_DC54(_174_v93);
        let _176_v95;
        _176_v95 = _dafny.MultiSet.fromElements((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], (_60_v10).f9);
        if ((new BigNumber((_dafny.Seq.Concat((_175_v94).dtor_cf84, _dafny.Seq.of(_module.D15.create_DC42(_176_v95), _module.D15.create_DC42(_176_v95), _173_v92))).length)).isLessThan(_49_v1)) {
          let _177_v96;
          let _nw19 = new _module.C2();
          _nw19.__ctor(_59_v9);
          _177_v96 = _nw19;
          let _178_v97;
          _178_v97 = _dafny.Map.Empty.slice().updateUnsafe(_177_v96,(_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]);
          let _179_v98;
          _179_v98 = _dafny.Seq.of(_56_v6);
          let _180_v99;
          _180_v99 = _dafny.Map.Empty.slice().updateUnsafe(_179_v98,(_60_v10).f10);
          (_60_v10).m1(new BigNumber((_178_v97).length), (((_180_v99).contains(_dafny.Seq.of(_56_v6, _56_v6))) ? ((_180_v99).get(_dafny.Seq.of(_56_v6, _56_v6))) : ((_60_v10).f10)), (_60_v10).f10, _53_globalState);
          let _181_v100;
          let _nw20 = Array((new BigNumber(11)).toNumber());
          _181_v100 = _nw20;
          let _index27 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_181_v100).length));
          (_181_v100)[_index27] = _60_v10;
          let _index28 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_181_v100).length));
          (_181_v100)[_index28] = _60_v10;
          let _index29 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          (_105_v49)[_index29] = (new BigNumber(569)).plus(new BigNumber((_113_v52).length));
          let _182_v101;
          _182_v101 = _dafny.Map.Empty.slice().updateUnsafe((_60_v10).f10,(_60_v10).f10);
          let _183_v102;
          _183_v102 = _module.D13.create_DC39((_dafny.Map.Empty.slice().updateUnsafe(_50_v2,_54_v4)).Merge(_55_v5), _56_v6, _113_v52, new BigNumber((_182_v101).length), _dafny.Seq.Concat(_172_v91, _172_v91));
          _183_v102 = _module.D13.create_DC39(_module.__default.fm37(_50_v2, _53_globalState), _56_v6, _113_v52, ((_60_v10).f9).minus(new BigNumber(134)), ((true) ? (_172_v91) : (_module.__default.fm18(_53_globalState))));
          _54_v4 = (_dafny.MultiSet.FromArray(_113_v52)).contains(!((_60_v10).f10) || (_54_v4));
        } else {
          _50_v2 = _50_v2;
          let _index30 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          (_105_v49)[_index30] = _49_v1;
          let _index31 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_56_v6).length));
          (_56_v6)[_index31] = (_60_v10).f10;
          let _index32 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_56_v6).length));
          (_56_v6)[_index32] = _dafny.Seq.IsPrefixOf(_48_v0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), _dafny.Seq.UnicodeFromString("btussn")));
          let _184_v103;
          let _nw21 = new _module.C6();
          _nw21.__ctor(_120_v60, _50_v2);
          _184_v103 = _nw21;
          let _185_v104;
          _185_v104 = _dafny.Map.Empty.slice().updateUnsafe(_184_v103,(_56_v6)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_56_v6).length))]);
          _185_v104 = (_185_v104).update(((_54_v4) ? (_184_v103) : (_184_v103)), (_60_v10).f10);
          let _index33 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          let _index34 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_56_v6).length));
          let _rhs26 = (_60_v10).f9;
          let _rhs27 = (_60_v10).f10;
          let _lhs16 = _105_v49;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          let _lhs18 = _56_v6;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_56_v6).length));
          _lhs16[_lhs17] = _rhs26;
          _lhs18[_lhs19] = _rhs27;
        }
        let _186_v105;
        let _nw22 = new _module.C3();
        _nw22.__ctor(_dafny.Seq.IsPrefixOf(_113_v52, _113_v52), _59_v9);
        _186_v105 = _nw22;
        _54_v4 = (_60_v10).f10;
      }
      let _187_v106;
      let _init4 = ((_188_v52) => function (_189_i9) {
        return _module.D6.create_DC17(_188_v52);
      })(_113_v52);
      let _nw23 = Array((new BigNumber(3)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw23.length); _i0_4++) {
        _nw23[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _187_v106 = _nw23;
      let _index35 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_187_v106).length));
      (_187_v106)[_index35] = _module.D6.create_DC17(_113_v52);
      let _190_v107;
      _190_v107 = _module.D6.create_DC17(_dafny.Seq.of(false));
      let _pat_let_tv0 = _113_v52;
      let _index36 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_187_v106).length));
      (_187_v106)[_index36] = function (_pat_let0_0) {
        return function (_191_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_192_dt__update_hcf24_h0) {
              return _module.D6.create_DC17(_192_dt__update_hcf24_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_190_v107);
      _54_v4 = _54_v4;
      let _193_v108;
      let _nw24 = new _module.C3();
      _nw24.__ctor((_60_v10).f10, _59_v9);
      _193_v108 = _nw24;
      let _194_v109;
      _194_v109 = _dafny.Set.fromElements(_193_v108, _193_v108, _193_v108);
      _194_v109 = _194_v109;
      let _195_v110;
      _195_v110 = _dafny.Set.fromElements((_60_v10).f10, (_60_v10).fm1(_54_v4, _49_v1, new BigNumber((_120_v60).length), _53_globalState), _54_v4);
      let _196_v111;
      _196_v111 = _dafny.Map.Empty.slice().updateUnsafe(_48_v0,false);
      let _index37 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      let _rhs28 = ((_60_v10).fm4(new BigNumber((_48_v0).length), _49_v1, _196_v111, new BigNumber((_48_v0).length), _53_globalState)).plus(new BigNumber(-538));
      let _rhs29 = (_60_v10).f10;
      let _rhs30 = _dafny.Set.fromElements(!((_60_v10).f10));
      let _rhs31 = (_113_v52)[_module.__default.safeIndex((_60_v10).f9, new BigNumber((_113_v52).length))];
      let _lhs20 = _105_v49;
      let _lhs21 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      _lhs20[_lhs21] = _rhs28;
      _54_v4 = _rhs29;
      _195_v110 = _rhs30;
      _54_v4 = _rhs31;
      let _197_v112;
      _197_v112 = _module.D17.create_DC47(_54_v4, (_60_v10).fm4((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_60_v10).f10,(_60_v10).fm4((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], new BigNumber((_48_v0).length), _196_v111, (_60_v10).f9, _53_globalState))).length), _196_v111, (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], _53_globalState), (_60_v10).f9);
      if ((_197_v112).dtor_cf74) {
        (_193_v108).m1(_49_v1, (_60_v10).f10, (_60_v10).f10, _53_globalState);
        let _198_v113;
        _198_v113 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wfd"),_195_v110);
        _198_v113 = (_198_v113).update(_dafny.Seq.Concat(_48_v0, _48_v0), _195_v110);
        (_60_v10).m1((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))], (_60_v10).f10, true, _53_globalState);
        if ((_113_v52)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_113_v52).length))]) {
          _54_v4 = ((_54_v4) || (_54_v4)) || (((_115_v54).update(new BigNumber(8), (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))])).contains(_49_v1));
          _54_v4 = (((_60_v10).f10) ? ((_60_v10).f10) : ((_54_v4) || ((_60_v10).f10)));
          let _199_v114;
          _199_v114 = _dafny.MultiSet.fromElements(_195_v110, _dafny.Set.fromElements(_54_v4), _195_v110, _195_v110);
          let _200_v115;
          let _nw25 = new _module.C3();
          _nw25.__ctor((_60_v10).fm1((_60_v10).f10, _49_v1, new BigNumber((_199_v114).cardinality()), _53_globalState), (_59_v9).Union(_59_v9));
          _200_v115 = _nw25;
          let _rhs32 = (_60_v10).f9;
          let _rhs33 = _module.__default.safeDivisionInt(_49_v1, (_dafny.ZERO).minus((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]));
          let _rhs34 = true;
          let _lhs22 = _53_globalState;
          let _lhs23 = _53_globalState;
          _lhs22.f0 = _rhs32;
          _lhs23.f0 = _rhs33;
          _54_v4 = _rhs34;
          let _201_v116;
          let _nw26 = Array((new BigNumber(16)).toNumber());
          _201_v116 = _nw26;
          let _202_v117;
          let _nw27 = new _module.C5();
          _nw27.__ctor((_60_v10).f9, (_60_v10).f10, _56_v6, (_60_v10).f10, _54_v4, _59_v9);
          _202_v117 = _nw27;
          let _index38 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_201_v116).length));
          (_201_v116)[_index38] = _202_v117;
          let _203_v118;
          let _nw28 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _203_v118 = _nw28;
          let _index39 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          let _index40 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_201_v116).length));
          let _rhs35 = (_60_v10).f9;
          let _rhs36 = _202_v117;
          let _rhs37 = _203_v118;
          let _rhs38 = _48_v0;
          let _lhs24 = _105_v49;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
          let _lhs26 = _201_v116;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_201_v116).length));
          _lhs24[_lhs25] = _rhs35;
          _lhs26[_lhs27] = _rhs36;
          _203_v118 = _rhs37;
          _48_v0 = _rhs38;
        } else {
          let _204_v119;
          let _nw29 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _204_v119 = _nw29;
          let _205_v120;
          _205_v120 = _dafny.Seq.of(_105_v49);
          let _index41 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_204_v119).length));
          (_204_v119)[_index41] = _dafny.Seq.Concat(_205_v120, _dafny.Seq.of(_105_v49));
          let _index42 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_204_v119).length));
          (_204_v119)[_index42] = _dafny.Seq.of(_105_v49, _105_v49);
          let _206_v121;
          _206_v121 = _dafny.MultiSet.fromElements(_49_v1);
          _206_v121 = _206_v121;
          let _207_v122;
          let _nw30 = new _module.C6();
          _nw30.__ctor(_dafny.Set.fromElements((_60_v10).f9), _50_v2);
          _207_v122 = _nw30;
          _54_v4 = (_60_v10).f10;
          (_53_globalState).f0 = (_60_v10).f9;
        }
        let _208_v123;
        _208_v123 = _dafny.Map.Empty.slice().updateUnsafe((_54_v4) || (!((_60_v10).f10)),(false) === ((_60_v10).f10));
        let _209_v124;
        _209_v124 = _module.D12.create_DC35(_54_v4);
        _208_v123 = (_208_v123).update(!_dafny.areEqual(_dafny.Seq.of(!((_209_v124).dtor_cf55)), _113_v52), _54_v4);
      } else {
        let _210_v125;
        _210_v125 = _module.D3.create_DC9(_105_v49);
        let _211_v126;
        _211_v126 = _module.D5.create_DC16(_210_v125, new BigNumber(941), (_60_v10).f9, new BigNumber(-528), _104_v48);
        let _212_v127;
        _212_v127 = _dafny.Map.Empty.slice().updateUnsafe(_211_v126,_105_v49);
        _105_v49 = (((_212_v127).contains(_211_v126)) ? ((_212_v127).get(_211_v126)) : (_105_v49));
        let _213_v128;
        let _nw31 = new _module.C3();
        _nw31.__ctor((_60_v10).f10, (_dafny.MultiSet.fromElements(_54_v4, !((_60_v10).f10), (_60_v10).f10, (_193_v108).fm3(_53_globalState))).Union(_59_v9));
        _213_v128 = _nw31;
        let _index43 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_56_v6).length));
        (_56_v6)[_index43] = _54_v4;
        let _index44 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_56_v6).length));
        (_56_v6)[_index44] = _54_v4;
        let _214_v129;
        let _nw32 = new _module.C2();
        _nw32.__ctor(_dafny.MultiSet.FromArray(_dafny.Seq.of(_54_v4)));
        _214_v129 = _nw32;
        let _215_v130;
        _215_v130 = _dafny.MultiSet.fromElements(_214_v129, _214_v129);
        let _216_v131;
        _216_v131 = _module.D22.create_DC57(_214_v129);
        let _217_v132;
        _217_v132 = _dafny.Seq.of(_214_v129);
        let _218_v133;
        let _init5 = ((_219_v6) => function (_220_i10) {
          return (_219_v6)[_module.__default.safeIndex(new BigNumber(989), new BigNumber((_219_v6).length))];
        })(_56_v6);
        let _nw33 = Array((new BigNumber(10)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
          _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _218_v133 = _nw33;
        let _221_v134;
        let _nw34 = new _module.C4();
        _nw34.__ctor(_218_v133, (_60_v10).f10, _54_v4, _59_v9);
        _221_v134 = _nw34;
        let _222_v135;
        _222_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_48_v0).length),_221_v134);
        let _223_v136;
        _223_v136 = _dafny.Map.Empty.slice().updateUnsafe((((_222_v135).contains(_49_v1)) ? ((_222_v135).get(_49_v1)) : (_221_v134)),_215_v130);
        let _224_v137;
        let _nw35 = Array((new BigNumber(23)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = (_215_v130).Union(_dafny.MultiSet.fromElements(_214_v129));
        _nw35[(_dafny.ONE).toNumber()] = ((_54_v4) ? (_215_v130) : (_dafny.MultiSet.fromElements(_214_v129, _214_v129, (_216_v131).dtor_cf88)));
        _nw35[(new BigNumber(2)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(3)).toNumber()] = (_215_v130).Intersect(_215_v130);
        _nw35[(new BigNumber(4)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(5)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.FromArray(_217_v132)).update(_214_v129, _module.__default.abs((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]));
        _nw35[(new BigNumber(7)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(8)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(9)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(10)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(11)).toNumber()] = (_215_v130).Union(_dafny.MultiSet.fromElements(_214_v129, _214_v129));
        _nw35[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_214_v129);
        _nw35[(new BigNumber(13)).toNumber()] = (_215_v130).Difference(_215_v130);
        _nw35[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.FromArray(_217_v132)).Union(_215_v130);
        _nw35[(new BigNumber(15)).toNumber()] = (((_223_v136).contains(_221_v134)) ? ((_223_v136).get(_221_v134)) : (_215_v130));
        _nw35[(new BigNumber(16)).toNumber()] = (_215_v130).Difference(_215_v130);
        _nw35[(new BigNumber(17)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(_214_v129))).Intersect((_215_v130).update(_214_v129, _module.__default.abs((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))])));
        _nw35[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(_214_v129);
        _nw35[(new BigNumber(19)).toNumber()] = (_215_v130).Intersect(_215_v130);
        _nw35[(new BigNumber(20)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(21)).toNumber()] = _215_v130;
        _nw35[(new BigNumber(22)).toNumber()] = _215_v130;
        _224_v137 = _nw35;
        let _index45 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_224_v137).length));
        (_224_v137)[_index45] = _dafny.MultiSet.FromArray(_217_v132);
        let _index46 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_224_v137).length));
        (_224_v137)[_index46] = _215_v130;
        let _225_v138;
        _225_v138 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_module.__default.fm11((_221_v134).fm1(true, new BigNumber(357), _49_v1, _53_globalState), _53_globalState), _48_v0),(_221_v134).fm6(_53_globalState));
        _225_v138 = (_225_v138).update((_60_v10).f10, new BigNumber(36));
      }
      (_53_globalState).f0 = _49_v1;
      let _index47 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length));
      (_105_v49)[_index47] = (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_56_v6).length))) {
        let _226_i11 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_226_i11)) && ((_226_i11).isLessThan(new BigNumber((_56_v6).length))))) {
          (_56_v6)[(_226_i11)] = (_60_v10).f10;
        }
      }
      let _227_i12;
      _227_i12 = _dafny.ZERO;
      L1: {
        while (_54_v4) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_227_i12)) {
              break L1;
            }
            _227_i12 = (_227_i12).plus(_dafny.ONE);
            let _228_v139;
            _228_v139 = _dafny.Seq.of(new BigNumber((_48_v0).length));
            let _229_v140;
            _229_v140 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Concat(_228_v139, _dafny.Seq.of((_60_v10).f9))).length), (_60_v10).f9, (_193_v108).fm4(_49_v1, _49_v1, _196_v111, (_60_v10).f9, _53_globalState), (_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]);
            let _230_v141;
            _230_v141 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), ((_231_v10) => function (_232_i13) {
              return (_231_v10).f9;
            })(_60_v10)));
            _229_v140 = _dafny.MultiSet.FromArray((_230_v141)[_module.__default.safeIndex((new BigNumber((_48_v0).length)).plus((_105_v49)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_105_v49).length))]), new BigNumber((_230_v141).length))]);
            _49_v1 = (_60_v10).f9;
            let _233_v142;
            _233_v142 = _dafny.Set.fromElements(_56_v6, _56_v6, _56_v6);
            _54_v4 = ((_233_v142).Intersect(_dafny.Set.fromElements(_56_v6, _56_v6))).IsDisjointFrom(_233_v142);
            let _234_v143;
            let _nw36 = Array((_dafny.ONE).toNumber()).fill([]);
            _234_v143 = _nw36;
            _234_v143 = _234_v143;
          }
        }
      }
      process.stdout.write((_48_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_49_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_50_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_53_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_53_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_53_globalState.f2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_54_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_55_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v7).equals(_dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_58_v8, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("rwdkj"), _dafny.Seq.UnicodeFromString("rwwkj")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_v9).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v10).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v10).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_61_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v46).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v47).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-434),new BigNumber(143)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_104_v48).f11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-434),new BigNumber(143)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v49)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(3), true, _dafny.Seq.of(true, true, true), new BigNumber(143), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_108_v51)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(143), true, _dafny.Seq.of(true), new BigNumber(959), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_113_v52, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v53).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v53).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_114_v53).dtor_cf36, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v53).dtor_cf37));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v53).dtor_cf38));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_115_v54).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(143),new BigNumber(143)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v55).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(3), true, _dafny.Seq.of(true, true, true), new BigNumber(143), new BigNumber(143)),_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_117_v57, _dafny.Seq.of(_module.D8.create_DC24(new BigNumber(3), true, _dafny.Seq.of(true, true, true), new BigNumber(143), new BigNumber(143)), _module.D8.create_DC24(new BigNumber(3), true, _dafny.Seq.of(true, true, true), new BigNumber(143), new BigNumber(143))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_118_v58, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC24(new BigNumber(3), true, _dafny.Seq.of(true, true, true), new BigNumber(143), new BigNumber(143)),_dafny.Seq.UnicodeFromString("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v60).equals(_dafny.Set.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_187_v106)[_dafny.ZERO]).dtor_cf24, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_187_v106)[_dafny.ONE]).dtor_cf24, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_187_v106)[new BigNumber(2)]).dtor_cf24, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_190_v107).dtor_cf24, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_194_v109).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v110).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v111).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v112).dtor_cf74));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v112).dtor_cf75));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v112).dtor_cf76));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_227_i12));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
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
    static create_DC3(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_module.D0.Default(), _dafny.ZERO);
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
    static create_DC5(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC7(cf6, cf7, cf8) {
      let $dt = new D2(2);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D2(3);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D2(4);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get is_DC8() { return this.$tag === 4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 4) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf9) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false);
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
    static create_DC10(cf11, cf12) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC11(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC9(cf10) {
      let $dt = new D3(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf10) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf10 === other.cf10;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, _dafny.ZERO);
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
    static create_DC13(cf15, cf16) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC12(cf14) {
      let $dt = new D4(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC14(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(false, _dafny.Map.Empty);
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
    static create_DC16(cf19, cf20, cf21, cf22, cf23) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC15(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + this.cf18.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(_module.D3.Default(), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, null);
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
    static create_DC18(cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC17(cf24) {
      let $dt = new D6(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(null, _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC20(cf29) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf30, cf31) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D7(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC22(cf32) {
      let $dt = new D7(3);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC22() { return this.$tag === 3; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(false);
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
    static create_DC24(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D8(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf39, cf40) {
      let $dt = new D8(1);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC23(cf33) {
      let $dt = new D8(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(_dafny.ZERO, false, _dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC27(cf42, cf43) {
      let $dt = new D9(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC26(cf41) {
      let $dt = new D9(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D9(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf41 === other.cf41;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC27(_dafny.Map.Empty, false);
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
    static create_DC30(cf46, cf47, cf48) {
      let $dt = new D10(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC31(cf49, cf50) {
      let $dt = new D10(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC29(cf45) {
      let $dt = new D10(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC30(false, false, false);
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
    static create_DC33(cf52, cf53) {
      let $dt = new D11(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC32(cf51) {
      let $dt = new D11(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && this.cf53 === other.cf53;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC33(false, false);
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
    static create_DC35(cf55) {
      let $dt = new D12(0);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC36(cf56, cf57) {
      let $dt = new D12(1);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC37(cf58) {
      let $dt = new D12(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC34(cf54) {
      let $dt = new D12(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf56, other.cf56) && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC35(false);
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
    static create_DC39(cf60, cf61, cf62, cf63, cf64) {
      let $dt = new D13(0);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC38(cf59) {
      let $dt = new D13(1);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62) && _dafny.areEqual(this.cf63, other.cf63) && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC39(_dafny.Map.Empty, [], _dafny.Seq.of(), _dafny.ZERO, _dafny.Seq.of());
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
    static create_DC40(cf65) {
      let $dt = new D14(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf65) + ")";
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
    static create_DC42(cf67) {
      let $dt = new D15(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC41(cf66) {
      let $dt = new D15(1);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC41" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC42(_dafny.MultiSet.Empty);
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
    static create_DC44(cf69, cf70, cf71) {
      let $dt = new D16(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC43(cf68) {
      let $dt = new D16(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC45(cf72) {
      let $dt = new D16(2);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC44(false, false, _dafny.ZERO);
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
    static create_DC47(cf74, cf75, cf76) {
      let $dt = new D17(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC46(cf73) {
      let $dt = new D17(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC47(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC49(cf78, cf79) {
      let $dt = new D18(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC48(cf77) {
      let $dt = new D18(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC50(cf80) {
      let $dt = new D18(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC49" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC48" + "(" + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC49(_dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC52(cf82) {
      let $dt = new D19(0);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC51(cf81) {
      let $dt = new D19(1);
      $dt.cf81 = cf81;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf81() { return this.cf81; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC52" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf81) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf81, other.cf81);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC52(_dafny.ZERO);
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
    static create_DC53(cf83) {
      let $dt = new D20(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC53" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf83, other.cf83);
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
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf85) {
      let $dt = new D21(0);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC56(cf86, cf87) {
      let $dt = new D21(1);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC54(cf84) {
      let $dt = new D21(2);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86) && _dafny.areEqual(this.cf87, other.cf87);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC55(false);
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
    static create_DC58() {
      let $dt = new D22(0);
      return $dt;
    }
    static create_DC57(cf88) {
      let $dt = new D22(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC58";
      } else if (this.$tag === 1) {
        return "D22.DC57" + "(" + _dafny.toString(this.cf88) + ")";
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
        return other.$tag === 1 && this.cf88 === other.cf88;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC58();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
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
      this.f0 = _dafny.ZERO;
      this.f2 = [];
      this._f1 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f11 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    fm10(p0, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("abyn"))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("tnbomydo")))).IsDisjointFrom((function () {
        let _coll20 = new _dafny.Set();
        for (const _compr_20 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bfuu"), _dafny.Seq.UnicodeFromString("qgnlm"))).Elements) {
          let _235_v0 = _compr_20;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bfuu"), _dafny.Seq.UnicodeFromString("qgnlm")), _235_v0)) {
            _coll20.add(_235_v0);
          }
        }
        return _coll20;
      }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hr"))));
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f6 = false;
      this._f5 = _dafny.MultiSet.Empty;
      this._f12 = false;
      this._f13 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f12, f13, f6, f5) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f6 = f6;
      (_this)._f5 = f5;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((_this).f13).length))).length))).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!((_this).f6), (_this).f12)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(323)), function (_236_i0) {
        return new BigNumber(579);
      })).length), new BigNumber(821), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_this).f12)).length), new BigNumber(562))).cardinality()), new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sejx"), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("sejx")).length)), new _dafny.CodePoint('i'.codePointAt(0)))).length)));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber(((((new BigNumber((_dafny.Set.fromElements((_this).f12, (_this).f6)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-463)))) ? ((_this).f5) : (((_this).f5).update((_this).f12, _module.__default.abs(new BigNumber(-762)))))).cardinality()));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm2(p0, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.of((_this).f6))).Union(_dafny.Set.fromElements(_dafny.Seq.of((_this).f6), _dafny.Seq.of((_this).f6, (_this).f6, (_this).f6)))).IsDisjointFrom(_dafny.Set.fromElements(_dafny.Seq.of((_this).f6), _dafny.Seq.of((_this).f12, true), _dafny.Seq.of(!((_this).f12)), _dafny.Seq.of((_this).f12)));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _237_v0;
      let _nw37 = Array((new BigNumber(6)).toNumber());
      _nw37[(_dafny.ZERO).toNumber()] = p2;
      _nw37[(_dafny.ONE).toNumber()] = p2;
      _nw37[(new BigNumber(2)).toNumber()] = (_this).f6;
      _nw37[(new BigNumber(3)).toNumber()] = p2;
      _nw37[(new BigNumber(4)).toNumber()] = (_this).f12;
      _nw37[(new BigNumber(5)).toNumber()] = (_this).f6;
      _237_v0 = _nw37;
      let _238_v1;
      _238_v1 = _dafny.MultiSet.fromElements(_237_v0);
      let _239_i0;
      _239_i0 = _dafny.ZERO;
      L2: {
        while (((_238_v1).Difference(_238_v1)).IsSubsetOf(_238_v1)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_239_i0)) {
              break L2;
            }
            _239_i0 = (_239_i0).plus(_dafny.ONE);
            let _240_v2;
            _240_v2 = _dafny.Set.fromElements(p0);
            let _241_v3;
            _241_v3 = _module.D2.create_DC4(_240_v2);
            let _242_v4;
            _242_v4 = _dafny.Seq.of(_241_v3);
            let _243_v5;
            _243_v5 = _dafny.Seq.of(_242_v4);
            let _244_v6;
            _244_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_242_v4);
            _243_v5 = _dafny.Seq.update(_dafny.Seq.of(_242_v4), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_242_v4)).length)), (((_244_v6).contains(p0)) ? ((_244_v6).get(p0)) : (_242_v4)));
            let _245_v7;
            let _nw38 = Array((new BigNumber(17)).toNumber()).fill(_module.D2.Default());
            _245_v7 = _nw38;
            let _246_v8;
            _246_v8 = new _dafny.CodePoint('a'.codePointAt(0));
            let _247_v9;
            _247_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-852));
            let _248_v10;
            _248_v10 = _module.D2.create_DC7(_246_v8, _247_v9, (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)));
            let _index48 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_245_v7).length));
            (_245_v7)[_index48] = _248_v10;
            let _index49 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_245_v7).length));
            (_245_v7)[_index49] = _248_v10;
            let _249_v11;
            _249_v11 = true;
            _249_v11 = (_this).f12;
            (globalState).f0 = p0;
          }
        }
      }
      let _250_v12;
      _250_v12 = new _dafny.CodePoint('n'.codePointAt(0));
      let _251_v13;
      _251_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f5).cardinality()),p0);
      let _252_v14;
      _252_v14 = _module.D2.create_DC7(_250_v12, _251_v13, p0);
      let _hi0 = p0;
      for (let _253_i1 = _module.__default.safeModuloInt(p0, (_252_v14).dtor_cf8); _253_i1.isLessThan(_hi0); _253_i1 = _253_i1.plus(_dafny.ONE)) {
        let _254_v15;
        _254_v15 = _dafny.Seq.UnicodeFromString("gmpwl");
        _254_v15 = _dafny.Seq.UnicodeFromString("ahtrhyr");
        let _255_v16;
        _255_v16 = _dafny.MultiSet.fromElements(p0, p0);
        (globalState).f0 = ((p1) ? (_253_i1) : (new BigNumber((_255_v16).cardinality())));
        let _256_v17;
        _256_v17 = _dafny.Seq.of(p2);
        let _257_v18;
        _257_v18 = _dafny.Set.fromElements(p1);
        _256_v17 = _dafny.Seq.of(((_this).f12) && ((_this).fm1(p2, p0, _253_i1, globalState)), (_this).f6, !(_257_v18).contains(p2));
        let _258_v19;
        _258_v19 = true;
        let _259_v20;
        _259_v20 = _dafny.Seq.of(_251_v13, (_252_v14).dtor_cf7);
        _258_v19 = !_dafny.Seq.contains(_259_v20, _251_v13);
      }
      (globalState).f0 = p0;
      let _260_v21;
      let _init6 = ((_261_p0) => function (_262_i3) {
        return (_262_i3).minus((_dafny.ZERO).minus(_261_p0));
      })(p0);
      let _nw39 = Array((new BigNumber(13)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw39.length); _i0_6++) {
        _nw39[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _260_v21 = _nw39;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_260_v21).length))) {
        let _263_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_263_i2)) && ((_263_i2).isLessThan(new BigNumber((_260_v21).length))))) {
          (_260_v21)[(_263_i2)] = _module.__default.safeModuloInt(_263_i2, p0);
        }
      }
      let _264_v22;
      _264_v22 = _module.D5.create_DC15((_this).f13);
      let _265_v23;
      _265_v23 = _dafny.MultiSet.fromElements(_250_v12, new _dafny.CodePoint('q'.codePointAt(0)));
      let _266_v24;
      _266_v24 = _dafny.Map.Empty.slice().updateUnsafe((_264_v22).dtor_cf18,_265_v23);
      _266_v24 = (_266_v24).update(_dafny.Seq.Concat((_this).f13, (_this).f13), _265_v23);
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_260_v21).length))) {
        let _267_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_267_i4)) && ((_267_i4).isLessThan(new BigNumber((_260_v21).length))))) {
          (_260_v21)[(_267_i4)] = (_267_i4).minus(p0);
        }
      }
      return;
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let _268_v0;
      _268_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f12)).length),(_dafny.ZERO).minus(p2));
      let _269_v1;
      _269_v1 = _dafny.Seq.of(p2);
      let _270_v2;
      _270_v2 = _dafny.Set.fromElements((_this).f6, !((_this).f12));
      let _271_v3;
      _271_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm14(_268_v0, (_this).f12, (_269_v1)[_module.__default.safeIndex(new BigNumber(((_this).f13).length), new BigNumber((_269_v1).length))], (_this).f12, globalState)).cardinality()),(_270_v2).IsSubsetOf(_270_v2));
      let _272_v4;
      _272_v4 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm2((_this).f12, globalState)),(_this).f12);
      _271_v3 = (_271_v3).update((_dafny.ZERO).minus((new BigNumber((_272_v4).length)).multipliedBy(p2)), (_this).f6);
      let _273_v5;
      _273_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(((_271_v3).contains(p2)) ? ((_271_v3).get(p2)) : ((_this).f6)));
      let _274_i0;
      _274_i0 = _dafny.ZERO;
      L3: {
        while ((p0).isLessThan((new BigNumber(((_this).f13).length)).minus((_this).fm4(new BigNumber(19), new BigNumber(((_this).f13).length), _273_v5, p0, globalState)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_274_i0)) {
              break L3;
            }
            _274_i0 = (_274_i0).plus(_dafny.ONE);
            let _275_v6;
            _275_v6 = _dafny.Map.Empty.slice().updateUnsafe(_268_v0,!((_this).f12));
            let _276_v7;
            let _nw40 = new _module.C0();
            _nw40.__ctor(_275_v6);
            _276_v7 = _nw40;
            let _277_v8;
            _277_v8 = _module.D2.create_DC7(new _dafny.CodePoint('m'.codePointAt(0)), _268_v0, p2);
            let _278_v9;
            _278_v9 = _dafny.MultiSet.fromElements(_277_v8);
            let _279_v10;
            _279_v10 = new _dafny.CodePoint('i'.codePointAt(0));
            let _pat_let_tv1 = _270_v2;
            let _pat_let_tv2 = p0;
            let _pat_let_tv3 = p1;
            let _pat_let_tv4 = p2;
            let _pat_let_tv5 = _268_v0;
            let _pat_let_tv6 = _268_v0;
            let _pat_let_tv7 = _270_v2;
            let _pat_let_tv8 = p0;
            let _pat_let_tv9 = p1;
            let _pat_let_tv10 = p2;
            let _pat_let_tv11 = _268_v0;
            let _pat_let_tv12 = _268_v0;
            (globalState).f0 = (((_278_v9).contains(function (_pat_let10_0) {
              return function (_294_dt__update__tmp_h2) {
                return function (_pat_let17_0) {
                  return function (_295_dt__update_hcf7_h2) {
                    return _module.D2.create_DC7((_294_dt__update__tmp_h2).dtor_cf6, _295_dt__update_hcf7_h2, (_294_dt__update__tmp_h2).dtor_cf8);
                  }(_pat_let17_0);
                }(_pat_let_tv12);
              }(_pat_let10_0);
            }(function (_pat_let11_0) {
              return function (_291_dt__update__tmp_h1) {
                return function (_pat_let15_0) {
                  return function (_292_dt__update_hcf8_h1) {
                    return function (_pat_let16_0) {
                      return function (_293_dt__update_hcf7_h1) {
                        return _module.D2.create_DC7((_291_dt__update__tmp_h1).dtor_cf6, _293_dt__update_hcf7_h1, _292_dt__update_hcf8_h1);
                      }(_pat_let16_0);
                    }(_pat_let_tv11);
                  }(_pat_let15_0);
                }(_pat_let_tv10);
              }(_pat_let11_0);
            }(function (_pat_let12_0) {
              return function (_288_dt__update__tmp_h0) {
                return function (_pat_let13_0) {
                  return function (_289_dt__update_hcf8_h0) {
                    return function (_pat_let14_0) {
                      return function (_290_dt__update_hcf7_h0) {
                        return _module.D2.create_DC7((_288_dt__update__tmp_h0).dtor_cf6, _290_dt__update_hcf7_h0, _289_dt__update_hcf8_h0);
                      }(_pat_let14_0);
                    }(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_pat_let_tv8),new BigNumber((_pat_let_tv9).cardinality())));
                  }(_pat_let13_0);
                }(new BigNumber((_pat_let_tv7).length));
              }(_pat_let12_0);
            }(_module.D2.create_DC7(_279_v10, _dafny.Map.Empty.slice().updateUnsafe(p2,p2), p2)))))) ? ((_278_v9).get(function (_pat_let2_0) {
              return function (_286_dt__update__tmp_h5) {
                return function (_pat_let9_0) {
                  return function (_287_dt__update_hcf7_h5) {
                    return _module.D2.create_DC7((_286_dt__update__tmp_h5).dtor_cf6, _287_dt__update_hcf7_h5, (_286_dt__update__tmp_h5).dtor_cf8);
                  }(_pat_let9_0);
                }(_pat_let_tv6);
              }(_pat_let2_0);
            }(function (_pat_let3_0) {
              return function (_283_dt__update__tmp_h4) {
                return function (_pat_let7_0) {
                  return function (_284_dt__update_hcf8_h3) {
                    return function (_pat_let8_0) {
                      return function (_285_dt__update_hcf7_h4) {
                        return _module.D2.create_DC7((_283_dt__update__tmp_h4).dtor_cf6, _285_dt__update_hcf7_h4, _284_dt__update_hcf8_h3);
                      }(_pat_let8_0);
                    }(_pat_let_tv5);
                  }(_pat_let7_0);
                }(_pat_let_tv4);
              }(_pat_let3_0);
            }(function (_pat_let4_0) {
              return function (_280_dt__update__tmp_h3) {
                return function (_pat_let5_0) {
                  return function (_281_dt__update_hcf8_h2) {
                    return function (_pat_let6_0) {
                      return function (_282_dt__update_hcf7_h3) {
                        return _module.D2.create_DC7((_280_dt__update__tmp_h3).dtor_cf6, _282_dt__update_hcf7_h3, _281_dt__update_hcf8_h2);
                      }(_pat_let6_0);
                    }(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_pat_let_tv2),new BigNumber((_pat_let_tv3).cardinality())));
                  }(_pat_let5_0);
                }(new BigNumber((_pat_let_tv1).length));
              }(_pat_let4_0);
            }(_module.D2.create_DC7(_279_v10, _dafny.Map.Empty.slice().updateUnsafe(p2,p2), p2)))))) : ((_dafny.ZERO).minus(p0)));
            let _296_v11;
            _296_v11 = false;
            _296_v11 = (p2).isLessThan(p0);
            let _297_v12;
            _297_v12 = _module.D3.create_DC10((_this).f6, p0);
            (globalState).f0 = (_297_v12).dtor_cf12;
          }
        }
      }
      (globalState).f0 = (_dafny.ZERO).minus(p0);
      let _298_v13;
      _298_v13 = _dafny.Seq.of(p1, _module.__default.fm15((_this).f6, globalState), (_dafny.MultiSet.FromArray(_269_v1)).Union(p1));
      let _299_v14;
      let _init7 = function (_300_i1) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-706)), function (_301_i2) {
          return (_this).f13;
        });
      };
      let _nw41 = Array((new BigNumber(20)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw41.length); _i0_7++) {
        _nw41[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _299_v14 = _nw41;
      let _302_v15;
      _302_v15 = _dafny.Seq.of((_this).f13);
      let _index50 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_299_v14).length));
      (_299_v14)[_index50] = _302_v15;
      let _303_v16;
      let _nw42 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _303_v16 = _nw42;
      let _304_v17;
      _304_v17 = _dafny.Seq.of((_this).f12, (_this).f12);
      let _index51 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_299_v14).length));
      let _rhs39 = _298_v13;
      let _rhs40 = p0;
      let _rhs41 = _303_v16;
      let _rhs42 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), function (_305_i3) {
        return (_this).f13;
      });
      let _rhs43 = ((((_this).fm2((_this).f6, globalState)) ? (p2) : (p0))).minus((((_this).f12) ? (new BigNumber((_304_v17).length)) : (p0)));
      let _lhs28 = globalState;
      let _lhs29 = globalState;
      let _lhs30 = _299_v14;
      let _lhs31 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_299_v14).length));
      let _lhs32 = globalState;
      _298_v13 = _rhs39;
      _lhs28.f0 = _rhs40;
      _lhs29.f2 = _rhs41;
      _lhs30[_lhs31] = _rhs42;
      _lhs32.f0 = _rhs43;
      let _306_v18;
      _306_v18 = false;
      let _307_v19;
      _307_v19 = _dafny.Seq.UnicodeFromString("gvyhp");
      let _rhs44 = ((_dafny.ZERO).minus((((_this).f12) ? (new BigNumber((_269_v1).length)) : (new BigNumber((p1).cardinality()))))).isLessThan(new BigNumber(60));
      let _rhs45 = _307_v19;
      _306_v18 = _rhs44;
      _307_v19 = _rhs45;
      (globalState).f0 = (_dafny.ZERO).minus(p2);
      return;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _308_v0;
      _308_v0 = new BigNumber(375);
      let _309_v1;
      _309_v1 = _dafny.Map.Empty.slice().updateUnsafe(_308_v0,new BigNumber(110));
      let _310_v2;
      _310_v2 = _dafny.Map.Empty.slice().updateUnsafe(_309_v1,(_this).f6);
      let _311_v3;
      let _nw43 = new _module.C0();
      _nw43.__ctor(_310_v2);
      _311_v3 = _nw43;
      let _312_v4;
      _312_v4 = _dafny.Seq.of(!((_this).f6), (_this).f6);
      let _313_v5;
      _313_v5 = _dafny.Map.Empty.slice().updateUnsafe(_311_v3,new BigNumber((_312_v4).length));
      _313_v5 = (_313_v5).update(_311_v3, _module.__default.safeModuloInt(_308_v0, _308_v0));
      let _314_v6;
      _314_v6 = _dafny.Seq.of(_308_v0);
      let _315_v7;
      _315_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_this).f12));
      let _316_v8;
      _316_v8 = _dafny.Map.Empty.slice().updateUnsafe((_314_v6)[_module.__default.safeIndex(_308_v0, new BigNumber((_314_v6).length))],(_315_v7).Merge(_315_v7));
      _316_v8 = (_316_v8).update((_308_v0).minus((_dafny.ZERO).minus(_308_v0)), _315_v7);
      r2 = !(p0) || ((_this).fm3(globalState));
      let _317_i0;
      _317_i0 = _dafny.ZERO;
      L4: {
        while ((_this).f6) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_317_i0)) {
              break L4;
            }
            _317_i0 = (_317_i0).plus(_dafny.ONE);
            _308_v0 = _308_v0;
            let _318_v9;
            let _init8 = function (_319_i1) {
              return (_this).f12;
            };
            let _nw44 = Array((new BigNumber(25)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw44.length); _i0_8++) {
              _nw44[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _318_v9 = _nw44;
            _318_v9 = _318_v9;
            _311_v3 = _311_v3;
            let _320_v10;
            _320_v10 = _dafny.Set.fromElements(_308_v0);
            let _321_v11;
            _321_v11 = _module.D6.create_DC18(_311_v3, _312_v4, new BigNumber((_320_v10).length));
            _315_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_dafny.MultiSet.FromArray(_312_v4)).equals(_dafny.MultiSet.FromArray((_321_v11).dtor_cf26)));
          }
        }
      }
      let _322_v12;
      let _nw45 = Array((new BigNumber(7)).toNumber());
      _nw45[(_dafny.ZERO).toNumber()] = _308_v0;
      _nw45[(_dafny.ONE).toNumber()] = _308_v0;
      _nw45[(new BigNumber(2)).toNumber()] = _308_v0;
      _nw45[(new BigNumber(3)).toNumber()] = new BigNumber(686);
      _nw45[(new BigNumber(4)).toNumber()] = (((_309_v1).contains(_308_v0)) ? ((_309_v1).get(_308_v0)) : (_308_v0));
      _nw45[(new BigNumber(5)).toNumber()] = _308_v0;
      _nw45[(new BigNumber(6)).toNumber()] = _308_v0;
      _322_v12 = _nw45;
      let _index52 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_322_v12).length));
      (_322_v12)[_index52] = new BigNumber((_dafny.Seq.update(_314_v6, _module.__default.safeIndex(_308_v0, new BigNumber((_314_v6).length)), _308_v0)).length);
      let _323_v13;
      _323_v13 = _dafny.Map.Empty.slice().updateUnsafe((_312_v4)[_module.__default.safeIndex(_308_v0, new BigNumber((_312_v4).length))],new BigNumber((_314_v6).length));
      let _324_v14;
      _324_v14 = _dafny.Set.fromElements(!((_this).f6), (_this).f6, (_this).f12, (_312_v4)[_module.__default.safeIndex(new BigNumber(((_323_v13).update((_this).f6, _308_v0)).length), new BigNumber((_312_v4).length))]);
      let _index53 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_322_v12).length));
      (_322_v12)[_index53] = ((p0) ? ((_308_v0).multipliedBy(new BigNumber((_324_v14).length))) : (_308_v0));
      let _325_i2;
      _325_i2 = _dafny.ZERO;
      L5: {
        while (p0) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_325_i2)) {
              break L5;
            }
            _325_i2 = (_325_i2).plus(_dafny.ONE);
            let _326_v15;
            let _nw46 = new _module.C0();
            _nw46.__ctor(((_311_v3).f11).Merge(((_311_v3).f11).update(_309_v1, (_this).f12)));
            _326_v15 = _nw46;
            r2 = (_this).f12;
            let _327_v16;
            let _nw47 = new _module.C0();
            _nw47.__ctor((_311_v3).f11);
            _327_v16 = _nw47;
            let _328_v17;
            let _nw48 = Array((new BigNumber(11)).toNumber()).fill(false);
            _328_v17 = _nw48;
            let _index54 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_328_v17).length));
            (_328_v17)[_index54] = (((_315_v7).contains(p0)) ? ((_315_v7).get(p0)) : ((_this).f12));
            let _index55 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_328_v17).length));
            (_328_v17)[_index55] = (_this).f12;
          }
        }
      }
      r0 = !((_this).f12) || ((_this).f12);
      r1 = (_308_v0).multipliedBy(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_308_v0), new BigNumber(585)));
      r2 = true;
      r3 = (_this).fm2(p0, globalState);
      return [r0, r1, r2, r3];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f5 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f5) {
      let _this = this;
      (_this)._f5 = f5;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return !(true) || (true);
    };
    fm2(p0, globalState) {
      let _this = this;
      return (new BigNumber(211)).isLessThan(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("uvl"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), function (_329_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("hxubpnhri"))).length));
    };
    m9(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _330_v0;
      _330_v0 = new _dafny.CodePoint('g'.codePointAt(0));
      let _331_v1;
      _331_v1 = _dafny.Seq.of(p0);
      let _332_v2;
      _332_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_331_v1).length));
      let _333_v3;
      _333_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_332_v2);
      _330_v0 = _module.__default.fm16(p0, (((_333_v3).contains(p0)) ? ((_333_v3).get(p0)) : (_332_v2)), _module.__default.fm17(globalState), globalState);
      let _334_v4;
      _334_v4 = _dafny.Seq.of(p1);
      _334_v4 = _module.__default.fm18(globalState);
      let _335_v5;
      _335_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _336_v6;
      _336_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
      let _337_v7;
      _337_v7 = _dafny.Seq.UnicodeFromString("ibtnyjf");
      let _338_v8;
      _338_v8 = _module.D7.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(p0,p1));
      let _339_v9;
      let _nw49 = Array((new BigNumber(23)).toNumber());
      _nw49[(_dafny.ZERO).toNumber()] = ((p0) ? (_335_v5) : (_dafny.Map.Empty.slice().updateUnsafe(true,p1)));
      _nw49[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      _nw49[(new BigNumber(2)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(3)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      _nw49[(new BigNumber(5)).toNumber()] = (_335_v5).Merge(_335_v5);
      _nw49[(new BigNumber(6)).toNumber()] = (_335_v5).Merge(_335_v5);
      _nw49[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
      _nw49[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_336_v6).contains(p0)) ? ((_336_v6).get(p0)) : (p0)),p1);
      _nw49[(new BigNumber(9)).toNumber()] = _module.__default.fm19((_module.D3.create_DC10(false, new BigNumber((_337_v7).length))).dtor_cf12, globalState);
      _nw49[(new BigNumber(10)).toNumber()] = (_335_v5).Merge(_335_v5);
      _nw49[(new BigNumber(11)).toNumber()] = ((p0) ? ((_335_v5).update(true, new BigNumber((_337_v7).length))) : (_335_v5));
      _nw49[(new BigNumber(12)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(13)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(14)).toNumber()] = ((p0) ? (_335_v5) : (_335_v5));
      _nw49[(new BigNumber(15)).toNumber()] = ((_338_v8).dtor_cf28).Merge(_335_v5);
      _nw49[(new BigNumber(16)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(17)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(18)).toNumber()] = ((!(true)) ? ((_335_v5).update((_this).fm2(false, globalState), p1)) : (_335_v5));
      _nw49[(new BigNumber(19)).toNumber()] = _335_v5;
      _nw49[(new BigNumber(20)).toNumber()] = _module.__default.fm19((_dafny.ZERO).minus(p1), globalState);
      _nw49[(new BigNumber(21)).toNumber()] = (_335_v5).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,p1));
      _nw49[(new BigNumber(22)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(false,p1)).Merge(_335_v5);
      _339_v9 = _nw49;
      let _340_v10;
      let _nw50 = Array((new BigNumber(18)).toNumber()).fill([]);
      _340_v10 = _nw50;
      let _341_v11;
      _341_v11 = true;
      let _342_v12;
      _342_v12 = _dafny.Set.fromElements(p1, new BigNumber((_336_v6).length));
      let _343_v13;
      _343_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_342_v12);
      let _344_v14;
      _344_v14 = _module.D5.create_DC15(_337_v7);
      let _345_v15;
      _345_v15 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), ((_346_v0) => function (_347_i0) {
        return _346_v0;
      })(_330_v0)), _337_v7);
      let _rhs46 = _339_v9;
      let _rhs47 = _340_v10;
      let _rhs48 = (((_336_v6).contains((_342_v12).IsDisjointFrom((((_343_v13).contains(p1)) ? ((_343_v13).get(p1)) : (_342_v12))))) ? ((_336_v6).get((_342_v12).IsDisjointFrom((((_343_v13).contains(p1)) ? ((_343_v13).get(p1)) : (_342_v12))))) : (_dafny.Seq.IsPrefixOf(_337_v7, (_344_v14).dtor_cf18)));
      let _rhs49 = ((_345_v15).IsProperSubsetOf(_345_v15)) && (false);
      _339_v9 = _rhs46;
      _340_v10 = _rhs47;
      _341_v11 = _rhs48;
      _341_v11 = _rhs49;
      let _348_v16;
      _348_v16 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), ((_349_v0) => function (_350_i1) {
        return _349_v0;
      })(_330_v0)));
      let _351_v17;
      let _nw51 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _351_v17 = _nw51;
      let _index56 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length));
      (_351_v17)[_index56] = p1;
      let _352_v19;
      _352_v19 = _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_341_v11),new BigNumber((_334_v4).length)));
      let _index57 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length));
      let _rhs50 = _dafny.Seq.Concat(_348_v16, _dafny.Seq.Concat(_348_v16, _348_v16));
      let _rhs51 = (p1).minus(p1);
      let _rhs52 = _module.__default.fm16(true, (function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of _dafny.IntegerRange(new BigNumber(366), new BigNumber(185))) {
          let _353_v18 = _compr_21;
          if (((new BigNumber(366)).isLessThanOrEqualTo(_353_v18)) && ((_353_v18).isLessThan(new BigNumber(185)))) {
            _coll21.push([(_353_v18).plus(new BigNumber((_dafny.MultiSet.fromElements(_330_v0)).cardinality())),p1]);
          }
        }
        return _coll21;
      }()).update(new BigNumber(((_352_v19).dtor_cf33).length), p1), _dafny.Seq.Concat(_337_v7, _337_v7), globalState);
      let _rhs53 = p1;
      let _rhs54 = p1;
      let _lhs33 = globalState;
      let _lhs34 = _351_v17;
      let _lhs35 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length));
      let _lhs36 = globalState;
      _348_v16 = _rhs50;
      _lhs33.f0 = _rhs51;
      _330_v0 = _rhs52;
      _lhs34[_lhs35] = _rhs53;
      _lhs36.f0 = _rhs54;
      let _354_v20;
      let _init9 = ((_355_v7) => function (_356_i2) {
        return _355_v7;
      })(_337_v7);
      let _nw52 = Array((new BigNumber(8)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw52.length); _i0_9++) {
        _nw52[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _354_v20 = _nw52;
      let _index58 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_354_v20).length));
      (_354_v20)[_index58] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(333)), ((_357_v0) => function (_358_i3) {
        return _357_v0;
      })(_330_v0));
      let _359_v21;
      _359_v21 = _module.D1.create_DC2(_341_v11);
      let _360_v22;
      _360_v22 = _dafny.Map.Empty.slice().updateUnsafe(_330_v0,_351_v17);
      let _index59 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length));
      let _index60 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_354_v20).length));
      let _rhs55 = p0;
      let _rhs56 = (_351_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length))];
      let _rhs57 = _dafny.Seq.update(_337_v7, _module.__default.safeIndex(new BigNumber((_360_v22).length), new BigNumber((_337_v7).length)), _module.__default.fm16(!(p0), _dafny.Map.Empty.slice().updateUnsafe((_351_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length))],(_351_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length))]), (_348_v16)[_module.__default.safeIndex(new BigNumber(-325), new BigNumber((_348_v16).length))], globalState));
      let _rhs58 = _359_v21;
      let _lhs37 = _351_v17;
      let _lhs38 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_351_v17).length));
      let _lhs39 = _354_v20;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_354_v20).length));
      _341_v11 = _rhs55;
      _lhs37[_lhs38] = _rhs56;
      _lhs39[_lhs40] = _rhs57;
      _359_v21 = _rhs58;
      (globalState).f0 = p1;
      let _361_v23;
      let _nw53 = Array((new BigNumber(19)).toNumber()).fill(false);
      _361_v23 = _nw53;
      let _362_v24;
      _362_v24 = _dafny.Map.Empty.slice().updateUnsafe(_361_v23,p0);
      let _363_v25;
      _363_v25 = _dafny.Seq.of(_362_v24, (_362_v24).Merge(_362_v24), _362_v24);
      r0 = (_363_v25)[_module.__default.safeIndex(new BigNumber((_342_v12).length), new BigNumber((_363_v25).length))];
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f6 = false;
      this._f5 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f6, f5) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f5 = f5;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f6), _dafny.Seq.of((_this).f6))).length)).multipliedBy((new BigNumber((_dafny.Seq.UnicodeFromString("rassq")).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cax")).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!((_this).f6))).length)))).length)));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      let _source3 = _module.D0.create_DC0(new BigNumber(-247));
      if (_source3.is_DC1) {
        return (_this).f6;
      } else {
        let _364___mcc_h0 = (_source3).cf0;
        let _365_cf0 = _364___mcc_h0;
        return !((_this).f6);
      }
    };
    fm2(p0, globalState) {
      let _this = this;
      return (_module.D1.create_DC2((_this).f6)).dtor_cf1;
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _366_v0;
      _366_v0 = _module.D0.create_DC1();
      _366_v0 = _366_v0;
      let _367_v1;
      _367_v1 = _dafny.Seq.UnicodeFromString("vekcmyumj");
      let _368_v2;
      _368_v2 = _module.D0.create_DC0(new BigNumber((_367_v1).length));
      let _pat_let_tv13 = p0;
      let _source4 = function (_pat_let18_0) {
        return function (_369_dt__update__tmp_h0) {
          return function (_pat_let19_0) {
            return function (_370_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_370_dt__update_hcf0_h0);
            }(_pat_let19_0);
          }(_pat_let_tv13);
        }(_pat_let18_0);
      }(_368_v2);
      if (_source4.is_DC1) {
        if (p1) {
          let _371_v3;
          _371_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _372_v4;
          _372_v4 = new _dafny.CodePoint('m'.codePointAt(0));
          let _373_v5;
          _373_v5 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-887)).isLessThanOrEqualTo(new BigNumber((_371_v3).length)),_372_v4);
          _373_v5 = (_373_v5).update(!(p0).isEqualTo(p0), _372_v4);
          let _374_v6;
          _374_v6 = _module.D1.create_DC2((_this).f6);
          let _375_v7;
          _375_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0);
          let _376_v8;
          _376_v8 = _dafny.Map.Empty.slice().updateUnsafe(_374_v6,_375_v7);
          let _377_v9;
          _377_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p0));
          (globalState).f0 = (_dafny.ZERO).minus((((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_376_v8).length),p0)).equals(_377_v9)) ? (p0) : (p0)));
          let _378_v10;
          _378_v10 = _dafny.Map.Empty.slice().updateUnsafe(_377_v9,p1);
          let _379_v11;
          let _nw54 = new _module.C0();
          _nw54.__ctor(_378_v10);
          _379_v11 = _nw54;
          let _380_v12;
          _380_v12 = _dafny.Map.Empty.slice().updateUnsafe(_379_v11,p0);
          _380_v12 = (_380_v12).update(_379_v11, p0);
          let _381_v13;
          let _nw55 = new _module.C0();
          _nw55.__ctor(((_379_v11).f11).Merge(_378_v10));
          _381_v13 = _nw55;
          (globalState).f0 = _module.__default.safeModuloInt(p0, p0);
        } else {
          let _382_v14;
          _382_v14 = true;
          _382_v14 = !(!((_this).f6));
          let _383_v15;
          let _nw56 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _383_v15 = _nw56;
          let _384_v16;
          _384_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _index61 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_383_v15).length));
          (_383_v15)[_index61] = (_384_v16).Merge(_384_v16);
          let _index62 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_383_v15).length));
          (_383_v15)[_index62] = _384_v16;
          let _385_v17;
          let _nw57 = Array((new BigNumber(2)).toNumber()).fill(false);
          _385_v17 = _nw57;
          let _index63 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_385_v17).length));
          (_385_v17)[_index63] = _382_v14;
          let _386_v18;
          let _nw58 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _386_v18 = _nw58;
          let _387_v19;
          _387_v19 = _dafny.Map.Empty.slice().updateUnsafe(_386_v18,p2);
          let _388_v20;
          _388_v20 = _dafny.Seq.of((_this).f6);
          let _index64 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_385_v17).length));
          let _rhs59 = (_this).f6;
          let _rhs60 = !(!(_387_v19).equals((_dafny.Map.Empty.slice().updateUnsafe(_386_v18,p1)).Merge(_387_v19)));
          let _rhs61 = !((_388_v20)[_module.__default.safeIndex(p0, new BigNumber((_388_v20).length))]);
          let _lhs41 = _385_v17;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_385_v17).length));
          _lhs41[_lhs42] = _rhs59;
          _382_v14 = _rhs60;
          _382_v14 = _rhs61;
          let _389_v21;
          _389_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,((_385_v17)[_module.__default.safeIndex(new BigNumber(682), new BigNumber((_385_v17).length))]) || ((_this).f6));
          _389_v21 = (_389_v21).update(false, (_385_v17)[_module.__default.safeIndex(new BigNumber(682), new BigNumber((_385_v17).length))]);
          (globalState).f0 = ((p0).minus(p0)).multipliedBy((p0).plus(new BigNumber((_388_v20).length)));
        }
        (globalState).f0 = p0;
        (globalState).f0 = (_dafny.ZERO).minus(new BigNumber(-971));
        let _390_v22;
        _390_v22 = true;
        let _391_v23;
        _391_v23 = _dafny.Set.fromElements(p0);
        let _392_v24;
        _392_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_366_v0);
        let _393_v26;
        _393_v26 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _rhs62 = !((_391_v23).IsDisjointFrom(function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_392_v24).Keys.Elements) {
            let _394_v25 = _compr_22;
            if ((_392_v24).contains(_394_v25)) {
              _coll22.add(_module.__default.safeDivisionInt(_394_v25, new BigNumber((_dafny.Seq.UnicodeFromString("i")).length)));
            }
          }
          return _coll22;
        }()));
        let _rhs63 = (_module.__default.safeModuloInt((((_393_v26).contains(false)) ? ((_393_v26).get(false)) : (p0)), new BigNumber(-562))).plus(p0);
        let _lhs43 = globalState;
        _390_v22 = _rhs62;
        _lhs43.f0 = _rhs63;
      } else {
        let _395___mcc_h0 = (_source4).cf0;
        let _396_cf0 = _395___mcc_h0;
        let _397_v27;
        _397_v27 = _dafny.MultiSet.fromElements(_396_cf0);
        _397_v27 = _397_v27;
        let _398_v28;
        _398_v28 = _module.D1.create_DC3(_366_v0, _module.__default.safeDivisionInt(new BigNumber(655), _396_cf0));
        let _399_v29;
        _399_v29 = true;
        let _400_v30;
        let _init10 = ((_401_v29) => function (_402_i0) {
          return _401_v29;
        })(_399_v29);
        let _nw59 = Array((new BigNumber(7)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw59.length); _i0_10++) {
          _nw59[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _400_v30 = _nw59;
        let _403_v31;
        _403_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ypkiarto"),(_this).f6);
        let _404_v32;
        _404_v32 = _dafny.Seq.of(p0, _396_cf0);
        let _rhs64 = _398_v28;
        let _rhs65 = (_this).fm4(_396_cf0, p0, (_403_v31).Merge(_403_v31), (_404_v32)[_module.__default.safeIndex((_dafny.ZERO).minus(_396_cf0), new BigNumber((_404_v32).length))], globalState);
        let _rhs66 = new BigNumber((_367_v1).length);
        let _rhs67 = p2;
        let _rhs68 = _400_v30;
        let _lhs44 = globalState;
        let _lhs45 = globalState;
        _398_v28 = _rhs64;
        _lhs44.f0 = _rhs65;
        _lhs45.f0 = _rhs66;
        _399_v29 = _rhs67;
        _400_v30 = _rhs68;
        let _405_v33;
        _405_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_404_v32, _404_v32)).length),_396_cf0);
        let _406_v34;
        _406_v34 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
        let _407_v35;
        _407_v35 = _dafny.Seq.of(_406_v34);
        let _408_v40;
        _408_v40 = _module.D2.create_DC4(_dafny.Set.fromElements(new BigNumber((_404_v32).length), new BigNumber(-705)));
        let _409_v42;
        let _nw60 = Array((new BigNumber(22)).toNumber());
        _nw60[(_dafny.ZERO).toNumber()] = _406_v34;
        _nw60[(_dafny.ONE).toNumber()] = (_406_v34).Difference(_406_v34);
        _nw60[(new BigNumber(2)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(3)).toNumber()] = (_407_v35)[_module.__default.safeIndex(_396_cf0, new BigNumber((_407_v35).length))];
        _nw60[(new BigNumber(4)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(5)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(p0);
        _nw60[(new BigNumber(7)).toNumber()] = function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of _dafny.IntegerRange(new BigNumber(753), new BigNumber(464))) {
            let _410_v36 = _compr_23;
            if (((new BigNumber(753)).isLessThanOrEqualTo(_410_v36)) && ((_410_v36).isLessThan(new BigNumber(464)))) {
              _coll23.add((_410_v36).plus(p0));
            }
          }
          return _coll23;
        }();
        _nw60[(new BigNumber(8)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(9)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(10)).toNumber()] = function () {
          let _coll24 = new _dafny.Set();
          for (const _compr_24 of (_397_v27).Elements) {
            let _411_v37 = _compr_24;
            if ((_397_v27).contains(_411_v37)) {
              _coll24.add(_module.__default.safeDivisionInt(_411_v37, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length),_dafny.Seq.of(new BigNumber(711), new BigNumber((function () {
                let _coll25 = new _dafny.Map();
                for (const _compr_25 of (_dafny.Set.fromElements(new BigNumber(113))).Elements) {
                  let _412_v38 = _compr_25;
                  if ((_dafny.Set.fromElements(new BigNumber(113))).contains(_412_v38)) {
                    _coll25.push([(_412_v38).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)),!(false)]);
                  }
                }
                return _coll25;
              }()).length)))).length),new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false))).length))).length)));
            }
          }
          return _coll24;
        }();
        _nw60[(new BigNumber(11)).toNumber()] = (_406_v34).Intersect(_406_v34);
        _nw60[(new BigNumber(12)).toNumber()] = function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of _dafny.IntegerRange(new BigNumber(636), new BigNumber(285))) {
            let _413_v39 = _compr_26;
            if (((new BigNumber(636)).isLessThanOrEqualTo(_413_v39)) && ((_413_v39).isLessThan(new BigNumber(285)))) {
              _coll26.add(_module.__default.safeDivisionInt(_413_v39, _396_cf0));
            }
          }
          return _coll26;
        }();
        _nw60[(new BigNumber(13)).toNumber()] = (_406_v34).Intersect(_dafny.Set.fromElements(p0));
        _nw60[(new BigNumber(14)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(p0);
        _nw60[(new BigNumber(16)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(17)).toNumber()] = (_408_v40).dtor_cf4;
        _nw60[(new BigNumber(18)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(19)).toNumber()] = function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-506), new BigNumber(339))) {
            let _414_v41 = _compr_27;
            if (((new BigNumber(-506)).isLessThanOrEqualTo(_414_v41)) && ((_414_v41).isLessThan(new BigNumber(339)))) {
              _coll27.add((_414_v41).multipliedBy(new BigNumber(25)));
            }
          }
          return _coll27;
        }();
        _nw60[(new BigNumber(20)).toNumber()] = _406_v34;
        _nw60[(new BigNumber(21)).toNumber()] = _406_v34;
        _409_v42 = _nw60;
        let _rhs69 = _405_v33;
        let _rhs70 = (new BigNumber(105)).multipliedBy(new BigNumber((((_this).f5).Union((_this).f5)).cardinality()));
        let _rhs71 = _409_v42;
        _405_v33 = _rhs69;
        _396_cf0 = _rhs70;
        _409_v42 = _rhs71;
        let _index65 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_400_v30).length));
        (_400_v30)[_index65] = !((_this).fm2(p2, globalState));
        let _415_v43;
        let _nw61 = Array((new BigNumber(25)).toNumber()).fill([]);
        _415_v43 = _nw61;
        let _416_v44;
        let _nw62 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _416_v44 = _nw62;
        let _417_v45;
        _417_v45 = _module.D3.create_DC9(_416_v44);
        let _418_v46;
        let _nw63 = Array((new BigNumber(11)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = _416_v44;
        _nw63[(_dafny.ONE).toNumber()] = _416_v44;
        _nw63[(new BigNumber(2)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(3)).toNumber()] = (_417_v45).dtor_cf10;
        _nw63[(new BigNumber(4)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(5)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(6)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(7)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(8)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(9)).toNumber()] = _416_v44;
        _nw63[(new BigNumber(10)).toNumber()] = _416_v44;
        _418_v46 = _nw63;
        let _index66 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_415_v43).length));
        (_415_v43)[_index66] = _418_v46;
        let _index67 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_416_v44).length));
        (_416_v44)[_index67] = p0;
        let _419_v47;
        let _nw64 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _419_v47 = _nw64;
        let _index68 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_419_v47).length));
        (_419_v47)[_index68] = _367_v1;
        let _index69 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_400_v30).length));
        let _index70 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_415_v43).length));
        let _index71 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_416_v44).length));
        let _index72 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_419_v47).length));
        let _rhs72 = _399_v29;
        let _rhs73 = _418_v46;
        let _rhs74 = (_396_cf0).multipliedBy(p0);
        let _rhs75 = _367_v1;
        let _rhs76 = (_dafny.ZERO).minus(p0);
        let _lhs46 = _400_v30;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_400_v30).length));
        let _lhs48 = _415_v43;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_415_v43).length));
        let _lhs50 = _416_v44;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_416_v44).length));
        let _lhs52 = _419_v47;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_419_v47).length));
        let _lhs54 = globalState;
        _lhs46[_lhs47] = _rhs72;
        _lhs48[_lhs49] = _rhs73;
        _lhs50[_lhs51] = _rhs74;
        _lhs52[_lhs53] = _rhs75;
        _lhs54.f0 = _rhs76;
      }
      let _420_v48;
      let _nw65 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _420_v48 = _nw65;
      let _421_v49;
      _421_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_367_v1, _module.__default.safeIndex(p0, new BigNumber((_367_v1).length)), new _dafny.CodePoint('r'.codePointAt(0))), _367_v1),_420_v48);
      _421_v49 = (_421_v49).update(_dafny.Seq.Concat(_367_v1, _367_v1), _420_v48);
      let _422_i1;
      _422_i1 = _dafny.ZERO;
      L6: {
        while (true) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_422_i1)) {
              break L6;
            }
            _422_i1 = (_422_i1).plus(_dafny.ONE);
            let _423_v50;
            _423_v50 = false;
            _423_v50 = !(((_this).f6) || (_423_v50));
            let _424_v51;
            _424_v51 = _dafny.Set.fromElements((_module.D2.create_DC5((_this).f6)).dtor_cf5);
            let _425_v52;
            _425_v52 = new _dafny.CodePoint('m'.codePointAt(0));
            (globalState).f0 = ((p1) ? ((new BigNumber((_424_v51).length)).minus(new BigNumber((_dafny.Set.fromElements(_425_v52, _425_v52, _425_v52)).length))) : (_module.__default.safeDivisionInt(p0, p0)));
            let _426_v53;
            _426_v53 = _dafny.Map.Empty.slice().updateUnsafe(_423_v50,(_this).fm3(globalState));
            let _427_v54;
            _427_v54 = _dafny.Set.fromElements(_425_v52);
            let _428_v55;
            _428_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_423_v50,(_this).f6),_427_v54);
            _423_v50 = !(_428_v55).contains(_426_v53);
            let _429_v56;
            _429_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _430_v57;
            _430_v57 = _dafny.Map.Empty.slice().updateUnsafe(_429_v56,!(p2));
            let _431_v58;
            let _nw66 = new _module.C0();
            _nw66.__ctor(_430_v57);
            _431_v58 = _nw66;
          }
        }
      }
      let _432_v59;
      let _nw67 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _432_v59 = _nw67;
      let _433_v60;
      _433_v60 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("f"));
      let _434_v61;
      _434_v61 = new _dafny.CodePoint('b'.codePointAt(0));
      let _435_v62;
      _435_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p0));
      let _436_v63;
      _436_v63 = _module.D2.create_DC7(_434_v61, _435_v62, p0);
      let _index73 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_432_v59).length));
      (_432_v59)[_index73] = (_433_v60)[_module.__default.safeIndex((_this).fm4((_436_v63).dtor_cf8, p0, _dafny.Map.Empty.slice().updateUnsafe(_367_v1,!((_this).f6)), p0, globalState), new BigNumber((_433_v60).length))];
      let _index74 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_432_v59).length));
      (_432_v59)[_index74] = _dafny.Seq.Concat(_module.__default.fm11((_this).f6, globalState), _367_v1);
      let _437_v64;
      _437_v64 = _dafny.Map.Empty.slice().updateUnsafe(_367_v1,p2);
      let _438_v65;
      _438_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f6);
      let _hi1 = (_this).fm4(p0, p0, _437_v64, new BigNumber((_438_v65).length), globalState);
      for (let _439_i2 = new BigNumber(((_module.__default.fm12(globalState)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(p0)))).cardinality()); _439_i2.isLessThan(_hi1); _439_i2 = _439_i2.plus(_dafny.ONE)) {
        let _440_v66;
        _440_v66 = false;
        let _441_v67;
        let _nw68 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _441_v67 = _nw68;
        let _442_v68;
        _442_v68 = _dafny.Seq.of(_441_v67);
        let _443_v69;
        let _nw69 = Array((new BigNumber(12)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = _441_v67;
        _nw69[(_dafny.ONE).toNumber()] = _441_v67;
        _nw69[(new BigNumber(2)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(3)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(4)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(5)).toNumber()] = (_442_v68)[_module.__default.safeIndex(_439_i2, new BigNumber((_442_v68).length))];
        _nw69[(new BigNumber(6)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(7)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(8)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(9)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(10)).toNumber()] = _441_v67;
        _nw69[(new BigNumber(11)).toNumber()] = _441_v67;
        _443_v69 = _nw69;
        let _444_v70;
        _444_v70 = _dafny.Seq.of((_this).f6);
        let _445_v71;
        _445_v71 = _dafny.Seq.of((_439_i2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_432_v59)[_module.__default.safeIndex(new BigNumber(599), new BigNumber((_432_v59).length))]).length))).length)), p0, (_this).fm4(new BigNumber(459), _439_i2, _module.__default.fm13(new BigNumber((_444_v70).length), _436_v63, p2, new BigNumber((_dafny.Set.fromElements(p1, (_this).f6, (_this).f6, p1)).length), globalState), new BigNumber(49), globalState));
        let _446_v72;
        _446_v72 = _dafny.Map.Empty.slice().updateUnsafe(_434_v61,_443_v69);
        let _447_v73;
        _447_v73 = _module.D4.create_DC12((((_446_v72).contains(_434_v61)) ? ((_446_v72).get(_434_v61)) : (_443_v69)));
        let _448_v74;
        _448_v74 = _dafny.Set.fromElements(p0, (_dafny.ZERO).minus(p0));
        let _449_v75;
        _449_v75 = _dafny.Seq.of(_445_v71, _445_v71);
        let _rhs77 = !(p1);
        let _rhs78 = (_447_v73).dtor_cf14;
        let _rhs79 = _439_i2;
        let _rhs80 = (_module.__default.safeModuloInt(new BigNumber((_448_v74).length), _439_i2)).multipliedBy(_module.__default.safeDivisionInt(_439_i2, _439_i2));
        let _rhs81 = _dafny.Seq.Concat((_449_v75)[_module.__default.safeIndex(_439_i2, new BigNumber((_449_v75).length))], _445_v71);
        let _lhs55 = globalState;
        let _lhs56 = globalState;
        _440_v66 = _rhs77;
        _443_v69 = _rhs78;
        _lhs55.f0 = _rhs79;
        _lhs56.f0 = _rhs80;
        _445_v71 = _rhs81;
        _434_v61 = _434_v61;
        _438_v65 = (_438_v65).update((_this).f6, _440_v66);
        let _450_v76;
        _450_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
        let _451_v77;
        let _nw70 = new _module.C1();
        _nw70.__ctor((_this).f6, _367_v1, true, _dafny.MultiSet.fromElements(false, _440_v66));
        _451_v77 = _nw70;
        let _452_v78;
        let _init11 = ((_453_v77) => function (_454_i3) {
          return !((_453_v77).f6);
        })(_451_v77);
        let _nw71 = Array((new BigNumber(18)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw71.length); _i0_11++) {
          _nw71[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _452_v78 = _nw71;
        let _index75 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_452_v78).length));
        (_452_v78)[_index75] = !((_this).f6);
        let _455_v80;
        _455_v80 = _dafny.MultiSet.fromElements(_368_v2);
        let _index76 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_452_v78).length));
        let _rhs82 = ((!((p0).isLessThanOrEqualTo(_439_i2))) ? (_450_v76) : ((((_451_v77).f6) ? (function () {
          let _coll28 = new _dafny.Map();
          for (const _compr_28 of _dafny.IntegerRange(new BigNumber(938), new BigNumber(-71))) {
            let _456_v79 = _compr_28;
            if (((new BigNumber(938)).isLessThanOrEqualTo(_456_v79)) && ((_456_v79).isLessThan(new BigNumber(-71)))) {
              _coll28.push([(_456_v79).multipliedBy(_439_i2),(_451_v77).f6]);
            }
          }
          return _coll28;
        }()) : (_450_v76))));
        let _rhs83 = _451_v77;
        let _rhs84 = _434_v61;
        let _rhs85 = ((_455_v80).update(_368_v2, _module.__default.abs(_439_i2))).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), ((_457_v2) => function (_458_i4) {
          return _457_v2;
        })(_368_v2))));
        let _lhs57 = _452_v78;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_452_v78).length));
        _450_v76 = _rhs82;
        _451_v77 = _rhs83;
        _434_v61 = _rhs84;
        _lhs57[_lhs58] = _rhs85;
      }
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = false;
      let _459_v0;
      let _nw72 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _459_v0 = _nw72;
      let _nw73 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _459_v0 = _nw73;
      r3 = p2;
      if (p2) {
        if ((_this).f6) {
          let _460_v1;
          let _nw74 = Array((new BigNumber(14)).toNumber()).fill(false);
          _460_v1 = _nw74;
          let _index77 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_460_v1).length));
          (_460_v1)[_index77] = (p0).isEqualTo(p0);
          let _461_v2;
          _461_v2 = _dafny.Seq.of((_this).fm1((_this).fm3(globalState), p0, p0, globalState));
          let _index78 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_460_v1).length));
          (_460_v1)[_index78] = (p2) && (!(p0).isEqualTo(new BigNumber((_461_v2).length)));
          (globalState).f0 = (_dafny.ZERO).minus(p0);
          let _462_v3;
          let _nw75 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
          _462_v3 = _nw75;
          let _463_v4;
          let _nw76 = new _module.C1();
          _nw76.__ctor(false, _dafny.Seq.UnicodeFromString("xy"), p2, _dafny.MultiSet.FromArray(_461_v2));
          _463_v4 = _nw76;
          let _464_v5;
          let _nw77 = Array((new BigNumber(29)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = _463_v4;
          _nw77[(_dafny.ONE).toNumber()] = _463_v4;
          _nw77[(new BigNumber(2)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(3)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(4)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(5)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(6)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(7)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(8)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(9)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(10)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(11)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(12)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(13)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(14)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(15)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(16)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(17)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(18)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(19)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(20)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(21)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(22)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(23)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(24)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(25)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(26)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(27)).toNumber()] = _463_v4;
          _nw77[(new BigNumber(28)).toNumber()] = _463_v4;
          _464_v5 = _nw77;
          let _465_v6;
          _465_v6 = _dafny.Map.Empty.slice().updateUnsafe(_460_v1,_464_v5);
          let _index79 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_462_v3).length));
          (_462_v3)[_index79] = _465_v6;
          let _index80 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_462_v3).length));
          (_462_v3)[_index80] = _465_v6;
          let _466_v7;
          _466_v7 = new _dafny.CodePoint('r'.codePointAt(0));
          let _467_v8;
          let _468_v9;
          let _469_v10;
          let _470_v11;
          let _out3;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector1 = (_463_v4).m8(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update((_463_v4).f13, _module.__default.safeIndex(p0, new BigNumber(((_463_v4).f13).length)), _466_v7), (_463_v4).f13), globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _out6 = _outcollector1[3];
          _467_v8 = _out3;
          _468_v9 = _out4;
          _469_v10 = _out5;
          _470_v11 = _out6;
          let _471_v12;
          _471_v12 = _dafny.Map.Empty.slice().updateUnsafe(_468_v9,(((_463_v4).f12) ? (false) : ((_460_v1)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_460_v1).length))])));
          _471_v12 = _471_v12;
        } else {
          (globalState).f0 = p0;
          let _472_v13;
          _472_v13 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _473_v14;
          _473_v14 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f6), _472_v13, _472_v13, _472_v13);
          let _474_v15;
          _474_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          let _475_v16;
          _475_v16 = _dafny.Seq.of(p2, (_this).f6);
          let _476_v17;
          _476_v17 = _dafny.Map.Empty.slice().updateUnsafe((_472_v13).equals((_473_v14)[_module.__default.safeIndex(p0, new BigNumber((_473_v14).length))]),(((_474_v15).contains((_this).f6)) ? ((_474_v15).get((_this).f6)) : (new BigNumber((_475_v16).length))));
          _476_v17 = _474_v15;
          r3 = p2;
          (globalState).f0 = p0;
          let _477_v18;
          let _nw78 = Array((new BigNumber(15)).toNumber());
          _477_v18 = _nw78;
          let _478_v19;
          let _nw79 = new _module.C2();
          _nw79.__ctor((_this).f5);
          _478_v19 = _nw79;
          let _index81 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_477_v18).length));
          (_477_v18)[_index81] = _478_v19;
          let _index82 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_477_v18).length));
          (_477_v18)[_index82] = ((false) ? (_478_v19) : (_478_v19));
        }
        if (!(p2)) {
          let _479_v20;
          _479_v20 = _dafny.Seq.of((_this).f6);
          let _480_v21;
          let _nw80 = new _module.C2();
          _nw80.__ctor(_dafny.MultiSet.FromArray(_479_v20));
          _480_v21 = _nw80;
          let _481_v22;
          let _out7;
          _out7 = (_480_v21).m9((p2) && ((_this).f6), p0, globalState);
          _481_v22 = _out7;
          let _482_v23;
          _482_v23 = _dafny.Seq.UnicodeFromString("edhl");
          let _483_v24;
          _483_v24 = new _dafny.CodePoint('q'.codePointAt(0));
          let _484_v25;
          _484_v25 = _dafny.Map.Empty.slice().updateUnsafe(_482_v23,(_this).f6);
          (globalState).f0 = (_this).fm4(_module.__default.safeDivisionInt(new BigNumber((_482_v23).length), new BigNumber(180)), (p0).minus(p0), (((_this).fm3(globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_482_v23, _module.__default.safeIndex(p0, new BigNumber((_482_v23).length)), _483_v24),false)) : (_484_v25)), ((((_this).f5).contains(p2)) ? (((_this).f5).get(p2)) : (p0)), globalState);
          let _485_v26;
          _485_v26 = _module.D1.create_DC2((_this).f6);
          _485_v26 = p3;
          let _486_v27;
          let _nw81 = new _module.C2();
          _nw81.__ctor(_dafny.MultiSet.fromElements(p2, (_this).fm3(globalState), (_this).fm2((_this).f6, globalState)));
          _486_v27 = _nw81;
        } else {
          let _487_v28;
          _487_v28 = _dafny.Seq.of((_this).f6);
          let _488_v29;
          _488_v29 = _dafny.Map.Empty.slice().updateUnsafe(_487_v28,p2);
          let _489_v30;
          _489_v30 = _dafny.Seq.UnicodeFromString("mflyhpk");
          let _490_v31;
          _490_v31 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm1(p2, p0, new BigNumber(46), globalState)),(_this).f6);
          let _491_v32;
          let _nw82 = new _module.C1();
          _nw82.__ctor((_this).f6, _489_v30, !(_490_v31).contains((_this).f6), (_this).f5);
          _491_v32 = _nw82;
          let _rhs86 = (_module.__default.fm20(p0, globalState)).update(_487_v28, p2);
          let _rhs87 = _491_v32;
          _488_v29 = _rhs86;
          _491_v32 = _rhs87;
          let _492_v33;
          let _init12 = function (_493_i0) {
            return (_this).f6;
          };
          let _nw83 = Array((new BigNumber(22)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw83.length); _i0_12++) {
            _nw83[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _492_v33 = _nw83;
          let _494_v34;
          _494_v34 = _dafny.Set.fromElements(_492_v33);
          _494_v34 = _494_v34;
          let _495_v35;
          _495_v35 = _dafny.Map.Empty.slice().updateUnsafe((_491_v32).f13,(p0).minus((_dafny.ZERO).minus(new BigNumber(((_491_v32).f13).length))));
          let _496_v37;
          _496_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _495_v35 = (_495_v35).update(_dafny.Seq.Concat((_491_v32).f13, _module.__default.fm11((_491_v32).f12, globalState)), (_dafny.ZERO).minus(((p2) ? (new BigNumber((_dafny.Set.fromElements(function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_496_v37).Keys.Elements) {
              let _497_v36 = _compr_29;
              if ((_496_v37).contains(_497_v36)) {
                _coll29.push([_module.__default.safeModuloInt(_497_v36, p0),(_491_v32).f12]);
              }
            }
            return _coll29;
          }())).length)) : ((_dafny.ZERO).minus(p0)))));
          let _498_v38;
          let _nw84 = Array((new BigNumber(7)).toNumber()).fill([]);
          _498_v38 = _nw84;
          let _499_v39;
          _499_v39 = _module.D3.create_DC10(!(false), new BigNumber(-542));
          let _500_v40;
          _500_v40 = _module.D0.create_DC0(p0);
          let _pat_let_tv14 = p0;
          let _pat_let_tv15 = p0;
          let _501_v41;
          let _nw85 = Array((new BigNumber(21)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _499_v39;
          _nw85[(_dafny.ONE).toNumber()] = _499_v39;
          _nw85[(new BigNumber(2)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(3)).toNumber()] = _module.D3.create_DC10(p2, new BigNumber(816));
          _nw85[(new BigNumber(4)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(5)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(6)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(7)).toNumber()] = function (_pat_let20_0) {
            return function (_502_dt__update__tmp_h0) {
              return function (_pat_let21_0) {
                return function (_503_dt__update_hcf12_h0) {
                  return _module.D3.create_DC10((_502_dt__update__tmp_h0).dtor_cf11, _503_dt__update_hcf12_h0);
                }(_pat_let21_0);
              }(_pat_let_tv14);
            }(_pat_let20_0);
          }(_module.__default.fm21((_491_v32).f12, globalState));
          _nw85[(new BigNumber(8)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(9)).toNumber()] = _module.D3.create_DC10((_491_v32).f12, p0);
          _nw85[(new BigNumber(10)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(11)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(12)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(13)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(14)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(15)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(16)).toNumber()] = _module.D3.create_DC10(p2, p0);
          _nw85[(new BigNumber(17)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(18)).toNumber()] = _499_v39;
          _nw85[(new BigNumber(19)).toNumber()] = function (_pat_let22_0) {
            return function (_504_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_505_dt__update_hcf12_h1) {
                  return _module.D3.create_DC10((_504_dt__update__tmp_h1).dtor_cf11, _505_dt__update_hcf12_h1);
                }(_pat_let23_0);
              }(_pat_let_tv15);
            }(_pat_let22_0);
          }(_499_v39);
          _nw85[(new BigNumber(20)).toNumber()] = _module.D3.create_DC10(p2, (_500_v40).dtor_cf0);
          _501_v41 = _nw85;
          let _index83 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_498_v38).length));
          (_498_v38)[_index83] = _501_v41;
          let _index84 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_498_v38).length));
          (_498_v38)[_index84] = ((!(false)) ? (_501_v41) : (_501_v41));
          _492_v33 = _492_v33;
        }
        let _506_v42;
        _506_v42 = _dafny.Seq.of((_this).f6);
        let _507_v43;
        _507_v43 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(p2, (_this).f6)).length), new BigNumber((_dafny.MultiSet.FromArray(_506_v42)).cardinality()), p0, p0, p0);
        let _508_v44;
        _508_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        let _509_v45;
        _509_v45 = _dafny.Set.fromElements(_459_v0);
        let _510_v46;
        _510_v46 = new _dafny.CodePoint('n'.codePointAt(0));
        let _511_v47;
        _511_v47 = _dafny.Seq.of(_510_v46, _510_v46);
        let _512_v48;
        _512_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_511_v47).length),(_this).f6);
        let _513_v49;
        _513_v49 = _dafny.Seq.of((_dafny.ZERO).minus(p0), p0);
        let _514_v50;
        let _nw86 = Array((new BigNumber(22)).toNumber());
        _nw86[(_dafny.ZERO).toNumber()] = !(true);
        _nw86[(_dafny.ONE).toNumber()] = p2;
        _nw86[(new BigNumber(2)).toNumber()] = !((_this).f6) || ((_this).f6);
        _nw86[(new BigNumber(3)).toNumber()] = ((_this).f6) === (p2);
        _nw86[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw86[(new BigNumber(5)).toNumber()] = (_this).f6;
        _nw86[(new BigNumber(6)).toNumber()] = p2;
        _nw86[(new BigNumber(7)).toNumber()] = p2;
        _nw86[(new BigNumber(8)).toNumber()] = (_507_v43).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(-874)));
        _nw86[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw86[(new BigNumber(10)).toNumber()] = p2;
        _nw86[(new BigNumber(11)).toNumber()] = p2;
        _nw86[(new BigNumber(12)).toNumber()] = p2;
        _nw86[(new BigNumber(13)).toNumber()] = p2;
        _nw86[(new BigNumber(14)).toNumber()] = (_module.__default.fm21(!((_this).f6), globalState)).dtor_cf11;
        _nw86[(new BigNumber(15)).toNumber()] = true;
        _nw86[(new BigNumber(16)).toNumber()] = (((_508_v44).contains(!(false))) ? ((_508_v44).get(!(false))) : ((_this).f6));
        _nw86[(new BigNumber(17)).toNumber()] = (_dafny.Set.fromElements(_459_v0, _459_v0)).IsSubsetOf(_509_v45);
        _nw86[(new BigNumber(18)).toNumber()] = (_this).fm3(globalState);
        _nw86[(new BigNumber(19)).toNumber()] = (new BigNumber((_512_v48).length)).isLessThanOrEqualTo((_513_v49)[_module.__default.safeIndex(p0, new BigNumber((_513_v49).length))]);
        _nw86[(new BigNumber(20)).toNumber()] = p2;
        _nw86[(new BigNumber(21)).toNumber()] = p2;
        _514_v50 = _nw86;
        let _index85 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length));
        (_514_v50)[_index85] = p2;
        let _index86 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length));
        (_514_v50)[_index86] = false;
        let _515_v51;
        let _init13 = ((_516_v47, _517_v50, _518_p0) => function (_519_i1) {
          return (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_516_v47).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_517_v50)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_517_v50).length))],_518_p0));
        })(_511_v47, _514_v50, p0);
        let _nw87 = Array((new BigNumber(28)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw87.length); _i0_13++) {
          _nw87[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _515_v51 = _nw87;
        let _520_v52;
        _520_v52 = _dafny.Map.Empty.slice().updateUnsafe(!((_514_v50)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length))]),p0);
        let _index87 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_515_v51).length));
        (_515_v51)[_index87] = _520_v52;
        let _index88 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_515_v51).length));
        let _rhs88 = (((_this).f6) ? (_507_v43) : ((_507_v43).Difference(function () {
          let _coll30 = new _dafny.Set();
          for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-695), new BigNumber(274))) {
            let _521_v53 = _compr_30;
            if (((new BigNumber(-695)).isLessThanOrEqualTo(_521_v53)) && ((_521_v53).isLessThan(new BigNumber(274)))) {
              _coll30.add((_521_v53).minus(p0));
            }
          }
          return _coll30;
        }())));
        let _rhs89 = (_this).fm3(globalState);
        let _rhs90 = (((p0).isLessThan(p0)) ? ((_514_v50)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length))]) : (p2));
        let _rhs91 = ((_520_v52).Merge(_module.__default.fm19(p0, globalState))).update(((_514_v50)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length))]) && ((_this).f6), (new BigNumber(438)).multipliedBy(new BigNumber((((_this).f5).update((_514_v50)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length))], _module.__default.abs(p0))).cardinality())));
        let _lhs59 = _515_v51;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_515_v51).length));
        _507_v43 = _rhs88;
        r3 = _rhs89;
        r3 = _rhs90;
        _lhs59[_lhs60] = _rhs91;
        let _index89 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_514_v50).length));
        (_514_v50)[_index89] = (_this).fm1(false, p0, _module.__default.safeDivisionInt(p0, p0), globalState);
      } else {
        let _522_v54;
        let _nw88 = Array((new BigNumber(8)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = !(!(p2));
        _nw88[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(824)), function (_523_i2) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("yaxq"));
        _nw88[(new BigNumber(2)).toNumber()] = (_this).f6;
        _nw88[(new BigNumber(3)).toNumber()] = !(p0).isEqualTo(p0);
        _nw88[(new BigNumber(4)).toNumber()] = p2;
        _nw88[(new BigNumber(5)).toNumber()] = !(p0).isEqualTo(p0);
        _nw88[(new BigNumber(6)).toNumber()] = p2;
        _nw88[(new BigNumber(7)).toNumber()] = (_this).f6;
        _522_v54 = _nw88;
        let _index90 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_522_v54).length));
        (_522_v54)[_index90] = (_this).f6;
        let _524_v55;
        _524_v55 = _dafny.Seq.of(p0, p0, new BigNumber(-775), p0);
        let _525_v56;
        _525_v56 = _dafny.Seq.UnicodeFromString("jrojrlvjv");
        let _526_v57;
        _526_v57 = _dafny.Map.Empty.slice().updateUnsafe(_525_v56,(_this).f6);
        let _527_v58;
        _527_v58 = _dafny.Set.fromElements(p2, (_this).f6);
        let _index91 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_522_v54).length));
        let _rhs92 = (_this).fm4(_module.__default.safeModuloInt((_524_v55)[_module.__default.safeIndex(p0, new BigNumber((_524_v55).length))], (_dafny.ZERO).minus(p0)), new BigNumber(512), (_526_v57).Merge(_526_v57), new BigNumber(108), globalState);
        let _rhs93 = p2;
        let _rhs94 = !((_527_v58).IsDisjointFrom(_527_v58));
        let _lhs61 = globalState;
        let _lhs62 = _522_v54;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_522_v54).length));
        _lhs61.f0 = _rhs92;
        _lhs62[_lhs63] = _rhs93;
        r3 = _rhs94;
        (globalState).f0 = p0;
        let _rhs95 = p0;
        let _rhs96 = ((new BigNumber((_527_v58).length)).plus(p0)).multipliedBy(p0);
        let _lhs64 = globalState;
        let _lhs65 = globalState;
        _lhs64.f0 = _rhs95;
        _lhs65.f0 = _rhs96;
        let _528_v59;
        _528_v59 = new _dafny.CodePoint('w'.codePointAt(0));
        _528_v59 = _528_v59;
        let _529_v60;
        _529_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p0, p0, p0, new BigNumber(966)),_dafny.Seq.of((_522_v54)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_522_v54).length))], true));
        let _530_v61;
        _530_v61 = _dafny.Map.Empty.slice().updateUnsafe(_529_v60,((_522_v54)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_522_v54).length))]) && (p2));
        _530_v61 = (_530_v61).update(_529_v60, !(!(!((_this).f6)) || (false)));
      }
      if ((_this).f6) {
        let _531_v62;
        _531_v62 = _dafny.MultiSet.fromElements((_this).f6);
        _531_v62 = ((_this).f5).Difference((_this).f5);
        let _532_v63;
        _532_v63 = new _dafny.CodePoint('t'.codePointAt(0));
        let _533_v64;
        _533_v64 = _module.D2.create_DC7(_532_v63, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), p0);
        let _534_v65;
        _534_v65 = _module.D2.create_DC8(_533_v64);
        let _535_v66;
        _535_v66 = _dafny.Map.Empty.slice().updateUnsafe(_534_v65,p0);
        let _536_v67;
        _536_v67 = _dafny.Map.Empty.slice().updateUnsafe((_535_v66).update(_534_v65, p0),p0);
        (globalState).f0 = (((_536_v67).contains(_535_v66)) ? ((_536_v67).get(_535_v66)) : (p0));
        r2 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ysdg"), _module.__default.fm11(p2, globalState));
        r3 = (_this).f6;
        r3 = (_this).f6;
      } else {
        let _537_v68;
        _537_v68 = _module.D3.create_DC10(p2, p0);
        let _538_v69;
        _538_v69 = _dafny.Seq.of((_537_v68).dtor_cf12);
        _538_v69 = _538_v69;
        let _539_v70;
        let _nw89 = new _module.C2();
        _nw89.__ctor((_this).f5);
        _539_v70 = _nw89;
        r3 = (_this).f6;
        let _540_v71;
        _540_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(622),new BigNumber(((_this).f5).cardinality()));
        r3 = ((_dafny.MultiSet.fromElements((_539_v70).fm1((_this).f6, new BigNumber((_540_v71).length), p0, globalState), p2)).IsSubsetOf((_this).f5)) === (!((_this).f6));
        let _index92 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_459_v0).length));
        (_459_v0)[_index92] = p0;
        let _541_v72;
        _541_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p2);
        let _542_v73;
        _542_v73 = _dafny.Seq.UnicodeFromString("qgee");
        let _543_v74;
        _543_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(157),_542_v73);
        let _index93 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_459_v0).length));
        let _rhs97 = (_module.__default.safeModuloInt(p0, new BigNumber((_541_v72).length))).plus(p0);
        let _rhs98 = new BigNumber(((((_543_v74).contains((((_this).f6) ? (p0) : (p0)))) ? ((_543_v74).get((((_this).f6) ? (p0) : (p0)))) : (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(666)), function (_544_i3) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _542_v73)))).length);
        let _rhs99 = (_this).fm2(p2, globalState);
        let _lhs66 = globalState;
        let _lhs67 = _459_v0;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_459_v0).length));
        _lhs66.f0 = _rhs97;
        _lhs67[_lhs68] = _rhs98;
        r3 = _rhs99;
      }
      r3 = (_this).f6;
      let _index94 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_459_v0).length));
      (_459_v0)[_index94] = p0;
      let _545_v75;
      _545_v75 = _dafny.Seq.UnicodeFromString("frah");
      let _546_v76;
      _546_v76 = _dafny.Map.Empty.slice().updateUnsafe(_545_v75,!(p2));
      let _547_v77;
      _547_v77 = new _dafny.CodePoint('t'.codePointAt(0));
      let _index95 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_459_v0).length));
      let _rhs100 = (_this).fm4(p0, (_dafny.ZERO).minus(p0), _546_v76, p0, globalState);
      let _rhs101 = ((((_this).f5).contains((_this).f6)) ? (((_this).f5).get((_this).f6)) : (new BigNumber((_dafny.Seq.update(_545_v75, _module.__default.safeIndex(new BigNumber((_545_v75).length), new BigNumber((_545_v75).length)), _547_v77)).length)));
      let _lhs69 = globalState;
      let _lhs70 = _459_v0;
      let _lhs71 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_459_v0).length));
      _lhs69.f0 = _rhs100;
      _lhs70[_lhs71] = _rhs101;
      r0 = _module.__default.fm22(globalState);
      r1 = _dafny.Seq.Concat(_545_v75, _545_v75);
      r2 = _dafny.Seq.Concat(_545_v75, _545_v75);
      r3 = (_this).f6;
      return [r0, r1, r2, r3];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f7 = [];
      this._f8 = false;
      this._f6 = false;
      this._f5 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    set f8(value) {
      let _this = this;
      _this._f8 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f7, f8, f6, f5) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f6 = f6;
      (_this)._f5 = f5;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(_this.f8, (_this).f6, _this.f8, _this.f8, _this.f8)), _dafny.Seq.of(_this.f8, (_this).f6, (_this).f6, (_this).f6, _this.f8));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(-256);
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f8;
    };
    fm2(p0, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm6(globalState) {
      let _this = this;
      return (new BigNumber(32)).minus((_dafny.ZERO).minus(((((_this).f5).contains(_this.f8)) ? (((_this).f5).get(_this.f8)) : (new BigNumber(-44)))));
    };
    fm7(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lsi"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jxsfieg"), _dafny.Seq.UnicodeFromString("isrelfj")));
    };
    m2(p0, globalState) {
      let _this = this;
      let _548_v0;
      _548_v0 = new _dafny.CodePoint('o'.codePointAt(0));
      let _rhs102 = false;
      let _rhs103 = (_this).fm6(globalState);
      let _rhs104 = _548_v0;
      let _lhs72 = _this;
      let _lhs73 = globalState;
      _lhs72.f8 = _rhs102;
      _lhs73.f0 = _rhs103;
      _548_v0 = _rhs104;
      let _549_v1;
      _549_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f8);
      _549_v1 = (_549_v1).update(((_this).f6) && ((_this).f6), (_this).f6);
      let _550_v2;
      _550_v2 = new BigNumber(725);
      let _551_v3;
      _551_v3 = _module.D0.create_DC1();
      let _552_v4;
      _552_v4 = _dafny.Seq.UnicodeFromString("elhkbkopj");
      let _553_v5;
      _553_v5 = _dafny.Seq.of(_552_v4, _552_v4);
      let _554_v6;
      let _nw90 = Array((new BigNumber(18)).toNumber());
      _nw90[(_dafny.ZERO).toNumber()] = _550_v2;
      _nw90[(_dafny.ONE).toNumber()] = new BigNumber((_552_v4).length);
      _nw90[(new BigNumber(2)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(3)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(4)).toNumber()] = new BigNumber(365);
      _nw90[(new BigNumber(5)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(6)).toNumber()] = (_this).fm6(globalState);
      _nw90[(new BigNumber(7)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(8)).toNumber()] = new BigNumber((_553_v5).length);
      _nw90[(new BigNumber(9)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(10)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(11)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(12)).toNumber()] = new BigNumber((_552_v4).length);
      _nw90[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.of(_550_v2, _550_v2)).length);
      _nw90[(new BigNumber(14)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("bjayaggc")).length);
      _nw90[(new BigNumber(16)).toNumber()] = _550_v2;
      _nw90[(new BigNumber(17)).toNumber()] = new BigNumber((_module.__default.fm8(_550_v2, false, _552_v4, globalState)).length);
      _554_v6 = _nw90;
      let _555_v7;
      _555_v7 = _dafny.Map.Empty.slice().updateUnsafe(_551_v3,_554_v6);
      let _556_v8;
      _556_v8 = _dafny.Map.Empty.slice().updateUnsafe(_550_v2,new BigNumber((_555_v7).length));
      let _557_v9;
      _557_v9 = _dafny.Set.fromElements((_this).fm4(_550_v2, _550_v2, _dafny.Map.Empty.slice().updateUnsafe(_552_v4,_this.f8), _550_v2, globalState), _550_v2);
      let _558_v10;
      let _nw91 = Array((new BigNumber(27)).toNumber());
      _nw91[(_dafny.ZERO).toNumber()] = _550_v2;
      _nw91[(_dafny.ONE).toNumber()] = (_550_v2).multipliedBy((_dafny.ZERO).minus(_550_v2));
      _nw91[(new BigNumber(2)).toNumber()] = new BigNumber((_556_v8).length);
      _nw91[(new BigNumber(3)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(4)).toNumber()] = new BigNumber(737);
      _nw91[(new BigNumber(5)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(6)).toNumber()] = new BigNumber(115);
      _nw91[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_552_v4).length), _550_v2);
      _nw91[(new BigNumber(8)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(9)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_557_v9).length),_550_v2)).length);
      _nw91[(new BigNumber(11)).toNumber()] = new BigNumber(-640);
      _nw91[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_550_v2).plus(_550_v2));
      _nw91[(new BigNumber(13)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(14)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(_550_v2, _550_v2);
      _nw91[(new BigNumber(16)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(17)).toNumber()] = new BigNumber(((_this).f5).cardinality());
      _nw91[(new BigNumber(18)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f8,_550_v2)).length);
      _nw91[(new BigNumber(20)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(21)).toNumber()] = new BigNumber((_module.__default.fm9(globalState)).length);
      _nw91[(new BigNumber(22)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(23)).toNumber()] = _550_v2;
      _nw91[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(_550_v2);
      _nw91[(new BigNumber(25)).toNumber()] = new BigNumber(518);
      _nw91[(new BigNumber(26)).toNumber()] = _550_v2;
      _558_v10 = _nw91;
      _558_v10 = _554_v6;
      let _559_i0;
      _559_i0 = _dafny.ZERO;
      L7: {
        while (!(_550_v2).isEqualTo((new BigNumber((_552_v4).length)).minus((_dafny.ZERO).minus(_550_v2)))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_559_i0)) {
              break L7;
            }
            _559_i0 = (_559_i0).plus(_dafny.ONE);
            if ((_this).f6) {
              let _560_v12;
              _560_v12 = _dafny.Map.Empty.slice().updateUnsafe(_551_v3,p0);
              let _561_v13;
              _561_v13 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_550_v2,new BigNumber((_560_v12).length)), _556_v8);
              let _562_v15;
              _562_v15 = _dafny.Seq.of(_556_v8, _556_v8);
              let _563_v16;
              let _nw92 = new _module.C0();
              _nw92.__ctor(((!(!(_this.f8))) ? (function () {
                let _coll31 = new _dafny.Map();
                for (const _compr_31 of (_561_v13).Elements) {
                  let _564_v11 = _compr_31;
                  if ((_561_v13).contains(_564_v11)) {
                    _coll31.push([_564_v11,false]);
                  }
                }
                return _coll31;
              }()) : (function () {
                let _coll32 = new _dafny.Map();
                for (const _compr_32 of (_562_v15).Elements) {
                  let _565_v14 = _compr_32;
                  if (_dafny.Seq.contains(_562_v15, _565_v14)) {
                    _coll32.push([_565_v14,_this.f8]);
                  }
                }
                return _coll32;
              }())));
              _563_v16 = _nw92;
              let _566_v17;
              let _nw93 = new _module.C3();
              _nw93.__ctor((_this).f6, _dafny.MultiSet.fromElements((_this).f6));
              _566_v17 = _nw93;
              let _567_v18;
              _567_v18 = _dafny.Set.fromElements(_566_v17, _566_v17);
              let _arr0 = _this.f7;
              let _index96 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_this.f7).length));
              _arr0[_index96] = _this.f8;
              let _568_v19;
              _568_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(913),!(((false) ? ((_566_v17).f6) : ((_this).f6))));
              let _569_v20;
              _569_v20 = _dafny.Seq.of(_550_v2);
              let _arr1 = _this.f7;
              let _index97 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_this.f7).length));
              let _rhs105 = _dafny.Set.fromElements(_566_v17, _566_v17, _566_v17);
              let _rhs106 = (((_568_v19).contains(new BigNumber((_dafny.Seq.Concat(_552_v4, _dafny.Seq.update(_552_v4, _module.__default.safeIndex(_550_v2, new BigNumber((_552_v4).length)), _548_v0))).length))) ? ((_568_v19).get(new BigNumber((_dafny.Seq.Concat(_552_v4, _dafny.Seq.update(_552_v4, _module.__default.safeIndex(_550_v2, new BigNumber((_552_v4).length)), _548_v0))).length))) : ((_550_v2).isLessThanOrEqualTo(new BigNumber((_569_v20).length))));
              let _rhs107 = (_566_v17).f6;
              let _rhs108 = _551_v3;
              let _lhs74 = _this;
              let _lhs75 = _this.f7;
              let _lhs76 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_this.f7).length));
              _567_v18 = _rhs105;
              _lhs74.f8 = _rhs106;
              _lhs75[_lhs76] = _rhs107;
              _551_v3 = _rhs108;
              let _arr2 = _this.f7;
              let _index98 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_this.f7).length));
              _arr2[_index98] = (_566_v17).f6;
              _550_v2 = new BigNumber((_dafny.Seq.UnicodeFromString("ulmddp")).length);
              let _570_v21;
              _570_v21 = _dafny.Map.Empty.slice().updateUnsafe(_552_v4,_this.f8);
              (globalState).f0 = (_this).fm4(_550_v2, _550_v2, _570_v21, _550_v2, globalState);
            } else {
              let _571_v22;
              _571_v22 = _dafny.MultiSet.fromElements(new BigNumber((_552_v4).length));
              (_this).f8 = (_571_v22).IsSubsetOf(_571_v22);
              let _arr3 = _this.f7;
              let _index99 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_this.f7).length));
              _arr3[_index99] = (_this).f6;
              let _572_v23;
              _572_v23 = _dafny.Seq.of(new BigNumber((_552_v4).length));
              let _arr4 = _this.f7;
              let _index100 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_this.f7).length));
              let _rhs109 = (_550_v2).minus(_550_v2);
              let _rhs110 = new _dafny.CodePoint('g'.codePointAt(0));
              let _rhs111 = (_module.__default.safeModuloInt(new BigNumber((_572_v23).length), new BigNumber(((_this).f5).cardinality()))).isEqualTo(_550_v2);
              let _rhs112 = _this.f7;
              let _rhs113 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm19((_dafny.ZERO).minus(_550_v2), globalState)).length));
              let _lhs77 = globalState;
              let _lhs78 = _this.f7;
              let _lhs79 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_this.f7).length));
              let _lhs80 = _this;
              _lhs77.f0 = _rhs109;
              _548_v0 = _rhs110;
              _lhs78[_lhs79] = _rhs111;
              _lhs80.f7 = _rhs112;
              _550_v2 = _rhs113;
              let _573_v24;
              _573_v24 = _dafny.Map.Empty.slice().updateUnsafe(_550_v2,p0);
              _548_v0 = (((_573_v24).contains((_dafny.ZERO).minus(_550_v2))) ? ((_573_v24).get((_dafny.ZERO).minus(_550_v2))) : (p0));
              let _574_v25;
              _574_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23(_550_v2, globalState),(_this).fm3(globalState));
              _574_v25 = _574_v25;
              let _rhs114 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_572_v23, _572_v23))).Union(_571_v22);
              let _rhs115 = (_dafny.ZERO).minus((_550_v2).multipliedBy(_550_v2));
              let _lhs81 = globalState;
              _571_v22 = _rhs114;
              _lhs81.f0 = _rhs115;
            }
            (globalState).f0 = (new BigNumber(((((_this).f6) ? (_dafny.MultiSet.fromElements((_this).f6)) : ((_this).f5))).cardinality())).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ckkg")).length));
            let _575_v26;
            _575_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,_550_v2);
            let _576_v27;
            _576_v27 = _module.D7.create_DC19(_575_v26);
            _576_v27 = _module.D7.create_DC19(_575_v26);
            let _577_v28;
            let _nw94 = new _module.C2();
            _nw94.__ctor((_this).f5);
            _577_v28 = _nw94;
            let _578_v29;
            _578_v29 = _module.D9.create_DC26(_577_v28);
            _577_v28 = (_578_v29).dtor_cf41;
          }
        }
      }
      let _579_v30;
      _579_v30 = _module.D3.create_DC10(false, _550_v2);
      let _580_v31;
      let _nw95 = new _module.C1();
      _nw95.__ctor(_this.f8, _module.__default.fm17(globalState), (_this).fm1((_579_v30).dtor_cf11, _550_v2, new BigNumber((_552_v4).length), globalState), ((_this).f5).Union(_dafny.MultiSet.fromElements(!(_this.f8))));
      _580_v31 = _nw95;
      _558_v10 = _558_v10;
      return;
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      (_this).f8 = (_this).f6;
      if (true) {
        let _arr5 = _this.f7;
        let _index101 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_this.f7).length));
        _arr5[_index101] = true;
        let _581_v0;
        let _init14 = function (_582_i0) {
          return (_582_i0).plus(new BigNumber(-680));
        };
        let _nw96 = Array((new BigNumber(4)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw96.length); _i0_14++) {
          _nw96[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _581_v0 = _nw96;
        let _583_v1;
        _583_v1 = new BigNumber(345);
        let _index102 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_581_v0).length));
        (_581_v0)[_index102] = (_583_v1).multipliedBy(_583_v1);
        let _584_v2;
        _584_v2 = _dafny.Seq.of((_this).fm2((_this).f6, globalState), (_this).f6, (_this).f6, _this.f8, _this.f8);
        let _arr6 = _this.f7;
        let _index103 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_this.f7).length));
        let _index104 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_581_v0).length));
        let _rhs116 = _this.f7;
        let _rhs117 = (_this).fm3(globalState);
        let _rhs118 = (new BigNumber((_584_v2).length)).isEqualTo((_this).fm6(globalState));
        let _rhs119 = _583_v1;
        let _lhs82 = _this;
        let _lhs83 = _this;
        let _lhs84 = _this.f7;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_this.f7).length));
        let _lhs86 = _581_v0;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_581_v0).length));
        _lhs82.f7 = _rhs116;
        _lhs83.f8 = _rhs117;
        _lhs84[_lhs85] = _rhs118;
        _lhs86[_lhs87] = _rhs119;
        _584_v2 = _584_v2;
        let _585_v3;
        let _init15 = function (_586_i1) {
          return false;
        };
        let _nw97 = Array((new BigNumber(18)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw97.length); _i0_15++) {
          _nw97[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _585_v3 = _nw97;
        let _587_v4;
        _587_v4 = _dafny.MultiSet.fromElements(_585_v3);
        let _588_v5;
        _588_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(232),_this.f8);
        let _589_v6;
        _589_v6 = _dafny.Map.Empty.slice().updateUnsafe(_587_v4,_588_v5);
        _589_v6 = (_589_v6).update(_587_v4, _dafny.Map.Empty.slice().updateUnsafe(_583_v1,(_this).f6));
        let _590_v7;
        let _init16 = function (_591_i2) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        };
        let _nw98 = Array((new BigNumber(28)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw98.length); _i0_16++) {
          _nw98[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _590_v7 = _nw98;
        (globalState).f2 = _590_v7;
        let _592_v8;
        _592_v8 = _dafny.Seq.of((_581_v0)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_581_v0).length))], new BigNumber((_584_v2).length), _583_v1, new BigNumber(778));
        let _arr7 = _this.f7;
        let _index105 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_this.f7).length));
        let _rhs120 = _581_v0;
        let _rhs121 = (_581_v0)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_581_v0).length))];
        let _rhs122 = _dafny.Seq.Concat(_dafny.Seq.Concat(_592_v8, _592_v8), _592_v8);
        let _rhs123 = (_this).f6;
        let _lhs88 = globalState;
        let _lhs89 = _this.f7;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_this.f7).length));
        _581_v0 = _rhs120;
        _lhs88.f0 = _rhs121;
        _592_v8 = _rhs122;
        _lhs89[_lhs90] = _rhs123;
      } else {
        let _593_v9;
        let _nw99 = new _module.C2();
        _nw99.__ctor((_this).f5);
        _593_v9 = _nw99;
        (_this).f8 = false;
        let _594_v10;
        _594_v10 = new _dafny.CodePoint('u'.codePointAt(0));
        _594_v10 = _594_v10;
        let _595_v11;
        _595_v11 = new BigNumber(442);
        (globalState).f0 = _595_v11;
        if (!(false)) {
          let _596_v12;
          _596_v12 = _dafny.Seq.UnicodeFromString("lh");
          let _597_v13;
          _597_v13 = _dafny.Seq.of(_596_v12);
          _596_v12 = _dafny.Seq.update((_597_v13)[_module.__default.safeIndex(_595_v11, new BigNumber((_597_v13).length))], _module.__default.safeIndex((_this).fm6(globalState), new BigNumber(((_597_v13)[_module.__default.safeIndex(_595_v11, new BigNumber((_597_v13).length))]).length)), _594_v10);
          _596_v12 = _596_v12;
          let _598_v14;
          let _nw100 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _598_v14 = _nw100;
          let _index106 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_598_v14).length));
          (_598_v14)[_index106] = _dafny.Seq.Concat(_596_v12, _596_v12);
          let _index107 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_598_v14).length));
          (_598_v14)[_index107] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("px"), _596_v12), _module.__default.safeIndex(_595_v11, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("px"), _596_v12)).length)), _594_v10), _596_v12);
          _594_v10 = _594_v10;
          let _599_v15;
          _599_v15 = _dafny.Seq.of((_this).f6);
          let _600_v16;
          _600_v16 = _dafny.MultiSet.fromElements(_599_v15);
          let _601_v17;
          _601_v17 = _module.D10.create_DC29(_600_v16);
          let _602_v18;
          let _nw101 = Array((new BigNumber(17)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = (_600_v16).Difference(_600_v16);
          _nw101[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_599_v15, _dafny.Seq.of((_this).f6));
          _nw101[(new BigNumber(2)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_599_v15);
          _nw101[(new BigNumber(4)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(5)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(6)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(_599_v15)).Union(_600_v16);
          _nw101[(new BigNumber(8)).toNumber()] = (_600_v16).update(_599_v15, _module.__default.abs(_595_v11));
          _nw101[(new BigNumber(9)).toNumber()] = (_600_v16).Union((_601_v17).dtor_cf45);
          _nw101[(new BigNumber(10)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(11)).toNumber()] = (_600_v16).Intersect(_dafny.MultiSet.fromElements(_599_v15, _599_v15, _599_v15, _599_v15));
          _nw101[(new BigNumber(12)).toNumber()] = (_600_v16).Difference(_600_v16);
          _nw101[(new BigNumber(13)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(14)).toNumber()] = _600_v16;
          _nw101[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_603_i3) {
            return _dafny.Seq.of(_this.f8, (_this).f6, true);
          }));
          _nw101[(new BigNumber(16)).toNumber()] = _600_v16;
          _602_v18 = _nw101;
          let _604_v19;
          _604_v19 = _dafny.Seq.of(_599_v15);
          let _index108 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_602_v18).length));
          (_602_v18)[_index108] = _dafny.MultiSet.FromArray(_604_v19);
          let _index109 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_602_v18).length));
          (_602_v18)[_index109] = _dafny.MultiSet.fromElements(_599_v15);
        } else {
          let _arr8 = _this.f7;
          let _index110 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f7).length));
          _arr8[_index110] = _this.f8;
          let _605_v21;
          _605_v21 = _module.D2.create_DC4(function () {
  let _coll33 = new _dafny.Set();
  for (const _compr_33 of _dafny.IntegerRange(new BigNumber(851), new BigNumber(-711))) {
    let _606_v20 = _compr_33;
    if (((new BigNumber(851)).isLessThanOrEqualTo(_606_v20)) && ((_606_v20).isLessThan(new BigNumber(-711)))) {
      _coll33.add((_606_v20).multipliedBy(_595_v11));
    }
  }
  return _coll33;
}());
          let _607_v22;
          _607_v22 = _dafny.Seq.of(_595_v11, _595_v11);
          let _608_v23;
          _608_v23 = _dafny.Set.fromElements(new BigNumber((_607_v22).length), _595_v11);
          let _arr9 = _this.f7;
          let _index111 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f7).length));
          _arr9[_index111] = (_608_v23).IsSubsetOf((_605_v21).dtor_cf4);
          let _609_v24;
          _609_v24 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("kypgb"));
          let _610_v25;
          _610_v25 = _dafny.Seq.of(true);
          let _611_v26;
          _611_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_610_v25, _module.__default.safeIndex(_595_v11, new BigNumber((_610_v25).length)), _this.f8),_dafny.Seq.Concat(_609_v24, _609_v24));
          _609_v24 = (((_611_v26).contains(_610_v25)) ? ((_611_v26).get(_610_v25)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), function (_612_i4) {
            return _dafny.Seq.UnicodeFromString("ykalvw");
          })));
          _609_v24 = _609_v24;
          let _613_v27;
          _613_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this.f7)[_module.__default.safeIndex(new BigNumber(240), new BigNumber((_this.f7).length))]),_this.f8);
          _613_v27 = (_613_v27).update((_this).f5, ((_dafny.ZERO).minus(_595_v11)).isLessThan(_595_v11));
          let _614_v28;
          let _out8;
          _out8 = (_593_v9).m9(!((_this).f6), _595_v11, globalState);
          _614_v28 = _out8;
        }
      }
      let _615_v29;
      _615_v29 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(250)), function (_616_i5) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }));
      let _arr10 = _this.f7;
      let _index112 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
      _arr10[_index112] = _dafny.areEqual(_615_v29, _615_v29);
      let _arr11 = _this.f7;
      let _index113 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
      _arr11[_index113] = !((_this).fm3(globalState));
      let _617_v30;
      _617_v30 = new BigNumber(967);
      let _hi2 = _617_v30;
      for (let _618_i6 = _617_v30; _618_i6.isLessThan(_hi2); _618_i6 = _618_i6.plus(_dafny.ONE)) {
        let _619_v31;
        _619_v31 = new _dafny.CodePoint('g'.codePointAt(0));
        let _620_v32;
        _620_v32 = _dafny.Map.Empty.slice().updateUnsafe(_617_v30,new BigNumber(((_this).f5).cardinality()));
        let _arr12 = _this.f7;
        let _index114 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        _arr12[_index114] = (_617_v30).isLessThanOrEqualTo((_618_i6).plus((_module.D2.create_DC7(_619_v31, _620_v32, new BigNumber(-433))).dtor_cf8));
        let _621_v33;
        let _init17 = ((_622_i6) => function (_623_i7) {
          return _module.__default.safeModuloInt(_623_i7, _622_i6);
        })(_618_i6);
        let _nw102 = Array((new BigNumber(29)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw102.length); _i0_17++) {
          _nw102[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _621_v33 = _nw102;
        let _index115 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_621_v33).length));
        (_621_v33)[_index115] = new BigNumber(728);
        let _arr13 = _this.f7;
        let _index116 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        let _index117 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_621_v33).length));
        let _rhs124 = _this.f8;
        let _rhs125 = (_618_i6).multipliedBy(_618_i6);
        let _lhs91 = _this.f7;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        let _lhs93 = _621_v33;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_621_v33).length));
        _lhs91[_lhs92] = _rhs124;
        _lhs93[_lhs94] = _rhs125;
        let _arr14 = _this.f7;
        let _index118 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        let _rhs126 = (_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))];
        let _rhs127 = (_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))];
        let _rhs128 = ((_this.f8) ? (_615_v29) : (_615_v29));
        let _lhs95 = _this;
        let _lhs96 = _this.f7;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        _lhs95.f8 = _rhs126;
        _lhs96[_lhs97] = _rhs127;
        _615_v29 = _rhs128;
        let _624_v34;
        _624_v34 = _dafny.Seq.of(_617_v30);
        let _625_v35;
        _625_v35 = _dafny.Seq.of(_624_v34);
        let _626_v36;
        _626_v36 = _dafny.Seq.UnicodeFromString("hj");
        let _627_v37;
        _627_v37 = _dafny.Map.Empty.slice().updateUnsafe((_625_v35)[_module.__default.safeIndex(new BigNumber((_626_v36).length), new BigNumber((_625_v35).length))],false);
        _627_v37 = _627_v37;
      }
      if (false) {
        (globalState).f0 = _617_v30;
        let _628_v38;
        let _nw103 = Array((new BigNumber(28)).toNumber()).fill(false);
        _628_v38 = _nw103;
        let _629_v39;
        _629_v39 = _dafny.Map.Empty.slice().updateUnsafe(_628_v38,(_this).f6);
        let _630_v40;
        _630_v40 = _module.D9.create_DC28(_module.D9.create_DC27(_629_v39, (_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))]));
        _630_v40 = _630_v40;
        (_this).f8 = _this.f8;
        (globalState).f0 = (((_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))]) ? (new BigNumber(560)) : (new BigNumber(789)));
        let _631_v41;
        _631_v41 = _dafny.MultiSet.fromElements(_617_v30, (_this).fm6(globalState), (_this).fm6(globalState));
        let _632_v42;
        _632_v42 = _dafny.Set.fromElements(_617_v30);
        let _633_v43;
        _633_v43 = _dafny.Seq.UnicodeFromString("gtark");
        let _634_v44;
        _634_v44 = _dafny.Map.Empty.slice().updateUnsafe(_633_v43,(_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))]);
        let _635_v45;
        _635_v45 = _dafny.Seq.of(_617_v30, (_this).fm4(_617_v30, _617_v30, _634_v44, _617_v30, globalState));
        let _636_v46;
        _636_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_617_v30),_this.f8);
        let _637_v47;
        _637_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_636_v46).length),_617_v30);
        let _638_v48;
        let _nw104 = Array((new BigNumber(28)).toNumber());
        _nw104[(_dafny.ZERO).toNumber()] = _617_v30;
        _nw104[(_dafny.ONE).toNumber()] = _617_v30;
        _nw104[(new BigNumber(2)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(3)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(4)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(5)).toNumber()] = new BigNumber((_631_v41).cardinality());
        _nw104[(new BigNumber(6)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(7)).toNumber()] = new BigNumber(433);
        _nw104[(new BigNumber(8)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(9)).toNumber()] = new BigNumber((_module.__default.fm15((_this).f6, globalState)).cardinality());
        _nw104[(new BigNumber(10)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(11)).toNumber()] = new BigNumber((_632_v42).length);
        _nw104[(new BigNumber(12)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(13)).toNumber()] = new BigNumber((_635_v45).length);
        _nw104[(new BigNumber(14)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(15)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(16)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(17)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(18)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(19)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(20)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(21)).toNumber()] = new BigNumber((((_this).f5).update(false, _module.__default.abs(new BigNumber((_636_v46).length)))).cardinality());
        _nw104[(new BigNumber(22)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(23)).toNumber()] = new BigNumber((_637_v47).length);
        _nw104[(new BigNumber(24)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(25)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(26)).toNumber()] = _617_v30;
        _nw104[(new BigNumber(27)).toNumber()] = _617_v30;
        _638_v48 = _nw104;
        let _639_v49;
        _639_v49 = _dafny.Seq.of(_638_v48);
        _639_v49 = _dafny.Seq.of(_638_v48, _638_v48, _638_v48);
      } else {
        let _640_v50;
        _640_v50 = _dafny.Map.Empty.slice().updateUnsafe((_617_v30).isLessThan(_617_v30),(_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))]);
        _640_v50 = _640_v50;
        let _641_v51;
        _641_v51 = _dafny.Seq.UnicodeFromString("ccqgh");
        let _642_v52;
        _642_v52 = _dafny.Set.fromElements(_641_v51);
        let _arr15 = _this.f7;
        let _index119 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length));
        _arr15[_index119] = (function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of (_615_v29).Elements) {
            let _643_v53 = _compr_34;
            if (_dafny.Seq.contains(_615_v29, _643_v53)) {
              _coll34.add(_643_v53);
            }
          }
          return _coll34;
        }()).IsSubsetOf(_642_v52);
        (globalState).f0 = _617_v30;
        let _644_v54;
        _644_v54 = _dafny.Map.Empty.slice().updateUnsafe(!(!(!((_this).fm3(globalState)))),_dafny.Seq.UnicodeFromString("pa"));
        _644_v54 = (_644_v54).update((_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))], _dafny.Seq.UnicodeFromString("qc"));
        _617_v30 = _module.__default.safeModuloInt(_617_v30, (_617_v30).multipliedBy(new BigNumber(167)));
      }
      let _645_v55;
      _645_v55 = _dafny.Seq.UnicodeFromString("djxkfjewt");
      let _646_v56;
      _646_v56 = _dafny.Seq.of(_617_v30, new BigNumber((_645_v55).length));
      let _647_v57;
      _647_v57 = _dafny.Map.Empty.slice().updateUnsafe(_646_v56,_617_v30);
      _647_v57 = (_647_v57).update(_dafny.Seq.of(_617_v30, _617_v30, new BigNumber((_645_v55).length)), (_617_v30).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))])).length))));
      let _648_v58;
      _648_v58 = _dafny.Set.fromElements(_617_v30, _617_v30, _617_v30, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), function (_649_i8) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length));
      let _650_v59;
      _650_v59 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f8),_this.f8);
      let _651_v60;
      _651_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("kdtsapr"),(_this.f7)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_this.f7).length))]);
      r0 = (_648_v58).Intersect((_dafny.Set.fromElements(_617_v30)).Intersect(_module.__default.fm24((_this).fm4(_617_v30, new BigNumber((_650_v59).length), _651_v60, _617_v30, globalState), (_this).f6, new BigNumber(111), globalState)));
      r1 = (_617_v30).minus(_617_v30);
      let _652_v61;
      _652_v61 = new _dafny.CodePoint('r'.codePointAt(0));
      let _653_v62;
      _653_v62 = _dafny.Map.Empty.slice().updateUnsafe(_652_v61,_this.f8);
      r2 = _module.__default.safeModuloInt(_617_v30, new BigNumber(((_653_v62).update(_652_v61, (_this).fm2(_this.f8, globalState))).length));
      return [r0, r1, r2];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      (_this).f8 = true;
      let _654_v0;
      _654_v0 = new _dafny.CodePoint('t'.codePointAt(0));
      let _655_v1;
      _655_v1 = _module.D7.create_DC21(p0, p0);
      let _656_v2;
      _656_v2 = _dafny.Map.Empty.slice().updateUnsafe((_655_v1).dtor_cf30,p0);
      let _657_v3;
      _657_v3 = _module.D2.create_DC7(_654_v0, _656_v2, p0);
      let _658_i0;
      _658_i0 = _dafny.ZERO;
      L8: {
        let _pat_let_tv16 = p0;
        let _pat_let_tv17 = p0;
        let _pat_let_tv18 = p0;
        while (function (_source5) {
          if (_source5.is_DC5) {
            let _664___mcc_h0 = (_source5).cf5;
            let _665_cf5 = _664___mcc_h0;
            return _665_cf5;
          } else if (_source5.is_DC6) {
            return (_this).f6;
          } else if (_source5.is_DC7) {
            let _666___mcc_h1 = (_source5).cf6;
            let _667___mcc_h2 = (_source5).cf7;
            let _668___mcc_h3 = (_source5).cf8;
            let _669_cf8 = _668___mcc_h3;
            let _670_cf7 = _667___mcc_h2;
            let _671_cf6 = _666___mcc_h1;
            return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv16,(_this).f6)).equals(_dafny.Map.Empty.slice().updateUnsafe(_669_cf8,(_this).f6));
          } else if (_source5.is_DC4) {
            let _672___mcc_h4 = (_source5).cf4;
            let _673_cf4 = _672___mcc_h4;
            return (_pat_let_tv17).isLessThanOrEqualTo(_pat_let_tv18);
          } else {
            let _674___mcc_h5 = (_source5).cf9;
            let _675_cf9 = _674___mcc_h5;
            return _this.f8;
          }
        }(_module.D2.create_DC8(_657_v3))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_658_i0)) {
              break L8;
            }
            _658_i0 = (_658_i0).plus(_dafny.ONE);
            (_this).f7 = _this.f7;
            let _659_v4;
            let _init18 = function (_660_i1) {
              return _module.D10.create_DC29(_dafny.MultiSet.fromElements(_dafny.Seq.of(_this.f8, (_this).f6), _dafny.Seq.of((_this).f6)));
            };
            let _nw105 = Array((new BigNumber(27)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
              _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _659_v4 = _nw105;
            _659_v4 = _659_v4;
            let _661_v5;
            _661_v5 = _dafny.Seq.UnicodeFromString("mvhgcimfk");
            let _662_v6;
            _662_v6 = _module.D5.create_DC15(_661_v5);
            let _663_v7;
            _663_v7 = _module.D3.create_DC10((_this).f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_662_v6,true)).length));
            (_this).f8 = (_663_v7).dtor_cf11;
            (globalState).f0 = p0;
          }
        }
      }
      let _676_v8;
      _676_v8 = _dafny.MultiSet.fromElements(p0);
      let _677_v9;
      let _nw106 = new _module.C2();
      _nw106.__ctor((_this).f5);
      _677_v9 = _nw106;
      let _rhs129 = _676_v8;
      let _rhs130 = p0;
      let _rhs131 = (p2) === (p2);
      let _rhs132 = p1;
      let _rhs133 = ((p1) ? ((((_this).f6) ? (_677_v9) : (_677_v9))) : (_677_v9));
      let _lhs98 = globalState;
      let _lhs99 = _this;
      let _lhs100 = _this;
      _676_v8 = _rhs129;
      _lhs98.f0 = _rhs130;
      _lhs99.f8 = _rhs131;
      _lhs100.f8 = _rhs132;
      _677_v9 = _rhs133;
      let _678_v10;
      _678_v10 = _dafny.Seq.of(p0, p0, p0, p0, p0);
      let _679_v11;
      _679_v11 = _dafny.Seq.UnicodeFromString("lpaexls");
      let _680_v12;
      _680_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_678_v10, _678_v10),_dafny.Seq.contains(_679_v11, _654_v0));
      if ((((_680_v12).contains(_678_v10)) ? ((_680_v12).get(_678_v10)) : (_this.f8))) {
        (globalState).f0 = p0;
        let _681_v13;
        _681_v13 = _module.D2.create_DC8(_657_v3);
        let _source6 = _681_v13;
        if (_source6.is_DC5) {
          let _682___mcc_h6 = (_source6).cf5;
          let _683_cf5 = _682___mcc_h6;
          _654_v0 = _654_v0;
          let _684_v14;
          let _nw107 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _684_v14 = _nw107;
          _684_v14 = _684_v14;
          _683_cf5 = p2;
          let _685_v15;
          _685_v15 = _module.D7.create_DC20(_683_cf5);
          (_this).f8 = (_685_v15).dtor_cf29;
        } else if (_source6.is_DC6) {
          (_this).f8 = p1;
          let _686_v17;
          _686_v17 = _module.D0.create_DC1();
          let _687_v18;
          _687_v18 = _module.D1.create_DC3(_686_v17, p0);
          let _rhs134 = !(!(function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of _dafny.IntegerRange(new BigNumber(687), new BigNumber(872))) {
              let _688_v16 = _compr_35;
              if (((new BigNumber(687)).isLessThanOrEqualTo(_688_v16)) && ((_688_v16).isLessThan(new BigNumber(872)))) {
                _coll35.push([_module.__default.safeDivisionInt(_688_v16, p0),(_this).f6]);
              }
            }
            return _coll35;
          }()).contains((p0).plus(p0)));
          let _rhs135 = (_687_v18).dtor_cf3;
          let _rhs136 = _this.f8;
          let _lhs101 = _this;
          let _lhs102 = globalState;
          let _lhs103 = _this;
          _lhs101.f8 = _rhs134;
          _lhs102.f0 = _rhs135;
          _lhs103.f8 = _rhs136;
          let _689_v19;
          _689_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f8);
          _689_v19 = (_689_v19).update((_this).f6, (_this).f6);
          (globalState).f0 = (_this).fm6(globalState);
        } else if (_source6.is_DC7) {
          let _690___mcc_h7 = (_source6).cf6;
          let _691___mcc_h8 = (_source6).cf7;
          let _692___mcc_h9 = (_source6).cf8;
          let _693_cf8 = _692___mcc_h9;
          let _694_cf7 = _691___mcc_h8;
          let _695_cf6 = _690___mcc_h7;
          (globalState).f0 = _693_cf8;
          let _696_v20;
          _696_v20 = _dafny.Seq.of(true, p1);
          let _697_v21;
          _697_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(!((_this).f6)),_696_v20);
          let _698_v22;
          _698_v22 = _module.D1.create_DC2((_this).f6);
          _697_v21 = (_697_v21).update(_698_v22, _module.__default.fm23(p0, globalState));
          let _arr16 = _this.f7;
          let _index120 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_this.f7).length));
          _arr16[_index120] = p1;
          let _699_v23;
          _699_v23 = _dafny.MultiSet.fromElements(_678_v10);
          let _700_v24;
          _700_v24 = _module.D0.create_DC1();
          let _701_v25;
          _701_v25 = _module.D1.create_DC3(_700_v24, p0);
          let _702_v26;
          _702_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f8);
          let _arr17 = _this.f7;
          let _index121 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_this.f7).length));
          let _rhs137 = (_677_v9).fm1((_699_v23).equals(_699_v23), (p0).minus(_693_cf8), (_701_v25).dtor_cf3, globalState);
          let _rhs138 = _dafny.Seq.Concat(_678_v10, _678_v10);
          let _rhs139 = (_this).fm2((_702_v26).equals(_702_v26), globalState);
          let _rhs140 = _678_v10;
          let _lhs104 = _this.f7;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_this.f7).length));
          let _lhs106 = _this;
          _lhs104[_lhs105] = _rhs137;
          _678_v10 = _rhs138;
          _lhs106.f8 = _rhs139;
          _678_v10 = _rhs140;
          (globalState).f0 = (_dafny.ZERO).minus(((_this).fm6(globalState)).plus(_693_cf8));
        } else if (_source6.is_DC4) {
          let _703___mcc_h10 = (_source6).cf4;
          let _704_cf4 = _703___mcc_h10;
          (_this).f8 = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vurhyya"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(844)), ((_705_v0) => function (_706_i2) {
            return _705_v0;
          })(_654_v0))), _654_v0);
          let _707_v27;
          let _nw108 = Array((new BigNumber(12)).toNumber());
          _707_v27 = _nw108;
          (_this).f8 = _dafny.Seq.contains(_679_v11, (_679_v11)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_707_v27)).cardinality()), new BigNumber((_679_v11).length))]);
          (_this).f8 = !(p2);
          let _708_v28;
          _708_v28 = _dafny.Map.Empty.slice().updateUnsafe((p1) && (_this.f8),(_this).fm6(globalState));
          let _709_v29;
          _709_v29 = _dafny.Seq.of((_this).f6);
          _708_v28 = (_708_v28).update((_this).f6, (new BigNumber((_dafny.Seq.update(_709_v29, _module.__default.safeIndex(p0, new BigNumber((_709_v29).length)), (_this).f6)).length)).plus(p0));
        } else {
          let _710___mcc_h11 = (_source6).cf9;
          let _711_cf9 = _710___mcc_h11;
          let _712_v30;
          _712_v30 = _dafny.Seq.of(_this.f8);
          _712_v30 = _dafny.Seq.Concat(_dafny.Seq.Concat(_712_v30, _712_v30), _dafny.Seq.of(false));
          (_this).f8 = _dafny.Seq.IsPrefixOf((_this).fm7(globalState), _679_v11);
          let _713_v31;
          let _nw109 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _713_v31 = _nw109;
          let _714_v32;
          _714_v32 = _dafny.MultiSet.fromElements(_713_v31);
          (_this).f8 = (_714_v32).equals(_714_v32);
          (_this).f7 = _this.f7;
        }
        let _715_v33;
        _715_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f8);
        _715_v33 = _715_v33;
        if (p2) {
          let _init19 = function (_716_i3) {
            return true;
          };
          let _nw110 = Array((new BigNumber(26)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw110.length); _i0_19++) {
            _nw110[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          (_this).f7 = _nw110;
          let _717_v34;
          _717_v34 = _module.D0.create_DC0(p0);
          let _718_v35;
          _718_v35 = _dafny.Seq.of(_module.D0.create_DC0(p0), _717_v34, _717_v34, ((p2) ? (_717_v34) : (_717_v34)), _module.__default.fm25((_this).f6, p0, p2, globalState));
          (globalState).f0 = new BigNumber((_718_v35).length);
          let _719_v36;
          _719_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p2);
          let _720_v37;
          _720_v37 = _module.D9.create_DC27(_719_v36, ((p2) ? (_this.f8) : (p1)));
          _720_v37 = _720_v37;
          (_this).f8 = (_this).f6;
          let _721_v38;
          let _nw111 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _721_v38 = _nw111;
          let _722_v40;
          _722_v40 = _dafny.Seq.of((_this).f5);
          let _723_v41;
          _723_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f6);
          let _index122 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_721_v38).length));
          (_721_v38)[_index122] = (function () {
            let _coll36 = new _dafny.Map();
            for (const _compr_36 of (_722_v40).Elements) {
              let _724_v39 = _compr_36;
              if (_dafny.Seq.contains(_722_v40, _724_v39)) {
                _coll36.push([_724_v39,_this.f8]);
              }
            }
            return _coll36;
          }()).Merge(_723_v41);
          let _index123 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_721_v38).length));
          (_721_v38)[_index123] = (_723_v41).Merge(_723_v41);
        } else {
          _654_v0 = _654_v0;
          let _725_v42;
          let _nw112 = new _module.C2();
          _nw112.__ctor((_677_v9).f5);
          _725_v42 = _nw112;
          let _726_v43;
          _726_v43 = _module.D2.create_DC7(_654_v0, _656_v2, p0);
          (globalState).f0 = (_726_v43).dtor_cf8;
          (_this).f7 = _this.f7;
          (globalState).f0 = (_dafny.ZERO).minus(p0);
        }
        let _727_v44;
        _727_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
        let _728_v45;
        let _nw113 = new _module.C1();
        _nw113.__ctor(p2, _dafny.Seq.Concat(_679_v11, _679_v11), (((_727_v44).contains(p0)) ? ((_727_v44).get(p0)) : (false)), (_this).f5);
        _728_v45 = _nw113;
      } else {
        let _729_v46;
        let _nw114 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _729_v46 = _nw114;
        let _730_v47;
        _730_v47 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f6);
        let _index124 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_729_v46).length));
        (_729_v46)[_index124] = (_730_v47).Merge(_730_v47);
        let _731_v48;
        _731_v48 = _module.D11.create_DC32(_730_v47);
        let _index125 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_729_v46).length));
        (_729_v46)[_index125] = (_730_v47).Merge(((_731_v48).dtor_cf51).Merge(_730_v47));
        let _732_v49;
        _732_v49 = _module.D1.create_DC2(p1);
        let _source7 = _732_v49;
        if (_source7.is_DC3) {
          let _733___mcc_h12 = (_source7).cf2;
          let _734___mcc_h13 = (_source7).cf3;
          let _735_cf3 = _734___mcc_h13;
          let _736_cf2 = _733___mcc_h12;
          _679_v11 = _679_v11;
          (globalState).f0 = (_735_cf3).minus(p0);
          let _737_v50;
          _737_v50 = _dafny.Seq.of(p2);
          let _738_v51;
          let _nw115 = new _module.C1();
          _nw115.__ctor(p2, _679_v11, _dafny.Seq.IsProperPrefixOf(_module.__default.fm23(_735_cf3, globalState), _737_v50), (_677_v9).f5);
          _738_v51 = _nw115;
          let _739_v52;
          _739_v52 = _module.D3.create_DC11(new BigNumber(58));
          let _740_v53;
          _740_v53 = _dafny.Map.Empty.slice().updateUnsafe(p2,_739_v52);
          _735_cf3 = (_735_cf3).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), new BigNumber((_740_v53).length)));
        } else {
          let _741___mcc_h14 = (_source7).cf1;
          let _742_cf1 = _741___mcc_h14;
          let _743_v54;
          _743_v54 = _dafny.Set.fromElements(p2);
          _743_v54 = _743_v54;
          let _744_v55;
          let _nw116 = Array((new BigNumber(28)).toNumber()).fill(_module.D2.Default());
          _744_v55 = _nw116;
          _744_v55 = _744_v55;
          (globalState).f0 = (_module.__default.safeDivisionInt(p0, p0)).plus(p0);
          let _745_v56;
          let _nw117 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _745_v56 = _nw117;
          let _746_v57;
          _746_v57 = _dafny.Map.Empty.slice().updateUnsafe(_654_v0,false);
          let _index126 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_745_v56).length));
          (_745_v56)[_index126] = _746_v57;
          let _747_v58;
          _747_v58 = _module.D12.create_DC34(_746_v57);
          let _index127 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_745_v56).length));
          (_745_v56)[_index127] = ((_747_v58).dtor_cf54).Merge(_746_v57);
        }
        let _748_v59;
        _748_v59 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_dafny.MultiSet.fromElements(p2));
        let _749_v60;
        let _init20 = function (_750_i4) {
          return _module.__default.safeModuloInt(_750_i4, new BigNumber(-705));
        };
        let _nw118 = Array((new BigNumber(7)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw118.length); _i0_20++) {
          _nw118[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _749_v60 = _nw118;
        let _index128 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_749_v60).length));
        (_749_v60)[_index128] = p0;
        let _751_v61;
        _751_v61 = _module.D13.create_DC38(_748_v59);
        let _752_v62;
        _752_v62 = _dafny.Set.fromElements((_this).f6, p2);
        let _753_v63;
        _753_v63 = _dafny.Seq.of(_this.f8);
        let _754_v64;
        _754_v64 = _dafny.Set.fromElements(new BigNumber((_753_v63).length));
        let _index129 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_749_v60).length));
        let _rhs141 = (_751_v61).dtor_cf59;
        let _rhs142 = new BigNumber(((_752_v62).Intersect((_752_v62).Intersect(_dafny.Set.fromElements((_753_v63)[_module.__default.safeIndex(p0, new BigNumber((_753_v63).length))])))).length);
        let _rhs143 = !(p1) || (p2);
        let _rhs144 = _module.__default.safeDivisionInt(p0, (p0).plus(new BigNumber((_754_v64).length)));
        let _lhs107 = _749_v60;
        let _lhs108 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_749_v60).length));
        let _lhs109 = _this;
        let _lhs110 = globalState;
        _748_v59 = _rhs141;
        _lhs107[_lhs108] = _rhs142;
        _lhs109.f8 = _rhs143;
        _lhs110.f0 = _rhs144;
        (globalState).f0 = ((_749_v60)[_module.__default.safeIndex(new BigNumber(48), new BigNumber((_749_v60).length))]).plus(p0);
        let _755_v65;
        let _nw119 = new _module.C2();
        _nw119.__ctor((_677_v9).f5);
        _755_v65 = _nw119;
      }
      if (_this.f8) {
        (_this).f8 = p2;
        let _756_v66;
        _756_v66 = _module.D5.create_DC15(_dafny.Seq.UnicodeFromString("uipsyskp"));
        let _pat_let_tv19 = _679_v11;
        let _757_v67;
        _757_v67 = _dafny.Map.Empty.slice().updateUnsafe(_678_v10,function (_pat_let24_0) {
          return function (_758_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_759_dt__update_hcf18_h0) {
                return _module.D5.create_DC15(_759_dt__update_hcf18_h0);
              }(_pat_let25_0);
            }(_pat_let_tv19);
          }(_pat_let24_0);
        }(_756_v66));
        _757_v67 = _757_v67;
        let _760_v68;
        _760_v68 = _dafny.Seq.of(!(p2));
        _656_v2 = (_656_v2).update(p0, (((_676_v8).contains(p0)) ? ((_676_v8).get(p0)) : (new BigNumber((_760_v68).length))));
        let _761_v69;
        _761_v69 = _module.D12.create_DC36(((_this.f8) ? (((((_this).f5).contains(_this.f8)) ? (((_this).f5).get(_this.f8)) : (p0))) : (p0)), p0);
        _761_v69 = _761_v69;
        (_this).f8 = true;
      } else {
        let _762_v70;
        _762_v70 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f7);
        _762_v70 = (_762_v70).Merge(_762_v70);
        _654_v0 = _654_v0;
        let _763_v71;
        _763_v71 = _module.D3.create_DC11(p0);
        let _764_v72;
        _764_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_763_v71).dtor_cf13),_679_v11);
        let _765_v73;
        _765_v73 = _dafny.Seq.of(p2);
        _764_v72 = (_764_v72).update(_module.__default.safeModuloInt(new BigNumber((_765_v73).length), p0), _679_v11);
        let _766_v74;
        let _init21 = ((_767_p0) => function (_768_i5) {
          return (_768_i5).multipliedBy(_767_p0);
        })(p0);
        let _nw120 = Array((new BigNumber(8)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw120.length); _i0_21++) {
          _nw120[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _766_v74 = _nw120;
        let _index130 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_766_v74).length));
        (_766_v74)[_index130] = p0;
        let _769_v75;
        _769_v75 = _dafny.Seq.of(_765_v73, _765_v73);
        let _index131 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_766_v74).length));
        let _rhs145 = (_678_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_769_v75, _769_v75)).length), new BigNumber((_678_v10).length))];
        let _rhs146 = (_module.__default.safeDivisionInt(p0, p0)).minus(new BigNumber((_765_v73).length));
        let _rhs147 = p1;
        let _rhs148 = _679_v11;
        let _rhs149 = _dafny.Seq.Concat(_dafny.Seq.update(_679_v11, _module.__default.safeIndex(p0, new BigNumber((_679_v11).length)), _654_v0), _dafny.Seq.Concat(_679_v11, _dafny.Seq.UnicodeFromString("uayxfpbbt")));
        let _lhs111 = _766_v74;
        let _lhs112 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_766_v74).length));
        let _lhs113 = globalState;
        let _lhs114 = _this;
        _lhs111[_lhs112] = _rhs145;
        _lhs113.f0 = _rhs146;
        _lhs114.f8 = _rhs147;
        _679_v11 = _rhs148;
        _679_v11 = _rhs149;
        let _770_v76;
        let _nw121 = new _module.C3();
        _nw121.__ctor(_this.f8, (_677_v9).f5);
        _770_v76 = _nw121;
      }
      let _771_v77;
      _771_v77 = _module.D12.create_DC36(p0, p0);
      let _source8 = _771_v77;
      if (_source8.is_DC35) {
        let _772___mcc_h15 = (_source8).cf55;
        let _773_cf55 = _772___mcc_h15;
        _654_v0 = _654_v0;
        let _774_v78;
        let _nw122 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _774_v78 = _nw122;
        let _index132 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length));
        (_774_v78)[_index132] = _679_v11;
        let _775_v79;
        let _nw123 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _775_v79 = _nw123;
        let _776_v80;
        _776_v80 = _dafny.Seq.of(_775_v79);
        let _index133 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length));
        (_774_v78)[_index133] = _dafny.Seq.update(_679_v11, _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_776_v80).length), p0), new BigNumber((_679_v11).length)), _654_v0);
        let _777_v81;
        _777_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(547),p2);
        let _778_v82;
        _778_v82 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_777_v81).length), new BigNumber((_678_v10).length)), _678_v10),p0);
        let _index134 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length));
        (_775_v79)[_index134] = (p0).minus(p0);
        let _arr18 = _this.f7;
        let _index135 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
        _arr18[_index135] = !(_773_cf55) || (!(p1));
        let _779_v83;
        _779_v83 = _dafny.Seq.of(p1);
        let _780_v84;
        _780_v84 = _dafny.Set.fromElements(p0, p0);
        let _781_v86;
        _781_v86 = _dafny.Seq.of(_656_v2, _656_v2, _656_v2, _656_v2, _656_v2);
        let _782_v87;
        let _nw124 = new _module.C0();
        _nw124.__ctor(function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_781_v86).Elements) {
            let _783_v85 = _compr_37;
            if (_dafny.Seq.contains(_781_v86, _783_v85)) {
              _coll37.push([_783_v85,_this.f8]);
            }
          }
          return _coll37;
        }());
        _782_v87 = _nw124;
        let _784_v88;
        _784_v88 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_677_v9).f5).cardinality())),_782_v87);
        let _index136 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length));
        let _arr19 = _this.f7;
        let _index137 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
        let _rhs150 = _dafny.Map.Empty.slice().updateUnsafe(_678_v10,new BigNumber((_779_v83).length));
        let _rhs151 = (_677_v9).fm1(p2, p0, p0, globalState);
        let _rhs152 = (p0).multipliedBy(((_678_v10)[_module.__default.safeIndex(new BigNumber((_780_v84).length), new BigNumber((_678_v10).length))]).minus(new BigNumber((_784_v88).length)));
        let _rhs153 = !((_this).f6);
        let _rhs154 = (p2) && (_dafny.Seq.IsPrefixOf(_679_v11, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("p"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)), _module.__default.fm16(p2, _656_v2, (_774_v78)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length))], globalState))));
        let _lhs115 = _this;
        let _lhs116 = _775_v79;
        let _lhs117 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length));
        let _lhs118 = _this.f7;
        let _lhs119 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
        _778_v82 = _rhs150;
        _lhs115.f8 = _rhs151;
        _lhs116[_lhs117] = _rhs152;
        _lhs118[_lhs119] = _rhs153;
        _773_cf55 = _rhs154;
        if (p1) {
          let _index138 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length));
          (_774_v78)[_index138] = _679_v11;
          (_this).f8 = _this.f8;
          let _785_v89;
          let _nw125 = new _module.C1();
          _nw125.__ctor(!(!(p2)) || (_this.f8), _679_v11, p1, (_677_v9).f5);
          _785_v89 = _nw125;
          let _786_v90;
          _786_v90 = _dafny.MultiSet.fromElements(_module.__default.fm18(globalState));
          let _787_v91;
          let _nw126 = Array((new BigNumber(14)).toNumber()).fill(_dafny.MultiSet.Empty);
          _787_v91 = _nw126;
          let _788_v92;
          _788_v92 = _dafny.Map.Empty.slice().updateUnsafe(((_773_cf55) ? (_787_v91) : (_787_v91)),p0);
          let _arr20 = _this.f7;
          let _index139 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
          let _rhs155 = (_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0))).multipliedBy((_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))]);
          let _rhs156 = !(_780_v84).contains((((_786_v90).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(810)), ((_791_p0) => function (_792_i6) {
            return _791_p0;
          })(p0)))) ? ((_786_v90).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(810)), ((_789_p0) => function (_790_i6) {
            return _789_p0;
          })(p0)))) : (p0)));
          let _rhs157 = (((_788_v92).contains(_787_v91)) ? ((_788_v92).get(_787_v91)) : (((_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))]).minus((_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))])));
          let _lhs120 = globalState;
          let _lhs121 = _this.f7;
          let _lhs122 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
          let _lhs123 = globalState;
          _lhs120.f0 = _rhs155;
          _lhs121[_lhs122] = _rhs156;
          _lhs123.f0 = _rhs157;
          let _793_v93;
          _793_v93 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yc"),(_785_v89).f12);
          let _794_v94;
          _794_v94 = _dafny.Seq.of((_785_v89).f13, (_774_v78)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length))], (_774_v78)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length))]);
          let _795_v95;
          _795_v95 = _dafny.Map.Empty.slice().updateUnsafe((_794_v94)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_796_v79) => function (_797_i7) {
            return (_796_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_796_v79).length))];
          })(_775_v79))).length), new BigNumber((_794_v94).length))],_654_v0);
          (globalState).f0 = (_785_v89).fm4(_module.__default.safeDivisionInt((_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))], (_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))]), (_775_v79)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_775_v79).length))], _793_v93, new BigNumber((_795_v95).length), globalState);
        } else {
          let _798_v96;
          let _nw127 = new _module.C0();
          _nw127.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_656_v2,(_this.f7)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length))]));
          _798_v96 = _nw127;
          let _799_v97;
          _799_v97 = _dafny.Map.Empty.slice().updateUnsafe(_773_cf55,true);
          let _arr21 = _this.f7;
          let _index140 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
          let _rhs158 = p1;
          let _rhs159 = (new BigNumber(((_799_v97).Merge(_799_v97)).length)).multipliedBy((_dafny.ZERO).minus(p0));
          let _rhs160 = _799_v97;
          let _lhs124 = _this.f7;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_this.f7).length));
          let _lhs126 = globalState;
          _lhs124[_lhs125] = _rhs158;
          _lhs126.f0 = _rhs159;
          _799_v97 = _rhs160;
          let _800_v98;
          _800_v98 = _dafny.Map.Empty.slice().updateUnsafe(_654_v0,_773_cf55);
          let _801_v99;
          _801_v99 = _dafny.MultiSet.fromElements(_800_v98);
          let _802_v100;
          let _nw128 = new _module.C1();
          _nw128.__ctor((_801_v99).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), ((_803_v98) => function (_804_i8) {
            return _803_v98;
          })(_800_v98)))), (_774_v78)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length))], _dafny.Seq.IsPrefixOf(_679_v11, (_774_v78)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_774_v78).length))]), (_677_v9).f5);
          _802_v100 = _nw128;
          _773_cf55 = p2;
          _775_v79 = _775_v79;
        }
      } else if (_source8.is_DC36) {
        let _805___mcc_h16 = (_source8).cf56;
        let _806___mcc_h17 = (_source8).cf57;
        let _807_cf57 = _806___mcc_h17;
        let _808_cf56 = _805___mcc_h16;
        (_this).f8 = !((_807_cf57).isLessThan(_807_cf57)) || (((p1) ? (_this.f8) : (p1)));
        _678_v10 = _678_v10;
        _807_cf57 = ((_678_v10)[_module.__default.safeIndex(_808_cf56, new BigNumber((_678_v10).length))]).minus(new BigNumber((_676_v8).cardinality()));
        (globalState).f0 = _808_cf56;
      } else if (_source8.is_DC37) {
        let _809___mcc_h18 = (_source8).cf58;
        let _810_cf58 = _809___mcc_h18;
        let _811_v101;
        let _nw129 = new _module.C1();
        _nw129.__ctor(p2, _679_v11, (_this).f6, _dafny.MultiSet.fromElements(p1));
        _811_v101 = _nw129;
        let _812_v102;
        _812_v102 = _dafny.Set.fromElements(_811_v101);
        if ((_812_v102).IsProperSubsetOf(_dafny.Set.fromElements(_811_v101, _811_v101))) {
          (_this).f8 = _this.f8;
          let _813_v103;
          let _nw130 = new _module.C0();
          _nw130.__ctor(_module.__default.fm26(p0, globalState));
          _813_v103 = _nw130;
          (globalState).f0 = ((_module.D6.create_DC18(_813_v103, _module.__default.fm23(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_813_v103,new BigNumber(179))).length), globalState), new BigNumber((_679_v11).length))).dtor_cf27).minus(p0);
          let _814_v104;
          let _nw131 = new _module.C1();
          _nw131.__ctor(((_677_v9).f5).IsDisjointFrom((_this).f5), _dafny.Seq.UnicodeFromString("j"), _this.f8, (_this).f5);
          _814_v104 = _nw131;
          _656_v2 = (_656_v2).update(p0, _module.__default.safeModuloInt(new BigNumber(((_677_v9).f5).cardinality()), p0));
          let _815_v105;
          let _nw132 = new _module.C3();
          _nw132.__ctor(p1, _dafny.MultiSet.fromElements((_811_v101).f12));
          _815_v105 = _nw132;
        } else {
          (_this).f8 = (p0).isLessThan((_dafny.ZERO).minus(p0));
          (_this).f8 = (p0).isEqualTo(_module.__default.safeModuloInt(p0, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(509),(_811_v101).f12))).length)));
          let _816_v106;
          _816_v106 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-876)), function (_817_i9) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }),(_811_v101).f12);
          let _818_v107;
          _818_v107 = _dafny.Map.Empty.slice().updateUnsafe((_811_v101).fm4(p0, new BigNumber(-167), _816_v106, (_811_v101).fm4((_dafny.ZERO).minus(p0), p0, _dafny.Map.Empty.slice().updateUnsafe(_679_v11,p1), p0, globalState), globalState),_810_cf58);
          _818_v107 = (_818_v107).update((_dafny.ZERO).minus(p0), _module.__default.fm16(_this.f8, _656_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-403)), ((_819_v0) => function (_820_i10) {
            return _819_v0;
          })(_654_v0)), globalState));
          let _821_v108;
          _821_v108 = _dafny.Seq.of((_811_v101).f12, p1);
          let _822_v109;
          _822_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(999),!(_this.f8));
          let _823_v110;
          _823_v110 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_811_v101).fm3(globalState));
          let _824_v111;
          _824_v111 = _module.D8.create_DC24(p0, (_811_v101).f12, _821_v108, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm8(p0, (((_822_v109).contains(new BigNumber((_823_v110).length))) ? ((_822_v109).get(new BigNumber((_823_v110).length))) : (p2)), _679_v11, globalState)).length)), p0);
          (globalState).f0 = new BigNumber(((_824_v111).dtor_cf36).length);
          let _825_v113;
          _825_v113 = _dafny.Set.fromElements(_dafny.Seq.update((_811_v101).f13, _module.__default.safeIndex(p0, new BigNumber(((_811_v101).f13).length)), _654_v0));
          let _826_v114;
          _826_v114 = _module.D15.create_DC41(_816_v106);
          (globalState).f0 = (_this).fm4((_this).fm4(p0, p0, function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_825_v113).Elements) {
              let _827_v112 = _compr_38;
              if ((_825_v113).contains(_827_v112)) {
                _coll38.push([_827_v112,true]);
              }
            }
            return _coll38;
          }(), (_dafny.ZERO).minus(p0), globalState), new BigNumber(184), (_826_v114).dtor_cf66, (p0).minus(new BigNumber((_818_v107).length)), globalState);
        }
        let _rhs161 = (_this).fm3(globalState);
        let _rhs162 = p2;
        let _lhs127 = _this;
        let _lhs128 = _this;
        _lhs127.f8 = _rhs161;
        _lhs128.f8 = _rhs162;
        _654_v0 = _810_cf58;
        _676_v8 = _676_v8;
      } else {
        let _828___mcc_h19 = (_source8).cf54;
        let _829_cf54 = _828___mcc_h19;
        _678_v10 = _678_v10;
        let _830_v115;
        let _nw133 = new _module.C2();
        _nw133.__ctor((_this).f5);
        _830_v115 = _nw133;
        let _831_v116;
        _831_v116 = _module.D9.create_DC26(_830_v115);
        let _source9 = _831_v116;
        if (_source9.is_DC27) {
          let _832___mcc_h20 = (_source9).cf42;
          let _833___mcc_h21 = (_source9).cf43;
          let _834_cf43 = _833___mcc_h21;
          let _835_cf42 = _832___mcc_h20;
          _679_v11 = _679_v11;
          let _836_v117;
          let _nw134 = Array((new BigNumber(14)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("sspyxy");
          _nw134[(_dafny.ONE).toNumber()] = _679_v11;
          _nw134[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("enh");
          _nw134[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), ((_837_v0) => function (_838_i11) {
            return _837_v0;
          })(_654_v0));
          _nw134[(new BigNumber(4)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(5)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(6)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(7)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(8)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(9)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(10)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(11)).toNumber()] = _679_v11;
          _nw134[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("uu");
          _nw134[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), ((_839_v0) => function (_840_i12) {
            return _839_v0;
          })(_654_v0));
          _836_v117 = _nw134;
          let _841_v118;
          _841_v118 = _dafny.Map.Empty.slice().updateUnsafe(_836_v117,(_this).f5);
          _841_v118 = (_841_v118).update(_836_v117, _dafny.MultiSet.fromElements((_this).f6));
          let _842_v119;
          _842_v119 = _dafny.Seq.of(false);
          _842_v119 = _842_v119;
          (globalState).f0 = (p0).minus((_dafny.ZERO).minus(((((_677_v9).f5).contains(p1)) ? (((_677_v9).f5).get(p1)) : (p0))));
        } else if (_source9.is_DC26) {
          let _843___mcc_h22 = (_source9).cf41;
          let _844_cf41 = _843___mcc_h22;
          let _845_v120;
          let _nw135 = new _module.C1();
          _nw135.__ctor((_this).f6, _679_v11, p1, (_677_v9).f5);
          _845_v120 = _nw135;
          _654_v0 = (((_845_v120).f12) ? (_654_v0) : (_654_v0));
          let _arr22 = _this.f7;
          let _index141 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_this.f7).length));
          _arr22[_index141] = (_845_v120).f12;
          let _846_v121;
          _846_v121 = _dafny.Map.Empty.slice().updateUnsafe((_845_v120).f13,p2);
          let _847_v122;
          _847_v122 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(634),_654_v0);
          let _848_v123;
          let _nw136 = Array((new BigNumber(13)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = p0;
          _nw136[(_dafny.ONE).toNumber()] = p0;
          _nw136[(new BigNumber(2)).toNumber()] = p0;
          _nw136[(new BigNumber(3)).toNumber()] = (_this).fm6(globalState);
          _nw136[(new BigNumber(4)).toNumber()] = (_845_v120).fm4(new BigNumber(109), p0, _846_v121, p0, globalState);
          _nw136[(new BigNumber(5)).toNumber()] = p0;
          _nw136[(new BigNumber(6)).toNumber()] = p0;
          _nw136[(new BigNumber(7)).toNumber()] = (new BigNumber((_847_v122).length)).minus((_dafny.ZERO).minus(p0));
          _nw136[(new BigNumber(8)).toNumber()] = p0;
          _nw136[(new BigNumber(9)).toNumber()] = p0;
          _nw136[(new BigNumber(10)).toNumber()] = p0;
          _nw136[(new BigNumber(11)).toNumber()] = p0;
          _nw136[(new BigNumber(12)).toNumber()] = (((_656_v2).contains(new BigNumber((_dafny.Seq.UnicodeFromString("nqbu")).length))) ? ((_656_v2).get(new BigNumber((_dafny.Seq.UnicodeFromString("nqbu")).length))) : (new BigNumber((_679_v11).length)));
          _848_v123 = _nw136;
          let _index142 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_848_v123).length));
          (_848_v123)[_index142] = (_this).fm4(p0, p0, _846_v121, (_678_v10)[_module.__default.safeIndex(p0, new BigNumber((_678_v10).length))], globalState);
          let _849_v124;
          _849_v124 = _dafny.Set.fromElements(new BigNumber(9));
          let _arr23 = _this.f7;
          let _index143 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_this.f7).length));
          let _index144 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_848_v123).length));
          let _rhs163 = p0;
          let _rhs164 = p2;
          let _rhs165 = _module.__default.safeModuloInt(new BigNumber(((_849_v124).Difference(_849_v124)).length), new BigNumber(((_677_v9).f5).cardinality()));
          let _rhs166 = (_this).f6;
          let _lhs129 = globalState;
          let _lhs130 = _this.f7;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_this.f7).length));
          let _lhs132 = _848_v123;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_848_v123).length));
          let _lhs134 = _this;
          _lhs129.f0 = _rhs163;
          _lhs130[_lhs131] = _rhs164;
          _lhs132[_lhs133] = _rhs165;
          _lhs134.f8 = _rhs166;
          let _850_v125;
          _850_v125 = _dafny.Set.fromElements((_this).f6);
          let _851_v126;
          _851_v126 = _dafny.Set.fromElements(_dafny.Set.fromElements(p2, true), _850_v125);
          let _852_v127;
          _852_v127 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_851_v126).length),(_845_v120).f12);
          _852_v127 = (_852_v127).update(p0, (_dafny.MultiSet.fromElements((_this).fm2((_this.f7)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_this.f7).length))], globalState), p1)).IsDisjointFrom((_677_v9).f5));
        } else {
          let _853___mcc_h23 = (_source9).cf44;
          let _854_cf44 = _853___mcc_h23;
          let _arr24 = _this.f7;
          let _index145 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_this.f7).length));
          _arr24[_index145] = true;
          let _855_v128;
          let _nw137 = new _module.C3();
          _nw137.__ctor((_this).f6, (_this).f5);
          _855_v128 = _nw137;
          let _856_v129;
          _856_v129 = _dafny.Map.Empty.slice().updateUnsafe(_855_v128,_this.f7);
          let _857_v130;
          _857_v130 = _dafny.Seq.of(_this.f7, (((_856_v129).contains(_855_v128)) ? ((_856_v129).get(_855_v128)) : (_this.f7)));
          let _858_v131;
          let _init22 = ((_859_v10) => function (_860_i13) {
            return _859_v10;
          })(_678_v10);
          let _nw138 = Array((new BigNumber(5)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw138.length); _i0_22++) {
            _nw138[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _858_v131 = _nw138;
          let _861_v132;
          let _nw139 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _861_v132 = _nw139;
          let _862_v133;
          _862_v133 = _dafny.Map.Empty.slice().updateUnsafe(_858_v131,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_861_v132)).length));
          let _arr25 = _this.f7;
          let _index146 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_this.f7).length));
          let _rhs167 = (_this).f6;
          let _rhs168 = new BigNumber((_dafny.Seq.Concat(_857_v130, _857_v130)).length);
          let _rhs169 = _654_v0;
          let _rhs170 = (((_862_v133).contains(_858_v131)) ? ((_862_v133).get(_858_v131)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ibbdsuctt")).length)));
          let _rhs171 = new BigNumber(635);
          let _lhs135 = _this.f7;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_this.f7).length));
          let _lhs137 = globalState;
          let _lhs138 = globalState;
          let _lhs139 = globalState;
          _lhs135[_lhs136] = _rhs167;
          _lhs137.f0 = _rhs168;
          _654_v0 = _rhs169;
          _lhs138.f0 = _rhs170;
          _lhs139.f0 = _rhs171;
          let _863_v134;
          let _nw140 = Array((new BigNumber(21)).toNumber()).fill(_module.D13.Default());
          _863_v134 = _nw140;
          let _864_v135;
          _864_v135 = _dafny.Seq.of(_863_v134, _863_v134, _863_v134);
          _863_v134 = (_864_v135)[_module.__default.safeIndex(new BigNumber(799), new BigNumber((_864_v135).length))];
          let _865_v136;
          _865_v136 = _dafny.Map.Empty.slice().updateUnsafe(p0,_654_v0);
          let _866_v137;
          _866_v137 = _dafny.Map.Empty.slice().updateUnsafe(_865_v136,p0);
          let _867_v138;
          let _nw141 = new _module.C1();
          _nw141.__ctor(false, _679_v11, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(p0, p0, p0, p0), _678_v10), (_677_v9).f5);
          _867_v138 = _nw141;
          let _868_v139;
          _868_v139 = _module.D12.create_DC35(((_this.f7)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_this.f7).length))]) === ((_867_v138).f12));
          let _869_v141;
          _869_v141 = _dafny.Map.Empty.slice().updateUnsafe(_865_v136,p1);
          let _870_v143;
          _870_v143 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(586), new BigNumber(-57))) {
              let _871_v142 = _compr_39;
              if (((new BigNumber(586)).isLessThanOrEqualTo(_871_v142)) && ((_871_v142).isLessThan(new BigNumber(-57)))) {
                _coll39.push([_module.__default.safeModuloInt(_871_v142, p0),new BigNumber(854)]);
              }
            }
            return _coll39;
          }(),true);
          let _872_v144;
          let _nw142 = new _module.C0();
          _nw142.__ctor(_870_v143);
          _872_v144 = _nw142;
          let _873_v145;
          _873_v145 = _dafny.Set.fromElements(_872_v144, _872_v144, _872_v144, _872_v144);
          let _rhs172 = (_866_v137).Merge((function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_869_v141).Keys.Elements) {
              let _874_v140 = _compr_40;
              if ((_869_v141).contains(_874_v140)) {
                _coll40.push([_874_v140,p0]);
              }
            }
            return _coll40;
          }()).Merge(_866_v137));
          let _rhs173 = _module.__default.safeDivisionInt(p0, (p0).plus((_dafny.ZERO).minus(p0)));
          let _rhs174 = _867_v138;
          let _rhs175 = new BigNumber((((((_867_v138).f12) || (_this.f8)) ? (_873_v145) : (_873_v145))).length);
          let _rhs176 = function (_pat_let26_0) {
            return function (_875_dt__update__tmp_h1) {
              return function (_pat_let27_0) {
                return function (_876_dt__update_hcf55_h0) {
                  return _module.D12.create_DC35(_876_dt__update_hcf55_h0);
                }(_pat_let27_0);
              }(_this.f8);
            }(_pat_let26_0);
          }(_868_v139);
          let _lhs140 = globalState;
          let _lhs141 = globalState;
          _866_v137 = _rhs172;
          _lhs140.f0 = _rhs173;
          _867_v138 = _rhs174;
          _lhs141.f0 = _rhs175;
          _868_v139 = _rhs176;
          (globalState).f0 = p0;
        }
        let _877_v146;
        _877_v146 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _679_v11),p0);
        let _878_v147;
        _878_v147 = _dafny.Seq.of(_877_v146);
        _877_v146 = (((_878_v147)[_module.__default.safeIndex(new BigNumber((_678_v10).length), new BigNumber((_878_v147).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(_679_v11,new BigNumber(777)))).Merge(_877_v146);
        if (((!(_656_v2).equals(function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of _dafny.IntegerRange(new BigNumber(849), new BigNumber(589))) {
            let _879_v148 = _compr_41;
            if (((new BigNumber(849)).isLessThanOrEqualTo(_879_v148)) && ((_879_v148).isLessThan(new BigNumber(589)))) {
              _coll41.push([_module.__default.safeDivisionInt(_879_v148, p0),p0]);
            }
          }
          return _coll41;
        }())) ? (p2) : (((_this).f6) && ((_this).f6)))) {
          let _880_v149;
          _880_v149 = _dafny.Set.fromElements(_679_v11, _679_v11, _679_v11, (_this).fm7(globalState));
          _880_v149 = _880_v149;
          let _881_v150;
          _881_v150 = _dafny.Set.fromElements(p2, _this.f8);
          let _882_v151;
          _882_v151 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(p1, p2, p1, _this.f8)).IsSubsetOf(_881_v150),p0);
          _882_v151 = (_882_v151).update(((false) ? (p2) : (_this.f8)), p0);
          let _nw143 = Array((new BigNumber(27)).toNumber()).fill(false);
          (_this).f7 = _nw143;
          let _883_v152;
          _883_v152 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f8);
          let _884_v153;
          _884_v153 = _module.D9.create_DC27(_883_v152, p1);
          let _885_v154;
          _885_v154 = _dafny.MultiSet.fromElements(_module.D9.create_DC27(_883_v152, _this.f8), _884_v153);
          let _886_v155;
          _886_v155 = _module.D16.create_DC43(_885_v154);
          let _887_v156;
          _887_v156 = _dafny.Seq.of((_dafny.MultiSet.fromElements(_884_v153, _884_v153)).Difference(_885_v154), (_886_v155).dtor_cf68, (_dafny.MultiSet.fromElements(_884_v153, _884_v153, _884_v153)).Difference(_885_v154));
          (globalState).f0 = new BigNumber(((_887_v156)[_module.__default.safeIndex(p0, new BigNumber((_887_v156).length))]).cardinality());
          let _888_v157;
          _888_v157 = _dafny.Map.Empty.slice().updateUnsafe(_656_v2,_this.f8);
          let _889_v158;
          let _nw144 = new _module.C0();
          _nw144.__ctor(_888_v157);
          _889_v158 = _nw144;
          _889_v158 = _889_v158;
        } else {
          (_this).f8 = (_this).f6;
          let _arr26 = _this.f7;
          let _index147 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_this.f7).length));
          _arr26[_index147] = (_dafny.Set.fromElements(p0)).contains(p0);
          let _arr27 = _this.f7;
          let _index148 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_this.f7).length));
          let _rhs177 = (_this).fm6(globalState);
          let _rhs178 = (new BigNumber(122)).plus(p0);
          let _rhs179 = _dafny.Seq.update(_679_v11, _module.__default.safeIndex((_this).fm6(globalState), new BigNumber((_679_v11).length)), _654_v0);
          let _rhs180 = _this.f8;
          let _rhs181 = _this.f8;
          let _lhs142 = globalState;
          let _lhs143 = globalState;
          let _lhs144 = _this;
          let _lhs145 = _this.f7;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_this.f7).length));
          _lhs142.f0 = _rhs177;
          _lhs143.f0 = _rhs178;
          _679_v11 = _rhs179;
          _lhs144.f8 = _rhs180;
          _lhs145[_lhs146] = _rhs181;
          let _890_v159;
          let _out9;
          _out9 = (_830_v115).m9(false, p0, globalState);
          _890_v159 = _out9;
          (globalState).f0 = (_this).fm6(globalState);
          let _891_v160;
          let _init23 = ((_892_p1, _893_v11, _894_p0) => function (_895_i14) {
            return (_895_i14).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_892_p1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_893_v11,_dafny.Set.fromElements(_894_p0, (_dafny.ZERO).minus(_894_p0)))).length))).length));
          })(p1, _679_v11, p0);
          let _nw145 = Array((new BigNumber(10)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw145.length); _i0_23++) {
            _nw145[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _891_v160 = _nw145;
          _891_v160 = _891_v160;
        }
      }
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f7 = [];
      this._f8 = false;
      this._f6 = false;
      this._f5 = _dafny.MultiSet.Empty;
      this._f9 = _dafny.ZERO;
      this._f10 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    set f8(value) {
      let _this = this;
      _this._f8 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f9, f10, f7, f8, f6, f5) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f6 = f6;
      (_this)._f5 = f5;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return ((_this).f9).isLessThanOrEqualTo((_this).f9);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rnpbenu"), _dafny.Seq.UnicodeFromString("ygnckedh")), _dafny.Seq.UnicodeFromString("klfmjbsc"))).length);
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f8;
    };
    fm2(p0, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt((_this).f9, (_this).f9)).isLessThan((_this).f9);
    };
    m2(p0, globalState) {
      let _this = this;
      let _896_v0;
      _896_v0 = new _dafny.CodePoint('r'.codePointAt(0));
      _896_v0 = _896_v0;
      let _897_v1;
      let _init24 = ((_898_p0) => function (_899_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-489)), ((_900_p0) => function (_901_i1) {
          return _900_p0;
        })(_898_p0));
      })(p0);
      let _nw146 = Array((new BigNumber(29)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw146.length); _i0_24++) {
        _nw146[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _897_v1 = _nw146;
      _897_v1 = _897_v1;
      let _902_v2;
      _902_v2 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f10);
      let _903_v3;
      _903_v3 = _dafny.Seq.of(!(_this.f8), (_this).fm1((_this).f6, (_this).f9, new BigNumber((_902_v2).length), globalState), (_this).f6, true);
      let _904_v4;
      _904_v4 = _dafny.Seq.UnicodeFromString("nhj");
      (_this).f8 = ((_this.f8) ? ((new BigNumber((_903_v3).length)).isLessThan((_this).f9)) : (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), function (_905_i2) {
        return _dafny.Seq.UnicodeFromString("khisohk");
      }), _dafny.Seq.of(_904_v4))));
      let _906_v5;
      _906_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this.f8) || (!((_this).f10)),_this.f7);
      _906_v5 = _906_v5;
      if (!(_this.f8)) {
        let _907_v6;
        _907_v6 = _dafny.Seq.of(_this.f7, _this.f7);
        let _908_v7;
        _908_v7 = _dafny.Map.Empty.slice().updateUnsafe((_907_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_this.f8)).length), new BigNumber((_907_v6).length))],_896_v0);
        _896_v0 = (((_908_v7).contains(_this.f7)) ? ((_908_v7).get(_this.f7)) : (p0));
        let _909_v8;
        let _nw147 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
        _909_v8 = _nw147;
        let _910_v9;
        _910_v9 = _dafny.Set.fromElements(new BigNumber((_903_v3).length), new BigNumber((_dafny.Seq.UnicodeFromString("rkohfvwpq")).length));
        let _index149 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_909_v8).length));
        (_909_v8)[_index149] = _910_v9;
        let _911_v10;
        _911_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Set.fromElements((_this).f9, (_this).f9, (_this).f9));
        let _912_v11;
        _912_v11 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_911_v10).contains((_this).f9)) ? ((_911_v10).get((_this).f9)) : (_dafny.Set.fromElements((_this).f9))));
        let _index150 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_909_v8).length));
        (_909_v8)[_index150] = (_910_v9).Intersect((((_912_v11).contains((_this).f6)) ? ((_912_v11).get((_this).f6)) : (_910_v9)));
        if (_this.f8) {
          (globalState).f0 = (_this).f9;
          let _index151 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_897_v1).length));
          (_897_v1)[_index151] = _904_v4;
          let _913_v12;
          _913_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(-811));
          let _914_v13;
          _914_v13 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jxhfdgy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), ((_915_p0) => function (_916_i3) {
            return _915_p0;
          })(p0))), _module.__default.safeIndex(new BigNumber((_904_v4).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jxhfdgy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), ((_917_p0) => function (_918_i3) {
            return _917_p0;
          })(p0)))).length)), p0));
          let _919_v15;
          _919_v15 = _dafny.Map.Empty.slice().updateUnsafe(_904_v4,_this.f8);
          let _920_v16;
          _920_v16 = _module.D0.create_DC1();
          let _921_v17;
          _921_v17 = _dafny.Set.fromElements(_920_v16);
          let _922_v18;
          _922_v18 = _dafny.Set.fromElements((((_902_v2).contains((_this).f10)) ? ((_902_v2).get((_this).f10)) : ((((_902_v2).contains((_this).f10)) ? ((_902_v2).get((_this).f10)) : (_this.f8)))), (_this).f10, _this.f8, _this.f8);
          let _index152 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_897_v1).length));
          let _rhs182 = _904_v4;
          let _rhs183 = function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of _dafny.IntegerRange(new BigNumber(-603), new BigNumber(543))) {
              let _923_v14 = _compr_42;
              if (((new BigNumber(-603)).isLessThanOrEqualTo(_923_v14)) && ((_923_v14).isLessThan(new BigNumber(543)))) {
                _coll42.push([_module.__default.safeDivisionInt(_923_v14, (_this).f9),(_module.D0.create_DC0(new BigNumber(22))).dtor_cf0]);
              }
            }
            return _coll42;
          }();
          let _rhs184 = (_this.f8) === ((_dafny.Set.fromElements(new BigNumber((_902_v2).length), new BigNumber(-810), (_this).f9)).IsDisjointFrom(_dafny.Set.fromElements((_this).fm4(new BigNumber((_dafny.Seq.of((_this).f9, (_this).f9, new BigNumber(-857), (_this).f9, (_this).f9)).length), (_this).f9, _919_v15, new BigNumber((_921_v17).length), globalState))));
          let _rhs185 = (_922_v18).IsProperSubsetOf(_922_v18);
          let _rhs186 = _dafny.Seq.Concat(_914_v13, _module.__default.fm5(globalState));
          let _lhs147 = _897_v1;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_897_v1).length));
          let _lhs149 = _this;
          let _lhs150 = _this;
          _lhs147[_lhs148] = _rhs182;
          _913_v12 = _rhs183;
          _lhs149.f8 = _rhs184;
          _lhs150.f8 = _rhs185;
          _914_v13 = _rhs186;
          let _924_v19;
          _924_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(172),(_this).f9));
          let _925_v20;
          _925_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_913_v12).length),(((_924_v19).contains((_this).f6)) ? ((_924_v19).get((_this).f6)) : (_913_v12)));
          let _926_v21;
          _926_v21 = _dafny.Seq.of((((_925_v20).contains((_this).f9)) ? ((_925_v20).get((_this).f9)) : (_913_v12)));
          let _927_v22;
          _927_v22 = _dafny.MultiSet.fromElements((_this).f9);
          let _928_v23;
          _928_v23 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,_896_v0)).length));
          let _929_v24;
          _929_v24 = _dafny.Seq.of(_928_v23);
          let _930_v25;
          _930_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Seq.update(_928_v23, _module.__default.safeIndex((_this).f9, new BigNumber((_928_v23).length)), (_this).f9));
          let _931_v26;
          let _nw148 = Array((new BigNumber(24)).toNumber());
          _nw148[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f9);
          _nw148[(_dafny.ONE).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(2)).toNumber()] = ((_this).f9).multipliedBy((_this).f9);
          _nw148[(new BigNumber(3)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(4)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(6)).toNumber()] = ((_this).f9).minus((_this).f9);
          _nw148[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus((_this).f9)).multipliedBy((_this).f9);
          _nw148[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_926_v21)[_module.__default.safeIndex(new BigNumber((_927_v22).cardinality()), new BigNumber((_926_v21).length))]).length));
          _nw148[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).f9);
          _nw148[(new BigNumber(10)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(11)).toNumber()] = new BigNumber(-585);
          _nw148[(new BigNumber(12)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(13)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(14)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(15)).toNumber()] = new BigNumber(((_929_v24)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_904_v4, _module.__default.safeIndex((_this).f9, new BigNumber((_904_v4).length)), p0)).length), new BigNumber((_929_v24).length))]).length);
          _nw148[(new BigNumber(16)).toNumber()] = ((_this).f9).plus(new BigNumber((_930_v25).length));
          _nw148[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(470), new BigNumber(97));
          _nw148[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_904_v4, _dafny.Seq.UnicodeFromString("vhffqpp")), _module.__default.safeIndex((_this).f9, new BigNumber((_dafny.Seq.Concat(_904_v4, _dafny.Seq.UnicodeFromString("vhffqpp"))).length)), p0)).length);
          _nw148[(new BigNumber(19)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(98), (_this).f9);
          _nw148[(new BigNumber(21)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(22)).toNumber()] = (_this).f9;
          _nw148[(new BigNumber(23)).toNumber()] = new BigNumber(-674);
          _931_v26 = _nw148;
          let _index153 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_931_v26).length));
          (_931_v26)[_index153] = (_dafny.ZERO).minus((_this).f9);
          let _932_v27;
          _932_v27 = _module.D0.create_DC0(new BigNumber((_903_v3).length));
          let _index154 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_931_v26).length));
          (_931_v26)[_index154] = (_932_v27).dtor_cf0;
          (globalState).f0 = (_this).f9;
          let _933_v28;
          let _nw149 = new _module.C4();
          _nw149.__ctor(_this.f7, (_this).f10, !(_this.f8), (_this).f5);
          _933_v28 = _nw149;
          let _934_v29;
          let _nw150 = Array((new BigNumber(10)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _933_v28;
          _nw150[(_dafny.ONE).toNumber()] = _933_v28;
          _nw150[(new BigNumber(2)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(3)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(4)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(5)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(6)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(7)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(8)).toNumber()] = _933_v28;
          _nw150[(new BigNumber(9)).toNumber()] = _933_v28;
          _934_v29 = _nw150;
          let _index155 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_934_v29).length));
          (_934_v29)[_index155] = _933_v28;
          let _index156 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_934_v29).length));
          let _nw151 = new _module.C4();
          _nw151.__ctor(_933_v28.f7, ((_933_v28).f5).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.update(_903_v3, _module.__default.safeIndex((_931_v26)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_931_v26).length))], new BigNumber((_903_v3).length)), (_933_v28).f6))), _933_v28.f8, (_dafny.MultiSet.fromElements(_this.f8, (_this).f6)).Difference((_this).f5));
          (_934_v29)[_index156] = _nw151;
        } else {
          let _935_v30;
          let _nw152 = Array((new BigNumber(14)).toNumber()).fill([]);
          _935_v30 = _nw152;
          _935_v30 = _935_v30;
          (_this).f8 = false;
          (_this).f8 = ((_this).f5).IsDisjointFrom((_this).f5);
          let _arr28 = _this.f7;
          let _index157 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_this.f7).length));
          _arr28[_index157] = (_this).f6;
          let _arr29 = _this.f7;
          let _index158 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_this.f7).length));
          let _rhs187 = _dafny.Seq.UnicodeFromString("ooxmqc");
          let _rhs188 = (_this).f6;
          let _rhs189 = _this.f8;
          let _rhs190 = (_this).f9;
          let _rhs191 = _896_v0;
          let _lhs151 = _this;
          let _lhs152 = _this.f7;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_this.f7).length));
          let _lhs154 = globalState;
          _904_v4 = _rhs187;
          _lhs151.f8 = _rhs188;
          _lhs152[_lhs153] = _rhs189;
          _lhs154.f0 = _rhs190;
          _896_v0 = _rhs191;
          let _arr30 = _this.f7;
          let _index159 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_this.f7).length));
          _arr30[_index159] = ((_this).f5).IsProperSubsetOf(((_this).f5).update(_this.f8, _module.__default.abs((_dafny.ZERO).minus((_this).f9))));
        }
        (globalState).f0 = (_this).f9;
        let _936_v31;
        let _init25 = function (_937_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f10);
        };
        let _nw153 = Array((new BigNumber(29)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw153.length); _i0_25++) {
          _nw153[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _936_v31 = _nw153;
        let _rhs192 = (_this).f9;
        let _rhs193 = _936_v31;
        let _lhs155 = globalState;
        _lhs155.f0 = _rhs192;
        _936_v31 = _rhs193;
      } else {
        (_this).f8 = false;
        let _938_v32;
        let _nw154 = new _module.C3();
        _nw154.__ctor(((((_902_v2).contains((_this).f6)) ? ((_902_v2).get((_this).f6)) : ((_this).f10))) === (_this.f8), (_this).f5);
        _938_v32 = _nw154;
        (_this).f8 = !(_this.f8) || ((_this).f10);
        let _939_v33;
        let _nw155 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _939_v33 = _nw155;
        let _940_v34;
        _940_v34 = _dafny.Seq.of(_this.f7);
        let _941_v35;
        _941_v35 = _dafny.Set.fromElements((_this).f9);
        let _942_v36;
        _942_v36 = _dafny.MultiSet.fromElements(_this.f7, _this.f7, (_940_v34)[_module.__default.safeIndex(new BigNumber((_941_v35).length), new BigNumber((_940_v34).length))], _this.f7, _this.f7);
        let _943_v37;
        _943_v37 = _dafny.Map.Empty.slice().updateUnsafe(_939_v33,(_dafny.ZERO).minus(((_this).f9).multipliedBy((((_942_v36).contains(_this.f7)) ? ((_942_v36).get(_this.f7)) : ((_this).f9)))));
        let _944_v38;
        _944_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(289));
        let _945_v39;
        let _nw156 = new _module.C1();
        _nw156.__ctor(_this.f8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), ((_946_p0) => function (_947_i5) {
          return _946_p0;
        })(p0)), (_938_v32).fm3(globalState), _module.__default.fm14(_944_v38, _this.f8, (_this).f9, (_this).f10, globalState));
        _945_v39 = _nw156;
        let _948_v40;
        _948_v40 = _dafny.MultiSet.fromElements(_945_v39);
        _943_v37 = (_943_v37).update(_939_v33, new BigNumber((_948_v40).cardinality()));
        let _949_v41;
        _949_v41 = _dafny.MultiSet.fromElements((_this).f9);
        let _950_v42;
        _950_v42 = _module.D2.create_DC7(p0, _944_v38, (_this).f9);
        let _951_v43;
        _951_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-980)), ((_952_p0) => function (_953_i6) {
          return _952_p0;
        })(p0)),_this.f8);
        (globalState).f0 = (_945_v39).fm4((((_949_v41).contains(new BigNumber(608))) ? ((_949_v41).get(new BigNumber(608))) : ((_this).f9)), (_this).f9, (_module.__default.fm13((_this).f9, _950_v42, (_this).f10, new BigNumber((_903_v3).length), globalState)).Merge(_951_v43), ((_this).f9).multipliedBy((_dafny.ZERO).minus((_this).f9)), globalState);
      }
      let _954_v44;
      _954_v44 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(588),(_this).f6)).length)),(_this).f9);
      let _955_v45;
      _955_v45 = _module.D2.create_DC7(_896_v0, _954_v44, (_this).f9);
      let _956_v46;
      _956_v46 = _module.D12.create_DC36(new BigNumber(321), (_this).f9);
      let _957_v47;
      _957_v47 = _module.D15.create_DC41(_module.__default.fm13((_this).f9, _955_v45, (_this).f10, (function (_pat_let28_0) {
  return function (_958_dt__update__tmp_h0) {
    return function (_pat_let29_0) {
      return function (_959_dt__update_hcf57_h0) {
        return _module.D12.create_DC36((_958_dt__update__tmp_h0).dtor_cf56, _959_dt__update_hcf57_h0);
      }(_pat_let29_0);
    }((_this).f9);
  }(_pat_let28_0);
}(_956_v46)).dtor_cf57, globalState));
      let _960_i7;
      _960_i7 = _dafny.ZERO;
      L9: {
        while (function (_source10) {
          if (_source10.is_DC42) {
            let _963___mcc_h0 = (_source10).cf67;
            let _964_cf67 = _963___mcc_h0;
            return (new BigNumber((_dafny.Set.fromElements((_this).f10, _this.f8)).length)).isEqualTo(new BigNumber((_dafny.Set.fromElements((_this).f10)).length));
          } else {
            let _965___mcc_h1 = (_source10).cf66;
            let _966_cf66 = _965___mcc_h1;
            return (_this).f6;
          }
        }(_957_v47)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_960_i7)) {
              break L9;
            }
            _960_i7 = (_960_i7).plus(_dafny.ONE);
            let _961_v48;
            let _nw157 = new _module.C1();
            _nw157.__ctor(((_this.f8) ? (_this.f8) : (_this.f8)), _dafny.Seq.Concat(_904_v4, _dafny.Seq.UnicodeFromString("ceqq")), (_this).f6, (_this).f5);
            _961_v48 = _nw157;
            let _962_v49;
            let _nw158 = new _module.C4();
            _nw158.__ctor(_this.f7, _this.f8, (_this).f10, ((_this).f5).Union(_dafny.MultiSet.fromElements(!((_this).f6))));
            _962_v49 = _nw158;
            (_this).f8 = (_this).f10;
            let _arr31 = _this.f7;
            let _index160 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_this.f7).length));
            _arr31[_index160] = !((_962_v49).fm3(globalState));
            let _arr32 = _this.f7;
            let _index161 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_this.f7).length));
            _arr32[_index161] = !_dafny.Seq.contains(_903_v3, (((_this).f10) ? ((_961_v48).f12) : ((_this).f6)));
          }
        }
      }
      return;
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _arr33 = _this.f7;
      let _index162 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length));
      _arr33[_index162] = !((_this).f10);
      let _967_v0;
      _967_v0 = _dafny.Seq.of((_this).f6, (_this).f10);
      let _968_v1;
      _968_v1 = _module.D10.create_DC29(_dafny.MultiSet.FromArray(_dafny.Seq.of(_967_v0)));
      let _969_v2;
      _969_v2 = _dafny.MultiSet.fromElements(_968_v1, _968_v1);
      let _arr34 = _this.f7;
      let _index163 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length));
      _arr34[_index163] = !(((_969_v2).Intersect(_969_v2)).IsSubsetOf(_969_v2));
      let _970_v3;
      _970_v3 = _module.D3.create_DC10(!(_this.f8), (_this).f9);
      let _arr35 = _this.f7;
      let _index164 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length));
      _arr35[_index164] = !(!(function (_source11) {
        if (_source11.is_DC10) {
          let _971___mcc_h0 = (_source11).cf11;
          let _972___mcc_h1 = (_source11).cf12;
          let _973_cf12 = _972___mcc_h1;
          let _974_cf11 = _971___mcc_h0;
          return (_this).f10;
        } else if (_source11.is_DC11) {
          let _975___mcc_h2 = (_source11).cf13;
          let _976_cf13 = _975___mcc_h2;
          return (_this).f10;
        } else {
          let _977___mcc_h3 = (_source11).cf10;
          let _978_cf10 = _977___mcc_h3;
          return (_this).f6;
        }
      }(_970_v3)));
      let _979_v4;
      _979_v4 = _dafny.Seq.UnicodeFromString("yie");
      let _980_v5;
      _980_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_979_v4, _dafny.Seq.UnicodeFromString("kynauahc")),(_this).f6);
      _980_v5 = (_980_v5).update(_979_v4, !(_this.f8));
      let _981_v6;
      _981_v6 = _dafny.MultiSet.fromElements((_this).f9);
      let _982_v7;
      let _nw159 = Array((new BigNumber(17)).toNumber());
      _nw159[(_dafny.ZERO).toNumber()] = _967_v0;
      _nw159[(_dafny.ONE).toNumber()] = _967_v0;
      _nw159[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_967_v0, _967_v0);
      _nw159[(new BigNumber(3)).toNumber()] = _967_v0;
      _nw159[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(!(true), (_this.f7)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length))]), _967_v0);
      _nw159[(new BigNumber(5)).toNumber()] = _module.__default.fm23((((_981_v6).contains((_this).f9)) ? ((_981_v6).get((_this).f9)) : ((_this).f9)), globalState);
      _nw159[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_967_v0, _967_v0);
      _nw159[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_967_v0, _module.__default.safeIndex((_this).f9, new BigNumber((_967_v0).length)), (_this.f7)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length))]);
      _nw159[(new BigNumber(8)).toNumber()] = _967_v0;
      _nw159[(new BigNumber(9)).toNumber()] = _dafny.Seq.of((_this).f6);
      _nw159[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_967_v0, _967_v0);
      _nw159[(new BigNumber(11)).toNumber()] = _967_v0;
      _nw159[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_967_v0, _967_v0);
      _nw159[(new BigNumber(13)).toNumber()] = _module.__default.fm23((_this).f9, globalState);
      _nw159[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_967_v0, _dafny.Seq.update(_967_v0, _module.__default.safeIndex((_this).f9, new BigNumber((_967_v0).length)), false));
      _nw159[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f8), _dafny.Seq.update(_dafny.Seq.of((_this).f10, !((_this.f7)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length))])), _module.__default.safeIndex((_this).f9, new BigNumber((_dafny.Seq.of((_this).f10, !((_this.f7)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length))]))).length)), (_this).f10));
      _nw159[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm23(new BigNumber(712), globalState), _967_v0);
      _982_v7 = _nw159;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_982_v7).length))) {
        let _983_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_983_i0)) && ((_983_i0).isLessThan(new BigNumber((_982_v7).length))))) {
          (_982_v7)[(_983_i0)] = _967_v0;
        }
      }
      let _984_v8;
      _984_v8 = _module.D0.create_DC1();
      _984_v8 = _984_v8;
      let _985_v9;
      _985_v9 = new _dafny.CodePoint('b'.codePointAt(0));
      let _986_v10;
      _986_v10 = _module.D10.create_DC31(_dafny.Seq.of((_this).f5), _985_v9);
      (_this).f8 = function (_source12) {
        if (_source12.is_DC30) {
          let _987___mcc_h4 = (_source12).cf46;
          let _988___mcc_h5 = (_source12).cf47;
          let _989___mcc_h6 = (_source12).cf48;
          let _990_cf48 = _989___mcc_h6;
          let _991_cf47 = _988___mcc_h5;
          let _992_cf46 = _987___mcc_h4;
          return (_this.f7)[_module.__default.safeIndex(new BigNumber(238), new BigNumber((_this.f7).length))];
        } else if (_source12.is_DC31) {
          let _993___mcc_h7 = (_source12).cf49;
          let _994___mcc_h8 = (_source12).cf50;
          let _995_cf50 = _994___mcc_h8;
          let _996_cf49 = _993___mcc_h7;
          return (_this).f10;
        } else {
          let _997___mcc_h9 = (_source12).cf45;
          let _998_cf45 = _997___mcc_h9;
          return (_this).f6;
        }
      }(_986_v10);
      let _999_v11;
      _999_v11 = _module.D12.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_985_v9,_this.f8));
      let _1000_v12;
      _1000_v12 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f9), new BigNumber(32), _module.__default.safeModuloInt(new BigNumber(39), (_this).f9), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_999_v11)).length), new BigNumber((_979_v4).length));
      r0 = _1000_v12;
      let _1001_v13;
      let _init26 = function (_1002_i1) {
        return _module.__default.safeModuloInt(_1002_i1, (_this).f9);
      };
      let _nw160 = Array((new BigNumber(4)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw160.length); _i0_26++) {
        _nw160[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _1001_v13 = _nw160;
      let _1003_v14;
      _1003_v14 = _dafny.MultiSet.fromElements(_1001_v13, _1001_v13, _1001_v13);
      let _1004_v15;
      _1004_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_1003_v14).update(_1001_v13, _module.__default.abs((_this).f9)));
      r1 = (_this).fm4(new BigNumber(217), (_this).f9, _980_v5, new BigNumber((((((_1004_v15).contains((_this).f9)) ? ((_1004_v15).get((_this).f9)) : (_1003_v14))).Union(_1003_v14)).cardinality()), globalState);
      r2 = (_this).f9;
      return [r0, r1, r2];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _1005_i0;
      _1005_i0 = _dafny.ZERO;
      L10: {
        while ((_this).f10) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1005_i0)) {
              break L10;
            }
            _1005_i0 = (_1005_i0).plus(_dafny.ONE);
            let _1006_v0;
            let _init27 = function (_1007_i1) {
              return (_1007_i1).multipliedBy((_this).f9);
            };
            let _nw161 = Array((new BigNumber(21)).toNumber());
            for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw161.length); _i0_27++) {
              _nw161[_i0_27] = _init27(new BigNumber(_i0_27));
            }
            _1006_v0 = _nw161;
            let _1008_v1;
            _1008_v1 = _module.D3.create_DC9(_1006_v0);
            let _1009_v2;
            _1009_v2 = _dafny.Seq.of(p0);
            let _1010_v3;
            _1010_v3 = _dafny.Seq.of(!((_this).f6));
            let _1011_v4;
            _1011_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(487),new BigNumber(828));
            let _1012_v5;
            _1012_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1011_v4,(_this).f6);
            let _1013_v6;
            let _nw162 = new _module.C0();
            _nw162.__ctor(_1012_v5);
            _1013_v6 = _nw162;
            let _1014_v7;
            _1014_v7 = _dafny.Set.fromElements((_this).f6, _this.f8, (_this).f6, (_this).f6, p2);
            let _rhs194 = (p0).multipliedBy((_module.D5.create_DC16(_1008_v1, new BigNumber((_1009_v2).length), new BigNumber(40), new BigNumber((_1010_v3).length), _1013_v6)).dtor_cf20);
            let _rhs195 = _dafny.Seq.contains(_dafny.Seq.of(p0, new BigNumber(-625)), (((_1011_v4).contains(p0)) ? ((_1011_v4).get(p0)) : (new BigNumber((_1014_v7).length))));
            let _rhs196 = (_1009_v2)[_module.__default.safeIndex((_this).f9, new BigNumber((_1009_v2).length))];
            let _rhs197 = (((((_this).f5).contains(p1)) ? (((_this).f5).get(p1)) : ((_this).f9))).isLessThan(p0);
            let _lhs156 = globalState;
            let _lhs157 = _this;
            let _lhs158 = globalState;
            let _lhs159 = _this;
            _lhs156.f0 = _rhs194;
            _lhs157.f8 = _rhs195;
            _lhs158.f0 = _rhs196;
            _lhs159.f8 = _rhs197;
            (globalState).f0 = (_this).f9;
            let _index165 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_1006_v0).length));
            (_1006_v0)[_index165] = (p0).minus(p0);
            let _index166 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_1006_v0).length));
            (_1006_v0)[_index166] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(p0))));
            let _1015_v8;
            _1015_v8 = _dafny.Seq.UnicodeFromString("bobaat");
            _1015_v8 = ((false) ? (_1015_v8) : (_dafny.Seq.Concat(_1015_v8, _1015_v8)));
          }
        }
      }
      let _1016_v9;
      _1016_v9 = _dafny.Set.fromElements((_this).f6);
      let _1017_v10;
      _1017_v10 = _dafny.Seq.UnicodeFromString("rhu");
      let _1018_v11;
      _1018_v11 = _module.D5.create_DC15(_1017_v10);
      let _1019_v12;
      _1019_v12 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v11).dtor_cf18,(_this).f10);
      (globalState).f0 = ((_this).fm4(new BigNumber((_1016_v9).length), (_this).f9, _1019_v12, (_dafny.ZERO).minus((_this).f9), globalState)).multipliedBy(p0);
      let _pat_let_tv20 = p0;
      let _pat_let_tv21 = p0;
      let _pat_let_tv22 = p1;
      let _pat_let_tv23 = p1;
      let _pat_let_tv24 = p0;
      (globalState).f0 = function (_source13) {
        if (_source13.is_DC20) {
          let _1020___mcc_h0 = (_source13).cf29;
          let _1021_cf29 = _1020___mcc_h0;
          return _module.__default.safeModuloInt(_pat_let_tv20, _pat_let_tv21);
        } else if (_source13.is_DC21) {
          let _1022___mcc_h1 = (_source13).cf30;
          let _1023___mcc_h2 = (_source13).cf31;
          let _1024_cf31 = _1023___mcc_h2;
          let _1025_cf30 = _1022___mcc_h1;
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_pat_let_tv22)).length);
        } else if (_source13.is_DC19) {
          let _1026___mcc_h3 = (_source13).cf28;
          let _1027_cf28 = _1026___mcc_h3;
          return (_this).f9;
        } else {
          let _1028___mcc_h4 = (_source13).cf32;
          let _1029_cf32 = _1028___mcc_h4;
          return _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(_pat_let_tv23, false)).cardinality()), _pat_let_tv24);
        }
      }(_module.D7.create_DC22(_module.D7.create_DC21(p0, new BigNumber(734))));
      let _arr36 = _this.f7;
      let _index167 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
      _arr36[_index167] = p1;
      let _arr37 = _this.f7;
      let _index168 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
      _arr37[_index168] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ibfacfd"), _1017_v10);
      let _1030_v13;
      _1030_v13 = _module.D7.create_DC21(p0, new BigNumber(383));
      let _1031_v14;
      _1031_v14 = _dafny.MultiSet.fromElements(_1030_v13);
      let _1032_v15;
      _1032_v15 = _dafny.Seq.of((_this).f9, new BigNumber((_1031_v14).cardinality()));
      let _1033_v16;
      _1033_v16 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(false) || (!((_this).f10)));
      let _arr38 = _this.f7;
      let _index169 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
      let _rhs198 = _dafny.Seq.Concat(_1032_v15, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-447)), ((_1034_p0) => function (_1035_i2) {
        return _1034_p0;
      })(p0)));
      let _rhs199 = !((((_1033_v16).contains((_1017_v10)[_module.__default.safeIndex((_this).fm4(p0, (_dafny.ZERO).minus((_this).f9), _1019_v12, new BigNumber(464), globalState), new BigNumber((_1017_v10).length))])) ? ((_1033_v16).get((_1017_v10)[_module.__default.safeIndex((_this).fm4(p0, (_dafny.ZERO).minus((_this).f9), _1019_v12, new BigNumber(464), globalState), new BigNumber((_1017_v10).length))])) : ((p0).isEqualTo(p0))));
      let _lhs160 = _this.f7;
      let _lhs161 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
      _1032_v15 = _rhs198;
      _lhs160[_lhs161] = _rhs199;
      let _1036_v17;
      _1036_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber((_1017_v10).length));
      let _1037_v18;
      _1037_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_1036_v17).contains((_this).f9)) ? ((_1036_v17).get((_this).f9)) : ((_this).f9)));
      let _1038_i3;
      _1038_i3 = _dafny.ZERO;
      L11: {
        while ((new BigNumber(((_module.__default.fm14(_1037_v18, false, new BigNumber(692), (_this.f7)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length))], globalState)).update(p1, _module.__default.abs((_1032_v15)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f9), new BigNumber((_1032_v15).length))]))).cardinality())).isLessThanOrEqualTo((new BigNumber((_1032_v15).length)).plus(new BigNumber((_1017_v10).length)))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1038_i3)) {
              break L11;
            }
            _1038_i3 = (_1038_i3).plus(_dafny.ONE);
            _1016_v9 = _dafny.Set.fromElements((_this).f6, p2);
            let _arr39 = _this.f7;
            let _index170 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
            _arr39[_index170] = (_this).f6;
            let _1039_v19;
            let _nw163 = new _module.C2();
            _nw163.__ctor(_dafny.MultiSet.fromElements((_this.f7)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length))], (_this.f7)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length))], (_this.f7)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length))], true));
            _1039_v19 = _nw163;
            let _1040_v20;
            _1040_v20 = _dafny.MultiSet.fromElements(_1039_v19, _1039_v19, _1039_v19);
            let _1041_v21;
            _1041_v21 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f8);
            let _1042_v22;
            _1042_v22 = _dafny.Seq.of(_1041_v21);
            (_this).f8 = (_this).fm1((_this).f6, (((_1040_v20).contains(_1039_v19)) ? ((_1040_v20).get(_1039_v19)) : (new BigNumber((_1042_v22).length))), (_this).f9, globalState);
            let _arr40 = _this.f7;
            let _index171 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_this.f7).length));
            _arr40[_index171] = true;
          }
        }
      }
      return;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let _1043_v0;
      _1043_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f9);
      _1043_v0 = (_1043_v0).update((_this).fm2((_this).f10, globalState), (_this).f9);
      let _1044_v1;
      _1044_v1 = _dafny.Seq.UnicodeFromString("nttgnqle");
      let _1045_v2;
      _1045_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1044_v1,true);
      let _1046_v3;
      _1046_v3 = _module.D16.create_DC44((_this).fm2((_this).f6, globalState), (_this).fm3(globalState), (_this).fm4((_this).f9, (_this).f9, _1045_v2, (_this).f9, globalState));
      let _source14 = function (_source15) {
        if (_source15.is_DC44) {
          let _1047___mcc_h4 = (_source15).cf69;
          let _1048___mcc_h5 = (_source15).cf70;
          let _1049___mcc_h6 = (_source15).cf71;
          let _1050_cf71 = _1049___mcc_h6;
          let _1051_cf70 = _1048___mcc_h5;
          let _1052_cf69 = _1047___mcc_h4;
          return _module.D6.create_DC17(_dafny.Seq.of(!(_1051_cf70), _1051_cf70, true));
        } else if (_source15.is_DC43) {
          let _1053___mcc_h7 = (_source15).cf68;
          let _1054_cf68 = _1053___mcc_h7;
          return _module.D6.create_DC17(_dafny.Seq.of(_this.f8));
        } else {
          let _1055___mcc_h8 = (_source15).cf72;
          let _1056_cf72 = _1055___mcc_h8;
          return _module.D6.create_DC17(_dafny.Seq.of(_this.f8));
        }
      }(_1046_v3);
      if (_source14.is_DC18) {
        let _1057___mcc_h0 = (_source14).cf25;
        let _1058___mcc_h1 = (_source14).cf26;
        let _1059___mcc_h2 = (_source14).cf27;
        let _1060_cf27 = _1059___mcc_h2;
        let _1061_cf26 = _1058___mcc_h1;
        let _1062_cf25 = _1057___mcc_h0;
        let _1063_v4;
        let _nw164 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1063_v4 = _nw164;
        let _1064_v5;
        _1064_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1063_v4);
        _1064_v5 = (_1064_v5).update(false, _1063_v4);
        let _1065_v6;
        _1065_v6 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1066_v7;
        let _nw165 = Array((new BigNumber(8)).toNumber());
        _nw165[(_dafny.ZERO).toNumber()] = _1065_v6;
        _nw165[(_dafny.ONE).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(2)).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(3)).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(4)).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(5)).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(6)).toNumber()] = _1065_v6;
        _nw165[(new BigNumber(7)).toNumber()] = _1065_v6;
        _1066_v7 = _nw165;
        let _1067_v8;
        let _nw166 = Array((new BigNumber(13)).toNumber());
        _nw166[(_dafny.ZERO).toNumber()] = _1066_v7;
        _nw166[(_dafny.ONE).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(2)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(3)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(4)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(5)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(6)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(7)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(8)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(9)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(10)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(11)).toNumber()] = _1066_v7;
        _nw166[(new BigNumber(12)).toNumber()] = _1066_v7;
        _1067_v8 = _nw166;
        let _1068_v9;
        _1068_v9 = _module.D4.create_DC12(_1067_v8);
        let _1069_v10;
        _1069_v10 = _module.D4.create_DC14(_1068_v9);
        let _1070_v11;
        _1070_v11 = _module.D4.create_DC14(_1069_v10);
        let _pat_let_tv25 = _1069_v10;
        let _pat_let_tv26 = _1069_v10;
        let _1071_v12;
        let _nw167 = Array((new BigNumber(20)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = _1070_v11;
        _nw167[(_dafny.ONE).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(2)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(3)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(4)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(5)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(6)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(7)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(8)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(9)).toNumber()] = _module.D4.create_DC14(_1069_v10);
        _nw167[(new BigNumber(10)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(11)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(12)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(13)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(14)).toNumber()] = _module.D4.create_DC14((_1070_v11).dtor_cf17);
        _nw167[(new BigNumber(15)).toNumber()] = function (_pat_let30_0) {
          return function (_1072_dt__update__tmp_h0) {
            return function (_pat_let31_0) {
              return function (_1073_dt__update_hcf17_h0) {
                return _module.D4.create_DC14(_1073_dt__update_hcf17_h0);
              }(_pat_let31_0);
            }(_pat_let_tv25);
          }(_pat_let30_0);
        }(_module.D4.create_DC14(_1069_v10));
        _nw167[(new BigNumber(16)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(17)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(18)).toNumber()] = _1070_v11;
        _nw167[(new BigNumber(19)).toNumber()] = function (_pat_let32_0) {
          return function (_1074_dt__update__tmp_h1) {
            return function (_pat_let33_0) {
              return function (_1075_dt__update_hcf17_h1) {
                return _module.D4.create_DC14(_1075_dt__update_hcf17_h1);
              }(_pat_let33_0);
            }(_pat_let_tv26);
          }(_pat_let32_0);
        }(_1070_v11);
        _1071_v12 = _nw167;
        let _index172 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1071_v12).length));
        (_1071_v12)[_index172] = _1070_v11;
        let _index173 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1071_v12).length));
        (_1071_v12)[_index173] = _1070_v11;
        let _1076_v13;
        let _nw168 = new _module.C3();
        _nw168.__ctor(((_this).f9).isLessThan(new BigNumber(984)), (_this).f5);
        _1076_v13 = _nw168;
        let _1077_v14;
        let _nw169 = new _module.C3();
        _nw169.__ctor((_this).f6, (_this).f5);
        _1077_v14 = _nw169;
      } else {
        let _1078___mcc_h3 = (_source14).cf24;
        let _1079_cf24 = _1078___mcc_h3;
        let _1080_v15;
        _1080_v15 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f8);
        let _1081_v16;
        let _nw170 = Array((new BigNumber(14)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = new BigNumber((_1080_v15).length);
        _nw170[(_dafny.ONE).toNumber()] = ((_this).f9).multipliedBy((_this).f9);
        _nw170[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(3)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(6)).toNumber()] = new BigNumber((p1).length);
        _nw170[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(((_this).f9).multipliedBy((_dafny.ZERO).minus((_this).f9)));
        _nw170[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw170[(new BigNumber(12)).toNumber()] = new BigNumber(635);
        _nw170[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_this).f9);
        _1081_v16 = _nw170;
        let _index174 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
        (_1081_v16)[_index174] = new BigNumber(600);
        let _index175 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
        (_1081_v16)[_index175] = (_this).f9;
        let _1082_v17;
        _1082_v17 = _dafny.Set.fromElements((_this).f10);
        let _1083_v18;
        _1083_v18 = _module.D0.create_DC0(new BigNumber((_1082_v17).length));
        let _1084_v19;
        _1084_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v18,(_1044_v1)[_module.__default.safeIndex((_this).f9, new BigNumber((_1044_v1).length))]);
        let _1085_v20;
        _1085_v20 = new _dafny.CodePoint('k'.codePointAt(0));
        _1084_v19 = (_1084_v19).update(((_this.f8) ? (_module.D0.create_DC0((_this).f9)) : (_1083_v18)), _1085_v20);
        let _1086_v21;
        _1086_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,!(p0));
        if (!((((((_1086_v21).contains((_this).f9)) ? ((_1086_v21).get((_this).f9)) : (_this.f8))) ? (true) : (_this.f8)))) {
          let _index176 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          (_1081_v16)[_index176] = (_this).f9;
          let _index177 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          let _rhs200 = (true) && ((_this).f10);
          let _rhs201 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1044_v1, _1044_v1), _dafny.Seq.Concat(_1044_v1, _1044_v1));
          let _rhs202 = (_this).f6;
          let _rhs203 = (_this).f6;
          let _rhs204 = (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))];
          let _lhs162 = _this;
          let _lhs163 = _this;
          let _lhs164 = _this;
          let _lhs165 = _1081_v16;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          _lhs162.f8 = _rhs200;
          _1044_v1 = _rhs201;
          _lhs163.f8 = _rhs202;
          _lhs164.f8 = _rhs203;
          _lhs165[_lhs166] = _rhs204;
          let _index178 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          let _rhs205 = _1044_v1;
          let _rhs206 = (_this).fm4((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], _1045_v2, (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], globalState);
          let _rhs207 = (((_1079_cf24)[_module.__default.safeIndex((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], new BigNumber((_1079_cf24).length))]) ? (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], new BigNumber((p1).length)), new BigNumber(-872))).length), (_this).fm4(new BigNumber((_1080_v15).length), new BigNumber((_1044_v1).length), _1045_v2, new BigNumber((_1079_cf24).length), globalState))) : ((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]));
          let _rhs208 = (!(_this.f8) || ((_this).f6)) || (true);
          let _rhs209 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_1087_v16, _1088_v0) => function (_1089_i0) {
            return (_dafny.Set.fromElements((_1087_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1087_v16).length))])).Union(function () {
              let _coll43 = new _dafny.Set();
              for (const _compr_43 of _dafny.IntegerRange(new BigNumber(314), new BigNumber(504))) {
                let _1090_v22 = _compr_43;
                if (((new BigNumber(314)).isLessThanOrEqualTo(_1090_v22)) && ((_1090_v22).isLessThan(new BigNumber(504)))) {
                  _coll43.add((_1090_v22).multipliedBy((((_1088_v0).contains(false)) ? ((_1088_v0).get(false)) : ((_1087_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1087_v16).length))]))));
                }
              }
              return _coll43;
            }());
          })(_1081_v16, _1043_v0))).length);
          let _lhs167 = globalState;
          let _lhs168 = globalState;
          let _lhs169 = _this;
          let _lhs170 = _1081_v16;
          let _lhs171 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          _1044_v1 = _rhs205;
          _lhs167.f0 = _rhs206;
          _lhs168.f0 = _rhs207;
          _lhs169.f8 = _rhs208;
          _lhs170[_lhs171] = _rhs209;
          let _index179 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length));
          (_1081_v16)[_index179] = ((p0) ? ((_this).f9) : ((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]));
          let _1091_v23;
          let _nw171 = new _module.C2();
          _nw171.__ctor((_this).f5);
          _1091_v23 = _nw171;
          let _1092_v24;
          _1092_v24 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f9));
          let _rhs210 = _1091_v23;
          let _rhs211 = _module.__default.safeDivisionInt(((_this).f9).multipliedBy((_this).fm4((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], _1045_v2, new BigNumber((_1092_v24).length), globalState)), (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]);
          let _lhs172 = globalState;
          _1091_v23 = _rhs210;
          _lhs172.f0 = _rhs211;
        } else {
          let _1093_v25;
          _1093_v25 = _dafny.Seq.of((_1080_v15).update(_this.f8, _this.f8));
          (_this).f8 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_1094_v15) => function (_1095_i1) {
            return _1094_v15;
          })(_1080_v15)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), ((_1096_cf24, _1097_v16) => function (_1098_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_this.f8,(_1096_cf24)[_module.__default.safeIndex((_1097_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1097_v16).length))], new BigNumber((_1096_cf24).length))]);
          })(_1079_cf24, _1081_v16))), _1093_v25);
          (_this).f8 = (_this).f6;
          let _arr41 = _this.f7;
          let _index180 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_this.f7).length));
          _arr41[_index180] = (_this).f6;
          let _arr42 = _this.f7;
          let _index181 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_this.f7).length));
          _arr42[_index181] = !(!(new BigNumber((_1082_v17).length)).isEqualTo(new BigNumber((_1079_cf24).length)));
          let _arr43 = _this.f7;
          let _index182 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_this.f7).length));
          _arr43[_index182] = (_this.f7)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_this.f7).length))];
          let _1099_v26;
          let _nw172 = new _module.C0();
          _nw172.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1079_cf24, _module.__default.safeIndex((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], new BigNumber((_1079_cf24).length)), (_this).f10), _module.__default.safeIndex((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], new BigNumber((_dafny.Seq.update(_1079_cf24, _module.__default.safeIndex((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], new BigNumber((_1079_cf24).length)), (_this).f10)).length)), p0)).length),(_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]),(_this).fm2((_this).f10, globalState)));
          _1099_v26 = _nw172;
        }
        let _1100_v27;
        _1100_v27 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]));
        let _1101_v28;
        _1101_v28 = _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber(-541), (_this).fm4((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], (_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))], _1045_v2, (_dafny.ZERO).minus((_1081_v16)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1081_v16).length))]), globalState)));
        let _rhs212 = (_1100_v27).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_1102_i3) {
          return (_this).f9;
        })));
        let _rhs213 = false;
        let _rhs214 = new BigNumber((_1101_v28).length);
        let _lhs173 = _this;
        let _lhs174 = _this;
        let _lhs175 = globalState;
        _lhs173.f8 = _rhs212;
        _lhs174.f8 = _rhs213;
        _lhs175.f0 = _rhs214;
      }
      if ((_this).f6) {
        let _1103_v29;
        _1103_v29 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1104_v30;
        _1104_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _1105_v31;
        _1105_v31 = _module.D2.create_DC7(_1103_v29, _1104_v30, (_this).f9);
        let _1106_v32;
        _1106_v32 = _module.D2.create_DC8(_1105_v31);
        let _source16 = _1106_v32;
        if (_source16.is_DC5) {
          let _1107___mcc_h9 = (_source16).cf5;
          let _1108_cf5 = _1107___mcc_h9;
          let _1109_v33;
          let _nw173 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1109_v33 = _nw173;
          let _index183 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length));
          (_1109_v33)[_index183] = (_this).f9;
          let _index184 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1109_v33).length));
          (_1109_v33)[_index184] = ((_this).f9).plus((_this).f9);
          let _index185 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length));
          let _index186 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1109_v33).length));
          let _rhs215 = (_this).f9;
          let _rhs216 = ((_this).f9).plus(new BigNumber(721));
          let _rhs217 = (new BigNumber((_1044_v1).length)).minus(new BigNumber(-353));
          let _rhs218 = !((((_this).f10) ? (p0) : ((_this).f10)));
          let _lhs176 = _1109_v33;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length));
          let _lhs178 = _1109_v33;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1109_v33).length));
          let _lhs180 = globalState;
          let _lhs181 = _this;
          _lhs176[_lhs177] = _rhs215;
          _lhs178[_lhs179] = _rhs216;
          _lhs180.f0 = _rhs217;
          _lhs181.f8 = _rhs218;
          (globalState).f0 = (_this).f9;
          let _1110_v34;
          _1110_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10);
          let _index187 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length));
          (_1109_v33)[_index187] = ((new BigNumber((_1110_v34).length)).minus(new BigNumber(301))).multipliedBy((_1109_v33)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length))]);
          let _1111_v35;
          _1111_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1103_v29);
          let _index188 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1109_v33).length));
          (_1109_v33)[_index188] = (_dafny.ZERO).minus((new BigNumber((_1111_v35).length)).plus((_this).f9));
        } else if (_source16.is_DC6) {
          let _1112_v36;
          let _nw174 = new _module.C2();
          _nw174.__ctor((_this).f5);
          _1112_v36 = _nw174;
          let _1113_v37;
          _1113_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1112_v36,!(true));
          _1113_v37 = (_1113_v37).Merge(_1113_v37);
          let _1114_v40;
          _1114_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1104_v30,p0);
          let _1115_v41;
          let _nw175 = new _module.C0();
          _nw175.__ctor((_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll44 = new _dafny.Map();
            for (const _compr_44 of (function () {
              let _coll45 = new _dafny.Map();
              for (const _compr_45 of _dafny.IntegerRange(new BigNumber(50), new BigNumber(-133))) {
                let _1116_v39 = _compr_45;
                if (((new BigNumber(50)).isLessThanOrEqualTo(_1116_v39)) && ((_1116_v39).isLessThan(new BigNumber(-133)))) {
                  _coll45.push([_module.__default.safeDivisionInt(_1116_v39, (_this).f9),_this.f8]);
                }
              }
              return _coll45;
            }()).Keys.Elements) {
              let _1117_v38 = _compr_44;
              if ((function () {
                let _coll46 = new _dafny.Map();
                for (const _compr_46 of _dafny.IntegerRange(new BigNumber(50), new BigNumber(-133))) {
                  let _1116_v39 = _compr_46;
                  if (((new BigNumber(50)).isLessThanOrEqualTo(_1116_v39)) && ((_1116_v39).isLessThan(new BigNumber(-133)))) {
                    _coll46.push([_module.__default.safeDivisionInt(_1116_v39, (_this).f9),_this.f8]);
                  }
                }
                return _coll46;
              }()).contains(_1117_v38)) {
                _coll44.push([(_1117_v38).plus((_this).f9),new BigNumber(175)]);
              }
            }
            return _coll44;
          }(),_this.f8)).Merge(_1114_v40));
          _1115_v41 = _nw175;
          let _1118_v42;
          _1118_v42 = _dafny.Seq.of(new BigNumber((p1).length), (_this).f9);
          let _1119_v43;
          let _nw176 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _1119_v43 = _nw176;
          let _1120_v44;
          _1120_v44 = _dafny.Seq.of(p1, _1118_v42, ((p0) ? (_1118_v42) : (_dafny.Seq.update(_1118_v42, _module.__default.safeIndex(new BigNumber((_1104_v30).length), new BigNumber((_1118_v42).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm4(new BigNumber((_1044_v1).length), (_this).f9, _dafny.Map.Empty.slice().updateUnsafe(_1044_v1,(_this).f6), (_this).f9, globalState))).length)))), p1, p1);
          let _1121_v45;
          _1121_v45 = _module.D3.create_DC11((_this).f9);
          let _1122_v46;
          _1122_v46 = _dafny.Set.fromElements(_1121_v45);
          let _1123_v47;
          _1123_v47 = _dafny.MultiSet.fromElements((_this).f9);
          let _rhs219 = (_1120_v44)[_module.__default.safeIndex((_this).f9, new BigNumber((_1120_v44).length))];
          let _rhs220 = _1119_v43;
          let _rhs221 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_module.D3.create_DC11((_this).f9), _1121_v45)).IsSubsetOf(_1122_v46),p0)).length);
          let _rhs222 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f9), new BigNumber((_1123_v47).cardinality()));
          let _lhs182 = globalState;
          let _lhs183 = globalState;
          _1118_v42 = _rhs219;
          _1119_v43 = _rhs220;
          _lhs182.f0 = _rhs221;
          _lhs183.f0 = _rhs222;
          let _1124_v48;
          _1124_v48 = _dafny.Seq.of(_1046_v3, _1046_v3);
          let _1125_v49;
          let _nw177 = new _module.C4();
          _nw177.__ctor(_this.f7, _this.f8, p0, (_dafny.MultiSet.fromElements(p0)).update(_this.f8, _module.__default.abs(new BigNumber((_1124_v48).length))));
          _1125_v49 = _nw177;
          let _1126_v50;
          _1126_v50 = _dafny.Set.fromElements(_1125_v49);
          let _1127_v51;
          _1127_v51 = _dafny.Seq.of((_1126_v50).IsSubsetOf(_dafny.Set.fromElements(_1125_v49, _1125_v49)));
          let _1128_v52;
          _1128_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1103_v29,(_this).f9);
          (_this).f8 = (_1127_v51)[_module.__default.safeIndex((((_1128_v52).contains(_1103_v29)) ? ((_1128_v52).get(_1103_v29)) : ((_this).f9)), new BigNumber((_1127_v51).length))];
        } else if (_source16.is_DC7) {
          let _1129___mcc_h10 = (_source16).cf6;
          let _1130___mcc_h11 = (_source16).cf7;
          let _1131___mcc_h12 = (_source16).cf8;
          let _1132_cf8 = _1131___mcc_h12;
          let _1133_cf7 = _1130___mcc_h11;
          let _1134_cf6 = _1129___mcc_h10;
          let _1135_v53;
          let _init28 = ((_1136_p0) => function (_1137_i4) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1136_p0,_1136_p0);
          })(p0);
          let _nw178 = Array((new BigNumber(18)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw178.length); _i0_28++) {
            _nw178[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1135_v53 = _nw178;
          let _init29 = function (_1138_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe(_this.f8,(_this).f6);
          };
          let _nw179 = Array((new BigNumber(19)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw179.length); _i0_29++) {
            _nw179[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1135_v53 = _nw179;
          (_this).f8 = _this.f8;
          (_this).f8 = (_this).f10;
          let _rhs223 = ((_this).f9).minus(_1132_cf8);
          let _rhs224 = (_this).f9;
          let _rhs225 = ((((_1043_v0).contains((_this).f6)) ? ((_1043_v0).get((_this).f6)) : ((_this).f9))).plus((_this).f9);
          let _lhs184 = globalState;
          let _lhs185 = globalState;
          _1132_cf8 = _rhs223;
          _lhs184.f0 = _rhs224;
          _lhs185.f0 = _rhs225;
        } else if (_source16.is_DC4) {
          let _1139___mcc_h13 = (_source16).cf4;
          let _1140_cf4 = _1139___mcc_h13;
          _1103_v29 = _1103_v29;
          (globalState).f0 = (_this).f9;
          let _1141_v54;
          _1141_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1104_v30,_this.f8);
          let _1142_v55;
          let _nw180 = new _module.C0();
          _nw180.__ctor(_1141_v54);
          _1142_v55 = _nw180;
          let _1143_v56;
          let _init30 = ((_1144_p0, _1145_v1) => function (_1146_i6) {
            return ((_1144_p0) ? (_1145_v1) : (_dafny.Seq.UnicodeFromString("esgdfxat")));
          })(p0, _1044_v1);
          let _nw181 = Array((new BigNumber(5)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw181.length); _i0_30++) {
            _nw181[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1143_v56 = _nw181;
          let _index189 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1143_v56).length));
          (_1143_v56)[_index189] = _1044_v1;
          let _index190 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1143_v56).length));
          (_1143_v56)[_index190] = _dafny.Seq.Concat(_1044_v1, _1044_v1);
        } else {
          let _1147___mcc_h14 = (_source16).cf9;
          let _1148_cf9 = _1147___mcc_h14;
          (globalState).f0 = ((_this).f9).plus(new BigNumber((p1).length));
          let _1149_v57;
          _1149_v57 = _module.D5.create_DC15(_1044_v1);
          let _1150_v58;
          _1150_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v57,_1044_v1);
          let _1151_v59;
          let _nw182 = new _module.C1();
          _nw182.__ctor(p0, (((_1150_v58).contains(_1149_v57)) ? ((_1150_v58).get(_1149_v57)) : (_1044_v1)), (_this).f6, (_dafny.MultiSet.fromElements((_this).f10)).Union((_this).f5));
          _1151_v59 = _nw182;
          let _1152_v60;
          _1152_v60 = _dafny.Seq.of((_this).f10, false);
          (_this).f8 = (_1152_v60)[_module.__default.safeIndex((_dafny.ZERO).minus(((_this).f9).minus((_this).f9)), new BigNumber((_1152_v60).length))];
          (globalState).f0 = (_dafny.ZERO).minus((_this).f9);
        }
        let _1153_v61;
        _1153_v61 = _dafny.Set.fromElements(_1046_v3);
        let _arr44 = _this.f7;
        let _index191 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
        _arr44[_index191] = p0;
        let _arr45 = _this.f7;
        let _index192 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
        let _rhs226 = _1153_v61;
        let _rhs227 = (_this).f9;
        let _rhs228 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("uqurwber"), _module.__default.fm17(globalState));
        let _rhs229 = !(_this.f8);
        let _rhs230 = (_this).f9;
        let _lhs186 = globalState;
        let _lhs187 = _this.f7;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
        let _lhs189 = _this;
        let _lhs190 = globalState;
        _1153_v61 = _rhs226;
        _lhs186.f0 = _rhs227;
        _lhs187[_lhs188] = _rhs228;
        _lhs189.f8 = _rhs229;
        _lhs190.f0 = _rhs230;
        if ((_this.f7)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length))]) {
          _1044_v1 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ism"), _1044_v1);
          let _1154_v62;
          let _init31 = function (_1155_i7) {
            return _module.D1.create_DC2((_this.f7)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length))]);
          };
          let _nw183 = Array((new BigNumber(27)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw183.length); _i0_31++) {
            _nw183[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1154_v62 = _nw183;
          let _1156_v63;
          _1156_v63 = _module.D1.create_DC2((_this).f10);
          let _index193 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_1154_v62).length));
          (_1154_v62)[_index193] = _1156_v63;
          let _1157_v64;
          let _init32 = function (_1158_i8) {
            return (_1158_i8).multipliedBy((_dafny.ZERO).minus((_this).f9));
          };
          let _nw184 = Array((new BigNumber(15)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw184.length); _i0_32++) {
            _nw184[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1157_v64 = _nw184;
          let _index194 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1157_v64).length));
          (_1157_v64)[_index194] = (_this).f9;
          let _arr46 = _this.f7;
          let _index195 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          let _index196 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_1154_v62).length));
          let _arr47 = _this.f7;
          let _index197 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          let _index198 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1157_v64).length));
          let _rhs231 = (_this.f7)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length))];
          let _rhs232 = ((_this).f9).plus(((_this).f9).minus((_this).f9));
          let _rhs233 = _module.D1.create_DC2(false);
          let _rhs234 = !(_dafny.Seq.IsPrefixOf(_1044_v1, _1044_v1)) || (p0);
          let _rhs235 = (_this).f9;
          let _lhs191 = _this.f7;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          let _lhs193 = globalState;
          let _lhs194 = _1154_v62;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_1154_v62).length));
          let _lhs196 = _this.f7;
          let _lhs197 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          let _lhs198 = _1157_v64;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1157_v64).length));
          _lhs191[_lhs192] = _rhs231;
          _lhs193.f0 = _rhs232;
          _lhs194[_lhs195] = _rhs233;
          _lhs196[_lhs197] = _rhs234;
          _lhs198[_lhs199] = _rhs235;
          let _index199 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1157_v64).length));
          (_1157_v64)[_index199] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f9, (_this).f9), (_dafny.ZERO).minus((_this).f9));
          let _index200 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1157_v64).length));
          (_1157_v64)[_index200] = _module.__default.safeDivisionInt((((_this).f10) ? ((_this).f9) : (new BigNumber(569))), (((_1104_v30).contains((_this).f9)) ? ((_1104_v30).get((_this).f9)) : (new BigNumber((_dafny.Set.fromElements(_1046_v3)).length))));
          let _1159_v65;
          let _init33 = function (_1160_i9) {
            return (_this).f6;
          };
          let _nw185 = Array((new BigNumber(20)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw185.length); _i0_33++) {
            _nw185[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1159_v65 = _nw185;
          let _1161_v66;
          _1161_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1159_v65);
          (globalState).f0 = (_this).fm4((_this).f9, new BigNumber((_1161_v66).length), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(405)), function (_1162_i10) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          }),(_this).f10), new BigNumber((_dafny.Set.fromElements(p0, (_this).f6, p0)).length), globalState);
        } else {
          (globalState).f0 = ((_this).f9).minus(((_this).f9).plus((_this).f9));
          (_this).f8 = p0;
          let _1163_v67;
          let _nw186 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _1163_v67 = _nw186;
          let _1164_v68;
          let _1165_v69;
          let _1166_v70;
          let _out10;
          let _out11;
          let _out12;
          let _outcollector2 = (_this).m5((_this).f9, _dafny.MultiSet.fromElements((_this).f10), _1163_v67, p0, globalState);
          _out10 = _outcollector2[0];
          _out11 = _outcollector2[1];
          _out12 = _outcollector2[2];
          _1164_v68 = _out10;
          _1165_v69 = _out11;
          _1166_v70 = _out12;
          let _1167_v71;
          let _init34 = function (_1168_i11) {
            return (_this).f10;
          };
          let _nw187 = Array((new BigNumber(20)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw187.length); _i0_34++) {
            _nw187[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1167_v71 = _nw187;
          _1167_v71 = _1167_v71;
          let _1169_v72;
          let _nw188 = new _module.C3();
          _nw188.__ctor(_1166_v70, (_this).f5);
          _1169_v72 = _nw188;
        }
        let _1170_v73;
        let _nw189 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1170_v73 = _nw189;
        let _1171_v74;
        _1171_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1170_v73,(_this).f10);
        let _1172_v75;
        _1172_v75 = _module.D9.create_DC27(_1171_v74, _this.f8);
        let _1173_v76;
        _1173_v76 = _dafny.Seq.of(true);
        let _1174_v77;
        _1174_v77 = _module.D8.create_DC24((((_1043_v0).contains((_this).f6)) ? ((_1043_v0).get((_this).f6)) : (new BigNumber((_dafny.Seq.UnicodeFromString("bwkuqlte")).length))), (_1172_v75).dtor_cf43, _1173_v76, (((_this).f6) ? ((_this).f9) : ((_this).f9)), (_this).f9);
        let _source17 = _1174_v77;
        if (_source17.is_DC24) {
          let _1175___mcc_h15 = (_source17).cf34;
          let _1176___mcc_h16 = (_source17).cf35;
          let _1177___mcc_h17 = (_source17).cf36;
          let _1178___mcc_h18 = (_source17).cf37;
          let _1179___mcc_h19 = (_source17).cf38;
          let _1180_cf38 = _1179___mcc_h19;
          let _1181_cf37 = _1178___mcc_h18;
          let _1182_cf36 = _1177___mcc_h17;
          let _1183_cf35 = _1176___mcc_h16;
          let _1184_cf34 = _1175___mcc_h15;
          let _1185_v79;
          _1185_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1184_cf34,_1183_cf35);
          _1103_v29 = _module.__default.fm16(p0, function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_1185_v79).Keys.Elements) {
              let _1186_v78 = _compr_47;
              if ((_1185_v79).contains(_1186_v78)) {
                _coll47.push([(_1186_v78).minus((_module.D0.create_DC0(_1184_cf34)).dtor_cf0),_1180_cf38]);
              }
            }
            return _coll47;
          }(), _1044_v1, globalState);
          _1173_v76 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1173_v76, _1182_cf36), _dafny.Seq.Concat(_1182_cf36, _1173_v76));
          let _1187_v80;
          _1187_v80 = _module.D9.create_DC27(_1171_v74, _this.f8);
          let _1188_v81;
          _1188_v81 = _module.D9.create_DC28(_1187_v80);
          let _1189_v82;
          let _nw190 = new _module.C4();
          _nw190.__ctor(_1170_v73, _1183_cf35, false, (_this).f5);
          _1189_v82 = _nw190;
          let _1190_v83;
          _1190_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1188_v81,_1189_v82);
          let _1191_v84;
          _1191_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1190_v83);
          let _1192_v85;
          let _nw191 = new _module.C2();
          _nw191.__ctor((_this).f5);
          _1192_v85 = _nw191;
          let _1193_v86;
          _1193_v86 = _dafny.Set.fromElements(((((_1191_v84).contains((_this).f10)) ? ((_1191_v84).get((_this).f10)) : (_1190_v83))).update(_module.D9.create_DC28(_module.D9.create_DC26(_1192_v85)), _1189_v82), _1190_v83);
          let _1194_v87;
          _1194_v87 = _dafny.Seq.of(_1193_v86);
          let _1195_v88;
          _1195_v88 = _dafny.MultiSet.fromElements(_1180_cf38);
          (globalState).f0 = (_dafny.ZERO).minus(new BigNumber(((_1194_v87)[_module.__default.safeIndex(new BigNumber(((_1195_v88).Intersect(_dafny.MultiSet.FromArray(p1))).cardinality()), new BigNumber((_1194_v87).length))]).length));
          (globalState).f0 = (_this).f9;
        } else if (_source17.is_DC25) {
          let _1196___mcc_h20 = (_source17).cf39;
          let _1197___mcc_h21 = (_source17).cf40;
          let _1198_cf40 = _1197___mcc_h21;
          let _1199_cf39 = _1196___mcc_h20;
          _1104_v30 = (_1104_v30).Merge(_1104_v30);
          _1044_v1 = _1044_v1;
          let _1200_v89;
          _1200_v89 = _dafny.Seq.of(_1173_v76);
          let _1201_v90;
          _1201_v90 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1200_v89);
          _1201_v90 = (_1201_v90).update(((_this).fm2(_1199_cf39, globalState)) === ((_this).f10), _dafny.Seq.Create(_module.__default.abs(new BigNumber(810)), ((_1202_v76) => function (_1203_i12) {
            return _1202_v76;
          })(_1173_v76)));
          (globalState).f0 = (_this).f9;
        } else {
          let _1204___mcc_h22 = (_source17).cf33;
          let _1205_cf33 = _1204___mcc_h22;
          (_this).f8 = (_this.f7)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length))];
          let _arr48 = _this.f7;
          let _index201 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          let _rhs236 = !(_this.f8);
          let _rhs237 = _1173_v76;
          let _lhs200 = _this.f7;
          let _lhs201 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length));
          _lhs200[_lhs201] = _rhs236;
          _1173_v76 = _rhs237;
          let _1206_v91;
          let _nw192 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _1206_v91 = _nw192;
          let _1207_v92;
          let _1208_v93;
          let _1209_v94;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = (_this).m5(_module.__default.safeModuloInt((_this).f9, (_this).f9), _dafny.MultiSet.fromElements(true, !((_this).f10), !(p0)), _1206_v91, p0, globalState);
          _out13 = _outcollector3[0];
          _out14 = _outcollector3[1];
          _out15 = _outcollector3[2];
          _1207_v92 = _out13;
          _1208_v93 = _out14;
          _1209_v94 = _out15;
          let _1210_v95;
          _1210_v95 = _dafny.Seq.of(_1173_v76);
          let _1211_v96;
          _1211_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), ((_1212_v93) => function (_1213_i13) {
            return new BigNumber((_dafny.Set.fromElements(_1212_v93)).length);
          })(_1208_v93)));
          let _1214_v97;
          _1214_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this.f7)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_this.f7).length))]);
          let _1215_v98;
          let _nw193 = Array((new BigNumber(24)).toNumber());
          _nw193[(_dafny.ZERO).toNumber()] = (new BigNumber(566)).multipliedBy((_this).f9);
          _nw193[(_dafny.ONE).toNumber()] = new BigNumber((_1210_v95).length);
          _nw193[(new BigNumber(2)).toNumber()] = _1208_v93;
          _nw193[(new BigNumber(3)).toNumber()] = (_this).f9;
          _nw193[(new BigNumber(4)).toNumber()] = (_this).f9;
          _nw193[(new BigNumber(5)).toNumber()] = (_1208_v93).minus((_this).f9);
          _nw193[(new BigNumber(6)).toNumber()] = ((p0) ? (_1208_v93) : (_1208_v93));
          _nw193[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_1208_v93);
          _nw193[(new BigNumber(8)).toNumber()] = _1208_v93;
          _nw193[(new BigNumber(9)).toNumber()] = new BigNumber(-602);
          _nw193[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_1208_v93, _1208_v93);
          _nw193[(new BigNumber(11)).toNumber()] = (_this).f9;
          _nw193[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(((_this).f9).multipliedBy((_this).f9));
          _nw193[(new BigNumber(13)).toNumber()] = ((_this).f9).minus((_this).f9);
          _nw193[(new BigNumber(14)).toNumber()] = new BigNumber((_1211_v96).length);
          _nw193[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(_1208_v93, new BigNumber(-348));
          _nw193[(new BigNumber(16)).toNumber()] = ((false) ? ((_this).f9) : ((_this).f9));
          _nw193[(new BigNumber(17)).toNumber()] = ((_this).f9).multipliedBy(_1208_v93);
          _nw193[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((((_this).f6) ? (_1208_v93) : (new BigNumber(-714))));
          _nw193[(new BigNumber(19)).toNumber()] = new BigNumber((_1043_v0).length);
          _nw193[(new BigNumber(20)).toNumber()] = _1208_v93;
          _nw193[(new BigNumber(21)).toNumber()] = _1208_v93;
          _nw193[(new BigNumber(22)).toNumber()] = (_this).f9;
          _nw193[(new BigNumber(23)).toNumber()] = (_this).fm4((_this).f9, (_this).f9, _1045_v2, new BigNumber((_1214_v97).length), globalState);
          _1215_v98 = _nw193;
          let _1216_v99;
          _1216_v99 = _dafny.Seq.of(_1215_v98, (((_this).f6) ? (_1215_v98) : (_1215_v98)), _1215_v98);
          _1215_v98 = (_1216_v99)[_module.__default.safeIndex((_this).f9, new BigNumber((_1216_v99).length))];
        }
        (globalState).f0 = (_this).f9;
      } else {
        (_this).f7 = _this.f7;
        let _1217_v100;
        _1217_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(335),(_this).f6);
        let _1218_v101;
        _1218_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _1219_v102;
        _1219_v102 = _dafny.MultiSet.fromElements(_module.__default.fm27(new BigNumber((_dafny.Seq.UnicodeFromString("ulvsylkks")).length), (((_1217_v100).contains((_this).f9)) ? ((_1217_v100).get((_this).f9)) : (!(_this.f8))), (_this).f9, _1217_v100, globalState), _1218_v101);
        let _1220_v103;
        let _nw194 = Array((new BigNumber(19)).toNumber()).fill(_module.D7.Default());
        _1220_v103 = _nw194;
        let _1221_v104;
        _1221_v104 = _module.D7.create_DC21((_this).f9, (_dafny.ZERO).minus((_this).f9));
        let _index202 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1220_v103).length));
        (_1220_v103)[_index202] = _1221_v104;
        let _1222_v105;
        let _nw195 = Array((new BigNumber(20)).toNumber()).fill([]);
        _1222_v105 = _nw195;
        let _1223_v106;
        let _init35 = function (_1224_i14) {
          return (_1224_i14).multipliedBy(new BigNumber(153));
        };
        let _nw196 = Array((new BigNumber(23)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw196.length); _i0_35++) {
          _nw196[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1223_v106 = _nw196;
        let _index203 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length));
        (_1223_v106)[_index203] = (_this).f9;
        let _1225_v107;
        _1225_v107 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4((_this).f9, new BigNumber((_1044_v1).length), _1045_v2, (_this).fm4((_this).f9, new BigNumber(46), _1045_v2, (_this).f9, globalState), globalState),_1223_v106);
        let _1226_v108;
        _1226_v108 = _dafny.MultiSet.fromElements(new BigNumber((_1225_v107).length));
        let _index204 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1220_v103).length));
        let _index205 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length));
        let _rhs238 = (_module.__default.fm28((_this).f9, _1218_v101, (_this).f10, (_this).f9, globalState)).Intersect(_1219_v102);
        let _rhs239 = _module.__default.fm29(_1044_v1, p1, _1226_v108, globalState);
        let _rhs240 = _1222_v105;
        let _rhs241 = (_this).f9;
        let _rhs242 = new BigNumber((_1043_v0).length);
        let _lhs202 = _1220_v103;
        let _lhs203 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1220_v103).length));
        let _lhs204 = _1223_v106;
        let _lhs205 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length));
        let _lhs206 = globalState;
        _1219_v102 = _rhs238;
        _lhs202[_lhs203] = _rhs239;
        _1222_v105 = _rhs240;
        _lhs204[_lhs205] = _rhs241;
        _lhs206.f0 = _rhs242;
        if (!((_this).f6)) {
          let _1227_v109;
          let _nw197 = new _module.C4();
          _nw197.__ctor(_this.f7, (_this).f10, true, (_this).f5);
          _1227_v109 = _nw197;
          _1227_v109 = _1227_v109;
          let _index206 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length));
          (_1223_v106)[_index206] = new BigNumber(-90);
          let _index207 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length));
          (_1223_v106)[_index207] = _module.__default.safeModuloInt((p1)[_module.__default.safeIndex((_this).f9, new BigNumber((p1).length))], (_this).f9);
          (_this).f8 = _this.f8;
          (_1227_v109).f8 = (_this).f6;
        } else {
          let _1228_v110;
          let _nw198 = Array((new BigNumber(27)).toNumber()).fill([]);
          _1228_v110 = _nw198;
          let _index208 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1228_v110).length));
          (_1228_v110)[_index208] = _this.f7;
          let _index209 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_1228_v110).length));
          (_1228_v110)[_index209] = ((((_this).f5).IsProperSubsetOf((_this).f5)) ? (_this.f7) : (_this.f7));
          (globalState).f0 = (_this).f9;
          (_this).f8 = p0;
          (globalState).f0 = (_1223_v106)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length))];
          let _1229_v111;
          _1229_v111 = _dafny.Seq.of(p1, p1, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), function (_1230_i15) {
            return new BigNumber((_dafny.Seq.of(false, _this.f8)).length);
          }), _module.__default.fm8((_this).f9, _this.f8, _1044_v1, globalState));
          _1229_v111 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-870)), ((_1231_p1) => function (_1232_i16) {
            return _1231_p1;
          })(p1)), _1229_v111);
        }
        let _1233_v112;
        _1233_v112 = _dafny.Seq.of(p1, _dafny.Seq.Concat(p1, p1), _dafny.Seq.Concat(p1, p1));
        _1233_v112 = _1233_v112;
        let _rhs243 = (_dafny.ZERO).minus(((_1223_v106)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_1223_v106).length))]).multipliedBy((((_1226_v108).contains(new BigNumber((_1217_v100).length))) ? ((_1226_v108).get(new BigNumber((_1217_v100).length))) : ((_this).f9))));
        let _rhs244 = _1044_v1;
        let _rhs245 = _1223_v106;
        let _lhs207 = globalState;
        _lhs207.f0 = _rhs243;
        _1044_v1 = _rhs244;
        _1223_v106 = _rhs245;
      }
      let _1234_v113;
      _1234_v113 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_module.__default.safeModuloInt((_this).f9, (_this).f9));
      _1234_v113 = (_1234_v113).update((_this).f9, (_this).f9);
      let _1235_v114;
      _1235_v114 = _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,false));
      let _1236_v115;
      _1236_v115 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,false);
      let _pat_let_tv27 = p0;
      let _1237_v116;
      let _nw199 = Array((new BigNumber(29)).toNumber());
      _nw199[(_dafny.ZERO).toNumber()] = _1235_v114;
      _nw199[(_dafny.ONE).toNumber()] = _module.D11.create_DC32(_module.__default.fm9(globalState));
      _nw199[(new BigNumber(2)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(3)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(4)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(5)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(6)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(7)).toNumber()] = _module.D11.create_DC32(_1236_v115);
      _nw199[(new BigNumber(8)).toNumber()] = function (_pat_let34_0) {
        return function (_1238_dt__update__tmp_h2) {
          return function (_pat_let35_0) {
            return function (_1239_dt__update_hcf51_h0) {
              return _module.D11.create_DC32(_1239_dt__update_hcf51_h0);
            }(_pat_let35_0);
          }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv27,true));
        }(_pat_let34_0);
      }(_1235_v114);
      _nw199[(new BigNumber(9)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(10)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(11)).toNumber()] = _module.__default.fm30(globalState);
      _nw199[(new BigNumber(12)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(13)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(14)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(15)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(16)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(17)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(18)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(19)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(20)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(21)).toNumber()] = _module.D11.create_DC32(_1236_v115);
      _nw199[(new BigNumber(22)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(23)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(24)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(25)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(26)).toNumber()] = _1235_v114;
      _nw199[(new BigNumber(27)).toNumber()] = _module.D11.create_DC32(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f8));
      _nw199[(new BigNumber(28)).toNumber()] = _1235_v114;
      _1237_v116 = _nw199;
      let _index210 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1237_v116).length));
      (_1237_v116)[_index210] = _1235_v114;
      let _index211 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1237_v116).length));
      (_1237_v116)[_index211] = _1235_v114;
      let _1240_v117;
      let _nw200 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1240_v117 = _nw200;
      let _1241_v118;
      _1241_v118 = new _dafny.CodePoint('q'.codePointAt(0));
      let _index212 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_1240_v117).length));
      (_1240_v117)[_index212] = _1241_v118;
      let _1242_v119;
      let _nw201 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
      _1242_v119 = _nw201;
      let _1243_v120;
      _1243_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1234_v113,_this.f8);
      let _1244_v121;
      let _nw202 = new _module.C0();
      _nw202.__ctor(_1243_v120);
      _1244_v121 = _nw202;
      let _1245_v122;
      _1245_v122 = _dafny.Set.fromElements(_1244_v121, _1244_v121, _1244_v121);
      let _index213 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1242_v119).length));
      (_1242_v119)[_index213] = _1245_v122;
      let _1246_v123;
      _1246_v123 = _module.D2.create_DC8(_module.D2.create_DC8(_module.D2.create_DC7(new _dafny.CodePoint('e'.codePointAt(0)), _1234_v113, (_this).f9)));
      let _1247_v124;
      _1247_v124 = _dafny.Set.fromElements((_this).f9, (_this).f9);
      let _1248_v125;
      _1248_v125 = _dafny.Seq.of((_this).f5, (_this).f5);
      let _1249_v126;
      _1249_v126 = _module.D10.create_DC31(_1248_v125, _1241_v118);
      let _pat_let_tv28 = _1241_v118;
      let _pat_let_tv29 = _1241_v118;
      let _pat_let_tv30 = _1241_v118;
      let _index214 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_1240_v117).length));
      let _index215 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1242_v119).length));
      let _rhs246 = function (_source18) {
        if (_source18.is_DC30) {
          let _1250___mcc_h23 = (_source18).cf46;
          let _1251___mcc_h24 = (_source18).cf47;
          let _1252___mcc_h25 = (_source18).cf48;
          let _1253_cf48 = _1252___mcc_h25;
          let _1254_cf47 = _1251___mcc_h24;
          let _1255_cf46 = _1250___mcc_h23;
          return _pat_let_tv28;
        } else if (_source18.is_DC31) {
          let _1256___mcc_h26 = (_source18).cf49;
          let _1257___mcc_h27 = (_source18).cf50;
          let _1258_cf50 = _1257___mcc_h27;
          let _1259_cf49 = _1256___mcc_h26;
          return _pat_let_tv29;
        } else {
          let _1260___mcc_h28 = (_source18).cf45;
          let _1261_cf45 = _1260___mcc_h28;
          return _pat_let_tv30;
        }
      }(_1249_v126);
      let _rhs247 = _1245_v122;
      let _rhs248 = _1246_v123;
      let _rhs249 = _1247_v124;
      let _rhs250 = (_1043_v0).Merge((((_this).f6) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber(-252))) : (_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f9))));
      let _lhs208 = _1240_v117;
      let _lhs209 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_1240_v117).length));
      let _lhs210 = _1242_v119;
      let _lhs211 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1242_v119).length));
      _lhs208[_lhs209] = _rhs246;
      _lhs210[_lhs211] = _rhs247;
      _1246_v123 = _rhs248;
      _1247_v124 = _rhs249;
      _1043_v0 = _rhs250;
      return;
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1262_v2;
      _1262_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f9);
      let _1263_v3;
      _1263_v3 = _dafny.Set.fromElements(_1262_v2, _1262_v2);
      let _1264_v4;
      let _nw203 = new _module.C0();
      _nw203.__ctor(function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of (function () {
          let _coll49 = new _dafny.Map();
          for (const _compr_49 of (_1263_v3).Elements) {
            let _1265_v1 = _compr_49;
            if ((_1263_v3).contains(_1265_v1)) {
              _coll49.push([_1265_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f6)).length)]);
            }
          }
          return _coll49;
        }()).Keys.Elements) {
          let _1266_v0 = _compr_48;
          if ((function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of (_1263_v3).Elements) {
              let _1265_v1 = _compr_50;
              if ((_1263_v3).contains(_1265_v1)) {
                _coll50.push([_1265_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f6)).length)]);
              }
            }
            return _coll50;
          }()).contains(_1266_v0)) {
            _coll48.push([_1266_v0,(_this).f6]);
          }
        }
        return _coll48;
      }());
      _1264_v4 = _nw203;
      let _1267_v6;
      _1267_v6 = _module.D2.create_DC5((_this).f6);
      let _1268_v7;
      _1268_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1267_v6,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), ((_1269_p1) => function (_1270_i0) {
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_1271_p1) => function (_1272_i1) {
          return _1271_p1;
        })(_1269_p1)))).cardinality());
      })(p1))).length));
      let _1273_v8;
      _1273_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v7,(_this).f10);
      let _1274_v9;
      _1274_v9 = _module.D4.create_DC13(_this.f8, function () {
  let _coll51 = new _dafny.Map();
  for (const _compr_51 of (_1273_v8).Keys.Elements) {
    let _1275_v5 = _compr_51;
    if ((_1273_v8).contains(_1275_v5)) {
      _coll51.push([_1275_v5,false]);
    }
  }
  return _coll51;
}());
      let _1276_v10;
      let _nw204 = new _module.C4();
      _nw204.__ctor(_this.f7, (_this).f6, (!(p3)) && ((_1274_v9).dtor_cf15), p1);
      _1276_v10 = _nw204;
      let _1277_v11;
      _1277_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f10);
      let _1278_v12;
      _1278_v12 = _dafny.Set.fromElements(_1277_v11, _1277_v11, (_1277_v11).Merge(_1277_v11));
      let _1279_v13;
      _1279_v13 = _dafny.MultiSet.fromElements(_1277_v11, _1277_v11);
      let _1280_v15;
      _1280_v15 = _dafny.Map.Empty.slice().updateUnsafe((p0).isLessThanOrEqualTo((_this).f9),function () {
        let _coll52 = new _dafny.Set();
        for (const _compr_52 of (_1279_v13).Elements) {
          let _1281_v14 = _compr_52;
          if ((_1279_v13).contains(_1281_v14)) {
            _coll52.add(_1281_v14);
          }
        }
        return _coll52;
      }());
      let _1282_v16;
      _1282_v16 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).length));
      let _1283_v17;
      _1283_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1282_v16).cardinality()),_1262_v2);
      let _rhs251 = ((false) ? (_1276_v10) : (_1276_v10));
      let _rhs252 = (((_1280_v15).contains((_1283_v17).equals((_1283_v17).update((_this).f9, _1262_v2)))) ? ((_1280_v15).get((_1283_v17).equals((_1283_v17).update((_this).f9, _1262_v2)))) : (_module.__default.fm31(globalState)));
      let _rhs253 = p0;
      let _lhs212 = globalState;
      _1276_v10 = _rhs251;
      _1278_v12 = _rhs252;
      _lhs212.f0 = _rhs253;
      let _1284_v18;
      _1284_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1276_v10.f7,p3);
      let _1285_v19;
      _1285_v19 = _module.D9.create_DC27(_1284_v18, (_1276_v10).f6);
      if (!(((((_this).f10) ? (_module.D9.create_DC27(_1284_v18, (_this).f10)) : (_1285_v19))).dtor_cf43)) {
        _1262_v2 = (_1262_v2).update(p0, ((_this).f9).minus((_this).f9));
        let _1286_v20;
        _1286_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1276_v10).f6,(_this).f6);
        let _1287_v21;
        _1287_v21 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_1286_v20).length));
        let _1288_v22;
        _1288_v22 = _module.D7.create_DC19(_1287_v21);
        let _1289_v23;
        _1289_v23 = _dafny.Set.fromElements(_1288_v22, _1288_v22, _1288_v22, _1288_v22);
        (_1276_v10).f8 = ((_1289_v23).Union(_1289_v23)).IsDisjointFrom((_1289_v23).Difference(_dafny.Set.fromElements(_1288_v22)));
        (globalState).f0 = p0;
        let _1290_v24;
        _1290_v24 = _dafny.Set.fromElements(_this.f8, _1276_v10.f8, (_this).f10);
        r0 = (_1290_v24).equals(_module.__default.fm32((_1276_v10).f6, (_this).f10, !(true), globalState));
        let _1291_v25;
        _1291_v25 = _dafny.Seq.of(_1276_v10.f8);
        let _1292_v26;
        let _nw205 = Array((new BigNumber(6)).toNumber());
        _nw205[(_dafny.ZERO).toNumber()] = _1291_v25;
        _nw205[(_dafny.ONE).toNumber()] = _1291_v25;
        _nw205[(new BigNumber(2)).toNumber()] = _1291_v25;
        _nw205[(new BigNumber(3)).toNumber()] = _1291_v25;
        _nw205[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1291_v25, _module.__default.safeIndex((_this).f9, new BigNumber((_1291_v25).length)), false);
        _nw205[(new BigNumber(5)).toNumber()] = _1291_v25;
        _1292_v26 = _nw205;
        let _1293_v27;
        _1293_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1292_v26,(_this).f9);
        _1293_v27 = (_1293_v27).update(_1292_v26, new BigNumber(854));
      } else {
        let _1294_v28;
        _1294_v28 = new _dafny.CodePoint('x'.codePointAt(0));
        _1294_v28 = new _dafny.CodePoint('d'.codePointAt(0));
        r1 = (_this).f9;
        let _1295_v29;
        let _init36 = ((_1296_v10, _1297_p0) => function (_1298_i2) {
          return _module.__default.safeDivisionInt(_1298_i2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1296_v10.f8,_1297_p0)).length));
        })(_1276_v10, p0);
        let _nw206 = Array((new BigNumber(8)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw206.length); _i0_36++) {
          _nw206[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1295_v29 = _nw206;
        let _1299_v30;
        _1299_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1295_v29,(_this).f10);
        let _arr49 = _this.f7;
        let _index216 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_this.f7).length));
        _arr49[_index216] = (((_1299_v30).contains(_1295_v29)) ? ((_1299_v30).get(_1295_v29)) : ((_1276_v10).f6));
        let _1300_v31;
        _1300_v31 = _dafny.Seq.UnicodeFromString("w");
        let _arr50 = _this.f7;
        let _index217 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_this.f7).length));
        _arr50[_index217] = !_dafny.Seq.contains(_dafny.Seq.Concat(_1300_v31, _1300_v31), _1294_v28);
        (globalState).f0 = (_this).f9;
        (_1276_v10).m2((_1300_v31)[_module.__default.safeIndex((_this).f9, new BigNumber((_1300_v31).length))], globalState);
      }
      if (true) {
        let _1301_v32;
        _1301_v32 = new _dafny.CodePoint('e'.codePointAt(0));
        (_1276_v10).m2(_1301_v32, globalState);
        if ((p1).equals(_dafny.MultiSet.fromElements((_1276_v10).fm3(globalState)))) {
          r1 = (_this).f9;
          r0 = (_1276_v10).f6;
          r2 = true;
          let _1302_v33;
          _1302_v33 = _dafny.Seq.UnicodeFromString("nfwac");
          let _1303_v34;
          let _nw207 = Array((new BigNumber(12)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1302_v33, _1302_v33);
          _nw207[(_dafny.ONE).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(2)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(3)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(4)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(5)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(6)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("grcliofu"), _1302_v33);
          _nw207[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1302_v33, _module.__default.safeIndex(p0, new BigNumber((_1302_v33).length)), new _dafny.CodePoint('g'.codePointAt(0)));
          _nw207[(new BigNumber(9)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(10)).toNumber()] = _1302_v33;
          _nw207[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), ((_1304_v32) => function (_1305_i3) {
            return _1304_v32;
          })(_1301_v32));
          _1303_v34 = _nw207;
          let _index218 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1303_v34).length));
          (_1303_v34)[_index218] = _1302_v33;
          let _index219 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1303_v34).length));
          (_1303_v34)[_index219] = _dafny.Seq.Concat(_dafny.Seq.Concat(_1302_v33, _1302_v33), _dafny.Seq.Concat(_1302_v33, _dafny.Seq.UnicodeFromString("qo")));
          let _1306_v35;
          let _nw208 = new _module.C4();
          _nw208.__ctor(_1276_v10.f7, _this.f8, _1276_v10.f8, _module.__default.fm14(_1262_v2, _this.f8, new BigNumber((_module.__default.fm17(globalState)).length), true, globalState));
          _1306_v35 = _nw208;
        } else {
          let _1307_v36;
          _1307_v36 = _dafny.Seq.of((_this).f10, (_this).f6, p3, (_1276_v10).f6);
          (_1276_v10).m1(p0, _this.f8, !_dafny.Seq.contains(_1307_v36, _1276_v10.f8), globalState);
          (_this).f8 = (_module.__default.fm14(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(622),p0), _1276_v10.f8, p0, (_1276_v10).f6, globalState)).IsSubsetOf((_this).f5);
          (globalState).f0 = (p0).minus((_this).fm4(new BigNumber((_dafny.Seq.of(_this.f8, (_1276_v10).f6)).length), (_this).f9, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wrmdqy"),(_this).f6), p0, globalState));
          (_1276_v10).m2(_1301_v32, globalState);
          let _arr51 = _1276_v10.f7;
          let _index220 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1276_v10.f7).length));
          _arr51[_index220] = true;
          let _1308_v37;
          let _nw209 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1308_v37 = _nw209;
          let _index221 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1308_v37).length));
          (_1308_v37)[_index221] = (_this).f9;
          let _index222 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1308_v37).length));
          (_1308_v37)[_index222] = (_dafny.ZERO).minus((_this).f9);
          let _1309_v38;
          _1309_v38 = _dafny.Seq.UnicodeFromString("r");
          let _1310_v39;
          _1310_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1309_v38,(_1276_v10).f6);
          let _1311_v40;
          _1311_v40 = _module.D3.create_DC9(_1308_v37);
          let _1312_v41;
          _1312_v41 = _module.D5.create_DC16(_1311_v40, new BigNumber((_dafny.Seq.UnicodeFromString("kd")).length), p0, (_this).f9, _1264_v4);
          let _arr52 = _1276_v10.f7;
          let _index223 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1276_v10.f7).length));
          let _index224 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1308_v37).length));
          let _index225 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1308_v37).length));
          let _rhs254 = (p1).IsProperSubsetOf((_1276_v10).f5);
          let _rhs255 = ((_1276_v10).fm4(p0, (_this).f9, _1310_v39, (_this).f9, globalState)).minus(((_1312_v41).dtor_cf22).minus((_this).f9));
          let _rhs256 = _module.__default.safeDivisionInt((_this).fm4((_this).f9, (_dafny.ZERO).minus(p0), _1310_v39, (_this).f9, globalState), (_this).f9);
          let _lhs213 = _1276_v10.f7;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1276_v10.f7).length));
          let _lhs215 = _1308_v37;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1308_v37).length));
          let _lhs217 = _1308_v37;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_1308_v37).length));
          _lhs213[_lhs214] = _rhs254;
          _lhs215[_lhs216] = _rhs255;
          _lhs217[_lhs218] = _rhs256;
        }
        let _1313_v42;
        let _init37 = function (_1314_i4) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ljs"), _dafny.Seq.UnicodeFromString("ow"));
        };
        let _nw210 = Array((new BigNumber(28)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw210.length); _i0_37++) {
          _nw210[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1313_v42 = _nw210;
        let _1315_v43;
        _1315_v43 = _dafny.Seq.UnicodeFromString("qwwlgc");
        let _index226 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1313_v42).length));
        (_1313_v42)[_index226] = _1315_v43;
        let _index227 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1313_v42).length));
        (_1313_v42)[_index227] = (((_this).f6) ? (_1315_v43) : (_1315_v43));
        let _1316_v44;
        let _nw211 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
        _1316_v44 = _nw211;
        let _1317_v45;
        _1317_v45 = _dafny.Seq.of(new BigNumber((_1315_v43).length), new BigNumber(776), p0, new BigNumber(-193), new BigNumber(933));
        let _index228 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1316_v44).length));
        (_1316_v44)[_index228] = _dafny.Seq.Concat(_1317_v45, _1317_v45);
        let _1318_v46;
        _1318_v46 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1277_v11).length), p0), _1317_v45));
        let _1319_v47;
        let _nw212 = Array((new BigNumber(8)).toNumber());
        _nw212[(_dafny.ZERO).toNumber()] = new BigNumber(929);
        _nw212[(_dafny.ONE).toNumber()] = p0;
        _nw212[(new BigNumber(2)).toNumber()] = p0;
        _nw212[(new BigNumber(3)).toNumber()] = p0;
        _nw212[(new BigNumber(4)).toNumber()] = p0;
        _nw212[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw212[(new BigNumber(6)).toNumber()] = p0;
        _nw212[(new BigNumber(7)).toNumber()] = new BigNumber((_1315_v43).length);
        _1319_v47 = _nw212;
        let _1320_v48;
        _1320_v48 = _dafny.Seq.of(_1319_v47, _1319_v47);
        let _index229 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1316_v44).length));
        (_1316_v44)[_index229] = _dafny.Seq.update((_1318_v46)[_module.__default.safeIndex(p0, new BigNumber((_1318_v46).length))], _module.__default.safeIndex((new BigNumber((_1320_v48).length)).plus((_dafny.ZERO).minus((_this).f9)), new BigNumber(((_1318_v46)[_module.__default.safeIndex(p0, new BigNumber((_1318_v46).length))]).length)), _module.__default.safeModuloInt(p0, new BigNumber(-440)));
        let _1321_v49;
        _1321_v49 = _dafny.MultiSet.fromElements(_1301_v32);
        let _1322_v50;
        _1322_v50 = _dafny.Seq.of(_1321_v49, _1321_v49);
        let _1323_v51;
        _1323_v51 = _module.D17.create_DC46(_1321_v49);
        let _1324_v52;
        let _nw213 = Array((new BigNumber(23)).toNumber());
        _nw213[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1301_v32));
        _nw213[(_dafny.ONE).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(2)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), ((_1325_v42) => function (_1326_i5) {
          return _dafny.MultiSet.FromArray((_1325_v42)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_1325_v42).length))]);
        })(_1313_v42));
        _nw213[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_1321_v49, _1321_v49);
        _nw213[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_1322_v50, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(p3, _1276_v10.f8)).cardinality()), new BigNumber((_1322_v50).length)), _1321_v49);
        _nw213[(new BigNumber(6)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1301_v32, _1301_v32), _1321_v49, _1321_v49, _1321_v49, _1321_v49);
        _nw213[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1322_v50, _1322_v50);
        _nw213[(new BigNumber(9)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1322_v50, _dafny.Seq.of(_module.__default.fm33(false, (_this).f9, globalState), _1321_v49));
        _nw213[(new BigNumber(11)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm34(globalState), _dafny.Seq.of(_1321_v49));
        _nw213[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1321_v49, _1321_v49), _dafny.Seq.update(_1322_v50, _module.__default.safeIndex((_this).f9, new BigNumber((_1322_v50).length)), _dafny.MultiSet.fromElements(_1301_v32)));
        _nw213[(new BigNumber(14)).toNumber()] = (((_this).f10) ? (_dafny.Seq.of(_1321_v49)) : (_1322_v50));
        _nw213[(new BigNumber(15)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(16)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm34(globalState), _1322_v50);
        _nw213[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1321_v49, _1321_v49, _dafny.MultiSet.fromElements(_1301_v32)), _1322_v50);
        _nw213[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1321_v49, (_1323_v51).dtor_cf73, _dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), _1301_v32, _1301_v32), _1321_v49, _dafny.MultiSet.fromElements(_1301_v32)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(130)), ((_1327_v49) => function (_1328_i6) {
          return _1327_v49;
        })(_1321_v49)));
        _nw213[(new BigNumber(20)).toNumber()] = _module.__default.fm34(globalState);
        _nw213[(new BigNumber(21)).toNumber()] = _1322_v50;
        _nw213[(new BigNumber(22)).toNumber()] = (((_1276_v10).f6) ? (_1322_v50) : (_1322_v50));
        _1324_v52 = _nw213;
        let _index230 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1324_v52).length));
        (_1324_v52)[_index230] = _dafny.Seq.update(_dafny.Seq.of((_1321_v49).update(_1301_v32, _module.__default.abs((_1317_v45)[_module.__default.safeIndex(new BigNumber(-144), new BigNumber((_1317_v45).length))])), _1321_v49, _1321_v49, _1321_v49), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of((_1321_v49).update(_1301_v32, _module.__default.abs((_1317_v45)[_module.__default.safeIndex(new BigNumber(-144), new BigNumber((_1317_v45).length))])), _1321_v49, _1321_v49, _1321_v49)).length)), _1321_v49);
        let _1329_v53;
        _1329_v53 = _dafny.Seq.of(_1322_v50);
        let _index231 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1324_v52).length));
        let _rhs257 = (_this).f9;
        let _rhs258 = (_1329_v53)[_module.__default.safeIndex((_this).f9, new BigNumber((_1329_v53).length))];
        let _rhs259 = new BigNumber((_1317_v45).length);
        let _lhs219 = globalState;
        let _lhs220 = _1324_v52;
        let _lhs221 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1324_v52).length));
        let _lhs222 = globalState;
        _lhs219.f0 = _rhs257;
        _lhs220[_lhs221] = _rhs258;
        _lhs222.f0 = _rhs259;
      } else {
        let _1330_v54;
        let _nw214 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1330_v54 = _nw214;
        let _index232 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
        (_1330_v54)[_index232] = p0;
        let _index233 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
        let _rhs260 = _this.f8;
        let _rhs261 = !((_this).f10);
        let _rhs262 = p0;
        let _lhs223 = _this;
        let _lhs224 = _1330_v54;
        let _lhs225 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
        r0 = _rhs260;
        _lhs223.f8 = _rhs261;
        _lhs224[_lhs225] = _rhs262;
        let _source19 = _module.D0.create_DC1();
        if (_source19.is_DC1) {
          (_this).f8 = (_1276_v10).f6;
          let _1331_v55;
          _1331_v55 = _dafny.Seq.UnicodeFromString("rq");
          let _1332_v56;
          _1332_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_1331_v55);
          let _1333_v57;
          let _nw215 = new _module.C1();
          _nw215.__ctor(_1276_v10.f8, (((_1332_v56).contains(_1276_v10.f7)) ? ((_1332_v56).get(_1276_v10.f7)) : (_1331_v55)), _this.f8, ((_1276_v10).f5).Intersect(p1));
          _1333_v57 = _nw215;
          let _1334_v58;
          _1334_v58 = _dafny.Seq.of(!(p3));
          let _index234 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
          (_1330_v54)[_index234] = _module.__default.safeModuloInt(new BigNumber((_1334_v58).length), (_this).f9);
          let _arr53 = _1276_v10.f7;
          let _index235 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_1276_v10.f7).length));
          _arr53[_index235] = !(!(_1276_v10.f8));
          let _1335_v59;
          _1335_v59 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1336_v60;
          _1336_v60 = _dafny.Seq.of(_dafny.Seq.update((_1333_v57).f13, _module.__default.safeIndex((_this).f9, new BigNumber(((_1333_v57).f13).length)), _1335_v59));
          let _1337_v61;
          _1337_v61 = _dafny.Seq.of(_1277_v11, _1277_v11, _1277_v11, _1277_v11);
          let _1338_v62;
          _1338_v62 = _dafny.Seq.of((_this).f9);
          let _arr54 = _1276_v10.f7;
          let _index236 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_1276_v10.f7).length));
          let _rhs263 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update((_1336_v60)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1337_v61, _module.__default.safeIndex((_this).f9, new BigNumber((_1337_v61).length)), _1277_v11))).cardinality()), new BigNumber((_1336_v60).length))], _module.__default.safeIndex((_1330_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length))], new BigNumber(((_1336_v60)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1337_v61, _module.__default.safeIndex((_this).f9, new BigNumber((_1337_v61).length)), _1277_v11))).cardinality()), new BigNumber((_1336_v60).length))]).length)), _1335_v59), (_1333_v57).f13)).length);
          let _rhs264 = p3;
          let _rhs265 = !(!(((_this).f9).plus((_this).f9)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1338_v62).length)))));
          let _lhs226 = globalState;
          let _lhs227 = _1276_v10.f7;
          let _lhs228 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_1276_v10.f7).length));
          _lhs226.f0 = _rhs263;
          r2 = _rhs264;
          _lhs227[_lhs228] = _rhs265;
        } else {
          let _1339___mcc_h0 = (_source19).cf0;
          let _1340_cf0 = _1339___mcc_h0;
          (_this).f8 = (((_1276_v10).f6) ? ((_this).f10) : (((_this).f6) === ((_this).f10)));
          let _1341_v63;
          _1341_v63 = _dafny.Seq.of(_1330_v54);
          let _index237 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
          let _index238 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
          let _rhs266 = _this.f8;
          let _rhs267 = (_this).f9;
          let _rhs268 = ((_this).f9).minus((_dafny.ZERO).minus((_1330_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length))]));
          let _rhs269 = _1340_cf0;
          let _rhs270 = (_1341_v63)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_1330_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length))], (_1330_v54)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length))]), new BigNumber((_1341_v63).length))];
          let _lhs229 = _this;
          let _lhs230 = _1330_v54;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
          let _lhs232 = _1330_v54;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_1330_v54).length));
          _lhs229.f8 = _rhs266;
          _lhs230[_lhs231] = _rhs267;
          r1 = _rhs268;
          _lhs232[_lhs233] = _rhs269;
          _1330_v54 = _rhs270;
          let _1342_v64;
          _1342_v64 = _module.D0.create_DC1();
          let _1343_v65;
          _1343_v65 = _dafny.Seq.of((_1276_v10).f6, (new BigNumber(364)).isLessThanOrEqualTo(p0), (_1340_cf0).isLessThan((_dafny.ZERO).minus(p0)));
          let _rhs271 = _1342_v64;
          let _rhs272 = _dafny.Seq.Concat(_1343_v65, _1343_v65);
          let _rhs273 = (_this).f9;
          let _lhs234 = globalState;
          _1342_v64 = _rhs271;
          _1343_v65 = _rhs272;
          _lhs234.f0 = _rhs273;
          r1 = _module.__default.safeDivisionInt(_1340_cf0, p0);
        }
        let _1344_v66;
        _1344_v66 = _dafny.Seq.of(_dafny.Seq.of(p1), _dafny.Seq.of((_1276_v10).f5));
        let _1345_v67;
        _1345_v67 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1346_v68;
        _1346_v68 = _module.D10.create_DC31((_1344_v66)[_module.__default.safeIndex((_this).f9, new BigNumber((_1344_v66).length))], _1345_v67);
        _1346_v68 = _1346_v68;
        let _1347_v69;
        _1347_v69 = _dafny.Seq.UnicodeFromString("mu");
        let _1348_v70;
        _1348_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_1347_v69).length));
        let _1349_v71;
        _1349_v71 = _dafny.Set.fromElements((_1276_v10).f6);
        let _1350_v72;
        let _nw216 = new _module.C3();
        _nw216.__ctor((_1276_v10).fm1((_this).fm3(globalState), new BigNumber((_1348_v70).length), new BigNumber((_1349_v71).length), globalState), _dafny.MultiSet.fromElements((_this).f6));
        _1350_v72 = _nw216;
        _1350_v72 = _1350_v72;
        let _1351_v73;
        _1351_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1264_v4).f11);
        let _1352_v74;
        let _nw217 = new _module.C0();
        _nw217.__ctor((((_1351_v73).contains((_1276_v10).f6)) ? ((_1351_v73).get((_1276_v10).f6)) : (_dafny.Map.Empty.slice().updateUnsafe(_1262_v2,(_1276_v10).f6))));
        _1352_v74 = _nw217;
      }
      let _1353_v75;
      _1353_v75 = _dafny.Seq.of(_1276_v10.f8);
      let _1354_v76;
      _1354_v76 = _module.D6.create_DC17(_1353_v75);
      let _1355_v77;
      _1355_v77 = _module.D11.create_DC33(_1276_v10.f8, false);
      let _pat_let_tv31 = _1276_v10;
      let _pat_let_tv32 = p3;
      let _1356_v78;
      _1356_v78 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let36_0) {
        return function (_1357_dt__update__tmp_h0) {
          return function (_pat_let37_0) {
            return function (_1358_dt__update_hcf24_h0) {
              return _module.D6.create_DC17(_1358_dt__update_hcf24_h0);
            }(_pat_let37_0);
          }(_dafny.Seq.of(_pat_let_tv31.f8, _pat_let_tv32));
        }(_pat_let36_0);
      }(_1354_v76),_1355_v77);
      _1356_v78 = (_1356_v78).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC17(_1353_v75),_1355_v77));
      let _1359_i7;
      _1359_i7 = _dafny.ZERO;
      L12: {
        while ((_this).f6) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1359_i7)) {
              break L12;
            }
            _1359_i7 = (_1359_i7).plus(_dafny.ONE);
            let _1360_v79;
            let _init38 = function (_1361_i8) {
              return (_1361_i8).minus((_this).f9);
            };
            let _nw218 = Array((new BigNumber(15)).toNumber());
            for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw218.length); _i0_38++) {
              _nw218[_i0_38] = _init38(new BigNumber(_i0_38));
            }
            _1360_v79 = _nw218;
            let _index239 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length));
            (_1360_v79)[_index239] = p0;
            let _index240 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length));
            (_1360_v79)[_index240] = _module.__default.safeModuloInt(p0, p0);
            let _1362_v80;
            _1362_v80 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1353_v75));
            let _1363_v81;
            _1363_v81 = _module.D10.create_DC31(_1362_v80, new _dafny.CodePoint('q'.codePointAt(0)));
            let _1364_v82;
            _1364_v82 = _dafny.Seq.UnicodeFromString("hm");
            if (!(!_dafny.Seq.contains(_1364_v82, (_1363_v81).dtor_cf50))) {
              let _1365_v83;
              _1365_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("uexbyodl"),(_this).f6);
              r1 = (_dafny.ZERO).minus((_this).fm4((_this).f9, (_1360_v79)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length))], _1365_v83, (_dafny.ZERO).minus(new BigNumber((_1364_v82).length)), globalState));
              (_1276_v10).f8 = p3;
              (globalState).f0 = p0;
              let _arr55 = _1276_v10.f7;
              let _index241 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1276_v10.f7).length));
              _arr55[_index241] = p3;
              let _arr56 = _1276_v10.f7;
              let _index242 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1276_v10.f7).length));
              _arr56[_index242] = true;
              (_1276_v10).m2(_module.__default.fm16((_1276_v10).f6, _1262_v2, _1364_v82, globalState), globalState);
            } else {
              (globalState).f0 = p0;
              r2 = (_this).f10;
              let _1366_v84;
              _1366_v84 = _dafny.MultiSet.fromElements(_1360_v79);
              let _1367_v85;
              _1367_v85 = _module.D18.create_DC48(_1366_v84);
              _1366_v84 = (_1366_v84).Intersect((_1366_v84).Intersect((_1367_v85).dtor_cf77));
              (_1276_v10).f8 = !((_this).f10);
              let _1368_v86;
              let _nw219 = Array((new BigNumber(24)).toNumber()).fill(_module.D4.Default());
              _1368_v86 = _nw219;
              let _1369_v87;
              _1369_v87 = new _dafny.CodePoint('c'.codePointAt(0));
              let _1370_v88;
              let _nw220 = Array((new BigNumber(22)).toNumber());
              _nw220[(_dafny.ZERO).toNumber()] = _1369_v87;
              _nw220[(_dafny.ONE).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(2)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(3)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(4)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
              _nw220[(new BigNumber(6)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
              _nw220[(new BigNumber(8)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(9)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(10)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(11)).toNumber()] = (_1363_v81).dtor_cf50;
              _nw220[(new BigNumber(12)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(13)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(14)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(15)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(16)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(17)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(18)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(19)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(20)).toNumber()] = _1369_v87;
              _nw220[(new BigNumber(21)).toNumber()] = _1369_v87;
              _1370_v88 = _nw220;
              let _1371_v89;
              let _nw221 = Array((new BigNumber(25)).toNumber());
              _nw221[(_dafny.ZERO).toNumber()] = _1370_v88;
              _nw221[(_dafny.ONE).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(2)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(3)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(4)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(5)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(6)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(7)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(8)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(9)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(10)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(11)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(12)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(13)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(14)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(15)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(16)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(17)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(18)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(19)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(20)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(21)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(22)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(23)).toNumber()] = _1370_v88;
              _nw221[(new BigNumber(24)).toNumber()] = _1370_v88;
              _1371_v89 = _nw221;
              let _1372_v90;
              _1372_v90 = _module.D4.create_DC12(_1371_v89);
              let _1373_v91;
              _1373_v91 = _module.D4.create_DC14(_1372_v90);
              let _index243 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1368_v86).length));
              (_1368_v86)[_index243] = _1373_v91;
              let _index244 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1368_v86).length));
              let _rhs274 = _1373_v91;
              let _rhs275 = _1262_v2;
              let _rhs276 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("lsl"), _module.__default.safeIndex((_1360_v79)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length))], new BigNumber((_dafny.Seq.UnicodeFromString("lsl")).length)), _1369_v87)).length)).multipliedBy((_1360_v79)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length))]);
              let _lhs235 = _1368_v86;
              let _lhs236 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1368_v86).length));
              _lhs235[_lhs236] = _rhs274;
              _1262_v2 = _rhs275;
              r1 = _rhs276;
            }
            let _1374_v92;
            let _nw222 = new _module.C4();
            _nw222.__ctor(_this.f7, _dafny.Seq.contains(_1353_v75, (_this).f6), (_this).f10, (_this).f5);
            _1374_v92 = _nw222;
            let _1375_v93;
            let _nw223 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
            _1375_v93 = _nw223;
            let _index245 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1375_v93).length));
            (_1375_v93)[_index245] = _dafny.Seq.of(new BigNumber(412));
            let _1376_v94;
            _1376_v94 = _dafny.Seq.of((_this).f9);
            let _1377_v95;
            _1377_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1364_v82,p3);
            let _index246 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1375_v93).length));
            let _rhs277 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1376_v94, _dafny.Seq.Create(_module.__default.abs(new BigNumber(708)), function (_1378_i9) {
              return (_this).f9;
            })), _1376_v94);
            let _rhs278 = (_module.__default.safeDivisionInt(new BigNumber(194), (_1276_v10).fm4((_1360_v79)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_1360_v79).length))], (_this).f9, _1377_v95, p0, globalState))).isLessThan(new BigNumber(752));
            let _rhs279 = _1364_v82;
            let _rhs280 = _1276_v10.f7;
            let _lhs237 = _1375_v93;
            let _lhs238 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1375_v93).length));
            let _lhs239 = _1276_v10;
            _lhs237[_lhs238] = _rhs277;
            r0 = _rhs278;
            _1364_v82 = _rhs279;
            _lhs239.f7 = _rhs280;
          }
        }
      }
      r0 = (_1276_v10).f6;
      r1 = (_this).f9;
      r2 = (_1276_v10).f6;
      return [r0, r1, r2];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f3 = _dafny.Set.Empty;
      this.f4 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f3, f4) {
      let _this = this;
      (_this).f3 = f3;
      (_this).f4 = f4;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!_dafny.areEqual(_dafny.Seq.UnicodeFromString("adv"), _dafny.Seq.UnicodeFromString("iy"))) {
        if (false) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(731)), function (_1379_i0) {
            return _this.f4;
          })).length)))).length);
        } else {
          return new BigNumber(737);
        }
      } else {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("sayxmsb")).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()));
      }
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1380_v0;
      _1380_v0 = false;
      let _1381_i0;
      _1381_i0 = _dafny.ZERO;
      L13: {
        while (_1380_v0) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1381_i0)) {
              break L13;
            }
            _1381_i0 = (_1381_i0).plus(_dafny.ONE);
            let _1382_v1;
            _1382_v1 = _dafny.Seq.of((p0).minus(p2));
            r2 = (_1382_v1)[_module.__default.safeIndex(p0, new BigNumber((_1382_v1).length))];
            let _1383_v2;
            let _nw224 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
            _1383_v2 = _nw224;
            let _index247 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1383_v2).length));
            (_1383_v2)[_index247] = p0;
            let _1384_v3;
            let _init39 = ((_1385_v0) => function (_1386_i1) {
              return _1385_v0;
            })(_1380_v0);
            let _nw225 = Array((new BigNumber(5)).toNumber());
            for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw225.length); _i0_39++) {
              _nw225[_i0_39] = _init39(new BigNumber(_i0_39));
            }
            _1384_v3 = _nw225;
            let _index248 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1383_v2).length));
            (_1383_v2)[_index248] = p2;
            let _1387_v4;
            _1387_v4 = _dafny.Set.fromElements(_1384_v3, _1384_v3);
            let _index249 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1383_v2).length));
            let _index250 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1383_v2).length));
            let _rhs281 = _1383_v2;
            let _rhs282 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("gutgmh"))).length), p0);
            let _rhs283 = _1384_v3;
            let _rhs284 = _module.__default.safeModuloInt(p0, new BigNumber(((_1387_v4).Intersect(_1387_v4)).length));
            let _lhs240 = _1383_v2;
            let _lhs241 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1383_v2).length));
            let _lhs242 = _1383_v2;
            let _lhs243 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1383_v2).length));
            _1383_v2 = _rhs281;
            _lhs240[_lhs241] = _rhs282;
            _1384_v3 = _rhs283;
            _lhs242[_lhs243] = _rhs284;
            let _1388_v5;
            _1388_v5 = _dafny.Seq.of(p1);
            let _1389_v6;
            let _nw226 = new _module.C4();
            _nw226.__ctor(_1384_v3, _module.__default.fm35(_1388_v5, _1380_v0, globalState), _1380_v0, _dafny.MultiSet.fromElements(_1380_v0));
            _1389_v6 = _nw226;
            let _1390_v7;
            _1390_v7 = _dafny.Set.fromElements(_1380_v0, (_1389_v6).fm3(globalState));
            let _1391_v8;
            _1391_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1390_v7,!(!(true)));
            _1391_v8 = (_1391_v8).update(_1390_v7, _1380_v0);
          }
        }
      }
      if (_module.__default.fm35(_dafny.Seq.of(p1), !(_1380_v0), globalState)) {
        let _1392_v9;
        _1392_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _1393_v10;
        _1393_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1392_v9,_1380_v0);
        let _1394_v11;
        _1394_v11 = _dafny.Seq.of(_1393_v10);
        _1380_v0 = (_this.f3).contains(new BigNumber(((_1394_v11)[_module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_1394_v11).length))]).length));
        let _1395_v12;
        let _nw227 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1395_v12 = _nw227;
        let _1396_v13;
        _1396_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1395_v12,_1380_v0);
        _1396_v13 = _1396_v13;
        let _1397_v14;
        _1397_v14 = _dafny.Seq.of(true);
        let _1398_v15;
        _1398_v15 = _module.D18.create_DC49(_dafny.Seq.update(_1397_v14, _module.__default.safeIndex(p0, new BigNumber((_1397_v14).length)), _1380_v0), new BigNumber(460));
        let _1399_v16;
        _1399_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1398_v15,false);
        if (!((((_1399_v16).contains(_1398_v15)) ? ((_1399_v16).get(_1398_v15)) : (_1380_v0)))) {
          (globalState).f0 = p0;
          let _1400_v17;
          let _nw228 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1400_v17 = _nw228;
          let _1401_v18;
          let _nw229 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1401_v18 = _nw229;
          let _index251 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1400_v17).length));
          (_1400_v17)[_index251] = _1401_v18;
          let _1402_v19;
          _1402_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1403_v20;
          _1403_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1402_v19,_1380_v0);
          let _1404_v21;
          let _nw230 = new _module.C0();
          _nw230.__ctor(_1403_v20);
          _1404_v21 = _nw230;
          let _1405_v22;
          _1405_v22 = _module.D6.create_DC18(_1404_v21, _1397_v14, p0);
          let _1406_v23;
          _1406_v23 = _dafny.Seq.of(new BigNumber(849), new BigNumber(655), (_1405_v22).dtor_cf27, p2);
          let _index252 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1400_v17).length));
          let _rhs285 = new BigNumber((_dafny.Set.fromElements(_1395_v12, _1395_v12)).length);
          let _rhs286 = _1401_v18;
          let _rhs287 = (new BigNumber((_dafny.Seq.Concat(_1406_v23, _dafny.Seq.of(p2))).length)).multipliedBy(new BigNumber((_1397_v14).length));
          let _lhs244 = _1400_v17;
          let _lhs245 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1400_v17).length));
          r2 = _rhs285;
          _lhs244[_lhs245] = _rhs286;
          r2 = _rhs287;
          r1 = _1380_v0;
          r1 = _1380_v0;
          let _1407_v25;
          _1407_v25 = _module.D2.create_DC7(_this.f4, _1402_v19, (_dafny.ZERO).minus(p2));
          let _1408_v26;
          _1408_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1398_v15);
          let _1409_v27;
          _1409_v27 = _dafny.Seq.of(function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of ((_1407_v25).dtor_cf7).Keys.Elements) {
              let _1410_v24 = _compr_53;
              if (((_1407_v25).dtor_cf7).contains(_1410_v24)) {
                _coll53.push([(_1410_v24).multipliedBy(p0),_module.D18.create_DC49(_1397_v14, new BigNumber((_1397_v14).length))]);
              }
            }
            return _coll53;
          }(), _1408_v26);
          let _pat_let_tv33 = p0;
          let _1411_v28;
          _1411_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_dafny.Map.Empty.slice().updateUnsafe(p0,function (_pat_let38_0) {
            return function (_1412_dt__update__tmp_h0) {
              return function (_pat_let39_0) {
                return function (_1413_dt__update_hcf79_h0) {
                  return _module.D18.create_DC49((_1412_dt__update__tmp_h0).dtor_cf78, _1413_dt__update_hcf79_h0);
                }(_pat_let39_0);
              }(_pat_let_tv33);
            }(_pat_let38_0);
          }(_1398_v15)));
          (_this).f4 = (p1)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1409_v27, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(829),_1398_v15), (((_1411_v28).contains(_this.f3)) ? ((_1411_v28).get(_this.f3)) : (_1408_v26))))).length)), new BigNumber((p1).length))];
        } else {
          let _1414_v29;
          _1414_v29 = _dafny.Seq.UnicodeFromString("sn");
          let _1415_v30;
          _1415_v30 = _dafny.Seq.of(p1, p1, p1, _1414_v29);
          _1414_v29 = _dafny.Seq.Concat((_1415_v30)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_1414_v29).length), p0, p2, p0))).cardinality()), new BigNumber((_1415_v30).length))], p1);
          (globalState).f0 = p0;
          let _1416_v32;
          _1416_v32 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Seq.UnicodeFromString("dtk")).length));
          let _1417_v33;
          _1417_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1416_v32,p2);
          let _1418_v34;
          let _nw231 = new _module.C0();
          _nw231.__ctor(function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of (_1417_v33).Keys.Elements) {
              let _1419_v31 = _compr_54;
              if ((_1417_v33).contains(_1419_v31)) {
                _coll54.push([_1419_v31,!(_1380_v0)]);
              }
            }
            return _coll54;
          }());
          _1418_v34 = _nw231;
          let _index253 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1395_v12).length));
          (_1395_v12)[_index253] = _1380_v0;
          let _index254 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1395_v12).length));
          (_1395_v12)[_index254] = false;
          let _rhs288 = (_1395_v12)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1395_v12).length))];
          let _rhs289 = _1380_v0;
          let _rhs290 = (_1395_v12)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1395_v12).length))];
          let _rhs291 = _this.f4;
          let _lhs246 = _this;
          _1380_v0 = _rhs288;
          _1380_v0 = _rhs289;
          r1 = _rhs290;
          _lhs246.f4 = _rhs291;
        }
        let _1420_v35;
        _1420_v35 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _1421_v36;
        _1421_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1420_v35,_1380_v0);
        let _1422_v37;
        let _nw232 = new _module.C0();
        _nw232.__ctor(_1421_v36);
        _1422_v37 = _nw232;
        let _1423_v38;
        _1423_v38 = _dafny.Set.fromElements(_1422_v37);
        let _1424_v39;
        _1424_v39 = _dafny.Map.Empty.slice().updateUnsafe(!(p2).isEqualTo(new BigNumber((_this.f3).length)),new BigNumber((_1423_v38).length));
        _1424_v39 = (_1424_v39).update((_1422_v37).fm10(new BigNumber((p1).length), globalState), p0);
        let _1425_v40;
        let _nw233 = Array((new BigNumber(17)).toNumber());
        _1425_v40 = _nw233;
        let _1426_v41;
        _1426_v41 = _module.D1.create_DC2(_1380_v0);
        let _1427_v42;
        _1427_v42 = _dafny.MultiSet.fromElements(_1380_v0, _1380_v0, false);
        let _1428_v43;
        let _nw234 = new _module.C3();
        _nw234.__ctor((_1426_v41).dtor_cf1, _1427_v42);
        _1428_v43 = _nw234;
        let _index255 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1425_v40).length));
        (_1425_v40)[_index255] = _1428_v43;
        let _index256 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1425_v40).length));
        (_1425_v40)[_index256] = _1428_v43;
      } else {
        let _1429_v44;
        let _init40 = ((_1430_v0) => function (_1431_i2) {
          return _1430_v0;
        })(_1380_v0);
        let _nw235 = Array((new BigNumber(26)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw235.length); _i0_40++) {
          _nw235[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1429_v44 = _nw235;
        let _1432_v45;
        _1432_v45 = _dafny.Seq.of(p1, p1, p1);
        let _1433_v46;
        _1433_v46 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm35(_1432_v45, _1380_v0, globalState));
        let _1434_v47;
        _1434_v47 = _dafny.MultiSet.fromElements(true);
        let _1435_v48;
        let _nw236 = new _module.C5();
        _nw236.__ctor((_dafny.ZERO).minus(p2), _1380_v0, _1429_v44, _1380_v0, !((((_1433_v46).contains(_1380_v0)) ? ((_1433_v46).get(_1380_v0)) : (_1380_v0))), (_1434_v47).Intersect(_1434_v47));
        _1435_v48 = _nw236;
        let _rhs292 = (_module.D8.create_DC25(_1380_v0, (_1435_v48).f10)).dtor_cf39;
        let _rhs293 = p0;
        _1380_v0 = _rhs292;
        r2 = _rhs293;
        if ((_dafny.Set.fromElements(_1429_v44, _1429_v44, _1429_v44)).IsDisjointFrom(_dafny.Set.fromElements(_1429_v44, _1429_v44, _1429_v44))) {
          let _1436_v49;
          _1436_v49 = _dafny.Set.fromElements((_1435_v48).f10, _1380_v0);
          let _1437_v50;
          _1437_v50 = _module.D2.create_DC7(_this.f4, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1436_v49).length),p2), (_1435_v48).f9);
          let _1438_v51;
          _1438_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          let _1439_v52;
          let _nw237 = Array((new BigNumber(29)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = new BigNumber(563);
          _nw237[(_dafny.ONE).toNumber()] = new BigNumber(288);
          _nw237[(new BigNumber(2)).toNumber()] = new BigNumber(712);
          _nw237[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_1435_v48).f9);
          _nw237[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(523));
          _nw237[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((_1435_v48).f9, (_1435_v48).f9);
          _nw237[(new BigNumber(6)).toNumber()] = p2;
          _nw237[(new BigNumber(7)).toNumber()] = (_1437_v50).dtor_cf8;
          _nw237[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_module.__default.fm17(globalState)).length), p2);
          _nw237[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_1435_v48).f9);
          _nw237[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pp")).length)), p2);
          _nw237[(new BigNumber(11)).toNumber()] = ((_1380_v0) ? (p0) : (new BigNumber((_1438_v51).length)));
          _nw237[(new BigNumber(12)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(13)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(14)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(15)).toNumber()] = p0;
          _nw237[(new BigNumber(16)).toNumber()] = new BigNumber(-644);
          _nw237[(new BigNumber(17)).toNumber()] = p2;
          _nw237[(new BigNumber(18)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(19)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(20)).toNumber()] = new BigNumber(-83);
          _nw237[(new BigNumber(21)).toNumber()] = ((_1435_v48).f9).plus(new BigNumber((_dafny.Set.fromElements((_1435_v48).f9)).length));
          _nw237[(new BigNumber(22)).toNumber()] = p0;
          _nw237[(new BigNumber(23)).toNumber()] = p2;
          _nw237[(new BigNumber(24)).toNumber()] = p0;
          _nw237[(new BigNumber(25)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(26)).toNumber()] = (_1435_v48).f9;
          _nw237[(new BigNumber(27)).toNumber()] = (p0).plus(p2);
          _nw237[(new BigNumber(28)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(77), p0);
          _1439_v52 = _nw237;
          let _index257 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length));
          (_1439_v52)[_index257] = p2;
          let _1440_v53;
          _1440_v53 = _dafny.Seq.of((_1435_v48).f10);
          let _1441_v54;
          _1441_v54 = _dafny.Map.Empty.slice().updateUnsafe((_1435_v48).f10,_dafny.Seq.Concat(_1440_v53, _1440_v53));
          let _index258 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length));
          let _rhs294 = p0;
          let _rhs295 = _module.__default.safeDivisionInt((_this).fm0(_1432_v45, (_1435_v48).f9, _1380_v0, !(!((_1435_v48).f10)), globalState), new BigNumber(845));
          let _rhs296 = (_1441_v54).Merge(_1441_v54);
          let _rhs297 = _module.__default.fm35(_dafny.Seq.Concat(_dafny.Seq.of(p1, p1), _1432_v45), _1380_v0, globalState);
          let _lhs247 = globalState;
          let _lhs248 = _1439_v52;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length));
          _lhs247.f0 = _rhs294;
          _lhs248[_lhs249] = _rhs295;
          _1441_v54 = _rhs296;
          _1380_v0 = _rhs297;
          let _1442_v55;
          _1442_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cbep"),_1380_v0);
          let _rhs298 = (_1435_v48).fm4((_1439_v52)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length))], p2, _1442_v55, new BigNumber((p1).length), globalState);
          let _rhs299 = p2;
          r2 = _rhs298;
          r2 = _rhs299;
          let _1443_v56;
          _1443_v56 = _dafny.Seq.UnicodeFromString("qwdwvuxp");
          let _1444_v57;
          let _nw238 = new _module.C4();
          _nw238.__ctor(_1429_v44, _1380_v0, _1380_v0, _1434_v47);
          _1444_v57 = _nw238;
          let _1445_v58;
          _1445_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1444_v57,_module.D7.create_DC21(p2, (_1439_v52)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length))]));
          let _1446_v59;
          _1446_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v58,(_1439_v52)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length))]);
          let _1447_v60;
          _1447_v60 = _dafny.Seq.of((_1435_v48).f9);
          let _rhs300 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-717)), function (_1448_i3) {
            return _this.f4;
          });
          let _rhs301 = (_1440_v53)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1447_v60)).cardinality()), new BigNumber((_1440_v53).length))];
          let _rhs302 = (_1446_v59).Merge((_1446_v59).Merge(_1446_v59));
          _1443_v56 = _rhs300;
          _1380_v0 = _rhs301;
          _1446_v59 = _rhs302;
          let _arr57 = _1444_v57.f7;
          let _index259 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length));
          _arr57[_index259] = ((!(_1444_v57.f8)) ? ((_1435_v48).f10) : ((_1444_v57).f6));
          let _arr58 = _1444_v57.f7;
          let _index260 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length));
          let _rhs303 = !((_1435_v48).f10);
          let _rhs304 = (_1439_v52)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1439_v52).length))];
          let _rhs305 = (_this.f3).IsSubsetOf(((false) ? (_dafny.Set.fromElements(p0, (_1435_v48).f9)) : (function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_1447_v60).Elements) {
              let _1449_v61 = _compr_55;
              if (_dafny.Seq.contains(_1447_v60, _1449_v61)) {
                _coll55.add((_1449_v61).multipliedBy(new BigNumber(133)));
              }
            }
            return _coll55;
          }())));
          let _lhs250 = _1444_v57.f7;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length));
          let _lhs252 = _1444_v57;
          _lhs250[_lhs251] = _rhs303;
          r2 = _rhs304;
          _lhs252.f8 = _rhs305;
          let _1450_v62;
          let _nw239 = new _module.C4();
          _nw239.__ctor(_1429_v44, (_1444_v57.f7)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length))], (_1444_v57.f7)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length))], _1434_v47);
          _1450_v62 = _nw239;
          let _1451_v63;
          _1451_v63 = _dafny.Seq.of(_1450_v62, _1450_v62);
          let _1452_v64;
          _1452_v64 = _dafny.Map.Empty.slice().updateUnsafe(!((_1444_v57).f6),(_1451_v63)[_module.__default.safeIndex(p2, new BigNumber((_1451_v63).length))]);
          let _arr59 = _1444_v57.f7;
          let _index261 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length));
          let _rhs306 = ((_1436_v49).Union(_1436_v49)).IsSubsetOf(_dafny.Set.fromElements(_1380_v0, (_1444_v57).f6, _1444_v57.f8, !((_1435_v48).f10), _1380_v0));
          let _rhs307 = _1452_v64;
          let _rhs308 = (_1447_v60)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1439_v52)).length), new BigNumber((_1447_v60).length))];
          let _lhs253 = _1444_v57.f7;
          let _lhs254 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1444_v57.f7).length));
          _lhs253[_lhs254] = _rhs306;
          _1452_v64 = _rhs307;
          r0 = _rhs308;
        } else {
          _1433_v46 = (_1433_v46).Merge(_1433_v46);
          let _index262 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1429_v44).length));
          (_1429_v44)[_index262] = (_1435_v48).f10;
          let _1453_v65;
          _1453_v65 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1435_v48).f9);
          let _1454_v66;
          _1454_v66 = _module.D7.create_DC21((_1435_v48).f9, (p0).multipliedBy(new BigNumber((_1453_v65).length)));
          let _index263 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1429_v44).length));
          let _rhs309 = ((_1435_v48).f9).isLessThanOrEqualTo(p2);
          let _rhs310 = (_1435_v48).f9;
          let _rhs311 = _1454_v66;
          let _rhs312 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p2));
          let _rhs313 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("jv")), p1), _module.__default.safeIndex(new BigNumber(228), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("jv")), p1)).length)), _this.f4)).length);
          let _lhs255 = _1429_v44;
          let _lhs256 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1429_v44).length));
          let _lhs257 = globalState;
          let _lhs258 = globalState;
          _lhs255[_lhs256] = _rhs309;
          r0 = _rhs310;
          _1454_v66 = _rhs311;
          _lhs257.f0 = _rhs312;
          _lhs258.f0 = _rhs313;
          let _1455_v67;
          let _init41 = ((_1456_v48) => function (_1457_i4) {
            return (_1456_v48).f10;
          })(_1435_v48);
          let _nw240 = Array((new BigNumber(26)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw240.length); _i0_41++) {
            _nw240[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1455_v67 = _nw240;
          let _1458_v68;
          let _nw241 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1458_v68 = _nw241;
          let _1459_v69;
          _1459_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1458_v68,_1455_v67);
          let _1460_v70;
          _1460_v70 = _dafny.Seq.of(_1455_v67, (((_1459_v69).contains(_1458_v68)) ? ((_1459_v69).get(_1458_v68)) : (_1455_v67)), _1455_v67);
          let _1461_v71;
          _1461_v71 = _dafny.MultiSet.fromElements(new BigNumber(391));
          let _rhs314 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), (_1435_v48).f9);
          let _rhs315 = !((_module.__default.fm12(globalState)).IsDisjointFrom((_1461_v71).update(new BigNumber(336), _module.__default.abs(p0))));
          let _rhs316 = _1460_v70;
          let _rhs317 = ((_1435_v48).f10) || (!((_1435_v48).f10));
          let _rhs318 = (_1435_v48).f9;
          r2 = _rhs314;
          _1380_v0 = _rhs315;
          _1460_v70 = _rhs316;
          _1380_v0 = _rhs317;
          r2 = _rhs318;
          let _1462_v72;
          _1462_v72 = _module.D1.create_DC2((_1435_v48).f10);
          r1 = (_1462_v72).dtor_cf1;
          r0 = _module.__default.safeModuloInt((p0).plus(p2), new BigNumber((_dafny.Seq.Concat(p1, p1)).length));
        }
        let _1463_v73;
        _1463_v73 = _dafny.Seq.UnicodeFromString("fieugd");
        _1463_v73 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_1464_i5) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }), _module.__default.fm11(_module.__default.fm35(_1432_v45, _1380_v0, globalState), globalState));
        (globalState).f0 = (_1435_v48).f9;
      }
      let _1465_v74;
      let _nw242 = Array((new BigNumber(29)).toNumber()).fill(false);
      _1465_v74 = _nw242;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1465_v74).length))) {
        let _1466_i6 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1466_i6)) && ((_1466_i6).isLessThan(new BigNumber((_1465_v74).length))))) {
          (_1465_v74)[(_1466_i6)] = (_module.D12.create_DC35(_1380_v0)).dtor_cf55;
        }
      }
      let _1467_i7;
      _1467_i7 = _dafny.ZERO;
      L14: {
        while ((_module.__default.safeModuloInt(p2, p0)).isLessThanOrEqualTo(p2)) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1467_i7)) {
              break L14;
            }
            _1467_i7 = (_1467_i7).plus(_dafny.ONE);
            let _1468_v75;
            _1468_v75 = _dafny.Seq.of(_1380_v0, _1380_v0, _1380_v0);
            let _1469_v76;
            _1469_v76 = _dafny.Set.fromElements(_1380_v0, false);
            let _1470_v77;
            _1470_v77 = _dafny.Map.Empty.slice().updateUnsafe((_1469_v76).contains((_1468_v75)[_module.__default.safeIndex(p0, new BigNumber((_1468_v75).length))]),_1380_v0);
            r1 = !((((_1470_v77).contains(_1380_v0)) ? ((_1470_v77).get(_1380_v0)) : (_1380_v0)));
            let _1471_v78;
            _1471_v78 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
            let _1472_v79;
            _1472_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1380_v0,p2);
            _1380_v0 = ((p0).minus((((_1472_v79).contains(_1380_v0)) ? ((_1472_v79).get(_1380_v0)) : (p2)))).isLessThanOrEqualTo((((_1471_v78).contains(p2)) ? ((_1471_v78).get(p2)) : (p0)));
            let _1473_v80;
            _1473_v80 = _module.D2.create_DC5(_1380_v0);
            let _1474_v81;
            _1474_v81 = _module.D2.create_DC8(_1473_v80);
            let _1475_v82;
            let _nw243 = Array((new BigNumber(5)).toNumber());
            _nw243[(_dafny.ZERO).toNumber()] = _1468_v75;
            _nw243[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_1380_v0);
            _nw243[(new BigNumber(2)).toNumber()] = _1468_v75;
            _nw243[(new BigNumber(3)).toNumber()] = _1468_v75;
            _nw243[(new BigNumber(4)).toNumber()] = _1468_v75;
            _1475_v82 = _nw243;
            let _index264 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_1475_v82).length));
            (_1475_v82)[_index264] = _1468_v75;
            let _1476_v83;
            _1476_v83 = _dafny.Seq.UnicodeFromString("pgaotoy");
            let _index265 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_1475_v82).length));
            let _rhs319 = _module.D2.create_DC8(_module.D2.create_DC5(_1380_v0));
            let _rhs320 = _1468_v75;
            let _rhs321 = _1476_v83;
            let _lhs259 = _1475_v82;
            let _lhs260 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_1475_v82).length));
            _1474_v81 = _rhs319;
            _lhs259[_lhs260] = _rhs320;
            _1476_v83 = _rhs321;
            let _1477_v84;
            _1477_v84 = _dafny.MultiSet.fromElements(_1380_v0, _1380_v0, _1380_v0);
            let _1478_v85;
            let _nw244 = new _module.C5();
            _nw244.__ctor(p2, _1380_v0, _1465_v74, _1380_v0, _1380_v0, _1477_v84);
            _1478_v85 = _nw244;
            let _1479_v86;
            _1479_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1478_v85,(_1478_v85).f9);
            (globalState).f0 = (p0).plus((((_1479_v86).contains(_1478_v85)) ? ((_1479_v86).get(_1478_v85)) : (new BigNumber((_module.__default.fm23(p2, globalState)).length))));
          }
        }
      }
      let _1480_v87;
      let _nw245 = new _module.C2();
      _nw245.__ctor(_dafny.MultiSet.fromElements(_1380_v0));
      _1480_v87 = _nw245;
      let _1481_v88;
      _1481_v88 = _dafny.Seq.of(_1380_v0);
      let _1482_v89;
      _1482_v89 = _dafny.MultiSet.fromElements(_1380_v0, _1380_v0, _1380_v0, false);
      if ((_1480_v87).fm1(_1380_v0, ((_1380_v0) ? (p2) : (new BigNumber((_1481_v88).length))), (((_1482_v89).contains(_1380_v0)) ? ((_1482_v89).get(_1380_v0)) : (new BigNumber(979))), globalState)) {
        let _1483_v90;
        _1483_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1380_v0,_1380_v0);
        let _1484_v91;
        _1484_v91 = _dafny.Seq.of((p2).multipliedBy(new BigNumber((_1483_v90).length)), (_dafny.ZERO).minus(p0), p2);
        _1484_v91 = _dafny.Seq.Concat(_1484_v91, _1484_v91);
        let _1485_v92;
        _1485_v92 = _dafny.Map.Empty.slice().updateUnsafe(!(false),p2);
        let _1486_v93;
        _1486_v93 = _dafny.MultiSet.fromElements(p2);
        _1485_v92 = (_1485_v92).update(_1380_v0, (((_1486_v93).contains(p0)) ? ((_1486_v93).get(p0)) : (new BigNumber(369))));
        let _1487_v94;
        _1487_v94 = _dafny.Seq.of(p1);
        let _1488_v95;
        _1488_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1465_v74,p0);
        let _1489_v96;
        let _nw246 = Array((new BigNumber(21)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm18(globalState)).length);
        _nw246[(_dafny.ONE).toNumber()] = (_1484_v91)[_module.__default.safeIndex(p0, new BigNumber((_1484_v91).length))];
        _nw246[(new BigNumber(2)).toNumber()] = p0;
        _nw246[(new BigNumber(3)).toNumber()] = p2;
        _nw246[(new BigNumber(4)).toNumber()] = p0;
        _nw246[(new BigNumber(5)).toNumber()] = p0;
        _nw246[(new BigNumber(6)).toNumber()] = ((false) ? (p2) : (p0));
        _nw246[(new BigNumber(7)).toNumber()] = (_this).fm0(_1487_v94, new BigNumber((_this.f3).length), _1380_v0, _1380_v0, globalState);
        _nw246[(new BigNumber(8)).toNumber()] = p0;
        _nw246[(new BigNumber(9)).toNumber()] = p2;
        _nw246[(new BigNumber(10)).toNumber()] = p2;
        _nw246[(new BigNumber(11)).toNumber()] = ((_1380_v0) ? (p0) : (p2));
        _nw246[(new BigNumber(12)).toNumber()] = p2;
        _nw246[(new BigNumber(13)).toNumber()] = p0;
        _nw246[(new BigNumber(14)).toNumber()] = p0;
        _nw246[(new BigNumber(15)).toNumber()] = new BigNumber((_1488_v95).length);
        _nw246[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(p2, p2);
        _nw246[(new BigNumber(17)).toNumber()] = p2;
        _nw246[(new BigNumber(18)).toNumber()] = p2;
        _nw246[(new BigNumber(19)).toNumber()] = (((_1482_v89).contains(_1380_v0)) ? ((_1482_v89).get(_1380_v0)) : (p2));
        _nw246[(new BigNumber(20)).toNumber()] = (_this).fm0(_dafny.Seq.of(p1), new BigNumber(-966), _1380_v0, _1380_v0, globalState);
        _1489_v96 = _nw246;
        _1489_v96 = _1489_v96;
        let _1490_v97;
        _1490_v97 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1491_v98;
        _1491_v98 = _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1483_v90).length),new BigNumber(-866))).Merge((_1490_v97).update(p2, p0)));
        let _1492_v99;
        _1492_v99 = _dafny.Seq.UnicodeFromString("wxp");
        let _1493_v100;
        _1493_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1380_v0,_1490_v97);
        let _rhs322 = (_1491_v98).Union(((!(_module.__default.fm35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), ((_1494_p1) => function (_1495_i8) {
          return _1494_p1;
        })(p1)), !(_1380_v0), globalState))) ? (_1491_v98) : (_dafny.Set.fromElements(_1490_v97, (((_1493_v100).contains(_1380_v0)) ? ((_1493_v100).get(_1380_v0)) : (_1490_v97))))));
        let _rhs323 = _dafny.Seq.UnicodeFromString("yohtecotu");
        _1491_v98 = _rhs322;
        _1492_v99 = _rhs323;
        let _1496_v101;
        let _nw247 = new _module.C4();
        _nw247.__ctor(_1465_v74, false, _1380_v0, _1482_v89);
        _1496_v101 = _nw247;
        let _1497_v102;
        _1497_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1496_v101,_1486_v93);
        let _1498_v103;
        let _nw248 = new _module.C4();
        _nw248.__ctor(_1465_v74, _1380_v0, (new BigNumber(875)).isLessThan(new BigNumber((_1497_v102).length)), _1482_v89);
        _1498_v103 = _nw248;
        let _rhs324 = _1498_v103.f8;
        let _rhs325 = p0;
        let _rhs326 = _1498_v103;
        let _rhs327 = p0;
        _1380_v0 = _rhs324;
        r0 = _rhs325;
        _1498_v103 = _rhs326;
        r2 = _rhs327;
      } else {
        let _1499_v104;
        let _init42 = ((_1500_p1) => function (_1501_i9) {
          return _1500_p1;
        })(p1);
        let _nw249 = Array((new BigNumber(19)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw249.length); _i0_42++) {
          _nw249[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1499_v104 = _nw249;
        let _index266 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1499_v104).length));
        (_1499_v104)[_index266] = p1;
        let _1502_v105;
        _1502_v105 = _module.D5.create_DC15(p1);
        let _index267 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1499_v104).length));
        (_1499_v104)[_index267] = _dafny.Seq.Concat(_dafny.Seq.Concat((_1502_v105).dtor_cf18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), function (_1503_i10) {
          return _this.f4;
        })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), function (_1504_i11) {
          return _this.f4;
        }));
        _1380_v0 = _1380_v0;
        _1380_v0 = _1380_v0;
        let _1505_v106;
        let _init43 = ((_1506_v0) => function (_1507_i12) {
          return (_1507_i12).multipliedBy(new BigNumber((_dafny.Set.fromElements(_1506_v0, _1506_v0)).length));
        })(_1380_v0);
        let _nw250 = Array((new BigNumber(25)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw250.length); _i0_43++) {
          _nw250[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1505_v106 = _nw250;
        let _index268 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1505_v106).length));
        (_1505_v106)[_index268] = p2;
        let _1508_v107;
        _1508_v107 = _dafny.Seq.of(_1481_v88, _dafny.Seq.of(true), _1481_v88, _1481_v88, _1481_v88);
        let _1509_v108;
        let _nw251 = new _module.C5();
        _nw251.__ctor(p0, _1380_v0, _1465_v74, false, _1380_v0, (_1482_v89).update(_1380_v0, _module.__default.abs(p2)));
        _1509_v108 = _nw251;
        let _1510_v109;
        _1510_v109 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v108,(_1509_v108).f10);
        let _1511_v110;
        _1511_v110 = _dafny.MultiSet.fromElements(_dafny.Seq.of((((_1510_v109).contains(_1509_v108)) ? ((_1510_v109).get(_1509_v108)) : (!(!(!((_1509_v108).f10))))), _1380_v0), _1481_v88);
        let _index269 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1505_v106).length));
        (_1505_v106)[_index269] = new BigNumber(((_dafny.MultiSet.FromArray(_1508_v107)).Difference(_1511_v110)).cardinality());
        r1 = !(!(false) || (!(_1380_v0)));
      }
      let _1512_v111;
      _1512_v111 = _module.D17.create_DC47(_1380_v0, p0, p2);
      let _pat_let_tv34 = p0;
      r0 = (function (_pat_let40_0) {
        return function (_1513_dt__update__tmp_h1) {
          return function (_pat_let41_0) {
            return function (_1514_dt__update_hcf75_h0) {
              return _module.D17.create_DC47((_1513_dt__update__tmp_h1).dtor_cf74, _1514_dt__update_hcf75_h0, (_1513_dt__update__tmp_h1).dtor_cf76);
            }(_pat_let41_0);
          }(_pat_let_tv34);
        }(_pat_let40_0);
      }(_1512_v111)).dtor_cf75;
      let _1515_v112;
      _1515_v112 = _dafny.MultiSet.fromElements(p0);
      let _1516_v113;
      _1516_v113 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1380_v0);
      let _1517_v114;
      _1517_v114 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("uksqkb"), p1);
      r1 = !((((_1516_v113).contains((_this).fm0(_1517_v114, new BigNumber((_dafny.Seq.of(p0, p0, p0)).length), _1380_v0, _1380_v0, globalState))) ? ((_1516_v113).get((_this).fm0(_1517_v114, new BigNumber((_dafny.Seq.of(p0, p0, p0)).length), _1380_v0, _1380_v0, globalState))) : (_1380_v0))) || ((_dafny.MultiSet.fromElements(p2)).IsDisjointFrom(_1515_v112));
      r2 = (_dafny.ZERO).minus(p2);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
