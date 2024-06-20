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
    static fm4(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('s'.codePointAt(0));
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("qcjolb"), _dafny.Seq.UnicodeFromString("jvvpxlotu"));
    };
    static fm6(p0, p1, p2, globalState) {
      return (new BigNumber(42)).minus(new BigNumber(-345));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), function (_0_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uus"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(725)), function (_1_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })));
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(692), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(604)), function (_2_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("j")).length);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_3_i1) {
        return new BigNumber(206);
      })));
    };
    static fm9(p0, p1, p2, globalState) {
      if (true) {
        return _dafny.Seq.of(!(!(!(!(false)))), !(false));
      } else {
        return _dafny.Seq.of(true);
      }
    };
    static m0(globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _4_v0;
      _4_v0 = new BigNumber(341);
      let _5_v1;
      _5_v1 = _dafny.Seq.of(_4_v0);
      let _6_v2;
      _6_v2 = new _dafny.CodePoint('n'.codePointAt(0));
      let _7_v3;
      _7_v3 = false;
      let _8_v4;
      let _nw0 = Array((new BigNumber(12)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _5_v1;
      _nw0[(_dafny.ONE).toNumber()] = _5_v1;
      _nw0[(new BigNumber(2)).toNumber()] = _5_v1;
      _nw0[(new BigNumber(3)).toNumber()] = _5_v1;
      _nw0[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("wuojgho")).length), _4_v0);
      _nw0[(new BigNumber(5)).toNumber()] = _5_v1;
      _nw0[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(439)), ((_9_v0) => function (_10_i0) {
        return (_dafny.ZERO).minus(_9_v0);
      })(_4_v0));
      _nw0[(new BigNumber(7)).toNumber()] = _5_v1;
      _nw0[(new BigNumber(8)).toNumber()] = _5_v1;
      _nw0[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_5_v1, _5_v1);
      _nw0[(new BigNumber(10)).toNumber()] = _module.__default.fm8(_6_v2, _7_v3, _5_v1, globalState);
      _nw0[(new BigNumber(11)).toNumber()] = _5_v1;
      _8_v4 = _nw0;
      let _11_v5;
      _11_v5 = _module.D3.create_DC11(_dafny.Seq.Concat(_dafny.Seq.of(_4_v0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), function (_12_i1) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("axdp")).length);
})));
      let _13_v6;
      _13_v6 = _dafny.Set.fromElements(_7_v3, _7_v3, _7_v3, _7_v3, _7_v3);
      let _14_v7;
      _14_v7 = _dafny.Seq.UnicodeFromString("ekcwr");
      let _rhs0 = _module.__default.fm5(_7_v3, (new BigNumber((_13_v6).length)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), ((_15_v2) => function (_16_i2) {
        return _15_v2;
      })(_6_v2))).length)), _4_v0, _4_v0, globalState);
      let _rhs1 = (((_7_v3) ? (_4_v0) : (_4_v0))).isEqualTo((_dafny.ZERO).minus(_module.__default.fm6(_7_v3, _7_v3, false, globalState)));
      let _rhs2 = _8_v4;
      let _rhs3 = (_module.__default.fm6(true, false, _7_v3, globalState)).plus(new BigNumber((_14_v7).length));
      let _rhs4 = _11_v5;
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      _lhs0.f20 = _rhs0;
      _lhs1.f20 = _rhs1;
      _8_v4 = _rhs2;
      r2 = _rhs3;
      _11_v5 = _rhs4;
      let _17_v8;
      let _nw1 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _17_v8 = _nw1;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_17_v8).length))) {
        let _18_i3 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_18_i3)) && ((_18_i3).isLessThan(new BigNumber((_17_v8).length))))) {
          (_17_v8)[(_18_i3)] = _module.__default.safeDivisionInt(_18_i3, _4_v0);
        }
      }
      let _19_v9;
      let _nw2 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
      _19_v9 = _nw2;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_19_v9).length))) {
        let _20_i4 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_20_i4)) && ((_20_i4).isLessThan(new BigNumber((_19_v9).length))))) {
          (_19_v9)[(_20_i4)] = _dafny.Map.Empty.slice().updateUnsafe(_4_v0,_module.D4.create_DC18());
        }
      }
      (globalState).f20 = !(_7_v3);
      if ((_4_v0).isLessThan(_module.__default.safeDivisionInt(_4_v0, _4_v0))) {
        let _21_v10;
        _21_v10 = _module.D0.create_DC1(!(_7_v3), _7_v3, _4_v0);
        let _22_v11;
        _22_v11 = _dafny.Seq.of(_7_v3);
        let _23_v12;
        _23_v12 = _dafny.Map.Empty.slice().updateUnsafe(_4_v0,_4_v0);
        let _24_v13;
        _24_v13 = _module.D5.create_DC21(_23_v12);
        let _25_v14;
        _25_v14 = _module.D5.create_DC22(_24_v13);
        let _26_v15;
        _26_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-355),false);
        let _27_v16;
        let _nw3 = Array((new BigNumber(28)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = !(true) || (_7_v3);
        _nw3[(_dafny.ONE).toNumber()] = _7_v3;
        _nw3[(new BigNumber(2)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(3)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(4)).toNumber()] = ((_7_v3) ? (_7_v3) : (_7_v3));
        _nw3[(new BigNumber(5)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(6)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(7)).toNumber()] = false;
        _nw3[(new BigNumber(8)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(9)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(10)).toNumber()] = (((_21_v10).dtor_cf1) ? (_7_v3) : (_7_v3));
        _nw3[(new BigNumber(11)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(12)).toNumber()] = (_7_v3) === (_7_v3);
        _nw3[(new BigNumber(13)).toNumber()] = true;
        _nw3[(new BigNumber(14)).toNumber()] = !(!(!(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_22_v11).length),_25_v14)).length)).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), ((_28_v2) => function (_29_i5) {
          return _28_v2;
        })(_6_v2))).length))));
        _nw3[(new BigNumber(15)).toNumber()] = !(_7_v3) || (_7_v3);
        _nw3[(new BigNumber(16)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(17)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(18)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(19)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(20)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(21)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(22)).toNumber()] = true;
        _nw3[(new BigNumber(23)).toNumber()] = !(_7_v3);
        _nw3[(new BigNumber(24)).toNumber()] = _7_v3;
        _nw3[(new BigNumber(25)).toNumber()] = !_dafny.Seq.contains(_module.__default.fm7(_7_v3, new BigNumber(110), _4_v0, false, globalState), _6_v2);
        _nw3[(new BigNumber(26)).toNumber()] = ((((_26_v15).contains(_4_v0)) ? ((_26_v15).get(_4_v0)) : (true))) && (_7_v3);
        _nw3[(new BigNumber(27)).toNumber()] = false;
        _27_v16 = _nw3;
        let _index0 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
        (_27_v16)[_index0] = true;
        let _index1 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
        let _rhs5 = (_4_v0).minus((_4_v0).minus(_4_v0));
        let _rhs6 = _7_v3;
        let _lhs2 = globalState;
        let _lhs3 = _27_v16;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
        _lhs2.f5 = _rhs5;
        _lhs3[_lhs4] = _rhs6;
        let _30_v17;
        let _nw4 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _30_v17 = _nw4;
        _30_v17 = (((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))]) ? (_30_v17) : (_30_v17));
        let _index2 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
        (_27_v16)[_index2] = (_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))];
        let _index3 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
        (_27_v16)[_index3] = !((_dafny.ZERO).minus(_4_v0)).isEqualTo(_4_v0);
        if ((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))]) {
          let _index4 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
          (_27_v16)[_index4] = ((_4_v0).minus(_4_v0)).isLessThan(_4_v0);
          let _index5 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length));
          (_27_v16)[_index5] = _7_v3;
          let _31_v18;
          _31_v18 = _dafny.Set.fromElements(_4_v0, new BigNumber((_5_v1).length));
          _7_v3 = (_31_v18).IsProperSubsetOf(_31_v18);
          let _32_v19;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_4_v0, !(_module.__default.fm5((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))], _7_v3, _4_v0, new BigNumber((_22_v11).length), globalState)), _dafny.Seq.Concat(_14_v7, _14_v7));
          _32_v19 = _nw5;
          let _33_v20;
          _33_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,_32_v19);
          let _34_v21;
          _34_v21 = _module.D1.create_DC4((((_33_v20).contains(_7_v3)) ? ((_33_v20).get(_7_v3)) : (_32_v19)));
          _34_v21 = _34_v21;
        } else {
          _22_v11 = _dafny.Seq.Concat(_dafny.Seq.of((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))], _module.__default.fm5(_7_v3, _7_v3, _4_v0, _4_v0, globalState)), _module.__default.fm9(_dafny.Seq.update(_22_v11, _module.__default.safeIndex(_4_v0, new BigNumber((_22_v11).length)), (_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))]), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_35_v2) => function (_36_i6) {
            return _35_v2;
          })(_6_v2))).length), (_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))], globalState));
          let _37_v22;
          let _nw6 = new _module.C0();
          _nw6.__ctor(new BigNumber((_13_v6).length), false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-589)), ((_38_v2) => function (_39_i7) {
            return _38_v2;
          })(_6_v2)));
          _37_v22 = _nw6;
          let _40_v23;
          _40_v23 = _dafny.Seq.of(_37_v22);
          let _41_v24;
          let _nw7 = Array((new BigNumber(18)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _37_v22;
          _nw7[(_dafny.ONE).toNumber()] = _37_v22;
          _nw7[(new BigNumber(2)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(3)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(4)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(5)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(6)).toNumber()] = (_40_v23)[_module.__default.safeIndex((_37_v22).f26, new BigNumber((_40_v23).length))];
          _nw7[(new BigNumber(7)).toNumber()] = (((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))]) ? (_37_v22) : (_37_v22));
          _nw7[(new BigNumber(8)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(9)).toNumber()] = ((_7_v3) ? (_37_v22) : (_37_v22));
          _nw7[(new BigNumber(10)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(11)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(12)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(13)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(14)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(15)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(16)).toNumber()] = _37_v22;
          _nw7[(new BigNumber(17)).toNumber()] = _37_v22;
          _41_v24 = _nw7;
          let _index6 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_41_v24).length));
          (_41_v24)[_index6] = (((_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))]) ? (_37_v22) : (_37_v22));
          let _42_v25;
          _42_v25 = _dafny.Map.Empty.slice().updateUnsafe(_30_v17,_37_v22);
          let _index7 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_41_v24).length));
          (_41_v24)[_index7] = ((((_37_v22).f26).isLessThan(_4_v0)) ? (((_7_v3) ? (_37_v22) : (_37_v22))) : ((((_42_v25).contains(_30_v17)) ? ((_42_v25).get(_30_v17)) : (_37_v22))));
          let _43_v26;
          _43_v26 = _dafny.MultiSet.fromElements((_41_v24)[_module.__default.safeIndex(new BigNumber(929), new BigNumber((_41_v24).length))], _37_v22);
          let _44_v27;
          let _nw8 = new _module.C0();
          _nw8.__ctor((new BigNumber(365)).plus(new BigNumber((_43_v26).cardinality())), _7_v3, _14_v7);
          _44_v27 = _nw8;
          let _45_v28;
          _45_v28 = _dafny.MultiSet.fromElements(_17_v8);
          let _46_v29;
          let _nw9 = new _module.C0();
          _nw9.__ctor(new BigNumber((_45_v28).cardinality()), (_27_v16)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_27_v16).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(126)), ((_47_v2) => function (_48_i8) {
            return _47_v2;
          })(_6_v2)));
          _46_v29 = _nw9;
          let _49_v30;
          _49_v30 = _dafny.Map.Empty.slice().updateUnsafe(_46_v29,(_44_v27).f26);
          _49_v30 = (_49_v30).update(_46_v29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-65)), ((_50_v2) => function (_51_i9) {
            return _50_v2;
          })(_6_v2))).length));
          let _52_v31;
          let _nw10 = new _module.C0();
          _nw10.__ctor((_44_v27).f26, _7_v3, _14_v7);
          _52_v31 = _nw10;
        }
      } else {
        (globalState).f5 = _4_v0;
        _5_v1 = _5_v1;
        let _53_v32;
        let _nw11 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
        _53_v32 = _nw11;
        let _54_v33;
        let _init0 = function (_55_i10) {
          return true;
        };
        let _nw12 = Array((new BigNumber(13)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw12.length); _i0_0++) {
          _nw12[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _54_v33 = _nw12;
        let _56_v34;
        _56_v34 = _module.D2.create_DC8(_54_v33);
        let _57_v35;
        _57_v35 = _module.D2.create_DC10(_56_v34);
        let _58_v36;
        _58_v36 = _dafny.Map.Empty.slice().updateUnsafe(_57_v35,_4_v0);
        let _59_v37;
        let _nw13 = new _module.C0();
        _nw13.__ctor((_dafny.ZERO).minus(new BigNumber((_58_v36).length)), _module.__default.fm5(_7_v3, _7_v3, _4_v0, _4_v0, globalState), _14_v7);
        _59_v37 = _nw13;
        let _60_v38;
        _60_v38 = _53_v32;
        let _rhs7 = (_60_v38);
        let _rhs8 = _59_v37;
        let _rhs9 = ((_59_v37).f26).multipliedBy(new BigNumber((_5_v1).length));
        let _lhs5 = globalState;
        _53_v32 = _rhs7;
        _59_v37 = _rhs8;
        _lhs5.f5 = _rhs9;
        let _61_v39;
        _61_v39 = _dafny.Map.Empty.slice().updateUnsafe(_7_v3,_7_v3);
        let _62_v40;
        _62_v40 = _dafny.Seq.of(_7_v3);
        _61_v39 = _dafny.Map.Empty.slice().updateUnsafe(_7_v3,(_62_v40)[_module.__default.safeIndex((_59_v37).f26, new BigNumber((_62_v40).length))]);
        let _63_v41;
        _63_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,_6_v2);
        let _64_v42;
        _64_v42 = _module.D4.create_DC17(new BigNumber((_63_v41).length), _59_v37);
        let _65_v43;
        _65_v43 = _dafny.Map.Empty.slice().updateUnsafe((_64_v42).dtor_cf29,_7_v3);
        _65_v43 = (_65_v43).update(_4_v0, _7_v3);
      }
      let _66_v44;
      let _nw14 = new _module.C0();
      _nw14.__ctor(_4_v0, _module.__default.fm5(_7_v3, _7_v3, new BigNumber(-857), _module.__default.fm6(_7_v3, _7_v3, _7_v3, globalState), globalState), _14_v7);
      _66_v44 = _nw14;
      let _67_v45;
      _67_v45 = _dafny.Map.Empty.slice().updateUnsafe(_66_v44,new BigNumber(-276));
      let _68_v46;
      _68_v46 = _dafny.Map.Empty.slice().updateUnsafe(_67_v45,_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_69_v2) => function (_70_i11) {
        return _69_v2;
      })(_6_v2)));
      (globalState).f4 = _module.__default.fm7(_7_v3, (_dafny.ZERO).minus(new BigNumber((_68_v46).length)), _4_v0, _7_v3, globalState);
      r0 = _dafny.Seq.Concat(_14_v7, _dafny.Seq.Concat(_14_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), ((_71_v2) => function (_72_i12) {
        return _71_v2;
      })(_6_v2))));
      let _73_v47;
      let _nw15 = Array((new BigNumber(9)).toNumber()).fill([]);
      _73_v47 = _nw15;
      r1 = _73_v47;
      r2 = _4_v0;
      let _74_v48;
      let _nw16 = new _module.C0();
      _nw16.__ctor(_4_v0, _7_v3, _14_v7);
      _74_v48 = _nw16;
      let _75_v49;
      _75_v49 = _dafny.MultiSet.fromElements(_74_v48, _74_v48, _74_v48);
      let _76_v50;
      _76_v50 = _dafny.Seq.of(_74_v48, _74_v48);
      r3 = ((_74_v48).f26).isLessThan((((_75_v49).contains((_76_v50)[_module.__default.safeIndex((_74_v48).f26, new BigNumber((_76_v50).length))])) ? ((_75_v49).get((_76_v50)[_module.__default.safeIndex((_74_v48).f26, new BigNumber((_76_v50).length))])) : (_4_v0)));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _77_v0;
      _77_v0 = _dafny.Seq.UnicodeFromString("mkyq");
      let _78_v1;
      _78_v1 = false;
      let _79_v2;
      _79_v2 = _dafny.Set.fromElements(_78_v1, _78_v1);
      let _80_v3;
      let _nw17 = Array((new BigNumber(15)).toNumber());
      _nw17[(_dafny.ZERO).toNumber()] = !(false);
      _nw17[(_dafny.ONE).toNumber()] = _78_v1;
      _nw17[(new BigNumber(2)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(3)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(4)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(5)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(6)).toNumber()] = true;
      _nw17[(new BigNumber(7)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(8)).toNumber()] = !(_78_v1);
      _nw17[(new BigNumber(9)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(10)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(11)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(12)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(13)).toNumber()] = _78_v1;
      _nw17[(new BigNumber(14)).toNumber()] = _78_v1;
      _80_v3 = _nw17;
      let _81_v4;
      _81_v4 = _module.D0.create_DC0(_77_v0);
      let _82_v6;
      _82_v6 = new BigNumber(462);
      let _83_v7;
      _83_v7 = _dafny.Seq.of(_82_v6, _82_v6, _82_v6);
      let _84_v8;
      _84_v8 = new _dafny.CodePoint('x'.codePointAt(0));
      let _85_v9;
      let _nw18 = Array((new BigNumber(27)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = _84_v8;
      _nw18[(_dafny.ONE).toNumber()] = _84_v8;
      _nw18[(new BigNumber(2)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
      _nw18[(new BigNumber(4)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(5)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(6)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(7)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(8)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(9)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(10)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(11)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(12)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(13)).toNumber()] = (_77_v0)[_module.__default.safeIndex(_82_v6, new BigNumber((_77_v0).length))];
      _nw18[(new BigNumber(14)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(15)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(16)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(17)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
      _nw18[(new BigNumber(19)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(20)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(21)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(22)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(23)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(24)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(25)).toNumber()] = _84_v8;
      _nw18[(new BigNumber(26)).toNumber()] = _84_v8;
      _85_v9 = _nw18;
      let _86_v10;
      _86_v10 = _dafny.Seq.of(_85_v9, _85_v9, _85_v9, _85_v9, _85_v9);
      let _87_v11;
      _87_v11 = _dafny.MultiSet.fromElements(_82_v6);
      let _88_v12;
      _88_v12 = _dafny.Seq.of(_87_v11);
      let _89_v13;
      let _init1 = ((_90_v1) => function (_91_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(_90_v1,new BigNumber(-176));
      })(_78_v1);
      let _nw19 = Array((new BigNumber(29)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw19.length); _i0_1++) {
        _nw19[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _89_v13 = _nw19;
      let _92_v14;
      _92_v14 = _dafny.Seq.of(_80_v3, _80_v3);
      let _93_globalState;
      let _nw20 = new _module.GlobalState();
      _nw20.__ctor(_77_v0, (_79_v2).Intersect(_79_v2), _80_v3, false, (_81_v4).dtor_cf0, new BigNumber(202), (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_83_v7).Elements) {
          let _94_v5 = _compr_0;
          if (_dafny.Seq.contains(_83_v7, _94_v5)) {
            _coll0.push([(_94_v5).multipliedBy(_82_v6),new _dafny.CodePoint('j'.codePointAt(0))]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(668)), ((_95_v6) => function (_96_i0) {
        return _95_v6;
      })(_82_v6))).length),_84_v8)), false, _dafny.Seq.Concat(_86_v10, _86_v10), false, new BigNumber(525), false, false, true, (_88_v12)[_module.__default.safeIndex(_82_v6, new BigNumber((_88_v12).length))], _89_v13, new BigNumber(588), false, (_92_v14)[_module.__default.safeIndex(_82_v6, new BigNumber((_92_v14).length))], _77_v0, true, true, new BigNumber(173), true);
      _93_globalState = _nw20;
      let _source0 = _81_v4;
      if (_source0.is_DC1) {
        let _97___mcc_h0 = (_source0).cf1;
        let _98___mcc_h1 = (_source0).cf2;
        let _99___mcc_h2 = (_source0).cf3;
        let _100_cf3 = _99___mcc_h2;
        let _101_cf2 = _98___mcc_h1;
        let _102_cf1 = _97___mcc_h0;
        let _103_v16;
        _103_v16 = _dafny.MultiSet.fromElements(_77_v0, _77_v0);
        let _104_v17;
        _104_v17 = _module.D0.create_DC1(_78_v1, _78_v1, _100_cf3);
        let _pat_let_tv0 = _101_cf2;
        let _pat_let_tv1 = _82_v6;
        let _rhs10 = (new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-99), new BigNumber(854))) {
            let _105_v15 = _compr_1;
            if (((new BigNumber(-99)).isLessThanOrEqualTo(_105_v15)) && ((_105_v15).isLessThan(new BigNumber(854)))) {
              _coll1.push([(_105_v15).multipliedBy(_82_v6),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-372)), ((_106_v8) => function (_107_i2) {
                return _106_v8;
              })(_84_v8))).length)]);
            }
          }
          return _coll1;
        }()).length)).minus(new BigNumber((_103_v16).cardinality()));
        let _rhs11 = ((_101_cf2) ? (_module.__default.safeModuloInt(_82_v6, _82_v6)) : (new BigNumber(620)));
        let _rhs12 = _module.__default.safeModuloInt(((_102_cf1) ? (_82_v6) : ((_83_v7)[_module.__default.safeIndex(_82_v6, new BigNumber((_83_v7).length))])), _100_cf3);
        let _rhs13 = (function (_pat_let0_0) {
          return function (_110_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_111_dt__update_hcf3_h0) {
                return function (_pat_let4_0) {
                  return function (_112_dt__update_hcf2_h1) {
                    return _module.D0.create_DC1((_110_dt__update__tmp_h1).dtor_cf1, _112_dt__update_hcf2_h1, _111_dt__update_hcf3_h0);
                  }(_pat_let4_0);
                }(false);
              }(_pat_let3_0);
            }((_dafny.ZERO).minus(_pat_let_tv1));
          }(_pat_let0_0);
        }(function (_pat_let1_0) {
          return function (_108_dt__update__tmp_h0) {
            return function (_pat_let2_0) {
              return function (_109_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_108_dt__update__tmp_h0).dtor_cf1, _109_dt__update_hcf2_h0, (_108_dt__update__tmp_h0).dtor_cf3);
              }(_pat_let2_0);
            }(_pat_let_tv0);
          }(_pat_let1_0);
        }(_104_v17))).dtor_cf3;
        let _rhs14 = ((!(_78_v1)) ? (_80_v3) : (_80_v3));
        let _lhs6 = _93_globalState;
        let _lhs7 = _93_globalState;
        _lhs6.f5 = _rhs10;
        _100_cf3 = _rhs11;
        _lhs7.f5 = _rhs12;
        _100_cf3 = _rhs13;
        _80_v3 = _rhs14;
        (_93_globalState).f5 = _82_v6;
        let _113_v18;
        _113_v18 = _dafny.Map.Empty.slice().updateUnsafe(_101_cf2,new BigNumber(-284));
        let _114_v19;
        let _nw21 = new _module.C0();
        _nw21.__ctor(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_78_v1,new BigNumber(-22))).Merge(_113_v18)).length), _101_cf2, _77_v0);
        _114_v19 = _nw21;
        let _115_v20;
        let _nw22 = new _module.C0();
        _nw22.__ctor(_82_v6, _101_cf2, _dafny.Seq.UnicodeFromString("nlwarcww"));
        _115_v20 = _nw22;
        let _116_v21;
        _116_v21 = _dafny.Map.Empty.slice().updateUnsafe(_115_v20,_100_cf3);
        (_93_globalState).f5 = (((_116_v21).contains(_115_v20)) ? ((_116_v21).get(_115_v20)) : (_100_cf3));
      } else if (_source0.is_DC2) {
        let _117___mcc_h3 = (_source0).cf4;
        let _118___mcc_h4 = (_source0).cf5;
        let _119_cf5 = _118___mcc_h4;
        let _120_cf4 = _117___mcc_h3;
        let _121_v22;
        let _nw23 = new _module.C0();
        _nw23.__ctor(new BigNumber((_83_v7).length), _78_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_122_v8) => function (_123_i3) {
          return _122_v8;
        })(_84_v8)));
        _121_v22 = _nw23;
        let _124_v23;
        _124_v23 = _dafny.Map.Empty.slice().updateUnsafe(_121_v22,_82_v6);
        let _125_v24;
        let _nw24 = new _module.C0();
        _nw24.__ctor(new BigNumber((_124_v23).length), _78_v1, _77_v0);
        _125_v24 = _nw24;
        let _126_v25;
        _126_v25 = _module.D1.create_DC4(_121_v22);
        let _127_v26;
        _127_v26 = _dafny.Map.Empty.slice().updateUnsafe((_126_v25).dtor_cf7,_121_v22);
        let _128_v27;
        _128_v27 = _dafny.Seq.of(_127_v26, _127_v26, (_127_v26).Merge(_127_v26), (_127_v26).Merge(_127_v26), _127_v26);
        _128_v27 = _128_v27;
        (_93_globalState).f20 = _78_v1;
        let _index8 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_80_v3).length));
        (_80_v3)[_index8] = _78_v1;
        let _129_v28;
        _129_v28 = _dafny.Seq.of(_78_v1, _78_v1);
        let _index9 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_80_v3).length));
        (_80_v3)[_index9] = (_129_v28)[_module.__default.safeIndex(_82_v6, new BigNumber((_129_v28).length))];
        let _130_v29;
        _130_v29 = _dafny.Set.fromElements(_82_v6, (_121_v22).f26);
        let _131_v30;
        _131_v30 = _dafny.Map.Empty.slice().updateUnsafe(_119_cf5,!((_80_v3)[_module.__default.safeIndex(new BigNumber(895), new BigNumber((_80_v3).length))]));
        let _rhs15 = _module.__default.fm4(new BigNumber((_130_v29).length), _82_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), ((_132_v8) => function (_133_i4) {
          return _132_v8;
        })(_84_v8)), _dafny.Seq.Concat(_dafny.Seq.of(_78_v1, _78_v1), _129_v28), _93_globalState);
        let _rhs16 = ((_125_v24).f26).multipliedBy(new BigNumber((_131_v30).length));
        let _lhs8 = _93_globalState;
        _84_v8 = _rhs15;
        _lhs8.f5 = _rhs16;
      } else if (_source0.is_DC0) {
        let _134___mcc_h5 = (_source0).cf0;
        let _135_cf0 = _134___mcc_h5;
        (_93_globalState).f5 = new BigNumber((_77_v0).length);
        let _136_v31;
        let _nw25 = new _module.C0();
        _nw25.__ctor(_82_v6, _78_v1, _135_cf0);
        _136_v31 = _nw25;
        let _137_v32;
        _137_v32 = _dafny.Seq.of(_136_v31);
        let _138_v33;
        _138_v33 = _dafny.Seq.of(_137_v32);
        _138_v33 = _dafny.Seq.of(_137_v32, _dafny.Seq.update(_137_v32, _module.__default.safeIndex(new BigNumber(477), new BigNumber((_137_v32).length)), _136_v31), _137_v32);
        _80_v3 = ((!(_78_v1)) ? (_80_v3) : (_80_v3));
        let _139_v34;
        let _140_v35;
        let _141_v36;
        let _142_v37;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(_93_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _139_v34 = _out0;
        _140_v35 = _out1;
        _141_v36 = _out2;
        _142_v37 = _out3;
      } else {
        let _143___mcc_h6 = (_source0).cf6;
        let _144_cf6 = _143___mcc_h6;
        let _source1 = _module.D0.create_DC0(((true) ? (_77_v0) : (_77_v0)));
        if (_source1.is_DC1) {
          let _145___mcc_h7 = (_source1).cf1;
          let _146___mcc_h8 = (_source1).cf2;
          let _147___mcc_h9 = (_source1).cf3;
          let _148_cf3 = _147___mcc_h9;
          let _149_cf2 = _146___mcc_h8;
          let _150_cf1 = _145___mcc_h7;
          let _151_v38;
          _151_v38 = _dafny.Seq.of(!(_149_cf2));
          _77_v0 = _dafny.Seq.update(_dafny.Seq.Concat(_77_v0, _77_v0), _module.__default.safeIndex(_module.__default.safeModuloInt(_82_v6, (_dafny.ZERO).minus(_82_v6)), new BigNumber((_dafny.Seq.Concat(_77_v0, _77_v0)).length)), _module.__default.fm4(new BigNumber((_77_v0).length), _148_cf3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), ((_152_v8) => function (_153_i5) {
            return _152_v8;
          })(_84_v8)), _151_v38, _93_globalState));
          _149_cf2 = _module.__default.fm5(!(_78_v1) || (_150_cf1), _150_cf1, _148_cf3, _82_v6, _93_globalState);
          _149_cf2 = _78_v1;
          let _154_v39;
          let _nw26 = new _module.C0();
          _nw26.__ctor(_82_v6, _150_cf1, _77_v0);
          _154_v39 = _nw26;
          let _155_v40;
          _155_v40 = _dafny.Map.Empty.slice().updateUnsafe(_154_v39,_module.__default.fm5(true, _149_cf2, _148_cf3, _148_cf3, _93_globalState));
          (_93_globalState).f5 = _module.__default.safeModuloInt(_module.__default.fm6(!(_78_v1), _149_cf2, _149_cf2, _93_globalState), new BigNumber(((_155_v40).update(_154_v39, false)).length));
        } else if (_source1.is_DC2) {
          let _156___mcc_h10 = (_source1).cf4;
          let _157___mcc_h11 = (_source1).cf5;
          let _158_cf5 = _157___mcc_h11;
          let _159_cf4 = _156___mcc_h10;
          let _160_v41;
          _160_v41 = _dafny.MultiSet.fromElements(_78_v1, false);
          let _161_v42;
          let _nw27 = new _module.C0();
          _nw27.__ctor(new BigNumber((_160_v41).cardinality()), _78_v1, _77_v0);
          _161_v42 = _nw27;
          let _162_v43;
          _162_v43 = _module.D1.create_DC6(_158_cf5, _161_v42, !(true), _78_v1);
          let _163_v44;
          _163_v44 = _dafny.Map.Empty.slice().updateUnsafe(false,_84_v8);
          let _164_v45;
          _164_v45 = _dafny.Map.Empty.slice().updateUnsafe(_82_v6,_module.__default.fm5(_78_v1, (_161_v42).f24, new BigNumber(609), _158_cf5, _93_globalState));
          let _165_v46;
          _165_v46 = _dafny.Map.Empty.slice().updateUnsafe(_158_cf5,_158_cf5);
          let _166_v48;
          let _nw28 = new _module.C0();
          _nw28.__ctor(_158_cf5, !(!(_78_v1)), (_161_v42).f25);
          _166_v48 = _nw28;
          let _167_v49;
          _167_v49 = _dafny.Seq.of(_166_v48, _166_v48, _166_v48);
          let _168_v50;
          _168_v50 = _module.D0.create_DC2(_159_cf4, _82_v6);
          let _169_v51;
          _169_v51 = _dafny.MultiSet.fromElements(_165_v46, _dafny.Map.Empty.slice().updateUnsafe(_82_v6,new BigNumber((_83_v7).length)), function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.Set.fromElements(new BigNumber((_167_v49).length), new BigNumber((_dafny.Seq.of((_168_v50).dtor_cf5)).length), (_166_v48).f26, (_162_v43).dtor_cf10, (_166_v48).f26)).Elements) {
              let _170_v47 = _compr_2;
              if ((_dafny.Set.fromElements(new BigNumber((_167_v49).length), new BigNumber((_dafny.Seq.of((_168_v50).dtor_cf5)).length), (_166_v48).f26, (_162_v43).dtor_cf10, (_166_v48).f26)).contains(_170_v47)) {
                _coll2.push([_module.__default.safeDivisionInt(_170_v47, _158_cf5),_82_v6]);
              }
            }
            return _coll2;
          }(), _165_v46);
          let _171_v52;
          _171_v52 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,new BigNumber((_169_v51).cardinality()));
          let _pat_let_tv2 = _161_v42;
          let _pat_let_tv3 = _78_v1;
          let _pat_let_tv4 = _161_v42;
          let _pat_let_tv5 = _158_cf5;
          (_93_globalState).f20 = ((_module.__default.fm5((function (_pat_let5_0) {
            return function (_175_dt__update__tmp_h3) {
              return function (_pat_let9_0) {
                return function (_176_dt__update_hcf13_h0) {
                  return function (_pat_let10_0) {
                    return function (_177_dt__update_hcf11_h1) {
                      return function (_pat_let11_0) {
                        return function (_178_dt__update_hcf10_h1) {
                          return _module.D1.create_DC6(_178_dt__update_hcf10_h1, _177_dt__update_hcf11_h1, (_175_dt__update__tmp_h3).dtor_cf12, _176_dt__update_hcf13_h0);
                        }(_pat_let11_0);
                      }(_pat_let_tv5);
                    }(_pat_let10_0);
                  }(_pat_let_tv4);
                }(_pat_let9_0);
              }(!(_pat_let_tv3));
            }(_pat_let5_0);
          }(function (_pat_let6_0) {
            return function (_172_dt__update__tmp_h2) {
              return function (_pat_let7_0) {
                return function (_173_dt__update_hcf11_h0) {
                  return function (_pat_let8_0) {
                    return function (_174_dt__update_hcf10_h0) {
                      return _module.D1.create_DC6(_174_dt__update_hcf10_h0, _173_dt__update_hcf11_h0, (_172_dt__update__tmp_h2).dtor_cf12, (_172_dt__update__tmp_h2).dtor_cf13);
                    }(_pat_let8_0);
                  }(new BigNumber(660));
                }(_pat_let7_0);
              }(_pat_let_tv2);
            }(_pat_let6_0);
          }(_162_v43))).dtor_cf12, _78_v1, _module.__default.fm6((_161_v42).fm1(_163_v44, _164_v45, _83_v7, _93_globalState), _78_v1, _78_v1, _93_globalState), _158_cf5, _93_globalState)) ? ((_161_v42).f24) : (!(_171_v52).equals(_171_v52)));
          let _179_v53;
          _179_v53 = _dafny.Set.fromElements(_158_cf5);
          let _180_v54;
          _180_v54 = _dafny.Seq.of((_161_v42).f24, (_161_v42).f24);
          let _rhs17 = ((!(_78_v1)) ? (_179_v53) : ((_dafny.Set.fromElements((_166_v48).f26)).Difference(_179_v53)));
          let _rhs18 = false;
          let _rhs19 = _dafny.areEqual(_180_v54, _dafny.Seq.Concat(_180_v54, _180_v54));
          let _rhs20 = (_161_v42).f24;
          let _lhs9 = _93_globalState;
          let _lhs10 = _93_globalState;
          let _lhs11 = _93_globalState;
          _179_v53 = _rhs17;
          _lhs9.f20 = _rhs18;
          _lhs10.f21 = _rhs19;
          _lhs11.f21 = _rhs20;
          let _181_v55;
          _181_v55 = _dafny.Map.Empty.slice().updateUnsafe(_179_v53,(_166_v48).f26);
          let _182_v56;
          let _nw29 = new _module.C0();
          _nw29.__ctor((_166_v48).f26, (_181_v55).equals(_181_v55), _77_v0);
          _182_v56 = _nw29;
          let _183_v57;
          let _184_v58;
          let _185_v59;
          let _186_v60;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(_93_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _183_v57 = _out4;
          _184_v58 = _out5;
          _185_v59 = _out6;
          _186_v60 = _out7;
        } else if (_source1.is_DC0) {
          let _187___mcc_h12 = (_source1).cf0;
          let _188_cf0 = _187___mcc_h12;
          (_93_globalState).f20 = _78_v1;
          let _189_v61;
          let _nw30 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _189_v61 = _nw30;
          let _index10 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_189_v61).length));
          (_189_v61)[_index10] = _82_v6;
          let _190_v62;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_82_v6, false, _188_cf0);
          _190_v62 = _nw31;
          let _191_v63;
          _191_v63 = _dafny.Seq.of(_190_v62, _190_v62);
          let _192_v64;
          _192_v64 = _dafny.Seq.of(_191_v63);
          let _index11 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_189_v61).length));
          (_189_v61)[_index11] = (_82_v6).plus(((false) ? (new BigNumber((_dafny.Seq.update((_192_v64)[_module.__default.safeIndex(new BigNumber(-355), new BigNumber((_192_v64).length))], _module.__default.safeIndex((_190_v62).fm3(_82_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), ((_193_v7) => function (_194_i6) {
            return _193_v7;
          })(_83_v7)), _82_v6, _dafny.Seq.update(_188_cf0, _module.__default.safeIndex(_82_v6, new BigNumber((_188_cf0).length)), _84_v8), _93_globalState), new BigNumber(((_192_v64)[_module.__default.safeIndex(new BigNumber(-355), new BigNumber((_192_v64).length))]).length)), _190_v62)).length)) : ((_dafny.ZERO).minus(new BigNumber((_188_cf0).length)))));
          _78_v1 = (new BigNumber((_188_cf0).length)).isLessThan(new BigNumber(659));
          (_93_globalState).f21 = _78_v1;
        } else {
          let _195___mcc_h13 = (_source1).cf6;
          let _196_cf6 = _195___mcc_h13;
          let _197_v65;
          let _198_v66;
          let _199_v67;
          let _200_v68;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(_93_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _197_v65 = _out8;
          _198_v66 = _out9;
          _199_v67 = _out10;
          _200_v68 = _out11;
          (_93_globalState).f5 = _199_v67;
          _78_v1 = !(_78_v1);
          _82_v6 = _module.__default.safeModuloInt(_199_v67, new BigNumber(874));
        }
        _78_v1 = false;
        let _201_v69;
        _201_v69 = _dafny.Seq.of(_78_v1);
        let _202_v71;
        let _nw32 = new _module.C0();
        _nw32.__ctor(new BigNumber(953), _78_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-753)), ((_203_v8) => function (_204_i7) {
          return _203_v8;
        })(_84_v8)));
        _202_v71 = _nw32;
        let _205_v72;
        _205_v72 = _dafny.Seq.of(_202_v71, _202_v71, _202_v71, _202_v71);
        let _206_v73;
        _206_v73 = _dafny.Map.Empty.slice().updateUnsafe((_202_v71).f26,_82_v6);
        let _207_v74;
        let _nw33 = Array((new BigNumber(25)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _82_v6;
        _nw33[(_dafny.ONE).toNumber()] = _82_v6;
        _nw33[(new BigNumber(2)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(3)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(4)).toNumber()] = new BigNumber((_201_v69).length);
        _nw33[(new BigNumber(5)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(6)).toNumber()] = new BigNumber(225);
        _nw33[(new BigNumber(7)).toNumber()] = new BigNumber((_77_v0).length);
        _nw33[(new BigNumber(8)).toNumber()] = new BigNumber(373);
        _nw33[(new BigNumber(9)).toNumber()] = new BigNumber((function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(970), new BigNumber(840))) {
            let _208_v70 = _compr_3;
            if (((new BigNumber(970)).isLessThanOrEqualTo(_208_v70)) && ((_208_v70).isLessThan(new BigNumber(840)))) {
              _coll3.add((_208_v70).multipliedBy(_82_v6));
            }
          }
          return _coll3;
        }()).length);
        _nw33[(new BigNumber(10)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(11)).toNumber()] = new BigNumber((_87_v11).cardinality());
        _nw33[(new BigNumber(12)).toNumber()] = new BigNumber((_77_v0).length);
        _nw33[(new BigNumber(13)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(14)).toNumber()] = new BigNumber((_205_v72).length);
        _nw33[(new BigNumber(15)).toNumber()] = (_202_v71).f26;
        _nw33[(new BigNumber(16)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(17)).toNumber()] = new BigNumber(768);
        _nw33[(new BigNumber(18)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(19)).toNumber()] = (_202_v71).f26;
        _nw33[(new BigNumber(20)).toNumber()] = _82_v6;
        _nw33[(new BigNumber(21)).toNumber()] = (_202_v71).f26;
        _nw33[(new BigNumber(22)).toNumber()] = new BigNumber((_206_v73).length);
        _nw33[(new BigNumber(23)).toNumber()] = new BigNumber(-224);
        _nw33[(new BigNumber(24)).toNumber()] = _82_v6;
        _207_v74 = _nw33;
        let _209_v75;
        _209_v75 = _module.D0.create_DC2(_207_v74, _82_v6);
        let _source2 = _209_v75;
        if (_source2.is_DC1) {
          let _210___mcc_h14 = (_source2).cf1;
          let _211___mcc_h15 = (_source2).cf2;
          let _212___mcc_h16 = (_source2).cf3;
          let _213_cf3 = _212___mcc_h16;
          let _214_cf2 = _211___mcc_h15;
          let _215_cf1 = _210___mcc_h14;
          let _216_v76;
          _216_v76 = _dafny.Map.Empty.slice().updateUnsafe(((_202_v71).f26).multipliedBy((_83_v7)[_module.__default.safeIndex(_213_cf3, new BigNumber((_83_v7).length))]),_80_v3);
          let _217_v77;
          _217_v77 = _dafny.Map.Empty.slice().updateUnsafe(_215_cf1,new BigNumber(-730));
          _216_v76 = (_216_v76).update(new BigNumber((_217_v77).length), _80_v3);
          let _218_v78;
          _218_v78 = _dafny.Map.Empty.slice().updateUnsafe(((_202_v71).f26).minus(_213_cf3),_214_cf2);
          _218_v78 = (_218_v78).update((_202_v71).f26, !(_215_cf1));
          _214_cf2 = _215_cf1;
          let _219_v79;
          let _220_v80;
          let _221_v81;
          let _222_v82;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m0(_93_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _219_v79 = _out12;
          _220_v80 = _out13;
          _221_v81 = _out14;
          _222_v82 = _out15;
        } else if (_source2.is_DC2) {
          let _223___mcc_h17 = (_source2).cf4;
          let _224___mcc_h18 = (_source2).cf5;
          let _225_cf5 = _224___mcc_h18;
          let _226_cf4 = _223___mcc_h17;
          let _227_v83;
          _227_v83 = _dafny.Map.Empty.slice().updateUnsafe(_79_v2,new BigNumber((_79_v2).length));
          let _228_v84;
          _228_v84 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_82_v6);
          let _229_v85;
          _229_v85 = _dafny.MultiSet.fromElements(_228_v84);
          let _230_v86;
          _230_v86 = _dafny.Map.Empty.slice().updateUnsafe((((_227_v83).contains(_79_v2)) ? ((_227_v83).get(_79_v2)) : (new BigNumber((_83_v7).length))),(_229_v85).IsSubsetOf((_229_v85).update(_228_v84, _module.__default.abs((_202_v71).f26))));
          _230_v86 = (_230_v86).update((_202_v71).f26, false);
          let _231_v87;
          let _nw34 = new _module.C0();
          _nw34.__ctor((_dafny.ZERO).minus((_202_v71).f26), _78_v1, _77_v0);
          _231_v87 = _nw34;
          let _index12 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_85_v9).length));
          (_85_v9)[_index12] = _84_v8;
          let _index13 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_85_v9).length));
          (_85_v9)[_index13] = _84_v8;
          _79_v2 = _79_v2;
        } else if (_source2.is_DC0) {
          let _232___mcc_h19 = (_source2).cf0;
          let _233_cf0 = _232___mcc_h19;
          let _234_v88;
          let _235_v89;
          let _236_v90;
          let _237_v91;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(_93_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _234_v88 = _out16;
          _235_v89 = _out17;
          _236_v90 = _out18;
          _237_v91 = _out19;
          let _238_v92;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_236_v90, _78_v1, _module.__default.fm7(_78_v1, _236_v90, (_83_v7)[_module.__default.safeIndex(_236_v90, new BigNumber((_83_v7).length))], _237_v91, _93_globalState));
          _238_v92 = _nw35;
          let _239_v93;
          _239_v93 = _dafny.Map.Empty.slice().updateUnsafe(_236_v90,_237_v91);
          _239_v93 = (_239_v93).update(((_202_v71).f26).plus(_82_v6), !((new BigNumber((_201_v69).length)).isLessThan(_82_v6)));
          (_93_globalState).f20 = _237_v91;
        } else {
          let _240___mcc_h20 = (_source2).cf6;
          let _241_cf6 = _240___mcc_h20;
          let _242_v94;
          let _nw36 = Array((new BigNumber(13)).toNumber()).fill(_module.D1.Default());
          _242_v94 = _nw36;
          let _243_v95;
          let _nw37 = new _module.C0();
          _nw37.__ctor((_202_v71).f26, !(_78_v1), _77_v0);
          _243_v95 = _nw37;
          let _index14 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_242_v94).length));
          (_242_v94)[_index14] = _module.D1.create_DC6((_202_v71).f26, _243_v95, _78_v1, (_243_v95).f24);
          let _244_v96;
          _244_v96 = _dafny.Map.Empty.slice().updateUnsafe((_202_v71).f26,true);
          let _index15 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_242_v94).length));
          (_242_v94)[_index15] = _module.D1.create_DC6((_202_v71).f26, _243_v95, (_243_v95).f24, (_243_v95).fm1(_dafny.Map.Empty.slice().updateUnsafe((_243_v95).f24,_84_v8), _244_v96, _83_v7, _93_globalState));
          let _245_v98;
          _245_v98 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-822), new BigNumber(224))) {
              let _246_v97 = _compr_4;
              if (((new BigNumber(-822)).isLessThanOrEqualTo(_246_v97)) && ((_246_v97).isLessThan(new BigNumber(224)))) {
                _coll4.push([_module.__default.safeDivisionInt(_246_v97, (_202_v71).f26),(_243_v95).f24]);
              }
            }
            return _coll4;
          }(),_dafny.MultiSet.fromElements((_243_v95).f24, _78_v1, _78_v1));
          let _247_v99;
          _247_v99 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_77_v0, _module.__default.safeIndex(new BigNumber((_245_v98).length), new BigNumber((_77_v0).length)), new _dafny.CodePoint('k'.codePointAt(0))));
          _247_v99 = (_247_v99).Intersect(_247_v99);
          _78_v1 = (new BigNumber(551)).isLessThanOrEqualTo(((_202_v71).f26).multipliedBy(new BigNumber(564)));
          let _248_v100;
          let _nw38 = Array((new BigNumber(29)).toNumber()).fill([]);
          _248_v100 = _nw38;
          let _249_v101;
          _249_v101 = _module.D2.create_DC8(_80_v3);
          let _250_v102;
          let _nw39 = Array((new BigNumber(14)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _80_v3;
          _nw39[(_dafny.ONE).toNumber()] = _80_v3;
          _nw39[(new BigNumber(2)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(3)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(4)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(5)).toNumber()] = (_249_v101).dtor_cf15;
          _nw39[(new BigNumber(6)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(7)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(8)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(9)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(10)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(11)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(12)).toNumber()] = _80_v3;
          _nw39[(new BigNumber(13)).toNumber()] = _80_v3;
          _250_v102 = _nw39;
          let _index16 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_248_v100).length));
          (_248_v100)[_index16] = _250_v102;
          let _index17 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_248_v100).length));
          (_248_v100)[_index17] = (((_243_v95).f24) ? (_250_v102) : (_250_v102));
        }
        let _251_v103;
        _251_v103 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_84_v8);
        let _252_v104;
        _252_v104 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm6(_78_v1, _78_v1, _78_v1, _93_globalState)),_201_v69);
        (_93_globalState).f21 = ((((_202_v71).fm1(_251_v103, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_252_v104).length),_78_v1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_253_v6) => function (_254_i8) {
          return _253_v6;
        })(_82_v6)), _93_globalState)) ? (_78_v1) : (!(false)))) && (_78_v1);
      }
      let _255_v105;
      let _nw40 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
      _255_v105 = _nw40;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_255_v105).length))) {
        let _256_i9 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_256_i9)) && ((_256_i9).isLessThan(new BigNumber((_255_v105).length))))) {
          (_255_v105)[(_256_i9)] = (_dafny.Map.Empty.slice().updateUnsafe(_78_v1,_84_v8)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_78_v1,new _dafny.CodePoint('t'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_78_v1,_84_v8)));
        }
      }
      let _257_v106;
      _257_v106 = _dafny.Seq.of(_78_v1, _78_v1);
      (_93_globalState).f21 = (!((_257_v106)[_module.__default.safeIndex(new BigNumber((_77_v0).length), new BigNumber((_257_v106).length))])) === (_78_v1);
      let _258_v107;
      _258_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(49),_82_v6);
      let _259_v108;
      _259_v108 = _module.D1.create_DC5(_78_v1, _82_v6);
      let _pat_let_tv6 = _82_v6;
      let _pat_let_tv7 = _82_v6;
      let _pat_let_tv8 = _82_v6;
      let _pat_let_tv9 = _82_v6;
      let _pat_let_tv10 = _82_v6;
      let _pat_let_tv11 = _82_v6;
      let _rhs21 = (_78_v1) === (_78_v1);
      let _rhs22 = (((_258_v107).contains(_82_v6)) ? ((_258_v107).get(_82_v6)) : (_82_v6));
      let _rhs23 = function (_source3) {
        if (_source3.is_DC5) {
          let _260___mcc_h21 = (_source3).cf8;
          let _261___mcc_h22 = (_source3).cf9;
          let _262_cf9 = _261___mcc_h22;
          let _263_cf8 = _260___mcc_h21;
          return (_pat_let_tv6).multipliedBy(_pat_let_tv7);
        } else if (_source3.is_DC6) {
          let _264___mcc_h23 = (_source3).cf10;
          let _265___mcc_h24 = (_source3).cf11;
          let _266___mcc_h25 = (_source3).cf12;
          let _267___mcc_h26 = (_source3).cf13;
          let _268_cf13 = _267___mcc_h26;
          let _269_cf12 = _266___mcc_h25;
          let _270_cf11 = _265___mcc_h24;
          let _271_cf10 = _264___mcc_h23;
          return (_pat_let_tv8).multipliedBy(_pat_let_tv9);
        } else if (_source3.is_DC4) {
          let _272___mcc_h27 = (_source3).cf7;
          let _273_cf7 = _272___mcc_h27;
          return (_273_cf7).f26;
        } else {
          let _274___mcc_h28 = (_source3).cf14;
          let _275_cf14 = _274___mcc_h28;
          return new BigNumber((function () {
            let _coll5 = new _dafny.Set();
            for (const _compr_5 of (_dafny.Set.fromElements(_pat_let_tv10)).Elements) {
              let _276_v109 = _compr_5;
              if ((_dafny.Set.fromElements(_pat_let_tv11)).contains(_276_v109)) {
                _coll5.add(_module.__default.safeDivisionInt(_276_v109, new BigNumber(53)));
              }
            }
            return _coll5;
          }()).length);
        }
      }(_259_v108);
      let _lhs12 = _93_globalState;
      _lhs12.f20 = _rhs21;
      _82_v6 = _rhs22;
      _82_v6 = _rhs23;
      let _277_v110;
      let _nw41 = new _module.C0();
      _nw41.__ctor(_82_v6, _78_v1, _77_v0);
      _277_v110 = _nw41;
      let _278_v111;
      _278_v111 = _module.D1.create_DC6(_82_v6, _277_v110, true, (_277_v110).f24);
      let _279_v112;
      _279_v112 = _dafny.Map.Empty.slice().updateUnsafe(false,(_82_v6).isLessThan((_278_v111).dtor_cf10));
      _279_v112 = (_279_v112).update((_277_v110).f24, !(_82_v6).isEqualTo(_82_v6));
      let _hi0 = _82_v6;
      for (let _280_i10 = new BigNumber(658); _280_i10.isLessThan(_hi0); _280_i10 = _280_i10.plus(_dafny.ONE)) {
        _82_v6 = _280_i10;
        let _281_v113;
        let _nw42 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _281_v113 = _nw42;
        _281_v113 = _281_v113;
        let _282_v114;
        _282_v114 = _dafny.Map.Empty.slice().updateUnsafe(_82_v6,_277_v110);
        _282_v114 = (_282_v114).update((_dafny.ZERO).minus(new BigNumber((_83_v7).length)), _277_v110);
        _279_v112 = (_279_v112).update(_78_v1, (_257_v106)[_module.__default.safeIndex(_82_v6, new BigNumber((_257_v106).length))]);
      }
      let _283_v115;
      _283_v115 = _dafny.Set.fromElements(_84_v8);
      let _284_v116;
      _284_v116 = _dafny.Set.fromElements(_283_v115);
      let _285_i11;
      _285_i11 = _dafny.ZERO;
      L0: {
        while (((_284_v116).Union(_dafny.Set.fromElements(_283_v115, _283_v115))).IsProperSubsetOf(_284_v116)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_285_i11)) {
              break L0;
            }
            _285_i11 = (_285_i11).plus(_dafny.ONE);
            let _286_v117;
            let _nw43 = Array((_dafny.ONE).toNumber());
            _nw43[(_dafny.ZERO).toNumber()] = _82_v6;
            _286_v117 = _nw43;
            let _index18 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_286_v117).length));
            (_286_v117)[_index18] = _82_v6;
            let _287_v118;
            _287_v118 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((_277_v110).f25, _77_v0),_82_v6);
            let _index19 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_286_v117).length));
            let _rhs24 = (_257_v106)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_82_v6, _82_v6), new BigNumber((_257_v106).length))];
            let _rhs25 = (((_287_v118).contains((_277_v110).f25)) ? ((_287_v118).get((_277_v110).f25)) : (_82_v6));
            let _rhs26 = (new BigNumber(674)).isEqualTo(_module.__default.fm6(_78_v1, (_277_v110).f24, true, _93_globalState));
            let _lhs13 = _93_globalState;
            let _lhs14 = _286_v117;
            let _lhs15 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_286_v117).length));
            let _lhs16 = _93_globalState;
            _lhs13.f21 = _rhs24;
            _lhs14[_lhs15] = _rhs25;
            _lhs16.f21 = _rhs26;
            let _288_v119;
            let _289_v120;
            let _290_v121;
            let _291_v122;
            let _out20;
            let _out21;
            let _out22;
            let _out23;
            let _outcollector5 = _module.__default.m0(_93_globalState);
            _out20 = _outcollector5[0];
            _out21 = _outcollector5[1];
            _out22 = _outcollector5[2];
            _out23 = _outcollector5[3];
            _288_v119 = _out20;
            _289_v120 = _out21;
            _290_v121 = _out22;
            _291_v122 = _out23;
            _78_v1 = ((new BigNumber(54)).plus(new BigNumber((_83_v7).length))).isEqualTo(_290_v121);
            let _292_v123;
            let _nw44 = new _module.C0();
            _nw44.__ctor((_dafny.ZERO).minus(_82_v6), _291_v122, _288_v119);
            _292_v123 = _nw44;
            let _293_v124;
            let _nw45 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
            _293_v124 = _nw45;
            let _294_v125;
            _294_v125 = _dafny.Map.Empty.slice().updateUnsafe(_87_v11,!((_277_v110).f24));
            let _index20 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_293_v124).length));
            (_293_v124)[_index20] = _294_v125;
            let _index21 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_293_v124).length));
            let _rhs27 = _292_v123;
            let _rhs28 = (function () {
              let _coll6 = new _dafny.Map();
              for (const _compr_6 of (_88_v12).Elements) {
                let _295_v126 = _compr_6;
                if (_dafny.Seq.contains(_88_v12, _295_v126)) {
                  _coll6.push([_295_v126,(_277_v110).f24]);
                }
              }
              return _coll6;
            }()).Merge((_294_v125).Merge(_294_v125));
            let _rhs29 = _80_v3;
            let _rhs30 = (_286_v117)[_module.__default.safeIndex(new BigNumber(6), new BigNumber((_286_v117).length))];
            let _lhs17 = _293_v124;
            let _lhs18 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_293_v124).length));
            _292_v123 = _rhs27;
            _lhs17[_lhs18] = _rhs28;
            _80_v3 = _rhs29;
            _290_v121 = _rhs30;
          }
        }
      }
      let _296_v127;
      let _nw46 = Array((new BigNumber(16)).toNumber());
      _nw46[(_dafny.ZERO).toNumber()] = (_277_v110).f25;
      _nw46[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_77_v0, _dafny.Seq.update((_277_v110).f25, _module.__default.safeIndex(_82_v6, new BigNumber(((_277_v110).f25).length)), _84_v8));
      _nw46[(new BigNumber(2)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vhh"), _77_v0);
      _nw46[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat((_277_v110).f25, _77_v0), _module.__default.safeIndex(_82_v6, new BigNumber((_dafny.Seq.Concat((_277_v110).f25, _77_v0)).length)), _84_v8);
      _nw46[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_277_v110).f25, (_277_v110).f25);
      _nw46[(new BigNumber(6)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(7)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(8)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(9)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(10)).toNumber()] = _module.__default.fm7((_277_v110).f24, _82_v6, _82_v6, (_277_v110).f24, _93_globalState);
      _nw46[(new BigNumber(11)).toNumber()] = _77_v0;
      _nw46[(new BigNumber(12)).toNumber()] = _77_v0;
      _nw46[(new BigNumber(13)).toNumber()] = _77_v0;
      _nw46[(new BigNumber(14)).toNumber()] = (_277_v110).f25;
      _nw46[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), ((_297_v8) => function (_298_i12) {
        return _297_v8;
      })(_84_v8));
      _296_v127 = _nw46;
      let _index22 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length));
      (_296_v127)[_index22] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("chxreot"), _module.__default.safeIndex(_82_v6, new BigNumber((_dafny.Seq.UnicodeFromString("chxreot")).length)), _84_v8);
      let _index23 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length));
      (_296_v127)[_index23] = (_277_v110).f25;
      if (_78_v1) {
        _80_v3 = _80_v3;
        let _299_v128;
        let _300_v129;
        let _301_v130;
        let _302_v131;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0(_93_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _299_v128 = _out24;
        _300_v129 = _out25;
        _301_v130 = _out26;
        _302_v131 = _out27;
        (_93_globalState).f21 = !(!((new BigNumber(991)).isLessThan(_82_v6))) || (_78_v1);
        let _303_v132;
        let _304_v133;
        let _305_v134;
        let _306_v135;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0(_93_globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _303_v132 = _out28;
        _304_v133 = _out29;
        _305_v134 = _out30;
        _306_v135 = _out31;
        let _307_v136;
        let _nw47 = new _module.C0();
        _nw47.__ctor(new BigNumber((_303_v132).length), true, (_277_v110).f25);
        _307_v136 = _nw47;
        let _308_v137;
        _308_v137 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_307_v136);
        let _309_v138;
        _309_v138 = _dafny.MultiSet.fromElements((_308_v137).update(_302_v131, _307_v136));
        let _310_v139;
        _310_v139 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("dcoudebbi"), _303_v132);
        let _311_v140;
        _311_v140 = _dafny.Seq.of((_310_v139)[_module.__default.safeIndex(_305_v134, new BigNumber((_310_v139).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), ((_312_v8) => function (_313_i13) {
          return _312_v8;
        })(_84_v8)));
        let _314_v141;
        _314_v141 = _dafny.Map.Empty.slice().updateUnsafe(_80_v3,_dafny.Seq.UnicodeFromString("yckgtwu"));
        let _315_v142;
        let _nw48 = Array((new BigNumber(17)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _82_v6;
        _nw48[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_77_v0, (_277_v110).f25)).length);
        _nw48[(new BigNumber(2)).toNumber()] = (_305_v134).plus(_305_v134);
        _nw48[(new BigNumber(3)).toNumber()] = new BigNumber((_309_v138).cardinality());
        _nw48[(new BigNumber(4)).toNumber()] = _305_v134;
        _nw48[(new BigNumber(5)).toNumber()] = new BigNumber(((_277_v110).f25).length);
        _nw48[(new BigNumber(6)).toNumber()] = ((_78_v1) ? (new BigNumber(((_310_v139)[_module.__default.safeIndex(_82_v6, new BigNumber((_310_v139).length))]).length)) : (new BigNumber(995)));
        _nw48[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(_301_v130, _301_v130);
        _nw48[(new BigNumber(8)).toNumber()] = (new BigNumber(782)).plus(_301_v130);
        _nw48[(new BigNumber(9)).toNumber()] = (_82_v6).minus(_301_v130);
        _nw48[(new BigNumber(10)).toNumber()] = new BigNumber(929);
        _nw48[(new BigNumber(11)).toNumber()] = new BigNumber(((_277_v110).f25).length);
        _nw48[(new BigNumber(12)).toNumber()] = new BigNumber(507);
        _nw48[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_314_v141).length), _305_v134);
        _nw48[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), ((_316_v8) => function (_317_i14) {
          return _316_v8;
        })(_84_v8))).length), _305_v134);
        _nw48[(new BigNumber(15)).toNumber()] = _module.__default.fm6((_277_v110).f24, false, _306_v135, _93_globalState);
        _nw48[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_83_v7, _dafny.Seq.of(_301_v130))).length);
        _315_v142 = _nw48;
        let _index24 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_315_v142).length));
        (_315_v142)[_index24] = _305_v134;
        let _index25 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_315_v142).length));
        (_315_v142)[_index25] = (_82_v6).multipliedBy(_82_v6);
      } else {
        let _318_v143;
        _318_v143 = _module.D2.create_DC9(_78_v1, _82_v6);
        let _319_v144;
        let _nw49 = new _module.C0();
        _nw49.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_81_v4,(_277_v110).f24)).length), (_277_v110).f24, (_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]);
        _319_v144 = _nw49;
        let _320_v145;
        _320_v145 = _dafny.Set.fromElements(_319_v144);
        let _321_v146;
        _321_v146 = _dafny.Map.Empty.slice().updateUnsafe(_78_v1,_320_v145);
        let _322_v147;
        _322_v147 = _dafny.Map.Empty.slice().updateUnsafe((_277_v110).f25,_dafny.Map.Empty.slice().updateUnsafe(_318_v143,new BigNumber((_321_v146).length)));
        let _323_v149;
        _323_v149 = _dafny.Map.Empty.slice().updateUnsafe(_77_v0,_dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,new BigNumber(428)));
        _322_v147 = function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_323_v149).Keys.Elements) {
            let _324_v148 = _compr_7;
            if ((_323_v149).contains(_324_v148)) {
              _coll7.push([_324_v148,_dafny.Map.Empty.slice().updateUnsafe(_318_v143,new BigNumber(((_277_v110).f25).length))]);
            }
          }
          return _coll7;
        }();
        let _325_v150;
        let _326_v151;
        let _327_v152;
        let _328_v153;
        let _out32;
        let _out33;
        let _out34;
        let _out35;
        let _outcollector8 = _module.__default.m0(_93_globalState);
        _out32 = _outcollector8[0];
        _out33 = _outcollector8[1];
        _out34 = _outcollector8[2];
        _out35 = _outcollector8[3];
        _325_v150 = _out32;
        _326_v151 = _out33;
        _327_v152 = _out34;
        _328_v153 = _out35;
        let _329_v154;
        _329_v154 = _dafny.Map.Empty.slice().updateUnsafe(_327_v152,_84_v8);
        let _330_v155;
        _330_v155 = _module.D0.create_DC1(!(_78_v1), (_277_v110).f24, _327_v152);
        let _331_v156;
        _331_v156 = _dafny.Seq.of(_module.D1.create_DC6((_330_v155).dtor_cf3, _277_v110, true, (_277_v110).f24));
        _329_v154 = (_329_v154).update(new BigNumber((_331_v156).length), _module.__default.fm4(new BigNumber((_329_v154).length), (_319_v144).f26, _dafny.Seq.UnicodeFromString("g"), _257_v106, _93_globalState));
        if ((_277_v110).f24) {
          let _332_v157;
          let _333_v158;
          let _334_v159;
          let _335_v160;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector9 = _module.__default.m0(_93_globalState);
          _out36 = _outcollector9[0];
          _out37 = _outcollector9[1];
          _out38 = _outcollector9[2];
          _out39 = _outcollector9[3];
          _332_v157 = _out36;
          _333_v158 = _out37;
          _334_v159 = _out38;
          _335_v160 = _out39;
          let _336_v161;
          let _nw50 = new _module.C0();
          _nw50.__ctor(_334_v159, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("v"), (_277_v110).f25), _dafny.Seq.Concat(_332_v157, _dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_337_v8) => function (_338_i15) {
            return _337_v8;
          })(_84_v8))));
          _336_v161 = _nw50;
          let _339_v162;
          _339_v162 = _module.D0.create_DC1(_78_v1, (_257_v106)[_module.__default.safeIndex(_327_v152, new BigNumber((_257_v106).length))], new BigNumber((_279_v112).length));
          let _340_v163;
          _340_v163 = _module.D0.create_DC3(_339_v162);
          let _341_v164;
          _341_v164 = _module.D0.create_DC3(_339_v162);
          let _342_v165;
          _342_v165 = _module.D0.create_DC3(_339_v162);
          let _pat_let_tv12 = _340_v163;
          let _rhs31 = _335_v160;
          let _rhs32 = (_277_v110).f24;
          let _rhs33 = ((!(_78_v1)) ? (_327_v152) : (_334_v159));
          let _rhs34 = new BigNumber(-766);
          let _rhs35 = function (_pat_let12_0) {
            return function (_343_dt__update__tmp_h4) {
              return function (_pat_let13_0) {
                return function (_344_dt__update_hcf6_h0) {
                  return _module.D0.create_DC3(_344_dt__update_hcf6_h0);
                }(_pat_let13_0);
              }(_module.D0.create_DC3(_pat_let_tv12));
            }(_pat_let12_0);
          }(_342_v165);
          let _lhs19 = _93_globalState;
          let _lhs20 = _93_globalState;
          let _lhs21 = _93_globalState;
          _335_v160 = _rhs31;
          _lhs19.f20 = _rhs32;
          _lhs20.f5 = _rhs33;
          _lhs21.f5 = _rhs34;
          _342_v165 = _rhs35;
          (_93_globalState).f20 = _78_v1;
          let _345_v166;
          _345_v166 = _module.D3.create_DC11(_83_v7);
          _335_v160 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_345_v166).dtor_cf19, _83_v7), _dafny.Seq.Create(_module.__default.abs(new BigNumber(796)), ((_346_v161, _347_v110) => function (_348_i16) {
            return new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_346_v161).f26,(_347_v110).f24), function () {
              let _coll8 = new _dafny.Map();
              for (const _compr_8 of _dafny.IntegerRange(new BigNumber(636), new BigNumber(445))) {
                let _349_v167 = _compr_8;
                if (((new BigNumber(636)).isLessThanOrEqualTo(_349_v167)) && ((_349_v167).isLessThan(new BigNumber(445)))) {
                  _coll8.push([(_349_v167).multipliedBy((_346_v161).f26),(_347_v110).f24]);
                }
              }
              return _coll8;
            }())).cardinality());
          })(_336_v161, _277_v110)));
        } else {
          let _350_v168;
          let _351_v169;
          let _352_v170;
          let _353_v171;
          let _out40;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector10 = _module.__default.m0(_93_globalState);
          _out40 = _outcollector10[0];
          _out41 = _outcollector10[1];
          _out42 = _outcollector10[2];
          _out43 = _outcollector10[3];
          _350_v168 = _out40;
          _351_v169 = _out41;
          _352_v170 = _out42;
          _353_v171 = _out43;
          let _354_v172;
          let _355_v173;
          let _356_v174;
          let _357_v175;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector11 = _module.__default.m0(_93_globalState);
          _out44 = _outcollector11[0];
          _out45 = _outcollector11[1];
          _out46 = _outcollector11[2];
          _out47 = _outcollector11[3];
          _354_v172 = _out44;
          _355_v173 = _out45;
          _356_v174 = _out46;
          _357_v175 = _out47;
          let _358_v176;
          let _359_v177;
          let _360_v178;
          let _361_v179;
          let _out48;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector12 = _module.__default.m0(_93_globalState);
          _out48 = _outcollector12[0];
          _out49 = _outcollector12[1];
          _out50 = _outcollector12[2];
          _out51 = _outcollector12[3];
          _358_v176 = _out48;
          _359_v177 = _out49;
          _360_v178 = _out50;
          _361_v179 = _out51;
          (_93_globalState).f20 = (_318_v143).dtor_cf16;
          (_93_globalState).f21 = !((_277_v110).f24);
        }
        let _362_v180;
        let _nw51 = new _module.C0();
        _nw51.__ctor(_327_v152, (_278_v111).dtor_cf13, _dafny.Seq.UnicodeFromString("iwdsggla"));
        _362_v180 = _nw51;
        _362_v180 = _362_v180;
      }
      let _363_v181;
      _363_v181 = _dafny.Map.Empty.slice().updateUnsafe(_84_v8,_82_v6);
      let _364_v182;
      _364_v182 = _dafny.MultiSet.fromElements((_277_v110).f24);
      let _365_v183;
      _365_v183 = _dafny.Seq.of(_364_v182);
      let _366_v185;
      let _nw52 = Array((new BigNumber(18)).toNumber());
      _nw52[(_dafny.ZERO).toNumber()] = _82_v6;
      _nw52[(_dafny.ONE).toNumber()] = new BigNumber((_363_v181).length);
      _nw52[(new BigNumber(2)).toNumber()] = new BigNumber(37);
      _nw52[(new BigNumber(3)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(4)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(5)).toNumber()] = ((((_258_v107).contains(new BigNumber(((_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]).length))) ? ((_258_v107).get(new BigNumber(((_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]).length))) : (_82_v6))).multipliedBy(_82_v6);
      _nw52[(new BigNumber(6)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(_82_v6, new BigNumber(655));
      _nw52[(new BigNumber(8)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(_82_v6, _82_v6);
      _nw52[(new BigNumber(10)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(11)).toNumber()] = new BigNumber(((_277_v110).f25).length);
      _nw52[(new BigNumber(12)).toNumber()] = new BigNumber(201);
      _nw52[(new BigNumber(13)).toNumber()] = (((_277_v110).f24) ? (_82_v6) : (new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_365_v183).Elements) {
          let _367_v184 = _compr_9;
          if (_dafny.Seq.contains(_365_v183, _367_v184)) {
            _coll9.add(_367_v184);
          }
        }
        return _coll9;
      }()).length)));
      _nw52[(new BigNumber(14)).toNumber()] = new BigNumber((_83_v7).length);
      _nw52[(new BigNumber(15)).toNumber()] = new BigNumber((((_78_v1) ? (_83_v7) : (_83_v7))).length);
      _nw52[(new BigNumber(16)).toNumber()] = _82_v6;
      _nw52[(new BigNumber(17)).toNumber()] = _82_v6;
      _366_v185 = _nw52;
      let _rhs36 = _82_v6;
      let _rhs37 = (_277_v110).f24;
      let _rhs38 = (_82_v6).minus((((_364_v182).contains(_78_v1)) ? ((_364_v182).get(_78_v1)) : (_82_v6)));
      let _rhs39 = _366_v185;
      let _lhs22 = _93_globalState;
      _82_v6 = _rhs36;
      _78_v1 = _rhs37;
      _lhs22.f5 = _rhs38;
      _366_v185 = _rhs39;
      let _368_i17;
      _368_i17 = _dafny.ZERO;
      L1: {
        while (_78_v1) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_368_i17)) {
              break L1;
            }
            _368_i17 = (_368_i17).plus(_dafny.ONE);
            (_93_globalState).f21 = !(((_277_v110).f24) || ((_277_v110).f24));
            (_93_globalState).f5 = _82_v6;
            (_93_globalState).f21 = (_277_v110).f24;
            (_93_globalState).f5 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_82_v6, _82_v6, new BigNumber(529), _82_v6, _82_v6)).length))).plus(_82_v6);
          }
        }
      }
      let _369_v186;
      let _nw53 = new _module.C0();
      _nw53.__ctor(_82_v6, _78_v1, (_277_v110).f25);
      _369_v186 = _nw53;
      let _370_v187;
      _370_v187 = _dafny.Seq.of(_369_v186);
      _82_v6 = (_module.__default.safeModuloInt(_82_v6, _82_v6)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_module.__default.fm7((_277_v110).f24, _82_v6, new BigNumber(25), _78_v1, _93_globalState), _module.__default.safeIndex(new BigNumber((_370_v187).length), new BigNumber((_module.__default.fm7((_277_v110).f24, _82_v6, new BigNumber(25), _78_v1, _93_globalState)).length)), _84_v8)).length)));
      if (_78_v1) {
        (_93_globalState).f5 = (new BigNumber(((_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]).length)).multipliedBy(new BigNumber((((_78_v1) ? (_dafny.MultiSet.fromElements((_83_v7)[_module.__default.safeIndex((_369_v186).f26, new BigNumber((_83_v7).length))], (_369_v186).f26, (_369_v186).f26)) : (_87_v11))).cardinality()));
        let _371_v188;
        _371_v188 = _dafny.Map.Empty.slice().updateUnsafe(_366_v185,_369_v186);
        _371_v188 = (_371_v188).Merge(_371_v188);
        let _372_v189;
        _372_v189 = _dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,_279_v112);
        let _373_v190;
        _373_v190 = _dafny.Set.fromElements((_369_v186).f26);
        let _rhs40 = new BigNumber(((_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]).length);
        let _rhs41 = (((_372_v189).contains((_79_v2).IsProperSubsetOf(_79_v2))) ? ((_372_v189).get((_79_v2).IsProperSubsetOf(_79_v2))) : (_279_v112));
        let _rhs42 = _module.__default.fm6((_277_v110).f24, (_277_v110).f24, !((_277_v110).f24), _93_globalState);
        let _rhs43 = ((_369_v186).f26).minus((new BigNumber((_373_v190).length)).minus(_82_v6));
        let _lhs23 = _93_globalState;
        let _lhs24 = _93_globalState;
        let _lhs25 = _93_globalState;
        _lhs23.f5 = _rhs40;
        _279_v112 = _rhs41;
        _lhs24.f5 = _rhs42;
        _lhs25.f5 = _rhs43;
        (_93_globalState).f5 = (((((_277_v110).f24) ? ((_277_v110).f24) : ((_277_v110).f24))) ? (_82_v6) : ((_369_v186).f26));
        let _374_v191;
        _374_v191 = _dafny.Seq.of(_277_v110, _277_v110, _277_v110, _277_v110, _277_v110);
        let _375_v192;
        _375_v192 = _dafny.Set.fromElements(_79_v2);
        let _376_v193;
        let _nw54 = new _module.C0();
        _nw54.__ctor(new BigNumber((_257_v106).length), _78_v1, _dafny.Seq.UnicodeFromString("bacmjqi"));
        _376_v193 = _nw54;
        let _377_v194;
        _377_v194 = _dafny.Seq.of(_376_v193);
        let _378_v195;
        _378_v195 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_375_v192).length)),(_377_v194)[_module.__default.safeIndex((_369_v186).f26, new BigNumber((_377_v194).length))]);
        let _rhs44 = _module.__default.safeDivisionInt((_369_v186).f26, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_277_v110), _374_v191)).length));
        let _rhs45 = _module.__default.safeDivisionInt(((_369_v186).f26).minus((_369_v186).f26), new BigNumber(((_378_v195).Merge(_dafny.Map.Empty.slice().updateUnsafe(_82_v6,_376_v193))).length));
        let _rhs46 = (_277_v110).f24;
        let _lhs26 = _93_globalState;
        let _lhs27 = _93_globalState;
        let _lhs28 = _93_globalState;
        _lhs26.f5 = _rhs44;
        _lhs27.f5 = _rhs45;
        _lhs28.f21 = _rhs46;
      } else {
        let _379_v196;
        _379_v196 = _dafny.Map.Empty.slice().updateUnsafe(_77_v0,_78_v1);
        (_93_globalState).f20 = (((_379_v196).contains(_dafny.Seq.UnicodeFromString("nxd"))) ? ((_379_v196).get(_dafny.Seq.UnicodeFromString("nxd"))) : (_78_v1));
        let _380_v197;
        _380_v197 = _module.D3.create_DC11(_83_v7);
        let _381_v198;
        _381_v198 = _module.D3.create_DC11((_380_v197).dtor_cf19);
        let _source4 = _380_v197;
        if (_source4.is_DC12) {
          let _382___mcc_h29 = (_source4).cf20;
          let _383___mcc_h30 = (_source4).cf21;
          let _384_cf21 = _383___mcc_h30;
          let _385_cf20 = _382___mcc_h29;
          (_93_globalState).f5 = _82_v6;
          _82_v6 = new BigNumber((_92_v14).length);
          _80_v3 = ((_78_v1) ? (_80_v3) : (_80_v3));
          _366_v185 = _366_v185;
        } else if (_source4.is_DC13) {
          let _386___mcc_h31 = (_source4).cf22;
          let _387___mcc_h32 = (_source4).cf23;
          let _388_cf23 = _387___mcc_h32;
          let _389_cf22 = _386___mcc_h31;
          let _390_v199;
          let _nw55 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
          _390_v199 = _nw55;
          let _391_v200;
          _391_v200 = _module.D4.create_DC15(_257_v106);
          let _index26 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_390_v199).length));
          (_390_v199)[_index26] = (_391_v200).dtor_cf25;
          let _index27 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_390_v199).length));
          (_390_v199)[_index27] = _257_v106;
          let _392_v201;
          let _393_v202;
          let _394_v203;
          let _395_v204;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector13 = _module.__default.m0(_93_globalState);
          _out52 = _outcollector13[0];
          _out53 = _outcollector13[1];
          _out54 = _outcollector13[2];
          _out55 = _outcollector13[3];
          _392_v201 = _out52;
          _393_v202 = _out53;
          _394_v203 = _out54;
          _395_v204 = _out55;
          (_93_globalState).f5 = new BigNumber(((_277_v110).f25).length);
          let _index28 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_80_v3).length));
          (_80_v3)[_index28] = (_277_v110).f24;
          let _396_v205;
          _396_v205 = _dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,_84_v8);
          let _index29 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_80_v3).length));
          (_80_v3)[_index29] = (_389_cf22).fm1(_396_v205, _dafny.Map.Empty.slice().updateUnsafe(_394_v203,(_277_v110).f24), _dafny.Seq.Concat(_83_v7, _83_v7), _93_globalState);
        } else if (_source4.is_DC11) {
          let _397___mcc_h33 = (_source4).cf19;
          let _398_cf19 = _397___mcc_h33;
          let _399_v206;
          _399_v206 = _dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,(_369_v186).f26);
          (_93_globalState).f5 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,(_369_v186).f26)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,_82_v6)).Merge(_399_v206))).length);
          let _400_v207;
          let _401_v208;
          let _402_v209;
          let _403_v210;
          let _out56;
          let _out57;
          let _out58;
          let _out59;
          let _outcollector14 = _module.__default.m0(_93_globalState);
          _out56 = _outcollector14[0];
          _out57 = _outcollector14[1];
          _out58 = _outcollector14[2];
          _out59 = _outcollector14[3];
          _400_v207 = _out56;
          _401_v208 = _out57;
          _402_v209 = _out58;
          _403_v210 = _out59;
          let _404_v211;
          let _405_v212;
          let _406_v213;
          let _407_v214;
          let _out60;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector15 = _module.__default.m0(_93_globalState);
          _out60 = _outcollector15[0];
          _out61 = _outcollector15[1];
          _out62 = _outcollector15[2];
          _out63 = _outcollector15[3];
          _404_v211 = _out60;
          _405_v212 = _out61;
          _406_v213 = _out62;
          _407_v214 = _out63;
          let _408_v215;
          _408_v215 = _dafny.Seq.of(_83_v7);
          let _409_v216;
          _409_v216 = _dafny.Map.Empty.slice().updateUnsafe((_277_v110).f24,_84_v8);
          let _410_v217;
          _410_v217 = _dafny.Map.Empty.slice().updateUnsafe(_82_v6,_83_v7);
          (_93_globalState).f5 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_408_v215)[_module.__default.safeIndex(new BigNumber(929), new BigNumber((_408_v215).length))],(_369_v186).fm1((_409_v216).update(_403_v210, _84_v8), _dafny.Map.Empty.slice().updateUnsafe((_369_v186).f26,_403_v210), _dafny.Seq.update((((_410_v217).contains(new BigNumber((_79_v2).length))) ? ((_410_v217).get(new BigNumber((_79_v2).length))) : (_398_cf19)), _module.__default.safeIndex((_369_v186).f26, new BigNumber(((((_410_v217).contains(new BigNumber((_79_v2).length))) ? ((_410_v217).get(new BigNumber((_79_v2).length))) : (_398_cf19))).length)), _module.__default.fm6(_78_v1, !(_403_v210), (_277_v110).f24, _93_globalState)), _93_globalState))).length);
        } else {
          let _411___mcc_h34 = (_source4).cf24;
          let _412_cf24 = _411___mcc_h34;
          let _413_v218;
          _413_v218 = _dafny.Map.Empty.slice().updateUnsafe(false,_84_v8);
          let _414_v219;
          _414_v219 = _dafny.Map.Empty.slice().updateUnsafe(_82_v6,(_277_v110).f24);
          let _415_v220;
          let _nw56 = new _module.C0();
          _nw56.__ctor(_82_v6, (_277_v110).fm1(_413_v218, _414_v219, _dafny.Seq.update(_83_v7, _module.__default.safeIndex(_82_v6, new BigNumber((_83_v7).length)), (_369_v186).f26), _93_globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_416_v8) => function (_417_i18) {
            return _416_v8;
          })(_84_v8)));
          _415_v220 = _nw56;
          let _418_v221;
          _418_v221 = _module.D4.create_DC16((_369_v186).f26, (_257_v106)[_module.__default.safeIndex(_82_v6, new BigNumber((_257_v106).length))], _415_v220);
          let _419_v222;
          _419_v222 = _dafny.Map.Empty.slice().updateUnsafe(_418_v221,(new BigNumber((_279_v112).length)).plus((_dafny.ZERO).minus(new BigNumber(((_277_v110).f25).length))));
          _419_v222 = (_419_v222).update(_418_v221, new BigNumber((_257_v106).length));
          let _nw57 = new _module.C0();
          _nw57.__ctor(((_369_v186).f26).minus((_369_v186).fm3((_369_v186).f26, _dafny.Seq.of(_dafny.Seq.of(new BigNumber((function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of (_83_v7).Elements) {
              let _420_v223 = _compr_10;
              if (_dafny.Seq.contains(_83_v7, _420_v223)) {
                _coll10.push([_module.__default.safeDivisionInt(_420_v223, (_369_v186).f26),(_277_v110).f24]);
              }
            }
            return _coll10;
          }()).length))), new BigNumber(638), (_277_v110).f25, _93_globalState)), (((_277_v110).f24) ? (_78_v1) : (false)), (_277_v110).f25);
          _369_v186 = _nw57;
          let _421_v224;
          _421_v224 = _dafny.Set.fromElements(_380_v197);
          (_93_globalState).f5 = (new BigNumber((_279_v112).length)).multipliedBy(new BigNumber((_421_v224).length));
          let _422_v225;
          _422_v225 = _dafny.Map.Empty.slice().updateUnsafe(_258_v107,(_277_v110).f24);
          let _423_v226;
          _423_v226 = _module.D5.create_DC20(_422_v225);
          let _rhs47 = ((_423_v226).dtor_cf32).contains(_258_v107);
          let _rhs48 = (_369_v186).f26;
          let _lhs29 = _93_globalState;
          let _lhs30 = _93_globalState;
          _lhs29.f20 = _rhs47;
          _lhs30.f5 = _rhs48;
        }
        let _424_v227;
        _424_v227 = _dafny.Map.Empty.slice().updateUnsafe(_80_v3,(_369_v186).f26);
        let _425_v228;
        _425_v228 = _dafny.Map.Empty.slice().updateUnsafe((((_277_v110).f24) ? (_82_v6) : (new BigNumber((_424_v227).length))),!(_82_v6).isEqualTo(new BigNumber((_92_v14).length)));
        _425_v228 = (_425_v228).update(new BigNumber(488), _78_v1);
        let _index30 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_80_v3).length));
        (_80_v3)[_index30] = (_277_v110).f24;
        let _426_v229;
        let _nw58 = new _module.C0();
        _nw58.__ctor((_369_v186).f26, _78_v1, (_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]);
        _426_v229 = _nw58;
        let _427_v230;
        _427_v230 = _dafny.Seq.of(_426_v229, _426_v229, _426_v229);
        let _index31 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_80_v3).length));
        let _rhs49 = (_369_v186).f26;
        let _rhs50 = (new BigNumber(-804)).isLessThanOrEqualTo((_369_v186).f26);
        let _rhs51 = !_dafny.areEqual(_83_v7, _dafny.Seq.update(_83_v7, _module.__default.safeIndex(new BigNumber((_427_v230).length), new BigNumber((_83_v7).length)), _82_v6));
        let _rhs52 = (_257_v106)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((_257_v106).length))];
        let _lhs31 = _80_v3;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_80_v3).length));
        _82_v6 = _rhs49;
        _78_v1 = _rhs50;
        _lhs31[_lhs32] = _rhs51;
        _78_v1 = _rhs52;
        (_93_globalState).f5 = (_369_v186).f26;
      }
      let _428_v231;
      let _nw59 = Array((new BigNumber(2)).toNumber()).fill([]);
      _428_v231 = _nw59;
      let _429_v232;
      _429_v232 = _dafny.Map.Empty.slice().updateUnsafe(_428_v231,_dafny.MultiSet.FromArray(_dafny.Seq.of(false, (_277_v110).f24)));
      _429_v232 = (_429_v232).update(_428_v231, _364_v182);
      let _430_v233;
      let _nw60 = new _module.C0();
      _nw60.__ctor(_82_v6, (_277_v110).f24, (_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))]);
      _430_v233 = _nw60;
      let _index32 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_80_v3).length));
      (_80_v3)[_index32] = (_277_v110).f24;
      let _431_v234;
      _431_v234 = _dafny.Set.fromElements(_80_v3, _80_v3, _80_v3);
      let _index33 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_80_v3).length));
      let _index34 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length));
      let _rhs53 = (_431_v234).IsDisjointFrom(_431_v234);
      let _rhs54 = _77_v0;
      let _rhs55 = _dafny.Seq.Concat((_296_v127)[_module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length))], _77_v0);
      let _lhs33 = _80_v3;
      let _lhs34 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_80_v3).length));
      let _lhs35 = _296_v127;
      let _lhs36 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_296_v127).length));
      let _lhs37 = _93_globalState;
      _lhs33[_lhs34] = _rhs53;
      _lhs35[_lhs36] = _rhs54;
      _lhs37.f19 = _rhs55;
      process.stdout.write((_77_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_78_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_79_v2).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v3)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_81_v4).dtor_cf0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_82_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_83_v7, _dafny.Seq.of(new BigNumber(462), new BigNumber(462), new BigNumber(462)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_84_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v9)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_86_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_v11).equals(_dafny.MultiSet.fromElements(new BigNumber(462)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_88_v12, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(462))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(26)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(27)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_89_v13)[new BigNumber(28)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_92_v14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_93_globalState).f0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState.f1).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f2)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_93_globalState.f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_93_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(213444),new _dafny.CodePoint('j'.codePointAt(0))).updateUnsafe(new BigNumber(668),new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_93_globalState.f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState.f14).equals(_dafny.MultiSet.fromElements(new BigNumber(462)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(26)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(27)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_93_globalState).f15)[new BigNumber(28)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_93_globalState).f18)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_93_globalState.f19).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_93_globalState.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_93_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v105)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_257_v106, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v107).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(49),new BigNumber(462)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v108).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v108).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v110).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_277_v110).f25).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_278_v111).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_278_v111).dtor_cf11).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_278_v111).dtor_cf11).f25).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_278_v111).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_278_v111).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_283_v115).equals(_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v116).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_285_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_296_v127)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_296_v127)[new BigNumber(15)], _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_363_v181).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(213443)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v182).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_365_v183, _dafny.Seq.of(_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v185)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_368_i17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_369_v186).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_370_v187).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_429_v232).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_430_v233).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_431_v234).length)));
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
    static create_DC2(cf4, cf5) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D0(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, _dafny.ZERO);
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
    static create_DC5(cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf10, cf11, cf12, cf13) {
      let $dt = new D1(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC7(cf14) {
      let $dt = new D1(3);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11 && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf7 === other.cf7;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false, _dafny.ZERO);
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
    static create_DC9(cf16, cf17) {
      let $dt = new D2(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D2(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D2(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC9(false, _dafny.ZERO);
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
    static create_DC12(cf20, cf21) {
      let $dt = new D3(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC13(cf22, cf23) {
      let $dt = new D3(1);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D3(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC14(cf24) {
      let $dt = new D3(3);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC13" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC12(false, _dafny.ZERO);
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
    static create_DC16(cf26, cf27, cf28) {
      let $dt = new D4(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC17(cf29, cf30) {
      let $dt = new D4(1);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18() {
      let $dt = new D4(2);
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D4(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC19(cf31) {
      let $dt = new D4(4);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get is_DC19() { return this.$tag === 4; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC17" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC18";
      } else if (this.$tag === 3) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27 && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC16(_dafny.ZERO, false, null);
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
    static create_DC21(cf33) {
      let $dt = new D5(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC20(cf32) {
      let $dt = new D5(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC22(cf34) {
      let $dt = new D5(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC20" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC22" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC21(_dafny.Map.Empty);
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
    static create_DC23(cf35) {
      let $dt = new D6(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35;
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

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.Set.Empty;
      this.f4 = _dafny.Seq.UnicodeFromString("");
      this.f5 = _dafny.ZERO;
      this.f8 = _dafny.Seq.of();
      this.f14 = _dafny.MultiSet.Empty;
      this.f19 = _dafny.Seq.UnicodeFromString("");
      this.f20 = false;
      this.f21 = false;
      this._f0 = _dafny.Seq.UnicodeFromString("");
      this._f2 = [];
      this._f3 = false;
      this._f6 = _dafny.Map.Empty;
      this._f7 = false;
      this._f9 = false;
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = false;
      this._f13 = false;
      this._f15 = [];
      this._f16 = _dafny.ZERO;
      this._f17 = false;
      this._f18 = [];
      this._f22 = _dafny.ZERO;
      this._f23 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
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
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
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
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f24 = false;
      this._f25 = _dafny.Seq.UnicodeFromString("");
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f26, f24, f25) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-128),(_this).f24)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f24));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f24) && ((_this).f24);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return _module.D0.create_DC1((_this).f24, (_this).f24, (_this).f26);
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f26;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
