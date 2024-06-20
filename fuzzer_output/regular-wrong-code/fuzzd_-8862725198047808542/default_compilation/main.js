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
    static fm3(p0, globalState) {
      return ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)).plus(new BigNumber(618))).multipliedBy((new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_0_i0) {
          return new BigNumber(887);
        })).Elements) {
          let _1_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), function (_0_i0) {
            return new BigNumber(887);
          }), _1_v0)) {
            _coll0.push([_module.__default.safeDivisionInt(_1_v0, new BigNumber((_dafny.Seq.UnicodeFromString("lll")).length)),true]);
          }
        }
        return _coll0;
      }()).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-226)), function (_2_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length)));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      let _source0 = ((!(!(false))) ? (_module.D3.create_DC11(_module.D3.create_DC10(new _dafny.CodePoint('y'.codePointAt(0))))) : (_module.D3.create_DC11(_module.D3.create_DC10(new _dafny.CodePoint('u'.codePointAt(0))))));
      if (_source0.is_DC9) {
        let _3___mcc_h0 = (_source0).cf8;
        let _4___mcc_h1 = (_source0).cf9;
        let _5___mcc_h2 = (_source0).cf10;
        let _6___mcc_h3 = (_source0).cf11;
        let _7_cf11 = _6___mcc_h3;
        let _8_cf10 = _5___mcc_h2;
        let _9_cf9 = _4___mcc_h1;
        let _10_cf8 = _3___mcc_h0;
        return new _dafny.CodePoint('o'.codePointAt(0));
      } else if (_source0.is_DC10) {
        let _11___mcc_h4 = (_source0).cf12;
        let _12_cf12 = _11___mcc_h4;
        return _12_cf12;
      } else if (_source0.is_DC8) {
        let _13___mcc_h5 = (_source0).cf7;
        let _14_cf7 = _13___mcc_h5;
        return new _dafny.CodePoint('b'.codePointAt(0));
      } else {
        let _15___mcc_h6 = (_source0).cf13;
        let _16_cf13 = _15___mcc_h6;
        return new _dafny.CodePoint('h'.codePointAt(0));
      }
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(!(true))), _dafny.Seq.of(true, true));
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nxm"), _dafny.Seq.UnicodeFromString("uaxq"));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      if ((_dafny.MultiSet.fromElements(new BigNumber(668))).equals(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-644), new BigNumber((_dafny.MultiSet.fromElements(false, true, !(false), false)).cardinality()))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),true)).length), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-321), new BigNumber(102), new BigNumber(974), new BigNumber(250))).length),new _dafny.CodePoint('h'.codePointAt(0))))).cardinality())))) {
        if (true) {
          return _module.D1.create_DC4(false);
        } else {
          return _module.D1.create_DC4(false);
        }
      } else if ((_module.D7.create_DC21(false, new _dafny.CodePoint('o'.codePointAt(0)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-439)), function (_17_i0) {
  return _module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(361))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("w")).length));
})).length)))).dtor_cf30) {
        return _module.D1.create_DC4(true);
      } else {
        return _module.D1.create_DC4(false);
      }
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(false),!(!(true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(false),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(true),true)));
    };
    static fm9(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(true, true, false, true, true));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),false)).Merge((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(896), new BigNumber(351))) {
          let _18_v0 = _compr_1;
          if (((new BigNumber(896)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(351)))) {
            _coll1.push([(_18_v0).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-24)), function (_19_i0) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length);
            }))).cardinality())),true]);
          }
        }
        return _coll1;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(821),false)));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-758), new BigNumber(774))) {
          let _20_v0 = _compr_2;
          if (((new BigNumber(-758)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(774)))) {
            _coll2.push([_module.__default.safeDivisionInt(_20_v0, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(531))).cardinality())), new BigNumber((_dafny.Seq.of(new BigNumber(-896))).length))).length)),new BigNumber((_dafny.Seq.UnicodeFromString("fixlptlef")).length)]);
          }
        }
        return _coll2;
      }()).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(49)))).length), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(829),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(748),false)).length))).length)).minus(new BigNumber(795)));
    };
    static fm12(p0, p1, p2, globalState) {
      if (true) {
        return (_dafny.MultiSet.fromElements(true, true, !(!(true)), true, true)).IsSubsetOf(_dafny.MultiSet.fromElements(false));
      } else {
        return (new BigNumber(761)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(748), new BigNumber(26))).length)));
      }
    };
    static fm13(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(!(!(false)), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(true), _dafny.Seq.of(true, !(!(false)))), false);
    };
    static fm14(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm15(globalState) {
      if (!((_module.D7.create_DC21(true, new _dafny.CodePoint('s'.codePointAt(0)), _dafny.MultiSet.fromElements(new BigNumber(756)))).dtor_cf30)) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nfbfbvbsb")).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(453),false));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jdxlfgub")).length)),false))).Merge(function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-111), new BigNumber(887))) {
            let _21_v0 = _compr_3;
            if (((new BigNumber(-111)).isLessThanOrEqualTo(_21_v0)) && ((_21_v0).isLessThan(new BigNumber(887)))) {
              _coll3.push([_module.__default.safeModuloInt(_21_v0, new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(false))).length),!(true))]);
            }
          }
          return _coll3;
        }());
      }
    };
    static fm16(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uxeiish")).length),new BigNumber(755))).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-863)), function (_22_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of(false, false)).length));
      })), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("fgghpfsb")).length))).length))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(344), new BigNumber(194))) {
          let _23_v0 = _compr_4;
          if (((new BigNumber(344)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(194)))) {
            _coll4.add((_23_v0).multipliedBy(new BigNumber(-608)));
          }
        }
        return _coll4;
      }()).length)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true, !(false))).length)))));
    };
    static fm17(globalState) {
      return _dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm18(p0, globalState) {
      let _source1 = _module.D3.create_DC8(_dafny.Seq.UnicodeFromString("ycfygtdb"));
      if (_source1.is_DC9) {
        let _24___mcc_h0 = (_source1).cf8;
        let _25___mcc_h1 = (_source1).cf9;
        let _26___mcc_h2 = (_source1).cf10;
        let _27___mcc_h3 = (_source1).cf11;
        let _28_cf11 = _27___mcc_h3;
        let _29_cf10 = _26___mcc_h2;
        let _30_cf9 = _25___mcc_h1;
        let _31_cf8 = _24___mcc_h0;
        return _module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC4(_30_cf9)));
      } else if (_source1.is_DC10) {
        let _32___mcc_h4 = (_source1).cf12;
        let _33_cf12 = _32___mcc_h4;
        return _module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC3(!(true))));
      } else if (_source1.is_DC8) {
        let _34___mcc_h5 = (_source1).cf7;
        let _35_cf7 = _34___mcc_h5;
        return _module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC4(false))));
      } else {
        let _36___mcc_h6 = (_source1).cf13;
        let _37_cf13 = _36___mcc_h6;
        return _module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC5(_module.D1.create_DC3(!(true)))));
      }
    };
    static fm19(p0, p1, globalState) {
      if (!(false) || ((_module.D7.create_DC21(false, new _dafny.CodePoint('f'.codePointAt(0)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(914)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(355), new BigNumber((_dafny.Seq.UnicodeFromString("lloh")).length), new BigNumber(384))).cardinality())))).length), new BigNumber(-576)))).dtor_cf30)) {
        return _module.D1.create_DC3(!(!(false)));
      } else {
        return _module.D1.create_DC3(false);
      }
    };
    static fm20(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("cytq"), _dafny.Seq.UnicodeFromString("r")), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("kcqqyceot")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("omkpbmjjq"))));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new _dafny.CodePoint('g'.codePointAt(0)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new _dafny.CodePoint('d'.codePointAt(0)))).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
          let _38_v0 = _compr_5;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).contains(_38_v0)) {
            _coll5.push([_38_v0,new _dafny.CodePoint('p'.codePointAt(0))]);
          }
        }
        return _coll5;
      }()));
    };
    static m5(p0, p1, globalState) {
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _39_v0;
      _39_v0 = _dafny.Seq.UnicodeFromString("ojcommo");
      let _40_v1;
      _40_v1 = _dafny.Set.fromElements(_39_v0);
      let _41_i0;
      _41_i0 = _dafny.ZERO;
      L0: {
        while ((new BigNumber((_40_v1).length)).isLessThan(new BigNumber(457))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_41_i0)) {
              break L0;
            }
            _41_i0 = (_41_i0).plus(_dafny.ONE);
            let _42_v2;
            let _nw0 = Array((new BigNumber(25)).toNumber()).fill(false);
            _42_v2 = _nw0;
            let _43_v3;
            _43_v3 = true;
            let _index0 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
            (_42_v2)[_index0] = _module.__default.fm12(_43_v3, true, _43_v3, globalState);
            let _44_v4;
            let _init0 = function (_45_i1) {
              return (_45_i1).plus(new BigNumber(364));
            };
            let _nw1 = Array((new BigNumber(16)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
              _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _44_v4 = _nw1;
            let _46_v5;
            _46_v5 = new BigNumber(521);
            let _index1 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
            (_44_v4)[_index1] = _46_v5;
            let _47_v6;
            _47_v6 = _module.D3.create_DC8(_39_v0);
            let _index2 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
            let _index3 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
            let _rhs0 = (_module.D1.create_DC3(_43_v3)).dtor_cf3;
            let _rhs1 = _46_v5;
            let _rhs2 = _47_v6;
            let _lhs0 = _42_v2;
            let _lhs1 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
            let _lhs2 = _44_v4;
            let _lhs3 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
            _lhs0[_lhs1] = _rhs0;
            _lhs2[_lhs3] = _rhs1;
            _47_v6 = _rhs2;
            if (true) {
              let _48_v7;
              let _nw2 = Array((new BigNumber(19)).toNumber()).fill([]);
              _48_v7 = _nw2;
              let _index4 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_48_v7).length));
              (_48_v7)[_index4] = _44_v4;
              let _49_v8;
              _49_v8 = _dafny.Set.fromElements((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], _43_v3);
              let _50_v10;
              _50_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
                let _coll6 = new _dafny.Map();
                for (const _compr_6 of (p1).Elements) {
                  let _51_v9 = _compr_6;
                  if ((p1).contains(_51_v9)) {
                    _coll6.push([(_51_v9).minus((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]),_43_v3]);
                  }
                }
                return _coll6;
              }()).length),_43_v3);
              let _index5 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_48_v7).length));
              (_48_v7)[_index5] = (((_49_v8).IsSubsetOf(_49_v8)) ? (_44_v4) : ((((((_50_v10).contains((p0)[_module.__default.safeIndex(_46_v5, new BigNumber((p0).length))])) ? ((_50_v10).get((p0)[_module.__default.safeIndex(_46_v5, new BigNumber((p0).length))])) : (_43_v3))) ? (_44_v4) : (_44_v4))));
              let _index6 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              (_44_v4)[_index6] = (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))];
              let _index7 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              (_44_v4)[_index7] = (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))];
              let _52_v11;
              _52_v11 = _dafny.Map.Empty.slice().updateUnsafe(_42_v2,_46_v5);
              _52_v11 = ((_52_v11).Merge(_52_v11)).Merge(_52_v11);
              _46_v5 = _module.__default.safeDivisionInt((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], _module.__default.fm3(!(false), globalState));
            } else {
              let _index8 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              (_44_v4)[_index8] = (_dafny.ZERO).minus((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]);
              let _53_v12;
              let _nw3 = new _module.C2();
              _nw3.__ctor((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], _46_v5, (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]);
              _53_v12 = _nw3;
              let _54_v13;
              _54_v13 = _dafny.Map.Empty.slice().updateUnsafe(_46_v5,_53_v12);
              let _55_v14;
              let _nw4 = Array((new BigNumber(28)).toNumber());
              _nw4[(_dafny.ZERO).toNumber()] = _53_v12;
              _nw4[(_dafny.ONE).toNumber()] = _53_v12;
              _nw4[(new BigNumber(2)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(3)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(4)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(5)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(6)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(7)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(8)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(9)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(10)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(11)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(12)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(13)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(14)).toNumber()] = (((_54_v13).contains(new BigNumber(-282))) ? ((_54_v13).get(new BigNumber(-282))) : (_53_v12));
              _nw4[(new BigNumber(15)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(16)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(17)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(18)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(19)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(20)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(21)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(22)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(23)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(24)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(25)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(26)).toNumber()] = _53_v12;
              _nw4[(new BigNumber(27)).toNumber()] = _53_v12;
              _55_v14 = _nw4;
              let _56_v15;
              _56_v15 = _dafny.Map.Empty.slice().updateUnsafe(_55_v14,(_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              let _57_v16;
              _57_v16 = _dafny.Map.Empty.slice().updateUnsafe(_56_v15,(p0)[_module.__default.safeIndex((_dafny.ZERO).minus((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]), new BigNumber((p0).length))]);
              _57_v16 = (_57_v16).update((_56_v15).Merge(_56_v15), (_53_v12).f4);
              let _58_v17;
              _58_v17 = _dafny.Seq.of(_39_v0);
              let _59_v18;
              let _nw5 = Array((new BigNumber(12)).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = _39_v0;
              _nw5[(_dafny.ONE).toNumber()] = (_58_v17)[_module.__default.safeIndex((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], new BigNumber((_58_v17).length))];
              _nw5[(new BigNumber(2)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_39_v0, _39_v0);
              _nw5[(new BigNumber(4)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-331)), function (_60_i2) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              });
              _nw5[(new BigNumber(6)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(7)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(8)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(9)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(10)).toNumber()] = _39_v0;
              _nw5[(new BigNumber(11)).toNumber()] = _39_v0;
              _59_v18 = _nw5;
              let _index9 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_59_v18).length));
              (_59_v18)[_index9] = _dafny.Seq.UnicodeFromString("owumuar");
              let _index10 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
              let _index11 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_59_v18).length));
              let _rhs3 = _43_v3;
              let _rhs4 = _43_v3;
              let _rhs5 = (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))];
              let _rhs6 = _39_v0;
              let _lhs4 = globalState;
              let _lhs5 = _42_v2;
              let _lhs6 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
              let _lhs7 = _53_v12;
              let _lhs8 = _59_v18;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_59_v18).length));
              _lhs4.f2 = _rhs3;
              _lhs5[_lhs6] = _rhs4;
              _lhs7.f5 = _rhs5;
              _lhs8[_lhs9] = _rhs6;
              let _61_v19;
              let _nw6 = new _module.C1();
              _nw6.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus(_53_v12.f5), _53_v12.f5), _53_v12.f5);
              _61_v19 = _nw6;
              let _62_v20;
              let _nw7 = new _module.C2();
              _nw7.__ctor(!(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_43_v3), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("cbblxtt")).length), new BigNumber((_dafny.Seq.of(_43_v3)).length)), _43_v3)).length)).isEqualTo((_53_v12).f4), _46_v5, _53_v12.f5);
              _62_v20 = _nw7;
            }
            let _63_v21;
            _63_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_43_v3, false)).cardinality()),new BigNumber(-512)),_43_v3);
            let _64_v22;
            _64_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], globalState),_46_v5);
            if ((((_63_v21).contains(_64_v22)) ? ((_63_v21).get(_64_v22)) : (_43_v3))) {
              let _65_v23;
              _65_v23 = _dafny.Seq.of((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], _43_v3, (_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], (_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], (_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              let _66_v24;
              _66_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm20((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], _39_v0, _65_v23, globalState),(_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              let _67_v25;
              _67_v25 = _dafny.Seq.of(_39_v0, _39_v0, _39_v0);
              _66_v24 = (_66_v24).update(_67_v25, (_65_v23)[_module.__default.safeIndex((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], new BigNumber((_65_v23).length))]);
              let _index12 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length));
              (_42_v2)[_index12] = false;
              let _index13 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _index14 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _rhs7 = (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))];
              let _rhs8 = (p0)[_module.__default.safeIndex((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], new BigNumber((p0).length))];
              let _lhs10 = _44_v4;
              let _lhs11 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _lhs12 = _44_v4;
              let _lhs13 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              _lhs10[_lhs11] = _rhs7;
              _lhs12[_lhs13] = _rhs8;
              let _68_v26;
              let _nw8 = new _module.C1();
              _nw8.__ctor((_46_v5).multipliedBy((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]), (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]);
              _68_v26 = _nw8;
              let _69_v27;
              let _nw9 = new _module.C3();
              _nw9.__ctor((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              _69_v27 = _nw9;
            } else {
              _39_v0 = _39_v0;
              let _70_v28;
              _70_v28 = _dafny.MultiSet.fromElements((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              let _71_v29;
              _71_v29 = _dafny.Seq.of((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))]);
              let _index15 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _index16 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _rhs9 = _module.__default.safeDivisionInt((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], (_dafny.ZERO).minus(new BigNumber((_70_v28).cardinality())));
              let _rhs10 = (_71_v29)[_module.__default.safeIndex(_module.__default.safeModuloInt((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]), new BigNumber((_71_v29).length))];
              let _rhs11 = new BigNumber(487);
              let _lhs14 = _44_v4;
              let _lhs15 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _lhs16 = globalState;
              let _lhs17 = _44_v4;
              let _lhs18 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              _lhs14[_lhs15] = _rhs9;
              _lhs16.f2 = _rhs10;
              _lhs17[_lhs18] = _rhs11;
              let _72_v30;
              _72_v30 = _dafny.Map.Empty.slice().updateUnsafe(_43_v3,(_dafny.ZERO).minus((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]));
              let _73_v31;
              _73_v31 = _72_v30;
              _72_v30 = ((_72_v30).update(_43_v3, _46_v5)).Merge((_73_v31));
              let _74_v32;
              _74_v32 = new _dafny.CodePoint('f'.codePointAt(0));
              let _75_v33;
              _75_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(true, (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))], _module.__default.fm12((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], _43_v3, _43_v3, globalState), new BigNumber((_39_v0).length), globalState),_74_v32);
              let _76_v34;
              _76_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_43_v3);
              let _77_v35;
              _77_v35 = _dafny.MultiSet.fromElements(_75_v33, _module.__default.fm21(_46_v5, (((_76_v34).contains(_43_v3)) ? ((_76_v34).get(_43_v3)) : (_43_v3)), new BigNumber((_76_v34).length), _46_v5, globalState), _75_v33, _dafny.Map.Empty.slice().updateUnsafe(_74_v32,_74_v32));
              let _78_v36;
              let _nw10 = new _module.C2();
              _nw10.__ctor((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))], _46_v5, _46_v5);
              _78_v36 = _nw10;
              let _79_v37;
              _79_v37 = _dafny.Seq.of(_78_v36);
              let _index17 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              let _rhs12 = (_77_v35).Union((_77_v35).update(_75_v33, _module.__default.abs(new BigNumber((_79_v37).length))));
              let _rhs13 = (_dafny.ZERO).minus(_46_v5);
              let _rhs14 = (new BigNumber(-531)).plus(((_dafny.ZERO).minus((_dafny.ZERO).minus(_46_v5))).plus((_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))]));
              let _lhs19 = _44_v4;
              let _lhs20 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length));
              _77_v35 = _rhs12;
              _46_v5 = _rhs13;
              _lhs19[_lhs20] = _rhs14;
              let _80_v38;
              _80_v38 = _dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC13((_42_v2)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_42_v2).length))])).dtor_cf15,_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0))));
              _80_v38 = (_80_v38).update(_43_v3, _dafny.MultiSet.FromArray(_39_v0));
            }
            _46_v5 = (_44_v4)[_module.__default.safeIndex(new BigNumber(965), new BigNumber((_44_v4).length))];
          }
        }
      }
      (globalState).f2 = false;
      let _81_v39;
      _81_v39 = false;
      r0 = _81_v39;
      let _82_v40;
      let _nw11 = Array((new BigNumber(4)).toNumber()).fill([]);
      _82_v40 = _nw11;
      _82_v40 = _82_v40;
      let _83_v41;
      _83_v41 = new BigNumber(-496);
      let _84_v42;
      _84_v42 = _dafny.Set.fromElements(_81_v39, _81_v39);
      let _hi0 = _module.__default.safeModuloInt(_83_v41, new BigNumber((_84_v42).length));
      for (let _85_i3 = _83_v41; _85_i3.isLessThan(_hi0); _85_i3 = _85_i3.plus(_dafny.ONE)) {
        let _86_v43;
        _86_v43 = _dafny.MultiSet.fromElements(_81_v39);
        let _87_v44;
        _87_v44 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_85_i3),_86_v43);
        let _88_v45;
        _88_v45 = _module.D6.create_DC18(_39_v0, _81_v39, (_dafny.ZERO).minus(_83_v41));
        let _89_v46;
        _89_v46 = _dafny.Set.fromElements((((_87_v44).contains((_88_v45).dtor_cf27)) ? ((_87_v44).get((_88_v45).dtor_cf27)) : (_86_v43)));
        _89_v46 = ((_89_v46).Union(_dafny.Set.fromElements(_86_v43))).Difference(_89_v46);
        let _90_v47;
        let _nw12 = Array((new BigNumber(26)).toNumber());
        _90_v47 = _nw12;
        let _rhs15 = _90_v47;
        let _rhs16 = ((!(_81_v39)) ? (_81_v39) : (_81_v39));
        let _lhs21 = globalState;
        _90_v47 = _rhs15;
        _lhs21.f2 = _rhs16;
        let _91_v48;
        _91_v48 = _dafny.Map.Empty.slice().updateUnsafe(_83_v41,_39_v0);
        let _92_v49;
        _92_v49 = new _dafny.CodePoint('l'.codePointAt(0));
        _91_v48 = (_91_v48).update(_85_i3, _dafny.Seq.update(_39_v0, _module.__default.safeIndex(_83_v41, new BigNumber((_39_v0).length)), _92_v49));
        let _93_v50;
        let _nw13 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _93_v50 = _nw13;
        let _94_v51;
        _94_v51 = _dafny.Seq.of(_81_v39, _81_v39);
        let _index18 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_93_v50).length));
        (_93_v50)[_index18] = (new BigNumber((_94_v51).length)).minus(_85_i3);
        let _95_v52;
        _95_v52 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)));
        let _96_v53;
        _96_v53 = _module.D3.create_DC9(new BigNumber((p0).length), _module.__default.fm12(_81_v39, _81_v39, _81_v39, globalState), new BigNumber((_95_v52).cardinality()), _81_v39);
        let _index19 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_93_v50).length));
        (_93_v50)[_index19] = (function (_pat_let0_0) {
          return function (_97_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_98_dt__update_hcf11_h0) {
                return _module.D3.create_DC9((_97_dt__update__tmp_h0).dtor_cf8, (_97_dt__update__tmp_h0).dtor_cf9, (_97_dt__update__tmp_h0).dtor_cf10, _98_dt__update_hcf11_h0);
              }(_pat_let1_0);
            }(false);
          }(_pat_let0_0);
        }(_96_v53)).dtor_cf8;
      }
      (globalState).f2 = _81_v39;
      let _99_v54;
      _99_v54 = _dafny.Map.Empty.slice().updateUnsafe(_83_v41,!(_81_v39));
      let _100_v55;
      _100_v55 = _module.D0.create_DC0(new BigNumber((_99_v54).length));
      let _pat_let_tv0 = _81_v39;
      let _pat_let_tv1 = _81_v39;
      let _pat_let_tv2 = _81_v39;
      let _pat_let_tv3 = _83_v41;
      r0 = function (_source2) {
        if (_source2.is_DC1) {
          let _101___mcc_h0 = (_source2).cf1;
          let _102___mcc_h1 = (_source2).cf2;
          let _103_cf2 = _102___mcc_h1;
          let _104_cf1 = _101___mcc_h0;
          return _pat_let_tv0;
        } else if (_source2.is_DC2) {
          return _pat_let_tv1;
        } else {
          let _105___mcc_h2 = (_source2).cf0;
          let _106_cf0 = _105___mcc_h2;
          return _pat_let_tv2;
        }
      }(function (_pat_let2_0) {
        return function (_107_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_108_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_108_dt__update_hcf0_h0);
            }(_pat_let3_0);
          }(_pat_let_tv3);
        }(_pat_let2_0);
      }(_100_v55));
      let _109_v56;
      _109_v56 = new _dafny.CodePoint('e'.codePointAt(0));
      let _110_v57;
      _110_v57 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_39_v0, _39_v0), _module.__default.safeIndex(_83_v41, new BigNumber((_dafny.Seq.Concat(_39_v0, _39_v0)).length)), _109_v56), _module.__default.safeIndex(_83_v41, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_39_v0, _39_v0), _module.__default.safeIndex(_83_v41, new BigNumber((_dafny.Seq.Concat(_39_v0, _39_v0)).length)), _109_v56)).length)), _109_v56),_module.__default.fm3(_81_v39, globalState));
      r1 = _110_v57;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _111_globalState;
      let _nw14 = new _module.GlobalState();
      _nw14.__ctor(true, false, true);
      _111_globalState = _nw14;
      let _112_v0;
      _112_v0 = new _dafny.CodePoint('h'.codePointAt(0));
      let _113_v1;
      _113_v1 = _dafny.MultiSet.fromElements(_112_v0, _112_v0, _112_v0, _112_v0, _112_v0);
      let _114_v2;
      let _nw15 = new _module.C0();
      _nw15.__ctor(_113_v1);
      _114_v2 = _nw15;
      let _115_v3;
      _115_v3 = new BigNumber(295);
      _115_v3 = new BigNumber(-786);
      let _116_v4;
      let _nw16 = Array((new BigNumber(29)).toNumber()).fill(false);
      _116_v4 = _nw16;
      let _117_v5;
      _117_v5 = true;
      let _index20 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
      (_116_v4)[_index20] = _117_v5;
      let _index21 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
      (_116_v4)[_index21] = (_115_v3).isLessThanOrEqualTo(_115_v3);
      let _118_v6;
      let _nw17 = Array((new BigNumber(21)).toNumber()).fill(_module.D4.Default());
      _118_v6 = _nw17;
      let _119_v7;
      _119_v7 = _dafny.Set.fromElements((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
      let _index22 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_118_v6).length));
      (_118_v6)[_index22] = _module.D4.create_DC12(_119_v7);
      let _120_v8;
      _120_v8 = _module.D4.create_DC12(_119_v7);
      let _index23 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_118_v6).length));
      (_118_v6)[_index23] = _120_v8;
      let _121_i0;
      _121_i0 = _dafny.ZERO;
      L1: {
        while (false) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_121_i0)) {
              break L1;
            }
            _121_i0 = (_121_i0).plus(_dafny.ONE);
            let _122_v9;
            let _init1 = ((_123_v0) => function (_124_i1) {
              return _123_v0;
            })(_112_v0);
            let _nw18 = Array((new BigNumber(18)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw18.length); _i0_1++) {
              _nw18[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _122_v9 = _nw18;
            let _index24 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_122_v9).length));
            (_122_v9)[_index24] = new _dafny.CodePoint('h'.codePointAt(0));
            let _125_v10;
            _125_v10 = _dafny.Map.Empty.slice().updateUnsafe(_115_v3,_115_v3);
            let _126_v11;
            _126_v11 = _dafny.Seq.of(new BigNumber((_125_v10).length));
            let _127_v12;
            _127_v12 = _dafny.Seq.of(_117_v5, _117_v5);
            let _index25 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_122_v9).length));
            let _rhs17 = (((_127_v12)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_115_v3, _115_v3)).length), new BigNumber((_127_v12).length))]) ? (new _dafny.CodePoint('t'.codePointAt(0))) : (_112_v0));
            let _rhs18 = _dafny.Seq.Concat(_126_v11, _126_v11);
            let _rhs19 = (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))];
            let _lhs22 = _122_v9;
            let _lhs23 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_122_v9).length));
            _lhs22[_lhs23] = _rhs17;
            _126_v11 = _rhs18;
            _117_v5 = _rhs19;
            let _128_v13;
            _128_v13 = _dafny.Seq.UnicodeFromString("mpqme");
            _128_v13 = _128_v13;
            let _129_v14;
            _129_v14 = _module.D2.create_DC7();
            _129_v14 = _129_v14;
            let _130_v15;
            _130_v15 = _dafny.Map.Empty.slice().updateUnsafe((_126_v11)[_module.__default.safeIndex(_115_v3, new BigNumber((_126_v11).length))],_117_v5);
            (_111_globalState).f2 = !(!(((_117_v5) ? (_130_v15) : (_130_v15))).equals(_dafny.Map.Empty.slice().updateUnsafe(_115_v3,_117_v5)));
          }
        }
      }
      _119_v7 = (_119_v7).Union(_119_v7);
      let _hi1 = new BigNumber((_119_v7).length);
      for (let _131_i2 = _115_v3; _131_i2.isLessThan(_hi1); _131_i2 = _131_i2.plus(_dafny.ONE)) {
        let _132_v16;
        _132_v16 = _dafny.MultiSet.fromElements(_module.__default.fm17(_111_globalState), _113_v1);
        if (!(!(((((_132_v16).contains((_114_v2).f7)) ? ((_132_v16).get((_114_v2).f7)) : (_115_v3))).isLessThan(_131_i2)))) {
          let _133_v17;
          _133_v17 = _dafny.Seq.of(_131_i2, _131_i2);
          let _134_v18;
          _134_v18 = _dafny.Seq.UnicodeFromString("nqlglyph");
          (_111_globalState).f2 = ((_133_v17)[_module.__default.safeIndex(new BigNumber((_134_v18).length), new BigNumber((_133_v17).length))]).isLessThanOrEqualTo(_131_i2);
          let _135_v19;
          let _nw19 = new _module.C3();
          _nw19.__ctor((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
          _135_v19 = _nw19;
          let _136_v20;
          _136_v20 = _module.D6.create_DC17(_135_v19);
          _135_v19 = (_136_v20).dtor_cf24;
          (_111_globalState).f2 = _117_v5;
          let _137_v21;
          let _nw20 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _137_v21 = _nw20;
          let _index26 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_137_v21).length));
          (_137_v21)[_index26] = new BigNumber((_134_v18).length);
          let _index27 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_137_v21).length));
          (_137_v21)[_index27] = _module.__default.fm3((_115_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(946)), ((_138_v18) => function (_139_i3) {
            return _138_v18;
          })(_134_v18))).length)), _111_globalState);
          let _140_v22;
          _140_v22 = _dafny.Map.Empty.slice().updateUnsafe(_112_v0,_135_v19);
          _140_v22 = _140_v22;
        } else {
          let _141_v23;
          _141_v23 = _dafny.Seq.of(new BigNumber(867));
          let _142_v24;
          _142_v24 = _dafny.Map.Empty.slice().updateUnsafe(_131_i2,_115_v3);
          let _143_v25;
          _143_v25 = _dafny.Map.Empty.slice().updateUnsafe((_141_v23)[_module.__default.safeIndex(new BigNumber((_142_v24).length), new BigNumber((_141_v23).length))],_116_v4);
          let _144_v26;
          _144_v26 = _dafny.Seq.of(_116_v4);
          _116_v4 = (((_143_v25).contains(_131_i2)) ? ((_143_v25).get(_131_i2)) : ((_144_v26)[_module.__default.safeIndex(_module.__default.fm3((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState), new BigNumber((_144_v26).length))]));
          let _145_v27;
          _145_v27 = _dafny.Map.Empty.slice().updateUnsafe(_117_v5,_112_v0);
          _145_v27 = (_145_v27).update(!((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]), _112_v0);
          let _146_v28;
          let _nw21 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _146_v28 = _nw21;
          let _index28 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_146_v28).length));
          (_146_v28)[_index28] = (_115_v3).plus(_131_i2);
          let _147_v29;
          _147_v29 = _dafny.Seq.UnicodeFromString("oasbjj");
          let _148_v30;
          _148_v30 = _dafny.Map.Empty.slice().updateUnsafe(_147_v29,new BigNumber(692));
          let _index29 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_146_v28).length));
          (_146_v28)[_index29] = (((_148_v30).contains(_147_v29)) ? ((_148_v30).get(_147_v29)) : (_131_i2));
          let _rhs20 = (_147_v29)[_module.__default.safeIndex(_115_v3, new BigNumber((_147_v29).length))];
          let _rhs21 = _131_i2;
          let _rhs22 = ((_dafny.ZERO).minus(_131_i2)).minus(_131_i2);
          _112_v0 = _rhs20;
          _115_v3 = _rhs21;
          _115_v3 = _rhs22;
          (_111_globalState).f2 = (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))];
        }
        let _149_v31;
        _149_v31 = _dafny.Seq.UnicodeFromString("tvpces");
        let _150_v32;
        _150_v32 = _module.D6.create_DC18(_149_v31, _117_v5, _131_i2);
        _150_v32 = _150_v32;
        let _151_v33;
        _151_v33 = _module.D3.create_DC8(_149_v31);
        _149_v31 = _dafny.Seq.Concat((_151_v33).dtor_cf7, _dafny.Seq.Concat(_module.__default.fm6((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_117_v5,_131_i2)).length), (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState), _149_v31));
        let _152_v34;
        let _init2 = ((_153_v4) => function (_154_i4) {
          return _module.D2.create_DC6(_dafny.Seq.of((_153_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_153_v4).length))]));
        })(_116_v4);
        let _nw22 = Array((new BigNumber(24)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
          _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _152_v34 = _nw22;
        let _155_v35;
        _155_v35 = _dafny.Map.Empty.slice().updateUnsafe(_152_v34,!(_117_v5));
        let _156_v36;
        let _nw23 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _156_v36 = _nw23;
        let _157_v37;
        _157_v37 = _dafny.Map.Empty.slice().updateUnsafe(_156_v36,_155_v35);
        _155_v35 = (((_157_v37).contains(_156_v36)) ? ((_157_v37).get(_156_v36)) : (_155_v35));
      }
      _115_v3 = _115_v3;
      _115_v3 = _115_v3;
      let _158_v38;
      _158_v38 = _dafny.Seq.of(_115_v3);
      let _159_v39;
      let _nw24 = Array((new BigNumber(5)).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = _115_v3;
      _nw24[(_dafny.ONE).toNumber()] = new BigNumber(134);
      _nw24[(new BigNumber(2)).toNumber()] = _115_v3;
      _nw24[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_115_v3);
      _nw24[(new BigNumber(4)).toNumber()] = new BigNumber((_158_v38).length);
      _159_v39 = _nw24;
      let _160_v40;
      _160_v40 = _dafny.Seq.of(_159_v39);
      _160_v40 = _160_v40;
      let _161_v41;
      _161_v41 = _dafny.Seq.UnicodeFromString("oayxq");
      let _162_v42;
      _162_v42 = _module.D0.create_DC2();
      let _pat_let_tv4 = _161_v41;
      let _pat_let_tv5 = _161_v41;
      let _pat_let_tv6 = _161_v41;
      let _pat_let_tv7 = _161_v41;
      _161_v41 = function (_source3) {
        if (_source3.is_DC1) {
          let _163___mcc_h0 = (_source3).cf1;
          let _164___mcc_h1 = (_source3).cf2;
          let _165_cf2 = _164___mcc_h1;
          let _166_cf1 = _163___mcc_h0;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("n"), _pat_let_tv4);
        } else if (_source3.is_DC2) {
          return _dafny.Seq.Concat(_pat_let_tv5, _pat_let_tv6);
        } else {
          let _167___mcc_h2 = (_source3).cf0;
          let _168_cf0 = _167___mcc_h2;
          return _pat_let_tv7;
        }
      }(_162_v42);
      let _index30 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
      (_116_v4)[_index30] = !((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
      if (_117_v5) {
        let _169_v43;
        _169_v43 = _dafny.Map.Empty.slice().updateUnsafe(_115_v3,_159_v39);
        _159_v39 = (((_169_v43).contains(_module.__default.safeDivisionInt(_115_v3, _115_v3))) ? ((_169_v43).get(_module.__default.safeDivisionInt(_115_v3, _115_v3))) : (_159_v39));
        let _170_v44;
        let _nw25 = new _module.C3();
        _nw25.__ctor((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        _170_v44 = _nw25;
        (_111_globalState).f2 = (_170_v44).f3;
        let _171_v45;
        _171_v45 = _dafny.Seq.of((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], true);
        let _172_v46;
        let _init3 = ((_173_v38) => function (_174_i5) {
          return _173_v38;
        })(_158_v38);
        let _nw26 = Array((new BigNumber(8)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw26.length); _i0_3++) {
          _nw26[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _172_v46 = _nw26;
        let _rhs23 = _171_v45;
        let _rhs24 = _172_v46;
        _171_v45 = _rhs23;
        _172_v46 = _rhs24;
        let _175_v47;
        _175_v47 = _module.D1.create_DC4(true);
        let _176_v48;
        _176_v48 = _module.D1.create_DC5(_175_v47);
        let _177_v49;
        _177_v49 = _module.D1.create_DC5(_175_v47);
        let _178_v50;
        _178_v50 = _module.D1.create_DC5(_176_v48);
        let _179_v51;
        let _nw27 = new _module.C1();
        _nw27.__ctor(_115_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_115_v3,_117_v5)).length));
        _179_v51 = _nw27;
        let _180_v52;
        _180_v52 = _dafny.Seq.of(_179_v51, _179_v51, _179_v51, _179_v51);
        _178_v50 = _module.__default.fm18(new BigNumber((_180_v52).length), _111_globalState);
      } else {
        let _181_v53;
        let _nw28 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _181_v53 = _nw28;
        let _182_v54;
        let _nw29 = new _module.C2();
        _nw29.__ctor(true, _115_v3, _115_v3);
        _182_v54 = _nw29;
        let _183_v55;
        _183_v55 = _dafny.Map.Empty.slice().updateUnsafe(_182_v54,true);
        let _index31 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_181_v53).length));
        (_181_v53)[_index31] = _183_v55;
        let _index32 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_181_v53).length));
        (_181_v53)[_index32] = _183_v55;
        _161_v41 = _161_v41;
        let _184_v56;
        _184_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_161_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-346)), ((_185_v0) => function (_186_i6) {
          return _185_v0;
        })(_112_v0))),_117_v5);
        let _index33 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
        (_116_v4)[_index33] = (((_184_v56).contains((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))])) ? ((_184_v56).get((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))])) : (_117_v5));
        let _187_v57;
        _187_v57 = _dafny.Seq.of((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _117_v5, (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _117_v5, (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        let _188_v58;
        _188_v58 = _dafny.Map.Empty.slice().updateUnsafe(_119_v7,(_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        if ((((_187_v57)[_module.__default.safeIndex((_182_v54).f4, new BigNumber((_187_v57).length))]) ? (true) : ((((_188_v58).contains(_module.__default.fm9((_182_v54).f4, _117_v5, new BigNumber((_161_v41).length), _111_globalState))) ? ((_188_v58).get(_module.__default.fm9((_182_v54).f4, _117_v5, new BigNumber((_161_v41).length), _111_globalState))) : (_117_v5))))) {
          let _189_v59;
          _189_v59 = _dafny.MultiSet.fromElements(_159_v39, _159_v39, _159_v39, _159_v39);
          let _190_v60;
          _190_v60 = _dafny.Seq.of(_189_v59, _189_v59, _189_v59);
          let _index34 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
          let _rhs25 = _112_v0;
          let _rhs26 = (_190_v60)[_module.__default.safeIndex(((_182_v54).f4).plus((_182_v54).f4), new BigNumber((_190_v60).length))];
          let _rhs27 = !(_117_v5) || ((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
          let _rhs28 = _117_v5;
          let _rhs29 = _182_v54.f5;
          let _lhs24 = _116_v4;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
          let _lhs26 = _182_v54;
          _112_v0 = _rhs25;
          _189_v59 = _rhs26;
          _117_v5 = _rhs27;
          _lhs24[_lhs25] = _rhs28;
          _lhs26.f5 = _rhs29;
          let _191_v61;
          let _nw30 = new _module.C3();
          _nw30.__ctor((_187_v57)[_module.__default.safeIndex(new BigNumber((_187_v57).length), new BigNumber((_187_v57).length))]);
          _191_v61 = _nw30;
          _191_v61 = _191_v61;
          _115_v3 = (_dafny.ZERO).minus(_115_v3);
          let _192_v62;
          _192_v62 = _dafny.Map.Empty.slice().updateUnsafe((_191_v61).f3,new BigNumber(372));
          (_182_v54).f5 = (((_192_v62).contains(_117_v5)) ? ((_192_v62).get(_117_v5)) : ((_182_v54).f4));
          let _193_v63;
          _193_v63 = _dafny.MultiSet.fromElements(_158_v38, _dafny.Seq.Concat(_158_v38, _158_v38), _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_194_v54, _195_v3) => function (_196_i7) {
            return (_module.D0.create_DC1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_194_v54).f4,_195_v3)).length), _194_v54.f5)).dtor_cf2;
          })(_182_v54, _115_v3)));
          (_182_v54).f5 = (((_193_v63).contains((((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]) ? (_158_v38) : (_dafny.Seq.of((_182_v54).f4, _115_v3))))) ? ((_193_v63).get((((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]) ? (_158_v38) : (_dafny.Seq.of((_182_v54).f4, _115_v3))))) : (_182_v54.f5));
        } else {
          let _index35 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
          (_116_v4)[_index35] = _module.__default.fm12(false, (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _117_v5, _111_globalState);
          (_182_v54).f5 = new BigNumber(605);
          let _index36 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_159_v39).length));
          (_159_v39)[_index36] = (_182_v54).f4;
          let _index37 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_159_v39).length));
          (_159_v39)[_index37] = _182_v54.f5;
          let _index38 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_159_v39).length));
          (_159_v39)[_index38] = (_115_v3).plus(_115_v3);
          let _197_v64;
          let _nw31 = Array((new BigNumber(5)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _115_v3;
          _nw31[(_dafny.ONE).toNumber()] = new BigNumber((_158_v38).length);
          _nw31[(new BigNumber(2)).toNumber()] = new BigNumber(745);
          _nw31[(new BigNumber(3)).toNumber()] = (_182_v54).f4;
          _nw31[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(728), (_182_v54).f4);
          _197_v64 = _nw31;
          let _198_v65;
          let _nw32 = Array((new BigNumber(10)).toNumber());
          _198_v65 = _nw32;
          let _199_v66;
          _199_v66 = _dafny.Set.fromElements(_198_v65);
          let _index39 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_159_v39).length));
          let _rhs30 = _197_v64;
          let _rhs31 = _115_v3;
          let _rhs32 = !((_dafny.Set.fromElements(_198_v65)).IsSubsetOf(_199_v66));
          let _lhs27 = _159_v39;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_159_v39).length));
          let _lhs29 = _111_globalState;
          _197_v64 = _rhs30;
          _lhs27[_lhs28] = _rhs31;
          _lhs29.f2 = _rhs32;
        }
        (_182_v54).f5 = _182_v54.f5;
      }
      let _source4 = _162_v42;
      if (_source4.is_DC1) {
        let _200___mcc_h3 = (_source4).cf1;
        let _201___mcc_h4 = (_source4).cf2;
        let _202_cf2 = _201___mcc_h4;
        let _203_cf1 = _200___mcc_h3;
        let _204_v67;
        _204_v67 = _dafny.Map.Empty.slice().updateUnsafe((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))],_dafny.Seq.contains(_dafny.Seq.of(_117_v5), _module.__default.fm12((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], !(_117_v5), _117_v5, _111_globalState)));
        let _205_v68;
        let _nw33 = new _module.C2();
        _nw33.__ctor(_117_v5, _202_cf2, _202_cf2);
        _205_v68 = _nw33;
        let _206_v69;
        _206_v69 = _dafny.Map.Empty.slice().updateUnsafe(_115_v3,_205_v68);
        _204_v67 = (_204_v67).update(!(!((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))])), !((new BigNumber((_206_v69).length)).isLessThanOrEqualTo(new BigNumber((_158_v38).length))));
        _202_cf2 = _202_cf2;
        let _207_v70;
        let _nw34 = new _module.C1();
        _nw34.__ctor(_module.__default.fm3((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState), _115_v3);
        _207_v70 = _nw34;
        _207_v70 = _207_v70;
        let _208_v71;
        _208_v71 = _dafny.MultiSet.fromElements(_117_v5);
        let _209_v72;
        _209_v72 = _module.D6.create_DC18(_161_v41, _117_v5, new BigNumber(((_208_v71).update(_117_v5, _module.__default.abs(_115_v3))).cardinality()));
        let _210_v73;
        _210_v73 = _module.D6.create_DC19(_209_v72);
        _210_v73 = _210_v73;
      } else if (_source4.is_DC2) {
        let _211_v74;
        let _nw35 = new _module.C1();
        _nw35.__ctor((_dafny.ZERO).minus(_115_v3), new BigNumber(933));
        _211_v74 = _nw35;
        _211_v74 = _211_v74;
        _115_v3 = (_115_v3).plus(_115_v3);
        let _212_v75;
        let _nw36 = Array((new BigNumber(4)).toNumber());
        _212_v75 = _nw36;
        let _index40 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_212_v75).length));
        (_212_v75)[_index40] = _114_v2;
        let _index41 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_212_v75).length));
        (_212_v75)[_index41] = _114_v2;
        let _213_v76;
        let _nw37 = Array((new BigNumber(14)).toNumber());
        _213_v76 = _nw37;
        _213_v76 = _213_v76;
      } else {
        let _214___mcc_h5 = (_source4).cf0;
        let _215_cf0 = _214___mcc_h5;
        let _216_v77;
        _216_v77 = _dafny.Map.Empty.slice().updateUnsafe(_115_v3,_115_v3);
        let _217_v78;
        _217_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_216_v77).length),_115_v3);
        let _218_v79;
        _218_v79 = _module.D0.create_DC0(new BigNumber((_217_v78).length));
        let _219_v80;
        _219_v80 = _dafny.Seq.of(_218_v79);
        _219_v80 = _dafny.Seq.update(_dafny.Seq.Concat(_219_v80, _219_v80), _module.__default.safeIndex((_215_cf0).plus(_215_cf0), new BigNumber((_dafny.Seq.Concat(_219_v80, _219_v80)).length)), _218_v79);
        let _220_v81;
        let _nw38 = new _module.C1();
        _nw38.__ctor((_dafny.ZERO).minus(_module.__default.fm3((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState)), new BigNumber((_119_v7).length));
        _220_v81 = _nw38;
        let _221_v82;
        _221_v82 = _dafny.Seq.of(_220_v81);
        let _222_v83;
        _222_v83 = _dafny.Seq.of((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _117_v5);
        let _223_v84;
        _223_v84 = _dafny.Set.fromElements(_222_v83);
        let _224_v85;
        _224_v85 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(629), (_dafny.ZERO).minus(_115_v3)),(_221_v82)[_module.__default.safeIndex(new BigNumber((_223_v84).length), new BigNumber((_221_v82).length))]);
        _224_v85 = ((_224_v85).Merge(_224_v85)).Merge(_224_v85);
        (_111_globalState).f2 = false;
        let _225_v86;
        let _226_v87;
        let _227_v88;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_220_v81).m1((((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]) ? (_159_v39) : (_159_v39)), new _dafny.CodePoint('u'.codePointAt(0)), _111_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _225_v86 = _out0;
        _226_v87 = _out1;
        _227_v88 = _out2;
      }
      let _source5 = _module.D3.create_DC8(_161_v41);
      if (_source5.is_DC9) {
        let _228___mcc_h6 = (_source5).cf8;
        let _229___mcc_h7 = (_source5).cf9;
        let _230___mcc_h8 = (_source5).cf10;
        let _231___mcc_h9 = (_source5).cf11;
        let _232_cf11 = _231___mcc_h9;
        let _233_cf10 = _230___mcc_h8;
        let _234_cf9 = _229___mcc_h7;
        let _235_cf8 = _228___mcc_h6;
        _234_cf9 = (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))];
        let _236_v89;
        _236_v89 = _dafny.MultiSet.fromElements(_235_cf8, new BigNumber(-952));
        let _237_v90;
        _237_v90 = _dafny.Set.fromElements((((_236_v89).contains(_235_cf8)) ? ((_236_v89).get(_235_cf8)) : (_115_v3)), _233_cf10, new BigNumber(-994), _115_v3);
        let _238_v91;
        let _239_v92;
        let _out3;
        let _out4;
        let _outcollector1 = _module.__default.m5(_dafny.Seq.Concat(_dafny.Seq.of(_233_cf10), _158_v38), _237_v90, _111_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _238_v91 = _out3;
        _239_v92 = _out4;
        let _240_v94;
        let _241_v95;
        let _out5;
        let _out6;
        let _outcollector2 = _module.__default.m5(_158_v38, function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-340), new BigNumber(994))) {
            let _242_v93 = _compr_7;
            if (((new BigNumber(-340)).isLessThanOrEqualTo(_242_v93)) && ((_242_v93).isLessThan(new BigNumber(994)))) {
              _coll7.add((_242_v93).minus(_235_cf8));
            }
          }
          return _coll7;
        }(), _111_globalState);
        _out5 = _outcollector2[0];
        _out6 = _outcollector2[1];
        _240_v94 = _out5;
        _241_v95 = _out6;
        if (!(_238_v91)) {
          let _243_v96;
          let _244_v97;
          let _out7;
          let _out8;
          let _outcollector3 = _module.__default.m5(_158_v38, _237_v90, _111_globalState);
          _out7 = _outcollector3[0];
          _out8 = _outcollector3[1];
          _243_v96 = _out7;
          _244_v97 = _out8;
          let _rhs33 = _dafny.Seq.of(((_234_cf9) ? (_159_v39) : (_159_v39)), _159_v39, _159_v39, _159_v39);
          let _rhs34 = _115_v3;
          _160_v40 = _rhs33;
          _235_cf8 = _rhs34;
          (_111_globalState).f2 = !(_232_cf11) || (_117_v5);
          let _245_v98;
          _245_v98 = _module.D1.create_DC3(_234_cf9);
          let _pat_let_tv8 = _116_v4;
          let _pat_let_tv9 = _116_v4;
          let _246_v99;
          _246_v99 = _dafny.MultiSet.fromElements(_module.__default.fm19(_234_cf9, !((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]), _111_globalState), _245_v98, _245_v98, _245_v98, function (_pat_let4_0) {
            return function (_247_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_248_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_248_dt__update_hcf3_h0);
                }(_pat_let5_0);
              }((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_pat_let_tv8).length))]);
            }(_pat_let4_0);
          }(_245_v98));
          let _249_v100;
          let _nw39 = Array((new BigNumber(6)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _246_v99;
          _nw39[(_dafny.ONE).toNumber()] = _246_v99;
          _nw39[(new BigNumber(2)).toNumber()] = (_246_v99).Intersect((_dafny.MultiSet.fromElements(_245_v98)).update(_module.D1.create_DC3(_240_v94), _module.__default.abs(_233_cf10)));
          _nw39[(new BigNumber(3)).toNumber()] = _246_v99;
          _nw39[(new BigNumber(4)).toNumber()] = _246_v99;
          _nw39[(new BigNumber(5)).toNumber()] = _246_v99;
          _249_v100 = _nw39;
          _249_v100 = _249_v100;
          let _250_v101;
          _250_v101 = _module.D7.create_DC20(_158_v38);
          let _index42 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
          let _rhs35 = _dafny.Seq.IsPrefixOf((_250_v101).dtor_cf29, _158_v38);
          let _rhs36 = _dafny.Seq.Concat(_module.__default.fm6(_117_v5, _233_cf10, (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState), _dafny.Seq.UnicodeFromString("ygn"));
          let _rhs37 = !((_115_v3).isLessThan(_233_cf10));
          let _lhs30 = _111_globalState;
          let _lhs31 = _116_v4;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
          _lhs30.f2 = _rhs35;
          _161_v41 = _rhs36;
          _lhs31[_lhs32] = _rhs37;
        } else {
          let _251_v102;
          _251_v102 = _dafny.Seq.of(_161_v41);
          let _252_v103;
          _252_v103 = _dafny.MultiSet.fromElements((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
          let _253_v104;
          _253_v104 = _module.D6.create_DC18(_161_v41, true, new BigNumber((_252_v103).cardinality()));
          let _254_v105;
          _254_v105 = _module.D6.create_DC18(_dafny.Seq.UnicodeFromString("q"), _238_v91, (_253_v104).dtor_cf27);
          let _255_v106;
          _255_v106 = _dafny.Seq.of(_161_v41, (_251_v102)[_module.__default.safeIndex(_115_v3, new BigNumber((_251_v102).length))], _dafny.Seq.Concat((_253_v104).dtor_cf25, _161_v41));
          _251_v102 = _255_v106;
          let _256_v107;
          _256_v107 = _module.D0.create_DC0(_115_v3);
          _235_cf8 = (_235_cf8).plus((_256_v107).dtor_cf0);
          let _257_v108;
          let _nw40 = new _module.C3();
          _nw40.__ctor(_232_cf11);
          _257_v108 = _nw40;
          let _258_v109;
          let _nw41 = new _module.C2();
          _nw41.__ctor(true, new BigNumber((_236_v89).cardinality()), _module.__default.safeModuloInt(new BigNumber((_161_v41).length), _233_cf10));
          _258_v109 = _nw41;
          let _259_v110;
          let _nw42 = new _module.C2();
          _nw42.__ctor((_257_v108).f3, ((_258_v109).f4).multipliedBy(_233_cf10), new BigNumber(685));
          _259_v110 = _nw42;
          let _260_v111;
          _260_v111 = _dafny.Map.Empty.slice().updateUnsafe(_232_cf11,(_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
          let _rhs38 = (_260_v111).equals((_dafny.Map.Empty.slice().updateUnsafe(_117_v5,_240_v94)).Merge(_260_v111));
          let _rhs39 = ((_238_v91) ? (_257_v108) : (_257_v108));
          let _rhs40 = _258_v109;
          let _rhs41 = _259_v110;
          let _lhs33 = _111_globalState;
          _lhs33.f2 = _rhs38;
          _257_v108 = _rhs39;
          _258_v109 = _rhs40;
          _259_v110 = _rhs41;
          _260_v111 = _module.__default.fm14(new BigNumber(543), _111_globalState);
          _234_cf9 = (_257_v108).f3;
        }
      } else if (_source5.is_DC10) {
        let _261___mcc_h10 = (_source5).cf12;
        let _262_cf12 = _261___mcc_h10;
        let _index43 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
        (_116_v4)[_index43] = !((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        let _263_v112;
        let _nw43 = new _module.C2();
        _nw43.__ctor((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _115_v3, _115_v3);
        _263_v112 = _nw43;
        let _264_v113;
        let _nw44 = new _module.C3();
        _nw44.__ctor((_263_v112).f6);
        _264_v113 = _nw44;
        let _265_v114;
        _265_v114 = _dafny.Map.Empty.slice().updateUnsafe((_263_v112).f6,_264_v113);
        let _266_v115;
        _266_v115 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-744),(_263_v112).f6);
        let _267_v116;
        _267_v116 = _dafny.Map.Empty.slice().updateUnsafe((((_265_v114).contains(_117_v5)) ? ((_265_v114).get(_117_v5)) : (_264_v113)),_266_v115);
        let _268_v118;
        _268_v118 = _dafny.Seq.of(_158_v38);
        let _269_v119;
        _269_v119 = _dafny.Seq.of((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], false, !(_117_v5), (_263_v112).f6);
        _267_v116 = (_267_v116).update(_264_v113, function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of ((_268_v118)[_module.__default.safeIndex(new BigNumber((_269_v119).length), new BigNumber((_268_v118).length))]).Elements) {
            let _270_v117 = _compr_8;
            if (_dafny.Seq.contains((_268_v118)[_module.__default.safeIndex(new BigNumber((_269_v119).length), new BigNumber((_268_v118).length))], _270_v117)) {
              _coll8.push([(_270_v117).minus(new BigNumber((_161_v41).length)),(_263_v112).f6]);
            }
          }
          return _coll8;
        }());
        _115_v3 = _module.__default.safeModuloInt(_115_v3, _module.__default.safeDivisionInt(new BigNumber((_158_v38).length), _115_v3));
      } else if (_source5.is_DC8) {
        let _271___mcc_h11 = (_source5).cf7;
        let _272_cf7 = _271___mcc_h11;
        let _index44 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_159_v39).length));
        (_159_v39)[_index44] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-902)), ((_273_v3) => function (_274_i8) {
          return _273_v3;
        })(_115_v3))).length);
        let _index45 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_159_v39).length));
        (_159_v39)[_index45] = _module.__default.safeModuloInt(_115_v3, _115_v3);
        let _275_v120;
        _275_v120 = _dafny.Seq.of((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        (_111_globalState).f2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_275_v120, _module.__default.safeIndex(_115_v3, new BigNumber((_275_v120).length)), (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]), _275_v120);
        let _276_v121;
        _276_v121 = _module.D7.create_DC22(new BigNumber((_275_v120).length));
        let _277_v122;
        _277_v122 = _module.D0.create_DC0((_276_v121).dtor_cf33);
        _115_v3 = (_277_v122).dtor_cf0;
        _158_v38 = _dafny.Seq.Concat(_dafny.Seq.Concat(_158_v38, _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _117_v5)).length))), _158_v38);
      } else {
        let _278___mcc_h12 = (_source5).cf13;
        let _279_cf13 = _278___mcc_h12;
        let _280_v123;
        _280_v123 = _dafny.Map.Empty.slice().updateUnsafe(_117_v5,(_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))]);
        let _281_v124;
        _281_v124 = _dafny.Map.Empty.slice().updateUnsafe((_280_v123).update(_117_v5, _117_v5),_115_v3);
        let _282_v125;
        _282_v125 = _dafny.Map.Empty.slice().updateUnsafe(_281_v124,_159_v39);
        _282_v125 = (_282_v125).update(_281_v124, _159_v39);
        let _283_v126;
        _283_v126 = _module.D8.create_DC23(_114_v2);
        _114_v2 = (_283_v126).dtor_cf34;
        let _284_v127;
        let _nw45 = new _module.C1();
        _nw45.__ctor((_115_v3).minus(_115_v3), (_115_v3).multipliedBy((_dafny.ZERO).minus(_115_v3)));
        _284_v127 = _nw45;
        _117_v5 = _117_v5;
      }
      let _hi2 = _115_v3;
      for (let _285_i9 = (_158_v38)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_158_v38).length))]; _285_i9.isLessThan(_hi2); _285_i9 = _285_i9.plus(_dafny.ONE)) {
        let _286_v128;
        _286_v128 = _dafny.Set.fromElements(_115_v3);
        let _287_v129;
        let _288_v130;
        let _out9;
        let _out10;
        let _outcollector4 = _module.__default.m5(_158_v38, _286_v128, _111_globalState);
        _out9 = _outcollector4[0];
        _out10 = _outcollector4[1];
        _287_v129 = _out9;
        _288_v130 = _out10;
        let _289_v131;
        let _nw46 = new _module.C2();
        _nw46.__ctor(false, _115_v3, _115_v3);
        _289_v131 = _nw46;
        let _nw47 = new _module.C2();
        _nw47.__ctor(_287_v129, _285_i9, _115_v3);
        _289_v131 = _nw47;
        let _290_v132;
        let _nw48 = new _module.C3();
        _nw48.__ctor(_module.__default.fm12(_117_v5, _117_v5, (_116_v4)[_module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length))], _111_globalState));
        _290_v132 = _nw48;
        let _index46 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
        let _rhs42 = _117_v5;
        let _rhs43 = (new BigNumber(((_dafny.Set.fromElements(_290_v132, _290_v132)).Intersect(_dafny.Set.fromElements(_290_v132))).length)).isLessThan(_289_v131.f5);
        let _rhs44 = _117_v5;
        let _rhs45 = _290_v132;
        let _lhs34 = _116_v4;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_116_v4).length));
        let _lhs36 = _111_globalState;
        _117_v5 = _rhs42;
        _lhs34[_lhs35] = _rhs43;
        _lhs36.f2 = _rhs44;
        _290_v132 = _rhs45;
        _115_v3 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(-302), (_dafny.ZERO).minus((_dafny.ZERO).minus(_289_v131.f5))), _module.__default.safeModuloInt((_dafny.ZERO).minus(_285_i9), (_289_v131).f4));
      }
      process.stdout.write(_dafny.toString((_111_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_111_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_111_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_112_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v1).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_114_v2).f7).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_115_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_117_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_118_v6)[new BigNumber(3)]).dtor_cf14).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v7).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v8).dtor_cf14).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_121_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_158_v38, _dafny.Seq.of(_dafny.ZERO, new BigNumber(2), _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v39)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v39)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v39)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v39)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v39)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_160_v40).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_161_v41).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC2() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D1(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false);
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
    static create_DC7() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7();
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
    static create_DC9(cf8, cf9, cf10, cf11) {
      let $dt = new D3(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC10(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8(cf7) {
      let $dt = new D3(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC11(cf13) {
      let $dt = new D3(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + this.cf7.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO, false, _dafny.ZERO, false);
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
    static create_DC13(cf15) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC14(cf16, cf17, cf18) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC12(cf14) {
      let $dt = new D4(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf14) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(false);
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
    static create_DC16(cf20, cf21, cf22, cf23) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC15(cf19) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(_dafny.ZERO, _dafny.MultiSet.Empty, _dafny.Seq.of(), _module.D3.Default());
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
    static create_DC19(cf28) {
      let $dt = new D6(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + this.cf25.toVerbatimString(true) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC21(cf30, cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC22(cf33) {
      let $dt = new D7(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC20(cf29) {
      let $dt = new D7(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21(false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.MultiSet.Empty);
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
    static create_DC24(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC23(cf34) {
      let $dt = new D8(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC25(cf40) {
      let $dt = new D8(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(_dafny.ZERO, [], [], _dafny.ZERO, false);
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
    static create_DC26(cf41) {
      let $dt = new D9(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41);
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

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = false;
      this._f0 = false;
      this._f1 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f7 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f7) {
      let _this = this;
      (_this)._f7 = f7;
      return;
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f5 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return true;
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f4)).length)))).IsProperSubsetOf((_dafny.MultiSet.fromElements((_this).f4)).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _291_v0;
      _291_v0 = _dafny.Seq.UnicodeFromString("ffyidwc");
      let _292_v1;
      _292_v1 = true;
      let _293_v2;
      _293_v2 = _dafny.Seq.of(_292_v1);
      (_this).f5 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_291_v0, _dafny.Seq.UnicodeFromString("ykiwthwa"))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_293_v2)[_module.__default.safeIndex(_this.f5, new BigNumber((_293_v2).length))])).length)));
      let _294_v3;
      _294_v3 = _dafny.Set.fromElements(!(_292_v1));
      r2 = ((_294_v3).Difference(_294_v3)).IsDisjointFrom(_294_v3);
      (_this).f5 = _this.f5;
      r1 = _292_v1;
      let _295_v4;
      _295_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,!(true));
      if (!(((new BigNumber((_294_v3).length)).minus(_this.f5)).isLessThanOrEqualTo(new BigNumber(((_295_v4).Merge(_295_v4)).length)))) {
        let _rhs46 = !(new BigNumber((_291_v0).length)).isEqualTo((_this).f4);
        let _rhs47 = (_this.f5).isLessThanOrEqualTo(_this.f5);
        r0 = _rhs46;
        r1 = _rhs47;
        let _296_v5;
        let _nw49 = Array((new BigNumber(8)).toNumber()).fill(false);
        _296_v5 = _nw49;
        let _297_v6;
        _297_v6 = _dafny.Seq.of(_296_v5, _296_v5, _296_v5, _296_v5);
        let _index47 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((p0).length));
        (p0)[_index47] = _module.__default.safeDivisionInt(new BigNumber((_297_v6).length), _this.f5);
        let _298_v7;
        _298_v7 = _dafny.Seq.of((_this).f4);
        let _299_v8;
        _299_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_298_v7, _module.__default.safeIndex((_this).f4, new BigNumber((_298_v7).length)), _this.f5), _dafny.Seq.of(new BigNumber((_291_v0).length))),_this.f5);
        let _300_v9;
        _300_v9 = _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("khdwjltpl")).length))).length));
        let _index48 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((p0).length));
        let _rhs48 = (_this).f4;
        let _rhs49 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_300_v9).dtor_cf0), (_this).f4);
        let _rhs50 = _299_v8;
        let _lhs37 = p0;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((p0).length));
        let _lhs39 = _this;
        _lhs37[_lhs38] = _rhs48;
        _lhs39.f5 = _rhs49;
        _299_v8 = _rhs50;
        (_this).f5 = (_this).f4;
        let _301_v10;
        _301_v10 = _module.D1.create_DC3(_292_v1);
        let _302_v11;
        _302_v11 = _dafny.Map.Empty.slice().updateUnsafe(_292_v1,(_this).f4);
        let _303_v12;
        _303_v12 = _dafny.MultiSet.fromElements(new BigNumber((_302_v11).length));
        let _304_v13;
        let _305_v14;
        let _306_v15;
        let _307_v16;
        let _out11;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector5 = (_this).m4((_301_v10).dtor_cf3, (((_303_v12).contains(new BigNumber(369))) ? ((_303_v12).get(new BigNumber(369))) : (_module.__default.fm3(_292_v1, globalState))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_308_p1) => function (_309_i0) {
          return _308_p1;
        })(p1)), globalState);
        _out11 = _outcollector5[0];
        _out12 = _outcollector5[1];
        _out13 = _outcollector5[2];
        _out14 = _outcollector5[3];
        _304_v13 = _out11;
        _305_v14 = _out12;
        _306_v15 = _out13;
        _307_v16 = _out14;
        let _310_v17;
        _310_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_304_v13, _304_v13),(_this.f5).plus(new BigNumber(933)));
        let _311_v18;
        _311_v18 = _dafny.MultiSet.fromElements(false);
        _310_v17 = (_310_v17).update(_module.__default.safeModuloInt((_this).f4, _this.f5), new BigNumber(((_dafny.MultiSet.fromElements(true, _292_v1, _292_v1, true, _292_v1)).Intersect((_311_v18).update(_292_v1, _module.__default.abs(_this.f5)))).cardinality()));
      } else {
        let _312_v19;
        _312_v19 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), p1);
        let _313_v20;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_312_v19);
        _313_v20 = _nw50;
        let _314_v21;
        let _nw51 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _314_v21 = _nw51;
        let _index49 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_314_v21).length));
        (_314_v21)[_index49] = p1;
        let _315_v22;
        _315_v22 = _dafny.Seq.of((_this).f4);
        let _index50 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_314_v21).length));
        (_314_v21)[_index50] = _module.__default.fm4((_this).fm2(!(_292_v1), (_this).f4, _292_v1, globalState), _this.f5, _292_v1, _module.__default.safeDivisionInt(_this.f5, (_315_v22)[_module.__default.safeIndex((_this).f4, new BigNumber((_315_v22).length))]), globalState);
        r2 = !(true);
        let _316_v23;
        _316_v23 = _dafny.Seq.of(_315_v22);
        r0 = !_dafny.areEqual(_316_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), ((_317_v22) => function (_318_i1) {
          return _317_v22;
        })(_315_v22)));
        let _index51 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((p0).length));
        (p0)[_index51] = _this.f5;
        let _index52 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((p0).length));
        (p0)[_index52] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm3(_292_v1, globalState)));
      }
      let _319_v24;
      _319_v24 = _module.D1.create_DC3(_292_v1);
      let _320_v25;
      _320_v25 = _dafny.Map.Empty.slice().updateUnsafe(true,_292_v1);
      let _321_v26;
      _321_v26 = _dafny.Seq.of(new BigNumber(160));
      let _322_v27;
      _322_v27 = _dafny.MultiSet.fromElements(_this.f5, _this.f5, new BigNumber((_321_v26).length), _module.__default.fm3(_292_v1, globalState));
      let _323_v28;
      let _nw52 = Array((new BigNumber(25)).toNumber());
      _nw52[(_dafny.ZERO).toNumber()] = ((_292_v1) ? (_292_v1) : ((_this).fm2(_292_v1, _this.f5, _292_v1, globalState)));
      _nw52[(_dafny.ONE).toNumber()] = !(_292_v1) || (_292_v1);
      _nw52[(new BigNumber(2)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(3)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(4)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(5)).toNumber()] = !(!((_319_v24).dtor_cf3));
      _nw52[(new BigNumber(6)).toNumber()] = (_this).fm1(_292_v1, _this.f5, _dafny.Set.fromElements(false, _292_v1), (((_320_v25).contains(_292_v1)) ? ((_320_v25).get(_292_v1)) : (_292_v1)), globalState);
      _nw52[(new BigNumber(7)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(8)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(9)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(10)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(11)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(12)).toNumber()] = (_322_v27).IsSubsetOf(_322_v27);
      _nw52[(new BigNumber(13)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(14)).toNumber()] = !(!(_module.__default.fm3(_292_v1, globalState)).isEqualTo(new BigNumber(252)));
      _nw52[(new BigNumber(15)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(16)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(17)).toNumber()] = (true) || (_292_v1);
      _nw52[(new BigNumber(18)).toNumber()] = !(_292_v1) || (_292_v1);
      _nw52[(new BigNumber(19)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(20)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(21)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(22)).toNumber()] = _292_v1;
      _nw52[(new BigNumber(23)).toNumber()] = (_319_v24).dtor_cf3;
      _nw52[(new BigNumber(24)).toNumber()] = (_293_v2)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_293_v2).length))];
      _323_v28 = _nw52;
      let _index53 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_323_v28).length));
      (_323_v28)[_index53] = !(_292_v1);
      let _index54 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_323_v28).length));
      (_323_v28)[_index54] = _292_v1;
      r0 = (_323_v28)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_323_v28).length))];
      let _324_v29;
      _324_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f5,_this.f5);
      r1 = (new BigNumber((_324_v29).length)).isLessThanOrEqualTo(new BigNumber(-967));
      r2 = (((_323_v28)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_323_v28).length))]) ? (_292_v1) : (_292_v1));
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _325_v0;
      _325_v0 = true;
      (globalState).f2 = _325_v0;
      if (_325_v0) {
        let _326_v1;
        _326_v1 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)));
        let _327_v2;
        let _nw53 = new _module.C0();
        _nw53.__ctor(_326_v1);
        _327_v2 = _nw53;
        let _328_v3;
        let _nw54 = new _module.C0();
        _nw54.__ctor((_327_v2).f7);
        _328_v3 = _nw54;
        let _329_v4;
        let _nw55 = Array((new BigNumber(25)).toNumber()).fill(false);
        _329_v4 = _nw55;
        let _index55 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length));
        (_329_v4)[_index55] = !(_325_v0);
        let _index56 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length));
        (_329_v4)[_index56] = _325_v0;
        _328_v3 = _328_v3;
        let _330_v5;
        _330_v5 = _module.D0.create_DC0(_this.f5);
        let _source6 = _330_v5;
        if (_source6.is_DC1) {
          let _331___mcc_h0 = (_source6).cf1;
          let _332___mcc_h1 = (_source6).cf2;
          let _333_cf2 = _332___mcc_h1;
          let _334_cf1 = _331___mcc_h0;
          let _335_v6;
          _335_v6 = _dafny.Seq.of(_333_cf2);
          let _336_v7;
          _336_v7 = _dafny.Seq.of((_this.f5).isLessThanOrEqualTo(new BigNumber((_335_v6).length)));
          let _337_v8;
          _337_v8 = _dafny.Seq.UnicodeFromString("omrnrrxhq");
          _336_v7 = _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.UnicodeFromString("hyq"), _337_v8));
          let _338_v9;
          _338_v9 = _dafny.Set.fromElements(_325_v0, (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]);
          let _339_v10;
          _339_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,(_338_v9).IsProperSubsetOf(_338_v9));
          let _index57 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length));
          (_329_v4)[_index57] = (((_339_v10).contains(!((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]))) ? ((_339_v10).get(!((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]))) : ((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]));
          let _340_v11;
          _340_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], new BigNumber(-72), (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], globalState),_dafny.Seq.of(new BigNumber(134), new BigNumber(287)));
          _335_v6 = (((_340_v11).contains(_325_v0)) ? ((_340_v11).get(_325_v0)) : (_335_v6));
          (globalState).f2 = _325_v0;
        } else if (_source6.is_DC2) {
          let _341_v12;
          let _nw56 = Array((new BigNumber(21)).toNumber()).fill(_module.D0.Default());
          _341_v12 = _nw56;
          let _index58 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_341_v12).length));
          (_341_v12)[_index58] = _module.D0.create_DC2();
          let _342_v13;
          _342_v13 = _dafny.Map.Empty.slice().updateUnsafe(!((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]) || (_325_v0),_325_v0);
          let _343_v14;
          _343_v14 = _dafny.Set.fromElements(_325_v0);
          let _344_v16;
          _344_v16 = _dafny.Seq.UnicodeFromString("cxo");
          let _345_v17;
          _345_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_344_v16);
          let _346_v18;
          let _nw57 = Array((new BigNumber(22)).toNumber());
          _nw57[(_dafny.ZERO).toNumber()] = p0;
          _nw57[(_dafny.ONE).toNumber()] = p0;
          _nw57[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_325_v0, (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]), _dafny.Seq.update(p0, _module.__default.safeIndex((_this).f4, new BigNumber((p0).length)), _325_v0));
          _nw57[(new BigNumber(3)).toNumber()] = p0;
          _nw57[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_325_v0);
          _nw57[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], false);
          _nw57[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex(_this.f5, new BigNumber((p0).length)), (_this).fm1(_325_v0, (_dafny.ZERO).minus((_this).f4), _343_v14, _325_v0, globalState));
          _nw57[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p0, p0);
          _nw57[(new BigNumber(8)).toNumber()] = p0;
          _nw57[(new BigNumber(9)).toNumber()] = _dafny.Seq.of((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]);
          _nw57[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((p0)[_module.__default.safeIndex(_this.f5, new BigNumber((p0).length))]);
          _nw57[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_325_v0);
          _nw57[(new BigNumber(12)).toNumber()] = p0;
          _nw57[(new BigNumber(13)).toNumber()] = _module.__default.fm5(function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-360), new BigNumber(250))) {
              let _347_v15 = _compr_9;
              if (((new BigNumber(-360)).isLessThanOrEqualTo(_347_v15)) && ((_347_v15).isLessThan(new BigNumber(250)))) {
                _coll9.push([(_347_v15).plus(_this.f5),_dafny.Seq.UnicodeFromString("dfrfshin")]);
              }
            }
            return _coll9;
          }(), _this.f5, _325_v0, (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], globalState);
          _nw57[(new BigNumber(14)).toNumber()] = p0;
          _nw57[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm5(_345_v17, _this.f5, (_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))], _325_v0, globalState), p0);
          _nw57[(new BigNumber(16)).toNumber()] = p0;
          _nw57[(new BigNumber(17)).toNumber()] = p0;
          _nw57[(new BigNumber(18)).toNumber()] = p0;
          _nw57[(new BigNumber(19)).toNumber()] = p0;
          _nw57[(new BigNumber(20)).toNumber()] = p0;
          _nw57[(new BigNumber(21)).toNumber()] = p0;
          _346_v18 = _nw57;
          let _index59 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_346_v18).length));
          (_346_v18)[_index59] = _dafny.Seq.Concat(p0, p0);
          let _348_v19;
          _348_v19 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f4),_this.f5);
          let _349_v20;
          _349_v20 = _module.D0.create_DC2();
          let _350_v21;
          _350_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]);
          let _index60 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_341_v12).length));
          let _index61 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_346_v18).length));
          let _rhs51 = (_this.f5).isLessThanOrEqualTo(new BigNumber((_348_v19).length));
          let _rhs52 = _349_v20;
          let _rhs53 = _342_v13;
          let _rhs54 = p0;
          let _rhs55 = (_dafny.ZERO).minus((_module.__default.safeModuloInt((_330_v5).dtor_cf0, _this.f5)).plus(_module.__default.safeModuloInt(new BigNumber((_350_v21).length), (_this).f4)));
          let _lhs40 = globalState;
          let _lhs41 = _341_v12;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_341_v12).length));
          let _lhs43 = _346_v18;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_346_v18).length));
          _lhs40.f2 = _rhs51;
          _lhs41[_lhs42] = _rhs52;
          _342_v13 = _rhs53;
          _lhs43[_lhs44] = _rhs54;
          r1 = _rhs55;
          _348_v19 = (_348_v19).update(_module.__default.safeModuloInt(_this.f5, new BigNumber((function () {
            let _coll10 = new _dafny.Set();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(814), new BigNumber(257))) {
              let _351_v22 = _compr_10;
              if (((new BigNumber(814)).isLessThanOrEqualTo(_351_v22)) && ((_351_v22).isLessThan(new BigNumber(257)))) {
                _coll10.add((_351_v22).plus(_this.f5));
              }
            }
            return _coll10;
          }()).length)), _module.__default.safeDivisionInt(_this.f5, (_this).f4));
          let _352_v23;
          _352_v23 = _module.D1.create_DC3(!(_325_v0) || ((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]));
          _352_v23 = _352_v23;
          (_this).f5 = (_module.__default.fm3(_325_v0, globalState)).minus((_this).f4);
        } else {
          let _353___mcc_h2 = (_source6).cf0;
          let _354_cf0 = _353___mcc_h2;
          _330_v5 = _330_v5;
          let _355_v24;
          _355_v24 = _dafny.Set.fromElements((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]);
          (_this).f5 = _module.__default.fm3(!(_355_v24).contains((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]), globalState);
          let _356_v25;
          _356_v25 = _dafny.Map.Empty.slice().updateUnsafe((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))],(_this).f4);
          _356_v25 = (_356_v25).update(!((_329_v4)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_329_v4).length))]), (_this).f4);
          let _357_v26;
          _357_v26 = _dafny.MultiSet.fromElements(_this.f5);
          let _358_v27;
          _358_v27 = _dafny.Seq.UnicodeFromString("x");
          r0 = ((_dafny.ZERO).minus(((_325_v0) ? (_this.f5) : (new BigNumber((_357_v26).cardinality()))))).plus(new BigNumber((_358_v27).length));
        }
      } else {
        let _359_v28;
        _359_v28 = new _dafny.CodePoint('e'.codePointAt(0));
        let _360_v29;
        let _nw58 = Array((new BigNumber(3)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = _359_v28;
        _nw58[(_dafny.ONE).toNumber()] = _359_v28;
        _nw58[(new BigNumber(2)).toNumber()] = _359_v28;
        _360_v29 = _nw58;
        let _rhs56 = _325_v0;
        let _rhs57 = ((_this).f4).isEqualTo((_this).f4);
        let _rhs58 = _module.__default.fm3(_dafny.Seq.IsProperPrefixOf(p0, p0), globalState);
        let _rhs59 = _360_v29;
        let _rhs60 = (_this).f4;
        let _lhs45 = globalState;
        _325_v0 = _rhs56;
        _lhs45.f2 = _rhs57;
        r0 = _rhs58;
        _360_v29 = _rhs59;
        r1 = _rhs60;
        let _361_v30;
        _361_v30 = _dafny.MultiSet.fromElements(_359_v28);
        let _362_v31;
        let _nw59 = new _module.C0();
        _nw59.__ctor(_361_v30);
        _362_v31 = _nw59;
        let _363_v32;
        _363_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f5);
        _363_v32 = (_363_v32).update(_325_v0, _this.f5);
        let _364_v33;
        let _nw60 = Array((new BigNumber(24)).toNumber()).fill([]);
        _364_v33 = _nw60;
        _364_v33 = _364_v33;
        (globalState).f2 = _325_v0;
      }
      let _365_v34;
      _365_v34 = new _dafny.CodePoint('b'.codePointAt(0));
      let _366_v35;
      _366_v35 = _dafny.Seq.of(_365_v34);
      let _367_v36;
      let _nw61 = new _module.C0();
      _nw61.__ctor(_dafny.MultiSet.FromArray(_366_v35));
      _367_v36 = _nw61;
      let _368_v37;
      let _init4 = ((_369_v34) => function (_370_i0) {
        return _369_v34;
      })(_365_v34);
      let _nw62 = Array((new BigNumber(21)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw62.length); _i0_4++) {
        _nw62[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _368_v37 = _nw62;
      let _371_v38;
      _371_v38 = _dafny.Map.Empty.slice().updateUnsafe(_365_v34,_368_v37);
      let _372_v39;
      _372_v39 = _dafny.Seq.of((((_371_v38).contains(_365_v34)) ? ((_371_v38).get(_365_v34)) : (_368_v37)), _368_v37, _368_v37);
      if (_dafny.areEqual(((_325_v0) ? (_372_v39) : (_372_v39)), _dafny.Seq.of(_368_v37, _368_v37))) {
        let _373_v40;
        let _nw63 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _373_v40 = _nw63;
        let _374_v41;
        _374_v41 = _dafny.Map.Empty.slice().updateUnsafe(_325_v0,_373_v40);
        _374_v41 = (_374_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(_325_v0,_373_v40));
        let _375_v42;
        _375_v42 = _dafny.Seq.of(new BigNumber((p0).length), new BigNumber((_dafny.Seq.UnicodeFromString("dgn")).length));
        (globalState).f2 = (_module.__default.safeDivisionInt(_this.f5, new BigNumber((_375_v42).length))).isLessThan((_this.f5).minus(_this.f5));
        let _376_v43;
        let _nw64 = new _module.C0();
        _nw64.__ctor(_dafny.MultiSet.FromArray(_366_v35));
        _376_v43 = _nw64;
        let _377_v44;
        _377_v44 = _module.D0.create_DC1((_this.f5).plus(_this.f5), _this.f5);
        let _source7 = _377_v44;
        if (_source7.is_DC1) {
          let _378___mcc_h3 = (_source7).cf1;
          let _379___mcc_h4 = (_source7).cf2;
          let _380_cf2 = _379___mcc_h4;
          let _381_cf1 = _378___mcc_h3;
          let _index62 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_373_v40).length));
          (_373_v40)[_index62] = _this.f5;
          let _index63 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_373_v40).length));
          (_373_v40)[_index63] = _380_cf2;
          let _382_v45;
          let _383_v46;
          let _384_v47;
          let _385_v48;
          let _out15;
          let _out16;
          let _out17;
          let _out18;
          let _outcollector6 = (_this).m4(_325_v0, (_this).f4, _module.__default.fm6(_325_v0, _this.f5, _325_v0, globalState), globalState);
          _out15 = _outcollector6[0];
          _out16 = _outcollector6[1];
          _out17 = _outcollector6[2];
          _out18 = _outcollector6[3];
          _382_v45 = _out15;
          _383_v46 = _out16;
          _384_v47 = _out17;
          _385_v48 = _out18;
          _374_v41 = _374_v41;
          let _386_v49;
          let _init5 = ((_387_cf1) => function (_388_i1) {
            return (_388_i1).multipliedBy(_387_cf1);
          })(_381_cf1);
          let _nw65 = Array((new BigNumber(27)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw65.length); _i0_5++) {
            _nw65[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _386_v49 = _nw65;
          _386_v49 = _386_v49;
        } else if (_source7.is_DC2) {
          let _389_v50;
          _389_v50 = _dafny.Map.Empty.slice().updateUnsafe((_325_v0) === (_325_v0),_325_v0);
          _389_v50 = (_389_v50).update(!(_325_v0), _325_v0);
          _377_v44 = _377_v44;
          _325_v0 = _325_v0;
          _389_v50 = (_389_v50).update(_325_v0, false);
        } else {
          let _390___mcc_h5 = (_source7).cf0;
          let _391_cf0 = _390___mcc_h5;
          (globalState).f2 = _325_v0;
          let _392_v51;
          _392_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f4);
          let _393_v52;
          _393_v52 = _module.D1.create_DC4(_325_v0);
          let _pat_let_tv10 = _325_v0;
          let _394_v53;
          _394_v53 = _dafny.Set.fromElements(_module.__default.fm7(_325_v0, _392_v51, _365_v34, _325_v0, globalState), _393_v52, _module.__default.fm7(_325_v0, _392_v51, _365_v34, _325_v0, globalState), _393_v52, function (_pat_let6_0) {
            return function (_395_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_396_dt__update_hcf4_h0) {
                  return _module.D1.create_DC4(_396_dt__update_hcf4_h0);
                }(_pat_let7_0);
              }(_pat_let_tv10);
            }(_pat_let6_0);
          }(_393_v52));
          _394_v53 = function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_module.__default.fm8(_365_v34, _325_v0, _391_cf0, _325_v0, globalState)).Keys.Elements) {
              let _397_v54 = _compr_11;
              if ((_module.__default.fm8(_365_v34, _325_v0, _391_cf0, _325_v0, globalState)).contains(_397_v54)) {
                _coll11.add(_397_v54);
              }
            }
            return _coll11;
          }();
          let _398_v55;
          let _init6 = ((_399_v0) => function (_400_i2) {
            return _399_v0;
          })(_325_v0);
          let _nw66 = Array((new BigNumber(10)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw66.length); _i0_6++) {
            _nw66[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _398_v55 = _nw66;
          let _401_v56;
          _401_v56 = _dafny.Map.Empty.slice().updateUnsafe(_391_cf0,_365_v34);
          let _rhs61 = (((_401_v56).contains((_this).f4)) ? ((_401_v56).get((_this).f4)) : (_365_v34));
          let _rhs62 = _365_v34;
          let _rhs63 = !(!(_325_v0) || (_325_v0)) || (!(_325_v0));
          let _rhs64 = _398_v55;
          _365_v34 = _rhs61;
          _365_v34 = _rhs62;
          _325_v0 = _rhs63;
          _398_v55 = _rhs64;
          (globalState).f2 = _dafny.Seq.contains(_366_v35, new _dafny.CodePoint('u'.codePointAt(0)));
        }
        r0 = (_375_v42)[_module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements(_325_v0, _325_v0)).cardinality())).multipliedBy((_this).f4)), new BigNumber((_375_v42).length))];
      } else {
        let _402_v57;
        let _init7 = ((_403_v0) => function (_404_i3) {
          return !(_403_v0) || (_403_v0);
        })(_325_v0);
        let _nw67 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw67.length); _i0_7++) {
          _nw67[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _402_v57 = _nw67;
        let _405_v58;
        let _nw68 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _405_v58 = _nw68;
        let _406_v59;
        _406_v59 = _dafny.Set.fromElements(_325_v0);
        let _rhs65 = _402_v57;
        let _rhs66 = !((_406_v59).Difference(_406_v59)).equals(_406_v59);
        let _rhs67 = _405_v58;
        let _rhs68 = (_this).fm2(_325_v0, _this.f5, _325_v0, globalState);
        let _lhs46 = globalState;
        _402_v57 = _rhs65;
        _lhs46.f2 = _rhs66;
        _405_v58 = _rhs67;
        _325_v0 = _rhs68;
        if (true) {
          let _407_v60;
          _407_v60 = _dafny.Seq.of((_367_v36).f7);
          let _408_v61;
          let _nw69 = new _module.C0();
          _nw69.__ctor((_407_v60)[_module.__default.safeIndex(_this.f5, new BigNumber((_407_v60).length))]);
          _408_v61 = _nw69;
          let _409_v62;
          let _nw70 = Array((new BigNumber(26)).toNumber()).fill([]);
          _409_v62 = _nw70;
          let _nw71 = Array((new BigNumber(22)).toNumber()).fill([]);
          _409_v62 = _nw71;
          let _410_v63;
          let _nw72 = new _module.C0();
          _nw72.__ctor((_367_v36).f7);
          _410_v63 = _nw72;
          r0 = ((_this).f4).plus((_this).f4);
          let _411_v64;
          _411_v64 = _module.D0.create_DC0(_this.f5);
          r0 = (_411_v64).dtor_cf0;
        } else {
          let _412_v65;
          let _nw73 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _412_v65 = _nw73;
          let _index64 = _module.__default.safeIndex(new BigNumber(826), new BigNumber((_412_v65).length));
          (_412_v65)[_index64] = _this.f5;
          let _index65 = _module.__default.safeIndex(new BigNumber(826), new BigNumber((_412_v65).length));
          (_412_v65)[_index65] = (_dafny.ZERO).minus((_this).f4);
          (globalState).f2 = !(!(_325_v0));
          _402_v57 = _402_v57;
          _412_v65 = _412_v65;
          (globalState).f2 = _dafny.areEqual(_366_v35, _366_v35);
        }
        let _413_v66;
        _413_v66 = _dafny.MultiSet.fromElements(_325_v0);
        let _414_v67;
        _414_v67 = _dafny.Seq.of(_366_v35);
        let _415_v68;
        _415_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm6(_325_v0, (_this).f4, _325_v0, globalState)).length),false);
        let _416_v69;
        _416_v69 = _dafny.Set.fromElements(_415_v68, _415_v68, _415_v68);
        let _417_v70;
        _417_v70 = _dafny.Map.Empty.slice().updateUnsafe((_413_v66).IsDisjointFrom(_413_v66),!(_dafny.Seq.IsProperPrefixOf(_module.__default.fm6(!(_325_v0), _this.f5, _325_v0, globalState), (_414_v67)[_module.__default.safeIndex(new BigNumber((_416_v69).length), new BigNumber((_414_v67).length))])));
        _417_v70 = (_417_v70).update(_325_v0, _325_v0);
        let _418_v71;
        _418_v71 = _dafny.MultiSet.fromElements(p0);
        let _rhs69 = (_325_v0) && (_325_v0);
        let _rhs70 = ((_dafny.ZERO).minus((((_418_v71).contains(p0)) ? ((_418_v71).get(p0)) : (_this.f5)))).minus(new BigNumber(215));
        let _lhs47 = globalState;
        _lhs47.f2 = _rhs69;
        r0 = _rhs70;
        (_this).f5 = _module.__default.safeModuloInt((_this).f4, new BigNumber((((_325_v0) ? (_366_v35) : (_dafny.Seq.update(_366_v35, _module.__default.safeIndex((_this).f4, new BigNumber((_366_v35).length)), _365_v34)))).length));
      }
      let _419_v72;
      _419_v72 = _dafny.MultiSet.fromElements(_325_v0);
      _419_v72 = _419_v72;
      let _420_v73;
      _420_v73 = _dafny.Map.Empty.slice().updateUnsafe(((_325_v0) ? (_366_v35) : (_366_v35)),_325_v0);
      _420_v73 = (_420_v73).update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), function (_421_i4) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("pu")), _325_v0);
      let _422_v74;
      _422_v74 = _dafny.Map.Empty.slice().updateUnsafe(_325_v0,_365_v34);
      r0 = _module.__default.safeDivisionInt((new BigNumber((_422_v74).length)).minus((_this).f4), (_this).f4);
      r1 = (_this).f4;
      return [r0, r1];
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.MultiSet.Empty;
      let r3 = _dafny.Seq.of();
      let _423_v0;
      _423_v0 = _dafny.Seq.of(p1);
      let _424_v1;
      _424_v1 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_423_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), ((_425_p1) => function (_426_i0) {
        return _425_p1;
      })(p1)))).length));
      let _427_v2;
      _427_v2 = _dafny.Seq.of((_424_v1).IsSubsetOf(_424_v1));
      (globalState).f2 = (_427_v2)[_module.__default.safeIndex(_this.f5, new BigNumber((_427_v2).length))];
      let _428_v3;
      _428_v3 = _module.D1.create_DC3(p0);
      let _429_v4;
      _429_v4 = _dafny.Set.fromElements(_module.D1.create_DC5(_428_v3));
      let _430_v5;
      _430_v5 = _module.D1.create_DC5(_428_v3);
      _429_v4 = _dafny.Set.fromElements(_430_v5, _430_v5);
      let _431_v6;
      _431_v6 = _dafny.MultiSet.fromElements(p0);
      (globalState).f2 = (_this).fm1(p0, _this.f5, _module.__default.fm9(p1, p0, new BigNumber((_431_v6).cardinality()), globalState), (_this.f5).isLessThanOrEqualTo(p1), globalState);
      if (p0) {
        let _432_v7;
        _432_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_423_v0);
        let _433_v8;
        _433_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _434_v9;
        _434_v9 = _dafny.Set.fromElements(p0);
        let _435_v10;
        _435_v10 = _module.D1.create_DC4(p0);
        let _436_v11;
        _436_v11 = _dafny.Seq.of(_dafny.Set.fromElements((_435_v10).dtor_cf4, p0), _434_v9, _434_v9);
        let _437_v12;
        _437_v12 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), ((_438_p0) => function (_439_i1) {
          return _module.D1.create_DC4(_438_p0);
        })(p0)));
        let _440_v13;
        _440_v13 = _dafny.Seq.of(_435_v10, _435_v10);
        let _441_v14;
        let _nw74 = Array((new BigNumber(21)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = p1;
        _nw74[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((p1).plus((_dafny.ZERO).minus((_423_v0)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_423_v0).length))])));
        _nw74[(new BigNumber(2)).toNumber()] = (new BigNumber((_432_v7).length)).plus(_this.f5);
        _nw74[(new BigNumber(3)).toNumber()] = new BigNumber((_433_v8).length);
        _nw74[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_434_v9).Intersect((_436_v11)[_module.__default.safeIndex(p1, new BigNumber((_436_v11).length))])).length));
        _nw74[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw74[(new BigNumber(6)).toNumber()] = (_this).f4;
        _nw74[(new BigNumber(7)).toNumber()] = p1;
        _nw74[(new BigNumber(8)).toNumber()] = (((_437_v12).contains(_440_v13)) ? ((_437_v12).get(_440_v13)) : ((_this).f4));
        _nw74[(new BigNumber(9)).toNumber()] = _this.f5;
        _nw74[(new BigNumber(10)).toNumber()] = p1;
        _nw74[(new BigNumber(11)).toNumber()] = _this.f5;
        _nw74[(new BigNumber(12)).toNumber()] = _this.f5;
        _nw74[(new BigNumber(13)).toNumber()] = new BigNumber(-836);
        _nw74[(new BigNumber(14)).toNumber()] = _module.__default.fm3((_this).fm1(!(p0), _module.__default.fm3(p0, globalState), _434_v9, p0, globalState), globalState);
        _nw74[(new BigNumber(15)).toNumber()] = new BigNumber(254);
        _nw74[(new BigNumber(16)).toNumber()] = (_this).f4;
        _nw74[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_this.f5).multipliedBy(new BigNumber((_dafny.Seq.of(p0)).length)));
        _nw74[(new BigNumber(18)).toNumber()] = p1;
        _nw74[(new BigNumber(19)).toNumber()] = (_423_v0)[_module.__default.safeIndex(_this.f5, new BigNumber((_423_v0).length))];
        _nw74[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_423_v0, _423_v0)).length);
        _441_v14 = _nw74;
        let _index66 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_441_v14).length));
        (_441_v14)[_index66] = (p1).minus(_this.f5);
        let _442_v15;
        let _init8 = ((_443_p0) => function (_444_i2) {
          return ((_443_p0) ? (new _dafny.CodePoint('v'.codePointAt(0))) : (new _dafny.CodePoint('f'.codePointAt(0))));
        })(p0);
        let _nw75 = Array((new BigNumber(15)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
          _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _442_v15 = _nw75;
        let _445_v16;
        _445_v16 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index67 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_442_v15).length));
        (_442_v15)[_index67] = _445_v16;
        let _446_v17;
        _446_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0);
        let _index68 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_441_v14).length));
        let _index69 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_442_v15).length));
        let _rhs71 = (_dafny.ZERO).minus((_this).f4);
        let _rhs72 = _module.__default.fm4(p0, new BigNumber(-652), !(p0) || (false), _module.__default.fm3(!(false), globalState), globalState);
        let _rhs73 = !(p0) || (p0);
        let _rhs74 = !(p0);
        let _rhs75 = ((p0) ? (_446_v17) : ((_446_v17).Merge(_module.__default.fm10(_this.f5, p0, (_this).f4, p1, globalState))));
        let _lhs48 = _441_v14;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_441_v14).length));
        let _lhs50 = _442_v15;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_442_v15).length));
        let _lhs52 = globalState;
        let _lhs53 = globalState;
        _lhs48[_lhs49] = _rhs71;
        _lhs50[_lhs51] = _rhs72;
        _lhs52.f2 = _rhs73;
        _lhs53.f2 = _rhs74;
        _446_v17 = _rhs75;
        let _447_v18;
        _447_v18 = _module.D0.create_DC1(_this.f5, (_441_v14)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_441_v14).length))]);
        _447_v18 = _447_v18;
        r0 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(-676), p1), (_this).f4);
        let _index70 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_441_v14).length));
        (_441_v14)[_index70] = (_this).f4;
        (globalState).f2 = p0;
      } else {
        (globalState).f2 = (_module.__default.fm11(new BigNumber(-981), p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f5)), globalState)).IsSubsetOf((_424_v1).Difference(_424_v1));
        (_this).f5 = (((_this).f4).minus((_dafny.ZERO).minus(p1))).minus(_this.f5);
        (globalState).f2 = (p1).isLessThan(new BigNumber(194));
        let _448_v19;
        _448_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        if ((_this).fm2(p0, _this.f5, (((_448_v19).contains(p0)) ? ((_448_v19).get(p0)) : (p0)), globalState)) {
          let _449_v20;
          let _nw76 = Array((new BigNumber(10)).toNumber());
          _449_v20 = _nw76;
          _449_v20 = _449_v20;
          let _450_v21;
          let _nw77 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _450_v21 = _nw77;
          let _rhs76 = p1;
          let _rhs77 = _450_v21;
          r0 = _rhs76;
          _450_v21 = _rhs77;
          r0 = p1;
          _423_v0 = _423_v0;
          let _451_v22;
          let _init9 = ((_452_p1) => function (_453_i3) {
            return (_453_i3).multipliedBy(_452_p1);
          })(p1);
          let _nw78 = Array((new BigNumber(18)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
            _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _451_v22 = _nw78;
          let _rhs78 = (_this).f4;
          let _rhs79 = _451_v22;
          let _rhs80 = p0;
          let _rhs81 = (_dafny.ZERO).minus((((_424_v1).IsDisjointFrom(_424_v1)) ? (new BigNumber(-201)) : ((_dafny.ZERO).minus(new BigNumber(((_431_v6).Intersect(_431_v6)).cardinality())))));
          let _lhs54 = _this;
          let _lhs55 = globalState;
          let _lhs56 = _this;
          _lhs54.f5 = _rhs78;
          _451_v22 = _rhs79;
          _lhs55.f2 = _rhs80;
          _lhs56.f5 = _rhs81;
        } else {
          r0 = new BigNumber((_dafny.Seq.Concat(_427_v2, _427_v2)).length);
          let _454_v23;
          let _nw79 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _454_v23 = _nw79;
          let _index71 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_454_v23).length));
          (_454_v23)[_index71] = _423_v0;
          let _index72 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_454_v23).length));
          (_454_v23)[_index72] = _dafny.Seq.update(_423_v0, _module.__default.safeIndex(_module.__default.fm3(p0, globalState), new BigNumber((_423_v0).length)), p1);
          let _455_v24;
          _455_v24 = _module.D0.create_DC0(_this.f5);
          (_this).f5 = (_455_v24).dtor_cf0;
          let _456_v25;
          let _init10 = ((_457_p0) => function (_458_i4) {
            return _457_p0;
          })(p0);
          let _nw80 = Array((new BigNumber(3)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw80.length); _i0_10++) {
            _nw80[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _456_v25 = _nw80;
          let _459_v26;
          _459_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_456_v25, _456_v25),p1);
          let _460_v27;
          _460_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,_456_v25);
          let _461_v28;
          _461_v28 = _dafny.Seq.of((((_460_v27).contains(false)) ? ((_460_v27).get(false)) : (_456_v25)), _456_v25, _456_v25, _456_v25, _456_v25);
          let _462_v29;
          _462_v29 = _dafny.Seq.of(_461_v28);
          (globalState).f2 = (p1).isLessThan((((_459_v26).contains((_462_v29)[_module.__default.safeIndex((_this).f4, new BigNumber((_462_v29).length))])) ? ((_459_v26).get((_462_v29)[_module.__default.safeIndex((_this).f4, new BigNumber((_462_v29).length))])) : (new BigNumber((p2).length))));
          let _463_v30;
          _463_v30 = _dafny.Set.fromElements(p0);
          let _index73 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_456_v25).length));
          (_456_v25)[_index73] = (_this).fm1(!((_this).fm2(false, p1, p0, globalState)), (_dafny.ZERO).minus(p1), _463_v30, p0, globalState);
          let _index74 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_456_v25).length));
          (_456_v25)[_index74] = p0;
        }
        r0 = _module.__default.safeDivisionInt((_this).f4, (_this).f4);
      }
      (globalState).f2 = p0;
      let _464_v31;
      _464_v31 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("k"));
      _464_v31 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nxhuqgf"), p2);
      r0 = (_this).f4;
      r1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_465_i5) {
        return (_this).f4;
      });
      r2 = _424_v1;
      r3 = _427_v2;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f5 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
      this._f6 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f6, f4, f5) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _466_v0;
      let _nw81 = Array((new BigNumber(15)).toNumber()).fill(false);
      _466_v0 = _nw81;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_466_v0).length))) {
        let _467_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_467_i0)) && ((_467_i0).isLessThan(new BigNumber((_466_v0).length))))) {
          (_466_v0)[(_467_i0)] = (_this).f6;
        }
      }
      let _468_v1;
      _468_v1 = new _dafny.CodePoint('o'.codePointAt(0));
      _468_v1 = _468_v1;
      (_this).f5 = _this.f5;
      let _rhs82 = _this.f5;
      let _rhs83 = (_this).f6;
      let _lhs57 = _this;
      let _lhs58 = globalState;
      _lhs57.f5 = _rhs82;
      _lhs58.f2 = _rhs83;
      let _469_v2;
      let _nw82 = new _module.C0();
      _nw82.__ctor(_dafny.MultiSet.fromElements(_468_v1));
      _469_v2 = _nw82;
      r0 = (_this).f6;
      let _470_v3;
      _470_v3 = _module.D1.create_DC4((_this).f6);
      r0 = !((_470_v3).dtor_cf4);
      let _471_v4;
      _471_v4 = _dafny.Seq.of(_module.__default.fm12((_this).f6, (_this).f6, (_this).f6, globalState));
      r1 = (_471_v4)[_module.__default.safeIndex((_this).f4, new BigNumber((_471_v4).length))];
      r2 = (_this).f6;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      if (false) {
        r0 = (_this).f6;
        if (((_this).f6) && ((_this).f6)) {
          r0 = !((_this).f6);
          let _472_v0;
          _472_v0 = _dafny.Seq.of((_this).f6);
          let _473_v1;
          _473_v1 = _dafny.MultiSet.fromElements(_472_v0);
          let _474_v2;
          _474_v2 = _module.D2.create_DC6(_472_v0);
          let _475_v3;
          _475_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
          let _476_v4;
          _476_v4 = _dafny.Seq.UnicodeFromString("quepdvot");
          let _477_v5;
          _477_v5 = _module.D3.create_DC8(_476_v4);
          let _478_v6;
          _478_v6 = _dafny.Set.fromElements((_this).f6);
          let _479_v7;
          _479_v7 = new _dafny.CodePoint('v'.codePointAt(0));
          let _480_v8;
          let _nw83 = Array((new BigNumber(18)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _this.f5;
          _nw83[(_dafny.ONE).toNumber()] = (((_473_v1).contains(_dafny.Seq.of(!((_this).f6)))) ? ((_473_v1).get(_dafny.Seq.of(!((_this).f6)))) : (_module.__default.fm3((_this).f6, globalState)));
          _nw83[(new BigNumber(2)).toNumber()] = _this.f5;
          _nw83[(new BigNumber(3)).toNumber()] = new BigNumber(((_474_v2).dtor_cf6).length);
          _nw83[(new BigNumber(4)).toNumber()] = (_this).f4;
          _nw83[(new BigNumber(5)).toNumber()] = new BigNumber((_475_v3).length);
          _nw83[(new BigNumber(6)).toNumber()] = (_this).f4;
          _nw83[(new BigNumber(7)).toNumber()] = p0;
          _nw83[(new BigNumber(8)).toNumber()] = _this.f5;
          _nw83[(new BigNumber(9)).toNumber()] = _this.f5;
          _nw83[(new BigNumber(10)).toNumber()] = new BigNumber(-265);
          _nw83[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update((_477_v5).dtor_cf7, _module.__default.safeIndex(new BigNumber((_478_v6).length), new BigNumber(((_477_v5).dtor_cf7).length)), _479_v7)).length);
          _nw83[(new BigNumber(12)).toNumber()] = (_this).f4;
          _nw83[(new BigNumber(13)).toNumber()] = p0;
          _nw83[(new BigNumber(14)).toNumber()] = (_this).f4;
          _nw83[(new BigNumber(15)).toNumber()] = p0;
          _nw83[(new BigNumber(16)).toNumber()] = _this.f5;
          _nw83[(new BigNumber(17)).toNumber()] = (_this).f4;
          _480_v8 = _nw83;
          let _481_v9;
          let _nw84 = Array((new BigNumber(14)).toNumber()).fill(false);
          _481_v9 = _nw84;
          let _482_v10;
          _482_v10 = _dafny.Map.Empty.slice().updateUnsafe(_480_v8,_481_v9);
          let _483_v11;
          _483_v11 = _module.D1.create_DC3((_this).f6);
          let _484_v12;
          _484_v12 = _dafny.Map.Empty.slice().updateUnsafe(_482_v10,!(((_this).f6) === (_module.__default.fm12((_483_v11).dtor_cf3, (_this).f6, (_this).f6, globalState))));
          let _485_v13;
          _485_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
          _484_v12 = (_484_v12).update((_482_v10).update(_480_v8, _481_v9), (((_485_v13).contains((_this).f6)) ? ((_485_v13).get((_this).f6)) : ((_this).f6)));
          _475_v3 = (_475_v3).update(_this.f5, (_this).f6);
          (_this).f5 = p0;
          let _486_v14;
          _486_v14 = _dafny.MultiSet.fromElements(_479_v7, _479_v7);
          let _487_v15;
          let _nw85 = new _module.C0();
          _nw85.__ctor((_486_v14).Difference(_dafny.MultiSet.fromElements(_479_v7)));
          _487_v15 = _nw85;
        } else {
          let _488_v16;
          _488_v16 = new _dafny.CodePoint('v'.codePointAt(0));
          let _489_v17;
          _489_v17 = _dafny.MultiSet.fromElements(_488_v16);
          let _490_v18;
          let _nw86 = new _module.C0();
          _nw86.__ctor(_489_v17);
          _490_v18 = _nw86;
          (globalState).f2 = (_this).f6;
          _488_v16 = _488_v16;
          let _491_v19;
          _491_v19 = _dafny.Seq.UnicodeFromString("ngmxo");
          let _492_v20;
          _492_v20 = _module.D3.create_DC8(_491_v19);
          _491_v19 = (_492_v20).dtor_cf7;
          _491_v19 = _module.__default.fm6((_this).f6, _this.f5, !((_this).f6), globalState);
        }
        (_this).f5 = (_dafny.ZERO).minus(_module.__default.fm3((_this).f6, globalState));
        let _493_v21;
        _493_v21 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, (_this).f6, (_this).f6, (_this).f6);
        let _494_v22;
        _494_v22 = _dafny.Seq.of((_this).f4);
        let _495_v23;
        _495_v23 = new _dafny.CodePoint('t'.codePointAt(0));
        let _496_v24;
        _496_v24 = _dafny.MultiSet.fromElements(_495_v23, _495_v23, _495_v23);
        let _497_v25;
        let _nw87 = new _module.C0();
        _nw87.__ctor(_496_v24);
        _497_v25 = _nw87;
        let _498_v26;
        _498_v26 = _module.D0.create_DC1((_this).f4, new BigNumber(469));
        let _499_v27;
        _499_v27 = _module.D0.create_DC0(p0);
        let _500_v28;
        _500_v28 = _module.D1.create_DC4((_this).f6);
        let _501_v29;
        _501_v29 = _module.D3.create_DC9(p0, (_this).f6, _this.f5, (_500_v28).dtor_cf4);
        let _502_v30;
        let _nw88 = Array((new BigNumber(19)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = new BigNumber(356);
        _nw88[(_dafny.ONE).toNumber()] = _this.f5;
        _nw88[(new BigNumber(2)).toNumber()] = new BigNumber((_493_v21).cardinality());
        _nw88[(new BigNumber(3)).toNumber()] = _module.__default.fm3((_this).f6, globalState);
        _nw88[(new BigNumber(4)).toNumber()] = _this.f5;
        _nw88[(new BigNumber(5)).toNumber()] = (_this).f4;
        _nw88[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_494_v22)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_497_v25)).length), new BigNumber((_494_v22).length))]);
        _nw88[(new BigNumber(7)).toNumber()] = _this.f5;
        _nw88[(new BigNumber(8)).toNumber()] = ((_498_v26).dtor_cf1).plus(_this.f5);
        _nw88[(new BigNumber(9)).toNumber()] = new BigNumber(195);
        _nw88[(new BigNumber(10)).toNumber()] = (_this).f4;
        _nw88[(new BigNumber(11)).toNumber()] = p0;
        _nw88[(new BigNumber(12)).toNumber()] = (_498_v26).dtor_cf2;
        _nw88[(new BigNumber(13)).toNumber()] = (_499_v27).dtor_cf0;
        _nw88[(new BigNumber(14)).toNumber()] = (_494_v22)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_494_v22).length))];
        _nw88[(new BigNumber(15)).toNumber()] = _this.f5;
        _nw88[(new BigNumber(16)).toNumber()] = (_501_v29).dtor_cf10;
        _nw88[(new BigNumber(17)).toNumber()] = _this.f5;
        _nw88[(new BigNumber(18)).toNumber()] = (_this).f4;
        _502_v30 = _nw88;
        let _index75 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_502_v30).length));
        (_502_v30)[_index75] = (_this).f4;
        let _index76 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_502_v30).length));
        (_502_v30)[_index76] = (_this).f4;
        let _503_v31;
        _503_v31 = _dafny.Seq.of((_this).f6, (_this).f6);
        let _504_v32;
        _504_v32 = _dafny.Map.Empty.slice().updateUnsafe(_502_v30,_503_v31);
        _493_v21 = _dafny.MultiSet.FromArray((((_504_v32).contains(_502_v30)) ? ((_504_v32).get(_502_v30)) : (_503_v31)));
      } else {
        (globalState).f2 = (_this).f6;
        let _505_v33;
        _505_v33 = new _dafny.CodePoint('e'.codePointAt(0));
        let _506_v34;
        _506_v34 = _dafny.MultiSet.fromElements(_505_v33);
        let _507_v35;
        let _nw89 = new _module.C0();
        _nw89.__ctor(_506_v34);
        _507_v35 = _nw89;
        let _508_v36;
        _508_v36 = _module.D1.create_DC4((_this).f6);
        let _509_v37;
        let _nw90 = Array((new BigNumber(10)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = _508_v36;
        _nw90[(_dafny.ONE).toNumber()] = _508_v36;
        _nw90[(new BigNumber(2)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(3)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(4)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(5)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(6)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(7)).toNumber()] = function (_pat_let8_0) {
          return function (_510_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_511_dt__update_hcf4_h0) {
                return _module.D1.create_DC4(_511_dt__update_hcf4_h0);
              }(_pat_let9_0);
            }((_this).f6);
          }(_pat_let8_0);
        }(_508_v36);
        _nw90[(new BigNumber(8)).toNumber()] = _508_v36;
        _nw90[(new BigNumber(9)).toNumber()] = _module.D1.create_DC4((_this).f6);
        _509_v37 = _nw90;
        let _512_v38;
        let _nw91 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _512_v38 = _nw91;
        let _513_v39;
        _513_v39 = _dafny.Seq.of(_512_v38);
        let _rhs84 = _507_v35;
        let _rhs85 = _509_v37;
        let _rhs86 = (_513_v39)[_module.__default.safeIndex(((_this).f4).minus(_module.__default.fm3((_this).f6, globalState)), new BigNumber((_513_v39).length))];
        let _rhs87 = ((_this).f6) && ((_this).f6);
        let _rhs88 = ((_dafny.ZERO).minus((p0).multipliedBy((_this).f4))).minus((_this).f4);
        let _lhs59 = globalState;
        let _lhs60 = _this;
        _507_v35 = _rhs84;
        _509_v37 = _rhs85;
        _512_v38 = _rhs86;
        _lhs59.f2 = _rhs87;
        _lhs60.f5 = _rhs88;
        r0 = !(!((_this).f6));
        let _514_v40;
        _514_v40 = _dafny.Set.fromElements((_this).f6);
        (globalState).f2 = !(_514_v40).equals(_514_v40);
        let _515_v41;
        let _nw92 = new _module.C1();
        _nw92.__ctor(p0, (_this).f4);
        _515_v41 = _nw92;
        let _516_v42;
        let _nw93 = Array((new BigNumber(7)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = (((_this).f6) ? (_515_v41) : (_515_v41));
        _nw93[(_dafny.ONE).toNumber()] = _515_v41;
        _nw93[(new BigNumber(2)).toNumber()] = _515_v41;
        _nw93[(new BigNumber(3)).toNumber()] = _515_v41;
        _nw93[(new BigNumber(4)).toNumber()] = _515_v41;
        _nw93[(new BigNumber(5)).toNumber()] = _515_v41;
        _nw93[(new BigNumber(6)).toNumber()] = _515_v41;
        _516_v42 = _nw93;
        let _index77 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_516_v42).length));
        (_516_v42)[_index77] = _515_v41;
        let _index78 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_516_v42).length));
        (_516_v42)[_index78] = _515_v41;
      }
      let _517_v43;
      let _nw94 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
      _517_v43 = _nw94;
      let _518_v44;
      let _init11 = ((_519_p0) => function (_520_i0) {
        return _module.__default.safeModuloInt(_520_i0, _519_p0);
      })(p0);
      let _nw95 = Array((new BigNumber(15)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw95.length); _i0_11++) {
        _nw95[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _518_v44 = _nw95;
      let _521_v45;
      _521_v45 = new _dafny.CodePoint('q'.codePointAt(0));
      let _522_v46;
      _522_v46 = _dafny.Map.Empty.slice().updateUnsafe(_518_v44,_521_v45);
      let _index79 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_517_v43).length));
      (_517_v43)[_index79] = _522_v46;
      let _523_v47;
      _523_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-702)), function (_524_i1) {
        return (_this).f4;
      }),_522_v46);
      let _525_v48;
      _525_v48 = _dafny.Seq.UnicodeFromString("spwqw");
      let _526_v49;
      _526_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_525_v48);
      let _527_v50;
      _527_v50 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qsprbn")).length)), p0, (_dafny.ZERO).minus(new BigNumber(((((_526_v49).contains((_this).f6)) ? ((_526_v49).get((_this).f6)) : (_525_v48))).length)), new BigNumber(703), (_this).f4);
      let _index80 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_517_v43).length));
      (_517_v43)[_index80] = (((_523_v47).contains(_527_v50)) ? ((_523_v47).get(_527_v50)) : (_522_v46));
      if ((_this).f6) {
        let _528_v51;
        _528_v51 = _dafny.Seq.of((_this).f6);
        let _529_v52;
        let _nw96 = Array((new BigNumber(10)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = _module.__default.fm12((_this).f6, _module.__default.fm12((_this).f6, (_this).f6, (_this).f6, globalState), (_this).f6, globalState);
        _nw96[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(2)).toNumber()] = (function (_pat_let10_0) {
          return function (_530_dt__update__tmp_h1) {
            return function (_pat_let11_0) {
              return function (_531_dt__update_hcf3_h0) {
                return _module.D1.create_DC3(_531_dt__update_hcf3_h0);
              }(_pat_let11_0);
            }((_this).f6);
          }(_pat_let10_0);
        }(_module.D1.create_DC3((_this).f6))).dtor_cf3;
        _nw96[(new BigNumber(3)).toNumber()] = true;
        _nw96[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(5)).toNumber()] = (_528_v51)[_module.__default.safeIndex(p0, new BigNumber((_528_v51).length))];
        _nw96[(new BigNumber(6)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(8)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(9)).toNumber()] = !((_this).f6) || (!((_this).f6));
        _529_v52 = _nw96;
        let _532_v53;
        _532_v53 = _module.D1.create_DC4((_this).f6);
        let _index81 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
        (_529_v52)[_index81] = (_532_v53).dtor_cf4;
        let _index82 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
        let _rhs89 = _this.f5;
        let _rhs90 = _529_v52;
        let _rhs91 = (_this).f6;
        let _lhs61 = _529_v52;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
        r1 = _rhs89;
        _529_v52 = _rhs90;
        _lhs61[_lhs62] = _rhs91;
        if ((_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))]) {
          let _533_v54;
          _533_v54 = _dafny.MultiSet.fromElements(_521_v45);
          let _534_v55;
          _534_v55 = _dafny.Seq.of(_533_v54);
          let _535_v56;
          let _nw97 = new _module.C0();
          _nw97.__ctor((_534_v55)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_534_v55).length))]);
          _535_v56 = _nw97;
          let _536_v57;
          _536_v57 = _dafny.Seq.of(_535_v56, _535_v56);
          let _537_v58;
          _537_v58 = _dafny.Map.Empty.slice().updateUnsafe((_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))],(_536_v57)[_module.__default.safeIndex((_this).f4, new BigNumber((_536_v57).length))]);
          _537_v58 = (_537_v58).update(true, _535_v56);
          let _index83 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          (_529_v52)[_index83] = (_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))];
          r1 = (_module.__default.fm3(!((_this).f6), globalState)).plus((_this).f4);
          let _538_v59;
          _538_v59 = _dafny.MultiSet.fromElements(_535_v56);
          let _539_v60;
          _539_v60 = _dafny.Map.Empty.slice().updateUnsafe((((_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))]) ? (p0) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_540_v45) => function (_541_i2) {
            return _540_v45;
          })(_521_v45))).length))),(_this).f4);
          let _index84 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          let _rhs92 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4));
          let _rhs93 = _dafny.MultiSet.fromElements(_535_v56, _535_v56, _535_v56, _535_v56, _535_v56);
          let _rhs94 = !((_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))]);
          let _rhs95 = (_539_v60).Merge(_539_v60);
          let _rhs96 = _module.__default.fm12((p0).isLessThanOrEqualTo(_this.f5), (_this).f6, (_this).f6, globalState);
          let _lhs63 = _529_v52;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          let _lhs65 = globalState;
          r1 = _rhs92;
          _538_v59 = _rhs93;
          _lhs63[_lhs64] = _rhs94;
          _539_v60 = _rhs95;
          _lhs65.f2 = _rhs96;
          let _542_v61;
          _542_v61 = _dafny.MultiSet.fromElements(new BigNumber(-358));
          let _index85 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          (_529_v52)[_index85] = ((_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))]) || ((_542_v61).IsProperSubsetOf(_542_v61));
        } else {
          (globalState).f2 = _dafny.areEqual(_dafny.Seq.update(_525_v48, _module.__default.safeIndex(p0, new BigNumber((_525_v48).length)), new _dafny.CodePoint('n'.codePointAt(0))), _module.__default.fm6((_this).f6, p0, (_this).f6, globalState));
          let _543_v62;
          _543_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))]);
          _543_v62 = (((_this).f6) ? (_543_v62) : (_543_v62));
          let _544_v63;
          let _nw98 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
          _544_v63 = _nw98;
          let _545_v64;
          _545_v64 = _dafny.Set.fromElements((_this).f4, (_this).f4, (_this).f4);
          let _index86 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_544_v63).length));
          (_544_v63)[_index86] = _dafny.MultiSet.fromElements(_545_v64);
          let _546_v65;
          _546_v65 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_525_v48).length), new BigNumber(548), (_module.D3.create_DC9((_this).f4, (_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))], p0, (_this).f6)).dtor_cf8));
          let _index87 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_544_v63).length));
          (_544_v63)[_index87] = _546_v65;
          _545_v64 = (_545_v64).Intersect(_545_v64);
          r1 = ((_this.f5).minus(_this.f5)).plus(_module.__default.fm3((_this).f6, globalState));
        }
        if ((_this).f6) {
          (_this).f5 = (_this).f4;
          (_this).f5 = (_dafny.ZERO).minus((_this).f4);
          let _index88 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          (_529_v52)[_index88] = (_this).f6;
          _521_v45 = _521_v45;
          (globalState).f2 = (_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))];
        } else {
          let _index89 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          (_529_v52)[_index89] = (_this).f6;
          let _index90 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length));
          (_529_v52)[_index90] = (_this).f6;
          (_this).f5 = (_dafny.ZERO).minus(_this.f5);
          let _547_v66;
          _547_v66 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(_dafny.Seq.of((_this).f6, (_529_v52)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_529_v52).length))])),(_this).f6);
          let _index91 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_518_v44).length));
          (_518_v44)[_index91] = ((_dafny.ZERO).minus(new BigNumber((_547_v66).length))).multipliedBy(_this.f5);
          let _index92 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_518_v44).length));
          (_518_v44)[_index92] = _module.__default.safeDivisionInt(p0, (_this).f4);
          let _548_v67;
          let _nw99 = new _module.C1();
          _nw99.__ctor((_this).f4, (_this).f4);
          _548_v67 = _nw99;
          let _549_v68;
          _549_v68 = _dafny.Seq.of(_548_v67, _548_v67);
          let _550_v69;
          _550_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_549_v68).length));
          let _551_v70;
          _551_v70 = _dafny.Set.fromElements(new BigNumber((_550_v69).length));
          let _552_v72;
          _552_v72 = _dafny.Seq.of(_551_v70, function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(390), new BigNumber(552))) {
              let _553_v71 = _compr_12;
              if (((new BigNumber(390)).isLessThanOrEqualTo(_553_v71)) && ((_553_v71).isLessThan(new BigNumber(552)))) {
                _coll12.add((_553_v71).minus(_this.f5));
              }
            }
            return _coll12;
          }());
          (globalState).f2 = !(((_552_v72)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("jucpayd")).length), new BigNumber((_552_v72).length))]).IsDisjointFrom(_551_v70));
        }
        (_this).f5 = (_dafny.ZERO).minus(p0);
        r1 = (_this).f4;
      } else {
        r1 = (p0).plus(((_this).f4).minus(new BigNumber(209)));
        let _554_v73;
        _554_v73 = _dafny.MultiSet.fromElements(_521_v45);
        r1 = (new BigNumber(((_554_v73).Difference(_554_v73)).cardinality())).minus(_module.__default.safeDivisionInt(new BigNumber(615), new BigNumber((_dafny.MultiSet.fromElements((_this).f6, (_this).f6)).cardinality())));
        let _555_v74;
        _555_v74 = _module.D3.create_DC9((_this).f4, (_this).f6, (_this).f4, (_this).f6);
        let _556_v75;
        _556_v75 = _dafny.Seq.of(true, (_555_v74).dtor_cf9);
        r0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_556_v75, _556_v75), _556_v75);
        let _557_v76;
        let _nw100 = Array((new BigNumber(10)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = (((_this).f6) ? (_521_v45) : (_521_v45));
        _nw100[(_dafny.ONE).toNumber()] = (((_this).f6) ? (_521_v45) : (_521_v45));
        _nw100[(new BigNumber(2)).toNumber()] = (((_this).f6) ? ((_525_v48)[_module.__default.safeIndex((_this).f4, new BigNumber((_525_v48).length))]) : (_521_v45));
        _nw100[(new BigNumber(3)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(4)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(5)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(6)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(7)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(8)).toNumber()] = _521_v45;
        _nw100[(new BigNumber(9)).toNumber()] = _521_v45;
        _557_v76 = _nw100;
        let _index93 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_557_v76).length));
        (_557_v76)[_index93] = _521_v45;
        let _index94 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_557_v76).length));
        (_557_v76)[_index94] = _521_v45;
        (_this).f5 = _module.__default.fm3((_this).f6, globalState);
      }
      r0 = (_this).f6;
      r0 = (_this).f6;
      let _index95 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_518_v44).length));
      (_518_v44)[_index95] = (_this.f5).multipliedBy(p0);
      let _558_v77;
      let _init12 = ((_559_v45) => function (_560_i3) {
        return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), ((_561_v45) => function (_562_i4) {
          return _561_v45;
        })(_559_v45))).length), (_this).f4, new BigNumber(-661), (_this).f4);
      })(_521_v45);
      let _nw101 = Array((new BigNumber(2)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw101.length); _i0_12++) {
        _nw101[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _558_v77 = _nw101;
      let _index96 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_518_v44).length));
      let _rhs97 = (_527_v50)[_module.__default.safeIndex((_this).f4, new BigNumber((_527_v50).length))];
      let _rhs98 = _558_v77;
      let _lhs66 = _518_v44;
      let _lhs67 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_518_v44).length));
      _lhs66[_lhs67] = _rhs97;
      _558_v77 = _rhs98;
      r0 = (_this).f6;
      r1 = (_this.f5).plus((_518_v44)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_518_v44).length))]);
      return [r0, r1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f3 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f3) {
      let _this = this;
      (_this)._f3 = f3;
      return;
    }
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.Map.Empty;
      let _563_v0;
      let _nw102 = new _module.C2();
      _nw102.__ctor(_module.__default.fm12(p2, p2, (_this).f3, globalState), p1, p1);
      _563_v0 = _nw102;
      let _564_v1;
      _564_v1 = _dafny.Seq.UnicodeFromString("iqwgvcpna");
      if (_dafny.Seq.IsProperPrefixOf(_564_v1, _564_v1)) {
        let _565_v2;
        _565_v2 = _dafny.Seq.of(!((_563_v0).f6), _module.__default.fm12((_563_v0).f6, p2, (_563_v0).f6, globalState), true);
        (globalState).f2 = (_565_v2)[_module.__default.safeIndex(p1, new BigNumber((_565_v2).length))];
        if (((p2) ? ((_563_v0).f6) : (p2))) {
          let _566_v3;
          _566_v3 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),(_this).f3);
          let _567_v4;
          _567_v4 = _dafny.MultiSet.fromElements(_566_v3);
          let _568_v5;
          _568_v5 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_567_v4).cardinality())));
          let _569_v6;
          _569_v6 = _dafny.MultiSet.fromElements((_563_v0).f6, (_563_v0).f6, _dafny.Seq.IsProperPrefixOf(_568_v5, _568_v5), (p1).isLessThanOrEqualTo(new BigNumber((_564_v1).length)), !((_563_v0).f6));
          let _570_v7;
          _570_v7 = _dafny.Seq.of(_565_v2);
          _569_v6 = (_569_v6).Union((_569_v6).update(p2, _module.__default.abs(new BigNumber((_570_v7).length))));
          let _571_v8;
          _571_v8 = new _dafny.CodePoint('r'.codePointAt(0));
          _571_v8 = ((p2) ? (_571_v8) : (new _dafny.CodePoint('d'.codePointAt(0))));
          let _572_v9;
          _572_v9 = _dafny.Map.Empty.slice().updateUnsafe(_571_v8,_564_v1);
          _572_v9 = (_572_v9).update(_571_v8, _564_v1);
          let _573_v10;
          _573_v10 = _dafny.Set.fromElements(p2);
          (globalState).f2 = ((_573_v10).Difference(_573_v10)).IsDisjointFrom(_573_v10);
          let _574_v11;
          _574_v11 = new BigNumber(-107);
          _574_v11 = p1;
        } else {
          let _575_v12;
          let _nw103 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _575_v12 = _nw103;
          let _576_v13;
          _576_v13 = new _dafny.CodePoint('p'.codePointAt(0));
          let _index97 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_575_v12).length));
          (_575_v12)[_index97] = _576_v13;
          let _index98 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_575_v12).length));
          (_575_v12)[_index98] = _576_v13;
          let _577_v14;
          _577_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_564_v1).length));
          let _578_v15;
          _578_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _579_v16;
          _579_v16 = _dafny.Set.fromElements((_563_v0).f6);
          let _580_v17;
          _580_v17 = _module.D4.create_DC12(_579_v16);
          _577_v14 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm13((_563_v0).f6, new BigNumber((_578_v15).length), p1, globalState)).cardinality()),new BigNumber(((_580_v17).dtor_cf14).length))).Merge(_577_v14);
          let _581_v18;
          _581_v18 = new BigNumber(822);
          let _582_v19;
          _582_v19 = _dafny.Map.Empty.slice().updateUnsafe((_563_v0).f6,_581_v18);
          let _583_v20;
          _583_v20 = _dafny.Seq.of(_581_v18);
          let _584_v21;
          let _nw104 = new _module.C2();
          _nw104.__ctor((_563_v0).f6, new BigNumber((_583_v20).length), _581_v18);
          _584_v21 = _nw104;
          let _585_v22;
          _585_v22 = _dafny.Map.Empty.slice().updateUnsafe(_581_v18,_584_v21);
          _581_v18 = (((_582_v19).contains((_563_v0).f6)) ? ((_582_v19).get((_563_v0).f6)) : ((new BigNumber((_dafny.Seq.update(_module.__default.fm6(false, new BigNumber((_585_v22).length), (_this).f3, globalState), _module.__default.safeIndex(new BigNumber(117), new BigNumber((_module.__default.fm6(false, new BigNumber((_585_v22).length), (_this).f3, globalState)).length)), (_575_v12)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_575_v12).length))])).length)).multipliedBy(p1)));
          let _586_v23;
          let _nw105 = Array((new BigNumber(26)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = p2;
          _nw105[(_dafny.ONE).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(2)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(3)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(4)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(5)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(6)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(7)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(8)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(9)).toNumber()] = p2;
          _nw105[(new BigNumber(10)).toNumber()] = p2;
          _nw105[(new BigNumber(11)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(12)).toNumber()] = p2;
          _nw105[(new BigNumber(13)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(14)).toNumber()] = p2;
          _nw105[(new BigNumber(15)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(16)).toNumber()] = !(p2);
          _nw105[(new BigNumber(17)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(18)).toNumber()] = p2;
          _nw105[(new BigNumber(19)).toNumber()] = (_this).f3;
          _nw105[(new BigNumber(20)).toNumber()] = (_this).f3;
          _nw105[(new BigNumber(21)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(22)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(23)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(24)).toNumber()] = (_563_v0).f6;
          _nw105[(new BigNumber(25)).toNumber()] = (_563_v0).f6;
          _586_v23 = _nw105;
          let _587_v24;
          _587_v24 = _dafny.Seq.of(_586_v23, _586_v23);
          let _588_v25;
          _588_v25 = _dafny.Seq.of(_587_v24);
          let _589_v26;
          _589_v26 = _dafny.Seq.of(_587_v24, _dafny.Seq.update((_588_v25)[_module.__default.safeIndex(p1, new BigNumber((_588_v25).length))], _module.__default.safeIndex((_584_v21).f4, new BigNumber(((_588_v25)[_module.__default.safeIndex(p1, new BigNumber((_588_v25).length))]).length)), _586_v23));
          let _590_v27;
          _590_v27 = _dafny.Map.Empty.slice().updateUnsafe(_576_v13,new BigNumber(((_589_v26)[_module.__default.safeIndex(_584_v21.f5, new BigNumber((_589_v26).length))]).length));
          let _591_v28;
          let _nw106 = Array((new BigNumber(3)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = p2;
          _nw106[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), ((_592_v13, _593_v20) => function (_594_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(_592_v13,new BigNumber((_593_v20).length));
          })(_576_v13, _583_v20)), _590_v27);
          _nw106[(new BigNumber(2)).toNumber()] = !(!((_584_v21).f4).isEqualTo(p1));
          _591_v28 = _nw106;
          let _index99 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_575_v12).length));
          let _rhs99 = (((((_563_v0).f6) ? ((_563_v0).f6) : (p2))) ? ((_564_v1)[_module.__default.safeIndex(p1, new BigNumber((_564_v1).length))]) : (_576_v13));
          let _rhs100 = (_module.D5.create_DC15(_591_v28)).dtor_cf19;
          let _rhs101 = (_dafny.ZERO).minus(_module.__default.fm3(p2, globalState));
          let _lhs68 = _575_v12;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_575_v12).length));
          _lhs68[_lhs69] = _rhs99;
          _586_v23 = _rhs100;
          _581_v18 = _rhs101;
          _585_v22 = (_585_v22).update((p1).plus((_dafny.ZERO).minus(p1)), _584_v21);
        }
        let _595_v29;
        _595_v29 = new BigNumber(446);
        let _596_v30;
        _596_v30 = _dafny.MultiSet.fromElements(!((_this).f3));
        _595_v29 = new BigNumber((_596_v30).cardinality());
        let _597_v31;
        _597_v31 = _dafny.Seq.of((((_563_v0).f6) ? (p1) : (p1)));
        _595_v29 = (_dafny.ZERO).minus((_597_v31)[_module.__default.safeIndex(p1, new BigNumber((_597_v31).length))]);
        (globalState).f2 = !((p1).isLessThan((((_563_v0).f6) ? (p1) : ((_dafny.ZERO).minus(new BigNumber((_564_v1).length))))));
      } else {
        let _598_v32;
        let _nw107 = new _module.C1();
        _nw107.__ctor(p1, p1);
        _598_v32 = _nw107;
        let _599_v33;
        _599_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_598_v32);
        let _600_v34;
        let _nw108 = Array((new BigNumber(7)).toNumber()).fill(false);
        _600_v34 = _nw108;
        let _601_v35;
        _601_v35 = _dafny.Seq.of(_600_v34);
        let _602_v36;
        _602_v36 = _module.D4.create_DC14((_599_v33).Merge(_599_v33), p2, _601_v35);
        let _pat_let_tv11 = _599_v33;
        _602_v36 = function (_pat_let12_0) {
          return function (_603_dt__update__tmp_h0) {
            return function (_pat_let13_0) {
              return function (_604_dt__update_hcf16_h0) {
                return function (_pat_let14_0) {
                  return function (_605_dt__update_hcf17_h0) {
                    return _module.D4.create_DC14(_604_dt__update_hcf16_h0, _605_dt__update_hcf17_h0, (_603_dt__update__tmp_h0).dtor_cf18);
                  }(_pat_let14_0);
                }((_this).f3);
              }(_pat_let13_0);
            }(_pat_let_tv11);
          }(_pat_let12_0);
        }(_602_v36);
        (_598_v32).f5 = _598_v32.f5;
        let _606_v37;
        let _init13 = ((_607_v0, _608_p2) => function (_609_i1) {
          return _dafny.MultiSet.fromElements((_607_v0).f6, _608_p2);
        })(_563_v0, p2);
        let _nw109 = Array((new BigNumber(6)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw109.length); _i0_13++) {
          _nw109[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _606_v37 = _nw109;
        let _610_v38;
        _610_v38 = _dafny.MultiSet.fromElements((_563_v0).f6);
        let _index100 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_606_v37).length));
        (_606_v37)[_index100] = (_610_v38).Intersect(_610_v38);
        let _index101 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_606_v37).length));
        (_606_v37)[_index101] = _dafny.MultiSet.fromElements(!((_563_v0).f6), (_this).f3);
        let _611_v39;
        _611_v39 = _dafny.Seq.of(p2);
        let _612_v40;
        _612_v40 = _module.D2.create_DC6(_611_v39);
        let _source8 = _612_v40;
        if (_source8.is_DC7) {
          let _613_v41;
          let _nw110 = Array((new BigNumber(26)).toNumber()).fill([]);
          _613_v41 = _nw110;
          let _614_v42;
          _614_v42 = new _dafny.CodePoint('e'.codePointAt(0));
          let _615_v43;
          _615_v43 = _dafny.Map.Empty.slice().updateUnsafe(_614_v42,_module.__default.fm6(p2, _598_v32.f5, !((_this).f3), globalState));
          let _616_v44;
          let _nw111 = new _module.C1();
          _nw111.__ctor(new BigNumber(((((_615_v43).contains(_614_v42)) ? ((_615_v43).get(_614_v42)) : (_564_v1))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), ((_617_v32, _618_v1, _619_p1) => function (_620_i2) {
            return new BigNumber((_dafny.Seq.of(new BigNumber(502), (_617_v32).f4, new BigNumber((_618_v1).length), _619_p1)).length);
          })(_598_v32, _564_v1, p1))).length));
          _616_v44 = _nw111;
          let _621_v45;
          _621_v45 = _dafny.Map.Empty.slice().updateUnsafe(_616_v44,_598_v32.f5);
          let _rhs102 = _600_v34;
          let _rhs103 = _613_v41;
          let _rhs104 = !((((_621_v45).contains(_616_v44)) ? ((_621_v45).get(_616_v44)) : (new BigNumber(372)))).isEqualTo(p1);
          let _lhs70 = globalState;
          _600_v34 = _rhs102;
          _613_v41 = _rhs103;
          _lhs70.f2 = _rhs104;
          let _622_v46;
          let _nw112 = Array((new BigNumber(17)).toNumber()).fill(_module.D3.Default());
          _622_v46 = _nw112;
          let _623_v47;
          _623_v47 = _dafny.Map.Empty.slice().updateUnsafe(_622_v46,p1);
          let _624_v48;
          _624_v48 = _dafny.Map.Empty.slice().updateUnsafe((_563_v0).f6,p2);
          let _625_v49;
          _625_v49 = _dafny.Map.Empty.slice().updateUnsafe((_623_v47).Merge(_623_v47),(_dafny.Map.Empty.slice().updateUnsafe((_563_v0).f6,p2)).Merge(_624_v48));
          _625_v49 = (_625_v49).update(_623_v47, _624_v48);
          (globalState).f2 = (_563_v0).f6;
          let _626_v50;
          let _nw113 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _626_v50 = _nw113;
          let _index102 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_626_v50).length));
          (_626_v50)[_index102] = (_624_v48).update(true, false);
          let _index103 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_626_v50).length));
          (_626_v50)[_index103] = (_module.__default.fm14(_598_v32.f5, globalState)).Merge(_624_v48);
        } else {
          let _627___mcc_h0 = (_source8).cf6;
          let _628_cf6 = _627___mcc_h0;
          let _629_v51;
          let _nw114 = new _module.C2();
          _nw114.__ctor(p2, (_dafny.ZERO).minus(_598_v32.f5), new BigNumber(645));
          _629_v51 = _nw114;
          let _630_v52;
          _630_v52 = _module.D3.create_DC8(_564_v1);
          let _631_v53;
          _631_v53 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_564_v1).length)).plus((_dafny.ZERO).minus(_598_v32.f5)),_module.D5.create_DC16(_598_v32.f5, (_606_v37)[_module.__default.safeIndex(new BigNumber(147), new BigNumber((_606_v37).length))], _dafny.Seq.of(_module.__default.fm12((_this).f3, (_563_v0).f6, (_629_v51).f6, globalState), _module.__default.fm12((_563_v0).f6, p2, !((_this).f3), globalState)), _630_v52));
          let _632_v54;
          _632_v54 = _module.D5.create_DC16(new BigNumber((_564_v1).length), _dafny.MultiSet.FromArray(_611_v39), _628_cf6, _630_v52);
          _631_v53 = (_631_v53).update(new BigNumber((_module.__default.fm15(globalState)).length), _632_v54);
          let _633_v55;
          _633_v55 = new _dafny.CodePoint('i'.codePointAt(0));
          let _634_v56;
          let _nw115 = new _module.C0();
          _nw115.__ctor(_dafny.MultiSet.fromElements(_633_v55));
          _634_v56 = _nw115;
          let _635_v57;
          let _nw116 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _635_v57 = _nw116;
          let _636_v58;
          _636_v58 = _dafny.Seq.of(_635_v57);
          _636_v58 = _dafny.Seq.of(_635_v57, _635_v57, _635_v57, _635_v57, _635_v57);
        }
        (globalState).f2 = (_this).f3;
      }
      let _hi3 = new BigNumber((_564_v1).length);
      for (let _637_i3 = p1; _637_i3.isLessThan(_hi3); _637_i3 = _637_i3.plus(_dafny.ONE)) {
        let _638_v59;
        let _init14 = function (_639_i4) {
          return !((_this).f3);
        };
        let _nw117 = Array((new BigNumber(9)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw117.length); _i0_14++) {
          _nw117[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _638_v59 = _nw117;
        let _index104 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_638_v59).length));
        (_638_v59)[_index104] = ((p2) ? ((_this).f3) : ((_this).f3));
        let _index105 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_638_v59).length));
        (_638_v59)[_index105] = (p2) === ((_563_v0).f6);
        let _640_v60;
        _640_v60 = new _dafny.CodePoint('w'.codePointAt(0));
        let _641_v61;
        let _nw118 = new _module.C0();
        _nw118.__ctor(_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), _640_v60));
        _641_v61 = _nw118;
        let _642_v62;
        _642_v62 = _dafny.MultiSet.fromElements(_641_v61, _641_v61);
        let _643_v63;
        let _nw119 = new _module.C1();
        _nw119.__ctor(new BigNumber((_642_v62).cardinality()), new BigNumber(831));
        _643_v63 = _nw119;
        let _nw120 = new _module.C1();
        _nw120.__ctor(_module.__default.fm3((_module.D1.create_DC4(!(false))).dtor_cf4, globalState), _637_i3);
        _643_v63 = _nw120;
        let _644_v64;
        _644_v64 = new BigNumber(149);
        let _645_v65;
        _645_v65 = _dafny.Set.fromElements((_563_v0).f6);
        let _646_v66;
        _646_v66 = _dafny.Set.fromElements(new BigNumber((_645_v65).length));
        let _647_v68;
        _647_v68 = _dafny.Set.fromElements(_646_v66, function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(595), _dafny.ONE)) {
            let _648_v67 = _compr_13;
            if (((new BigNumber(595)).isLessThanOrEqualTo(_648_v67)) && ((_648_v67).isLessThan(_dafny.ONE))) {
              _coll13.add(_module.__default.safeDivisionInt(_648_v67, new BigNumber(-992)));
            }
          }
          return _coll13;
        }());
        _644_v64 = ((_dafny.ZERO).minus(new BigNumber((_647_v68).length))).minus(_637_i3);
        let _649_v69;
        let _nw121 = new _module.C0();
        _nw121.__ctor((_641_v61).f7);
        _649_v69 = _nw121;
      }
      let _650_v70;
      _650_v70 = _dafny.Seq.of(true);
      let _651_v71;
      _651_v71 = _module.D1.create_DC4((_650_v70)[_module.__default.safeIndex(p1, new BigNumber((_650_v70).length))]);
      let _652_v72;
      _652_v72 = _module.D1.create_DC5(_651_v71);
      _652_v72 = _652_v72;
      let _653_v73;
      _653_v73 = new _dafny.CodePoint('y'.codePointAt(0));
      let _654_v74;
      _654_v74 = _dafny.Map.Empty.slice().updateUnsafe(p2,_653_v73);
      _654_v74 = (_654_v74).update(p2, _653_v73);
      let _hi4 = p1;
      for (let _655_i5 = p1; _655_i5.isLessThan(_hi4); _655_i5 = _655_i5.plus(_dafny.ONE)) {
        let _656_v75;
        _656_v75 = _dafny.Map.Empty.slice().updateUnsafe(_655_i5,(_563_v0).f6);
        let _657_v76;
        _657_v76 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm16((_563_v0).f6, _module.__default.fm3((_563_v0).f6, globalState), (_563_v0).f6, globalState)).length), new BigNumber(504), _655_i5);
        let _658_v77;
        let _nw122 = new _module.C1();
        _nw122.__ctor(new BigNumber(-609), (new BigNumber((_656_v75).length)).minus(new BigNumber((_657_v76).cardinality())));
        _658_v77 = _nw122;
        let _659_v78;
        _659_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_655_i5);
        _659_v78 = (_659_v78).update(true, (((_563_v0).f6) ? (p1) : (p1)));
        let _660_v79;
        _660_v79 = new BigNumber(35);
        _660_v79 = (_dafny.ZERO).minus(_660_v79);
        let _661_v80;
        _661_v80 = _module.D0.create_DC2();
        _661_v80 = _661_v80;
      }
      let _662_v81;
      _662_v81 = _dafny.Seq.of(new BigNumber((_564_v1).length), (_dafny.ZERO).minus(p1));
      r0 = _dafny.Seq.Concat(_662_v81, _662_v81);
      let _663_v82;
      _663_v82 = _dafny.Map.Empty.slice().updateUnsafe(p1,_653_v73);
      let _664_v83;
      _664_v83 = _dafny.Map.Empty.slice().updateUnsafe(_663_v82,(_563_v0).f6);
      r1 = _664_v83;
      return [r0, r1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
