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
    static fm0(p0, p1, p2, p3, globalState) {
      return ((new BigNumber((_dafny.Seq.UnicodeFromString("fbkunt")).length)).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ich")).length), new BigNumber(910))).cardinality())), new BigNumber(208)))).length))).length))).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("cgadqenod")).length));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(!(!(false)))).IsProperSubsetOf(_dafny.Set.fromElements(true, true, !(!(false)))),new BigNumber(646));
    };
    static fm5(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(!(false)));
    };
    static fm6(globalState) {
      return ((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-866), new BigNumber(9))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(-866)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(9)))) {
            _coll0.add((_0_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(_module.D2.create_DC7(_dafny.Seq.of(false)))).cardinality())));
          }
        }
        return _coll0;
      }()).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("k"),new BigNumber(696))).length)))).Difference((_dafny.Set.fromElements(new BigNumber(267))).Intersect(_dafny.Set.fromElements(new BigNumber(305))));
    };
    static fm7(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of((_module.D9.create_DC30(new _dafny.CodePoint('q'.codePointAt(0)), false, new BigNumber(-248), _dafny.Set.fromElements(new BigNumber(588), new BigNumber(628), new BigNumber((_dafny.Seq.UnicodeFromString("slllsh")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber(-740)), new BigNumber((_dafny.Seq.UnicodeFromString("pbe")).length))).dtor_cf54))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(-855))));
    };
    static fm13(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(620)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vsrbg")).length), new BigNumber(617), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_1_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, !(false))).length));
      }), _dafny.Seq.of(new BigNumber(-578))))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),true)).length)), _dafny.Seq.of(new BigNumber(-82), new BigNumber(371))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-616))))));
    };
    static fm14(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-188),new BigNumber(202))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(330),new BigNumber(681))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), function (_2_i0) {
        return new BigNumber(524);
      })).length),new BigNumber((_dafny.Seq.of(false)).length))));
    };
    static fm19(p0, p1, globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))));
    };
    static fm20(p0, p1, p2, globalState) {
      let _source0 = ((false) ? (_module.D7.create_DC23(false, new BigNumber(842), false)) : (_module.D7.create_DC23(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(-164))).length), !(true))));
      if (_source0.is_DC22) {
        let _3___mcc_h0 = (_source0).cf37;
        let _4___mcc_h1 = (_source0).cf38;
        let _5___mcc_h2 = (_source0).cf39;
        let _6___mcc_h3 = (_source0).cf40;
        let _7___mcc_h4 = (_source0).cf41;
        let _8_cf41 = _7___mcc_h4;
        let _9_cf40 = _6___mcc_h3;
        let _10_cf39 = _5___mcc_h2;
        let _11_cf38 = _4___mcc_h1;
        let _12_cf37 = _3___mcc_h0;
        return _module.D0.create_DC0(true);
      } else if (_source0.is_DC23) {
        let _13___mcc_h5 = (_source0).cf42;
        let _14___mcc_h6 = (_source0).cf43;
        let _15___mcc_h7 = (_source0).cf44;
        let _16_cf44 = _15___mcc_h7;
        let _17_cf43 = _14___mcc_h6;
        let _18_cf42 = _13___mcc_h5;
        return _module.D0.create_DC0(_18_cf42);
      } else {
        let _19___mcc_h8 = (_source0).cf36;
        let _20_cf36 = _19___mcc_h8;
        return _module.D0.create_DC0(true);
      }
    };
    static fm21(p0, p1, p2, globalState) {
      let _source1 = _module.D4.create_DC14(true, (_module.D7.create_DC23(true, new BigNumber(781), true)).dtor_cf44);
      if (_source1.is_DC13) {
        let _21___mcc_h0 = (_source1).cf22;
        let _22___mcc_h1 = (_source1).cf23;
        let _23___mcc_h2 = (_source1).cf24;
        let _24___mcc_h3 = (_source1).cf25;
        let _25_cf25 = _24___mcc_h3;
        let _26_cf24 = _23___mcc_h2;
        let _27_cf23 = _22___mcc_h1;
        let _28_cf22 = _21___mcc_h0;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), ((_29_cf23) => function (_30_i0) {
          return _29_cf23;
        })(_27_cf23));
      } else if (_source1.is_DC14) {
        let _31___mcc_h4 = (_source1).cf26;
        let _32___mcc_h5 = (_source1).cf27;
        let _33_cf27 = _32___mcc_h5;
        let _34_cf26 = _31___mcc_h4;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-807)), ((_35_cf26) => function (_36_i1) {
          return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_35_cf26)).length), new BigNumber((function () {
            let _coll1 = new _dafny.Set();
            for (const _compr_1 of _dafny.IntegerRange(new BigNumber(752), new BigNumber(428))) {
              let _37_v0 = _compr_1;
              if (((new BigNumber(752)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(428)))) {
                _coll1.add((_37_v0).minus(new BigNumber(708)));
              }
            }
            return _coll1;
          }()).length))).cardinality()))).length), _dafny.ZERO)).length))).length);
        })(_34_cf26));
      } else if (_source1.is_DC12) {
        let _38___mcc_h6 = (_source1).cf21;
        let _39_cf21 = _38___mcc_h6;
        return _dafny.Seq.of(new BigNumber(106), new BigNumber((_dafny.Set.fromElements(new BigNumber(-163), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(937),new BigNumber(667))).length), new BigNumber(795), new BigNumber((_dafny.Set.fromElements(true, !(true), !(true))).length))).length));
      } else {
        let _40___mcc_h7 = (_source1).cf28;
        let _41_cf28 = _40___mcc_h7;
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fdqjqr")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(141))).cardinality()), new BigNumber((_dafny.Seq.of(!(true), false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("pvuwwxen")).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, !(true), true)).cardinality()))));
      }
    };
    static fm22(p0, p1, p2, p3, globalState) {
      if (!(!(true) || (true))) {
        return (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("nuqceuqr")).length), new BigNumber(-359))).length))).length), new BigNumber(870))).length)))).Union(function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(768), new BigNumber(-704))) {
            let _42_v0 = _compr_2;
            if (((new BigNumber(768)).isLessThanOrEqualTo(_42_v0)) && ((_42_v0).isLessThan(new BigNumber(-704)))) {
              _coll2.add(_module.__default.safeDivisionInt(_42_v0, new BigNumber(236)));
            }
          }
          return _coll2;
        }());
      } else {
        return _dafny.Set.fromElements(new BigNumber(569), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(false)))).cardinality()), new BigNumber((function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)))).Elements) {
            let _43_v1 = _compr_3;
            if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)))).contains(_43_v1)) {
              _coll3.add(_module.__default.safeModuloInt(_43_v1, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(688), new BigNumber(355), new BigNumber((_dafny.Seq.of(!(true), true)).length))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_44_i0) {
                return new BigNumber(384);
              })).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("dp")).length), new BigNumber(-318), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)), new BigNumber(813))).length)))).length)));
            }
          }
          return _coll3;
        }()).length), new BigNumber(-904));
      }
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return true;
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),(_dafny.ZERO).minus(new BigNumber(-778)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),(_module.D18.create_DC56(true, true, new BigNumber((_dafny.Set.fromElements(new BigNumber(754), new BigNumber((_dafny.Seq.UnicodeFromString("kk")).length))).length))).dtor_cf92)).length))));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    static fm26(p0, p1, p2, globalState) {
      let _source2 = _module.D12.create_DC38((_module.D9.create_DC30(new _dafny.CodePoint('y'.codePointAt(0)), true, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-170))).length)), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-48)), new BigNumber(262))).cardinality())), new BigNumber((_dafny.Set.fromElements(true)).length)), new BigNumber(-96))).dtor_cf52);
      if (_source2.is_DC38) {
        let _45___mcc_h0 = (_source2).cf64;
        let _46_cf64 = _45___mcc_h0;
        return _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))));
      } else if (_source2.is_DC39) {
        let _47___mcc_h1 = (_source2).cf65;
        let _48_cf65 = _47___mcc_h1;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(881)), function (_49_i0) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), function (_50_i1) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          });
        });
      } else if (_source2.is_DC37) {
        let _51___mcc_h2 = (_source2).cf63;
        let _52_cf63 = _51___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), function (_53_i2) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(87)), function (_54_i3) {
          return _dafny.Seq.of((_module.D9.create_DC30(new _dafny.CodePoint('j'.codePointAt(0)), true, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), function () {
  let _coll4 = new _dafny.Set();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(750), new BigNumber(208))) {
    let _55_v0 = _compr_4;
    if (((new BigNumber(750)).isLessThanOrEqualTo(_55_v0)) && ((_55_v0).isLessThan(new BigNumber(208)))) {
      _coll4.add(_module.__default.safeModuloInt(_55_v0, new BigNumber((function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length),true)).Keys.Elements) {
          let _56_v1 = _compr_5;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length),true)).contains(_56_v1)) {
            _coll5.add(_module.__default.safeModuloInt(_56_v1, new BigNumber(881)));
          }
        }
        return _coll5;
      }()).length)));
    }
  }
  return _coll4;
}(), new BigNumber(-338))).dtor_cf50);
        }));
      } else {
        let _57___mcc_h3 = (_source2).cf66;
        let _58_cf66 = _57___mcc_h3;
        return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), function (_59_i4) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }));
      }
    };
    static fm27(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vkgblbh"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(122)), function (_60_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }));
    };
    static fm28(p0, globalState) {
      return _module.D1.create_DC4((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),false)).length)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("hyccnriii")).length)), (new BigNumber(-573)).minus(new BigNumber(618)));
    };
    static fm29(globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true, !(true))));
    };
    static fm30(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(-690))).Difference(_dafny.MultiSet.fromElements(new BigNumber(-639), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(false))).cardinality()), new BigNumber(-736), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-596)), function (_61_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("cmgn")).length))).length);
      })).length), new BigNumber(465)));
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ymdokrv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), function (_62_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }))).length)),new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm32(p0, p1, globalState) {
      return function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_dafny.MultiSet.fromElements(false)))).Keys.Elements) {
          let _63_v0 = _compr_6;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_dafny.MultiSet.fromElements(false)))).contains(_63_v0)) {
            _coll6.add(_63_v0);
          }
        }
        return _coll6;
      }();
    };
    static fm33(p0, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(!(false)), !(!(true))), _dafny.Seq.of(true, true)), _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(true, false, true, false))));
    };
    static fm34(globalState) {
      return function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of ((function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_64_i0) {
            return (_module.D22.create_DC70(true, !(false), _dafny.Seq.UnicodeFromString("rgiu"), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(331), new BigNumber((_dafny.Seq.UnicodeFromString("ia")).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-47),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_65_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).length))).dtor_cf114;
          })).Elements) {
            let _66_v0 = _compr_8;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_64_i0) {
              return (_module.D22.create_DC70(true, !(false), _dafny.Seq.UnicodeFromString("rgiu"), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(331), new BigNumber((_dafny.Seq.UnicodeFromString("ia")).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-47),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_65_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).length))).dtor_cf114;
            }), _66_v0)) {
              _coll8.push([_66_v0,new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, !(true)))).length)]);
            }
          }
          return _coll8;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("elx"),new BigNumber(919)))).Keys.Elements) {
          let _67_v1 = _compr_7;
          if (((function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_64_i0) {
              return (_module.D22.create_DC70(true, !(false), _dafny.Seq.UnicodeFromString("rgiu"), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(331), new BigNumber((_dafny.Seq.UnicodeFromString("ia")).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-47),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_65_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).length))).dtor_cf114;
            })).Elements) {
              let _66_v0 = _compr_9;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_64_i0) {
                return (_module.D22.create_DC70(true, !(false), _dafny.Seq.UnicodeFromString("rgiu"), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(331), new BigNumber((_dafny.Seq.UnicodeFromString("ia")).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-47),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_65_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).length))).dtor_cf114;
              }), _66_v0)) {
                _coll9.push([_66_v0,new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, !(true)))).length)]);
              }
            }
            return _coll9;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("elx"),new BigNumber(919)))).contains(_67_v1)) {
            _coll7.add(_67_v1);
          }
        }
        return _coll7;
      }();
    };
    static fm35(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1(false)).dtor_cf1,false)).length),_module.D5.create_DC18(false))).Merge((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(527), new BigNumber(315))) {
          let _68_v0 = _compr_10;
          if (((new BigNumber(527)).isLessThanOrEqualTo(_68_v0)) && ((_68_v0).isLessThan(new BigNumber(315)))) {
            _coll10.push([(_68_v0).plus(new BigNumber(909)),_module.D5.create_DC18(!(false))]);
          }
        }
        return _coll10;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("niflmco")).length),_module.D5.create_DC18(true))));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(510))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length),new BigNumber(940))).length)))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("xcail")).length)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(954), new BigNumber(68))) {
          let _69_v0 = _compr_11;
          if (((new BigNumber(954)).isLessThanOrEqualTo(_69_v0)) && ((_69_v0).isLessThan(new BigNumber(68)))) {
            _coll11.push([(_69_v0).plus(new BigNumber((_dafny.Set.fromElements(false)).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('f'.codePointAt(0)))).length))).length)]);
          }
        }
        return _coll11;
      }()).length))));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1(!((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),true)).length)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("tvidhsf")).length))));
    };
    static fm38(globalState) {
      let _source3 = _module.D5.create_DC17(_dafny.Seq.of(_dafny.Seq.of(true, false, false)), false, _dafny.MultiSet.fromElements(new BigNumber(372), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(142)), function (_70_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length),new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(new BigNumber(839)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(327),false)).length)));
      if (_source3.is_DC17) {
        let _71___mcc_h0 = (_source3).cf30;
        let _72___mcc_h1 = (_source3).cf31;
        let _73___mcc_h2 = (_source3).cf32;
        let _74_cf32 = _73___mcc_h2;
        let _75_cf31 = _72___mcc_h1;
        let _76_cf30 = _71___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_75_cf31), _dafny.Seq.of(_75_cf31));
      } else if (_source3.is_DC18) {
        let _77___mcc_h3 = (_source3).cf33;
        let _78_cf33 = _77___mcc_h3;
        return _dafny.Seq.of(_78_cf33, true);
      } else if (_source3.is_DC16) {
        let _79___mcc_h4 = (_source3).cf29;
        let _80_cf29 = _79___mcc_h4;
        return _dafny.Seq.of(true, true);
      } else {
        let _81___mcc_h5 = (_source3).cf34;
        let _82_cf34 = _81___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.of(false, true, !(false), true, true), _dafny.Seq.of(true));
      }
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return _module.D9.create_DC30(new _dafny.CodePoint('w'.codePointAt(0)), ((true) ? (false) : (true)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sfdadptf")).length),true)).length)).minus(new BigNumber(844)), function () {
  let _coll12 = new _dafny.Set();
  for (const _compr_12 of _dafny.IntegerRange(new BigNumber(626), new BigNumber(349))) {
    let _83_v0 = _compr_12;
    if (((new BigNumber(626)).isLessThanOrEqualTo(_83_v0)) && ((_83_v0).isLessThan(new BigNumber(349)))) {
      _coll12.add(_module.__default.safeModuloInt(_83_v0, new BigNumber(168)));
    }
  }
  return _coll12;
}(), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_module.D7.create_DC23(false, new BigNumber(232), !(false))).dtor_cf43, new BigNumber((_dafny.Seq.UnicodeFromString("ixq")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-233),true)).length))).length),true)).length));
    };
    static fm40(globalState) {
      return _dafny.Seq.of(_module.D2.create_DC7(_dafny.Seq.of(false, true)));
    };
    static fm41(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(!(false), !(true), !(true))).Intersect(_dafny.Set.fromElements(!(true)))).Union((_dafny.Set.fromElements(!(!(!(false))))).Union(_dafny.Set.fromElements(false)));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      let _source4 = _module.D11.create_DC33(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))));
      if (_source4.is_DC34) {
        let _84___mcc_h0 = (_source4).cf58;
        let _85___mcc_h1 = (_source4).cf59;
        let _86___mcc_h2 = (_source4).cf60;
        let _87___mcc_h3 = (_source4).cf61;
        let _88_cf61 = _87___mcc_h3;
        let _89_cf60 = _86___mcc_h2;
        let _90_cf59 = _85___mcc_h1;
        let _91_cf58 = _84___mcc_h0;
        return _module.D8.create_DC26(true, function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of ((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("di")).length))))).Elements) {
    let _92_v0 = _compr_13;
    if (_dafny.Seq.contains((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("di")).length)))), _92_v0)) {
      _coll13.push([(_92_v0).minus(new BigNumber(-650)),_dafny.Seq.UnicodeFromString("pc")]);
    }
  }
  return _coll13;
}());
      } else if (_source4.is_DC35) {
        return _module.D8.create_DC26(false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(66),_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_93_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})));
      } else if (_source4.is_DC33) {
        let _94___mcc_h4 = (_source4).cf57;
        let _95_cf57 = _94___mcc_h4;
        return _module.D8.create_DC26(!(true), function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of _dafny.IntegerRange(new BigNumber(528), new BigNumber(791))) {
    let _96_v1 = _compr_14;
    if (((new BigNumber(528)).isLessThanOrEqualTo(_96_v1)) && ((_96_v1).isLessThan(new BigNumber(791)))) {
      _coll14.push([(_96_v1).minus(new BigNumber(586)),_dafny.Seq.UnicodeFromString("hsqevyqb")]);
    }
  }
  return _coll14;
}());
      } else {
        let _97___mcc_h5 = (_source4).cf62;
        let _98_cf62 = _97___mcc_h5;
        if (!(true)) {
          return _module.D8.create_DC26(true, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-119),_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), function (_99_i1) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})));
        } else {
          return _module.D8.create_DC26(false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-953),_dafny.Seq.UnicodeFromString("m")));
        }
      }
    };
    static fm43(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements((_module.D0.create_DC0(true)).dtor_cf0, !(true), true, true)), _dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.FromArray(_dafny.Seq.of(true, false)), _dafny.MultiSet.fromElements(false), _dafny.MultiSet.FromArray(_dafny.Seq.of(true)), _dafny.MultiSet.fromElements(!(false)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-251)), function (_100_i0) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(false, false));
      })));
    };
    static fm44(p0, p1, p2, globalState) {
      let _source5 = _module.D13.create_DC41(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(689))));
      if (_source5.is_DC42) {
        let _101___mcc_h0 = (_source5).cf68;
        let _102___mcc_h1 = (_source5).cf69;
        let _103_cf69 = _102___mcc_h1;
        let _104_cf68 = _101___mcc_h0;
        return _module.D2.create_DC8((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), ((_105_cf68) => function (_106_i0) {
  return _105_cf68;
})(_104_cf68)))).length)), new BigNumber(-147), new BigNumber(-214));
      } else if (_source5.is_DC41) {
        let _107___mcc_h2 = (_source5).cf67;
        let _108_cf67 = _107___mcc_h2;
        return _module.D2.create_DC8(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(117))).length), new BigNumber(293), new BigNumber(-852));
      } else {
        let _109___mcc_h3 = (_source5).cf70;
        let _110_cf70 = _109___mcc_h3;
        return _module.D2.create_DC8(new BigNumber((_dafny.Set.fromElements(new BigNumber(188))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.of(false, false, true)).length));
      }
    };
    static fm45(p0, p1, p2, globalState) {
      let _source6 = _module.D9.create_DC30(new _dafny.CodePoint('q'.codePointAt(0)), false, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))),_dafny.ZERO)).length), new BigNumber((_dafny.MultiSet.fromElements(false, true, false)).cardinality()))).length), _dafny.Set.fromElements(new BigNumber(335), new BigNumber(218), new BigNumber(-979)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-175),new BigNumber((_dafny.Seq.of(!(false))).length))).length));
      if (_source6.is_DC30) {
        let _111___mcc_h0 = (_source6).cf50;
        let _112___mcc_h1 = (_source6).cf51;
        let _113___mcc_h2 = (_source6).cf52;
        let _114___mcc_h3 = (_source6).cf53;
        let _115___mcc_h4 = (_source6).cf54;
        let _116_cf54 = _115___mcc_h4;
        let _117_cf53 = _114___mcc_h3;
        let _118_cf52 = _113___mcc_h2;
        let _119_cf51 = _112___mcc_h1;
        let _120_cf50 = _111___mcc_h0;
        return _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_119_cf51,_119_cf51));
      } else if (_source6.is_DC29) {
        let _121___mcc_h5 = (_source6).cf49;
        let _122_cf49 = _121___mcc_h5;
        return _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      } else {
        let _123___mcc_h6 = (_source6).cf55;
        let _124_cf55 = _123___mcc_h6;
        return _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      }
    };
    static fm46(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),!(!(true)))).Merge(function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).Elements) {
          let _125_v0 = _compr_15;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))), _125_v0)) {
            _coll15.push([_125_v0,true]);
          }
        }
        return _coll15;
      }());
    };
    static fm47(globalState) {
      return ((_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(false))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, false, false)), _dafny.MultiSet.fromElements(!(false))))).Union(function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, false, true)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(false, !(false)), _dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true, true, true)))).Elements) {
          let _126_v0 = _compr_16;
          if ((_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, false, true)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(false, !(false)), _dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true, true, true)))).contains(_126_v0)) {
            _coll16.add(_126_v0);
          }
        }
        return _coll16;
      }());
    };
    static fm48(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(true))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).length)),(new BigNumber(871)).plus(new BigNumber(758)));
    };
    static fm49(p0, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(_module.D9.create_DC31(_module.D9.create_DC30(new _dafny.CodePoint('h'.codePointAt(0)), false, new BigNumber((function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(386)),_dafny.Map.Empty.slice().updateUnsafe(!(true),true))).Keys.Elements) {
    let _127_v0 = _compr_17;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(386)),_dafny.Map.Empty.slice().updateUnsafe(!(true),true))).contains(_127_v0)) {
      _coll17.push([_127_v0,false]);
    }
  }
  return _coll17;
}()).length), function () {
  let _coll18 = new _dafny.Set();
  for (const _compr_18 of (_dafny.MultiSet.fromElements(new BigNumber(954))).Elements) {
    let _128_v1 = _compr_18;
    if ((_dafny.MultiSet.fromElements(new BigNumber(954))).contains(_128_v1)) {
      _coll18.add(_module.__default.safeDivisionInt(_128_v1, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-253),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(880))).length))).length),_module.D1.create_DC4(new BigNumber((_dafny.Set.fromElements(new BigNumber(-15))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(974), new BigNumber((_dafny.Seq.of(!(false), false)).length), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(106), new BigNumber((_dafny.Seq.UnicodeFromString("xpse")).length))).cardinality())))).length), new BigNumber(162), new BigNumber((_dafny.Seq.of(false, false, !(false))).length), new BigNumber(929))).length)));
    }
  }
  return _coll18;
}(), new BigNumber(277)))), _dafny.Seq.of(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC29(function () {
  let _coll19 = new _dafny.Map();
  for (const _compr_19 of _dafny.IntegerRange(new BigNumber(140), new BigNumber(644))) {
    let _129_v2 = _compr_19;
    if (((new BigNumber(140)).isLessThanOrEqualTo(_129_v2)) && ((_129_v2).isLessThan(new BigNumber(644)))) {
      _coll19.push([(_129_v2).plus(new BigNumber(293)),new BigNumber((_dafny.Seq.UnicodeFromString("phlpopw")).length)]);
    }
  }
  return _coll19;
}())))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-447)), function (_130_i0) {
        return _module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC30(new _dafny.CodePoint('i'.codePointAt(0)), true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true, true, true)).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-825)), function (_131_i1) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length))));
      }))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(_module.D9.create_DC31((_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC30(new _dafny.CodePoint('w'.codePointAt(0)), false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length))).length), _dafny.Set.fromElements(new BigNumber(271), new BigNumber(444)), new BigNumber(168))))).dtor_cf55), _module.D9.create_DC31(_module.D9.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(733),new BigNumber(-491)))))))).Intersect((function () {
        let _coll20 = new _dafny.Set();
        for (const _compr_20 of (_dafny.Seq.of(_dafny.Seq.of(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC29(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of _dafny.IntegerRange(new BigNumber(394), new BigNumber(-187))) {
    let _132_v3 = _compr_21;
    if (((new BigNumber(394)).isLessThanOrEqualTo(_132_v3)) && ((_132_v3).isLessThan(new BigNumber(-187)))) {
      _coll21.push([_module.__default.safeDivisionInt(_132_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_133_i2) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-559), new BigNumber(990))).cardinality())]);
    }
  }
  return _coll21;
}()))), _module.D9.create_DC31(_module.D9.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_134_i3) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})).length))))))).Elements) {
          let _135_v4 = _compr_20;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC29(function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of _dafny.IntegerRange(new BigNumber(394), new BigNumber(-187))) {
    let _132_v3 = _compr_22;
    if (((new BigNumber(394)).isLessThanOrEqualTo(_132_v3)) && ((_132_v3).isLessThan(new BigNumber(-187)))) {
      _coll22.push([_module.__default.safeDivisionInt(_132_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_133_i2) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-559), new BigNumber(990))).cardinality())]);
    }
  }
  return _coll22;
}()))), _module.D9.create_DC31(_module.D9.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_134_i3) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})).length)))))), _135_v4)) {
            _coll20.add(_135_v4);
          }
        }
        return _coll20;
      }()).Difference(_dafny.Set.fromElements(_dafny.Seq.of(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-463),new BigNumber(601)))))))), _module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC31(_module.D9.create_DC29(function () {
  let _coll23 = new _dafny.Map();
  for (const _compr_23 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()),false)).Keys.Elements) {
    let _136_v5 = _compr_23;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()),false)).contains(_136_v5)) {
      _coll23.push([_module.__default.safeDivisionInt(_136_v5, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(126),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fjfiraa")).length))).length))).length))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("to")).length))]);
    }
  }
  return _coll23;
}()))))))));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of _dafny.IntegerRange(new BigNumber(614), new BigNumber(484))) {
          let _137_v0 = _compr_24;
          if (((new BigNumber(614)).isLessThanOrEqualTo(_137_v0)) && ((_137_v0).isLessThan(new BigNumber(484)))) {
            _coll24.push([(_137_v0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-399)), function (_138_i0) {
              return new _dafny.CodePoint('b'.codePointAt(0));
            })).length)),true]);
          }
        }
        return _coll24;
      }()).length)))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length))));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _module.D13.create_DC43(_module.D13.create_DC43(_module.D13.create_DC42(new _dafny.CodePoint('k'.codePointAt(0)), true)));
    };
    static fm52(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-309)), function (_139_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("jlkqqbehd")).length);
      })).length), new BigNumber(589)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(578),_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-859))))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-541),_dafny.Seq.of(new BigNumber(83))));
    };
    static fm53(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-116), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(!(false))).length)), _dafny.Seq.of(new BigNumber(50)))).length), new BigNumber(945))).length)), _dafny.Set.fromElements(new BigNumber((function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of ((_module.D22.create_DC67(_dafny.Seq.UnicodeFromString("hac"))).dtor_cf108).Elements) {
          let _140_v0 = _compr_25;
          if (_dafny.Seq.contains((_module.D22.create_DC67(_dafny.Seq.UnicodeFromString("hac"))).dtor_cf108, _140_v0)) {
            _coll25.push([_140_v0,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(426)), function (_141_i0) {
              return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_142_i1) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              })).length));
            }))).cardinality())]);
          }
        }
        return _coll25;
      }()).length)), _dafny.Set.fromElements(new BigNumber(106), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(427))).cardinality()))).length)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))).cardinality()), new BigNumber(326))), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(887)))));
    };
    static fm54(p0, p1, globalState) {
      if (!(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), function (_143_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("pgi")).length)))) {
        return _dafny.Set.fromElements((_module.D8.create_DC24(_dafny.Set.fromElements((_module.D1.create_DC5(true, false)).dtor_cf7))).dtor_cf45);
      } else {
        return (function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of (_dafny.Set.fromElements(_dafny.Set.fromElements(true, false))).Elements) {
            let _144_v0 = _compr_26;
            if ((_dafny.Set.fromElements(_dafny.Set.fromElements(true, false))).contains(_144_v0)) {
              _coll26.add(_144_v0);
            }
          }
          return _coll26;
        }()).Intersect(_dafny.Set.fromElements(_dafny.Set.fromElements(true, false)));
      }
    };
    static fm55(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), function (_145_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("wvjfiqb")).length));
      }), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-532), new BigNumber(-470)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(328)), function (_146_i1) {
        return new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality());
      })), _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), function (_147_i2) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-994)), function (_148_i3) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("xoxncefj")).length);
      }), _dafny.Seq.of(new BigNumber(809)))));
    };
    static m11(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let _149_v0;
      _149_v0 = _dafny.Seq.UnicodeFromString("plwnbed");
      let _150_v1;
      let _nw0 = Array((_dafny.ONE).toNumber()).fill(false);
      _150_v1 = _nw0;
      let _151_v2;
      _151_v2 = _module.D4.create_DC13(_149_v0, p1, false, _150_v1);
      let _source7 = function (_pat_let0_0) {
        return function (_152_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_153_dt__update_hcf22_h0) {
              return _module.D4.create_DC13(_153_dt__update_hcf22_h0, (_152_dt__update__tmp_h0).dtor_cf23, (_152_dt__update__tmp_h0).dtor_cf24, (_152_dt__update__tmp_h0).dtor_cf25);
            }(_pat_let1_0);
          }(_dafny.Seq.UnicodeFromString("lkgqgh"));
        }(_pat_let0_0);
      }(((false) ? (_151_v2) : (_151_v2)));
      if (_source7.is_DC13) {
        let _154___mcc_h0 = (_source7).cf22;
        let _155___mcc_h1 = (_source7).cf23;
        let _156___mcc_h2 = (_source7).cf24;
        let _157___mcc_h3 = (_source7).cf25;
        let _158_cf25 = _157___mcc_h3;
        let _159_cf24 = _156___mcc_h2;
        let _160_cf23 = _155___mcc_h1;
        let _161_cf22 = _154___mcc_h0;
        _159_cf24 = ((true) ? (p3) : (p2));
        (globalState).f9 = p2;
        let _162_v3;
        _162_v3 = _dafny.MultiSet.fromElements(p3);
        let _163_v4;
        let _nw1 = new _module.C6();
        _nw1.__ctor(_159_cf24, (_162_v3).Union((_162_v3).update(!(p3), _module.__default.abs(_160_cf23))));
        _163_v4 = _nw1;
        _163_v4 = _163_v4;
        let _164_v5;
        let _nw2 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _164_v5 = _nw2;
        let _165_v6;
        _165_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-374),_164_v5);
        let _166_v7;
        _166_v7 = _module.D18.create_DC56(p3, _163_v4.f14, _160_cf23);
        let _167_v8;
        _167_v8 = _dafny.Seq.of(p2);
        let _168_v9;
        let _nw3 = new _module.C0();
        _nw3.__ctor((_167_v8)[_module.__default.safeIndex(_160_cf23, new BigNumber((_167_v8).length))], p2);
        _168_v9 = _nw3;
        let _169_v10;
        _169_v10 = _dafny.Seq.of(_168_v9);
        let _170_v11;
        _170_v11 = _dafny.Map.Empty.slice().updateUnsafe((_166_v7).dtor_cf93,new BigNumber((_169_v10).length));
        let _171_v12;
        _171_v12 = _dafny.Seq.of(_170_v11, _dafny.Map.Empty.slice().updateUnsafe(p3,_160_cf23), (_170_v11).update((_168_v9).f24, p1));
        let _172_v13;
        let _nw4 = new _module.C8();
        _nw4.__ctor(_165_v6, _171_v12, p2, _162_v3);
        _172_v13 = _nw4;
        let _173_v14;
        let _nw5 = Array((new BigNumber(5)).toNumber());
        _nw5[(_dafny.ZERO).toNumber()] = _172_v13;
        _nw5[(_dafny.ONE).toNumber()] = _172_v13;
        _nw5[(new BigNumber(2)).toNumber()] = _172_v13;
        _nw5[(new BigNumber(3)).toNumber()] = _172_v13;
        _nw5[(new BigNumber(4)).toNumber()] = _172_v13;
        _173_v14 = _nw5;
        _173_v14 = ((_159_cf24) ? (_173_v14) : (_173_v14));
      } else if (_source7.is_DC14) {
        let _174___mcc_h4 = (_source7).cf26;
        let _175___mcc_h5 = (_source7).cf27;
        let _176_cf27 = _175___mcc_h5;
        let _177_cf26 = _174___mcc_h4;
        let _178_v15;
        _178_v15 = _dafny.Seq.of(p2);
        let _179_v16;
        let _nw6 = new _module.C6();
        _nw6.__ctor(true, _dafny.MultiSet.FromArray(_dafny.Seq.of(_177_cf26)));
        _179_v16 = _nw6;
        let _180_v17;
        _180_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,_179_v16);
        let _rhs0 = _module.__default.safeDivisionInt(new BigNumber((((p3) ? (_178_v15) : (_178_v15))).length), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),_177_cf26)).length)).minus(new BigNumber(((_180_v17).update((_178_v15)[_module.__default.safeIndex(p1, new BigNumber((_178_v15).length))], _179_v16)).length)));
        let _rhs1 = (_179_v16).fm2(_dafny.Seq.IsPrefixOf(_178_v15, _178_v15), p2, globalState);
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _lhs0.f3 = _rhs0;
        _lhs1.f9 = _rhs1;
        let _181_v18;
        _181_v18 = _dafny.MultiSet.fromElements(_177_cf26, _176_cf27);
        let _182_v19;
        let _nw7 = new _module.C4();
        _nw7.__ctor(_176_cf27, _177_cf26, _181_v18);
        _182_v19 = _nw7;
        let _183_v20;
        _183_v20 = _dafny.Map.Empty.slice().updateUnsafe(_182_v19,p2);
        let _184_v21;
        let _nw8 = new _module.C5();
        _nw8.__ctor((p1).plus(new BigNumber((_183_v20).length)), p2, _dafny.MultiSet.fromElements(_177_cf26, true));
        _184_v21 = _nw8;
        _184_v21 = _184_v21;
        let _185_v22;
        let _nw9 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _185_v22 = _nw9;
        let _index0 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_185_v22).length));
        (_185_v22)[_index0] = ((_184_v21).f18).multipliedBy((_dafny.ZERO).minus(p1));
        let _index1 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_150_v1).length));
        (_150_v1)[_index1] = !(p1).isEqualTo(p1);
        let _186_v23;
        let _nw10 = Array((new BigNumber(9)).toNumber());
        _186_v23 = _nw10;
        let _187_v24;
        _187_v24 = _dafny.Set.fromElements(_186_v23, _186_v23, _186_v23);
        let _188_v25;
        _188_v25 = _dafny.Map.Empty.slice().updateUnsafe((_179_v16).fm2(p2, (_182_v19).f19, globalState),_module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), new BigNumber((_187_v24).length)));
        let _index2 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_185_v22).length));
        let _index3 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_150_v1).length));
        let _rhs2 = _dafny.Seq.Concat(_149_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_189_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
        let _rhs3 = _177_cf26;
        let _rhs4 = new BigNumber((_188_v25).length);
        let _rhs5 = p2;
        let _lhs2 = globalState;
        let _lhs3 = _185_v22;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_185_v22).length));
        let _lhs5 = _150_v1;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_150_v1).length));
        _149_v0 = _rhs2;
        _lhs2.f9 = _rhs3;
        _lhs3[_lhs4] = _rhs4;
        _lhs5[_lhs6] = _rhs5;
        let _index4 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_150_v1).length));
        (_150_v1)[_index4] = (_178_v15)[_module.__default.safeIndex((_185_v22)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_185_v22).length))], new BigNumber((_178_v15).length))];
      } else if (_source7.is_DC12) {
        let _190___mcc_h6 = (_source7).cf21;
        let _191_cf21 = _190___mcc_h6;
        let _index5 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
        (_150_v1)[_index5] = (p1).isLessThan(p1);
        let _192_v26;
        _192_v26 = _dafny.Set.fromElements(p3, p2, false);
        let _193_v27;
        _193_v27 = _dafny.Set.fromElements(_192_v26);
        let _index6 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
        let _rhs6 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p1));
        let _rhs7 = ((_193_v27).Union(_193_v27)).IsSubsetOf(_module.__default.fm54(new BigNumber(850), _149_v0, globalState));
        let _lhs7 = globalState;
        let _lhs8 = _150_v1;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
        _lhs7.f3 = _rhs6;
        _lhs8[_lhs9] = _rhs7;
        if ((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]) {
          (globalState).f3 = p1;
          let _194_v28;
          _194_v28 = _dafny.Seq.of(p1);
          let _195_v29;
          _195_v29 = _dafny.MultiSet.fromElements(p3);
          let _196_v30;
          let _nw11 = new _module.C6();
          _nw11.__ctor((p1).isLessThan(new BigNumber((_194_v28).length)), _195_v29);
          _196_v30 = _nw11;
          (globalState).f3 = (p1).plus(p1);
          let _197_v31;
          let _nw12 = new _module.C3();
          _nw12.__ctor(p1, _149_v0, p3, _195_v29);
          _197_v31 = _nw12;
          let _198_v32;
          _198_v32 = _dafny.Seq.of(_197_v31);
          let _199_v33;
          let _nw13 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _199_v33 = _nw13;
          let _200_v34;
          _200_v34 = _dafny.Set.fromElements(_199_v33, _199_v33, _199_v33);
          let _201_v35;
          _201_v35 = _dafny.Set.fromElements(new BigNumber((_200_v34).length));
          (globalState).f9 = !(_201_v35).contains(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_198_v32)[_module.__default.safeIndex(p1, new BigNumber((_198_v32).length))], _197_v31, _197_v31, _197_v31, _197_v31), _198_v32)).length));
          let _index7 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
          (_150_v1)[_index7] = !(true);
        } else {
          let _index8 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
          (_150_v1)[_index8] = (p1).isLessThanOrEqualTo(p1);
          let _202_v36;
          _202_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]);
          let _203_v37;
          _203_v37 = _module.D1.create_DC5(p3, true);
          let _204_v38;
          _204_v38 = _dafny.Map.Empty.slice().updateUnsafe((((_202_v36).contains(p1)) ? ((_202_v36).get(p1)) : (true)),_203_v37);
          _204_v38 = _204_v38;
          let _205_v39;
          _205_v39 = _dafny.Set.fromElements(p1, p1);
          let _206_v40;
          let _nw14 = Array((new BigNumber(28)).toNumber());
          _206_v40 = _nw14;
          let _207_v41;
          _207_v41 = _dafny.Seq.of(_205_v39, _module.__default.fm6(globalState));
          let _208_v42;
          _208_v42 = _dafny.Seq.of(p2, p3);
          let _209_v43;
          _209_v43 = _dafny.Map.Empty.slice().updateUnsafe((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))],(_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]);
          let _210_v44;
          _210_v44 = new _dafny.CodePoint('c'.codePointAt(0));
          let _211_v45;
          _211_v45 = _dafny.Map.Empty.slice().updateUnsafe(_210_v44,p1);
          let _212_v46;
          _212_v46 = _dafny.Seq.of(p1, p1, p1, p1, p1);
          let _213_v47;
          _213_v47 = _dafny.Set.fromElements(_212_v46, _212_v46, _212_v46, _212_v46);
          let _index9 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
          let _rhs8 = ((_207_v41)[_module.__default.safeIndex(new BigNumber((_208_v42).length), new BigNumber((_207_v41).length))]).Difference(_205_v39);
          let _rhs9 = (((_209_v43).contains(!((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]))) ? ((_209_v43).get(!((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]))) : (_module.__default.fm23(new BigNumber((_202_v36).length), (((_202_v36).contains(p1)) ? ((_202_v36).get(p1)) : (!(p3))), _211_v45, _213_v47, globalState)));
          let _rhs10 = _206_v40;
          let _lhs10 = _150_v1;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
          _205_v39 = _rhs8;
          _lhs10[_lhs11] = _rhs9;
          _206_v40 = _rhs10;
          let _214_v48;
          let _nw15 = Array((new BigNumber(8)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _191_cf21;
          _nw15[(_dafny.ONE).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(2)).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(3)).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(4)).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(5)).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(6)).toNumber()] = _191_cf21;
          _nw15[(new BigNumber(7)).toNumber()] = _191_cf21;
          _214_v48 = _nw15;
          let _index10 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_214_v48).length));
          (_214_v48)[_index10] = _191_cf21;
          let _index11 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_214_v48).length));
          let _nw16 = Array((new BigNumber(29)).toNumber()).fill(false);
          (_214_v48)[_index11] = _nw16;
          let _215_v49;
          _215_v49 = _dafny.MultiSet.fromElements(p3, p2);
          let _216_v50;
          let _nw17 = new _module.C5();
          _nw17.__ctor(p1, p3, (_215_v49).Difference((_215_v49).update(p3, _module.__default.abs(new BigNumber((_215_v49).cardinality())))));
          _216_v50 = _nw17;
        }
        let _217_v51;
        _217_v51 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
        let _218_v52;
        _218_v52 = _dafny.MultiSet.fromElements((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))], (_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))], p2);
        let _219_v53;
        _219_v53 = _dafny.MultiSet.fromElements(new BigNumber(-561), p1);
        let _220_v54;
        _220_v54 = _dafny.Seq.of(p1);
        let _221_v55;
        _221_v55 = _dafny.Seq.of(_220_v54);
        let _222_v56;
        _222_v56 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(622));
        let _223_v57;
        let _nw18 = Array((new BigNumber(23)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(p1, p1);
        _nw18[(_dafny.ONE).toNumber()] = _module.__default.fm0(p2, (_dafny.ZERO).minus(p1), p1, (_dafny.MultiSet.fromElements(new BigNumber((_149_v0).length), new BigNumber(-622), p1, p1, (_dafny.ZERO).minus(p1))).update(p1, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_217_v51).length)))), globalState);
        _nw18[(new BigNumber(2)).toNumber()] = p1;
        _nw18[(new BigNumber(3)).toNumber()] = (p1).plus((_dafny.ZERO).minus(new BigNumber((_218_v52).cardinality())));
        _nw18[(new BigNumber(4)).toNumber()] = p1;
        _nw18[(new BigNumber(5)).toNumber()] = new BigNumber(978);
        _nw18[(new BigNumber(6)).toNumber()] = _module.__default.fm0(false, p1, p1, _219_v53, globalState);
        _nw18[(new BigNumber(7)).toNumber()] = new BigNumber((_221_v55).length);
        _nw18[(new BigNumber(8)).toNumber()] = p1;
        _nw18[(new BigNumber(9)).toNumber()] = ((p2) ? (p1) : (p1));
        _nw18[(new BigNumber(10)).toNumber()] = (((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]) ? ((_220_v54)[_module.__default.safeIndex(p1, new BigNumber((_220_v54).length))]) : (p1));
        _nw18[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
        _nw18[(new BigNumber(12)).toNumber()] = p1;
        _nw18[(new BigNumber(13)).toNumber()] = new BigNumber(929);
        _nw18[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_220_v54).length));
        _nw18[(new BigNumber(15)).toNumber()] = p1;
        _nw18[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_220_v54).length));
        _nw18[(new BigNumber(17)).toNumber()] = p1;
        _nw18[(new BigNumber(18)).toNumber()] = (((_222_v56).contains((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))])) ? ((_222_v56).get((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))])) : (new BigNumber(171)));
        _nw18[(new BigNumber(19)).toNumber()] = p1;
        _nw18[(new BigNumber(20)).toNumber()] = p1;
        _nw18[(new BigNumber(21)).toNumber()] = p1;
        _nw18[(new BigNumber(22)).toNumber()] = p1;
        _223_v57 = _nw18;
        _223_v57 = _223_v57;
        if (((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, p1))).isLessThanOrEqualTo(p1)) {
          (globalState).f9 = !(((true) ? ((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]) : (p2)));
          let _224_v58;
          _224_v58 = _dafny.Map.Empty.slice().updateUnsafe(_191_cf21,(_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]);
          let _225_v59;
          _225_v59 = new _dafny.CodePoint('p'.codePointAt(0));
          let _226_v60;
          _226_v60 = _dafny.Map.Empty.slice().updateUnsafe(_225_v59,new BigNumber((_dafny.Seq.of(p1, new BigNumber((_dafny.Set.fromElements(p1)).length))).length));
          let _227_v61;
          _227_v61 = _dafny.Set.fromElements(_220_v54);
          _224_v58 = _dafny.Map.Empty.slice().updateUnsafe(_191_cf21,_module.__default.fm23(p1, false, _module.__default.fm24(new BigNumber(-773), p3, true, p1, globalState), function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (_module.__default.fm55(_module.__default.fm23(p1, p2, _226_v60, _227_v61, globalState), globalState)).Elements) {
              let _228_v62 = _compr_27;
              if (_dafny.Seq.contains(_module.__default.fm55(_module.__default.fm23(p1, p2, _226_v60, _227_v61, globalState), globalState), _228_v62)) {
                _coll27.add(_228_v62);
              }
            }
            return _coll27;
          }(), globalState));
          let _229_v64;
          _229_v64 = _dafny.Map.Empty.slice().updateUnsafe(_220_v54,p2);
          let _230_v66;
          let _nw19 = Array((new BigNumber(15)).toNumber());
          _230_v66 = _nw19;
          let _231_v67;
          let _nw20 = new _module.C3();
          _nw20.__ctor(_module.__default.fm0((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))], p1, p1, _219_v53, globalState), _dafny.Seq.Concat(_dafny.Seq.update(_149_v0, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_module.__default.fm23(p1, false, function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_149_v0).Elements) {
              let _232_v63 = _compr_28;
              if (_dafny.Seq.contains(_149_v0, _232_v63)) {
                _coll28.push([_232_v63,p1]);
              }
            }
            return _coll28;
          }(), function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of (_229_v64).Keys.Elements) {
              let _233_v65 = _compr_29;
              if ((_229_v64).contains(_233_v65)) {
                _coll29.add(_233_v65);
              }
            }
            return _coll29;
          }(), globalState), p1, p1, _219_v53, globalState),_230_v66)).length), new BigNumber((_149_v0).length)), _225_v59), _dafny.Seq.update(_149_v0, _module.__default.safeIndex(p1, new BigNumber((_149_v0).length)), _225_v59)), (_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))], _218_v52);
          _231_v67 = _nw20;
          let _nw21 = new _module.C3();
          _nw21.__ctor((_231_v67).f20, (_231_v67).f21, p3, _218_v52);
          _231_v67 = _nw21;
          let _234_v68;
          _234_v68 = _dafny.Set.fromElements((_231_v67).f20, (_231_v67).f20);
          let _235_v69;
          let _236_v70;
          let _237_v71;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_231_v67).m1(_234_v68, (_231_v67).f20, globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _235_v69 = _out0;
          _236_v70 = _out1;
          _237_v71 = _out2;
          let _238_v72;
          _238_v72 = _module.D5.create_DC18(false);
          let _239_v73;
          _239_v73 = _dafny.Map.Empty.slice().updateUnsafe(p3,_238_v72);
          _239_v73 = (_239_v73).update(p3, _module.D5.create_DC18((_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))]));
        } else {
          let _240_v75;
          _240_v75 = new _dafny.CodePoint('f'.codePointAt(0));
          let _241_v77;
          _241_v77 = _dafny.Seq.of(p2);
          let _242_v78;
          _242_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_149_v0).length),new BigNumber((function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_dafny.MultiSet.fromElements(_241_v77)).Elements) {
              let _243_v76 = _compr_30;
              if ((_dafny.MultiSet.fromElements(_241_v77)).contains(_243_v76)) {
                _coll30.push([_243_v76,_dafny.Seq.of(p1)]);
              }
            }
            return _coll30;
          }()).length));
          let _244_v79;
          _244_v79 = _dafny.Map.Empty.slice().updateUnsafe(_240_v75,(_242_v78).update(new BigNumber((_241_v77).length), new BigNumber(165)));
          let _245_v80;
          _245_v80 = _dafny.Set.fromElements(p1, p1);
          let _246_v81;
          _246_v81 = _module.D9.create_DC30(_module.__default.fm25(p1, p1, p3, _module.__default.fm1(p2, p1, true, _245_v80, globalState), globalState), (_150_v1)[_module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length))], p1, _245_v80, new BigNumber(961));
          let _247_v82;
          _247_v82 = _dafny.Map.Empty.slice().updateUnsafe((_246_v81).dtor_cf50,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), ((_248_v75) => function (_249_i1) {
            return _248_v75;
          })(_240_v75))).length)));
          let _250_v84;
          _250_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(101)),p1);
          let _index12 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_150_v1).length));
          (_150_v1)[_index12] = _module.__default.fm23(p1, false, (function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_244_v79).Keys.Elements) {
              let _251_v74 = _compr_31;
              if ((_244_v79).contains(_251_v74)) {
                _coll31.push([_251_v74,new BigNumber(781)]);
              }
            }
            return _coll31;
          }()).Merge(_247_v82), function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (function () {
              let _coll33 = new _dafny.Map();
              for (const _compr_33 of (_250_v84).Keys.Elements) {
                let _252_v83 = _compr_33;
                if ((_250_v84).contains(_252_v83)) {
                  _coll33.push([_252_v83,new BigNumber(540)]);
                }
              }
              return _coll33;
            }()).Keys.Elements) {
              let _253_v85 = _compr_32;
              if ((function () {
                let _coll34 = new _dafny.Map();
                for (const _compr_34 of (_250_v84).Keys.Elements) {
                  let _252_v83 = _compr_34;
                  if ((_250_v84).contains(_252_v83)) {
                    _coll34.push([_252_v83,new BigNumber(540)]);
                  }
                }
                return _coll34;
              }()).contains(_253_v85)) {
                _coll32.add(_253_v85);
              }
            }
            return _coll32;
          }(), globalState);
          let _254_v86;
          let _nw22 = Array((_dafny.ONE).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _149_v0;
          _254_v86 = _nw22;
          let _index13 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length));
          (_254_v86)[_index13] = _dafny.Seq.Concat(_149_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(869)), ((_255_v75) => function (_256_i2) {
            return _255_v75;
          })(_240_v75)));
          let _index14 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length));
          (_254_v86)[_index14] = _149_v0;
          let _rhs11 = new BigNumber(667);
          let _rhs12 = _191_cf21;
          let _rhs13 = (p1).minus(new BigNumber((_dafny.Seq.Concat((_254_v86)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length))], _dafny.Seq.update((_254_v86)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length))], _module.__default.safeIndex(p1, new BigNumber(((_254_v86)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length))]).length)), ((_254_v86)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length))])[_module.__default.safeIndex(new BigNumber((_222_v56).length), new BigNumber(((_254_v86)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_254_v86).length))]).length))]))).length));
          let _lhs12 = globalState;
          _lhs12.f3 = _rhs11;
          _191_cf21 = _rhs12;
          r0 = _rhs13;
          _245_v80 = ((_245_v80).Intersect(_245_v80)).Union((_dafny.Set.fromElements(p1)).Difference(_245_v80));
          let _257_v87;
          _257_v87 = _module.D1.create_DC3(p1);
          let _index15 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_223_v57).length));
          (_223_v57)[_index15] = (_257_v87).dtor_cf3;
          let _index16 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_223_v57).length));
          (_223_v57)[_index16] = (_module.__default.safeDivisionInt(p1, new BigNumber(474))).minus(p1);
        }
      } else {
        let _258___mcc_h7 = (_source7).cf28;
        let _259_cf28 = _258___mcc_h7;
        let _260_v88;
        _260_v88 = new _dafny.CodePoint('i'.codePointAt(0));
        let _261_v89;
        _261_v89 = _dafny.Map.Empty.slice().updateUnsafe(_260_v88,p1);
        let _262_v90;
        _262_v90 = _dafny.MultiSet.fromElements(p3, p3);
        let _263_v91;
        _263_v91 = _dafny.Seq.of(new BigNumber((_262_v90).cardinality()), new BigNumber(555));
        let _264_v92;
        _264_v92 = _dafny.Set.fromElements(_263_v91, _dafny.Seq.of(p1));
        let _265_v93;
        _265_v93 = _dafny.Seq.of(_module.__default.fm23(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-800)), p1)).length), p2, _261_v89, _264_v92, globalState));
        let _266_v94;
        let _nw23 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _266_v94 = _nw23;
        let _267_v95;
        _267_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_265_v93)).cardinality()),_266_v94);
        let _268_v96;
        _268_v96 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_dafny.ZERO).minus(new BigNumber((_149_v0).length)));
        let _269_v97;
        _269_v97 = _dafny.Seq.of(_268_v96);
        let _270_v98;
        let _nw24 = new _module.C8();
        _nw24.__ctor(_267_v95, _269_v97, (_265_v93)[_module.__default.safeIndex(new BigNumber((_265_v93).length), new BigNumber((_265_v93).length))], (_262_v90).update(p3, _module.__default.abs((_dafny.ZERO).minus(p1))));
        _270_v98 = _nw24;
        let _271_v99;
        _271_v99 = _module.D2.create_DC9(_270_v98.f14, new BigNumber(19), (_149_v0)[_module.__default.safeIndex(p1, new BigNumber((_149_v0).length))]);
        let _nw25 = new _module.C3();
        _nw25.__ctor((_271_v99).dtor_cf14, _module.__default.fm27(p1, _270_v98.f14, globalState), _270_v98.f14, ((_270_v98).f15).Intersect(_dafny.MultiSet.FromArray(_265_v93)));
        _270_v98 = _nw25;
        (_270_v98).f14 = false;
        let _272_v100;
        let _nw26 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _272_v100 = _nw26;
        let _273_v101;
        _273_v101 = _dafny.Map.Empty.slice().updateUnsafe(_270_v98.f14,p3);
        let _index17 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_272_v100).length));
        (_272_v100)[_index17] = _273_v101;
        let _274_v102;
        _274_v102 = _dafny.MultiSet.fromElements(new BigNumber((_149_v0).length));
        let _275_v103;
        let _nw27 = Array((new BigNumber(3)).toNumber());
        _nw27[(_dafny.ZERO).toNumber()] = new BigNumber(-46);
        _nw27[(_dafny.ONE).toNumber()] = _module.__default.fm0(_270_v98.f14, p1, p1, _274_v102, globalState);
        _nw27[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p1);
        _275_v103 = _nw27;
        let _index18 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
        (_275_v103)[_index18] = p1;
        let _index19 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_150_v1).length));
        (_150_v1)[_index19] = _270_v98.f14;
        let _index20 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_266_v94).length));
        (_266_v94)[_index20] = _260_v88;
        let _276_v104;
        _276_v104 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,_270_v98.f14), _273_v101, _dafny.Map.Empty.slice().updateUnsafe(true,p2));
        let _277_v105;
        let _nw28 = new _module.C8();
        _nw28.__ctor(_dafny.Map.Empty.slice().updateUnsafe(p1,_266_v94), _269_v97, _270_v98.f14, (_dafny.MultiSet.fromElements(p2)).update(_270_v98.f14, _module.__default.abs(p1)));
        _277_v105 = _nw28;
        let _278_v106;
        _278_v106 = _dafny.MultiSet.fromElements(_277_v105);
        let _279_v107;
        _279_v107 = _dafny.Map.Empty.slice().updateUnsafe(p1,_270_v98.f14);
        let _index21 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_272_v100).length));
        let _index22 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
        let _index23 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_150_v1).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_266_v94).length));
        let _rhs14 = (_276_v104)[_module.__default.safeIndex(p1, new BigNumber((_276_v104).length))];
        let _rhs15 = (_dafny.ZERO).minus(((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_278_v106).cardinality())))).cardinality())).multipliedBy(p1))).plus(p1));
        let _rhs16 = p1;
        let _rhs17 = (((_279_v107).contains(new BigNumber((_module.__default.fm27(new BigNumber(467), _270_v98.f14, globalState)).length))) ? ((_279_v107).get(new BigNumber((_module.__default.fm27(new BigNumber(467), _270_v98.f14, globalState)).length))) : (p2));
        let _rhs18 = _260_v88;
        let _lhs13 = _272_v100;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_272_v100).length));
        let _lhs15 = _275_v103;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
        let _lhs17 = _150_v1;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_150_v1).length));
        let _lhs19 = _266_v94;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_266_v94).length));
        _lhs13[_lhs14] = _rhs14;
        r0 = _rhs15;
        _lhs15[_lhs16] = _rhs16;
        _lhs17[_lhs18] = _rhs17;
        _lhs19[_lhs20] = _rhs18;
        if (!((_150_v1)[_module.__default.safeIndex(new BigNumber(672), new BigNumber((_150_v1).length))])) {
          let _280_v108;
          let _nw29 = new _module.C2();
          _nw29.__ctor(_279_v107);
          _280_v108 = _nw29;
          let _index25 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
          (_275_v103)[_index25] = (_263_v91)[_module.__default.safeIndex((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))], new BigNumber((_263_v91).length))];
          let _281_v109;
          let _nw30 = Array((new BigNumber(12)).toNumber()).fill([]);
          _281_v109 = _nw30;
          let _282_v110;
          _282_v110 = _dafny.Map.Empty.slice().updateUnsafe((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))],_275_v103);
          let _283_v111;
          _283_v111 = _dafny.Map.Empty.slice().updateUnsafe(p2,_275_v103);
          let _284_v112;
          let _nw31 = Array((new BigNumber(25)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _275_v103;
          _nw31[(_dafny.ONE).toNumber()] = _275_v103;
          _nw31[(new BigNumber(2)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(3)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(4)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(5)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(6)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(7)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(8)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(9)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(10)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(11)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(12)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(13)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(14)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(15)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(16)).toNumber()] = (((_282_v110).contains(p1)) ? ((_282_v110).get(p1)) : (_275_v103));
          _nw31[(new BigNumber(17)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(18)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(19)).toNumber()] = (((_283_v111).contains(_270_v98.f14)) ? ((_283_v111).get(_270_v98.f14)) : (_275_v103));
          _nw31[(new BigNumber(20)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(21)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(22)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(23)).toNumber()] = _275_v103;
          _nw31[(new BigNumber(24)).toNumber()] = _275_v103;
          _284_v112 = _nw31;
          let _index26 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_281_v109).length));
          (_281_v109)[_index26] = _284_v112;
          let _index27 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_281_v109).length));
          let _rhs19 = _284_v112;
          let _rhs20 = (_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))];
          let _rhs21 = _module.__default.safeDivisionInt(p1, (_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))]);
          let _lhs21 = _281_v109;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_281_v109).length));
          let _lhs23 = globalState;
          let _lhs24 = globalState;
          _lhs21[_lhs22] = _rhs19;
          _lhs23.f3 = _rhs20;
          _lhs24.f3 = _rhs21;
          let _285_v113;
          let _nw32 = new _module.C1();
          _nw32.__ctor();
          _285_v113 = _nw32;
          let _286_v114;
          _286_v114 = _module.D0.create_DC1(p3);
          _286_v114 = _286_v114;
        } else {
          let _287_v115;
          let _nw33 = Array((new BigNumber(12)).toNumber()).fill(false);
          _287_v115 = _nw33;
          let _288_v116;
          _288_v116 = _dafny.Map.Empty.slice().updateUnsafe(p1,_287_v115);
          (globalState).f2 = (((_288_v116).contains(((_270_v98.f14) ? ((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))]) : (new BigNumber(-369))))) ? ((_288_v116).get(((_270_v98.f14) ? ((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))]) : (new BigNumber(-369))))) : (_287_v115));
          let _289_v117;
          _289_v117 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))],_270_v98.f14),_270_v98.f14);
          let _290_v118;
          _290_v118 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_289_v117).length),_149_v0);
          let _291_v119;
          _291_v119 = _module.D8.create_DC26(_module.__default.fm23(p1, _270_v98.f14, _261_v89, _264_v92, globalState), _290_v118);
          let _292_v120;
          _292_v120 = _module.D1.create_DC3(new BigNumber(808));
          let _293_v121;
          _293_v121 = _module.D1.create_DC6(_292_v120);
          let _294_v122;
          _294_v122 = _module.D1.create_DC6(_292_v120);
          let _295_v123;
          _295_v123 = _dafny.Map.Empty.slice().updateUnsafe(_291_v119,_293_v121);
          let _296_v124;
          _296_v124 = _module.D1.create_DC6((((_295_v123).contains(_291_v119)) ? ((_295_v123).get(_291_v119)) : (_292_v120)));
          let _297_v125;
          let _nw34 = Array((_dafny.ONE).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = _296_v124;
          _297_v125 = _nw34;
          let _index28 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_297_v125).length));
          (_297_v125)[_index28] = _296_v124;
          let _index29 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_297_v125).length));
          (_297_v125)[_index29] = _296_v124;
          let _index30 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
          let _rhs22 = _module.__default.safeModuloInt(new BigNumber(667), _module.__default.fm0(_270_v98.f14, new BigNumber(173), p1, _274_v102, globalState));
          let _rhs23 = p3;
          let _lhs25 = _275_v103;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length));
          let _lhs27 = _270_v98;
          _lhs25[_lhs26] = _rhs22;
          _lhs27.f14 = _rhs23;
          (globalState).f2 = _287_v115;
          let _298_v126;
          _298_v126 = _dafny.Set.fromElements(new BigNumber((_268_v96).length));
          let _rhs24 = _270_v98;
          let _rhs25 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber((_279_v107).length)).plus((_275_v103)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_275_v103).length))]), p1));
          let _rhs26 = (_266_v94)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_266_v94).length))];
          let _rhs27 = (_298_v126).Union((_module.__default.fm6(globalState)).Union(_298_v126));
          let _lhs28 = globalState;
          _270_v98 = _rhs24;
          _lhs28.f3 = _rhs25;
          _260_v88 = _rhs26;
          _298_v126 = _rhs27;
        }
      }
      let _299_v127;
      let _nw35 = new _module.C7();
      _nw35.__ctor();
      _299_v127 = _nw35;
      _299_v127 = _299_v127;
      r0 = (_dafny.ZERO).minus(p1);
      _149_v0 = (_151_v2).dtor_cf22;
      let _300_v128;
      _300_v128 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
      let _301_v129;
      let _nw36 = new _module.C4();
      _nw36.__ctor(p3, p3, _dafny.MultiSet.fromElements(p2));
      _301_v129 = _nw36;
      let _302_v130;
      _302_v130 = _module.D20.create_DC60(_301_v129);
      let _303_v131;
      _303_v131 = _dafny.MultiSet.fromElements(p1, new BigNumber((_module.__default.fm7(globalState)).length), p1, new BigNumber((_300_v128).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_302_v130).dtor_cf99,p3)).length));
      let _304_v132;
      _304_v132 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(p1))),(_299_v127).fm10(_300_v128, _303_v131, (_301_v129).f19, _149_v0, globalState));
      if ((((_304_v132).contains(_module.__default.safeDivisionInt(p1, p1))) ? ((_304_v132).get(_module.__default.safeDivisionInt(p1, p1))) : (p3))) {
        (globalState).f3 = (p1).plus(p1);
        let _305_v133;
        _305_v133 = _dafny.MultiSet.fromElements(p2, !(p2), (_301_v129).f19);
        let _306_v134;
        let _nw37 = new _module.C5();
        _nw37.__ctor(new BigNumber(709), true, _305_v133);
        _306_v134 = _nw37;
        _306_v134 = (((_301_v129).f19) ? (_306_v134) : (_306_v134));
        let _307_v135;
        let _nw38 = new _module.C0();
        _nw38.__ctor((_301_v129).f19, p2);
        _307_v135 = _nw38;
        (globalState).f9 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("wbhurb"), (_151_v2).dtor_cf22)) || ((_307_v135).f23);
        let _308_v136;
        _308_v136 = _module.D1.create_DC5((_301_v129).f19, (_307_v135).f23);
        let _309_v137;
        _309_v137 = _module.D1.create_DC6(_module.D1.create_DC6(_308_v136));
        let _pat_let_tv0 = _308_v136;
        let _310_v138;
        _310_v138 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let2_0) {
          return function (_311_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_312_dt__update_hcf8_h0) {
                return _module.D1.create_DC6(_312_dt__update_hcf8_h0);
              }(_pat_let3_0);
            }(_module.D1.create_DC6(_pat_let_tv0));
          }(_pat_let2_0);
        }(_309_v137),p3);
        _310_v138 = (_310_v138).update(_309_v137, p3);
      } else {
        let _313_v139;
        let _init0 = ((_314_p1) => function (_315_i3) {
          return _module.__default.safeDivisionInt(_315_i3, _314_p1);
        })(p1);
        let _nw39 = Array((new BigNumber(4)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw39.length); _i0_0++) {
          _nw39[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _313_v139 = _nw39;
        let _index31 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_313_v139).length));
        (_313_v139)[_index31] = p1;
        let _316_v140;
        _316_v140 = _dafny.Seq.of((new BigNumber(421)).plus(p1));
        let _index32 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_313_v139).length));
        (_313_v139)[_index32] = new BigNumber((_316_v140).length);
        let _317_v141;
        _317_v141 = _dafny.Set.fromElements((_299_v127).fm10(_300_v128, _303_v131, (_301_v129).f19, _149_v0, globalState), p2);
        (globalState).f3 = (_module.__default.safeModuloInt(p1, p1)).minus(_module.__default.safeModuloInt((_313_v139)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_313_v139).length))], (_dafny.ZERO).minus(new BigNumber((_317_v141).length))));
        (globalState).f9 = ((((_301_v129).f19) ? ((_301_v129).f19) : (true))) || ((_301_v129).f19);
        let _318_v142;
        _318_v142 = _module.D0.create_DC0(p2);
        let _319_v143;
        _319_v143 = _dafny.MultiSet.fromElements((_318_v142).dtor_cf0, (_301_v129).f19);
        let _320_v144;
        let _nw40 = new _module.C3();
        _nw40.__ctor((_313_v139)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_313_v139).length))], _149_v0, (_301_v129).f19, _319_v143);
        _320_v144 = _nw40;
        let _321_v145;
        let _nw41 = new _module.C6();
        _nw41.__ctor(p3, (_320_v144).f15);
        _321_v145 = _nw41;
        let _322_v146;
        _322_v146 = _module.D21.create_DC63(_321_v145);
        let _323_v147;
        _323_v147 = _dafny.Map.Empty.slice().updateUnsafe(_320_v144,(_322_v146).dtor_cf101);
        _323_v147 = (_323_v147).update(_320_v144, _321_v145);
        let _324_v148;
        _324_v148 = new _dafny.CodePoint('o'.codePointAt(0));
        let _325_v149;
        _325_v149 = _dafny.Set.fromElements((_320_v144).f15);
        let _rhs28 = (_325_v149).contains((_320_v144).f15);
        let _rhs29 = _324_v148;
        let _lhs29 = globalState;
        _lhs29.f9 = _rhs28;
        _324_v148 = _rhs29;
      }
      let _326_v150;
      _326_v150 = _module.D5.create_DC18((_299_v127).fm10(_300_v128, _303_v131, p2, _149_v0, globalState));
      let _pat_let_tv1 = p3;
      let _pat_let_tv2 = p3;
      let _pat_let_tv3 = p1;
      let _pat_let_tv4 = p1;
      let _pat_let_tv5 = p1;
      let _pat_let_tv6 = p2;
      let _pat_let_tv7 = _301_v129;
      let _pat_let_tv8 = _301_v129;
      let _pat_let_tv9 = _301_v129;
      let _pat_let_tv10 = p2;
      (globalState).f9 = function (_source8) {
        if (_source8.is_DC17) {
          let _327___mcc_h8 = (_source8).cf30;
          let _328___mcc_h9 = (_source8).cf31;
          let _329___mcc_h10 = (_source8).cf32;
          let _330_cf32 = _329___mcc_h10;
          let _331_cf31 = _328___mcc_h9;
          let _332_cf30 = _327___mcc_h8;
          return _pat_let_tv1;
        } else if (_source8.is_DC18) {
          let _333___mcc_h11 = (_source8).cf33;
          let _334_cf33 = _333___mcc_h11;
          return _pat_let_tv2;
        } else if (_source8.is_DC16) {
          let _335___mcc_h12 = (_source8).cf29;
          let _336_cf29 = _335___mcc_h12;
          return (new BigNumber(459)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv3,(_module.D1.create_DC4(_pat_let_tv4, _pat_let_tv5)).dtor_cf4)).length)));
        } else {
          let _337___mcc_h13 = (_source8).cf34;
          let _338_cf34 = _337___mcc_h13;
          return !_dafny.areEqual(_dafny.Seq.of(_pat_let_tv6), _dafny.Seq.of((_pat_let_tv7).f19, (_pat_let_tv8).f19, (_pat_let_tv9).f19, _pat_let_tv10));
        }
      }(_326_v150);
      let _339_v151;
      _339_v151 = _dafny.Seq.of((_301_v129).f19);
      r0 = (_module.__default.safeModuloInt(new BigNumber((_339_v151).length), p1)).multipliedBy(p1);
      let _340_v152;
      let _nw42 = Array((new BigNumber(13)).toNumber());
      _340_v152 = _nw42;
      let _341_v153;
      _341_v153 = _dafny.Map.Empty.slice().updateUnsafe(p1,_340_v152);
      r1 = _341_v153;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _342_v0;
      _342_v0 = new BigNumber(206);
      let _343_v1;
      _343_v1 = _dafny.MultiSet.fromElements(_342_v0);
      let _344_v2;
      let _nw43 = Array((new BigNumber(3)).toNumber()).fill(false);
      _344_v2 = _nw43;
      let _345_v3;
      _345_v3 = _dafny.Seq.UnicodeFromString("bwiunjo");
      let _346_v4;
      _346_v4 = true;
      let _347_v5;
      _347_v5 = _dafny.Set.fromElements(_346_v4, _346_v4);
      let _348_v6;
      let _nw44 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
      _348_v6 = _nw44;
      let _349_v7;
      let _init1 = ((_350_v3) => function (_351_i0) {
        return _350_v3;
      })(_345_v3);
      let _nw45 = Array((new BigNumber(11)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw45.length); _i0_1++) {
        _nw45[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _349_v7 = _nw45;
      let _352_globalState;
      let _nw46 = new _module.GlobalState();
      _nw46.__ctor(new BigNumber(946), (_343_v1).Difference(_343_v1), _344_v2, new BigNumber(664), true, _345_v3, _347_v5, false, new BigNumber(762), false, new BigNumber(-427), _348_v6, _349_v7, _dafny.Set.fromElements(_346_v4, _346_v4, false));
      _352_globalState = _nw46;
      let _353_v8;
      _353_v8 = _dafny.Seq.of(_346_v4);
      _353_v8 = _dafny.Seq.Concat(_353_v8, _353_v8);
      let _354_v9;
      _354_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(_342_v0).isEqualTo(_342_v0),_347_v5);
      _354_v9 = (_354_v9).update(_346_v4, _347_v5);
      let _355_v10;
      _355_v10 = _dafny.Seq.of(_342_v0, _342_v0);
      let _hi0 = _342_v0;
      for (let _356_i1 = _module.__default.fm0(_346_v4, _342_v0, _342_v0, (_dafny.MultiSet.FromArray(_355_v10)).update(_342_v0, _module.__default.abs(_342_v0)), _352_globalState); _356_i1.isLessThan(_hi0); _356_i1 = _356_i1.plus(_dafny.ONE)) {
        let _357_v11;
        _357_v11 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_349_v7);
        _349_v7 = (((_357_v11).contains((_342_v0).multipliedBy(new BigNumber(-478)))) ? ((_357_v11).get((_342_v0).multipliedBy(new BigNumber(-478)))) : (_349_v7));
        let _358_v12;
        _358_v12 = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,_module.__default.safeModuloInt(_356_i1, _356_i1));
        let _359_v13;
        let _nw47 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _359_v13 = _nw47;
        let _360_v14;
        _360_v14 = _dafny.MultiSet.fromElements(_359_v13);
        let _361_v15;
        _361_v15 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_359_v13);
        let _362_v16;
        _362_v16 = _dafny.Set.fromElements(new BigNumber((_355_v10).length), _342_v0, _342_v0);
        let _rhs30 = _module.__default.safeModuloInt((_356_i1).minus(_356_i1), _module.__default.safeDivisionInt(_342_v0, _342_v0));
        let _rhs31 = new BigNumber((_dafny.Seq.Concat(_345_v3, _dafny.Seq.Concat(_345_v3, _345_v3))).length);
        let _rhs32 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((((_346_v4) ? (_dafny.MultiSet.fromElements(_359_v13)) : (_360_v14))).update((((_361_v15).contains(_356_i1)) ? ((_361_v15).get(_356_i1)) : (_359_v13)), _module.__default.abs((_dafny.ZERO).minus(_342_v0)))).cardinality())));
        let _rhs33 = (_358_v12).Merge(_module.__default.fm1(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_356_i1,_346_v4)).length), _346_v4, _362_v16, _352_globalState));
        let _lhs30 = _352_globalState;
        let _lhs31 = _352_globalState;
        _342_v0 = _rhs30;
        _lhs30.f3 = _rhs31;
        _lhs31.f3 = _rhs32;
        _358_v12 = _rhs33;
        let _363_v17;
        _363_v17 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_346_v4);
        if ((((_363_v17).contains(_342_v0)) ? ((_363_v17).get(_342_v0)) : (true))) {
          let _index33 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_359_v13).length));
          (_359_v13)[_index33] = _342_v0;
          let _index34 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_359_v13).length));
          (_359_v13)[_index34] = new BigNumber(971);
          let _rhs34 = _345_v3;
          let _rhs35 = _346_v4;
          let _lhs32 = _352_globalState;
          _345_v3 = _rhs34;
          _lhs32.f9 = _rhs35;
          (_352_globalState).f3 = (_dafny.ZERO).minus(new BigNumber(-824));
          _346_v4 = _346_v4;
          let _364_v18;
          let _nw48 = new _module.C1();
          _nw48.__ctor();
          _364_v18 = _nw48;
        } else {
          let _365_v19;
          _365_v19 = new _dafny.CodePoint('r'.codePointAt(0));
          let _366_v20;
          _366_v20 = _dafny.Map.Empty.slice().updateUnsafe(_365_v19,_342_v0);
          _366_v20 = (_366_v20).update(_365_v19, _342_v0);
          let _367_v21;
          _367_v21 = _module.D5.create_DC16(_359_v13);
          let _pat_let_tv11 = _359_v13;
          let _pat_let_tv12 = _359_v13;
          let _pat_let_tv13 = _359_v13;
          let _368_v22;
          let _nw49 = Array((new BigNumber(15)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _367_v21;
          _nw49[(_dafny.ONE).toNumber()] = _367_v21;
          _nw49[(new BigNumber(2)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(3)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(4)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(5)).toNumber()] = function (_pat_let4_0) {
            return function (_369_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_370_dt__update_hcf29_h0) {
                  return _module.D5.create_DC16(_370_dt__update_hcf29_h0);
                }(_pat_let5_0);
              }(_pat_let_tv11);
            }(_pat_let4_0);
          }(_367_v21);
          _nw49[(new BigNumber(6)).toNumber()] = function (_pat_let6_0) {
            return function (_373_dt__update__tmp_h2) {
              return function (_pat_let9_0) {
                return function (_374_dt__update_hcf29_h2) {
                  return _module.D5.create_DC16(_374_dt__update_hcf29_h2);
                }(_pat_let9_0);
              }(_pat_let_tv13);
            }(_pat_let6_0);
          }(function (_pat_let7_0) {
            return function (_371_dt__update__tmp_h1) {
              return function (_pat_let8_0) {
                return function (_372_dt__update_hcf29_h1) {
                  return _module.D5.create_DC16(_372_dt__update_hcf29_h1);
                }(_pat_let8_0);
              }(_pat_let_tv12);
            }(_pat_let7_0);
          }(_367_v21));
          _nw49[(new BigNumber(7)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(8)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(9)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(10)).toNumber()] = _module.D5.create_DC16(_359_v13);
          _nw49[(new BigNumber(11)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(12)).toNumber()] = _module.D5.create_DC16(_359_v13);
          _nw49[(new BigNumber(13)).toNumber()] = _367_v21;
          _nw49[(new BigNumber(14)).toNumber()] = _367_v21;
          _368_v22 = _nw49;
          let _index35 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_368_v22).length));
          (_368_v22)[_index35] = _367_v21;
          let _375_v23;
          _375_v23 = _dafny.MultiSet.fromElements(((_346_v4) ? (_346_v4) : (false)), !_dafny.Seq.contains(_dafny.Seq.of(_356_i1), _342_v0), _346_v4, (_342_v0).isLessThanOrEqualTo(_356_i1), _346_v4);
          let _index36 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_344_v2).length));
          (_344_v2)[_index36] = _346_v4;
          let _index37 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_368_v22).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_344_v2).length));
          let _rhs36 = (_dafny.ZERO).minus(_342_v0);
          let _rhs37 = _367_v21;
          let _rhs38 = (((_346_v4) ? (_375_v23) : (_375_v23))).Intersect(_375_v23);
          let _rhs39 = _346_v4;
          let _lhs33 = _352_globalState;
          let _lhs34 = _368_v22;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_368_v22).length));
          let _lhs36 = _344_v2;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_344_v2).length));
          _lhs33.f3 = _rhs36;
          _lhs34[_lhs35] = _rhs37;
          _375_v23 = _rhs38;
          _lhs36[_lhs37] = _rhs39;
          (_352_globalState).f3 = ((_346_v4) ? (_module.__default.safeModuloInt(_342_v0, _356_i1)) : (_356_i1));
          let _376_v24;
          _376_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_356_i1),_342_v0);
          _376_v24 = (_376_v24).update((_342_v0).plus(_356_i1), _342_v0);
          _346_v4 = _346_v4;
        }
        let _377_v25;
        let _nw50 = new _module.C5();
        _nw50.__ctor(new BigNumber((_353_v8).length), _346_v4, _dafny.MultiSet.fromElements(false, _346_v4));
        _377_v25 = _nw50;
        let _378_v26;
        _378_v26 = _dafny.Seq.of(_377_v25, _377_v25);
        let _379_v27;
        let _nw51 = Array((new BigNumber(24)).toNumber());
        _nw51[(_dafny.ZERO).toNumber()] = _377_v25;
        _nw51[(_dafny.ONE).toNumber()] = _377_v25;
        _nw51[(new BigNumber(2)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(3)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(4)).toNumber()] = (_378_v26)[_module.__default.safeIndex(_356_i1, new BigNumber((_378_v26).length))];
        _nw51[(new BigNumber(5)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(6)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(7)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(8)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(9)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(10)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(11)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(12)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(13)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(14)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(15)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(16)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(17)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(18)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(19)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(20)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(21)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(22)).toNumber()] = _377_v25;
        _nw51[(new BigNumber(23)).toNumber()] = _377_v25;
        _379_v27 = _nw51;
        let _380_v28;
        _380_v28 = _dafny.MultiSet.fromElements(_379_v27, _379_v27, _379_v27);
        let _index39 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_344_v2).length));
        (_344_v2)[_index39] = !((_380_v28).IsSubsetOf(_380_v28));
        let _index40 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_344_v2).length));
        (_344_v2)[_index40] = _346_v4;
      }
      let _381_v29;
      _381_v29 = _module.D13.create_DC43(_module.__default.fm51(_346_v4, _346_v4, _346_v4, _342_v0, _352_globalState));
      let _382_v30;
      _382_v30 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_342_v0);
      let _383_v31;
      _383_v31 = new _dafny.CodePoint('m'.codePointAt(0));
      let _384_v32;
      let _385_v33;
      let _out3;
      let _out4;
      let _outcollector1 = _module.__default.m11(_381_v29, new BigNumber(((_382_v30).update(new BigNumber((_dafny.Seq.update(_345_v3, _module.__default.safeIndex(_342_v0, new BigNumber((_345_v3).length)), _383_v31)).length), _342_v0)).length), _346_v4, _346_v4, _352_globalState);
      _out3 = _outcollector1[0];
      _out4 = _outcollector1[1];
      _384_v32 = _out3;
      _385_v33 = _out4;
      if (((!(!((true) || (_346_v4)))) ? (_346_v4) : ((_384_v32).isLessThanOrEqualTo(_384_v32)))) {
        let _386_v34;
        let _nw52 = Array((new BigNumber(18)).toNumber()).fill([]);
        _386_v34 = _nw52;
        let _387_v35;
        _387_v35 = _module.D16.create_DC52(_342_v0, true, _386_v34, _346_v4, _342_v0);
        let _388_v36;
        _388_v36 = _module.D2.create_DC9(_346_v4, _342_v0, new _dafny.CodePoint('l'.codePointAt(0)));
        let _389_v37;
        let _390_v38;
        let _out5;
        let _out6;
        let _outcollector2 = _module.__default.m11(_381_v29, (_384_v32).multipliedBy((_387_v35).dtor_cf88), _dafny.Seq.contains(_345_v3, (_388_v36).dtor_cf15), _346_v4, _352_globalState);
        _out5 = _outcollector2[0];
        _out6 = _outcollector2[1];
        _389_v37 = _out5;
        _390_v38 = _out6;
        let _391_v39;
        let _init2 = ((_392_v0) => function (_393_i2) {
          return _module.__default.safeModuloInt(_393_i2, (_module.D1.create_DC4(_392_v0, _dafny.ZERO)).dtor_cf5);
        })(_342_v0);
        let _nw53 = Array((new BigNumber(15)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw53.length); _i0_2++) {
          _nw53[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _391_v39 = _nw53;
        let _index41 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length));
        (_391_v39)[_index41] = _384_v32;
        let _index42 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length));
        (_391_v39)[_index42] = (_dafny.ZERO).minus(_384_v32);
        if (_346_v4) {
          let _394_v40;
          _394_v40 = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,_353_v8);
          let _395_v41;
          _395_v41 = _dafny.Map.Empty.slice().updateUnsafe(_383_v31,new BigNumber((_343_v1).cardinality()));
          let _396_v42;
          _396_v42 = _dafny.Set.fromElements(_355_v10);
          let _397_v43;
          _397_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pshth"),_dafny.MultiSet.FromArray(_353_v8));
          let _398_v44;
          _398_v44 = _dafny.MultiSet.fromElements(_346_v4);
          let _399_v45;
          _399_v45 = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,(_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))]);
          let _400_v46;
          _400_v46 = _dafny.Map.Empty.slice().updateUnsafe((((_399_v45).contains(_346_v4)) ? ((_399_v45).get(_346_v4)) : (new BigNumber((_dafny.Seq.update(_355_v10, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-625)), function (_401_i4) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          })).length), new BigNumber((_355_v10).length)), _389_v37)).length))),_346_v4);
          let _402_v47;
          _402_v47 = _module.D1.create_DC5(_346_v4, false);
          let _pat_let_tv14 = _346_v4;
          let _403_v48;
          _403_v48 = _module.D11.create_DC34(_398_v44, _343_v1, _344_v2, function (_pat_let10_0) {
  return function (_404_dt__update__tmp_h3) {
    return function (_pat_let11_0) {
      return function (_405_dt__update_hcf7_h0) {
        return _module.D1.create_DC5((_404_dt__update__tmp_h3).dtor_cf6, _405_dt__update_hcf7_h0);
      }(_pat_let11_0);
    }(_pat_let_tv14);
  }(_pat_let10_0);
}(_402_v47));
          let _pat_let_tv15 = _398_v44;
          let _pat_let_tv16 = _344_v2;
          let _406_v49;
          let _nw54 = Array((new BigNumber(26)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray((((_394_v40).contains(_module.__default.fm23(new BigNumber((_353_v8).length), _346_v4, _395_v41, _396_v42, _352_globalState))) ? ((_394_v40).get(_module.__default.fm23(new BigNumber((_353_v8).length), _346_v4, _395_v41, _396_v42, _352_globalState))) : (_dafny.Seq.update(_353_v8, _module.__default.safeIndex((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))], new BigNumber((_353_v8).length)), _module.__default.fm23(new BigNumber((_353_v8).length), !(_346_v4), _module.__default.fm24(_389_v37, _346_v4, _346_v4, (_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))], _352_globalState), _396_v42, _352_globalState)))));
          _nw54[(_dafny.ONE).toNumber()] = ((((_397_v43).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), ((_409_v31) => function (_410_i3) {
            return _409_v31;
          })(_383_v31)))) ? ((_397_v43).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), ((_407_v31) => function (_408_i3) {
            return _407_v31;
          })(_383_v31)))) : (_398_v44))).Intersect(_dafny.MultiSet.fromElements(_346_v4));
          _nw54[(new BigNumber(2)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(3)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(!(_346_v4), _346_v4));
          _nw54[(new BigNumber(5)).toNumber()] = (_module.__default.fm33(!(false), _352_globalState)).Union(_398_v44);
          _nw54[(new BigNumber(6)).toNumber()] = (_398_v44).Intersect(_398_v44);
          _nw54[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(_353_v8);
          _nw54[(new BigNumber(8)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(false);
          _nw54[(new BigNumber(10)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(11)).toNumber()] = (_398_v44).update(_346_v4, _module.__default.abs(_384_v32));
          _nw54[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_346_v4, !(_346_v4), (((_400_v46).contains(_389_v37)) ? ((_400_v46).get(_389_v37)) : (false)));
          _nw54[(new BigNumber(13)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(true);
          _nw54[(new BigNumber(15)).toNumber()] = ((_398_v44).update(_346_v4, _module.__default.abs(_389_v37))).update(_346_v4, _module.__default.abs(_342_v0));
          _nw54[(new BigNumber(16)).toNumber()] = (_398_v44).Union(_398_v44);
          _nw54[(new BigNumber(17)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(18)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(19)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(20)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.fromElements(_346_v4, !(_346_v4), true, _346_v4);
          _nw54[(new BigNumber(22)).toNumber()] = _398_v44;
          _nw54[(new BigNumber(23)).toNumber()] = (function (_pat_let12_0) {
            return function (_411_dt__update__tmp_h4) {
              return function (_pat_let13_0) {
                return function (_412_dt__update_hcf58_h0) {
                  return function (_pat_let14_0) {
                    return function (_413_dt__update_hcf60_h0) {
                      return _module.D11.create_DC34(_412_dt__update_hcf58_h0, (_411_dt__update__tmp_h4).dtor_cf59, _413_dt__update_hcf60_h0, (_411_dt__update__tmp_h4).dtor_cf61);
                    }(_pat_let14_0);
                  }(_pat_let_tv16);
                }(_pat_let13_0);
              }(_pat_let_tv15);
            }(_pat_let12_0);
          }(_403_v48)).dtor_cf58;
          _nw54[(new BigNumber(24)).toNumber()] = (_398_v44).Difference(_398_v44);
          _nw54[(new BigNumber(25)).toNumber()] = _398_v44;
          _406_v49 = _nw54;
          let _414_v50;
          let _nw55 = Array((new BigNumber(9)).toNumber());
          _414_v50 = _nw55;
          let _415_v51;
          _415_v51 = _dafny.Map.Empty.slice().updateUnsafe(_414_v50,_389_v37);
          let _416_v52;
          let _init3 = ((_417_v31) => function (_418_i5) {
            return _417_v31;
          })(_383_v31);
          let _nw56 = Array((_dafny.ONE).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw56.length); _i0_3++) {
            _nw56[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _416_v52 = _nw56;
          let _419_v53;
          _419_v53 = _module.D16.create_DC51(_416_v52);
          let _rhs40 = _module.__default.fm0((new BigNumber((_415_v51).length)).isLessThanOrEqualTo(_384_v32), (_dafny.ZERO).minus(_384_v32), _384_v32, _343_v1, _352_globalState);
          let _rhs41 = _389_v37;
          let _rhs42 = _383_v31;
          let _rhs43 = _406_v49;
          let _rhs44 = (_353_v8)[_module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.fromElements(_416_v52)).update((_419_v53).dtor_cf83, _module.__default.abs(new BigNumber(35)))).cardinality()), new BigNumber((_353_v8).length))];
          let _lhs38 = _352_globalState;
          let _lhs39 = _352_globalState;
          _lhs38.f3 = _rhs40;
          _389_v37 = _rhs41;
          _383_v31 = _rhs42;
          _406_v49 = _rhs43;
          _lhs39.f9 = _rhs44;
          let _420_v54;
          let _nw57 = Array((new BigNumber(7)).toNumber());
          _420_v54 = _nw57;
          let _421_v55;
          _421_v55 = _module.D14.create_DC44(_349_v7);
          let _index43 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_420_v54).length));
          (_420_v54)[_index43] = _421_v55;
          let _pat_let_tv17 = _349_v7;
          let _index44 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_420_v54).length));
          (_420_v54)[_index44] = function (_pat_let15_0) {
            return function (_422_dt__update__tmp_h5) {
              return function (_pat_let16_0) {
                return function (_423_dt__update_hcf71_h0) {
                  return _module.D14.create_DC44(_423_dt__update_hcf71_h0);
                }(_pat_let16_0);
              }(_pat_let_tv17);
            }(_pat_let15_0);
          }(_421_v55);
          let _424_v56;
          _424_v56 = _dafny.Map.Empty.slice().updateUnsafe((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))],(_398_v44).update(_346_v4, _module.__default.abs(_384_v32)));
          let _425_v57;
          let _nw58 = new _module.C5();
          _nw58.__ctor(_389_v37, _346_v4, ((((_424_v56).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_428_v31) => function (_429_i6) {
            return _428_v31;
          })(_383_v31))).length))) ? ((_424_v56).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_426_v31) => function (_427_i6) {
            return _426_v31;
          })(_383_v31))).length))) : (_398_v44))).Union(_dafny.MultiSet.FromArray(_353_v8)));
          _425_v57 = _nw58;
          _383_v31 = _383_v31;
          let _430_v58;
          let _init4 = ((_431_v4, _432_v3, _433_v8) => function (_434_i7) {
            return _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_431_v4, _431_v4), _module.__default.safeIndex(new BigNumber((_432_v3).length), new BigNumber((_dafny.Seq.of(_431_v4, _431_v4)).length)), _431_v4), _433_v8);
          })(_346_v4, _345_v3, _353_v8);
          let _nw59 = Array((new BigNumber(28)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw59.length); _i0_4++) {
            _nw59[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _430_v58 = _nw59;
          let _index45 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_430_v58).length));
          (_430_v58)[_index45] = _353_v8;
          let _index46 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_430_v58).length));
          (_430_v58)[_index46] = _353_v8;
        } else {
          (_352_globalState).f9 = _346_v4;
          _391_v39 = _391_v39;
          (_352_globalState).f3 = _module.__default.fm0(!(_346_v4), _389_v37, _342_v0, (_343_v1).Intersect(_343_v1), _352_globalState);
          let _435_v59;
          let _nw60 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _435_v59 = _nw60;
          let _index47 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_435_v59).length));
          (_435_v59)[_index47] = _353_v8;
          let _436_v60;
          _436_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_353_v8, _dafny.Seq.of(false)),((_346_v4) ? ((_dafny.ZERO).minus(new BigNumber((_module.__default.fm53(_346_v4, _346_v4, _352_globalState)).length))) : ((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))])));
          let _437_v61;
          let _nw61 = new _module.C4();
          _nw61.__ctor(_346_v4, _346_v4, (_dafny.MultiSet.fromElements(_346_v4, _346_v4)).update(_346_v4, _module.__default.abs(new BigNumber((_343_v1).cardinality()))));
          _437_v61 = _nw61;
          let _438_v62;
          _438_v62 = _dafny.Map.Empty.slice().updateUnsafe(_437_v61,_384_v32);
          let _index48 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_435_v59).length));
          let _rhs45 = _module.__default.fm38(_352_globalState);
          let _rhs46 = ((_dafny.Map.Empty.slice().updateUnsafe(_353_v8,new BigNumber((_438_v62).length))).update(_353_v8, new BigNumber(267))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_353_v8, _module.__default.safeIndex(_342_v0, new BigNumber((_353_v8).length)), (_437_v61).f19),_384_v32));
          let _lhs40 = _435_v59;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_435_v59).length));
          _lhs40[_lhs41] = _rhs45;
          _436_v60 = _rhs46;
          let _index49 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length));
          (_391_v39)[_index49] = new BigNumber(-956);
        }
        let _439_v64;
        _439_v64 = _dafny.Map.Empty.slice().updateUnsafe((_355_v10)[_module.__default.safeIndex((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))], new BigNumber((_355_v10).length))],_dafny.Set.fromElements(_346_v4));
        let _440_v65;
        _440_v65 = _module.D0.create_DC2(_module.__default.fm37(function () {
  let _coll35 = new _dafny.Set();
  for (const _compr_35 of (_dafny.Seq.update(_355_v10, _module.__default.safeIndex(new BigNumber((_345_v3).length), new BigNumber((_355_v10).length)), new BigNumber((_343_v1).cardinality()))).Elements) {
    let _441_v63 = _compr_35;
    if (_dafny.Seq.contains(_dafny.Seq.update(_355_v10, _module.__default.safeIndex(new BigNumber((_345_v3).length), new BigNumber((_355_v10).length)), new BigNumber((_343_v1).cardinality())), _441_v63)) {
      _coll35.add(_module.__default.safeDivisionInt(_441_v63, new BigNumber(240)));
    }
  }
  return _coll35;
}(), _439_v64, _346_v4, _346_v4, _352_globalState));
        let _source9 = _440_v65;
        if (_source9.is_DC1) {
          let _442___mcc_h0 = (_source9).cf1;
          let _443_cf1 = _442___mcc_h0;
          (_352_globalState).f3 = _389_v37;
          _384_v32 = (_342_v0).plus(_module.__default.fm0(!(false), _389_v37, new BigNumber((_347_v5).length), _343_v1, _352_globalState));
          let _444_v66;
          let _445_v67;
          let _out7;
          let _out8;
          let _outcollector3 = _module.__default.m11(_381_v29, _389_v37, (new BigNumber(633)).isLessThan((_dafny.ZERO).minus(_384_v32)), false, _352_globalState);
          _out7 = _outcollector3[0];
          _out8 = _outcollector3[1];
          _444_v66 = _out7;
          _445_v67 = _out8;
          _444_v66 = _384_v32;
        } else if (_source9.is_DC0) {
          let _446___mcc_h1 = (_source9).cf0;
          let _447_cf0 = _446___mcc_h1;
          _346_v4 = _447_cf0;
          _389_v37 = _384_v32;
          let _448_v68;
          _448_v68 = _dafny.MultiSet.fromElements(_346_v4, _447_cf0);
          let _449_v69;
          _449_v69 = _dafny.Set.fromElements(_384_v32, (_dafny.ZERO).minus(_342_v0), (((_448_v68).contains(false)) ? ((_448_v68).get(false)) : (_389_v37)));
          (_352_globalState).f9 = (_449_v69).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(981), _342_v0));
          let _450_v70;
          _450_v70 = _dafny.Map.Empty.slice().updateUnsafe(_384_v32,_dafny.Seq.of(_346_v4, _447_cf0));
          (_352_globalState).f9 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((((_450_v70).contains((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))])) ? ((_450_v70).get((_391_v39)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_391_v39).length))])) : (_353_v8)), _353_v8), _353_v8);
        } else {
          let _451___mcc_h2 = (_source9).cf2;
          let _452_cf2 = _451___mcc_h2;
          let _453_v71;
          _453_v71 = _module.D13.create_DC42(_383_v31, !(true));
          let _454_v72;
          let _455_v73;
          let _out9;
          let _out10;
          let _outcollector4 = _module.__default.m11(_module.D13.create_DC43(_453_v71), _389_v37, _346_v4, _346_v4, _352_globalState);
          _out9 = _outcollector4[0];
          _out10 = _outcollector4[1];
          _454_v72 = _out9;
          _455_v73 = _out10;
          _346_v4 = (_343_v1).equals(_343_v1);
          let _456_v74;
          _456_v74 = _module.D5.create_DC16(_391_v39);
          _391_v39 = (_456_v74).dtor_cf29;
          (_352_globalState).f9 = _346_v4;
        }
        _342_v0 = _342_v0;
      } else {
        (_352_globalState).f9 = _dafny.Seq.contains(_dafny.Seq.Concat(_353_v8, _353_v8), ((_346_v4) ? (!(_346_v4)) : (_346_v4)));
        let _457_v75;
        _457_v75 = _dafny.Map.Empty.slice().updateUnsafe(_383_v31,_342_v0);
        let _458_v77;
        let _459_v78;
        let _out11;
        let _out12;
        let _outcollector5 = _module.__default.m11(_381_v29, _342_v0, _346_v4, _module.__default.fm23(_342_v0, _346_v4, _457_v75, function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_342_v0),_342_v0)).Keys.Elements) {
            let _460_v76 = _compr_36;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_342_v0),_342_v0)).contains(_460_v76)) {
              _coll36.add(_460_v76);
            }
          }
          return _coll36;
        }(), _352_globalState), _352_globalState);
        _out11 = _outcollector5[0];
        _out12 = _outcollector5[1];
        _458_v77 = _out11;
        _459_v78 = _out12;
        let _461_v79;
        let _nw62 = new _module.C4();
        _nw62.__ctor(_346_v4, _346_v4, _module.__default.fm33(_346_v4, _352_globalState));
        _461_v79 = _nw62;
        let _462_v80;
        _462_v80 = _dafny.MultiSet.fromElements(_461_v79);
        let _463_v81;
        _463_v81 = _dafny.MultiSet.fromElements(_346_v4, _346_v4);
        let _464_v82;
        _464_v82 = _dafny.Map.Empty.slice().updateUnsafe(_462_v80,(((_463_v81).contains((_461_v79).f19)) ? ((_463_v81).get((_461_v79).f19)) : (_384_v32)));
        _464_v82 = (_464_v82).update(_462_v80, _module.__default.safeModuloInt(_384_v32, _458_v77));
        if (((_346_v4) ? (true) : ((((_461_v79).fm2((_461_v79).f19, _346_v4, _352_globalState)) ? ((_461_v79).f19) : (_346_v4))))) {
          let _465_v83;
          let _init5 = ((_466_v1) => function (_467_i8) {
            return (_467_i8).multipliedBy(new BigNumber((_466_v1).cardinality()));
          })(_343_v1);
          let _nw63 = Array((new BigNumber(12)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw63.length); _i0_5++) {
            _nw63[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _465_v83 = _nw63;
          let _index50 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_465_v83).length));
          (_465_v83)[_index50] = _458_v77;
          let _468_v84;
          _468_v84 = _module.D4.create_DC12(_344_v2);
          let _469_v85;
          _469_v85 = _module.D4.create_DC15(_468_v84);
          let _470_v86;
          let _nw64 = new _module.C6();
          _nw64.__ctor(_346_v4, _463_v81);
          _470_v86 = _nw64;
          let _471_v87;
          _471_v87 = _dafny.Seq.of(_470_v86);
          let _472_v88;
          _472_v88 = _dafny.Map.Empty.slice().updateUnsafe(_471_v87,_module.__default.fm6(_352_globalState));
          let _pat_let_tv18 = _468_v84;
          let _index51 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_465_v83).length));
          let _rhs47 = (_module.__default.safeDivisionInt(_342_v0, new BigNumber(632))).multipliedBy(new BigNumber((_472_v88).length));
          let _rhs48 = _384_v32;
          let _rhs49 = _384_v32;
          let _rhs50 = function (_pat_let17_0) {
            return function (_473_dt__update__tmp_h6) {
              return function (_pat_let18_0) {
                return function (_474_dt__update_hcf28_h0) {
                  return _module.D4.create_DC15(_474_dt__update_hcf28_h0);
                }(_pat_let18_0);
              }(_pat_let_tv18);
            }(_pat_let17_0);
          }(_469_v85);
          let _lhs42 = _465_v83;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_465_v83).length));
          let _lhs44 = _352_globalState;
          let _lhs45 = _352_globalState;
          _lhs42[_lhs43] = _rhs47;
          _lhs44.f3 = _rhs48;
          _lhs45.f3 = _rhs49;
          _469_v85 = _rhs50;
          let _index52 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_465_v83).length));
          (_465_v83)[_index52] = (_465_v83)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_465_v83).length))];
          let _475_v89;
          let _476_v90;
          let _477_v91;
          let _478_v92;
          let _out13;
          let _out14;
          let _out15;
          let _out16;
          let _outcollector6 = (_470_v86).m5(_352_globalState);
          _out13 = _outcollector6[0];
          _out14 = _outcollector6[1];
          _out15 = _outcollector6[2];
          _out16 = _outcollector6[3];
          _475_v89 = _out13;
          _476_v90 = _out14;
          _477_v91 = _out15;
          _478_v92 = _out16;
          let _479_v93;
          let _480_v94;
          let _481_v95;
          let _482_v96;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector7 = (_461_v79).m0(_352_globalState);
          _out17 = _outcollector7[0];
          _out18 = _outcollector7[1];
          _out19 = _outcollector7[2];
          _out20 = _outcollector7[3];
          _479_v93 = _out17;
          _480_v94 = _out18;
          _481_v95 = _out19;
          _482_v96 = _out20;
          let _483_v97;
          let _nw65 = new _module.C7();
          _nw65.__ctor();
          _483_v97 = _nw65;
        } else {
          let _484_v98;
          let _init6 = ((_485_v31) => function (_486_i9) {
            return _485_v31;
          })(_383_v31);
          let _nw66 = Array((new BigNumber(6)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw66.length); _i0_6++) {
            _nw66[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _484_v98 = _nw66;
          let _487_v99;
          _487_v99 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_484_v98);
          let _488_v100;
          _488_v100 = _dafny.Map.Empty.slice().updateUnsafe((_461_v79).f19,_384_v32);
          let _489_v101;
          let _nw67 = Array((new BigNumber(15)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _484_v98;
          _nw67[(_dafny.ONE).toNumber()] = _484_v98;
          _nw67[(new BigNumber(2)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(3)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(4)).toNumber()] = (((_487_v99).contains(_module.__default.fm25(_342_v0, new BigNumber((_355_v10).length), (_461_v79).f19, _488_v100, _352_globalState))) ? ((_487_v99).get(_module.__default.fm25(_342_v0, new BigNumber((_355_v10).length), (_461_v79).f19, _488_v100, _352_globalState))) : (_484_v98));
          _nw67[(new BigNumber(5)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(6)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(7)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(8)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(9)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(10)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(11)).toNumber()] = (((_461_v79).f19) ? (_484_v98) : (_484_v98));
          _nw67[(new BigNumber(12)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(13)).toNumber()] = _484_v98;
          _nw67[(new BigNumber(14)).toNumber()] = _484_v98;
          _489_v101 = _nw67;
          let _index53 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_489_v101).length));
          (_489_v101)[_index53] = _484_v98;
          let _490_v102;
          _490_v102 = _module.D16.create_DC51(_484_v98);
          let _index54 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_489_v101).length));
          (_489_v101)[_index54] = (_490_v102).dtor_cf83;
          let _491_v103;
          let _nw68 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _491_v103 = _nw68;
          let _index55 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_491_v103).length));
          (_491_v103)[_index55] = (((_461_v79).f19) ? (_458_v77) : (_342_v0));
          let _492_v104;
          _492_v104 = _dafny.Map.Empty.slice().updateUnsafe(_491_v103,((_346_v4) ? ((_353_v8)[_module.__default.safeIndex(new BigNumber((_382_v30).length), new BigNumber((_353_v8).length))]) : (_346_v4)));
          let _493_v105;
          _493_v105 = _module.D5.create_DC18((_461_v79).f19);
          let _494_v106;
          _494_v106 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_module.D11.create_DC35());
          let _495_v107;
          _495_v107 = _dafny.Set.fromElements(_345_v3);
          let _index56 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_491_v103).length));
          let _rhs51 = (_461_v79).f19;
          let _rhs52 = (_493_v105).dtor_cf33;
          let _rhs53 = ((!(_494_v106).equals(_494_v106)) ? ((_461_v79).f19) : ((_495_v107).IsDisjointFrom(_495_v107)));
          let _rhs54 = _module.__default.safeModuloInt(_342_v0, _342_v0);
          let _rhs55 = _492_v104;
          let _lhs46 = _352_globalState;
          let _lhs47 = _352_globalState;
          let _lhs48 = _491_v103;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_491_v103).length));
          _lhs46.f9 = _rhs51;
          _lhs47.f9 = _rhs52;
          _346_v4 = _rhs53;
          _lhs48[_lhs49] = _rhs54;
          _492_v104 = _rhs55;
          let _496_v108;
          let _nw69 = new _module.C3();
          _nw69.__ctor((_491_v103)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_491_v103).length))], _345_v3, _346_v4, _463_v81);
          _496_v108 = _nw69;
          let _497_v109;
          _497_v109 = _module.D18.create_DC55(_496_v108);
          let _498_v110;
          _498_v110 = _dafny.MultiSet.fromElements((_497_v109).dtor_cf91, _496_v108);
          let _499_v111;
          _499_v111 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_498_v110).cardinality()),(_496_v108).f20)).update(new BigNumber((_dafny.Seq.update(_353_v8, _module.__default.safeIndex((_491_v103)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_491_v103).length))], new BigNumber((_353_v8).length)), (_461_v79).f19)).length), (_dafny.ZERO).minus((_496_v108).f20)));
          _499_v111 = _499_v111;
          let _index57 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_344_v2).length));
          (_344_v2)[_index57] = true;
          let _index58 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_344_v2).length));
          (_344_v2)[_index58] = (_461_v79).f19;
          let _500_v112;
          _500_v112 = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,_383_v31);
          let _501_v113;
          _501_v113 = _dafny.MultiSet.fromElements((_500_v112).update((_461_v79).f19, _383_v31), _500_v112);
          (_352_globalState).f9 = (_501_v113).IsProperSubsetOf(_501_v113);
        }
        let _502_v114;
        _502_v114 = _dafny.Map.Empty.slice().updateUnsafe(_344_v2,_383_v31);
        let _503_v115;
        _503_v115 = _module.D15.create_DC49((_502_v114).Merge(_502_v114), new BigNumber(-88));
        _503_v115 = _module.D15.create_DC49(_502_v114, _384_v32);
      }
      _342_v0 = _384_v32;
      _346_v4 = _346_v4;
      let _504_v116;
      _504_v116 = _dafny.MultiSet.fromElements(_346_v4, _346_v4, _346_v4, _346_v4, _346_v4);
      let _505_v117;
      let _nw70 = new _module.C6();
      _nw70.__ctor(_346_v4, _504_v116);
      _505_v117 = _nw70;
      if (!_dafny.areEqual(_355_v10, _355_v10)) {
        _342_v0 = _module.__default.fm0(_346_v4, ((_dafny.ZERO).minus(_384_v32)).plus(new BigNumber(359)), _384_v32, _343_v1, _352_globalState);
        let _506_v118;
        _506_v118 = _module.D8.create_DC24(_dafny.Set.fromElements(_346_v4));
        let _507_v119;
        _507_v119 = _dafny.Map.Empty.slice().updateUnsafe(_384_v32,_347_v5);
        let _508_v120;
        _508_v120 = _dafny.Map.Empty.slice().updateUnsafe(_342_v0,_346_v4);
        let _509_v121;
        _509_v121 = _dafny.Seq.of(_347_v5, _dafny.Set.fromElements(_346_v4, _346_v4, true), _dafny.Set.fromElements(_346_v4));
        let _510_v122;
        let _nw71 = Array((new BigNumber(16)).toNumber());
        _nw71[(_dafny.ZERO).toNumber()] = _347_v5;
        _nw71[(_dafny.ONE).toNumber()] = _347_v5;
        _nw71[(new BigNumber(2)).toNumber()] = _347_v5;
        _nw71[(new BigNumber(3)).toNumber()] = (_506_v118).dtor_cf45;
        _nw71[(new BigNumber(4)).toNumber()] = _347_v5;
        _nw71[(new BigNumber(5)).toNumber()] = (((_507_v119).contains(_342_v0)) ? ((_507_v119).get(_342_v0)) : (_347_v5));
        _nw71[(new BigNumber(6)).toNumber()] = (_347_v5).Union(_module.__default.fm41(_384_v32, new BigNumber((_508_v120).length), _352_globalState));
        _nw71[(new BigNumber(7)).toNumber()] = ((_509_v121)[_module.__default.safeIndex(_384_v32, new BigNumber((_509_v121).length))]).Intersect(_dafny.Set.fromElements(_346_v4));
        _nw71[(new BigNumber(8)).toNumber()] = ((_346_v4) ? (_347_v5) : (_347_v5));
        _nw71[(new BigNumber(9)).toNumber()] = _module.__default.fm41(new BigNumber((_355_v10).length), _342_v0, _352_globalState);
        _nw71[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_346_v4, _346_v4);
        _nw71[(new BigNumber(11)).toNumber()] = _347_v5;
        _nw71[(new BigNumber(12)).toNumber()] = (_347_v5).Difference(_dafny.Set.fromElements(_346_v4, _346_v4, _346_v4, _346_v4, _346_v4));
        _nw71[(new BigNumber(13)).toNumber()] = (_347_v5).Union(_347_v5);
        _nw71[(new BigNumber(14)).toNumber()] = _347_v5;
        _nw71[(new BigNumber(15)).toNumber()] = _347_v5;
        _510_v122 = _nw71;
        let _index59 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_510_v122).length));
        (_510_v122)[_index59] = (_347_v5).Intersect(_347_v5);
        let _511_v123;
        let _nw72 = Array((new BigNumber(10)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_342_v0);
        _nw72[(_dafny.ONE).toNumber()] = (_355_v10)[_module.__default.safeIndex(new BigNumber((_345_v3).length), new BigNumber((_355_v10).length))];
        _nw72[(new BigNumber(2)).toNumber()] = new BigNumber((_353_v8).length);
        _nw72[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_384_v32, _342_v0);
        _nw72[(new BigNumber(4)).toNumber()] = _342_v0;
        _nw72[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_342_v0).plus(_342_v0));
        _nw72[(new BigNumber(6)).toNumber()] = _384_v32;
        _nw72[(new BigNumber(7)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_346_v4,_383_v31)).update(_346_v4, _383_v31)).length);
        _nw72[(new BigNumber(8)).toNumber()] = (_384_v32).minus(_384_v32);
        _nw72[(new BigNumber(9)).toNumber()] = _342_v0;
        _511_v123 = _nw72;
        let _index60 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_511_v123).length));
        (_511_v123)[_index60] = _342_v0;
        let _512_v124;
        _512_v124 = _module.D12.create_DC39(_384_v32);
        let _index61 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_510_v122).length));
        let _index62 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_511_v123).length));
        let _rhs56 = _module.__default.fm41(new BigNumber((_345_v3).length), (_512_v124).dtor_cf65, _352_globalState);
        let _rhs57 = (_347_v5).Intersect(_347_v5);
        let _rhs58 = (_dafny.ZERO).minus((_342_v0).plus(new BigNumber((_345_v3).length)));
        let _rhs59 = _module.__default.fm0(_346_v4, _342_v0, new BigNumber(-571), _343_v1, _352_globalState);
        let _lhs50 = _510_v122;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_510_v122).length));
        let _lhs52 = _352_globalState;
        let _lhs53 = _511_v123;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_511_v123).length));
        _lhs50[_lhs51] = _rhs56;
        _lhs52.f6 = _rhs57;
        _lhs53[_lhs54] = _rhs58;
        _384_v32 = _rhs59;
        _346_v4 = !(_346_v4);
        _346_v4 = true;
        let _index63 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_511_v123).length));
        (_511_v123)[_index63] = _342_v0;
      } else {
        if (_346_v4) {
          _504_v116 = _504_v116;
          let _513_v125;
          _513_v125 = _dafny.Set.fromElements(new BigNumber(78), _342_v0);
          let _514_v126;
          _514_v126 = _dafny.Map.Empty.slice().updateUnsafe(_383_v31,_513_v125);
          let _515_v127;
          let _516_v128;
          let _517_v129;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector8 = (_505_v117).m1(((((_514_v126).contains(_383_v31)) ? ((_514_v126).get(_383_v31)) : (_513_v125))).Union(_dafny.Set.fromElements(new BigNumber(-222))), new BigNumber(30), _352_globalState);
          _out21 = _outcollector8[0];
          _out22 = _outcollector8[1];
          _out23 = _outcollector8[2];
          _515_v127 = _out21;
          _516_v128 = _out22;
          _517_v129 = _out23;
          let _518_v130;
          let _nw73 = new _module.C3();
          _nw73.__ctor(_516_v128, _345_v3, true, _504_v116);
          _518_v130 = _nw73;
          let _519_v131;
          _519_v131 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_516_v128, _516_v128),_518_v130);
          _519_v131 = (_519_v131).update(_384_v32, _518_v130);
          let _520_v132;
          _520_v132 = _dafny.Map.Empty.slice().updateUnsafe(_516_v128,_344_v2);
          let _521_v133;
          let _nw74 = new _module.C2();
          _nw74.__ctor((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_520_v132).length),_346_v4)).update(_384_v32, _346_v4));
          _521_v133 = _nw74;
          let _522_v134;
          _522_v134 = _dafny.Map.Empty.slice().updateUnsafe(_517_v129,_342_v0);
          _522_v134 = (_522_v134).update(_517_v129, _342_v0);
        } else {
          let _523_v135;
          let _524_v136;
          let _525_v137;
          let _526_v138;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector9 = (_505_v117).m5(_352_globalState);
          _out24 = _outcollector9[0];
          _out25 = _outcollector9[1];
          _out26 = _outcollector9[2];
          _out27 = _outcollector9[3];
          _523_v135 = _out24;
          _524_v136 = _out25;
          _525_v137 = _out26;
          _526_v138 = _out27;
          let _527_v139;
          let _nw75 = new _module.C7();
          _nw75.__ctor();
          _527_v139 = _nw75;
          let _nw76 = new _module.C7();
          _nw76.__ctor();
          _527_v139 = _nw76;
          let _528_v140;
          _528_v140 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_346_v4,_346_v4)).length),_526_v138);
          let _529_v141;
          let _nw77 = Array((new BigNumber(28)).toNumber()).fill([]);
          _529_v141 = _nw77;
          let _530_v142;
          _530_v142 = _module.D16.create_DC52(_342_v0, _526_v138, (_module.D16.create_DC52(_524_v136, _526_v138, _529_v141, true, _342_v0)).dtor_cf86, _346_v4, _524_v136);
          _528_v140 = (_528_v140).update(((_530_v142).dtor_cf88).multipliedBy(_384_v32), _346_v4);
          let _531_v143;
          let _init7 = ((_532_v136) => function (_533_i10) {
            return _module.__default.safeDivisionInt(_533_i10, _532_v136);
          })(_524_v136);
          let _nw78 = Array((new BigNumber(11)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw78.length); _i0_7++) {
            _nw78[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _531_v143 = _nw78;
          let _index64 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_531_v143).length));
          (_531_v143)[_index64] = (_355_v10)[_module.__default.safeIndex(_342_v0, new BigNumber((_355_v10).length))];
          let _534_v144;
          _534_v144 = _dafny.Map.Empty.slice().updateUnsafe(_526_v138,_384_v32);
          let _535_v145;
          _535_v145 = _dafny.Set.fromElements(_384_v32);
          let _536_v147;
          let _nw79 = new _module.C5();
          _nw79.__ctor(_524_v136, _526_v138, _504_v116);
          _536_v147 = _nw79;
          let _537_v148;
          _537_v148 = _module.D19.create_DC59(_536_v147, _534_v144);
          let _538_v149;
          let _nw80 = Array((new BigNumber(29)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = (_534_v144).update(_526_v138, _523_v135);
          _nw80[(_dafny.ONE).toNumber()] = _534_v144;
          _nw80[(new BigNumber(2)).toNumber()] = (_534_v144).Merge(_534_v144);
          _nw80[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_526_v138, _384_v32, _346_v4, _535_v145, _352_globalState);
          _nw80[(new BigNumber(4)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,_384_v32);
          _nw80[(new BigNumber(6)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(7)).toNumber()] = ((!(_346_v4)) ? (_module.__default.fm1(_346_v4, _384_v32, !(_526_v138), function () {
            let _coll37 = new _dafny.Set();
            for (const _compr_37 of _dafny.IntegerRange(new BigNumber(994), new BigNumber(775))) {
              let _539_v146 = _compr_37;
              if (((new BigNumber(994)).isLessThanOrEqualTo(_539_v146)) && ((_539_v146).isLessThan(new BigNumber(775)))) {
                _coll37.add(_module.__default.safeDivisionInt(_539_v146, _384_v32));
              }
            }
            return _coll37;
          }(), _352_globalState)) : (_534_v144));
          _nw80[(new BigNumber(8)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(9)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(10)).toNumber()] = _module.__default.fm1(_346_v4, new BigNumber(705), _346_v4, _535_v145, _352_globalState);
          _nw80[(new BigNumber(11)).toNumber()] = (_537_v148).dtor_cf98;
          _nw80[(new BigNumber(12)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_526_v138,_523_v135)).Merge(_534_v144);
          _nw80[(new BigNumber(14)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(15)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(16)).toNumber()] = (_534_v144).Merge(_dafny.Map.Empty.slice().updateUnsafe(_526_v138,_384_v32));
          _nw80[(new BigNumber(17)).toNumber()] = (_module.__default.fm1(!(true), _524_v136, _526_v138, _535_v145, _352_globalState)).Merge(_534_v144);
          _nw80[(new BigNumber(18)).toNumber()] = (_534_v144).Merge((_534_v144).update(_346_v4, _384_v32));
          _nw80[(new BigNumber(19)).toNumber()] = (_module.D19.create_DC59(_536_v147, _534_v144)).dtor_cf98;
          _nw80[(new BigNumber(20)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(21)).toNumber()] = (_534_v144).Merge(_534_v144);
          _nw80[(new BigNumber(22)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(23)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(24)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(25)).toNumber()] = _534_v144;
          _nw80[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_342_v0);
          _nw80[(new BigNumber(27)).toNumber()] = (_534_v144).Merge(_module.__default.fm1(_526_v138, _523_v135, _526_v138, _535_v145, _352_globalState));
          _nw80[(new BigNumber(28)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_536_v147.f14,_342_v0);
          _538_v149 = _nw80;
          let _540_v150;
          _540_v150 = _dafny.Seq.of(_534_v144);
          let _index65 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_538_v149).length));
          (_538_v149)[_index65] = ((_540_v150)[_module.__default.safeIndex(_523_v135, new BigNumber((_540_v150).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(_536_v147.f14,_524_v136));
          let _index66 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_531_v143).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_538_v149).length));
          let _rhs60 = _383_v31;
          let _rhs61 = _384_v32;
          let _rhs62 = ((_534_v144).Merge(_534_v144)).Merge(((_534_v144).update(_346_v4, _523_v135)).Merge(_534_v144));
          let _rhs63 = _342_v0;
          let _rhs64 = _383_v31;
          let _lhs55 = _531_v143;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_531_v143).length));
          let _lhs57 = _538_v149;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_538_v149).length));
          _383_v31 = _rhs60;
          _lhs55[_lhs56] = _rhs61;
          _lhs57[_lhs58] = _rhs62;
          _524_v136 = _rhs63;
          _383_v31 = _rhs64;
          _342_v0 = (_524_v136).minus(_module.__default.safeDivisionInt(_523_v135, new BigNumber((_345_v3).length)));
        }
        let _541_v151;
        let _nw81 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _541_v151 = _nw81;
        _541_v151 = _541_v151;
        let _index68 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_541_v151).length));
        (_541_v151)[_index68] = new BigNumber(815);
        let _index69 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_541_v151).length));
        (_541_v151)[_index69] = ((true) ? (_384_v32) : (_384_v32));
        _345_v3 = _345_v3;
        let _index70 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_541_v151).length));
        (_541_v151)[_index70] = (_541_v151)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_541_v151).length))];
      }
      let _542_i11;
      _542_i11 = _dafny.ZERO;
      L0: {
        while (_346_v4) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_542_i11)) {
              break L0;
            }
            _542_i11 = (_542_i11).plus(_dafny.ONE);
            let _543_v152;
            let _init8 = ((_544_v4) => function (_545_i12) {
              return (_545_i12).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_544_v4,_544_v4))))).cardinality()));
            })(_346_v4);
            let _nw82 = Array((new BigNumber(25)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw82.length); _i0_8++) {
              _nw82[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _543_v152 = _nw82;
            let _index71 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_543_v152).length));
            (_543_v152)[_index71] = _384_v32;
            let _index72 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_543_v152).length));
            (_543_v152)[_index72] = _module.__default.safeModuloInt(new BigNumber((_343_v1).cardinality()), _384_v32);
            let _546_v153;
            let _nw83 = Array((new BigNumber(15)).toNumber());
            _546_v153 = _nw83;
            let _547_v154;
            let _nw84 = new _module.C3();
            _nw84.__ctor(_342_v0, _dafny.Seq.UnicodeFromString("ihvjxhpxy"), _346_v4, _504_v116);
            _547_v154 = _nw84;
            let _index73 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_546_v153).length));
            (_546_v153)[_index73] = _547_v154;
            let _index74 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_546_v153).length));
            let _nw85 = new _module.C3();
            _nw85.__ctor(new BigNumber((_dafny.Seq.of(true, _346_v4, _346_v4, _346_v4)).length), _dafny.Seq.Concat((_547_v154).f21, (_547_v154).f21), _346_v4, _504_v116);
            (_546_v153)[_index74] = _nw85;
            let _548_v155;
            _548_v155 = _module.D2.create_DC7(_353_v8);
            let _549_v156;
            _549_v156 = _dafny.Set.fromElements(_548_v155, _548_v155, _548_v155);
            (_352_globalState).f9 = !(_549_v156).contains(_module.D2.create_DC7(_353_v8));
            let _index75 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_344_v2).length));
            (_344_v2)[_index75] = _346_v4;
            let _550_v157;
            _550_v157 = _dafny.Map.Empty.slice().updateUnsafe(_346_v4,((_547_v154).f21)[_module.__default.safeIndex(new BigNumber(-660), new BigNumber(((_547_v154).f21).length))]);
            let _551_v158;
            _551_v158 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_550_v157).length)));
            let _index76 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_344_v2).length));
            (_344_v2)[_index76] = (_module.__default.fm6(_352_globalState)).IsSubsetOf(_551_v158);
          }
        }
      }
      let _552_v159;
      let _init9 = ((_553_v8) => function (_554_i14) {
        return (_554_i14).minus(new BigNumber((_553_v8).length));
      })(_353_v8);
      let _nw86 = Array((new BigNumber(13)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw86.length); _i0_9++) {
        _nw86[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _552_v159 = _nw86;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_552_v159).length))) {
        let _555_i13 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_555_i13)) && ((_555_i13).isLessThan(new BigNumber((_552_v159).length))))) {
          (_552_v159)[(_555_i13)] = (_555_i13).minus(new BigNumber(525));
        }
      }
      let _index77 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
      (_344_v2)[_index77] = _346_v4;
      let _index78 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
      (_344_v2)[_index78] = _346_v4;
      let _556_v160;
      let _557_v161;
      let _558_v162;
      let _559_v163;
      let _out28;
      let _out29;
      let _out30;
      let _out31;
      let _outcollector10 = (_505_v117).m0(_352_globalState);
      _out28 = _outcollector10[0];
      _out29 = _outcollector10[1];
      _out30 = _outcollector10[2];
      _out31 = _outcollector10[3];
      _556_v160 = _out28;
      _557_v161 = _out29;
      _558_v162 = _out30;
      _559_v163 = _out31;
      let _560_v164;
      _560_v164 = _dafny.Map.Empty.slice().updateUnsafe(_556_v160,_346_v4);
      _560_v164 = (_560_v164).update((_dafny.ZERO).minus(_342_v0), (_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))]);
      let _561_i15;
      _561_i15 = _dafny.ZERO;
      L1: {
        while ((new BigNumber((_355_v10).length)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_342_v0, _342_v0))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_561_i15)) {
              break L1;
            }
            _561_i15 = (_561_i15).plus(_dafny.ONE);
            let _562_v165;
            _562_v165 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), function (_563_i16) {
              return new BigNumber(188);
            }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(731)), ((_564_v32) => function (_565_i17) {
              return _564_v32;
            })(_384_v32)), _355_v10);
            _562_v165 = _562_v165;
            let _566_v166;
            _566_v166 = _module.D1.create_DC5(_346_v4, _557_v161);
            let _index79 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
            let _rhs65 = (((_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))]) ? ((_566_v166).dtor_cf6) : (_557_v161));
            let _rhs66 = _559_v163;
            let _rhs67 = (_module.D4.create_DC13(_345_v3, _556_v160, _557_v161, _344_v2)).dtor_cf22;
            let _lhs59 = _344_v2;
            let _lhs60 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
            _559_v163 = _rhs65;
            _lhs59[_lhs60] = _rhs66;
            _345_v3 = _rhs67;
            let _567_v167;
            let _568_v168;
            let _out32;
            let _out33;
            let _outcollector11 = _module.__default.m11(_381_v29, new BigNumber((_345_v3).length), true, (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(307), _384_v32))).IsProperSubsetOf(_343_v1), _352_globalState);
            _out32 = _outcollector11[0];
            _out33 = _outcollector11[1];
            _567_v167 = _out32;
            _568_v168 = _out33;
            let _index80 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
            (_344_v2)[_index80] = false;
          }
        }
      }
      if ((_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))]) {
        let _569_v169;
        _569_v169 = _dafny.Set.fromElements(new BigNumber(600));
        let _570_v170;
        let _571_v171;
        let _572_v172;
        let _out34;
        let _out35;
        let _out36;
        let _outcollector12 = (_505_v117).m1(_569_v169, new BigNumber((_560_v164).length), _352_globalState);
        _out34 = _outcollector12[0];
        _out35 = _outcollector12[1];
        _out36 = _outcollector12[2];
        _570_v170 = _out34;
        _571_v171 = _out35;
        _572_v172 = _out36;
        let _index81 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
        (_344_v2)[_index81] = _559_v163;
        let _573_v173;
        let _nw87 = new _module.C1();
        _nw87.__ctor();
        _573_v173 = _nw87;
        let _574_v174;
        let _nw88 = new _module.C5();
        _nw88.__ctor(_556_v160, !(_557_v161) || (_557_v161), _dafny.MultiSet.fromElements(_346_v4));
        _574_v174 = _nw88;
        let _index82 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
        let _rhs68 = _572_v172;
        let _rhs69 = _574_v174;
        let _lhs61 = _344_v2;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
        _lhs61[_lhs62] = _rhs68;
        _574_v174 = _rhs69;
        _559_v163 = _557_v161;
      } else {
        let _575_v175;
        let _576_v176;
        let _out37;
        let _out38;
        let _outcollector13 = _module.__default.m11(_381_v29, _module.__default.fm0((_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))], _342_v0, _342_v0, _module.__default.fm30(_557_v161, _352_globalState), _352_globalState), (_dafny.MultiSet.FromArray(_353_v8)).IsSubsetOf(_504_v116), _559_v163, _352_globalState);
        _out37 = _outcollector13[0];
        _out38 = _outcollector13[1];
        _575_v175 = _out37;
        _576_v176 = _out38;
        let _577_v177;
        _577_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_560_v164).length),_345_v3);
        let _578_v178;
        _578_v178 = _module.D8.create_DC26(_557_v161, _577_v177);
        let _579_v179;
        _579_v179 = _dafny.Set.fromElements(_578_v178, _578_v178, _578_v178, _578_v178);
        let _580_v180;
        _580_v180 = _dafny.Map.Empty.slice().updateUnsafe(_559_v163,!(_557_v161));
        let _index83 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
        let _rhs70 = _dafny.Seq.Concat(_dafny.Seq.of(_559_v163), _dafny.Seq.of(!(_559_v163), (((_580_v180).contains((_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))])) ? ((_580_v180).get((_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))])) : (_557_v161))));
        let _rhs71 = _559_v163;
        let _rhs72 = _579_v179;
        let _rhs73 = _343_v1;
        let _rhs74 = (_344_v2)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length))];
        let _lhs63 = _344_v2;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_344_v2).length));
        _353_v8 = _rhs70;
        _559_v163 = _rhs71;
        _579_v179 = _rhs72;
        _343_v1 = _rhs73;
        _lhs63[_lhs64] = _rhs74;
        let _581_v181;
        _581_v181 = _dafny.Set.fromElements(_342_v0);
        let _582_v182;
        let _583_v183;
        let _584_v184;
        let _out39;
        let _out40;
        let _out41;
        let _outcollector14 = (_505_v117).m1((_dafny.Set.fromElements((_dafny.ZERO).minus(_556_v160), new BigNumber(314))).Union(_581_v181), _384_v32, _352_globalState);
        _out39 = _outcollector14[0];
        _out40 = _outcollector14[1];
        _out41 = _outcollector14[2];
        _582_v182 = _out39;
        _583_v183 = _out40;
        _584_v184 = _out41;
        (_352_globalState).f3 = _556_v160;
        let _585_v185;
        let _nw89 = new _module.C0();
        _nw89.__ctor(!(_559_v163), !(false));
        _585_v185 = _nw89;
      }
      process.stdout.write(_dafny.toString(_342_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_343_v1).equals(_dafny.MultiSet.fromElements(new BigNumber(206)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_344_v2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_345_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_346_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_347_v5).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_349_v7)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_352_globalState).f1).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState.f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState.f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState.f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_352_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_352_globalState).f5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState.f6).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_352_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_352_globalState).f12)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_352_globalState.f13).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_353_v8, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_354_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(true)).updateUnsafe(true,_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_355_v10, _dafny.Seq.of(new BigNumber(206), new BigNumber(206)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_381_v29).dtor_cf70).dtor_cf70).dtor_cf70).dtor_cf68));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_381_v29).dtor_cf70).dtor_cf70).dtor_cf70).dtor_cf69));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_382_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_383_v31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_384_v32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_385_v33).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_504_v116).equals(_dafny.MultiSet.fromElements(true, true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_542_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_552_v159)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_556_v160));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_557_v161));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_558_v162).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-153),true),new BigNumber(-804)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_559_v163));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_560_v164).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-153539075),true).updateUnsafe(new BigNumber(-2),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_561_i15));
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
    static create_DC2(cf2) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false);
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
    static create_DC4(cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC5(cf6, cf7) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D1(3);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6 && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC8(cf10, cf11, cf12) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC9(cf13, cf14, cf15) {
      let $dt = new D2(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf9) {
      let $dt = new D2(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC11(cf17, cf18, cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(_dafny.ZERO, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC13(cf22, cf23, cf24, cf25) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf26, cf27) {
      let $dt = new D4(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D4(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D4(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + this.cf22.toVerbatimString(true) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf21 === other.cf21;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false, []);
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
    static create_DC17(cf30, cf31, cf32) {
      let $dt = new D5(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D5(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D5(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D5(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf33 === other.cf33;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf29 === other.cf29;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.Seq.of(), false, _dafny.MultiSet.Empty);
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
    static create_DC20(cf35) {
      let $dt = new D6(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf37, cf38, cf39, cf40, cf41) {
      let $dt = new D7(0);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf42, cf43, cf44) {
      let $dt = new D7(1);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC21(cf36) {
      let $dt = new D7(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC23" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22([], _dafny.Seq.of(), _dafny.Seq.of(), [], _module.D5.Default());
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
    static create_DC25() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC26(cf46, cf47) {
      let $dt = new D8(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC27() {
      let $dt = new D8(2);
      return $dt;
    }
    static create_DC24(cf45) {
      let $dt = new D8(3);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC28(cf48) {
      let $dt = new D8(4);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get is_DC28() { return this.$tag === 4; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC25";
      } else if (this.$tag === 1) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC27";
      } else if (this.$tag === 3) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC28" + "(" + _dafny.toString(this.cf48) + ")";
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
        return other.$tag === 1 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC25();
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
    static create_DC30(cf50, cf51, cf52, cf53, cf54) {
      let $dt = new D9(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC29(cf49) {
      let $dt = new D9(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC31(cf55) {
      let $dt = new D9(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC30" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC29" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC31" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50) && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC30(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC32(cf56) {
      let $dt = new D10(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC32" + "(" + _dafny.toString(this.cf56) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC34(cf58, cf59, cf60, cf61) {
      let $dt = new D11(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC35() {
      let $dt = new D11(1);
      return $dt;
    }
    static create_DC33(cf57) {
      let $dt = new D11(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC36(cf62) {
      let $dt = new D11(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC35";
      } else if (this.$tag === 2) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC36" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC34(_dafny.MultiSet.Empty, _dafny.MultiSet.Empty, [], _module.D1.Default());
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
    static create_DC38(cf64) {
      let $dt = new D12(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC39(cf65) {
      let $dt = new D12(1);
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC37(cf63) {
      let $dt = new D12(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC40(cf66) {
      let $dt = new D12(3);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC40() { return this.$tag === 3; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC38" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC39" + "(" + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC40" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC38(_dafny.ZERO);
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
    static create_DC42(cf68, cf69) {
      let $dt = new D13(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC41(cf67) {
      let $dt = new D13(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC43(cf70) {
      let $dt = new D13(2);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC42" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC41" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC43" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC42(new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC45(cf72, cf73, cf74) {
      let $dt = new D14(0);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC44(cf71) {
      let $dt = new D14(1);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC46(cf75) {
      let $dt = new D14(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC45" + "(" + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC44" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC46" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf72, other.cf72) && this.cf73 === other.cf73 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf71 === other.cf71;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf75, other.cf75);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC45(_dafny.Map.Empty, null, _dafny.ZERO);
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
    static create_DC48(cf77, cf78, cf79) {
      let $dt = new D15(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC49(cf80, cf81) {
      let $dt = new D15(1);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC47(cf76) {
      let $dt = new D15(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC50(cf82) {
      let $dt = new D15(3);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get is_DC50() { return this.$tag === 3; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC48" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC49" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC47" + "(" + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC50" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC48(_dafny.ZERO, null, _dafny.ZERO);
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
    static create_DC52(cf84, cf85, cf86, cf87, cf88) {
      let $dt = new D16(0);
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC51(cf83) {
      let $dt = new D16(1);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC53(cf89) {
      let $dt = new D16(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC53() { return this.$tag === 2; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC52" + "(" + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC51" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC53" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84) && this.cf85 === other.cf85 && this.cf86 === other.cf86 && this.cf87 === other.cf87 && _dafny.areEqual(this.cf88, other.cf88);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf83 === other.cf83;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC52(_dafny.ZERO, false, [], false, _dafny.ZERO);
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
    static create_DC54(cf90) {
      let $dt = new D17(0);
      $dt.cf90 = cf90;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get dtor_cf90() { return this.cf90; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC54" + "(" + _dafny.toString(this.cf90) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf90 === other.cf90;
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC56(cf92, cf93, cf94) {
      let $dt = new D18(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC57(cf95) {
      let $dt = new D18(1);
      $dt.cf95 = cf95;
      return $dt;
    }
    static create_DC55(cf91) {
      let $dt = new D18(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC56" + "(" + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC57" + "(" + _dafny.toString(this.cf95) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC55" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92 && this.cf93 === other.cf93 && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf95 === other.cf95;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf91 === other.cf91;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC56(false, false, _dafny.ZERO);
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
    static create_DC59(cf97, cf98) {
      let $dt = new D19(0);
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC58(cf96) {
      let $dt = new D19(1);
      $dt.cf96 = cf96;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf96() { return this.cf96; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC59" + "(" + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC58" + "(" + _dafny.toString(this.cf96) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf97 === other.cf97 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf96, other.cf96);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC59(null, _dafny.Map.Empty);
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
    static create_DC61() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC60(cf99) {
      let $dt = new D20(1);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC62(cf100) {
      let $dt = new D20(2);
      $dt.cf100 = cf100;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC61";
      } else if (this.$tag === 1) {
        return "D20.DC60" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC62" + "(" + _dafny.toString(this.cf100) + ")";
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
        return other.$tag === 1 && this.cf99 === other.cf99;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf100, other.cf100);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC61();
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
    static create_DC64(cf102, cf103) {
      let $dt = new D21(0);
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC65(cf104, cf105, cf106) {
      let $dt = new D21(1);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      return $dt;
    }
    static create_DC63(cf101) {
      let $dt = new D21(2);
      $dt.cf101 = cf101;
      return $dt;
    }
    static create_DC66(cf107) {
      let $dt = new D21(3);
      $dt.cf107 = cf107;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC63() { return this.$tag === 2; }
    get is_DC66() { return this.$tag === 3; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf107() { return this.cf107; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC64" + "(" + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC65" + "(" + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC63" + "(" + _dafny.toString(this.cf101) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC66" + "(" + _dafny.toString(this.cf107) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf102 === other.cf102 && this.cf103 === other.cf103;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf104, other.cf104) && _dafny.areEqual(this.cf105, other.cf105) && this.cf106 === other.cf106;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf101 === other.cf101;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf107, other.cf107);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC64(false, null);
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
    static create_DC68(cf109, cf110) {
      let $dt = new D22(0);
      $dt.cf109 = cf109;
      $dt.cf110 = cf110;
      return $dt;
    }
    static create_DC69(cf111) {
      let $dt = new D22(1);
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC70(cf112, cf113, cf114, cf115, cf116) {
      let $dt = new D22(2);
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      $dt.cf115 = cf115;
      $dt.cf116 = cf116;
      return $dt;
    }
    static create_DC67(cf108) {
      let $dt = new D22(3);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC69() { return this.$tag === 1; }
    get is_DC70() { return this.$tag === 2; }
    get is_DC67() { return this.$tag === 3; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC68" + "(" + _dafny.toString(this.cf109) + ", " + _dafny.toString(this.cf110) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC69" + "(" + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC70" + "(" + _dafny.toString(this.cf112) + ", " + _dafny.toString(this.cf113) + ", " + this.cf114.toVerbatimString(true) + ", " + _dafny.toString(this.cf115) + ", " + _dafny.toString(this.cf116) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC67" + "(" + this.cf108.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf109 === other.cf109 && _dafny.areEqual(this.cf110, other.cf110);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf111 === other.cf111;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf112 === other.cf112 && this.cf113 === other.cf113 && _dafny.areEqual(this.cf114, other.cf114) && _dafny.areEqual(this.cf115, other.cf115) && _dafny.areEqual(this.cf116, other.cf116);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC68(false, _dafny.ZERO);
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = [];
      this.f3 = _dafny.ZERO;
      this.f6 = _dafny.Set.Empty;
      this.f9 = false;
      this.f11 = [];
      this.f13 = _dafny.Set.Empty;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.MultiSet.Empty;
      this._f4 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
      this._f7 = false;
      this._f8 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f12 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f23 = false;
      this._f24 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f23, f24) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      return;
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
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
    m10(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      if (p2) {
        let _586_v0;
        let _nw90 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _586_v0 = _nw90;
        let _587_v1;
        _587_v1 = _module.D5.create_DC16(p1);
        _586_v0 = (_587_v1).dtor_cf29;
        let _588_v2;
        _588_v2 = _dafny.Seq.UnicodeFromString("dhcjdj");
        let _589_v3;
        _589_v3 = _dafny.Seq.of(_588_v2);
        let _590_v4;
        _590_v4 = new BigNumber(-738);
        let _591_v5;
        let _nw91 = new _module.C0();
        _nw91.__ctor(_dafny.Seq.IsPrefixOf((_589_v3)[_module.__default.safeIndex(_590_v4, new BigNumber((_589_v3).length))], _dafny.Seq.UnicodeFromString("rylyonx")), true);
        _591_v5 = _nw91;
        let _592_v6;
        let _nw92 = new _module.C0();
        _nw92.__ctor(!((_591_v5).f23), (true) || ((_591_v5).f23));
        _592_v6 = _nw92;
        let _593_v7;
        _593_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,_590_v4);
        r3 = ((_dafny.ZERO).minus(_590_v4)).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((((_593_v7).contains(p2)) ? ((_593_v7).get(p2)) : (_590_v4))), _590_v4));
        let _594_v8;
        _594_v8 = _dafny.Seq.of(_590_v4);
        _594_v8 = _594_v8;
      } else {
        let _595_v9;
        let _nw93 = Array((new BigNumber(7)).toNumber()).fill(false);
        _595_v9 = _nw93;
        (globalState).f2 = _595_v9;
        let _596_v10;
        _596_v10 = new BigNumber(489);
        let _597_v11;
        _597_v11 = _dafny.Seq.of(_596_v10);
        let _598_v12;
        _598_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _599_v13;
        _599_v13 = _dafny.MultiSet.fromElements(p2, p2);
        let _600_v14;
        let _nw94 = Array((new BigNumber(27)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = new BigNumber((_597_v11).length);
        _nw94[(_dafny.ONE).toNumber()] = _596_v10;
        _nw94[(new BigNumber(2)).toNumber()] = (_596_v10).minus(new BigNumber(620));
        _nw94[(new BigNumber(3)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(4)).toNumber()] = new BigNumber((_598_v12).length);
        _nw94[(new BigNumber(5)).toNumber()] = _module.__default.fm0(p2, _596_v10, _596_v10, _dafny.MultiSet.fromElements(_596_v10, _596_v10, _596_v10), globalState);
        _nw94[(new BigNumber(6)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(7)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(8)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(9)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(10)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(11)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(12)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(13)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(14)).toNumber()] = _module.__default.fm0(p2, _596_v10, _596_v10, _dafny.MultiSet.fromElements(new BigNumber(550), _596_v10), globalState);
        _nw94[(new BigNumber(15)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_596_v10, _596_v10);
        _nw94[(new BigNumber(17)).toNumber()] = new BigNumber(((_599_v13).Intersect(_599_v13)).cardinality());
        _nw94[(new BigNumber(18)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(19)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(20)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(21)).toNumber()] = new BigNumber((_597_v11).length);
        _nw94[(new BigNumber(22)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(23)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(24)).toNumber()] = new BigNumber(589);
        _nw94[(new BigNumber(25)).toNumber()] = _596_v10;
        _nw94[(new BigNumber(26)).toNumber()] = _596_v10;
        _600_v14 = _nw94;
        _600_v14 = p1;
        let _601_v15;
        _601_v15 = _dafny.Seq.of(p2, p2);
        let _index84 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length));
        (p1)[_index84] = new BigNumber((_601_v15).length);
        let _index85 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length));
        (p1)[_index85] = _596_v10;
        let _602_v16;
        _602_v16 = _dafny.Set.fromElements(p2);
        let _603_v17;
        _603_v17 = _dafny.MultiSet.fromElements(new BigNumber((_602_v16).length));
        let _604_v18;
        _604_v18 = _module.D2.create_DC7(_601_v15);
        let _605_v19;
        _605_v19 = _dafny.Seq.of(_604_v18, _604_v18, _604_v18, _604_v18, _604_v18);
        _597_v11 = _module.__default.fm21(_module.D1.create_DC3(_module.__default.fm0(p2, _596_v10, (p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))], (_603_v17).update(_596_v10, _module.__default.abs(_596_v10)), globalState)), (_601_v15)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))], new BigNumber((_601_v15).length))], _605_v19, globalState);
        let _606_v20;
        _606_v20 = _dafny.Set.fromElements(_597_v11, _597_v11, _597_v11, _597_v11);
        let _607_v21;
        _607_v21 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p2)),new BigNumber((_606_v20).length));
        let _608_v22;
        _608_v22 = _dafny.Seq.UnicodeFromString("ye");
        let _609_v23;
        _609_v23 = _dafny.Set.fromElements(_596_v10);
        let _610_v24;
        _610_v24 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((((_603_v17).contains(new BigNumber(651))) ? ((_603_v17).get(new BigNumber(651))) : (new BigNumber((_609_v23).length))),p2));
        let _611_v25;
        let _nw95 = Array((new BigNumber(15)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = _607_v21;
        _nw95[(_dafny.ONE).toNumber()] = ((p2) ? (_607_v21) : (_607_v21));
        _nw95[(new BigNumber(2)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(3)).toNumber()] = (_607_v21).Merge(_607_v21);
        _nw95[(new BigNumber(4)).toNumber()] = (_module.__default.fm1((((_598_v12).contains(p2)) ? ((_598_v12).get(p2)) : (p2)), (p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))], true, _module.__default.fm22(new BigNumber((_608_v22).length), _610_v24, p2, p2, globalState), globalState)).update(p2, _596_v10);
        _nw95[(new BigNumber(5)).toNumber()] = (_607_v21).update(false, (p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))]);
        _nw95[(new BigNumber(6)).toNumber()] = (_607_v21).Merge(_607_v21);
        _nw95[(new BigNumber(7)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(8)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(9)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(10)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(11)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(12)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(13)).toNumber()] = _607_v21;
        _nw95[(new BigNumber(14)).toNumber()] = (_607_v21).Merge(_607_v21);
        _611_v25 = _nw95;
        let _index86 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_611_v25).length));
        (_611_v25)[_index86] = (_607_v21).Merge((_607_v21).update(p2, (p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))]));
        let _612_v26;
        _612_v26 = _module.D1.create_DC5(true, p2);
        let _613_v27;
        _613_v27 = _dafny.Set.fromElements(_module.D1.create_DC5(!(p2), p2), _612_v26, _612_v26, _612_v26, _612_v26);
        let _index87 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_611_v25).length));
        let _rhs75 = new BigNumber(((_613_v27).Union(_613_v27)).length);
        let _rhs76 = p2;
        let _rhs77 = (_dafny.Map.Empty.slice().updateUnsafe(p2,(p1)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((p1).length))])).Merge(_607_v21);
        let _rhs78 = _dafny.Seq.Concat(_597_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), function (_614_i0) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new _dafny.CodePoint('o'.codePointAt(0)))).length);
        }));
        let _lhs65 = globalState;
        let _lhs66 = globalState;
        let _lhs67 = _611_v25;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_611_v25).length));
        _lhs65.f3 = _rhs75;
        _lhs66.f9 = _rhs76;
        _lhs67[_lhs68] = _rhs77;
        _597_v11 = _rhs78;
      }
      let _615_v28;
      _615_v28 = new BigNumber(267);
      let _616_v29;
      _616_v29 = _dafny.Seq.of(_615_v28, _615_v28);
      let _617_v30;
      _617_v30 = new _dafny.CodePoint('j'.codePointAt(0));
      let _618_v31;
      _618_v31 = _dafny.Map.Empty.slice().updateUnsafe(_617_v30,_615_v28);
      let _619_v32;
      _619_v32 = _dafny.Set.fromElements(_616_v29);
      if (_module.__default.fm23(_module.__default.safeDivisionInt(_615_v28, _615_v28), _module.__default.fm23(_615_v28, p2, _module.__default.fm24(_615_v28, p2, p2, _615_v28, globalState), _dafny.Set.fromElements(_616_v29, _616_v29), globalState), (_618_v31).Merge((_618_v31).update(_617_v30, _615_v28)), _619_v32, globalState)) {
        let _620_v33;
        _620_v33 = _dafny.Seq.of(p2, _module.__default.fm23(new BigNumber(-497), false, _618_v31, _619_v32, globalState));
        let _621_v34;
        _621_v34 = _module.D2.create_DC7(_620_v33);
        _620_v33 = (_621_v34).dtor_cf9;
        let _622_v35;
        _622_v35 = _dafny.Seq.UnicodeFromString("yjf");
        let _623_v36;
        _623_v36 = _dafny.Map.Empty.slice().updateUnsafe(_615_v28,new BigNumber((_622_v35).length));
        let _624_v37;
        _624_v37 = _module.D1.create_DC4(new BigNumber((_623_v36).length), new BigNumber(112));
        let _source10 = _624_v37;
        if (_source10.is_DC4) {
          let _625___mcc_h0 = (_source10).cf4;
          let _626___mcc_h1 = (_source10).cf5;
          let _627_cf5 = _626___mcc_h1;
          let _628_cf4 = _625___mcc_h0;
          let _629_v38;
          _629_v38 = _module.D1.create_DC3(_627_cf5);
          let _630_v39;
          let _nw96 = Array((new BigNumber(29)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = p2;
          _nw96[(_dafny.ONE).toNumber()] = !(!(p2));
          _nw96[(new BigNumber(2)).toNumber()] = p2;
          _nw96[(new BigNumber(3)).toNumber()] = false;
          _nw96[(new BigNumber(4)).toNumber()] = _module.__default.fm23(_615_v28, true, _618_v31, _dafny.Set.fromElements(_module.__default.fm21(_629_v38, p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-918)), ((_631_v34) => function (_632_i1) {
            return _631_v34;
          })(_621_v34)), globalState)), globalState);
          _nw96[(new BigNumber(5)).toNumber()] = p2;
          _nw96[(new BigNumber(6)).toNumber()] = p2;
          _nw96[(new BigNumber(7)).toNumber()] = p2;
          _nw96[(new BigNumber(8)).toNumber()] = p2;
          _nw96[(new BigNumber(9)).toNumber()] = p2;
          _nw96[(new BigNumber(10)).toNumber()] = false;
          _nw96[(new BigNumber(11)).toNumber()] = p2;
          _nw96[(new BigNumber(12)).toNumber()] = p2;
          _nw96[(new BigNumber(13)).toNumber()] = p2;
          _nw96[(new BigNumber(14)).toNumber()] = !((_620_v33)[_module.__default.safeIndex(_627_cf5, new BigNumber((_620_v33).length))]);
          _nw96[(new BigNumber(15)).toNumber()] = false;
          _nw96[(new BigNumber(16)).toNumber()] = p2;
          _nw96[(new BigNumber(17)).toNumber()] = true;
          _nw96[(new BigNumber(18)).toNumber()] = p2;
          _nw96[(new BigNumber(19)).toNumber()] = true;
          _nw96[(new BigNumber(20)).toNumber()] = p2;
          _nw96[(new BigNumber(21)).toNumber()] = p2;
          _nw96[(new BigNumber(22)).toNumber()] = p2;
          _nw96[(new BigNumber(23)).toNumber()] = (_620_v33)[_module.__default.safeIndex(_628_cf4, new BigNumber((_620_v33).length))];
          _nw96[(new BigNumber(24)).toNumber()] = !(false);
          _nw96[(new BigNumber(25)).toNumber()] = p2;
          _nw96[(new BigNumber(26)).toNumber()] = p2;
          _nw96[(new BigNumber(27)).toNumber()] = p2;
          _nw96[(new BigNumber(28)).toNumber()] = true;
          _630_v39 = _nw96;
          let _633_v40;
          _633_v40 = _dafny.Map.Empty.slice().updateUnsafe((_620_v33)[_module.__default.safeIndex(_628_cf4, new BigNumber((_620_v33).length))],_630_v39);
          let _634_v41;
          let _nw97 = new _module.C0();
          _nw97.__ctor(p2, !(p2));
          _634_v41 = _nw97;
          let _635_v42;
          _635_v42 = _dafny.Seq.of(_634_v41, _634_v41, _634_v41);
          let _index88 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_630_v39).length));
          (_630_v39)[_index88] = _dafny.areEqual(_635_v42, _635_v42);
          let _index89 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_630_v39).length));
          let _rhs79 = _622_v35;
          let _rhs80 = _633_v40;
          let _rhs81 = (((_628_cf4).isLessThan(_628_cf4)) ? ((_634_v41).f23) : (p2));
          let _lhs69 = _630_v39;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_630_v39).length));
          _622_v35 = _rhs79;
          _633_v40 = _rhs80;
          _lhs69[_lhs70] = _rhs81;
          let _636_v43;
          let _nw98 = Array((new BigNumber(28)).toNumber()).fill(_module.D2.Default());
          _636_v43 = _nw98;
          let _index90 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_636_v43).length));
          (_636_v43)[_index90] = _621_v34;
          let _637_v44;
          _637_v44 = _dafny.Map.Empty.slice().updateUnsafe((_630_v39)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_630_v39).length))],new BigNumber(284));
          let _index91 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_636_v43).length));
          let _rhs82 = _module.__default.fm25(_628_cf4, (_dafny.ZERO).minus((_628_cf4).plus(_628_cf4)), (_634_v41).f23, _637_v44, globalState);
          let _rhs83 = _634_v41;
          let _rhs84 = _617_v30;
          let _rhs85 = new BigNumber(948);
          let _rhs86 = _621_v34;
          let _lhs71 = _636_v43;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_636_v43).length));
          _617_v30 = _rhs82;
          _634_v41 = _rhs83;
          _617_v30 = _rhs84;
          _627_cf5 = _rhs85;
          _lhs71[_lhs72] = _rhs86;
          let _638_v45;
          _638_v45 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_634_v41).f23);
          _638_v45 = (_638_v45).update(p1, (_634_v41).f23);
          let _639_v46;
          _639_v46 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_615_v28));
          let _640_v47;
          _640_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Set.fromElements((_630_v39)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_630_v39).length))]));
          let _rhs87 = _639_v46;
          let _rhs88 = _630_v39;
          let _rhs89 = _622_v35;
          let _rhs90 = _640_v47;
          let _rhs91 = new BigNumber(79);
          let _lhs73 = globalState;
          _639_v46 = _rhs87;
          _lhs73.f2 = _rhs88;
          _622_v35 = _rhs89;
          _640_v47 = _rhs90;
          r3 = _rhs91;
        } else if (_source10.is_DC5) {
          let _641___mcc_h2 = (_source10).cf6;
          let _642___mcc_h3 = (_source10).cf7;
          let _643_cf7 = _642___mcc_h3;
          let _644_cf6 = _641___mcc_h2;
          let _645_v48;
          _645_v48 = _module.D4.create_DC14(_643_cf7, false);
          let _646_v49;
          _646_v49 = _dafny.Map.Empty.slice().updateUnsafe(_645_v48,((_644_cf6) ? (_643_cf7) : (true)));
          _646_v49 = (_646_v49).update(_645_v48, (_620_v33)[_module.__default.safeIndex(_615_v28, new BigNumber((_620_v33).length))]);
          let _647_v50;
          let _nw99 = Array((new BigNumber(17)).toNumber()).fill([]);
          _647_v50 = _nw99;
          let _648_v51;
          let _nw100 = Array((new BigNumber(8)).toNumber()).fill(false);
          _648_v51 = _nw100;
          let _index92 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length));
          (_647_v50)[_index92] = _648_v51;
          let _index93 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length));
          (_647_v50)[_index93] = _648_v51;
          let _arr0 = (_647_v50)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length))];
          let _index94 = _module.__default.safeIndex(new BigNumber(855), new BigNumber(((_647_v50)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length))]).length));
          _arr0[_index94] = p2;
          let _arr1 = (_647_v50)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length))];
          let _index95 = _module.__default.safeIndex(new BigNumber(855), new BigNumber(((_647_v50)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_647_v50).length))]).length));
          _arr1[_index95] = p2;
          let _649_v52;
          _649_v52 = _dafny.Seq.of(_622_v35, _dafny.Seq.UnicodeFromString("tbocxqw"), _622_v35);
          let _650_v53;
          _650_v53 = _dafny.Map.Empty.slice().updateUnsafe(_616_v29,(_649_v52)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_615_v28, _615_v28, _615_v28)).length), new BigNumber((_649_v52).length))]);
          (globalState).f9 = !(_650_v53).contains(_616_v29);
        } else if (_source10.is_DC3) {
          let _651___mcc_h4 = (_source10).cf3;
          let _652_cf3 = _651___mcc_h4;
          let _rhs92 = _615_v28;
          let _rhs93 = (_module.__default.safeModuloInt(_652_cf3, new BigNumber(324))).multipliedBy(new BigNumber((_622_v35).length));
          r3 = _rhs92;
          r3 = _rhs93;
          let _653_v54;
          let _init10 = ((_654_p2) => function (_655_i2) {
            return _654_p2;
          })(p2);
          let _nw101 = Array((new BigNumber(24)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw101.length); _i0_10++) {
            _nw101[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _653_v54 = _nw101;
          let _656_v55;
          _656_v55 = _module.D4.create_DC13(_622_v35, new BigNumber(930), p2, _653_v54);
          (globalState).f2 = (_656_v55).dtor_cf25;
          let _index96 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_653_v54).length));
          (_653_v54)[_index96] = p2;
          let _index97 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_653_v54).length));
          (_653_v54)[_index97] = p2;
          _615_v28 = _615_v28;
        } else {
          let _657___mcc_h5 = (_source10).cf8;
          let _658_cf8 = _657___mcc_h5;
          let _659_v56;
          _659_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
          let _660_v57;
          let _nw102 = Array((new BigNumber(27)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = p2;
          _nw102[(_dafny.ONE).toNumber()] = p2;
          _nw102[(new BigNumber(2)).toNumber()] = p2;
          _nw102[(new BigNumber(3)).toNumber()] = p2;
          _nw102[(new BigNumber(4)).toNumber()] = p2;
          _nw102[(new BigNumber(5)).toNumber()] = p2;
          _nw102[(new BigNumber(6)).toNumber()] = true;
          _nw102[(new BigNumber(7)).toNumber()] = p2;
          _nw102[(new BigNumber(8)).toNumber()] = p2;
          _nw102[(new BigNumber(9)).toNumber()] = p2;
          _nw102[(new BigNumber(10)).toNumber()] = (((_659_v56).contains(p2)) ? ((_659_v56).get(p2)) : (p2));
          _nw102[(new BigNumber(11)).toNumber()] = p2;
          _nw102[(new BigNumber(12)).toNumber()] = p2;
          _nw102[(new BigNumber(13)).toNumber()] = p2;
          _nw102[(new BigNumber(14)).toNumber()] = p2;
          _nw102[(new BigNumber(15)).toNumber()] = p2;
          _nw102[(new BigNumber(16)).toNumber()] = p2;
          _nw102[(new BigNumber(17)).toNumber()] = (_620_v33)[_module.__default.safeIndex(_615_v28, new BigNumber((_620_v33).length))];
          _nw102[(new BigNumber(18)).toNumber()] = p2;
          _nw102[(new BigNumber(19)).toNumber()] = p2;
          _nw102[(new BigNumber(20)).toNumber()] = !(false);
          _nw102[(new BigNumber(21)).toNumber()] = p2;
          _nw102[(new BigNumber(22)).toNumber()] = (_620_v33)[_module.__default.safeIndex(_615_v28, new BigNumber((_620_v33).length))];
          _nw102[(new BigNumber(23)).toNumber()] = true;
          _nw102[(new BigNumber(24)).toNumber()] = p2;
          _nw102[(new BigNumber(25)).toNumber()] = p2;
          _nw102[(new BigNumber(26)).toNumber()] = false;
          _660_v57 = _nw102;
          let _661_v58;
          _661_v58 = _module.D4.create_DC13(_622_v35, new BigNumber((_dafny.Seq.update(_622_v35, _module.__default.safeIndex(new BigNumber(747), new BigNumber((_622_v35).length)), _617_v30)).length), !(false), _660_v57);
          let _662_v59;
          _662_v59 = _module.D4.create_DC15(_661_v58);
          let _663_v60;
          _663_v60 = _dafny.Seq.of(_662_v59, _662_v59);
          r1 = (_615_v28).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_663_v60, _663_v60))).cardinality()));
          r2 = !(new BigNumber(434)).isEqualTo(new BigNumber(22));
          let _664_v61;
          _664_v61 = _dafny.MultiSet.fromElements(_615_v28);
          let _665_v62;
          _665_v62 = _module.D4.create_DC13(_622_v35, new BigNumber((_664_v61).cardinality()), true, _660_v57);
          (globalState).f3 = (_665_v62).dtor_cf23;
          let _666_v63;
          _666_v63 = _dafny.Map.Empty.slice().updateUnsafe(_615_v28,p2);
          let _667_v64;
          _667_v64 = _dafny.Seq.of(p1);
          let _668_v65;
          _668_v65 = _dafny.Map.Empty.slice().updateUnsafe(true,_615_v28);
          let _669_v66;
          _669_v66 = _dafny.Map.Empty.slice().updateUnsafe(_666_v63,!(new BigNumber((_667_v64).length)).isEqualTo((((_668_v65).contains(p2)) ? ((_668_v65).get(p2)) : (new BigNumber((_668_v65).length)))));
          _669_v66 = _669_v66;
        }
        let _670_v67;
        _670_v67 = _module.D1.create_DC5(p2, p2);
        _670_v67 = _670_v67;
        (globalState).f3 = _615_v28;
        let _671_v68;
        let _nw103 = new _module.C0();
        _nw103.__ctor(p2, p2);
        _671_v68 = _nw103;
      } else {
        let _672_v69;
        let _nw104 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _672_v69 = _nw104;
        let _673_v70;
        _673_v70 = _dafny.MultiSet.fromElements(_672_v69, _672_v69, _672_v69, _672_v69, _672_v69);
        _673_v70 = _dafny.MultiSet.fromElements(_672_v69);
        let _674_v71;
        _674_v71 = _dafny.MultiSet.fromElements(_615_v28);
        let _675_v72;
        _675_v72 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), ((_676_v30) => function (_677_i3) {
          return _676_v30;
        })(_617_v30)), _module.__default.safeIndex(_module.__default.fm0(p2, _615_v28, _615_v28, _674_v71, globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), ((_678_v30) => function (_679_i3) {
          return _678_v30;
        })(_617_v30))).length)), _617_v30));
        let _rhs94 = ((_675_v72).Merge(_675_v72)).Merge(_675_v72);
        let _rhs95 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), ((_680_v28) => function (_681_i4) {
          return (_module.D1.create_DC4(new BigNumber(-838), _680_v28)).dtor_cf4;
        })(_615_v28));
        _675_v72 = _rhs94;
        _616_v29 = _rhs95;
        let _682_v73;
        let _init11 = ((_683_v30) => function (_684_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)))), _dafny.Seq.of(_dafny.Seq.of(_683_v30), _dafny.Seq.of(_683_v30, _683_v30)));
        })(_617_v30);
        let _nw105 = Array((new BigNumber(7)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw105.length); _i0_11++) {
          _nw105[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _682_v73 = _nw105;
        let _685_v74;
        _685_v74 = _dafny.Seq.of(_617_v30, _617_v30);
        let _686_v75;
        _686_v75 = _dafny.Seq.of(_685_v74, _dafny.Seq.of(_617_v30, _617_v30), _685_v74);
        let _index98 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_682_v73).length));
        (_682_v73)[_index98] = _686_v75;
        let _index99 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_682_v73).length));
        (_682_v73)[_index99] = _dafny.Seq.Concat(_686_v75, _dafny.Seq.Concat(_686_v75, _module.__default.fm26(_615_v28, p2, new BigNumber((_685_v74).length), globalState)));
        let _687_v76;
        let _init12 = function (_688_i6) {
          return _dafny.Seq.UnicodeFromString("tulofkoow");
        };
        let _nw106 = Array((new BigNumber(6)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw106.length); _i0_12++) {
          _nw106[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _687_v76 = _nw106;
        let _index100 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_687_v76).length));
        (_687_v76)[_index100] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), ((_689_v30) => function (_690_i7) {
          return _689_v30;
        })(_617_v30)), _dafny.Seq.update(_685_v74, _module.__default.safeIndex(_615_v28, new BigNumber((_685_v74).length)), _617_v30));
        let _index101 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_687_v76).length));
        (_687_v76)[_index101] = _dafny.Seq.Concat(_module.__default.fm27((_dafny.ZERO).minus(_615_v28), p2, globalState), _685_v74);
        let _691_v77;
        _691_v77 = _dafny.Seq.of(true, p2, p2, p2, p2);
        let _692_v78;
        _692_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_691_v77).length),_615_v28);
        let _693_v79;
        _693_v79 = _dafny.Map.Empty.slice().updateUnsafe(_692_v78,_615_v28);
        (globalState).f3 = new BigNumber((((p2) ? (_693_v79) : (_693_v79))).length);
      }
      let _index102 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length));
      (p1)[_index102] = new BigNumber((_616_v29).length);
      let _index103 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length));
      (p1)[_index103] = _615_v28;
      r3 = (_615_v28).multipliedBy((p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))]);
      let _hi1 = _615_v28;
      for (let _694_i8 = (p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))]; _694_i8.isLessThan(_hi1); _694_i8 = _694_i8.plus(_dafny.ONE)) {
        let _695_v80;
        let _nw107 = new _module.C0();
        _nw107.__ctor(((p2) ? (p2) : (false)), (p2) && (false));
        _695_v80 = _nw107;
        if ((_615_v28).isLessThanOrEqualTo(_694_i8)) {
          r2 = (_695_v80).f24;
          let _696_v81;
          _696_v81 = _dafny.Seq.UnicodeFromString("lpi");
          _696_v81 = _696_v81;
          _696_v81 = _696_v81;
          let _697_v82;
          _697_v82 = _dafny.Map.Empty.slice().updateUnsafe(_695_v80,_module.__default.fm27((_dafny.ZERO).minus(_694_i8), p2, globalState));
          let _698_v83;
          let _nw108 = Array((new BigNumber(14)).toNumber()).fill(false);
          _698_v83 = _nw108;
          let _699_v84;
          _699_v84 = _module.D4.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(103)), ((_700_v30) => function (_701_i9) {
  return _700_v30;
})(_617_v30)), new BigNumber(12), (_695_v80).f23, _698_v83);
          let _702_v85;
          _702_v85 = _dafny.Seq.of(_696_v81, _696_v81);
          let _703_v86;
          let _nw109 = Array((new BigNumber(25)).toNumber());
          _nw109[(_dafny.ZERO).toNumber()] = _696_v81;
          _nw109[(_dafny.ONE).toNumber()] = (((_697_v82).contains(_695_v80)) ? ((_697_v82).get(_695_v80)) : (_696_v81));
          _nw109[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pmsifhi"), _696_v81);
          _nw109[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("sta");
          _nw109[(new BigNumber(4)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_696_v81, _module.__default.fm27(new BigNumber((_696_v81).length), (_695_v80).f24, globalState));
          _nw109[(new BigNumber(6)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(7)).toNumber()] = _module.__default.fm27(_615_v28, (_695_v80).f23, globalState);
          _nw109[(new BigNumber(8)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(9)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(10)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(11)).toNumber()] = (_699_v84).dtor_cf22;
          _nw109[(new BigNumber(12)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(13)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_696_v81, _dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_704_i10) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }));
          _nw109[(new BigNumber(15)).toNumber()] = (_702_v85)[_module.__default.safeIndex(_615_v28, new BigNumber((_702_v85).length))];
          _nw109[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_696_v81, _module.__default.safeIndex(_694_i8, new BigNumber((_696_v81).length)), _617_v30), _module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))], new BigNumber((_dafny.Seq.update(_696_v81, _module.__default.safeIndex(_694_i8, new BigNumber((_696_v81).length)), _617_v30)).length)), _617_v30);
          _nw109[(new BigNumber(17)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(18)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(19)).toNumber()] = (_699_v84).dtor_cf22;
          _nw109[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_696_v81, _module.__default.safeIndex(new BigNumber((_696_v81).length), new BigNumber((_696_v81).length)), _617_v30);
          _nw109[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_696_v81, _696_v81);
          _nw109[(new BigNumber(22)).toNumber()] = _696_v81;
          _nw109[(new BigNumber(23)).toNumber()] = ((p2) ? (_696_v81) : (_696_v81));
          _nw109[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("yhvna");
          _703_v86 = _nw109;
          _703_v86 = _703_v86;
          let _705_v87;
          let _init13 = ((_706_v80) => function (_707_i11) {
            return _dafny.Seq.of((_706_v80).f24, true, !(!((_706_v80).f23)));
          })(_695_v80);
          let _nw110 = Array((new BigNumber(21)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw110.length); _i0_13++) {
            _nw110[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _705_v87 = _nw110;
          let _708_v88;
          _708_v88 = _dafny.Seq.of((_695_v80).f24, false);
          let _index104 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_705_v87).length));
          (_705_v87)[_index104] = _708_v88;
          let _index105 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_705_v87).length));
          (_705_v87)[_index105] = _708_v88;
        } else {
          let _709_v89;
          _709_v89 = _dafny.Seq.of(_617_v30);
          (globalState).f3 = _module.__default.fm0(_module.__default.fm23(_694_i8, (_695_v80).f23, _618_v31, _dafny.Set.fromElements(_616_v29), globalState), (new BigNumber((_709_v89).length)).plus(_615_v28), new BigNumber((_dafny.Set.fromElements((p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))])).length), _dafny.MultiSet.fromElements(_694_i8), globalState);
          (globalState).f9 = false;
          (globalState).f9 = (_615_v28).isLessThan(_615_v28);
          let _710_v90;
          let _nw111 = new _module.C0();
          _nw111.__ctor((_695_v80).f23, (_695_v80).f24);
          _710_v90 = _nw111;
          let _711_v91;
          let _nw112 = Array((new BigNumber(11)).toNumber()).fill(false);
          _711_v91 = _nw112;
          let _712_v92;
          _712_v92 = _dafny.MultiSet.fromElements(_711_v91, _711_v91);
          (globalState).f9 = !((((_712_v92).update(_711_v91, _module.__default.abs(new BigNumber(911)))).Difference(_712_v92)).equals(_712_v92));
        }
        r3 = (p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))];
        let _713_v93;
        _713_v93 = _module.D5.create_DC18((_695_v80).f24);
        let _714_v94;
        _714_v94 = _dafny.Map.Empty.slice().updateUnsafe(_713_v93,_615_v28);
        (globalState).f9 = ((_615_v28).multipliedBy(new BigNumber((_714_v94).length))).isLessThan(_615_v28);
      }
      let _index106 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length));
      (p1)[_index106] = (p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))];
      r0 = _module.__default.fm28(p2, globalState);
      r1 = (p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))];
      let _715_v95;
      _715_v95 = _dafny.Seq.of(!(p2));
      let _716_v96;
      _716_v96 = _dafny.Seq.of(_715_v95, _dafny.Seq.update(_715_v95, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))],(p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))])).length), new BigNumber((_715_v95).length)), p2));
      let _717_v97;
      _717_v97 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_615_v28));
      let _pat_let_tv19 = p2;
      let _718_v98;
      _718_v98 = _dafny.Seq.of(function (_pat_let19_0) {
        return function (_719_dt__update__tmp_h0) {
          return function (_pat_let20_0) {
            return function (_720_dt__update_hcf31_h0) {
              return _module.D5.create_DC17((_719_dt__update__tmp_h0).dtor_cf30, _720_dt__update_hcf31_h0, (_719_dt__update__tmp_h0).dtor_cf32);
            }(_pat_let20_0);
          }(_pat_let_tv19);
        }(_pat_let19_0);
      }(_module.D5.create_DC17(_716_v96, p2, _717_v97)));
      r2 = _module.__default.fm23((p1)[_module.__default.safeIndex(new BigNumber(15), new BigNumber((p1).length))], p2, (_dafny.Map.Empty.slice().updateUnsafe(_617_v30,new BigNumber((_718_v98).length))).Merge(_618_v31), _619_v32, globalState);
      r3 = _615_v28;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f22 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22) {
      let _this = this;
      (_this).f22 = f22;
      return;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _module.D3.Default();
      let _721_v0;
      _721_v0 = true;
      let _722_i0;
      _722_i0 = _dafny.ZERO;
      L2: {
        while (_721_v0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_722_i0)) {
              break L2;
            }
            _722_i0 = (_722_i0).plus(_dafny.ONE);
            (globalState).f9 = _721_v0;
            let _723_v1;
            _723_v1 = new _dafny.CodePoint('y'.codePointAt(0));
            r0 = _723_v1;
            let _724_v2;
            _724_v2 = _module.D0.create_DC1(_721_v0);
            let _source11 = _724_v2;
            if (_source11.is_DC1) {
              let _725___mcc_h0 = (_source11).cf1;
              let _726_cf1 = _725___mcc_h0;
              let _727_v3;
              _727_v3 = new BigNumber(8);
              (globalState).f3 = _727_v3;
              let _728_v4;
              let _init14 = ((_729_v3, _730_cf1, _731_v0) => function (_732_i1) {
                return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), ((_733_v3) => function (_734_i2) {
                  return _module.D1.create_DC4(_733_v3, _733_v3);
                })(_729_v3)), _dafny.Seq.of(_module.D1.create_DC4(new BigNumber((_this.f22).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_729_v3,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_730_cf1,_731_v0)).length))).length))));
              })(_727_v3, _726_cf1, _721_v0);
              let _nw113 = Array((new BigNumber(24)).toNumber());
              for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw113.length); _i0_14++) {
                _nw113[_i0_14] = _init14(new BigNumber(_i0_14));
              }
              _728_v4 = _nw113;
              let _735_v5;
              _735_v5 = _dafny.Map.Empty.slice().updateUnsafe(_726_cf1,_726_cf1);
              let _736_v6;
              _736_v6 = _dafny.Map.Empty.slice().updateUnsafe(_723_v1,new BigNumber((_735_v5).length));
              let _737_v7;
              _737_v7 = _dafny.Seq.UnicodeFromString("pr");
              let _738_v8;
              _738_v8 = _dafny.Seq.of((((_736_v6).contains(_723_v1)) ? ((_736_v6).get(_723_v1)) : (_727_v3)), _727_v3, new BigNumber((_737_v7).length));
              let _739_v9;
              let _nw114 = Array((new BigNumber(17)).toNumber()).fill(false);
              _739_v9 = _nw114;
              let _740_v10;
              _740_v10 = _dafny.Map.Empty.slice().updateUnsafe(_739_v9,_723_v1);
              let _741_v11;
              _741_v11 = _dafny.Seq.of(_721_v0, _726_cf1);
              let _742_v12;
              _742_v12 = _dafny.Seq.of(_741_v11, _741_v11);
              let _743_v13;
              let _nw115 = Array((new BigNumber(8)).toNumber());
              _nw115[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("cudqbqnau")).length);
              _nw115[(_dafny.ONE).toNumber()] = new BigNumber(-98);
              _nw115[(new BigNumber(2)).toNumber()] = (_727_v3).plus(_727_v3);
              _nw115[(new BigNumber(3)).toNumber()] = _727_v3;
              _nw115[(new BigNumber(4)).toNumber()] = new BigNumber((((false) ? (_dafny.Seq.of(_727_v3)) : (_dafny.Seq.update(_738_v8, _module.__default.safeIndex(_727_v3, new BigNumber((_738_v8).length)), new BigNumber((_740_v10).length))))).length);
              _nw115[(new BigNumber(5)).toNumber()] = (_727_v3).minus(new BigNumber((_738_v8).length));
              _nw115[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm0(_726_cf1, _727_v3, new BigNumber(441), _dafny.MultiSet.FromArray(_738_v8), globalState), (_dafny.ZERO).minus(new BigNumber(((_742_v12)[_module.__default.safeIndex(_727_v3, new BigNumber((_742_v12).length))]).length)));
              _nw115[(new BigNumber(7)).toNumber()] = (new BigNumber((_738_v8).length)).plus(new BigNumber(558));
              _743_v13 = _nw115;
              let _744_v14;
              let _nw116 = Array((new BigNumber(13)).toNumber());
              _nw116[(_dafny.ZERO).toNumber()] = _724_v2;
              _nw116[(_dafny.ONE).toNumber()] = _724_v2;
              _nw116[(new BigNumber(2)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(3)).toNumber()] = _module.D0.create_DC1((((_this.f22).contains(_727_v3)) ? ((_this.f22).get(_727_v3)) : (false)));
              _nw116[(new BigNumber(4)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(5)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(6)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(7)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(8)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(9)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(10)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(11)).toNumber()] = _724_v2;
              _nw116[(new BigNumber(12)).toNumber()] = _724_v2;
              _744_v14 = _nw116;
              let _index107 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_744_v14).length));
              (_744_v14)[_index107] = _724_v2;
              let _745_v15;
              _745_v15 = _dafny.Seq.of(_728_v4);
              let _746_v16;
              _746_v16 = _dafny.Map.Empty.slice().updateUnsafe(_727_v3,p0);
              let _747_v17;
              _747_v17 = _dafny.MultiSet.fromElements(new BigNumber((_this.f22).length));
              let _pat_let_tv20 = _726_cf1;
              let _index108 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_744_v14).length));
              let _rhs96 = (_745_v15)[_module.__default.safeIndex((_727_v3).multipliedBy(new BigNumber((_746_v16).length)), new BigNumber((_745_v15).length))];
              let _rhs97 = _module.__default.fm0(false, (_727_v3).plus(new BigNumber(358)), _727_v3, _747_v17, globalState);
              let _rhs98 = _743_v13;
              let _rhs99 = (_module.D4.create_DC12(_739_v9)).dtor_cf21;
              let _rhs100 = function (_pat_let21_0) {
                return function (_748_dt__update__tmp_h0) {
                  return function (_pat_let22_0) {
                    return function (_749_dt__update_hcf1_h0) {
                      return _module.D0.create_DC1(_749_dt__update_hcf1_h0);
                    }(_pat_let22_0);
                  }(_pat_let_tv20);
                }(_pat_let21_0);
              }(_724_v2);
              let _lhs74 = _744_v14;
              let _lhs75 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_744_v14).length));
              _728_v4 = _rhs96;
              _727_v3 = _rhs97;
              _743_v13 = _rhs98;
              _739_v9 = _rhs99;
              _lhs74[_lhs75] = _rhs100;
              let _750_v18;
              let _nw117 = new _module.C1();
              _nw117.__ctor();
              _750_v18 = _nw117;
              let _index109 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((p0).length));
              (p0)[_index109] = _727_v3;
              let _index110 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((p0).length));
              (p0)[_index110] = ((_727_v3).multipliedBy(_727_v3)).multipliedBy((_727_v3).minus(_727_v3));
            } else if (_source11.is_DC0) {
              let _751___mcc_h1 = (_source11).cf0;
              let _752_cf0 = _751___mcc_h1;
              let _753_v19;
              _753_v19 = _module.D0.create_DC0(_721_v0);
              _753_v19 = _753_v19;
              let _754_v20;
              _754_v20 = _dafny.Map.Empty.slice().updateUnsafe((_752_cf0) || (_721_v0),_dafny.Seq.Create(_module.__default.abs(new BigNumber(529)), ((_755_v1) => function (_756_i3) {
                return _755_v1;
              })(_723_v1)));
              let _757_v21;
              _757_v21 = _dafny.Seq.UnicodeFromString("lksdhhcbj");
              _754_v20 = (_754_v20).update(_752_cf0, _757_v21);
              let _758_v23;
              let _out42;
              _out42 = (_this).m9(new BigNumber((function () {
                let _coll38 = new _dafny.Set();
                for (const _compr_38 of _dafny.IntegerRange(new BigNumber(502), new BigNumber(693))) {
                  let _759_v22 = _compr_38;
                  if (((new BigNumber(502)).isLessThanOrEqualTo(_759_v22)) && ((_759_v22).isLessThan(new BigNumber(693)))) {
                    _coll38.add(_module.__default.safeDivisionInt(_759_v22, new BigNumber((_757_v21).length)));
                  }
                }
                return _coll38;
              }()).length), globalState);
              _758_v23 = _out42;
              let _760_v24;
              let _nw118 = Array((new BigNumber(3)).toNumber());
              _nw118[(_dafny.ZERO).toNumber()] = _752_cf0;
              _nw118[(_dafny.ONE).toNumber()] = _752_cf0;
              _nw118[(new BigNumber(2)).toNumber()] = _752_cf0;
              _760_v24 = _nw118;
              let _index111 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_760_v24).length));
              (_760_v24)[_index111] = _721_v0;
              let _index112 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_760_v24).length));
              (_760_v24)[_index112] = _721_v0;
            } else {
              let _761___mcc_h2 = (_source11).cf2;
              let _762_cf2 = _761___mcc_h2;
              let _763_v25;
              _763_v25 = new BigNumber(921);
              r1 = _763_v25;
              let _index113 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((p0).length));
              (p0)[_index113] = new BigNumber(887);
              let _764_v26;
              _764_v26 = _dafny.Seq.UnicodeFromString("wfyhouxd");
              let _index114 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((p0).length));
              (p0)[_index114] = _module.__default.safeDivisionInt(_763_v25, (new BigNumber(-526)).plus(new BigNumber((_764_v26).length)));
              let _765_v27;
              let _nw119 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
              _765_v27 = _nw119;
              _765_v27 = _765_v27;
              let _766_v28;
              let _nw120 = Array((new BigNumber(14)).toNumber()).fill(_module.D4.Default());
              _766_v28 = _nw120;
              let _767_v29;
              _767_v29 = _dafny.Map.Empty.slice().updateUnsafe(_763_v25,_766_v28);
              let _768_v30;
              let _nw121 = Array((new BigNumber(10)).toNumber());
              _nw121[(_dafny.ZERO).toNumber()] = _766_v28;
              _nw121[(_dafny.ONE).toNumber()] = _766_v28;
              _nw121[(new BigNumber(2)).toNumber()] = _766_v28;
              _nw121[(new BigNumber(3)).toNumber()] = _766_v28;
              _nw121[(new BigNumber(4)).toNumber()] = _766_v28;
              _nw121[(new BigNumber(5)).toNumber()] = ((true) ? (_766_v28) : (_766_v28));
              _nw121[(new BigNumber(6)).toNumber()] = _766_v28;
              _nw121[(new BigNumber(7)).toNumber()] = _766_v28;
              _nw121[(new BigNumber(8)).toNumber()] = (((_767_v29).contains((p0)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((p0).length))])) ? ((_767_v29).get((p0)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((p0).length))])) : (_766_v28));
              _nw121[(new BigNumber(9)).toNumber()] = _766_v28;
              _768_v30 = _nw121;
              let _index115 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_768_v30).length));
              (_768_v30)[_index115] = _766_v28;
              let _index116 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_768_v30).length));
              (_768_v30)[_index116] = _766_v28;
            }
            (globalState).f9 = (_721_v0) && (_721_v0);
          }
        }
      }
      let _769_v31;
      _769_v31 = new BigNumber(86);
      let _770_v32;
      _770_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(_769_v31).isEqualTo(new BigNumber(-586)),true);
      let _771_v33;
      _771_v33 = new _dafny.CodePoint('f'.codePointAt(0));
      let _772_v34;
      _772_v34 = _dafny.Map.Empty.slice().updateUnsafe(_771_v33,_769_v31);
      let _773_v35;
      _773_v35 = _dafny.Set.fromElements(_dafny.Seq.of(_769_v31, _769_v31));
      _770_v32 = (_770_v32).update(_721_v0, _module.__default.fm23(_769_v31, true, _772_v34, _773_v35, globalState));
      r0 = _771_v33;
      let _774_v36;
      let _nw122 = new _module.C1();
      _nw122.__ctor();
      _774_v36 = _nw122;
      let _775_i4;
      _775_i4 = _dafny.ZERO;
      L3: {
        while (_721_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_775_i4)) {
              break L3;
            }
            _775_i4 = (_775_i4).plus(_dafny.ONE);
            let _776_v37;
            _776_v37 = _dafny.Set.fromElements(new BigNumber(365), _769_v31, _769_v31);
            let _777_v38;
            _777_v38 = _dafny.Seq.of(_776_v37, _776_v37, _776_v37);
            let _778_v39;
            _778_v39 = _dafny.Map.Empty.slice().updateUnsafe(_769_v31,(_777_v38)[_module.__default.safeIndex(_769_v31, new BigNumber((_777_v38).length))]);
            _778_v39 = (_778_v39).update(_769_v31, _776_v37);
            let _779_v40;
            _779_v40 = _dafny.Seq.UnicodeFromString("ncnudk");
            let _780_v41;
            _780_v41 = _dafny.MultiSet.fromElements(new BigNumber((_779_v40).length), _769_v31, _769_v31, _769_v31, _769_v31);
            let _index117 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((p0).length));
            (p0)[_index117] = new BigNumber((_780_v41).cardinality());
            let _index118 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((p0).length));
            (p0)[_index118] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), ((_781_v31) => function (_782_i5) {
              return _781_v31;
            })(_769_v31))).length)).plus(_module.__default.fm0(_721_v0, _769_v31, _769_v31, _780_v41, globalState));
            let _783_v42;
            let _nw123 = new _module.C0();
            _nw123.__ctor(_721_v0, _721_v0);
            _783_v42 = _nw123;
            let _784_v43;
            _784_v43 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(215), new BigNumber((p0).length))],_769_v31);
            r1 = (new BigNumber((_dafny.Seq.of(new BigNumber((_784_v43).length))).length)).plus((p0)[_module.__default.safeIndex(new BigNumber(215), new BigNumber((p0).length))]);
          }
        }
      }
      let _785_v44;
      _785_v44 = _dafny.MultiSet.fromElements(!(_721_v0));
      let _786_v45;
      _786_v45 = _dafny.Seq.of(_769_v31, _769_v31, _769_v31, _769_v31, new BigNumber(494));
      (globalState).f3 = _module.__default.safeModuloInt(new BigNumber((_785_v44).cardinality()), new BigNumber((((_721_v0) ? (_786_v45) : (_786_v45))).length));
      r0 = _771_v33;
      r1 = _769_v31;
      let _787_v46;
      _787_v46 = _dafny.Seq.of(_772_v34, _772_v34, _772_v34, _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(701)), _772_v34);
      let _788_v47;
      _788_v47 = _dafny.Seq.UnicodeFromString("ymak");
      let _789_v48;
      _789_v48 = _module.D3.create_DC10(_770_v32);
      r2 = ((_module.__default.fm23(_769_v31, _721_v0, (_787_v46)[_module.__default.safeIndex(new BigNumber((_788_v47).length), new BigNumber((_787_v46).length))], _773_v35, globalState)) ? (_789_v48) : (_789_v48));
      return [r0, r1, r2];
    }
    m9(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _790_v0;
      _790_v0 = true;
      let _791_v1;
      _791_v1 = _dafny.MultiSet.fromElements(p0, p0, new BigNumber(917), p0);
      let _792_v2;
      _792_v2 = _dafny.Seq.of(_791_v1);
      let _793_v3;
      _793_v3 = _dafny.Seq.UnicodeFromString("p");
      let _794_v4;
      _794_v4 = _module.D5.create_DC19(_module.D5.create_DC17(_module.__default.fm29(globalState), _790_v0, (_792_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_793_v3).length)), new BigNumber((_792_v2).length))]));
      let _795_v5;
      _795_v5 = _dafny.Seq.of(_794_v4);
      let _796_v6;
      _796_v6 = _dafny.Seq.of(_790_v0);
      let _797_v7;
      _797_v7 = _module.D2.create_DC7(_796_v6);
      let _798_v8;
      let _nw124 = Array((new BigNumber(17)).toNumber());
      _nw124[(_dafny.ZERO).toNumber()] = p0;
      _nw124[(_dafny.ONE).toNumber()] = p0;
      _nw124[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw124[(new BigNumber(3)).toNumber()] = p0;
      _nw124[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_794_v4), _795_v5)).length);
      _nw124[(new BigNumber(5)).toNumber()] = p0;
      _nw124[(new BigNumber(6)).toNumber()] = p0;
      _nw124[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw124[(new BigNumber(8)).toNumber()] = p0;
      _nw124[(new BigNumber(9)).toNumber()] = p0;
      _nw124[(new BigNumber(10)).toNumber()] = (p0).plus(p0);
      _nw124[(new BigNumber(11)).toNumber()] = (new BigNumber(((_797_v7).dtor_cf9).length)).plus(p0);
      _nw124[(new BigNumber(12)).toNumber()] = p0;
      _nw124[(new BigNumber(13)).toNumber()] = p0;
      _nw124[(new BigNumber(14)).toNumber()] = p0;
      _nw124[(new BigNumber(15)).toNumber()] = p0;
      _nw124[(new BigNumber(16)).toNumber()] = p0;
      _798_v8 = _nw124;
      let _index119 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      (_798_v8)[_index119] = _module.__default.safeModuloInt(p0, p0);
      let _index120 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      (_798_v8)[_index120] = new BigNumber(-989);
      let _799_v9;
      _799_v9 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lpyalxhrf"), _793_v3);
      (globalState).f9 = (_796_v6)[_module.__default.safeIndex(new BigNumber(((_799_v9).Union(_799_v9)).cardinality()), new BigNumber((_796_v6).length))];
      let _800_v10;
      let _nw125 = new _module.C1();
      _nw125.__ctor();
      _800_v10 = _nw125;
      let _801_v11;
      let _init15 = function (_802_i0) {
        return (_module.D5.create_DC18(true)).dtor_cf33;
      };
      let _nw126 = Array((new BigNumber(23)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw126.length); _i0_15++) {
        _nw126[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _801_v11 = _nw126;
      let _index121 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
      (_801_v11)[_index121] = _790_v0;
      let _index122 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_801_v11).length));
      (_801_v11)[_index122] = !(_790_v0);
      let _index123 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
      let _index124 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      let _index125 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_801_v11).length));
      let _rhs101 = _790_v0;
      let _rhs102 = !((_module.__default.safeDivisionInt(p0, (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))])).isLessThanOrEqualTo((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]));
      let _rhs103 = _module.__default.safeModuloInt((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], _module.__default.fm0(_790_v0, new BigNumber(177), (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], _791_v1, globalState));
      let _rhs104 = _790_v0;
      let _lhs76 = _801_v11;
      let _lhs77 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
      let _lhs78 = _798_v8;
      let _lhs79 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      let _lhs80 = _801_v11;
      let _lhs81 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_801_v11).length));
      _lhs76[_lhs77] = _rhs101;
      _790_v0 = _rhs102;
      _lhs78[_lhs79] = _rhs103;
      _lhs80[_lhs81] = _rhs104;
      let _index126 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
      let _index127 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      let _rhs105 = (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))];
      let _rhs106 = new BigNumber(174);
      let _lhs82 = _801_v11;
      let _lhs83 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
      let _lhs84 = _798_v8;
      let _lhs85 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
      _lhs82[_lhs83] = _rhs105;
      _lhs84[_lhs85] = _rhs106;
      let _803_v12;
      _803_v12 = new _dafny.CodePoint('q'.codePointAt(0));
      let _804_v13;
      _804_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(_790_v0),(_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]);
      let _805_v14;
      _805_v14 = _dafny.Seq.of(p0, new BigNumber((_804_v13).length), p0, (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]);
      let _806_v15;
      _806_v15 = _dafny.Set.fromElements(_805_v14);
      let _807_v16;
      _807_v16 = _dafny.Set.fromElements(_790_v0, _module.__default.fm23((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], false, _dafny.Map.Empty.slice().updateUnsafe(_803_v12,(_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]), _806_v15, globalState));
      if ((_807_v16).IsProperSubsetOf((_807_v16).Intersect(_807_v16))) {
        if ((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]) {
          let _808_v17;
          _808_v17 = _dafny.Map.Empty.slice().updateUnsafe(_790_v0,p0);
          let _809_v18;
          _809_v18 = _dafny.Map.Empty.slice().updateUnsafe(_793_v3,_808_v17);
          let _810_v19;
          _810_v19 = _module.D3.create_DC11(p0, new BigNumber((_804_v13).length), _790_v0, (_805_v14)[_module.__default.safeIndex(p0, new BigNumber((_805_v14).length))]);
          _809_v18 = (_809_v18).update(((_790_v0) ? (_dafny.Seq.UnicodeFromString("byt")) : (_dafny.Seq.UnicodeFromString("e"))), _module.__default.fm1((_810_v19).dtor_cf19, (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], !(_790_v0), _module.__default.fm22(p0, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,_790_v0), _dafny.Map.Empty.slice().updateUnsafe(p0,(_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))])), _790_v0, !((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]), globalState), globalState));
          let _811_v20;
          _811_v20 = _dafny.MultiSet.fromElements((_808_v17).update((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], p0), _dafny.Map.Empty.slice().updateUnsafe(_790_v0,p0));
          let _index128 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
          (_798_v8)[_index128] = new BigNumber((_811_v20).cardinality());
          (globalState).f3 = p0;
          (globalState).f9 = !(_790_v0);
          (globalState).f3 = _module.__default.fm0(true, new BigNumber(312), p0, _791_v1, globalState);
        } else {
          let _index129 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
          (_801_v11)[_index129] = (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))];
          r0 = (_805_v14)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_805_v14).length))];
          let _812_v21;
          _812_v21 = _module.D1.create_DC5(!_dafny.Seq.contains(_805_v14, p0), false);
          let _813_v22;
          _813_v22 = _dafny.Seq.of(_796_v6);
          let _814_v23;
          _814_v23 = _module.D5.create_DC17(_813_v22, (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _791_v1);
          let _index130 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
          let _index131 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
          let _rhs107 = (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))];
          let _rhs108 = _812_v21;
          let _rhs109 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], p0)).minus((_dafny.ZERO).minus((new BigNumber((_791_v1).cardinality())).multipliedBy((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]))));
          let _rhs110 = true;
          let _rhs111 = !(((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]) || (!((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]))) || ((_814_v23).dtor_cf31);
          let _lhs86 = globalState;
          let _lhs87 = _798_v8;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
          let _lhs89 = globalState;
          let _lhs90 = _801_v11;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
          _lhs86.f9 = _rhs107;
          _812_v21 = _rhs108;
          _lhs87[_lhs88] = _rhs109;
          _lhs89.f9 = _rhs110;
          _lhs90[_lhs91] = _rhs111;
          (globalState).f3 = new BigNumber(31);
          let _index132 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
          (_801_v11)[_index132] = (new BigNumber(882)).isEqualTo((p0).multipliedBy((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]));
        }
        let _index133 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
        (_798_v8)[_index133] = (p0).multipliedBy(((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]).multipliedBy(p0));
        if ((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]) {
          let _815_v24;
          _815_v24 = _dafny.Map.Empty.slice().updateUnsafe((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))],p0);
          _815_v24 = ((_dafny.Map.Empty.slice().updateUnsafe((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))],(_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))])).update((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], (_dafny.ZERO).minus(p0))).Merge(_815_v24);
          (globalState).f9 = (p0).isLessThanOrEqualTo(p0);
          let _816_v25;
          _816_v25 = _module.D4.create_DC14(_790_v0, (new BigNumber(446)).isLessThanOrEqualTo(p0));
          _816_v25 = _816_v25;
          _790_v0 = _790_v0;
          let _index134 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
          (_798_v8)[_index134] = (_dafny.ZERO).minus((((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]) ? ((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]) : (_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), new BigNumber((_793_v3).length)))));
        } else {
          (globalState).f9 = !(!(_790_v0));
          let _817_v26;
          let _init16 = ((_818_v12) => function (_819_i1) {
            return _818_v12;
          })(_803_v12);
          let _nw127 = Array((new BigNumber(16)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw127.length); _i0_16++) {
            _nw127[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _817_v26 = _nw127;
          let _index135 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_817_v26).length));
          (_817_v26)[_index135] = _803_v12;
          let _index136 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_817_v26).length));
          (_817_v26)[_index136] = _803_v12;
          let _820_v27;
          _820_v27 = _dafny.Map.Empty.slice().updateUnsafe((_817_v26)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_817_v26).length))],new BigNumber(118));
          let _821_v28;
          _821_v28 = _module.D1.create_DC5(_790_v0, _module.__default.fm23((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _820_v27, _806_v15, globalState));
          let _822_v29;
          _822_v29 = _module.D1.create_DC6(_821_v28);
          let _823_v30;
          _823_v30 = _module.D1.create_DC6(_821_v28);
          let _824_v31;
          _824_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_823_v30);
          _824_v31 = (_824_v31).update(p0, _823_v30);
          (globalState).f3 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(_module.__default.fm0(_790_v0, new BigNumber((_791_v1).cardinality()), (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], _791_v1, globalState), (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], _790_v0, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23(new BigNumber((_791_v1).cardinality()), _790_v0, _820_v27, _806_v15, globalState),(_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]), globalState),_796_v6)).length);
          (globalState).f3 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-204)), _module.__default.safeDivisionInt((_dafny.ZERO).minus((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]), (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]));
        }
        _790_v0 = true;
        let _825_v32;
        _825_v32 = _module.D2.create_DC9((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], p0, _803_v12);
        let _826_v33;
        _826_v33 = _module.D4.create_DC13(_793_v3, (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _801_v11);
        let _index137 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
        let _index138 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
        let _rhs112 = (_825_v32).dtor_cf14;
        let _rhs113 = (_826_v33).dtor_cf22;
        let _rhs114 = false;
        let _rhs115 = (_module.__default.safeModuloInt(p0, p0)).plus(_module.__default.fm0((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], p0, p0, _791_v1, globalState));
        let _lhs92 = _798_v8;
        let _lhs93 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
        let _lhs94 = _801_v11;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length));
        let _lhs96 = globalState;
        _lhs92[_lhs93] = _rhs112;
        _793_v3 = _rhs113;
        _lhs94[_lhs95] = _rhs114;
        _lhs96.f3 = _rhs115;
      } else {
        (globalState).f3 = new BigNumber(835);
        (globalState).f3 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]), (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))])));
        let _827_v34;
        _827_v34 = _module.D4.create_DC13(_793_v3, new BigNumber(452), _790_v0, _801_v11);
        let _828_v35;
        _828_v35 = _dafny.Seq.of(_827_v34, _827_v34, _827_v34);
        _790_v0 = _dafny.Seq.IsPrefixOf(_828_v35, _828_v35);
        let _829_v36;
        _829_v36 = _module.D2.create_DC9(_790_v0, (_dafny.ZERO).minus(p0), _803_v12);
        let _830_v37;
        _830_v37 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(true)), _dafny.MultiSet.fromElements((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]));
        let _831_v38;
        _831_v38 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? ((_829_v36).dtor_cf15) : (_803_v12)),((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]).minus(new BigNumber((_830_v37).length)));
        let _832_v39;
        _832_v39 = _dafny.MultiSet.fromElements(_798_v8);
        _831_v38 = (_831_v38).update(new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber((_832_v39).cardinality()));
        if (!(((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]).isEqualTo((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]))) {
          (globalState).f13 = _807_v16;
          let _833_v40;
          _833_v40 = _dafny.Map.Empty.slice().updateUnsafe((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))],_803_v12);
          let _834_v43;
          _834_v43 = _833_v40;
          let _835_v44;
          let _nw128 = Array((new BigNumber(28)).toNumber());
          _nw128[(_dafny.ZERO).toNumber()] = _833_v40;
          _nw128[(_dafny.ONE).toNumber()] = (_833_v40).update(p0, new _dafny.CodePoint('o'.codePointAt(0)));
          _nw128[(new BigNumber(2)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(3)).toNumber()] = function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_module.__default.fm30((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], globalState)).Elements) {
              let _836_v41 = _compr_39;
              if ((_module.__default.fm30((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], globalState)).contains(_836_v41)) {
                _coll39.push([_module.__default.safeDivisionInt(_836_v41, new BigNumber((_796_v6).length)),_803_v12]);
              }
            }
            return _coll39;
          }();
          _nw128[(new BigNumber(4)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(5)).toNumber()] = function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(-39), new BigNumber(-511))) {
              let _837_v42 = _compr_40;
              if (((new BigNumber(-39)).isLessThanOrEqualTo(_837_v42)) && ((_837_v42).isLessThan(new BigNumber(-511)))) {
                _coll40.push([_module.__default.safeModuloInt(_837_v42, (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]),new _dafny.CodePoint('v'.codePointAt(0))]);
              }
            }
            return _coll40;
          }();
          _nw128[(new BigNumber(6)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(7)).toNumber()] = _module.__default.fm31(_790_v0, _790_v0, globalState);
          _nw128[(new BigNumber(8)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(9)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(10)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(11)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(12)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(13)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(14)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_803_v12);
          _nw128[(new BigNumber(16)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(17)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(18)).toNumber()] = (_834_v43);
          _nw128[(new BigNumber(19)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(20)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(21)).toNumber()] = _module.__default.fm31(true, _790_v0, globalState);
          _nw128[(new BigNumber(22)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(23)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(24)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))],_803_v12);
          _nw128[(new BigNumber(26)).toNumber()] = _833_v40;
          _nw128[(new BigNumber(27)).toNumber()] = _833_v40;
          _835_v44 = _nw128;
          let _838_v45;
          _838_v45 = _module.D0.create_DC1(_790_v0);
          let _839_v46;
          let _840_v47;
          let _841_v48;
          let _842_v49;
          let _out43;
          let _out44;
          let _out45;
          let _out46;
          let _outcollector15 = (_800_v10).m10(_835_v44, _798_v8, ((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]) && ((_838_v45).dtor_cf1), globalState);
          _out43 = _outcollector15[0];
          _out44 = _outcollector15[1];
          _out45 = _outcollector15[2];
          _out46 = _outcollector15[3];
          _839_v46 = _out43;
          _840_v47 = _out44;
          _841_v48 = _out45;
          _842_v49 = _out46;
          let _843_v50;
          _843_v50 = _dafny.MultiSet.fromElements(_803_v12, _803_v12, _803_v12, _803_v12, _803_v12);
          let _844_v51;
          _844_v51 = _module.D4.create_DC14(!(new BigNumber((_module.__default.fm29(globalState)).length)).isEqualTo(new BigNumber((_843_v50).cardinality())), (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]);
          let _845_v52;
          _845_v52 = _dafny.MultiSet.fromElements(_841_v48, (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _790_v0, (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _841_v48);
          let _pat_let_tv21 = _845_v52;
          let _pat_let_tv22 = _845_v52;
          _844_v51 = function (_pat_let23_0) {
            return function (_846_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_847_dt__update_hcf26_h0) {
                  return _module.D4.create_DC14(_847_dt__update_hcf26_h0, (_846_dt__update__tmp_h0).dtor_cf27);
                }(_pat_let24_0);
              }((_pat_let_tv21).IsProperSubsetOf(_pat_let_tv22));
            }(_pat_let23_0);
          }(_module.D4.create_DC14((((_this.f22).contains(_842_v49)) ? ((_this.f22).get(_842_v49)) : (_790_v0)), _790_v0));
          (globalState).f3 = _840_v47;
          let _848_v53;
          let _nw129 = new _module.C0();
          _nw129.__ctor(_790_v0, _790_v0);
          _848_v53 = _nw129;
        } else {
          let _index139 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length));
          (_798_v8)[_index139] = new BigNumber(288);
          (globalState).f9 = (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))];
          (globalState).f9 = _790_v0;
          let _849_v54;
          _849_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-811)), ((_850_v12) => function (_851_i2) {
            return _850_v12;
          })(_803_v12)));
          _793_v3 = _dafny.Seq.Concat(_793_v3, (((_849_v54).contains((_dafny.ZERO).minus((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]))) ? ((_849_v54).get((_dafny.ZERO).minus((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]))) : (_dafny.Seq.UnicodeFromString("tecdvlbj"))));
          let _852_v55;
          _852_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(327),_796_v6);
          let _853_v56;
          _853_v56 = _dafny.MultiSet.fromElements(_790_v0, _790_v0, _790_v0);
          let _rhs116 = new BigNumber((_dafny.Set.fromElements((new BigNumber((_852_v55).length)).isLessThanOrEqualTo(new BigNumber(970)), (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], _790_v0, (_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))])).length);
          let _rhs117 = (((((_804_v13).contains((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))])) ? ((_804_v13).get((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))])) : (false))) ? (new BigNumber((_799_v9).cardinality())) : ((((_853_v56).contains(_790_v0)) ? ((_853_v56).get(_790_v0)) : (_module.__default.fm0((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], p0, p0, (_791_v1).update(new BigNumber(554), _module.__default.abs(new BigNumber((_805_v14).length))), globalState)))));
          r0 = _rhs116;
          r0 = _rhs117;
        }
      }
      let _854_v58;
      _854_v58 = _dafny.Set.fromElements(_803_v12);
      r0 = _module.__default.fm0(!(_790_v0), _module.__default.fm0((_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))], (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], (_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))], _dafny.MultiSet.fromElements((_798_v8)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_798_v8).length))]), globalState), new BigNumber((function () {
        let _coll41 = new _dafny.Map();
        for (const _compr_41 of (_854_v58).Elements) {
          let _855_v57 = _compr_41;
          if ((_854_v58).contains(_855_v57)) {
            _coll41.push([_855_v57,(_801_v11)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_801_v11).length))]]);
          }
        }
        return _coll41;
      }()).length), ((false) ? (_791_v1) : (_791_v1)), globalState);
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f14 = false;
      this._f15 = _dafny.MultiSet.Empty;
      this._f20 = _dafny.ZERO;
      this._f21 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f20, f21, f14, f15) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      if (_this.f14) {
        return _this.f14;
      } else {
        return !(new BigNumber(-49)).isEqualTo(new BigNumber(((_this).f21).length));
      }
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dowcmd"), _dafny.Seq.Concat(_dafny.Seq.update((_this).f21, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), function (_856_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length), new BigNumber(((_this).f21).length)), new _dafny.CodePoint('y'.codePointAt(0))), (_this).f21)), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dowcmd"), _dafny.Seq.Concat(_dafny.Seq.update((_this).f21, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), function (_857_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length), new BigNumber(((_this).f21).length)), new _dafny.CodePoint('y'.codePointAt(0))), (_this).f21))).length)), new _dafny.CodePoint('a'.codePointAt(0)));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _858_v0;
      _858_v0 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14));
      let _859_v1;
      _859_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      let _860_v2;
      _860_v2 = _dafny.Map.Empty.slice().updateUnsafe((((_859_v1).contains(!(_this.f14))) ? ((_859_v1).get(!(_this.f14))) : (_this.f14)),_this.f14);
      let _861_v3;
      _861_v3 = _dafny.MultiSet.fromElements((_this).f20);
      _858_v0 = (_858_v0).Difference((_dafny.MultiSet.fromElements((_860_v2).update(false, _this.f14), _860_v2)).update(_859_v1, _module.__default.abs(_module.__default.fm0(_this.f14, (_this).f20, (_this).f20, _861_v3, globalState))));
      let _862_v4;
      _862_v4 = _module.D2.create_DC7(_dafny.Seq.of(_this.f14));
      let _863_v5;
      _863_v5 = _dafny.MultiSet.fromElements(_862_v4);
      _863_v5 = _863_v5;
      (globalState).f3 = _module.__default.safeDivisionInt(new BigNumber(713), ((_this).f20).multipliedBy(new BigNumber(115)));
      let _864_v6;
      _864_v6 = _dafny.Seq.of(_this.f14, _this.f14);
      let _hi2 = new BigNumber((_864_v6).length);
      for (let _865_i0 = (_this).f20; _865_i0.isLessThan(_hi2); _865_i0 = _865_i0.plus(_dafny.ONE)) {
        let _866_v7;
        let _nw130 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _866_v7 = _nw130;
        let _index140 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        (_866_v7)[_index140] = _865_i0;
        let _867_v8;
        let _nw131 = Array((new BigNumber(11)).toNumber()).fill(false);
        _867_v8 = _nw131;
        let _868_v9;
        _868_v9 = _module.D1.create_DC4(_865_i0, new BigNumber(((_this).fm3((_this).f20, new BigNumber((_dafny.Seq.UnicodeFromString("fo")).length), _this.f14, (_this).f20, globalState)).length));
        let _index141 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        let _rhs118 = _867_v8;
        let _rhs119 = !(_this.f14);
        let _rhs120 = ((((_this).f15).contains((_this).fm2(_this.f14, _this.f14, globalState))) ? (((_this).f15).get((_this).fm2(_this.f14, _this.f14, globalState))) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_869_i1) {
          return (_this).f21;
        })).length)));
        let _rhs121 = (_868_v9).dtor_cf4;
        let _lhs97 = globalState;
        let _lhs98 = globalState;
        let _lhs99 = _866_v7;
        let _lhs100 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        let _lhs101 = globalState;
        _lhs97.f2 = _rhs118;
        _lhs98.f9 = _rhs119;
        _lhs99[_lhs100] = _rhs120;
        _lhs101.f3 = _rhs121;
        let _870_v10;
        let _nw132 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _870_v10 = _nw132;
        let _index142 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_870_v10).length));
        (_870_v10)[_index142] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-993)), function (_871_i2) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        });
        let _872_v11;
        let _init17 = ((_873_v9, _874_v7) => function (_875_i3) {
          return function (_pat_let25_0) {
            return function (_876_dt__update__tmp_h0) {
              return function (_pat_let26_0) {
                return function (_877_dt__update_hcf4_h0) {
                  return _module.D1.create_DC4(_877_dt__update_hcf4_h0, (_876_dt__update__tmp_h0).dtor_cf5);
                }(_pat_let26_0);
              }((_dafny.ZERO).minus((_874_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_874_v7).length))]));
            }(_pat_let25_0);
          }(_873_v9);
        })(_868_v9, _866_v7);
        let _nw133 = Array((new BigNumber(8)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw133.length); _i0_17++) {
          _nw133[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _872_v11 = _nw133;
        let _index143 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_872_v11).length));
        (_872_v11)[_index143] = _868_v9;
        let _index144 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        let _index145 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_870_v10).length));
        let _index146 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_872_v11).length));
        let _rhs122 = _865_i0;
        let _rhs123 = (_this).f21;
        let _rhs124 = (((true) ? (false) : (_this.f14))) || (_this.f14);
        let _rhs125 = _865_i0;
        let _rhs126 = _868_v9;
        let _lhs102 = _866_v7;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        let _lhs104 = _870_v10;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_870_v10).length));
        let _lhs106 = globalState;
        let _lhs107 = _872_v11;
        let _lhs108 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_872_v11).length));
        _lhs102[_lhs103] = _rhs122;
        _lhs104[_lhs105] = _rhs123;
        r1 = _rhs124;
        _lhs106.f3 = _rhs125;
        _lhs107[_lhs108] = _rhs126;
        let _index147 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
        (_866_v7)[_index147] = new BigNumber(578);
        if (true) {
          let _878_v12;
          _878_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f14);
          _878_v12 = (_dafny.Map.Empty.slice().updateUnsafe(_865_i0,true)).Merge((_878_v12).Merge(_878_v12));
          let _879_v13;
          _879_v13 = _module.D0.create_DC0(!(_this.f14));
          let _880_v14;
          let _nw134 = Array((new BigNumber(13)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = _879_v13;
          _nw134[(_dafny.ONE).toNumber()] = _879_v13;
          _nw134[(new BigNumber(2)).toNumber()] = _module.__default.fm20(new BigNumber(-575), (_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))], (_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))], globalState);
          _nw134[(new BigNumber(3)).toNumber()] = ((_this.f14) ? (_879_v13) : (_879_v13));
          _nw134[(new BigNumber(4)).toNumber()] = _module.__default.fm20(new BigNumber(600), (_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))], (_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))], globalState);
          _nw134[(new BigNumber(5)).toNumber()] = _879_v13;
          _nw134[(new BigNumber(6)).toNumber()] = ((_this.f14) ? (_879_v13) : (_879_v13));
          _nw134[(new BigNumber(7)).toNumber()] = _879_v13;
          _nw134[(new BigNumber(8)).toNumber()] = _879_v13;
          _nw134[(new BigNumber(9)).toNumber()] = ((_this.f14) ? (_879_v13) : (_879_v13));
          _nw134[(new BigNumber(10)).toNumber()] = _879_v13;
          _nw134[(new BigNumber(11)).toNumber()] = _879_v13;
          _nw134[(new BigNumber(12)).toNumber()] = _879_v13;
          _880_v14 = _nw134;
          let _index148 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_880_v14).length));
          (_880_v14)[_index148] = _879_v13;
          let _index149 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_880_v14).length));
          (_880_v14)[_index149] = _879_v13;
          let _881_v15;
          _881_v15 = new _dafny.CodePoint('e'.codePointAt(0));
          let _882_v16;
          _882_v16 = _dafny.MultiSet.fromElements(_881_v15, new _dafny.CodePoint('p'.codePointAt(0)));
          let _883_v17;
          _883_v17 = _dafny.Seq.of((_this).f20);
          let _rhs127 = (_this).f20;
          let _rhs128 = _882_v16;
          let _rhs129 = _865_i0;
          let _rhs130 = (_this).f20;
          let _rhs131 = (_883_v17)[_module.__default.safeIndex((_this).f20, new BigNumber((_883_v17).length))];
          let _lhs109 = globalState;
          let _lhs110 = globalState;
          let _lhs111 = globalState;
          let _lhs112 = globalState;
          _lhs109.f3 = _rhs127;
          _882_v16 = _rhs128;
          _lhs110.f3 = _rhs129;
          _lhs111.f3 = _rhs130;
          _lhs112.f3 = _rhs131;
          let _884_v18;
          _884_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_dafny.MultiSet.fromElements(_this.f14, _this.f14, _this.f14, _this.f14, false));
          let _index150 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_867_v8).length));
          (_867_v8)[_index150] = !(new BigNumber(((((_884_v18).contains(!(_this.f14))) ? ((_884_v18).get(!(_this.f14))) : ((_this).f15))).cardinality())).isEqualTo(_865_i0);
          let _index151 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_867_v8).length));
          (_867_v8)[_index151] = true;
          (globalState).f3 = ((_this).f20).minus((_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))]);
        } else {
          let _885_v19;
          let _nw135 = new _module.C2();
          _nw135.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f14));
          _885_v19 = _nw135;
          let _886_v20;
          _886_v20 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f14) ? (_this.f14) : (_this.f14)),(_this).f20);
          _886_v20 = (_886_v20).update(_this.f14, ((_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))]).multipliedBy(_865_i0));
          let _index152 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length));
          (_866_v7)[_index152] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(657), (_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))]), new BigNumber(631));
          let _887_v21;
          _887_v21 = _dafny.Map.Empty.slice().updateUnsafe((_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))],_865_i0);
          let _888_v22;
          _888_v22 = _dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21);
          let _889_v23;
          _889_v23 = _dafny.Seq.of((_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))]);
          let _890_v24;
          let _nw136 = Array((new BigNumber(25)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = _888_v22;
          _nw136[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21)).Merge(_888_v22);
          _nw136[(new BigNumber(2)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21);
          _nw136[(new BigNumber(4)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21);
          _nw136[(new BigNumber(6)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(7)).toNumber()] = (_888_v22).Merge(_888_v22);
          _nw136[(new BigNumber(8)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(9)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(10)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(11)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(12)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(13)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(14)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(15)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(16)).toNumber()] = _888_v22;
          _nw136[(new BigNumber(17)).toNumber()] = (_888_v22).Merge(_888_v22);
          _nw136[(new BigNumber(18)).toNumber()] = (_888_v22).update(_866_v7, _887_v21);
          _nw136[(new BigNumber(19)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21)).Merge(_888_v22);
          _nw136[(new BigNumber(20)).toNumber()] = (_888_v22).Merge(_888_v22);
          _nw136[(new BigNumber(21)).toNumber()] = (_888_v22).Merge(_888_v22);
          _nw136[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_866_v7,_887_v21);
          _nw136[(new BigNumber(23)).toNumber()] = (_888_v22).Merge(_888_v22);
          _nw136[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_866_v7,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(135),_module.__default.fm0(_this.f14, _865_i0, (_dafny.ZERO).minus((_889_v23)[_module.__default.safeIndex((_866_v7)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_866_v7).length))], new BigNumber((_889_v23).length))]), _dafny.MultiSet.FromArray(_889_v23), globalState)));
          _890_v24 = _nw136;
          let _index153 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_890_v24).length));
          (_890_v24)[_index153] = (_888_v22).Merge(_888_v22);
          let _index154 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_890_v24).length));
          (_890_v24)[_index154] = (_888_v22).update(_866_v7, _887_v21);
          (_this).f14 = (((_864_v6)[_module.__default.safeIndex(_865_i0, new BigNumber((_864_v6).length))]) ? (_this.f14) : (_this.f14));
        }
      }
      if (!(_this.f14)) {
        let _891_v25;
        _891_v25 = _dafny.Seq.of(new BigNumber(301), ((_this).f20).multipliedBy((_dafny.ZERO).minus((_this).f20)), (_this).f20);
        _891_v25 = _891_v25;
        let _892_v26;
        _892_v26 = _module.D3.create_DC11((_this).f20, ((_this).f20).multipliedBy((_this).f20), _this.f14, (new BigNumber(326)).plus(new BigNumber((_module.__default.fm32((_this).f21, _861_v3, globalState)).length)));
        _892_v26 = _892_v26;
        let _893_v27;
        let _nw137 = Array((new BigNumber(8)).toNumber());
        _nw137[(_dafny.ZERO).toNumber()] = new BigNumber(898);
        _nw137[(_dafny.ONE).toNumber()] = (_this).f20;
        _nw137[(new BigNumber(2)).toNumber()] = (_this).f20;
        _nw137[(new BigNumber(3)).toNumber()] = (_this).f20;
        _nw137[(new BigNumber(4)).toNumber()] = (_this).f20;
        _nw137[(new BigNumber(5)).toNumber()] = (_this).f20;
        _nw137[(new BigNumber(6)).toNumber()] = new BigNumber((_864_v6).length);
        _nw137[(new BigNumber(7)).toNumber()] = (_this).f20;
        _893_v27 = _nw137;
        let _894_v28;
        _894_v28 = _dafny.Seq.of(_893_v27);
        _894_v28 = _dafny.Seq.Concat(_894_v28, _894_v28);
        r3 = (((_dafny.ZERO).minus((_this).f20)).isLessThan((_this).f20)) && (_this.f14);
        (globalState).f9 = _this.f14;
      } else {
        let _895_v29;
        let _nw138 = Array((new BigNumber(29)).toNumber()).fill([]);
        _895_v29 = _nw138;
        let _896_v30;
        _896_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wnwodjwai"), (_this).f21),_895_v29);
        _896_v30 = (_896_v30).update((_this).f21, _895_v29);
        let _897_v31;
        let _nw139 = new _module.C0();
        _nw139.__ctor(_this.f14, _this.f14);
        _897_v31 = _nw139;
        let _898_v32;
        let _nw140 = Array((new BigNumber(24)).toNumber()).fill(false);
        _898_v32 = _nw140;
        let _index155 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_898_v32).length));
        (_898_v32)[_index155] = _this.f14;
        let _index156 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_898_v32).length));
        (_898_v32)[_index156] = (_897_v31).f23;
        let _899_v33;
        let _nw141 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _899_v33 = _nw141;
        let _900_v34;
        _900_v34 = new _dafny.CodePoint('p'.codePointAt(0));
        let _901_v35;
        _901_v35 = _dafny.Map.Empty.slice().updateUnsafe(_899_v33,_dafny.Seq.update((_this).f21, _module.__default.safeIndex((_this).f20, new BigNumber(((_this).f21).length)), _900_v34));
        _901_v35 = _901_v35;
        (globalState).f3 = new BigNumber(255);
      }
      (globalState).f3 = (_this).f20;
      r0 = (((_861_v3).contains((_this).f20)) ? ((_861_v3).get((_this).f20)) : ((_this).f20));
      r1 = _this.f14;
      let _902_v36;
      _902_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f14);
      let _903_v37;
      _903_v37 = _dafny.Map.Empty.slice().updateUnsafe(_902_v36,(_this).f20);
      r2 = (_module.D7.create_DC21(_903_v37)).dtor_cf36;
      r3 = true;
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      (_this).f14 = _this.f14;
      let _904_v0;
      let _nw142 = Array((new BigNumber(14)).toNumber()).fill(false);
      _904_v0 = _nw142;
      let _index157 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length));
      (_904_v0)[_index157] = _this.f14;
      let _index158 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length));
      (_904_v0)[_index158] = _this.f14;
      let _905_v1;
      let _nw143 = new _module.C0();
      _nw143.__ctor(_this.f14, (_this.f14) === (_this.f14));
      _905_v1 = _nw143;
      let _906_v2;
      let _nw144 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _906_v2 = _nw144;
      let _index159 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_906_v2).length));
      (_906_v2)[_index159] = new BigNumber((_module.__default.fm27(new BigNumber((p0).length), (_905_v1).f24, globalState)).length);
      let _index160 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_906_v2).length));
      (_906_v2)[_index160] = _module.__default.safeDivisionInt(p1, _module.__default.fm0((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))], (_this).f20, p1, _module.__default.fm30(false, globalState), globalState));
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_906_v2).length))) {
        let _907_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_907_i0)) && ((_907_i0).isLessThan(new BigNumber((_906_v2).length))))) {
          (_906_v2)[(_907_i0)] = (_907_i0).plus((_module.D3.create_DC11(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_dafny.ZERO).minus((_this).f20))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(p1,_module.D2.create_DC9(!((_905_v1).f24), p1, new _dafny.CodePoint('w'.codePointAt(0)))))).length), (_905_v1).f24, (_this).f20)).dtor_cf20);
        }
      }
      if (false) {
        r2 = (((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))]) ? ((new BigNumber(((_this).f21).length)).isLessThanOrEqualTo(p1)) : (_this.f14));
        let _908_v3;
        let _nw145 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _908_v3 = _nw145;
        let _index161 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_908_v3).length));
        (_908_v3)[_index161] = (_this).f21;
        let _index162 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_908_v3).length));
        (_908_v3)[_index162] = (_this).f21;
        let _909_v4;
        _909_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_905_v1).f23);
        let _910_v5;
        _910_v5 = _dafny.Seq.of((_905_v1).f23);
        (globalState).f6 = _dafny.Set.fromElements(true, (((_905_v1).f23) ? ((((_909_v4).contains((_908_v3)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_908_v3).length))])) ? ((_909_v4).get((_908_v3)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_908_v3).length))])) : ((_910_v5)[_module.__default.safeIndex(p1, new BigNumber((_910_v5).length))]))) : ((_905_v1).f24)), !(!(false)));
        (_this).f14 = !((_this).fm2(_this.f14, !(((_905_v1).f23) === ((_905_v1).f24)), globalState));
        let _911_v6;
        _911_v6 = new _dafny.CodePoint('f'.codePointAt(0));
        let _index163 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length));
        let _rhs132 = ((_905_v1).f24) && ((_905_v1).f24);
        let _rhs133 = _911_v6;
        let _rhs134 = (_905_v1).f24;
        let _lhs113 = _this;
        let _lhs114 = _904_v0;
        let _lhs115 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length));
        _lhs113.f14 = _rhs132;
        _911_v6 = _rhs133;
        _lhs114[_lhs115] = _rhs134;
      } else {
        let _912_v7;
        let _nw146 = Array((new BigNumber(20)).toNumber()).fill([]);
        _912_v7 = _nw146;
        let _index164 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_912_v7).length));
        (_912_v7)[_index164] = _906_v2;
        let _index165 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_912_v7).length));
        let _nw147 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        (_912_v7)[_index165] = _nw147;
        (globalState).f9 = (_905_v1).f24;
        let _913_v8;
        _913_v8 = _dafny.Map.Empty.slice().updateUnsafe((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))],!((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))]));
        let _914_v9;
        _914_v9 = _module.D3.create_DC10((_913_v8).update((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))], (_905_v1).f23));
        let _915_v10;
        _915_v10 = _dafny.Map.Empty.slice().updateUnsafe(_914_v9,((_906_v2)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_906_v2).length))]).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f14, (_905_v1).f24)).length))));
        _915_v10 = (_915_v10).update(_914_v9, (_905_v1).f23);
        let _916_v11;
        _916_v11 = _dafny.Set.fromElements(p1);
        _916_v11 = (_916_v11).Difference((((_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))]) ? (_916_v11) : (_916_v11)));
        let _index166 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length));
        (_904_v0)[_index166] = (_905_v1).f23;
      }
      let _917_v12;
      _917_v12 = _dafny.Seq.of(_this.f14, false);
      let _918_v13;
      _918_v13 = new _dafny.CodePoint('h'.codePointAt(0));
      let _919_v14;
      _919_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_917_v12).length),_918_v13);
      r0 = (new BigNumber((((_this).f15).Difference((_this).f15)).cardinality())).minus(new BigNumber((_919_v14).length));
      let _920_v15;
      _920_v15 = _dafny.MultiSet.fromElements((_906_v2)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_906_v2).length))], p1);
      r1 = ((new BigNumber((_dafny.Set.fromElements(p1)).length)).multipliedBy(new BigNumber((_920_v15).cardinality()))).minus(_module.__default.safeDivisionInt((_this).f20, new BigNumber(546)));
      r2 = !(!(!(false)) || ((new BigNumber((_module.__default.fm27((_dafny.ZERO).minus(new BigNumber((_917_v12).length)), (_904_v0)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_904_v0).length))], globalState)).length)).isLessThan(p1)));
      return [r0, r1, r2];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f14 = false;
      this._f15 = _dafny.MultiSet.Empty;
      this._f19 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f19, f14, f15) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return (_this).f19;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-13)), function (_921_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("xlscb")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fu"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), function (_922_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })));
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((((_this).f15).Intersect((_this).f15)).cardinality());
    };
    fm18(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f19))).dtor_cf16).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f19,_this.f14));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      if (!((_this).f19)) {
        let _923_v0;
        let _init18 = function (_924_i0) {
          return false;
        };
        let _nw148 = Array((new BigNumber(10)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw148.length); _i0_18++) {
          _nw148[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _923_v0 = _nw148;
        let _index167 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_923_v0).length));
        (_923_v0)[_index167] = !((_this).f19) || ((_this).f19);
        let _925_v1;
        _925_v1 = new BigNumber(-689);
        let _926_v2;
        _926_v2 = _dafny.Seq.of(_this.f14);
        let _927_v3;
        _927_v3 = new _dafny.CodePoint('v'.codePointAt(0));
        let _928_v4;
        _928_v4 = _dafny.Set.fromElements(_927_v3);
        let _929_v5;
        _929_v5 = _dafny.Map.Empty.slice().updateUnsafe(_927_v3,(_this).f19);
        let _index168 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_923_v0).length));
        (_923_v0)[_index168] = ((_module.__default.fm19(_925_v1, (_926_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(_925_v1), new BigNumber((_926_v2).length))], globalState)).Difference(_928_v4)).IsDisjointFrom(function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of (_929_v5).Keys.Elements) {
            let _930_v6 = _compr_42;
            if ((_929_v5).contains(_930_v6)) {
              _coll42.add(_930_v6);
            }
          }
          return _coll42;
        }());
        (globalState).f9 = _this.f14;
        let _931_v7;
        _931_v7 = _dafny.Map.Empty.slice().updateUnsafe((_923_v0)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_923_v0).length))],_this.f14);
        let _932_v8;
        _932_v8 = _dafny.Seq.of(_931_v7);
        (globalState).f9 = _dafny.Seq.contains(_932_v8, _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14));
        let _933_v9;
        _933_v9 = _dafny.Seq.of((_this).f15);
        let _934_v10;
        let _init19 = ((_935_v1) => function (_936_i1) {
          return (_936_i1).multipliedBy(_935_v1);
        })(_925_v1);
        let _nw149 = Array((new BigNumber(13)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw149.length); _i0_19++) {
          _nw149[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _934_v10 = _nw149;
        let _937_v11;
        _937_v11 = _dafny.Set.fromElements(_934_v10, _934_v10, _934_v10);
        let _938_v12;
        _938_v12 = _dafny.MultiSet.fromElements(_925_v1);
        (_this).m7(_933_v9, (_dafny.Set.fromElements(_934_v10)).IsDisjointFrom(_937_v11), (((_this).f19) ? (new BigNumber(772)) : (_925_v1)), (_module.__default.fm0(_this.f14, new BigNumber(-955), _925_v1, _938_v12, globalState)).isLessThanOrEqualTo(_925_v1), globalState);
        let _939_v13;
        _939_v13 = _dafny.Map.Empty.slice().updateUnsafe(_925_v1,_927_v3);
        let _940_v14;
        _940_v14 = _module.D3.create_DC11(_925_v1, _925_v1, (_this).f19, new BigNumber((_939_v13).length));
        (globalState).f3 = _module.__default.fm0((new BigNumber((((_this).f15).update((_this).f19, _module.__default.abs(_925_v1))).cardinality())).isLessThanOrEqualTo(_925_v1), _925_v1, (_940_v14).dtor_cf20, _938_v12, globalState);
      } else {
        let _941_v15;
        _941_v15 = _module.D1.create_DC5(_this.f14, _this.f14);
        let _942_v16;
        let _nw150 = Array((new BigNumber(18)).toNumber());
        _nw150[(_dafny.ZERO).toNumber()] = _941_v15;
        _nw150[(_dafny.ONE).toNumber()] = _941_v15;
        _nw150[(new BigNumber(2)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(3)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(4)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(5)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(6)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(7)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(8)).toNumber()] = _module.D1.create_DC5(false, true);
        _nw150[(new BigNumber(9)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(10)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(11)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(12)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(13)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(14)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(15)).toNumber()] = _941_v15;
        _nw150[(new BigNumber(16)).toNumber()] = _module.D1.create_DC5(_this.f14, _this.f14);
        _nw150[(new BigNumber(17)).toNumber()] = _941_v15;
        _942_v16 = _nw150;
        let _943_v17;
        _943_v17 = _dafny.Seq.of(_942_v16, _942_v16, _942_v16);
        let _944_v18;
        _944_v18 = new BigNumber(-618);
        let _945_v19;
        let _nw151 = Array((new BigNumber(12)).toNumber());
        _nw151[(_dafny.ZERO).toNumber()] = _943_v17;
        _nw151[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_943_v17, _943_v17);
        _nw151[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_943_v17, _943_v17), _module.__default.safeIndex(_944_v18, new BigNumber((_dafny.Seq.Concat(_943_v17, _943_v17)).length)), _942_v16);
        _nw151[(new BigNumber(3)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(4)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(5)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(6)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_942_v16, _942_v16);
        _nw151[(new BigNumber(8)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_943_v17, _943_v17);
        _nw151[(new BigNumber(10)).toNumber()] = _943_v17;
        _nw151[(new BigNumber(11)).toNumber()] = _943_v17;
        _945_v19 = _nw151;
        let _index169 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_945_v19).length));
        (_945_v19)[_index169] = _943_v17;
        let _index170 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_945_v19).length));
        (_945_v19)[_index170] = _dafny.Seq.Concat(_dafny.Seq.Concat(_943_v17, _943_v17), _dafny.Seq.Concat(_943_v17, _943_v17));
        (globalState).f9 = (_944_v18).isLessThanOrEqualTo((_dafny.ZERO).minus(_944_v18));
        let _946_v20;
        let _nw152 = Array((new BigNumber(26)).toNumber()).fill(false);
        _946_v20 = _nw152;
        let _947_v21;
        let _nw153 = Array((new BigNumber(8)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = _946_v20;
        _nw153[(_dafny.ONE).toNumber()] = _946_v20;
        _nw153[(new BigNumber(2)).toNumber()] = _946_v20;
        _nw153[(new BigNumber(3)).toNumber()] = _946_v20;
        _nw153[(new BigNumber(4)).toNumber()] = _946_v20;
        _nw153[(new BigNumber(5)).toNumber()] = _946_v20;
        _nw153[(new BigNumber(6)).toNumber()] = ((true) ? (_946_v20) : (_946_v20));
        _nw153[(new BigNumber(7)).toNumber()] = _946_v20;
        _947_v21 = _nw153;
        let _index171 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_947_v21).length));
        (_947_v21)[_index171] = _946_v20;
        let _index172 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_947_v21).length));
        let _nw154 = Array((new BigNumber(7)).toNumber()).fill(false);
        (_947_v21)[_index172] = _nw154;
        let _948_v22;
        _948_v22 = _dafny.Seq.UnicodeFromString("gkgn");
        let _949_v23;
        let _nw155 = new _module.C3();
        _nw155.__ctor(_944_v18, _948_v22, !((_this).f19), _module.__default.fm33((_this).f19, globalState));
        _949_v23 = _nw155;
        let _950_v24;
        _950_v24 = _dafny.Map.Empty.slice().updateUnsafe((_944_v18).plus(new BigNumber((_dafny.Seq.of(_944_v18)).length)),_949_v23);
        let _951_v25;
        _951_v25 = _dafny.Map.Empty.slice().updateUnsafe(_944_v18,new BigNumber(-353));
        let _952_v26;
        _952_v26 = _dafny.MultiSet.fromElements(new BigNumber((_951_v25).length));
        let _953_v27;
        _953_v27 = _module.D1.create_DC4(_944_v18, _module.__default.fm0(_949_v23.f14, (_dafny.ZERO).minus(new BigNumber((_948_v22).length)), new BigNumber(548), _952_v26, globalState));
        _950_v24 = (_950_v24).update((_953_v27).dtor_cf5, _949_v23);
        _944_v18 = _module.__default.safeModuloInt(_944_v18, (_944_v18).multipliedBy(_944_v18));
      }
      let _954_v28;
      _954_v28 = new BigNumber(-847);
      let _955_v29;
      _955_v29 = _dafny.Seq.of(_954_v28);
      let _956_v30;
      _956_v30 = _dafny.MultiSet.fromElements((_955_v29)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-609)), function (_957_i2) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length), new BigNumber((_955_v29).length))]);
      let _958_v31;
      _958_v31 = _dafny.Seq.of(_954_v28, _954_v28, new BigNumber((_956_v30).cardinality()));
      _955_v29 = _958_v31;
      let _959_v32;
      _959_v32 = _dafny.Seq.UnicodeFromString("hsfh");
      (_this).f14 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("eduiltxa"), _959_v32)) || (_this.f14);
      let _960_v33;
      _960_v33 = _dafny.Map.Empty.slice().updateUnsafe(_954_v28,_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_dafny.Seq.of(_this.f14, false, _this.f14)).length)));
      let _961_v34;
      _961_v34 = _dafny.Map.Empty.slice().updateUnsafe(_960_v33,(_this).f19);
      let _962_v35;
      _962_v35 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-201)), function (_963_i4) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }));
      let _964_v36;
      let _nw156 = Array((new BigNumber(22)).toNumber());
      _nw156[(_dafny.ZERO).toNumber()] = (((_961_v34).contains(_960_v33)) ? ((_961_v34).get(_960_v33)) : (_this.f14));
      _nw156[(_dafny.ONE).toNumber()] = (new BigNumber(377)).isLessThanOrEqualTo(_954_v28);
      _nw156[(new BigNumber(2)).toNumber()] = (_this).f19;
      _nw156[(new BigNumber(3)).toNumber()] = (_954_v28).isLessThanOrEqualTo((_dafny.ZERO).minus(_954_v28));
      _nw156[(new BigNumber(4)).toNumber()] = (_this).f19;
      _nw156[(new BigNumber(5)).toNumber()] = (_this).fm2((_this).f19, _this.f14, globalState);
      _nw156[(new BigNumber(6)).toNumber()] = !((_this).f19);
      _nw156[(new BigNumber(7)).toNumber()] = (_this).f19;
      _nw156[(new BigNumber(8)).toNumber()] = _this.f14;
      _nw156[(new BigNumber(9)).toNumber()] = !_dafny.areEqual(_959_v32, (_962_v35)[_module.__default.safeIndex(_954_v28, new BigNumber((_962_v35).length))]);
      _nw156[(new BigNumber(10)).toNumber()] = (_this).fm2((_this).f19, (_this).f19, globalState);
      _nw156[(new BigNumber(11)).toNumber()] = !((_this).f19);
      _nw156[(new BigNumber(12)).toNumber()] = ((_this).f19) === (_this.f14);
      _nw156[(new BigNumber(13)).toNumber()] = !(_this.f14) || (_this.f14);
      _nw156[(new BigNumber(14)).toNumber()] = false;
      _nw156[(new BigNumber(15)).toNumber()] = (true) || (false);
      _nw156[(new BigNumber(16)).toNumber()] = ((_this.f14) ? (_this.f14) : ((_this).f19));
      _nw156[(new BigNumber(17)).toNumber()] = _this.f14;
      _nw156[(new BigNumber(18)).toNumber()] = false;
      _nw156[(new BigNumber(19)).toNumber()] = _this.f14;
      _nw156[(new BigNumber(20)).toNumber()] = (_dafny.MultiSet.fromElements(_959_v32)).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lpfkgsqvw")));
      _nw156[(new BigNumber(21)).toNumber()] = !(_this.f14);
      _964_v36 = _nw156;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_964_v36).length))) {
        let _965_i3 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_965_i3)) && ((_965_i3).isLessThan(new BigNumber((_964_v36).length))))) {
          (_964_v36)[(_965_i3)] = (_dafny.Set.fromElements(_this.f14, !(true), _this.f14, _this.f14)).IsSubsetOf((_dafny.Set.fromElements(false, _this.f14)).Difference(_dafny.Set.fromElements(true, _this.f14)));
        }
      }
      let _966_v37;
      let _nw157 = new _module.C1();
      _nw157.__ctor();
      _966_v37 = _nw157;
      let _967_v38;
      _967_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
      let _968_v39;
      _968_v39 = _dafny.Set.fromElements(_954_v28, new BigNumber(608), new BigNumber((_967_v38).length), _954_v28);
      r1 = !((_968_v39).IsSubsetOf(_968_v39));
      r0 = new BigNumber((_959_v32).length);
      let _969_v40;
      _969_v40 = new _dafny.CodePoint('s'.codePointAt(0));
      let _970_v41;
      _970_v41 = _dafny.Map.Empty.slice().updateUnsafe(_969_v40,_954_v28);
      let _971_v42;
      _971_v42 = _dafny.Set.fromElements(_955_v29, _958_v31);
      r1 = _module.__default.fm23((_954_v28).multipliedBy(_954_v28), (_this).f19, (_970_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(_969_v40,_954_v28)), _971_v42, globalState);
      let _972_v43;
      _972_v43 = _dafny.Map.Empty.slice().updateUnsafe(_954_v28,_this.f14);
      r2 = _dafny.Map.Empty.slice().updateUnsafe((_972_v43).Merge(_dafny.Map.Empty.slice().updateUnsafe(_954_v28,(_this).fm2(_this.f14, _this.f14, globalState))),_954_v28);
      r3 = (_954_v28).isEqualTo(_954_v28);
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      r1 = p1;
      if (!(!(!((_this).f19) || (_this.f14))) || ((_this).f19)) {
        let _973_v0;
        let _init20 = function (_974_i0) {
          return _dafny.Seq.UnicodeFromString("mpehub");
        };
        let _nw158 = Array((new BigNumber(3)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw158.length); _i0_20++) {
          _nw158[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _973_v0 = _nw158;
        let _975_v2;
        _975_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _rhs135 = new BigNumber((function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of (_975_v2).Keys.Elements) {
            let _976_v1 = _compr_43;
            if ((_975_v2).contains(_976_v1)) {
              _coll43.push([(_976_v1).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC9(_this.f14, p1, new _dafny.CodePoint('q'.codePointAt(0)))).dtor_cf13))).cardinality())),_this.f14]);
            }
          }
          return _coll43;
        }()).length);
        let _rhs136 = _973_v0;
        r1 = _rhs135;
        _973_v0 = _rhs136;
        r1 = p1;
        r2 = true;
        let _977_v3;
        _977_v3 = _dafny.Seq.UnicodeFromString("iycqg");
        let _978_v4;
        let _nw159 = new _module.C3();
        _nw159.__ctor(p1, _977_v3, _this.f14, (_this).f15);
        _978_v4 = _nw159;
        let _979_v5;
        let _init21 = ((_980_p1) => function (_981_i1) {
          return (_981_i1).multipliedBy(new BigNumber((_dafny.Seq.of(_980_p1)).length));
        })(p1);
        let _nw160 = Array((new BigNumber(27)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw160.length); _i0_21++) {
          _nw160[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _979_v5 = _nw160;
        let _982_v6;
        _982_v6 = _dafny.Map.Empty.slice().updateUnsafe(_979_v5,_978_v4);
        if (!(_dafny.Set.fromElements(_978_v4, (((_982_v6).contains(_979_v5)) ? ((_982_v6).get(_979_v5)) : (_978_v4)), _978_v4, _978_v4)).contains(_978_v4)) {
          let _983_v7;
          _983_v7 = new _dafny.CodePoint('v'.codePointAt(0));
          let _984_v8;
          let _nw161 = new _module.C3();
          _nw161.__ctor(p1, _dafny.Seq.update(_977_v3, _module.__default.safeIndex(p1, new BigNumber((_977_v3).length)), _983_v7), _this.f14, (_978_v4).f15);
          _984_v8 = _nw161;
          let _index173 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_973_v0).length));
          (_973_v0)[_index173] = (_984_v8).f21;
          let _index174 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_973_v0).length));
          (_973_v0)[_index174] = ((!((_this).f19)) ? ((_984_v8).f21) : ((_984_v8).f21));
          let _985_v9;
          _985_v9 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(749)), ((_986_v8) => function (_987_i2) {
            return new BigNumber(((_986_v8).f21).length);
          })(_984_v8)));
          let _988_v10;
          _988_v10 = _dafny.Seq.of((_this).f19);
          let _989_v11;
          _989_v11 = _dafny.Seq.of((_988_v10)[_module.__default.safeIndex(p1, new BigNumber((_988_v10).length))]);
          let _990_v12;
          _990_v12 = _dafny.Set.fromElements((_984_v8).f21);
          r2 = ((!(_module.__default.fm23((_984_v8).f20, _978_v4.f14, _dafny.Map.Empty.slice().updateUnsafe(_983_v7,p1), _985_v9, globalState))) ? (_dafny.areEqual(_989_v11, _dafny.Seq.update(_988_v10, _module.__default.safeIndex((_984_v8).f20, new BigNumber((_988_v10).length)), (_this).f19))) : ((_990_v12).IsProperSubsetOf(_module.__default.fm34(globalState))));
          let _991_v13;
          let _nw162 = Array((new BigNumber(8)).toNumber()).fill(false);
          _991_v13 = _nw162;
          (globalState).f2 = _991_v13;
          let _992_v14;
          let _init22 = ((_993_v4) => function (_994_i3) {
            return _dafny.Seq.update(_dafny.Seq.of(_993_v4.f14), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_this.f14)).length), new BigNumber((_dafny.Seq.of(_993_v4.f14)).length)), true);
          })(_978_v4);
          let _nw163 = Array((new BigNumber(23)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw163.length); _i0_22++) {
            _nw163[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _992_v14 = _nw163;
          let _index175 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_992_v14).length));
          (_992_v14)[_index175] = _dafny.Seq.of(true, _978_v4.f14, _978_v4.f14);
          let _index176 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_979_v5).length));
          (_979_v5)[_index176] = ((_984_v8).f20).multipliedBy(p1);
          let _index177 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_992_v14).length));
          let _index178 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_979_v5).length));
          let _rhs137 = _988_v10;
          let _rhs138 = _983_v7;
          let _rhs139 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber((_977_v3).length)).plus(p1)), new BigNumber((((_978_v4).f15).Difference((_978_v4).f15)).cardinality()));
          let _rhs140 = _978_v4.f14;
          let _lhs116 = _992_v14;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_992_v14).length));
          let _lhs118 = _979_v5;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_979_v5).length));
          let _lhs120 = _this;
          _lhs116[_lhs117] = _rhs137;
          _983_v7 = _rhs138;
          _lhs118[_lhs119] = _rhs139;
          _lhs120.f14 = _rhs140;
        } else {
          r0 = (p1).multipliedBy((_dafny.ZERO).minus(p1));
          r0 = (_dafny.ZERO).minus((new BigNumber(621)).minus(((_978_v4.f14) ? (p1) : ((_dafny.ZERO).minus(p1)))));
          let _995_v15;
          _995_v15 = _dafny.MultiSet.fromElements((_this).f19, !(p1).isEqualTo(p1), false, false, true);
          _995_v15 = (((_dafny.MultiSet.fromElements(_978_v4.f14)).update(_978_v4.f14, _module.__default.abs(p1))).Difference(_dafny.MultiSet.fromElements(_978_v4.f14, _this.f14, (_this).f19, (_this).f19))).Difference(((_978_v4).f15).Union(_995_v15));
          let _index179 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_979_v5).length));
          (_979_v5)[_index179] = p1;
          let _index180 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_979_v5).length));
          (_979_v5)[_index180] = p1;
          r2 = _978_v4.f14;
        }
        let _996_v16;
        let _nw164 = Array((new BigNumber(21)).toNumber()).fill(false);
        _996_v16 = _nw164;
        let _997_v17;
        _997_v17 = _module.D5.create_DC18(_978_v4.f14);
        let _index181 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_996_v16).length));
        (_996_v16)[_index181] = (_997_v17).dtor_cf33;
        let _index182 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_996_v16).length));
        (_996_v16)[_index182] = !((!((_this).f19)) && (_this.f14));
      } else {
        (globalState).f9 = (new BigNumber(819)).isLessThan(p1);
        let _998_v18;
        _998_v18 = new _dafny.CodePoint('c'.codePointAt(0));
        let _999_v19;
        _999_v19 = _dafny.Map.Empty.slice().updateUnsafe(_998_v18,p1);
        let _1000_v20;
        _1000_v20 = _dafny.Seq.of(p1, p1);
        let _1001_v21;
        _1001_v21 = _dafny.MultiSet.fromElements(new BigNumber(-711));
        let _1002_v22;
        _1002_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        if (!(!(_1002_v22).contains(_module.__default.fm0(true, p1, _module.__default.fm0((_this).f19, new BigNumber(877), p1, _1001_v21, globalState), _1001_v21, globalState))) || (_module.__default.fm23(p1, _this.f14, _999_v19, _dafny.Set.fromElements(_1000_v20), globalState))) {
          (globalState).f3 = p1;
          let _1003_v23;
          let _nw165 = new _module.C1();
          _nw165.__ctor();
          _1003_v23 = _nw165;
          _1003_v23 = _1003_v23;
          let _1004_v24;
          _1004_v24 = _dafny.Seq.of(_this.f14, (_this).f19, _this.f14, _this.f14);
          let _1005_v25;
          _1005_v25 = _dafny.MultiSet.fromElements(_1004_v24);
          let _1006_v26;
          _1006_v26 = _dafny.Seq.of(_1001_v21);
          r1 = (_this).fm17(p1, _this.f14, (_dafny.MultiSet.fromElements(_1004_v24)).Difference(_1005_v25), new BigNumber((_1006_v26).length), globalState);
          (_this).m6(globalState);
          let _1007_v27;
          let _nw166 = Array((new BigNumber(6)).toNumber()).fill(_module.D1.Default());
          _1007_v27 = _nw166;
          let _1008_v28;
          _1008_v28 = _dafny.Seq.of(_1007_v27, _1007_v27, _1007_v27, _1007_v27);
          r1 = (_dafny.ZERO).minus(new BigNumber((_1008_v28).length));
        } else {
          let _1009_v29;
          _1009_v29 = _dafny.Seq.UnicodeFromString("erx");
          let _1010_v30;
          _1010_v30 = _dafny.Seq.of((_this).f19, _this.f14);
          let _1011_v31;
          let _nw167 = new _module.C3();
          _nw167.__ctor(p1, _1009_v29, !(_this.f14), _dafny.MultiSet.FromArray(_1010_v30));
          _1011_v31 = _nw167;
          let _rhs141 = !(new BigNumber(-958)).isEqualTo(p1);
          let _rhs142 = _dafny.Seq.of(p1);
          let _rhs143 = _module.__default.fm0((_1011_v31).fm2(_this.f14, (_this).f19, globalState), p1, new BigNumber((_module.__default.fm35(p1, p1, globalState)).length), _1001_v21, globalState);
          let _rhs144 = (_this).f19;
          let _lhs121 = globalState;
          let _lhs122 = globalState;
          r2 = _rhs141;
          _1000_v20 = _rhs142;
          _lhs121.f3 = _rhs143;
          _lhs122.f9 = _rhs144;
          r0 = p1;
          (globalState).f3 = (_1011_v31).f20;
          let _1012_v32;
          let _nw168 = new _module.C3();
          _nw168.__ctor((_1011_v31).f20, _dafny.Seq.UnicodeFromString("oek"), _this.f14, (_this).f15);
          _1012_v32 = _nw168;
        }
        _998_v18 = _998_v18;
        let _1013_v33;
        _1013_v33 = _module.D0.create_DC0((_this).f19);
        let _1014_v34;
        _1014_v34 = _module.D0.create_DC2(_1013_v33);
        let _1015_v35;
        _1015_v35 = _module.D0.create_DC2(_1014_v34);
        let _source12 = _1015_v35;
        if (_source12.is_DC1) {
          let _1016___mcc_h0 = (_source12).cf1;
          let _1017_cf1 = _1016___mcc_h0;
          let _1018_v36;
          _1018_v36 = _dafny.Seq.UnicodeFromString("ukldksby");
          let _1019_v37;
          _1019_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,p1);
          _998_v18 = _module.__default.fm25(p1, _module.__default.safeDivisionInt(p1, p1), _dafny.Seq.contains(_1018_v36, _998_v18), _1019_v37, globalState);
          (globalState).f3 = p1;
          (globalState).f9 = _1017_cf1;
          let _1020_v38;
          let _init23 = ((_1021_cf1) => function (_1022_i4) {
            return _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC5(_this.f14, (_this).f19)).dtor_cf6,_1021_cf1));
          })(_1017_cf1);
          let _nw169 = Array((new BigNumber(8)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw169.length); _i0_23++) {
            _nw169[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1020_v38 = _nw169;
          _1020_v38 = _1020_v38;
        } else if (_source12.is_DC0) {
          let _1023___mcc_h1 = (_source12).cf0;
          let _1024_cf0 = _1023___mcc_h1;
          (globalState).f3 = _module.__default.safeDivisionInt(p1, (((_1002_v22).contains(p1)) ? ((_1002_v22).get(p1)) : (p1)));
          let _1025_v39;
          _1025_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1024_cf0,p1);
          let _1026_v40;
          _1026_v40 = _dafny.Seq.of(_1025_v39);
          let _1027_v41;
          _1027_v41 = _dafny.Seq.of(_module.__default.fm36(_1001_v21, p1, _1024_cf0, (_this).f19, globalState));
          _1026_v40 = _dafny.Seq.Concat((_1027_v41)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1027_v41).length))], _1026_v40);
          let _1028_v42;
          let _nw170 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1028_v42 = _nw170;
          let _index183 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1028_v42).length));
          (_1028_v42)[_index183] = _this.f14;
          let _1029_v43;
          _1029_v43 = _dafny.Seq.of(_1024_cf0);
          let _index184 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1028_v42).length));
          (_1028_v42)[_index184] = (!((_this).f19) || ((_this).f19)) && ((_1029_v43)[_module.__default.safeIndex(p1, new BigNumber((_1029_v43).length))]);
          (_this).m6(globalState);
        } else {
          let _1030___mcc_h2 = (_source12).cf2;
          let _1031_cf2 = _1030___mcc_h2;
          let _1032_v44;
          let _nw171 = Array((new BigNumber(3)).toNumber()).fill([]);
          _1032_v44 = _nw171;
          let _1033_v45;
          let _init24 = function (_1034_i5) {
            return _module.__default.safeDivisionInt(_1034_i5, new BigNumber(588));
          };
          let _nw172 = Array((new BigNumber(25)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw172.length); _i0_24++) {
            _nw172[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1033_v45 = _nw172;
          let _index185 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length));
          (_1032_v44)[_index185] = _1033_v45;
          let _index186 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length));
          (_1032_v44)[_index186] = _1033_v45;
          let _1035_v46;
          _1035_v46 = _module.D7.create_DC23(_this.f14, new BigNumber((_dafny.Seq.update(_1000_v20, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1000_v20).length)), p1)).length), _this.f14);
          let _1036_v47;
          _1036_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1035_v46,(_this).fm2((_this).f19, (_this).f19, globalState));
          _1036_v47 = (_1036_v47).update(_1035_v46, (_this).f19);
          let _1037_v48;
          _1037_v48 = _dafny.Seq.of(_this.f14);
          let _1038_v49;
          let _init25 = function (_1039_i6) {
            return (_this).f19;
          };
          let _nw173 = Array((_dafny.ONE).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw173.length); _i0_25++) {
            _nw173[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1038_v49 = _nw173;
          let _1040_v50;
          _1040_v50 = _dafny.Set.fromElements((_this).f19);
          let _1041_v51;
          _1041_v51 = _module.D8.create_DC24(_1040_v50);
          let _1042_v52;
          let _nw174 = new _module.C1();
          _nw174.__ctor();
          _1042_v52 = _nw174;
          let _1043_v53;
          _1043_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1042_v52);
          let _1044_v54;
          _1044_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1043_v53,_dafny.Set.fromElements((_this).fm2(true, (_this).f19, globalState), !(false)));
          let _1045_v55;
          _1045_v55 = _dafny.MultiSet.fromElements(_1037_v48, _1037_v48);
          let _1046_v56;
          let _nw175 = Array((new BigNumber(25)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = (_this).f19;
          _nw175[(_dafny.ONE).toNumber()] = (_this).f19;
          _nw175[(new BigNumber(2)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(3)).toNumber()] = !((((_this).f19) ? ((_this).f19) : ((_this).f19)));
          _nw175[(new BigNumber(4)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(5)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(6)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(7)).toNumber()] = !((_this).f19);
          _nw175[(new BigNumber(8)).toNumber()] = true;
          _nw175[(new BigNumber(9)).toNumber()] = (p1).isEqualTo(new BigNumber((_1037_v48).length));
          _nw175[(new BigNumber(10)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_1038_v49,_998_v18)).contains(_1038_v49);
          _nw175[(new BigNumber(11)).toNumber()] = !((_this).f19);
          _nw175[(new BigNumber(12)).toNumber()] = !(_this.f14);
          _nw175[(new BigNumber(13)).toNumber()] = ((((_1002_v22).contains(p1)) ? ((_1002_v22).get(p1)) : (p1))).isLessThanOrEqualTo(new BigNumber((p0).length));
          _nw175[(new BigNumber(14)).toNumber()] = (((_this).f19) ? ((_this).f19) : ((_this).f19));
          _nw175[(new BigNumber(15)).toNumber()] = (_this).f19;
          _nw175[(new BigNumber(16)).toNumber()] = (_this).f19;
          _nw175[(new BigNumber(17)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(18)).toNumber()] = ((_1041_v51).dtor_cf45).IsDisjointFrom((((_1044_v54).contains(_1043_v53)) ? ((_1044_v54).get(_1043_v53)) : (_1040_v50)));
          _nw175[(new BigNumber(19)).toNumber()] = true;
          _nw175[(new BigNumber(20)).toNumber()] = !((((_this).f19) ? ((_this).f19) : ((_this).fm2((_this).f19, _this.f14, globalState))));
          _nw175[(new BigNumber(21)).toNumber()] = !((_this).f19) || (_this.f14);
          _nw175[(new BigNumber(22)).toNumber()] = (p1).isEqualTo((_this).fm17(p1, (_this).f19, _1045_v55, new BigNumber(29), globalState));
          _nw175[(new BigNumber(23)).toNumber()] = _this.f14;
          _nw175[(new BigNumber(24)).toNumber()] = (_this).f19;
          _1046_v56 = _nw175;
          let _index187 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1038_v49).length));
          (_1038_v49)[_index187] = (_this).f19;
          let _1047_v57;
          let _nw176 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1047_v57 = _nw176;
          let _index188 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_1047_v57).length));
          (_1047_v57)[_index188] = _1038_v49;
          let _1048_v58;
          _1048_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1038_v49);
          let _index189 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1038_v49).length));
          let _index190 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_1047_v57).length));
          let _rhs145 = !(!(!(_this.f14)));
          let _rhs146 = (((_1048_v58).contains((_this).f19)) ? ((_1048_v58).get((_this).f19)) : (_1046_v56));
          let _lhs123 = _1038_v49;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1038_v49).length));
          let _lhs125 = _1047_v57;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_1047_v57).length));
          _lhs123[_lhs124] = _rhs145;
          _lhs125[_lhs126] = _rhs146;
          let _1049_v59;
          _1049_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-263),false);
          let _arr2 = (_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))];
          let _index191 = _module.__default.safeIndex(new BigNumber(932), new BigNumber(((_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))]).length));
          _arr2[_index191] = (((_1001_v21).contains(p1)) ? ((_1001_v21).get(p1)) : (new BigNumber((_1049_v59).length)));
          let _1050_v60;
          _1050_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1038_v49)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_1038_v49).length))],new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), ((_1051_p1) => function (_1052_i7) {
            return _1051_p1;
          })(p1))).length));
          let _arr3 = (_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))];
          let _index192 = _module.__default.safeIndex(new BigNumber(932), new BigNumber(((_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))]).length));
          let _rhs147 = (((_1050_v60).contains((_this).f19)) ? ((_1050_v60).get((_this).f19)) : (p1));
          let _rhs148 = (p1).minus(p1);
          let _lhs127 = globalState;
          let _lhs128 = (_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))];
          let _lhs129 = _module.__default.safeIndex(new BigNumber(932), new BigNumber(((_1032_v44)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1032_v44).length))]).length));
          _lhs127.f3 = _rhs147;
          _lhs128[_lhs129] = _rhs148;
        }
        let _1053_v61;
        let _init26 = ((_1054_p1) => function (_1055_i8) {
          return (_1055_i8).plus(_1054_p1);
        })(p1);
        let _nw177 = Array((new BigNumber(6)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw177.length); _i0_26++) {
          _nw177[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1053_v61 = _nw177;
        let _index193 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1053_v61).length));
        (_1053_v61)[_index193] = p1;
        let _1056_v62;
        _1056_v62 = _dafny.Seq.of((_this).fm2(_this.f14, false, globalState));
        let _index194 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1053_v61).length));
        let _rhs149 = p1;
        let _rhs150 = p1;
        let _rhs151 = (p1).multipliedBy(new BigNumber((_dafny.Seq.Concat(_1056_v62, _1056_v62)).length));
        let _rhs152 = (_this).fm2(_this.f14, (_this).f19, globalState);
        let _rhs153 = (((_this).fm2(_this.f14, _this.f14, globalState)) ? (new BigNumber((_dafny.Seq.Concat(_1056_v62, _1056_v62)).length)) : (p1));
        let _lhs130 = _1053_v61;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1053_v61).length));
        let _lhs132 = globalState;
        let _lhs133 = globalState;
        r0 = _rhs149;
        _lhs130[_lhs131] = _rhs150;
        _lhs132.f3 = _rhs151;
        r2 = _rhs152;
        _lhs133.f3 = _rhs153;
      }
      let _1057_v63;
      _1057_v63 = _dafny.Seq.UnicodeFromString("iep");
      _1057_v63 = _1057_v63;
      if ((_this).f19) {
        (globalState).f9 = (_this).f19;
        let _1058_v64;
        _1058_v64 = _module.D4.create_DC14(false, (_this).f19);
        let _1059_v65;
        let _nw178 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _1059_v65 = _nw178;
        let _1060_v66;
        _1060_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1059_v65,!((_this).f19));
        let _1061_v67;
        _1061_v67 = _module.D5.create_DC16(_1059_v65);
        let _1062_v68;
        _1062_v68 = _dafny.Set.fromElements(false);
        let _1063_v69;
        let _nw179 = Array((new BigNumber(22)).toNumber());
        _nw179[(_dafny.ZERO).toNumber()] = !((_1058_v64).dtor_cf26);
        _nw179[(_dafny.ONE).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(2)).toNumber()] = false;
        _nw179[(new BigNumber(3)).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(4)).toNumber()] = ((true) ? (_this.f14) : ((_this).f19));
        _nw179[(new BigNumber(5)).toNumber()] = (_this).fm2(_this.f14, (_this).f19, globalState);
        _nw179[(new BigNumber(6)).toNumber()] = _this.f14;
        _nw179[(new BigNumber(7)).toNumber()] = _this.f14;
        _nw179[(new BigNumber(8)).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(9)).toNumber()] = _this.f14;
        _nw179[(new BigNumber(10)).toNumber()] = (_this).fm2(_this.f14, _this.f14, globalState);
        _nw179[(new BigNumber(11)).toNumber()] = (((_this).f19) ? ((_this).f19) : ((_this).f19));
        _nw179[(new BigNumber(12)).toNumber()] = ((_this).f19) === ((((_1060_v66).contains((_1061_v67).dtor_cf29)) ? ((_1060_v66).get((_1061_v67).dtor_cf29)) : (_this.f14)));
        _nw179[(new BigNumber(13)).toNumber()] = !(_this.f14);
        _nw179[(new BigNumber(14)).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(15)).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(16)).toNumber()] = (_dafny.Set.fromElements(!((_this).f19), _this.f14, (_this).f19, _this.f14)).IsSubsetOf(_1062_v68);
        _nw179[(new BigNumber(17)).toNumber()] = (_this).f19;
        _nw179[(new BigNumber(18)).toNumber()] = (_this).fm2(_this.f14, (_this).f19, globalState);
        _nw179[(new BigNumber(19)).toNumber()] = _this.f14;
        _nw179[(new BigNumber(20)).toNumber()] = (p1).isEqualTo(p1);
        _nw179[(new BigNumber(21)).toNumber()] = !((_this).f19);
        _1063_v69 = _nw179;
        let _index195 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1063_v69).length));
        (_1063_v69)[_index195] = _this.f14;
        let _index196 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1063_v69).length));
        (_1063_v69)[_index196] = _this.f14;
        let _1064_v70;
        _1064_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1063_v69);
        _1064_v70 = (_1064_v70).update((_this).f19, _1063_v69);
        let _1065_v71;
        _1065_v71 = _dafny.Seq.of((_this).f15);
        (_this).m7(_1065_v71, (_1063_v69)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1063_v69).length))], p1, (_1063_v69)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1063_v69).length))], globalState);
        r2 = (p1).isEqualTo(p1);
      } else {
        (globalState).f3 = _module.__default.safeDivisionInt(new BigNumber(-458), new BigNumber(715));
        let _1066_v72;
        _1066_v72 = _dafny.Set.fromElements(true);
        if (!(_1066_v72).contains(_this.f14)) {
          let _1067_v73;
          let _nw180 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _1067_v73 = _nw180;
          _1067_v73 = _1067_v73;
          let _1068_v74;
          _1068_v74 = _dafny.Seq.of(_this.f14, (_this).f19);
          (globalState).f9 = (_1068_v74)[_module.__default.safeIndex(p1, new BigNumber((_1068_v74).length))];
          let _1069_v75;
          _1069_v75 = _module.D0.create_DC1(!(_this.f14));
          let _1070_v76;
          _1070_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(500),_dafny.Set.fromElements((_this).f19, _this.f14));
          let _1071_v77;
          let _nw181 = Array((new BigNumber(25)).toNumber());
          _nw181[(_dafny.ZERO).toNumber()] = _1069_v75;
          _nw181[(_dafny.ONE).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(2)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(3)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(4)).toNumber()] = function (_pat_let27_0) {
            return function (_1072_dt__update__tmp_h0) {
              return function (_pat_let28_0) {
                return function (_1073_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_1073_dt__update_hcf1_h0);
                }(_pat_let28_0);
              }((_this).f19);
            }(_pat_let27_0);
          }(_1069_v75);
          _nw181[(new BigNumber(5)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(6)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(7)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(8)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(9)).toNumber()] = _module.__default.fm37(p0, _1070_v76, (_this).f19, (_module.D0.create_DC1(!(_this.f14))).dtor_cf1, globalState);
          _nw181[(new BigNumber(10)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(11)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(12)).toNumber()] = function (_pat_let29_0) {
            return function (_1074_dt__update__tmp_h1) {
              return function (_pat_let30_0) {
                return function (_1075_dt__update_hcf1_h1) {
                  return _module.D0.create_DC1(_1075_dt__update_hcf1_h1);
                }(_pat_let30_0);
              }(_this.f14);
            }(_pat_let29_0);
          }(_module.__default.fm37(p0, _1070_v76, _this.f14, (_this).f19, globalState));
          _nw181[(new BigNumber(13)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(14)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(15)).toNumber()] = _module.D0.create_DC1(_this.f14);
          _nw181[(new BigNumber(16)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(17)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(18)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(19)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(20)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(21)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(22)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(23)).toNumber()] = _1069_v75;
          _nw181[(new BigNumber(24)).toNumber()] = function (_pat_let31_0) {
            return function (_1076_dt__update__tmp_h2) {
              return function (_pat_let32_0) {
                return function (_1077_dt__update_hcf1_h2) {
                  return _module.D0.create_DC1(_1077_dt__update_hcf1_h2);
                }(_pat_let32_0);
              }(_this.f14);
            }(_pat_let31_0);
          }(_1069_v75);
          _1071_v77 = _nw181;
          _1071_v77 = _1071_v77;
          (globalState).f9 = false;
          _1057_v63 = _1057_v63;
        } else {
          (globalState).f9 = ((_this.f14) ? (_this.f14) : ((_this).f19));
          let _1078_v78;
          let _init27 = function (_1079_i9) {
            return ((_this).f15).IsDisjointFrom((_this).f15);
          };
          let _nw182 = Array((new BigNumber(6)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw182.length); _i0_27++) {
            _nw182[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1078_v78 = _nw182;
          (globalState).f2 = _1078_v78;
          r2 = _this.f14;
          let _1080_v79;
          _1080_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_this.f14);
          let _1081_v80;
          _1081_v80 = _dafny.Set.fromElements(_1080_v79, _1080_v79, _1080_v79);
          let _1082_v81;
          _1082_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1081_v80,(_dafny.ZERO).minus((p1).minus(new BigNumber(48))));
          _1082_v81 = (_1082_v81).update(_1081_v80, p1);
          let _1083_v82;
          let _nw183 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1083_v82 = _nw183;
          let _index197 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1083_v82).length));
          (_1083_v82)[_index197] = _module.__default.safeModuloInt(p1, p1);
          let _index198 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1083_v82).length));
          (_1083_v82)[_index198] = new BigNumber(-600);
        }
        let _1084_v83;
        _1084_v83 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1085_v84;
        _1085_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1084_v83,p1);
        let _1086_v85;
        _1086_v85 = _dafny.Seq.of(p1);
        let _1087_v86;
        _1087_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1086_v85,(_this).f19);
        let _1088_v88;
        _1088_v88 = _module.D8.create_DC24(_dafny.Set.fromElements(_this.f14, _module.__default.fm23(p1, _this.f14, _1085_v84, function () {
  let _coll44 = new _dafny.Set();
  for (const _compr_44 of (_1087_v86).Keys.Elements) {
    let _1089_v87 = _compr_44;
    if ((_1087_v86).contains(_1089_v87)) {
      _coll44.add(_1089_v87);
    }
  }
  return _coll44;
}(), globalState)));
        _1088_v88 = _1088_v88;
        let _1090_v89;
        let _init28 = function (_1091_i10) {
          return _module.__default.safeDivisionInt(_1091_i10, new BigNumber(380));
        };
        let _nw184 = Array((_dafny.ONE).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw184.length); _i0_28++) {
          _nw184[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1090_v89 = _nw184;
        let _1092_v90;
        _1092_v90 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1057_v63);
        let _1093_v91;
        _1093_v91 = _module.D8.create_DC28(_module.D8.create_DC26(!(_this.f14), _1092_v90));
        let _1094_v92;
        _1094_v92 = _dafny.MultiSet.fromElements(_1093_v91);
        let _1095_v93;
        _1095_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1090_v89,p1);
        let _1096_v94;
        _1096_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1095_v93);
        let _1097_v95;
        _1097_v95 = _dafny.Seq.of(_this.f14, !(_this.f14));
        let _1098_v96;
        _1098_v96 = _dafny.Seq.of(_1090_v89, _1090_v89);
        let _1099_v97;
        _1099_v97 = _module.D3.create_DC11(new BigNumber((_1094_v92).cardinality()), new BigNumber(((((_1096_v94).contains((_1097_v95)[_module.__default.safeIndex(p1, new BigNumber((_1097_v95).length))])) ? ((_1096_v94).get((_1097_v95)[_module.__default.safeIndex(p1, new BigNumber((_1097_v95).length))])) : (_dafny.Map.Empty.slice().updateUnsafe((_1098_v96)[_module.__default.safeIndex(p1, new BigNumber((_1098_v96).length))],p1)))).length), (_this).f19, p1);
        let _rhs154 = _1090_v89;
        let _rhs155 = (_this).f19;
        let _rhs156 = _1099_v97;
        let _rhs157 = !((_this).f19);
        let _lhs134 = globalState;
        let _lhs135 = _this;
        _1090_v89 = _rhs154;
        _lhs134.f9 = _rhs155;
        _1099_v97 = _rhs156;
        _lhs135.f14 = _rhs157;
        (globalState).f9 = (_1097_v95)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_1097_v95).length))];
      }
      if (((_this).f19) && (!(_this.f14))) {
        let _1100_v98;
        _1100_v98 = _dafny.Seq.of((_this).f19);
        r1 = new BigNumber((((_dafny.MultiSet.fromElements(_this.f14)).Intersect((_this).f15)).Union((_dafny.MultiSet.FromArray(_1100_v98)).Union((_this).f15))).cardinality());
        if (!((_this).fm2((p1).isLessThan(p1), (_this).f19, globalState))) {
          (globalState).f9 = (_this).f19;
          _1057_v63 = _1057_v63;
          r1 = (p1).plus(p1);
          let _1101_v99;
          let _init29 = ((_1102_p1) => function (_1103_i11) {
            return (_1103_i11).plus((_dafny.ZERO).minus(_1102_p1));
          })(p1);
          let _nw185 = Array((new BigNumber(10)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw185.length); _i0_29++) {
            _nw185[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1101_v99 = _nw185;
          let _index199 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length));
          (_1101_v99)[_index199] = (p1).plus(p1);
          let _1104_v100;
          _1104_v100 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f19, _this.f14)).length), p1);
          let _1105_v101;
          _1105_v101 = _dafny.MultiSet.fromElements(p1);
          let _1106_v102;
          _1106_v102 = _dafny.Seq.of((p1).multipliedBy(_module.__default.fm0(_this.f14, new BigNumber((_1104_v100).length), p1, _1105_v101, globalState)));
          let _1107_v103;
          let _nw186 = Array((new BigNumber(9)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = !((_this).f19);
          _nw186[(_dafny.ONE).toNumber()] = (_this).f19;
          _nw186[(new BigNumber(2)).toNumber()] = _this.f14;
          _nw186[(new BigNumber(3)).toNumber()] = _this.f14;
          _nw186[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw186[(new BigNumber(5)).toNumber()] = (_this).f19;
          _nw186[(new BigNumber(6)).toNumber()] = !(_this.f14);
          _nw186[(new BigNumber(7)).toNumber()] = _this.f14;
          _nw186[(new BigNumber(8)).toNumber()] = (_this).f19;
          _1107_v103 = _nw186;
          let _1108_v104;
          _1108_v104 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1109_v105;
          _1109_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1108_v104,new BigNumber((_1057_v63).length));
          let _1110_v106;
          _1110_v106 = _dafny.Set.fromElements(_1104_v100);
          let _1111_v107;
          _1111_v107 = _module.D7.create_DC23(_module.__default.fm23(new BigNumber((_dafny.Seq.of(_1107_v103)).length), (_this).f19, _1109_v105, _1110_v106, globalState), p1, _this.f14);
          let _index200 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length));
          let _rhs158 = (_1106_v102)[_module.__default.safeIndex((_1111_v107).dtor_cf43, new BigNumber((_1106_v102).length))];
          let _rhs159 = p1;
          let _lhs136 = _1101_v99;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length));
          _lhs136[_lhs137] = _rhs158;
          r0 = _rhs159;
          let _1112_v108;
          _1112_v108 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _1113_v109;
          _1113_v109 = _dafny.Map.Empty.slice().updateUnsafe(_1112_v108,(_1101_v99)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length))]);
          let _1114_v110;
          _1114_v110 = _module.D7.create_DC21(_1113_v109);
          let _1115_v111;
          _1115_v111 = _dafny.Set.fromElements(_this.f14, _this.f14, (_this).f19, _this.f14, _this.f14);
          let _index201 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length));
          let _rhs160 = new BigNumber((((_dafny.Set.fromElements(_this.f14, false, _this.f14, _this.f14)).Intersect(_dafny.Set.fromElements((_this).f19))).Intersect(_1115_v111)).length);
          let _rhs161 = (_1101_v99)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length))];
          let _rhs162 = (_1101_v99)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length))];
          let _rhs163 = _1114_v110;
          let _lhs138 = globalState;
          let _lhs139 = _1101_v99;
          let _lhs140 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1101_v99).length));
          _lhs138.f3 = _rhs160;
          r0 = _rhs161;
          _lhs139[_lhs140] = _rhs162;
          _1114_v110 = _rhs163;
        } else {
          let _1116_v112;
          _1116_v112 = _dafny.MultiSet.fromElements(_module.D8.create_DC25());
          let _1117_v113;
          _1117_v113 = _module.D2.create_DC7(_1100_v98);
          let _rhs164 = new BigNumber(((_1117_v113).dtor_cf9).length);
          let _rhs165 = _module.__default.safeModuloInt((p1).plus(p1), p1);
          let _rhs166 = (_1116_v112).Intersect(_1116_v112);
          let _rhs167 = _dafny.Seq.UnicodeFromString("e");
          let _lhs141 = globalState;
          _lhs141.f3 = _rhs164;
          r1 = _rhs165;
          _1116_v112 = _rhs166;
          _1057_v63 = _rhs167;
          (globalState).f3 = p1;
          let _1118_v114;
          _1118_v114 = _dafny.Seq.of(p1);
          let _1119_v115;
          _1119_v115 = _dafny.MultiSet.fromElements(p1);
          let _1120_v116;
          _1120_v116 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_dafny.MultiSet.FromArray(_1118_v114)).Intersect(_1119_v115));
          _1120_v116 = (_1120_v116).update(_this.f14, (_1119_v115).Intersect(_1119_v115));
          let _1121_v117;
          _1121_v117 = _module.D3.create_DC11(p1, p1, (_this).f19, p1);
          let _1122_v118;
          let _nw187 = new _module.C3();
          _nw187.__ctor(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), function (_1123_i12) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }), (_this).f19, (_this).f15);
          _1122_v118 = _nw187;
          let _1124_v119;
          _1124_v119 = _dafny.MultiSet.fromElements(_1122_v118, _1122_v118);
          let _1125_v120;
          _1125_v120 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1126_v121;
          _1126_v121 = _dafny.Set.fromElements(_1125_v120, _1125_v120);
          let _1127_v122;
          _1127_v122 = _dafny.MultiSet.fromElements(_1126_v121);
          let _1128_v123;
          _1128_v123 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_1127_v122).cardinality()));
          let _1129_v124;
          _1129_v124 = _module.D8.create_DC24(_dafny.Set.fromElements((_this).f19, false));
          let _1130_v125;
          _1130_v125 = _dafny.Map.Empty.slice().updateUnsafe(_1129_v124,_this.f14);
          let _1131_v126;
          _1131_v126 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_this).f19);
          let _pat_let_tv23 = _1130_v125;
          let _pat_let_tv24 = p1;
          let _pat_let_tv25 = _1128_v123;
          let _pat_let_tv26 = _1122_v118;
          let _pat_let_tv27 = _1119_v115;
          let _pat_let_tv28 = globalState;
          let _pat_let_tv29 = _1122_v118;
          let _pat_let_tv30 = _1119_v115;
          let _pat_let_tv31 = globalState;
          let _pat_let_tv32 = _1122_v118;
          let _1132_v127;
          let _nw188 = Array((new BigNumber(20)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = _1121_v117;
          _nw188[(_dafny.ONE).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(2)).toNumber()] = (((_this).f19) ? (_module.D3.create_DC11(p1, p1, (_this).f19, new BigNumber(967))) : (function (_pat_let33_0) {
            return function (_1133_dt__update__tmp_h3) {
              return function (_pat_let34_0) {
                return function (_1134_dt__update_hcf19_h0) {
                  return _module.D3.create_DC11((_1133_dt__update__tmp_h3).dtor_cf17, (_1133_dt__update__tmp_h3).dtor_cf18, _1134_dt__update_hcf19_h0, (_1133_dt__update__tmp_h3).dtor_cf20);
                }(_pat_let34_0);
              }(_this.f14);
            }(_pat_let33_0);
          }(_1121_v117)));
          _nw188[(new BigNumber(3)).toNumber()] = _module.D3.create_DC11(p1, p1, _this.f14, new BigNumber(485));
          _nw188[(new BigNumber(4)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(5)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(6)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(7)).toNumber()] = _module.D3.create_DC11(p1, p1, true, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1124_v119).cardinality()))));
          _nw188[(new BigNumber(8)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(9)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(10)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(11)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(12)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(13)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(14)).toNumber()] = function (_pat_let35_0) {
            return function (_1135_dt__update__tmp_h4) {
              return function (_pat_let36_0) {
                return function (_1136_dt__update_hcf19_h1) {
                  return function (_pat_let37_0) {
                    return function (_1137_dt__update_hcf18_h0) {
                      return function (_pat_let38_0) {
                        return function (_1138_dt__update_hcf20_h0) {
                          return _module.D3.create_DC11((_1135_dt__update__tmp_h4).dtor_cf17, _1137_dt__update_hcf18_h0, _1136_dt__update_hcf19_h1, _1138_dt__update_hcf20_h0);
                        }(_pat_let38_0);
                      }(_module.__default.fm0((_this).f19, _module.__default.fm0(_this.f14, new BigNumber((_dafny.Set.fromElements(_pat_let_tv24, new BigNumber((_pat_let_tv25).length))).length), (_pat_let_tv26).f20, _pat_let_tv27, _pat_let_tv28), (_pat_let_tv29).f20, _pat_let_tv30, _pat_let_tv31));
                    }(_pat_let37_0);
                  }(new BigNumber((_pat_let_tv23).length));
                }(_pat_let36_0);
              }((_this).f19);
            }(_pat_let35_0);
          }(_module.D3.create_DC11(p1, _module.__default.fm0(true, _module.__default.fm0((_this).f19, new BigNumber(535), p1, _1119_v115, globalState), (_1122_v118).f20, _dafny.MultiSet.FromArray(_1118_v114), globalState), _this.f14, p1));
          _nw188[(new BigNumber(15)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(16)).toNumber()] = _1121_v117;
          _nw188[(new BigNumber(17)).toNumber()] = _module.D3.create_DC11(new BigNumber((_1128_v123).length), new BigNumber(((_1131_v126).update((_this).f19, (_this).f19)).length), _this.f14, (_1122_v118).f20);
          _nw188[(new BigNumber(18)).toNumber()] = (((_this).f19) ? (_1121_v117) : (function (_pat_let39_0) {
            return function (_1139_dt__update__tmp_h5) {
              return function (_pat_let40_0) {
                return function (_1140_dt__update_hcf20_h1) {
                  return _module.D3.create_DC11((_1139_dt__update__tmp_h5).dtor_cf17, (_1139_dt__update__tmp_h5).dtor_cf18, (_1139_dt__update__tmp_h5).dtor_cf19, _1140_dt__update_hcf20_h1);
                }(_pat_let40_0);
              }((_pat_let_tv32).f20);
            }(_pat_let39_0);
          }(_1121_v117)));
          _nw188[(new BigNumber(19)).toNumber()] = _1121_v117;
          _1132_v127 = _nw188;
          let _index202 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1132_v127).length));
          (_1132_v127)[_index202] = _1121_v117;
          let _1141_v128;
          let _init30 = function (_1142_i13) {
            return (_this).f15;
          };
          let _nw189 = Array((_dafny.ONE).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw189.length); _i0_30++) {
            _nw189[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1141_v128 = _nw189;
          let _index203 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1141_v128).length));
          (_1141_v128)[_index203] = (_this).f15;
          let _1143_v129;
          _1143_v129 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(p1, new BigNumber((_1100_v98).length), _this.f14, _1128_v123, globalState),p1);
          let _1144_v130;
          _1144_v130 = _dafny.Set.fromElements(_1118_v114);
          let _1145_v133;
          _1145_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1100_v98,false);
          let _pat_let_tv33 = _1143_v129;
          let _pat_let_tv34 = _1125_v120;
          let _pat_let_tv35 = p1;
          let _index204 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1132_v127).length));
          let _index205 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1141_v128).length));
          let _rhs168 = function (_pat_let41_0) {
            return function (_1146_dt__update__tmp_h6) {
              return function (_pat_let42_0) {
                return function (_1147_dt__update_hcf19_h2) {
                  return function (_pat_let43_0) {
                    return function (_1148_dt__update_hcf18_h1) {
                      return _module.D3.create_DC11((_1146_dt__update__tmp_h6).dtor_cf17, _1148_dt__update_hcf18_h1, _1147_dt__update_hcf19_h2, (_1146_dt__update__tmp_h6).dtor_cf20);
                    }(_pat_let43_0);
                  }(_pat_let_tv35);
                }(_pat_let42_0);
              }(!(_pat_let_tv33).contains(_pat_let_tv34));
            }(_pat_let41_0);
          }(_1121_v117);
          let _rhs169 = (_this).f15;
          let _rhs170 = (((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.__default.fm23(new BigNumber(((_1122_v118).f21).length), _module.__default.fm23((_1122_v118).f20, (_this).f19, _1143_v129, _1144_v130, globalState), function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of ((_1122_v118).f21).Elements) {
              let _1149_v131 = _compr_45;
              if (_dafny.Seq.contains((_1122_v118).f21, _1149_v131)) {
                _coll45.push([_1149_v131,p1]);
              }
            }
            return _coll45;
          }(), function () {
            let _coll46 = new _dafny.Set();
            for (const _compr_46 of (_dafny.MultiSet.fromElements(_1118_v114, _dafny.Seq.of(new BigNumber(-546)))).Elements) {
              let _1150_v132 = _compr_46;
              if ((_dafny.MultiSet.fromElements(_1118_v114, _dafny.Seq.of(new BigNumber(-546)))).contains(_1150_v132)) {
                _coll46.add(_1150_v132);
              }
            }
            return _coll46;
          }(), globalState), _this.f14),(_this).f19)).update(_dafny.Seq.of((_this).f19), (_this).f19)).update(_1100_v98, (_this).f19)).equals(_1145_v133);
          let _lhs142 = _1132_v127;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1132_v127).length));
          let _lhs144 = _1141_v128;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1141_v128).length));
          let _lhs146 = globalState;
          _lhs142[_lhs143] = _rhs168;
          _lhs144[_lhs145] = _rhs169;
          _lhs146.f9 = _rhs170;
          let _1151_v134;
          _1151_v134 = _dafny.Seq.of((_1131_v126).Merge(_1131_v126), _1131_v126, _dafny.Map.Empty.slice().updateUnsafe(true,false), _1131_v126);
          _1151_v134 = _dafny.Seq.Concat(_1151_v134, _1151_v134);
        }
        let _1152_v135;
        _1152_v135 = _module.D7.create_DC23((_this).f19, p1, _this.f14);
        let _1153_v136;
        _1153_v136 = _dafny.Seq.of(_1100_v98, _module.__default.fm38(globalState), _1100_v98);
        (globalState).f3 = ((_1152_v135).dtor_cf43).multipliedBy(new BigNumber((_1153_v136).length));
        (globalState).f3 = p1;
        let _1154_v137;
        let _nw190 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1154_v137 = _nw190;
        let _index206 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1154_v137).length));
        (_1154_v137)[_index206] = (_dafny.ZERO).minus(new BigNumber((_1100_v98).length));
        let _1155_v138;
        _1155_v138 = _dafny.Set.fromElements((_this).f19, _this.f14);
        let _index207 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1154_v137).length));
        let _rhs171 = (_dafny.ZERO).minus(p1);
        let _rhs172 = (_1155_v138).IsProperSubsetOf(_1155_v138);
        let _rhs173 = _module.__default.safeDivisionInt(p1, p1);
        let _lhs147 = globalState;
        let _lhs148 = _this;
        let _lhs149 = _1154_v137;
        let _lhs150 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1154_v137).length));
        _lhs147.f3 = _rhs171;
        _lhs148.f14 = _rhs172;
        _lhs149[_lhs150] = _rhs173;
      } else {
        let _1156_v139;
        _1156_v139 = _dafny.MultiSet.fromElements(((_this).f19) || ((_this).f19));
        let _1157_v140;
        let _nw191 = Array((new BigNumber(20)).toNumber()).fill([]);
        _1157_v140 = _nw191;
        let _1158_v141;
        _1158_v141 = _dafny.Seq.of((_this).f19);
        let _1159_v142;
        _1159_v142 = _module.D0.create_DC1(_this.f14);
        let _1160_v143;
        let _nw192 = Array((new BigNumber(12)).toNumber());
        _nw192[(_dafny.ZERO).toNumber()] = _this.f14;
        _nw192[(_dafny.ONE).toNumber()] = _this.f14;
        _nw192[(new BigNumber(2)).toNumber()] = _this.f14;
        _nw192[(new BigNumber(3)).toNumber()] = _this.f14;
        _nw192[(new BigNumber(4)).toNumber()] = !(_this.f14);
        _nw192[(new BigNumber(5)).toNumber()] = (_1158_v141)[_module.__default.safeIndex(p1, new BigNumber((_1158_v141).length))];
        _nw192[(new BigNumber(6)).toNumber()] = (_this).f19;
        _nw192[(new BigNumber(7)).toNumber()] = _this.f14;
        _nw192[(new BigNumber(8)).toNumber()] = (_1159_v142).dtor_cf1;
        _nw192[(new BigNumber(9)).toNumber()] = (_this).f19;
        _nw192[(new BigNumber(10)).toNumber()] = (_this).f19;
        _nw192[(new BigNumber(11)).toNumber()] = _this.f14;
        _1160_v143 = _nw192;
        let _index208 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1157_v140).length));
        (_1157_v140)[_index208] = _1160_v143;
        let _index209 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1157_v140).length));
        let _rhs174 = (_1156_v139).Union((_this).f15);
        let _rhs175 = ((false) ? (_1160_v143) : (_1160_v143));
        let _lhs151 = _1157_v140;
        let _lhs152 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1157_v140).length));
        _1156_v139 = _rhs174;
        _lhs151[_lhs152] = _rhs175;
        let _index210 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length));
        (_1160_v143)[_index210] = (_this).f19;
        let _1161_v144;
        _1161_v144 = _dafny.Set.fromElements((_this).f19, (_this).f19, false);
        let _index211 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length));
        (_1160_v143)[_index211] = !((new BigNumber(70)).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1161_v144).length))))) || (!(_this.f14));
        let _1162_v145;
        _1162_v145 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(709),p1);
        let _1163_v147;
        _1163_v147 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1057_v63).length),(_this).f19);
        let _1164_v148;
        _1164_v148 = _dafny.Seq.of((function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of (p0).Elements) {
            let _1165_v146 = _compr_47;
            if ((p0).contains(_1165_v146)) {
              _coll47.push([(_1165_v146).plus(p1),p1]);
            }
          }
          return _coll47;
        }()).update(new BigNumber(889), new BigNumber((_1163_v147).length)), _1162_v145);
        let _rhs176 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1057_v63).length)), p1);
        let _rhs177 = (_1162_v145).Merge((_1164_v148)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(581), new BigNumber(439))) {
            let _1166_v149 = _compr_48;
            if (((new BigNumber(581)).isLessThanOrEqualTo(_1166_v149)) && ((_1166_v149).isLessThan(new BigNumber(439)))) {
              _coll48.push([_module.__default.safeDivisionInt(_1166_v149, new BigNumber((_1161_v144).length)),(_this).f19]);
            }
          }
          return _coll48;
        }()).length), new BigNumber((_1164_v148).length))]);
        let _rhs178 = _1057_v63;
        let _lhs153 = globalState;
        _lhs153.f3 = _rhs176;
        _1162_v145 = _rhs177;
        _1057_v63 = _rhs178;
        let _1167_v150;
        _1167_v150 = _dafny.MultiSet.fromElements(p1, p1);
        let _1168_v151;
        _1168_v151 = _module.D1.create_DC4(_module.__default.fm0((_1160_v143)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length))], p1, p1, _1167_v150, globalState), p1);
        let _1169_v153;
        _1169_v153 = _module.D3.create_DC11((_1168_v151).dtor_cf5, new BigNumber((function () {
  let _coll49 = new _dafny.Map();
  for (const _compr_49 of (_1167_v150).Elements) {
    let _1170_v152 = _compr_49;
    if ((_1167_v150).contains(_1170_v152)) {
      _coll49.push([(_1170_v152).multipliedBy(p1),_1162_v145]);
    }
  }
  return _coll49;
}()).length), _this.f14, p1);
        let _source13 = _1169_v153;
        if (_source13.is_DC11) {
          let _1171___mcc_h3 = (_source13).cf17;
          let _1172___mcc_h4 = (_source13).cf18;
          let _1173___mcc_h5 = (_source13).cf19;
          let _1174___mcc_h6 = (_source13).cf20;
          let _1175_cf20 = _1174___mcc_h6;
          let _1176_cf19 = _1173___mcc_h5;
          let _1177_cf18 = _1172___mcc_h4;
          let _1178_cf17 = _1171___mcc_h3;
          let _1179_v154;
          _1179_v154 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1057_v63, _1057_v63),_1162_v145);
          let _1180_v155;
          _1180_v155 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1181_v156;
          _1181_v156 = _module.D9.create_DC29(_1162_v145);
          _1179_v154 = (_1179_v154).update(_dafny.Seq.update(_1057_v63, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(345)), function (_1182_i14) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length), new BigNumber((_1057_v63).length)), _1180_v155), (_1181_v156).dtor_cf49);
          r2 = (((_1163_v147).contains((_1175_cf20).multipliedBy(_1178_cf17))) ? ((_1163_v147).get((_1175_cf20).multipliedBy(_1178_cf17))) : (_1176_cf19));
          let _1183_v157;
          _1183_v157 = _dafny.Set.fromElements(_1180_v155, _1180_v155);
          let _1184_v158;
          _1184_v158 = _dafny.Seq.of(_1183_v157, (_1183_v157).Difference(_1183_v157));
          _1184_v158 = _1184_v158;
          (_this).f14 = !(_1176_cf19) || (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), ((_1185_v155) => function (_1186_i15) {
            return _1185_v155;
          })(_1180_v155)), _1057_v63));
        } else {
          let _1187___mcc_h7 = (_source13).cf16;
          let _1188_cf16 = _1187___mcc_h7;
          (_this).m6(globalState);
          let _1189_v159;
          let _nw193 = Array((_dafny.ONE).toNumber());
          _nw193[(_dafny.ZERO).toNumber()] = (p1).minus(p1);
          _1189_v159 = _nw193;
          let _index212 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1189_v159).length));
          (_1189_v159)[_index212] = new BigNumber(966);
          let _index213 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_1189_v159).length));
          (_1189_v159)[_index213] = ((p1).plus(_module.__default.fm0(_this.f14, p1, p1, _module.__default.fm30((_this).f19, globalState), globalState))).minus(p1);
          r2 = true;
          let _index214 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length));
          let _rhs179 = ((_1168_v151).dtor_cf4).isLessThanOrEqualTo((_1189_v159)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_1189_v159).length))]);
          let _rhs180 = _1160_v143;
          let _rhs181 = new BigNumber(-942);
          let _rhs182 = (p1).multipliedBy((_1189_v159)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_1189_v159).length))]);
          let _lhs154 = _1160_v143;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length));
          let _lhs156 = globalState;
          let _lhs157 = globalState;
          _lhs154[_lhs155] = _rhs179;
          _lhs156.f2 = _rhs180;
          r0 = _rhs181;
          _lhs157.f3 = _rhs182;
        }
        let _1190_v160;
        _1190_v160 = _dafny.Map.Empty.slice().updateUnsafe((_1157_v140)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_1157_v140).length))],p1);
        let _1191_v161;
        _1191_v161 = _module.D4.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_1192_i16) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), new BigNumber(786), (_this).f19, (_1157_v140)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_1157_v140).length))]);
        let _1193_v162;
        _1193_v162 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_this.f14, (_this).f19));
        let _rhs183 = (new BigNumber((_module.__default.fm30((_1160_v143)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_1160_v143).length))], globalState)).cardinality())).plus(p1);
        let _rhs184 = (_dafny.ZERO).minus(((!((_dafny.ZERO).minus(new BigNumber((_1190_v160).length))).isEqualTo(p1)) ? (_module.__default.safeDivisionInt(p1, p1)) : (p1)));
        let _rhs185 = ((p1).multipliedBy(p1)).minus((_dafny.ZERO).minus(p1));
        let _rhs186 = (p1).plus(((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_1191_v161).dtor_cf22).length)))).multipliedBy((_dafny.ZERO).minus(p1)));
        let _rhs187 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).fm17(p1, _this.f14, _1193_v162, p1, globalState), p1), p1);
        let _lhs158 = globalState;
        r1 = _rhs183;
        _lhs158.f3 = _rhs184;
        r1 = _rhs185;
        r0 = _rhs186;
        r1 = _rhs187;
      }
      let _1194_v163;
      let _nw194 = Array((new BigNumber(15)).toNumber()).fill(false);
      _1194_v163 = _nw194;
      let _1195_v164;
      _1195_v164 = _dafny.Seq.of(_1194_v163);
      let _1196_v165;
      _1196_v165 = _module.D4.create_DC13(_dafny.Seq.UnicodeFromString("uoqpauedn"), p1, _this.f14, (_1195_v164)[_module.__default.safeIndex(p1, new BigNumber((_1195_v164).length))]);
      let _source14 = function (_pat_let44_0) {
        return function (_1197_dt__update__tmp_h7) {
          return function (_pat_let45_0) {
            return function (_1199_dt__update_hcf22_h0) {
              return _module.D4.create_DC13(_1199_dt__update_hcf22_h0, (_1197_dt__update__tmp_h7).dtor_cf23, (_1197_dt__update__tmp_h7).dtor_cf24, (_1197_dt__update__tmp_h7).dtor_cf25);
            }(_pat_let45_0);
          }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), function (_1198_i17) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }));
        }(_pat_let44_0);
      }(_1196_v165);
      if (_source14.is_DC13) {
        let _1200___mcc_h8 = (_source14).cf22;
        let _1201___mcc_h9 = (_source14).cf23;
        let _1202___mcc_h10 = (_source14).cf24;
        let _1203___mcc_h11 = (_source14).cf25;
        let _1204_cf25 = _1203___mcc_h11;
        let _1205_cf24 = _1202___mcc_h10;
        let _1206_cf23 = _1201___mcc_h9;
        let _1207_cf22 = _1200___mcc_h8;
        let _1208_v166;
        _1208_v166 = _dafny.Map.Empty.slice().updateUnsafe(_1206_cf23,_1205_cf24);
        let _1209_v167;
        _1209_v167 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1057_v63).length),_1208_v166);
        _1209_v167 = (_1209_v167).update(_module.__default.safeDivisionInt(_1206_cf23, _1206_cf23), _1208_v166);
        let _index215 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length));
        (_1204_cf25)[_index215] = _1205_cf24;
        let _index216 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length));
        (_1204_cf25)[_index216] = (_this).fm2((_this).f19, _this.f14, globalState);
        if ((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))]) {
          let _1210_v168;
          let _init31 = ((_1211_p1) => function (_1212_i18) {
            return (_1212_i18).plus(_1211_p1);
          })(p1);
          let _nw195 = Array((new BigNumber(17)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw195.length); _i0_31++) {
            _nw195[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1210_v168 = _nw195;
          let _1213_v169;
          _1213_v169 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1210_v168);
          (globalState).f3 = (_dafny.ZERO).minus(new BigNumber((((_1213_v169).Merge(_1213_v169)).Merge(_1213_v169)).length));
          let _1214_v170;
          let _nw196 = new _module.C3();
          _nw196.__ctor(p1, _dafny.Seq.update(_1057_v63, _module.__default.safeIndex(_1206_cf23, new BigNumber((_1057_v63).length)), new _dafny.CodePoint('t'.codePointAt(0))), false, (_this).f15);
          _1214_v170 = _nw196;
          _1214_v170 = _1214_v170;
          let _1215_v171;
          let _nw197 = new _module.C3();
          _nw197.__ctor((p1).plus((_1214_v170).f20), _dafny.Seq.Concat(_1057_v63, _1207_cf22), false, (_this).f15);
          _1215_v171 = _nw197;
          (globalState).f3 = new BigNumber(((_this).f15).cardinality());
          let _1216_v172;
          _1216_v172 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1217_v173;
          _1217_v173 = _module.D9.create_DC30(_1216_v172, _1205_cf24, p1, p0, (_1215_v171).f20);
          let _1218_v174;
          _1218_v174 = _dafny.Seq.of(p1);
          let _1219_v175;
          _1219_v175 = _module.D2.create_DC9((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))], (_1215_v171).f20, _1216_v172);
          let _1220_v176;
          _1220_v176 = _dafny.Seq.of(_this.f14);
          let _1221_v177;
          _1221_v177 = _dafny.Set.fromElements(_1057_v63);
          let _pat_let_tv36 = _1216_v172;
          let _1222_v179;
          let _nw198 = Array((new BigNumber(21)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = _module.D9.create_DC30(_1216_v172, (_this).f19, _1206_cf23, p0, p1);
          _nw198[(_dafny.ONE).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(2)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(3)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(4)).toNumber()] = function (_pat_let46_0) {
            return function (_1223_dt__update__tmp_h8) {
              return function (_pat_let47_0) {
                return function (_1224_dt__update_hcf50_h0) {
                  return _module.D9.create_DC30(_1224_dt__update_hcf50_h0, (_1223_dt__update__tmp_h8).dtor_cf51, (_1223_dt__update__tmp_h8).dtor_cf52, (_1223_dt__update__tmp_h8).dtor_cf53, (_1223_dt__update__tmp_h8).dtor_cf54);
                }(_pat_let47_0);
              }(_pat_let_tv36);
            }(_pat_let46_0);
          }(_1217_v173);
          _nw198[(new BigNumber(5)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(6)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(7)).toNumber()] = _module.D9.create_DC30(_1216_v172, _1205_cf24, (_1214_v170).f20, p0, new BigNumber(746));
          _nw198[(new BigNumber(8)).toNumber()] = _module.D9.create_DC30(_1216_v172, (_this).f19, (_1215_v171).f20, p0, p1);
          _nw198[(new BigNumber(9)).toNumber()] = _module.D9.create_DC30(_1216_v172, false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1218_v174, _1218_v174)).length)), _dafny.Set.fromElements((_1215_v171).f20, new BigNumber(5)), (_1215_v171).f20);
          _nw198[(new BigNumber(10)).toNumber()] = _module.D9.create_DC30((_1219_v175).dtor_cf15, (_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))], (_1214_v170).f20, p0, (_1214_v170).f20);
          _nw198[(new BigNumber(11)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(12)).toNumber()] = _module.__default.fm39(new BigNumber((_dafny.Seq.update(_1220_v176, _module.__default.safeIndex(p1, new BigNumber((_1220_v176).length)), (_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))])).length), (_1214_v170).f20, p1, new BigNumber((_1221_v177).length), globalState);
          _nw198[(new BigNumber(13)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(14)).toNumber()] = _module.D9.create_DC30(_1216_v172, (_this).f19, (_1214_v170).f20, p0, p1);
          _nw198[(new BigNumber(15)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(16)).toNumber()] = _module.D9.create_DC30(_1216_v172, (_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))], new BigNumber((function () {
  let _coll50 = new _dafny.Map();
  for (const _compr_50 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-580)), ((_1225_v172) => function (_1226_i19) {
    return _1225_v172;
  })(_1216_v172))).Elements) {
    let _1227_v178 = _compr_50;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-580)), ((_1228_v172) => function (_1226_i19) {
      return _1228_v172;
    })(_1216_v172)), _1227_v178)) {
      _coll50.push([_1227_v178,false]);
    }
  }
  return _coll50;
}()).length), _dafny.Set.fromElements(new BigNumber(((_1214_v170).f21).length)), (_1215_v171).f20);
          _nw198[(new BigNumber(17)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(18)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(19)).toNumber()] = _1217_v173;
          _nw198[(new BigNumber(20)).toNumber()] = _1217_v173;
          _1222_v179 = _nw198;
          let _1229_v180;
          _1229_v180 = _dafny.Map.Empty.slice().updateUnsafe(_1222_v179,(_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))]);
          let _rhs188 = (((_1229_v180).contains(_1222_v179)) ? ((_1229_v180).get(_1222_v179)) : (!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_1230_cf24) => function (_1231_i20) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1230_cf24,_1230_cf24)).length);
          })(_1205_cf24)), (_module.D1.create_DC4((_1214_v170).f20, (_1214_v170).f20)).dtor_cf4)));
          let _rhs189 = (_1215_v171).fm2(false, _1205_cf24, globalState);
          r2 = _rhs188;
          _1205_cf24 = _rhs189;
        } else {
          (_this).f14 = !((p1).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qd")).length))).isEqualTo(_1206_cf23);
          let _1232_v181;
          _1232_v181 = _module.D4.create_DC14(_this.f14, (_this).f19);
          let _1233_v182;
          _1233_v182 = _module.D4.create_DC15(_1232_v181);
          let _1234_v183;
          _1234_v183 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v182,_1206_cf23);
          let _1235_v184;
          _1235_v184 = _dafny.Set.fromElements(_this.f14, _this.f14);
          let _1236_v185;
          _1236_v185 = _dafny.Map.Empty.slice().updateUnsafe((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))],_1235_v184);
          let _1237_v186;
          _1237_v186 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1234_v183).update(_module.D4.create_DC15(_1232_v181), new BigNumber((_1236_v185).length)));
          _1237_v186 = (_1237_v186).update(_1206_cf23, _dafny.Map.Empty.slice().updateUnsafe(_1233_v182,p1));
          let _1238_v187;
          _1238_v187 = _dafny.Seq.of((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))]);
          let _1239_v188;
          _1239_v188 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1057_v63, (_this).fm3(new BigNumber((_1238_v187).length), _1206_cf23, !((_this).f19), _1206_cf23, globalState)));
          _1239_v188 = _dafny.MultiSet.fromElements(_1207_cf22);
          let _1240_v189;
          _1240_v189 = _dafny.Seq.of(new BigNumber((_1057_v63).length), p1);
          let _1241_v190;
          _1241_v190 = _dafny.Map.Empty.slice().updateUnsafe(_1204_cf25,new BigNumber((_1240_v189).length));
          let _1242_v191;
          _1242_v191 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1243_v192;
          _1243_v192 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v191,_1206_cf23);
          let _1244_v193;
          _1244_v193 = _dafny.Set.fromElements(_1240_v189);
          let _1245_v194;
          _1245_v194 = _dafny.Seq.of(_1244_v193);
          let _index217 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length));
          (_1204_cf25)[_index217] = !(_module.__default.fm23(p1, (((_1208_v166).contains(new BigNumber((_1241_v190).length))) ? ((_1208_v166).get(new BigNumber((_1241_v190).length))) : (_this.f14)), _1243_v192, (_1245_v194)[_module.__default.safeIndex(new BigNumber(940), new BigNumber((_1245_v194).length))], globalState)) || ((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))]);
          let _1246_v195;
          let _nw199 = new _module.C2();
          _nw199.__ctor((_1208_v166).Merge(_1208_v166));
          _1246_v195 = _nw199;
        }
        let _1247_v196;
        _1247_v196 = _dafny.Map.Empty.slice().updateUnsafe(_1205_cf24,(_this).f15);
        let _1248_v197;
        let _nw200 = new _module.C3();
        _nw200.__ctor(p1, _dafny.Seq.Concat(_1207_cf22, _1207_cf22), _1205_cf24, (((_1247_v196).contains((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))])) ? ((_1247_v196).get((_1204_cf25)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1204_cf25).length))])) : ((_this).f15)));
        _1248_v197 = _nw200;
      } else if (_source14.is_DC14) {
        let _1249___mcc_h12 = (_source14).cf26;
        let _1250___mcc_h13 = (_source14).cf27;
        let _1251_cf27 = _1250___mcc_h13;
        let _1252_cf26 = _1249___mcc_h12;
        let _1253_v198;
        _1253_v198 = _dafny.MultiSet.fromElements(_module.__default.fm0(_this.f14, p1, new BigNumber((_1057_v63).length), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-852)), function (_1254_i21) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length)), globalState), new BigNumber((_1057_v63).length), new BigNumber(869));
        let _1255_v199;
        _1255_v199 = _dafny.Seq.of(_1253_v198, _dafny.MultiSet.fromElements(p1));
        r1 = _module.__default.fm0(_1252_cf26, new BigNumber((_1057_v63).length), p1, ((_1255_v199)[_module.__default.safeIndex(p1, new BigNumber((_1255_v199).length))]).Intersect(_1253_v198), globalState);
        let _1256_v200;
        _1256_v200 = _dafny.Seq.of((_this).f15);
        let _1257_v201;
        _1257_v201 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p1, p1)).length)), p1);
        (_this).m7(_dafny.Seq.Concat(_1256_v200, _dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), function (_1258_i22) {
          return (_this).f15;
        })), (new BigNumber((_1257_v201).length)).isLessThanOrEqualTo(p1), new BigNumber(-452), _1251_cf27, globalState);
        (_this).f14 = _dafny.Seq.IsProperPrefixOf(_1057_v63, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vlgqwvqgo"), _1057_v63));
        (globalState).f3 = p1;
      } else if (_source14.is_DC12) {
        let _1259___mcc_h14 = (_source14).cf21;
        let _1260_cf21 = _1259___mcc_h14;
        let _1261_v202;
        _1261_v202 = _dafny.Set.fromElements(_this.f14);
        let _1262_v203;
        _1262_v203 = _dafny.Seq.of(new BigNumber((_1261_v202).length), p1);
        let _1263_v204;
        _1263_v204 = _module.D1.create_DC3(new BigNumber((_1057_v63).length));
        let _1264_v205;
        _1264_v205 = _1262_v203;
        let _1265_v206;
        _1265_v206 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f19);
        let _1266_v207;
        let _nw201 = new _module.C2();
        _nw201.__ctor(_1265_v206);
        _1266_v207 = _nw201;
        let _1267_v208;
        _1267_v208 = _dafny.Map.Empty.slice().updateUnsafe(_1266_v207,_1262_v203);
        let _1268_v209;
        _1268_v209 = _dafny.Seq.of((_this).f19, _this.f14);
        let _1269_v210;
        let _nw202 = Array((new BigNumber(28)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), function (_1270_i23) {
          return new BigNumber(681);
        });
        _nw202[(_dafny.ONE).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(2)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(p1);
        _nw202[(new BigNumber(4)).toNumber()] = _module.__default.fm21(_1263_v204, _this.f14, _module.__default.fm40(globalState), globalState);
        _nw202[(new BigNumber(5)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(6)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(7)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(8)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(9)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(10)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(11)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(12)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(13)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(14)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(new BigNumber(189));
        _nw202[(new BigNumber(16)).toNumber()] = (_1264_v205);
        _nw202[(new BigNumber(17)).toNumber()] = (((_1267_v208).contains(_1266_v207)) ? ((_1267_v208).get(_1266_v207)) : (_1262_v203));
        _nw202[(new BigNumber(18)).toNumber()] = (((_this).f19) ? (_1262_v203) : (_1262_v203));
        _nw202[(new BigNumber(19)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(20)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_1262_v203, _1262_v203);
        _nw202[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(p1, p1);
        _nw202[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-99)), function (_1271_i24) {
          return new BigNumber((_dafny.Seq.of(!((_this).f19), (_this).f19, _this.f14, (_this).f19, (_this).f19)).length);
        });
        _nw202[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_1262_v203, _dafny.Seq.update(_1262_v203, _module.__default.safeIndex(new BigNumber((_1268_v209).length), new BigNumber((_1262_v203).length)), new BigNumber(790)));
        _nw202[(new BigNumber(25)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(26)).toNumber()] = _1262_v203;
        _nw202[(new BigNumber(27)).toNumber()] = _1262_v203;
        _1269_v210 = _nw202;
        _1269_v210 = _1269_v210;
        let _1272_v211;
        _1272_v211 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1261_v202);
        let _1273_v212;
        let _nw203 = Array((new BigNumber(9)).toNumber());
        _nw203[(_dafny.ZERO).toNumber()] = _1261_v202;
        _nw203[(_dafny.ONE).toNumber()] = _1261_v202;
        _nw203[(new BigNumber(2)).toNumber()] = _module.__default.fm41(new BigNumber((_1262_v203).length), p1, globalState);
        _nw203[(new BigNumber(3)).toNumber()] = (((_1272_v211).contains(new BigNumber(996))) ? ((_1272_v211).get(new BigNumber(996))) : (_1261_v202));
        _nw203[(new BigNumber(4)).toNumber()] = (_1261_v202).Intersect(_1261_v202);
        _nw203[(new BigNumber(5)).toNumber()] = _1261_v202;
        _nw203[(new BigNumber(6)).toNumber()] = (_1261_v202).Union(_dafny.Set.fromElements(false, false));
        _nw203[(new BigNumber(7)).toNumber()] = _1261_v202;
        _nw203[(new BigNumber(8)).toNumber()] = _1261_v202;
        _1273_v212 = _nw203;
        let _1274_v213;
        _1274_v213 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1057_v63).length)), p1);
        let _1275_v215;
        _1275_v215 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), ((_1276_p1) => function (_1277_i25) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(366),new BigNumber((function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(210), new BigNumber(972))) {
              let _1278_v214 = _compr_51;
              if (((new BigNumber(210)).isLessThanOrEqualTo(_1278_v214)) && ((_1278_v214).isLessThan(new BigNumber(972)))) {
                _coll51.push([(_1278_v214).multipliedBy(_1276_p1),true]);
              }
            }
            return _coll51;
          }()).length))).length);
        })(p1)));
        let _index218 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1273_v212).length));
        (_1273_v212)[_index218] = _dafny.Set.fromElements(!(true), _module.__default.fm23(_module.__default.fm0(false, new BigNumber(279), p1, _dafny.MultiSet.fromElements((((_1274_v213).contains(p1)) ? ((_1274_v213).get(p1)) : ((_dafny.ZERO).minus(p1))), new BigNumber((_1268_v209).length), p1), globalState), _this.f14, _module.__default.fm24(p1, (_this).f19, false, p1, globalState), _1275_v215, globalState), (_this).f19, (_this).f19, (_this).f19);
        let _index219 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1273_v212).length));
        let _rhs190 = p1;
        let _rhs191 = !(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1057_v63, _dafny.Seq.UnicodeFromString("nhxiqa")), _1057_v63)));
        let _rhs192 = ((_this.f14) ? (_dafny.Set.fromElements(!(_this.f14), _this.f14, _this.f14, (_this).f19, (_this).f19)) : ((_1261_v202).Difference(_1261_v202)));
        let _lhs159 = globalState;
        let _lhs160 = _1273_v212;
        let _lhs161 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1273_v212).length));
        _lhs159.f3 = _rhs190;
        r2 = _rhs191;
        _lhs160[_lhs161] = _rhs192;
        let _1279_v216;
        let _nw204 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1279_v216 = _nw204;
        let _1280_v217;
        let _nw205 = new _module.C3();
        _nw205.__ctor(new BigNumber(6), _1057_v63, (_this).f19, (_this).f15);
        _1280_v217 = _nw205;
        let _1281_v218;
        _1281_v218 = _dafny.Map.Empty.slice().updateUnsafe(_1279_v216,_1280_v217);
        _1281_v218 = (_1281_v218).update((((_this).f19) ? (_1279_v216) : (_1279_v216)), _1280_v217);
        let _1282_v219;
        _1282_v219 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_1280_v217).f20),(_this).f19);
        let _1283_v220;
        let _init32 = ((_1284_p1, _1285_v217) => function (_1286_i26) {
          return (_1286_i26).minus((_module.D1.create_DC4((_dafny.ZERO).minus(_1284_p1), (_dafny.ZERO).minus((_dafny.ZERO).minus((_1285_v217).f20)))).dtor_cf4);
        })(p1, _1280_v217);
        let _nw206 = Array((new BigNumber(14)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw206.length); _i0_32++) {
          _nw206[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1283_v220 = _nw206;
        let _1287_v221;
        _1287_v221 = _dafny.Map.Empty.slice().updateUnsafe(_1282_v219,_1283_v220);
        _1287_v221 = (_1287_v221).update(_1282_v219, _1283_v220);
      } else {
        let _1288___mcc_h15 = (_source14).cf28;
        let _1289_cf28 = _1288___mcc_h15;
        if ((new BigNumber(595)).isLessThan(p1)) {
          let _index220 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length));
          (_1194_v163)[_index220] = !(_this.f14) || ((_this).f19);
          let _index221 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length));
          (_1194_v163)[_index221] = (_this).f19;
          let _1290_v222;
          _1290_v222 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1291_v223;
          _1291_v223 = _dafny.Map.Empty.slice().updateUnsafe(_1290_v222,(_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))]);
          _1291_v223 = (_1291_v223).update(_1290_v222, (_this).f19);
          let _1292_v224;
          _1292_v224 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!((_this).f19)),new BigNumber((_1195_v164).length));
          let _1293_v226;
          _1293_v226 = _dafny.Seq.of((_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))]);
          let _1294_v227;
          _1294_v227 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1057_v63);
          let _1295_v228;
          _1295_v228 = _module.D8.create_DC26((_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))], _1294_v227);
          let _1296_v229;
          _1296_v229 = _dafny.MultiSet.fromElements(_module.__default.fm42(false, !((_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))]), new BigNumber(265), (_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))], globalState), _1295_v228);
          let _1297_v230;
          _1297_v230 = _module.D3.create_DC11(p1, p1, _this.f14, new BigNumber(-278));
          let _1298_v231;
          _1298_v231 = _dafny.MultiSet.fromElements((_1297_v230).dtor_cf18, new BigNumber((_1057_v63).length));
          let _1299_v232;
          _1299_v232 = _dafny.Set.fromElements(_1298_v231, _1298_v231);
          let _index222 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length));
          let _rhs193 = function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_dafny.Seq.of(_1293_v226)).Elements) {
              let _1300_v225 = _compr_52;
              if (_dafny.Seq.contains(_dafny.Seq.of(_1293_v226), _1300_v225)) {
                _coll52.push([_1300_v225,p1]);
              }
            }
            return _coll52;
          }();
          let _rhs194 = !((_1296_v229).IsProperSubsetOf(_1296_v229));
          let _rhs195 = (_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))];
          let _rhs196 = (_1299_v232).IsDisjointFrom(_1299_v232);
          let _lhs162 = globalState;
          let _lhs163 = _this;
          let _lhs164 = _1194_v163;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length));
          _1292_v224 = _rhs193;
          _lhs162.f9 = _rhs194;
          _lhs163.f14 = _rhs195;
          _lhs164[_lhs165] = _rhs196;
          let _1301_v233;
          let _nw207 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1301_v233 = _nw207;
          let _1302_v234;
          let _init33 = ((_1303_p1) => function (_1304_i27) {
            return _module.__default.safeDivisionInt(_1304_i27, _1303_p1);
          })(p1);
          let _nw208 = Array((new BigNumber(28)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw208.length); _i0_33++) {
            _nw208[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1302_v234 = _nw208;
          let _index223 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1301_v233).length));
          (_1301_v233)[_index223] = _1302_v234;
          let _1305_v235;
          _1305_v235 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1290_v222);
          let _1306_v236;
          _1306_v236 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1290_v222);
          let _index224 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1301_v233).length));
          let _rhs197 = p1;
          let _rhs198 = p1;
          let _rhs199 = _1302_v234;
          let _rhs200 = _dafny.Seq.update((_this).fm3((_dafny.ZERO).minus(p1), p1, !((((_1291_v223).contains(_1290_v222)) ? ((_1291_v223).get(_1290_v222)) : (_this.f14))), _module.__default.safeModuloInt(new BigNumber((_1305_v235).length), p1), globalState), _module.__default.safeIndex(p1, new BigNumber(((_this).fm3((_dafny.ZERO).minus(p1), p1, !((((_1291_v223).contains(_1290_v222)) ? ((_1291_v223).get(_1290_v222)) : (_this.f14))), _module.__default.safeModuloInt(new BigNumber((_1305_v235).length), p1), globalState)).length)), (((_1306_v236).contains(!(_this.f14))) ? ((_1306_v236).get(!(_this.f14))) : (_1290_v222)));
          let _lhs166 = globalState;
          let _lhs167 = globalState;
          let _lhs168 = _1301_v233;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_1301_v233).length));
          _lhs166.f3 = _rhs197;
          _lhs167.f3 = _rhs198;
          _lhs168[_lhs169] = _rhs199;
          _1057_v63 = _rhs200;
          let _1307_v237;
          _1307_v237 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_this).f19, (_this).f19, !((_1194_v163)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((_1194_v163).length))]))).length), (_1297_v230).dtor_cf18, p1);
          _1307_v237 = _1307_v237;
        } else {
          let _1308_v238;
          let _init34 = ((_1309_p1) => function (_1310_i28) {
            return _dafny.Seq.of(_1309_p1);
          })(p1);
          let _nw209 = Array((new BigNumber(14)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw209.length); _i0_34++) {
            _nw209[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1308_v238 = _nw209;
          let _1311_v239;
          _1311_v239 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),p1);
          let _1312_v240;
          _1312_v240 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1311_v239);
          let _1313_v241;
          _1313_v241 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).fm2((_this).f19, _this.f14, globalState));
          let _1314_v242;
          _1314_v242 = _dafny.Seq.of(p1, new BigNumber(((((_1312_v240).contains(new BigNumber((_1313_v241).length))) ? ((_1312_v240).get(new BigNumber((_1313_v241).length))) : (_1311_v239))).length), p1, p1, p1);
          let _index225 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1308_v238).length));
          (_1308_v238)[_index225] = ((_this.f14) ? (_1314_v242) : (_1314_v242));
          let _index226 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1308_v238).length));
          (_1308_v238)[_index226] = _1314_v242;
          let _1315_v243;
          _1315_v243 = _dafny.Seq.of(_1057_v63);
          r1 = (((_this).f19) ? (_module.__default.safeDivisionInt(p1, new BigNumber((_1057_v63).length))) : (new BigNumber((_1315_v243).length)));
          let _1316_v244;
          _1316_v244 = _module.D4.create_DC12(_1194_v163);
          let _1317_v245;
          let _nw210 = Array((new BigNumber(4)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = _1316_v244;
          _nw210[(_dafny.ONE).toNumber()] = _1316_v244;
          _nw210[(new BigNumber(2)).toNumber()] = _1316_v244;
          _nw210[(new BigNumber(3)).toNumber()] = _1316_v244;
          _1317_v245 = _nw210;
          let _1318_v246;
          _1318_v246 = _dafny.Map.Empty.slice().updateUnsafe(_1317_v245,true);
          (globalState).f9 = !(_1318_v246).contains(_1317_v245);
          let _1319_v247;
          _1319_v247 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1057_v63).length),_this.f14);
          (_this).f14 = ((new BigNumber((_1319_v247).length)).isEqualTo(p1)) === ((_this).f19);
          let _1320_v248;
          _1320_v248 = _dafny.Seq.of((_this).f19);
          (_this).f14 = !((_1320_v248)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1320_v248).length))]) || ((_this).f19);
        }
        if (!(true)) {
          let _1321_v249;
          _1321_v249 = _dafny.Seq.of(p1, p1);
          let _1322_v250;
          _1322_v250 = _dafny.MultiSet.fromElements(p1);
          let _1323_v251;
          let _nw211 = new _module.C3();
          _nw211.__ctor(_module.__default.fm0((_this).f19, new BigNumber((_1321_v249).length), p1, _1322_v250, globalState), _1057_v63, (p1).isLessThanOrEqualTo(p1), (_this).f15);
          _1323_v251 = _nw211;
          let _1324_v252;
          _1324_v252 = _dafny.Seq.of(_1322_v250, _1322_v250, _1322_v250, _1322_v250);
          (_this).m7(_module.__default.fm43((_1323_v251).f20, p1, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_module.__default.fm0(!((_this).f19), (_1323_v251).f20, p1, _1322_v250, globalState)))).length), globalState), _this.f14, new BigNumber(((_1322_v250).Intersect((_1324_v252)[_module.__default.safeIndex(p1, new BigNumber((_1324_v252).length))])).cardinality()), _this.f14, globalState);
          let _1325_v253;
          _1325_v253 = _module.D2.create_DC8(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-565)), function (_1326_i29) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})).length), _module.__default.safeDivisionInt((_1323_v251).f20, p1), p1);
          let _1327_v254;
          let _nw212 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1327_v254 = _nw212;
          let _index227 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1327_v254).length));
          (_1327_v254)[_index227] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_1323_v251).f20, p1));
          let _index228 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1327_v254).length));
          (_1327_v254)[_index228] = (new BigNumber((_1321_v249).length)).minus(p1);
          let _index229 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1194_v163).length));
          (_1194_v163)[_index229] = true;
          let _1328_v255;
          _1328_v255 = _dafny.Set.fromElements(_1327_v254);
          let _index230 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1327_v254).length));
          let _index231 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1327_v254).length));
          let _index232 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1194_v163).length));
          let _rhs201 = (_this).f19;
          let _rhs202 = _module.__default.fm44((_this).f19, (_1323_v251).f20, (_1323_v251).f20, globalState);
          let _rhs203 = ((_this.f14) ? (p1) : (p1));
          let _rhs204 = (_1323_v251).f20;
          let _rhs205 = ((_1328_v255).Union(_1328_v255)).equals(_dafny.Set.fromElements(_1327_v254));
          let _lhs170 = _1327_v254;
          let _lhs171 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1327_v254).length));
          let _lhs172 = _1327_v254;
          let _lhs173 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1327_v254).length));
          let _lhs174 = _1194_v163;
          let _lhs175 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1194_v163).length));
          r2 = _rhs201;
          _1325_v253 = _rhs202;
          _lhs170[_lhs171] = _rhs203;
          _lhs172[_lhs173] = _rhs204;
          _lhs174[_lhs175] = _rhs205;
          let _1329_v256;
          _1329_v256 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(560));
          let _index233 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1327_v254).length));
          (_1327_v254)[_index233] = ((_1323_v251).f20).multipliedBy((((_1329_v256).contains(_this.f14)) ? ((_1329_v256).get(_this.f14)) : ((_1323_v251).f20)));
          (globalState).f3 = new BigNumber((_dafny.Seq.UnicodeFromString("owt")).length);
        } else {
          let _1330_v257;
          let _nw213 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1330_v257 = _nw213;
          _1330_v257 = _1330_v257;
          let _1331_v258;
          _1331_v258 = _module.D8.create_DC26((_this).f19, _dafny.Map.Empty.slice().updateUnsafe(p1,_1057_v63));
          let _1332_v259;
          let _nw214 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1332_v259 = _nw214;
          let _index234 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length));
          (_1332_v259)[_index234] = (((_this).f19) ? (p1) : (p1));
          let _1333_v260;
          _1333_v260 = _module.D4.create_DC14(_this.f14, (_this).fm2(_this.f14, !(_this.f14), globalState));
          let _1334_v261;
          _1334_v261 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(p1, p1, (_1333_v260).dtor_cf26, p1, globalState),_1332_v259);
          let _1335_v262;
          _1335_v262 = _dafny.Seq.of(_1057_v63, _1057_v63, _1057_v63);
          let _1336_v263;
          _1336_v263 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1335_v262)[_module.__default.safeIndex(p1, new BigNumber((_1335_v262).length))]);
          let _pat_let_tv37 = _1336_v263;
          let _index235 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length));
          let _rhs206 = p1;
          let _rhs207 = new BigNumber((_1334_v261).length);
          let _rhs208 = function (_pat_let48_0) {
            return function (_1337_dt__update__tmp_h9) {
              return function (_pat_let49_0) {
                return function (_1338_dt__update_hcf47_h0) {
                  return _module.D8.create_DC26((_1337_dt__update__tmp_h9).dtor_cf46, _1338_dt__update_hcf47_h0);
                }(_pat_let49_0);
              }(_pat_let_tv37);
            }(_pat_let48_0);
          }(_1331_v258);
          let _rhs209 = _module.__default.safeModuloInt((p1).multipliedBy(new BigNumber(((_this).f15).cardinality())), p1);
          let _lhs176 = globalState;
          let _lhs177 = _1332_v259;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length));
          r0 = _rhs206;
          _lhs176.f3 = _rhs207;
          _1331_v258 = _rhs208;
          _lhs177[_lhs178] = _rhs209;
          let _1339_v264;
          _1339_v264 = _module.D2.create_DC9((_this).f19, (_1332_v259)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length))], new _dafny.CodePoint('n'.codePointAt(0)));
          r0 = (_1339_v264).dtor_cf14;
          r2 = !(_this.f14);
          let _1340_v265;
          _1340_v265 = _dafny.Map.Empty.slice().updateUnsafe((_1332_v259)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length))],new BigNumber((p0).length));
          let _1341_v266;
          _1341_v266 = _dafny.Map.Empty.slice().updateUnsafe(_1340_v265,(_1332_v259)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length))]);
          _1341_v266 = (_1341_v266).update((_1340_v265).Merge(_1340_v265), (_1332_v259)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1332_v259).length))]);
        }
        (_this).f14 = !(_this.f14);
        if ((_this).f19) {
          let _1342_v267;
          _1342_v267 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f14);
          let _1343_v268;
          let _nw215 = new _module.C2();
          _nw215.__ctor((_1342_v267).update(p1, (_this).f19));
          _1343_v268 = _nw215;
          let _1344_v269;
          let _nw216 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1344_v269 = _nw216;
          let _1345_v270;
          _1345_v270 = new _dafny.CodePoint('a'.codePointAt(0));
          let _index236 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_1344_v269).length));
          (_1344_v269)[_index236] = _1345_v270;
          let _1346_v271;
          _1346_v271 = _dafny.Set.fromElements((_this).f15, (_this).f15);
          let _index237 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_1344_v269).length));
          let _rhs210 = (new BigNumber(241)).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)));
          let _rhs211 = (_1346_v271).IsSubsetOf(_1346_v271);
          let _rhs212 = (_1057_v63)[_module.__default.safeIndex(p1, new BigNumber((_1057_v63).length))];
          let _lhs179 = globalState;
          let _lhs180 = globalState;
          let _lhs181 = _1344_v269;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_1344_v269).length));
          _lhs179.f9 = _rhs210;
          _lhs180.f9 = _rhs211;
          _lhs181[_lhs182] = _rhs212;
          let _1347_v272;
          _1347_v272 = _dafny.MultiSet.fromElements(p1, p1);
          let _1348_v273;
          _1348_v273 = _dafny.Seq.of(new BigNumber((_1347_v272).cardinality()));
          let _1349_v274;
          _1349_v274 = _dafny.Set.fromElements(_1348_v273, _dafny.Seq.of(p1), _1348_v273);
          let _1350_v275;
          _1350_v275 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f19);
          let _pat_let_tv38 = _1350_v275;
          let _1351_v276;
          _1351_v276 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wwmccp")).length)),((_this.f14) ? (new BigNumber((_dafny.Set.fromElements(p1, p1)).length)) : (new BigNumber(((function (_pat_let50_0) {
            return function (_1352_dt__update__tmp_h10) {
              return function (_pat_let51_0) {
                return function (_1353_dt__update_hcf16_h0) {
                  return _module.D3.create_DC10(_1353_dt__update_hcf16_h0);
                }(_pat_let51_0);
              }(_pat_let_tv38);
            }(_pat_let50_0);
          }(_module.__default.fm45(p1, _module.__default.fm23(p1, true, _dafny.Map.Empty.slice().updateUnsafe(_1345_v270,p1), _1349_v274, globalState), p1, globalState))).dtor_cf16).length))));
          _1351_v276 = (_1351_v276).update(p1, new BigNumber((_1057_v63).length));
          (globalState).f3 = p1;
          (globalState).f9 = (_1347_v272).IsDisjointFrom((_1347_v272).Intersect(_1347_v272));
        } else {
          (globalState).f3 = p1;
          let _1354_v277;
          _1354_v277 = _dafny.Map.Empty.slice().updateUnsafe(false,(p1).multipliedBy(p1));
          _1354_v277 = _1354_v277;
          r1 = p1;
          _1057_v63 = _dafny.Seq.Concat(_1057_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-916)), function (_1355_i30) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }));
          let _1356_v278;
          _1356_v278 = _dafny.MultiSet.fromElements(p1, p1);
          (globalState).f3 = _module.__default.fm0(!((_this).f19) || (_this.f14), p1, (((_this).f19) ? (p1) : (p1)), _1356_v278, globalState);
        }
      }
      r0 = p1;
      r1 = _module.__default.safeDivisionInt(p1, p1);
      r2 = _this.f14;
      return [r0, r1, r2];
    }
    m6(globalState) {
      let _this = this;
      if (_this.f14) {
        let _1357_v0;
        let _init35 = function (_1358_i0) {
          return _module.__default.safeDivisionInt(_1358_i0, new BigNumber((_dafny.Seq.of(_this.f14)).length));
        };
        let _nw217 = Array((new BigNumber(22)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw217.length); _i0_35++) {
          _nw217[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1357_v0 = _nw217;
        let _1359_v1;
        _1359_v1 = new BigNumber(835);
        let _index238 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1357_v0).length));
        (_1357_v0)[_index238] = (_1359_v1).multipliedBy(new BigNumber(698));
        let _1360_v2;
        _1360_v2 = _dafny.Seq.UnicodeFromString("dchmesu");
        let _1361_v3;
        _1361_v3 = _dafny.Seq.of(_1359_v1);
        let _1362_v4;
        _1362_v4 = _dafny.Seq.of(_this.f14);
        let _1363_v5;
        _1363_v5 = _dafny.Set.fromElements(_1359_v1);
        let _index239 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1357_v0).length));
        let _rhs213 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1361_v3)[_module.__default.safeIndex(_1359_v1, new BigNumber((_1361_v3).length))]));
        let _rhs214 = _dafny.Seq.Concat(_1360_v2, _1360_v2);
        let _rhs215 = (_1363_v5).IsProperSubsetOf(_dafny.Set.fromElements(_1359_v1, (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_1359_v1))), new BigNumber((_1362_v4).length)));
        let _rhs216 = new BigNumber(711);
        let _rhs217 = _module.__default.safeModuloInt(_1359_v1, _1359_v1);
        let _lhs183 = _1357_v0;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1357_v0).length));
        let _lhs185 = globalState;
        let _lhs186 = globalState;
        let _lhs187 = globalState;
        _lhs183[_lhs184] = _rhs213;
        _1360_v2 = _rhs214;
        _lhs185.f9 = _rhs215;
        _lhs186.f3 = _rhs216;
        _lhs187.f3 = _rhs217;
        let _1364_v6;
        let _init36 = function (_1365_i1) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14)).contains((_this).f19);
        };
        let _nw218 = Array((new BigNumber(29)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw218.length); _i0_36++) {
          _nw218[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1364_v6 = _nw218;
        (globalState).f2 = _1364_v6;
        let _index240 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1364_v6).length));
        (_1364_v6)[_index240] = _this.f14;
        let _index241 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1364_v6).length));
        (_1364_v6)[_index241] = _this.f14;
        let _1366_v7;
        _1366_v7 = _dafny.Map.Empty.slice().updateUnsafe((_1357_v0)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_1357_v0).length))],false);
        _1363_v5 = (_module.__default.fm22(new BigNumber((_dafny.Seq.UnicodeFromString("epm")).length), _dafny.Seq.of(_1366_v7), (_this).f19, (_this).f19, globalState)).Intersect((_1363_v5).Intersect(_1363_v5));
        let _1367_v8;
        _1367_v8 = new _dafny.CodePoint('h'.codePointAt(0));
        let _rhs218 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1360_v2, _module.__default.safeIndex(_1359_v1, new BigNumber((_1360_v2).length)), _1367_v8), _dafny.Seq.UnicodeFromString("ob"))).length));
        let _rhs219 = (((_1364_v6)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1364_v6).length))]) ? ((_1357_v0)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_1357_v0).length))]) : (new BigNumber((_dafny.Seq.of((_1364_v6)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1364_v6).length))], (_this).f19, _this.f14)).length)));
        let _rhs220 = _1357_v0;
        let _rhs221 = _1360_v2;
        let _lhs188 = globalState;
        _1359_v1 = _rhs218;
        _lhs188.f3 = _rhs219;
        _1357_v0 = _rhs220;
        _1360_v2 = _rhs221;
      } else {
        let _1368_v9;
        _1368_v9 = new BigNumber(819);
        (globalState).f3 = _1368_v9;
        let _1369_v10;
        let _init37 = function (_1370_i2) {
          return (_this).f19;
        };
        let _nw219 = Array((new BigNumber(2)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw219.length); _i0_37++) {
          _nw219[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1369_v10 = _nw219;
        let _1371_v11;
        _1371_v11 = _dafny.Seq.of(_1369_v10, _1369_v10, _1369_v10);
        let _1372_v12;
        _1372_v12 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1368_v9));
        (globalState).f2 = (_1371_v11)[_module.__default.safeIndex(_module.__default.fm0(_this.f14, (_dafny.ZERO).minus(_1368_v9), _1368_v9, _1372_v12, globalState), new BigNumber((_1371_v11).length))];
        let _1373_v13;
        let _nw220 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1373_v13 = _nw220;
        let _nw221 = Array((new BigNumber(8)).toNumber()).fill([]);
        _1373_v13 = _nw221;
        let _index242 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_1369_v10).length));
        (_1369_v10)[_index242] = _this.f14;
        let _index243 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_1369_v10).length));
        (_1369_v10)[_index243] = ((_this).f15).IsProperSubsetOf((_this).f15);
        (_this).f14 = (_1369_v10)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_1369_v10).length))];
      }
      let _1374_v14;
      _1374_v14 = new BigNumber(202);
      let _hi3 = _1374_v14;
      for (let _1375_i3 = _1374_v14; _1375_i3.isLessThan(_hi3); _1375_i3 = _1375_i3.plus(_dafny.ONE)) {
        let _1376_v15;
        _1376_v15 = _dafny.Seq.UnicodeFromString("gmbu");
        let _1377_v16;
        _1377_v16 = _dafny.MultiSet.fromElements((_this).f19);
        let _1378_v17;
        _1378_v17 = _module.D8.create_DC25();
        let _rhs222 = _1376_v15;
        let _rhs223 = _1377_v16;
        let _rhs224 = _this.f14;
        let _rhs225 = _1378_v17;
        let _lhs189 = globalState;
        _1376_v15 = _rhs222;
        _1377_v16 = _rhs223;
        _lhs189.f9 = _rhs224;
        _1378_v17 = _rhs225;
        let _1379_v18;
        let _init38 = ((_1380_i3, _1381_v14) => function (_1382_i4) {
          return !(!(_1380_i3).isEqualTo(_1381_v14));
        })(_1375_i3, _1374_v14);
        let _nw222 = Array((new BigNumber(25)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw222.length); _i0_38++) {
          _nw222[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1379_v18 = _nw222;
        let _index244 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_1379_v18).length));
        (_1379_v18)[_index244] = (_this).f19;
        let _index245 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_1379_v18).length));
        (_1379_v18)[_index245] = _this.f14;
        let _1383_v19;
        _1383_v19 = _dafny.Seq.of((_1379_v18)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_1379_v18).length))]);
        let _1384_v20;
        _1384_v20 = _dafny.MultiSet.fromElements(_module.__default.fm38(globalState));
        _1374_v14 = (_this).fm17(_1375_i3, !((_1379_v18)[_module.__default.safeIndex(new BigNumber(997), new BigNumber((_1379_v18).length))]), (_dafny.MultiSet.fromElements(_1383_v19)).Intersect(_1384_v20), new BigNumber(-899), globalState);
        (globalState).f9 = (_1383_v19)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_1375_i3, _1375_i3), new BigNumber((_1383_v19).length))];
      }
      if ((_this).f19) {
        let _1385_v21;
        let _init39 = function (_1386_i5) {
          return _dafny.Set.fromElements(_this.f14, _this.f14);
        };
        let _nw223 = Array((new BigNumber(24)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw223.length); _i0_39++) {
          _nw223[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1385_v21 = _nw223;
        let _1387_v22;
        _1387_v22 = _dafny.Set.fromElements(_this.f14);
        let _index246 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1385_v21).length));
        (_1385_v21)[_index246] = (_1387_v22).Difference(_1387_v22);
        let _index247 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1385_v21).length));
        (_1385_v21)[_index247] = _1387_v22;
        if (_this.f14) {
          let _1388_v23;
          let _init40 = ((_1389_v14) => function (_1390_i6) {
            return _dafny.Seq.of((_dafny.ZERO).minus(_1389_v14), _1389_v14, _1389_v14);
          })(_1374_v14);
          let _nw224 = Array((new BigNumber(2)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw224.length); _i0_40++) {
            _nw224[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1388_v23 = _nw224;
          let _1391_v24;
          let _nw225 = Array((new BigNumber(24)).toNumber()).fill(_module.D0.Default());
          _1391_v24 = _nw225;
          let _1392_v25;
          _1392_v25 = _module.D0.create_DC1((_this).fm2(_this.f14, true, globalState));
          let _1393_v26;
          _1393_v26 = _module.D0.create_DC2(_1392_v25);
          let _index248 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1391_v24).length));
          (_1391_v24)[_index248] = _1393_v26;
          let _1394_v27;
          let _init41 = function (_1395_i7) {
            return _module.D0.create_DC1(_this.f14);
          };
          let _nw226 = Array((new BigNumber(25)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw226.length); _i0_41++) {
            _nw226[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1394_v27 = _nw226;
          let _1396_v28;
          _1396_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v14,_dafny.Set.fromElements(_this.f14, (_this).f19));
          let _index249 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1394_v27).length));
          (_1394_v27)[_index249] = _module.__default.fm37(_dafny.Set.fromElements(_1374_v14), _1396_v28, _this.f14, (_this).f19, globalState);
          let _1397_v29;
          let _nw227 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1397_v29 = _nw227;
          let _index250 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1397_v29).length));
          (_1397_v29)[_index250] = !((_this).f19);
          let _1398_v30;
          _1398_v30 = _dafny.Set.fromElements(_1374_v14);
          let _1399_v31;
          _1399_v31 = _dafny.Seq.UnicodeFromString("pw");
          let _1400_v32;
          _1400_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(343)), function (_1401_i8) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length),_1399_v31);
          let _1402_v33;
          _1402_v33 = _module.D8.create_DC26(_this.f14, _1400_v32);
          let _pat_let_tv39 = _1402_v33;
          let _index251 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1391_v24).length));
          let _index252 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1394_v27).length));
          let _index253 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1397_v29).length));
          let _rhs226 = _1388_v23;
          let _rhs227 = _1393_v26;
          let _rhs228 = (_1398_v30).IsDisjointFrom(_1398_v30);
          let _rhs229 = function (_pat_let52_0) {
            return function (_1405_dt__update__tmp_h1) {
              return function (_pat_let55_0) {
                return function (_1406_dt__update_hcf1_h1) {
                  return _module.D0.create_DC1(_1406_dt__update_hcf1_h1);
                }(_pat_let55_0);
              }(!((_this).f19));
            }(_pat_let52_0);
          }(function (_pat_let53_0) {
            return function (_1403_dt__update__tmp_h0) {
              return function (_pat_let54_0) {
                return function (_1404_dt__update_hcf1_h0) {
                  return _module.D0.create_DC1(_1404_dt__update_hcf1_h0);
                }(_pat_let54_0);
              }((_pat_let_tv39).dtor_cf46);
            }(_pat_let53_0);
          }(_module.D0.create_DC1((_this).fm2(!((_this).f19), _this.f14, globalState))));
          let _rhs230 = _this.f14;
          let _lhs190 = _1391_v24;
          let _lhs191 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_1391_v24).length));
          let _lhs192 = globalState;
          let _lhs193 = _1394_v27;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_1394_v27).length));
          let _lhs195 = _1397_v29;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1397_v29).length));
          _1388_v23 = _rhs226;
          _lhs190[_lhs191] = _rhs227;
          _lhs192.f9 = _rhs228;
          _lhs193[_lhs194] = _rhs229;
          _lhs195[_lhs196] = _rhs230;
          (globalState).f2 = _1397_v29;
          (_this).f14 = !((_1397_v29)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_1397_v29).length))]);
          let _1407_v34;
          let _nw228 = new _module.C1();
          _nw228.__ctor();
          _1407_v34 = _nw228;
          let _1408_v35;
          _1408_v35 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("uwcabk")).length));
          let _1409_v36;
          _1409_v36 = _dafny.Seq.of(_1408_v35, _1408_v35);
          _1409_v36 = _1409_v36;
        } else {
          let _1410_v37;
          _1410_v37 = _dafny.Seq.UnicodeFromString("nenovha");
          let _1411_v38;
          _1411_v38 = _dafny.Seq.of(_this.f14, (_this).f19, _this.f14, _this.f14, _this.f14);
          let _1412_v39;
          _1412_v39 = new _dafny.CodePoint('y'.codePointAt(0));
          (globalState).f3 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1410_v37, _1410_v37), _dafny.Seq.update((((_1411_v38)[_module.__default.safeIndex(_1374_v14, new BigNumber((_1411_v38).length))]) ? (_1410_v37) : (_1410_v37)), _module.__default.safeIndex((_dafny.ZERO).minus(_1374_v14), new BigNumber(((((_1411_v38)[_module.__default.safeIndex(_1374_v14, new BigNumber((_1411_v38).length))]) ? (_1410_v37) : (_1410_v37))).length)), _1412_v39))).length);
          let _1413_v40;
          _1413_v40 = _dafny.Seq.of(new BigNumber(571));
          let _1414_v41;
          _1414_v41 = _1413_v40;
          let _1415_v42;
          _1415_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1414_v41,_this.f14);
          let _1416_v43;
          _1416_v43 = _dafny.Seq.of(_1410_v37, _1410_v37, _dafny.Seq.UnicodeFromString("qvqxuld"));
          let _1417_v44;
          let _nw229 = Array((new BigNumber(10)).toNumber());
          _nw229[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1410_v37, _1410_v37);
          _nw229[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("ebteocn");
          _nw229[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("lronodbd");
          _nw229[(new BigNumber(3)).toNumber()] = _1410_v37;
          _nw229[(new BigNumber(4)).toNumber()] = _1410_v37;
          _nw229[(new BigNumber(5)).toNumber()] = _module.__default.fm27(_1374_v14, true, globalState);
          _nw229[(new BigNumber(6)).toNumber()] = (((((_1415_v42).contains(_1414_v41)) ? ((_1415_v42).get(_1414_v41)) : ((_this).f19))) ? (_1410_v37) : (_1410_v37));
          _nw229[(new BigNumber(7)).toNumber()] = (_1416_v43)[_module.__default.safeIndex(_1374_v14, new BigNumber((_1416_v43).length))];
          _nw229[(new BigNumber(8)).toNumber()] = (((_this).f19) ? (_1410_v37) : (_dafny.Seq.UnicodeFromString("evmfvq")));
          _nw229[(new BigNumber(9)).toNumber()] = _1410_v37;
          _1417_v44 = _nw229;
          let _index254 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1417_v44).length));
          (_1417_v44)[_index254] = _1410_v37;
          let _index255 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1417_v44).length));
          (_1417_v44)[_index255] = _1410_v37;
          (globalState).f3 = _1374_v14;
          let _1418_v45;
          _1418_v45 = _dafny.Seq.of((_this).f15);
          (_this).m7(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(!(_this.f14))), _1418_v45), _dafny.areEqual(_1413_v40, _1413_v40), _1374_v14, _this.f14, globalState);
          let _1419_v46;
          let _nw230 = Array((new BigNumber(24)).toNumber()).fill(false);
          _1419_v46 = _nw230;
          let _1420_v47;
          _1420_v47 = _dafny.MultiSet.fromElements(_1412_v39, _1412_v39, _1412_v39, _1412_v39, _1412_v39);
          let _index256 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_1419_v46).length));
          (_1419_v46)[_index256] = (_1374_v14).isEqualTo(new BigNumber((_1420_v47).cardinality()));
          let _1421_v48;
          _1421_v48 = _dafny.MultiSet.fromElements(new BigNumber((_1413_v40).length), new BigNumber(180));
          let _1422_v49;
          let _init42 = ((_1423_v14) => function (_1424_i9) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1425_v14) => function (_1426_i10) {
              return _dafny.Map.Empty.slice().updateUnsafe(_1425_v14,false);
            })(_1423_v14));
          })(_1374_v14);
          let _nw231 = Array((new BigNumber(28)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw231.length); _i0_42++) {
            _nw231[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1422_v49 = _nw231;
          let _1427_v50;
          _1427_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v14,_this.f14);
          let _1428_v51;
          _1428_v51 = _dafny.Seq.of(_1427_v50);
          let _index257 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1422_v49).length));
          (_1422_v49)[_index257] = _1428_v51;
          let _1429_v52;
          _1429_v52 = _dafny.Set.fromElements(new BigNumber((_1421_v48).cardinality()), _1374_v14);
          let _index258 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_1419_v46).length));
          let _index259 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1422_v49).length));
          let _rhs231 = _this.f14;
          let _rhs232 = (new BigNumber((function () {
            let _coll53 = new _dafny.Set();
            for (const _compr_53 of (_1429_v52).Elements) {
              let _1430_v53 = _compr_53;
              if ((_1429_v52).contains(_1430_v53)) {
                _coll53.add((_1430_v53).minus(new BigNumber(216)));
              }
            }
            return _coll53;
          }()).length)).isLessThan((_1374_v14).minus(_1374_v14));
          let _rhs233 = _1421_v48;
          let _rhs234 = _dafny.Seq.of(_1427_v50, (_1427_v50).Merge(_1427_v50), _1427_v50, ((_1427_v50).update(new BigNumber(298), (_this).f19)).Merge((_1427_v50).update(_1374_v14, (_this).f19)));
          let _lhs197 = _1419_v46;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_1419_v46).length));
          let _lhs199 = globalState;
          let _lhs200 = _1422_v49;
          let _lhs201 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1422_v49).length));
          _lhs197[_lhs198] = _rhs231;
          _lhs199.f9 = _rhs232;
          _1421_v48 = _rhs233;
          _lhs200[_lhs201] = _rhs234;
        }
        let _1431_v54;
        _1431_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1374_v14);
        let _1432_v58;
        _1432_v58 = _dafny.Seq.of(_1431_v54);
        let _1433_v59;
        _1433_v59 = _dafny.Seq.of(_1374_v14, _1374_v14);
        let _1434_v60;
        _1434_v60 = _dafny.Set.fromElements(_1374_v14, (_1433_v59)[_module.__default.safeIndex(new BigNumber(((_this).f15).cardinality()), new BigNumber((_1433_v59).length))], _1374_v14);
        let _1435_v61;
        _1435_v61 = _dafny.Seq.UnicodeFromString("chvabiuu");
        let _1436_v62;
        let _nw232 = Array((new BigNumber(23)).toNumber());
        _nw232[(_dafny.ZERO).toNumber()] = (_1431_v54).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1374_v14));
        _nw232[(_dafny.ONE).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(2)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(3)).toNumber()] = (_1431_v54).Merge(_module.__default.fm1(false, _1374_v14, _this.f14, function () {
          let _coll54 = new _dafny.Set();
          for (const _compr_54 of _dafny.IntegerRange(new BigNumber(-621), new BigNumber(879))) {
            let _1437_v55 = _compr_54;
            if (((new BigNumber(-621)).isLessThanOrEqualTo(_1437_v55)) && ((_1437_v55).isLessThan(new BigNumber(879)))) {
              _coll54.add((_1437_v55).plus(new BigNumber((function () {
                let _coll55 = new _dafny.Map();
                for (const _compr_55 of (function () {
                  let _coll56 = new _dafny.Map();
                  for (const _compr_56 of (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1374_v14),_this.f14)).Keys.Elements) {
                    let _1438_v57 = _compr_56;
                    if ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1374_v14),_this.f14)).contains(_1438_v57)) {
                      _coll56.push([_module.__default.safeDivisionInt(_1438_v57, _1374_v14),(_this).f19]);
                    }
                  }
                  return _coll56;
                }()).Keys.Elements) {
                  let _1439_v56 = _compr_55;
                  if ((function () {
                    let _coll57 = new _dafny.Map();
                    for (const _compr_57 of (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1374_v14),_this.f14)).Keys.Elements) {
                      let _1438_v57 = _compr_57;
                      if ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1374_v14),_this.f14)).contains(_1438_v57)) {
                        _coll57.push([_module.__default.safeDivisionInt(_1438_v57, _1374_v14),(_this).f19]);
                      }
                    }
                    return _coll57;
                  }()).contains(_1439_v56)) {
                    _coll55.push([_module.__default.safeDivisionInt(_1439_v56, new BigNumber(403)),_1374_v14]);
                  }
                }
                return _coll55;
              }()).length)));
            }
          }
          return _coll54;
        }(), globalState));
        _nw232[(new BigNumber(4)).toNumber()] = ((_1432_v58)[_module.__default.safeIndex(_1374_v14, new BigNumber((_1432_v58).length))]).Merge(_1431_v54);
        _nw232[(new BigNumber(5)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(6)).toNumber()] = (((_this).f19) ? (_1431_v54) : (_1431_v54));
        _nw232[(new BigNumber(7)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1374_v14);
        _nw232[(new BigNumber(9)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(10)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(11)).toNumber()] = (_1431_v54).update(false, _1374_v14);
        _nw232[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1374_v14)).Merge(_module.__default.fm1((_this).f19, _1374_v14, _this.f14, _1434_v60, globalState));
        _nw232[(new BigNumber(13)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(14)).toNumber()] = (_1431_v54).Merge(_1431_v54);
        _nw232[(new BigNumber(15)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(16)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(17)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(18)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(19)).toNumber()] = _1431_v54;
        _nw232[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1374_v14);
        _nw232[(new BigNumber(21)).toNumber()] = (_1431_v54).Merge(_1431_v54);
        _nw232[(new BigNumber(22)).toNumber()] = (_1431_v54).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_this).f19),new BigNumber((_1435_v61).length)));
        _1436_v62 = _nw232;
        let _index260 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1436_v62).length));
        (_1436_v62)[_index260] = _1431_v54;
        let _index261 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1436_v62).length));
        (_1436_v62)[_index261] = _1431_v54;
        _1374_v14 = (_dafny.ZERO).minus((((_this).f19) ? (_1374_v14) : ((_dafny.ZERO).minus(_1374_v14))));
        let _1440_v63;
        let _nw233 = new _module.C1();
        _nw233.__ctor();
        _1440_v63 = _nw233;
      } else {
        (globalState).f9 = _this.f14;
        if ((_this).f19) {
          let _1441_v64;
          _1441_v64 = _dafny.Seq.UnicodeFromString("frcf");
          _1441_v64 = _dafny.Seq.Concat(_1441_v64, _1441_v64);
          let _1442_v65;
          _1442_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,((_this).f15).Intersect((_this).f15));
          _1442_v65 = (_1442_v65).update(!((_this).f19) || (!((_this).f19)), (_this).f15);
          let _1443_v66;
          let _nw234 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1443_v66 = _nw234;
          let _1444_v67;
          _1444_v67 = _module.D4.create_DC13(_1441_v64, _1374_v14, (_this).f19, _1443_v66);
          let _1445_v68;
          _1445_v68 = _dafny.MultiSet.fromElements(_1374_v14, _1374_v14, _1374_v14);
          _1441_v64 = (_this).fm3(_module.__default.fm0(_this.f14, new BigNumber(((_1444_v67).dtor_cf22).length), _1374_v14, _1445_v68, globalState), new BigNumber((_dafny.Set.fromElements(_1374_v14)).length), _this.f14, _1374_v14, globalState);
          let _1446_v69;
          _1446_v69 = new _dafny.CodePoint('m'.codePointAt(0));
          let _1447_v70;
          _1447_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1446_v69,_this.f14);
          let _index262 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1443_v66).length));
          (_1443_v66)[_index262] = false;
          let _1448_v71;
          _1448_v71 = _dafny.Seq.of((_this).f19);
          let _1449_v72;
          _1449_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1441_v64).length),true);
          let _1450_v73;
          _1450_v73 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22(_1374_v14, _dafny.Seq.of(_1449_v72), _this.f14, (_this).f19, globalState),_1443_v66);
          let _1451_v74;
          _1451_v74 = _dafny.Set.fromElements(_1445_v68);
          let _1452_v75;
          _1452_v75 = _dafny.MultiSet.fromElements(_1448_v71, _1448_v71, _1448_v71, _1448_v71);
          let _index263 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1443_v66).length));
          let _rhs235 = (_this).f19;
          let _rhs236 = (((_1448_v71)[_module.__default.safeIndex(new BigNumber((_1450_v73).length), new BigNumber((_1448_v71).length))]) ? ((_1447_v70).Merge(_1447_v70)) : ((_module.__default.fm46(_1374_v14, (_this).f19, globalState)).Merge(_1447_v70)));
          let _rhs237 = true;
          let _rhs238 = (_this).fm17(_1374_v14, (_1451_v74).IsDisjointFrom(_1451_v74), _1452_v75, _1374_v14, globalState);
          let _lhs202 = globalState;
          let _lhs203 = _1443_v66;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1443_v66).length));
          _lhs202.f9 = _rhs235;
          _1447_v70 = _rhs236;
          _lhs203[_lhs204] = _rhs237;
          _1374_v14 = _rhs238;
          (globalState).f3 = _module.__default.safeDivisionInt(_1374_v14, _1374_v14);
        } else {
          let _1453_v76;
          let _nw235 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1453_v76 = _nw235;
          let _1454_v77;
          _1454_v77 = _dafny.Set.fromElements(_1453_v76, _1453_v76);
          (globalState).f3 = (_1374_v14).multipliedBy(new BigNumber((_1454_v77).length));
          let _1455_v78;
          _1455_v78 = _dafny.Seq.UnicodeFromString("wkk");
          let _1456_v79;
          _1456_v79 = _dafny.Set.fromElements(_dafny.Seq.IsProperPrefixOf(_1455_v78, _1455_v78));
          (globalState).f6 = _1456_v79;
          (globalState).f3 = (new BigNumber((_1455_v78).length)).multipliedBy(_1374_v14);
          let _1457_v80;
          _1457_v80 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_1374_v14);
          let _1458_v81;
          _1458_v81 = _dafny.Seq.of(_1374_v14, new BigNumber((_1455_v78).length));
          let _1459_v82;
          _1459_v82 = _dafny.Set.fromElements(_1458_v81);
          let _1460_v83;
          _1460_v83 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm23(_1374_v14, _this.f14, _1457_v80, _1459_v82, globalState)) || (_this.f14),_1374_v14);
          _1460_v83 = (_1460_v83).update(_this.f14, _1374_v14);
          (globalState).f9 = (_this).f19;
        }
        let _1461_v84;
        let _init43 = function (_1462_i11) {
          return _dafny.Seq.Concat(_dafny.Seq.of((_this).f19), _dafny.Seq.of(!((_this).f19)));
        };
        let _nw236 = Array((new BigNumber(9)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw236.length); _i0_43++) {
          _nw236[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1461_v84 = _nw236;
        let _1463_v85;
        _1463_v85 = _dafny.Seq.of(_this.f14, _this.f14, (_this).f19);
        let _index264 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1461_v84).length));
        (_1461_v84)[_index264] = _1463_v85;
        let _1464_v86;
        _1464_v86 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_this.f14)),_1374_v14);
        let _1465_v87;
        _1465_v87 = _dafny.Seq.UnicodeFromString("pw");
        let _1466_v88;
        _1466_v88 = _dafny.Seq.of(_1463_v85);
        let _1467_v89;
        _1467_v89 = _dafny.Seq.of(_1374_v14, (_this).fm17((((_1464_v86).contains(_this.f14)) ? ((_1464_v86).get(_this.f14)) : (_1374_v14)), _this.f14, _dafny.MultiSet.fromElements(_1463_v85, _dafny.Seq.update(_1463_v85, _module.__default.safeIndex(new BigNumber((_1465_v87).length), new BigNumber((_1463_v85).length)), (_this).f19), _1463_v85, _1463_v85, _1463_v85), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_1466_v88, _module.__default.safeIndex((_dafny.ZERO).minus(_1374_v14), new BigNumber((_1466_v88).length)), _dafny.Seq.of((_this).f19))).length)), globalState));
        let _index265 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1461_v84).length));
        let _rhs239 = _dafny.Seq.update(_1463_v85, _module.__default.safeIndex(_1374_v14, new BigNumber((_1463_v85).length)), (_this).f19);
        let _rhs240 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), ((_1468_v14) => function (_1469_i12) {
          return _module.__default.safeModuloInt(_1468_v14, new BigNumber((_dafny.Seq.UnicodeFromString("ldnplt")).length));
        })(_1374_v14));
        let _lhs205 = _1461_v84;
        let _lhs206 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1461_v84).length));
        _lhs205[_lhs206] = _rhs239;
        _1467_v89 = _rhs240;
        (globalState).f9 = !(!((_this.f14) === (!(_this.f14))));
        (globalState).f9 = _this.f14;
      }
      let _1470_v90;
      _1470_v90 = new _dafny.CodePoint('o'.codePointAt(0));
      _1470_v90 = _1470_v90;
      let _1471_v91;
      let _nw237 = Array((new BigNumber(7)).toNumber()).fill(false);
      _1471_v91 = _nw237;
      let _1472_v92;
      _1472_v92 = _dafny.Seq.of((_this).f19, true, (_this).f19, _this.f14, (_this).f19);
      let _1473_v93;
      _1473_v93 = _module.D1.create_DC3(new BigNumber(338));
      let _1474_v94;
      _1474_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v14,_1473_v93);
      let _index266 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1471_v91).length));
      (_1471_v91)[_index266] = (_1472_v92)[_module.__default.safeIndex(new BigNumber((_1474_v94).length), new BigNumber((_1472_v92).length))];
      let _index267 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1471_v91).length));
      (_1471_v91)[_index267] = (_this).f19;
      let _1475_v95;
      _1475_v95 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber(101));
      let _1476_v96;
      _1476_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v14,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("voqhvj"), _module.__default.safeIndex(_1374_v14, new BigNumber((_dafny.Seq.UnicodeFromString("voqhvj")).length)), _1470_v90));
      let _1477_v97;
      _1477_v97 = _dafny.MultiSet.fromElements(new BigNumber((_1475_v95).length), _1374_v14, new BigNumber((_1476_v96).length));
      _1470_v90 = _module.__default.fm25(_1374_v14, ((_dafny.ZERO).minus(_1374_v14)).minus(new BigNumber((_1477_v97).cardinality())), !((_1471_v91)[_module.__default.safeIndex(new BigNumber(821), new BigNumber((_1471_v91).length))]) || ((_this).f19), _1475_v95, globalState);
      return;
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1478_v0;
      let _nw238 = new _module.C1();
      _nw238.__ctor();
      _1478_v0 = _nw238;
      let _1479_v1;
      let _nw239 = Array((new BigNumber(20)).toNumber()).fill(false);
      _1479_v1 = _nw239;
      let _index268 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length));
      (_1479_v1)[_index268] = p1;
      let _1480_v2;
      _1480_v2 = _dafny.Set.fromElements((_this).f15);
      let _1481_v3;
      let _nw240 = Array((new BigNumber(7)).toNumber());
      _nw240[(_dafny.ZERO).toNumber()] = _1480_v2;
      _nw240[(_dafny.ONE).toNumber()] = (_module.D11.create_DC33(_1480_v2)).dtor_cf57;
      _nw240[(new BigNumber(2)).toNumber()] = _1480_v2;
      _nw240[(new BigNumber(3)).toNumber()] = _1480_v2;
      _nw240[(new BigNumber(4)).toNumber()] = (_1480_v2).Difference(_1480_v2);
      _nw240[(new BigNumber(5)).toNumber()] = _1480_v2;
      _nw240[(new BigNumber(6)).toNumber()] = (_1480_v2).Difference(_dafny.Set.fromElements((_this).f15, (_this).f15));
      _1481_v3 = _nw240;
      let _index269 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1481_v3).length));
      (_1481_v3)[_index269] = (_module.__default.fm47(globalState)).Difference(_1480_v2);
      let _1482_v4;
      _1482_v4 = _dafny.Set.fromElements(p2);
      let _1483_v5;
      _1483_v5 = _dafny.Seq.UnicodeFromString("ovck");
      let _1484_v6;
      _1484_v6 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1482_v4).length)), (_dafny.ZERO).minus(new BigNumber(-295)), p2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1483_v5).length),_1478_v0)).length));
      let _index270 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length));
      let _index271 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1481_v3).length));
      let _rhs241 = (!(p3) || (p3)) && (_dafny.Seq.IsPrefixOf(_1484_v6, _1484_v6));
      let _rhs242 = _1480_v2;
      let _rhs243 = true;
      let _lhs207 = _1479_v1;
      let _lhs208 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length));
      let _lhs209 = _1481_v3;
      let _lhs210 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1481_v3).length));
      let _lhs211 = globalState;
      _lhs207[_lhs208] = _rhs241;
      _lhs209[_lhs210] = _rhs242;
      _lhs211.f9 = _rhs243;
      let _rhs244 = p2;
      let _rhs245 = new BigNumber(538);
      let _lhs212 = globalState;
      let _lhs213 = globalState;
      _lhs212.f3 = _rhs244;
      _lhs213.f3 = _rhs245;
      let _1485_v7;
      _1485_v7 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1486_v8;
      _1486_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(new BigNumber((_1483_v5).length)));
      let _1487_v9;
      _1487_v9 = _module.D4.create_DC13(_dafny.Seq.UnicodeFromString("sikdur"), p2, (_this).f19, _1479_v1);
      let _1488_v11;
      _1488_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
      let _1489_v12;
      _1489_v12 = _dafny.Seq.of(function () {
        let _coll58 = new _dafny.Map();
        for (const _compr_58 of _dafny.IntegerRange(new BigNumber(-314), new BigNumber(-878))) {
          let _1490_v10 = _compr_58;
          if (((new BigNumber(-314)).isLessThanOrEqualTo(_1490_v10)) && ((_1490_v10).isLessThan(new BigNumber(-878)))) {
            _coll58.push([(_1490_v10).plus(p2),_this.f14]);
          }
        }
        return _coll58;
      }(), _1488_v11, _1488_v11, _1488_v11);
      let _1491_v13;
      _1491_v13 = _module.D9.create_DC30(_1485_v7, p1, new BigNumber((_1486_v8).length), _module.__default.fm22((_1487_v9).dtor_cf23, _1489_v12, (_this).f19, (_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))], globalState), p2);
      let _1492_v14;
      _1492_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1491_v13,p2);
      let _hi4 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), ((_1493_v7) => function (_1494_i1) {
        return _1493_v7;
      })(_1485_v7))).length)).plus(p2);
      for (let _1495_i0 = new BigNumber(((_1492_v14).update(_1491_v13, p2)).length); _1495_i0.isLessThan(_hi4); _1495_i0 = _1495_i0.plus(_dafny.ONE)) {
        let _1496_v15;
        _1496_v15 = _dafny.MultiSet.fromElements(p2);
        let _1497_v16;
        _1497_v16 = _module.D2.create_DC9((_this).f19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), function (_1498_i2) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length), _1485_v7);
        let _1499_v17;
        let _nw241 = Array((new BigNumber(6)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = _1497_v16;
        _nw241[(_dafny.ONE).toNumber()] = _1497_v16;
        _nw241[(new BigNumber(2)).toNumber()] = _1497_v16;
        _nw241[(new BigNumber(3)).toNumber()] = _1497_v16;
        _nw241[(new BigNumber(4)).toNumber()] = _module.D2.create_DC9(p3, p2, _1485_v7);
        _nw241[(new BigNumber(5)).toNumber()] = _module.D2.create_DC9(p1, _1495_i0, _1485_v7);
        _1499_v17 = _nw241;
        let _1500_v18;
        _1500_v18 = _dafny.Seq.of(_1499_v17);
        let _1501_v19;
        let _nw242 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1501_v19 = _nw242;
        let _1502_v20;
        _1502_v20 = _module.D5.create_DC16(_1501_v19);
        let _1503_v21;
        _1503_v21 = _module.D5.create_DC19(_1502_v20);
        let _1504_v22;
        _1504_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(_this.f14, p3, globalState),_1502_v20);
        let _1505_v23;
        _1505_v23 = _module.D5.create_DC19((((_1504_v22).contains((_this).f19)) ? ((_1504_v22).get((_this).f19)) : (_1503_v21)));
        let _1506_v24;
        _1506_v24 = _module.D7.create_DC22(_1479_v1, _1500_v18, _dafny.Seq.of(_1479_v1, _1479_v1), _1479_v1, _1505_v23);
        let _1507_v25;
        _1507_v25 = _dafny.Seq.of(_1506_v24);
        let _1508_v26;
        let _nw243 = Array((new BigNumber(11)).toNumber());
        _nw243[(_dafny.ZERO).toNumber()] = p2;
        _nw243[(_dafny.ONE).toNumber()] = (p2).plus(p2);
        _nw243[(new BigNumber(2)).toNumber()] = p2;
        _nw243[(new BigNumber(3)).toNumber()] = _1495_i0;
        _nw243[(new BigNumber(4)).toNumber()] = (((((_this).f15).contains((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))])) ? (((_this).f15).get((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))])) : (_module.__default.fm0((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))], new BigNumber(24), _1495_i0, _1496_v15, globalState)))).multipliedBy(p2);
        _nw243[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1495_i0,p2)).length));
        _nw243[(new BigNumber(6)).toNumber()] = new BigNumber((_1483_v5).length);
        _nw243[(new BigNumber(7)).toNumber()] = _module.__default.fm0(_this.f14, p2, _1495_i0, _1496_v15, globalState);
        _nw243[(new BigNumber(8)).toNumber()] = p2;
        _nw243[(new BigNumber(9)).toNumber()] = (_1495_i0).minus(new BigNumber((_1483_v5).length));
        _nw243[(new BigNumber(10)).toNumber()] = new BigNumber((_1507_v25).length);
        _1508_v26 = _nw243;
        let _1509_v27;
        _1509_v27 = _dafny.Seq.of(_this.f14, _this.f14, _this.f14, p3, _this.f14);
        let _1510_v28;
        _1510_v28 = _dafny.MultiSet.fromElements(_1509_v27);
        let _index272 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1501_v19).length));
        (_1501_v19)[_index272] = (_this).fm17(p2, !((_this).f19), _1510_v28, (_dafny.ZERO).minus(_module.__default.fm0((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))], _1495_i0, _1495_i0, _1496_v15, globalState)), globalState);
        let _1511_v29;
        _1511_v29 = _module.D12.create_DC37(_1510_v28);
        let _index273 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_1501_v19).length));
        (_1501_v19)[_index273] = _module.__default.safeModuloInt((p2).plus(new BigNumber(528)), (_this).fm17(_1495_i0, (_1509_v27)[_module.__default.safeIndex(p2, new BigNumber((_1509_v27).length))], (_1511_v29).dtor_cf63, _1495_i0, globalState));
        let _1512_v30;
        _1512_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_1501_v19)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_1501_v19).length))]);
        (globalState).f9 = ((_1495_i0).plus((((_1512_v30).contains((_this).f19)) ? ((_1512_v30).get((_this).f19)) : (p2)))).isLessThan(new BigNumber(128));
        let _rhs246 = _1485_v7;
        let _rhs247 = _1485_v7;
        let _rhs248 = _1501_v19;
        let _rhs249 = (((_1512_v30).contains(p1)) ? ((_1512_v30).get(p1)) : ((p2).minus(p2)));
        let _lhs214 = globalState;
        _1485_v7 = _rhs246;
        _1485_v7 = _rhs247;
        _1501_v19 = _rhs248;
        _lhs214.f3 = _rhs249;
        (globalState).f9 = _dafny.areEqual(_1509_v27, _1509_v27);
      }
      let _1513_v31;
      let _nw244 = new _module.C2();
      _nw244.__ctor(_1488_v11);
      _1513_v31 = _nw244;
      _1513_v31 = _1513_v31;
      let _1514_v32;
      _1514_v32 = _dafny.Seq.of(p1);
      let _1515_v33;
      _1515_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))],_dafny.Seq.Concat(_1514_v32, _1514_v32));
      let _rhs250 = new BigNumber(((((_1515_v33).contains((_this.f14) && (p1))) ? ((_1515_v33).get((_this.f14) && (p1))) : (_1514_v32))).length);
      let _rhs251 = !((_1479_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_1479_v1).length))]) || ((_this).f19);
      let _rhs252 = (_this).f19;
      let _lhs215 = globalState;
      let _lhs216 = globalState;
      let _lhs217 = globalState;
      _lhs215.f3 = _rhs250;
      _lhs216.f9 = _rhs251;
      _lhs217.f9 = _rhs252;
      return;
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f14 = false;
      this._f15 = _dafny.MultiSet.Empty;
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f18, f14, f15) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return _this.f14;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("e"));
    };
    fm15(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("y"), _dafny.Seq.UnicodeFromString("kfx")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), function (_1516_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_1517_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })));
    };
    fm16(p0, p1, p2, globalState) {
      let _this = this;
      return !(_this.f14);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _1518_v0;
      _1518_v0 = _dafny.Seq.of(_this.f14);
      let _1519_v1;
      _1519_v1 = _module.D2.create_DC7(_1518_v0);
      let _1520_v2;
      _1520_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1519_v1).dtor_cf9);
      let _1521_v3;
      _1521_v3 = _dafny.Seq.of(_1518_v0, _1518_v0, _1518_v0, _1518_v0, _1518_v0);
      _1520_v2 = (_1520_v2).update((_this).f18, (_1521_v3)[_module.__default.safeIndex(new BigNumber(840), new BigNumber((_1521_v3).length))]);
      let _1522_v4;
      _1522_v4 = _dafny.Seq.of(_dafny.Set.fromElements(true, _this.f14, _this.f14, true, _this.f14));
      let _1523_v5;
      let _nw245 = new _module.C4();
      _nw245.__ctor(_this.f14, (new BigNumber((_1522_v4).length)).isLessThan(new BigNumber((_1521_v3).length)), (_this).f15);
      _1523_v5 = _nw245;
      r0 = new BigNumber(((_1521_v3)[_module.__default.safeIndex((_this).f18, new BigNumber((_1521_v3).length))]).length);
      r0 = (_this).f18;
      let _1524_v6;
      let _nw246 = Array((_dafny.ONE).toNumber()).fill(false);
      _1524_v6 = _nw246;
      let _index274 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length));
      (_1524_v6)[_index274] = _this.f14;
      let _index275 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length));
      (_1524_v6)[_index275] = ((_this.f14) ? (_this.f14) : ((_1523_v5).f19));
      let _1525_v7;
      _1525_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
      let _1526_v8;
      _1526_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1523_v5).f19,new BigNumber((_1525_v7).length));
      let _1527_v9;
      _1527_v9 = _dafny.Seq.UnicodeFromString("lm");
      let _1528_v10;
      _1528_v10 = _dafny.Seq.of((_this).f18);
      let _1529_v11;
      let _nw247 = Array((new BigNumber(26)).toNumber());
      _nw247[(_dafny.ZERO).toNumber()] = (_this).f18;
      _nw247[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(2)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(3)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(4)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(5)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(6)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(7)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(8)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(9)).toNumber()] = new BigNumber((_1526_v8).length);
      _nw247[(new BigNumber(10)).toNumber()] = new BigNumber((_1527_v9).length);
      _nw247[(new BigNumber(11)).toNumber()] = new BigNumber(-604);
      _nw247[(new BigNumber(12)).toNumber()] = new BigNumber((_1528_v10).length);
      _nw247[(new BigNumber(13)).toNumber()] = new BigNumber(817);
      _nw247[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.of((_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))], (_1523_v5).f19)).length);
      _nw247[(new BigNumber(15)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(16)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(17)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(18)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(19)).toNumber()] = new BigNumber(975);
      _nw247[(new BigNumber(20)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(21)).toNumber()] = new BigNumber(-187);
      _nw247[(new BigNumber(22)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(23)).toNumber()] = (_this).f18;
      _nw247[(new BigNumber(24)).toNumber()] = (_1528_v10)[_module.__default.safeIndex((_this).f18, new BigNumber((_1528_v10).length))];
      _nw247[(new BigNumber(25)).toNumber()] = (_this).f18;
      _1529_v11 = _nw247;
      let _1530_v12;
      _1530_v12 = _module.D5.create_DC16(_1529_v11);
      let _1531_v13;
      _1531_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1527_v9,_1529_v11);
      let _1532_v14;
      _1532_v14 = _dafny.Map.Empty.slice().updateUnsafe((_1523_v5).f19,_1529_v11);
      let _pat_let_tv40 = _1529_v11;
      let _1533_v15;
      let _nw248 = Array((new BigNumber(26)).toNumber());
      _nw248[(_dafny.ZERO).toNumber()] = (function (_pat_let56_0) {
        return function (_1534_dt__update__tmp_h0) {
          return function (_pat_let57_0) {
            return function (_1535_dt__update_hcf29_h0) {
              return _module.D5.create_DC16(_1535_dt__update_hcf29_h0);
            }(_pat_let57_0);
          }(_pat_let_tv40);
        }(_pat_let56_0);
      }(_1530_v12)).dtor_cf29;
      _nw248[(_dafny.ONE).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(2)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(3)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(4)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(5)).toNumber()] = (((_1523_v5).f19) ? (_1529_v11) : (_1529_v11));
      _nw248[(new BigNumber(6)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(7)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(8)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(9)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(10)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(11)).toNumber()] = (((_1531_v13).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(979)), function (_1537_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }))) ? ((_1531_v13).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(979)), function (_1536_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }))) : (_1529_v11));
      _nw248[(new BigNumber(12)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(13)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(14)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(15)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(16)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(17)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(18)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(19)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(20)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(21)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(22)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(23)).toNumber()] = _1529_v11;
      _nw248[(new BigNumber(24)).toNumber()] = (((_1532_v14).contains((_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))])) ? ((_1532_v14).get((_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))])) : (_1529_v11));
      _nw248[(new BigNumber(25)).toNumber()] = _1529_v11;
      _1533_v15 = _nw248;
      let _nw249 = Array((new BigNumber(9)).toNumber()).fill([]);
      _1533_v15 = _nw249;
      r0 = (_this).f18;
      r1 = (_this.f14) && (_this.f14);
      let _1538_v16;
      _1538_v16 = _dafny.Seq.of(_1524_v6);
      let _1539_v17;
      _1539_v17 = _module.D4.create_DC13(_1527_v9, (_this).f18, (_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))], (_1538_v16)[_module.__default.safeIndex((_this).f18, new BigNumber((_1538_v16).length))]);
      let _1540_v18;
      _1540_v18 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))])).update((_this).f18, (_1523_v5).f19)).update((_1539_v17).dtor_cf23, _this.f14),(_this).f18);
      r2 = _1540_v18;
      r3 = (_1524_v6)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_1524_v6).length))];
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1541_v0;
      let _init44 = function (_1542_i1) {
        return (_1542_i1).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)))).cardinality()));
      };
      let _nw250 = Array((new BigNumber(27)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw250.length); _i0_44++) {
        _nw250[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1541_v0 = _nw250;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1541_v0).length))) {
        let _1543_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1543_i0)) && ((_1543_i0).isLessThan(new BigNumber((_1541_v0).length))))) {
          (_1541_v0)[(_1543_i0)] = _module.__default.safeModuloInt(_1543_i0, p1);
        }
      }
      let _1544_v1;
      _1544_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      let _1545_v2;
      _1545_v2 = _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14));
      _1544_v1 = function (_source15) {
        if (_source15.is_DC11) {
          let _1546___mcc_h0 = (_source15).cf17;
          let _1547___mcc_h1 = (_source15).cf18;
          let _1548___mcc_h2 = (_source15).cf19;
          let _1549___mcc_h3 = (_source15).cf20;
          let _1550_cf20 = _1549___mcc_h3;
          let _1551_cf19 = _1548___mcc_h2;
          let _1552_cf18 = _1547___mcc_h1;
          let _1553_cf17 = _1546___mcc_h0;
          return _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_1551_cf19);
        } else {
          let _1554___mcc_h4 = (_source15).cf16;
          let _1555_cf16 = _1554___mcc_h4;
          return (_1555_cf16).Merge(_1555_cf16);
        }
      }(_1545_v2);
      let _1556_v3;
      _1556_v3 = _dafny.Seq.of((_this).f18);
      let _1557_v4;
      _1557_v4 = _dafny.Seq.of(_this.f14);
      let _1558_v5;
      let _nw251 = new _module.C4();
      _nw251.__ctor(_dafny.Seq.IsPrefixOf(_1556_v3, _1556_v3), _this.f14, _dafny.MultiSet.fromElements((_1557_v4)[_module.__default.safeIndex((_this).f18, new BigNumber((_1557_v4).length))], _this.f14));
      _1558_v5 = _nw251;
      let _1559_v6;
      _1559_v6 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1560_v7;
      _1560_v7 = _dafny.Seq.UnicodeFromString("h");
      let _1561_v8;
      _1561_v8 = _dafny.Seq.of(_1560_v7, _1560_v7);
      let _1562_v9;
      let _nw252 = Array((new BigNumber(21)).toNumber());
      _nw252[(_dafny.ZERO).toNumber()] = (_1558_v5).f19;
      _nw252[(_dafny.ONE).toNumber()] = (_1558_v5).f19;
      _nw252[(new BigNumber(2)).toNumber()] = ((_this).f18).isLessThanOrEqualTo(new BigNumber(720));
      _nw252[(new BigNumber(3)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(4)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(5)).toNumber()] = ((_this).f15).IsSubsetOf((_this).f15);
      _nw252[(new BigNumber(6)).toNumber()] = (_1558_v5).f19;
      _nw252[(new BigNumber(7)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(8)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(9)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(10)).toNumber()] = true;
      _nw252[(new BigNumber(11)).toNumber()] = _this.f14;
      _nw252[(new BigNumber(12)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(815)), ((_1563_v6) => function (_1564_i2) {
        return _1563_v6;
      })(_1559_v6)), _1559_v6);
      _nw252[(new BigNumber(13)).toNumber()] = (p1).isLessThanOrEqualTo(new BigNumber((_1560_v7).length));
      _nw252[(new BigNumber(14)).toNumber()] = ((_this).f18).isLessThan(new BigNumber((_1561_v8).length));
      _nw252[(new BigNumber(15)).toNumber()] = !(p0).equals(p0);
      _nw252[(new BigNumber(16)).toNumber()] = !((_1558_v5).f19);
      _nw252[(new BigNumber(17)).toNumber()] = (p0).IsSubsetOf(p0);
      _nw252[(new BigNumber(18)).toNumber()] = (_1558_v5).f19;
      _nw252[(new BigNumber(19)).toNumber()] = (_this).fm2((_1558_v5).f19, (_1558_v5).f19, globalState);
      _nw252[(new BigNumber(20)).toNumber()] = _this.f14;
      _1562_v9 = _nw252;
      let _index276 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
      (_1562_v9)[_index276] = _this.f14;
      let _index277 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
      (_1541_v0)[_index277] = new BigNumber(-934);
      let _1565_v10;
      let _nw253 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
      _1565_v10 = _nw253;
      let _1566_v11;
      _1566_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(p1, p1, _this.f14, p1, globalState),_this.f14);
      let _index278 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1565_v10).length));
      (_1565_v10)[_index278] = _1566_v11;
      let _1567_v12;
      _1567_v12 = _dafny.MultiSet.fromElements(new BigNumber(-837));
      let _index279 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
      let _index280 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
      let _index281 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1565_v10).length));
      let _rhs253 = (_1567_v12).IsDisjointFrom((_dafny.MultiSet.fromElements(p1)).update((_this).f18, _module.__default.abs((_this).f18)));
      let _rhs254 = (_dafny.ZERO).minus((_this).f18);
      let _rhs255 = (_1558_v5).f19;
      let _rhs256 = _1566_v11;
      let _lhs218 = _1562_v9;
      let _lhs219 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
      let _lhs220 = _1541_v0;
      let _lhs221 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
      let _lhs222 = _this;
      let _lhs223 = _1565_v10;
      let _lhs224 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1565_v10).length));
      _lhs218[_lhs219] = _rhs253;
      _lhs220[_lhs221] = _rhs254;
      _lhs222.f14 = _rhs255;
      _lhs223[_lhs224] = _rhs256;
      let _1568_v13;
      _1568_v13 = _dafny.Seq.of(_1567_v12, _dafny.MultiSet.fromElements((_this).f18));
      let _1569_v14;
      _1569_v14 = _dafny.Set.fromElements((_1568_v13)[_module.__default.safeIndex(p1, new BigNumber((_1568_v13).length))]);
      let _hi5 = _module.__default.safeDivisionInt((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))], new BigNumber((_1569_v14).length));
      for (let _1570_i3 = (_dafny.ZERO).minus((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))]); _1570_i3.isLessThan(_hi5); _1570_i3 = _1570_i3.plus(_dafny.ONE)) {
        let _1571_v15;
        _1571_v15 = _dafny.Set.fromElements(_1557_v4);
        let _1572_v17;
        _1572_v17 = _dafny.Set.fromElements((_1558_v5).fm2((_1558_v5).f19, _this.f14, globalState));
        let _1573_v18;
        _1573_v18 = _dafny.Map.Empty.slice().updateUnsafe((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))],_1572_v17);
        let _1574_v19;
        _1574_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1556_v3,new BigNumber(((((_1573_v18).contains((_1558_v5).f19)) ? ((_1573_v18).get((_1558_v5).f19)) : (_1572_v17))).length));
        let _index282 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
        let _index283 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
        let _rhs257 = _1560_v7;
        let _rhs258 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(((_dafny.ZERO).minus(p1)).multipliedBy(_1570_i3), new BigNumber(((function () {
          let _coll59 = new _dafny.Set();
          for (const _compr_59 of (_1571_v15).Elements) {
            let _1575_v16 = _compr_59;
            if ((_1571_v15).contains(_1575_v16)) {
              _coll59.add(_1575_v16);
            }
          }
          return _coll59;
        }()).Intersect(_1571_v15)).length)));
        let _rhs259 = (_this).f18;
        let _rhs260 = (_this).fm15((_1558_v5).f19, ((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))]).isEqualTo(p1), (((_1574_v19).contains(_1556_v3)) ? ((_1574_v19).get(_1556_v3)) : (p1)), globalState);
        let _rhs261 = (_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))];
        let _lhs225 = _1541_v0;
        let _lhs226 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
        let _lhs227 = _1541_v0;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
        let _lhs229 = globalState;
        _1560_v7 = _rhs257;
        _lhs225[_lhs226] = _rhs258;
        _lhs227[_lhs228] = _rhs259;
        _1560_v7 = _rhs260;
        _lhs229.f9 = _rhs261;
        let _1576_v20;
        let _init45 = function (_1577_i4) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        };
        let _nw254 = Array((new BigNumber(7)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw254.length); _i0_45++) {
          _nw254[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1576_v20 = _nw254;
        let _1578_v21;
        _1578_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))],new BigNumber(95));
        let _index284 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1576_v20).length));
        (_1576_v20)[_index284] = (((_1558_v5).f19) ? (_module.__default.fm25(p1, p1, false, _1578_v21, globalState)) : (_1559_v6));
        let _index285 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1576_v20).length));
        (_1576_v20)[_index285] = new _dafny.CodePoint('w'.codePointAt(0));
        let _1579_v22;
        _1579_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1578_v21,(_this).fm2((_1558_v5).f19, (_1558_v5).f19, globalState));
        let _1580_v24;
        _1580_v24 = _dafny.Set.fromElements(_1578_v21);
        if ((function () {
          let _coll60 = new _dafny.Set();
          for (const _compr_60 of (_1579_v22).Keys.Elements) {
            let _1581_v23 = _compr_60;
            if ((_1579_v22).contains(_1581_v23)) {
              _coll60.add(_1581_v23);
            }
          }
          return _coll60;
        }()).IsSubsetOf(_1580_v24)) {
          r1 = _1570_i3;
          let _index286 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1576_v20).length));
          (_1576_v20)[_index286] = new _dafny.CodePoint('h'.codePointAt(0));
          let _index287 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
          let _rhs262 = (_1558_v5).f19;
          let _rhs263 = !(((_1558_v5).f19) === (_this.f14));
          let _rhs264 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_1582_p0) => function (_1583_i5) {
            return new BigNumber((_1582_p0).length);
          })(p0))).length);
          let _lhs230 = globalState;
          let _lhs231 = globalState;
          let _lhs232 = _1541_v0;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length));
          _lhs230.f9 = _rhs262;
          _lhs231.f9 = _rhs263;
          _lhs232[_lhs233] = _rhs264;
          let _nw255 = Array((new BigNumber(29)).toNumber()).fill(false);
          (globalState).f2 = _nw255;
          let _1584_v25;
          let _nw256 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1584_v25 = _nw256;
          let _1585_v26;
          _1585_v26 = _module.D0.create_DC1((_1558_v5).f19);
          let _rhs265 = (_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))];
          let _rhs266 = (((_1558_v5).f19) ? (!(_1578_v21).contains((_1585_v26).dtor_cf1)) : (((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))]) || ((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))])));
          let _rhs267 = ((new BigNumber(23)).multipliedBy(_1570_i3)).minus((new BigNumber(63)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)));
          let _rhs268 = _1584_v25;
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          _lhs234.f3 = _rhs265;
          r2 = _rhs266;
          _lhs235.f3 = _rhs267;
          _1584_v25 = _rhs268;
        } else {
          let _1586_v27;
          let _init46 = ((_1587_v3) => function (_1588_i6) {
            return _dafny.Seq.Concat(_1587_v3, _1587_v3);
          })(_1556_v3);
          let _nw257 = Array((new BigNumber(11)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw257.length); _i0_46++) {
            _nw257[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1586_v27 = _nw257;
          _1586_v27 = _1586_v27;
          let _index288 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
          (_1562_v9)[_index288] = !(!(!(true)));
          (_this).f14 = !((p1).isLessThanOrEqualTo((_this).f18));
          (globalState).f3 = p1;
          _1541_v0 = _1541_v0;
        }
        let _1589_v28;
        _1589_v28 = _dafny.Seq.of(_1562_v9);
        let _1590_v29;
        _1590_v29 = _module.D1.create_DC5((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))], (_this).fm2(false, _this.f14, globalState));
        let _1591_v30;
        _1591_v30 = _module.D11.create_DC34((_this).f15, _1567_v12, _1562_v9, _1590_v29);
        let _1592_v31;
        let _nw258 = Array((new BigNumber(18)).toNumber());
        _nw258[(_dafny.ZERO).toNumber()] = _1562_v9;
        _nw258[(_dafny.ONE).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(2)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(3)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(4)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(5)).toNumber()] = (_1589_v28)[_module.__default.safeIndex(p1, new BigNumber((_1589_v28).length))];
        _nw258[(new BigNumber(6)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(7)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(8)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(9)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(10)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(11)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(12)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(13)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(14)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(15)).toNumber()] = (_1591_v30).dtor_cf60;
        _nw258[(new BigNumber(16)).toNumber()] = _1562_v9;
        _nw258[(new BigNumber(17)).toNumber()] = _1562_v9;
        _1592_v31 = _nw258;
        let _index289 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1592_v31).length));
        (_1592_v31)[_index289] = _1562_v9;
        let _1593_v32;
        _1593_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1576_v20)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_1576_v20).length))],_1557_v4);
        let _index290 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1592_v31).length));
        let _index291 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
        let _rhs269 = (_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))];
        let _rhs270 = _1562_v9;
        let _rhs271 = (((_this).f15).Difference((_this).f15)).IsDisjointFrom((_this).f15);
        let _rhs272 = _1593_v32;
        let _lhs236 = _1592_v31;
        let _lhs237 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1592_v31).length));
        let _lhs238 = _1562_v9;
        let _lhs239 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length));
        r2 = _rhs269;
        _lhs236[_lhs237] = _rhs270;
        _lhs238[_lhs239] = _rhs271;
        _1593_v32 = _rhs272;
      }
      let _1594_v33;
      _1594_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))],_1562_v9);
      let _1595_v34;
      _1595_v34 = _dafny.MultiSet.fromElements(_1557_v4, _1557_v4);
      let _1596_v35;
      _1596_v35 = _dafny.Seq.of(_1562_v9);
      let _1597_v36;
      _1597_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1559_v6,_1562_v9);
      let _1598_v37;
      _1598_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1562_v9)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_1562_v9).length))],(_this).f18);
      let _1599_v38;
      let _nw259 = Array((new BigNumber(24)).toNumber());
      _nw259[(_dafny.ZERO).toNumber()] = _1562_v9;
      _nw259[(_dafny.ONE).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(2)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(3)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(4)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(5)).toNumber()] = (((_1594_v33).contains((_1558_v5).fm17((_this).f18, (_1558_v5).f19, _1595_v34, (_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))], globalState))) ? ((_1594_v33).get((_1558_v5).fm17((_this).f18, (_1558_v5).f19, _1595_v34, (_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))], globalState))) : (_1562_v9));
      _nw259[(new BigNumber(6)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(7)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(8)).toNumber()] = (_1596_v35)[_module.__default.safeIndex((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))], new BigNumber((_1596_v35).length))];
      _nw259[(new BigNumber(9)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(10)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(11)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(12)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(13)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(14)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(15)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(16)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(17)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(18)).toNumber()] = (((_1597_v36).contains(_module.__default.fm25(p1, p1, (_1558_v5).f19, _1598_v37, globalState))) ? ((_1597_v36).get(_module.__default.fm25(p1, p1, (_1558_v5).f19, _1598_v37, globalState))) : (_1562_v9));
      _nw259[(new BigNumber(19)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(20)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(21)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(22)).toNumber()] = _1562_v9;
      _nw259[(new BigNumber(23)).toNumber()] = _1562_v9;
      _1599_v38 = _nw259;
      let _1600_v39;
      _1600_v39 = _module.D4.create_DC12(_1562_v9);
      let _1601_v40;
      _1601_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1600_v39,_1599_v38);
      let _1602_v41;
      _1602_v41 = _dafny.Seq.of(_1599_v38);
      _1599_v38 = (((_1601_v40).contains(_1600_v39)) ? ((_1601_v40).get(_1600_v39)) : ((_1602_v41)[_module.__default.safeIndex((_1556_v3)[_module.__default.safeIndex(p1, new BigNumber((_1556_v3).length))], new BigNumber((_1602_v41).length))]));
      r0 = new BigNumber(561);
      r1 = p1;
      r2 = ((_1541_v0)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1541_v0).length))]).isLessThan((_dafny.ZERO).minus(p1));
      return [r0, r1, r2];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f14 = false;
      this._f15 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f14, f15) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return !(function () {
        let _coll61 = new _dafny.Set();
        for (const _compr_61 of _dafny.IntegerRange(new BigNumber(490), new BigNumber(573))) {
          let _1603_v0 = _compr_61;
          if (((new BigNumber(490)).isLessThanOrEqualTo(_1603_v0)) && ((_1603_v0).isLessThan(new BigNumber(573)))) {
            _coll61.add(_module.__default.safeDivisionInt(_1603_v0, new BigNumber(621)));
          }
        }
        return _coll61;
      }()).equals((function () {
        let _coll62 = new _dafny.Set();
        for (const _compr_62 of (_dafny.Set.fromElements(new BigNumber(-224))).Elements) {
          let _1604_v1 = _compr_62;
          if ((_dafny.Set.fromElements(new BigNumber(-224))).contains(_1604_v1)) {
            _coll62.add((_1604_v1).multipliedBy(new BigNumber(544)));
          }
        }
        return _coll62;
      }()).Difference(function () {
        let _coll63 = new _dafny.Set();
        for (const _compr_63 of _dafny.IntegerRange(new BigNumber(613), new BigNumber(145))) {
          let _1605_v2 = _compr_63;
          if (((new BigNumber(613)).isLessThanOrEqualTo(_1605_v2)) && ((_1605_v2).isLessThan(new BigNumber(145)))) {
            _coll63.add((_1605_v2).minus(new BigNumber(943)));
          }
        }
        return _coll63;
      }()));
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("evj");
    };
    fm11(p0, p1, globalState) {
      let _this = this;
      return (_this.f14) && ((function () {
        let _coll64 = new _dafny.Set();
        for (const _compr_64 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(857),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-494)), function (_1606_i0) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }))).cardinality()))).length))).cardinality()), new BigNumber(339))).Elements) {
          let _1607_v0 = _compr_64;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(857),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-494)), function (_1606_i0) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          }))).cardinality()))).length))).cardinality()), new BigNumber(339))).contains(_1607_v0)) {
            _coll64.add((_1607_v0).minus(new BigNumber(177)));
          }
        }
        return _coll64;
      }()).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(864), new BigNumber(476), new BigNumber(180), (_module.D1.create_DC3(new BigNumber(646))).dtor_cf3, new BigNumber(305))).length), new BigNumber((_dafny.Seq.UnicodeFromString("njb")).length))));
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(new BigNumber((_dafny.Seq.of(_this.f14)).length)),(_this).f15)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("aksalss")).length)),(_this).f15))).equals(function () {
        let _coll65 = new _dafny.Map();
        for (const _compr_65 of (function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_dafny.Seq.of(_module.D1.create_DC3(new BigNumber(-509)))).Elements) {
            let _1608_v1 = _compr_66;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC3(new BigNumber(-509))), _1608_v1)) {
              _coll66.push([_1608_v1,new BigNumber(121)]);
            }
          }
          return _coll66;
        }()).Keys.Elements) {
          let _1609_v0 = _compr_65;
          if ((function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of (_dafny.Seq.of(_module.D1.create_DC3(new BigNumber(-509)))).Elements) {
              let _1608_v1 = _compr_67;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC3(new BigNumber(-509))), _1608_v1)) {
                _coll67.push([_1608_v1,new BigNumber(121)]);
              }
            }
            return _coll67;
          }()).contains(_1609_v0)) {
            _coll65.push([_1609_v0,(_this).f15]);
          }
        }
        return _coll65;
      }());
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _1610_v0;
      _1610_v0 = new BigNumber(-804);
      let _hi6 = _1610_v0;
      for (let _1611_i0 = _1610_v0; _1611_i0.isLessThan(_hi6); _1611_i0 = _1611_i0.plus(_dafny.ONE)) {
        let _1612_v1;
        _1612_v1 = _dafny.Seq.UnicodeFromString("li");
        let _1613_v2;
        _1613_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), function (_1614_i1) {
          return _dafny.Seq.UnicodeFromString("pf");
        }), _module.__default.safeIndex(_1610_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), function (_1615_i1) {
          return _dafny.Seq.UnicodeFromString("pf");
        })).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(343)), function (_1616_i2) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }))).length),_1612_v1);
        _1613_v2 = (_1613_v2).update(_1610_v0, ((_this.f14) ? (_1612_v1) : (_1612_v1)));
        let _1617_v3;
        _1617_v3 = _module.D0.create_DC0(_this.f14);
        if (((((_1617_v3).dtor_cf0) ? (_1617_v3) : (_1617_v3))).dtor_cf0) {
          let _1618_v4;
          _1618_v4 = _dafny.Seq.of(new BigNumber((_1612_v1).length));
          let _1619_v5;
          _1619_v5 = _dafny.MultiSet.fromElements(_1611_i0);
          let _1620_v6;
          _1620_v6 = _dafny.MultiSet.fromElements(((_this.f14) ? (_dafny.Seq.of(_module.__default.fm0(_this.f14, _1611_i0, new BigNumber((_1618_v4).length), _1619_v5, globalState))) : (_dafny.Seq.of(_1611_i0))), _1618_v4, _dafny.Seq.update(_1618_v4, _module.__default.safeIndex(_1610_v0, new BigNumber((_1618_v4).length)), _1610_v0));
          _1620_v6 = _module.__default.fm13((_this).f15, (new BigNumber(488)).multipliedBy(_1611_i0), globalState);
          let _1621_v7;
          let _nw260 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _1621_v7 = _nw260;
          let _1622_v8;
          let _nw261 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1622_v8 = _nw261;
          let _1623_v9;
          _1623_v9 = _dafny.Seq.of(_1622_v8, _1622_v8, _1622_v8);
          let _index292 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1621_v7).length));
          (_1621_v7)[_index292] = _1623_v9;
          let _index293 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1621_v7).length));
          (_1621_v7)[_index293] = _dafny.Seq.of(_1622_v8);
          let _1624_v10;
          _1624_v10 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(false)).equals((_this).f15),_1618_v4);
          _1624_v10 = (_1624_v10).update(_this.f14, _1618_v4);
          let _1625_v11;
          let _1626_v12;
          let _1627_v13;
          let _1628_v14;
          let _out47;
          let _out48;
          let _out49;
          let _out50;
          let _outcollector16 = (_this).m5(globalState);
          _out47 = _outcollector16[0];
          _out48 = _outcollector16[1];
          _out49 = _outcollector16[2];
          _out50 = _outcollector16[3];
          _1625_v11 = _out47;
          _1626_v12 = _out48;
          _1627_v13 = _out49;
          _1628_v14 = _out50;
          let _1629_v15;
          _1629_v15 = _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(512)), function (_1630_i3) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }), _1612_v1), true, _1628_v14, _1628_v14, _1628_v14);
          let _rhs273 = _dafny.Seq.Concat(_1629_v15, _1629_v15);
          let _rhs274 = _1622_v8;
          let _rhs275 = _this.f14;
          let _lhs240 = globalState;
          _1629_v15 = _rhs273;
          _1622_v8 = _rhs274;
          _lhs240.f9 = _rhs275;
        } else {
          let _1631_v16;
          _1631_v16 = _dafny.Set.fromElements(_this.f14);
          let _1632_v17;
          _1632_v17 = _dafny.Seq.of(_1631_v16);
          let _1633_v18;
          _1633_v18 = _dafny.MultiSet.fromElements(_1611_i0);
          let _1634_v19;
          _1634_v19 = _dafny.MultiSet.fromElements(_module.__default.fm0(_this.f14, (_dafny.ZERO).minus(new BigNumber((_1633_v18).cardinality())), _1611_i0, _1633_v18, globalState), new BigNumber(-334), _1611_i0, (_dafny.ZERO).minus(_1610_v0), _1611_i0);
          let _1635_v20;
          _1635_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_1632_v17, _dafny.Seq.of(_1631_v16, _1631_v16)),(_1634_v19).IsProperSubsetOf(_1633_v18));
          let _1636_v21;
          _1636_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14(globalState),_1610_v0);
          let _1637_v22;
          _1637_v22 = _dafny.Seq.of(_1636_v21);
          _1635_v20 = (_1635_v20).update(!((_this).fm12(_this.f14, (_1637_v22)[_module.__default.safeIndex(_1610_v0, new BigNumber((_1637_v22).length))], globalState)), _this.f14);
          let _1638_v23;
          _1638_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1612_v1,_module.D0.create_DC0(_this.f14));
          let _1639_v24;
          _1639_v24 = _dafny.Seq.of(_this.f14, (_this).fm12(_this.f14, _module.__default.fm48(_1611_i0, _1610_v0, _this.f14, globalState), globalState), _this.f14);
          let _1640_v25;
          let _nw262 = new _module.C5();
          _nw262.__ctor(_1610_v0, true, _dafny.MultiSet.FromArray(_1639_v24));
          _1640_v25 = _nw262;
          let _1641_v26;
          _1641_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1640_v25,_1612_v1);
          let _1642_v27;
          let _nw263 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1642_v27 = _nw263;
          _1638_v23 = (_1638_v23).update(_dafny.Seq.Concat((((_1641_v26).contains(_1640_v25)) ? ((_1641_v26).get(_1640_v25)) : ((_module.D4.create_DC13(_1612_v1, _1610_v0, _this.f14, _1642_v27)).dtor_cf22)), _1612_v1), _1617_v3);
          let _1643_v28;
          _1643_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jvdqndys"),_1611_i0);
          r3 = !(_1643_v28).contains(_dafny.Seq.Concat(_1612_v1, _1612_v1));
          r0 = _1610_v0;
          let _1644_v29;
          _1644_v29 = _dafny.Seq.of(_1635_v20);
          let _1645_v30;
          _1645_v30 = _dafny.MultiSet.fromElements(_1635_v20);
          let _1646_v31;
          _1646_v31 = _dafny.Seq.of(_1645_v30);
          let _1647_v32;
          _1647_v32 = new _dafny.CodePoint('b'.codePointAt(0));
          let _1648_v33;
          _1648_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1611_i0,_1647_v32);
          let _1649_v34;
          let _nw264 = Array((new BigNumber(19)).toNumber());
          _nw264[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.FromArray(_1644_v29)).Union(_1645_v30);
          _nw264[(_dafny.ONE).toNumber()] = (_1645_v30).Intersect(_1645_v30);
          _nw264[(new BigNumber(2)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(3)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(4)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(5)).toNumber()] = (_1646_v31)[_module.__default.safeIndex(new BigNumber((_1648_v33).length), new BigNumber((_1646_v31).length))];
          _nw264[(new BigNumber(6)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.update(_1644_v29, _module.__default.safeIndex(new BigNumber((_1631_v16).length), new BigNumber((_1644_v29).length)), _1635_v20));
          _nw264[(new BigNumber(8)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(9)).toNumber()] = (_1645_v30).Difference(_1645_v30);
          _nw264[(new BigNumber(10)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(11)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(12)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(13)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(14)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(15)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(16)).toNumber()] = _1645_v30;
          _nw264[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements(_1635_v20, _1635_v20);
          _nw264[(new BigNumber(18)).toNumber()] = (_1645_v30).update(_dafny.Map.Empty.slice().updateUnsafe(false,_1640_v25.f14), _module.__default.abs(new BigNumber((_1639_v24).length)));
          _1649_v34 = _nw264;
          let _index294 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1649_v34).length));
          (_1649_v34)[_index294] = (_1645_v30).Union(_1645_v30);
          let _index295 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_1649_v34).length));
          (_1649_v34)[_index295] = _dafny.MultiSet.FromArray(_1644_v29);
        }
        if (_this.f14) {
          let _1650_v35;
          let _nw265 = Array((new BigNumber(27)).toNumber());
          _1650_v35 = _nw265;
          _1650_v35 = _1650_v35;
          (globalState).f9 = _this.f14;
          let _1651_v36;
          _1651_v36 = _dafny.Seq.of(_1610_v0);
          (globalState).f3 = _module.__default.safeModuloInt(_1611_i0, new BigNumber((_1651_v36).length));
          let _1652_v37;
          _1652_v37 = _dafny.Set.fromElements((_this).f15);
          let _1653_v38;
          _1653_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.D11.create_DC33(_1652_v37),_1610_v0);
          let _1654_v39;
          _1654_v39 = _module.D11.create_DC33(_1652_v37);
          _1653_v38 = (_1653_v38).update(_1654_v39, ((_dafny.ZERO).minus(_1611_i0)).multipliedBy(_1610_v0));
          let _1655_v40;
          let _nw266 = new _module.C1();
          _nw266.__ctor();
          _1655_v40 = _nw266;
        } else {
          let _1656_v41;
          let _nw267 = Array((new BigNumber(2)).toNumber());
          _nw267[(_dafny.ZERO).toNumber()] = _1611_i0;
          _nw267[(_dafny.ONE).toNumber()] = _1610_v0;
          _1656_v41 = _nw267;
          let _index296 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1656_v41).length));
          (_1656_v41)[_index296] = _1610_v0;
          let _1657_v42;
          let _nw268 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1657_v42 = _nw268;
          let _1658_v43;
          _1658_v43 = new _dafny.CodePoint('t'.codePointAt(0));
          let _index297 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1657_v42).length));
          (_1657_v42)[_index297] = _1658_v43;
          let _index298 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1656_v41).length));
          let _index299 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1657_v42).length));
          let _rhs276 = (_1610_v0).plus(_1611_i0);
          let _rhs277 = _1611_i0;
          let _rhs278 = _1658_v43;
          let _rhs279 = (_dafny.ZERO).minus(_1611_i0);
          let _lhs241 = globalState;
          let _lhs242 = _1656_v41;
          let _lhs243 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_1656_v41).length));
          let _lhs244 = _1657_v42;
          let _lhs245 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1657_v42).length));
          _lhs241.f3 = _rhs276;
          _lhs242[_lhs243] = _rhs277;
          _lhs244[_lhs245] = _rhs278;
          _1610_v0 = _rhs279;
          let _1659_v45;
          let _init47 = ((_1660_v0, _1661_i0) => function (_1662_i4) {
            return (function () {
              let _coll68 = new _dafny.Set();
              for (const _compr_68 of (_dafny.Set.fromElements(_1661_i0)).Elements) {
                let _1663_v44 = _compr_68;
                if ((_dafny.Set.fromElements(_1661_i0)).contains(_1663_v44)) {
                  _coll68.add((_1663_v44).minus(new BigNumber(274)));
                }
              }
              return _coll68;
            }()).IsDisjointFrom(_dafny.Set.fromElements(_1660_v0));
          })(_1610_v0, _1611_i0);
          let _nw269 = Array((new BigNumber(3)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw269.length); _i0_47++) {
            _nw269[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1659_v45 = _nw269;
          let _index300 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_1659_v45).length));
          (_1659_v45)[_index300] = true;
          let _1664_v46;
          _1664_v46 = _dafny.Seq.of(_this.f14, (_this.f14) || (true), _this.f14);
          let _1665_v47;
          _1665_v47 = _dafny.Set.fromElements(_this.f14);
          let _index301 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_1659_v45).length));
          let _rhs280 = _this.f14;
          let _rhs281 = (_1664_v46)[_module.__default.safeIndex(new BigNumber((((_this.f14) ? (_dafny.Set.fromElements(_this.f14, true)) : (_1665_v47))).length), new BigNumber((_1664_v46).length))];
          let _rhs282 = _this.f14;
          let _lhs246 = _1659_v45;
          let _lhs247 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_1659_v45).length));
          r3 = _rhs280;
          r1 = _rhs281;
          _lhs246[_lhs247] = _rhs282;
          let _1666_v48;
          _1666_v48 = _dafny.Map.Empty.slice().updateUnsafe((_1657_v42)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_1657_v42).length))],new BigNumber(891));
          let _1667_v49;
          _1667_v49 = _dafny.Seq.of(_1610_v0, _1611_i0);
          let _1668_v50;
          _1668_v50 = _dafny.Set.fromElements(_1667_v49, _1667_v49, _dafny.Seq.update(_1667_v49, _module.__default.safeIndex((_1656_v41)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_1656_v41).length))], new BigNumber((_1667_v49).length)), (_1656_v41)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_1656_v41).length))]), _1667_v49);
          (globalState).f9 = !(((true) ? (_module.__default.fm23((_dafny.ZERO).minus(_1610_v0), _this.f14, _1666_v48, _1668_v50, globalState)) : (false)));
          let _1669_v51;
          let _1670_v52;
          let _1671_v53;
          let _1672_v54;
          let _out51;
          let _out52;
          let _out53;
          let _out54;
          let _outcollector17 = (_this).m5(globalState);
          _out51 = _outcollector17[0];
          _out52 = _outcollector17[1];
          _out53 = _outcollector17[2];
          _out54 = _outcollector17[3];
          _1669_v51 = _out51;
          _1670_v52 = _out52;
          _1671_v53 = _out53;
          _1672_v54 = _out54;
          _1612_v1 = _dafny.Seq.Concat(_dafny.Seq.of(_1658_v43), _1612_v1);
        }
        let _1673_v55;
        let _nw270 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1673_v55 = _nw270;
        let _1674_v56;
        _1674_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1610_v0,_1610_v0);
        let _index302 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1673_v55).length));
        (_1673_v55)[_index302] = new BigNumber(((_1674_v56).Merge((function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of _dafny.IntegerRange(new BigNumber(964), new BigNumber(-867))) {
            let _1675_v57 = _compr_69;
            if (((new BigNumber(964)).isLessThanOrEqualTo(_1675_v57)) && ((_1675_v57).isLessThan(new BigNumber(-867)))) {
              _coll69.push([(_1675_v57).minus(new BigNumber((_1612_v1).length)),_1611_i0]);
            }
          }
          return _coll69;
        }()).update((_dafny.ZERO).minus(_1610_v0), _1611_i0))).length);
        let _1676_v58;
        _1676_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1612_v1,(_dafny.ZERO).minus(_1610_v0));
        let _1677_v59;
        _1677_v59 = _dafny.MultiSet.fromElements(new BigNumber((_1612_v1).length));
        let _1678_v60;
        _1678_v60 = _dafny.Seq.of(_1610_v0);
        let _1679_v61;
        _1679_v61 = _dafny.Set.fromElements(_dafny.Seq.of(_module.__default.fm0(_this.f14, _1611_i0, _1611_i0, _1677_v59, globalState)), _1678_v60, _1678_v60);
        let _index303 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1673_v55).length));
        (_1673_v55)[_index303] = _module.__default.safeDivisionInt((((_1676_v58).contains(_1612_v1)) ? ((_1676_v58).get(_1612_v1)) : (new BigNumber(((_module.D13.create_DC41(_1679_v61)).dtor_cf67).length))), _1611_i0);
      }
      let _1680_v62;
      let _nw271 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
      _1680_v62 = _nw271;
      let _1681_v63;
      _1681_v63 = _dafny.Set.fromElements(_1610_v0);
      let _index304 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1680_v62).length));
      (_1680_v62)[_index304] = _1681_v63;
      let _index305 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1680_v62).length));
      (_1680_v62)[_index305] = _1681_v63;
      let _1682_v64;
      _1682_v64 = new _dafny.CodePoint('h'.codePointAt(0));
      let _1683_v65;
      _1683_v65 = _dafny.MultiSet.fromElements(_1682_v64);
      _1683_v65 = (_1683_v65).Intersect((_1683_v65).Difference(_1683_v65));
      let _1684_v66;
      let _nw272 = Array((new BigNumber(10)).toNumber()).fill(false);
      _1684_v66 = _nw272;
      let _index306 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length));
      (_1684_v66)[_index306] = !(_this.f14);
      let _index307 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length));
      (_1684_v66)[_index307] = _this.f14;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1684_v66).length))) {
        let _1685_i5 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1685_i5)) && ((_1685_i5).isLessThan(new BigNumber((_1684_v66).length))))) {
          (_1684_v66)[(_1685_i5)] = !((_1610_v0).plus(_1610_v0)).isEqualTo(_1610_v0);
        }
      }
      let _1686_v67;
      let _1687_v68;
      let _1688_v69;
      let _1689_v70;
      let _out55;
      let _out56;
      let _out57;
      let _out58;
      let _outcollector18 = (_this).m5(globalState);
      _out55 = _outcollector18[0];
      _out56 = _outcollector18[1];
      _out57 = _outcollector18[2];
      _out58 = _outcollector18[3];
      _1686_v67 = _out55;
      _1687_v68 = _out56;
      _1688_v69 = _out57;
      _1689_v70 = _out58;
      r0 = ((_1610_v0).multipliedBy(_1687_v68)).plus(new BigNumber((function () {
        let _coll70 = new _dafny.Map();
        for (const _compr_70 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), ((_1690_v0) => function (_1691_i6) {
          return _1690_v0;
        })(_1610_v0))).Elements) {
          let _1692_v71 = _compr_70;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), ((_1693_v0) => function (_1691_i6) {
            return _1693_v0;
          })(_1610_v0)), _1692_v71)) {
            _coll70.push([_module.__default.safeDivisionInt(_1692_v71, _1687_v68),_1610_v0]);
          }
        }
        return _coll70;
      }()).length));
      let _1694_v72;
      _1694_v72 = _module.D1.create_DC3((_dafny.ZERO).minus(_1610_v0));
      let _1695_v73;
      _1695_v73 = _module.D1.create_DC6(_1694_v72);
      let _pat_let_tv41 = _1682_v64;
      let _pat_let_tv42 = _1682_v64;
      r1 = function (_source16) {
        if (_source16.is_DC4) {
          let _1696___mcc_h0 = (_source16).cf4;
          let _1697___mcc_h1 = (_source16).cf5;
          let _1698_cf5 = _1697___mcc_h1;
          let _1699_cf4 = _1696___mcc_h0;
          return (_module.D13.create_DC42(new _dafny.CodePoint('r'.codePointAt(0)), _this.f14)).dtor_cf69;
        } else if (_source16.is_DC5) {
          let _1700___mcc_h2 = (_source16).cf6;
          let _1701___mcc_h3 = (_source16).cf7;
          let _1702_cf7 = _1701___mcc_h3;
          let _1703_cf6 = _1700___mcc_h2;
          return _this.f14;
        } else if (_source16.is_DC3) {
          let _1704___mcc_h4 = (_source16).cf3;
          let _1705_cf3 = _1704___mcc_h4;
          return _this.f14;
        } else {
          let _1706___mcc_h5 = (_source16).cf8;
          let _1707_cf8 = _1706___mcc_h5;
          return (_dafny.Set.fromElements(_pat_let_tv41)).IsSubsetOf(_dafny.Set.fromElements(_pat_let_tv42));
        }
      }(((true) ? (_1695_v73) : (_1695_v73)));
      let _1708_v74;
      _1708_v74 = _dafny.Map.Empty.slice().updateUnsafe((_1684_v66)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length))],_1686_v67);
      let _1709_v75;
      _1709_v75 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((((_1708_v74).contains((_1684_v66)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length))])) ? ((_1708_v74).get((_1684_v66)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length))])) : ((_dafny.ZERO).minus(_1687_v68))),(_1684_v66)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_1684_v66).length))]),_1610_v0);
      r2 = _1709_v75;
      r3 = _1689_v70;
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1710_v0;
      let _init48 = function (_1711_i0) {
        return _this.f14;
      };
      let _nw273 = Array((new BigNumber(5)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw273.length); _i0_48++) {
        _nw273[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1710_v0 = _nw273;
      let _1712_v1;
      _1712_v1 = _dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)));
      let _index308 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
      (_1710_v0)[_index308] = (p1).isLessThan(new BigNumber((_1712_v1).length));
      let _index309 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
      (_1710_v0)[_index309] = (_this).fm2(_this.f14, _this.f14, globalState);
      (globalState).f9 = (_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))];
      let _1713_v2;
      let _nw274 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1713_v2 = _nw274;
      let _1714_v3;
      let _nw275 = Array((new BigNumber(17)).toNumber());
      _nw275[(_dafny.ZERO).toNumber()] = _1713_v2;
      _nw275[(_dafny.ONE).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(2)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(3)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(4)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(5)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(6)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(7)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(8)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(9)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(10)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(11)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(12)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(13)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(14)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(15)).toNumber()] = _1713_v2;
      _nw275[(new BigNumber(16)).toNumber()] = (_module.D14.create_DC44(_1713_v2)).dtor_cf71;
      _1714_v3 = _nw275;
      let _index310 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1714_v3).length));
      (_1714_v3)[_index310] = _1713_v2;
      let _index311 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1714_v3).length));
      (_1714_v3)[_index311] = (((_module.D1.create_DC5(_this.f14, true)).dtor_cf7) ? (_1713_v2) : ((((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))]) ? (_1713_v2) : (_1713_v2))));
      let _1715_v4;
      _1715_v4 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1716_v5;
      _1716_v5 = _dafny.Set.fromElements(_1715_v4);
      let _1717_v6;
      _1717_v6 = _dafny.MultiSet.fromElements(p1);
      _1716_v5 = _module.__default.fm32(_1712_v1, _1717_v6, globalState);
      let _1718_v7;
      _1718_v7 = _module.D4.create_DC12(_1710_v0);
      let _1719_v8;
      _1719_v8 = _dafny.Seq.of(_1718_v7);
      let _1720_v9;
      let _nw276 = new _module.C4();
      _nw276.__ctor(!_dafny.areEqual(_1719_v8, _1719_v8), (_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))], (_this).f15);
      _1720_v9 = _nw276;
      if ((_1720_v9).f19) {
        if ((_1720_v9).f19) {
          _1710_v0 = _1710_v0;
          let _1721_v10;
          _1721_v10 = _dafny.Seq.of(_1712_v1, _1712_v1, _dafny.Seq.Concat(_1712_v1, _1712_v1), _1712_v1);
          _1721_v10 = _1721_v10;
          (globalState).f9 = (_1720_v9).f19;
          let _1722_v11;
          _1722_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((_this).f15).cardinality()));
          let _1723_v12;
          _1723_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1712_v1).length),_1722_v11);
          let _1724_v13;
          _1724_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p1,!(!((_1720_v9).f19))),p1);
          let _1725_v14;
          _1725_v14 = _module.D7.create_DC21(_1724_v13);
          let _1726_v15;
          _1726_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1725_v14,_1716_v5);
          let _1727_v16;
          _1727_v16 = _dafny.Seq.of(_1723_v12, _1723_v12);
          let _rhs283 = ((((_1726_v15).contains(_1725_v14)) ? ((_1726_v15).get(_1725_v14)) : (_1716_v5))).Intersect(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0))));
          let _rhs284 = p1;
          let _rhs285 = !((p1).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_module.__default.fm0((_1720_v9).f19, (_dafny.ZERO).minus(p1), p1, _1717_v6, globalState)), p1)).length)));
          let _rhs286 = (_1727_v16)[_module.__default.safeIndex(_module.__default.safeModuloInt(p1, p1), new BigNumber((_1727_v16).length))];
          let _lhs248 = globalState;
          _1716_v5 = _rhs283;
          _lhs248.f3 = _rhs284;
          r2 = _rhs285;
          _1723_v12 = _rhs286;
          let _1728_v17;
          let _nw277 = new _module.C1();
          _nw277.__ctor();
          _1728_v17 = _nw277;
        } else {
          let _1729_v18;
          _1729_v18 = _module.D4.create_DC13(_1712_v1, p1, (_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))], _1710_v0);
          let _rhs287 = _dafny.Seq.Concat((_1729_v18).dtor_cf22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), ((_1730_v4) => function (_1731_i1) {
            return _1730_v4;
          })(_1715_v4)));
          let _rhs288 = _1712_v1;
          _1712_v1 = _rhs287;
          _1712_v1 = _rhs288;
          let _1732_v19;
          _1732_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1720_v9).f19,p1);
          let _1733_v20;
          let _nw278 = Array((new BigNumber(19)).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = _1715_v4;
          _nw278[(_dafny.ONE).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(2)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(3)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(4)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(5)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(6)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(7)).toNumber()] = (_1712_v1)[_module.__default.safeIndex(new BigNumber((_1717_v6).cardinality()), new BigNumber((_1712_v1).length))];
          _nw278[(new BigNumber(8)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(9)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(10)).toNumber()] = _module.__default.fm25(p1, new BigNumber(129), _this.f14, _1732_v19, globalState);
          _nw278[(new BigNumber(11)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(12)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw278[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw278[(new BigNumber(15)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(16)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(17)).toNumber()] = _1715_v4;
          _nw278[(new BigNumber(18)).toNumber()] = _1715_v4;
          _1733_v20 = _nw278;
          let _index312 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1733_v20).length));
          (_1733_v20)[_index312] = _1715_v4;
          let _index313 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1733_v20).length));
          (_1733_v20)[_index313] = _1715_v4;
          let _1734_v21;
          let _nw279 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _1734_v21 = _nw279;
          let _1735_v22;
          _1735_v22 = _dafny.Set.fromElements(_module.D8.create_DC25());
          let _index314 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1734_v21).length));
          (_1734_v21)[_index314] = _1735_v22;
          let _index315 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1734_v21).length));
          (_1734_v21)[_index315] = _1735_v22;
          let _1736_v24;
          _1736_v24 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_1712_v1, (_1733_v20)[_module.__default.safeIndex(new BigNumber(774), new BigNumber((_1733_v20).length))]),(p0).Intersect(function () {
            let _coll71 = new _dafny.Set();
            for (const _compr_71 of _dafny.IntegerRange(new BigNumber(636), new BigNumber(181))) {
              let _1737_v23 = _compr_71;
              if (((new BigNumber(636)).isLessThanOrEqualTo(_1737_v23)) && ((_1737_v23).isLessThan(new BigNumber(181)))) {
                _coll71.add(_module.__default.safeDivisionInt(_1737_v23, p1));
              }
            }
            return _coll71;
          }()));
          _1736_v24 = _1736_v24;
          let _1738_v25;
          _1738_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1712_v1,new BigNumber(((_this).f15).cardinality()));
          _1738_v25 = (_1738_v25).update(_1712_v1, (new BigNumber((_1712_v1).length)).minus((_dafny.ZERO).minus(p1)));
        }
        let _1739_v26;
        _1739_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1720_v9).f19,!((p1).isLessThanOrEqualTo(p1)));
        _1739_v26 = (_1739_v26).update(_this.f14, (p1).isLessThan(p1));
        _1712_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-802)), ((_1740_v4) => function (_1741_i2) {
          return _1740_v4;
        })(_1715_v4));
        let _1742_v27;
        let _nw280 = Array((new BigNumber(5)).toNumber());
        _nw280[(_dafny.ZERO).toNumber()] = p0;
        _nw280[(_dafny.ONE).toNumber()] = p0;
        _nw280[(new BigNumber(2)).toNumber()] = p0;
        _nw280[(new BigNumber(3)).toNumber()] = p0;
        _nw280[(new BigNumber(4)).toNumber()] = (p0).Union(p0);
        _1742_v27 = _nw280;
        let _1743_v28;
        _1743_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1720_v9).f19);
        let _1744_v29;
        _1744_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1743_v28).length),p1);
        let _1745_v30;
        _1745_v30 = _dafny.Seq.of(p1);
        let _1746_v31;
        _1746_v31 = _dafny.Seq.of(new BigNumber((_1744_v29).length), p1, _module.__default.fm0(_this.f14, p1, p1, _1717_v6, globalState), new BigNumber((_1745_v30).length), p1);
        let _index316 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1742_v27).length));
        (_1742_v27)[_index316] = (_dafny.Set.fromElements(p1, p1)).Union(function () {
          let _coll72 = new _dafny.Set();
          for (const _compr_72 of (_1745_v30).Elements) {
            let _1747_v32 = _compr_72;
            if (_dafny.Seq.contains(_1745_v30, _1747_v32)) {
              _coll72.add((_1747_v32).plus(new BigNumber(876)));
            }
          }
          return _coll72;
        }());
        let _1748_v33;
        let _nw281 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1748_v33 = _nw281;
        let _index317 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length));
        (_1748_v33)[_index317] = (((_1744_v29).contains(p1)) ? ((_1744_v29).get(p1)) : (p1));
        let _1749_v34;
        _1749_v34 = _dafny.Set.fromElements(_1712_v1);
        let _index318 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1742_v27).length));
        let _index319 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length));
        let _rhs289 = _dafny.Set.fromElements(p1);
        let _rhs290 = p1;
        let _rhs291 = p1;
        let _rhs292 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-551)), ((_1750_v4) => function (_1751_i3) {
          return _1750_v4;
        })(_1715_v4));
        let _rhs293 = _1749_v34;
        let _lhs249 = _1742_v27;
        let _lhs250 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1742_v27).length));
        let _lhs251 = _1748_v33;
        let _lhs252 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length));
        _lhs249[_lhs250] = _rhs289;
        r1 = _rhs290;
        _lhs251[_lhs252] = _rhs291;
        _1712_v1 = _rhs292;
        _1749_v34 = _rhs293;
        let _1752_v35;
        _1752_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))],(_1748_v33)[_module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length))]),_1748_v33);
        let _1753_v36;
        _1753_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_1748_v33)[_module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length))]);
        let _1754_v37;
        _1754_v37 = _dafny.Seq.of((((_1752_v35).contains(_1753_v36)) ? ((_1752_v35).get(_1753_v36)) : (_1748_v33)), _1748_v33);
        let _1755_v38;
        _1755_v38 = _module.D5.create_DC16(_1748_v33);
        let _index320 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
        let _rhs294 = (_module.D0.create_DC0((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))])).dtor_cf0;
        let _rhs295 = _dafny.Seq.of(_1748_v33, _1748_v33, _1748_v33, (_1755_v38).dtor_cf29, _1748_v33);
        let _rhs296 = (_1748_v33)[_module.__default.safeIndex(new BigNumber(453), new BigNumber((_1748_v33).length))];
        let _rhs297 = !(true);
        let _lhs253 = _1710_v0;
        let _lhs254 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
        let _lhs255 = globalState;
        _lhs253[_lhs254] = _rhs294;
        _1754_v37 = _rhs295;
        r1 = _rhs296;
        _lhs255.f9 = _rhs297;
      } else {
        _1715_v4 = _1715_v4;
        let _1756_v39;
        let _nw282 = Array((new BigNumber(26)).toNumber()).fill([]);
        _1756_v39 = _nw282;
        let _nw283 = Array((new BigNumber(4)).toNumber()).fill([]);
        _1756_v39 = _nw283;
        r1 = (p1).minus(p1);
        if ((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))]) {
          let _1757_v40;
          let _init49 = ((_1758_p1) => function (_1759_i4) {
            return (_1759_i4).minus(_1758_p1);
          })(p1);
          let _nw284 = Array((new BigNumber(3)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw284.length); _i0_49++) {
            _nw284[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1757_v40 = _nw284;
          let _index321 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length));
          (_1757_v40)[_index321] = p1;
          let _index322 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length));
          (_1757_v40)[_index322] = p1;
          let _index323 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
          (_1710_v0)[_index323] = (_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))];
          let _1760_v41;
          _1760_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), ((_1761_v4) => function (_1762_i5) {
            return _1761_v4;
          })(_1715_v4))).length),(((_1720_v9).f19) ? (_1712_v1) : (_1712_v1)));
          _1760_v41 = (_1760_v41).update((new BigNumber((_1712_v1).length)).minus(new BigNumber((_1712_v1).length)), (((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))]) ? (_dafny.Seq.UnicodeFromString("aupfiojbx")) : (_1712_v1)));
          let _1763_v42;
          _1763_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))],(_1757_v40)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length))]);
          let _1764_v43;
          _1764_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1715_v4,new BigNumber(720));
          let _1765_v44;
          _1765_v44 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1720_v9).f19,_1715_v4)).length), (_dafny.ZERO).minus(p1));
          let _1766_v45;
          _1766_v45 = _dafny.Set.fromElements(_1765_v44, _1765_v44);
          let _1767_v46;
          _1767_v46 = _module.D13.create_DC41(_1766_v45);
          let _1768_v47;
          _1768_v47 = _dafny.Map.Empty.slice().updateUnsafe((((_1763_v42).contains((_1720_v9).f19)) ? ((_1763_v42).get((_1720_v9).f19)) : ((_1757_v40)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length))])),_module.__default.fm23(p1, (_1720_v9).f19, _1764_v43, (_1767_v46).dtor_cf67, globalState));
          let _1769_v48;
          let _nw285 = new _module.C2();
          _nw285.__ctor(_1768_v47);
          _1769_v48 = _nw285;
          let _1770_v49;
          let _nw286 = new _module.C5();
          _nw286.__ctor(((_1757_v40)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length))]).multipliedBy(p1), (p1).isLessThan((_1757_v40)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_1757_v40).length))]), ((_this).f15).Union((_this).f15));
          _1770_v49 = _nw286;
        } else {
          (_1720_v9).m6(globalState);
          let _1771_v50;
          _1771_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1717_v6).cardinality()),_this.f14);
          let _1772_v51;
          _1772_v51 = _dafny.Seq.of((_1720_v9).f19);
          let _1773_v52;
          let _nw287 = new _module.C2();
          _nw287.__ctor((_1771_v50).Merge((_1771_v50).update(new BigNumber((_dafny.Seq.of(_1772_v51)).length), (_1720_v9).f19)));
          _1773_v52 = _nw287;
          _1773_v52 = _1773_v52;
          let _1774_v53;
          _1774_v53 = _dafny.Seq.of(p1, p1);
          let _1775_v54;
          let _nw288 = new _module.C0();
          _nw288.__ctor(_this.f14, (_1720_v9).f19);
          _1775_v54 = _nw288;
          let _1776_v55;
          _1776_v55 = _dafny.Set.fromElements(_1775_v54);
          let _1777_v56;
          _1777_v56 = _dafny.Seq.of(_1776_v55, _1776_v55);
          let _index324 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length));
          (_1710_v0)[_index324] = !(p1).isEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_1774_v53)[_module.__default.safeIndex(new BigNumber((_1712_v1).length), new BigNumber((_1774_v53).length))])), new BigNumber((_1777_v56).length)));
          let _1778_v57;
          let _nw289 = new _module.C3();
          _nw289.__ctor(new BigNumber(360), _1712_v1, (_1775_v54).f23, (_this).f15);
          _1778_v57 = _nw289;
          let _1779_v58;
          _1779_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1778_v57,new BigNumber((_dafny.Seq.UnicodeFromString("tfija")).length));
          let _1780_v59;
          _1780_v59 = _module.D15.create_DC47(_1779_v58);
          r1 = (_module.D3.create_DC11((_dafny.ZERO).minus(p1), (_1720_v9).fm17(new BigNumber(((_1780_v59).dtor_cf76).length), _this.f14, _dafny.MultiSet.fromElements(_1772_v51), (_1778_v57).f20, globalState), (_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))], (_1778_v57).f20)).dtor_cf17;
          _1774_v53 = _1774_v53;
        }
        _1715_v4 = _1715_v4;
      }
      let _1781_v60;
      _1781_v60 = _dafny.Set.fromElements((_1720_v9).f19, true, (_1720_v9).f19, true, true);
      r0 = ((((((_1717_v6).contains(new BigNumber((_1781_v60).length))) ? ((_1717_v6).get(new BigNumber((_1781_v60).length))) : (p1))).isLessThanOrEqualTo(p1)) ? (p1) : (p1));
      r1 = ((p1).plus(p1)).plus(new BigNumber((_dafny.Seq.Concat(_1712_v1, _1712_v1)).length));
      let _1782_v61;
      let _nw290 = new _module.C0();
      _nw290.__ctor((_1710_v0)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1710_v0).length))], !(true));
      _1782_v61 = _nw290;
      let _1783_v62;
      _1783_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1782_v61,_1781_v60);
      r2 = (_1783_v62).equals(_1783_v62);
      return [r0, r1, r2];
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.MultiSet.Empty;
      let r3 = false;
      let _1784_v0;
      let _init50 = function (_1785_i1) {
        return _this.f14;
      };
      let _nw291 = Array((new BigNumber(19)).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw291.length); _i0_50++) {
        _nw291[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _1784_v0 = _nw291;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1784_v0).length))) {
        let _1786_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1786_i0)) && ((_1786_i0).isLessThan(new BigNumber((_1784_v0).length))))) {
          (_1784_v0)[(_1786_i0)] = !(_module.__default.safeModuloInt(new BigNumber(-320), new BigNumber((_dafny.Seq.UnicodeFromString("vuy")).length))).isEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-304)), function (_1787_i2) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("qdea"))).length));
        }
      }
      if (_this.f14) {
        r3 = _this.f14;
        r3 = _this.f14;
        let _1788_v1;
        _1788_v1 = _module.D0.create_DC0(_this.f14);
        let _source17 = _1788_v1;
        if (_source17.is_DC1) {
          let _1789___mcc_h0 = (_source17).cf1;
          let _1790_cf1 = _1789___mcc_h0;
          let _1791_v2;
          _1791_v2 = new BigNumber(-981);
          r0 = _1791_v2;
          let _1792_v3;
          _1792_v3 = _dafny.Seq.UnicodeFromString("ikpk");
          let _1793_v4;
          _1793_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1790_cf1,_1792_v3);
          let _1794_v5;
          _1794_v5 = _dafny.Seq.of(_1792_v3, _1792_v3, _1792_v3, _1792_v3, _1792_v3);
          let _1795_v6;
          _1795_v6 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1796_v7;
          _1796_v7 = _dafny.Set.fromElements(_1791_v2);
          let _1797_v8;
          _1797_v8 = _module.D9.create_DC30(_1795_v6, _1790_cf1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), ((_1798_v6) => function (_1799_i3) {
  return _1798_v6;
})(_1795_v6))).length), _1796_v7, _1791_v2);
          let _1800_v9;
          let _nw292 = new _module.C3();
          _nw292.__ctor(new BigNumber(((_1793_v4).Merge(_1793_v4)).length), (_1794_v5)[_module.__default.safeIndex(_1791_v2, new BigNumber((_1794_v5).length))], ((_1790_cf1) ? ((_1797_v8).dtor_cf51) : (_1790_cf1)), (_this).f15);
          _1800_v9 = _nw292;
          let _1801_v10;
          let _nw293 = new _module.C3();
          _nw293.__ctor(new BigNumber((_dafny.Seq.of((_1800_v9).f20, (_1800_v9).f20, (_1800_v9).f20)).length), _dafny.Seq.Concat(_1792_v3, _1792_v3), true, _dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f14, !(_1790_cf1), _this.f14)));
          _1801_v10 = _nw293;
          let _1802_v11;
          _1802_v11 = _dafny.Seq.of((_1801_v10).f20, _1791_v2);
          r0 = new BigNumber((_dafny.Seq.Concat(_1802_v11, _dafny.Seq.Concat(_1802_v11, _1802_v11))).length);
        } else if (_source17.is_DC0) {
          let _1803___mcc_h1 = (_source17).cf0;
          let _1804_cf0 = _1803___mcc_h1;
          let _1805_v12;
          _1805_v12 = new BigNumber(83);
          let _1806_v13;
          let _nw294 = new _module.C5();
          _nw294.__ctor((_1805_v12).plus(_1805_v12), _this.f14, _module.__default.fm33(_this.f14, globalState));
          _1806_v13 = _nw294;
          let _1807_v14;
          _1807_v14 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1808_v15;
          _1808_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1807_v14,new BigNumber(66));
          let _1809_v16;
          _1809_v16 = _dafny.Seq.of((_1806_v13).f18);
          let _1810_v17;
          _1810_v17 = _dafny.Set.fromElements(_1809_v16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), ((_1811_v13) => function (_1812_i4) {
            return (_1811_v13).f18;
          })(_1806_v13)), _1809_v16, _1809_v16, _1809_v16);
          let _1813_v18;
          let _nw295 = new _module.C4();
          _nw295.__ctor(_1804_cf0, !(_1804_cf0) || (!(_module.__default.fm23(_1805_v12, _1804_cf0, _1808_v15, _1810_v17, globalState))), (_this).f15);
          _1813_v18 = _nw295;
          let _1814_v19;
          _1814_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1805_v12,_this.f14);
          let _1815_v20;
          let _nw296 = new _module.C2();
          _nw296.__ctor(_1814_v19);
          _1815_v20 = _nw296;
          (globalState).f3 = _1805_v12;
        } else {
          let _1816___mcc_h2 = (_source17).cf2;
          let _1817_cf2 = _1816___mcc_h2;
          let _1818_v21;
          _1818_v21 = _dafny.Set.fromElements(_this.f14);
          r1 = new BigNumber((_1818_v21).length);
          let _1819_v22;
          _1819_v22 = new BigNumber(999);
          let _1820_v23;
          _1820_v23 = _module.D12.create_DC38(_1819_v22);
          _1820_v23 = _module.D12.create_DC38(_1819_v22);
          (globalState).f2 = _1784_v0;
          let _1821_v24;
          let _init51 = ((_1822_v21) => function (_1823_i5) {
            return _module.D8.create_DC24(_1822_v21);
          })(_1818_v21);
          let _nw297 = Array((new BigNumber(9)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw297.length); _i0_51++) {
            _nw297[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1821_v24 = _nw297;
          let _1824_v25;
          _1824_v25 = _module.D8.create_DC24(_1818_v21);
          let _index325 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1821_v24).length));
          (_1821_v24)[_index325] = function (_pat_let58_0) {
            return function (_1825_dt__update__tmp_h0) {
              return function (_pat_let59_0) {
                return function (_1826_dt__update_hcf45_h0) {
                  return _module.D8.create_DC24(_1826_dt__update_hcf45_h0);
                }(_pat_let59_0);
              }(_dafny.Set.fromElements(_this.f14));
            }(_pat_let58_0);
          }(_1824_v25);
          let _index326 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1821_v24).length));
          (_1821_v24)[_index326] = _1824_v25;
        }
        let _1827_v26;
        let _init52 = function (_1828_i6) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tqux"), _dafny.Seq.UnicodeFromString("mdsulisdp"));
        };
        let _nw298 = Array((new BigNumber(11)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw298.length); _i0_52++) {
          _nw298[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1827_v26 = _nw298;
        let _1829_v27;
        _1829_v27 = _dafny.Set.fromElements((_this).f15);
        let _1830_v28;
        _1830_v28 = _module.D11.create_DC33(_1829_v27);
        let _pat_let_tv43 = _1829_v27;
        let _rhs298 = _1827_v26;
        let _rhs299 = function (_pat_let60_0) {
          return function (_1831_dt__update__tmp_h1) {
            return function (_pat_let61_0) {
              return function (_1832_dt__update_hcf57_h0) {
                return _module.D11.create_DC33(_1832_dt__update_hcf57_h0);
              }(_pat_let61_0);
            }(_pat_let_tv43);
          }(_pat_let60_0);
        }(_module.D11.create_DC33(_1829_v27));
        _1827_v26 = _rhs298;
        _1830_v28 = _rhs299;
        let _1833_v29;
        let _nw299 = new _module.C1();
        _nw299.__ctor();
        _1833_v29 = _nw299;
      } else {
        let _1834_v30;
        _1834_v30 = new BigNumber(140);
        let _1835_v31;
        _1835_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1834_v30,new BigNumber(41));
        _1835_v31 = (_1835_v31).update(_1834_v30, (_1834_v30).multipliedBy(_1834_v30));
        let _1836_v32;
        _1836_v32 = _dafny.Seq.UnicodeFromString("shmrn");
        if (!(_this.f14) || (!(_dafny.Seq.IsPrefixOf(_1836_v32, _1836_v32)))) {
          let _1837_v33;
          let _nw300 = new _module.C0();
          _nw300.__ctor(_this.f14, _this.f14);
          _1837_v33 = _nw300;
          _1834_v30 = new BigNumber(3);
          let _1838_v34;
          _1838_v34 = _dafny.Set.fromElements(_module.__default.fm27(_1834_v30, _this.f14, globalState));
          _1838_v34 = ((_1838_v34).Union(_1838_v34)).Difference((((_1837_v33).f24) ? (_1838_v34) : (_1838_v34)));
          let _1839_v35;
          _1839_v35 = new _dafny.CodePoint('n'.codePointAt(0));
          _1839_v35 = _1839_v35;
          let _1840_v36;
          _1840_v36 = _dafny.MultiSet.fromElements(_1834_v30);
          let _1841_v37;
          _1841_v37 = _dafny.Seq.of((_1837_v33).f23);
          let _1842_v38;
          _1842_v38 = _dafny.Set.fromElements((_dafny.MultiSet.fromElements(_this.f14, (_1837_v33).f24)).update(_this.f14, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), ((_1843_v35) => function (_1844_i7) {
            return _1843_v35;
          })(_1839_v35))).length))), (_this).f15, _dafny.MultiSet.fromElements((_1837_v33).f23), _dafny.MultiSet.FromArray(_1841_v37));
          let _rhs300 = _1834_v30;
          let _rhs301 = ((new BigNumber((_1836_v32).length)).minus(new BigNumber((_1840_v36).cardinality()))).plus((_1834_v30).minus(new BigNumber((_1836_v32).length)));
          let _rhs302 = (_1842_v38).IsDisjointFrom(_1842_v38);
          let _lhs256 = globalState;
          let _lhs257 = globalState;
          _lhs256.f3 = _rhs300;
          r0 = _rhs301;
          _lhs257.f9 = _rhs302;
        } else {
          let _1845_v39;
          _1845_v39 = _dafny.Map.Empty.slice().updateUnsafe(false,_1834_v30);
          let _1846_v40;
          _1846_v40 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1847_v41;
          _1847_v41 = _dafny.Set.fromElements(_module.__default.fm25(_1834_v30, _1834_v30, _this.f14, _1845_v39, globalState), _1846_v40);
          let _1848_v42;
          _1848_v42 = _dafny.Seq.of(_1847_v41);
          let _1849_v44;
          _1849_v44 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),_this.f14);
          let _1850_v45;
          _1850_v45 = _dafny.Seq.of(_1834_v30, _1834_v30);
          let _1851_v46;
          _1851_v46 = _dafny.Set.fromElements(_1850_v45);
          let _rhs303 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1834_v30).plus(new BigNumber(658))));
          let _rhs304 = !((_module.__default.fm23(_1834_v30, _this.f14, function () {
            let _coll73 = new _dafny.Map();
            for (const _compr_73 of (_1849_v44).Keys.Elements) {
              let _1852_v43 = _compr_73;
              if ((_1849_v44).contains(_1852_v43)) {
                _coll73.push([_1852_v43,new BigNumber((_dafny.Seq.of(_1834_v30, _1834_v30)).length)]);
              }
            }
            return _coll73;
          }(), _1851_v46, globalState)) && (!(_this.f14) || (_this.f14)));
          let _rhs305 = (((_1835_v31).contains(new BigNumber((_1836_v32).length))) ? ((_1835_v31).get(new BigNumber((_1836_v32).length))) : (_1834_v30));
          let _rhs306 = _1848_v42;
          let _lhs258 = globalState;
          let _lhs259 = globalState;
          _lhs258.f3 = _rhs303;
          _lhs259.f9 = _rhs304;
          _1834_v30 = _rhs305;
          _1848_v42 = _rhs306;
          let _1853_v47;
          _1853_v47 = _dafny.Set.fromElements(!(_this.f14));
          let _1854_v48;
          _1854_v48 = _dafny.Seq.of(_this.f14, _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("dpidlr"), _module.__default.fm25(_1834_v30, new BigNumber((_1853_v47).length), _this.f14, _1845_v39, globalState)), _this.f14, _this.f14);
          _1854_v48 = _1854_v48;
          let _1855_v49;
          _1855_v49 = _module.D0.create_DC1(_this.f14);
          let _1856_v50;
          _1856_v50 = _module.D0.create_DC2(_1855_v49);
          let _1857_v51;
          _1857_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1856_v50,_this.f14);
          let _1858_v52;
          _1858_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1834_v30,(_this).fm11(_1857_v51, _1834_v30, globalState));
          let _1859_v53;
          _1859_v53 = _dafny.Seq.of(_1858_v52);
          let _1860_v54;
          let _nw301 = new _module.C2();
          _nw301.__ctor(((true) ? ((_1859_v53)[_module.__default.safeIndex(_1834_v30, new BigNumber((_1859_v53).length))]) : (_1858_v52)));
          _1860_v54 = _nw301;
          let _1861_v55;
          _1861_v55 = _dafny.MultiSet.fromElements(_1834_v30, _1834_v30);
          r0 = _module.__default.fm0(_this.f14, _1834_v30, _1834_v30, _1861_v55, globalState);
          let _1862_v56;
          let _nw302 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1862_v56 = _nw302;
          let _index327 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1862_v56).length));
          (_1862_v56)[_index327] = _1834_v30;
          let _index328 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1862_v56).length));
          (_1862_v56)[_index328] = (_1834_v30).plus(_1834_v30);
        }
        let _1863_v57;
        _1863_v57 = _dafny.Seq.of(_this.f14);
        let _1864_v58;
        _1864_v58 = _dafny.Seq.of(_1863_v57);
        if (!_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(true, true, _this.f14, _this.f14), (_1864_v58)[_module.__default.safeIndex(new BigNumber(654), new BigNumber((_1864_v58).length))]), _dafny.Seq.update(_1863_v57, _module.__default.safeIndex(_1834_v30, new BigNumber((_1863_v57).length)), _this.f14))) {
          let _1865_v59;
          _1865_v59 = _dafny.MultiSet.fromElements(_1834_v30, new BigNumber(943), _1834_v30, new BigNumber(97));
          r1 = (((_1865_v59).contains(_module.__default.safeModuloInt(_1834_v30, _module.__default.fm0(_this.f14, new BigNumber((_1836_v32).length), _1834_v30, _1865_v59, globalState)))) ? ((_1865_v59).get(_module.__default.safeModuloInt(_1834_v30, _module.__default.fm0(_this.f14, new BigNumber((_1836_v32).length), _1834_v30, _1865_v59, globalState)))) : ((_dafny.ZERO).minus(_1834_v30)));
          let _1866_v60;
          let _nw303 = new _module.C3();
          _nw303.__ctor(_module.__default.safeModuloInt(_1834_v30, _1834_v30), _1836_v32, !(_this.f14) || (_this.f14), ((_this).f15).update(_this.f14, _module.__default.abs(_1834_v30)));
          _1866_v60 = _nw303;
          _1836_v32 = _dafny.Seq.UnicodeFromString("srr");
          let _1867_v61;
          _1867_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1866_v60).f20,((_this.f14) ? (_this.f14) : (_this.f14)));
          _1867_v61 = (_1867_v61).update(((_this.f14) ? (new BigNumber(737)) : (_1834_v30)), true);
          let _1868_v62;
          _1868_v62 = _module.D5.create_DC18(_this.f14);
          (globalState).f9 = (_1868_v62).dtor_cf33;
        } else {
          (globalState).f9 = ((_dafny.ZERO).minus((_1834_v30).minus(new BigNumber(-320)))).isLessThan(_1834_v30);
          let _1869_v63;
          _1869_v63 = _module.D1.create_DC5(_this.f14, _this.f14);
          let _1870_v64;
          let _nw304 = new _module.C5();
          _nw304.__ctor(new BigNumber((_1863_v57).length), (_1869_v63).dtor_cf6, (_this).f15);
          _1870_v64 = _nw304;
          let _1871_v65;
          _1871_v65 = _module.D15.create_DC48(_1834_v30, _1870_v64, new BigNumber(491));
          (globalState).f3 = (_1871_v65).dtor_cf79;
          let _1872_v66;
          _1872_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1784_v0,_this.f14);
          let _index329 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1784_v0).length));
          (_1784_v0)[_index329] = !((((_1872_v66).contains(_1784_v0)) ? ((_1872_v66).get(_1784_v0)) : (_this.f14)));
          let _index330 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1784_v0).length));
          (_1784_v0)[_index330] = _this.f14;
          let _1873_v67;
          _1873_v67 = _module.D0.create_DC0((_1784_v0)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1784_v0).length))]);
          (globalState).f3 = (((_1873_v67).dtor_cf0) ? (_1834_v30) : ((_dafny.ZERO).minus((_1870_v64).f18)));
          let _1874_v68;
          _1874_v68 = _module.D7.create_DC23(_this.f14, _1834_v30, _this.f14);
          _1874_v68 = _1874_v68;
        }
        (globalState).f2 = _1784_v0;
        let _1875_v69;
        _1875_v69 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1876_v70;
        _1876_v70 = _module.D2.create_DC9(_this.f14, _1834_v30, new _dafny.CodePoint('v'.codePointAt(0)));
        (globalState).f9 = (((_this.f14) ? (_module.D2.create_DC9(_this.f14, _1834_v30, _1875_v69)) : (_1876_v70))).dtor_cf13;
      }
      let _1877_v71;
      let _nw305 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1877_v71 = _nw305;
      let _1878_v72;
      _1878_v72 = _dafny.Seq.UnicodeFromString("hau");
      let _index331 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1877_v71).length));
      (_1877_v71)[_index331] = _1878_v72;
      let _1879_v73;
      _1879_v73 = new BigNumber(437);
      let _index332 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1877_v71).length));
      (_1877_v71)[_index332] = _dafny.Seq.Concat(_dafny.Seq.Concat(_1878_v72, _dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_1880_i8) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })), _dafny.Seq.Concat(_module.__default.fm27(_1879_v73, _this.f14, globalState), _1878_v72));
      let _1881_v74;
      _1881_v74 = _dafny.Set.fromElements(_1879_v73);
      let _hi7 = _1879_v73;
      for (let _1882_i9 = new BigNumber((_1881_v74).length); _1882_i9.isLessThan(_hi7); _1882_i9 = _1882_i9.plus(_dafny.ONE)) {
        let _1883_v75;
        _1883_v75 = _module.D5.create_DC18(!(true));
        let _1884_v76;
        _1884_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1882_i9,_1883_v75);
        _1884_v76 = (_1884_v76).update((new BigNumber((_1881_v74).length)).minus(_1879_v73), _1883_v75);
        let _index333 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1784_v0).length));
        (_1784_v0)[_index333] = ((_this.f14) ? (_this.f14) : (!(_this.f14)));
        let _index334 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1784_v0).length));
        (_1784_v0)[_index334] = _this.f14;
        let _1885_v77;
        _1885_v77 = new _dafny.CodePoint('h'.codePointAt(0));
        let _index335 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1784_v0).length));
        let _index336 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1784_v0).length));
        let _index337 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1877_v71).length));
        let _rhs307 = true;
        let _rhs308 = _this.f14;
        let _rhs309 = _dafny.Seq.Concat(_dafny.Seq.update(_1878_v72, _module.__default.safeIndex(_1879_v73, new BigNumber((_1878_v72).length)), _1885_v77), (_1877_v71)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_1877_v71).length))]);
        let _lhs260 = _1784_v0;
        let _lhs261 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1784_v0).length));
        let _lhs262 = _1784_v0;
        let _lhs263 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1784_v0).length));
        let _lhs264 = _1877_v71;
        let _lhs265 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_1877_v71).length));
        _lhs260[_lhs261] = _rhs307;
        _lhs262[_lhs263] = _rhs308;
        _lhs264[_lhs265] = _rhs309;
        let _1886_v78;
        _1886_v78 = _dafny.Seq.of(true, _this.f14);
        (globalState).f3 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1886_v78, _1886_v78), _1886_v78)).length);
        let _1887_v79;
        _1887_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1882_i9,_1882_i9);
        let _1888_v80;
        _1888_v80 = _module.D9.create_DC29(_1887_v79);
        let _1889_v81;
        _1889_v81 = _module.D9.create_DC31(_1888_v80);
        let _1890_v82;
        _1890_v82 = _module.D9.create_DC31(_1889_v81);
        let _1891_v83;
        _1891_v83 = _dafny.Seq.of(_1890_v82);
        let _1892_v84;
        _1892_v84 = _dafny.Set.fromElements(_1891_v83);
        _1892_v84 = _module.__default.fm49(_1879_v73, globalState);
      }
      let _1893_v85;
      let _init53 = ((_1894_v73) => function (_1895_i10) {
        return _module.__default.safeModuloInt(_1895_i10, _1894_v73);
      })(_1879_v73);
      let _nw306 = Array((new BigNumber(22)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw306.length); _i0_53++) {
        _nw306[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _1893_v85 = _nw306;
      let _index338 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length));
      (_1893_v85)[_index338] = _1879_v73;
      let _index339 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length));
      (_1893_v85)[_index339] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(((_this.f14) ? (_1879_v73) : (_1879_v73)), ((_this.f14) ? (new BigNumber(284)) : (new BigNumber(269)))));
      let _1896_v86;
      _1896_v86 = _dafny.Seq.of((_1893_v85)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length))]);
      let _1897_v87;
      _1897_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1896_v86, _dafny.Seq.of((_1893_v85)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length))]))).length),_this.f14);
      _1897_v87 = (_1897_v87).update(_1879_v73, ((_1893_v85)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length))]).isLessThanOrEqualTo(new BigNumber(510)));
      r0 = (_1893_v85)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length))];
      r1 = (_1879_v73).multipliedBy(_1879_v73);
      let _1898_v88;
      _1898_v88 = _dafny.MultiSet.fromElements(_1879_v73, new BigNumber((_1878_v72).length), _1879_v73, _module.__default.safeModuloInt(_1879_v73, (_1893_v85)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1893_v85).length))]));
      r2 = _1898_v88;
      let _1899_v90;
      _1899_v90 = _dafny.Seq.of(_1897_v87);
      let _1900_v91;
      _1900_v91 = _module.D7.create_DC21(function () {
  let _coll74 = new _dafny.Map();
  for (const _compr_74 of (_1899_v90).Elements) {
    let _1901_v89 = _compr_74;
    if (_dafny.Seq.contains(_1899_v90, _1901_v89)) {
      _coll74.push([_1901_v89,_1879_v73]);
    }
  }
  return _coll74;
}());
      r3 = function (_source18) {
        if (_source18.is_DC22) {
          let _1902___mcc_h3 = (_source18).cf37;
          let _1903___mcc_h4 = (_source18).cf38;
          let _1904___mcc_h5 = (_source18).cf39;
          let _1905___mcc_h6 = (_source18).cf40;
          let _1906___mcc_h7 = (_source18).cf41;
          let _1907_cf41 = _1906___mcc_h7;
          let _1908_cf40 = _1905___mcc_h6;
          let _1909_cf39 = _1904___mcc_h5;
          let _1910_cf38 = _1903___mcc_h4;
          let _1911_cf37 = _1902___mcc_h3;
          return !(_this.f14);
        } else if (_source18.is_DC23) {
          let _1912___mcc_h8 = (_source18).cf42;
          let _1913___mcc_h9 = (_source18).cf43;
          let _1914___mcc_h10 = (_source18).cf44;
          let _1915_cf44 = _1914___mcc_h10;
          let _1916_cf43 = _1913___mcc_h9;
          let _1917_cf42 = _1912___mcc_h8;
          return _1915_cf44;
        } else {
          let _1918___mcc_h11 = (_source18).cf36;
          let _1919_cf36 = _1918___mcc_h11;
          return _this.f14;
        }
      }(_1900_v91);
      return [r0, r1, r2, r3];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm9(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kko"), _dafny.Seq.UnicodeFromString("ddx"))).length);
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return false;
    };
    m3(p0, globalState) {
      let _this = this;
      let _hi8 = p0;
      for (let _1920_i0 = p0; _1920_i0.isLessThan(_hi8); _1920_i0 = _1920_i0.plus(_dafny.ONE)) {
        let _1921_v0;
        _1921_v0 = _module.D1.create_DC3(new BigNumber(-181));
        let _1922_v1;
        _1922_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,_1920_i0);
        let _1923_v2;
        _1923_v2 = _dafny.Seq.of(new BigNumber((_1922_v1).length), p0);
        let _1924_v3;
        _1924_v3 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1925_v4;
        _1925_v4 = _dafny.Seq.of(_1924_v3);
        let _1926_v5;
        let _nw307 = Array((new BigNumber(15)).toNumber());
        _nw307[(_dafny.ZERO).toNumber()] = _1923_v2;
        _nw307[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1923_v2, _1923_v2);
        _nw307[(new BigNumber(2)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(3)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(4)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_1923_v2, _module.__default.safeIndex(_1920_i0, new BigNumber((_1923_v2).length)), (_1921_v0).dtor_cf3);
        _nw307[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(new BigNumber((_1925_v4).length), p0);
        _nw307[(new BigNumber(7)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(8)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(9)).toNumber()] = ((true) ? (_1923_v2) : (_1923_v2));
        _nw307[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_1923_v2, _module.__default.safeIndex(_1920_i0, new BigNumber((_1923_v2).length)), _1920_i0);
        _nw307[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1923_v2, _dafny.Seq.of(p0, p0, new BigNumber((_1925_v4).length)));
        _nw307[(new BigNumber(12)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(13)).toNumber()] = _1923_v2;
        _nw307[(new BigNumber(14)).toNumber()] = _1923_v2;
        _1926_v5 = _nw307;
        let _index340 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1926_v5).length));
        (_1926_v5)[_index340] = _dafny.Seq.update(_1923_v2, _module.__default.safeIndex(_1920_i0, new BigNumber((_1923_v2).length)), p0);
        let _1927_v6;
        let _nw308 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1927_v6 = _nw308;
        let _1928_v7;
        _1928_v7 = true;
        let _index341 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1927_v6).length));
        (_1927_v6)[_index341] = _1928_v7;
        let _index342 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1926_v5).length));
        let _index343 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1927_v6).length));
        let _rhs310 = _1928_v7;
        let _rhs311 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1923_v2, _1923_v2), _dafny.Seq.Concat(_1923_v2, _1923_v2));
        let _rhs312 = !(_1928_v7);
        let _lhs266 = globalState;
        let _lhs267 = _1926_v5;
        let _lhs268 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1926_v5).length));
        let _lhs269 = _1927_v6;
        let _lhs270 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1927_v6).length));
        _lhs266.f9 = _rhs310;
        _lhs267[_lhs268] = _rhs311;
        _lhs269[_lhs270] = _rhs312;
        let _1929_v8;
        _1929_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1924_v3);
        let _1930_v9;
        _1930_v9 = _dafny.MultiSet.fromElements(_1929_v8);
        let _1931_v10;
        let _nw309 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
        _1931_v10 = _nw309;
        let _1932_v11;
        _1932_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), ((_1933_v6) => function (_1934_i1) {
          return new BigNumber((_dafny.Seq.of((_1933_v6)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1933_v6).length))])).length);
        })(_1927_v6)),_1924_v3);
        let _1935_v12;
        _1935_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1928_v7,_1927_v6);
        let _rhs313 = (_1930_v9).Difference(_1930_v9);
        let _rhs314 = !(_1928_v7);
        let _rhs315 = (((_1935_v12).contains((_1927_v6)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1927_v6).length))])) ? ((_1935_v12).get((_1927_v6)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1927_v6).length))])) : (_1927_v6));
        let _rhs316 = _1931_v10;
        let _rhs317 = _dafny.Map.Empty.slice().updateUnsafe((_1926_v5)[_module.__default.safeIndex(new BigNumber(789), new BigNumber((_1926_v5).length))],_1924_v3);
        let _lhs271 = globalState;
        _1930_v9 = _rhs313;
        _1928_v7 = _rhs314;
        _lhs271.f2 = _rhs315;
        _1931_v10 = _rhs316;
        _1932_v11 = _rhs317;
        _1925_v4 = _dafny.Seq.Concat(_1925_v4, _dafny.Seq.UnicodeFromString("reyaln"));
        (globalState).f9 = !(_1920_i0).isEqualTo(_1920_i0);
      }
      let _hi9 = p0;
      for (let _1936_i2 = new BigNumber(-126); _1936_i2.isLessThan(_hi9); _1936_i2 = _1936_i2.plus(_dafny.ONE)) {
        let _1937_v13;
        let _nw310 = Array((new BigNumber(2)).toNumber()).fill(_module.D1.Default());
        _1937_v13 = _nw310;
        let _1938_v14;
        _1938_v14 = _module.D1.create_DC4(new BigNumber(573), p0);
        let _index344 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1937_v13).length));
        (_1937_v13)[_index344] = function (_pat_let62_0) {
          return function (_1939_dt__update__tmp_h0) {
            return function (_pat_let63_0) {
              return function (_1940_dt__update_hcf4_h0) {
                return _module.D1.create_DC4(_1940_dt__update_hcf4_h0, (_1939_dt__update__tmp_h0).dtor_cf5);
              }(_pat_let63_0);
            }(_1936_i2);
          }(_pat_let62_0);
        }(_1938_v14);
        let _index345 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1937_v13).length));
        (_1937_v13)[_index345] = _1938_v14;
        let _1941_v15;
        _1941_v15 = true;
        let _1942_v16;
        let _nw311 = new _module.C0();
        _nw311.__ctor(_1941_v15, _1941_v15);
        _1942_v16 = _nw311;
        let _1943_v17;
        _1943_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-189),p0);
        (globalState).f9 = !(_1943_v17).equals(_module.__default.fm14(globalState));
        let _1944_v18;
        _1944_v18 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1945_v19;
        _1945_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1944_v18,(_dafny.ZERO).minus(_1936_i2));
        let _1946_v20;
        _1946_v20 = _dafny.Seq.of(_1936_i2, _1936_i2);
        let _1947_v21;
        _1947_v21 = _dafny.Set.fromElements(_1946_v20, _1946_v20);
        let _1948_v22;
        let _init54 = ((_1949_v16) => function (_1950_i3) {
          return (_1949_v16).f24;
        })(_1942_v16);
        let _nw312 = Array((new BigNumber(6)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw312.length); _i0_54++) {
          _nw312[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1948_v22 = _nw312;
        (globalState).f2 = ((!(((_module.__default.fm23(p0, _1941_v15, _1945_v19, _1947_v21, globalState)) ? (_1941_v15) : (false)))) ? (_1948_v22) : (_1948_v22));
      }
      let _1951_v23;
      _1951_v23 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), new BigNumber(704));
      (globalState).f9 = !(!(((_1951_v23).Intersect(_1951_v23)).IsDisjointFrom(_1951_v23)));
      let _1952_i4;
      _1952_i4 = _dafny.ZERO;
      L4: {
        while (!(p0).isEqualTo(p0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1952_i4)) {
              break L4;
            }
            _1952_i4 = (_1952_i4).plus(_dafny.ONE);
            let _1953_v24;
            let _nw313 = Array((new BigNumber(14)).toNumber()).fill(false);
            _1953_v24 = _nw313;
            let _1954_v25;
            _1954_v25 = _dafny.Seq.of(_1953_v24, _1953_v24, _1953_v24, _1953_v24, _1953_v24);
            let _1955_v26;
            _1955_v26 = false;
            _1954_v25 = ((_1955_v26) ? (_1954_v25) : (_1954_v25));
            _1955_v26 = _1955_v26;
            let _1956_v27;
            _1956_v27 = _dafny.Seq.of(_1955_v26);
            let _1957_v28;
            _1957_v28 = _dafny.Seq.of(p0, new BigNumber((_dafny.Seq.update(_1956_v27, _module.__default.safeIndex(p0, new BigNumber((_1956_v27).length)), !(_1955_v26))).length), new BigNumber(-573), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1955_v26,_1955_v26)).length), (_dafny.ZERO).minus(p0));
            let _1958_v29;
            _1958_v29 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_1959_i5) {
              return new BigNumber(506);
            }), _1957_v28);
            let _source19 = _1958_v29;
            let _1960___mcc_h0 = _source19;
            let _1961_cf56 = _1960___mcc_h0;
            let _1962_v30;
            let _nw314 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1962_v30 = _nw314;
            let _1963_v31;
            _1963_v31 = _dafny.Seq.of(_1962_v30);
            let _1964_v32;
            _1964_v32 = _module.D16.create_DC51(_1962_v30);
            _1963_v31 = _dafny.Seq.of(_1962_v30, (_1964_v32).dtor_cf83);
            let _1965_v33;
            _1965_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1958_v29,p0);
            let _1966_v34;
            _1966_v34 = _dafny.MultiSet.fromElements(_1965_v33);
            (globalState).f9 = (_1966_v34).contains(_1965_v33);
            let _1967_v35;
            _1967_v35 = _dafny.Set.fromElements(_1955_v26);
            (globalState).f6 = (_1967_v35).Intersect(_module.__default.fm41(p0, p0, globalState));
            let _1968_v36;
            _1968_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_1955_v26));
            let _1969_v37;
            _1969_v37 = _dafny.Seq.UnicodeFromString("o");
            let _1970_v38;
            _1970_v38 = _dafny.MultiSet.fromElements(_1955_v26);
            let _1971_v39;
            let _nw315 = new _module.C3();
            _nw315.__ctor(p0, _1969_v37, _1955_v26, _1970_v38);
            _1971_v39 = _nw315;
            let _1972_v40;
            _1972_v40 = _dafny.Seq.of(_1971_v39, _1971_v39, _1971_v39, _1971_v39, _1971_v39);
            let _1973_v41;
            _1973_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1972_v40,_1970_v38);
            let _1974_v42;
            _1974_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1955_v26,_1973_v41);
            let _1975_v43;
            _1975_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1956_v27).length),_1973_v41);
            _1968_v36 = (_1968_v36).update((_dafny.ZERO).minus(new BigNumber((((((_1974_v42).contains(_1955_v26)) ? ((_1974_v42).get(_1955_v26)) : ((((_1975_v43).contains(p0)) ? ((_1975_v43).get(p0)) : (_dafny.Map.Empty.slice().updateUnsafe(_1972_v40,_1970_v38)))))).update(_1972_v40, _1970_v38)).length)), _1955_v26);
            let _1976_v44;
            _1976_v44 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(866), p0));
            let _1977_v45;
            _1977_v45 = _module.D13.create_DC41(_1976_v44);
            let _source20 = _1977_v45;
            if (_source20.is_DC42) {
              let _1978___mcc_h1 = (_source20).cf68;
              let _1979___mcc_h2 = (_source20).cf69;
              let _1980_cf69 = _1979___mcc_h2;
              let _1981_cf68 = _1978___mcc_h1;
              let _1982_v46;
              _1982_v46 = _dafny.Set.fromElements(_1980_cf69, _1980_cf69, _1980_cf69);
              (globalState).f6 = ((_1955_v26) ? (_1982_v46) : (_1982_v46));
              let _1983_v47;
              let _nw316 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
              _1983_v47 = _nw316;
              let _index346 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1983_v47).length));
              (_1983_v47)[_index346] = (_this).fm9(p0, globalState);
              let _index347 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1983_v47).length));
              (_1983_v47)[_index347] = new BigNumber((_dafny.Seq.Concat(_1954_v25, _dafny.Seq.Concat(_1954_v25, _1954_v25))).length);
              let _1984_v48;
              _1984_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1980_cf69,_dafny.Seq.Concat(_1956_v27, _dafny.Seq.update(_dafny.Seq.of(_1955_v26), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_1955_v26)).length)), _1980_cf69)));
              let _1985_v49;
              _1985_v49 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),(_1983_v47)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1983_v47).length))]);
              let _1986_v50;
              _1986_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1980_cf69,new BigNumber(202));
              let _1987_v51;
              _1987_v51 = _dafny.Seq.UnicodeFromString("jrpyg");
              _1984_v48 = (_1984_v48).update(_module.__default.fm23((_1983_v47)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1983_v47).length))], false, _1985_v49, _dafny.Set.fromElements(_1957_v28, _1957_v28, _1957_v28), globalState), (((_1984_v48).contains(_1955_v26)) ? ((_1984_v48).get(_1955_v26)) : (_dafny.Seq.of(_1955_v26, _1980_cf69, !((_this).fm10(_1986_v50, _1951_v23, false, _1987_v51, globalState))))));
              let _1988_v52;
              _1988_v52 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("sct"));
              let _1989_v53;
              _1989_v53 = _module.D4.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), ((_1990_cf68) => function (_1991_i8) {
  return _1990_cf68;
})(_1981_cf68)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-776)), ((_1992_cf68, _1993_cf69, _1994_p0, _1995_v47) => function (_1996_i9) {
  return (_module.D13.create_DC42(_1992_cf68, (_module.D9.create_DC30(_1992_cf68, _1993_cf69, _1994_p0, _dafny.Set.fromElements((_1995_v47)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1995_v47).length))]), (_1995_v47)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1995_v47).length))])).dtor_cf51)).dtor_cf68;
})(_1981_cf68, _1980_cf69, p0, _1983_v47))).length), _1955_v26, _1953_v24);
              let _1997_v54;
              let _nw317 = Array((new BigNumber(23)).toNumber());
              _nw317[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_1988_v52)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1988_v52).length))], _1987_v51);
              _nw317[(_dafny.ONE).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(2)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1987_v51, _1987_v51);
              _nw317[(new BigNumber(4)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1987_v51, _1987_v51);
              _nw317[(new BigNumber(6)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1987_v51, _1987_v51);
              _nw317[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1987_v51, _module.__default.safeIndex((_1983_v47)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1983_v47).length))], new BigNumber((_1987_v51).length)), _1981_cf68);
              _nw317[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_1998_cf68) => function (_1999_i6) {
                return _1998_cf68;
              })(_1981_cf68));
              _nw317[(new BigNumber(10)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), ((_2000_cf68) => function (_2001_i7) {
                return _2000_cf68;
              })(_1981_cf68));
              _nw317[(new BigNumber(12)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(13)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(14)).toNumber()] = (_1989_v53).dtor_cf22;
              _nw317[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_2002_cf68) => function (_2003_i10) {
                return _2002_cf68;
              })(_1981_cf68));
              _nw317[(new BigNumber(16)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(17)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("abhvu");
              _nw317[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1987_v51, _dafny.Seq.UnicodeFromString("u"));
              _nw317[(new BigNumber(20)).toNumber()] = _1987_v51;
              _nw317[(new BigNumber(21)).toNumber()] = ((_1980_cf69) ? (_1987_v51) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(390)), function (_2004_i11) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              })));
              _nw317[(new BigNumber(22)).toNumber()] = _1987_v51;
              _1997_v54 = _nw317;
              let _index348 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_1997_v54).length));
              (_1997_v54)[_index348] = _dafny.Seq.UnicodeFromString("i");
              let _index349 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_1997_v54).length));
              (_1997_v54)[_index349] = _dafny.Seq.Concat(_1987_v51, _dafny.Seq.UnicodeFromString("k"));
            } else if (_source20.is_DC41) {
              let _2005___mcc_h3 = (_source20).cf67;
              let _2006_cf67 = _2005___mcc_h3;
              let _2007_v55;
              _2007_v55 = _dafny.Seq.UnicodeFromString("gv");
              let _2008_v56;
              _2008_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2007_v55).length));
              (globalState).f3 = _module.__default.safeDivisionInt(((_1955_v26) ? (p0) : (p0)), new BigNumber(((_2008_v56).Merge((function () {
                let _coll75 = new _dafny.Map();
                for (const _compr_75 of (_1951_v23).Elements) {
                  let _2009_v57 = _compr_75;
                  if ((_1951_v23).contains(_2009_v57)) {
                    _coll75.push([_module.__default.safeModuloInt(_2009_v57, new BigNumber((_2007_v55).length)),p0]);
                  }
                }
                return _coll75;
              }()).update(p0, p0))).length));
              let _2010_v58;
              _2010_v58 = new _dafny.CodePoint('f'.codePointAt(0));
              let _2011_v59;
              let _nw318 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
              _2011_v59 = _nw318;
              let _2012_v60;
              _2012_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2011_v59,_1955_v26);
              let _2013_v61;
              _2013_v61 = _dafny.Set.fromElements(_module.__default.fm0(_1955_v26, new BigNumber(418), new BigNumber((_2012_v60).length), _1951_v23, globalState), new BigNumber((_2007_v55).length));
              let _2014_v62;
              _2014_v62 = _dafny.Set.fromElements(_module.D9.create_DC30(_2010_v58, _1955_v26, p0, _2013_v61, p0));
              let _2015_v63;
              _2015_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2010_v58,p0);
              let _2016_v64;
              _2016_v64 = _module.D0.create_DC1(_1955_v26);
              let _rhs318 = (_2014_v62).contains(_module.D9.create_DC30(_2010_v58, _1955_v26, p0, _2013_v61, p0));
              let _rhs319 = ((_module.__default.fm23(p0, _1955_v26, _2015_v63, _module.__default.fm50(p0, (_2016_v64).dtor_cf1, p0, p0, globalState), globalState)) ? (_1953_v24) : (_1953_v24));
              let _lhs272 = globalState;
              _1955_v26 = _rhs318;
              _lhs272.f2 = _rhs319;
              let _2017_v65;
              let _nw319 = Array((new BigNumber(13)).toNumber());
              _nw319[(_dafny.ZERO).toNumber()] = _1953_v24;
              _nw319[(_dafny.ONE).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(2)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(3)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(4)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(5)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(6)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(7)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(8)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(9)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(10)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(11)).toNumber()] = _1953_v24;
              _nw319[(new BigNumber(12)).toNumber()] = _1953_v24;
              _2017_v65 = _nw319;
              let _2018_v66;
              _2018_v66 = _2017_v65;
              let _2019_v67;
              let _nw320 = Array((new BigNumber(3)).toNumber());
              _nw320[(_dafny.ZERO).toNumber()] = ((_1955_v26) ? (_2017_v65) : (_2017_v65));
              _nw320[(_dafny.ONE).toNumber()] = (_2018_v66);
              _nw320[(new BigNumber(2)).toNumber()] = _2017_v65;
              _2019_v67 = _nw320;
              let _index350 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_2019_v67).length));
              (_2019_v67)[_index350] = ((_1955_v26) ? (_2017_v65) : (_2017_v65));
              let _index351 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_2019_v67).length));
              (_2019_v67)[_index351] = _2017_v65;
              let _index352 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2011_v59).length));
              (_2011_v59)[_index352] = p0;
              let _2020_v68;
              _2020_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.UnicodeFromString("itn"));
              let _index353 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2011_v59).length));
              (_2011_v59)[_index353] = _module.__default.safeModuloInt(new BigNumber(-772), new BigNumber((_2020_v68).length));
            } else {
              let _2021___mcc_h4 = (_source20).cf70;
              let _2022_cf70 = _2021___mcc_h4;
              let _2023_v69;
              let _nw321 = new _module.C0();
              _nw321.__ctor(_1955_v26, !(_1955_v26));
              _2023_v69 = _nw321;
              (globalState).f9 = (_2023_v69).f24;
              (globalState).f9 = (p0).isLessThan((new BigNumber(53)).minus(new BigNumber(175)));
              let _2024_v70;
              let _nw322 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _2024_v70 = _nw322;
              let _2025_v71;
              let _nw323 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
              _2025_v71 = _nw323;
              let _2026_v72;
              _2026_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1955_v26,(_2023_v69).f23);
              let _rhs320 = _2025_v71;
              let _rhs321 = (((((_2026_v72).contains(_1955_v26)) ? ((_2026_v72).get(_1955_v26)) : (_1955_v26))) ? (((false) ? (_2024_v70) : (_2024_v70))) : (_2024_v70));
              let _lhs273 = globalState;
              _lhs273.f11 = _rhs320;
              _2024_v70 = _rhs321;
            }
          }
        }
      }
      let _2027_v73;
      _2027_v73 = _dafny.Seq.UnicodeFromString("g");
      let _2028_v74;
      _2028_v74 = _module.D0.create_DC2(_module.__default.fm20(p0, new BigNumber((_2027_v73).length), p0, globalState));
      let _2029_v75;
      _2029_v75 = _module.D0.create_DC2(_2028_v74);
      let _pat_let_tv44 = p0;
      if (function (_source21) {
        if (_source21.is_DC1) {
          let _2030___mcc_h5 = (_source21).cf1;
          let _2031_cf1 = _2030___mcc_h5;
          return !(!(_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), function (_2032_i12) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length), _pat_let_tv44), new BigNumber(952))));
        } else if (_source21.is_DC0) {
          let _2033___mcc_h6 = (_source21).cf0;
          let _2034_cf0 = _2033___mcc_h6;
          return _2034_cf0;
        } else {
          let _2035___mcc_h7 = (_source21).cf2;
          let _2036_cf2 = _2035___mcc_h7;
          return !(_dafny.Map.Empty.slice().updateUnsafe(true,false)).equals(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)));
        }
      }(_2029_v75)) {
        let _2037_v76;
        _2037_v76 = false;
        let _2038_v77;
        _2038_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2027_v73,_2037_v76);
        _2038_v77 = _2038_v77;
        (globalState).f9 = (_2037_v76) || (_2037_v76);
        let _2039_v78;
        _2039_v78 = _dafny.MultiSet.fromElements(_2037_v76, _2037_v76);
        let _2040_v79;
        let _nw324 = new _module.C5();
        _nw324.__ctor((_dafny.ZERO).minus(p0), (new BigNumber((_2027_v73).length)).isLessThanOrEqualTo(p0), _2039_v78);
        _2040_v79 = _nw324;
        let _2041_v80;
        _2041_v80 = _dafny.Map.Empty.slice().updateUnsafe(true,_2039_v78);
        let _2042_v81;
        _2042_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_2041_v80).contains(!(_2037_v76))) ? ((_2041_v80).get(!(_2037_v76))) : (_module.__default.fm33(_2037_v76, globalState))),((_2037_v76) ? (_2037_v76) : (_2037_v76)));
        _2042_v81 = (_2042_v81).update(_2039_v78, false);
        _2037_v76 = (_module.__default.safeModuloInt((_2040_v79).f18, p0)).isLessThan(new BigNumber((_2027_v73).length));
      } else {
        (globalState).f3 = p0;
        let _2043_v82;
        _2043_v82 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2044_v83;
        _2044_v83 = _module.D2.create_DC9(true, p0, _2043_v82);
        (globalState).f9 = (_2044_v83).dtor_cf13;
        let _2045_v84;
        _2045_v84 = true;
        let _2046_v85;
        _2046_v85 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), ((_2047_v82) => function (_2048_i13) {
          return _2047_v82;
        })(_2043_v82)),_2045_v84);
        _2046_v85 = ((_2046_v85).Merge(_2046_v85)).Merge(_2046_v85);
        let _2049_v86;
        _2049_v86 = _dafny.MultiSet.fromElements(_2045_v84, _2045_v84);
        let _2050_v87;
        let _nw325 = new _module.C4();
        _nw325.__ctor(_2045_v84, !(_2045_v84), (_2049_v86).Difference(_2049_v86));
        _2050_v87 = _nw325;
        (globalState).f9 = (_2050_v87).f19;
      }
      let _2051_v88;
      _2051_v88 = true;
      if (_2051_v88) {
        let _2052_v89;
        _2052_v89 = _dafny.MultiSet.fromElements(false);
        let _2053_v90;
        _2053_v90 = _dafny.Map.Empty.slice().updateUnsafe(_2051_v88,_2051_v88);
        let _2054_v91;
        let _init55 = ((_2055_v88) => function (_2056_i14) {
          return _2055_v88;
        })(_2051_v88);
        let _nw326 = Array((new BigNumber(15)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw326.length); _i0_55++) {
          _nw326[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2054_v91 = _nw326;
        let _2057_v92;
        _2057_v92 = _module.D11.create_DC34(_2052_v89, _dafny.MultiSet.fromElements(new BigNumber((_2053_v90).length)), _2054_v91, _module.D1.create_DC5(_2051_v88, false));
        let _2058_v93;
        _2058_v93 = _dafny.MultiSet.fromElements((_2057_v92).dtor_cf60);
        let _2059_v94;
        _2059_v94 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2060_v95;
        _2060_v95 = _dafny.Map.Empty.slice().updateUnsafe((p0).minus(new BigNumber((_2058_v93).cardinality())),_2059_v94);
        let _2061_v96;
        _2061_v96 = _dafny.Seq.of(p0, new BigNumber((_2052_v89).cardinality()));
        _2060_v95 = (_2060_v95).update(new BigNumber((_2061_v96).length), _2059_v94);
        let _2062_v97;
        let _nw327 = new _module.C0();
        _nw327.__ctor(_2051_v88, _2051_v88);
        _2062_v97 = _nw327;
        let _2063_v98;
        _2063_v98 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2062_v97);
        _2063_v98 = _2063_v98;
        let _2064_v99;
        _2064_v99 = _module.D14.create_DC45(_2053_v90, _2062_v97, p0);
        let _2065_v100;
        _2065_v100 = _module.D14.create_DC46(_2064_v99);
        let _pat_let_tv45 = _2064_v99;
        let _pat_let_tv46 = _2064_v99;
        let _2066_v101;
        let _nw328 = Array((new BigNumber(26)).toNumber());
        _nw328[(_dafny.ZERO).toNumber()] = _2065_v100;
        _nw328[(_dafny.ONE).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(2)).toNumber()] = function (_pat_let64_0) {
          return function (_2067_dt__update__tmp_h1) {
            return function (_pat_let65_0) {
              return function (_2068_dt__update_hcf75_h0) {
                return _module.D14.create_DC46(_2068_dt__update_hcf75_h0);
              }(_pat_let65_0);
            }(_pat_let_tv45);
          }(_pat_let64_0);
        }(_2065_v100);
        _nw328[(new BigNumber(3)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(4)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(5)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(6)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(7)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(8)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(9)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(10)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(11)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(12)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(13)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(14)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(15)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(16)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(17)).toNumber()] = _module.D14.create_DC46(_2064_v99);
        _nw328[(new BigNumber(18)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(19)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(20)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(21)).toNumber()] = ((_2051_v88) ? (_2065_v100) : (_2065_v100));
        _nw328[(new BigNumber(22)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(23)).toNumber()] = _2065_v100;
        _nw328[(new BigNumber(24)).toNumber()] = function (_pat_let66_0) {
          return function (_2069_dt__update__tmp_h2) {
            return function (_pat_let67_0) {
              return function (_2070_dt__update_hcf75_h1) {
                return _module.D14.create_DC46(_2070_dt__update_hcf75_h1);
              }(_pat_let67_0);
            }(_pat_let_tv46);
          }(_pat_let66_0);
        }(_2065_v100);
        _nw328[(new BigNumber(25)).toNumber()] = _2065_v100;
        _2066_v101 = _nw328;
        let _index354 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2066_v101).length));
        (_2066_v101)[_index354] = _2065_v100;
        let _index355 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2066_v101).length));
        let _rhs322 = (_2062_v97).f23;
        let _rhs323 = (((_2062_v97).f23) ? (_2065_v100) : (_2065_v100));
        let _rhs324 = _module.__default.fm27(p0, _2051_v88, globalState);
        let _rhs325 = p0;
        let _lhs274 = globalState;
        let _lhs275 = _2066_v101;
        let _lhs276 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2066_v101).length));
        let _lhs277 = globalState;
        _lhs274.f9 = _rhs322;
        _lhs275[_lhs276] = _rhs323;
        _2027_v73 = _rhs324;
        _lhs277.f3 = _rhs325;
        let _2071_v102;
        _2071_v102 = _dafny.Seq.of((_2062_v97).f23);
        let _2072_v103;
        _2072_v103 = _module.D2.create_DC7(_2071_v102);
        if (_dafny.areEqual(_dafny.Seq.of(true), _dafny.Seq.update((_2072_v103).dtor_cf9, _module.__default.safeIndex(p0, new BigNumber(((_2072_v103).dtor_cf9).length)), (_2062_v97).f23))) {
          let _2073_v104;
          let _nw329 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _2073_v104 = _nw329;
          let _2074_v106;
          let _init56 = ((_2075_v89, _2076_v102) => function (_2077_i15) {
            return _module.D11.create_DC33(function () {
  let _coll76 = new _dafny.Set();
  for (const _compr_76 of (_dafny.Seq.of(_2075_v89, _2075_v89, _dafny.MultiSet.FromArray(_2076_v102))).Elements) {
    let _2078_v105 = _compr_76;
    if (_dafny.Seq.contains(_dafny.Seq.of(_2075_v89, _2075_v89, _dafny.MultiSet.FromArray(_2076_v102)), _2078_v105)) {
      _coll76.add(_2078_v105);
    }
  }
  return _coll76;
}());
          })(_2052_v89, _2071_v102);
          let _nw330 = Array((new BigNumber(24)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw330.length); _i0_56++) {
            _nw330[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2074_v106 = _nw330;
          let _2079_v107;
          _2079_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2059_v94,(_dafny.ZERO).minus(p0));
          let _2080_v108;
          _2080_v108 = _dafny.Set.fromElements(_2061_v96);
          let _2081_v109;
          _2081_v109 = _dafny.Map.Empty.slice().updateUnsafe(_2074_v106,_module.__default.fm23(p0, (_2062_v97).f23, _2079_v107, _2080_v108, globalState));
          let _index356 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2073_v104).length));
          (_2073_v104)[_index356] = _2081_v109;
          let _index357 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_2073_v104).length));
          (_2073_v104)[_index357] = (_2081_v109).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2074_v106,(_2062_v97).f23));
          (globalState).f9 = !((p0).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("yacasvvoi")).length))) || (_2051_v88);
          let _2082_v110;
          _2082_v110 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p0, p0),p0);
          _2082_v110 = (_2082_v110).Merge(_2082_v110);
          let _2083_v111;
          _2083_v111 = _module.D13.create_DC42(_2059_v94, (_2062_v97).f24);
          let _2084_v112;
          _2084_v112 = _module.D13.create_DC43(_2083_v111);
          let _2085_v113;
          _2085_v113 = _module.D13.create_DC43(_module.D13.create_DC43(_2083_v111));
          let _2086_v114;
          _2086_v114 = _module.D13.create_DC43(_2085_v113);
          let _2087_v115;
          _2087_v115 = _module.D13.create_DC43(_2084_v112);
          let _2088_v116;
          _2088_v116 = _dafny.Set.fromElements(p0);
          let _2089_v118;
          _2089_v118 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(p0));
          let _2090_v119;
          _2090_v119 = _dafny.Seq.of(_2027_v73);
          let _2091_v120;
          _2091_v120 = _module.D14.create_DC45(_2053_v90, _2062_v97, new BigNumber(((_2090_v119)[_module.__default.safeIndex(p0, new BigNumber((_2090_v119).length))]).length));
          let _2092_v124;
          _2092_v124 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2062_v97).f24);
          let _2093_v126;
          _2093_v126 = _dafny.Seq.of(_2092_v124, (function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of _dafny.IntegerRange(new BigNumber(206), new BigNumber(40))) {
              let _2094_v125 = _compr_77;
              if (((new BigNumber(206)).isLessThanOrEqualTo(_2094_v125)) && ((_2094_v125).isLessThan(new BigNumber(40)))) {
                _coll77.push([_module.__default.safeDivisionInt(_2094_v125, new BigNumber(-866)),(_2062_v97).f23]);
              }
            }
            return _coll77;
          }()).update(p0, true));
          let _2095_v128;
          let _nw331 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2095_v128 = _nw331;
          let _2096_v130;
          let _nw332 = Array((new BigNumber(19)).toNumber());
          _nw332[(_dafny.ZERO).toNumber()] = _2088_v116;
          _nw332[(_dafny.ONE).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(2)).toNumber()] = function () {
            let _coll78 = new _dafny.Set();
            for (const _compr_78 of _dafny.IntegerRange(new BigNumber(778), new BigNumber(596))) {
              let _2097_v117 = _compr_78;
              if (((new BigNumber(778)).isLessThanOrEqualTo(_2097_v117)) && ((_2097_v117).isLessThan(new BigNumber(596)))) {
                _coll78.add((_2097_v117).multipliedBy(p0));
              }
            }
            return _coll78;
          }();
          _nw332[(new BigNumber(3)).toNumber()] = (((_2062_v97).f23) ? ((((_2089_v118).contains(new BigNumber(((_2091_v120).dtor_cf72).length))) ? ((_2089_v118).get(new BigNumber(((_2091_v120).dtor_cf72).length))) : (_2088_v116))) : (_2088_v116));
          _nw332[(new BigNumber(4)).toNumber()] = function () {
            let _coll79 = new _dafny.Set();
            for (const _compr_79 of _dafny.IntegerRange(new BigNumber(407), new BigNumber(-440))) {
              let _2098_v121 = _compr_79;
              if (((new BigNumber(407)).isLessThanOrEqualTo(_2098_v121)) && ((_2098_v121).isLessThan(new BigNumber(-440)))) {
                _coll79.add(_module.__default.safeModuloInt(_2098_v121, p0));
              }
            }
            return _coll79;
          }();
          _nw332[(new BigNumber(5)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(6)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(7)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(8)).toNumber()] = function () {
            let _coll80 = new _dafny.Set();
            for (const _compr_80 of _dafny.IntegerRange(new BigNumber(-243), new BigNumber(713))) {
              let _2099_v122 = _compr_80;
              if (((new BigNumber(-243)).isLessThanOrEqualTo(_2099_v122)) && ((_2099_v122).isLessThan(new BigNumber(713)))) {
                _coll80.add((_2099_v122).multipliedBy(p0));
              }
            }
            return _coll80;
          }();
          _nw332[(new BigNumber(9)).toNumber()] = (function () {
            let _coll81 = new _dafny.Set();
            for (const _compr_81 of _dafny.IntegerRange(new BigNumber(437), new BigNumber(-919))) {
              let _2100_v123 = _compr_81;
              if (((new BigNumber(437)).isLessThanOrEqualTo(_2100_v123)) && ((_2100_v123).isLessThan(new BigNumber(-919)))) {
                _coll81.add((_2100_v123).minus(new BigNumber((_dafny.Set.fromElements((_2062_v97).f23)).length)));
              }
            }
            return _coll81;
          }()).Difference(_2088_v116);
          _nw332[(new BigNumber(10)).toNumber()] = _module.__default.fm22(new BigNumber((_2027_v73).length), _2093_v126, (_2062_v97).f23, (_2062_v97).f24, globalState);
          _nw332[(new BigNumber(11)).toNumber()] = function () {
            let _coll82 = new _dafny.Set();
            for (const _compr_82 of _dafny.IntegerRange(new BigNumber(-200), new BigNumber(104))) {
              let _2101_v127 = _compr_82;
              if (((new BigNumber(-200)).isLessThanOrEqualTo(_2101_v127)) && ((_2101_v127).isLessThan(new BigNumber(104)))) {
                _coll82.add((_2101_v127).plus(new BigNumber((_1951_v23).cardinality())));
              }
            }
            return _coll82;
          }();
          _nw332[(new BigNumber(12)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(13)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(14)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(p0, new BigNumber((_2093_v126).length), _module.__default.fm0(_2051_v88, new BigNumber((_dafny.MultiSet.fromElements(_2095_v128)).cardinality()), new BigNumber((_dafny.Seq.of((_2062_v97).f23, (_2062_v97).f24)).length), _1951_v23, globalState), p0, p0);
          _nw332[(new BigNumber(16)).toNumber()] = function () {
            let _coll83 = new _dafny.Set();
            for (const _compr_83 of _dafny.IntegerRange(new BigNumber(311), new BigNumber(608))) {
              let _2102_v129 = _compr_83;
              if (((new BigNumber(311)).isLessThanOrEqualTo(_2102_v129)) && ((_2102_v129).isLessThan(new BigNumber(608)))) {
                _coll83.add((_2102_v129).multipliedBy(new BigNumber((_dafny.Seq.of((_2062_v97).f24, (_2062_v97).f24)).length)));
              }
            }
            return _coll83;
          }();
          _nw332[(new BigNumber(17)).toNumber()] = _2088_v116;
          _nw332[(new BigNumber(18)).toNumber()] = _2088_v116;
          _2096_v130 = _nw332;
          let _pat_let_tv47 = _2086_v114;
          let _rhs326 = _2096_v130;
          let _rhs327 = (_2062_v97).f23;
          let _rhs328 = p0;
          let _rhs329 = function (_pat_let68_0) {
            return function (_2103_dt__update__tmp_h3) {
              return function (_pat_let69_0) {
                return function (_2104_dt__update_hcf70_h0) {
                  return _module.D13.create_DC43(_2104_dt__update_hcf70_h0);
                }(_pat_let69_0);
              }(_pat_let_tv47);
            }(_pat_let68_0);
          }(_module.__default.fm51(_2051_v88, (_2062_v97).f23, _2051_v88, new BigNumber((_2092_v124).length), globalState));
          let _lhs278 = globalState;
          let _lhs279 = globalState;
          let _lhs280 = globalState;
          _lhs278.f11 = _rhs326;
          _lhs279.f9 = _rhs327;
          _lhs280.f3 = _rhs328;
          _2087_v115 = _rhs329;
          _2027_v73 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-323)), ((_2105_v73, _2106_v97) => function (_2107_i16) {
            return (_2105_v73)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_2106_v97).f24)).length), new BigNumber((_2105_v73).length))];
          })(_2027_v73, _2062_v97));
        } else {
          let _2108_v131;
          let _init57 = ((_2109_p0) => function (_2110_i17) {
            return (_2110_i17).minus(_2109_p0);
          })(p0);
          let _nw333 = Array((new BigNumber(22)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw333.length); _i0_57++) {
            _nw333[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2108_v131 = _nw333;
          let _index358 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_2108_v131).length));
          (_2108_v131)[_index358] = p0;
          let _index359 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_2108_v131).length));
          (_2108_v131)[_index359] = p0;
          let _2111_v132;
          let _nw334 = new _module.C3();
          _nw334.__ctor(_module.__default.safeModuloInt(new BigNumber(845), p0), _dafny.Seq.Concat(_2027_v73, _2027_v73), false, (_2052_v89).Union(_2052_v89));
          _2111_v132 = _nw334;
          let _2112_v134;
          _2112_v134 = _module.D9.create_DC30(_2059_v94, (_2062_v97).f23, new BigNumber((_2071_v102).length), function () {
  let _coll84 = new _dafny.Set();
  for (const _compr_84 of (_2061_v96).Elements) {
    let _2113_v133 = _compr_84;
    if (_dafny.Seq.contains(_2061_v96, _2113_v133)) {
      _coll84.add(_module.__default.safeModuloInt(_2113_v133, new BigNumber(979)));
    }
  }
  return _coll84;
}(), (_2108_v131)[_module.__default.safeIndex(new BigNumber(927), new BigNumber((_2108_v131).length))]);
          (globalState).f9 = (_2112_v134).dtor_cf51;
          let _index360 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_2108_v131).length));
          (_2108_v131)[_index360] = (_2108_v131)[_module.__default.safeIndex(new BigNumber(927), new BigNumber((_2108_v131).length))];
          let _index361 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_2054_v91).length));
          (_2054_v91)[_index361] = ((_2062_v97).f23) && ((_2062_v97).f24);
          let _index362 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_2054_v91).length));
          (_2054_v91)[_index362] = (_2062_v97).f24;
        }
        (globalState).f9 = (_2062_v97).f24;
      } else {
        (globalState).f9 = _2051_v88;
        let _2114_v135;
        _2114_v135 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),!(!(_2051_v88)));
        let _2115_v136;
        _2115_v136 = _dafny.Seq.of(_2114_v135, (_2114_v135).Merge(_module.__default.fm46(new BigNumber((_2027_v73).length), _2051_v88, globalState)), _2114_v135, _2114_v135, _2114_v135);
        (globalState).f3 = new BigNumber((_2115_v136).length);
        let _2116_v137;
        let _nw335 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2116_v137 = _nw335;
        let _index363 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2116_v137).length));
        (_2116_v137)[_index363] = _2027_v73;
        let _2117_v138;
        _2117_v138 = new _dafny.CodePoint('e'.codePointAt(0));
        let _2118_v139;
        let _init58 = ((_2119_v88) => function (_2120_i18) {
          return _2119_v88;
        })(_2051_v88);
        let _nw336 = Array((new BigNumber(14)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw336.length); _i0_58++) {
          _nw336[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2118_v139 = _nw336;
        let _2121_v140;
        _2121_v140 = _dafny.Seq.of(_2118_v139);
        let _2122_v141;
        _2122_v141 = _dafny.MultiSet.fromElements((_2121_v140)[_module.__default.safeIndex(p0, new BigNumber((_2121_v140).length))]);
        let _index364 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2116_v137).length));
        (_2116_v137)[_index364] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm27(p0, _2051_v88, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm27(p0, _2051_v88, globalState)).length)), _2117_v138), _module.__default.safeIndex(new BigNumber(-50), new BigNumber((_dafny.Seq.update(_module.__default.fm27(p0, _2051_v88, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm27(p0, _2051_v88, globalState)).length)), _2117_v138)).length)), _2117_v138), _module.__default.safeIndex(new BigNumber((_2122_v141).cardinality()), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm27(p0, _2051_v88, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm27(p0, _2051_v88, globalState)).length)), _2117_v138), _module.__default.safeIndex(new BigNumber(-50), new BigNumber((_dafny.Seq.update(_module.__default.fm27(p0, _2051_v88, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm27(p0, _2051_v88, globalState)).length)), _2117_v138)).length)), _2117_v138)).length)), _2117_v138);
        let _2123_v142;
        let _nw337 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _2123_v142 = _nw337;
        let _index365 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2123_v142).length));
        (_2123_v142)[_index365] = p0;
        let _index366 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2123_v142).length));
        (_2123_v142)[_index366] = ((new BigNumber(((_2116_v137)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2116_v137).length))]).length)).plus(p0)).minus(p0);
        _2117_v138 = _2117_v138;
      }
      return;
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f9 = p0;
      let _2124_v0;
      let _init59 = function (_2125_i1) {
        return _module.__default.safeDivisionInt(_2125_i1, new BigNumber(150));
      };
      let _nw338 = Array((new BigNumber(27)).toNumber());
      for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw338.length); _i0_59++) {
        _nw338[_i0_59] = _init59(new BigNumber(_i0_59));
      }
      _2124_v0 = _nw338;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2124_v0).length))) {
        let _2126_i0 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2126_i0)) && ((_2126_i0).isLessThan(new BigNumber((_2124_v0).length))))) {
          (_2124_v0)[(_2126_i0)] = (_2126_i0).plus((_module.D3.create_DC11(new BigNumber(463), p2, p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p2)).length))).dtor_cf18);
        }
      }
      let _2127_v1;
      _2127_v1 = _dafny.Set.fromElements(p0, p0);
      let _2128_v2;
      _2128_v2 = _dafny.Map.Empty.slice().updateUnsafe((_2127_v1).Union(_module.__default.fm41(p2, p1, globalState)),p0);
      (globalState).f3 = new BigNumber((_2128_v2).length);
      let _2129_v3;
      let _nw339 = Array((new BigNumber(2)).toNumber());
      _nw339[(_dafny.ZERO).toNumber()] = false;
      _nw339[(_dafny.ONE).toNumber()] = p0;
      _2129_v3 = _nw339;
      let _index367 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length));
      (_2129_v3)[_index367] = p0;
      let _index368 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length));
      (_2129_v3)[_index368] = p0;
      let _2130_v4;
      _2130_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2129_v3)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length))]);
      let _index369 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      (_2124_v0)[_index369] = (_dafny.ZERO).minus(p1);
      let _index370 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length));
      let _index371 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      let _rhs330 = _2130_v4;
      let _rhs331 = !(p1).isEqualTo((_dafny.ZERO).minus((p1).minus(p2)));
      let _rhs332 = p1;
      let _rhs333 = p2;
      let _lhs281 = _2129_v3;
      let _lhs282 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length));
      let _lhs283 = _2124_v0;
      let _lhs284 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      let _lhs285 = globalState;
      _2130_v4 = _rhs330;
      _lhs281[_lhs282] = _rhs331;
      _lhs283[_lhs284] = _rhs332;
      _lhs285.f3 = _rhs333;
      let _2131_v5;
      _2131_v5 = new _dafny.CodePoint('c'.codePointAt(0));
      let _2132_v6;
      _2132_v6 = _dafny.MultiSet.fromElements((_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], new BigNumber(441));
      let _2133_v7;
      _2133_v7 = _dafny.MultiSet.fromElements(false);
      let _2134_v8;
      let _nw340 = new _module.C5();
      _nw340.__ctor((_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], p0, _2133_v7);
      _2134_v8 = _nw340;
      let _2135_v9;
      _2135_v9 = _dafny.MultiSet.fromElements(_2134_v8, _2134_v8);
      let _2136_v10;
      _2136_v10 = _dafny.Seq.UnicodeFromString("nnbtloi");
      let _2137_v11;
      _2137_v11 = _dafny.Set.fromElements((_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], _module.__default.fm0((_2129_v3)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length))], (_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], (_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], (_2132_v6).update(new BigNumber((_2135_v9).cardinality()), _module.__default.abs(p1)), globalState), new BigNumber((_2136_v10).length), p1);
      let _index372 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      let _index373 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      let _rhs334 = (_2129_v3)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length))];
      let _rhs335 = _module.__default.safeModuloInt((_2124_v0)[_module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length))], new BigNumber((_2137_v11).length));
      let _rhs336 = _2131_v5;
      let _rhs337 = (_dafny.ZERO).minus(p2);
      let _lhs286 = globalState;
      let _lhs287 = _2124_v0;
      let _lhs288 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      let _lhs289 = _2124_v0;
      let _lhs290 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2124_v0).length));
      _lhs286.f9 = _rhs334;
      _lhs287[_lhs288] = _rhs335;
      _2131_v5 = _rhs336;
      _lhs289[_lhs290] = _rhs337;
      let _2138_v12;
      _2138_v12 = _dafny.Seq.of(p0, true, true);
      r0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_2129_v3)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_2129_v3).length))]), _2138_v12);
      return r0;
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f14 = false;
      this._f15 = _dafny.MultiSet.Empty;
      this._f16 = _dafny.Map.Empty;
      this._f17 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f16, f17, f14, f15) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return _this.f14;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((_this.f14) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-738)), function (_2139_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), function (_2140_i1) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }))), _dafny.Seq.UnicodeFromString("jshowx"));
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return _this.f14;
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _2141_v0;
      let _nw341 = Array((new BigNumber(8)).toNumber());
      _nw341[(_dafny.ZERO).toNumber()] = !(_this.f14);
      _nw341[(_dafny.ONE).toNumber()] = _this.f14;
      _nw341[(new BigNumber(2)).toNumber()] = _this.f14;
      _nw341[(new BigNumber(3)).toNumber()] = _this.f14;
      _nw341[(new BigNumber(4)).toNumber()] = _this.f14;
      _nw341[(new BigNumber(5)).toNumber()] = _this.f14;
      _nw341[(new BigNumber(6)).toNumber()] = _this.f14;
      _nw341[(new BigNumber(7)).toNumber()] = _this.f14;
      _2141_v0 = _nw341;
      let _2142_v1;
      _2142_v1 = _dafny.MultiSet.fromElements(_2141_v0, _2141_v0);
      if (!(_2142_v1).contains(_2141_v0)) {
        let _2143_v2;
        _2143_v2 = _dafny.Set.fromElements(!(_this.f14), _this.f14, _this.f14, _this.f14);
        (globalState).f9 = !(_2143_v2).equals(_2143_v2);
        if (((_this.f14) ? (true) : (_this.f14))) {
          r3 = !(_this.f14);
          let _2144_v3;
          _2144_v3 = _dafny.Seq.UnicodeFromString("okwamfa");
          r3 = (_this).fm2(_this.f14, _dafny.areEqual(_2144_v3, _2144_v3), globalState);
          (globalState).f13 = _dafny.Set.fromElements(_this.f14);
          let _2145_v4;
          _2145_v4 = new BigNumber(431);
          (globalState).f3 = (_dafny.ZERO).minus(_2145_v4);
          let _2146_v5;
          _2146_v5 = _dafny.Seq.of(_2145_v4);
          let _2147_v6;
          _2147_v6 = _dafny.Seq.of(_2146_v5, _2146_v5, _2146_v5);
          let _index374 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_2141_v0).length));
          (_2141_v0)[_index374] = _dafny.areEqual(_dafny.Seq.UnicodeFromString("ahnxhipsh"), _2144_v3);
          let _index375 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_2141_v0).length));
          let _rhs338 = _dafny.Seq.Concat(((_this.f14) ? (_2147_v6) : (_2147_v6)), _2147_v6);
          let _rhs339 = ((((_this).f15).equals((_this).f15)) ? (_this.f14) : (_this.f14));
          let _lhs291 = _2141_v0;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_2141_v0).length));
          _2147_v6 = _rhs338;
          _lhs291[_lhs292] = _rhs339;
        } else {
          let _2148_v7;
          let _nw342 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _2148_v7 = _nw342;
          let _2149_v8;
          _2149_v8 = new BigNumber(726);
          let _index376 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length));
          (_2148_v7)[_index376] = _module.__default.fm0(_this.f14, new BigNumber(325), _2149_v8, _dafny.MultiSet.fromElements(_2149_v8, new BigNumber(370), _2149_v8, _2149_v8, _2149_v8), globalState);
          let _2150_v9;
          _2150_v9 = _dafny.Seq.of(new BigNumber(-348));
          let _2151_v10;
          _2151_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v8,_2149_v8);
          let _2152_v11;
          _2152_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2150_v9,_2151_v10);
          let _2153_v12;
          _2153_v12 = _dafny.MultiSet.fromElements(_2149_v8, new BigNumber((_module.__default.fm5(_this.f14, globalState)).length), new BigNumber((_2152_v11).length));
          let _index377 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length));
          (_2148_v7)[_index377] = _module.__default.fm0(_this.f14, _2149_v8, _2149_v8, _2153_v12, globalState);
          let _2154_v13;
          _2154_v13 = _module.D0.create_DC1(_this.f14);
          let _2155_v14;
          _2155_v14 = _dafny.Seq.of((_2154_v13).dtor_cf1);
          let _2156_v15;
          _2156_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2150_v9,new BigNumber((_2155_v14).length));
          let _2157_v16;
          _2157_v16 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f14) ? ((_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))]) : ((((_2156_v15).contains(_2150_v9)) ? ((_2156_v15).get(_2150_v9)) : ((_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))])))),(_this.f14) === (_this.f14));
          let _2158_v17;
          _2158_v17 = _module.D0.create_DC0(_this.f14);
          _2157_v16 = (_2157_v16).update(_module.__default.safeModuloInt(_2149_v8, _2149_v8), (_2158_v17).dtor_cf0);
          let _2159_v18;
          _2159_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2148_v7);
          let _2160_v19;
          _2160_v19 = _dafny.Set.fromElements(_2149_v8, new BigNumber((_2159_v18).length), (_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))], _2149_v8);
          let _2161_v20;
          _2161_v20 = _dafny.Seq.UnicodeFromString("h");
          let _2162_v21;
          let _out59;
          _out59 = (_this).m2(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_2160_v19).length))).length), _2149_v8), _dafny.Seq.Concat(_2161_v20, _dafny.Seq.UnicodeFromString("fmlv")), _this.f14, !(_this.f14) || (_this.f14), globalState);
          _2162_v21 = _out59;
          let _2163_v23;
          _2163_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mkl"),true);
          let _2164_v24;
          _2164_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(false),(new BigNumber((_2161_v20).length)).isLessThan(new BigNumber((function () {
            let _coll85 = new _dafny.Map();
            for (const _compr_85 of (_2163_v23).Keys.Elements) {
              let _2165_v22 = _compr_85;
              if ((_2163_v23).contains(_2165_v22)) {
                _coll85.push([_2165_v22,_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), ((_2166_v21) => function (_2167_i0) {
                  return _2166_v21;
                })(_2162_v21))]);
              }
            }
            return _coll85;
          }()).length)));
          (_this).f14 = (((_2164_v24).contains(_2158_v17)) ? ((_2164_v24).get(_2158_v17)) : ((new BigNumber(264)).isLessThan(new BigNumber((_2161_v20).length))));
          let _2168_v25;
          _2168_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber(562));
          let _2169_v26;
          _2169_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))],_2168_v25);
          _2168_v25 = (((_2169_v26).contains((_dafny.ZERO).minus((new BigNumber(945)).minus((_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))])))) ? ((_2169_v26).get((_dafny.ZERO).minus((new BigNumber(945)).minus((_2148_v7)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2148_v7).length))])))) : ((_2168_v25).update(_this.f14, new BigNumber((_dafny.Set.fromElements(_this.f14, _this.f14)).length))));
        }
        let _2170_v27;
        _2170_v27 = new BigNumber(-905);
        let _2171_v28;
        _2171_v28 = _dafny.Seq.of(((true) ? (_2170_v27) : (new BigNumber(3))), _2170_v27);
        _2171_v28 = _2171_v28;
        r0 = _2170_v27;
        r3 = !(((_this.f14) ? (_this.f14) : (_this.f14)));
      } else {
        let _2172_v29;
        _2172_v29 = _dafny.Seq.UnicodeFromString("vo");
        let _2173_v30;
        let _init60 = function (_2174_i1) {
          return _module.D0.create_DC0(_this.f14);
        };
        let _nw343 = Array((new BigNumber(17)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw343.length); _i0_60++) {
          _nw343[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2173_v30 = _nw343;
        let _2175_v31;
        let _nw344 = Array((new BigNumber(28)).toNumber());
        _nw344[(_dafny.ZERO).toNumber()] = _2173_v30;
        _nw344[(_dafny.ONE).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(2)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(3)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(4)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(5)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(6)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(7)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(8)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(9)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(10)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(11)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(12)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(13)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(14)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(15)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(16)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(17)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(18)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(19)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(20)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(21)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(22)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(23)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(24)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(25)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(26)).toNumber()] = _2173_v30;
        _nw344[(new BigNumber(27)).toNumber()] = _2173_v30;
        _2175_v31 = _nw344;
        let _2176_v32;
        _2176_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2172_v29,_2175_v31);
        _2176_v32 = (_2176_v32).update(_2172_v29, _2175_v31);
        let _2177_v33;
        _2177_v33 = new BigNumber(411);
        let _2178_v34;
        _2178_v34 = _dafny.Seq.of(_2177_v33);
        let _2179_v35;
        _2179_v35 = _dafny.Set.fromElements(new BigNumber((_2178_v34).length), _2177_v33, _2177_v33);
        (globalState).f3 = new BigNumber(((_module.__default.fm6(globalState)).Difference(_2179_v35)).length);
        r0 = new BigNumber(-607);
        let _2180_v36;
        let _nw345 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2180_v36 = _nw345;
        let _2181_v37;
        _2181_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
        let _2182_v38;
        _2182_v38 = _dafny.Map.Empty.slice().updateUnsafe(((!(_this.f14)) ? (_module.__default.fm7(globalState)) : (_2181_v37)),_2177_v33);
        let _rhs340 = _2180_v36;
        let _rhs341 = function () {
          let _coll86 = new _dafny.Map();
          for (const _compr_86 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(646)), ((_2183_v37) => function (_2184_i2) {
            return (_2183_v37).update(false, _this.f14);
          })(_2181_v37))).Elements) {
            let _2185_v39 = _compr_86;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(646)), ((_2186_v37) => function (_2184_i2) {
              return (_2186_v37).update(false, _this.f14);
            })(_2181_v37)), _2185_v39)) {
              _coll86.push([_2185_v39,_2177_v33]);
            }
          }
          return _coll86;
        }();
        let _rhs342 = _dafny.Seq.Concat(_2172_v29, _2172_v29);
        _2180_v36 = _rhs340;
        _2182_v38 = _rhs341;
        _2172_v29 = _rhs342;
        let _2187_v40;
        let _nw346 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _2187_v40 = _nw346;
        let _2188_v41;
        _2188_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2177_v33,_2187_v40);
        _2188_v41 = (_2188_v41).update(_2177_v33, _2187_v40);
      }
      let _2189_v42;
      _2189_v42 = _dafny.Set.fromElements(_this.f14);
      let _2190_v43;
      _2190_v43 = new BigNumber(703);
      let _2191_v44;
      _2191_v44 = new _dafny.CodePoint('h'.codePointAt(0));
      let _2192_v45;
      _2192_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2190_v43,_2191_v44);
      let _2193_v46;
      _2193_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_this.f14)).IsDisjointFrom(_2189_v42),_2192_v45);
      _2193_v46 = (_2193_v46).update(_this.f14, _dafny.Map.Empty.slice().updateUnsafe(_2190_v43,_2191_v44));
      let _2194_v47;
      _2194_v47 = _dafny.Seq.of(_this.f14, _this.f14);
      let _2195_v48;
      _2195_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2194_v47,!(_this.f14));
      _2195_v48 = (_2195_v48).update(_2194_v47, (_this).fm2(_this.f14, _this.f14, globalState));
      let _2196_v49;
      _2196_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,!(_this.f14));
      let _2197_i3;
      _2197_i3 = _dafny.ZERO;
      L5: {
        while ((_this).fm4(_2196_v49, _this.f14, globalState)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2197_i3)) {
              break L5;
            }
            _2197_i3 = (_2197_i3).plus(_dafny.ONE);
            let _2198_v50;
            _2198_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2190_v43);
            let _2199_v51;
            _2199_v51 = _dafny.MultiSet.fromElements(new BigNumber(501));
            (_this).f14 = (_2199_v51).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_2198_v50).length)));
            let _2200_v52;
            _2200_v52 = _dafny.Map.Empty.slice().updateUnsafe((_2199_v51).Difference(_2199_v51),_this.f14);
            _2200_v52 = _2200_v52;
            if ((_module.D0.create_DC0((_2194_v47)[_module.__default.safeIndex(_2190_v43, new BigNumber((_2194_v47).length))])).dtor_cf0) {
              let _2201_v53;
              _2201_v53 = _module.D1.create_DC4(_2190_v43, _2190_v43);
              (globalState).f9 = ((_dafny.ZERO).minus(_2190_v43)).isEqualTo((_2201_v53).dtor_cf4);
              let _2202_v54;
              _2202_v54 = _dafny.Seq.UnicodeFromString("wfaglrkli");
              let _2203_v55;
              _2203_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_dafny.Seq.of(_2190_v43, new BigNumber((_2202_v54).length)));
              let _2204_v56;
              _2204_v56 = _dafny.Seq.of(_2190_v43, _2190_v43);
              let _2205_v57;
              _2205_v57 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_dafny.Seq.update((((_2203_v55).contains(_this.f14)) ? ((_2203_v55).get(_this.f14)) : (_2204_v56)), _module.__default.safeIndex(_2190_v43, new BigNumber(((((_2203_v55).contains(_this.f14)) ? ((_2203_v55).get(_this.f14)) : (_2204_v56))).length)), _2190_v43));
              _2203_v55 = ((_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_dafny.Seq.of(_2190_v43))).Merge(_module.__default.fm8(_this.f14, _this.f14, new BigNumber((_module.__default.fm6(globalState)).length), _this.f14, globalState))).Merge(((_this.f14) ? (_2205_v57) : (_2203_v55)));
              let _2206_v58;
              _2206_v58 = _dafny.Map.Empty.slice().updateUnsafe((((_2196_v49).contains(_this.f14)) ? ((_2196_v49).get(_this.f14)) : (_this.f14)),_2201_v53);
              _2206_v58 = (_2206_v58).update((_2194_v47)[_module.__default.safeIndex(_2190_v43, new BigNumber((_2194_v47).length))], _module.D1.create_DC4(_2190_v43, _2190_v43));
              (globalState).f9 = _this.f14;
              (globalState).f9 = false;
            } else {
              let _2207_v59;
              let _nw347 = Array((new BigNumber(15)).toNumber()).fill(_module.D1.Default());
              _2207_v59 = _nw347;
              let _2208_v60;
              let _nw348 = Array((new BigNumber(28)).toNumber());
              _nw348[(_dafny.ZERO).toNumber()] = _2207_v59;
              _nw348[(_dafny.ONE).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(2)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(3)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(4)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(5)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(6)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(7)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(8)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(9)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(10)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(11)).toNumber()] = ((_this.f14) ? (_2207_v59) : (_2207_v59));
              _nw348[(new BigNumber(12)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(13)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(14)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(15)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(16)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(17)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(18)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(19)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(20)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(21)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(22)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(23)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(24)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(25)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(26)).toNumber()] = _2207_v59;
              _nw348[(new BigNumber(27)).toNumber()] = _2207_v59;
              _2208_v60 = _nw348;
              let _index378 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2208_v60).length));
              (_2208_v60)[_index378] = _2207_v59;
              let _2209_v61;
              _2209_v61 = _module.D1.create_DC4(_2190_v43, _2190_v43);
              let _index379 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2208_v60).length));
              let _rhs343 = (((false) && (_this.f14)) ? (_2207_v59) : (_2207_v59));
              let _rhs344 = _2209_v61;
              let _rhs345 = _2190_v43;
              let _lhs293 = _2208_v60;
              let _lhs294 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_2208_v60).length));
              _lhs293[_lhs294] = _rhs343;
              _2209_v61 = _rhs344;
              _2190_v43 = _rhs345;
              let _2210_v62;
              _2210_v62 = _dafny.Seq.of(_2141_v0, _2141_v0, _2141_v0);
              _2210_v62 = _dafny.Seq.update(_dafny.Seq.update(_2210_v62, _module.__default.safeIndex(new BigNumber(-990), new BigNumber((_2210_v62).length)), _2141_v0), _module.__default.safeIndex(_2190_v43, new BigNumber((_dafny.Seq.update(_2210_v62, _module.__default.safeIndex(new BigNumber(-990), new BigNumber((_2210_v62).length)), _2141_v0)).length)), _2141_v0);
              let _2211_v63;
              _2211_v63 = _dafny.Seq.of(_2194_v47);
              let _2212_v64;
              let _nw349 = new _module.C6();
              _nw349.__ctor(_this.f14, _dafny.MultiSet.FromArray((_2211_v63)[_module.__default.safeIndex(_2190_v43, new BigNumber((_2211_v63).length))]));
              _2212_v64 = _nw349;
              let _2213_v65;
              _2213_v65 = _dafny.Seq.UnicodeFromString("ehllenung");
              _2213_v65 = _2213_v65;
              let _2214_v66;
              let _nw350 = new _module.C7();
              _nw350.__ctor();
              _2214_v66 = _nw350;
              _2214_v66 = _2214_v66;
            }
            let _2215_v67;
            let _nw351 = Array((new BigNumber(22)).toNumber());
            _2215_v67 = _nw351;
            _2215_v67 = _2215_v67;
          }
        }
      }
      let _2216_v69;
      let _init61 = ((_2217_v44, _2218_v43) => function (_2219_i4) {
        return ((_this.f14) ? (function () {
          let _coll87 = new _dafny.Map();
          for (const _compr_87 of _dafny.IntegerRange(new BigNumber(-553), new BigNumber(967))) {
            let _2220_v68 = _compr_87;
            if (((new BigNumber(-553)).isLessThanOrEqualTo(_2220_v68)) && ((_2220_v68).isLessThan(new BigNumber(967)))) {
              _coll87.push([(_2220_v68).multipliedBy(_2218_v43),_dafny.Seq.Create(_module.__default.abs(new BigNumber(636)), ((_2221_v43) => function (_2222_i5) {
                return _2221_v43;
              })(_2218_v43))]);
            }
          }
          return _coll87;
        }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_2217_v44, new _dafny.CodePoint('v'.codePointAt(0)))).cardinality()),(_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), ((_2223_v43) => function (_2224_i6) {
          return _2223_v43;
        })(_2218_v43))))));
      })(_2191_v44, _2190_v43);
      let _nw352 = Array((new BigNumber(10)).toNumber());
      for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw352.length); _i0_61++) {
        _nw352[_i0_61] = _init61(new BigNumber(_i0_61));
      }
      _2216_v69 = _nw352;
      let _2225_v70;
      let _nw353 = new _module.C1();
      _nw353.__ctor();
      _2225_v70 = _nw353;
      let _2226_v71;
      _2226_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_2225_v70);
      let _2227_v72;
      _2227_v72 = _dafny.Seq.of(new BigNumber((_2226_v71).length));
      let _2228_v73;
      _2228_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2190_v43,_2227_v72);
      let _index380 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_2216_v69).length));
      (_2216_v69)[_index380] = (_2228_v73).Merge(_2228_v73);
      let _2229_v74;
      _2229_v74 = _dafny.Set.fromElements(_2190_v43);
      let _2230_v75;
      _2230_v75 = _module.D9.create_DC30(_2191_v44, _this.f14, _2190_v43, _2229_v74, _2190_v43);
      let _pat_let_tv48 = _2190_v43;
      let _pat_let_tv49 = _2190_v43;
      let _index381 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_2216_v69).length));
      let _rhs346 = true;
      let _rhs347 = (function (_pat_let70_0) {
        return function (_2231_dt__update__tmp_h0) {
          return function (_pat_let71_0) {
            return function (_2232_dt__update_hcf54_h0) {
              return function (_pat_let72_0) {
                return function (_2233_dt__update_hcf51_h0) {
                  return function (_pat_let73_0) {
                    return function (_2234_dt__update_hcf52_h0) {
                      return _module.D9.create_DC30((_2231_dt__update__tmp_h0).dtor_cf50, _2233_dt__update_hcf51_h0, _2234_dt__update_hcf52_h0, (_2231_dt__update__tmp_h0).dtor_cf53, _2232_dt__update_hcf54_h0);
                    }(_pat_let73_0);
                  }(_pat_let_tv49);
                }(_pat_let72_0);
              }(_this.f14);
            }(_pat_let71_0);
          }(_pat_let_tv48);
        }(_pat_let70_0);
      }(_2230_v75)).dtor_cf50;
      let _rhs348 = _module.__default.fm52(((_dafny.ZERO).minus(_2190_v43)).multipliedBy(_2190_v43), globalState);
      let _lhs295 = globalState;
      let _lhs296 = _2216_v69;
      let _lhs297 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_2216_v69).length));
      _lhs295.f9 = _rhs346;
      _2191_v44 = _rhs347;
      _lhs296[_lhs297] = _rhs348;
      let _2235_v76;
      let _nw354 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _2235_v76 = _nw354;
      let _2236_v77;
      _2236_v77 = _dafny.Seq.of(_2235_v76, _2235_v76);
      _2235_v76 = ((((_this.f14) ? (true) : (_this.f14))) ? (_2235_v76) : ((_2236_v77)[_module.__default.safeIndex(_2190_v43, new BigNumber((_2236_v77).length))]));
      r0 = new BigNumber((_2229_v74).length);
      let _2237_v78;
      _2237_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(287),_this.f14);
      let _2238_v79;
      _2238_v79 = _dafny.Seq.of(_2237_v78);
      r1 = (_module.D9.create_DC30(_2191_v44, _this.f14, _2190_v43, _module.__default.fm22(_2190_v43, _2238_v79, _this.f14, _this.f14, globalState), _2190_v43)).dtor_cf51;
      let _2239_v80;
      _2239_v80 = _dafny.Map.Empty.slice().updateUnsafe(_2237_v78,_2190_v43);
      r2 = _2239_v80;
      r3 = !(_this.f14);
      return [r0, r1, r2, r3];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _2240_v0;
      _2240_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      _2240_v0 = _2240_v0;
      let _2241_v1;
      let _nw355 = Array((new BigNumber(10)).toNumber());
      _nw355[(_dafny.ZERO).toNumber()] = _this.f14;
      _nw355[(_dafny.ONE).toNumber()] = _this.f14;
      _nw355[(new BigNumber(2)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(3)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(4)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(5)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(6)).toNumber()] = false;
      _nw355[(new BigNumber(7)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(8)).toNumber()] = _this.f14;
      _nw355[(new BigNumber(9)).toNumber()] = _this.f14;
      _2241_v1 = _nw355;
      let _index382 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2241_v1).length));
      (_2241_v1)[_index382] = _this.f14;
      let _index383 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2241_v1).length));
      (_2241_v1)[_index383] = !(_this.f14);
      let _2242_v2;
      let _nw356 = new _module.C1();
      _nw356.__ctor();
      _2242_v2 = _nw356;
      (globalState).f9 = !((p1).isLessThanOrEqualTo(p1));
      let _2243_v3;
      let _nw357 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _2243_v3 = _nw357;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2243_v3).length))) {
        let _2244_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2244_i0)) && ((_2244_i0).isLessThan(new BigNumber((_2243_v3).length))))) {
          (_2243_v3)[(_2244_i0)] = _module.__default.safeDivisionInt(_2244_i0, p1);
        }
      }
      let _2245_v4;
      _2245_v4 = _dafny.MultiSet.fromElements((new BigNumber((_2240_v0).length)).isEqualTo(p1));
      _2245_v4 = (_this).f15;
      let _2246_v5;
      _2246_v5 = _dafny.Seq.UnicodeFromString("cxh");
      r0 = (new BigNumber((_dafny.Seq.Concat(_2246_v5, _dafny.Seq.UnicodeFromString("u"))).length)).multipliedBy(p1);
      r1 = p1;
      r2 = (_2241_v1)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2241_v1).length))];
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _2247_v0;
      _2247_v0 = _module.D8.create_DC27();
      _2247_v0 = _2247_v0;
      (globalState).f3 = p0;
      let _2248_v1;
      let _nw358 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _2248_v1 = _nw358;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2248_v1).length))) {
        let _2249_i0 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2249_i0)) && ((_2249_i0).isLessThan(new BigNumber((_2248_v1).length))))) {
          (_2248_v1)[(_2249_i0)] = _module.__default.safeDivisionInt(_2249_i0, p0);
        }
      }
      let _2250_v2;
      _2250_v2 = _dafny.Seq.of(_this.f14, p3, p3);
      let _2251_v3;
      let _nw359 = new _module.C3();
      _nw359.__ctor(p0, p1, p3, (_this).f15);
      _2251_v3 = _nw359;
      let _2252_v4;
      _2252_v4 = _dafny.MultiSet.fromElements(_2251_v3);
      let _2253_v5;
      _2253_v5 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2254_v6;
      _2254_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f14);
      let _2255_v7;
      _2255_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2253_v5,new BigNumber(((_2254_v6).update(p0, _this.f14)).length));
      let _2256_v8;
      _2256_v8 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_2257_v3) => function (_2258_i1) {
        return (_2257_v3).f20;
      })(_2251_v3)));
      let _2259_v9;
      _2259_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23(p0, (_2250_v2)[_module.__default.safeIndex(new BigNumber((_2252_v4).cardinality()), new BigNumber((_2250_v2).length))], _2255_v7, _2256_v8, globalState),(_2251_v3).f20);
      let _2260_v10;
      _2260_v10 = _dafny.MultiSet.fromElements(p0);
      (globalState).f3 = (((_2259_v9).contains(!((_this.f14) && (_this.f14)))) ? ((_2259_v9).get(!((_this.f14) && (_this.f14)))) : ((_dafny.ZERO).minus(_module.__default.fm0(_this.f14, new BigNumber(((_2251_v3).f21).length), p0, _2260_v10, globalState))));
      let _2261_v11;
      let _nw360 = Array((new BigNumber(14)).toNumber()).fill(false);
      _2261_v11 = _nw360;
      (globalState).f2 = _2261_v11;
      let _hi10 = (_2251_v3).f20;
      for (let _2262_i2 = p0; _2262_i2.isLessThan(_hi10); _2262_i2 = _2262_i2.plus(_dafny.ONE)) {
        let _2263_v12;
        let _nw361 = new _module.C5();
        _nw361.__ctor(p0, p3, ((_this).f15).Intersect((_this).f15));
        _2263_v12 = _nw361;
        let _2264_v13;
        _2264_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(751),p1);
        let _2265_v14;
        _2265_v14 = _module.D8.create_DC26(p3, _2264_v13);
        let _2266_v15;
        _2266_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_2263_v12).fm2(p2, p3, globalState));
        let _2267_v16;
        _2267_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,(((_2266_v15).contains(p2)) ? ((_2266_v15).get(p2)) : (p3)));
        let _rhs349 = _2263_v12;
        let _rhs350 = (_dafny.ZERO).minus((((_2265_v14).dtor_cf46) ? ((_2263_v12).f18) : (_module.__default.safeModuloInt((_2263_v12).f18, (_dafny.ZERO).minus(new BigNumber((_2267_v16).length))))));
        let _lhs298 = globalState;
        _2263_v12 = _rhs349;
        _lhs298.f3 = _rhs350;
        let _2268_v17;
        _2268_v17 = _module.D4.create_DC12(_2261_v11);
        _2268_v17 = _2268_v17;
        let _2269_v18;
        let _nw362 = new _module.C5();
        _nw362.__ctor(p0, p3, (_this).f15);
        _2269_v18 = _nw362;
        let _2270_v19;
        _2270_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2269_v18,p2);
        _2270_v19 = _2270_v19;
        let _2271_v20;
        let _nw363 = Array((new BigNumber(9)).toNumber()).fill([]);
        _2271_v20 = _nw363;
        let _index384 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_2271_v20).length));
        (_2271_v20)[_index384] = _2248_v1;
        let _index385 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_2271_v20).length));
        (_2271_v20)[_index385] = _2248_v1;
      }
      r0 = _2253_v5;
      return r0;
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
