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
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (!((_dafny.MultiSet.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(true, true))).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(true))))) === (true);
    };
    static fm2(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D27.create_DC66(_dafny.Seq.of(false), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length));
      if (_source0.is_DC65) {
        let _0___mcc_h0 = (_source0).cf107;
        let _1___mcc_h1 = (_source0).cf108;
        let _2___mcc_h2 = (_source0).cf109;
        let _3___mcc_h3 = (_source0).cf110;
        let _4_cf110 = _3___mcc_h3;
        let _5_cf109 = _2___mcc_h2;
        let _6_cf108 = _1___mcc_h1;
        let _7_cf107 = _0___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_7_cf107,_4_cf110);
      } else if (_source0.is_DC66) {
        let _8___mcc_h4 = (_source0).cf111;
        let _9___mcc_h5 = (_source0).cf112;
        let _10_cf112 = _9___mcc_h5;
        let _11_cf111 = _8___mcc_h4;
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(true,false);
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(true,true);
        }
      } else if (_source0.is_DC67) {
        let _12___mcc_h6 = (_source0).cf113;
        let _13___mcc_h7 = (_source0).cf114;
        let _14___mcc_h8 = (_source0).cf115;
        let _15_cf115 = _14___mcc_h8;
        let _16_cf114 = _13___mcc_h7;
        let _17_cf113 = _12___mcc_h6;
        return (_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
      } else {
        let _18___mcc_h9 = (_source0).cf106;
        let _19_cf106 = _18___mcc_h9;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),false));
      }
    };
    static fm11(globalState) {
      let _source1 = _module.D5.create_DC12();
      if (_source1.is_DC12) {
        return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of (_dafny.Set.fromElements(_module.D8.create_DC18(_dafny.MultiSet.fromElements(new BigNumber(-307), new BigNumber((_dafny.Seq.of(false)).length))))).Elements) {
            let _20_v0 = _compr_0;
            if ((_dafny.Set.fromElements(_module.D8.create_DC18(_dafny.MultiSet.fromElements(new BigNumber(-307), new BigNumber((_dafny.Seq.of(false)).length))))).contains(_20_v0)) {
              _coll0.push([_20_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(558))).length))).cardinality()),new BigNumber(156))).length))).length))]);
            }
          }
          return _coll0;
        }()).length)));
      } else if (_source1.is_DC13) {
        let _21___mcc_h0 = (_source1).cf16;
        let _22___mcc_h1 = (_source1).cf17;
        let _23___mcc_h2 = (_source1).cf18;
        let _24_cf18 = _23___mcc_h2;
        let _25_cf17 = _22___mcc_h1;
        let _26_cf16 = _21___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber(646), _25_cf17);
      } else {
        let _27___mcc_h3 = (_source1).cf15;
        let _28_cf15 = _27___mcc_h3;
        return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(688))).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), (_dafny.ZERO).minus(new BigNumber(((_module.D5.create_DC13(_dafny.Seq.UnicodeFromString("pbfo"), new BigNumber(514), false)).dtor_cf16).length)), new BigNumber(832))));
      }
    };
    static fm12(globalState) {
      return _dafny.Seq.UnicodeFromString("bt");
    };
    static fm13(p0, p1, p2, globalState) {
      return new BigNumber(369);
    };
    static fm14(globalState) {
      return (new BigNumber(((_dafny.MultiSet.fromElements(true, false, true)).Union(_dafny.MultiSet.fromElements(true))).cardinality())).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(164), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(572)), function (_29_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(891)),!(false));
      })).length))).length));
    };
    static fm15(p0, p1, p2, globalState) {
      if (((true) ? (true) : (false))) {
        return (_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(true));
      } else {
        return _dafny.MultiSet.fromElements(true);
      }
    };
    static fm16(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("xxrlr");
    };
    static fm19(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
    };
    static fm20(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false, true, true, false, !(true)))));
    };
    static fm21(p0, p1, p2, globalState) {
      let _source2 = function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(267))).cardinality()), new BigNumber(946), new BigNumber(-175)))).Elements) {
          let _30_v0 = _compr_1;
          if ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(267))).cardinality()), new BigNumber(946), new BigNumber(-175)))).contains(_30_v0)) {
            _coll1.push([_30_v0,true]);
          }
        }
        return _coll1;
      }();
      let _31___mcc_h0 = _source2;
      let _32_cf94 = _31___mcc_h0;
      return _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false));
    };
    static fm22(p0, p1, p2, globalState) {
      if (!(!(!(true)) || (!(false)))) {
        return (function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-471), new BigNumber(969))) {
            let _33_v0 = _compr_2;
            if (((new BigNumber(-471)).isLessThanOrEqualTo(_33_v0)) && ((_33_v0).isLessThan(new BigNumber(969)))) {
              _coll2.add(_module.__default.safeDivisionInt(_33_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), function (_34_i0) {
                return new BigNumber(917);
              })).length)));
            }
          }
          return _coll2;
        }()).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), function (_35_i1) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), function (_36_i2) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(576))).length))).length))).length)));
      } else {
        return _dafny.Set.fromElements(new BigNumber(168), new BigNumber(-164));
      }
    };
    static fm23(p0, globalState) {
      return _dafny.MultiSet.fromElements(((true) ? (new _dafny.CodePoint('x'.codePointAt(0))) : (new _dafny.CodePoint('o'.codePointAt(0)))));
    };
    static fm24(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gtft"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-262)), function (_37_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("lpf")));
    };
    static fm25(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-921),_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rlrlyjs")).length)),new BigNumber((_dafny.Seq.UnicodeFromString("odsdhie")).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("lxobxhhst")).length),function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(611), new BigNumber(962))) {
          let _38_v0 = _compr_3;
          if (((new BigNumber(611)).isLessThanOrEqualTo(_38_v0)) && ((_38_v0).isLessThan(new BigNumber(962)))) {
            _coll3.push([(_38_v0).multipliedBy(new BigNumber(283)),new BigNumber(331)]);
          }
        }
        return _coll3;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qhycailwn")).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nimexgr")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hosqaqqwm")).length),true)).length),new BigNumber(964))).length)))));
    };
    static fm26(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true, !(true)))).Union(((false) ? (_dafny.Set.fromElements(true, false, !(!(!(true))))) : (_dafny.Set.fromElements(false))));
    };
    static fm27(globalState) {
      return _module.D13.create_DC29(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("t"), _dafny.Seq.UnicodeFromString("iusforw")), false);
    };
    static fm28(globalState) {
      return ((_dafny.MultiSet.fromElements(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(false, true),true)).Keys.Elements) {
          let _39_v0 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(false, true),true)).contains(_39_v0)) {
            _coll4.push([_39_v0,true]);
          }
        }
        return _coll4;
      }())).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(false, false),false), _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(true, false),true), _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(false, true),true)))).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(true, false),false)));
    };
    static fm29(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(true)))),new BigNumber(464))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(856)), function (_40_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(221),_dafny.Seq.of(true, false))).length);
      }))).cardinality())))));
    };
    static fm30(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(228))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(650),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(169)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("qshnukah")).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_module.D23.create_DC56(false, true, true, false)).dtor_cf92),false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(116))).cardinality())), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-65))).length), new BigNumber(-43), new BigNumber(249)));
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-879), new BigNumber(774), new BigNumber(-161), new BigNumber(-747)), _dafny.Seq.of(new BigNumber(-625)));
    };
    static fm32(p0, p1, globalState) {
      return _module.D7.create_DC17(!(!(true) || (true)), (function () {
  let _coll5 = new _dafny.Set();
  for (const _compr_5 of (_dafny.Seq.of(new BigNumber(917), new BigNumber(878), new BigNumber(489), new BigNumber(542))).Elements) {
    let _41_v0 = _compr_5;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(917), new BigNumber(878), new BigNumber(489), new BigNumber(542)), _41_v0)) {
      _coll5.add(_module.__default.safeModuloInt(_41_v0, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length)));
    }
  }
  return _coll5;
}()).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(839))).cardinality()))).length))), true, true);
    };
    static fm33(p0, p1, p2, globalState) {
      return ((_module.D31.create_DC74(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))),new _dafny.CodePoint('g'.codePointAt(0))))).dtor_cf126).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0))),new _dafny.CodePoint('h'.codePointAt(0))));
    };
    static fm34(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nkhgoyrc")).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(669))).length),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(102),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(725),true)));
    };
    static fm35(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,true),new BigNumber(18))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),new BigNumber((_dafny.Seq.of(true, false)).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),new BigNumber(41))).Merge(function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),false), _dafny.Map.Empty.slice().updateUnsafe(false,true))).Elements) {
          let _42_v0 = _compr_6;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),false), _dafny.Map.Empty.slice().updateUnsafe(false,true))).contains(_42_v0)) {
            _coll6.push([_42_v0,new BigNumber(-213)]);
          }
        }
        return _coll6;
      }()));
    };
    static fm36(p0, p1, globalState) {
      return _module.D10.create_DC21((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(!(false))));
    };
    static fm37(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true, false, false, false, !(false)), _dafny.Set.fromElements(false, true), _dafny.Set.fromElements(!(!(true)))), _dafny.Seq.of(_dafny.Set.fromElements(!(true), false, false), _dafny.Set.fromElements(true))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-284)), function (_43_i0) {
        return _dafny.Set.fromElements(false, !(false), true);
      }));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      if (true) {
        return _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('c'.codePointAt(0))));
      } else {
        return _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('v'.codePointAt(0))));
      }
    };
    static fm39(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Set.fromElements(false));
    };
    static fm40(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),new BigNumber(-975))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_module.D20.create_DC48(false, _dafny.Seq.of(true), _dafny.Seq.of(!(true), false), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length))).dtor_cf75,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), function (_44_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(302)), function (_45_i1) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).length));
      })).length),new _dafny.CodePoint('v'.codePointAt(0)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),new BigNumber((_dafny.Seq.UnicodeFromString("dfpqhopax")).length))));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return _module.D16.create_DC38(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sbmyxmjv"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(853)), function (_46_i0) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})), true, ((!(false)) ? (new BigNumber((_dafny.Set.fromElements(new BigNumber(438), new BigNumber((_dafny.Seq.UnicodeFromString("qra")).length))).length)) : (new BigNumber((_dafny.Seq.of(false)).length))), true);
    };
    static fm42(p0, p1, p2, p3, globalState) {
      let _source3 = _module.D26.create_DC63(false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-795)), function (_47_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(863))).length), !(false));
      if (_source3.is_DC62) {
        let _48___mcc_h0 = (_source3).cf100;
        let _49___mcc_h1 = (_source3).cf101;
        let _50_cf101 = _49___mcc_h1;
        let _51_cf100 = _48___mcc_h0;
        return true;
      } else if (_source3.is_DC63) {
        let _52___mcc_h2 = (_source3).cf102;
        let _53___mcc_h3 = (_source3).cf103;
        let _54___mcc_h4 = (_source3).cf104;
        let _55___mcc_h5 = (_source3).cf105;
        let _56_cf105 = _55___mcc_h5;
        let _57_cf104 = _54___mcc_h4;
        let _58_cf103 = _53___mcc_h3;
        let _59_cf102 = _52___mcc_h2;
        return false;
      } else {
        let _60___mcc_h6 = (_source3).cf99;
        let _61_cf99 = _60___mcc_h6;
        return true;
      }
    };
    static fm43(p0, p1, globalState) {
      return _dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true, true));
    };
    static fm44(globalState) {
      return _dafny.MultiSet.fromElements(_module.D16.create_DC39(_module.D16.create_DC37(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length)))));
    };
    static fm45(p0, globalState) {
      return _module.D25.create_DC60(true, !(false), _dafny.Seq.UnicodeFromString("rdydk"));
    };
    static fm46(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("mhr"))).Intersect(((_module.D32.create_DC77(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fk"), _dafny.Seq.UnicodeFromString("oiba"), _dafny.Seq.UnicodeFromString("orlal")))).dtor_cf129).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("cgoxp"))));
    };
    static fm47(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(false),false),(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(88)), function (_62_i0) {
        return new BigNumber(-990);
      })).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(824))).cardinality())));
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('v'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('e'.codePointAt(0))));
    };
    static fm49(globalState) {
      return _module.D20.create_DC48((new BigNumber((_dafny.Seq.UnicodeFromString("djfgxi")).length)).isLessThan(new BigNumber(4)), _dafny.Seq.Concat(_dafny.Seq.of(!(false), !(true)), _dafny.Seq.of(!(false), false)), _dafny.Seq.of(false), new BigNumber((function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(23), new BigNumber(-251))) {
    let _63_v0 = _compr_7;
    if (((new BigNumber(23)).isLessThanOrEqualTo(_63_v0)) && ((_63_v0).isLessThan(new BigNumber(-251)))) {
      _coll7.add(_module.__default.safeDivisionInt(_63_v0, new BigNumber(-694)));
    }
  }
  return _coll7;
}()).length));
    };
    static fm50(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.of(true));
    };
    static fm51(p0, globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Intersect((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,!(!(true))))).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)), _dafny.Map.Empty.slice().updateUnsafe(false,!(false)))));
    };
    static fm52(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC19(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-113), new BigNumber(655))).cardinality()), true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kyprqgfxf")).length),new BigNumber((_dafny.Set.fromElements(true)).length))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC19(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-318)), function (_64_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length), false),new BigNumber(411)))).Merge((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC19(new BigNumber((function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-986), new BigNumber(207))) {
    let _65_v1 = _compr_9;
    if (((new BigNumber(-986)).isLessThanOrEqualTo(_65_v1)) && ((_65_v1).isLessThan(new BigNumber(207)))) {
      _coll9.push([(_65_v1).plus(new BigNumber(183)),_dafny.Seq.UnicodeFromString("qchgrep")]);
    }
  }
  return _coll9;
}()).length), true),(_dafny.ZERO).minus(new BigNumber(-893)))).Keys.Elements) {
          let _66_v0 = _compr_8;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC19(new BigNumber((function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(-986), new BigNumber(207))) {
    let _65_v1 = _compr_10;
    if (((new BigNumber(-986)).isLessThanOrEqualTo(_65_v1)) && ((_65_v1).isLessThan(new BigNumber(207)))) {
      _coll10.push([(_65_v1).plus(new BigNumber(183)),_dafny.Seq.UnicodeFromString("qchgrep")]);
    }
  }
  return _coll10;
}()).length), true),(_dafny.ZERO).minus(new BigNumber(-893)))).contains(_66_v0)) {
            _coll8.push([_66_v0,new BigNumber(-258)]);
          }
        }
        return _coll8;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC19(new BigNumber(993), false),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("wtdlx")).length))).length))));
    };
    static fm53(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-166))).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber((_dafny.Seq.of(new BigNumber(733), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(930))).length), new BigNumber((function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Set.fromElements(_dafny.Seq.of(true))).Elements) {
          let _67_v0 = _compr_11;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(true))).contains(_67_v0)) {
            _coll11.add(_67_v0);
          }
        }
        return _coll11;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),false)).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)))).length))).length))).length)));
    };
    static fm54(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_68_i0) {
        return _dafny.Seq.UnicodeFromString("arlqe");
      }), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yi"), _dafny.Seq.UnicodeFromString("pps"), _dafny.Seq.UnicodeFromString("mgyhqbpjr"), _dafny.Seq.UnicodeFromString("iyuepp"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), function (_69_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("fljjpgts")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("vfbd"))));
    };
    static m0(p0, p1, globalState) {
      let _hi0 = (p0).plus(p0);
      for (let _70_i0 = (_dafny.ZERO).minus(p0); _70_i0.isLessThan(_hi0); _70_i0 = _70_i0.plus(_dafny.ONE)) {
        let _71_v0;
        _71_v0 = _dafny.Seq.UnicodeFromString("ngu");
        let _72_v1;
        _72_v1 = _module.D16.create_DC38(_71_v0, !(!(p1)), p0, p1);
        let _73_v2;
        _73_v2 = _dafny.Seq.of((_72_v1).dtor_cf57, _71_v0, _71_v0);
        let _74_v3;
        _74_v3 = _dafny.MultiSet.fromElements(p1);
        let _75_v4;
        _75_v4 = _module.D1.create_DC1(_74_v3);
        _73_v2 = _module.__default.fm54(p0, _75_v4, _module.__default.fm14(globalState), globalState);
        let _76_v5;
        let _init0 = ((_77_p1) => function (_78_i1) {
          return _dafny.Set.fromElements(_77_p1);
        })(p1);
        let _nw0 = Array((new BigNumber(10)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _76_v5 = _nw0;
        let _79_v6;
        _79_v6 = new _dafny.CodePoint('w'.codePointAt(0));
        let _80_v7;
        let _nw1 = new _module.C9();
        _nw1.__ctor(((p1) ? (_76_v5) : (_76_v5)), _79_v6);
        _80_v7 = _nw1;
        let _81_v8;
        _81_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_79_v6);
        let _82_v9;
        _82_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_81_v8).length),p1);
        let _source4 = _module.D14.create_DC32(((_dafny.ZERO).minus(_module.__default.fm14(globalState))).minus(p0), (((_82_v9).contains(p0)) ? ((_82_v9).get(p0)) : (true)), p1);
        if (_source4.is_DC31) {
          let _83___mcc_h0 = (_source4).cf48;
          let _84_cf48 = _83___mcc_h0;
          let _85_v10;
          _85_v10 = true;
          _85_v10 = p1;
          _85_v10 = p1;
          let _86_v11;
          _86_v11 = _dafny.Seq.of(_70_i0, p0);
          let _87_v12;
          _87_v12 = _dafny.Seq.of((_86_v11)[_module.__default.safeIndex(_84_cf48, new BigNumber((_86_v11).length))], _70_i0, _84_cf48);
          let _88_v13;
          _88_v13 = _dafny.Map.Empty.slice().updateUnsafe(_87_v12,p1);
          let _89_v14;
          _89_v14 = _dafny.Seq.of(_85_v10, _85_v10, p1);
          _88_v13 = (_88_v13).update(_87_v12, (_89_v14)[_module.__default.safeIndex(_70_i0, new BigNumber((_89_v14).length))]);
          let _90_v15;
          let _nw2 = new _module.C0();
          _nw2.__ctor();
          _90_v15 = _nw2;
          let _91_v16;
          _91_v16 = _90_v15;
          _91_v16 = _91_v16;
        } else if (_source4.is_DC32) {
          let _92___mcc_h1 = (_source4).cf49;
          let _93___mcc_h2 = (_source4).cf50;
          let _94___mcc_h3 = (_source4).cf51;
          let _95_cf51 = _94___mcc_h3;
          let _96_cf50 = _93___mcc_h2;
          let _97_cf49 = _92___mcc_h1;
          let _98_v17;
          let _init1 = ((_99_cf51) => function (_100_i2) {
            return _99_cf51;
          })(_95_cf51);
          let _nw3 = Array((_dafny.ONE).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
            _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _98_v17 = _nw3;
          let _index0 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length));
          (_98_v17)[_index0] = p1;
          let _101_v18;
          let _nw4 = new _module.C7();
          _nw4.__ctor();
          _101_v18 = _nw4;
          let _index1 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length));
          let _rhs0 = _95_cf51;
          let _rhs1 = _101_v18;
          let _lhs0 = _98_v17;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length));
          _lhs0[_lhs1] = _rhs0;
          _101_v18 = _rhs1;
          let _102_v19;
          _102_v19 = _dafny.Seq.of(!(_96_cf50), true, _95_cf51, p1, (_98_v17)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length))]);
          let _103_v20;
          _103_v20 = _dafny.Seq.of(_97_cf49);
          let _104_v21;
          _104_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_103_v20).length), (_dafny.ZERO).minus(_70_i0), p0),_96_cf50);
          let _105_v22;
          _105_v22 = _dafny.Seq.of((_102_v19)[_module.__default.safeIndex(_70_i0, new BigNumber((_102_v19).length))], _95_cf51, !((((_104_v21).contains(_module.__default.fm31(new BigNumber(-181), new BigNumber((_dafny.Seq.UnicodeFromString("otxi")).length), globalState))) ? ((_104_v21).get(_module.__default.fm31(new BigNumber(-181), new BigNumber((_dafny.Seq.UnicodeFromString("otxi")).length), globalState))) : (_96_cf50))) || ((_98_v17)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length))]));
          _105_v22 = _dafny.Seq.update(_dafny.Seq.Concat(_102_v19, _102_v19), _module.__default.safeIndex(_97_cf49, new BigNumber((_dafny.Seq.Concat(_102_v19, _102_v19)).length)), !(_70_i0).isEqualTo(_70_i0));
          let _106_v23;
          _106_v23 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(840)).isLessThanOrEqualTo(p0),_70_i0);
          _106_v23 = (_106_v23).update(!(((_96_cf50) ? (_96_cf50) : (_95_cf51))), _70_i0);
          let _107_v24;
          _107_v24 = _dafny.Map.Empty.slice().updateUnsafe((_98_v17)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_98_v17).length))],_79_v6);
          let _108_v25;
          _108_v25 = _dafny.Map.Empty.slice().updateUnsafe(_103_v20,(((_105_v22)[_module.__default.safeIndex(_97_cf49, new BigNumber((_105_v22).length))]) ? (_107_v24) : ((_107_v24).update(!(p1), _79_v6))));
          _108_v25 = (_108_v25).Merge(_108_v25);
        } else if (_source4.is_DC30) {
          let _109___mcc_h4 = (_source4).cf47;
          let _110_cf47 = _109___mcc_h4;
          let _111_v26;
          let _nw5 = Array((new BigNumber(4)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = p1;
          _nw5[(_dafny.ONE).toNumber()] = !(p1);
          _nw5[(new BigNumber(2)).toNumber()] = ((p1) ? ((_80_v7).fm3(new BigNumber(-674), _70_i0, globalState)) : (p1));
          _nw5[(new BigNumber(3)).toNumber()] = true;
          _111_v26 = _nw5;
          let _index2 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length));
          (_111_v26)[_index2] = false;
          let _112_v27;
          _112_v27 = new BigNumber(681);
          let _113_v28;
          _113_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_70_i0, _112_v27)).cardinality()),_70_i0);
          let _114_v29;
          let _nw6 = new _module.C6();
          _nw6.__ctor(_113_v28);
          _114_v29 = _nw6;
          let _115_v30;
          _115_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_114_v29);
          let _116_v31;
          let _nw7 = Array((new BigNumber(28)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _114_v29;
          _nw7[(_dafny.ONE).toNumber()] = _114_v29;
          _nw7[(new BigNumber(2)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(3)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(4)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(5)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(6)).toNumber()] = (((_115_v30).contains(p1)) ? ((_115_v30).get(p1)) : (_114_v29));
          _nw7[(new BigNumber(7)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(8)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(9)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(10)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(11)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(12)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(13)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(14)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(15)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(16)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(17)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(18)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(19)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(20)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(21)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(22)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(23)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(24)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(25)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(26)).toNumber()] = _114_v29;
          _nw7[(new BigNumber(27)).toNumber()] = _114_v29;
          _116_v31 = _nw7;
          let _117_v32;
          _117_v32 = _module.D28.create_DC69(p0, _116_v31, _79_v6);
          let _index3 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length));
          let _rhs2 = p1;
          let _rhs3 = (new BigNumber((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(-295), new BigNumber(165))) {
              let _118_v33 = _compr_12;
              if (((new BigNumber(-295)).isLessThanOrEqualTo(_118_v33)) && ((_118_v33).isLessThan(new BigNumber(165)))) {
                _coll12.push([(_118_v33).minus(_112_v27),_dafny.Map.Empty.slice().updateUnsafe(_70_i0,p1)]);
              }
            }
            return _coll12;
          }()).length)).multipliedBy(p0);
          let _rhs4 = new BigNumber(334);
          let _rhs5 = _117_v32;
          let _rhs6 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), ((_119_v6) => function (_120_i3) {
            return _119_v6;
          })(_79_v6));
          let _lhs2 = _111_v26;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length));
          _lhs2[_lhs3] = _rhs2;
          _112_v27 = _rhs3;
          _112_v27 = _rhs4;
          _117_v32 = _rhs5;
          _71_v0 = _rhs6;
          let _121_v34;
          let _nw8 = new _module.C3();
          _nw8.__ctor(_76_v5, _module.__default.fm0(p0, (_dafny.ZERO).minus(_112_v27), globalState));
          _121_v34 = _nw8;
          let _122_v35;
          _122_v35 = _module.D20.create_DC47(_76_v5);
          let _123_v36;
          let _nw9 = new _module.C10();
          _nw9.__ctor(p0, p1, (_122_v35).dtor_cf73, _79_v6);
          _123_v36 = _nw9;
          let _124_v37;
          _124_v37 = (_123_v36).f9;
          let _125_v38;
          let _nw10 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _125_v38 = _nw10;
          let _126_v39;
          _126_v39 = _dafny.Map.Empty.slice().updateUnsafe((_123_v36).f9,_125_v38);
          let _127_v40;
          let _nw11 = Array((new BigNumber(21)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _125_v38;
          _nw11[(_dafny.ONE).toNumber()] = ((false) ? (_125_v38) : ((((_126_v39).contains((_123_v36).f9)) ? ((_126_v39).get((_123_v36).f9)) : (_125_v38))));
          _nw11[(new BigNumber(2)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(3)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(4)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(5)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(6)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(7)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(8)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(9)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(10)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(11)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(12)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(13)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(14)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(15)).toNumber()] = (((_126_v39).contains(p1)) ? ((_126_v39).get(p1)) : (_125_v38));
          _nw11[(new BigNumber(16)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(17)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(18)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(19)).toNumber()] = _125_v38;
          _nw11[(new BigNumber(20)).toNumber()] = _125_v38;
          _127_v40 = _nw11;
          let _index4 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_127_v40).length));
          (_127_v40)[_index4] = _125_v38;
          let _index5 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length));
          let _index6 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_127_v40).length));
          let _rhs7 = _123_v36;
          let _rhs8 = _124_v37;
          let _rhs9 = ((_dafny.ZERO).minus((_dafny.ZERO).minus((_123_v36).f8))).plus(_112_v27);
          let _rhs10 = !(!((_111_v26)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length))]) || (false)) || ((new BigNumber(-330)).isLessThanOrEqualTo(new BigNumber((_71_v0).length)));
          let _rhs11 = _125_v38;
          let _lhs4 = _111_v26;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_111_v26).length));
          let _lhs6 = _127_v40;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_127_v40).length));
          _123_v36 = _rhs7;
          _124_v37 = _rhs8;
          _112_v27 = _rhs9;
          _lhs4[_lhs5] = _rhs10;
          _lhs6[_lhs7] = _rhs11;
          _71_v0 = _71_v0;
        } else {
          let _128___mcc_h5 = (_source4).cf52;
          let _129_cf52 = _128___mcc_h5;
          let _130_v41;
          _130_v41 = true;
          _130_v41 = _130_v41;
          let _131_v42;
          let _132_v43;
          let _out0;
          let _out1;
          let _outcollector0 = (_80_v7).m1(globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _131_v42 = _out0;
          _132_v43 = _out1;
          let _133_v44;
          let _nw12 = new _module.C0();
          _nw12.__ctor();
          _133_v44 = _nw12;
          let _134_v45;
          _134_v45 = _dafny.Map.Empty.slice().updateUnsafe(_133_v44,_71_v0);
          _134_v45 = (_134_v45).update(_133_v44, _module.__default.fm12(globalState));
          let _135_v46;
          let _nw13 = Array((new BigNumber(7)).toNumber());
          _135_v46 = _nw13;
          let _index7 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_135_v46).length));
          (_135_v46)[_index7] = _80_v7;
          let _index8 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_135_v46).length));
          (_135_v46)[_index8] = _80_v7;
        }
        let _136_v47;
        _136_v47 = false;
        _136_v47 = p1;
      }
      let _137_i4;
      _137_i4 = _dafny.ZERO;
      L0: {
        while ((((p1) || (p1)) ? (!(p1)) : (p1))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_137_i4)) {
              break L0;
            }
            _137_i4 = (_137_i4).plus(_dafny.ONE);
            let _138_v48;
            let _nw14 = Array((new BigNumber(16)).toNumber());
            _138_v48 = _nw14;
            let _139_v49;
            let _nw15 = new _module.C5();
            _nw15.__ctor();
            _139_v49 = _nw15;
            let _index9 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_138_v48).length));
            (_138_v48)[_index9] = _139_v49;
            let _index10 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_138_v48).length));
            (_138_v48)[_index10] = _139_v49;
            let _140_v50;
            _140_v50 = true;
            _140_v50 = p1;
            let _141_v51;
            _141_v51 = _dafny.MultiSet.fromElements(new BigNumber(686));
            let _142_v52;
            _142_v52 = _module.D8.create_DC18(_141_v51);
            let _source5 = _142_v52;
            if (_source5.is_DC19) {
              let _143___mcc_h6 = (_source5).cf29;
              let _144___mcc_h7 = (_source5).cf30;
              let _145_cf30 = _144___mcc_h7;
              let _146_cf29 = _143___mcc_h6;
              _146_cf29 = _module.__default.fm13(p1, (_145_cf30) === (_140_v50), ((p1) ? (p1) : (p1)), globalState);
              let _147_v53;
              let _nw16 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Set.Empty);
              _147_v53 = _nw16;
              let _148_v54;
              _148_v54 = new _dafny.CodePoint('x'.codePointAt(0));
              let _149_v55;
              let _nw17 = new _module.C9();
              _nw17.__ctor(_147_v53, _148_v54);
              _149_v55 = _nw17;
              let _150_v56;
              _150_v56 = _dafny.Map.Empty.slice().updateUnsafe(_146_cf29,p0);
              let _151_v57;
              _151_v57 = _dafny.Map.Empty.slice().updateUnsafe(_149_v55,_150_v56);
              let _152_v59;
              let _nw18 = new _module.C6();
              _nw18.__ctor((((_151_v57).contains(_149_v55)) ? ((_151_v57).get(_149_v55)) : (function () {
                let _coll13 = new _dafny.Map();
                for (const _compr_13 of _dafny.IntegerRange(new BigNumber(146), new BigNumber(160))) {
                  let _153_v58 = _compr_13;
                  if (((new BigNumber(146)).isLessThanOrEqualTo(_153_v58)) && ((_153_v58).isLessThan(new BigNumber(160)))) {
                    _coll13.push([_module.__default.safeDivisionInt(_153_v58, p0),_146_cf29]);
                  }
                }
                return _coll13;
              }())));
              _152_v59 = _nw18;
              let _154_v60;
              let _nw19 = Array((new BigNumber(16)).toNumber()).fill([]);
              _154_v60 = _nw19;
              let _155_v61;
              let _nw20 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
              _155_v61 = _nw20;
              let _index11 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_154_v60).length));
              (_154_v60)[_index11] = _155_v61;
              let _156_v62;
              let _nw21 = new _module.C1();
              _nw21.__ctor(_147_v53, _148_v54);
              _156_v62 = _nw21;
              let _157_v63;
              _157_v63 = _dafny.Seq.of(_156_v62);
              let _158_v64;
              _158_v64 = _dafny.Map.Empty.slice().updateUnsafe(_146_cf29,_157_v63);
              let _index12 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_154_v60).length));
              let _rhs12 = _155_v61;
              let _rhs13 = _module.__default.safeModuloInt(p0, _146_cf29);
              let _rhs14 = new BigNumber((((p1) ? ((_158_v64).update(p0, _157_v63)) : (_158_v64))).length);
              let _lhs8 = _154_v60;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_154_v60).length));
              _lhs8[_lhs9] = _rhs12;
              _146_cf29 = _rhs13;
              _146_cf29 = _rhs14;
              let _159_v65;
              let _nw22 = new _module.C2();
              _nw22.__ctor(_147_v53, _148_v54);
              _159_v65 = _nw22;
              let _nw23 = new _module.C2();
              _nw23.__ctor(_147_v53, _148_v54);
              _159_v65 = _nw23;
            } else {
              let _160___mcc_h8 = (_source5).cf28;
              let _161_cf28 = _160___mcc_h8;
              _140_v50 = false;
              let _162_v66;
              let _init2 = ((_163_v50) => function (_164_i5) {
                return _163_v50;
              })(_140_v50);
              let _nw24 = Array((new BigNumber(24)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
                _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _162_v66 = _nw24;
              let _165_v67;
              _165_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_dafny.ZERO).minus(p0), p0, _140_v50, false, globalState),_162_v66);
              let _166_v68;
              _166_v68 = _dafny.Set.fromElements((_module.D30.create_DC73(p0, _140_v50, _162_v66)).dtor_cf125, _162_v66, _162_v66, (((_165_v67).contains(_140_v50)) ? ((_165_v67).get(_140_v50)) : (_162_v66)));
              _166_v68 = (_166_v68).Union(_166_v68);
              let _167_v69;
              let _nw25 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
              _167_v69 = _nw25;
              let _index13 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_167_v69).length));
              (_167_v69)[_index13] = (p0).multipliedBy(new BigNumber(276));
              let _168_v70;
              _168_v70 = new BigNumber(-223);
              let _169_v71;
              _169_v71 = _module.D1.create_DC1(_dafny.MultiSet.fromElements(_140_v50));
              let _170_v72;
              _170_v72 = _dafny.MultiSet.fromElements(p1);
              let _index14 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_167_v69).length));
              let _rhs15 = (p0).multipliedBy(p0);
              let _rhs16 = _module.__default.safeDivisionInt(_168_v70, p0);
              let _rhs17 = _167_v69;
              let _rhs18 = (_170_v72).IsProperSubsetOf((_dafny.MultiSet.fromElements(p1, _140_v50)).Union((_169_v71).dtor_cf1));
              let _lhs10 = _167_v69;
              let _lhs11 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_167_v69).length));
              _lhs10[_lhs11] = _rhs15;
              _168_v70 = _rhs16;
              _167_v69 = _rhs17;
              _140_v50 = _rhs18;
              let _171_v73;
              _171_v73 = _dafny.Seq.of(_140_v50, _140_v50, p1, p1, p1);
              let _172_v74;
              _172_v74 = _dafny.Set.fromElements(_140_v50);
              let _173_v75;
              _173_v75 = _module.D20.create_DC48(true, _dafny.Seq.Concat(_171_v73, _module.__default.fm21(new BigNumber((_170_v72).cardinality()), _168_v70, _172_v74, globalState)), _171_v73, (_167_v69)[_module.__default.safeIndex(new BigNumber(654), new BigNumber((_167_v69).length))]);
              let _pat_let_tv0 = _167_v69;
              let _pat_let_tv1 = _167_v69;
              let _pat_let_tv2 = p1;
              _173_v75 = function (_pat_let0_0) {
                return function (_174_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_175_dt__update_hcf77_h0) {
                      return function (_pat_let2_0) {
                        return function (_176_dt__update_hcf74_h0) {
                          return _module.D20.create_DC48(_176_dt__update_hcf74_h0, (_174_dt__update__tmp_h0).dtor_cf75, (_174_dt__update__tmp_h0).dtor_cf76, _175_dt__update_hcf77_h0);
                        }(_pat_let2_0);
                      }(_pat_let_tv2);
                    }(_pat_let1_0);
                  }((_pat_let_tv1)[_module.__default.safeIndex(new BigNumber(654), new BigNumber((_pat_let_tv0).length))]);
                }(_pat_let0_0);
              }(_module.__default.fm49(globalState));
            }
            let _177_v76;
            _177_v76 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, _140_v50));
            let _178_v77;
            _178_v77 = _dafny.MultiSet.fromElements(_140_v50, p1);
            let _179_v78;
            _179_v78 = _dafny.Seq.UnicodeFromString("dasn");
            let _180_v79;
            _180_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_179_v78).length));
            let _181_v80;
            _181_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_180_v79).length),p1);
            let _182_v81;
            _182_v81 = _dafny.Seq.of(_181_v80);
            let _183_v82;
            _183_v82 = _dafny.Seq.of(_180_v79);
            let _184_v83;
            _184_v83 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
            let _185_v84;
            let _nw26 = Array((new BigNumber(20)).toNumber());
            _nw26[(_dafny.ZERO).toNumber()] = new BigNumber(248);
            _nw26[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(p0, (((_177_v76).contains(_178_v77)) ? ((_177_v76).get(_178_v77)) : (p0)));
            _nw26[(new BigNumber(2)).toNumber()] = p0;
            _nw26[(new BigNumber(3)).toNumber()] = p0;
            _nw26[(new BigNumber(4)).toNumber()] = p0;
            _nw26[(new BigNumber(5)).toNumber()] = new BigNumber((_179_v78).length);
            _nw26[(new BigNumber(6)).toNumber()] = new BigNumber(238);
            _nw26[(new BigNumber(7)).toNumber()] = p0;
            _nw26[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm14(globalState));
            _nw26[(new BigNumber(9)).toNumber()] = p0;
            _nw26[(new BigNumber(10)).toNumber()] = (new BigNumber(641)).multipliedBy(new BigNumber(665));
            _nw26[(new BigNumber(11)).toNumber()] = _module.__default.fm13(false, _140_v50, p1, globalState);
            _nw26[(new BigNumber(12)).toNumber()] = p0;
            _nw26[(new BigNumber(13)).toNumber()] = p0;
            _nw26[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_182_v81).length));
            _nw26[(new BigNumber(15)).toNumber()] = p0;
            _nw26[(new BigNumber(16)).toNumber()] = (p0).plus((_dafny.ZERO).minus(p0));
            _nw26[(new BigNumber(17)).toNumber()] = p0;
            _nw26[(new BigNumber(18)).toNumber()] = (_module.__default.fm14(globalState)).plus((_dafny.ZERO).minus(new BigNumber((_183_v82).length)));
            _nw26[(new BigNumber(19)).toNumber()] = new BigNumber((_184_v83).length);
            _185_v84 = _nw26;
            let _186_v85;
            _186_v85 = _dafny.Seq.of(_140_v50, _140_v50, false, _140_v50);
            let _index15 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_185_v84).length));
            (_185_v84)[_index15] = new BigNumber((_module.__default.fm30(_140_v50, _186_v85, globalState)).cardinality());
            let _index16 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_185_v84).length));
            (_185_v84)[_index16] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.fm13(_140_v50, p1, false, globalState)), p0);
          }
        }
      }
      let _187_v86;
      _187_v86 = _dafny.Seq.UnicodeFromString("vnttcvusf");
      let _188_v87;
      _188_v87 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("k"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), function (_189_i7) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _187_v86, _187_v86, _187_v86);
      let _hi1 = new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm1(new BigNumber((_188_v87).length), p0, p1, p1, globalState))).cardinality());
      for (let _190_i6 = p0; _190_i6.isLessThan(_hi1); _190_i6 = _190_i6.plus(_dafny.ONE)) {
        let _191_v88;
        _191_v88 = new BigNumber(204);
        _191_v88 = p0;
        let _192_v89;
        _192_v89 = _dafny.Set.fromElements(false, p1);
        let _193_v90;
        _193_v90 = _dafny.Seq.of(p1);
        let _194_v91;
        let _nw27 = Array((new BigNumber(19)).toNumber());
        _nw27[(_dafny.ZERO).toNumber()] = _192_v89;
        _nw27[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_module.__default.fm1(p0, new BigNumber((_193_v90).length), p1, true, globalState));
        _nw27[(new BigNumber(2)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(3)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(4)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(5)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(6)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(7)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(8)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(9)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(false);
        _nw27[(new BigNumber(11)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(12)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(13)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(p1);
        _nw27[(new BigNumber(15)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(16)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(17)).toNumber()] = _192_v89;
        _nw27[(new BigNumber(18)).toNumber()] = _192_v89;
        _194_v91 = _nw27;
        let _195_v92;
        _195_v92 = _dafny.Seq.of(_190_i6);
        let _196_v93;
        let _nw28 = new _module.C2();
        _nw28.__ctor(_194_v91, _module.__default.fm0(new BigNumber((_dafny.MultiSet.FromArray(_195_v92)).cardinality()), _191_v88, globalState));
        _196_v93 = _nw28;
        let _197_v94;
        _197_v94 = _dafny.Map.Empty.slice().updateUnsafe(p1,_196_v93);
        let _198_v95;
        _198_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_197_v94).length),_191_v88);
        let _199_v96;
        let _nw29 = new _module.C6();
        _nw29.__ctor(_198_v95);
        _199_v96 = _nw29;
        let _200_v97;
        _200_v97 = _dafny.Map.Empty.slice().updateUnsafe(_199_v96,!(p1));
        _200_v97 = (_200_v97).update(_199_v96, p1);
        let _201_v98;
        let _nw30 = Array((new BigNumber(24)).toNumber()).fill(false);
        _201_v98 = _nw30;
        let _index17 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_201_v98).length));
        (_201_v98)[_index17] = p1;
        let _index18 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_201_v98).length));
        (_201_v98)[_index18] = p1;
        let _202_v99;
        let _nw31 = new _module.C2();
        _nw31.__ctor(_194_v91, new _dafny.CodePoint('u'.codePointAt(0)));
        _202_v99 = _nw31;
        let _203_v100;
        _203_v100 = _dafny.Set.fromElements(_202_v99, _202_v99, _202_v99);
        let _204_v101;
        _204_v101 = _dafny.Set.fromElements(_203_v100);
        let _index19 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_201_v98).length));
        let _index20 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_201_v98).length));
        let _rhs19 = ((new BigNumber((_192_v89).length)).plus(_191_v88)).isEqualTo(new BigNumber((_module.__default.fm24(p0, globalState)).length));
        let _rhs20 = _187_v86;
        let _rhs21 = (p1) === (false);
        let _rhs22 = _dafny.Set.fromElements((_203_v100).Intersect(_203_v100));
        let _lhs12 = _201_v98;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_201_v98).length));
        let _lhs14 = _201_v98;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_201_v98).length));
        _lhs12[_lhs13] = _rhs19;
        _187_v86 = _rhs20;
        _lhs14[_lhs15] = _rhs21;
        _204_v101 = _rhs22;
        let _205_v102;
        _205_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_201_v98)[_module.__default.safeIndex(new BigNumber(256), new BigNumber((_201_v98).length))]);
        let _index21 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_201_v98).length));
        (_201_v98)[_index21] = !((p1) && ((_202_v99).fm3(_module.__default.fm14(globalState), new BigNumber((_205_v102).length), globalState))) || (((p1) ? (p1) : (p1)));
      }
      let _206_v103;
      _206_v103 = _dafny.Set.fromElements(p1, p1, _module.__default.fm1(p0, new BigNumber((_187_v86).length), p1, false, globalState), p1, p1);
      let _207_v104;
      let _nw32 = Array((new BigNumber(5)).toNumber());
      _nw32[(_dafny.ZERO).toNumber()] = _206_v103;
      _nw32[(_dafny.ONE).toNumber()] = _206_v103;
      _nw32[(new BigNumber(2)).toNumber()] = _206_v103;
      _nw32[(new BigNumber(3)).toNumber()] = _206_v103;
      _nw32[(new BigNumber(4)).toNumber()] = _206_v103;
      _207_v104 = _nw32;
      let _208_v105;
      let _nw33 = new _module.C2();
      _nw33.__ctor(_207_v104, new _dafny.CodePoint('i'.codePointAt(0)));
      _208_v105 = _nw33;
      let _hi2 = p0;
      for (let _209_i8 = p0; _209_i8.isLessThan(_hi2); _209_i8 = _209_i8.plus(_dafny.ONE)) {
        let _210_v106;
        _210_v106 = false;
        let _211_v107;
        _211_v107 = _dafny.Seq.of(p0);
        _210_v106 = !(_module.__default.fm1(new BigNumber((_211_v107).length), _209_i8, _210_v106, !(p1), globalState));
        if (_210_v106) {
          _210_v106 = p1;
          let _212_v108;
          _212_v108 = new _dafny.CodePoint('a'.codePointAt(0));
          let _213_v109;
          let _nw34 = new _module.C3();
          _nw34.__ctor(_207_v104, _212_v108);
          _213_v109 = _nw34;
          _213_v109 = _213_v109;
          _210_v106 = (_213_v109).fm5(p0, p1, p0, p0, globalState);
          let _214_v110;
          let _nw35 = new _module.C9();
          _nw35.__ctor(((_210_v106) ? (_207_v104) : (_207_v104)), _212_v108);
          _214_v110 = _nw35;
          let _215_v111;
          let _nw36 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _215_v111 = _nw36;
          let _216_v112;
          _216_v112 = _dafny.Map.Empty.slice().updateUnsafe(_212_v108,_212_v108);
          let _index22 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_215_v111).length));
          (_215_v111)[_index22] = (((_216_v112).contains(_212_v108)) ? ((_216_v112).get(_212_v108)) : (_212_v108));
          let _217_v113;
          _217_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(491),p1);
          let _218_v114;
          _218_v114 = _dafny.Seq.of(!((_208_v105).fm5(_209_i8, (((_217_v113).contains(_209_i8)) ? ((_217_v113).get(_209_i8)) : (false)), p0, new BigNumber(840), globalState)), p1);
          let _219_v115;
          _219_v115 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_218_v114).length), new BigNumber((_217_v113).length)),_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("s"), _187_v86));
          let _index23 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_215_v111).length));
          let _rhs23 = _dafny.Seq.Concat(_187_v86, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-524)), ((_220_v108) => function (_221_i9) {
            return _220_v108;
          })(_212_v108)));
          let _rhs24 = _212_v108;
          let _rhs25 = (((_219_v115).contains(new BigNumber((((p1) ? (_187_v86) : (_187_v86))).length))) ? ((_219_v115).get(new BigNumber((((p1) ? (_187_v86) : (_187_v86))).length))) : (p1));
          let _lhs16 = _215_v111;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_215_v111).length));
          _187_v86 = _rhs23;
          _lhs16[_lhs17] = _rhs24;
          _210_v106 = _rhs25;
        } else {
          let _222_v116;
          _222_v116 = new BigNumber(401);
          let _223_v117;
          _223_v117 = _dafny.MultiSet.fromElements(new BigNumber(878));
          let _224_v118;
          _224_v118 = new _dafny.CodePoint('s'.codePointAt(0));
          let _225_v119;
          let _nw37 = Array((new BigNumber(3)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _224_v118;
          _nw37[(_dafny.ONE).toNumber()] = _224_v118;
          _nw37[(new BigNumber(2)).toNumber()] = _224_v118;
          _225_v119 = _nw37;
          let _226_v120;
          _226_v120 = _dafny.Set.fromElements(_225_v119);
          let _227_v121;
          _227_v121 = _dafny.MultiSet.fromElements(_210_v106, _210_v106);
          let _rhs26 = (((_210_v106) ? (new BigNumber(-464)) : (_209_i8))).plus(_module.__default.safeModuloInt(_222_v116, p0));
          let _rhs27 = (_dafny.ZERO).minus(p0);
          let _rhs28 = (_209_i8).isLessThan((((_223_v117).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), function (_229_i10) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length))) ? ((_223_v117).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), function (_228_i10) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length))) : (new BigNumber((_226_v120).length))));
          let _rhs29 = ((false) ? (p1) : (!((_227_v121).IsSubsetOf(_dafny.MultiSet.fromElements(p1)))));
          _222_v116 = _rhs26;
          _222_v116 = _rhs27;
          _210_v106 = _rhs28;
          _210_v106 = _rhs29;
          _222_v116 = new BigNumber((function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(674), new BigNumber(620))) {
              let _230_v122 = _compr_14;
              if (((new BigNumber(674)).isLessThanOrEqualTo(_230_v122)) && ((_230_v122).isLessThan(new BigNumber(620)))) {
                _coll14.push([(_230_v122).minus(p0),p0]);
              }
            }
            return _coll14;
          }()).length);
          let _231_v123;
          let _nw38 = new _module.C9();
          _nw38.__ctor(((_210_v106) ? (_207_v104) : (_207_v104)), new _dafny.CodePoint('u'.codePointAt(0)));
          _231_v123 = _nw38;
          _231_v123 = _231_v123;
          let _232_v124;
          let _nw39 = new _module.C9();
          _nw39.__ctor(_207_v104, ((true) ? (_224_v118) : (_224_v118)));
          _232_v124 = _nw39;
          _222_v116 = _222_v116;
        }
        let _233_v125;
        _233_v125 = new BigNumber(894);
        _233_v125 = (_209_i8).plus(new BigNumber((_187_v86).length));
        _233_v125 = p0;
      }
      let _234_v126;
      _234_v126 = _module.D1.create_DC2(p0, p0);
      _234_v126 = _234_v126;
      return;
    }
    static Main(__noArgsParameter) {
      let _235_v0;
      _235_v0 = true;
      let _236_v1;
      _236_v1 = _dafny.Set.fromElements(_235_v0);
      let _237_v2;
      let _nw40 = Array((new BigNumber(25)).toNumber()).fill(false);
      _237_v2 = _nw40;
      let _238_globalState;
      let _nw41 = new _module.GlobalState();
      _nw41.__ctor((_236_v1).Difference(_236_v1), new BigNumber(214), _dafny.Seq.Create(_module.__default.abs(new BigNumber(358)), ((_239_v0) => function (_240_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(836), new BigNumber(40))).length),_239_v0)).length);
      })(_235_v0)), ((_235_v0) ? (_237_v2) : (_237_v2)));
      _238_globalState = _nw41;
      let _index24 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
      (_237_v2)[_index24] = _235_v0;
      let _241_v3;
      _241_v3 = new BigNumber(263);
      let _index25 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
      (_237_v2)[_index25] = (_241_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(_235_v0, false, _235_v0)).length));
      let _242_v4;
      _242_v4 = new _dafny.CodePoint('r'.codePointAt(0));
      let _243_v5;
      let _nw42 = Array((new BigNumber(12)).toNumber());
      _nw42[(_dafny.ZERO).toNumber()] = _242_v4;
      _nw42[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
      _nw42[(new BigNumber(2)).toNumber()] = _242_v4;
      _nw42[(new BigNumber(3)).toNumber()] = _242_v4;
      _nw42[(new BigNumber(4)).toNumber()] = _242_v4;
      _nw42[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
      _nw42[(new BigNumber(6)).toNumber()] = _242_v4;
      _nw42[(new BigNumber(7)).toNumber()] = ((_235_v0) ? (_242_v4) : (_242_v4));
      _nw42[(new BigNumber(8)).toNumber()] = (((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) ? (_module.__default.fm0(_241_v3, _241_v3, _238_globalState)) : (_module.__default.fm0(_241_v3, _241_v3, _238_globalState)));
      _nw42[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
      _nw42[(new BigNumber(10)).toNumber()] = _242_v4;
      _nw42[(new BigNumber(11)).toNumber()] = _242_v4;
      _243_v5 = _nw42;
      let _index26 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
      (_243_v5)[_index26] = _242_v4;
      let _244_v6;
      _244_v6 = _dafny.Seq.of((_dafny.ZERO).minus(_241_v3));
      let _index27 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
      let _rhs30 = new _dafny.CodePoint('m'.codePointAt(0));
      let _rhs31 = (_244_v6)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_244_v6).length))];
      let _rhs32 = (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))];
      let _lhs18 = _243_v5;
      let _lhs19 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
      _lhs18[_lhs19] = _rhs30;
      _241_v3 = _rhs31;
      _235_v0 = _rhs32;
      if ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) {
        let _245_v7;
        _245_v7 = _dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]));
        let _246_v8;
        _246_v8 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,_245_v7);
        let _247_v9;
        _247_v9 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,(_245_v7).IsSubsetOf((((_246_v8).contains(new BigNumber(389))) ? ((_246_v8).get(new BigNumber(389))) : (_dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))])))));
        let _248_v10;
        _248_v10 = _module.D1.create_DC1(_245_v7);
        let _249_v11;
        _249_v11 = _module.D2.create_DC3(_237_v2);
        let _index28 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        let _rhs33 = (((_247_v9).contains(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(((_248_v10).dtor_cf1).cardinality()))).length)), _241_v3))) ? ((_247_v9).get(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(((_248_v10).dtor_cf1).cardinality()))).length)), _241_v3))) : (_235_v0));
        let _rhs34 = false;
        let _rhs35 = (_249_v11).dtor_cf4;
        let _lhs20 = _237_v2;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        _235_v0 = _rhs33;
        _lhs20[_lhs21] = _rhs34;
        _237_v2 = _rhs35;
        let _250_v12;
        _250_v12 = _dafny.Set.fromElements(_237_v2);
        let _251_v13;
        _251_v13 = _250_v12;
        let _index29 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index29] = ((_251_v13)).IsSubsetOf((_250_v12).Difference(_250_v12));
        let _252_v14;
        _252_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-845),_237_v2);
        let _253_v15;
        _253_v15 = _dafny.Seq.UnicodeFromString("q");
        let _254_v16;
        _254_v16 = _dafny.Set.fromElements(_241_v3, _241_v3);
        let _255_v17;
        let _nw43 = Array((new BigNumber(13)).toNumber());
        _nw43[(_dafny.ZERO).toNumber()] = _237_v2;
        _nw43[(_dafny.ONE).toNumber()] = (((_252_v14).contains(new BigNumber((_253_v15).length))) ? ((_252_v14).get(new BigNumber((_253_v15).length))) : (_237_v2));
        _nw43[(new BigNumber(2)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(3)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(4)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(5)).toNumber()] = ((_module.__default.fm1(_241_v3, new BigNumber((_254_v16).length), true, _235_v0, _238_globalState)) ? (_237_v2) : (_237_v2));
        _nw43[(new BigNumber(6)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(7)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(8)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(9)).toNumber()] = (_249_v11).dtor_cf4;
        _nw43[(new BigNumber(10)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(11)).toNumber()] = _237_v2;
        _nw43[(new BigNumber(12)).toNumber()] = _237_v2;
        _255_v17 = _nw43;
        let _index30 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_255_v17).length));
        (_255_v17)[_index30] = _237_v2;
        let _index31 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_255_v17).length));
        (_255_v17)[_index31] = (_249_v11).dtor_cf4;
        let _256_v18;
        let _init3 = ((_257_v16) => function (_258_i1) {
          return _257_v16;
        })(_254_v16);
        let _nw44 = Array((new BigNumber(18)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw44.length); _i0_3++) {
          _nw44[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _256_v18 = _nw44;
        let _259_v19;
        _259_v19 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,_256_v18);
        _259_v19 = (_259_v19).update(_241_v3, _256_v18);
        let _260_v20;
        _260_v20 = _dafny.Set.fromElements(_module.__default.fm2(_241_v3, _235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], new BigNumber((_253_v15).length), _238_globalState));
        _254_v16 = _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber((_260_v20).length), _241_v3));
      } else {
        let _261_v21;
        _261_v21 = _dafny.Map.Empty.slice().updateUnsafe(_237_v2,((_dafny.ZERO).minus(new BigNumber((_244_v6).length))).isLessThan(new BigNumber(-994)));
        _241_v3 = new BigNumber((_261_v21).length);
        let _262_v22;
        _262_v22 = _dafny.Seq.of(_235_v0);
        if ((new BigNumber(((_dafny.MultiSet.FromArray(_262_v22)).Difference(_dafny.MultiSet.fromElements(true))).cardinality())).isEqualTo(new BigNumber(463))) {
          let _263_v23;
          _263_v23 = _dafny.Map.Empty.slice().updateUnsafe((_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))],!(_235_v0));
          _263_v23 = (_263_v23).update((_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          let _264_v24;
          _264_v24 = _dafny.Seq.UnicodeFromString("sp");
          let _265_v25;
          _265_v25 = _dafny.MultiSet.fromElements(_241_v3, _241_v3);
          let _266_v26;
          _266_v26 = _dafny.Seq.of(_265_v25);
          let _rhs36 = new BigNumber((_dafny.Seq.update(_264_v24, _module.__default.safeIndex((new BigNumber((_264_v24).length)).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("tslisom"), _module.__default.safeIndex(new BigNumber((_264_v24).length), new BigNumber((_dafny.Seq.UnicodeFromString("tslisom")).length)), (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))])).length)), new BigNumber((_264_v24).length)), (_264_v24)[_module.__default.safeIndex(_241_v3, new BigNumber((_264_v24).length))])).length);
          let _rhs37 = (new BigNumber(-232)).multipliedBy(_241_v3);
          let _rhs38 = (((_262_v22)[_module.__default.safeIndex(new BigNumber(((_266_v26)[_module.__default.safeIndex(_241_v3, new BigNumber((_266_v26).length))]).cardinality()), new BigNumber((_262_v22).length))]) ? ((new BigNumber((_dafny.MultiSet.fromElements(_235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0)).cardinality())).plus(_241_v3)) : ((_dafny.ZERO).minus(_241_v3)));
          let _rhs39 = _241_v3;
          let _rhs40 = _235_v0;
          _241_v3 = _rhs36;
          _241_v3 = _rhs37;
          _241_v3 = _rhs38;
          _241_v3 = _rhs39;
          _235_v0 = _rhs40;
          let _267_v27;
          _267_v27 = _dafny.Seq.of(_236_v1);
          _module.__default.m0(new BigNumber(((_236_v1).Intersect((_267_v27)[_module.__default.safeIndex((_dafny.ZERO).minus(_241_v3), new BigNumber((_267_v27).length))])).length), _235_v0, _238_globalState);
          let _268_v28;
          _268_v28 = _module.D2.create_DC5(_264_v24);
          _268_v28 = _268_v28;
          let _269_v29;
          let _nw45 = Array((new BigNumber(25)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _236_v1;
          _nw45[(_dafny.ONE).toNumber()] = _module.__default.fm26(_241_v3, _241_v3, _238_globalState);
          _nw45[(new BigNumber(2)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(3)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(4)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(5)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(6)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(7)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(8)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(9)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_235_v0);
          _nw45[(new BigNumber(11)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(12)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(13)).toNumber()] = _module.__default.fm26(_241_v3, new BigNumber((_262_v22).length), _238_globalState);
          _nw45[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_235_v0);
          _nw45[(new BigNumber(15)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(16)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(17)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(18)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(19)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(20)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], false);
          _nw45[(new BigNumber(22)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(23)).toNumber()] = _236_v1;
          _nw45[(new BigNumber(24)).toNumber()] = _module.__default.fm26(new BigNumber((_264_v24).length), new BigNumber((_236_v1).length), _238_globalState);
          _269_v29 = _nw45;
          let _270_v30;
          let _nw46 = new _module.C4();
          _nw46.__ctor(_241_v3, new BigNumber(116), _269_v29, _242_v4);
          _270_v30 = _nw46;
        } else {
          let _271_v31;
          _271_v31 = _dafny.Map.Empty.slice().updateUnsafe(_236_v1,(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          let _272_v32;
          _272_v32 = _dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_271_v31).equals(_271_v31));
          _241_v3 = (_dafny.ZERO).minus(new BigNumber((_272_v32).cardinality()));
          _module.__default.m0(_241_v3, (_241_v3).isLessThanOrEqualTo((_244_v6)[_module.__default.safeIndex(new BigNumber(681), new BigNumber((_244_v6).length))]), _238_globalState);
          let _273_v33;
          _273_v33 = _dafny.Seq.UnicodeFromString("yfq");
          _273_v33 = _273_v33;
          let _274_v35;
          _274_v35 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_235_v0, _235_v0));
          let _275_v36;
          _275_v36 = _module.D28.create_DC68(function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of (_274_v35).Elements) {
    let _276_v34 = _compr_15;
    if ((_274_v35).contains(_276_v34)) {
      _coll15.push([_276_v34,_241_v3]);
    }
  }
  return _coll15;
}());
          _module.__default.m0(_241_v3, (_262_v22)[_module.__default.safeIndex(new BigNumber(((_275_v36).dtor_cf116).length), new BigNumber((_262_v22).length))], _238_globalState);
          let _277_v37;
          let _init4 = ((_278_v1) => function (_279_i2) {
            return _278_v1;
          })(_236_v1);
          let _nw47 = Array((new BigNumber(9)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw47.length); _i0_4++) {
            _nw47[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _277_v37 = _nw47;
          let _280_v38;
          let _nw48 = new _module.C3();
          _nw48.__ctor(_277_v37, _242_v4);
          _280_v38 = _nw48;
        }
        _241_v3 = (_241_v3).plus(_241_v3);
        if (_235_v0) {
          let _281_v39;
          _281_v39 = _dafny.Seq.UnicodeFromString("fvwankh");
          _281_v39 = _dafny.Seq.UnicodeFromString("qf");
          let _282_v40;
          let _nw49 = Array((new BigNumber(14)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _236_v1;
          _nw49[(_dafny.ONE).toNumber()] = _236_v1;
          _nw49[(new BigNumber(2)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(3)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(4)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(5)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          _nw49[(new BigNumber(7)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(8)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(9)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(10)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(11)).toNumber()] = _236_v1;
          _nw49[(new BigNumber(12)).toNumber()] = _module.__default.fm26(_241_v3, new BigNumber(944), _238_globalState);
          _nw49[(new BigNumber(13)).toNumber()] = _236_v1;
          _282_v40 = _nw49;
          let _283_v41;
          let _nw50 = new _module.C2();
          _nw50.__ctor(_282_v40, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
          _283_v41 = _nw50;
          _281_v39 = _281_v39;
          let _index32 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          (_237_v2)[_index32] = _235_v0;
          let _284_v42;
          let _nw51 = new _module.C9();
          _nw51.__ctor(_282_v40, _242_v4);
          _284_v42 = _nw51;
          let _nw52 = new _module.C7();
          _nw52.__ctor();
          _284_v42 = _nw52;
        } else {
          let _285_v43;
          _285_v43 = _dafny.Map.Empty.slice().updateUnsafe(_235_v0,true);
          let _286_v44;
          _286_v44 = _dafny.Set.fromElements(_285_v43);
          let _287_v45;
          _287_v45 = _dafny.Seq.UnicodeFromString("t");
          let _288_v46;
          _288_v46 = _dafny.Map.Empty.slice().updateUnsafe(((!(_235_v0)) ? (_286_v44) : (_dafny.Set.fromElements(_285_v43))),_287_v45);
          _288_v46 = (_288_v46).update(_module.__default.fm51(_235_v0, _238_globalState), _dafny.Seq.Concat(_287_v45, _287_v45));
          _module.__default.m0(_module.__default.safeModuloInt(_241_v3, new BigNumber(832)), _235_v0, _238_globalState);
          let _289_v47;
          _289_v47 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,new BigNumber(912));
          _289_v47 = _289_v47;
          let _290_v48;
          _290_v48 = _module.D5.create_DC13(_287_v45, _241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          let _index33 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          let _rhs41 = _module.__default.fm14(_238_globalState);
          let _rhs42 = (_290_v48).dtor_cf18;
          let _rhs43 = (((_285_v43).contains(_235_v0)) ? ((_285_v43).get(_235_v0)) : (!((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))])));
          let _rhs44 = _241_v3;
          let _rhs45 = !((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          let _lhs22 = _237_v2;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          _241_v3 = _rhs41;
          _235_v0 = _rhs42;
          _235_v0 = _rhs43;
          _241_v3 = _rhs44;
          _lhs22[_lhs23] = _rhs45;
          let _291_v49;
          _291_v49 = _module.D27.create_DC65(_module.__default.fm1(_241_v3, _241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, _238_globalState), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), ((_292_v5) => function (_293_i3) {
  return (_292_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_292_v5).length))];
})(_243_v5)), _287_v45), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          _291_v49 = _module.D27.create_DC65(!_dafny.Seq.contains(_287_v45, _module.__default.fm0(_241_v3, new BigNumber(890), _238_globalState)), _dafny.Seq.Concat(_287_v45, _dafny.Seq.UnicodeFromString("uia")), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], !(true));
        }
        let _294_v50;
        _294_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm34((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState)).length),_241_v3);
        let _295_v51;
        _295_v51 = _dafny.Map.Empty.slice().updateUnsafe((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))],_294_v50);
        let _296_v52;
        let _nw53 = new _module.C5();
        _nw53.__ctor();
        _296_v52 = _nw53;
        let _297_v53;
        _297_v53 = _dafny.Seq.of(_296_v52);
        let _298_v54;
        _298_v54 = _module.D27.create_DC64(_297_v53);
        let _rhs46 = _295_v51;
        let _rhs47 = _241_v3;
        let _rhs48 = (new BigNumber((_dafny.Seq.UnicodeFromString("nacylgo")).length)).isLessThanOrEqualTo(_241_v3);
        let _rhs49 = _298_v54;
        _295_v51 = _rhs46;
        _241_v3 = _rhs47;
        _235_v0 = _rhs48;
        _298_v54 = _rhs49;
      }
      _module.__default.m0(_241_v3, _235_v0, _238_globalState);
      let _299_v55;
      _299_v55 = _dafny.Seq.of(_235_v0);
      let _hi3 = new BigNumber((_dafny.Seq.Concat(_299_v55, _299_v55)).length);
      for (let _300_i4 = _241_v3; _300_i4.isLessThan(_hi3); _300_i4 = _300_i4.plus(_dafny.ONE)) {
        let _index34 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index34] = ((_235_v0) ? ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) : (_235_v0));
        let _301_v56;
        _301_v56 = _dafny.Map.Empty.slice().updateUnsafe((((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) ? (_237_v2) : (_237_v2)),(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _301_v56 = (_301_v56).update(_237_v2, true);
        let _index35 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
        (_243_v5)[_index35] = _242_v4;
        let _302_v57;
        _302_v57 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        let _303_v58;
        let _nw54 = Array((new BigNumber(9)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _236_v1;
        _nw54[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements((_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(582)), ((_304_v4) => function (_305_i5) {
  return _304_v4;
})(_242_v4)), new BigNumber((_302_v57).length), _235_v0)).dtor_cf18, _235_v0);
        _nw54[(new BigNumber(2)).toNumber()] = _236_v1;
        _nw54[(new BigNumber(3)).toNumber()] = _236_v1;
        _nw54[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, _235_v0);
        _nw54[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(!(_235_v0));
        _nw54[(new BigNumber(6)).toNumber()] = _236_v1;
        _nw54[(new BigNumber(7)).toNumber()] = _236_v1;
        _nw54[(new BigNumber(8)).toNumber()] = _236_v1;
        _303_v58 = _nw54;
        let _306_v59;
        let _nw55 = new _module.C3();
        _nw55.__ctor(_303_v58, _242_v4);
        _306_v59 = _nw55;
      }
      let _307_v60;
      _307_v60 = _dafny.Map.Empty.slice().updateUnsafe((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))],_244_v6);
      _module.__default.m0(new BigNumber((((_307_v60).update((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), function (_308_i6) {
        return new BigNumber(313);
      }))).Merge((_307_v60).update(_235_v0, _244_v6))).length), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState);
      let _309_v61;
      _309_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-494),_241_v3);
      let _310_v62;
      let _nw56 = new _module.C6();
      _nw56.__ctor(_309_v61);
      _310_v62 = _nw56;
      let _311_v63;
      _311_v63 = _dafny.Seq.UnicodeFromString("waupnh");
      let _312_v64;
      _312_v64 = _dafny.Map.Empty.slice().updateUnsafe(!((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]),_235_v0);
      let _hi4 = _module.__default.safeModuloInt(new BigNumber((_311_v63).length), new BigNumber((_312_v64).length));
      for (let _313_i7 = (new BigNumber(148)).plus(_241_v3); _313_i7.isLessThan(_hi4); _313_i7 = _313_i7.plus(_dafny.ONE)) {
        let _314_v65;
        let _nw57 = new _module.C5();
        _nw57.__ctor();
        _314_v65 = _nw57;
        let _315_v66;
        _315_v66 = _dafny.Seq.of(_314_v65);
        let _316_v67;
        _316_v67 = _module.D27.create_DC64(_315_v66);
        _316_v67 = _316_v67;
        if (_235_v0) {
          let _index36 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          (_237_v2)[_index36] = !((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) || (_235_v0);
          let _317_v68;
          _317_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_241_v3, _241_v3),false);
          _317_v68 = ((_317_v68).update(_313_i7, _235_v0)).Merge(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(424), new BigNumber(861))) {
              let _318_v69 = _compr_16;
              if (((new BigNumber(424)).isLessThanOrEqualTo(_318_v69)) && ((_318_v69).isLessThan(new BigNumber(861)))) {
                _coll16.push([(_318_v69).plus(_241_v3),(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]]);
              }
            }
            return _coll16;
          }());
          let _319_v70;
          _319_v70 = _dafny.Seq.of(_312_v64, _312_v64);
          let _320_v71;
          let _321_v72;
          let _322_v73;
          let _323_v74;
          let _out2;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = (_310_v62).m9(_319_v70, false, _238_globalState);
          _out2 = _outcollector1[0];
          _out3 = _outcollector1[1];
          _out4 = _outcollector1[2];
          _out5 = _outcollector1[3];
          _320_v71 = _out2;
          _321_v72 = _out3;
          _322_v73 = _out4;
          _323_v74 = _out5;
          let _324_v75;
          _324_v75 = _dafny.Map.Empty.slice().updateUnsafe((_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))],_313_i7);
          _324_v75 = (_324_v75).update((_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))], (_dafny.ZERO).minus(_313_i7));
          _311_v63 = _311_v63;
        } else {
          let _325_v76;
          let _init5 = ((_326_v63) => function (_327_i8) {
            return _326_v63;
          })(_311_v63);
          let _nw58 = Array((new BigNumber(25)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw58.length); _i0_5++) {
            _nw58[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _325_v76 = _nw58;
          let _index37 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_325_v76).length));
          (_325_v76)[_index37] = _311_v63;
          let _328_v77;
          let _nw59 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
          _328_v77 = _nw59;
          let _329_v78;
          let _nw60 = new _module.C1();
          _nw60.__ctor((((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) ? (_328_v77) : (_328_v77)), (_311_v63)[_module.__default.safeIndex(_313_i7, new BigNumber((_311_v63).length))]);
          _329_v78 = _nw60;
          let _330_v79;
          let _nw61 = new _module.C0();
          _nw61.__ctor();
          _330_v79 = _nw61;
          let _331_v80;
          _331_v80 = _dafny.MultiSet.fromElements(new BigNumber(136));
          let _332_v81;
          _332_v81 = _330_v79;
          let _index38 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_325_v76).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          let _rhs50 = _dafny.Seq.update(_311_v63, _module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber((_331_v80).cardinality())).plus(_module.__default.fm13((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, _235_v0, _238_globalState))), new BigNumber((_311_v63).length)), (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
          let _rhs51 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("rinkuq"), _311_v63);
          let _rhs52 = _329_v78;
          let _rhs53 = (((_235_v0) ? (_332_v81) : (_332_v81)));
          let _lhs24 = _325_v76;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_325_v76).length));
          let _lhs26 = _237_v2;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          _lhs24[_lhs25] = _rhs50;
          _lhs26[_lhs27] = _rhs51;
          _329_v78 = _rhs52;
          _330_v79 = _rhs53;
          let _333_v82;
          let _nw62 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _333_v82 = _nw62;
          let _334_v83;
          let _nw63 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _334_v83 = _nw63;
          let _index40 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_333_v82).length));
          (_333_v82)[_index40] = (_dafny.Map.Empty.slice().updateUnsafe(_334_v83,new BigNumber(((_325_v76)[_module.__default.safeIndex(new BigNumber(31), new BigNumber((_325_v76).length))]).length))).update(_334_v83, _313_i7);
          let _335_v84;
          let _nw64 = new _module.C7();
          _nw64.__ctor();
          _335_v84 = _nw64;
          let _336_v85;
          _336_v85 = _dafny.Map.Empty.slice().updateUnsafe(_335_v84,_241_v3);
          let _index41 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_333_v82).length));
          (_333_v82)[_index41] = _dafny.Map.Empty.slice().updateUnsafe(_334_v83,_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_336_v85).length)), _313_i7));
          let _337_v86;
          _337_v86 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          _337_v86 = (_337_v86).update(_241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          let _338_v87;
          let _nw65 = new _module.C7();
          _nw65.__ctor();
          _338_v87 = _nw65;
          _235_v0 = (_235_v0) || ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        }
        _241_v3 = new BigNumber(868);
        let _339_v90;
        let _nw66 = Array((new BigNumber(23)).toNumber());
        _nw66[(_dafny.ZERO).toNumber()] = _236_v1;
        _nw66[(_dafny.ONE).toNumber()] = _236_v1;
        _nw66[(new BigNumber(2)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(3)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(4)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(5)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_235_v0);
        _nw66[(new BigNumber(7)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(8)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw66[(new BigNumber(10)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw66[(new BigNumber(12)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(13)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(_235_v0);
        _nw66[(new BigNumber(15)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(16)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(17)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_235_v0);
        _nw66[(new BigNumber(19)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(20)).toNumber()] = _236_v1;
        _nw66[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw66[(new BigNumber(22)).toNumber()] = _236_v1;
        _339_v90 = _nw66;
        let _340_v91;
        let _nw67 = Array((new BigNumber(24)).toNumber());
        _340_v91 = _nw67;
        let _341_v92;
        _341_v92 = _module.D28.create_DC69(_241_v3, _340_v91, _242_v4);
        let _342_v93;
        let _nw68 = new _module.C11();
        _nw68.__ctor(function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of ((_310_v62).f10).Keys.Elements) {
            let _343_v88 = _compr_17;
            if (((_310_v62).f10).contains(_343_v88)) {
              _coll17.push([(_343_v88).plus(_241_v3),function () {
                let _coll18 = new _dafny.Map();
                for (const _compr_18 of (_dafny.Seq.of(_241_v3)).Elements) {
                  let _344_v89 = _compr_18;
                  if (_dafny.Seq.contains(_dafny.Seq.of(_241_v3), _344_v89)) {
                    _coll18.push([_module.__default.safeModuloInt(_344_v89, _241_v3),(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]]);
                  }
                }
                return _coll18;
              }()]);
            }
          }
          return _coll17;
        }(), _313_i7, _339_v90, (_341_v92).dtor_cf119);
        _342_v93 = _nw68;
        let _345_v94;
        _345_v94 = _dafny.Map.Empty.slice().updateUnsafe(_342_v93,(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        let _346_v95;
        let _nw69 = Array((new BigNumber(22)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(false, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw69[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_235_v0);
        _nw69[(new BigNumber(2)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(3)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(4)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(5)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw69[(new BigNumber(7)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(!(_235_v0), false);
        _nw69[(new BigNumber(9)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(10)).toNumber()] = _module.__default.fm26(_241_v3, _313_i7, _238_globalState);
        _nw69[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((((_345_v94).contains(_342_v93)) ? ((_345_v94).get(_342_v93)) : (_235_v0)));
        _nw69[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw69[(new BigNumber(13)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(14)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(15)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(16)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(17)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
        _nw69[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(_235_v0);
        _nw69[(new BigNumber(20)).toNumber()] = _236_v1;
        _nw69[(new BigNumber(21)).toNumber()] = _236_v1;
        _346_v95 = _nw69;
        let _347_v96;
        let _nw70 = new _module.C4();
        _nw70.__ctor(_313_i7, _241_v3, (_342_v93).f4, (_342_v93).f5);
        _347_v96 = _nw70;
        let _348_v97;
        _348_v97 = _dafny.Set.fromElements(_314_v65, _314_v65);
        let _nw71 = new _module.C10();
        _nw71.__ctor((_313_i7).multipliedBy(new BigNumber((_244_v6).length)), !(_dafny.Set.fromElements(_314_v65)).equals(_348_v97), (_342_v93).f4, new _dafny.CodePoint('l'.codePointAt(0)));
        _342_v93 = _nw71;
      }
      let _349_v98;
      _349_v98 = _dafny.Seq.of(_237_v2, _237_v2);
      let _350_v99;
      _350_v99 = _module.D18.create_DC45((_349_v98)[_module.__default.safeIndex(_241_v3, new BigNumber((_349_v98).length))], true, _241_v3, new BigNumber(-410));
      _241_v3 = (_dafny.ZERO).minus((_350_v99).dtor_cf70);
      _241_v3 = _module.__default.safeModuloInt(_241_v3, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_241_v3, new BigNumber((_311_v63).length))));
      let _351_v100;
      _351_v100 = _dafny.Map.Empty.slice().updateUnsafe((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))],new BigNumber(199));
      _module.__default.m0(new BigNumber(((_351_v100).update((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _241_v3)).length), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState);
      if ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) {
        let _352_v101;
        let _nw72 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Set.Empty);
        _352_v101 = _nw72;
        let _353_v102;
        let _nw73 = new _module.C9();
        _nw73.__ctor(_352_v101, _242_v4);
        _353_v102 = _nw73;
        let _354_v103;
        _354_v103 = _dafny.Map.Empty.slice().updateUnsafe(_353_v102,_241_v3);
        let _355_v104;
        let _nw74 = Array((new BigNumber(23)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _236_v1;
        _nw74[(_dafny.ONE).toNumber()] = _236_v1;
        _nw74[(new BigNumber(2)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(3)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(4)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(5)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(6)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(7)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(8)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(9)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(10)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(11)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(12)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(13)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(14)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(15)).toNumber()] = _module.__default.fm26(_241_v3, _241_v3, _238_globalState);
        _nw74[(new BigNumber(16)).toNumber()] = _module.__default.fm26((((_354_v103).contains(_353_v102)) ? ((_354_v103).get(_353_v102)) : (_241_v3)), _241_v3, _238_globalState);
        _nw74[(new BigNumber(17)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(18)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(19)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(20)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(21)).toNumber()] = _236_v1;
        _nw74[(new BigNumber(22)).toNumber()] = _dafny.Set.fromElements(!((_353_v102).fm5(new BigNumber((_311_v63).length), _235_v0, _241_v3, _241_v3, _238_globalState)));
        _355_v104 = _nw74;
        let _356_v105;
        let _nw75 = new _module.C3();
        _nw75.__ctor(_355_v104, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
        _356_v105 = _nw75;
        let _357_v106;
        _357_v106 = _dafny.Map.Empty.slice().updateUnsafe(_242_v4,_356_v105);
        let _358_v107;
        _358_v107 = _dafny.Map.Empty.slice().updateUnsafe(((((_310_v62).f10).contains(_241_v3)) ? (((_310_v62).f10).get(_241_v3)) : (new BigNumber((_357_v106).length))),_237_v2);
        let _359_v108;
        _359_v108 = _dafny.Map.Empty.slice().updateUnsafe(_311_v63,_237_v2);
        let _360_v109;
        _360_v109 = _module.D2.create_DC3(_237_v2);
        let _361_v110;
        _361_v110 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-303)), ((_362_v2, _363_v0) => function (_364_i9) {
          return _module.D7.create_DC17((_362_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_362_v2).length))], _363_v0, false, _363_v0);
        })(_237_v2, _235_v0)));
        let _365_v111;
        let _init6 = ((_366_v3) => function (_367_i10) {
          return (_367_i10).plus(_366_v3);
        })(_241_v3);
        let _nw76 = Array((new BigNumber(26)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw76.length); _i0_6++) {
          _nw76[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _365_v111 = _nw76;
        let _368_v112;
        _368_v112 = _dafny.Map.Empty.slice().updateUnsafe(_365_v111,_241_v3);
        let _369_v113;
        let _nw77 = Array((new BigNumber(22)).toNumber());
        _nw77[(_dafny.ZERO).toNumber()] = _310_v62;
        _nw77[(_dafny.ONE).toNumber()] = _310_v62;
        _nw77[(new BigNumber(2)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(3)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(4)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(5)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(6)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(7)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(8)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(9)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(10)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(11)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(12)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(13)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(14)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(15)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(16)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(17)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(18)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(19)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(20)).toNumber()] = _310_v62;
        _nw77[(new BigNumber(21)).toNumber()] = _310_v62;
        _369_v113 = _nw77;
        let _370_v114;
        let _nw78 = new _module.C5();
        _nw78.__ctor();
        _370_v114 = _nw78;
        let _371_v115;
        _371_v115 = _dafny.Seq.of(_370_v114);
        let _index42 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        let _rhs54 = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,(((_359_v108).contains(_311_v63)) ? ((_359_v108).get(_311_v63)) : ((_360_v109).dtor_cf4)));
        let _rhs55 = (((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) ? (!(((_dafny.ZERO).minus(_241_v3)).isEqualTo(_241_v3))) : (!(_241_v3).isEqualTo(_241_v3)));
        let _rhs56 = ((((_310_v62).f10).contains((new BigNumber((_361_v110).length)).minus((((_368_v112).contains(_365_v111)) ? ((_368_v112).get(_365_v111)) : (_241_v3))))) ? (((_310_v62).f10).get((new BigNumber((_361_v110).length)).minus((((_368_v112).contains(_365_v111)) ? ((_368_v112).get(_365_v111)) : (_241_v3))))) : ((_module.D28.create_DC69(_241_v3, _369_v113, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))])).dtor_cf117));
        let _rhs57 = _module.__default.safeModuloInt(new BigNumber(-421), _module.__default.safeModuloInt((_dafny.ZERO).minus(_241_v3), new BigNumber((_371_v115).length)));
        let _lhs28 = _237_v2;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        _358_v107 = _rhs54;
        _lhs28[_lhs29] = _rhs55;
        _241_v3 = _rhs56;
        _241_v3 = _rhs57;
        let _372_v116;
        let _nw79 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _372_v116 = _nw79;
        let _index43 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_372_v116).length));
        (_372_v116)[_index43] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), function (_373_i11) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _311_v63);
        let _index44 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_372_v116).length));
        (_372_v116)[_index44] = _dafny.Seq.Concat(_311_v63, _311_v63);
        let _374_v117;
        let _nw80 = new _module.C7();
        _nw80.__ctor();
        _374_v117 = _nw80;
        let _index45 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index45] = _235_v0;
        let _375_v118;
        _375_v118 = _dafny.Seq.of(_365_v111);
        _241_v3 = (_module.D11.create_DC25(!(!((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))])), new BigNumber(558), (_375_v118)[_module.__default.safeIndex(new BigNumber((_311_v63).length), new BigNumber((_375_v118).length))], _241_v3)).dtor_cf39;
      } else {
        let _index46 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index46] = (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))];
        let _376_v119;
        let _nw81 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
        _376_v119 = _nw81;
        let _377_v120;
        let _nw82 = new _module.C9();
        _nw82.__ctor(_376_v119, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
        _377_v120 = _nw82;
        let _378_v121;
        _378_v121 = _dafny.Seq.of(_377_v120, _377_v120);
        let _379_v122;
        let _nw83 = Array((new BigNumber(6)).toNumber());
        _nw83[(_dafny.ZERO).toNumber()] = _377_v120;
        _nw83[(_dafny.ONE).toNumber()] = _377_v120;
        _nw83[(new BigNumber(2)).toNumber()] = _377_v120;
        _nw83[(new BigNumber(3)).toNumber()] = (_378_v121)[_module.__default.safeIndex(_241_v3, new BigNumber((_378_v121).length))];
        _nw83[(new BigNumber(4)).toNumber()] = _377_v120;
        _nw83[(new BigNumber(5)).toNumber()] = _377_v120;
        _379_v122 = _nw83;
        let _index47 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_379_v122).length));
        (_379_v122)[_index47] = _377_v120;
        let _index48 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_379_v122).length));
        (_379_v122)[_index48] = _377_v120;
        _241_v3 = new BigNumber(-721);
        _236_v1 = _236_v1;
        let _380_v123;
        _380_v123 = _dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0, _235_v0, _235_v0);
        let _index49 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index49] = (new BigNumber((_380_v123).cardinality())).isLessThanOrEqualTo((_dafny.ZERO).minus(_241_v3));
      }
      let _381_v124;
      _381_v124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_244_v6, _dafny.Seq.of(_241_v3, _241_v3, _241_v3, _241_v3)),_241_v3);
      _381_v124 = (_381_v124).update(_244_v6, _241_v3);
      _351_v100 = _351_v100;
      let _source6 = _module.D1.create_DC2(_241_v3, ((_244_v6)[_module.__default.safeIndex(new BigNumber(-276), new BigNumber((_244_v6).length))]).plus(_241_v3));
      if (_source6.is_DC2) {
        let _382___mcc_h0 = (_source6).cf2;
        let _383___mcc_h1 = (_source6).cf3;
        let _384_cf3 = _383___mcc_h1;
        let _385_cf2 = _382___mcc_h0;
        if ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) {
          let _386_v125;
          _386_v125 = _dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], !(new BigNumber((_244_v6).length)).isEqualTo(new BigNumber(340)), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], !((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]));
          _386_v125 = _386_v125;
          let _index50 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
          (_243_v5)[_index50] = _242_v4;
          _module.__default.m0(new BigNumber(962), _235_v0, _238_globalState);
          let _387_v126;
          let _init7 = ((_388_cf3) => function (_389_i12) {
            return _module.__default.safeModuloInt(_389_i12, _388_cf3);
          })(_384_cf3);
          let _nw84 = Array((_dafny.ONE).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw84.length); _i0_7++) {
            _nw84[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _387_v126 = _nw84;
          let _390_v127;
          _390_v127 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))],new BigNumber((_386_v125).cardinality()))).length));
          let _index51 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_387_v126).length));
          (_387_v126)[_index51] = new BigNumber((_390_v127).length);
          let _391_v128;
          _391_v128 = _dafny.MultiSet.fromElements(_387_v126);
          let _index52 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_387_v126).length));
          let _rhs58 = new BigNumber((((_391_v128).Union(_391_v128)).Union(_391_v128)).cardinality());
          let _rhs59 = (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))];
          let _rhs60 = (_385_cf2).minus((new BigNumber(-812)).plus(_385_cf2));
          let _lhs30 = _387_v126;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_387_v126).length));
          _385_cf2 = _rhs58;
          _235_v0 = _rhs59;
          _lhs30[_lhs31] = _rhs60;
          let _392_v129;
          let _nw85 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _392_v129 = _nw85;
          let _index53 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_392_v129).length));
          (_392_v129)[_index53] = _386_v125;
          let _index54 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_392_v129).length));
          (_392_v129)[_index54] = _386_v125;
        } else {
          let _393_v130;
          let _nw86 = new _module.C0();
          _nw86.__ctor();
          _393_v130 = _nw86;
          let _394_v131;
          _394_v131 = _393_v130;
          _394_v131 = _393_v130;
          let _395_v132;
          let _init8 = ((_396_v1) => function (_397_i13) {
            return _396_v1;
          })(_236_v1);
          let _nw87 = Array((new BigNumber(4)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw87.length); _i0_8++) {
            _nw87[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _395_v132 = _nw87;
          let _398_v133;
          let _nw88 = new _module.C10();
          _nw88.__ctor(_384_cf3, _235_v0, _395_v132, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
          _398_v133 = _nw88;
          let _399_v134;
          _399_v134 = _module.D30.create_DC72(_398_v133);
          let _400_v135;
          let _nw89 = Array((new BigNumber(27)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = (_399_v134).dtor_cf122;
          _nw89[(_dafny.ONE).toNumber()] = _398_v133;
          _nw89[(new BigNumber(2)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(3)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(4)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(5)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(6)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(7)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(8)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(9)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(10)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(11)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(12)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(13)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(14)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(15)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(16)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(17)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(18)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(19)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(20)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(21)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(22)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(23)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(24)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(25)).toNumber()] = _398_v133;
          _nw89[(new BigNumber(26)).toNumber()] = _398_v133;
          _400_v135 = _nw89;
          let _index55 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_400_v135).length));
          (_400_v135)[_index55] = _398_v133;
          let _index56 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_400_v135).length));
          let _rhs61 = _398_v133;
          let _rhs62 = _385_cf2;
          let _rhs63 = _381_v124;
          let _lhs32 = _400_v135;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_400_v135).length));
          _lhs32[_lhs33] = _rhs61;
          _385_cf2 = _rhs62;
          _381_v124 = _rhs63;
          let _401_v136;
          let _nw90 = new _module.C9();
          _nw90.__ctor(_395_v132, _242_v4);
          _401_v136 = _nw90;
          _401_v136 = _401_v136;
          let _402_v137;
          _402_v137 = _module.D8.create_DC19(_384_cf3, _235_v0);
          let _403_v138;
          _403_v138 = _dafny.Map.Empty.slice().updateUnsafe((((_398_v133).f9) ? (_402_v137) : (_402_v137)),(_398_v133).f8);
          let _404_v139;
          _404_v139 = _module.D6.create_DC14(_244_v6);
          let _index57 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          let _index58 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
          let _rhs64 = _dafny.Seq.Concat(_299_v55, _dafny.Seq.update(_299_v55, _module.__default.safeIndex(_384_cf3, new BigNumber((_299_v55).length)), false));
          let _rhs65 = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(716), _385_cf2), (_404_v139).dtor_cf19), (_398_v133).f8);
          let _rhs66 = (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))];
          let _rhs67 = _241_v3;
          let _rhs68 = (_module.__default.fm52((_398_v133).f8, new BigNumber(973), _241_v3, (_dafny.ZERO).minus((_398_v133).f8), _238_globalState)).Merge(_403_v138);
          let _lhs34 = _237_v2;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          let _lhs36 = _243_v5;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length));
          _299_v55 = _rhs64;
          _lhs34[_lhs35] = _rhs65;
          _lhs36[_lhs37] = _rhs66;
          _385_cf2 = _rhs67;
          _403_v138 = _rhs68;
          let _405_v140;
          _405_v140 = _dafny.Set.fromElements(_module.__default.fm26((_dafny.ZERO).minus(new BigNumber(-20)), (_398_v133).f8, _238_globalState), _236_v1, _236_v1, _236_v1);
          let _406_v141;
          _406_v141 = _module.D18.create_DC44(_405_v140);
          _406_v141 = _module.D18.create_DC44((_405_v140).Intersect(_405_v140));
        }
        _311_v63 = _311_v63;
        _385_cf2 = (_dafny.ZERO).minus((new BigNumber(514)).multipliedBy(new BigNumber((_244_v6).length)));
        if (((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) || (_235_v0)) {
          _384_cf3 = _module.__default.fm13(_235_v0, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState);
          _module.__default.m0(_241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState);
          let _407_v142;
          let _nw91 = new _module.C7();
          _nw91.__ctor();
          _407_v142 = _nw91;
          let _index59 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          let _rhs69 = (_dafny.ZERO).minus((_385_cf2).minus(new BigNumber((_299_v55).length)));
          let _rhs70 = ((_235_v0) ? (_242_v4) : (_module.__default.fm0(_241_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_235_v0,!(false))).length), _238_globalState)));
          let _rhs71 = _235_v0;
          let _lhs38 = _237_v2;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
          _384_cf3 = _rhs69;
          _242_v4 = _rhs70;
          _lhs38[_lhs39] = _rhs71;
          let _408_v143;
          let _init9 = function (_409_i14) {
            return _module.__default.safeDivisionInt(_409_i14, new BigNumber(724));
          };
          let _nw92 = Array((new BigNumber(21)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw92.length); _i0_9++) {
            _nw92[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _408_v143 = _nw92;
          let _index60 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_408_v143).length));
          (_408_v143)[_index60] = _385_cf2;
          let _410_v144;
          _410_v144 = _dafny.MultiSet.fromElements(_385_cf2);
          let _index61 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_408_v143).length));
          (_408_v143)[_index61] = new BigNumber((((_410_v144).Union(_410_v144)).Difference((_410_v144).Difference(_410_v144))).cardinality());
        } else {
          _311_v63 = _311_v63;
          let _411_v145;
          let _nw93 = new _module.C8();
          _nw93.__ctor();
          _411_v145 = _nw93;
          _411_v145 = _411_v145;
          let _412_v146;
          let _nw94 = new _module.C7();
          _nw94.__ctor();
          _412_v146 = _nw94;
          let _413_v147;
          let _nw95 = new _module.C7();
          _nw95.__ctor();
          _413_v147 = _nw95;
          let _414_v148;
          let _nw96 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _414_v148 = _nw96;
          let _415_v149;
          _415_v149 = _dafny.Set.fromElements(_237_v2, _237_v2);
          let _416_v150;
          _416_v150 = _415_v149;
          let _417_v151;
          _417_v151 = _dafny.Map.Empty.slice().updateUnsafe(_385_cf2,_235_v0);
          let _418_v154;
          let _nw97 = Array((new BigNumber(22)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _417_v151;
          _nw97[(_dafny.ONE).toNumber()] = _417_v151;
          _nw97[(new BigNumber(2)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(3)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(4)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(5)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(6)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(7)).toNumber()] = (_417_v151).update(new BigNumber(702), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]);
          _nw97[(new BigNumber(8)).toNumber()] = (_417_v151).update((_dafny.ZERO).minus(_241_v3), !(_235_v0));
          _nw97[(new BigNumber(9)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(10)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(11)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(12)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm11(_238_globalState)).cardinality()),_235_v0);
          _nw97[(new BigNumber(14)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_241_v3,_235_v0);
          _nw97[(new BigNumber(16)).toNumber()] = function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(976), new BigNumber(470))) {
              let _419_v152 = _compr_19;
              if (((new BigNumber(976)).isLessThanOrEqualTo(_419_v152)) && ((_419_v152).isLessThan(new BigNumber(470)))) {
                _coll19.push([(_419_v152).plus(_384_cf3),(_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]]);
              }
            }
            return _coll19;
          }();
          _nw97[(new BigNumber(17)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(18)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(19)).toNumber()] = _417_v151;
          _nw97[(new BigNumber(20)).toNumber()] = function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(410), new BigNumber(813))) {
              let _420_v153 = _compr_20;
              if (((new BigNumber(410)).isLessThanOrEqualTo(_420_v153)) && ((_420_v153).isLessThan(new BigNumber(813)))) {
                _coll20.push([(_420_v153).plus(_384_cf3),_235_v0]);
              }
            }
            return _coll20;
          }();
          _nw97[(new BigNumber(21)).toNumber()] = _417_v151;
          _418_v154 = _nw97;
          let _421_v155;
          _421_v155 = _module.D4.create_DC10(_416_v150, _384_cf3, _311_v63, _418_v154);
          let _rhs72 = (_421_v155).dtor_cf13;
          let _rhs73 = _384_cf3;
          let _rhs74 = _412_v146;
          let _rhs75 = _413_v147;
          let _rhs76 = _414_v148;
          _311_v63 = _rhs72;
          _384_cf3 = _rhs73;
          _412_v146 = _rhs74;
          _413_v147 = _rhs75;
          _414_v148 = _rhs76;
          let _422_v156;
          let _nw98 = new _module.C0();
          _nw98.__ctor();
          _422_v156 = _nw98;
          let _423_v157;
          _423_v157 = _dafny.Map.Empty.slice().updateUnsafe((((_351_v100).contains(_235_v0)) ? ((_351_v100).get(_235_v0)) : (_384_cf3)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uqvlgxso")).length),!(_235_v0)));
          let _424_v158;
          let _nw99 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
          _424_v158 = _nw99;
          let _425_v159;
          _425_v159 = _module.D20.create_DC47(_424_v158);
          let _426_v160;
          let _nw100 = new _module.C11();
          _nw100.__ctor(_423_v157, _384_cf3, (_425_v159).dtor_cf73, (_243_v5)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_243_v5).length))]);
          _426_v160 = _nw100;
        }
      } else {
        let _427___mcc_h2 = (_source6).cf1;
        let _428_cf1 = _427___mcc_h2;
        _235_v0 = !(!((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]) || ((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))])) || (!_dafny.areEqual(_dafny.Seq.of(new BigNumber((_dafny.Seq.update(_311_v63, _module.__default.safeIndex(_241_v3, new BigNumber((_311_v63).length)), _242_v4)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-668)), ((_429_v3) => function (_430_i15) {
          return _429_v3;
        })(_241_v3))));
        let _431_v161;
        _431_v161 = _dafny.MultiSet.fromElements(_241_v3, _241_v3);
        let _432_v162;
        _432_v162 = _module.D6.create_DC14(_244_v6);
        _431_v161 = ((_431_v161).Union(_dafny.MultiSet.FromArray((_432_v162).dtor_cf19))).update(new BigNumber((_299_v55).length), _module.__default.abs(_module.__default.fm13(_235_v0, !(_module.__default.fm1(_241_v3, _241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState)), _module.__default.fm1(_241_v3, _241_v3, (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState), _238_globalState)));
        let _433_v163;
        _433_v163 = _dafny.Map.Empty.slice().updateUnsafe(_242_v4,_241_v3);
        _module.__default.m0(new BigNumber((_433_v163).length), (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _238_globalState);
        let _434_v164;
        _434_v164 = _dafny.Map.Empty.slice().updateUnsafe(_235_v0,_module.__default.fm53(_238_globalState));
        let _435_v165;
        let _nw101 = Array((new BigNumber(11)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = _241_v3;
        _nw101[(_dafny.ONE).toNumber()] = _241_v3;
        _nw101[(new BigNumber(2)).toNumber()] = _241_v3;
        _nw101[(new BigNumber(3)).toNumber()] = new BigNumber((_434_v164).length);
        _nw101[(new BigNumber(4)).toNumber()] = _241_v3;
        _nw101[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("gy")).length);
        _nw101[(new BigNumber(6)).toNumber()] = _241_v3;
        _nw101[(new BigNumber(7)).toNumber()] = _241_v3;
        _nw101[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))])).cardinality());
        _nw101[(new BigNumber(9)).toNumber()] = new BigNumber(670);
        _nw101[(new BigNumber(10)).toNumber()] = _241_v3;
        _435_v165 = _nw101;
        let _436_v166;
        _436_v166 = _dafny.Map.Empty.slice().updateUnsafe(_299_v55,_435_v165);
        _436_v166 = (_436_v166).update(_299_v55, _435_v165);
      }
      let _437_v167;
      _437_v167 = _module.D14.create_DC31(_241_v3);
      let _438_v168;
      _438_v168 = _module.D14.create_DC33(_437_v167);
      let _pat_let_tv3 = _241_v3;
      let _pat_let_tv4 = _312_v64;
      let _pat_let_tv5 = _237_v2;
      let _pat_let_tv6 = _237_v2;
      let _pat_let_tv7 = _312_v64;
      let _pat_let_tv8 = _237_v2;
      let _pat_let_tv9 = _237_v2;
      let _pat_let_tv10 = _241_v3;
      let _pat_let_tv11 = _312_v64;
      let _pat_let_tv12 = _237_v2;
      let _pat_let_tv13 = _237_v2;
      let _pat_let_tv14 = _312_v64;
      let _pat_let_tv15 = _237_v2;
      let _pat_let_tv16 = _237_v2;
      let _pat_let_tv17 = _235_v0;
      let _pat_let_tv18 = _241_v3;
      let _source7 = function (_source8) {
        if (_source8.is_DC31) {
          let _439___mcc_h6 = (_source8).cf48;
          let _440_cf48 = _439___mcc_h6;
          return _module.D15.create_DC35(_440_cf48);
        } else if (_source8.is_DC32) {
          let _441___mcc_h7 = (_source8).cf49;
          let _442___mcc_h8 = (_source8).cf50;
          let _443___mcc_h9 = (_source8).cf51;
          let _444_cf51 = _443___mcc_h9;
          let _445_cf50 = _442___mcc_h8;
          let _446_cf49 = _441___mcc_h7;
          return _module.D15.create_DC35(_pat_let_tv3);
        } else if (_source8.is_DC30) {
          let _447___mcc_h10 = (_source8).cf47;
          let _448_cf47 = _447___mcc_h10;
          return _module.D15.create_DC35(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((((_pat_let_tv7).contains((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_pat_let_tv8).length))])) ? ((_pat_let_tv4).get((_pat_let_tv6)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_pat_let_tv5).length))])) : (false))), _module.__default.safeIndex(_pat_let_tv10, new BigNumber((_dafny.Seq.of((((_pat_let_tv14).contains((_pat_let_tv16)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_pat_let_tv15).length))])) ? ((_pat_let_tv11).get((_pat_let_tv13)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_pat_let_tv12).length))])) : (false)))).length)), _pat_let_tv17)).length));
        } else {
          let _449___mcc_h11 = (_source8).cf52;
          let _450_cf52 = _449___mcc_h11;
          return _module.D15.create_DC35(_pat_let_tv18);
        }
      }(_438_v168);
      if (_source7.is_DC35) {
        let _451___mcc_h3 = (_source7).cf54;
        let _452_cf54 = _451___mcc_h3;
        let _453_v169;
        _453_v169 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_244_v6)).cardinality()), ((((_310_v62).f10).contains(_452_cf54)) ? (((_310_v62).f10).get(_452_cf54)) : (new BigNumber((_311_v63).length))));
        let _454_v170;
        _454_v170 = _dafny.MultiSet.fromElements(_452_cf54, _241_v3, (((_453_v169).contains(_241_v3)) ? ((_453_v169).get(_241_v3)) : (_241_v3)));
        let _455_v171;
        _455_v171 = _dafny.Seq.of(_454_v170, _454_v170);
        let _456_v172;
        let _nw102 = Array((new BigNumber(4)).toNumber());
        _nw102[(_dafny.ZERO).toNumber()] = new BigNumber(((_455_v171)[_module.__default.safeIndex(_241_v3, new BigNumber((_455_v171).length))]).cardinality());
        _nw102[(_dafny.ONE).toNumber()] = _452_cf54;
        _nw102[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_311_v63).length), _452_cf54);
        _nw102[(new BigNumber(3)).toNumber()] = _241_v3;
        _456_v172 = _nw102;
        let _index62 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_456_v172).length));
        (_456_v172)[_index62] = _241_v3;
        let _index63 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_456_v172).length));
        (_456_v172)[_index63] = (_452_cf54).multipliedBy((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((((_351_v100).contains(_235_v0)) ? ((_351_v100).get(_235_v0)) : ((_244_v6)[_module.__default.safeIndex(_452_cf54, new BigNumber((_244_v6).length))]))), _452_cf54)));
        let _index64 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_456_v172).length));
        let _rhs77 = new BigNumber(718);
        let _rhs78 = _299_v55;
        let _lhs40 = _456_v172;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_456_v172).length));
        _lhs40[_lhs41] = _rhs77;
        _299_v55 = _rhs78;
        let _457_v173;
        _457_v173 = _dafny.Seq.of(_236_v1, _dafny.Set.fromElements((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _235_v0), _236_v1);
        let _458_v174;
        _458_v174 = _dafny.Seq.of(_244_v6, _244_v6);
        let _459_v175;
        _459_v175 = _dafny.MultiSet.fromElements(!_dafny.areEqual(_458_v174, _458_v174), !((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))]), _235_v0);
        let _460_v176;
        _460_v176 = _module.D2.create_DC5(_311_v63);
        let _461_v177;
        _461_v177 = _dafny.MultiSet.fromElements(_module.D2.create_DC5(_311_v63), _460_v176, _460_v176, _460_v176);
        let _index65 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        let _rhs79 = _452_cf54;
        let _rhs80 = _457_v173;
        let _rhs81 = (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))];
        let _rhs82 = ((_456_v172)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_456_v172).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus(_241_v3));
        let _rhs83 = ((_459_v175).update((_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))], _module.__default.abs(new BigNumber((_461_v177).cardinality())))).Union(_459_v175);
        let _lhs42 = _237_v2;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        _452_cf54 = _rhs79;
        _457_v173 = _rhs80;
        _lhs42[_lhs43] = _rhs81;
        _235_v0 = _rhs82;
        _459_v175 = _rhs83;
        let _index66 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        (_237_v2)[_index66] = (_237_v2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length))];
      } else if (_source7.is_DC34) {
        let _462___mcc_h4 = (_source7).cf53;
        let _463_cf53 = _462___mcc_h4;
        _241_v3 = new BigNumber(748);
        (_238_globalState).f3 = _237_v2;
        let _index67 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        let _rhs84 = _310_v62;
        let _rhs85 = (_dafny.Seq.IsPrefixOf(_311_v63, _module.__default.fm12(_238_globalState))) === (_dafny.Seq.contains(_311_v63, new _dafny.CodePoint('d'.codePointAt(0))));
        let _rhs86 = _241_v3;
        let _rhs87 = _311_v63;
        let _lhs44 = _237_v2;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_237_v2).length));
        _310_v62 = _rhs84;
        _lhs44[_lhs45] = _rhs85;
        _241_v3 = _rhs86;
        _311_v63 = _rhs87;
        let _464_v178;
        let _nw103 = new _module.C7();
        _nw103.__ctor();
        _464_v178 = _nw103;
      } else {
        let _465___mcc_h5 = (_source7).cf55;
        let _466_cf55 = _465___mcc_h5;
        _241_v3 = (_241_v3).plus((new BigNumber(((_dafny.MultiSet.fromElements(_235_v0)).update(_235_v0, _module.__default.abs(_241_v3))).cardinality())).minus(_241_v3));
        let _467_v179;
        _467_v179 = _dafny.Seq.of(_312_v64);
        let _468_v180;
        let _469_v181;
        let _470_v182;
        let _471_v183;
        let _out6;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector2 = (_310_v62).m9(_dafny.Seq.Concat(_467_v179, _467_v179), !((_299_v55)[_module.__default.safeIndex(_241_v3, new BigNumber((_299_v55).length))]), _238_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _out9 = _outcollector2[3];
        _468_v180 = _out6;
        _469_v181 = _out7;
        _470_v182 = _out8;
        _471_v183 = _out9;
        let _472_v184;
        let _nw104 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _472_v184 = _nw104;
        let _473_v185;
        _473_v185 = _dafny.Seq.of(_472_v184);
        _473_v185 = _473_v185;
        _470_v182 = _235_v0;
      }
      process.stdout.write(_dafny.toString(_235_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v2)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_238_globalState).f0).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_238_globalState.f2, _dafny.Seq.of(_dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE, _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_globalState.f3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_241_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_242_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_244_v6, _dafny.Seq.of(new BigNumber(-263)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_299_v55, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_307_v60).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(-263))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_v61).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-494),new BigNumber(-4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_310_v62).f10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-494),new BigNumber(-4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_311_v63).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v64).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_349_v98).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_350_v99).dtor_cf68)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v99).dtor_cf69));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v99).dtor_cf70));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v99).dtor_cf71));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_351_v100).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(199)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_381_v124).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-263), new BigNumber(-721), new BigNumber(-721), new BigNumber(-721), new BigNumber(-721)),new BigNumber(-721)).updateUnsafe(_dafny.Seq.of(new BigNumber(-263)),new BigNumber(-721)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_437_v167).dtor_cf48));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_438_v168).dtor_cf52).dtor_cf48));
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
    static create_DC2(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC4(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC6(cf7) {
      let $dt = new D2(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D2(3);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC7(cf8) {
      let $dt = new D2(4);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get is_DC7() { return this.$tag === 4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + this.cf6.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 4) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf8) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf4 === other.cf4;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(false);
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
    static create_DC8(cf9) {
      let $dt = new D3(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf11, cf12, cf13, cf14) {
      let $dt = new D4(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC9(cf10) {
      let $dt = new D4(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + this.cf13.toVerbatimString(true) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf10 === other.cf10;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(_dafny.Set.Empty, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), []);
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
    static create_DC12() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC13(cf16, cf17, cf18) {
      let $dt = new D5(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC11(cf15) {
      let $dt = new D5(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + this.cf16.toVerbatimString(true) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12();
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
    static create_DC15(cf20, cf21, cf22) {
      let $dt = new D6(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf19) {
      let $dt = new D6(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(false, false, false);
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
    static create_DC17(cf24, cf25, cf26, cf27) {
      let $dt = new D7(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC16(cf23) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(false, false, false, false);
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
    static create_DC19(cf29, cf30) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D8(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC19(_dafny.ZERO, false);
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
    static create_DC20(cf31) {
      let $dt = new D9(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
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
    static create_DC22(cf33, cf34, cf35) {
      let $dt = new D10(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC21(cf32) {
      let $dt = new D10(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC23(cf36) {
      let $dt = new D10(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + this.cf33.toVerbatimString(true) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(_dafny.Seq.UnicodeFromString(""), [], _dafny.ZERO);
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
    static create_DC25(cf38, cf39, cf40, cf41) {
      let $dt = new D11(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC24(cf37) {
      let $dt = new D11(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC26(cf42) {
      let $dt = new D11(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(false, _dafny.ZERO, [], _dafny.ZERO);
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
    static create_DC27(cf43) {
      let $dt = new D12(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43;
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf45, cf46) {
      let $dt = new D13(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D13(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf44 === other.cf44;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC29(false, false);
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
    static create_DC31(cf48) {
      let $dt = new D14(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC32(cf49, cf50, cf51) {
      let $dt = new D14(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC30(cf47) {
      let $dt = new D14(2);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC33(cf52) {
      let $dt = new D14(3);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC33" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && this.cf51 === other.cf51;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC31(_dafny.ZERO);
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
    static create_DC35(cf54) {
      let $dt = new D15(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC34(cf53) {
      let $dt = new D15(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC36(cf55) {
      let $dt = new D15(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
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
    static create_DC38(cf57, cf58, cf59, cf60) {
      let $dt = new D16(0);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC37(cf56) {
      let $dt = new D16(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC39(cf61) {
      let $dt = new D16(2);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC38" + "(" + this.cf57.toVerbatimString(true) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC37" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC38(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO, false);
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
    static create_DC41(cf63) {
      let $dt = new D17(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC42(cf64, cf65) {
      let $dt = new D17(1);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC40(cf62) {
      let $dt = new D17(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC43(cf66) {
      let $dt = new D17(3);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf64 === other.cf64 && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC41(_dafny.ZERO);
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
    static create_DC45(cf68, cf69, cf70, cf71) {
      let $dt = new D18(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC44(cf67) {
      let $dt = new D18(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC45" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf68 === other.cf68 && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC45([], false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC46(cf72) {
      let $dt = new D19(0);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf72, other.cf72);
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48(cf74, cf75, cf76, cf77) {
      let $dt = new D20(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC47(cf73) {
      let $dt = new D20(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC49(cf78) {
      let $dt = new D20(2);
      $dt.cf78 = cf78;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf78() { return this.cf78; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC47" + "(" + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf78) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf73 === other.cf73;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf78, other.cf78);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC48(false, _dafny.Seq.of(), _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC51(cf80, cf81) {
      let $dt = new D21(0);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC50(cf79) {
      let $dt = new D21(1);
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC52(cf82) {
      let $dt = new D21(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC50" + "(" + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf80 === other.cf80 && this.cf81 === other.cf81;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC51(false, null);
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
    static create_DC53(cf83) {
      let $dt = new D22(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf83) + ")";
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
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf85, cf86, cf87, cf88, cf89) {
      let $dt = new D23(0);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC56(cf90, cf91, cf92, cf93) {
      let $dt = new D23(1);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC57() {
      let $dt = new D23(2);
      return $dt;
    }
    static create_DC54(cf84) {
      let $dt = new D23(3);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC55" + "(" + this.cf85.toVerbatimString(true) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC56" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC57";
      } else if (this.$tag === 3) {
        return "D23.DC54" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85) && this.cf86 === other.cf86 && this.cf87 === other.cf87 && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf90 === other.cf90 && this.cf91 === other.cf91 && this.cf92 === other.cf92 && this.cf93 === other.cf93;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC55(_dafny.Seq.UnicodeFromString(""), false, false, false, _dafny.Seq.of());
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
    static create_DC58(cf94) {
      let $dt = new D24(0);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC58" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf94, other.cf94);
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
    static create_DC60(cf96, cf97, cf98) {
      let $dt = new D25(0);
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC59(cf95) {
      let $dt = new D25(1);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC60" + "(" + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ", " + this.cf98.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC59" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf96 === other.cf96 && this.cf97 === other.cf97 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf95, other.cf95);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC60(false, false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC62(cf100, cf101) {
      let $dt = new D26(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      return $dt;
    }
    static create_DC63(cf102, cf103, cf104, cf105) {
      let $dt = new D26(1);
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC61(cf99) {
      let $dt = new D26(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC62" + "(" + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC63" + "(" + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC61" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf100, other.cf100) && _dafny.areEqual(this.cf101, other.cf101);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf102 === other.cf102 && _dafny.areEqual(this.cf103, other.cf103) && _dafny.areEqual(this.cf104, other.cf104) && this.cf105 === other.cf105;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC62(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC65(cf107, cf108, cf109, cf110) {
      let $dt = new D27(0);
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      $dt.cf110 = cf110;
      return $dt;
    }
    static create_DC66(cf111, cf112) {
      let $dt = new D27(1);
      $dt.cf111 = cf111;
      $dt.cf112 = cf112;
      return $dt;
    }
    static create_DC67(cf113, cf114, cf115) {
      let $dt = new D27(2);
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC64(cf106) {
      let $dt = new D27(3);
      $dt.cf106 = cf106;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get is_DC66() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get is_DC64() { return this.$tag === 3; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf106() { return this.cf106; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC65" + "(" + _dafny.toString(this.cf107) + ", " + this.cf108.toVerbatimString(true) + ", " + _dafny.toString(this.cf109) + ", " + _dafny.toString(this.cf110) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC66" + "(" + _dafny.toString(this.cf111) + ", " + _dafny.toString(this.cf112) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC67" + "(" + _dafny.toString(this.cf113) + ", " + _dafny.toString(this.cf114) + ", " + _dafny.toString(this.cf115) + ")";
      } else if (this.$tag === 3) {
        return "D27.DC64" + "(" + _dafny.toString(this.cf106) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf107 === other.cf107 && _dafny.areEqual(this.cf108, other.cf108) && this.cf109 === other.cf109 && this.cf110 === other.cf110;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf111, other.cf111) && _dafny.areEqual(this.cf112, other.cf112);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf113, other.cf113) && this.cf114 === other.cf114 && _dafny.areEqual(this.cf115, other.cf115);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf106, other.cf106);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC65(false, _dafny.Seq.UnicodeFromString(""), false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC69(cf117, cf118, cf119) {
      let $dt = new D28(0);
      $dt.cf117 = cf117;
      $dt.cf118 = cf118;
      $dt.cf119 = cf119;
      return $dt;
    }
    static create_DC68(cf116) {
      let $dt = new D28(1);
      $dt.cf116 = cf116;
      return $dt;
    }
    static create_DC70(cf120) {
      let $dt = new D28(2);
      $dt.cf120 = cf120;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get is_DC68() { return this.$tag === 1; }
    get is_DC70() { return this.$tag === 2; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf120() { return this.cf120; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC69" + "(" + _dafny.toString(this.cf117) + ", " + _dafny.toString(this.cf118) + ", " + _dafny.toString(this.cf119) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC68" + "(" + _dafny.toString(this.cf116) + ")";
      } else if (this.$tag === 2) {
        return "D28.DC70" + "(" + _dafny.toString(this.cf120) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf117, other.cf117) && this.cf118 === other.cf118 && _dafny.areEqual(this.cf119, other.cf119);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf116, other.cf116);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf120, other.cf120);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC69(_dafny.ZERO, [], new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC71(cf121) {
      let $dt = new D29(0);
      $dt.cf121 = cf121;
      return $dt;
    }
    get is_DC71() { return this.$tag === 0; }
    get dtor_cf121() { return this.cf121; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC71" + "(" + _dafny.toString(this.cf121) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf121 === other.cf121;
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
          return D29.Default();
        }
      };
    }
  }

  $module.D30 = class D30 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC73(cf123, cf124, cf125) {
      let $dt = new D30(0);
      $dt.cf123 = cf123;
      $dt.cf124 = cf124;
      $dt.cf125 = cf125;
      return $dt;
    }
    static create_DC72(cf122) {
      let $dt = new D30(1);
      $dt.cf122 = cf122;
      return $dt;
    }
    get is_DC73() { return this.$tag === 0; }
    get is_DC72() { return this.$tag === 1; }
    get dtor_cf123() { return this.cf123; }
    get dtor_cf124() { return this.cf124; }
    get dtor_cf125() { return this.cf125; }
    get dtor_cf122() { return this.cf122; }
    toString() {
      if (this.$tag === 0) {
        return "D30.DC73" + "(" + _dafny.toString(this.cf123) + ", " + _dafny.toString(this.cf124) + ", " + _dafny.toString(this.cf125) + ")";
      } else if (this.$tag === 1) {
        return "D30.DC72" + "(" + _dafny.toString(this.cf122) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf123, other.cf123) && this.cf124 === other.cf124 && this.cf125 === other.cf125;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf122 === other.cf122;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D30.create_DC73(_dafny.ZERO, false, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D30.Default();
        }
      };
    }
  }

  $module.D31 = class D31 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC75(cf127) {
      let $dt = new D31(0);
      $dt.cf127 = cf127;
      return $dt;
    }
    static create_DC74(cf126) {
      let $dt = new D31(1);
      $dt.cf126 = cf126;
      return $dt;
    }
    static create_DC76(cf128) {
      let $dt = new D31(2);
      $dt.cf128 = cf128;
      return $dt;
    }
    get is_DC75() { return this.$tag === 0; }
    get is_DC74() { return this.$tag === 1; }
    get is_DC76() { return this.$tag === 2; }
    get dtor_cf127() { return this.cf127; }
    get dtor_cf126() { return this.cf126; }
    get dtor_cf128() { return this.cf128; }
    toString() {
      if (this.$tag === 0) {
        return "D31.DC75" + "(" + _dafny.toString(this.cf127) + ")";
      } else if (this.$tag === 1) {
        return "D31.DC74" + "(" + _dafny.toString(this.cf126) + ")";
      } else if (this.$tag === 2) {
        return "D31.DC76" + "(" + _dafny.toString(this.cf128) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf127 === other.cf127;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf126, other.cf126);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf128, other.cf128);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D31.create_DC75(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D31.Default();
        }
      };
    }
  }

  $module.D32 = class D32 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC78(cf130, cf131, cf132, cf133, cf134) {
      let $dt = new D32(0);
      $dt.cf130 = cf130;
      $dt.cf131 = cf131;
      $dt.cf132 = cf132;
      $dt.cf133 = cf133;
      $dt.cf134 = cf134;
      return $dt;
    }
    static create_DC79(cf135, cf136, cf137) {
      let $dt = new D32(1);
      $dt.cf135 = cf135;
      $dt.cf136 = cf136;
      $dt.cf137 = cf137;
      return $dt;
    }
    static create_DC77(cf129) {
      let $dt = new D32(2);
      $dt.cf129 = cf129;
      return $dt;
    }
    static create_DC80(cf138) {
      let $dt = new D32(3);
      $dt.cf138 = cf138;
      return $dt;
    }
    get is_DC78() { return this.$tag === 0; }
    get is_DC79() { return this.$tag === 1; }
    get is_DC77() { return this.$tag === 2; }
    get is_DC80() { return this.$tag === 3; }
    get dtor_cf130() { return this.cf130; }
    get dtor_cf131() { return this.cf131; }
    get dtor_cf132() { return this.cf132; }
    get dtor_cf133() { return this.cf133; }
    get dtor_cf134() { return this.cf134; }
    get dtor_cf135() { return this.cf135; }
    get dtor_cf136() { return this.cf136; }
    get dtor_cf137() { return this.cf137; }
    get dtor_cf129() { return this.cf129; }
    get dtor_cf138() { return this.cf138; }
    toString() {
      if (this.$tag === 0) {
        return "D32.DC78" + "(" + _dafny.toString(this.cf130) + ", " + _dafny.toString(this.cf131) + ", " + _dafny.toString(this.cf132) + ", " + _dafny.toString(this.cf133) + ", " + _dafny.toString(this.cf134) + ")";
      } else if (this.$tag === 1) {
        return "D32.DC79" + "(" + _dafny.toString(this.cf135) + ", " + _dafny.toString(this.cf136) + ", " + _dafny.toString(this.cf137) + ")";
      } else if (this.$tag === 2) {
        return "D32.DC77" + "(" + _dafny.toString(this.cf129) + ")";
      } else if (this.$tag === 3) {
        return "D32.DC80" + "(" + _dafny.toString(this.cf138) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf130 === other.cf130 && _dafny.areEqual(this.cf131, other.cf131) && this.cf132 === other.cf132 && _dafny.areEqual(this.cf133, other.cf133) && this.cf134 === other.cf134;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf135 === other.cf135 && this.cf136 === other.cf136 && this.cf137 === other.cf137;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf129, other.cf129);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf138, other.cf138);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D32.create_DC78(false, _dafny.ZERO, false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D32.Default();
        }
      };
    }
  }

  $module.D33 = class D33 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC81(cf139) {
      let $dt = new D33(0);
      $dt.cf139 = cf139;
      return $dt;
    }
    get is_DC81() { return this.$tag === 0; }
    get dtor_cf139() { return this.cf139; }
    toString() {
      if (this.$tag === 0) {
        return "D33.DC81" + "(" + _dafny.toString(this.cf139) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf139, other.cf139);
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
          return D33.Default();
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
      this.f2 = _dafny.Seq.of();
      this.f3 = [];
      this._f0 = _dafny.Set.Empty;
      this._f1 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
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
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return (((false) ? (_dafny.MultiSet.fromElements(true, !(!(false)), true, true, false)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(true), !(false)))))).IsDisjointFrom((_dafny.MultiSet.fromElements(!(false), true, false)).Intersect(_dafny.MultiSet.fromElements(true)));
    };
    fm17(p0, p1, p2, globalState) {
      let _this = this;
      return _module.D8.create_DC18((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(224),!(true))).length), new BigNumber((_dafny.Seq.UnicodeFromString("phmadbtnq")).length))).Difference((_module.D8.create_DC18(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tlf")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length)))).dtor_cf28));
    };
    fm18(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_474_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("aaymeth")).length));
      }), _dafny.Seq.of(new BigNumber(-14), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(203),new BigNumber(-39))).length))))).length)).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber(686), new BigNumber(737)));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-466), new BigNumber((function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gpqxlhwy"),new BigNumber(356))).Keys.Elements) {
          let _475_v0 = _compr_21;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gpqxlhwy"),new BigNumber(356))).contains(_475_v0)) {
            _coll21.push([_475_v0,!(false)]);
          }
        }
        return _coll21;
      }()).length)))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pwts"), _dafny.Seq.UnicodeFromString("g"))).length));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return true;
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let _476_v0;
      let _nw105 = Array((new BigNumber(20)).toNumber());
      _476_v0 = _nw105;
      let _477_v1;
      _477_v1 = _476_v0;
      let _478_v2;
      _478_v2 = true;
      let _479_v3;
      let _nw106 = Array((new BigNumber(25)).toNumber());
      _nw106[(_dafny.ZERO).toNumber()] = _476_v0;
      _nw106[(_dafny.ONE).toNumber()] = _476_v0;
      _nw106[(new BigNumber(2)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(3)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(4)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(5)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(6)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(7)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(8)).toNumber()] = (_477_v1);
      _nw106[(new BigNumber(9)).toNumber()] = ((_478_v2) ? (_476_v0) : (_476_v0));
      _nw106[(new BigNumber(10)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(11)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(12)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(13)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(14)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(15)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(16)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(17)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(18)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(19)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(20)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(21)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(22)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(23)).toNumber()] = _476_v0;
      _nw106[(new BigNumber(24)).toNumber()] = _476_v0;
      _479_v3 = _nw106;
      let _index68 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_479_v3).length));
      (_479_v3)[_index68] = _476_v0;
      let _480_v4;
      let _nw107 = new _module.C0();
      _nw107.__ctor();
      _480_v4 = _nw107;
      let _index69 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_479_v3).length));
      let _nw108 = Array((new BigNumber(7)).toNumber());
      _nw108[(_dafny.ZERO).toNumber()] = _480_v4;
      _nw108[(_dafny.ONE).toNumber()] = _480_v4;
      _nw108[(new BigNumber(2)).toNumber()] = _480_v4;
      _nw108[(new BigNumber(3)).toNumber()] = _480_v4;
      _nw108[(new BigNumber(4)).toNumber()] = _480_v4;
      _nw108[(new BigNumber(5)).toNumber()] = _480_v4;
      _nw108[(new BigNumber(6)).toNumber()] = _480_v4;
      (_479_v3)[_index69] = _nw108;
      let _481_i0;
      _481_i0 = _dafny.ZERO;
      L1: {
        while (_478_v2) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_481_i0)) {
              break L1;
            }
            _481_i0 = (_481_i0).plus(_dafny.ONE);
            let _482_v5;
            let _nw109 = Array((new BigNumber(29)).toNumber()).fill(_module.D11.Default());
            _482_v5 = _nw109;
            let _483_v6;
            let _init10 = function (_484_i1) {
              return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ewp")).length));
            };
            let _nw110 = Array((new BigNumber(22)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw110.length); _i0_10++) {
              _nw110[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _483_v6 = _nw110;
            let _index70 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_482_v5).length));
            (_482_v5)[_index70] = _module.D11.create_DC24(_483_v6);
            let _485_v7;
            _485_v7 = _module.D11.create_DC24(_483_v6);
            let _index71 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_482_v5).length));
            (_482_v5)[_index71] = _485_v7;
            let _486_v8;
            _486_v8 = new BigNumber(6);
            r1 = (_486_v8).minus(_486_v8);
            let _487_v9;
            let _nw111 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
            _487_v9 = _nw111;
            let _488_v10;
            _488_v10 = _module.D4.create_DC9(_487_v9);
            let _489_v11;
            _489_v11 = _dafny.Seq.of(_478_v2);
            _488_v10 = (((_486_v8).isLessThan(new BigNumber((_489_v11).length))) ? (_488_v10) : (_488_v10));
            _478_v2 = false;
          }
        }
      }
      let _490_v12;
      let _init11 = function (_491_i2) {
        return (_491_i2).minus(new BigNumber(-463));
      };
      let _nw112 = Array((new BigNumber(13)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw112.length); _i0_11++) {
        _nw112[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _490_v12 = _nw112;
      let _492_v13;
      _492_v13 = _dafny.Map.Empty.slice().updateUnsafe(_478_v2,_490_v12);
      _492_v13 = ((_478_v2) ? (_492_v13) : (_492_v13));
      let _493_v14;
      let _init12 = ((_494_v2) => function (_495_i3) {
        return (_dafny.MultiSet.fromElements(_494_v2)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_494_v2, _494_v2, _494_v2));
      })(_478_v2);
      let _nw113 = Array((new BigNumber(6)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw113.length); _i0_12++) {
        _nw113[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _493_v14 = _nw113;
      (globalState).f3 = _493_v14;
      let _496_v15;
      _496_v15 = new BigNumber(282);
      r1 = _496_v15;
      let _497_v16;
      _497_v16 = _dafny.Map.Empty.slice().updateUnsafe(_478_v2,_478_v2);
      let _498_v17;
      _498_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_497_v16).length),!(_478_v2));
      _498_v17 = (_498_v17).update((_dafny.ZERO).minus(_496_v15), _478_v2);
      let _499_v18;
      _499_v18 = _dafny.Map.Empty.slice().updateUnsafe(_490_v12,_496_v15);
      let _500_v19;
      _500_v19 = _dafny.Seq.of(new BigNumber((_499_v18).length));
      r0 = _dafny.Seq.Concat(_500_v19, _500_v19);
      r1 = _496_v15;
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _501_v0;
      _501_v0 = true;
      let _502_v1;
      _502_v1 = _dafny.Seq.UnicodeFromString("mwx");
      let _503_v2;
      _503_v2 = new BigNumber(425);
      let _504_v3;
      _504_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_502_v1)).length),_503_v2);
      let _505_v4;
      _505_v4 = _dafny.Map.Empty.slice().updateUnsafe(_501_v0,_504_v3);
      let _506_i0;
      _506_i0 = _dafny.ZERO;
      L2: {
        while (!(_505_v4).contains(_501_v0)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_506_i0)) {
              break L2;
            }
            _506_i0 = (_506_i0).plus(_dafny.ONE);
            let _507_v5;
            let _nw114 = new _module.C0();
            _nw114.__ctor();
            _507_v5 = _nw114;
            _503_v2 = (new BigNumber(285)).multipliedBy(_503_v2);
            _501_v0 = _501_v0;
            let _508_v6;
            _508_v6 = _dafny.Seq.of(_503_v2, _503_v2);
            let _509_v7;
            _509_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_508_v6, globalState),_501_v0);
            let _510_v8;
            _510_v8 = _dafny.MultiSet.fromElements(_509_v7);
            if ((_510_v8).IsProperSubsetOf(_510_v8)) {
              let _511_v9;
              _511_v9 = _dafny.MultiSet.fromElements(false);
              let _512_v10;
              _512_v10 = _dafny.Set.fromElements((_this).fm4(_508_v6, globalState), _503_v2, _503_v2, (((_511_v9).contains(_501_v0)) ? ((_511_v9).get(_501_v0)) : ((_dafny.ZERO).minus(new BigNumber((_502_v1).length)))), _503_v2);
              let _513_v11;
              _513_v11 = _dafny.Seq.of(_501_v0);
              let _rhs88 = (_module.__default.fm22((_this).f5, _501_v0, _503_v2, globalState)).Intersect(_512_v10);
              let _rhs89 = !((_513_v11)[_module.__default.safeIndex(_503_v2, new BigNumber((_513_v11).length))]);
              let _rhs90 = ((_501_v0) ? (_503_v2) : (_503_v2));
              _512_v10 = _rhs88;
              _501_v0 = _rhs89;
              _503_v2 = _rhs90;
              let _514_v12;
              _514_v12 = _dafny.Seq.of(_dafny.Seq.update(_513_v11, _module.__default.safeIndex(new BigNumber((_508_v6).length), new BigNumber((_513_v11).length)), _501_v0), _513_v11, _dafny.Seq.of(_501_v0, _501_v0, _501_v0, _501_v0, _501_v0));
              let _515_v13;
              _515_v13 = _dafny.Map.Empty.slice().updateUnsafe(_501_v0,_501_v0);
              let _516_v14;
              _516_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("hcajlchw"), (_this).f5)),new BigNumber((_515_v13).length));
              let _rhs91 = (_513_v11)[_module.__default.safeIndex(_503_v2, new BigNumber((_513_v11).length))];
              let _rhs92 = _module.__default.safeDivisionInt((_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(36)), ((_517_v2) => function (_518_i1) {
                return _517_v2;
              })(_503_v2)), globalState), _503_v2);
              let _rhs93 = _dafny.Seq.of(_513_v11, _513_v11);
              let _rhs94 = _516_v14;
              _501_v0 = _rhs91;
              _503_v2 = _rhs92;
              _514_v12 = _rhs93;
              _516_v14 = _rhs94;
              let _519_v15;
              let _init13 = function (_520_i2) {
                return (_this).f5;
              };
              let _nw115 = Array((new BigNumber(12)).toNumber());
              for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw115.length); _i0_13++) {
                _nw115[_i0_13] = _init13(new BigNumber(_i0_13));
              }
              _519_v15 = _nw115;
              let _521_v16;
              _521_v16 = _module.D10.create_DC22(_dafny.Seq.UnicodeFromString("fenhui"), _519_v15, _503_v2);
              let _522_v17;
              let _init14 = ((_523_v2) => function (_524_i3) {
                return (_524_i3).multipliedBy(_523_v2);
              })(_503_v2);
              let _nw116 = Array((new BigNumber(6)).toNumber());
              for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw116.length); _i0_14++) {
                _nw116[_i0_14] = _init14(new BigNumber(_i0_14));
              }
              _522_v17 = _nw116;
              let _525_v18;
              _525_v18 = _dafny.Map.Empty.slice().updateUnsafe(_522_v17,_503_v2);
              let _526_v19;
              _526_v19 = _dafny.Map.Empty.slice().updateUnsafe(_503_v2,_525_v18);
              _501_v0 = (_module.__default.safeDivisionInt(_503_v2, (_521_v16).dtor_cf35)).isLessThan((_503_v2).minus(new BigNumber(((((_526_v19).contains(_503_v2)) ? ((_526_v19).get(_503_v2)) : ((_525_v18).update(_522_v17, (_dafny.ZERO).minus(new BigNumber((_513_v11).length)))))).length)));
              let _527_v20;
              _527_v20 = _dafny.Map.Empty.slice().updateUnsafe(((_501_v0) ? (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_502_v1).length)))) : (_508_v6)),_501_v0);
              _527_v20 = _527_v20;
              let _528_v21;
              let _nw117 = Array((new BigNumber(16)).toNumber()).fill([]);
              _528_v21 = _nw117;
              let _529_v22;
              let _init15 = ((_530_v0) => function (_531_i4) {
                return _530_v0;
              })(_501_v0);
              let _nw118 = Array((new BigNumber(25)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw118.length); _i0_15++) {
                _nw118[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _529_v22 = _nw118;
              let _index72 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_528_v21).length));
              (_528_v21)[_index72] = _529_v22;
              let _index73 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_528_v21).length));
              (_528_v21)[_index73] = _529_v22;
            } else {
              let _532_v23;
              let _nw119 = new _module.C0();
              _nw119.__ctor();
              _532_v23 = _nw119;
              _502_v1 = _502_v1;
              let _533_v24;
              let _nw120 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _533_v24 = _nw120;
              _533_v24 = _533_v24;
              let _534_v25;
              _534_v25 = _dafny.MultiSet.fromElements((_this).f5, (_this).f5);
              _501_v0 = (_534_v25).equals(_module.__default.fm23(_501_v0, globalState));
              let _535_v26;
              _535_v26 = _dafny.Map.Empty.slice().updateUnsafe(_501_v0,_503_v2);
              _535_v26 = (_535_v26).update(true, _503_v2);
            }
          }
        }
      }
      let _536_v27;
      let _init16 = function (_537_i5) {
        return (_this).f5;
      };
      let _nw121 = Array((new BigNumber(25)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw121.length); _i0_16++) {
        _nw121[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _536_v27 = _nw121;
      let _index74 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_536_v27).length));
      (_536_v27)[_index74] = (_this).f5;
      let _538_v28;
      _538_v28 = _dafny.Seq.of(!(_501_v0), true);
      let _index75 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_536_v27).length));
      (_536_v27)[_index75] = ((_dafny.Seq.contains(_538_v28, _501_v0)) ? ((_this).f5) : ((_this).f5));
      _501_v0 = _501_v0;
      _503_v2 = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_503_v2)).length), _503_v2);
      let _hi5 = _503_v2;
      for (let _539_i6 = new BigNumber((_502_v1).length); _539_i6.isLessThan(_hi5); _539_i6 = _539_i6.plus(_dafny.ONE)) {
        let _540_v29;
        let _nw122 = new _module.C0();
        _nw122.__ctor();
        _540_v29 = _nw122;
        _503_v2 = _503_v2;
        let _541_v30;
        let _nw123 = Array((new BigNumber(23)).toNumber()).fill(false);
        _541_v30 = _nw123;
        let _542_v31;
        _542_v31 = _dafny.Map.Empty.slice().updateUnsafe(_501_v0,_541_v30);
        let _543_v32;
        let _nw124 = Array((new BigNumber(14)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = _541_v30;
        _nw124[(_dafny.ONE).toNumber()] = _541_v30;
        _nw124[(new BigNumber(2)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(3)).toNumber()] = ((_501_v0) ? ((((_542_v31).contains(_501_v0)) ? ((_542_v31).get(_501_v0)) : (_541_v30))) : (_541_v30));
        _nw124[(new BigNumber(4)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(5)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(6)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(7)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(8)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(9)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(10)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(11)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(12)).toNumber()] = _541_v30;
        _nw124[(new BigNumber(13)).toNumber()] = _541_v30;
        _543_v32 = _nw124;
        let _rhs95 = _543_v32;
        let _rhs96 = _541_v30;
        _543_v32 = _rhs95;
        _541_v30 = _rhs96;
        let _index76 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_541_v30).length));
        (_541_v30)[_index76] = _501_v0;
        let _544_v33;
        let _nw125 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
        _544_v33 = _nw125;
        let _545_v34;
        _545_v34 = _module.D13.create_DC28(_544_v33);
        let _546_v35;
        let _nw126 = Array((new BigNumber(6)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = (_545_v34).dtor_cf44;
        _nw126[(_dafny.ONE).toNumber()] = _544_v33;
        _nw126[(new BigNumber(2)).toNumber()] = _544_v33;
        _nw126[(new BigNumber(3)).toNumber()] = _544_v33;
        _nw126[(new BigNumber(4)).toNumber()] = ((_501_v0) ? (_544_v33) : (_544_v33));
        _nw126[(new BigNumber(5)).toNumber()] = _544_v33;
        _546_v35 = _nw126;
        let _index77 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_546_v35).length));
        (_546_v35)[_index77] = _544_v33;
        let _547_v36;
        _547_v36 = _dafny.Seq.of(_544_v33, _544_v33);
        let _index78 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_541_v30).length));
        let _index79 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_546_v35).length));
        let _rhs97 = _503_v2;
        let _rhs98 = _501_v0;
        let _rhs99 = !(!(!(_501_v0)));
        let _rhs100 = ((_501_v0) ? (_544_v33) : ((_547_v36)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_547_v36).length))]));
        let _lhs46 = _541_v30;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_541_v30).length));
        let _lhs48 = _546_v35;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_546_v35).length));
        _503_v2 = _rhs97;
        _lhs46[_lhs47] = _rhs98;
        _501_v0 = _rhs99;
        _lhs48[_lhs49] = _rhs100;
      }
      let _548_v37;
      _548_v37 = _dafny.Map.Empty.slice().updateUnsafe(_501_v0,(_536_v27)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_536_v27).length))]);
      let _549_v38;
      _549_v38 = _module.D5.create_DC11(_548_v37);
      let _550_v39;
      _550_v39 = _dafny.Map.Empty.slice().updateUnsafe(_503_v2,_549_v38);
      let _551_v40;
      _551_v40 = _dafny.Seq.of(_550_v39);
      _551_v40 = _dafny.Seq.Concat(_551_v40, _551_v40);
      let _552_v41;
      let _init17 = ((_553_v0) => function (_554_i7) {
        return (_554_i7).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_553_v0, _553_v0))).cardinality()));
      })(_501_v0);
      let _nw127 = Array((new BigNumber(3)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw127.length); _i0_17++) {
        _nw127[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _552_v41 = _nw127;
      r0 = ((false) ? (_552_v41) : (_552_v41));
      return r0;
    }
    m14(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _555_v0;
      _555_v0 = false;
      let _556_v1;
      _556_v1 = _dafny.Seq.of(p1);
      let _557_v2;
      _557_v2 = _dafny.MultiSet.fromElements(!_dafny.Seq.contains(_556_v1, new BigNumber((_dafny.MultiSet.fromElements(_555_v0, _555_v0)).cardinality())), _555_v0);
      let _558_v3;
      let _init18 = ((_559_v0) => function (_560_i0) {
        return (_559_v0) === (_559_v0);
      })(_555_v0);
      let _nw128 = Array((new BigNumber(12)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw128.length); _i0_18++) {
        _nw128[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _558_v3 = _nw128;
      let _index80 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
      (_558_v3)[_index80] = _555_v0;
      let _561_v4;
      _561_v4 = _dafny.Seq.of(((_555_v0) ? (_557_v2) : (_557_v2)));
      let _562_v5;
      _562_v5 = _dafny.Set.fromElements(_555_v0, !(_555_v0), _555_v0);
      let _563_v6;
      _563_v6 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_562_v5).length))));
      let _index81 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
      let _rhs101 = (_561_v4)[_module.__default.safeIndex((((_563_v6).contains(!(_555_v0))) ? ((_563_v6).get(!(_555_v0))) : ((_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(154)), ((_564_p1) => function (_565_i1) {
        return _564_p1;
      })(p1)), globalState))), new BigNumber((_561_v4).length))];
      let _rhs102 = _555_v0;
      let _lhs50 = _558_v3;
      let _lhs51 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
      _557_v2 = _rhs101;
      _lhs50[_lhs51] = _rhs102;
      if ((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]) {
        let _566_v7;
        _566_v7 = _dafny.Set.fromElements((_this).f5);
        let _567_v8;
        _567_v8 = _dafny.Seq.UnicodeFromString("oyekk");
        let _568_v9;
        let _nw129 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _568_v9 = _nw129;
        let _569_v10;
        _569_v10 = _dafny.Map.Empty.slice().updateUnsafe(_567_v8,_568_v9);
        let _570_v11;
        let _nw130 = Array((new BigNumber(11)).toNumber());
        _nw130[(_dafny.ZERO).toNumber()] = p0;
        _nw130[(_dafny.ONE).toNumber()] = (_this).fm4(_556_v1, globalState);
        _nw130[(new BigNumber(2)).toNumber()] = p0;
        _nw130[(new BigNumber(3)).toNumber()] = (p0).minus(p0);
        _nw130[(new BigNumber(4)).toNumber()] = (p0).minus(p0);
        _nw130[(new BigNumber(5)).toNumber()] = p0;
        _nw130[(new BigNumber(6)).toNumber()] = new BigNumber(((_563_v6).Merge(_563_v6)).length);
        _nw130[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw130[(new BigNumber(8)).toNumber()] = new BigNumber((_566_v7).length);
        _nw130[(new BigNumber(9)).toNumber()] = new BigNumber(((((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]) ? (_569_v10) : ((_569_v10).update(_567_v8, _568_v9)))).length);
        _nw130[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p1);
        _570_v11 = _nw130;
        _570_v11 = ((false) ? (_570_v11) : (_570_v11));
        if (true) {
          let _index82 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index82] = (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))];
          r2 = p0;
          _555_v0 = ((!_dafny.Seq.contains(_567_v8, (_this).f5)) ? (!((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))])) : (true));
          let _index83 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index83] = !(_555_v0) || (!((_this).fm5(p1, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], p1, p1, globalState)));
          let _571_v12;
          _571_v12 = _dafny.Seq.of(_556_v1, _556_v1);
          let _572_v13;
          _572_v13 = _dafny.Set.fromElements(_556_v1, _556_v1, _dafny.Seq.Concat(_556_v1, _556_v1), _dafny.Seq.Concat(_556_v1, (_571_v12)[_module.__default.safeIndex(p0, new BigNumber((_571_v12).length))]), _556_v1);
          _572_v13 = _572_v13;
        } else {
          let _573_v14;
          _573_v14 = _dafny.Set.fromElements(_567_v8, _567_v8, _567_v8);
          _573_v14 = ((((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]) ? (_573_v14) : (function () {
            let _coll22 = new _dafny.Set();
            for (const _compr_22 of (_573_v14).Elements) {
              let _574_v15 = _compr_22;
              if ((_573_v14).contains(_574_v15)) {
                _coll22.add(_574_v15);
              }
            }
            return _coll22;
          }()))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("kw")));
          let _575_v16;
          let _nw131 = new _module.C0();
          _nw131.__ctor();
          _575_v16 = _nw131;
          r1 = p0;
          let _index84 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_570_v11).length));
          (_570_v11)[_index84] = p0;
          let _index85 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_570_v11).length));
          (_570_v11)[_index85] = (_this).fm4(_dafny.Seq.Concat(_556_v1, _556_v1), globalState);
          let _576_v17;
          _576_v17 = _dafny.Map.Empty.slice().updateUnsafe((_570_v11)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_570_v11).length))],((_575_v16).fm17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-198)), function (_577_i2) {
            return new BigNumber(608);
          }), _555_v0, _567_v8, globalState)).dtor_cf28);
          _563_v6 = (_563_v6).update(_555_v0, (_dafny.ZERO).minus(new BigNumber((_576_v17).length)));
        }
        let _578_v19;
        _578_v19 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), ((_579_p1) => function (_580_i3) {
            return _579_p1;
          })(p1))).Elements) {
            let _581_v18 = _compr_23;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), ((_582_p1) => function (_580_i3) {
              return _582_p1;
            })(p1)), _581_v18)) {
              _coll23.add((_581_v18).minus(new BigNumber(492)));
            }
          }
          return _coll23;
        }(),(_dafny.ZERO).minus(new BigNumber((_562_v5).length)));
        _555_v0 = !(_module.__default.fm1(new BigNumber(91), p1, !(new BigNumber((_578_v19).length)).isEqualTo((_dafny.ZERO).minus(p1)), _555_v0, globalState));
        let _583_v20;
        let _init19 = ((_584_v1) => function (_585_i4) {
          return _584_v1;
        })(_556_v1);
        let _nw132 = Array((new BigNumber(17)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw132.length); _i0_19++) {
          _nw132[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _583_v20 = _nw132;
        let _source9 = _module.D13.create_DC28(_583_v20);
        if (_source9.is_DC29) {
          let _586___mcc_h0 = (_source9).cf45;
          let _587___mcc_h1 = (_source9).cf46;
          let _588_cf46 = _587___mcc_h1;
          let _589_cf45 = _586___mcc_h0;
          let _590_v21;
          let _nw133 = new _module.C0();
          _nw133.__ctor();
          _590_v21 = _nw133;
          let _591_v22;
          let _nw134 = Array((new BigNumber(16)).toNumber()).fill([]);
          _591_v22 = _nw134;
          _591_v22 = _591_v22;
          r0 = (_dafny.ZERO).minus(p0);
          let _592_v23;
          _592_v23 = _module.D13.create_DC28(_583_v20);
          let _593_v24;
          _593_v24 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
          let _594_v26;
          _594_v26 = _dafny.Seq.of((function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(945), new BigNumber(999))) {
              let _595_v25 = _compr_24;
              if (((new BigNumber(945)).isLessThanOrEqualTo(_595_v25)) && ((_595_v25).isLessThan(new BigNumber(999)))) {
                _coll24.add(_module.__default.safeModuloInt(_595_v25, p1));
              }
            }
            return _coll24;
          }()).IsSubsetOf(_593_v24), _589_cf45, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _rhs103 = _592_v23;
          let _rhs104 = (_594_v26)[_module.__default.safeIndex(p1, new BigNumber((_594_v26).length))];
          _592_v23 = _rhs103;
          _589_cf45 = _rhs104;
        } else {
          let _596___mcc_h2 = (_source9).cf44;
          let _597_cf44 = _596___mcc_h2;
          let _rhs105 = !((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _rhs106 = (_this).fm3((((_563_v6).contains(_555_v0)) ? ((_563_v6).get(_555_v0)) : ((_dafny.ZERO).minus(p1))), (_dafny.ZERO).minus(p0), globalState);
          let _rhs107 = (((_555_v0) ? ((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]) : (_555_v0))) && ((_this).fm3(p1, p1, globalState));
          _555_v0 = _rhs105;
          r3 = _rhs106;
          r3 = _rhs107;
          let _598_v27;
          _598_v27 = _dafny.Seq.of(_570_v11, _570_v11, _570_v11);
          _570_v11 = (_598_v27)[_module.__default.safeIndex(p0, new BigNumber((_598_v27).length))];
          let _599_v28;
          _599_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          _599_v28 = (_599_v28).update((_this).f5, (p1).isEqualTo(p0));
          r0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_567_v8, _module.__default.safeIndex(p0, new BigNumber((_567_v8).length)), new _dafny.CodePoint('d'.codePointAt(0))), _567_v8)).length);
        }
        let _rhs108 = _module.__default.fm24((_this).fm4(_556_v1, globalState), globalState);
        let _rhs109 = _555_v0;
        _567_v8 = _rhs108;
        _555_v0 = _rhs109;
      } else {
        if (_555_v0) {
          let _index86 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index86] = _555_v0;
          let _600_v29;
          _600_v29 = _dafny.Map.Empty.slice().updateUnsafe(_555_v0,false);
          let _601_v30;
          _601_v30 = _dafny.Map.Empty.slice().updateUnsafe((_600_v29).Merge(_600_v29),_557_v2);
          _601_v30 = _601_v30;
          let _602_v31;
          _602_v31 = _dafny.Seq.UnicodeFromString("laj");
          _602_v31 = _602_v31;
          r2 = _module.__default.safeDivisionInt(p0, p0);
          let _603_v32;
          _603_v32 = _module.D1.create_DC1(_dafny.MultiSet.fromElements(_555_v0));
          let _604_v33;
          _604_v33 = _dafny.MultiSet.fromElements(_603_v32);
          let _605_v34;
          _605_v34 = _dafny.Seq.of(_module.__default.fm1(p1, p1, (_this).fm5(p1, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], new BigNumber(787), new BigNumber(((_dafny.MultiSet.fromElements((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))])).update((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], _module.__default.abs(p1))).cardinality()), globalState), _555_v0, globalState));
          let _606_v35;
          _606_v35 = _dafny.Set.fromElements((_this).f5);
          let _607_v36;
          let _nw135 = Array((new BigNumber(20)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = new BigNumber(65);
          _nw135[(_dafny.ONE).toNumber()] = p1;
          _nw135[(new BigNumber(2)).toNumber()] = (_this).fm4(_556_v1, globalState);
          _nw135[(new BigNumber(3)).toNumber()] = p0;
          _nw135[(new BigNumber(4)).toNumber()] = p1;
          _nw135[(new BigNumber(5)).toNumber()] = p1;
          _nw135[(new BigNumber(6)).toNumber()] = p0;
          _nw135[(new BigNumber(7)).toNumber()] = new BigNumber((_604_v33).cardinality());
          _nw135[(new BigNumber(8)).toNumber()] = p0;
          _nw135[(new BigNumber(9)).toNumber()] = new BigNumber((_557_v2).cardinality());
          _nw135[(new BigNumber(10)).toNumber()] = p0;
          _nw135[(new BigNumber(11)).toNumber()] = p0;
          _nw135[(new BigNumber(12)).toNumber()] = p1;
          _nw135[(new BigNumber(13)).toNumber()] = new BigNumber((_602_v31).length);
          _nw135[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw135[(new BigNumber(15)).toNumber()] = new BigNumber((_605_v34).length);
          _nw135[(new BigNumber(16)).toNumber()] = p1;
          _nw135[(new BigNumber(17)).toNumber()] = new BigNumber(72);
          _nw135[(new BigNumber(18)).toNumber()] = new BigNumber((_606_v35).length);
          _nw135[(new BigNumber(19)).toNumber()] = p1;
          _607_v36 = _nw135;
          let _608_v37;
          _608_v37 = _dafny.Map.Empty.slice().updateUnsafe(_607_v36,_dafny.Seq.IsProperPrefixOf(_605_v34, _605_v34));
          _608_v37 = _608_v37;
        } else {
          r1 = (_dafny.ZERO).minus(new BigNumber((_557_v2).cardinality()));
          let _609_v38;
          let _nw136 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _609_v38 = _nw136;
          let _610_v39;
          _610_v39 = _dafny.Map.Empty.slice().updateUnsafe(_609_v38,(_562_v5).IsProperSubsetOf(_562_v5));
          _610_v39 = (_610_v39).update(_609_v38, _555_v0);
          let _611_v40;
          let _nw137 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
          _611_v40 = _nw137;
          let _612_v41;
          _612_v41 = _module.D13.create_DC29(false, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _613_v42;
          _613_v42 = _dafny.Seq.of((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _pat_let_tv19 = _555_v0;
          let _614_v43;
          let _nw138 = Array((new BigNumber(18)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _612_v41;
          _nw138[(_dafny.ONE).toNumber()] = _612_v41;
          _nw138[(new BigNumber(2)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(3)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(4)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(5)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(6)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(7)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(8)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(9)).toNumber()] = _module.D13.create_DC29(_555_v0, (_613_v42)[_module.__default.safeIndex(p1, new BigNumber((_613_v42).length))]);
          _nw138[(new BigNumber(10)).toNumber()] = _module.D13.create_DC29(!(_555_v0), true);
          _nw138[(new BigNumber(11)).toNumber()] = _module.D13.create_DC29((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], _555_v0);
          _nw138[(new BigNumber(12)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(13)).toNumber()] = function (_pat_let3_0) {
            return function (_615_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_616_dt__update_hcf46_h0) {
                  return _module.D13.create_DC29((_615_dt__update__tmp_h0).dtor_cf45, _616_dt__update_hcf46_h0);
                }(_pat_let4_0);
              }(_pat_let_tv19);
            }(_pat_let3_0);
          }(_612_v41);
          _nw138[(new BigNumber(14)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(15)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(16)).toNumber()] = _612_v41;
          _nw138[(new BigNumber(17)).toNumber()] = _612_v41;
          _614_v43 = _nw138;
          let _617_v44;
          _617_v44 = _dafny.Set.fromElements(_614_v43);
          let _index87 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_611_v40).length));
          (_611_v40)[_index87] = (_617_v44).Union(_617_v44);
          let _618_v45;
          _618_v45 = _dafny.Seq.UnicodeFromString("xo");
          let _index88 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          let _index89 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_611_v40).length));
          let _rhs110 = (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))];
          let _rhs111 = (((_555_v0) ? (_617_v44) : (_dafny.Set.fromElements(_614_v43, _614_v43)))).Union(_617_v44);
          let _rhs112 = _618_v45;
          let _rhs113 = (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))];
          let _rhs114 = (_610_v39).Merge(_610_v39);
          let _lhs52 = _558_v3;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          let _lhs54 = _611_v40;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_611_v40).length));
          _lhs52[_lhs53] = _rhs110;
          _lhs54[_lhs55] = _rhs111;
          _618_v45 = _rhs112;
          _555_v0 = _rhs113;
          _610_v39 = _rhs114;
          let _619_v46;
          _619_v46 = !(_module.__default.fm1(p0, (_dafny.ZERO).minus(p1), false, _555_v0, globalState));
          r3 = (_619_v46);
          r0 = (_dafny.ZERO).minus(p0);
        }
        _555_v0 = !(false) || (_555_v0);
        let _620_v47;
        let _nw139 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _620_v47 = _nw139;
        let _621_v48;
        _621_v48 = _module.D11.create_DC25(false, p0, _620_v47, p1);
        if ((_621_v48).dtor_cf38) {
          let _622_v49;
          _622_v49 = new _dafny.CodePoint('y'.codePointAt(0));
          _622_v49 = _622_v49;
          let _623_v50;
          _623_v50 = _dafny.Seq.UnicodeFromString("iyoke");
          _555_v0 = (_this).fm3(new BigNumber((_623_v50).length), (_this).fm4(_556_v1, globalState), globalState);
          let _624_v51;
          _624_v51 = _dafny.Map.Empty.slice().updateUnsafe(p1,_558_v3);
          let _625_v52;
          _625_v52 = _dafny.Set.fromElements(_558_v3, (((_624_v51).contains(new BigNumber(710))) ? ((_624_v51).get(new BigNumber(710))) : (_558_v3)), _558_v3);
          let _index90 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index90] = (_625_v52).IsProperSubsetOf(_625_v52);
          let _626_v53;
          _626_v53 = _module.D2.create_DC3(_558_v3);
          _626_v53 = _626_v53;
          _623_v50 = _dafny.Seq.update(_623_v50, _module.__default.safeIndex(p1, new BigNumber((_623_v50).length)), _622_v49);
        } else {
          let _627_v54;
          _627_v54 = _module.D8.create_DC19(p0, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _628_v55;
          _628_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(605),_555_v0);
          let _pat_let_tv20 = _620_v47;
          let _629_v56;
          let _nw140 = Array((new BigNumber(24)).toNumber());
          _nw140[(_dafny.ZERO).toNumber()] = _621_v48;
          _nw140[(_dafny.ONE).toNumber()] = _module.D11.create_DC25(_555_v0, p1, _620_v47, p0);
          _nw140[(new BigNumber(2)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(3)).toNumber()] = _module.D11.create_DC25(_555_v0, p1, _620_v47, new BigNumber(826));
          _nw140[(new BigNumber(4)).toNumber()] = _module.D11.create_DC25(false, p0, _620_v47, p1);
          _nw140[(new BigNumber(5)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(6)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(7)).toNumber()] = function (_pat_let5_0) {
            return function (_630_dt__update__tmp_h1) {
              return function (_pat_let6_0) {
                return function (_631_dt__update_hcf40_h0) {
                  return _module.D11.create_DC25((_630_dt__update__tmp_h1).dtor_cf38, (_630_dt__update__tmp_h1).dtor_cf39, _631_dt__update_hcf40_h0, (_630_dt__update__tmp_h1).dtor_cf41);
                }(_pat_let6_0);
              }(_pat_let_tv20);
            }(_pat_let5_0);
          }(_module.D11.create_DC25((_module.D8.create_DC19(new BigNumber((_556_v1).length), (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))])).dtor_cf30, p1, _620_v47, p0));
          _nw140[(new BigNumber(8)).toNumber()] = (((_this).fm3(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_627_v54).dtor_cf29,_555_v0)).length), globalState)) ? (_621_v48) : (_621_v48));
          _nw140[(new BigNumber(9)).toNumber()] = _module.D11.create_DC25(!((((_628_v55).contains(p1)) ? ((_628_v55).get(p1)) : (_555_v0))), new BigNumber(181), _620_v47, p0);
          _nw140[(new BigNumber(10)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(11)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(12)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(13)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(14)).toNumber()] = _module.D11.create_DC25((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))], p1, _620_v47, new BigNumber((_dafny.Seq.of(p0)).length));
          _nw140[(new BigNumber(15)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(16)).toNumber()] = _module.D11.create_DC25(_555_v0, p0, _620_v47, p0);
          _nw140[(new BigNumber(17)).toNumber()] = _module.D11.create_DC25((_this).fm5(new BigNumber((_dafny.Set.fromElements((_this).f5)).length), !(true), p0, p0, globalState), p1, _620_v47, (_this).fm4(_556_v1, globalState));
          _nw140[(new BigNumber(18)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(19)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(20)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(21)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(22)).toNumber()] = _621_v48;
          _nw140[(new BigNumber(23)).toNumber()] = _621_v48;
          _629_v56 = _nw140;
          let _index91 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_629_v56).length));
          (_629_v56)[_index91] = _621_v48;
          let _632_v57;
          _632_v57 = new _dafny.CodePoint('x'.codePointAt(0));
          let _633_v58;
          _633_v58 = _dafny.Set.fromElements(p0, p1, (_dafny.ZERO).minus(p0));
          let _634_v59;
          _634_v59 = _module.D2.create_DC6(p0);
          let _635_v60;
          _635_v60 = _dafny.Set.fromElements(_634_v59, _634_v59);
          let _index92 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_629_v56).length));
          let _rhs115 = true;
          let _rhs116 = _621_v48;
          let _rhs117 = (_633_v58).IsDisjointFrom((_633_v58).Difference(_633_v58));
          let _rhs118 = new BigNumber((function () {
            let _coll25 = new _dafny.Set();
            for (const _compr_25 of (_635_v60).Elements) {
              let _636_v61 = _compr_25;
              if ((_635_v60).contains(_636_v61)) {
                _coll25.add(_636_v61);
              }
            }
            return _coll25;
          }()).length);
          let _rhs119 = ((_555_v0) ? ((_this).f5) : (new _dafny.CodePoint('p'.codePointAt(0))));
          let _lhs56 = _629_v56;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_629_v56).length));
          r3 = _rhs115;
          _lhs56[_lhs57] = _rhs116;
          r3 = _rhs117;
          r1 = _rhs118;
          _632_v57 = _rhs119;
          _632_v57 = _632_v57;
          let _637_v62;
          _637_v62 = _dafny.Set.fromElements(_620_v47, _620_v47, _620_v47);
          let _638_v63;
          let _nw141 = new _module.C0();
          _nw141.__ctor();
          _638_v63 = _nw141;
          let _639_v64;
          let _nw142 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _639_v64 = _nw142;
          let _index93 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_639_v64).length));
          (_639_v64)[_index93] = (_this).f5;
          let _pat_let_tv21 = _555_v0;
          let _index94 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_639_v64).length));
          let _rhs120 = _637_v62;
          let _rhs121 = (((function (_pat_let7_0) {
            return function (_640_dt__update__tmp_h2) {
              return function (_pat_let8_0) {
                return function (_641_dt__update_hcf30_h0) {
                  return _module.D8.create_DC19((_640_dt__update__tmp_h2).dtor_cf29, _641_dt__update_hcf30_h0);
                }(_pat_let8_0);
              }(_pat_let_tv21);
            }(_pat_let7_0);
          }(_627_v54)).dtor_cf30) ? (_638_v63) : (_638_v63));
          let _rhs122 = (_this).f5;
          let _rhs123 = (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))];
          let _lhs58 = _639_v64;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_639_v64).length));
          _637_v62 = _rhs120;
          _638_v63 = _rhs121;
          _lhs58[_lhs59] = _rhs122;
          _555_v0 = _rhs123;
          _632_v57 = _632_v57;
          let _642_v65;
          let _init20 = function (_643_i5) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("apibiv"), _dafny.Seq.UnicodeFromString("lkanpg"));
          };
          let _nw143 = Array((new BigNumber(26)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw143.length); _i0_20++) {
            _nw143[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _642_v65 = _nw143;
          let _644_v66;
          _644_v66 = _dafny.Seq.UnicodeFromString("xnotyjd");
          let _index95 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_642_v65).length));
          (_642_v65)[_index95] = _644_v66;
          let _index96 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          let _index97 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_642_v65).length));
          let _rhs124 = ((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]) === (_555_v0);
          let _rhs125 = _644_v66;
          let _rhs126 = (p0).plus(p0);
          let _lhs60 = _558_v3;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          let _lhs62 = _642_v65;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_642_v65).length));
          _lhs60[_lhs61] = _rhs124;
          _lhs62[_lhs63] = _rhs125;
          r1 = _rhs126;
        }
        let _index98 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
        (_558_v3)[_index98] = _555_v0;
        if (_555_v0) {
          let _645_v67;
          let _nw144 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _645_v67 = _nw144;
          let _646_v68;
          _646_v68 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _647_v69;
          _647_v69 = _module.D8.create_DC19(p0, _555_v0);
          let _index99 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_645_v67).length));
          (_645_v67)[_index99] = _dafny.Set.fromElements((((_646_v68).contains(p1)) ? ((_646_v68).get(p1)) : ((_647_v69).dtor_cf29)));
          let _648_v70;
          _648_v70 = _dafny.Set.fromElements(p1, new BigNumber(669), (_dafny.ZERO).minus(p0));
          let _index100 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_645_v67).length));
          (_645_v67)[_index100] = (_648_v70).Union(_648_v70);
          let _649_v71;
          let _nw145 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _649_v71 = _nw145;
          _649_v71 = _649_v71;
          let _index101 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_620_v47).length));
          (_620_v47)[_index101] = p1;
          let _index102 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_620_v47).length));
          (_620_v47)[_index102] = p1;
          let _650_v72;
          _650_v72 = _dafny.Seq.of(true, (_this).fm3(p0, (_620_v47)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_620_v47).length))], globalState));
          let _index103 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index103] = _dafny.Seq.IsPrefixOf(_module.__default.fm24(new BigNumber((_650_v72).length), globalState), _module.__default.fm24(new BigNumber(857), globalState));
          _module.__default.m0(p1, _dafny.Seq.contains(_module.__default.fm24(p0, globalState), (_this).f5), globalState);
        } else {
          let _651_v73;
          let _nw146 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _651_v73 = _nw146;
          let _652_v74;
          _652_v74 = _dafny.Seq.UnicodeFromString("qpwliykl");
          let _index104 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_651_v73).length));
          (_651_v73)[_index104] = _dafny.Seq.update(_652_v74, _module.__default.safeIndex(p1, new BigNumber((_652_v74).length)), (_this).f5);
          let _653_v75;
          _653_v75 = _module.D2.create_DC4((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
          let _654_v76;
          _654_v76 = _dafny.Map.Empty.slice().updateUnsafe(_653_v75,_dafny.Seq.Create(_module.__default.abs(new BigNumber(410)), function (_655_i6) {
            return (_this).f5;
          }));
          let _656_v77;
          _656_v77 = _dafny.Map.Empty.slice().updateUnsafe((((_654_v76).contains(_module.D2.create_DC4(_555_v0))) ? ((_654_v76).get(_module.D2.create_DC4(_555_v0))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_657_i7) {
            return (_this).f5;
          }))),_652_v74);
          let _index105 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_651_v73).length));
          (_651_v73)[_index105] = _dafny.Seq.Concat((((_656_v77).contains(_dafny.Seq.UnicodeFromString("dr"))) ? ((_656_v77).get(_dafny.Seq.UnicodeFromString("dr"))) : (_652_v74)), _dafny.Seq.Concat(_652_v74, _652_v74));
          let _658_v78;
          let _nw147 = Array((new BigNumber(19)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = (_this).f5;
          _nw147[(_dafny.ONE).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(2)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(3)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(4)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(5)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
          _nw147[(new BigNumber(7)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(8)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(9)).toNumber()] = _module.__default.fm0(p0, p0, globalState);
          _nw147[(new BigNumber(10)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(11)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
          _nw147[(new BigNumber(13)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(14)).toNumber()] = _module.__default.fm0(p1, (_dafny.ZERO).minus(p1), globalState);
          _nw147[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw147[(new BigNumber(16)).toNumber()] = _module.__default.fm0(p0, (_dafny.ZERO).minus(p1), globalState);
          _nw147[(new BigNumber(17)).toNumber()] = ((_651_v73)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_651_v73).length))])[_module.__default.safeIndex(new BigNumber((_556_v1).length), new BigNumber(((_651_v73)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_651_v73).length))]).length))];
          _nw147[(new BigNumber(18)).toNumber()] = (_this).f5;
          _658_v78 = _nw147;
          let _index106 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_658_v78).length));
          (_658_v78)[_index106] = (_652_v74)[_module.__default.safeIndex(p1, new BigNumber((_652_v74).length))];
          let _index107 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_658_v78).length));
          (_658_v78)[_index107] = (_this).f5;
          let _index108 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
          (_558_v3)[_index108] = _555_v0;
          let _index109 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_620_v47).length));
          (_620_v47)[_index109] = new BigNumber((_dafny.Seq.Concat(_652_v74, _652_v74)).length);
          let _index110 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_620_v47).length));
          (_620_v47)[_index110] = (p1).multipliedBy(new BigNumber(((_651_v73)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_651_v73).length))]).length));
          r3 = (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))];
        }
      }
      let _hi6 = (_this).fm4(_556_v1, globalState);
      for (let _659_i8 = p0; _659_i8.isLessThan(_hi6); _659_i8 = _659_i8.plus(_dafny.ONE)) {
        let _660_v79;
        _660_v79 = _dafny.Seq.UnicodeFromString("j");
        let _661_v80;
        _661_v80 = _module.D1.create_DC2(new BigNumber((_660_v79).length), p0);
        let _662_v81;
        _662_v81 = _dafny.Seq.of((_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))]);
        let _663_v82;
        _663_v82 = _module.D7.create_DC16((_this).f5);
        let _664_v83;
        let _nw148 = Array((new BigNumber(24)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = (_this).f5;
        _nw148[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
        _nw148[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
        _nw148[(new BigNumber(5)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(6)).toNumber()] = (_663_v82).dtor_cf23;
        _nw148[(new BigNumber(7)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(8)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(9)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(10)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(11)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(12)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(13)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw148[(new BigNumber(15)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(16)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(17)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(18)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
        _nw148[(new BigNumber(20)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(21)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(22)).toNumber()] = (_this).f5;
        _nw148[(new BigNumber(23)).toNumber()] = (_this).f5;
        _664_v83 = _nw148;
        let _665_v84;
        _665_v84 = _dafny.Seq.of((_module.D10.create_DC22((_module.D5.create_DC13(_660_v79, _659_i8, (_558_v3)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length))])).dtor_cf16, _664_v83, new BigNumber(239))).dtor_cf33, _660_v79, _dafny.Seq.of((_this).f5));
        let _666_v85;
        _666_v85 = _dafny.MultiSet.fromElements(new BigNumber((_562_v5).length), p1, new BigNumber((_556_v1).length), p0, _659_i8);
        let _667_v86;
        let _nw149 = Array((new BigNumber(10)).toNumber());
        _nw149[(_dafny.ZERO).toNumber()] = _659_i8;
        _nw149[(_dafny.ONE).toNumber()] = p1;
        _nw149[(new BigNumber(2)).toNumber()] = p1;
        _nw149[(new BigNumber(3)).toNumber()] = (_661_v80).dtor_cf2;
        _nw149[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(415), (_556_v1)[_module.__default.safeIndex(new BigNumber((_662_v81).length), new BigNumber((_556_v1).length))]);
        _nw149[(new BigNumber(5)).toNumber()] = new BigNumber((_665_v84).length);
        _nw149[(new BigNumber(6)).toNumber()] = (new BigNumber((_666_v85).cardinality())).plus(new BigNumber((_556_v1).length));
        _nw149[(new BigNumber(7)).toNumber()] = p1;
        _nw149[(new BigNumber(8)).toNumber()] = (_this).fm4(_556_v1, globalState);
        _nw149[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_660_v79, (_665_v84)[_module.__default.safeIndex((_dafny.ZERO).minus(_659_i8), new BigNumber((_665_v84).length))])).length));
        _667_v86 = _nw149;
        _667_v86 = _667_v86;
        r1 = _659_i8;
        let _index111 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
        (_558_v3)[_index111] = (new BigNumber(361)).isEqualTo(p1);
        let _index112 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
        (_558_v3)[_index112] = (_this).fm3(p1, p1, globalState);
      }
      let _668_v87;
      _668_v87 = _dafny.Seq.UnicodeFromString("cwenu");
      let _669_v88;
      _669_v88 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_668_v87).length)));
      let _670_v89;
      _670_v89 = _dafny.Map.Empty.slice().updateUnsafe(_669_v88,p0);
      r3 = !(_dafny.Map.Empty.slice().updateUnsafe(_669_v88,p1)).equals(_670_v89);
      let _671_v90;
      _671_v90 = _module.D1.create_DC1(_557_v2);
      _671_v90 = _module.D1.create_DC1(_557_v2);
      let _672_v91;
      _672_v91 = _module.D8.create_DC19(p1, _555_v0);
      let _pat_let_tv22 = _555_v0;
      let _index113 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_558_v3).length));
      (_558_v3)[_index113] = function (_source10) {
        if (_source10.is_DC19) {
          let _673___mcc_h3 = (_source10).cf29;
          let _674___mcc_h4 = (_source10).cf30;
          let _675_cf30 = _674___mcc_h4;
          let _676_cf29 = _673___mcc_h3;
          return _675_cf30;
        } else {
          let _677___mcc_h5 = (_source10).cf28;
          let _678_cf28 = _677___mcc_h5;
          return _pat_let_tv22;
        }
      }(_672_v91);
      let _679_v92;
      _679_v92 = _dafny.MultiSet.fromElements(_556_v1);
      let _680_v93;
      _680_v93 = _dafny.Seq.of(_679_v92, _679_v92);
      r0 = new BigNumber(((_680_v93)[_module.__default.safeIndex(p1, new BigNumber((_680_v93).length))]).cardinality());
      r1 = p1;
      r2 = (_dafny.ZERO).minus((p1).minus((_dafny.ZERO).minus(p1)));
      r3 = _555_v0;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber(916), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements((_this).f5, (_this).f5, (_this).f5)).cardinality()))).length))).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), (_this).f5), _dafny.Seq.of((_this).f5))).length));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(true, true, true, false, false)).Intersect(_dafny.MultiSet.fromElements(true))).IsDisjointFrom((_dafny.MultiSet.fromElements(false, true, false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))));
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(803)).isLessThan(_module.__default.safeModuloInt(new BigNumber(320), new BigNumber(224)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let _681_v0;
      _681_v0 = false;
      let _682_v1;
      let _nw150 = Array((new BigNumber(11)).toNumber()).fill(false);
      _682_v1 = _nw150;
      let _index114 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
      (_682_v1)[_index114] = _681_v0;
      let _683_v2;
      _683_v2 = new BigNumber(559);
      let _684_v3;
      _684_v3 = _dafny.Seq.UnicodeFromString("hrpoqg");
      let _685_v4;
      _685_v4 = _dafny.MultiSet.fromElements(_683_v2, ((!(_module.__default.fm1(_683_v2, _683_v2, _681_v0, _681_v0, globalState))) ? (_683_v2) : ((_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), ((_686_v3) => function (_687_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_686_v3).length));
      })(_684_v3)), globalState))), (_683_v2).minus(new BigNumber((_684_v3).length)));
      let _688_v5;
      _688_v5 = _dafny.Set.fromElements(_681_v0);
      let _index115 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
      let _rhs127 = (_module.__default.safeDivisionInt(_683_v2, new BigNumber((_684_v3).length))).isLessThan((new BigNumber((_module.__default.fm15(_681_v0, _681_v0, _683_v2, globalState)).cardinality())).minus(new BigNumber(-447)));
      let _rhs128 = (((_685_v4).contains(new BigNumber((_688_v5).length))) ? ((_685_v4).get(new BigNumber((_688_v5).length))) : ((_683_v2).multipliedBy(_683_v2)));
      let _rhs129 = _683_v2;
      let _rhs130 = !_dafny.Seq.contains(_684_v3, (_this).f5);
      let _lhs64 = _682_v1;
      let _lhs65 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
      _681_v0 = _rhs127;
      r1 = _rhs128;
      r1 = _rhs129;
      _lhs64[_lhs65] = _rhs130;
      let _689_v6;
      _689_v6 = _dafny.MultiSet.fromElements(_681_v0);
      if ((_681_v0) && (!((_dafny.ZERO).minus(new BigNumber((_689_v6).cardinality()))).isEqualTo(_683_v2))) {
        _683_v2 = (_dafny.ZERO).minus(_683_v2);
        let _690_v7;
        let _init21 = function (_691_i1) {
          return _module.__default.safeDivisionInt(_691_i1, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(909)), function (_692_i2) {
            return (_this).f5;
          }))).cardinality()));
        };
        let _nw151 = Array((new BigNumber(24)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw151.length); _i0_21++) {
          _nw151[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _690_v7 = _nw151;
        let _index116 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length));
        (_690_v7)[_index116] = _683_v2;
        let _693_v8;
        _693_v8 = _dafny.Set.fromElements(_682_v1, _682_v1);
        let _694_v9;
        _694_v9 = _693_v8;
        let _695_v10;
        _695_v10 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_683_v2),(_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
        let _696_v11;
        let _nw152 = Array((new BigNumber(7)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = _695_v10;
        _nw152[(_dafny.ONE).toNumber()] = _695_v10;
        _nw152[(new BigNumber(2)).toNumber()] = _695_v10;
        _nw152[(new BigNumber(3)).toNumber()] = _695_v10;
        _nw152[(new BigNumber(4)).toNumber()] = _695_v10;
        _nw152[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_684_v3).length),_681_v0);
        _nw152[(new BigNumber(6)).toNumber()] = _695_v10;
        _696_v11 = _nw152;
        let _697_v12;
        _697_v12 = _module.D4.create_DC10(_694_v9, _683_v2, _684_v3, _696_v11);
        let _index117 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length));
        (_690_v7)[_index117] = (_dafny.ZERO).minus((_697_v12).dtor_cf12);
        let _698_v13;
        _698_v13 = _dafny.Set.fromElements(_683_v2, new BigNumber(-787), new BigNumber(-779), new BigNumber(99));
        _683_v2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_698_v13).length), _683_v2));
        let _index118 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
        (_682_v1)[_index118] = _681_v0;
        if ((_this).fm3(_683_v2, (_690_v7)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length))], globalState)) {
          (globalState).f2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), ((_699_v2) => function (_700_i3) {
            return _699_v2;
          })(_683_v2));
          let _701_v14;
          _701_v14 = _dafny.Seq.of(false, (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          let _index119 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _index120 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _rhs131 = !_dafny.areEqual(_701_v14, _dafny.Seq.of((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]));
          let _rhs132 = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
          let _rhs133 = _690_v7;
          let _rhs134 = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
          let _rhs135 = ((_681_v0) ? ((_dafny.ZERO).minus(_683_v2)) : (_683_v2));
          let _lhs66 = _682_v1;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _lhs68 = _682_v1;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          _lhs66[_lhs67] = _rhs131;
          _681_v0 = _rhs132;
          _690_v7 = _rhs133;
          _lhs68[_lhs69] = _rhs134;
          r1 = _rhs135;
          let _702_v15;
          _702_v15 = _dafny.Set.fromElements(_684_v3, _module.__default.fm16((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _683_v2, globalState));
          _702_v15 = _dafny.Set.fromElements(_dafny.Seq.update(_684_v3, _module.__default.safeIndex((_690_v7)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length))], new BigNumber((_684_v3).length)), (_this).f5));
          let _703_v16;
          let _nw153 = Array((new BigNumber(10)).toNumber()).fill([]);
          _703_v16 = _nw153;
          let _index121 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_703_v16).length));
          (_703_v16)[_index121] = _682_v1;
          let _index122 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_703_v16).length));
          (_703_v16)[_index122] = _682_v1;
          let _704_v17;
          let _nw154 = new _module.C0();
          _nw154.__ctor();
          _704_v17 = _nw154;
        } else {
          let _index123 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length));
          (_690_v7)[_index123] = ((_dafny.ZERO).minus((_690_v7)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length))])).multipliedBy(_module.__default.safeModuloInt(_683_v2, _683_v2));
          let _705_v18;
          _705_v18 = _module.D5.create_DC13(_684_v3, _683_v2, (_681_v0) || ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]));
          let _706_v19;
          _706_v19 = _dafny.Seq.of(_682_v1, _682_v1);
          let _pat_let_tv23 = _684_v3;
          let _index124 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _rhs136 = false;
          let _rhs137 = function (_pat_let9_0) {
            return function (_709_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_710_dt__update_hcf16_h1) {
                  return _module.D5.create_DC13(_710_dt__update_hcf16_h1, (_709_dt__update__tmp_h1).dtor_cf17, (_709_dt__update__tmp_h1).dtor_cf18);
                }(_pat_let12_0);
              }(_pat_let_tv23);
            }(_pat_let9_0);
          }(function (_pat_let10_0) {
            return function (_707_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_708_dt__update_hcf16_h0) {
                  return _module.D5.create_DC13(_708_dt__update_hcf16_h0, (_707_dt__update__tmp_h0).dtor_cf17, (_707_dt__update__tmp_h0).dtor_cf18);
                }(_pat_let11_0);
              }(_dafny.Seq.UnicodeFromString("xyfu"));
            }(_pat_let10_0);
          }(_705_v18));
          let _rhs138 = _dafny.Seq.of(_682_v1, _682_v1);
          let _rhs139 = (_module.D6.create_DC15((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _681_v0, false)).dtor_cf20;
          let _lhs70 = _682_v1;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          _681_v0 = _rhs136;
          _705_v18 = _rhs137;
          _706_v19 = _rhs138;
          _lhs70[_lhs71] = _rhs139;
          let _711_v20;
          let _nw155 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _711_v20 = _nw155;
          let _index125 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_711_v20).length));
          (_711_v20)[_index125] = (_this).f5;
          let _712_v21;
          _712_v21 = _dafny.Seq.of(_681_v0, _681_v0, (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          let _713_v22;
          _713_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_712_v21, _712_v21),_681_v0);
          let _714_v23;
          _714_v23 = _dafny.Seq.of(new BigNumber(-282));
          let _index126 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_711_v20).length));
          let _index127 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _rhs140 = (_this).f5;
          let _rhs141 = (new BigNumber((_714_v23).length)).isEqualTo(_683_v2);
          let _rhs142 = !_dafny.areEqual(_684_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(444)), function (_715_i4) {
            return (_this).f5;
          }));
          let _rhs143 = ((_713_v22).Merge((_module.__default.fm19(_681_v0, new BigNumber((_685_v4).cardinality()), (((_713_v22).contains(_681_v0)) ? ((_713_v22).get(_681_v0)) : (_681_v0)), globalState)))).update((new BigNumber((function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(81), new BigNumber(72))) {
              let _716_v24 = _compr_26;
              if (((new BigNumber(81)).isLessThanOrEqualTo(_716_v24)) && ((_716_v24).isLessThan(new BigNumber(72)))) {
                _coll26.push([_module.__default.safeModuloInt(_716_v24, _683_v2),true]);
              }
            }
            return _coll26;
          }()).length)).isLessThanOrEqualTo(_683_v2), ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) || ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]));
          let _rhs144 = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
          let _lhs72 = _711_v20;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_711_v20).length));
          let _lhs74 = _682_v1;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          _lhs72[_lhs73] = _rhs140;
          _681_v0 = _rhs141;
          _lhs74[_lhs75] = _rhs142;
          _713_v22 = _rhs143;
          _681_v0 = _rhs144;
          let _index128 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_690_v7).length));
          (_690_v7)[_index128] = new BigNumber((_dafny.Seq.Concat(_684_v3, _684_v3)).length);
          let _index129 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index129] = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
        }
      } else {
        let _717_v25;
        let _nw156 = new _module.C0();
        _nw156.__ctor();
        _717_v25 = _nw156;
        _681_v0 = _681_v0;
        _684_v3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(661)), function (_718_i5) {
          return (_this).f5;
        });
        _683_v2 = _683_v2;
        let _719_v26;
        let _init22 = ((_720_v2) => function (_721_i6) {
          return _module.__default.safeDivisionInt(_721_i6, _720_v2);
        })(_683_v2);
        let _nw157 = Array((new BigNumber(3)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw157.length); _i0_22++) {
          _nw157[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _719_v26 = _nw157;
        let _722_v27;
        _722_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC18(_685_v4),_719_v26);
        let _723_v28;
        _723_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_module.__default.fm16(_681_v0, _683_v2, globalState), _module.__default.safeIndex(new BigNumber((_722_v27).length), new BigNumber((_module.__default.fm16(_681_v0, _683_v2, globalState)).length)), (_this).f5),_681_v0);
        let _724_v29;
        _724_v29 = _dafny.Map.Empty.slice().updateUnsafe((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))],(_this).fm3(new BigNumber(625), _683_v2, globalState));
        let _725_v30;
        _725_v30 = _dafny.Map.Empty.slice().updateUnsafe(_723_v28,!(!(!((((_724_v29).contains((_this).fm3(_683_v2, _dafny.ONE, globalState))) ? ((_724_v29).get((_this).fm3(_683_v2, _dafny.ONE, globalState))) : (_681_v0))) || (_681_v0))));
        _725_v30 = (_725_v30).update((_723_v28).Merge(_723_v28), (_683_v2).isEqualTo(new BigNumber((_689_v6).cardinality())));
      }
      let _index130 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
      (_682_v1)[_index130] = !(true);
      let _726_v31;
      _726_v31 = _dafny.Seq.of((((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) : ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))])), (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], false, _module.__default.fm1(_683_v2, _683_v2, (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _681_v0, globalState));
      let _index131 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
      (_682_v1)[_index131] = !((_726_v31)[_module.__default.safeIndex(_683_v2, new BigNumber((_726_v31).length))]);
      let _727_v32;
      _727_v32 = _module.D2.create_DC3(_682_v1);
      let _pat_let_tv24 = _682_v1;
      let _source11 = function (_pat_let13_0) {
        return function (_728_dt__update__tmp_h2) {
          return function (_pat_let14_0) {
            return function (_729_dt__update_hcf4_h0) {
              return _module.D2.create_DC3(_729_dt__update_hcf4_h0);
            }(_pat_let14_0);
          }(_pat_let_tv24);
        }(_pat_let13_0);
      }(_727_v32);
      if (_source11.is_DC4) {
        let _730___mcc_h0 = (_source11).cf5;
        let _731_cf5 = _730___mcc_h0;
        let _732_v33;
        _732_v33 = _dafny.Set.fromElements(_683_v2, _683_v2, _683_v2);
        let _733_v34;
        _733_v34 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_683_v2, _683_v2));
        let _734_v35;
        _734_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.MultiSet.fromElements(_732_v33)).Difference(_733_v34)).cardinality()),new BigNumber(139));
        _734_v35 = (_734_v35).update(_683_v2, (_dafny.ZERO).minus(_683_v2));
        let _735_v36;
        _735_v36 = _dafny.Map.Empty.slice().updateUnsafe(_684_v3,_681_v0);
        let _736_v37;
        _736_v37 = _dafny.Seq.of(_688_v5);
        let _index132 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
        let _rhs145 = !(!(((_736_v37)[_module.__default.safeIndex(new BigNumber((_732_v33).length), new BigNumber((_736_v37).length))]).IsSubsetOf(_dafny.Set.fromElements(_731_cf5))));
        let _rhs146 = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
        let _rhs147 = (_dafny.Map.Empty.slice().updateUnsafe(_684_v3,_731_cf5)).Merge(_735_v36);
        let _rhs148 = _684_v3;
        let _lhs76 = _682_v1;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
        _731_cf5 = _rhs145;
        _lhs76[_lhs77] = _rhs146;
        _735_v36 = _rhs147;
        _684_v3 = _rhs148;
        let _737_v38;
        let _nw158 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _737_v38 = _nw158;
        let _index133 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length));
        (_737_v38)[_index133] = (_683_v2).plus(new BigNumber(983));
        let _738_v39;
        _738_v39 = _dafny.Map.Empty.slice().updateUnsafe((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))],new BigNumber(-502));
        let _index134 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length));
        (_737_v38)[_index134] = (new BigNumber(((_738_v39).Merge(_738_v39)).length)).plus(new BigNumber((_684_v3).length));
        let _739_v40;
        _739_v40 = _dafny.Seq.of(_689_v6, _689_v6, _689_v6, _689_v6);
        let _740_v41;
        _740_v41 = _dafny.Seq.of((_737_v38)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length))], (_737_v38)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length))], new BigNumber(361));
        let _741_v42;
        _741_v42 = _dafny.Seq.of(_740_v41, _740_v41, _dafny.Seq.of(new BigNumber(547)));
        let _index135 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length));
        let _rhs149 = ((_737_v38)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length))]).isLessThanOrEqualTo(new BigNumber(277));
        let _rhs150 = _683_v2;
        let _rhs151 = ((_737_v38)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length))]).plus(new BigNumber((_dafny.Seq.Concat(_739_v40, _module.__default.fm20(_683_v2, (_741_v42)[_module.__default.safeIndex(_683_v2, new BigNumber((_741_v42).length))], (_this).f5, globalState))).length));
        let _lhs78 = _737_v38;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_737_v38).length));
        _681_v0 = _rhs149;
        r1 = _rhs150;
        _lhs78[_lhs79] = _rhs151;
      } else if (_source11.is_DC5) {
        let _742___mcc_h1 = (_source11).cf6;
        let _743_cf6 = _742___mcc_h1;
        let _744_v43;
        _744_v43 = _dafny.Seq.of(_683_v2);
        r1 = new BigNumber(((((_this).fm5(new BigNumber((_744_v43).length), (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _683_v2, new BigNumber(-1), globalState)) ? (_dafny.Seq.UnicodeFromString("ccxgxm")) : (_dafny.Seq.update(_743_cf6, _module.__default.safeIndex((_this).fm4(_744_v43, globalState), new BigNumber((_743_cf6).length)), (_this).f5)))).length);
        let _745_v44;
        _745_v44 = _dafny.Map.Empty.slice().updateUnsafe(_726_v31,_744_v43);
        let _index136 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
        (_682_v1)[_index136] = _module.__default.fm1(_683_v2, _683_v2, true, !_dafny.Seq.contains((((_745_v44).contains(_module.__default.fm21(_683_v2, (_744_v43)[_module.__default.safeIndex(_683_v2, new BigNumber((_744_v43).length))], _688_v5, globalState))) ? ((_745_v44).get(_module.__default.fm21(_683_v2, (_744_v43)[_module.__default.safeIndex(_683_v2, new BigNumber((_744_v43).length))], _688_v5, globalState))) : (_dafny.Seq.of(_683_v2))), _683_v2), globalState);
        _684_v3 = _743_cf6;
        if (false) {
          let _746_v45;
          let _nw159 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _746_v45 = _nw159;
          let _747_v46;
          _747_v46 = _dafny.Map.Empty.slice().updateUnsafe(_746_v45,_681_v0);
          _683_v2 = new BigNumber(((_747_v46).Merge(_747_v46)).length);
          let _748_v47;
          let _init23 = ((_749_v2) => function (_750_i7) {
            return (_750_i7).minus(_749_v2);
          })(_683_v2);
          let _nw160 = Array((new BigNumber(14)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw160.length); _i0_23++) {
            _nw160[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _748_v47 = _nw160;
          let _index137 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length));
          (_748_v47)[_index137] = (_683_v2).multipliedBy(_683_v2);
          let _index138 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length));
          (_748_v47)[_index138] = _module.__default.safeModuloInt(_683_v2, (_this).fm4(_744_v43, globalState));
          let _index139 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index139] = (_685_v4).IsDisjointFrom(_dafny.MultiSet.fromElements(_683_v2));
          let _index140 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length));
          (_748_v47)[_index140] = (_683_v2).multipliedBy(((!(_681_v0)) ? (_683_v2) : (new BigNumber(216))));
          let _751_v48;
          _751_v48 = _dafny.Map.Empty.slice().updateUnsafe((_748_v47)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length))],(_748_v47)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length))]);
          _751_v48 = (_751_v48).update((new BigNumber((_684_v3).length)).plus((_748_v47)[_module.__default.safeIndex(new BigNumber(917), new BigNumber((_748_v47).length))]), new BigNumber((_744_v43).length));
        } else {
          let _752_v49;
          let _init24 = ((_753_v2) => function (_754_i8) {
            return (_754_i8).multipliedBy(_753_v2);
          })(_683_v2);
          let _nw161 = Array((new BigNumber(22)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw161.length); _i0_24++) {
            _nw161[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _752_v49 = _nw161;
          let _755_v50;
          _755_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("f")).length),_752_v49);
          _752_v49 = ((_681_v0) ? ((((_755_v50).contains(new BigNumber(355))) ? ((_755_v50).get(new BigNumber(355))) : (_752_v49))) : (_752_v49));
          (globalState).f2 = _744_v43;
          let _index141 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_752_v49).length));
          (_752_v49)[_index141] = _683_v2;
          let _756_v51;
          _756_v51 = _dafny.Seq.of(_752_v49, _752_v49);
          let _index142 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_752_v49).length));
          let _index143 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _index144 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _rhs152 = ((_681_v0) ? (_683_v2) : (new BigNumber(763)));
          let _rhs153 = _683_v2;
          let _rhs154 = _683_v2;
          let _rhs155 = true;
          let _rhs156 = _dafny.areEqual(_756_v51, _756_v51);
          let _lhs80 = _752_v49;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_752_v49).length));
          let _lhs82 = _682_v1;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _lhs84 = _682_v1;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          _lhs80[_lhs81] = _rhs152;
          r1 = _rhs153;
          r1 = _rhs154;
          _lhs82[_lhs83] = _rhs155;
          _lhs84[_lhs85] = _rhs156;
          let _757_v52;
          _757_v52 = _dafny.Map.Empty.slice().updateUnsafe((_752_v49)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_752_v49).length))],new _dafny.CodePoint('h'.codePointAt(0)));
          _757_v52 = (_757_v52).update(new BigNumber((_dafny.Seq.Concat(_726_v31, _726_v31)).length), (_this).f5);
          let _758_v53;
          let _nw162 = new _module.C0();
          _nw162.__ctor();
          _758_v53 = _nw162;
          let _759_v54;
          _759_v54 = _dafny.Map.Empty.slice().updateUnsafe(_681_v0,_758_v53);
          let _760_v55;
          _760_v55 = _681_v0;
          _759_v54 = (_759_v54).update((_760_v55), _758_v53);
        }
      } else if (_source11.is_DC6) {
        let _761___mcc_h2 = (_source11).cf7;
        let _762_cf7 = _761___mcc_h2;
        let _763_v56;
        _763_v56 = _dafny.Seq.of(_683_v2);
        let _764_v57;
        _764_v57 = _module.D6.create_DC14((((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? (_763_v56) : (_763_v56)));
        _764_v57 = _764_v57;
        let _765_v58;
        let _nw163 = new _module.C0();
        _nw163.__ctor();
        _765_v58 = _nw163;
        if (_681_v0) {
          let _index145 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index145] = _681_v0;
          let _index146 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index146] = _681_v0;
          let _index147 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index147] = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
          r1 = _683_v2;
          let _766_v59;
          _766_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? ((_this).f5) : ((_this).f5)));
          let _767_v60;
          _767_v60 = _module.D7.create_DC16((_this).f5);
          _766_v59 = (_766_v59).update(((false) ? ((_767_v60).dtor_cf23) : (new _dafny.CodePoint('b'.codePointAt(0)))), _module.__default.fm0(_762_cf7, _683_v2, globalState));
        } else {
          let _768_v61;
          _768_v61 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,(_dafny.ZERO).minus(_762_cf7));
          let _769_v62;
          _769_v62 = _dafny.Map.Empty.slice().updateUnsafe(_689_v6,_768_v61);
          let _770_v63;
          let _init25 = ((_771_v2) => function (_772_i9) {
            return (_772_i9).minus((_dafny.ZERO).minus(_771_v2));
          })(_683_v2);
          let _nw164 = Array((new BigNumber(14)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw164.length); _i0_25++) {
            _nw164[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _770_v63 = _nw164;
          let _773_v64;
          let _nw165 = Array((_dafny.ONE).toNumber()).fill([]);
          _773_v64 = _nw165;
          let _774_v65;
          let _init26 = ((_775_v3) => function (_776_i10) {
            return _775_v3;
          })(_684_v3);
          let _nw166 = Array((new BigNumber(10)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw166.length); _i0_26++) {
            _nw166[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _774_v65 = _nw166;
          let _index148 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_773_v64).length));
          (_773_v64)[_index148] = _774_v65;
          let _index149 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_773_v64).length));
          let _rhs157 = _762_cf7;
          let _rhs158 = _769_v62;
          let _rhs159 = (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))];
          let _rhs160 = _770_v63;
          let _rhs161 = _774_v65;
          let _lhs86 = _773_v64;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_773_v64).length));
          r1 = _rhs157;
          _769_v62 = _rhs158;
          _681_v0 = _rhs159;
          _770_v63 = _rhs160;
          _lhs86[_lhs87] = _rhs161;
          let _index150 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index150] = (new BigNumber(((_689_v6).Union(_689_v6)).cardinality())).isLessThan((_683_v2).multipliedBy(_683_v2));
          _683_v2 = _683_v2;
          let _777_v66;
          _777_v66 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_684_v3, _dafny.Seq.UnicodeFromString("fusnvwurp")),(_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          _777_v66 = _777_v66;
          let _index151 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_770_v63).length));
          (_770_v63)[_index151] = new BigNumber((_684_v3).length);
          let _index152 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_770_v63).length));
          (_770_v63)[_index152] = (new BigNumber((_726_v31).length)).plus(_683_v2);
        }
        let _778_v67;
        _778_v67 = _module.D8.create_DC18(_685_v4);
        let _779_v68;
        _779_v68 = _dafny.Seq.of(_685_v4, _685_v4);
        let _780_v69;
        _780_v69 = _dafny.Map.Empty.slice().updateUnsafe(_689_v6,(_dafny.ZERO).minus(_762_cf7));
        let _781_v70;
        let _nw167 = Array((new BigNumber(10)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = (_685_v4).Difference((_778_v67).dtor_cf28);
        _nw167[(_dafny.ONE).toNumber()] = _685_v4;
        _nw167[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_685_v4).cardinality()), (_dafny.ZERO).minus((_763_v56)[_module.__default.safeIndex(_683_v2, new BigNumber((_763_v56).length))]), _762_cf7);
        _nw167[(new BigNumber(3)).toNumber()] = _685_v4;
        _nw167[(new BigNumber(4)).toNumber()] = (_685_v4).update(new BigNumber(404), _module.__default.abs(_683_v2));
        _nw167[(new BigNumber(5)).toNumber()] = _685_v4;
        _nw167[(new BigNumber(6)).toNumber()] = _685_v4;
        _nw167[(new BigNumber(7)).toNumber()] = (_779_v68)[_module.__default.safeIndex(_683_v2, new BigNumber((_779_v68).length))];
        _nw167[(new BigNumber(8)).toNumber()] = _685_v4;
        _nw167[(new BigNumber(9)).toNumber()] = (_685_v4).update(_762_cf7, _module.__default.abs((((_780_v69).contains(_dafny.MultiSet.fromElements((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _681_v0))) ? ((_780_v69).get(_dafny.MultiSet.fromElements((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _681_v0))) : (_762_cf7))));
        _781_v70 = _nw167;
        let _index153 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_781_v70).length));
        (_781_v70)[_index153] = _685_v4;
        let _index154 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_781_v70).length));
        (_781_v70)[_index154] = _dafny.MultiSet.fromElements((((_689_v6).contains(_681_v0)) ? ((_689_v6).get(_681_v0)) : (_683_v2)), (((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? (new BigNumber((_689_v6).cardinality())) : (_683_v2)));
      } else if (_source11.is_DC3) {
        let _782___mcc_h3 = (_source11).cf4;
        let _783_cf4 = _782___mcc_h3;
        let _784_v71;
        _784_v71 = _dafny.Map.Empty.slice().updateUnsafe(((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) && ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]),_681_v0);
        _784_v71 = (_784_v71).update(_module.__default.fm1(_683_v2, _683_v2, (_this).fm3(_683_v2, _683_v2, globalState), _681_v0, globalState), _681_v0);
        let _785_v72;
        _785_v72 = _dafny.Seq.of(_684_v3, _dafny.Seq.UnicodeFromString("hwrcaq"));
        let _786_v73;
        _786_v73 = _module.D5.create_DC13((_785_v72)[_module.__default.safeIndex(new BigNumber((_688_v5).length), new BigNumber((_785_v72).length))], _683_v2, (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
        let _787_v74;
        _787_v74 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,(_786_v73).dtor_cf17);
        _787_v74 = (_787_v74).update(_module.__default.safeDivisionInt(_683_v2, new BigNumber((_dafny.Seq.of(new BigNumber((_688_v5).length))).length)), _module.__default.safeModuloInt(new BigNumber(802), _683_v2));
        if (_681_v0) {
          let _788_v75;
          _788_v75 = _module.D1.create_DC1(_689_v6);
          _788_v75 = ((!((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))])) ? (_module.D1.create_DC1(_689_v6)) : (_788_v75));
          let _789_v76;
          let _init27 = ((_790_v2) => function (_791_i11) {
            return (_791_i11).minus(new BigNumber((_dafny.MultiSet.fromElements(_790_v2)).cardinality()));
          })(_683_v2);
          let _nw168 = Array((new BigNumber(20)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw168.length); _i0_27++) {
            _nw168[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _789_v76 = _nw168;
          let _index155 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length));
          (_789_v76)[_index155] = _683_v2;
          let _index156 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length));
          (_789_v76)[_index156] = _module.__default.safeDivisionInt(_683_v2, _module.__default.safeModuloInt(_683_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_792_i12) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)));
          _681_v0 = _681_v0;
          _787_v74 = (_787_v74).update((_789_v76)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length))], (_789_v76)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length))]);
          let _index157 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length));
          (_789_v76)[_index157] = ((!((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))])) ? ((_789_v76)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_789_v76).length))]) : (new BigNumber((_dafny.Seq.Concat(_726_v31, _726_v31)).length)));
        } else {
          let _793_v77;
          _793_v77 = _module.D10.create_DC21(_dafny.Set.fromElements(_681_v0, _681_v0));
          let _index158 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index158] = (_module.__default.safeModuloInt(_683_v2, new BigNumber(((_793_v77).dtor_cf32).length))).isEqualTo((new BigNumber(667)).plus(new BigNumber((_787_v74).length)));
          let _794_v79;
          let _init28 = ((_795_v2) => function (_796_i13) {
            return function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), ((_797_v2) => function (_798_i14) {
                return _797_v2;
              })(_795_v2))).Elements) {
                let _799_v78 = _compr_27;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), ((_800_v2) => function (_798_i14) {
                  return _800_v2;
                })(_795_v2)), _799_v78)) {
                  _coll27.add((_799_v78).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ykjtoq")).length))).length),true)).length),new BigNumber(792))).length)));
                }
              }
              return _coll27;
            }();
          })(_683_v2);
          let _nw169 = Array((new BigNumber(23)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw169.length); _i0_28++) {
            _nw169[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _794_v79 = _nw169;
          let _801_v80;
          _801_v80 = _module.D11.create_DC24(_794_v79);
          let _802_v81;
          _802_v81 = _dafny.Seq.of(_794_v79, _794_v79, _794_v79, _794_v79, _794_v79);
          let _803_v82;
          let _nw170 = Array((new BigNumber(25)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _794_v79;
          _nw170[(_dafny.ONE).toNumber()] = _794_v79;
          _nw170[(new BigNumber(2)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(3)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(4)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(5)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(6)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(7)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(8)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(9)).toNumber()] = (_801_v80).dtor_cf37;
          _nw170[(new BigNumber(10)).toNumber()] = (((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? (_794_v79) : (_794_v79));
          _nw170[(new BigNumber(11)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(12)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(13)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(14)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(15)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(16)).toNumber()] = (_802_v81)[_module.__default.safeIndex((_dafny.ZERO).minus(_683_v2), new BigNumber((_802_v81).length))];
          _nw170[(new BigNumber(17)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(18)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(19)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(20)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(21)).toNumber()] = (_801_v80).dtor_cf37;
          _nw170[(new BigNumber(22)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(23)).toNumber()] = _794_v79;
          _nw170[(new BigNumber(24)).toNumber()] = _794_v79;
          _803_v82 = _nw170;
          let _index159 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_803_v82).length));
          (_803_v82)[_index159] = _794_v79;
          let _804_v83;
          _804_v83 = _dafny.MultiSet.fromElements(_684_v3, _dafny.Seq.Concat(_684_v3, _684_v3), _dafny.Seq.UnicodeFromString("lmaps"), _dafny.Seq.UnicodeFromString("s"));
          let _index160 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_803_v82).length));
          let _rhs162 = _794_v79;
          let _rhs163 = (_688_v5).IsProperSubsetOf(_dafny.Set.fromElements((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]));
          let _rhs164 = _804_v83;
          let _rhs165 = _dafny.Seq.Concat(_684_v3, _dafny.Seq.Concat(_dafny.Seq.update(_684_v3, _module.__default.safeIndex(new BigNumber(276), new BigNumber((_684_v3).length)), (_this).f5), _684_v3));
          let _lhs88 = _803_v82;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_803_v82).length));
          _lhs88[_lhs89] = _rhs162;
          _681_v0 = _rhs163;
          _804_v83 = _rhs164;
          _684_v3 = _rhs165;
          let _805_v84;
          _805_v84 = new _dafny.CodePoint('p'.codePointAt(0));
          let _806_v85;
          let _nw171 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _806_v85 = _nw171;
          let _index161 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_806_v85).length));
          (_806_v85)[_index161] = (_this).f5;
          let _index162 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_806_v85).length));
          let _rhs166 = (_this).f5;
          let _rhs167 = _module.__default.fm0(_683_v2, _683_v2, globalState);
          let _lhs90 = _806_v85;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_806_v85).length));
          _805_v84 = _rhs166;
          _lhs90[_lhs91] = _rhs167;
          let _807_v86;
          let _nw172 = new _module.C0();
          _nw172.__ctor();
          _807_v86 = _nw172;
          let _808_v87;
          _808_v87 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_dafny.Seq.of(_807_v86, _807_v86));
          let _809_v88;
          _809_v88 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_808_v87);
          _809_v88 = (_809_v88).update(_683_v2, _808_v87);
          _683_v2 = new BigNumber(740);
        }
        let _810_v89;
        let _init29 = function (_811_i15) {
          return _module.__default.safeModuloInt(_811_i15, new BigNumber(445));
        };
        let _nw173 = Array((new BigNumber(16)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw173.length); _i0_29++) {
          _nw173[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _810_v89 = _nw173;
        let _index163 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_810_v89).length));
        (_810_v89)[_index163] = new BigNumber(930);
        let _812_v90;
        _812_v90 = _module.D11.create_DC25((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _683_v2, _810_v89, _683_v2);
        let _813_v91;
        let _nw174 = Array((new BigNumber(5)).toNumber());
        _nw174[(_dafny.ZERO).toNumber()] = _812_v90;
        _nw174[(_dafny.ONE).toNumber()] = _module.D11.create_DC25((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))], _683_v2, _810_v89, _683_v2);
        _nw174[(new BigNumber(2)).toNumber()] = _812_v90;
        _nw174[(new BigNumber(3)).toNumber()] = _812_v90;
        _nw174[(new BigNumber(4)).toNumber()] = _812_v90;
        _813_v91 = _nw174;
        let _814_v92;
        _814_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-171),_813_v91);
        let _815_v93;
        _815_v93 = _dafny.MultiSet.fromElements((((_814_v92).contains(new BigNumber(708))) ? ((_814_v92).get(new BigNumber(708))) : (_813_v91)), _813_v91);
        let _index164 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_810_v89).length));
        (_810_v89)[_index164] = new BigNumber((((_815_v93).Union(_815_v93)).Union(_dafny.MultiSet.fromElements(_813_v91))).cardinality());
      } else {
        let _816___mcc_h4 = (_source11).cf8;
        let _817_cf8 = _816___mcc_h4;
        if ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) {
          let _818_v94;
          let _nw175 = new _module.C0();
          _nw175.__ctor();
          _818_v94 = _nw175;
          let _nw176 = new _module.C0();
          _nw176.__ctor();
          _818_v94 = _nw176;
          let _819_v95;
          let _nw177 = new _module.C1();
          _nw177.__ctor((_this).f4, (_this).f5);
          _819_v95 = _nw177;
          let _820_v96;
          let _init30 = ((_821_v2) => function (_822_i16) {
            return _module.__default.safeDivisionInt(_822_i16, _821_v2);
          })(_683_v2);
          let _nw178 = Array((new BigNumber(12)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw178.length); _i0_30++) {
            _nw178[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _820_v96 = _nw178;
          let _index165 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v96).length));
          (_820_v96)[_index165] = new BigNumber(612);
          let _823_v97;
          _823_v97 = _dafny.Seq.of(_683_v2);
          let _index166 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v96).length));
          let _rhs168 = (_dafny.ZERO).minus((((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) ? (_683_v2) : (new BigNumber((_dafny.Seq.Concat(_823_v97, _dafny.Seq.of(_683_v2))).length))));
          let _rhs169 = ((_681_v0) ? (_819_v95) : (_819_v95));
          let _rhs170 = (new BigNumber(329)).plus(_683_v2);
          let _rhs171 = _683_v2;
          let _lhs92 = _820_v96;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v96).length));
          _683_v2 = _rhs168;
          _819_v95 = _rhs169;
          r1 = _rhs170;
          _lhs92[_lhs93] = _rhs171;
          let _824_v98;
          let _nw179 = Array((new BigNumber(16)).toNumber());
          _824_v98 = _nw179;
          let _825_v99;
          _825_v99 = _824_v98;
          let _826_v100;
          _826_v100 = _dafny.Set.fromElements(_824_v98, _825_v99, _825_v99);
          let _index167 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          (_682_v1)[_index167] = !(((_826_v100).Difference(_826_v100)).IsProperSubsetOf(_dafny.Set.fromElements(_825_v99)));
          r1 = _683_v2;
          let _827_v101;
          _827_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-331)), function (_828_i17) {
            return (_this).f5;
          }), _684_v3),(_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          let _index168 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          let _rhs172 = ((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]) && (_681_v0);
          let _rhs173 = _827_v101;
          let _rhs174 = (_689_v6).Difference((_689_v6).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]))));
          let _lhs94 = _682_v1;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length));
          _lhs94[_lhs95] = _rhs172;
          _827_v101 = _rhs173;
          _689_v6 = _rhs174;
        } else {
          let _829_v102;
          let _nw180 = new _module.C1();
          _nw180.__ctor(((_681_v0) ? ((_this).f4) : ((_this).f4)), (_684_v3)[_module.__default.safeIndex(_683_v2, new BigNumber((_684_v3).length))]);
          _829_v102 = _nw180;
          let _830_v103;
          _830_v103 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_683_v2);
          let _831_v104;
          _831_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_688_v5).length),_830_v103);
          let _832_v107;
          _832_v107 = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_831_v104);
          let _833_v109;
          _833_v109 = _module.D2.create_DC4((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          let _834_v110;
          _834_v110 = _dafny.Seq.of(_683_v2);
          let _835_v113;
          _835_v113 = _module.D14.create_DC30(function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(-267), new BigNumber(748))) {
    let _836_v111 = _compr_28;
    if (((new BigNumber(-267)).isLessThanOrEqualTo(_836_v111)) && ((_836_v111).isLessThan(new BigNumber(748)))) {
      _coll28.push([_module.__default.safeModuloInt(_836_v111, _683_v2),function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(570), new BigNumber(540))) {
          let _837_v112 = _compr_29;
          if (((new BigNumber(570)).isLessThanOrEqualTo(_837_v112)) && ((_837_v112).isLessThan(new BigNumber(540)))) {
            _coll29.push([(_837_v112).minus(new BigNumber(294)),_683_v2]);
          }
        }
        return _coll29;
      }()]);
    }
  }
  return _coll28;
}());
          let _838_v114;
          let _nw181 = Array((new BigNumber(29)).toNumber());
          _nw181[(_dafny.ZERO).toNumber()] = (_831_v104).Merge(_831_v104);
          _nw181[(_dafny.ONE).toNumber()] = (_831_v104).Merge(_831_v104);
          _nw181[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_830_v103);
          _nw181[(new BigNumber(3)).toNumber()] = function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(16), new BigNumber(730))) {
              let _839_v105 = _compr_30;
              if (((new BigNumber(16)).isLessThanOrEqualTo(_839_v105)) && ((_839_v105).isLessThan(new BigNumber(730)))) {
                _coll30.push([(_839_v105).minus(_683_v2),_830_v103]);
              }
            }
            return _coll30;
          }();
          _nw181[(new BigNumber(4)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(5)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(6)).toNumber()] = (_831_v104).Merge(_831_v104);
          _nw181[(new BigNumber(7)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(8)).toNumber()] = (_831_v104).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_684_v3).length),_830_v103));
          _nw181[(new BigNumber(9)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(10)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_683_v2,_dafny.Map.Empty.slice().updateUnsafe(_683_v2,_683_v2));
          _nw181[(new BigNumber(12)).toNumber()] = function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(504), new BigNumber(891))) {
              let _840_v106 = _compr_31;
              if (((new BigNumber(504)).isLessThanOrEqualTo(_840_v106)) && ((_840_v106).isLessThan(new BigNumber(891)))) {
                _coll31.push([_module.__default.safeDivisionInt(_840_v106, _683_v2),_830_v103]);
              }
            }
            return _coll31;
          }();
          _nw181[(new BigNumber(13)).toNumber()] = (((_832_v107).contains(_683_v2)) ? ((_832_v107).get(_683_v2)) : (function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(26), new BigNumber(851))) {
              let _841_v108 = _compr_32;
              if (((new BigNumber(26)).isLessThanOrEqualTo(_841_v108)) && ((_841_v108).isLessThan(new BigNumber(851)))) {
                _coll32.push([(_841_v108).multipliedBy(_683_v2),_830_v103]);
              }
            }
            return _coll32;
          }()));
          _nw181[(new BigNumber(14)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(15)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(16)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(17)).toNumber()] = (_831_v104).Merge(_831_v104);
          _nw181[(new BigNumber(18)).toNumber()] = ((_831_v104).update(_683_v2, _830_v103)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_683_v2,_830_v103));
          _nw181[(new BigNumber(19)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(20)).toNumber()] = (_831_v104).Merge(_831_v104);
          _nw181[(new BigNumber(21)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(22)).toNumber()] = _module.__default.fm25(_681_v0, (_833_v109).dtor_cf5, _834_v110, globalState);
          _nw181[(new BigNumber(23)).toNumber()] = (_831_v104).update(new BigNumber(-582), _830_v103);
          _nw181[(new BigNumber(24)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(25)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(26)).toNumber()] = (_835_v113).dtor_cf47;
          _nw181[(new BigNumber(27)).toNumber()] = _831_v104;
          _nw181[(new BigNumber(28)).toNumber()] = (_831_v104).Merge(_831_v104);
          _838_v114 = _nw181;
          let _index169 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_838_v114).length));
          (_838_v114)[_index169] = (_831_v104).Merge(_831_v104);
          let _index170 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_838_v114).length));
          (_838_v114)[_index170] = _831_v104;
          _684_v3 = _684_v3;
          let _842_v115;
          _842_v115 = _module.D15.create_DC34(_726_v31);
          let _843_v116;
          let _nw182 = Array((new BigNumber(13)).toNumber());
          _nw182[(_dafny.ZERO).toNumber()] = _module.__default.fm21(new BigNumber((_684_v3).length), _683_v2, _688_v5, globalState);
          _nw182[(_dafny.ONE).toNumber()] = _dafny.Seq.update((_842_v115).dtor_cf53, _module.__default.safeIndex(_683_v2, new BigNumber(((_842_v115).dtor_cf53).length)), _681_v0);
          _nw182[(new BigNumber(2)).toNumber()] = _726_v31;
          _nw182[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]), _module.__default.safeIndex(_683_v2, new BigNumber((_dafny.Seq.of((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))])).length)), _681_v0), _726_v31);
          _nw182[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_726_v31, _726_v31);
          _nw182[(new BigNumber(5)).toNumber()] = _726_v31;
          _nw182[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_726_v31, _726_v31);
          _nw182[(new BigNumber(7)).toNumber()] = _726_v31;
          _nw182[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_726_v31, _726_v31);
          _nw182[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_681_v0), _module.__default.fm21(_683_v2, _683_v2, _dafny.Set.fromElements(_681_v0), globalState));
          _nw182[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]);
          _nw182[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_726_v31, _dafny.Seq.update(_726_v31, _module.__default.safeIndex(_683_v2, new BigNumber((_726_v31).length)), (_682_v1)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_682_v1).length))]));
          _nw182[(new BigNumber(12)).toNumber()] = _726_v31;
          _843_v116 = _nw182;
          let _index171 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_843_v116).length));
          (_843_v116)[_index171] = _module.__default.fm21(_683_v2, _683_v2, _688_v5, globalState);
          let _index172 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_843_v116).length));
          (_843_v116)[_index172] = _726_v31;
          r1 = _683_v2;
        }
        let _844_v117;
        let _init31 = function (_845_i18) {
          return (_845_i18).plus(new BigNumber(887));
        };
        let _nw183 = Array((new BigNumber(13)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw183.length); _i0_31++) {
          _nw183[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _844_v117 = _nw183;
        let _index173 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_844_v117).length));
        (_844_v117)[_index173] = _683_v2;
        let _index174 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_844_v117).length));
        (_844_v117)[_index174] = _683_v2;
        _684_v3 = _684_v3;
        r1 = new BigNumber(-575);
      }
      let _846_v118;
      _846_v118 = _dafny.Seq.of(_dafny.Seq.Concat(_726_v31, _726_v31));
      _846_v118 = _846_v118;
      r0 = _dafny.Seq.of(new BigNumber(((_685_v4).Intersect(_685_v4)).cardinality()), _module.__default.safeModuloInt(_683_v2, new BigNumber((_dafny.Seq.of(_683_v2)).length)), new BigNumber(270), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_681_v0,_683_v2)).length));
      r1 = _683_v2;
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _847_v0;
      _847_v0 = true;
      let _rhs175 = false;
      let _rhs176 = _847_v0;
      let _rhs177 = _847_v0;
      _847_v0 = _rhs175;
      _847_v0 = _rhs176;
      _847_v0 = _rhs177;
      let _848_v1;
      _848_v1 = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,_847_v0);
      let _849_v2;
      _849_v2 = new BigNumber(-346);
      let _850_v3;
      _850_v3 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_848_v1).length), _849_v2, globalState));
      let _851_v4;
      _851_v4 = _dafny.Seq.of((_this).f5);
      if (!(((_850_v3).Difference(_850_v3)).IsDisjointFrom((function () {
        let _coll33 = new _dafny.Set();
        for (const _compr_33 of (_851_v4).Elements) {
          let _852_v5 = _compr_33;
          if (_dafny.Seq.contains(_851_v4, _852_v5)) {
            _coll33.add(_852_v5);
          }
        }
        return _coll33;
      }()).Intersect(_850_v3)))) {
        let _853_v6;
        let _init32 = ((_854_v0) => function (_855_i0) {
          return !(_854_v0) || (true);
        })(_847_v0);
        let _nw184 = Array((new BigNumber(26)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw184.length); _i0_32++) {
          _nw184[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _853_v6 = _nw184;
        let _index175 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length));
        (_853_v6)[_index175] = _847_v0;
        let _index176 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length));
        (_853_v6)[_index176] = !(_847_v0);
        let _856_v7;
        _856_v7 = _dafny.Set.fromElements(_849_v2, new BigNumber(804));
        let _857_v8;
        _857_v8 = _dafny.Seq.of(_849_v2);
        let _858_v9;
        _858_v9 = _dafny.Map.Empty.slice().updateUnsafe(_856_v7,_857_v8);
        _858_v9 = (_858_v9).Merge(_858_v9);
        _849_v2 = _849_v2;
        let _859_v10;
        let _init33 = ((_860_v2) => function (_861_i1) {
          return (_861_i1).multipliedBy(_860_v2);
        })(_849_v2);
        let _nw185 = Array((new BigNumber(21)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw185.length); _i0_33++) {
          _nw185[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _859_v10 = _nw185;
        let _index177 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_859_v10).length));
        (_859_v10)[_index177] = _849_v2;
        let _862_v11;
        _862_v11 = _dafny.Set.fromElements(!(!(_847_v0)));
        let _863_v12;
        _863_v12 = new _dafny.CodePoint('c'.codePointAt(0));
        let _864_v13;
        _864_v13 = _module.D6.create_DC15((_853_v6)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length))], _847_v0, true);
        let _865_v14;
        _865_v14 = _dafny.Seq.of(_847_v0);
        let _index178 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_859_v10).length));
        let _index179 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length));
        let _rhs178 = (_849_v2).multipliedBy(new BigNumber(792));
        let _rhs179 = (_dafny.Set.fromElements((_864_v13).dtor_cf22, (_853_v6)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length))])).Union(_module.__default.fm26(_849_v2, _849_v2, globalState));
        let _rhs180 = (new BigNumber((_865_v14).length)).isLessThan(_849_v2);
        let _rhs181 = _863_v12;
        let _rhs182 = _849_v2;
        let _lhs96 = _859_v10;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_859_v10).length));
        let _lhs98 = _853_v6;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_853_v6).length));
        _lhs96[_lhs97] = _rhs178;
        _862_v11 = _rhs179;
        _lhs98[_lhs99] = _rhs180;
        _863_v12 = _rhs181;
        _849_v2 = _rhs182;
        let _866_v15;
        _866_v15 = _dafny.MultiSet.fromElements((_859_v10)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_859_v10).length))]);
        let _867_v16;
        _867_v16 = _dafny.Map.Empty.slice().updateUnsafe(_849_v2,(_dafny.MultiSet.fromElements(_849_v2, _849_v2)).Union(_866_v15));
        _867_v16 = (_867_v16).update((_859_v10)[_module.__default.safeIndex(new BigNumber(401), new BigNumber((_859_v10).length))], _866_v15);
      } else {
        if (!(_847_v0)) {
          let _868_v17;
          let _nw186 = Array((new BigNumber(3)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = _848_v1;
          _nw186[(_dafny.ONE).toNumber()] = _848_v1;
          _nw186[(new BigNumber(2)).toNumber()] = _848_v1;
          _868_v17 = _nw186;
          let _869_v18;
          _869_v18 = _dafny.Map.Empty.slice().updateUnsafe(_849_v2,_868_v17);
          let _870_v19;
          let _nw187 = Array((new BigNumber(22)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = _868_v17;
          _nw187[(_dafny.ONE).toNumber()] = _868_v17;
          _nw187[(new BigNumber(2)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(3)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(4)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(5)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(6)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(7)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(8)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(9)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(10)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(11)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(12)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(13)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(14)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(15)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(16)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(17)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(18)).toNumber()] = (((_869_v18).contains(new BigNumber(-100))) ? ((_869_v18).get(new BigNumber(-100))) : (_868_v17));
          _nw187[(new BigNumber(19)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(20)).toNumber()] = _868_v17;
          _nw187[(new BigNumber(21)).toNumber()] = _868_v17;
          _870_v19 = _nw187;
          let _index180 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_870_v19).length));
          (_870_v19)[_index180] = _868_v17;
          let _index181 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_870_v19).length));
          (_870_v19)[_index181] = _868_v17;
          _847_v0 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), function (_871_i2) {
            return (_this).f5;
          }), (_module.D7.create_DC16((_this).f5)).dtor_cf23);
          let _872_v20;
          _872_v20 = _dafny.Seq.of(_849_v2);
          let _873_v21;
          _873_v21 = _dafny.MultiSet.fromElements(_872_v20, _872_v20);
          let _874_v22;
          _874_v22 = _dafny.Set.fromElements(new BigNumber(59));
          let _875_v23;
          let _nw188 = Array((new BigNumber(21)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = ((_847_v0) ? (false) : (_847_v0));
          _nw188[(_dafny.ONE).toNumber()] = _847_v0;
          _nw188[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_849_v2, new BigNumber((_873_v21).cardinality()), _847_v0, _847_v0, globalState);
          _nw188[(new BigNumber(3)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(4)).toNumber()] = !(!(!(!(_847_v0))));
          _nw188[(new BigNumber(5)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(6)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(7)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(8)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(9)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(10)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(11)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(12)).toNumber()] = !(_847_v0);
          _nw188[(new BigNumber(13)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(14)).toNumber()] = !(_847_v0);
          _nw188[(new BigNumber(15)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(16)).toNumber()] = !(_849_v2).isEqualTo(_849_v2);
          _nw188[(new BigNumber(17)).toNumber()] = !(_847_v0);
          _nw188[(new BigNumber(18)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(19)).toNumber()] = _847_v0;
          _nw188[(new BigNumber(20)).toNumber()] = !(_874_v22).equals(_dafny.Set.fromElements(new BigNumber(161)));
          _875_v23 = _nw188;
          let _index182 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_875_v23).length));
          (_875_v23)[_index182] = _847_v0;
          let _876_v24;
          _876_v24 = _dafny.MultiSet.fromElements((_this).f5);
          let _877_v25;
          _877_v25 = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,_849_v2);
          let _index183 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_875_v23).length));
          let _rhs183 = (((_876_v24).contains((_this).f5)) ? ((_876_v24).get((_this).f5)) : ((((_877_v25).contains(!(_847_v0))) ? ((_877_v25).get(!(_847_v0))) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_849_v2))))));
          let _rhs184 = (_876_v24).IsProperSubsetOf((_876_v24).Intersect(_876_v24));
          let _lhs100 = _875_v23;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_875_v23).length));
          _849_v2 = _rhs183;
          _lhs100[_lhs101] = _rhs184;
          _849_v2 = _849_v2;
          let _index184 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_875_v23).length));
          (_875_v23)[_index184] = (_847_v0) || ((_875_v23)[_module.__default.safeIndex(new BigNumber(138), new BigNumber((_875_v23).length))]);
        } else {
          let _878_v26;
          let _nw189 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _878_v26 = _nw189;
          let _879_v27;
          _879_v27 = _dafny.Set.fromElements(_878_v26, _878_v26);
          _879_v27 = _879_v27;
          _847_v0 = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(369)), function (_880_i3) {
            return (_this).f5;
          }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), function (_881_i4) {
            return (_this).f5;
          }), _851_v4));
          let _882_v28;
          let _nw190 = Array((new BigNumber(24)).toNumber()).fill(false);
          _882_v28 = _nw190;
          let _index185 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_882_v28).length));
          (_882_v28)[_index185] = !(!(_847_v0) || (_847_v0));
          let _883_v29;
          _883_v29 = _module.D2.create_DC4(_847_v0);
          let _index186 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_882_v28).length));
          (_882_v28)[_index186] = (_847_v0) && ((_883_v29).dtor_cf5);
          let _884_v30;
          let _nw191 = new _module.C0();
          _nw191.__ctor();
          _884_v30 = _nw191;
          let _885_v31;
          _885_v31 = _module.D7.create_DC16(_module.__default.fm0(_849_v2, _849_v2, globalState));
          let _886_v32;
          _886_v32 = _dafny.Seq.of(_885_v31, _885_v31);
          let _887_v33;
          _887_v33 = _dafny.MultiSet.fromElements(_886_v32);
          _887_v33 = _887_v33;
        }
        let _888_v34;
        let _nw192 = Array((new BigNumber(28)).toNumber()).fill([]);
        _888_v34 = _nw192;
        let _889_v35;
        let _nw193 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _889_v35 = _nw193;
        let _index187 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_888_v34).length));
        (_888_v34)[_index187] = _889_v35;
        let _index188 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_888_v34).length));
        (_888_v34)[_index188] = ((_847_v0) ? (_889_v35) : (_889_v35));
        let _890_v36;
        _890_v36 = _dafny.Map.Empty.slice().updateUnsafe(_849_v2,_847_v0);
        let _891_v37;
        _891_v37 = _dafny.MultiSet.fromElements(_849_v2, _849_v2);
        let _892_v38;
        let _nw194 = Array((new BigNumber(8)).toNumber());
        _nw194[(_dafny.ZERO).toNumber()] = (((_890_v36).contains((_dafny.ZERO).minus(_849_v2))) ? ((_890_v36).get((_dafny.ZERO).minus(_849_v2))) : (_847_v0));
        _nw194[(_dafny.ONE).toNumber()] = ((_847_v0) ? (false) : (!(_847_v0)));
        _nw194[(new BigNumber(2)).toNumber()] = (_849_v2).isLessThan(_849_v2);
        _nw194[(new BigNumber(3)).toNumber()] = (_891_v37).IsSubsetOf(_dafny.MultiSet.fromElements(_849_v2));
        _nw194[(new BigNumber(4)).toNumber()] = _847_v0;
        _nw194[(new BigNumber(5)).toNumber()] = _847_v0;
        _nw194[(new BigNumber(6)).toNumber()] = !(false);
        _nw194[(new BigNumber(7)).toNumber()] = _847_v0;
        _892_v38 = _nw194;
        let _index189 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length));
        (_892_v38)[_index189] = _847_v0;
        let _index190 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length));
        (_892_v38)[_index190] = _847_v0;
        if (_847_v0) {
          _849_v2 = _849_v2;
          let _index191 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length));
          (_892_v38)[_index191] = (_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))];
          let _893_v39;
          _893_v39 = _module.D13.create_DC29((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))], !((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]) || ((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]));
          _893_v39 = _module.__default.fm27(globalState);
          let _894_v40;
          let _nw195 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
          _894_v40 = _nw195;
          let _index192 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_894_v40).length));
          (_894_v40)[_index192] = function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of (_850_v3).Elements) {
              let _895_v41 = _compr_34;
              if ((_850_v3).contains(_895_v41)) {
                _coll34.add(_895_v41);
              }
            }
            return _coll34;
          }();
          let _index193 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_894_v40).length));
          (_894_v40)[_index193] = _dafny.Set.fromElements((_this).f5);
          let _896_v42;
          _896_v42 = _dafny.MultiSet.fromElements((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]);
          let _897_v43;
          _897_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_896_v42).cardinality()),_849_v2);
          _847_v0 = (new BigNumber(-377)).isEqualTo(new BigNumber((_897_v43).length));
        } else {
          let _898_v44;
          let _nw196 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _898_v44 = _nw196;
          let _index194 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_898_v44).length));
          (_898_v44)[_index194] = _dafny.Map.Empty.slice().updateUnsafe(_851_v4,_849_v2);
          let _899_v46;
          _899_v46 = _dafny.MultiSet.fromElements(_851_v4, _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0))), _851_v4);
          let _900_v47;
          _900_v47 = _dafny.Set.fromElements((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]);
          let _901_v48;
          _901_v48 = _dafny.Map.Empty.slice().updateUnsafe(_851_v4,new BigNumber((_900_v47).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_898_v44).length));
          (_898_v44)[_index195] = ((function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_899_v46).Elements) {
              let _902_v45 = _compr_35;
              if ((_899_v46).contains(_902_v45)) {
                _coll35.push([_902_v45,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]))).length)]);
              }
            }
            return _coll35;
          }()).Merge(_901_v48)).Merge((_901_v48).Merge(_901_v48));
          let _903_v49;
          _903_v49 = _dafny.Map.Empty.slice().updateUnsafe(_849_v2,new BigNumber(623));
          let _904_v50;
          _904_v50 = _dafny.MultiSet.fromElements(_847_v0);
          _903_v49 = (_903_v49).update((((_904_v50).contains((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))])) ? ((_904_v50).get((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))])) : (_849_v2)), _849_v2);
          _849_v2 = _849_v2;
          let _905_v51;
          let _nw197 = new _module.C0();
          _nw197.__ctor();
          _905_v51 = _nw197;
          let _906_v52;
          _906_v52 = _dafny.Map.Empty.slice().updateUnsafe(((_847_v0) ? (new BigNumber((_851_v4).length)) : (new BigNumber(-133))),_dafny.MultiSet.fromElements(_849_v2));
          let _rhs185 = _905_v51;
          let _rhs186 = (_849_v2).isLessThan(new BigNumber(113));
          let _rhs187 = _906_v52;
          let _rhs188 = (_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))];
          _905_v51 = _rhs185;
          _847_v0 = _rhs186;
          _906_v52 = _rhs187;
          _847_v0 = _rhs188;
          let _907_v53;
          _907_v53 = _dafny.Seq.of(_892_v38);
          _907_v53 = _907_v53;
        }
        if (_847_v0) {
          let _908_v54;
          _908_v54 = _dafny.Seq.of(_849_v2);
          let _909_v55;
          _909_v55 = _dafny.Map.Empty.slice().updateUnsafe((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))],_908_v54);
          _847_v0 = ((_909_v55).Merge(_909_v55)).equals(_909_v55);
          let _910_v56;
          let _nw198 = new _module.C1();
          _nw198.__ctor((_this).f4, (_this).f5);
          _910_v56 = _nw198;
          let _index196 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length));
          (_892_v38)[_index196] = !((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]);
          let _911_v57;
          let _init34 = ((_912_v4) => function (_913_i5) {
            return (_913_i5).minus(new BigNumber((_912_v4).length));
          })(_851_v4);
          let _nw199 = Array((new BigNumber(21)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw199.length); _i0_34++) {
            _nw199[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _911_v57 = _nw199;
          let _index197 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_911_v57).length));
          (_911_v57)[_index197] = (((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]) ? (_849_v2) : (_849_v2));
          let _index198 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_911_v57).length));
          (_911_v57)[_index198] = _module.__default.safeDivisionInt((_910_v56).fm4(_dafny.Seq.of(_849_v2, _849_v2), globalState), _849_v2);
          let _index199 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_911_v57).length));
          (_911_v57)[_index199] = (_dafny.ZERO).minus(_849_v2);
        } else {
          let _index200 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length));
          (_892_v38)[_index200] = (((_848_v1).contains(_847_v0)) ? ((_848_v1).get(_847_v0)) : ((_849_v2).isEqualTo(_849_v2)));
          let _914_v58;
          _914_v58 = _dafny.Seq.of(_849_v2);
          let _915_v59;
          _915_v59 = _dafny.Seq.of(_914_v58, _914_v58);
          (globalState).f2 = _dafny.Seq.Concat(_914_v58, _dafny.Seq.Concat((_915_v59)[_module.__default.safeIndex(_849_v2, new BigNumber((_915_v59).length))], _914_v58));
          _851_v4 = _851_v4;
          let _916_v60;
          let _nw200 = Array((new BigNumber(22)).toNumber()).fill([]);
          _916_v60 = _nw200;
          let _917_v61;
          let _nw201 = Array((new BigNumber(5)).toNumber());
          _917_v61 = _nw201;
          let _index201 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_916_v60).length));
          (_916_v60)[_index201] = _917_v61;
          let _918_v62;
          _918_v62 = _917_v61;
          let _index202 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_916_v60).length));
          (_916_v60)[_index202] = _918_v62;
          let _919_v63;
          _919_v63 = _dafny.Map.Empty.slice().updateUnsafe(_914_v58,_847_v0);
          _919_v63 = (_919_v63).update(_914_v58, !((_892_v38)[_module.__default.safeIndex(new BigNumber(664), new BigNumber((_892_v38).length))]) || (_847_v0));
        }
      }
      let _920_v64;
      _920_v64 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), function (_921_i6) {
        return (_this).f5;
      })).length), _849_v2);
      _849_v2 = (_849_v2).minus(_module.__default.safeModuloInt(new BigNumber((_920_v64).length), _849_v2));
      let _rhs189 = (((_847_v0) ? (_847_v0) : ((_this).fm3(_849_v2, _849_v2, globalState)))) === (_847_v0);
      let _rhs190 = _849_v2;
      _847_v0 = _rhs189;
      _849_v2 = _rhs190;
      if (!((_module.D14.create_DC32(_849_v2, _847_v0, false)).dtor_cf51)) {
        _849_v2 = _849_v2;
        let _rhs191 = _dafny.Seq.IsPrefixOf(_851_v4, _851_v4);
        let _rhs192 = _849_v2;
        _847_v0 = _rhs191;
        _849_v2 = _rhs192;
        _847_v0 = _847_v0;
        _849_v2 = _849_v2;
        let _922_v65;
        let _nw202 = new _module.C1();
        _nw202.__ctor((_this).f4, new _dafny.CodePoint('a'.codePointAt(0)));
        _922_v65 = _nw202;
      } else {
        let _923_v66;
        _923_v66 = _dafny.MultiSet.fromElements(true);
        let _924_v67;
        _924_v67 = _dafny.Map.Empty.slice().updateUnsafe(_923_v66,_847_v0);
        _924_v67 = (_924_v67).update(_923_v66, (_dafny.MultiSet.fromElements(_847_v0)).IsProperSubsetOf(_923_v66));
        _847_v0 = !(new BigNumber((_851_v4).length)).isEqualTo(_849_v2);
        let _925_v69;
        _925_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update(_920_v64, _module.__default.safeIndex(new BigNumber(992), new BigNumber((_920_v64).length)), _849_v2), _920_v64),(_849_v2).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-714)), function (_926_i7) {
          return (_this).f5;
        }), _module.__default.safeIndex(new BigNumber(((function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_dafny.Map.Empty.slice().updateUnsafe(_849_v2,_849_v2)).Keys.Elements) {
            let _927_v68 = _compr_36;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_849_v2,_849_v2)).contains(_927_v68)) {
              _coll36.push([(_927_v68).minus((_dafny.ZERO).minus(new BigNumber((_851_v4).length))),_847_v0]);
            }
          }
          return _coll36;
        }()).update(_849_v2, _847_v0)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-714)), function (_928_i7) {
          return (_this).f5;
        })).length)), (_this).f5)).length)));
        let _929_v70;
        _929_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_847_v0, _847_v0, _847_v0)).length),_849_v2);
        _925_v69 = (_925_v69).update(((_847_v0) ? (_920_v64) : (_920_v64)), ((_this).fm4(_dafny.Seq.of(new BigNumber((_929_v70).length), _849_v2), globalState)).minus(_849_v2));
        _847_v0 = !(_847_v0) || (true);
        _925_v69 = _925_v69;
      }
      let _930_v71;
      _930_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_847_v0);
      let _931_i8;
      _931_i8 = _dafny.ZERO;
      L3: {
        while ((((_930_v71).contains((_851_v4)[_module.__default.safeIndex(_849_v2, new BigNumber((_851_v4).length))])) ? ((_930_v71).get((_851_v4)[_module.__default.safeIndex(_849_v2, new BigNumber((_851_v4).length))])) : (_847_v0))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_931_i8)) {
              break L3;
            }
            _931_i8 = (_931_i8).plus(_dafny.ONE);
            let _932_v72;
            _932_v72 = _dafny.Map.Empty.slice().updateUnsafe(!(_847_v0),_847_v0);
            let _933_v73;
            let _nw203 = Array((new BigNumber(25)).toNumber());
            _nw203[(_dafny.ZERO).toNumber()] = _848_v1;
            _nw203[(_dafny.ONE).toNumber()] = _848_v1;
            _nw203[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,_847_v0);
            _nw203[(new BigNumber(3)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(4)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(5)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(6)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(7)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(8)).toNumber()] = ((_932_v72)).update(_847_v0, _847_v0);
            _nw203[(new BigNumber(9)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(10)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(11)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_847_v0);
            _nw203[(new BigNumber(13)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,_847_v0);
            _nw203[(new BigNumber(15)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(16)).toNumber()] = (_848_v1).update(_847_v0, _module.__default.fm1(_849_v2, (_920_v64)[_module.__default.safeIndex(_849_v2, new BigNumber((_920_v64).length))], false, _847_v0, globalState));
            _nw203[(new BigNumber(17)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(18)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(19)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(20)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_847_v0);
            _nw203[(new BigNumber(22)).toNumber()] = (_848_v1).update(_847_v0, _847_v0);
            _nw203[(new BigNumber(23)).toNumber()] = _848_v1;
            _nw203[(new BigNumber(24)).toNumber()] = _848_v1;
            _933_v73 = _nw203;
            let _934_v74;
            let _935_v75;
            let _936_v76;
            let _out10;
            let _out11;
            let _out12;
            let _outcollector3 = (_this).m13(_933_v73, globalState);
            _out10 = _outcollector3[0];
            _out11 = _outcollector3[1];
            _out12 = _outcollector3[2];
            _934_v74 = _out10;
            _935_v75 = _out11;
            _936_v76 = _out12;
            let _937_v77;
            _937_v77 = _module.D8.create_DC19(_849_v2, _847_v0);
            let _938_v78;
            let _nw204 = Array((new BigNumber(7)).toNumber());
            _nw204[(_dafny.ZERO).toNumber()] = _module.__default.fm1(new BigNumber(648), _934_v74, _847_v0, _847_v0, globalState);
            _nw204[(_dafny.ONE).toNumber()] = false;
            _nw204[(new BigNumber(2)).toNumber()] = true;
            _nw204[(new BigNumber(3)).toNumber()] = _847_v0;
            _nw204[(new BigNumber(4)).toNumber()] = _847_v0;
            _nw204[(new BigNumber(5)).toNumber()] = ((_937_v77).dtor_cf29).isLessThan(_934_v74);
            _nw204[(new BigNumber(6)).toNumber()] = _847_v0;
            _938_v78 = _nw204;
            let _index203 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length));
            (_938_v78)[_index203] = _847_v0;
            let _index204 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length));
            (_938_v78)[_index204] = false;
            _849_v2 = ((new BigNumber(-428)).plus((_this).fm4(_920_v64, globalState))).multipliedBy(_849_v2);
            let _939_v79;
            _939_v79 = _dafny.MultiSet.fromElements((_938_v78)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length))]);
            let _940_v80;
            _940_v80 = _dafny.Map.Empty.slice().updateUnsafe(_934_v74,new BigNumber((_939_v79).cardinality()));
            let _941_v81;
            _941_v81 = _dafny.Set.fromElements(_940_v80, _940_v80, _940_v80);
            let _942_v82;
            _942_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(new BigNumber((_941_v81).length), _847_v0, _849_v2, new BigNumber(508), globalState),new BigNumber((_920_v64).length));
            let _943_v83;
            _943_v83 = _dafny.Seq.of(_942_v82, _dafny.Map.Empty.slice().updateUnsafe(false,_849_v2), _942_v82);
            let _944_v84;
            _944_v84 = _dafny.Map.Empty.slice().updateUnsafe(_934_v74,_942_v82);
            let _945_v85;
            _945_v85 = _dafny.Map.Empty.slice().updateUnsafe((((_938_v78)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length))]) ? (_943_v83) : (_943_v83)),(((_944_v84).contains(new BigNumber((_939_v79).cardinality()))) ? ((_944_v84).get(new BigNumber((_939_v79).cardinality()))) : (_942_v82)));
            let _946_v86;
            _946_v86 = _dafny.Set.fromElements(false, _847_v0, _module.__default.fm1(_934_v74, (_dafny.ZERO).minus(_849_v2), true, true, globalState), _847_v0, (_938_v78)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length))]);
            _945_v85 = (_945_v85).update(_dafny.Seq.Concat(_943_v83, _dafny.Seq.update(_943_v83, _module.__default.safeIndex(_849_v2, new BigNumber((_943_v83).length)), (_942_v82).update((_938_v78)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_938_v78).length))], _849_v2))), _dafny.Map.Empty.slice().updateUnsafe(_847_v0,new BigNumber((_946_v86).length)));
          }
        }
      }
      let _947_v87;
      let _nw205 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _947_v87 = _nw205;
      r0 = _947_v87;
      return r0;
    }
    m13(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let _948_v0;
      _948_v0 = false;
      let _949_v1;
      let _nw206 = Array((new BigNumber(22)).toNumber());
      _nw206[(_dafny.ZERO).toNumber()] = _948_v0;
      _nw206[(_dafny.ONE).toNumber()] = _948_v0;
      _nw206[(new BigNumber(2)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(3)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(4)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(5)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(6)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(7)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(8)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(9)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(10)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(11)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(12)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(13)).toNumber()] = true;
      _nw206[(new BigNumber(14)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(15)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(16)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(17)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(18)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(19)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(20)).toNumber()] = _948_v0;
      _nw206[(new BigNumber(21)).toNumber()] = _948_v0;
      _949_v1 = _nw206;
      (globalState).f3 = ((_948_v0) ? (_949_v1) : (((!(_948_v0)) ? (_949_v1) : (_949_v1))));
      let _950_v2;
      _950_v2 = _module.D2.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), function (_951_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
}));
      let _952_v3;
      _952_v3 = _dafny.Seq.UnicodeFromString("dt");
      let _953_v4;
      _953_v4 = new BigNumber(897);
      let _954_v5;
      _954_v5 = _module.D5.create_DC13(_952_v3, _953_v4, _948_v0);
      _950_v2 = _module.D2.create_DC5((_954_v5).dtor_cf16);
      if (_948_v0) {
        let _955_v6;
        _955_v6 = _dafny.Seq.of(_949_v1, _949_v1, _949_v1, _949_v1, _949_v1);
        _948_v0 = _dafny.Seq.IsProperPrefixOf(_955_v6, _955_v6);
        let _956_v7;
        let _nw207 = Array((new BigNumber(3)).toNumber()).fill([]);
        _956_v7 = _nw207;
        let _957_v8;
        let _init35 = ((_958_v4) => function (_959_i1) {
          return _module.__default.safeDivisionInt(_959_i1, _958_v4);
        })(_953_v4);
        let _nw208 = Array((new BigNumber(7)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw208.length); _i0_35++) {
          _nw208[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _957_v8 = _nw208;
        let _index205 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length));
        (_956_v7)[_index205] = _957_v8;
        let _index206 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length));
        (_956_v7)[_index206] = _957_v8;
        r0 = _953_v4;
        let _960_v9;
        _960_v9 = _dafny.Seq.of((_956_v7)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length))], (_956_v7)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length))], (_956_v7)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length))]);
        _957_v8 = ((_948_v0) ? ((_956_v7)[_module.__default.safeIndex(new BigNumber(475), new BigNumber((_956_v7).length))]) : ((_960_v9)[_module.__default.safeIndex(_953_v4, new BigNumber((_960_v9).length))]));
        r0 = _953_v4;
      } else {
        let _961_v10;
        _961_v10 = _module.D2.create_DC4(_948_v0);
        let _index207 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
        (_949_v1)[_index207] = !((_961_v10).dtor_cf5);
        let _index208 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
        (_949_v1)[_index208] = _module.__default.fm1(_953_v4, (_dafny.ZERO).minus(_953_v4), false, false, globalState);
        let _962_v11;
        _962_v11 = _module.D2.create_DC6(_953_v4);
        let _source12 = _module.D2.create_DC7(_962_v11);
        if (_source12.is_DC4) {
          let _963___mcc_h0 = (_source12).cf5;
          let _964_cf5 = _963___mcc_h0;
          let _965_v12;
          _965_v12 = _dafny.Map.Empty.slice().updateUnsafe(false,(_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))]);
          _965_v12 = (_965_v12).update(_964_cf5, _948_v0);
          let _966_v13;
          _966_v13 = _dafny.Seq.of(_952_v3, _952_v3);
          _953_v4 = new BigNumber((_dafny.Seq.Concat(_966_v13, _dafny.Seq.Concat(_966_v13, _966_v13))).length);
          let _967_v14;
          _967_v14 = new _dafny.CodePoint('l'.codePointAt(0));
          let _968_v15;
          let _nw209 = Array((new BigNumber(26)).toNumber()).fill(_module.D13.Default());
          _968_v15 = _nw209;
          let _969_v16;
          let _nw210 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _969_v16 = _nw210;
          let _970_v17;
          _970_v17 = _module.D13.create_DC28(_969_v16);
          let _971_v18;
          _971_v18 = _dafny.Map.Empty.slice().updateUnsafe(_953_v4,_970_v17);
          let _972_v19;
          _972_v19 = _module.D13.create_DC29(_948_v0, _module.__default.fm1(new BigNumber((_971_v18).length), (_dafny.ZERO).minus(_953_v4), (_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))], false, globalState));
          let _973_v20;
          _973_v20 = _dafny.Map.Empty.slice().updateUnsafe(_972_v19,false);
          let _974_v22;
          _974_v22 = _dafny.MultiSet.fromElements(_973_v20, _973_v20);
          let _975_v23;
          let _nw211 = Array((new BigNumber(8)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_973_v20, _973_v20, function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_973_v20).Keys.Elements) {
              let _976_v21 = _compr_37;
              if ((_973_v20).contains(_976_v21)) {
                _coll37.push([_976_v21,_964_cf5]);
              }
            }
            return _coll37;
          }());
          _nw211[(_dafny.ONE).toNumber()] = _974_v22;
          _nw211[(new BigNumber(2)).toNumber()] = (_974_v22).Intersect(_974_v22);
          _nw211[(new BigNumber(3)).toNumber()] = _974_v22;
          _nw211[(new BigNumber(4)).toNumber()] = _974_v22;
          _nw211[(new BigNumber(5)).toNumber()] = (_974_v22).Union(_974_v22);
          _nw211[(new BigNumber(6)).toNumber()] = _module.__default.fm28(globalState);
          _nw211[(new BigNumber(7)).toNumber()] = _974_v22;
          _975_v23 = _nw211;
          let _index209 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v23).length));
          (_975_v23)[_index209] = (_module.__default.fm28(globalState)).Difference(_974_v22);
          let _977_v25;
          _977_v25 = _dafny.Set.fromElements(_972_v19);
          let _index210 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v23).length));
          let _rhs193 = _967_v14;
          let _rhs194 = _968_v15;
          let _rhs195 = _dafny.MultiSet.FromArray(_dafny.Seq.of((_973_v20).Merge(function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_977_v25).Elements) {
              let _978_v24 = _compr_38;
              if ((_977_v25).contains(_978_v24)) {
                _coll38.push([_978_v24,_948_v0]);
              }
            }
            return _coll38;
          }())));
          let _lhs102 = _975_v23;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_975_v23).length));
          _967_v14 = _rhs193;
          _968_v15 = _rhs194;
          _lhs102[_lhs103] = _rhs195;
          _964_cf5 = !(false);
        } else if (_source12.is_DC5) {
          let _979___mcc_h1 = (_source12).cf6;
          let _980_cf6 = _979___mcc_h1;
          let _981_v26;
          _981_v26 = _dafny.Seq.of(_948_v0);
          let _982_v28;
          _982_v28 = _dafny.Seq.of(new BigNumber(-267));
          let _index211 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
          (_949_v1)[_index211] = (!((_981_v26)[_module.__default.safeIndex(_953_v4, new BigNumber((_981_v26).length))])) && (!(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_980_cf6).length),new BigNumber(243))).equals(function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_982_v28).Elements) {
              let _983_v27 = _compr_39;
              if (_dafny.Seq.contains(_982_v28, _983_v27)) {
                _coll39.push([_module.__default.safeModuloInt(_983_v27, _953_v4),new BigNumber(75)]);
              }
            }
            return _coll39;
          }()));
          let _984_v29;
          let _nw212 = new _module.C0();
          _nw212.__ctor();
          _984_v29 = _nw212;
          _982_v28 = _982_v28;
          let _985_v30;
          _985_v30 = _dafny.Set.fromElements(_953_v4);
          let _986_v31;
          _986_v31 = _module.D8.create_DC19(_953_v4, (_953_v4).isLessThan(new BigNumber((_985_v30).length)));
          _986_v31 = _986_v31;
        } else if (_source12.is_DC6) {
          let _987___mcc_h2 = (_source12).cf7;
          let _988_cf7 = _987___mcc_h2;
          let _989_v32;
          _989_v32 = _module.D14.create_DC33(_module.D14.create_DC31(_988_cf7));
          let _990_v33;
          _990_v33 = _module.D14.create_DC33(_989_v32);
          let _991_v34;
          _991_v34 = _module.D14.create_DC33(_990_v33);
          let _pat_let_tv25 = _990_v33;
          let _992_v35;
          _992_v35 = _dafny.Seq.of(_991_v34, _991_v34, function (_pat_let15_0) {
            return function (_993_dt__update__tmp_h0) {
              return function (_pat_let16_0) {
                return function (_994_dt__update_hcf52_h0) {
                  return _module.D14.create_DC33(_994_dt__update_hcf52_h0);
                }(_pat_let16_0);
              }(_pat_let_tv25);
            }(_pat_let15_0);
          }(_991_v34), _991_v34, _991_v34);
          let _995_v36;
          _995_v36 = _dafny.Seq.of(_992_v35, _992_v35, _992_v35, _992_v35, _992_v35);
          let _996_v37;
          _996_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_995_v36);
          _996_v37 = (_996_v37).update((_this).f5, _995_v36);
          let _997_v38;
          _997_v38 = _dafny.Map.Empty.slice().updateUnsafe(_948_v0,_988_cf7);
          let _998_v39;
          _998_v39 = _dafny.Seq.of(new BigNumber((_997_v38).length), (_dafny.ZERO).minus(_953_v4));
          let _999_v40;
          _999_v40 = _dafny.Seq.of((_this).fm4(_998_v39, globalState));
          (globalState).f2 = _999_v40;
          (globalState).f2 = _dafny.Seq.Concat(_998_v39, _dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), ((_1000_v4) => function (_1001_i2) {
            return _1000_v4;
          })(_953_v4)));
          let _index212 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
          (_949_v1)[_index212] = !((_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))]);
        } else if (_source12.is_DC3) {
          let _1002___mcc_h3 = (_source12).cf4;
          let _1003_cf4 = _1002___mcc_h3;
          _948_v0 = false;
          _948_v0 = (_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))];
          _948_v0 = (_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))];
          let _1004_v41;
          _1004_v41 = _module.D14.create_DC31(_953_v4);
          r0 = (_1004_v41).dtor_cf48;
        } else {
          let _1005___mcc_h4 = (_source12).cf8;
          let _1006_cf8 = _1005___mcc_h4;
          let _index213 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
          (_949_v1)[_index213] = ((_953_v4).isLessThanOrEqualTo(_953_v4)) === (_948_v0);
          _948_v0 = _module.__default.fm1(_953_v4, _953_v4, _948_v0, _948_v0, globalState);
          _953_v4 = _953_v4;
          r0 = _953_v4;
        }
        let _1007_v42;
        _1007_v42 = _dafny.Map.Empty.slice().updateUnsafe(_948_v0,(_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))]);
        let _index214 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length));
        (_949_v1)[_index214] = (((_1007_v42).contains((_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))])) ? ((_1007_v42).get((_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))])) : ((_this).fm3(_953_v4, _953_v4, globalState)));
        let _1008_v43;
        _1008_v43 = _dafny.MultiSet.fromElements(_953_v4, _953_v4);
        let _1009_v44;
        _1009_v44 = _dafny.Map.Empty.slice().updateUnsafe(((_948_v0) ? (_1008_v43) : (_1008_v43)),new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("faqhwwpjf"), _952_v3)).length));
        _1009_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1008_v43,_953_v4);
        let _1010_v45;
        _1010_v45 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),(_949_v1)[_module.__default.safeIndex(new BigNumber(455), new BigNumber((_949_v1).length))]);
        _1010_v45 = (_1010_v45).update((_this).f5, _948_v0);
      }
      let _1011_v46;
      let _nw213 = Array((new BigNumber(28)).toNumber()).fill(_module.D1.Default());
      _1011_v46 = _nw213;
      let _1012_v47;
      _1012_v47 = _dafny.MultiSet.fromElements(_948_v0, _948_v0);
      let _1013_v48;
      _1013_v48 = _module.D1.create_DC1(_1012_v47);
      let _pat_let_tv26 = _1012_v47;
      let _index215 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_1011_v46).length));
      (_1011_v46)[_index215] = function (_pat_let17_0) {
        return function (_1014_dt__update__tmp_h1) {
          return function (_pat_let18_0) {
            return function (_1015_dt__update_hcf1_h0) {
              return _module.D1.create_DC1(_1015_dt__update_hcf1_h0);
            }(_pat_let18_0);
          }(_pat_let_tv26);
        }(_pat_let17_0);
      }(_1013_v48);
      let _pat_let_tv27 = _1012_v47;
      let _pat_let_tv28 = _948_v0;
      let _pat_let_tv29 = _1012_v47;
      let _index216 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_1011_v46).length));
      (_1011_v46)[_index216] = function (_pat_let19_0) {
        return function (_1018_dt__update__tmp_h3) {
          return function (_pat_let22_0) {
            return function (_1019_dt__update_hcf1_h2) {
              return _module.D1.create_DC1(_1019_dt__update_hcf1_h2);
            }(_pat_let22_0);
          }((_dafny.MultiSet.fromElements(_pat_let_tv28)).Difference(_pat_let_tv29));
        }(_pat_let19_0);
      }(function (_pat_let20_0) {
        return function (_1016_dt__update__tmp_h2) {
          return function (_pat_let21_0) {
            return function (_1017_dt__update_hcf1_h1) {
              return _module.D1.create_DC1(_1017_dt__update_hcf1_h1);
            }(_pat_let21_0);
          }(_pat_let_tv27);
        }(_pat_let20_0);
      }(_1013_v48));
      r0 = (new BigNumber(366)).minus(_953_v4);
      let _1020_v49;
      let _nw214 = new _module.C0();
      _nw214.__ctor();
      _1020_v49 = _nw214;
      let _1021_v50;
      _1021_v50 = _dafny.Set.fromElements((_this).f5, (_this).f5);
      let _1022_v51;
      _1022_v51 = _dafny.MultiSet.fromElements(new BigNumber((_1021_v50).length));
      let _1023_v52;
      _1023_v52 = _dafny.Seq.of(_953_v4);
      let _1024_v53;
      _1024_v53 = _dafny.Map.Empty.slice().updateUnsafe(_948_v0,(_this).fm4(_1023_v52, globalState));
      r0 = new BigNumber((((_module.__default.fm1(new BigNumber((_1022_v51).cardinality()), new BigNumber((_1024_v53).length), _948_v0, false, globalState)) ? (_952_v3) : (_952_v3))).length);
      let _1025_v54;
      let _nw215 = Array((new BigNumber(22)).toNumber()).fill([]);
      _1025_v54 = _nw215;
      r1 = _1025_v54;
      let _1026_v55;
      _1026_v55 = _dafny.Map.Empty.slice().updateUnsafe(_948_v0,_948_v0);
      let _1027_v56;
      _1027_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1026_v55).length),_dafny.Seq.update(_dafny.Seq.of(!(true)), _module.__default.safeIndex(_953_v4, new BigNumber((_dafny.Seq.of(!(true))).length)), _948_v0));
      r2 = (_1027_v56).Merge(_1027_v56);
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_module.D1.create_DC2(new BigNumber(325), new BigNumber((_dafny.Seq.UnicodeFromString("asvmv")).length))).dtor_cf3)).minus(_module.__default.safeModuloInt(new BigNumber(878), new BigNumber(778)));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-780),true))).contains(!(true))) && ((_dafny.MultiSet.fromElements(new BigNumber(999), new BigNumber((_dafny.Seq.UnicodeFromString("bimvtv")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("js")).length))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("tg")).length))).IsProperSubsetOf((_module.D8.create_DC18(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("i")).length)))).dtor_cf28));
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return !(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("jcrfnbjfd"), (_this).f5));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let _1028_v0;
      _1028_v0 = true;
      let _1029_v1;
      let _init36 = ((_1030_v0) => function (_1031_i0) {
        return _1030_v0;
      })(_1028_v0);
      let _nw216 = Array((new BigNumber(3)).toNumber());
      for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw216.length); _i0_36++) {
        _nw216[_i0_36] = _init36(new BigNumber(_i0_36));
      }
      _1029_v1 = _nw216;
      let _rhs196 = !(_1028_v0);
      let _rhs197 = _1029_v1;
      let _lhs104 = globalState;
      _1028_v0 = _rhs196;
      _lhs104.f3 = _rhs197;
      let _index217 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1029_v1).length));
      (_1029_v1)[_index217] = _1028_v0;
      let _1032_v2;
      _1032_v2 = _dafny.Seq.UnicodeFromString("wjluy");
      let _index218 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1029_v1).length));
      (_1029_v1)[_index218] = _dafny.Seq.contains(_1032_v2, (_this).f5);
      let _1033_v3;
      _1033_v3 = new BigNumber(733);
      let _1034_v4;
      _1034_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v3,(_this).f4);
      let _1035_v5;
      let _nw217 = new _module.C2();
      _nw217.__ctor((((_1034_v4).contains(_1033_v3)) ? ((_1034_v4).get(_1033_v3)) : ((_this).f4)), (_this).f5);
      _1035_v5 = _nw217;
      let _1036_v6;
      let _nw218 = Array((new BigNumber(7)).toNumber());
      _nw218[(_dafny.ZERO).toNumber()] = (_this).f5;
      _nw218[(_dafny.ONE).toNumber()] = (_this).f5;
      _nw218[(new BigNumber(2)).toNumber()] = (_this).f5;
      _nw218[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
      _nw218[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
      _nw218[(new BigNumber(5)).toNumber()] = (_this).f5;
      _nw218[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
      _1036_v6 = _nw218;
      let _index219 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1036_v6).length));
      (_1036_v6)[_index219] = (_this).f5;
      let _index220 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1036_v6).length));
      (_1036_v6)[_index220] = (_this).f5;
      (globalState).f3 = _1029_v1;
      let _1037_v7;
      let _nw219 = new _module.C1();
      _nw219.__ctor((_this).f4, (_this).f5);
      _1037_v7 = _nw219;
      let _1038_v8;
      _1038_v8 = _dafny.Seq.of(new BigNumber(607));
      let _1039_v9;
      _1039_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1029_v1)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1029_v1).length))],(_1029_v1)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1029_v1).length))]);
      r0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_1040_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-712),_dafny.Set.fromElements(true))).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_1033_v3, _1033_v3, new BigNumber((_1038_v8).length), _1033_v3, new BigNumber((_1039_v9).length))).length)));
      r1 = _1033_v3;
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _1041_v0;
      _1041_v0 = false;
      let _1042_i0;
      _1042_i0 = _dafny.ZERO;
      L4: {
        while (_1041_v0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1042_i0)) {
              break L4;
            }
            _1042_i0 = (_1042_i0).plus(_dafny.ONE);
            let _1043_v1;
            _1043_v1 = new BigNumber(-202);
            let _1044_v2;
            _1044_v2 = _dafny.Seq.UnicodeFromString("adlfp");
            let _1045_v3;
            _1045_v3 = _dafny.Seq.of(_1043_v1, _1043_v1);
            let _1046_v4;
            _1046_v4 = _dafny.Seq.of(_1043_v1, (_dafny.ZERO).minus(new BigNumber((_1044_v2).length)), (_dafny.ZERO).minus((_1045_v3)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_1045_v3).length))]));
            _1043_v1 = (_this).fm4(_1046_v4, globalState);
            let _1047_v5;
            _1047_v5 = _dafny.Seq.of(_1041_v0);
            let _1048_v6;
            _1048_v6 = _dafny.Set.fromElements((_this).fm4(_1046_v4, globalState), new BigNumber((_1044_v2).length), _1043_v1, new BigNumber((_1047_v5).length));
            let _1049_v7;
            _1049_v7 = _dafny.Seq.of(_1048_v6, (_1048_v6).Intersect(_1048_v6));
            _1049_v7 = _1049_v7;
            let _1050_v8;
            let _init37 = ((_1051_v1) => function (_1052_i1) {
              return _module.__default.safeDivisionInt(_1052_i1, _1051_v1);
            })(_1043_v1);
            let _nw220 = Array((new BigNumber(22)).toNumber());
            for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw220.length); _i0_37++) {
              _nw220[_i0_37] = _init37(new BigNumber(_i0_37));
            }
            _1050_v8 = _nw220;
            let _1053_v9;
            _1053_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(_1043_v1).isEqualTo((_dafny.ZERO).minus(_1043_v1)),_1050_v8);
            _1053_v9 = _1053_v9;
            _1043_v1 = _1043_v1;
          }
        }
      }
      let _1054_v10;
      _1054_v10 = new BigNumber(929);
      let _1055_v11;
      _1055_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v10,_1054_v10);
      let _1056_v12;
      _1056_v12 = _dafny.Seq.UnicodeFromString("yjdpd");
      let _1057_v13;
      _1057_v13 = _dafny.Seq.of(_1056_v12);
      if (((_dafny.ZERO).minus(new BigNumber((_1057_v13).length))).isLessThanOrEqualTo(new BigNumber((_1055_v11).length))) {
        let _1058_v14;
        let _nw221 = new _module.C1();
        _nw221.__ctor((_this).f4, _module.__default.fm0((_dafny.ZERO).minus(_1054_v10), new BigNumber(957), globalState));
        _1058_v14 = _nw221;
        let _1059_v15;
        let _init38 = ((_1060_v0) => function (_1061_i2) {
          return _1060_v0;
        })(_1041_v0);
        let _nw222 = Array((new BigNumber(13)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw222.length); _i0_38++) {
          _nw222[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1059_v15 = _nw222;
        let _index221 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length));
        (_1059_v15)[_index221] = true;
        let _index222 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length));
        (_1059_v15)[_index222] = (_1041_v0) && (_1041_v0);
        _1054_v10 = _1054_v10;
        let _1062_v16;
        _1062_v16 = _dafny.Seq.of(_1054_v10, _1054_v10);
        let _1063_v17;
        _1063_v17 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1054_v10,new BigNumber((_1062_v16).length)));
        let _1064_v18;
        _1064_v18 = _dafny.Set.fromElements(_1041_v0);
        let _1065_v19;
        _1065_v19 = _dafny.MultiSet.fromElements((_1059_v15)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length))], (_1059_v15)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length))]);
        let _1066_v20;
        let _nw223 = Array((new BigNumber(21)).toNumber());
        _nw223[(_dafny.ZERO).toNumber()] = _1054_v10;
        _nw223[(_dafny.ONE).toNumber()] = (_1054_v10).minus(_1054_v10);
        _nw223[(new BigNumber(2)).toNumber()] = (_1054_v10).multipliedBy((((_1063_v17).contains(_1055_v11)) ? ((_1063_v17).get(_1055_v11)) : (_1054_v10)));
        _nw223[(new BigNumber(3)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(4)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_1054_v10, _1054_v10);
        _nw223[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ypjujowm")).length);
        _nw223[(new BigNumber(7)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_1054_v10, new BigNumber((_1064_v18).length));
        _nw223[(new BigNumber(9)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(10)).toNumber()] = new BigNumber(-406);
        _nw223[(new BigNumber(11)).toNumber()] = (((_1065_v19).contains(_1041_v0)) ? ((_1065_v19).get(_1041_v0)) : (_1054_v10));
        _nw223[(new BigNumber(12)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(13)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_1054_v10);
        _nw223[(new BigNumber(15)).toNumber()] = new BigNumber((_1056_v12).length);
        _nw223[(new BigNumber(16)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(17)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(18)).toNumber()] = _1054_v10;
        _nw223[(new BigNumber(19)).toNumber()] = (new BigNumber(336)).minus(_1054_v10);
        _nw223[(new BigNumber(20)).toNumber()] = _1054_v10;
        _1066_v20 = _nw223;
        let _index223 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1066_v20).length));
        (_1066_v20)[_index223] = new BigNumber(468);
        let _index224 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1066_v20).length));
        (_1066_v20)[_index224] = _module.__default.safeModuloInt(_1054_v10, _1054_v10);
        let _index225 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length));
        let _rhs198 = (_this).fm5(_1054_v10, (_1054_v10).isLessThanOrEqualTo((((_1055_v11).contains((_1066_v20)[_module.__default.safeIndex(new BigNumber(643), new BigNumber((_1066_v20).length))])) ? ((_1055_v11).get((_1066_v20)[_module.__default.safeIndex(new BigNumber(643), new BigNumber((_1066_v20).length))])) : (new BigNumber((_1056_v12).length)))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_1059_v15)[_module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length))])).length), (_1066_v20)[_module.__default.safeIndex(new BigNumber(643), new BigNumber((_1066_v20).length))], globalState);
        let _rhs199 = _1066_v20;
        let _rhs200 = _1041_v0;
        let _lhs105 = _1059_v15;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_1059_v15).length));
        _1041_v0 = _rhs198;
        _1066_v20 = _rhs199;
        _lhs105[_lhs106] = _rhs200;
      } else {
        _1056_v12 = _1056_v12;
        _1041_v0 = true;
        let _1067_v21;
        _1067_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1041_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-749)), function (_1068_i3) {
          return (_this).f5;
        }));
        let _1069_v22;
        _1069_v22 = _dafny.Set.fromElements(_1041_v0);
        _1041_v0 = _module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1056_v12, _module.__default.safeIndex(_1054_v10, new BigNumber((_1056_v12).length)), (_this).f5), _1056_v12)).length)), ((true) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((((_1067_v21).contains(_1041_v0)) ? ((_1067_v21).get(_1041_v0)) : (_1056_v12))).length))))) : (_1054_v10)), (_dafny.Set.fromElements(_1041_v0, _1041_v0, _1041_v0)).IsProperSubsetOf(_1069_v22), false, globalState);
        let _1070_v23;
        _1070_v23 = _module.D14.create_DC32(_1054_v10, _1041_v0, _1041_v0);
        let _1071_v24;
        _1071_v24 = _dafny.Seq.of((_1070_v23).dtor_cf49);
        let _1072_v25;
        let _nw224 = Array((new BigNumber(23)).toNumber()).fill(false);
        _1072_v25 = _nw224;
        let _1073_v26;
        _1073_v26 = _dafny.Seq.of(_1072_v25);
        let _1074_v27;
        let _nw225 = Array((new BigNumber(3)).toNumber());
        _nw225[(_dafny.ZERO).toNumber()] = new BigNumber((_1071_v24).length);
        _nw225[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1054_v10,new BigNumber((_1073_v26).length))).length);
        _nw225[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_1054_v10);
        _1074_v27 = _nw225;
        let _index226 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1074_v27).length));
        (_1074_v27)[_index226] = _1054_v10;
        let _index227 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1074_v27).length));
        (_1074_v27)[_index227] = _1054_v10;
        _1054_v10 = (_dafny.ZERO).minus(new BigNumber((_1073_v26).length));
      }
      let _1075_v29;
      _1075_v29 = _dafny.Seq.of(_1054_v10, new BigNumber((function () {
        let _coll40 = new _dafny.Set();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(678), new BigNumber(-935))) {
          let _1076_v28 = _compr_40;
          if (((new BigNumber(678)).isLessThanOrEqualTo(_1076_v28)) && ((_1076_v28).isLessThan(new BigNumber(-935)))) {
            _coll40.add(_module.__default.safeDivisionInt(_1076_v28, _1054_v10));
          }
        }
        return _coll40;
      }()).length));
      _1054_v10 = ((_1041_v0) ? (_1054_v10) : (new BigNumber((_dafny.Seq.Concat(_1075_v29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-255)), ((_1077_v10) => function (_1078_i4) {
        return _1077_v10;
      })(_1054_v10)))).length)));
      _1054_v10 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yjknrytek"), _dafny.Seq.Concat(_1056_v12, _1056_v12))).length);
      let _1079_v30;
      _1079_v30 = _dafny.MultiSet.fromElements(_1054_v10);
      _1054_v10 = new BigNumber((_1079_v30).cardinality());
      _1054_v10 = (((_1055_v11).contains(_1054_v10)) ? ((_1055_v11).get(_1054_v10)) : (_1054_v10));
      let _1080_v31;
      let _nw226 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _1080_v31 = _nw226;
      r0 = _1080_v31;
      return r0;
    }
    m11(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1081_v0;
      let _init39 = function (_1082_i0) {
        return _dafny.Seq.of(!(false), false, !(false));
      };
      let _nw227 = Array((new BigNumber(23)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw227.length); _i0_39++) {
        _nw227[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1081_v0 = _nw227;
      let _1083_v1;
      _1083_v1 = false;
      let _1084_v2;
      _1084_v2 = _dafny.Seq.of(_1083_v1, _1083_v1, _1083_v1, false);
      let _index228 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length));
      (_1081_v0)[_index228] = _1084_v2;
      let _1085_v3;
      _1085_v3 = new BigNumber(794);
      let _1086_v4;
      _1086_v4 = _dafny.Set.fromElements(_1083_v1);
      let _index229 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length));
      (_1081_v0)[_index229] = _dafny.Seq.update(_dafny.Seq.Concat(_1084_v2, _1084_v2), _module.__default.safeIndex(_1085_v3, new BigNumber((_dafny.Seq.Concat(_1084_v2, _1084_v2)).length)), (_1086_v4).IsDisjointFrom(_1086_v4));
      _1083_v1 = _1083_v1;
      let _1087_i1;
      _1087_i1 = _dafny.ZERO;
      L5: {
        while (false) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1087_i1)) {
              break L5;
            }
            _1087_i1 = (_1087_i1).plus(_dafny.ONE);
            let _1088_v5;
            _1088_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f5);
            let _1089_v6;
            _1089_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1088_v5).length),_1083_v1);
            if ((_1083_v1) || (!(true) || ((((_1089_v6).contains(_1085_v3)) ? ((_1089_v6).get(_1085_v3)) : (_1083_v1))))) {
              let _1090_v7;
              _1090_v7 = _dafny.Seq.UnicodeFromString("al");
              let _1091_v8;
              _1091_v8 = new _dafny.CodePoint('d'.codePointAt(0));
              let _1092_v9;
              _1092_v9 = _dafny.Seq.of(_1085_v3, _1085_v3);
              let _rhs201 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("wbrvy"), _module.__default.safeIndex((_1092_v9)[_module.__default.safeIndex(_1085_v3, new BigNumber((_1092_v9).length))], new BigNumber((_dafny.Seq.UnicodeFromString("wbrvy")).length)), _1091_v8);
              let _rhs202 = !((_1083_v1) || (_1083_v1));
              let _rhs203 = _1091_v8;
              _1090_v7 = _rhs201;
              _1083_v1 = _rhs202;
              _1091_v8 = _rhs203;
              _1083_v1 = _1083_v1;
              _1083_v1 = _1083_v1;
              let _1093_v10;
              _1093_v10 = _dafny.Set.fromElements(_1085_v3);
              let _1094_v11;
              _1094_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(!(_1083_v1), _1083_v1, _1083_v1), _dafny.Seq.update((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))], _module.__default.safeIndex(new BigNumber((_1093_v10).length), new BigNumber(((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))]).length)), _1083_v1)), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(174)), ((_1095_v1, _1096_v9) => function (_1097_i2) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1095_v1,(_dafny.ZERO).minus(new BigNumber((_1096_v9).length)));
              })(_1083_v1, _1092_v9))).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(_1083_v1), _1083_v1, _1083_v1), _dafny.Seq.update((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))], _module.__default.safeIndex(new BigNumber((_1093_v10).length), new BigNumber(((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))]).length)), _1083_v1))).length)), _1083_v1),((_1083_v1) ? (_1083_v1) : (!(_1083_v1))));
              let _1098_v12;
              _1098_v12 = _module.D6.create_DC15(_1083_v1, _1083_v1, _1083_v1);
              _1083_v1 = !(!((((_1094_v11).contains(_1084_v2)) ? ((_1094_v11).get(_1084_v2)) : ((_1098_v12).dtor_cf21))));
              _1085_v3 = (_dafny.ZERO).minus(_1085_v3);
            } else {
              let _1099_v13;
              let _nw228 = new _module.C1();
              _nw228.__ctor((_this).f4, (_this).f5);
              _1099_v13 = _nw228;
              let _1100_v14;
              let _nw229 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
              _1100_v14 = _nw229;
              let _1101_v15;
              _1101_v15 = _module.D11.create_DC25(_1083_v1, new BigNumber(157), _1100_v14, (_dafny.ZERO).minus(_1085_v3));
              let _1102_v16;
              _1102_v16 = _dafny.Set.fromElements(new BigNumber(89), new BigNumber(918), _1085_v3);
              let _1103_v17;
              _1103_v17 = _dafny.Seq.UnicodeFromString("v");
              let _1104_v18;
              _1104_v18 = _dafny.MultiSet.fromElements((_this).fm5(_1085_v3, _1083_v1, new BigNumber(-836), _1085_v3, globalState));
              let _1105_v19;
              _1105_v19 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), (_this).f5, (_this).f5, (_this).f5);
              let _1106_v20;
              _1106_v20 = _dafny.Seq.of(new BigNumber((_1103_v17).length), new BigNumber((_1105_v19).cardinality()));
              let _1107_v21;
              _1107_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1083_v1);
              let _1108_v22;
              let _nw230 = Array((new BigNumber(26)).toNumber());
              _nw230[(_dafny.ZERO).toNumber()] = _1083_v1;
              _nw230[(_dafny.ONE).toNumber()] = true;
              _nw230[(new BigNumber(2)).toNumber()] = false;
              _nw230[(new BigNumber(3)).toNumber()] = false;
              _nw230[(new BigNumber(4)).toNumber()] = _module.__default.fm1((_dafny.ZERO).minus(_1085_v3), _1085_v3, _1083_v1, _1083_v1, globalState);
              _nw230[(new BigNumber(5)).toNumber()] = (_1101_v15).dtor_cf38;
              _nw230[(new BigNumber(6)).toNumber()] = (_1085_v3).isLessThanOrEqualTo((_dafny.ZERO).minus(_1085_v3));
              _nw230[(new BigNumber(7)).toNumber()] = !(_1083_v1);
              _nw230[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.FromArray(_1084_v2)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1083_v1));
              _nw230[(new BigNumber(9)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(10)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(11)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(12)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(13)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(14)).toNumber()] = (_this).fm5(new BigNumber((_1102_v16).length), _1083_v1, new BigNumber((_1103_v17).length), _1085_v3, globalState);
              _nw230[(new BigNumber(15)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(16)).toNumber()] = (_1086_v4).equals(_1086_v4);
              _nw230[(new BigNumber(17)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(18)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(19)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(20)).toNumber()] = !(_1083_v1);
              _nw230[(new BigNumber(21)).toNumber()] = (_1085_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1104_v18)).length));
              _nw230[(new BigNumber(22)).toNumber()] = (_1099_v13).fm5(_1085_v3, _1083_v1, (_1106_v20)[_module.__default.safeIndex(new BigNumber(((_1107_v21).update(_1083_v1, _1083_v1)).length), new BigNumber((_1106_v20).length))], _1085_v3, globalState);
              _nw230[(new BigNumber(23)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(24)).toNumber()] = _1083_v1;
              _nw230[(new BigNumber(25)).toNumber()] = false;
              _1108_v22 = _nw230;
              let _index230 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              (_1108_v22)[_index230] = _1083_v1;
              let _index231 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              (_1108_v22)[_index231] = _1083_v1;
              let _1109_v23;
              _1109_v23 = _module.D1.create_DC2(_1085_v3, new BigNumber((_dafny.Seq.of(true, (_1108_v22)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length))])).length));
              let _1110_v24;
              _1110_v24 = _dafny.Map.Empty.slice().updateUnsafe(true,_1085_v3);
              let _1111_v25;
              _1111_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v23,(((_1110_v24).contains(true)) ? ((_1110_v24).get(true)) : (_1085_v3)));
              let _1112_v26;
              _1112_v26 = _dafny.MultiSet.fromElements(new BigNumber((_1106_v20).length));
              let _1113_v27;
              _1113_v27 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yofklgq"));
              let _index232 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              let _index233 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              let _rhs204 = (new BigNumber(807)).plus(((_dafny.ZERO).minus(_1085_v3)).plus(new BigNumber(-678)));
              let _rhs205 = ((_1099_v13).fm4(_dafny.Seq.of(new BigNumber((_1112_v26).cardinality())), globalState)).multipliedBy((_1085_v3).plus(_1085_v3));
              let _rhs206 = ((_1111_v25).Merge(_1111_v25)).update(_1109_v23, _1085_v3);
              let _rhs207 = _module.__default.fm1(_1085_v3, new BigNumber((_module.__default.fm29(new BigNumber((_1113_v27).length), _1083_v1, _1085_v3, globalState)).length), (_1108_v22)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length))], (_1108_v22)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length))], globalState);
              let _rhs208 = _1083_v1;
              let _lhs107 = _1108_v22;
              let _lhs108 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              let _lhs109 = _1108_v22;
              let _lhs110 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              _1085_v3 = _rhs204;
              _1085_v3 = _rhs205;
              _1111_v25 = _rhs206;
              _lhs107[_lhs108] = _rhs207;
              _lhs109[_lhs110] = _rhs208;
              let _1114_v28;
              _1114_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v1,(_this).f5);
              let _1115_v29;
              _1115_v29 = _module.D5.create_DC11(_1114_v28);
              let _1116_v30;
              _1116_v30 = _module.D2.create_DC3(_1108_v22);
              let _1117_v31;
              let _nw231 = Array((new BigNumber(4)).toNumber());
              _nw231[(_dafny.ZERO).toNumber()] = _1108_v22;
              _nw231[(_dafny.ONE).toNumber()] = _1108_v22;
              _nw231[(new BigNumber(2)).toNumber()] = (_1116_v30).dtor_cf4;
              _nw231[(new BigNumber(3)).toNumber()] = _1108_v22;
              _1117_v31 = _nw231;
              let _1118_v32;
              _1118_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1117_v31,_1083_v1);
              let _rhs209 = (((_1118_v32).contains(_1117_v31)) ? ((_1118_v32).get(_1117_v31)) : ((!((_1108_v22)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length))])) === (_1083_v1)));
              let _rhs210 = _1085_v3;
              let _rhs211 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1106_v20, _1106_v20)).length)), _1085_v3);
              let _rhs212 = _1100_v14;
              let _rhs213 = _1115_v29;
              _1083_v1 = _rhs209;
              _1085_v3 = _rhs210;
              _1085_v3 = _rhs211;
              _1100_v14 = _rhs212;
              _1115_v29 = _rhs213;
              let _index234 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length));
              (_1108_v22)[_index234] = (_1108_v22)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_1108_v22).length))];
            }
            let _1119_v33;
            _1119_v33 = _dafny.Seq.UnicodeFromString("rvdnequn");
            let _1120_v34;
            _1120_v34 = _module.D2.create_DC6(new BigNumber((_1119_v33).length));
            let _1121_v35;
            _1121_v35 = _module.D2.create_DC7(_1120_v34);
            let _1122_v36;
            _1122_v36 = _module.D2.create_DC7(_1121_v35);
            let _1123_v37;
            _1123_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1122_v36);
            let _1124_v38;
            _1124_v38 = _dafny.Set.fromElements(_1123_v37);
            let _1125_v39;
            _1125_v39 = _module.D7.create_DC17(_1083_v1, _1083_v1, _1083_v1, _1083_v1);
            let _1126_v40;
            _1126_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-176),_1124_v38);
            let _1127_v41;
            _1127_v41 = _dafny.Seq.of(_1124_v38, _1124_v38, _1124_v38, _1124_v38);
            let _1128_v42;
            let _nw232 = Array((new BigNumber(27)).toNumber());
            _nw232[(_dafny.ZERO).toNumber()] = _1124_v38;
            _nw232[(_dafny.ONE).toNumber()] = (_1124_v38).Difference(_dafny.Set.fromElements(_1123_v37, _dafny.Map.Empty.slice().updateUnsafe((_1125_v39).dtor_cf24,_1122_v36), _1123_v37));
            _nw232[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_1123_v37);
            _nw232[(new BigNumber(3)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(4)).toNumber()] = (_1124_v38).Intersect(_1124_v38);
            _nw232[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_1123_v37)).Union(_1124_v38);
            _nw232[(new BigNumber(6)).toNumber()] = (((_1126_v40).contains(_1085_v3)) ? ((_1126_v40).get(_1085_v3)) : (_1124_v38));
            _nw232[(new BigNumber(7)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(8)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(9)).toNumber()] = (_1124_v38).Union(_1124_v38);
            _nw232[(new BigNumber(10)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(11)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(12)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_1123_v37);
            _nw232[(new BigNumber(14)).toNumber()] = (_1127_v41)[_module.__default.safeIndex(new BigNumber((_1119_v33).length), new BigNumber((_1127_v41).length))];
            _nw232[(new BigNumber(15)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(_1123_v37, _dafny.Map.Empty.slice().updateUnsafe(true,_1122_v36), _dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1122_v36));
            _nw232[(new BigNumber(17)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(18)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(19)).toNumber()] = (_1124_v38).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1122_v36)));
            _nw232[(new BigNumber(20)).toNumber()] = (_1124_v38).Intersect(_dafny.Set.fromElements(_1123_v37, _1123_v37));
            _nw232[(new BigNumber(21)).toNumber()] = (_1124_v38).Union(_1124_v38);
            _nw232[(new BigNumber(22)).toNumber()] = (_1124_v38).Union(_1124_v38);
            _nw232[(new BigNumber(23)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(24)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(25)).toNumber()] = _1124_v38;
            _nw232[(new BigNumber(26)).toNumber()] = _1124_v38;
            _1128_v42 = _nw232;
            let _index235 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1128_v42).length));
            (_1128_v42)[_index235] = _1124_v38;
            let _1129_v43;
            _1129_v43 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1085_v3,_1083_v1));
            let _1130_v44;
            _1130_v44 = _dafny.MultiSet.fromElements(_1089_v6);
            let _1131_v45;
            _1131_v45 = _dafny.Seq.of(_1085_v3, _1085_v3);
            let _1132_v46;
            _1132_v46 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_1131_v45, _1131_v45, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-754)), ((_1133_v3) => function (_1134_i3) {
              return _1133_v3;
            })(_1085_v3)))).length));
            let _index236 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1128_v42).length));
            let _rhs214 = _1124_v38;
            let _rhs215 = _1119_v33;
            let _rhs216 = (_dafny.ZERO).minus(((_1083_v1) ? (new BigNumber(((_dafny.MultiSet.FromArray(_1129_v43)).Intersect(_1130_v44)).cardinality())) : (_1085_v3)));
            let _rhs217 = new BigNumber((_1132_v46).length);
            let _rhs218 = _module.__default.safeDivisionInt(new BigNumber(((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))]).length), _1085_v3);
            let _lhs111 = _1128_v42;
            let _lhs112 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1128_v42).length));
            _lhs111[_lhs112] = _rhs214;
            _1119_v33 = _rhs215;
            _1085_v3 = _rhs216;
            _1085_v3 = _rhs217;
            _1085_v3 = _rhs218;
            let _1135_v47;
            let _nw233 = Array((new BigNumber(7)).toNumber()).fill(false);
            _1135_v47 = _nw233;
            let _index237 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1135_v47).length));
            (_1135_v47)[_index237] = (new BigNumber((_1131_v45).length)).isEqualTo(_1085_v3);
            let _index238 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1135_v47).length));
            (_1135_v47)[_index238] = _1083_v1;
            let _1136_v48;
            let _init40 = function (_1137_i4) {
              return _module.__default.safeDivisionInt(_1137_i4, new BigNumber(793));
            };
            let _nw234 = Array((new BigNumber(23)).toNumber());
            for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw234.length); _i0_40++) {
              _nw234[_i0_40] = _init40(new BigNumber(_i0_40));
            }
            _1136_v48 = _nw234;
            let _index239 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length));
            (_1136_v48)[_index239] = _1085_v3;
            let _1138_v49;
            _1138_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v1,_1085_v3);
            let _index240 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1135_v47).length));
            let _index241 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1135_v47).length));
            let _index242 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length));
            let _rhs219 = (((_1138_v49).equals(_1138_v49)) ? (((_1083_v1) ? (_1083_v1) : (_1083_v1))) : (true));
            let _rhs220 = _1083_v1;
            let _rhs221 = ((_1083_v1) ? (_1085_v3) : ((_1085_v3).multipliedBy(_1085_v3)));
            let _lhs113 = _1135_v47;
            let _lhs114 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1135_v47).length));
            let _lhs115 = _1135_v47;
            let _lhs116 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1135_v47).length));
            let _lhs117 = _1136_v48;
            let _lhs118 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length));
            _lhs113[_lhs114] = _rhs219;
            _lhs115[_lhs116] = _rhs220;
            _lhs117[_lhs118] = _rhs221;
            let _1139_v50;
            _1139_v50 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("shaofskxb"));
            let _index243 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1135_v47).length));
            let _rhs222 = !((_1139_v50).IsProperSubsetOf((_1139_v50).update(_dafny.Seq.UnicodeFromString("tmd"), _module.__default.abs((_1136_v48)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length))]))));
            let _rhs223 = _module.__default.safeModuloInt((_1136_v48)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length))], (_1136_v48)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1136_v48).length))]);
            let _lhs119 = _1135_v47;
            let _lhs120 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_1135_v47).length));
            _lhs119[_lhs120] = _rhs222;
            _1085_v3 = _rhs223;
          }
        }
      }
      let _1140_v51;
      _1140_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1085_v3,_1083_v1);
      let _1141_v52;
      _1141_v52 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_1140_v51, _1140_v51)).length), _1085_v3, _1085_v3, _1085_v3);
      let _1142_v53;
      let _init41 = ((_1143_v3) => function (_1144_i5) {
        return _module.__default.safeDivisionInt(_1144_i5, _1143_v3);
      })(_1085_v3);
      let _nw235 = Array((new BigNumber(17)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw235.length); _i0_41++) {
        _nw235[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1142_v53 = _nw235;
      let _1145_v54;
      _1145_v54 = _dafny.MultiSet.fromElements((_this).f5);
      let _1146_v55;
      _1146_v55 = _dafny.Seq.of(_1085_v3);
      let _rhs224 = ((!((_1145_v54).IsDisjointFrom(_1145_v54))) ? ((_1141_v52).Intersect(_1141_v52)) : ((_1141_v52).Union(_1141_v52)));
      let _rhs225 = _1142_v53;
      let _rhs226 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1085_v3, _1085_v3))).multipliedBy(_module.__default.safeDivisionInt((_1146_v55)[_module.__default.safeIndex(_1085_v3, new BigNumber((_1146_v55).length))], _1085_v3)));
      _1141_v52 = _rhs224;
      _1142_v53 = _rhs225;
      _1085_v3 = _rhs226;
      let _hi7 = _1085_v3;
      for (let _1147_i6 = (new BigNumber(((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))]).length)).plus(_1085_v3); _1147_i6.isLessThan(_hi7); _1147_i6 = _1147_i6.plus(_dafny.ONE)) {
        let _1148_v56;
        _1148_v56 = _module.D14.create_DC31(_1147_i6);
        let _1149_v57;
        _1149_v57 = _dafny.MultiSet.fromElements(_1148_v56);
        let _1150_v58;
        _1150_v58 = _dafny.Seq.of(_1148_v56, _1148_v56, _1148_v56, _1148_v56);
        _1083_v1 = (_1149_v57).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1150_v58));
        let _1151_v59;
        let _nw236 = new _module.C1();
        _nw236.__ctor((_this).f4, (_this).f5);
        _1151_v59 = _nw236;
        let _1152_v60;
        _1152_v60 = _dafny.Seq.UnicodeFromString("cglegqjc");
        _1152_v60 = _1152_v60;
        let _1153_v61;
        _1153_v61 = _module.D6.create_DC15(_1083_v1, !(_1083_v1) || (_1083_v1), _1083_v1);
        _1153_v61 = _1153_v61;
      }
      _1083_v1 = (new BigNumber(((_1081_v0)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_1081_v0).length))]).length)).isLessThanOrEqualTo(_1085_v3);
      r0 = (_1140_v51).update(_1085_v3, !(_1083_v1));
      return r0;
    }
    m12(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _1154_v0;
      _1154_v0 = false;
      let _1155_v1;
      _1155_v1 = _dafny.Seq.UnicodeFromString("ms");
      let _1156_v2;
      _1156_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_1155_v1);
      let _1157_v3;
      _1157_v3 = _dafny.MultiSet.fromElements(p0, p0, new BigNumber((_1156_v2).length), p0);
      let _1158_v4;
      _1158_v4 = _dafny.MultiSet.fromElements(false, true);
      let _1159_v5;
      _1159_v5 = _dafny.Seq.of(true);
      let _1160_v6;
      _1160_v6 = _dafny.Seq.of((_1159_v5)[_module.__default.safeIndex((_this).fm4(_dafny.Seq.of(p0), globalState), new BigNumber((_1159_v5).length))]);
      let _1161_v7;
      let _nw237 = Array((new BigNumber(10)).toNumber());
      _nw237[(_dafny.ZERO).toNumber()] = p0;
      _nw237[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
      _nw237[(new BigNumber(2)).toNumber()] = p0;
      _nw237[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1157_v3).cardinality()), p0);
      _nw237[(new BigNumber(4)).toNumber()] = p0;
      _nw237[(new BigNumber(5)).toNumber()] = p0;
      _nw237[(new BigNumber(6)).toNumber()] = p0;
      _nw237[(new BigNumber(7)).toNumber()] = (((_1158_v4).contains(_1154_v0)) ? ((_1158_v4).get(_1154_v0)) : (p0));
      _nw237[(new BigNumber(8)).toNumber()] = new BigNumber(((_module.__default.fm30(_1154_v0, _1160_v6, globalState)).update(p0, _module.__default.abs(p0))).cardinality());
      _nw237[(new BigNumber(9)).toNumber()] = p0;
      _1161_v7 = _nw237;
      let _index244 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
      (_1161_v7)[_index244] = new BigNumber((_1155_v1).length);
      let _index245 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
      (_1161_v7)[_index245] = p0;
      let _1162_v8;
      let _nw238 = Array((new BigNumber(23)).toNumber()).fill(false);
      _1162_v8 = _nw238;
      (globalState).f3 = _1162_v8;
      let _1163_v9;
      let _nw239 = new _module.C0();
      _nw239.__ctor();
      _1163_v9 = _nw239;
      _1163_v9 = _1163_v9;
      let _hi8 = (_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))];
      for (let _1164_i0 = p0; _1164_i0.isLessThan(_hi8); _1164_i0 = _1164_i0.plus(_dafny.ONE)) {
        let _1165_v10;
        _1165_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(p0,(_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))]));
        let _1166_v11;
        _1166_v11 = _module.D14.create_DC30(_1165_v10);
        let _pat_let_tv30 = _1165_v10;
        let _1167_v12;
        _1167_v12 = _dafny.Seq.of(function (_pat_let23_0) {
          return function (_1168_dt__update__tmp_h0) {
            return function (_pat_let24_0) {
              return function (_1169_dt__update_hcf47_h0) {
                return _module.D14.create_DC30(_1169_dt__update_hcf47_h0);
              }(_pat_let24_0);
            }(_pat_let_tv30);
          }(_pat_let23_0);
        }(_1166_v11), _1166_v11, _1166_v11, _1166_v11, _1166_v11);
        let _1170_v13;
        _1170_v13 = _dafny.Seq.of(_1167_v12, _1167_v12, _1167_v12, _1167_v12);
        _1167_v12 = (_1170_v13)[_module.__default.safeIndex(_1164_i0, new BigNumber((_1170_v13).length))];
        let _1171_v14;
        _1171_v14 = _dafny.Set.fromElements(_1154_v0);
        let _index246 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
        let _rhs227 = _dafny.Seq.update(_dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, _1164_i0)), _module.__default.safeDivisionInt((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))], new BigNumber((_1171_v14).length)), new BigNumber((_1157_v3).cardinality())), _module.__default.safeIndex((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))], new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, _1164_i0)), _module.__default.safeDivisionInt((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))], new BigNumber((_1171_v14).length)), new BigNumber((_1157_v3).cardinality()))).length)), new BigNumber((((_1154_v0) ? (_1155_v1) : (_1155_v1))).length));
        let _rhs228 = (_1164_i0).plus(_1164_i0);
        let _lhs121 = globalState;
        let _lhs122 = _1161_v7;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
        _lhs121.f2 = _rhs227;
        _lhs122[_lhs123] = _rhs228;
        let _1172_v15;
        let _nw240 = new _module.C2();
        _nw240.__ctor((_this).f4, _module.__default.fm0(new BigNumber(684), p0, globalState));
        _1172_v15 = _nw240;
        let _1173_v16;
        _1173_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_1154_v0);
        let _1174_v17;
        let _nw241 = new _module.C1();
        _nw241.__ctor((_this).f4, (_this).f5);
        _1174_v17 = _nw241;
        let _1175_v18;
        _1175_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_1174_v17);
        let _1176_v19;
        _1176_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_1154_v0) ? (_1164_i0) : (new BigNumber((_1175_v18).length))),(_dafny.ZERO).minus((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))]));
        let _1177_v20;
        _1177_v20 = _dafny.MultiSet.fromElements(_1162_v8, _1162_v8, _1162_v8);
        let _index247 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
        let _rhs229 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1164_i0, (((_1177_v20).contains(_1162_v8)) ? ((_1177_v20).get(_1162_v8)) : (_1164_i0))));
        let _rhs230 = _1172_v15;
        let _rhs231 = _1173_v16;
        let _rhs232 = !(!((((_1173_v16).contains(((_dafny.ZERO).minus(_1164_i0)).isLessThanOrEqualTo((_dafny.ZERO).minus((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))])))) ? ((_1173_v16).get(((_dafny.ZERO).minus(_1164_i0)).isLessThanOrEqualTo((_dafny.ZERO).minus((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))])))) : ((false) && (_1154_v0)))));
        let _rhs233 = ((_1176_v19).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))],p0))).Merge((_1176_v19).Merge(function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of _dafny.IntegerRange(new BigNumber(28), new BigNumber(15))) {
            let _1178_v21 = _compr_41;
            if (((new BigNumber(28)).isLessThanOrEqualTo(_1178_v21)) && ((_1178_v21).isLessThan(new BigNumber(15)))) {
              _coll41.push([_module.__default.safeDivisionInt(_1178_v21, new BigNumber((_dafny.Seq.UnicodeFromString("nukgrcx")).length)),p0]);
            }
          }
          return _coll41;
        }()));
        let _lhs124 = _1161_v7;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length));
        _lhs124[_lhs125] = _rhs229;
        _1172_v15 = _rhs230;
        _1173_v16 = _rhs231;
        _1154_v0 = _rhs232;
        _1176_v19 = _rhs233;
        let _1179_v22;
        _1179_v22 = new _dafny.CodePoint('e'.codePointAt(0));
        _1179_v22 = _1179_v22;
      }
      _1154_v0 = _1154_v0;
      r1 = _1154_v0;
      r0 = _module.__default.fm16(false, p0, globalState);
      r1 = _1154_v0;
      let _1180_v23;
      _1180_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_module.__default.safeDivisionInt(new BigNumber(-68), (_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))]));
      r2 = (((_1180_v23).contains(false)) ? ((_1180_v23).get(false)) : (_module.__default.safeDivisionInt(p0, (_1161_v7)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_1161_v7).length))])));
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f11 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f11, f12, f4, f5) {
      let _this = this;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(!(false), false)).length), _this.f11);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (((!(true)) ? (false) : (!(true)))) {
        return true;
      } else {
        return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(true), _dafny.Seq.of(true));
      }
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return !((_module.__default.safeModuloInt(_this.f11, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).isLessThan((new BigNumber(361)).plus((_this).f12)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      (_this).f11 = _this.f11;
      let _1181_v0;
      _1181_v0 = true;
      _1181_v0 = !((_module.__default.safeModuloInt(_this.f11, _this.f11)).isLessThanOrEqualTo((_this).f12));
      let _1182_v1;
      _1182_v1 = _dafny.Seq.of(_dafny.Seq.of(false));
      let _1183_v2;
      _1183_v2 = _dafny.Seq.UnicodeFromString("mhg");
      let _rhs234 = _1182_v1;
      let _rhs235 = !_dafny.areEqual(_1183_v2, _1183_v2);
      _1182_v1 = _rhs234;
      _1181_v0 = _rhs235;
      let _1184_v3;
      _1184_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_1181_v0);
      let _1185_v4;
      _1185_v4 = _dafny.MultiSet.fromElements(_this.f11, (_dafny.ZERO).minus(_this.f11), new BigNumber(691), (_this).f12, new BigNumber((_1184_v3).length));
      _1185_v4 = _1185_v4;
      let _hi9 = ((_1181_v0) ? ((_dafny.ZERO).minus((_this).f12)) : ((_dafny.ZERO).minus(_this.f11)));
      for (let _1186_i0 = (_this).f12; _1186_i0.isLessThan(_hi9); _1186_i0 = _1186_i0.plus(_dafny.ONE)) {
        let _1187_v5;
        let _nw242 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1187_v5 = _nw242;
        let _index248 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
        (_1187_v5)[_index248] = _1181_v0;
        let _1188_v6;
        _1188_v6 = !(_1181_v0);
        let _1189_v7;
        _1189_v7 = _dafny.Seq.of(_1181_v0, (_1188_v6), _1181_v0);
        let _index249 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
        let _rhs236 = _1181_v0;
        let _rhs237 = _1189_v7;
        let _lhs126 = _1187_v5;
        let _lhs127 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
        _lhs126[_lhs127] = _rhs236;
        _1189_v7 = _rhs237;
        let _1190_v8;
        _1190_v8 = _dafny.Set.fromElements(_1187_v5, _1187_v5);
        let _1191_v9;
        _1191_v9 = _1190_v8;
        let _1192_v13;
        let _nw243 = Array((new BigNumber(17)).toNumber());
        _nw243[(_dafny.ZERO).toNumber()] = function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(919), new BigNumber(303))) {
            let _1193_v10 = _compr_42;
            if (((new BigNumber(919)).isLessThanOrEqualTo(_1193_v10)) && ((_1193_v10).isLessThan(new BigNumber(303)))) {
              _coll42.push([_module.__default.safeDivisionInt(_1193_v10, _1186_i0),true]);
            }
          }
          return _coll42;
        }();
        _nw243[(_dafny.ONE).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(2)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(3)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(4)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(5)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1186_i0,_1181_v0);
        _nw243[(new BigNumber(7)).toNumber()] = (_1184_v3).update(_1186_i0, (((_1184_v3).contains(_this.f11)) ? ((_1184_v3).get(_this.f11)) : (_1181_v0)));
        _nw243[(new BigNumber(8)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(9)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(10)).toNumber()] = function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(737), new BigNumber(-102))) {
            let _1194_v11 = _compr_43;
            if (((new BigNumber(737)).isLessThanOrEqualTo(_1194_v11)) && ((_1194_v11).isLessThan(new BigNumber(-102)))) {
              _coll43.push([_module.__default.safeModuloInt(_1194_v11, _this.f11),_1181_v0]);
            }
          }
          return _coll43;
        }();
        _nw243[(new BigNumber(11)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(12)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(13)).toNumber()] = function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(-117), new BigNumber(-575))) {
            let _1195_v12 = _compr_44;
            if (((new BigNumber(-117)).isLessThanOrEqualTo(_1195_v12)) && ((_1195_v12).isLessThan(new BigNumber(-575)))) {
              _coll44.push([(_1195_v12).minus(_1186_i0),(_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]]);
            }
          }
          return _coll44;
        }();
        _nw243[(new BigNumber(14)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(15)).toNumber()] = _1184_v3;
        _nw243[(new BigNumber(16)).toNumber()] = _1184_v3;
        _1192_v13 = _nw243;
        let _1196_v14;
        _1196_v14 = _module.D4.create_DC10(_1191_v9, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), function (_1197_i1) {
  return (_this).f12;
})).length), _1183_v2, _1192_v13);
        let _1198_v15;
        _1198_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1187_v5,(_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]);
        if ((_module.__default.safeDivisionInt(_1186_i0, (_1196_v14).dtor_cf12)).isLessThanOrEqualTo(new BigNumber((_1198_v15).length))) {
          let _1199_v16;
          let _init42 = ((_1200_v0) => function (_1201_i2) {
            return _dafny.Seq.of(_dafny.MultiSet.fromElements(_1200_v0));
          })(_1181_v0);
          let _nw244 = Array((new BigNumber(13)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw244.length); _i0_42++) {
            _nw244[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1199_v16 = _nw244;
          let _1202_v17;
          _1202_v17 = _dafny.MultiSet.fromElements((_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))], (_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]);
          let _1203_v18;
          _1203_v18 = _dafny.Seq.of(_1202_v17);
          let _index250 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1199_v16).length));
          (_1199_v16)[_index250] = _1203_v18;
          let _1204_v19;
          _1204_v19 = _module.D1.create_DC1(_1202_v17);
          let _index251 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1199_v16).length));
          (_1199_v16)[_index251] = _dafny.Seq.Concat(_dafny.Seq.of((_1204_v19).dtor_cf1, _1202_v17, _1202_v17), _1203_v18);
          let _index252 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
          let _rhs238 = true;
          let _rhs239 = _module.__default.fm1((_1186_i0).multipliedBy(_1186_i0), _this.f11, ((_1181_v0) ? (_1181_v0) : (!(_1181_v0))), (_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))], globalState);
          let _lhs128 = _1187_v5;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
          _lhs128[_lhs129] = _rhs238;
          _1181_v0 = _rhs239;
          let _1205_v20;
          let _nw245 = Array((new BigNumber(3)).toNumber());
          _nw245[(_dafny.ZERO).toNumber()] = _1183_v2;
          _nw245[(_dafny.ONE).toNumber()] = _1183_v2;
          _nw245[(new BigNumber(2)).toNumber()] = _1183_v2;
          _1205_v20 = _nw245;
          _1205_v20 = _1205_v20;
          let _1206_v21;
          _1206_v21 = _dafny.Seq.of((_this).f12);
          let _1207_v22;
          _1207_v22 = _dafny.Seq.of(_this.f11, new BigNumber((_1206_v21).length));
          let _1208_v23;
          _1208_v23 = _dafny.Set.fromElements((_module.D6.create_DC14(_1206_v21)).dtor_cf19);
          _1208_v23 = _1208_v23;
          _1181_v0 = (_this.f11).isEqualTo(_1186_i0);
        } else {
          let _index253 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length));
          (_1187_v5)[_index253] = !(!(!(!((_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]))));
          let _1209_v24;
          let _nw246 = Array((new BigNumber(28)).toNumber()).fill([]);
          _1209_v24 = _nw246;
          let _1210_v25;
          let _nw247 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1210_v25 = _nw247;
          let _1211_v26;
          _1211_v26 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_1210_v25, _1210_v25, _1210_v25)).cardinality()));
          let _1212_v27;
          _1212_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1181_v0,_1183_v2);
          let _1213_v28;
          _1213_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v24,(_this).fm5(new BigNumber((_dafny.Seq.update(_1211_v26, _module.__default.safeIndex(_this.f11, new BigNumber((_1211_v26).length)), new BigNumber(755))).length), !(_1181_v0), new BigNumber((_dafny.Set.fromElements(!((_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]), _1181_v0)).length), new BigNumber((_1212_v27).length), globalState));
          _1213_v28 = (_1213_v28).update((((_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]) ? (_1209_v24) : (_1209_v24)), (_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]);
          let _1214_v29;
          _1214_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))]);
          let _1215_v30;
          _1215_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber((_1189_v7).length));
          let _1216_v31;
          let _nw248 = Array((new BigNumber(4)).toNumber());
          _nw248[(_dafny.ZERO).toNumber()] = _this.f11;
          _nw248[(_dafny.ONE).toNumber()] = new BigNumber((_1215_v30).length);
          _nw248[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw248[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_1186_i0, (_this).fm4(_1211_v26, globalState));
          _1216_v31 = _nw248;
          let _rhs240 = _1214_v29;
          let _rhs241 = _1216_v31;
          let _rhs242 = (new BigNumber(154)).isLessThan(new BigNumber(437));
          let _rhs243 = (new BigNumber(-847)).minus((new BigNumber((_1183_v2).length)).minus((_this).f12));
          let _lhs130 = _this;
          _1214_v29 = _rhs240;
          _1216_v31 = _rhs241;
          _1181_v0 = _rhs242;
          _lhs130.f11 = _rhs243;
          let _index254 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_1216_v31).length));
          (_1216_v31)[_index254] = _module.__default.safeModuloInt((_this).f12, _this.f11);
          let _index255 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_1216_v31).length));
          let _rhs244 = _1210_v25;
          let _rhs245 = (((_this.f11).isLessThanOrEqualTo((_this).f12)) ? (_1216_v31) : (_1216_v31));
          let _rhs246 = (new BigNumber(869)).minus(_1186_i0);
          let _rhs247 = (_1187_v5)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1187_v5).length))];
          let _lhs131 = _1216_v31;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_1216_v31).length));
          _1210_v25 = _rhs244;
          _1216_v31 = _rhs245;
          _lhs131[_lhs132] = _rhs246;
          _1181_v0 = _rhs247;
          (_this).f11 = _this.f11;
        }
        let _1217_v32;
        let _init43 = function (_1218_i3) {
          return (_1218_i3).multipliedBy((_this).f12);
        };
        let _nw249 = Array((new BigNumber(20)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw249.length); _i0_43++) {
          _nw249[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1217_v32 = _nw249;
        let _1219_v33;
        _1219_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1183_v2,_1217_v32);
        let _1220_v34;
        _1220_v34 = _dafny.Seq.of(new BigNumber((_1219_v33).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(56)), function (_1221_i4) {
          return (_this).f5;
        })).length));
        let _1222_v35;
        _1222_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_module.__default.safeDivisionInt((_dafny.ZERO).minus((_1220_v34)[_module.__default.safeIndex(new BigNumber((_1189_v7).length), new BigNumber((_1220_v34).length))]), new BigNumber(537)));
        _1222_v35 = (_1222_v35).update(_this.f11, new BigNumber(916));
        let _1223_v36;
        let _nw250 = new _module.C1();
        _nw250.__ctor((_this).f4, (_this).f5);
        _1223_v36 = _nw250;
        _1223_v36 = _1223_v36;
      }
      let _1224_v37;
      let _nw251 = Array((new BigNumber(27)).toNumber());
      _nw251[(_dafny.ZERO).toNumber()] = (_this).f5;
      _nw251[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
      _nw251[(new BigNumber(2)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(3)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(4)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(5)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
      _nw251[(new BigNumber(7)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(8)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(9)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(10)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(11)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(12)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(13)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(14)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(15)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_this.f11, (_this).f12, globalState);
      _nw251[(new BigNumber(17)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(18)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(19)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(20)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(21)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(22)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
      _nw251[(new BigNumber(23)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(24)).toNumber()] = (_this).f5;
      _nw251[(new BigNumber(25)).toNumber()] = _module.__default.fm0(new BigNumber(294), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(497)), function (_1225_i5) {
        return (_this).f5;
      })).length))).length), globalState);
      _nw251[(new BigNumber(26)).toNumber()] = (_this).f5;
      _1224_v37 = _nw251;
      let _index256 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1224_v37).length));
      (_1224_v37)[_index256] = (_this).f5;
      let _index257 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1224_v37).length));
      (_1224_v37)[_index257] = (_this).f5;
      let _1226_v38;
      _1226_v38 = _dafny.Seq.of((_this).f12);
      let _1227_v39;
      _1227_v39 = _module.D6.create_DC14(_1226_v38);
      r0 = _dafny.Seq.Concat(_module.__default.fm31((_this).f12, (_this).f12, globalState), (_1227_v39).dtor_cf19);
      r1 = (_this).fm4(_dafny.Seq.of(_this.f11), globalState);
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _1228_v0;
      let _init44 = function (_1229_i1) {
        return false;
      };
      let _nw252 = Array((new BigNumber(23)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw252.length); _i0_44++) {
        _nw252[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1228_v0 = _nw252;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1228_v0).length))) {
        let _1230_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1230_i0)) && ((_1230_i0).isLessThan(new BigNumber((_1228_v0).length))))) {
          (_1228_v0)[(_1230_i0)] = false;
        }
      }
      let _1231_v1;
      let _init45 = function (_1232_i3) {
        return _dafny.Seq.UnicodeFromString("q");
      };
      let _nw253 = Array((new BigNumber(4)).toNumber());
      for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw253.length); _i0_45++) {
        _nw253[_i0_45] = _init45(new BigNumber(_i0_45));
      }
      _1231_v1 = _nw253;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1231_v1).length))) {
        let _1233_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1233_i2)) && ((_1233_i2).isLessThan(new BigNumber((_1231_v1).length))))) {
          (_1231_v1)[(_1233_i2)] = _dafny.Seq.UnicodeFromString("baooybn");
        }
      }
      let _1234_v2;
      let _nw254 = new _module.C0();
      _nw254.__ctor();
      _1234_v2 = _nw254;
      let _1235_v3;
      _1235_v3 = _dafny.Seq.UnicodeFromString("ljvlju");
      _1235_v3 = _1235_v3;
      _1228_v0 = _1228_v0;
      let _1236_v4;
      let _init46 = function (_1237_i4) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_dafny.ZERO).minus((_this).f12))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f12,_this.f11));
      };
      let _nw255 = Array((new BigNumber(16)).toNumber());
      for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw255.length); _i0_46++) {
        _nw255[_i0_46] = _init46(new BigNumber(_i0_46));
      }
      _1236_v4 = _nw255;
      let _index258 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1236_v4).length));
      (_1236_v4)[_index258] = _dafny.Map.Empty.slice().updateUnsafe(_this.f11,new BigNumber(-171));
      let _1238_v5;
      _1238_v5 = _dafny.Seq.of(_this.f11, new BigNumber(55));
      let _1239_v6;
      _1239_v6 = true;
      let _1240_v7;
      _1240_v7 = _dafny.Set.fromElements(_1239_v6, _1239_v6);
      let _1241_v8;
      _1241_v8 = _dafny.MultiSet.fromElements(new BigNumber((_1240_v7).length), (_this).f12, _this.f11);
      let _1242_v9;
      _1242_v9 = _dafny.MultiSet.fromElements(_1239_v6);
      let _1243_v10;
      _1243_v10 = _dafny.MultiSet.fromElements((_this).fm4(_dafny.Seq.of((_this).f12), globalState), (_1238_v5)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_1238_v5).length))], (((_1241_v8).contains(new BigNumber((_1242_v9).cardinality()))) ? ((_1241_v8).get(new BigNumber((_1242_v9).cardinality()))) : (new BigNumber((_1240_v7).length))));
      let _1244_v11;
      _1244_v11 = _dafny.Map.Empty.slice().updateUnsafe((((_1241_v8).contains((_dafny.ZERO).minus(_this.f11))) ? ((_1241_v8).get((_dafny.ZERO).minus(_this.f11))) : (_this.f11)),new BigNumber(369));
      let _index259 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1236_v4).length));
      let _rhs248 = new BigNumber((_1235_v3).length);
      let _rhs249 = ((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f12)).Merge(_1244_v11)).update((_this).f12, new BigNumber(267));
      let _lhs133 = _this;
      let _lhs134 = _1236_v4;
      let _lhs135 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1236_v4).length));
      _lhs133.f11 = _rhs248;
      _lhs134[_lhs135] = _rhs249;
      let _1245_v12;
      let _nw256 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _1245_v12 = _nw256;
      let _1246_v13;
      _1246_v13 = _module.D11.create_DC25(_1239_v6, _this.f11, _1245_v12, (_this).f12);
      r0 = (_1246_v13).dtor_cf40;
      return r0;
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
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
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.ZERO;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = _dafny.ZERO;
      let _1247_v0;
      let _init47 = ((_1248_p2) => function (_1249_i0) {
        return _1248_p2;
      })(p2);
      let _nw257 = Array((new BigNumber(11)).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw257.length); _i0_47++) {
        _nw257[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _1247_v0 = _nw257;
      let _1250_v1;
      _1250_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(260),_1247_v0);
      _1250_v1 = ((p2) ? (_1250_v1) : (_1250_v1));
      let _1251_v2;
      _1251_v2 = _dafny.MultiSet.fromElements(p0);
      let _1252_v3;
      _1252_v3 = _dafny.Seq.of(_1251_v2, _1251_v2);
      r1 = (new BigNumber(((_1252_v3)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_1252_v3).length))]).cardinality())).multipliedBy(p0);
      let _1253_v4;
      _1253_v4 = _dafny.Seq.UnicodeFromString("lcyi");
      r0 = _1253_v4;
      let _1254_v5;
      let _nw258 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _1254_v5 = _nw258;
      let _index260 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
      (_1254_v5)[_index260] = p0;
      let _index261 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
      let _rhs250 = _dafny.Seq.Concat(_1253_v4, ((p2) ? (_1253_v4) : (_dafny.Seq.UnicodeFromString("qbpbcw"))));
      let _rhs251 = p0;
      let _rhs252 = (p0).plus(p0);
      let _lhs136 = _1254_v5;
      let _lhs137 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
      _1253_v4 = _rhs250;
      _lhs136[_lhs137] = _rhs251;
      r1 = _rhs252;
      let _1255_v6;
      _1255_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(p0));
      let _hi10 = (((_1255_v6).contains(!(p2))) ? ((_1255_v6).get(!(p2))) : (p0));
      for (let _1256_i1 = p0; _1256_i1.isLessThan(_hi10); _1256_i1 = _1256_i1.plus(_dafny.ONE)) {
        let _1257_v7;
        _1257_v7 = false;
        _1257_v7 = !(_1257_v7) || (p2);
        let _1258_v8;
        let _init48 = ((_1259_i1, _1260_p2, _1261_v5, _1262_p0) => function (_1263_i2) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_1259_i1, new BigNumber((_dafny.Seq.of(_1260_p2)).length), new BigNumber(-80), (_1261_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1261_v5).length))], _1262_p0), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1261_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1261_v5).length))],(_1261_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1261_v5).length))])).length)));
        })(_1256_i1, p2, _1254_v5, p0);
        let _nw259 = Array((new BigNumber(29)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw259.length); _i0_48++) {
          _nw259[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1258_v8 = _nw259;
        let _1264_v9;
        _1264_v9 = _dafny.Seq.of(_1256_i1);
        let _index262 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1258_v8).length));
        (_1258_v8)[_index262] = _dafny.Seq.Concat(_1264_v9, _1264_v9);
        let _index263 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1258_v8).length));
        (_1258_v8)[_index263] = _1264_v9;
        let _1265_v10;
        _1265_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1256_i1, _1256_i1),(_dafny.ZERO).minus(_1256_i1));
        let _1266_v11;
        _1266_v11 = _module.D6.create_DC14(_dafny.Seq.of((_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))]));
        r3 = new BigNumber((((_1265_v10).update((_1266_v11).dtor_cf19, p0)).Merge(_1265_v10)).length);
        let _1267_v12;
        _1267_v12 = _dafny.Set.fromElements(new BigNumber(846), _module.__default.fm14(globalState));
        if (((_1267_v12).IsDisjointFrom(_1267_v12)) || (!(p2))) {
          let _index264 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
          let _rhs253 = _1253_v4;
          let _rhs254 = _1256_i1;
          let _lhs138 = _1254_v5;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
          _1253_v4 = _rhs253;
          _lhs138[_lhs139] = _rhs254;
          let _1268_v13;
          let _nw260 = Array((new BigNumber(29)).toNumber()).fill(_module.D5.Default());
          _1268_v13 = _nw260;
          let _1269_v14;
          _1269_v14 = _module.D5.create_DC12();
          let _index265 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1268_v13).length));
          (_1268_v13)[_index265] = _1269_v14;
          let _index266 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1268_v13).length));
          (_1268_v13)[_index266] = _module.D5.create_DC12();
          let _1270_v15;
          _1270_v15 = _dafny.Seq.of(p2);
          let _1271_v16;
          _1271_v16 = _module.D2.create_DC3(_1247_v0);
          let _1272_v17;
          _1272_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1271_v16,p0);
          let _1273_v18;
          _1273_v18 = _dafny.Set.fromElements(_1255_v6, _1255_v6);
          let _index267 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
          let _rhs255 = !((_1256_i1).isLessThanOrEqualTo(new BigNumber((_1270_v15).length)));
          let _rhs256 = (_dafny.ZERO).minus(p0);
          let _rhs257 = new BigNumber((_1272_v17).length);
          let _rhs258 = (p0).multipliedBy(_module.__default.fm14(globalState));
          let _rhs259 = ((((_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))]).isLessThan((_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))])) ? (!(_1273_v18).equals(function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of (_1273_v18).Elements) {
              let _1274_v19 = _compr_45;
              if ((_1273_v18).contains(_1274_v19)) {
                _coll45.add(_1274_v19);
              }
            }
            return _coll45;
          }())) : (false));
          let _lhs140 = _1254_v5;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
          _1257_v7 = _rhs255;
          _lhs140[_lhs141] = _rhs256;
          r1 = _rhs257;
          r1 = _rhs258;
          _1257_v7 = _rhs259;
          let _index268 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1247_v0).length));
          (_1247_v0)[_index268] = _1257_v7;
          let _index269 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1247_v0).length));
          (_1247_v0)[_index269] = _1257_v7;
          let _1275_v20;
          _1275_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1257_v7);
          let _1276_v21;
          _1276_v21 = _dafny.Set.fromElements((((_1275_v20).contains(p0)) ? ((_1275_v20).get(p0)) : ((_1247_v0)[_module.__default.safeIndex(new BigNumber(135), new BigNumber((_1247_v0).length))])), _1257_v7);
          let _1277_v22;
          let _nw261 = Array((new BigNumber(11)).toNumber());
          _nw261[(_dafny.ZERO).toNumber()] = _module.__default.fm26((_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))], p0, globalState);
          _nw261[(_dafny.ONE).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(2)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(3)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(4)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(5)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(6)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(7)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(8)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(9)).toNumber()] = _1276_v21;
          _nw261[(new BigNumber(10)).toNumber()] = _1276_v21;
          _1277_v22 = _nw261;
          let _1278_v23;
          _1278_v23 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1279_v24;
          let _nw262 = new _module.C2();
          _nw262.__ctor(_1277_v22, _1278_v23);
          _1279_v24 = _nw262;
        } else {
          let _1280_v25;
          _1280_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1256_i1,new BigNumber((_dafny.Seq.of(_1257_v7, _1257_v7, true, !(_1257_v7))).length));
          r3 = _module.__default.safeModuloInt(new BigNumber((_1253_v4).length), (((_1280_v25).contains(_1256_i1)) ? ((_1280_v25).get(_1256_i1)) : ((_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))])));
          let _1281_v26;
          _1281_v26 = _dafny.Set.fromElements(_module.__default.fm16(p2, _1256_i1, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_1282_i3) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), ((false) ? (_dafny.Seq.UnicodeFromString("okftu")) : (_1253_v4)));
          _1281_v26 = _1281_v26;
          _1267_v12 = (function () {
            let _coll46 = new _dafny.Set();
            for (const _compr_46 of (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-187)))).Elements) {
              let _1283_v27 = _compr_46;
              if ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-187)))).contains(_1283_v27)) {
                _coll46.add(_module.__default.safeModuloInt(_1283_v27, new BigNumber(((_module.D16.create_DC37(function () {
  let _coll47 = new _dafny.Set();
  for (const _compr_47 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_1284_i4) {
    return new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(940))).cardinality()))).length);
  })).Elements) {
    let _1285_v28 = _compr_47;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_1284_i4) {
      return new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(940))).cardinality()))).length);
    }), _1285_v28)) {
      _coll47.add((_1285_v28).minus(new BigNumber(-160)));
    }
  }
  return _coll47;
}())).dtor_cf56).length)));
              }
            }
            return _coll46;
          }()).Difference(_1267_v12);
          let _1286_v29;
          _1286_v29 = _module.D6.create_DC15(p2, _1257_v7, _1257_v7);
          let _1287_v30;
          _1287_v30 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1286_v29,new BigNumber((_1281_v26).length)));
          let _1288_v31;
          _1288_v31 = _module.D8.create_DC19(new BigNumber(((_1287_v30).Intersect(_1287_v30)).length), p2);
          let _1289_v32;
          _1289_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1254_v5,_1253_v4);
          let _1290_v33;
          _1290_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1257_v7,_1253_v4);
          let _rhs260 = ((_1251_v2).Union(_1251_v2)).Union(_1251_v2);
          let _rhs261 = _1288_v31;
          let _rhs262 = (p0).multipliedBy(new BigNumber(((((_1289_v32).contains(_1254_v5)) ? ((_1289_v32).get(_1254_v5)) : ((((_1290_v33).contains(p2)) ? ((_1290_v33).get(p2)) : (_dafny.Seq.UnicodeFromString("evqfa")))))).length));
          _1251_v2 = _rhs260;
          _1288_v31 = _rhs261;
          r3 = _rhs262;
          let _1291_v34;
          let _nw263 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _1291_v34 = _nw263;
          let _1292_v35;
          _1292_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("itrexs")).length)),new _dafny.CodePoint('u'.codePointAt(0)));
          let _index270 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1291_v34).length));
          (_1291_v34)[_index270] = _1292_v35;
          let _index271 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1291_v34).length));
          (_1291_v34)[_index271] = _1292_v35;
        }
      }
      let _index272 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length));
      (_1254_v5)[_index272] = (_1254_v5)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((_1254_v5).length))];
      let _1293_v36;
      _1293_v36 = _dafny.MultiSet.fromElements(!(p2));
      r0 = _module.__default.fm16(p2, new BigNumber((_1293_v36).cardinality()), globalState);
      let _1294_v37;
      _1294_v37 = _dafny.Seq.of(p2, p2);
      r1 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1294_v37).length), new BigNumber(945)))).minus(new BigNumber(442));
      let _1295_v38;
      _1295_v38 = new _dafny.CodePoint('s'.codePointAt(0));
      r2 = ((!_dafny.areEqual(_1253_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(598)), function (_1296_i5) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }))) ? (_1295_v38) : (_1295_v38));
      r3 = _module.__default.fm14(globalState);
      return [r0, r1, r2, r3];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f10 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10) {
      let _this = this;
      (_this)._f10 = f10;
      return;
    }
    m9(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = [];
      let _1297_v0;
      _1297_v0 = new BigNumber(266);
      let _hi11 = _1297_v0;
      for (let _1298_i0 = _1297_v0; _1298_i0.isLessThan(_hi11); _1298_i0 = _1298_i0.plus(_dafny.ONE)) {
        let _1299_v1;
        _1299_v1 = _dafny.Seq.of(p1);
        r2 = (_1299_v1)[_module.__default.safeIndex(_module.__default.fm13(p1, true, p1, globalState), new BigNumber((_1299_v1).length))];
        let _1300_v3;
        let _init49 = ((_1301_i0) => function (_1302_i1) {
          return (_1302_i1).multipliedBy(_1301_i0);
        })(_1298_i0);
        let _nw264 = Array((new BigNumber(17)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw264.length); _i0_49++) {
          _nw264[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1300_v3 = _nw264;
        let _1303_v4;
        _1303_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_dafny.ZERO).minus(_1297_v0), _1298_i0, (_dafny.ZERO).minus(_1298_i0)),_1300_v3);
        _1297_v0 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll48 = new _dafny.Set();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(10), new BigNumber(-345))) {
            let _1304_v2 = _compr_48;
            if (((new BigNumber(10)).isLessThanOrEqualTo(_1304_v2)) && ((_1304_v2).isLessThan(new BigNumber(-345)))) {
              _coll48.add((_1304_v2).plus(_1298_i0));
            }
          }
          return _coll48;
        }(),_1300_v3)).Merge(_1303_v4)).length);
        _1297_v0 = _1298_i0;
        r1 = (_1298_i0).isLessThanOrEqualTo(new BigNumber(990));
      }
      let _1305_v5;
      _1305_v5 = _dafny.Seq.UnicodeFromString("pcoyirdm");
      let _1306_v6;
      _1306_v6 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1307_v7;
      _1307_v7 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_1305_v5, _module.__default.safeIndex(_1297_v0, new BigNumber((_1305_v5).length)), _1306_v6), _1305_v5);
      let _1308_v8;
      _1308_v8 = _module.D5.create_DC13(_dafny.Seq.UnicodeFromString("tbin"), _1297_v0, p1);
      let _hi12 = _1297_v0;
      for (let _1309_i2 = (((_1307_v7).contains((_1308_v8).dtor_cf16)) ? ((_1307_v7).get((_1308_v8).dtor_cf16)) : (_1297_v0)); _1309_i2.isLessThan(_hi12); _1309_i2 = _1309_i2.plus(_dafny.ONE)) {
        let _1310_v9;
        let _nw265 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1310_v9 = _nw265;
        let _index273 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length));
        (_1310_v9)[_index273] = p1;
        let _index274 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length));
        (_1310_v9)[_index274] = p1;
        r1 = !((_module.D7.create_DC17((_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))], p1, p1, (_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))])).dtor_cf27);
        let _1311_v10;
        let _init50 = ((_1312_v0) => function (_1313_i3) {
          return _module.__default.safeDivisionInt(_1313_i3, _1312_v0);
        })(_1297_v0);
        let _nw266 = Array((new BigNumber(16)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw266.length); _i0_50++) {
          _nw266[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1311_v10 = _nw266;
        let _index275 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1311_v10).length));
        (_1311_v10)[_index275] = _1297_v0;
        let _1314_v11;
        _1314_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1305_v5,_1310_v9);
        let _1315_v12;
        _1315_v12 = _dafny.Seq.of(_1310_v9, (((_1314_v11).contains(_dafny.Seq.UnicodeFromString("woupfmq"))) ? ((_1314_v11).get(_dafny.Seq.UnicodeFromString("woupfmq"))) : (_1310_v9)), _1310_v9);
        let _1316_v13;
        let _nw267 = Array((new BigNumber(16)).toNumber());
        _nw267[(_dafny.ZERO).toNumber()] = _1310_v9;
        _nw267[(_dafny.ONE).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(2)).toNumber()] = (((_1314_v11).contains(_1305_v5)) ? ((_1314_v11).get(_1305_v5)) : (_1310_v9));
        _nw267[(new BigNumber(3)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(4)).toNumber()] = (_1315_v12)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_1315_v12).length))];
        _nw267[(new BigNumber(5)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(6)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(7)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(8)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(9)).toNumber()] = (((_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))]) ? (_1310_v9) : (_1310_v9));
        _nw267[(new BigNumber(10)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(11)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(12)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(13)).toNumber()] = (_1315_v12)[_module.__default.safeIndex(_1309_i2, new BigNumber((_1315_v12).length))];
        _nw267[(new BigNumber(14)).toNumber()] = _1310_v9;
        _nw267[(new BigNumber(15)).toNumber()] = _1310_v9;
        _1316_v13 = _nw267;
        let _1317_v14;
        _1317_v14 = _dafny.MultiSet.fromElements((_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))]);
        let _1318_v15;
        _1318_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1317_v14,_1297_v0);
        let _1319_v16;
        _1319_v16 = _dafny.MultiSet.fromElements(_module.__default.fm13(p1, (_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))], p1, globalState), new BigNumber((_1318_v15).length));
        let _1320_v17;
        _1320_v17 = _dafny.Seq.of(((((_this).f10).contains(new BigNumber((_1319_v16).cardinality()))) ? (((_this).f10).get(new BigNumber((_1319_v16).cardinality()))) : (_1309_i2)), new BigNumber((_dafny.Seq.of(new BigNumber(497))).length));
        let _1321_v18;
        _1321_v18 = _dafny.Seq.of(_1320_v17, _1320_v17);
        let _1322_v20;
        _1322_v20 = _module.D2.create_DC5(_1305_v5);
        let _1323_v21;
        _1323_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1320_v17).length),_1322_v20);
        let _index276 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1311_v10).length));
        let _rhs263 = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-173)), ((_1324_i2) => function (_1325_i4) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), ((_1326_i2) => function (_1327_i5) {
            return _1326_i2;
          })(_1324_i2));
        })(_1309_i2)), ((!((_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))])) ? (_1321_v18) : (_1321_v18)));
        let _rhs264 = _1297_v0;
        let _rhs265 = p1;
        let _rhs266 = _module.__default.fm1(new BigNumber(((function () {
          let _coll49 = new _dafny.Map();
          for (const _compr_49 of (_1319_v16).Elements) {
            let _1328_v19 = _compr_49;
            if ((_1319_v16).contains(_1328_v19)) {
              _coll49.push([_module.__default.safeModuloInt(_1328_v19, _1297_v0),_module.D2.create_DC5(_1305_v5)]);
            }
          }
          return _coll49;
        }()).Merge(_1323_v21)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1310_v9)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_1310_v9).length))],_1297_v0)).length), p1, p1, globalState);
        let _rhs267 = _1316_v13;
        let _lhs142 = _1311_v10;
        let _lhs143 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1311_v10).length));
        r1 = _rhs263;
        _lhs142[_lhs143] = _rhs264;
        r2 = _rhs265;
        r1 = _rhs266;
        _1316_v13 = _rhs267;
        let _1329_v22;
        _1329_v22 = _dafny.Set.fromElements(_1309_i2);
        _1329_v22 = _1329_v22;
      }
      let _1330_v23;
      _1330_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1305_v5,p1);
      _1330_v23 = (_1330_v23).update(_1305_v5, p1);
      let _1331_v24;
      _1331_v24 = _dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)));
      let _1332_v25;
      _1332_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1331_v24).Difference(_1331_v24),_1297_v0);
      let _rhs268 = _module.__default.safeModuloInt(_1297_v0, _1297_v0);
      let _rhs269 = p1;
      let _rhs270 = p1;
      let _rhs271 = (_1297_v0).minus(_1297_v0);
      let _rhs272 = _1332_v25;
      _1297_v0 = _rhs268;
      r2 = _rhs269;
      r2 = _rhs270;
      _1297_v0 = _rhs271;
      _1332_v25 = _rhs272;
      let _1333_v26;
      let _init51 = ((_1334_p1) => function (_1335_i6) {
        return _dafny.Set.fromElements(_1334_p1, true);
      })(p1);
      let _nw268 = Array((new BigNumber(19)).toNumber());
      for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw268.length); _i0_51++) {
        _nw268[_i0_51] = _init51(new BigNumber(_i0_51));
      }
      _1333_v26 = _nw268;
      let _1336_v27;
      let _nw269 = new _module.C1();
      _nw269.__ctor(_1333_v26, _module.__default.fm0(_1297_v0, _1297_v0, globalState));
      _1336_v27 = _nw269;
      let _1337_v28;
      _1337_v28 = _dafny.MultiSet.fromElements(p1);
      let _1338_i7;
      _1338_i7 = _dafny.ZERO;
      L6: {
        while ((_1337_v28).IsDisjointFrom(_1337_v28)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1338_i7)) {
              break L6;
            }
            _1338_i7 = (_1338_i7).plus(_dafny.ONE);
            let _1339_v29;
            let _1340_v30;
            let _1341_v31;
            let _1342_v32;
            let _out13;
            let _out14;
            let _out15;
            let _out16;
            let _outcollector4 = (_1336_v27).m14(_module.__default.safeDivisionInt(_1297_v0, _1297_v0), _1297_v0, globalState);
            _out13 = _outcollector4[0];
            _out14 = _outcollector4[1];
            _out15 = _outcollector4[2];
            _out16 = _outcollector4[3];
            _1339_v29 = _out13;
            _1340_v30 = _out14;
            _1341_v31 = _out15;
            _1342_v32 = _out16;
            let _1343_v33;
            _1343_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1342_v32,_1342_v32);
            let _1344_v34;
            _1344_v34 = _dafny.Seq.of((_1343_v33).Merge(_1343_v33), (_1343_v33).Merge(_1343_v33), _module.__default.fm2(new BigNumber(796), _1342_v32, _1342_v32, _1297_v0, globalState));
            _1344_v34 = _dafny.Seq.Concat(_dafny.Seq.Concat(p0, p0), _dafny.Seq.of((_1343_v33).update(true, _1342_v32), _1343_v33));
            let _1345_v35;
            let _nw270 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _1345_v35 = _nw270;
            let _index277 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length));
            (_1345_v35)[_index277] = ((_1342_v32) ? (_1339_v29) : (new BigNumber(5)));
            let _index278 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length));
            (_1345_v35)[_index278] = _module.__default.safeModuloInt((new BigNumber((_1305_v5).length)).minus(_1339_v29), _1339_v29);
            let _1346_v36;
            _1346_v36 = _dafny.Seq.of(p1);
            let _1347_v37;
            _1347_v37 = _module.D11.create_DC25(p1, _1339_v29, _1345_v35, new BigNumber((_1346_v36).length));
            let _source13 = _1347_v37;
            if (_source13.is_DC25) {
              let _1348___mcc_h0 = (_source13).cf38;
              let _1349___mcc_h1 = (_source13).cf39;
              let _1350___mcc_h2 = (_source13).cf40;
              let _1351___mcc_h3 = (_source13).cf41;
              let _1352_cf41 = _1351___mcc_h3;
              let _1353_cf40 = _1350___mcc_h2;
              let _1354_cf39 = _1349___mcc_h1;
              let _1355_cf38 = _1348___mcc_h0;
              let _1356_v38;
              _1356_v38 = _dafny.Seq.of(_1297_v0, _1297_v0);
              let _1357_v39;
              _1357_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1343_v33,(_1336_v27).fm5((_1336_v27).fm4(_1356_v38, globalState), _1355_cf38, _1352_cf41, _1297_v0, globalState));
              _1357_v39 = _1357_v39;
              let _index279 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length));
              (_1345_v35)[_index279] = ((_1345_v35)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length))]).multipliedBy(_1354_cf39);
              let _1358_v40;
              let _nw271 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
              _1358_v40 = _nw271;
              let _index280 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1358_v40).length));
              (_1358_v40)[_index280] = function () {
                let _coll50 = new _dafny.Map();
                for (const _compr_50 of (_1356_v38).Elements) {
                  let _1359_v41 = _compr_50;
                  if (_dafny.Seq.contains(_1356_v38, _1359_v41)) {
                    _coll50.push([(_1359_v41).multipliedBy(_1297_v0),new BigNumber((_dafny.Seq.update(_1305_v5, _module.__default.safeIndex(_1352_cf41, new BigNumber((_1305_v5).length)), _1306_v6)).length)]);
                  }
                }
                return _coll50;
              }();
              let _1360_v42;
              let _nw272 = Array((new BigNumber(8)).toNumber()).fill(false);
              _1360_v42 = _nw272;
              let _index281 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1358_v40).length));
              let _rhs273 = _1360_v42;
              let _rhs274 = (_this).f10;
              let _rhs275 = _1355_cf38;
              let _lhs144 = globalState;
              let _lhs145 = _1358_v40;
              let _lhs146 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1358_v40).length));
              _lhs144.f3 = _rhs273;
              _lhs145[_lhs146] = _rhs274;
              _1342_v32 = _rhs275;
              _1297_v0 = _1339_v29;
            } else if (_source13.is_DC24) {
              let _1361___mcc_h4 = (_source13).cf37;
              let _1362_cf37 = _1361___mcc_h4;
              r2 = false;
              r2 = ((p1) ? (_1342_v32) : ((_1342_v32) === (!(_1342_v32))));
              let _1363_v43;
              _1363_v43 = _dafny.Set.fromElements(new BigNumber(750));
              _1343_v33 = (_1343_v33).update(_1342_v32, (_1363_v43).IsProperSubsetOf(_1363_v43));
              let _1364_v44;
              _1364_v44 = _dafny.Seq.of(_dafny.Seq.Concat(_module.__default.fm31(_1297_v0, _1297_v0, globalState), _module.__default.fm31(new BigNumber((_dafny.Seq.UnicodeFromString("axjc")).length), _1340_v30, globalState)));
              let _1365_v45;
              _1365_v45 = _dafny.Seq.of((_1345_v35)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length))], _1340_v30);
              _1339_v29 = new BigNumber((_dafny.Seq.update(_1364_v44, _module.__default.safeIndex((_1345_v35)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length))], new BigNumber((_1364_v44).length)), _1365_v45)).length);
            } else {
              let _1366___mcc_h5 = (_source13).cf42;
              let _1367_cf42 = _1366___mcc_h5;
              let _index282 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1345_v35).length));
              (_1345_v35)[_index282] = _1339_v29;
              let _1368_v46;
              _1368_v46 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_1305_v5, _1306_v6),_1343_v33);
              let _1369_v47;
              _1369_v47 = _dafny.Set.fromElements(false);
              let _1370_v48;
              _1370_v48 = _dafny.Seq.of(_1345_v35, _1345_v35);
              _1368_v46 = (_1368_v46).update((_1369_v47).IsProperSubsetOf(_1369_v47), (_module.__default.fm2(_1341_v31, p1, _1342_v32, (_dafny.ZERO).minus(new BigNumber((_1370_v48).length)), globalState)).update(_1342_v32, p1));
              _1342_v32 = ((_1346_v36)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1305_v5).length))).length)), new BigNumber((_1346_v36).length))]) || (_1342_v32);
              let _1371_v49;
              _1371_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p1,_1341_v31),false);
              _1342_v32 = (_1340_v30).isLessThan(new BigNumber((_1371_v49).length));
            }
          }
        }
      }
      let _1372_v51;
      _1372_v51 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1297_v0);
      let _1373_v52;
      _1373_v52 = _module.D2.create_DC6(new BigNumber((_1372_v51).length));
      let _1374_v53;
      _1374_v53 = _dafny.Seq.of((_1373_v52).dtor_cf7);
      let _1375_v54;
      _1375_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v53,new BigNumber((_dafny.Seq.of(_1297_v0)).length));
      r0 = function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of (((p1) ? ((_1375_v54).update(_dafny.Seq.of(_1297_v0), _1297_v0)) : (_1375_v54))).Keys.Elements) {
          let _1376_v50 = _compr_51;
          if ((((p1) ? ((_1375_v54).update(_dafny.Seq.of(_1297_v0), _1297_v0)) : (_1375_v54))).contains(_1376_v50)) {
            _coll51.push([_1376_v50,!(p1)]);
          }
        }
        return _coll51;
      }();
      let _1377_v55;
      let _init52 = ((_1378_p1) => function (_1379_i8) {
        return _1378_p1;
      })(p1);
      let _nw273 = Array((new BigNumber(20)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw273.length); _i0_52++) {
        _nw273[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _1377_v55 = _nw273;
      let _1380_v56;
      _1380_v56 = _dafny.Set.fromElements(_1377_v55);
      r1 = (_1380_v56).IsSubsetOf(_1380_v56);
      r2 = p1;
      let _1381_v57;
      let _init53 = ((_1382_v28) => function (_1383_i9) {
        return _1382_v28;
      })(_1337_v28);
      let _nw274 = Array((new BigNumber(17)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw274.length); _i0_53++) {
        _nw274[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _1381_v57 = _nw274;
      r3 = _1381_v57;
      return [r0, r1, r2, r3];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return !(((_module.D2.create_DC4(false)).dtor_cf5) === (true)) || (true);
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(913);
    };
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _hi13 = p1;
      for (let _1384_i0 = p1; _1384_i0.isLessThan(_hi13); _1384_i0 = _1384_i0.plus(_dafny.ONE)) {
        let _1385_v0;
        _1385_v0 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1386_v1;
        _1386_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,p2);
        let _1387_v2;
        let _init54 = ((_1388_p2) => function (_1389_i1) {
          return _1388_p2;
        })(p2);
        let _nw275 = Array((new BigNumber(9)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw275.length); _i0_54++) {
          _nw275[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1387_v2 = _nw275;
        let _1390_v3;
        _1390_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1387_v2,p2);
        _1386_v1 = (_1386_v1).update(_1385_v0, _module.__default.fm1(_1384_i0, _1384_i0, (((_1390_v3).contains(_1387_v2)) ? ((_1390_v3).get(_1387_v2)) : (!(true))), p2, globalState));
        let _index283 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_1387_v2).length));
        (_1387_v2)[_index283] = (_this).fm3(_1384_i0, _1384_i0, globalState);
        let _1391_v4;
        _1391_v4 = _module.D1.create_DC2(p1, _1384_i0);
        let _index284 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_1387_v2).length));
        (_1387_v2)[_index284] = ((false) ? (!((_dafny.ZERO).minus(_1384_i0)).isEqualTo(new BigNumber(299))) : (((_1391_v4).dtor_cf2).isLessThanOrEqualTo(p1)));
        let _1392_v5;
        _1392_v5 = new BigNumber(671);
        let _1393_v6;
        _1393_v6 = _dafny.Seq.UnicodeFromString("ph");
        let _1394_v7;
        _1394_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1393_v6,_module.__default.safeModuloInt(new BigNumber(134), p1));
        _1392_v5 = (((_1394_v7).contains(_1393_v6)) ? ((_1394_v7).get(_1393_v6)) : (p1));
        let _1395_v8;
        _1395_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,_1385_v0);
        _1395_v8 = (_1395_v8).update(_1385_v0, _1385_v0);
      }
      if (p2) {
        r0 = p2;
        let _1396_v9;
        let _nw276 = Array((new BigNumber(29)).toNumber()).fill(_module.D2.Default());
        _1396_v9 = _nw276;
        let _1397_v10;
        _1397_v10 = _module.D2.create_DC4(p2);
        let _1398_v11;
        _1398_v11 = _module.D2.create_DC7(_1397_v10);
        let _1399_v12;
        _1399_v12 = _module.D2.create_DC7(_1398_v11);
        let _index285 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_1396_v9).length));
        (_1396_v9)[_index285] = _1399_v12;
        let _index286 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_1396_v9).length));
        (_1396_v9)[_index286] = _module.D2.create_DC7(_1397_v10);
        if (_module.__default.fm1(p1, p1, p2, !((p2) || (p2)), globalState)) {
          let _1400_v13;
          _1400_v13 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1401_v14;
          _1401_v14 = _dafny.Seq.UnicodeFromString("tcfkl");
          let _1402_v15;
          _1402_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1400_v13, _1400_v13),_dafny.Seq.IsProperPrefixOf(_1401_v14, _1401_v14));
          _1402_v15 = _1402_v15;
          let _1403_v16;
          let _init55 = ((_1404_v14, _1405_p2, _1406_p1) => function (_1407_i2) {
            return (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-893), new BigNumber((_1404_v14).length)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1405_p2,_1406_p1)).length), _1406_p1)));
          })(_1401_v14, p2, p1);
          let _nw277 = Array((new BigNumber(6)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw277.length); _i0_55++) {
            _nw277[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _1403_v16 = _nw277;
          let _1408_v17;
          _1408_v17 = _dafny.MultiSet.fromElements(p1);
          let _index287 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1403_v16).length));
          (_1403_v16)[_index287] = (_1408_v17).Difference(_1408_v17);
          let _index288 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_1403_v16).length));
          (_1403_v16)[_index288] = _module.__default.fm11(globalState);
          let _1409_v18;
          _1409_v18 = _dafny.Seq.of(p1);
          (globalState).f2 = _1409_v18;
          r0 = !((_module.D5.create_DC13(_1401_v14, new BigNumber(531), p2)).dtor_cf18);
          r0 = p2;
        } else {
          let _1410_v19;
          _1410_v19 = new BigNumber(520);
          _1410_v19 = new BigNumber(280);
          let _1411_v20;
          _1411_v20 = _dafny.MultiSet.fromElements(p2);
          let _1412_v21;
          _1412_v21 = _module.D1.create_DC1(_1411_v20);
          let _1413_v22;
          _1413_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _1414_v23;
          _1414_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm10(_1412_v21, p2, _1413_v22, globalState),_1411_v20);
          let _1415_v24;
          _1415_v24 = _dafny.Seq.UnicodeFromString("gxjqhtn");
          let _1416_v25;
          _1416_v25 = _dafny.MultiSet.fromElements(_1415_v24);
          let _1417_v26;
          _1417_v26 = _module.D1.create_DC1((((_1414_v23).contains(new BigNumber((_1416_v25).cardinality()))) ? ((_1414_v23).get(new BigNumber((_1416_v25).cardinality()))) : (_1411_v20)));
          _1417_v26 = _1412_v21;
          _1410_v19 = (p1).multipliedBy(p1);
          let _1418_v27;
          _1418_v27 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1419_v28;
          _1419_v28 = _dafny.Seq.of(p1, p1, _1410_v19);
          let _1420_v29;
          let _nw278 = Array((_dafny.ONE).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1419_v28, _1419_v28);
          _1420_v29 = _nw278;
          let _index289 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1420_v29).length));
          (_1420_v29)[_index289] = _dafny.Seq.update(_1419_v28, _module.__default.safeIndex(_1410_v19, new BigNumber((_1419_v28).length)), _1410_v19);
          let _index290 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1420_v29).length));
          let _rhs276 = _1418_v27;
          let _rhs277 = _1410_v19;
          let _rhs278 = _dafny.Seq.Concat(_1419_v28, (_module.D6.create_DC14(_dafny.Seq.of(_1410_v19, _1410_v19))).dtor_cf19);
          let _lhs147 = _1420_v29;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1420_v29).length));
          _1418_v27 = _rhs276;
          _1410_v19 = _rhs277;
          _lhs147[_lhs148] = _rhs278;
          let _1421_v30;
          _1421_v30 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), function (_1422_i3) {
            return (_module.D7.create_DC16(new _dafny.CodePoint('j'.codePointAt(0)))).dtor_cf23;
          })).length)).isEqualTo(_1410_v19),p1);
          let _1423_v31;
          let _init56 = ((_1424_p2) => function (_1425_i4) {
            return _1424_p2;
          })(p2);
          let _nw279 = Array((new BigNumber(5)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw279.length); _i0_56++) {
            _nw279[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _1423_v31 = _nw279;
          let _index291 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1423_v31).length));
          (_1423_v31)[_index291] = p2;
          let _1426_v32;
          _1426_v32 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1410_v19), (_dafny.ZERO).minus(p1), _1410_v19);
          let _1427_v33;
          _1427_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rblvmkb"),(_dafny.ZERO).minus((((_1426_v32).contains(_1410_v19)) ? ((_1426_v32).get(_1410_v19)) : (p1))));
          let _index292 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1423_v31).length));
          let _rhs279 = _1421_v30;
          let _rhs280 = !(!((_1426_v32).IsProperSubsetOf(_1426_v32)));
          let _rhs281 = (((_module.D2.create_DC4(p2)).dtor_cf5) ? (_1418_v27) : ((_1415_v24)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1410_v19,_1410_v19)).length), new BigNumber((_1415_v24).length))]));
          let _rhs282 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mkagjtkuc"),p1);
          let _lhs149 = _1423_v31;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1423_v31).length));
          _1421_v30 = _rhs279;
          _lhs149[_lhs150] = _rhs280;
          _1418_v27 = _rhs281;
          _1427_v33 = _rhs282;
        }
        let _1428_v34;
        _1428_v34 = new BigNumber(879);
        _1428_v34 = _1428_v34;
        let _1429_v35;
        _1429_v35 = _dafny.Seq.UnicodeFromString("vx");
        let _1430_v36;
        let _nw280 = Array((new BigNumber(7)).toNumber());
        _nw280[(_dafny.ZERO).toNumber()] = _1429_v35;
        _nw280[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1429_v35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), function (_1431_i5) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }));
        _nw280[(new BigNumber(2)).toNumber()] = _1429_v35;
        _nw280[(new BigNumber(3)).toNumber()] = ((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(182)), function (_1432_i6) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })) : (_1429_v35));
        _nw280[(new BigNumber(4)).toNumber()] = _1429_v35;
        _nw280[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("kpley");
        _nw280[(new BigNumber(6)).toNumber()] = _module.__default.fm12(globalState);
        _1430_v36 = _nw280;
        _1430_v36 = _1430_v36;
      } else {
        let _1433_v37;
        _1433_v37 = _dafny.Set.fromElements(p2);
        let _1434_v38;
        _1434_v38 = _dafny.MultiSet.fromElements((_1433_v37).IsProperSubsetOf(_1433_v37));
        _1434_v38 = _1434_v38;
        let _1435_v39;
        let _nw281 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1435_v39 = _nw281;
        let _1436_v40;
        _1436_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _index293 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1435_v39).length));
        (_1435_v39)[_index293] = (((_1436_v40).contains(p1)) ? ((_1436_v40).get(p1)) : (p1));
        let _index294 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1435_v39).length));
        (_1435_v39)[_index294] = (p1).multipliedBy(p1);
        let _1437_v41;
        let _nw282 = Array((new BigNumber(28)).toNumber()).fill([]);
        _1437_v41 = _nw282;
        let _1438_v42;
        _1438_v42 = _dafny.Seq.of(_1437_v41);
        let _1439_v43;
        _1439_v43 = _dafny.Seq.of(_1437_v41, (_1438_v42)[_module.__default.safeIndex(p1, new BigNumber((_1438_v42).length))]);
        let _1440_v44;
        let _nw283 = Array((new BigNumber(12)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = _1437_v41;
        _nw283[(_dafny.ONE).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(2)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(3)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(4)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(5)).toNumber()] = ((p2) ? (_1437_v41) : (_1437_v41));
        _nw283[(new BigNumber(6)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(7)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(8)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(9)).toNumber()] = (_1439_v43)[_module.__default.safeIndex((_1435_v39)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1435_v39).length))], new BigNumber((_1439_v43).length))];
        _nw283[(new BigNumber(10)).toNumber()] = _1437_v41;
        _nw283[(new BigNumber(11)).toNumber()] = _1437_v41;
        _1440_v44 = _nw283;
        let _index295 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1440_v44).length));
        (_1440_v44)[_index295] = _1437_v41;
        let _index296 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1440_v44).length));
        (_1440_v44)[_index296] = _1437_v41;
        r0 = p2;
        let _1441_v45;
        _1441_v45 = _dafny.Seq.UnicodeFromString("rsco");
        _1441_v45 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), function (_1442_i7) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        });
      }
      let _1443_v46;
      let _nw284 = new _module.C5();
      _nw284.__ctor();
      _1443_v46 = _nw284;
      r0 = p2;
      let _1444_v47;
      _1444_v47 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1445_v48;
      _1445_v48 = _dafny.Seq.of(_1444_v47, _1444_v47);
      let _1446_v49;
      _1446_v49 = _dafny.Seq.of(new BigNumber((_1445_v48).length));
      let _1447_v50;
      _1447_v50 = _module.D6.create_DC14(_1446_v49);
      let _source14 = _1447_v50;
      if (_source14.is_DC15) {
        let _1448___mcc_h0 = (_source14).cf20;
        let _1449___mcc_h1 = (_source14).cf21;
        let _1450___mcc_h2 = (_source14).cf22;
        let _1451_cf22 = _1450___mcc_h2;
        let _1452_cf21 = _1449___mcc_h1;
        let _1453_cf20 = _1448___mcc_h0;
        let _1454_v51;
        _1454_v51 = new BigNumber(171);
        let _rhs283 = p1;
        let _rhs284 = _1454_v51;
        let _rhs285 = (_dafny.ZERO).minus((((_1454_v51).isLessThanOrEqualTo(new BigNumber((_1445_v48).length))) ? (new BigNumber((_1446_v49).length)) : (new BigNumber(915))));
        _1454_v51 = _rhs283;
        _1454_v51 = _rhs284;
        _1454_v51 = _rhs285;
        if (_1453_cf20) {
          let _1455_v52;
          let _nw285 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
          _1455_v52 = _nw285;
          let _1456_v53;
          let _nw286 = new _module.C4();
          _nw286.__ctor(p1, (_dafny.ZERO).minus(p1), _1455_v52, _1444_v47);
          _1456_v53 = _nw286;
          let _1457_v54;
          let _nw287 = new _module.C3();
          _nw287.__ctor(_1455_v52, _1444_v47);
          _1457_v54 = _nw287;
          _1457_v54 = _1457_v54;
          r0 = (p0).IsDisjointFrom((p0).Intersect(p0));
          let _1458_v55;
          _1458_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1456_v53,_dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-737)), ((_1459_v47) => function (_1460_i8) {
            return _1459_v47;
          })(_1444_v47)), _module.__default.safeIndex(_1456_v53.f11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-737)), ((_1461_v47) => function (_1462_i8) {
            return _1461_v47;
          })(_1444_v47))).length)), _1444_v47)));
          let _1463_v56;
          _1463_v56 = _dafny.Set.fromElements(_1445_v48);
          _1458_v55 = (_1458_v55).update(_1456_v53, _1463_v56);
          let _1464_v57;
          let _init57 = ((_1465_cf20) => function (_1466_i9) {
            return _1465_cf20;
          })(_1453_cf20);
          let _nw288 = Array((new BigNumber(29)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw288.length); _i0_57++) {
            _nw288[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _1464_v57 = _nw288;
          let _1467_v58;
          _1467_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1464_v57,_1451_cf22);
          let _1468_v59;
          _1468_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1467_v58,_1454_v51);
          _1468_v59 = (_1468_v59).update(_1467_v58, _module.__default.safeModuloInt((_1456_v53).f12, (_1456_v53).f12));
        } else {
          let _1469_v60;
          _1469_v60 = _module.D7.create_DC17(true, _1452_cf21, _1451_cf22, _1453_cf20);
          let _1470_v61;
          _1470_v61 = _dafny.Seq.of(_1469_v60, _module.__default.fm32(_1453_cf20, _1452_cf21, globalState));
          _1470_v61 = ((_1453_cf20) ? (_1470_v61) : (_dafny.Seq.Concat(_1470_v61, _dafny.Seq.update(_1470_v61, _module.__default.safeIndex(p1, new BigNumber((_1470_v61).length)), _1469_v60))));
          let _1471_v62;
          let _init58 = ((_1472_cf20, _1473_cf22) => function (_1474_i10) {
            return _dafny.Set.fromElements(_1472_cf20, _1473_cf22);
          })(_1453_cf20, _1451_cf22);
          let _nw289 = Array((new BigNumber(16)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw289.length); _i0_58++) {
            _nw289[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1471_v62 = _nw289;
          let _1475_v63;
          let _nw290 = new _module.C1();
          _nw290.__ctor(_1471_v62, new _dafny.CodePoint('s'.codePointAt(0)));
          _1475_v63 = _nw290;
          let _1476_v64;
          let _nw291 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _1476_v64 = _nw291;
          let _1477_v65;
          _1477_v65 = _dafny.Seq.of(p2);
          let _1478_v66;
          _1478_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-110),_1452_cf21);
          let _1479_v67;
          _1479_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1451_cf22,p1);
          let _1480_v68;
          _1480_v68 = _dafny.MultiSet.fromElements(_1479_v67, _1479_v67);
          let _1481_v69;
          _1481_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1480_v68).cardinality()),p1);
          let _1482_v70;
          _1482_v70 = _dafny.Set.fromElements(p2);
          let _1483_v71;
          _1483_v71 = _dafny.MultiSet.fromElements(p2, p2);
          let _1484_v72;
          let _nw292 = Array((new BigNumber(24)).toNumber());
          _nw292[(_dafny.ZERO).toNumber()] = new BigNumber((p0).length);
          _nw292[(_dafny.ONE).toNumber()] = new BigNumber(755);
          _nw292[(new BigNumber(2)).toNumber()] = new BigNumber((_1477_v65).length);
          _nw292[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_1454_v51)).length);
          _nw292[(new BigNumber(4)).toNumber()] = p1;
          _nw292[(new BigNumber(5)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(6)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(7)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(8)).toNumber()] = new BigNumber((_1445_v48).length);
          _nw292[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_1454_v51);
          _nw292[(new BigNumber(10)).toNumber()] = p1;
          _nw292[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_1454_v51);
          _nw292[(new BigNumber(12)).toNumber()] = new BigNumber((_1478_v66).length);
          _nw292[(new BigNumber(13)).toNumber()] = new BigNumber(554);
          _nw292[(new BigNumber(14)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality());
          _nw292[(new BigNumber(16)).toNumber()] = p1;
          _nw292[(new BigNumber(17)).toNumber()] = p1;
          _nw292[(new BigNumber(18)).toNumber()] = p1;
          _nw292[(new BigNumber(19)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(20)).toNumber()] = new BigNumber((_1481_v69).length);
          _nw292[(new BigNumber(21)).toNumber()] = new BigNumber((_1482_v70).length);
          _nw292[(new BigNumber(22)).toNumber()] = _1454_v51;
          _nw292[(new BigNumber(23)).toNumber()] = new BigNumber((_1483_v71).cardinality());
          _1484_v72 = _nw292;
          let _1485_v73;
          _1485_v73 = _dafny.Seq.of(_1484_v72);
          let _index297 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_1476_v64).length));
          (_1476_v64)[_index297] = _dafny.Seq.Concat(_dafny.Seq.of(_1484_v72), _1485_v73);
          let _1486_v74;
          _1486_v74 = _module.D1.create_DC1(_dafny.MultiSet.fromElements(_1451_cf22));
          let _index298 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_1476_v64).length));
          let _rhs286 = _1485_v73;
          let _rhs287 = _dafny.MultiSet.fromElements(_1452_cf21);
          let _rhs288 = p1;
          let _rhs289 = !((_1475_v63).fm3(((_1452_cf21) ? (_1454_v51) : (new BigNumber(498))), _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), (_this).fm10(_1486_v74, _1451_cf22, _1478_v66, globalState)), globalState));
          let _lhs151 = _1476_v64;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_1476_v64).length));
          _lhs151[_lhs152] = _rhs286;
          _1483_v71 = _rhs287;
          _1454_v51 = _rhs288;
          _1451_cf22 = _rhs289;
          _1453_cf20 = _1452_cf21;
          let _1487_v75;
          _1487_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1484_v72,(_1452_cf21) && (_1452_cf21));
          _1487_v75 = (_1487_v75).update(_1484_v72, _1453_cf20);
        }
        if (false) {
          let _1488_v76;
          _1488_v76 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1453_cf20);
          let _1489_v77;
          let _nw293 = Array((new BigNumber(6)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = new BigNumber((_1488_v76).length);
          _nw293[(_dafny.ONE).toNumber()] = _1454_v51;
          _nw293[(new BigNumber(2)).toNumber()] = p1;
          _nw293[(new BigNumber(3)).toNumber()] = new BigNumber(491);
          _nw293[(new BigNumber(4)).toNumber()] = p1;
          _nw293[(new BigNumber(5)).toNumber()] = _1454_v51;
          _1489_v77 = _nw293;
          let _1490_v78;
          _1490_v78 = _dafny.Seq.of(_1489_v77);
          _1490_v78 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1490_v78, _1490_v78), _1490_v78);
          _1454_v51 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_1445_v48, _1445_v48)).length), p1);
          let _1491_v79;
          let _nw294 = Array((new BigNumber(8)).toNumber());
          _nw294[(_dafny.ZERO).toNumber()] = _1453_cf20;
          _nw294[(_dafny.ONE).toNumber()] = false;
          _nw294[(new BigNumber(2)).toNumber()] = p2;
          _nw294[(new BigNumber(3)).toNumber()] = p2;
          _nw294[(new BigNumber(4)).toNumber()] = _1453_cf20;
          _nw294[(new BigNumber(5)).toNumber()] = true;
          _nw294[(new BigNumber(6)).toNumber()] = _1451_cf22;
          _nw294[(new BigNumber(7)).toNumber()] = p2;
          _1491_v79 = _nw294;
          let _1492_v80;
          let _init59 = ((_1493_v76) => function (_1494_i11) {
            return _1493_v76;
          })(_1488_v76);
          let _nw295 = Array((new BigNumber(14)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw295.length); _i0_59++) {
            _nw295[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _1492_v80 = _nw295;
          let _1495_v81;
          _1495_v81 = _module.D4.create_DC10(_dafny.Set.fromElements(_1491_v79, _1491_v79), new BigNumber((_1488_v76).length), _1445_v48, _1492_v80);
          _1453_cf20 = (_this).fm3(_1454_v51, (_1495_v81).dtor_cf12, globalState);
          let _1496_v82;
          _1496_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1451_cf22,_1445_v48);
          let _1497_v83;
          _1497_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1446_v49).length),_dafny.Seq.of(new BigNumber((_1496_v82).length)));
          _1497_v83 = (_1497_v83).update((p1).minus(_1454_v51), _1446_v49);
          let _nw296 = Array((new BigNumber(5)).toNumber());
          _nw296[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm14(globalState));
          _nw296[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1445_v48, _1445_v48)).length);
          _nw296[(new BigNumber(2)).toNumber()] = p1;
          _nw296[(new BigNumber(3)).toNumber()] = p1;
          _nw296[(new BigNumber(4)).toNumber()] = _1454_v51;
          _1489_v77 = _nw296;
        } else {
          let _1498_v84;
          let _nw297 = new _module.C5();
          _nw297.__ctor();
          _1498_v84 = _nw297;
          let _1499_v85;
          let _nw298 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
          _1499_v85 = _nw298;
          let _1500_v86;
          _1500_v86 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1452_cf21);
          let _1501_v87;
          let _nw299 = new _module.C1();
          _nw299.__ctor(_1499_v85, _module.__default.fm0(new BigNumber((p0).length), new BigNumber((_1500_v86).length), globalState));
          _1501_v87 = _nw299;
          _1501_v87 = _1501_v87;
          _1444_v47 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1502_v88;
          _1502_v88 = _dafny.MultiSet.fromElements(p2);
          let _1503_v89;
          let _nw300 = Array((new BigNumber(7)).toNumber());
          _nw300[(_dafny.ZERO).toNumber()] = _1451_cf22;
          _nw300[(_dafny.ONE).toNumber()] = _1451_cf22;
          _nw300[(new BigNumber(2)).toNumber()] = _1452_cf21;
          _nw300[(new BigNumber(3)).toNumber()] = _1451_cf22;
          _nw300[(new BigNumber(4)).toNumber()] = _1452_cf21;
          _nw300[(new BigNumber(5)).toNumber()] = _1453_cf20;
          _nw300[(new BigNumber(6)).toNumber()] = !(_module.__default.fm1(new BigNumber(-183), new BigNumber((_1502_v88).cardinality()), !(_1453_cf20), _1453_cf20, globalState)) || (_1453_cf20);
          _1503_v89 = _nw300;
          let _index299 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1503_v89).length));
          (_1503_v89)[_index299] = p2;
          let _1504_v90;
          _1504_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1503_v89,((_1451_cf22) ? (p0) : (p0)));
          let _1505_v91;
          _1505_v91 = _dafny.Set.fromElements(_1452_cf21);
          let _index300 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1503_v89).length));
          let _rhs290 = (((_1500_v86).contains(_1451_cf22)) ? ((_1500_v86).get(_1451_cf22)) : (_1453_cf20));
          let _rhs291 = _dafny.Seq.Concat(_1445_v48, _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), function (_1506_i12) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), function (_1507_i12) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length)), new _dafny.CodePoint('g'.codePointAt(0))), _1445_v48));
          let _rhs292 = new BigNumber((_1504_v90).length);
          let _rhs293 = (_1505_v91).IsSubsetOf(_dafny.Set.fromElements(_1451_cf22, true, _1453_cf20));
          let _rhs294 = _1451_cf22;
          let _lhs153 = _1503_v89;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1503_v89).length));
          _1451_cf22 = _rhs290;
          _1445_v48 = _rhs291;
          _1454_v51 = _rhs292;
          _lhs153[_lhs154] = _rhs293;
          _1453_cf20 = _rhs294;
          let _1508_v92;
          _1508_v92 = _dafny.Map.Empty.slice().updateUnsafe((_1503_v89)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_1503_v89).length))],p1);
          _1508_v92 = (_1508_v92).update(!(p1).isEqualTo((_dafny.ZERO).minus(_1454_v51)), _1454_v51);
        }
        let _1509_v93;
        _1509_v93 = _dafny.Seq.of(_1451_cf22, _1453_cf20);
        let _1510_v94;
        _1510_v94 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1509_v93);
        let _1511_v95;
        _1511_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1451_cf22,new BigNumber((_1445_v48).length));
        let _1512_v96;
        _1512_v96 = _dafny.MultiSet.fromElements(_1444_v47, _1444_v47, new _dafny.CodePoint('r'.codePointAt(0)));
        let _1513_v97;
        _1513_v97 = _dafny.MultiSet.fromElements(false, p2, _1453_cf20, _1452_cf21);
        let _1514_v98;
        _1514_v98 = _dafny.MultiSet.fromElements(_1454_v51, p1, new BigNumber((_dafny.Seq.update(_1445_v48, _module.__default.safeIndex(_1454_v51, new BigNumber((_1445_v48).length)), _1444_v47)).length), _1454_v51);
        let _1515_v99;
        let _nw301 = Array((new BigNumber(26)).toNumber());
        _nw301[(_dafny.ZERO).toNumber()] = _1454_v51;
        _nw301[(_dafny.ONE).toNumber()] = p1;
        _nw301[(new BigNumber(2)).toNumber()] = _1454_v51;
        _nw301[(new BigNumber(3)).toNumber()] = new BigNumber(((((_1510_v94).contains(_1453_cf20)) ? ((_1510_v94).get(_1453_cf20)) : (_1509_v93))).length);
        _nw301[(new BigNumber(4)).toNumber()] = p1;
        _nw301[(new BigNumber(5)).toNumber()] = p1;
        _nw301[(new BigNumber(6)).toNumber()] = (_1446_v49)[_module.__default.safeIndex(p1, new BigNumber((_1446_v49).length))];
        _nw301[(new BigNumber(7)).toNumber()] = new BigNumber((_1511_v95).length);
        _nw301[(new BigNumber(8)).toNumber()] = new BigNumber(741);
        _nw301[(new BigNumber(9)).toNumber()] = new BigNumber(52);
        _nw301[(new BigNumber(10)).toNumber()] = (((_1512_v96).contains(_1444_v47)) ? ((_1512_v96).get(_1444_v47)) : (p1));
        _nw301[(new BigNumber(11)).toNumber()] = new BigNumber(813);
        _nw301[(new BigNumber(12)).toNumber()] = _1454_v51;
        _nw301[(new BigNumber(13)).toNumber()] = new BigNumber((_1513_v97).cardinality());
        _nw301[(new BigNumber(14)).toNumber()] = new BigNumber((_1509_v93).length);
        _nw301[(new BigNumber(15)).toNumber()] = new BigNumber((_1514_v98).cardinality());
        _nw301[(new BigNumber(16)).toNumber()] = new BigNumber((_module.__default.fm33(_1453_cf20, !(_1452_cf21), (_dafny.ZERO).minus(_1454_v51), globalState)).length);
        _nw301[(new BigNumber(17)).toNumber()] = _1454_v51;
        _nw301[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_1509_v93, _module.__default.safeIndex(_1454_v51, new BigNumber((_1509_v93).length)), true)).length);
        _nw301[(new BigNumber(19)).toNumber()] = p1;
        _nw301[(new BigNumber(20)).toNumber()] = _1454_v51;
        _nw301[(new BigNumber(21)).toNumber()] = p1;
        _nw301[(new BigNumber(22)).toNumber()] = _module.__default.fm14(globalState);
        _nw301[(new BigNumber(23)).toNumber()] = new BigNumber(823);
        _nw301[(new BigNumber(24)).toNumber()] = new BigNumber((p0).length);
        _nw301[(new BigNumber(25)).toNumber()] = new BigNumber((_1445_v48).length);
        _1515_v99 = _nw301;
        let _1516_v100;
        _1516_v100 = _dafny.Set.fromElements(_1515_v99, _1515_v99);
        _1452_cf21 = !((_1516_v100).IsProperSubsetOf((_1516_v100).Union(_1516_v100)));
      } else {
        let _1517___mcc_h3 = (_source14).cf19;
        let _1518_cf19 = _1517___mcc_h3;
        r0 = p2;
        let _1519_v101;
        let _nw302 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
        _1519_v101 = _nw302;
        let _index301 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1519_v101).length));
        (_1519_v101)[_index301] = _dafny.Set.fromElements(p2);
        let _1520_v102;
        _1520_v102 = _dafny.Seq.of(true);
        let _1521_v103;
        _1521_v103 = _dafny.Set.fromElements(p2, _dafny.Seq.contains(_1520_v102, (_this).fm3(new BigNumber(-959), new BigNumber(610), globalState)), p2);
        let _index302 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1519_v101).length));
        (_1519_v101)[_index302] = _1521_v103;
        let _1522_v104;
        _1522_v104 = _module.D10.create_DC21(_dafny.Set.fromElements(p2));
        let _source15 = _1522_v104;
        if (_source15.is_DC22) {
          let _1523___mcc_h4 = (_source15).cf33;
          let _1524___mcc_h5 = (_source15).cf34;
          let _1525___mcc_h6 = (_source15).cf35;
          let _1526_cf35 = _1525___mcc_h6;
          let _1527_cf34 = _1524___mcc_h5;
          let _1528_cf33 = _1523___mcc_h4;
          let _1529_v105;
          _1529_v105 = _module.D5.create_DC13(_1528_cf33, p1, p2);
          let _1530_v106;
          _1530_v106 = _dafny.MultiSet.fromElements(_1526_cf35);
          let _1531_v107;
          _1531_v107 = _dafny.Set.fromElements(_1530_v106);
          let _pat_let_tv31 = p2;
          let _pat_let_tv32 = _1526_cf35;
          let _1532_v108;
          let _nw303 = Array((new BigNumber(23)).toNumber());
          _nw303[(_dafny.ZERO).toNumber()] = p2;
          _nw303[(_dafny.ONE).toNumber()] = p2;
          _nw303[(new BigNumber(2)).toNumber()] = p2;
          _nw303[(new BigNumber(3)).toNumber()] = (_1526_cf35).isLessThan(_1526_cf35);
          _nw303[(new BigNumber(4)).toNumber()] = !(p2);
          _nw303[(new BigNumber(5)).toNumber()] = p2;
          _nw303[(new BigNumber(6)).toNumber()] = p2;
          _nw303[(new BigNumber(7)).toNumber()] = false;
          _nw303[(new BigNumber(8)).toNumber()] = p2;
          _nw303[(new BigNumber(9)).toNumber()] = p2;
          _nw303[(new BigNumber(10)).toNumber()] = p2;
          _nw303[(new BigNumber(11)).toNumber()] = p2;
          _nw303[(new BigNumber(12)).toNumber()] = p2;
          _nw303[(new BigNumber(13)).toNumber()] = p2;
          _nw303[(new BigNumber(14)).toNumber()] = !(!(_module.__default.fm1(_1526_cf35, _1526_cf35, p2, !(p2), globalState)));
          _nw303[(new BigNumber(15)).toNumber()] = (p2) && ((function (_pat_let25_0) {
            return function (_1533_dt__update__tmp_h0) {
              return function (_pat_let26_0) {
                return function (_1534_dt__update_hcf18_h0) {
                  return function (_pat_let27_0) {
                    return function (_1535_dt__update_hcf17_h0) {
                      return _module.D5.create_DC13((_1533_dt__update__tmp_h0).dtor_cf16, _1535_dt__update_hcf17_h0, _1534_dt__update_hcf18_h0);
                    }(_pat_let27_0);
                  }((_dafny.ZERO).minus(_pat_let_tv32));
                }(_pat_let26_0);
              }(_pat_let_tv31);
            }(_pat_let25_0);
          }(_1529_v105)).dtor_cf18);
          _nw303[(new BigNumber(16)).toNumber()] = p2;
          _nw303[(new BigNumber(17)).toNumber()] = (_1531_v107).equals(_1531_v107);
          _nw303[(new BigNumber(18)).toNumber()] = true;
          _nw303[(new BigNumber(19)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1520_v102, _1520_v102);
          _nw303[(new BigNumber(20)).toNumber()] = !(p1).isEqualTo(p1);
          _nw303[(new BigNumber(21)).toNumber()] = p2;
          _nw303[(new BigNumber(22)).toNumber()] = p2;
          _1532_v108 = _nw303;
          let _index303 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1532_v108).length));
          (_1532_v108)[_index303] = p2;
          let _index304 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_1532_v108).length));
          (_1532_v108)[_index304] = p2;
          let _1536_v109;
          _1536_v109 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),(_1532_v108)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_1532_v108).length))]);
          let _1537_v110;
          _1537_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1444_v47,_1536_v109);
          let _1538_v111;
          _1538_v111 = _dafny.Seq.of(_1537_v110);
          _1537_v110 = ((_1538_v111)[_module.__default.safeIndex(_1526_cf35, new BigNumber((_1538_v111).length))]).update(new _dafny.CodePoint('s'.codePointAt(0)), (_1536_v109).Merge(_1536_v109));
          let _1539_v112;
          let _nw304 = new _module.C3();
          _nw304.__ctor(_1519_v101, new _dafny.CodePoint('u'.codePointAt(0)));
          _1539_v112 = _nw304;
          _1539_v112 = _1539_v112;
          _1526_cf35 = _1526_cf35;
        } else if (_source15.is_DC21) {
          let _1540___mcc_h7 = (_source15).cf32;
          let _1541_cf32 = _1540___mcc_h7;
          let _1542_v113;
          _1542_v113 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), _1444_v47);
          let _1543_v114;
          _1543_v114 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_1444_v47)).Intersect(_1542_v113),p2);
          _1543_v114 = (_1543_v114).update(_1542_v113, p2);
          let _1544_v115;
          let _nw305 = new _module.C1();
          _nw305.__ctor(_1519_v101, _1444_v47);
          _1544_v115 = _nw305;
          let _1545_v116;
          _1545_v116 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),p2);
          r0 = !(((_1544_v115).fm5((_dafny.ZERO).minus(p1), p2, p1, new BigNumber((_1545_v116).length), globalState)) === (p2));
          let _1546_v117;
          _1546_v117 = new BigNumber(55);
          _1546_v117 = p1;
        } else {
          let _1547___mcc_h8 = (_source15).cf36;
          let _1548_cf36 = _1547___mcc_h8;
          let _1549_v118;
          _1549_v118 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          let _1550_v119;
          _1550_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1549_v118,((_module.__default.fm19(p2, p1, (_1520_v102)[_module.__default.safeIndex(p1, new BigNumber((_1520_v102).length))], globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2)));
          let _1551_v120;
          _1551_v120 = _module.D17.create_DC40(_1550_v119);
          _1550_v119 = (_1551_v120).dtor_cf62;
          let _1552_v121;
          _1552_v121 = new BigNumber(302);
          _1552_v121 = _1552_v121;
          let _1553_v122;
          _1553_v122 = _module.D1.create_DC1(_dafny.MultiSet.fromElements(!(false)));
          let _1554_v123;
          _1554_v123 = _dafny.MultiSet.fromElements(p2);
          let _1555_v124;
          _1555_v124 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1554_v123).cardinality()),!(p2));
          let _1556_v125;
          _1556_v125 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm10(_1553_v122, (((_1555_v124).contains(p1)) ? ((_1555_v124).get(p1)) : (p2)), _1555_v124, globalState)),_1552_v121);
          _1556_v125 = (_1556_v125).update(_1552_v121, (_this).fm10(_1553_v122, p2, _module.__default.fm34(p2, globalState), globalState));
          _1552_v121 = _1552_v121;
        }
        let _1557_v126;
        _1557_v126 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p2, p2),_1521_v103);
        _1557_v126 = _1557_v126;
      }
      let _1558_v127;
      let _init60 = function (_1559_i13) {
        return (_1559_i13).multipliedBy(new BigNumber(-729));
      };
      let _nw306 = Array((new BigNumber(23)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw306.length); _i0_60++) {
        _nw306[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _1558_v127 = _nw306;
      let _1560_v128;
      _1560_v128 = _dafny.Seq.of(true);
      let _1561_v129;
      _1561_v129 = _module.D15.create_DC34(_1560_v128);
      let _1562_v130;
      _1562_v130 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1558_v127);
      let _1563_v131;
      _1563_v131 = _dafny.Map.Empty.slice().updateUnsafe(_1561_v129,(((_1562_v130).contains(_module.__default.fm1(new BigNumber(-171), new BigNumber((_1560_v128).length), p2, p2, globalState))) ? ((_1562_v130).get(_module.__default.fm1(new BigNumber(-171), new BigNumber((_1560_v128).length), p2, p2, globalState))) : (_1558_v127)));
      let _1564_v132;
      let _nw307 = Array((new BigNumber(3)).toNumber());
      _nw307[(_dafny.ZERO).toNumber()] = _1558_v127;
      _nw307[(_dafny.ONE).toNumber()] = (((_1563_v131).contains(_1561_v129)) ? ((_1563_v131).get(_1561_v129)) : (_1558_v127));
      _nw307[(new BigNumber(2)).toNumber()] = _1558_v127;
      _1564_v132 = _nw307;
      let _index305 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1564_v132).length));
      (_1564_v132)[_index305] = _1558_v127;
      let _index306 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1564_v132).length));
      (_1564_v132)[_index306] = _1558_v127;
      r0 = _dafny.Seq.IsProperPrefixOf(((p2) ? (_1560_v128) : (_1560_v128)), _1560_v128);
      return r0;
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = false;
      let _1565_v0;
      let _nw308 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1565_v0 = _nw308;
      let _1566_v1;
      _1566_v1 = true;
      let _index307 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
      (_1565_v0)[_index307] = _1566_v1;
      let _1567_v2;
      _1567_v2 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1568_v3;
      _1568_v3 = _dafny.Seq.UnicodeFromString("vteybig");
      let _index308 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
      (_1565_v0)[_index308] = !_dafny.Seq.contains(_1568_v3, _1567_v2);
      let _1569_v4;
      let _nw309 = new _module.C7();
      _nw309.__ctor();
      _1569_v4 = _nw309;
      _1569_v4 = _1569_v4;
      if (!(p2).isEqualTo((new BigNumber(187)).multipliedBy(p0))) {
        let _1570_v5;
        _1570_v5 = _dafny.Set.fromElements(false);
        let _1571_v6;
        _1571_v6 = _dafny.Set.fromElements(_1570_v5, _1570_v5, (_dafny.Set.fromElements(_1566_v1)).Intersect(_1570_v5));
        let _1572_v7;
        _1572_v7 = _module.D18.create_DC44(_1571_v6);
        _1571_v6 = (_1572_v7).dtor_cf67;
        r2 = _1566_v1;
        let _1573_v8;
        _1573_v8 = new BigNumber(970);
        let _1574_v9;
        _1574_v9 = _dafny.Set.fromElements(_1568_v3);
        _1573_v8 = new BigNumber(((_1574_v9).Union(_1574_v9)).length);
        _1573_v8 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p2));
        let _1575_v10;
        _1575_v10 = _dafny.Seq.of(_1573_v8);
        let _1576_v11;
        _1576_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.IsProperPrefixOf(_1575_v10, _1575_v10));
        _1576_v11 = (_1576_v11).update(p0, true);
      } else {
        let _1577_v12;
        _1577_v12 = new BigNumber(980);
        let _1578_v13;
        _1578_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1567_v2,_1577_v12);
        _1577_v12 = ((new BigNumber((_1578_v13).length)).minus(p2)).minus((p3).multipliedBy(p3));
        let _1579_v14;
        _1579_v14 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _1580_v15;
        let _nw310 = new _module.C6();
        _nw310.__ctor((_1579_v14).Merge(_1579_v14));
        _1580_v15 = _nw310;
        let _1581_v16;
        let _nw311 = new _module.C5();
        _nw311.__ctor();
        _1581_v16 = _nw311;
        let _1582_v17;
        _1582_v17 = _dafny.Set.fromElements(true);
        let _1583_v18;
        _1583_v18 = _dafny.MultiSet.fromElements(new BigNumber((_1582_v17).length), p0, (_dafny.ZERO).minus(p2), p0);
        let _1584_v19;
        let _nw312 = Array((new BigNumber(16)).toNumber()).fill([]);
        _1584_v19 = _nw312;
        let _1585_v20;
        let _init61 = ((_1586_v12) => function (_1587_i0) {
          return _module.__default.safeDivisionInt(_1587_i0, _1586_v12);
        })(_1577_v12);
        let _nw313 = Array((new BigNumber(15)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw313.length); _i0_61++) {
          _nw313[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _1585_v20 = _nw313;
        let _index309 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1584_v19).length));
        (_1584_v19)[_index309] = _1585_v20;
        let _1588_v21;
        let _init62 = ((_1589_v3) => function (_1590_i1) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nut"), _1589_v3);
        })(_1568_v3);
        let _nw314 = Array((new BigNumber(9)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw314.length); _i0_62++) {
          _nw314[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _1588_v21 = _nw314;
        let _index310 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_1588_v21).length));
        (_1588_v21)[_index310] = _1568_v3;
        let _1591_v22;
        _1591_v22 = _dafny.Seq.of((_dafny.ZERO).minus(p3));
        let _index311 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1584_v19).length));
        let _index312 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_1588_v21).length));
        let _rhs295 = p0;
        let _rhs296 = ((_1583_v18).update(p3, _module.__default.abs(_1577_v12))).Union(_dafny.MultiSet.FromArray(((_1566_v1) ? (_1591_v22) : (_1591_v22))));
        let _rhs297 = _1585_v20;
        let _rhs298 = (p2).multipliedBy(p2);
        let _rhs299 = _1568_v3;
        let _lhs155 = _1584_v19;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1584_v19).length));
        let _lhs157 = _1588_v21;
        let _lhs158 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_1588_v21).length));
        _1577_v12 = _rhs295;
        _1583_v18 = _rhs296;
        _lhs155[_lhs156] = _rhs297;
        _1577_v12 = _rhs298;
        _lhs157[_lhs158] = _rhs299;
        if (!_dafny.areEqual(_1568_v3, _1568_v3)) {
          let _1592_v23;
          let _nw315 = new _module.C6();
          _nw315.__ctor((_1580_v15).f10);
          _1592_v23 = _nw315;
          let _1593_v24;
          let _nw316 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
          _1593_v24 = _nw316;
          _1593_v24 = (((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) ? (_1593_v24) : ((((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) ? (_1593_v24) : (_1593_v24))));
          let _1594_v25;
          _1594_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))],p3);
          let _1595_v26;
          _1595_v26 = _dafny.Seq.of(_1594_v25);
          let _1596_v27;
          _1596_v27 = _dafny.Seq.of(true);
          _1579_v14 = (_1579_v14).update((_dafny.ZERO).minus(new BigNumber((_1595_v26).length)), ((_dafny.ZERO).minus(new BigNumber((_1596_v27).length))).plus(p0));
          let _index313 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          (_1565_v0)[_index313] = (_1569_v4).fm3(p2, p3, globalState);
          let _1597_v28;
          _1597_v28 = _dafny.Set.fromElements(_1577_v12);
          _1597_v28 = _1597_v28;
        } else {
          let _1598_v29;
          let _nw317 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1598_v29 = _nw317;
          let _1599_v30;
          _1599_v30 = _module.D10.create_DC22(_1568_v3, _1598_v29, new BigNumber(((_module.__default.fm23((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], globalState)).update(_1567_v2, _module.__default.abs(p0))).cardinality()));
          let _1600_v31;
          _1600_v31 = _dafny.Seq.of(_1566_v1, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]);
          let _1601_v32;
          _1601_v32 = _dafny.Seq.of(_1588_v21, ((_module.__default.fm1(new BigNumber(686), p2, false, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], globalState)) ? (_1588_v21) : (_1588_v21)), _1588_v21, _1588_v21, _1588_v21);
          let _1602_v33;
          _1602_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1566_v1,_1600_v31);
          let _1603_v34;
          _1603_v34 = _dafny.Map.Empty.slice().updateUnsafe(((((_1580_v15).f10).contains(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_1608_v2) => function (_1609_i2) {
            return _1608_v2;
          })(_1567_v2)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_1610_v2) => function (_1611_i2) {
            return _1610_v2;
          })(_1567_v2))).length)), new _dafny.CodePoint('u'.codePointAt(0)))).length))) ? (((_1580_v15).f10).get(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_1604_v2) => function (_1605_i2) {
            return _1604_v2;
          })(_1567_v2)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_1606_v2) => function (_1607_i2) {
            return _1606_v2;
          })(_1567_v2))).length)), new _dafny.CodePoint('u'.codePointAt(0)))).length))) : (new BigNumber((_1602_v33).length))),_1581_v16);
          let _index314 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          let _index315 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          let _rhs300 = (_1600_v31)[_module.__default.safeIndex(p3, new BigNumber((_1600_v31).length))];
          let _rhs301 = (_1601_v32)[_module.__default.safeIndex(p2, new BigNumber((_1601_v32).length))];
          let _rhs302 = (p2).isLessThanOrEqualTo(p2);
          let _rhs303 = (((_1603_v34).contains((p0).multipliedBy(_module.__default.fm14(globalState)))) ? ((_1603_v34).get((p0).multipliedBy(_module.__default.fm14(globalState)))) : (_1581_v16));
          let _rhs304 = _1599_v30;
          let _lhs159 = _1565_v0;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          let _lhs161 = _1565_v0;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          _lhs159[_lhs160] = _rhs300;
          _1588_v21 = _rhs301;
          _lhs161[_lhs162] = _rhs302;
          _1581_v16 = _rhs303;
          _1599_v30 = _rhs304;
          _1577_v12 = new BigNumber((_module.__default.fm35(globalState)).length);
          r1 = !((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]);
          r3 = true;
          _1566_v1 = _1566_v1;
        }
      }
      let _index316 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
      (_1565_v0)[_index316] = !(_1566_v1);
      let _1612_v35;
      _1612_v35 = _dafny.Set.fromElements((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]);
      let _1613_v36;
      _1613_v36 = _module.D10.create_DC21(_1612_v35);
      let _1614_v37;
      _1614_v37 = _dafny.Seq.of(_1612_v35, _1612_v35);
      let _1615_v38;
      _1615_v38 = _module.D17.create_DC41(p3);
      let _pat_let_tv33 = _1612_v35;
      let _pat_let_tv34 = _1565_v0;
      let _pat_let_tv35 = _1565_v0;
      let _1616_v39;
      let _nw318 = Array((new BigNumber(24)).toNumber());
      _nw318[(_dafny.ZERO).toNumber()] = _1613_v36;
      _nw318[(_dafny.ONE).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(2)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(3)).toNumber()] = function (_pat_let28_0) {
        return function (_1617_dt__update__tmp_h0) {
          return function (_pat_let29_0) {
            return function (_1618_dt__update_hcf32_h0) {
              return _module.D10.create_DC21(_1618_dt__update_hcf32_h0);
            }(_pat_let29_0);
          }(_pat_let_tv33);
        }(_pat_let28_0);
      }(_1613_v36);
      _nw318[(new BigNumber(4)).toNumber()] = _module.D10.create_DC21((_1614_v37)[_module.__default.safeIndex(p0, new BigNumber((_1614_v37).length))]);
      _nw318[(new BigNumber(5)).toNumber()] = function (_pat_let30_0) {
        return function (_1619_dt__update__tmp_h1) {
          return function (_pat_let31_0) {
            return function (_1620_dt__update_hcf32_h1) {
              return _module.D10.create_DC21(_1620_dt__update_hcf32_h1);
            }(_pat_let31_0);
          }(_dafny.Set.fromElements((_pat_let_tv35)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_pat_let_tv34).length))]));
        }(_pat_let30_0);
      }(_1613_v36);
      _nw318[(new BigNumber(6)).toNumber()] = _module.D10.create_DC21(_1612_v35);
      _nw318[(new BigNumber(7)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(8)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(9)).toNumber()] = (((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) ? (_1613_v36) : (_1613_v36));
      _nw318[(new BigNumber(10)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(11)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(12)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(13)).toNumber()] = _module.__default.fm36(_1615_v38, false, globalState);
      _nw318[(new BigNumber(14)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(15)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(16)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(17)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(18)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(19)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(20)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(21)).toNumber()] = _module.D10.create_DC21(_1612_v35);
      _nw318[(new BigNumber(22)).toNumber()] = _1613_v36;
      _nw318[(new BigNumber(23)).toNumber()] = _1613_v36;
      _1616_v39 = _nw318;
      let _pat_let_tv36 = _1612_v35;
      let _pat_let_tv37 = p3;
      let _pat_let_tv38 = p3;
      let _pat_let_tv39 = globalState;
      let _index317 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1616_v39).length));
      (_1616_v39)[_index317] = function (_pat_let32_0) {
        return function (_1623_dt__update__tmp_h3) {
          return function (_pat_let35_0) {
            return function (_1624_dt__update_hcf32_h3) {
              return _module.D10.create_DC21(_1624_dt__update_hcf32_h3);
            }(_pat_let35_0);
          }(_module.__default.fm26(_pat_let_tv37, _pat_let_tv38, _pat_let_tv39));
        }(_pat_let32_0);
      }(function (_pat_let33_0) {
        return function (_1621_dt__update__tmp_h2) {
          return function (_pat_let34_0) {
            return function (_1622_dt__update_hcf32_h2) {
              return _module.D10.create_DC21(_1622_dt__update_hcf32_h2);
            }(_pat_let34_0);
          }(_pat_let_tv36);
        }(_pat_let33_0);
      }(_module.D10.create_DC21(_dafny.Set.fromElements(_1566_v1, false))));
      let _1625_v40;
      let _nw319 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _1625_v40 = _nw319;
      let _1626_v41;
      _1626_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1625_v40,(_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]);
      let _1627_v42;
      _1627_v42 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ojrqkow"), _1568_v3);
      let _1628_v43;
      _1628_v43 = _dafny.Seq.of(new BigNumber(((_1626_v41).update(_1625_v40, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))])).length), new BigNumber((_1627_v42).length), new BigNumber(-350), p2, new BigNumber(159));
      let _1629_v44;
      let _nw320 = Array((_dafny.ONE).toNumber());
      _nw320[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1628_v43, _1628_v43);
      _1629_v44 = _nw320;
      let _index318 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1629_v44).length));
      (_1629_v44)[_index318] = _1628_v43;
      let _index319 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1616_v39).length));
      let _index320 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1629_v44).length));
      let _rhs305 = _1613_v36;
      let _rhs306 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-906)), ((_1630_v0, _1631_p0) => function (_1632_i3) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1630_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1630_v0).length))],_1631_p0)).length);
      })(_1565_v0, p0));
      let _lhs163 = _1616_v39;
      let _lhs164 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1616_v39).length));
      let _lhs165 = _1629_v44;
      let _lhs166 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_1629_v44).length));
      _lhs163[_lhs164] = _rhs305;
      _lhs165[_lhs166] = _rhs306;
      if (_dafny.Seq.contains(_1568_v3, _1567_v2)) {
        let _index321 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length));
        (_1625_v40)[_index321] = p2;
        let _index322 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length));
        (_1625_v40)[_index322] = new BigNumber(858);
        r3 = !(_1566_v1);
        let _1633_v45;
        let _nw321 = Array((new BigNumber(7)).toNumber()).fill(_module.D6.Default());
        _1633_v45 = _nw321;
        let _1634_v46;
        _1634_v46 = _dafny.Seq.of(_1566_v1);
        let _1635_v47;
        _1635_v47 = _module.D6.create_DC15((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], (_1634_v46)[_module.__default.safeIndex(_module.__default.fm13(true, _1566_v1, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], globalState), new BigNumber((_1634_v46).length))]);
        let _index323 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1633_v45).length));
        (_1633_v45)[_index323] = _1635_v47;
        let _index324 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1633_v45).length));
        (_1633_v45)[_index324] = _1635_v47;
        let _index325 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length));
        (_1625_v40)[_index325] = (_dafny.ZERO).minus(p2);
        if ((_1569_v4).fm3((_1625_v40)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length))], new BigNumber(124), globalState)) {
          let _1636_v48;
          _1636_v48 = _dafny.MultiSet.fromElements(false);
          let _index326 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length));
          (_1625_v40)[_index326] = (((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1568_v3).length)))).plus(p0)).minus((new BigNumber((_1636_v48).cardinality())).plus(new BigNumber((_1636_v48).cardinality())));
          let _1637_v49;
          let _nw322 = new _module.C5();
          _nw322.__ctor();
          _1637_v49 = _nw322;
          let _1638_v50;
          _1638_v50 = _module.D16.create_DC38(_1568_v3, _1566_v1, p3, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]);
          let _1639_v51;
          let _nw323 = Array((new BigNumber(17)).toNumber());
          _nw323[(_dafny.ZERO).toNumber()] = _1568_v3;
          _nw323[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("ilyxr");
          _nw323[(new BigNumber(2)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(3)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(4)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_1627_v42)[_module.__default.safeIndex(p0, new BigNumber((_1627_v42).length))], _1568_v3);
          _nw323[(new BigNumber(6)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(7)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(8)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(9)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(10)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(11)).toNumber()] = _1568_v3;
          _nw323[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_1568_v3, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], _1566_v1)).cardinality()), new BigNumber((_1568_v3).length)), _1567_v2);
          _nw323[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("pkarnsgh");
          _nw323[(new BigNumber(14)).toNumber()] = ((_1566_v1) ? ((_1638_v50).dtor_cf57) : (_dafny.Seq.update(_1568_v3, _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_1568_v3).length)), new _dafny.CodePoint('y'.codePointAt(0)))));
          _nw323[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1568_v3, _1568_v3);
          _nw323[(new BigNumber(16)).toNumber()] = ((_module.__default.fm1((_1625_v40)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length))], p0, _1566_v1, _1566_v1, globalState)) ? (_1568_v3) : (_1568_v3));
          _1639_v51 = _nw323;
          let _index327 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1639_v51).length));
          (_1639_v51)[_index327] = _1568_v3;
          let _index328 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_1639_v51).length));
          (_1639_v51)[_index328] = _1568_v3;
          let _index329 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
          (_1565_v0)[_index329] = _1566_v1;
          let _1640_v52;
          _1640_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p0, new BigNumber(((_1639_v51)[_module.__default.safeIndex(new BigNumber(836), new BigNumber((_1639_v51).length))]).length)),_1566_v1);
          _1640_v52 = (_1640_v52).update(new BigNumber(433), !(!(_1566_v1)) || (_1566_v1));
        } else {
          r3 = _1566_v1;
          let _index330 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length));
          (_1625_v40)[_index330] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), ((_1641_p0) => function (_1642_i4) {
            return _1641_p0;
          })(p0)), _module.__default.safeIndex((_1625_v40)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), ((_1643_p0) => function (_1644_i4) {
            return _1643_p0;
          })(p0))).length)), p2)).length);
          let _1645_v54;
          _1645_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1634_v46,_1566_v1);
          let _1646_v55;
          let _nw324 = Array((new BigNumber(28)).toNumber());
          _nw324[(_dafny.ZERO).toNumber()] = _1614_v37;
          _nw324[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_1612_v35);
          _nw324[(new BigNumber(2)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(3)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(4)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(5)).toNumber()] = _module.__default.fm37(_1627_v42, globalState);
          _nw324[(new BigNumber(6)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), ((_1647_v35) => function (_1648_i5) {
            return _1647_v35;
          })(_1612_v35));
          _nw324[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _1614_v37);
          _nw324[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _1614_v37);
          _nw324[(new BigNumber(10)).toNumber()] = ((!(false)) ? (_1614_v37) : (_1614_v37));
          _nw324[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _dafny.Seq.of(_1612_v35));
          _nw324[(new BigNumber(12)).toNumber()] = _module.__default.fm37(_1627_v42, globalState);
          _nw324[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_1612_v35);
          _nw324[(new BigNumber(14)).toNumber()] = (((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) ? (_1614_v37) : (_1614_v37));
          _nw324[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-724)), ((_1649_v35) => function (_1650_i6) {
            return _1649_v35;
          })(_1612_v35)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-724)), ((_1651_v35) => function (_1652_i6) {
            return _1651_v35;
          })(_1612_v35))).length)), _module.__default.fm26((_dafny.ZERO).minus(p2), new BigNumber((function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_1645_v54).Keys.Elements) {
              let _1653_v53 = _compr_52;
              if ((_1645_v54).contains(_1653_v53)) {
                _coll52.push([_1653_v53,_1566_v1]);
              }
            }
            return _coll52;
          }()).length), globalState));
          _nw324[(new BigNumber(16)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(17)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_1612_v35);
          _nw324[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _1614_v37);
          _nw324[(new BigNumber(20)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_module.__default.fm37(_dafny.Seq.Create(_module.__default.abs(new BigNumber(197)), ((_1654_v3) => function (_1655_i7) {
            return _1654_v3;
          })(_1568_v3)), globalState), _module.__default.safeIndex(new BigNumber((_1568_v3).length), new BigNumber((_module.__default.fm37(_dafny.Seq.Create(_module.__default.abs(new BigNumber(197)), ((_1656_v3) => function (_1657_i7) {
            return _1656_v3;
          })(_1568_v3)), globalState)).length)), _1612_v35);
          _nw324[(new BigNumber(22)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _1614_v37);
          _nw324[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(_1566_v1, (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], !(_module.__default.fm1(p2, (_1625_v40)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_1625_v40).length))], (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))], _1566_v1, globalState)))), _1614_v37);
          _nw324[(new BigNumber(25)).toNumber()] = _1614_v37;
          _nw324[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_1614_v37, _1614_v37);
          _nw324[(new BigNumber(27)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(458)), ((_1658_v35) => function (_1659_i8) {
            return _1658_v35;
          })(_1612_v35)), _1614_v37);
          _1646_v55 = _nw324;
          let _index331 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1646_v55).length));
          (_1646_v55)[_index331] = _1614_v37;
          let _1660_v56;
          _1660_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1634_v46,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-262)), ((_1661_v35) => function (_1662_i9) {
            return _1661_v35;
          })(_1612_v35)));
          let _index332 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1646_v55).length));
          (_1646_v55)[_index332] = (((_1660_v56).contains(_1634_v46)) ? ((_1660_v56).get(_1634_v46)) : (_1614_v37));
          let _1663_v57;
          _1663_v57 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1566_v1);
          let _1664_v58;
          _1664_v58 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ktvdbuymx"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_1665_v2) => function (_1666_i10) {
            return _1665_v2;
          })(_1567_v2))), _dafny.Seq.Concat(_1568_v3, _module.__default.fm16(_1566_v1, new BigNumber((_dafny.MultiSet.fromElements(_1566_v1)).cardinality()), globalState)), _1568_v3, _dafny.Seq.Concat(_1568_v3, _module.__default.fm16(_1566_v1, new BigNumber((_1663_v57).length), globalState)));
          _1664_v58 = (_dafny.MultiSet.fromElements(_1568_v3, _1568_v3, _dafny.Seq.UnicodeFromString("sl"), _1568_v3, _1568_v3)).Intersect(_1664_v58);
          let _1667_v59;
          _1667_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1568_v3, _1568_v3),_1566_v1);
          _1667_v59 = (_1667_v59).update(_dafny.Seq.UnicodeFromString("yukmtwjs"), _1566_v1);
        }
      } else {
        _1568_v3 = _1568_v3;
        let _1668_v60;
        _1668_v60 = _module.D15.create_DC35(p3);
        let _1669_v61;
        _1669_v61 = _module.D15.create_DC36(_1668_v60);
        let _1670_v62;
        _1670_v62 = _dafny.Map.Empty.slice().updateUnsafe(((_1566_v1) ? (_1669_v61) : (_1669_v61)),_dafny.Seq.Concat(_1568_v3, _1568_v3));
        _1670_v62 = (_1670_v62).update(_1669_v61, _1568_v3);
        let _1671_v63;
        _1671_v63 = new BigNumber(-609);
        _1671_v63 = p3;
        _1671_v63 = ((p3).multipliedBy(p3)).plus(_1671_v63);
        let _1672_v64;
        _1672_v64 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1566_v1);
        let _index333 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length));
        (_1565_v0)[_index333] = (_1672_v64).contains(new BigNumber(612));
      }
      let _1673_v65;
      _1673_v65 = _dafny.Map.Empty.slice().updateUnsafe(true,_1567_v2);
      let _1674_v66;
      _1674_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1566_v1,_module.D5.create_DC11(_1673_v65));
      r0 = (((((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) ? ((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]) : ((_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))]))) ? (_1674_v66) : (_1674_v66));
      r1 = _1566_v1;
      r2 = (_1565_v0)[_module.__default.safeIndex(new BigNumber(220), new BigNumber((_1565_v0).length))];
      let _1675_v67;
      _1675_v67 = _module.D2.create_DC4(_1566_v1);
      r3 = (_1675_v67).dtor_cf5;
      return [r0, r1, r2, r3];
    }
    m7(globalState) {
      let _this = this;
      let r0 = _module.D2.Default();
      let _1676_v0;
      _1676_v0 = true;
      let _1677_v1;
      _1677_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,!(_1676_v0));
      let _1678_v2;
      _1678_v2 = new BigNumber(830);
      let _hi14 = _1678_v2;
      for (let _1679_i0 = _module.__default.fm13(_1676_v0, _1676_v0, (((_1677_v1).contains(_1676_v0)) ? ((_1677_v1).get(_1676_v0)) : (false)), globalState); _1679_i0.isLessThan(_hi14); _1679_i0 = _1679_i0.plus(_dafny.ONE)) {
        let _1680_v3;
        _1680_v3 = _dafny.Seq.UnicodeFromString("uri");
        let _1681_v4;
        let _nw325 = Array((new BigNumber(6)).toNumber());
        _nw325[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-647)), function (_1682_i1) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        });
        _nw325[(_dafny.ONE).toNumber()] = _1680_v3;
        _nw325[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1680_v3, _dafny.Seq.UnicodeFromString("dcvosk"));
        _nw325[(new BigNumber(3)).toNumber()] = _1680_v3;
        _nw325[(new BigNumber(4)).toNumber()] = _1680_v3;
        _nw325[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1680_v3, _1680_v3);
        _1681_v4 = _nw325;
        _1681_v4 = _1681_v4;
        let _1683_v5;
        _1683_v5 = _module.D15.create_DC34(_dafny.Seq.of(!(_1676_v0), _1676_v0, _1676_v0, _module.__default.fm1(_1678_v2, _1678_v2, _1676_v0, _1676_v0, globalState)));
        let _1684_v6;
        _1684_v6 = _dafny.Seq.of(true, _1676_v0, _1676_v0, _1676_v0, !(_1676_v0));
        let _pat_let_tv40 = _1684_v6;
        let _source16 = function (_pat_let36_0) {
          return function (_1685_dt__update__tmp_h0) {
            return function (_pat_let37_0) {
              return function (_1686_dt__update_hcf53_h0) {
                return _module.D15.create_DC34(_1686_dt__update_hcf53_h0);
              }(_pat_let37_0);
            }(_pat_let_tv40);
          }(_pat_let36_0);
        }(_1683_v5);
        if (_source16.is_DC35) {
          let _1687___mcc_h0 = (_source16).cf54;
          let _1688_cf54 = _1687___mcc_h0;
          let _1689_v7;
          _1689_v7 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1690_v8;
          _1690_v8 = _dafny.Set.fromElements(_dafny.Seq.contains(_1680_v3, new _dafny.CodePoint('m'.codePointAt(0))), _1676_v0, _1676_v0);
          let _rhs307 = _module.__default.fm24((_dafny.ZERO).minus(((_dafny.ZERO).minus(_1679_i0)).multipliedBy(_1679_i0)), globalState);
          let _rhs308 = _1689_v7;
          let _rhs309 = _1690_v8;
          _1680_v3 = _rhs307;
          _1689_v7 = _rhs308;
          _1690_v8 = _rhs309;
          let _1691_v9;
          _1691_v9 = _dafny.Seq.of(_module.__default.fm13(_1676_v0, _1676_v0, !(_1676_v0), globalState));
          _1676_v0 = (((function () {
            let _coll53 = new _dafny.Set();
            for (const _compr_53 of (_1691_v9).Elements) {
              let _1692_v10 = _compr_53;
              if (_dafny.Seq.contains(_1691_v9, _1692_v10)) {
                _coll53.add((_1692_v10).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(668))).cardinality())));
              }
            }
            return _coll53;
          }()).equals(function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of _dafny.IntegerRange(new BigNumber(281), new BigNumber(-758))) {
              let _1693_v11 = _compr_54;
              if (((new BigNumber(281)).isLessThanOrEqualTo(_1693_v11)) && ((_1693_v11).isLessThan(new BigNumber(-758)))) {
                _coll54.add((_1693_v11).minus(new BigNumber(-561)));
              }
            }
            return _coll54;
          }())) ? (_1676_v0) : (_1676_v0));
          let _1694_v12;
          let _nw326 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Set.Empty);
          _1694_v12 = _nw326;
          let _1695_v13;
          let _nw327 = new _module.C3();
          _nw327.__ctor(_1694_v12, (_1680_v3)[_module.__default.safeIndex(_1679_i0, new BigNumber((_1680_v3).length))]);
          _1695_v13 = _nw327;
          _1676_v0 = (((_1677_v1).contains(true)) ? ((_1677_v1).get(true)) : (_1676_v0));
        } else if (_source16.is_DC34) {
          let _1696___mcc_h1 = (_source16).cf53;
          let _1697_cf53 = _1696___mcc_h1;
          let _1698_v14;
          let _init63 = ((_1699_v0) => function (_1700_i2) {
            return _1699_v0;
          })(_1676_v0);
          let _nw328 = Array((new BigNumber(9)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw328.length); _i0_63++) {
            _nw328[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _1698_v14 = _nw328;
          let _1701_v15;
          _1701_v15 = _dafny.MultiSet.fromElements(_1698_v14, _1698_v14);
          let _1702_v16;
          _1702_v16 = _module.D2.create_DC3(_1698_v14);
          let _pat_let_tv41 = _1698_v14;
          let _1703_v17;
          _1703_v17 = _module.D2.create_DC3((function (_pat_let38_0) {
  return function (_1704_dt__update__tmp_h1) {
    return function (_pat_let39_0) {
      return function (_1705_dt__update_hcf4_h0) {
        return _module.D2.create_DC3(_1705_dt__update_hcf4_h0);
      }(_pat_let39_0);
    }(_pat_let_tv41);
  }(_pat_let38_0);
}(_1702_v16)).dtor_cf4);
          let _1706_v18;
          _1706_v18 = _dafny.Seq.of(_1702_v16);
          let _1707_v19;
          _1707_v19 = _dafny.Set.fromElements(_1676_v0);
          let _1708_v20;
          _1708_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm31(_1678_v2, new BigNumber((_1707_v19).length), globalState),new _dafny.CodePoint('c'.codePointAt(0)));
          let _pat_let_tv42 = _1698_v14;
          let _rhs310 = (_1679_i0).isEqualTo(_1679_i0);
          let _rhs311 = (((_1677_v1).contains(_1676_v0)) ? ((_1677_v1).get(_1676_v0)) : ((_1701_v15).IsDisjointFrom(_1701_v15)));
          let _rhs312 = _module.__default.fm1(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1706_v18, _module.__default.safeIndex(_1678_v2, new BigNumber((_1706_v18).length)), _1703_v17), _dafny.Seq.of(_1702_v16, function (_pat_let40_0) {
            return function (_1709_dt__update__tmp_h2) {
              return function (_pat_let41_0) {
                return function (_1710_dt__update_hcf4_h1) {
                  return _module.D2.create_DC3(_1710_dt__update_hcf4_h1);
                }(_pat_let41_0);
              }(_pat_let_tv42);
            }(_pat_let40_0);
          }(_1703_v17), _1703_v17))).length), new BigNumber((_1708_v20).length), _1676_v0, !(_1676_v0) || (_1676_v0), globalState);
          _1676_v0 = _rhs310;
          _1676_v0 = _rhs311;
          _1676_v0 = _rhs312;
          let _index334 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1698_v14).length));
          (_1698_v14)[_index334] = false;
          let _index335 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1698_v14).length));
          (_1698_v14)[_index335] = _1676_v0;
          let _1711_v21;
          let _nw329 = new _module.C0();
          _nw329.__ctor();
          _1711_v21 = _nw329;
          let _1712_v22;
          _1712_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1698_v14)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_1698_v14).length))],_1678_v2);
          let _1713_v24;
          _1713_v24 = _dafny.Seq.of(_1678_v2, _1679_i0);
          let _1714_v25;
          _1714_v25 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of _dafny.IntegerRange(new BigNumber(-212), new BigNumber(277))) {
              let _1715_v23 = _compr_55;
              if (((new BigNumber(-212)).isLessThanOrEqualTo(_1715_v23)) && ((_1715_v23).isLessThan(new BigNumber(277)))) {
                _coll55.push([_module.__default.safeModuloInt(_1715_v23, _1679_i0),_1679_i0]);
              }
            }
            return _coll55;
          }(),(_1713_v24)[_module.__default.safeIndex(_1678_v2, new BigNumber((_1713_v24).length))]);
          let _1716_v26;
          _1716_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1679_i0,new BigNumber(76));
          let _1717_v27;
          _1717_v27 = _module.D10.create_DC21(_module.__default.fm26((((_1712_v22).contains(true)) ? ((_1712_v22).get(true)) : (_1679_i0)), (((_1714_v25).contains(_1716_v26)) ? ((_1714_v25).get(_1716_v26)) : (_1679_i0)), globalState));
          _1707_v19 = ((_1707_v19).Intersect(_1707_v19)).Difference(((_1717_v27).dtor_cf32).Union(_1707_v19));
        } else {
          let _1718___mcc_h2 = (_source16).cf55;
          let _1719_cf55 = _1718___mcc_h2;
          let _1720_v28;
          let _nw330 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1720_v28 = _nw330;
          let _index336 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1720_v28).length));
          (_1720_v28)[_index336] = _1676_v0;
          let _index337 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1720_v28).length));
          (_1720_v28)[_index337] = _1676_v0;
          let _index338 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1720_v28).length));
          (_1720_v28)[_index338] = _1676_v0;
          let _1721_v29;
          _1721_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_1678_v2, _1678_v2, _1676_v0, _1676_v0, globalState),_1679_i0);
          _1721_v29 = (_1721_v29).update(_1676_v0, (_1679_i0).plus(new BigNumber((_1684_v6).length)));
          let _1722_v30;
          _1722_v30 = new _dafny.CodePoint('l'.codePointAt(0));
          _1722_v30 = _1722_v30;
        }
        _1677_v1 = (_1677_v1).Merge((_1677_v1).update(true, _1676_v0));
        let _1723_v31;
        _1723_v31 = _dafny.MultiSet.fromElements(_1678_v2);
        let _1724_v32;
        _1724_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(861),(_dafny.ZERO).minus(_1678_v2));
        let _1725_v33;
        _1725_v33 = _dafny.Seq.of(new BigNumber((_1684_v6).length), _1678_v2, (_dafny.ZERO).minus(new BigNumber((_1724_v32).length)));
        let _1726_v34;
        _1726_v34 = _dafny.Seq.of((_1725_v33)[_module.__default.safeIndex(_1678_v2, new BigNumber((_1725_v33).length))], (_dafny.ZERO).minus(_1679_i0), _1679_i0);
        let _1727_v35;
        _1727_v35 = _dafny.Seq.of(_1678_v2, new BigNumber((_1677_v1).length), new BigNumber((_1723_v31).cardinality()), new BigNumber((_1725_v33).length), new BigNumber(-724));
        let _1728_v36;
        let _nw331 = Array((new BigNumber(28)).toNumber());
        _nw331[(_dafny.ZERO).toNumber()] = _1676_v0;
        _nw331[(_dafny.ONE).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(2)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(3)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(4)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(5)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(6)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(7)).toNumber()] = !(_1676_v0);
        _nw331[(new BigNumber(8)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(9)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(10)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(11)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(12)).toNumber()] = !(_1676_v0);
        _nw331[(new BigNumber(13)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(14)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(15)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(16)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(17)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(18)).toNumber()] = false;
        _nw331[(new BigNumber(19)).toNumber()] = false;
        _nw331[(new BigNumber(20)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(21)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(22)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(23)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(24)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(25)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(26)).toNumber()] = _1676_v0;
        _nw331[(new BigNumber(27)).toNumber()] = _1676_v0;
        _1728_v36 = _nw331;
        let _1729_v37;
        _1729_v37 = _dafny.MultiSet.fromElements(_1728_v36, _1728_v36, _1728_v36);
        let _1730_v38;
        let _nw332 = Array((new BigNumber(19)).toNumber());
        _nw332[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1727_v35, _module.__default.safeIndex(_1679_i0, new BigNumber((_1727_v35).length)), _1679_i0);
        _nw332[(_dafny.ONE).toNumber()] = _1725_v33;
        _nw332[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1726_v34, _dafny.Seq.of(_1678_v2, (_dafny.ZERO).minus(_1679_i0)));
        _nw332[(new BigNumber(3)).toNumber()] = _1727_v35;
        _nw332[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), ((_1731_i0) => function (_1732_i3) {
          return _1731_i0;
        })(_1679_i0));
        _nw332[(new BigNumber(5)).toNumber()] = _1725_v33;
        _nw332[(new BigNumber(6)).toNumber()] = _1727_v35;
        _nw332[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_1679_i0);
        _nw332[(new BigNumber(8)).toNumber()] = _1726_v34;
        _nw332[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1725_v33, _1725_v33);
        _nw332[(new BigNumber(10)).toNumber()] = _1726_v34;
        _nw332[(new BigNumber(11)).toNumber()] = _1725_v33;
        _nw332[(new BigNumber(12)).toNumber()] = _1726_v34;
        _nw332[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_1679_i0, _1679_i0);
        _nw332[(new BigNumber(14)).toNumber()] = _1727_v35;
        _nw332[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1725_v33, _1726_v34), _module.__default.safeIndex(_1678_v2, new BigNumber((_dafny.Seq.Concat(_1725_v33, _1726_v34)).length)), new BigNumber((_1729_v37).cardinality()));
        _nw332[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(new BigNumber((_1684_v6).length), new BigNumber(188), _1679_i0);
        _nw332[(new BigNumber(17)).toNumber()] = _1726_v34;
        _nw332[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-432)), ((_1733_i0) => function (_1734_i4) {
          return _1733_i0;
        })(_1679_i0)), _1727_v35);
        _1730_v38 = _nw332;
        let _index339 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_1730_v38).length));
        (_1730_v38)[_index339] = _dafny.Seq.of(_1679_i0, _1678_v2, _1678_v2);
        let _1735_v39;
        _1735_v39 = _dafny.Seq.of(_1725_v33, _1727_v35, _1727_v35);
        let _1736_v40;
        let _init64 = function (_1737_i5) {
          return _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)));
        };
        let _nw333 = Array((new BigNumber(4)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw333.length); _i0_64++) {
          _nw333[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _1736_v40 = _nw333;
        let _1738_v41;
        _1738_v41 = _dafny.MultiSet.fromElements(_1736_v40);
        let _index340 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_1730_v38).length));
        (_1730_v38)[_index340] = (_1735_v39)[_module.__default.safeIndex((((_1738_v41).contains(_1736_v40)) ? ((_1738_v41).get(_1736_v40)) : (_1678_v2)), new BigNumber((_1735_v39).length))];
      }
      let _1739_v42;
      _1739_v42 = new _dafny.CodePoint('v'.codePointAt(0));
      let _1740_v43;
      _1740_v43 = _dafny.MultiSet.fromElements(_1739_v42);
      let _hi15 = new BigNumber((_1740_v43).cardinality());
      for (let _1741_i6 = _1678_v2; _1741_i6.isLessThan(_hi15); _1741_i6 = _1741_i6.plus(_dafny.ONE)) {
        let _1742_v44;
        _1742_v44 = _dafny.Seq.UnicodeFromString("vxncedkul");
        _1676_v0 = _dafny.Seq.IsPrefixOf(_1742_v44, _dafny.Seq.Concat(_1742_v44, _1742_v44));
        let _1743_v45;
        let _nw334 = new _module.C7();
        _nw334.__ctor();
        _1743_v45 = _nw334;
        let _1744_v46;
        let _nw335 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
        _1744_v46 = _nw335;
        let _1745_v47;
        _1745_v47 = _dafny.Seq.of(_1676_v0);
        let _1746_v48;
        _1746_v48 = _dafny.MultiSet.fromElements(_1745_v47, _dafny.Seq.of(_1676_v0));
        let _1747_v49;
        _1747_v49 = _dafny.Seq.of(new BigNumber((_1746_v48).cardinality()), _1741_i6);
        let _index341 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1744_v46).length));
        (_1744_v46)[_index341] = _1747_v49;
        let _index342 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1744_v46).length));
        (_1744_v46)[_index342] = _dafny.Seq.of(_module.__default.fm13(_1676_v0, _1676_v0, _1676_v0, globalState));
        let _1748_v50;
        let _nw336 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
        _1748_v50 = _nw336;
        let _1749_v51;
        let _nw337 = new _module.C2();
        _nw337.__ctor(_1748_v50, new _dafny.CodePoint('o'.codePointAt(0)));
        _1749_v51 = _nw337;
        let _1750_v52;
        _1750_v52 = _dafny.Seq.of(_1749_v51, _1749_v51);
        _1676_v0 = _dafny.Seq.IsProperPrefixOf(_1750_v52, _1750_v52);
      }
      let _1751_v53;
      _1751_v53 = _dafny.Set.fromElements(_1678_v2, _1678_v2, _1678_v2);
      let _1752_v54;
      _1752_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1676_v0,new BigNumber((_1751_v53).length));
      let _1753_v55;
      _1753_v55 = _dafny.Seq.of(new BigNumber((_1752_v54).length), new BigNumber((_1677_v1).length), _1678_v2);
      let _1754_v56;
      let _nw338 = new _module.C0();
      _nw338.__ctor();
      _1754_v56 = _nw338;
      let _1755_v57;
      _1755_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-64),_1676_v0);
      let _1756_v58;
      _1756_v58 = _dafny.MultiSet.fromElements((_1753_v55)[_module.__default.safeIndex(new BigNumber((_1755_v57).length), new BigNumber((_1753_v55).length))], new BigNumber((_dafny.Seq.update(_1753_v55, _module.__default.safeIndex((_dafny.ZERO).minus(_1678_v2), new BigNumber((_1753_v55).length)), _1678_v2)).length), new BigNumber((_1753_v55).length), _1678_v2);
      let _1757_v59;
      let _nw339 = Array((new BigNumber(23)).toNumber());
      _nw339[(_dafny.ZERO).toNumber()] = _1676_v0;
      _nw339[(_dafny.ONE).toNumber()] = _module.__default.fm1(new BigNumber((_1753_v55).length), new BigNumber((_dafny.Set.fromElements(_1754_v56, _1754_v56)).length), _1676_v0, _1676_v0, globalState);
      _nw339[(new BigNumber(2)).toNumber()] = !(_1676_v0);
      _nw339[(new BigNumber(3)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(4)).toNumber()] = (_1754_v56).fm18(_1676_v0, _1678_v2, _1676_v0, (_dafny.ZERO).minus(new BigNumber((_1756_v58).cardinality())), globalState);
      _nw339[(new BigNumber(5)).toNumber()] = !(_1676_v0);
      _nw339[(new BigNumber(6)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(7)).toNumber()] = true;
      _nw339[(new BigNumber(8)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(9)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(10)).toNumber()] = !(_1676_v0);
      _nw339[(new BigNumber(11)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(12)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(13)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(14)).toNumber()] = false;
      _nw339[(new BigNumber(15)).toNumber()] = !((_1754_v56).fm18(_1676_v0, _1678_v2, _1676_v0, _1678_v2, globalState));
      _nw339[(new BigNumber(16)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(17)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(18)).toNumber()] = !(_1676_v0);
      _nw339[(new BigNumber(19)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(20)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(21)).toNumber()] = _1676_v0;
      _nw339[(new BigNumber(22)).toNumber()] = _1676_v0;
      _1757_v59 = _nw339;
      let _1758_v60;
      _1758_v60 = _module.D2.create_DC3(_1757_v59);
      let _source17 = _1758_v60;
      if (_source17.is_DC4) {
        let _1759___mcc_h3 = (_source17).cf5;
        let _1760_cf5 = _1759___mcc_h3;
        let _1761_v61;
        _1761_v61 = _dafny.Seq.UnicodeFromString("smxtmhf");
        let _1762_v62;
        _1762_v62 = _dafny.Seq.of(_1761_v61, _1761_v61);
        let _1763_v63;
        let _nw340 = Array((new BigNumber(7)).toNumber());
        _nw340[(_dafny.ZERO).toNumber()] = _module.__default.fm13(_1760_cf5, _1760_cf5, true, globalState);
        _nw340[(_dafny.ONE).toNumber()] = _1678_v2;
        _nw340[(new BigNumber(2)).toNumber()] = _1678_v2;
        _nw340[(new BigNumber(3)).toNumber()] = _1678_v2;
        _nw340[(new BigNumber(4)).toNumber()] = _1678_v2;
        _nw340[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1762_v62).length));
        _nw340[(new BigNumber(6)).toNumber()] = _1678_v2;
        _1763_v63 = _nw340;
        let _1764_v64;
        _1764_v64 = _dafny.Seq.of(_1763_v63, _1763_v63, _1763_v63);
        let _1765_v65;
        _1765_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1739_v42,(_1764_v64)[_module.__default.safeIndex(_1678_v2, new BigNumber((_1764_v64).length))]);
        let _1766_v66;
        let _nw341 = Array((new BigNumber(28)).toNumber());
        _nw341[(_dafny.ZERO).toNumber()] = _1763_v63;
        _nw341[(_dafny.ONE).toNumber()] = (_1764_v64)[_module.__default.safeIndex(new BigNumber((_1761_v61).length), new BigNumber((_1764_v64).length))];
        _nw341[(new BigNumber(2)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(3)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(4)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(5)).toNumber()] = ((_1676_v0) ? (_1763_v63) : (_1763_v63));
        _nw341[(new BigNumber(6)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(7)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(8)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(9)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(10)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(11)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(12)).toNumber()] = ((_1760_cf5) ? (_1763_v63) : (_1763_v63));
        _nw341[(new BigNumber(13)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(14)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(15)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(16)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(17)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(18)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(19)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(20)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(21)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(22)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(23)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(24)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(25)).toNumber()] = _1763_v63;
        _nw341[(new BigNumber(26)).toNumber()] = ((_1676_v0) ? ((((_1765_v65).contains(_1739_v42)) ? ((_1765_v65).get(_1739_v42)) : (_1763_v63))) : (_1763_v63));
        _nw341[(new BigNumber(27)).toNumber()] = _1763_v63;
        _1766_v66 = _nw341;
        let _index343 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1766_v66).length));
        (_1766_v66)[_index343] = _1763_v63;
        let _1767_v67;
        _1767_v67 = _module.D7.create_DC17(_1760_cf5, _1760_cf5, _1760_cf5, _1676_v0);
        let _index344 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1766_v66).length));
        let _rhs313 = _1678_v2;
        let _rhs314 = (((_1760_cf5) ? (_module.D7.create_DC17(_1676_v0, _1676_v0, _1760_cf5, _1760_cf5)) : (_1767_v67))).dtor_cf25;
        let _rhs315 = _1763_v63;
        let _rhs316 = _module.__default.fm14(globalState);
        let _lhs167 = _1766_v66;
        let _lhs168 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_1766_v66).length));
        _1678_v2 = _rhs313;
        _1676_v0 = _rhs314;
        _lhs167[_lhs168] = _rhs315;
        _1678_v2 = _rhs316;
        let _1768_v68;
        _1768_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1676_v0,_1761_v61);
        _1761_v61 = (((_1768_v68).contains(_1676_v0)) ? ((_1768_v68).get(_1676_v0)) : (_1761_v61));
        _1678_v2 = _1678_v2;
        _1678_v2 = _1678_v2;
      } else if (_source17.is_DC5) {
        let _1769___mcc_h4 = (_source17).cf6;
        let _1770_cf6 = _1769___mcc_h4;
        _1676_v0 = !(!((_module.D7.create_DC17(_1676_v0, _1676_v0, _1676_v0, !(!(_1676_v0)))).dtor_cf27));
        let _index345 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index345] = _1676_v0;
        let _index346 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index346] = _1676_v0;
        let _1771_v69;
        let _nw342 = new _module.C0();
        _nw342.__ctor();
        _1771_v69 = _nw342;
        let _index347 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index347] = (((_1755_v57).contains(_1678_v2)) ? ((_1755_v57).get(_1678_v2)) : ((_1757_v59)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1757_v59).length))]));
      } else if (_source17.is_DC6) {
        let _1772___mcc_h5 = (_source17).cf7;
        let _1773_cf7 = _1772___mcc_h5;
        let _1774_v70;
        _1774_v70 = _dafny.Seq.UnicodeFromString("xvnt");
        if (!(_dafny.Seq.IsPrefixOf(_1774_v70, _1774_v70))) {
          _1774_v70 = _dafny.Seq.Concat(_1774_v70, _1774_v70);
          _1755_v57 = (_1755_v57).update(_module.__default.safeDivisionInt(_1678_v2, new BigNumber(-564)), _1676_v0);
          _1678_v2 = _module.__default.safeDivisionInt(_1773_cf7, new BigNumber((_dafny.Seq.Concat(_1774_v70, _1774_v70)).length));
          _1676_v0 = _1676_v0;
          _1676_v0 = _1676_v0;
        } else {
          let _1775_v71;
          _1775_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1774_v70,_1678_v2);
          _1773_cf7 = (_dafny.ZERO).minus((((_1775_v71).contains(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_1780_v42) => function (_1781_i7) {
            return _1780_v42;
          })(_1739_v42)), _dafny.Seq.UnicodeFromString("mvdekoxc")), _module.__default.safeIndex(_1678_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_1782_v42) => function (_1783_i7) {
            return _1782_v42;
          })(_1739_v42)), _dafny.Seq.UnicodeFromString("mvdekoxc"))).length)), _1739_v42))) ? ((_1775_v71).get(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_1776_v42) => function (_1777_i7) {
            return _1776_v42;
          })(_1739_v42)), _dafny.Seq.UnicodeFromString("mvdekoxc")), _module.__default.safeIndex(_1678_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(985)), ((_1778_v42) => function (_1779_i7) {
            return _1778_v42;
          })(_1739_v42)), _dafny.Seq.UnicodeFromString("mvdekoxc"))).length)), _1739_v42))) : (_module.__default.fm13(_1676_v0, _1676_v0, !(_1676_v0), globalState))));
          _1676_v0 = _1676_v0;
          let _1784_v72;
          let _init65 = ((_1785_cf7) => function (_1786_i8) {
            return (_1786_i8).plus(_1785_cf7);
          })(_1773_cf7);
          let _nw343 = Array((new BigNumber(20)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw343.length); _i0_65++) {
            _nw343[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _1784_v72 = _nw343;
          let _1787_v73;
          _1787_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1784_v72,_1773_cf7);
          _1787_v73 = (_1787_v73).update(_1784_v72, _1773_cf7);
          _1676_v0 = _1676_v0;
          _1676_v0 = (_1676_v0) || ((_1678_v2).isEqualTo(_1773_cf7));
        }
        let _1788_v74;
        _1788_v74 = _dafny.Seq.of((_1754_v56).fm3(new BigNumber(-252), new BigNumber((_1677_v1).length), globalState));
        _1788_v74 = _1788_v74;
        let _1789_v75;
        _1789_v75 = _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(_1676_v0,_1739_v42));
        let _1790_v76;
        _1790_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1676_v0,_1739_v42);
        let _1791_v77;
        _1791_v77 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1678_v2, _1773_cf7, globalState),_dafny.Map.Empty.slice().updateUnsafe(_1676_v0,_1739_v42));
        let _pat_let_tv43 = _1791_v77;
        let _pat_let_tv44 = _1739_v42;
        let _pat_let_tv45 = _1790_v76;
        let _pat_let_tv46 = _1791_v77;
        let _pat_let_tv47 = _1739_v42;
        let _pat_let_tv48 = _1789_v75;
        let _1792_v78;
        let _nw344 = Array((new BigNumber(26)).toNumber());
        _nw344[(_dafny.ZERO).toNumber()] = _1789_v75;
        _nw344[(_dafny.ONE).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(2)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(3)).toNumber()] = _module.D5.create_DC11(_1790_v76);
        _nw344[(new BigNumber(4)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(5)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(6)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(7)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(8)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(9)).toNumber()] = _module.D5.create_DC11((_1790_v76).update(_1676_v0, _1739_v42));
        _nw344[(new BigNumber(10)).toNumber()] = _module.__default.fm38(_1789_v75, _1773_cf7, new BigNumber((_dafny.Seq.of(_1773_cf7)).length), _1678_v2, globalState);
        _nw344[(new BigNumber(11)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(12)).toNumber()] = ((_1676_v0) ? (_module.D5.create_DC11((_1790_v76).update(_1676_v0, (_1774_v70)[_module.__default.safeIndex(_1773_cf7, new BigNumber((_1774_v70).length))]))) : (_1789_v75));
        _nw344[(new BigNumber(13)).toNumber()] = function (_pat_let42_0) {
          return function (_1793_dt__update__tmp_h3) {
            return function (_pat_let43_0) {
              return function (_1794_dt__update_hcf15_h0) {
                return _module.D5.create_DC11(_1794_dt__update_hcf15_h0);
              }(_pat_let43_0);
            }((((_pat_let_tv46).contains(_pat_let_tv47)) ? ((_pat_let_tv43).get(_pat_let_tv44)) : (_pat_let_tv45)));
          }(_pat_let42_0);
        }(_1789_v75);
        _nw344[(new BigNumber(14)).toNumber()] = _module.D5.create_DC11(_1790_v76);
        _nw344[(new BigNumber(15)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(16)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(17)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(18)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(19)).toNumber()] = _module.D5.create_DC11(_1790_v76);
        _nw344[(new BigNumber(20)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(21)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(22)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(23)).toNumber()] = _module.D5.create_DC11(_1790_v76);
        _nw344[(new BigNumber(24)).toNumber()] = _1789_v75;
        _nw344[(new BigNumber(25)).toNumber()] = function (_pat_let44_0) {
          return function (_1795_dt__update__tmp_h4) {
            return function (_pat_let45_0) {
              return function (_1796_dt__update_hcf15_h1) {
                return _module.D5.create_DC11(_1796_dt__update_hcf15_h1);
              }(_pat_let45_0);
            }((_pat_let_tv48).dtor_cf15);
          }(_pat_let44_0);
        }(_1789_v75);
        _1792_v78 = _nw344;
        let _index348 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1792_v78).length));
        (_1792_v78)[_index348] = _1789_v75;
        let _index349 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1792_v78).length));
        let _rhs317 = _1789_v75;
        let _rhs318 = _1739_v42;
        let _lhs169 = _1792_v78;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1792_v78).length));
        _lhs169[_lhs170] = _rhs317;
        _1739_v42 = _rhs318;
        let _index350 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index350] = true;
        let _index351 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index351] = (((_1677_v1).contains(_1676_v0)) ? ((_1677_v1).get(_1676_v0)) : (_dafny.Seq.IsProperPrefixOf(_1774_v70, _1774_v70)));
      } else if (_source17.is_DC3) {
        let _1797___mcc_h6 = (_source17).cf4;
        let _1798_cf4 = _1797___mcc_h6;
        let _1799_v79;
        let _nw345 = new _module.C7();
        _nw345.__ctor();
        _1799_v79 = _nw345;
        let _1800_v80;
        _1800_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1678_v2,_dafny.Map.Empty.slice().updateUnsafe(_1676_v0,new BigNumber(481)));
        _1800_v80 = _1800_v80;
        let _1801_v81;
        _1801_v81 = _dafny.Seq.of(_1676_v0, _1676_v0, (_1754_v56).fm18(_1676_v0, _1678_v2, _1676_v0, new BigNumber(54), globalState), _1676_v0, _1676_v0);
        let _rhs319 = _1678_v2;
        let _rhs320 = ((_1801_v81)[_module.__default.safeIndex(_1678_v2, new BigNumber((_1801_v81).length))]) || (_1676_v0);
        let _rhs321 = _1676_v0;
        let _rhs322 = _1739_v42;
        _1678_v2 = _rhs319;
        _1676_v0 = _rhs320;
        _1676_v0 = _rhs321;
        _1739_v42 = _rhs322;
        _1801_v81 = _dafny.Seq.update(_1801_v81, _module.__default.safeIndex(_1678_v2, new BigNumber((_1801_v81).length)), _1676_v0);
      } else {
        let _1802___mcc_h7 = (_source17).cf8;
        let _1803_cf8 = _1802___mcc_h7;
        let _1804_v82;
        _1804_v82 = _dafny.MultiSet.fromElements(_1753_v55, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-841)), ((_1805_v2) => function (_1806_i9) {
          return (_dafny.ZERO).minus(_1805_v2);
        })(_1678_v2)), _1753_v55);
        _1676_v0 = (_1804_v82).IsSubsetOf(_1804_v82);
        _1676_v0 = (((_1755_v57).contains(new BigNumber((_dafny.Seq.UnicodeFromString("uadctpfh")).length))) ? ((_1755_v57).get(new BigNumber((_dafny.Seq.UnicodeFromString("uadctpfh")).length))) : ((_1754_v56).fm3(_1678_v2, _1678_v2, globalState)));
        let _1807_v83;
        _1807_v83 = _dafny.Seq.of(_1676_v0);
        let _rhs323 = (_1754_v56).fm18(_1676_v0, (_module.__default.fm13(!(_1676_v0), _1676_v0, false, globalState)).multipliedBy(_1678_v2), _1676_v0, _1678_v2, globalState);
        let _rhs324 = (_1678_v2).minus(_module.__default.safeModuloInt(_1678_v2, _1678_v2));
        let _rhs325 = !(_1678_v2).isEqualTo((_dafny.ZERO).minus(_module.__default.fm13((_1807_v83)[_module.__default.safeIndex(_1678_v2, new BigNumber((_1807_v83).length))], _1676_v0, _1676_v0, globalState)));
        let _rhs326 = (_1756_v58).IsProperSubsetOf(_1756_v58);
        _1676_v0 = _rhs323;
        _1678_v2 = _rhs324;
        _1676_v0 = _rhs325;
        _1676_v0 = _rhs326;
        _1678_v2 = _1678_v2;
      }
      let _1808_v84;
      let _init66 = ((_1809_v42, _1810_v55, _1811_v0) => function (_1812_i10) {
        return _dafny.Set.fromElements(true, (_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_1813_v42) => function (_1814_i11) {
  return _1813_v42;
})(_1809_v42)), new BigNumber((_1810_v55).length), _1811_v0)).dtor_cf18, false);
      })(_1739_v42, _1753_v55, _1676_v0);
      let _nw346 = Array((new BigNumber(14)).toNumber());
      for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw346.length); _i0_66++) {
        _nw346[_i0_66] = _init66(new BigNumber(_i0_66));
      }
      _1808_v84 = _nw346;
      let _1815_v85;
      let _nw347 = new _module.C1();
      _nw347.__ctor(_1808_v84, _1739_v42);
      _1815_v85 = _nw347;
      let _1816_i12;
      _1816_i12 = _dafny.ZERO;
      L7: {
        while (_1676_v0) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1816_i12)) {
              break L7;
            }
            _1816_i12 = (_1816_i12).plus(_dafny.ONE);
            let _1817_v86;
            let _nw348 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _1817_v86 = _nw348;
            let _index352 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1817_v86).length));
            (_1817_v86)[_index352] = _1678_v2;
            let _index353 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1817_v86).length));
            (_1817_v86)[_index353] = (((_1752_v54).contains(_1676_v0)) ? ((_1752_v54).get(_1676_v0)) : (_1678_v2));
            _1676_v0 = _1676_v0;
            let _1818_v87;
            _1818_v87 = _dafny.Map.Empty.slice().updateUnsafe(!(_1676_v0),_1739_v42);
            let _1819_v88;
            _1819_v88 = _module.D5.create_DC11(_1818_v87);
            let _1820_v89;
            _1820_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1819_v88,_1676_v0);
            let _1821_v90;
            _1821_v90 = _dafny.Seq.of(_1820_v89);
            _1821_v90 = _1821_v90;
            let _1822_v91;
            _1822_v91 = _dafny.Map.Empty.slice().updateUnsafe((_1817_v86)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_1817_v86).length))],(_dafny.ZERO).minus(new BigNumber(-668)));
            _1822_v91 = (_1822_v91).update(new BigNumber(279), (((_1752_v54).contains(_1676_v0)) ? ((_1752_v54).get(_1676_v0)) : ((_1817_v86)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_1817_v86).length))])));
          }
        }
      }
      let _1823_v92;
      _1823_v92 = _dafny.MultiSet.fromElements(_1676_v0, _1676_v0, _1676_v0);
      if ((_1823_v92).IsProperSubsetOf(_1823_v92)) {
        let _index354 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index354] = _1676_v0;
        let _index355 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1757_v59).length));
        (_1757_v59)[_index355] = _1676_v0;
        _1676_v0 = _module.__default.fm1(_1678_v2, _1678_v2, (_1757_v59)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_1757_v59).length))], _1676_v0, globalState);
        let _1824_v93;
        _1824_v93 = _dafny.Set.fromElements(!(_1676_v0));
        let _rhs327 = _1824_v93;
        let _rhs328 = _1739_v42;
        _1824_v93 = _rhs327;
        _1739_v42 = _rhs328;
        let _1825_v94;
        let _nw349 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1825_v94 = _nw349;
        _1825_v94 = ((_1676_v0) ? (_1825_v94) : (_1825_v94));
        let _1826_v95;
        _1826_v95 = _dafny.Seq.UnicodeFromString("smj");
        _1676_v0 = _dafny.Seq.IsPrefixOf(_1826_v95, _1826_v95);
      } else {
        let _1827_v96;
        _1827_v96 = _dafny.MultiSet.fromElements(_1757_v59, _1757_v59, _1757_v59);
        _1827_v96 = _dafny.MultiSet.fromElements(_1757_v59, _1757_v59);
        _1677_v1 = (_1677_v1).update(((_dafny.ZERO).minus(_1678_v2)).isLessThanOrEqualTo((_dafny.ZERO).minus(_1678_v2)), !(_1676_v0));
        let _1828_v97;
        let _nw350 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _1828_v97 = _nw350;
        _1828_v97 = _1828_v97;
        let _1829_v98;
        _1829_v98 = _dafny.Seq.of(true, _1676_v0);
        let _1830_v99;
        _1830_v99 = _dafny.Map.Empty.slice().updateUnsafe(_1676_v0,_dafny.Seq.Concat(_1829_v98, _1829_v98));
        let _1831_v100;
        _1831_v100 = _dafny.Set.fromElements(_1676_v0);
        _1830_v99 = (_1830_v99).update((_1815_v85).fm5(_1678_v2, _1676_v0, new BigNumber((_1831_v100).length), _1678_v2, globalState), _dafny.Seq.Concat(_dafny.Seq.update(_1829_v98, _module.__default.safeIndex(_1678_v2, new BigNumber((_1829_v98).length)), !((_1815_v85).fm5(_1678_v2, _1676_v0, new BigNumber((_dafny.Seq.UnicodeFromString("f")).length), _1678_v2, globalState))), _1829_v98));
        let _1832_v101;
        _1832_v101 = _module.D7.create_DC17(_1676_v0, true, !(_1676_v0), _1676_v0);
        let _1833_v102;
        _1833_v102 = _dafny.Seq.of(_1832_v101);
        _1833_v102 = _dafny.Seq.of(_1832_v101);
      }
      let _1834_v103;
      _1834_v103 = _module.D2.create_DC6(new BigNumber(374));
      r0 = _1834_v103;
      return r0;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return !(_dafny.MultiSet.fromElements(false)).contains((true) || (true));
    };
    fm4(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(false, false)).length), (new BigNumber((_dafny.Seq.of((_module.D5.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-635)), function (_1835_i0) {
  return (_this).f5;
}), new BigNumber(556), false)).dtor_cf17)).length)).minus(new BigNumber((_dafny.Seq.of(true, true)).length)));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber(703)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let _1836_v0;
      _1836_v0 = true;
      let _1837_v1;
      _1837_v1 = new BigNumber(-526);
      let _source18 = _module.D2.create_DC4((_1836_v0) || (_module.__default.fm1(_1837_v1, _1837_v1, _1836_v0, _1836_v0, globalState)));
      if (_source18.is_DC4) {
        let _1838___mcc_h0 = (_source18).cf5;
        let _1839_cf5 = _1838___mcc_h0;
        let _1840_v2;
        let _nw351 = new _module.C5();
        _nw351.__ctor();
        _1840_v2 = _nw351;
        _1839_cf5 = _1836_v0;
        if (_1839_cf5) {
          let _1841_v3;
          let _nw352 = new _module.C5();
          _nw352.__ctor();
          _1841_v3 = _nw352;
          let _1842_v4;
          _1842_v4 = _dafny.Seq.UnicodeFromString("n");
          let _1843_v5;
          let _nw353 = Array((new BigNumber(8)).toNumber());
          _nw353[(_dafny.ZERO).toNumber()] = _1836_v0;
          _nw353[(_dafny.ONE).toNumber()] = _1836_v0;
          _nw353[(new BigNumber(2)).toNumber()] = _1839_cf5;
          _nw353[(new BigNumber(3)).toNumber()] = true;
          _nw353[(new BigNumber(4)).toNumber()] = _1839_cf5;
          _nw353[(new BigNumber(5)).toNumber()] = _1839_cf5;
          _nw353[(new BigNumber(6)).toNumber()] = _1839_cf5;
          _nw353[(new BigNumber(7)).toNumber()] = _1836_v0;
          _1843_v5 = _nw353;
          let _1844_v6;
          _1844_v6 = _dafny.Set.fromElements(_1843_v5);
          let _1845_v7;
          _1845_v7 = _1844_v6;
          let _1846_v8;
          _1846_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1839_cf5,_1845_v7);
          let _1847_v9;
          let _1848_v10;
          let _1849_v11;
          let _1850_v12;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector5 = (_this).m5(_1842_v4, (_1846_v8).update(!(_1836_v0), _1845_v7), _1837_v1, globalState);
          _out17 = _outcollector5[0];
          _out18 = _outcollector5[1];
          _out19 = _outcollector5[2];
          _out20 = _outcollector5[3];
          _1847_v9 = _out17;
          _1848_v10 = _out18;
          _1849_v11 = _out19;
          _1850_v12 = _out20;
          _1849_v11 = _1849_v11;
          let _1851_v13;
          _1851_v13 = _dafny.Set.fromElements(_1849_v11, true, _1849_v11);
          let _1852_v14;
          _1852_v14 = _dafny.Set.fromElements(_1851_v13, _dafny.Set.fromElements(_1849_v11, true, _1839_cf5), (_1851_v13).Union(_dafny.Set.fromElements(_1849_v11)));
          _1852_v14 = _module.__default.fm39(_dafny.Seq.Concat(_1842_v4, _1842_v4), globalState);
          let _1853_v15;
          let _init67 = ((_1854_v1) => function (_1855_i0) {
            return (_1855_i0).plus(_1854_v1);
          })(_1837_v1);
          let _nw354 = Array((new BigNumber(17)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw354.length); _i0_67++) {
            _nw354[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _1853_v15 = _nw354;
          let _1856_v16;
          _1856_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,_1853_v15);
          let _1857_v17;
          _1857_v17 = _dafny.MultiSet.fromElements(_1836_v0);
          _1856_v16 = (_1856_v16).update(_module.__default.safeDivisionInt(_1848_v10, (((_1857_v17).contains(true)) ? ((_1857_v17).get(true)) : (_1848_v10))), _1853_v15);
        } else {
          let _1858_v18;
          _1858_v18 = _dafny.Seq.UnicodeFromString("lcer");
          _1837_v1 = new BigNumber((_1858_v18).length);
          let _1859_v19;
          let _nw355 = new _module.C2();
          _nw355.__ctor((_this).f4, ((_1839_cf5) ? ((_this).f5) : (new _dafny.CodePoint('p'.codePointAt(0)))));
          _1859_v19 = _nw355;
          let _1860_v20;
          _1860_v20 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs329 = ((true) ? (!(!(_1836_v0)) || (false)) : (!(_1836_v0) || (true)));
          let _rhs330 = _1859_v19;
          let _rhs331 = _1860_v20;
          _1836_v0 = _rhs329;
          _1859_v19 = _rhs330;
          _1860_v20 = _rhs331;
          _1839_cf5 = _1839_cf5;
          let _1861_v21;
          _1861_v21 = _dafny.Seq.of(_1837_v1);
          let _1862_v22;
          _1862_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1861_v21,_1837_v1);
          _1836_v0 = !(((_1862_v22)).update(_1861_v21, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(631),!(_1839_cf5))).length))).contains(_1861_v21);
          r1 = _module.__default.safeDivisionInt(_1837_v1, _1837_v1);
        }
        let _1863_v23;
        _1863_v23 = _dafny.Set.fromElements(_1837_v1, _1837_v1, _1837_v1);
        let _1864_v24;
        _1864_v24 = _dafny.Seq.of(_1839_cf5, !(new BigNumber(635)).isEqualTo(_1837_v1), (_1863_v23).IsSubsetOf(_1863_v23));
        let _1865_v25;
        let _nw356 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1865_v25 = _nw356;
        let _index356 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1865_v25).length));
        (_1865_v25)[_index356] = _1837_v1;
        let _1866_v26;
        _1866_v26 = _dafny.Seq.UnicodeFromString("p");
        let _1867_v27;
        _1867_v27 = new _dafny.CodePoint('n'.codePointAt(0));
        let _index357 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1865_v25).length));
        let _rhs332 = _1864_v24;
        let _rhs333 = _1837_v1;
        let _rhs334 = _1837_v1;
        let _rhs335 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("nisbhej"), _module.__default.safeIndex(_1837_v1, new BigNumber((_dafny.Seq.UnicodeFromString("nisbhej")).length)), (_this).f5);
        let _rhs336 = new _dafny.CodePoint('b'.codePointAt(0));
        let _lhs171 = _1865_v25;
        let _lhs172 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1865_v25).length));
        _1864_v24 = _rhs332;
        _lhs171[_lhs172] = _rhs333;
        r1 = _rhs334;
        _1866_v26 = _rhs335;
        _1867_v27 = _rhs336;
      } else if (_source18.is_DC5) {
        let _1868___mcc_h1 = (_source18).cf6;
        let _1869_cf6 = _1868___mcc_h1;
        _1836_v0 = true;
        _1836_v0 = _1836_v0;
        let _1870_v28;
        _1870_v28 = _module.D20.create_DC47((_this).f4);
        let _1871_v29;
        let _nw357 = new _module.C3();
        _nw357.__ctor((_1870_v28).dtor_cf73, (_this).f5);
        _1871_v29 = _nw357;
        if (_1836_v0) {
          r1 = _1837_v1;
          let _1872_v30;
          _1872_v30 = _dafny.MultiSet.fromElements(_1836_v0);
          let _1873_v31;
          _1873_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,(_module.__default.fm15(_1836_v0, _1836_v0, _1837_v1, globalState)).equals(_1872_v30));
          _1873_v31 = (_1873_v31).update(_1837_v1, _1836_v0);
          let _1874_v32;
          _1874_v32 = _dafny.Seq.of(_1836_v0, _1836_v0, !(true), _1836_v0);
          let _1875_v33;
          _1875_v33 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1874_v32), (_1872_v30).Union(_1872_v30));
          _1875_v33 = _dafny.Seq.Concat(_dafny.Seq.of(_1872_v30, _1872_v30), _1875_v33);
          _1874_v32 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1874_v32, _1874_v32), _1874_v32), _module.__default.safeIndex((_dafny.ZERO).minus((_1837_v1).multipliedBy(_1837_v1)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1874_v32, _1874_v32), _1874_v32)).length)), (_1836_v0) === (_1836_v0));
          let _1876_v34;
          let _nw358 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1876_v34 = _nw358;
          let _1877_v35;
          _1877_v35 = _dafny.Set.fromElements(_1876_v34, _1876_v34, _1876_v34);
          let _1878_v36;
          _1878_v36 = _dafny.Seq.of(_1837_v1);
          let _index358 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1876_v34).length));
          (_1876_v34)[_index358] = (_1837_v1).isEqualTo((_this).fm4(_1878_v36, globalState));
          let _1879_v37;
          let _init68 = ((_1880_v1) => function (_1881_i1) {
            return (_1881_i1).multipliedBy(_1880_v1);
          })(_1837_v1);
          let _nw359 = Array((new BigNumber(10)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw359.length); _i0_68++) {
            _nw359[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _1879_v37 = _nw359;
          let _1882_v38;
          _1882_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1877_v35);
          let _index359 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1876_v34).length));
          let _rhs337 = ((((_1882_v38).contains(true)) ? ((_1882_v38).get(true)) : (_1877_v35))).Union(_1877_v35);
          let _rhs338 = _1836_v0;
          let _rhs339 = _1869_cf6;
          let _rhs340 = _1879_v37;
          let _rhs341 = _1836_v0;
          let _lhs173 = _1876_v34;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1876_v34).length));
          _1877_v35 = _rhs337;
          _lhs173[_lhs174] = _rhs338;
          _1869_cf6 = _rhs339;
          _1879_v37 = _rhs340;
          _1836_v0 = _rhs341;
        } else {
          r1 = _module.__default.safeDivisionInt(_1837_v1, _1837_v1);
          let _1883_v39;
          _1883_v39 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1837_v1));
          let _1884_v40;
          _1884_v40 = _dafny.Set.fromElements(_1883_v39, _1883_v39);
          let _1885_v41;
          _1885_v41 = _dafny.Seq.of(_1837_v1, new BigNumber((_1884_v40).length), _1837_v1, _module.__default.fm13(false, _1836_v0, _1836_v0, globalState));
          let _1886_v42;
          _1886_v42 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(960)), ((_1887_v1) => function (_1888_i2) {
            return (_module.D14.create_DC31(_1887_v1)).dtor_cf48;
          })(_1837_v1)), _1885_v41, _dafny.Seq.Concat(_1885_v41, _1885_v41));
          _1886_v42 = (_1886_v42).Difference(_1886_v42);
          let _1889_v43;
          _1889_v43 = _dafny.MultiSet.fromElements(_1836_v0);
          _1836_v0 = !((_1889_v43).IsSubsetOf((_1889_v43).Difference(_1889_v43)));
          let _1890_v44;
          _1890_v44 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1836_v0, _1836_v0, _1836_v0),_1837_v1);
          let _1891_v45;
          _1891_v45 = _module.D14.create_DC32(_1837_v1, _1836_v0, _1836_v0);
          let _1892_v46;
          _1892_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1891_v45,_dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), function (_1893_i3) {
            return (_this).f5;
          }));
          let _1894_v48;
          _1894_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1891_v45,_dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1836_v0));
          let _1895_v49;
          _1895_v49 = _module.D21.create_DC50(_1894_v48);
          let _rhs342 = _module.__default.fm40(_1837_v1, globalState);
          let _rhs343 = ((!(_1837_v1).isEqualTo(_1837_v1)) ? (function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of ((_1895_v49).dtor_cf79).Keys.Elements) {
              let _1896_v47 = _compr_56;
              if (((_1895_v49).dtor_cf79).contains(_1896_v47)) {
                _coll56.push([_1896_v47,_1869_cf6]);
              }
            }
            return _coll56;
          }()) : (_dafny.Map.Empty.slice().updateUnsafe(_1891_v45,_1869_cf6)));
          let _rhs344 = _dafny.Seq.Concat(_1869_cf6, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1869_cf6, _module.__default.safeIndex(_1837_v1, new BigNumber((_1869_cf6).length)), (_this).f5), _1869_cf6), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1889_v43).cardinality())), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1869_cf6, _module.__default.safeIndex(_1837_v1, new BigNumber((_1869_cf6).length)), (_this).f5), _1869_cf6)).length)), new _dafny.CodePoint('u'.codePointAt(0))));
          let _rhs345 = _1837_v1;
          let _rhs346 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), function (_1897_i4) {
            return (_this).f5;
          });
          _1890_v44 = _rhs342;
          _1892_v46 = _rhs343;
          _1869_cf6 = _rhs344;
          r1 = _rhs345;
          _1869_cf6 = _rhs346;
          r1 = _1837_v1;
        }
      } else if (_source18.is_DC6) {
        let _1898___mcc_h2 = (_source18).cf7;
        let _1899_cf7 = _1898___mcc_h2;
        let _1900_v50;
        _1900_v50 = _dafny.MultiSet.fromElements(_1836_v0);
        let _1901_v51;
        _1901_v51 = _dafny.Set.fromElements(_1837_v1, new BigNumber((_1900_v50).cardinality()));
        let _1902_v52;
        _1902_v52 = _module.D16.create_DC37(_1901_v51);
        let _1903_v53;
        _1903_v53 = _module.D16.create_DC39(_module.D16.create_DC39(_1902_v52));
        let _source19 = _1903_v53;
        if (_source19.is_DC38) {
          let _1904___mcc_h5 = (_source19).cf57;
          let _1905___mcc_h6 = (_source19).cf58;
          let _1906___mcc_h7 = (_source19).cf59;
          let _1907___mcc_h8 = (_source19).cf60;
          let _1908_cf60 = _1907___mcc_h8;
          let _1909_cf59 = _1906___mcc_h7;
          let _1910_cf58 = _1905___mcc_h6;
          let _1911_cf57 = _1904___mcc_h5;
          let _1912_v54;
          _1912_v54 = _dafny.Seq.of(_1837_v1, _1909_cf59, _1837_v1);
          let _nw360 = Array((new BigNumber(4)).toNumber());
          _nw360[(_dafny.ZERO).toNumber()] = _1910_cf58;
          _nw360[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_1912_v54, _1899_cf7);
          _nw360[(new BigNumber(2)).toNumber()] = _1836_v0;
          _nw360[(new BigNumber(3)).toNumber()] = ((_1908_cf60) ? (_module.__default.fm1(_1899_cf7, _1899_cf7, true, _1910_cf58, globalState)) : (true));
          (globalState).f3 = _nw360;
          let _1913_v55;
          _1913_v55 = _dafny.MultiSet.fromElements(((true) ? (_1911_cf57) : (_1911_cf57)), _dafny.Seq.Concat(_1911_cf57, _1911_cf57));
          r1 = (((_1913_v55).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_1915_i5) {
            return (_this).f5;
          }))) ? ((_1913_v55).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_1914_i5) {
            return (_this).f5;
          }))) : ((new BigNumber(-840)).minus(_1899_cf7)));
          let _1916_v56;
          let _nw361 = new _module.C2();
          _nw361.__ctor((_this).f4, (_this).f5);
          _1916_v56 = _nw361;
          _1916_v56 = _1916_v56;
          let _1917_v57;
          let _init69 = ((_1918_v0) => function (_1919_i6) {
            return _1918_v0;
          })(_1836_v0);
          let _nw362 = Array((_dafny.ONE).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw362.length); _i0_69++) {
            _nw362[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _1917_v57 = _nw362;
          let _index360 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1917_v57).length));
          (_1917_v57)[_index360] = _1908_cf60;
          let _index361 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1917_v57).length));
          (_1917_v57)[_index361] = _1908_cf60;
        } else if (_source19.is_DC37) {
          let _1920___mcc_h9 = (_source19).cf56;
          let _1921_cf56 = _1920___mcc_h9;
          let _1922_v58;
          let _nw363 = new _module.C0();
          _nw363.__ctor();
          _1922_v58 = _nw363;
          let _1923_v59;
          _1923_v59 = _dafny.Seq.of(_1922_v58, _1922_v58, _1922_v58, _1922_v58);
          _1923_v59 = _dafny.Seq.of(_1922_v58);
          let _1924_v60;
          _1924_v60 = new _dafny.CodePoint('n'.codePointAt(0));
          _1924_v60 = (_this).f5;
          _1836_v0 = !((_1899_cf7).isEqualTo(_1837_v1)) || (!(_1836_v0) || (_1836_v0));
          let _1925_v61;
          _1925_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1837_v1),_1924_v60);
          let _1926_v62;
          _1926_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1899_cf7, _1837_v1, globalState),_1925_v61);
          let _1927_v63;
          let _nw364 = Array((new BigNumber(4)).toNumber());
          _nw364[(_dafny.ZERO).toNumber()] = _1925_v61;
          _nw364[(_dafny.ONE).toNumber()] = _1925_v61;
          _nw364[(new BigNumber(2)).toNumber()] = (_1925_v61).Merge(_1925_v61);
          _nw364[(new BigNumber(3)).toNumber()] = (_1925_v61).Merge((((_1926_v62).contains(_1924_v60)) ? ((_1926_v62).get(_1924_v60)) : (_1925_v61)));
          _1927_v63 = _nw364;
          _1927_v63 = _1927_v63;
        } else {
          let _1928___mcc_h10 = (_source19).cf61;
          let _1929_cf61 = _1928___mcc_h10;
          _1899_cf7 = (_module.__default.fm41(_1836_v0, _1837_v1, _1899_cf7, _1837_v1, globalState)).dtor_cf59;
          let _1930_v64;
          let _nw365 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1930_v64 = _nw365;
          let _1931_v65;
          _1931_v65 = _dafny.MultiSet.fromElements(_1899_cf7, _1837_v1);
          let _1932_v66;
          _1932_v66 = _dafny.Seq.of(_1837_v1, new BigNumber((_1931_v65).cardinality()));
          let _index362 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_1930_v64).length));
          (_1930_v64)[_index362] = (_this).fm4(_1932_v66, globalState);
          let _index363 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_1930_v64).length));
          (_1930_v64)[_index363] = _1899_cf7;
          let _1933_v67;
          _1933_v67 = _dafny.Seq.UnicodeFromString("milgoae");
          _1933_v67 = _dafny.Seq.Concat(_1933_v67, _1933_v67);
          let _1934_v68;
          _1934_v68 = _1836_v0;
          let _1935_v69;
          _1935_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,(_1930_v64)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_1930_v64).length))]);
          let _1936_v70;
          let _nw366 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1936_v70 = _nw366;
          let _1937_v71;
          _1937_v71 = _module.D18.create_DC45(_1936_v70, _1836_v0, _1837_v1, new BigNumber(474));
          _1934_v68 = _module.__default.fm42((_1930_v64)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_1930_v64).length))], (((_1935_v69).contains(_1836_v0)) ? ((_1935_v69).get(_1836_v0)) : (_1899_cf7)), _1836_v0, (_1937_v71).dtor_cf70, globalState);
        }
        let _1938_v72;
        _1938_v72 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber(-467)));
        let _1939_v73;
        let _nw367 = new _module.C4();
        _nw367.__ctor(new BigNumber((_dafny.Seq.of((_this).f5)).length), new BigNumber((_1938_v72).length), (_this).f4, (_this).f5);
        _1939_v73 = _nw367;
        (_1939_v73).f11 = new BigNumber(-506);
        _1836_v0 = !(_1836_v0) || (_1836_v0);
      } else if (_source18.is_DC3) {
        let _1940___mcc_h3 = (_source18).cf4;
        let _1941_cf4 = _1940___mcc_h3;
        let _1942_v74;
        _1942_v74 = new _dafny.CodePoint('g'.codePointAt(0));
        _1942_v74 = _1942_v74;
        let _1943_v75;
        _1943_v75 = _dafny.Set.fromElements(_1836_v0, _1836_v0, _1836_v0);
        _1837_v1 = _module.__default.fm13(_1836_v0, _1836_v0, (_1837_v1).isEqualTo(new BigNumber((_1943_v75).length)), globalState);
        let _1944_v76;
        _1944_v76 = _dafny.Seq.UnicodeFromString("hh");
        let _1945_v77;
        _1945_v77 = _module.D7.create_DC16(_1942_v74);
        let _1946_v78;
        let _init70 = ((_1947_v0) => function (_1948_i7) {
          return _dafny.Seq.of(_1947_v0);
        })(_1836_v0);
        let _nw368 = Array((new BigNumber(2)).toNumber());
        for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw368.length); _i0_70++) {
          _nw368[_i0_70] = _init70(new BigNumber(_i0_70));
        }
        _1946_v78 = _nw368;
        let _1949_v79;
        _1949_v79 = _module.D4.create_DC9(_1946_v78);
        let _1950_v80;
        _1950_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1945_v77,_1949_v79);
        let _1951_v81;
        _1951_v81 = _dafny.Map.Empty.slice().updateUnsafe((_1950_v80).update(_1945_v77, _1949_v79),(_dafny.ZERO).minus(_1837_v1));
        let _rhs347 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cyyph"), _1944_v76);
        let _rhs348 = new _dafny.CodePoint('p'.codePointAt(0));
        let _rhs349 = _dafny.Seq.of(_module.__default.fm13(_1836_v0, _1836_v0, _1836_v0, globalState), (((_1951_v81).contains(_1950_v80)) ? ((_1951_v81).get(_1950_v80)) : (_1837_v1)), _1837_v1);
        _1944_v76 = _rhs347;
        _1942_v74 = _rhs348;
        r0 = _rhs349;
        let _1952_v82;
        let _nw369 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1952_v82 = _nw369;
        let _index364 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1952_v82).length));
        (_1952_v82)[_index364] = _1837_v1;
        let _index365 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1952_v82).length));
        (_1952_v82)[_index365] = _module.__default.safeDivisionInt(_1837_v1, _module.__default.fm14(globalState));
      } else {
        let _1953___mcc_h4 = (_source18).cf8;
        let _1954_cf8 = _1953___mcc_h4;
        _1836_v0 = true;
        let _1955_v83;
        let _nw370 = new _module.C3();
        _nw370.__ctor((_module.D20.create_DC47((_this).f4)).dtor_cf73, new _dafny.CodePoint('s'.codePointAt(0)));
        _1955_v83 = _nw370;
        let _1956_v84;
        _1956_v84 = _dafny.Seq.UnicodeFromString("enprlnjp");
        let _1957_v85;
        _1957_v85 = _dafny.Seq.of(_1837_v1);
        let _1958_v86;
        _1958_v86 = _module.D14.create_DC32(_1837_v1, _1836_v0, _1836_v0);
        let _1959_v87;
        _1959_v87 = _dafny.Seq.of((_dafny.ZERO).minus((_1957_v85)[_module.__default.safeIndex((_dafny.ZERO).minus((_1958_v86).dtor_cf49), new BigNumber((_1957_v85).length))]));
        _1956_v84 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("l"), _module.__default.safeIndex((_1959_v87)[_module.__default.safeIndex(_1837_v1, new BigNumber((_1959_v87).length))], new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)), (_this).f5);
        _1959_v87 = _dafny.Seq.update(_1957_v85, _module.__default.safeIndex(_1837_v1, new BigNumber((_1957_v85).length)), _1837_v1);
      }
      let _hi16 = (_1837_v1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("lrob")).length));
      for (let _1960_i8 = (_dafny.ZERO).minus(_1837_v1); _1960_i8.isLessThan(_hi16); _1960_i8 = _1960_i8.plus(_dafny.ONE)) {
        let _1961_v88;
        _1961_v88 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1960_i8),new _dafny.CodePoint('b'.codePointAt(0)));
        _1961_v88 = (_1961_v88).update(_1960_i8, (_this).f5);
        let _1962_v89;
        _1962_v89 = _dafny.Seq.UnicodeFromString("eh");
        let _rhs350 = _1962_v89;
        let _rhs351 = !(_1960_i8).isEqualTo(_1960_i8);
        _1962_v89 = _rhs350;
        _1836_v0 = _rhs351;
        let _1963_v90;
        _1963_v90 = _dafny.Set.fromElements(_1836_v0, _1836_v0);
        let _1964_v91;
        _1964_v91 = _dafny.Seq.of(_1836_v0, _1836_v0, _1836_v0, _1836_v0, !((_dafny.Set.fromElements(_1836_v0)).IsSubsetOf(_1963_v90)));
        let _1965_v92;
        _1965_v92 = _dafny.Seq.of(_1960_i8);
        _1836_v0 = (_1964_v91)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1965_v92)).cardinality()), new BigNumber((_1964_v91).length))];
        let _1966_v93;
        _1966_v93 = _dafny.MultiSet.fromElements((_this).f5);
        let _1967_v94;
        _1967_v94 = _dafny.MultiSet.fromElements(_1836_v0, _1836_v0, _1836_v0, (_1966_v93).IsSubsetOf(_1966_v93), _1836_v0);
        _1967_v94 = _1967_v94;
      }
      let _1968_v95;
      _1968_v95 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_1836_v0)).length), _1837_v1);
      let _1969_v96;
      _1969_v96 = _dafny.Seq.of(_1968_v95);
      let _1970_v97;
      _1970_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,_1968_v95);
      let _1971_v98;
      _1971_v98 = _dafny.Seq.of(!(new BigNumber((_1969_v96).length)).isEqualTo(new BigNumber((_1970_v97).length)));
      _1836_v0 = (_1971_v98)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1837_v1, (_dafny.ZERO).minus(_1837_v1)), new BigNumber((_1971_v98).length))];
      if (_1836_v0) {
        if (!(_1836_v0) || (_1836_v0)) {
          let _1972_v99;
          _1972_v99 = _dafny.MultiSet.fromElements(_1837_v1);
          let _1973_v100;
          _1973_v100 = _dafny.Seq.of(_1972_v99);
          let _1974_v101;
          _1974_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(_1837_v1, _1836_v0, (_dafny.ZERO).minus(_1837_v1), _1837_v1, globalState),new BigNumber(((_1973_v100)[_module.__default.safeIndex(_1837_v1, new BigNumber((_1973_v100).length))]).cardinality()));
          let _1975_v102;
          let _nw371 = Array((_dafny.ONE).toNumber()).fill(false);
          _1975_v102 = _nw371;
          let _index366 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1975_v102).length));
          (_1975_v102)[_index366] = _1836_v0;
          let _1976_v103;
          _1976_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,new BigNumber(540));
          let _1977_v104;
          _1977_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1836_v0);
          let _1978_v105;
          _1978_v105 = _dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), (_this).f5, (_this).f5);
          let _1979_v106;
          let _nw372 = new _module.C2();
          _nw372.__ctor((_this).f4, (_1978_v105)[_module.__default.safeIndex(_1837_v1, new BigNumber((_1978_v105).length))]);
          _1979_v106 = _nw372;
          let _1980_v107;
          _1980_v107 = _dafny.Seq.of(_1979_v106);
          let _1981_v108;
          _1981_v108 = _dafny.Set.fromElements(_1836_v0);
          let _1982_v109;
          let _nw373 = Array((new BigNumber(24)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = new BigNumber(613);
          _nw373[(_dafny.ONE).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(2)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(3)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(4)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(5)).toNumber()] = new BigNumber((_1976_v103).length);
          _nw373[(new BigNumber(6)).toNumber()] = new BigNumber((_1977_v104).length);
          _nw373[(new BigNumber(7)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(8)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(9)).toNumber()] = new BigNumber((_1980_v107).length);
          _nw373[(new BigNumber(10)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(11)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(12)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(13)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(14)).toNumber()] = new BigNumber((_1981_v108).length);
          _nw373[(new BigNumber(15)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(16)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(17)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(18)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(19)).toNumber()] = new BigNumber((_1978_v105).length);
          _nw373[(new BigNumber(20)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(21)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(22)).toNumber()] = _1837_v1;
          _nw373[(new BigNumber(23)).toNumber()] = _1837_v1;
          _1982_v109 = _nw373;
          let _1983_v110;
          _1983_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1982_v109,_1837_v1);
          let _index367 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1975_v102).length));
          let _rhs352 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1837_v1);
          let _rhs353 = (_1837_v1).isLessThanOrEqualTo(_1837_v1);
          let _rhs354 = _1836_v0;
          let _rhs355 = (_1837_v1).isLessThan(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1982_v109,_1837_v1)).Merge(_1983_v110)).length));
          let _lhs175 = _1975_v102;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1975_v102).length));
          _1974_v101 = _rhs352;
          _lhs175[_lhs176] = _rhs353;
          _1836_v0 = _rhs354;
          _1836_v0 = _rhs355;
          _1837_v1 = (_this).fm4(_dafny.Seq.of((_module.D1.create_DC2(new BigNumber(-871), _1837_v1)).dtor_cf2, (_dafny.ZERO).minus(_1837_v1)), globalState);
          _1837_v1 = (_1837_v1).minus((_1837_v1).minus(new BigNumber((_module.__default.fm34(false, globalState)).length)));
          let _1984_v111;
          let _nw374 = new _module.C1();
          _nw374.__ctor((_this).f4, (_1979_v106).f5);
          _1984_v111 = _nw374;
          _1836_v0 = (_1971_v98)[_module.__default.safeIndex((_module.__default.fm14(globalState)).multipliedBy(_1837_v1), new BigNumber((_1971_v98).length))];
        } else {
          let _1985_v112;
          let _nw375 = new _module.C7();
          _nw375.__ctor();
          _1985_v112 = _nw375;
          let _rhs356 = _1985_v112;
          let _rhs357 = _1837_v1;
          let _rhs358 = _1836_v0;
          _1985_v112 = _rhs356;
          r1 = _rhs357;
          _1836_v0 = _rhs358;
          let _1986_v113;
          let _nw376 = new _module.C2();
          _nw376.__ctor((_this).f4, (_this).f5);
          _1986_v113 = _nw376;
          let _1987_v114;
          _1987_v114 = _dafny.Set.fromElements(_1836_v0, _1836_v0);
          let _1988_v115;
          _1988_v115 = _module.D10.create_DC21(_1987_v114);
          let _1989_v116;
          _1989_v116 = _module.D10.create_DC23(_1988_v115);
          let _1990_v117;
          let _nw377 = Array((new BigNumber(10)).toNumber()).fill(false);
          _1990_v117 = _nw377;
          let _1991_v118;
          _1991_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1990_v117,_1836_v0);
          let _1992_v119;
          _1992_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1989_v116,!(_1991_v118).contains(_1990_v117));
          _1992_v119 = (_1992_v119).update(_1989_v116, _1836_v0);
          let _1993_v120;
          let _nw378 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1993_v120 = _nw378;
          let _1994_v121;
          let _init71 = function (_1995_i9) {
            return _module.__default.safeModuloInt(_1995_i9, new BigNumber(-59));
          };
          let _nw379 = Array((new BigNumber(2)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw379.length); _i0_71++) {
            _nw379[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _1994_v121 = _nw379;
          let _index368 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1993_v120).length));
          (_1993_v120)[_index368] = _1994_v121;
          let _1996_v122;
          _1996_v122 = _dafny.Seq.of(_1994_v121, _1994_v121);
          let _index369 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1993_v120).length));
          let _rhs359 = (_1996_v122)[_module.__default.safeIndex(new BigNumber(261), new BigNumber((_1996_v122).length))];
          let _rhs360 = _1837_v1;
          let _rhs361 = ((_1836_v0) ? (_1837_v1) : ((_1837_v1).plus(_1837_v1)));
          let _lhs177 = _1993_v120;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_1993_v120).length));
          _lhs177[_lhs178] = _rhs359;
          r1 = _rhs360;
          _1837_v1 = _rhs361;
          let _1997_v123;
          _1997_v123 = _dafny.Seq.UnicodeFromString("evcvcc");
          let _1998_v124;
          _1998_v124 = _module.D16.create_DC38(((_1836_v0) ? (_1997_v123) : (_dafny.Seq.UnicodeFromString("nbw"))), _1836_v0, new BigNumber(797), _1836_v0);
          let _1999_v125;
          _1999_v125 = new _dafny.CodePoint('d'.codePointAt(0));
          let _rhs362 = _1998_v124;
          let _rhs363 = (_this).f5;
          _1998_v124 = _rhs362;
          _1999_v125 = _rhs363;
        }
        let _2000_v126;
        _2000_v126 = _dafny.Seq.UnicodeFromString("cyqox");
        let _2001_v127;
        _2001_v127 = _module.D16.create_DC38(_2000_v126, _1836_v0, _1837_v1, _1836_v0);
        let _2002_v128;
        _2002_v128 = _module.D7.create_DC17((_2001_v127).dtor_cf58, _1836_v0, true, _1836_v0);
        let _pat_let_tv49 = _1836_v0;
        let _2003_v129;
        _2003_v129 = _module.D7.create_DC17((_this).fm3(new BigNumber((_2000_v126).length), new BigNumber(261), globalState), (function (_pat_let46_0) {
  return function (_2004_dt__update__tmp_h0) {
    return function (_pat_let47_0) {
      return function (_2005_dt__update_hcf26_h0) {
        return function (_pat_let48_0) {
          return function (_2006_dt__update_hcf25_h0) {
            return _module.D7.create_DC17((_2004_dt__update__tmp_h0).dtor_cf24, _2006_dt__update_hcf25_h0, _2005_dt__update_hcf26_h0, (_2004_dt__update__tmp_h0).dtor_cf27);
          }(_pat_let48_0);
        }(true);
      }(_pat_let47_0);
    }(_pat_let_tv49);
  }(_pat_let46_0);
}(_2002_v128)).dtor_cf24, _1836_v0, _1836_v0);
        let _rhs364 = true;
        let _rhs365 = _module.D7.create_DC17(_1836_v0, _1836_v0, _1836_v0, _module.__default.fm1(_1837_v1, new BigNumber((_2000_v126).length), _1836_v0, _1836_v0, globalState));
        _1836_v0 = _rhs364;
        _2003_v129 = _rhs365;
        let _2007_v130;
        let _nw380 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _2007_v130 = _nw380;
        let _index370 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2007_v130).length));
        (_2007_v130)[_index370] = _1837_v1;
        let _index371 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2007_v130).length));
        (_2007_v130)[_index371] = _module.__default.safeDivisionInt((_1837_v1).multipliedBy(_1837_v1), (new BigNumber((_dafny.Seq.UnicodeFromString("qgucd")).length)).plus(_1837_v1));
        let _2008_v131;
        _2008_v131 = _dafny.Map.Empty.slice().updateUnsafe((_1837_v1).minus((_dafny.ZERO).minus(_1837_v1)),!(_1837_v1).isEqualTo((_2007_v130)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_2007_v130).length))]));
        _1836_v0 = (((_2008_v131).contains(_1837_v1)) ? ((_2008_v131).get(_1837_v1)) : (_1836_v0));
        let _2009_v132;
        let _nw381 = new _module.C2();
        _nw381.__ctor((_this).f4, (_this).f5);
        _2009_v132 = _nw381;
      } else {
        r1 = _1837_v1;
        let _2010_v133;
        _2010_v133 = _dafny.Seq.UnicodeFromString("lmep");
        let _2011_v134;
        let _nw382 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _2011_v134 = _nw382;
        let _2012_v135;
        _2012_v135 = _module.D11.create_DC25(_1836_v0, new BigNumber((_2010_v133).length), _2011_v134, new BigNumber((_dafny.Set.fromElements(_1837_v1, new BigNumber(-789))).length));
        let _2013_v136;
        _2013_v136 = _module.D11.create_DC26(_2012_v135);
        let _pat_let_tv50 = _2012_v135;
        let _source20 = (((new BigNumber(798)).isLessThanOrEqualTo(_1837_v1)) ? (function (_pat_let49_0) {
          return function (_2014_dt__update__tmp_h1) {
            return function (_pat_let50_0) {
              return function (_2015_dt__update_hcf42_h0) {
                return _module.D11.create_DC26(_2015_dt__update_hcf42_h0);
              }(_pat_let50_0);
            }(_module.D11.create_DC26(_pat_let_tv50));
          }(_pat_let49_0);
        }(_2013_v136)) : (_2013_v136));
        if (_source20.is_DC25) {
          let _2016___mcc_h11 = (_source20).cf38;
          let _2017___mcc_h12 = (_source20).cf39;
          let _2018___mcc_h13 = (_source20).cf40;
          let _2019___mcc_h14 = (_source20).cf41;
          let _2020_cf41 = _2019___mcc_h14;
          let _2021_cf40 = _2018___mcc_h13;
          let _2022_cf39 = _2017___mcc_h12;
          let _2023_cf38 = _2016___mcc_h11;
          let _2024_v137;
          _2024_v137 = _dafny.Set.fromElements(_2023_cf38, _2023_cf38, _1836_v0, _1836_v0, true);
          _1971_v98 = _dafny.Seq.Concat(_module.__default.fm21(_2022_cf39, _2020_cf41, _2024_v137, globalState), _1971_v98);
          let _2025_v138;
          let _nw383 = new _module.C1();
          _nw383.__ctor((_this).f4, (_this).f5);
          _2025_v138 = _nw383;
          _2025_v138 = _2025_v138;
          _1836_v0 = _1836_v0;
          _2023_cf38 = (_2023_cf38) === (_1836_v0);
        } else if (_source20.is_DC24) {
          let _2026___mcc_h15 = (_source20).cf37;
          let _2027_cf37 = _2026___mcc_h15;
          let _2028_v139;
          _2028_v139 = _module.D16.create_DC38(_2010_v133, _1836_v0, (_dafny.ZERO).minus(_1837_v1), _1836_v0);
          let _2029_v140;
          _2029_v140 = _dafny.MultiSet.fromElements(_1837_v1);
          let _pat_let_tv51 = _1971_v98;
          let _pat_let_tv52 = _2029_v140;
          let _pat_let_tv53 = _1971_v98;
          _1836_v0 = (function (_pat_let51_0) {
            return function (_2030_dt__update__tmp_h2) {
              return function (_pat_let52_0) {
                return function (_2031_dt__update_hcf58_h0) {
                  return _module.D16.create_DC38((_2030_dt__update__tmp_h2).dtor_cf57, _2031_dt__update_hcf58_h0, (_2030_dt__update__tmp_h2).dtor_cf59, (_2030_dt__update__tmp_h2).dtor_cf60);
                }(_pat_let52_0);
              }((_pat_let_tv51)[_module.__default.safeIndex(new BigNumber((_pat_let_tv52).cardinality()), new BigNumber((_pat_let_tv53).length))]);
            }(_pat_let51_0);
          }(_2028_v139)).dtor_cf58;
          let _2032_v141;
          let _nw384 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2032_v141 = _nw384;
          _2032_v141 = _2032_v141;
          let _2033_v142;
          _2033_v142 = new _dafny.CodePoint('l'.codePointAt(0));
          _2033_v142 = _2033_v142;
          let _2034_v143;
          _2034_v143 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm14(globalState),new BigNumber(666));
          _1836_v0 = ((new BigNumber((_2034_v143).length)).multipliedBy(_module.__default.fm13(_1836_v0, true, _1836_v0, globalState))).isEqualTo(_1837_v1);
        } else {
          let _2035___mcc_h16 = (_source20).cf42;
          let _2036_cf42 = _2035___mcc_h16;
          let _index372 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_2011_v134).length));
          (_2011_v134)[_index372] = _1837_v1;
          let _2037_v144;
          _2037_v144 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1837_v1);
          let _index373 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_2011_v134).length));
          (_2011_v134)[_index373] = (((_2037_v144).contains((_this).f5)) ? ((_2037_v144).get((_this).f5)) : (_module.__default.safeModuloInt(new BigNumber(-855), new BigNumber((_2010_v133).length))));
          let _2038_v145;
          let _init72 = ((_2039_v134) => function (_2040_i10) {
            return (_dafny.MultiSet.fromElements((_2039_v134)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_2039_v134).length))])).Union(_dafny.MultiSet.fromElements(new BigNumber(696), (_2039_v134)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_2039_v134).length))]));
          })(_2011_v134);
          let _nw385 = Array((new BigNumber(23)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw385.length); _i0_72++) {
            _nw385[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2038_v145 = _nw385;
          _2038_v145 = _2038_v145;
          let _2041_v146;
          _2041_v146 = _dafny.Set.fromElements(_1837_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-678)), ((_2042_v134) => function (_2043_i11) {
            return (_2042_v134)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_2042_v134).length))];
          })(_2011_v134))).length), _module.__default.safeDivisionInt(_1837_v1, _1837_v1), (_2011_v134)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_2011_v134).length))]);
          _2041_v146 = _2041_v146;
          let _2044_v147;
          _2044_v147 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1836_v0);
          _1836_v0 = (((_2044_v147).contains(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), function (_2046_i12) {
            return (_this).f5;
          }), _2010_v133))) ? ((_2044_v147).get(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), function (_2045_i12) {
            return (_this).f5;
          }), _2010_v133))) : (_1836_v0));
        }
        let _2047_v148;
        let _init73 = ((_2048_v1, _2049_v0) => function (_2050_i13) {
          return (_module.D14.create_DC32(_2048_v1, _2049_v0, _2049_v0)).dtor_cf50;
        })(_1837_v1, _1836_v0);
        let _nw386 = Array((new BigNumber(16)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw386.length); _i0_73++) {
          _nw386[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2047_v148 = _nw386;
        let _index374 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_2047_v148).length));
        (_2047_v148)[_index374] = false;
        let _index375 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_2047_v148).length));
        (_2047_v148)[_index375] = !(_module.__default.fm14(globalState)).isEqualTo(_1837_v1);
        let _2051_v149;
        _2051_v149 = _dafny.Map.Empty.slice().updateUnsafe(_2010_v133,(_2047_v148)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_2047_v148).length))]);
        _2051_v149 = (_2051_v149).update(_2010_v133, (_2047_v148)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_2047_v148).length))]);
        let _2052_v150;
        _2052_v150 = _dafny.MultiSet.fromElements((_2047_v148)[_module.__default.safeIndex(new BigNumber(417), new BigNumber((_2047_v148).length))]);
        let _2053_v151;
        _2053_v151 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2052_v150).cardinality()),_1837_v1);
        let _2054_v153;
        let _nw387 = new _module.C6();
        _nw387.__ctor((_2053_v151).Merge(function () {
          let _coll57 = new _dafny.Map();
          for (const _compr_57 of _dafny.IntegerRange(new BigNumber(221), new BigNumber(62))) {
            let _2055_v152 = _compr_57;
            if (((new BigNumber(221)).isLessThanOrEqualTo(_2055_v152)) && ((_2055_v152).isLessThan(new BigNumber(62)))) {
              _coll57.push([(_2055_v152).plus(new BigNumber((_1971_v98).length)),new BigNumber(6)]);
            }
          }
          return _coll57;
        }()));
        _2054_v153 = _nw387;
      }
      let _2056_v154;
      let _nw388 = Array((new BigNumber(4)).toNumber());
      _nw388[(_dafny.ZERO).toNumber()] = _1836_v0;
      _nw388[(_dafny.ONE).toNumber()] = _1836_v0;
      _nw388[(new BigNumber(2)).toNumber()] = !(false);
      _nw388[(new BigNumber(3)).toNumber()] = _1836_v0;
      _2056_v154 = _nw388;
      let _2057_v155;
      _2057_v155 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), ((_2058_v1) => function (_2059_i14) {
        return _2058_v1;
      })(_1837_v1)));
      let _2060_v156;
      _2060_v156 = _dafny.Seq.UnicodeFromString("rtcajlcn");
      let _2061_v157;
      _2061_v157 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_1837_v1, new BigNumber(((((_2057_v155).contains(_1836_v0)) ? ((_2057_v155).get(_1836_v0)) : (_1968_v95))).length), globalState),new BigNumber((_2060_v156).length));
      let _2062_v158;
      _2062_v158 = _dafny.Seq.of(_2061_v157, _2061_v157, _2061_v157, _2061_v157);
      let _2063_v159;
      _2063_v159 = _module.D17.create_DC42(_2056_v154, new BigNumber((_2062_v158).length));
      let _source21 = _2063_v159;
      if (_source21.is_DC41) {
        let _2064___mcc_h17 = (_source21).cf63;
        let _2065_cf63 = _2064___mcc_h17;
        let _2066_v160;
        _2066_v160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2065_cf63);
        let _2067_v161;
        _2067_v161 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_2065_cf63)).length), _2065_cf63);
        let _2068_v162;
        let _nw389 = Array((new BigNumber(19)).toNumber());
        _nw389[(_dafny.ZERO).toNumber()] = _1837_v1;
        _nw389[(_dafny.ONE).toNumber()] = _1837_v1;
        _nw389[(new BigNumber(2)).toNumber()] = (((_2066_v160).contains((_this).f5)) ? ((_2066_v160).get((_this).f5)) : (new BigNumber((_module.__default.fm16(false, new BigNumber((_2067_v161).cardinality()), globalState)).length)));
        _nw389[(new BigNumber(3)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(4)).toNumber()] = _1837_v1;
        _nw389[(new BigNumber(5)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(6)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(7)).toNumber()] = (((_2066_v160).contains(new _dafny.CodePoint('o'.codePointAt(0)))) ? ((_2066_v160).get(new _dafny.CodePoint('o'.codePointAt(0)))) : (_2065_cf63));
        _nw389[(new BigNumber(8)).toNumber()] = new BigNumber(712);
        _nw389[(new BigNumber(9)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(10)).toNumber()] = _1837_v1;
        _nw389[(new BigNumber(11)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(12)).toNumber()] = _1837_v1;
        _nw389[(new BigNumber(13)).toNumber()] = new BigNumber(-138);
        _nw389[(new BigNumber(14)).toNumber()] = _1837_v1;
        _nw389[(new BigNumber(15)).toNumber()] = (_this).fm4(_1968_v95, globalState);
        _nw389[(new BigNumber(16)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(17)).toNumber()] = _2065_cf63;
        _nw389[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_1837_v1);
        _2068_v162 = _nw389;
        let _2069_v163;
        _2069_v163 = _dafny.MultiSet.fromElements(_2068_v162, _2068_v162, _2068_v162, _2068_v162, _2068_v162);
        _2069_v163 = _2069_v163;
        let _2070_v165;
        _2070_v165 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,function () {
          let _coll58 = new _dafny.Set();
          for (const _compr_58 of _dafny.IntegerRange(new BigNumber(777), new BigNumber(78))) {
            let _2071_v164 = _compr_58;
            if (((new BigNumber(777)).isLessThanOrEqualTo(_2071_v164)) && ((_2071_v164).isLessThan(new BigNumber(78)))) {
              _coll58.add((_2071_v164).plus(_1837_v1));
            }
          }
          return _coll58;
        }());
        let _2072_v166;
        _2072_v166 = _dafny.Set.fromElements(_2065_cf63);
        let _2073_v167;
        _2073_v167 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_2070_v165).contains(_2065_cf63)) ? ((_2070_v165).get(_2065_cf63)) : (_2072_v166))).length),_2065_cf63);
        let _2074_v168;
        _2074_v168 = _dafny.Map.Empty.slice().updateUnsafe(_2060_v156,(_2073_v167).contains(_2065_cf63));
        let _index376 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
        (_2056_v154)[_index376] = !(!(_1837_v1).isEqualTo(_1837_v1));
        let _2075_v169;
        _2075_v169 = _1836_v0;
        let _index377 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
        let _rhs366 = _2074_v168;
        let _rhs367 = (_2075_v169);
        let _lhs179 = _2056_v154;
        let _lhs180 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
        _2074_v168 = _rhs366;
        _lhs179[_lhs180] = _rhs367;
        if (_dafny.Seq.contains(_dafny.Seq.Concat(_2060_v156, _2060_v156), (_this).f5)) {
          _1837_v1 = (_2065_cf63).plus(new BigNumber(-47));
          r1 = _2065_cf63;
          _1837_v1 = (_2065_cf63).minus((_1968_v95)[_module.__default.safeIndex(new BigNumber((_1971_v98).length), new BigNumber((_1968_v95).length))]);
          let _2076_v170;
          _2076_v170 = _dafny.Seq.of(_2060_v156);
          _2076_v170 = _dafny.Seq.Concat(_2076_v170, _2076_v170);
          let _index378 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
          (_2056_v154)[_index378] = ((_2067_v161).Union(_2067_v161)).IsDisjointFrom(_module.__default.fm30((_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))], _1971_v98, globalState));
        } else {
          _1837_v1 = _2065_cf63;
          let _2077_v171;
          _2077_v171 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,_1836_v0);
          let _index379 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
          (_2056_v154)[_index379] = !(new BigNumber((((_1836_v0) ? (_2077_v171) : (_2077_v171))).length)).isEqualTo(_2065_cf63);
          r1 = _module.__default.fm13(_1836_v0, (_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))], (_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))], globalState);
          let _2078_v173;
          _2078_v173 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2060_v156, _2060_v156),function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_dafny.Map.Empty.slice().updateUnsafe((_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))],_2065_cf63)).update(false, _2065_cf63))).Keys.Elements) {
              let _2079_v172 = _compr_59;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_dafny.Map.Empty.slice().updateUnsafe((_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))],_2065_cf63)).update(false, _2065_cf63))).contains(_2079_v172)) {
                _coll59.push([_2079_v172,_module.D8.create_DC19(_1837_v1, (_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))])]);
              }
            }
            return _coll59;
          }());
          let _2080_v174;
          _2080_v174 = _module.D8.create_DC19(_2065_cf63, false);
          let _2081_v175;
          _2081_v175 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2080_v174);
          _2078_v173 = (_2078_v173).update(_dafny.Seq.Concat(_2060_v156, _2060_v156), _2081_v175);
          let _2082_v176;
          let _nw390 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _2082_v176 = _nw390;
          _2082_v176 = (_this).f4;
        }
        let _index380 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2068_v162).length));
        (_2068_v162)[_index380] = _2065_cf63;
        let _2083_v177;
        let _init74 = function (_2084_i15) {
          return _module.D7.create_DC16((_this).f5);
        };
        let _nw391 = Array((new BigNumber(18)).toNumber());
        for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw391.length); _i0_74++) {
          _nw391[_i0_74] = _init74(new BigNumber(_i0_74));
        }
        _2083_v177 = _nw391;
        let _2085_v178;
        _2085_v178 = _module.D7.create_DC16((_this).f5);
        let _index381 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2083_v177).length));
        (_2083_v177)[_index381] = _2085_v178;
        let _index382 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
        let _index383 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2068_v162).length));
        let _index384 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2083_v177).length));
        let _rhs368 = _1836_v0;
        let _rhs369 = _2065_cf63;
        let _rhs370 = _2065_cf63;
        let _rhs371 = (((_2056_v154)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length))]) ? (_module.D7.create_DC16((_this).f5)) : (_2085_v178));
        let _lhs181 = _2056_v154;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_2056_v154).length));
        let _lhs183 = _2068_v162;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_2068_v162).length));
        let _lhs185 = _2083_v177;
        let _lhs186 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_2083_v177).length));
        _lhs181[_lhs182] = _rhs368;
        r1 = _rhs369;
        _lhs183[_lhs184] = _rhs370;
        _lhs185[_lhs186] = _rhs371;
      } else if (_source21.is_DC42) {
        let _2086___mcc_h18 = (_source21).cf64;
        let _2087___mcc_h19 = (_source21).cf65;
        let _2088_cf65 = _2087___mcc_h19;
        let _2089_cf64 = _2086___mcc_h18;
        let _2090_v179;
        let _nw392 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2090_v179 = _nw392;
        let _index385 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2090_v179).length));
        (_2090_v179)[_index385] = _1837_v1;
        let _index386 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2090_v179).length));
        (_2090_v179)[_index386] = _1837_v1;
        let _2091_v180;
        let _nw393 = new _module.C0();
        _nw393.__ctor();
        _2091_v180 = _nw393;
        _2060_v156 = _dafny.Seq.UnicodeFromString("srfmniqdo");
        let _2092_v181;
        let _init75 = ((_2093_v95) => function (_2094_i16) {
          return _2093_v95;
        })(_1968_v95);
        let _nw394 = Array((new BigNumber(26)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw394.length); _i0_75++) {
          _nw394[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _2092_v181 = _nw394;
        let _2095_v182;
        _2095_v182 = _dafny.Map.Empty.slice().updateUnsafe(_2061_v157,_2092_v181);
        let _2096_v183;
        _2096_v183 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_1836_v0);
        let _2097_v184;
        let _nw395 = Array((new BigNumber(29)).toNumber());
        _nw395[(_dafny.ZERO).toNumber()] = _2092_v181;
        _nw395[(_dafny.ONE).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(2)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(3)).toNumber()] = (((_2095_v182).contains(_2061_v157)) ? ((_2095_v182).get(_2061_v157)) : (_2092_v181));
        _nw395[(new BigNumber(4)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(5)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(6)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(7)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(8)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(9)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(10)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(11)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(12)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(13)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(14)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(15)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(16)).toNumber()] = ((_1836_v0) ? (_2092_v181) : (_2092_v181));
        _nw395[(new BigNumber(17)).toNumber()] = ((_1836_v0) ? (_2092_v181) : (_2092_v181));
        _nw395[(new BigNumber(18)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(19)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(20)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(21)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(22)).toNumber()] = ((_1836_v0) ? (_2092_v181) : (_2092_v181));
        _nw395[(new BigNumber(23)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(24)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(25)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(26)).toNumber()] = _2092_v181;
        _nw395[(new BigNumber(27)).toNumber()] = (((((_2096_v183).contains(!(!(_module.__default.fm1(new BigNumber(866), _1837_v1, _1836_v0, _1836_v0, globalState))))) ? ((_2096_v183).get(!(!(_module.__default.fm1(new BigNumber(866), _1837_v1, _1836_v0, _1836_v0, globalState))))) : (_1836_v0))) ? (_2092_v181) : (_2092_v181));
        _nw395[(new BigNumber(28)).toNumber()] = _2092_v181;
        _2097_v184 = _nw395;
        let _index387 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_2097_v184).length));
        (_2097_v184)[_index387] = _2092_v181;
        let _2098_v185;
        _2098_v185 = _dafny.MultiSet.fromElements(_1836_v0);
        let _index388 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2090_v179).length));
        let _index389 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_2097_v184).length));
        let _rhs372 = !(_1836_v0) || (_1836_v0);
        let _rhs373 = _2088_cf65;
        let _rhs374 = (_module.__default.safeModuloInt((_2090_v179)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_2090_v179).length))], _2088_cf65)).multipliedBy(new BigNumber((_2098_v185).cardinality()));
        let _rhs375 = _2092_v181;
        let _rhs376 = _1836_v0;
        let _lhs187 = _2090_v179;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2090_v179).length));
        let _lhs189 = _2097_v184;
        let _lhs190 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_2097_v184).length));
        _1836_v0 = _rhs372;
        _lhs187[_lhs188] = _rhs373;
        r1 = _rhs374;
        _lhs189[_lhs190] = _rhs375;
        _1836_v0 = _rhs376;
      } else if (_source21.is_DC40) {
        let _2099___mcc_h20 = (_source21).cf62;
        let _2100_cf62 = _2099___mcc_h20;
        let _2101_v186;
        let _nw396 = new _module.C0();
        _nw396.__ctor();
        _2101_v186 = _nw396;
        _1836_v0 = ((_dafny.ZERO).minus(_1837_v1)).isLessThan((_dafny.ZERO).minus(_1837_v1));
        if (_1836_v0) {
          _2060_v156 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tulyymy"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-503)), function (_2102_i17) {
            return (_this).f5;
          }), _2060_v156));
          let _2103_v187;
          _2103_v187 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber((_1968_v95).length));
          _2103_v187 = (_2103_v187).update((_this).f5, (_dafny.ZERO).minus(_1837_v1));
          let _2104_v188;
          let _nw397 = new _module.C5();
          _nw397.__ctor();
          _2104_v188 = _nw397;
          let _2105_v189;
          let _nw398 = new _module.C6();
          _nw398.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1837_v1,_1837_v1));
          _2105_v189 = _nw398;
          let _2106_v190;
          let _nw399 = new _module.C7();
          _nw399.__ctor();
          _2106_v190 = _nw399;
          let _2107_v191;
          _2107_v191 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_2106_v190);
          _2107_v191 = (_2107_v191).update(_1836_v0, _2106_v190);
        } else {
          _1837_v1 = _1837_v1;
          let _2108_v192;
          let _nw400 = Array((new BigNumber(9)).toNumber()).fill([]);
          _2108_v192 = _nw400;
          let _nw401 = Array((new BigNumber(19)).toNumber()).fill([]);
          _2108_v192 = _nw401;
          let _2109_v193;
          _2109_v193 = _dafny.Map.Empty.slice().updateUnsafe(_1836_v0,_2060_v156);
          _2109_v193 = (_2109_v193).update(_1836_v0, _2060_v156);
          r1 = _module.__default.safeModuloInt(new BigNumber(194), (_1837_v1).multipliedBy(_1837_v1));
          _1837_v1 = _module.__default.safeModuloInt(_1837_v1, _1837_v1);
        }
        _1836_v0 = false;
      } else {
        let _2110___mcc_h21 = (_source21).cf66;
        let _2111_cf66 = _2110___mcc_h21;
        _1971_v98 = _dafny.Seq.Concat(_1971_v98, _1971_v98);
        (globalState).f3 = _2056_v154;
        let _2112_v194;
        let _nw402 = Array((new BigNumber(3)).toNumber());
        _2112_v194 = _nw402;
        let _2113_v195;
        _2113_v195 = _dafny.Seq.of(_2112_v194);
        let _2114_v196;
        let _nw403 = Array((new BigNumber(5)).toNumber());
        _nw403[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2113_v195, _2113_v195);
        _nw403[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2113_v195, _2113_v195);
        _nw403[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_2112_v194, _2112_v194), _2113_v195);
        _nw403[(new BigNumber(3)).toNumber()] = _2113_v195;
        _nw403[(new BigNumber(4)).toNumber()] = _2113_v195;
        _2114_v196 = _nw403;
        let _index390 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2114_v196).length));
        (_2114_v196)[_index390] = ((_1836_v0) ? (_2113_v195) : (_2113_v195));
        let _index391 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2114_v196).length));
        let _rhs377 = _dafny.Seq.Concat(_1968_v95, _1968_v95);
        let _rhs378 = _2113_v195;
        let _lhs191 = globalState;
        let _lhs192 = _2114_v196;
        let _lhs193 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2114_v196).length));
        _lhs191.f2 = _rhs377;
        _lhs192[_lhs193] = _rhs378;
        let _rhs379 = !(!(_1836_v0) || (!(!(_1836_v0))));
        let _rhs380 = _1837_v1;
        _1836_v0 = _rhs379;
        _1837_v1 = _rhs380;
      }
      if ((_this).fm5(_1837_v1, _1836_v0, _1837_v1, _1837_v1, globalState)) {
        let _2115_v197;
        let _nw404 = Array((new BigNumber(4)).toNumber());
        _nw404[(_dafny.ZERO).toNumber()] = _2056_v154;
        _nw404[(_dafny.ONE).toNumber()] = _2056_v154;
        _nw404[(new BigNumber(2)).toNumber()] = _2056_v154;
        _nw404[(new BigNumber(3)).toNumber()] = _2056_v154;
        _2115_v197 = _nw404;
        let _rhs381 = new BigNumber((_2060_v156).length);
        let _rhs382 = _2115_v197;
        let _rhs383 = (!(!(true))) || (_dafny.Seq.contains(_2060_v156, (_this).f5));
        let _rhs384 = !(_1836_v0);
        _1837_v1 = _rhs381;
        _2115_v197 = _rhs382;
        _1836_v0 = _rhs383;
        _1836_v0 = _rhs384;
        _1836_v0 = (_this).fm3(new BigNumber((_2060_v156).length), new BigNumber((_dafny.Seq.update(_1971_v98, _module.__default.safeIndex((_dafny.ZERO).minus(_1837_v1), new BigNumber((_1971_v98).length)), _1836_v0)).length), globalState);
        let _2116_v198;
        _2116_v198 = _dafny.Set.fromElements(_1836_v0, _1836_v0);
        let _2117_v199;
        _2117_v199 = _dafny.Seq.of(_2116_v198);
        _1837_v1 = _module.__default.safeModuloInt(new BigNumber(((_2117_v199)[_module.__default.safeIndex(_1837_v1, new BigNumber((_2117_v199).length))]).length), _1837_v1);
        let _2118_v200;
        let _nw405 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2118_v200 = _nw405;
        let _2119_v201;
        _2119_v201 = _dafny.Map.Empty.slice().updateUnsafe(true,_2118_v200);
        let _2120_v202;
        let _nw406 = new _module.C3();
        _nw406.__ctor((_this).f4, (_this).f5);
        _2120_v202 = _nw406;
        let _rhs385 = _1836_v0;
        let _rhs386 = (_2119_v201).Merge(_2119_v201);
        let _rhs387 = _2120_v202;
        _1836_v0 = _rhs385;
        _2119_v201 = _rhs386;
        _2120_v202 = _rhs387;
        _1837_v1 = new BigNumber((_1971_v98).length);
      } else {
        let _2121_v203;
        let _nw407 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _2121_v203 = _nw407;
        let _index392 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2121_v203).length));
        (_2121_v203)[_index392] = _1837_v1;
        let _2122_v204;
        _2122_v204 = _dafny.Set.fromElements(_1837_v1, _1837_v1, _1837_v1);
        let _2123_v205;
        let _nw408 = new _module.C8();
        _nw408.__ctor();
        _2123_v205 = _nw408;
        let _2124_v206;
        _2124_v206 = _dafny.Map.Empty.slice().updateUnsafe((_2122_v204).Difference(_2122_v204),_2123_v205);
        let _2125_v207;
        _2125_v207 = _dafny.Map.Empty.slice().updateUnsafe(_1837_v1,_1836_v0);
        let _2126_v208;
        _2126_v208 = _dafny.Seq.of((_2125_v207).Merge(_2125_v207), (_2125_v207).Merge(_2125_v207));
        let _2127_v209;
        _2127_v209 = _module.D7.create_DC17(_1836_v0, _1836_v0, !(_1836_v0), _1836_v0);
        let _2128_v210;
        _2128_v210 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13((_2127_v209).dtor_cf25, _1836_v0, !(_1836_v0), globalState),new BigNumber((_2060_v156).length));
        let _index393 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2121_v203).length));
        let _rhs388 = new BigNumber((_2126_v208).length);
        let _rhs389 = _1836_v0;
        let _rhs390 = _2124_v206;
        let _rhs391 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_1837_v1, (((_2128_v210).contains(_1837_v1)) ? ((_2128_v210).get(_1837_v1)) : (_1837_v1))), _1837_v1);
        let _rhs392 = (new BigNumber((_1968_v95).length)).minus(_module.__default.safeModuloInt(new BigNumber((_2125_v207).length), _1837_v1));
        let _lhs194 = _2121_v203;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2121_v203).length));
        _lhs194[_lhs195] = _rhs388;
        _1836_v0 = _rhs389;
        _2124_v206 = _rhs390;
        _1837_v1 = _rhs391;
        r1 = _rhs392;
        let _rhs393 = _1836_v0;
        let _rhs394 = ((_1836_v0) ? (_1836_v0) : (true));
        _1836_v0 = _rhs393;
        _1836_v0 = _rhs394;
        let _index394 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2056_v154).length));
        (_2056_v154)[_index394] = _1836_v0;
        let _2129_v211;
        _2129_v211 = _dafny.MultiSet.fromElements(_1836_v0);
        let _index395 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2056_v154).length));
        (_2056_v154)[_index395] = (_2129_v211).equals(_2129_v211);
        let _2130_v212;
        let _nw409 = new _module.C7();
        _nw409.__ctor();
        _2130_v212 = _nw409;
        _2130_v212 = _2130_v212;
        let _index396 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2056_v154).length));
        (_2056_v154)[_index396] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-532)), function (_2131_i18) {
          return (_this).f5;
        }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), function (_2132_i19) {
          return (_this).f5;
        }), _2060_v156));
      }
      r0 = _1968_v95;
      r1 = _1837_v1;
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _2133_v0;
      _2133_v0 = true;
      _2133_v0 = _2133_v0;
      let _2134_v1;
      _2134_v1 = new BigNumber(376);
      let _2135_v2;
      _2135_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,(_this).f5);
      let _2136_v3;
      _2136_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_dafny.ZERO).minus(_2134_v1));
      let _2137_v4;
      _2137_v4 = _dafny.Seq.of(((_2133_v0) ? (_2134_v1) : (_2134_v1)), (_dafny.ZERO).minus(_2134_v1), (_2134_v1).plus(new BigNumber(((_2135_v2).update(!(_2133_v0), (_this).f5)).length)), _module.__default.safeDivisionInt((((_2136_v3).contains((_this).f5)) ? ((_2136_v3).get((_this).f5)) : (new BigNumber(966))), _2134_v1));
      _2134_v1 = (_2137_v4)[_module.__default.safeIndex(_module.__default.safeModuloInt(_module.__default.fm13(_2133_v0, _2133_v0, _2133_v0, globalState), new BigNumber(857)), new BigNumber((_2137_v4).length))];
      let _2138_v5;
      _2138_v5 = _dafny.Seq.of(_2133_v0);
      let _hi17 = new BigNumber((_2138_v5).length);
      for (let _2139_i0 = ((_dafny.ZERO).minus(_2134_v1)).minus(_2134_v1); _2139_i0.isLessThan(_hi17); _2139_i0 = _2139_i0.plus(_dafny.ONE)) {
        let _2140_v6;
        _2140_v6 = _dafny.Seq.UnicodeFromString("hycyid");
        _2134_v1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2140_v6, _dafny.Seq.UnicodeFromString("snftantu")), _dafny.Seq.UnicodeFromString("tfkcaacp"))).length));
        if (_2133_v0) {
          let _2141_v7;
          let _nw410 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2141_v7 = _nw410;
          let _index397 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_2141_v7).length));
          (_2141_v7)[_index397] = _dafny.Seq.UnicodeFromString("udrapuaxq");
          let _index398 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_2141_v7).length));
          (_2141_v7)[_index398] = _2140_v6;
          _2133_v0 = _2133_v0;
          _2133_v0 = _2133_v0;
          let _2142_v8;
          _2142_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,_2134_v1);
          let _2143_v9;
          _2143_v9 = _dafny.Set.fromElements(_2133_v0, !(_2133_v0), true);
          let _2144_v10;
          _2144_v10 = _dafny.MultiSet.fromElements(_2133_v0);
          let _2145_v11;
          let _nw411 = Array((new BigNumber(13)).toNumber());
          _nw411[(_dafny.ZERO).toNumber()] = _2139_i0;
          _nw411[(_dafny.ONE).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(2)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(3)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(4)).toNumber()] = new BigNumber(((_2141_v7)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_2141_v7).length))]).length);
          _nw411[(new BigNumber(5)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(6)).toNumber()] = _2134_v1;
          _nw411[(new BigNumber(7)).toNumber()] = _2134_v1;
          _nw411[(new BigNumber(8)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(9)).toNumber()] = _2134_v1;
          _nw411[(new BigNumber(10)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(11)).toNumber()] = _2139_i0;
          _nw411[(new BigNumber(12)).toNumber()] = new BigNumber((_2140_v6).length);
          _2145_v11 = _nw411;
          let _2146_v12;
          _2146_v12 = _dafny.Set.fromElements(new BigNumber((_2137_v4).length));
          let _2147_v13;
          let _nw412 = Array((new BigNumber(16)).toNumber());
          _nw412[(_dafny.ZERO).toNumber()] = _2139_i0;
          _nw412[(_dafny.ONE).toNumber()] = new BigNumber(((_2141_v7)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_2141_v7).length))]).length);
          _nw412[(new BigNumber(2)).toNumber()] = new BigNumber(800);
          _nw412[(new BigNumber(3)).toNumber()] = _2134_v1;
          _nw412[(new BigNumber(4)).toNumber()] = _2134_v1;
          _nw412[(new BigNumber(5)).toNumber()] = new BigNumber((_2142_v8).length);
          _nw412[(new BigNumber(6)).toNumber()] = _2139_i0;
          _nw412[(new BigNumber(7)).toNumber()] = new BigNumber((_2143_v9).length);
          _nw412[(new BigNumber(8)).toNumber()] = (_2134_v1).multipliedBy(_2134_v1);
          _nw412[(new BigNumber(9)).toNumber()] = new BigNumber(848);
          _nw412[(new BigNumber(10)).toNumber()] = _2134_v1;
          _nw412[(new BigNumber(11)).toNumber()] = _2139_i0;
          _nw412[(new BigNumber(12)).toNumber()] = _2134_v1;
          _nw412[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2144_v10).cardinality()), (_module.D11.create_DC25(_2133_v0, _2134_v1, _2145_v11, _2134_v1)).dtor_cf41);
          _nw412[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2146_v12).length));
          _nw412[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_2134_v1).multipliedBy(new BigNumber((_2138_v5).length)));
          _2147_v13 = _nw412;
          let _index399 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_2145_v11).length));
          (_2145_v11)[_index399] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_2139_i0),_2133_v0)).length);
          let _index400 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_2145_v11).length));
          (_2145_v11)[_index400] = _2134_v1;
          _module.__default.m0(_2134_v1, (_2134_v1).isLessThan(_2134_v1), globalState);
        } else {
          _2134_v1 = _2139_i0;
          let _2148_v14;
          _2148_v14 = _module.D8.create_DC19(_2139_i0, true);
          let _2149_v15;
          _2149_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2139_i0,_2134_v1);
          _2134_v1 = _module.__default.safeModuloInt((_2148_v14).dtor_cf29, new BigNumber((_2149_v15).length));
          let _2150_v16;
          _2150_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2134_v1,_2133_v0);
          let _2151_v17;
          _2151_v17 = _module.D15.create_DC35(_2134_v1);
          let _rhs395 = (new BigNumber((_2137_v4).length)).plus(_2139_i0);
          let _rhs396 = (((_2150_v16).contains((_2151_v17).dtor_cf54)) ? ((_2150_v16).get((_2151_v17).dtor_cf54)) : (_2133_v0));
          _2134_v1 = _rhs395;
          _2133_v0 = _rhs396;
          _2133_v0 = _2133_v0;
          _2134_v1 = new BigNumber((_2140_v6).length);
        }
        let _2152_v18;
        _2152_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,_module.__default.fm14(globalState));
        _2152_v18 = (_2152_v18).update(_2133_v0, _2139_i0);
        _2134_v1 = (((_2139_i0).isLessThanOrEqualTo(_2139_i0)) ? (_2139_i0) : ((_dafny.ZERO).minus((_module.__default.fm14(globalState)).plus(new BigNumber(408)))));
      }
      if ((new BigNumber(-362)).isEqualTo(_2134_v1)) {
        _2134_v1 = ((_2137_v4)[_module.__default.safeIndex(_2134_v1, new BigNumber((_2137_v4).length))]).minus(_2134_v1);
        _2134_v1 = _2134_v1;
        _2134_v1 = ((_2133_v0) ? (_2134_v1) : (_module.__default.safeDivisionInt(_2134_v1, _2134_v1)));
        let _2153_v19;
        let _nw413 = new _module.C0();
        _nw413.__ctor();
        _2153_v19 = _nw413;
        _2153_v19 = _2153_v19;
        _2133_v0 = _2133_v0;
      } else {
        let _2154_v20;
        let _nw414 = new _module.C0();
        _nw414.__ctor();
        _2154_v20 = _nw414;
        let _2155_v21;
        let _nw415 = new _module.C2();
        _nw415.__ctor((_this).f4, (_this).f5);
        _2155_v21 = _nw415;
        _2155_v21 = _2155_v21;
        let _2156_v22;
        _2156_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2134_v1,_2154_v20);
        let _2157_v23;
        _2157_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_2156_v22).contains(_2134_v1)) ? ((_2156_v22).get(_2134_v1)) : (_2154_v20)),new BigNumber(630));
        _2134_v1 = _module.__default.safeModuloInt(_2134_v1, new BigNumber((_2157_v23).length));
        let _2158_v24;
        let _nw416 = new _module.C7();
        _nw416.__ctor();
        _2158_v24 = _nw416;
        if ((_2133_v0) && (_2133_v0)) {
          let _2159_v25;
          let _init76 = function (_2160_i1) {
            return (_this).f5;
          };
          let _nw417 = Array((new BigNumber(20)).toNumber());
          for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw417.length); _i0_76++) {
            _nw417[_i0_76] = _init76(new BigNumber(_i0_76));
          }
          _2159_v25 = _nw417;
          let _index401 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2159_v25).length));
          (_2159_v25)[_index401] = (_this).f5;
          let _2161_v26;
          _2161_v26 = _dafny.Seq.UnicodeFromString("ehiaxp");
          let _index402 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2159_v25).length));
          let _rhs397 = (_this).f5;
          let _rhs398 = (((new BigNumber((_2161_v26).length)).isEqualTo(_2134_v1)) ? (!(_2133_v0)) : (!((_this).fm3(_2134_v1, (_dafny.ZERO).minus(_2134_v1), globalState))));
          let _rhs399 = (_2158_v24).fm3(((true) ? (_2134_v1) : (_2134_v1)), (new BigNumber((_dafny.Seq.UnicodeFromString("gnxs")).length)).minus(_2134_v1), globalState);
          let _lhs196 = _2159_v25;
          let _lhs197 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2159_v25).length));
          _lhs196[_lhs197] = _rhs397;
          _2133_v0 = _rhs398;
          _2133_v0 = _rhs399;
          let _2162_v27;
          let _nw418 = Array((new BigNumber(12)).toNumber()).fill(false);
          _2162_v27 = _nw418;
          let _index403 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length));
          (_2162_v27)[_index403] = (_2134_v1).isLessThanOrEqualTo(_2134_v1);
          let _index404 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length));
          (_2162_v27)[_index404] = (_2158_v24).fm3(new BigNumber((_2161_v26).length), (_dafny.ZERO).minus(_2134_v1), globalState);
          let _2163_v28;
          let _nw419 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _2163_v28 = _nw419;
          let _index405 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2163_v28).length));
          (_2163_v28)[_index405] = _2134_v1;
          let _2164_v29;
          _2164_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2162_v27,_2161_v26);
          let _2165_v30;
          _2165_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_2134_v1, new BigNumber(304), globalState),_dafny.Seq.Create(_module.__default.abs(new BigNumber(762)), function (_2166_i2) {
            return (_this).f5;
          }));
          let _2167_v31;
          _2167_v31 = _dafny.Seq.of(_dafny.Seq.Concat((((_2164_v29).contains(_2162_v27)) ? ((_2164_v29).get(_2162_v27)) : ((((_2165_v30).contains((_2162_v27)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length))])) ? ((_2165_v30).get((_2162_v27)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length))])) : (_2161_v26)))), _2161_v26), _2161_v26);
          let _2168_v32;
          _2168_v32 = _dafny.MultiSet.fromElements((_2162_v27)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length))]);
          let _index406 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2163_v28).length));
          let _index407 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length));
          let _rhs400 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm16(false, _2134_v1, globalState), _dafny.Seq.UnicodeFromString("rsnh")), _2161_v26);
          let _rhs401 = new BigNumber(((_2167_v31)[_module.__default.safeIndex((new BigNumber((_2168_v32).cardinality())).minus(_2134_v1), new BigNumber((_2167_v31).length))]).length);
          let _rhs402 = _2134_v1;
          let _rhs403 = _dafny.areEqual(_2138_v5, _dafny.Seq.Concat(_2138_v5, _2138_v5));
          let _rhs404 = true;
          let _lhs198 = _2163_v28;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2163_v28).length));
          let _lhs200 = _2162_v27;
          let _lhs201 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length));
          _2161_v26 = _rhs400;
          _2134_v1 = _rhs401;
          _lhs198[_lhs199] = _rhs402;
          _2133_v0 = _rhs403;
          _lhs200[_lhs201] = _rhs404;
          _2161_v26 = _dafny.Seq.Concat(_2161_v26, (((_2162_v27)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length))]) ? (_dafny.Seq.update(_2161_v26, _module.__default.safeIndex(_2134_v1, new BigNumber((_2161_v26).length)), (_2159_v25)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_2159_v25).length))])) : (_2161_v26)));
          _2133_v0 = (_2162_v27)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_2162_v27).length))];
        } else {
          _2133_v0 = _2133_v0;
          let _2169_v33;
          _2169_v33 = _dafny.MultiSet.fromElements((_this).f5, new _dafny.CodePoint('x'.codePointAt(0)));
          let _2170_v34;
          _2170_v34 = _dafny.Set.fromElements(_2134_v1, new BigNumber(((_2169_v33).Intersect(_2169_v33)).cardinality()));
          _2170_v34 = _2170_v34;
          let _2171_v35;
          let _init77 = ((_2172_v0) => function (_2173_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2172_v0,_dafny.Seq.UnicodeFromString("mhbmkh"));
          })(_2133_v0);
          let _nw420 = Array((new BigNumber(16)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw420.length); _i0_77++) {
            _nw420[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2171_v35 = _nw420;
          let _2174_v36;
          _2174_v36 = _dafny.Seq.UnicodeFromString("j");
          let _2175_v37;
          _2175_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,_2174_v36);
          let _index408 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_2171_v35).length));
          (_2171_v35)[_index408] = _2175_v37;
          let _index409 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_2171_v35).length));
          (_2171_v35)[_index409] = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ap"), _module.__default.safeIndex(_2134_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ap")).length)), (_this).f5));
          let _2176_v38;
          _2176_v38 = new _dafny.CodePoint('l'.codePointAt(0));
          _2176_v38 = (_this).f5;
          let _rhs405 = _2174_v36;
          let _rhs406 = _2133_v0;
          _2174_v36 = _rhs405;
          _2133_v0 = _rhs406;
        }
      }
      let _2177_v39;
      _2177_v39 = _dafny.MultiSet.fromElements(_2133_v0, _2133_v0);
      let _hi18 = _2134_v1;
      for (let _2178_i4 = (_dafny.ZERO).minus((new BigNumber(((_2177_v39).update(_2133_v0, _module.__default.abs(_2134_v1))).cardinality())).minus(_2134_v1)); _2178_i4.isLessThan(_hi18); _2178_i4 = _2178_i4.plus(_dafny.ONE)) {
        _2133_v0 = (_2178_i4).isEqualTo(_2134_v1);
        let _2179_v40;
        _2179_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm43(_2133_v0, _dafny.MultiSet.fromElements(_2134_v1, (_dafny.ZERO).minus(_2134_v1), _2134_v1, _module.__default.fm13(_2133_v0, _2133_v0, _2133_v0, globalState)), globalState),_2133_v0);
        let _2180_v41;
        let _nw421 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _2180_v41 = _nw421;
        let _2181_v42;
        _2181_v42 = _module.D11.create_DC25(_2133_v0, _2178_i4, _2180_v41, new BigNumber((_dafny.Seq.of(_2134_v1)).length));
        let _2182_v43;
        _2182_v43 = _dafny.Set.fromElements(_2177_v39, (_2177_v39).update(_2133_v0, _module.__default.abs((_2181_v42).dtor_cf41)), _2177_v39);
        let _2183_v44;
        _2183_v44 = _dafny.Seq.of(_2182_v43);
        let _2184_v45;
        _2184_v45 = _dafny.Seq.UnicodeFromString("auc");
        _2133_v0 = (((_2179_v40).contains(((_2183_v44)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_2183_v44).length))]).Union(_2182_v43))) ? ((_2179_v40).get(((_2183_v44)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_2183_v44).length))]).Union(_2182_v43))) : (!(new BigNumber((_2184_v45).length)).isEqualTo(_2134_v1)));
        if (_2133_v0) {
          let _2185_v46;
          _2185_v46 = _dafny.Map.Empty.slice().updateUnsafe(_2180_v41,false);
          _2185_v46 = (_2185_v46).update(_2180_v41, (_this).fm3(new BigNumber(885), _2178_i4, globalState));
          let _2186_v47;
          let _nw422 = Array((new BigNumber(4)).toNumber());
          _nw422[(_dafny.ZERO).toNumber()] = _2133_v0;
          _nw422[(_dafny.ONE).toNumber()] = _2133_v0;
          _nw422[(new BigNumber(2)).toNumber()] = _2133_v0;
          _nw422[(new BigNumber(3)).toNumber()] = _2133_v0;
          _2186_v47 = _nw422;
          let _2187_v48;
          _2187_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2186_v47,true);
          _2187_v48 = (_2187_v48).update(_2186_v47, _2133_v0);
          let _2188_v49;
          let _init78 = ((_2189_i4, _2190_v0) => function (_2191_i5) {
            return _module.D21.create_DC50(_dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC32(new BigNumber((_dafny.MultiSet.fromElements(_2189_i4)).cardinality()), _2190_v0, _2190_v0),_dafny.Map.Empty.slice().updateUnsafe(_2190_v0,_2190_v0)));
          })(_2178_i4, _2133_v0);
          let _nw423 = Array((new BigNumber(25)).toNumber());
          for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw423.length); _i0_78++) {
            _nw423[_i0_78] = _init78(new BigNumber(_i0_78));
          }
          _2188_v49 = _nw423;
          _2188_v49 = _2188_v49;
          let _index410 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2186_v47).length));
          (_2186_v47)[_index410] = _2133_v0;
          let _index411 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_2186_v47).length));
          (_2186_v47)[_index411] = _2133_v0;
          let _index412 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2186_v47).length));
          (_2186_v47)[_index412] = !_dafny.areEqual(_2137_v4, _2137_v4);
          let _index413 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2186_v47).length));
          let _index414 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_2186_v47).length));
          let _index415 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2186_v47).length));
          let _rhs407 = _2133_v0;
          let _rhs408 = _2177_v39;
          let _rhs409 = (_2138_v5)[_module.__default.safeIndex(_2178_i4, new BigNumber((_2138_v5).length))];
          let _rhs410 = !((_this).fm3(_module.__default.fm13(_2133_v0, _2133_v0, _2133_v0, globalState), _module.__default.fm14(globalState), globalState)) || (_2133_v0);
          let _rhs411 = _2133_v0;
          let _lhs202 = _2186_v47;
          let _lhs203 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_2186_v47).length));
          let _lhs204 = _2186_v47;
          let _lhs205 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_2186_v47).length));
          let _lhs206 = _2186_v47;
          let _lhs207 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2186_v47).length));
          _2133_v0 = _rhs407;
          _2177_v39 = _rhs408;
          _lhs202[_lhs203] = _rhs409;
          _lhs204[_lhs205] = _rhs410;
          _lhs206[_lhs207] = _rhs411;
          let _2192_v50;
          _2192_v50 = _module.D18.create_DC45(_2186_v47, _2133_v0, _2134_v1, new BigNumber(128));
          let _2193_v51;
          _2193_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC11(_2135_v2),(_2192_v50).dtor_cf69);
          _2134_v1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_2184_v45, _module.__default.safeIndex(new BigNumber((_module.__default.fm31(new BigNumber(180), new BigNumber((_2193_v51).length), globalState)).length), new BigNumber((_2184_v45).length)), (_this).f5)).length));
        } else {
          _2133_v0 = _2133_v0;
          _2184_v45 = ((_2133_v0) ? (_module.__default.fm12(globalState)) : (((_2133_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-992)), function (_2194_i6) {
            return (_this).f5;
          })) : (_2184_v45))));
          let _2195_v52;
          _2195_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kx"), _2184_v45),_2184_v45);
          _2195_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_2196_i7) {
            return (_this).f5;
          }), _2184_v45),_dafny.Seq.UnicodeFromString("kmmasuhrp"));
          let _2197_v53;
          let _init79 = function (_2198_i8) {
            return false;
          };
          let _nw424 = Array((new BigNumber(22)).toNumber());
          for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw424.length); _i0_79++) {
            _nw424[_i0_79] = _init79(new BigNumber(_i0_79));
          }
          _2197_v53 = _nw424;
          let _2199_v54;
          _2199_v54 = _dafny.Set.fromElements(_2197_v53);
          let _2200_v55;
          _2200_v55 = _2199_v54;
          let _2201_v56;
          _2201_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2133_v0,_2200_v55);
          let _2202_v57;
          let _2203_v58;
          let _2204_v59;
          let _2205_v60;
          let _out21;
          let _out22;
          let _out23;
          let _out24;
          let _outcollector6 = (_this).m5(_2184_v45, _2201_v56, _2178_i4, globalState);
          _out21 = _outcollector6[0];
          _out22 = _outcollector6[1];
          _out23 = _outcollector6[2];
          _out24 = _outcollector6[3];
          _2202_v57 = _out21;
          _2203_v58 = _out22;
          _2204_v59 = _out23;
          _2205_v60 = _out24;
          let _2206_v61;
          _2206_v61 = _dafny.Map.Empty.slice().updateUnsafe(!((_2203_v58).isLessThan(_2134_v1)),((_2133_v0) ? (!(_2204_v59)) : (_2204_v59)));
          _2206_v61 = (_2206_v61).update(_2204_v59, _2204_v59);
        }
        _2134_v1 = (_2134_v1).minus(_2178_i4);
      }
      let _2207_v62;
      let _init80 = ((_2208_v1) => function (_2209_i9) {
        return (_2209_i9).multipliedBy(_2208_v1);
      })(_2134_v1);
      let _nw425 = Array((new BigNumber(3)).toNumber());
      for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw425.length); _i0_80++) {
        _nw425[_i0_80] = _init80(new BigNumber(_i0_80));
      }
      _2207_v62 = _nw425;
      let _2210_v63;
      _2210_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2207_v62,_2134_v1);
      _2210_v63 = _2210_v63;
      r0 = _2207_v62;
      return r0;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.MultiSet.Empty;
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _2211_v0;
      _2211_v0 = true;
      let _2212_i0;
      _2212_i0 = _dafny.ZERO;
      L8: {
        while (_2211_v0) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2212_i0)) {
              break L8;
            }
            _2212_i0 = (_2212_i0).plus(_dafny.ONE);
            _2211_v0 = !(_2211_v0);
            r3 = (_this).f5;
            let _2213_v1;
            _2213_v1 = new BigNumber(444);
            let _2214_v2;
            _2214_v2 = _dafny.Map.Empty.slice().updateUnsafe((_2213_v1).multipliedBy(p1),new BigNumber((_dafny.Seq.of(_2211_v0, _2211_v0)).length));
            let _2215_v3;
            let _nw426 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
            _2215_v3 = _nw426;
            let _2216_v4;
            _2216_v4 = _dafny.MultiSet.fromElements(_2215_v3, _2215_v3);
            let _2217_v5;
            _2217_v5 = _dafny.Seq.of((((_2216_v4).contains(_2215_v3)) ? ((_2216_v4).get(_2215_v3)) : (p0)), _2213_v1, p1);
            let _rhs412 = (((_2211_v0) === (_2211_v0)) ? ((_this).fm5(_2213_v1, _2211_v0, p0, p1, globalState)) : (_2211_v0));
            let _rhs413 = (((_2214_v2).contains((_2217_v5)[_module.__default.safeIndex(_2213_v1, new BigNumber((_2217_v5).length))])) ? ((_2214_v2).get((_2217_v5)[_module.__default.safeIndex(_2213_v1, new BigNumber((_2217_v5).length))])) : (_2213_v1));
            r0 = _rhs412;
            _2213_v1 = _rhs413;
            let _2218_v6;
            let _init81 = ((_2219_p1) => function (_2220_i1) {
              return (_dafny.Set.fromElements(_2219_p1)).IsProperSubsetOf(_dafny.Set.fromElements(_2219_p1));
            })(p1);
            let _nw427 = Array((new BigNumber(5)).toNumber());
            for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw427.length); _i0_81++) {
              _nw427[_i0_81] = _init81(new BigNumber(_i0_81));
            }
            _2218_v6 = _nw427;
            let _index416 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_2218_v6).length));
            (_2218_v6)[_index416] = (_2213_v1).isLessThan(p0);
            let _index417 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_2218_v6).length));
            (_2218_v6)[_index417] = !(_module.__default.fm1(_2213_v1, p1, true, _2211_v0, globalState)) || (_2211_v0);
          }
        }
      }
      let _2221_v7;
      _2221_v7 = _dafny.Seq.UnicodeFromString("bveyecq");
      let _2222_v8;
      let _nw428 = Array((new BigNumber(19)).toNumber());
      _nw428[(_dafny.ZERO).toNumber()] = _module.__default.fm16(_2211_v0, p0, globalState);
      _nw428[(_dafny.ONE).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), function (_2223_i2) {
        return (_this).f5;
      });
      _nw428[(new BigNumber(3)).toNumber()] = _module.__default.fm12(globalState);
      _nw428[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nhgvbpyh"), _2221_v7);
      _nw428[(new BigNumber(5)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(6)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(7)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(8)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2221_v7, _dafny.Seq.UnicodeFromString("qvhh"));
      _nw428[(new BigNumber(10)).toNumber()] = _module.__default.fm24(new BigNumber(-430), globalState);
      _nw428[(new BigNumber(11)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(12)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(13)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), function (_2224_i3) {
        return (_this).f5;
      }), _2221_v7);
      _nw428[(new BigNumber(15)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(16)).toNumber()] = _2221_v7;
      _nw428[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_2221_v7, _2221_v7);
      _nw428[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_2221_v7, _module.__default.safeIndex(p0, new BigNumber((_2221_v7).length)), new _dafny.CodePoint('l'.codePointAt(0)));
      _2222_v8 = _nw428;
      let _index418 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length));
      (_2222_v8)[_index418] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_2225_i4) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      });
      let _2226_v9;
      _2226_v9 = new BigNumber(334);
      let _2227_v10;
      _2227_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2221_v7);
      let _2228_v11;
      _2228_v11 = _dafny.Seq.of(_2227_v10);
      let _2229_v12;
      _2229_v12 = _dafny.Seq.of(_2226_v9);
      let _2230_v13;
      _2230_v13 = _module.D2.create_DC6(p0);
      let _index419 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length));
      let _rhs414 = _2221_v7;
      let _rhs415 = (((_dafny.ZERO).minus(_2226_v9)).minus(_2226_v9)).minus((_dafny.ZERO).minus(_2226_v9));
      let _rhs416 = (_2228_v11)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2228_v11).length))];
      let _rhs417 = ((_this).fm4(_2229_v12, globalState)).plus(_2226_v9);
      let _rhs418 = ((_2230_v13).dtor_cf7).isLessThanOrEqualTo(_module.__default.fm14(globalState));
      let _lhs208 = _2222_v8;
      let _lhs209 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length));
      _lhs208[_lhs209] = _rhs414;
      _2226_v9 = _rhs415;
      _2227_v10 = _rhs416;
      _2226_v9 = _rhs417;
      _2211_v0 = _rhs418;
      let _2231_v14;
      _2231_v14 = _module.D5.create_DC13(_2221_v7, new BigNumber((_2221_v7).length), _2211_v0);
      if ((_2231_v14).dtor_cf18) {
        let _2232_v15;
        let _nw429 = Array((new BigNumber(5)).toNumber()).fill(false);
        _2232_v15 = _nw429;
        let _2233_v16;
        _2233_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2232_v15,_2211_v0);
        let _2234_v17;
        _2234_v17 = _dafny.Seq.of(_2211_v0);
        let _2235_v18;
        _2235_v18 = _dafny.MultiSet.fromElements(_2211_v0);
        let _2236_v19;
        let _nw430 = Array((new BigNumber(25)).toNumber());
        _nw430[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_2229_v12)).cardinality());
        _nw430[(_dafny.ONE).toNumber()] = new BigNumber((((_2233_v16).update(_2232_v15, false)).Merge(_2233_v16)).length);
        _nw430[(new BigNumber(2)).toNumber()] = (_2226_v9).multipliedBy((_dafny.ZERO).minus((_this).fm4(_2229_v12, globalState)));
        _nw430[(new BigNumber(3)).toNumber()] = _2226_v9;
        _nw430[(new BigNumber(4)).toNumber()] = p1;
        _nw430[(new BigNumber(5)).toNumber()] = (p0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("e")).length));
        _nw430[(new BigNumber(6)).toNumber()] = (p0).multipliedBy(p1);
        _nw430[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2229_v12, _2229_v12)).length);
        _nw430[(new BigNumber(8)).toNumber()] = new BigNumber(917);
        _nw430[(new BigNumber(9)).toNumber()] = new BigNumber((_2234_v17).length);
        _nw430[(new BigNumber(10)).toNumber()] = _2226_v9;
        _nw430[(new BigNumber(11)).toNumber()] = ((((_2235_v18).contains(_module.__default.fm1(new BigNumber(-304), p1, _2211_v0, _2211_v0, globalState))) ? ((_2235_v18).get(_module.__default.fm1(new BigNumber(-304), p1, _2211_v0, _2211_v0, globalState))) : (p1))).plus(new BigNumber((_dafny.Seq.of(_2211_v0, _2211_v0, _2211_v0)).length));
        _nw430[(new BigNumber(12)).toNumber()] = new BigNumber(58);
        _nw430[(new BigNumber(13)).toNumber()] = p1;
        _nw430[(new BigNumber(14)).toNumber()] = (p0).plus((_dafny.ZERO).minus(p1));
        _nw430[(new BigNumber(15)).toNumber()] = _2226_v9;
        _nw430[(new BigNumber(16)).toNumber()] = p1;
        _nw430[(new BigNumber(17)).toNumber()] = (p1).minus(_2226_v9);
        _nw430[(new BigNumber(18)).toNumber()] = (_2226_v9).minus(p1);
        _nw430[(new BigNumber(19)).toNumber()] = _2226_v9;
        _nw430[(new BigNumber(20)).toNumber()] = (_this).fm4(_2229_v12, globalState);
        _nw430[(new BigNumber(21)).toNumber()] = p0;
        _nw430[(new BigNumber(22)).toNumber()] = _2226_v9;
        _nw430[(new BigNumber(23)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(970));
        _nw430[(new BigNumber(24)).toNumber()] = (p0).plus(p1);
        _2236_v19 = _nw430;
        let _index420 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        (_2236_v19)[_index420] = p1;
        let _2237_v20;
        _2237_v20 = _dafny.Set.fromElements(_2211_v0);
        let _2238_v21;
        _2238_v21 = _module.D16.create_DC38((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], !(_2211_v0), new BigNumber(((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]).length), _module.__default.fm1(p1, p0, _2211_v0, _2211_v0, globalState));
        let _2239_v22;
        _2239_v22 = _dafny.Set.fromElements(_2237_v20, _dafny.Set.fromElements((_2238_v21).dtor_cf58));
        let _2240_v23;
        _2240_v23 = _module.D18.create_DC44(_2239_v22);
        let _2241_v24;
        _2241_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2211_v0);
        let _2242_v25;
        let _nw431 = Array((new BigNumber(20)).toNumber());
        _nw431[(_dafny.ZERO).toNumber()] = (_this).f5;
        _nw431[(_dafny.ONE).toNumber()] = _module.__default.fm0(new BigNumber(966), (_2229_v12)[_module.__default.safeIndex(p1, new BigNumber((_2229_v12).length))], globalState);
        _nw431[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(4)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(5)).toNumber()] = _module.__default.fm0(_2226_v9, _2226_v9, globalState);
        _nw431[(new BigNumber(6)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(7)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(8)).toNumber()] = _module.__default.fm0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2241_v24)).length), _2226_v9, globalState);
        _nw431[(new BigNumber(9)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(10)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(11)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(12)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(13)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(14)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(15)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(16)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(17)).toNumber()] = (_this).f5;
        _nw431[(new BigNumber(18)).toNumber()] = ((!(_2211_v0)) ? ((_this).f5) : ((_this).f5));
        _nw431[(new BigNumber(19)).toNumber()] = (_this).f5;
        _2242_v25 = _nw431;
        let _index421 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2242_v25).length));
        (_2242_v25)[_index421] = (_this).f5;
        let _2243_v26;
        _2243_v26 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),_2211_v0);
        let _2244_v27;
        _2244_v27 = _dafny.Map.Empty.slice().updateUnsafe((_2243_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2211_v0)),p1);
        let _index422 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        let _index423 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2242_v25).length));
        let _rhs419 = new BigNumber((_dafny.Seq.Concat(_2229_v12, _dafny.Seq.of(_2226_v9))).length);
        let _rhs420 = _2240_v23;
        let _rhs421 = (_this).f5;
        let _rhs422 = _2244_v27;
        let _lhs210 = _2236_v19;
        let _lhs211 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        let _lhs212 = _2242_v25;
        let _lhs213 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2242_v25).length));
        _lhs210[_lhs211] = _rhs419;
        _2240_v23 = _rhs420;
        _lhs212[_lhs213] = _rhs421;
        _2244_v27 = _rhs422;
        let _index424 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        (_2236_v19)[_index424] = p1;
        let _index425 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2242_v25).length));
        (_2242_v25)[_index425] = new _dafny.CodePoint('g'.codePointAt(0));
        let _2245_v28;
        _2245_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2232_v15);
        let _2246_v29;
        _2246_v29 = _module.D17.create_DC42(_2232_v15, _2226_v9);
        let _2247_v30;
        _2247_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(921),_2232_v15);
        let _2248_v31;
        let _nw432 = Array((new BigNumber(26)).toNumber());
        _nw432[(_dafny.ZERO).toNumber()] = _2232_v15;
        _nw432[(_dafny.ONE).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(2)).toNumber()] = (((_2245_v28).contains(_2211_v0)) ? ((_2245_v28).get(_2211_v0)) : (_2232_v15));
        _nw432[(new BigNumber(3)).toNumber()] = (_2246_v29).dtor_cf64;
        _nw432[(new BigNumber(4)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(5)).toNumber()] = ((!(!(true))) ? (_2232_v15) : (_2232_v15));
        _nw432[(new BigNumber(6)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(7)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(8)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(9)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(10)).toNumber()] = (((_2247_v30).contains(p1)) ? ((_2247_v30).get(p1)) : (_2232_v15));
        _nw432[(new BigNumber(11)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(12)).toNumber()] = ((_2211_v0) ? (_2232_v15) : (_2232_v15));
        _nw432[(new BigNumber(13)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(14)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(15)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(16)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(17)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(18)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(19)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(20)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(21)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(22)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(23)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(24)).toNumber()] = _2232_v15;
        _nw432[(new BigNumber(25)).toNumber()] = _2232_v15;
        _2248_v31 = _nw432;
        let _index426 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2248_v31).length));
        (_2248_v31)[_index426] = _2232_v15;
        let _2249_v32;
        _2249_v32 = _module.D18.create_DC45(_2232_v15, _2211_v0, _2226_v9, (_2236_v19)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length))]);
        let _index427 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2248_v31).length));
        let _index428 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        let _index429 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        let _rhs423 = _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus((_2236_v19)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length))]));
        let _rhs424 = (_2249_v32).dtor_cf68;
        let _rhs425 = _module.__default.fm14(globalState);
        let _rhs426 = (_2236_v19)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length))];
        let _rhs427 = (_2236_v19)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length))];
        let _lhs214 = _2248_v31;
        let _lhs215 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2248_v31).length));
        let _lhs216 = _2236_v19;
        let _lhs217 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        let _lhs218 = _2236_v19;
        let _lhs219 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2236_v19).length));
        _2226_v9 = _rhs423;
        _lhs214[_lhs215] = _rhs424;
        _lhs216[_lhs217] = _rhs425;
        _2226_v9 = _rhs426;
        _lhs218[_lhs219] = _rhs427;
        let _2250_v33;
        let _nw433 = Array((new BigNumber(18)).toNumber()).fill([]);
        _2250_v33 = _nw433;
        let _index430 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2250_v33).length));
        (_2250_v33)[_index430] = _2236_v19;
        let _index431 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2250_v33).length));
        let _init82 = ((_2251_v0) => function (_2252_i5) {
          return (_2252_i5).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_2251_v0),_dafny.Map.Empty.slice().updateUnsafe(_2251_v0,new BigNumber(730)))).length));
        })(_2211_v0);
        let _nw434 = Array((new BigNumber(10)).toNumber());
        for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw434.length); _i0_82++) {
          _nw434[_i0_82] = _init82(new BigNumber(_i0_82));
        }
        (_2250_v33)[_index431] = _nw434;
      } else {
        let _2253_v34;
        _2253_v34 = _module.D16.create_DC38((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], _2211_v0, p1, _2211_v0);
        let _2254_v35;
        _2254_v35 = _module.D16.create_DC39(_2253_v34);
        let _2255_v36;
        _2255_v36 = _dafny.MultiSet.fromElements(_2254_v35, _2254_v35, _2254_v35);
        let _2256_v37;
        _2256_v37 = _dafny.Seq.of(_2255_v36, _2255_v36, _2255_v36, _dafny.MultiSet.fromElements(_2254_v35, _2254_v35, _2254_v35, _2254_v35));
        let _2257_v38;
        _2257_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2211_v0);
        let _2258_v39;
        _2258_v39 = _dafny.Seq.of(_2211_v0, _2211_v0);
        let _2259_v40;
        _2259_v40 = _dafny.Seq.of(_2255_v36, _dafny.MultiSet.fromElements(_2254_v35, _2254_v35, _2254_v35), (_2255_v36).Union((_2256_v37)[_module.__default.safeIndex(new BigNumber((_2257_v38).length), new BigNumber((_2256_v37).length))]), (_module.__default.fm44(globalState)), ((_2211_v0) ? ((_2255_v36).update(_2254_v35, _module.__default.abs(new BigNumber((_2258_v39).length)))) : (_2255_v36)));
        _2259_v40 = _dafny.Seq.Concat(_2259_v40, _2256_v37);
        let _2260_v41;
        let _nw435 = new _module.C7();
        _nw435.__ctor();
        _2260_v41 = _nw435;
        let _2261_v42;
        _2261_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2260_v41);
        let _2262_v43;
        _2262_v43 = _dafny.Map.Empty.slice().updateUnsafe((_2261_v42).Merge(_2261_v42),_2211_v0);
        _2262_v43 = _2262_v43;
        let _2263_v44;
        let _init83 = function (_2264_i6) {
          return true;
        };
        let _nw436 = Array((new BigNumber(15)).toNumber());
        for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw436.length); _i0_83++) {
          _nw436[_i0_83] = _init83(new BigNumber(_i0_83));
        }
        _2263_v44 = _nw436;
        let _rhs428 = _2263_v44;
        let _rhs429 = _2211_v0;
        let _lhs220 = globalState;
        _lhs220.f3 = _rhs428;
        r0 = _rhs429;
        _2226_v9 = new BigNumber((_2258_v39).length);
        let _2265_v45;
        _2265_v45 = _dafny.MultiSet.fromElements(((_2211_v0) ? (_2263_v44) : (_2263_v44)), _2263_v44, _2263_v44, _2263_v44, _2263_v44);
        _2265_v45 = _2265_v45;
      }
      if (_2211_v0) {
        _2226_v9 = ((_2211_v0) ? (_2226_v9) : (_module.__default.fm14(globalState)));
        _2226_v9 = p1;
        let _2266_v46;
        _2266_v46 = _dafny.Seq.of(true);
        let _2267_v47;
        _2267_v47 = _dafny.Seq.of(_module.__default.fm1(_2226_v9, _2226_v9, _2211_v0, (_2266_v46)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2266_v46).length))], globalState));
        r3 = (((_2211_v0) === ((_2267_v47)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2267_v47).length))])) ? ((_this).f5) : ((_this).f5));
        r0 = _2211_v0;
        _2226_v9 = p1;
      } else {
        let _2268_v48;
        let _init84 = ((_2269_p0) => function (_2270_i7) {
          return (_2270_i7).multipliedBy(_2269_p0);
        })(p0);
        let _nw437 = Array((new BigNumber(10)).toNumber());
        for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw437.length); _i0_84++) {
          _nw437[_i0_84] = _init84(new BigNumber(_i0_84));
        }
        _2268_v48 = _nw437;
        let _index432 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
        (_2268_v48)[_index432] = p0;
        let _index433 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
        (_2268_v48)[_index433] = new BigNumber(878);
        _2211_v0 = (_this).fm5(p0, _2211_v0, p0, (_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))], globalState);
        if (false) {
          let _2271_v49;
          _2271_v49 = _dafny.MultiSet.fromElements(_2211_v0, _2211_v0, _2211_v0, _2211_v0);
          _2271_v49 = _2271_v49;
          let _2272_v50;
          _2272_v50 = _module.D7.create_DC16((_this).f5);
          let _2273_v51;
          _2273_v51 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), (_2272_v50).dtor_cf23);
          let _2274_v52;
          _2274_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),new BigNumber((_2273_v51).cardinality()));
          let _2275_v53;
          _2275_v53 = _dafny.Seq.of(_2271_v49, _dafny.MultiSet.fromElements(_2211_v0, _2211_v0));
          let _2276_v54;
          _2276_v54 = _dafny.Seq.of(_2211_v0, _2211_v0, _2211_v0);
          let _2277_v55;
          _2277_v55 = _module.D11.create_DC25(_2211_v0, new BigNumber(-927), _2268_v48, p1);
          let _2278_v56;
          let _nw438 = Array((new BigNumber(28)).toNumber());
          _nw438[(_dafny.ZERO).toNumber()] = _2211_v0;
          _nw438[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(p0)).isLessThan(_2226_v9);
          _nw438[(new BigNumber(2)).toNumber()] = ((!(_2211_v0)) ? (_2211_v0) : ((_this).fm5((_dafny.ZERO).minus(_2226_v9), _2211_v0, p0, p1, globalState)));
          _nw438[(new BigNumber(3)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(4)).toNumber()] = (_this).fm3((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))], (((_2274_v52).contains(new BigNumber((_2221_v7).length))) ? ((_2274_v52).get(new BigNumber((_2221_v7).length))) : (new BigNumber((_dafny.Seq.of(p1)).length))), globalState);
          _nw438[(new BigNumber(5)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(6)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(7)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(8)).toNumber()] = false;
          _nw438[(new BigNumber(9)).toNumber()] = (_2226_v9).isEqualTo(p0);
          _nw438[(new BigNumber(10)).toNumber()] = (_2273_v51).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f5));
          _nw438[(new BigNumber(11)).toNumber()] = false;
          _nw438[(new BigNumber(12)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(13)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(14)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(15)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(16)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(17)).toNumber()] = (((_this).fm5(_2226_v9, _2211_v0, new BigNumber((_2271_v49).cardinality()), new BigNumber((_2221_v7).length), globalState)) ? (false) : (_2211_v0));
          _nw438[(new BigNumber(18)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(19)).toNumber()] = ((!(_2211_v0)) ? (_2211_v0) : (_2211_v0));
          _nw438[(new BigNumber(20)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(21)).toNumber()] = !(_2211_v0);
          _nw438[(new BigNumber(22)).toNumber()] = (_dafny.MultiSet.FromArray(_2276_v54)).IsProperSubsetOf((_2275_v53)[_module.__default.safeIndex(p0, new BigNumber((_2275_v53).length))]);
          _nw438[(new BigNumber(23)).toNumber()] = ((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))]).isLessThan(p1);
          _nw438[(new BigNumber(24)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(25)).toNumber()] = _2211_v0;
          _nw438[(new BigNumber(26)).toNumber()] = (_2277_v55).dtor_cf38;
          _nw438[(new BigNumber(27)).toNumber()] = _dafny.Seq.IsPrefixOf(_2276_v54, _2276_v54);
          _2278_v56 = _nw438;
          let _index434 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2278_v56).length));
          (_2278_v56)[_index434] = (_2271_v49).IsDisjointFrom(_2271_v49);
          let _2279_v57;
          _2279_v57 = _dafny.Set.fromElements(_2276_v54);
          let _2280_v58;
          _2280_v58 = _dafny.Map.Empty.slice().updateUnsafe((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))],_2279_v57);
          let _2281_v59;
          _2281_v59 = _module.D23.create_DC54((((_2280_v58).contains((_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_2284_v48) => function (_2285_i8) {
  return (_2284_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2284_v48).length))];
})(_2268_v48)), globalState))) ? ((_2280_v58).get((_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_2282_v48) => function (_2283_i8) {
  return (_2282_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2282_v48).length))];
})(_2268_v48)), globalState))) : (_2279_v57)));
          let _index435 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2278_v56).length));
          let _rhs430 = _2268_v48;
          let _rhs431 = (_2279_v57).IsProperSubsetOf((_2281_v59).dtor_cf84);
          let _lhs221 = _2278_v56;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2278_v56).length));
          _2268_v48 = _rhs430;
          _lhs221[_lhs222] = _rhs431;
          let _index436 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
          (_2268_v48)[_index436] = (((_2274_v52).contains((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))])) ? ((_2274_v52).get((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))])) : (_module.__default.safeModuloInt(new BigNumber((_2221_v7).length), p0)));
          let _2286_v60;
          let _nw439 = new _module.C5();
          _nw439.__ctor();
          _2286_v60 = _nw439;
          let _2287_v61;
          _2287_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))])),(_2278_v56)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2278_v56).length))]);
          let _2288_v62;
          _2288_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2287_v61,(_this).f4);
          let _2289_v63;
          let _nw440 = new _module.C3();
          _nw440.__ctor((((_2288_v62).contains(_2287_v61)) ? ((_2288_v62).get(_2287_v61)) : ((_this).f4)), _module.__default.fm0(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("pufjmt"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("pufjmt")).length)), (_this).f5)).length), (_2268_v48)[_module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length))], globalState));
          _2289_v63 = _nw440;
        } else {
          let _2290_v64;
          _2290_v64 = _dafny.Seq.of(_2268_v48, _2268_v48, _2268_v48);
          let _2291_v65;
          _2291_v65 = _dafny.Seq.of(false, _2211_v0);
          _2290_v64 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_2290_v64, _2290_v64), _module.__default.safeIndex(new BigNumber((_2291_v65).length), new BigNumber((_dafny.Seq.Concat(_2290_v64, _2290_v64)).length)), _2268_v48), _2290_v64);
          let _2292_v66;
          _2292_v66 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_2229_v12),(_2291_v65)[_module.__default.safeIndex(new BigNumber((_2221_v7).length), new BigNumber((_2291_v65).length))]);
          let _2293_v67;
          _2293_v67 = _2292_v66;
          let _index437 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
          let _rhs432 = new BigNumber((((_2292_v66).Merge((_2293_v67))).Merge(_2292_v66)).length);
          let _rhs433 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), function (_2294_i9) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          let _lhs223 = _2268_v48;
          let _lhs224 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
          _lhs223[_lhs224] = _rhs432;
          _2221_v7 = _rhs433;
          let _index438 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_2268_v48).length));
          (_2268_v48)[_index438] = p0;
          let _2295_v68;
          let _nw441 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
          _2295_v68 = _nw441;
          let _index439 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2295_v68).length));
          (_2295_v68)[_index439] = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,false);
          let _2296_v69;
          _2296_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2211_v0);
          let _index440 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2295_v68).length));
          (_2295_v68)[_index440] = ((_2211_v0) ? (_2296_v69) : ((_dafny.Map.Empty.slice().updateUnsafe(_2211_v0,false)).Merge(_2296_v69)));
          let _2297_v70;
          let _nw442 = Array((new BigNumber(18)).toNumber()).fill(false);
          _2297_v70 = _nw442;
          let _2298_v71;
          _2298_v71 = _dafny.MultiSet.fromElements(_2297_v70, _2297_v70);
          let _2299_v72;
          _2299_v72 = _module.D25.create_DC59(_2298_v71);
          _2298_v71 = ((_2299_v72).dtor_cf95).Difference(_2298_v71);
        }
        let _2300_v73;
        let _init85 = ((_2301_v0) => function (_2302_i10) {
          return _2301_v0;
        })(_2211_v0);
        let _nw443 = Array((new BigNumber(27)).toNumber());
        for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw443.length); _i0_85++) {
          _nw443[_i0_85] = _init85(new BigNumber(_i0_85));
        }
        _2300_v73 = _nw443;
        let _index441 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_2300_v73).length));
        (_2300_v73)[_index441] = _2211_v0;
        let _2303_v74;
        _2303_v74 = _dafny.MultiSet.fromElements((_this).f5, (_this).f5);
        let _index442 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_2300_v73).length));
        (_2300_v73)[_index442] = (_2303_v74).equals(((_2211_v0) ? (_2303_v74) : (_2303_v74)));
        r0 = (_2300_v73)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_2300_v73).length))];
      }
      if ((_module.__default.fm45(new BigNumber((_dafny.Seq.UnicodeFromString("dtaqiowis")).length), globalState)).dtor_cf97) {
        r3 = (_this).f5;
        let _2304_v75;
        _2304_v75 = _dafny.Seq.of(_2211_v0, _2211_v0, _2211_v0, false, _2211_v0);
        let _2305_v76;
        _2305_v76 = _dafny.Seq.of(_module.__default.fm1(_2226_v9, p1, (_2304_v75)[_module.__default.safeIndex(p0, new BigNumber((_2304_v75).length))], _2211_v0, globalState), _2211_v0, true);
        let _2306_v77;
        _2306_v77 = _module.D23.create_DC55((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], _2211_v0, _2211_v0, (_2211_v0) || (_2211_v0), _2304_v75);
        let _source22 = _2306_v77;
        if (_source22.is_DC55) {
          let _2307___mcc_h0 = (_source22).cf85;
          let _2308___mcc_h1 = (_source22).cf86;
          let _2309___mcc_h2 = (_source22).cf87;
          let _2310___mcc_h3 = (_source22).cf88;
          let _2311___mcc_h4 = (_source22).cf89;
          let _2312_cf89 = _2311___mcc_h4;
          let _2313_cf88 = _2310___mcc_h3;
          let _2314_cf87 = _2309___mcc_h2;
          let _2315_cf86 = _2308___mcc_h1;
          let _2316_cf85 = _2307___mcc_h0;
          _2313_cf88 = (_2314_cf87) === ((_2211_v0) && (_2211_v0));
          r2 = _module.__default.fm11(globalState);
          let _2317_v78;
          _2317_v78 = _module.D16.create_DC38(_2316_cf85, _2313_cf88, p0, _2211_v0);
          let _2318_v79;
          let _nw444 = Array((new BigNumber(20)).toNumber());
          _nw444[(_dafny.ZERO).toNumber()] = _2314_cf87;
          _nw444[(_dafny.ONE).toNumber()] = _2211_v0;
          _nw444[(new BigNumber(2)).toNumber()] = true;
          _nw444[(new BigNumber(3)).toNumber()] = _2211_v0;
          _nw444[(new BigNumber(4)).toNumber()] = _2315_cf86;
          _nw444[(new BigNumber(5)).toNumber()] = _2211_v0;
          _nw444[(new BigNumber(6)).toNumber()] = _2314_cf87;
          _nw444[(new BigNumber(7)).toNumber()] = _2315_cf86;
          _nw444[(new BigNumber(8)).toNumber()] = _2315_cf86;
          _nw444[(new BigNumber(9)).toNumber()] = (_2317_v78).dtor_cf58;
          _nw444[(new BigNumber(10)).toNumber()] = _2314_cf87;
          _nw444[(new BigNumber(11)).toNumber()] = false;
          _nw444[(new BigNumber(12)).toNumber()] = _2313_cf88;
          _nw444[(new BigNumber(13)).toNumber()] = !(_2314_cf87);
          _nw444[(new BigNumber(14)).toNumber()] = _2315_cf86;
          _nw444[(new BigNumber(15)).toNumber()] = _2314_cf87;
          _nw444[(new BigNumber(16)).toNumber()] = _2315_cf86;
          _nw444[(new BigNumber(17)).toNumber()] = _2211_v0;
          _nw444[(new BigNumber(18)).toNumber()] = _2314_cf87;
          _nw444[(new BigNumber(19)).toNumber()] = _2211_v0;
          _2318_v79 = _nw444;
          let _2319_v80;
          _2319_v80 = _module.D17.create_DC42(_2318_v79, p1);
          _2319_v80 = _2319_v80;
          let _2320_v81;
          let _nw445 = Array((new BigNumber(24)).toNumber()).fill(_module.D11.Default());
          _2320_v81 = _nw445;
          let _2321_v82;
          let _nw446 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _2321_v82 = _nw446;
          let _2322_v83;
          _2322_v83 = _module.D18.create_DC45(_2318_v79, _2211_v0, new BigNumber((_dafny.Seq.UnicodeFromString("mjvmy")).length), p1);
          let _2323_v84;
          _2323_v84 = _module.D11.create_DC25(_2315_cf86, p1, _2321_v82, (_2322_v83).dtor_cf70);
          let _index443 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_2320_v81).length));
          (_2320_v81)[_index443] = _2323_v84;
          let _index444 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2318_v79).length));
          (_2318_v79)[_index444] = (false) && (!(_2211_v0));
          let _index445 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_2320_v81).length));
          let _index446 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2318_v79).length));
          let _rhs434 = _2323_v84;
          let _rhs435 = _2313_cf88;
          let _lhs225 = _2320_v81;
          let _lhs226 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_2320_v81).length));
          let _lhs227 = _2318_v79;
          let _lhs228 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_2318_v79).length));
          _lhs225[_lhs226] = _rhs434;
          _lhs227[_lhs228] = _rhs435;
        } else if (_source22.is_DC56) {
          let _2324___mcc_h5 = (_source22).cf90;
          let _2325___mcc_h6 = (_source22).cf91;
          let _2326___mcc_h7 = (_source22).cf92;
          let _2327___mcc_h8 = (_source22).cf93;
          let _2328_cf93 = _2327___mcc_h8;
          let _2329_cf92 = _2326___mcc_h7;
          let _2330_cf91 = _2325___mcc_h6;
          let _2331_cf90 = _2324___mcc_h5;
          let _2332_v85;
          let _init86 = function (_2333_i11) {
            return (_dafny.MultiSet.fromElements(_module.D7.create_DC16(new _dafny.CodePoint('x'.codePointAt(0))), _module.D7.create_DC16((_this).f5))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D7.create_DC16((_this).f5))));
          };
          let _nw447 = Array((new BigNumber(18)).toNumber());
          for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw447.length); _i0_86++) {
            _nw447[_i0_86] = _init86(new BigNumber(_i0_86));
          }
          _2332_v85 = _nw447;
          let _index447 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2332_v85).length));
          (_2332_v85)[_index447] = _2329_cf92;
          let _index448 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2332_v85).length));
          (_2332_v85)[_index448] = !(_2331_cf90) || (!(_2211_v0));
          let _2334_v86;
          let _nw448 = new _module.C0();
          _nw448.__ctor();
          _2334_v86 = _nw448;
          let _2335_v87;
          _2335_v87 = _dafny.Map.Empty.slice().updateUnsafe(_2328_cf93,_2334_v86);
          let _2336_v88;
          _2336_v88 = _dafny.Seq.of(_2334_v86, _2334_v86, (((_2335_v87).contains(_2211_v0)) ? ((_2335_v87).get(_2211_v0)) : (_2334_v86)), (((_2335_v87).contains(_2211_v0)) ? ((_2335_v87).get(_2211_v0)) : (_2334_v86)));
          let _2337_v89;
          let _nw449 = Array((new BigNumber(24)).toNumber());
          _nw449[(_dafny.ZERO).toNumber()] = _2334_v86;
          _nw449[(_dafny.ONE).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(2)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(3)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(4)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(5)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(6)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(7)).toNumber()] = (_2336_v88)[_module.__default.safeIndex(p1, new BigNumber((_2336_v88).length))];
          _nw449[(new BigNumber(8)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(9)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(10)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(11)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(12)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(13)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(14)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(15)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(16)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(17)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(18)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(19)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(20)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(21)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(22)).toNumber()] = _2334_v86;
          _nw449[(new BigNumber(23)).toNumber()] = _2334_v86;
          _2337_v89 = _nw449;
          let _2338_v90;
          _2338_v90 = _2337_v89;
          let _2339_v91;
          _2339_v91 = _dafny.MultiSet.fromElements(_2337_v89, _2338_v90, _2338_v90, _2338_v90);
          _2339_v91 = _2339_v91;
          _2226_v9 = _2226_v9;
          let _2340_v92;
          _2340_v92 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2226_v9);
          let _2341_v93;
          _2341_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2329_cf92,_2328_cf93);
          let _2342_v94;
          _2342_v94 = _module.D17.create_DC40(_dafny.Map.Empty.slice().updateUnsafe(_2340_v92,_2341_v93));
          let _2343_v95;
          _2343_v95 = _module.D17.create_DC43(_2342_v94);
          let _2344_v96;
          let _nw450 = Array((new BigNumber(3)).toNumber());
          _nw450[(_dafny.ZERO).toNumber()] = (_this).f5;
          _nw450[(_dafny.ONE).toNumber()] = (_this).f5;
          _nw450[(new BigNumber(2)).toNumber()] = (_this).f5;
          _2344_v96 = _nw450;
          let _2345_v97;
          _2345_v97 = _module.D7.create_DC16((_this).f5);
          let _index449 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2344_v96).length));
          (_2344_v96)[_index449] = (_2345_v97).dtor_cf23;
          let _index450 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2332_v85).length));
          let _index451 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2344_v96).length));
          let _rhs436 = _2343_v95;
          let _rhs437 = ((!(true) || (_2331_cf90)) ? (_2330_cf91) : (!_dafny.areEqual((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], _dafny.Seq.UnicodeFromString("pwsytbx"))));
          let _rhs438 = _2211_v0;
          let _rhs439 = (_this).f5;
          let _rhs440 = (_dafny.ZERO).minus(_2226_v9);
          let _lhs229 = _2332_v85;
          let _lhs230 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2332_v85).length));
          let _lhs231 = _2344_v96;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2344_v96).length));
          _2343_v95 = _rhs436;
          _2330_cf91 = _rhs437;
          _lhs229[_lhs230] = _rhs438;
          _lhs231[_lhs232] = _rhs439;
          _2226_v9 = _rhs440;
        } else if (_source22.is_DC57) {
          let _2346_v98;
          _2346_v98 = _dafny.Set.fromElements(p1, (_2229_v12)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2229_v12).length))], p1);
          let _2347_v99;
          _2347_v99 = _dafny.MultiSet.fromElements(_2211_v0, _2211_v0);
          let _2348_v100;
          _2348_v100 = _dafny.Map.Empty.slice().updateUnsafe(_2346_v98,_2347_v99);
          _2226_v9 = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_2346_v98,_2347_v99)).Merge(_2348_v100)).Merge(_2348_v100)).length);
          let _2349_v101;
          let _nw451 = Array((new BigNumber(22)).toNumber());
          _nw451[(_dafny.ZERO).toNumber()] = _2211_v0;
          _nw451[(_dafny.ONE).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(2)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(3)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(4)).toNumber()] = true;
          _nw451[(new BigNumber(5)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(6)).toNumber()] = true;
          _nw451[(new BigNumber(7)).toNumber()] = !(_2211_v0);
          _nw451[(new BigNumber(8)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(9)).toNumber()] = !(_2211_v0);
          _nw451[(new BigNumber(10)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(11)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(12)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(13)).toNumber()] = !(!(true));
          _nw451[(new BigNumber(14)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(15)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(16)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(17)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(18)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(19)).toNumber()] = false;
          _nw451[(new BigNumber(20)).toNumber()] = _2211_v0;
          _nw451[(new BigNumber(21)).toNumber()] = _2211_v0;
          _2349_v101 = _nw451;
          let _2350_v102;
          _2350_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2349_v101);
          let _2351_v103;
          _2351_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,p1);
          let _2352_v104;
          _2352_v104 = _dafny.Seq.of(_2351_v103);
          _2350_v102 = (_2350_v102).update(new BigNumber((((_2352_v104)[_module.__default.safeIndex(new BigNumber(83), new BigNumber((_2352_v104).length))]).Merge(_2351_v103)).length), _2349_v101);
          let _2353_v105;
          _2353_v105 = _dafny.Set.fromElements(_2211_v0);
          let _2354_v106;
          _2354_v106 = _dafny.Seq.of(_dafny.Set.fromElements(_2211_v0), (_dafny.Set.fromElements(_2211_v0)).Difference(_2353_v105));
          let _rhs441 = !(!(_2211_v0)) || (_2211_v0);
          let _rhs442 = new BigNumber((_2354_v106).length);
          let _rhs443 = _module.__default.safeModuloInt(p0, p1);
          _2211_v0 = _rhs441;
          _2226_v9 = _rhs442;
          _2226_v9 = _rhs443;
          let _2355_v107;
          let _nw452 = Array((new BigNumber(17)).toNumber()).fill([]);
          _2355_v107 = _nw452;
          let _2356_v108;
          let _nw453 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _2356_v108 = _nw453;
          let _index452 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2355_v107).length));
          (_2355_v107)[_index452] = _2356_v108;
          let _index453 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2355_v107).length));
          let _nw454 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          (_2355_v107)[_index453] = _nw454;
        } else {
          let _2357___mcc_h9 = (_source22).cf84;
          let _2358_cf84 = _2357___mcc_h9;
          let _2359_v109;
          let _init87 = ((_2360_v0) => function (_2361_i12) {
            return _2360_v0;
          })(_2211_v0);
          let _nw455 = Array((new BigNumber(6)).toNumber());
          for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw455.length); _i0_87++) {
            _nw455[_i0_87] = _init87(new BigNumber(_i0_87));
          }
          _2359_v109 = _nw455;
          let _2362_v110;
          _2362_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_module.D17.create_DC42(_2359_v109, p0));
          let _2363_v111;
          _2363_v111 = _dafny.Map.Empty.slice().updateUnsafe(true,!(_2211_v0));
          let _2364_v112;
          _2364_v112 = _module.D17.create_DC42(_2359_v109, p1);
          let _pat_let_tv54 = _2229_v12;
          _2362_v110 = (_2362_v110).update((_this).fm5(_2226_v9, _2211_v0, p1, new BigNumber((_2363_v111).length), globalState), function (_pat_let53_0) {
            return function (_2365_dt__update__tmp_h1) {
              return function (_pat_let54_0) {
                return function (_2366_dt__update_hcf65_h0) {
                  return _module.D17.create_DC42((_2365_dt__update__tmp_h1).dtor_cf64, _2366_dt__update_hcf65_h0);
                }(_pat_let54_0);
              }(new BigNumber((_pat_let_tv54).length));
            }(_pat_let53_0);
          }(_2364_v112));
          let _2367_v113;
          _2367_v113 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(774)), function (_2368_i13) {
            return (_this).f5;
          }), _2221_v7);
          r0 = (_module.__default.fm46(globalState)).IsSubsetOf((_2367_v113).update((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], _module.__default.abs(new BigNumber(903))));
          let _2369_v114;
          _2369_v114 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2211_v0), _2363_v111);
          let _2370_v115;
          _2370_v115 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,new BigNumber(99));
          let _2371_v116;
          _2371_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,((_2369_v114)[_module.__default.safeIndex(new BigNumber((_2370_v115).length), new BigNumber((_2369_v114).length))]).Merge(_2363_v111));
          let _rhs444 = new BigNumber(595);
          let _rhs445 = _2371_v116;
          _2226_v9 = _rhs444;
          _2371_v116 = _rhs445;
          _2226_v9 = _2226_v9;
        }
        let _2372_v118;
        let _nw456 = new _module.C6();
        _nw456.__ctor((function () {
          let _coll60 = new _dafny.Map();
          for (const _compr_60 of _dafny.IntegerRange(new BigNumber(-726), new BigNumber(473))) {
            let _2373_v117 = _compr_60;
            if (((new BigNumber(-726)).isLessThanOrEqualTo(_2373_v117)) && ((_2373_v117).isLessThan(new BigNumber(473)))) {
              _coll60.push([(_2373_v117).plus(_2226_v9),(_dafny.ZERO).minus(p0)]);
            }
          }
          return _coll60;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_2229_v12).length))));
        _2372_v118 = _nw456;
        r0 = _2211_v0;
        let _2374_v119;
        let _nw457 = Array((new BigNumber(15)).toNumber()).fill(false);
        _2374_v119 = _nw457;
        let _2375_v120;
        _2375_v120 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-552),_2374_v119);
        let _2376_v121;
        _2376_v121 = _dafny.MultiSet.fromElements((((_2375_v120).contains(p0)) ? ((_2375_v120).get(p0)) : (_2374_v119)));
        let _2377_v122;
        _2377_v122 = _module.D25.create_DC59(_2376_v121);
        let _source23 = _2377_v122;
        if (_source23.is_DC60) {
          let _2378___mcc_h10 = (_source23).cf96;
          let _2379___mcc_h11 = (_source23).cf97;
          let _2380___mcc_h12 = (_source23).cf98;
          let _2381_cf98 = _2380___mcc_h12;
          let _2382_cf97 = _2379___mcc_h11;
          let _2383_cf96 = _2378___mcc_h10;
          let _2384_v123;
          _2384_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2383_cf96,_2226_v9);
          let _2385_v124;
          _2385_v124 = _dafny.Seq.of((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]);
          let _rhs446 = _2221_v7;
          let _rhs447 = (_2385_v124)[_module.__default.safeIndex(p0, new BigNumber((_2385_v124).length))];
          let _rhs448 = (_2384_v123).update(_2211_v0, _2226_v9);
          let _rhs449 = _2211_v0;
          r1 = _rhs446;
          _2221_v7 = _rhs447;
          _2384_v123 = _rhs448;
          _2382_cf97 = _rhs449;
          _2227_v10 = function () {
            let _coll61 = new _dafny.Map();
            for (const _compr_61 of ((_2372_v118).f10).Keys.Elements) {
              let _2386_v125 = _compr_61;
              if (((_2372_v118).f10).contains(_2386_v125)) {
                _coll61.push([(_2386_v125).minus(p1),_dafny.Seq.update(((_2383_cf96) ? (_2381_cf98) : (_2221_v7)), _module.__default.safeIndex(p0, new BigNumber((((_2383_cf96) ? (_2381_cf98) : (_2221_v7))).length)), (_this).f5)]);
              }
            }
            return _coll61;
          }();
          let _2387_v126;
          let _init88 = ((_2388_v9) => function (_2389_i14) {
            return (_2389_i14).plus(_2388_v9);
          })(_2226_v9);
          let _nw458 = Array((new BigNumber(6)).toNumber());
          for (let _i0_88 = 0; _i0_88 < new BigNumber(_nw458.length); _i0_88++) {
            _nw458[_i0_88] = _init88(new BigNumber(_i0_88));
          }
          _2387_v126 = _nw458;
          let _index454 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2387_v126).length));
          (_2387_v126)[_index454] = p1;
          let _index455 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2387_v126).length));
          (_2387_v126)[_index455] = ((_this).fm4(_2229_v12, globalState)).multipliedBy(p1);
          let _index456 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2387_v126).length));
          let _rhs450 = _module.__default.fm14(globalState);
          let _rhs451 = (new BigNumber((_2381_cf98).length)).multipliedBy((p0).multipliedBy(p1));
          let _lhs233 = _2387_v126;
          let _lhs234 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2387_v126).length));
          _lhs233[_lhs234] = _rhs450;
          _2226_v9 = _rhs451;
        } else {
          let _2390___mcc_h13 = (_source23).cf95;
          let _2391_cf95 = _2390___mcc_h13;
          let _2392_v127;
          let _nw459 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _2392_v127 = _nw459;
          let _2393_v128;
          _2393_v128 = _dafny.Map.Empty.slice().updateUnsafe(!(_2211_v0),(_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]);
          let _2394_v129;
          _2394_v129 = _dafny.Seq.of(_2221_v7, (((_2393_v128).contains(_2211_v0)) ? ((_2393_v128).get(_2211_v0)) : (_dafny.Seq.UnicodeFromString("qna"))));
          let _index457 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2392_v127).length));
          (_2392_v127)[_index457] = _2394_v129;
          let _index458 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2392_v127).length));
          (_2392_v127)[_index458] = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("vkkyq"), _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(true, _2211_v0, _2211_v0)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("vkkyq")).length)), (_this).f5), (_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], _2221_v7);
          let _2395_v130;
          _2395_v130 = _dafny.MultiSet.fromElements(false, false, _2211_v0);
          let _2396_v131;
          _2396_v131 = _dafny.Map.Empty.slice().updateUnsafe(_2395_v130,(_2229_v12)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2229_v12).length))]);
          let _2397_v132;
          let _nw460 = Array((new BigNumber(11)).toNumber());
          _nw460[(_dafny.ZERO).toNumber()] = p0;
          _nw460[(_dafny.ONE).toNumber()] = p0;
          _nw460[(new BigNumber(2)).toNumber()] = new BigNumber(-17);
          _nw460[(new BigNumber(3)).toNumber()] = p1;
          _nw460[(new BigNumber(4)).toNumber()] = _2226_v9;
          _nw460[(new BigNumber(5)).toNumber()] = p0;
          _nw460[(new BigNumber(6)).toNumber()] = p1;
          _nw460[(new BigNumber(7)).toNumber()] = new BigNumber((_2396_v131).length);
          _nw460[(new BigNumber(8)).toNumber()] = _2226_v9;
          _nw460[(new BigNumber(9)).toNumber()] = p1;
          _nw460[(new BigNumber(10)).toNumber()] = p0;
          _2397_v132 = _nw460;
          let _2398_v133;
          _2398_v133 = _dafny.MultiSet.fromElements(_2226_v9, p0, new BigNumber((_dafny.Seq.of(_2397_v132, _2397_v132)).length), p0);
          r2 = (_2398_v133).Union(_2398_v133);
          _2226_v9 = new BigNumber(((_2372_v118).f10).length);
          let _2399_v134;
          _2399_v134 = _module.D15.create_DC35(_2226_v9);
          let _2400_v135;
          _2400_v135 = _module.D15.create_DC36(_2399_v134);
          let _2401_v136;
          _2401_v136 = _dafny.Set.fromElements(_2374_v119);
          let _2402_v137;
          _2402_v137 = _2401_v136;
          let _2403_v138;
          let _nw461 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _2403_v138 = _nw461;
          let _2404_v139;
          _2404_v139 = _module.D4.create_DC10(_2402_v137, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(834)), function (_2405_i15) {
  return new _dafny.CodePoint('d'.codePointAt(0));
}), _2403_v138);
          let _2406_v140;
          let _nw462 = new _module.C4();
          _nw462.__ctor(_2226_v9, _2226_v9, (_this).f4, new _dafny.CodePoint('w'.codePointAt(0)));
          _2406_v140 = _nw462;
          let _2407_v141;
          _2407_v141 = _dafny.Map.Empty.slice().updateUnsafe((_2404_v139).dtor_cf12,_2406_v140);
          let _2408_v142;
          _2408_v142 = _dafny.Map.Empty.slice().updateUnsafe(_2400_v135,_2407_v141);
          _2408_v142 = (_2408_v142).update(_2400_v135, _2407_v141);
        }
      } else {
        let _2409_v143;
        _2409_v143 = _dafny.Set.fromElements(p1, p0);
        let _2410_v145;
        _2410_v145 = _module.D14.create_DC32(new BigNumber((_2221_v7).length), _2211_v0, _2211_v0);
        let _2411_v146;
        let _nw463 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _2411_v146 = _nw463;
        let _2412_v147;
        _2412_v147 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2411_v146);
        let _2413_v150;
        let _nw464 = Array((new BigNumber(17)).toNumber());
        _nw464[(_dafny.ZERO).toNumber()] = (_dafny.Set.fromElements(_2226_v9)).Union(_2409_v143);
        _nw464[(_dafny.ONE).toNumber()] = ((_2211_v0) ? (_dafny.Set.fromElements(p0)) : (_2409_v143));
        _nw464[(new BigNumber(2)).toNumber()] = function () {
          let _coll62 = new _dafny.Set();
          for (const _compr_62 of _dafny.IntegerRange(new BigNumber(-228), new BigNumber(726))) {
            let _2414_v144 = _compr_62;
            if (((new BigNumber(-228)).isLessThanOrEqualTo(_2414_v144)) && ((_2414_v144).isLessThan(new BigNumber(726)))) {
              _coll62.add((_2414_v144).minus(p0));
            }
          }
          return _coll62;
        }();
        _nw464[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_2226_v9, p0);
        _nw464[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(p0);
        _nw464[(new BigNumber(5)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(6)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(7)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(8)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_2410_v145).dtor_cf49, _2226_v9);
        _nw464[(new BigNumber(10)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(11)).toNumber()] = (_2409_v143).Intersect(function () {
          let _coll63 = new _dafny.Set();
          for (const _compr_63 of (_dafny.MultiSet.fromElements(new BigNumber((_2412_v147).length))).Elements) {
            let _2415_v148 = _compr_63;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_2412_v147).length))).contains(_2415_v148)) {
              _coll63.add((_2415_v148).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(true))).length)));
            }
          }
          return _coll63;
        }());
        _nw464[(new BigNumber(12)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(13)).toNumber()] = function () {
          let _coll64 = new _dafny.Set();
          for (const _compr_64 of (_dafny.Seq.of(_2226_v9)).Elements) {
            let _2416_v149 = _compr_64;
            if (_dafny.Seq.contains(_dafny.Seq.of(_2226_v9), _2416_v149)) {
              _coll64.add((_2416_v149).plus((_module.D20.create_DC48(false, _dafny.Seq.of(true, false), _dafny.Seq.of(!(true), !(true), true), new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).dtor_cf77));
            }
          }
          return _coll64;
        }();
        _nw464[(new BigNumber(14)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(15)).toNumber()] = _2409_v143;
        _nw464[(new BigNumber(16)).toNumber()] = _2409_v143;
        _2413_v150 = _nw464;
        let _nw465 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
        _2413_v150 = _nw465;
        let _2417_v151;
        let _init89 = ((_2418_v0) => function (_2419_i16) {
          return _2418_v0;
        })(_2211_v0);
        let _nw466 = Array((new BigNumber(13)).toNumber());
        for (let _i0_89 = 0; _i0_89 < new BigNumber(_nw466.length); _i0_89++) {
          _nw466[_i0_89] = _init89(new BigNumber(_i0_89));
        }
        _2417_v151 = _nw466;
        let _index459 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_2417_v151).length));
        (_2417_v151)[_index459] = !(_2211_v0);
        let _2420_v152;
        let _nw467 = Array((new BigNumber(13)).toNumber()).fill(_module.D25.Default());
        _2420_v152 = _nw467;
        let _2421_v153;
        _2421_v153 = _dafny.Seq.of(_2420_v152);
        let _index460 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_2417_v151).length));
        (_2417_v151)[_index460] = !_dafny.Seq.contains(_2421_v153, _2420_v152);
        let _2422_v154;
        _2422_v154 = _module.D16.create_DC38((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))], (_2417_v151)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_2417_v151).length))], p1, (_2417_v151)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_2417_v151).length))]);
        let _2423_v155;
        _2423_v155 = _dafny.MultiSet.fromElements(_module.__default.fm41(_2211_v0, new BigNumber((_dafny.MultiSet.fromElements(_2226_v9, _2226_v9, p0)).cardinality()), new BigNumber(974), new BigNumber(((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]).length), globalState), _2422_v154, _2422_v154);
        _2211_v0 = !(_2423_v155).equals(_2423_v155);
        r0 = _2211_v0;
        let _2424_v156;
        _2424_v156 = _dafny.Set.fromElements(_2417_v151);
        let _2425_v157;
        _2425_v157 = _dafny.Map.Empty.slice().updateUnsafe((_2417_v151)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_2417_v151).length))],_2424_v156);
        let _2426_v158;
        let _2427_v159;
        let _2428_v160;
        let _2429_v161;
        let _out25;
        let _out26;
        let _out27;
        let _out28;
        let _outcollector7 = (_this).m5(_2221_v7, _2425_v157, (p1).minus(_2226_v9), globalState);
        _out25 = _outcollector7[0];
        _out26 = _outcollector7[1];
        _out27 = _outcollector7[2];
        _out28 = _outcollector7[3];
        _2426_v158 = _out25;
        _2427_v159 = _out26;
        _2428_v160 = _out27;
        _2429_v161 = _out28;
      }
      let _2430_v162;
      _2430_v162 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2221_v7);
      let _2431_v163;
      _2431_v163 = _dafny.MultiSet.fromElements(new BigNumber(((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]).length), _2226_v9, _2226_v9, _2226_v9, new BigNumber((_2430_v162).length));
      let _2432_v164;
      _2432_v164 = _module.D8.create_DC18(_2431_v163);
      let _source24 = _2432_v164;
      if (_source24.is_DC19) {
        let _2433___mcc_h14 = (_source24).cf29;
        let _2434___mcc_h15 = (_source24).cf30;
        let _2435_cf30 = _2434___mcc_h15;
        let _2436_cf29 = _2433___mcc_h14;
        _module.__default.m0(new BigNumber(((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]).length), _2435_cf30, globalState);
        _2211_v0 = _2435_cf30;
        _2226_v9 = _module.__default.safeDivisionInt(p1, _2436_cf29);
        let _2437_v165;
        _2437_v165 = _dafny.Map.Empty.slice().updateUnsafe(_2229_v12,false);
        _2437_v165 = (_2437_v165).update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-998)), ((_2438_p1) => function (_2439_i17) {
          return _2438_p1;
        })(p1)), _2229_v12), !_dafny.areEqual(_2229_v12, _2229_v12));
      } else {
        let _2440___mcc_h16 = (_source24).cf28;
        let _2441_cf28 = _2440___mcc_h16;
        let _2442_v166;
        let _nw468 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _2442_v166 = _nw468;
        let _2443_v167;
        _2443_v167 = _dafny.Seq.of(_2211_v0);
        let _index461 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
        (_2442_v166)[_index461] = (_2226_v9).multipliedBy(new BigNumber((_2443_v167).length));
        let _index462 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2442_v166).length));
        (_2442_v166)[_index462] = p1;
        let _2444_v168;
        _2444_v168 = _module.D7.create_DC17(_2211_v0, _2211_v0, _2211_v0, _2211_v0);
        let _index463 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
        let _index464 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2442_v166).length));
        let _rhs452 = !(_2211_v0);
        let _rhs453 = (((_2444_v168).dtor_cf24) ? (_2211_v0) : (_2211_v0));
        let _rhs454 = new BigNumber(-611);
        let _rhs455 = _module.__default.safeModuloInt(_2226_v9, p1);
        let _lhs235 = _2442_v166;
        let _lhs236 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
        let _lhs237 = _2442_v166;
        let _lhs238 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_2442_v166).length));
        _2211_v0 = _rhs452;
        _2211_v0 = _rhs453;
        _lhs235[_lhs236] = _rhs454;
        _lhs237[_lhs238] = _rhs455;
        if (_2211_v0) {
          let _index465 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
          let _rhs456 = _2442_v166;
          let _rhs457 = _module.__default.fm14(globalState);
          let _rhs458 = false;
          let _rhs459 = (_2443_v167)[_module.__default.safeIndex(_2226_v9, new BigNumber((_2443_v167).length))];
          let _lhs239 = _2442_v166;
          let _lhs240 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
          _2442_v166 = _rhs456;
          _lhs239[_lhs240] = _rhs457;
          r0 = _rhs458;
          r0 = _rhs459;
          let _2445_v169;
          let _init90 = ((_2446_v7, _2447_v0, _2448_v167) => function (_2449_i18) {
            return _dafny.Seq.of(_module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe((_module.D23.create_DC55(_2446_v7, _2447_v0, true, _2447_v0, _2448_v167)).dtor_cf88,(_this).f5)), _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(_2447_v0,(_this).f5)), _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f5)), _module.D5.create_DC11(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('u'.codePointAt(0)))));
          })(_2221_v7, _2211_v0, _2443_v167);
          let _nw469 = Array((new BigNumber(8)).toNumber());
          for (let _i0_90 = 0; _i0_90 < new BigNumber(_nw469.length); _i0_90++) {
            _nw469[_i0_90] = _init90(new BigNumber(_i0_90));
          }
          _2445_v169 = _nw469;
          let _2450_v170;
          _2450_v170 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,(_this).f5);
          let _2451_v171;
          _2451_v171 = _module.D5.create_DC11(_2450_v170);
          let _2452_v172;
          _2452_v172 = _dafny.Seq.of(_2451_v171, _2451_v171, _module.D5.create_DC11(_2450_v170));
          let _index466 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_2445_v169).length));
          (_2445_v169)[_index466] = _2452_v172;
          let _pat_let_tv55 = _2450_v170;
          let _pat_let_tv56 = _2211_v0;
          let _index467 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_2445_v169).length));
          (_2445_v169)[_index467] = _dafny.Seq.Concat(_dafny.Seq.of(_2451_v171), _dafny.Seq.of(_2451_v171, _2451_v171, _2451_v171, function (_pat_let55_0) {
            return function (_2453_dt__update__tmp_h2) {
              return function (_pat_let56_0) {
                return function (_2454_dt__update_hcf15_h0) {
                  return _module.D5.create_DC11(_2454_dt__update_hcf15_h0);
                }(_pat_let56_0);
              }((_pat_let_tv55).update(_pat_let_tv56, (_this).f5));
            }(_pat_let55_0);
          }(_2451_v171)));
          _2211_v0 = !(_2211_v0);
          let _2455_v173;
          _2455_v173 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,!(_2211_v0));
          _2455_v173 = (_2455_v173).update(_2211_v0, _2211_v0);
          let _2456_v174;
          let _nw470 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2456_v174 = _nw470;
          let _2457_v175;
          _2457_v175 = _dafny.Set.fromElements(_2456_v174, _2456_v174);
          let _2458_v176;
          _2458_v176 = _2457_v175;
          let _2459_v177;
          _2459_v177 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v0,_2458_v176);
          let _2460_v178;
          _2460_v178 = _dafny.Seq.of(_2459_v177);
          let _2461_v179;
          let _2462_v180;
          let _2463_v181;
          let _2464_v182;
          let _out29;
          let _out30;
          let _out31;
          let _out32;
          let _outcollector8 = (_this).m5(_module.__default.fm12(globalState), (_2460_v178)[_module.__default.safeIndex(p1, new BigNumber((_2460_v178).length))], p0, globalState);
          _out29 = _outcollector8[0];
          _out30 = _outcollector8[1];
          _out31 = _outcollector8[2];
          _out32 = _outcollector8[3];
          _2461_v179 = _out29;
          _2462_v180 = _out30;
          _2463_v181 = _out31;
          _2464_v182 = _out32;
        } else {
          let _index468 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
          (_2442_v166)[_index468] = (_this).fm4(_module.__default.fm31(new BigNumber(623), p0, globalState), globalState);
          let _2465_v183;
          let _nw471 = Array((new BigNumber(14)).toNumber()).fill(false);
          _2465_v183 = _nw471;
          let _index469 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2465_v183).length));
          (_2465_v183)[_index469] = _2211_v0;
          let _2466_v184;
          _2466_v184 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]).length),p0);
          let _index470 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2465_v183).length));
          (_2465_v183)[_index470] = !(_module.__default.safeModuloInt((_2442_v166)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length))], (_2442_v166)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length))])).isEqualTo((((_2466_v184).contains((_2442_v166)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length))])) ? ((_2466_v184).get((_2442_v166)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length))])) : (p1)));
          let _index471 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_2465_v183).length));
          (_2465_v183)[_index471] = (_2211_v0) || (_2211_v0);
          let _index472 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_2442_v166).length));
          (_2442_v166)[_index472] = (_2229_v12)[_module.__default.safeIndex((_dafny.ZERO).minus((p1).multipliedBy(_2226_v9)), new BigNumber((_2229_v12).length))];
          _2442_v166 = _2442_v166;
        }
        let _2467_v185;
        let _nw472 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
        _2467_v185 = _nw472;
        let _rhs460 = _2211_v0;
        let _rhs461 = _2467_v185;
        r0 = _rhs460;
        _2467_v185 = _rhs461;
        (globalState).f2 = ((_2211_v0) ? (_dafny.Seq.Concat(_2229_v12, _2229_v12)) : (_2229_v12));
      }
      let _2468_v186;
      let _nw473 = Array((new BigNumber(8)).toNumber()).fill(false);
      _2468_v186 = _nw473;
      let _2469_v187;
      _2469_v187 = _dafny.Seq.of(_2468_v186);
      r0 = !(!_dafny.Seq.contains(_dafny.Seq.Concat(_2469_v187, _dafny.Seq.update(_2469_v187, _module.__default.safeIndex(p1, new BigNumber((_2469_v187).length)), _2468_v186)), _2468_v186));
      let _2470_v188;
      _2470_v188 = _dafny.Map.Empty.slice().updateUnsafe(_2468_v186,(_2222_v8)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2222_v8).length))]);
      r1 = (((_2470_v188).contains(_2468_v186)) ? ((_2470_v188).get(_2468_v186)) : (_module.__default.fm12(globalState)));
      r2 = _2431_v163;
      r3 = _module.__default.fm0((_dafny.ZERO).minus((p0).plus(_2226_v9)), p0, globalState);
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.Map.Empty;
      let _2471_v0;
      let _init91 = function (_2472_i0) {
        return _module.__default.safeDivisionInt(_2472_i0, new BigNumber(277));
      };
      let _nw474 = Array((new BigNumber(13)).toNumber());
      for (let _i0_91 = 0; _i0_91 < new BigNumber(_nw474.length); _i0_91++) {
        _nw474[_i0_91] = _init91(new BigNumber(_i0_91));
      }
      _2471_v0 = _nw474;
      _2471_v0 = _2471_v0;
      r1 = (p2).plus(new BigNumber((p0).length));
      let _2473_v1;
      _2473_v1 = new _dafny.CodePoint('k'.codePointAt(0));
      _2473_v1 = _2473_v1;
      let _2474_v2;
      _2474_v2 = _dafny.Seq.of(p2);
      if (!_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_2474_v2, _2474_v2), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("uf")).length), new BigNumber((_dafny.Seq.Concat(_2474_v2, _2474_v2)).length)), p2), _2474_v2)) {
        let _2475_v3;
        _2475_v3 = true;
        let _2476_v4;
        _2476_v4 = _dafny.Seq.of(false, _2475_v3, _2475_v3, _2475_v3);
        let _2477_v5;
        _2477_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v3,_2476_v4);
        let _2478_v6;
        _2478_v6 = _dafny.MultiSet.fromElements((((_2477_v5).contains(_2475_v3)) ? ((_2477_v5).get(_2475_v3)) : (_2476_v4)), _2476_v4);
        let _index473 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length));
        (_2471_v0)[_index473] = new BigNumber((_2478_v6).cardinality());
        let _index474 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length));
        (_2471_v0)[_index474] = (((_dafny.ZERO).minus(p2)).plus((_dafny.ZERO).minus(p2))).multipliedBy(new BigNumber(-393));
        let _2479_v7;
        let _nw475 = Array((new BigNumber(3)).toNumber());
        _nw475[(_dafny.ZERO).toNumber()] = (_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))];
        _nw475[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))]);
        _nw475[(new BigNumber(2)).toNumber()] = (_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))];
        _2479_v7 = _nw475;
        let _2480_v8;
        _2480_v8 = _dafny.Seq.of(_2479_v7, _2479_v7, _2479_v7, _2479_v7, _2479_v7);
        _2480_v8 = _dafny.Seq.of(_2479_v7);
        r1 = (_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))];
        let _2481_v9;
        let _nw476 = Array((new BigNumber(26)).toNumber()).fill(false);
        _2481_v9 = _nw476;
        let _index475 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2481_v9).length));
        (_2481_v9)[_index475] = !(p2).isEqualTo((_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))]);
        let _index476 = _module.__default.safeIndex(new BigNumber(827), new BigNumber((_2481_v9).length));
        (_2481_v9)[_index476] = (_module.__default.safeModuloInt((_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))], (_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))])).isLessThan(p2);
        let _index477 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length));
        (_2471_v0)[_index477] = (_2471_v0)[_module.__default.safeIndex(new BigNumber(310), new BigNumber((_2471_v0).length))];
      } else {
        r1 = _module.__default.safeDivisionInt(p2, new BigNumber(408));
        let _2482_v10;
        _2482_v10 = true;
        r2 = _2482_v10;
        let _2483_v11;
        _2483_v11 = _dafny.MultiSet.fromElements(p2);
        let _2484_v12;
        _2484_v12 = _dafny.Set.fromElements(_2483_v11, _2483_v11, _dafny.MultiSet.FromArray(_2474_v2));
        let _2485_v13;
        _2485_v13 = _dafny.Seq.of(_2484_v12);
        let _2486_v14;
        _2486_v14 = _dafny.Seq.of(_module.__default.fm22(_2473_v1, _2482_v10, new BigNumber(((_2485_v13)[_module.__default.safeIndex(p2, new BigNumber((_2485_v13).length))]).length), globalState));
        let _2487_v15;
        _2487_v15 = _dafny.Seq.UnicodeFromString("gq");
        let _rhs462 = _2482_v10;
        let _rhs463 = _2486_v14;
        let _rhs464 = _2487_v15;
        _2482_v10 = _rhs462;
        _2486_v14 = _rhs463;
        _2487_v15 = _rhs464;
        let _2488_v16;
        _2488_v16 = _dafny.Seq.of(_2482_v10, _2482_v10);
        let _2489_v17;
        _2489_v17 = _dafny.Set.fromElements(!(_2482_v10));
        let _2490_v18;
        let _nw477 = Array((new BigNumber(7)).toNumber());
        _nw477[(_dafny.ZERO).toNumber()] = _2488_v16;
        _nw477[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2488_v16, _2488_v16);
        _nw477[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_2488_v16, _2488_v16);
        _nw477[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2488_v16, _2488_v16);
        _nw477[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_module.__default.fm21(new BigNumber((_dafny.Seq.UnicodeFromString("emr")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2482_v10,_2482_v10)).length), _2489_v17, globalState), _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_module.__default.fm21(new BigNumber((_dafny.Seq.UnicodeFromString("emr")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2482_v10,_2482_v10)).length), _2489_v17, globalState)).length)), _2482_v10);
        _nw477[(new BigNumber(5)).toNumber()] = _module.__default.fm21(p2, p2, _2489_v17, globalState);
        _nw477[(new BigNumber(6)).toNumber()] = _2488_v16;
        _2490_v18 = _nw477;
        let _index478 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2490_v18).length));
        (_2490_v18)[_index478] = _2488_v16;
        let _index479 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_2490_v18).length));
        (_2490_v18)[_index479] = _dafny.Seq.Concat(_dafny.Seq.Concat(_2488_v16, _2488_v16), _2488_v16);
        let _2491_v19;
        let _nw478 = Array((new BigNumber(6)).toNumber()).fill(false);
        _2491_v19 = _nw478;
        (globalState).f3 = _2491_v19;
      }
      let _2492_v20;
      let _nw479 = new _module.C0();
      _nw479.__ctor();
      _2492_v20 = _nw479;
      let _2493_v21;
      _2493_v21 = _dafny.Set.fromElements(_2492_v20, _2492_v20);
      let _2494_v22;
      _2494_v22 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_2493_v21).length));
      r3 = _2494_v22;
      let _hi19 = new BigNumber((_dafny.Seq.of((_this).fm4(_2474_v2, globalState), p2, new BigNumber(600))).length);
      for (let _2495_i1 = p2; _2495_i1.isLessThan(_hi19); _2495_i1 = _2495_i1.plus(_dafny.ONE)) {
        let _2496_v23;
        _2496_v23 = false;
        if (!(_dafny.areEqual(_dafny.Seq.of(p2), _2474_v2)) || ((_module.D7.create_DC17((_2492_v20).fm3(new BigNumber(846), p2, globalState), _2496_v23, _2496_v23, _2496_v23)).dtor_cf27)) {
          let _2497_v24;
          _2497_v24 = _dafny.Set.fromElements(_2495_i1);
          let _2498_v25;
          _2498_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber(277), _2495_i1, globalState),_2497_v24);
          let _2499_v26;
          _2499_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2498_v25).length),_dafny.Seq.of((_dafny.ZERO).minus(p2)));
          _2499_v26 = (_2499_v26).update(_2495_i1, _2474_v2);
          let _2500_v27;
          _2500_v27 = _dafny.MultiSet.fromElements(_2496_v23, _2496_v23);
          let _2501_v28;
          _2501_v28 = _dafny.Map.Empty.slice().updateUnsafe((((_2494_v22).contains(new BigNumber(932))) ? ((_2494_v22).get(new BigNumber(932))) : (_2495_i1)),_2500_v27);
          let _2502_v29;
          _2502_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2501_v28).length),_2496_v23);
          let _2503_v30;
          _2503_v30 = _dafny.MultiSet.fromElements(false, (((_2502_v29).contains(p2)) ? ((_2502_v29).get(p2)) : (!(_2496_v23))), _2496_v23, !((new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2496_v23,_2495_i1)).length), new BigNumber((p0).length)), new _dafny.CodePoint('c'.codePointAt(0)))).length)).isLessThan(new BigNumber((_2474_v2).length))), _2496_v23);
          _2500_v27 = ((_2503_v30).Union(_2500_v27)).Union(_dafny.MultiSet.fromElements(_2496_v23));
          r2 = _2496_v23;
          let _2504_v31;
          let _nw480 = Array((new BigNumber(2)).toNumber()).fill(false);
          _2504_v31 = _nw480;
          let _index480 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_2504_v31).length));
          (_2504_v31)[_index480] = true;
          let _index481 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_2504_v31).length));
          (_2504_v31)[_index481] = _2496_v23;
          r1 = p2;
        } else {
          _2496_v23 = (new BigNumber((_2474_v2).length)).isEqualTo(((!(_2496_v23)) ? (new BigNumber((p0).length)) : (new BigNumber(-897))));
          r1 = p2;
          _2496_v23 = _2496_v23;
          r1 = p2;
          r1 = (_2495_i1).plus(p2);
        }
        _2496_v23 = _2496_v23;
        if (_2496_v23) {
          let _rhs465 = _2496_v23;
          let _rhs466 = (_dafny.Map.Empty.slice().updateUnsafe(_2495_i1,(_this).fm4(_2474_v2, globalState))).Merge(_2494_v22);
          let _rhs467 = _2473_v1;
          let _rhs468 = !(_2495_i1).isEqualTo((_dafny.ZERO).minus((_2495_i1).multipliedBy(p2)));
          let _rhs469 = _2496_v23;
          r2 = _rhs465;
          _2494_v22 = _rhs466;
          _2473_v1 = _rhs467;
          _2496_v23 = _rhs468;
          _2496_v23 = _rhs469;
          let _2505_v32;
          let _nw481 = new _module.C7();
          _nw481.__ctor();
          _2505_v32 = _nw481;
          let _2506_v33;
          let _nw482 = Array((new BigNumber(17)).toNumber());
          _nw482[(_dafny.ZERO).toNumber()] = _2505_v32;
          _nw482[(_dafny.ONE).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(2)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(3)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(4)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(5)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(6)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(7)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(8)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(9)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(10)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(11)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(12)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(13)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(14)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(15)).toNumber()] = _2505_v32;
          _nw482[(new BigNumber(16)).toNumber()] = _2505_v32;
          _2506_v33 = _nw482;
          let _2507_v34;
          _2507_v34 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2506_v33);
          let _2508_v35;
          let _nw483 = Array((new BigNumber(29)).toNumber());
          _nw483[(_dafny.ZERO).toNumber()] = _2506_v33;
          _nw483[(_dafny.ONE).toNumber()] = (((_2507_v34).contains((_dafny.ZERO).minus(_2495_i1))) ? ((_2507_v34).get((_dafny.ZERO).minus(_2495_i1))) : (_2506_v33));
          _nw483[(new BigNumber(2)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(3)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(4)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(5)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(6)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(7)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(8)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(9)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(10)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(11)).toNumber()] = (((_2507_v34).contains(_2495_i1)) ? ((_2507_v34).get(_2495_i1)) : (_2506_v33));
          _nw483[(new BigNumber(12)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(13)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(14)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(15)).toNumber()] = ((true) ? (_2506_v33) : (_2506_v33));
          _nw483[(new BigNumber(16)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(17)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(18)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(19)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(20)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(21)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(22)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(23)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(24)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(25)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(26)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(27)).toNumber()] = _2506_v33;
          _nw483[(new BigNumber(28)).toNumber()] = _2506_v33;
          _2508_v35 = _nw483;
          _2508_v35 = ((_2496_v23) ? (_2508_v35) : (_2508_v35));
          let _2509_v36;
          _2509_v36 = _dafny.Seq.of(_2496_v23, _2496_v23, _2496_v23, _2496_v23, _2496_v23);
          let _2510_v37;
          _2510_v37 = _module.D15.create_DC36(_module.D15.create_DC34(_2509_v36));
          let _2511_v38;
          _2511_v38 = _module.D15.create_DC34(_2509_v36);
          let _rhs470 = _2496_v23;
          let _rhs471 = ((_2496_v23) ? (_module.D15.create_DC36(_2511_v38)) : (_module.D15.create_DC36(_2511_v38)));
          let _rhs472 = _2496_v23;
          let _rhs473 = ((_2496_v23) ? (_2473_v1) : (new _dafny.CodePoint('y'.codePointAt(0))));
          r2 = _rhs470;
          _2510_v37 = _rhs471;
          _2496_v23 = _rhs472;
          _2473_v1 = _rhs473;
          let _index482 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_2471_v0).length));
          (_2471_v0)[_index482] = ((_2496_v23) ? (new BigNumber((p0).length)) : (_2495_i1));
          let _index483 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_2471_v0).length));
          (_2471_v0)[_index483] = (_dafny.ZERO).minus((_this).fm4(_2474_v2, globalState));
          r2 = _2496_v23;
        } else {
          r1 = _module.__default.safeModuloInt(_2495_i1, (new BigNumber((_2494_v22).length)).multipliedBy(new BigNumber(238)));
          let _2512_v39;
          let _nw484 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2512_v39 = _nw484;
          let _index484 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_2512_v39).length));
          (_2512_v39)[_index484] = _dafny.Seq.Concat(p0, p0);
          let _index485 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_2512_v39).length));
          (_2512_v39)[_index485] = _module.__default.fm12(globalState);
          _2471_v0 = _2471_v0;
          let _2513_v40;
          _2513_v40 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_2496_v23);
          _2513_v40 = (_2513_v40).update(p2, _2496_v23);
          r2 = _2496_v23;
        }
        r1 = _module.__default.safeModuloInt(_module.__default.fm13(_2496_v23, _2496_v23, _module.__default.fm1(_2495_i1, p2, false, _2496_v23, globalState), globalState), (new BigNumber((p0).length)).minus(_2495_i1));
      }
      let _2514_v41;
      _2514_v41 = true;
      let _2515_v42;
      _2515_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2514_v41,p2);
      let _2516_v43;
      _2516_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2514_v41,_module.__default.safeDivisionInt((((_2515_v42).contains(true)) ? ((_2515_v42).get(true)) : (p2)), p2));
      r0 = _2516_v43;
      r1 = p2;
      r2 = (p2).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), function (_2517_i2) {
        return (_this).f5;
      }), _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), function (_2518_i2) {
        return (_this).f5;
      })).length)), _2473_v1)).length));
      r3 = (_2494_v22).Merge(_2494_v22);
      return [r0, r1, r2, r3];
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.ZERO;
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f8, f9, f4, f5) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber(4)).multipliedBy((_this).f8)), (_this).f8);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("chgqrdurm"), _dafny.Seq.UnicodeFromString("dr")));
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f8;
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_dafny.ZERO).minus((_this).f8)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f9))).cardinality()),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kckfhdycy")).length),(_this).f9)).length))));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      if ((_this).f9) {
        r1 = (new BigNumber(-907)).multipliedBy((_this).f8);
        let _2519_v0;
        _2519_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_dafny.Seq.UnicodeFromString("fxaegu"));
        let _2520_v2;
        _2520_v2 = _module.D2.create_DC4((_this).f9);
        let _2521_v3;
        _2521_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_2519_v0).Merge(function () {
          let _coll65 = new _dafny.Map();
          for (const _compr_65 of _dafny.IntegerRange(new BigNumber(443), new BigNumber(341))) {
            let _2522_v1 = _compr_65;
            if (((new BigNumber(443)).isLessThanOrEqualTo(_2522_v1)) && ((_2522_v1).isLessThan(new BigNumber(341)))) {
              _coll65.push([_module.__default.safeDivisionInt(_2522_v1, new BigNumber(550)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), function (_2523_i0) {
                return (_this).f5;
              })]);
            }
          }
          return _coll65;
        }())).length)),function (_pat_let57_0) {
          return function (_2524_dt__update__tmp_h0) {
            return function (_pat_let58_0) {
              return function (_2525_dt__update_hcf5_h0) {
                return _module.D2.create_DC4(_2525_dt__update_hcf5_h0);
              }(_pat_let58_0);
            }((_this).f9);
          }(_pat_let57_0);
        }(_2520_v2));
        _2521_v3 = (_2521_v3).Merge(_2521_v3);
        let _2526_v4;
        _2526_v4 = _dafny.Seq.UnicodeFromString("lpeh");
        _2526_v4 = _2526_v4;
        r1 = (_this).f8;
        r1 = (_this).f8;
      } else {
        let _2527_v5;
        let _nw485 = Array((new BigNumber(26)).toNumber()).fill(false);
        _2527_v5 = _nw485;
        let _2528_v6;
        _2528_v6 = _dafny.Set.fromElements(_2527_v5);
        let _2529_v7;
        _2529_v7 = _2528_v6;
        let _2530_v8;
        _2530_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2529_v7,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), function (_2531_i1) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length));
        let _2532_v9;
        _2532_v9 = _dafny.Seq.of((_this).f5, (_this).f5);
        let _2533_v10;
        _2533_v10 = _dafny.Seq.of((_2530_v8).update(_2529_v7, new BigNumber((_2532_v9).length)));
        let _index486 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_2527_v5).length));
        (_2527_v5)[_index486] = false;
        let _2534_v11;
        _2534_v11 = false;
        let _2535_v12;
        _2535_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2534_v11,(_this).f5);
        let _2536_v13;
        _2536_v13 = _module.D5.create_DC11(_2535_v12);
        let _2537_v15;
        _2537_v15 = _dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC13(_2532_v9, (_this).f8, (_this).f9)).dtor_cf17,(_this).f8);
        let _2538_v16;
        _2538_v16 = _dafny.MultiSet.fromElements((_this).f8, new BigNumber((_2537_v15).length));
        let _2539_v17;
        _2539_v17 = _module.D2.create_DC4(!(_2534_v11));
        let _index487 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_2527_v5).length));
        let _rhs474 = _dafny.Seq.update(_dafny.Seq.of(_2530_v8), _module.__default.safeIndex((_this).fm8((_2536_v13).dtor_cf15, function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_2538_v16).Elements) {
            let _2540_v14 = _compr_66;
            if ((_2538_v16).contains(_2540_v14)) {
              _coll66.push([_module.__default.safeModuloInt(_2540_v14, (_dafny.ZERO).minus((_this).f8)),new BigNumber((_dafny.MultiSet.fromElements((_this).f5, (_this).f5, (_this).f5, (_this).f5)).cardinality())]);
            }
          }
          return _coll66;
        }(), _dafny.Seq.UnicodeFromString("ltvuc"), globalState), new BigNumber((_dafny.Seq.of(_2530_v8)).length)), _2530_v8);
        let _rhs475 = (_2539_v17).dtor_cf5;
        let _rhs476 = (_this).f9;
        let _rhs477 = _2534_v11;
        let _lhs241 = _2527_v5;
        let _lhs242 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_2527_v5).length));
        _2533_v10 = _rhs474;
        _lhs241[_lhs242] = _rhs475;
        _2534_v11 = _rhs476;
        _2534_v11 = _rhs477;
        let _2541_v18;
        let _nw486 = new _module.C0();
        _nw486.__ctor();
        _2541_v18 = _nw486;
        _2537_v15 = (_2537_v15).update((_this).f8, (_this).f8);
        _2534_v11 = (_2541_v18).fm3((_this).f8, new BigNumber(-990), globalState);
        _2534_v11 = !(_dafny.areEqual(_2532_v9, _2532_v9));
      }
      r1 = (_this).f8;
      let _2542_v19;
      _2542_v19 = true;
      let _2543_v20;
      let _nw487 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _2543_v20 = _nw487;
      let _index488 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length));
      (_2543_v20)[_index488] = new BigNumber(-895);
      let _2544_v21;
      _2544_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_2542_v19);
      let _2545_v22;
      _2545_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_2544_v21).length));
      let _2546_v23;
      _2546_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_2545_v22);
      let _2547_v24;
      _2547_v24 = _module.D14.create_DC30((_2546_v23).update((_this).f8, _2545_v22));
      let _2548_v25;
      _2548_v25 = _dafny.Seq.of(_2547_v24);
      let _2549_v26;
      _2549_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f5);
      let _2550_v27;
      _2550_v27 = _dafny.Seq.of(_2548_v25);
      let _index489 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length));
      let _rhs478 = (_this).fm3(_module.__default.safeDivisionInt((_this).f8, _module.__default.fm14(globalState)), (new BigNumber((_2549_v26).length)).plus((_this).f8), globalState);
      let _rhs479 = (_this).f8;
      let _rhs480 = (_2550_v27)[_module.__default.safeIndex((_this).f8, new BigNumber((_2550_v27).length))];
      let _rhs481 = (_module.__default.safeModuloInt((_this).f8, new BigNumber(229))).minus((_this).f8);
      let _lhs243 = _2543_v20;
      let _lhs244 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length));
      _2542_v19 = _rhs478;
      _lhs243[_lhs244] = _rhs479;
      _2548_v25 = _rhs480;
      r1 = _rhs481;
      let _2551_v28;
      let _nw488 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
      _2551_v28 = _nw488;
      let _2552_v29;
      _2552_v29 = _dafny.Seq.UnicodeFromString("rt");
      let _2553_v30;
      _2553_v30 = _dafny.Seq.of((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))], (((_2545_v22).contains(new BigNumber((_2552_v29).length))) ? ((_2545_v22).get(new BigNumber((_2552_v29).length))) : ((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))])), (_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]);
      let _index490 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_2551_v28).length));
      (_2551_v28)[_index490] = _2553_v30;
      let _2554_v31;
      _2554_v31 = _dafny.Map.Empty.slice().updateUnsafe((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))],_2552_v29);
      let _index491 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_2551_v28).length));
      (_2551_v28)[_index491] = _module.__default.fm31(new BigNumber(((((_2554_v31).contains((_this).f8)) ? ((_2554_v31).get((_this).f8)) : (_2552_v29))).length), new BigNumber(690), globalState);
      let _2555_v32;
      _2555_v32 = _dafny.Seq.of(_2542_v19);
      let _2556_v33;
      _2556_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2555_v32,(_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]);
      _2556_v33 = ((_dafny.Map.Empty.slice().updateUnsafe(_2555_v32,new BigNumber(891))).Merge(_2556_v33)).Merge((_2556_v33).update(_2555_v32, (_this).f8));
      let _2557_v34;
      let _nw489 = Array((new BigNumber(11)).toNumber()).fill(false);
      _2557_v34 = _nw489;
      let _2558_v35;
      _2558_v35 = _module.D18.create_DC45(_2557_v34, (_this).f9, (_this).f8, new BigNumber(-585));
      let _2559_v36;
      _2559_v36 = _dafny.Seq.of(_2557_v34, (_2558_v35).dtor_cf68);
      let _2560_i2;
      _2560_i2 = _dafny.ZERO;
      L9: {
        while (!((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]).isEqualTo(new BigNumber((_2559_v36).length))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2560_i2)) {
              break L9;
            }
            _2560_i2 = (_2560_i2).plus(_dafny.ONE);
            let _2561_v37;
            let _init92 = ((_2562_v29, _2563_v20) => function (_2564_i3) {
              return (_2562_v29)[_module.__default.safeIndex((_2563_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2563_v20).length))], new BigNumber((_2562_v29).length))];
            })(_2552_v29, _2543_v20);
            let _nw490 = Array((new BigNumber(20)).toNumber());
            for (let _i0_92 = 0; _i0_92 < new BigNumber(_nw490.length); _i0_92++) {
              _nw490[_i0_92] = _init92(new BigNumber(_i0_92));
            }
            _2561_v37 = _nw490;
            _2561_v37 = _2561_v37;
            let _2565_v38;
            _2565_v38 = _dafny.MultiSet.fromElements((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))], (_this).f8);
            let _index492 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length));
            let _rhs482 = _2543_v20;
            let _rhs483 = (_2565_v38).IsProperSubsetOf(_2565_v38);
            let _rhs484 = ((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]).multipliedBy((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]);
            let _lhs245 = _2543_v20;
            let _lhs246 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length));
            _2543_v20 = _rhs482;
            _2542_v19 = _rhs483;
            _lhs245[_lhs246] = _rhs484;
            let _2566_v39;
            _2566_v39 = _dafny.Seq.of(((_2542_v19) ? (_2553_v30) : (_2553_v30)));
            _2566_v39 = _dafny.Seq.of(_2553_v30, _dafny.Seq.update(_2553_v30, _module.__default.safeIndex((_dafny.ZERO).minus((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]), new BigNumber((_2553_v30).length)), (_this).f8), _2553_v30);
            let _2567_v40;
            _2567_v40 = new _dafny.CodePoint('o'.codePointAt(0));
            _2567_v40 = (_this).f5;
          }
        }
      }
      r0 = _dafny.Seq.update(_dafny.Seq.Concat((_2551_v28)[_module.__default.safeIndex(new BigNumber(650), new BigNumber((_2551_v28).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), function (_2568_i4) {
        return (_this).f8;
      })), _module.__default.safeIndex((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))], new BigNumber((_dafny.Seq.Concat((_2551_v28)[_module.__default.safeIndex(new BigNumber(650), new BigNumber((_2551_v28).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), function (_2569_i4) {
        return (_this).f8;
      }))).length)), new BigNumber(330));
      let _2570_v41;
      _2570_v41 = _module.D11.create_DC25((_this).f9, (_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))], _2543_v20, (_this).f8);
      let _2571_v42;
      _2571_v42 = _module.D11.create_DC26(_2570_v41);
      let _2572_v43;
      _2572_v43 = _module.D11.create_DC26(_2571_v42);
      let _2573_v44;
      _2573_v44 = _module.D11.create_DC26(_2571_v42);
      let _2574_v45;
      _2574_v45 = _module.D11.create_DC26(_2572_v43);
      let _2575_v46;
      _2575_v46 = _dafny.MultiSet.fromElements((_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))], (_2543_v20)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_2543_v20).length))]);
      let _2576_v47;
      _2576_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2574_v45,_2575_v46);
      r1 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_2576_v47).Merge(_2576_v47)).length)))).plus((_this).f8);
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      if ((_this).f9) {
        if ((_this).f9) {
          let _2577_v0;
          let _init93 = function (_2578_i0) {
            return false;
          };
          let _nw491 = Array((new BigNumber(10)).toNumber());
          for (let _i0_93 = 0; _i0_93 < new BigNumber(_nw491.length); _i0_93++) {
            _nw491[_i0_93] = _init93(new BigNumber(_i0_93));
          }
          _2577_v0 = _nw491;
          let _2579_v1;
          _2579_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2577_v0,(_this).f9);
          _2579_v1 = (_2579_v1).update(_2577_v0, (_this).f9);
          let _2580_v2;
          _2580_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          let _2581_v3;
          _2581_v3 = _dafny.Seq.of((_this).f8, (_this).f8);
          let _2582_v5;
          _2582_v5 = _dafny.Seq.of(_2580_v2);
          let _2583_v6;
          let _nw492 = Array((new BigNumber(8)).toNumber());
          _nw492[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f8, (_this).f8, (_this).f8, (_dafny.ZERO).minus((_this).f8), _module.__default.fm14(globalState))).length);
          _nw492[(_dafny.ONE).toNumber()] = (_this).f8;
          _nw492[(new BigNumber(2)).toNumber()] = new BigNumber((_2580_v2).length);
          _nw492[(new BigNumber(3)).toNumber()] = (_this).f8;
          _nw492[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_this).f8, (_this).f8);
          _nw492[(new BigNumber(5)).toNumber()] = (_2581_v3)[_module.__default.safeIndex((_this).f8, new BigNumber((_2581_v3).length))];
          _nw492[(new BigNumber(6)).toNumber()] = new BigNumber(((function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of (_2582_v5).Elements) {
              let _2584_v4 = _compr_67;
              if (_dafny.Seq.contains(_2582_v5, _2584_v4)) {
                _coll67.push([_2584_v4,(_this).f8]);
              }
            }
            return _coll67;
          }()).Merge(_module.__default.fm47(true, (_this).f8, globalState))).length);
          _nw492[(new BigNumber(7)).toNumber()] = (_this).f8;
          _2583_v6 = _nw492;
          let _index493 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length));
          (_2583_v6)[_index493] = (_dafny.ZERO).minus(new BigNumber((function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of _dafny.IntegerRange(new BigNumber(172), new BigNumber(986))) {
              let _2585_v7 = _compr_68;
              if (((new BigNumber(172)).isLessThanOrEqualTo(_2585_v7)) && ((_2585_v7).isLessThan(new BigNumber(986)))) {
                _coll68.add(_module.__default.safeModuloInt(_2585_v7, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), function (_2586_i1) {
                  return (_this).f5;
                }))).cardinality())));
              }
            }
            return _coll68;
          }()).length));
          let _index494 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length));
          (_2583_v6)[_index494] = ((_dafny.ZERO).minus((_this).f8)).multipliedBy((_this).f8);
          let _index495 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length));
          (_2583_v6)[_index495] = (_dafny.ZERO).minus((_2583_v6)[_module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length))]);
          let _2587_v8;
          _2587_v8 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(340), (_this).f8));
          let _2588_v9;
          _2588_v9 = _dafny.Seq.of((_this).f9, (new BigNumber((_module.__default.fm29((_2583_v6)[_module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length))], (_this).f9, (_this).f8, globalState)).length)).isLessThanOrEqualTo(new BigNumber((_2587_v8).length)), (_this).f9);
          _2588_v9 = _2588_v9;
          let _index496 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length));
          (_2583_v6)[_index496] = (new BigNumber(-975)).multipliedBy((_2583_v6)[_module.__default.safeIndex(new BigNumber(62), new BigNumber((_2583_v6).length))]);
        } else {
          let _2589_v10;
          _2589_v10 = _dafny.Seq.of((_this).f9, (_this).f9);
          let _2590_v11;
          _2590_v11 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9)).length));
          let _2591_v12;
          _2591_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          let _2592_v13;
          _2592_v13 = _dafny.Seq.of((_this).f5, (_this).f5);
          let _2593_v14;
          _2593_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_2592_v13).length));
          let _2594_v15;
          let _nw493 = new _module.C6();
          _nw493.__ctor(_2593_v14);
          _2594_v15 = _nw493;
          let _2595_v16;
          _2595_v16 = _dafny.Set.fromElements(_2594_v15, _2594_v15);
          let _2596_v17;
          _2596_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f8);
          let _2597_v18;
          _2597_v18 = _dafny.MultiSet.fromElements(new BigNumber((_2596_v17).length));
          let _2598_v19;
          _2598_v19 = _dafny.Seq.of((_this).f8, (_this).f8);
          let _2599_v20;
          let _init94 = function (_2600_i2) {
            return (_this).f9;
          };
          let _nw494 = Array((new BigNumber(24)).toNumber());
          for (let _i0_94 = 0; _i0_94 < new BigNumber(_nw494.length); _i0_94++) {
            _nw494[_i0_94] = _init94(new BigNumber(_i0_94));
          }
          _2599_v20 = _nw494;
          let _2601_v21;
          _2601_v21 = _dafny.Seq.of(_2599_v20);
          let _2602_v22;
          _2602_v22 = _module.D2.create_DC4((_this).f9);
          let _2603_v23;
          let _nw495 = Array((new BigNumber(14)).toNumber());
          _nw495[(_dafny.ZERO).toNumber()] = !(!((_this).f9));
          _nw495[(_dafny.ONE).toNumber()] = (new BigNumber((_2589_v10).length)).isLessThan((_this).f8);
          _nw495[(new BigNumber(2)).toNumber()] = ((_this).f8).isLessThan((_this).f8);
          _nw495[(new BigNumber(3)).toNumber()] = (_2590_v11).IsSubsetOf(_2590_v11);
          _nw495[(new BigNumber(4)).toNumber()] = (((_2591_v12).contains((_this).f9)) ? ((_2591_v12).get((_this).f9)) : ((_this).f9));
          _nw495[(new BigNumber(5)).toNumber()] = !(_2595_v16).contains(_2594_v15);
          _nw495[(new BigNumber(6)).toNumber()] = ((_this).f9) && ((_this).f9);
          _nw495[(new BigNumber(7)).toNumber()] = (new BigNumber((_2597_v18).cardinality())).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f8));
          _nw495[(new BigNumber(8)).toNumber()] = !(!(((_this).f8).isLessThanOrEqualTo((_2598_v19)[_module.__default.safeIndex((_this).f8, new BigNumber((_2598_v19).length))])));
          _nw495[(new BigNumber(9)).toNumber()] = !_dafny.Seq.contains(_2601_v21, _2599_v20);
          _nw495[(new BigNumber(10)).toNumber()] = (_this).f9;
          _nw495[(new BigNumber(11)).toNumber()] = (_this).f9;
          _nw495[(new BigNumber(12)).toNumber()] = (_2602_v22).dtor_cf5;
          _nw495[(new BigNumber(13)).toNumber()] = false;
          _2603_v23 = _nw495;
          let _2604_v24;
          _2604_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2598_v19);
          let _index497 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length));
          (_2603_v23)[_index497] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2598_v19)).equals(_2604_v24);
          let _index498 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length));
          (_2603_v23)[_index498] = (_this).f9;
          let _2605_v25;
          _2605_v25 = _dafny.MultiSet.fromElements(((_2603_v23)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length))]) === ((_this).f9), (_this).f9, (_this).f9);
          _2605_v25 = (_2605_v25).Difference(_2605_v25);
          let _2606_v26;
          _2606_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2593_v14).Merge(_2593_v14),(_this).f9);
          _2606_v26 = (_2606_v26).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_2592_v13).length)), (_2603_v23)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length))]);
          let _2607_v27;
          _2607_v27 = new BigNumber(569);
          let _2608_v28;
          let _nw496 = new _module.C0();
          _nw496.__ctor();
          _2608_v28 = _nw496;
          let _2609_v29;
          let _init95 = ((_2610_v13) => function (_2611_i3) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), function (_2612_i4) {
              return (_this).f5;
            }), _2610_v13);
          })(_2592_v13);
          let _nw497 = Array((new BigNumber(9)).toNumber());
          for (let _i0_95 = 0; _i0_95 < new BigNumber(_nw497.length); _i0_95++) {
            _nw497[_i0_95] = _init95(new BigNumber(_i0_95));
          }
          _2609_v29 = _nw497;
          let _index499 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_2609_v29).length));
          (_2609_v29)[_index499] = _2592_v13;
          let _index500 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_2609_v29).length));
          let _rhs485 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2598_v19, _module.__default.safeIndex(new BigNumber((_2598_v19).length), new BigNumber((_2598_v19).length)), (_dafny.ZERO).minus(_2607_v27)), _module.__default.fm31(_2607_v27, (_this).f8, globalState))).length);
          let _rhs486 = _2608_v28;
          let _rhs487 = (((_this).f9) ? (((_module.__default.fm1(_2607_v27, (_this).f8, (_this).f9, true, globalState)) ? (_2592_v13) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_2613_i5) {
            return (_this).f5;
          })))) : (_2592_v13));
          let _lhs247 = _2609_v29;
          let _lhs248 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_2609_v29).length));
          _2607_v27 = _rhs485;
          _2608_v28 = _rhs486;
          _lhs247[_lhs248] = _rhs487;
          let _index501 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length));
          let _rhs488 = (((_this).f9) ? ((((_this).f9) ? ((_this).f9) : ((_this).f9))) : (true));
          let _rhs489 = _dafny.Seq.Concat(_2598_v19, _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_2590_v11).length)), _2607_v27, (_dafny.ZERO).minus(_2607_v27)));
          let _rhs490 = _2607_v27;
          let _lhs249 = _2603_v23;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_2603_v23).length));
          let _lhs251 = globalState;
          _lhs249[_lhs250] = _rhs488;
          _lhs251.f2 = _rhs489;
          _2607_v27 = _rhs490;
        }
        if (!(_module.__default.fm1((_this).f8, (_this).f8, (_this).f9, (_this).f9, globalState)) || ((_this).f9)) {
          let _2614_v30;
          _2614_v30 = new BigNumber(729);
          let _2615_v31;
          _2615_v31 = _dafny.Seq.UnicodeFromString("gc");
          let _2616_v32;
          let _nw498 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _2616_v32 = _nw498;
          let _2617_v33;
          _2617_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2614_v30,_2616_v32);
          let _2618_v34;
          _2618_v34 = _dafny.Seq.of(_2616_v32);
          let _2619_v35;
          let _nw499 = Array((new BigNumber(23)).toNumber());
          _nw499[(_dafny.ZERO).toNumber()] = (((_2617_v33).contains(_2614_v30)) ? ((_2617_v33).get(_2614_v30)) : (_2616_v32));
          _nw499[(_dafny.ONE).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(2)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(3)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(4)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(5)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(6)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(7)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(8)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(9)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(10)).toNumber()] = (_2618_v34)[_module.__default.safeIndex(_2614_v30, new BigNumber((_2618_v34).length))];
          _nw499[(new BigNumber(11)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(12)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(13)).toNumber()] = (((_this).f9) ? (_2616_v32) : (_2616_v32));
          _nw499[(new BigNumber(14)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(15)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(16)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(17)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(18)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(19)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(20)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(21)).toNumber()] = _2616_v32;
          _nw499[(new BigNumber(22)).toNumber()] = _2616_v32;
          _2619_v35 = _nw499;
          let _index502 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2619_v35).length));
          (_2619_v35)[_index502] = _2616_v32;
          let _index503 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length));
          (_2616_v32)[_index503] = (_dafny.ZERO).minus((_this).f8);
          let _2620_v36;
          _2620_v36 = new _dafny.CodePoint('t'.codePointAt(0));
          let _index504 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2619_v35).length));
          let _index505 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length));
          let _rhs491 = (_this).f8;
          let _rhs492 = _2615_v31;
          let _rhs493 = _2616_v32;
          let _rhs494 = (_this).f8;
          let _rhs495 = (_this).f5;
          let _lhs252 = _2619_v35;
          let _lhs253 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2619_v35).length));
          let _lhs254 = _2616_v32;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length));
          _2614_v30 = _rhs491;
          _2615_v31 = _rhs492;
          _lhs252[_lhs253] = _rhs493;
          _lhs254[_lhs255] = _rhs494;
          _2620_v36 = _rhs495;
          let _2621_v37;
          _2621_v37 = false;
          _2621_v37 = (!((_2616_v32)[_module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length))]).isEqualTo((_this).f8)) === (_2621_v37);
          _2621_v37 = ((_2616_v32)[_module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length))]).isLessThan((_2616_v32)[_module.__default.safeIndex(new BigNumber(32), new BigNumber((_2616_v32).length))]);
          let _2622_v38;
          let _nw500 = new _module.C9();
          _nw500.__ctor((_this).f4, (_this).f5);
          _2622_v38 = _nw500;
          let _2623_v39;
          _2623_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_dafny.ZERO).minus((_this).f8));
          _2623_v39 = _2623_v39;
        } else {
          let _2624_v40;
          _2624_v40 = _dafny.Seq.UnicodeFromString("fodfeqy");
          let _2625_v41;
          _2625_v41 = true;
          let _2626_v42;
          let _nw501 = new _module.C7();
          _nw501.__ctor();
          _2626_v42 = _nw501;
          let _2627_v43;
          _2627_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_2626_v42);
          let _2628_v44;
          _2628_v44 = _dafny.Seq.of(_2626_v42);
          let _2629_v45;
          _2629_v45 = _module.D16.create_DC38(_2624_v40, (_this).f9, new BigNumber(92), (_this).f9);
          let _2630_v46;
          _2630_v46 = _dafny.Map.Empty.slice().updateUnsafe((((_2627_v43).contains(_module.__default.fm13(_2625_v41, _2625_v41, false, globalState))) ? ((_2627_v43).get(_module.__default.fm13(_2625_v41, _2625_v41, false, globalState))) : ((_2628_v44)[_module.__default.safeIndex(new BigNumber(942), new BigNumber((_2628_v44).length))])),(_2629_v45).dtor_cf57);
          let _2631_v47;
          _2631_v47 = _dafny.Seq.of((_this).f9);
          let _2632_v48;
          _2632_v48 = _dafny.Set.fromElements((_this).f8);
          let _2633_v49;
          _2633_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2631_v47,_2632_v48);
          let _2634_v50;
          _2634_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f8);
          let _2635_v51;
          let _init96 = function (_2636_i7) {
            return (_2636_i7).minus((_this).f8);
          };
          let _nw502 = Array((new BigNumber(16)).toNumber());
          for (let _i0_96 = 0; _i0_96 < new BigNumber(_nw502.length); _i0_96++) {
            _nw502[_i0_96] = _init96(new BigNumber(_i0_96));
          }
          _2635_v51 = _nw502;
          let _rhs496 = _dafny.Seq.UnicodeFromString("cgbxxirqm");
          let _rhs497 = !(!(_2633_v49).equals(_2633_v49)) || (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("emovxn"), (((_2630_v46).contains(_2626_v42)) ? ((_2630_v46).get(_2626_v42)) : (_2624_v40))));
          let _rhs498 = (_2631_v47)[_module.__default.safeIndex(new BigNumber((_2634_v50).length), new BigNumber((_2631_v47).length))];
          let _rhs499 = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-134)), function (_2637_i6) {
            return (_this).f5;
          }), _2624_v40);
          let _rhs500 = _2635_v51;
          _2624_v40 = _rhs496;
          _2625_v41 = _rhs497;
          _2625_v41 = _rhs498;
          _2625_v41 = _rhs499;
          r0 = _rhs500;
          let _2638_v52;
          _2638_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2625_v41,((_this).f8).isLessThan((_this).f8));
          _2625_v41 = (((_2638_v52).contains((_this).f9)) ? ((_2638_v52).get((_this).f9)) : ((_this).f9));
          let _2639_v53;
          _2639_v53 = new _dafny.CodePoint('u'.codePointAt(0));
          let _2640_v54;
          _2640_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2639_v53);
          let _2641_v55;
          _2641_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(209));
          let _index506 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2635_v51).length));
          (_2635_v51)[_index506] = ((_this).f8).multipliedBy((_this).fm8(_2640_v54, _2641_v55, _module.__default.fm12(globalState), globalState));
          let _index507 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2635_v51).length));
          let _rhs501 = _2635_v51;
          let _rhs502 = _2639_v53;
          let _rhs503 = (_this).f8;
          let _lhs256 = _2635_v51;
          let _lhs257 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2635_v51).length));
          _2635_v51 = _rhs501;
          _2639_v53 = _rhs502;
          _lhs256[_lhs257] = _rhs503;
          _2625_v41 = _2625_v41;
          let _index508 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2635_v51).length));
          (_2635_v51)[_index508] = _module.__default.safeDivisionInt((_this).f8, (_dafny.ZERO).minus((new BigNumber((_2632_v48).length)).multipliedBy(new BigNumber(146))));
        }
        let _2642_v56;
        _2642_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_module.__default.fm14(globalState));
        _module.__default.m0((_this).f8, ((_this).f8).isLessThan(new BigNumber((_2642_v56).length)), globalState);
        let _2643_v57;
        _2643_v57 = _module.D1.create_DC1(_module.__default.fm15((_this).f9, (_this).f9, (_this).f8, globalState));
        let _source25 = (((_this).f9) ? (_2643_v57) : (_2643_v57));
        if (_source25.is_DC2) {
          let _2644___mcc_h0 = (_source25).cf2;
          let _2645___mcc_h1 = (_source25).cf3;
          let _2646_cf3 = _2645___mcc_h1;
          let _2647_cf2 = _2644___mcc_h0;
          let _2648_v58;
          let _init97 = ((_2649_cf2) => function (_2650_i8) {
            return _module.__default.safeModuloInt(_2650_i8, _2649_cf2);
          })(_2647_cf2);
          let _nw503 = Array((new BigNumber(4)).toNumber());
          for (let _i0_97 = 0; _i0_97 < new BigNumber(_nw503.length); _i0_97++) {
            _nw503[_i0_97] = _init97(new BigNumber(_i0_97));
          }
          _2648_v58 = _nw503;
          r0 = _2648_v58;
          _module.__default.m0(_2647_cf2, (_this).f9, globalState);
          let _2651_v59;
          let _nw504 = new _module.C8();
          _nw504.__ctor();
          _2651_v59 = _nw504;
          let _2652_v60;
          _2652_v60 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f9) || ((_this).f9),_2651_v59);
          let _2653_v61;
          _2653_v61 = _dafny.Seq.UnicodeFromString("s");
          let _2654_v62;
          _2654_v62 = _dafny.Seq.of((_this).f9);
          let _2655_v63;
          _2655_v63 = _module.D23.create_DC55(_2653_v61, (_this).f9, (_this).f9, (_this).f9, _2654_v62);
          _2651_v59 = (((_2652_v60).contains(((!((_2655_v63).dtor_cf88)) ? ((_this).f9) : ((_this).f9)))) ? ((_2652_v60).get(((!((_2655_v63).dtor_cf88)) ? ((_this).f9) : ((_this).f9)))) : (_2651_v59));
          let _2656_v64;
          _2656_v64 = new _dafny.CodePoint('x'.codePointAt(0));
          _2656_v64 = _2656_v64;
        } else {
          let _2657___mcc_h2 = (_source25).cf1;
          let _2658_cf1 = _2657___mcc_h2;
          let _2659_v65;
          _2659_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f8);
          _2659_v65 = (_2659_v65).Merge((((_this).f9) ? (_2659_v65) : (_2659_v65)));
          let _2660_v66;
          let _init98 = function (_2661_i9) {
            return !((_this).f9);
          };
          let _nw505 = Array((new BigNumber(25)).toNumber());
          for (let _i0_98 = 0; _i0_98 < new BigNumber(_nw505.length); _i0_98++) {
            _nw505[_i0_98] = _init98(new BigNumber(_i0_98));
          }
          _2660_v66 = _nw505;
          let _2662_v67;
          _2662_v67 = _dafny.Seq.UnicodeFromString("idq");
          let _index509 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2660_v66).length));
          (_2660_v66)[_index509] = !_dafny.Seq.contains(_2662_v67, _module.__default.fm0(new BigNumber(74), (_this).f8, globalState));
          let _2663_v68;
          _2663_v68 = _dafny.Set.fromElements(_2662_v67);
          let _2664_v69;
          _2664_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2660_v66,false);
          let _2665_v70;
          let _nw506 = Array((new BigNumber(24)).toNumber());
          _nw506[(_dafny.ZERO).toNumber()] = ((_this).f8).multipliedBy((_this).f8);
          _nw506[(_dafny.ONE).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(2)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_this).f8, (_this).f8);
          _nw506[(new BigNumber(4)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(5)).toNumber()] = new BigNumber(467);
          _nw506[(new BigNumber(6)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f8));
          _nw506[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2662_v67, _2662_v67)).length));
          _nw506[(new BigNumber(9)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2662_v67).length), (_dafny.ZERO).minus(new BigNumber((_2663_v68).length)));
          _nw506[(new BigNumber(11)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(12)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(13)).toNumber()] = new BigNumber(963);
          _nw506[(new BigNumber(14)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(15)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(16)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(17)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2664_v69).length));
          _nw506[(new BigNumber(19)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(20)).toNumber()] = (_this).f8;
          _nw506[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
          _nw506[(new BigNumber(22)).toNumber()] = ((_this).f8).minus((_dafny.ZERO).minus((_this).f8));
          _nw506[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
          _2665_v70 = _nw506;
          let _2666_v71;
          _2666_v71 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wu"));
          let _index510 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_2665_v70).length));
          (_2665_v70)[_index510] = new BigNumber((_2666_v71).length);
          let _2667_v72;
          _2667_v72 = new BigNumber(695);
          let _2668_v73;
          _2668_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2667_v72,_2660_v66);
          let _index511 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2660_v66).length));
          let _index512 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_2665_v70).length));
          let _rhs504 = (_this).f9;
          let _rhs505 = _dafny.Seq.update(_2662_v67, _module.__default.safeIndex((_this).f8, new BigNumber((_2662_v67).length)), (_this).f5);
          let _rhs506 = new BigNumber(((_2668_v73).Merge(_2668_v73)).length);
          let _rhs507 = ((_this).f8).minus(_2667_v72);
          let _lhs258 = _2660_v66;
          let _lhs259 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_2660_v66).length));
          let _lhs260 = _2665_v70;
          let _lhs261 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_2665_v70).length));
          _lhs258[_lhs259] = _rhs504;
          _2662_v67 = _rhs505;
          _lhs260[_lhs261] = _rhs506;
          _2667_v72 = _rhs507;
          _module.__default.m0((_this).f8, (_2660_v66)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_2660_v66).length))], globalState);
          let _2669_v74;
          _2669_v74 = _module.D17.create_DC42(_2660_v66, (_this).f8);
          _2659_v65 = (_2659_v65).update(new BigNumber((_dafny.Set.fromElements(_2669_v74, _2669_v74)).length), (_this).f8);
        }
        let _2670_v75;
        _2670_v75 = true;
        _2670_v75 = !((((_this).f9) ? ((_this).f9) : (_2670_v75)));
      } else {
        let _2671_v76;
        _2671_v76 = new BigNumber(303);
        _2671_v76 = (_dafny.ZERO).minus(((_dafny.ZERO).multipliedBy(_2671_v76)).multipliedBy(_2671_v76));
        let _2672_v77;
        let _nw507 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
        _2672_v77 = _nw507;
        let _2673_v78;
        _2673_v78 = _dafny.Seq.of(_2671_v76, (_this).f8, _2671_v76);
        let _index513 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_2672_v77).length));
        (_2672_v77)[_index513] = _2673_v78;
        let _index514 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_2672_v77).length));
        (_2672_v77)[_index514] = _dafny.Seq.Concat(_2673_v78, _2673_v78);
        let _2674_v79;
        _2674_v79 = _dafny.Seq.UnicodeFromString("y");
        let _2675_v80;
        _2675_v80 = _dafny.Seq.of((_this).f9);
        _2671_v76 = _module.__default.safeDivisionInt((_this).f8, (new BigNumber((_2674_v79).length)).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_2675_v80)).cardinality())));
        _2671_v76 = _module.__default.safeDivisionInt(_2671_v76, new BigNumber(861));
        let _2676_v81;
        _2676_v81 = _dafny.Set.fromElements((_this).f9, (_this).f9);
        let _2677_v82;
        _2677_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2676_v81);
        let _2678_v83;
        _2678_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,(_this).f9);
        let _2679_v84;
        let _nw508 = new _module.C0();
        _nw508.__ctor();
        _2679_v84 = _nw508;
        let _2680_v85;
        _2680_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_2679_v84);
        let _rhs508 = _module.__default.safeDivisionInt((_this).f8, _2671_v76);
        let _rhs509 = (_2676_v81).Difference(((((_2677_v82).contains((_this).f5)) ? ((_2677_v82).get((_this).f5)) : (_2676_v81))).Union(_dafny.Set.fromElements(!(_module.__default.fm1((_this).f8, _2671_v76, (_this).f9, false, globalState)), (((_2678_v83).contains(new BigNumber((_2680_v85).length))) ? ((_2678_v83).get(new BigNumber((_2680_v85).length))) : ((_2679_v84).fm18((_this).f9, _2671_v76, !((_this).f9), (_this).f8, globalState))))));
        _2671_v76 = _rhs508;
        _2676_v81 = _rhs509;
      }
      let _2681_v86;
      _2681_v86 = _dafny.Seq.UnicodeFromString("wgwfnpj");
      let _2682_v87;
      _2682_v87 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ehknsls"));
      let _2683_v88;
      _2683_v88 = _dafny.MultiSet.fromElements(_2681_v86, _2681_v86, (_2682_v87)[_module.__default.safeIndex(new BigNumber(-467), new BigNumber((_2682_v87).length))], _2681_v86, (_2682_v87)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_2682_v87).length))]);
      if (((_this).f9) && ((_2683_v88).IsDisjointFrom(_2683_v88))) {
        let _2684_v89;
        let _nw509 = Array((new BigNumber(6)).toNumber()).fill([]);
        _2684_v89 = _nw509;
        _2684_v89 = _2684_v89;
        let _2685_v90;
        let _nw510 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _2685_v90 = _nw510;
        let _2686_v91;
        _2686_v91 = _dafny.Map.Empty.slice().updateUnsafe(_2685_v90,_dafny.Seq.Concat(_2681_v86, _dafny.Seq.UnicodeFromString("lnuurg")));
        _2686_v91 = _2686_v91;
        let _2687_v92;
        _2687_v92 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(globalState),((_this).f9) === ((_this).f9));
        _2687_v92 = _2687_v92;
        let _2688_v93;
        _2688_v93 = new BigNumber(978);
        let _2689_v94;
        _2689_v94 = _dafny.MultiSet.fromElements(_2685_v90);
        let _2690_v95;
        let _init99 = function (_2691_i10) {
          return (_this).f9;
        };
        let _nw511 = Array((new BigNumber(28)).toNumber());
        for (let _i0_99 = 0; _i0_99 < new BigNumber(_nw511.length); _i0_99++) {
          _nw511[_i0_99] = _init99(new BigNumber(_i0_99));
        }
        _2690_v95 = _nw511;
        let _2692_v96;
        _2692_v96 = _dafny.Map.Empty.slice().updateUnsafe(_2690_v95,(_this).f9);
        let _2693_v97;
        _2693_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2688_v93,(_this).f9);
        let _2694_v98;
        _2694_v98 = _dafny.Seq.of(false);
        let _2695_v99;
        _2695_v99 = _dafny.MultiSet.fromElements((_this).f9, (_2694_v98)[_module.__default.safeIndex(_2688_v93, new BigNumber((_2694_v98).length))]);
        let _2696_v100;
        _2696_v100 = _dafny.MultiSet.fromElements((_2692_v96).equals(_2692_v96), (_this).f9, (_this).f9, ((_dafny.ZERO).minus(new BigNumber((_2693_v97).length))).isLessThanOrEqualTo((((_2695_v99).contains((_this).f9)) ? ((_2695_v99).get((_this).f9)) : (_2688_v93))));
        let _rhs510 = _2683_v88;
        let _rhs511 = (((_2689_v94).contains(_2685_v90)) ? ((_2689_v94).get(_2685_v90)) : (_2688_v93));
        let _rhs512 = new BigNumber((_2695_v99).cardinality());
        _2683_v88 = _rhs510;
        _2688_v93 = _rhs511;
        _2688_v93 = _rhs512;
        _2688_v93 = _2688_v93;
      } else {
        let _2697_v101;
        let _nw512 = new _module.C3();
        _nw512.__ctor((_this).f4, (_this).f5);
        _2697_v101 = _nw512;
        let _2698_v102;
        _2698_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(683));
        let _2699_v103;
        let _nw513 = Array((new BigNumber(19)).toNumber()).fill(false);
        _2699_v103 = _nw513;
        let _2700_v104;
        _2700_v104 = _module.D17.create_DC42(_2699_v103, new BigNumber(-284));
        let _2701_v105;
        _2701_v105 = _dafny.Seq.of((_2700_v104).dtor_cf65);
        let _2702_v106;
        _2702_v106 = _dafny.Set.fromElements((_2697_v101).fm5(new BigNumber((_2698_v102).length), (_this).f9, (_this).f8, new BigNumber((_2701_v105).length), globalState));
        _2702_v106 = _2702_v106;
        let _2703_v107;
        _2703_v107 = _module.D10.create_DC21(_dafny.Set.fromElements((_this).f9));
        let _2704_v108;
        _2704_v108 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21((_this).f8, new BigNumber(562), (_2703_v107).dtor_cf32, globalState),(_this).f8);
        let _2705_v109;
        _2705_v109 = _dafny.Seq.of((_this).f9);
        _2704_v108 = (_2704_v108).update(_dafny.Seq.update(_dafny.Seq.Concat(_2705_v109, _dafny.Seq.update(_2705_v109, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("chfoomfv")).length)), new BigNumber((_2705_v109).length)), (_this).f9)), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.Concat(_2705_v109, _dafny.Seq.update(_2705_v109, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("chfoomfv")).length)), new BigNumber((_2705_v109).length)), (_this).f9))).length)), (_this).f9), (_this).f8);
        let _2706_v110;
        _2706_v110 = _dafny.Set.fromElements(_dafny.Set.fromElements((_this).f9));
        let _2707_v111;
        _2707_v111 = _dafny.Seq.of(_module.__default.fm26((_this).f8, new BigNumber((_2681_v86).length), globalState), _dafny.Set.fromElements((_this).f9));
        _2706_v110 = function () {
          let _coll69 = new _dafny.Set();
          for (const _compr_69 of (_2707_v111).Elements) {
            let _2708_v112 = _compr_69;
            if (_dafny.Seq.contains(_2707_v111, _2708_v112)) {
              _coll69.add(_2708_v112);
            }
          }
          return _coll69;
        }();
        let _2709_v113;
        _2709_v113 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(((_this).f9) ? ((_dafny.ZERO).minus((_this).f8)) : ((_this).f8)));
        _2709_v113 = (_2709_v113).update((_this).f5, (_this).f8);
      }
      if ((_this).f9) {
        let _2710_v114;
        let _init100 = function (_2711_i11) {
          return _module.__default.safeModuloInt(_2711_i11, (_this).f8);
        };
        let _nw514 = Array((new BigNumber(21)).toNumber());
        for (let _i0_100 = 0; _i0_100 < new BigNumber(_nw514.length); _i0_100++) {
          _nw514[_i0_100] = _init100(new BigNumber(_i0_100));
        }
        _2710_v114 = _nw514;
        let _index515 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        (_2710_v114)[_index515] = (_this).f8;
        let _2712_v115;
        _2712_v115 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f8);
        let _index516 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        (_2710_v114)[_index516] = new BigNumber((_dafny.Seq.of(((_dafny.ZERO).minus((((_2712_v115).contains(true)) ? ((_2712_v115).get(true)) : ((_this).f8)))).isLessThan((_this).f8), ((_this).f8).isLessThan((_this).f8))).length);
        _2681_v86 = _2681_v86;
        let _2713_v116;
        _2713_v116 = _module.D16.create_DC38(_dafny.Seq.UnicodeFromString("mbqpxnm"), (_this).f9, (_2710_v114)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length))], (_this).f9);
        let _2714_v117;
        _2714_v117 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_module.D16.create_DC39(_2713_v116));
        let _2715_v118;
        _2715_v118 = _dafny.Seq.of(_2714_v117);
        let _2716_v119;
        _2716_v119 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f8, (_this).f8));
        let _2717_v120;
        _2717_v120 = _module.D26.create_DC61(_2716_v119);
        let _index517 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        (_2710_v114)[_index517] = ((_this).f8).plus((((_this).f9) ? (new BigNumber(((_2715_v118)[_module.__default.safeIndex(new BigNumber(-739), new BigNumber((_2715_v118).length))]).length)) : (new BigNumber(((_2717_v120).dtor_cf99).length))));
        let _2718_v121;
        _2718_v121 = _dafny.Seq.of((_this).f9);
        let _2719_v122;
        _2719_v122 = false;
        let _2720_v123;
        let _nw515 = Array((new BigNumber(3)).toNumber());
        _nw515[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw515[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw515[(new BigNumber(2)).toNumber()] = ((_2710_v114)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f8));
        _2720_v123 = _nw515;
        let _index518 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2720_v123).length));
        (_2720_v123)[_index518] = _2719_v122;
        let _2721_v124;
        _2721_v124 = _dafny.Set.fromElements((_this).f9, (_this).f9);
        let _index519 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        let _index520 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        let _index521 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2720_v123).length));
        let _rhs513 = _module.__default.fm21((_this).f8, (_this).f8, _2721_v124, globalState);
        let _rhs514 = (_this).f9;
        let _rhs515 = (_2710_v114)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length))];
        let _rhs516 = (_2710_v114)[_module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length))];
        let _rhs517 = _module.__default.fm1(_module.__default.safeDivisionInt((_this).f8, (_this).f8), new BigNumber(344), _dafny.Seq.IsProperPrefixOf(_2681_v86, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-310)), function (_2722_i12) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })), (_this).f9, globalState);
        let _lhs262 = _2710_v114;
        let _lhs263 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        let _lhs264 = _2710_v114;
        let _lhs265 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2710_v114).length));
        let _lhs266 = _2720_v123;
        let _lhs267 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2720_v123).length));
        _2718_v121 = _rhs513;
        _2719_v122 = _rhs514;
        _lhs262[_lhs263] = _rhs515;
        _lhs264[_lhs265] = _rhs516;
        _lhs266[_lhs267] = _rhs517;
        let _2723_v125;
        _2723_v125 = _dafny.Seq.of((_this).f8, (_dafny.ZERO).minus((_this).f8));
        let _2724_v126;
        _2724_v126 = _dafny.Map.Empty.slice().updateUnsafe(_2723_v125,(_this).f8);
        let _2725_v127;
        _2725_v127 = _2724_v126;
        let _2726_v128;
        _2726_v128 = _dafny.Map.Empty.slice().updateUnsafe(((!((_this).f9)) ? (_2725_v127) : (_2725_v127)),_2719_v122);
        let _2727_v129;
        _2727_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_2724_v126);
        let _2728_v130;
        _2728_v130 = _dafny.Map.Empty.slice().updateUnsafe((_2720_v123)[_module.__default.safeIndex(new BigNumber(980), new BigNumber((_2720_v123).length))],(_this).f5);
        let _2729_v131;
        _2729_v131 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(311));
        _2726_v128 = (_2726_v128).update((((_2727_v129).contains((_this).fm8(_2728_v130, _2729_v131, _2681_v86, globalState))) ? ((_2727_v129).get((_this).fm8(_2728_v130, _2729_v131, _2681_v86, globalState))) : (_2724_v126)), (_2720_v123)[_module.__default.safeIndex(new BigNumber(980), new BigNumber((_2720_v123).length))]);
      } else {
        let _2730_v132;
        _2730_v132 = new BigNumber(-444);
        let _2731_v133;
        _2731_v133 = _dafny.Seq.of((_this).f8);
        let _2732_v135;
        let _nw516 = new _module.C8();
        _nw516.__ctor();
        _2732_v135 = _nw516;
        let _2733_v136;
        _2733_v136 = _dafny.Map.Empty.slice().updateUnsafe(_2732_v135,(_this).f9);
        let _2734_v137;
        _2734_v137 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f8);
        let _2735_v138;
        _2735_v138 = _dafny.Set.fromElements(new BigNumber(-145), new BigNumber((_2734_v137).length));
        let _2736_v139;
        let _nw517 = Array((new BigNumber(7)).toNumber());
        _nw517[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(_module.__default.fm14(globalState))).isLessThan((_2731_v133)[_module.__default.safeIndex(_2730_v132, new BigNumber((_2731_v133).length))]);
        _nw517[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw517[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw517[(new BigNumber(3)).toNumber()] = !_dafny.Seq.contains(_module.__default.fm31(new BigNumber((function () {
          let _coll70 = new _dafny.Map();
          for (const _compr_70 of (_2681_v86).Elements) {
            let _2737_v134 = _compr_70;
            if (_dafny.Seq.contains(_2681_v86, _2737_v134)) {
              _coll70.push([_2737_v134,_2730_v132]);
            }
          }
          return _coll70;
        }()).length), new BigNumber((_2681_v86).length), globalState), (_this).f8);
        _nw517[(new BigNumber(4)).toNumber()] = (((_2733_v136).contains(_2732_v135)) ? ((_2733_v136).get(_2732_v135)) : ((_this).f9));
        _nw517[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw517[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements((_this).f8, new BigNumber((_2735_v138).length))).IsDisjointFrom(_dafny.Set.fromElements((_this).f8));
        _2736_v139 = _nw517;
        let _index522 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length));
        (_2736_v139)[_index522] = (_this).f9;
        let _2738_v140;
        let _nw518 = Array((_dafny.ONE).toNumber());
        _nw518[(_dafny.ZERO).toNumber()] = new BigNumber(351);
        _2738_v140 = _nw518;
        let _2739_v141;
        _2739_v141 = _dafny.MultiSet.fromElements((_this).f9);
        let _2740_v142;
        _2740_v142 = _module.D1.create_DC1(_2739_v141);
        let _index523 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2738_v140).length));
        (_2738_v140)[_index523] = new BigNumber(((_dafny.MultiSet.fromElements((_this).f9)).Intersect((_2740_v142).dtor_cf1)).cardinality());
        let _2741_v143;
        _2741_v143 = _dafny.Set.fromElements((_this).f9);
        let _2742_v144;
        _2742_v144 = _dafny.MultiSet.fromElements((_this).f8, (_this).f8, new BigNumber(-770), _2730_v132, new BigNumber((_2741_v143).length));
        let _index524 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length));
        let _index525 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2738_v140).length));
        let _rhs518 = new BigNumber(35);
        let _rhs519 = (_this).f8;
        let _rhs520 = (_2742_v144).IsProperSubsetOf((_2742_v144).Intersect(_2742_v144));
        let _rhs521 = new BigNumber((_2731_v133).length);
        let _rhs522 = _module.__default.safeDivisionInt((new BigNumber((_dafny.Seq.UnicodeFromString("uxesrvvq")).length)).plus(_2730_v132), _2730_v132);
        let _lhs268 = _2736_v139;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length));
        let _lhs270 = _2738_v140;
        let _lhs271 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2738_v140).length));
        _2730_v132 = _rhs518;
        _2730_v132 = _rhs519;
        _lhs268[_lhs269] = _rhs520;
        _lhs270[_lhs271] = _rhs521;
        _2730_v132 = _rhs522;
        _2681_v86 = _2681_v86;
        let _index526 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2738_v140).length));
        (_2738_v140)[_index526] = _module.__default.safeDivisionInt(_2730_v132, (_this).f8);
        let _2743_v145;
        _2743_v145 = _dafny.Map.Empty.slice().updateUnsafe((_2736_v139)[_module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length))],(_this).f5);
        let _2744_v146;
        _2744_v146 = _module.D1.create_DC2(new BigNumber((_module.__default.fm21(_2730_v132, _2730_v132, _2741_v143, globalState)).length), _2730_v132);
        let _2745_v147;
        _2745_v147 = _dafny.Map.Empty.slice().updateUnsafe(_2739_v141,(_2738_v140)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_2738_v140).length))]);
        let _index527 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length));
        (_2736_v139)[_index527] = ((_this).fm8(_2743_v145, (_this).fm9(_2744_v146, _2681_v86, globalState), _2681_v86, globalState)).isEqualTo(new BigNumber((_2745_v147).length));
        let _2746_v148;
        _2746_v148 = _module.D6.create_DC15(true, true, false);
        let _2747_v149;
        _2747_v149 = _dafny.Map.Empty.slice().updateUnsafe((_2746_v148).dtor_cf20,!((_this).f9));
        let _2748_v150;
        _2748_v150 = _2747_v149;
        let _2749_v151;
        _2749_v151 = _dafny.Map.Empty.slice().updateUnsafe(_2681_v86,_2747_v149);
        let _2750_v152;
        _2750_v152 = _dafny.Seq.of(_2738_v140, _2738_v140);
        let _2751_v153;
        _2751_v153 = _dafny.Map.Empty.slice().updateUnsafe(_2750_v152,_2749_v151);
        let _index528 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_2736_v139).length));
        (_2736_v139)[_index528] = ((_2749_v151).Merge((((_2751_v153).contains(_2750_v152)) ? ((_2751_v153).get(_2750_v152)) : ((_2749_v151).update(_2681_v86, _2748_v150))))).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("skiyru"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-117)), function (_2752_i13) {
          return (_this).f5;
        })));
      }
      let _hi20 = new BigNumber(-926);
      for (let _2753_i14 = new BigNumber((_dafny.Seq.Concat(_2681_v86, _dafny.Seq.of((_this).f5, (_this).f5))).length); _2753_i14.isLessThan(_hi20); _2753_i14 = _2753_i14.plus(_dafny.ONE)) {
        let _2754_v154;
        let _init101 = ((_2755_i14) => function (_2756_i15) {
          return (_2756_i15).multipliedBy(_2755_i14);
        })(_2753_i14);
        let _nw519 = Array((new BigNumber(2)).toNumber());
        for (let _i0_101 = 0; _i0_101 < new BigNumber(_nw519.length); _i0_101++) {
          _nw519[_i0_101] = _init101(new BigNumber(_i0_101));
        }
        _2754_v154 = _nw519;
        let _index529 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
        (_2754_v154)[_index529] = (_this).f8;
        let _2757_v155;
        let _nw520 = Array((new BigNumber(2)).toNumber());
        _nw520[(_dafny.ZERO).toNumber()] = _2754_v154;
        _nw520[(_dafny.ONE).toNumber()] = _2754_v154;
        _2757_v155 = _nw520;
        let _2758_v156;
        _2758_v156 = _dafny.Seq.of((_this).f9);
        let _2759_v157;
        _2759_v157 = _dafny.Seq.of((_this).f9, true, (_this).f9, !_dafny.areEqual(_2758_v156, _2758_v156));
        let _2760_v158;
        _2760_v158 = false;
        let _2761_v159;
        _2761_v159 = _dafny.Map.Empty.slice().updateUnsafe(_2760_v158,(_this).f8);
        let _2762_v160;
        _2762_v160 = _dafny.Seq.of(new BigNumber((_2761_v159).length), (_2753_i14).minus(_module.__default.fm14(globalState)));
        let _2763_v161;
        _2763_v161 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber(335));
        let _index530 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
        let _rhs523 = _2762_v160;
        let _rhs524 = _module.__default.fm13(!(_2760_v158), _2760_v158, ((_this).f8).isLessThanOrEqualTo(new BigNumber((_2763_v161).length)), globalState);
        let _rhs525 = _2757_v155;
        let _rhs526 = _dafny.Seq.of(_2760_v158);
        let _rhs527 = _module.__default.fm1((_dafny.ZERO).minus((_this).f8), _2753_i14, (_this).f9, ((_2760_v158) ? (_2760_v158) : (true)), globalState);
        let _lhs272 = globalState;
        let _lhs273 = _2754_v154;
        let _lhs274 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
        _lhs272.f2 = _rhs523;
        _lhs273[_lhs274] = _rhs524;
        _2757_v155 = _rhs525;
        _2759_v157 = _rhs526;
        _2760_v158 = _rhs527;
        if ((_this).f9) {
          let _2764_v162;
          _2764_v162 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_2753_i14));
          let _index531 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          let _index532 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          let _index533 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          let _rhs528 = (_this).fm4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(270)), function (_2765_i16) {
            return (_this).f8;
          }), globalState);
          let _rhs529 = _2753_i14;
          let _rhs530 = _2754_v154;
          let _rhs531 = new BigNumber((_2764_v162).cardinality());
          let _rhs532 = _2681_v86;
          let _lhs275 = _2754_v154;
          let _lhs276 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          let _lhs277 = _2754_v154;
          let _lhs278 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          let _lhs279 = _2754_v154;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          _lhs275[_lhs276] = _rhs528;
          _lhs277[_lhs278] = _rhs529;
          _2754_v154 = _rhs530;
          _lhs279[_lhs280] = _rhs531;
          _2681_v86 = _rhs532;
          let _2766_v163;
          let _nw521 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2766_v163 = _nw521;
          let _index534 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2766_v163).length));
          (_2766_v163)[_index534] = new _dafny.CodePoint('l'.codePointAt(0));
          let _index535 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2766_v163).length));
          (_2766_v163)[_index535] = (_this).f5;
          let _2767_v164;
          _2767_v164 = _module.D23.create_DC55(_2681_v86, (_this).f9, (_this).f9, (_this).f9, _2759_v157);
          let _2768_v165;
          _2768_v165 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber(((_2767_v164).dtor_cf85).length));
          let _2769_v166;
          _2769_v166 = _dafny.Seq.of(_2768_v165);
          let _2770_v167;
          _2770_v167 = _dafny.Set.fromElements(new BigNumber(465));
          let _2771_v168;
          _2771_v168 = _module.D1.create_DC2((_dafny.ZERO).minus(_2753_i14), new BigNumber((_2770_v167).length));
          let _2772_v169;
          _2772_v169 = _dafny.Map.Empty.slice().updateUnsafe(_2760_v158,(_2766_v163)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_2766_v163).length))]);
          _2769_v166 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2753_i14,new BigNumber((_2681_v86).length)), _2768_v165, _2768_v165), _2769_v166), _2769_v166), _module.__default.safeIndex((_this).f8, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2753_i14,new BigNumber((_2681_v86).length)), _2768_v165, _2768_v165), _2769_v166), _2769_v166)).length)), ((_this).fm9(_2771_v168, _dafny.Seq.UnicodeFromString("wpjskhwm"), globalState)).update(_2753_i14, (_this).fm8(_2772_v169, _2768_v165, _2681_v86, globalState)));
          let _2773_v170;
          _2773_v170 = _dafny.Map.Empty.slice().updateUnsafe(_2753_i14,_2754_v154);
          let _2774_v171;
          _2774_v171 = _dafny.Map.Empty.slice().updateUnsafe(_2760_v158,_2773_v170);
          let _2775_v172;
          let _nw522 = Array((new BigNumber(17)).toNumber());
          _nw522[(_dafny.ZERO).toNumber()] = _2773_v170;
          _nw522[(_dafny.ONE).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(2)).toNumber()] = (_2773_v170).Merge(_2773_v170);
          _nw522[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))],_2754_v154);
          _nw522[(new BigNumber(4)).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(5)).toNumber()] = (_2773_v170).Merge(_2773_v170);
          _nw522[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))],_2754_v154)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9)).length),_2754_v154));
          _nw522[(new BigNumber(7)).toNumber()] = (_2773_v170).Merge(_2773_v170);
          _nw522[(new BigNumber(8)).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(9)).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(10)).toNumber()] = (((_2774_v171).contains(_2760_v158)) ? ((_2774_v171).get(_2760_v158)) : (_2773_v170));
          _nw522[(new BigNumber(11)).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(12)).toNumber()] = _2773_v170;
          _nw522[(new BigNumber(13)).toNumber()] = (((_this).f9) ? (_2773_v170) : (_2773_v170));
          _nw522[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))],_2754_v154)).update((_this).f8, _2754_v154);
          _nw522[(new BigNumber(15)).toNumber()] = (_2773_v170).Merge(_2773_v170);
          _nw522[(new BigNumber(16)).toNumber()] = _2773_v170;
          _2775_v172 = _nw522;
          let _index536 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_2775_v172).length));
          (_2775_v172)[_index536] = _2773_v170;
          let _2776_v173;
          _2776_v173 = _dafny.Seq.of(_2773_v170);
          let _index537 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_2775_v172).length));
          (_2775_v172)[_index537] = (((_2776_v173)[_module.__default.safeIndex(_2753_i14, new BigNumber((_2776_v173).length))]).Merge(_2773_v170)).Merge(_2773_v170);
          let _2777_v174;
          _2777_v174 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,!((_this).f9));
          _2777_v174 = ((_2777_v174).Merge(_2777_v174)).Merge(_2777_v174);
        } else {
          _2760_v158 = !((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))]).isEqualTo(_2753_i14);
          let _2778_v175;
          let _nw523 = new _module.C4();
          _nw523.__ctor(((_this).f8).multipliedBy((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))]), ((_this).f8).minus((_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))]), (_this).f4, (_this).f5);
          _2778_v175 = _nw523;
          let _2779_v176;
          let _nw524 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _2779_v176 = _nw524;
          let _index538 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_2779_v176).length));
          (_2779_v176)[_index538] = function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of _dafny.IntegerRange(new BigNumber(170), new BigNumber(-457))) {
              let _2780_v177 = _compr_71;
              if (((new BigNumber(170)).isLessThanOrEqualTo(_2780_v177)) && ((_2780_v177).isLessThan(new BigNumber(-457)))) {
                _coll71.push([_module.__default.safeModuloInt(_2780_v177, new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())),_2760_v158]);
              }
            }
            return _coll71;
          }();
          let _2781_v178;
          _2781_v178 = _dafny.Map.Empty.slice().updateUnsafe((_2753_i14).multipliedBy((_this).f8),_2760_v158);
          let _index539 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_2779_v176).length));
          (_2779_v176)[_index539] = _2781_v178;
          let _2782_v179;
          _2782_v179 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f5);
          let _2783_v180;
          _2783_v180 = _module.D5.create_DC11((_2782_v179).update(_2760_v158, new _dafny.CodePoint('y'.codePointAt(0))));
          let _2784_v181;
          _2784_v181 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2681_v86).length),(_2778_v175).f12);
          let _rhs533 = _2783_v180;
          let _rhs534 = (_this).f9;
          let _rhs535 = (_this).fm8(_module.__default.fm48(_2760_v158, false, _module.__default.fm1(_2778_v175.f11, _2753_i14, (_this).f9, _2760_v158, globalState), (_this).f8, globalState), _2784_v181, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cwjmmom"), _2681_v86), globalState);
          let _lhs281 = _2778_v175;
          _2783_v180 = _rhs533;
          _2760_v158 = _rhs534;
          _lhs281.f11 = _rhs535;
          let _index540 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
          (_2754_v154)[_index540] = (_2778_v175).f12;
        }
        let _index541 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length));
        (_2754_v154)[_index541] = (_2754_v154)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2754_v154).length))];
        let _2785_v182;
        let _init102 = function (_2786_i17) {
          return (_this).f5;
        };
        let _nw525 = Array((new BigNumber(13)).toNumber());
        for (let _i0_102 = 0; _i0_102 < new BigNumber(_nw525.length); _i0_102++) {
          _nw525[_i0_102] = _init102(new BigNumber(_i0_102));
        }
        _2785_v182 = _nw525;
        let _2787_v183;
        _2787_v183 = _dafny.Seq.of(_2785_v182);
        let _rhs536 = _2760_v158;
        let _rhs537 = (_this).f9;
        let _rhs538 = _2787_v183;
        _2760_v158 = _rhs536;
        _2760_v158 = _rhs537;
        _2787_v183 = _rhs538;
      }
      let _2788_v184;
      let _nw526 = new _module.C2();
      _nw526.__ctor((_this).f4, (_2681_v86)[_module.__default.safeIndex((_this).f8, new BigNumber((_2681_v86).length))]);
      _2788_v184 = _nw526;
      if (false) {
        let _2789_v185;
        _2789_v185 = new BigNumber(92);
        let _2790_v186;
        let _nw527 = new _module.C5();
        _nw527.__ctor();
        _2790_v186 = _nw527;
        let _2791_v187;
        _2791_v187 = _module.D21.create_DC51((_this).f9, _2790_v186);
        let _2792_v188;
        _2792_v188 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _2793_v189;
        _2793_v189 = _dafny.Seq.of((((_2792_v188).contains((_this).f9)) ? ((_2792_v188).get((_this).f9)) : ((_this).f9)));
        let _2794_v190;
        _2794_v190 = _module.D23.create_DC55(_2681_v86, (_this).f9, (((_this).f9) ? ((_2791_v187).dtor_cf80) : ((_this).f9)), (_2788_v184).fm5((_this).f8, (_this).f9, new BigNumber(262), new BigNumber((_2681_v86).length), globalState), _2793_v189);
        let _2795_v191;
        _2795_v191 = _dafny.Set.fromElements(_2792_v188);
        let _2796_v192;
        _2796_v192 = _dafny.Seq.of((_this).f8, _2789_v185);
        let _rhs539 = _module.__default.safeDivisionInt((_this).f8, new BigNumber(((((_this).f9) ? (_2795_v191) : (_2795_v191))).length));
        let _rhs540 = _2794_v190;
        let _rhs541 = ((_2796_v192)[_module.__default.safeIndex(_2789_v185, new BigNumber((_2796_v192).length))]).plus((_this).f8);
        let _rhs542 = (_dafny.ZERO).minus(_2789_v185);
        _2789_v185 = _rhs539;
        _2794_v190 = _rhs540;
        _2789_v185 = _rhs541;
        _2789_v185 = _rhs542;
        let _2797_v193;
        let _nw528 = Array((new BigNumber(28)).toNumber()).fill(false);
        _2797_v193 = _nw528;
        let _index542 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length));
        (_2797_v193)[_index542] = (_this).f9;
        let _2798_v194;
        _2798_v194 = false;
        let _index543 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_2797_v193).length));
        (_2797_v193)[_index543] = _dafny.Seq.IsPrefixOf(_2793_v189, _2793_v189);
        let _2799_v195;
        _2799_v195 = _dafny.MultiSet.fromElements(_2789_v185);
        let _2800_v196;
        _2800_v196 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f5);
        let _2801_v197;
        _2801_v197 = _dafny.Map.Empty.slice().updateUnsafe(_2793_v189,_2798_v194);
        let _2802_v198;
        _2802_v198 = _dafny.Map.Empty.slice().updateUnsafe(_2789_v185,new BigNumber((_2801_v197).length));
        let _2803_v199;
        _2803_v199 = _dafny.Set.fromElements(_2798_v194, (_this).f9);
        let _2804_v200;
        let _nw529 = Array((new BigNumber(28)).toNumber());
        _nw529[(_dafny.ZERO).toNumber()] = _2789_v185;
        _nw529[(_dafny.ONE).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(2)).toNumber()] = new BigNumber(601);
        _nw529[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2681_v86).length));
        _nw529[(new BigNumber(4)).toNumber()] = _2789_v185;
        _nw529[(new BigNumber(5)).toNumber()] = (((_2799_v195).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2789_v185,(_this).f9)).length))) ? ((_2799_v195).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2789_v185,(_this).f9)).length))) : (_2789_v185));
        _nw529[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw529[(new BigNumber(7)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(8)).toNumber()] = (_this).fm8(_2800_v196, _2802_v198, _2681_v86, globalState);
        _nw529[(new BigNumber(9)).toNumber()] = _2789_v185;
        _nw529[(new BigNumber(10)).toNumber()] = new BigNumber(795);
        _nw529[(new BigNumber(11)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(12)).toNumber()] = _2789_v185;
        _nw529[(new BigNumber(13)).toNumber()] = new BigNumber((_2803_v199).length);
        _nw529[(new BigNumber(14)).toNumber()] = _2789_v185;
        _nw529[(new BigNumber(15)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(16)).toNumber()] = new BigNumber((_2796_v192).length);
        _nw529[(new BigNumber(17)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2798_v194,(_2796_v192)[_module.__default.safeIndex(new BigNumber((_2792_v188).length), new BigNumber((_2796_v192).length))])).length);
        _nw529[(new BigNumber(19)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(20)).toNumber()] = new BigNumber(69);
        _nw529[(new BigNumber(21)).toNumber()] = new BigNumber(-586);
        _nw529[(new BigNumber(22)).toNumber()] = new BigNumber(941);
        _nw529[(new BigNumber(23)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(24)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(25)).toNumber()] = (_this).f8;
        _nw529[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus((_this).f8);
        _nw529[(new BigNumber(27)).toNumber()] = new BigNumber((_2796_v192).length);
        _2804_v200 = _nw529;
        let _2805_v201;
        _2805_v201 = _module.D11.create_DC25((_this).f9, _2789_v185, _2804_v200, _2789_v185);
        let _2806_v202;
        _2806_v202 = _dafny.Map.Empty.slice().updateUnsafe(_2793_v189,(_this).f8);
        let _index544 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length));
        let _index545 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_2797_v193).length));
        let _rhs543 = new BigNumber((_dafny.Seq.Concat(_2681_v86, _2681_v86)).length);
        let _rhs544 = !((_this).f9) || (_2798_v194);
        let _rhs545 = (_2805_v201).dtor_cf40;
        let _rhs546 = !(!((_2789_v185).isLessThanOrEqualTo((((_2806_v202).contains(_dafny.Seq.of(_2798_v194, _2798_v194))) ? ((_2806_v202).get(_dafny.Seq.of(_2798_v194, _2798_v194))) : (new BigNumber(-173)))))) || ((_2789_v185).isLessThanOrEqualTo(_2789_v185));
        let _rhs547 = !((_2794_v190).dtor_cf87);
        let _lhs282 = _2797_v193;
        let _lhs283 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length));
        let _lhs284 = _2797_v193;
        let _lhs285 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_2797_v193).length));
        _2789_v185 = _rhs543;
        _lhs282[_lhs283] = _rhs544;
        r0 = _rhs545;
        _2798_v194 = _rhs546;
        _lhs284[_lhs285] = _rhs547;
        let _2807_v203;
        _2807_v203 = _module.D15.create_DC36(_module.D15.create_DC35(_2789_v185));
        let _2808_v204;
        _2808_v204 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f9) || ((_2797_v193)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length))]),_2807_v203);
        _2808_v204 = (_2808_v204).update(_2798_v194, _2807_v203);
        let _2809_v205;
        _2809_v205 = _dafny.MultiSet.fromElements((_this).f5);
        let _2810_v206;
        _2810_v206 = _dafny.Map.Empty.slice().updateUnsafe(_2802_v198,(_this).f9);
        let _index546 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length));
        let _rhs548 = (_2797_v193)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length))];
        let _rhs549 = (_dafny.ZERO).minus((_this).f8);
        let _rhs550 = _module.__default.fm13((_2809_v205).IsSubsetOf(_2809_v205), (((_2810_v206).contains(_2802_v198)) ? ((_2810_v206).get(_2802_v198)) : ((_2797_v193)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length))])), (_2797_v193)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length))], globalState);
        let _lhs286 = _2797_v193;
        let _lhs287 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2797_v193).length));
        _lhs286[_lhs287] = _rhs548;
        _2789_v185 = _rhs549;
        _2789_v185 = _rhs550;
        let _2811_v207;
        _2811_v207 = _dafny.Seq.of(_2804_v200);
        let _2812_v208;
        let _nw530 = Array((new BigNumber(8)).toNumber());
        _nw530[(_dafny.ZERO).toNumber()] = (_this).f5;
        _nw530[(_dafny.ONE).toNumber()] = (_this).f5;
        _nw530[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw530[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw530[(new BigNumber(4)).toNumber()] = (_this).f5;
        _nw530[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
        _nw530[(new BigNumber(6)).toNumber()] = (_this).f5;
        _nw530[(new BigNumber(7)).toNumber()] = (_this).f5;
        _2812_v208 = _nw530;
        let _2813_v209;
        let _nw531 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _2813_v209 = _nw531;
        let _2814_v210;
        _2814_v210 = _dafny.Map.Empty.slice().updateUnsafe(_2799_v195,new BigNumber((_dafny.Seq.UnicodeFromString("vvi")).length));
        let _index547 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_2813_v209).length));
        (_2813_v209)[_index547] = (_2814_v210).Merge(_2814_v210);
        let _2815_v211;
        _2815_v211 = _dafny.Set.fromElements(_2804_v200, _2804_v200, _2804_v200);
        let _index548 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_2813_v209).length));
        let _rhs551 = _2811_v207;
        let _rhs552 = _2812_v208;
        let _rhs553 = (_2815_v211).IsProperSubsetOf(_dafny.Set.fromElements(_2804_v200, _2804_v200));
        let _rhs554 = _2814_v210;
        let _lhs288 = _2813_v209;
        let _lhs289 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_2813_v209).length));
        _2811_v207 = _rhs551;
        _2812_v208 = _rhs552;
        _2798_v194 = _rhs553;
        _lhs288[_lhs289] = _rhs554;
      } else {
        let _2816_v212;
        let _init103 = function (_2817_i18) {
          return (_2817_i18).minus((_this).f8);
        };
        let _nw532 = Array((new BigNumber(13)).toNumber());
        for (let _i0_103 = 0; _i0_103 < new BigNumber(_nw532.length); _i0_103++) {
          _nw532[_i0_103] = _init103(new BigNumber(_i0_103));
        }
        _2816_v212 = _nw532;
        let _index549 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2816_v212).length));
        (_2816_v212)[_index549] = (_this).f8;
        let _index550 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2816_v212).length));
        (_2816_v212)[_index550] = (_this).f8;
        let _2818_v213;
        _2818_v213 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2681_v86);
        let _2819_v214;
        let _nw533 = Array((new BigNumber(6)).toNumber());
        _nw533[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("nxlj");
        _nw533[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2681_v86, _2681_v86);
        _nw533[(new BigNumber(2)).toNumber()] = _2681_v86;
        _nw533[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2681_v86, _2681_v86);
        _nw533[(new BigNumber(4)).toNumber()] = _2681_v86;
        _nw533[(new BigNumber(5)).toNumber()] = (((_2818_v213).contains((_this).f5)) ? ((_2818_v213).get((_this).f5)) : (_2681_v86));
        _2819_v214 = _nw533;
        let _index551 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2819_v214).length));
        (_2819_v214)[_index551] = _2681_v86;
        let _index552 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2819_v214).length));
        (_2819_v214)[_index552] = _dafny.Seq.Concat(_2681_v86, _2681_v86);
        let _2820_v215;
        let _nw534 = new _module.C7();
        _nw534.__ctor();
        _2820_v215 = _nw534;
        let _2821_v216;
        _2821_v216 = _dafny.Map.Empty.slice().updateUnsafe((_2816_v212)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_2816_v212).length))],(_this).f9);
        let _2822_v217;
        _2822_v217 = _dafny.Seq.of((_this).f9, (_this).f9);
        _2821_v216 = (_2821_v216).update((_this).f8, !_dafny.areEqual(_2822_v217, _2822_v217));
        let _2823_v218;
        _2823_v218 = true;
        _2823_v218 = !(_dafny.Seq.IsProperPrefixOf(_2681_v86, _dafny.Seq.UnicodeFromString("hufpuyy")));
      }
      let _2824_v219;
      let _nw535 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _2824_v219 = _nw535;
      r0 = _2824_v219;
      return r0;
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f4 = [];
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f7 = _dafny.ZERO;
      this._f6 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f6, f7, f4, f5) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return ((_this).f7).multipliedBy((_this).f7);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(((false) ? (true) : ((new BigNumber((_dafny.Seq.UnicodeFromString("sfwsjafps")).length)).isLessThan((_this).f7))));
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      if (false) {
        return !(_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll72 = new _dafny.Map();
          for (const _compr_72 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f5)).Keys.Elements) {
            let _2825_v0 = _compr_72;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f5)).contains(_2825_v0)) {
              _coll72.push([(_2825_v0).minus(new BigNumber(-860)),_module.D2.create_DC4(true)]);
            }
          }
          return _coll72;
        }(),true)).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_module.D2.create_DC4(false)),true));
      } else {
        return true;
      }
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f7)).cardinality()),_dafny.Seq.of(!((false))))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f7),_dafny.Seq.of(true, false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_dafny.Seq.of(true, false))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_dafny.Seq.of(false, false, false))))).length);
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f7;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      if (false) {
        let _2826_v0;
        _2826_v0 = _dafny.Seq.UnicodeFromString("nnsxyet");
        let _2827_v1;
        let _nw536 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2827_v1 = _nw536;
        let _2828_v2;
        _2828_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2827_v1,_dafny.Seq.UnicodeFromString("gspoppq"));
        _2826_v0 = _dafny.Seq.update((((_2828_v2).contains(_2827_v1)) ? ((_2828_v2).get(_2827_v1)) : (_2826_v0)), _module.__default.safeIndex(new BigNumber((_2826_v0).length), new BigNumber(((((_2828_v2).contains(_2827_v1)) ? ((_2828_v2).get(_2827_v1)) : (_2826_v0))).length)), (_this).f5);
        let _2829_v3;
        let _nw537 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _2829_v3 = _nw537;
        let _2830_v4;
        _2830_v4 = _module.D4.create_DC9(_2829_v3);
        _2829_v3 = (_2830_v4).dtor_cf10;
        let _2831_v5;
        _2831_v5 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f7));
        let _2832_v6;
        _2832_v6 = _module.D1.create_DC2((_this).f7, (((_2831_v5).contains((_this).f7)) ? ((_2831_v5).get((_this).f7)) : ((_this).f7)));
        let _2833_v7;
        _2833_v7 = true;
        let _2834_v8;
        _2834_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2830_v4,(_this).f7);
        r1 = (_this).fm7(function (_pat_let59_0) {
          return function (_2835_dt__update__tmp_h0) {
            return function (_pat_let60_0) {
              return function (_2836_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_2836_dt__update_hcf2_h0, (_2835_dt__update__tmp_h0).dtor_cf3);
              }(_pat_let60_0);
            }((_this).f7);
          }(_pat_let59_0);
        }(_2832_v6), _2833_v7, (new BigNumber(779)).plus((_this).f7), !(_module.__default.fm1((_this).f7, new BigNumber(((_2834_v8).update(_2830_v4, (_this).f7)).length), _2833_v7, _2833_v7, globalState)), globalState);
        r1 = (_this).f7;
        if ((_2833_v7) || (false)) {
          _2833_v7 = _2833_v7;
          let _2837_v9;
          let _init104 = ((_2838_v0) => function (_2839_i0) {
            return _2838_v0;
          })(_2826_v0);
          let _nw538 = Array((new BigNumber(14)).toNumber());
          for (let _i0_104 = 0; _i0_104 < new BigNumber(_nw538.length); _i0_104++) {
            _nw538[_i0_104] = _init104(new BigNumber(_i0_104));
          }
          _2837_v9 = _nw538;
          _2837_v9 = ((true) ? (_2837_v9) : (_2837_v9));
          let _2840_v10;
          let _nw539 = new _module.C1();
          _nw539.__ctor((_this).f4, (_this).f5);
          _2840_v10 = _nw539;
          let _2841_v11;
          _2841_v11 = _dafny.Seq.of(_2840_v10, _2840_v10);
          let _2842_v12;
          _2842_v12 = _dafny.Seq.of(_2833_v7, (_this).fm3((_dafny.ZERO).minus(new BigNumber((_2841_v11).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), ((_2843_v0) => function (_2844_i1) {
            return (_2843_v0)[_module.__default.safeIndex((_this).f7, new BigNumber((_2843_v0).length))];
          })(_2826_v0))).length), globalState));
          _2842_v12 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2842_v12, _dafny.Seq.of(false, _2833_v7)), _2842_v12);
          _2833_v7 = (_2842_v12)[_module.__default.safeIndex(new BigNumber((_2842_v12).length), new BigNumber((_2842_v12).length))];
          let _2845_v13;
          _2845_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2833_v7,_module.__default.safeDivisionInt(new BigNumber(964), new BigNumber(292)));
          let _2846_v14;
          _2846_v14 = _dafny.Set.fromElements((_this).f7);
          let _2847_v15;
          _2847_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2826_v0,_2833_v7);
          let _2848_v16;
          _2848_v16 = _dafny.Set.fromElements(_2833_v7);
          let _2849_v17;
          _2849_v17 = _dafny.Seq.of((_this).f7, new BigNumber((_module.__default.fm21(new BigNumber((_2847_v15).length), (_this).f7, _2848_v16, globalState)).length));
          _2845_v13 = (_2845_v13).update((_2833_v7) && ((_this).fm3(new BigNumber((_2846_v14).length), (_this).f7, globalState)), (_this).fm4(_2849_v17, globalState));
        } else {
          let _index553 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2827_v1).length));
          (_2827_v1)[_index553] = (_this).f5;
          let _2850_v18;
          let _nw540 = new _module.C9();
          _nw540.__ctor((_this).f4, (_this).f5);
          _2850_v18 = _nw540;
          let _index554 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2827_v1).length));
          let _rhs555 = ((_2833_v7) ? (((_2833_v7) ? ((_this).f5) : ((_this).f5))) : ((_this).f5));
          let _rhs556 = _2850_v18;
          let _lhs290 = _2827_v1;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2827_v1).length));
          _lhs290[_lhs291] = _rhs555;
          _2850_v18 = _rhs556;
          let _2851_v19;
          _2851_v19 = _dafny.Seq.of((_this).f7);
          let _2852_v20;
          let _nw541 = Array((new BigNumber(6)).toNumber());
          _nw541[(_dafny.ZERO).toNumber()] = (_this).f7;
          _nw541[(_dafny.ONE).toNumber()] = (_this).f7;
          _nw541[(new BigNumber(2)).toNumber()] = (_this).f7;
          _nw541[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw541[(new BigNumber(4)).toNumber()] = (_this).f7;
          _nw541[(new BigNumber(5)).toNumber()] = new BigNumber((_2851_v19).length);
          _2852_v20 = _nw541;
          let _2853_v21;
          _2853_v21 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_2852_v20);
          _2853_v21 = (_2853_v21).update((_this).f5, _2852_v20);
          let _index555 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2827_v1).length));
          (_2827_v1)[_index555] = (_this).f5;
          let _2854_v22;
          _2854_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,false);
          let _2855_v23;
          _2855_v23 = (((_2854_v22).contains(new BigNumber(264))) ? ((_2854_v22).get(new BigNumber(264))) : (_2833_v7));
          _2855_v23 = _2855_v23;
          let _2856_v24;
          let _nw542 = new _module.C1();
          _nw542.__ctor((_this).f4, (_this).f5);
          _2856_v24 = _nw542;
        }
      } else {
        let _2857_v25;
        let _init105 = function (_2858_i2) {
          return (_2858_i2).multipliedBy((_module.D14.create_DC31((_this).f7)).dtor_cf48);
        };
        let _nw543 = Array((new BigNumber(18)).toNumber());
        for (let _i0_105 = 0; _i0_105 < new BigNumber(_nw543.length); _i0_105++) {
          _nw543[_i0_105] = _init105(new BigNumber(_i0_105));
        }
        _2857_v25 = _nw543;
        let _index556 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2857_v25).length));
        (_2857_v25)[_index556] = (_this).f7;
        let _2859_v26;
        _2859_v26 = true;
        let _2860_v27;
        _2860_v27 = _module.D13.create_DC29(_2859_v26, _2859_v26);
        let _2861_v28;
        _2861_v28 = _dafny.MultiSet.fromElements((_2860_v27).dtor_cf46);
        let _index557 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_2857_v25).length));
        (_2857_v25)[_index557] = new BigNumber((_2861_v28).cardinality());
        let _2862_v29;
        let _nw544 = new _module.C2();
        _nw544.__ctor((_this).f4, (_this).f5);
        _2862_v29 = _nw544;
        r1 = (_2857_v25)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_2857_v25).length))];
        let _2863_v30;
        let _init106 = function (_2864_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,true);
        };
        let _nw545 = Array((new BigNumber(16)).toNumber());
        for (let _i0_106 = 0; _i0_106 < new BigNumber(_nw545.length); _i0_106++) {
          _nw545[_i0_106] = _init106(new BigNumber(_i0_106));
        }
        _2863_v30 = _nw545;
        _2863_v30 = _2863_v30;
        let _2865_v31;
        _2865_v31 = _dafny.Set.fromElements(_2859_v26, _2859_v26);
        let _2866_v32;
        _2866_v32 = _dafny.Seq.of(((_this).f7).isEqualTo(new BigNumber((_2865_v31).length)));
        let _2867_v33;
        let _nw546 = new _module.C5();
        _nw546.__ctor();
        _2867_v33 = _nw546;
        let _2868_v34;
        let _nw547 = Array((new BigNumber(7)).toNumber());
        _nw547[(_dafny.ZERO).toNumber()] = _2867_v33;
        _nw547[(_dafny.ONE).toNumber()] = _2867_v33;
        _nw547[(new BigNumber(2)).toNumber()] = _2867_v33;
        _nw547[(new BigNumber(3)).toNumber()] = _2867_v33;
        _nw547[(new BigNumber(4)).toNumber()] = _2867_v33;
        _nw547[(new BigNumber(5)).toNumber()] = _2867_v33;
        _nw547[(new BigNumber(6)).toNumber()] = _2867_v33;
        _2868_v34 = _nw547;
        let _2869_v35;
        _2869_v35 = _module.D21.create_DC51(_2859_v26, _2867_v33);
        let _2870_v36;
        _2870_v36 = _dafny.Seq.of(_2867_v33, _2867_v33, (_2869_v35).dtor_cf81);
        let _index558 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_2868_v34).length));
        (_2868_v34)[_index558] = (_2870_v36)[_module.__default.safeIndex(_module.__default.fm14(globalState), new BigNumber((_2870_v36).length))];
        let _2871_v37;
        _2871_v37 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_this).fm5((_this).f7, !(_2859_v26), (_this).f7, (_this).f7, globalState));
        let _index559 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_2868_v34).length));
        let _rhs557 = _dafny.Seq.update(_dafny.Seq.update(_2866_v32, _module.__default.safeIndex((_this).f7, new BigNumber((_2866_v32).length)), _2859_v26), _module.__default.safeIndex((_2857_v25)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_2857_v25).length))], new BigNumber((_dafny.Seq.update(_2866_v32, _module.__default.safeIndex((_this).f7, new BigNumber((_2866_v32).length)), _2859_v26)).length)), _2859_v26);
        let _rhs558 = ((_2859_v26) ? (_2859_v26) : ((_2862_v29).fm5((_2857_v25)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_2857_v25).length))], (((_2871_v37).contains((_this).f5)) ? ((_2871_v37).get((_this).f5)) : (_2859_v26)), new BigNumber((_module.__default.fm2((_this).f7, _2859_v26, _2859_v26, new BigNumber(457), globalState)).length), (_this).f7, globalState)));
        let _rhs559 = _2867_v33;
        let _lhs292 = _2868_v34;
        let _lhs293 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_2868_v34).length));
        _2866_v32 = _rhs557;
        _2859_v26 = _rhs558;
        _lhs292[_lhs293] = _rhs559;
      }
      r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f7, (_this).f7)));
      let _2872_v38;
      _2872_v38 = false;
      let _2873_v39;
      _2873_v39 = _dafny.Seq.of(_2872_v38, false);
      let _2874_v40;
      _2874_v40 = _dafny.Set.fromElements((_this).f7);
      let _2875_v41;
      _2875_v41 = _dafny.Seq.of(_2874_v40, _2874_v40);
      let _2876_v42;
      _2876_v42 = _dafny.Seq.of(_2873_v39);
      let _2877_v43;
      _2877_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_2873_v39), _module.__default.safeIndex(new BigNumber((_2875_v41).length), new BigNumber((_dafny.Seq.of(_2873_v39)).length)), _dafny.Seq.of(_2872_v38, _2872_v38, _2872_v38)), _2876_v42));
      _2877_v43 = (_2877_v43).update((_this).f7, _dafny.Seq.of(_2873_v39, _2873_v39));
      let _2878_v44;
      _2878_v44 = _dafny.Map.Empty.slice().updateUnsafe(false,_2872_v38);
      let _2879_v45;
      _2879_v45 = _2878_v44;
      let _2880_v46;
      let _nw548 = Array((new BigNumber(28)).toNumber());
      _nw548[(_dafny.ZERO).toNumber()] = _2878_v44;
      _nw548[(_dafny.ONE).toNumber()] = (_2878_v44).Merge(_2878_v44);
      _nw548[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38)).Merge(_2878_v44);
      _nw548[(new BigNumber(3)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(4)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(5)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38)).update((_2873_v39)[_module.__default.safeIndex((_this).f7, new BigNumber((_2873_v39).length))], _2872_v38)).Merge(_2878_v44);
      _nw548[(new BigNumber(6)).toNumber()] = (_2878_v44).update(_2872_v38, _2872_v38);
      _nw548[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38);
      _nw548[(new BigNumber(8)).toNumber()] = ((_2872_v38) ? (_2878_v44) : (_2878_v44));
      _nw548[(new BigNumber(9)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(10)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(11)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(12)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38);
      _nw548[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38);
      _nw548[(new BigNumber(15)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38));
      _nw548[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38);
      _nw548[(new BigNumber(18)).toNumber()] = (_2878_v44).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2872_v38));
      _nw548[(new BigNumber(19)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(20)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(21)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(22)).toNumber()] = (_2878_v44).Merge(_2878_v44);
      _nw548[(new BigNumber(23)).toNumber()] = ((_2872_v38) ? (_2878_v44) : (_dafny.Map.Empty.slice().updateUnsafe(_2872_v38,true)));
      _nw548[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,!(true));
      _nw548[(new BigNumber(25)).toNumber()] = (_2879_v45);
      _nw548[(new BigNumber(26)).toNumber()] = _2878_v44;
      _nw548[(new BigNumber(27)).toNumber()] = _2878_v44;
      _2880_v46 = _nw548;
      let _index560 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_2880_v46).length));
      (_2880_v46)[_index560] = _2878_v44;
      let _2881_v47;
      _2881_v47 = _dafny.Map.Empty.slice().updateUnsafe(((_2872_v38) ? (false) : (!(_2872_v38))),_2878_v44);
      let _index561 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_2880_v46).length));
      (_2880_v46)[_index561] = (((_2881_v47).contains(_2872_v38)) ? ((_2881_v47).get(_2872_v38)) : (_2878_v44));
      let _2882_v48;
      let _nw549 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
      _2882_v48 = _nw549;
      let _2883_v49;
      _2883_v49 = _module.D13.create_DC28(_2882_v48);
      let _source26 = _2883_v49;
      if (_source26.is_DC29) {
        let _2884___mcc_h0 = (_source26).cf45;
        let _2885___mcc_h1 = (_source26).cf46;
        let _2886_cf46 = _2885___mcc_h1;
        let _2887_cf45 = _2884___mcc_h0;
        let _2888_v50;
        let _nw550 = new _module.C0();
        _nw550.__ctor();
        _2888_v50 = _nw550;
        let _2889_v51;
        let _nw551 = Array((new BigNumber(20)).toNumber()).fill(_module.D26.Default());
        _2889_v51 = _nw551;
        _2889_v51 = _2889_v51;
        if (_2886_cf46) {
          _2886_cf46 = _2886_cf46;
          let _2890_v52;
          _2890_v52 = _dafny.Seq.UnicodeFromString("wa");
          let _2891_v53;
          _2891_v53 = _dafny.Seq.of((_this).f7);
          let _2892_v54;
          _2892_v54 = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,_2891_v53);
          let _2893_v55;
          _2893_v55 = _module.D1.create_DC2((_this).f7, (_this).f7);
          let _2894_v56;
          _2894_v56 = _dafny.Set.fromElements(_2890_v52);
          let _2895_v57;
          let _nw552 = new _module.C5();
          _nw552.__ctor();
          _2895_v57 = _nw552;
          let _2896_v58;
          _2896_v58 = _dafny.Seq.of(_2895_v57);
          let _2897_v59;
          _2897_v59 = _module.D27.create_DC64(_2896_v58);
          let _2898_v60;
          let _nw553 = Array((new BigNumber(22)).toNumber());
          _nw553[(_dafny.ZERO).toNumber()] = (_this).f7;
          _nw553[(_dafny.ONE).toNumber()] = new BigNumber((_2878_v44).length);
          _nw553[(new BigNumber(2)).toNumber()] = new BigNumber(144);
          _nw553[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(4)).toNumber()] = new BigNumber((_2890_v52).length);
          _nw553[(new BigNumber(5)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(6)).toNumber()] = ((_this).f7).minus(new BigNumber(((((_2892_v54).contains(!(_2886_cf46))) ? ((_2892_v54).get(!(_2886_cf46))) : (_2891_v53))).length));
          _nw553[(new BigNumber(7)).toNumber()] = ((_this).f7).plus(new BigNumber((_2891_v53).length));
          _nw553[(new BigNumber(8)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(9)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(10)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(11)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(12)).toNumber()] = new BigNumber((_2890_v52).length);
          _nw553[(new BigNumber(13)).toNumber()] = _module.__default.fm14(globalState);
          _nw553[(new BigNumber(14)).toNumber()] = (new BigNumber(493)).multipliedBy((_this).f7);
          _nw553[(new BigNumber(15)).toNumber()] = (_this).fm7(_2893_v55, _2887_cf45, (_this).f7, _2872_v38, globalState);
          _nw553[(new BigNumber(16)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(17)).toNumber()] = (_this).f7;
          _nw553[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2894_v56).length), (_this).f7);
          _nw553[(new BigNumber(19)).toNumber()] = ((_this).f7).minus((_this).f7);
          _nw553[(new BigNumber(20)).toNumber()] = new BigNumber(((_2897_v59).dtor_cf106).length);
          _nw553[(new BigNumber(21)).toNumber()] = (_this).f7;
          _2898_v60 = _nw553;
          _2898_v60 = _2898_v60;
          _2886_cf46 = _2872_v38;
          let _2899_v61;
          let _nw554 = new _module.C10();
          _nw554.__ctor((_dafny.ZERO).minus((_this).f7), _2887_cf45, (_this).f4, (_this).f5);
          _2899_v61 = _nw554;
          r1 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_2899_v61)).length), (_this).f7)).multipliedBy(new BigNumber(895));
          let _2900_v62;
          let _nw555 = Array((_dafny.ONE).toNumber()).fill(false);
          _2900_v62 = _nw555;
          let _index562 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_2900_v62).length));
          (_2900_v62)[_index562] = _2872_v38;
          let _index563 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_2900_v62).length));
          (_2900_v62)[_index563] = _2872_v38;
        } else {
          let _2901_v63;
          _2901_v63 = new _dafny.CodePoint('g'.codePointAt(0));
          _2901_v63 = _2901_v63;
          r1 = (new BigNumber(75)).minus((_this).f7);
          let _2902_v64;
          _2902_v64 = _dafny.Seq.of((_this).f5);
          let _2903_v65;
          let _nw556 = new _module.C2();
          _nw556.__ctor((_this).f4, (_2902_v64)[_module.__default.safeIndex((_this).f7, new BigNumber((_2902_v64).length))]);
          _2903_v65 = _nw556;
          _2886_cf46 = _2887_cf45;
          let _2904_v66;
          _2904_v66 = _module.D26.create_DC63(_2886_cf46, (_this).fm6((_this).f7, (_this).f7, globalState), new BigNumber(219), _2887_cf45);
          let _rhs560 = !(_2886_cf46) || (_2886_cf46);
          let _rhs561 = (_2904_v66).dtor_cf104;
          _2887_cf45 = _rhs560;
          r1 = _rhs561;
        }
        let _2905_v67;
        let _nw557 = new _module.C8();
        _nw557.__ctor();
        _2905_v67 = _nw557;
      } else {
        let _2906___mcc_h2 = (_source26).cf44;
        let _2907_cf44 = _2906___mcc_h2;
        let _2908_v68;
        let _nw558 = Array((new BigNumber(2)).toNumber());
        _nw558[(_dafny.ZERO).toNumber()] = _2872_v38;
        _nw558[(_dafny.ONE).toNumber()] = _2872_v38;
        _2908_v68 = _nw558;
        let _index564 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length));
        (_2908_v68)[_index564] = (_this).fm5((_this).f7, _2872_v38, (_this).f7, (_this).f7, globalState);
        let _2909_v69;
        _2909_v69 = _dafny.Seq.UnicodeFromString("lxiotcre");
        let _index565 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length));
        (_2908_v68)[_index565] = _dafny.areEqual(_dafny.Seq.UnicodeFromString("fhybnu"), _2909_v69);
        let _2910_v70;
        _2910_v70 = _dafny.MultiSet.fromElements((_this).f7);
        let _2911_v71;
        _2911_v71 = _module.D2.create_DC4(!((_this).fm5((_this).f7, (_2908_v68)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length))], (((_2910_v70).contains((_this).f7)) ? ((_2910_v70).get((_this).f7)) : (new BigNumber(663))), (_this).f7, globalState)));
        _2911_v71 = _2911_v71;
        let _2912_v72;
        _2912_v72 = _module.D23.create_DC56(!(_2872_v38), (_2908_v68)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length))], _module.__default.fm1((_this).f7, (_this).f7, (_2908_v68)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length))], _2872_v38, globalState), _2872_v38);
        _2872_v38 = (((_2878_v44).contains((_2908_v68)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length))])) ? ((_2878_v44).get((_2908_v68)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_2908_v68).length))])) : ((_2912_v72).dtor_cf92));
        let _2913_v73;
        let _nw559 = new _module.C7();
        _nw559.__ctor();
        _2913_v73 = _nw559;
        let _nw560 = new _module.C7();
        _nw560.__ctor();
        _2913_v73 = _nw560;
      }
      let _2914_v74;
      let _init107 = function (_2915_i4) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      };
      let _nw561 = Array((new BigNumber(18)).toNumber());
      for (let _i0_107 = 0; _i0_107 < new BigNumber(_nw561.length); _i0_107++) {
        _nw561[_i0_107] = _init107(new BigNumber(_i0_107));
      }
      _2914_v74 = _nw561;
      let _2916_v75;
      _2916_v75 = _dafny.Seq.UnicodeFromString("cs");
      let _index566 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_2914_v74).length));
      (_2914_v74)[_index566] = (_2916_v75)[_module.__default.safeIndex((_this).f7, new BigNumber((_2916_v75).length))];
      let _index567 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_2914_v74).length));
      let _rhs562 = ((_2872_v38) ? ((_this).f5) : ((_this).f5));
      let _rhs563 = _module.__default.safeDivisionInt((_this).f7, (_this).f7);
      let _rhs564 = ((((_this).f7).isEqualTo((_this).f7)) ? (!(_2872_v38)) : (_2872_v38));
      let _lhs294 = _2914_v74;
      let _lhs295 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_2914_v74).length));
      _lhs294[_lhs295] = _rhs562;
      r1 = _rhs563;
      _2872_v38 = _rhs564;
      let _2917_v76;
      _2917_v76 = _dafny.Map.Empty.slice().updateUnsafe(_2872_v38,(_this).f7);
      let _2918_v77;
      _2918_v77 = _dafny.Seq.of((_this).f7, (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f7, new BigNumber((_2917_v76).length))));
      r0 = _2918_v77;
      r1 = (_this).f7;
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = [];
      let _2919_v0;
      let _nw562 = new _module.C0();
      _nw562.__ctor();
      _2919_v0 = _nw562;
      let _2920_v1;
      _2920_v1 = false;
      if (_2920_v1) {
        let _2921_v2;
        let _nw563 = new _module.C10();
        _nw563.__ctor(_module.__default.safeDivisionInt((_this).f7, (_this).f7), _2920_v1, (_this).f4, new _dafny.CodePoint('y'.codePointAt(0)));
        _2921_v2 = _nw563;
        _2921_v2 = _2921_v2;
        let _2922_v3;
        let _nw564 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _2922_v3 = _nw564;
        let _2923_v4;
        _2923_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2922_v3,_2920_v1);
        let _2924_v5;
        _2924_v5 = _dafny.Seq.of((_this).f7, new BigNumber((_dafny.Seq.of(_2920_v1)).length));
        let _2925_v6;
        _2925_v6 = _dafny.Set.fromElements(_2920_v1, !(!(_2923_v4).contains(_2922_v3)), !((_this).f7).isEqualTo((_2924_v5)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_2924_v5).length))]));
        _2925_v6 = (_2925_v6).Difference((_2925_v6).Difference(_2925_v6));
        let _index568 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        (_2922_v3)[_index568] = (_this).f7;
        let _index569 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        (_2922_v3)[_index569] = (_this).f7;
        let _index570 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        let _rhs565 = _2920_v1;
        let _rhs566 = (_2922_v3)[_module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length))];
        let _lhs296 = _2922_v3;
        let _lhs297 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        _2920_v1 = _rhs565;
        _lhs296[_lhs297] = _rhs566;
        let _index571 = _module.__default.safeIndex(new BigNumber(7), new BigNumber(((_2921_v2).f4).length));
        ((_2921_v2).f4)[_index571] = _2925_v6;
        let _2926_v7;
        _2926_v7 = _module.D15.create_DC35((_2921_v2).fm4(_dafny.Seq.of((_this).f7, (_2922_v3)[_module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length))]), globalState));
        let _index572 = _module.__default.safeIndex(new BigNumber(7), new BigNumber(((_2921_v2).f4).length));
        let _index573 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        let _rhs567 = _2920_v1;
        let _rhs568 = !(_2920_v1) || (_2920_v1);
        let _rhs569 = (_2925_v6).Union(_dafny.Set.fromElements(true));
        let _rhs570 = (_2926_v7).dtor_cf54;
        let _lhs298 = (_2921_v2).f4;
        let _lhs299 = _module.__default.safeIndex(new BigNumber(7), new BigNumber(((_2921_v2).f4).length));
        let _lhs300 = _2922_v3;
        let _lhs301 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_2922_v3).length));
        _2920_v1 = _rhs567;
        _2920_v1 = _rhs568;
        _lhs298[_lhs299] = _rhs569;
        _lhs300[_lhs301] = _rhs570;
      } else {
        let _2927_v8;
        let _nw565 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _2927_v8 = _nw565;
        let _2928_v9;
        _2928_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2920_v1,new BigNumber((_dafny.Set.fromElements((_this).f7)).length));
        let _index574 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        (_2927_v8)[_index574] = new BigNumber(((_2928_v9).Merge((_2928_v9).update(_2920_v1, (_dafny.ZERO).minus((_this).f7)))).length);
        let _index575 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        (_2927_v8)[_index575] = (_this).f7;
        let _2929_v10;
        _2929_v10 = _module.D11.create_DC25(_2920_v1, (_this).f7, _2927_v8, (_this).f7);
        let _2930_v11;
        _2930_v11 = _dafny.Seq.of(_2927_v8, _2927_v8, _2927_v8);
        let _2931_v12;
        let _nw566 = Array((new BigNumber(21)).toNumber());
        _nw566[(_dafny.ZERO).toNumber()] = _2927_v8;
        _nw566[(_dafny.ONE).toNumber()] = (_2929_v10).dtor_cf40;
        _nw566[(new BigNumber(2)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(3)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(4)).toNumber()] = (_2929_v10).dtor_cf40;
        _nw566[(new BigNumber(5)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(6)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(7)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(8)).toNumber()] = (_2929_v10).dtor_cf40;
        _nw566[(new BigNumber(9)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(10)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(11)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(12)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(13)).toNumber()] = (_2930_v11)[_module.__default.safeIndex((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], new BigNumber((_2930_v11).length))];
        _nw566[(new BigNumber(14)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(15)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(16)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(17)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(18)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(19)).toNumber()] = _2927_v8;
        _nw566[(new BigNumber(20)).toNumber()] = _2927_v8;
        _2931_v12 = _nw566;
        let _2932_v13;
        let _nw567 = Array((new BigNumber(15)).toNumber()).fill(false);
        _2932_v13 = _nw567;
        let _2933_v14;
        _2933_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2931_v12,_2932_v13);
        let _2934_v15;
        _2934_v15 = _dafny.Set.fromElements((_this).f7);
        let _index576 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        let _index577 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        let _rhs571 = (new BigNumber(-119)).minus((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))]);
        let _rhs572 = (((_2933_v14).contains(_2931_v12)) ? ((_2933_v14).get(_2931_v12)) : (_2932_v13));
        let _rhs573 = _module.__default.safeDivisionInt(new BigNumber((_2934_v15).length), (_dafny.ZERO).minus((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))]));
        let _lhs302 = _2927_v8;
        let _lhs303 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        let _lhs304 = globalState;
        let _lhs305 = _2927_v8;
        let _lhs306 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
        _lhs302[_lhs303] = _rhs571;
        _lhs304.f3 = _rhs572;
        _lhs305[_lhs306] = _rhs573;
        _2920_v1 = !(_2920_v1);
        let _2935_v16;
        _2935_v16 = _dafny.Set.fromElements(_2920_v1);
        if ((_2935_v16).IsProperSubsetOf((_2935_v16).Intersect(_2935_v16))) {
          let _2936_v17;
          let _nw568 = new _module.C3();
          _nw568.__ctor((_this).f4, (_this).f5);
          _2936_v17 = _nw568;
          let _2937_v18;
          _2937_v18 = _module.D26.create_DC63(_2920_v1, (_this).f7, (_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], _2920_v1);
          _2920_v1 = (_2937_v18).dtor_cf102;
          let _2938_v19;
          _2938_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-874)), ((_2939_v8) => function (_2940_i0) {
            return (_2939_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2939_v8).length))];
          })(_2927_v8)),(_this).f7);
          let _2941_v20;
          _2941_v20 = _dafny.Seq.of((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], (_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))]);
          let _2942_v21;
          _2942_v21 = _module.D1.create_DC2((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], (_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))]);
          _2938_v19 = (_2938_v19).update(_2941_v20, (_this).fm7(_2942_v21, _module.__default.fm1((_this).f7, (_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], _2920_v1, _2920_v1, globalState), (_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], _2920_v1, globalState));
          let _index578 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
          (_2927_v8)[_index578] = (_this).f7;
          let _index579 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
          (_2927_v8)[_index579] = (_this).f7;
        } else {
          let _2943_v22;
          _2943_v22 = _dafny.MultiSet.fromElements((_this).f7);
          _2920_v1 = !(!(((_2943_v22).update((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], _module.__default.abs(new BigNumber(50)))).IsSubsetOf(_2943_v22)) || ((_2920_v1) || (_2920_v1)));
          let _2944_v23;
          _2944_v23 = _dafny.Seq.of(_2920_v1, !(_2920_v1), _2920_v1, false, _2920_v1);
          _2944_v23 = _dafny.Seq.Concat(_2944_v23, _dafny.Seq.of(!(_2920_v1), _2920_v1, _2920_v1, _2920_v1));
          let _index580 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length));
          (_2927_v8)[_index580] = (_this).f7;
          let _2945_v24;
          _2945_v24 = _dafny.Seq.UnicodeFromString("fjves");
          let _2946_v25;
          _2946_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2945_v24,_2920_v1);
          let _2947_v26;
          let _nw569 = new _module.C5();
          _nw569.__ctor();
          _2947_v26 = _nw569;
          let _2948_v27;
          _2948_v27 = _dafny.Seq.of(_2947_v26);
          let _2949_v28;
          _2949_v28 = _module.D21.create_DC51((_2944_v23)[_module.__default.safeIndex((_this).f7, new BigNumber((_2944_v23).length))], (_2948_v27)[_module.__default.safeIndex((_this).f7, new BigNumber((_2948_v27).length))]);
          let _2950_v29;
          _2950_v29 = _module.D21.create_DC51(!(_2920_v1), (_2949_v28).dtor_cf81);
          let _rhs574 = ((_2920_v1) && ((_2950_v29).dtor_cf80)) || ((((_2944_v23)[_module.__default.safeIndex((_2927_v8)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_2927_v8).length))], new BigNumber((_2944_v23).length))]) ? (_2920_v1) : (_2920_v1)));
          let _rhs575 = (_2946_v25).Merge(_2946_v25);
          _2920_v1 = _rhs574;
          _2946_v25 = _rhs575;
          let _2951_v30;
          _2951_v30 = _dafny.MultiSet.fromElements(_2920_v1);
          _2920_v1 = !((_2951_v30).IsSubsetOf((_2951_v30).update(_2920_v1, _module.__default.abs((_this).f7))));
        }
        let _2952_v32;
        let _nw570 = new _module.C6();
        _nw570.__ctor(function () {
          let _coll73 = new _dafny.Map();
          for (const _compr_73 of _dafny.IntegerRange(new BigNumber(72), new BigNumber(862))) {
            let _2953_v31 = _compr_73;
            if (((new BigNumber(72)).isLessThanOrEqualTo(_2953_v31)) && ((_2953_v31).isLessThan(new BigNumber(862)))) {
              _coll73.push([(_2953_v31).multipliedBy(new BigNumber(465)),(_dafny.ZERO).minus((_this).f7)]);
            }
          }
          return _coll73;
        }());
        _2952_v32 = _nw570;
        _2952_v32 = _2952_v32;
      }
      let _2954_v33;
      _2954_v33 = _module.D7.create_DC17(_2920_v1, _2920_v1, _2920_v1, _2920_v1);
      let _source27 = _2954_v33;
      if (_source27.is_DC17) {
        let _2955___mcc_h0 = (_source27).cf24;
        let _2956___mcc_h1 = (_source27).cf25;
        let _2957___mcc_h2 = (_source27).cf26;
        let _2958___mcc_h3 = (_source27).cf27;
        let _2959_cf27 = _2958___mcc_h3;
        let _2960_cf26 = _2957___mcc_h2;
        let _2961_cf25 = _2956___mcc_h1;
        let _2962_cf24 = _2955___mcc_h0;
        _2961_cf25 = _2920_v1;
        let _2963_v34;
        _2963_v34 = _dafny.Seq.UnicodeFromString("qybbhcq");
        let _2964_v35;
        let _nw571 = new _module.C10();
        _nw571.__ctor(new BigNumber((_2963_v34).length), _2960_cf26, (_this).f4, (_this).f5);
        _2964_v35 = _nw571;
        let _2965_v36;
        let _nw572 = Array((_dafny.ONE).toNumber()).fill(false);
        _2965_v36 = _nw572;
        let _2966_v37;
        _2966_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2964_v35,_2965_v36);
        _2966_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2964_v35,_2965_v36);
        _2961_cf25 = _2960_cf26;
        let _2967_v38;
        _2967_v38 = new BigNumber(567);
        let _2968_v39;
        _2968_v39 = _dafny.Set.fromElements(_2962_cf24, _2960_cf26);
        let _index581 = _module.__default.safeIndex(new BigNumber(138), new BigNumber(((_2964_v35).f4).length));
        ((_2964_v35).f4)[_index581] = _2968_v39;
        let _index582 = _module.__default.safeIndex(new BigNumber(138), new BigNumber(((_2964_v35).f4).length));
        let _rhs576 = _2960_cf26;
        let _rhs577 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), function (_2969_i1) {
          return (_this).f5;
        })).length);
        let _rhs578 = (_2968_v39).Difference(_2968_v39);
        let _lhs307 = (_2964_v35).f4;
        let _lhs308 = _module.__default.safeIndex(new BigNumber(138), new BigNumber(((_2964_v35).f4).length));
        _2962_cf24 = _rhs576;
        _2967_v38 = _rhs577;
        _lhs307[_lhs308] = _rhs578;
      } else {
        let _2970___mcc_h4 = (_source27).cf23;
        let _2971_cf23 = _2970___mcc_h4;
        let _2972_v40;
        let _init108 = function (_2973_i2) {
          return _module.__default.safeModuloInt(_2973_i2, (_this).f7);
        };
        let _nw573 = Array((new BigNumber(29)).toNumber());
        for (let _i0_108 = 0; _i0_108 < new BigNumber(_nw573.length); _i0_108++) {
          _nw573[_i0_108] = _init108(new BigNumber(_i0_108));
        }
        _2972_v40 = _nw573;
        let _2974_v41;
        let _2975_v42;
        let _2976_v43;
        let _out33;
        let _out34;
        let _out35;
        let _outcollector9 = (_this).m3((_this).f7, _2972_v40, globalState);
        _out33 = _outcollector9[0];
        _out34 = _outcollector9[1];
        _out35 = _outcollector9[2];
        _2974_v41 = _out33;
        _2975_v42 = _out34;
        _2976_v43 = _out35;
        let _2977_v44;
        let _nw574 = new _module.C5();
        _nw574.__ctor();
        _2977_v44 = _nw574;
        let _2978_v45;
        _2978_v45 = _dafny.MultiSet.fromElements((_this).f7, (_dafny.ZERO).minus((_this).f7), (_this).f7, (_this).f7);
        let _2979_v46;
        _2979_v46 = _dafny.MultiSet.fromElements(new BigNumber((_2978_v45).cardinality()), (_this).f7, (_dafny.ZERO).minus((_this).f7), (_this).f7, (_this).f7);
        let _2980_v47;
        _2980_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(((_2978_v45).contains((_this).f7)) ? ((_2978_v45).get((_this).f7)) : (new BigNumber(-441))));
        let _2981_v48;
        _2981_v48 = _dafny.Seq.of((_this).f7);
        let _2982_v49;
        let _nw575 = Array((new BigNumber(14)).toNumber());
        _nw575[(_dafny.ZERO).toNumber()] = (_2920_v1) === (_2920_v1);
        _nw575[(_dafny.ONE).toNumber()] = _2920_v1;
        _nw575[(new BigNumber(2)).toNumber()] = _2974_v41;
        _nw575[(new BigNumber(3)).toNumber()] = false;
        _nw575[(new BigNumber(4)).toNumber()] = _module.__default.fm1((((_2980_v47).contains((_this).f7)) ? ((_2980_v47).get((_this).f7)) : ((_this).f7)), new BigNumber((_dafny.MultiSet.FromArray(_2981_v48)).cardinality()), _2974_v41, _2920_v1, globalState);
        _nw575[(new BigNumber(5)).toNumber()] = _2974_v41;
        _nw575[(new BigNumber(6)).toNumber()] = _2975_v42;
        _nw575[(new BigNumber(7)).toNumber()] = !(_2974_v41);
        _nw575[(new BigNumber(8)).toNumber()] = _2920_v1;
        _nw575[(new BigNumber(9)).toNumber()] = _2920_v1;
        _nw575[(new BigNumber(10)).toNumber()] = _2975_v42;
        _nw575[(new BigNumber(11)).toNumber()] = _2920_v1;
        _nw575[(new BigNumber(12)).toNumber()] = _2920_v1;
        _nw575[(new BigNumber(13)).toNumber()] = _2920_v1;
        _2982_v49 = _nw575;
        (globalState).f3 = _2982_v49;
        let _2983_v50;
        _2983_v50 = _dafny.Seq.of((_module.__default.fm49(globalState)).dtor_cf75);
        let _2984_v52;
        _2984_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_2974_v41);
        let _2985_v53;
        _2985_v53 = _dafny.Set.fromElements(_2975_v42, _2920_v1, (_this).fm3((_this).f7, new BigNumber((_2984_v52).length), globalState));
        let _2986_v54;
        _2986_v54 = _dafny.Set.fromElements(_module.__default.fm21((_this).f7, new BigNumber((_2981_v48).length), _2985_v53, globalState), _dafny.Seq.of(_2974_v41, _2974_v41));
        let _2987_v55;
        _2987_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2974_v41,_2986_v54);
        if ((function () {
          let _coll74 = new _dafny.Set();
          for (const _compr_74 of (_2983_v50).Elements) {
            let _2988_v51 = _compr_74;
            if (_dafny.Seq.contains(_2983_v50, _2988_v51)) {
              _coll74.add(_2988_v51);
            }
          }
          return _coll74;
        }()).IsDisjointFrom((((_2987_v55).contains(_2920_v1)) ? ((_2987_v55).get(_2920_v1)) : (_module.__default.fm50(_2975_v42, _2920_v1, _2975_v42, globalState))))) {
          _2975_v42 = _2974_v41;
          _2979_v46 = _dafny.MultiSet.FromArray(_2981_v48);
          let _2989_v56;
          let _nw576 = new _module.C6();
          _nw576.__ctor(_2980_v47);
          _2989_v56 = _nw576;
          let _2990_v57;
          _2990_v57 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2989_v56,(_this).fm5((_this).f7, !(!(_2920_v1)), (_this).f7, (_this).f7, globalState)));
          let _2991_v58;
          _2991_v58 = _dafny.Seq.of(false);
          let _2992_v59;
          _2992_v59 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2991_v58)[_module.__default.safeIndex((_this).f7, new BigNumber((_2991_v58).length))]);
          let _2993_v60;
          let _2994_v61;
          let _2995_v62;
          let _2996_v63;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector10 = (_2977_v44).m10(new BigNumber((_2990_v57).cardinality()), _2992_v59, _2975_v42, globalState);
          _out36 = _outcollector10[0];
          _out37 = _outcollector10[1];
          _out38 = _outcollector10[2];
          _out39 = _outcollector10[3];
          _2993_v60 = _out36;
          _2994_v61 = _out37;
          _2995_v62 = _out38;
          _2996_v63 = _out39;
          let _index583 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2982_v49).length));
          (_2982_v49)[_index583] = _2920_v1;
          let _index584 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2982_v49).length));
          (_2982_v49)[_index584] = _2920_v1;
          let _2997_v64;
          let _nw577 = new _module.C2();
          _nw577.__ctor((_this).f4, (_this).f5);
          _2997_v64 = _nw577;
        } else {
          _2975_v42 = !(_2975_v42);
          r0 = _2972_v40;
          let _2998_v65;
          let _nw578 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2998_v65 = _nw578;
          let _index585 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2998_v65).length));
          (_2998_v65)[_index585] = _2976_v43;
          let _2999_v66;
          _2999_v66 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("insl"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), function (_3000_i3) {
            return (_this).f5;
          }), _2976_v43, _2976_v43);
          let _3001_v67;
          let _nw579 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _3001_v67 = _nw579;
          let _3002_v68;
          _3002_v68 = _dafny.MultiSet.fromElements(_3001_v67);
          let _3003_v69;
          _3003_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2974_v41,(_this).f7);
          let _index586 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_2998_v65).length));
          (_2998_v65)[_index586] = (_2999_v66)[_module.__default.safeIndex(_module.__default.safeModuloInt((((_3002_v68).contains(_3001_v67)) ? ((_3002_v68).get(_3001_v67)) : ((((_3003_v69).contains(_2975_v42)) ? ((_3003_v69).get(_2975_v42)) : ((_dafny.ZERO).minus((_this).f7))))), (_this).f7), new BigNumber((_2999_v66).length))];
          let _3004_v70;
          _3004_v70 = _module.D7.create_DC16(_2971_cf23);
          let _3005_v71;
          let _nw580 = new _module.C9();
          _nw580.__ctor((_this).f4, (_3004_v70).dtor_cf23);
          _3005_v71 = _nw580;
          let _3006_v72;
          _3006_v72 = new BigNumber(889);
          let _3007_v73;
          _3007_v73 = _dafny.Seq.of(_2977_v44);
          let _3008_v74;
          _3008_v74 = _dafny.Seq.of((_3007_v73)[_module.__default.safeIndex(new BigNumber(444), new BigNumber((_3007_v73).length))], _2977_v44);
          let _3009_v75;
          _3009_v75 = _module.D18.create_DC44(_dafny.Set.fromElements(_2985_v53));
          let _3010_v76;
          _3010_v76 = _dafny.Seq.of(_3009_v75);
          let _3011_v77;
          _3011_v77 = _dafny.Seq.of(!(_2974_v41));
          _3006_v72 = ((_dafny.areEqual(_3008_v74, _3007_v73)) ? ((new BigNumber((_3010_v76).length)).multipliedBy(new BigNumber(((_2998_v65)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_2998_v65).length))]).length))) : (new BigNumber((((_2920_v1) ? (_3011_v77) : (_3011_v77))).length)));
        }
      }
      let _3012_v78;
      _3012_v78 = new BigNumber(824);
      let _3013_v79;
      _3013_v79 = _dafny.Seq.UnicodeFromString("a");
      let _rhs579 = new BigNumber(116);
      let _rhs580 = _dafny.Seq.Concat(((_2920_v1) ? (_3013_v79) : (_3013_v79)), _3013_v79);
      let _rhs581 = _2920_v1;
      let _rhs582 = _3012_v78;
      _3012_v78 = _rhs579;
      _3013_v79 = _rhs580;
      _2920_v1 = _rhs581;
      _3012_v78 = _rhs582;
      let _3014_v80;
      _3014_v80 = _dafny.Seq.of(_2920_v1);
      let _3015_v81;
      _3015_v81 = _module.D20.create_DC48(_2920_v1, _3014_v80, _3014_v80, _3012_v78);
      let _3016_v82;
      _3016_v82 = _module.D20.create_DC49(_3015_v81);
      let _source28 = _module.D20.create_DC49(_3015_v81);
      if (_source28.is_DC48) {
        let _3017___mcc_h5 = (_source28).cf74;
        let _3018___mcc_h6 = (_source28).cf75;
        let _3019___mcc_h7 = (_source28).cf76;
        let _3020___mcc_h8 = (_source28).cf77;
        let _3021_cf77 = _3020___mcc_h8;
        let _3022_cf76 = _3019___mcc_h7;
        let _3023_cf75 = _3018___mcc_h6;
        let _3024_cf74 = _3017___mcc_h5;
        _3024_cf74 = _2920_v1;
        let _3025_v83;
        let _nw581 = new _module.C5();
        _nw581.__ctor();
        _3025_v83 = _nw581;
        let _3026_v84;
        _3026_v84 = _dafny.Seq.of(_3025_v83, _3025_v83);
        let _3027_v85;
        _3027_v85 = _module.D27.create_DC64(_3026_v84);
        let _pat_let_tv57 = _3026_v84;
        let _pat_let_tv58 = _3021_cf77;
        let _pat_let_tv59 = _3026_v84;
        let _pat_let_tv60 = _3025_v83;
        _3027_v85 = function (_pat_let61_0) {
          return function (_3028_dt__update__tmp_h0) {
            return function (_pat_let62_0) {
              return function (_3029_dt__update_hcf106_h0) {
                return _module.D27.create_DC64(_3029_dt__update_hcf106_h0);
              }(_pat_let62_0);
            }(_dafny.Seq.update(_pat_let_tv57, _module.__default.safeIndex(_pat_let_tv58, new BigNumber((_pat_let_tv59).length)), _pat_let_tv60));
          }(_pat_let61_0);
        }(_3027_v85);
        let _3030_v86;
        _3030_v86 = _module.D13.create_DC29(_2920_v1, _2920_v1);
        let _3031_v87;
        _3031_v87 = _dafny.Set.fromElements((_this).f7);
        let _3032_v88;
        _3032_v88 = _module.D16.create_DC39(_module.D16.create_DC37(_3031_v87));
        let _3033_v89;
        _3033_v89 = _module.D16.create_DC37(_3031_v87);
        let _3034_v90;
        _3034_v90 = _dafny.MultiSet.fromElements(_3032_v88, _module.D16.create_DC39(_3033_v89), _3032_v88);
        let _3035_v91;
        _3035_v91 = _3034_v90;
        let _3036_v92;
        _3036_v92 = (_3035_v91);
        let _3037_v93;
        _3037_v93 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_3036_v92);
        let _rhs583 = _module.__default.safeDivisionInt((_3021_cf77).minus((_this).f7), _module.__default.safeModuloInt((_this).f7, _3012_v78));
        let _rhs584 = _module.__default.safeDivisionInt(_3012_v78, _3021_cf77);
        let _rhs585 = _3030_v86;
        let _rhs586 = _3012_v78;
        let _rhs587 = (new BigNumber(-443)).isLessThan(new BigNumber((function () {
          let _coll75 = new _dafny.Set();
          for (const _compr_75 of (_3037_v93).Keys.Elements) {
            let _3038_v94 = _compr_75;
            if ((_3037_v93).contains(_3038_v94)) {
              _coll75.add(_3038_v94);
            }
          }
          return _coll75;
        }()).length));
        _3012_v78 = _rhs583;
        _3012_v78 = _rhs584;
        _3030_v86 = _rhs585;
        _3012_v78 = _rhs586;
        _2920_v1 = _rhs587;
        let _3039_v95;
        let _nw582 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _3039_v95 = _nw582;
        let _index587 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_3039_v95).length));
        (_3039_v95)[_index587] = (_this).f7;
        let _3040_v96;
        let _nw583 = Array((new BigNumber(14)).toNumber());
        _nw583[(_dafny.ZERO).toNumber()] = _3024_cf74;
        _nw583[(_dafny.ONE).toNumber()] = _2920_v1;
        _nw583[(new BigNumber(2)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(3)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(4)).toNumber()] = true;
        _nw583[(new BigNumber(5)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(6)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(7)).toNumber()] = _2920_v1;
        _nw583[(new BigNumber(8)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(9)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(10)).toNumber()] = _2920_v1;
        _nw583[(new BigNumber(11)).toNumber()] = _2920_v1;
        _nw583[(new BigNumber(12)).toNumber()] = _3024_cf74;
        _nw583[(new BigNumber(13)).toNumber()] = _2920_v1;
        _3040_v96 = _nw583;
        let _3041_v97;
        _3041_v97 = _dafny.Set.fromElements(_3040_v96);
        let _3042_v98;
        _3042_v98 = _3041_v97;
        let _3043_v99;
        let _init109 = ((_3044_cf77, _3045_v1) => function (_3046_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe(_3044_cf77,_3045_v1);
        })(_3021_cf77, _2920_v1);
        let _nw584 = Array((new BigNumber(18)).toNumber());
        for (let _i0_109 = 0; _i0_109 < new BigNumber(_nw584.length); _i0_109++) {
          _nw584[_i0_109] = _init109(new BigNumber(_i0_109));
        }
        _3043_v99 = _nw584;
        let _3047_v100;
        _3047_v100 = _module.D4.create_DC10(((false) ? (_3042_v98) : (_3042_v98)), new BigNumber((_3031_v87).length), _module.__default.fm16(false, _3012_v78, globalState), _3043_v99);
        let _3048_v101;
        _3048_v101 = _dafny.Seq.of(_3012_v78);
        let _index588 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_3039_v95).length));
        let _rhs588 = new BigNumber((((_3024_cf74) ? (_3048_v101) : (_3048_v101))).length);
        let _rhs589 = _3012_v78;
        let _rhs590 = _3047_v100;
        let _rhs591 = _dafny.Seq.update(_3014_v80, _module.__default.safeIndex(((_3024_cf74) ? (_3012_v78) : (_3021_cf77)), new BigNumber((_3014_v80).length)), !(_2920_v1));
        let _rhs592 = _3012_v78;
        let _lhs309 = _3039_v95;
        let _lhs310 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_3039_v95).length));
        _lhs309[_lhs310] = _rhs588;
        _3012_v78 = _rhs589;
        _3047_v100 = _rhs590;
        _3022_cf76 = _rhs591;
        _3012_v78 = _rhs592;
      } else if (_source28.is_DC47) {
        let _3049___mcc_h9 = (_source28).cf73;
        let _3050_cf73 = _3049___mcc_h9;
        _2920_v1 = !(_2920_v1);
        let _3051_v102;
        let _nw585 = new _module.C8();
        _nw585.__ctor();
        _3051_v102 = _nw585;
        let _3052_v103;
        _3052_v103 = _module.D8.create_DC19(new BigNumber((_3013_v79).length), true);
        let _3053_v104;
        _3053_v104 = _module.D8.create_DC19(new BigNumber((_3014_v80).length), (_3052_v103).dtor_cf30);
        let _3054_v105;
        let _nw586 = Array((new BigNumber(9)).toNumber());
        _nw586[(_dafny.ZERO).toNumber()] = (_this).f7;
        _nw586[(_dafny.ONE).toNumber()] = (_this).f7;
        _nw586[(new BigNumber(2)).toNumber()] = _3012_v78;
        _nw586[(new BigNumber(3)).toNumber()] = _3012_v78;
        _nw586[(new BigNumber(4)).toNumber()] = (_this).f7;
        _nw586[(new BigNumber(5)).toNumber()] = (_this).f7;
        _nw586[(new BigNumber(6)).toNumber()] = (_this).f7;
        _nw586[(new BigNumber(7)).toNumber()] = (_3053_v104).dtor_cf29;
        _nw586[(new BigNumber(8)).toNumber()] = (_this).f7;
        _3054_v105 = _nw586;
        let _3055_v106;
        _3055_v106 = _dafny.MultiSet.fromElements(_3054_v105);
        let _3056_v107;
        _3056_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2920_v1,_dafny.MultiSet.fromElements(_3054_v105));
        _3055_v106 = ((_2920_v1) ? (_3055_v106) : (((((_3056_v107).contains(_2920_v1)) ? ((_3056_v107).get(_2920_v1)) : (_3055_v106))).Difference(_3055_v106)));
        _3012_v78 = _3012_v78;
      } else {
        let _3057___mcc_h10 = (_source28).cf78;
        let _3058_cf78 = _3057___mcc_h10;
        let _3059_v108;
        _3059_v108 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-718)), function (_3060_i5) {
          return (_this).f5;
        }), _module.__default.safeIndex((_this).f7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-718)), function (_3061_i5) {
          return (_this).f5;
        })).length)), (_this).f5)).length), new BigNumber((_dafny.MultiSet.fromElements(_2920_v1)).cardinality()), _3012_v78);
        let _3062_v109;
        _3062_v109 = _dafny.MultiSet.fromElements(_2920_v1, _2920_v1);
        let _3063_v110;
        _3063_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_3059_v108);
        let _3064_v111;
        let _nw587 = Array((new BigNumber(29)).toNumber());
        _nw587[(_dafny.ZERO).toNumber()] = _3059_v108;
        _nw587[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_3012_v78);
        _nw587[(new BigNumber(2)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_3065_v78) => function (_3066_i6) {
          return _3065_v78;
        })(_3012_v78));
        _nw587[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_3067_v78) => function (_3068_i7) {
          return _3067_v78;
        })(_3012_v78));
        _nw587[(new BigNumber(5)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_3059_v108, _module.__default.safeIndex(_3012_v78, new BigNumber((_3059_v108).length)), _3012_v78), _module.__default.fm31(_3012_v78, new BigNumber(763), globalState));
        _nw587[(new BigNumber(7)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(8)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_3059_v108, _3059_v108);
        _nw587[(new BigNumber(10)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(11)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), ((_3069_v79) => function (_3070_i8) {
          return new BigNumber((_3069_v79).length);
        })(_3013_v79));
        _nw587[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((((_3062_v109).contains(_2920_v1)) ? ((_3062_v109).get(_2920_v1)) : (_3012_v78))), _3059_v108);
        _nw587[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_3059_v108, _3059_v108);
        _nw587[(new BigNumber(15)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(16)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(17)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(18)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_3059_v108, _module.__default.safeIndex(new BigNumber(((((_3063_v110).contains(_3012_v78)) ? ((_3063_v110).get(_3012_v78)) : (_3059_v108))).length), new BigNumber((_3059_v108).length)), (_this).f7), _3059_v108);
        _nw587[(new BigNumber(20)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(21)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(22)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(23)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(24)).toNumber()] = _dafny.Seq.of((_this).f7);
        _nw587[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_3059_v108, _3059_v108);
        _nw587[(new BigNumber(26)).toNumber()] = _3059_v108;
        _nw587[(new BigNumber(27)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), ((_3071_v78) => function (_3072_i9) {
          return (_dafny.ZERO).minus(_3071_v78);
        })(_3012_v78));
        _nw587[(new BigNumber(28)).toNumber()] = _dafny.Seq.of((_this).fm6(new BigNumber((_3014_v80).length), new BigNumber(148), globalState));
        _3064_v111 = _nw587;
        let _index589 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_3064_v111).length));
        (_3064_v111)[_index589] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_3073_v78) => function (_3074_i10) {
          return _3073_v78;
        })(_3012_v78));
        let _index590 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_3064_v111).length));
        (_3064_v111)[_index590] = _3059_v108;
        _3012_v78 = (_this).f7;
        let _3075_v112;
        let _nw588 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _3075_v112 = _nw588;
        r0 = ((_2920_v1) ? (_3075_v112) : (_3075_v112));
        _2920_v1 = _2920_v1;
      }
      _3012_v78 = (_this).f7;
      let _3076_v113;
      let _nw589 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _3076_v113 = _nw589;
      r0 = _3076_v113;
      return r0;
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _3077_v0;
      _3077_v0 = false;
      let _3078_v1;
      _3078_v1 = _dafny.Set.fromElements(_3077_v0, _3077_v0, !(_3077_v0), _3077_v0);
      let _3079_v2;
      _3079_v2 = _dafny.Seq.of(_3077_v0);
      let _3080_v3;
      let _nw590 = Array((new BigNumber(13)).toNumber());
      _nw590[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_3077_v0, _3077_v0, _3077_v0);
      _nw590[(_dafny.ONE).toNumber()] = _3078_v1;
      _nw590[(new BigNumber(2)).toNumber()] = _3078_v1;
      _nw590[(new BigNumber(3)).toNumber()] = _3078_v1;
      _nw590[(new BigNumber(4)).toNumber()] = _3078_v1;
      _nw590[(new BigNumber(5)).toNumber()] = (_3078_v1).Intersect(_3078_v1);
      _nw590[(new BigNumber(6)).toNumber()] = _3078_v1;
      _nw590[(new BigNumber(7)).toNumber()] = (_3078_v1).Difference(_3078_v1);
      _nw590[(new BigNumber(8)).toNumber()] = (_dafny.Set.fromElements(_3077_v0, _3077_v0, _3077_v0)).Intersect(_3078_v1);
      _nw590[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_3077_v0, _3077_v0, (_3079_v2)[_module.__default.safeIndex((_this).f7, new BigNumber((_3079_v2).length))]);
      _nw590[(new BigNumber(10)).toNumber()] = (_3078_v1).Union(_3078_v1);
      _nw590[(new BigNumber(11)).toNumber()] = (_3078_v1).Difference(_3078_v1);
      _nw590[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(_module.__default.fm1((_this).f7, (_this).f7, _3077_v0, _module.__default.fm1(new BigNumber((_3078_v1).length), new BigNumber(483), _3077_v0, _3077_v0, globalState), globalState), _3077_v0, _3077_v0, false, _3077_v0);
      _3080_v3 = _nw590;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_3080_v3).length))) {
        let _3081_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_3081_i0)) && ((_3081_i0).isLessThan(new BigNumber((_3080_v3).length))))) {
          (_3080_v3)[(_3081_i0)] = _3078_v1;
        }
      }
      let _hi21 = p0;
      for (let _3082_i1 = (_this).f7; _3082_i1.isLessThan(_hi21); _3082_i1 = _3082_i1.plus(_dafny.ONE)) {
        if (_3077_v0) {
          let _3083_v4;
          _3083_v4 = new BigNumber(693);
          _3083_v4 = new BigNumber(967);
          _3077_v0 = _3077_v0;
          let _3084_v5;
          _3084_v5 = _dafny.Seq.of((_this).f7);
          let _3085_v6;
          _3085_v6 = _dafny.MultiSet.fromElements((_this).f7);
          _3083_v4 = new BigNumber((((_dafny.MultiSet.fromElements((_3084_v5)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_3084_v5).length))], _3083_v4, (((_3085_v6).contains(_3083_v4)) ? ((_3085_v6).get(_3083_v4)) : (_3082_i1)), new BigNumber(745))).Intersect(_3085_v6)).update((_this).f7, _module.__default.abs(new BigNumber(979)))).cardinality());
          let _3086_v7;
          _3086_v7 = _module.D2.create_DC4(_3077_v0);
          _3086_v7 = _3086_v7;
          _3078_v1 = (_3078_v1).Difference(_dafny.Set.fromElements(false, _3077_v0));
        } else {
          let _3087_v8;
          _3087_v8 = _module.D15.create_DC35(p0);
          let _3088_v9;
          _3088_v9 = new BigNumber(405);
          let _3089_v10;
          _3089_v10 = _dafny.MultiSet.fromElements(_module.__default.fm15(_3077_v0, _3077_v0, _3082_i1, globalState));
          let _3090_v11;
          _3090_v11 = _dafny.Seq.UnicodeFromString("ceearghr");
          let _rhs593 = _3087_v8;
          let _rhs594 = (((_3089_v10).contains(_dafny.MultiSet.fromElements(_3077_v0, _3077_v0, false))) ? ((_3089_v10).get(_dafny.MultiSet.fromElements(_3077_v0, _3077_v0, false))) : (new BigNumber((_3090_v11).length)));
          _3087_v8 = _rhs593;
          _3088_v9 = _rhs594;
          let _3091_v12;
          let _nw591 = new _module.C7();
          _nw591.__ctor();
          _3091_v12 = _nw591;
          _3088_v9 = p0;
          _3088_v9 = (_dafny.ZERO).minus(_3088_v9);
          _3088_v9 = (p0).minus((p0).minus((_this).f7));
        }
        r0 = _3077_v0;
        let _3092_v13;
        _3092_v13 = new BigNumber(607);
        _3092_v13 = (_this).f7;
        _3092_v13 = ((new BigNumber(54)).plus(new BigNumber(-703))).multipliedBy(new BigNumber(293));
      }
      _3077_v0 = !((_3077_v0) === (_3077_v0));
      let _3093_v14;
      let _nw592 = Array((new BigNumber(7)).toNumber());
      _nw592[(_dafny.ZERO).toNumber()] = _3077_v0;
      _nw592[(_dafny.ONE).toNumber()] = _3077_v0;
      _nw592[(new BigNumber(2)).toNumber()] = _3077_v0;
      _nw592[(new BigNumber(3)).toNumber()] = _3077_v0;
      _nw592[(new BigNumber(4)).toNumber()] = _3077_v0;
      _nw592[(new BigNumber(5)).toNumber()] = _3077_v0;
      _nw592[(new BigNumber(6)).toNumber()] = (_3079_v2)[_module.__default.safeIndex(p0, new BigNumber((_3079_v2).length))];
      _3093_v14 = _nw592;
      let _3094_v15;
      _3094_v15 = _dafny.Seq.of(_3093_v14, _3093_v14, _3093_v14, _3093_v14, _3093_v14);
      let _3095_v16;
      _3095_v16 = _dafny.Map.Empty.slice().updateUnsafe((_3094_v15)[_module.__default.safeIndex(p0, new BigNumber((_3094_v15).length))],_3077_v0);
      let _3096_v17;
      _3096_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_3077_v0);
      _3077_v0 = (((_3095_v16).contains(_3093_v14)) ? ((_3095_v16).get(_3093_v14)) : ((((_3096_v17).contains(_module.__default.fm14(globalState))) ? ((_3096_v17).get(_module.__default.fm14(globalState))) : (_3077_v0))));
      let _3097_v18;
      let _nw593 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _3097_v18 = _nw593;
      let _3098_v19;
      _3098_v19 = _dafny.Seq.of(new BigNumber(-904));
      let _index591 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_3097_v18).length));
      (_3097_v18)[_index591] = _dafny.Seq.Concat(_3098_v19, _3098_v19);
      let _index592 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_3097_v18).length));
      (_3097_v18)[_index592] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_3099_i2) {
        return (_this).f7;
      }), _dafny.Seq.Concat(_3098_v19, _3098_v19));
      let _3100_v20;
      _3100_v20 = _dafny.Seq.UnicodeFromString("wcjvago");
      r2 = _3100_v20;
      r0 = _3077_v0;
      r1 = _3077_v0;
      r2 = _dafny.Seq.UnicodeFromString("vkqgfndn");
      return [r0, r1, r2];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
