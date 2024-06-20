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
    static fm4(p0, globalState) {
      return (_module.D7.create_DC17(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(691),new BigNumber(995))).length), !(!(false)), true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ruhwlvnnb")).length)))).dtor_cf27;
    };
    static fm6(p0, globalState) {
      return !((_dafny.Set.fromElements(_dafny.Set.fromElements(!(true), false, true), _dafny.Set.fromElements(false), _dafny.Set.fromElements(!(true)))).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.Set.fromElements(true, true), _dafny.Set.fromElements(true))));
    };
    static fm8(p0, globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1(!(true), ((false) ? (new _dafny.CodePoint('y'.codePointAt(0))) : (new _dafny.CodePoint('j'.codePointAt(0)))));
    };
    static fm12(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(false, new _dafny.CodePoint('x'.codePointAt(0))),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(false, new _dafny.CodePoint('v'.codePointAt(0))),true));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      let _source0 = ((!(!(false))) ? (_module.D23.create_DC59(_dafny.Seq.UnicodeFromString("bhpqe"), new BigNumber(300), new BigNumber(162))) : (_module.D23.create_DC59(_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), function (_0_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
}), new BigNumber(-580), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(583))).cardinality())))));
      if (_source0.is_DC59) {
        let _1___mcc_h0 = (_source0).cf93;
        let _2___mcc_h1 = (_source0).cf94;
        let _3___mcc_h2 = (_source0).cf95;
        let _4_cf95 = _3___mcc_h2;
        let _5_cf94 = _2___mcc_h1;
        let _6_cf93 = _1___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_6_cf93,_5_cf94)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("f"),_4_cf95));
      } else if (_source0.is_DC60) {
        let _7___mcc_h3 = (_source0).cf96;
        let _8___mcc_h4 = (_source0).cf97;
        let _9_cf97 = _8___mcc_h4;
        let _10_cf96 = _7___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bk"),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_9_cf97,false)).length))).Merge(function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), function (_11_i1) {
            return _dafny.Seq.UnicodeFromString("qvbfh");
          })).Elements) {
            let _12_v0 = _compr_0;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), function (_11_i1) {
              return _dafny.Seq.UnicodeFromString("qvbfh");
            }), _12_v0)) {
              _coll0.push([_12_v0,new BigNumber(447)]);
            }
          }
          return _coll0;
        }());
      } else if (_source0.is_DC61) {
        let _13___mcc_h5 = (_source0).cf98;
        let _14_cf98 = _13___mcc_h5;
        return function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("aubgkgk"))).Elements) {
            let _15_v1 = _compr_1;
            if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("aubgkgk"))).contains(_15_v1)) {
              _coll1.push([_15_v1,_14_cf98]);
            }
          }
          return _coll1;
        }();
      } else {
        let _16___mcc_h6 = (_source0).cf92;
        let _17_cf92 = _16___mcc_h6;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(694)), function (_18_i2) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }),new BigNumber(-715));
      }
    };
    static fm14(globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eb"), _dafny.Seq.UnicodeFromString("mehe")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-217)), function (_19_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("lypkvvier")));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kgftfxi")).length),false));
    };
    static fm18(p0, p1, p2, globalState) {
      return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), function (_20_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length);
      }), _dafny.Seq.of(new BigNumber(485), new BigNumber(908), new BigNumber((_dafny.Set.fromElements(new BigNumber(14))).length), new BigNumber(375))));
    };
    static fm19(globalState) {
      return _dafny.Seq.of(new BigNumber(705));
    };
    static fm20(p0, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false, false, false));
    };
    static fm21(p0, p1, globalState) {
      return _dafny.areEqual(_dafny.Seq.UnicodeFromString("cjryf"), _dafny.Seq.UnicodeFromString("ri"));
    };
    static fm22(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(true)))).cardinality()),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jmue")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(308))).length),new BigNumber(-232))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(280),new BigNumber(540)));
    };
    static fm23(p0, p1, globalState) {
      let _source1 = _module.D10.create_DC23();
      if (_source1.is_DC23) {
        if (false) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }
      } else {
        let _21___mcc_h0 = (_source1).cf36;
        let _22_cf36 = _21___mcc_h0;
        return new _dafny.CodePoint('v'.codePointAt(0));
      }
    };
    static fm24(globalState) {
      return _dafny.MultiSet.FromArray(((((false) ? (true) : (!(!(false))))) ? (((!(true)) ? (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ptxr")).length))) : (_dafny.Seq.of(new BigNumber(622))))) : (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), function (_23_i0) {
        return new BigNumber(-587);
      }), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))))));
    };
    static fm25(globalState) {
      let _source2 = ((false) ? (_module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(349),false)).length),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(850))).cardinality())))) : (_module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hflpwwb"))).length), new BigNumber(749), new BigNumber(-493))).cardinality()),new BigNumber(856)))));
      if (_source2.is_DC25) {
        let _24___mcc_h0 = (_source2).cf38;
        let _25_cf38 = _24___mcc_h0;
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-330),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),!(true))).length));
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pneydj")).length),new BigNumber(553));
        }
      } else {
        let _26___mcc_h1 = (_source2).cf37;
        let _27_cf37 = _26___mcc_h1;
        return _27_cf37;
      }
    };
    static fm26(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!((!(!(true))) || (false)),true);
    };
    static fm27(p0, p1, p2, globalState) {
      return function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).Elements) {
          let _28_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0))), _28_v0)) {
            _coll2.add(_28_v0);
          }
        }
        return _coll2;
      }();
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(false)).equals(_dafny.Set.fromElements(true)),(_dafny.ZERO).minus((new BigNumber(-977)).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(168), new BigNumber((_dafny.Seq.UnicodeFromString("ljpjdt")).length), new BigNumber((_dafny.Seq.UnicodeFromString("iwdyqpvqq")).length), new BigNumber(-599))).length))));
    };
    static fm29(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-939)), function (_29_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-25))).length));
      }), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(153), (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(376)), function (_30_i1) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }))).Elements) {
          let _31_v0 = _compr_3;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(376)), function (_30_i1) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }))).contains(_31_v0)) {
            _coll3.push([_31_v0,new BigNumber(640)]);
          }
        }
        return _coll3;
      }()).length)))));
    };
    static fm30(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_32_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("h"));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return new BigNumber((function (_source3) {
        if (_source3.is_DC25) {
          let _33___mcc_h0 = (_source3).cf38;
          let _34_cf38 = _33___mcc_h0;
          return (_34_cf38).update(new BigNumber(-436), false);
        } else {
          let _35___mcc_h1 = (_source3).cf37;
          let _36_cf37 = _35___mcc_h1;
          return function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.Seq.of(new BigNumber(-199))).Elements) {
              let _37_v0 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-199)), _37_v0)) {
                _coll4.push([(_37_v0).plus(new BigNumber((_dafny.MultiSet.fromElements((_module.D19.create_DC45(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Set.fromElements(true)).length), true, false, _dafny.Seq.of(true, false))).dtor_cf71)).cardinality())),false]);
              }
            }
            return _coll4;
          }();
        }
      }(_module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(943),new BigNumber(383))))).length);
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(!(false), true, false)).length), new BigNumber((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-555)), function (_38_i0) {
        return new BigNumber(-693);
      }), _dafny.Seq.of(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(634), new BigNumber(314))) {
          let _39_v0 = _compr_5;
          if (((new BigNumber(634)).isLessThanOrEqualTo(_39_v0)) && ((_39_v0).isLessThan(new BigNumber(314)))) {
            _coll5.push([_module.__default.safeDivisionInt(_39_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("crmga")).length)))),false]);
          }
        }
        return _coll5;
      }()).length), new BigNumber((_dafny.Seq.of(true, true)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(572)), function (_40_i1) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(130)), function (_41_i2) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length);
      }))).length))), new BigNumber(-951), (new BigNumber(739)).multipliedBy(new BigNumber(313)));
    };
    static fm33(p0, p1, globalState) {
      return (_dafny.Set.fromElements(true, (_module.D22.create_DC56((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of (_dafny.Seq.of(new BigNumber(322))).Elements) {
    let _42_v0 = _compr_6;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(322)), _42_v0)) {
      _coll6.add(_module.__default.safeDivisionInt(_42_v0, new BigNumber(244)));
    }
  }
  return _coll6;
}()).length)), false)).dtor_cf90, true)).Union(_dafny.Set.fromElements(true));
    };
    static fm34(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(!(false), false));
    };
    static fm35(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("yg")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("strrna"))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(_dafny.ONE), function (_43_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ifw"))));
    };
    static fm36(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(500)).plus(new BigNumber((_dafny.Seq.of(!(false))).length)),_dafny.Set.fromElements(true, true, !(true)));
    };
    static fm37(p0, globalState) {
      if (((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),false),true)).length))).length), new BigNumber(607), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), function (_44_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_45_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }))).cardinality())))).length))))).isEqualTo(new BigNumber(-718))) {
        return _module.D16.create_DC36(_dafny.Set.fromElements(new BigNumber(-887)), !(true), true, false);
      } else {
        return _module.D16.create_DC36(_dafny.Set.fromElements(new BigNumber(141)), false, false, !(false));
      }
    };
    static fm38(p0, p1, globalState) {
      return _module.D12.create_DC28(_module.D12.create_DC28(_module.D12.create_DC28(_module.D12.create_DC28(_module.D12.create_DC27(new BigNumber(-792))))));
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC10((_module.D3.create_DC10(_dafny.MultiSet.fromElements(true, false), _module.D1.create_DC4(new BigNumber(-992)))).dtor_cf16, _module.D1.create_DC4(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), function (_46_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length),new BigNumber(12))).length)));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('i'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0))));
    };
    static fm41(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("emqvugbr"), _dafny.Seq.UnicodeFromString("kwm")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("byfhiw"), _dafny.Seq.UnicodeFromString("thtg")));
    };
    static fm42(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D4.create_DC12(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), function (_47_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("bpkjcwei")).length);
})).length), new BigNumber((_dafny.Seq.UnicodeFromString("kkajua")).length), new BigNumber((_dafny.Seq.of(true)).length))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length))).length), false))).Difference(_dafny.MultiSet.fromElements(_module.D4.create_DC12(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber(-41), (_module.D22.create_DC56(new BigNumber(560), false)).dtor_cf90)))).Difference(_dafny.MultiSet.FromArray((_module.D27.create_DC68(_dafny.Seq.of(_module.D4.create_DC12((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pw")).length)), new BigNumber(-595), true), _module.D4.create_DC12(new BigNumber((_dafny.Seq.UnicodeFromString("nmeenbgie")).length), new BigNumber(885), true)))).dtor_cf110));
    };
    static fm43(p0, p1, p2, globalState) {
      return _module.D4.create_DC13(_module.D4.create_DC11(_dafny.Seq.of(false, true)));
    };
    static fm44(p0, p1, globalState) {
      return _dafny.Set.fromElements((_dafny.Set.fromElements(_module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("a")).length), new BigNumber(55)),new _dafny.CodePoint('b'.codePointAt(0)))), _module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(456), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length))),new _dafny.CodePoint('x'.codePointAt(0)))), _module.D17.create_DC37(function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(374)))).Elements) {
    let _48_v0 = _compr_7;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(374))), _48_v0)) {
      _coll7.push([_48_v0,new _dafny.CodePoint('o'.codePointAt(0))]);
    }
  }
  return _coll7;
}()))).Difference(_dafny.Set.fromElements(_module.D17.create_DC37(function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(846), new BigNumber(-631))).cardinality()))).length), new BigNumber(542)),new BigNumber((_dafny.Seq.of(!(!(false)), false, false)).length))).Keys.Elements) {
    let _49_v1 = _compr_8;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(846), new BigNumber(-631))).cardinality()))).length), new BigNumber(542)),new BigNumber((_dafny.Seq.of(!(!(false)), false, false)).length))).contains(_49_v1)) {
      _coll8.push([_49_v1,new _dafny.CodePoint('f'.codePointAt(0))]);
    }
  }
  return _coll8;
}()), _module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), function (_50_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("mcdjuq")).length);
}),new _dafny.CodePoint('q'.codePointAt(0)))))), _dafny.Set.fromElements(_module.D17.create_DC37(function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_51_i1) {
    return _dafny.Seq.of(new BigNumber(-617));
  }))).Elements) {
    let _52_v2 = _compr_9;
    if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_51_i1) {
      return _dafny.Seq.of(new BigNumber(-617));
    }))).contains(_52_v2)) {
      _coll9.push([_52_v2,new _dafny.CodePoint('i'.codePointAt(0))]);
    }
  }
  return _coll9;
}())));
    };
    static fm45(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new _dafny.CodePoint('u'.codePointAt(0)));
    };
    static fm46(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(499)).isLessThan((_dafny.ZERO).minus((_module.D1.create_DC5(false, false, new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), (_dafny.ZERO).minus(new BigNumber(-179)))).dtor_cf11)),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), function (_53_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("dwrbjb")));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ji"), _dafny.Seq.UnicodeFromString("yns"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), function (_54_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_55_i1) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        })))).Elements) {
          let _56_v0 = _compr_10;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ji"), _dafny.Seq.UnicodeFromString("yns"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), function (_54_i0) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_55_i1) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })))).contains(_56_v0)) {
            _coll10.push([_56_v0,false]);
          }
        }
        return _coll10;
      }()).length)),new BigNumber(630));
    };
    static fm48(p0, globalState) {
      if ((_dafny.Set.fromElements(!(true))).IsProperSubsetOf(_dafny.Set.fromElements(true))) {
        return _module.D1.create_DC5(true, true, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-114)), function (_57_i0) {
  return new BigNumber((_dafny.Set.fromElements(false)).length);
})).length), new BigNumber(941), new BigNumber(-254), new BigNumber(841))).length), new BigNumber(117));
      } else {
        return _module.D1.create_DC5(!(true), true, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new _dafny.CodePoint('x'.codePointAt(0)))).length),false)).length))).cardinality()), new BigNumber(-864))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false)).length));
      }
    };
    static fm49(p0, globalState) {
      return _module.D23.create_DC59(_dafny.Seq.Create(_module.__default.abs(new BigNumber(507)), function (_58_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}), new BigNumber(8), (new BigNumber((_dafny.Set.fromElements(true, true)).length)).multipliedBy(new BigNumber(487)));
    };
    static fm50(p0, globalState) {
      return function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-751))).length)),true)).Keys.Elements) {
          let _59_v0 = _compr_11;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-751))).length)),true)).contains(_59_v0)) {
            _coll11.push([_59_v0,new _dafny.CodePoint('p'.codePointAt(0))]);
          }
        }
        return _coll11;
      }();
    };
    static fm51(globalState) {
      return _module.D0.create_DC2((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, true)), ((true) ? (new BigNumber((_dafny.Seq.of(false)).length)) : (new BigNumber((_dafny.Set.fromElements(true)).length))), (_dafny.ZERO).minus(((false) ? (new BigNumber((_dafny.Seq.UnicodeFromString("yj")).length)) : (new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())))));
    };
    static m12(p0, p1, p2, p3, globalState) {
      let r0 = _module.D11.Default();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let r3 = undefined;
      let _60_v0;
      _60_v0 = _dafny.Seq.of(p2, !(!(p2)));
      let _hi0 = _module.__default.fm31(p2, p0, !(p3), _60_v0, globalState);
      for (let _61_i0 = p0; _61_i0.isLessThan(_hi0); _61_i0 = _61_i0.plus(_dafny.ONE)) {
        r1 = (_61_i0).plus((_61_i0).minus(p0));
        let _62_v1;
        _62_v1 = _dafny.Seq.UnicodeFromString("wowdbx");
        let _63_v2;
        let _init0 = ((_64_p2) => function (_65_i1) {
          return _64_p2;
        })(p2);
        let _nw0 = Array((new BigNumber(16)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _63_v2 = _nw0;
        let _66_v3;
        _66_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_63_v2);
        let _67_v4;
        let _nw1 = new _module.C1();
        _nw1.__ctor(_dafny.Seq.UnicodeFromString("q"), p3);
        _67_v4 = _nw1;
        let _68_v5;
        _68_v5 = _module.D15.create_DC33(_63_v2, _67_v4);
        let _69_v6;
        _69_v6 = _dafny.Seq.of(_63_v2);
        let _70_v7;
        let _nw2 = Array((new BigNumber(18)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = (((_66_v3).contains(_61_i0)) ? ((_66_v3).get(_61_i0)) : (_63_v2));
        _nw2[(_dafny.ONE).toNumber()] = _63_v2;
        _nw2[(new BigNumber(2)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(3)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(4)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(5)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(6)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(7)).toNumber()] = (_68_v5).dtor_cf47;
        _nw2[(new BigNumber(8)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(9)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(10)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(11)).toNumber()] = ((p2) ? (_63_v2) : (_63_v2));
        _nw2[(new BigNumber(12)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(13)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(14)).toNumber()] = (_69_v6)[_module.__default.safeIndex(new BigNumber(-869), new BigNumber((_69_v6).length))];
        _nw2[(new BigNumber(15)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(16)).toNumber()] = _63_v2;
        _nw2[(new BigNumber(17)).toNumber()] = _63_v2;
        _70_v7 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_70_v7).length));
        (_70_v7)[_index0] = _63_v2;
        let _71_v8;
        let _nw3 = new _module.C3();
        _nw3.__ctor((_67_v4).f13, p2);
        _71_v8 = _nw3;
        let _72_v9;
        _72_v9 = _module.D22.create_DC55(_71_v8);
        let _index1 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_70_v7).length));
        let _rhs0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(753)), ((_73_v4) => function (_74_i2) {
          return (((_73_v4).f13) ? (new _dafny.CodePoint('m'.codePointAt(0))) : (new _dafny.CodePoint('o'.codePointAt(0))));
        })(_67_v4));
        let _rhs1 = _63_v2;
        let _rhs2 = p3;
        let _rhs3 = _63_v2;
        let _rhs4 = _72_v9;
        let _lhs0 = _70_v7;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_70_v7).length));
        let _lhs2 = globalState;
        _62_v1 = _rhs0;
        _lhs0[_lhs1] = _rhs1;
        _lhs2.f0 = _rhs2;
        _63_v2 = _rhs3;
        _72_v9 = _rhs4;
        let _75_v10;
        let _nw4 = new _module.C4();
        _nw4.__ctor(p1, _61_i0, false, p1);
        _75_v10 = _nw4;
        _75_v10 = _75_v10;
        let _76_v11;
        let _nw5 = new _module.C6();
        _nw5.__ctor(p1, p2, p1);
        _76_v11 = _nw5;
        let _77_v12;
        let _nw6 = new _module.C5();
        _nw6.__ctor(_dafny.Seq.UnicodeFromString("j"), (_75_v10).f9, _75_v10.f8, p3);
        _77_v12 = _nw6;
        let _78_v13;
        _78_v13 = _dafny.Map.Empty.slice().updateUnsafe(_61_i0,_77_v12);
        let _79_v14;
        _79_v14 = _dafny.Set.fromElements(_78_v13, _78_v13);
        let _80_v15;
        _80_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Map.Empty.slice().updateUnsafe(_77_v12.f7,_77_v12));
        let _81_v16;
        _81_v16 = _dafny.Map.Empty.slice().updateUnsafe(_61_i0,_77_v12.f7);
        let _82_v17;
        _82_v17 = _dafny.MultiSet.fromElements(_81_v16, _81_v16);
        let _83_v18;
        _83_v18 = _module.D2.create_DC8((_70_v7)[_module.__default.safeIndex(new BigNumber(44), new BigNumber((_70_v7).length))], (_67_v4).f12);
        let _84_v19;
        _84_v19 = new _dafny.CodePoint('i'.codePointAt(0));
        let _85_v20;
        let _nw7 = new _module.C0();
        _nw7.__ctor();
        _85_v20 = _nw7;
        let _86_v21;
        _86_v21 = _dafny.Seq.of(_85_v20);
        let _87_v22;
        _87_v22 = _dafny.MultiSet.fromElements(_77_v12.f7, new BigNumber(870));
        let _88_v23;
        _88_v23 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(341)), ((_89_v19) => function (_90_i3) {
          return _89_v19;
        })(_84_v19))).length), (_dafny.ZERO).minus(p0));
        let _91_v24;
        let _nw8 = Array((new BigNumber(20)).toNumber());
        _nw8[(_dafny.ZERO).toNumber()] = (p0).multipliedBy(_61_i0);
        _nw8[(_dafny.ONE).toNumber()] = _61_i0;
        _nw8[(new BigNumber(2)).toNumber()] = _61_i0;
        _nw8[(new BigNumber(3)).toNumber()] = (new BigNumber(24)).minus(new BigNumber((_dafny.Seq.of(_76_v11)).length));
        _nw8[(new BigNumber(4)).toNumber()] = (_75_v10).f9;
        _nw8[(new BigNumber(5)).toNumber()] = new BigNumber(((_79_v14).Intersect(_dafny.Set.fromElements((((_80_v15).contains(_module.__default.fm21(_61_i0, _82_v17, globalState))) ? ((_80_v15).get(_module.__default.fm21(_61_i0, _82_v17, globalState))) : (_78_v13))))).length);
        _nw8[(new BigNumber(6)).toNumber()] = p0;
        _nw8[(new BigNumber(7)).toNumber()] = _61_i0;
        _nw8[(new BigNumber(8)).toNumber()] = new BigNumber((_module.__default.fm30(_module.D1.create_DC5((_67_v4).f13, (_67_v4).f13, _61_i0, (_75_v10).f9), (_75_v10).fm5((_67_v4).f13, _dafny.Seq.update((_83_v18).dtor_cf14, _module.__default.safeIndex(_61_i0, new BigNumber(((_83_v18).dtor_cf14).length)), _84_v19), globalState), new BigNumber(((_77_v12).f6).length), globalState)).length);
        _nw8[(new BigNumber(9)).toNumber()] = (_77_v12.f7).multipliedBy(_77_v12.f7);
        _nw8[(new BigNumber(10)).toNumber()] = (new BigNumber((_60_v0).length)).multipliedBy(p0);
        _nw8[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_85_v20, _85_v20, _85_v20, _85_v20), _dafny.Seq.update(_86_v21, _module.__default.safeIndex(new BigNumber((_87_v22).cardinality()), new BigNumber((_86_v21).length)), _85_v20))).length);
        _nw8[(new BigNumber(12)).toNumber()] = (_75_v10).fm0(globalState);
        _nw8[(new BigNumber(13)).toNumber()] = (_75_v10).f9;
        _nw8[(new BigNumber(14)).toNumber()] = (_77_v12.f7).minus(p0);
        _nw8[(new BigNumber(15)).toNumber()] = _77_v12.f7;
        _nw8[(new BigNumber(16)).toNumber()] = new BigNumber(-490);
        _nw8[(new BigNumber(17)).toNumber()] = _61_i0;
        _nw8[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), new BigNumber((_88_v23).length));
        _nw8[(new BigNumber(19)).toNumber()] = ((_75_v10).fm5(p2, _62_v1, globalState)).multipliedBy(_61_i0);
        _91_v24 = _nw8;
        let _index2 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_91_v24).length));
        (_91_v24)[_index2] = _77_v12.f7;
        let _index3 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_91_v24).length));
        (_91_v24)[_index3] = p0;
      }
      let _92_v25;
      let _nw9 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _92_v25 = _nw9;
      let _93_v26;
      let _init1 = ((_94_p1) => function (_95_i4) {
        return _94_p1;
      })(p1);
      let _nw10 = Array((new BigNumber(2)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw10.length); _i0_1++) {
        _nw10[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _93_v26 = _nw10;
      let _index4 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length));
      (_93_v26)[_index4] = p1;
      let _96_v27;
      _96_v27 = _dafny.Seq.UnicodeFromString("fkcu");
      let _97_v28;
      let _nw11 = new _module.C1();
      _nw11.__ctor(_96_v27, p2);
      _97_v28 = _nw11;
      let _98_v29;
      _98_v29 = _module.D15.create_DC33(_93_v26, _97_v28);
      let _99_v30;
      _99_v30 = _dafny.Seq.of(_module.D15.create_DC33(_93_v26, _97_v28), _98_v29, _module.D15.create_DC33(_93_v26, _97_v28));
      let _100_v31;
      _100_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
      let _101_v32;
      _101_v32 = _dafny.Set.fromElements(_100_v31);
      let _102_v33;
      _102_v33 = _module.D8.create_DC18(_100_v31);
      let _103_v34;
      _103_v34 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,true));
      let _104_v35;
      _104_v35 = _dafny.Seq.of((_103_v34)[_module.__default.safeIndex(new BigNumber(((_97_v28).f12).length), new BigNumber((_103_v34).length))], _100_v31);
      let _index5 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length));
      let _rhs5 = !(_dafny.Seq.IsPrefixOf(_99_v30, _dafny.Seq.Concat(_99_v30, _99_v30)));
      let _rhs6 = p0;
      let _rhs7 = _92_v25;
      let _rhs8 = ((_101_v32).Union(_dafny.Set.fromElements(_100_v31, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_102_v33).dtor_cf30).length),(_97_v28).f13)))).IsDisjointFrom(function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of (_104_v35).Elements) {
          let _105_v36 = _compr_12;
          if (_dafny.Seq.contains(_104_v35, _105_v36)) {
            _coll12.add(_105_v36);
          }
        }
        return _coll12;
      }());
      let _lhs3 = globalState;
      let _lhs4 = _93_v26;
      let _lhs5 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length));
      _lhs3.f0 = _rhs5;
      r1 = _rhs6;
      _92_v25 = _rhs7;
      _lhs4[_lhs5] = _rhs8;
      let _106_v37;
      let _nw12 = new _module.C4();
      _nw12.__ctor(!(!(false)), p0, (_97_v28).f13, (_93_v26)[_module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length))]);
      _106_v37 = _nw12;
      let _107_v38;
      _107_v38 = _dafny.Map.Empty.slice().updateUnsafe(_106_v37,_106_v37.f8);
      _107_v38 = (_107_v38).update(_106_v37, _dafny.Seq.IsProperPrefixOf((_97_v28).f12, _96_v27));
      let _hi1 = (_dafny.ZERO).minus(p0);
      for (let _108_i5 = new BigNumber(838); _108_i5.isLessThan(_hi1); _108_i5 = _108_i5.plus(_dafny.ONE)) {
        let _index6 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length));
        (_92_v25)[_index6] = (_106_v37).f9;
        let _109_v39;
        _109_v39 = _module.D19.create_DC45(_108_i5, (_dafny.ZERO).minus(_108_i5), _module.__default.fm6((_106_v37).f9, globalState), p1, _60_v0);
        let _110_v40;
        let _nw13 = new _module.C6();
        _nw13.__ctor(!((_97_v28).f13), (_97_v28).f13, (_109_v39).dtor_cf72);
        _110_v40 = _nw13;
        let _111_v41;
        _111_v41 = _dafny.Set.fromElements(_110_v40, _110_v40, _110_v40);
        let _index7 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length));
        let _rhs9 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lg"), (_97_v28).f12)).length)).plus((((_93_v26)[_module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length))]) ? (new BigNumber((_111_v41).length)) : ((_106_v37).f9)));
        let _rhs10 = (_108_i5).isLessThanOrEqualTo(new BigNumber((_96_v27).length));
        let _rhs11 = p0;
        let _rhs12 = (_106_v37).f9;
        let _lhs6 = _106_v37;
        let _lhs7 = _92_v25;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length));
        r1 = _rhs9;
        _lhs6.f8 = _rhs10;
        _lhs7[_lhs8] = _rhs11;
        r1 = _rhs12;
        let _index8 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length));
        (_92_v25)[_index8] = (_dafny.ZERO).minus(p0);
        let _112_v42;
        _112_v42 = _dafny.Seq.of(p0, (_106_v37).f9, _108_i5, p0, p0);
        r1 = _module.__default.safeModuloInt(((_110_v40).fm1((_97_v28).f13, (_92_v25)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length))], true, _module.__default.fm29(_108_i5, globalState), globalState)).plus(new BigNumber(948)), (new BigNumber((_112_v42).length)).multipliedBy(new BigNumber(-952)));
        let _113_v43;
        _113_v43 = new _dafny.CodePoint('r'.codePointAt(0));
        let _114_v44;
        let _nw14 = Array((new BigNumber(15)).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_115_i6) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(197)), function (_116_i7) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }));
        _nw14[(_dafny.ONE).toNumber()] = (_97_v28).f12;
        _nw14[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_96_v27, _module.__default.safeIndex((_92_v25)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_92_v25).length))], new BigNumber((_96_v27).length)), _113_v43);
        _nw14[(new BigNumber(3)).toNumber()] = (_97_v28).f12;
        _nw14[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_96_v27, (_97_v28).f12);
        _nw14[(new BigNumber(5)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_97_v28).f12, (_97_v28).f12);
        _nw14[(new BigNumber(7)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(8)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(9)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(10)).toNumber()] = (_97_v28).f12;
        _nw14[(new BigNumber(11)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_96_v27, (_97_v28).f12);
        _nw14[(new BigNumber(13)).toNumber()] = _96_v27;
        _nw14[(new BigNumber(14)).toNumber()] = _96_v27;
        _114_v44 = _nw14;
        let _index9 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_114_v44).length));
        (_114_v44)[_index9] = (_97_v28).f12;
        let _index10 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_114_v44).length));
        (_114_v44)[_index10] = (((_module.__default.fm51(globalState)).dtor_cf3) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ugxsphir"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-287)), ((_117_v43) => function (_118_i8) {
          return _117_v43;
        })(_113_v43)))) : (_dafny.Seq.Concat(_96_v27, _96_v27)));
      }
      let _119_v45;
      _119_v45 = _dafny.MultiSet.fromElements((_97_v28).f13);
      let _index11 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length));
      (_92_v25)[_index11] = (((_119_v45).contains(_106_v37.f8)) ? ((_119_v45).get(_106_v37.f8)) : (p0));
      let _index12 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length));
      let _rhs13 = p0;
      let _rhs14 = ((_106_v37).f9).plus((_dafny.ZERO).minus((_106_v37).fm5(true, _96_v27, globalState)));
      let _rhs15 = ((p1) ? (_93_v26) : (_93_v26));
      let _lhs9 = _92_v25;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length));
      r1 = _rhs13;
      _lhs9[_lhs10] = _rhs14;
      _93_v26 = _rhs15;
      let _120_i9;
      _120_i9 = _dafny.ZERO;
      L0: {
        while ((_93_v26)[_module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_120_i9)) {
              break L0;
            }
            _120_i9 = (_120_i9).plus(_dafny.ONE);
            let _121_v46;
            let _init2 = ((_122_v0) => function (_123_i10) {
              return _module.D4.create_DC11(_122_v0);
            })(_60_v0);
            let _nw15 = Array((_dafny.ONE).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw15.length); _i0_2++) {
              _nw15[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _121_v46 = _nw15;
            let _124_v47;
            _124_v47 = _module.D4.create_DC11(_60_v0);
            let _index13 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_121_v46).length));
            (_121_v46)[_index13] = _124_v47;
            let _pat_let_tv0 = _106_v37;
            let _index14 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length));
            let _index15 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_121_v46).length));
            let _rhs16 = (_92_v25)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length))];
            let _rhs17 = function (_pat_let0_0) {
              return function (_125_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_126_dt__update_hcf18_h0) {
                    return _module.D4.create_DC11(_126_dt__update_hcf18_h0);
                  }(_pat_let1_0);
                }(_dafny.Seq.of(_pat_let_tv0.f8));
              }(_pat_let0_0);
            }(_124_v47);
            let _lhs11 = _92_v25;
            let _lhs12 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length));
            let _lhs13 = _121_v46;
            let _lhs14 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_121_v46).length));
            _lhs11[_lhs12] = _rhs16;
            _lhs13[_lhs14] = _rhs17;
            let _index16 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length));
            (_93_v26)[_index16] = p1;
            let _127_v48;
            _127_v48 = _dafny.Seq.of((((_119_v45).contains(p3)) ? ((_119_v45).get(p3)) : ((_106_v37).f9)), new BigNumber(593), (_106_v37).f9, p0);
            r1 = ((new BigNumber((_127_v48).length)).plus((_106_v37).f9)).multipliedBy((_dafny.ZERO).minus((_106_v37).f9));
            let _128_v49;
            let _nw16 = new _module.C1();
            _nw16.__ctor(_dafny.Seq.Concat((_97_v28).f12, _96_v27), (_97_v28).f13);
            _128_v49 = _nw16;
          }
        }
      }
      let _129_v50;
      _129_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_92_v25)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length))]);
      let _130_v51;
      _130_v51 = _dafny.Seq.of(_129_v50);
      let _131_v52;
      _131_v52 = _dafny.MultiSet.fromElements(_96_v27);
      r0 = _module.D11.create_DC24((_130_v51)[_module.__default.safeIndex(new BigNumber((_131_v52).cardinality()), new BigNumber((_130_v51).length))]);
      r1 = (_92_v25)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length))];
      let _132_v53;
      let _nw17 = new _module.C3();
      _nw17.__ctor(p2, !(!(_module.__default.fm21((_92_v25)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_92_v25).length))], _dafny.MultiSet.FromArray(_dafny.Seq.of(_129_v50, _129_v50, _129_v50)), globalState))));
      _132_v53 = _nw17;
      let _133_v54;
      _133_v54 = _dafny.Seq.of(_132_v53);
      r2 = _133_v54;
      let _nw18 = new _module.C2();
      _nw18.__ctor((_93_v26)[_module.__default.safeIndex(new BigNumber(969), new BigNumber((_93_v26).length))], _106_v37.f8, ((_dafny.ZERO).minus((_106_v37).f9)).isLessThanOrEqualTo((_dafny.ZERO).minus(p0)), p1);
      r3 = _nw18;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _134_globalState;
      let _nw19 = new _module.GlobalState();
      _nw19.__ctor(true, true, true);
      _134_globalState = _nw19;
      let _135_v0;
      _135_v0 = _dafny.Seq.of(true);
      let _136_v1;
      _136_v1 = false;
      let _137_i0;
      _137_i0 = _dafny.ZERO;
      L1: {
        while ((_dafny.Seq.IsPrefixOf(_135_v0, _dafny.Seq.of(_136_v1, _136_v1))) && (_136_v1)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_137_i0)) {
              break L1;
            }
            _137_i0 = (_137_i0).plus(_dafny.ONE);
            let _138_v2;
            let _nw20 = new _module.C6();
            _nw20.__ctor(_136_v1, _136_v1, _136_v1);
            _138_v2 = _nw20;
            let _139_v3;
            _139_v3 = _dafny.MultiSet.fromElements(_138_v2.f5, true);
            let _140_v4;
            _140_v4 = new BigNumber(194);
            let _141_v5;
            let _nw21 = new _module.C4();
            _nw21.__ctor(_module.__default.fm6(_140_v4, _134_globalState), _140_v4, _136_v1, _138_v2.f5);
            _141_v5 = _nw21;
            let _142_v6;
            _142_v6 = _dafny.Map.Empty.slice().updateUnsafe(_141_v5,_141_v5.f8);
            let _143_v7;
            _143_v7 = _dafny.Map.Empty.slice().updateUnsafe(!((((_142_v6).contains(_141_v5)) ? ((_142_v6).get(_141_v5)) : (_141_v5.f8))),_136_v1);
            let _144_v8;
            _144_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_143_v7).contains(!(_141_v5.f8))) ? ((_143_v7).get(!(_141_v5.f8))) : (_141_v5.f8)),_139_v3);
            let _145_v9;
            _145_v9 = _dafny.Seq.of(_module.__default.fm20(_138_v2.f5, _134_globalState), _139_v3, _139_v3, _139_v3, _139_v3);
            let _146_v10;
            let _nw22 = Array((new BigNumber(14)).toNumber());
            _nw22[(_dafny.ZERO).toNumber()] = (_139_v3).Intersect(_139_v3);
            _nw22[(_dafny.ONE).toNumber()] = _139_v3;
            _nw22[(new BigNumber(2)).toNumber()] = (((_144_v8).contains(_141_v5.f8)) ? ((_144_v8).get(_141_v5.f8)) : (_module.__default.fm20(_138_v2.f5, _134_globalState)));
            _nw22[(new BigNumber(3)).toNumber()] = _139_v3;
            _nw22[(new BigNumber(4)).toNumber()] = (_139_v3).Difference(_139_v3);
            _nw22[(new BigNumber(5)).toNumber()] = _module.__default.fm20(_138_v2.f5, _134_globalState);
            _nw22[(new BigNumber(6)).toNumber()] = _139_v3;
            _nw22[(new BigNumber(7)).toNumber()] = _139_v3;
            _nw22[(new BigNumber(8)).toNumber()] = _139_v3;
            _nw22[(new BigNumber(9)).toNumber()] = _139_v3;
            _nw22[(new BigNumber(10)).toNumber()] = (_145_v9)[_module.__default.safeIndex(new BigNumber(-596), new BigNumber((_145_v9).length))];
            _nw22[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_138_v2.f5);
            _nw22[(new BigNumber(12)).toNumber()] = (_139_v3).Union(_139_v3);
            _nw22[(new BigNumber(13)).toNumber()] = _139_v3;
            _146_v10 = _nw22;
            let _index17 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_146_v10).length));
            (_146_v10)[_index17] = _139_v3;
            let _index18 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_146_v10).length));
            (_146_v10)[_index18] = _module.__default.fm20(_138_v2.f5, _134_globalState);
            let _147_v11;
            let _nw23 = new _module.C3();
            _nw23.__ctor(!(_141_v5.f8), _141_v5.f8);
            _147_v11 = _nw23;
            let _nw24 = new _module.C3();
            _nw24.__ctor(_141_v5.f8, _141_v5.f8);
            _147_v11 = _nw24;
            _140_v4 = (new BigNumber(786)).minus((_dafny.ZERO).minus(_140_v4));
          }
        }
      }
      if (((_136_v1) ? ((_136_v1) === (_136_v1)) : (_136_v1))) {
        let _148_v12;
        _148_v12 = new BigNumber(958);
        let _149_v13;
        _149_v13 = _dafny.MultiSet.fromElements(_148_v12);
        let _150_v14;
        _150_v14 = _module.D23.create_DC61(new BigNumber((_149_v13).cardinality()));
        let _pat_let_tv1 = _148_v12;
        let _pat_let_tv2 = _148_v12;
        let _pat_let_tv3 = _136_v1;
        let _source4 = function (_pat_let2_0) {
          return function (_151_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_152_dt__update_hcf98_h0) {
                return _module.D23.create_DC61(_152_dt__update_hcf98_h0);
              }(_pat_let3_0);
            }(((_pat_let_tv3) ? (_pat_let_tv1) : (_pat_let_tv2)));
          }(_pat_let2_0);
        }(_150_v14);
        if (_source4.is_DC59) {
          let _153___mcc_h0 = (_source4).cf93;
          let _154___mcc_h1 = (_source4).cf94;
          let _155___mcc_h2 = (_source4).cf95;
          let _156_cf95 = _155___mcc_h2;
          let _157_cf94 = _154___mcc_h1;
          let _158_cf93 = _153___mcc_h0;
          _136_v1 = !(_148_v12).isEqualTo((_dafny.ZERO).minus((new BigNumber(882)).multipliedBy((_dafny.ZERO).minus(_156_cf95))));
          (_134_globalState).f0 = ((_136_v1) ? (_136_v1) : (_136_v1));
          _157_cf94 = _157_cf94;
          _158_cf93 = _dafny.Seq.Concat(_158_cf93, _158_cf93);
        } else if (_source4.is_DC60) {
          let _159___mcc_h3 = (_source4).cf96;
          let _160___mcc_h4 = (_source4).cf97;
          let _161_cf97 = _160___mcc_h4;
          let _162_cf96 = _159___mcc_h3;
          let _163_v15;
          _163_v15 = _dafny.Seq.UnicodeFromString("b");
          let _164_v17;
          _164_v17 = new _dafny.CodePoint('u'.codePointAt(0));
          let _165_v18;
          let _nw25 = Array((new BigNumber(13)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_148_v12);
          _nw25[(_dafny.ONE).toNumber()] = _148_v12;
          _nw25[(new BigNumber(2)).toNumber()] = new BigNumber(856);
          _nw25[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_dafny.Seq.of(new BigNumber(342))).Elements) {
              let _166_v16 = _compr_13;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(342)), _166_v16)) {
                _coll13.push([(_166_v16).multipliedBy(_148_v12),_148_v12]);
              }
            }
            return _coll13;
          }()).length));
          _nw25[(new BigNumber(4)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(5)).toNumber()] = _module.__default.fm31(_136_v1, new BigNumber(924), _162_cf96, _dafny.Seq.of(_136_v1), _134_globalState);
          _nw25[(new BigNumber(6)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(7)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(8)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_164_v17,_148_v12)).length);
          _nw25[(new BigNumber(10)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(11)).toNumber()] = _148_v12;
          _nw25[(new BigNumber(12)).toNumber()] = _148_v12;
          _165_v18 = _nw25;
          let _167_v19;
          _167_v19 = _module.D17.create_DC39(_165_v18, _136_v1, _136_v1, _148_v12, _148_v12);
          let _168_v20;
          _168_v20 = _module.D10.create_DC22(_165_v18);
          let _169_v21;
          _169_v21 = _dafny.Seq.of(_165_v18, _165_v18);
          let _170_v22;
          let _nw26 = Array((new BigNumber(27)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = (_167_v19).dtor_cf57;
          _nw26[(_dafny.ONE).toNumber()] = _165_v18;
          _nw26[(new BigNumber(2)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(3)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(4)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(5)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(6)).toNumber()] = (_168_v20).dtor_cf36;
          _nw26[(new BigNumber(7)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(8)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(9)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(10)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(11)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(12)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(13)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(14)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(15)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(16)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(17)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(18)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(19)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(20)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(21)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(22)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(23)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(24)).toNumber()] = (_169_v21)[_module.__default.safeIndex(_module.__default.fm31(_162_cf96, _148_v12, _161_cf97, _135_v0, _134_globalState), new BigNumber((_169_v21).length))];
          _nw26[(new BigNumber(25)).toNumber()] = _165_v18;
          _nw26[(new BigNumber(26)).toNumber()] = _165_v18;
          _170_v22 = _nw26;
          let _171_v23;
          let _nw27 = Array((new BigNumber(17)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _170_v22;
          _nw27[(_dafny.ONE).toNumber()] = _170_v22;
          _nw27[(new BigNumber(2)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(3)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(4)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(5)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(6)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(7)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(8)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(9)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(10)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(11)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(12)).toNumber()] = ((_162_cf96) ? (_170_v22) : (_170_v22));
          _nw27[(new BigNumber(13)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(14)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(15)).toNumber()] = _170_v22;
          _nw27[(new BigNumber(16)).toNumber()] = _170_v22;
          _171_v23 = _nw27;
          let _index19 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_171_v23).length));
          (_171_v23)[_index19] = _170_v22;
          let _172_v24;
          _172_v24 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_148_v12,_148_v12));
          let _173_v25;
          _173_v25 = _dafny.MultiSet.fromElements(_162_cf96, _module.__default.fm21(_148_v12, _172_v24, _134_globalState));
          let _index20 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_171_v23).length));
          let _rhs18 = _163_v15;
          let _rhs19 = _170_v22;
          let _rhs20 = (_dafny.MultiSet.fromElements(_161_cf97)).IsProperSubsetOf((_dafny.MultiSet.fromElements(_161_cf97)).Intersect(_173_v25));
          let _lhs15 = _171_v23;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_171_v23).length));
          let _lhs17 = _134_globalState;
          _163_v15 = _rhs18;
          _lhs15[_lhs16] = _rhs19;
          _lhs17.f0 = _rhs20;
          let _index21 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_165_v18).length));
          (_165_v18)[_index21] = _148_v12;
          let _174_v26;
          _174_v26 = _dafny.Map.Empty.slice().updateUnsafe(_148_v12,_module.__default.safeModuloInt((_dafny.ZERO).minus(_148_v12), _148_v12));
          let _index22 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_165_v18).length));
          (_165_v18)[_index22] = (((_174_v26).contains(new BigNumber(298))) ? ((_174_v26).get(new BigNumber(298))) : (_148_v12));
          let _175_v27;
          let _nw28 = new _module.C5();
          _nw28.__ctor(_163_v15, new BigNumber((_dafny.Set.fromElements(true)).length), _161_cf97, true);
          _175_v27 = _nw28;
          let _176_v28;
          _176_v28 = _dafny.Map.Empty.slice().updateUnsafe(_161_cf97,_175_v27);
          let _177_v29;
          _177_v29 = _dafny.Seq.of(_148_v12, (_165_v18)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_165_v18).length))]);
          let _178_v30;
          _178_v30 = _module.D23.create_DC59(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("xa"), _module.__default.safeIndex(new BigNumber((_177_v29).length), new BigNumber((_dafny.Seq.UnicodeFromString("xa")).length)), (_163_v15)[_module.__default.safeIndex((_module.D8.create_DC19(_148_v12, (_165_v18)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_165_v18).length))])).dtor_cf32, new BigNumber((_163_v15).length))]), new BigNumber((_173_v25).cardinality()), _148_v12);
          let _179_v31;
          _179_v31 = _module.D1.create_DC5(_161_cf97, _module.__default.fm4(_136_v1, _134_globalState), new BigNumber((_dafny.Set.fromElements(_module.__default.fm49(new BigNumber((_176_v28).length), _134_globalState), _178_v30)).length), (((_149_v13).contains(new BigNumber(430))) ? ((_149_v13).get(new BigNumber(430))) : (new BigNumber((_174_v26).length))));
          _163_v15 = _dafny.Seq.update(_dafny.Seq.Concat(((_136_v1) ? (_163_v15) : (_163_v15)), _module.__default.fm30(_179_v31, new BigNumber((_163_v15).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_148_v12,(_dafny.ZERO).minus(_148_v12))).length), _134_globalState)), _module.__default.safeIndex(_148_v12, new BigNumber((_dafny.Seq.Concat(((_136_v1) ? (_163_v15) : (_163_v15)), _module.__default.fm30(_179_v31, new BigNumber((_163_v15).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_148_v12,(_dafny.ZERO).minus(_148_v12))).length), _134_globalState))).length)), _164_v17);
          _163_v15 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_180_v17) => function (_181_i1) {
            return _180_v17;
          })(_164_v17)), _module.__default.safeIndex(_148_v12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_182_v17) => function (_183_i1) {
            return _182_v17;
          })(_164_v17))).length)), _164_v17), _module.__default.safeIndex((_165_v18)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_165_v18).length))], new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_184_v17) => function (_185_i1) {
            return _184_v17;
          })(_164_v17)), _module.__default.safeIndex(_148_v12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(452)), ((_186_v17) => function (_187_i1) {
            return _186_v17;
          })(_164_v17))).length)), _164_v17)).length)), (_163_v15)[_module.__default.safeIndex(_148_v12, new BigNumber((_163_v15).length))]);
        } else if (_source4.is_DC61) {
          let _188___mcc_h5 = (_source4).cf98;
          let _189_cf98 = _188___mcc_h5;
          let _190_v32;
          let _nw29 = Array((new BigNumber(27)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _136_v1;
          _nw29[(_dafny.ONE).toNumber()] = _136_v1;
          _nw29[(new BigNumber(2)).toNumber()] = false;
          _nw29[(new BigNumber(3)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(4)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(5)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(6)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(7)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(8)).toNumber()] = !(_136_v1);
          _nw29[(new BigNumber(9)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(10)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(11)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(12)).toNumber()] = false;
          _nw29[(new BigNumber(13)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(14)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(15)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(16)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(17)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(18)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(19)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(20)).toNumber()] = true;
          _nw29[(new BigNumber(21)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(22)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(23)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(24)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(25)).toNumber()] = _136_v1;
          _nw29[(new BigNumber(26)).toNumber()] = _136_v1;
          _190_v32 = _nw29;
          let _191_v33;
          _191_v33 = _dafny.Map.Empty.slice().updateUnsafe(_190_v32,_148_v12);
          _191_v33 = (_191_v33).Merge(_191_v33);
          _148_v12 = _148_v12;
          let _192_v34;
          _192_v34 = _dafny.Seq.UnicodeFromString("kpqwyym");
          let _193_v35;
          _193_v35 = new _dafny.CodePoint('a'.codePointAt(0));
          _192_v34 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("bedn"), _module.__default.safeIndex((_189_cf98).minus(_189_cf98), new BigNumber((_dafny.Seq.UnicodeFromString("bedn")).length)), _193_v35);
          let _194_v36;
          _194_v36 = _module.D2.create_DC8(_190_v32, _dafny.Seq.UnicodeFromString("pvrbrm"));
          let _195_v37;
          _195_v37 = _module.D1.create_DC5(_136_v1, _136_v1, _148_v12, _189_cf98);
          let _196_v38;
          let _nw30 = new _module.C5();
          _nw30.__ctor(_192_v34, _189_cf98, !(false), _136_v1);
          _196_v38 = _nw30;
          let _197_v39;
          _197_v39 = _dafny.Map.Empty.slice().updateUnsafe(_136_v1,_196_v38);
          let _198_v40;
          let _nw31 = Array((new BigNumber(18)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _192_v34;
          _nw31[(_dafny.ONE).toNumber()] = ((_136_v1) ? (_192_v34) : (_192_v34));
          _nw31[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(((_136_v1) ? (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), _193_v35)) : (_192_v34)), _module.__default.safeIndex(_189_cf98, new BigNumber((((_136_v1) ? (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), _193_v35)) : (_192_v34))).length)), _193_v35);
          _nw31[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_193_v35), _192_v34);
          _nw31[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_192_v34, (_194_v36).dtor_cf14);
          _nw31[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(531)), ((_199_v35) => function (_200_i2) {
            return _199_v35;
          })(_193_v35));
          _nw31[(new BigNumber(6)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), ((_201_v35) => function (_202_i3) {
            return _201_v35;
          })(_193_v35)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_135_v0, _module.__default.safeIndex(_148_v12, new BigNumber((_135_v0).length)), !(_136_v1))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), ((_203_v35) => function (_204_i3) {
            return _203_v35;
          })(_193_v35))).length)), _193_v35);
          _nw31[(new BigNumber(8)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(9)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(10)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(11)).toNumber()] = _module.__default.fm30(_195_v37, new BigNumber(698), (_dafny.ZERO).minus(_148_v12), _134_globalState);
          _nw31[(new BigNumber(12)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(13)).toNumber()] = ((_module.__default.fm4(true, _134_globalState)) ? (_192_v34) : (_192_v34));
          _nw31[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_192_v34, _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), _193_v35, (_192_v34)[_module.__default.safeIndex(new BigNumber(((_197_v39).update((_196_v38).f4, _196_v38)).length), new BigNumber((_192_v34).length))]));
          _nw31[(new BigNumber(15)).toNumber()] = _192_v34;
          _nw31[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_193_v35);
          _nw31[(new BigNumber(17)).toNumber()] = _192_v34;
          _198_v40 = _nw31;
          let _index23 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_198_v40).length));
          (_198_v40)[_index23] = _192_v34;
          let _205_v41;
          let _nw32 = new _module.C3();
          _nw32.__ctor(_136_v1, (_196_v38).f4);
          _205_v41 = _nw32;
          let _index24 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_198_v40).length));
          let _rhs21 = _dafny.Seq.Concat(_dafny.Seq.Concat(_192_v34, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-641)), ((_206_v35) => function (_207_i4) {
            return _206_v35;
          })(_193_v35))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), ((_208_v35) => function (_209_i5) {
            return _208_v35;
          })(_193_v35)), _192_v34));
          let _rhs22 = _192_v34;
          let _rhs23 = _205_v41;
          let _rhs24 = _189_cf98;
          let _rhs25 = _dafny.Seq.Concat(_192_v34, _192_v34);
          let _lhs18 = _198_v40;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_198_v40).length));
          _lhs18[_lhs19] = _rhs21;
          _192_v34 = _rhs22;
          _205_v41 = _rhs23;
          _189_cf98 = _rhs24;
          _192_v34 = _rhs25;
        } else {
          let _210___mcc_h6 = (_source4).cf92;
          let _211_cf92 = _210___mcc_h6;
          let _212_v43;
          _212_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_module.__default.fm50(_136_v1, _134_globalState)).Keys.Elements) {
              let _213_v42 = _compr_14;
              if ((_module.__default.fm50(_136_v1, _134_globalState)).contains(_213_v42)) {
                _coll14.add(_213_v42);
              }
            }
            return _coll14;
          }()).length));
          let _214_v44;
          _214_v44 = _dafny.Seq.of(_module.__default.fm28(_136_v1, _136_v1, false, (_dafny.ZERO).minus(_148_v12), _134_globalState), (_212_v43).Merge(_212_v43), _212_v43);
          let _215_v45;
          _215_v45 = _dafny.Set.fromElements(new BigNumber((_135_v0).length));
          let _216_v46;
          _216_v46 = _dafny.Map.Empty.slice().updateUnsafe(_148_v12,_136_v1);
          _214_v44 = (((_dafny.Set.fromElements(new BigNumber((_216_v46).length))).IsSubsetOf(_215_v45)) ? (_dafny.Seq.of(_212_v43, _212_v43, _212_v43, _212_v43)) : (_214_v44));
          let _217_v47;
          let _init3 = ((_218_v12) => function (_219_i6) {
            return _module.__default.safeModuloInt(_219_i6, _218_v12);
          })(_148_v12);
          let _nw33 = Array((new BigNumber(27)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw33.length); _i0_3++) {
            _nw33[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _217_v47 = _nw33;
          let _220_v48;
          let _nw34 = new _module.C5();
          _nw34.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_221_i7) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }), _148_v12, _136_v1, !(_136_v1));
          _220_v48 = _nw34;
          let _222_v49;
          _222_v49 = _dafny.Map.Empty.slice().updateUnsafe(_220_v48,_217_v47);
          let _223_v50;
          _223_v50 = _module.D21.create_DC52(_148_v12, _217_v47);
          let _224_v51;
          _224_v51 = _dafny.Seq.of(_217_v47, _217_v47, (_223_v50).dtor_cf85);
          let _225_v52;
          let _nw35 = Array((new BigNumber(14)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _217_v47;
          _nw35[(_dafny.ONE).toNumber()] = _217_v47;
          _nw35[(new BigNumber(2)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(3)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(4)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(5)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(6)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(7)).toNumber()] = ((_136_v1) ? (_217_v47) : (_217_v47));
          _nw35[(new BigNumber(8)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(9)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(10)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(11)).toNumber()] = (((_222_v49).contains(_220_v48)) ? ((_222_v49).get(_220_v48)) : (_217_v47));
          _nw35[(new BigNumber(12)).toNumber()] = _217_v47;
          _nw35[(new BigNumber(13)).toNumber()] = (_224_v51)[_module.__default.safeIndex(_148_v12, new BigNumber((_224_v51).length))];
          _225_v52 = _nw35;
          _225_v52 = _225_v52;
          let _226_v53;
          _226_v53 = _dafny.Seq.UnicodeFromString("gvdye");
          let _227_v54;
          let _nw36 = new _module.C0();
          _nw36.__ctor();
          _227_v54 = _nw36;
          let _228_v55;
          _228_v55 = _dafny.Seq.of(_227_v54);
          let _229_v56;
          let _nw37 = Array((new BigNumber(8)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _227_v54;
          _nw37[(_dafny.ONE).toNumber()] = _227_v54;
          _nw37[(new BigNumber(2)).toNumber()] = _227_v54;
          _nw37[(new BigNumber(3)).toNumber()] = _227_v54;
          _nw37[(new BigNumber(4)).toNumber()] = _227_v54;
          _nw37[(new BigNumber(5)).toNumber()] = _227_v54;
          _nw37[(new BigNumber(6)).toNumber()] = _227_v54;
          _nw37[(new BigNumber(7)).toNumber()] = (_228_v55)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_148_v12, _148_v12)).length), new BigNumber((_228_v55).length))];
          _229_v56 = _nw37;
          let _230_v57;
          _230_v57 = _dafny.Map.Empty.slice().updateUnsafe(true,_229_v56);
          let _231_v58;
          _231_v58 = _module.D19.create_DC44(_230_v57);
          let _index25 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length));
          (_217_v47)[_index25] = (_148_v12).plus(new BigNumber((_dafny.Seq.of(_220_v48)).length));
          let _232_v59;
          _232_v59 = _dafny.Seq.of(_148_v12);
          let _index26 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length));
          let _rhs26 = _226_v53;
          let _rhs27 = _136_v1;
          let _rhs28 = _231_v58;
          let _rhs29 = (_dafny.ZERO).minus(_148_v12);
          let _rhs30 = (_227_v54).fm10((_232_v59)[_module.__default.safeIndex((_dafny.ZERO).minus(_148_v12), new BigNumber((_232_v59).length))], _136_v1, _149_v13, _134_globalState);
          let _lhs20 = _134_globalState;
          let _lhs21 = _217_v47;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length));
          _226_v53 = _rhs26;
          _lhs20.f0 = _rhs27;
          _231_v58 = _rhs28;
          _148_v12 = _rhs29;
          _lhs21[_lhs22] = _rhs30;
          let _233_v61;
          _233_v61 = _module.D1.create_DC5(_module.__default.fm4(_136_v1, _134_globalState), true, new BigNumber((function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of _dafny.IntegerRange(new BigNumber(25), new BigNumber(-331))) {
    let _234_v60 = _compr_15;
    if (((new BigNumber(25)).isLessThanOrEqualTo(_234_v60)) && ((_234_v60).isLessThan(new BigNumber(-331)))) {
      _coll15.push([_module.__default.safeDivisionInt(_234_v60, _148_v12),new BigNumber(869)]);
    }
  }
  return _coll15;
}()).length), (_227_v54).fm10(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_217_v47)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length))],new BigNumber(-742))).length), _136_v1, _149_v13, _134_globalState));
          let _235_v62;
          _235_v62 = _module.D19.create_DC46((_dafny.ZERO).minus((_217_v47)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length))]), _136_v1, _136_v1, _233_v61, (_dafny.ZERO).minus((_217_v47)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length))]));
          let _index27 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_217_v47).length));
          (_217_v47)[_index27] = (_235_v62).dtor_cf78;
        }
        let _236_v63;
        _236_v63 = new _dafny.CodePoint('y'.codePointAt(0));
        let _237_v64;
        let _nw38 = new _module.C4();
        _nw38.__ctor(_136_v1, _148_v12, _136_v1, _136_v1);
        _237_v64 = _nw38;
        let _238_v65;
        _238_v65 = _dafny.Map.Empty.slice().updateUnsafe(_237_v64,!(_237_v64.f8));
        let _239_v66;
        let _nw39 = Array((new BigNumber(25)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = _136_v1;
        _nw39[(_dafny.ONE).toNumber()] = false;
        _nw39[(new BigNumber(2)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(3)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(4)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(5)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(6)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(7)).toNumber()] = _module.__default.fm18(_148_v12, _236_v63, _148_v12, _134_globalState);
        _nw39[(new BigNumber(8)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(9)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(10)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(11)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(12)).toNumber()] = !(_136_v1);
        _nw39[(new BigNumber(13)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(14)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(15)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(16)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(17)).toNumber()] = (((_238_v65).contains(_237_v64)) ? ((_238_v65).get(_237_v64)) : (_136_v1));
        _nw39[(new BigNumber(18)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(19)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(20)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(21)).toNumber()] = _237_v64.f8;
        _nw39[(new BigNumber(22)).toNumber()] = _237_v64.f8;
        _nw39[(new BigNumber(23)).toNumber()] = _136_v1;
        _nw39[(new BigNumber(24)).toNumber()] = _136_v1;
        _239_v66 = _nw39;
        let _240_v67;
        _240_v67 = _module.D2.create_DC8(_239_v66, _dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), ((_241_v63) => function (_242_i8) {
  return _241_v63;
})(_236_v63)));
        let _pat_let_tv4 = _239_v66;
        let _source5 = function (_pat_let4_0) {
          return function (_243_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_244_dt__update_hcf13_h0) {
                return _module.D2.create_DC8(_244_dt__update_hcf13_h0, (_243_dt__update__tmp_h1).dtor_cf14);
              }(_pat_let5_0);
            }(_pat_let_tv4);
          }(_pat_let4_0);
        }(_240_v67);
        if (_source5.is_DC8) {
          let _245___mcc_h7 = (_source5).cf13;
          let _246___mcc_h8 = (_source5).cf14;
          let _247_cf14 = _246___mcc_h8;
          let _248_cf13 = _245___mcc_h7;
          let _249_v68;
          let _250_v69;
          let _251_v70;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_237_v64).m6(_237_v64.f8, _134_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _249_v68 = _out0;
          _250_v69 = _out1;
          _251_v70 = _out2;
          _148_v12 = (_dafny.ZERO).minus(((_237_v64).f9).multipliedBy(_250_v69));
          _237_v64 = _237_v64;
          let _252_v71;
          let _nw40 = new _module.C4();
          _nw40.__ctor(_136_v1, (_237_v64).f9, _237_v64.f8, _237_v64.f8);
          _252_v71 = _nw40;
          let _253_v72;
          _253_v72 = _dafny.Set.fromElements(_252_v71);
          let _254_v74;
          _254_v74 = _dafny.Map.Empty.slice().updateUnsafe(_250_v69,_136_v1);
          let _255_v75;
          _255_v75 = _dafny.Map.Empty.slice().updateUnsafe((_237_v64).f9,new BigNumber((_254_v74).length));
          let _256_v76;
          _256_v76 = _dafny.Map.Empty.slice().updateUnsafe(_249_v68,_148_v12);
          let _257_v78;
          _257_v78 = _dafny.Seq.of(_252_v71);
          let _258_v79;
          _258_v79 = _dafny.Seq.of(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of (_255_v75).Keys.Elements) {
              let _259_v73 = _compr_16;
              if ((_255_v75).contains(_259_v73)) {
                _coll16.push([_module.__default.safeDivisionInt(_259_v73, _148_v12),(_252_v71).f4]);
              }
            }
            return _coll16;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_256_v76).length),_237_v64.f8), _254_v74, (function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_254_v74).Keys.Elements) {
              let _260_v77 = _compr_17;
              if ((_254_v74).contains(_260_v77)) {
                _coll17.push([_module.__default.safeDivisionInt(_260_v77, new BigNumber(600)),_237_v64.f8]);
              }
            }
            return _coll17;
          }()).update(_148_v12, _136_v1), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_257_v78).length),_252_v71.f3));
          let _261_v80;
          let _262_v81;
          let _out3;
          let _out4;
          let _outcollector1 = (_237_v64).m0((_253_v72).IsSubsetOf(_253_v72), ((_258_v79)[_module.__default.safeIndex((_237_v64).f9, new BigNumber((_258_v79).length))]).update((_237_v64).f9, _237_v64.f8), (_237_v64).f9, _134_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _261_v80 = _out3;
          _262_v81 = _out4;
        } else {
          let _263___mcc_h9 = (_source5).cf12;
          let _264_cf12 = _263___mcc_h9;
          let _265_v82;
          let _nw41 = new _module.C2();
          _nw41.__ctor(_237_v64.f8, false, true, _136_v1);
          _265_v82 = _nw41;
          let _266_v83;
          let _out5;
          _out5 = (_237_v64).m1(((_237_v64).f9).multipliedBy((_237_v64).f9), _265_v82, _134_globalState);
          _266_v83 = _out5;
          let _index28 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_239_v66).length));
          (_239_v66)[_index28] = (_265_v82).f4;
          let _267_v84;
          _267_v84 = _dafny.Seq.of(_266_v83);
          let _268_v85;
          _268_v85 = _dafny.MultiSet.fromElements(_136_v1);
          let _269_v86;
          _269_v86 = _dafny.Map.Empty.slice().updateUnsafe(_267_v84,_268_v85);
          let _270_v87;
          let _nw42 = Array((_dafny.ONE).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _269_v86;
          _270_v87 = _nw42;
          let _index29 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_270_v87).length));
          (_270_v87)[_index29] = _269_v86;
          let _271_v88;
          _271_v88 = _dafny.Map.Empty.slice().updateUnsafe(_237_v64.f8,_237_v64.f8);
          let _272_v89;
          _272_v89 = _269_v86;
          let _index30 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_239_v66).length));
          let _index31 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_270_v87).length));
          let _rhs31 = (_237_v64).f9;
          let _rhs32 = (((_271_v88).contains(_237_v64.f8)) ? ((_271_v88).get(_237_v64.f8)) : (_237_v64.f8));
          let _rhs33 = (_272_v89);
          let _lhs23 = _239_v66;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_239_v66).length));
          let _lhs25 = _270_v87;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_270_v87).length));
          _148_v12 = _rhs31;
          _lhs23[_lhs24] = _rhs32;
          _lhs25[_lhs26] = _rhs33;
          let _273_v90;
          _273_v90 = _dafny.Seq.UnicodeFromString("y");
          let _274_v91;
          _274_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_273_v90).length),(_239_v66)[_module.__default.safeIndex(new BigNumber(754), new BigNumber((_239_v66).length))]);
          let _275_v92;
          let _276_v93;
          let _out6;
          let _out7;
          let _outcollector2 = (_265_v82).m0(_237_v64.f8, (_274_v91).Merge(_274_v91), new BigNumber(-15), _134_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _275_v92 = _out6;
          _276_v93 = _out7;
          let _277_v94;
          _277_v94 = _module.D17.create_DC38(_273_v90);
          _273_v90 = (_277_v94).dtor_cf56;
        }
        let _278_v95;
        _278_v95 = _dafny.Seq.UnicodeFromString("qgg");
        let _279_v96;
        let _nw43 = new _module.C5();
        _nw43.__ctor(_278_v95, (_dafny.ZERO).minus(_148_v12), true, ((_237_v64).f9).isEqualTo(_148_v12));
        _279_v96 = _nw43;
        _279_v96 = _279_v96;
        let _280_v97;
        let _281_v98;
        let _282_v99;
        let _out8;
        let _out9;
        let _out10;
        let _outcollector3 = (_237_v64).m6(_136_v1, _134_globalState);
        _out8 = _outcollector3[0];
        _out9 = _outcollector3[1];
        _out10 = _outcollector3[2];
        _280_v97 = _out8;
        _281_v98 = _out9;
        _282_v99 = _out10;
        let _283_v100;
        _283_v100 = _dafny.Map.Empty.slice().updateUnsafe(_237_v64.f8,_279_v96.f7);
        let _284_v101;
        _284_v101 = _module.D1.create_DC5(_136_v1, true, new BigNumber((_dafny.Seq.update(_278_v95, _module.__default.safeIndex((((_283_v100).contains(_136_v1)) ? ((_283_v100).get(_136_v1)) : (_281_v98)), new BigNumber((_278_v95).length)), _236_v63)).length), new BigNumber((_dafny.Seq.UnicodeFromString("wgwyb")).length));
        let _285_v102;
        _285_v102 = _dafny.Seq.of(_module.__default.fm31(_237_v64.f8, _279_v96.f7, _136_v1, _135_v0, _134_globalState));
        let _286_v103;
        _286_v103 = _module.D0.create_DC0(_148_v12);
        _278_v95 = _module.__default.fm30(_284_v101, (_285_v102)[_module.__default.safeIndex(_279_v96.f7, new BigNumber((_285_v102).length))], ((_280_v97) ? ((_286_v103).dtor_cf0) : (new BigNumber((_dafny.Seq.UnicodeFromString("vcroj")).length))), _134_globalState);
      } else {
        let _287_v104;
        _287_v104 = new BigNumber(29);
        let _288_v105;
        _288_v105 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_287_v104));
        _287_v104 = (((_288_v105).contains(_287_v104)) ? ((_288_v105).get(_287_v104)) : ((_287_v104).minus(_287_v104)));
        let _289_v106;
        _289_v106 = _dafny.Map.Empty.slice().updateUnsafe(_287_v104,_287_v104);
        let _290_v107;
        _290_v107 = _dafny.MultiSet.fromElements(_289_v106, _289_v106);
        let _291_v108;
        let _nw44 = Array((new BigNumber(21)).toNumber());
        _nw44[(_dafny.ZERO).toNumber()] = _136_v1;
        _nw44[(_dafny.ONE).toNumber()] = _136_v1;
        _nw44[(new BigNumber(2)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(3)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(4)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(5)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(6)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(7)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(8)).toNumber()] = true;
        _nw44[(new BigNumber(9)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(10)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(11)).toNumber()] = !(_136_v1);
        _nw44[(new BigNumber(12)).toNumber()] = (_135_v0)[_module.__default.safeIndex(_287_v104, new BigNumber((_135_v0).length))];
        _nw44[(new BigNumber(13)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(14)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(15)).toNumber()] = _module.__default.fm21(_287_v104, _290_v107, _134_globalState);
        _nw44[(new BigNumber(16)).toNumber()] = false;
        _nw44[(new BigNumber(17)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(18)).toNumber()] = _136_v1;
        _nw44[(new BigNumber(19)).toNumber()] = !(_136_v1);
        _nw44[(new BigNumber(20)).toNumber()] = _136_v1;
        _291_v108 = _nw44;
        let _292_v109;
        _292_v109 = _dafny.Map.Empty.slice().updateUnsafe(_291_v108,_136_v1);
        let _293_v110;
        _293_v110 = _module.D19.create_DC45(_287_v104, _287_v104, _136_v1, _136_v1, _135_v0);
        let _294_v111;
        let _nw45 = new _module.C0();
        _nw45.__ctor();
        _294_v111 = _nw45;
        let _295_v112;
        _295_v112 = _dafny.Map.Empty.slice().updateUnsafe(_293_v110,_294_v111);
        _292_v109 = (_292_v109).update(_291_v108, (_295_v112).equals(_295_v112));
        let _296_v113;
        let _297_v114;
        let _298_v115;
        let _299_v116;
        let _out11;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m12((_module.__default.fm31(_136_v1, _287_v104, _136_v1, _module.__default.fm34(_287_v104, _module.__default.fm31(_136_v1, (_294_v111).fm10(_287_v104, _136_v1, _288_v105, _134_globalState), _136_v1, _135_v0, _134_globalState), _134_globalState), _134_globalState)).minus(_287_v104), _136_v1, (_136_v1) && (!(!(!(_136_v1)))), (new BigNumber((_292_v109).length)).isLessThanOrEqualTo(_287_v104), _134_globalState);
        _out11 = _outcollector4[0];
        _out12 = _outcollector4[1];
        _out13 = _outcollector4[2];
        _out14 = _outcollector4[3];
        _296_v113 = _out11;
        _297_v114 = _out12;
        _298_v115 = _out13;
        _299_v116 = _out14;
        if (!((_299_v116).f10) || ((_299_v116).f10)) {
          let _300_v117;
          _300_v117 = _dafny.Seq.of((((_288_v105).contains(new BigNumber((_135_v0).length))) ? ((_288_v105).get(new BigNumber((_135_v0).length))) : (_297_v114)), _297_v114, _297_v114, _297_v114);
          let _301_v118;
          _301_v118 = _dafny.Map.Empty.slice().updateUnsafe(_287_v104,_299_v116);
          let _302_v119;
          _302_v119 = _dafny.Map.Empty.slice().updateUnsafe(_301_v118,_135_v0);
          let _303_v120;
          _303_v120 = _dafny.Map.Empty.slice().updateUnsafe((_300_v117)[_module.__default.safeIndex(_287_v104, new BigNumber((_300_v117).length))],_302_v119);
          _303_v120 = (_303_v120).update(_287_v104, (_302_v119).Merge(_302_v119));
          let _304_v121;
          _304_v121 = _dafny.Seq.of(_291_v108);
          let _305_v122;
          let _init4 = ((_306_v104) => function (_307_i9) {
            return (_307_i9).plus(_306_v104);
          })(_287_v104);
          let _nw46 = Array((new BigNumber(29)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw46.length); _i0_4++) {
            _nw46[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _305_v122 = _nw46;
          let _308_v123;
          _308_v123 = _module.D21.create_DC52(_287_v104, _305_v122);
          let _309_v124;
          _309_v124 = _dafny.Seq.of(_308_v123);
          _291_v108 = (_304_v121)[_module.__default.safeIndex((new BigNumber((_309_v124).length)).multipliedBy(_287_v104), new BigNumber((_304_v121).length))];
          _136_v1 = !((_299_v116).f11);
          let _310_v125;
          let _nw47 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _310_v125 = _nw47;
          let _311_v126;
          _311_v126 = _module.D21.create_DC51(_310_v125);
          let _312_v127;
          _312_v127 = _dafny.MultiSet.fromElements(_311_v126, _311_v126, _311_v126, _311_v126);
          let _313_v128;
          let _nw48 = new _module.C3();
          _nw48.__ctor((_299_v116).f11, (_299_v116).f10);
          _313_v128 = _nw48;
          let _314_v129;
          _314_v129 = _dafny.Map.Empty.slice().updateUnsafe((_299_v116).f10,_313_v128);
          let _315_v130;
          _315_v130 = new _dafny.CodePoint('u'.codePointAt(0));
          let _316_v131;
          _316_v131 = _dafny.MultiSet.fromElements(_315_v130);
          let _317_v132;
          _317_v132 = _dafny.Map.Empty.slice().updateUnsafe(_297_v114,true);
          let _318_v133;
          _318_v133 = _dafny.Map.Empty.slice().updateUnsafe((_299_v116).f11,_136_v1);
          let _rhs34 = _312_v127;
          let _rhs35 = _136_v1;
          let _rhs36 = (((_299_v116).fm16(new BigNumber((_314_v129).length), (_299_v116).f10, new BigNumber((_316_v131).cardinality()), _317_v132, _134_globalState)).plus(_297_v114)).plus((_dafny.ZERO).minus(((false) ? (_297_v114) : (new BigNumber((_318_v133).length)))));
          let _lhs27 = _134_globalState;
          _312_v127 = _rhs34;
          _lhs27.f0 = _rhs35;
          _287_v104 = _rhs36;
          _315_v130 = _315_v130;
        } else {
          let _319_v134;
          let _nw49 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _319_v134 = _nw49;
          let _320_v135;
          _320_v135 = _dafny.MultiSet.fromElements(_319_v134, _319_v134);
          let _321_v136;
          _321_v136 = _dafny.MultiSet.fromElements(_320_v135, _320_v135, _dafny.MultiSet.fromElements(_319_v134, _319_v134));
          let _322_v137;
          _322_v137 = _dafny.MultiSet.fromElements((_299_v116).f11, _136_v1);
          _297_v114 = (((_321_v136).contains(_320_v135)) ? ((_321_v136).get(_320_v135)) : ((_dafny.ZERO).minus((((_322_v137).contains(false)) ? ((_322_v137).get(false)) : (_287_v104)))));
          let _index32 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_319_v134).length));
          (_319_v134)[_index32] = _287_v104;
          let _index33 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_291_v108).length));
          (_291_v108)[_index33] = (_287_v104).isLessThan(_297_v114);
          let _index34 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_319_v134).length));
          let _index35 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_291_v108).length));
          let _rhs37 = _297_v114;
          let _rhs38 = _136_v1;
          let _lhs28 = _319_v134;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_319_v134).length));
          let _lhs30 = _291_v108;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_291_v108).length));
          _lhs28[_lhs29] = _rhs37;
          _lhs30[_lhs31] = _rhs38;
          let _323_v138;
          _323_v138 = _dafny.Seq.UnicodeFromString("g");
          let _324_v139;
          let _nw50 = new _module.C1();
          _nw50.__ctor(_323_v138, (false) || ((_291_v108)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_291_v108).length))]));
          _324_v139 = _nw50;
          let _325_v140;
          let _nw51 = Array((new BigNumber(9)).toNumber());
          _nw51[(_dafny.ZERO).toNumber()] = _319_v134;
          _nw51[(_dafny.ONE).toNumber()] = _319_v134;
          _nw51[(new BigNumber(2)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(3)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(4)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(5)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(6)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(7)).toNumber()] = _319_v134;
          _nw51[(new BigNumber(8)).toNumber()] = _319_v134;
          _325_v140 = _nw51;
          let _index36 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_325_v140).length));
          (_325_v140)[_index36] = (((_299_v116).f10) ? (_319_v134) : (_319_v134));
          let _index37 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_325_v140).length));
          (_325_v140)[_index37] = _319_v134;
          _287_v104 = _297_v114;
        }
        let _326_v141;
        let _nw52 = new _module.C6();
        _nw52.__ctor((_299_v116).f10, (_299_v116).f11, (_299_v116).f10);
        _326_v141 = _nw52;
        let _327_v142;
        _327_v142 = _dafny.Map.Empty.slice().updateUnsafe(_326_v141,_136_v1);
        let _328_v143;
        _328_v143 = _dafny.Map.Empty.slice().updateUnsafe(_297_v114,(_299_v116).f10);
        _327_v142 = (_327_v142).update(_326_v141, (((_328_v143).contains(_297_v114)) ? ((_328_v143).get(_297_v114)) : ((_299_v116).f11)));
      }
      let _329_v144;
      _329_v144 = new BigNumber(937);
      let _330_v145;
      _330_v145 = _dafny.Set.fromElements(_136_v1, _136_v1);
      let _331_v146;
      _331_v146 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_330_v145).length)));
      _329_v144 = _module.__default.safeModuloInt(_329_v144, _module.__default.safeDivisionInt(_329_v144, (_331_v146)[_module.__default.safeIndex(new BigNumber((_135_v0).length), new BigNumber((_331_v146).length))]));
      let _332_v147;
      let _nw53 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _332_v147 = _nw53;
      let _index38 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
      (_332_v147)[_index38] = _329_v144;
      let _333_v148;
      _333_v148 = _dafny.Seq.UnicodeFromString("d");
      let _334_v149;
      _334_v149 = _module.D1.create_DC5(_136_v1, _136_v1, _329_v144, new BigNumber((_333_v148).length));
      let _pat_let_tv5 = _329_v144;
      let _335_v150;
      _335_v150 = _module.D19.create_DC46(_329_v144, (_module.D22.create_DC56(_329_v144, false)).dtor_cf90, _136_v1, function (_pat_let6_0) {
  return function (_336_dt__update__tmp_h2) {
    return function (_pat_let7_0) {
      return function (_337_dt__update_hcf10_h0) {
        return function (_pat_let8_0) {
          return function (_338_dt__update_hcf11_h0) {
            return _module.D1.create_DC5((_336_dt__update__tmp_h2).dtor_cf8, (_336_dt__update__tmp_h2).dtor_cf9, _337_dt__update_hcf10_h0, _338_dt__update_hcf11_h0);
          }(_pat_let8_0);
        }(_pat_let_tv5);
      }(_pat_let7_0);
    }(new BigNumber(747));
  }(_pat_let6_0);
}(_334_v149), (_dafny.ZERO).minus(_329_v144));
      let _pat_let_tv6 = _333_v148;
      let _pat_let_tv7 = _329_v144;
      let _index39 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
      let _rhs39 = function (_source6) {
        if (_source6.is_DC45) {
          let _339___mcc_h10 = (_source6).cf69;
          let _340___mcc_h11 = (_source6).cf70;
          let _341___mcc_h12 = (_source6).cf71;
          let _342___mcc_h13 = (_source6).cf72;
          let _343___mcc_h14 = (_source6).cf73;
          let _344_cf73 = _343___mcc_h14;
          let _345_cf72 = _342___mcc_h13;
          let _346_cf71 = _341___mcc_h12;
          let _347_cf70 = _340___mcc_h11;
          let _348_cf69 = _339___mcc_h10;
          return _347_cf70;
        } else if (_source6.is_DC46) {
          let _349___mcc_h15 = (_source6).cf74;
          let _350___mcc_h16 = (_source6).cf75;
          let _351___mcc_h17 = (_source6).cf76;
          let _352___mcc_h18 = (_source6).cf77;
          let _353___mcc_h19 = (_source6).cf78;
          let _354_cf78 = _353___mcc_h19;
          let _355_cf77 = _352___mcc_h18;
          let _356_cf76 = _351___mcc_h17;
          let _357_cf75 = _350___mcc_h16;
          let _358_cf74 = _349___mcc_h15;
          return (new BigNumber(964)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length));
        } else if (_source6.is_DC44) {
          let _359___mcc_h20 = (_source6).cf68;
          let _360_cf68 = _359___mcc_h20;
          return new BigNumber((_pat_let_tv6).length);
        } else {
          let _361___mcc_h21 = (_source6).cf79;
          let _362_cf79 = _361___mcc_h21;
          return (_dafny.ZERO).minus(_pat_let_tv7);
        }
      }(_335_v150);
      let _rhs40 = _329_v144;
      let _rhs41 = _329_v144;
      let _lhs32 = _332_v147;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
      _lhs32[_lhs33] = _rhs39;
      _329_v144 = _rhs40;
      _329_v144 = _rhs41;
      let _363_v151;
      let _nw54 = Array((new BigNumber(12)).toNumber()).fill(false);
      _363_v151 = _nw54;
      _363_v151 = _363_v151;
      let _364_v152;
      let _nw55 = new _module.C4();
      _nw55.__ctor((new BigNumber((_330_v145).length)).isLessThanOrEqualTo((_332_v147)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length))]), (_332_v147)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length))], _136_v1, _136_v1);
      _364_v152 = _nw55;
      _364_v152 = _364_v152;
      let _365_v153;
      let _nw56 = new _module.C4();
      _nw56.__ctor(_136_v1, _329_v144, false, _136_v1);
      _365_v153 = _nw56;
      let _366_v154;
      let _out15;
      _out15 = (_364_v152).m1(((_332_v147)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length))]).multipliedBy(new BigNumber(-84)), _365_v153, _134_globalState);
      _366_v154 = _out15;
      let _hi2 = (_366_v154).minus(_329_v144);
      for (let _367_i10 = _366_v154; _367_i10.isLessThan(_hi2); _367_i10 = _367_i10.plus(_dafny.ONE)) {
        let _index40 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
        (_332_v147)[_index40] = (_364_v152).fm0(_134_globalState);
        let _index41 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
        (_332_v147)[_index41] = _366_v154;
        _366_v154 = new BigNumber(763);
        let _368_v155;
        let _nw57 = new _module.C6();
        _nw57.__ctor(_364_v152.f8, false, (_365_v153).f4);
        _368_v155 = _nw57;
        let _369_v156;
        _369_v156 = _dafny.Set.fromElements(_368_v155);
        let _370_v157;
        _370_v157 = _dafny.Map.Empty.slice().updateUnsafe(!(_136_v1),_369_v156);
        let _371_v158;
        _371_v158 = _module.D7.create_DC17(_367_i10, _365_v153.f3, _364_v152.f8, new BigNumber(((((_370_v157).contains(true)) ? ((_370_v157).get(true)) : (_369_v156))).length));
        let _pat_let_tv8 = _334_v149;
        let _source7 = function (_pat_let9_0) {
          return function (_372_dt__update__tmp_h3) {
            return function (_pat_let10_0) {
              return function (_373_dt__update_hcf27_h0) {
                return _module.D7.create_DC17((_372_dt__update__tmp_h3).dtor_cf26, _373_dt__update_hcf27_h0, (_372_dt__update__tmp_h3).dtor_cf28, (_372_dt__update__tmp_h3).dtor_cf29);
              }(_pat_let10_0);
            }((_pat_let_tv8).dtor_cf9);
          }(_pat_let9_0);
        }(_371_v158);
        if (_source7.is_DC17) {
          let _374___mcc_h22 = (_source7).cf26;
          let _375___mcc_h23 = (_source7).cf27;
          let _376___mcc_h24 = (_source7).cf28;
          let _377___mcc_h25 = (_source7).cf29;
          let _378_cf29 = _377___mcc_h25;
          let _379_cf28 = _376___mcc_h24;
          let _380_cf27 = _375___mcc_h23;
          let _381_cf26 = _374___mcc_h22;
          let _382_v159;
          _382_v159 = _dafny.Map.Empty.slice().updateUnsafe(_368_v155,(_364_v152).f9);
          _382_v159 = (_382_v159).update(_368_v155, _module.__default.safeDivisionInt((_332_v147)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length))], _381_cf26));
          (_365_v153).f3 = ((_368_v155.f5) ? ((_dafny.Set.fromElements(_365_v153.f3)).IsDisjointFrom(_330_v145)) : (_365_v153.f3));
          _364_v152 = _364_v152;
          let _383_v160;
          let _nw58 = new _module.C1();
          _nw58.__ctor(_dafny.Seq.Concat(_333_v148, _dafny.Seq.UnicodeFromString("e")), _364_v152.f8);
          _383_v160 = _nw58;
        } else {
          let _384___mcc_h26 = (_source7).cf25;
          let _385_cf25 = _384___mcc_h26;
          let _index42 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
          (_332_v147)[_index42] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_333_v148, _dafny.Seq.UnicodeFromString("po")), _333_v148)).length);
          let _index43 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length));
          (_332_v147)[_index43] = _module.__default.safeModuloInt((_364_v152).f9, _366_v154);
          let _386_v161;
          _386_v161 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_333_v148).length),_365_v153.f3);
          let _387_v162;
          _387_v162 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_386_v161).length),_365_v153.f3);
          let _388_v163;
          let _389_v164;
          let _out16;
          let _out17;
          let _outcollector5 = (_365_v153).m0(_365_v153.f3, ((_387_v162).update((_364_v152).f9, _364_v152.f8)).Merge(_386_v161), _329_v144, _134_globalState);
          _out16 = _outcollector5[0];
          _out17 = _outcollector5[1];
          _388_v163 = _out16;
          _389_v164 = _out17;
          let _390_v165;
          let _nw59 = new _module.C1();
          _nw59.__ctor(_333_v148, _136_v1);
          _390_v165 = _nw59;
          _390_v165 = _390_v165;
        }
      }
      let _391_v166;
      _391_v166 = _dafny.Map.Empty.slice().updateUnsafe(_364_v152.f8,_136_v1);
      (_365_v153).f3 = (((_391_v166).contains(_136_v1)) ? ((_391_v166).get(_136_v1)) : (_136_v1));
      let _392_v169;
      _392_v169 = _module.D17.create_DC37(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_dafny.Seq.of(_331_v146)).Elements) {
    let _393_v168 = _compr_18;
    if (_dafny.Seq.contains(_dafny.Seq.of(_331_v146), _393_v168)) {
      _coll18.push([_393_v168,new _dafny.CodePoint('l'.codePointAt(0))]);
    }
  }
  return _coll18;
}());
      let _394_v170;
      _394_v170 = _dafny.Set.fromElements(_392_v169);
      let _395_v171;
      _395_v171 = _dafny.Seq.of(_394_v170);
      let _396_v173;
      _396_v173 = _module.D23.create_DC58(function () {
  let _coll19 = new _dafny.Set();
  for (const _compr_19 of (function () {
    let _coll20 = new _dafny.Map();
    for (const _compr_20 of (_395_v171).Elements) {
      let _397_v167 = _compr_20;
      if (_dafny.Seq.contains(_395_v171, _397_v167)) {
        _coll20.push([_397_v167,_366_v154]);
      }
    }
    return _coll20;
  }()).Keys.Elements) {
    let _398_v172 = _compr_19;
    if ((function () {
      let _coll21 = new _dafny.Map();
      for (const _compr_21 of (_395_v171).Elements) {
        let _397_v167 = _compr_21;
        if (_dafny.Seq.contains(_395_v171, _397_v167)) {
          _coll21.push([_397_v167,_366_v154]);
        }
      }
      return _coll21;
    }()).contains(_398_v172)) {
      _coll19.add(_398_v172);
    }
  }
  return _coll19;
}());
      let _pat_let_tv9 = _364_v152;
      let _pat_let_tv10 = _333_v148;
      let _pat_let_tv11 = _364_v152;
      let _pat_let_tv12 = _331_v146;
      let _pat_let_tv13 = _329_v144;
      let _pat_let_tv14 = _331_v146;
      let _pat_let_tv15 = _364_v152;
      let _pat_let_tv16 = _333_v148;
      let _pat_let_tv17 = _365_v153;
      let _pat_let_tv18 = _329_v144;
      let _pat_let_tv19 = _329_v144;
      _366_v154 = function (_source8) {
        if (_source8.is_DC59) {
          let _399___mcc_h27 = (_source8).cf93;
          let _400___mcc_h28 = (_source8).cf94;
          let _401___mcc_h29 = (_source8).cf95;
          let _402_cf95 = _401___mcc_h29;
          let _403_cf94 = _400___mcc_h28;
          let _404_cf93 = _399___mcc_h27;
          return ((_pat_let_tv9).f9).minus(new BigNumber((_pat_let_tv10).length));
        } else if (_source8.is_DC60) {
          let _405___mcc_h30 = (_source8).cf96;
          let _406___mcc_h31 = (_source8).cf97;
          let _407_cf97 = _406___mcc_h31;
          let _408_cf96 = _405___mcc_h30;
          return ((_dafny.ZERO).minus((_pat_let_tv11).f9)).minus((_pat_let_tv12)[_module.__default.safeIndex(_pat_let_tv13, new BigNumber((_pat_let_tv14).length))]);
        } else if (_source8.is_DC61) {
          let _409___mcc_h32 = (_source8).cf98;
          let _410_cf98 = _409___mcc_h32;
          return (_module.D24.create_DC63((_pat_let_tv15).f9, _pat_let_tv16, (_pat_let_tv17).f4, new _dafny.CodePoint('s'.codePointAt(0)))).dtor_cf100;
        } else {
          let _411___mcc_h33 = (_source8).cf92;
          let _412_cf92 = _411___mcc_h33;
          return (_pat_let_tv18).multipliedBy(_pat_let_tv19);
        }
      }(_396_v173);
      let _413_v174;
      let _nw60 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _413_v174 = _nw60;
      let _index44 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_413_v174).length));
      (_413_v174)[_index44] = _dafny.Seq.UnicodeFromString("qvlmx");
      let _index45 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_413_v174).length));
      (_413_v174)[_index45] = _333_v148;
      _366_v154 = _366_v154;
      let _index46 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_363_v151).length));
      (_363_v151)[_index46] = _364_v152.f8;
      let _index47 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_363_v151).length));
      (_363_v151)[_index47] = _dafny.Seq.IsPrefixOf(_333_v148, _333_v148);
      let _414_v175;
      let _nw61 = Array((new BigNumber(7)).toNumber());
      _nw61[(_dafny.ZERO).toNumber()] = _331_v146;
      _nw61[(_dafny.ONE).toNumber()] = _331_v146;
      _nw61[(new BigNumber(2)).toNumber()] = _331_v146;
      _nw61[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_331_v146, _dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), ((_415_v152) => function (_416_i11) {
        return (_415_v152).f9;
      })(_364_v152)));
      _nw61[(new BigNumber(4)).toNumber()] = _331_v146;
      _nw61[(new BigNumber(5)).toNumber()] = _module.__default.fm19(_134_globalState);
      _nw61[(new BigNumber(6)).toNumber()] = _331_v146;
      _414_v175 = _nw61;
      _414_v175 = _414_v175;
      let _417_v176;
      _417_v176 = _dafny.Map.Empty.slice().updateUnsafe((_365_v153).f4,(_364_v152).f9);
      let _418_v177;
      _418_v177 = _dafny.MultiSet.fromElements((_417_v176).update(true, new BigNumber(160)));
      let _419_v178;
      let _nw62 = new _module.C2();
      _nw62.__ctor(!(_366_v154).isEqualTo((_332_v147)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_332_v147).length))]), _365_v153.f3, (_418_v177).IsProperSubsetOf(_418_v177), _365_v153.f3);
      _419_v178 = _nw62;
      let _420_v179;
      _420_v179 = new _dafny.CodePoint('a'.codePointAt(0));
      let _421_v180;
      let _out18;
      _out18 = (_364_v152).m2(_420_v179, _365_v153, _134_globalState);
      _421_v180 = _out18;
      process.stdout.write(_dafny.toString(_134_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_135_v0, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_136_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_329_v144));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_330_v145).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_331_v146, _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_332_v147)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_333_v148).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v149).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v149).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v149).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_334_v149).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v150).dtor_cf74));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v150).dtor_cf75));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v150).dtor_cf76));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_335_v150).dtor_cf77).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_335_v150).dtor_cf77).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_335_v150).dtor_cf77).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_335_v150).dtor_cf77).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v150).dtor_cf78));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_363_v151)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_364_v152.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v152).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_365_v153.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_365_v153).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_366_v154));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_391_v166).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_392_v169).dtor_cf55).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-1)),new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_394_v170).equals(_dafny.Set.fromElements(_module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-1)),new _dafny.CodePoint('l'.codePointAt(0))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_395_v171, _dafny.Seq.of(_dafny.Set.fromElements(_module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-1)),new _dafny.CodePoint('l'.codePointAt(0)))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_396_v173).dtor_cf92).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(_module.D17.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-1)),new _dafny.CodePoint('l'.codePointAt(0)))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_413_v174)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[_dafny.ONE], _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(-1), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965), new BigNumber(965)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(705)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_414_v175)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_417_v176).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(965)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_418_v177).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(160))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_419_v178).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_419_v178).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_420_v179));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_421_v180, _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))))));
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
    static create_DC2(cf3, cf4, cf5) {
      let $dt = new D0(1);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
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
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC4(cf7) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9, cf10, cf11) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D1(2);
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf6 === other.cf6;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO);
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
    static create_DC8(cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ", " + this.cf14.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8([], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC10(cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D3(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.MultiSet.Empty, _module.D1.Default());
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
    static create_DC12(cf19, cf20, cf21) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D4(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC13(cf22) {
      let $dt = new D4(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC14(cf23) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf24) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24);
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
    static create_DC17(cf26, cf27, cf28, cf29) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC16(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(_dafny.ZERO, false, false, _dafny.ZERO);
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
    static create_DC19(cf31, cf32) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC20(cf33, cf34) {
      let $dt = new D8(1);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC18(cf30) {
      let $dt = new D8(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC19(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC21(cf35) {
      let $dt = new D9(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf35) + ")";
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
    static create_DC23() {
      let $dt = new D10(0);
      return $dt;
    }
    static create_DC22(cf36) {
      let $dt = new D10(1);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC23";
      } else if (this.$tag === 1) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf36) + ")";
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
        return other.$tag === 1 && this.cf36 === other.cf36;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC23();
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
    static create_DC25(cf38) {
      let $dt = new D11(0);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC24(cf37) {
      let $dt = new D11(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(_dafny.Map.Empty);
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
    static create_DC27(cf40) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf39) {
      let $dt = new D12(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC28(cf41) {
      let $dt = new D12(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC27(_dafny.ZERO);
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
    static create_DC29(cf42) {
      let $dt = new D13(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42;
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf44, cf45) {
      let $dt = new D14(0);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC30(cf43) {
      let $dt = new D14(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC31(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC33(cf47, cf48) {
      let $dt = new D15(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC32(cf46) {
      let $dt = new D15(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC34(cf49) {
      let $dt = new D15(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf46 === other.cf46;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC33([], null);
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
    static create_DC36(cf51, cf52, cf53, cf54) {
      let $dt = new D16(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC35(cf50) {
      let $dt = new D16(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC36" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && this.cf53 === other.cf53 && this.cf54 === other.cf54;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC36(_dafny.Set.Empty, false, false, false);
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
    static create_DC38(cf56) {
      let $dt = new D17(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC39(cf57, cf58, cf59, cf60, cf61) {
      let $dt = new D17(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC37(cf55) {
      let $dt = new D17(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC40(cf62) {
      let $dt = new D17(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC40() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC38" + "(" + this.cf56.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC39" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC37" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf57 === other.cf57 && this.cf58 === other.cf58 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC38(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC42(cf64) {
      let $dt = new D18(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC43(cf65, cf66, cf67) {
      let $dt = new D18(1);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC41(cf63) {
      let $dt = new D18(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC42" + "(" + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC41" + "(" + _dafny.toString(this.cf63) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf65, other.cf65) && _dafny.areEqual(this.cf66, other.cf66) && this.cf67 === other.cf67;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC42(_dafny.ZERO);
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
    static create_DC45(cf69, cf70, cf71, cf72, cf73) {
      let $dt = new D19(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC46(cf74, cf75, cf76, cf77, cf78) {
      let $dt = new D19(1);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC44(cf68) {
      let $dt = new D19(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC47(cf79) {
      let $dt = new D19(3);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC44" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && this.cf72 === other.cf72 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf74, other.cf74) && this.cf75 === other.cf75 && this.cf76 === other.cf76 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC45(_dafny.ZERO, _dafny.ZERO, false, false, _dafny.Seq.of());
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
    static create_DC49() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC50(cf81, cf82) {
      let $dt = new D20(1);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC48(cf80) {
      let $dt = new D20(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC49";
      } else if (this.$tag === 1) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf80) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC49();
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
    static create_DC52(cf84, cf85) {
      let $dt = new D21(0);
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC53(cf86) {
      let $dt = new D21(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC51(cf83) {
      let $dt = new D21(2);
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC54(cf87) {
      let $dt = new D21(3);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84) && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf86 === other.cf86;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf83 === other.cf83;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf87, other.cf87);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC52(_dafny.ZERO, []);
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
    static create_DC56(cf89, cf90) {
      let $dt = new D22(0);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC55(cf88) {
      let $dt = new D22(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC57(cf91) {
      let $dt = new D22(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC55" + "(" + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC57" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf89, other.cf89) && this.cf90 === other.cf90;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf88 === other.cf88;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf91, other.cf91);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC56(_dafny.ZERO, false);
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
    static create_DC59(cf93, cf94, cf95) {
      let $dt = new D23(0);
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      return $dt;
    }
    static create_DC60(cf96, cf97) {
      let $dt = new D23(1);
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC61(cf98) {
      let $dt = new D23(2);
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC58(cf92) {
      let $dt = new D23(3);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get is_DC58() { return this.$tag === 3; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC59" + "(" + this.cf93.toVerbatimString(true) + ", " + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC60" + "(" + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC61" + "(" + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 3) {
        return "D23.DC58" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94) && _dafny.areEqual(this.cf95, other.cf95);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf96 === other.cf96 && this.cf97 === other.cf97;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf92, other.cf92);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC59(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC63(cf100, cf101, cf102, cf103) {
      let $dt = new D24(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC64(cf104, cf105, cf106) {
      let $dt = new D24(1);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      return $dt;
    }
    static create_DC62(cf99) {
      let $dt = new D24(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC65(cf107) {
      let $dt = new D24(3);
      $dt.cf107 = cf107;
      return $dt;
    }
    get is_DC63() { return this.$tag === 0; }
    get is_DC64() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get is_DC65() { return this.$tag === 3; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf107() { return this.cf107; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC63" + "(" + _dafny.toString(this.cf100) + ", " + this.cf101.toVerbatimString(true) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC64" + "(" + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC62" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC65" + "(" + _dafny.toString(this.cf107) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf100, other.cf100) && _dafny.areEqual(this.cf101, other.cf101) && this.cf102 === other.cf102 && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf104 === other.cf104 && _dafny.areEqual(this.cf105, other.cf105) && this.cf106 === other.cf106;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf99, other.cf99);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf107, other.cf107);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC63(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC66(cf108) {
      let $dt = new D25(0);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC66" + "(" + _dafny.toString(this.cf108) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf108, other.cf108);
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
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC67(cf109) {
      let $dt = new D26(0);
      $dt.cf109 = cf109;
      return $dt;
    }
    get is_DC67() { return this.$tag === 0; }
    get dtor_cf109() { return this.cf109; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC67" + "(" + _dafny.toString(this.cf109) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf109, other.cf109);
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
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC69(cf111) {
      let $dt = new D27(0);
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC68(cf110) {
      let $dt = new D27(1);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get is_DC68() { return this.$tag === 1; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC69" + "(" + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC68" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf111, other.cf111);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf110, other.cf110);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC69(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
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
      this.f0 = false;
      this._f1 = false;
      this._f2 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
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
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.UnicodeFromString("hmyjnxlg")).length);
    };
    fm11(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("nbgjfgxoq");
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = _dafny.Seq.UnicodeFromString("");
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if ((p0).isLessThan(p0)) {
        let _422_v0;
        let _nw63 = Array((new BigNumber(8)).toNumber()).fill(false);
        _422_v0 = _nw63;
        _422_v0 = _422_v0;
        if (p1.f3) {
          let _423_v1;
          _423_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,new BigNumber(((_this).f12).length));
          _423_v1 = (_423_v1).update(p1.f3, (p1).fm0(globalState));
          let _424_v3;
          _424_v3 = _dafny.Seq.of(new BigNumber(164));
          let _425_v4;
          _425_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _426_v5;
          _426_v5 = _dafny.MultiSet.fromElements(function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_424_v3).Elements) {
              let _427_v2 = _compr_22;
              if (_dafny.Seq.contains(_424_v3, _427_v2)) {
                _coll22.push([(_427_v2).multipliedBy(p0),p0]);
              }
            }
            return _coll22;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(443),p0), _425_v4);
          let _rhs42 = !(_module.__default.fm21(new BigNumber(647), _426_v5, globalState)) || (p1.f3);
          let _rhs43 = p0;
          let _rhs44 = p0;
          let _rhs45 = (p1).fm0(globalState);
          let _lhs34 = globalState;
          _lhs34.f0 = _rhs42;
          r0 = _rhs43;
          r0 = _rhs44;
          r0 = _rhs45;
          (globalState).f0 = (_module.__default.safeModuloInt(p0, p0)).isLessThan(p0);
          let _428_v6;
          let _nw64 = new _module.C0();
          _nw64.__ctor();
          _428_v6 = _nw64;
          let _429_v7;
          _429_v7 = _dafny.Map.Empty.slice().updateUnsafe(_425_v4,(p1).f4);
          let _430_v8;
          _430_v8 = _dafny.MultiSet.fromElements(p1.f3);
          _429_v7 = (_429_v7).update(_425_v4, (_430_v8).IsProperSubsetOf(_430_v8));
        } else {
          let _431_v9;
          _431_v9 = _module.D1.create_DC6();
          _431_v9 = _431_v9;
          let _432_v10;
          let _nw65 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _432_v10 = _nw65;
          let _index48 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_432_v10).length));
          (_432_v10)[_index48] = _dafny.MultiSet.fromElements((p1).f4);
          let _433_v11;
          _433_v11 = _dafny.MultiSet.fromElements(p1.f3);
          let _index49 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_432_v10).length));
          (_432_v10)[_index49] = (_433_v11).Difference(_dafny.MultiSet.fromElements((p1).f4, (p1).f4, (p1).f4, (p1).f4));
          let _434_v12;
          let _init5 = ((_435_p0) => function (_436_i0) {
            return _dafny.Set.fromElements(_435_p0);
          })(p0);
          let _nw66 = Array((new BigNumber(13)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw66.length); _i0_5++) {
            _nw66[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _434_v12 = _nw66;
          _434_v12 = _434_v12;
          let _437_v13;
          _437_v13 = _dafny.Set.fromElements((p1).f4, !(p1.f3));
          (p1).f3 = (_437_v13).IsDisjointFrom(_437_v13);
          (p1).f3 = (p1).f4;
        }
        r0 = p0;
        let _438_v14;
        let _init6 = function (_439_i1) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        };
        let _nw67 = Array((new BigNumber(22)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw67.length); _i0_6++) {
          _nw67[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _438_v14 = _nw67;
        let _440_v15;
        _440_v15 = new _dafny.CodePoint('f'.codePointAt(0));
        let _index50 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_438_v14).length));
        (_438_v14)[_index50] = _440_v15;
        let _index51 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_438_v14).length));
        (_438_v14)[_index51] = _440_v15;
        let _441_v16;
        _441_v16 = _dafny.Set.fromElements((p1).f4, (p1).f4);
        let _442_v17;
        _442_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_this).f12).length));
        let _443_v18;
        _443_v18 = _dafny.Seq.of(p0, p0, p0);
        let _444_v19;
        _444_v19 = _dafny.Seq.of(p0, new BigNumber((_441_v16).length), new BigNumber(660), (_dafny.ZERO).minus((new BigNumber(131)).minus((((_442_v17).contains(p0)) ? ((_442_v17).get(p0)) : (new BigNumber((_443_v18).length))))), p0);
        let _rhs46 = new BigNumber(293);
        let _rhs47 = _dafny.areEqual((_this).f12, (_this).f12);
        let _rhs48 = _dafny.Seq.of(p0, (p1).fm0(globalState), p0, new BigNumber(-118));
        let _lhs35 = globalState;
        r0 = _rhs46;
        _lhs35.f0 = _rhs47;
        _443_v18 = _rhs48;
      } else {
        if (p1.f3) {
          let _445_v20;
          _445_v20 = new _dafny.CodePoint('u'.codePointAt(0));
          let _446_v21;
          _446_v21 = _dafny.Set.fromElements((p1).f4, (_this).f13);
          let _447_v22;
          let _nw68 = Array((new BigNumber(9)).toNumber()).fill(false);
          _447_v22 = _nw68;
          let _index52 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_447_v22).length));
          (_447_v22)[_index52] = (p1).f4;
          let _448_v23;
          _448_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0);
          let _index53 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_447_v22).length));
          let _rhs49 = ((_this).f12)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_448_v23, (_448_v23).update((p1).f4, (_dafny.ZERO).minus(p0)))).length), new BigNumber(((_this).f12).length))];
          let _rhs50 = _dafny.Set.fromElements((p1).f4, p1.f3);
          let _rhs51 = (p0).isLessThanOrEqualTo((p0).minus(p0));
          let _rhs52 = _module.__default.fm21(new BigNumber(339), _module.__default.fm22(p0, p0, new BigNumber((_448_v23).length), globalState), globalState);
          let _lhs36 = _447_v22;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_447_v22).length));
          let _lhs38 = globalState;
          _445_v20 = _rhs49;
          _446_v21 = _rhs50;
          _lhs36[_lhs37] = _rhs51;
          _lhs38.f0 = _rhs52;
          let _449_v24;
          let _nw69 = Array((new BigNumber(25)).toNumber()).fill([]);
          _449_v24 = _nw69;
          let _450_v25;
          let _nw70 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _450_v25 = _nw70;
          let _index54 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_449_v24).length));
          (_449_v24)[_index54] = _450_v25;
          let _index55 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_449_v24).length));
          let _nw71 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
          (_449_v24)[_index55] = _nw71;
          let _451_v26;
          _451_v26 = _dafny.Seq.of(((p1.f3) ? ((p1).f4) : (true)));
          let _452_v27;
          _452_v27 = _dafny.Seq.UnicodeFromString("cw");
          let _453_v28;
          _453_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_445_v20);
          let _454_v29;
          _454_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_445_v20);
          let _rhs53 = _dafny.Seq.update(_dafny.Seq.of((p1).f4, (p1).f4, false, (_447_v22)[_module.__default.safeIndex(new BigNumber(930), new BigNumber((_447_v22).length))], _dafny.areEqual((_this).f12, _dafny.Seq.UnicodeFromString("ter"))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of((p1).f4, (p1).f4, false, (_447_v22)[_module.__default.safeIndex(new BigNumber(930), new BigNumber((_447_v22).length))], _dafny.areEqual((_this).f12, _dafny.Seq.UnicodeFromString("ter")))).length)), false);
          let _rhs54 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-167)), ((_455_v20) => function (_456_i2) {
            return _455_v20;
          })(_445_v20)), _dafny.Seq.UnicodeFromString("kjmrx")), (_this).f12);
          let _rhs55 = p0;
          let _rhs56 = (((_453_v28).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_459_v20) => function (_460_i3) {
            return _459_v20;
          })(_445_v20))).length))) ? ((_453_v28).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_457_v20) => function (_458_i3) {
            return _457_v20;
          })(_445_v20))).length))) : ((((_454_v29).contains((_this).f13)) ? ((_454_v29).get((_this).f13)) : (_module.__default.fm23((p1).f4, new BigNumber(2), globalState)))));
          let _rhs57 = _dafny.areEqual(_dafny.Seq.Concat((_this).f12, (_this).f12), _dafny.Seq.Concat(_452_v27, (_this).f12));
          let _lhs39 = p1;
          _451_v26 = _rhs53;
          _452_v27 = _rhs54;
          r0 = _rhs55;
          _445_v20 = _rhs56;
          _lhs39.f3 = _rhs57;
          let _461_v30;
          _461_v30 = _dafny.MultiSet.fromElements(p0, new BigNumber(134), new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_452_v27)).cardinality()), new BigNumber(382), new BigNumber((_448_v23).length)), _module.__default.safeIndex(new BigNumber((_452_v27).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_452_v27)).cardinality()), new BigNumber(382), new BigNumber((_448_v23).length))).length)), p0)).length));
          let _462_v31;
          _462_v31 = _dafny.Set.fromElements(new BigNumber((_461_v30).cardinality()), p0, p0, p0);
          _462_v31 = _462_v31;
          _452_v27 = (_this).f12;
        } else {
          let _463_v32;
          _463_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1).f4);
          let _464_v33;
          _464_v33 = _dafny.Set.fromElements(p1.f3);
          (p1).f3 = (_464_v33).contains((((_463_v32).contains(p0)) ? ((_463_v32).get(p0)) : ((p1).f4)));
          let _465_v34;
          _465_v34 = _module.D1.create_DC5(p1.f3, (p1).f4, _module.__default.safeModuloInt(p0, p0), p0);
          let _rhs58 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).isLessThan(p0);
          let _rhs59 = _465_v34;
          let _rhs60 = (p1).fm0(globalState);
          let _lhs40 = p1;
          _lhs40.f3 = _rhs58;
          _465_v34 = _rhs59;
          r0 = _rhs60;
          let _466_v35;
          _466_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _467_v37;
          let _nw72 = Array((new BigNumber(9)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _466_v35;
          _nw72[(_dafny.ONE).toNumber()] = (_466_v35).update(new BigNumber(-880), p0);
          _nw72[(new BigNumber(2)).toNumber()] = _466_v35;
          _nw72[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _nw72[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(256))).Merge(_466_v35);
          _nw72[(new BigNumber(5)).toNumber()] = _466_v35;
          _nw72[(new BigNumber(6)).toNumber()] = function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(360))) {
              let _468_v36 = _compr_23;
              if (((new BigNumber(958)).isLessThanOrEqualTo(_468_v36)) && ((_468_v36).isLessThan(new BigNumber(360)))) {
                _coll23.push([(_468_v36).plus(p0),p0]);
              }
            }
            return _coll23;
          }();
          _nw72[(new BigNumber(7)).toNumber()] = _466_v35;
          _nw72[(new BigNumber(8)).toNumber()] = _466_v35;
          _467_v37 = _nw72;
          let _rhs61 = (p1).f4;
          let _rhs62 = p0;
          let _rhs63 = _467_v37;
          let _rhs64 = new BigNumber(290);
          let _lhs41 = globalState;
          _lhs41.f0 = _rhs61;
          r0 = _rhs62;
          _467_v37 = _rhs63;
          r0 = _rhs64;
          (p1).f3 = (p1).f4;
          (p1).f3 = false;
        }
        let _469_v38;
        let _nw73 = new _module.C0();
        _nw73.__ctor();
        _469_v38 = _nw73;
        let _470_v39;
        _470_v39 = _dafny.MultiSet.fromElements(_469_v38);
        let _471_v40;
        _471_v40 = _dafny.Seq.of(_469_v38);
        let _472_v41;
        _472_v41 = _module.D2.create_DC7(_470_v39);
        let _473_v42;
        let _nw74 = Array((new BigNumber(10)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _470_v39;
        _nw74[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_471_v40)[_module.__default.safeIndex(p0, new BigNumber((_471_v40).length))], _469_v38, _469_v38, _469_v38, _469_v38);
        _nw74[(new BigNumber(2)).toNumber()] = _470_v39;
        _nw74[(new BigNumber(3)).toNumber()] = (_470_v39).Union(_dafny.MultiSet.fromElements(_469_v38));
        _nw74[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_469_v38));
        _nw74[(new BigNumber(5)).toNumber()] = (_470_v39).Difference((_dafny.MultiSet.fromElements(_469_v38)).update(_469_v38, _module.__default.abs(p0)));
        _nw74[(new BigNumber(6)).toNumber()] = (_470_v39).update(_469_v38, _module.__default.abs(p0));
        _nw74[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(((!(false)) ? (_471_v40) : (_471_v40)));
        _nw74[(new BigNumber(8)).toNumber()] = (_472_v41).dtor_cf12;
        _nw74[(new BigNumber(9)).toNumber()] = _470_v39;
        _473_v42 = _nw74;
        let _index56 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_473_v42).length));
        (_473_v42)[_index56] = _470_v39;
        let _index57 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_473_v42).length));
        (_473_v42)[_index57] = _470_v39;
        let _474_v43;
        _474_v43 = _dafny.Set.fromElements(p0, p0);
        let _475_v44;
        _475_v44 = _dafny.Map.Empty.slice().updateUnsafe(((!((p1).f4)) ? (!((_this).f13)) : ((p1).f4)),new BigNumber((_474_v43).length));
        _475_v44 = (_475_v44).update(false, (p0).minus(p0));
        r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        let _476_v45;
        _476_v45 = _dafny.Seq.of((_this).f12, _dafny.Seq.UnicodeFromString("nlvktxw"), (_this).f12, (_this).f12);
        let _477_v46;
        _477_v46 = new _dafny.CodePoint('o'.codePointAt(0));
        let _478_v47;
        _478_v47 = _dafny.Set.fromElements((_this).f13);
        let _479_v48;
        _479_v48 = _dafny.Seq.of((_476_v45)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(p1.f3, p1.f3, p1.f3)).length), new BigNumber((_476_v45).length))], _dafny.Seq.update(_dafny.Seq.update((_469_v38).fm11(_477_v46, _478_v47, globalState), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber(((_469_v38).fm11(_477_v46, _478_v47, globalState)).length)), ((_this).f12)[_module.__default.safeIndex(p0, new BigNumber(((_this).f12).length))]), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update((_469_v38).fm11(_477_v46, _478_v47, globalState), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber(((_469_v38).fm11(_477_v46, _478_v47, globalState)).length)), ((_this).f12)[_module.__default.safeIndex(p0, new BigNumber(((_this).f12).length))])).length)), _477_v46));
        let _480_v49;
        _480_v49 = _dafny.Seq.of(true);
        _479_v48 = _dafny.Seq.of(_dafny.Seq.Concat((_this).f12, (_this).f12), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("a"), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((p1).f4, (_this).f13, (p1).f4, (_480_v49)[_module.__default.safeIndex(new BigNumber(267), new BigNumber((_480_v49).length))], true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("a")).length)), _477_v46));
      }
      let _481_v51;
      _481_v51 = new _dafny.CodePoint('v'.codePointAt(0));
      let _482_v52;
      _482_v52 = _dafny.Set.fromElements(_481_v51, _481_v51, _481_v51, _481_v51, _481_v51);
      let _483_v53;
      _483_v53 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),_module.__default.fm24(globalState));
      if (!((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_482_v52).Elements) {
          let _484_v50 = _compr_24;
          if ((_482_v52).contains(_484_v50)) {
            _coll24.push([_484_v50,_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(((_this).f12).length), p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(p1.f3))).length), p0, p0))]);
          }
        }
        return _coll24;
      }()).Merge(_483_v53)).equals(_483_v53)) {
        let _485_v54;
        _485_v54 = _dafny.Set.fromElements(p1.f3);
        if ((new BigNumber((_485_v54).length)).isEqualTo(p0)) {
          let _486_v55;
          _486_v55 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nfsrcr"), (_this).f12, (_this).f12, (_this).f12, (_this).f12);
          r0 = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(new BigNumber((_486_v55).cardinality())));
          r0 = p0;
          let _487_v56;
          _487_v56 = _dafny.Seq.of(p0, p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_488_i4) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })).length));
          let _489_v57;
          _489_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-484),_487_v56);
          _489_v57 = (_489_v57).update(p0, _dafny.Seq.Concat(_487_v56, _dafny.Seq.of(p0)));
          let _490_v58;
          _490_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D0.create_DC2(false, p0, new BigNumber(607))).dtor_cf3);
          _490_v58 = (_490_v58).update(p0, ((p1).f4) === (p1.f3));
          let _491_v59;
          let _nw75 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
          _491_v59 = _nw75;
          let _index58 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_491_v59).length));
          (_491_v59)[_index58] = _487_v56;
          let _index59 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_491_v59).length));
          (_491_v59)[_index59] = _487_v56;
        } else {
          _481_v51 = _481_v51;
          let _492_v60;
          _492_v60 = _dafny.MultiSet.fromElements(p0, p0);
          _492_v60 = _492_v60;
          let _493_v61;
          let _init7 = function (_494_i5) {
            return true;
          };
          let _nw76 = Array((new BigNumber(2)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw76.length); _i0_7++) {
            _nw76[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _493_v61 = _nw76;
          let _index60 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_493_v61).length));
          (_493_v61)[_index60] = (p1).f4;
          let _index61 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_493_v61).length));
          (_493_v61)[_index61] = !(((!(true)) ? ((_dafny.ZERO).minus(p0)) : (p0))).isEqualTo(p0);
          (p1).f3 = p1.f3;
          let _495_v62;
          _495_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _496_v63;
          _496_v63 = _dafny.MultiSet.fromElements(_495_v62);
          (globalState).f0 = _module.__default.fm21(p0, _496_v63, globalState);
        }
        let _497_v64;
        _497_v64 = _dafny.Map.Empty.slice().updateUnsafe(_481_v51,(p1).f4);
        r0 = new BigNumber((_497_v64).length);
        r0 = p0;
        let _498_v65;
        let _nw77 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _498_v65 = _nw77;
        let _index62 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_498_v65).length));
        (_498_v65)[_index62] = p0;
        let _index63 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_498_v65).length));
        let _rhs65 = p0;
        let _rhs66 = p0;
        let _lhs42 = _498_v65;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_498_v65).length));
        r0 = _rhs65;
        _lhs42[_lhs43] = _rhs66;
        let _499_v66;
        let _nw78 = new _module.C0();
        _nw78.__ctor();
        _499_v66 = _nw78;
      } else {
        let _500_v67;
        let _nw79 = Array((new BigNumber(4)).toNumber()).fill(false);
        _500_v67 = _nw79;
        let _501_v68;
        _501_v68 = _dafny.Seq.of(_500_v67, _500_v67, _500_v67);
        r0 = new BigNumber((_501_v68).length);
        (globalState).f0 = (p1).f4;
        let _502_v69;
        _502_v69 = _dafny.Seq.of(p1.f3);
        let _503_v70;
        _503_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_502_v69).length),p1.f3);
        let _504_v71;
        _504_v71 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _505_v72;
        _505_v72 = _dafny.MultiSet.fromElements(_504_v71, _504_v71);
        let _506_v73;
        _506_v73 = _dafny.Set.fromElements(p1.f3, false, _module.__default.fm21(new BigNumber((_503_v70).length), _505_v72, globalState));
        (globalState).f0 = (_506_v73).IsProperSubsetOf(_506_v73);
        r0 = p0;
        r0 = p0;
      }
      let _507_v74;
      let _nw80 = Array((new BigNumber(14)).toNumber()).fill(false);
      _507_v74 = _nw80;
      _507_v74 = _507_v74;
      let _508_v75;
      _508_v75 = _dafny.Seq.of(p1.f3);
      let _509_v76;
      _509_v76 = _dafny.Set.fromElements((p1).f4);
      let _510_v77;
      _510_v77 = _dafny.MultiSet.fromElements(p1.f3, (p1).f4, (_this).f13);
      let _511_v78;
      _511_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, p1.f3)).length),p0);
      let _512_v79;
      _512_v79 = _dafny.MultiSet.fromElements(p0);
      let _513_v80;
      _513_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_module.D0.create_DC0((((_512_v79).contains(p0)) ? ((_512_v79).get(p0)) : (p0)))).dtor_cf0);
      let _514_v81;
      _514_v81 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,new BigNumber((_513_v80).length));
      let _515_v82;
      _515_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
      let _516_v83;
      _516_v83 = _dafny.Seq.of(new BigNumber((_515_v82).length), p0);
      let _517_v84;
      let _nw81 = Array((new BigNumber(24)).toNumber());
      _nw81[(_dafny.ZERO).toNumber()] = p0;
      _nw81[(_dafny.ONE).toNumber()] = (((_508_v75)[_module.__default.safeIndex(p0, new BigNumber((_508_v75).length))]) ? (p0) : (new BigNumber((_509_v76).length)));
      _nw81[(new BigNumber(2)).toNumber()] = p0;
      _nw81[(new BigNumber(3)).toNumber()] = p0;
      _nw81[(new BigNumber(4)).toNumber()] = p0;
      _nw81[(new BigNumber(5)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("myku")).length)).minus(p0);
      _nw81[(new BigNumber(6)).toNumber()] = (((_510_v77).contains((p1).f4)) ? ((_510_v77).get((p1).f4)) : (p0));
      _nw81[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((((_511_v78).contains(p0)) ? ((_511_v78).get(p0)) : (p0)), p0);
      _nw81[(new BigNumber(8)).toNumber()] = p0;
      _nw81[(new BigNumber(9)).toNumber()] = (((p1).f4) ? (p0) : (p0));
      _nw81[(new BigNumber(10)).toNumber()] = p0;
      _nw81[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((p0).multipliedBy(p0));
      _nw81[(new BigNumber(12)).toNumber()] = new BigNumber(373);
      _nw81[(new BigNumber(13)).toNumber()] = new BigNumber(974);
      _nw81[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qrv")).length);
      _nw81[(new BigNumber(15)).toNumber()] = (p0).minus(p0);
      _nw81[(new BigNumber(16)).toNumber()] = new BigNumber(((_this).f12).length);
      _nw81[(new BigNumber(17)).toNumber()] = p0;
      _nw81[(new BigNumber(18)).toNumber()] = _dafny.ONE;
      _nw81[(new BigNumber(19)).toNumber()] = (((_513_v80).contains(p1.f3)) ? ((_513_v80).get(p1.f3)) : (p0));
      _nw81[(new BigNumber(20)).toNumber()] = new BigNumber(-114);
      _nw81[(new BigNumber(21)).toNumber()] = new BigNumber((_module.__default.fm25(globalState)).length);
      _nw81[(new BigNumber(22)).toNumber()] = p0;
      _nw81[(new BigNumber(23)).toNumber()] = new BigNumber((_516_v83).length);
      _517_v84 = _nw81;
      let _index64 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
      (_517_v84)[_index64] = new BigNumber((_dafny.Seq.Concat((_this).f12, _dafny.Seq.UnicodeFromString("ywt"))).length);
      let _518_v85;
      _518_v85 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat((_this).f12, _dafny.Seq.UnicodeFromString("x")));
      let _index65 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
      (_517_v84)[_index65] = (((_518_v85).contains((_this).f12)) ? ((_518_v85).get((_this).f12)) : (p0));
      if ((p1).f4) {
        let _519_v86;
        _519_v86 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_509_v76);
        _519_v86 = ((_519_v86).Merge((_519_v86).update((p1).f4, _509_v76))).update((p1).f4, _509_v76);
        let _520_v87;
        _520_v87 = _dafny.Seq.of(_517_v84, _517_v84);
        let _rhs67 = (_520_v87)[_module.__default.safeIndex(p0, new BigNumber((_520_v87).length))];
        let _rhs68 = (p1).f4;
        let _lhs44 = globalState;
        _517_v84 = _rhs67;
        _lhs44.f0 = _rhs68;
        let _521_v88;
        _521_v88 = _dafny.Set.fromElements(p0);
        let _index66 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_507_v74).length));
        (_507_v74)[_index66] = (_521_v88).IsSubsetOf(_521_v88);
        let _index67 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_507_v74).length));
        (_507_v74)[_index67] = (_this).f13;
        let _522_v89;
        _522_v89 = _dafny.Seq.UnicodeFromString("kttjwxwja");
        _522_v89 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), ((_523_v51) => function (_524_i6) {
          return (_module.D0.create_DC1(false, _523_v51)).dtor_cf2;
        })(_481_v51)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), ((_525_v51) => function (_526_i6) {
          return (_module.D0.create_DC1(false, _525_v51)).dtor_cf2;
        })(_481_v51))).length)), _module.__default.fm23(p1.f3, (_dafny.ZERO).minus(p0), globalState));
        r0 = (_dafny.ZERO).minus((new BigNumber((_509_v76).length)).minus(_module.__default.safeModuloInt((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))], new BigNumber((_522_v89).length))));
      } else {
        let _527_v90;
        _527_v90 = _dafny.Seq.UnicodeFromString("lvigjoi");
        _527_v90 = _dafny.Seq.Concat(_dafny.Seq.Concat(_527_v90, _dafny.Seq.update(_527_v90, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_527_v90).length)), _481_v51)), _dafny.Seq.Concat(_527_v90, (_this).f12));
        let _528_v91;
        let _nw82 = new _module.C0();
        _nw82.__ctor();
        _528_v91 = _nw82;
        let _529_v92;
        _529_v92 = _module.D1.create_DC6();
        let _source9 = _529_v92;
        if (_source9.is_DC4) {
          let _530___mcc_h0 = (_source9).cf7;
          let _531_cf7 = _530___mcc_h0;
          r0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_531_cf7, (_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))]));
          _531_cf7 = _531_cf7;
          let _532_v93;
          let _nw83 = Array((new BigNumber(14)).toNumber()).fill(_module.D1.Default());
          _532_v93 = _nw83;
          let _index68 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_532_v93).length));
          (_532_v93)[_index68] = _529_v92;
          let _index69 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_532_v93).length));
          (_532_v93)[_index69] = _529_v92;
          (globalState).f0 = false;
        } else if (_source9.is_DC5) {
          let _533___mcc_h1 = (_source9).cf8;
          let _534___mcc_h2 = (_source9).cf9;
          let _535___mcc_h3 = (_source9).cf10;
          let _536___mcc_h4 = (_source9).cf11;
          let _537_cf11 = _536___mcc_h4;
          let _538_cf10 = _535___mcc_h3;
          let _539_cf9 = _534___mcc_h2;
          let _540_cf8 = _533___mcc_h1;
          (globalState).f0 = _539_cf9;
          r0 = (p0).multipliedBy(_module.__default.safeDivisionInt((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))], (_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))]));
          let _541_v94;
          let _init8 = ((_542_v90) => function (_543_i7) {
            return _dafny.Set.fromElements(new BigNumber((_542_v90).length));
          })(_527_v90);
          let _nw84 = Array((new BigNumber(24)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw84.length); _i0_8++) {
            _nw84[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _541_v94 = _nw84;
          let _544_v95;
          _544_v95 = _dafny.Set.fromElements((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))]);
          let _index70 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_541_v94).length));
          (_541_v94)[_index70] = (_544_v95).Union(_544_v95);
          let _545_v96;
          _545_v96 = _dafny.MultiSet.fromElements(_528_v91);
          let _546_v97;
          _546_v97 = _module.D2.create_DC7(_545_v96);
          let _547_v98;
          _547_v98 = _dafny.Map.Empty.slice().updateUnsafe(_546_v97,(p1).fm0(globalState));
          let _index71 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
          let _index72 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_541_v94).length));
          let _index73 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
          let _rhs69 = (_528_v91).fm10(_537_cf11, (p1).f4, _512_v79, globalState);
          let _rhs70 = _544_v95;
          let _rhs71 = (((true) ? (new BigNumber((_512_v79).cardinality())) : ((((_547_v98).contains(_546_v97)) ? ((_547_v98).get(_546_v97)) : ((_dafny.ZERO).minus(new BigNumber((_512_v79).cardinality()))))))).plus(_538_cf10);
          let _rhs72 = p1.f3;
          let _lhs45 = _517_v84;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
          let _lhs47 = _541_v94;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_541_v94).length));
          let _lhs49 = _517_v84;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
          _lhs45[_lhs46] = _rhs69;
          _lhs47[_lhs48] = _rhs70;
          _lhs49[_lhs50] = _rhs71;
          _539_cf9 = _rhs72;
          _537_cf11 = _537_cf11;
        } else if (_source9.is_DC6) {
          (globalState).f0 = p1.f3;
          let _548_v99;
          _548_v99 = _dafny.Set.fromElements((_516_v83)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("qyylnovqs")).length), new BigNumber((_516_v83).length))]);
          _513_v80 = (_513_v80).update((p1).f4, new BigNumber(((_548_v99).Intersect(_dafny.Set.fromElements(new BigNumber((_509_v76).length)))).length));
          let _index74 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
          (_517_v84)[_index74] = p0;
          let _549_v100;
          _549_v100 = _dafny.Map.Empty.slice().updateUnsafe(_508_v75,_511_v78);
          let _550_v101;
          _550_v101 = _dafny.Map.Empty.slice().updateUnsafe(_549_v100,p0);
          _550_v101 = (_550_v101).update(_549_v100, (_dafny.ZERO).minus(p0));
        } else {
          let _551___mcc_h5 = (_source9).cf6;
          let _552_cf6 = _551___mcc_h5;
          let _index75 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_507_v74).length));
          (_507_v74)[_index75] = (_this).f13;
          let _index76 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_507_v74).length));
          (_507_v74)[_index76] = ((((true) ? (p1.f3) : (p1.f3))) ? ((p1).f4) : (!(p1.f3)));
          let _553_v102;
          let _nw85 = new _module.C0();
          _nw85.__ctor();
          _553_v102 = _nw85;
          let _554_v103;
          _554_v103 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1).f4);
          _514_v81 = (_514_v81).update((((_554_v103).contains((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))])) ? ((_554_v103).get((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))])) : (!((p1).f4))), new BigNumber((_dafny.Seq.of((p1).f4)).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_507_v74).length));
          (_507_v74)[_index77] = ((!(p1.f3) || (p1.f3)) ? (p1.f3) : (!(((p1).f4) || (p1.f3))));
        }
        let _index78 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length));
        (_517_v84)[_index78] = (((((_511_v78).contains((p1).fm0(globalState))) ? ((_511_v78).get((p1).fm0(globalState))) : (p0))).plus(p0)).plus((_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))]);
        let _555_v104;
        let _nw86 = new _module.C0();
        _nw86.__ctor();
        _555_v104 = _nw86;
      }
      let _556_v105;
      _556_v105 = _module.D0.create_DC2((p1).f4, new BigNumber(((_this).f12).length), p0);
      let _pat_let_tv20 = p1;
      let _pat_let_tv21 = p1;
      let _pat_let_tv22 = p0;
      let _pat_let_tv23 = p0;
      let _pat_let_tv24 = p1;
      let _pat_let_tv25 = _482_v52;
      let _pat_let_tv26 = p1;
      let _pat_let_tv27 = p1;
      let _pat_let_tv28 = _517_v84;
      let _pat_let_tv29 = _517_v84;
      let _pat_let_tv30 = _511_v78;
      (globalState).f0 = function (_source10) {
        if (_source10.is_DC1) {
          let _557___mcc_h6 = (_source10).cf1;
          let _558___mcc_h7 = (_source10).cf2;
          let _559_cf2 = _558___mcc_h7;
          let _560_cf1 = _557___mcc_h6;
          return (_dafny.MultiSet.fromElements(function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(377), new BigNumber(23))) {
              let _561_v106 = _compr_25;
              if (((new BigNumber(377)).isLessThanOrEqualTo(_561_v106)) && ((_561_v106).isLessThan(new BigNumber(23)))) {
                _coll25.push([_module.__default.safeDivisionInt(_561_v106, new BigNumber(441)),_pat_let_tv20.f3]);
              }
            }
            return _coll25;
          }(), function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(-707), new BigNumber(92))) {
              let _562_v107 = _compr_26;
              if (((new BigNumber(-707)).isLessThanOrEqualTo(_562_v107)) && ((_562_v107).isLessThan(new BigNumber(92)))) {
                _coll26.push([_module.__default.safeModuloInt(_562_v107, _pat_let_tv22),(_pat_let_tv21).f4]);
              }
            }
            return _coll26;
          }())).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv23,_pat_let_tv24.f3), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_pat_let_tv25).length),(_pat_let_tv26).f4)));
        } else if (_source10.is_DC2) {
          let _563___mcc_h8 = (_source10).cf3;
          let _564___mcc_h9 = (_source10).cf4;
          let _565___mcc_h10 = (_source10).cf5;
          let _566_cf5 = _565___mcc_h10;
          let _567_cf4 = _564___mcc_h9;
          let _568_cf3 = _563___mcc_h8;
          return !(((_pat_let_tv27).f4) || (true));
        } else {
          let _569___mcc_h11 = (_source10).cf0;
          let _570_cf0 = _569___mcc_h11;
          return (_dafny.Set.fromElements((_dafny.ZERO).minus((_pat_let_tv29)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_pat_let_tv28).length))]))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_pat_let_tv30).length)));
        }
      }(_556_v105);
      r0 = (_517_v84)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_517_v84).length))];
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      if (!((_this).f13)) {
        let _571_v0;
        let _init9 = ((_572_p1) => function (_573_i0) {
          return _572_p1.f3;
        })(p1);
        let _nw87 = Array((new BigNumber(22)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw87.length); _i0_9++) {
          _nw87[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _571_v0 = _nw87;
        let _574_v1;
        _574_v1 = _dafny.MultiSet.fromElements((p1).fm0(globalState));
        let _575_v2;
        _575_v2 = _dafny.Map.Empty.slice().updateUnsafe(_574_v1,p1.f3);
        let _index79 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length));
        (_571_v0)[_index79] = (((_575_v2).contains(_574_v1)) ? ((_575_v2).get(_574_v1)) : ((_this).f13));
        let _576_v3;
        _576_v3 = _dafny.Seq.of(!((p1).f4), true, p1.f3, (_this).f13);
        let _577_v4;
        _577_v4 = _dafny.Seq.of(_576_v3, _576_v3);
        let _578_v5;
        _578_v5 = new BigNumber(252);
        let _579_v6;
        _579_v6 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_578_v5);
        let _580_v7;
        _580_v7 = _dafny.Map.Empty.slice().updateUnsafe(false,p1.f3);
        let _index80 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length));
        let _rhs73 = !(_module.__default.fm26((((_579_v6).contains(p1.f3)) ? ((_579_v6).get(p1.f3)) : (new BigNumber(434))), p0, false, globalState)).equals(_580_v7);
        let _rhs74 = !(_578_v5).isEqualTo((p1).fm0(globalState));
        let _rhs75 = p1.f3;
        let _rhs76 = _577_v4;
        let _lhs51 = _571_v0;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length));
        let _lhs53 = p1;
        let _lhs54 = globalState;
        _lhs51[_lhs52] = _rhs73;
        _lhs53.f3 = _rhs74;
        _lhs54.f0 = _rhs75;
        _577_v4 = _rhs76;
        let _581_v8;
        let _nw88 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _581_v8 = _nw88;
        _581_v8 = ((p1.f3) ? (_581_v8) : (_581_v8));
        _581_v8 = _581_v8;
        let _582_v9;
        let _nw89 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
        _582_v9 = _nw89;
        let _583_v10;
        _583_v10 = _dafny.Set.fromElements(new BigNumber(-462), (((_574_v1).contains(_578_v5)) ? ((_574_v1).get(_578_v5)) : (_578_v5)));
        let _index81 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_582_v9).length));
        (_582_v9)[_index81] = _583_v10;
        let _584_v11;
        _584_v11 = _dafny.Set.fromElements(p0);
        let _index82 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_582_v9).length));
        let _rhs77 = _578_v5;
        let _rhs78 = _583_v10;
        let _rhs79 = _module.__default.fm27((((_571_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length))]) ? (new BigNumber((_579_v6).length)) : ((p1).fm0(globalState))), p1.f3, _578_v5, globalState);
        let _lhs55 = _582_v9;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_582_v9).length));
        _578_v5 = _rhs77;
        _lhs55[_lhs56] = _rhs78;
        _584_v11 = _rhs79;
        let _585_v12;
        _585_v12 = _module.D0.create_DC2((_571_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length))], new BigNumber((_579_v6).length), _578_v5);
        let _rhs80 = ((((_571_v0)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_571_v0).length))]) ? (_585_v12) : (_585_v12))).dtor_cf5;
        let _rhs81 = (p1).f4;
        let _lhs57 = globalState;
        _578_v5 = _rhs80;
        _lhs57.f0 = _rhs81;
      } else {
        let _586_v13;
        _586_v13 = new BigNumber(500);
        let _587_v14;
        _587_v14 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_586_v13,_586_v13));
        let _588_v15;
        _588_v15 = _dafny.MultiSet.fromElements(_586_v13, new BigNumber(-52));
        let _589_v16;
        _589_v16 = _dafny.MultiSet.fromElements(false, !((p1).f4));
        let _590_v17;
        _590_v17 = _module.D1.create_DC5(p1.f3, !(p1.f3), _586_v13, _586_v13);
        let _591_v18;
        _591_v18 = _dafny.Set.fromElements(true, p1.f3);
        let _592_v19;
        let _nw90 = Array((new BigNumber(28)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = p1.f3;
        _nw90[(_dafny.ONE).toNumber()] = _module.__default.fm21(_586_v13, _587_v14, globalState);
        _nw90[(new BigNumber(2)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(3)).toNumber()] = true;
        _nw90[(new BigNumber(4)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(5)).toNumber()] = (_this).f13;
        _nw90[(new BigNumber(6)).toNumber()] = (p1).f4;
        _nw90[(new BigNumber(7)).toNumber()] = (_this).f13;
        _nw90[(new BigNumber(8)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(9)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(10)).toNumber()] = (p1).f4;
        _nw90[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("fhyqkwo"), _module.__default.safeIndex(_586_v13, new BigNumber((_dafny.Seq.UnicodeFromString("fhyqkwo")).length)), p0), (_this).f12);
        _nw90[(new BigNumber(12)).toNumber()] = false;
        _nw90[(new BigNumber(13)).toNumber()] = _module.__default.fm21(_586_v13, _587_v14, globalState);
        _nw90[(new BigNumber(14)).toNumber()] = (_588_v15).IsDisjointFrom(_588_v15);
        _nw90[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements((p1).f4, !((p1).f4), (p1).f4)).IsProperSubsetOf(_589_v16);
        _nw90[(new BigNumber(16)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(17)).toNumber()] = (_590_v17).dtor_cf8;
        _nw90[(new BigNumber(18)).toNumber()] = true;
        _nw90[(new BigNumber(19)).toNumber()] = (p1).f4;
        _nw90[(new BigNumber(20)).toNumber()] = ((p1).f4) === (p1.f3);
        _nw90[(new BigNumber(21)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(22)).toNumber()] = !((p1).f4);
        _nw90[(new BigNumber(23)).toNumber()] = !(_591_v18).equals(_591_v18);
        _nw90[(new BigNumber(24)).toNumber()] = p1.f3;
        _nw90[(new BigNumber(25)).toNumber()] = (_this).f13;
        _nw90[(new BigNumber(26)).toNumber()] = (_this).f13;
        _nw90[(new BigNumber(27)).toNumber()] = p1.f3;
        _592_v19 = _nw90;
        _592_v19 = _592_v19;
        _586_v13 = _module.__default.safeModuloInt(new BigNumber(70), new BigNumber(-157));
        let _593_v20;
        _593_v20 = _dafny.Map.Empty.slice().updateUnsafe(_586_v13,(p1).f4);
        let _594_v21;
        _594_v21 = _dafny.MultiSet.fromElements(_593_v20);
        let _595_v22;
        _595_v22 = _module.D3.create_DC9(_594_v21);
        _586_v13 = new BigNumber(((_595_v22).dtor_cf15).cardinality());
        let _596_v23;
        _596_v23 = _dafny.Seq.of((p1).f4, _module.__default.fm21(_586_v13, _587_v14, globalState), !((p1).f4));
        let _597_v24;
        _597_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_589_v16).cardinality()),_596_v23);
        let _598_v25;
        _598_v25 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length));
        _597_v24 = (_597_v24).update((_598_v25)[_module.__default.safeIndex(_586_v13, new BigNumber((_598_v25).length))], _596_v23);
        r0 = _dafny.Seq.Concat((_this).f12, (_this).f12);
      }
      let _599_v26;
      _599_v26 = new BigNumber(859);
      let _600_v27;
      _600_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm28((p1).f4, (_this).f13, p1.f3, _599_v26, globalState),true);
      let _601_v28;
      _601_v28 = _dafny.Set.fromElements(new BigNumber(-226));
      let _602_v29;
      _602_v29 = _dafny.MultiSet.fromElements(new BigNumber(684), _599_v26);
      let _603_v30;
      _603_v30 = _dafny.MultiSet.fromElements(p0, _module.__default.fm23(false, new BigNumber(179), globalState), p0, p0, p0);
      let _604_v31;
      _604_v31 = _dafny.Seq.of(p1.f3);
      let _605_v32;
      let _nw91 = Array((new BigNumber(14)).toNumber());
      _nw91[(_dafny.ZERO).toNumber()] = (_599_v26).minus(new BigNumber((_600_v27).length));
      _nw91[(_dafny.ONE).toNumber()] = _599_v26;
      _nw91[(new BigNumber(2)).toNumber()] = _599_v26;
      _nw91[(new BigNumber(3)).toNumber()] = new BigNumber(477);
      _nw91[(new BigNumber(4)).toNumber()] = _599_v26;
      _nw91[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_599_v26).plus(new BigNumber((_601_v28).length)));
      _nw91[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), function (_606_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length);
      _nw91[(new BigNumber(7)).toNumber()] = (new BigNumber(((_this).f12).length)).minus(_599_v26);
      _nw91[(new BigNumber(8)).toNumber()] = ((((_602_v29).contains(_599_v26)) ? ((_602_v29).get(_599_v26)) : (new BigNumber(((_this).f12).length)))).multipliedBy(_599_v26);
      _nw91[(new BigNumber(9)).toNumber()] = (((_603_v30).contains(p0)) ? ((_603_v30).get(p0)) : (_599_v26));
      _nw91[(new BigNumber(10)).toNumber()] = new BigNumber((_602_v29).cardinality());
      _nw91[(new BigNumber(11)).toNumber()] = new BigNumber(972);
      _nw91[(new BigNumber(12)).toNumber()] = _599_v26;
      _nw91[(new BigNumber(13)).toNumber()] = new BigNumber((_604_v31).length);
      _605_v32 = _nw91;
      let _607_v33;
      _607_v33 = _module.D4.create_DC11(_dafny.Seq.of((p1).f4));
      let _index83 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length));
      (_605_v32)[_index83] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_604_v31, (_607_v33).dtor_cf18)).length));
      let _index84 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length));
      (_605_v32)[_index84] = (p1).fm0(globalState);
      let _608_v34;
      _608_v34 = _dafny.MultiSet.fromElements(_604_v31, _604_v31);
      let _hi3 = _module.__default.safeModuloInt(new BigNumber(((_this).f12).length), (((_608_v34).contains(_604_v31)) ? ((_608_v34).get(_604_v31)) : ((_dafny.ZERO).minus((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]))));
      for (let _609_i2 = new BigNumber(472); _609_i2.isLessThan(_hi3); _609_i2 = _609_i2.plus(_dafny.ONE)) {
        (globalState).f0 = (p1).f4;
        let _610_v35;
        _610_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(866),_609_i2);
        let _611_v36;
        _611_v36 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]),(_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]), _610_v35, _610_v35);
        let _612_v37;
        _612_v37 = _dafny.MultiSet.fromElements(p1.f3, _module.__default.fm21(_609_i2, _611_v36, globalState), (p1).f4);
        let _613_v38;
        _613_v38 = _dafny.Seq.of(_612_v37);
        let _614_v39;
        let _nw92 = Array((new BigNumber(26)).toNumber());
        _nw92[(_dafny.ZERO).toNumber()] = _612_v37;
        _nw92[(_dafny.ONE).toNumber()] = _612_v37;
        _nw92[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(p1.f3);
        _nw92[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements((p1).f4, p1.f3, p1.f3, (_this).f13, p1.f3);
        _nw92[(new BigNumber(4)).toNumber()] = (_612_v37).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(p1.f3, p1.f3, p1.f3, p1.f3)));
        _nw92[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(p1.f3, false);
        _nw92[(new BigNumber(6)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(7)).toNumber()] = (_613_v38)[_module.__default.safeIndex((_dafny.ZERO).minus((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]), new BigNumber((_613_v38).length))];
        _nw92[(new BigNumber(8)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(9)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(10)).toNumber()] = (_612_v37).Difference(_612_v37);
        _nw92[(new BigNumber(11)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(12)).toNumber()] = (_612_v37).Union(_612_v37);
        _nw92[(new BigNumber(13)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(14)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(p1.f3, (p1).f4)).Intersect(_612_v37);
        _nw92[(new BigNumber(16)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(17)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(18)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(19)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements((p1).f4, (p1).f4, (p1).f4, p1.f3);
        _nw92[(new BigNumber(21)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(22)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(23)).toNumber()] = _612_v37;
        _nw92[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.fromElements((p1).f4);
        _nw92[(new BigNumber(25)).toNumber()] = _612_v37;
        _614_v39 = _nw92;
        let _615_v40;
        _615_v40 = _dafny.Seq.of(_614_v39);
        let _616_v41;
        let _init10 = ((_617_p1) => function (_618_i3) {
          return (_617_p1).f4;
        })(p1);
        let _nw93 = Array((new BigNumber(17)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw93.length); _i0_10++) {
          _nw93[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _616_v41 = _nw93;
        let _619_v42;
        _619_v42 = _dafny.Set.fromElements(_616_v41, _616_v41);
        let _620_v43;
        _620_v43 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_619_v42);
        let _621_v44;
        _621_v44 = _dafny.Seq.of(_619_v42, _619_v42, (((_620_v43).contains(false)) ? ((_620_v43).get(false)) : (_619_v42)));
        _614_v39 = (_615_v40)[_module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber(((_621_v44)[_module.__default.safeIndex(_609_i2, new BigNumber((_621_v44).length))]).length)).minus(_599_v26)), new BigNumber((_615_v40).length))];
        (globalState).f0 = !(!(((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]).isEqualTo((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))])));
        let _index85 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length));
        (_605_v32)[_index85] = ((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]).plus((p1).fm0(globalState));
      }
      let _622_v45;
      _622_v45 = _dafny.Seq.of((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]);
      let _623_v46;
      _623_v46 = _dafny.Seq.of(_622_v45, _622_v45);
      let _624_v47;
      _624_v47 = _dafny.Seq.of(_623_v46, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-127)), ((_625_v45) => function (_626_i7) {
        return _625_v45;
      })(_622_v45)), _623_v46, _623_v46, _623_v46);
      let _627_v48;
      let _nw94 = Array((new BigNumber(21)).toNumber());
      _nw94[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), ((_628_p1) => function (_629_i4) {
        return _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_628_p1.f3, _628_p1.f3)).cardinality()));
      })(p1));
      _nw94[(_dafny.ONE).toNumber()] = _623_v46;
      _nw94[(new BigNumber(2)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(3)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(856)), ((_630_v45) => function (_631_i5) {
        return _630_v45;
      })(_622_v45)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), ((_632_v45) => function (_633_i6) {
        return _632_v45;
      })(_622_v45)));
      _nw94[(new BigNumber(5)).toNumber()] = _module.__default.fm29(_599_v26, globalState);
      _nw94[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_623_v46, (_624_v47)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_624_v47).length))]), _module.__default.safeIndex((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))], new BigNumber((_dafny.Seq.Concat(_623_v46, (_624_v47)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_624_v47).length))])).length)), _622_v45);
      _nw94[(new BigNumber(7)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-488)), ((_634_v45) => function (_635_i8) {
        return _634_v45;
      })(_622_v45));
      _nw94[(new BigNumber(9)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(10)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(11)).toNumber()] = (((p1).f4) ? (_623_v46) : (_module.__default.fm29((_dafny.ZERO).minus(_599_v26), globalState)));
      _nw94[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_dafny.Seq.of(_599_v26));
      _nw94[(new BigNumber(13)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(14)).toNumber()] = (((p1).f4) ? (_623_v46) : (_623_v46));
      _nw94[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_623_v46, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p1.f3)).length)), new BigNumber((_623_v46).length)), _622_v45), _dafny.Seq.update(_dafny.Seq.of(_dafny.Seq.of(_599_v26, (_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]), _622_v45, _dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_636_i9) {
        return new BigNumber(282);
      })), _module.__default.safeIndex(new BigNumber(((_this).f12).length), new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_599_v26, (_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]), _622_v45, _dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_637_i9) {
        return new BigNumber(282);
      }))).length)), _622_v45));
      _nw94[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_623_v46, _dafny.Seq.update(_623_v46, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update((_this).f12, _module.__default.safeIndex(_599_v26, new BigNumber(((_this).f12).length)), _module.__default.fm23(p1.f3, (_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))], globalState))).length)), new BigNumber((_623_v46).length)), _622_v45));
      _nw94[(new BigNumber(17)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(18)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(19)).toNumber()] = _623_v46;
      _nw94[(new BigNumber(20)).toNumber()] = _623_v46;
      _627_v48 = _nw94;
      let _index86 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_627_v48).length));
      (_627_v48)[_index86] = _623_v46;
      let _index87 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_627_v48).length));
      (_627_v48)[_index87] = _dafny.Seq.Concat(_623_v46, _623_v46);
      let _638_v49;
      _638_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_622_v45).length),_599_v26);
      let _639_v50;
      _639_v50 = _dafny.MultiSet.fromElements(_module.__default.fm25(globalState), _dafny.Map.Empty.slice().updateUnsafe((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))],new BigNumber((_622_v45).length)), _638_v49, _638_v49, _dafny.Map.Empty.slice().updateUnsafe(_599_v26,_599_v26));
      (globalState).f0 = _module.__default.fm21(new BigNumber(280), (_639_v50).Difference(_639_v50), globalState);
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_605_v32).length))) {
        let _640_i10 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_640_i10)) && ((_640_i10).isLessThan(new BigNumber((_605_v32).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_605_v32, (_640_i10).toNumber(), (_640_i10).multipliedBy((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))])));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _641_v51;
      _641_v51 = _module.D1.create_DC5(p1.f3, (_this).f13, new BigNumber(((_this).f12).length), (_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]);
      r0 = _dafny.Seq.Concat(_module.__default.fm30(_641_v51, (_dafny.ZERO).minus((_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))]), (_605_v32)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_605_v32).length))], globalState), (_this).f12);
      return r0;
    }
    m10(globalState) {
      let _this = this;
      let _642_v0;
      _642_v0 = new BigNumber(391);
      _642_v0 = _642_v0;
      let _643_v1;
      _643_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_642_v0);
      let _644_v2;
      _644_v2 = _dafny.Set.fromElements((_this).f13);
      let _645_v3;
      let _out19;
      _out19 = (_this).m11((_this).f13, new BigNumber((_643_v1).length), (_644_v2).Difference(_644_v2), globalState);
      _645_v3 = _out19;
      let _646_v4;
      _646_v4 = _dafny.Seq.of((_this).f13);
      let _647_v5;
      let _nw95 = Array((new BigNumber(11)).toNumber());
      _nw95[(_dafny.ZERO).toNumber()] = _642_v0;
      _nw95[(_dafny.ONE).toNumber()] = new BigNumber(((_this).f12).length);
      _nw95[(new BigNumber(2)).toNumber()] = new BigNumber(((_this).f12).length);
      _nw95[(new BigNumber(3)).toNumber()] = new BigNumber(204);
      _nw95[(new BigNumber(4)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(5)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(6)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(7)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(8)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(9)).toNumber()] = _642_v0;
      _nw95[(new BigNumber(10)).toNumber()] = _module.__default.fm31((_this).f13, _642_v0, (_this).f13, _646_v4, globalState);
      _647_v5 = _nw95;
      let _648_v6;
      _648_v6 = _dafny.Seq.of(_647_v5, _647_v5, _647_v5);
      let _649_v7;
      _649_v7 = _648_v6;
      let _650_v8;
      let _nw96 = Array((new BigNumber(14)).toNumber());
      _nw96[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_648_v6, _module.__default.safeIndex(new BigNumber(590), new BigNumber((_648_v6).length)), _647_v5), _module.__default.safeIndex(_642_v0, new BigNumber((_dafny.Seq.update(_648_v6, _module.__default.safeIndex(new BigNumber(590), new BigNumber((_648_v6).length)), _647_v5)).length)), _647_v5), _dafny.Seq.update(_dafny.Seq.of(_647_v5), _module.__default.safeIndex(_642_v0, new BigNumber((_dafny.Seq.of(_647_v5)).length)), _647_v5));
      _nw96[(_dafny.ONE).toNumber()] = (_649_v7);
      _nw96[(new BigNumber(2)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(3)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(4)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_648_v6, _648_v6);
      _nw96[(new BigNumber(6)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(7)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(8)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(9)).toNumber()] = _648_v6;
      _nw96[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_648_v6, _648_v6);
      _nw96[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_647_v5);
      _nw96[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_648_v6, _dafny.Seq.update(_dafny.Seq.of(_647_v5), _module.__default.safeIndex(_642_v0, new BigNumber((_dafny.Seq.of(_647_v5)).length)), _647_v5));
      _nw96[(new BigNumber(13)).toNumber()] = _648_v6;
      _650_v8 = _nw96;
      let _index88 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_650_v8).length));
      (_650_v8)[_index88] = _dafny.Seq.Concat(_648_v6, _648_v6);
      let _index89 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_650_v8).length));
      (_650_v8)[_index89] = _648_v6;
      _642_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_642_v0, _642_v0));
      let _651_v9;
      _651_v9 = _dafny.Seq.UnicodeFromString("ssfq");
      _651_v9 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uqxt"), ((true) ? ((_this).f12) : (_651_v9)));
      (globalState).f0 = (_this).f13;
      return;
    }
    m11(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _652_v0;
      let _init11 = ((_653_p0) => function (_654_i0) {
        return _653_p0;
      })(p0);
      let _nw97 = Array((new BigNumber(5)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw97.length); _i0_11++) {
        _nw97[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _652_v0 = _nw97;
      _652_v0 = _652_v0;
      let _655_v1;
      let _nw98 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _655_v1 = _nw98;
      _655_v1 = _655_v1;
      let _656_i1;
      _656_i1 = _dafny.ZERO;
      L2: {
        while (p0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_656_i1)) {
              break L2;
            }
            _656_i1 = (_656_i1).plus(_dafny.ONE);
            let _657_v2;
            _657_v2 = new _dafny.CodePoint('t'.codePointAt(0));
            let _658_v3;
            _658_v3 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), _657_v2);
            _658_v3 = _658_v3;
            let _659_v4;
            _659_v4 = new BigNumber(-429);
            _659_v4 = p1;
            let _660_v5;
            _660_v5 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(253)), ((_661_v2) => function (_662_i2) {
              return _661_v2;
            })(_657_v2)), (_this).f12), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gbd"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("gbd")).length)), _657_v2), (_this).f12), (_this).f12, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lkugwmyv"), (_this).f12));
            _660_v5 = _dafny.Seq.Concat(_660_v5, _660_v5);
            let _init12 = ((_663_v4) => function (_664_i3) {
              return (_663_v4).isEqualTo(_663_v4);
            })(_659_v4);
            let _nw99 = Array((new BigNumber(11)).toNumber());
            for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw99.length); _i0_12++) {
              _nw99[_i0_12] = _init12(new BigNumber(_i0_12));
            }
            _652_v0 = _nw99;
          }
        }
      }
      let _665_v6;
      _665_v6 = _dafny.Seq.of(p1);
      (globalState).f0 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(p1, p1, p1, p1, p1), _dafny.Seq.Concat(_665_v6, _665_v6)));
      let _666_v7;
      _666_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _667_v8;
      _667_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _668_v9;
      _668_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_667_v8).length),_666_v7);
      _666_v7 = ((((_668_v9).contains(p1)) ? ((_668_v9).get(p1)) : (_666_v7))).update(p1, p1);
      let _669_v10;
      _669_v10 = _dafny.Seq.of((_this).f12);
      let _670_v11;
      _670_v11 = _669_v10;
      let _671_v12;
      _671_v12 = _module.D0.create_DC2((_this).f13, new BigNumber(((_670_v11)).length), p1);
      let _pat_let_tv31 = p0;
      if (function (_source11) {
        if (_source11.is_DC1) {
          let _672___mcc_h6 = (_source11).cf1;
          let _673___mcc_h7 = (_source11).cf2;
          let _674_cf2 = _673___mcc_h7;
          let _675_cf1 = _672___mcc_h6;
          return _675_cf1;
        } else if (_source11.is_DC2) {
          let _676___mcc_h8 = (_source11).cf3;
          let _677___mcc_h9 = (_source11).cf4;
          let _678___mcc_h10 = (_source11).cf5;
          let _679_cf5 = _678___mcc_h10;
          let _680_cf4 = _677___mcc_h9;
          let _681_cf3 = _676___mcc_h8;
          return (_this).f13;
        } else {
          let _682___mcc_h11 = (_source11).cf0;
          let _683_cf0 = _682___mcc_h11;
          return _pat_let_tv31;
        }
      }(_671_v12)) {
        let _684_v13;
        _684_v13 = _module.D1.create_DC3(_652_v0);
        let _source12 = _684_v13;
        if (_source12.is_DC4) {
          let _685___mcc_h0 = (_source12).cf7;
          let _686_cf7 = _685___mcc_h0;
          let _687_v14;
          _687_v14 = _dafny.Set.fromElements(p1);
          let _688_v15;
          _688_v15 = _dafny.Seq.of((_this).f13);
          let _689_v16;
          _689_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f12).length),(_this).f12);
          let _690_v17;
          _690_v17 = _dafny.Seq.of(_module.__default.fm32(true, false, p0, (_this).f13, globalState), _687_v14, _dafny.Set.fromElements(_686_cf7), _dafny.Set.fromElements(_686_cf7, (_dafny.ZERO).minus(new BigNumber((_688_v15).length))), _dafny.Set.fromElements(p1, new BigNumber((_689_v16).length)));
          let _691_v19;
          _691_v19 = _dafny.MultiSet.fromElements(_module.__default.fm25(globalState));
          (globalState).f0 = _module.__default.fm21(new BigNumber((_dafny.Seq.Concat(_690_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), ((_692_p1) => function (_693_i4) {
            return function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of (_dafny.MultiSet.fromElements(_692_p1)).Elements) {
                let _694_v18 = _compr_27;
                if ((_dafny.MultiSet.fromElements(_692_p1)).contains(_694_v18)) {
                  _coll27.add((_694_v18).minus(new BigNumber(762)));
                }
              }
              return _coll27;
            }();
          })(p1)))).length), _691_v19, globalState);
          (globalState).f0 = (_688_v15)[_module.__default.safeIndex(_686_cf7, new BigNumber((_688_v15).length))];
          let _695_v20;
          let _nw100 = new _module.C0();
          _nw100.__ctor();
          _695_v20 = _nw100;
          let _696_v21;
          _696_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-805)), function (_697_i5) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("wcygaqfhi")),_667_v8);
          let _698_v22;
          let _nw101 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
          _698_v22 = _nw101;
          let _699_v23;
          _699_v23 = _dafny.MultiSet.fromElements(_687_v14);
          let _index90 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_698_v22).length));
          (_698_v22)[_index90] = (_699_v23).Union(_699_v23);
          let _700_v25;
          _700_v25 = _dafny.Set.fromElements((_this).f12);
          let _701_v28;
          _701_v28 = _dafny.Map.Empty.slice().updateUnsafe(_670_v11,function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(-500), new BigNumber(373))) {
              let _702_v27 = _compr_28;
              if (((new BigNumber(-500)).isLessThanOrEqualTo(_702_v27)) && ((_702_v27).isLessThan(new BigNumber(373)))) {
                _coll28.add((_702_v27).plus(new BigNumber((_dafny.Set.fromElements((_this).f13)).length)));
              }
            }
            return _coll28;
          }());
          let _index91 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_698_v22).length));
          let _rhs82 = (function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_700_v25).Elements) {
              let _703_v24 = _compr_29;
              if ((_700_v25).contains(_703_v24)) {
                _coll29.push([_703_v24,_667_v8]);
              }
            }
            return _coll29;
          }()).Merge((function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_669_v10).Elements) {
              let _704_v26 = _compr_30;
              if (_dafny.Seq.contains(_669_v10, _704_v26)) {
                _coll30.push([_704_v26,_667_v8]);
              }
            }
            return _coll30;
          }()).update((_this).f12, _667_v8));
          let _rhs83 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(_686_cf7, new BigNumber(-787)))).isLessThanOrEqualTo(p1);
          let _rhs84 = _686_cf7;
          let _rhs85 = _dafny.MultiSet.fromElements((((_701_v28).contains(_670_v11)) ? ((_701_v28).get(_670_v11)) : (_687_v14)));
          let _lhs58 = globalState;
          let _lhs59 = _698_v22;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_698_v22).length));
          _696_v21 = _rhs82;
          _lhs58.f0 = _rhs83;
          _686_cf7 = _rhs84;
          _lhs59[_lhs60] = _rhs85;
        } else if (_source12.is_DC5) {
          let _705___mcc_h1 = (_source12).cf8;
          let _706___mcc_h2 = (_source12).cf9;
          let _707___mcc_h3 = (_source12).cf10;
          let _708___mcc_h4 = (_source12).cf11;
          let _709_cf11 = _708___mcc_h4;
          let _710_cf10 = _707___mcc_h3;
          let _711_cf9 = _706___mcc_h2;
          let _712_cf8 = _705___mcc_h1;
          let _713_v29;
          _713_v29 = _dafny.Map.Empty.slice().updateUnsafe(((_711_cf9) ? (true) : ((_this).f13)),_712_cf8);
          _713_v29 = (_713_v29).update((_666_v7).contains(_709_cf11), false);
          let _714_v30;
          _714_v30 = _dafny.MultiSet.fromElements(_666_v7, _666_v7);
          let _index92 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_652_v0).length));
          (_652_v0)[_index92] = _module.__default.fm21(_709_cf11, _714_v30, globalState);
          let _index93 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_652_v0).length));
          (_652_v0)[_index93] = _712_cf8;
          let _715_v31;
          _715_v31 = _module.D1.create_DC5(_711_cf9, (_652_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_652_v0).length))], _709_cf11, (_dafny.ZERO).minus((_665_v6)[_module.__default.safeIndex(_710_cf10, new BigNumber((_665_v6).length))]));
          let _716_v32;
          _716_v32 = _dafny.Map.Empty.slice().updateUnsafe(_715_v31,p1);
          _716_v32 = _716_v32;
          _710_cf10 = _710_cf10;
        } else if (_source12.is_DC6) {
          let _index94 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length));
          (_655_v1)[_index94] = p1;
          let _index95 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length));
          (_655_v1)[_index95] = _module.__default.safeModuloInt(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bsas"), (_this).f12)).length));
          let _717_v33;
          let _nw102 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _717_v33 = _nw102;
          _717_v33 = _717_v33;
          let _718_v34;
          _718_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f12);
          let _719_v36;
          _719_v36 = _module.D7.create_DC16(function () {
  let _coll31 = new _dafny.Map();
  for (const _compr_31 of _dafny.IntegerRange(new BigNumber(483), new BigNumber(839))) {
    let _720_v35 = _compr_31;
    if (((new BigNumber(483)).isLessThanOrEqualTo(_720_v35)) && ((_720_v35).isLessThan(new BigNumber(839)))) {
      _coll31.push([_module.__default.safeDivisionInt(_720_v35, new BigNumber((_dafny.Seq.of(p0)).length)),p2]);
    }
  }
  return _coll31;
}());
          let _721_v37;
          _721_v37 = _dafny.Map.Empty.slice().updateUnsafe((((_718_v34).contains(!((_this).f13))) ? ((_718_v34).get(!((_this).f13))) : (_dafny.Seq.UnicodeFromString("bchjxx"))),(_719_v36).dtor_cf25);
          let _722_v38;
          _722_v38 = _module.D1.create_DC5(true, p0, new BigNumber(((_this).f12).length), (_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))]);
          let _723_v39;
          _723_v39 = _dafny.MultiSet.fromElements((_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))]);
          let _724_v40;
          _724_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _pat_let_tv32 = p0;
          _721_v37 = (_721_v37).update(_module.__default.fm30(function (_pat_let11_0) {
            return function (_725_dt__update__tmp_h0) {
              return function (_pat_let12_0) {
                return function (_726_dt__update_hcf9_h0) {
                  return function (_pat_let13_0) {
                    return function (_727_dt__update_hcf8_h0) {
                      return _module.D1.create_DC5(_727_dt__update_hcf8_h0, _726_dt__update_hcf9_h0, (_725_dt__update__tmp_h0).dtor_cf10, (_725_dt__update__tmp_h0).dtor_cf11);
                    }(_pat_let13_0);
                  }(true);
                }(_pat_let12_0);
              }(_pat_let_tv32);
            }(_pat_let11_0);
          }(_722_v38), (_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))], new BigNumber((_723_v39).cardinality()), globalState), _724_v40);
          let _728_v41;
          _728_v41 = _module.D0.create_DC0((_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))]);
          (globalState).f0 = (_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))]), p1, (_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))], (_655_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_655_v1).length))], (_728_v41).dtor_cf0)).IsSubsetOf(((_723_v39).update(new BigNumber(211), _module.__default.abs(new BigNumber(178)))).Intersect(_723_v39));
        } else {
          let _729___mcc_h5 = (_source12).cf6;
          let _730_cf6 = _729___mcc_h5;
          let _731_v42;
          _731_v42 = new BigNumber(-940);
          let _732_v43;
          let _nw103 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _732_v43 = _nw103;
          let _733_v44;
          _733_v44 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index96 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_732_v43).length));
          (_732_v43)[_index96] = _733_v44;
          let _734_v45;
          _734_v45 = _dafny.MultiSet.fromElements(_module.D7.create_DC17(p1, p0, p0, p1));
          let _index97 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_732_v43).length));
          let _rhs86 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat((_this).f12, (_this).f12), (_this).f12)).length);
          let _rhs87 = ((_this).f12)[_module.__default.safeIndex(_731_v42, new BigNumber(((_this).f12).length))];
          let _rhs88 = !((_734_v45).IsSubsetOf(_734_v45));
          let _rhs89 = _652_v0;
          let _lhs61 = _732_v43;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_732_v43).length));
          let _lhs63 = globalState;
          _731_v42 = _rhs86;
          _lhs61[_lhs62] = _rhs87;
          _lhs63.f0 = _rhs88;
          _730_cf6 = _rhs89;
          (globalState).f0 = p0;
          _732_v43 = _732_v43;
          let _735_v46;
          _735_v46 = _dafny.MultiSet.fromElements(_655_v1);
          let _736_v47;
          _736_v47 = _dafny.MultiSet.fromElements(p0, p0);
          let _737_v48;
          _737_v48 = _dafny.MultiSet.fromElements(new BigNumber(((_736_v47).Union(_736_v47)).cardinality()));
          let _738_v49;
          _738_v49 = _dafny.Seq.UnicodeFromString("jnrvw");
          let _rhs90 = (_735_v46).Union((_dafny.MultiSet.fromElements(_655_v1)).Union(_dafny.MultiSet.fromElements(_655_v1)));
          let _rhs91 = _module.__default.fm24(globalState);
          let _rhs92 = p1;
          let _rhs93 = _dafny.Seq.Concat(_738_v49, _dafny.Seq.UnicodeFromString("gwbrdrf"));
          let _rhs94 = p1;
          _735_v46 = _rhs90;
          _737_v48 = _rhs91;
          _731_v42 = _rhs92;
          _738_v49 = _rhs93;
          _731_v42 = _rhs94;
        }
        let _739_v50;
        _739_v50 = new BigNumber(-678);
        let _740_v51;
        _740_v51 = _dafny.Seq.of((_this).f13, p0);
        _739_v50 = (_module.D7.create_DC17(p1, p0, false, _module.__default.fm31(p0, _739_v50, (_this).f13, _740_v51, globalState))).dtor_cf26;
        let _index98 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length));
        (_652_v0)[_index98] = (_this).f13;
        let _index99 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length));
        (_652_v0)[_index99] = (_this).f13;
        let _index100 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length));
        (_652_v0)[_index100] = (_this).f13;
        if (!(p1).isEqualTo(p1)) {
          let _index101 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length));
          (_652_v0)[_index101] = (_652_v0)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length))];
          let _741_v52;
          _741_v52 = _dafny.MultiSet.fromElements((_this).f13);
          let _742_v53;
          _742_v53 = _dafny.Seq.of(_741_v52);
          (globalState).f0 = ((_742_v53)[_module.__default.safeIndex(_739_v50, new BigNumber((_742_v53).length))]).IsDisjointFrom(_741_v52);
          let _index102 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_655_v1).length));
          (_655_v1)[_index102] = new BigNumber(709);
          let _index103 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_655_v1).length));
          (_655_v1)[_index103] = _739_v50;
          let _index104 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_655_v1).length));
          (_655_v1)[_index104] = _739_v50;
          let _743_v54;
          _743_v54 = _dafny.Set.fromElements((_652_v0)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length))]);
          let _744_v55;
          _744_v55 = _module.D1.create_DC5(false, p0, _739_v50, p1);
          let _745_v56;
          _745_v56 = _module.D0.create_DC0(p1);
          let _746_v57;
          _746_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_667_v8);
          let _747_v58;
          _747_v58 = _dafny.Seq.of(_746_v57, _746_v57);
          let _index105 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_655_v1).length));
          let _index106 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_655_v1).length));
          let _index107 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_655_v1).length));
          let _rhs95 = _module.__default.safeDivisionInt(p1, _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_652_v0)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_652_v0).length))],new BigNumber(152))).length), _739_v50));
          let _rhs96 = _dafny.areEqual(_module.__default.fm30(_744_v55, (_745_v56).dtor_cf0, _739_v50, globalState), _dafny.Seq.Concat((_this).f12, _module.__default.fm30(_744_v55, p1, p1, globalState)));
          let _rhs97 = _739_v50;
          let _rhs98 = new BigNumber((((_dafny.Seq.IsPrefixOf(_740_v51, _740_v51)) ? (((_747_v58)[_module.__default.safeIndex(new BigNumber((_666_v7).length), new BigNumber((_747_v58).length))]).Merge(_746_v57)) : (_746_v57))).length);
          let _rhs99 = (p2).Difference(p2);
          let _lhs64 = _655_v1;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_655_v1).length));
          let _lhs66 = globalState;
          let _lhs67 = _655_v1;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_655_v1).length));
          let _lhs69 = _655_v1;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_655_v1).length));
          _lhs64[_lhs65] = _rhs95;
          _lhs66.f0 = _rhs96;
          _lhs67[_lhs68] = _rhs97;
          _lhs69[_lhs70] = _rhs98;
          _743_v54 = _rhs99;
          let _748_v59;
          _748_v59 = _dafny.MultiSet.fromElements(_739_v50);
          let _749_v60;
          _749_v60 = _dafny.MultiSet.fromElements(_748_v59);
          let _750_v61;
          _750_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_749_v60);
          let _751_v62;
          _751_v62 = _dafny.Seq.of(_748_v59);
          let _rhs100 = new BigNumber(108);
          let _rhs101 = _dafny.Seq.of(new BigNumber(((_667_v8).Merge(_667_v8)).length));
          let _rhs102 = (((_750_v61).contains((_655_v1)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_655_v1).length))])) ? ((_750_v61).get((_655_v1)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_655_v1).length))])) : (_dafny.MultiSet.FromArray(_751_v62)));
          _739_v50 = _rhs100;
          _665_v6 = _rhs101;
          _749_v60 = _rhs102;
          let _752_v63;
          let _init13 = ((_753_v1) => function (_754_i6) {
            return (_754_i6).minus((_753_v1)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_753_v1).length))]);
          })(_655_v1);
          let _nw104 = Array((new BigNumber(11)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw104.length); _i0_13++) {
            _nw104[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _752_v63 = _nw104;
          _752_v63 = (((_this).f13) ? (_752_v63) : (_752_v63));
        } else {
          let _755_v64;
          _755_v64 = _dafny.Map.Empty.slice().updateUnsafe(_739_v50,!(p0));
          let _756_v66;
          _756_v66 = _dafny.Set.fromElements(p1);
          let _757_v67;
          _757_v67 = _dafny.MultiSet.fromElements(_755_v64, function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_756_v66).Elements) {
              let _758_v65 = _compr_32;
              if ((_756_v66).contains(_758_v65)) {
                _coll32.push([_module.__default.safeDivisionInt(_758_v65, p1),p0]);
              }
            }
            return _coll32;
          }());
          let _759_v68;
          _759_v68 = _module.D3.create_DC9(_757_v67);
          _759_v68 = _759_v68;
          let _760_v69;
          _760_v69 = new _dafny.CodePoint('x'.codePointAt(0));
          let _761_v70;
          let _init14 = ((_762_v69) => function (_763_i7) {
            return _762_v69;
          })(_760_v69);
          let _nw105 = Array((_dafny.ONE).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw105.length); _i0_14++) {
            _nw105[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _761_v70 = _nw105;
          let _764_v71;
          _764_v71 = _dafny.Set.fromElements(_761_v70);
          _760_v69 = _module.__default.fm23((_764_v71).IsDisjointFrom(_764_v71), new BigNumber((_665_v6).length), globalState);
          let _765_v72;
          let _nw106 = new _module.C0();
          _nw106.__ctor();
          _765_v72 = _nw106;
          let _index108 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_655_v1).length));
          (_655_v1)[_index108] = _739_v50;
          let _766_v73;
          _766_v73 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_760_v69);
          let _767_v74;
          _767_v74 = _dafny.Map.Empty.slice().updateUnsafe(_739_v50,_766_v73);
          let _768_v75;
          _768_v75 = _dafny.MultiSet.fromElements((((_767_v74).contains(new BigNumber(899))) ? ((_767_v74).get(new BigNumber(899))) : (_766_v73)));
          let _769_v76;
          _769_v76 = _dafny.Seq.UnicodeFromString("xslucuycd");
          let _index109 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_655_v1).length));
          let _rhs103 = _739_v50;
          let _rhs104 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p1, _739_v50), (p1).multipliedBy(p1));
          let _rhs105 = (_768_v75).Difference((_768_v75).Union(_768_v75));
          let _rhs106 = _760_v69;
          let _rhs107 = _769_v76;
          let _lhs71 = _655_v1;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_655_v1).length));
          _739_v50 = _rhs103;
          _lhs71[_lhs72] = _rhs104;
          _768_v75 = _rhs105;
          _760_v69 = _rhs106;
          _769_v76 = _rhs107;
          (globalState).f0 = (_this).f13;
        }
      } else {
        let _770_v77;
        _770_v77 = _dafny.Seq.UnicodeFromString("cvlcjvnp");
        _770_v77 = (_this).f12;
        let _771_v78;
        _771_v78 = new _dafny.CodePoint('q'.codePointAt(0));
        let _772_v79;
        _772_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_771_v78);
        let _773_v80;
        _773_v80 = _dafny.Map.Empty.slice().updateUnsafe(!((p1).isLessThan(p1)),_772_v79);
        _773_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,_772_v79);
        let _774_v81;
        _774_v81 = new BigNumber(337);
        let _775_v82;
        _775_v82 = _dafny.Seq.of(false);
        let _776_v83;
        _776_v83 = _dafny.MultiSet.fromElements(_666_v7);
        let _777_v84;
        _777_v84 = _dafny.Seq.of((_this).f13, _module.__default.fm21(new BigNumber((_dafny.Seq.of(_775_v82, _dafny.Seq.of(p0))).length), _776_v83, globalState));
        _774_v81 = ((new BigNumber(886)).multipliedBy(p1)).plus(_module.__default.fm31(p0, _774_v81, false, _777_v84, globalState));
        let _778_v85;
        let _init15 = function (_779_i8) {
          return (_this).f12;
        };
        let _nw107 = Array((new BigNumber(12)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw107.length); _i0_15++) {
          _nw107[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _778_v85 = _nw107;
        let _index110 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_778_v85).length));
        (_778_v85)[_index110] = _dafny.Seq.UnicodeFromString("xoffjr");
        let _index111 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_778_v85).length));
        (_778_v85)[_index111] = (_this).f12;
        let _index112 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_655_v1).length));
        (_655_v1)[_index112] = new BigNumber(267);
        let _index113 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_655_v1).length));
        (_655_v1)[_index113] = p1;
      }
      let _780_v86;
      _780_v86 = _dafny.Map.Empty.slice().updateUnsafe(_669_v10,_655_v1);
      let _781_v87;
      _781_v87 = _module.D1.create_DC3(_652_v0);
      let _782_v88;
      _782_v88 = _dafny.Map.Empty.slice().updateUnsafe((((_780_v86).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), function (_784_i9) {
        return _dafny.Seq.UnicodeFromString("tyathfrj");
      }))) ? ((_780_v86).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), function (_783_i9) {
        return _dafny.Seq.UnicodeFromString("tyathfrj");
      }))) : (_655_v1)),_781_v87);
      r0 = _782_v88;
      return r0;
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
      this._f3 = false;
      this._f4 = false;
      this._f10 = false;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f10, f11, f3, f4) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(338),new BigNumber(547)))).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length),new BigNumber(597))))).cardinality())).multipliedBy((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(832))).length))).cardinality()),_this.f3)).length))).multipliedBy(new BigNumber(337))));
    };
    fm16(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source13 = _module.D0.create_DC0(new BigNumber(565));
      if (_source13.is_DC1) {
        let _785___mcc_h0 = (_source13).cf1;
        let _786___mcc_h1 = (_source13).cf2;
        let _787_cf2 = _786___mcc_h1;
        let _788_cf1 = _785___mcc_h0;
        return new BigNumber(482);
      } else if (_source13.is_DC2) {
        let _789___mcc_h2 = (_source13).cf3;
        let _790___mcc_h3 = (_source13).cf4;
        let _791___mcc_h4 = (_source13).cf5;
        let _792_cf5 = _791___mcc_h4;
        let _793_cf4 = _790___mcc_h3;
        let _794_cf3 = _789___mcc_h2;
        return (_793_cf4).minus(new BigNumber(-730));
      } else {
        let _795___mcc_h5 = (_source13).cf0;
        let _796_cf0 = _795___mcc_h5;
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_dafny.ZERO).minus(_796_cf0))).length);
      }
    };
    fm17(globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("jdorts");
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _797_v0;
      _797_v0 = _dafny.Seq.of(p2);
      let _hi4 = (_797_v0)[_module.__default.safeIndex(p2, new BigNumber((_797_v0).length))];
      for (let _798_i0 = new BigNumber((_dafny.Seq.UnicodeFromString("gqk")).length); _798_i0.isLessThan(_hi4); _798_i0 = _798_i0.plus(_dafny.ONE)) {
        let _799_v1;
        let _init16 = ((_800_p2) => function (_801_i1) {
          return (_800_p2).isLessThanOrEqualTo(new BigNumber(33));
        })(p2);
        let _nw108 = Array((new BigNumber(15)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw108.length); _i0_16++) {
          _nw108[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _799_v1 = _nw108;
        let _index114 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_799_v1).length));
        (_799_v1)[_index114] = true;
        let _index115 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_799_v1).length));
        (_799_v1)[_index115] = !((_this).f4);
        r1 = ((false) ? (p2) : (p2));
        let _802_v2;
        _802_v2 = new _dafny.CodePoint('v'.codePointAt(0));
        let _803_v3;
        _803_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_799_v1);
        let _804_v4;
        _804_v4 = _module.D1.create_DC3(_799_v1);
        let _805_v5;
        let _nw109 = Array((new BigNumber(28)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = _799_v1;
        _nw109[(_dafny.ONE).toNumber()] = _799_v1;
        _nw109[(new BigNumber(2)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(3)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(4)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(5)).toNumber()] = ((_module.__default.fm18(_798_i0, _802_v2, _798_i0, globalState)) ? (_799_v1) : (_799_v1));
        _nw109[(new BigNumber(6)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(7)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(8)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(9)).toNumber()] = (((_803_v3).contains((_this).f11)) ? ((_803_v3).get((_this).f11)) : (_799_v1));
        _nw109[(new BigNumber(10)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(11)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(12)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(13)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(14)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(15)).toNumber()] = ((p0) ? ((((_803_v3).contains((_this).f4)) ? ((_803_v3).get((_this).f4)) : ((_804_v4).dtor_cf6))) : (_799_v1));
        _nw109[(new BigNumber(16)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(17)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(18)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(19)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(20)).toNumber()] = (((_799_v1)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_799_v1).length))]) ? (_799_v1) : (_799_v1));
        _nw109[(new BigNumber(21)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(22)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(23)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(24)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(25)).toNumber()] = ((true) ? (_799_v1) : (_799_v1));
        _nw109[(new BigNumber(26)).toNumber()] = _799_v1;
        _nw109[(new BigNumber(27)).toNumber()] = _799_v1;
        _805_v5 = _nw109;
        let _806_v6;
        _806_v6 = _dafny.Seq.of(_this.f3, !((_this).f4));
        let _807_v7;
        _807_v7 = _dafny.Seq.UnicodeFromString("y");
        let _808_v8;
        _808_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_806_v6, _module.__default.safeIndex(new BigNumber((_807_v7).length), new BigNumber((_806_v6).length)), false)).length),_799_v1);
        let _index116 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_805_v5).length));
        (_805_v5)[_index116] = (((_808_v8).contains(_798_i0)) ? ((_808_v8).get(_798_i0)) : (_799_v1));
        let _index117 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_805_v5).length));
        (_805_v5)[_index117] = _799_v1;
        r1 = (_dafny.ZERO).minus(((_798_i0).plus(new BigNumber(819))).multipliedBy(_798_i0));
      }
      let _809_v9;
      _809_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
      r1 = _module.__default.safeDivisionInt(p2, new BigNumber((_809_v9).length));
      (_this).f3 = !(true);
      let _810_v10;
      _810_v10 = _dafny.Set.fromElements((_this).f11);
      let _811_v11;
      _811_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_810_v10).length));
      let _812_v12;
      _812_v12 = _dafny.MultiSet.fromElements((_this).f4, (_this).f10, !((_this).f4));
      let _813_v13;
      _813_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p2);
      let _814_v14;
      _814_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_813_v13);
      let _815_v15;
      let _nw110 = Array((new BigNumber(13)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = (((_811_v11).contains(p2)) ? ((_811_v11).get(p2)) : (p2));
      _nw110[(_dafny.ONE).toNumber()] = p2;
      _nw110[(new BigNumber(2)).toNumber()] = new BigNumber(491);
      _nw110[(new BigNumber(3)).toNumber()] = (_797_v0)[_module.__default.safeIndex(p2, new BigNumber((_797_v0).length))];
      _nw110[(new BigNumber(4)).toNumber()] = (((_812_v12).contains((_this).f4)) ? ((_812_v12).get((_this).f4)) : (p2));
      _nw110[(new BigNumber(5)).toNumber()] = p2;
      _nw110[(new BigNumber(6)).toNumber()] = p2;
      _nw110[(new BigNumber(7)).toNumber()] = new BigNumber((_814_v14).length);
      _nw110[(new BigNumber(8)).toNumber()] = p2;
      _nw110[(new BigNumber(9)).toNumber()] = p2;
      _nw110[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw110[(new BigNumber(11)).toNumber()] = p2;
      _nw110[(new BigNumber(12)).toNumber()] = p2;
      _815_v15 = _nw110;
      let _816_v16;
      _816_v16 = _dafny.Map.Empty.slice().updateUnsafe(_815_v15,false);
      (globalState).f0 = (_816_v16).contains(_815_v15);
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_815_v15).length))) {
        let _817_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_817_i2)) && ((_817_i2).isLessThan(new BigNumber((_815_v15).length))))) {
          (_815_v15)[(_817_i2)] = _module.__default.safeModuloInt(_817_i2, p2);
        }
      }
      if ((_this).f4) {
        let _818_v17;
        _818_v17 = new _dafny.CodePoint('n'.codePointAt(0));
        let _819_v18;
        _819_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm18(p2, _818_v17, p2, globalState)),(_this).fm17(globalState));
        _819_v18 = (_dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), ((_820_v17) => function (_821_i3) {
          return _820_v17;
        })(_818_v17)))).Merge((_819_v18).Merge(_819_v18));
        let _822_v19;
        _822_v19 = _dafny.MultiSet.fromElements(_809_v9);
        _797_v0 = _dafny.Seq.of((p2).plus(new BigNumber((_822_v19).cardinality())));
        let _823_v20;
        _823_v20 = _dafny.Seq.of(_813_v13);
        r1 = (_this).fm16(p2, !(_dafny.MultiSet.fromElements(_this.f3, (_this).f4, (_this).f11, (_this).f11, (_this).f4)).contains(p0), (_dafny.ZERO).minus(new BigNumber((_823_v20).length)), (p1).Merge(p1), globalState);
        (_this).f3 = _this.f3;
        let _index118 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_815_v15).length));
        (_815_v15)[_index118] = (new BigNumber(270)).plus(p2);
        let _index119 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_815_v15).length));
        (_815_v15)[_index119] = p2;
      } else {
        let _824_v21;
        let _nw111 = new _module.C0();
        _nw111.__ctor();
        _824_v21 = _nw111;
        let _825_v22;
        let _nw112 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
        _825_v22 = _nw112;
        let _826_v23;
        _826_v23 = _dafny.Seq.UnicodeFromString("pgmkqgjiw");
        let _index120 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_825_v22).length));
        (_825_v22)[_index120] = _dafny.Seq.update(_module.__default.fm19(globalState), _module.__default.safeIndex(new BigNumber((_826_v23).length), new BigNumber((_module.__default.fm19(globalState)).length)), p2);
        let _827_v24;
        _827_v24 = new _dafny.CodePoint('e'.codePointAt(0));
        let _828_v25;
        _828_v25 = _dafny.Seq.of(_module.__default.fm20(p0, globalState), _812_v12);
        let _index121 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_825_v22).length));
        let _rhs108 = (((_this).f10) ? (_824_v21) : (_824_v21));
        let _rhs109 = false;
        let _rhs110 = !(_module.__default.fm18(p2, _827_v24, (_824_v21).fm10(new BigNumber(771), (_this).f10, _dafny.MultiSet.fromElements(new BigNumber((_828_v25).length)), globalState), globalState));
        let _rhs111 = _module.__default.fm19(globalState);
        let _rhs112 = _this.f3;
        let _lhs73 = _this;
        let _lhs74 = globalState;
        let _lhs75 = _825_v22;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_825_v22).length));
        let _lhs77 = globalState;
        _824_v21 = _rhs108;
        _lhs73.f3 = _rhs109;
        _lhs74.f0 = _rhs110;
        _lhs75[_lhs76] = _rhs111;
        _lhs77.f0 = _rhs112;
        _827_v24 = _827_v24;
        let _829_v26;
        _829_v26 = _module.D1.create_DC5(_module.__default.fm18(p2, _827_v24, p2, globalState), (_this).f4, p2, (_dafny.ZERO).minus(p2));
        (globalState).f0 = (_829_v26).dtor_cf9;
        _826_v23 = _dafny.Seq.UnicodeFromString("tgljdqvut");
        (globalState).f0 = (_this).f11;
      }
      r0 = ((_this.f3) ? (_815_v15) : (_815_v15));
      r1 = p2;
      return [r0, r1];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _830_v0;
      _830_v0 = new _dafny.CodePoint('y'.codePointAt(0));
      _830_v0 = new _dafny.CodePoint('p'.codePointAt(0));
      let _831_v1;
      _831_v1 = _dafny.Seq.UnicodeFromString("tenl");
      let _832_v2;
      let _nw113 = new _module.C1();
      _nw113.__ctor(_831_v1, (_this).f4);
      _832_v2 = _nw113;
      _832_v2 = _832_v2;
      let _833_v3;
      let _nw114 = Array((new BigNumber(14)).toNumber()).fill(false);
      _833_v3 = _nw114;
      let _index122 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length));
      (_833_v3)[_index122] = false;
      let _834_v4;
      let _nw115 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
      _834_v4 = _nw115;
      let _835_v5;
      _835_v5 = _dafny.Set.fromElements((p1).f4);
      let _index123 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length));
      let _rhs113 = ((_this.f3) ? ((p1).f4) : ((p1).f4));
      let _rhs114 = ((_835_v5).Union(_835_v5)).IsSubsetOf(_835_v5);
      let _rhs115 = _834_v4;
      let _lhs78 = _this;
      let _lhs79 = _833_v3;
      let _lhs80 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length));
      _lhs78.f3 = _rhs113;
      _lhs79[_lhs80] = _rhs114;
      _834_v4 = _rhs115;
      let _836_v6;
      _836_v6 = _dafny.Seq.of((_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))], (p1).f4, (p1).f4, (_this).f4, (_this).f4);
      let _837_v7;
      _837_v7 = _dafny.Map.Empty.slice().updateUnsafe((_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))],new BigNumber((_836_v6).length));
      let _838_v8;
      _838_v8 = _dafny.Seq.of(p0, p0, new BigNumber((_837_v7).length));
      r0 = new BigNumber((_838_v8).length);
      let _839_v9;
      _839_v9 = _dafny.Set.fromElements(_830_v0);
      let _840_v10;
      _840_v10 = _dafny.Seq.of(_839_v9, _839_v9);
      let _841_v11;
      let _nw116 = new _module.C0();
      _nw116.__ctor();
      _841_v11 = _nw116;
      let _842_v12;
      _842_v12 = _dafny.Map.Empty.slice().updateUnsafe(_841_v11,(_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))]);
      let _843_v13;
      _843_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), function (_844_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length),(_dafny.Map.Empty.slice().updateUnsafe(_841_v11,(_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))])).Merge(_842_v12));
      let _rhs116 = p0;
      let _rhs117 = _840_v10;
      let _rhs118 = _830_v0;
      let _rhs119 = _843_v13;
      r0 = _rhs116;
      _840_v10 = _rhs117;
      _830_v0 = _rhs118;
      _843_v13 = _rhs119;
      if ((_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))]) {
        (globalState).f0 = !(new BigNumber(249)).isEqualTo(p0);
        if (p1.f3) {
          let _845_v14;
          _845_v14 = _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(p0)).length), p0), new BigNumber((_dafny.Seq.Concat(_836_v6, _dafny.Seq.of((_this).f11, p1.f3))).length));
          _845_v14 = _845_v14;
          let _846_v15;
          _846_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          _846_v15 = (_846_v15).update(p1, p0);
          let _847_v16;
          let _nw117 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _847_v16 = _nw117;
          let _index124 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_847_v16).length));
          (_847_v16)[_index124] = _831_v1;
          let _index125 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_847_v16).length));
          (_847_v16)[_index125] = _dafny.Seq.Concat(_831_v1, _dafny.Seq.Concat(_831_v1, _831_v1));
          let _848_v17;
          _848_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _849_v18;
          _849_v18 = _dafny.Seq.of(_848_v17, _848_v17);
          let _850_v19;
          _850_v19 = _dafny.Seq.of(_848_v17, (_849_v18)[_module.__default.safeIndex(new BigNumber((_836_v6).length), new BigNumber((_849_v18).length))]);
          let _851_v20;
          _851_v20 = _dafny.Map.Empty.slice().updateUnsafe(_850_v19,true);
          (p1).f3 = (((_851_v20).contains(_dafny.Seq.Concat(_850_v19, _849_v18))) ? ((_851_v20).get(_dafny.Seq.Concat(_850_v19, _849_v18))) : (((p1.f3) ? ((_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))]) : (p1.f3))));
          let _rhs120 = p0;
          let _rhs121 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ccdmppcba"), _831_v1), _831_v1);
          r0 = _rhs120;
          _831_v1 = _rhs121;
        } else {
          let _852_v21;
          _852_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-386));
          let _853_v22;
          _853_v22 = _dafny.MultiSet.fromElements(_852_v21, _852_v21);
          let _index126 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length));
          (_833_v3)[_index126] = !(_module.__default.fm21(p0, _853_v22, globalState));
          _830_v0 = _830_v0;
          (globalState).f0 = (p1).f4;
          _830_v0 = ((!(p1.f3)) ? (_830_v0) : (_830_v0));
          _838_v8 = _dafny.Seq.update(_dafny.Seq.Concat(_838_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), ((_854_p0) => function (_855_i1) {
            return _854_p0;
          })(p0))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_838_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), ((_856_p0) => function (_857_i1) {
            return _856_p0;
          })(p0)))).length)), (p0).minus(p0));
        }
        let _858_v23;
        _858_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        let _859_v24;
        _859_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_858_v23).length),p0);
        let _860_v25;
        _860_v25 = _dafny.MultiSet.fromElements(new BigNumber((_859_v24).length));
        _837_v7 = (_837_v7).update((p1).f4, (((_860_v25).contains((_dafny.ZERO).minus(p0))) ? ((_860_v25).get((_dafny.ZERO).minus(p0))) : ((_this).fm16(p0, (_this).f4, p0, (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_this.f3)).update(new BigNumber((_831_v1).length), (_this).f4), globalState))));
        let _index127 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length));
        (_833_v3)[_index127] = (_this).f11;
        (globalState).f0 = p1.f3;
      } else {
        let _861_v26;
        let _init17 = ((_862_p0) => function (_863_i2) {
          return (_863_i2).multipliedBy(_862_p0);
        })(p0);
        let _nw118 = Array((new BigNumber(28)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw118.length); _i0_17++) {
          _nw118[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _861_v26 = _nw118;
        let _864_v27;
        _864_v27 = _dafny.MultiSet.fromElements(_861_v26, _861_v26);
        _836_v6 = _dafny.Seq.of((_864_v27).IsProperSubsetOf(_864_v27));
        r0 = p0;
        _861_v26 = ((!((p1).f4)) ? (_861_v26) : (_861_v26));
        let _865_v28;
        _865_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,(_this).f10);
        _865_v28 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,!(!((_833_v3)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_833_v3).length))]) || (!((p1).f4))));
        (globalState).f0 = !(_this.f3);
      }
      let _866_v29;
      _866_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0),p0);
      let _867_v30;
      _867_v30 = _dafny.MultiSet.fromElements(p0, new BigNumber(682), p0, (((_866_v29).contains(_838_v8)) ? ((_866_v29).get(_838_v8)) : (p0)));
      r0 = (_841_v11).fm10(new BigNumber((_dafny.Seq.Concat(_831_v1, _831_v1)).length), true, (_867_v30).Intersect(_867_v30), globalState);
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _868_v0;
      _868_v0 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_this.f3));
      let _869_i0;
      _869_i0 = _dafny.ZERO;
      L3: {
        while ((_868_v0).IsDisjointFrom(_868_v0)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_869_i0)) {
              break L3;
            }
            _869_i0 = (_869_i0).plus(_dafny.ONE);
            (p1).f3 = p1.f3;
            let _870_v1;
            _870_v1 = _dafny.Map.Empty.slice().updateUnsafe(!((p1).f4),new BigNumber(-940));
            _870_v1 = _870_v1;
            let _871_v3;
            let _init18 = ((_872_p1) => function (_873_i1) {
              return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f4)).cardinality()),(_this).f10)).Merge(function () {
                let _coll33 = new _dafny.Map();
                for (const _compr_33 of (_dafny.Set.fromElements(new BigNumber(413))).Elements) {
                  let _874_v2 = _compr_33;
                  if ((_dafny.Set.fromElements(new BigNumber(413))).contains(_874_v2)) {
                    _coll33.push([(_874_v2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_872_p1.f3,(_872_p1).f4)).length)),true]);
                  }
                }
                return _coll33;
              }());
            })(p1);
            let _nw119 = Array((new BigNumber(27)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw119.length); _i0_18++) {
              _nw119[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _871_v3 = _nw119;
            let _875_v4;
            _875_v4 = _dafny.Seq.of(_871_v3);
            let _876_v5;
            _876_v5 = new BigNumber(395);
            _871_v3 = (_875_v4)[_module.__default.safeIndex((_876_v5).minus(_876_v5), new BigNumber((_875_v4).length))];
            let _877_v6;
            _877_v6 = _dafny.Seq.UnicodeFromString("o");
            let _878_v7;
            _878_v7 = _dafny.Map.Empty.slice().updateUnsafe(_876_v5,p1.f3);
            let _879_v8;
            _879_v8 = _dafny.Seq.of(_876_v5);
            let _880_v9;
            _880_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_879_v8);
            let _881_v10;
            _881_v10 = _module.D7.create_DC17(_876_v5, (_this).f4, (p1).f4, _876_v5);
            let _882_v11;
            _882_v11 = _dafny.Map.Empty.slice().updateUnsafe(_881_v10,(_this).f11);
            let _883_v12;
            let _nw120 = Array((new BigNumber(3)).toNumber());
            _nw120[(_dafny.ZERO).toNumber()] = _876_v5;
            _nw120[(_dafny.ONE).toNumber()] = new BigNumber((_882_v11).length);
            _nw120[(new BigNumber(2)).toNumber()] = _876_v5;
            _883_v12 = _nw120;
            let _884_v13;
            _884_v13 = _dafny.Seq.of(_883_v12, _883_v12);
            let _885_v15;
            _885_v15 = _dafny.MultiSet.fromElements(_876_v5);
            let _886_v16;
            _886_v16 = _dafny.Seq.of(_module.__default.fm21(_876_v5, _module.__default.fm22(_876_v5, (_this).fm16(_876_v5, p1.f3, _876_v5, function () {
              let _coll34 = new _dafny.Map();
              for (const _compr_34 of (_885_v15).Elements) {
                let _887_v14 = _compr_34;
                if ((_885_v15).contains(_887_v14)) {
                  _coll34.push([_module.__default.safeDivisionInt(_887_v14, _876_v5),p1.f3]);
                }
              }
              return _coll34;
            }(), globalState), _876_v5, globalState), globalState), (p1).f4, (_this).f11);
            let _888_v18;
            _888_v18 = _dafny.Map.Empty.slice().updateUnsafe(_876_v5,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-868)), ((_889_p0) => function (_890_i2) {
              return _889_p0;
            })(p0)));
            let _891_v19;
            _891_v19 = _dafny.Map.Empty.slice().updateUnsafe(true,_883_v12);
            let _892_v20;
            _892_v20 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_dafny.MultiSet.fromElements(_883_v12, (((_891_v19).contains((p1).f4)) ? ((_891_v19).get((p1).f4)) : (_883_v12))));
            let _893_v21;
            _893_v21 = _dafny.Map.Empty.slice().updateUnsafe(_876_v5,_883_v12);
            let _894_v22;
            _894_v22 = _dafny.MultiSet.fromElements((((_893_v21).contains(_876_v5)) ? ((_893_v21).get(_876_v5)) : (_883_v12)));
            let _895_v23;
            let _nw121 = Array((new BigNumber(27)).toNumber());
            _nw121[(_dafny.ZERO).toNumber()] = !(p1.f3);
            _nw121[(_dafny.ONE).toNumber()] = !(_this.f3);
            _nw121[(new BigNumber(2)).toNumber()] = (p1).f4;
            _nw121[(new BigNumber(3)).toNumber()] = (_this).f4;
            _nw121[(new BigNumber(4)).toNumber()] = !((p1).f4);
            _nw121[(new BigNumber(5)).toNumber()] = _this.f3;
            _nw121[(new BigNumber(6)).toNumber()] = false;
            _nw121[(new BigNumber(7)).toNumber()] = (p1).f4;
            _nw121[(new BigNumber(8)).toNumber()] = (_this).f4;
            _nw121[(new BigNumber(9)).toNumber()] = p1.f3;
            _nw121[(new BigNumber(10)).toNumber()] = !((p1).f4);
            _nw121[(new BigNumber(11)).toNumber()] = (_this).f11;
            _nw121[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_877_v6, _877_v6);
            _nw121[(new BigNumber(13)).toNumber()] = (((((_878_v7).contains(_876_v5)) ? ((_878_v7).get(_876_v5)) : ((_this).f4))) ? ((p1).f4) : (!(p1.f3)));
            _nw121[(new BigNumber(14)).toNumber()] = !_dafny.areEqual(_879_v8, (((_880_v9).contains(true)) ? ((_880_v9).get(true)) : (_dafny.Seq.of(_876_v5))));
            _nw121[(new BigNumber(15)).toNumber()] = (p1).f4;
            _nw121[(new BigNumber(16)).toNumber()] = !(_dafny.Seq.contains(_884_v13, _883_v12));
            _nw121[(new BigNumber(17)).toNumber()] = (_this).f4;
            _nw121[(new BigNumber(18)).toNumber()] = (p1).f4;
            _nw121[(new BigNumber(19)).toNumber()] = (p1).f4;
            _nw121[(new BigNumber(20)).toNumber()] = false;
            _nw121[(new BigNumber(21)).toNumber()] = (p1.f3) === ((_886_v16)[_module.__default.safeIndex(_876_v5, new BigNumber((_886_v16).length))]);
            _nw121[(new BigNumber(22)).toNumber()] = (function () {
              let _coll35 = new _dafny.Map();
              for (const _compr_35 of _dafny.IntegerRange(new BigNumber(118), new BigNumber(515))) {
                let _896_v17 = _compr_35;
                if (((new BigNumber(118)).isLessThanOrEqualTo(_896_v17)) && ((_896_v17).isLessThan(new BigNumber(515)))) {
                  _coll35.push([_module.__default.safeDivisionInt(_896_v17, new BigNumber(579)),_877_v6]);
                }
              }
              return _coll35;
            }()).equals(_888_v18);
            _nw121[(new BigNumber(23)).toNumber()] = (_dafny.MultiSet.fromElements(_876_v5)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_884_v13).length)));
            _nw121[(new BigNumber(24)).toNumber()] = (_this).f11;
            _nw121[(new BigNumber(25)).toNumber()] = ((((_892_v20).contains((p1).f4)) ? ((_892_v20).get((p1).f4)) : (_894_v22))).IsDisjointFrom(_894_v22);
            _nw121[(new BigNumber(26)).toNumber()] = (((_886_v16)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p1.f3, (_this).f11)).length), new BigNumber((_886_v16).length))]) ? (!((p1).f4)) : ((p1).f4));
            _895_v23 = _nw121;
            let _index128 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_895_v23).length));
            (_895_v23)[_index128] = _this.f3;
            let _897_v24;
            _897_v24 = _dafny.Map.Empty.slice().updateUnsafe(_895_v23,_dafny.Seq.UnicodeFromString("v"));
            let _898_v25;
            _898_v25 = _dafny.Map.Empty.slice().updateUnsafe(_876_v5,_876_v5);
            let _899_v26;
            _899_v26 = _dafny.MultiSet.fromElements(_898_v25, ((_898_v25).update(_876_v5, _876_v5)).update(_876_v5, _876_v5), _898_v25);
            let _index129 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_895_v23).length));
            (_895_v23)[_index129] = _module.__default.fm21(new BigNumber((_897_v24).length), _899_v26, globalState);
            let _900_v27;
            _900_v27 = _module.D0.create_DC2(p1.f3, _876_v5, new BigNumber((_877_v6).length));
            let _901_v28;
            _901_v28 = _dafny.Seq.of(_895_v23);
            let _index130 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_895_v23).length));
            let _index131 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_895_v23).length));
            let _rhs122 = (_900_v27).dtor_cf3;
            let _rhs123 = !((_this).f4) || (!_dafny.areEqual(_dafny.Seq.of(_895_v23), _901_v28));
            let _lhs81 = _895_v23;
            let _lhs82 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_895_v23).length));
            let _lhs83 = _895_v23;
            let _lhs84 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_895_v23).length));
            _lhs81[_lhs82] = _rhs122;
            _lhs83[_lhs84] = _rhs123;
          }
        }
      }
      let _902_v29;
      _902_v29 = new BigNumber(443);
      let _903_v30;
      _903_v30 = _dafny.Seq.UnicodeFromString("j");
      let _904_v31;
      _904_v31 = _dafny.Seq.of(new BigNumber((_903_v30).length), (_dafny.ZERO).minus(_902_v29));
      let _905_v32;
      _905_v32 = _dafny.Map.Empty.slice().updateUnsafe(_902_v29,(_904_v31)[_module.__default.safeIndex(_902_v29, new BigNumber((_904_v31).length))]);
      let _906_v33;
      _906_v33 = _dafny.MultiSet.fromElements(_905_v32);
      let _907_v34;
      _907_v34 = _dafny.Seq.of(true);
      let _908_v36;
      _908_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-61),(_this).f4);
      let _909_v37;
      let _910_v38;
      let _out20;
      let _out21;
      let _outcollector6 = (p1).m0((_module.__default.fm31(_module.__default.fm21(_902_v29, _906_v33, globalState), _902_v29, (_this).f10, _907_v34, globalState)).isLessThanOrEqualTo(_902_v29), (_module.D8.create_DC18(function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of (_908_v36).Keys.Elements) {
    let _911_v35 = _compr_36;
    if ((_908_v36).contains(_911_v35)) {
      _coll36.push([(_911_v35).plus(_902_v29),(p1).f4]);
    }
  }
  return _coll36;
}())).dtor_cf30, (new BigNumber(209)).plus(_902_v29), globalState);
      _out20 = _outcollector6[0];
      _out21 = _outcollector6[1];
      _909_v37 = _out20;
      _910_v38 = _out21;
      (p1).f3 = (_907_v34)[_module.__default.safeIndex(_module.__default.safeModuloInt(_910_v38, (_dafny.ZERO).minus(_902_v29)), new BigNumber((_907_v34).length))];
      _902_v29 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_903_v30).length), (_910_v38).multipliedBy(_902_v29)));
      let _912_v39;
      let _nw122 = new _module.C1();
      _nw122.__ctor(_903_v30, (_this).f10);
      _912_v39 = _nw122;
      let _913_v40;
      _913_v40 = new _dafny.CodePoint('y'.codePointAt(0));
      _913_v40 = p0;
      r0 = _903_v30;
      return r0;
    }
    m9(globalState) {
      let _this = this;
      let _914_v0;
      _914_v0 = new BigNumber(44);
      let _915_v1;
      _915_v1 = _dafny.Set.fromElements(_914_v0);
      let _916_v2;
      _916_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("htgmqesq")).length),_915_v1);
      let _917_v3;
      _917_v3 = _915_v1;
      if (((((_916_v2).contains(_914_v0)) ? ((_916_v2).get(_914_v0)) : (_915_v1))).IsDisjointFrom((_917_v3))) {
        let _918_v4;
        _918_v4 = _dafny.Seq.of((_this).f4);
        if (!((_918_v4)[_module.__default.safeIndex(_914_v0, new BigNumber((_918_v4).length))])) {
          let _919_v5;
          let _nw123 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _919_v5 = _nw123;
          let _920_v6;
          let _nw124 = Array((new BigNumber(9)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = _919_v5;
          _nw124[(_dafny.ONE).toNumber()] = _919_v5;
          _nw124[(new BigNumber(2)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(3)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(4)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(5)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(6)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(7)).toNumber()] = _919_v5;
          _nw124[(new BigNumber(8)).toNumber()] = _919_v5;
          _920_v6 = _nw124;
          let _index132 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_920_v6).length));
          (_920_v6)[_index132] = _919_v5;
          let _index133 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_920_v6).length));
          let _init19 = ((_921_v0) => function (_922_i0) {
            return (_922_i0).minus((_dafny.ZERO).minus(_921_v0));
          })(_914_v0);
          let _nw125 = Array((new BigNumber(23)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw125.length); _i0_19++) {
            _nw125[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          (_920_v6)[_index133] = _nw125;
          let _923_v7;
          _923_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_914_v0);
          let _924_v8;
          _924_v8 = _dafny.Set.fromElements(_923_v7);
          (globalState).f0 = (_924_v8).IsDisjointFrom(_924_v8);
          let _925_v9;
          _925_v9 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0);
          (_this).f3 = (_925_v9).equals(_925_v9);
          let _926_v10;
          let _nw126 = Array((new BigNumber(24)).toNumber()).fill(false);
          _926_v10 = _nw126;
          let _927_v11;
          _927_v11 = _dafny.Map.Empty.slice().updateUnsafe(_926_v10,_dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0));
          _927_v11 = (_927_v11).update(_926_v10, (_925_v9).Merge(_925_v9));
          (_this).f3 = !(_this.f3);
        } else {
          let _928_v12;
          _928_v12 = _dafny.Seq.UnicodeFromString("ldgivosl");
          let _929_v13;
          _929_v13 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0);
          let _930_v14;
          _930_v14 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,new BigNumber((_929_v13).length));
          let _931_v15;
          _931_v15 = _dafny.MultiSet.fromElements(_930_v14, _930_v14);
          let _932_v16;
          _932_v16 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_module.__default.fm21(new BigNumber((_928_v12).length), _931_v15, globalState));
          _932_v16 = (_932_v16).update((((_930_v14).contains(_914_v0)) ? ((_930_v14).get(_914_v0)) : (new BigNumber(360))), (_this).f4);
          let _933_v17;
          let _nw127 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _933_v17 = _nw127;
          let _index134 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length));
          (_933_v17)[_index134] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_928_v12).length)), _914_v0);
          let _934_v18;
          _934_v18 = _dafny.MultiSet.fromElements(_914_v0, _914_v0);
          let _index135 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length));
          (_933_v17)[_index135] = new BigNumber(((_934_v18).Union((_934_v18).Difference(_934_v18))).cardinality());
          let _935_v19;
          _935_v19 = new _dafny.CodePoint('u'.codePointAt(0));
          let _936_v20;
          _936_v20 = _module.D8.create_DC20(_935_v19, true);
          let _pat_let_tv33 = _935_v19;
          _936_v20 = function (_pat_let14_0) {
            return function (_937_dt__update__tmp_h0) {
              return function (_pat_let15_0) {
                return function (_938_dt__update_hcf33_h0) {
                  return _module.D8.create_DC20(_938_dt__update_hcf33_h0, (_937_dt__update__tmp_h0).dtor_cf34);
                }(_pat_let15_0);
              }(_pat_let_tv33);
            }(_pat_let14_0);
          }(_936_v20);
          let _939_v21;
          let _nw128 = new _module.C1();
          _nw128.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), ((_940_v19) => function (_941_i1) {
            return _940_v19;
          })(_935_v19)), _this.f3);
          _939_v21 = _nw128;
          _939_v21 = _939_v21;
          let _942_v22;
          let _init20 = ((_943_v19) => function (_944_i2) {
            return _943_v19;
          })(_935_v19);
          let _nw129 = Array((new BigNumber(25)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw129.length); _i0_20++) {
            _nw129[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _942_v22 = _nw129;
          let _945_v23;
          _945_v23 = _dafny.MultiSet.fromElements(_939_v21);
          let _index136 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_942_v22).length));
          (_942_v22)[_index136] = _module.__default.fm23((_939_v21).f13, new BigNumber(((_945_v23).update(_939_v21, _module.__default.abs((_this).fm16(_914_v0, (_939_v21).f13, _914_v0, _932_v16, globalState)))).cardinality()), globalState);
          let _946_v24;
          _946_v24 = _module.D0.create_DC2(_this.f3, (_dafny.ZERO).minus(_914_v0), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_914_v0))).length));
          let _947_v25;
          _947_v25 = _dafny.Seq.of(_dafny.Seq.Concat(_928_v12, (_939_v21).f12), (((_946_v24).dtor_cf3) ? (_928_v12) : (_dafny.Seq.UnicodeFromString("yvkiceyku"))));
          let _948_v26;
          let _nw130 = new _module.C0();
          _nw130.__ctor();
          _948_v26 = _nw130;
          let _949_v27;
          _949_v27 = _dafny.MultiSet.fromElements(_948_v26, _948_v26);
          let _950_v28;
          _950_v28 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_918_v4);
          let _951_v29;
          _951_v29 = _dafny.Set.fromElements((_this).f11, (_this).f4);
          let _index137 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length));
          let _index138 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_942_v22).length));
          let _rhs124 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm31((_949_v27).IsProperSubsetOf(_949_v27), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))])).update(_this.f3, (_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))])).length), (_this).f4, _918_v4, globalState)));
          let _rhs125 = _935_v19;
          let _rhs126 = (_module.D10.create_DC22(_933_v17)).dtor_cf36;
          let _rhs127 = (_948_v26).fm11(_935_v19, _module.__default.fm33((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))], _dafny.Seq.update((((_950_v28).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_939_v21).f13)).length))) ? ((_950_v28).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_939_v21).f13)).length))) : (_module.__default.fm34(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_this).f11)).length), new BigNumber((_951_v29).length), globalState))), _module.__default.safeIndex(new BigNumber(151), new BigNumber(((((_950_v28).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_939_v21).f13)).length))) ? ((_950_v28).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_939_v21).f13)).length))) : (_module.__default.fm34(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_933_v17)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length))],(_this).f11)).length), new BigNumber((_951_v29).length), globalState)))).length)), (_939_v21).f13), globalState), globalState);
          let _rhs128 = _module.__default.fm35(!(_this.f3), globalState);
          let _lhs85 = _933_v17;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_933_v17).length));
          let _lhs87 = _942_v22;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_942_v22).length));
          _lhs85[_lhs86] = _rhs124;
          _lhs87[_lhs88] = _rhs125;
          _933_v17 = _rhs126;
          _928_v12 = _rhs127;
          _947_v25 = _rhs128;
        }
        let _952_v30;
        _952_v30 = _dafny.Seq.UnicodeFromString("px");
        _952_v30 = _952_v30;
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_952_v30, _952_v30), _dafny.Seq.UnicodeFromString("edlbvgnpp"))) {
          let _953_v31;
          _953_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,true);
          let _954_v32;
          let _nw131 = new _module.C1();
          _nw131.__ctor(_952_v30, (((_953_v31).contains(_this.f3)) ? ((_953_v31).get(_this.f3)) : (_this.f3)));
          _954_v32 = _nw131;
          let _955_v33;
          _955_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_915_v1);
          (globalState).f0 = (((_this).f4) ? (_this.f3) : (((((_955_v33).contains((_954_v32).f13)) ? ((_955_v33).get((_954_v32).f13)) : (_915_v1))).contains(new BigNumber((_915_v1).length))));
          let _956_v34;
          _956_v34 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0);
          let _957_v36;
          _957_v36 = _dafny.Seq.of(_914_v0, _914_v0);
          let _958_v37;
          _958_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(346),_module.__default.fm25(globalState));
          let _959_v39;
          let _nw132 = Array((new BigNumber(20)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = _956_v34;
          _nw132[(_dafny.ONE).toNumber()] = _956_v34;
          _nw132[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0);
          _nw132[(new BigNumber(3)).toNumber()] = function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_957_v36).Elements) {
              let _960_v35 = _compr_37;
              if (_dafny.Seq.contains(_957_v36, _960_v35)) {
                _coll37.push([_module.__default.safeModuloInt(_960_v35, _914_v0),_914_v0]);
              }
            }
            return _coll37;
          }();
          _nw132[(new BigNumber(4)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(5)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(6)).toNumber()] = (_956_v34).Merge(_dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0));
          _nw132[(new BigNumber(7)).toNumber()] = (_module.D11.create_DC24(_956_v34)).dtor_cf37;
          _nw132[(new BigNumber(8)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(9)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(10)).toNumber()] = (_956_v34).Merge(_956_v34);
          _nw132[(new BigNumber(11)).toNumber()] = _module.__default.fm25(globalState);
          _nw132[(new BigNumber(12)).toNumber()] = (_module.__default.fm25(globalState)).Merge((((_958_v37).contains(_914_v0)) ? ((_958_v37).get(_914_v0)) : (_956_v34)));
          _nw132[(new BigNumber(13)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(14)).toNumber()] = (_module.__default.fm25(globalState)).Merge(function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(924), _dafny.ZERO)) {
              let _961_v38 = _compr_38;
              if (((new BigNumber(924)).isLessThanOrEqualTo(_961_v38)) && ((_961_v38).isLessThan(_dafny.ZERO))) {
                _coll38.push([_module.__default.safeDivisionInt(_961_v38, new BigNumber(((_954_v32).f12).length)),_914_v0]);
              }
            }
            return _coll38;
          }());
          _nw132[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_914_v0);
          _nw132[(new BigNumber(16)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(17)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(18)).toNumber()] = _956_v34;
          _nw132[(new BigNumber(19)).toNumber()] = _956_v34;
          _959_v39 = _nw132;
          _959_v39 = _959_v39;
          _952_v30 = (_954_v32).f12;
          let _962_v40;
          _962_v40 = _module.D10.create_DC23();
          _962_v40 = _962_v40;
        } else {
          let _963_v41;
          _963_v41 = _module.D0.create_DC2((_this).f4, _914_v0, _914_v0);
          _914_v0 = (_963_v41).dtor_cf4;
          let _964_v42;
          let _nw133 = new _module.C0();
          _nw133.__ctor();
          _964_v42 = _nw133;
          let _965_v43;
          _965_v43 = new _dafny.CodePoint('g'.codePointAt(0));
          let _966_v44;
          _966_v44 = _dafny.MultiSet.fromElements(_this.f3);
          let _967_v45;
          _967_v45 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).Union(_dafny.Set.fromElements(_965_v43, _965_v43)),(_966_v44).Union(_966_v44));
          _967_v45 = _967_v45;
          let _968_v46;
          let _init21 = function (_969_i3) {
            return !((_this).f11);
          };
          let _nw134 = Array((new BigNumber(18)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw134.length); _i0_21++) {
            _nw134[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _968_v46 = _nw134;
          let _970_v47;
          _970_v47 = _dafny.MultiSet.fromElements(_968_v46);
          (globalState).f0 = (_970_v47).equals(_970_v47);
          let _971_v48;
          let _nw135 = new _module.C0();
          _nw135.__ctor();
          _971_v48 = _nw135;
        }
        let _972_v49;
        _972_v49 = _module.D4.create_DC12(_914_v0, new BigNumber(950), (_this).f4);
        let _pat_let_tv34 = _914_v0;
        (globalState).f0 = ((_this).f10) || (((_this).f10) === (!((function (_pat_let16_0) {
          return function (_973_dt__update__tmp_h1) {
            return function (_pat_let17_0) {
              return function (_974_dt__update_hcf20_h0) {
                return function (_pat_let18_0) {
                  return function (_975_dt__update_hcf19_h0) {
                    return _module.D4.create_DC12(_975_dt__update_hcf19_h0, _974_dt__update_hcf20_h0, (_973_dt__update__tmp_h1).dtor_cf21);
                  }(_pat_let18_0);
                }(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).length));
              }(_pat_let17_0);
            }(_pat_let_tv34);
          }(_pat_let16_0);
        }(_972_v49)).dtor_cf21)));
        let _976_v50;
        let _nw136 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _976_v50 = _nw136;
        let _index139 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_976_v50).length));
        (_976_v50)[_index139] = new BigNumber((_915_v1).length);
        let _977_v51;
        _977_v51 = _dafny.MultiSet.fromElements(new BigNumber((_916_v2).length));
        let _index140 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_976_v50).length));
        (_976_v50)[_index140] = (_module.__default.safeModuloInt(new BigNumber((_977_v51).cardinality()), new BigNumber(547))).minus(_914_v0);
      } else {
        let _978_v52;
        let _nw137 = new _module.C0();
        _nw137.__ctor();
        _978_v52 = _nw137;
        let _979_v53;
        _979_v53 = _dafny.Seq.UnicodeFromString("ptxas");
        let _980_v54;
        _980_v54 = _dafny.Seq.of(_979_v53, _979_v53);
        let _981_v55;
        _981_v55 = _dafny.MultiSet.fromElements(_979_v53, _979_v53, _979_v53, _dafny.Seq.update(_979_v53, _module.__default.safeIndex(_914_v0, new BigNumber((_979_v53).length)), new _dafny.CodePoint('w'.codePointAt(0))), _979_v53);
        let _982_v56;
        _982_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_979_v53, _module.__default.safeIndex(new BigNumber((_980_v54).length), new BigNumber((_979_v53).length)), new _dafny.CodePoint('j'.codePointAt(0))),new BigNumber((_981_v55).cardinality()));
        _982_v56 = (_982_v56).update(_979_v53, _914_v0);
        (globalState).f0 = (_this).f11;
        let _983_v57;
        _983_v57 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,_this.f3);
        let _984_v58;
        _984_v58 = _dafny.Seq.of((_this).f4, (_this).f11, (_this).f10, (_this).f10);
        let _985_v59;
        _985_v59 = _module.D1.create_DC4(new BigNumber((_979_v53).length));
        let _986_v60;
        _986_v60 = _dafny.Set.fromElements(!((_this).f11));
        let _987_v61;
        _987_v61 = _module.D12.create_DC26(_986_v60);
        let _988_v62;
        _988_v62 = _dafny.Seq.of(_914_v0);
        let _989_v64;
        _989_v64 = _dafny.Map.Empty.slice().updateUnsafe(_914_v0,new BigNumber((_979_v53).length));
        let _990_v65;
        _990_v65 = _dafny.MultiSet.fromElements(_989_v64);
        let _991_v66;
        let _nw138 = Array((new BigNumber(22)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = (((_983_v57).contains(new BigNumber((_979_v53).length))) ? ((_983_v57).get(new BigNumber((_979_v53).length))) : ((_this).f4));
        _nw138[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_984_v58, !(_this.f3));
        _nw138[(new BigNumber(2)).toNumber()] = _dafny.Seq.contains(_984_v58, false);
        _nw138[(new BigNumber(3)).toNumber()] = !_dafny.Seq.contains(_984_v58, _this.f3);
        _nw138[(new BigNumber(4)).toNumber()] = (_this).f4;
        _nw138[(new BigNumber(5)).toNumber()] = true;
        _nw138[(new BigNumber(6)).toNumber()] = !((_985_v59).dtor_cf7).isEqualTo(_914_v0);
        _nw138[(new BigNumber(7)).toNumber()] = (false) === ((_this).f10);
        _nw138[(new BigNumber(8)).toNumber()] = !((_dafny.Set.fromElements(false, _this.f3, (_this).f4, _this.f3, (_this).f11)).IsProperSubsetOf((_987_v61).dtor_cf39));
        _nw138[(new BigNumber(9)).toNumber()] = (_this).f11;
        _nw138[(new BigNumber(10)).toNumber()] = (((_this).f11) ? (_this.f3) : ((_this).f10));
        _nw138[(new BigNumber(11)).toNumber()] = (((_this).f4) ? (!((_this).f10)) : (false));
        _nw138[(new BigNumber(12)).toNumber()] = ((_dafny.ZERO).minus((_988_v62)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_988_v62).length))])).isLessThanOrEqualTo(new BigNumber(233));
        _nw138[(new BigNumber(13)).toNumber()] = !((function () {
          let _coll39 = new _dafny.Set();
          for (const _compr_39 of _dafny.IntegerRange(new BigNumber(403), new BigNumber(-422))) {
            let _992_v63 = _compr_39;
            if (((new BigNumber(403)).isLessThanOrEqualTo(_992_v63)) && ((_992_v63).isLessThan(new BigNumber(-422)))) {
              _coll39.add((_992_v63).multipliedBy(_914_v0));
            }
          }
          return _coll39;
        }()).IsDisjointFrom(_915_v1));
        _nw138[(new BigNumber(14)).toNumber()] = (_this).f10;
        _nw138[(new BigNumber(15)).toNumber()] = (((_this).f4) ? (!((_this).f4)) : ((_this).f10));
        _nw138[(new BigNumber(16)).toNumber()] = (_this).f10;
        _nw138[(new BigNumber(17)).toNumber()] = (_this).f11;
        _nw138[(new BigNumber(18)).toNumber()] = (_914_v0).isLessThanOrEqualTo(_914_v0);
        _nw138[(new BigNumber(19)).toNumber()] = (_dafny.Set.fromElements(_module.__default.fm21(new BigNumber(-514), _990_v65, globalState))).IsSubsetOf(_dafny.Set.fromElements((_this).f11));
        _nw138[(new BigNumber(20)).toNumber()] = (_this).f10;
        _nw138[(new BigNumber(21)).toNumber()] = !_dafny.areEqual(_dafny.Seq.of((_this).f11), _dafny.Seq.of((_this).f4, (_this).f4));
        _991_v66 = _nw138;
        let _index141 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_991_v66).length));
        (_991_v66)[_index141] = !((_this).f4);
        let _index142 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_991_v66).length));
        (_991_v66)[_index142] = (_this).f10;
        let _993_v67;
        let _nw139 = Array((new BigNumber(18)).toNumber()).fill(_module.D2.Default());
        _993_v67 = _nw139;
        let _index143 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_993_v67).length));
        (_993_v67)[_index143] = _module.D2.create_DC8(_991_v66, _dafny.Seq.UnicodeFromString("iuu"));
        let _994_v68;
        _994_v68 = new _dafny.CodePoint('v'.codePointAt(0));
        let _995_v69;
        _995_v69 = _module.D2.create_DC8(_991_v66, (_978_v52).fm11(_994_v68, _986_v60, globalState));
        let _index144 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_993_v67).length));
        (_993_v67)[_index144] = _995_v69;
      }
      let _hi5 = (_this).fm0(globalState);
      for (let _996_i4 = _914_v0; _996_i4.isLessThan(_hi5); _996_i4 = _996_i4.plus(_dafny.ONE)) {
        let _997_v70;
        let _init22 = function (_998_i5) {
          return (_this).f10;
        };
        let _nw140 = Array((new BigNumber(25)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw140.length); _i0_22++) {
          _nw140[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _997_v70 = _nw140;
        let _999_v71;
        _999_v71 = _dafny.Set.fromElements(_997_v70);
        let _rhs129 = true;
        let _rhs130 = (_999_v71).Union(_999_v71);
        let _lhs89 = globalState;
        _lhs89.f0 = _rhs129;
        _999_v71 = _rhs130;
        _914_v0 = (_996_i4).multipliedBy(_914_v0);
        let _index145 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_997_v70).length));
        (_997_v70)[_index145] = (_this).f10;
        let _1000_v72;
        _1000_v72 = _dafny.Seq.UnicodeFromString("jhy");
        let _index146 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_997_v70).length));
        (_997_v70)[_index146] = (((_this).f4) ? ((_this).f11) : (((_dafny.ZERO).minus(_996_i4)).isLessThanOrEqualTo(new BigNumber((_1000_v72).length))));
        let _index147 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_997_v70).length));
        (_997_v70)[_index147] = (_997_v70)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_997_v70).length))];
      }
      (globalState).f0 = (_this).f4;
      _914_v0 = _module.__default.safeDivisionInt(_914_v0, _914_v0);
      _914_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_914_v0).minus(_914_v0), _914_v0));
      let _1001_v73;
      _1001_v73 = _dafny.Seq.UnicodeFromString("hd");
      let _1002_v74;
      let _nw141 = Array((new BigNumber(7)).toNumber());
      _nw141[(_dafny.ZERO).toNumber()] = _914_v0;
      _nw141[(_dafny.ONE).toNumber()] = new BigNumber(989);
      _nw141[(new BigNumber(2)).toNumber()] = new BigNumber((_1001_v73).length);
      _nw141[(new BigNumber(3)).toNumber()] = _914_v0;
      _nw141[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_914_v0, _914_v0);
      _nw141[(new BigNumber(5)).toNumber()] = _914_v0;
      _nw141[(new BigNumber(6)).toNumber()] = _914_v0;
      _1002_v74 = _nw141;
      let _index148 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1002_v74).length));
      (_1002_v74)[_index148] = _module.__default.safeModuloInt(_914_v0, _914_v0);
      let _index149 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_1002_v74).length));
      (_1002_v74)[_index149] = (_dafny.ZERO).minus(_914_v0);
      return;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f3 = false;
      this._f4 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f3, f4) {
      let _this = this;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, _this.f3)).cardinality())),new BigNumber(314))).length),_this.f3),_this.f3)).length)).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gett")).length))).length), new BigNumber(850), new BigNumber((_dafny.Seq.of(false)).length))).length));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements((_this).f4)).Intersect((_dafny.Set.fromElements(_this.f3)).Union(_dafny.Set.fromElements(_this.f3)));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1003_v0;
      _1003_v0 = _module.D0.create_DC0(p0);
      let _1004_v1;
      _1004_v1 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1005_v2;
      _1005_v2 = _dafny.MultiSet.fromElements(_1004_v1);
      let _1006_v3;
      _1006_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1005_v2).cardinality()),p0);
      let _1007_v4;
      _1007_v4 = _dafny.MultiSet.fromElements((_this).f4);
      let _1008_v5;
      _1008_v5 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_module.__default.fm8(new BigNumber((_1007_v4).cardinality()), globalState));
      let _1009_v6;
      _1009_v6 = _dafny.Seq.of(_this.f3, (_this).f4, (_this).f4);
      let _1010_v7;
      let _nw142 = Array((new BigNumber(28)).toNumber());
      _nw142[(_dafny.ZERO).toNumber()] = _1004_v1;
      _nw142[(_dafny.ONE).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(2)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(3)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(4)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(5)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(6)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(7)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(8)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
      _nw142[(new BigNumber(10)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
      _nw142[(new BigNumber(12)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(13)).toNumber()] = _module.__default.fm8(new BigNumber(589), globalState);
      _nw142[(new BigNumber(14)).toNumber()] = _module.__default.fm8(p0, globalState);
      _nw142[(new BigNumber(15)).toNumber()] = (((_1008_v5).contains((_this).f4)) ? ((_1008_v5).get((_this).f4)) : (_1004_v1));
      _nw142[(new BigNumber(16)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(17)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(18)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
      _nw142[(new BigNumber(20)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
      _nw142[(new BigNumber(22)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(23)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(24)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(25)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(26)).toNumber()] = _1004_v1;
      _nw142[(new BigNumber(27)).toNumber()] = (_module.__default.fm9(p0, _1009_v6, p1.f3, p0, globalState)).dtor_cf2;
      _1010_v7 = _nw142;
      let _1011_v8;
      _1011_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_1006_v3).contains(new BigNumber(816))) ? ((_1006_v3).get(new BigNumber(816))) : (p0)),_1010_v7);
      let _hi6 = new BigNumber((_1011_v8).length);
      for (let _1012_i0 = ((_1003_v0).dtor_cf0).multipliedBy((_this).fm0(globalState)); _1012_i0.isLessThan(_hi6); _1012_i0 = _1012_i0.plus(_dafny.ONE)) {
        r0 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).multipliedBy(p0);
        r0 = p0;
        (globalState).f0 = false;
        let _1013_v9;
        _1013_v9 = _dafny.Seq.of(_1004_v1, _1004_v1);
        let _1014_v10;
        _1014_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1013_v9,(p1).f4);
        let _1015_v11;
        let _nw143 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1015_v11 = _nw143;
        let _1016_v12;
        _1016_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1015_v11,_1012_i0);
        let _1017_v13;
        _1017_v13 = _dafny.Seq.of(p0, new BigNumber((_1016_v12).length));
        if (((_1014_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), ((_1018_v1) => function (_1019_i1) {
          return _1018_v1;
        })(_1004_v1)),(p1).f4))) || ((new BigNumber((_1013_v9).length)).isLessThanOrEqualTo((_1017_v13)[_module.__default.safeIndex((_dafny.ZERO).minus(_1012_i0), new BigNumber((_1017_v13).length))]))) {
          let _1020_v14;
          let _init23 = ((_1021_p1) => function (_1022_i2) {
            return (_1021_p1).f4;
          })(p1);
          let _nw144 = Array((new BigNumber(19)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw144.length); _i0_23++) {
            _nw144[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1020_v14 = _nw144;
          _1020_v14 = _1020_v14;
          let _1023_v15;
          _1023_v15 = _module.D0.create_DC2(((false) ? (p1.f3) : (_this.f3)), new BigNumber((_1007_v4).cardinality()), p0);
          _1023_v15 = _module.D0.create_DC2((p1).f4, _1012_i0, _module.__default.safeDivisionInt(p0, p0));
          _1023_v15 = _1023_v15;
          let _1024_v16;
          _1024_v16 = _dafny.Seq.of(_1003_v0);
          let _1025_v17;
          _1025_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1012_i0,_1024_v16);
          let _1026_v19;
          _1026_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
          let _rhs131 = p0;
          let _rhs132 = (((p1).f4) ? (function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_1026_v19).Keys.Elements) {
              let _1027_v18 = _compr_40;
              if ((_1026_v19).contains(_1027_v18)) {
                _coll40.push([(_1027_v18).multipliedBy(p0),_1024_v16]);
              }
            }
            return _coll40;
          }()) : ((_1025_v17).Merge(_1025_v17)));
          r0 = _rhs131;
          _1025_v17 = _rhs132;
          let _1028_v20;
          _1028_v20 = _dafny.MultiSet.fromElements(p0);
          let _index150 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1015_v11).length));
          (_1015_v11)[_index150] = (_dafny.ZERO).minus((((_1028_v20).contains(p0)) ? ((_1028_v20).get(p0)) : (p0)));
          let _index151 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1015_v11).length));
          (_1015_v11)[_index151] = (_1023_v15).dtor_cf5;
        } else {
          let _1029_v21;
          let _nw145 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1029_v21 = _nw145;
          let _1030_v22;
          let _nw146 = Array((new BigNumber(24)).toNumber()).fill(false);
          _1030_v22 = _nw146;
          let _index152 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1029_v21).length));
          (_1029_v21)[_index152] = _1030_v22;
          let _1031_v23;
          _1031_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_1012_i0);
          let _1032_v24;
          _1032_v24 = _dafny.MultiSet.fromElements(_1031_v23);
          let _index153 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1029_v21).length));
          let _rhs133 = true;
          let _rhs134 = _1030_v22;
          let _rhs135 = (_1032_v24).IsProperSubsetOf(_1032_v24);
          let _rhs136 = (p1).f4;
          let _rhs137 = _1030_v22;
          let _lhs90 = p1;
          let _lhs91 = _1029_v21;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1029_v21).length));
          let _lhs93 = p1;
          let _lhs94 = globalState;
          _lhs90.f3 = _rhs133;
          _lhs91[_lhs92] = _rhs134;
          _lhs93.f3 = _rhs135;
          _lhs94.f0 = _rhs136;
          _1030_v22 = _rhs137;
          let _1033_v25;
          _1033_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(p1).f4);
          (globalState).f0 = ((new BigNumber((_1009_v6).length)).plus(new BigNumber((_1033_v25).length))).isLessThan((_1012_i0).plus(_1012_i0));
          _1014_v10 = (_1014_v10).update(_dafny.Seq.Concat(_1013_v9, _1013_v9), ((p1).f4) === (p1.f3));
          let _1034_v26;
          _1034_v26 = _dafny.Set.fromElements(new BigNumber((_1017_v13).length));
          (globalState).f0 = !((new BigNumber((_1013_v9).length)).plus(new BigNumber((_1034_v26).length))).isEqualTo((p1).fm0(globalState));
          r0 = _1012_i0;
        }
      }
      if ((p1).f4) {
        let _1035_v27;
        let _nw147 = Array((new BigNumber(9)).toNumber()).fill(false);
        _1035_v27 = _nw147;
        let _1036_v28;
        _1036_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,p0);
        let _1037_v29;
        _1037_v29 = _dafny.Set.fromElements(_1036_v28);
        let _1038_v31;
        _1038_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1037_v29).Difference(function () {
          let _coll41 = new _dafny.Set();
          for (const _compr_41 of (_dafny.Seq.of(_1036_v28, _1036_v28, _1036_v28)).Elements) {
            let _1039_v30 = _compr_41;
            if (_dafny.Seq.contains(_dafny.Seq.of(_1036_v28, _1036_v28, _1036_v28), _1039_v30)) {
              _coll41.add(_1039_v30);
            }
          }
          return _coll41;
        }()),_1035_v27);
        _1035_v27 = (((_1038_v31).contains((_1037_v29).Intersect(_1037_v29))) ? ((_1038_v31).get((_1037_v29).Intersect(_1037_v29))) : (_1035_v27));
        let _1040_v32;
        _1040_v32 = _dafny.Set.fromElements(new BigNumber(470), p0);
        let _rhs138 = (_module.D0.create_DC0(p0)).dtor_cf0;
        let _rhs139 = _1004_v1;
        let _rhs140 = (_1040_v32).IsProperSubsetOf(_1040_v32);
        let _rhs141 = (_dafny.ZERO).minus(p0);
        let _lhs95 = p1;
        r0 = _rhs138;
        _1004_v1 = _rhs139;
        _lhs95.f3 = _rhs140;
        r0 = _rhs141;
        if ((_this).f4) {
          let _1041_v33;
          let _init24 = ((_1042_p0) => function (_1043_i3) {
            return (_1043_i3).minus(_1042_p0);
          })(p0);
          let _nw148 = Array((new BigNumber(23)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw148.length); _i0_24++) {
            _nw148[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1041_v33 = _nw148;
          let _1044_v34;
          _1044_v34 = _dafny.Seq.UnicodeFromString("yowfhhl");
          let _index154 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length));
          (_1041_v33)[_index154] = new BigNumber((_1044_v34).length);
          let _pat_let_tv35 = p0;
          let _1045_v35;
          _1045_v35 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,(function (_pat_let19_0) {
            return function (_1046_dt__update__tmp_h0) {
              return function (_pat_let20_0) {
                return function (_1047_dt__update_hcf4_h0) {
                  return _module.D0.create_DC2((_1046_dt__update__tmp_h0).dtor_cf3, _1047_dt__update_hcf4_h0, (_1046_dt__update__tmp_h0).dtor_cf5);
                }(_pat_let20_0);
              }((_module.D0.create_DC0(_pat_let_tv35)).dtor_cf0);
            }(_pat_let19_0);
          }(_module.D0.create_DC2((_this).f4, p0, p0))).dtor_cf3);
          let _1048_v36;
          _1048_v36 = _dafny.Seq.of(_1045_v35);
          let _index155 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length));
          let _rhs142 = true;
          let _rhs143 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1048_v36, _1048_v36), _dafny.Seq.update(_dafny.Seq.update(_1048_v36, _module.__default.safeIndex(new BigNumber(890), new BigNumber((_1048_v36).length)), _1045_v35), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_1048_v36, _module.__default.safeIndex(new BigNumber(890), new BigNumber((_1048_v36).length)), _1045_v35)).length)), _1045_v35))).length));
          let _lhs96 = globalState;
          let _lhs97 = _1041_v33;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length));
          _lhs96.f0 = _rhs142;
          _lhs97[_lhs98] = _rhs143;
          let _1049_v37;
          _1049_v37 = _module.D0.create_DC2(!_dafny.Seq.contains(_1044_v34, new _dafny.CodePoint('u'.codePointAt(0))), new BigNumber(-262), (new BigNumber((_1007_v4).cardinality())).plus(new BigNumber((_dafny.Seq.of(_1044_v34)).length)));
          let _1050_v38;
          _1050_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(90),p1.f3);
          _1049_v37 = _module.D0.create_DC2((((_1050_v38).contains(new BigNumber(449))) ? ((_1050_v38).get(new BigNumber(449))) : ((p1).f4)), ((_1041_v33)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length))]).plus(new BigNumber((_1044_v34).length)), (_this).fm0(globalState));
          let _1051_v39;
          _1051_v39 = _dafny.Set.fromElements((p1).f4);
          let _1052_v40;
          _1052_v40 = _dafny.MultiSet.fromElements(_1051_v39);
          _1052_v40 = _dafny.MultiSet.fromElements((_1051_v39).Intersect(_1051_v39), _1051_v39, _1051_v39, _1051_v39);
          let _1053_v41;
          _1053_v41 = _dafny.Map.Empty.slice().updateUnsafe((((_1007_v4).contains(p1.f3)) ? ((_1007_v4).get(p1.f3)) : ((_1041_v33)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length))])),_1004_v1);
          let _1054_v42;
          _1054_v42 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(374)).plus((_1041_v33)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_1041_v33).length))]),_1053_v41);
          _1054_v42 = _1054_v42;
          let _1055_v43;
          let _nw149 = new _module.C0();
          _nw149.__ctor();
          _1055_v43 = _nw149;
        } else {
          let _1056_v44;
          _1056_v44 = _dafny.Seq.of(_1035_v27, _1035_v27);
          let _1057_v45;
          _1057_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1035_v27);
          let _1058_v46;
          let _nw150 = Array((new BigNumber(25)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _1035_v27;
          _nw150[(_dafny.ONE).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(2)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(3)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(4)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(5)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(6)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(7)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(8)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(9)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(10)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(11)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(12)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(13)).toNumber()] = (_1056_v44)[_module.__default.safeIndex(p0, new BigNumber((_1056_v44).length))];
          _nw150[(new BigNumber(14)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(15)).toNumber()] = (((_1057_v45).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("juonh"),(p1).f4)).length))) ? ((_1057_v45).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("juonh"),(p1).f4)).length))) : (_1035_v27));
          _nw150[(new BigNumber(16)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(17)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(18)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(19)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(20)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(21)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(22)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(23)).toNumber()] = _1035_v27;
          _nw150[(new BigNumber(24)).toNumber()] = _1035_v27;
          _1058_v46 = _nw150;
          let _index156 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1058_v46).length));
          (_1058_v46)[_index156] = _1035_v27;
          let _1059_v47;
          _1059_v47 = _dafny.Seq.UnicodeFromString("mstiflvo");
          let _1060_v48;
          _1060_v48 = _dafny.Set.fromElements(_1059_v47, _1059_v47, _1059_v47);
          let _1061_v49;
          _1061_v49 = _dafny.Set.fromElements((p1).f4);
          let _1062_v50;
          _1062_v50 = _dafny.Seq.of(_1061_v49);
          let _1063_v51;
          _1063_v51 = _dafny.Seq.of(p0, p0, new BigNumber((_1062_v50).length), _module.__default.safeDivisionInt(p0, p0));
          let _index157 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1058_v46).length));
          let _rhs144 = _1035_v27;
          let _rhs145 = (p0).multipliedBy(new BigNumber(705));
          let _rhs146 = _1060_v48;
          let _rhs147 = new BigNumber((_1063_v51).length);
          let _lhs99 = _1058_v46;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1058_v46).length));
          _lhs99[_lhs100] = _rhs144;
          r0 = _rhs145;
          _1060_v48 = _rhs146;
          r0 = _rhs147;
          let _1064_v52;
          _1064_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1059_v47,p0);
          let _1065_v55;
          let _init25 = ((_1066_p1, _1067_v2, _1068_p0) => function (_1069_i4) {
            return ((_1066_p1.f3) ? (function () {
              let _coll42 = new _dafny.Map();
              for (const _compr_42 of (_1067_v2).Elements) {
                let _1070_v53 = _compr_42;
                if ((_1067_v2).contains(_1070_v53)) {
                  _coll42.push([_1070_v53,_1068_p0]);
                }
              }
              return _coll42;
            }()) : (function () {
              let _coll43 = new _dafny.Map();
              for (const _compr_43 of (_1067_v2).Elements) {
                let _1071_v54 = _compr_43;
                if ((_1067_v2).contains(_1071_v54)) {
                  _coll43.push([_1071_v54,_1068_p0]);
                }
              }
              return _coll43;
            }()));
          })(p1, _1005_v2, p0);
          let _nw151 = Array((new BigNumber(18)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw151.length); _i0_25++) {
            _nw151[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1065_v55 = _nw151;
          let _1072_v56;
          _1072_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1004_v1,new BigNumber((_dafny.Seq.of(p0, p0)).length));
          let _index158 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_1065_v55).length));
          (_1065_v55)[_index158] = _1072_v56;
          let _1073_v57;
          _1073_v57 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_1059_v47);
          let _1074_v58;
          _1074_v58 = _dafny.MultiSet.fromElements(new BigNumber((_1009_v6).length));
          let _index159 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_1065_v55).length));
          let _rhs148 = _1064_v52;
          let _rhs149 = (_dafny.Map.Empty.slice().updateUnsafe(_1004_v1,p0)).update(_1004_v1, p0);
          let _rhs150 = (_dafny.ZERO).minus((((_1074_v58).contains(p0)) ? ((_1074_v58).get(p0)) : (p0)));
          let _rhs151 = (new BigNumber(((((p1).f4) ? (_1006_v3) : (_1006_v3))).length)).plus((_this).fm0(globalState));
          let _rhs152 = ((p1.f3) ? (_1073_v57) : ((_1073_v57).Merge(_1073_v57)));
          let _lhs101 = _1065_v55;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_1065_v55).length));
          _1064_v52 = _rhs148;
          _lhs101[_lhs102] = _rhs149;
          r0 = _rhs150;
          r0 = _rhs151;
          _1073_v57 = _rhs152;
          r0 = new BigNumber(385);
          _1036_v28 = (_1036_v28).update((_this).f4, (((_this).f4) ? (p0) : ((_1063_v51)[_module.__default.safeIndex(p0, new BigNumber((_1063_v51).length))])));
          (globalState).f0 = _this.f3;
        }
        let _1075_v59;
        _1075_v59 = _module.D0.create_DC1((_this).f4, new _dafny.CodePoint('t'.codePointAt(0)));
        let _1076_v60;
        _1076_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1075_v59);
        _1076_v60 = _1076_v60;
        let _1077_v61;
        let _init26 = ((_1078_p0) => function (_1079_i5) {
          return _module.__default.safeDivisionInt(_1079_i5, _1078_p0);
        })(p0);
        let _nw152 = Array((new BigNumber(21)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw152.length); _i0_26++) {
          _nw152[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1077_v61 = _nw152;
        let _1080_v62;
        _1080_v62 = _dafny.Seq.of(_1035_v27, _1035_v27, _1035_v27);
        let _index160 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1077_v61).length));
        (_1077_v61)[_index160] = new BigNumber((_1080_v62).length);
        let _index161 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1077_v61).length));
        (_1077_v61)[_index161] = (new BigNumber(382)).multipliedBy(new BigNumber(((_module.__default.fm12(p0, _this.f3, p0, globalState)).update(_1075_v59, (_this).f4)).length));
      } else {
        r0 = (p1).fm0(globalState);
        let _1081_v64;
        let _init27 = ((_1082_p0, _1083_v1) => function (_1084_i6) {
          return (_dafny.Set.fromElements(_1082_p0, _1082_p0)).IsDisjointFrom(function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of _dafny.IntegerRange(new BigNumber(550), new BigNumber(-61))) {
              let _1085_v63 = _compr_44;
              if (((new BigNumber(550)).isLessThanOrEqualTo(_1085_v63)) && ((_1085_v63).isLessThan(new BigNumber(-61)))) {
                _coll44.add((_1085_v63).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), ((_1086_v1) => function (_1087_i7) {
                  return _1086_v1;
                })(_1083_v1))).length)));
              }
            }
            return _coll44;
          }());
        })(p0, _1004_v1);
        let _nw153 = Array((new BigNumber(20)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw153.length); _i0_27++) {
          _nw153[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1081_v64 = _nw153;
        let _1088_v65;
        _1088_v65 = _dafny.MultiSet.fromElements(_1003_v0);
        let _1089_v66;
        _1089_v66 = _dafny.Set.fromElements(new BigNumber(508), p0);
        let _index162 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1081_v64).length));
        (_1081_v64)[_index162] = (_1089_v66).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_1088_v65).cardinality())));
        let _1090_v67;
        _1090_v67 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1).f4);
        let _index163 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1081_v64).length));
        (_1081_v64)[_index163] = !(((((_1090_v67).contains(new BigNumber(302))) ? ((_1090_v67).get(new BigNumber(302))) : ((p1).f4))) && (false));
        let _1091_v68;
        _1091_v68 = _dafny.Seq.UnicodeFromString("kqqftlo");
        let _1092_v69;
        _1092_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1091_v68,p0);
        _1092_v69 = _module.__default.fm13((p0).multipliedBy(p0), (_dafny.ZERO).minus(new BigNumber((_1006_v3).length)), p0, (_this).f4, globalState);
        (p1).f3 = p1.f3;
        let _index164 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1081_v64).length));
        (_1081_v64)[_index164] = (_1009_v6)[_module.__default.safeIndex((((_1007_v4).contains(false)) ? ((_1007_v4).get(false)) : (p0)), new BigNumber((_1009_v6).length))];
      }
      let _hi7 = p0;
      for (let _1093_i8 = ((_this.f3) ? (p0) : (new BigNumber(-876))); _1093_i8.isLessThan(_hi7); _1093_i8 = _1093_i8.plus(_dafny.ONE)) {
        let _1094_v70;
        let _nw154 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1094_v70 = _nw154;
        let _index165 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1094_v70).length));
        (_1094_v70)[_index165] = _module.__default.safeModuloInt(p0, new BigNumber((_1006_v3).length));
        let _index166 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1094_v70).length));
        (_1094_v70)[_index166] = _module.__default.safeModuloInt(_1093_i8, (_1093_i8).multipliedBy(p0));
        let _index167 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1094_v70).length));
        (_1094_v70)[_index167] = p0;
        let _1095_v71;
        let _nw155 = new _module.C0();
        _nw155.__ctor();
        _1095_v71 = _nw155;
        let _index168 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1094_v70).length));
        (_1094_v70)[_index168] = new BigNumber((_1005_v2).cardinality());
      }
      let _1096_v72;
      let _init28 = ((_1097_p0) => function (_1098_i9) {
        return (_1098_i9).plus(_1097_p0);
      })(p0);
      let _nw156 = Array((new BigNumber(14)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw156.length); _i0_28++) {
        _nw156[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1096_v72 = _nw156;
      let _1099_v73;
      _1099_v73 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_1096_v72);
      let _1100_v74;
      let _nw157 = Array((new BigNumber(17)).toNumber());
      _nw157[(_dafny.ZERO).toNumber()] = _1096_v72;
      _nw157[(_dafny.ONE).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(2)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(3)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(4)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(5)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(6)).toNumber()] = (((p1).f4) ? (_1096_v72) : (_1096_v72));
      _nw157[(new BigNumber(7)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(8)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(9)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(10)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(11)).toNumber()] = ((_this.f3) ? (_1096_v72) : (_1096_v72));
      _nw157[(new BigNumber(12)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(13)).toNumber()] = (((_1099_v73).contains((_this).f4)) ? ((_1099_v73).get((_this).f4)) : (_1096_v72));
      _nw157[(new BigNumber(14)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(15)).toNumber()] = _1096_v72;
      _nw157[(new BigNumber(16)).toNumber()] = _1096_v72;
      _1100_v74 = _nw157;
      let _1101_v75;
      _1101_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1004_v1,_1096_v72);
      let _index169 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1100_v74).length));
      (_1100_v74)[_index169] = (((_1101_v75).contains(_1004_v1)) ? ((_1101_v75).get(_1004_v1)) : (_1096_v72));
      let _1102_v76;
      _1102_v76 = _dafny.Seq.of(_1096_v72, _1096_v72, _1096_v72, _1096_v72, _1096_v72);
      let _1103_v77;
      _1103_v77 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f4);
      let _index170 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_1100_v74).length));
      (_1100_v74)[_index170] = (_1102_v76)[_module.__default.safeIndex((p0).minus(new BigNumber((_1103_v77).length)), new BigNumber((_1102_v76).length))];
      let _1104_v78;
      _1104_v78 = _dafny.Seq.of(p0);
      let _hi8 = p0;
      for (let _1105_i10 = (new BigNumber((_1104_v78).length)).multipliedBy(p0); _1105_i10.isLessThan(_hi8); _1105_i10 = _1105_i10.plus(_dafny.ONE)) {
        let _1106_v79;
        _1106_v79 = _dafny.Seq.UnicodeFromString("lnyygcmow");
        let _1107_v80;
        _1107_v80 = _dafny.Seq.of(_1106_v79);
        _1107_v80 = _1107_v80;
        let _1108_v81;
        let _nw158 = Array((new BigNumber(23)).toNumber()).fill(false);
        _1108_v81 = _nw158;
        let _1109_v82;
        _1109_v82 = _dafny.MultiSet.fromElements(_1108_v81, _1108_v81);
        _1109_v82 = _1109_v82;
        let _init29 = ((_1110_p1, _1111_v1) => function (_1112_i11) {
          return !((_module.D0.create_DC1((_1110_p1).f4, _1111_v1)).dtor_cf1);
        })(p1, _1004_v1);
        let _nw159 = Array((new BigNumber(24)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw159.length); _i0_29++) {
          _nw159[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1108_v81 = _nw159;
        (p1).f3 = _this.f3;
      }
      let _1113_v83;
      let _nw160 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1113_v83 = _nw160;
      let _1114_v84;
      _1114_v84 = _dafny.Seq.of(_1113_v83, _1113_v83);
      let _1115_v85;
      _1115_v85 = _dafny.Seq.of(_1114_v84);
      let _1116_v86;
      _1116_v86 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
      _1114_v84 = (((p0).isLessThanOrEqualTo(p0)) ? ((_1115_v85)[_module.__default.safeIndex(new BigNumber((_1116_v86).length), new BigNumber((_1115_v85).length))]) : (_1114_v84));
      r0 = p0;
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _1117_v0;
      _1117_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_this.f3);
      let _1118_v1;
      _1118_v1 = _module.D0.create_DC1((((_1117_v0).contains(_module.__default.fm14(globalState))) ? ((_1117_v0).get(_module.__default.fm14(globalState))) : (!((_this).f4))), p0);
      _1118_v1 = _1118_v1;
      if (_this.f3) {
        let _1119_v2;
        _1119_v2 = _dafny.Seq.UnicodeFromString("ua");
        let _1120_v3;
        let _nw161 = Array((new BigNumber(4)).toNumber());
        _nw161[(_dafny.ZERO).toNumber()] = _this.f3;
        _nw161[(_dafny.ONE).toNumber()] = p1.f3;
        _nw161[(new BigNumber(2)).toNumber()] = (_this.f3) === ((_this).f4);
        _nw161[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1119_v2, _1119_v2);
        _1120_v3 = _nw161;
        _1120_v3 = _1120_v3;
        let _1121_v4;
        _1121_v4 = new BigNumber(902);
        let _1122_v5;
        _1122_v5 = _dafny.Seq.of(_this.f3);
        let _1123_v6;
        let _nw162 = Array((new BigNumber(10)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = new BigNumber((_1122_v5).length);
        _nw162[(_dafny.ONE).toNumber()] = _1121_v4;
        _nw162[(new BigNumber(2)).toNumber()] = _1121_v4;
        _nw162[(new BigNumber(3)).toNumber()] = new BigNumber(241);
        _nw162[(new BigNumber(4)).toNumber()] = _1121_v4;
        _nw162[(new BigNumber(5)).toNumber()] = _1121_v4;
        _nw162[(new BigNumber(6)).toNumber()] = _1121_v4;
        _nw162[(new BigNumber(7)).toNumber()] = new BigNumber(733);
        _nw162[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-98));
        _nw162[(new BigNumber(9)).toNumber()] = _1121_v4;
        _1123_v6 = _nw162;
        let _1124_v7;
        _1124_v7 = _dafny.Seq.of(_1123_v6);
        let _1125_v8;
        _1125_v8 = _dafny.Set.fromElements(_1121_v4, new BigNumber((_1124_v7).length));
        (_this).f3 = (new BigNumber((_1125_v8).length)).isLessThan(_1121_v4);
        let _1126_v9;
        _1126_v9 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(34), (_dafny.ZERO).minus(_1121_v4)));
        _1126_v9 = _1126_v9;
        let _1127_v10;
        let _nw163 = new _module.C0();
        _nw163.__ctor();
        _1127_v10 = _nw163;
        let _1128_v11;
        let _nw164 = new _module.C0();
        _nw164.__ctor();
        _1128_v11 = _nw164;
      } else {
        let _1129_v12;
        _1129_v12 = new BigNumber(123);
        _1129_v12 = _1129_v12;
        let _1130_v13;
        _1130_v13 = _dafny.MultiSet.fromElements((p1).f4);
        let _1131_v14;
        _1131_v14 = _dafny.Seq.of(_1129_v12, new BigNumber(((_1130_v13).update((p1).f4, _module.__default.abs(_1129_v12))).cardinality()), _1129_v12);
        _1129_v12 = (_1131_v14)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_1131_v14).length), (_1131_v14)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), ((_1132_p0) => function (_1133_i0) {
          return _1132_p0;
        })(p0))).length), new BigNumber((_1131_v14).length))]), new BigNumber((_1131_v14).length))];
        let _1134_v15;
        let _nw165 = Array((new BigNumber(12)).toNumber()).fill(_module.D0.Default());
        _1134_v15 = _nw165;
        _1134_v15 = _1134_v15;
        let _1135_v16;
        _1135_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14(globalState),_1129_v12);
        _1135_v16 = (_1135_v16).update((_this).f4, _1129_v12);
        let _1136_v17;
        let _nw166 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _1136_v17 = _nw166;
        let _1137_v18;
        _1137_v18 = _dafny.Seq.of(_dafny.Seq.of(_1131_v14));
        let _index171 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1136_v17).length));
        (_1136_v17)[_index171] = (_1137_v18)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1137_v18).length))];
        let _1138_v19;
        let _nw167 = Array((new BigNumber(10)).toNumber()).fill(false);
        _1138_v19 = _nw167;
        let _1139_v20;
        _1139_v20 = _module.D1.create_DC3(_1138_v19);
        let _1140_v21;
        _1140_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1139_v20).dtor_cf6,(_1129_v12).minus(_1129_v12));
        let _index172 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1136_v17).length));
        let _rhs153 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(453)), ((_1141_v14) => function (_1142_i1) {
          return _dafny.Seq.Concat(_1141_v14, _1141_v14);
        })(_1131_v14));
        let _rhs154 = new BigNumber((_1140_v21).length);
        let _lhs103 = _1136_v17;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1136_v17).length));
        _lhs103[_lhs104] = _rhs153;
        _1129_v12 = _rhs154;
      }
      let _1143_i2;
      _1143_i2 = _dafny.ZERO;
      L4: {
        while (_this.f3) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1143_i2)) {
              break L4;
            }
            _1143_i2 = (_1143_i2).plus(_dafny.ONE);
            let _1144_v22;
            let _nw168 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
            _1144_v22 = _nw168;
            let _1145_v23;
            _1145_v23 = new BigNumber(63);
            let _index173 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1144_v22).length));
            (_1144_v22)[_index173] = _1145_v23;
            let _index174 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1144_v22).length));
            (_1144_v22)[_index174] = ((_1145_v23).multipliedBy(new BigNumber(764))).multipliedBy(_1145_v23);
            (p1).f3 = (p1).f4;
            _1145_v23 = (p1).fm0(globalState);
            let _index175 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1144_v22).length));
            (_1144_v22)[_index175] = (p1).fm0(globalState);
          }
        }
      }
      if (!((_this).f4)) {
        let _1146_v24;
        _1146_v24 = new BigNumber(225);
        _1146_v24 = _1146_v24;
        let _1147_v25;
        let _nw169 = new _module.C0();
        _nw169.__ctor();
        _1147_v25 = _nw169;
        let _1148_v26;
        _1148_v26 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,p0);
        let _1149_v27;
        _1149_v27 = _module.D0.create_DC2(true, new BigNumber((_1148_v26).length), (_dafny.ZERO).minus(_1146_v24));
        let _1150_v28;
        _1150_v28 = _module.D0.create_DC2((_1149_v27).dtor_cf3, (_dafny.ZERO).minus(_1146_v24), _1146_v24);
        let _rhs155 = _1147_v25;
        let _rhs156 = _1150_v28;
        _1147_v25 = _rhs155;
        _1149_v27 = _rhs156;
        _1147_v25 = (((((_1117_v0).contains(_module.__default.fm14(globalState))) ? ((_1117_v0).get(_module.__default.fm14(globalState))) : ((p1).f4))) ? (_1147_v25) : (_1147_v25));
        let _1151_v29;
        _1151_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(982),((p1).f4) === ((_this).f4));
        let _1152_v30;
        let _init30 = function (_1153_i3) {
          return (_this).f4;
        };
        let _nw170 = Array((new BigNumber(22)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw170.length); _i0_30++) {
          _nw170[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1152_v30 = _nw170;
        let _1154_v31;
        _1154_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1152_v30,_1117_v0);
        let _1155_v32;
        _1155_v32 = _dafny.Seq.of(!((p1).f4));
        let _rhs157 = (((_1151_v29).contains(_1146_v24)) ? ((_1151_v29).get(_1146_v24)) : ((new BigNumber((_1154_v31).length)).isEqualTo(new BigNumber((_1155_v32).length))));
        let _rhs158 = false;
        let _rhs159 = _this.f3;
        let _rhs160 = false;
        let _rhs161 = p1.f3;
        let _lhs105 = _this;
        let _lhs106 = p1;
        let _lhs107 = globalState;
        let _lhs108 = p1;
        let _lhs109 = globalState;
        _lhs105.f3 = _rhs157;
        _lhs106.f3 = _rhs158;
        _lhs107.f0 = _rhs159;
        _lhs108.f3 = _rhs160;
        _lhs109.f0 = _rhs161;
        let _1156_v33;
        _1156_v33 = _dafny.MultiSet.fromElements(_1146_v24);
        let _1157_v34;
        _1157_v34 = _dafny.Seq.of((_1147_v25).fm10(_1146_v24, (p1).f4, _1156_v33, globalState), _1146_v24, _1146_v24, _1146_v24, _1146_v24);
        _1146_v24 = (_1157_v34)[_module.__default.safeIndex((_dafny.ZERO).minus((_1146_v24).minus((_1157_v34)[_module.__default.safeIndex(_1146_v24, new BigNumber((_1157_v34).length))])), new BigNumber((_1157_v34).length))];
      } else {
        let _1158_v35;
        _1158_v35 = new BigNumber(877);
        let _1159_v36;
        _1159_v36 = _dafny.Seq.of(_1158_v35);
        let _1160_v37;
        _1160_v37 = _dafny.Seq.of(p1.f3);
        let _1161_v38;
        _1161_v38 = _dafny.Seq.UnicodeFromString("ywhnqq");
        _1159_v36 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1159_v36, _dafny.Seq.of(_1158_v35, _1158_v35, _1158_v35, new BigNumber((_1160_v37).length), new BigNumber((_1161_v38).length))), _1159_v36);
        let _1162_v39;
        _1162_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1158_v35,_1158_v35);
        _1162_v39 = (_1162_v39).update(_1158_v35, (_dafny.ZERO).minus(_1158_v35));
        let _1163_v40;
        _1163_v40 = _dafny.Set.fromElements(p1.f3);
        let _1164_v41;
        _1164_v41 = _dafny.Set.fromElements(_1163_v40);
        _1164_v41 = _1164_v41;
        let _1165_v42;
        _1165_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1158_v35,p1.f3);
        _1165_v42 = ((_1165_v42).Merge(_module.__default.fm15((p1).f4, _1158_v35, _1158_v35, (p1).f4, globalState))).Merge((_1165_v42).Merge(_1165_v42));
        let _1166_v43;
        _1166_v43 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.D0.create_DC2(p1.f3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-986)), ((_1167_p0) => function (_1168_i4) {
  return _1167_p0;
})(p0))).length), _1158_v35));
        let _1169_v44;
        _1169_v44 = _dafny.Seq.of(_1166_v43);
        (_this).f3 = !_dafny.areEqual(_1169_v44, _1169_v44);
      }
      let _1170_v45;
      let _init31 = ((_1171_p1) => function (_1172_i5) {
        return ((_1171_p1.f3) ? ((_1171_p1).f4) : (_1171_p1.f3));
      })(p1);
      let _nw171 = Array((new BigNumber(28)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw171.length); _i0_31++) {
        _nw171[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1170_v45 = _nw171;
      let _1173_v46;
      _1173_v46 = new BigNumber(270);
      let _1174_v47;
      _1174_v47 = _dafny.Seq.UnicodeFromString("vg");
      let _rhs162 = !((_dafny.ZERO).minus(_1173_v46)).isEqualTo(((_this.f3) ? (_1173_v46) : (_1173_v46)));
      let _rhs163 = ((p1.f3) ? (_1170_v45) : (_1170_v45));
      let _rhs164 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1174_v47, _1174_v47), _1174_v47)).length);
      let _lhs110 = _this;
      _lhs110.f3 = _rhs162;
      _1170_v45 = _rhs163;
      _1173_v46 = _rhs164;
      let _hi9 = _1173_v46;
      for (let _1175_i6 = (_1173_v46).multipliedBy(_1173_v46); _1175_i6.isLessThan(_hi9); _1175_i6 = _1175_i6.plus(_dafny.ONE)) {
        let _1176_v48;
        _1176_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
        let _1177_v49;
        _1177_v49 = _dafny.Seq.of(_1176_v48);
        let _1178_v50;
        let _nw172 = Array((new BigNumber(12)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = (_1173_v46).multipliedBy(_1173_v46);
        _nw172[(_dafny.ONE).toNumber()] = new BigNumber((_1174_v47).length);
        _nw172[(new BigNumber(2)).toNumber()] = new BigNumber((_1174_v47).length);
        _nw172[(new BigNumber(3)).toNumber()] = _1175_i6;
        _nw172[(new BigNumber(4)).toNumber()] = new BigNumber((_1174_v47).length);
        _nw172[(new BigNumber(5)).toNumber()] = (_1173_v46).plus(_1175_i6);
        _nw172[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1177_v49).length), _1173_v46))));
        _nw172[(new BigNumber(7)).toNumber()] = (p1).fm0(globalState);
        _nw172[(new BigNumber(8)).toNumber()] = new BigNumber((_1174_v47).length);
        _nw172[(new BigNumber(9)).toNumber()] = (_1175_i6).plus(new BigNumber(767));
        _nw172[(new BigNumber(10)).toNumber()] = _1173_v46;
        _nw172[(new BigNumber(11)).toNumber()] = _1175_i6;
        _1178_v50 = _nw172;
        let _index176 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length));
        (_1178_v50)[_index176] = _1173_v46;
        let _1179_v51;
        let _nw173 = new _module.C0();
        _nw173.__ctor();
        _1179_v51 = _nw173;
        let _1180_v52;
        _1180_v52 = _dafny.Set.fromElements(_1179_v51, _1179_v51);
        let _index177 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length));
        (_1178_v50)[_index177] = (_1175_i6).multipliedBy((_1175_i6).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1180_v52).length))));
        let _1181_v53;
        _1181_v53 = _dafny.Set.fromElements(p1.f3, p1.f3, _this.f3, (p1).f4, _module.__default.fm14(globalState));
        let _1182_v54;
        _1182_v54 = _module.D0.create_DC0(new BigNumber(((_1179_v51).fm11(p0, _1181_v53, globalState)).length));
        let _source14 = _1182_v54;
        if (_source14.is_DC1) {
          let _1183___mcc_h0 = (_source14).cf1;
          let _1184___mcc_h1 = (_source14).cf2;
          let _1185_cf2 = _1184___mcc_h1;
          let _1186_cf1 = _1183___mcc_h0;
          let _1187_v55;
          let _nw174 = new _module.C0();
          _nw174.__ctor();
          _1187_v55 = _nw174;
          let _1188_v56;
          _1188_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))], (_dafny.ZERO).minus(_1173_v46), _1173_v46, _1173_v46, (_dafny.ZERO).minus((_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))]))).length),(_this).f4);
          _1188_v56 = (function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of _dafny.IntegerRange(new BigNumber(652), new BigNumber(-955))) {
              let _1189_v57 = _compr_45;
              if (((new BigNumber(652)).isLessThanOrEqualTo(_1189_v57)) && ((_1189_v57).isLessThan(new BigNumber(-955)))) {
                _coll45.push([_module.__default.safeModuloInt(_1189_v57, _1175_i6),true]);
              }
            }
            return _coll45;
          }()).Merge(_1188_v56);
          let _1190_v58;
          _1190_v58 = _module.D1.create_DC3(_1170_v45);
          _1190_v58 = _1190_v58;
          let _1191_v59;
          let _nw175 = new _module.C0();
          _nw175.__ctor();
          _1191_v59 = _nw175;
        } else if (_source14.is_DC2) {
          let _1192___mcc_h2 = (_source14).cf3;
          let _1193___mcc_h3 = (_source14).cf4;
          let _1194___mcc_h4 = (_source14).cf5;
          let _1195_cf5 = _1194___mcc_h4;
          let _1196_cf4 = _1193___mcc_h3;
          let _1197_cf3 = _1192___mcc_h2;
          let _1198_v60;
          _1198_v60 = _dafny.Seq.of((p1).f4, _this.f3);
          let _1199_v61;
          _1199_v61 = _dafny.Set.fromElements((_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))], (_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))]);
          (p1).f3 = (_1199_v61).contains(new BigNumber((_1198_v60).length));
          _1179_v51 = _1179_v51;
          _1117_v0 = (_1117_v0).update(!((p1).f4), ((_dafny.ZERO).minus(_1173_v46)).isEqualTo((_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))]));
          let _index178 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length));
          (_1178_v50)[_index178] = _module.__default.safeDivisionInt(_1196_cf4, _1175_i6);
        } else {
          let _1200___mcc_h5 = (_source14).cf0;
          let _1201_cf0 = _1200___mcc_h5;
          let _1202_v62;
          let _nw176 = new _module.C1();
          _nw176.__ctor(_1174_v47, p1.f3);
          _1202_v62 = _nw176;
          let _1203_v63;
          _1203_v63 = _1202_v62;
          let _1204_v64;
          _1204_v64 = _dafny.Seq.of(_1202_v62, (_1203_v63));
          let _1205_v65;
          _1205_v65 = _dafny.Map.Empty.slice().updateUnsafe((!((p1).f4)) === (true),_1204_v64);
          let _1206_v66;
          _1206_v66 = _dafny.Map.Empty.slice().updateUnsafe((_1178_v50)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length))],_1205_v65);
          _1205_v65 = ((((_1206_v66).contains(_1201_cf0)) ? ((_1206_v66).get(_1201_cf0)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1204_v64)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_1204_v64)).Merge(_1205_v65));
          (globalState).f0 = p1.f3;
          (globalState).f0 = (p1).f4;
          let _index179 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1178_v50).length));
          (_1178_v50)[_index179] = (((p1).f4) ? ((_1201_cf0).minus((_this).fm0(globalState))) : (_1175_i6));
        }
        _1117_v0 = (_1117_v0).update((p1).f4, _this.f3);
        let _1207_v67;
        let _nw177 = new _module.C1();
        _nw177.__ctor(_1174_v47, ((p1).f4) || (p1.f3));
        _1207_v67 = _nw177;
        _1207_v67 = _1207_v67;
      }
      r0 = _dafny.Seq.Concat(((_this.f3) ? (_1174_v47) : (_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_1208_p0) => function (_1209_i7) {
        return _1208_p0;
      })(p0)), _module.__default.safeIndex(_1173_v46, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_1210_p0) => function (_1211_i7) {
        return _1210_p0;
      })(p0))).length)), p0), _module.__default.safeIndex(_1173_v46, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_1212_p0) => function (_1213_i7) {
        return _1212_p0;
      })(p0)), _module.__default.safeIndex(_1173_v46, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), ((_1214_p0) => function (_1215_i7) {
        return _1214_p0;
      })(p0))).length)), p0)).length)), p0))), _1174_v47);
      return r0;
    }
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _1216_v0;
      let _init32 = ((_1217_p2) => function (_1218_i0) {
        return (_1218_i0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1217_p2,_1217_p2)).length));
      })(p2);
      let _nw178 = Array((new BigNumber(5)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw178.length); _i0_32++) {
        _nw178[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1216_v0 = _nw178;
      let _index180 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length));
      (_1216_v0)[_index180] = p2;
      let _index181 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length));
      (_1216_v0)[_index181] = p2;
      let _1219_v1;
      let _init33 = function (_1220_i1) {
        return !((_this.f3) || (true));
      };
      let _nw179 = Array((new BigNumber(28)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw179.length); _i0_33++) {
        _nw179[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1219_v1 = _nw179;
      let _index182 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length));
      (_1219_v1)[_index182] = false;
      let _1221_v2;
      _1221_v2 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1222_v3;
      _1222_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1221_v2);
      let _1223_v4;
      _1223_v4 = _dafny.Seq.of(_1221_v2, (((_1222_v3).contains(p1)) ? ((_1222_v3).get(p1)) : (new _dafny.CodePoint('x'.codePointAt(0)))), _module.__default.fm23((_this).f4, (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], globalState));
      let _1224_v5;
      _1224_v5 = _dafny.Seq.of(new BigNumber(386));
      let _1225_v6;
      _1225_v6 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_1226_p2) => function (_1227_i2) {
        return (_dafny.ZERO).minus(_1226_p2);
      })(p2)), _1224_v5);
      let _1228_v7;
      _1228_v7 = _dafny.MultiSet.fromElements(new BigNumber((_1223_v4).length), (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], new BigNumber((_1225_v6).length), (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], new BigNumber(756));
      let _index183 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length));
      let _index184 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length));
      let _rhs165 = (_this.f3) && (_this.f3);
      let _rhs166 = (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))];
      let _rhs167 = ((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]).plus((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]);
      let _rhs168 = ((p2).plus((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))])).isLessThanOrEqualTo(new BigNumber((_1228_v7).cardinality()));
      let _rhs169 = (_dafny.ZERO).minus((p2).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p2, (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]))));
      let _lhs111 = _this;
      let _lhs112 = _1216_v0;
      let _lhs113 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length));
      let _lhs114 = _1219_v1;
      let _lhs115 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length));
      _lhs111.f3 = _rhs165;
      _lhs112[_lhs113] = _rhs166;
      r1 = _rhs167;
      _lhs114[_lhs115] = _rhs168;
      r1 = _rhs169;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1219_v1).length))) {
        let _1229_i3 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1229_i3)) && ((_1229_i3).isLessThan(new BigNumber((_1219_v1).length))))) {
          (_1219_v1)[(_1229_i3)] = !((_this).f4);
        }
      }
      let _1230_i4;
      _1230_i4 = _dafny.ZERO;
      L5: {
        while (((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]).isLessThanOrEqualTo((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))])) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1230_i4)) {
              break L5;
            }
            _1230_i4 = (_1230_i4).plus(_dafny.ONE);
            let _1231_v8;
            _1231_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
            _1231_v8 = _1231_v8;
            let _1232_v9;
            _1232_v9 = _dafny.Seq.of((_this).f4, true);
            let _1233_v10;
            _1233_v10 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(450)), ((_1234_v4) => function (_1235_i5) {
              return _1234_v4;
            })(_1223_v4));
            let _1236_v11;
            _1236_v11 = _module.D1.create_DC5(false, _this.f3, _module.__default.fm31(false, (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], (_this).f4, _1232_v9, globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_1233_v10)).length));
            let _index185 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length));
            (_1219_v1)[_index185] = (_1236_v11).dtor_cf8;
            let _1237_v12;
            let _init34 = ((_1238_v1, _1239_p0, _1240_v0) => function (_1241_i6) {
              return _dafny.Seq.update(_dafny.Seq.of((_1238_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1238_v1).length))], _1239_p0, _this.f3), _module.__default.safeIndex((_1240_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1240_v0).length))], new BigNumber((_dafny.Seq.of((_1238_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1238_v1).length))], _1239_p0, _this.f3)).length)), (_1238_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1238_v1).length))]);
            })(_1219_v1, p0, _1216_v0);
            let _nw180 = Array((new BigNumber(6)).toNumber());
            for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw180.length); _i0_34++) {
              _nw180[_i0_34] = _init34(new BigNumber(_i0_34));
            }
            _1237_v12 = _nw180;
            let _rhs170 = _dafny.Seq.contains(_1224_v5, _module.__default.safeDivisionInt(new BigNumber(684), new BigNumber((_dafny.Seq.of((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], p2)).length)));
            let _rhs171 = _1237_v12;
            let _rhs172 = (p2).isLessThan(new BigNumber(158));
            let _lhs116 = globalState;
            let _lhs117 = globalState;
            _lhs116.f0 = _rhs170;
            _1237_v12 = _rhs171;
            _lhs117.f0 = _rhs172;
            let _1242_v13;
            let _nw181 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
            _1242_v13 = _nw181;
            let _1243_v14;
            _1243_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]));
            let _index186 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_1242_v13).length));
            (_1242_v13)[_index186] = (_1243_v14).Merge(_1243_v14);
            let _index187 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_1242_v13).length));
            (_1242_v13)[_index187] = _1243_v14;
          }
        }
      }
      let _index188 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length));
      (_1219_v1)[_index188] = (_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))];
      let _1244_v15;
      _1244_v15 = _dafny.Set.fromElements(p2);
      let _1245_v17;
      _1245_v17 = _module.D7.create_DC16(_module.__default.fm36(new BigNumber((function () {
  let _coll46 = new _dafny.Set();
  for (const _compr_46 of (_1244_v15).Elements) {
    let _1246_v16 = _compr_46;
    if ((_1244_v15).contains(_1246_v16)) {
      _coll46.add((_1246_v16).plus(new BigNumber((_dafny.Seq.of(true, false)).length)));
    }
  }
  return _coll46;
}()).length), (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], globalState));
      let _1247_v18;
      _1247_v18 = _module.D1.create_DC4((_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))]);
      let _1248_v19;
      _1248_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1245_v17,_1247_v18);
      let _1249_v20;
      _1249_v20 = _dafny.Set.fromElements((_module.D1.create_DC5((_this).f4, (_this).f4, (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))], new BigNumber(-772))).dtor_cf9);
      let _1250_v21;
      _1250_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_1249_v20);
      let _pat_let_tv36 = _1250_v21;
      let _pat_let_tv37 = _1250_v21;
      let _pat_let_tv38 = p2;
      let _pat_let_tv39 = _1249_v20;
      _1248_v19 = (_1248_v19).update(function (_pat_let21_0) {
        return function (_1253_dt__update__tmp_h1) {
          return function (_pat_let24_0) {
            return function (_1254_dt__update_hcf25_h1) {
              return _module.D7.create_DC16(_1254_dt__update_hcf25_h1);
            }(_pat_let24_0);
          }((_pat_let_tv37).update(_pat_let_tv38, _pat_let_tv39));
        }(_pat_let21_0);
      }(function (_pat_let22_0) {
        return function (_1251_dt__update__tmp_h0) {
          return function (_pat_let23_0) {
            return function (_1252_dt__update_hcf25_h0) {
              return _module.D7.create_DC16(_1252_dt__update_hcf25_h0);
            }(_pat_let23_0);
          }(_pat_let_tv36);
        }(_pat_let22_0);
      }(_1245_v17)), _1247_v18);
      let _1255_v22;
      _1255_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(p0) || ((_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))]),_1216_v0);
      r0 = (((_1255_v22).contains((((_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))]) ? ((_this).f4) : ((_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))])))) ? ((_1255_v22).get((((_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))]) ? ((_this).f4) : ((_1219_v1)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_1219_v1).length))])))) : (_1216_v0));
      r1 = (_1216_v0)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1216_v0).length))];
      return [r0, r1];
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      if ((_this).f4) {
        r0 = false;
        let _1256_v0;
        _1256_v0 = _module.D1.create_DC5((_this).f4, (_this).f4, p0, new BigNumber(-259));
        (_this).f3 = (p0).isLessThan((new BigNumber((_module.__default.fm30(function (_pat_let25_0) {
          return function (_1257_dt__update__tmp_h0) {
            return function (_pat_let26_0) {
              return function (_1258_dt__update_hcf8_h0) {
                return function (_pat_let27_0) {
                  return function (_1259_dt__update_hcf9_h0) {
                    return _module.D1.create_DC5(_1258_dt__update_hcf8_h0, _1259_dt__update_hcf9_h0, (_1257_dt__update__tmp_h0).dtor_cf10, (_1257_dt__update__tmp_h0).dtor_cf11);
                  }(_pat_let27_0);
                }(false);
              }(_pat_let26_0);
            }(_this.f3);
          }(_pat_let25_0);
        }(_1256_v0), p0, p0, globalState)).length)).plus(p0));
        let _1260_v1;
        _1260_v1 = _dafny.Seq.of((_this).f4);
        let _1261_v2;
        let _nw182 = new _module.C1();
        _nw182.__ctor(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pavx"), _dafny.Seq.UnicodeFromString("wibc")), (((_1260_v1)[_module.__default.safeIndex(p0, new BigNumber((_1260_v1).length))]) ? (!((_this).f4)) : ((_this).f4)));
        _1261_v2 = _nw182;
        let _1262_v3;
        _1262_v3 = _dafny.MultiSet.fromElements(p0);
        let _1263_v4;
        _1263_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1262_v3);
        let _1264_v5;
        _1264_v5 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1263_v4).length)).plus(p0),p0);
        _1264_v5 = (_1264_v5).update((_this).fm0(globalState), (p0).minus(p0));
        let _1265_v6;
        let _nw183 = new _module.C0();
        _nw183.__ctor();
        _1265_v6 = _nw183;
      } else {
        let _1266_v7;
        let _nw184 = new _module.C0();
        _nw184.__ctor();
        _1266_v7 = _nw184;
        let _1267_v8;
        let _1268_v9;
        let _out22;
        let _out23;
        let _outcollector7 = (_this).m8(new BigNumber(716), p0, globalState);
        _out22 = _outcollector7[0];
        _out23 = _outcollector7[1];
        _1267_v8 = _out22;
        _1268_v9 = _out23;
        (_this).f3 = (_this).f4;
        let _1269_v10;
        _1269_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_dafny.MultiSet.fromElements(true));
        _1269_v10 = (_1269_v10).update((_this.f3) && (_this.f3), _module.__default.fm20(_this.f3, globalState));
        let _1270_v11;
        _1270_v11 = _dafny.Seq.UnicodeFromString("rgbmpb");
        _1270_v11 = _1270_v11;
      }
      let _1271_v12;
      _1271_v12 = new BigNumber(-138);
      let _1272_v13;
      _1272_v13 = _dafny.Seq.UnicodeFromString("pvdeve");
      let _1273_v14;
      _1273_v14 = _dafny.Seq.of(_1272_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(99)), function (_1274_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
      let _1275_v15;
      _1275_v15 = _dafny.Seq.of((_this).f4, (_this).f4);
      let _1276_v16;
      _1276_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f4);
      let _1277_v17;
      _1277_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1271_v12);
      let _1278_v19;
      let _nw185 = Array((new BigNumber(22)).toNumber());
      _nw185[(_dafny.ZERO).toNumber()] = _1271_v12;
      _nw185[(_dafny.ONE).toNumber()] = p0;
      _nw185[(new BigNumber(2)).toNumber()] = new BigNumber((_1275_v15).length);
      _nw185[(new BigNumber(3)).toNumber()] = new BigNumber((_1275_v15).length);
      _nw185[(new BigNumber(4)).toNumber()] = new BigNumber(-836);
      _nw185[(new BigNumber(5)).toNumber()] = new BigNumber((_1276_v16).length);
      _nw185[(new BigNumber(6)).toNumber()] = new BigNumber(-414);
      _nw185[(new BigNumber(7)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(8)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(9)).toNumber()] = (((_1277_v17).contains((_this).f4)) ? ((_1277_v17).get((_this).f4)) : (p0));
      _nw185[(new BigNumber(10)).toNumber()] = _module.__default.fm31(true, _1271_v12, _this.f3, _1275_v15, globalState);
      _nw185[(new BigNumber(11)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(12)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(13)).toNumber()] = new BigNumber((_1272_v13).length);
      _nw185[(new BigNumber(14)).toNumber()] = p0;
      _nw185[(new BigNumber(15)).toNumber()] = new BigNumber((function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(-59), new BigNumber(797))) {
          let _1279_v18 = _compr_47;
          if (((new BigNumber(-59)).isLessThanOrEqualTo(_1279_v18)) && ((_1279_v18).isLessThan(new BigNumber(797)))) {
            _coll47.push([(_1279_v18).minus((_dafny.ZERO).minus(p0)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1271_v12,new BigNumber(902))).length))).length)]);
          }
        }
        return _coll47;
      }()).length);
      _nw185[(new BigNumber(16)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(17)).toNumber()] = p0;
      _nw185[(new BigNumber(18)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(19)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(20)).toNumber()] = _1271_v12;
      _nw185[(new BigNumber(21)).toNumber()] = new BigNumber((_1272_v13).length);
      _1278_v19 = _nw185;
      let _rhs173 = function (_source15) {
        let _1280___mcc_h0 = _source15;
        let _1281_cf24 = _1280___mcc_h0;
        return _this.f3;
      }(_1273_v14);
      let _rhs174 = new BigNumber((_dafny.MultiSet.fromElements(_1278_v19)).cardinality());
      let _lhs118 = globalState;
      _lhs118.f0 = _rhs173;
      _1271_v12 = _rhs174;
      let _hi10 = _1271_v12;
      for (let _1282_i1 = _1271_v12; _1282_i1.isLessThan(_hi10); _1282_i1 = _1282_i1.plus(_dafny.ONE)) {
        let _index189 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length));
        (_1278_v19)[_index189] = _1271_v12;
        let _index190 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length));
        (_1278_v19)[_index190] = _1271_v12;
        let _1283_v20;
        _1283_v20 = _module.D1.create_DC6();
        _1283_v20 = _1283_v20;
        let _index191 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length));
        (_1278_v19)[_index191] = new BigNumber(614);
        if (false) {
          let _1284_v21;
          _1284_v21 = _dafny.Set.fromElements((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))]);
          let _index192 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index192] = _module.__default.fm31(_this.f3, _1271_v12, (_1284_v21).IsProperSubsetOf(_1284_v21), _1275_v15, globalState);
          let _1285_v22;
          let _nw186 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _1285_v22 = _nw186;
          let _1286_v23;
          _1286_v23 = _dafny.Seq.of(_1271_v12, (_dafny.ZERO).minus((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))]));
          let _index193 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1285_v22).length));
          (_1285_v22)[_index193] = _1286_v23;
          let _index194 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1285_v22).length));
          (_1285_v22)[_index194] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))]), _module.__default.safeIndex((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], new BigNumber((_dafny.Seq.of((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))])).length)), new BigNumber((_1275_v15).length)), _1286_v23);
          let _index195 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index195] = _1282_i1;
          (globalState).f0 = (_this).f4;
          let _1287_v24;
          _1287_v24 = _dafny.MultiSet.fromElements(p0, _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(_this.f3)).cardinality()), (_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))]), p0, _1282_i1);
          _1287_v24 = (_1287_v24).update(_module.__default.safeModuloInt((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], new BigNumber((_1284_v21).length)), _module.__default.abs(_1282_i1));
        } else {
          let _1288_v25;
          _1288_v25 = _dafny.MultiSet.fromElements((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], (_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))]);
          let _1289_v26;
          _1289_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-624),_1288_v25);
          let _1290_v27;
          _1290_v27 = _dafny.Seq.of(p0);
          let _1291_v28;
          _1291_v28 = _module.D14.create_DC30(_dafny.MultiSet.fromElements(_1271_v12));
          (globalState).f0 = (((((_1289_v26).contains((_1290_v27)[_module.__default.safeIndex((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], new BigNumber((_1290_v27).length))])) ? ((_1289_v26).get((_1290_v27)[_module.__default.safeIndex((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], new BigNumber((_1290_v27).length))])) : ((_1291_v28).dtor_cf43))).Intersect(_1288_v25)).IsSubsetOf((_module.__default.fm24(globalState)).Union(_1288_v25));
          let _1292_v29;
          _1292_v29 = _dafny.Set.fromElements((_this).f4, _this.f3, (_this).f4, _this.f3);
          _1292_v29 = _1292_v29;
          let _1293_v30;
          _1293_v30 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1294_v31;
          _1294_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1293_v30,_1293_v30);
          _1294_v31 = (_1294_v31).update(_1293_v30, _1293_v30);
          let _1295_v32;
          _1295_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))],_1293_v30);
          _1295_v32 = (_1295_v32).update((_1278_v19)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1278_v19).length))], _1293_v30);
          _1276_v16 = _1276_v16;
        }
      }
      let _1296_v33;
      _1296_v33 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1297_v34;
      _1297_v34 = _dafny.MultiSet.fromElements(_1296_v33, _1296_v33);
      (globalState).f0 = (_dafny.Set.fromElements(_1297_v34)).contains(_1297_v34);
      let _hi11 = _module.__default.safeModuloInt(_1271_v12, new BigNumber((_1272_v13).length));
      for (let _1298_i2 = _1271_v12; _1298_i2.isLessThan(_hi11); _1298_i2 = _1298_i2.plus(_dafny.ONE)) {
        let _1299_v35;
        let _nw187 = new _module.C0();
        _nw187.__ctor();
        _1299_v35 = _nw187;
        _1271_v12 = _1271_v12;
        let _1300_v36;
        _1300_v36 = _module.D4.create_DC11(_dafny.Seq.Concat(_dafny.Seq.of(false, false), _1275_v15));
        let _source16 = _1300_v36;
        if (_source16.is_DC12) {
          let _1301___mcc_h1 = (_source16).cf19;
          let _1302___mcc_h2 = (_source16).cf20;
          let _1303___mcc_h3 = (_source16).cf21;
          let _1304_cf21 = _1303___mcc_h3;
          let _1305_cf20 = _1302___mcc_h2;
          let _1306_cf19 = _1301___mcc_h1;
          let _1307_v37;
          _1307_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(375),_this.f3);
          let _1308_v38;
          _1308_v38 = _dafny.MultiSet.fromElements(_1272_v13);
          let _1309_v39;
          _1309_v39 = _dafny.MultiSet.fromElements((_this).f4);
          let _1310_v40;
          let _nw188 = Array((new BigNumber(26)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = (((_1307_v37).contains(_1271_v12)) ? ((_1307_v37).get(_1271_v12)) : ((_this).f4));
          _nw188[(_dafny.ONE).toNumber()] = _1304_cf21;
          _nw188[(new BigNumber(2)).toNumber()] = _module.__default.fm14(globalState);
          _nw188[(new BigNumber(3)).toNumber()] = (_this).f4;
          _nw188[(new BigNumber(4)).toNumber()] = (_this).f4;
          _nw188[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw188[(new BigNumber(6)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(7)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(8)).toNumber()] = _1304_cf21;
          _nw188[(new BigNumber(9)).toNumber()] = !(_this.f3) || (_1304_cf21);
          _nw188[(new BigNumber(10)).toNumber()] = _1304_cf21;
          _nw188[(new BigNumber(11)).toNumber()] = ((_this.f3) ? ((_this).f4) : ((_this).f4));
          _nw188[(new BigNumber(12)).toNumber()] = ((_this.f3) ? ((_this).f4) : ((_this).f4));
          _nw188[(new BigNumber(13)).toNumber()] = _module.__default.fm14(globalState);
          _nw188[(new BigNumber(14)).toNumber()] = _module.__default.fm18(new BigNumber(722), _1296_v33, new BigNumber((_1308_v38).cardinality()), globalState);
          _nw188[(new BigNumber(15)).toNumber()] = ((((_1309_v39).contains((_this).f4)) ? ((_1309_v39).get((_this).f4)) : (_1305_cf20))).isLessThan(_1305_cf20);
          _nw188[(new BigNumber(16)).toNumber()] = true;
          _nw188[(new BigNumber(17)).toNumber()] = _1304_cf21;
          _nw188[(new BigNumber(18)).toNumber()] = ((_dafny.ZERO).minus(_1271_v12)).isLessThanOrEqualTo(p0);
          _nw188[(new BigNumber(19)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(20)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(21)).toNumber()] = false;
          _nw188[(new BigNumber(22)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(23)).toNumber()] = _this.f3;
          _nw188[(new BigNumber(24)).toNumber()] = (_this).f4;
          _nw188[(new BigNumber(25)).toNumber()] = _1304_cf21;
          _1310_v40 = _nw188;
          let _index196 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1310_v40).length));
          (_1310_v40)[_index196] = _this.f3;
          let _1311_v41;
          _1311_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1305_cf20,_1306_cf19);
          let _index197 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1310_v40).length));
          (_1310_v40)[_index197] = _module.__default.fm21(_1271_v12, _dafny.MultiSet.fromElements(_1311_v41, _1311_v41), globalState);
          (_this).f3 = _this.f3;
          _1271_v12 = _1306_cf19;
          (globalState).f0 = (_this).f4;
        } else if (_source16.is_DC11) {
          let _1312___mcc_h4 = (_source16).cf18;
          let _1313_cf18 = _1312___mcc_h4;
          let _1314_v42;
          let _nw189 = Array((new BigNumber(7)).toNumber()).fill([]);
          _1314_v42 = _nw189;
          let _1315_v43;
          _1315_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f3),_1313_cf18);
          let _1316_v44;
          let _nw190 = Array((new BigNumber(18)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = _1275_v15;
          _nw190[(_dafny.ONE).toNumber()] = _1275_v15;
          _nw190[(new BigNumber(2)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(3)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_this).f4);
          _nw190[(new BigNumber(5)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_1313_cf18, _module.__default.safeIndex(_1271_v12, new BigNumber((_1313_cf18).length)), _this.f3);
          _nw190[(new BigNumber(7)).toNumber()] = _1275_v15;
          _nw190[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_this.f3, (_this).f4);
          _nw190[(new BigNumber(9)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(10)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(11)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(12)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(13)).toNumber()] = _1275_v15;
          _nw190[(new BigNumber(14)).toNumber()] = _1275_v15;
          _nw190[(new BigNumber(15)).toNumber()] = _module.__default.fm34(_1298_i2, _1271_v12, globalState);
          _nw190[(new BigNumber(16)).toNumber()] = _1313_cf18;
          _nw190[(new BigNumber(17)).toNumber()] = (((_1315_v43).contains(_this.f3)) ? ((_1315_v43).get(_this.f3)) : (_1313_cf18));
          _1316_v44 = _nw190;
          let _index198 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_1314_v42).length));
          (_1314_v42)[_index198] = _1316_v44;
          let _index199 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_1314_v42).length));
          let _rhs175 = _1296_v33;
          let _rhs176 = _1316_v44;
          let _lhs119 = _1314_v42;
          let _lhs120 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_1314_v42).length));
          _1296_v33 = _rhs175;
          _lhs119[_lhs120] = _rhs176;
          let _1317_v45;
          _1317_v45 = _dafny.Seq.of(new BigNumber((_1272_v13).length));
          let _1318_v46;
          _1318_v46 = _dafny.Seq.of(_1317_v45, _dafny.Seq.Create(_module.__default.abs(new BigNumber(39)), ((_1319_v12) => function (_1320_i3) {
            return _1319_v12;
          })(_1271_v12)));
          let _1321_v47;
          _1321_v47 = _dafny.Set.fromElements((_1318_v46)[_module.__default.safeIndex(_1271_v12, new BigNumber((_1318_v46).length))], _dafny.Seq.of(p0, _1271_v12), _dafny.Seq.of(_1298_i2));
          _1321_v47 = _1321_v47;
          let _1322_v48;
          _1322_v48 = _dafny.MultiSet.fromElements((_this).f4);
          _1322_v48 = _1322_v48;
          let _1323_v49;
          _1323_v49 = _module.D7.create_DC17(_1298_i2, (_this).f4, (_this).f4, new BigNumber((_1272_v13).length));
          let _1324_v50;
          _1324_v50 = _dafny.Set.fromElements((_1323_v49).dtor_cf28, _this.f3);
          let _1325_v51;
          _1325_v51 = _module.D7.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1317_v45).length),_1324_v50));
          let _1326_v52;
          _1326_v52 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1271_v12),_1324_v50);
          _1325_v51 = _module.D7.create_DC16(_1326_v52);
        } else {
          let _1327___mcc_h5 = (_source16).cf22;
          let _1328_cf22 = _1327___mcc_h5;
          let _index200 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index200] = _1271_v12;
          let _index201 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index201] = _module.__default.safeModuloInt(_1271_v12, _1271_v12);
          let _1329_v53;
          _1329_v53 = _dafny.Seq.of(p0);
          _1329_v53 = _1329_v53;
          (globalState).f0 = _this.f3;
          let _index202 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index202] = (_1278_v19)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_1278_v19).length))];
        }
        r1 = _this.f3;
      }
      let _1330_v54;
      _1330_v54 = _module.D1.create_DC5(_this.f3, (p0).isLessThanOrEqualTo(p0), (_1271_v12).minus(_module.__default.fm31((_this).f4, new BigNumber((_1277_v17).length), (_this).f4, _1275_v15, globalState)), (_1271_v12).minus(_1271_v12));
      let _source17 = _1330_v54;
      if (_source17.is_DC4) {
        let _1331___mcc_h6 = (_source17).cf7;
        let _1332_cf7 = _1331___mcc_h6;
        (_this).f3 = (_this).f4;
        let _1333_v55;
        let _1334_v56;
        let _out24;
        let _out25;
        let _outcollector8 = (_this).m8(new BigNumber((_1297_v34).cardinality()), (_1271_v12).minus(p0), globalState);
        _out24 = _outcollector8[0];
        _out25 = _outcollector8[1];
        _1333_v55 = _out24;
        _1334_v56 = _out25;
        r0 = _this.f3;
        let _1335_v57;
        let _nw191 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _1335_v57 = _nw191;
        let _index203 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1335_v57).length));
        (_1335_v57)[_index203] = _1276_v16;
        let _index204 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1335_v57).length));
        (_1335_v57)[_index204] = _1276_v16;
      } else if (_source17.is_DC5) {
        let _1336___mcc_h7 = (_source17).cf8;
        let _1337___mcc_h8 = (_source17).cf9;
        let _1338___mcc_h9 = (_source17).cf10;
        let _1339___mcc_h10 = (_source17).cf11;
        let _1340_cf11 = _1339___mcc_h10;
        let _1341_cf10 = _1338___mcc_h9;
        let _1342_cf9 = _1337___mcc_h8;
        let _1343_cf8 = _1336___mcc_h7;
        (globalState).f0 = _this.f3;
        _1342_cf9 = _dafny.areEqual(_1275_v15, _1275_v15);
        (globalState).f0 = !(_1343_cf8) || (!(_1342_cf9));
        let _1344_v58;
        let _nw192 = new _module.C2();
        _nw192.__ctor(!(false), (_this).f4, _this.f3, _1343_cf8);
        _1344_v58 = _nw192;
      } else if (_source17.is_DC6) {
        (globalState).f0 = (_this).f4;
        if (_this.f3) {
          (globalState).f0 = (_this).f4;
          let _1345_v59;
          let _nw193 = new _module.C0();
          _nw193.__ctor();
          _1345_v59 = _nw193;
          let _1346_v60;
          _1346_v60 = _dafny.Seq.of((((_1277_v17).contains((_this).f4)) ? ((_1277_v17).get((_this).f4)) : (new BigNumber((_dafny.Seq.of(_this.f3, _this.f3)).length))));
          let _1347_v61;
          _1347_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1346_v60)[_module.__default.safeIndex(_1271_v12, new BigNumber((_1346_v60).length))]),(_dafny.ZERO).minus(p0));
          _1347_v61 = (_1347_v61).update((_1271_v12).minus(p0), p0);
          let _1348_v62;
          _1348_v62 = _dafny.Set.fromElements(true);
          let _1349_v63;
          _1349_v63 = _module.D12.create_DC26(_1348_v62);
          let _1350_v64;
          _1350_v64 = _module.D12.create_DC28(_1349_v63);
          let _1351_v65;
          _1351_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1350_v64);
          let _1352_v66;
          _1352_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1351_v65,(_1345_v59).fm11(new _dafny.CodePoint('m'.codePointAt(0)), _1348_v62, globalState));
          let _rhs177 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1350_v64),_1272_v13);
          let _rhs178 = _1275_v15;
          let _rhs179 = (_this).f4;
          let _lhs121 = globalState;
          _1352_v66 = _rhs177;
          _1275_v15 = _rhs178;
          _lhs121.f0 = _rhs179;
          let _1353_v67;
          let _nw194 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1353_v67 = _nw194;
          let _index205 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_1353_v67).length));
          (_1353_v67)[_index205] = new _dafny.CodePoint('a'.codePointAt(0));
          let _index206 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_1353_v67).length));
          (_1353_v67)[_index206] = _1296_v33;
        } else {
          let _index207 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index207] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p0));
          let _index208 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index208] = _1271_v12;
          let _index209 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index209] = p0;
          _1271_v12 = (_1278_v19)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length))];
          let _1354_v68;
          let _nw195 = Array((new BigNumber(22)).toNumber()).fill(false);
          _1354_v68 = _nw195;
          let _1355_v69;
          _1355_v69 = _module.D0.create_DC2(_this.f3, new BigNumber(656), p0);
          let _1356_v70;
          _1356_v70 = _dafny.Seq.of(new BigNumber((_1272_v13).length));
          let _index210 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length));
          let _rhs180 = (_this.f3) === (false);
          let _rhs181 = _module.__default.safeModuloInt(((_1278_v19)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length))]).multipliedBy((_1355_v69).dtor_cf4), _1271_v12);
          let _rhs182 = (_1356_v70)[_module.__default.safeIndex((_1278_v19)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length))], new BigNumber((_1356_v70).length))];
          let _rhs183 = _dafny.Seq.IsPrefixOf(_module.__default.fm30(_1330_v54, (_1278_v19)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length))], p0, globalState), _dafny.Seq.update(_1272_v13, _module.__default.safeIndex((_1278_v19)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length))], new BigNumber((_1272_v13).length)), _1296_v33));
          let _rhs184 = _1354_v68;
          let _lhs122 = globalState;
          let _lhs123 = _1278_v19;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1278_v19).length));
          let _lhs125 = globalState;
          _lhs122.f0 = _rhs180;
          _lhs123[_lhs124] = _rhs181;
          _1271_v12 = _rhs182;
          _lhs125.f0 = _rhs183;
          _1354_v68 = _rhs184;
          let _1357_v71;
          let _nw196 = new _module.C0();
          _nw196.__ctor();
          _1357_v71 = _nw196;
        }
        let _1358_v72;
        let _1359_v73;
        let _out26;
        let _out27;
        let _outcollector9 = (_this).m8((_1271_v12).plus((_this).fm0(globalState)), _module.__default.fm31((_this).f4, new BigNumber((_1272_v13).length), _this.f3, _1275_v15, globalState), globalState);
        _out26 = _outcollector9[0];
        _out27 = _outcollector9[1];
        _1358_v72 = _out26;
        _1359_v73 = _out27;
        let _1360_v74;
        _1360_v74 = _dafny.Set.fromElements(_1271_v12);
        _1360_v74 = _1360_v74;
      } else {
        let _1361___mcc_h11 = (_source17).cf6;
        let _1362_cf6 = _1361___mcc_h11;
        _1271_v12 = p0;
        if (!(!(!((_1271_v12).isLessThan(_1271_v12))) || (!(p0).isEqualTo((_dafny.ZERO).minus(_1271_v12))))) {
          r1 = (_this).f4;
          let _1363_v75;
          _1363_v75 = _dafny.Set.fromElements(_this.f3);
          let _1364_v76;
          _1364_v76 = _dafny.Set.fromElements(!(_1271_v12).isEqualTo((_dafny.ZERO).minus(_1271_v12)), !(_1363_v75).contains(!((_this).f4)), (_this).f4);
          let _1365_v77;
          _1365_v77 = _dafny.MultiSet.fromElements((((_1276_v16).contains((_this).f4)) ? ((_1276_v16).get((_this).f4)) : ((_this).f4)));
          _1364_v76 = _dafny.Set.fromElements((_this).f4, _this.f3, _this.f3, (_1365_v77).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1275_v15)), (_this).f4);
          let _1366_v78;
          let _init35 = function (_1367_i4) {
            return _dafny.Set.fromElements(new BigNumber(709));
          };
          let _nw197 = Array((new BigNumber(29)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw197.length); _i0_35++) {
            _nw197[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1366_v78 = _nw197;
          let _1368_v79;
          _1368_v79 = _module.D15.create_DC32(_1366_v78);
          let _1369_v80;
          _1369_v80 = _dafny.Seq.of(_1366_v78, _1366_v78, _1366_v78);
          let _1370_v81;
          let _nw198 = Array((new BigNumber(26)).toNumber());
          _nw198[(_dafny.ZERO).toNumber()] = (((_this).f4) ? (_1366_v78) : (_1366_v78));
          _nw198[(_dafny.ONE).toNumber()] = (_1368_v79).dtor_cf46;
          _nw198[(new BigNumber(2)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(3)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(4)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(5)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(6)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(7)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(8)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(9)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(10)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(11)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(12)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(13)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(14)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(15)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(16)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(17)).toNumber()] = (_1369_v80)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), ((_1371_p0) => function (_1372_i5) {
            return _1371_p0;
          })(p0)))).cardinality()), new BigNumber((_1369_v80).length))];
          _nw198[(new BigNumber(18)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(19)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(20)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(21)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(22)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(23)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(24)).toNumber()] = _1366_v78;
          _nw198[(new BigNumber(25)).toNumber()] = (((_this).f4) ? (_1366_v78) : (_1366_v78));
          _1370_v81 = _nw198;
          let _1373_v82;
          _1373_v82 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1366_v78);
          let _index211 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1370_v81).length));
          (_1370_v81)[_index211] = (((_1373_v82).contains((_this).f4)) ? ((_1373_v82).get((_this).f4)) : (_1366_v78));
          let _1374_v83;
          let _nw199 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
          _1374_v83 = _nw199;
          let _1375_v84;
          _1375_v84 = _dafny.Set.fromElements(p0);
          let _1376_v85;
          _1376_v85 = _dafny.Seq.of(new BigNumber((_module.__default.fm34(p0, p0, globalState)).length));
          let _index212 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1374_v83).length));
          (_1374_v83)[_index212] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1375_v84).length), (_dafny.ZERO).minus(_1271_v12), p0, new BigNumber(33), new BigNumber((_1275_v15).length)), _1376_v85);
          let _index213 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1370_v81).length));
          let _index214 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1374_v83).length));
          let _rhs185 = _1366_v78;
          let _rhs186 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(959)), ((_1377_p0) => function (_1378_i6) {
            return _1377_p0;
          })(p0));
          let _lhs126 = _1370_v81;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1370_v81).length));
          let _lhs128 = _1374_v83;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1374_v83).length));
          _lhs126[_lhs127] = _rhs185;
          _lhs128[_lhs129] = _rhs186;
          let _init36 = ((_1379_p0) => function (_1380_i7) {
            return (_1380_i7).multipliedBy(_1379_p0);
          })(p0);
          let _nw200 = Array((new BigNumber(9)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw200.length); _i0_36++) {
            _nw200[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1278_v19 = _nw200;
          let _1381_v86;
          let _nw201 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1381_v86 = _nw201;
          _1381_v86 = _1381_v86;
        } else {
          _1362_cf6 = _1362_cf6;
          _1272_v13 = _dafny.Seq.Concat(_1272_v13, _dafny.Seq.update(_1272_v13, _module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_1275_v15).length))).update(_this.f3, p0)).length), new BigNumber((_1272_v13).length)), _1296_v33));
          let _index215 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index215] = (p0).multipliedBy(p0);
          let _1382_v87;
          _1382_v87 = _module.D14.create_DC31((_dafny.ZERO).minus(p0), p0);
          let _index216 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_1278_v19).length));
          (_1278_v19)[_index216] = (_1382_v87).dtor_cf44;
          r1 = true;
          let _1383_v88;
          let _nw202 = Array((new BigNumber(17)).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = _1277_v17;
          _nw202[(_dafny.ONE).toNumber()] = (_1277_v17).Merge(_1277_v17);
          _nw202[(new BigNumber(2)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0));
          _nw202[(new BigNumber(4)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(5)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(6)).toNumber()] = (_1277_v17).update((_this).f4, (_1278_v19)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_1278_v19).length))]);
          _nw202[(new BigNumber(7)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(8)).toNumber()] = (_1277_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_1271_v12));
          _nw202[(new BigNumber(9)).toNumber()] = _module.__default.fm28(false, _this.f3, _this.f3, new BigNumber((_1272_v13).length), globalState);
          _nw202[(new BigNumber(10)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(11)).toNumber()] = (_1277_v17).update((_this).f4, new BigNumber(848));
          _nw202[(new BigNumber(12)).toNumber()] = (_module.D16.create_DC35(_1277_v17)).dtor_cf50;
          _nw202[(new BigNumber(13)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(14)).toNumber()] = _1277_v17;
          _nw202[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_1278_v19)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_1278_v19).length))])).update((_this).f4, (_1278_v19)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_1278_v19).length))]);
          _nw202[(new BigNumber(16)).toNumber()] = _1277_v17;
          _1383_v88 = _nw202;
          let _index217 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1383_v88).length));
          (_1383_v88)[_index217] = _1277_v17;
          let _index218 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1383_v88).length));
          (_1383_v88)[_index218] = _1277_v17;
        }
        let _1384_v89;
        _1384_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1271_v12, p0),_dafny.MultiSet.fromElements(_this.f3));
        let _index219 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1278_v19).length));
        (_1278_v19)[_index219] = _1271_v12;
        let _index220 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_1278_v19).length));
        (_1278_v19)[_index220] = p0;
        let _index221 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1278_v19).length));
        let _index222 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_1278_v19).length));
        let _rhs187 = new BigNumber(954);
        let _rhs188 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1275_v15).length),_dafny.MultiSet.fromElements(true, (_this).f4));
        let _rhs189 = (_1271_v12).multipliedBy((_1271_v12).multipliedBy(p0));
        let _rhs190 = ((_1271_v12).multipliedBy(p0)).minus((p0).plus(_1271_v12));
        let _rhs191 = _1271_v12;
        let _lhs130 = _1278_v19;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1278_v19).length));
        let _lhs132 = _1278_v19;
        let _lhs133 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_1278_v19).length));
        _1271_v12 = _rhs187;
        _1384_v89 = _rhs188;
        _lhs130[_lhs131] = _rhs189;
        _lhs132[_lhs133] = _rhs190;
        _1271_v12 = _rhs191;
        let _1385_v90;
        let _nw203 = new _module.C0();
        _nw203.__ctor();
        _1385_v90 = _nw203;
      }
      r0 = _module.__default.fm14(globalState);
      let _1386_v91;
      _1386_v91 = _dafny.Set.fromElements((_this).f4);
      r1 = !((_1386_v91).IsDisjointFrom(_1386_v91)) || ((p0).isLessThanOrEqualTo(p0));
      return [r0, r1];
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = [];
      let _1387_i0;
      _1387_i0 = _dafny.ZERO;
      L6: {
        while ((_module.__default.safeModuloInt(p1, p0)).isLessThanOrEqualTo((_this).fm0(globalState))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1387_i0)) {
              break L6;
            }
            _1387_i0 = (_1387_i0).plus(_dafny.ONE);
            let _1388_v0;
            _1388_v0 = _dafny.Seq.of((_this).f4, _this.f3);
            let _1389_v1;
            _1389_v1 = _module.D1.create_DC5((_this).f4, (_this).f4, p1, p1);
            let _1390_v2;
            let _nw204 = Array((new BigNumber(7)).toNumber());
            _nw204[(_dafny.ZERO).toNumber()] = p0;
            _nw204[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p1);
            _nw204[(new BigNumber(2)).toNumber()] = (((_this).f4) ? (p1) : (new BigNumber(631)));
            _nw204[(new BigNumber(3)).toNumber()] = p0;
            _nw204[(new BigNumber(4)).toNumber()] = (_module.__default.fm31(false, p0, (_this).f4, _1388_v0, globalState)).multipliedBy(p0);
            _nw204[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1388_v0, _1388_v0)).length);
            _nw204[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1389_v1,(_this).f4)).length);
            _1390_v2 = _nw204;
            let _1391_v3;
            _1391_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber(83));
            let _index223 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1390_v2).length));
            (_1390_v2)[_index223] = (new BigNumber((_1391_v3).length)).multipliedBy(new BigNumber(84));
            let _index224 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1390_v2).length));
            (_1390_v2)[_index224] = p0;
            (_this).f3 = !(!(false)) || ((_this).f4);
            let _1392_v4;
            let _nw205 = new _module.C2();
            _nw205.__ctor((_this).f4, _this.f3, (_this).f4, (_this).f4);
            _1392_v4 = _nw205;
            let _1393_v5;
            _1393_v5 = _dafny.Set.fromElements(new BigNumber((_1391_v3).length), p1);
            let _1394_v6;
            _1394_v6 = _1393_v5;
            _1394_v6 = _1394_v6;
          }
        }
      }
      let _1395_v7;
      _1395_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(260),_this.f3);
      let _1396_v8;
      _1396_v8 = _dafny.Seq.UnicodeFromString("ihae");
      let _1397_v9;
      _1397_v9 = _dafny.Seq.of(false);
      _1395_v7 = (_1395_v7).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm31(_this.f3, new BigNumber((_1396_v8).length), _this.f3, _1397_v9, globalState),true)).Merge(function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of (_dafny.Set.fromElements(p0, p0, p1, new BigNumber((_1396_v8).length), new BigNumber((_1396_v8).length))).Elements) {
          let _1398_v10 = _compr_48;
          if ((_dafny.Set.fromElements(p0, p0, p1, new BigNumber((_1396_v8).length), new BigNumber((_1396_v8).length))).contains(_1398_v10)) {
            _coll48.push([(_1398_v10).plus(p0),_this.f3]);
          }
        }
        return _coll48;
      }()));
      let _1399_v11;
      _1399_v11 = new BigNumber(-243);
      _1399_v11 = _module.__default.safeDivisionInt(_1399_v11, (new BigNumber(-33)).multipliedBy(p1));
      let _1400_v12;
      let _init37 = function (_1401_i2) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,_this.f3);
      };
      let _nw206 = Array((new BigNumber(14)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw206.length); _i0_37++) {
        _nw206[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1400_v12 = _nw206;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1400_v12).length))) {
        let _1402_i1 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1402_i1)) && ((_1402_i1).isLessThan(new BigNumber((_1400_v12).length))))) {
          (_1400_v12)[(_1402_i1)] = (((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), ((_1403_p1) => function (_1404_i3) {
            return _1403_p1;
          })(p1)))).IsSubsetOf(_dafny.MultiSet.fromElements(_1399_v11))) ? ((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,false))) : ((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3))));
        }
      }
      let _1405_v13;
      _1405_v13 = _dafny.Seq.of(p0, (_this).fm0(globalState), p1, _1399_v11, _1399_v11);
      let _1406_v14;
      _1406_v14 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber((_1405_v13).length), new BigNumber(55)));
      let _1407_v15;
      _1407_v15 = _dafny.Set.fromElements(new BigNumber(-890));
      _1406_v14 = ((_1406_v14).Difference(_module.__default.fm24(globalState))).update(new BigNumber((_1407_v15).length), _module.__default.abs(new BigNumber(498)));
      if (_this.f3) {
        _1399_v11 = (p1).minus(_1399_v11);
        let _1408_v16;
        let _init38 = function (_1409_i4) {
          return _this.f3;
        };
        let _nw207 = Array((_dafny.ONE).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw207.length); _i0_38++) {
          _nw207[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1408_v16 = _nw207;
        let _index225 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
        (_1408_v16)[_index225] = _module.__default.fm14(globalState);
        let _1410_v17;
        _1410_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
        let _1411_v18;
        _1411_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(664),new BigNumber((_1410_v17).length));
        let _1412_v19;
        _1412_v19 = _dafny.MultiSet.fromElements(_1411_v18, _1411_v18);
        let _index226 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
        (_1408_v16)[_index226] = _module.__default.fm21((p0).multipliedBy(p0), (_dafny.MultiSet.fromElements(_1411_v18)).Intersect(_1412_v19), globalState);
        if ((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))]) {
          let _1413_v20;
          let _init39 = ((_1414_v9) => function (_1415_i5) {
            return _1414_v9;
          })(_1397_v9);
          let _nw208 = Array((new BigNumber(23)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw208.length); _i0_39++) {
            _nw208[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1413_v20 = _nw208;
          let _index227 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1413_v20).length));
          (_1413_v20)[_index227] = _1397_v9;
          let _index228 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1413_v20).length));
          (_1413_v20)[_index228] = _1397_v9;
          let _1416_v21;
          _1416_v21 = _dafny.MultiSet.fromElements((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))]);
          (globalState).f0 = ((((((_1416_v21).contains((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))])) ? ((_1416_v21).get((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))])) : (_1399_v11))).isEqualTo(new BigNumber((_1397_v9).length))) ? (true) : ((_this).f4));
          _1408_v16 = _1408_v16;
          (globalState).f0 = ((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))]) || (_module.__default.fm14(globalState));
          _1399_v11 = new BigNumber(808);
        } else {
          let _1417_v23;
          _1417_v23 = _dafny.Set.fromElements((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))]);
          let _1418_v24;
          _1418_v24 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1419_v25;
          _1419_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1418_v24,_1417_v23);
          let _1420_v26;
          _1420_v26 = _dafny.MultiSet.fromElements(_1405_v13, _1405_v13, _1405_v13, _1405_v13);
          (globalState).f0 = (((function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_dafny.Seq.of(new BigNumber(206))).Elements) {
              let _1421_v22 = _compr_49;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(206)), _1421_v22)) {
                _coll49.push([_module.__default.safeDivisionInt(_1421_v22, (_dafny.ZERO).minus(p0)),!(_this.f3)]);
              }
            }
            return _coll49;
          }()).contains(p1)) ? (((((_1419_v25).contains(_1418_v24)) ? ((_1419_v25).get(_1418_v24)) : (_1417_v23))).IsSubsetOf(_1417_v23)) : ((_1420_v26).IsDisjointFrom(_1420_v26)));
          let _index229 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
          let _index230 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
          let _rhs192 = (new BigNumber(734)).plus(p1);
          let _rhs193 = (_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))];
          let _rhs194 = (_this).f4;
          let _rhs195 = _1418_v24;
          let _lhs134 = _1408_v16;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
          let _lhs136 = _1408_v16;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
          _1399_v11 = _rhs192;
          _lhs134[_lhs135] = _rhs193;
          _lhs136[_lhs137] = _rhs194;
          _1418_v24 = _rhs195;
          let _1422_v27;
          let _init40 = ((_1423_v11) => function (_1424_i6) {
            return (_1424_i6).multipliedBy(_1423_v11);
          })(_1399_v11);
          let _nw209 = Array((new BigNumber(15)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw209.length); _i0_40++) {
            _nw209[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1422_v27 = _nw209;
          let _index231 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1422_v27).length));
          (_1422_v27)[_index231] = (_dafny.ZERO).minus(p0);
          let _1425_v28;
          _1425_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1406_v14).IsProperSubsetOf((_1406_v14).update(_1399_v11, _module.__default.abs(new BigNumber(211)))),_1399_v11);
          let _index232 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1422_v27).length));
          (_1422_v27)[_index232] = new BigNumber((_1425_v28).length);
          let _index233 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length));
          (_1408_v16)[_index233] = _this.f3;
          let _index234 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1422_v27).length));
          (_1422_v27)[_index234] = ((_1422_v27)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1422_v27).length))]).minus((_1422_v27)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1422_v27).length))]);
        }
        let _1426_v29;
        let _nw210 = new _module.C0();
        _nw210.__ctor();
        _1426_v29 = _nw210;
        let _1427_v30;
        let _nw211 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
        _1427_v30 = _nw211;
        let _1428_v31;
        _1428_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1427_v30);
        _1427_v30 = (((_1408_v16)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1408_v16).length))]) ? ((((_1428_v31).contains(_this.f3)) ? ((_1428_v31).get(_this.f3)) : (_1427_v30))) : ((((_1428_v31).contains(true)) ? ((_1428_v31).get(true)) : (_1427_v30))));
      } else {
        let _1429_v32;
        _1429_v32 = new _dafny.CodePoint('i'.codePointAt(0));
        _1429_v32 = _1429_v32;
        _1399_v11 = p0;
        let _1430_v33;
        let _nw212 = Array((new BigNumber(4)).toNumber()).fill(false);
        _1430_v33 = _nw212;
        let _index235 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length));
        (_1430_v33)[_index235] = _this.f3;
        let _1431_v34;
        _1431_v34 = _dafny.Set.fromElements((_this).f4);
        let _1432_v35;
        _1432_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_1431_v34);
        let _1433_v36;
        _1433_v36 = _module.D7.create_DC16((_1432_v35).Merge(_1432_v35));
        let _1434_v37;
        _1434_v37 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("lhxyt"));
        let _index236 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length));
        let _rhs196 = ((_1434_v37).Union(_1434_v37)).equals(_1434_v37);
        let _rhs197 = new BigNumber(145);
        let _rhs198 = _1433_v36;
        let _lhs138 = _1430_v33;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length));
        _lhs138[_lhs139] = _rhs196;
        _1399_v11 = _rhs197;
        _1433_v36 = _rhs198;
        let _1435_v38;
        _1435_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1397_v9,p1);
        let _1436_v39;
        _1436_v39 = _module.D7.create_DC17(_module.__default.fm31(_this.f3, new BigNumber((_1435_v38).length), _this.f3, _1397_v9, globalState), !(_this.f3), (_1430_v33)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length))], new BigNumber((_dafny.MultiSet.fromElements((_1430_v33)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length))])).cardinality()));
        let _1437_v40;
        let _nw213 = new _module.C2();
        _nw213.__ctor((_1430_v33)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1430_v33).length))], false, _this.f3, (_1436_v39).dtor_cf27);
        _1437_v40 = _nw213;
        let _1438_v41;
        _1438_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1437_v40);
        let _1439_v42;
        _1439_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1399_v11);
        let _1440_v43;
        _1440_v43 = _dafny.Map.Empty.slice().updateUnsafe((((_1438_v41).contains(p0)) ? ((_1438_v41).get(p0)) : (_1437_v40)),_1439_v42);
        _1440_v43 = (_1440_v43).update(_1437_v40, _1439_v42);
        let _1441_v44;
        let _nw214 = new _module.C0();
        _nw214.__ctor();
        _1441_v44 = _nw214;
        _1441_v44 = _1441_v44;
      }
      r0 = _1397_v9;
      let _1442_v45;
      let _nw215 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1442_v45 = _nw215;
      r1 = _1442_v45;
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f3 = false;
      this._f4 = false;
      this.f8 = false;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f8, f9, f3, f4) {
      let _this = this;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (_this).f9;
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (_module.D0.create_DC2((_this).f4, (_this).f9, (_this).f9)).dtor_cf4;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1443_v0;
      _1443_v0 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1444_v1;
      _1444_v1 = _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), _1443_v0, _1443_v0, _1443_v0);
      if (((p1).f4) === (_module.__default.fm6((_this).fm5(false, _1444_v1, globalState), globalState))) {
        let _1445_v2;
        _1445_v2 = _module.D0.create_DC2(!((p1).f4), (_this).f9, p0);
        let _1446_v3;
        _1446_v3 = _dafny.Seq.of(_1445_v2, ((_this.f3) ? (_1445_v2) : (_module.D0.create_DC2(_this.f3, p0, new BigNumber((_dafny.Seq.UnicodeFromString("irpitaca")).length)))), _1445_v2, _1445_v2);
        _1446_v3 = _dafny.Seq.Concat(_1446_v3, _1446_v3);
        (globalState).f0 = true;
        let _1447_v4;
        _1447_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _1448_v5;
        _1448_v5 = _dafny.MultiSet.fromElements(_1447_v4);
        let _1449_v6;
        let _nw216 = new _module.C3();
        _nw216.__ctor(_module.__default.fm21((_this).f9, _1448_v5, globalState), !(!(((!(p1.f3)) ? (p1.f3) : (false)))));
        _1449_v6 = _nw216;
        r0 = p0;
        let _1450_v7;
        let _nw217 = new _module.C3();
        _nw217.__ctor(_this.f8, (p1).f4);
        _1450_v7 = _nw217;
        _1450_v7 = p1;
      } else {
        let _1451_v8;
        _1451_v8 = _dafny.Seq.of((p1).f4, p1.f3, p1.f3);
        let _1452_v9;
        _1452_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,p1.f3);
        r0 = _module.__default.fm31(false, (_this).f9, (_1451_v8)[_module.__default.safeIndex(new BigNumber((_1452_v9).length), new BigNumber((_1451_v8).length))], _1451_v8, globalState);
        let _1453_v10;
        let _init41 = ((_1454_v1, _1455_v0, _1456_p0) => function (_1457_i0) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1454_v1, _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1454_v1).length)), _1455_v0),_1456_p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1454_v1,(_this).f9));
        })(_1444_v1, _1443_v0, p0);
        let _nw218 = Array((new BigNumber(10)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw218.length); _i0_41++) {
          _nw218[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1453_v10 = _nw218;
        let _index237 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1453_v10).length));
        (_1453_v10)[_index237] = _dafny.Map.Empty.slice().updateUnsafe(_1444_v1,new BigNumber(608));
        let _1458_v11;
        _1458_v11 = _dafny.Set.fromElements(_this.f8, p1.f3, (p1).f4, _this.f8);
        let _1459_v12;
        _1459_v12 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_1458_v11);
        let _1460_v13;
        _1460_v13 = _module.D7.create_DC16(_1459_v12);
        let _1461_v14;
        let _init42 = ((_1462_p1, _1463_v1) => function (_1464_i1) {
          return (((_1462_p1).f4) ? (_dafny.Map.Empty.slice().updateUnsafe(_1462_p1.f3,_1463_v1)) : (_dafny.Map.Empty.slice().updateUnsafe((_1462_p1).f4,_1463_v1)));
        })(p1, _1444_v1);
        let _nw219 = Array((new BigNumber(14)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw219.length); _i0_42++) {
          _nw219[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1461_v14 = _nw219;
        let _1465_v15;
        _1465_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,_dafny.Seq.UnicodeFromString("oh"));
        let _index238 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1461_v14).length));
        (_1461_v14)[_index238] = _1465_v15;
        let _1466_v16;
        let _nw220 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1466_v16 = _nw220;
        let _1467_v17;
        _1467_v17 = _module.D4.create_DC12((_this).f9, (_this).f9, (_this).f4);
        let _index239 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1466_v16).length));
        (_1466_v16)[_index239] = (_1467_v17).dtor_cf20;
        let _1468_v18;
        _1468_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jnwad"), _1444_v1),(p0).minus((_this).f9));
        let _1469_v19;
        _1469_v19 = _dafny.MultiSet.fromElements((_this).f4, (p1).f4);
        let _1470_v20;
        _1470_v20 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(new BigNumber((_1469_v19).cardinality())), (_this).f9, new BigNumber(145), (_this).f9);
        let _1471_v21;
        _1471_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _1472_v22;
        _1472_v22 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,_1471_v21);
        let _1473_v23;
        _1473_v23 = _dafny.Seq.of(_1444_v1, _dafny.Seq.UnicodeFromString("uwnoak"));
        let _pat_let_tv40 = _1460_v13;
        let _index240 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1453_v10).length));
        let _index241 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1461_v14).length));
        let _index242 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1466_v16).length));
        let _rhs199 = _1468_v18;
        let _rhs200 = function (_pat_let28_0) {
          return function (_1474_dt__update__tmp_h0) {
            return function (_pat_let29_0) {
              return function (_1475_dt__update_hcf25_h0) {
                return _module.D7.create_DC16(_1475_dt__update_hcf25_h0);
              }(_pat_let29_0);
            }((_pat_let_tv40).dtor_cf25);
          }(_pat_let28_0);
        }(_1460_v13);
        let _rhs201 = !(_1469_v19).contains((p1).f4);
        let _rhs202 = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(new BigNumber((_1458_v11).length), globalState),_1444_v1)).Merge(_1465_v15);
        let _rhs203 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_1476_i2) {
          return (_this).f9;
        }), _dafny.Seq.Concat(_1470_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_1477_i3) {
          return (_this).f9;
        }))), _module.__default.safeIndex((_this).fm5((_this).f4, _1444_v1, globalState), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_1478_i2) {
          return (_this).f9;
        }), _dafny.Seq.Concat(_1470_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_1479_i3) {
          return (_this).f9;
        })))).length)), (new BigNumber((_1472_v22).length)).multipliedBy(new BigNumber(((_1473_v23)).length))), _module.__default.safeIndex(((_this).f9).multipliedBy(p0), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_1480_i2) {
          return (_this).f9;
        }), _dafny.Seq.Concat(_1470_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_1481_i3) {
          return (_this).f9;
        }))), _module.__default.safeIndex((_this).fm5((_this).f4, _1444_v1, globalState), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_1482_i2) {
          return (_this).f9;
        }), _dafny.Seq.Concat(_1470_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_1483_i3) {
          return (_this).f9;
        })))).length)), (new BigNumber((_1472_v22).length)).multipliedBy(new BigNumber(((_1473_v23)).length)))).length)), p0)).length);
        let _lhs140 = _1453_v10;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1453_v10).length));
        let _lhs142 = p1;
        let _lhs143 = _1461_v14;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1461_v14).length));
        let _lhs145 = _1466_v16;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1466_v16).length));
        _lhs140[_lhs141] = _rhs199;
        _1460_v13 = _rhs200;
        _lhs142.f3 = _rhs201;
        _lhs143[_lhs144] = _rhs202;
        _lhs145[_lhs146] = _rhs203;
        _1444_v1 = ((!_dafny.Seq.contains(_1444_v1, _1443_v0)) ? (_dafny.Seq.UnicodeFromString("vrmn")) : (_1444_v1));
        (globalState).f0 = (p0).isLessThan(_module.__default.safeModuloInt((_this).f9, p0));
        let _1484_v24;
        _1484_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1466_v16)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1466_v16).length))],p1.f3);
        let _1485_v25;
        _1485_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f4) ? ((_1484_v24).update(p0, (p1).f4)) : (_1484_v24)),(_this).f9);
        _1485_v25 = (_1485_v25).Merge(_1485_v25);
      }
      r0 = new BigNumber((_1444_v1).length);
      let _hi12 = ((_this).f9).plus(new BigNumber(91));
      for (let _1486_i4 = p0; _1486_i4.isLessThan(_hi12); _1486_i4 = _1486_i4.plus(_dafny.ONE)) {
        if (true) {
          let _1487_v26;
          let _init43 = ((_1488_p1) => function (_1489_i5) {
            return (_1488_p1).f4;
          })(p1);
          let _nw221 = Array((new BigNumber(29)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw221.length); _i0_43++) {
            _nw221[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1487_v26 = _nw221;
          let _index243 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1487_v26).length));
          (_1487_v26)[_index243] = p1.f3;
          let _1490_v27;
          let _init44 = ((_1491_v0) => function (_1492_i6) {
            return _1491_v0;
          })(_1443_v0);
          let _nw222 = Array((new BigNumber(25)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw222.length); _i0_44++) {
            _nw222[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1490_v27 = _nw222;
          let _1493_v28;
          _1493_v28 = _dafny.MultiSet.fromElements(_this.f3, !(p1.f3), (_this).f4, _this.f8);
          let _1494_v29;
          _1494_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1490_v27,(_1493_v28).equals(_1493_v28));
          let _index244 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1487_v26).length));
          (_1487_v26)[_index244] = (((_1494_v29).contains(_1490_v27)) ? ((_1494_v29).get(_1490_v27)) : (!((_this).f4)));
          _1493_v28 = _dafny.MultiSet.fromElements((p1).f4, (_1487_v26)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1487_v26).length))], _this.f8);
          let _1495_v30;
          _1495_v30 = _dafny.Seq.of(_this.f3, p1.f3);
          let _1496_v31;
          _1496_v31 = _dafny.Seq.of(_this.f8, (_1495_v30)[_module.__default.safeIndex(p0, new BigNumber((_1495_v30).length))]);
          _1496_v31 = _1496_v31;
          (_this).f3 = (p1).f4;
          let _1497_v32;
          let _nw223 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _1497_v32 = _nw223;
          let _index245 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1497_v32).length));
          (_1497_v32)[_index245] = p0;
          let _1498_v33;
          _1498_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_1444_v1)[_module.__default.safeIndex(_1486_i4, new BigNumber((_1444_v1).length))]);
          let _1499_v34;
          _1499_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_1498_v33).contains((p1).f4)) ? ((_1498_v33).get((p1).f4)) : (new _dafny.CodePoint('a'.codePointAt(0)))),_dafny.Seq.Concat(_1444_v1, _1444_v1));
          let _index246 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1497_v32).length));
          let _rhs204 = (p1).f4;
          let _rhs205 = new BigNumber(175);
          let _rhs206 = new BigNumber((_1499_v34).length);
          let _rhs207 = _dafny.Seq.UnicodeFromString("b");
          let _lhs147 = globalState;
          let _lhs148 = _1497_v32;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1497_v32).length));
          _lhs147.f0 = _rhs204;
          r0 = _rhs205;
          _lhs148[_lhs149] = _rhs206;
          _1444_v1 = _rhs207;
        } else {
          r0 = (p1).fm0(globalState);
          let _1500_v35;
          let _init45 = function (_1501_i7) {
            return _dafny.Set.fromElements((_this).f9, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3)).length));
          };
          let _nw224 = Array((new BigNumber(17)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw224.length); _i0_45++) {
            _nw224[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1500_v35 = _nw224;
          let _1502_v36;
          _1502_v36 = _module.D15.create_DC32(_1500_v35);
          let _1503_v37;
          _1503_v37 = _module.D15.create_DC34(_1502_v36);
          let _1504_v38;
          _1504_v38 = _module.D15.create_DC34(_1503_v37);
          _1504_v38 = _1504_v38;
          r0 = p0;
          let _1505_v39;
          _1505_v39 = _dafny.Seq.of((_this).f4, p1.f3, p1.f3, (p1).f4);
          let _1506_v40;
          _1506_v40 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1443_v0,(p1).f4)).length), new BigNumber((_1505_v39).length), (_this).f9);
          let _1507_v41;
          _1507_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          let _1508_v42;
          let _nw225 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1508_v42 = _nw225;
          let _1509_v43;
          _1509_v43 = _dafny.Seq.of(_1508_v42);
          let _1510_v44;
          _1510_v44 = _dafny.Seq.of((_this).f9);
          let _1511_v45;
          let _nw226 = Array((new BigNumber(24)).toNumber());
          _nw226[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f9);
          _nw226[(_dafny.ONE).toNumber()] = (new BigNumber((_1506_v40).cardinality())).multipliedBy(new BigNumber(453));
          _nw226[(new BigNumber(2)).toNumber()] = p0;
          _nw226[(new BigNumber(3)).toNumber()] = p0;
          _nw226[(new BigNumber(4)).toNumber()] = p0;
          _nw226[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(6)).toNumber()] = p0;
          _nw226[(new BigNumber(7)).toNumber()] = new BigNumber(-944);
          _nw226[(new BigNumber(8)).toNumber()] = (new BigNumber(395)).plus(new BigNumber((_1507_v41).length));
          _nw226[(new BigNumber(9)).toNumber()] = new BigNumber((_module.__default.fm34(_1486_i4, (_this).f9, globalState)).length);
          _nw226[(new BigNumber(10)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(11)).toNumber()] = _1486_i4;
          _nw226[(new BigNumber(12)).toNumber()] = (new BigNumber((_1509_v43).length)).plus(p0);
          _nw226[(new BigNumber(13)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(14)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt((_this).f9, new BigNumber(-732));
          _nw226[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(98), (_this).f9);
          _nw226[(new BigNumber(17)).toNumber()] = _1486_i4;
          _nw226[(new BigNumber(18)).toNumber()] = new BigNumber(660);
          _nw226[(new BigNumber(19)).toNumber()] = p0;
          _nw226[(new BigNumber(20)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(21)).toNumber()] = _1486_i4;
          _nw226[(new BigNumber(22)).toNumber()] = (_this).f9;
          _nw226[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1505_v39, _module.__default.fm34(new BigNumber((_1510_v44).length), p0, globalState))).length);
          _1511_v45 = _nw226;
          let _index247 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v42).length));
          (_1508_v42)[_index247] = ((p1.f3) ? (p0) : (_1486_i4));
          let _index248 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v42).length));
          (_1508_v42)[_index248] = p0;
          let _1512_v46;
          _1512_v46 = _dafny.MultiSet.fromElements(false);
          let _1513_v47;
          _1513_v47 = _dafny.Seq.of(_1505_v39);
          let _1514_v48;
          _1514_v48 = _dafny.Seq.of(_1510_v44, _dafny.Seq.update(_1510_v44, _module.__default.safeIndex((_1508_v42)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v42).length))], new BigNumber((_1510_v44).length)), (_1508_v42)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_1508_v42).length))]), _1510_v44, _dafny.Seq.of(p0));
          let _1515_v49;
          _1515_v49 = _dafny.Map.Empty.slice().updateUnsafe((_1486_i4).minus((((_1512_v46).contains(_this.f8)) ? ((_1512_v46).get(_this.f8)) : (_module.__default.fm31(_this.f3, (_dafny.ZERO).minus(new BigNumber((_1505_v39).length)), !((p1).f4), (_1513_v47)[_module.__default.safeIndex(p0, new BigNumber((_1513_v47).length))], globalState)))),_1514_v48);
          _1515_v49 = (_1515_v49).update(_1486_i4, _dafny.Seq.of(_1510_v44, _1510_v44, _1510_v44));
        }
        if ((_module.__default.fm37(_1486_i4, globalState)).dtor_cf53) {
          let _1516_v50;
          let _nw227 = new _module.C2();
          _nw227.__ctor(p1.f3, p1.f3, _this.f3, true);
          _1516_v50 = _nw227;
          let _1517_v51;
          _1517_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(p1).f4);
          let _1518_v52;
          _1518_v52 = _module.D0.create_DC2((_1516_v50).f10, p0, (_this).f9);
          let _1519_v53;
          _1519_v53 = _dafny.Seq.of(_this.f3, (_1518_v52).dtor_cf3);
          let _1520_v54;
          let _nw228 = new _module.C1();
          _nw228.__ctor(_1444_v1, (((((_1517_v51).contains(p1.f3)) ? ((_1517_v51).get(p1.f3)) : ((_this).f4))) ? ((_this).f4) : ((_1519_v53)[_module.__default.safeIndex((_this).f9, new BigNumber((_1519_v53).length))])));
          _1520_v54 = _nw228;
          let _1521_v55;
          let _nw229 = Array((new BigNumber(10)).toNumber());
          _1521_v55 = _nw229;
          _1521_v55 = _1521_v55;
          let _1522_v56;
          _1522_v56 = _dafny.Set.fromElements(false);
          let _1523_v57;
          _1523_v57 = _dafny.MultiSet.fromElements(_this.f3, (_1519_v53)[_module.__default.safeIndex(new BigNumber((_1522_v56).length), new BigNumber((_1519_v53).length))], _this.f8);
          let _1524_v58;
          _1524_v58 = _module.D12.create_DC26(_1522_v56);
          let _1525_v59;
          _1525_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_1524_v58);
          let _1526_v60;
          _1526_v60 = _dafny.Seq.of((((_1523_v57).contains((_this).f4)) ? ((_1523_v57).get((_this).f4)) : (_1486_i4)), new BigNumber((_1525_v59).length));
          let _1527_v61;
          _1527_v61 = _module.D1.create_DC4(new BigNumber(866));
          let _1528_v62;
          _1528_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1527_v61).dtor_cf7,p1);
          let _1529_v63;
          _1529_v63 = _dafny.Seq.of((((_1528_v62).contains(_1486_i4)) ? ((_1528_v62).get(_1486_i4)) : (p1)), p1);
          let _1530_v64;
          _1530_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1529_v63).length));
          let _1531_v65;
          _1531_v65 = _dafny.Seq.of(_1530_v64, _1530_v64);
          let _1532_v66;
          _1532_v66 = _dafny.Seq.of(_1531_v65, _1531_v65);
          let _1533_v67;
          let _nw230 = Array((new BigNumber(25)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = ((_this).f9).isLessThanOrEqualTo(new BigNumber(364));
          _nw230[(_dafny.ONE).toNumber()] = (_1520_v54).f13;
          _nw230[(new BigNumber(2)).toNumber()] = (_1520_v54).f13;
          _nw230[(new BigNumber(3)).toNumber()] = (p1).f4;
          _nw230[(new BigNumber(4)).toNumber()] = (_1516_v50).f10;
          _nw230[(new BigNumber(5)).toNumber()] = _module.__default.fm18(p0, _1443_v0, p0, globalState);
          _nw230[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1526_v60, _1526_v60);
          _nw230[(new BigNumber(7)).toNumber()] = (_this).f4;
          _nw230[(new BigNumber(8)).toNumber()] = !_dafny.Seq.contains((_1520_v54).f12, _1443_v0);
          _nw230[(new BigNumber(9)).toNumber()] = p1.f3;
          _nw230[(new BigNumber(10)).toNumber()] = (p1).f4;
          _nw230[(new BigNumber(11)).toNumber()] = _module.__default.fm21((_this).f9, _dafny.MultiSet.FromArray(_dafny.Seq.update((_1532_v66)[_module.__default.safeIndex(_1486_i4, new BigNumber((_1532_v66).length))], _module.__default.safeIndex(p0, new BigNumber(((_1532_v66)[_module.__default.safeIndex(_1486_i4, new BigNumber((_1532_v66).length))]).length)), _1530_v64)), globalState);
          _nw230[(new BigNumber(12)).toNumber()] = _this.f8;
          _nw230[(new BigNumber(13)).toNumber()] = (_1516_v50).f10;
          _nw230[(new BigNumber(14)).toNumber()] = !((p1).f4);
          _nw230[(new BigNumber(15)).toNumber()] = _this.f8;
          _nw230[(new BigNumber(16)).toNumber()] = p1.f3;
          _nw230[(new BigNumber(17)).toNumber()] = false;
          _nw230[(new BigNumber(18)).toNumber()] = (_this).f4;
          _nw230[(new BigNumber(19)).toNumber()] = (p1).f4;
          _nw230[(new BigNumber(20)).toNumber()] = (_this.f3) || (p1.f3);
          _nw230[(new BigNumber(21)).toNumber()] = !((_1516_v50).f11) || (_this.f8);
          _nw230[(new BigNumber(22)).toNumber()] = _module.__default.fm18(_1486_i4, new _dafny.CodePoint('d'.codePointAt(0)), new BigNumber((_dafny.Seq.of((p1).f4, (_1519_v53)[_module.__default.safeIndex(new BigNumber((_1517_v51).length), new BigNumber((_1519_v53).length))], (_1516_v50).f10)).length), globalState);
          _nw230[(new BigNumber(23)).toNumber()] = !((_1516_v50).f11);
          _nw230[(new BigNumber(24)).toNumber()] = !((p1).f4) || ((_1520_v54).f13);
          _1533_v67 = _nw230;
          _1533_v67 = _1533_v67;
          let _1534_v68;
          _1534_v68 = _module.D2.create_DC8(_1533_v67, _dafny.Seq.UnicodeFromString("kmomdi"));
          let _pat_let_tv41 = _1533_v67;
          let _pat_let_tv42 = _1533_v67;
          let _1535_v69;
          _1535_v69 = _dafny.Set.fromElements(_1534_v68, _1534_v68, function (_pat_let30_0) {
            return function (_1538_dt__update__tmp_h2) {
              return function (_pat_let33_0) {
                return function (_1539_dt__update_hcf13_h1) {
                  return _module.D2.create_DC8(_1539_dt__update_hcf13_h1, (_1538_dt__update__tmp_h2).dtor_cf14);
                }(_pat_let33_0);
              }(_pat_let_tv42);
            }(_pat_let30_0);
          }(function (_pat_let31_0) {
            return function (_1536_dt__update__tmp_h1) {
              return function (_pat_let32_0) {
                return function (_1537_dt__update_hcf13_h0) {
                  return _module.D2.create_DC8(_1537_dt__update_hcf13_h0, (_1536_dt__update__tmp_h1).dtor_cf14);
                }(_pat_let32_0);
              }(_pat_let_tv41);
            }(_pat_let31_0);
          }(_1534_v68)));
          _1535_v69 = _1535_v69;
        } else {
          let _1540_v70;
          let _nw231 = new _module.C1();
          _nw231.__ctor(_dafny.Seq.UnicodeFromString("ipej"), p1.f3);
          _1540_v70 = _nw231;
          let _1541_v71;
          _1541_v71 = _dafny.Set.fromElements(_1540_v70, _1540_v70, _1540_v70);
          (p1).f3 = (_1541_v71).IsSubsetOf(_1541_v71);
          let _1542_v72;
          _1542_v72 = _dafny.Set.fromElements(false);
          let _1543_v73;
          _1543_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1542_v72).IsSubsetOf(_dafny.Set.fromElements(p1.f3, (p1).f4)),(_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((_1540_v70).f12).length), (_this).f9)));
          _1543_v73 = (_1543_v73).update((_1540_v70).f13, new BigNumber((_1444_v1).length));
          let _1544_v74;
          _1544_v74 = _dafny.Seq.of((p1).f4, false, _this.f3);
          let _1545_v75;
          _1545_v75 = _module.D4.create_DC11(_1544_v74);
          let _1546_v76;
          _1546_v76 = _module.D4.create_DC13(_1545_v75);
          let _1547_v77;
          _1547_v77 = _module.D4.create_DC13(_1545_v75);
          _1547_v77 = _1547_v77;
          let _1548_v78;
          let _init46 = function (_1549_i8) {
            return _module.__default.safeDivisionInt(_1549_i8, new BigNumber(478));
          };
          let _nw232 = Array((new BigNumber(17)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw232.length); _i0_46++) {
            _nw232[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1548_v78 = _nw232;
          let _1550_v79;
          _1550_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1548_v78,(new BigNumber(-516)).multipliedBy(_1486_i4));
          _1550_v79 = (_1550_v79).update(_1548_v78, p0);
          let _1551_v80;
          let _nw233 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1551_v80 = _nw233;
          let _1552_v81;
          let _nw234 = Array((new BigNumber(13)).toNumber()).fill(false);
          _1552_v81 = _nw234;
          let _1553_v82;
          _1553_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1551_v80,_1552_v81);
          _1553_v82 = (_1553_v82).update(_1551_v80, _1552_v81);
        }
        let _1554_v84;
        _1554_v84 = _dafny.Set.fromElements(_1486_i4);
        let _1555_v85;
        _1555_v85 = _dafny.MultiSet.fromElements(function () {
          let _coll50 = new _dafny.Map();
          for (const _compr_50 of (_1554_v84).Elements) {
            let _1556_v83 = _compr_50;
            if ((_1554_v84).contains(_1556_v83)) {
              _coll50.push([_module.__default.safeModuloInt(_1556_v83, (_this).f9),(_this).f9]);
            }
          }
          return _coll50;
        }());
        let _1557_v86;
        let _nw235 = new _module.C3();
        _nw235.__ctor((p1).f4, _module.__default.fm21((_dafny.ZERO).minus(new BigNumber(-822)), _1555_v85, globalState));
        _1557_v86 = _nw235;
        let _1558_v87;
        let _nw236 = new _module.C1();
        _nw236.__ctor(_dafny.Seq.Concat(_1444_v1, _1444_v1), _this.f8);
        _1558_v87 = _nw236;
      }
      let _1559_v88;
      _1559_v88 = _dafny.Seq.of((_this).f4, _this.f8);
      let _1560_v89;
      _1560_v89 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1559_v88).length)), (_dafny.ZERO).minus(p0));
      let _1561_v90;
      _1561_v90 = _module.D4.create_DC12((_this).f9, new BigNumber((_1560_v89).length), _this.f3);
      r0 = (_1561_v90).dtor_cf19;
      let _1562_i9;
      _1562_i9 = _dafny.ZERO;
      L7: {
        while ((new BigNumber((_dafny.Seq.Concat(_1444_v1, _dafny.Seq.of(_1443_v0))).length)).isLessThan(((_this).f9).multipliedBy((_this).f9))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1562_i9)) {
              break L7;
            }
            _1562_i9 = (_1562_i9).plus(_dafny.ONE);
            let _1563_v91;
            _1563_v91 = _module.D7.create_DC17(p0, p1.f3, false, (_this).f9);
            let _1564_v92;
            _1564_v92 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(860));
            let _1565_v93;
            _1565_v93 = _dafny.MultiSet.fromElements(p1.f3, !(p1.f3));
            let _1566_v94;
            let _nw237 = Array((new BigNumber(14)).toNumber());
            _nw237[(_dafny.ZERO).toNumber()] = (_this).f9;
            _nw237[(_dafny.ONE).toNumber()] = p0;
            _nw237[(new BigNumber(2)).toNumber()] = ((_this).fm0(globalState)).plus((_1563_v91).dtor_cf29);
            _nw237[(new BigNumber(3)).toNumber()] = ((p1.f3) ? (p0) : (new BigNumber(-207)));
            _nw237[(new BigNumber(4)).toNumber()] = new BigNumber((_1564_v92).length);
            _nw237[(new BigNumber(5)).toNumber()] = ((_this).f9).plus(p0);
            _nw237[(new BigNumber(6)).toNumber()] = new BigNumber((_1565_v93).cardinality());
            _nw237[(new BigNumber(7)).toNumber()] = p0;
            _nw237[(new BigNumber(8)).toNumber()] = (((_1565_v93).contains(p1.f3)) ? ((_1565_v93).get(p1.f3)) : (p0));
            _nw237[(new BigNumber(9)).toNumber()] = (_this).f9;
            _nw237[(new BigNumber(10)).toNumber()] = (_this).f9;
            _nw237[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
            _nw237[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1.f3,_module.__default.fm14(globalState))).length));
            _nw237[(new BigNumber(13)).toNumber()] = (_this).f9;
            _1566_v94 = _nw237;
            let _index249 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_1566_v94).length));
            (_1566_v94)[_index249] = new BigNumber((_module.__default.fm19(globalState)).length);
            let _1567_v95;
            _1567_v95 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,(p1).f4);
            let _index250 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_1566_v94).length));
            let _rhs208 = true;
            let _rhs209 = new BigNumber(((_1567_v95).Merge(_dafny.Map.Empty.slice().updateUnsafe((p1).f4,(_this).f4))).length);
            let _rhs210 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yvhykh"), _dafny.Seq.UnicodeFromString("fccwodxf"));
            let _rhs211 = _module.__default.safeDivisionInt((((_1564_v92).contains(new BigNumber((_dafny.Seq.update(_1559_v88, _module.__default.safeIndex(new BigNumber((_1559_v88).length), new BigNumber((_1559_v88).length)), _module.__default.fm6(p0, globalState))).length))) ? ((_1564_v92).get(new BigNumber((_dafny.Seq.update(_1559_v88, _module.__default.safeIndex(new BigNumber((_1559_v88).length), new BigNumber((_1559_v88).length)), _module.__default.fm6(p0, globalState))).length))) : (p0)), new BigNumber((_1444_v1).length));
            let _rhs212 = _1443_v0;
            let _lhs150 = globalState;
            let _lhs151 = _1566_v94;
            let _lhs152 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_1566_v94).length));
            _lhs150.f0 = _rhs208;
            r0 = _rhs209;
            _1444_v1 = _rhs210;
            _lhs151[_lhs152] = _rhs211;
            _1443_v0 = _rhs212;
            let _1568_v96;
            _1568_v96 = _dafny.Set.fromElements(true, _this.f8);
            let _1569_v97;
            _1569_v97 = _module.D12.create_DC26(_1568_v96);
            let _1570_v98;
            _1570_v98 = _module.D12.create_DC28(_1569_v97);
            let _pat_let_tv43 = _1569_v97;
            let _1571_v99;
            let _nw238 = Array((new BigNumber(16)).toNumber());
            _nw238[(_dafny.ZERO).toNumber()] = ((_this.f8) ? (_module.D12.create_DC28(_1569_v97)) : (_1570_v98));
            _nw238[(_dafny.ONE).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(2)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(3)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(4)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(5)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(6)).toNumber()] = function (_pat_let34_0) {
              return function (_1572_dt__update__tmp_h3) {
                return function (_pat_let35_0) {
                  return function (_1573_dt__update_hcf41_h0) {
                    return _module.D12.create_DC28(_1573_dt__update_hcf41_h0);
                  }(_pat_let35_0);
                }(_pat_let_tv43);
              }(_pat_let34_0);
            }(_1570_v98);
            _nw238[(new BigNumber(7)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(8)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(9)).toNumber()] = _module.__default.fm38(p0, p1.f3, globalState);
            _nw238[(new BigNumber(10)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(11)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(12)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(13)).toNumber()] = _1570_v98;
            _nw238[(new BigNumber(14)).toNumber()] = _module.__default.fm38((_1566_v94)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_1566_v94).length))], !(p1.f3), globalState);
            _nw238[(new BigNumber(15)).toNumber()] = _module.__default.fm38((_1566_v94)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_1566_v94).length))], p1.f3, globalState);
            _1571_v99 = _nw238;
            let _index251 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1571_v99).length));
            (_1571_v99)[_index251] = _1570_v98;
            let _index252 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_1571_v99).length));
            (_1571_v99)[_index252] = _module.D12.create_DC28(_module.__default.fm38(p0, (p1).f4, globalState));
            _1444_v1 = _1444_v1;
            _1567_v95 = (_1567_v95).update(_dafny.Seq.IsProperPrefixOf(_1444_v1, _1444_v1), true);
          }
        }
      }
      let _1574_v100;
      let _nw239 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
      _1574_v100 = _nw239;
      _1574_v100 = _1574_v100;
      r0 = p0;
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _1575_v0;
      let _1576_v1;
      let _out28;
      let _out29;
      let _outcollector10 = (_this).m5(p1.f3, _this.f8, globalState);
      _out28 = _outcollector10[0];
      _out29 = _outcollector10[1];
      _1575_v0 = _out28;
      _1576_v1 = _out29;
      let _1577_v2;
      let _init47 = ((_1578_v0, _1579_p0, _1580_p1) => function (_1581_i0) {
        return _dafny.Seq.of(_1578_v0, new BigNumber((_dafny.Seq.of(_1579_p0, _1579_p0)).length), new BigNumber((_dafny.Seq.of((_1580_p1).f4, (_1580_p1).f4)).length), new BigNumber((_dafny.Set.fromElements((_this).f9, new BigNumber((_dafny.MultiSet.fromElements(_1578_v0, (_this).f9)).cardinality()))).length));
      })(_1575_v0, p0, p1);
      let _nw240 = Array((new BigNumber(19)).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw240.length); _i0_47++) {
        _nw240[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _1577_v2 = _nw240;
      let _1582_v3;
      let _init48 = ((_1583_v0) => function (_1584_i1) {
        return _dafny.Set.fromElements(_1583_v0, (_this).f9);
      })(_1575_v0);
      let _nw241 = Array((new BigNumber(24)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw241.length); _i0_48++) {
        _nw241[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1582_v3 = _nw241;
      let _1585_v4;
      _1585_v4 = _module.D15.create_DC32(_1582_v3);
      let _1586_v5;
      _1586_v5 = _module.D15.create_DC34(_1585_v4);
      let _1587_v6;
      _1587_v6 = _module.D15.create_DC34(_1585_v4);
      let _1588_v7;
      _1588_v7 = _module.D15.create_DC34(_1587_v6);
      let _1589_v8;
      _1589_v8 = _module.D15.create_DC34(_1588_v7);
      let _1590_v9;
      _1590_v9 = _module.D15.create_DC34(_1585_v4);
      let _1591_v10;
      _1591_v10 = _module.D15.create_DC34(_1589_v8);
      let _1592_v11;
      _1592_v11 = _module.D15.create_DC34(_1589_v8);
      let _1593_v12;
      _1593_v12 = _dafny.Set.fromElements(_1592_v11, _1592_v11);
      let _1594_v13;
      _1594_v13 = _dafny.Seq.of(new BigNumber((_1593_v12).length), (_this).f9, (_dafny.ZERO).minus(_1575_v0), new BigNumber(442), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,_this.f3)).length));
      let _1595_v14;
      _1595_v14 = _dafny.Seq.of((p1).f4);
      let _index253 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length));
      (_1577_v2)[_index253] = _dafny.Seq.Concat(_1594_v13, _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_1595_v14)).cardinality())));
      let _index254 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length));
      (_1577_v2)[_index254] = _1594_v13;
      let _1596_v15;
      let _nw242 = Array((_dafny.ONE).toNumber()).fill(false);
      _1596_v15 = _nw242;
      let _index255 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
      (_1596_v15)[_index255] = (_this).f4;
      let _1597_v16;
      _1597_v16 = _dafny.MultiSet.fromElements(!(_1576_v1), true);
      let _index256 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
      (_1596_v15)[_index256] = (_1597_v16).IsProperSubsetOf(_1597_v16);
      _1575_v0 = (_this).f9;
      if ((((_this).f4) ? (false) : (((true) ? (true) : ((p1).f4))))) {
        let _index257 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
        (_1596_v15)[_index257] = !(_module.__default.fm14(globalState));
        let _1598_v17;
        _1598_v17 = _module.D12.create_DC27((_this).f9);
        let _1599_v18;
        _1599_v18 = _module.D12.create_DC28(_1598_v17);
        let _rhs213 = (_this).f9;
        let _rhs214 = _1599_v18;
        _1575_v0 = _rhs213;
        _1599_v18 = _rhs214;
        _1575_v0 = _module.__default.safeDivisionInt((_this).f9, _1575_v0);
        let _1600_v19;
        _1600_v19 = _dafny.Seq.UnicodeFromString("abbgbhwi");
        let _1601_v20;
        _1601_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1575_v0,((_this).f9).isLessThanOrEqualTo(new BigNumber((_1600_v19).length)));
        _1601_v20 = (_1601_v20).update(_module.__default.safeModuloInt(new BigNumber(221), new BigNumber((_1597_v16).cardinality())), _this.f3);
        (_this).f8 = p1.f3;
      } else {
        let _1602_v21;
        _1602_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_this.f3);
        let _1603_v22;
        _1603_v22 = _module.D11.create_DC25(_1602_v21);
        _1603_v22 = ((_this.f3) ? (_1603_v22) : (_1603_v22));
        if ((((_1602_v21).contains(new BigNumber((_1597_v16).cardinality()))) ? ((_1602_v21).get(new BigNumber((_1597_v16).cardinality()))) : ((_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]))) {
          let _1604_v23;
          let _nw243 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _1604_v23 = _nw243;
          let _rhs215 = ((_1594_v13)[_module.__default.safeIndex(_1575_v0, new BigNumber((_1594_v13).length))]).multipliedBy(_1575_v0);
          let _rhs216 = _1604_v23;
          _1575_v0 = _rhs215;
          _1604_v23 = _rhs216;
          let _1605_v24;
          _1605_v24 = _module.D16.create_DC35(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1602_v21).length)));
          let _1606_v25;
          _1606_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1605_v24).dtor_cf50).length),(_this).f9);
          let _1607_v26;
          _1607_v26 = _dafny.Seq.UnicodeFromString("w");
          let _1608_v27;
          _1608_v27 = _dafny.Seq.of(_1607_v26);
          let _1609_v28;
          _1609_v28 = _dafny.Set.fromElements((_this).f9);
          let _1610_v29;
          _1610_v29 = _module.D4.create_DC12(new BigNumber((_1606_v25).length), new BigNumber(((_1608_v27)[_module.__default.safeIndex(new BigNumber((_1609_v28).length), new BigNumber((_1608_v27).length))]).length), (_1595_v14)[_module.__default.safeIndex((_dafny.ZERO).minus(_1575_v0), new BigNumber((_1595_v14).length))]);
          let _1611_v30;
          _1611_v30 = _module.D7.create_DC17(_1575_v0, (_this).f4, true, new BigNumber(316));
          let _1612_v31;
          _1612_v31 = _module.D8.create_DC20(p0, (p1).f4);
          let _1613_v32;
          _1613_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1612_v31);
          let _pat_let_tv44 = _1613_v32;
          let _pat_let_tv45 = _1613_v32;
          let _pat_let_tv46 = _1611_v30;
          (p1).f3 = (function (_pat_let36_0) {
            return function (_1614_dt__update__tmp_h0) {
              return function (_pat_let37_0) {
                return function (_1616_dt__update_hcf19_h0) {
                  return function (_pat_let38_0) {
                    return function (_1617_dt__update_hcf21_h0) {
                      return _module.D4.create_DC12(_1616_dt__update_hcf19_h0, (_1614_dt__update__tmp_h0).dtor_cf20, _1617_dt__update_hcf21_h0);
                    }(_pat_let38_0);
                  }((_pat_let_tv46).dtor_cf28);
                }(_pat_let37_0);
              }(new BigNumber((function () {
                let _coll51 = new _dafny.Set();
                for (const _compr_51 of (_pat_let_tv44).Keys.Elements) {
                  let _1615_v33 = _compr_51;
                  if ((_pat_let_tv45).contains(_1615_v33)) {
                    _coll51.add(_1615_v33);
                  }
                }
                return _coll51;
              }()).length));
            }(_pat_let36_0);
          }(_1610_v29)).dtor_cf21;
          let _1618_v34;
          _1618_v34 = _dafny.MultiSet.fromElements((_this).f9);
          let _1619_v35;
          let _nw244 = new _module.C2();
          _nw244.__ctor(_this.f8, ((_dafny.MultiSet.fromElements((_this).f9)).update(new BigNumber((_1607_v26).length), _module.__default.abs((_this).f9))).equals(_1618_v34), (_1595_v14)[_module.__default.safeIndex((_this).f9, new BigNumber((_1595_v14).length))], _1576_v1);
          _1619_v35 = _nw244;
          _1576_v1 = _module.__default.fm14(globalState);
          let _index258 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
          (_1596_v15)[_index258] = _this.f8;
        } else {
          let _1620_v36;
          let _nw245 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1620_v36 = _nw245;
          let _1621_v37;
          _1621_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1620_v36);
          _1621_v37 = (_1621_v37).update(p0, _1620_v36);
          let _1622_v38;
          let _nw246 = new _module.C0();
          _nw246.__ctor();
          _1622_v38 = _nw246;
          _1575_v0 = _module.__default.safeModuloInt(_1575_v0, _1575_v0);
          (_this).f8 = _module.__default.fm18((_1575_v0).multipliedBy((_this).f9), p0, (_this).f9, globalState);
          let _1623_v39;
          _1623_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC27(_1575_v0),(_dafny.ZERO).minus(_1575_v0));
          let _1624_v40;
          _1624_v40 = _dafny.Seq.UnicodeFromString("kcg");
          _1623_v39 = (_1623_v39).update(_module.D12.create_DC27(new BigNumber(948)), new BigNumber((_1624_v40).length));
        }
        let _source18 = _1603_v22;
        if (_source18.is_DC25) {
          let _1625___mcc_h0 = (_source18).cf38;
          let _1626_cf38 = _1625___mcc_h0;
          let _1627_v41;
          _1627_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-533),_module.__default.fm34((_this).f9, (_this).f9, globalState));
          let _1628_v42;
          _1628_v42 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          _1627_v41 = (_1627_v41).update((new BigNumber((_dafny.Set.fromElements(new BigNumber((_1628_v42).length))).length)).plus((_this).f9), _1595_v14);
          let _1629_v43;
          _1629_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          let _1630_v44;
          _1630_v44 = _dafny.MultiSet.fromElements(_1629_v43);
          let _1631_v45;
          _1631_v45 = _dafny.Seq.UnicodeFromString("jophvjl");
          let _1632_v46;
          _1632_v46 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm21(new BigNumber(402), _1630_v44, globalState)),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jhflclhlo"), _1631_v45));
          let _1633_v47;
          _1633_v47 = _module.D1.create_DC5((p1).f4, p1.f3, new BigNumber((_1594_v13).length), (_this).f9);
          _1632_v46 = (_1632_v46).update(((p1).f4) && (p1.f3), _module.__default.fm30(_1633_v47, _1575_v0, (_this).f9, globalState));
          let _1634_v48;
          _1634_v48 = _module.D1.create_DC4((_this).f9);
          let _1635_v49;
          _1635_v49 = _module.D3.create_DC10(_1597_v16, _1634_v48);
          let _1636_v50;
          _1636_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1635_v49,p0);
          _1636_v50 = (_1636_v50).update(_module.__default.fm39((_this).f4, _1575_v0, (_this).f9, _1575_v0, globalState), p0);
          let _1637_v51;
          let _nw247 = new _module.C3();
          _nw247.__ctor(_module.__default.fm14(globalState), (p1).f4);
          _1637_v51 = _nw247;
          let _1638_v52;
          _1638_v52 = _dafny.MultiSet.fromElements(_1637_v51);
          let _1639_v53;
          _1639_v53 = _dafny.MultiSet.fromElements((_this).f9, _1575_v0);
          let _rhs217 = ((((_1638_v52).contains(_1637_v51)) ? ((_1638_v52).get(_1637_v51)) : (new BigNumber(((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))]).length)))).minus(_module.__default.safeModuloInt(_1575_v0, (((_1639_v53).contains(new BigNumber(360))) ? ((_1639_v53).get(new BigNumber(360))) : ((_this).f9))));
          let _rhs218 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), ((_1640_p0) => function (_1641_i2) {
            return _1640_p0;
          })(p0));
          let _rhs219 = _module.__default.fm6(((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))])[_module.__default.safeIndex(new BigNumber(862), new BigNumber(((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))]).length))], globalState);
          let _lhs153 = p1;
          _1575_v0 = _rhs217;
          r0 = _rhs218;
          _lhs153.f3 = _rhs219;
        } else {
          let _1642___mcc_h1 = (_source18).cf37;
          let _1643_cf37 = _1642___mcc_h1;
          let _1644_v54;
          _1644_v54 = _dafny.Seq.UnicodeFromString("grqnosdc");
          (globalState).f0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1644_v54, _1644_v54), _1644_v54);
          let _1645_v55;
          let _nw248 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
          _1645_v55 = _nw248;
          let _1646_v56;
          _1646_v56 = _module.D0.create_DC2(_1576_v1, new BigNumber(-720), _1575_v0);
          _1645_v55 = (((_1646_v56).dtor_cf3) ? (_1645_v55) : (_1645_v55));
          let _1647_v57;
          let _init49 = ((_1648_p1) => function (_1649_i3) {
            return _module.D7.create_DC17((_this).f9, (_1648_p1).f4, _this.f3, (_this).f9);
          })(p1);
          let _nw249 = Array((new BigNumber(6)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw249.length); _i0_49++) {
            _nw249[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1647_v57 = _nw249;
          let _init50 = ((_1650_v13, _1651_p1, _1652_v1) => function (_1653_i4) {
            return _module.D7.create_DC17(new BigNumber((_1650_v13).length), (_1651_p1).f4, _1652_v1, (_this).f9);
          })(_1594_v13, p1, _1576_v1);
          let _nw250 = Array((new BigNumber(26)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw250.length); _i0_50++) {
            _nw250[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1647_v57 = _nw250;
          (globalState).f0 = (_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))];
        }
        let _1654_v58;
        _1654_v58 = _dafny.Set.fromElements((_this).f9);
        let _1655_v59;
        _1655_v59 = _module.D7.create_DC17((_this).f9, _module.__default.fm18(new BigNumber((_1654_v58).length), p0, _1575_v0, globalState), true, _1575_v0);
        let _1656_v60;
        _1656_v60 = _dafny.Seq.of(_1655_v59, _1655_v59);
        (_this).f8 = ((p1.f3) ? ((new BigNumber((_1595_v14).length)).isEqualTo(_1575_v0)) : (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1656_v60, _module.__default.safeIndex(_1575_v0, new BigNumber((_1656_v60).length)), _1655_v59), _1656_v60)));
        _1575_v0 = (_this).f9;
      }
      let _1657_v61;
      _1657_v61 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p1.f3)),_1575_v0);
      let _1658_v62;
      _1658_v62 = _module.D16.create_DC35((_1657_v61).Merge(_1657_v61));
      let _source19 = _1658_v62;
      if (_source19.is_DC36) {
        let _1659___mcc_h2 = (_source19).cf51;
        let _1660___mcc_h3 = (_source19).cf52;
        let _1661___mcc_h4 = (_source19).cf53;
        let _1662___mcc_h5 = (_source19).cf54;
        let _1663_cf54 = _1662___mcc_h5;
        let _1664_cf53 = _1661___mcc_h4;
        let _1665_cf52 = _1660___mcc_h3;
        let _1666_cf51 = _1659___mcc_h2;
        let _1667_v63;
        _1667_v63 = _dafny.Seq.of(_1596_v15);
        let _1668_v64;
        _1668_v64 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1575_v0));
        let _1669_v65;
        _1669_v65 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1575_v0),(p1).f4);
        let _1670_v66;
        _1670_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1575_v0,new BigNumber(-818));
        let _1671_v67;
        _1671_v67 = _dafny.Seq.UnicodeFromString("nnfqnbnws");
        let _1672_v69;
        _1672_v69 = _module.D1.create_DC5(_1576_v1, (p1).f4, (_this).f9, new BigNumber((_1671_v67).length));
        let _1673_v70;
        let _nw251 = Array((new BigNumber(28)).toNumber());
        _nw251[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw251[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1667_v63, _1667_v63), _module.__default.safeIndex(new BigNumber((_1668_v64).cardinality()), new BigNumber((_dafny.Seq.Concat(_1667_v63, _1667_v63)).length)), _1596_v15)).length);
        _nw251[(new BigNumber(2)).toNumber()] = new BigNumber(101);
        _nw251[(new BigNumber(3)).toNumber()] = (new BigNumber((_1669_v65).length)).plus(new BigNumber((_1670_v66).length));
        _nw251[(new BigNumber(4)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(5)).toNumber()] = (_1575_v0).plus(_1575_v0);
        _nw251[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_1575_v0).plus((_this).f9));
        _nw251[(new BigNumber(7)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(8)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw251[(new BigNumber(10)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_this).f9);
        _nw251[(new BigNumber(12)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1671_v67,(_this).f4)).length);
        _nw251[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1671_v67, _dafny.Seq.UnicodeFromString("swjrixy"))).length);
        _nw251[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_this).f9);
        _nw251[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt((((_1597_v16).contains(_this.f3)) ? ((_1597_v16).get(_this.f3)) : (_1575_v0)), (_this).f9);
        _nw251[(new BigNumber(17)).toNumber()] = (_this).f9;
        _nw251[(new BigNumber(18)).toNumber()] = (_this).f9;
        _nw251[(new BigNumber(19)).toNumber()] = new BigNumber((function () {
          let _coll52 = new _dafny.Set();
          for (const _compr_52 of _dafny.IntegerRange(new BigNumber(283), new BigNumber(923))) {
            let _1674_v68 = _compr_52;
            if (((new BigNumber(283)).isLessThanOrEqualTo(_1674_v68)) && ((_1674_v68).isLessThan(new BigNumber(923)))) {
              _coll52.add(_module.__default.safeDivisionInt(_1674_v68, new BigNumber((_1670_v66).length)));
            }
          }
          return _coll52;
        }()).length);
        _nw251[(new BigNumber(20)).toNumber()] = ((_this).f9).multipliedBy((_1672_v69).dtor_cf10);
        _nw251[(new BigNumber(21)).toNumber()] = ((_this).f9).minus((_this).f9);
        _nw251[(new BigNumber(22)).toNumber()] = (_this).f9;
        _nw251[(new BigNumber(23)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(24)).toNumber()] = (_this).f9;
        _nw251[(new BigNumber(25)).toNumber()] = _1575_v0;
        _nw251[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt((_1594_v13)[_module.__default.safeIndex(_1575_v0, new BigNumber((_1594_v13).length))], (_this).fm5((_this).f4, _1671_v67, globalState));
        _nw251[(new BigNumber(27)).toNumber()] = _1575_v0;
        _1673_v70 = _nw251;
        let _index259 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length));
        (_1673_v70)[_index259] = _1575_v0;
        let _1675_v71;
        _1675_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1664_cf53,(p1).f4);
        let _1676_v72;
        _1676_v72 = _dafny.Set.fromElements(_1675_v71, _1675_v71);
        let _1677_v73;
        _1677_v73 = _dafny.MultiSet.fromElements(_1675_v71);
        let _index260 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length));
        (_1673_v70)[_index260] = new BigNumber(((_1676_v72).Union((_1676_v72).Intersect(function () {
          let _coll53 = new _dafny.Set();
          for (const _compr_53 of ((_1677_v73).update(_1675_v71, _module.__default.abs(_1575_v0))).Elements) {
            let _1678_v74 = _compr_53;
            if (((_1677_v73).update(_1675_v71, _module.__default.abs(_1575_v0))).contains(_1678_v74)) {
              _coll53.add(_1678_v74);
            }
          }
          return _coll53;
        }()))).length);
        if (_1576_v1) {
          let _index261 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length));
          (_1673_v70)[_index261] = (_this).f9;
          let _1679_v75;
          let _nw252 = new _module.C1();
          _nw252.__ctor(_1671_v67, (new BigNumber(83)).isLessThan((_this).f9));
          _1679_v75 = _nw252;
          _1679_v75 = _1679_v75;
          _1575_v0 = ((_this).f9).plus(((_this).f9).minus(new BigNumber(265)));
          _1575_v0 = ((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))]).minus((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))]);
          let _1680_v76;
          _1680_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))],p0);
          let _1681_v77;
          _1681_v77 = _module.D17.create_DC37(_1680_v76);
          let _1682_v79;
          _1682_v79 = _dafny.Set.fromElements(_dafny.Seq.of(_1575_v0));
          let _1683_v80;
          _1683_v80 = _module.D18.create_DC41(_1682_v79);
          let _1684_v81;
          _1684_v81 = _dafny.Seq.of(_1680_v76, _1680_v76, _1680_v76);
          let _1685_v83;
          _1685_v83 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(964)), ((_1686_v0) => function (_1687_i6) {
            return _1686_v0;
          })(_1575_v0)));
          let _1688_v85;
          _1688_v85 = _dafny.Map.Empty.slice().updateUnsafe((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))],(p1).f4);
          let _1689_v87;
          let _nw253 = Array((new BigNumber(27)).toNumber());
          _nw253[(_dafny.ZERO).toNumber()] = (_1680_v76).update(_module.__default.fm19(globalState), p0);
          _nw253[(_dafny.ONE).toNumber()] = (_1681_v77).dtor_cf55;
          _nw253[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_1690_i5) {
            return new BigNumber(-261);
          }),p0);
          _nw253[(new BigNumber(3)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(4)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(5)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(6)).toNumber()] = function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of ((_1683_v80).dtor_cf63).Elements) {
              let _1691_v78 = _compr_54;
              if (((_1683_v80).dtor_cf63).contains(_1691_v78)) {
                _coll54.push([_1691_v78,p0]);
              }
            }
            return _coll54;
          }();
          _nw253[(new BigNumber(7)).toNumber()] = (_1684_v81)[_module.__default.safeIndex((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))], new BigNumber((_1684_v81).length))];
          _nw253[(new BigNumber(8)).toNumber()] = (_1680_v76).update((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))], p0);
          _nw253[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1594_v13,p0)).Merge(_1680_v76);
          _nw253[(new BigNumber(10)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))],p0);
          _nw253[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_1671_v67).length)),p0);
          _nw253[(new BigNumber(13)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(14)).toNumber()] = (_1680_v76).Merge(_1680_v76);
          _nw253[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1594_v13,p0);
          _nw253[(new BigNumber(16)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(17)).toNumber()] = function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of (_1685_v83).Elements) {
              let _1692_v82 = _compr_55;
              if ((_1685_v83).contains(_1692_v82)) {
                _coll55.push([_1692_v82,p0]);
              }
            }
            return _coll55;
          }();
          _nw253[(new BigNumber(18)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(19)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(20)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(21)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(22)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(23)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1594_v13,p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))],p0));
          _nw253[(new BigNumber(24)).toNumber()] = (function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of (_1688_v85).Keys.Elements) {
              let _1693_v84 = _compr_56;
              if ((_1688_v85).contains(_1693_v84)) {
                _coll56.push([_1693_v84,new _dafny.CodePoint('p'.codePointAt(0))]);
              }
            }
            return _coll56;
          }()).Merge(function () {
            let _coll57 = new _dafny.Map();
            for (const _compr_57 of (_1685_v83).Elements) {
              let _1694_v86 = _compr_57;
              if ((_1685_v83).contains(_1694_v86)) {
                _coll57.push([_1694_v86,p0]);
              }
            }
            return _coll57;
          }());
          _nw253[(new BigNumber(25)).toNumber()] = _1680_v76;
          _nw253[(new BigNumber(26)).toNumber()] = ((_1680_v76).update(_1594_v13, p0)).Merge(_1680_v76);
          _1689_v87 = _nw253;
          let _index262 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1689_v87).length));
          (_1689_v87)[_index262] = _1680_v76;
          let _index263 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1689_v87).length));
          (_1689_v87)[_index263] = (_1680_v76).Merge(_1680_v76);
        } else {
          let _1695_v89;
          _1695_v89 = _dafny.Seq.of(_1669_v65, function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_dafny.Seq.of((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))])).Elements) {
              let _1696_v88 = _compr_58;
              if (_dafny.Seq.contains(_dafny.Seq.of((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))]), _1696_v88)) {
                _coll58.push([_module.__default.safeDivisionInt(_1696_v88, (_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))]),!(true)]);
              }
            }
            return _coll58;
          }(), (((p1).f4) ? (_1669_v65) : (_1669_v65)), (_1669_v65).Merge(_1669_v65), _dafny.Map.Empty.slice().updateUnsafe((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))],_1663_cf54));
          let _rhs220 = _1695_v89;
          let _rhs221 = (function () {
            let _coll59 = new _dafny.Set();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(626), new BigNumber(615))) {
              let _1697_v90 = _compr_59;
              if (((new BigNumber(626)).isLessThanOrEqualTo(_1697_v90)) && ((_1697_v90).isLessThan(new BigNumber(615)))) {
                _coll59.add(_module.__default.safeDivisionInt(_1697_v90, new BigNumber((_dafny.Seq.of(_this.f3, !(_this.f3), _this.f3)).length)));
              }
            }
            return _coll59;
          }()).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))],_1663_cf54)).length)));
          let _rhs222 = ((((p1).f4) ? ((p1).f4) : (p1.f3))) || ((_1665_cf52) === (!(true)));
          let _rhs223 = (_this).f9;
          let _lhs154 = globalState;
          _1695_v89 = _rhs220;
          _lhs154.f0 = _rhs221;
          _1576_v1 = _rhs222;
          _1575_v0 = _rhs223;
          let _1698_v92;
          _1698_v92 = _module.D8.create_DC18(function () {
  let _coll60 = new _dafny.Map();
  for (const _compr_60 of (_1668_v64).Elements) {
    let _1699_v91 = _compr_60;
    if ((_1668_v64).contains(_1699_v91)) {
      _coll60.push([(_1699_v91).plus((_this).f9),_1663_cf54]);
    }
  }
  return _coll60;
}());
          let _pat_let_tv47 = _1669_v65;
          let _1700_v93;
          _1700_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1575_v0,function (_pat_let39_0) {
            return function (_1701_dt__update__tmp_h1) {
              return function (_pat_let40_0) {
                return function (_1702_dt__update_hcf30_h0) {
                  return _module.D8.create_DC18(_1702_dt__update_hcf30_h0);
                }(_pat_let40_0);
              }(_pat_let_tv47);
            }(_pat_let39_0);
          }(_1698_v92));
          let _1703_v94;
          _1703_v94 = _dafny.MultiSet.fromElements(_1700_v93);
          _1703_v94 = (((p1).f4) ? (_1703_v94) : (_1703_v94));
          let _1704_v95;
          _1704_v95 = _module.D4.create_DC12((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))], (_this).f9, _this.f8);
          (globalState).f0 = !((_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]) || ((_1704_v95).dtor_cf21);
          let _index264 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length));
          (_1673_v70)[_index264] = _1575_v0;
          let _index265 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length));
          (_1673_v70)[_index265] = (_dafny.ZERO).minus((_this).fm0(globalState));
        }
        (p1).f3 = !(_1576_v1);
        let _1705_v96;
        _1705_v96 = _module.D0.create_DC1(p1.f3, p0);
        let _1706_v97;
        _1706_v97 = _dafny.Seq.of(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-829)), function (_1707_i7) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), ((_1708_p0) => function (_1709_i8) {
          return _1708_p0;
        })(p0))));
        let _1710_v98;
        _1710_v98 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ykv"));
        let _1711_v99;
        _1711_v99 = _1710_v98;
        let _index266 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
        let _rhs224 = _1705_v96;
        let _rhs225 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_1706_v97, _module.__default.safeIndex(_1575_v0, new BigNumber((_1706_v97).length)), _1711_v99)).length), _1575_v0)).isLessThanOrEqualTo(((_this).f9).plus((_1673_v70)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1673_v70).length))]));
        let _lhs155 = _1596_v15;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
        _1705_v96 = _rhs224;
        _lhs155[_lhs156] = _rhs225;
      } else {
        let _1712___mcc_h6 = (_source19).cf50;
        let _1713_cf50 = _1712___mcc_h6;
        if (_this.f8) {
          let _1714_v100;
          let _1715_v101;
          let _out30;
          let _out31;
          let _outcollector11 = (_this).m5((p1).f4, p1.f3, globalState);
          _out30 = _outcollector11[0];
          _out31 = _outcollector11[1];
          _1714_v100 = _out30;
          _1715_v101 = _out31;
          let _1716_v102;
          _1716_v102 = _dafny.Set.fromElements(new BigNumber((_1713_cf50).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((p1).f4,p0)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-156)), ((_1717_p0) => function (_1718_i9) {
            return _1717_p0;
          })(p0))).length));
          let _1719_v103;
          let _init51 = ((_1720_v100) => function (_1721_i10) {
            return (_1721_i10).minus(_1720_v100);
          })(_1714_v100);
          let _nw254 = Array((new BigNumber(26)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw254.length); _i0_51++) {
            _nw254[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1719_v103 = _nw254;
          let _1722_v104;
          _1722_v104 = _dafny.Seq.of(_1719_v103);
          let _1723_v105;
          let _nw255 = Array((new BigNumber(21)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = (_1722_v104)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_1722_v104).length))];
          _nw255[(_dafny.ONE).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(2)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(3)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(4)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(5)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(6)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(7)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(8)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(9)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(10)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(11)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(12)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(13)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(14)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(15)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(16)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(17)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(18)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(19)).toNumber()] = _1719_v103;
          _nw255[(new BigNumber(20)).toNumber()] = _1719_v103;
          _1723_v105 = _nw255;
          let _1724_v106;
          _1724_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1723_v105,p1.f3);
          let _1725_v107;
          _1725_v107 = _module.D8.create_DC19(_1575_v0, (_this).f9);
          let _1726_v108;
          _1726_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          let _1727_v109;
          let _nw256 = new _module.C3();
          _nw256.__ctor((_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))], true);
          _1727_v109 = _nw256;
          let _1728_v110;
          _1728_v110 = _module.D0.create_DC2(p1.f3, (_this).f9, _1714_v100);
          let _1729_v111;
          let _nw257 = Array((new BigNumber(13)).toNumber());
          _nw257[(_dafny.ZERO).toNumber()] = _module.__default.fm31((_this).f4, new BigNumber(170), !(false), _dafny.Seq.of(_module.__default.fm18(new BigNumber((_1716_v102).length), new _dafny.CodePoint('e'.codePointAt(0)), _1575_v0, globalState), (((_1724_v106).contains(_1723_v105)) ? ((_1724_v106).get(_1723_v105)) : (_1576_v1))), globalState);
          _nw257[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt((_this).f9, _1575_v0);
          _nw257[(new BigNumber(2)).toNumber()] = _1714_v100;
          _nw257[(new BigNumber(3)).toNumber()] = _1714_v100;
          _nw257[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(((p1.f3) ? (_1575_v0) : ((_1725_v107).dtor_cf31)))));
          _nw257[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw257[(new BigNumber(6)).toNumber()] = (_this).f9;
          _nw257[(new BigNumber(7)).toNumber()] = (_this).f9;
          _nw257[(new BigNumber(8)).toNumber()] = new BigNumber(112);
          _nw257[(new BigNumber(9)).toNumber()] = new BigNumber(((_1726_v108).Merge(_1726_v108)).length);
          _nw257[(new BigNumber(10)).toNumber()] = (new BigNumber(-166)).multipliedBy(new BigNumber((_dafny.Seq.of(_1727_v109)).length));
          _nw257[(new BigNumber(11)).toNumber()] = (_1575_v0).minus(_1575_v0);
          _nw257[(new BigNumber(12)).toNumber()] = (_1728_v110).dtor_cf4;
          _1729_v111 = _nw257;
          _1729_v111 = _1729_v111;
          let _1730_v112;
          _1730_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1594_v13,_1575_v0);
          let _1731_v114;
          _1731_v114 = _dafny.Seq.of(_1594_v13);
          _1730_v112 = (_1730_v112).Merge(function () {
            let _coll61 = new _dafny.Map();
            for (const _compr_61 of (function () {
              let _coll62 = new _dafny.Set();
              for (const _compr_62 of (_1731_v114).Elements) {
                let _1732_v115 = _compr_62;
                if (_dafny.Seq.contains(_1731_v114, _1732_v115)) {
                  _coll62.add(_1732_v115);
                }
              }
              return _coll62;
            }()).Elements) {
              let _1733_v113 = _compr_61;
              if ((function () {
                let _coll63 = new _dafny.Set();
                for (const _compr_63 of (_1731_v114).Elements) {
                  let _1734_v115 = _compr_63;
                  if (_dafny.Seq.contains(_1731_v114, _1734_v115)) {
                    _coll63.add(_1734_v115);
                  }
                }
                return _coll63;
              }()).contains(_1733_v113)) {
                _coll61.push([_1733_v113,(_this).f9]);
              }
            }
            return _coll61;
          }());
          let _1735_v116;
          let _nw258 = new _module.C3();
          _nw258.__ctor(_1576_v1, !(_1575_v0).isEqualTo(_1714_v100));
          _1735_v116 = _nw258;
          _1735_v116 = _1735_v116;
          _1726_v108 = (_1726_v108).update(_1714_v100, new BigNumber(420));
        } else {
          let _1736_v117;
          let _nw259 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _1736_v117 = _nw259;
          let _index267 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length));
          (_1736_v117)[_index267] = _module.__default.safeModuloInt(_1575_v0, new BigNumber(-683));
          let _1737_v118;
          _1737_v118 = _module.D18.create_DC43(_1575_v0, (_this).f9, (p1).f4);
          let _index268 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length));
          (_1736_v117)[_index268] = (_dafny.ZERO).minus((_1737_v118).dtor_cf65);
          let _1738_v119;
          _1738_v119 = _dafny.Seq.UnicodeFromString("lhsbcmbd");
          let _1739_v120;
          _1739_v120 = _dafny.Seq.of(_1738_v119, _1738_v119);
          let _1740_v121;
          _1740_v121 = _dafny.Seq.of((_1739_v120)[_module.__default.safeIndex((_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))], new BigNumber((_1739_v120).length))], _dafny.Seq.Concat(_1738_v119, _1738_v119), _1738_v119);
          let _1741_v122;
          _1741_v122 = _module.D7.create_DC17((((_1713_cf50).contains((_1595_v14)[_module.__default.safeIndex((_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))], new BigNumber((_1595_v14).length))])) ? ((_1713_cf50).get((_1595_v14)[_module.__default.safeIndex((_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))], new BigNumber((_1595_v14).length))])) : ((_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))])), (_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))], _this.f3, new BigNumber(((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))]).length));
          r0 = (_1740_v121)[_module.__default.safeIndex(_module.__default.safeModuloInt((_1741_v122).dtor_cf26, _1575_v0), new BigNumber((_1740_v121).length))];
          let _1742_v123;
          _1742_v123 = _dafny.MultiSet.fromElements((_this).f9, (_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))], _1575_v0);
          let _1743_v124;
          let _nw260 = new _module.C2();
          _nw260.__ctor((_module.D8.create_DC20(p0, (p1).f4)).dtor_cf34, (_this).f4, (_1575_v0).isEqualTo(new BigNumber((_1742_v123).cardinality())), p1.f3);
          _1743_v124 = _nw260;
          r0 = _dafny.Seq.Concat(_1738_v119, _dafny.Seq.Create(_module.__default.abs(new BigNumber(92)), ((_1744_p0) => function (_1745_i11) {
            return _1744_p0;
          })(p0)));
          let _index269 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length));
          (_1736_v117)[_index269] = _module.__default.safeModuloInt(new BigNumber(915), (_1736_v117)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1736_v117).length))]);
        }
        let _1746_v125;
        _1746_v125 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,p0);
        let _1747_v126;
        _1747_v126 = _module.D14.create_DC31(_1575_v0, new BigNumber(((((_1746_v125).update((_this).f4, p0)).update(_this.f8, p0)).Merge(_module.__default.fm40(p0, _this.f3, _1575_v0, _1575_v0, globalState))).length));
        let _source20 = _1747_v126;
        if (_source20.is_DC31) {
          let _1748___mcc_h7 = (_source20).cf44;
          let _1749___mcc_h8 = (_source20).cf45;
          let _1750_cf45 = _1749___mcc_h8;
          let _1751_cf44 = _1748___mcc_h7;
          let _1752_v127;
          let _nw261 = new _module.C3();
          _nw261.__ctor(((_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]) === ((p1).f4), (p1).f4);
          _1752_v127 = _nw261;
          _1575_v0 = _1751_cf44;
          let _1753_v128;
          _1753_v128 = _dafny.Seq.UnicodeFromString("xxfvhd");
          let _1754_v129;
          _1754_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1753_v128,(_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]);
          let _1755_v130;
          _1755_v130 = _module.D1.create_DC5(p1.f3, (p1).f4, _1575_v0, new BigNumber((_dafny.Set.fromElements(_1750_cf45)).length));
          let _1756_v131;
          _1756_v131 = _dafny.Set.fromElements(_1575_v0, _1751_cf44, _1750_cf45);
          let _1757_v132;
          _1757_v132 = _dafny.Map.Empty.slice().updateUnsafe(_1750_cf45,(_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]);
          let _pat_let_tv48 = _1595_v14;
          let _pat_let_tv49 = p1;
          let _pat_let_tv50 = globalState;
          let _pat_let_tv51 = _1595_v14;
          let _pat_let_tv52 = _1753_v128;
          let _pat_let_tv53 = globalState;
          _1754_v129 = (_1754_v129).update(_module.__default.fm30(function (_pat_let41_0) {
            return function (_1758_dt__update__tmp_h2) {
              return function (_pat_let42_0) {
                return function (_1759_dt__update_hcf11_h0) {
                  return _module.D1.create_DC5((_1758_dt__update__tmp_h2).dtor_cf8, (_1758_dt__update__tmp_h2).dtor_cf9, (_1758_dt__update__tmp_h2).dtor_cf10, _1759_dt__update_hcf11_h0);
                }(_pat_let42_0);
              }((_dafny.ZERO).minus((_this).fm5((_pat_let_tv48)[_module.__default.safeIndex((_pat_let_tv49).fm0(_pat_let_tv50), new BigNumber((_pat_let_tv51).length))], _pat_let_tv52, _pat_let_tv53)));
            }(_pat_let41_0);
          }(_1755_v130), new BigNumber((_1756_v131).length), new BigNumber((_1757_v132).length), globalState), (_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]);
          let _1760_v133;
          let _init52 = ((_1761_p0) => function (_1762_i12) {
            return _1761_p0;
          })(p0);
          let _nw262 = Array((new BigNumber(12)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw262.length); _i0_52++) {
            _nw262[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _1760_v133 = _nw262;
          let _index270 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1760_v133).length));
          (_1760_v133)[_index270] = (_1753_v128)[_module.__default.safeIndex(new BigNumber(((_1577_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1577_v2).length))]).length), new BigNumber((_1753_v128).length))];
          let _1763_v134;
          _1763_v134 = new _dafny.CodePoint('x'.codePointAt(0));
          let _index271 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
          let _index272 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1760_v133).length));
          let _rhs226 = !(!(p1.f3));
          let _rhs227 = (_module.D0.create_DC1((p1).f4, p0)).dtor_cf2;
          let _rhs228 = _module.__default.fm8((_dafny.ZERO).minus(_1575_v0), globalState);
          let _lhs157 = _1596_v15;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
          let _lhs159 = _1760_v133;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1760_v133).length));
          _lhs157[_lhs158] = _rhs226;
          _lhs159[_lhs160] = _rhs227;
          _1763_v134 = _rhs228;
        } else {
          let _1764___mcc_h9 = (_source20).cf43;
          let _1765_cf43 = _1764___mcc_h9;
          _1575_v0 = _1575_v0;
          let _1766_v135;
          _1766_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-101),_1597_v16);
          _1597_v16 = ((((_1766_v135).contains((_this).fm0(globalState))) ? ((_1766_v135).get((_this).fm0(globalState))) : (_1597_v16))).update(p1.f3, _module.__default.abs((_this).f9));
          let _index273 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length));
          (_1596_v15)[_index273] = _this.f8;
          let _1767_v136;
          let _nw263 = new _module.C0();
          _nw263.__ctor();
          _1767_v136 = _nw263;
        }
        let _1768_v137;
        _1768_v137 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_1576_v1);
        _1768_v137 = (_1768_v137).update(_1576_v1, (_1596_v15)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1596_v15).length))]);
        _1575_v0 = _1575_v0;
      }
      r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-423)), function (_1769_i13) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      });
      return r0;
    }
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      r1 = p2;
      let _1770_v0;
      let _nw264 = Array((new BigNumber(8)).toNumber());
      _nw264[(_dafny.ZERO).toNumber()] = _module.__default.fm6(p2, globalState);
      _nw264[(_dafny.ONE).toNumber()] = !((_this).f4);
      _nw264[(new BigNumber(2)).toNumber()] = _this.f3;
      _nw264[(new BigNumber(3)).toNumber()] = p0;
      _nw264[(new BigNumber(4)).toNumber()] = (_this.f3) === (_this.f3);
      _nw264[(new BigNumber(5)).toNumber()] = p0;
      _nw264[(new BigNumber(6)).toNumber()] = (_this).f4;
      _nw264[(new BigNumber(7)).toNumber()] = (_this).f4;
      _1770_v0 = _nw264;
      _1770_v0 = _1770_v0;
      let _1771_v1;
      _1771_v1 = _dafny.Set.fromElements(_this.f3, p0, _module.__default.fm6(p2, globalState));
      let _1772_v2;
      _1772_v2 = _dafny.MultiSet.fromElements((_1771_v1).Intersect(_1771_v1), _1771_v1);
      let _index274 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
      (_1770_v0)[_index274] = p0;
      let _1773_v3;
      _1773_v3 = _dafny.Seq.of(_1771_v1);
      let _1774_v4;
      _1774_v4 = _module.D12.create_DC26(_1771_v1);
      let _1775_v5;
      _1775_v5 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1776_v6;
      _1776_v6 = _dafny.Seq.of(_1775_v5, _1775_v5);
      let _1777_v7;
      _1777_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1776_v6).length),_1771_v1);
      let _1778_v8;
      _1778_v8 = _module.D7.create_DC16(_1777_v7);
      let _pat_let_tv54 = _1775_v5;
      let _index275 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
      let _rhs229 = ((_1772_v2).Union(_1772_v2)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1773_v3, _dafny.Seq.of((_1774_v4).dtor_cf39))));
      let _rhs230 = function (_source21) {
        if (_source21.is_DC17) {
          let _1779___mcc_h0 = (_source21).cf26;
          let _1780___mcc_h1 = (_source21).cf27;
          let _1781___mcc_h2 = (_source21).cf28;
          let _1782___mcc_h3 = (_source21).cf29;
          let _1783_cf29 = _1782___mcc_h3;
          let _1784_cf28 = _1781___mcc_h2;
          let _1785_cf27 = _1780___mcc_h1;
          let _1786_cf26 = _1779___mcc_h0;
          return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f8,new BigNumber(521)),_this.f8)).contains(_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv54,_1784_cf28)).length))));
        } else {
          let _1787___mcc_h4 = (_source21).cf25;
          let _1788_cf25 = _1787___mcc_h4;
          return (_this).f4;
        }
      }(_1778_v8);
      let _lhs161 = _1770_v0;
      let _lhs162 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
      _1772_v2 = _rhs229;
      _lhs161[_lhs162] = _rhs230;
      let _index276 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
      (_1770_v0)[_index276] = _this.f3;
      let _1789_v9;
      _1789_v9 = _module.D7.create_DC17((_this).f9, _this.f3, _this.f8, (_this).f9);
      let _1790_v10;
      _1790_v10 = _dafny.MultiSet.fromElements(_1789_v9, _1789_v9, _1789_v9, _module.D7.create_DC17(new BigNumber((_dafny.Seq.UnicodeFromString("bydy")).length), false, _this.f8, p2), _1789_v9);
      let _1791_v11;
      let _nw265 = new _module.C2();
      _nw265.__ctor(p0, !(_this.f8) || ((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))]), ((_this.f3) ? (!(_module.__default.fm14(globalState))) : ((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))])), ((_1790_v10).update(_1789_v9, _module.__default.abs((_this).f9))).contains(_1789_v9));
      _1791_v11 = _nw265;
      if (!(_this.f3)) {
        let _1792_v12;
        _1792_v12 = _module.D10.create_DC23();
        _1792_v12 = _1792_v12;
        let _1793_v13;
        let _nw266 = new _module.C1();
        _nw266.__ctor(_1776_v6, _this.f3);
        _1793_v13 = _nw266;
        let _1794_v14;
        _1794_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1771_v1,(p2).multipliedBy((_this).f9));
        _1794_v14 = (_1794_v14).update(_dafny.Set.fromElements(_module.__default.fm18((_this).f9, _1775_v5, p2, globalState), (_1791_v11).f11, p0, (_1791_v11).f10, _this.f8), p2);
        let _1795_v15;
        _1795_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1793_v13).f13,new BigNumber(-268));
        let _1796_v16;
        _1796_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1770_v0,(((_1795_v15).contains(p0)) ? ((_1795_v15).get(p0)) : ((_this).f9)));
        let _1797_v17;
        _1797_v17 = _dafny.Seq.of(_1776_v6, _dafny.Seq.UnicodeFromString("gwvhmvke"));
        let _1798_v18;
        _1798_v18 = _dafny.Seq.of(new BigNumber(((_1797_v17)[_module.__default.safeIndex((_this).f9, new BigNumber((_1797_v17).length))]).length), p2);
        let _1799_v19;
        _1799_v19 = _dafny.Seq.of(_dafny.Seq.of(p2), _1798_v18, _1798_v18, _1798_v18);
        let _1800_v20;
        _1800_v20 = _dafny.MultiSet.fromElements(p2, new BigNumber(((_1799_v19)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f9)), new BigNumber((_1799_v19).length))]).length));
        _1796_v16 = (_dafny.Map.Empty.slice().updateUnsafe(_1770_v0,(_this).f9)).Merge((((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))]) ? (_1796_v16) : (_dafny.Map.Empty.slice().updateUnsafe(_1770_v0,new BigNumber((_1800_v20).cardinality())))));
        r1 = new BigNumber((_1776_v6).length);
      } else {
        let _1801_v21;
        _1801_v21 = _dafny.Set.fromElements((_this).f9, new BigNumber((_1776_v6).length), new BigNumber(702), p2, (_this).f9);
        let _1802_v23;
        let _nw267 = new _module.C3();
        _nw267.__ctor((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))], (function () {
          let _coll64 = new _dafny.Set();
          for (const _compr_64 of _dafny.IntegerRange(new BigNumber(157), new BigNumber(522))) {
            let _1803_v22 = _compr_64;
            if (((new BigNumber(157)).isLessThanOrEqualTo(_1803_v22)) && ((_1803_v22).isLessThan(new BigNumber(522)))) {
              _coll64.add((_1803_v22).minus(p2));
            }
          }
          return _coll64;
        }()).IsSubsetOf(_1801_v21));
        _1802_v23 = _nw267;
        _1802_v23 = _1802_v23;
        let _1804_v24;
        let _init53 = function (_1805_i0) {
          return (_1805_i0).multipliedBy((_this).f9);
        };
        let _nw268 = Array((new BigNumber(29)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw268.length); _i0_53++) {
          _nw268[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1804_v24 = _nw268;
        let _1806_v25;
        _1806_v25 = _dafny.Seq.of(_1804_v24, _1804_v24);
        let _1807_v26;
        _1807_v26 = _1806_v25;
        let _source22 = _1807_v26;
        let _1808___mcc_h5 = _source22;
        let _1809_cf23 = _1808___mcc_h5;
        (_1791_v11).m9(globalState);
        let _1810_v27;
        _1810_v27 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1791_v11).f11);
        let _1811_v28;
        _1811_v28 = _dafny.MultiSet.fromElements(_1810_v27);
        let _index277 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1804_v24).length));
        (_1804_v24)[_index277] = p2;
        let _index278 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1804_v24).length));
        let _rhs231 = _1811_v28;
        let _rhs232 = ((_this).f9).minus((_this).f9);
        let _lhs163 = _1804_v24;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1804_v24).length));
        _1811_v28 = _rhs231;
        _lhs163[_lhs164] = _rhs232;
        let _index279 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1804_v24).length));
        (_1804_v24)[_index279] = (_1791_v11).fm16(p2, true, p2, p1, globalState);
        let _1812_v29;
        _1812_v29 = _dafny.Seq.of((_1791_v11).f10);
        let _1813_v30;
        _1813_v30 = _dafny.MultiSet.fromElements((_1812_v29)[_module.__default.safeIndex((_this).f9, new BigNumber((_1812_v29).length))], _this.f8);
        let _1814_v31;
        _1814_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1812_v29,_1813_v30);
        _1814_v31 = ((_dafny.Map.Empty.slice().updateUnsafe(_1812_v29,_dafny.MultiSet.FromArray(_1812_v29))).Merge(_1814_v31)).Merge(_1814_v31);
        (globalState).f0 = (_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))];
        let _1815_v32;
        _1815_v32 = _dafny.Set.fromElements(_1770_v0);
        if ((_dafny.Set.fromElements(_1770_v0)).IsProperSubsetOf(_1815_v32)) {
          let _index280 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _rhs233 = (_1791_v11).f10;
          let _rhs234 = new BigNumber(220);
          let _rhs235 = _1775_v5;
          let _lhs165 = _1770_v0;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          _lhs165[_lhs166] = _rhs233;
          r1 = _rhs234;
          _1775_v5 = _rhs235;
          let _1816_v34;
          _1816_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(261),p2);
          let _1817_v35;
          _1817_v35 = _dafny.MultiSet.fromElements(function () {
            let _coll65 = new _dafny.Map();
            for (const _compr_65 of (_1816_v34).Keys.Elements) {
              let _1818_v33 = _compr_65;
              if ((_1816_v34).contains(_1818_v33)) {
                _coll65.push([_module.__default.safeModuloInt(_1818_v33, new BigNumber((_1771_v1).length)),p2]);
              }
            }
            return _coll65;
          }(), _1816_v34, _1816_v34);
          let _1819_v36;
          _1819_v36 = _dafny.MultiSet.fromElements(!((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))]), (_1791_v11).f11);
          let _pat_let_tv55 = p2;
          r1 = ((function (_pat_let43_0) {
            return function (_1820_dt__update__tmp_h0) {
              return function (_pat_let44_0) {
                return function (_1821_dt__update_hcf61_h0) {
                  return function (_pat_let45_0) {
                    return function (_1822_dt__update_hcf58_h0) {
                      return _module.D17.create_DC39((_1820_dt__update__tmp_h0).dtor_cf57, _1822_dt__update_hcf58_h0, (_1820_dt__update__tmp_h0).dtor_cf59, (_1820_dt__update__tmp_h0).dtor_cf60, _1821_dt__update_hcf61_h0);
                    }(_pat_let45_0);
                  }((_this).f4);
                }(_pat_let44_0);
              }(_pat_let_tv55);
            }(_pat_let43_0);
          }(_module.D17.create_DC39(_1804_v24, p0, (_1791_v11).f10, new BigNumber(((_1802_v23).fm7((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))], _1777_v7, _module.__default.fm21((_this).f9, _1817_v35, globalState), _1819_v36, globalState)).length), p2))).dtor_cf60).minus(_module.__default.safeModuloInt((_this).f9, (_this).f9));
          let _index281 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1804_v24).length));
          (_1804_v24)[_index281] = (_this).f9;
          let _index282 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1804_v24).length));
          let _index283 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _index284 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _rhs236 = _module.__default.safeDivisionInt(p2, (p2).plus(p2));
          let _rhs237 = (_this).f4;
          let _rhs238 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p2, p2))).plus(p2);
          let _rhs239 = (_1791_v11).f10;
          let _lhs167 = _1804_v24;
          let _lhs168 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1804_v24).length));
          let _lhs169 = _1770_v0;
          let _lhs170 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _lhs171 = _1770_v0;
          let _lhs172 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          _lhs167[_lhs168] = _rhs236;
          _lhs169[_lhs170] = _rhs237;
          r1 = _rhs238;
          _lhs171[_lhs172] = _rhs239;
          let _1823_v37;
          let _nw269 = new _module.C3();
          _nw269.__ctor((p2).isLessThanOrEqualTo(p2), (_1791_v11).f10);
          _1823_v37 = _nw269;
          let _1824_v38;
          let _nw270 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _1824_v38 = _nw270;
          let _1825_v39;
          _1825_v39 = _dafny.Seq.of((_1791_v11).f10, true);
          let _1826_v40;
          _1826_v40 = _module.D10.create_DC22(_1824_v38);
          let _1827_v41;
          let _nw271 = Array((new BigNumber(21)).toNumber());
          _nw271[(_dafny.ZERO).toNumber()] = _1824_v38;
          _nw271[(_dafny.ONE).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(2)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(3)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(4)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(5)).toNumber()] = (((_1825_v39)[_module.__default.safeIndex(new BigNumber((_1816_v34).length), new BigNumber((_1825_v39).length))]) ? (_1824_v38) : (_1824_v38));
          _nw271[(new BigNumber(6)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(7)).toNumber()] = (_1806_v25)[_module.__default.safeIndex((_1804_v24)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1804_v24).length))], new BigNumber((_1806_v25).length))];
          _nw271[(new BigNumber(8)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(9)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(10)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(11)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(12)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(13)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(14)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(15)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(16)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(17)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(18)).toNumber()] = (_1826_v40).dtor_cf36;
          _nw271[(new BigNumber(19)).toNumber()] = _1824_v38;
          _nw271[(new BigNumber(20)).toNumber()] = _1824_v38;
          _1827_v41 = _nw271;
          let _1828_v42;
          _1828_v42 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1804_v24)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_1804_v24).length))],_1775_v5)).length), p2, new BigNumber((_1825_v39).length));
          let _index285 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1827_v41).length));
          (_1827_v41)[_index285] = (_1806_v25)[_module.__default.safeIndex(new BigNumber((_1828_v42).length), new BigNumber((_1806_v25).length))];
          let _index286 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_1827_v41).length));
          (_1827_v41)[_index286] = _1824_v38;
        } else {
          let _1829_v43;
          _1829_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_1775_v5);
          let _1830_v44;
          _1830_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm5((_1791_v11).f10, _1776_v6, globalState));
          let _1831_v45;
          _1831_v45 = _dafny.MultiSet.fromElements(new BigNumber((_1830_v44).length));
          let _1832_v46;
          _1832_v46 = _dafny.MultiSet.fromElements(new BigNumber((_1829_v43).length), new BigNumber((_1831_v45).cardinality()));
          let _index287 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _rhs240 = _this.f8;
          let _rhs241 = (_1831_v45).IsSubsetOf((_1832_v46).Difference(_1832_v46));
          let _rhs242 = _1776_v6;
          let _rhs243 = ((!(_module.__default.fm18((_this).f9, _1775_v5, p2, globalState)) || ((_1791_v11).f11)) ? ((p2).isLessThanOrEqualTo(new BigNumber(295))) : (p0));
          let _lhs173 = _1770_v0;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length));
          let _lhs175 = globalState;
          let _lhs176 = _this;
          _lhs173[_lhs174] = _rhs240;
          _lhs175.f0 = _rhs241;
          _1776_v6 = _rhs242;
          _lhs176.f3 = _rhs243;
          let _1833_v47;
          _1833_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1791_v11).f10,(_this).f4);
          let _1834_v48;
          let _nw272 = new _module.C2();
          _nw272.__ctor((_1791_v11).f10, !_dafny.areEqual(_1776_v6, _dafny.Seq.UnicodeFromString("ffkltwrcq")), (_1791_v11).f11, (_module.__default.fm18(new BigNumber((_1833_v47).length), _1775_v5, new BigNumber(((_1791_v11).fm17(globalState)).length), globalState)) || ((_1770_v0)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_1770_v0).length))]));
          _1834_v48 = _nw272;
          let _1835_v49;
          let _nw273 = new _module.C3();
          _nw273.__ctor((_1791_v11).f11, !_dafny.Seq.contains((_1791_v11).fm17(globalState), _1775_v5));
          _1835_v49 = _nw273;
          let _1836_v50;
          _1836_v50 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          _1836_v50 = (_1836_v50).update(new BigNumber((_dafny.Seq.Concat(_1776_v6, _1776_v6)).length), (_1791_v11).f11);
          let _1837_v51;
          _1837_v51 = _module.D0.create_DC1((_1791_v11).f11, _1775_v5);
          let _1838_v52;
          _1838_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),_1775_v5);
          let _1839_v53;
          _1839_v53 = _dafny.Seq.of(_this.f3, true, (_1791_v11).f10, (_1791_v11).f11);
          let _1840_v54;
          let _nw274 = Array((new BigNumber(25)).toNumber());
          _nw274[(_dafny.ZERO).toNumber()] = _1775_v5;
          _nw274[(_dafny.ONE).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(2)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(3)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(4)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(5)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(6)).toNumber()] = (_1837_v51).dtor_cf2;
          _nw274[(new BigNumber(7)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
          _nw274[(new BigNumber(9)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(10)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(11)).toNumber()] = _module.__default.fm8(new BigNumber(60), globalState);
          _nw274[(new BigNumber(12)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(13)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(14)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(15)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(16)).toNumber()] = (((_1838_v52).contains(new BigNumber((_1839_v53).length))) ? ((_1838_v52).get(new BigNumber((_1839_v53).length))) : (new _dafny.CodePoint('w'.codePointAt(0))));
          _nw274[(new BigNumber(17)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(18)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(19)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(20)).toNumber()] = _module.__default.fm8(new BigNumber((_dafny.Set.fromElements(p2)).length), globalState);
          _nw274[(new BigNumber(21)).toNumber()] = _1775_v5;
          _nw274[(new BigNumber(22)).toNumber()] = (_1776_v6)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_1776_v6).length))];
          _nw274[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _nw274[(new BigNumber(24)).toNumber()] = _module.__default.fm23((_1791_v11).f11, new BigNumber((_1776_v6).length), globalState);
          _1840_v54 = _nw274;
          let _index288 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1840_v54).length));
          (_1840_v54)[_index288] = _1775_v5;
          let _index289 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1840_v54).length));
          (_1840_v54)[_index289] = new _dafny.CodePoint('v'.codePointAt(0));
        }
        let _1841_v55;
        _1841_v55 = _dafny.MultiSet.fromElements(_1776_v6);
        r1 = ((((_1841_v55).contains(_1776_v6)) ? ((_1841_v55).get(_1776_v6)) : ((_this).f9))).multipliedBy(_module.__default.safeDivisionInt(p2, p2));
      }
      let _1842_v56;
      let _init54 = ((_1843_p2) => function (_1844_i1) {
        return (_1844_i1).minus(_1843_p2);
      })(p2);
      let _nw275 = Array((new BigNumber(12)).toNumber());
      for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw275.length); _i0_54++) {
        _nw275[_i0_54] = _init54(new BigNumber(_i0_54));
      }
      _1842_v56 = _nw275;
      r0 = _1842_v56;
      let _1845_v57;
      let _nw276 = Array((new BigNumber(24)).toNumber());
      _nw276[(_dafny.ZERO).toNumber()] = _1791_v11;
      _nw276[(_dafny.ONE).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(2)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(3)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(4)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(5)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(6)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(7)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(8)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(9)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(10)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(11)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(12)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(13)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(14)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(15)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(16)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(17)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(18)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(19)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(20)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(21)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(22)).toNumber()] = _1791_v11;
      _nw276[(new BigNumber(23)).toNumber()] = _1791_v11;
      _1845_v57 = _nw276;
      let _1846_v58;
      _1846_v58 = _dafny.Seq.of(_1845_v57, _1845_v57);
      let _1847_v59;
      _1847_v59 = _dafny.Set.fromElements(p2, (_this).f9);
      let _1848_v60;
      _1848_v60 = _dafny.Seq.of((_1846_v58)[_module.__default.safeIndex(new BigNumber((_1847_v59).length), new BigNumber((_1846_v58).length))], _1845_v57, _1845_v57);
      let _1849_v61;
      _1849_v61 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f9);
      r1 = (new BigNumber((_1848_v60).length)).plus(new BigNumber(((_1849_v61).Merge(_1849_v61)).length));
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1850_v0;
      _1850_v0 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1851_v1;
      _1851_v1 = _module.D8.create_DC20(_1850_v0, p1);
      let _1852_v2;
      _1852_v2 = _dafny.Seq.UnicodeFromString("kxqsrb");
      let _1853_v4;
      let _nw277 = new _module.C0();
      _nw277.__ctor();
      _1853_v4 = _nw277;
      let _1854_v5;
      let _nw278 = Array((new BigNumber(19)).toNumber());
      _nw278[(_dafny.ZERO).toNumber()] = _1853_v4;
      _nw278[(_dafny.ONE).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(2)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(3)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(4)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(5)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(6)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(7)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(8)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(9)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(10)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(11)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(12)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(13)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(14)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(15)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(16)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(17)).toNumber()] = _1853_v4;
      _nw278[(new BigNumber(18)).toNumber()] = _1853_v4;
      _1854_v5 = _nw278;
      let _1855_v6;
      _1855_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1854_v5);
      let _1856_v7;
      _1856_v7 = _module.D19.create_DC44(_1855_v6);
      let _1857_v8;
      let _nw279 = Array((new BigNumber(10)).toNumber());
      _nw279[(_dafny.ZERO).toNumber()] = (_1851_v1).dtor_cf34;
      _nw279[(_dafny.ONE).toNumber()] = _this.f8;
      _nw279[(new BigNumber(2)).toNumber()] = p0;
      _nw279[(new BigNumber(3)).toNumber()] = (_module.D7.create_DC17(new BigNumber((_1852_v2).length), _this.f8, (_this).f4, (_this).f9)).dtor_cf28;
      _nw279[(new BigNumber(4)).toNumber()] = _this.f8;
      _nw279[(new BigNumber(5)).toNumber()] = _this.f8;
      _nw279[(new BigNumber(6)).toNumber()] = ((_this).f9).isLessThan(new BigNumber((function () {
        let _coll66 = new _dafny.Map();
        for (const _compr_66 of _dafny.IntegerRange(new BigNumber(128), new BigNumber(194))) {
          let _1858_v3 = _compr_66;
          if (((new BigNumber(128)).isLessThanOrEqualTo(_1858_v3)) && ((_1858_v3).isLessThan(new BigNumber(194)))) {
            _coll66.push([_module.__default.safeDivisionInt(_1858_v3, (_this).f9),(_this).f9]);
          }
        }
        return _coll66;
      }()).length));
      _nw279[(new BigNumber(7)).toNumber()] = !((_1856_v7).dtor_cf68).contains((_this).f4);
      _nw279[(new BigNumber(8)).toNumber()] = (((_module.D0.create_DC1(true, _1850_v0)).dtor_cf1) ? (_this.f3) : (_this.f8));
      _nw279[(new BigNumber(9)).toNumber()] = (_this).f4;
      _1857_v8 = _nw279;
      let _index290 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
      (_1857_v8)[_index290] = _module.__default.fm14(globalState);
      let _index291 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
      (_1857_v8)[_index291] = _this.f3;
      if (_this.f3) {
        let _1859_v9;
        _1859_v9 = _dafny.Set.fromElements(new BigNumber(-848));
        let _1860_v10;
        _1860_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jvjswysav"),new BigNumber((_1859_v9).length));
        r0 = (((_1860_v10).contains(_1852_v2)) ? ((_1860_v10).get(_1852_v2)) : ((_dafny.ZERO).minus((_this).f9)));
        let _1861_v11;
        let _nw280 = new _module.C3();
        _nw280.__ctor((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))], (_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]);
        _1861_v11 = _nw280;
        (globalState).f0 = p1;
        let _1862_v12;
        _1862_v12 = _dafny.Seq.of((_this).f9);
        let _1863_v14;
        _1863_v14 = _module.D8.create_DC18(function () {
  let _coll67 = new _dafny.Map();
  for (const _compr_67 of (_1862_v12).Elements) {
    let _1864_v13 = _compr_67;
    if (_dafny.Seq.contains(_1862_v12, _1864_v13)) {
      _coll67.push([(_1864_v13).plus((_this).f9),false]);
    }
  }
  return _coll67;
}());
        let _index292 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
        let _rhs244 = !(_dafny.Seq.IsProperPrefixOf(_1862_v12, _1862_v12)) || (false);
        let _rhs245 = (_dafny.ZERO).minus((((_this).f9).multipliedBy(new BigNumber(((_1863_v14).dtor_cf30).length))).multipliedBy((_this).f9));
        let _lhs177 = _1857_v8;
        let _lhs178 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
        _lhs177[_lhs178] = _rhs244;
        r0 = _rhs245;
        let _1865_v15;
        let _nw281 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1865_v15 = _nw281;
        let _1866_v16;
        _1866_v16 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_1867_i0) {
          return (_this).f9;
        })).length));
        let _index293 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1865_v15).length));
        (_1865_v15)[_index293] = _1866_v16;
        let _1868_v17;
        let _init55 = function (_1869_i1) {
          return _module.__default.safeDivisionInt(_1869_i1, (_this).f9);
        };
        let _nw282 = Array((new BigNumber(9)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw282.length); _i0_55++) {
          _nw282[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _1868_v17 = _nw282;
        let _1870_v18;
        _1870_v18 = _dafny.Seq.of(!(false));
        let _1871_v19;
        _1871_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1870_v18).length),p0);
        let _1872_v20;
        _1872_v20 = _module.D11.create_DC25(_1871_v19);
        let _1873_v21;
        _1873_v21 = _dafny.Seq.of(_1872_v20);
        let _1874_v22;
        _1874_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1873_v21,(_this).f4);
        let _index294 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1868_v17).length));
        (_1868_v17)[_index294] = (((((_1874_v22).contains(_1873_v21)) ? ((_1874_v22).get(_1873_v21)) : (p0))) ? (new BigNumber(917)) : ((_this).f9));
        let _index295 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1865_v15).length));
        let _index296 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1868_v17).length));
        let _rhs246 = _1866_v16;
        let _rhs247 = (_this).f9;
        let _rhs248 = (_this).f9;
        let _lhs179 = _1865_v15;
        let _lhs180 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_1865_v15).length));
        let _lhs181 = _1868_v17;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1868_v17).length));
        _lhs179[_lhs180] = _rhs246;
        _lhs181[_lhs182] = _rhs247;
        r0 = _rhs248;
      } else {
        r0 = (_this).f9;
        r0 = ((_this).f9).plus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-505), (_this).f9)));
        let _1875_v23;
        let _nw283 = new _module.C3();
        _nw283.__ctor(p0, (_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]);
        _1875_v23 = _nw283;
        let _1876_v24;
        _1876_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_this.f8);
        let _1877_v25;
        _1877_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1876_v24).length),!(!((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))])));
        let _1878_v26;
        _1878_v26 = _module.D8.create_DC18(_1876_v24);
        _1878_v26 = _1878_v26;
        let _1879_v27;
        _1879_v27 = _dafny.Seq.of(p0, p1);
        let _1880_v28;
        _1880_v28 = _dafny.Seq.of(_1879_v27, _1879_v27, _1879_v27, _dafny.Seq.Concat(_module.__default.fm34(new BigNumber(654), (_this).f9, globalState), _1879_v27));
        _1880_v28 = _dafny.Seq.of(_1879_v27, _dafny.Seq.Concat(_module.__default.fm34((_this).f9, (_this).f9, globalState), _dafny.Seq.of(!(false), (_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))], p1)));
      }
      let _1881_v29;
      let _init56 = function (_1882_i3) {
        return (_1882_i3).multipliedBy((_this).f9);
      };
      let _nw284 = Array((new BigNumber(2)).toNumber());
      for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw284.length); _i0_56++) {
        _nw284[_i0_56] = _init56(new BigNumber(_i0_56));
      }
      _1881_v29 = _nw284;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1881_v29).length))) {
        let _1883_i2 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1883_i2)) && ((_1883_i2).isLessThan(new BigNumber((_1881_v29).length))))) {
          (_1881_v29)[(_1883_i2)] = (_1883_i2).minus(((p0) ? ((_this).f9) : (new BigNumber(4))));
        }
      }
      _1852_v2 = _1852_v2;
      if (_this.f3) {
        let _1884_v30;
        _1884_v30 = _module.D1.create_DC3(_1857_v8);
        let _pat_let_tv56 = _1857_v8;
        let _source23 = function (_pat_let46_0) {
          return function (_1885_dt__update__tmp_h0) {
            return function (_pat_let47_0) {
              return function (_1886_dt__update_hcf6_h0) {
                return _module.D1.create_DC3(_1886_dt__update_hcf6_h0);
              }(_pat_let47_0);
            }(_pat_let_tv56);
          }(_pat_let46_0);
        }(_1884_v30);
        if (_source23.is_DC4) {
          let _1887___mcc_h0 = (_source23).cf7;
          let _1888_cf7 = _1887___mcc_h0;
          let _nw285 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1881_v29 = _nw285;
          let _1889_v31;
          _1889_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,(_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]);
          _1889_v31 = (_1889_v31).update(p0, p0);
          let _1890_v32;
          let _1891_v33;
          let _1892_v34;
          let _out32;
          let _out33;
          let _out34;
          let _outcollector12 = (_this).m6(p1, globalState);
          _out32 = _outcollector12[0];
          _out33 = _outcollector12[1];
          _out34 = _outcollector12[2];
          _1890_v32 = _out32;
          _1891_v33 = _out33;
          _1892_v34 = _out34;
          let _1893_v35;
          _1893_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1891_v33,_1888_cf7);
          _1891_v33 = (_dafny.ZERO).minus((((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]) ? ((_this).f9) : (new BigNumber((_1893_v35).length))));
        } else if (_source23.is_DC5) {
          let _1894___mcc_h1 = (_source23).cf8;
          let _1895___mcc_h2 = (_source23).cf9;
          let _1896___mcc_h3 = (_source23).cf10;
          let _1897___mcc_h4 = (_source23).cf11;
          let _1898_cf11 = _1897___mcc_h4;
          let _1899_cf10 = _1896___mcc_h3;
          let _1900_cf9 = _1895___mcc_h2;
          let _1901_cf8 = _1894___mcc_h1;
          _1900_cf9 = _module.__default.fm14(globalState);
          _1900_cf9 = false;
          let _1902_v36;
          _1902_v36 = _dafny.Map.Empty.slice().updateUnsafe(true,_1899_cf10);
          let _index297 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1881_v29).length));
          (_1881_v29)[_index297] = (new BigNumber(((_1902_v36).update(p1, (_this).f9)).length)).plus(_1898_cf11);
          let _index298 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_1881_v29).length));
          (_1881_v29)[_index298] = (_this).f9;
          let _1903_v37;
          let _nw286 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1903_v37 = _nw286;
          let _1904_v38;
          _1904_v38 = _dafny.Set.fromElements(_1899_cf10, _1899_cf10);
          let _1905_v39;
          _1905_v39 = _module.D16.create_DC36(_1904_v38, _1901_cf8, (_this).f4, _1901_cf8);
          let _1906_v40;
          _1906_v40 = _dafny.MultiSet.fromElements(_1905_v39, _1905_v39);
          let _index299 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_1903_v37).length));
          (_1903_v37)[_index299] = _1906_v40;
          let _index300 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_1903_v37).length));
          (_1903_v37)[_index300] = _1906_v40;
        } else if (_source23.is_DC6) {
          (globalState).f0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_1907_v0) => function (_1908_i4) {
            return _1907_v0;
          })(_1850_v0)), _dafny.Seq.UnicodeFromString("ng")), _dafny.Seq.Concat(_1852_v2, _1852_v2));
          let _1909_v41;
          let _1910_v42;
          let _1911_v43;
          let _out35;
          let _out36;
          let _out37;
          let _outcollector13 = (_this).m6(p1, globalState);
          _out35 = _outcollector13[0];
          _out36 = _outcollector13[1];
          _out37 = _outcollector13[2];
          _1909_v41 = _out35;
          _1910_v42 = _out36;
          _1911_v43 = _out37;
          r0 = (_this).f9;
          let _1912_v44;
          let _nw287 = new _module.C1();
          _nw287.__ctor(_1852_v2, _1909_v41);
          _1912_v44 = _nw287;
        } else {
          let _1913___mcc_h5 = (_source23).cf6;
          let _1914_cf6 = _1913___mcc_h5;
          let _1915_v45;
          _1915_v45 = _dafny.MultiSet.fromElements((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))], _module.__default.fm6((_this).f9, globalState));
          let _1916_v46;
          _1916_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]);
          let _1917_v47;
          _1917_v47 = _dafny.Seq.of(_this.f3);
          let _rhs249 = (((_1916_v46).contains(_this.f3)) ? ((_1916_v46).get(_this.f3)) : ((_1917_v47)[_module.__default.safeIndex(new BigNumber(818), new BigNumber((_1917_v47).length))]));
          let _rhs250 = ((_dafny.MultiSet.fromElements(!((_this).f4))).Intersect(_1915_v45)).Difference(_1915_v45);
          let _lhs183 = _this;
          _lhs183.f8 = _rhs249;
          _1915_v45 = _rhs250;
          r0 = (_this).f9;
          let _1918_v48;
          _1918_v48 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f9);
          let _1919_v49;
          _1919_v49 = _dafny.Seq.of((_this).f9);
          let _1920_v50;
          _1920_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_1853_v4).fm10((((_1918_v48).contains(_this.f3)) ? ((_1918_v48).get(_this.f3)) : ((_this).f9)), p0, _dafny.MultiSet.FromArray(_1919_v49), globalState));
          let _1921_v51;
          _1921_v51 = _dafny.Set.fromElements(_this.f8, p1, (_this).f4);
          r0 = (((_1920_v50).contains((_this).f9)) ? ((_1920_v50).get((_this).f9)) : (new BigNumber((_1921_v51).length)));
          let _1922_v52;
          _1922_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1857_v8);
          _1922_v52 = (_1922_v52).update((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))], _1914_cf6);
        }
        let _1923_v53;
        _1923_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_this.f8);
        let _1924_v54;
        let _nw288 = new _module.C2();
        _nw288.__ctor(!((((_1923_v53).contains((_this).f9)) ? ((_1923_v53).get((_this).f9)) : ((_this).f4))) || (p1), p1, p1, true);
        _1924_v54 = _nw288;
        let _1925_v55;
        _1925_v55 = _dafny.Seq.of(p0);
        let _1926_v56;
        _1926_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1925_v55,_dafny.Seq.Concat((_1924_v54).fm17(globalState), _dafny.Seq.UnicodeFromString("g")));
        _1926_v56 = (_1926_v56).update(_1925_v55, _dafny.Seq.update(_1852_v2, _module.__default.safeIndex((_this).f9, new BigNumber((_1852_v2).length)), _1850_v0));
        let _index301 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length));
        (_1881_v29)[_index301] = new BigNumber(85);
        let _index302 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length));
        (_1881_v29)[_index302] = (_this).f9;
        if ((_this).f4) {
          let _1927_v57;
          _1927_v57 = _module.D19.create_DC45(_module.__default.safeModuloInt(new BigNumber(47), (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]), (_this).f9, (_1924_v54).f10, (_1924_v54).f11, _1925_v55);
          _1927_v57 = _1927_v57;
          let _1928_v58;
          _1928_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1852_v2, _1852_v2),!((_this).f4));
          let _1929_v59;
          _1929_v59 = _module.D1.create_DC5((_this).f4, (_1924_v54).f11, (_this).f9, (_this).f9);
          let _index303 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
          (_1857_v8)[_index303] = (((_1928_v58).contains(_module.__default.fm30(_1929_v59, new BigNumber((_1923_v53).length), (_this).f9, globalState))) ? ((_1928_v58).get(_module.__default.fm30(_1929_v59, new BigNumber((_1923_v53).length), (_this).f9, globalState))) : (false));
          let _1930_v60;
          _1930_v60 = _module.D2.create_DC8(_1857_v8, _1852_v2);
          let _1931_v61;
          _1931_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1930_v60,(_this).f9);
          let _index304 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length));
          (_1857_v8)[_index304] = ((_this).fm5((_1924_v54).f10, _1852_v2, globalState)).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1931_v61,p0)).length)));
          let _index305 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length));
          (_1881_v29)[_index305] = (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))];
          let _1932_v62;
          _1932_v62 = _dafny.Seq.of((_this).f9);
          let _1933_v63;
          _1933_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1932_v62,_1850_v0);
          let _1934_v64;
          _1934_v64 = _module.D17.create_DC37(_1933_v63);
          let _1935_v65;
          _1935_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1934_v64,(_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]);
          _1935_v65 = (_1935_v65).update(_module.D17.create_DC37(_1933_v63), (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]);
        } else {
          _1852_v2 = _dafny.Seq.UnicodeFromString("empl");
          let _1936_v66;
          _1936_v66 = _dafny.Seq.of((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]);
          let _1937_v67;
          _1937_v67 = _dafny.Map.Empty.slice().updateUnsafe((_1924_v54).f10,(_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]);
          r0 = _module.__default.fm31(!((((_1923_v53).contains((_this).f9)) ? ((_1923_v53).get((_this).f9)) : ((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]))), (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))], !_dafny.Seq.contains(_1936_v66, (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]), _dafny.Seq.update(_1925_v55, _module.__default.safeIndex((((_1937_v67).contains((((_1923_v53).contains((_dafny.ZERO).minus(new BigNumber((_1852_v2).length)))) ? ((_1923_v53).get((_dafny.ZERO).minus(new BigNumber((_1852_v2).length)))) : ((_this).f4)))) ? ((_1937_v67).get((((_1923_v53).contains((_dafny.ZERO).minus(new BigNumber((_1852_v2).length)))) ? ((_1923_v53).get((_dafny.ZERO).minus(new BigNumber((_1852_v2).length)))) : ((_this).f4)))) : (new BigNumber((_1925_v55).length))), new BigNumber((_1925_v55).length)), (_this).f4), globalState);
          let _1938_v68;
          _1938_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1857_v8,(_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]);
          let _1939_v69;
          _1939_v69 = _module.D18.create_DC42((_this).f9);
          _1938_v68 = (_1938_v68).update(_1857_v8, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_this).f9), _module.__default.safeIndex((_this).f9, new BigNumber((_dafny.Seq.of((_this).f9)).length)), (_this).f9), _dafny.Seq.of((_1939_v69).dtor_cf64))).length));
          let _1940_v70;
          let _nw289 = Array((new BigNumber(13)).toNumber());
          _nw289[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1936_v66, _1936_v66);
          _nw289[(_dafny.ONE).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(2)).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(3)).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_1941_v29) => function (_1942_i5) {
            return (_1941_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1941_v29).length))];
          })(_1881_v29));
          _nw289[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f9, (_this).f9, (_this).f9), _1936_v66);
          _nw289[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_module.__default.fm19(globalState), _module.__default.safeIndex(new BigNumber((_1936_v66).length), new BigNumber((_module.__default.fm19(globalState)).length)), new BigNumber(158));
          _nw289[(new BigNumber(7)).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(8)).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-327)), ((_1943_v29) => function (_1944_i6) {
            return new BigNumber(((_module.D20.create_DC48(_dafny.Seq.of((_1943_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1943_v29).length))]))).dtor_cf80).length);
          })(_1881_v29));
          _nw289[(new BigNumber(10)).toNumber()] = _1936_v66;
          _nw289[(new BigNumber(11)).toNumber()] = ((p1) ? (_1936_v66) : (_1936_v66));
          _nw289[(new BigNumber(12)).toNumber()] = _1936_v66;
          _1940_v70 = _nw289;
          let _1945_v72;
          let _init57 = ((_1946_v55, _1947_v2, _1948_v54, _1949_v0, _1950_v29) => function (_1951_i7) {
            return (function () {
              let _coll68 = new _dafny.Map();
              for (const _compr_68 of _dafny.IntegerRange(new BigNumber(137), new BigNumber(-50))) {
                let _1952_v71 = _compr_68;
                if (((new BigNumber(137)).isLessThanOrEqualTo(_1952_v71)) && ((_1952_v71).isLessThan(new BigNumber(-50)))) {
                  _coll68.push([_module.__default.safeDivisionInt(_1952_v71, new BigNumber((_1947_v2).length)),_dafny.Map.Empty.slice().updateUnsafe((_1948_v54).f10,_1949_v0)]);
                }
              }
              return _coll68;
            }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_1946_v55)[_module.__default.safeIndex((_1950_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1950_v29).length))], new BigNumber((_1946_v55).length))])).length),_dafny.Map.Empty.slice().updateUnsafe((_1948_v54).f11,new _dafny.CodePoint('p'.codePointAt(0)))));
          })(_1925_v55, _1852_v2, _1924_v54, _1850_v0, _1881_v29);
          let _nw290 = Array((new BigNumber(24)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw290.length); _i0_57++) {
            _nw290[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _1945_v72 = _nw290;
          let _1953_v73;
          _1953_v73 = _dafny.Map.Empty.slice().updateUnsafe(!((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))]),_1850_v0);
          let _index306 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1945_v72).length));
          (_1945_v72)[_index306] = _dafny.Map.Empty.slice().updateUnsafe((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))],_1953_v73);
          let _1954_v74;
          let _init58 = ((_1955_v29) => function (_1956_i8) {
            return _module.D12.create_DC27((_1955_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1955_v29).length))]);
          })(_1881_v29);
          let _nw291 = Array((new BigNumber(20)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw291.length); _i0_58++) {
            _nw291[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1954_v74 = _nw291;
          let _1957_v75;
          _1957_v75 = _dafny.Set.fromElements(new BigNumber((_1925_v55).length));
          let _1958_v76;
          _1958_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))],_1957_v75);
          let _1959_v77;
          _1959_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1953_v73);
          let _1960_v78;
          _1960_v78 = _dafny.Seq.of(_1954_v74, ((p0) ? (_1954_v74) : (_1954_v74)), _1954_v74, _1954_v74, _1954_v74);
          let _index307 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1945_v72).length));
          let _rhs251 = ((_1957_v75).Intersect((((_1958_v76).contains((_this).f9)) ? ((_1958_v76).get((_this).f9)) : (_dafny.Set.fromElements((_this).f9, (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]))))).IsProperSubsetOf(_1957_v75);
          let _rhs252 = _1940_v70;
          let _rhs253 = _1959_v77;
          let _rhs254 = (_1960_v78)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_1852_v2).length), (_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))])).length), new BigNumber((_1960_v78).length))];
          let _lhs184 = _1945_v72;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_1945_v72).length));
          r1 = _rhs251;
          _1940_v70 = _rhs252;
          _lhs184[_lhs185] = _rhs253;
          _1954_v74 = _rhs254;
          let _1961_v79;
          _1961_v79 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qsnalpc"));
          let _1962_v80;
          _1962_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))],_1961_v79);
          let _1963_v81;
          _1963_v81 = (((_1962_v80).contains((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))])) ? ((_1962_v80).get((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))])) : (_1961_v79));
          let _1964_v82;
          let _nw292 = Array((new BigNumber(28)).toNumber());
          _nw292[(_dafny.ZERO).toNumber()] = _1963_v81;
          _nw292[(_dafny.ONE).toNumber()] = _module.__default.fm35((_1924_v54).f10, globalState);
          _nw292[(new BigNumber(2)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(3)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(4)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(5)).toNumber()] = _1961_v79;
          _nw292[(new BigNumber(6)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(7)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(8)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(9)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(10)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(11)).toNumber()] = _1961_v79;
          _nw292[(new BigNumber(12)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(13)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(14)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(15)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(16)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(17)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(18)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(19)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(20)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(21)).toNumber()] = _1961_v79;
          _nw292[(new BigNumber(22)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(23)).toNumber()] = _1963_v81;
          _nw292[(new BigNumber(24)).toNumber()] = _1961_v79;
          _nw292[(new BigNumber(25)).toNumber()] = _1961_v79;
          _nw292[(new BigNumber(26)).toNumber()] = _module.__default.fm41((_this).f9, globalState);
          _nw292[(new BigNumber(27)).toNumber()] = _1963_v81;
          _1964_v82 = _nw292;
          let _index308 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length));
          let _rhs255 = _1964_v82;
          let _rhs256 = _module.__default.safeModuloInt((_this).f9, ((_1881_v29)[_module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length))]).multipliedBy(new BigNumber((_1925_v55).length)));
          let _rhs257 = (_this).f9;
          let _rhs258 = _this.f3;
          let _lhs186 = _1881_v29;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_1881_v29).length));
          _1964_v82 = _rhs255;
          _lhs186[_lhs187] = _rhs256;
          r0 = _rhs257;
          r1 = _rhs258;
        }
      } else {
        let _1965_v83;
        _1965_v83 = _dafny.Seq.of((_this).f9, (_this).f9);
        let _index309 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1881_v29).length));
        (_1881_v29)[_index309] = new BigNumber((_1965_v83).length);
        let _index310 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1881_v29).length));
        (_1881_v29)[_index310] = new BigNumber(258);
        let _1966_v84;
        let _nw293 = new _module.C3();
        _nw293.__ctor(_this.f3, p0);
        _1966_v84 = _nw293;
        let _1967_v85;
        let _nw294 = new _module.C2();
        _nw294.__ctor((_this).f4, true, _this.f3, (_this).f4);
        _1967_v85 = _nw294;
        let _1968_v86;
        _1968_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1850_v0,_1967_v85);
        let _1969_v87;
        _1969_v87 = _dafny.MultiSet.fromElements((_1881_v29)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_1881_v29).length))], (_1881_v29)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_1881_v29).length))], new BigNumber((_1968_v86).length), (_this).f9);
        let _1970_v88;
        _1970_v88 = _dafny.Seq.of(true, !(!((_1857_v8)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_1857_v8).length))])), (_dafny.MultiSet.fromElements(new BigNumber((_1965_v83).length))).IsSubsetOf((_1969_v87).update((_this).f9, _module.__default.abs(new BigNumber(292)))));
        _1970_v88 = _dafny.Seq.update(_1970_v88, _module.__default.safeIndex(new BigNumber((_1852_v2).length), new BigNumber((_1970_v88).length)), _1967_v85.f3);
        (globalState).f0 = _1967_v85.f3;
        (_this).f8 = (_1967_v85.f3) === (p0);
      }
      let _index311 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1881_v29).length));
      (_1881_v29)[_index311] = (_this).f9;
      let _1971_v89;
      _1971_v89 = _dafny.MultiSet.fromElements(_this.f3);
      let _index312 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1881_v29).length));
      (_1881_v29)[_index312] = new BigNumber((_1971_v89).cardinality());
      let _1972_v90;
      _1972_v90 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1852_v2);
      r0 = (((_1972_v90).contains((_1881_v29)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_1881_v29).length))])) ? ((_1881_v29)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_1881_v29).length))]) : ((_this).f9));
      r1 = p1;
      return [r0, r1];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      (_this).f8 = p0;
      let _1973_v0;
      let _nw295 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
      _1973_v0 = _nw295;
      let _1974_v1;
      let _nw296 = new _module.C3();
      _nw296.__ctor(p0, !(_this.f3));
      _1974_v1 = _nw296;
      let _1975_v2;
      _1975_v2 = _dafny.Seq.of(_1974_v1, _1974_v1);
      let _1976_v3;
      _1976_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1975_v2);
      let _1977_v4;
      _1977_v4 = _dafny.Seq.UnicodeFromString("jrkmvyivs");
      let _1978_v5;
      _1978_v5 = _dafny.Set.fromElements(_1975_v2, _dafny.Seq.update((((_1976_v3).contains((_this).f9)) ? ((_1976_v3).get((_this).f9)) : (_1975_v2)), _module.__default.safeIndex(new BigNumber((_1977_v4).length), new BigNumber(((((_1976_v3).contains((_this).f9)) ? ((_1976_v3).get((_this).f9)) : (_1975_v2))).length)), _1974_v1));
      let _index313 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1973_v0).length));
      (_1973_v0)[_index313] = _1978_v5;
      let _index314 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1973_v0).length));
      let _rhs259 = _this.f3;
      let _rhs260 = ((_1978_v5).Difference(_1978_v5)).Intersect(_1978_v5);
      let _lhs188 = _1973_v0;
      let _lhs189 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1973_v0).length));
      r0 = _rhs259;
      _lhs188[_lhs189] = _rhs260;
      let _1979_v6;
      let _nw297 = Array((new BigNumber(6)).toNumber()).fill([]);
      _1979_v6 = _nw297;
      let _1980_v7;
      let _init59 = function (_1981_i0) {
        return true;
      };
      let _nw298 = Array((new BigNumber(19)).toNumber());
      for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw298.length); _i0_59++) {
        _nw298[_i0_59] = _init59(new BigNumber(_i0_59));
      }
      _1980_v7 = _nw298;
      let _1982_v8;
      _1982_v8 = _dafny.Set.fromElements((_this).f9, new BigNumber((_1977_v4).length));
      let _1983_v9;
      _1983_v9 = _module.D16.create_DC36(_1982_v8, (_this).f4, (_this).f4, false);
      let _nw299 = Array((new BigNumber(23)).toNumber());
      _nw299[(_dafny.ZERO).toNumber()] = _1980_v7;
      _nw299[(_dafny.ONE).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(2)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(3)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(4)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(5)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(6)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(7)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(8)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(9)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(10)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(11)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(12)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(13)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(14)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(15)).toNumber()] = ((_this.f8) ? (_1980_v7) : (_1980_v7));
      _nw299[(new BigNumber(16)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(17)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(18)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(19)).toNumber()] = _1980_v7;
      _nw299[(new BigNumber(20)).toNumber()] = ((p0) ? (_1980_v7) : (_1980_v7));
      _nw299[(new BigNumber(21)).toNumber()] = (((_1983_v9).dtor_cf52) ? (_1980_v7) : (_1980_v7));
      _nw299[(new BigNumber(22)).toNumber()] = _1980_v7;
      _1979_v6 = _nw299;
      _1982_v8 = _1982_v8;
      let _1984_v10;
      _1984_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1974_v1,_1974_v1);
      let _1985_v11;
      _1985_v11 = _dafny.Seq.of(_1984_v10, _1984_v10);
      let _1986_v12;
      let _nw300 = Array((new BigNumber(3)).toNumber());
      _nw300[(_dafny.ZERO).toNumber()] = _1985_v11;
      _nw300[(_dafny.ONE).toNumber()] = _1985_v11;
      _nw300[(new BigNumber(2)).toNumber()] = _1985_v11;
      _1986_v12 = _nw300;
      let _index315 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1986_v12).length));
      (_1986_v12)[_index315] = _1985_v11;
      let _index316 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1986_v12).length));
      (_1986_v12)[_index316] = _dafny.Seq.Concat(_dafny.Seq.of((_1984_v10).update(_1974_v1, _1974_v1)), _1985_v11);
      let _1987_v13;
      let _init60 = function (_1988_i2) {
        return _dafny.Seq.of((_this).f9, (_this).f9, (_this).f9);
      };
      let _nw301 = Array((new BigNumber(21)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw301.length); _i0_60++) {
        _nw301[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _1987_v13 = _nw301;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1987_v13).length))) {
        let _1989_i1 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1989_i1)) && ((_1989_i1).isLessThan(new BigNumber((_1987_v13).length))))) {
          (_1987_v13)[(_1989_i1)] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(297)), function (_1990_i3) {
            return _module.__default.safeDivisionInt((_this).f9, (_this).f9);
          });
        }
      }
      r0 = _1974_v1.f3;
      let _1991_v14;
      _1991_v14 = _dafny.MultiSet.fromElements(_module.__default.fm14(globalState), !(((_this).f4) || (_this.f3)), (_1974_v1.f3) && (_this.f3), !((_1974_v1).f4) || (p0));
      r1 = (((_1991_v14).contains(_module.__default.fm6(new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()), globalState))) ? ((_1991_v14).get(_module.__default.fm6(new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()), globalState))) : ((_this).f9));
      let _1992_v15;
      let _nw302 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _1992_v15 = _nw302;
      r2 = _1992_v15;
      return [r0, r1, r2];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f3 = false;
      this._f4 = false;
      this.f7 = _dafny.ZERO;
      this._f6 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f6, f7, f3, f4) {
      let _this = this;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm0(globalState) {
      let _this = this;
      if (_this.f3) {
        return (new BigNumber((_dafny.MultiSet.fromElements((_this).f4)).cardinality())).multipliedBy(_this.f7);
      } else {
        return new BigNumber(192);
      }
    };
    fm3(globalState) {
      let _this = this;
      return (_this.f7).multipliedBy((_dafny.ZERO).minus(_this.f7));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      (_this).f7 = _this.f7;
      let _1993_v0;
      _1993_v0 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f7), _dafny.MultiSet.fromElements(p2));
      let _1994_v1;
      _1994_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1993_v0,p0);
      (_this).f3 = (((_1994_v1).contains(_1993_v0)) ? ((_1994_v1).get(_1993_v0)) : (_module.__default.fm4(_this.f3, globalState)));
      let _1995_v2;
      _1995_v2 = _module.D1.create_DC5((_this).f4, _this.f3, p2, _this.f7);
      let _1996_v3;
      let _nw303 = new _module.C4();
      _nw303.__ctor(false, new BigNumber(367), _this.f3, ((_this.f3) ? (_this.f3) : ((_1995_v2).dtor_cf8)));
      _1996_v3 = _nw303;
      let _1997_v4;
      _1997_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
      if ((new BigNumber(((_1997_v4).Merge(_1997_v4)).length)).isLessThanOrEqualTo(((_1996_v3.f8) ? ((_1996_v3).f9) : (p2)))) {
        (_1996_v3).f8 = p0;
        if (_1996_v3.f8) {
          let _1998_v5;
          _1998_v5 = _dafny.Set.fromElements(p0);
          let _rhs261 = (_1998_v5).IsProperSubsetOf(_1998_v5);
          let _rhs262 = new BigNumber(735);
          let _rhs263 = (_this).f4;
          let _lhs190 = globalState;
          let _lhs191 = globalState;
          _lhs190.f0 = _rhs261;
          r1 = _rhs262;
          _lhs191.f0 = _rhs263;
          (_1996_v3).f8 = _1996_v3.f8;
          let _1999_v6;
          _1999_v6 = _dafny.Seq.UnicodeFromString("yuup");
          _1999_v6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1999_v6, _module.__default.fm30(_1995_v2, (_dafny.ZERO).minus(p2), p2, globalState)), (_this).f6);
          let _2000_v7;
          _2000_v7 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, _1999_v6);
          r1 = (((_2000_v7).contains((_this).f6)) ? ((_2000_v7).get((_this).f6)) : (_this.f7));
          let _2001_v8;
          _2001_v8 = _dafny.MultiSet.fromElements(p2, _this.f7);
          _1999_v6 = _dafny.Seq.Concat(_dafny.Seq.update(_1999_v6, _module.__default.safeIndex(new BigNumber((_2001_v8).cardinality()), new BigNumber((_1999_v6).length)), new _dafny.CodePoint('r'.codePointAt(0))), _dafny.Seq.Concat(_1999_v6, (_this).f6));
        } else {
          let _2002_v9;
          let _nw304 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _2002_v9 = _nw304;
          _2002_v9 = _2002_v9;
          let _2003_v10;
          _2003_v10 = _dafny.Seq.UnicodeFromString("xbfhlbhiw");
          let _2004_v11;
          let _nw305 = Array((new BigNumber(8)).toNumber());
          _2004_v11 = _nw305;
          let _2005_v12;
          let _init61 = ((_2006_v3) => function (_2007_i0) {
            return _dafny.Set.fromElements((_2006_v3).f9);
          })(_1996_v3);
          let _nw306 = Array((new BigNumber(16)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw306.length); _i0_61++) {
            _nw306[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2005_v12 = _nw306;
          let _2008_v13;
          _2008_v13 = _module.D15.create_DC32(_2005_v12);
          let _index317 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_2004_v11).length));
          (_2004_v11)[_index317] = _2008_v13;
          let _2009_v14;
          let _nw307 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2009_v14 = _nw307;
          let _2010_v15;
          let _nw308 = Array((new BigNumber(17)).toNumber());
          _nw308[(_dafny.ZERO).toNumber()] = _2009_v14;
          _nw308[(_dafny.ONE).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(2)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(3)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(4)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(5)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(6)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(7)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(8)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(9)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(10)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(11)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(12)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(13)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(14)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(15)).toNumber()] = _2009_v14;
          _nw308[(new BigNumber(16)).toNumber()] = _2009_v14;
          _2010_v15 = _nw308;
          let _index318 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_2010_v15).length));
          (_2010_v15)[_index318] = _2009_v14;
          let _2011_v16;
          let _init62 = function (_2012_i1) {
            return (_2012_i1).minus(new BigNumber((_dafny.Seq.of(new BigNumber(601))).length));
          };
          let _nw309 = Array((new BigNumber(9)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw309.length); _i0_62++) {
            _nw309[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2011_v16 = _nw309;
          let _pat_let_tv57 = _2005_v12;
          let _index319 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_2004_v11).length));
          let _index320 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_2010_v15).length));
          let _rhs264 = _2011_v16;
          let _rhs265 = _dafny.Seq.Concat(((_this.f3) ? (_2003_v10) : ((_this).f6)), _2003_v10);
          let _rhs266 = function (_pat_let48_0) {
            return function (_2013_dt__update__tmp_h0) {
              return function (_pat_let49_0) {
                return function (_2014_dt__update_hcf46_h0) {
                  return _module.D15.create_DC32(_2014_dt__update_hcf46_h0);
                }(_pat_let49_0);
              }(_pat_let_tv57);
            }(_pat_let48_0);
          }(_2008_v13);
          let _rhs267 = _2009_v14;
          let _lhs192 = _2004_v11;
          let _lhs193 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_2004_v11).length));
          let _lhs194 = _2010_v15;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_2010_v15).length));
          r0 = _rhs264;
          _2003_v10 = _rhs265;
          _lhs192[_lhs193] = _rhs266;
          _lhs194[_lhs195] = _rhs267;
          let _2015_v17;
          _2015_v17 = _dafny.Seq.of((_this).f4);
          let _2016_v18;
          _2016_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_2015_v17).length))).length),(_1996_v3).f9);
          let _2017_v19;
          _2017_v19 = _dafny.MultiSet.fromElements(_2016_v18, _2016_v18);
          let _2018_v20;
          let _nw310 = new _module.C4();
          _nw310.__ctor(!(_module.__default.fm21((_this).fm3(globalState), _2017_v19, globalState)), (_dafny.ZERO).minus(_this.f7), !(_1996_v3.f8), (_this).f4);
          _2018_v20 = _nw310;
          let _2019_v21;
          _2019_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1996_v3).f9);
          r1 = _module.__default.safeDivisionInt((_1996_v3).f9, _module.__default.safeModuloInt((_2018_v20).f9, (((_2019_v21).contains(_1996_v3.f8)) ? ((_2019_v21).get(_1996_v3.f8)) : ((_1996_v3).f9))));
          let _2020_v22;
          let _nw311 = new _module.C3();
          _nw311.__ctor(true, !(_1996_v3.f8));
          _2020_v22 = _nw311;
          let _2021_v23;
          _2021_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2020_v22,_1996_v3.f8);
          let _2022_v24;
          _2022_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f4),_2011_v16);
          let _2023_v25;
          _2023_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_2022_v24).contains(_dafny.Set.fromElements(_this.f3, _this.f3, _1996_v3.f8, !(true)))) ? ((_2022_v24).get(_dafny.Set.fromElements(_this.f3, _this.f3, _1996_v3.f8, !(true)))) : (_2011_v16)),_2018_v20.f8);
          let _2024_v26;
          _2024_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm33((_2018_v20).f9, _2015_v17, globalState)).length),(_this).f6);
          let _2025_v27;
          let _nw312 = Array((new BigNumber(20)).toNumber());
          _nw312[(_dafny.ZERO).toNumber()] = !(p0);
          _nw312[(_dafny.ONE).toNumber()] = !(_this.f3);
          _nw312[(new BigNumber(2)).toNumber()] = _1996_v3.f8;
          _nw312[(new BigNumber(3)).toNumber()] = _this.f3;
          _nw312[(new BigNumber(4)).toNumber()] = _1996_v3.f8;
          _nw312[(new BigNumber(5)).toNumber()] = _1996_v3.f8;
          _nw312[(new BigNumber(6)).toNumber()] = _1996_v3.f8;
          _nw312[(new BigNumber(7)).toNumber()] = !(new BigNumber(509)).isEqualTo(new BigNumber(232));
          _nw312[(new BigNumber(8)).toNumber()] = (((_2021_v23).contains(_2020_v22)) ? ((_2021_v23).get(_2020_v22)) : (!(_1996_v3.f8)));
          _nw312[(new BigNumber(9)).toNumber()] = ((_1996_v3).f9).isLessThanOrEqualTo((_1996_v3).f9);
          _nw312[(new BigNumber(10)).toNumber()] = _this.f3;
          _nw312[(new BigNumber(11)).toNumber()] = false;
          _nw312[(new BigNumber(12)).toNumber()] = ((_1996_v3).f9).isLessThan((_1996_v3).f9);
          _nw312[(new BigNumber(13)).toNumber()] = !(_1996_v3.f8);
          _nw312[(new BigNumber(14)).toNumber()] = _1996_v3.f8;
          _nw312[(new BigNumber(15)).toNumber()] = _2018_v20.f8;
          _nw312[(new BigNumber(16)).toNumber()] = p0;
          _nw312[(new BigNumber(17)).toNumber()] = (new BigNumber((_2024_v26).length)).isLessThan(new BigNumber((_2023_v25).length));
          _nw312[(new BigNumber(18)).toNumber()] = p0;
          _nw312[(new BigNumber(19)).toNumber()] = _1996_v3.f8;
          _2025_v27 = _nw312;
          let _index321 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2025_v27).length));
          (_2025_v27)[_index321] = _1996_v3.f8;
          let _index322 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2025_v27).length));
          (_2025_v27)[_index322] = !(!(_module.__default.fm21((_1996_v3).f9, (_2017_v19).Union((_2017_v19).update(function () {
            let _coll69 = new _dafny.Map();
            for (const _compr_69 of (p1).Keys.Elements) {
              let _2026_v28 = _compr_69;
              if ((p1).contains(_2026_v28)) {
                _coll69.push([(_2026_v28).multipliedBy((_1996_v3).f9),_this.f7]);
              }
            }
            return _coll69;
          }(), _module.__default.abs((_1996_v3).f9))), globalState)));
        }
        let _2027_v29;
        let _nw313 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2027_v29 = _nw313;
        let _index323 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2027_v29).length));
        (_2027_v29)[_index323] = p2;
        let _2028_v30;
        _2028_v30 = _dafny.Seq.of(p2);
        let _index324 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2027_v29).length));
        (_2027_v29)[_index324] = ((_module.D19.create_DC46((_1996_v3).f9, _this.f3, (_this).f4, _module.D1.create_DC5(_this.f3, _1996_v3.f8, new BigNumber(((_this).f6).length), (_1996_v3).f9), new BigNumber((_2028_v30).length))).dtor_cf78).plus(((_1996_v3).f9).multipliedBy((_1996_v3).f9));
        let _2029_v31;
        _2029_v31 = _dafny.Seq.UnicodeFromString("aqudafc");
        _2029_v31 = (_this).f6;
        let _2030_v32;
        _2030_v32 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_this.f7, (_2027_v29)[_module.__default.safeIndex(new BigNumber(322), new BigNumber((_2027_v29).length))]), (p2).plus(p2));
        _2030_v32 = _dafny.MultiSet.fromElements(_this.f7, new BigNumber((_2028_v30).length), _this.f7, (_1996_v3).f9, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2028_v30, _2028_v30)).length)));
      } else {
        let _2031_v33;
        _2031_v33 = _dafny.MultiSet.fromElements((_1996_v3).f9);
        r1 = _module.__default.safeDivisionInt((_1996_v3).f9, (((_2031_v33).contains((_1996_v3).f9)) ? ((_2031_v33).get((_1996_v3).f9)) : (_this.f7)));
        let _2032_v34;
        _2032_v34 = _dafny.Set.fromElements(p0);
        let _2033_v35;
        _2033_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gfiwjkr"),_2032_v34);
        let _2034_v36;
        _2034_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f7);
        let _2035_v37;
        _2035_v37 = _dafny.Seq.of(_2034_v36, _2034_v36, _dafny.Map.Empty.slice().updateUnsafe(false,(_1996_v3).f9));
        let _2036_v38;
        _2036_v38 = new _dafny.CodePoint('g'.codePointAt(0));
        let _2037_v39;
        _2037_v39 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2036_v38);
        _2033_v35 = (_2033_v35).update(_module.__default.fm30(_1995_v2, _this.f7, new BigNumber(((_2035_v37)[_module.__default.safeIndex(new BigNumber((_2037_v39).length), new BigNumber((_2035_v37).length))]).length), globalState), _2032_v34);
        (_this).f7 = _module.__default.safeDivisionInt((_this.f7).multipliedBy(new BigNumber(((_this).f6).length)), _this.f7);
        let _2038_v40;
        let _init63 = ((_2039_v33) => function (_2040_i2) {
          return _2039_v33;
        })(_2031_v33);
        let _nw314 = Array((new BigNumber(15)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw314.length); _i0_63++) {
          _nw314[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2038_v40 = _nw314;
        _2038_v40 = _2038_v40;
        let _2041_v41;
        let _init64 = function (_2042_i4) {
          return (_this).f4;
        };
        let _nw315 = Array((new BigNumber(4)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw315.length); _i0_64++) {
          _nw315[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _2041_v41 = _nw315;
        let _2043_v42;
        _2043_v42 = _module.D2.create_DC8(_2041_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-310)), ((_2044_v38) => function (_2045_i5) {
  return _2044_v38;
})(_2036_v38)));
        let _2046_v43;
        _2046_v43 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-888)), ((_2047_v38) => function (_2048_i3) {
          return _2047_v38;
        })(_2036_v38)), (_this).f6), (_2043_v42).dtor_cf14, (_this).f6, _dafny.Seq.Concat((_this).f6, (_this).f6));
        let _2049_v44;
        _2049_v44 = _dafny.Seq.UnicodeFromString("tnqupg");
        let _rhs268 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), ((_2050_v44) => function (_2051_i6) {
          return _2050_v44;
        })(_2049_v44));
        let _rhs269 = (((_this.f3) ? (p0) : ((_this).f4))) && ((_dafny.MultiSet.fromElements((_1996_v3).f9)).IsProperSubsetOf(_2031_v33));
        let _rhs270 = _this.f7;
        let _rhs271 = _dafny.Seq.Concat((_2043_v42).dtor_cf14, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vqw"), _dafny.Seq.UnicodeFromString("fetrpyq")));
        let _lhs196 = globalState;
        _2046_v43 = _rhs268;
        _lhs196.f0 = _rhs269;
        r1 = _rhs270;
        _2049_v44 = _rhs271;
      }
      let _source24 = _1995_v2;
      if (_source24.is_DC4) {
        let _2052___mcc_h0 = (_source24).cf7;
        let _2053_cf7 = _2052___mcc_h0;
        let _2054_v45;
        _2054_v45 = _dafny.Set.fromElements(_1996_v3.f8, _this.f3);
        let _2055_v46;
        _2055_v46 = _dafny.Seq.of(_2054_v45, _2054_v45, _2054_v45);
        let _2056_v47;
        _2056_v47 = new _dafny.CodePoint('p'.codePointAt(0));
        let _2057_v48;
        _2057_v48 = _dafny.MultiSet.fromElements(_2053_cf7);
        let _2058_v49;
        _2058_v49 = _dafny.Seq.of(_module.__default.fm18(new BigNumber(((_2055_v46)[_module.__default.safeIndex((_1996_v3).f9, new BigNumber((_2055_v46).length))]).length), _2056_v47, (_1996_v3).f9, globalState), p0, _1996_v3.f8, (_2057_v48).IsSubsetOf(_2057_v48), false);
        let _rhs272 = (_2058_v49)[_module.__default.safeIndex((_1996_v3).f9, new BigNumber((_2058_v49).length))];
        let _rhs273 = _1996_v3.f8;
        let _lhs197 = globalState;
        let _lhs198 = globalState;
        _lhs197.f0 = _rhs272;
        _lhs198.f0 = _rhs273;
        let _2059_v50;
        _2059_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v3,p0);
        let _2060_v51;
        let _nw316 = Array((new BigNumber(28)).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = _1996_v3.f8;
        _nw316[(_dafny.ONE).toNumber()] = !((_this).f4);
        _nw316[(new BigNumber(2)).toNumber()] = (_this).f4;
        _nw316[(new BigNumber(3)).toNumber()] = true;
        _nw316[(new BigNumber(4)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(5)).toNumber()] = false;
        _nw316[(new BigNumber(6)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(7)).toNumber()] = p0;
        _nw316[(new BigNumber(8)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(9)).toNumber()] = p0;
        _nw316[(new BigNumber(10)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(11)).toNumber()] = p0;
        _nw316[(new BigNumber(12)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(13)).toNumber()] = (((_2059_v50).contains(_1996_v3)) ? ((_2059_v50).get(_1996_v3)) : (_1996_v3.f8));
        _nw316[(new BigNumber(14)).toNumber()] = p0;
        _nw316[(new BigNumber(15)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(16)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(17)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(18)).toNumber()] = !(true);
        _nw316[(new BigNumber(19)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(20)).toNumber()] = !(_1996_v3.f8);
        _nw316[(new BigNumber(21)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(22)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(23)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(24)).toNumber()] = _1996_v3.f8;
        _nw316[(new BigNumber(25)).toNumber()] = p0;
        _nw316[(new BigNumber(26)).toNumber()] = _this.f3;
        _nw316[(new BigNumber(27)).toNumber()] = (_this).f4;
        _2060_v51 = _nw316;
        let _2061_v52;
        _2061_v52 = _module.D4.create_DC11(_2058_v49);
        let _2062_v53;
        _2062_v53 = _module.D19.create_DC45((_1996_v3).f9, _2053_cf7, _1996_v3.f8, false, (_2061_v52).dtor_cf18);
        let _2063_v54;
        _2063_v54 = _module.D19.create_DC47(_2062_v53);
        let _2064_v55;
        _2064_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2053_cf7,_2063_v54);
        let _2065_v56;
        _2065_v56 = _module.D18.create_DC43(_2053_cf7, p2, _1996_v3.f8);
        let _2066_v57;
        _2066_v57 = _dafny.Seq.of(new BigNumber(195));
        let _2067_v58;
        _2067_v58 = _module.D8.create_DC19(new BigNumber((_2066_v57).length), p2);
        let _2068_v59;
        _2068_v59 = _dafny.MultiSet.fromElements(_1996_v3.f8);
        let _2069_v60;
        _2069_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1996_v3).f9,_2058_v49);
        let _2070_v61;
        _2070_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2069_v60);
        let _2071_v62;
        _2071_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2057_v48,_module.__default.fm31(_1996_v3.f8, _2053_cf7, _this.f3, _2058_v49, globalState));
        let _2072_v63;
        let _nw317 = Array((new BigNumber(26)).toNumber());
        _nw317[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_2053_cf7);
        _nw317[(_dafny.ONE).toNumber()] = p2;
        _nw317[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_1996_v3).f9, _2053_cf7);
        _nw317[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_module.D2.create_DC8(_2060_v51, (_this).f6)).dtor_cf14, (_this).f6)).length);
        _nw317[(new BigNumber(4)).toNumber()] = new BigNumber(((_this).f6).length);
        _nw317[(new BigNumber(5)).toNumber()] = new BigNumber(((_2064_v55).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2065_v56,(_this).f4)).length), _2063_v54)).length);
        _nw317[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt((_1996_v3).f9, new BigNumber((_module.__default.fm42(_2067_v58, _1996_v3.f8, new BigNumber(557), globalState)).cardinality()));
        _nw317[(new BigNumber(7)).toNumber()] = (_1996_v3).f9;
        _nw317[(new BigNumber(8)).toNumber()] = new BigNumber(((_module.__default.fm20((_this).f4, globalState)).Difference(_2068_v59)).cardinality());
        _nw317[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_this.f7);
        _nw317[(new BigNumber(10)).toNumber()] = _this.f7;
        _nw317[(new BigNumber(11)).toNumber()] = new BigNumber(843);
        _nw317[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_1996_v3).f9);
        _nw317[(new BigNumber(13)).toNumber()] = _2053_cf7;
        _nw317[(new BigNumber(14)).toNumber()] = (_1996_v3).f9;
        _nw317[(new BigNumber(15)).toNumber()] = (_this.f7).minus(_this.f7);
        _nw317[(new BigNumber(16)).toNumber()] = (_1996_v3).f9;
        _nw317[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(((((_2070_v61).contains(p2)) ? ((_2070_v61).get(p2)) : (_2069_v60))).length), (_dafny.ZERO).minus((_1996_v3).f9));
        _nw317[(new BigNumber(18)).toNumber()] = new BigNumber((_module.__default.fm34(new BigNumber(((_this).f6).length), (_1996_v3).f9, globalState)).length);
        _nw317[(new BigNumber(19)).toNumber()] = _2053_cf7;
        _nw317[(new BigNumber(20)).toNumber()] = (_1996_v3).f9;
        _nw317[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_this.f7);
        _nw317[(new BigNumber(22)).toNumber()] = (_module.__default.fm31((_this).f4, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(p2, (_1996_v3).f9)).cardinality())), _1996_v3.f8, _dafny.Seq.of(true, (_this).f4, _this.f3, _1996_v3.f8), globalState)).multipliedBy((_1996_v3).f9);
        _nw317[(new BigNumber(23)).toNumber()] = _this.f7;
        _nw317[(new BigNumber(24)).toNumber()] = _this.f7;
        _nw317[(new BigNumber(25)).toNumber()] = ((_dafny.ZERO).minus(_this.f7)).minus(new BigNumber((_2071_v62).length));
        _2072_v63 = _nw317;
        let _2073_v64;
        _2073_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2056_v47,(_this).f4);
        let _index325 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_2072_v63).length));
        (_2072_v63)[_index325] = new BigNumber((_2073_v64).length);
        let _index326 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_2072_v63).length));
        (_2072_v63)[_index326] = (((_this.f3) ? (new BigNumber(((_this).f6).length)) : (_2053_cf7))).multipliedBy((_1996_v3).fm5(_1996_v3.f8, (_this).f6, globalState));
        let _pat_let_tv58 = _1996_v3;
        let _pat_let_tv59 = _1996_v3;
        let _source25 = function (_pat_let50_0) {
          return function (_2074_dt__update__tmp_h1) {
            return function (_pat_let51_0) {
              return function (_2075_dt__update_hcf30_h0) {
                return _module.D8.create_DC18(_2075_dt__update_hcf30_h0);
              }(_pat_let51_0);
            }(_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv58).f9,_pat_let_tv59.f8));
          }(_pat_let50_0);
        }(_module.D8.create_DC18(p1));
        if (_source25.is_DC19) {
          let _2076___mcc_h6 = (_source25).cf31;
          let _2077___mcc_h7 = (_source25).cf32;
          let _2078_cf32 = _2077___mcc_h7;
          let _2079_cf31 = _2076___mcc_h6;
          let _2080_v65;
          _2080_v65 = _dafny.Set.fromElements((_this).f6);
          let _2081_v66;
          _2081_v66 = _dafny.MultiSet.fromElements((_this).f6, _dafny.Seq.UnicodeFromString("cwncwwteh"));
          let _2082_v68;
          _2082_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2078_cf32,function () {
            let _coll70 = new _dafny.Set();
            for (const _compr_70 of (_2081_v66).Elements) {
              let _2083_v67 = _compr_70;
              if ((_2081_v66).contains(_2083_v67)) {
                _coll70.add(_2083_v67);
              }
            }
            return _coll70;
          }());
          let _2084_v70;
          _2084_v70 = _dafny.Seq.of(_2080_v65, _dafny.Set.fromElements((_this).f6), (((_2082_v68).contains(_2079_cf31)) ? ((_2082_v68).get(_2079_cf31)) : (function () {
            let _coll71 = new _dafny.Set();
            for (const _compr_71 of (_2080_v65).Elements) {
              let _2085_v69 = _compr_71;
              if ((_2080_v65).contains(_2085_v69)) {
                _coll71.add(_2085_v69);
              }
            }
            return _coll71;
          }())), _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("uurou")));
          let _2086_v71;
          _2086_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v3.f8,_2054_v45);
          let _2087_v72;
          _2087_v72 = _dafny.Map.Empty.slice().updateUnsafe((_1996_v3).f9,(_1996_v3).f9);
          let _2088_v73;
          _2088_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2086_v71).contains(_this.f3)) ? ((_2086_v71).get(_this.f3)) : (_2054_v45))).length),_2087_v72);
          let _rhs274 = ((_2080_v65).Difference((_2084_v70)[_module.__default.safeIndex(new BigNumber((_2088_v73).length), new BigNumber((_2084_v70).length))])).IsDisjointFrom(_2080_v65);
          let _rhs275 = _1995_v2;
          let _rhs276 = _2053_cf7;
          let _lhs199 = globalState;
          _lhs199.f0 = _rhs274;
          _1995_v2 = _rhs275;
          r1 = _rhs276;
          _2058_v49 = _module.__default.fm34((_1996_v3).f9, ((_1996_v3.f8) ? (_2079_cf31) : (_2053_cf7)), globalState);
          (_1996_v3).f8 = !(p0);
          (globalState).f0 = _1996_v3.f8;
        } else if (_source25.is_DC20) {
          let _2089___mcc_h8 = (_source25).cf33;
          let _2090___mcc_h9 = (_source25).cf34;
          let _2091_cf34 = _2090___mcc_h9;
          let _2092_cf33 = _2089___mcc_h8;
          (globalState).f0 = _1996_v3.f8;
          _2060_v51 = ((p0) ? (_2060_v51) : (((!(false)) ? (_2060_v51) : (_2060_v51))));
          let _2093_v74;
          _2093_v74 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1996_v3).f9);
          let _2094_v75;
          _2094_v75 = _dafny.Set.fromElements((_1996_v3).f9, (((_this).f4) ? (new BigNumber(459)) : (p2)), ((_dafny.ZERO).minus((_1996_v3).f9)).multipliedBy(_this.f7), new BigNumber((_2093_v74).length));
          _2094_v75 = _2094_v75;
          let _2095_v76;
          let _2096_v77;
          let _out38;
          let _out39;
          let _outcollector14 = (_1996_v3).m5(_1996_v3.f8, _1996_v3.f8, globalState);
          _out38 = _outcollector14[0];
          _out39 = _outcollector14[1];
          _2095_v76 = _out38;
          _2096_v77 = _out39;
        } else {
          let _2097___mcc_h10 = (_source25).cf30;
          let _2098_cf30 = _2097___mcc_h10;
          let _rhs277 = (_this).f4;
          let _rhs278 = (((_1996_v3.f8) ? ((_1996_v3).f9) : (new BigNumber(294)))).isEqualTo(p2);
          let _lhs200 = globalState;
          let _lhs201 = _this;
          _lhs200.f0 = _rhs277;
          _lhs201.f3 = _rhs278;
          (globalState).f0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_module.__default.fm30(_1995_v2, new BigNumber(((_this).f6).length), (_2072_v63)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_2072_v63).length))], globalState), _dafny.Seq.UnicodeFromString("dnoyfw")), (_this).f6);
          _1997_v4 = (_1997_v4).update(_dafny.areEqual((_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), function (_2099_i7) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          })), !(!(!(_this.f3))));
          let _2100_v78;
          _2100_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2060_v51,new BigNumber(169));
          _2100_v78 = (_2100_v78).update(_2060_v51, (_1996_v3).f9);
        }
        let _2101_v79;
        let _nw318 = new _module.C0();
        _nw318.__ctor();
        _2101_v79 = _nw318;
      } else if (_source24.is_DC5) {
        let _2102___mcc_h1 = (_source24).cf8;
        let _2103___mcc_h2 = (_source24).cf9;
        let _2104___mcc_h3 = (_source24).cf10;
        let _2105___mcc_h4 = (_source24).cf11;
        let _2106_cf11 = _2105___mcc_h4;
        let _2107_cf10 = _2104___mcc_h3;
        let _2108_cf9 = _2103___mcc_h2;
        let _2109_cf8 = _2102___mcc_h1;
        if (false) {
          let _2110_v80;
          let _init65 = function (_2111_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber(-863), new BigNumber(((_this).f6).length))).length));
          };
          let _nw319 = Array((new BigNumber(23)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw319.length); _i0_65++) {
            _nw319[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2110_v80 = _nw319;
          let _2112_v81;
          _2112_v81 = new _dafny.CodePoint('b'.codePointAt(0));
          let _2113_v82;
          _2113_v82 = _dafny.Map.Empty.slice().updateUnsafe(_2112_v81,(_1996_v3).f9);
          let _index327 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2110_v80).length));
          (_2110_v80)[_index327] = _2113_v82;
          let _index328 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_2110_v80).length));
          (_2110_v80)[_index328] = _2113_v82;
          let _2114_v83;
          _2114_v83 = _dafny.MultiSet.fromElements((_this).f4, _2109_cf8);
          (globalState).f0 = (_2114_v83).IsSubsetOf(_2114_v83);
          _2107_cf10 = ((_1996_v3).f9).multipliedBy(_2106_cf11);
          let _2115_v84;
          let _nw320 = Array((new BigNumber(18)).toNumber()).fill([]);
          _2115_v84 = _nw320;
          let _2116_v85;
          let _init66 = ((_2117_v3) => function (_2118_i9) {
            return _2117_v3.f8;
          })(_1996_v3);
          let _nw321 = Array((new BigNumber(7)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw321.length); _i0_66++) {
            _nw321[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2116_v85 = _nw321;
          let _index329 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_2115_v84).length));
          (_2115_v84)[_index329] = _2116_v85;
          let _index330 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_2115_v84).length));
          (_2115_v84)[_index330] = _2116_v85;
          let _2119_v86;
          let _nw322 = new _module.C3();
          _nw322.__ctor((_this.f3) || (_2108_cf9), (_this).f4);
          _2119_v86 = _nw322;
        } else {
          let _2120_v87;
          let _nw323 = Array((new BigNumber(26)).toNumber()).fill([]);
          _2120_v87 = _nw323;
          let _2121_v88;
          let _nw324 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2121_v88 = _nw324;
          let _index331 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_2120_v87).length));
          (_2120_v87)[_index331] = _2121_v88;
          let _2122_v89;
          let _init67 = ((_2123_cf8) => function (_2124_i10) {
            return _2123_cf8;
          })(_2109_cf8);
          let _nw325 = Array((new BigNumber(7)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw325.length); _i0_67++) {
            _nw325[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2122_v89 = _nw325;
          let _2125_v90;
          _2125_v90 = _dafny.MultiSet.fromElements(!(_1996_v3.f8), true);
          let _index332 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_2122_v89).length));
          (_2122_v89)[_index332] = (_2125_v90).IsSubsetOf(_2125_v90);
          let _2126_v91;
          _2126_v91 = _dafny.Seq.UnicodeFromString("wfcgkb");
          let _index333 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_2120_v87).length));
          let _index334 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_2122_v89).length));
          let _rhs279 = ((_1996_v3.f8) ? (_2121_v88) : (_2121_v88));
          let _rhs280 = _module.__default.safeDivisionInt((_2107_cf10).multipliedBy(p2), new BigNumber(276));
          let _rhs281 = !((_this).f4);
          let _rhs282 = (_this).f6;
          let _rhs283 = _2122_v89;
          let _lhs202 = _2120_v87;
          let _lhs203 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_2120_v87).length));
          let _lhs204 = _2122_v89;
          let _lhs205 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_2122_v89).length));
          _lhs202[_lhs203] = _rhs279;
          _2106_cf11 = _rhs280;
          _lhs204[_lhs205] = _rhs281;
          _2126_v91 = _rhs282;
          _2122_v89 = _rhs283;
          let _2127_v92;
          _2127_v92 = _dafny.MultiSet.fromElements((_1996_v3).f9, _2107_cf10);
          (_this).f3 = ((((_2127_v92).contains((_1996_v3).f9)) ? ((_2127_v92).get((_1996_v3).f9)) : (_2106_cf11))).isLessThanOrEqualTo(_2106_cf11);
          let _2128_v93;
          _2128_v93 = _dafny.Seq.of((_1997_v4).update(_this.f3, _2109_cf8), _1997_v4);
          r1 = new BigNumber((_2128_v93).length);
          let _2129_v94;
          _2129_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2107_cf10,(_dafny.ZERO).minus(_this.f7));
          _2129_v94 = (_2129_v94).update(p2, _2106_cf11);
          let _2130_v95;
          let _2131_v96;
          let _2132_v97;
          let _2133_v98;
          let _out40;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector15 = (_this).m3(_dafny.Seq.UnicodeFromString("pkowqfgre"), ((_1996_v3).f9).isLessThanOrEqualTo(_this.f7), (_2108_cf9) && (_this.f3), (_1996_v3).f9, globalState);
          _out40 = _outcollector15[0];
          _out41 = _outcollector15[1];
          _out42 = _outcollector15[2];
          _out43 = _outcollector15[3];
          _2130_v95 = _out40;
          _2131_v96 = _out41;
          _2132_v97 = _out42;
          _2133_v98 = _out43;
        }
        let _2134_v99;
        let _nw326 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2134_v99 = _nw326;
        let _index335 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_2134_v99).length));
        (_2134_v99)[_index335] = _dafny.MultiSet.fromElements(_2108_cf9);
        let _2135_v100;
        _2135_v100 = _dafny.MultiSet.fromElements(_2109_cf8, _2109_cf8);
        let _index336 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_2134_v99).length));
        (_2134_v99)[_index336] = _2135_v100;
        let _2136_v101;
        let _2137_v102;
        let _2138_v103;
        let _out44;
        let _out45;
        let _out46;
        let _outcollector16 = (_1996_v3).m6(_1996_v3.f8, globalState);
        _out44 = _outcollector16[0];
        _out45 = _outcollector16[1];
        _out46 = _outcollector16[2];
        _2136_v101 = _out44;
        _2137_v102 = _out45;
        _2138_v103 = _out46;
        r1 = p2;
      } else if (_source24.is_DC6) {
        let _2139_v104;
        _2139_v104 = _dafny.Seq.of(false, !(_this.f3), _1996_v3.f8, (_this).f4);
        let _2140_v105;
        _2140_v105 = _module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2139_v104).length),new BigNumber(180)));
        let _2141_v106;
        _2141_v106 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_1996_v3).f9);
        _2140_v105 = ((_1996_v3.f8) ? (_2140_v105) : (_module.D11.create_DC24(_2141_v106)));
        let _2142_v107;
        _2142_v107 = _dafny.Set.fromElements((_this).f4);
        (_1996_v3).f8 = (_2142_v107).IsProperSubsetOf(_dafny.Set.fromElements((_2139_v104)[_module.__default.safeIndex((_1996_v3).f9, new BigNumber((_2139_v104).length))]));
        if ((_module.__default.fm15((_this).f4, _this.f7, p2, _module.__default.fm4(false, globalState), globalState)).contains((_dafny.ZERO).minus(p2))) {
          let _2143_v108;
          _2143_v108 = _dafny.Map.Empty.slice().updateUnsafe((_1996_v3).f9,_module.__default.fm30(_1995_v2, (_1996_v3).f9, (_1996_v3).f9, globalState));
          let _2144_v109;
          _2144_v109 = _dafny.Seq.of((_1996_v3).f9);
          _2143_v108 = (_2143_v108).update((_dafny.ZERO).minus(((_1996_v3).f9).plus((_2144_v109)[_module.__default.safeIndex(p2, new BigNumber((_2144_v109).length))])), _dafny.Seq.Concat((_this).f6, (_this).f6));
          let _2145_v110;
          _2145_v110 = _dafny.Set.fromElements(_this.f7, _this.f7, (_1996_v3).f9);
          let _2146_v111;
          let _nw327 = Array((_dafny.ONE).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = new BigNumber((_2145_v110).length);
          _2146_v111 = _nw327;
          let _index337 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length));
          (_2146_v111)[_index337] = new BigNumber(((_this).f6).length);
          let _index338 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2146_v111).length));
          (_2146_v111)[_index338] = (_1996_v3).f9;
          let _index339 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length));
          let _index340 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2146_v111).length));
          let _rhs284 = (_dafny.ZERO).minus(p2);
          let _rhs285 = true;
          let _rhs286 = p2;
          let _lhs206 = _2146_v111;
          let _lhs207 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length));
          let _lhs208 = _1996_v3;
          let _lhs209 = _2146_v111;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2146_v111).length));
          _lhs206[_lhs207] = _rhs284;
          _lhs208.f8 = _rhs285;
          _lhs209[_lhs210] = _rhs286;
          let _pat_let_tv60 = _1996_v3;
          let _pat_let_tv61 = _2141_v106;
          let _pat_let_tv62 = globalState;
          let _pat_let_tv63 = _1996_v3;
          let _2147_v112;
          _2147_v112 = _module.D19.create_DC46(((_1996_v3).f9).multipliedBy(new BigNumber(626)), p0, _1996_v3.f8, function (_pat_let52_0) {
  return function (_2148_dt__update__tmp_h2) {
    return function (_pat_let53_0) {
      return function (_2149_dt__update_hcf8_h0) {
        return function (_pat_let54_0) {
          return function (_2150_dt__update_hcf9_h0) {
            return function (_pat_let55_0) {
              return function (_2151_dt__update_hcf10_h0) {
                return _module.D1.create_DC5(_2149_dt__update_hcf8_h0, _2150_dt__update_hcf9_h0, _2151_dt__update_hcf10_h0, (_2148_dt__update__tmp_h2).dtor_cf11);
              }(_pat_let55_0);
            }((_pat_let_tv63).f9);
          }(_pat_let54_0);
        }(!(_module.__default.fm21((_pat_let_tv60).f9, _dafny.MultiSet.fromElements(_pat_let_tv61), _pat_let_tv62)));
      }(_pat_let53_0);
    }((_this).f4);
  }(_pat_let52_0);
}(_1995_v2), (_2146_v111)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length))]);
          let _pat_let_tv64 = _1995_v2;
          let _pat_let_tv65 = _1996_v3;
          _2147_v112 = function (_pat_let56_0) {
            return function (_2152_dt__update__tmp_h3) {
              return function (_pat_let57_0) {
                return function (_2153_dt__update_hcf75_h0) {
                  return function (_pat_let58_0) {
                    return function (_2154_dt__update_hcf77_h0) {
                      return function (_pat_let59_0) {
                        return function (_2155_dt__update_hcf74_h0) {
                          return _module.D19.create_DC46(_2155_dt__update_hcf74_h0, _2153_dt__update_hcf75_h0, (_2152_dt__update__tmp_h3).dtor_cf76, _2154_dt__update_hcf77_h0, (_2152_dt__update__tmp_h3).dtor_cf78);
                        }(_pat_let59_0);
                      }((_pat_let_tv65).f9);
                    }(_pat_let58_0);
                  }(_pat_let_tv64);
                }(_pat_let57_0);
              }((_this).f4);
            }(_pat_let56_0);
          }(_module.D19.create_DC46(p2, _1996_v3.f8, p0, _1995_v2, (_dafny.ZERO).minus((_2146_v111)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length))])));
          let _index341 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_2146_v111).length));
          (_2146_v111)[_index341] = new BigNumber(932);
          (_1996_v3).f8 = (_this).f4;
        } else {
          let _2156_v113;
          let _nw328 = Array((new BigNumber(26)).toNumber()).fill(false);
          _2156_v113 = _nw328;
          let _index342 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_2156_v113).length));
          (_2156_v113)[_index342] = ((!((_this).f4)) ? (_this.f3) : (_1996_v3.f8));
          let _index343 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_2156_v113).length));
          (_2156_v113)[_index343] = _1996_v3.f8;
          let _2157_v114;
          _2157_v114 = _dafny.MultiSet.fromElements(_this.f7, p2);
          let _2158_v115;
          let _nw329 = new _module.C2();
          _nw329.__ctor((_2157_v114).IsSubsetOf(_2157_v114), !(_module.__default.fm6((_1996_v3).f9, globalState)), (_2156_v113)[_module.__default.safeIndex(new BigNumber(269), new BigNumber((_2156_v113).length))], _this.f3);
          _2158_v115 = _nw329;
          let _2159_v116;
          _2159_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v3.f8,new BigNumber((_2141_v106).length));
          (globalState).f0 = ((((_2159_v116).contains(!(p0))) ? ((_2159_v116).get(!(p0))) : (_this.f7))).isLessThanOrEqualTo(p2);
          let _2160_v117;
          _2160_v117 = _dafny.Seq.of(p2, (_dafny.ZERO).minus(_this.f7), p2);
          let _2161_v118;
          _2161_v118 = _module.D8.create_DC19(new BigNumber(((_this).f6).length), (_2158_v115).fm16((_1996_v3).f9, (_2139_v104)[_module.__default.safeIndex((_2160_v117)[_module.__default.safeIndex(_this.f7, new BigNumber((_2160_v117).length))], new BigNumber((_2139_v104).length))], _this.f7, p1, globalState));
          let _2162_v119;
          _2162_v119 = _dafny.Map.Empty.slice().updateUnsafe(_2161_v118,(_1996_v3).f9);
          _2162_v119 = (_2162_v119).update(_module.D8.create_DC19(_this.f7, _this.f7), new BigNumber(((_this).f6).length));
          (_1996_v3).f8 = (_2142_v107).IsDisjointFrom(_2142_v107);
        }
        r1 = _this.f7;
      } else {
        let _2163___mcc_h5 = (_source24).cf6;
        let _2164_cf6 = _2163___mcc_h5;
        let _2165_v120;
        _2165_v120 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f7);
        let _2166_v121;
        _2166_v121 = new _dafny.CodePoint('y'.codePointAt(0));
        _2165_v120 = (_2165_v120).update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-557)), function (_2167_i11) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }), _module.__default.safeIndex((_1996_v3).f9, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-557)), function (_2168_i11) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length)), _2166_v121), (new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)).minus(new BigNumber(-532)));
        (_this).f7 = _this.f7;
        let _2169_v122;
        _2169_v122 = _dafny.Set.fromElements(_1996_v3.f8);
        (_this).f7 = new BigNumber((_2169_v122).length);
        let _2170_v123;
        _2170_v123 = _dafny.Seq.of(new BigNumber(236));
        let _2171_v124;
        _2171_v124 = _dafny.Map.Empty.slice().updateUnsafe(_2170_v123,_2166_v121);
        let _2172_v125;
        _2172_v125 = _module.D17.create_DC37((_2171_v124).Merge(_2171_v124));
        let _source26 = _2172_v125;
        if (_source26.is_DC38) {
          let _2173___mcc_h11 = (_source26).cf56;
          let _2174_cf56 = _2173___mcc_h11;
          let _2175_v126;
          _2175_v126 = _dafny.MultiSet.fromElements((_this).f4);
          let _2176_v127;
          _2176_v127 = _dafny.Seq.of(_2164_cf6);
          let _2177_v128;
          _2177_v128 = _module.D17.create_DC38(_2174_cf56);
          let _2178_v129;
          let _nw330 = Array((new BigNumber(18)).toNumber());
          _nw330[(_dafny.ZERO).toNumber()] = (_1996_v3).f9;
          _nw330[(_dafny.ONE).toNumber()] = _this.f7;
          _nw330[(new BigNumber(2)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(3)).toNumber()] = (_2170_v123)[_module.__default.safeIndex((((_2175_v126).contains((_this).f4)) ? ((_2175_v126).get((_this).f4)) : (p2)), new BigNumber((_2170_v123).length))];
          _nw330[(new BigNumber(4)).toNumber()] = _this.f7;
          _nw330[(new BigNumber(5)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(6)).toNumber()] = new BigNumber((((_1996_v3.f8) ? (_2176_v127) : (_dafny.Seq.update(_2176_v127, _module.__default.safeIndex((_1996_v3).f9, new BigNumber((_2176_v127).length)), _2164_cf6)))).length);
          _nw330[(new BigNumber(7)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(8)).toNumber()] = ((_1996_v3).f9).minus((_1996_v3).f9);
          _nw330[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(((_1996_v3).f9).multipliedBy((_1996_v3).f9));
          _nw330[(new BigNumber(10)).toNumber()] = ((_1996_v3.f8) ? (p2) : (new BigNumber(((_2177_v128).dtor_cf56).length)));
          _nw330[(new BigNumber(11)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(12)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(13)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(14)).toNumber()] = (_1996_v3).f9;
          _nw330[(new BigNumber(15)).toNumber()] = ((!(_1996_v3.f8)) ? (new BigNumber((_2170_v123).length)) : ((_1996_v3).f9));
          _nw330[(new BigNumber(16)).toNumber()] = (_2170_v123)[_module.__default.safeIndex(new BigNumber((_2170_v123).length), new BigNumber((_2170_v123).length))];
          _nw330[(new BigNumber(17)).toNumber()] = _this.f7;
          _2178_v129 = _nw330;
          let _index344 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2178_v129).length));
          (_2178_v129)[_index344] = (p2).plus(p2);
          let _2179_v130;
          _2179_v130 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f4),new BigNumber(((_this).f6).length));
          let _index345 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2178_v129).length));
          (_2178_v129)[_index345] = new BigNumber(((_2179_v130).Merge(_2179_v130)).length);
          let _index346 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_2178_v129).length));
          (_2178_v129)[_index346] = _module.__default.safeModuloInt((_this.f7).plus(p2), (_1996_v3).f9);
          (_1996_v3).f8 = p0;
          let _2180_v131;
          _2180_v131 = _module.D1.create_DC4((_1996_v3).f9);
          let _2181_v132;
          _2181_v132 = _module.D3.create_DC10((_2175_v126).Difference((_2175_v126).update(!((_this).f4), _module.__default.abs((_1996_v3).f9))), function (_pat_let60_0) {
  return function (_2182_dt__update__tmp_h4) {
    return function (_pat_let61_0) {
      return function (_2183_dt__update_hcf7_h0) {
        return _module.D1.create_DC4(_2183_dt__update_hcf7_h0);
      }(_pat_let61_0);
    }(_this.f7);
  }(_pat_let60_0);
}(_2180_v131));
          let _2184_v133;
          _2184_v133 = _dafny.Seq.of((_this).f4, false, _1996_v3.f8);
          let _2185_v134;
          _2185_v134 = _dafny.MultiSet.fromElements(new BigNumber((_2184_v133).length), (_1996_v3).f9);
          _2181_v132 = _module.__default.fm39(_1996_v3.f8, (_1996_v3).f9, new BigNumber(((_dafny.MultiSet.fromElements(_this.f7, new BigNumber((_dafny.Seq.UnicodeFromString("kjili")).length), (_1996_v3).f9)).Intersect(_2185_v134)).cardinality()), _this.f7, globalState);
        } else if (_source26.is_DC39) {
          let _2186___mcc_h12 = (_source26).cf57;
          let _2187___mcc_h13 = (_source26).cf58;
          let _2188___mcc_h14 = (_source26).cf59;
          let _2189___mcc_h15 = (_source26).cf60;
          let _2190___mcc_h16 = (_source26).cf61;
          let _2191_cf61 = _2190___mcc_h16;
          let _2192_cf60 = _2189___mcc_h15;
          let _2193_cf59 = _2188___mcc_h14;
          let _2194_cf58 = _2187___mcc_h13;
          let _2195_cf57 = _2186___mcc_h12;
          (_this).f7 = (_1996_v3).f9;
          _2191_cf61 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_2196_v3) => function (_2197_i12) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,(_2196_v3).f9);
          })(_1996_v3))).length);
          let _index347 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_2164_cf6).length));
          (_2164_cf6)[_index347] = !(_2194_cf58);
          let _index348 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_2164_cf6).length));
          (_2164_cf6)[_index348] = _module.__default.fm18(p2, _2166_v121, _module.__default.safeDivisionInt(_2191_cf61, _this.f7), globalState);
          let _2198_v135;
          let _nw331 = Array((new BigNumber(28)).toNumber()).fill([]);
          _2198_v135 = _nw331;
          let _2199_v136;
          _2199_v136 = _dafny.Map.Empty.slice().updateUnsafe(_2198_v135,_this.f7);
          _2199_v136 = (_2199_v136).update(_2198_v135, (new BigNumber((p1).length)).minus(new BigNumber(((_this).f6).length)));
        } else if (_source26.is_DC37) {
          let _2200___mcc_h17 = (_source26).cf55;
          let _2201_cf55 = _2200___mcc_h17;
          let _2202_v137;
          _2202_v137 = _dafny.Seq.of(_1996_v3.f8, _1996_v3.f8, p0, _1996_v3.f8, false);
          let _2203_v138;
          _2203_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v3.f8,_2166_v121);
          let _2204_v139;
          _2204_v139 = _dafny.Map.Empty.slice().updateUnsafe((_2202_v137)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_2203_v138).update(_1996_v3.f8, _2166_v121)).length)), new BigNumber((_2202_v137).length))],p2);
          _2204_v139 = _2204_v139;
          (_this).f7 = _module.__default.safeDivisionInt((new BigNumber((_dafny.Seq.of(_1996_v3.f8, _1996_v3.f8)).length)).plus(_this.f7), (_1996_v3).f9);
          let _2205_v140;
          let _init68 = function (_2206_i13) {
            return (_2206_i13).plus(_this.f7);
          };
          let _nw332 = Array((new BigNumber(6)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw332.length); _i0_68++) {
            _nw332[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2205_v140 = _nw332;
          let _2207_v141;
          _2207_v141 = _dafny.Seq.of(_2205_v140);
          let _2208_v142;
          let _nw333 = Array((new BigNumber(7)).toNumber());
          _nw333[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2207_v141, _2207_v141);
          _nw333[(_dafny.ONE).toNumber()] = _2207_v141;
          _nw333[(new BigNumber(2)).toNumber()] = _2207_v141;
          _nw333[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2207_v141, _dafny.Seq.update(_2207_v141, _module.__default.safeIndex(p2, new BigNumber((_2207_v141).length)), _2205_v140));
          _nw333[(new BigNumber(4)).toNumber()] = _2207_v141;
          _nw333[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2207_v141, _dafny.Seq.of(_2205_v140, _2205_v140));
          _nw333[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_2205_v140, _2205_v140);
          _2208_v142 = _nw333;
          let _index349 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_2208_v142).length));
          (_2208_v142)[_index349] = _2207_v141;
          let _index350 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_2208_v142).length));
          (_2208_v142)[_index350] = _dafny.Seq.Concat(_2207_v141, _dafny.Seq.Concat(_2207_v141, _2207_v141));
          let _2209_v143;
          let _nw334 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _2209_v143 = _nw334;
          let _2210_v144;
          _2210_v144 = _dafny.Map.Empty.slice().updateUnsafe((_1996_v3).f9,p2);
          let _index351 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2209_v143).length));
          (_2209_v143)[_index351] = _2210_v144;
          let _index352 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2209_v143).length));
          let _rhs287 = _this.f7;
          let _rhs288 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f6).length),new BigNumber(((_this).f6).length))).Merge(_2210_v144);
          let _lhs211 = _this;
          let _lhs212 = _2209_v143;
          let _lhs213 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_2209_v143).length));
          _lhs211.f7 = _rhs287;
          _lhs212[_lhs213] = _rhs288;
        } else {
          let _2211___mcc_h18 = (_source26).cf62;
          let _2212_cf62 = _2211___mcc_h18;
          let _2213_v145;
          let _nw335 = new _module.C3();
          _nw335.__ctor(true, (_this).f4);
          _2213_v145 = _nw335;
          let _index353 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_2164_cf6).length));
          (_2164_cf6)[_index353] = p0;
          let _index354 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_2164_cf6).length));
          (_2164_cf6)[_index354] = (_this).f4;
          (_this).f7 = new BigNumber(282);
          r1 = (((_this).fm0(globalState)).multipliedBy((_1996_v3).f9)).plus(_module.__default.safeModuloInt(new BigNumber(((_this).f6).length), (_1996_v3).f9));
        }
      }
      (globalState).f0 = ((_1996_v3).f9).isLessThan(new BigNumber((_module.__default.fm24(globalState)).cardinality()));
      let _2214_v146;
      _2214_v146 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f7);
      let _2215_v147;
      _2215_v147 = _dafny.Seq.of(false, _1996_v3.f8, _1996_v3.f8, p0);
      let _2216_v148;
      _2216_v148 = _dafny.Set.fromElements(p2);
      let _2217_v149;
      _2217_v149 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_2216_v148).length));
      let _2218_v150;
      let _nw336 = Array((new BigNumber(3)).toNumber());
      _nw336[(_dafny.ZERO).toNumber()] = _1996_v3.f8;
      _nw336[(_dafny.ONE).toNumber()] = !(p0);
      _nw336[(new BigNumber(2)).toNumber()] = p0;
      _2218_v150 = _nw336;
      let _2219_v151;
      _2219_v151 = _module.D1.create_DC3(_2218_v150);
      let _2220_v152;
      _2220_v152 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-315)), function (_2221_i14) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }));
      let _2222_v153;
      _2222_v153 = _dafny.Set.fromElements(_this.f3);
      let _nw337 = Array((new BigNumber(27)).toNumber());
      _nw337[(_dafny.ZERO).toNumber()] = _this.f7;
      _nw337[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_1996_v3).f9);
      _nw337[(new BigNumber(2)).toNumber()] = (((_2214_v146).contains(_dafny.Seq.UnicodeFromString("ochh"))) ? ((_2214_v146).get(_dafny.Seq.UnicodeFromString("ochh"))) : (new BigNumber((_2215_v147).length)));
      _nw337[(new BigNumber(3)).toNumber()] = (_1996_v3).f9;
      _nw337[(new BigNumber(4)).toNumber()] = _this.f7;
      _nw337[(new BigNumber(5)).toNumber()] = (_1996_v3).f9;
      _nw337[(new BigNumber(6)).toNumber()] = (_1996_v3).f9;
      _nw337[(new BigNumber(7)).toNumber()] = new BigNumber((p1).length);
      _nw337[(new BigNumber(8)).toNumber()] = (((_2217_v149).contains(p0)) ? ((_2217_v149).get(p0)) : (p2));
      _nw337[(new BigNumber(9)).toNumber()] = (_1996_v3).f9;
      _nw337[(new BigNumber(10)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_1996_v3).f9,_2218_v150)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,(_2219_v151).dtor_cf6))).length);
      _nw337[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(p2, _this.f7);
      _nw337[(new BigNumber(12)).toNumber()] = new BigNumber(((_this).f6).length);
      _nw337[(new BigNumber(13)).toNumber()] = new BigNumber((_2220_v152).length);
      _nw337[(new BigNumber(14)).toNumber()] = p2;
      _nw337[(new BigNumber(15)).toNumber()] = _this.f7;
      _nw337[(new BigNumber(16)).toNumber()] = (new BigNumber((_1997_v4).length)).plus(_this.f7);
      _nw337[(new BigNumber(17)).toNumber()] = _this.f7;
      _nw337[(new BigNumber(18)).toNumber()] = _this.f7;
      _nw337[(new BigNumber(19)).toNumber()] = ((p0) ? (p2) : (_this.f7));
      _nw337[(new BigNumber(20)).toNumber()] = (_1996_v3).f9;
      _nw337[(new BigNumber(21)).toNumber()] = (p2).minus(p2);
      _nw337[(new BigNumber(22)).toNumber()] = p2;
      _nw337[(new BigNumber(23)).toNumber()] = new BigNumber(714);
      _nw337[(new BigNumber(24)).toNumber()] = _module.__default.safeModuloInt((_1996_v3).f9, (((_2217_v149).contains(p0)) ? ((_2217_v149).get(p0)) : (_this.f7)));
      _nw337[(new BigNumber(25)).toNumber()] = p2;
      _nw337[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_1996_v3).f9, p2, new BigNumber((_2222_v153).length))).length);
      r0 = _nw337;
      let _2223_v154;
      _2223_v154 = _dafny.MultiSet.fromElements(!(!(_1996_v3.f8) || ((_this).f4)));
      let _2224_v155;
      _2224_v155 = new _dafny.CodePoint('m'.codePointAt(0));
      let _2225_v156;
      let _nw338 = new _module.C2();
      _nw338.__ctor(_this.f3, p0, (_this).f4, _this.f3);
      _2225_v156 = _nw338;
      let _2226_v157;
      _2226_v157 = _dafny.MultiSet.fromElements(_2225_v156);
      let _2227_v158;
      _2227_v158 = _dafny.Seq.of(new BigNumber((_2226_v157).cardinality()));
      r1 = (((_2223_v154).contains(_module.__default.fm18((_dafny.ZERO).minus(_this.f7), _2224_v155, p2, globalState))) ? ((_2223_v154).get(_module.__default.fm18((_dafny.ZERO).minus(_this.f7), _2224_v155, p2, globalState))) : (new BigNumber((_2227_v158).length)));
      return [r0, r1];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2228_i0;
      _2228_i0 = _dafny.ZERO;
      L8: {
        while (!(p1.f3) || (!(!(false)) || (_module.__default.fm4(p1.f3, globalState)))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2228_i0)) {
              break L8;
            }
            _2228_i0 = (_2228_i0).plus(_dafny.ONE);
            let _2229_v0;
            let _nw339 = new _module.C1();
            _nw339.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(497)), function (_2230_i1) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            }), p1.f3);
            _2229_v0 = _nw339;
            let _2231_v1;
            _2231_v1 = _dafny.Map.Empty.slice().updateUnsafe((_2229_v0).f13,(p1).f4);
            let _2232_v2;
            _2232_v2 = _dafny.Seq.of(p0, (p0).plus(_this.f7), (_this).fm0(globalState), p0, _this.f7);
            let _2233_v3;
            _2233_v3 = _dafny.Seq.UnicodeFromString("fscgsd");
            let _rhs289 = _2231_v1;
            let _rhs290 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p0, _this.f7), _2232_v2), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, _this.f7), _2232_v2)).length)), _module.__default.safeDivisionInt(p0, p0));
            let _rhs291 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("psukcoa"), _2233_v3);
            _2231_v1 = _rhs289;
            _2232_v2 = _rhs290;
            _2233_v3 = _rhs291;
            (p1).f3 = true;
            let _2234_v4;
            _2234_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(494),(p1).f4);
            let _2235_v5;
            _2235_v5 = _dafny.MultiSet.fromElements(p1.f3, (_this).f4, p1.f3, (p1).f4, (p1).f4);
            let _2236_v6;
            _2236_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2235_v5,new BigNumber((_2233_v3).length));
            let _2237_v7;
            let _nw340 = new _module.C4();
            _nw340.__ctor((((_2234_v4).contains(p0)) ? ((_2234_v4).get(p0)) : (p1.f3)), (((_2236_v6).contains(_2235_v5)) ? ((_2236_v6).get(_2235_v5)) : (_this.f7)), p1.f3, !((_2229_v0).f13));
            _2237_v7 = _nw340;
          }
        }
      }
      (_this).f3 = false;
      let _2238_v8;
      let _init69 = ((_2239_p0) => function (_2240_i2) {
        return _module.D14.create_DC30(_dafny.MultiSet.fromElements(_2239_p0, _2239_p0, new BigNumber(349), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),false)).length)));
      })(p0);
      let _nw341 = Array((new BigNumber(7)).toNumber());
      for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw341.length); _i0_69++) {
        _nw341[_i0_69] = _init69(new BigNumber(_i0_69));
      }
      _2238_v8 = _nw341;
      _2238_v8 = _2238_v8;
      let _2241_v9;
      let _nw342 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2241_v9 = _nw342;
      let _2242_v10;
      let _nw343 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _2242_v10 = _nw343;
      let _index355 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
      (_2242_v10)[_index355] = new BigNumber(665);
      let _2243_v11;
      _2243_v11 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,p0);
      let _2244_v12;
      _2244_v12 = _module.D16.create_DC35((((p1).f4) ? (_2243_v11) : (_2243_v11)));
      let _2245_v13;
      let _nw344 = Array((new BigNumber(2)).toNumber());
      _nw344[(_dafny.ZERO).toNumber()] = (_this).f6;
      _nw344[(_dafny.ONE).toNumber()] = (_this).f6;
      _2245_v13 = _nw344;
      let _index356 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length));
      (_2245_v13)[_index356] = (_this).f6;
      let _2246_v14;
      _2246_v14 = _module.D21.create_DC51(_2241_v9);
      let _2247_v15;
      _2247_v15 = _dafny.MultiSet.fromElements(_this.f7, _this.f7);
      let _2248_v16;
      _2248_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1.f3);
      let _2249_v17;
      _2249_v17 = _dafny.Seq.of(p0);
      let _2250_v18;
      let _nw345 = Array((new BigNumber(28)).toNumber());
      _nw345[(_dafny.ZERO).toNumber()] = !(_this.f3);
      _nw345[(_dafny.ONE).toNumber()] = true;
      _nw345[(new BigNumber(2)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(3)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(4)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(5)).toNumber()] = p1.f3;
      _nw345[(new BigNumber(6)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(7)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(8)).toNumber()] = p1.f3;
      _nw345[(new BigNumber(9)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(10)).toNumber()] = p1.f3;
      _nw345[(new BigNumber(11)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(12)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(13)).toNumber()] = p1.f3;
      _nw345[(new BigNumber(14)).toNumber()] = !(_this.f3);
      _nw345[(new BigNumber(15)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(16)).toNumber()] = !(p1.f3);
      _nw345[(new BigNumber(17)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(18)).toNumber()] = (p1).f4;
      _nw345[(new BigNumber(19)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(20)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(21)).toNumber()] = false;
      _nw345[(new BigNumber(22)).toNumber()] = _this.f3;
      _nw345[(new BigNumber(23)).toNumber()] = !(!(p1.f3));
      _nw345[(new BigNumber(24)).toNumber()] = (((_2248_v16).contains(new BigNumber((_2249_v17).length))) ? ((_2248_v16).get(new BigNumber((_2249_v17).length))) : (true));
      _nw345[(new BigNumber(25)).toNumber()] = false;
      _nw345[(new BigNumber(26)).toNumber()] = (_this).f4;
      _nw345[(new BigNumber(27)).toNumber()] = (_this).f4;
      _2250_v18 = _nw345;
      let _2251_v19;
      _2251_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14(globalState),_2250_v18);
      let _2252_v20;
      _2252_v20 = _module.D19.create_DC45((_dafny.ZERO).minus(_this.f7), new BigNumber((_2251_v19).length), (_this).f4, p1.f3, _dafny.Seq.of(p1.f3));
      let _index357 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
      let _index358 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length));
      let _rhs292 = (_2246_v14).dtor_cf83;
      let _rhs293 = (new BigNumber(-408)).isLessThanOrEqualTo((((_2247_v15).contains((_dafny.ZERO).minus(_this.f7))) ? ((_2247_v15).get((_dafny.ZERO).minus(_this.f7))) : (_this.f7)));
      let _rhs294 = (_2252_v20).dtor_cf70;
      let _rhs295 = _2244_v12;
      let _rhs296 = (_this).f6;
      let _lhs214 = p1;
      let _lhs215 = _2242_v10;
      let _lhs216 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
      let _lhs217 = _2245_v13;
      let _lhs218 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length));
      _2241_v9 = _rhs292;
      _lhs214.f3 = _rhs293;
      _lhs215[_lhs216] = _rhs294;
      _2244_v12 = _rhs295;
      _lhs217[_lhs218] = _rhs296;
      let _source27 = _2244_v12;
      if (_source27.is_DC36) {
        let _2253___mcc_h0 = (_source27).cf51;
        let _2254___mcc_h1 = (_source27).cf52;
        let _2255___mcc_h2 = (_source27).cf53;
        let _2256___mcc_h3 = (_source27).cf54;
        let _2257_cf54 = _2256___mcc_h3;
        let _2258_cf53 = _2255___mcc_h2;
        let _2259_cf52 = _2254___mcc_h1;
        let _2260_cf51 = _2253___mcc_h0;
        let _2261_v21;
        let _nw346 = new _module.C2();
        _nw346.__ctor(false, _2258_cf53, ((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]).isLessThanOrEqualTo(_this.f7), p1.f3);
        _2261_v21 = _nw346;
        (_this).f7 = p0;
        let _2262_v22;
        let _nw347 = new _module.C0();
        _nw347.__ctor();
        _2262_v22 = _nw347;
        if ((_2247_v15).equals((_2247_v15).Intersect(_2247_v15))) {
          let _rhs297 = (_2261_v21).f11;
          let _rhs298 = function () {
            let _coll72 = new _dafny.Set();
            for (const _compr_72 of (function () {
              let _coll73 = new _dafny.Set();
              for (const _compr_73 of (_module.__default.fm19(globalState)).Elements) {
                let _2263_v23 = _compr_73;
                if (_dafny.Seq.contains(_module.__default.fm19(globalState), _2263_v23)) {
                  _coll73.add((_2263_v23).multipliedBy(new BigNumber(-208)));
                }
              }
              return _coll73;
            }()).Elements) {
              let _2264_v24 = _compr_72;
              if ((function () {
                let _coll74 = new _dafny.Set();
                for (const _compr_74 of (_module.__default.fm19(globalState)).Elements) {
                  let _2265_v23 = _compr_74;
                  if (_dafny.Seq.contains(_module.__default.fm19(globalState), _2265_v23)) {
                    _coll74.add((_2265_v23).multipliedBy(new BigNumber(-208)));
                  }
                }
                return _coll74;
              }()).contains(_2264_v24)) {
                _coll72.add(_module.__default.safeModuloInt(_2264_v24, ((!(false)) ? (new BigNumber(678)) : (new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(!(false)))).length), new BigNumber(824), new BigNumber((_dafny.Seq.UnicodeFromString("hyhfjwks")).length), new BigNumber((_dafny.Set.fromElements(!(true))).length)), _dafny.Seq.of(new BigNumber(719), new BigNumber((function () {
                  let _coll75 = new _dafny.Map();
                  for (const _compr_75 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("oxtq")).length), new BigNumber((_dafny.Seq.UnicodeFromString("cb")).length))).length), (_dafny.ZERO).minus(new BigNumber(-986)))).Elements) {
                    let _2266_v25 = _compr_75;
                    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("oxtq")).length), new BigNumber((_dafny.Seq.UnicodeFromString("cb")).length))).length), (_dafny.ZERO).minus(new BigNumber(-986))), _2266_v25)) {
                      _coll75.push([_module.__default.safeModuloInt(_2266_v25, new BigNumber(178)),false]);
                    }
                  }
                  return _coll75;
                }()).length)))).length)))));
              }
            }
            return _coll72;
          }();
          let _rhs299 = ((!((_this).f4)) ? ((_2261_v21).f10) : ((_2261_v21).f10));
          let _rhs300 = ((_dafny.ZERO).minus((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))])).isLessThanOrEqualTo(((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]).plus(new BigNumber((_2247_v15).cardinality())));
          let _lhs219 = p1;
          let _lhs220 = p1;
          _lhs219.f3 = _rhs297;
          _2260_cf51 = _rhs298;
          _2258_cf53 = _rhs299;
          _lhs220.f3 = _rhs300;
          let _2267_v26;
          _2267_v26 = _module.D1.create_DC5(p1.f3, _2258_cf53, (_this).fm3(globalState), p0);
          let _2268_v27;
          _2268_v27 = _dafny.Map.Empty.slice().updateUnsafe((_2267_v26).dtor_cf11,(_module.D1.create_DC3(_2250_v18)).dtor_cf6);
          _2268_v27 = (_2268_v27).update((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))], _2250_v18);
          (_this).f7 = (p0).multipliedBy((_2249_v17)[_module.__default.safeIndex(new BigNumber((_2260_cf51).length), new BigNumber((_2249_v17).length))]);
          (_2261_v21).m9(globalState);
          let _index359 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2250_v18).length));
          (_2250_v18)[_index359] = p1.f3;
          let _2269_v29;
          _2269_v29 = _dafny.Seq.of(_2250_v18);
          let _2270_v30;
          _2270_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2269_v29).length),p0);
          let _2271_v31;
          _2271_v31 = _dafny.Map.Empty.slice().updateUnsafe((function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of _dafny.IntegerRange(new BigNumber(952), new BigNumber(-100))) {
              let _2272_v28 = _compr_76;
              if (((new BigNumber(952)).isLessThanOrEqualTo(_2272_v28)) && ((_2272_v28).isLessThan(new BigNumber(-100)))) {
                _coll76.push([(_2272_v28).multipliedBy(new BigNumber(((_2245_v13)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length))]).length)),new BigNumber(105)]);
              }
            }
            return _coll76;
          }()).Merge(_2270_v30),_2257_cf54);
          let _index360 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
          let _index361 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2250_v18).length));
          let _rhs301 = p0;
          let _rhs302 = _2245_v13;
          let _rhs303 = ((_this.f7).isLessThan(_this.f7)) && (_2259_cf52);
          let _rhs304 = (new BigNumber(983)).multipliedBy((((_2261_v21).f10) ? (new BigNumber(836)) : (p0)));
          let _rhs305 = _2271_v31;
          let _lhs221 = _2242_v10;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
          let _lhs223 = _2250_v18;
          let _lhs224 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2250_v18).length));
          _lhs221[_lhs222] = _rhs301;
          _2245_v13 = _rhs302;
          _lhs223[_lhs224] = _rhs303;
          r0 = _rhs304;
          _2271_v31 = _rhs305;
        } else {
          let _2273_v32;
          _2273_v32 = _module.D1.create_DC5(p1.f3, p1.f3, (_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))], (_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]);
          let _2274_v33;
          _2274_v33 = _dafny.Seq.of(_2242_v10, _2242_v10, _2242_v10);
          let _2275_v34;
          _2275_v34 = _module.D19.create_DC46(p0, _2259_cf52, p1.f3, _2273_v32, new BigNumber((_2274_v33).length));
          let _2276_v35;
          _2276_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2258_cf53,(_2261_v21).f11);
          let _2277_v36;
          _2277_v36 = _dafny.Set.fromElements(false);
          _2275_v34 = _module.D19.create_DC46((new BigNumber(735)).minus((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]), (p1).f4, (_2277_v36).IsProperSubsetOf(_dafny.Set.fromElements((((_2276_v35).contains(_2258_cf53)) ? ((_2276_v35).get(_2258_cf53)) : ((((_2248_v16).contains((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))])) ? ((_2248_v16).get((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))])) : (false)))))), _module.D1.create_DC5(p1.f3, (p1).f4, _this.f7, (_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]), (_2262_v22).fm10((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))], _2257_cf54, _2247_v15, globalState));
          let _2278_v37;
          let _nw348 = Array((new BigNumber(19)).toNumber());
          _2278_v37 = _nw348;
          let _2279_v38;
          _2279_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_2278_v37);
          let _2280_v39;
          _2280_v39 = _module.D19.create_DC44(_2279_v38);
          let _2281_v40;
          _2281_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2280_v39,_module.__default.safeDivisionInt(_this.f7, new BigNumber((_2243_v11).length)));
          _2281_v40 = (_2281_v40).update(_2280_v39, new BigNumber((_2247_v15).cardinality()));
          (_this).f7 = _module.__default.safeModuloInt(p0, p0);
          _2276_v35 = (_2276_v35).update(!(p1.f3), (_2261_v21).f11);
          let _2282_v41;
          _2282_v41 = _dafny.Seq.of(_2258_cf53);
          let _2283_v42;
          _2283_v42 = _module.D12.create_DC26(_2277_v36);
          let _2284_v43;
          _2284_v43 = _module.D12.create_DC28(_2283_v42);
          let _2285_v44;
          _2285_v44 = _module.D0.create_DC0((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]);
          let _2286_v45;
          _2286_v45 = _dafny.Set.fromElements(_2284_v43);
          let _pat_let_tv66 = _2283_v42;
          let _rhs306 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p1.f3), _2282_v41)).length);
          let _rhs307 = _2242_v10;
          let _rhs308 = (_dafny.Set.fromElements(_2284_v43, _module.__default.fm38((_2285_v44).dtor_cf0, _this.f3, globalState), _2284_v43, function (_pat_let62_0) {
            return function (_2287_dt__update__tmp_h0) {
              return function (_pat_let63_0) {
                return function (_2288_dt__update_hcf41_h0) {
                  return _module.D12.create_DC28(_2288_dt__update_hcf41_h0);
                }(_pat_let63_0);
              }(_pat_let_tv66);
            }(_pat_let62_0);
          }(_2284_v43), _2284_v43)).IsSubsetOf(_2286_v45);
          let _lhs225 = _this;
          _lhs225.f7 = _rhs306;
          _2242_v10 = _rhs307;
          _2259_cf52 = _rhs308;
        }
      } else {
        let _2289___mcc_h4 = (_source27).cf50;
        let _2290_cf50 = _2289___mcc_h4;
        let _2291_v46;
        _2291_v46 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2292_v47;
        _2292_v47 = _dafny.Map.Empty.slice().updateUnsafe(true,(((_2248_v16).contains((_this).fm3(globalState))) ? ((_2248_v16).get((_this).fm3(globalState))) : ((p1).f4)));
        let _2293_v48;
        _2293_v48 = _module.D8.create_DC20(_2291_v46, (((_2292_v47).contains(false)) ? ((_2292_v47).get(false)) : (p1.f3)));
        let _index362 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
        (_2250_v18)[_index362] = (p1).f4;
        let _index363 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
        let _rhs309 = _2293_v48;
        let _rhs310 = (_this).f4;
        let _lhs226 = _2250_v18;
        let _lhs227 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
        _2293_v48 = _rhs309;
        _lhs226[_lhs227] = _rhs310;
        (p1).f3 = !(_this.f3);
        let _2294_v49;
        _2294_v49 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(355)), ((_2295_v46) => function (_2296_i3) {
          return _2295_v46;
        })(_2291_v46)), (_2245_v13)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length))]), (_this).f6, _dafny.Seq.UnicodeFromString("hxq"));
        _2294_v49 = _2294_v49;
        let _2297_v50;
        _2297_v50 = _dafny.Map.Empty.slice().updateUnsafe(_2247_v15,!(_this.f3));
        let _2298_v51;
        _2298_v51 = _module.D17.create_DC39(_2242_v10, (p1).f4, (((_2250_v18)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length))]) ? ((p1).f4) : ((_this).f4)), (new BigNumber(((_2297_v50).update(_2247_v15, _this.f3)).length)).plus(p0), _this.f7);
        let _source28 = _2298_v51;
        if (_source28.is_DC38) {
          let _2299___mcc_h5 = (_source28).cf56;
          let _2300_cf56 = _2299___mcc_h5;
          let _index364 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
          (_2242_v10)[_index364] = p0;
          let _index365 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          (_2250_v18)[_index365] = !(_module.__default.fm14(globalState));
          let _2301_v52;
          _2301_v52 = _dafny.MultiSet.fromElements((p1).f4, (p1).f4);
          let _2302_v53;
          _2302_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2301_v52);
          let _2303_v54;
          _2303_v54 = _dafny.Set.fromElements(false, false, ((((_2302_v53).contains(_2300_cf56)) ? ((_2302_v53).get(_2300_cf56)) : (_module.__default.fm20(p1.f3, globalState)))).IsProperSubsetOf(_2301_v52), ((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]).isEqualTo(_this.f7), (_this).f4);
          let _2304_v55;
          _2304_v55 = _dafny.Seq.of(p1.f3);
          let _pat_let_tv67 = _2303_v54;
          _2303_v54 = ((function (_pat_let64_0) {
            return function (_2305_dt__update__tmp_h1) {
              return function (_pat_let65_0) {
                return function (_2306_dt__update_hcf39_h0) {
                  return _module.D12.create_DC26(_2306_dt__update_hcf39_h0);
                }(_pat_let65_0);
              }(_pat_let_tv67);
            }(_pat_let64_0);
          }(_module.D12.create_DC26(_dafny.Set.fromElements((_this).f4)))).dtor_cf39).Difference((_2303_v54).Intersect(_module.__default.fm33((((_2247_v15).contains(_this.f7)) ? ((_2247_v15).get(_this.f7)) : (new BigNumber(((_2245_v13)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length))]).length))), _2304_v55, globalState)));
          (_this).f3 = p1.f3;
        } else if (_source28.is_DC39) {
          let _2307___mcc_h6 = (_source28).cf57;
          let _2308___mcc_h7 = (_source28).cf58;
          let _2309___mcc_h8 = (_source28).cf59;
          let _2310___mcc_h9 = (_source28).cf60;
          let _2311___mcc_h10 = (_source28).cf61;
          let _2312_cf61 = _2311___mcc_h10;
          let _2313_cf60 = _2310___mcc_h9;
          let _2314_cf59 = _2309___mcc_h8;
          let _2315_cf58 = _2308___mcc_h7;
          let _2316_cf57 = _2307___mcc_h6;
          let _2317_v56;
          let _nw349 = Array((new BigNumber(10)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2317_v56 = _nw349;
          let _2318_v57;
          _2318_v57 = _dafny.Seq.of(_2315_cf58, _2314_cf59);
          let _index366 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_2317_v56).length));
          (_2317_v56)[_index366] = _module.__default.fm20(_module.__default.fm6(_module.__default.fm31(p1.f3, p0, _2315_cf58, _2318_v57, globalState), globalState), globalState);
          let _2319_v58;
          _2319_v58 = _dafny.MultiSet.fromElements(!((_2250_v18)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length))]));
          let _index367 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_2317_v56).length));
          (_2317_v56)[_index367] = _2319_v58;
          (globalState).f0 = (p1).f4;
          let _2320_v59;
          let _nw350 = new _module.C1();
          _nw350.__ctor((_this).f6, p1.f3);
          _2320_v59 = _nw350;
          let _2321_v60;
          _2321_v60 = _dafny.Seq.of(_2320_v59);
          let _2322_v61;
          let _init70 = function (_2323_i4) {
            return true;
          };
          let _nw351 = Array((new BigNumber(7)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw351.length); _i0_70++) {
            _nw351[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2322_v61 = _nw351;
          let _2324_v62;
          let _nw352 = Array((new BigNumber(12)).toNumber());
          _nw352[(_dafny.ZERO).toNumber()] = _2320_v59;
          _nw352[(_dafny.ONE).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(2)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(3)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(4)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(5)).toNumber()] = (_2321_v60)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_2321_v60).length))];
          _nw352[(new BigNumber(6)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(7)).toNumber()] = (_module.D15.create_DC33(_2322_v61, _2320_v59)).dtor_cf48;
          _nw352[(new BigNumber(8)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(9)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(10)).toNumber()] = _2320_v59;
          _nw352[(new BigNumber(11)).toNumber()] = _2320_v59;
          _2324_v62 = _nw352;
          let _2325_v63;
          _2325_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2320_v59).f12,_2324_v62);
          let _2326_v64;
          _2326_v64 = _dafny.Map.Empty.slice().updateUnsafe((_2317_v56)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_2317_v56).length))],_2324_v62);
          let _2327_v65;
          let _nw353 = Array((new BigNumber(12)).toNumber());
          _nw353[(_dafny.ZERO).toNumber()] = _2324_v62;
          _nw353[(_dafny.ONE).toNumber()] = (((_2325_v63).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_2330_v46) => function (_2331_i5) {
            return _2330_v46;
          })(_2291_v46)))) ? ((_2325_v63).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_2328_v46) => function (_2329_i5) {
            return _2328_v46;
          })(_2291_v46)))) : (_2324_v62));
          _nw353[(new BigNumber(2)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(3)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(4)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(5)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(6)).toNumber()] = (((_this).f4) ? (_2324_v62) : (_2324_v62));
          _nw353[(new BigNumber(7)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(8)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(9)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(10)).toNumber()] = _2324_v62;
          _nw353[(new BigNumber(11)).toNumber()] = (((_2326_v64).contains(_2319_v58)) ? ((_2326_v64).get(_2319_v58)) : (_2324_v62));
          _2327_v65 = _nw353;
          let _2332_v66;
          _2332_v66 = _dafny.Seq.of(_2324_v62, _2324_v62, _2324_v62, _2324_v62, _2324_v62);
          let _index368 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2327_v65).length));
          (_2327_v65)[_index368] = (_2332_v66)[_module.__default.safeIndex(_2313_cf60, new BigNumber((_2332_v66).length))];
          let _2333_v67;
          _2333_v67 = _dafny.Set.fromElements(_2313_cf60);
          let _2334_v68;
          _2334_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2291_v46);
          let _2335_v70;
          _2335_v70 = _dafny.Seq.of(_2334_v68, _2334_v68, function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of (_2249_v17).Elements) {
              let _2336_v69 = _compr_77;
              if (_dafny.Seq.contains(_2249_v17, _2336_v69)) {
                _coll77.push([_module.__default.safeDivisionInt(_2336_v69, _2312_cf61),_2291_v46]);
              }
            }
            return _coll77;
          }());
          let _2337_v71;
          _2337_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2333_v67,((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_2335_v70, _module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber((_2335_v70).length)), _dafny.Map.Empty.slice().updateUnsafe((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))],_2291_v46))).length))));
          let _2338_v72;
          _2338_v72 = _module.D1.create_DC4((_dafny.ZERO).minus(_2312_cf61));
          let _2339_v73;
          _2339_v73 = _module.D1.create_DC5(_2314_cf59, (p1).f4, (_2338_v72).dtor_cf7, new BigNumber(-848));
          let _index369 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2327_v65).length));
          let _index370 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          let _rhs311 = _2242_v10;
          let _rhs312 = _2324_v62;
          let _rhs313 = ((_2337_v71).Merge(_2337_v71)).Merge(_2337_v71);
          let _rhs314 = (_2339_v73).dtor_cf9;
          let _lhs228 = _2327_v65;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2327_v65).length));
          let _lhs230 = _2250_v18;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          _2316_cf57 = _rhs311;
          _lhs228[_lhs229] = _rhs312;
          _2337_v71 = _rhs313;
          _lhs230[_lhs231] = _rhs314;
          let _index371 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          (_2250_v18)[_index371] = p1.f3;
        } else if (_source28.is_DC37) {
          let _2340___mcc_h11 = (_source28).cf55;
          let _2341_cf55 = _2340___mcc_h11;
          let _2342_v74;
          _2342_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2242_v10,true);
          let _index372 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          let _rhs315 = !(!((_2342_v74).contains(_2242_v10)));
          let _rhs316 = !((p0).isEqualTo(_module.__default.safeModuloInt(p0, p0)));
          let _lhs232 = _2250_v18;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          let _lhs234 = p1;
          _lhs232[_lhs233] = _rhs315;
          _lhs234.f3 = _rhs316;
          let _2343_v75;
          let _nw354 = new _module.C1();
          _nw354.__ctor(_dafny.Seq.Concat((_this).f6, (_2245_v13)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2245_v13).length))]), p1.f3);
          _2343_v75 = _nw354;
          let _index373 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_2250_v18).length));
          (_2250_v18)[_index373] = _this.f3;
          r0 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), ((_2344_p0) => function (_2345_i6) {
            return _2344_p0;
          })(p0))).length)).plus((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))]);
        } else {
          let _2346___mcc_h12 = (_source28).cf62;
          let _2347_cf62 = _2346___mcc_h12;
          let _index374 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length));
          (_2242_v10)[_index374] = (((_2293_v48).dtor_cf34) ? (new BigNumber(-692)) : (new BigNumber(973)));
          let _2348_v76;
          _2348_v76 = _dafny.Set.fromElements((p1).f4);
          let _2349_v77;
          _2349_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2348_v76,_this.f7);
          _2349_v77 = (_2349_v77).update(_2348_v76, _this.f7);
          let _index375 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2241_v9).length));
          (_2241_v9)[_index375] = _2291_v46;
          let _index376 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2241_v9).length));
          (_2241_v9)[_index376] = new _dafny.CodePoint('e'.codePointAt(0));
          r0 = (_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))];
        }
      }
      let _2350_v78;
      _2350_v78 = _dafny.Seq.of((_this).f4);
      let _hi13 = new BigNumber((_dafny.Seq.UnicodeFromString("pjx")).length);
      for (let _2351_i7 = new BigNumber((_2350_v78).length); _2351_i7.isLessThan(_hi13); _2351_i7 = _2351_i7.plus(_dafny.ONE)) {
        (_this).f7 = p0;
        (_this).f7 = (((p1).f4) ? (p0) : ((p0).minus((_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))])));
        let _2352_v79;
        let _nw355 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
        _2352_v79 = _nw355;
        let _index377 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_2352_v79).length));
        (_2352_v79)[_index377] = _2350_v78;
        let _index378 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_2352_v79).length));
        (_2352_v79)[_index378] = _2350_v78;
        (_this).f7 = new BigNumber(-240);
      }
      r0 = (_2242_v10)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_2242_v10).length))];
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _2353_v0;
      let _init71 = ((_2354_p1) => function (_2355_i0) {
        return !(_2354_p1.f3);
      })(p1);
      let _nw356 = Array((new BigNumber(24)).toNumber());
      for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw356.length); _i0_71++) {
        _nw356[_i0_71] = _init71(new BigNumber(_i0_71));
      }
      _2353_v0 = _nw356;
      let _2356_v1;
      _2356_v1 = _module.D1.create_DC3(_2353_v0);
      let _source29 = _2356_v1;
      if (_source29.is_DC4) {
        let _2357___mcc_h0 = (_source29).cf7;
        let _2358_cf7 = _2357___mcc_h0;
        let _2359_v2;
        _2359_v2 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2360_v3;
        _2360_v3 = _dafny.Set.fromElements(_this.f7, _this.f7);
        let _2361_v4;
        _2361_v4 = _dafny.Set.fromElements(new BigNumber((_2360_v3).length), _2358_cf7);
        let _2362_v5;
        _2362_v5 = _dafny.Set.fromElements(_2361_v4, _2361_v4, _2361_v4);
        let _2363_v6;
        _2363_v6 = _dafny.Map.Empty.slice().updateUnsafe((p1).f4,false);
        let _rhs317 = (_2362_v5).IsDisjointFrom(_2362_v5);
        let _rhs318 = new BigNumber((((_2363_v6).Merge(_2363_v6)).Merge(_2363_v6)).length);
        let _rhs319 = _2358_cf7;
        let _rhs320 = p0;
        let _rhs321 = _this.f3;
        let _lhs235 = p1;
        let _lhs236 = _this;
        let _lhs237 = _this;
        let _lhs238 = _this;
        _lhs235.f3 = _rhs317;
        _lhs236.f7 = _rhs318;
        _lhs237.f7 = _rhs319;
        _2359_v2 = _rhs320;
        _lhs238.f3 = _rhs321;
        let _2364_v7;
        _2364_v7 = _dafny.Seq.of(_this.f3, p1.f3);
        let _2365_v8;
        _2365_v8 = _dafny.Seq.of(_this.f7, _this.f7, _this.f7);
        let _2366_v9;
        let _nw357 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _2366_v9 = _nw357;
        let _2367_v10;
        _2367_v10 = _module.D17.create_DC39(_2366_v9, _this.f3, (p1).f4, _2358_cf7, _2358_cf7);
        let _2368_v11;
        _2368_v11 = _dafny.MultiSet.fromElements(_this.f7, _2358_cf7, (_2367_v10).dtor_cf61, _this.f7, new BigNumber((_dafny.MultiSet.fromElements(p0, p0)).cardinality()));
        let _rhs322 = (p1).f4;
        let _rhs323 = (_2363_v6).Merge((_2363_v6).Merge(_2363_v6));
        let _rhs324 = _module.__default.fm32(_this.f3, true, (_2364_v7)[_module.__default.safeIndex(_2358_cf7, new BigNumber((_2364_v7).length))], !((_dafny.MultiSet.FromArray(_2365_v8)).IsDisjointFrom(_2368_v11)), globalState);
        let _lhs239 = p1;
        _lhs239.f3 = _rhs322;
        _2363_v6 = _rhs323;
        _2360_v3 = _rhs324;
        let _2369_v12;
        let _nw358 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
        _2369_v12 = _nw358;
        let _2370_v13;
        _2370_v13 = _module.D15.create_DC32(_2369_v12);
        let _2371_v14;
        _2371_v14 = _module.D15.create_DC34(_module.D15.create_DC34(_2370_v13));
        let _2372_v15;
        _2372_v15 = _module.D15.create_DC34(_2371_v14);
        let _source30 = _2372_v15;
        if (_source30.is_DC33) {
          let _2373___mcc_h6 = (_source30).cf47;
          let _2374___mcc_h7 = (_source30).cf48;
          let _2375_cf48 = _2374___mcc_h7;
          let _2376_cf47 = _2373___mcc_h6;
          let _2377_v16;
          _2377_v16 = _module.D20.create_DC48(_2365_v8);
          let _2378_v17;
          _2378_v17 = _dafny.Map.Empty.slice().updateUnsafe((_2377_v16).dtor_cf80,_this.f7);
          let _rhs325 = new BigNumber(((((((p1).f4) ? (!((_2375_cf48).f13)) : (true))) ? (_2364_v7) : (_2364_v7))).length);
          let _rhs326 = _2359_v2;
          let _rhs327 = (_this).f4;
          let _rhs328 = ((p1.f3) ? ((_2378_v17).Merge(_2378_v17)) : (_2378_v17));
          let _lhs240 = _this;
          let _lhs241 = p1;
          _lhs240.f7 = _rhs325;
          _2359_v2 = _rhs326;
          _lhs241.f3 = _rhs327;
          _2378_v17 = _rhs328;
          let _init72 = function (_2379_i1) {
            return (_2379_i1).minus(_this.f7);
          };
          let _nw359 = Array((new BigNumber(15)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw359.length); _i0_72++) {
            _nw359[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2366_v9 = _nw359;
          let _2380_v18;
          _2380_v18 = _module.D17.create_DC38(_dafny.Seq.Concat((_this).f6, (_this).f6));
          let _rhs329 = _2380_v18;
          let _rhs330 = _2363_v6;
          let _rhs331 = ((!((_2360_v3).IsDisjointFrom(_2360_v3))) ? (_2358_cf7) : (_2358_cf7));
          _2380_v18 = _rhs329;
          _2363_v6 = _rhs330;
          _2358_cf7 = _rhs331;
          (_this).f7 = _this.f7;
        } else if (_source30.is_DC32) {
          let _2381___mcc_h8 = (_source30).cf46;
          let _2382_cf46 = _2381___mcc_h8;
          r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(186)), ((_2383_p1, _2384_v2) => function (_2385_i2) {
            return (_module.D0.create_DC1((_2383_p1).f4, _2384_v2)).dtor_cf2;
          })(p1, _2359_v2));
          (_this).f3 = (p1).f4;
          (p1).f3 = (p1).f4;
          let _2386_v19;
          _2386_v19 = _dafny.MultiSet.fromElements(p1.f3);
          let _2387_v20;
          _2387_v20 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_2365_v8), _2368_v11);
          let _2388_v21;
          _2388_v21 = _dafny.Seq.of(_dafny.Seq.update(_2365_v8, _module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber((_2365_v8).length)), _this.f7));
          let _2389_v22;
          let _nw360 = Array((new BigNumber(26)).toNumber());
          _nw360[(_dafny.ZERO).toNumber()] = _2368_v11;
          _nw360[(_dafny.ONE).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(2)).toNumber()] = ((_2368_v11).update(new BigNumber(-532), _module.__default.abs(_2358_cf7))).Union(_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).fm3(globalState))));
          _nw360[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_2365_v8);
          _nw360[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_2365_v8);
          _nw360[(new BigNumber(5)).toNumber()] = (_2368_v11).Intersect(_2368_v11);
          _nw360[(new BigNumber(6)).toNumber()] = ((_2368_v11).update(_this.f7, _module.__default.abs((((_2386_v19).contains(false)) ? ((_2386_v19).get(false)) : (_this.f7))))).Intersect((_2387_v20)[_module.__default.safeIndex(_2358_cf7, new BigNumber((_2387_v20).length))]);
          _nw360[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(400), _this.f7)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_2358_cf7, _2358_cf7, new BigNumber((_2386_v19).cardinality()))));
          _nw360[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(_2358_cf7, _this.f7, _this.f7);
          _nw360[(new BigNumber(9)).toNumber()] = (_2387_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), ((_2390_v2) => function (_2391_i3) {
            return _2390_v2;
          })(_2359_v2))).length), new BigNumber((_2387_v20).length))];
          _nw360[(new BigNumber(10)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(11)).toNumber()] = (_2368_v11).Intersect(_2368_v11);
          _nw360[(new BigNumber(12)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(13)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(14)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(15)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(16)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.FromArray(_2365_v8);
          _nw360[(new BigNumber(18)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_2358_cf7, _this.f7, _2358_cf7);
          _nw360[(new BigNumber(20)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(21)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.FromArray((_2388_v21)[_module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber((_2388_v21).length))]);
          _nw360[(new BigNumber(23)).toNumber()] = _dafny.MultiSet.fromElements(_this.f7);
          _nw360[(new BigNumber(24)).toNumber()] = _2368_v11;
          _nw360[(new BigNumber(25)).toNumber()] = _2368_v11;
          _2389_v22 = _nw360;
          let _index379 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_2389_v22).length));
          (_2389_v22)[_index379] = _2368_v11;
          let _index380 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_2389_v22).length));
          let _rhs332 = _this.f7;
          let _rhs333 = ((!(_2358_cf7).isEqualTo(_2358_cf7)) ? (_2368_v11) : (_2368_v11));
          let _lhs242 = _2389_v22;
          let _lhs243 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_2389_v22).length));
          _2358_cf7 = _rhs332;
          _lhs242[_lhs243] = _rhs333;
        } else {
          let _2392___mcc_h9 = (_source30).cf49;
          let _2393_cf49 = _2392___mcc_h9;
          let _2394_v23;
          let _nw361 = new _module.C1();
          _nw361.__ctor((_this).f6, p1.f3);
          _2394_v23 = _nw361;
          let _2395_v24;
          _2395_v24 = _module.D1.create_DC5((p1).f4, true, _this.f7, _this.f7);
          let _2396_v25;
          _2396_v25 = _module.D19.create_DC46(_this.f7, p1.f3, _this.f3, _2395_v24, _2358_cf7);
          r0 = ((!((_this).f4) || ((_2396_v25).dtor_cf75)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_2397_v2) => function (_2398_i4) {
            return _2397_v2;
          })(_2359_v2))) : ((_2394_v23).f12));
          let _2399_v26;
          let _nw362 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2399_v26 = _nw362;
          let _2400_v27;
          let _nw363 = Array((new BigNumber(13)).toNumber());
          _nw363[(_dafny.ZERO).toNumber()] = _2399_v26;
          _nw363[(_dafny.ONE).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(2)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(3)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(4)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(5)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(6)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(7)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(8)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(9)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(10)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(11)).toNumber()] = _2399_v26;
          _nw363[(new BigNumber(12)).toNumber()] = _2399_v26;
          _2400_v27 = _nw363;
          let _2401_v28;
          let _nw364 = Array((new BigNumber(25)).toNumber());
          _nw364[(_dafny.ZERO).toNumber()] = _2353_v0;
          _nw364[(_dafny.ONE).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(2)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(3)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(4)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(5)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(6)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(7)).toNumber()] = (((p1).f4) ? (_2353_v0) : (_2353_v0));
          _nw364[(new BigNumber(8)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(9)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(10)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(11)).toNumber()] = (((p1).f4) ? (_2353_v0) : (_2353_v0));
          _nw364[(new BigNumber(12)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(13)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(14)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(15)).toNumber()] = ((false) ? (_2353_v0) : (_2353_v0));
          _nw364[(new BigNumber(16)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(17)).toNumber()] = ((p1.f3) ? (_2353_v0) : (_2353_v0));
          _nw364[(new BigNumber(18)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(19)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(20)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(21)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(22)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(23)).toNumber()] = _2353_v0;
          _nw364[(new BigNumber(24)).toNumber()] = _2353_v0;
          _2401_v28 = _nw364;
          let _index381 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2401_v28).length));
          (_2401_v28)[_index381] = _2353_v0;
          let _index382 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2401_v28).length));
          let _rhs334 = _2365_v8;
          let _rhs335 = _2400_v27;
          let _rhs336 = _2353_v0;
          let _rhs337 = _this.f7;
          let _lhs244 = _2401_v28;
          let _lhs245 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2401_v28).length));
          let _lhs246 = _this;
          _2365_v8 = _rhs334;
          _2400_v27 = _rhs335;
          _lhs244[_lhs245] = _rhs336;
          _lhs246.f7 = _rhs337;
          (globalState).f0 = _dafny.Seq.IsPrefixOf(_module.__default.fm34(_2358_cf7, _this.f7, globalState), _dafny.Seq.Concat(_2364_v7, _dafny.Seq.of(p1.f3, p1.f3)));
        }
        let _2402_v29;
        let _nw365 = Array((_dafny.ONE).toNumber());
        _2402_v29 = _nw365;
        let _2403_v30;
        let _nw366 = new _module.C3();
        _nw366.__ctor(p1.f3, _this.f3);
        _2403_v30 = _nw366;
        let _index383 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_2402_v29).length));
        (_2402_v29)[_index383] = _2403_v30;
        let _2404_v31;
        _2404_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f6).length),new BigNumber(128));
        let _2405_v32;
        _2405_v32 = _dafny.MultiSet.fromElements(_2404_v31);
        let _2406_v33;
        _2406_v33 = _module.D22.create_DC55(_2403_v30);
        let _index384 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_2402_v29).length));
        (_2402_v29)[_index384] = ((_dafny.areEqual(_2364_v7, _dafny.Seq.of((((_2363_v6).contains(_module.__default.fm21(_2358_cf7, _2405_v32, globalState))) ? ((_2363_v6).get(_module.__default.fm21(_2358_cf7, _2405_v32, globalState))) : ((_this).f4))))) ? (_2403_v30) : (((true) ? (_2403_v30) : ((_2406_v33).dtor_cf88))));
      } else if (_source29.is_DC5) {
        let _2407___mcc_h1 = (_source29).cf8;
        let _2408___mcc_h2 = (_source29).cf9;
        let _2409___mcc_h3 = (_source29).cf10;
        let _2410___mcc_h4 = (_source29).cf11;
        let _2411_cf11 = _2410___mcc_h4;
        let _2412_cf10 = _2409___mcc_h3;
        let _2413_cf9 = _2408___mcc_h2;
        let _2414_cf8 = _2407___mcc_h1;
        let _2415_v34;
        let _nw367 = Array((new BigNumber(18)).toNumber());
        _2415_v34 = _nw367;
        (globalState).f0 = !(!(new BigNumber((_dafny.Seq.of(_2415_v34, _2415_v34, _2415_v34, _2415_v34)).length)).isEqualTo(_this.f7));
        let _2416_v35;
        let _nw368 = new _module.C1();
        _nw368.__ctor((_this).f6, (p1).f4);
        _2416_v35 = _nw368;
        _2411_cf11 = _2411_cf11;
        let _2417_v36;
        _2417_v36 = _dafny.Seq.of(_2413_cf9, !(!(p1.f3)));
        let _2418_v37;
        _2418_v37 = _dafny.Seq.of(false, (_2417_v36)[_module.__default.safeIndex((_this).fm3(globalState), new BigNumber((_2417_v36).length))], (_2412_cf10).isLessThanOrEqualTo(_2411_cf11));
        _2417_v36 = _module.__default.fm34(_module.__default.fm31(p1.f3, _2411_cf11, !(_this.f3), _dafny.Seq.of(!((p1).f4), !((p1).f4), (_2416_v35).f13), globalState), _this.f7, globalState);
      } else if (_source29.is_DC6) {
        if (_this.f3) {
          let _2419_v38;
          let _nw369 = new _module.C0();
          _nw369.__ctor();
          _2419_v38 = _nw369;
          (_this).f7 = _this.f7;
          let _2420_v39;
          let _nw370 = new _module.C2();
          _nw370.__ctor(true, (_this).f4, p1.f3, p1.f3);
          _2420_v39 = _nw370;
          let _2421_v40;
          _2421_v40 = _dafny.Seq.of(_2420_v39);
          let _2422_v41;
          _2422_v41 = _dafny.Seq.of(_this.f7, _this.f7, new BigNumber((_2421_v40).length), _this.f7);
          let _2423_v42;
          _2423_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2422_v41,!((_2420_v39).f10));
          let _2424_v43;
          _2424_v43 = _dafny.Seq.of((_this).f4, p1.f3, p1.f3);
          let _2425_v44;
          _2425_v44 = _dafny.Set.fromElements(p1.f3, false);
          let _2426_v45;
          _2426_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p1.f3);
          let _2427_v46;
          let _nw371 = Array((new BigNumber(14)).toNumber());
          _nw371[(_dafny.ZERO).toNumber()] = _this.f7;
          _nw371[(_dafny.ONE).toNumber()] = ((_module.__default.fm4(false, globalState)) ? (_this.f7) : (_this.f7));
          _nw371[(new BigNumber(2)).toNumber()] = _this.f7;
          _nw371[(new BigNumber(3)).toNumber()] = _this.f7;
          _nw371[(new BigNumber(4)).toNumber()] = _this.f7;
          _nw371[(new BigNumber(5)).toNumber()] = _module.__default.fm31((((_2423_v42).contains(_2422_v41)) ? ((_2423_v42).get(_2422_v41)) : ((_2424_v43)[_module.__default.safeIndex(_this.f7, new BigNumber((_2424_v43).length))])), _this.f7, (_this).f4, _2424_v43, globalState);
          _nw371[(new BigNumber(6)).toNumber()] = new BigNumber((_2424_v43).length);
          _nw371[(new BigNumber(7)).toNumber()] = new BigNumber((_2425_v44).length);
          _nw371[(new BigNumber(8)).toNumber()] = _this.f7;
          _nw371[(new BigNumber(9)).toNumber()] = (_this.f7).multipliedBy(_this.f7);
          _nw371[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(_this.f7, new BigNumber((_2426_v45).length));
          _nw371[(new BigNumber(11)).toNumber()] = new BigNumber(-452);
          _nw371[(new BigNumber(12)).toNumber()] = _this.f7;
          _nw371[(new BigNumber(13)).toNumber()] = new BigNumber((_2422_v41).length);
          _2427_v46 = _nw371;
          let _2428_v47;
          _2428_v47 = _dafny.Set.fromElements(_this.f7);
          let _2429_v48;
          _2429_v48 = _dafny.Seq.of((_this).f6, (_this).f6, (_this).f6, (_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), ((_2430_p0) => function (_2431_i5) {
            return _2430_p0;
          })(p0)));
          let _rhs338 = (_2420_v39).fm0(globalState);
          let _rhs339 = (((p1).f4) ? (_this.f7) : (new BigNumber((_2428_v47).length)));
          let _rhs340 = (new BigNumber(((_2429_v48)[_module.__default.safeIndex(_this.f7, new BigNumber((_2429_v48).length))]).length)).multipliedBy(_this.f7);
          let _rhs341 = _2427_v46;
          let _lhs247 = _this;
          let _lhs248 = _this;
          let _lhs249 = _this;
          _lhs247.f7 = _rhs338;
          _lhs248.f7 = _rhs339;
          _lhs249.f7 = _rhs340;
          _2427_v46 = _rhs341;
          let _2432_v49;
          _2432_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f7);
          let _2433_v50;
          _2433_v50 = _module.D0.create_DC2((_this).f4, (_2422_v41)[_module.__default.safeIndex(new BigNumber((_2432_v49).length), new BigNumber((_2422_v41).length))], _this.f7);
          let _2434_v51;
          _2434_v51 = _dafny.MultiSet.fromElements((_2433_v50).dtor_cf5, new BigNumber(341), new BigNumber((_2424_v43).length));
          _2434_v51 = _dafny.MultiSet.fromElements(_this.f7);
          let _2435_v52;
          let _nw372 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _2435_v52 = _nw372;
          let _2436_v53;
          _2436_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p1.f3);
          let _2437_v54;
          _2437_v54 = _dafny.Map.Empty.slice().updateUnsafe(_2436_v53,p1.f3);
          let _index385 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_2435_v52).length));
          (_2435_v52)[_index385] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm15(_this.f3, _this.f7, _this.f7, p1.f3, globalState),(_this).f4)).Merge(_2437_v54);
          let _2438_v56;
          _2438_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), function (_2439_i6) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length));
          let _index386 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_2435_v52).length));
          (_2435_v52)[_index386] = function () {
            let _coll78 = new _dafny.Map();
            for (const _compr_78 of ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(p1).f4),_this.f7)).update(_2436_v53, new BigNumber((_2438_v56).length))).Keys.Elements) {
              let _2440_v55 = _compr_78;
              if (((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(p1).f4),_this.f7)).update(_2436_v53, new BigNumber((_2438_v56).length))).contains(_2440_v55)) {
                _coll78.push([_2440_v55,!(false)]);
              }
            }
            return _coll78;
          }();
        } else {
          let _2441_v57;
          _2441_v57 = _dafny.Seq.of((p1).f4);
          (_this).f3 = !(_dafny.Seq.IsPrefixOf(_2441_v57, _2441_v57));
          let _rhs342 = p1.f3;
          let _rhs343 = _this.f7;
          let _lhs250 = globalState;
          let _lhs251 = _this;
          _lhs250.f0 = _rhs342;
          _lhs251.f7 = _rhs343;
          let _2442_v58;
          _2442_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(p1).f4);
          let _2443_v59;
          _2443_v59 = _dafny.Map.Empty.slice().updateUnsafe((((_2442_v58).contains((p1).fm0(globalState))) ? ((_2442_v58).get((p1).fm0(globalState))) : (p1.f3)),new BigNumber(365));
          let _2444_v60;
          _2444_v60 = _module.D14.create_DC31(_this.f7, _this.f7);
          let _pat_let_tv68 = globalState;
          let _2445_v61;
          _2445_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(((_2443_v59).contains((p1).f4)) ? ((_2443_v59).get((p1).f4)) : ((function (_pat_let66_0) {
            return function (_2446_dt__update__tmp_h0) {
              return function (_pat_let67_0) {
                return function (_2447_dt__update_hcf44_h0) {
                  return _module.D14.create_DC31(_2447_dt__update_hcf44_h0, (_2446_dt__update__tmp_h0).dtor_cf45);
                }(_pat_let67_0);
              }((_this).fm3(_pat_let_tv68));
            }(_pat_let66_0);
          }(_2444_v60)).dtor_cf44)));
          let _2448_v62;
          _2448_v62 = _dafny.MultiSet.fromElements(_2445_v61);
          _2448_v62 = _2448_v62;
          (globalState).f0 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(175), _this.f7))).isEqualTo(new BigNumber(216));
          let _2449_v63;
          _2449_v63 = _dafny.MultiSet.fromElements(p1.f3);
          let _2450_v64;
          _2450_v64 = new _dafny.CodePoint('r'.codePointAt(0));
          let _rhs344 = (_2449_v63).Intersect(_2449_v63);
          let _rhs345 = _2450_v64;
          let _rhs346 = _this.f7;
          let _rhs347 = (p1).f4;
          let _lhs252 = _this;
          let _lhs253 = globalState;
          _2449_v63 = _rhs344;
          _2450_v64 = _rhs345;
          _lhs252.f7 = _rhs346;
          _lhs253.f0 = _rhs347;
        }
        (_this).f7 = _this.f7;
        let _2451_v65;
        _2451_v65 = _dafny.MultiSet.fromElements(_this.f7);
        _2451_v65 = _2451_v65;
        let _2452_v66;
        let _nw373 = new _module.C4();
        _nw373.__ctor(p1.f3, _this.f7, (p1).f4, (p1).f4);
        _2452_v66 = _nw373;
        let _2453_v67;
        _2453_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2452_v66);
        let _2454_v68;
        let _nw374 = Array((new BigNumber(10)).toNumber());
        _2454_v68 = _nw374;
        let _2455_v69;
        _2455_v69 = _dafny.MultiSet.fromElements(_2454_v68, _2454_v68, _2454_v68, _2454_v68, _2454_v68);
        let _2456_v70;
        let _nw375 = new _module.C2();
        _nw375.__ctor(_this.f3, (new BigNumber((_dafny.Set.fromElements(_this.f7, new BigNumber((_2453_v67).length), _this.f7)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_this.f7)), (_2455_v69).IsSubsetOf(_2455_v69), !(_module.__default.fm4(false, globalState)));
        _2456_v70 = _nw375;
      } else {
        let _2457___mcc_h5 = (_source29).cf6;
        let _2458_cf6 = _2457___mcc_h5;
        let _index387 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2458_cf6).length));
        (_2458_cf6)[_index387] = (_this).f4;
        let _index388 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2458_cf6).length));
        (_2458_cf6)[_index388] = p1.f3;
        let _2459_v71;
        let _nw376 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _2459_v71 = _nw376;
        let _2460_v72;
        _2460_v72 = _dafny.Set.fromElements(_2459_v71);
        (_this).f7 = new BigNumber(((_dafny.Set.fromElements(_2459_v71)).Intersect(_2460_v72)).length);
        let _2461_v73;
        let _nw377 = Array((new BigNumber(13)).toNumber());
        _nw377[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_this).f6, (_this).f6);
        _nw377[(_dafny.ONE).toNumber()] = _dafny.Seq.update((_this).f6, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("urawnxuq")).length), new BigNumber(((_this).f6).length)), _module.__default.fm23(_this.f3, _this.f7, globalState));
        _nw377[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat((_this).f6, (_this).f6);
        _nw377[(new BigNumber(3)).toNumber()] = ((p1.f3) ? ((_this).f6) : ((_this).f6));
        _nw377[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat((_this).f6, (_this).f6);
        _nw377[(new BigNumber(5)).toNumber()] = _dafny.Seq.update((_module.D17.create_DC38((_this).f6)).dtor_cf56, _module.__default.safeIndex(new BigNumber(219), new BigNumber(((_module.D17.create_DC38((_this).f6)).dtor_cf56).length)), p0);
        _nw377[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("doqvtbv"), (_this).f6);
        _nw377[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_2462_p0) => function (_2463_i7) {
          return _2462_p0;
        })(p0));
        _nw377[(new BigNumber(8)).toNumber()] = (_this).f6;
        _nw377[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), ((_2464_p0) => function (_2465_i8) {
          return _2464_p0;
        })(p0));
        _nw377[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("porrudmcw");
        _nw377[(new BigNumber(11)).toNumber()] = (_this).f6;
        _nw377[(new BigNumber(12)).toNumber()] = (_this).f6;
        _2461_v73 = _nw377;
        let _index389 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2461_v73).length));
        (_2461_v73)[_index389] = (_this).f6;
        let _index390 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2461_v73).length));
        (_2461_v73)[_index390] = _dafny.Seq.Concat((_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-63)), ((_2466_p0) => function (_2467_i9) {
          return _2466_p0;
        })(p0)));
        let _2468_v74;
        _2468_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,new BigNumber(707));
        let _2469_v75;
        _2469_v75 = _dafny.Seq.of((p1).f4, !(true), !(_this.f3));
        _2468_v74 = (_2468_v74).update((p1).f4, new BigNumber((_2469_v75).length));
      }
      let _2470_v76;
      let _nw378 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _2470_v76 = _nw378;
      let _2471_v77;
      _2471_v77 = _dafny.Seq.of(_2470_v76);
      let _2472_v78;
      _2472_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber(810));
      if (((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2471_v77).length), _this.f7))).isLessThan(_module.__default.safeModuloInt(new BigNumber((_2472_v78).length), _this.f7))) {
        let _2473_v79;
        _2473_v79 = _dafny.Set.fromElements(_this.f7);
        let _2474_v80;
        _2474_v80 = _module.D16.create_DC36(_2473_v79, p1.f3, false, !(_module.__default.fm14(globalState)));
        let _2475_v81;
        _2475_v81 = _dafny.Seq.of(true, (p1).f4);
        let _pat_let_tv69 = _2475_v81;
        let _pat_let_tv70 = _2475_v81;
        let _source31 = function (_pat_let68_0) {
          return function (_2476_dt__update__tmp_h1) {
            return function (_pat_let69_0) {
              return function (_2477_dt__update_hcf53_h0) {
                return function (_pat_let70_0) {
                  return function (_2478_dt__update_hcf51_h0) {
                    return _module.D16.create_DC36(_2478_dt__update_hcf51_h0, (_2476_dt__update__tmp_h1).dtor_cf52, _2477_dt__update_hcf53_h0, (_2476_dt__update__tmp_h1).dtor_cf54);
                  }(_pat_let70_0);
                }(_dafny.Set.fromElements(new BigNumber(706)));
              }(_pat_let69_0);
            }((_pat_let_tv69)[_module.__default.safeIndex(_this.f7, new BigNumber((_pat_let_tv70).length))]);
          }(_pat_let68_0);
        }(_2474_v80);
        if (_source31.is_DC36) {
          let _2479___mcc_h10 = (_source31).cf51;
          let _2480___mcc_h11 = (_source31).cf52;
          let _2481___mcc_h12 = (_source31).cf53;
          let _2482___mcc_h13 = (_source31).cf54;
          let _2483_cf54 = _2482___mcc_h13;
          let _2484_cf53 = _2481___mcc_h12;
          let _2485_cf52 = _2480___mcc_h11;
          let _2486_cf51 = _2479___mcc_h10;
          let _2487_v82;
          let _nw379 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2487_v82 = _nw379;
          let _index391 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2487_v82).length));
          (_2487_v82)[_index391] = (_this).f6;
          let _index392 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_2487_v82).length));
          (_2487_v82)[_index392] = (_this).f6;
          _2470_v76 = _2470_v76;
          _2470_v76 = _2470_v76;
          (_this).m4(globalState);
        } else {
          let _2488___mcc_h14 = (_source31).cf50;
          let _2489_cf50 = _2488___mcc_h14;
          let _2490_v83;
          let _init73 = function (_2491_i10) {
            return _module.D12.create_DC28(_module.D12.create_DC27(_this.f7));
          };
          let _nw380 = Array((new BigNumber(14)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw380.length); _i0_73++) {
            _nw380[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2490_v83 = _nw380;
          let _2492_v84;
          _2492_v84 = _module.D12.create_DC27(_this.f7);
          let _2493_v85;
          _2493_v85 = _module.D12.create_DC28(_2492_v84);
          let _2494_v86;
          _2494_v86 = _module.D12.create_DC28(_2492_v84);
          let _index393 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2490_v83).length));
          (_2490_v83)[_index393] = _2494_v86;
          let _pat_let_tv71 = _2493_v85;
          let _index394 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2490_v83).length));
          (_2490_v83)[_index394] = function (_pat_let71_0) {
            return function (_2495_dt__update__tmp_h2) {
              return function (_pat_let72_0) {
                return function (_2496_dt__update_hcf41_h0) {
                  return _module.D12.create_DC28(_2496_dt__update_hcf41_h0);
                }(_pat_let72_0);
              }(_pat_let_tv71);
            }(_pat_let71_0);
          }(_2494_v86);
          (globalState).f0 = ((_this.f3) ? (!(!((p1.f3) === (p1.f3)))) : (p1.f3));
          let _2497_v87;
          _2497_v87 = _module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2475_v81).length),_module.__default.fm31(p1.f3, new BigNumber(175), _this.f3, _dafny.Seq.of((_this).f4, (_this).f4), globalState)));
          let _2498_v88;
          _2498_v88 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f7),(_2497_v87).dtor_cf37);
          let _2499_v89;
          let _nw381 = Array((new BigNumber(11)).toNumber());
          _2499_v89 = _nw381;
          let _2500_v90;
          _2500_v90 = _dafny.Seq.of(_2499_v89);
          let _2501_v91;
          _2501_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
          let _2502_v92;
          _2502_v92 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p0);
          (_this).f3 = _module.__default.fm18(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2498_v88).contains(new BigNumber((_dafny.Seq.update(_2500_v90, _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm31((p1).f4, _this.f7, true, _2475_v81, globalState)), new BigNumber((_2500_v90).length)), _2499_v89)).length))) ? ((_2498_v88).get(new BigNumber((_dafny.Seq.update(_2500_v90, _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm31((p1).f4, _this.f7, true, _2475_v81, globalState)), new BigNumber((_2500_v90).length)), _2499_v89)).length))) : (_2501_v91))).length),new BigNumber((_2502_v92).length))).length), p0, _this.f7, globalState);
          let _2503_v93;
          _2503_v93 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,false);
          (p1).f3 = ((_dafny.ZERO).minus(new BigNumber((_2503_v93).length))).isLessThanOrEqualTo(_this.f7);
        }
        let _index395 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2470_v76).length));
        (_2470_v76)[_index395] = _this.f7;
        let _2504_v95;
        let _init74 = function (_2505_i11) {
          return function () {
            let _coll79 = new _dafny.Set();
            for (const _compr_79 of _dafny.IntegerRange(new BigNumber(351), new BigNumber(239))) {
              let _2506_v94 = _compr_79;
              if (((new BigNumber(351)).isLessThanOrEqualTo(_2506_v94)) && ((_2506_v94).isLessThan(new BigNumber(239)))) {
                _coll79.add(_module.__default.safeDivisionInt(_2506_v94, _this.f7));
              }
            }
            return _coll79;
          }();
        };
        let _nw382 = Array((new BigNumber(15)).toNumber());
        for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw382.length); _i0_74++) {
          _nw382[_i0_74] = _init74(new BigNumber(_i0_74));
        }
        _2504_v95 = _nw382;
        let _2507_v96;
        _2507_v96 = _dafny.Seq.of(_2504_v95);
        let _2508_v97;
        _2508_v97 = _module.D15.create_DC32(_2504_v95);
        let _2509_v98;
        _2509_v98 = _dafny.MultiSet.fromElements(_module.D15.create_DC32((_2507_v96)[_module.__default.safeIndex(_this.f7, new BigNumber((_2507_v96).length))]), _2508_v97, _2508_v97, _2508_v97);
        let _2510_v99;
        _2510_v99 = _dafny.Map.Empty.slice().updateUnsafe(_2509_v98,(_this.f7).isLessThan(_this.f7));
        let _index396 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2470_v76).length));
        let _rhs348 = _module.__default.safeModuloInt((new BigNumber((_2472_v78).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cpdfirkd")).length))), new BigNumber((_module.__default.fm19(globalState)).length));
        let _rhs349 = ((_dafny.Map.Empty.slice().updateUnsafe(_2509_v98,(p1).f4)).update(_2509_v98, (p1).f4)).Merge(_2510_v99);
        let _lhs254 = _2470_v76;
        let _lhs255 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2470_v76).length));
        _lhs254[_lhs255] = _rhs348;
        _2510_v99 = _rhs349;
        let _2511_v100;
        _2511_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(p1).f4);
        _2511_v100 = (_2511_v100).update(_dafny.Seq.Concat((_this).f6, (_this).f6), !(true));
        _2353_v0 = (_2356_v1).dtor_cf6;
        let _2512_v101;
        _2512_v101 = _dafny.MultiSet.fromElements(_this.f3);
        let _2513_v102;
        _2513_v102 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,new BigNumber((_2512_v101).cardinality()));
        let _2514_v103;
        _2514_v103 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!((p1).f4),_module.__default.fm18((_2470_v76)[_module.__default.safeIndex(new BigNumber(336), new BigNumber((_2470_v76).length))], p0, new BigNumber(5), globalState)),_2513_v102);
        _2514_v103 = _2514_v103;
      } else {
        (_this).f7 = (((_2472_v78).contains(p1.f3)) ? ((_2472_v78).get(p1.f3)) : ((_this.f7).multipliedBy(_this.f7)));
        let _2515_v104;
        _2515_v104 = _module.D1.create_DC5((p1).f4, p1.f3, _this.f7, _this.f7);
        let _2516_v105;
        _2516_v105 = _dafny.Seq.of((_this).f6);
        let _rhs350 = !_dafny.areEqual(_module.__default.fm30(_2515_v104, _this.f7, _this.f7, globalState), (_2516_v105)[_module.__default.safeIndex(_this.f7, new BigNumber((_2516_v105).length))]);
        let _rhs351 = _2353_v0;
        let _rhs352 = true;
        let _rhs353 = _this.f7;
        let _lhs256 = p1;
        let _lhs257 = globalState;
        let _lhs258 = _this;
        _lhs256.f3 = _rhs350;
        _2353_v0 = _rhs351;
        _lhs257.f0 = _rhs352;
        _lhs258.f7 = _rhs353;
        if (((p1.f3) ? ((_this.f7).isEqualTo(new BigNumber((_dafny.Seq.of(p1)).length))) : ((p1).f4))) {
          let _2517_v106;
          let _nw383 = Array((new BigNumber(22)).toNumber()).fill(_module.D17.Default());
          _2517_v106 = _nw383;
          _2517_v106 = _2517_v106;
          (_this).m4(globalState);
          let _2518_v107;
          _2518_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2353_v0,(_this.f7).minus(_this.f7));
          _2518_v107 = (_2518_v107).update(_2353_v0, _this.f7);
          let _2519_v108;
          _2519_v108 = _dafny.Seq.of(p1.f3);
          let _2520_v109;
          _2520_v109 = _module.D4.create_DC13(_module.D4.create_DC11(_2519_v108));
          let _2521_v110;
          _2521_v110 = _dafny.MultiSet.fromElements(_2520_v109, _2520_v109, _2520_v109, _module.__default.fm43(_this.f3, _this.f7, (p1).f4, globalState), _2520_v109);
          let _2522_v111;
          let _nw384 = new _module.C2();
          _nw384.__ctor((_2521_v110).contains(_2520_v109), !_dafny.areEqual((_this).f6, _dafny.Seq.update(_module.__default.fm30(_2515_v104, new BigNumber(86), _this.f7, globalState), _module.__default.safeIndex(_this.f7, new BigNumber((_module.__default.fm30(_2515_v104, new BigNumber(86), _this.f7, globalState)).length)), new _dafny.CodePoint('i'.codePointAt(0)))), _this.f3, (p1).f4);
          _2522_v111 = _nw384;
          let _2523_v112;
          _2523_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2470_v76,(_2522_v111).f11);
          let _2524_v113;
          _2524_v113 = _dafny.Set.fromElements(new BigNumber((_2523_v112).length), _this.f7);
          _2524_v113 = _2524_v113;
        } else {
          let _2525_v114;
          let _nw385 = new _module.C1();
          _nw385.__ctor((_this).f6, _this.f3);
          _2525_v114 = _nw385;
          let _2526_v115;
          _2526_v115 = _dafny.Map.Empty.slice().updateUnsafe(_2525_v114,(_this).f6);
          let _2527_v116;
          _2527_v116 = _dafny.Map.Empty.slice().updateUnsafe(true,(p1).f4);
          let _2528_v117;
          _2528_v117 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2525_v114).f12);
          let _2529_v118;
          _2529_v118 = _module.D17.create_DC38(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ymgcfot"), _module.__default.safeIndex(_this.f7, new BigNumber((_dafny.Seq.UnicodeFromString("ymgcfot")).length)), p0));
          let _2530_v119;
          let _nw386 = Array((new BigNumber(29)).toNumber());
          _nw386[(_dafny.ZERO).toNumber()] = (_this).f6;
          _nw386[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("tufjnmcu");
          _nw386[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(3)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), ((_2531_p0) => function (_2532_i12) {
            return _2531_p0;
          })(p0));
          _nw386[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("drlwyki");
          _nw386[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((((_2526_v115).contains(_2525_v114)) ? ((_2526_v115).get(_2525_v114)) : (_dafny.Seq.update((_2525_v114).f12, _module.__default.safeIndex(new BigNumber((_2527_v116).length), new BigNumber(((_2525_v114).f12).length)), new _dafny.CodePoint('p'.codePointAt(0))))), (_this).f6);
          _nw386[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-178)), ((_2533_p0) => function (_2534_i13) {
            return _2533_p0;
          })(p0));
          _nw386[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat((_this).f6, (_2525_v114).f12);
          _nw386[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iv"), (_2525_v114).f12);
          _nw386[(new BigNumber(10)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("kagcscus");
          _nw386[(new BigNumber(12)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(13)).toNumber()] = (_2525_v114).f12;
          _nw386[(new BigNumber(14)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(15)).toNumber()] = (((_this).f4) ? ((_2525_v114).f12) : (_dafny.Seq.UnicodeFromString("ldrf")));
          _nw386[(new BigNumber(16)).toNumber()] = ((_this.f3) ? ((_2525_v114).f12) : ((_this).f6));
          _nw386[(new BigNumber(17)).toNumber()] = (_2525_v114).f12;
          _nw386[(new BigNumber(18)).toNumber()] = (_2525_v114).f12;
          _nw386[(new BigNumber(19)).toNumber()] = (_2525_v114).f12;
          _nw386[(new BigNumber(20)).toNumber()] = (_2525_v114).f12;
          _nw386[(new BigNumber(21)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(22)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(23)).toNumber()] = (((_2528_v117).contains((p1).f4)) ? ((_2528_v117).get((p1).f4)) : ((_this).f6));
          _nw386[(new BigNumber(24)).toNumber()] = (_this).f6;
          _nw386[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lfmg"), (_2529_v118).dtor_cf56);
          _nw386[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat((_this).f6, _dafny.Seq.UnicodeFromString("ukbecno"));
          _nw386[(new BigNumber(27)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lxn"), (_this).f6);
          _nw386[(new BigNumber(28)).toNumber()] = (_this).f6;
          _2530_v119 = _nw386;
          let _index397 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2530_v119).length));
          (_2530_v119)[_index397] = (_2525_v114).f12;
          let _index398 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2530_v119).length));
          (_2530_v119)[_index398] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_2535_p0) => function (_2536_i14) {
            return _2535_p0;
          })(p0)), (_2525_v114).f12), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pnqsnqy"), (_this).f6));
          (p1).f3 = (p1).f4;
          (p1).f3 = !(_this.f3);
          let _2537_v120;
          let _nw387 = Array((new BigNumber(11)).toNumber());
          _nw387[(_dafny.ZERO).toNumber()] = _2470_v76;
          _nw387[(_dafny.ONE).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(2)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(3)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(4)).toNumber()] = (_2471_v77)[_module.__default.safeIndex(_this.f7, new BigNumber((_2471_v77).length))];
          _nw387[(new BigNumber(5)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(6)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(7)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(8)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(9)).toNumber()] = _2470_v76;
          _nw387[(new BigNumber(10)).toNumber()] = _2470_v76;
          _2537_v120 = _nw387;
          let _2538_v121;
          _2538_v121 = _module.D20.create_DC50(_this.f7, (p1).fm0(globalState));
          let _2539_v122;
          _2539_v122 = _dafny.MultiSet.fromElements(p1.f3, (_this).f4, !((p1).f4), p1.f3, p1.f3);
          let _rhs354 = _module.__default.safeDivisionInt(_module.__default.fm31((_this).f4, new BigNumber(172), (p1).f4, _dafny.Seq.update(_module.__default.fm34(_this.f7, new BigNumber(294), globalState), _module.__default.safeIndex(new BigNumber(((_this).f6).length), new BigNumber((_module.__default.fm34(_this.f7, new BigNumber(294), globalState)).length)), p1.f3), globalState), (((_2472_v78).contains((p1).f4)) ? ((_2472_v78).get((p1).f4)) : (_this.f7)));
          let _rhs355 = (p1).f4;
          let _rhs356 = _2537_v120;
          let _rhs357 = ((_2538_v121).dtor_cf82).plus((_this.f7).multipliedBy(_this.f7));
          let _rhs358 = ((p1.f3) ? ((_2539_v122).contains((p1).f4)) : (((!((_2525_v114).f13)) ? (!(p1.f3)) : (_this.f3))));
          let _lhs259 = _this;
          let _lhs260 = p1;
          let _lhs261 = _this;
          let _lhs262 = _this;
          _lhs259.f7 = _rhs354;
          _lhs260.f3 = _rhs355;
          _2537_v120 = _rhs356;
          _lhs261.f7 = _rhs357;
          _lhs262.f3 = _rhs358;
          (p1).f3 = (_2525_v114).f13;
        }
        (_this).f7 = new BigNumber(91);
        (_this).f7 = _this.f7;
      }
      (_this).f7 = _this.f7;
      let _index399 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_2470_v76).length));
      (_2470_v76)[_index399] = new BigNumber(((_this).f6).length);
      let _index400 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_2470_v76).length));
      (_2470_v76)[_index400] = _module.__default.safeModuloInt(_this.f7, new BigNumber(750));
      let _2540_v123;
      _2540_v123 = _dafny.MultiSet.fromElements((_2470_v76)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_2470_v76).length))]);
      let _2541_v124;
      let _nw388 = new _module.C1();
      _nw388.__ctor((_this).f6, (_2540_v123).IsSubsetOf(_2540_v123));
      _2541_v124 = _nw388;
      let _2542_v125;
      _2542_v125 = _dafny.Seq.of((_2470_v76)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_2470_v76).length))]);
      let _2543_v128;
      let _init75 = ((_2544_p0) => function (_2545_i15) {
        return function () {
          let _coll80 = new _dafny.Set();
          for (const _compr_80 of (function () {
            let _coll81 = new _dafny.Set();
            for (const _compr_81 of (_dafny.Map.Empty.slice().updateUnsafe(_2544_p0,_this.f3)).Keys.Elements) {
              let _2546_v126 = _compr_81;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_2544_p0,_this.f3)).contains(_2546_v126)) {
                _coll81.add(_2546_v126);
              }
            }
            return _coll81;
          }()).Elements) {
            let _2547_v127 = _compr_80;
            if ((function () {
              let _coll82 = new _dafny.Set();
              for (const _compr_82 of (_dafny.Map.Empty.slice().updateUnsafe(_2544_p0,_this.f3)).Keys.Elements) {
                let _2548_v126 = _compr_82;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_2544_p0,_this.f3)).contains(_2548_v126)) {
                  _coll82.add(_2548_v126);
                }
              }
              return _coll82;
            }()).contains(_2547_v127)) {
              _coll80.add(_2547_v127);
            }
          }
          return _coll80;
        }();
      })(p0);
      let _nw389 = Array((new BigNumber(20)).toNumber());
      for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw389.length); _i0_75++) {
        _nw389[_i0_75] = _init75(new BigNumber(_i0_75));
      }
      _2543_v128 = _nw389;
      let _2549_v129;
      _2549_v129 = _dafny.Seq.of(_2543_v128, _2543_v128);
      let _2550_v130;
      _2550_v130 = _dafny.Map.Empty.slice().updateUnsafe(_2542_v125,(_2549_v129)[_module.__default.safeIndex((_2470_v76)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_2470_v76).length))], new BigNumber((_2549_v129).length))]);
      _2550_v130 = (_2550_v130).update(_dafny.Seq.of((p1).fm0(globalState)), _2543_v128);
      r0 = (_2541_v124).f12;
      return r0;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _2551_v0;
      _2551_v0 = new _dafny.CodePoint('y'.codePointAt(0));
      let _2552_v1;
      _2552_v1 = _dafny.Seq.of(_2551_v0, _2551_v0, _2551_v0, _2551_v0);
      let _2553_v2;
      _2553_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(_2551_v0, _2551_v0));
      _2552_v1 = _dafny.Seq.Concat(_2552_v1, (((_2553_v2).contains(_this.f3)) ? ((_2553_v2).get(_this.f3)) : (p0)));
      if ((_this).f4) {
        let _2554_v3;
        let _nw390 = new _module.C3();
        _nw390.__ctor(p1, (_this).f4);
        _2554_v3 = _nw390;
        let _2555_v4;
        let _init76 = function (_2556_i0) {
          return ((false) ? (!((_this).f4)) : ((_this).f4));
        };
        let _nw391 = Array((new BigNumber(17)).toNumber());
        for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw391.length); _i0_76++) {
          _nw391[_i0_76] = _init76(new BigNumber(_i0_76));
        }
        _2555_v4 = _nw391;
        let _2557_v5;
        let _nw392 = new _module.C0();
        _nw392.__ctor();
        _2557_v5 = _nw392;
        let _2558_v6;
        _2558_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(138),p2);
        let _2559_v7;
        _2559_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2557_v5,new BigNumber((_2558_v6).length));
        let _2560_v8;
        _2560_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,p3);
        let _2561_v9;
        _2561_v9 = _module.D11.create_DC24(_2560_v8);
        let _pat_let_tv72 = _2560_v8;
        let _2562_v10;
        _2562_v10 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(((_2559_v7).contains(_2557_v5)) ? ((_2559_v7).get(_2557_v5)) : (new BigNumber((_dafny.Seq.of(new BigNumber(335))).length)))), (function (_pat_let73_0) {
          return function (_2563_dt__update__tmp_h0) {
            return function (_pat_let74_0) {
              return function (_2564_dt__update_hcf37_h0) {
                return _module.D11.create_DC24(_2564_dt__update_hcf37_h0);
              }(_pat_let74_0);
            }(_pat_let_tv72);
          }(_pat_let73_0);
        }(_2561_v9)).dtor_cf37, _2560_v8);
        let _2565_v11;
        _2565_v11 = _dafny.Seq.of(_this.f7);
        let _2566_v12;
        _2566_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2565_v11,_2551_v0);
        let _2567_v13;
        _2567_v13 = _module.D17.create_DC37(_2566_v12);
        let _2568_v14;
        _2568_v14 = _dafny.Seq.of(_2567_v13);
        let _2569_v16;
        _2569_v16 = _dafny.Seq.of(function () {
          let _coll83 = new _dafny.Set();
          for (const _compr_83 of (_2568_v14).Elements) {
            let _2570_v15 = _compr_83;
            if (_dafny.Seq.contains(_2568_v14, _2570_v15)) {
              _coll83.add(_2570_v15);
            }
          }
          return _coll83;
        }());
        let _2571_v17;
        _2571_v17 = _dafny.Seq.of((_2569_v16)[_module.__default.safeIndex(_this.f7, new BigNumber((_2569_v16).length))]);
        let _2572_v19;
        _2572_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2551_v0);
        let _2573_v20;
        _2573_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2572_v19);
        let _2574_v21;
        _2574_v21 = _dafny.Set.fromElements(_2567_v13);
        let _2575_v23;
        _2575_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2573_v20,function () {
          let _coll84 = new _dafny.Set();
          for (const _compr_84 of (_dafny.Map.Empty.slice().updateUnsafe(_2574_v21,_2551_v0)).Keys.Elements) {
            let _2576_v22 = _compr_84;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_2574_v21,_2551_v0)).contains(_2576_v22)) {
              _coll84.add(_2576_v22);
            }
          }
          return _coll84;
        }());
        let _2577_v24;
        _2577_v24 = _module.D23.create_DC58(_module.__default.fm44(_this.f3, _this.f3, globalState));
        let _2578_v25;
        _2578_v25 = _dafny.Set.fromElements((_this).f6);
        let _nw393 = Array((new BigNumber(9)).toNumber());
        _nw393[(_dafny.ZERO).toNumber()] = _this.f3;
        _nw393[(_dafny.ONE).toNumber()] = _module.__default.fm21(_this.f7, _2562_v10, globalState);
        _nw393[(new BigNumber(2)).toNumber()] = ((p2) ? (_this.f3) : (_this.f3));
        _nw393[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(p0, (_this).f6);
        _nw393[(new BigNumber(4)).toNumber()] = (_this).f4;
        _nw393[(new BigNumber(5)).toNumber()] = ((((_2575_v23).contains(_2573_v20)) ? ((_2575_v23).get(_2573_v20)) : ((_2577_v24).dtor_cf92))).IsSubsetOf(function () {
          let _coll85 = new _dafny.Set();
          for (const _compr_85 of (_2571_v17).Elements) {
            let _2579_v18 = _compr_85;
            if (_dafny.Seq.contains(_2571_v17, _2579_v18)) {
              _coll85.add(_2579_v18);
            }
          }
          return _coll85;
        }());
        _nw393[(new BigNumber(6)).toNumber()] = p2;
        _nw393[(new BigNumber(7)).toNumber()] = !(_2578_v25).equals(_2578_v25);
        _nw393[(new BigNumber(8)).toNumber()] = (_this).f4;
        _2555_v4 = _nw393;
        let _2580_v26;
        _2580_v26 = _dafny.Set.fromElements(p1, p1);
        let _2581_v27;
        _2581_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2580_v26,_this.f3);
        r3 = new BigNumber(((_2581_v27).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),(_this).f4))).length);
        let _2582_v28;
        _2582_v28 = _module.D0.create_DC2(p2, new BigNumber(999), ((true) ? (new BigNumber(861)) : (p3)));
        let _source32 = _2582_v28;
        if (_source32.is_DC1) {
          let _2583___mcc_h0 = (_source32).cf1;
          let _2584___mcc_h1 = (_source32).cf2;
          let _2585_cf2 = _2584___mcc_h1;
          let _2586_cf1 = _2583___mcc_h0;
          let _index401 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_2555_v4).length));
          (_2555_v4)[_index401] = (_this).f4;
          let _2587_v29;
          let _init77 = ((_2588_p2, _2589_p0, _2590_p3) => function (_2591_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2588_p2,_dafny.MultiSet.fromElements(new BigNumber((_2589_p0).length), _2590_p3));
          })(p2, p0, p3);
          let _nw394 = Array((new BigNumber(5)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw394.length); _i0_77++) {
            _nw394[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2587_v29 = _nw394;
          let _2592_v30;
          _2592_v30 = _dafny.MultiSet.fromElements(p3);
          let _2593_v31;
          _2593_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2592_v30);
          let _index402 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_2587_v29).length));
          (_2587_v29)[_index402] = _2593_v31;
          let _2594_v32;
          let _nw395 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2594_v32 = _nw395;
          let _2595_v33;
          _2595_v33 = _module.D17.create_DC39(_2594_v32, _2586_cf1, true, _this.f7, p3);
          let _index403 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_2555_v4).length));
          let _index404 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_2587_v29).length));
          let _rhs359 = _this.f7;
          let _rhs360 = (_this).f6;
          let _rhs361 = p1;
          let _rhs362 = (_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_dafny.MultiSet.fromElements((_2595_v33).dtor_cf60)).update(p3, _module.__default.abs(p3)))).Merge(_2593_v31);
          let _rhs363 = _2551_v0;
          let _lhs263 = _this;
          let _lhs264 = _2555_v4;
          let _lhs265 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_2555_v4).length));
          let _lhs266 = _2587_v29;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_2587_v29).length));
          _lhs263.f7 = _rhs359;
          _2552_v1 = _rhs360;
          _lhs264[_lhs265] = _rhs361;
          _lhs266[_lhs267] = _rhs362;
          _2585_cf2 = _rhs363;
          (globalState).f0 = _module.__default.fm18(p3, new _dafny.CodePoint('v'.codePointAt(0)), _module.__default.fm31(_this.f3, p3, false, _dafny.Seq.of(_this.f3), globalState), globalState);
          let _2596_v34;
          _2596_v34 = _dafny.Seq.of((_2592_v30).update(p3, _module.__default.abs(p3)));
          let _2597_v35;
          _2597_v35 = _dafny.Set.fromElements(_this.f7);
          (_this).f7 = (_2557_v5).fm10((_this.f7).multipliedBy(p3), (_this).f4, (_2596_v34)[_module.__default.safeIndex(new BigNumber((_2597_v35).length), new BigNumber((_2596_v34).length))], globalState);
          let _2598_v36;
          let _nw396 = new _module.C2();
          _nw396.__ctor(_2586_cf1, _this.f3, !(p2), p1);
          _2598_v36 = _nw396;
        } else if (_source32.is_DC2) {
          let _2599___mcc_h2 = (_source32).cf3;
          let _2600___mcc_h3 = (_source32).cf4;
          let _2601___mcc_h4 = (_source32).cf5;
          let _2602_cf5 = _2601___mcc_h4;
          let _2603_cf4 = _2600___mcc_h3;
          let _2604_cf3 = _2599___mcc_h2;
          let _rhs364 = (new BigNumber((_2565_v11).length)).plus((_this.f7).multipliedBy(_this.f7));
          let _rhs365 = _this.f3;
          let _rhs366 = _2555_v4;
          let _lhs268 = _this;
          let _lhs269 = _this;
          _lhs268.f7 = _rhs364;
          _lhs269.f3 = _rhs365;
          _2555_v4 = _rhs366;
          _2552_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(828)), function (_2605_i2) {
            return (_module.D0.create_DC1((_this).f4, new _dafny.CodePoint('m'.codePointAt(0)))).dtor_cf2;
          });
          let _2606_v37;
          _2606_v37 = _dafny.Map.Empty.slice().updateUnsafe((_2552_v1)[_module.__default.safeIndex(p3, new BigNumber((_2552_v1).length))],_2551_v0);
          let _2607_v39;
          _2607_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_module.__default.fm45(p3, _this.f3, globalState));
          let _2608_v40;
          _2608_v40 = _dafny.Seq.of(_module.__default.fm21(new BigNumber(316), _2562_v10, globalState), p2, _2604_cf3, _this.f3);
          let _2609_v41;
          _2609_v41 = _dafny.Seq.of(!(_dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_2608_v40).length))).equals(_2560_v8));
          let _2610_v42;
          _2610_v42 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_2604_cf3, !(_2604_cf3))).length));
          let _rhs367 = ((function () {
            let _coll86 = new _dafny.Map();
            for (const _compr_86 of ((_this).f6).Elements) {
              let _2611_v38 = _compr_86;
              if (_dafny.Seq.contains((_this).f6, _2611_v38)) {
                _coll86.push([_2611_v38,_2551_v0]);
              }
            }
            return _coll86;
          }()).Merge(_2606_v37)).Merge((((_2607_v39).contains(p1)) ? ((_2607_v39).get(p1)) : (_2606_v37)));
          let _rhs368 = new BigNumber((_2609_v41).length);
          let _rhs369 = !((_dafny.ZERO).minus(new BigNumber((_2610_v42).cardinality()))).isEqualTo(new BigNumber(-14));
          let _lhs270 = globalState;
          _2606_v37 = _rhs367;
          r2 = _rhs368;
          _lhs270.f0 = _rhs369;
          let _2612_v43;
          _2612_v43 = _dafny.MultiSet.fromElements(_2604_cf3);
          let _2613_v44;
          _2613_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2612_v43).update((_this).f4, _module.__default.abs(new BigNumber(565)))).cardinality()),_2565_v11);
          _2602_cf5 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this.f7).multipliedBy(new BigNumber((_2613_v44).length)), (_this).fm3(globalState)));
        } else {
          let _2614___mcc_h5 = (_source32).cf0;
          let _2615_cf0 = _2614___mcc_h5;
          _2558_v6 = (_2558_v6).update(_this.f7, !_dafny.areEqual((_this).f6, _2552_v1));
          r1 = p1;
          _2615_cf0 = _this.f7;
          let _2616_v45;
          _2616_v45 = _dafny.Seq.of(false);
          let _2617_v46;
          _2617_v46 = _dafny.MultiSet.fromElements(p1, !((_this).f4));
          r1 = ((_2617_v46).update(p1, _module.__default.abs(new BigNumber(15)))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_2616_v45));
        }
        let _2618_v47;
        let _init78 = function (_2619_i3) {
          return _module.__default.safeModuloInt(_2619_i3, _this.f7);
        };
        let _nw397 = Array((new BigNumber(20)).toNumber());
        for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw397.length); _i0_78++) {
          _nw397[_i0_78] = _init78(new BigNumber(_i0_78));
        }
        _2618_v47 = _nw397;
        let _index405 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2618_v47).length));
        (_2618_v47)[_index405] = p3;
        let _2620_v48;
        _2620_v48 = _dafny.Seq.of(!(_this.f3), p1);
        let _2621_v49;
        _2621_v49 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(_2620_v48, _module.__default.safeIndex(_this.f7, new BigNumber((_2620_v48).length)), (_this).f4), _2620_v48));
        let _index406 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2618_v47).length));
        (_2618_v47)[_index406] = new BigNumber(((_2621_v49)[_module.__default.safeIndex(p3, new BigNumber((_2621_v49).length))]).length);
      } else {
        r2 = p3;
        if (!(_this.f7).isEqualTo(new BigNumber((_module.__default.fm20((_this).f4, globalState)).cardinality()))) {
          let _2622_v50;
          _2622_v50 = _dafny.Seq.of(_this.f3);
          let _2623_v51;
          _2623_v51 = _dafny.Seq.of(_2622_v50);
          let _2624_v52;
          _2624_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2623_v51).length),(_this).f4);
          let _2625_v53;
          _2625_v53 = _dafny.Seq.of(_2624_v52, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(531),p1), _2624_v52);
          _2625_v53 = _2625_v53;
          (_this).f3 = _this.f3;
          let _2626_v54;
          let _nw398 = Array((new BigNumber(21)).toNumber()).fill(false);
          _2626_v54 = _nw398;
          let _index407 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_2626_v54).length));
          (_2626_v54)[_index407] = !((p1) || ((_this).f4));
          let _2627_v55;
          let _nw399 = Array((new BigNumber(23)).toNumber());
          _2627_v55 = _nw399;
          let _2628_v56;
          _2628_v56 = _dafny.MultiSet.fromElements(_2627_v55, _2627_v55, _2627_v55);
          let _index408 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_2626_v54).length));
          (_2626_v54)[_index408] = (_2628_v56).IsDisjointFrom(_dafny.MultiSet.fromElements(_2627_v55, _2627_v55));
          let _2629_v57;
          _2629_v57 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).fm3(globalState));
          let _2630_v58;
          _2630_v58 = _dafny.Set.fromElements((_2626_v54)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_2626_v54).length))]);
          let _2631_v59;
          _2631_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2630_v58,_2624_v52);
          _2629_v57 = (_2629_v57).update(p2, ((p2) ? (new BigNumber((_2631_v59).length)) : (_this.f7)));
          let _2632_v60;
          _2632_v60 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f7);
          _2632_v60 = (((!((_2626_v54)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_2626_v54).length))])) ? (_2632_v60) : (_2632_v60))).update(p3, p3);
        } else {
          let _2633_v61;
          let _nw400 = new _module.C0();
          _nw400.__ctor();
          _2633_v61 = _nw400;
          let _2634_v62;
          _2634_v62 = _dafny.MultiSet.fromElements(_2633_v61);
          let _2635_v63;
          let _nw401 = new _module.C3();
          _nw401.__ctor((_2634_v62).IsSubsetOf(_2634_v62), _this.f3);
          _2635_v63 = _nw401;
          _2635_v63 = _2635_v63;
          let _2636_v64;
          let _nw402 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2636_v64 = _nw402;
          let _2637_v65;
          let _nw403 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2637_v65 = _nw403;
          let _2638_v66;
          _2638_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2637_v65);
          let _2639_v67;
          _2639_v67 = _module.D21.create_DC51(_2637_v65);
          let _2640_v68;
          let _nw404 = Array((new BigNumber(14)).toNumber());
          _nw404[(_dafny.ZERO).toNumber()] = _2637_v65;
          _nw404[(_dafny.ONE).toNumber()] = (((_2638_v66).contains(p1)) ? ((_2638_v66).get(p1)) : (_2637_v65));
          _nw404[(new BigNumber(2)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(3)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(4)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(5)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(6)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(7)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(8)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(9)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(10)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(11)).toNumber()] = _2637_v65;
          _nw404[(new BigNumber(12)).toNumber()] = (_2639_v67).dtor_cf83;
          _nw404[(new BigNumber(13)).toNumber()] = _2637_v65;
          _2640_v68 = _nw404;
          let _index409 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2636_v64).length));
          (_2636_v64)[_index409] = _2640_v68;
          let _2641_v69;
          _2641_v69 = _module.D21.create_DC51(_2637_v65);
          let _2642_v70;
          _2642_v70 = _module.D21.create_DC54(_2641_v69);
          let _2643_v71;
          _2643_v71 = _dafny.Seq.of(_2635_v63.f3);
          let _2644_v72;
          _2644_v72 = _dafny.Seq.of(_2641_v69);
          let _pat_let_tv73 = _2641_v69;
          let _pat_let_tv74 = _2641_v69;
          let _pat_let_tv75 = _2641_v69;
          let _pat_let_tv76 = _2641_v69;
          let _2645_v73;
          let _nw405 = Array((new BigNumber(20)).toNumber());
          _nw405[(_dafny.ZERO).toNumber()] = _2642_v70;
          _nw405[(_dafny.ONE).toNumber()] = function (_pat_let75_0) {
            return function (_2646_dt__update__tmp_h1) {
              return function (_pat_let76_0) {
                return function (_2647_dt__update_hcf87_h0) {
                  return _module.D21.create_DC54(_2647_dt__update_hcf87_h0);
                }(_pat_let76_0);
              }(_pat_let_tv73);
            }(_pat_let75_0);
          }(_2642_v70);
          _nw405[(new BigNumber(2)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(3)).toNumber()] = (((_this).f4) ? (_2642_v70) : (_2642_v70));
          _nw405[(new BigNumber(4)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(5)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(6)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(7)).toNumber()] = ((_2635_v63.f3) ? (function (_pat_let77_0) {
            return function (_2648_dt__update__tmp_h2) {
              return function (_pat_let78_0) {
                return function (_2649_dt__update_hcf87_h1) {
                  return _module.D21.create_DC54(_2649_dt__update_hcf87_h1);
                }(_pat_let78_0);
              }(_pat_let_tv74);
            }(_pat_let77_0);
          }(_module.D21.create_DC54(_2641_v69))) : (_module.D21.create_DC54(_2641_v69)));
          _nw405[(new BigNumber(8)).toNumber()] = (((_2643_v71)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2635_v63.f3,_this.f3)).length), new BigNumber((_2643_v71).length))]) ? (_2642_v70) : (_2642_v70));
          _nw405[(new BigNumber(9)).toNumber()] = function (_pat_let79_0) {
            return function (_2650_dt__update__tmp_h3) {
              return function (_pat_let80_0) {
                return function (_2651_dt__update_hcf87_h2) {
                  return _module.D21.create_DC54(_2651_dt__update_hcf87_h2);
                }(_pat_let80_0);
              }(_pat_let_tv75);
            }(_pat_let79_0);
          }(_module.D21.create_DC54(_2641_v69));
          _nw405[(new BigNumber(10)).toNumber()] = ((p2) ? (_module.D21.create_DC54(_2641_v69)) : (_2642_v70));
          _nw405[(new BigNumber(11)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(12)).toNumber()] = function (_pat_let81_0) {
            return function (_2652_dt__update__tmp_h4) {
              return function (_pat_let82_0) {
                return function (_2653_dt__update_hcf87_h3) {
                  return _module.D21.create_DC54(_2653_dt__update_hcf87_h3);
                }(_pat_let82_0);
              }(_pat_let_tv76);
            }(_pat_let81_0);
          }(_2642_v70);
          _nw405[(new BigNumber(13)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(14)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(15)).toNumber()] = _module.D21.create_DC54(_2641_v69);
          _nw405[(new BigNumber(16)).toNumber()] = _2642_v70;
          _nw405[(new BigNumber(17)).toNumber()] = ((_this.f3) ? (_module.D21.create_DC54(_2641_v69)) : (_module.D21.create_DC54(_2641_v69)));
          _nw405[(new BigNumber(18)).toNumber()] = _module.D21.create_DC54((_2644_v72)[_module.__default.safeIndex(_this.f7, new BigNumber((_2644_v72).length))]);
          _nw405[(new BigNumber(19)).toNumber()] = _2642_v70;
          _2645_v73 = _nw405;
          let _index410 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_2645_v73).length));
          (_2645_v73)[_index410] = _2642_v70;
          let _2654_v74;
          let _nw406 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _2654_v74 = _nw406;
          let _2655_v75;
          _2655_v75 = _dafny.Seq.of(_2654_v74);
          let _index411 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2636_v64).length));
          let _index412 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_2645_v73).length));
          let _rhs370 = false;
          let _rhs371 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p3, _this.f7));
          let _rhs372 = _2640_v68;
          let _rhs373 = new BigNumber((_2655_v75).length);
          let _rhs374 = _2642_v70;
          let _lhs271 = _2636_v64;
          let _lhs272 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2636_v64).length));
          let _lhs273 = _this;
          let _lhs274 = _2645_v73;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_2645_v73).length));
          r1 = _rhs370;
          r2 = _rhs371;
          _lhs271[_lhs272] = _rhs372;
          _lhs273.f7 = _rhs373;
          _lhs274[_lhs275] = _rhs374;
          let _2656_v76;
          _2656_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_2635_v63).f4);
          _2656_v76 = _2656_v76;
          let _2657_v77;
          _2657_v77 = _dafny.Set.fromElements(!((_this).f4));
          (_2635_v63).f3 = (_2657_v77).IsProperSubsetOf(_2657_v77);
          _2551_v0 = _2551_v0;
        }
        let _2658_v78;
        _2658_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2552_v1).length),p2);
        let _2659_v79;
        _2659_v79 = _module.D11.create_DC25(_2658_v78);
        let _pat_let_tv77 = _2658_v78;
        let _pat_let_tv78 = _2658_v78;
        _2659_v79 = function (_pat_let83_0) {
          return function (_2662_dt__update__tmp_h6) {
            return function (_pat_let86_0) {
              return function (_2663_dt__update_hcf38_h1) {
                return _module.D11.create_DC25(_2663_dt__update_hcf38_h1);
              }(_pat_let86_0);
            }(_pat_let_tv78);
          }(_pat_let83_0);
        }(function (_pat_let84_0) {
          return function (_2660_dt__update__tmp_h5) {
            return function (_pat_let85_0) {
              return function (_2661_dt__update_hcf38_h0) {
                return _module.D11.create_DC25(_2661_dt__update_hcf38_h0);
              }(_pat_let85_0);
            }(_pat_let_tv77);
          }(_pat_let84_0);
        }(_2659_v79));
        let _2664_v80;
        _2664_v80 = _module.D21.create_DC53(_this.f3);
        let _2665_v81;
        _2665_v81 = _dafny.Seq.of((_2664_v80).dtor_cf86, p1, _this.f3, p2);
        let _2666_v82;
        _2666_v82 = _dafny.Set.fromElements((_this).f4);
        let _2667_v83;
        _2667_v83 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_this.f3, (_this).f4));
        let _2668_v84;
        _2668_v84 = _dafny.Seq.of(_2667_v83, _2667_v83);
        let _rhs375 = _module.__default.fm31(p1, new BigNumber((_dafny.Seq.Concat(_2552_v1, p0)).length), false, _2665_v81, globalState);
        let _rhs376 = _2551_v0;
        let _rhs377 = ((((_this).f4) ? (_2667_v83) : ((_2668_v84)[_module.__default.safeIndex(p3, new BigNumber((_2668_v84).length))]))).contains(_2666_v82);
        let _rhs378 = p2;
        let _rhs379 = _module.__default.safeModuloInt(p3, (_this).fm0(globalState));
        let _lhs276 = _this;
        let _lhs277 = globalState;
        _lhs276.f7 = _rhs375;
        _2551_v0 = _rhs376;
        _lhs277.f0 = _rhs377;
        r1 = _rhs378;
        r3 = _rhs379;
        (globalState).f0 = _dafny.Seq.contains(p0, _2551_v0);
      }
      (_this).f7 = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Concat(_2552_v1, p0))).length);
      (_this).m4(globalState);
      let _2669_v85;
      _2669_v85 = _module.D8.create_DC20(_2551_v0, (_this).f4);
      let _2670_v86;
      let _nw407 = new _module.C4();
      _nw407.__ctor((p2) && (true), new BigNumber(-394), (_2669_v85).dtor_cf34, _module.__default.fm6(new BigNumber(54), globalState));
      _2670_v86 = _nw407;
      (_this).f3 = _2670_v86.f8;
      r0 = _this.f3;
      r1 = _2670_v86.f8;
      r2 = p3;
      let _2671_v87;
      let _nw408 = Array((new BigNumber(10)).toNumber()).fill(false);
      _2671_v87 = _nw408;
      let _2672_v88;
      _2672_v88 = _dafny.MultiSet.fromElements(p1, _2670_v86.f8);
      let _2673_v89;
      _2673_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2671_v87,_2672_v88);
      r3 = (_dafny.ZERO).minus(new BigNumber(((((_2673_v89).contains(((p2) ? (_2671_v87) : (_2671_v87)))) ? ((_2673_v89).get(((p2) ? (_2671_v87) : (_2671_v87)))) : ((_2672_v88).Difference(_2672_v88)))).cardinality()));
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let _2674_v0;
      let _nw409 = new _module.C0();
      _nw409.__ctor();
      _2674_v0 = _nw409;
      let _2675_v1;
      _2675_v1 = _module.D2.create_DC7((_dafny.MultiSet.fromElements(_2674_v0, _2674_v0)).update(_2674_v0, _module.__default.abs(_this.f7)));
      let _2676_v2;
      _2676_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2675_v1,(new BigNumber((_dafny.Seq.of(_this.f3)).length)).isLessThanOrEqualTo(_this.f7));
      _2676_v2 = (_2676_v2).update(_2675_v1, (_this).f4);
      let _hi14 = _this.f7;
      for (let _2677_i0 = _this.f7; _2677_i0.isLessThan(_hi14); _2677_i0 = _2677_i0.plus(_dafny.ONE)) {
        let _2678_v3;
        let _nw410 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _2678_v3 = _nw410;
        let _index413 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_2678_v3).length));
        (_2678_v3)[_index413] = _2677_i0;
        let _2679_v4;
        let _nw411 = new _module.C4();
        _nw411.__ctor(_this.f3, new BigNumber((_dafny.Set.fromElements(_2677_i0)).length), !((_this).f4), _this.f3);
        _2679_v4 = _nw411;
        let _2680_v5;
        _2680_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2679_v4,(_2679_v4).f9);
        let _2681_v6;
        _2681_v6 = _dafny.Set.fromElements(_2680_v5);
        let _2682_v7;
        _2682_v7 = _module.D24.create_DC62(_2681_v6);
        let _index414 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_2678_v3).length));
        let _rhs380 = !((_2681_v6).Difference(_2681_v6)).equals((_2682_v7).dtor_cf99);
        let _rhs381 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-804)), function (_2683_i1) {
          return ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_this.f7, new BigNumber(153))).length))).multipliedBy(new BigNumber(-334));
        })).length);
        let _lhs278 = _this;
        let _lhs279 = _2678_v3;
        let _lhs280 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_2678_v3).length));
        _lhs278.f3 = _rhs380;
        _lhs279[_lhs280] = _rhs381;
        let _2684_v8;
        _2684_v8 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2685_v9;
        _2685_v9 = _dafny.MultiSet.fromElements((_2679_v4).f9);
        let _2686_v10;
        _2686_v10 = _dafny.Seq.of(new BigNumber((_2685_v9).cardinality()));
        let _2687_v11;
        _2687_v11 = _dafny.Seq.of(_2679_v4.f8, _2679_v4.f8);
        let _index415 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_2678_v3).length));
        let _rhs382 = (((!(_this.f3)) ? (_2677_i0) : (new BigNumber(793)))).plus((_this.f7).minus(new BigNumber((_2686_v10).length)));
        let _rhs383 = (_this).f4;
        let _rhs384 = _2684_v8;
        let _rhs385 = _module.__default.safeDivisionInt((new BigNumber((_2687_v11).length)).plus(_2677_i0), _2677_i0);
        let _lhs281 = _2678_v3;
        let _lhs282 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_2678_v3).length));
        let _lhs283 = _this;
        let _lhs284 = _this;
        _lhs281[_lhs282] = _rhs382;
        _lhs283.f3 = _rhs383;
        _2684_v8 = _rhs384;
        _lhs284.f7 = _rhs385;
        let _2688_v12;
        let _nw412 = Array((new BigNumber(4)).toNumber()).fill(_module.D16.Default());
        _2688_v12 = _nw412;
        let _2689_v13;
        _2689_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2677_i0);
        let _2690_v14;
        _2690_v14 = _module.D16.create_DC35(_2689_v13);
        let _index416 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2688_v12).length));
        (_2688_v12)[_index416] = _2690_v14;
        let _index417 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2688_v12).length));
        (_2688_v12)[_index417] = _2690_v14;
        (_this).f7 = (_module.D14.create_DC31(new BigNumber(((_module.__default.fm46(globalState)).update(_2679_v4.f8, (_this).f6)).length), new BigNumber((_2687_v11).length))).dtor_cf45;
      }
      let _2691_v15;
      _2691_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f3);
      let _2692_v16;
      let _nw413 = Array((_dafny.ONE).toNumber());
      _nw413[(_dafny.ZERO).toNumber()] = new BigNumber((_2691_v15).length);
      _2692_v16 = _nw413;
      let _2693_v17;
      _2693_v17 = _dafny.MultiSet.fromElements(!(_this.f3));
      let _2694_v18;
      _2694_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_this.f7);
      let _2695_v19;
      _2695_v19 = _dafny.MultiSet.fromElements(_2694_v18);
      let _2696_v20;
      _2696_v20 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f3) ? (_2692_v16) : (_2692_v16)),new BigNumber((_dafny.Set.fromElements(_2693_v17, _dafny.MultiSet.fromElements((_module.D0.create_DC2(_this.f3, new BigNumber(491), new BigNumber((_2691_v15).length))).dtor_cf3, _module.__default.fm21(_this.f7, _2695_v19, globalState), _this.f3), _2693_v17, _2693_v17, (_2693_v17).update(_this.f3, _module.__default.abs(_this.f7)))).length));
      _2696_v20 = (_2696_v20).update(_2692_v16, (_this.f7).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(976)), function (_2697_i2) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length)));
      (_this).f7 = _this.f7;
      let _2698_v21;
      let _nw414 = Array((new BigNumber(6)).toNumber());
      _nw414[(_dafny.ZERO).toNumber()] = _2691_v15;
      _nw414[(_dafny.ONE).toNumber()] = _2691_v15;
      _nw414[(new BigNumber(2)).toNumber()] = _2691_v15;
      _nw414[(new BigNumber(3)).toNumber()] = _2691_v15;
      _nw414[(new BigNumber(4)).toNumber()] = _2691_v15;
      _nw414[(new BigNumber(5)).toNumber()] = _2691_v15;
      _2698_v21 = _nw414;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2698_v21).length))) {
        let _2699_i3 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2699_i3)) && ((_2699_i3).isLessThan(new BigNumber((_2698_v21).length))))) {
          (_2698_v21)[(_2699_i3)] = ((_2691_v15).Merge(_2691_v15)).Merge(_2691_v15);
        }
      }
      let _2700_v22;
      let _nw415 = new _module.C4();
      _nw415.__ctor(_module.__default.fm21(_this.f7, _2695_v19, globalState), new BigNumber(759), ((!(_this.f3)) ? (_this.f3) : ((_this).f4)), _module.__default.fm21(_this.f7, (_2695_v19).update(_2694_v18, _module.__default.abs(_this.f7)), globalState));
      _2700_v22 = _nw415;
      _2700_v22 = _2700_v22;
      return;
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f3 = false;
      this._f4 = false;
      this.f5 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f5, f3, f4) {
      let _this = this;
      (_this).f5 = f5;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-455)))).length);
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(828), new BigNumber((_dafny.Seq.UnicodeFromString("x")).length)), new BigNumber(522)));
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.of(new BigNumber(900), ((!(_this.f3)) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(694))).length))) : (new BigNumber(643))))).length);
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _2701_v0;
      _2701_v0 = _dafny.Seq.of(p2, new BigNumber(-135));
      r1 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-472),_dafny.Seq.Concat(_2701_v0, _2701_v0))).length);
      let _2702_v1;
      _2702_v1 = _dafny.Seq.UnicodeFromString("bxmev");
      let _2703_v2;
      _2703_v2 = _dafny.Seq.of(_2702_v1);
      let _2704_v3;
      let _nw416 = new _module.C5();
      _nw416.__ctor((_2703_v2)[_module.__default.safeIndex(p2, new BigNumber((_2703_v2).length))], (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2702_v1, _2702_v1)).length)), (_this).f4, (_this).f4);
      _2704_v3 = _nw416;
      let _hi15 = p2;
      for (let _2705_i0 = p2; _2705_i0.isLessThan(_hi15); _2705_i0 = _2705_i0.plus(_dafny.ONE)) {
        r1 = (_2701_v0)[_module.__default.safeIndex(_2704_v3.f7, new BigNumber((_2701_v0).length))];
        let _2706_v4;
        let _nw417 = Array((_dafny.ONE).toNumber());
        _nw417[(_dafny.ZERO).toNumber()] = (_this.f3) && (p0);
        _2706_v4 = _nw417;
        let _index418 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length));
        (_2706_v4)[_index418] = _this.f5;
        let _2707_v5;
        _2707_v5 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2708_v6;
        _2708_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f5);
        let _index419 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length));
        let _rhs386 = (_2704_v3).f6;
        let _rhs387 = _module.__default.fm18(new BigNumber(996), _2707_v5, new BigNumber((_2708_v6).length), globalState);
        let _lhs285 = _2706_v4;
        let _lhs286 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length));
        _2702_v1 = _rhs386;
        _lhs285[_lhs286] = _rhs387;
        if ((_this).f4) {
          let _2709_v7;
          let _nw418 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _2709_v7 = _nw418;
          let _index420 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2709_v7).length));
          (_2709_v7)[_index420] = _2705_i0;
          let _index421 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2709_v7).length));
          (_2709_v7)[_index421] = ((p2).minus(p2)).plus(_2704_v3.f7);
          let _2710_v8;
          let _nw419 = Array((new BigNumber(22)).toNumber()).fill([]);
          _2710_v8 = _nw419;
          let _index422 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_2710_v8).length));
          (_2710_v8)[_index422] = _2709_v7;
          let _index423 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_2710_v8).length));
          (_2710_v8)[_index423] = _2709_v7;
          let _index424 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2709_v7).length));
          (_2709_v7)[_index424] = _2704_v3.f7;
          let _2711_v9;
          _2711_v9 = _dafny.Seq.of(_2701_v0, _2701_v0);
          let _2712_v10;
          let _nw420 = new _module.C4();
          _nw420.__ctor((_2705_i0).isLessThan((_this).fm1(!(_this.f5), _2704_v3.f7, p0, _2711_v9, globalState)), (_dafny.ZERO).minus(_2705_i0), true, ((_this.f5) ? (!((_this).f4)) : (false)));
          _2712_v10 = _nw420;
          _2712_v10 = _2712_v10;
          let _2713_v11;
          _2713_v11 = _dafny.Map.Empty.slice().updateUnsafe((_2704_v3.f7).multipliedBy((_2709_v7)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_2709_v7).length))]),_2704_v3.f7);
          _2713_v11 = (_2713_v11).update(_module.__default.safeModuloInt((_2709_v7)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_2709_v7).length))], _2704_v3.f7), _2704_v3.f7);
        } else {
          let _2714_v12;
          let _nw421 = new _module.C2();
          _nw421.__ctor(_this.f3, (_2706_v4)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length))], (_this).f4, _this.f5);
          _2714_v12 = _nw421;
          let _2715_v13;
          _2715_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2714_v12,_2707_v5);
          let _2716_v14;
          _2716_v14 = _dafny.Map.Empty.slice().updateUnsafe((((_2715_v13).contains(_2714_v12)) ? ((_2715_v13).get(_2714_v12)) : (_2707_v5)),(_2714_v12).f10);
          (_this).f5 = !(_dafny.Map.Empty.slice().updateUnsafe(_2707_v5,false)).equals(_2716_v14);
          let _2717_v15;
          let _init79 = ((_2718_v3) => function (_2719_i1) {
            return _module.__default.safeModuloInt(_2719_i1, new BigNumber(((_2718_v3).f6).length));
          })(_2704_v3);
          let _nw422 = Array((new BigNumber(12)).toNumber());
          for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw422.length); _i0_79++) {
            _nw422[_i0_79] = _init79(new BigNumber(_i0_79));
          }
          _2717_v15 = _nw422;
          r0 = _2717_v15;
          let _2720_v16;
          let _nw423 = new _module.C0();
          _nw423.__ctor();
          _2720_v16 = _nw423;
          let _2721_v17;
          _2721_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2720_v16,(_2714_v12).f10);
          let _2722_v18;
          let _nw424 = new _module.C4();
          _nw424.__ctor(((_2706_v4)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length))]) === (p0), new BigNumber((_2721_v17).length), (_2714_v12).f10, (_2714_v12).f10);
          _2722_v18 = _nw424;
          _2722_v18 = _2722_v18;
          let _2723_v19;
          _2723_v19 = _dafny.MultiSet.fromElements(_2705_i0, new BigNumber((_2701_v0).length));
          let _2724_v20;
          let _nw425 = new _module.C3();
          _nw425.__ctor(true, (_2706_v4)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length))]);
          _2724_v20 = _nw425;
          let _2725_v21;
          _2725_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2724_v20,_dafny.MultiSet.fromElements(_2705_i0, _2704_v3.f7, new BigNumber(858)));
          let _2726_v22;
          let _nw426 = new _module.C3();
          _nw426.__ctor((_2706_v4)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_2706_v4).length))], ((((_2725_v21).contains(_2724_v20)) ? ((_2725_v21).get(_2724_v20)) : (_dafny.MultiSet.FromArray(_2701_v0)))).IsSubsetOf(_2723_v19));
          _2726_v22 = _nw426;
          let _2727_v23;
          _2727_v23 = _dafny.Map.Empty.slice().updateUnsafe((_2714_v12).f10,(_2701_v0)[_module.__default.safeIndex(p2, new BigNumber((_2701_v0).length))]);
          _2727_v23 = (_2727_v23).update((p0) || ((_2722_v18).f4), new BigNumber(868));
        }
        let _rhs388 = (_2704_v3).f6;
        let _rhs389 = _this.f3;
        let _lhs287 = globalState;
        _2702_v1 = _rhs388;
        _lhs287.f0 = _rhs389;
      }
      let _2728_v24;
      _2728_v24 = _dafny.Set.fromElements(_2704_v3.f7, _2704_v3.f7, _2704_v3.f7, new BigNumber((_dafny.Seq.UnicodeFromString("dnsnspud")).length));
      let _2729_v25;
      _2729_v25 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(228)).multipliedBy(new BigNumber((_2728_v24).length)),(_dafny.ZERO).minus(p2));
      _2729_v25 = _2729_v25;
      r1 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_2704_v3.f7, p2), (p2).minus(_2704_v3.f7));
      let _2730_v26;
      let _nw427 = Array((_dafny.ONE).toNumber());
      _nw427[(_dafny.ZERO).toNumber()] = p2;
      _2730_v26 = _nw427;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2730_v26).length))) {
        let _2731_i2 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2731_i2)) && ((_2731_i2).isLessThan(new BigNumber((_2730_v26).length))))) {
          (_2730_v26)[(_2731_i2)] = (_2731_i2).multipliedBy(_2704_v3.f7);
        }
      }
      let _init80 = ((_2732_v0) => function (_2733_i3) {
        return (_2733_i3).multipliedBy(new BigNumber((_2732_v0).length));
      })(_2701_v0);
      let _nw428 = Array((new BigNumber(11)).toNumber());
      for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw428.length); _i0_80++) {
        _nw428[_i0_80] = _init80(new BigNumber(_i0_80));
      }
      r0 = _nw428;
      r1 = (_dafny.ZERO).minus(_2704_v3.f7);
      return [r0, r1];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2734_v0;
      _2734_v0 = _dafny.Set.fromElements(p0, p0);
      let _2735_v1;
      _2735_v1 = _module.D16.create_DC36(_2734_v0, (p1).f4, p1.f3, (p1).f4);
      let _pat_let_tv79 = p0;
      let _pat_let_tv80 = p1;
      let _2736_v2;
      _2736_v2 = _dafny.Set.fromElements(function (_pat_let87_0) {
        return function (_2737_dt__update__tmp_h0) {
          return function (_pat_let88_0) {
            return function (_2738_dt__update_hcf51_h0) {
              return function (_pat_let89_0) {
                return function (_2739_dt__update_hcf53_h0) {
                  return function (_pat_let90_0) {
                    return function (_2740_dt__update_hcf52_h0) {
                      return _module.D16.create_DC36(_2738_dt__update_hcf51_h0, _2740_dt__update_hcf52_h0, _2739_dt__update_hcf53_h0, (_2737_dt__update__tmp_h0).dtor_cf54);
                    }(_pat_let90_0);
                  }((_pat_let_tv80).f4);
                }(_pat_let89_0);
              }(_this.f3);
            }(_pat_let88_0);
          }(_dafny.Set.fromElements(_pat_let_tv79));
        }(_pat_let87_0);
      }(_2735_v1));
      let _2741_v3;
      _2741_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2736_v2,false);
      (p1).f3 = (((_2741_v3).contains(_2736_v2)) ? ((_2741_v3).get(_2736_v2)) : (p1.f3));
      let _2742_v4;
      _2742_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32(p1.f3, p1.f3, false, p1.f3, globalState),((p1.f3) ? (p0) : (new BigNumber(746))));
      let _2743_v5;
      _2743_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_this.f5, globalState),p1);
      let _2744_v6;
      _2744_v6 = _dafny.Seq.UnicodeFromString("wtokh");
      let _2745_v7;
      _2745_v7 = _dafny.MultiSet.fromElements(p0, new BigNumber((_2743_v5).length), new BigNumber((_2744_v6).length));
      let _2746_v8;
      _2746_v8 = _dafny.Seq.of(new BigNumber((_2744_v6).length), p0);
      let _2747_v9;
      _2747_v9 = _module.D20.create_DC48(_2746_v8);
      let _2748_v10;
      _2748_v10 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2749_v11;
      _2749_v11 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-105)), function (_2750_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.update(_2744_v6, _module.__default.safeIndex(p0, new BigNumber((_2744_v6).length)), _2748_v10));
      _2742_v4 = (_dafny.Map.Empty.slice().updateUnsafe(_2734_v0,p0)).Merge(_module.__default.fm47(new BigNumber((_2745_v7).cardinality()), _this.f3, _2747_v9, (_2749_v11)[_module.__default.safeIndex(p0, new BigNumber((_2749_v11).length))], globalState));
      _2744_v6 = _dafny.Seq.update(_2744_v6, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2744_v6).length)), _2748_v10);
      let _2751_i1;
      _2751_i1 = _dafny.ZERO;
      L9: {
        while (p1.f3) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2751_i1)) {
              break L9;
            }
            _2751_i1 = (_2751_i1).plus(_dafny.ONE);
            let _2752_v12;
            _2752_v12 = _dafny.Seq.of(_this.f5);
            _2752_v12 = _dafny.Seq.Concat(_2752_v12, _2752_v12);
            if ((p1).f4) {
              let _2753_v13;
              _2753_v13 = _dafny.MultiSet.fromElements(_2748_v10, _2748_v10, _2748_v10, _2748_v10, _2748_v10);
              let _2754_v14;
              let _nw429 = new _module.C4();
              _nw429.__ctor(!((p1).f4) || ((p1).f4), (((_2753_v13).contains(_2748_v10)) ? ((_2753_v13).get(_2748_v10)) : (p0)), (p1).f4, _this.f3);
              _2754_v14 = _nw429;
              let _2755_v15;
              let _nw430 = new _module.C3();
              _nw430.__ctor(p1.f3, p1.f3);
              _2755_v15 = _nw430;
              let _2756_v16;
              _2756_v16 = _dafny.MultiSet.fromElements(true);
              let _2757_v17;
              _2757_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2756_v16,!(_module.__default.fm4((p1).f4, globalState)));
              _2757_v17 = (_2757_v17).Merge(_2757_v17);
              let _2758_v18;
              _2758_v18 = _dafny.Seq.of(_2752_v12);
              let _2759_v19;
              let _nw431 = Array((new BigNumber(15)).toNumber());
              _nw431[(_dafny.ZERO).toNumber()] = _2752_v12;
              _nw431[(_dafny.ONE).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f4, true, (p1).f4, p1.f3, true);
              _nw431[(new BigNumber(3)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(4)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2752_v12, _2752_v12);
              _nw431[(new BigNumber(6)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(7)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(8)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2752_v12, _2752_v12);
              _nw431[(new BigNumber(10)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_2752_v12, (_2758_v18)[_module.__default.safeIndex(p0, new BigNumber((_2758_v18).length))]), _module.__default.safeIndex(new BigNumber((_2752_v12).length), new BigNumber((_dafny.Seq.Concat(_2752_v12, (_2758_v18)[_module.__default.safeIndex(p0, new BigNumber((_2758_v18).length))])).length)), p1.f3);
              _nw431[(new BigNumber(12)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(13)).toNumber()] = _2752_v12;
              _nw431[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(false);
              _2759_v19 = _nw431;
              let _index425 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2759_v19).length));
              (_2759_v19)[_index425] = _dafny.Seq.Concat(_2752_v12, _2752_v12);
              let _index426 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2759_v19).length));
              (_2759_v19)[_index426] = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm34(p0, (_2754_v14).f9, globalState), _2752_v12), _2752_v12);
              r0 = _module.__default.safeModuloInt(p0, (p0).multipliedBy(new BigNumber((_2744_v6).length)));
            } else {
              let _2760_v20;
              _2760_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2752_v12).length),(p1).f4);
              let _2761_v21;
              let _2762_v22;
              let _out47;
              let _out48;
              let _outcollector17 = (p1).m0((p1).f4, _2760_v20, _module.__default.safeDivisionInt(p0, p0), globalState);
              _out47 = _outcollector17[0];
              _out48 = _outcollector17[1];
              _2761_v21 = _out47;
              _2762_v22 = _out48;
              let _2763_v23;
              _2763_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2762_v22,new BigNumber(-272));
              let _2764_v24;
              _2764_v24 = _dafny.MultiSet.fromElements(_this.f3);
              (globalState).f0 = (_module.D0.create_DC2(p1.f3, new BigNumber(-562), new BigNumber(((_2763_v23).update(new BigNumber((_2746_v8).length), new BigNumber((_2764_v24).cardinality()))).length))).dtor_cf3;
              let _index427 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length));
              (_2761_v21)[_index427] = _2762_v22;
              let _2765_v25;
              _2765_v25 = _dafny.Seq.of(_2746_v8);
              let _2766_v26;
              _2766_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_2765_v25);
              let _2767_v27;
              _2767_v27 = (((_2766_v26).contains(p0)) ? ((_2766_v26).get(p0)) : (_2765_v25));
              let _index428 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length));
              let _rhs390 = (_this).fm1(_module.__default.fm4(p1.f3, globalState), _2762_v22, (((_2760_v20).contains(_2762_v22)) ? ((_2760_v20).get(_2762_v22)) : (p1.f3)), (_2767_v27), globalState);
              let _rhs391 = (_dafny.ZERO).minus(((((p1).f4) ? ((_dafny.ZERO).minus(p0)) : (new BigNumber((_2746_v8).length)))).minus(_2762_v22));
              let _lhs288 = _2761_v21;
              let _lhs289 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length));
              _2762_v22 = _rhs390;
              _lhs288[_lhs289] = _rhs391;
              let _2768_v28;
              _2768_v28 = _module.D17.create_DC39(_2761_v21, p1.f3, p1.f3, p0, (_2761_v21)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length))]);
              let _index429 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length));
              (_2761_v21)[_index429] = (_2768_v28).dtor_cf60;
              let _index430 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_2761_v21).length));
              (_2761_v21)[_index430] = p0;
            }
            _2746_v8 = _2746_v8;
            let _2769_v29;
            _2769_v29 = _dafny.MultiSet.fromElements(_module.__default.fm6(p0, globalState), _this.f3, p1.f3, _this.f3);
            (_this).f5 = (_2769_v29).IsDisjointFrom(_dafny.MultiSet.fromElements(_this.f3));
          }
        }
      }
      let _2770_v30;
      let _nw432 = new _module.C2();
      _nw432.__ctor((_this).f4, _this.f5, p1.f3, !(_this.f3) || (p1.f3));
      _2770_v30 = _nw432;
      let _2771_v31;
      let _nw433 = new _module.C3();
      _nw433.__ctor(_this.f5, (_this).f4);
      _2771_v31 = _nw433;
      let _2772_v32;
      _2772_v32 = _module.D22.create_DC55(_2771_v31);
      let _pat_let_tv81 = _2771_v31;
      let _pat_let_tv82 = _2771_v31;
      let _source33 = function (_pat_let91_0) {
        return function (_2775_dt__update__tmp_h2) {
          return function (_pat_let94_0) {
            return function (_2776_dt__update_hcf88_h1) {
              return _module.D22.create_DC55(_2776_dt__update_hcf88_h1);
            }(_pat_let94_0);
          }(_pat_let_tv82);
        }(_pat_let91_0);
      }(function (_pat_let92_0) {
        return function (_2773_dt__update__tmp_h1) {
          return function (_pat_let93_0) {
            return function (_2774_dt__update_hcf88_h0) {
              return _module.D22.create_DC55(_2774_dt__update_hcf88_h0);
            }(_pat_let93_0);
          }(_pat_let_tv81);
        }(_pat_let92_0);
      }(_2772_v32));
      if (_source33.is_DC56) {
        let _2777___mcc_h0 = (_source33).cf89;
        let _2778___mcc_h1 = (_source33).cf90;
        let _2779_cf90 = _2778___mcc_h1;
        let _2780_cf89 = _2777___mcc_h0;
        let _2781_v33;
        let _nw434 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _2781_v33 = _nw434;
        let _2782_v34;
        _2782_v34 = _dafny.Seq.of(_2781_v33);
        let _2783_v35;
        _2783_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2748_v10,(_2782_v34)[_module.__default.safeIndex(_2780_cf89, new BigNumber((_2782_v34).length))]);
        let _2784_v36;
        let _nw435 = Array((new BigNumber(20)).toNumber());
        _nw435[(_dafny.ZERO).toNumber()] = _2781_v33;
        _nw435[(_dafny.ONE).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(2)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(3)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(4)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(5)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(6)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(7)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(8)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(9)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(10)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(11)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(12)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(13)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(14)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(15)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(16)).toNumber()] = (((_2783_v35).contains(_2748_v10)) ? ((_2783_v35).get(_2748_v10)) : (_2781_v33));
        _nw435[(new BigNumber(17)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(18)).toNumber()] = _2781_v33;
        _nw435[(new BigNumber(19)).toNumber()] = _2781_v33;
        _2784_v36 = _nw435;
        let _index431 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length));
        (_2784_v36)[_index431] = _2781_v33;
        let _2785_v37;
        let _nw436 = new _module.C0();
        _nw436.__ctor();
        _2785_v37 = _nw436;
        let _2786_v38;
        let _nw437 = Array((new BigNumber(13)).toNumber());
        _nw437[(_dafny.ZERO).toNumber()] = _2785_v37;
        _nw437[(_dafny.ONE).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(2)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(3)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(4)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(5)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(6)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(7)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(8)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(9)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(10)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(11)).toNumber()] = _2785_v37;
        _nw437[(new BigNumber(12)).toNumber()] = _2785_v37;
        _2786_v38 = _nw437;
        let _2787_v39;
        _2787_v39 = _dafny.Map.Empty.slice().updateUnsafe((_2770_v30).f11,_2786_v38);
        let _2788_v40;
        _2788_v40 = _module.D19.create_DC44((_2787_v39).Merge((_dafny.Map.Empty.slice().updateUnsafe(p1.f3,_2786_v38)).update((p1).f4, _2786_v38)));
        let _2789_v41;
        let _nw438 = Array((new BigNumber(5)).toNumber()).fill(false);
        _2789_v41 = _nw438;
        let _index432 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_2789_v41).length));
        (_2789_v41)[_index432] = (_2770_v30).f11;
        let _2790_v42;
        _2790_v42 = _module.D0.create_DC0(new BigNumber(335));
        let _2791_v43;
        _2791_v43 = _dafny.Set.fromElements((_2770_v30).f11, p1.f3);
        let _index433 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length));
        let _index434 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_2789_v41).length));
        let _rhs392 = _2781_v33;
        let _rhs393 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_2780_cf89, _2780_cf89));
        let _rhs394 = ((_module.__default.fm18(p0, _2748_v10, (_2790_v42).dtor_cf0, globalState)) ? (new BigNumber((_2791_v43).length)) : (new BigNumber((_dafny.Set.fromElements(_module.__default.fm4((p1).f4, globalState), (p1).f4)).length)));
        let _rhs395 = _2788_v40;
        let _rhs396 = (_this).f4;
        let _lhs290 = _2784_v36;
        let _lhs291 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length));
        let _lhs292 = _2789_v41;
        let _lhs293 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_2789_v41).length));
        _lhs290[_lhs291] = _rhs392;
        _2745_v7 = _rhs393;
        r0 = _rhs394;
        _2788_v40 = _rhs395;
        _lhs292[_lhs293] = _rhs396;
        let _2792_v44;
        let _init81 = ((_2793_v10) => function (_2794_i2) {
          return _2793_v10;
        })(_2748_v10);
        let _nw439 = Array((new BigNumber(24)).toNumber());
        for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw439.length); _i0_81++) {
          _nw439[_i0_81] = _init81(new BigNumber(_i0_81));
        }
        _2792_v44 = _nw439;
        let _index435 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_2792_v44).length));
        (_2792_v44)[_index435] = _2748_v10;
        let _index436 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_2792_v44).length));
        (_2792_v44)[_index436] = _2748_v10;
        let _index437 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_2792_v44).length));
        (_2792_v44)[_index437] = _2748_v10;
        let _arr0 = (_2784_v36)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length))];
        let _index438 = _module.__default.safeIndex(new BigNumber(209), new BigNumber(((_2784_v36)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length))]).length));
        _arr0[_index438] = new BigNumber((_2746_v8).length);
        let _arr1 = (_2784_v36)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length))];
        let _index439 = _module.__default.safeIndex(new BigNumber(209), new BigNumber(((_2784_v36)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_2784_v36).length))]).length));
        _arr1[_index439] = _2780_cf89;
      } else if (_source33.is_DC55) {
        let _2795___mcc_h2 = (_source33).cf88;
        let _2796_cf88 = _2795___mcc_h2;
        let _2797_v45;
        let _init82 = ((_2798_v8, _2799_p0, _2800_v30, _2801_p1) => function (_2802_i3) {
          return (_module.D18.create_DC43((_2798_v8)[_module.__default.safeIndex(_2799_p0, new BigNumber((_2798_v8).length))], (_module.D19.create_DC45(_2799_p0, new BigNumber(167), (_2800_v30).f10, _this.f5, _dafny.Seq.of(_this.f5))).dtor_cf70, (_2801_p1).f4)).dtor_cf67;
        })(_2746_v8, p0, _2770_v30, p1);
        let _nw440 = Array((new BigNumber(18)).toNumber());
        for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw440.length); _i0_82++) {
          _nw440[_i0_82] = _init82(new BigNumber(_i0_82));
        }
        _2797_v45 = _nw440;
        _2797_v45 = _2797_v45;
        let _2803_v46;
        _2803_v46 = _dafny.Seq.of((_2770_v30).f10, (_this).f4, (p1).f4, !(_this.f5), (p1).f4);
        let _2804_v47;
        _2804_v47 = _module.D18.create_DC43(_module.__default.fm31(_this.f3, new BigNumber((_dafny.Seq.UnicodeFromString("wqlotnfij")).length), _this.f5, _2803_v46, globalState), (p0).minus(p0), _this.f3);
        let _2805_v48;
        let _nw441 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _2805_v48 = _nw441;
        let _2806_v49;
        _2806_v49 = _module.D17.create_DC39(_2805_v48, _this.f5, (_2770_v30).f11, p0, p0);
        let _pat_let_tv83 = _2770_v30;
        let _rhs397 = !((function (_pat_let95_0) {
          return function (_2807_dt__update__tmp_h3) {
            return function (_pat_let96_0) {
              return function (_2808_dt__update_hcf58_h0) {
                return function (_pat_let97_0) {
                  return function (_2809_dt__update_hcf59_h0) {
                    return _module.D17.create_DC39((_2807_dt__update__tmp_h3).dtor_cf57, _2808_dt__update_hcf58_h0, _2809_dt__update_hcf59_h0, (_2807_dt__update__tmp_h3).dtor_cf60, (_2807_dt__update__tmp_h3).dtor_cf61);
                  }(_pat_let97_0);
                }((_this).f4);
              }(_pat_let96_0);
            }((_pat_let_tv83).f10);
          }(_pat_let95_0);
        }(_2806_v49)).dtor_cf58);
        let _rhs398 = _2804_v47;
        let _lhs294 = p1;
        _lhs294.f3 = _rhs397;
        _2804_v47 = _rhs398;
        let _index440 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_2797_v45).length));
        (_2797_v45)[_index440] = p1.f3;
        let _2810_v50;
        _2810_v50 = _dafny.MultiSet.fromElements(_2744_v6);
        let _2811_v51;
        _2811_v51 = _dafny.Map.Empty.slice().updateUnsafe((_2770_v30).f11,new BigNumber(665));
        let _2812_v52;
        _2812_v52 = _module.D1.create_DC5(p1.f3, (p1).f4, new BigNumber(((_2810_v50).update(_dafny.Seq.UnicodeFromString("fiktmwl"), _module.__default.abs(new BigNumber((_2811_v51).length)))).cardinality()), p0);
        let _2813_v53;
        let _nw442 = Array((new BigNumber(8)).toNumber());
        _nw442[(_dafny.ZERO).toNumber()] = _2812_v52;
        _nw442[(_dafny.ONE).toNumber()] = _module.D1.create_DC5(!(true), (_2770_v30).f11, new BigNumber(-544), p0);
        _nw442[(new BigNumber(2)).toNumber()] = _module.D1.create_DC5(p1.f3, p1.f3, p0, p0);
        _nw442[(new BigNumber(3)).toNumber()] = _2812_v52;
        _nw442[(new BigNumber(4)).toNumber()] = _module.__default.fm48((_2770_v30).f10, globalState);
        _nw442[(new BigNumber(5)).toNumber()] = _2812_v52;
        _nw442[(new BigNumber(6)).toNumber()] = _2812_v52;
        _nw442[(new BigNumber(7)).toNumber()] = _2812_v52;
        _2813_v53 = _nw442;
        let _index441 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_2797_v45).length));
        let _rhs399 = (p1).f4;
        let _rhs400 = ((((_2770_v30).f11) || (true)) ? (_2813_v53) : (_2813_v53));
        let _rhs401 = (p1).f4;
        let _lhs295 = _2797_v45;
        let _lhs296 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_2797_v45).length));
        let _lhs297 = p1;
        _lhs295[_lhs296] = _rhs399;
        _2813_v53 = _rhs400;
        _lhs297.f3 = _rhs401;
        r0 = (((_2811_v51).contains((_2803_v46)[_module.__default.safeIndex(p0, new BigNumber((_2803_v46).length))])) ? ((_2811_v51).get((_2803_v46)[_module.__default.safeIndex(p0, new BigNumber((_2803_v46).length))])) : ((p0).plus(p0)));
      } else {
        let _2814___mcc_h3 = (_source33).cf91;
        let _2815_cf91 = _2814___mcc_h3;
        (p1).f3 = (p1).f4;
        r0 = (_dafny.ZERO).minus((new BigNumber(837)).plus(p0));
        _2748_v10 = _2748_v10;
        let _2816_v54;
        _2816_v54 = _dafny.Seq.of((_2770_v30).f10, (_2770_v30).f11);
        let _2817_v55;
        _2817_v55 = _dafny.MultiSet.fromElements(_2816_v54);
        let _2818_v56;
        _2818_v56 = _module.D19.create_DC45(new BigNumber((_2744_v6).length), (_2746_v8)[_module.__default.safeIndex(new BigNumber((_2817_v55).cardinality()), new BigNumber((_2746_v8).length))], true, false, _2816_v54);
        (globalState).f0 = !((((_this.f3) ? (_2818_v56) : (_2818_v56))).dtor_cf72);
      }
      r0 = p0;
      return r0;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _2819_v0;
      _2819_v0 = new BigNumber(304);
      let _2820_v1;
      _2820_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2819_v0,(p1).f4);
      let _2821_v2;
      _2821_v2 = _dafny.MultiSet.fromElements(p1.f3);
      let _2822_v3;
      _2822_v3 = _dafny.Seq.of(_2819_v0);
      let _2823_v4;
      _2823_v4 = _dafny.Seq.UnicodeFromString("oq");
      let _2824_v5;
      _2824_v5 = _dafny.Seq.of(_dafny.Seq.of((_module.D19.create_DC46(_2819_v0, (p1).f4, (p1).f4, _module.D1.create_DC5((p1).f4, (_this).f4, new BigNumber(119), _2819_v0), new BigNumber(795))).dtor_cf74));
      let _2825_v6;
      let _2826_v7;
      let _out49;
      let _out50;
      let _outcollector18 = (p1).m0(_this.f5, _2820_v1, new BigNumber(((_2821_v2).update(p1.f3, _module.__default.abs((_dafny.ZERO).minus((_2822_v3)[_module.__default.safeIndex((_this).fm1(true, new BigNumber((_2823_v4).length), _this.f5, _2824_v5, globalState), new BigNumber((_2822_v3).length))])))).cardinality()), globalState);
      _out49 = _outcollector18[0];
      _out50 = _outcollector18[1];
      _2825_v6 = _out49;
      _2826_v7 = _out50;
      let _2827_v8;
      _2827_v8 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_2821_v2).cardinality())).isLessThanOrEqualTo(_2826_v7),_this.f5);
      (_this).f3 = (((_2827_v8).contains((((_2827_v8).contains((_this).f4)) ? ((_2827_v8).get((_this).f4)) : (_this.f5)))) ? ((_2827_v8).get((((_2827_v8).contains((_this).f4)) ? ((_2827_v8).get((_this).f4)) : (_this.f5)))) : (_this.f5));
      let _2828_v9;
      _2828_v9 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), ((_2829_p0) => function (_2830_i0) {
        return _2829_p0;
      })(p0)), _2823_v4, _module.__default.fm30(_module.D1.create_DC5(!(_this.f3), _this.f5, _2819_v0, (_this).fm1((p1).f4, new BigNumber((_dafny.Seq.of(_this.f5)).length), p1.f3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(383)), ((_2831_v0) => function (_2832_i1) {
  return _dafny.Seq.of(_2831_v0);
})(_2819_v0)), globalState)), _2826_v7, _2826_v7, globalState), _2823_v4);
      let _2833_v10;
      _2833_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2819_v0,_2822_v3);
      let _2834_v11;
      _2834_v11 = _module.D17.create_DC39(_2825_v6, (_this).f4, _this.f3, _2826_v7, _2826_v7);
      let _2835_v12;
      _2835_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1.f3,_2834_v11);
      let _nw443 = Array((new BigNumber(13)).toNumber());
      _nw443[(_dafny.ZERO).toNumber()] = _2826_v7;
      _nw443[(_dafny.ONE).toNumber()] = ((!((_this).f4)) ? (_2826_v7) : (new BigNumber(-46)));
      _nw443[(new BigNumber(2)).toNumber()] = (((_2828_v9).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-891)), ((_2838_p0) => function (_2839_i2) {
        return _2838_p0;
      })(p0)))) ? ((_2828_v9).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-891)), ((_2836_p0) => function (_2837_i2) {
        return _2836_p0;
      })(p0)))) : (_2826_v7));
      _nw443[(new BigNumber(3)).toNumber()] = new BigNumber(((((_2833_v10).contains(_2819_v0)) ? ((_2833_v10).get(_2819_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(630)), ((_2840_v0) => function (_2841_i3) {
        return _2840_v0;
      })(_2819_v0))))).length);
      _nw443[(new BigNumber(4)).toNumber()] = new BigNumber(((_2835_v12).update((_this).f4, _2834_v11)).length);
      _nw443[(new BigNumber(5)).toNumber()] = _2819_v0;
      _nw443[(new BigNumber(6)).toNumber()] = _2819_v0;
      _nw443[(new BigNumber(7)).toNumber()] = _2819_v0;
      _nw443[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_2819_v0, (_dafny.ZERO).minus(_2819_v0));
      _nw443[(new BigNumber(9)).toNumber()] = new BigNumber((_2823_v4).length);
      _nw443[(new BigNumber(10)).toNumber()] = _2819_v0;
      _nw443[(new BigNumber(11)).toNumber()] = _2826_v7;
      _nw443[(new BigNumber(12)).toNumber()] = _2826_v7;
      _2825_v6 = _nw443;
      let _hi16 = (_2826_v7).minus(_2819_v0);
      for (let _2842_i4 = _2819_v0; _2842_i4.isLessThan(_hi16); _2842_i4 = _2842_i4.plus(_dafny.ONE)) {
        _2819_v0 = _2842_i4;
        _2826_v7 = _2819_v0;
        let _2843_v13;
        _2843_v13 = _dafny.Seq.of((_this).f4);
        _2821_v2 = _module.__default.fm20(_dafny.areEqual(_dafny.Seq.of((_this).f4), _2843_v13), globalState);
        _2843_v13 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2843_v13, _dafny.Seq.of(_this.f5, p1.f3)), ((p1.f3) ? (_2843_v13) : (_2843_v13)));
      }
      let _2844_i5;
      _2844_i5 = _dafny.ZERO;
      L10: {
        while (_this.f5) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2844_i5)) {
              break L10;
            }
            _2844_i5 = (_2844_i5).plus(_dafny.ONE);
            (p1).f3 = (p1).f4;
            let _index442 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2825_v6).length));
            (_2825_v6)[_index442] = (new BigNumber(23)).plus(_2826_v7);
            let _index443 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2825_v6).length));
            (_2825_v6)[_index443] = new BigNumber((function () {
              let _coll87 = new _dafny.Set();
              for (const _compr_87 of (_2820_v1).Keys.Elements) {
                let _2845_v14 = _compr_87;
                if ((_2820_v1).contains(_2845_v14)) {
                  _coll87.add((_2845_v14).multipliedBy((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('w'.codePointAt(0)))).length))))).length)).multipliedBy((_module.D1.create_DC5(!(true), !(true), new BigNumber(803), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).dtor_cf11)));
                }
              }
              return _coll87;
            }()).length);
            _2819_v0 = _2826_v7;
            (_this).f5 = false;
          }
        }
      }
      let _2846_v15;
      let _2847_v16;
      let _out51;
      let _out52;
      let _outcollector19 = (p1).m0(p1.f3, _2820_v1, _2819_v0, globalState);
      _out51 = _outcollector19[0];
      _out52 = _outcollector19[1];
      _2846_v15 = _out51;
      _2847_v16 = _out52;
      r0 = _dafny.Seq.Concat(_2823_v4, _2823_v4);
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
