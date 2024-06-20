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
    static fm0(globalState) {
      return (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(421), new BigNumber(316))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(421)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(316)))) {
            _coll0.push([_module.__default.safeModuloInt(_0_v0, new BigNumber(-169)),false]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(296),!(false)));
    };
    static fm1(globalState) {
      return new _dafny.CodePoint('l'.codePointAt(0));
    };
    static fm2(globalState) {
      return !(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(573),new BigNumber((_dafny.Seq.of(!(!(true)), !(true), false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-447)),new BigNumber(235)))).equals((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true)).length),new BigNumber(271))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-452)), function (_1_i0) {
        return new BigNumber(321);
      })).length),new BigNumber(4)))));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(new BigNumber(369))).Elements) {
          let _2_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(369)), _2_v0)) {
            _coll1.add(_module.__default.safeDivisionInt(_2_v0, new BigNumber(526)));
          }
        }
        return _coll1;
      }()).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(261))).length))),false);
    };
    static fm4(globalState) {
      return _dafny.MultiSet.fromElements(((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length)), new BigNumber(247));
    };
    static fm5(globalState) {
      return _module.D1.create_DC3((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, false, true, true, !(false)),new BigNumber(376))).length)).isLessThan(new BigNumber(-423)), true, _module.__default.safeDivisionInt(new BigNumber(117), new BigNumber((_dafny.Seq.UnicodeFromString("sebrnl")).length)), true);
    };
    static fm6(p0, globalState) {
      return (_dafny.ZERO).minus(new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),true))).length));
    };
    static fm7(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _dafny.Seq.UnicodeFromString("wfhne"));
    };
    static fm8(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(((_module.D3.create_DC9(new BigNumber(80))).dtor_cf17).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("gy")).length),new BigNumber((_dafny.Seq.of(false, true)).length))).length)));
    };
    static fm9(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(!(true))).Union((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(false)));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC2(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(972), new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), function (_3_i0) {
    return new _dafny.CodePoint('e'.codePointAt(0));
  }),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(211))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("xom")).length), new BigNumber(419))).length))).length))).Keys.Elements) {
    let _4_v0 = _compr_2;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), function (_3_i0) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    }),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(211))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("xom")).length), new BigNumber(419))).length))).length))).contains(_4_v0)) {
      _coll2.push([_4_v0,true]);
    }
  }
  return _coll2;
}()).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(785)), function (_5_i1) {
  return new _dafny.CodePoint('c'.codePointAt(0));
})).length), new BigNumber(-884)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(849))).length)))).length), new _dafny.CodePoint('r'.codePointAt(0)), (new BigNumber(-240)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(534)), function (_6_i2) {
  return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(593)), function (_7_i3) {
    return new _dafny.CodePoint('t'.codePointAt(0));
  })).length);
})).length)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_8_i4) {
    return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_9_i5) {
      return new BigNumber(352);
    })).length);
  })).Elements) {
    let _10_v1 = _compr_3;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), function (_8_i4) {
      return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_9_i5) {
        return new BigNumber(352);
      })).length);
    }), _10_v1)) {
      _coll3.push([_module.__default.safeDivisionInt(_10_v1, new BigNumber(967)),true]);
    }
  }
  return _coll3;
}()).length), new BigNumber(-539)), (_module.D1.create_DC2(new BigNumber(70), new _dafny.CodePoint('v'.codePointAt(0)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(515)), function (_11_i6) {
  return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-71), new BigNumber(213))).cardinality()));
})).length), _dafny.Seq.of(new BigNumber(66)), new _dafny.CodePoint('d'.codePointAt(0)))).dtor_cf7), new _dafny.CodePoint('m'.codePointAt(0)));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-935), new BigNumber(330))) {
          let _12_v0 = _compr_4;
          if (((new BigNumber(-935)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(330)))) {
            _coll4.push([(_12_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(!(false))).length)),(_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.fromElements(!(!(true))))]);
          }
        }
        return _coll4;
      }();
    };
    static m0(p0, globalState) {
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Set.Empty;
      let _13_v0;
      _13_v0 = false;
      let _14_v1;
      let _nw0 = new _module.C0();
      _nw0.__ctor();
      _14_v1 = _nw0;
      let _15_v2;
      _15_v2 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_13_v0);
      let _16_v3;
      _16_v3 = _dafny.Map.Empty.slice().updateUnsafe(_14_v1,_15_v2);
      r1 = ((new BigNumber((_dafny.Seq.of(_13_v0)).length)).multipliedBy(new BigNumber((_16_v3).length))).isLessThanOrEqualTo(p0);
      let _17_v4;
      _17_v4 = _dafny.Seq.UnicodeFromString("d");
      let _18_v5;
      _18_v5 = _module.D3.create_DC10(!(_13_v0), _17_v4, _13_v0);
      let _19_v6;
      _19_v6 = _module.D3.create_DC11(_18_v5);
      _19_v6 = _19_v6;
      let _20_v7;
      _20_v7 = _dafny.Set.fromElements(new BigNumber((_17_v4).length));
      let _21_i0;
      _21_i0 = _dafny.ZERO;
      L0: {
        while ((_20_v7).IsProperSubsetOf(_20_v7)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_21_i0)) {
              break L0;
            }
            _21_i0 = (_21_i0).plus(_dafny.ONE);
            let _22_v8;
            _22_v8 = new _dafny.CodePoint('n'.codePointAt(0));
            _22_v8 = _22_v8;
            if (!(!(_module.__default.fm2(globalState)))) {
              let _23_v9;
              _23_v9 = _dafny.Set.fromElements(_13_v0);
              let _24_v10;
              _24_v10 = _dafny.Seq.of(_23_v9);
              _24_v10 = _24_v10;
              let _25_v11;
              _25_v11 = _module.D3.create_DC10(_13_v0, _dafny.Seq.UnicodeFromString("yfugb"), _13_v0);
              let _26_v12;
              let _nw1 = Array((new BigNumber(19)).toNumber());
              _nw1[(_dafny.ZERO).toNumber()] = _13_v0;
              _nw1[(_dafny.ONE).toNumber()] = _13_v0;
              _nw1[(new BigNumber(2)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(3)).toNumber()] = true;
              _nw1[(new BigNumber(4)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(5)).toNumber()] = _module.__default.fm2(globalState);
              _nw1[(new BigNumber(6)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(7)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(8)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(9)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(10)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(11)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(12)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(13)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(14)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(15)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(16)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(17)).toNumber()] = _13_v0;
              _nw1[(new BigNumber(18)).toNumber()] = _13_v0;
              _26_v12 = _nw1;
              let _27_v13;
              _27_v13 = _dafny.Map.Empty.slice().updateUnsafe(_26_v12,_13_v0);
              let _pat_let_tv0 = _27_v13;
              let _pat_let_tv1 = _26_v12;
              let _pat_let_tv2 = _13_v0;
              let _pat_let_tv3 = _27_v13;
              let _pat_let_tv4 = _26_v12;
              let _pat_let_tv5 = _20_v7;
              let _pat_let_tv6 = _20_v7;
              _25_v11 = function (_pat_let0_0) {
                return function (_30_dt__update__tmp_h1) {
                  return function (_pat_let3_0) {
                    return function (_31_dt__update_hcf20_h1) {
                      return _module.D3.create_DC10((_30_dt__update__tmp_h1).dtor_cf18, (_30_dt__update__tmp_h1).dtor_cf19, _31_dt__update_hcf20_h1);
                    }(_pat_let3_0);
                  }((_pat_let_tv5).IsProperSubsetOf(_pat_let_tv6));
                }(_pat_let0_0);
              }(function (_pat_let1_0) {
                return function (_28_dt__update__tmp_h0) {
                  return function (_pat_let2_0) {
                    return function (_29_dt__update_hcf20_h0) {
                      return _module.D3.create_DC10((_28_dt__update__tmp_h0).dtor_cf18, (_28_dt__update__tmp_h0).dtor_cf19, _29_dt__update_hcf20_h0);
                    }(_pat_let2_0);
                  }(!((((_pat_let_tv3).contains(_pat_let_tv4)) ? ((_pat_let_tv0).get(_pat_let_tv1)) : (_pat_let_tv2))));
                }(_pat_let1_0);
              }(_25_v11));
              _13_v0 = _13_v0;
              _14_v1 = _14_v1;
              let _32_v14;
              _32_v14 = _module.D5.create_DC14(_13_v0, _13_v0, _14_v1, new BigNumber(151));
              let _pat_let_tv7 = _15_v2;
              let _pat_let_tv8 = _13_v0;
              let _pat_let_tv9 = _13_v0;
              let _pat_let_tv10 = _15_v2;
              let _pat_let_tv11 = _13_v0;
              let _33_v15;
              _33_v15 = _dafny.Seq.of((function (_pat_let4_0) {
                return function (_34_dt__update__tmp_h2) {
                  return function (_pat_let5_0) {
                    return function (_35_dt__update_hcf25_h0) {
                      return _module.D5.create_DC14((_34_dt__update__tmp_h2).dtor_cf24, _35_dt__update_hcf25_h0, (_34_dt__update__tmp_h2).dtor_cf26, (_34_dt__update__tmp_h2).dtor_cf27);
                    }(_pat_let5_0);
                  }((((_pat_let_tv10).contains(_pat_let_tv11)) ? ((_pat_let_tv7).get(_pat_let_tv8)) : (_pat_let_tv9)));
                }(_pat_let4_0);
              }(_32_v14)).dtor_cf26);
              _14_v1 = (_33_v15)[_module.__default.safeIndex(p0, new BigNumber((_33_v15).length))];
            } else {
              let _36_v16;
              let _nw2 = Array((new BigNumber(6)).toNumber()).fill(false);
              _36_v16 = _nw2;
              let _index0 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length));
              (_36_v16)[_index0] = (p0).isLessThan(p0);
              let _37_v17;
              _37_v17 = _dafny.MultiSet.fromElements(!(_13_v0), _13_v0);
              let _38_v18;
              _38_v18 = _dafny.Set.fromElements(_13_v0, _13_v0);
              let _39_v19;
              let _nw3 = Array((new BigNumber(5)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = (p0).minus(p0);
              _nw3[(_dafny.ONE).toNumber()] = new BigNumber((_37_v17).cardinality());
              _nw3[(new BigNumber(2)).toNumber()] = p0;
              _nw3[(new BigNumber(3)).toNumber()] = p0;
              _nw3[(new BigNumber(4)).toNumber()] = ((false) ? ((_dafny.ZERO).minus(new BigNumber((_38_v18).length))) : (p0));
              _39_v19 = _nw3;
              let _index1 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length));
              (_39_v19)[_index1] = p0;
              let _40_v20;
              let _nw4 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _40_v20 = _nw4;
              let _41_v21;
              _41_v21 = _dafny.Map.Empty.slice().updateUnsafe(false,_40_v20);
              let _42_v22;
              _42_v22 = _41_v21;
              let _43_v23;
              _43_v23 = _dafny.MultiSet.fromElements(p0);
              let _index2 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_39_v19).length));
              (_39_v19)[_index2] = (new BigNumber(((_42_v22)).length)).multipliedBy(new BigNumber((_43_v23).cardinality()));
              let _index3 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length));
              let _index4 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length));
              let _index5 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_39_v19).length));
              let _rhs0 = _dafny.Seq.IsPrefixOf(_module.__default.fm7(p0, globalState), _17_v4);
              let _rhs1 = p0;
              let _rhs2 = _module.__default.safeDivisionInt(new BigNumber(-370), p0);
              let _rhs3 = (_dafny.ZERO).minus(p0);
              let _lhs0 = _36_v16;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length));
              let _lhs2 = _39_v19;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length));
              let _lhs4 = globalState;
              let _lhs5 = _39_v19;
              let _lhs6 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_39_v19).length));
              _lhs0[_lhs1] = _rhs0;
              _lhs2[_lhs3] = _rhs1;
              _lhs4.f2 = _rhs2;
              _lhs5[_lhs6] = _rhs3;
              let _44_v24;
              let _init0 = ((_45_v8) => function (_46_i1) {
                return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_47_v8) => function (_48_i2) {
                  return _47_v8;
                })(_45_v8));
              })(_22_v8);
              let _nw5 = Array((new BigNumber(22)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
                _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _44_v24 = _nw5;
              let _rhs4 = _44_v24;
              let _rhs5 = (_36_v16)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length))];
              let _rhs6 = (_43_v23).Difference(_43_v23);
              let _rhs7 = _module.__default.fm7((((_36_v16)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length))]) ? ((_39_v19)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length))]) : ((_39_v19)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length))])), globalState);
              let _lhs7 = globalState;
              _44_v24 = _rhs4;
              r1 = _rhs5;
              _43_v23 = _rhs6;
              _lhs7.f11 = _rhs7;
              let _49_v25;
              _49_v25 = _dafny.Seq.of((_39_v19)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length))]);
              let _50_v26;
              _50_v26 = _dafny.Seq.of(p0, (_39_v19)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length))], new BigNumber((_49_v25).length));
              let _51_v27;
              _51_v27 = _dafny.Seq.of(_50_v26, _49_v25);
              let _52_v28;
              _52_v28 = _dafny.MultiSet.fromElements(_50_v26, (_51_v27)[_module.__default.safeIndex((_39_v19)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_39_v19).length))], new BigNumber((_51_v27).length))]);
              let _53_v29;
              _53_v29 = _dafny.Map.Empty.slice().updateUnsafe(_17_v4,_52_v28);
              _53_v29 = (_53_v29).update(_17_v4, _52_v28);
              let _rhs8 = _14_v1;
              let _rhs9 = (_36_v16)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length))];
              let _rhs10 = (((_36_v16)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length))]) ? (!(_13_v0)) : ((_36_v16)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_36_v16).length))]));
              _14_v1 = _rhs8;
              r1 = _rhs9;
              _13_v0 = _rhs10;
              let _54_v30;
              let _nw6 = new _module.C0();
              _nw6.__ctor();
              _54_v30 = _nw6;
            }
            let _55_v32;
            _55_v32 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(p0));
            let _56_v33;
            _56_v33 = _module.D1.create_DC2(new BigNumber(214), _22_v8, new BigNumber((function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of (_20_v7).Elements) {
    let _57_v31 = _compr_5;
    if ((_20_v7).contains(_57_v31)) {
      _coll5.push([(_57_v31).plus(p0),(((_15_v2).contains(true)) ? ((_15_v2).get(true)) : (_13_v0))]);
    }
  }
  return _coll5;
}()).length), _55_v32, _22_v8);
            let _source0 = _module.D1.create_DC4(_56_v33);
            if (_source0.is_DC2) {
              let _58___mcc_h0 = (_source0).cf4;
              let _59___mcc_h1 = (_source0).cf5;
              let _60___mcc_h2 = (_source0).cf6;
              let _61___mcc_h3 = (_source0).cf7;
              let _62___mcc_h4 = (_source0).cf8;
              let _63_cf8 = _62___mcc_h4;
              let _64_cf7 = _61___mcc_h3;
              let _65_cf6 = _60___mcc_h2;
              let _66_cf5 = _59___mcc_h1;
              let _67_cf4 = _58___mcc_h0;
              let _68_v34;
              let _init1 = ((_69_cf4) => function (_70_i3) {
                return (_70_i3).plus(_69_cf4);
              })(_67_cf4);
              let _nw7 = Array((new BigNumber(5)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
                _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _68_v34 = _nw7;
              let _index6 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_68_v34).length));
              (_68_v34)[_index6] = _67_cf4;
              let _71_v35;
              _71_v35 = _module.D3.create_DC9(_67_cf4);
              let _index7 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_68_v34).length));
              (_68_v34)[_index7] = (_71_v35).dtor_cf17;
              let _72_v36;
              let _nw8 = new _module.C0();
              _nw8.__ctor();
              _72_v36 = _nw8;
              let _73_v37;
              _73_v37 = _dafny.Map.Empty.slice().updateUnsafe(_67_cf4,_65_cf6);
              _73_v37 = (_73_v37).update(_65_cf6, (_65_cf6).multipliedBy(_67_cf4));
              let _74_v38;
              let _nw9 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
              _74_v38 = _nw9;
              let _75_v39;
              let _nw10 = Array((new BigNumber(6)).toNumber());
              _nw10[(_dafny.ZERO).toNumber()] = _module.__default.fm2(globalState);
              _nw10[(_dafny.ONE).toNumber()] = _13_v0;
              _nw10[(new BigNumber(2)).toNumber()] = _13_v0;
              _nw10[(new BigNumber(3)).toNumber()] = _13_v0;
              _nw10[(new BigNumber(4)).toNumber()] = _13_v0;
              _nw10[(new BigNumber(5)).toNumber()] = _13_v0;
              _75_v39 = _nw10;
              let _76_v40;
              _76_v40 = _dafny.Map.Empty.slice().updateUnsafe(_75_v39,_72_v36);
              let _index8 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_74_v38).length));
              (_74_v38)[_index8] = _76_v40;
              let _77_v41;
              _77_v41 = _dafny.Seq.of(_13_v0, _13_v0);
              let _78_v42;
              _78_v42 = _module.D2.create_DC5(_77_v41);
              let _79_v43;
              _79_v43 = _dafny.Map.Empty.slice().updateUnsafe(_78_v42,_64_cf7);
              let _80_v44;
              _80_v44 = _dafny.Map.Empty.slice().updateUnsafe(_72_v36,(_65_cf6).multipliedBy(new BigNumber(((((_79_v43).contains(_78_v42)) ? ((_79_v43).get(_78_v42)) : (_64_cf7))).length)));
              let _index9 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_74_v38).length));
              let _index10 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_68_v34).length));
              let _rhs11 = _14_v1;
              let _rhs12 = _76_v40;
              let _rhs13 = (new BigNumber(((_20_v7).Union(_20_v7)).length)).minus((_dafny.ZERO).minus(p0));
              let _rhs14 = _67_cf4;
              let _rhs15 = _80_v44;
              let _lhs8 = _74_v38;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_74_v38).length));
              let _lhs10 = globalState;
              let _lhs11 = _68_v34;
              let _lhs12 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_68_v34).length));
              _72_v36 = _rhs11;
              _lhs8[_lhs9] = _rhs12;
              _lhs10.f2 = _rhs13;
              _lhs11[_lhs12] = _rhs14;
              _80_v44 = _rhs15;
            } else if (_source0.is_DC3) {
              let _81___mcc_h5 = (_source0).cf9;
              let _82___mcc_h6 = (_source0).cf10;
              let _83___mcc_h7 = (_source0).cf11;
              let _84___mcc_h8 = (_source0).cf12;
              let _85_cf12 = _84___mcc_h8;
              let _86_cf11 = _83___mcc_h7;
              let _87_cf10 = _82___mcc_h6;
              let _88_cf9 = _81___mcc_h5;
              r1 = _88_cf9;
              let _89_v45;
              let _nw11 = new _module.C0();
              _nw11.__ctor();
              _89_v45 = _nw11;
              let _90_v47;
              let _init2 = ((_91_cf10, _92_p0, _93_cf9) => function (_94_i4) {
                return ((!(_91_cf10)) ? (_dafny.Map.Empty.slice().updateUnsafe(_92_p0,_dafny.MultiSet.fromElements(_93_cf9, _91_cf10))) : (function () {
                  let _coll6 = new _dafny.Map();
                  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-475), new BigNumber(301))) {
                    let _95_v46 = _compr_6;
                    if (((new BigNumber(-475)).isLessThanOrEqualTo(_95_v46)) && ((_95_v46).isLessThan(new BigNumber(301)))) {
                      _coll6.push([(_95_v46).multipliedBy(_92_p0),_dafny.MultiSet.fromElements(_91_cf10)]);
                    }
                  }
                  return _coll6;
                }()));
              })(_87_cf10, p0, _88_cf9);
              let _nw12 = Array((new BigNumber(19)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw12.length); _i0_2++) {
                _nw12[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _90_v47 = _nw12;
              let _96_v49;
              _96_v49 = _dafny.Set.fromElements(_85_cf12, _87_cf10, _13_v0, _88_cf9, _87_cf10);
              let _97_v50;
              _97_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_17_v4).length),_96_v49);
              let _index11 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_90_v47).length));
              (_90_v47)[_index11] = function () {
                let _coll7 = new _dafny.Map();
                for (const _compr_7 of (_97_v50).Keys.Elements) {
                  let _98_v48 = _compr_7;
                  if ((_97_v50).contains(_98_v48)) {
                    _coll7.push([(_98_v48).multipliedBy(new BigNumber((_dafny.Seq.of(_87_cf10)).length)),_dafny.MultiSet.fromElements(_87_cf10, !(_85_cf12), _88_cf9)]);
                  }
                }
                return _coll7;
              }();
              let _99_v51;
              _99_v51 = _dafny.Seq.of(_85_cf12);
              let _100_v52;
              _100_v52 = _dafny.MultiSet.fromElements(_86_cf11);
              let _101_v53;
              let _nw13 = Array((new BigNumber(11)).toNumber());
              _nw13[(_dafny.ZERO).toNumber()] = p0;
              _nw13[(_dafny.ONE).toNumber()] = p0;
              _nw13[(new BigNumber(2)).toNumber()] = _86_cf11;
              _nw13[(new BigNumber(3)).toNumber()] = p0;
              _nw13[(new BigNumber(4)).toNumber()] = new BigNumber(221);
              _nw13[(new BigNumber(5)).toNumber()] = _86_cf11;
              _nw13[(new BigNumber(6)).toNumber()] = _86_cf11;
              _nw13[(new BigNumber(7)).toNumber()] = new BigNumber((_17_v4).length);
              _nw13[(new BigNumber(8)).toNumber()] = _86_cf11;
              _nw13[(new BigNumber(9)).toNumber()] = _86_cf11;
              _nw13[(new BigNumber(10)).toNumber()] = _86_cf11;
              _101_v53 = _nw13;
              let _102_v54;
              _102_v54 = _101_v53;
              let _103_v55;
              _103_v55 = _dafny.Map.Empty.slice().updateUnsafe(_102_v54,_13_v0);
              let _104_v56;
              _104_v56 = _module.D2.create_DC7(_module.D2.create_DC7(_module.D2.create_DC6()));
              let _105_v57;
              _105_v57 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_104_v56);
              let _106_v58;
              let _nw14 = Array((new BigNumber(23)).toNumber());
              _nw14[(_dafny.ZERO).toNumber()] = new BigNumber(858);
              _nw14[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dnsnwbotg"), _dafny.Seq.UnicodeFromString("xvnoa"))).length);
              _nw14[(new BigNumber(2)).toNumber()] = ((_85_cf12) ? (_86_cf11) : (_86_cf11));
              _nw14[(new BigNumber(3)).toNumber()] = p0;
              _nw14[(new BigNumber(4)).toNumber()] = (new BigNumber(975)).plus(new BigNumber(655));
              _nw14[(new BigNumber(5)).toNumber()] = p0;
              _nw14[(new BigNumber(6)).toNumber()] = _86_cf11;
              _nw14[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_17_v4).length), p0);
              _nw14[(new BigNumber(8)).toNumber()] = new BigNumber((_100_v52).cardinality());
              _nw14[(new BigNumber(9)).toNumber()] = _86_cf11;
              _nw14[(new BigNumber(10)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("j")).length)).multipliedBy(p0);
              _nw14[(new BigNumber(11)).toNumber()] = new BigNumber((_20_v7).length);
              _nw14[(new BigNumber(12)).toNumber()] = p0;
              _nw14[(new BigNumber(13)).toNumber()] = new BigNumber((_103_v55).length);
              _nw14[(new BigNumber(14)).toNumber()] = p0;
              _nw14[(new BigNumber(15)).toNumber()] = (_86_cf11).multipliedBy(p0);
              _nw14[(new BigNumber(16)).toNumber()] = ((_module.__default.fm2(globalState)) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-769)), ((_107_v8) => function (_108_i5) {
                return _107_v8;
              })(_22_v8))).length)) : (new BigNumber((_105_v57).length)));
              _nw14[(new BigNumber(17)).toNumber()] = _86_cf11;
              _nw14[(new BigNumber(18)).toNumber()] = _86_cf11;
              _nw14[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), new BigNumber((_17_v4).length));
              _nw14[(new BigNumber(20)).toNumber()] = _86_cf11;
              _nw14[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("pjqydcgsg")).length);
              _nw14[(new BigNumber(22)).toNumber()] = _86_cf11;
              _106_v58 = _nw14;
              let _index12 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_106_v58).length));
              (_106_v58)[_index12] = p0;
              let _109_v59;
              _109_v59 = _module.D3.create_DC9(p0);
              let _index13 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_90_v47).length));
              let _index14 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_106_v58).length));
              let _rhs16 = _module.__default.fm11(!(_87_cf10) || (_13_v0), new BigNumber(749), _109_v59, !_dafny.Seq.contains(_17_v4, _22_v8), globalState);
              let _rhs17 = _99_v51;
              let _rhs18 = _module.__default.safeModuloInt(_86_cf11, _86_cf11);
              let _lhs13 = _90_v47;
              let _lhs14 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_90_v47).length));
              let _lhs15 = _106_v58;
              let _lhs16 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_106_v58).length));
              _lhs13[_lhs14] = _rhs16;
              _99_v51 = _rhs17;
              _lhs15[_lhs16] = _rhs18;
              _88_cf9 = _87_cf10;
            } else if (_source0.is_DC1) {
              let _110___mcc_h9 = (_source0).cf3;
              let _111_cf3 = _110___mcc_h9;
              _15_v2 = (_15_v2).update(_module.__default.fm2(globalState), _111_cf3);
              let _112_v60;
              let _nw15 = new _module.C0();
              _nw15.__ctor();
              _112_v60 = _nw15;
              let _113_v61;
              let _nw16 = Array((new BigNumber(4)).toNumber());
              _nw16[(_dafny.ZERO).toNumber()] = (_15_v2).Merge(_15_v2);
              _nw16[(_dafny.ONE).toNumber()] = _15_v2;
              _nw16[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_111_cf3);
              _nw16[(new BigNumber(3)).toNumber()] = _15_v2;
              _113_v61 = _nw16;
              _113_v61 = _113_v61;
              let _114_v62;
              let _init3 = ((_115_p0) => function (_116_i6) {
                return (_116_i6).plus(_115_p0);
              })(p0);
              let _nw17 = Array((new BigNumber(6)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw17.length); _i0_3++) {
                _nw17[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _114_v62 = _nw17;
              _114_v62 = _114_v62;
            } else {
              let _117___mcc_h10 = (_source0).cf13;
              let _118_cf13 = _117___mcc_h10;
              (globalState).f2 = p0;
              let _119_v63;
              let _nw18 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _119_v63 = _nw18;
              let _index15 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_119_v63).length));
              (_119_v63)[_index15] = _dafny.Seq.UnicodeFromString("lbrfdtt");
              let _index16 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_119_v63).length));
              (_119_v63)[_index16] = _dafny.Seq.UnicodeFromString("rkrne");
              let _120_v64;
              let _nw19 = new _module.C0();
              _nw19.__ctor();
              _120_v64 = _nw19;
              let _rhs19 = _55_v32;
              let _rhs20 = _13_v0;
              _55_v32 = _rhs19;
              _13_v0 = _rhs20;
            }
            _22_v8 = _22_v8;
          }
        }
      }
      let _121_v65;
      let _init4 = ((_122_v4) => function (_123_i8) {
        return _dafny.Seq.IsProperPrefixOf((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), function (_124_i9) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }), _122_v4)), _dafny.Seq.of(_122_v4));
      })(_17_v4);
      let _nw20 = Array((new BigNumber(17)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw20.length); _i0_4++) {
        _nw20[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _121_v65 = _nw20;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_121_v65).length))) {
        let _125_i7 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_125_i7)) && ((_125_i7).isLessThan(new BigNumber((_121_v65).length))))) {
          (_121_v65)[(_125_i7)] = (new BigNumber(613)).isLessThan((new BigNumber((_15_v2).length)).minus(p0));
        }
      }
      let _126_v66;
      let _nw21 = Array((new BigNumber(6)).toNumber());
      _126_v66 = _nw21;
      let _index17 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_126_v66).length));
      (_126_v66)[_index17] = _14_v1;
      let _index18 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_126_v66).length));
      (_126_v66)[_index18] = _14_v1;
      let _127_v67;
      let _nw22 = Array((new BigNumber(4)).toNumber());
      _nw22[(_dafny.ZERO).toNumber()] = _121_v65;
      _nw22[(_dafny.ONE).toNumber()] = _121_v65;
      _nw22[(new BigNumber(2)).toNumber()] = _121_v65;
      _nw22[(new BigNumber(3)).toNumber()] = _121_v65;
      _127_v67 = _nw22;
      let _index19 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_127_v67).length));
      (_127_v67)[_index19] = _121_v65;
      let _128_v68;
      _128_v68 = _module.D3.create_DC8(_121_v65);
      let _index20 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_127_v67).length));
      (_127_v67)[_index20] = (_128_v68).dtor_cf16;
      let _129_v70;
      _129_v70 = _module.D1.create_DC3(_13_v0, _13_v0, p0, _13_v0);
      let _130_v71;
      _130_v71 = _dafny.MultiSet.fromElements(_129_v70);
      let _131_v72;
      _131_v72 = _dafny.Seq.of(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_130_v71).Elements) {
          let _132_v69 = _compr_8;
          if ((_130_v71).contains(_132_v69)) {
            _coll8.push([_132_v69,_13_v0]);
          }
        }
        return _coll8;
      }(), _dafny.Map.Empty.slice().updateUnsafe(_129_v70,true));
      r0 = _131_v72;
      r1 = !(_13_v0);
      r2 = _13_v0;
      let _133_v73;
      _133_v73 = _dafny.Set.fromElements(((_13_v0) ? (_13_v0) : (_13_v0)));
      r3 = _133_v73;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _134_v0;
      let _nw23 = Array((new BigNumber(19)).toNumber()).fill(false);
      _134_v0 = _nw23;
      let _135_v1;
      let _init5 = function (_136_i0) {
        return _dafny.Set.fromElements(new BigNumber(-34));
      };
      let _nw24 = Array((new BigNumber(21)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw24.length); _i0_5++) {
        _nw24[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _135_v1 = _nw24;
      let _137_globalState;
      let _nw25 = new _module.GlobalState();
      _nw25.__ctor(false, _134_v0, new BigNumber(212), false, new BigNumber(3), true, false, new BigNumber(743), true, _135_v1, new BigNumber(700), _dafny.Seq.Create(_module.__default.abs(new BigNumber(199)), function (_138_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }));
      _137_globalState = _nw25;
      let _139_v2;
      _139_v2 = true;
      if (_139_v2) {
        let _140_v3;
        _140_v3 = new BigNumber(249);
        let _141_v4;
        _141_v4 = _dafny.Map.Empty.slice().updateUnsafe(_140_v3,_139_v2);
        let _142_v5;
        _142_v5 = _dafny.Seq.of((_module.__default.fm0(_137_globalState)).Merge(_141_v4));
        _142_v5 = ((_139_v2) ? (_142_v5) : (_142_v5));
        let _143_v6;
        _143_v6 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(_140_v3));
        let _144_v7;
        _144_v7 = _dafny.Seq.UnicodeFromString("xbl");
        let _145_v8;
        _145_v8 = _dafny.MultiSet.fromElements(new BigNumber(-795), new BigNumber((_144_v7).length), _140_v3, _140_v3, _140_v3);
        let _146_v9;
        _146_v9 = _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_143_v6).length)), _145_v8, _140_v3);
        let _source1 = _146_v9;
        let _147___mcc_h0 = (_source1).cf0;
        let _148___mcc_h1 = (_source1).cf1;
        let _149___mcc_h2 = (_source1).cf2;
        let _150_cf2 = _149___mcc_h2;
        let _151_cf1 = _148___mcc_h1;
        let _152_cf0 = _147___mcc_h0;
        let _153_v10;
        _153_v10 = new _dafny.CodePoint('g'.codePointAt(0));
        let _154_v11;
        let _nw26 = Array((new BigNumber(9)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _153_v10;
        _nw26[(_dafny.ONE).toNumber()] = _153_v10;
        _nw26[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_137_globalState);
        _nw26[(new BigNumber(3)).toNumber()] = _153_v10;
        _nw26[(new BigNumber(4)).toNumber()] = _153_v10;
        _nw26[(new BigNumber(5)).toNumber()] = _153_v10;
        _nw26[(new BigNumber(6)).toNumber()] = ((_139_v2) ? (_153_v10) : (_153_v10));
        _nw26[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_137_globalState);
        _nw26[(new BigNumber(8)).toNumber()] = _153_v10;
        _154_v11 = _nw26;
        let _index21 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_154_v11).length));
        (_154_v11)[_index21] = _153_v10;
        let _index22 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_154_v11).length));
        (_154_v11)[_index22] = _153_v10;
        let _155_v12;
        _155_v12 = _dafny.Map.Empty.slice().updateUnsafe(false,false);
        let _156_v13;
        _156_v13 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_155_v12).length)));
        let _157_v14;
        let _nw27 = Array((new BigNumber(23)).toNumber()).fill([]);
        _157_v14 = _nw27;
        let _158_v15;
        _158_v15 = _dafny.Set.fromElements(_139_v2, _139_v2);
        let _159_v16;
        let _nw28 = Array((_dafny.ONE).toNumber());
        _nw28[(_dafny.ZERO).toNumber()] = _158_v15;
        _159_v16 = _nw28;
        let _160_v17;
        _160_v17 = _dafny.Map.Empty.slice().updateUnsafe(_159_v16,_157_v14);
        let _161_v18;
        _161_v18 = _dafny.Seq.of(_159_v16);
        let _162_v19;
        _162_v19 = _module.D1.create_DC1(_139_v2);
        let _rhs21 = _156_v13;
        let _rhs22 = (((_160_v17).contains((_161_v18)[_module.__default.safeIndex(_150_cf2, new BigNumber((_161_v18).length))])) ? ((_160_v17).get((_161_v18)[_module.__default.safeIndex(_150_cf2, new BigNumber((_161_v18).length))])) : (_157_v14));
        let _rhs23 = _139_v2;
        let _rhs24 = (((_150_cf2).isEqualTo(_150_cf2)) ? (_139_v2) : ((_162_v19).dtor_cf3));
        _156_v13 = _rhs21;
        _157_v14 = _rhs22;
        _139_v2 = _rhs23;
        _139_v2 = _rhs24;
        _140_v3 = _150_cf2;
        let _163_v20;
        _163_v20 = _dafny.Seq.of(_139_v2);
        let _164_v21;
        _164_v21 = _dafny.Map.Empty.slice().updateUnsafe(_153_v10,_139_v2);
        let _165_v22;
        _165_v22 = _dafny.Seq.of((_163_v20)[_module.__default.safeIndex(_140_v3, new BigNumber((_163_v20).length))], false, (((_164_v21).contains((_154_v11)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_154_v11).length))])) ? ((_164_v21).get((_154_v11)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_154_v11).length))])) : (_139_v2)), !(_139_v2), !(_139_v2));
        let _166_v23;
        let _nw29 = Array((new BigNumber(21)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = _151_cf1;
        _nw29[(_dafny.ONE).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(2)).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(3)).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_156_v13);
        _nw29[(new BigNumber(5)).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(6)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(7)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_156_v13);
        _nw29[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_150_cf2, _152_cf0);
        _nw29[(new BigNumber(10)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(11)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(12)).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(13)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(14)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(15)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(_150_cf2);
        _nw29[(new BigNumber(17)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(18)).toNumber()] = _145_v8;
        _nw29[(new BigNumber(19)).toNumber()] = _151_cf1;
        _nw29[(new BigNumber(20)).toNumber()] = _145_v8;
        _166_v23 = _nw29;
        let _167_v24;
        _167_v24 = _dafny.Map.Empty.slice().updateUnsafe(_165_v22,_166_v23);
        _167_v24 = (_167_v24).update(_163_v20, _166_v23);
        _141_v4 = (_141_v4).update(new BigNumber((_143_v6).length), (!(_139_v2)) && (_139_v2));
        let _168_v25;
        let _nw30 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _168_v25 = _nw30;
        let _169_v26;
        _169_v26 = _dafny.Map.Empty.slice().updateUnsafe(_168_v25,((_139_v2) ? (_139_v2) : (_139_v2)));
        let _170_v29;
        _170_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_144_v7).length),_140_v3);
        let _171_v30;
        _171_v30 = _dafny.MultiSet.fromElements(function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of (_145_v8).Elements) {
            let _172_v27 = _compr_9;
            if ((_145_v8).contains(_172_v27)) {
              _coll9.push([(_172_v27).minus(_140_v3),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(((_module.D2.create_DC5(_dafny.Seq.of(_139_v2, _139_v2))).dtor_cf14).length))).length))).length)]);
            }
          }
          return _coll9;
        }(), function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_141_v4).Keys.Elements) {
            let _173_v28 = _compr_10;
            if ((_141_v4).contains(_173_v28)) {
              _coll10.push([_module.__default.safeDivisionInt(_173_v28, _140_v3),_140_v3]);
            }
          }
          return _coll10;
        }(), _170_v29);
        _169_v26 = (_169_v26).update(_168_v25, (_171_v30).IsDisjointFrom(_171_v30));
        let _174_v31;
        _174_v31 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.of(new BigNumber(725)), _dafny.Seq.of(_140_v3, _140_v3)),(_170_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(_140_v3,_140_v3)));
        _170_v29 = (((_174_v31).contains(_139_v2)) ? ((_174_v31).get(_139_v2)) : (_170_v29));
      } else {
        let _175_v32;
        _175_v32 = new _dafny.CodePoint('i'.codePointAt(0));
        _175_v32 = _175_v32;
        _139_v2 = _139_v2;
        let _176_v33;
        _176_v33 = new BigNumber(880);
        let _177_v34;
        _177_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_176_v33).isEqualTo(_176_v33),new BigNumber(636));
        _177_v34 = (_177_v34).update(_139_v2, _176_v33);
        _175_v32 = new _dafny.CodePoint('n'.codePointAt(0));
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(535)), ((_178_v32) => function (_179_i2) {
          return _178_v32;
        })(_175_v32)), _dafny.Seq.UnicodeFromString("jrvkbowy"))) {
          let _180_v35;
          _180_v35 = _dafny.MultiSet.fromElements(_175_v32, _175_v32);
          let _181_v36;
          _181_v36 = _dafny.MultiSet.fromElements(_139_v2, _module.__default.fm2(_137_globalState));
          let _182_v37;
          _182_v37 = _dafny.Map.Empty.slice().updateUnsafe(_176_v33,_176_v33);
          let _183_v38;
          let _nw31 = Array((new BigNumber(10)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = new BigNumber((_180_v35).cardinality());
          _nw31[(_dafny.ONE).toNumber()] = _176_v33;
          _nw31[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(903)), ((_184_v33) => function (_185_i3) {
            return (_dafny.ZERO).minus(_184_v33);
          })(_176_v33))).length);
          _nw31[(new BigNumber(3)).toNumber()] = _176_v33;
          _nw31[(new BigNumber(4)).toNumber()] = _176_v33;
          _nw31[(new BigNumber(5)).toNumber()] = _176_v33;
          _nw31[(new BigNumber(6)).toNumber()] = _176_v33;
          _nw31[(new BigNumber(7)).toNumber()] = _176_v33;
          _nw31[(new BigNumber(8)).toNumber()] = new BigNumber((_181_v36).cardinality());
          _nw31[(new BigNumber(9)).toNumber()] = new BigNumber((_182_v37).length);
          _183_v38 = _nw31;
          let _186_v39;
          _186_v39 = _dafny.Set.fromElements(_183_v38, _183_v38);
          let _rhs25 = new BigNumber((_186_v39).length);
          let _rhs26 = _module.__default.fm2(_137_globalState);
          let _rhs27 = _module.__default.fm2(_137_globalState);
          let _rhs28 = _139_v2;
          let _rhs29 = _176_v33;
          let _lhs17 = _137_globalState;
          let _lhs18 = _137_globalState;
          _lhs17.f2 = _rhs25;
          _139_v2 = _rhs26;
          _139_v2 = _rhs27;
          _139_v2 = _rhs28;
          _lhs18.f2 = _rhs29;
          let _187_v40;
          let _nw32 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _187_v40 = _nw32;
          let _index23 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_187_v40).length));
          (_187_v40)[_index23] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(266)), ((_188_v32) => function (_189_i4) {
            return _188_v32;
          })(_175_v32));
          let _index24 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_187_v40).length));
          (_187_v40)[_index24] = _dafny.Seq.UnicodeFromString("yjcvetkn");
          let _190_v41;
          let _nw33 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _190_v41 = _nw33;
          let _191_v42;
          _191_v42 = _dafny.Seq.of(_dafny.Seq.of(_176_v33, _176_v33));
          let _index25 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_190_v41).length));
          (_190_v41)[_index25] = _191_v42;
          let _index26 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_190_v41).length));
          (_190_v41)[_index26] = _dafny.Seq.Concat(_dafny.Seq.Concat(_191_v42, _191_v42), _dafny.Seq.Concat(_191_v42, _191_v42));
          let _index27 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_183_v38).length));
          (_183_v38)[_index27] = _176_v33;
          let _index28 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_183_v38).length));
          (_183_v38)[_index28] = new BigNumber(59);
          let _192_v43;
          _192_v43 = _dafny.Seq.of(new BigNumber(-970));
          (_137_globalState).f2 = (((_181_v36).contains(_dafny.Seq.IsProperPrefixOf(_192_v43, _192_v43))) ? ((_181_v36).get(_dafny.Seq.IsProperPrefixOf(_192_v43, _192_v43))) : ((_dafny.ZERO).minus(_176_v33)));
        } else {
          let _193_v44;
          _193_v44 = _dafny.MultiSet.fromElements(!(_module.__default.fm2(_137_globalState)));
          (_137_globalState).f2 = new BigNumber((_193_v44).cardinality());
          let _194_v45;
          _194_v45 = _dafny.Map.Empty.slice().updateUnsafe(_139_v2,_139_v2);
          _139_v2 = (new BigNumber((_194_v45).length)).isEqualTo(_176_v33);
          _194_v45 = (_194_v45).update(_139_v2, !(_139_v2) || (_139_v2));
          let _195_v46;
          let _nw34 = new _module.C0();
          _nw34.__ctor();
          _195_v46 = _nw34;
          let _rhs30 = _139_v2;
          let _rhs31 = (_176_v33).multipliedBy(_176_v33);
          let _rhs32 = _module.__default.fm3(_139_v2, _dafny.MultiSet.fromElements(false, _139_v2, _139_v2), _139_v2, _176_v33, _137_globalState);
          let _rhs33 = new BigNumber((((_194_v45).Merge(_194_v45)).Merge(_194_v45)).length);
          let _lhs19 = _137_globalState;
          _139_v2 = _rhs30;
          _lhs19.f2 = _rhs31;
          _194_v45 = _rhs32;
          _176_v33 = _rhs33;
        }
      }
      let _index29 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
      (_134_v0)[_index29] = _139_v2;
      let _index30 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
      (_134_v0)[_index30] = _139_v2;
      let _196_v47;
      let _nw35 = Array((new BigNumber(2)).toNumber()).fill(_module.D1.Default());
      _196_v47 = _nw35;
      let _ingredients0 = [];
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_196_v47).length))) {
        let _197_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_197_i5)) && ((_197_i5).isLessThan(new BigNumber((_196_v47).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_196_v47, (_197_i5).toNumber(), ((_139_v2) ? (_module.D1.create_DC4(_module.D1.create_DC1((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]))) : (_module.D1.create_DC4(_module.D1.create_DC1(_139_v2))))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      if (!((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])) {
        let _198_v48;
        _198_v48 = new BigNumber(890);
        let _199_v49;
        _199_v49 = _dafny.Map.Empty.slice().updateUnsafe(_198_v48,_module.__default.fm2(_137_globalState));
        let _200_v50;
        let _201_v51;
        let _202_v52;
        let _203_v53;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(new BigNumber((_199_v49).length), _137_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _200_v50 = _out0;
        _201_v51 = _out1;
        _202_v52 = _out2;
        _203_v53 = _out3;
        (_137_globalState).f2 = _198_v48;
        let _204_v54;
        let _nw36 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _204_v54 = _nw36;
        let _index31 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_204_v54).length));
        (_204_v54)[_index31] = _198_v48;
        let _205_v55;
        _205_v55 = _dafny.Set.fromElements(_198_v48, _198_v48);
        let _index32 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_204_v54).length));
        (_204_v54)[_index32] = (new BigNumber((_205_v55).length)).plus(_198_v48);
        _201_v51 = ((_dafny.Set.fromElements(_201_v51, true)).Union(_203_v53)).IsProperSubsetOf((_203_v53).Difference(_203_v53));
        let _206_v56;
        let _nw37 = new _module.C0();
        _nw37.__ctor();
        _206_v56 = _nw37;
      } else {
        let _207_v57;
        _207_v57 = new BigNumber(970);
        let _208_v58;
        _208_v58 = _dafny.MultiSet.fromElements(_207_v57, _207_v57);
        _208_v58 = (_module.__default.fm4(_137_globalState)).Intersect(_208_v58);
        let _index33 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        (_134_v0)[_index33] = (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))];
        let _209_v59;
        let _210_v60;
        let _211_v61;
        let _212_v62;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = _module.__default.m0(_207_v57, _137_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _209_v59 = _out4;
        _210_v60 = _out5;
        _211_v61 = _out6;
        _212_v62 = _out7;
        let _index34 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        (_134_v0)[_index34] = !(_139_v2);
        let _213_v63;
        _213_v63 = _dafny.Map.Empty.slice().updateUnsafe(_207_v57,false);
        let _214_v64;
        _214_v64 = _dafny.Seq.UnicodeFromString("ljigfrq");
        let _215_v65;
        _215_v65 = _dafny.Map.Empty.slice().updateUnsafe(_207_v57,_214_v64);
        if ((new BigNumber((((_213_v63).update(new BigNumber(((((_215_v65).contains(_207_v57)) ? ((_215_v65).get(_207_v57)) : (_214_v64))).length), (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])).Merge(_213_v63)).length)).isLessThanOrEqualTo(_207_v57)) {
          _210_v60 = _139_v2;
          (_137_globalState).f2 = new BigNumber(387);
          let _216_v66;
          let _nw38 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _216_v66 = _nw38;
          _216_v66 = _216_v66;
          let _217_v67;
          let _nw39 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _217_v67 = _nw39;
          let _218_v68;
          let _init6 = function (_219_i6) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          };
          let _nw40 = Array((new BigNumber(25)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw40.length); _i0_6++) {
            _nw40[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _218_v68 = _nw40;
          let _220_v69;
          _220_v69 = new _dafny.CodePoint('j'.codePointAt(0));
          let _index35 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_218_v68).length));
          (_218_v68)[_index35] = _220_v69;
          let _221_v70;
          let _nw41 = Array((new BigNumber(8)).toNumber()).fill(_module.D1.Default());
          _221_v70 = _nw41;
          let _index36 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_221_v70).length));
          (_221_v70)[_index36] = _module.__default.fm5(_137_globalState);
          let _222_v71;
          let _nw42 = new _module.C0();
          _nw42.__ctor();
          _222_v71 = _nw42;
          let _223_v72;
          _223_v72 = _dafny.Map.Empty.slice().updateUnsafe(_222_v71,_207_v57);
          let _224_v73;
          _224_v73 = _dafny.Seq.of((_207_v57).plus((((_223_v72).contains(_222_v71)) ? ((_223_v72).get(_222_v71)) : (new BigNumber(777)))));
          let _225_v74;
          _225_v74 = _module.D1.create_DC3(_211_v61, _211_v61, new BigNumber((_dafny.Set.fromElements(_210_v60)).length), _139_v2);
          let _pat_let_tv12 = _207_v57;
          let _pat_let_tv13 = _134_v0;
          let _pat_let_tv14 = _134_v0;
          let _index37 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_218_v68).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_221_v70).length));
          let _rhs34 = (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))];
          let _rhs35 = _217_v67;
          let _rhs36 = _220_v69;
          let _rhs37 = function (_pat_let6_0) {
            return function (_226_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_227_dt__update_hcf9_h0) {
                  return function (_pat_let8_0) {
                    return function (_228_dt__update_hcf12_h0) {
                      return _module.D1.create_DC3(_227_dt__update_hcf9_h0, (_226_dt__update__tmp_h0).dtor_cf10, (_226_dt__update__tmp_h0).dtor_cf11, _228_dt__update_hcf12_h0);
                    }(_pat_let8_0);
                  }((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_pat_let_tv13).length))]);
                }(_pat_let7_0);
              }((new BigNumber(574)).isLessThan(_pat_let_tv12));
            }(_pat_let6_0);
          }(_225_v74);
          let _rhs38 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-651)), ((_229_v57) => function (_230_i7) {
            return _229_v57;
          })(_207_v57));
          let _lhs20 = _134_v0;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
          let _lhs22 = _218_v68;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_218_v68).length));
          let _lhs24 = _221_v70;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_221_v70).length));
          _lhs20[_lhs21] = _rhs34;
          _217_v67 = _rhs35;
          _lhs22[_lhs23] = _rhs36;
          _lhs24[_lhs25] = _rhs37;
          _224_v73 = _rhs38;
          let _231_v75;
          _231_v75 = _module.D0.create_DC0(_207_v57, _dafny.MultiSet.fromElements(_207_v57), _207_v57);
          let _232_v76;
          _232_v76 = _dafny.Seq.of(_134_v0);
          let _233_v77;
          _233_v77 = _module.D1.create_DC2(_module.__default.safeModuloInt(_207_v57, _207_v57), _module.__default.fm1(_137_globalState), ((_231_v75).dtor_cf0).multipliedBy(new BigNumber((_232_v76).length)), _224_v73, (((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]) ? (_220_v69) : ((_218_v68)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_218_v68).length))])));
          _233_v77 = _233_v77;
        } else {
          let _234_v78;
          _234_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(351),_207_v57);
          let _235_v79;
          _235_v79 = _dafny.MultiSet.fromElements(!(_139_v2), _211_v61, (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]);
          let _236_v80;
          _236_v80 = _dafny.Seq.of(_139_v2);
          let _237_v81;
          let _nw43 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _237_v81 = _nw43;
          let _238_v82;
          _238_v82 = _dafny.Set.fromElements(_237_v81, _237_v81, _237_v81);
          let _239_v83;
          let _nw44 = Array((new BigNumber(28)).toNumber());
          _nw44[(_dafny.ZERO).toNumber()] = _207_v57;
          _nw44[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_207_v57);
          _nw44[(new BigNumber(2)).toNumber()] = (_207_v57).multipliedBy(new BigNumber((_208_v58).cardinality()));
          _nw44[(new BigNumber(3)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(4)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(5)).toNumber()] = (new BigNumber((_214_v64).length)).plus(new BigNumber(352));
          _nw44[(new BigNumber(6)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(7)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(8)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(9)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(10)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(11)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(12)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(13)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(14)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(15)).toNumber()] = (_207_v57).plus((((_208_v58).contains(_207_v57)) ? ((_208_v58).get(_207_v57)) : (_207_v57)));
          _nw44[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(_207_v57, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-382)), function (_240_i8) {
            return new BigNumber(53);
          })).length));
          _nw44[(new BigNumber(17)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(18)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(19)).toNumber()] = (_207_v57).minus(_207_v57);
          _nw44[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_234_v78).length), _207_v57);
          _nw44[(new BigNumber(21)).toNumber()] = new BigNumber(94);
          _nw44[(new BigNumber(22)).toNumber()] = _207_v57;
          _nw44[(new BigNumber(23)).toNumber()] = (((_235_v79).contains(_211_v61)) ? ((_235_v79).get(_211_v61)) : (_207_v57));
          _nw44[(new BigNumber(24)).toNumber()] = new BigNumber((((_210_v60) ? (_236_v80) : (_236_v80))).length);
          _nw44[(new BigNumber(25)).toNumber()] = (new BigNumber(599)).minus(_207_v57);
          _nw44[(new BigNumber(26)).toNumber()] = (((_235_v79).contains((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])) ? ((_235_v79).get((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])) : (_207_v57));
          _nw44[(new BigNumber(27)).toNumber()] = new BigNumber((_238_v82).length);
          _239_v83 = _nw44;
          let _index40 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length));
          (_237_v81)[_index40] = _207_v57;
          let _index41 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
          let _index42 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length));
          let _rhs39 = _module.__default.fm2(_137_globalState);
          let _rhs40 = true;
          let _rhs41 = new BigNumber(821);
          let _rhs42 = _210_v60;
          let _lhs26 = _134_v0;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
          let _lhs28 = _237_v81;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length));
          _211_v61 = _rhs39;
          _lhs26[_lhs27] = _rhs40;
          _lhs28[_lhs29] = _rhs41;
          _211_v61 = _rhs42;
          (_137_globalState).f2 = (_dafny.ZERO).minus((_237_v81)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length))]);
          let _241_v84;
          _241_v84 = _module.D1.create_DC3(!(!((_236_v80)[_module.__default.safeIndex(_207_v57, new BigNumber((_236_v80).length))])), _211_v61, (_237_v81)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length))], _211_v61);
          _211_v61 = (_241_v84).dtor_cf12;
          let _index43 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length));
          (_237_v81)[_index43] = (_dafny.ZERO).minus(((_210_v60) ? (_207_v57) : ((_237_v81)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_237_v81).length))])));
          _207_v57 = new BigNumber(-825);
        }
      }
      let _242_v85;
      _242_v85 = new BigNumber(-845);
      let _hi0 = _module.__default.safeDivisionInt(_242_v85, _module.__default.fm6(_242_v85, _137_globalState));
      for (let _243_i9 = _242_v85; _243_i9.isLessThan(_hi0); _243_i9 = _243_i9.plus(_dafny.ONE)) {
        let _244_v86;
        _244_v86 = _dafny.Seq.UnicodeFromString("fjukyxth");
        let _245_v87;
        _245_v87 = _dafny.Map.Empty.slice().updateUnsafe(_243_i9,_244_v86);
        _245_v87 = (_245_v87).update((((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]) ? (new BigNumber((_dafny.Set.fromElements(_242_v85, _242_v85, new BigNumber((_module.__default.fm7(_243_i9, _137_globalState)).length))).length)) : (_242_v85)), _244_v86);
        let _246_v88;
        let _nw45 = new _module.C0();
        _nw45.__ctor();
        _246_v88 = _nw45;
        let _247_v89;
        _247_v89 = _module.D1.create_DC1(_139_v2);
        let _248_v90;
        _248_v90 = _dafny.Seq.of((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))], (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]);
        let _pat_let_tv15 = _139_v2;
        let _index44 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _index45 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _rhs43 = false;
        let _rhs44 = (_248_v90)[_module.__default.safeIndex(((_139_v2) ? (_242_v85) : (_242_v85)), new BigNumber((_248_v90).length))];
        let _rhs45 = function (_pat_let9_0) {
          return function (_249_dt__update__tmp_h1) {
            return function (_pat_let10_0) {
              return function (_250_dt__update_hcf3_h0) {
                return _module.D1.create_DC1(_250_dt__update_hcf3_h0);
              }(_pat_let10_0);
            }(_pat_let_tv15);
          }(_pat_let9_0);
        }(_module.D1.create_DC1((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]));
        let _rhs46 = (_139_v2) === (((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]) && (_139_v2));
        let _rhs47 = _246_v88;
        let _lhs30 = _134_v0;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _lhs32 = _134_v0;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        _139_v2 = _rhs43;
        _lhs30[_lhs31] = _rhs44;
        _247_v89 = _rhs45;
        _lhs32[_lhs33] = _rhs46;
        _246_v88 = _rhs47;
        let _pat_let_tv16 = _134_v0;
        let _pat_let_tv17 = _134_v0;
        let _index46 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        (_134_v0)[_index46] = (function (_pat_let11_0) {
          return function (_251_dt__update__tmp_h2) {
            return function (_pat_let12_0) {
              return function (_252_dt__update_hcf3_h1) {
                return _module.D1.create_DC1(_252_dt__update_hcf3_h1);
              }(_pat_let12_0);
            }((_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_pat_let_tv16).length))]);
          }(_pat_let11_0);
        }(_247_v89)).dtor_cf3;
      }
      let _253_v91;
      _253_v91 = _dafny.Seq.UnicodeFromString("dbygq");
      let _254_v92;
      _254_v92 = _dafny.Seq.of(_242_v85, _242_v85, _242_v85, new BigNumber((_253_v91).length));
      let _255_v93;
      let _256_v94;
      let _257_v95;
      let _258_v96;
      let _out8;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector2 = _module.__default.m0(new BigNumber((_dafny.Seq.Concat(_254_v92, _254_v92)).length), _137_globalState);
      _out8 = _outcollector2[0];
      _out9 = _outcollector2[1];
      _out10 = _outcollector2[2];
      _out11 = _outcollector2[3];
      _255_v93 = _out8;
      _256_v94 = _out9;
      _257_v95 = _out10;
      _258_v96 = _out11;
      let _259_v97;
      _259_v97 = _module.D3.create_DC8(_134_v0);
      let _260_v98;
      _260_v98 = _dafny.Map.Empty.slice().updateUnsafe(_242_v85,(_259_v97).dtor_cf16);
      _260_v98 = (_260_v98).update(_242_v85, _134_v0);
      let _261_v99;
      let _nw46 = new _module.C0();
      _nw46.__ctor();
      _261_v99 = _nw46;
      let _262_v100;
      _262_v100 = _dafny.Map.Empty.slice().updateUnsafe((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))],_261_v99);
      _261_v99 = (((_262_v100).contains(_139_v2)) ? ((_262_v100).get(_139_v2)) : (_261_v99));
      let _263_v101;
      _263_v101 = _dafny.MultiSet.fromElements(_242_v85);
      let _264_v102;
      let _265_v103;
      let _266_v104;
      let _267_v105;
      let _out12;
      let _out13;
      let _out14;
      let _out15;
      let _outcollector3 = _module.__default.m0(new BigNumber(((_263_v101).Difference(_263_v101)).cardinality()), _137_globalState);
      _out12 = _outcollector3[0];
      _out13 = _outcollector3[1];
      _out14 = _outcollector3[2];
      _out15 = _outcollector3[3];
      _264_v102 = _out12;
      _265_v103 = _out13;
      _266_v104 = _out14;
      _267_v105 = _out15;
      _254_v92 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_242_v85), _254_v92), _254_v92);
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_134_v0).length))) {
        let _268_i10 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_268_i10)) && ((_268_i10).isLessThan(new BigNumber((_134_v0).length))))) {
          (_134_v0)[(_268_i10)] = _257_v95;
        }
      }
      let _ingredients1 = [];
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_134_v0).length))) {
        let _269_i11 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_269_i11)) && ((_269_i11).isLessThan(new BigNumber((_134_v0).length))))) {
          _ingredients1.push(_dafny.Tuple.of(_134_v0, (_269_i11).toNumber(), (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]));
        }
      }
      for (const _tup1 of _ingredients1) {
        _tup1[0][_tup1[1]] = _tup1[2];
      }
      let _hi1 = _242_v85;
      for (let _270_i12 = _242_v85; _270_i12.isLessThan(_hi1); _270_i12 = _270_i12.plus(_dafny.ONE)) {
        let _271_v106;
        _271_v106 = new _dafny.CodePoint('n'.codePointAt(0));
        let _272_v107;
        _272_v107 = _dafny.Seq.of(_254_v92, _254_v92);
        let _273_v108;
        _273_v108 = _module.D1.create_DC2(_242_v85, _271_v106, _270_i12, _dafny.Seq.Concat((_272_v107)[_module.__default.safeIndex(new BigNumber((_253_v91).length), new BigNumber((_272_v107).length))], _254_v92), _271_v106);
        let _274_v109;
        _274_v109 = _module.D3.create_DC10(_module.__default.fm2(_137_globalState), _253_v91, _265_v103);
        let _275_v110;
        _275_v110 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), _253_v91), _dafny.Seq.Concat(_253_v91, (_274_v109).dtor_cf19));
        let _276_v111;
        _276_v111 = _dafny.Seq.of(_256_v94);
        let _277_v112;
        _277_v112 = _dafny.Seq.of(_276_v111);
        let _rhs48 = (_275_v110)[_module.__default.safeIndex((((_263_v101).contains(new BigNumber(((_277_v112)[_module.__default.safeIndex(_242_v85, new BigNumber((_277_v112).length))]).length))) ? ((_263_v101).get(new BigNumber(((_277_v112)[_module.__default.safeIndex(_242_v85, new BigNumber((_277_v112).length))]).length))) : (_270_i12)), new BigNumber((_275_v110).length))];
        let _rhs49 = _273_v108;
        let _rhs50 = _module.__default.fm2(_137_globalState);
        let _rhs51 = _270_i12;
        let _lhs34 = _137_globalState;
        _lhs34.f11 = _rhs48;
        _273_v108 = _rhs49;
        _265_v103 = _rhs50;
        _242_v85 = _rhs51;
        let _278_v113;
        _278_v113 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(new BigNumber(54), _137_globalState),(_242_v85).isLessThan(new BigNumber(148)));
        let _279_v114;
        _279_v114 = _dafny.Map.Empty.slice().updateUnsafe(_270_i12,(_dafny.ZERO).minus(_242_v85));
        let _280_v115;
        _280_v115 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,(_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]);
        let _281_v116;
        _281_v116 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,(_dafny.ZERO).minus(new BigNumber((_280_v115).length)));
        let _282_v117;
        let _nw47 = Array((new BigNumber(17)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _242_v85;
        _nw47[(_dafny.ONE).toNumber()] = _242_v85;
        _nw47[(new BigNumber(2)).toNumber()] = _242_v85;
        _nw47[(new BigNumber(3)).toNumber()] = new BigNumber(((_279_v114).Merge(_279_v114)).length);
        _nw47[(new BigNumber(4)).toNumber()] = _270_i12;
        _nw47[(new BigNumber(5)).toNumber()] = _270_i12;
        _nw47[(new BigNumber(6)).toNumber()] = _270_i12;
        _nw47[(new BigNumber(7)).toNumber()] = _242_v85;
        _nw47[(new BigNumber(8)).toNumber()] = (((_281_v116).contains(_134_v0)) ? ((_281_v116).get(_134_v0)) : (_270_i12));
        _nw47[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_270_i12, _242_v85);
        _nw47[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(878), (_dafny.ZERO).minus(_242_v85));
        _nw47[(new BigNumber(11)).toNumber()] = _270_i12;
        _nw47[(new BigNumber(12)).toNumber()] = new BigNumber((_253_v91).length);
        _nw47[(new BigNumber(13)).toNumber()] = (_242_v85).multipliedBy(new BigNumber(((_274_v109).dtor_cf19).length));
        _nw47[(new BigNumber(14)).toNumber()] = _242_v85;
        _nw47[(new BigNumber(15)).toNumber()] = _242_v85;
        _nw47[(new BigNumber(16)).toNumber()] = _270_i12;
        _282_v117 = _nw47;
        let _283_v118;
        _283_v118 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm8(_242_v85, _266_v104, _139_v2, _137_globalState),_278_v113);
        let _284_v120;
        _284_v120 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_270_i12,_266_v104), (((_283_v118).contains(_dafny.Set.fromElements((_dafny.ZERO).minus(_242_v85)))) ? ((_283_v118).get(_dafny.Set.fromElements((_dafny.ZERO).minus(_242_v85)))) : (_278_v113)), _278_v113, function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_263_v101).Elements) {
            let _285_v119 = _compr_11;
            if ((_263_v101).contains(_285_v119)) {
              _coll11.push([(_285_v119).multipliedBy(_242_v85),true]);
            }
          }
          return _coll11;
        }(), (_278_v113).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-579),(_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])));
        let _index47 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _rhs52 = (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))];
        let _rhs53 = (_284_v120)[_module.__default.safeIndex(_270_i12, new BigNumber((_284_v120).length))];
        let _rhs54 = _282_v117;
        let _lhs35 = _134_v0;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        _lhs35[_lhs36] = _rhs52;
        _278_v113 = _rhs53;
        _282_v117 = _rhs54;
        let _286_v121;
        let _nw48 = Array((new BigNumber(17)).toNumber()).fill([]);
        _286_v121 = _nw48;
        let _287_v122;
        _287_v122 = _258_v96;
        let _288_v123;
        _288_v123 = _dafny.Map.Empty.slice().updateUnsafe(!((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]),_267_v105);
        let _289_v124;
        let _nw49 = Array((new BigNumber(26)).toNumber());
        _nw49[(_dafny.ZERO).toNumber()] = _258_v96;
        _nw49[(_dafny.ONE).toNumber()] = _258_v96;
        _nw49[(new BigNumber(2)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(3)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(4)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(false, false);
        _nw49[(new BigNumber(6)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(7)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(8)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(9)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(10)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(11)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(12)).toNumber()] = (_287_v122);
        _nw49[(new BigNumber(13)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(14)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(15)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(16)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(17)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(18)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(19)).toNumber()] = _258_v96;
        _nw49[(new BigNumber(20)).toNumber()] = (((_288_v123).contains(true)) ? ((_288_v123).get(true)) : (_258_v96));
        _nw49[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_257_v95, _256_v94);
        _nw49[(new BigNumber(22)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(23)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(24)).toNumber()] = _267_v105;
        _nw49[(new BigNumber(25)).toNumber()] = _267_v105;
        _289_v124 = _nw49;
        let _index48 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_286_v121).length));
        (_286_v121)[_index48] = _289_v124;
        let _290_v125;
        _290_v125 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(683)), ((_291_v106) => function (_292_i13) {
          return _291_v106;
        })(_271_v106)),_289_v124);
        let _index49 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_286_v121).length));
        (_286_v121)[_index49] = (((_290_v125).contains(_253_v91)) ? ((_290_v125).get(_253_v91)) : (_289_v124));
        (_137_globalState).f2 = _270_i12;
      }
      if (((_dafny.ZERO).minus(_242_v85)).isLessThanOrEqualTo(new BigNumber(745))) {
        let _293_v126;
        let _nw50 = Array((new BigNumber(4)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _254_v92;
        _nw50[(_dafny.ONE).toNumber()] = _254_v92;
        _nw50[(new BigNumber(2)).toNumber()] = _254_v92;
        _nw50[(new BigNumber(3)).toNumber()] = ((_module.__default.fm2(_137_globalState)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_294_v85) => function (_295_i14) {
          return _294_v85;
        })(_242_v85))) : (_254_v92));
        _293_v126 = _nw50;
        let _index50 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_293_v126).length));
        (_293_v126)[_index50] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(688)), ((_296_v85) => function (_297_i15) {
          return _296_v85;
        })(_242_v85));
        let _index51 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_293_v126).length));
        (_293_v126)[_index51] = _dafny.Seq.of((_dafny.ZERO).minus(_242_v85), (_254_v92)[_module.__default.safeIndex(_242_v85, new BigNumber((_254_v92).length))]);
        let _298_v127;
        _298_v127 = new _dafny.CodePoint('u'.codePointAt(0));
        let _299_v128;
        _299_v128 = _dafny.Set.fromElements(_298_v127, _298_v127);
        let _300_v129;
        _300_v129 = _dafny.Map.Empty.slice().updateUnsafe(_242_v85,_299_v128);
        let _301_v130;
        _301_v130 = _module.D5.create_DC13(_299_v128);
        _265_v103 = ((((_300_v129).contains(_242_v85)) ? ((_300_v129).get(_242_v85)) : ((_301_v130).dtor_cf23))).IsDisjointFrom((_299_v128).Intersect(_299_v128));
        let _302_v131;
        _302_v131 = _dafny.Seq.of(_139_v2);
        let _303_v132;
        let _init7 = ((_304_v96) => function (_305_i16) {
          return _304_v96;
        })(_258_v96);
        let _nw51 = Array((new BigNumber(5)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw51.length); _i0_7++) {
          _nw51[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _303_v132 = _nw51;
        let _306_v133;
        _306_v133 = _258_v96;
        let _index52 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_303_v132).length));
        (_303_v132)[_index52] = _306_v133;
        let _index53 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _index54 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_303_v132).length));
        let _rhs55 = _dafny.Seq.Concat(_302_v131, _302_v131);
        let _rhs56 = _139_v2;
        let _rhs57 = _267_v105;
        let _rhs58 = new BigNumber((_253_v91).length);
        let _lhs37 = _134_v0;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
        let _lhs39 = _303_v132;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_303_v132).length));
        let _lhs41 = _137_globalState;
        _302_v131 = _rhs55;
        _lhs37[_lhs38] = _rhs56;
        _lhs39[_lhs40] = _rhs57;
        _lhs41.f2 = _rhs58;
        _257_v95 = _266_v104;
        (_137_globalState).f2 = _242_v85;
      } else {
        if (!(!(_139_v2))) {
          let _307_v134;
          _307_v134 = _dafny.MultiSet.fromElements(_253_v91);
          let _308_v135;
          let _309_v136;
          let _310_v137;
          let _311_v138;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(new BigNumber((_307_v134).cardinality()), _137_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _308_v135 = _out16;
          _309_v136 = _out17;
          _310_v137 = _out18;
          _311_v138 = _out19;
          let _312_v139;
          let _init8 = ((_313_v85) => function (_314_i17) {
            return (_314_i17).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(948), _313_v85, _313_v85)).length));
          })(_242_v85);
          let _nw52 = Array((new BigNumber(13)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw52.length); _i0_8++) {
            _nw52[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _312_v139 = _nw52;
          let _init9 = function (_315_i18) {
            return (_315_i18).multipliedBy(new BigNumber(336));
          };
          let _nw53 = Array((new BigNumber(9)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw53.length); _i0_9++) {
            _nw53[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _312_v139 = _nw53;
          let _316_v140;
          _316_v140 = new _dafny.CodePoint('h'.codePointAt(0));
          _316_v140 = _316_v140;
          _139_v2 = _266_v104;
          let _index55 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
          (_134_v0)[_index55] = _256_v94;
        } else {
          let _317_v141;
          let _318_v142;
          let _319_v143;
          let _320_v144;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0(_242_v85, _137_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _317_v141 = _out20;
          _318_v142 = _out21;
          _319_v143 = _out22;
          _320_v144 = _out23;
          _261_v99 = _261_v99;
          _261_v99 = ((_266_v104) ? (_261_v99) : (_261_v99));
          let _321_v145;
          _321_v145 = _dafny.Map.Empty.slice().updateUnsafe(_242_v85,(new BigNumber(-784)).plus(_242_v85));
          _321_v145 = (_321_v145).update(_242_v85, _242_v85);
          let _322_v146;
          _322_v146 = _dafny.Set.fromElements(_263_v101);
          let _323_v147;
          _323_v147 = _dafny.Set.fromElements(_261_v99, _261_v99);
          let _324_v148;
          _324_v148 = _dafny.Map.Empty.slice().updateUnsafe(_266_v104,_139_v2);
          let _325_v149;
          let _nw54 = Array((new BigNumber(20)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = _242_v85;
          _nw54[(_dafny.ONE).toNumber()] = _242_v85;
          _nw54[(new BigNumber(2)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(3)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(4)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_242_v85, _module.__default.fm6(_242_v85, _137_globalState));
          _nw54[(new BigNumber(6)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(7)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(8)).toNumber()] = _module.__default.fm6(_242_v85, _137_globalState);
          _nw54[(new BigNumber(9)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(10)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_322_v146).length))).plus(_242_v85);
          _nw54[(new BigNumber(11)).toNumber()] = _module.__default.fm6(_242_v85, _137_globalState);
          _nw54[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_242_v85, _module.__default.fm6(new BigNumber((_253_v91).length), _137_globalState));
          _nw54[(new BigNumber(13)).toNumber()] = new BigNumber((_323_v147).length);
          _nw54[(new BigNumber(14)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(15)).toNumber()] = _242_v85;
          _nw54[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_242_v85);
          _nw54[(new BigNumber(17)).toNumber()] = new BigNumber(-986);
          _nw54[(new BigNumber(18)).toNumber()] = new BigNumber((_324_v148).length);
          _nw54[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(_242_v85, _242_v85);
          _325_v149 = _nw54;
          let _326_v150;
          _326_v150 = _dafny.Seq.of(_325_v149, _325_v149, _325_v149, _325_v149, _325_v149);
          _325_v149 = (_326_v150)[_module.__default.safeIndex((_242_v85).multipliedBy(_242_v85), new BigNumber((_326_v150).length))];
        }
        let _327_v151;
        _327_v151 = _dafny.Map.Empty.slice().updateUnsafe(_242_v85,_265_v103);
        let _328_v152;
        _328_v152 = _dafny.MultiSet.fromElements(true, _266_v104);
        let _329_v153;
        _329_v153 = _dafny.Map.Empty.slice().updateUnsafe(_261_v99,_328_v152);
        let _rhs59 = !((((_327_v151).contains(_module.__default.safeModuloInt(_242_v85, new BigNumber(((((_329_v153).contains(_261_v99)) ? ((_329_v153).get(_261_v99)) : (_dafny.MultiSet.fromElements((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])))).cardinality())))) ? ((_327_v151).get(_module.__default.safeModuloInt(_242_v85, new BigNumber(((((_329_v153).contains(_261_v99)) ? ((_329_v153).get(_261_v99)) : (_dafny.MultiSet.fromElements((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))])))).cardinality())))) : (_139_v2)));
        let _rhs60 = _module.__default.safeModuloInt(_242_v85, _242_v85);
        let _lhs42 = _137_globalState;
        _266_v104 = _rhs59;
        _lhs42.f2 = _rhs60;
        _242_v85 = _242_v85;
        (_137_globalState).f2 = new BigNumber((_253_v91).length);
        (_137_globalState).f2 = ((_256_v94) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_242_v85)))) : (((_266_v104) ? (_module.__default.fm6(_242_v85, _137_globalState)) : (_242_v85))));
      }
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_134_v0).length))) {
        let _330_i19 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_330_i19)) && ((_330_i19).isLessThan(new BigNumber((_134_v0).length))))) {
          (_134_v0)[(_330_i19)] = _139_v2;
        }
      }
      let _331_i20;
      _331_i20 = _dafny.ZERO;
      L1: {
        while (_265_v103) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_331_i20)) {
              break L1;
            }
            _331_i20 = (_331_i20).plus(_dafny.ONE);
            let _332_v154;
            let _nw55 = Array((new BigNumber(5)).toNumber());
            _nw55[(_dafny.ZERO).toNumber()] = _253_v91;
            _nw55[(_dafny.ONE).toNumber()] = _253_v91;
            _nw55[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("cugllwi");
            _nw55[(new BigNumber(3)).toNumber()] = _253_v91;
            _nw55[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_253_v91, _253_v91);
            _332_v154 = _nw55;
            let _index56 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_332_v154).length));
            (_332_v154)[_index56] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_333_i21) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            });
            let _index57 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_332_v154).length));
            (_332_v154)[_index57] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-999)), ((_334_v104) => function (_335_i22) {
              return ((_334_v104) ? (new _dafny.CodePoint('m'.codePointAt(0))) : (new _dafny.CodePoint('u'.codePointAt(0))));
            })(_266_v104));
            if ((new BigNumber(911)).isLessThanOrEqualTo(new BigNumber((_254_v92).length))) {
              let _336_v155;
              let _337_v156;
              let _338_v157;
              let _339_v158;
              let _out24;
              let _out25;
              let _out26;
              let _out27;
              let _outcollector6 = _module.__default.m0(_module.__default.fm6(_242_v85, _137_globalState), _137_globalState);
              _out24 = _outcollector6[0];
              _out25 = _outcollector6[1];
              _out26 = _outcollector6[2];
              _out27 = _outcollector6[3];
              _336_v155 = _out24;
              _337_v156 = _out25;
              _338_v157 = _out26;
              _339_v158 = _out27;
              let _340_v159;
              _340_v159 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_257_v95) ? (_242_v85) : (_242_v85))),_337_v156);
              _340_v159 = (_340_v159).update((_dafny.ZERO).minus((_242_v85).minus(_242_v85)), !((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]) || (_257_v95));
              (_137_globalState).f2 = _242_v85;
              let _341_v160;
              _341_v160 = _dafny.MultiSet.fromElements((_332_v154)[_module.__default.safeIndex(new BigNumber(240), new BigNumber((_332_v154).length))]);
              let _342_v161;
              _342_v161 = new _dafny.CodePoint('q'.codePointAt(0));
              let _rhs61 = _dafny.MultiSet.fromElements(_253_v91, _253_v91);
              let _rhs62 = _342_v161;
              _341_v160 = _rhs61;
              _342_v161 = _rhs62;
              let _343_v162;
              let _nw56 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
              _343_v162 = _nw56;
              let _344_v163;
              _344_v163 = _dafny.Set.fromElements(_343_v162);
              let _345_v164;
              _345_v164 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_344_v163).length),_242_v85);
              (_137_globalState).f2 = _module.__default.fm6((_242_v85).minus(new BigNumber((_345_v164).length)), _137_globalState);
            } else {
              let _346_v165;
              let _347_v166;
              let _348_v167;
              let _349_v168;
              let _out28;
              let _out29;
              let _out30;
              let _out31;
              let _outcollector7 = _module.__default.m0(_242_v85, _137_globalState);
              _out28 = _outcollector7[0];
              _out29 = _outcollector7[1];
              _out30 = _outcollector7[2];
              _out31 = _outcollector7[3];
              _346_v165 = _out28;
              _347_v166 = _out29;
              _348_v167 = _out30;
              _349_v168 = _out31;
              _347_v166 = _266_v104;
              (_137_globalState).f2 = _242_v85;
              let _350_v169;
              let _nw57 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
              _350_v169 = _nw57;
              let _351_v170;
              _351_v170 = _dafny.Map.Empty.slice().updateUnsafe(!((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]),_350_v169);
              let _352_v171;
              _352_v171 = _module.D2.create_DC5(_dafny.Seq.of(_139_v2, _139_v2, (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]));
              let _353_v172;
              _353_v172 = _module.D2.create_DC7(_352_v171);
              let _354_v173;
              _354_v173 = _dafny.Seq.of(_353_v172, _module.D2.create_DC7(_module.D2.create_DC6()));
              let _355_v174;
              _355_v174 = _dafny.MultiSet.fromElements(_354_v173);
              let _356_v175;
              _356_v175 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_242_v85);
              let _357_v176;
              _357_v176 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,_347_v166);
              let _358_v177;
              let _nw58 = Array((new BigNumber(21)).toNumber());
              _nw58[(_dafny.ZERO).toNumber()] = (_254_v92)[_module.__default.safeIndex(new BigNumber((_351_v170).length), new BigNumber((_254_v92).length))];
              _nw58[(_dafny.ONE).toNumber()] = _242_v85;
              _nw58[(new BigNumber(2)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(3)).toNumber()] = new BigNumber((_355_v174).cardinality());
              _nw58[(new BigNumber(4)).toNumber()] = new BigNumber((_356_v175).length);
              _nw58[(new BigNumber(5)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_242_v85);
              _nw58[(new BigNumber(7)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(8)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(9)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(10)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(11)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(12)).toNumber()] = new BigNumber(-668);
              _nw58[(new BigNumber(13)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(14)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(15)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(16)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(17)).toNumber()] = _242_v85;
              _nw58[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_348_v167,new BigNumber(288))).length);
              _nw58[(new BigNumber(19)).toNumber()] = new BigNumber((_357_v176).length);
              _nw58[(new BigNumber(20)).toNumber()] = _242_v85;
              _358_v177 = _nw58;
              let _359_v178;
              _359_v178 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-755),(_350_v169));
              let _360_v179;
              let _nw59 = Array((new BigNumber(15)).toNumber());
              _nw59[(_dafny.ZERO).toNumber()] = _358_v177;
              _nw59[(_dafny.ONE).toNumber()] = _358_v177;
              _nw59[(new BigNumber(2)).toNumber()] = (((_359_v178).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_242_v85)).length)))) ? ((_359_v178).get((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_242_v85)).length)))) : (_350_v169));
              _nw59[(new BigNumber(3)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(4)).toNumber()] = _350_v169;
              _nw59[(new BigNumber(5)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(6)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(7)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(8)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(9)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(10)).toNumber()] = _358_v177;
              _nw59[(new BigNumber(11)).toNumber()] = _350_v169;
              _nw59[(new BigNumber(12)).toNumber()] = _350_v169;
              _nw59[(new BigNumber(13)).toNumber()] = _350_v169;
              _nw59[(new BigNumber(14)).toNumber()] = _350_v169;
              _360_v179 = _nw59;
              let _index58 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_360_v179).length));
              (_360_v179)[_index58] = _350_v169;
              let _index59 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_360_v179).length));
              let _index60 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              let _index61 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              let _rhs63 = _256_v94;
              let _rhs64 = _350_v169;
              let _rhs65 = (((_242_v85).isLessThan(_242_v85)) ? (_347_v166) : (_257_v95));
              let _rhs66 = _139_v2;
              let _lhs43 = _360_v179;
              let _lhs44 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_360_v179).length));
              let _lhs45 = _134_v0;
              let _lhs46 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              let _lhs47 = _134_v0;
              let _lhs48 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              _256_v94 = _rhs63;
              _lhs43[_lhs44] = _rhs64;
              _lhs45[_lhs46] = _rhs65;
              _lhs47[_lhs48] = _rhs66;
              let _index62 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              (_134_v0)[_index62] = _265_v103;
            }
            if (_265_v103) {
              let _361_v180;
              _361_v180 = _dafny.MultiSet.fromElements(_256_v94, !(_265_v103), false, _266_v104);
              let _362_v181;
              _362_v181 = _dafny.Seq.of(_361_v180, _361_v180, _361_v180, _361_v180);
              _362_v181 = _dafny.Seq.Concat(_362_v181, _dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_363_v95, _364_v104) => function (_365_i23) {
                return _dafny.MultiSet.fromElements(_363_v95, _364_v104);
              })(_257_v95, _266_v104)));
              _361_v180 = (_362_v181)[_module.__default.safeIndex(_242_v85, new BigNumber((_362_v181).length))];
              let _index63 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              (_134_v0)[_index63] = !(_256_v94) || (!(((_dafny.ZERO).minus(_242_v85)).isLessThanOrEqualTo(_242_v85)));
              let _index64 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              (_134_v0)[_index64] = false;
              _257_v95 = _266_v104;
            } else {
              let _366_v182;
              _366_v182 = _dafny.Seq.of(_267_v105, _258_v96, _267_v105);
              let _367_v183;
              _367_v183 = _module.__default.fm9(_263_v101, _266_v104, _139_v2, _137_globalState);
              let _368_v184;
              _368_v184 = new _dafny.CodePoint('w'.codePointAt(0));
              let _369_v185;
              _369_v185 = _dafny.Map.Empty.slice().updateUnsafe(_368_v184,_dafny.Set.fromElements(_265_v103, _256_v94, _265_v103, _139_v2, _257_v95));
              let _370_v186;
              _370_v186 = _dafny.Map.Empty.slice().updateUnsafe(_368_v184,_368_v184);
              let _371_v187;
              let _nw60 = Array((new BigNumber(22)).toNumber());
              _nw60[(_dafny.ZERO).toNumber()] = (_267_v105).Union(_258_v96);
              _nw60[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(false, _265_v103)).Union(_267_v105);
              _nw60[(new BigNumber(2)).toNumber()] = _module.__default.fm9(_dafny.MultiSet.fromElements(_242_v85, _242_v85), true, _266_v104, _137_globalState);
              _nw60[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(false);
              _nw60[(new BigNumber(4)).toNumber()] = (_366_v182)[_module.__default.safeIndex(_module.__default.fm6(_242_v85, _137_globalState), new BigNumber((_366_v182).length))];
              _nw60[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_265_v103, _266_v104)).Intersect(_dafny.Set.fromElements(true, _256_v94, _139_v2, _266_v104));
              _nw60[(new BigNumber(6)).toNumber()] = _267_v105;
              _nw60[(new BigNumber(7)).toNumber()] = (_367_v183);
              _nw60[(new BigNumber(8)).toNumber()] = (_258_v96).Intersect(_dafny.Set.fromElements(_139_v2, _266_v104));
              _nw60[(new BigNumber(9)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(10)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(11)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(12)).toNumber()] = (((_369_v185).contains((((_370_v186).contains(_368_v184)) ? ((_370_v186).get(_368_v184)) : (_368_v184)))) ? ((_369_v185).get((((_370_v186).contains(_368_v184)) ? ((_370_v186).get(_368_v184)) : (_368_v184)))) : (_258_v96));
              _nw60[(new BigNumber(13)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(14)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(15)).toNumber()] = _267_v105;
              _nw60[(new BigNumber(16)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(17)).toNumber()] = ((_256_v94) ? (_258_v96) : ((_367_v183)));
              _nw60[(new BigNumber(18)).toNumber()] = _module.__default.fm9(_263_v101, !(_266_v104), _256_v94, _137_globalState);
              _nw60[(new BigNumber(19)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(20)).toNumber()] = _258_v96;
              _nw60[(new BigNumber(21)).toNumber()] = _267_v105;
              _371_v187 = _nw60;
              let _index65 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_371_v187).length));
              (_371_v187)[_index65] = _267_v105;
              let _372_v188;
              _372_v188 = _dafny.Map.Empty.slice().updateUnsafe(_368_v184,_263_v101);
              let _index66 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_371_v187).length));
              (_371_v187)[_index66] = (((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]) ? (_258_v96) : (_module.__default.fm9((((_372_v188).contains(_368_v184)) ? ((_372_v188).get(_368_v184)) : (_263_v101)), (_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))], _257_v95, _137_globalState)));
              let _index67 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length));
              (_134_v0)[_index67] = false;
              let _373_v189;
              _373_v189 = _dafny.Seq.of((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]);
              (_137_globalState).f2 = _module.__default.safeDivisionInt((_242_v85).minus(new BigNumber(254)), (_242_v85).plus((((_263_v101).contains(_242_v85)) ? ((_263_v101).get(_242_v85)) : (new BigNumber((_373_v189).length)))));
              _139_v2 = (_256_v94) || ((_134_v0)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_134_v0).length))]);
              (_137_globalState).f2 = (_242_v85).multipliedBy(_module.__default.fm6(_242_v85, _137_globalState));
            }
            let _374_v190;
            let _nw61 = Array((new BigNumber(3)).toNumber());
            _nw61[(_dafny.ZERO).toNumber()] = _242_v85;
            _nw61[(_dafny.ONE).toNumber()] = (new BigNumber(-198)).multipliedBy(_242_v85);
            _nw61[(new BigNumber(2)).toNumber()] = _module.__default.fm6(new BigNumber((_253_v91).length), _137_globalState);
            _374_v190 = _nw61;
            let _375_v191;
            _375_v191 = _dafny.Map.Empty.slice().updateUnsafe(_139_v2,_242_v85);
            let _376_v192;
            _376_v192 = _dafny.Map.Empty.slice().updateUnsafe(_375_v191,_266_v104);
            let _index68 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_374_v190).length));
            (_374_v190)[_index68] = (_242_v85).multipliedBy((_dafny.ZERO).minus(new BigNumber((_376_v192).length)));
            let _377_v193;
            _377_v193 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(196),_242_v85);
            let _378_v194;
            _378_v194 = new _dafny.CodePoint('t'.codePointAt(0));
            let _index69 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_374_v190).length));
            (_374_v190)[_index69] = ((_module.__default.fm10(_377_v193, _265_v103, _266_v104, false, _137_globalState)).dtor_cf4).minus(_module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_378_v194,_261_v99)).update(_378_v194, _261_v99)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dwaa")).length)));
          }
        }
      }
      process.stdout.write(_dafny.toString((_134_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(11)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(12)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(13)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(14)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(15)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(16)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(17)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(18)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(19)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v1)[new BigNumber(20)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(11)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(12)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(13)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(14)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(15)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(16)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(17)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(18)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(19)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_137_globalState).f9)[new BigNumber(20)]).equals(_dafny.Set.fromElements(new BigNumber(-34)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_137_globalState.f11, _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_196_v47)[_dafny.ZERO]).dtor_cf13).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_196_v47)[_dafny.ONE]).dtor_cf13).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_242_v85));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_253_v91).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_254_v92, _dafny.Seq.of(new BigNumber(-845), new BigNumber(-845), new BigNumber(-845), new BigNumber(-845), new BigNumber(5), new BigNumber(-845), new BigNumber(-845), new BigNumber(-845), new BigNumber(5)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_255_v93, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(false, false, new BigNumber(8), false),false), _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(false, false, new BigNumber(8), false),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_256_v94));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_257_v95));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v96).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v97).dtor_cf16)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_260_v98).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_262_v100).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v101).equals(_dafny.MultiSet.fromElements(new BigNumber(-845)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_264_v102, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(false, false, _dafny.ZERO, false),false), _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(false, false, _dafny.ZERO, false),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_265_v103));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_266_v104));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v105).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_331_i20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, _dafny.MultiSet.Empty, _dafny.ZERO);
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
    static create_DC2(cf4, cf5, cf6, cf7, cf8) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf9, cf10, cf11, cf12) {
      let $dt = new D1(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC1(cf3) {
      let $dt = new D1(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC4(cf13) {
      let $dt = new D1(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC1() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf9 === other.cf9 && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf3 === other.cf3;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC6() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC5(cf14) {
      let $dt = new D2(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf15) {
      let $dt = new D2(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf15) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6();
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
    static create_DC9(cf17) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf18, cf19, cf20) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC8(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC11(cf21) {
      let $dt = new D3(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf18) + ", " + this.cf19.toVerbatimString(true) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf16 === other.cf16;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO);
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
    static create_DC12(cf22) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf24, cf25, cf26, cf27) {
      let $dt = new D5(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC13(cf23) {
      let $dt = new D5(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(false, false, null, _dafny.ZERO);
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
    static create_DC15(cf28) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
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
    static create_DC16(cf29) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf29) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC17(cf30) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30);
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.ZERO;
      this.f11 = _dafny.Seq.UnicodeFromString("");
      this._f0 = false;
      this._f1 = [];
      this._f3 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f9 = [];
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
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
    get f3() {
      let _this = this;
      return _this._f3;
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
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
