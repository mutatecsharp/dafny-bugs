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
      return (new BigNumber(856)).minus((new BigNumber(964)).plus(new BigNumber((_dafny.Seq.of(false)).length)));
    };
    static fm2(p0, p1, p2, globalState) {
      return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("jxdfowin"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("imhcdljya"), _dafny.Seq.UnicodeFromString("gng"))));
    };
    static fm3(globalState) {
      return _dafny.Seq.of(new BigNumber(-386));
    };
    static fm4(globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm7(p0, p1, p2, globalState) {
      return !(((true) ? (new BigNumber((_dafny.Seq.UnicodeFromString("oydcocpi")).length)) : (new BigNumber(904)))).isEqualTo((new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(-604))).Keys.Elements) {
          let _0_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(-604))).contains(_0_v0)) {
            _coll0.push([_0_v0,!(false)]);
          }
        }
        return _coll0;
      }()).length)).plus(new BigNumber((_dafny.Seq.of(new BigNumber(820))).length)));
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(176)), function (_1_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0))));
    };
    static fm9(p0, globalState) {
      return true;
    };
    static fm10(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
    };
    static fm11(p0, globalState) {
      return (_module.D0.create_DC2(!(true), new BigNumber((_dafny.Set.fromElements(new BigNumber(185))).length), _dafny.Seq.of(true, false, true), new BigNumber(273), new BigNumber((_dafny.Seq.of(!(true), false)).length))).dtor_cf6;
    };
    static fm12(p0, p1, globalState) {
      if (true) {
        return _dafny.MultiSet.fromElements(new BigNumber(676), new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(492), new BigNumber(328))) {
            let _2_v0 = _compr_1;
            if (((new BigNumber(492)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(328)))) {
              _coll1.push([(_2_v0).plus(new BigNumber(956)),new BigNumber((_dafny.Seq.of(true)).length)]);
            }
          }
          return _coll1;
        }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber(49));
      } else {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-771)),new BigNumber((_dafny.Seq.of(true, true)).length))).length));
      }
    };
    static fm15(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm16(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('i'.codePointAt(0)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(491)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(true))),new BigNumber(644))));
    };
    static fm17(p0, globalState) {
      return (_dafny.Set.fromElements(true, !(false), true, true)).Difference((_dafny.Set.fromElements(!(true), !(true), false)).Intersect(_dafny.Set.fromElements(true)));
    };
    static fm18(p0, globalState) {
      return _dafny.MultiSet.fromElements(_module.D3.create_DC9(_dafny.Seq.UnicodeFromString("hsnhlv")));
    };
    static fm19(p0, globalState) {
      return (function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), function (_3_i0) {
          return new BigNumber((_dafny.Seq.of(true)).length);
        })).Elements) {
          let _4_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), function (_3_i0) {
            return new BigNumber((_dafny.Seq.of(true)).length);
          }), _4_v0)) {
            _coll2.push([_module.__default.safeModuloInt(_4_v0, new BigNumber((_dafny.Seq.UnicodeFromString("pkhneayht")).length)),new BigNumber(583)]);
          }
        }
        return _coll2;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-379),new BigNumber(887))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D5.create_DC17((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-257), new BigNumber(187), new BigNumber(911), new BigNumber(21))).cardinality()), _module.D0.create_DC1(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, false, false)).length),false)).length),new _dafny.CodePoint('w'.codePointAt(0)))).length))).length), new BigNumber(729))).length), new BigNumber(798)), false, true))).length),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))));
    };
    static fm20(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, true))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, true, true, true)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(!(false), false, false, false), _dafny.Seq.of(false, true))));
    };
    static fm21(p0, globalState) {
      if (!((new BigNumber(192)).isEqualTo(new BigNumber(358)))) {
        return _module.D0.create_DC0(_dafny.MultiSet.fromElements(false, false));
      } else {
        return _module.D0.create_DC0(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false)));
      }
    };
    static fm22(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_dafny.MultiSet.fromElements(false, true)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D5.create_DC17((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length))), new BigNumber((_dafny.Seq.UnicodeFromString("naexdwq")).length), _module.D0.create_DC1(true, new BigNumber((_dafny.Seq.of(new BigNumber(376))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(752),new BigNumber((_dafny.Seq.UnicodeFromString("famhky")).length))).length))).length)), true, true)).dtor_cf33)))),(new BigNumber(889)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("hwit")).length))).length)));
    };
    static fm23(globalState) {
      let _source0 = _module.D1.create_DC4(_dafny.Seq.of(_module.D0.create_DC2(true, new BigNumber((_dafny.Set.fromElements(true, false)).length), _dafny.Seq.of(true), new BigNumber(467), new BigNumber(-669)), _module.D0.create_DC2(true, new BigNumber(195), _dafny.Seq.of(false), new BigNumber(-382), new BigNumber((_dafny.Seq.UnicodeFromString("rqavgwus")).length))));
      if (_source0.is_DC5) {
        let _5___mcc_h0 = (_source0).cf11;
        let _6_cf11 = _5___mcc_h0;
        return _module.D2.create_DC8(new BigNumber(231), (_dafny.ZERO).minus(new BigNumber(-575)));
      } else if (_source0.is_DC4) {
        let _7___mcc_h1 = (_source0).cf10;
        let _8_cf10 = _7___mcc_h1;
        return _module.D2.create_DC8(new BigNumber(527), new BigNumber(49));
      } else {
        let _9___mcc_h2 = (_source0).cf12;
        let _10_cf12 = _9___mcc_h2;
        return _module.D2.create_DC8(new BigNumber((_dafny.Seq.UnicodeFromString("etcx")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("djn"),new BigNumber((_dafny.Seq.UnicodeFromString("vdomg")).length))).length));
      }
    };
    static fm24(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC2(true, new BigNumber((_dafny.MultiSet.fromElements(false, true, true, true, !(false))).cardinality()), _dafny.Seq.of(false), new BigNumber(74), new BigNumber((_dafny.Seq.UnicodeFromString("nrcsueoj")).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), function (_11_i0) {
        return _module.D0.create_DC2(false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-536), new BigNumber((_dafny.Seq.UnicodeFromString("jva")).length))).length)), _dafny.Seq.of(false, false), new BigNumber(546), new BigNumber(-596));
      }));
    };
    static fm25(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, (_module.D0.create_DC2(false, new BigNumber((_dafny.Seq.of(new BigNumber(541))).length), _dafny.Seq.of(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true)).length))).length), new BigNumber(810))).length)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ats")).length))).dtor_cf4, !(!(false)), true)).length),false)).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(660), new BigNumber(885))) {
          let _12_v0 = _compr_3;
          if (((new BigNumber(660)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(885)))) {
            _coll3.push([(_12_v0).minus(new BigNumber(641)),false]);
          }
        }
        return _coll3;
      }());
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC5(!(false) || (false));
    };
    static fm27(p0, p1, globalState) {
      return _module.D0.create_DC1((false) === (false), new BigNumber(192), ((false) ? (new BigNumber((_dafny.Seq.UnicodeFromString("hnqykwbvy")).length)) : (new BigNumber(443))));
    };
    static fm28(p0, p1, globalState) {
      return _dafny.Set.fromElements((_dafny.Set.fromElements((_module.D0.create_DC1(true, new BigNumber(2), new BigNumber(-271))).dtor_cf1, false, false, true, true)).Intersect(_dafny.Set.fromElements(!(true))), (_dafny.Set.fromElements(true, false)).Union(_dafny.Set.fromElements(false, true)), (_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true, true)), (_dafny.Set.fromElements(!(true))).Union(_dafny.Set.fromElements(true)));
    };
    static Main(__noArgsParameter) {
      let _13_v0;
      _13_v0 = new BigNumber(888);
      let _14_v1;
      _14_v1 = _dafny.Seq.of(_13_v0, _13_v0);
      let _15_v2;
      _15_v2 = _dafny.Seq.of(_13_v0, new BigNumber((_14_v1).length));
      let _16_v3;
      let _nw0 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _16_v3 = _nw0;
      let _17_v4;
      _17_v4 = _dafny.Set.fromElements(_13_v0);
      let _18_v5;
      _18_v5 = _dafny.Map.Empty.slice().updateUnsafe(_17_v4,_13_v0);
      let _19_v6;
      let _init0 = function (_20_i0) {
        return true;
      };
      let _nw1 = Array((new BigNumber(14)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _19_v6 = _nw1;
      let _21_v7;
      _21_v7 = _dafny.Seq.of(_16_v3);
      let _22_v8;
      let _nw2 = Array((new BigNumber(3)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = _16_v3;
      _nw2[(_dafny.ONE).toNumber()] = _16_v3;
      _nw2[(new BigNumber(2)).toNumber()] = (_21_v7)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_21_v7).length))];
      _22_v8 = _nw2;
      let _23_globalState;
      let _nw3 = new _module.GlobalState();
      _nw3.__ctor(true, _15_v2, new BigNumber(957), _16_v3, _18_v5, new BigNumber(116), new BigNumber(-467), _dafny.Map.Empty.slice().updateUnsafe(_19_v6,_13_v0), true, _22_v8, _16_v3, true, false, new BigNumber(-352), _16_v3, new BigNumber(833), false, true, new BigNumber(294));
      _23_globalState = _nw3;
      let _24_v9;
      _24_v9 = true;
      let _hi0 = new BigNumber(379);
      for (let _25_i1 = _module.__default.fm0(_24_v9, _24_v9, _23_globalState); _25_i1.isLessThan(_hi0); _25_i1 = _25_i1.plus(_dafny.ONE)) {
        (_23_globalState).f18 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(_24_v9, _24_v9, _23_globalState)));
        let _26_v10;
        let _nw4 = new _module.C0();
        _nw4.__ctor(_13_v0);
        _26_v10 = _nw4;
        let _27_v11;
        let _nw5 = new _module.C2();
        _nw5.__ctor(_26_v10, _24_v9, _19_v6, _24_v9);
        _27_v11 = _nw5;
        let _28_v12;
        _28_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_27_v11.f26),_24_v9);
        _24_v9 = !((_dafny.Map.Empty.slice().updateUnsafe(_27_v11.f26,_module.__default.fm9(_24_v9, _23_globalState))).equals((_28_v12).Merge(_28_v12)));
        (_23_globalState).f13 = (_dafny.ZERO).minus(_25_i1);
      }
      (_23_globalState).f0 = _module.__default.fm7(_13_v0, _13_v0, _13_v0, _23_globalState);
      let _29_v13;
      _29_v13 = new _dafny.CodePoint('j'.codePointAt(0));
      let _30_v14;
      _30_v14 = _dafny.Seq.of(_29_v13);
      let _31_v15;
      let _nw6 = new _module.C3();
      _nw6.__ctor(_module.__default.fm8(_30_v14, _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_24_v9), _13_v0, _23_globalState), _19_v6);
      _31_v15 = _nw6;
      let _32_v16;
      _32_v16 = _dafny.Seq.of(_31_v15, _31_v15, _31_v15, _31_v15);
      let _hi1 = new BigNumber((_17_v4).length);
      for (let _33_i2 = (_13_v0).multipliedBy(new BigNumber((_dafny.Seq.update(_32_v16, _module.__default.safeIndex(_13_v0, new BigNumber((_32_v16).length)), _31_v15)).length)); _33_i2.isLessThan(_hi1); _33_i2 = _33_i2.plus(_dafny.ONE)) {
        (_23_globalState).f0 = !(false) || (_24_v9);
        (_31_v15).f22 = _31_v15.f22;
        let _34_v17;
        _34_v17 = _dafny.Map.Empty.slice().updateUnsafe(_24_v9,_24_v9);
        _34_v17 = _34_v17;
        (_23_globalState).f0 = !(!(_13_v0).isEqualTo(_13_v0));
      }
      let _35_i3;
      _35_i3 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm2(_24_v9, _13_v0, _13_v0, _23_globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_35_i3)) {
              break L0;
            }
            _35_i3 = (_35_i3).plus(_dafny.ONE);
            let _36_v18;
            let _init1 = ((_37_v1) => function (_38_i4) {
              return _37_v1;
            })(_14_v1);
            let _nw7 = Array((new BigNumber(2)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
              _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _36_v18 = _nw7;
            let _39_v19;
            _39_v19 = _module.D4.create_DC12(_29_v13);
            let _40_v20;
            _40_v20 = _dafny.Map.Empty.slice().updateUnsafe(_36_v18,_39_v19);
            _40_v20 = (_40_v20).update(_36_v18, _39_v19);
            let _41_v21;
            _41_v21 = _dafny.Seq.of(_24_v9);
            let _42_v22;
            let _nw8 = new _module.C1();
            _nw8.__ctor(_31_v15.f22, (_31_v15).f23, _24_v9);
            _42_v22 = _nw8;
            let _43_v23;
            _43_v23 = _dafny.Map.Empty.slice().updateUnsafe(_42_v22,_13_v0);
            let _44_v24;
            let _nw9 = new _module.C4();
            _nw9.__ctor((_31_v15).fm5(_module.__default.fm7(new BigNumber((_41_v21).length), (((_43_v23).contains(_42_v22)) ? ((_43_v23).get(_42_v22)) : (_13_v0)), new BigNumber((_30_v14).length), _23_globalState), _23_globalState), _19_v6, (_42_v22).f20);
            _44_v24 = _nw9;
            _44_v24 = _44_v24;
            let _index0 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_31_v15).f23).length));
            ((_31_v15).f23)[_index0] = (_24_v9) === ((_42_v22).f20);
            let _45_v25;
            _45_v25 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_24_v9);
            let _index1 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_31_v15).f23).length));
            ((_31_v15).f23)[_index1] = !_dafny.areEqual(_30_v14, _module.__default.fm8(_30_v14, _45_v25, (_44_v24).f21, _23_globalState));
            let _46_v26;
            let _nw10 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _46_v26 = _nw10;
            let _index2 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_46_v26).length));
            (_46_v26)[_index2] = _31_v15.f22;
            let _index3 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_46_v26).length));
            (_46_v26)[_index3] = _31_v15.f22;
          }
        }
      }
      let _47_v27;
      let _48_v28;
      let _out0;
      let _out1;
      let _outcollector0 = (_31_v15).m2(_23_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _47_v27 = _out0;
      _48_v28 = _out1;
      let _49_i5;
      _49_i5 = _dafny.ZERO;
      L1: {
        while (_24_v9) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_49_i5)) {
              break L1;
            }
            _49_i5 = (_49_i5).plus(_dafny.ONE);
            (_23_globalState).f0 = (_module.__default.safeModuloInt(_13_v0, new BigNumber(155))).isLessThanOrEqualTo(new BigNumber(845));
            let _50_v29;
            _50_v29 = _dafny.Map.Empty.slice().updateUnsafe(_48_v28,_48_v28);
            (_23_globalState).f18 = _module.__default.safeDivisionInt(_13_v0, (_31_v15).fm5((((_50_v29).contains(_48_v28)) ? ((_50_v29).get(_48_v28)) : (_24_v9)), _23_globalState));
            if (!_dafny.Seq.contains(_14_v1, _13_v0)) {
              (_23_globalState).f13 = _13_v0;
              let _51_v30;
              let _nw11 = new _module.C1();
              _nw11.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(963)), ((_52_v13) => function (_53_i6) {
                return _52_v13;
              })(_29_v13)), (_31_v15).f23, _24_v9);
              _51_v30 = _nw11;
              let _54_v31;
              _54_v31 = _dafny.Seq.of(_51_v30, _51_v30);
              let _rhs0 = _dafny.Seq.Concat(_dafny.Seq.of(_51_v30, _51_v30, _51_v30), _54_v31);
              let _rhs1 = (_dafny.ZERO).minus(_13_v0);
              let _lhs0 = _23_globalState;
              _54_v31 = _rhs0;
              _lhs0.f18 = _rhs1;
              let _55_v32;
              let _56_v33;
              let _out2;
              let _out3;
              let _outcollector1 = (_31_v15).m2(_23_globalState);
              _out2 = _outcollector1[0];
              _out3 = _outcollector1[1];
              _55_v32 = _out2;
              _56_v33 = _out3;
              (_23_globalState).f18 = (_module.__default.fm0(_24_v9, _24_v9, _23_globalState)).multipliedBy(_module.__default.safeModuloInt(_13_v0, _13_v0));
              _30_v14 = (((_13_v0).isLessThanOrEqualTo(_13_v0)) ? (_30_v14) : (_31_v15.f22));
            } else {
              let _57_v34;
              let _nw12 = new _module.C1();
              _nw12.__ctor(_dafny.Seq.Concat(_30_v14, _module.__default.fm8(_dafny.Seq.of(_module.__default.fm4(_23_globalState)), _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_24_v9), new BigNumber((_30_v14).length), _23_globalState)), (_31_v15).f23, _24_v9);
              _57_v34 = _nw12;
              _48_v28 = _48_v28;
              let _58_v35;
              let _nw13 = new _module.C1();
              _nw13.__ctor(_30_v14, (_31_v15).f23, _24_v9);
              _58_v35 = _nw13;
              let _59_v36;
              let _nw14 = new _module.C0();
              _nw14.__ctor(new BigNumber((_dafny.Seq.of(new BigNumber((_47_v27).length), _13_v0)).length));
              _59_v36 = _nw14;
              let _60_v37;
              let _nw15 = new _module.C2();
              _nw15.__ctor(_59_v36, _24_v9, (_31_v15).f23, _48_v28);
              _60_v37 = _nw15;
              let _61_v38;
              _61_v38 = _module.D8.create_DC22(_60_v37);
              let _62_v39;
              _62_v39 = _dafny.Set.fromElements((_61_v38).dtor_cf42, _60_v37, _60_v37, _60_v37, _60_v37);
              _62_v39 = (_62_v39).Intersect(_62_v39);
              let _63_v40;
              let _64_v41;
              let _65_v42;
              let _66_v43;
              let _out4;
              let _out5;
              let _out6;
              let _out7;
              let _outcollector2 = (_60_v37).m4(((_59_v36).f24).isLessThan((_59_v36).f24), _23_globalState);
              _out4 = _outcollector2[0];
              _out5 = _outcollector2[1];
              _out6 = _outcollector2[2];
              _out7 = _outcollector2[3];
              _63_v40 = _out4;
              _64_v41 = _out5;
              _65_v42 = _out6;
              _66_v43 = _out7;
            }
            (_31_v15).m3(_23_globalState);
          }
        }
      }
      let _67_v44;
      let _nw16 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
      _67_v44 = _nw16;
      let _68_v45;
      _68_v45 = _dafny.MultiSet.fromElements(_13_v0, new BigNumber(-409));
      let _index4 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_67_v44).length));
      (_67_v44)[_index4] = _68_v45;
      let _index5 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_67_v44).length));
      (_67_v44)[_index5] = (_68_v45).Intersect((_68_v45).Difference(_68_v45));
      let _69_v46;
      _69_v46 = _dafny.Map.Empty.slice().updateUnsafe(_24_v9,_13_v0);
      let _70_v47;
      _70_v47 = _module.D1.create_DC6(_module.__default.fm26(_13_v0, _24_v9, _24_v9, _69_v46, _23_globalState));
      let _71_v48;
      _71_v48 = _module.D1.create_DC6(_70_v47);
      let _pat_let_tv0 = _17_v4;
      let _pat_let_tv1 = _17_v4;
      let _pat_let_tv2 = _17_v4;
      let _pat_let_tv3 = _67_v44;
      let _pat_let_tv4 = _67_v44;
      let _pat_let_tv5 = _67_v44;
      let _pat_let_tv6 = _67_v44;
      _17_v4 = function (_source1) {
        if (_source1.is_DC5) {
          let _72___mcc_h0 = (_source1).cf11;
          let _73_cf11 = _72___mcc_h0;
          return _pat_let_tv0;
        } else if (_source1.is_DC4) {
          let _74___mcc_h1 = (_source1).cf10;
          let _75_cf10 = _74___mcc_h1;
          return _pat_let_tv1;
        } else {
          let _76___mcc_h2 = (_source1).cf12;
          let _77_cf12 = _76___mcc_h2;
          return (_pat_let_tv2).Union(function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of ((_pat_let_tv4)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_pat_let_tv3).length))]).Elements) {
              let _78_v49 = _compr_4;
              if (((_pat_let_tv6)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_pat_let_tv5).length))]).contains(_78_v49)) {
                _coll4.add((_78_v49).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(258))).length)));
              }
            }
            return _coll4;
          }());
        }
      }(_71_v48);
      let _index6 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
      ((_31_v15).f23)[_index6] = false;
      let _79_v50;
      let _nw17 = new _module.C0();
      _nw17.__ctor(_13_v0);
      _79_v50 = _nw17;
      let _80_v51;
      let _nw18 = new _module.C2();
      _nw18.__ctor(_79_v50, false, (_31_v15).f23, false);
      _80_v51 = _nw18;
      let _81_v52;
      _81_v52 = _dafny.Map.Empty.slice().updateUnsafe(_80_v51,_module.__default.fm7((_dafny.ZERO).minus(new BigNumber((_30_v14).length)), (_dafny.ZERO).minus((_79_v50).f24), (_79_v50).f24, _23_globalState));
      let _82_v53;
      _82_v53 = _dafny.MultiSet.fromElements(_29_v13, new _dafny.CodePoint('o'.codePointAt(0)), _module.__default.fm15((_79_v50).f24, (_80_v51).f20, (_79_v50).f24, _23_globalState));
      let _index7 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
      ((_31_v15).f23)[_index7] = !(new BigNumber((_81_v52).length)).isEqualTo(new BigNumber((_82_v53).cardinality()));
      let _source2 = _module.__default.fm27(!(_48_v28), function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(712), new BigNumber(713))) {
          let _83_v54 = _compr_5;
          if (((new BigNumber(712)).isLessThanOrEqualTo(_83_v54)) && ((_83_v54).isLessThan(new BigNumber(713)))) {
            _coll5.push([_module.__default.safeDivisionInt(_83_v54, (_dafny.ZERO).minus((_dafny.ZERO).minus(_13_v0))),_module.D3.create_DC11(_13_v0, _13_v0, _48_v28)]);
          }
        }
        return _coll5;
      }(), _23_globalState);
      if (_source2.is_DC1) {
        let _84___mcc_h3 = (_source2).cf1;
        let _85___mcc_h4 = (_source2).cf2;
        let _86___mcc_h5 = (_source2).cf3;
        let _87_cf3 = _86___mcc_h5;
        let _88_cf2 = _85___mcc_h4;
        let _89_cf1 = _84___mcc_h3;
        _14_v1 = _dafny.Seq.Concat(_15_v2, _dafny.Seq.Concat(_15_v2, _dafny.Seq.of(new BigNumber(648), new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(828), new BigNumber(775))) {
            let _90_v55 = _compr_6;
            if (((new BigNumber(828)).isLessThanOrEqualTo(_90_v55)) && ((_90_v55).isLessThan(new BigNumber(775)))) {
              _coll6.push([(_90_v55).minus((_79_v50).f24),_dafny.Set.fromElements(_29_v13)]);
            }
          }
          return _coll6;
        }()).length))));
        _87_cf3 = _88_cf2;
        let _91_v56;
        let _nw19 = new _module.C2();
        _nw19.__ctor(_79_v50, (_80_v51).f20, (_31_v15).f23, ((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))]);
        _91_v56 = _nw19;
        if ((_91_v56.f26) === ((_80_v51).f20)) {
          let _92_v57;
          _92_v57 = _dafny.MultiSet.fromElements((_80_v51).f20);
          _88_cf2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((_92_v57).update(_91_v56.f26, _module.__default.abs((_14_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_89_cf1,true)).length), new BigNumber((_14_v1).length))]))).cardinality()), (_79_v50).f24));
          let _93_v58;
          _93_v58 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_dafny.Seq.UnicodeFromString("gvadfbvx"));
          let _94_v59;
          _94_v59 = _dafny.Map.Empty.slice().updateUnsafe(_93_v58,(new BigNumber((_31_v15.f22).length)).minus((_79_v50).f24));
          _94_v59 = (_94_v59).update(_dafny.Map.Empty.slice().updateUnsafe((_79_v50).f24,_31_v15.f22), (_88_cf2).plus((_79_v50).f24));
          let _95_v60;
          let _init2 = ((_96_v4) => function (_97_i7) {
            return _96_v4;
          })(_17_v4);
          let _nw20 = Array((new BigNumber(26)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
            _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _95_v60 = _nw20;
          let _index8 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_95_v60).length));
          (_95_v60)[_index8] = _17_v4;
          let _98_v61;
          _98_v61 = _dafny.Set.fromElements(_13_v0);
          let _index9 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
          let _index10 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_95_v60).length));
          let _rhs2 = _module.__default.safeModuloInt((_79_v50).f24, _13_v0);
          let _rhs3 = (_79_v50).f24;
          let _rhs4 = (_80_v51).f20;
          let _rhs5 = (_98_v61);
          let _lhs1 = _23_globalState;
          let _lhs2 = (_31_v15).f23;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
          let _lhs4 = _95_v60;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_95_v60).length));
          _87_cf3 = _rhs2;
          _lhs1.f13 = _rhs3;
          _lhs2[_lhs3] = _rhs4;
          _lhs4[_lhs5] = _rhs5;
          let _99_v62;
          _99_v62 = _dafny.Seq.of((_31_v15).f23);
          let _100_v63;
          let _nw21 = new _module.C3();
          _nw21.__ctor(_dafny.Seq.UnicodeFromString("kfcjm"), (_99_v62)[_module.__default.safeIndex(_13_v0, new BigNumber((_99_v62).length))]);
          _100_v63 = _nw21;
          let _101_v64;
          let _nw22 = new _module.C1();
          _nw22.__ctor(_31_v15.f22, (_31_v15).f23, true);
          _101_v64 = _nw22;
          let _102_v65;
          _102_v65 = _dafny.Seq.of(_dafny.Seq.of(_101_v64, _101_v64));
          _17_v4 = ((_95_v60)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_95_v60).length))]).Difference(_dafny.Set.fromElements(new BigNumber(((_102_v65)[_module.__default.safeIndex((_79_v50).f24, new BigNumber((_102_v65).length))]).length)));
        } else {
          let _103_v66;
          let _init3 = ((_104_v15) => function (_105_i8) {
            return _104_v15.f22;
          })(_31_v15);
          let _nw23 = Array((new BigNumber(21)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
            _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _103_v66 = _nw23;
          let _index11 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_103_v66).length));
          (_103_v66)[_index11] = _30_v14;
          let _index12 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_103_v66).length));
          (_103_v66)[_index12] = _31_v15.f22;
          let _106_v67;
          _106_v67 = _dafny.Set.fromElements(_19_v6);
          _106_v67 = (_106_v67).Intersect(_106_v67);
          _24_v9 = _24_v9;
          (_91_v56).f26 = !(!(_dafny.Set.fromElements((_80_v51).f20, _89_cf1, _91_v56.f26)).contains(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_88_cf2, (_79_v50).f24), _14_v1))));
          let _107_v68;
          let _init4 = ((_108_v27, _109_v50) => function (_110_i9) {
            return _dafny.Seq.update(_108_v27, _module.__default.safeIndex((_109_v50).f24, new BigNumber((_108_v27).length)), true);
          })(_47_v27, _79_v50);
          let _nw24 = Array((_dafny.ONE).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw24.length); _i0_4++) {
            _nw24[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _107_v68 = _nw24;
          _107_v68 = _107_v68;
        }
      } else if (_source2.is_DC2) {
        let _111___mcc_h6 = (_source2).cf4;
        let _112___mcc_h7 = (_source2).cf5;
        let _113___mcc_h8 = (_source2).cf6;
        let _114___mcc_h9 = (_source2).cf7;
        let _115___mcc_h10 = (_source2).cf8;
        let _116_cf8 = _115___mcc_h10;
        let _117_cf7 = _114___mcc_h9;
        let _118_cf6 = _113___mcc_h8;
        let _119_cf5 = _112___mcc_h7;
        let _120_cf4 = _111___mcc_h6;
        (_23_globalState).f0 = _120_cf4;
        let _index13 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_16_v3).length));
        (_16_v3)[_index13] = (_79_v50).f24;
        let _index14 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_16_v3).length));
        (_16_v3)[_index14] = (_79_v50).f24;
        (_23_globalState).f18 = (_79_v50).f24;
        (_31_v15).f22 = _31_v15.f22;
      } else if (_source2.is_DC3) {
        let _121___mcc_h11 = (_source2).cf9;
        let _122_cf9 = _121___mcc_h11;
        _19_v6 = (_31_v15).f23;
        let _123_v69;
        _123_v69 = _dafny.Set.fromElements((_80_v51).f20);
        let _124_v70;
        _124_v70 = _dafny.Seq.of(_123_v69, (_123_v69).Difference(_123_v69), (_123_v69).Intersect(_123_v69), (_module.__default.fm17(new BigNumber((_31_v15.f22).length), _23_globalState)).Intersect(_123_v69));
        let _index15 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
        let _index16 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
        let _rhs6 = _124_v70;
        let _rhs7 = (_80_v51).f20;
        let _rhs8 = !(_24_v9);
        let _rhs9 = true;
        let _rhs10 = ((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))];
        let _lhs6 = _23_globalState;
        let _lhs7 = (_31_v15).f23;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
        let _lhs9 = (_31_v15).f23;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
        _124_v70 = _rhs6;
        _lhs6.f0 = _rhs7;
        _lhs7[_lhs8] = _rhs8;
        _lhs9[_lhs10] = _rhs9;
        _24_v9 = _rhs10;
        let _125_v71;
        let _nw25 = new _module.C0();
        _nw25.__ctor((_13_v0).minus(_13_v0));
        _125_v71 = _nw25;
        (_31_v15).m3(_23_globalState);
      } else {
        let _126___mcc_h12 = (_source2).cf0;
        let _127_cf0 = _126___mcc_h12;
        let _128_v72;
        _128_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm19(_24_v9, _23_globalState)).length), (_79_v50).f24),false);
        _128_v72 = (_128_v72).update(new BigNumber((_dafny.Seq.update(_30_v14, _module.__default.safeIndex((_31_v15).fm5((_80_v51).f20, _23_globalState), new BigNumber((_30_v14).length)), _29_v13)).length), !(_48_v28));
        let _129_v73;
        _129_v73 = _dafny.Set.fromElements(!((_80_v51).f20), _24_v9, _48_v28, _24_v9);
        let _130_v74;
        _130_v74 = _module.D4.create_DC14(_24_v9, _129_v73);
        _130_v74 = _130_v74;
        let _131_v75;
        _131_v75 = _module.D1.create_DC5((_80_v51).f20);
        (_23_globalState).f0 = _module.__default.fm2((_131_v75).dtor_cf11, (((_80_v51).f20) ? (_13_v0) : ((_79_v50).f24)), ((_79_v50).f24).multipliedBy((_79_v50).f24), _23_globalState);
        let _index17 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_16_v3).length));
        (_16_v3)[_index17] = (_79_v50).f24;
        let _132_v76;
        _132_v76 = _dafny.Set.fromElements(_31_v15);
        let _index18 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_16_v3).length));
        (_16_v3)[_index18] = (((((_127_cf0).contains(_24_v9)) ? ((_127_cf0).get(_24_v9)) : (new BigNumber((_17_v4).length)))).plus(new BigNumber((_132_v76).length))).plus((_dafny.ZERO).minus(_13_v0));
      }
      _15_v2 = _14_v1;
      (_31_v15).m3(_23_globalState);
      let _133_v77;
      _133_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_31_v15.f22,_80_v51),((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))]);
      let _134_v78;
      _134_v78 = _dafny.Map.Empty.slice().updateUnsafe((_80_v51).f20,_80_v51);
      let _135_v79;
      _135_v79 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sqqlffck"),(((_134_v78).contains((_80_v51).f20)) ? ((_134_v78).get((_80_v51).f20)) : (_80_v51)));
      let _136_v80;
      let _nw26 = new _module.C4();
      _nw26.__ctor(_module.__default.safeModuloInt(_13_v0, _13_v0), _19_v6, (((_133_v77).contains(_135_v79)) ? ((_133_v77).get(_135_v79)) : (_48_v28)));
      _136_v80 = _nw26;
      let _index19 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
      let _rhs11 = _29_v13;
      let _rhs12 = _13_v0;
      let _rhs13 = _136_v80;
      let _rhs14 = _48_v28;
      let _lhs11 = _23_globalState;
      let _lhs12 = (_31_v15).f23;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
      _29_v13 = _rhs11;
      _lhs11.f13 = _rhs12;
      _136_v80 = _rhs13;
      _lhs12[_lhs13] = _rhs14;
      _16_v3 = _16_v3;
      let _hi2 = _13_v0;
      for (let _137_i10 = _13_v0; _137_i10.isLessThan(_hi2); _137_i10 = _137_i10.plus(_dafny.ONE)) {
        let _138_v81;
        let _nw27 = new _module.C2();
        _nw27.__ctor(_79_v50, (_80_v51).f20, (_80_v51).f19, true);
        _138_v81 = _nw27;
        let _139_v82;
        _139_v82 = _dafny.Seq.of(_138_v81, _138_v81);
        let _140_v83;
        _140_v83 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_139_v82);
        let _141_v84;
        _141_v84 = _dafny.Seq.of(_dafny.Seq.Concat(_139_v82, _139_v82), _139_v82, (((_140_v83).contains(_13_v0)) ? ((_140_v83).get(_13_v0)) : (_139_v82)));
        _141_v84 = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(_138_v81), _139_v82, _139_v82, _139_v82, (_141_v84)[_module.__default.safeIndex(_13_v0, new BigNumber((_141_v84).length))]), _dafny.Seq.Concat(_141_v84, _141_v84));
        let _142_v85;
        _142_v85 = _module.D0.create_DC2(_138_v81.f26, new BigNumber((_dafny.Set.fromElements(_31_v15)).length), _dafny.Seq.of(((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))]), (_136_v80).f21, new BigNumber(945));
        let _143_v86;
        _143_v86 = _module.D0.create_DC2((_142_v85).dtor_cf4, _137_i10, _dafny.Seq.update(_47_v27, _module.__default.safeIndex(_137_i10, new BigNumber((_47_v27).length)), _24_v9), (_136_v80).f21, (_79_v50).f24);
        let _144_v87;
        _144_v87 = _dafny.Map.Empty.slice().updateUnsafe(_31_v15,_138_v81);
        let _145_v88;
        _145_v88 = _dafny.Seq.of(_142_v85, _142_v85, _142_v85, _module.D0.create_DC2(((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))], _137_i10, _47_v27, (_dafny.ZERO).minus(new BigNumber((_30_v14).length)), _137_i10), _module.D0.create_DC2(_138_v81.f26, _137_i10, _dafny.Seq.of(false), new BigNumber((_144_v87).length), new BigNumber((_15_v2).length)));
        let _source3 = _module.D1.create_DC4(_145_v88);
        if (_source3.is_DC5) {
          let _146___mcc_h13 = (_source3).cf11;
          let _147_cf11 = _146___mcc_h13;
          let _nw28 = Array((new BigNumber(10)).toNumber()).fill([]);
          _22_v8 = _nw28;
          (_23_globalState).f13 = (((_80_v51).f20) ? ((_31_v15).fm6((_80_v51).f20, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_47_v27).length),new BigNumber(40))).update((_136_v80).f21, (_136_v80).f21), _dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_148_v13) => function (_149_i11) {
            return _148_v13;
          })(_29_v13)), false, _23_globalState)) : ((_79_v50).f24));
          let _150_v89;
          let _nw29 = new _module.C1();
          _nw29.__ctor(_31_v15.f22, _19_v6, ((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))]);
          _150_v89 = _nw29;
          let _151_v90;
          _151_v90 = _dafny.Map.Empty.slice().updateUnsafe((_module.D8.create_DC23(_48_v28, new BigNumber((_15_v2).length), (_136_v80).f21, _24_v9)).dtor_cf45,_150_v89);
          (_23_globalState).f18 = (((_79_v50).f24).multipliedBy(new BigNumber((_151_v90).length))).multipliedBy(_13_v0);
          _13_v0 = new BigNumber(75);
        } else if (_source3.is_DC4) {
          let _152___mcc_h14 = (_source3).cf10;
          let _153_cf10 = _152___mcc_h14;
          let _154_v91;
          _154_v91 = _dafny.Seq.of(_136_v80);
          let _155_v92;
          let _nw30 = Array((new BigNumber(8)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _136_v80;
          _nw30[(_dafny.ONE).toNumber()] = _136_v80;
          _nw30[(new BigNumber(2)).toNumber()] = _136_v80;
          _nw30[(new BigNumber(3)).toNumber()] = _136_v80;
          _nw30[(new BigNumber(4)).toNumber()] = _136_v80;
          _nw30[(new BigNumber(5)).toNumber()] = _136_v80;
          _nw30[(new BigNumber(6)).toNumber()] = (_154_v91)[_module.__default.safeIndex((_79_v50).f24, new BigNumber((_154_v91).length))];
          _nw30[(new BigNumber(7)).toNumber()] = _136_v80;
          _155_v92 = _nw30;
          let _index20 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_155_v92).length));
          (_155_v92)[_index20] = _136_v80;
          let _index21 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_155_v92).length));
          (_155_v92)[_index21] = _136_v80;
          (_23_globalState).f0 = ((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))];
          let _156_v93;
          let _157_v94;
          let _158_v95;
          let _159_v96;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = (_138_v81).m4((_80_v51).f20, _23_globalState);
          _out8 = _outcollector3[0];
          _out9 = _outcollector3[1];
          _out10 = _outcollector3[2];
          _out11 = _outcollector3[3];
          _156_v93 = _out8;
          _157_v94 = _out9;
          _158_v95 = _out10;
          _159_v96 = _out11;
          _24_v9 = (_80_v51).f20;
        } else {
          let _160___mcc_h15 = (_source3).cf12;
          let _161_cf12 = _160___mcc_h15;
          let _index22 = _module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length));
          ((_31_v15).f23)[_index22] = _48_v28;
          let _162_v97;
          let _nw31 = new _module.C2();
          _nw31.__ctor(_138_v81.f25, (_80_v51).f20, (_80_v51).f19, _138_v81.f26);
          _162_v97 = _nw31;
          let _163_v98;
          _163_v98 = _dafny.Set.fromElements((_80_v51).f20);
          let _164_v99;
          _164_v99 = _dafny.Set.fromElements((_dafny.Set.fromElements((_80_v51).f20)).Difference(_163_v98));
          _164_v99 = _module.__default.fm28(_137_i10, _module.__default.safeModuloInt(_module.__default.fm0((_80_v51).f20, _162_v97.f26, _23_globalState), (_136_v80).f21), _23_globalState);
          (_23_globalState).f13 = _13_v0;
        }
        let _165_v100;
        let _166_v101;
        let _167_v102;
        let _168_v103;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector4 = (_138_v81).m4(!(_48_v28) || (_138_v81.f26), _23_globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _out15 = _outcollector4[3];
        _165_v100 = _out12;
        _166_v101 = _out13;
        _167_v102 = _out14;
        _168_v103 = _out15;
        let _index23 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_16_v3).length));
        (_16_v3)[_index23] = _168_v103;
        let _index24 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_16_v3).length));
        let _rhs15 = (_31_v15);
        let _rhs16 = _168_v103;
        let _lhs14 = _16_v3;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_16_v3).length));
        _31_v15 = _rhs15;
        _lhs14[_lhs15] = _rhs16;
      }
      let _169_v104;
      _169_v104 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_module.__default.fm7(_13_v0, _13_v0, new BigNumber((_30_v14).length), _23_globalState));
      let _170_v107;
      _170_v107 = _module.D0.create_DC1(true, (_136_v80).f21, (_79_v50).f24);
      let _171_v108;
      _171_v108 = _dafny.MultiSet.fromElements((_169_v104).Merge(_dafny.Map.Empty.slice().updateUnsafe((_79_v50).f24,_48_v28)), function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(929), new BigNumber(294))) {
          let _172_v105 = _compr_7;
          if (((new BigNumber(929)).isLessThanOrEqualTo(_172_v105)) && ((_172_v105).isLessThan(new BigNumber(294)))) {
            _coll7.push([(_172_v105).plus((_136_v80).f21),((_31_v15).f23)[_module.__default.safeIndex(new BigNumber(23), new BigNumber(((_31_v15).f23).length))]]);
          }
        }
        return _coll7;
      }(), (function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_14_v1).Elements) {
          let _173_v106 = _compr_8;
          if (_dafny.Seq.contains(_14_v1, _173_v106)) {
            _coll8.push([(_173_v106).multipliedBy((_136_v80).f21),_48_v28]);
          }
        }
        return _coll8;
      }()).Merge((_169_v104).update((_79_v50).f24, (_170_v107).dtor_cf1)));
      (_23_globalState).f13 = (((_171_v108).contains((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_79_v50).f24),!(_48_v28))).Merge(_169_v104))) ? ((_171_v108).get((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_79_v50).f24),!(_48_v28))).Merge(_169_v104))) : ((_136_v80).f21));
      process.stdout.write(_dafny.toString(_13_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_14_v1, _dafny.Seq.of(new BigNumber(888), new BigNumber(2), new BigNumber(888), new BigNumber(2), new BigNumber(648), _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_15_v2, _dafny.Seq.of(new BigNumber(888), new BigNumber(2), new BigNumber(888), new BigNumber(2), new BigNumber(648), _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v4).equals(_dafny.Set.fromElements(new BigNumber(888)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(888)),new BigNumber(888)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_21_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_23_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_23_globalState).f1, _dafny.Seq.of(new BigNumber(888), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState.f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(888)),new BigNumber(888)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_23_globalState).f7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_23_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_23_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_23_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_24_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_29_v13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_30_v14, _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_31_v15.f22, _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_31_v15).f23)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_32_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_35_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_47_v27, _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_48_v28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_49_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_67_v44)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_68_v45).equals(_dafny.MultiSet.fromElements(new BigNumber(888), new BigNumber(-409)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v46).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(888)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_70_v47).dtor_cf12).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_v48).dtor_cf12).dtor_cf12).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_79_v50).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v51).f19)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v51).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_81_v52).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_82_v53).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_133_v77).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_134_v78).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_135_v79).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v80).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v104).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(888),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v107).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v107).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v107).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v108).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(888),true), _dafny.Map.Empty.slice(), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true).updateUnsafe(new BigNumber(888),true)))));
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
    static create_DC2(cf4, cf5, cf6, cf7, cf8) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf9) {
      let $dt = new D0(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC5(cf11) {
      let $dt = new D1(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf10) {
      let $dt = new D1(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D1(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
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
    static create_DC8(cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D2(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC10(cf17, cf18, cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC11(cf21, cf22, cf23) {
      let $dt = new D3(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC9(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + this.cf16.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10([], _dafny.Map.Empty, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC13() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC14(cf25, cf26) {
      let $dt = new D4(1);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC12(cf24) {
      let $dt = new D4(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf27) {
      let $dt = new D4(3);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf27) + ")";
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
        return other.$tag === 1 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13();
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
    static create_DC17(cf29, cf30, cf31, cf32, cf33) {
      let $dt = new D5(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC18(cf34) {
      let $dt = new D5(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC19(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D5(2);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D5(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && this.cf37 === other.cf37 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf28 === other.cf28;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.ZERO, _dafny.ZERO, _module.D0.Default(), false, false);
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
    static create_DC20(cf40) {
      let $dt = new D6(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
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
    static create_DC21(cf41) {
      let $dt = new D7(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41;
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
    static create_DC23(cf43, cf44, cf45, cf46) {
      let $dt = new D8(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC24(cf47) {
      let $dt = new D8(1);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC22(cf42) {
      let $dt = new D8(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC25(cf48) {
      let $dt = new D8(3);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf47 === other.cf47;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf42 === other.cf42;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(false, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC26(cf49) {
      let $dt = new D9(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf50) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50;
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
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
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
          return D11.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f3 = [];
      this.f4 = _dafny.Map.Empty;
      this.f10 = [];
      this.f13 = _dafny.ZERO;
      this.f18 = _dafny.ZERO;
      this._f1 = _dafny.Seq.of();
      this._f2 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Map.Empty;
      this._f8 = false;
      this._f9 = [];
      this._f11 = false;
      this._f12 = false;
      this._f14 = [];
      this._f15 = _dafny.ZERO;
      this._f16 = false;
      this._f17 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
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
    get f9() {
      let _this = this;
      return _this._f9;
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f24 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f19 = [];
      this._f20 = false;
      this._f27 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f27, f19, f20) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return (_this).f20;
    };
    fm14(p0, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), function (_174_i0) {
  return _module.D0.create_DC2((_this).f20, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber(((_this).f27).length))).length), _dafny.Seq.of(!((_this).f20), (_this).f20), new BigNumber(572), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(293),_module.D0.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f20, (_this).f20)),new BigNumber((_dafny.MultiSet.fromElements((_this).f20, (_this).f20, (_this).f20, (_this).f20)).cardinality()))))).length));
})),(_this).f20)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(398)), function (_175_i1) {
  return _module.D0.create_DC2((_this).f20, new BigNumber(756), _dafny.Seq.of((_this).f20), new BigNumber((_dafny.Seq.of(new BigNumber(319))).length), new BigNumber(449));
})),(_this).f20)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), function (_176_i2) {
  return _module.D0.create_DC2((_this).f20, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_module.D0.create_DC1((_this).f20, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f27).length),new BigNumber((_dafny.MultiSet.fromElements((_this).f20)).cardinality()))).length), new BigNumber(-88))).length), new BigNumber(((_this).f27).length))).dtor_cf1)).length), _dafny.Seq.of(false), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(951),new BigNumber((_dafny.Seq.UnicodeFromString("yhnf")).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20)).length));
})),!((_this).f20))));
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f19 = [];
      this._f20 = false;
      this.f25 = undefined;
      this.f26 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f25, f26, f19, f20) {
      let _this = this;
      (_this).f25 = f25;
      (_this).f26 = f26;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _177_v0;
      _177_v0 = _dafny.Seq.of((_this.f25).f24);
      let _hi3 = ((_this.f25).f24).plus(new BigNumber((_177_v0).length));
      for (let _178_i0 = (_this.f25).f24; _178_i0.isLessThan(_hi3); _178_i0 = _178_i0.plus(_dafny.ONE)) {
        let _179_v1;
        _179_v1 = _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)));
        let _180_v2;
        _180_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),(_this).f20);
        _179_v1 = _module.__default.fm8(_179_v1, _180_v2, _178_i0, globalState);
        r1 = (_dafny.ZERO).minus(((_this.f25).f24).plus((_this.f25).f24));
        let _181_v3;
        _181_v3 = _dafny.Seq.of(_this.f26, p0);
        let _182_v4;
        let _nw32 = new _module.C0();
        _nw32.__ctor(new BigNumber((_181_v3).length));
        _182_v4 = _nw32;
        (globalState).f0 = _this.f26;
      }
      let _183_v5;
      _183_v5 = _dafny.Seq.of(_this.f26);
      let _184_v6;
      _184_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this.f25).f24);
      let _185_v7;
      _185_v7 = _dafny.MultiSet.fromElements(_184_v6, _184_v6);
      let _186_v8;
      _186_v8 = _module.D0.create_DC2(true, (_this.f25).f24, _183_v5, (((_185_v7).contains(_184_v6)) ? ((_185_v7).get(_184_v6)) : (new BigNumber((_dafny.Set.fromElements(new BigNumber(137))).length))), (_this.f25).f24);
      let _187_v9;
      _187_v9 = _dafny.Seq.of(_186_v8, _186_v8);
      let _188_v10;
      _188_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_this.f26, (_this).f20)).length),(_this.f25).f24);
      let _189_v11;
      _189_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_188_v10).length),p0);
      let _190_v12;
      _190_v12 = _module.D1.create_DC4(_dafny.Seq.of(_module.D0.create_DC2((_this).f20, (_this.f25).f24, _183_v5, (_dafny.ZERO).minus((_this.f25).f24), new BigNumber((_189_v11).length))));
      let _191_v13;
      let _nw33 = Array((new BigNumber(2)).toNumber());
      _nw33[(_dafny.ZERO).toNumber()] = (((_this).f20) ? (_module.D1.create_DC4(_dafny.Seq.update(_187_v9, _module.__default.safeIndex((_this.f25).f24, new BigNumber((_187_v9).length)), _186_v8))) : (_190_v12));
      _nw33[(_dafny.ONE).toNumber()] = _190_v12;
      _191_v13 = _nw33;
      let _pat_let_tv7 = _187_v9;
      let _index25 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_191_v13).length));
      (_191_v13)[_index25] = function (_pat_let0_0) {
        return function (_192_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_193_dt__update_hcf10_h0) {
              return _module.D1.create_DC4(_193_dt__update_hcf10_h0);
            }(_pat_let1_0);
          }(_pat_let_tv7);
        }(_pat_let0_0);
      }(_190_v12);
      let _index26 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_191_v13).length));
      (_191_v13)[_index26] = _190_v12;
      if ((_this).f20) {
        r0 = _module.__default.fm9((_183_v5)[_module.__default.safeIndex(new BigNumber(-479), new BigNumber((_183_v5).length))], globalState);
        let _194_v14;
        _194_v14 = _dafny.Set.fromElements((_this).f20, _this.f26);
        let _195_v15;
        _195_v15 = _module.D1.create_DC5(false);
        let _196_v16;
        _196_v16 = _dafny.MultiSet.fromElements(_195_v15);
        let _197_v17;
        _197_v17 = _dafny.Map.Empty.slice().updateUnsafe(_194_v14,_196_v16);
        let _198_v18;
        let _nw34 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _198_v18 = _nw34;
        let _199_v19;
        _199_v19 = _dafny.Seq.of(_198_v18, _198_v18);
        let _200_v20;
        _200_v20 = _dafny.Map.Empty.slice().updateUnsafe(_197_v17,(_199_v19)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_199_v19).length))]);
        let _201_v21;
        _201_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this.f25).f24,_194_v14);
        _200_v20 = (_200_v20).update(_dafny.Map.Empty.slice().updateUnsafe((((_201_v21).contains((_this.f25).f24)) ? ((_201_v21).get((_this.f25).f24)) : (_194_v14)),_dafny.MultiSet.fromElements(_195_v15, _module.D1.create_DC5((_this).f20), _195_v15)), _198_v18);
        (globalState).f13 = ((_this.f25).f24).multipliedBy((_this.f25).f24);
        let _202_v22;
        let _nw35 = new _module.C0();
        _nw35.__ctor((_this.f25).f24);
        _202_v22 = _nw35;
        let _203_v23;
        let _nw36 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _203_v23 = _nw36;
        let _index27 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_203_v23).length));
        (_203_v23)[_index27] = (_202_v22).f24;
        let _index28 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_203_v23).length));
        (_203_v23)[_index28] = (_this.f25).f24;
      } else {
        let _204_v24;
        let _nw37 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _204_v24 = _nw37;
        let _205_v25;
        _205_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
        let _index29 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_204_v24).length));
        (_204_v24)[_index29] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_this.f26)).Merge(_205_v25);
        let _206_v26;
        _206_v26 = _dafny.Seq.of((_module.__default.fm10((_this.f25).f24, new BigNumber((_177_v0).length), (_this.f25).f24, globalState)).Merge(_205_v25));
        let _207_v27;
        _207_v27 = _dafny.MultiSet.fromElements(p0, !(_this.f26));
        let _208_v28;
        let _nw38 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _208_v28 = _nw38;
        let _209_v29;
        _209_v29 = _dafny.MultiSet.fromElements(_208_v28);
        let _210_v30;
        _210_v30 = _dafny.Set.fromElements((_this.f25).f24, (((_209_v29).contains(_208_v28)) ? ((_209_v29).get(_208_v28)) : ((_this.f25).f24)));
        let _211_v31;
        _211_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_207_v27).cardinality())),_210_v30);
        let _212_v32;
        _212_v32 = _dafny.Seq.UnicodeFromString("r");
        let _index30 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_204_v24).length));
        (_204_v24)[_index30] = (_206_v26)[_module.__default.safeIndex((new BigNumber((_211_v31).length)).multipliedBy(new BigNumber((_212_v32).length)), new BigNumber((_206_v26).length))];
        _188_v10 = (_188_v10).Merge(((p0) ? (_188_v10) : ((_188_v10).update(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this.f25).f24), _module.__default.safeIndex((_this.f25).f24, new BigNumber((_dafny.Seq.of((_this.f25).f24)).length)), (_this.f25).f24)).length), (_this.f25).f24))));
        let _213_v33;
        _213_v33 = _module.D2.create_DC7(_dafny.Seq.of((_this.f25).f24));
        let _214_v34;
        _214_v34 = _dafny.Set.fromElements(_213_v33);
        let _215_v35;
        _215_v35 = _dafny.Seq.of(_214_v34, _214_v34);
        (_this).f26 = (_dafny.Set.fromElements(_213_v33, _213_v33)).IsDisjointFrom((_215_v35)[_module.__default.safeIndex(new BigNumber((_212_v32).length), new BigNumber((_215_v35).length))]);
        let _216_v36;
        let _init5 = function (_217_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        };
        let _nw39 = Array((new BigNumber(16)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw39.length); _i0_5++) {
          _nw39[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _216_v36 = _nw39;
        _216_v36 = _216_v36;
        (globalState).f0 = (_this).f20;
      }
      let _218_v37;
      let _init6 = function (_219_i2) {
        return (_219_i2).plus((_this.f25).f24);
      };
      let _nw40 = Array((new BigNumber(8)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw40.length); _i0_6++) {
        _nw40[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _218_v37 = _nw40;
      let _index31 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length));
      (_218_v37)[_index31] = new BigNumber((_module.__default.fm11((_this.f25).f24, globalState)).length);
      let _index32 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length));
      (_218_v37)[_index32] = (_this.f25).f24;
      let _220_v38;
      _220_v38 = _dafny.Seq.UnicodeFromString("nmahe");
      let _221_v39;
      _221_v39 = new _dafny.CodePoint('m'.codePointAt(0));
      let _222_v40;
      _222_v40 = _module.D3.create_DC9(_220_v38);
      let _223_v42;
      _223_v42 = _dafny.Set.fromElements((_218_v37)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length))]);
      let _224_v43;
      let _nw41 = Array((new BigNumber(17)).toNumber());
      _nw41[(_dafny.ZERO).toNumber()] = _220_v38;
      _nw41[(_dafny.ONE).toNumber()] = _220_v38;
      _nw41[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_220_v38, _220_v38);
      _nw41[(new BigNumber(3)).toNumber()] = _220_v38;
      _nw41[(new BigNumber(4)).toNumber()] = _220_v38;
      _nw41[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("cnp");
      _nw41[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("iqfjf"), _module.__default.safeIndex((_this.f25).f24, new BigNumber((_dafny.Seq.UnicodeFromString("iqfjf")).length)), _221_v39), _220_v38);
      _nw41[(new BigNumber(7)).toNumber()] = _220_v38;
      _nw41[(new BigNumber(8)).toNumber()] = (_222_v40).dtor_cf16;
      _nw41[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_220_v38, _220_v38);
      _nw41[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_220_v38, _module.__default.safeIndex((_this.f25).f24, new BigNumber((_220_v38).length)), _221_v39);
      _nw41[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("cbiytnt");
      _nw41[(new BigNumber(12)).toNumber()] = _220_v38;
      _nw41[(new BigNumber(13)).toNumber()] = _220_v38;
      _nw41[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uc"), _dafny.Seq.UnicodeFromString("b"));
      _nw41[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("uusrltbmy");
      _nw41[(new BigNumber(16)).toNumber()] = _module.__default.fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(88)), ((_225_v39) => function (_226_i4) {
        return _225_v39;
      })(_221_v39)), (function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_177_v0).Elements) {
          let _227_v41 = _compr_9;
          if (_dafny.Seq.contains(_177_v0, _227_v41)) {
            _coll9.push([(_227_v41).multipliedBy((_this.f25).f24),(_this).f20]);
          }
        }
        return _coll9;
      }()).update(new BigNumber((_223_v42).length), _this.f26), (_218_v37)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length))], globalState);
      _224_v43 = _nw41;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_224_v43).length))) {
        let _228_i3 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_228_i3)) && ((_228_i3).isLessThan(new BigNumber((_224_v43).length))))) {
          (_224_v43)[(_228_i3)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _220_v38), _dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), ((_229_v39) => function (_230_i5) {
            return _229_v39;
          })(_221_v39)));
        }
      }
      let _hi4 = (_218_v37)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length))];
      for (let _231_i6 = ((_this.f25).f24).minus((_this.f25).f24); _231_i6.isLessThan(_hi4); _231_i6 = _231_i6.plus(_dafny.ONE)) {
        r0 = _this.f26;
        (globalState).f13 = (((_this.f26) === (!((_this).f20))) ? ((_this.f25).f24) : (new BigNumber(128)));
        let _232_v44;
        _232_v44 = _dafny.Map.Empty.slice().updateUnsafe(_231_i6,(_this).f19);
        let _233_v45;
        let _nw42 = Array((new BigNumber(13)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = (_this).f19;
        _nw42[(_dafny.ONE).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(2)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(3)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(4)).toNumber()] = (((_232_v44).contains(_module.__default.fm0(p0, _this.f26, globalState))) ? ((_232_v44).get(_module.__default.fm0(p0, _this.f26, globalState))) : ((_this).f19));
        _nw42[(new BigNumber(5)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(6)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(7)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(8)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(9)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(10)).toNumber()] = (((_this).f20) ? ((_this).f19) : ((_this).f19));
        _nw42[(new BigNumber(11)).toNumber()] = (_this).f19;
        _nw42[(new BigNumber(12)).toNumber()] = (_this).f19;
        _233_v45 = _nw42;
        let _index33 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_233_v45).length));
        (_233_v45)[_index33] = (_this).f19;
        let _234_v46;
        _234_v46 = _dafny.Set.fromElements(p0);
        let _235_v47;
        _235_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _236_v48;
        _236_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_177_v0).length),_235_v47);
        let _index34 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_233_v45).length));
        let _nw43 = Array((new BigNumber(7)).toNumber());
        _nw43[(_dafny.ZERO).toNumber()] = _this.f26;
        _nw43[(_dafny.ONE).toNumber()] = (_234_v46).IsProperSubsetOf(_234_v46);
        _nw43[(new BigNumber(2)).toNumber()] = _dafny.Seq.contains(_220_v38, _221_v39);
        _nw43[(new BigNumber(3)).toNumber()] = (_this).f20;
        _nw43[(new BigNumber(4)).toNumber()] = (_this).f20;
        _nw43[(new BigNumber(5)).toNumber()] = (_223_v42).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(((((_236_v48).contains(new BigNumber(120))) ? ((_236_v48).get(new BigNumber(120))) : (_235_v47))).length)));
        _nw43[(new BigNumber(6)).toNumber()] = p0;
        (_233_v45)[_index34] = _nw43;
        let _237_v49;
        _237_v49 = _dafny.MultiSet.fromElements(_module.__default.fm0(_this.f26, _this.f26, globalState), (_this.f25).f24, (_this.f25).f24);
        let _238_v50;
        _238_v50 = _dafny.MultiSet.fromElements(_this.f26, p0);
        (globalState).f0 = (_module.__default.fm12(_237_v49, (((_238_v50).contains(!((_this).f20))) ? ((_238_v50).get(!((_this).f20))) : ((_this.f25).f24)), globalState)).contains(new BigNumber((_234_v46).length));
      }
      r0 = _module.__default.fm9(p0, globalState);
      r1 = new BigNumber(893);
      r2 = (((_184_v6).contains(!(((_this.f25).f24).isLessThanOrEqualTo((_this.f25).f24)))) ? ((_184_v6).get(!(((_this.f25).f24).isLessThanOrEqualTo((_this.f25).f24)))) : ((((_188_v10).contains((_this.f25).f24)) ? ((_188_v10).get((_this.f25).f24)) : ((_this.f25).f24))));
      r3 = (_218_v37)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_218_v37).length))];
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _239_v0;
      _239_v0 = _dafny.Seq.of(_this.f26);
      let _240_v1;
      _240_v1 = _dafny.Seq.of(_module.D0.create_DC2((_this).f20, (_this.f25).f24, _239_v0, (_this.f25).f24, (_this.f25).f24));
      let _241_v2;
      _241_v2 = _module.D1.create_DC4(_240_v1);
      let _pat_let_tv8 = _240_v1;
      let _source4 = function (_pat_let2_0) {
        return function (_242_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_243_dt__update_hcf10_h0) {
              return _module.D1.create_DC4(_243_dt__update_hcf10_h0);
            }(_pat_let3_0);
          }(_pat_let_tv8);
        }(_pat_let2_0);
      }(_241_v2);
      if (_source4.is_DC5) {
        let _244___mcc_h0 = (_source4).cf11;
        let _245_cf11 = _244___mcc_h0;
        let _246_v3;
        let _nw44 = new _module.C0();
        _nw44.__ctor(_module.__default.fm0((_this).f20, _245_cf11, globalState));
        _246_v3 = _nw44;
        let _247_v4;
        _247_v4 = _module.D1.create_DC5(_module.__default.fm9(_this.f26, globalState));
        let _248_v5;
        _248_v5 = _module.D1.create_DC6(_247_v4);
        _248_v5 = _248_v5;
        _245_cf11 = _245_cf11;
        let _249_v6;
        _249_v6 = _dafny.Seq.UnicodeFromString("tclfl");
        let _250_v7;
        let _nw45 = new _module.C1();
        _nw45.__ctor(_249_v6, (_this).f19, _245_cf11);
        _250_v7 = _nw45;
        let _251_v8;
        _251_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_250_v7);
        _251_v8 = (_251_v8).update(true, ((_245_cf11) ? (_250_v7) : (_250_v7)));
      } else if (_source4.is_DC4) {
        let _252___mcc_h1 = (_source4).cf10;
        let _253_cf10 = _252___mcc_h1;
        let _254_v9;
        let _nw46 = Array((new BigNumber(15)).toNumber()).fill([]);
        _254_v9 = _nw46;
        let _255_v10;
        _255_v10 = new _dafny.CodePoint('i'.codePointAt(0));
        let _256_v11;
        _256_v11 = _dafny.MultiSet.fromElements((_this).f20, _this.f26);
        let _257_v12;
        let _nw47 = Array((new BigNumber(26)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _255_v10;
        _nw47[(_dafny.ONE).toNumber()] = _255_v10;
        _nw47[(new BigNumber(2)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(3)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(4)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(5)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(6)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(7)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(8)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(9)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(10)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(11)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(12)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(13)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(14)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw47[(new BigNumber(16)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(17)).toNumber()] = _module.__default.fm15(new BigNumber((_256_v11).cardinality()), _this.f26, (_this.f25).f24, globalState);
        _nw47[(new BigNumber(18)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(19)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
        _nw47[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw47[(new BigNumber(22)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(23)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(24)).toNumber()] = _255_v10;
        _nw47[(new BigNumber(25)).toNumber()] = _255_v10;
        _257_v12 = _nw47;
        let _index35 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_254_v9).length));
        (_254_v9)[_index35] = _257_v12;
        let _index36 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_254_v9).length));
        (_254_v9)[_index36] = _257_v12;
        let _258_v13;
        let _nw48 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _258_v13 = _nw48;
        _258_v13 = _258_v13;
        let _259_v14;
        _259_v14 = _dafny.Seq.UnicodeFromString("w");
        let _260_v15;
        _260_v15 = _module.D3.create_DC9(_259_v14);
        let _261_v16;
        _261_v16 = _module.D3.create_DC9(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), ((_262_v10) => function (_263_i0) {
  return _262_v10;
})(_255_v10)), (_260_v15).dtor_cf16));
        let _source5 = _260_v15;
        if (_source5.is_DC10) {
          let _264___mcc_h3 = (_source5).cf17;
          let _265___mcc_h4 = (_source5).cf18;
          let _266___mcc_h5 = (_source5).cf19;
          let _267___mcc_h6 = (_source5).cf20;
          let _268_cf20 = _267___mcc_h6;
          let _269_cf19 = _266___mcc_h5;
          let _270_cf18 = _265___mcc_h4;
          let _271_cf17 = _264___mcc_h3;
          let _272_v17;
          _272_v17 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), ((_273_v10) => function (_274_i1) {
            return _273_v10;
          })(_255_v10))).length), (_this.f25).f24);
          let _275_v18;
          _275_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_272_v17, _272_v17),((_this.f25).f24).multipliedBy(new BigNumber((_259_v14).length)));
          _275_v18 = (_275_v18).update(_272_v17, (_this.f25).f24);
          let _276_v19;
          _276_v19 = _dafny.Map.Empty.slice().updateUnsafe(_255_v10,(_this.f25).f24);
          let _277_v20;
          _277_v20 = _dafny.MultiSet.fromElements((_this.f25).f24, new BigNumber(((_270_cf18).update(_this.f26, new BigNumber(850))).length), (((_256_v11).contains((_this).f20)) ? ((_256_v11).get((_this).f20)) : (_268_cf20)), (_this.f25).f24, (((_276_v19).contains(_255_v10)) ? ((_276_v19).get(_255_v10)) : ((_this.f25).f24)));
          r1 = (_277_v20).IsDisjointFrom((_277_v20).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_270_cf18).length), (_this.f25).f24)));
          let _index37 = _module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index37] = (_this).f20;
          let _index38 = _module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index38] = !((_this).f20);
          let _278_v21;
          _278_v21 = _dafny.Set.fromElements(_dafny.Seq.contains(_259_v14, _255_v10), ((_this).f20) && (((_this).f19)[_module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f19).length))]));
          _278_v21 = _dafny.Set.fromElements(((_this).f19)[_module.__default.safeIndex(new BigNumber(103), new BigNumber(((_this).f19).length))], (_239_v0)[_module.__default.safeIndex((_this.f25).f24, new BigNumber((_239_v0).length))], true);
        } else if (_source5.is_DC11) {
          let _279___mcc_h7 = (_source5).cf21;
          let _280___mcc_h8 = (_source5).cf22;
          let _281___mcc_h9 = (_source5).cf23;
          let _282_cf23 = _281___mcc_h9;
          let _283_cf22 = _280___mcc_h8;
          let _284_cf21 = _279___mcc_h7;
          _283_cf22 = new BigNumber((_dafny.Seq.update(_239_v0, _module.__default.safeIndex(_module.__default.safeModuloInt(_284_cf21, (_this.f25).f24), new BigNumber((_239_v0).length)), (_283_cf22).isLessThanOrEqualTo((_this.f25).f24))).length);
          let _index39 = _module.__default.safeIndex(new BigNumber(317), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index39] = ((_this).f20) && (_this.f26);
          let _285_v22;
          let _nw49 = new _module.C1();
          _nw49.__ctor(_259_v14, (_this).f19, _282_cf23);
          _285_v22 = _nw49;
          let _286_v23;
          _286_v23 = _dafny.Seq.of(_285_v22, _285_v22);
          let _287_v24;
          let _nw50 = Array((new BigNumber(26)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = (_this.f25).f24;
          _nw50[(_dafny.ONE).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(2)).toNumber()] = new BigNumber(858);
          _nw50[(new BigNumber(3)).toNumber()] = new BigNumber((_259_v14).length);
          _nw50[(new BigNumber(4)).toNumber()] = new BigNumber(46);
          _nw50[(new BigNumber(5)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(6)).toNumber()] = _284_cf21;
          _nw50[(new BigNumber(7)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(8)).toNumber()] = new BigNumber((_259_v14).length);
          _nw50[(new BigNumber(9)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_255_v10, _255_v10)).cardinality());
          _nw50[(new BigNumber(11)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(12)).toNumber()] = new BigNumber((_286_v23).length);
          _nw50[(new BigNumber(13)).toNumber()] = new BigNumber(479);
          _nw50[(new BigNumber(14)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(15)).toNumber()] = new BigNumber(708);
          _nw50[(new BigNumber(16)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(17)).toNumber()] = _283_cf22;
          _nw50[(new BigNumber(18)).toNumber()] = new BigNumber(777);
          _nw50[(new BigNumber(19)).toNumber()] = _284_cf21;
          _nw50[(new BigNumber(20)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(21)).toNumber()] = (_this.f25).f24;
          _nw50[(new BigNumber(22)).toNumber()] = new BigNumber((_239_v0).length);
          _nw50[(new BigNumber(23)).toNumber()] = _283_cf22;
          _nw50[(new BigNumber(24)).toNumber()] = new BigNumber(29);
          _nw50[(new BigNumber(25)).toNumber()] = (_this.f25).f24;
          _287_v24 = _nw50;
          let _288_v25;
          _288_v25 = _module.D0.create_DC0(_256_v11);
          let _289_v26;
          _289_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this.f25).f24,_288_v25);
          let _290_v27;
          _290_v27 = _dafny.Map.Empty.slice().updateUnsafe(_287_v24,_289_v26);
          let _index40 = _module.__default.safeIndex(new BigNumber(317), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index40] = !(_290_v27).equals(_290_v27);
          let _291_v28;
          _291_v28 = _dafny.Seq.of(_module.__default.fm8(_259_v14, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(944),_this.f26), (_this.f25).f24, globalState));
          let _292_v29;
          let _nw51 = new _module.C0();
          _nw51.__ctor(new BigNumber((_291_v28).length));
          _292_v29 = _nw51;
          let _293_v30;
          _293_v30 = _module.D4.create_DC12(_255_v10);
          _255_v10 = (_293_v30).dtor_cf24;
        } else {
          let _294___mcc_h10 = (_source5).cf16;
          let _295_cf16 = _294___mcc_h10;
          let _296_v31;
          let _nw52 = new _module.C1();
          _nw52.__ctor(_259_v14, (_this).f19, ((_this.f25).f24).isLessThan((_this.f25).f24));
          _296_v31 = _nw52;
          _296_v31 = _296_v31;
          let _297_v32;
          _297_v32 = _dafny.Map.Empty.slice().updateUnsafe(_296_v31,_this.f26);
          let _298_v33;
          _298_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_297_v32).length),_this.f26);
          let _299_v34;
          _299_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_298_v33).length),(_this.f25).f24);
          _299_v34 = (_299_v34).update((_this.f25).f24, ((_this.f25).f24).plus((((_256_v11).contains(_this.f26)) ? ((_256_v11).get(_this.f26)) : (new BigNumber(((_296_v31).f27).length)))));
          (globalState).f0 = _this.f26;
          let _300_v35;
          _300_v35 = _dafny.Set.fromElements((_this.f25).f24);
          let _index41 = _module.__default.safeIndex(new BigNumber(666), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index41] = !(!(new BigNumber(670)).isEqualTo((((_256_v11).contains((_this).f20)) ? ((_256_v11).get((_this).f20)) : (new BigNumber((_300_v35).length)))));
          let _301_v36;
          let _nw53 = Array((new BigNumber(16)).toNumber()).fill([]);
          _301_v36 = _nw53;
          let _302_v37;
          _302_v37 = _dafny.Seq.of((_this.f25).f24);
          let _303_v38;
          _303_v38 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f20) ? (new BigNumber((_302_v37).length)) : ((_this.f25).f24)),_301_v36);
          let _304_v39;
          _304_v39 = _dafny.Seq.of(_259_v14);
          let _305_v40;
          _305_v40 = _dafny.Map.Empty.slice().updateUnsafe(_304_v39,_301_v36);
          let _index42 = _module.__default.safeIndex(new BigNumber(666), new BigNumber(((_this).f19).length));
          let _rhs17 = _module.__default.fm0((_256_v11).IsSubsetOf(_256_v11), _this.f26, globalState);
          let _rhs18 = (_this).f20;
          let _rhs19 = (_this.f25).f24;
          let _rhs20 = (((_303_v38).contains((_this.f25).f24)) ? ((_303_v38).get((_this.f25).f24)) : ((((_305_v40).contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("mti")))) ? ((_305_v40).get(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("mti")))) : (_301_v36))));
          let _lhs16 = globalState;
          let _lhs17 = (_this).f19;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(666), new BigNumber(((_this).f19).length));
          let _lhs19 = globalState;
          _lhs16.f18 = _rhs17;
          _lhs17[_lhs18] = _rhs18;
          _lhs19.f18 = _rhs19;
          _301_v36 = _rhs20;
        }
        (globalState).f0 = _this.f26;
      } else {
        let _306___mcc_h2 = (_source4).cf12;
        let _307_cf12 = _306___mcc_h2;
        let _index43 = _module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index43] = (_this).f20;
        let _308_v41;
        _308_v41 = _dafny.Seq.of((_this.f25).f24);
        let _index44 = _module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index44] = _dafny.Seq.IsProperPrefixOf(_308_v41, _308_v41);
        let _309_v42;
        _309_v42 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))],_this.f25);
        let _310_v43;
        _310_v43 = _dafny.Map.Empty.slice().updateUnsafe(_this.f26,(_dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))],_this.f25)).update((_this).f20, _this.f25));
        let _311_v44;
        let _nw54 = Array((new BigNumber(28)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _309_v42;
        _nw54[(_dafny.ONE).toNumber()] = _309_v42;
        _nw54[(new BigNumber(2)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(3)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f25);
        _nw54[(new BigNumber(5)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(6)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(7)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(8)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(9)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(10)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(11)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(12)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(13)).toNumber()] = (_309_v42).update(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))], _this.f25);
        _nw54[(new BigNumber(14)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(15)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f25);
        _nw54[(new BigNumber(17)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(18)).toNumber()] = (((_310_v43).contains((_this).f20)) ? ((_310_v43).get((_this).f20)) : (_309_v42));
        _nw54[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))],_this.f25);
        _nw54[(new BigNumber(20)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(21)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(22)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(23)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(24)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(25)).toNumber()] = _309_v42;
        _nw54[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_239_v0)[_module.__default.safeIndex((_this.f25).f24, new BigNumber((_239_v0).length))],_this.f25);
        _nw54[(new BigNumber(27)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f25);
        _311_v44 = _nw54;
        let _312_v45;
        _312_v45 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))],_module.__default.fm0(_this.f26, ((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))], globalState));
        let _313_v46;
        _313_v46 = _dafny.Seq.of(_312_v45);
        let _314_v47;
        _314_v47 = _module.D3.create_DC10(_311_v44, (_313_v46)[_module.__default.safeIndex(new BigNumber(-519), new BigNumber((_313_v46).length))], _dafny.Set.fromElements(_241_v2), (new BigNumber(932)).minus((_this.f25).f24));
        let _315_v48;
        _315_v48 = new _dafny.CodePoint('l'.codePointAt(0));
        let _316_v49;
        _316_v49 = _dafny.Seq.of(((_this.f26) ? (_315_v48) : (_315_v48)), _315_v48);
        let _317_v50;
        _317_v50 = _dafny.MultiSet.fromElements(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))]);
        let _318_v51;
        _318_v51 = _dafny.Set.fromElements((_317_v50).Intersect(_317_v50));
        let _319_v52;
        _319_v52 = _dafny.MultiSet.fromElements(_241_v2, _241_v2);
        let _320_v53;
        _320_v53 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this.f25).f24), (_this.f25).f24);
        let _321_v54;
        let _nw55 = Array((new BigNumber(24)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = (((_319_v52).contains(_241_v2)) ? ((_319_v52).get(_241_v2)) : ((_this.f25).f24));
        _nw55[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this.f25).f24);
        _nw55[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_this.f25).f24, (_this.f25).f24);
        _nw55[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this.f25).f24), (_dafny.ZERO).minus((_this.f25).f24)));
        _nw55[(new BigNumber(4)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(5)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(6)).toNumber()] = new BigNumber(40);
        _nw55[(new BigNumber(7)).toNumber()] = new BigNumber(670);
        _nw55[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this.f25).f24), (_this.f25).f24);
        _nw55[(new BigNumber(9)).toNumber()] = new BigNumber((_316_v49).length);
        _nw55[(new BigNumber(10)).toNumber()] = new BigNumber(337);
        _nw55[(new BigNumber(11)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(12)).toNumber()] = ((_this.f25).f24).plus((_this.f25).f24);
        _nw55[(new BigNumber(13)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(14)).toNumber()] = new BigNumber((_316_v49).length);
        _nw55[(new BigNumber(15)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(16)).toNumber()] = ((((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))]) ? (new BigNumber((_dafny.Seq.UnicodeFromString("mxgapglw")).length)) : ((_this.f25).f24));
        _nw55[(new BigNumber(17)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(18)).toNumber()] = ((_this.f26) ? (new BigNumber(611)) : ((_dafny.ZERO).minus(_module.__default.fm0(_this.f26, _this.f26, globalState))));
        _nw55[(new BigNumber(19)).toNumber()] = ((_this.f25).f24).multipliedBy((_this.f25).f24);
        _nw55[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("x")).length);
        _nw55[(new BigNumber(21)).toNumber()] = (new BigNumber(283)).minus((_this.f25).f24);
        _nw55[(new BigNumber(22)).toNumber()] = (_this.f25).f24;
        _nw55[(new BigNumber(23)).toNumber()] = (new BigNumber((_317_v50).cardinality())).plus(new BigNumber((_320_v53).cardinality()));
        _321_v54 = _nw55;
        let _index45 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_321_v54).length));
        (_321_v54)[_index45] = (_this.f25).f24;
        let _322_v55;
        _322_v55 = _dafny.Seq.of(_311_v44);
        let _323_v56;
        _323_v56 = _module.D0.create_DC2((_this).f20, new BigNumber(682), _239_v0, new BigNumber((_316_v49).length), (_this.f25).f24);
        let _324_v57;
        _324_v57 = _dafny.Map.Empty.slice().updateUnsafe((_323_v56).dtor_cf7,((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))]);
        let _325_v58;
        _325_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_324_v57).length),_dafny.Seq.of(_315_v48, _315_v48, _315_v48));
        let _pat_let_tv9 = _312_v45;
        let _pat_let_tv10 = _322_v55;
        let _pat_let_tv11 = _322_v55;
        let _index46 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_321_v54).length));
        let _rhs21 = function (_pat_let4_0) {
          return function (_326_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_327_dt__update_hcf18_h0) {
                return function (_pat_let6_0) {
                  return function (_328_dt__update_hcf17_h0) {
                    return _module.D3.create_DC10(_328_dt__update_hcf17_h0, _327_dt__update_hcf18_h0, (_326_dt__update__tmp_h1).dtor_cf19, (_326_dt__update__tmp_h1).dtor_cf20);
                  }(_pat_let6_0);
                }((_pat_let_tv10)[_module.__default.safeIndex((_this.f25).f24, new BigNumber((_pat_let_tv11).length))]);
              }(_pat_let5_0);
            }(_pat_let_tv9);
          }(_pat_let4_0);
        }(_314_v47);
        let _rhs22 = (((_325_v58).contains(new BigNumber((_239_v0).length))) ? ((_325_v58).get(new BigNumber((_239_v0).length))) : (_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0))), _dafny.Seq.of(_315_v48, _315_v48))));
        let _rhs23 = _318_v51;
        let _rhs24 = (_dafny.ZERO).minus((_this.f25).f24);
        let _lhs20 = _321_v54;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_321_v54).length));
        _314_v47 = _rhs21;
        _316_v49 = _rhs22;
        _318_v51 = _rhs23;
        _lhs20[_lhs21] = _rhs24;
        let _329_v59;
        _329_v59 = _module.D3.create_DC11((_this.f25).f24, (_this.f25).f24, !(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))]) || (((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))]));
        let _source6 = _329_v59;
        if (_source6.is_DC10) {
          let _330___mcc_h11 = (_source6).cf17;
          let _331___mcc_h12 = (_source6).cf18;
          let _332___mcc_h13 = (_source6).cf19;
          let _333___mcc_h14 = (_source6).cf20;
          let _334_cf20 = _333___mcc_h14;
          let _335_cf19 = _332___mcc_h13;
          let _336_cf18 = _331___mcc_h12;
          let _337_cf17 = _330___mcc_h11;
          let _338_v60;
          _338_v60 = _module.D4.create_DC12(_315_v48);
          let _pat_let_tv12 = _315_v48;
          let _339_v61;
          _339_v61 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let7_0) {
            return function (_340_dt__update__tmp_h2) {
              return function (_pat_let8_0) {
                return function (_341_dt__update_hcf24_h0) {
                  return _module.D4.create_DC12(_341_dt__update_hcf24_h0);
                }(_pat_let8_0);
              }(_pat_let_tv12);
            }(_pat_let7_0);
          }(_338_v60),(_this.f25).f24);
          _339_v61 = (_339_v61).update(_338_v60, _module.__default.safeModuloInt((_this.f25).f24, (_this.f25).f24));
          (globalState).f0 = !((((_324_v57).contains((_dafny.ZERO).minus((new BigNumber((_239_v0).length)).plus(new BigNumber((_dafny.MultiSet.fromElements((_this.f25).f24)).cardinality()))))) ? ((_324_v57).get((_dafny.ZERO).minus((new BigNumber((_239_v0).length)).plus(new BigNumber((_dafny.MultiSet.fromElements((_this.f25).f24)).cardinality()))))) : ((_this).f20)));
          (_this).f26 = ((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))];
          let _342_v62;
          _342_v62 = _dafny.Set.fromElements(_this.f26, !((_317_v50).update(_this.f26, _module.__default.abs((_this.f25).f24))).equals(_317_v50), (_this).f20, ((_this.f25).f24).isLessThanOrEqualTo((_this.f25).f24));
          let _343_v63;
          _343_v63 = _module.D4.create_DC14(((_this).f19)[_module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f19).length))], _342_v62);
          let _pat_let_tv13 = _342_v62;
          _342_v62 = (function (_pat_let9_0) {
            return function (_344_dt__update__tmp_h3) {
              return function (_pat_let10_0) {
                return function (_345_dt__update_hcf26_h0) {
                  return _module.D4.create_DC14((_344_dt__update__tmp_h3).dtor_cf25, _345_dt__update_hcf26_h0);
                }(_pat_let10_0);
              }(_pat_let_tv13);
            }(_pat_let9_0);
          }(_343_v63)).dtor_cf26;
        } else if (_source6.is_DC11) {
          let _346___mcc_h15 = (_source6).cf21;
          let _347___mcc_h16 = (_source6).cf22;
          let _348___mcc_h17 = (_source6).cf23;
          let _349_cf23 = _348___mcc_h17;
          let _350_cf22 = _347___mcc_h16;
          let _351_cf21 = _346___mcc_h15;
          let _352_v64;
          let _init7 = ((_353_cf22) => function (_354_i2) {
            return (_353_cf22).isLessThan((_this.f25).f24);
          })(_350_cf22);
          let _nw56 = Array((new BigNumber(14)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw56.length); _i0_7++) {
            _nw56[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _352_v64 = _nw56;
          _352_v64 = _352_v64;
          (globalState).f18 = (_321_v54)[_module.__default.safeIndex(new BigNumber(929), new BigNumber((_321_v54).length))];
          let _355_v65;
          let _nw57 = new _module.C1();
          _nw57.__ctor(_316_v49, _352_v64, _this.f26);
          _355_v65 = _nw57;
          let _356_v66;
          _356_v66 = _dafny.Map.Empty.slice().updateUnsafe((_355_v65).f27,_this.f26);
          _356_v66 = (_356_v66).update((_355_v65).f27, !((_this).f20));
        } else {
          let _357___mcc_h18 = (_source6).cf16;
          let _358_cf16 = _357___mcc_h18;
          let _359_v68;
          let _init8 = ((_360_v48) => function (_361_i3) {
            return (_dafny.Set.fromElements(_360_v48)).IsDisjointFrom(function () {
              let _coll10 = new _dafny.Set();
              for (const _compr_10 of (_dafny.Set.fromElements(_360_v48, _360_v48)).Elements) {
                let _362_v67 = _compr_10;
                if ((_dafny.Set.fromElements(_360_v48, _360_v48)).contains(_362_v67)) {
                  _coll10.add(_362_v67);
                }
              }
              return _coll10;
            }());
          })(_315_v48);
          let _nw58 = Array((new BigNumber(8)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw58.length); _i0_8++) {
            _nw58[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _359_v68 = _nw58;
          _359_v68 = _359_v68;
          (globalState).f10 = _321_v54;
          let _363_v69;
          let _nw59 = new _module.C0();
          _nw59.__ctor(((_this.f25).f24).minus((_this.f25).f24));
          _363_v69 = _nw59;
          let _364_v70;
          _364_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f26);
          _364_v70 = _364_v70;
        }
        let _365_v71;
        let _366_v72;
        let _367_v73;
        let _368_v74;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector5 = (_this).m4(!((_this).f20), globalState);
        _out16 = _outcollector5[0];
        _out17 = _outcollector5[1];
        _out18 = _outcollector5[2];
        _out19 = _outcollector5[3];
        _365_v71 = _out16;
        _366_v72 = _out17;
        _367_v73 = _out18;
        _368_v74 = _out19;
      }
      let _369_v75;
      _369_v75 = _dafny.Seq.UnicodeFromString("eneiyaye");
      _369_v75 = _dafny.Seq.UnicodeFromString("rbiwky");
      if (_this.f26) {
        let _370_v76;
        _370_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(955),(_this.f25).f24);
        let _371_v77;
        _371_v77 = _dafny.MultiSet.fromElements((_this.f25).f24, (((_370_v76).contains((_this.f25).f24)) ? ((_370_v76).get((_this.f25).f24)) : (new BigNumber((_369_v75).length))), new BigNumber(403), (_this.f25).f24);
        let _372_v78;
        _372_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_371_v77).update(new BigNumber((_dafny.MultiSet.FromArray(_239_v0)).cardinality()), _module.__default.abs(new BigNumber((_239_v0).length))));
        _371_v77 = ((((_372_v78).contains(_this.f26)) ? ((_372_v78).get(_this.f26)) : (_371_v77))).Difference(_371_v77);
        let _373_v79;
        let _nw60 = new _module.C1();
        _nw60.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), function (_374_i4) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }), (_this).f19, _this.f26);
        _373_v79 = _nw60;
        let _375_v80;
        _375_v80 = _dafny.Seq.of(_373_v79);
        let _376_v81;
        _376_v81 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f26),(_this).f20);
        let _377_v82;
        _377_v82 = _module.D0.create_DC2(_module.__default.fm9((_this).f20, globalState), new BigNumber((_376_v81).length), _dafny.Seq.of((_this).f20, _this.f26), (_this.f25).f24, (_this.f25).f24);
        let _378_v83;
        let _nw61 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _378_v83 = _nw61;
        let _pat_let_tv14 = _239_v0;
        let _379_v84;
        _379_v84 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_375_v80).length)).minus((function (_pat_let11_0) {
          return function (_380_dt__update__tmp_h4) {
            return function (_pat_let12_0) {
              return function (_381_dt__update_hcf5_h0) {
                return function (_pat_let13_0) {
                  return function (_382_dt__update_hcf4_h0) {
                    return function (_pat_let14_0) {
                      return function (_383_dt__update_hcf6_h0) {
                        return _module.D0.create_DC2(_382_dt__update_hcf4_h0, _381_dt__update_hcf5_h0, _383_dt__update_hcf6_h0, (_380_dt__update__tmp_h4).dtor_cf7, (_380_dt__update__tmp_h4).dtor_cf8);
                      }(_pat_let14_0);
                    }(_pat_let_tv14);
                  }(_pat_let13_0);
                }(true);
              }(_pat_let12_0);
            }((_this.f25).f24);
          }(_pat_let11_0);
        }(_377_v82)).dtor_cf7),_378_v83);
        let _384_v85;
        _384_v85 = _module.D3.create_DC9(_369_v75);
        let _385_v86;
        _385_v86 = _dafny.MultiSet.fromElements(_module.D3.create_DC9((_373_v79).f27), _384_v85, _module.D3.create_DC9(_369_v75), _384_v85, _384_v85);
        _379_v84 = (_379_v84).update(new BigNumber((_385_v86).cardinality()), _378_v83);
        let _386_v87;
        _386_v87 = _dafny.MultiSet.fromElements((_this).f20);
        _386_v87 = _386_v87;
        if (!((_this.f26) || ((_this.f26) || (!((_this).f20))))) {
          (globalState).f0 = (_this).f20;
          let _387_v88;
          _387_v88 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this.f25).f24));
          let _rhs25 = (_387_v88).IsProperSubsetOf((_387_v88).Difference(_387_v88));
          let _rhs26 = (_this.f25).f24;
          let _lhs22 = _this;
          let _lhs23 = globalState;
          _lhs22.f26 = _rhs25;
          _lhs23.f13 = _rhs26;
          let _388_v89;
          _388_v89 = _module.D2.create_DC8((_this.f25).f24, _module.__default.fm0(_this.f26, _this.f26, globalState));
          let _rhs27 = (_388_v89).dtor_cf15;
          let _rhs28 = (((_this).f20) ? (_378_v83) : (_378_v83));
          r3 = _rhs27;
          _378_v83 = _rhs28;
          (globalState).f18 = (_this.f25).f24;
          let _389_v90;
          _389_v90 = _dafny.Map.Empty.slice().updateUnsafe((_this.f25).f24,(_this).f20);
          (globalState).f18 = _module.__default.safeDivisionInt((_this.f25).f24, new BigNumber((_module.__default.fm12(_371_v77, new BigNumber(((_389_v90).update(new BigNumber(((_373_v79).f27).length), (_this).f20)).length), globalState)).cardinality()));
        } else {
          let _index47 = _module.__default.safeIndex(new BigNumber(291), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index47] = _this.f26;
          let _390_v91;
          _390_v91 = _module.D1.create_DC5(true);
          let _391_v92;
          _391_v92 = new _dafny.CodePoint('o'.codePointAt(0));
          let _392_v93;
          let _nw62 = new _module.C1();
          _nw62.__ctor(_dafny.Seq.update((_373_v79).f27, _module.__default.safeIndex((_dafny.ZERO).minus((_this.f25).f24), new BigNumber(((_373_v79).f27).length)), _391_v92), (_this).f19, _this.f26);
          _392_v93 = _nw62;
          let _393_v94;
          _393_v94 = _dafny.Seq.of(_392_v93, _392_v93);
          let _394_v95;
          _394_v95 = _dafny.MultiSet.fromElements(_392_v93, (_393_v94)[_module.__default.safeIndex((_this.f25).f24, new BigNumber((_393_v94).length))], _392_v93);
          let _index48 = _module.__default.safeIndex(new BigNumber(291), new BigNumber(((_this).f19).length));
          let _rhs29 = !(new BigNumber((_394_v95).cardinality())).isEqualTo(((_this.f25).f24).multipliedBy((_this.f25).f24));
          let _rhs30 = (_this.f25).f24;
          let _rhs31 = ((_this.f26) ? (_module.D1.create_DC5((_392_v93).f20)) : (_390_v91));
          let _lhs24 = (_this).f19;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(291), new BigNumber(((_this).f19).length));
          _lhs24[_lhs25] = _rhs29;
          r3 = _rhs30;
          _390_v91 = _rhs31;
          r3 = (_module.__default.safeModuloInt(new BigNumber((_369_v75).length), (_dafny.ZERO).minus((_this.f25).f24))).minus((_this.f25).f24);
          let _395_v96;
          _395_v96 = _dafny.Set.fromElements((_this).f20);
          let _396_v97;
          let _nw63 = new _module.C0();
          _nw63.__ctor(_module.__default.safeDivisionInt(new BigNumber((_395_v96).length), (_this.f25).f24));
          _396_v97 = _nw63;
          let _397_v98;
          let _nw64 = Array((new BigNumber(26)).toNumber()).fill(_module.D1.Default());
          _397_v98 = _nw64;
          let _index49 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_397_v98).length));
          (_397_v98)[_index49] = _module.D1.create_DC4(_240_v1);
          let _index50 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_397_v98).length));
          (_397_v98)[_index50] = _241_v2;
          (globalState).f10 = _378_v83;
        }
        let _index51 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index51] = (_this).f20;
        let _index52 = _module.__default.safeIndex(new BigNumber(534), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index52] = (_this).f20;
      } else {
        let _index53 = _module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index53] = true;
        let _index54 = _module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length));
        ((_this).f19)[_index54] = _this.f26;
        let _398_v99;
        let _nw65 = Array((new BigNumber(10)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this.f25).f24);
        _nw65[(_dafny.ONE).toNumber()] = (_this.f25).f24;
        _nw65[(new BigNumber(2)).toNumber()] = (_this.f25).f24;
        _nw65[(new BigNumber(3)).toNumber()] = new BigNumber(226);
        _nw65[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_this.f25).f24)).cardinality());
        _nw65[(new BigNumber(5)).toNumber()] = (_this.f25).f24;
        _nw65[(new BigNumber(6)).toNumber()] = new BigNumber(67);
        _nw65[(new BigNumber(7)).toNumber()] = new BigNumber(418);
        _nw65[(new BigNumber(8)).toNumber()] = (_this.f25).f24;
        _nw65[(new BigNumber(9)).toNumber()] = (_this.f25).f24;
        _398_v99 = _nw65;
        (globalState).f10 = (_module.D5.create_DC16(_398_v99)).dtor_cf28;
        let _399_v100;
        let _nw66 = Array((new BigNumber(5)).toNumber()).fill([]);
        _399_v100 = _nw66;
        let _400_v101;
        let _init9 = function (_401_i5) {
          return _module.D0.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(((_this).f19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length))], (_this).f20, (_this).f20, ((_this).f19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length))]),new BigNumber(118)));
        };
        let _nw67 = Array((new BigNumber(16)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw67.length); _i0_9++) {
          _nw67[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _400_v101 = _nw67;
        let _index55 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_399_v100).length));
        (_399_v100)[_index55] = _400_v101;
        let _index56 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_399_v100).length));
        (_399_v100)[_index56] = _400_v101;
        let _402_v102;
        let _init10 = function (_403_i6) {
          return (_dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length))],(_this).f20)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,((_this).f19)[_module.__default.safeIndex(new BigNumber(329), new BigNumber(((_this).f19).length))]));
        };
        let _nw68 = Array((new BigNumber(11)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw68.length); _i0_10++) {
          _nw68[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _402_v102 = _nw68;
        _402_v102 = _402_v102;
        let _404_v103;
        let _nw69 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _404_v103 = _nw69;
        let _405_v104;
        _405_v104 = _dafny.Set.fromElements(_241_v2);
        let _406_v105;
        _406_v105 = _module.D3.create_DC10(_404_v103, _module.__default.fm16(_this.f26, _this.f26, new BigNumber((_369_v75).length), globalState), _405_v104, (_dafny.ZERO).minus((_this.f25).f24));
        r2 = (_dafny.ZERO).minus((_406_v105).dtor_cf20);
      }
      let _407_v106;
      let _nw70 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _407_v106 = _nw70;
      (globalState).f3 = _407_v106;
      let _408_v107;
      let _nw71 = Array((new BigNumber(29)).toNumber()).fill([]);
      _408_v107 = _nw71;
      let _409_v108;
      let _nw72 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _409_v108 = _nw72;
      let _index57 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_408_v107).length));
      (_408_v107)[_index57] = _409_v108;
      let _index58 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_408_v107).length));
      (_408_v107)[_index58] = _409_v108;
      r3 = ((_this.f25).f24).minus(((_this.f25).f24).plus((_this.f25).f24));
      r0 = _this.f26;
      r1 = ((false) ? ((_this).f20) : (_this.f26));
      r2 = new BigNumber((_dafny.Seq.Concat(_369_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_410_i7) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }))).length);
      r3 = new BigNumber((_dafny.Seq.Concat(_369_v75, _dafny.Seq.Concat(_369_v75, _369_v75))).length);
      return [r0, r1, r2, r3];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this.f22 = _dafny.Seq.UnicodeFromString("");
      this._f23 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22, f23) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source7 = _module.D0.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true)),new BigNumber((_dafny.Seq.of(true, false, true)).length)));
      if (_source7.is_DC1) {
        let _411___mcc_h0 = (_source7).cf1;
        let _412___mcc_h1 = (_source7).cf2;
        let _413___mcc_h2 = (_source7).cf3;
        let _414_cf3 = _413___mcc_h2;
        let _415_cf2 = _412___mcc_h1;
        let _416_cf1 = _411___mcc_h0;
        return new BigNumber(543);
      } else if (_source7.is_DC2) {
        let _417___mcc_h3 = (_source7).cf4;
        let _418___mcc_h4 = (_source7).cf5;
        let _419___mcc_h5 = (_source7).cf6;
        let _420___mcc_h6 = (_source7).cf7;
        let _421___mcc_h7 = (_source7).cf8;
        let _422_cf8 = _421___mcc_h7;
        let _423_cf7 = _420___mcc_h6;
        let _424_cf6 = _419___mcc_h5;
        let _425_cf5 = _418___mcc_h4;
        let _426_cf4 = _417___mcc_h3;
        return (new BigNumber((_dafny.Seq.of(_426_cf4)).length)).minus(_425_cf5);
      } else if (_source7.is_DC3) {
        let _427___mcc_h8 = (_source7).cf9;
        let _428_cf9 = _427___mcc_h8;
        return new BigNumber(4);
      } else {
        let _429___mcc_h9 = (_source7).cf0;
        let _430_cf0 = _429___mcc_h9;
        return _module.__default.safeModuloInt(new BigNumber(206), new BigNumber(-935));
      }
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_431_i0) {
        return _module.D0.create_DC2(!(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(283),!(true))).length), _dafny.Seq.of(true), new BigNumber(78), new BigNumber(231));
      }), _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC2(!(true), new BigNumber(872), _dafny.Seq.of(false, true), new BigNumber(945), new BigNumber((_dafny.Set.fromElements(true)).length))), (_module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_432_i1) {
  return _module.D0.create_DC2(false, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(!(false), false)).length))).length))).cardinality()), _dafny.Seq.of(false), new BigNumber(-264), new BigNumber((_dafny.Seq.of(new BigNumber(428))).length));
}))).dtor_cf10))).length);
    };
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _433_v0;
      _433_v0 = new BigNumber(419);
      let _434_v1;
      _434_v1 = true;
      let _435_v2;
      _435_v2 = _dafny.Set.fromElements(false, _434_v1, _434_v1, _434_v1);
      let _436_v3;
      _436_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-146),_dafny.Set.fromElements(_434_v1));
      let _437_v4;
      _437_v4 = _dafny.Seq.of(new BigNumber((_435_v2).length), new BigNumber(((((_436_v3).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_435_v2).length),_this.f22)).length))) ? ((_436_v3).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_435_v2).length),_this.f22)).length))) : (_435_v2))).length), _433_v0, _433_v0);
      let _438_v5;
      let _nw73 = Array((new BigNumber(4)).toNumber());
      _nw73[(_dafny.ZERO).toNumber()] = _433_v0;
      _nw73[(_dafny.ONE).toNumber()] = (new BigNumber(581)).plus(_433_v0);
      _nw73[(new BigNumber(2)).toNumber()] = (_437_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(_433_v0), new BigNumber((_437_v4).length))];
      _nw73[(new BigNumber(3)).toNumber()] = _433_v0;
      _438_v5 = _nw73;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_438_v5).length))) {
        let _439_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_439_i0)) && ((_439_i0).isLessThan(new BigNumber((_438_v5).length))))) {
          (_438_v5)[(_439_i0)] = _module.__default.safeDivisionInt(_439_i0, _433_v0);
        }
      }
      let _440_v6;
      let _nw74 = new _module.C0();
      _nw74.__ctor(_module.__default.safeModuloInt(_433_v0, _433_v0));
      _440_v6 = _nw74;
      let _441_v7;
      _441_v7 = _dafny.MultiSet.fromElements(_434_v1);
      let _442_v8;
      _442_v8 = _module.D0.create_DC0(_441_v7);
      _442_v8 = _module.D0.create_DC0(_441_v7);
      let _443_v9;
      _443_v9 = _module.D2.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), function (_444_i1) {
  return new BigNumber((_this.f22).length);
}));
      _437_v4 = (_443_v9).dtor_cf13;
      (globalState).f18 = (_433_v0).minus(_433_v0);
      r1 = (new BigNumber((_435_v2).length)).isLessThan(new BigNumber((_dafny.Seq.of(false, _434_v1, true, _434_v1, _434_v1)).length));
      let _445_v10;
      _445_v10 = _dafny.Seq.of(_module.__default.fm7(_433_v0, (_dafny.ZERO).minus(_433_v0), _433_v0, globalState), _dafny.Seq.IsProperPrefixOf(_437_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_446_v0) => function (_447_i2) {
        return _446_v0;
      })(_433_v0))));
      r0 = _445_v10;
      r1 = _434_v1;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _448_v0;
      _448_v0 = true;
      let _449_v1;
      _449_v1 = _dafny.Seq.of(_448_v0, false, _448_v0, _448_v0);
      (globalState).f13 = new BigNumber((((_448_v0) ? (_449_v1) : (_449_v1))).length);
      let _450_v2;
      let _nw75 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _450_v2 = _nw75;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_450_v2).length))) {
        let _451_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_451_i0)) && ((_451_i0).isLessThan(new BigNumber((_450_v2).length))))) {
          (_450_v2)[(_451_i0)] = new _dafny.CodePoint('d'.codePointAt(0));
        }
      }
      let _452_v3;
      let _nw76 = new _module.C1();
      _nw76.__ctor(_dafny.Seq.UnicodeFromString("xxxd"), (_this).f23, _448_v0);
      _452_v3 = _nw76;
      let _453_v4;
      _453_v4 = _dafny.Map.Empty.slice().updateUnsafe(_452_v3,(_this).f23);
      _453_v4 = _453_v4;
      let _454_v5;
      _454_v5 = (_this).f23;
      let _455_v6;
      _455_v6 = _dafny.Set.fromElements((_this).f23, (_this).f23);
      let _456_v7;
      _456_v7 = new BigNumber(-719);
      let _457_v8;
      _457_v8 = _dafny.Map.Empty.slice().updateUnsafe((_455_v6).IsProperSubsetOf(_dafny.Set.fromElements((_this).f23, (_454_v5), (_this).f23, (_this).f23, (_452_v3).f19)),_456_v7);
      let _458_v9;
      _458_v9 = _dafny.Seq.of(_457_v8);
      let _459_v10;
      _459_v10 = _dafny.Map.Empty.slice().updateUnsafe(_457_v8,(_458_v9)[_module.__default.safeIndex(_456_v7, new BigNumber((_458_v9).length))]);
      let _460_v11;
      _460_v11 = _dafny.Map.Empty.slice().updateUnsafe(false,_457_v8);
      let _461_v12;
      let _nw77 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
      _461_v12 = _nw77;
      let _462_v13;
      _462_v13 = _dafny.Set.fromElements(_module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), ((_463_v3, _464_v7, _465_v1) => function (_466_i1) {
  return _module.D0.create_DC2(!((_463_v3).f20), _464_v7, _dafny.Seq.of((_465_v1)[_module.__default.safeIndex(_464_v7, new BigNumber((_465_v1).length))]), _464_v7, _464_v7);
})(_452_v3, _456_v7, _449_v1))));
      let _467_v14;
      _467_v14 = _module.D3.create_DC10(_461_v12, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(_456_v7, new BigNumber(844), _456_v7, globalState),_456_v7), _462_v13, _456_v7);
      _457_v8 = ((((_452_v3).f20) ? ((((_459_v10).contains(_dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,_456_v7))) ? ((_459_v10).get(_dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,_456_v7))) : (_module.__default.fm16((_452_v3).f20, !(_448_v0), _456_v7, globalState)))) : ((((_460_v11).contains((_452_v3).f20)) ? ((_460_v11).get((_452_v3).f20)) : (_457_v8))))).Merge((_467_v14).dtor_cf18);
      if ((_452_v3).f20) {
        (globalState).f18 = _456_v7;
        (globalState).f0 = _448_v0;
        let _rhs32 = (new BigNumber(-257)).minus(_456_v7);
        let _rhs33 = _456_v7;
        let _lhs26 = globalState;
        _456_v7 = _rhs32;
        _lhs26.f18 = _rhs33;
        let _468_v15;
        _468_v15 = _module.D1.create_DC5((_452_v3).f20);
        let _rhs34 = _448_v0;
        let _rhs35 = _468_v15;
        let _rhs36 = !(false) || ((_452_v3).f20);
        let _lhs27 = globalState;
        let _lhs28 = globalState;
        _lhs27.f0 = _rhs34;
        _468_v15 = _rhs35;
        _lhs28.f0 = _rhs36;
        let _469_v16;
        _469_v16 = _dafny.Seq.of((_452_v3).f19, (_452_v3).f19, (_452_v3).f19);
        let _470_v17;
        _470_v17 = new _dafny.CodePoint('o'.codePointAt(0));
        let _471_v18;
        _471_v18 = _dafny.Map.Empty.slice().updateUnsafe(_470_v17,_456_v7);
        let _472_v19;
        _472_v19 = _dafny.Seq.of(_471_v18);
        let _473_v20;
        _473_v20 = _module.D0.create_DC1((_452_v3).f20, new BigNumber(514), _456_v7);
        let _474_v21;
        _474_v21 = _module.D5.create_DC17(new BigNumber((_449_v1).length), _456_v7, _473_v20, (_452_v3).f20, (_452_v3).f20);
        let _475_v22;
        let _init11 = function (_476_i2) {
          return (_476_i2).multipliedBy(new BigNumber((_this.f22).length));
        };
        let _nw78 = Array((new BigNumber(13)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw78.length); _i0_11++) {
          _nw78[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _475_v22 = _nw78;
        let _477_v23;
        _477_v23 = _dafny.MultiSet.fromElements(_475_v22, _475_v22);
        let _478_v24;
        _478_v24 = _dafny.MultiSet.fromElements(new BigNumber((_477_v23).cardinality()), _456_v7);
        let _479_v25;
        _479_v25 = _dafny.Seq.of(_456_v7);
        let _480_v26;
        _480_v26 = _dafny.Seq.of(_dafny.Seq.of((((_478_v24).contains(new BigNumber((_457_v8).length))) ? ((_478_v24).get(new BigNumber((_457_v8).length))) : (_456_v7)), _456_v7, new BigNumber((_479_v25).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), ((_481_v8) => function (_482_i3) {
          return new BigNumber((_481_v8).length);
        })(_457_v8)));
        let _pat_let_tv15 = _456_v7;
        let _pat_let_tv16 = _449_v1;
        let _483_v27;
        _483_v27 = _dafny.Seq.of(function (_pat_let15_0) {
          return function (_484_dt__update__tmp_h0) {
            return function (_pat_let16_0) {
              return function (_485_dt__update_hcf8_h0) {
                return function (_pat_let17_0) {
                  return function (_486_dt__update_hcf6_h0) {
                    return _module.D0.create_DC2((_484_dt__update__tmp_h0).dtor_cf4, (_484_dt__update__tmp_h0).dtor_cf5, _486_dt__update_hcf6_h0, (_484_dt__update__tmp_h0).dtor_cf7, _485_dt__update_hcf8_h0);
                  }(_pat_let17_0);
                }(_pat_let_tv16);
              }(_pat_let16_0);
            }(_pat_let_tv15);
          }(_pat_let15_0);
        }(_module.D0.create_DC2((_452_v3).f20, _456_v7, _449_v1, _456_v7, new BigNumber((_this.f22).length))));
        let _487_v28;
        _487_v28 = _module.D1.create_DC4(_483_v27);
        let _488_v29;
        _488_v29 = _dafny.Seq.of(_this.f22);
        let _pat_let_tv17 = _452_v3;
        let _pat_let_tv18 = _456_v7;
        let _pat_let_tv19 = _449_v1;
        let _pat_let_tv20 = _456_v7;
        let _pat_let_tv21 = _488_v29;
        let _pat_let_tv22 = _456_v7;
        let _pat_let_tv23 = _488_v29;
        let _489_v30;
        _489_v30 = _dafny.MultiSet.fromElements(function (_pat_let18_0) {
          return function (_490_dt__update__tmp_h1) {
            return function (_pat_let19_0) {
              return function (_491_dt__update_hcf10_h0) {
                return _module.D1.create_DC4(_491_dt__update_hcf10_h0);
              }(_pat_let19_0);
            }(_dafny.Seq.of(_module.D0.create_DC2(!((_pat_let_tv17).f20), _pat_let_tv18, _pat_let_tv19, _pat_let_tv20, new BigNumber(((_pat_let_tv21)[_module.__default.safeIndex(_pat_let_tv22, new BigNumber((_pat_let_tv23).length))]).length))));
          }(_pat_let18_0);
        }(_487_v28), _487_v28);
        let _492_v31;
        let _nw79 = Array((new BigNumber(27)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = new BigNumber((_469_v16).length);
        _nw79[(_dafny.ONE).toNumber()] = new BigNumber((_this.f22).length);
        _nw79[(new BigNumber(2)).toNumber()] = ((_448_v0) ? (_456_v7) : (new BigNumber(851)));
        _nw79[(new BigNumber(3)).toNumber()] = (_module.__default.fm0(false, (_452_v3).f20, globalState)).multipliedBy(_456_v7);
        _nw79[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_456_v7);
        _nw79[(new BigNumber(5)).toNumber()] = (_456_v7).plus(_456_v7);
        _nw79[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_456_v7);
        _nw79[(new BigNumber(7)).toNumber()] = new BigNumber(398);
        _nw79[(new BigNumber(8)).toNumber()] = new BigNumber(((_dafny.Set.fromElements((_452_v3).f20)).Difference(_module.__default.fm17(new BigNumber((_this.f22).length), globalState))).length);
        _nw79[(new BigNumber(9)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qnxjphmtm")).length);
        _nw79[(new BigNumber(11)).toNumber()] = new BigNumber((_472_v19).length);
        _nw79[(new BigNumber(12)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(13)).toNumber()] = _module.__default.fm0((_452_v3).f20, (_474_v21).dtor_cf33, globalState);
        _nw79[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(_456_v7, _456_v7);
        _nw79[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_456_v7)).multipliedBy(_456_v7));
        _nw79[(new BigNumber(16)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(17)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(18)).toNumber()] = new BigNumber(((_480_v26)[_module.__default.safeIndex(_456_v7, new BigNumber((_480_v26).length))]).length);
        _nw79[(new BigNumber(19)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(20)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(21)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(22)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(23)).toNumber()] = _456_v7;
        _nw79[(new BigNumber(24)).toNumber()] = (_456_v7).multipliedBy(new BigNumber((_489_v30).cardinality()));
        _nw79[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_this.f22, _dafny.Seq.UnicodeFromString("w"))).length);
        _nw79[(new BigNumber(26)).toNumber()] = _456_v7;
        _492_v31 = _nw79;
        let _index59 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_475_v22).length));
        (_475_v22)[_index59] = _456_v7;
        let _493_v32;
        let _nw80 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _493_v32 = _nw80;
        let _494_v33;
        _494_v33 = _module.D5.create_DC19(new BigNumber((_478_v24).cardinality()), (_452_v3).f20, _493_v32, (_452_v3).f20, _470_v17);
        let _index60 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_475_v22).length));
        let _rhs37 = (_456_v7).minus(_456_v7);
        let _rhs38 = (_449_v1)[_module.__default.safeIndex(_456_v7, new BigNumber((_449_v1).length))];
        let _rhs39 = (_494_v33).dtor_cf36;
        let _rhs40 = _456_v7;
        let _lhs29 = _475_v22;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_475_v22).length));
        let _lhs31 = globalState;
        let _lhs32 = globalState;
        let _lhs33 = globalState;
        _lhs29[_lhs30] = _rhs37;
        _lhs31.f0 = _rhs38;
        _lhs32.f0 = _rhs39;
        _lhs33.f13 = _rhs40;
      } else {
        (globalState).f0 = ((_dafny.MultiSet.fromElements((_452_v3).f20)).update((_452_v3).f20, _module.__default.abs(_456_v7))).IsDisjointFrom(_dafny.MultiSet.fromElements(_448_v0, !((_452_v3).f20), _448_v0));
        if (_448_v0) {
          let _495_v34;
          let _nw81 = new _module.C0();
          _nw81.__ctor(_456_v7);
          _495_v34 = _nw81;
          let _496_v35;
          _496_v35 = new _dafny.CodePoint('p'.codePointAt(0));
          let _497_v36;
          _497_v36 = _dafny.Map.Empty.slice().updateUnsafe(_496_v35,(_452_v3).f19);
          let _498_v37;
          let _nw82 = new _module.C2();
          _nw82.__ctor(_495_v34, (_452_v3).f20, ((!((_452_v3).f20)) ? ((_452_v3).f19) : ((((_497_v36).contains(_496_v35)) ? ((_497_v36).get(_496_v35)) : ((_452_v3).f19)))), _module.__default.fm9((_452_v3).f20, globalState));
          _498_v37 = _nw82;
          let _499_v38;
          let _nw83 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _499_v38 = _nw83;
          let _500_v39;
          _500_v39 = _dafny.MultiSet.fromElements(false, (_452_v3).f20);
          let _index61 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_499_v38).length));
          (_499_v38)[_index61] = new BigNumber((_dafny.Seq.of(_500_v39)).length);
          let _index62 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_499_v38).length));
          (_499_v38)[_index62] = _456_v7;
          _456_v7 = _module.__default.safeModuloInt((_456_v7).plus((_495_v34).f24), (_499_v38)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_499_v38).length))]);
          let _index63 = _module.__default.safeIndex(new BigNumber(300), new BigNumber(((_452_v3).f19).length));
          ((_452_v3).f19)[_index63] = !(((_452_v3).f20) || ((_452_v3).f20));
          let _index64 = _module.__default.safeIndex(new BigNumber(300), new BigNumber(((_452_v3).f19).length));
          ((_452_v3).f19)[_index64] = false;
          let _501_v40;
          _501_v40 = _dafny.Seq.of(((_499_v38)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_499_v38).length))]).multipliedBy((_499_v38)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_499_v38).length))]), _456_v7);
          (globalState).f13 = (_501_v40)[_module.__default.safeIndex((_495_v34).f24, new BigNumber((_501_v40).length))];
        } else {
          let _502_v41;
          _502_v41 = _module.D3.create_DC9(_this.f22);
          let _503_v42;
          _503_v42 = _dafny.Map.Empty.slice().updateUnsafe(_448_v0,_dafny.MultiSet.fromElements(_502_v41, _502_v41));
          let _504_v43;
          _504_v43 = _dafny.MultiSet.fromElements(_502_v41);
          let _505_v44;
          _505_v44 = _dafny.Seq.of(_456_v7, new BigNumber(448), _456_v7, _456_v7);
          (globalState).f0 = (_module.__default.fm18(_505_v44, globalState)).IsSubsetOf(((((_503_v42).contains((_452_v3).f20)) ? ((_503_v42).get((_452_v3).f20)) : (_504_v43))).Union(_504_v43));
          let _506_v45;
          _506_v45 = _dafny.Map.Empty.slice().updateUnsafe(_448_v0,(_452_v3).f20);
          _506_v45 = (_506_v45).update((_452_v3).f20, (_452_v3).f20);
          let _507_v46;
          _507_v46 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,!((_452_v3).f20));
          let _508_v47;
          let _nw84 = new _module.C0();
          _nw84.__ctor(new BigNumber((_507_v46).length));
          _508_v47 = _nw84;
          let _509_v48;
          let _nw85 = new _module.C2();
          _nw85.__ctor(_508_v47, (_452_v3).f20, (_452_v3).f19, _dafny.Seq.IsProperPrefixOf(_this.f22, _this.f22));
          _509_v48 = _nw85;
          let _510_v49;
          let _nw86 = Array((_dafny.ONE).toNumber()).fill(_module.D0.Default());
          _510_v49 = _nw86;
          let _511_v50;
          _511_v50 = _510_v49;
          let _512_v51;
          let _nw87 = Array((new BigNumber(6)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = _510_v49;
          _nw87[(_dafny.ONE).toNumber()] = _510_v49;
          _nw87[(new BigNumber(2)).toNumber()] = _510_v49;
          _nw87[(new BigNumber(3)).toNumber()] = _510_v49;
          _nw87[(new BigNumber(4)).toNumber()] = _510_v49;
          _nw87[(new BigNumber(5)).toNumber()] = (_511_v50);
          _512_v51 = _nw87;
          let _index65 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_512_v51).length));
          (_512_v51)[_index65] = _510_v49;
          let _index66 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_512_v51).length));
          (_512_v51)[_index66] = _510_v49;
          let _513_v52;
          _513_v52 = _dafny.MultiSet.fromElements(!(_448_v0), (_452_v3).f20, (_452_v3).f20, (_452_v3).f20, (_452_v3).f20);
          (globalState).f0 = ((((_513_v52).contains(!(_448_v0))) ? ((_513_v52).get(!(_448_v0))) : (new BigNumber((_module.__default.fm19(true, globalState)).length)))).isLessThan((_508_v47).f24);
        }
        let _514_v53;
        let _515_v54;
        let _out20;
        let _out21;
        let _outcollector6 = (_this).m2(globalState);
        _out20 = _outcollector6[0];
        _out21 = _outcollector6[1];
        _514_v53 = _out20;
        _515_v54 = _out21;
        let _516_v55;
        _516_v55 = new _dafny.CodePoint('o'.codePointAt(0));
        _516_v55 = _516_v55;
        (globalState).f13 = ((_dafny.ZERO).minus(_456_v7)).minus((_456_v7).plus(_456_v7));
      }
      if (!(_448_v0)) {
        _456_v7 = new BigNumber(-455);
        let _517_v56;
        _517_v56 = _dafny.Set.fromElements(_456_v7);
        let _rhs41 = _456_v7;
        let _rhs42 = (_452_v3).f20;
        let _rhs43 = (_452_v3).f20;
        let _rhs44 = ((((_452_v3).f20) ? (_517_v56) : (_517_v56))).IsSubsetOf(_517_v56);
        let _lhs34 = globalState;
        let _lhs35 = globalState;
        let _lhs36 = globalState;
        let _lhs37 = globalState;
        _lhs34.f13 = _rhs41;
        _lhs35.f0 = _rhs42;
        _lhs36.f0 = _rhs43;
        _lhs37.f0 = _rhs44;
        (globalState).f13 = (((_452_v3).f20) ? (new BigNumber((_457_v8).length)) : (_456_v7));
        let _518_v57;
        _518_v57 = new _dafny.CodePoint('p'.codePointAt(0));
        (globalState).f13 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_this.f22, _dafny.Seq.update(_this.f22, _module.__default.safeIndex(_456_v7, new BigNumber((_this.f22).length)), _518_v57)), _dafny.Seq.UnicodeFromString("usturyl"))).length);
        if ((_452_v3).f20) {
          let _519_v58;
          _519_v58 = _module.D4.create_DC13();
          _519_v58 = _519_v58;
          let _520_v59;
          _520_v59 = _dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,_518_v57);
          let _521_v60;
          _521_v60 = _dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,(_520_v59).contains(_448_v0));
          _521_v60 = ((_521_v60).Merge(_521_v60)).Merge(_521_v60);
          let _522_v61;
          _522_v61 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("i"));
          let _523_v62;
          let _nw88 = Array((new BigNumber(5)).toNumber()).fill([]);
          _523_v62 = _nw88;
          let _524_v63;
          _524_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_522_v61).length),_523_v62);
          let _525_v64;
          _525_v64 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,(((_524_v63).contains(new BigNumber(680))) ? ((_524_v63).get(new BigNumber(680))) : (_523_v62)));
          _525_v64 = (_525_v64).update((((_452_v3).f20) ? (_456_v7) : (_456_v7)), _523_v62);
          (globalState).f0 = ((_module.__default.fm7(new BigNumber((_module.__default.fm19(true, globalState)).length), _456_v7, _456_v7, globalState)) ? (!((_452_v3).f20)) : (_dafny.Seq.contains(_449_v1, _448_v0)));
          let _526_v65;
          _526_v65 = _module.D0.create_DC1((_452_v3).f20, _456_v7, _456_v7);
          let _527_v66;
          let _nw89 = Array((new BigNumber(9)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _456_v7;
          _nw89[(_dafny.ONE).toNumber()] = _module.__default.fm0((_526_v65).dtor_cf1, (_452_v3).f20, globalState);
          _nw89[(new BigNumber(2)).toNumber()] = _456_v7;
          _nw89[(new BigNumber(3)).toNumber()] = _456_v7;
          _nw89[(new BigNumber(4)).toNumber()] = new BigNumber((_449_v1).length);
          _nw89[(new BigNumber(5)).toNumber()] = new BigNumber(139);
          _nw89[(new BigNumber(6)).toNumber()] = _456_v7;
          _nw89[(new BigNumber(7)).toNumber()] = _456_v7;
          _nw89[(new BigNumber(8)).toNumber()] = _456_v7;
          _527_v66 = _nw89;
          let _528_v67;
          _528_v67 = _module.D5.create_DC16(_527_v66);
          let _529_v68;
          _529_v68 = _dafny.Map.Empty.slice().updateUnsafe(_528_v67,(_452_v3).f19);
          let _530_v69;
          let _nw90 = new _module.C1();
          _nw90.__ctor(_dafny.Seq.UnicodeFromString("sjlkirkt"), (((_529_v68).contains(_528_v67)) ? ((_529_v68).get(_528_v67)) : ((_452_v3).f19)), (_449_v1)[_module.__default.safeIndex(_456_v7, new BigNumber((_449_v1).length))]);
          _530_v69 = _nw90;
        } else {
          (globalState).f18 = (((_457_v8).contains((_452_v3).f20)) ? ((_457_v8).get((_452_v3).f20)) : (new BigNumber((_dafny.Seq.of((_452_v3).f20)).length)));
          _456_v7 = new BigNumber((_this.f22).length);
          let _531_v70;
          _531_v70 = _module.D0.create_DC1((_452_v3).f20, new BigNumber(310), (_dafny.ZERO).minus(_456_v7));
          let _532_v71;
          _532_v71 = _dafny.MultiSet.fromElements(_module.D0.create_DC1((_452_v3).f20, _456_v7, _456_v7), _531_v70, _531_v70, _531_v70);
          let _533_v72;
          let _nw91 = new _module.C1();
          _nw91.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), ((_534_v57) => function (_535_i4) {
            return _534_v57;
          })(_518_v57)), (_this).f23, (_452_v3).f20);
          _533_v72 = _nw91;
          let _536_v73;
          _536_v73 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,_533_v72);
          let _537_v74;
          _537_v74 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,new BigNumber((_536_v73).length));
          let _538_v75;
          _538_v75 = _dafny.Seq.of((((_532_v71).contains(_module.D0.create_DC1((_452_v3).f20, (((_537_v74).contains(_456_v7)) ? ((_537_v74).get(_456_v7)) : (_456_v7)), _456_v7))) ? ((_532_v71).get(_module.D0.create_DC1((_452_v3).f20, (((_537_v74).contains(_456_v7)) ? ((_537_v74).get(_456_v7)) : (_456_v7)), _456_v7))) : (_456_v7)), _456_v7, new BigNumber((_dafny.MultiSet.fromElements(_456_v7)).cardinality()));
          let _539_v76;
          let _nw92 = new _module.C0();
          _nw92.__ctor(_456_v7);
          _539_v76 = _nw92;
          let _540_v77;
          _540_v77 = _dafny.Seq.of(_539_v76);
          let _541_v78;
          _541_v78 = _module.D5.create_DC18(_540_v77);
          let _542_v79;
          _542_v79 = _dafny.Map.Empty.slice().updateUnsafe(_541_v78,(_452_v3).f20);
          let _543_v80;
          let _nw93 = Array((new BigNumber(16)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _456_v7;
          _nw93[(_dafny.ONE).toNumber()] = _456_v7;
          _nw93[(new BigNumber(2)).toNumber()] = _456_v7;
          _nw93[(new BigNumber(3)).toNumber()] = _456_v7;
          _nw93[(new BigNumber(4)).toNumber()] = _456_v7;
          _nw93[(new BigNumber(5)).toNumber()] = (_538_v75)[_module.__default.safeIndex(_456_v7, new BigNumber((_538_v75).length))];
          _nw93[(new BigNumber(6)).toNumber()] = _456_v7;
          _nw93[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_456_v7, _456_v7));
          _nw93[(new BigNumber(8)).toNumber()] = (_456_v7).minus(_456_v7);
          _nw93[(new BigNumber(9)).toNumber()] = new BigNumber((_542_v79).length);
          _nw93[(new BigNumber(10)).toNumber()] = (_this).fm5((_452_v3).f20, globalState);
          _nw93[(new BigNumber(11)).toNumber()] = (new BigNumber((_449_v1).length)).plus((_539_v76).f24);
          _nw93[(new BigNumber(12)).toNumber()] = _456_v7;
          _nw93[(new BigNumber(13)).toNumber()] = new BigNumber(((((_452_v3).f20) ? (_449_v1) : (_449_v1))).length);
          _nw93[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_this.f22, _this.f22), _module.__default.safeIndex((_539_v76).f24, new BigNumber((_dafny.Seq.Concat(_this.f22, _this.f22)).length)), _518_v57)).length);
          _nw93[(new BigNumber(15)).toNumber()] = (_456_v7).multipliedBy((_dafny.ZERO).minus((_539_v76).f24));
          _543_v80 = _nw93;
          let _index67 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_543_v80).length));
          (_543_v80)[_index67] = (_539_v76).f24;
          let _index68 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_543_v80).length));
          (_543_v80)[_index68] = _module.__default.safeDivisionInt((_539_v76).f24, new BigNumber((_dafny.Seq.of((_452_v3).f20, _448_v0)).length));
          let _index69 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_543_v80).length));
          let _index70 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_543_v80).length));
          let _rhs45 = (_456_v7).minus((_539_v76).f24);
          let _rhs46 = _456_v7;
          let _rhs47 = _456_v7;
          let _lhs38 = globalState;
          let _lhs39 = _543_v80;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_543_v80).length));
          let _lhs41 = _543_v80;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_543_v80).length));
          _lhs38.f13 = _rhs45;
          _lhs39[_lhs40] = _rhs46;
          _lhs41[_lhs42] = _rhs47;
          let _544_v81;
          _544_v81 = _module.D4.create_DC14((_452_v3).f20, _dafny.Set.fromElements((_452_v3).f20));
          (globalState).f0 = (((_452_v3).f20) ? ((_544_v81).dtor_cf25) : ((_452_v3).f20));
          (globalState).f13 = _456_v7;
        }
      } else {
        let _545_v82;
        _545_v82 = _dafny.Set.fromElements((_452_v3).f20);
        let _546_v83;
        _546_v83 = _dafny.MultiSet.fromElements(_456_v7, _456_v7);
        let _547_v84;
        _547_v84 = _dafny.Seq.of(new BigNumber((_546_v83).cardinality()));
        let _548_v85;
        _548_v85 = _dafny.Map.Empty.slice().updateUnsafe(_545_v82,(_547_v84)[_module.__default.safeIndex(_456_v7, new BigNumber((_547_v84).length))]);
        let _549_v86;
        _549_v86 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,_456_v7);
        let _550_v87;
        _550_v87 = _dafny.Map.Empty.slice().updateUnsafe(_549_v86,true);
        if (((((_548_v85).contains(_545_v82)) ? ((_548_v85).get(_545_v82)) : (_456_v7))).isLessThan((_dafny.ZERO).minus(new BigNumber((_550_v87).length)))) {
          let _551_v88;
          _551_v88 = _dafny.Set.fromElements(_module.__default.safeModuloInt(_456_v7, _456_v7), _456_v7, _module.__default.fm0((_452_v3).f20, (_452_v3).f20, globalState), _456_v7);
          _551_v88 = (function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_dafny.MultiSet.fromElements(new BigNumber((_this.f22).length))).Elements) {
              let _552_v89 = _compr_11;
              if ((_dafny.MultiSet.fromElements(new BigNumber((_this.f22).length))).contains(_552_v89)) {
                _coll11.add(_module.__default.safeModuloInt(_552_v89, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(false)))).cardinality())));
              }
            }
            return _coll11;
          }()).Intersect(_551_v88);
          let _553_v90;
          _553_v90 = new _dafny.CodePoint('p'.codePointAt(0));
          let _554_v91;
          _554_v91 = _dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,_553_v90);
          _554_v91 = (_554_v91).update(_448_v0, _553_v90);
          let _555_v92;
          _555_v92 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,(_this).f23);
          let _rhs48 = _555_v92;
          let _rhs49 = !(!((_452_v3).f20));
          _555_v92 = _rhs48;
          _448_v0 = _rhs49;
          let _556_v93;
          _556_v93 = _dafny.Set.fromElements(_545_v82);
          let _557_v94;
          _557_v94 = _dafny.MultiSet.fromElements(_553_v90);
          let _558_v95;
          _558_v95 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,_557_v94);
          _449_v1 = _dafny.Seq.of((_556_v93).IsProperSubsetOf(_556_v93), (_452_v3).f20, (_452_v3).f20, (_452_v3).f20, !((((_558_v95).contains(new BigNumber((_547_v84).length))) ? ((_558_v95).get(new BigNumber((_547_v84).length))) : (_557_v94))).equals(_557_v94));
          let _559_v96;
          let _init12 = ((_560_v1, _561_v3) => function (_562_i5) {
            return _dafny.Seq.of(_module.D0.create_DC0(_dafny.MultiSet.FromArray(_dafny.Seq.update(_560_v1, _module.__default.safeIndex(new BigNumber((_560_v1).length), new BigNumber((_560_v1).length)), (_561_v3).f20))), _module.D0.create_DC0(_dafny.MultiSet.fromElements((_561_v3).f20, (_561_v3).f20, (_561_v3).f20, !((_561_v3).f20), (_561_v3).f20)), _module.D0.create_DC0(_dafny.MultiSet.FromArray(_560_v1)));
          })(_449_v1, _452_v3);
          let _nw94 = Array((new BigNumber(23)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw94.length); _i0_12++) {
            _nw94[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _559_v96 = _nw94;
          let _563_v97;
          _563_v97 = _dafny.MultiSet.fromElements(_448_v0);
          let _564_v98;
          _564_v98 = _module.D0.create_DC0(_563_v97);
          let _565_v99;
          _565_v99 = _dafny.Seq.of(_module.D0.create_DC0(_563_v97), _564_v98);
          let _index71 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_559_v96).length));
          (_559_v96)[_index71] = (((_452_v3).f20) ? (_565_v99) : (_565_v99));
          let _index72 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_559_v96).length));
          (_559_v96)[_index72] = _dafny.Seq.update(_dafny.Seq.update(_565_v99, _module.__default.safeIndex((_dafny.ZERO).minus(_456_v7), new BigNumber((_565_v99).length)), _module.D0.create_DC0(_module.__default.fm20((_452_v3).f20, _456_v7, globalState))), _module.__default.safeIndex(_456_v7, new BigNumber((_dafny.Seq.update(_565_v99, _module.__default.safeIndex((_dafny.ZERO).minus(_456_v7), new BigNumber((_565_v99).length)), _module.D0.create_DC0(_module.__default.fm20((_452_v3).f20, _456_v7, globalState)))).length)), (((_452_v3).f20) ? (_564_v98) : (_module.__default.fm21(_456_v7, globalState))));
        } else {
          let _566_v100;
          _566_v100 = _dafny.MultiSet.fromElements(_this.f22, _this.f22);
          let _567_v102;
          let _nw95 = new _module.C1();
          _nw95.__ctor(_this.f22, (_452_v3).f19, _448_v0);
          _567_v102 = _nw95;
          let _568_v103;
          _568_v103 = _dafny.MultiSet.fromElements(_567_v102);
          let _569_v104;
          _569_v104 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22((((_568_v103).contains(_567_v102)) ? ((_568_v103).get(_567_v102)) : (_456_v7)), _456_v7, globalState),new BigNumber(((_567_v102).f27).length));
          let _570_v105;
          _570_v105 = _dafny.MultiSet.fromElements(_448_v0, (_452_v3).f20, (_452_v3).f20, (_452_v3).f20, (_452_v3).f20);
          let _571_v106;
          _571_v106 = new _dafny.CodePoint('d'.codePointAt(0));
          let _572_v107;
          _572_v107 = _dafny.Map.Empty.slice().updateUnsafe(_571_v106,_448_v0);
          let _573_v108;
          _573_v108 = _dafny.Set.fromElements(new BigNumber((_572_v107).length));
          let _574_v109;
          _574_v109 = _dafny.Map.Empty.slice().updateUnsafe(_570_v105,new BigNumber((_573_v108).length));
          let _575_v110;
          let _nw96 = Array((new BigNumber(28)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = new BigNumber((_this.f22).length);
          _nw96[(_dafny.ONE).toNumber()] = _456_v7;
          _nw96[(new BigNumber(2)).toNumber()] = (((_452_v3).f20) ? (_456_v7) : ((_dafny.ZERO).minus(_456_v7)));
          _nw96[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_456_v7), new BigNumber(645));
          _nw96[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_456_v7);
          _nw96[(new BigNumber(5)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(6)).toNumber()] = new BigNumber(-369);
          _nw96[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("brlgftm"),false)).length);
          _nw96[(new BigNumber(8)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(9)).toNumber()] = new BigNumber(122);
          _nw96[(new BigNumber(10)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(11)).toNumber()] = (_456_v7).multipliedBy((_dafny.ZERO).minus(_456_v7));
          _nw96[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_456_v7), _456_v7);
          _nw96[(new BigNumber(13)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(14)).toNumber()] = (_456_v7).multipliedBy(new BigNumber((function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_566_v100).Elements) {
              let _576_v101 = _compr_12;
              if ((_566_v100).contains(_576_v101)) {
                _coll12.add(_576_v101);
              }
            }
            return _coll12;
          }()).length));
          _nw96[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_452_v3).f20) ? (_456_v7) : (_456_v7))));
          _nw96[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_this.f22).length), new BigNumber((_460_v11).length));
          _nw96[(new BigNumber(17)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(18)).toNumber()] = new BigNumber((_566_v100).cardinality());
          _nw96[(new BigNumber(19)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(20)).toNumber()] = (new BigNumber(942)).minus((_dafny.ZERO).minus(_456_v7));
          _nw96[(new BigNumber(21)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(406), _456_v7);
          _nw96[(new BigNumber(23)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(24)).toNumber()] = (((_569_v104).contains(_574_v109)) ? ((_569_v104).get(_574_v109)) : (_456_v7));
          _nw96[(new BigNumber(25)).toNumber()] = (_456_v7).minus(_456_v7);
          _nw96[(new BigNumber(26)).toNumber()] = _456_v7;
          _nw96[(new BigNumber(27)).toNumber()] = (_456_v7).multipliedBy(_456_v7);
          _575_v110 = _nw96;
          let _index73 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length));
          (_575_v110)[_index73] = _456_v7;
          let _index74 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length));
          (_575_v110)[_index74] = _456_v7;
          let _577_v111;
          let _nw97 = new _module.C1();
          _nw97.__ctor((_567_v102).f27, (_this).f23, (_452_v3).f20);
          _577_v111 = _nw97;
          let _578_v112;
          _578_v112 = _dafny.Map.Empty.slice().updateUnsafe(((_575_v110)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length))]).plus((_575_v110)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length))]),false);
          _578_v112 = (_578_v112).update((_456_v7).multipliedBy((_575_v110)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length))]), (_452_v3).f20);
          _545_v82 = ((_545_v82).Intersect(_545_v82)).Intersect(_545_v82);
          _457_v8 = (_457_v8).update((_452_v3).f20, (_575_v110)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_575_v110).length))]);
        }
        let _579_v113;
        _579_v113 = _module.D4.create_DC14((_452_v3).f20, _545_v82);
        let _580_v114;
        _580_v114 = _dafny.Seq.of(_dafny.Set.fromElements(_448_v0, (_579_v113).dtor_cf25));
        let _581_v115;
        _581_v115 = new _dafny.CodePoint('y'.codePointAt(0));
        let _582_v116;
        _582_v116 = _dafny.Map.Empty.slice().updateUnsafe(_581_v115,(_452_v3).f20);
        let _583_v117;
        _583_v117 = _dafny.Map.Empty.slice().updateUnsafe((_452_v3).f20,(_452_v3).f20);
        let _584_v118;
        _584_v118 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(917),(_452_v3).f20);
        let _585_v119;
        _585_v119 = _dafny.Map.Empty.slice().updateUnsafe(_456_v7,_dafny.Set.fromElements(!((((_584_v118).contains(new BigNumber(-341))) ? ((_584_v118).get(new BigNumber(-341))) : ((_449_v1)[_module.__default.safeIndex(_456_v7, new BigNumber((_449_v1).length))])))));
        let _586_v120;
        let _nw98 = Array((new BigNumber(19)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = (_580_v114)[_module.__default.safeIndex((_dafny.ZERO).minus(_456_v7), new BigNumber((_580_v114).length))];
        _nw98[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements((_452_v3).f20, (((_582_v116).contains(new _dafny.CodePoint('w'.codePointAt(0)))) ? ((_582_v116).get(new _dafny.CodePoint('w'.codePointAt(0)))) : ((_452_v3).f20)));
        _nw98[(new BigNumber(2)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(3)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(4)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(5)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(6)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(7)).toNumber()] = _module.__default.fm17(new BigNumber((_583_v117).length), globalState);
        _nw98[(new BigNumber(8)).toNumber()] = (_545_v82).Difference(_545_v82);
        _nw98[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_452_v3).f20, (_452_v3).f20, (_452_v3).f20);
        _nw98[(new BigNumber(10)).toNumber()] = (_545_v82).Union((((_585_v119).contains(new BigNumber((_this.f22).length))) ? ((_585_v119).get(new BigNumber((_this.f22).length))) : (_545_v82)));
        _nw98[(new BigNumber(11)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(12)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(13)).toNumber()] = (_module.__default.fm17(_456_v7, globalState)).Difference(_dafny.Set.fromElements((_452_v3).f20, (_452_v3).f20, (_452_v3).f20, (_452_v3).f20, (_452_v3).f20));
        _nw98[(new BigNumber(14)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(15)).toNumber()] = (_580_v114)[_module.__default.safeIndex(_456_v7, new BigNumber((_580_v114).length))];
        _nw98[(new BigNumber(16)).toNumber()] = _545_v82;
        _nw98[(new BigNumber(17)).toNumber()] = (_545_v82).Difference(_545_v82);
        _nw98[(new BigNumber(18)).toNumber()] = (_545_v82).Difference(_545_v82);
        _586_v120 = _nw98;
        let _index75 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_586_v120).length));
        (_586_v120)[_index75] = _545_v82;
        let _index76 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_586_v120).length));
        (_586_v120)[_index76] = ((_545_v82).Union(_545_v82)).Intersect(_dafny.Set.fromElements((_452_v3).f20, _448_v0));
        let _587_v121;
        let _init13 = ((_588_v7, _589_v118, _590_v3, _591_v84, _592_v117) => function (_593_i6) {
          return _dafny.Seq.of(_module.D3.create_DC11(_588_v7, new BigNumber((_589_v118).length), (_590_v3).f20), _module.D3.create_DC11(new BigNumber((_591_v84).length), new BigNumber((_592_v117).length), (_590_v3).f20), _module.D3.create_DC11(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-641)), ((_594_v117) => function (_595_i7) {
  return new BigNumber((_594_v117).length);
})(_592_v117))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), function (_596_i8) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length), (_590_v3).f20), _module.D3.create_DC11(_588_v7, _588_v7, true), _module.D3.create_DC11(_588_v7, _588_v7, (_590_v3).f20));
        })(_456_v7, _584_v118, _452_v3, _547_v84, _583_v117);
        let _nw99 = Array((new BigNumber(24)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw99.length); _i0_13++) {
          _nw99[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _587_v121 = _nw99;
        let _597_v122;
        _597_v122 = _module.D5.create_DC19(_456_v7, (_452_v3).f20, _587_v121, _448_v0, _581_v115);
        _583_v117 = (_583_v117).update((_597_v122).dtor_cf38, !(true));
        let _rhs50 = _456_v7;
        let _rhs51 = (_449_v1)[_module.__default.safeIndex(new BigNumber((_this.f22).length), new BigNumber((_449_v1).length))];
        let _lhs43 = globalState;
        let _lhs44 = globalState;
        _lhs43.f18 = _rhs50;
        _lhs44.f0 = _rhs51;
        _456_v7 = _456_v7;
      }
      return;
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f19 = [];
      this._f20 = false;
      this._f21 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    __ctor(f21, f19, f20) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      if ((_this).f20) {
        return _dafny.MultiSet.fromElements((_this).f20);
      } else {
        return ((_module.D0.create_DC0(_dafny.MultiSet.fromElements((_this).f20))).dtor_cf0).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f20)));
      }
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let _598_v0;
      _598_v0 = _dafny.Seq.UnicodeFromString("jrrji");
      (globalState).f13 = new BigNumber((_dafny.Seq.Concat(_598_v0, _dafny.Seq.UnicodeFromString("qqvxm"))).length);
      _598_v0 = _598_v0;
      (globalState).f0 = !((_this).f20);
      let _599_v1;
      _599_v1 = _dafny.MultiSet.fromElements((_this).f20, p0);
      let _600_v2;
      _600_v2 = _module.D0.create_DC0(_599_v1);
      let _pat_let_tv24 = p0;
      let _pat_let_tv25 = p0;
      let _pat_let_tv26 = p0;
      if (function (_source8) {
        if (_source8.is_DC1) {
          let _601___mcc_h0 = (_source8).cf1;
          let _602___mcc_h1 = (_source8).cf2;
          let _603___mcc_h2 = (_source8).cf3;
          let _604_cf3 = _603___mcc_h2;
          let _605_cf2 = _602___mcc_h1;
          let _606_cf1 = _601___mcc_h0;
          return _pat_let_tv24;
        } else if (_source8.is_DC2) {
          let _607___mcc_h3 = (_source8).cf4;
          let _608___mcc_h4 = (_source8).cf5;
          let _609___mcc_h5 = (_source8).cf6;
          let _610___mcc_h6 = (_source8).cf7;
          let _611___mcc_h7 = (_source8).cf8;
          let _612_cf8 = _611___mcc_h7;
          let _613_cf7 = _610___mcc_h6;
          let _614_cf6 = _609___mcc_h5;
          let _615_cf5 = _608___mcc_h4;
          let _616_cf4 = _607___mcc_h3;
          return _616_cf4;
        } else if (_source8.is_DC3) {
          let _617___mcc_h8 = (_source8).cf9;
          let _618_cf9 = _617___mcc_h8;
          return _pat_let_tv25;
        } else {
          let _619___mcc_h9 = (_source8).cf0;
          let _620_cf0 = _619___mcc_h9;
          return _pat_let_tv26;
        }
      }(_600_v2)) {
        (globalState).f0 = p0;
        (globalState).f0 = !_dafny.areEqual(_598_v0, _dafny.Seq.UnicodeFromString("sqfpsses"));
        let _621_v3;
        _621_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
        let _622_v4;
        _622_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_598_v0, _598_v0),_621_v3);
        _622_v4 = _622_v4;
        if (p0) {
          (globalState).f13 = (_this).f21;
          let _623_v5;
          _623_v5 = _dafny.Seq.of(p0);
          let _624_v6;
          _624_v6 = _module.D0.create_DC2((_this).f20, _module.__default.fm0((_this).f20, true, globalState), _623_v5, (_this).f21, (_this).f21);
          (globalState).f13 = (_624_v6).dtor_cf5;
          (globalState).f0 = !((true) && (((_this).f21).isLessThan((_this).f21)));
          (globalState).f0 = p0;
          let _625_v7;
          _625_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_598_v0).length)),_dafny.MultiSet.fromElements((_this).f20, p0, p0, p0));
          let _626_v8;
          _626_v8 = _dafny.Seq.of(_599_v1, _599_v1, _599_v1, _599_v1, _599_v1);
          let _627_v9;
          _627_v9 = _dafny.Seq.of((_this).f21);
          let _628_v10;
          let _nw100 = Array((new BigNumber(26)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(p0, (_this).f20, p0);
          _nw100[(_dafny.ONE).toNumber()] = ((((_625_v7).contains((_this).f21)) ? ((_625_v7).get((_this).f21)) : (_599_v1))).update(!((_this).f20), _module.__default.abs(new BigNumber(138)));
          _nw100[(new BigNumber(2)).toNumber()] = (_599_v1).Difference(_599_v1);
          _nw100[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.FromArray(_623_v5)).update(p0, _module.__default.abs((_this).f21));
          _nw100[(new BigNumber(4)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(5)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_623_v5);
          _nw100[(new BigNumber(7)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(8)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(9)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(10)).toNumber()] = ((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs((_this).f21))).Difference(_599_v1);
          _nw100[(new BigNumber(11)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(12)).toNumber()] = (((_this).f20) ? (_599_v1) : (_599_v1));
          _nw100[(new BigNumber(13)).toNumber()] = (_626_v8)[_module.__default.safeIndex(new BigNumber((_627_v9).length), new BigNumber((_626_v8).length))];
          _nw100[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.fromElements(_module.__default.fm2((_this).f20, (_this).f21, (_this).f21, globalState))).Union(_599_v1);
          _nw100[(new BigNumber(15)).toNumber()] = (_this).fm1((_this).f21, (_this).f21, (_this).f20, globalState);
          _nw100[(new BigNumber(16)).toNumber()] = ((_this).fm1((_this).f21, new BigNumber((p1).length), true, globalState)).Intersect(_dafny.MultiSet.fromElements((_this).f20, false));
          _nw100[(new BigNumber(17)).toNumber()] = (_600_v2).dtor_cf0;
          _nw100[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_623_v5);
          _nw100[(new BigNumber(19)).toNumber()] = (_this).fm1((_this).f21, (_this).f21, p0, globalState);
          _nw100[(new BigNumber(20)).toNumber()] = (_this).fm1((_this).f21, (_this).f21, (_this).f20, globalState);
          _nw100[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.fromElements((_this).f20, (_this).f20, !(true), false, p0);
          _nw100[(new BigNumber(22)).toNumber()] = (_dafny.MultiSet.fromElements(false, (_this).f20)).Union(_599_v1);
          _nw100[(new BigNumber(23)).toNumber()] = (_626_v8)[_module.__default.safeIndex(new BigNumber((_599_v1).cardinality()), new BigNumber((_626_v8).length))];
          _nw100[(new BigNumber(24)).toNumber()] = _599_v1;
          _nw100[(new BigNumber(25)).toNumber()] = (_600_v2).dtor_cf0;
          _628_v10 = _nw100;
          let _index77 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_628_v10).length));
          (_628_v10)[_index77] = (_this).fm1((_dafny.ZERO).minus((_this).f21), (_this).f21, _module.__default.fm2(p0, (_this).f21, (_this).f21, globalState), globalState);
          let _index78 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_628_v10).length));
          (_628_v10)[_index78] = _599_v1;
        } else {
          let _629_v11;
          _629_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _index79 = _module.__default.safeIndex(new BigNumber(883), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index79] = p0;
          let _index80 = _module.__default.safeIndex(new BigNumber(883), new BigNumber(((_this).f19).length));
          let _rhs52 = _629_v11;
          let _rhs53 = _dafny.areEqual(_598_v0, _dafny.Seq.Concat(_598_v0, _598_v0));
          let _rhs54 = true;
          let _rhs55 = !((_this).f20) || (!((_this).f20));
          let _lhs45 = globalState;
          let _lhs46 = globalState;
          let _lhs47 = (_this).f19;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(883), new BigNumber(((_this).f19).length));
          _629_v11 = _rhs52;
          _lhs45.f0 = _rhs53;
          _lhs46.f0 = _rhs54;
          _lhs47[_lhs48] = _rhs55;
          let _630_v12;
          _630_v12 = _dafny.Seq.of((_this).f21);
          let _631_v13;
          _631_v13 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f19)[_module.__default.safeIndex(new BigNumber(883), new BigNumber(((_this).f19).length))],_630_v12);
          let _632_v14;
          _632_v14 = _dafny.Set.fromElements(_dafny.Seq.Concat((((_631_v13).contains(p0)) ? ((_631_v13).get(p0)) : (_dafny.Seq.of((_this).f21, (_this).f21))), _630_v12), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), function (_633_i0) {
            return new BigNumber(430);
          }), _module.__default.safeIndex((_this).f21, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), function (_634_i0) {
            return new BigNumber(430);
          })).length)), (_this).f21), _630_v12));
          let _rhs56 = (((_dafny.ZERO).minus((_this).f21)).minus((_this).f21)).isLessThan(new BigNumber(((_629_v11).Merge(_629_v11)).length));
          let _rhs57 = _dafny.Set.fromElements(_module.__default.fm3(globalState));
          let _rhs58 = _module.__default.safeDivisionInt((_this).f21, new BigNumber((_621_v3).length));
          let _lhs49 = globalState;
          let _lhs50 = globalState;
          _lhs49.f0 = _rhs56;
          _632_v14 = _rhs57;
          _lhs50.f13 = _rhs58;
          let _635_v15;
          let _nw101 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _635_v15 = _nw101;
          let _index81 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_635_v15).length));
          (_635_v15)[_index81] = new BigNumber(789);
          let _index82 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_635_v15).length));
          let _rhs59 = p0;
          let _rhs60 = (_this).f21;
          let _lhs51 = globalState;
          let _lhs52 = _635_v15;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_635_v15).length));
          _lhs51.f0 = _rhs59;
          _lhs52[_lhs53] = _rhs60;
          (globalState).f0 = (_this).f20;
          _598_v0 = _598_v0;
        }
        let _636_v16;
        _636_v16 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false);
        let _637_v17;
        _637_v17 = new _dafny.CodePoint('j'.codePointAt(0));
        (globalState).f0 = (((_636_v16).contains(_637_v17)) ? ((_636_v16).get(_637_v17)) : ((_this).f20));
      } else {
        let _638_v18;
        _638_v18 = new _dafny.CodePoint('x'.codePointAt(0));
        let _639_v19;
        _639_v19 = _dafny.Map.Empty.slice().updateUnsafe(_638_v18,p0);
        let _640_v20;
        _640_v20 = _module.D0.create_DC1(p0, (_dafny.ZERO).minus(new BigNumber((_599_v1).cardinality())), new BigNumber(360));
        let _641_v21;
        _641_v21 = _dafny.MultiSet.fromElements(_640_v20);
        let _642_v22;
        _642_v22 = _dafny.Seq.of(p0, (_this).f20);
        let _643_v23;
        _643_v23 = _module.D0.create_DC2((((_639_v19).contains(_module.__default.fm4(globalState))) ? ((_639_v19).get(_module.__default.fm4(globalState))) : ((_this).f20)), new BigNumber((_641_v21).cardinality()), _642_v22, new BigNumber((_598_v0).length), (_this).f21);
        _643_v23 = _643_v23;
        (globalState).f0 = !(!((_this).f20)) || ((((_this).f20) ? ((_this).f20) : ((_this).f20)));
        let _644_v24;
        let _nw102 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _644_v24 = _nw102;
        (globalState).f3 = _644_v24;
        if ((_this).f20) {
          let _index83 = _module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length));
          ((_this).f19)[_index83] = (true) === (!(p0));
          let _645_v25;
          _645_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f21);
          let _index84 = _module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length));
          let _rhs61 = p0;
          let _rhs62 = !(_645_v25).equals(_645_v25);
          let _rhs63 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f20), _642_v22);
          let _lhs54 = globalState;
          let _lhs55 = (_this).f19;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length));
          _lhs54.f0 = _rhs61;
          _lhs55[_lhs56] = _rhs62;
          _642_v22 = _rhs63;
          let _646_v26;
          let _nw103 = Array((new BigNumber(23)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(_dafny.ONE).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(2)).toNumber()] = false;
          _nw103[(new BigNumber(3)).toNumber()] = true;
          _nw103[(new BigNumber(4)).toNumber()] = false;
          _nw103[(new BigNumber(5)).toNumber()] = p0;
          _nw103[(new BigNumber(6)).toNumber()] = false;
          _nw103[(new BigNumber(7)).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(8)).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(9)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(new BigNumber(10)).toNumber()] = p0;
          _nw103[(new BigNumber(11)).toNumber()] = p0;
          _nw103[(new BigNumber(12)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(new BigNumber(13)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(new BigNumber(14)).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(15)).toNumber()] = true;
          _nw103[(new BigNumber(16)).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(17)).toNumber()] = (_this).f20;
          _nw103[(new BigNumber(18)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(new BigNumber(19)).toNumber()] = p0;
          _nw103[(new BigNumber(20)).toNumber()] = ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))];
          _nw103[(new BigNumber(21)).toNumber()] = p0;
          _nw103[(new BigNumber(22)).toNumber()] = (_this).f20;
          _646_v26 = _nw103;
          let _647_v27;
          let _nw104 = new _module.C1();
          _nw104.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ucn"), _598_v0), _646_v26, false);
          _647_v27 = _nw104;
          let _648_v28;
          _648_v28 = _dafny.Map.Empty.slice().updateUnsafe(_644_v24,new _dafny.CodePoint('e'.codePointAt(0)));
          _648_v28 = (_648_v28).update(_644_v24, new _dafny.CodePoint('g'.codePointAt(0)));
          let _649_v29;
          _649_v29 = _module.D4.create_DC13();
          let _650_v30;
          _650_v30 = _module.D4.create_DC15(_649_v29);
          _650_v30 = _650_v30;
          let _651_v31;
          let _nw105 = new _module.C0();
          _nw105.__ctor((_this).f21);
          _651_v31 = _nw105;
          let _652_v32;
          let _nw106 = new _module.C2();
          _nw106.__ctor(_651_v31, ((_this).f19)[_module.__default.safeIndex(new BigNumber(410), new BigNumber(((_this).f19).length))], _646_v26, (_this).f20);
          _652_v32 = _nw106;
        } else {
          let _653_v33;
          _653_v33 = _dafny.MultiSet.fromElements((_this).f21);
          let _654_v34;
          _654_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_653_v33).update((_this).f21, _module.__default.abs((_this).f21))).cardinality())),(_this).f20);
          let _655_v35;
          _655_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,new BigNumber((_module.__default.fm8(_598_v0, _654_v34, (_this).f21, globalState)).length));
          let _656_v36;
          let _nw107 = new _module.C0();
          _nw107.__ctor((_this).f21);
          _656_v36 = _nw107;
          let _657_v37;
          let _nw108 = new _module.C2();
          _nw108.__ctor(_656_v36, (_this).f20, (_this).f19, p0);
          _657_v37 = _nw108;
          let _658_v38;
          _658_v38 = _dafny.Seq.of(_657_v37, _657_v37, _657_v37, _657_v37, _657_v37);
          let _659_v39;
          _659_v39 = _dafny.MultiSet.fromElements(new BigNumber((_642_v22).length), (_this).f21, new BigNumber((_655_v35).length), new BigNumber((_658_v38).length));
          (globalState).f13 = (new BigNumber((_659_v39).cardinality())).minus((_656_v36).f24);
          let _660_v40;
          _660_v40 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(126));
          _660_v40 = (_660_v40).update((_this).f20, (_656_v36).f24);
          let _index85 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length));
          (_644_v24)[_index85] = _module.__default.safeModuloInt((_this).f21, (_this).f21);
          let _index86 = _module.__default.safeIndex(new BigNumber(627), new BigNumber(((_657_v37).f19).length));
          ((_657_v37).f19)[_index86] = p0;
          let _661_v41;
          _661_v41 = _dafny.Map.Empty.slice().updateUnsafe((_657_v37).f20,_598_v0);
          let _index87 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(627), new BigNumber(((_657_v37).f19).length));
          let _rhs64 = new BigNumber(948);
          let _rhs65 = new BigNumber(542);
          let _rhs66 = (_this).f21;
          let _rhs67 = _module.__default.fm7(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((((_661_v41).contains(_module.__default.fm7((_this).f21, (_this).f21, (_this).f21, globalState))) ? ((_661_v41).get(_module.__default.fm7((_this).f21, (_this).f21, (_this).f21, globalState))) : (_dafny.Seq.UnicodeFromString("bxledcrb")))).length)), (((_655_v35).contains((_this).f21)) ? ((_655_v35).get((_this).f21)) : ((_this).f21))), (_656_v36).f24, (_this).f21, globalState);
          let _rhs68 = (_dafny.ZERO).minus((_this).f21);
          let _lhs57 = globalState;
          let _lhs58 = globalState;
          let _lhs59 = _644_v24;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length));
          let _lhs61 = (_657_v37).f19;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(627), new BigNumber(((_657_v37).f19).length));
          let _lhs63 = globalState;
          _lhs57.f18 = _rhs64;
          _lhs58.f13 = _rhs65;
          _lhs59[_lhs60] = _rhs66;
          _lhs61[_lhs62] = _rhs67;
          _lhs63.f18 = _rhs68;
          let _662_v42;
          _662_v42 = _module.D3.create_DC9(_598_v0);
          let _663_v43;
          let _nw109 = Array((new BigNumber(17)).toNumber());
          _nw109[(_dafny.ZERO).toNumber()] = _module.__default.fm0(((_657_v37).f19)[_module.__default.safeIndex(new BigNumber(627), new BigNumber(((_657_v37).f19).length))], (_this).f20, globalState);
          _nw109[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("xw")).length);
          _nw109[(new BigNumber(2)).toNumber()] = new BigNumber(836);
          _nw109[(new BigNumber(3)).toNumber()] = new BigNumber(350);
          _nw109[(new BigNumber(4)).toNumber()] = (_this).f21;
          _nw109[(new BigNumber(5)).toNumber()] = new BigNumber((_598_v0).length);
          _nw109[(new BigNumber(6)).toNumber()] = (_644_v24)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length))];
          _nw109[(new BigNumber(7)).toNumber()] = (_this).f21;
          _nw109[(new BigNumber(8)).toNumber()] = (_this).f21;
          _nw109[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(233)), ((_664_v18) => function (_665_i1) {
            return _664_v18;
          })(_638_v18))).length)));
          _nw109[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("yqgtfnndn")).length);
          _nw109[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt((_656_v36).f24, (_656_v36).f24);
          _nw109[(new BigNumber(12)).toNumber()] = (new BigNumber(((_662_v42).dtor_cf16).length)).plus(new BigNumber((p1).length));
          _nw109[(new BigNumber(13)).toNumber()] = (_644_v24)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length))];
          _nw109[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_this).f21);
          _nw109[(new BigNumber(15)).toNumber()] = (_656_v36).f24;
          _nw109[(new BigNumber(16)).toNumber()] = ((_this).f21).minus((_644_v24)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_644_v24).length))]);
          _663_v43 = _nw109;
          (globalState).f10 = _663_v43;
          let _index89 = _module.__default.safeIndex(new BigNumber(627), new BigNumber(((_657_v37).f19).length));
          ((_657_v37).f19)[_index89] = !((p0) === (_module.__default.fm7((_this).f21, new BigNumber((_dafny.Set.fromElements((_657_v37).f20, (_657_v37).f20, false)).length), (_656_v36).f24, globalState)));
        }
        let _666_v44;
        _666_v44 = _module.D3.create_DC9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), ((_667_v18) => function (_668_i2) {
  return _667_v18;
})(_638_v18)));
        let _pat_let_tv27 = _666_v44;
        let _pat_let_tv28 = _598_v0;
        let _pat_let_tv29 = _598_v0;
        let _pat_let_tv30 = _638_v18;
        let _669_v45;
        let _nw110 = Array((new BigNumber(22)).toNumber());
        _nw110[(_dafny.ZERO).toNumber()] = _666_v44;
        _nw110[(_dafny.ONE).toNumber()] = _666_v44;
        _nw110[(new BigNumber(2)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(3)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(4)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(5)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(6)).toNumber()] = _module.D3.create_DC9(_598_v0);
        _nw110[(new BigNumber(7)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(8)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(9)).toNumber()] = function (_pat_let20_0) {
          return function (_670_dt__update__tmp_h0) {
            return function (_pat_let21_0) {
              return function (_671_dt__update_hcf16_h0) {
                return _module.D3.create_DC9(_671_dt__update_hcf16_h0);
              }(_pat_let21_0);
            }((_pat_let_tv27).dtor_cf16);
          }(_pat_let20_0);
        }(_666_v44);
        _nw110[(new BigNumber(10)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(11)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(12)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(13)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(14)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(15)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(16)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(17)).toNumber()] = ((p0) ? (_666_v44) : (_666_v44));
        _nw110[(new BigNumber(18)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(19)).toNumber()] = _666_v44;
        _nw110[(new BigNumber(20)).toNumber()] = function (_pat_let22_0) {
          return function (_672_dt__update__tmp_h1) {
            return function (_pat_let23_0) {
              return function (_673_dt__update_hcf16_h1) {
                return _module.D3.create_DC9(_673_dt__update_hcf16_h1);
              }(_pat_let23_0);
            }(_dafny.Seq.update(_pat_let_tv28, _module.__default.safeIndex((_this).f21, new BigNumber((_pat_let_tv29).length)), _pat_let_tv30));
          }(_pat_let22_0);
        }(_666_v44);
        _nw110[(new BigNumber(21)).toNumber()] = _666_v44;
        _669_v45 = _nw110;
        _669_v45 = _669_v45;
      }
      (globalState).f18 = (_this).f21;
      let _674_v46;
      let _nw111 = new _module.C0();
      _nw111.__ctor(new BigNumber(107));
      _674_v46 = _nw111;
      let _675_v47;
      let _nw112 = new _module.C2();
      _nw112.__ctor(_674_v46, (_this).f20, (_this).f19, (_this).f20);
      _675_v47 = _nw112;
      let _676_v48;
      _676_v48 = _dafny.Map.Empty.slice().updateUnsafe(_675_v47,(_this).f20);
      let _677_v49;
      _677_v49 = _dafny.Map.Empty.slice().updateUnsafe((((_676_v48).contains(_675_v47)) ? ((_676_v48).get(_675_v47)) : (!(p0))),((_674_v46).f24).isEqualTo((_674_v46).f24));
      let _678_v50;
      _678_v50 = _module.D4.create_DC14(p0, p1);
      let _pat_let_tv31 = p1;
      _677_v49 = (_677_v49).update((function (_pat_let24_0) {
        return function (_679_dt__update__tmp_h2) {
          return function (_pat_let25_0) {
            return function (_680_dt__update_hcf26_h0) {
              return _module.D4.create_DC14((_679_dt__update__tmp_h2).dtor_cf25, _680_dt__update_hcf26_h0);
            }(_pat_let25_0);
          }(_pat_let_tv31);
        }(_pat_let24_0);
      }(_678_v50)).dtor_cf25, p0);
      return;
    }
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f13 = ((_this).f21).multipliedBy((_this).f21);
      let _hi5 = (_this).f21;
      for (let _681_i0 = (_this).f21; _681_i0.isLessThan(_hi5); _681_i0 = _681_i0.plus(_dafny.ONE)) {
        let _682_v0;
        _682_v0 = _dafny.Seq.UnicodeFromString("v");
        let _683_v1;
        let _nw113 = new _module.C1();
        _nw113.__ctor(_682_v0, (_this).f19, (_this).f20);
        _683_v1 = _nw113;
        let _684_v2;
        let _nw114 = new _module.C3();
        _nw114.__ctor(_dafny.Seq.Concat((_683_v1).f27, _682_v0), (_this).f19);
        _684_v2 = _nw114;
        let _685_v3;
        _685_v3 = _module.D2.create_DC8((((_this).f20) ? ((_this).f21) : (_681_i0)), (_this).f21);
        let _686_v4;
        let _nw115 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _686_v4 = _nw115;
        let _index90 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length));
        (_686_v4)[_index90] = _681_i0;
        let _index91 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_686_v4).length));
        (_686_v4)[_index91] = _681_i0;
        let _687_v5;
        _687_v5 = _dafny.Seq.of((_this).f20, (_this).f20);
        let _688_v6;
        _688_v6 = _dafny.MultiSet.fromElements((_this).f21);
        let _pat_let_tv32 = _688_v6;
        let _index92 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length));
        let _index93 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_686_v4).length));
        let _rhs69 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_687_v5, _687_v5)).length));
        let _rhs70 = function (_pat_let26_0) {
          return function (_689_dt__update__tmp_h0) {
            return function (_pat_let27_0) {
              return function (_690_dt__update_hcf15_h0) {
                return _module.D2.create_DC8((_689_dt__update__tmp_h0).dtor_cf14, _690_dt__update_hcf15_h0);
              }(_pat_let27_0);
            }(new BigNumber(((_pat_let_tv32).Union(_dafny.MultiSet.fromElements(new BigNumber(28)))).cardinality()));
          }(_pat_let26_0);
        }(_module.__default.fm23(globalState));
        let _rhs71 = new BigNumber((_684_v2.f22).length);
        let _rhs72 = _module.__default.safeModuloInt(_681_i0, (_this).f21);
        let _rhs73 = (_this).f20;
        let _lhs64 = globalState;
        let _lhs65 = _686_v4;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length));
        let _lhs67 = _686_v4;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_686_v4).length));
        let _lhs69 = globalState;
        _lhs64.f13 = _rhs69;
        _685_v3 = _rhs70;
        _lhs65[_lhs66] = _rhs71;
        _lhs67[_lhs68] = _rhs72;
        _lhs69.f0 = _rhs73;
        if ((_this).f20) {
          let _691_v7;
          _691_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20);
          let _692_v8;
          _692_v8 = _dafny.Seq.of((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))], (_this).f21, (_this).f21, (_this).f21, (_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]);
          let _693_v9;
          let _nw116 = Array((new BigNumber(28)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = (_this).f20;
          _nw116[(_dafny.ONE).toNumber()] = !((((_691_v7).contains((_this).f20)) ? ((_691_v7).get((_this).f20)) : ((_this).f20)));
          _nw116[(new BigNumber(2)).toNumber()] = !_dafny.areEqual(_684_v2.f22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), function (_694_i1) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }));
          _nw116[(new BigNumber(3)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(4)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(5)).toNumber()] = !(true);
          _nw116[(new BigNumber(6)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(7)).toNumber()] = !((_dafny.MultiSet.FromArray(_692_v8)).IsSubsetOf(_dafny.MultiSet.fromElements((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))])));
          _nw116[(new BigNumber(8)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(9)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(10)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(11)).toNumber()] = false;
          _nw116[(new BigNumber(12)).toNumber()] = (_683_v1).fm13(globalState);
          _nw116[(new BigNumber(13)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(14)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(15)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(16)).toNumber()] = (_688_v6).IsDisjointFrom(_688_v6);
          _nw116[(new BigNumber(17)).toNumber()] = (_681_i0).isLessThan((_this).f21);
          _nw116[(new BigNumber(18)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(19)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(20)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(21)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(22)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(23)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(24)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(25)).toNumber()] = !(true);
          _nw116[(new BigNumber(26)).toNumber()] = (_this).f20;
          _nw116[(new BigNumber(27)).toNumber()] = (_this).f20;
          _693_v9 = _nw116;
          _693_v9 = (_this).f19;
          let _695_v10;
          _695_v10 = new _dafny.CodePoint('q'.codePointAt(0));
          let _696_v12;
          _696_v12 = _dafny.Seq.of(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_692_v8).Elements) {
              let _697_v11 = _compr_13;
              if (_dafny.Seq.contains(_692_v8, _697_v11)) {
                _coll13.push([_module.__default.safeModuloInt(_697_v11, (_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]),(_this).f20]);
              }
            }
            return _coll13;
          }());
          let _698_v13;
          _698_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f20);
          let _699_v14;
          let _nw117 = new _module.C3();
          _nw117.__ctor(_module.__default.fm8(_module.__default.fm8(_dafny.Seq.of(_695_v10), (_696_v12)[_module.__default.safeIndex(_681_i0, new BigNumber((_696_v12).length))], new BigNumber((_684_v2.f22).length), globalState), _698_v13, (_this).f21, globalState), _693_v9);
          _699_v14 = _nw117;
          let _700_v15;
          _700_v15 = _dafny.Set.fromElements((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]);
          let _701_v16;
          _701_v16 = _dafny.Map.Empty.slice().updateUnsafe(_695_v10,(_700_v15).IsDisjointFrom(_dafny.Set.fromElements((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))], new BigNumber((_684_v2.f22).length))));
          _701_v16 = _701_v16;
          let _index94 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length));
          (_686_v4)[_index94] = (_dafny.ZERO).minus(new BigNumber(-178));
          let _702_v17;
          let _nw118 = Array((new BigNumber(2)).toNumber()).fill(_module.D1.Default());
          _702_v17 = _nw118;
          let _703_v18;
          _703_v18 = _module.D1.create_DC5((_this).f20);
          let _index95 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_702_v17).length));
          (_702_v17)[_index95] = _703_v18;
          let _704_v19;
          _704_v19 = _module.D0.create_DC2((_this).f20, (_this).f21, _687_v5, _module.__default.safeDivisionInt((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))], new BigNumber(853)), (_dafny.ZERO).minus(new BigNumber((_684_v2.f22).length)));
          let _705_v20;
          _705_v20 = _dafny.Map.Empty.slice().updateUnsafe((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))],_681_i0);
          let _706_v21;
          _706_v21 = _dafny.MultiSet.fromElements(!_dafny.areEqual((_683_v1).f27, _684_v2.f22), (_687_v5)[_module.__default.safeIndex(_681_i0, new BigNumber((_687_v5).length))], !((_this).f20), (_681_i0).isEqualTo((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]), (_this).f20);
          let _707_v22;
          _707_v22 = _dafny.MultiSet.fromElements(_695_v10, _695_v10);
          let _index96 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_702_v17).length));
          let _rhs74 = _module.D1.create_DC5((_this).f20);
          let _rhs75 = _704_v19;
          let _rhs76 = (_705_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe((((_707_v22).contains(_695_v10)) ? ((_707_v22).get(_695_v10)) : (new BigNumber(817))),_681_i0));
          let _rhs77 = (_this).f20;
          let _rhs78 = (_706_v21).Intersect((_706_v21).Intersect(_706_v21));
          let _lhs70 = _702_v17;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_702_v17).length));
          let _lhs72 = globalState;
          _lhs70[_lhs71] = _rhs74;
          _704_v19 = _rhs75;
          _705_v20 = _rhs76;
          _lhs72.f0 = _rhs77;
          _706_v21 = _rhs78;
        } else {
          r0 = _681_i0;
          let _rhs79 = _687_v5;
          let _rhs80 = (_this).f20;
          let _lhs73 = globalState;
          _687_v5 = _rhs79;
          _lhs73.f0 = _rhs80;
          (_684_v2).m3(globalState);
          let _708_v23;
          _708_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("iqw"),(_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]);
          let _709_v24;
          _709_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_683_v1).f27);
          let _710_v25;
          _710_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,(_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]);
          let _711_v26;
          _711_v26 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))]), (_710_v25).update((_this).f20, _681_i0));
          let _712_v27;
          _712_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-831),(((_711_v26).contains(_710_v25)) ? ((_711_v26).get(_710_v25)) : (new BigNumber(-915))));
          _708_v23 = (_708_v23).update(_dafny.Seq.Concat(_682_v0, (((_709_v24).contains(_681_i0)) ? ((_709_v24).get(_681_i0)) : (_682_v0))), new BigNumber(((_712_v27).update((_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))], (_686_v4)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_686_v4).length))])).length));
          let _713_v28;
          _713_v28 = new _dafny.CodePoint('y'.codePointAt(0));
          _713_v28 = _713_v28;
        }
      }
      let _714_v29;
      let _nw119 = Array((new BigNumber(4)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = (_this).f21;
      _nw119[(_dafny.ONE).toNumber()] = (_this).f21;
      _nw119[(new BigNumber(2)).toNumber()] = (_this).f21;
      _nw119[(new BigNumber(3)).toNumber()] = _module.__default.fm0((_this).f20, (_this).f20, globalState);
      _714_v29 = _nw119;
      let _index97 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
      (_714_v29)[_index97] = (_this).f21;
      let _index98 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
      (_714_v29)[_index98] = (_this).f21;
      let _715_v30;
      _715_v30 = new _dafny.CodePoint('a'.codePointAt(0));
      let _716_v31;
      _716_v31 = _module.D4.create_DC12(_715_v30);
      let _717_v32;
      _717_v32 = _dafny.Seq.of(_715_v30, (_716_v31).dtor_cf24);
      let _718_v33;
      _718_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(968),false);
      let _719_v34;
      let _nw120 = new _module.C1();
      _nw120.__ctor(_dafny.Seq.Concat(_module.__default.fm8(_717_v32, _718_v33, (_this).f21, globalState), _717_v32), ((false) ? ((_this).f19) : ((_this).f19)), (_this).f20);
      _719_v34 = _nw120;
      let _720_v35;
      _720_v35 = _module.D0.create_DC1(true, (_this).f21, (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
      let _nw121 = new _module.C1();
      _nw121.__ctor(_717_v32, (((_module.D5.create_DC17((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], new BigNumber((_dafny.Seq.UnicodeFromString("qedpt")).length), _720_v35, (_this).f20, true)).dtor_cf32) ? ((_719_v34).f19) : ((_this).f19)), (_719_v34).f20);
      _719_v34 = _nw121;
      let _hi6 = (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))];
      for (let _721_i2 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-527)), ((_729_v30) => function (_730_i3) {
        return _729_v30;
      })(_715_v30))).length)).minus((_this).f21); _721_i2.isLessThan(_hi6); _721_i2 = _721_i2.plus(_dafny.ONE)) {
        let _722_v36;
        _722_v36 = _dafny.MultiSet.fromElements((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], _module.__default.fm0((_719_v34).f20, (_this).f20, globalState));
        let _723_v37;
        _723_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], _721_i2, globalState),_722_v36);
        _723_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_722_v36).Union(_722_v36));
        (globalState).f0 = (_this).f20;
        let _724_v38;
        _724_v38 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_717_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_725_v30) => function (_726_i4) {
          return _725_v30;
        })(_715_v30))), _dafny.Seq.Concat(_717_v32, _717_v32), _717_v32, _717_v32, _dafny.Seq.update((((_719_v34).f20) ? (_717_v32) : (_717_v32)), _module.__default.safeIndex((_this).f21, new BigNumber(((((_719_v34).f20) ? (_717_v32) : (_717_v32))).length)), _715_v30));
        _724_v38 = _724_v38;
        let _727_v39;
        let _nw122 = Array((new BigNumber(12)).toNumber()).fill(false);
        _727_v39 = _nw122;
        let _index99 = _module.__default.safeIndex(new BigNumber(306), new BigNumber(((_719_v34).f19).length));
        ((_719_v34).f19)[_index99] = !((_this).f20);
        let _728_v40;
        _728_v40 = _dafny.MultiSet.fromElements((_719_v34).f20, (_719_v34).f20, false);
        let _index100 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
        let _index101 = _module.__default.safeIndex(new BigNumber(306), new BigNumber(((_719_v34).f19).length));
        let _rhs81 = (_719_v34).f20;
        let _rhs82 = (_719_v34).f19;
        let _rhs83 = _module.__default.safeDivisionInt((_this).f21, (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
        let _rhs84 = (_728_v40).IsDisjointFrom((_728_v40).update((_719_v34).f20, _module.__default.abs(_721_i2)));
        let _lhs74 = globalState;
        let _lhs75 = _714_v29;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
        let _lhs77 = (_719_v34).f19;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(306), new BigNumber(((_719_v34).f19).length));
        _lhs74.f0 = _rhs81;
        _727_v39 = _rhs82;
        _lhs75[_lhs76] = _rhs83;
        _lhs77[_lhs78] = _rhs84;
      }
      let _hi7 = (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))];
      for (let _731_i5 = (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]; _731_i5.isLessThan(_hi7); _731_i5 = _731_i5.plus(_dafny.ONE)) {
        if ((((_this).f21).isLessThanOrEqualTo((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))])) === ((_this).f20)) {
          let _732_v41;
          _732_v41 = _dafny.Map.Empty.slice().updateUnsafe((_719_v34).f20,(_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
          let _733_v42;
          _733_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
          (globalState).f13 = (((_732_v41).contains(((false) ? (true) : ((_719_v34).f20)))) ? ((_732_v41).get(((false) ? (true) : ((_719_v34).f20)))) : (new BigNumber(((_733_v42).Merge(_733_v42)).length)));
          let _index102 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          (_714_v29)[_index102] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]), (_dafny.ZERO).minus((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]));
          (globalState).f0 = (_this).f20;
          let _734_v43;
          let _init14 = ((_735_v34) => function (_736_i6) {
            return (_735_v34).f20;
          })(_719_v34);
          let _nw123 = Array((new BigNumber(5)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw123.length); _i0_14++) {
            _nw123[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _734_v43 = _nw123;
          let _737_v44;
          let _nw124 = new _module.C0();
          _nw124.__ctor((_this).f21);
          _737_v44 = _nw124;
          let _738_v45;
          let _nw125 = new _module.C2();
          _nw125.__ctor(_737_v44, (_719_v34).f20, _734_v43, (_719_v34).f20);
          _738_v45 = _nw125;
          let _739_v46;
          let _nw126 = Array((new BigNumber(2)).toNumber()).fill(_module.D0.Default());
          _739_v46 = _nw126;
          let _740_v47;
          _740_v47 = _739_v46;
          let _741_v48;
          let _nw127 = Array((new BigNumber(18)).toNumber());
          _nw127[(_dafny.ZERO).toNumber()] = _740_v47;
          _nw127[(_dafny.ONE).toNumber()] = _739_v46;
          _nw127[(new BigNumber(2)).toNumber()] = _739_v46;
          _nw127[(new BigNumber(3)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(4)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(5)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(6)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(7)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(8)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(9)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(10)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(11)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(12)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(13)).toNumber()] = _739_v46;
          _nw127[(new BigNumber(14)).toNumber()] = (((_719_v34).f20) ? (_740_v47) : (_739_v46));
          _nw127[(new BigNumber(15)).toNumber()] = _739_v46;
          _nw127[(new BigNumber(16)).toNumber()] = _740_v47;
          _nw127[(new BigNumber(17)).toNumber()] = (((_719_v34).f20) ? (_739_v46) : (_739_v46));
          _741_v48 = _nw127;
          let _index103 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_741_v48).length));
          (_741_v48)[_index103] = _740_v47;
          let _index104 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          let _index105 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_741_v48).length));
          let _rhs85 = _734_v43;
          let _rhs86 = _738_v45;
          let _rhs87 = (_this).f21;
          let _rhs88 = _740_v47;
          let _lhs79 = _714_v29;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          let _lhs81 = _741_v48;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_741_v48).length));
          _734_v43 = _rhs85;
          _738_v45 = _rhs86;
          _lhs79[_lhs80] = _rhs87;
          _lhs81[_lhs82] = _rhs88;
          _732_v41 = (_732_v41).update(_dafny.Seq.IsPrefixOf(_717_v32, _dafny.Seq.UnicodeFromString("jnij")), (_737_v44).f24);
        } else {
          let _742_v49;
          _742_v49 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20);
          let _743_v50;
          _743_v50 = _module.D0.create_DC0(_742_v49);
          _743_v50 = _743_v50;
          (globalState).f0 = (_this).f20;
          let _744_v51;
          let _nw128 = Array((new BigNumber(24)).toNumber()).fill([]);
          _744_v51 = _nw128;
          _744_v51 = _744_v51;
          let _745_v52;
          _745_v52 = _dafny.Map.Empty.slice().updateUnsafe((_719_v34).f20,(_dafny.ZERO).minus((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]));
          _742_v49 = ((_742_v49).update((_719_v34).f20, _module.__default.abs(new BigNumber((_745_v52).length)))).Union(_742_v49);
          let _746_v53;
          let _nw129 = new _module.C0();
          _nw129.__ctor(new BigNumber((_module.__default.fm19(false, globalState)).length));
          _746_v53 = _nw129;
        }
        let _747_v54;
        _747_v54 = _dafny.Seq.of((_this).f20);
        let _748_v55;
        _748_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_717_v32, _717_v32),(_module.__default.fm0((_this).f20, (_719_v34).f20, globalState)).plus(new BigNumber((_module.__default.fm10(new BigNumber((_747_v54).length), new BigNumber(462), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-433))), globalState)).length)));
        let _rhs89 = _748_v55;
        let _rhs90 = (_719_v34).f20;
        let _lhs83 = globalState;
        _748_v55 = _rhs89;
        _lhs83.f0 = _rhs90;
        let _749_v56;
        _749_v56 = _dafny.Seq.of(_717_v32, _717_v32);
        let _750_v57;
        _750_v57 = _dafny.Map.Empty.slice().updateUnsafe(_749_v56,_dafny.Seq.Create(_module.__default.abs(new BigNumber(182)), ((_751_v30) => function (_752_i7) {
          return _751_v30;
        })(_715_v30)));
        _750_v57 = (_750_v57).update(((_module.__default.fm9((_this).f20, globalState)) ? (_749_v56) : (_749_v56)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-291)), ((_753_v30) => function (_754_i8) {
          return _753_v30;
        })(_715_v30)));
        if ((_this).f20) {
          let _755_v58;
          _755_v58 = _dafny.MultiSet.fromElements((_this).f20);
          let _756_v59;
          _756_v59 = _module.D0.create_DC0(_755_v58);
          let _pat_let_tv33 = _755_v58;
          let _757_v60;
          _757_v60 = _dafny.Map.Empty.slice().updateUnsafe(false,function (_pat_let28_0) {
            return function (_758_dt__update__tmp_h4) {
              return function (_pat_let29_0) {
                return function (_759_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_759_dt__update_hcf0_h0);
                }(_pat_let29_0);
              }(_pat_let_tv33);
            }(_pat_let28_0);
          }(_756_v59));
          let _760_v61;
          _760_v61 = _dafny.Seq.of(_757_v60, _dafny.Map.Empty.slice().updateUnsafe(false,_756_v59));
          (globalState).f0 = !_dafny.areEqual(_dafny.Seq.Concat(_760_v61, _760_v61), _760_v61);
          (globalState).f18 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.fm0((_719_v34).f20, true, globalState))).minus((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]));
          let _index106 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          (_714_v29)[_index106] = new BigNumber(260);
          let _761_v62;
          let _nw130 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _761_v62 = _nw130;
          let _762_v63;
          _762_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm9((_719_v34).f20, globalState),_731_i5);
          let _763_v64;
          _763_v64 = _module.D0.create_DC2((_this).f20, _731_i5, _747_v54, (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], new BigNumber((_dafny.Seq.of((_this).f20)).length));
          let _764_v65;
          _764_v65 = _dafny.Seq.of(_763_v64);
          let _765_v66;
          _765_v66 = _dafny.Set.fromElements(_module.D1.create_DC4(_764_v65));
          let _766_v67;
          _766_v67 = _dafny.Set.fromElements((_this).f20, (_this).f20, (_719_v34).f20);
          let _767_v68;
          _767_v68 = _module.D4.create_DC14(!(false), _766_v67);
          _718_v33 = (_718_v33).update((_module.D3.create_DC10(_761_v62, _762_v63, _765_v66, (_this).f21)).dtor_cf20, (_767_v68).dtor_cf25);
          let _768_v69;
          let _nw131 = Array((new BigNumber(5)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = _715_v30;
          _nw131[(_dafny.ONE).toNumber()] = _715_v30;
          _nw131[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
          _nw131[(new BigNumber(3)).toNumber()] = _715_v30;
          _nw131[(new BigNumber(4)).toNumber()] = (_717_v32)[_module.__default.safeIndex((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))], new BigNumber((_717_v32).length))];
          _768_v69 = _nw131;
          let _index107 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_768_v69).length));
          (_768_v69)[_index107] = _715_v30;
          let _index108 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_768_v69).length));
          (_768_v69)[_index108] = new _dafny.CodePoint('u'.codePointAt(0));
        } else {
          let _769_v70;
          _769_v70 = _dafny.MultiSet.fromElements((_731_i5).isLessThan(new BigNumber(548)), (_719_v34).f20, (_719_v34).f20, ((_719_v34).f20) === ((_this).f20));
          _769_v70 = (_769_v70).Union((((_719_v34).f20) ? (_769_v70) : (_769_v70)));
          let _770_v71;
          let _nw132 = new _module.C0();
          _nw132.__ctor((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
          _770_v71 = _nw132;
          let _771_v72;
          _771_v72 = _dafny.Map.Empty.slice().updateUnsafe((_719_v34).f20,_770_v71);
          let _772_v73;
          _772_v73 = _dafny.Seq.of(_770_v71);
          let _773_v74;
          _773_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
          let _774_v75;
          let _nw133 = Array((new BigNumber(17)).toNumber());
          _nw133[(_dafny.ZERO).toNumber()] = _771_v72;
          _nw133[(_dafny.ONE).toNumber()] = _771_v72;
          _nw133[(new BigNumber(2)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(3)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(4)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(5)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(6)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(7)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_719_v34).f20,(_772_v73)[_module.__default.safeIndex(new BigNumber((_773_v74).length), new BigNumber((_772_v73).length))]);
          _nw133[(new BigNumber(9)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(10)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(11)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(12)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(13)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(14)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(15)).toNumber()] = _771_v72;
          _nw133[(new BigNumber(16)).toNumber()] = _771_v72;
          _774_v75 = _nw133;
          let _775_v76;
          _775_v76 = _module.D1.create_DC4(_module.__default.fm24(globalState));
          let _776_v77;
          _776_v77 = _module.D0.create_DC2((_719_v34).f20, (_770_v71).f24, _747_v54, new BigNumber(604), (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]);
          let _777_v78;
          _777_v78 = _dafny.Seq.of(_776_v77);
          let _pat_let_tv34 = _777_v78;
          let _778_v79;
          _778_v79 = _dafny.Set.fromElements(function (_pat_let30_0) {
            return function (_779_dt__update__tmp_h5) {
              return function (_pat_let31_0) {
                return function (_780_dt__update_hcf10_h0) {
                  return _module.D1.create_DC4(_780_dt__update_hcf10_h0);
                }(_pat_let31_0);
              }(_pat_let_tv34);
            }(_pat_let30_0);
          }(_775_v76), _module.D1.create_DC4(_777_v78));
          let _781_v80;
          _781_v80 = _module.D3.create_DC10(_774_v75, _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21), _778_v79, new BigNumber(-727));
          let _index109 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          (_714_v29)[_index109] = (_781_v80).dtor_cf20;
          (globalState).f0 = ((!_dafny.Seq.contains(_dafny.Seq.of(_module.__default.fm25(globalState)), function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(527), new BigNumber(117))) {
              let _782_v81 = _compr_14;
              if (((new BigNumber(527)).isLessThanOrEqualTo(_782_v81)) && ((_782_v81).isLessThan(new BigNumber(117)))) {
                _coll14.push([(_782_v81).multipliedBy((_770_v71).f24),(_719_v34).f20]);
              }
            }
            return _coll14;
          }())) ? ((_719_v34).f20) : ((_719_v34).f20));
          let _index110 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length));
          (_714_v29)[_index110] = (_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))];
          _769_v70 = _dafny.MultiSet.fromElements((_this).f20);
        }
      }
      r0 = _module.__default.safeModuloInt(((_this).f21).plus((_714_v29)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_714_v29).length))]), (_this).f21);
      return r0;
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
