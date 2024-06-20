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
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements((_module.D4.create_DC13(new BigNumber((_dafny.Seq.of(new BigNumber(47))).length), new BigNumber(-267), true)).dtor_cf20, false, false, !(true), false)).cardinality()), (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(39)), function (_0_i0) {
        return _dafny.Seq.UnicodeFromString("hrdfpvvp");
      })).length)).minus(new BigNumber((_dafny.Seq.of(true, (_module.D0.create_DC0(true)).dtor_cf0)).length)));
    };
    static fm1(globalState) {
      let _source0 = _module.D0.create_DC2(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true)), new BigNumber(484)));
      if (_source0.is_DC1) {
        let _1___mcc_h0 = (_source0).cf1;
        let _2___mcc_h1 = (_source0).cf2;
        let _3_cf2 = _2___mcc_h1;
        let _4_cf1 = _1___mcc_h0;
        return _module.D0.create_DC1(_4_cf1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_3_cf2)).length));
      } else if (_source0.is_DC0) {
        let _5___mcc_h2 = (_source0).cf0;
        let _6_cf0 = _5___mcc_h2;
        return _module.D0.create_DC1(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(_6_cf0))), new BigNumber(170));
      } else {
        let _7___mcc_h3 = (_source0).cf3;
        let _8_cf3 = _7___mcc_h3;
        return _module.D0.create_DC1(_dafny.MultiSet.fromElements(_dafny.Seq.of(!(false))), new BigNumber(261));
      }
    };
    static fm2(p0, p1, globalState) {
      return (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-262))).length)))).Union(function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(481), new BigNumber(21))) {
          let _9_v0 = _compr_0;
          if (((new BigNumber(481)).isLessThanOrEqualTo(_9_v0)) && ((_9_v0).isLessThan(new BigNumber(21)))) {
            _coll0.add((_9_v0).minus(new BigNumber(-863)));
          }
        }
        return _coll0;
      }());
    };
    static fm3(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('n'.codePointAt(0))),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('i'.codePointAt(0))),true));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return (_module.D3.create_DC9(_dafny.Seq.of(false))).dtor_cf10;
    };
    static fm7(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(((_module.D11.create_DC32(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), function (_10_i0) {
  return _dafny.Seq.UnicodeFromString("re");
})))).dtor_cf55).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber(517), new BigNumber(44))).length)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber(833), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false))).length)))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge((_module.D12.create_DC35(_dafny.Map.Empty.slice().updateUnsafe(false,true))).dtor_cf58)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)));
    };
    static fm9(globalState) {
      return _module.D1.create_DC3(_dafny.Seq.of(new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-10), new BigNumber(691))) {
    let _11_v0 = _compr_1;
    if (((new BigNumber(-10)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(691)))) {
      _coll1.push([(_11_v0).plus(new BigNumber(155)),new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)))).length)]);
    }
  }
  return _coll1;
}()).length)));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hrux"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(346)), function (_12_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('l'.codePointAt(0));
    };
    static fm12(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D3.create_DC11(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ihqlrtng")).length), true, !(true))).dtor_cf15)));
    };
    static fm13(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-488))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-168))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-929))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-715))));
    };
    static fm14(globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).Difference((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
          let _13_v0 = _compr_2;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).contains(_13_v0)) {
            _coll2.add(_13_v0);
          }
        }
        return _coll2;
      }()).Union(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))));
    };
    static fm15(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(738),(true) === (!(!(false))));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-350),new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).length))).Merge((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Set.fromElements(new BigNumber(260))).Elements) {
          let _14_v0 = _compr_3;
          if ((_dafny.Set.fromElements(new BigNumber(260))).contains(_14_v0)) {
            _coll3.push([_module.__default.safeModuloInt(_14_v0, new BigNumber((_dafny.Seq.UnicodeFromString("uvbomwi")).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_15_i0) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll3;
      }()).Merge(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(970), new BigNumber(-978))) {
          let _16_v1 = _compr_4;
          if (((new BigNumber(970)).isLessThanOrEqualTo(_16_v1)) && ((_16_v1).isLessThan(new BigNumber(-978)))) {
            _coll4.push([(_16_v1).multipliedBy(new BigNumber(137)),new BigNumber(484)]);
          }
        }
        return _coll4;
      }()));
    };
    static fm17(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(true));
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _hi0 = p0;
      for (let _17_i0 = p0; _17_i0.isLessThan(_hi0); _17_i0 = _17_i0.plus(_dafny.ONE)) {
        let _18_v0;
        let _init0 = ((_19_p0, _20_i0) => function (_21_i1) {
          return ((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(899)), ((_22_p0) => function (_23_i2) {
            return _dafny.Seq.of(_22_p0, new BigNumber((_dafny.Set.fromElements(_22_p0, new BigNumber((_dafny.Seq.UnicodeFromString("wq")).length))).length), new BigNumber((_dafny.MultiSet.fromElements(_22_p0)).cardinality()));
          })(_19_p0))) : (_dafny.Seq.of(_dafny.Seq.of(_20_i0, _20_i0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), function (_24_i3) {
            return new BigNumber(143);
          }), _dafny.Seq.of(_19_p0))));
        })(p0, _17_i0);
        let _nw0 = Array((new BigNumber(25)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _18_v0 = _nw0;
        let _25_v1;
        _25_v1 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
        let _26_v2;
        _26_v2 = _dafny.Seq.of(p0, new BigNumber((_25_v1).length), p0);
        let _27_v3;
        _27_v3 = _dafny.Seq.of(_26_v2);
        let _index0 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_18_v0).length));
        (_18_v0)[_index0] = _27_v3;
        let _28_v4;
        _28_v4 = false;
        let _29_v5;
        let _nw1 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
        _29_v5 = _nw1;
        let _index1 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_29_v5).length));
        (_29_v5)[_index1] = _25_v1;
        let _30_v6;
        _30_v6 = new _dafny.CodePoint('n'.codePointAt(0));
        let _31_v7;
        _31_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_30_v6);
        let _32_v8;
        let _nw2 = new _module.C0();
        _nw2.__ctor(_17_i0, (_dafny.ZERO).minus(_17_i0), true);
        _32_v8 = _nw2;
        let _33_v9;
        _33_v9 = _module.D2.create_DC7(new BigNumber((_31_v7).length), _32_v8);
        let _index2 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_18_v0).length));
        let _index3 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_29_v5).length));
        let _rhs0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_34_v2) => function (_35_i4) {
          return _34_v2;
        })(_26_v2)), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-914)), ((_36_p0) => function (_37_i5) {
          return _36_p0;
        })(p0))));
        let _rhs1 = _28_v4;
        let _rhs2 = (_33_v9).dtor_cf7;
        let _rhs3 = (_32_v8).f9;
        let _rhs4 = _25_v1;
        let _lhs0 = _18_v0;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_18_v0).length));
        let _lhs2 = globalState;
        let _lhs3 = globalState;
        let _lhs4 = _29_v5;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_29_v5).length));
        _lhs0[_lhs1] = _rhs0;
        _28_v4 = _rhs1;
        _lhs2.f3 = _rhs2;
        _lhs3.f6 = _rhs3;
        _lhs4[_lhs5] = _rhs4;
        let _38_v10;
        let _nw3 = Array((new BigNumber(21)).toNumber());
        _38_v10 = _nw3;
        let _index4 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length));
        (_38_v10)[_index4] = _32_v8;
        let _index5 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length));
        (_38_v10)[_index5] = _32_v8;
        _28_v4 = _28_v4;
        if (_28_v4) {
          let _39_v11;
          let _init1 = ((_40_p1) => function (_41_i6) {
            return _module.__default.safeModuloInt(_41_i6, new BigNumber((_40_p1).cardinality()));
          })(p1);
          let _nw4 = Array((new BigNumber(3)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw4.length); _i0_1++) {
            _nw4[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _39_v11 = _nw4;
          let _42_v12;
          _42_v12 = _dafny.Seq.UnicodeFromString("kn");
          let _43_v13;
          _43_v13 = _dafny.Set.fromElements(_32_v8);
          let _44_v14;
          _44_v14 = _dafny.MultiSet.fromElements(_43_v13);
          let _index6 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length));
          (_39_v11)[_index6] = _module.__default.safeDivisionInt(new BigNumber((_42_v12).length), (((_44_v14).contains(_dafny.Set.fromElements((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))], _32_v8))) ? ((_44_v14).get(_dafny.Set.fromElements((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))], _32_v8))) : ((_26_v2)[_module.__default.safeIndex(new BigNumber((_42_v12).length), new BigNumber((_26_v2).length))])));
          let _index7 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length));
          (_39_v11)[_index7] = new BigNumber(551);
          let _45_v15;
          _45_v15 = _dafny.Map.Empty.slice().updateUnsafe((_32_v8).f9,(_32_v8).f9);
          (globalState).f6 = (new BigNumber(42)).multipliedBy(_module.__default.safeDivisionInt((_32_v8).f9, new BigNumber((_45_v15).length)));
          let _46_v16;
          _46_v16 = _dafny.Seq.of(_28_v4, _28_v4, !(p1).contains(!(!(_28_v4))));
          let _47_v17;
          _47_v17 = _module.D4.create_DC14((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))], _28_v4);
          let _48_v18;
          _48_v18 = _dafny.Map.Empty.slice().updateUnsafe((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))],_30_v6);
          let _49_v19;
          _49_v19 = _dafny.Map.Empty.slice().updateUnsafe(_47_v17,(((_48_v18).contains((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))])) ? ((_48_v18).get((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))])) : (_30_v6)));
          let _50_v20;
          _50_v20 = _dafny.MultiSet.fromElements(new BigNumber((_49_v19).length));
          let _51_v21;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_17_i0, new BigNumber((_46_v16).length), _28_v4);
          _51_v21 = _nw5;
          let _52_v22;
          _52_v22 = _dafny.Seq.of(_51_v21, _51_v21, _51_v21);
          let _index8 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length));
          let _rhs5 = _46_v16;
          let _rhs6 = _30_v6;
          let _rhs7 = (_dafny.ZERO).minus(new BigNumber((_50_v20).cardinality()));
          let _rhs8 = (((_dafny.ZERO).minus(new BigNumber((_52_v22).length))).multipliedBy((_39_v11)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length))])).plus((_32_v8).f8);
          let _lhs6 = globalState;
          let _lhs7 = _39_v11;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length));
          _46_v16 = _rhs5;
          _30_v6 = _rhs6;
          _lhs6.f6 = _rhs7;
          _lhs7[_lhs8] = _rhs8;
          let _53_v23;
          _53_v23 = _module.D6.create_DC17(_51_v21.f7, false, _28_v4);
          let _54_v24;
          _54_v24 = _dafny.Seq.of(_53_v23, _53_v23, _53_v23);
          _28_v4 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_54_v24, _dafny.Seq.of(_53_v23)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(576)), ((_55_v23) => function (_56_i7) {
            return _55_v23;
          })(_53_v23)), _dafny.Seq.update(_dafny.Seq.update(_54_v24, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ljh")).length), new BigNumber((_54_v24).length)), _53_v23), _module.__default.safeIndex((_32_v8).f8, new BigNumber((_dafny.Seq.update(_54_v24, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ljh")).length), new BigNumber((_54_v24).length)), _53_v23)).length)), _module.D6.create_DC17(_51_v21.f7, _51_v21.f7, _51_v21.f7))));
          let _57_v25;
          _57_v25 = _dafny.Seq.of(_46_v16);
          let _58_v26;
          _58_v26 = _dafny.MultiSet.fromElements(_46_v16, _46_v16, p2, p2);
          let _59_v27;
          _59_v27 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_57_v25), _58_v26);
          let _60_v28;
          _60_v28 = _module.D0.create_DC1((_59_v27)[_module.__default.safeIndex(new BigNumber((_42_v12).length), new BigNumber((_59_v27).length))], _module.__default.safeModuloInt(new BigNumber(233), p0));
          let _rhs9 = _46_v16;
          let _rhs10 = ((_module.__default.fm3((_39_v11)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_39_v11).length))], globalState)) ? (_32_v8) : ((_38_v10)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_38_v10).length))]));
          let _rhs11 = _60_v28;
          _46_v16 = _rhs9;
          _32_v8 = _rhs10;
          _60_v28 = _rhs11;
        } else {
          _25_v1 = (_29_v5)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_29_v5).length))];
          let _61_v29;
          _61_v29 = _dafny.Seq.UnicodeFromString("dled");
          let _62_v30;
          _62_v30 = _dafny.Seq.of(_61_v29, _dafny.Seq.UnicodeFromString("vsxctdar"), _61_v29, _dafny.Seq.UnicodeFromString("gmyfssb"));
          _62_v30 = _62_v30;
          let _63_v31;
          let _nw6 = new _module.C0();
          _nw6.__ctor((_32_v8).f8, (_32_v8).f9, (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_64_v6) => function (_65_i8) {
            return _64_v6;
          })(_30_v6))).length)).isLessThan(p0));
          _63_v31 = _nw6;
          let _66_v32;
          let _nw7 = new _module.C0();
          _nw7.__ctor(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), ((_67_v6) => function (_68_i9) {
            return _67_v6;
          })(_30_v6)), _module.__default.safeIndex((_32_v8).f8, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), ((_69_v6) => function (_70_i9) {
            return _69_v6;
          })(_30_v6))).length)), _30_v6)).length), new BigNumber(403), _28_v4);
          _66_v32 = _nw7;
          let _71_v33;
          _71_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
          let _72_v34;
          _72_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_71_v33).contains((_dafny.ZERO).minus(_17_i0))) ? ((_71_v33).get((_dafny.ZERO).minus(_17_i0))) : (_28_v4)),(_26_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_26_v2).length))]);
          let _73_v35;
          _73_v35 = _module.D3.create_DC10(_30_v6, _28_v4);
          _72_v34 = (_72_v34).update((_73_v35).dtor_cf12, _module.__default.safeModuloInt((_32_v8).f9, new BigNumber(-391)));
        }
      }
      let _74_v36;
      _74_v36 = true;
      let _75_v37;
      _75_v37 = _dafny.Seq.UnicodeFromString("deqbibs");
      let _hi1 = p0;
      for (let _76_i10 = ((!(_74_v36)) ? (new BigNumber((_75_v37).length)) : (new BigNumber((_75_v37).length))); _76_i10.isLessThan(_hi1); _76_i10 = _76_i10.plus(_dafny.ONE)) {
        let _77_v38;
        let _nw8 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
        _77_v38 = _nw8;
        let _78_v39;
        _78_v39 = _dafny.Seq.of(p0);
        let _79_v40;
        _79_v40 = new _dafny.CodePoint('x'.codePointAt(0));
        let _index9 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_77_v38).length));
        (_77_v38)[_index9] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber(340), _74_v36, _78_v39, _79_v40, globalState), p0, _76_i10)));
        let _80_v41;
        let _init2 = ((_81_v40) => function (_82_i11) {
          return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("uyntxwc"), _81_v40);
        })(_79_v40);
        let _nw9 = Array((new BigNumber(5)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw9.length); _i0_2++) {
          _nw9[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _80_v41 = _nw9;
        let _index10 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_80_v41).length));
        (_80_v41)[_index10] = (p0).isEqualTo(p0);
        let _83_v42;
        let _nw10 = Array((new BigNumber(10)).toNumber()).fill(_dafny.MultiSet.Empty);
        _83_v42 = _nw10;
        let _index11 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_83_v42).length));
        (_83_v42)[_index11] = p1;
        let _84_v43;
        _84_v43 = _module.D6.create_DC17(_74_v36, _74_v36, _74_v36);
        let _85_v44;
        _85_v44 = _dafny.Map.Empty.slice().updateUnsafe(_84_v43,p0);
        let _86_v45;
        _86_v45 = _dafny.MultiSet.fromElements(_85_v44, _85_v44, _85_v44, _85_v44);
        let _87_v46;
        let _nw11 = new _module.C0();
        _nw11.__ctor(new BigNumber(666), _76_i10, false);
        _87_v46 = _nw11;
        let _88_v47;
        _88_v47 = _module.D2.create_DC7(p0, _87_v46);
        let _89_v48;
        _89_v48 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), (_88_v47).dtor_cf7, p0, p0);
        let _90_v49;
        _90_v49 = _module.D1.create_DC3(_78_v39);
        let _91_v50;
        _91_v50 = _module.D7.create_DC19(_dafny.MultiSet.FromArray(_78_v39));
        let _92_v51;
        _92_v51 = _dafny.MultiSet.fromElements(_89_v48, _dafny.MultiSet.FromArray((_90_v49).dtor_cf4), (_91_v50).dtor_cf29, _89_v48);
        let _index12 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_77_v38).length));
        let _index13 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_80_v41).length));
        let _index14 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_83_v42).length));
        let _rhs12 = !(!(false)) || ((_86_v45).IsDisjointFrom(_86_v45));
        let _rhs13 = _92_v51;
        let _rhs14 = _80_v41;
        let _rhs15 = !_dafny.areEqual(_module.__default.fm7(globalState), _78_v39);
        let _rhs16 = _dafny.MultiSet.FromArray(p2);
        let _lhs9 = _77_v38;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_77_v38).length));
        let _lhs11 = _80_v41;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_80_v41).length));
        let _lhs13 = _83_v42;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_83_v42).length));
        _74_v36 = _rhs12;
        _lhs9[_lhs10] = _rhs13;
        _80_v41 = _rhs14;
        _lhs11[_lhs12] = _rhs15;
        _lhs13[_lhs14] = _rhs16;
        let _93_v52;
        let _nw12 = new _module.C0();
        _nw12.__ctor((_87_v46).f9, p0, !((_80_v41)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_80_v41).length))]));
        _93_v52 = _nw12;
        let _94_v53;
        _94_v53 = _dafny.Seq.of(_93_v52);
        let _95_v54;
        let _nw13 = Array((new BigNumber(4)).toNumber());
        _nw13[(_dafny.ZERO).toNumber()] = p0;
        _nw13[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt((_87_v46).f8, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xpqepkn"),(_94_v53)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_94_v53).length))])).length));
        _nw13[(new BigNumber(2)).toNumber()] = (_87_v46).f9;
        _nw13[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_76_i10, (_87_v46).f9);
        _95_v54 = _nw13;
        let _index15 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_95_v54).length));
        (_95_v54)[_index15] = p0;
        let _index16 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_95_v54).length));
        (_95_v54)[_index16] = (_87_v46).f8;
        let _96_v55;
        _96_v55 = _dafny.Seq.of(_87_v46);
        let _97_v56;
        _97_v56 = _dafny.Set.fromElements(_76_i10, new BigNumber((_96_v55).length));
        let _98_v57;
        _98_v57 = _dafny.Set.fromElements((_97_v56).Union(_dafny.Set.fromElements(new BigNumber((p2).length))));
        _98_v57 = _dafny.Set.fromElements(_97_v56);
        (globalState).f3 = new BigNumber(-293);
      }
      let _99_v58;
      _99_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,_75_v37);
      _75_v37 = (((_99_v58).contains((_dafny.ZERO).minus((((p1).contains(false)) ? ((p1).get(false)) : (p0))))) ? ((_99_v58).get((_dafny.ZERO).minus((((p1).contains(false)) ? ((p1).get(false)) : (p0))))) : (_75_v37));
      if (!(p0).isEqualTo(new BigNumber(784))) {
        let _100_v59;
        _100_v59 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lmjjoibd"), _75_v37), _75_v37, _75_v37, _75_v37);
        _100_v59 = ((_100_v59).Intersect(_100_v59)).Difference(_100_v59);
        if ((new BigNumber(106)).isLessThanOrEqualTo(new BigNumber((_75_v37).length))) {
          let _101_v60;
          _101_v60 = _dafny.Seq.of(p0);
          let _102_v61;
          _102_v61 = new _dafny.CodePoint('d'.codePointAt(0));
          let _103_v62;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_module.__default.fm0(p0, _74_v36, _101_v60, _102_v61, globalState), new BigNumber((_75_v37).length), _74_v36);
          _103_v62 = _nw14;
          let _104_v63;
          _104_v63 = _dafny.MultiSet.fromElements(_103_v62, _103_v62);
          _74_v36 = (((_74_v36) ? (_104_v63) : (_104_v63))).IsProperSubsetOf(_104_v63);
          let _105_v64;
          let _nw15 = new _module.C0();
          _nw15.__ctor(new BigNumber(520), p0, _74_v36);
          _105_v64 = _nw15;
          let _106_v65;
          _106_v65 = _dafny.Seq.of(_105_v64);
          let _107_v66;
          let _nw16 = Array((new BigNumber(16)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _module.__default.fm0(p0, _74_v36, _dafny.Seq.of(new BigNumber((p2).length)), _102_v61, globalState);
          _nw16[(_dafny.ONE).toNumber()] = new BigNumber(483);
          _nw16[(new BigNumber(2)).toNumber()] = new BigNumber((_106_v65).length);
          _nw16[(new BigNumber(3)).toNumber()] = p0;
          _nw16[(new BigNumber(4)).toNumber()] = (_105_v64).f9;
          _nw16[(new BigNumber(5)).toNumber()] = (_105_v64).f8;
          _nw16[(new BigNumber(6)).toNumber()] = (_105_v64).f9;
          _nw16[(new BigNumber(7)).toNumber()] = (_105_v64).f8;
          _nw16[(new BigNumber(8)).toNumber()] = new BigNumber(-166);
          _nw16[(new BigNumber(9)).toNumber()] = (_105_v64).f9;
          _nw16[(new BigNumber(10)).toNumber()] = (_105_v64).f9;
          _nw16[(new BigNumber(11)).toNumber()] = (_105_v64).f8;
          _nw16[(new BigNumber(12)).toNumber()] = (_105_v64).f9;
          _nw16[(new BigNumber(13)).toNumber()] = (_105_v64).f8;
          _nw16[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_101_v60)[_module.__default.safeIndex(p0, new BigNumber((_101_v60).length))]);
          _nw16[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_105_v64).f8);
          _107_v66 = _nw16;
          let _108_v67;
          _108_v67 = _module.D7.create_DC20(_74_v36, _103_v62.f7, _105_v64);
          let _109_v68;
          _109_v68 = _dafny.Map.Empty.slice().updateUnsafe(_107_v66,((_74_v36) ? (_108_v67) : (_108_v67)));
          let _110_v69;
          _110_v69 = _dafny.Map.Empty.slice().updateUnsafe(_74_v36,_107_v66);
          let _111_v70;
          _111_v70 = _dafny.Seq.of(_109_v68);
          _109_v68 = (_dafny.Map.Empty.slice().updateUnsafe((((_110_v69).contains(_74_v36)) ? ((_110_v69).get(_74_v36)) : (_107_v66)),_108_v67)).Merge((_111_v70)[_module.__default.safeIndex((_105_v64).f8, new BigNumber((_111_v70).length))]);
          let _112_v71;
          _112_v71 = _dafny.Map.Empty.slice().updateUnsafe(_102_v61,_103_v62);
          let _113_v72;
          _113_v72 = _dafny.Map.Empty.slice().updateUnsafe(_103_v62.f7,(_105_v64).f9);
          let _114_v73;
          _114_v73 = _dafny.MultiSet.fromElements(_113_v72, _113_v72);
          let _115_v74;
          let _nw17 = Array((new BigNumber(15)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = (_112_v71).update(_module.__default.fm11(_114_v73, p0, _75_v37, _dafny.Seq.UnicodeFromString("aicyyk"), globalState), _103_v62);
          _nw17[(_dafny.ONE).toNumber()] = _112_v71;
          _nw17[(new BigNumber(2)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(3)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(4)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_102_v61,_103_v62);
          _nw17[(new BigNumber(6)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(7)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(8)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(9)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(10)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(11)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(12)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(13)).toNumber()] = _112_v71;
          _nw17[(new BigNumber(14)).toNumber()] = (_112_v71).update(_102_v61, _103_v62);
          _115_v74 = _nw17;
          let _116_v75;
          _116_v75 = _dafny.Set.fromElements(_101_v60, _101_v60);
          let _117_v76;
          _117_v76 = _dafny.Map.Empty.slice().updateUnsafe(_115_v74,_116_v75);
          let _118_v77;
          _118_v77 = _module.D8.create_DC22(_116_v75);
          _117_v76 = (_117_v76).update(_115_v74, ((_118_v77).dtor_cf34).Intersect(_116_v75));
          let _119_v78;
          let _init3 = ((_120_v62) => function (_121_i12) {
            return _120_v62.f7;
          })(_103_v62);
          let _nw18 = Array((new BigNumber(28)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw18.length); _i0_3++) {
            _nw18[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _119_v78 = _nw18;
          let _122_v79;
          _122_v79 = _dafny.Map.Empty.slice().updateUnsafe(_75_v37,_119_v78);
          _122_v79 = ((_122_v79).Merge(_122_v79)).Merge(_122_v79);
          let _123_v80;
          _123_v80 = _dafny.Seq.of(_119_v78, _119_v78);
          _123_v80 = _123_v80;
        } else {
          let _124_v81;
          _124_v81 = _dafny.Map.Empty.slice().updateUnsafe(p0,_74_v36);
          let _125_v82;
          _125_v82 = _dafny.Seq.of(_124_v81, _module.__default.fm15(true, _74_v36, p0, globalState));
          _74_v36 = !_dafny.areEqual(_dafny.Seq.of(_124_v81), _dafny.Seq.Concat(_dafny.Seq.of(_124_v81, _124_v81, _module.__default.fm15(_74_v36, _module.__default.fm3(p0, globalState), new BigNumber(137), globalState), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_74_v36)), _125_v82));
          let _126_v83;
          let _nw19 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _126_v83 = _nw19;
          let _127_v84;
          _127_v84 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index17 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_126_v83).length));
          (_126_v83)[_index17] = _127_v84;
          let _index18 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_126_v83).length));
          (_126_v83)[_index18] = _127_v84;
          let _128_v85;
          _128_v85 = _dafny.Map.Empty.slice().updateUnsafe(_74_v36,false);
          _128_v85 = (_128_v85).update(true, (_74_v36) || (_74_v36));
          let _129_v86;
          _129_v86 = _dafny.Set.fromElements(p0, new BigNumber(-138));
          let _130_v88;
          _130_v88 = _dafny.MultiSet.fromElements((_129_v86).Intersect(function () {
            let _coll5 = new _dafny.Set();
            for (const _compr_5 of (_129_v86).Elements) {
              let _131_v87 = _compr_5;
              if ((_129_v86).contains(_131_v87)) {
                _coll5.add(_module.__default.safeDivisionInt(_131_v87, new BigNumber(948)));
              }
            }
            return _coll5;
          }()));
          let _132_v90;
          _132_v90 = _dafny.Set.fromElements(_74_v36);
          let _133_v91;
          _133_v91 = _dafny.MultiSet.fromElements(_132_v90);
          let _134_v92;
          _134_v92 = _module.D9.create_DC27(_133_v91);
          let _135_v93;
          _135_v93 = _dafny.Seq.of(_133_v91, _133_v91, (_134_v92).dtor_cf43, _133_v91, _133_v91);
          let _136_v94;
          _136_v94 = _dafny.MultiSet.fromElements(p0);
          let _137_v95;
          _137_v95 = _dafny.Map.Empty.slice().updateUnsafe((((_136_v94).contains(new BigNumber(123))) ? ((_136_v94).get(new BigNumber(123))) : (p0)),p0);
          let _rhs17 = _74_v36;
          let _rhs18 = _dafny.Seq.UnicodeFromString("uownf");
          let _rhs19 = (((_130_v88).contains((_129_v86).Union(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(113), new BigNumber(926))) {
              let _139_v89 = _compr_7;
              if (((new BigNumber(113)).isLessThanOrEqualTo(_139_v89)) && ((_139_v89).isLessThan(new BigNumber(926)))) {
                _coll7.add((_139_v89).minus(p0));
              }
            }
            return _coll7;
          }()))) ? ((_130_v88).get((_129_v86).Union(function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(113), new BigNumber(926))) {
              let _138_v89 = _compr_6;
              if (((new BigNumber(113)).isLessThanOrEqualTo(_138_v89)) && ((_138_v89).isLessThan(new BigNumber(926)))) {
                _coll6.add((_138_v89).minus(p0));
              }
            }
            return _coll6;
          }()))) : (new BigNumber(((_135_v93)[_module.__default.safeIndex(new BigNumber((_137_v95).length), new BigNumber((_135_v93).length))]).cardinality())));
          let _lhs15 = globalState;
          _74_v36 = _rhs17;
          r0 = _rhs18;
          _lhs15.f6 = _rhs19;
          let _140_v96;
          let _nw20 = new _module.C0();
          _nw20.__ctor(new BigNumber((_module.__default.fm16(_74_v36, new BigNumber((_75_v37).length), _74_v36, _74_v36, globalState)).length), p0, _74_v36);
          _140_v96 = _nw20;
          let _141_v97;
          _141_v97 = _module.D6.create_DC18(_140_v96);
          let _142_v98;
          let _nw21 = Array((new BigNumber(6)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _141_v97;
          _nw21[(_dafny.ONE).toNumber()] = _141_v97;
          _nw21[(new BigNumber(2)).toNumber()] = _module.D6.create_DC18(_140_v96);
          _nw21[(new BigNumber(3)).toNumber()] = _141_v97;
          _nw21[(new BigNumber(4)).toNumber()] = _141_v97;
          _nw21[(new BigNumber(5)).toNumber()] = _141_v97;
          _142_v98 = _nw21;
          let _index19 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_142_v98).length));
          (_142_v98)[_index19] = _module.D6.create_DC18(_140_v96);
          let _pat_let_tv0 = _140_v96;
          let _index20 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_142_v98).length));
          (_142_v98)[_index20] = function (_pat_let0_0) {
            return function (_143_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_144_dt__update_hcf28_h0) {
                  return _module.D6.create_DC18(_144_dt__update_hcf28_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_141_v97);
        }
        let _145_v99;
        let _nw22 = new _module.C0();
        _nw22.__ctor((_dafny.ZERO).minus(p0), p0, _74_v36);
        _145_v99 = _nw22;
        let _146_v100;
        let _nw23 = new _module.C0();
        _nw23.__ctor(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_145_v99, _145_v99)).length), (((p1).contains(_145_v99.f7)) ? ((p1).get(_145_v99.f7)) : (p0))), p0, _145_v99.f7);
        _146_v100 = _nw23;
        (globalState).f6 = (_146_v100).f8;
        (_145_v99).f7 = _74_v36;
      } else {
        (globalState).f3 = p0;
        _74_v36 = _74_v36;
        let _147_v101;
        let _nw24 = new _module.C0();
        _nw24.__ctor(new BigNumber((p1).cardinality()), p0, _module.__default.fm3(p0, globalState));
        _147_v101 = _nw24;
        let _148_v102;
        _148_v102 = _dafny.Set.fromElements(_147_v101);
        if (!(_148_v102).contains(_147_v101)) {
          (globalState).f3 = (p0).minus(p0);
          let _149_v103;
          let _init4 = ((_150_v101) => function (_151_i13) {
            return _150_v101.f7;
          })(_147_v101);
          let _nw25 = Array((new BigNumber(28)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw25.length); _i0_4++) {
            _nw25[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _149_v103 = _nw25;
          let _152_v104;
          _152_v104 = _dafny.MultiSet.fromElements(_149_v103);
          _152_v104 = (_dafny.MultiSet.fromElements(_149_v103)).Union(_152_v104);
          let _153_v105;
          let _nw26 = new _module.C0();
          _nw26.__ctor(p0, (_dafny.ZERO).minus(p0), _74_v36);
          _153_v105 = _nw26;
          let _154_v106;
          _154_v106 = _module.D7.create_DC20(true, false, _153_v105);
          _153_v105 = (_154_v106).dtor_cf32;
          let _155_v107;
          let _nw27 = new _module.C0();
          _nw27.__ctor(_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_147_v101)).length)), _module.__default.safeDivisionInt((_153_v105).f9, (_153_v105).f8), _module.__default.fm3((_153_v105).f8, globalState));
          _155_v107 = _nw27;
          _153_v105 = _153_v105;
        } else {
          (_147_v101).f7 = (p2)[_module.__default.safeIndex(p0, new BigNumber((p2).length))];
          let _156_v108;
          let _nw28 = Array((new BigNumber(27)).toNumber()).fill(false);
          _156_v108 = _nw28;
          let _index21 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_156_v108).length));
          (_156_v108)[_index21] = false;
          let _157_v109;
          _157_v109 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _index22 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_156_v108).length));
          (_156_v108)[_index22] = !(p1).equals(((!(_147_v101.f7)) ? ((((_157_v109).contains(p0)) ? ((_157_v109).get(p0)) : (p1))) : (p1)));
          let _158_v110;
          _158_v110 = _dafny.Map.Empty.slice().updateUnsafe(true,(_156_v108)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_156_v108).length))]);
          let _159_v111;
          _159_v111 = _dafny.Seq.of(new BigNumber((_158_v110).length));
          let _160_v112;
          _160_v112 = _dafny.Seq.of(new BigNumber((_159_v111).length));
          let _161_v113;
          _161_v113 = _module.D8.create_DC22(_dafny.Set.fromElements(_159_v111));
          let _162_v114;
          _162_v114 = _dafny.Set.fromElements(_160_v112, _159_v111);
          _161_v113 = _module.D8.create_DC22(_162_v114);
          let _163_v115;
          let _nw29 = new _module.C0();
          _nw29.__ctor((_147_v101).fm4(globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,p2)).length), (_156_v108)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_156_v108).length))]);
          _163_v115 = _nw29;
          let _164_v116;
          _164_v116 = _dafny.Map.Empty.slice().updateUnsafe(_163_v115,p0);
          _164_v116 = (_164_v116).Merge(_164_v116);
          _147_v101 = _147_v101;
        }
        let _165_v117;
        let _nw30 = new _module.C0();
        _nw30.__ctor(new BigNumber((_75_v37).length), p0, _147_v101.f7);
        _165_v117 = _nw30;
        let _166_v118;
        _166_v118 = _module.D4.create_DC14(_165_v117, _147_v101.f7);
        let _167_v119;
        let _nw31 = new _module.C0();
        _nw31.__ctor(p0, (p0).minus(p0), (_166_v118).dtor_cf22);
        _167_v119 = _nw31;
        let _168_v120;
        _168_v120 = _module.D0.create_DC0(_74_v36);
        let _pat_let_tv1 = _147_v101;
        (_147_v101).f7 = (function (_pat_let2_0) {
          return function (_169_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_170_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_170_dt__update_hcf0_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1.f7);
          }(_pat_let2_0);
        }(_168_v120)).dtor_cf0;
      }
      let _171_v121;
      let _init5 = function (_172_i14) {
        return (_172_i14).multipliedBy(new BigNumber(-771));
      };
      let _nw32 = Array((new BigNumber(7)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw32.length); _i0_5++) {
        _nw32[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _171_v121 = _nw32;
      let _173_v122;
      _173_v122 = _dafny.Map.Empty.slice().updateUnsafe(_74_v36,_171_v121);
      _173_v122 = (_173_v122).update(_module.__default.fm3(p0, globalState), _171_v121);
      let _174_v123;
      _174_v123 = _dafny.MultiSet.fromElements(p2);
      let _175_v124;
      _175_v124 = _module.D0.create_DC1(_dafny.MultiSet.fromElements(p2, p2), (_dafny.ZERO).minus(p0));
      let _176_v125;
      _176_v125 = _dafny.MultiSet.fromElements(_175_v124);
      let _hi2 = new BigNumber(((_dafny.MultiSet.fromElements(_module.D0.create_DC1(_174_v123, p0))).Union(_176_v125)).cardinality());
      for (let _177_i15 = p0; _177_i15.isLessThan(_hi2); _177_i15 = _177_i15.plus(_dafny.ONE)) {
        let _178_v126;
        _178_v126 = _dafny.Seq.of(p0);
        let _source1 = _module.D1.create_DC3(_dafny.Seq.Concat(_178_v126, _178_v126));
        if (_source1.is_DC4) {
          r0 = _75_v37;
          let _179_v127;
          let _nw33 = Array((new BigNumber(27)).toNumber());
          _179_v127 = _nw33;
          let _180_v128;
          let _nw34 = new _module.C0();
          _nw34.__ctor(_177_i15, (_178_v126)[_module.__default.safeIndex(_177_i15, new BigNumber((_178_v126).length))], _74_v36);
          _180_v128 = _nw34;
          let _181_v129;
          _181_v129 = _dafny.Map.Empty.slice().updateUnsafe(_180_v128,_180_v128);
          let _182_v130;
          _182_v130 = _dafny.Map.Empty.slice().updateUnsafe(_74_v36,_181_v129);
          let _183_v131;
          _183_v131 = new _dafny.CodePoint('q'.codePointAt(0));
          let _184_v132;
          _184_v132 = _dafny.MultiSet.fromElements(new BigNumber(480), _module.__default.fm0(new BigNumber(((((_182_v130).contains(_74_v36)) ? ((_182_v130).get(_74_v36)) : (_181_v129))).length), _74_v36, _dafny.Seq.Create(_module.__default.abs(new BigNumber(927)), ((_185_v128) => function (_186_i16) {
            return (_185_v128).f9;
          })(_180_v128)), _183_v131, globalState), (_dafny.ZERO).minus(new BigNumber((_178_v126).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(474)), ((_187_v131) => function (_188_i17) {
            return _187_v131;
          })(_183_v131))).length));
          let _189_v133;
          let _nw35 = new _module.C0();
          _nw35.__ctor(p0, new BigNumber((_184_v132).cardinality()), _74_v36);
          _189_v133 = _nw35;
          let _index23 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_179_v127).length));
          (_179_v127)[_index23] = _180_v128;
          let _index24 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_179_v127).length));
          (_179_v127)[_index24] = _189_v133;
          (globalState).f6 = ((_180_v128).f9).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(-448), _177_i15));
          (globalState).f3 = (_189_v133).f9;
        } else if (_source1.is_DC3) {
          let _190___mcc_h0 = (_source1).cf4;
          let _191_cf4 = _190___mcc_h0;
          let _192_v134;
          _192_v134 = _75_v37;
          let _193_v135;
          _193_v135 = _dafny.MultiSet.fromElements(_191_cf4, _191_cf4);
          let _194_v136;
          _194_v136 = new _dafny.CodePoint('t'.codePointAt(0));
          let _195_v137;
          let _nw36 = Array((new BigNumber(29)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _75_v37;
          _nw36[(_dafny.ONE).toNumber()] = _75_v37;
          _nw36[(new BigNumber(2)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-246)), function (_196_i18) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          });
          _nw36[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("iqglrvh");
          _nw36[(new BigNumber(5)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(6)).toNumber()] = ((_74_v36) ? (_75_v37) : (_75_v37));
          _nw36[(new BigNumber(7)).toNumber()] = (_192_v134);
          _nw36[(new BigNumber(8)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(9)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_75_v37, _75_v37);
          _nw36[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_75_v37, _75_v37);
          _nw36[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("ugemjeiny");
          _nw36[(new BigNumber(13)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(14)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(15)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(16)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(17)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(18)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm10(p0, _74_v36, _177_i15, _193_v135, globalState), _75_v37);
          _nw36[(new BigNumber(20)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("foi"), _75_v37);
          _nw36[(new BigNumber(22)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ptuxtb"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("ptuxtb")).length)), _194_v136);
          _nw36[(new BigNumber(23)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(24)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(25)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(26)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(27)).toNumber()] = _75_v37;
          _nw36[(new BigNumber(28)).toNumber()] = _75_v37;
          _195_v137 = _nw36;
          let _index25 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_195_v137).length));
          (_195_v137)[_index25] = _75_v37;
          let _index26 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_195_v137).length));
          (_195_v137)[_index26] = _dafny.Seq.UnicodeFromString("jquu");
          let _index27 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_171_v121).length));
          (_171_v121)[_index27] = p0;
          let _index28 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_171_v121).length));
          (_171_v121)[_index28] = _177_i15;
          let _197_v138;
          _197_v138 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_171_v121)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_171_v121).length))]);
          let _198_v139;
          _198_v139 = _dafny.Map.Empty.slice().updateUnsafe(_197_v138,_194_v136);
          _194_v136 = (((_198_v139).contains(function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_module.__default.fm17(_74_v36, _74_v36, _74_v36, globalState)).Elements) {
              let _200_v140 = _compr_9;
              if (_dafny.Seq.contains(_module.__default.fm17(_74_v36, _74_v36, _74_v36, globalState), _200_v140)) {
                _coll9.push([_200_v140,_177_i15]);
              }
            }
            return _coll9;
          }())) ? ((_198_v139).get(function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_module.__default.fm17(_74_v36, _74_v36, _74_v36, globalState)).Elements) {
              let _199_v140 = _compr_8;
              if (_dafny.Seq.contains(_module.__default.fm17(_74_v36, _74_v36, _74_v36, globalState), _199_v140)) {
                _coll8.push([_199_v140,_177_i15]);
              }
            }
            return _coll8;
          }())) : (new _dafny.CodePoint('g'.codePointAt(0))));
          let _201_v141;
          let _nw37 = new _module.C0();
          _nw37.__ctor(new BigNumber((p2).length), _177_i15, false);
          _201_v141 = _nw37;
        } else {
          let _202___mcc_h1 = (_source1).cf5;
          let _203_cf5 = _202___mcc_h1;
          _74_v36 = _74_v36;
          _74_v36 = !(_dafny.MultiSet.fromElements(_75_v37, _75_v37, _75_v37, _75_v37, _75_v37)).contains(_75_v37);
          let _204_v142;
          _204_v142 = _dafny.MultiSet.fromElements(_74_v36);
          _204_v142 = p1;
          let _205_v143;
          _205_v143 = _dafny.Map.Empty.slice().updateUnsafe(_74_v36,_177_i15);
          _204_v142 = ((_module.__default.fm12(_177_i15, _177_i15, _205_v143, globalState)).update(_74_v36, _module.__default.abs(_177_i15))).Union(p1);
        }
        let _206_v144;
        _206_v144 = _dafny.Set.fromElements(_177_i15);
        let _207_v145;
        _207_v145 = _dafny.MultiSet.fromElements(_206_v144);
        let _208_v146;
        let _nw38 = new _module.C0();
        _nw38.__ctor(((((_207_v145).contains(_206_v144)) ? ((_207_v145).get(_206_v144)) : (p0))).plus(_177_i15), (_dafny.ZERO).minus(_177_i15), _74_v36);
        _208_v146 = _nw38;
        let _209_v147;
        _209_v147 = _dafny.Map.Empty.slice().updateUnsafe(_177_i15,_208_v146);
        _208_v146 = (((_209_v147).contains(_177_i15)) ? ((_209_v147).get(_177_i15)) : (_208_v146));
        let _210_v148;
        let _nw39 = Array((new BigNumber(26)).toNumber()).fill([]);
        _210_v148 = _nw39;
        let _index29 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_210_v148).length));
        (_210_v148)[_index29] = _171_v121;
        let _index30 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_210_v148).length));
        let _nw40 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        (_210_v148)[_index30] = _nw40;
        (_208_v146).f7 = _208_v146.f7;
      }
      r0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_75_v37, _75_v37), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("oc"), _75_v37));
      return r0;
    }
    static Main(__noArgsParameter) {
      let _211_v0;
      _211_v0 = _dafny.Seq.UnicodeFromString("qfxbbb");
      let _212_v1;
      _212_v1 = new BigNumber(779);
      let _213_v2;
      _213_v2 = new _dafny.CodePoint('p'.codePointAt(0));
      let _214_v3;
      _214_v3 = _dafny.Map.Empty.slice().updateUnsafe(_212_v1,_213_v2);
      let _215_v4;
      _215_v4 = false;
      let _216_v5;
      _216_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_214_v3).contains(new BigNumber((_211_v0).length))) ? ((_214_v3).get(new BigNumber((_211_v0).length))) : (_213_v2)),_215_v4);
      let _217_v6;
      _217_v6 = _dafny.MultiSet.fromElements(new BigNumber(-166));
      let _218_v7;
      _218_v7 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_217_v6).cardinality())));
      let _219_globalState;
      let _nw41 = new _module.GlobalState();
      _nw41.__ctor(true, _211_v0, true, new BigNumber(225), (_216_v5).Merge(_216_v5), _218_v7, new BigNumber(852));
      _219_globalState = _nw41;
      (_219_globalState).f6 = new BigNumber(155);
      let _hi3 = _212_v1;
      for (let _220_i0 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_212_v1), _module.__default.safeIndex(new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_211_v0).Elements) {
          let _253_v8 = _compr_10;
          if (_dafny.Seq.contains(_211_v0, _253_v8)) {
            _coll10.push([_253_v8,_212_v1]);
          }
        }
        return _coll10;
      }()).length), new BigNumber((_dafny.Seq.of(_212_v1)).length)), _212_v1)).length)).plus(new BigNumber(83)); _220_i0.isLessThan(_hi3); _220_i0 = _220_i0.plus(_dafny.ONE)) {
        let _221_v9;
        _221_v9 = _dafny.Seq.of(_220_i0, new BigNumber(832), _212_v1, _212_v1);
        if (_dafny.Seq.IsProperPrefixOf(_221_v9, _221_v9)) {
          let _222_v10;
          let _init6 = ((_223_v6) => function (_224_i1) {
            return _223_v6;
          })(_217_v6);
          let _nw42 = Array((new BigNumber(8)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw42.length); _i0_6++) {
            _nw42[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _222_v10 = _nw42;
          _222_v10 = _222_v10;
          let _225_v11;
          _225_v11 = _dafny.MultiSet.fromElements(_215_v4, !(_215_v4));
          let _226_v12;
          let _out0;
          _out0 = _module.__default.m0(_220_i0, (_225_v11).Intersect(_225_v11), _dafny.Seq.update(_dafny.Seq.of(_215_v4), _module.__default.safeIndex(new BigNumber((_221_v9).length), new BigNumber((_dafny.Seq.of(_215_v4)).length)), _215_v4), _219_globalState);
          _226_v12 = _out0;
          let _rhs20 = ((_215_v4) ? (_220_i0) : (new BigNumber((_211_v0).length)));
          let _rhs21 = _220_i0;
          let _lhs16 = _219_globalState;
          let _lhs17 = _219_globalState;
          _lhs16.f6 = _rhs20;
          _lhs17.f3 = _rhs21;
          _215_v4 = _215_v4;
          let _227_v13;
          _227_v13 = _module.D0.create_DC0(_215_v4);
          _215_v4 = (_227_v13).dtor_cf0;
        } else {
          let _228_v14;
          _228_v14 = _dafny.Set.fromElements(_220_i0, _212_v1);
          let _rhs22 = _module.__default.fm0((_module.__default.fm1(_219_globalState)).dtor_cf2, true, _221_v9, _213_v2, _219_globalState);
          let _rhs23 = _module.__default.safeDivisionInt(_212_v1, new BigNumber((_228_v14).length));
          let _lhs18 = _219_globalState;
          let _lhs19 = _219_globalState;
          _lhs18.f3 = _rhs22;
          _lhs19.f3 = _rhs23;
          let _229_v15;
          _229_v15 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,_dafny.MultiSet.fromElements(new BigNumber(928), _212_v1));
          let _230_v16;
          _230_v16 = _dafny.Map.Empty.slice().updateUnsafe((((_229_v15).contains(_215_v4)) ? ((_229_v15).get(_215_v4)) : (_217_v6)),false);
          _228_v14 = _module.__default.fm2(_230_v16, new _dafny.CodePoint('d'.codePointAt(0)), _219_globalState);
          let _231_v17;
          _231_v17 = _dafny.Map.Empty.slice().updateUnsafe(_220_i0,_212_v1);
          let _232_v18;
          _232_v18 = _dafny.Seq.of(_215_v4, _215_v4);
          let _233_v19;
          _233_v19 = _dafny.MultiSet.fromElements(_232_v18);
          let _234_v20;
          _234_v20 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,false);
          let _235_v21;
          _235_v21 = _module.D0.create_DC1(_233_v19, new BigNumber((_234_v20).length));
          let _236_v22;
          _236_v22 = _dafny.Map.Empty.slice().updateUnsafe(_211_v0,_215_v4);
          let _237_v23;
          _237_v23 = _dafny.Map.Empty.slice().updateUnsafe(_236_v22,_215_v4);
          let _238_v24;
          _238_v24 = _dafny.Set.fromElements(_215_v4);
          let _rhs24 = ((_module.__default.fm3(_212_v1, _219_globalState)) ? (((_215_v4) ? (_212_v1) : (_220_i0))) : ((_235_v21).dtor_cf2));
          let _rhs25 = _231_v17;
          let _rhs26 = (((_237_v23).contains(_236_v22)) ? ((_237_v23).get(_236_v22)) : ((_238_v24).IsProperSubsetOf(_238_v24)));
          let _rhs27 = _215_v4;
          let _lhs20 = _219_globalState;
          _lhs20.f3 = _rhs24;
          _231_v17 = _rhs25;
          _215_v4 = _rhs26;
          _215_v4 = _rhs27;
          _215_v4 = _215_v4;
          let _239_v25;
          let _nw43 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _239_v25 = _nw43;
          let _index31 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_239_v25).length));
          (_239_v25)[_index31] = _220_i0;
          let _index32 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_239_v25).length));
          (_239_v25)[_index32] = _212_v1;
        }
        let _240_v26;
        _240_v26 = _dafny.Seq.of(_215_v4);
        let _241_v27;
        _241_v27 = _module.D0.create_DC0(false);
        let _242_v28;
        let _nw44 = Array((new BigNumber(10)).toNumber());
        _nw44[(_dafny.ZERO).toNumber()] = _215_v4;
        _nw44[(_dafny.ONE).toNumber()] = (_240_v26)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_240_v26).length))];
        _nw44[(new BigNumber(2)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(3)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(4)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(5)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(6)).toNumber()] = true;
        _nw44[(new BigNumber(7)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(8)).toNumber()] = _215_v4;
        _nw44[(new BigNumber(9)).toNumber()] = (_241_v27).dtor_cf0;
        _242_v28 = _nw44;
        let _index33 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_242_v28).length));
        (_242_v28)[_index33] = _215_v4;
        let _index34 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_242_v28).length));
        (_242_v28)[_index34] = _215_v4;
        let _243_v29;
        _243_v29 = _dafny.Set.fromElements(_213_v2);
        let _244_v30;
        _244_v30 = _dafny.MultiSet.fromElements((_243_v29).Intersect(_243_v29), _243_v29);
        let _245_v31;
        _245_v31 = _dafny.Seq.of(_243_v29);
        _244_v30 = ((_dafny.MultiSet.FromArray(_245_v31)).Union(_244_v30)).update(_243_v29, _module.__default.abs(_220_i0));
        let _246_v32;
        _246_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,_220_i0);
        let _247_v33;
        _247_v33 = _dafny.Seq.of(_242_v28, _242_v28);
        let _248_v34;
        _248_v34 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_212_v1)), _217_v6);
        let _249_v35;
        _249_v35 = _module.D1.create_DC3(_221_v9);
        let _250_v36;
        let _nw45 = Array((new BigNumber(21)).toNumber());
        _nw45[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-332)), function (_251_i2) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
        }), _module.__default.safeIndex(new BigNumber((_211_v0).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-332)), function (_252_i2) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
        })).length)), new BigNumber((_246_v32).length))).length);
        _nw45[(_dafny.ONE).toNumber()] = new BigNumber(742);
        _nw45[(new BigNumber(2)).toNumber()] = new BigNumber((_246_v32).length);
        _nw45[(new BigNumber(3)).toNumber()] = _212_v1;
        _nw45[(new BigNumber(4)).toNumber()] = _220_i0;
        _nw45[(new BigNumber(5)).toNumber()] = _212_v1;
        _nw45[(new BigNumber(6)).toNumber()] = _220_i0;
        _nw45[(new BigNumber(7)).toNumber()] = new BigNumber(967);
        _nw45[(new BigNumber(8)).toNumber()] = _220_i0;
        _nw45[(new BigNumber(9)).toNumber()] = new BigNumber(568);
        _nw45[(new BigNumber(10)).toNumber()] = new BigNumber((_247_v33).length);
        _nw45[(new BigNumber(11)).toNumber()] = _212_v1;
        _nw45[(new BigNumber(12)).toNumber()] = new BigNumber(869);
        _nw45[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(_212_v1, _220_i0);
        _nw45[(new BigNumber(14)).toNumber()] = _212_v1;
        _nw45[(new BigNumber(15)).toNumber()] = (_220_i0).minus(_212_v1);
        _nw45[(new BigNumber(16)).toNumber()] = _212_v1;
        _nw45[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm0(new BigNumber((_248_v34).length), (_242_v28)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_242_v28).length))], _221_v9, _213_v2, _219_globalState), _module.__default.fm0(new BigNumber((_211_v0).length), (_242_v28)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_242_v28).length))], (_249_v35).dtor_cf4, _213_v2, _219_globalState));
        _nw45[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_220_i0);
        _nw45[(new BigNumber(19)).toNumber()] = new BigNumber(975);
        _nw45[(new BigNumber(20)).toNumber()] = ((_dafny.ZERO).minus(_220_i0)).plus(_220_i0);
        _250_v36 = _nw45;
        let _index35 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_250_v36).length));
        (_250_v36)[_index35] = _212_v1;
        let _index36 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_250_v36).length));
        let _rhs28 = _220_i0;
        let _rhs29 = _220_i0;
        let _rhs30 = _220_i0;
        let _lhs21 = _250_v36;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_250_v36).length));
        let _lhs23 = _219_globalState;
        _lhs21[_lhs22] = _rhs28;
        _212_v1 = _rhs29;
        _lhs23.f6 = _rhs30;
      }
      let _254_v37;
      _254_v37 = _dafny.Set.fromElements(_213_v2);
      let _255_v38;
      _255_v38 = _dafny.Set.fromElements(_254_v37, _254_v37, _254_v37, _dafny.Set.fromElements(_213_v2, _213_v2), _254_v37);
      let _256_v39;
      _256_v39 = _dafny.Seq.of(_255_v38);
      let _257_v40;
      _257_v40 = _dafny.Map.Empty.slice().updateUnsafe(_212_v1,_212_v1);
      let _hi4 = (_212_v1).minus(new BigNumber((_257_v40).length));
      for (let _258_i3 = new BigNumber(((_256_v39)[_module.__default.safeIndex(_212_v1, new BigNumber((_256_v39).length))]).length); _258_i3.isLessThan(_hi4); _258_i3 = _258_i3.plus(_dafny.ONE)) {
        if (_215_v4) {
          let _259_v41;
          let _nw46 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _259_v41 = _nw46;
          let _260_v42;
          let _init7 = ((_261_v1) => function (_262_i4) {
            return _module.__default.safeModuloInt(_262_i4, _261_v1);
          })(_212_v1);
          let _nw47 = Array((new BigNumber(18)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw47.length); _i0_7++) {
            _nw47[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _260_v42 = _nw47;
          let _263_v43;
          _263_v43 = _dafny.Map.Empty.slice().updateUnsafe(_259_v41,_260_v42);
          _263_v43 = (_263_v43).update(_259_v41, _260_v42);
          let _264_v44;
          let _nw48 = new _module.C0();
          _nw48.__ctor(_212_v1, _module.__default.safeModuloInt(_258_i3, new BigNumber(163)), _215_v4);
          _264_v44 = _nw48;
          let _index37 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_260_v42).length));
          (_260_v42)[_index37] = _module.__default.safeDivisionInt((_264_v44).f8, _258_i3);
          let _index38 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_260_v42).length));
          (_260_v42)[_index38] = (_dafny.ZERO).minus(_212_v1);
          let _265_v45;
          _265_v45 = _dafny.Set.fromElements(_215_v4);
          let _266_v46;
          _266_v46 = _dafny.MultiSet.fromElements(_module.__default.fm3(_212_v1, _219_globalState));
          let _267_v47;
          _267_v47 = _module.D1.create_DC4();
          let _268_v48;
          _268_v48 = _dafny.Map.Empty.slice().updateUnsafe(_267_v47,_215_v4);
          let _269_v49;
          let _out1;
          _out1 = _module.__default.m0((new BigNumber((_265_v45).length)).plus(_258_i3), _266_v46, _dafny.Seq.Concat(_module.__default.fm6((_264_v44).f9, new _dafny.CodePoint('p'.codePointAt(0)), (((_268_v48).contains(_267_v47)) ? ((_268_v48).get(_267_v47)) : (_215_v4)), _215_v4, _219_globalState), _module.__default.fm6(_212_v1, _213_v2, _215_v4, _215_v4, _219_globalState)), _219_globalState);
          _269_v49 = _out1;
          let _270_v50;
          _270_v50 = _dafny.Seq.of(_215_v4, _215_v4);
          let _271_v51;
          _271_v51 = _dafny.Seq.of((_270_v50)[_module.__default.safeIndex((_264_v44).f8, new BigNumber((_270_v50).length))]);
          let _272_v52;
          let _out2;
          _out2 = _module.__default.m0(new BigNumber(496), _266_v46, _271_v51, _219_globalState);
          _272_v52 = _out2;
        } else {
          _215_v4 = (_217_v6).IsDisjointFrom(_217_v6);
          let _273_v53;
          let _init8 = ((_274_v4) => function (_275_i5) {
            return _274_v4;
          })(_215_v4);
          let _nw49 = Array((new BigNumber(26)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw49.length); _i0_8++) {
            _nw49[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _273_v53 = _nw49;
          let _276_v54;
          _276_v54 = _dafny.Map.Empty.slice().updateUnsafe(_258_i3,_273_v53);
          _212_v1 = new BigNumber((((_276_v54).Merge(_dafny.Map.Empty.slice().updateUnsafe(_212_v1,_273_v53))).Merge(_276_v54)).length);
          let _277_v55;
          _277_v55 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,_258_i3);
          let _278_v56;
          _278_v56 = _dafny.Seq.of(!(true));
          let _279_v57;
          _279_v57 = _dafny.Seq.of(_212_v1, _212_v1);
          let _280_v58;
          _280_v58 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_277_v55).length), (_278_v56)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_211_v0).length)), new BigNumber((_278_v56).length))], _279_v57, _213_v2, _219_globalState),_215_v4);
          let _281_v59;
          let _nw50 = new _module.C0();
          _nw50.__ctor((new BigNumber(969)).minus(_258_i3), _module.__default.safeDivisionInt(_212_v1, new BigNumber((_277_v55).length)), (((_280_v58).contains(_212_v1)) ? ((_280_v58).get(_212_v1)) : (_215_v4)));
          _281_v59 = _nw50;
          let _282_v60;
          let _nw51 = Array((new BigNumber(12)).toNumber()).fill(_module.D1.Default());
          _282_v60 = _nw51;
          let _283_v61;
          _283_v61 = _module.D1.create_DC3(_dafny.Seq.of(new BigNumber(146)));
          let _index39 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_282_v60).length));
          (_282_v60)[_index39] = _283_v61;
          let _index40 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_282_v60).length));
          (_282_v60)[_index40] = _283_v61;
          let _284_v62;
          _284_v62 = _dafny.MultiSet.fromElements(_215_v4);
          let _285_v63;
          let _out3;
          _out3 = _module.__default.m0(((_215_v4) ? ((_281_v59).f9) : (new BigNumber((_278_v56).length))), (_dafny.MultiSet.FromArray(_278_v56)).Intersect(_284_v62), _dafny.Seq.update(_278_v56, _module.__default.safeIndex(_258_i3, new BigNumber((_278_v56).length)), _215_v4), _219_globalState);
          _285_v63 = _out3;
        }
        (_219_globalState).f6 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_258_i3), new BigNumber((_211_v0).length)), _212_v1);
        let _286_v64;
        _286_v64 = _dafny.Seq.of(_212_v1);
        let _287_v65;
        let _nw52 = Array((new BigNumber(11)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = true;
        _nw52[(_dafny.ONE).toNumber()] = _215_v4;
        _nw52[(new BigNumber(2)).toNumber()] = _215_v4;
        _nw52[(new BigNumber(3)).toNumber()] = false;
        _nw52[(new BigNumber(4)).toNumber()] = !(_215_v4);
        _nw52[(new BigNumber(5)).toNumber()] = _215_v4;
        _nw52[(new BigNumber(6)).toNumber()] = !(_215_v4);
        _nw52[(new BigNumber(7)).toNumber()] = _215_v4;
        _nw52[(new BigNumber(8)).toNumber()] = _module.__default.fm3((_dafny.ZERO).minus(_258_i3), _219_globalState);
        _nw52[(new BigNumber(9)).toNumber()] = !(false);
        _nw52[(new BigNumber(10)).toNumber()] = _215_v4;
        _287_v65 = _nw52;
        let _288_v66;
        _288_v66 = _dafny.Map.Empty.slice().updateUnsafe((_286_v64)[_module.__default.safeIndex(_258_i3, new BigNumber((_286_v64).length))],_287_v65);
        let _289_v67;
        _289_v67 = _dafny.Set.fromElements(_215_v4);
        _288_v66 = (_288_v66).update(new BigNumber((_289_v67).length), _287_v65);
        if (_215_v4) {
          let _290_v68;
          _290_v68 = _dafny.MultiSet.fromElements(_215_v4);
          let _rhs31 = (_212_v1).multipliedBy(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(_258_i3), _212_v1, _212_v1)).length));
          let _rhs32 = !(_290_v68).equals(_dafny.MultiSet.fromElements(_215_v4));
          let _lhs24 = _219_globalState;
          _lhs24.f3 = _rhs31;
          _215_v4 = _rhs32;
          let _291_v69;
          let _nw53 = new _module.C0();
          _nw53.__ctor(_258_i3, _258_i3, (_215_v4) === (_215_v4));
          _291_v69 = _nw53;
          _215_v4 = !((_module.D0.create_DC0(_215_v4)).dtor_cf0);
          let _292_v70;
          _292_v70 = _dafny.Seq.of(_290_v68);
          let _293_v71;
          _293_v71 = _dafny.Seq.of(_215_v4);
          let _294_v72;
          let _out4;
          _out4 = _module.__default.m0(_module.__default.safeDivisionInt(_258_i3, _module.__default.fm0(_212_v1, _215_v4, _286_v64, _213_v2, _219_globalState)), ((_290_v68).update(_215_v4, _module.__default.abs((_dafny.ZERO).minus(_258_i3)))).Difference((_292_v70)[_module.__default.safeIndex((_291_v69).f9, new BigNumber((_292_v70).length))]), _293_v71, _219_globalState);
          _294_v72 = _out4;
          _215_v4 = _module.__default.fm3((_291_v69).f9, _219_globalState);
        } else {
          _215_v4 = _215_v4;
          let _295_v73;
          _295_v73 = _dafny.Seq.of(_215_v4);
          let _296_v74;
          _296_v74 = _dafny.Seq.of(_215_v4, (_295_v73)[_module.__default.safeIndex(new BigNumber((_217_v6).cardinality()), new BigNumber((_295_v73).length))]);
          let _index41 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_287_v65).length));
          (_287_v65)[_index41] = (_295_v73)[_module.__default.safeIndex(_212_v1, new BigNumber((_295_v73).length))];
          let _index42 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_287_v65).length));
          (_287_v65)[_index42] = _215_v4;
          let _297_v75;
          _297_v75 = _dafny.MultiSet.fromElements(_295_v73);
          let _298_v76;
          _298_v76 = _module.D0.create_DC1(_297_v75, _212_v1);
          let _299_v77;
          _299_v77 = _module.D0.create_DC2(_298_v76);
          let _pat_let_tv2 = _219_globalState;
          _299_v77 = function (_pat_let4_0) {
            return function (_300_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_301_dt__update_hcf3_h0) {
                  return _module.D0.create_DC2(_301_dt__update_hcf3_h0);
                }(_pat_let5_0);
              }(_module.__default.fm1(_pat_let_tv2));
            }(_pat_let4_0);
          }(((_215_v4) ? (_299_v77) : (_299_v77)));
          let _index43 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_287_v65).length));
          (_287_v65)[_index43] = _module.__default.fm3(((_215_v4) ? (_258_i3) : (_258_i3)), _219_globalState);
          (_219_globalState).f3 = _212_v1;
        }
      }
      let _hi5 = _212_v1;
      for (let _302_i6 = _212_v1; _302_i6.isLessThan(_hi5); _302_i6 = _302_i6.plus(_dafny.ONE)) {
        let _303_v78;
        let _nw54 = new _module.C0();
        _nw54.__ctor((_dafny.ZERO).minus((_dafny.ZERO).minus(_212_v1)), _302_i6, _215_v4);
        _303_v78 = _nw54;
        _214_v3 = (_214_v3).update((_303_v78).f9, _213_v2);
        let _304_v79;
        _304_v79 = _dafny.Seq.of(_215_v4, _215_v4, _215_v4);
        let _305_v80;
        let _out5;
        _out5 = _module.__default.m0((_dafny.ZERO).minus((_303_v78).f8), _dafny.MultiSet.fromElements(_215_v4, false), _304_v79, _219_globalState);
        _305_v80 = _out5;
        (_219_globalState).f3 = (_dafny.ZERO).minus(_302_i6);
      }
      let _306_v81;
      _306_v81 = _dafny.Seq.of(_212_v1);
      let _307_v82;
      _307_v82 = _dafny.Seq.of(_212_v1, (_306_v81)[_module.__default.safeIndex(_212_v1, new BigNumber((_306_v81).length))]);
      let _308_v83;
      _308_v83 = _module.D1.create_DC3(_306_v81);
      let _309_i7;
      _309_i7 = _dafny.ZERO;
      L0: {
        let _pat_let_tv3 = _215_v4;
        let _pat_let_tv4 = _215_v4;
        let _pat_let_tv5 = _212_v1;
        while (function (_source2) {
          if (_source2.is_DC4) {
            return _pat_let_tv3;
          } else if (_source2.is_DC3) {
            let _322___mcc_h0 = (_source2).cf4;
            let _323_cf4 = _322___mcc_h0;
            return _pat_let_tv4;
          } else {
            let _324___mcc_h1 = (_source2).cf5;
            let _325_cf5 = _324___mcc_h1;
            return (_pat_let_tv5).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("nol")).length));
          }
        }(_308_v83)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_309_i7)) {
              break L0;
            }
            _309_i7 = (_309_i7).plus(_dafny.ONE);
            let _310_v84;
            let _nw55 = Array((new BigNumber(14)).toNumber()).fill(false);
            _310_v84 = _nw55;
            let _311_v85;
            _311_v85 = _dafny.Seq.of(_215_v4);
            let _index44 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length));
            (_310_v84)[_index44] = !_dafny.areEqual(_311_v85, _311_v85);
            let _312_v86;
            _312_v86 = _module.D0.create_DC0(_215_v4);
            let _313_v87;
            let _nw56 = new _module.C0();
            _nw56.__ctor(_212_v1, new BigNumber(715), (_312_v86).dtor_cf0);
            _313_v87 = _nw56;
            let _314_v88;
            _314_v88 = _dafny.Map.Empty.slice().updateUnsafe(_313_v87,_215_v4);
            let _315_v89;
            _315_v89 = _dafny.MultiSet.fromElements(_dafny.Seq.of(!(_215_v4), _215_v4), _311_v85);
            let _index45 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length));
            (_310_v84)[_index45] = (((_314_v88).contains(_313_v87)) ? ((_314_v88).get(_313_v87)) : ((_315_v89).IsSubsetOf(_315_v89)));
            let _316_v90;
            let _nw57 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _316_v90 = _nw57;
            let _index46 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_316_v90).length));
            (_316_v90)[_index46] = (_313_v87).f9;
            let _317_v91;
            _317_v91 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,(_310_v84)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length))]);
            let _318_v92;
            _318_v92 = _dafny.Set.fromElements(_215_v4, (_312_v86).dtor_cf0, !(_317_v91).equals(_317_v91), _module.__default.fm3((_313_v87).f9, _219_globalState), _dafny.Seq.contains(_311_v85, (_310_v84)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length))]));
            let _index47 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length));
            let _index48 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_316_v90).length));
            let _rhs33 = _215_v4;
            let _rhs34 = (_313_v87).f9;
            let _rhs35 = new BigNumber((_318_v92).length);
            let _lhs25 = _310_v84;
            let _lhs26 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length));
            let _lhs27 = _219_globalState;
            let _lhs28 = _316_v90;
            let _lhs29 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_316_v90).length));
            _lhs25[_lhs26] = _rhs33;
            _lhs27.f6 = _rhs34;
            _lhs28[_lhs29] = _rhs35;
            let _319_v93;
            let _nw58 = Array((new BigNumber(11)).toNumber());
            _319_v93 = _nw58;
            let _nw59 = Array((new BigNumber(27)).toNumber());
            _319_v93 = _nw59;
            let _320_v94;
            _320_v94 = _dafny.MultiSet.fromElements(!(false), !(!(!(_215_v4))), _215_v4, !(true), (_310_v84)[_module.__default.safeIndex(new BigNumber(798), new BigNumber((_310_v84).length))]);
            let _321_v95;
            let _out6;
            _out6 = _module.__default.m0((_dafny.ZERO).minus(_212_v1), _320_v94, _311_v85, _219_globalState);
            _321_v95 = _out6;
          }
        }
      }
      _215_v4 = _215_v4;
      let _326_v96;
      _326_v96 = _module.D0.create_DC0(_module.__default.fm3(_212_v1, _219_globalState));
      if ((_326_v96).dtor_cf0) {
        let _327_v97;
        let _nw60 = Array((new BigNumber(18)).toNumber()).fill(false);
        _327_v97 = _nw60;
        let _index49 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
        (_327_v97)[_index49] = _215_v4;
        let _328_v98;
        _328_v98 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,_212_v1);
        let _329_v99;
        let _nw61 = Array((new BigNumber(4)).toNumber());
        _nw61[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_212_v1);
        _nw61[(_dafny.ONE).toNumber()] = _212_v1;
        _nw61[(new BigNumber(2)).toNumber()] = (((_328_v98).contains(_215_v4)) ? ((_328_v98).get(_215_v4)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_327_v97,_212_v1)).length)));
        _nw61[(new BigNumber(3)).toNumber()] = _212_v1;
        _329_v99 = _nw61;
        let _index50 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
        (_329_v99)[_index50] = new BigNumber(736);
        let _330_v100;
        _330_v100 = _dafny.Seq.of(_215_v4, _215_v4);
        let _331_v101;
        _331_v101 = _dafny.MultiSet.fromElements(_330_v100);
        let _332_v102;
        _332_v102 = _module.D0.create_DC1(_331_v101, _212_v1);
        let _333_v103;
        _333_v103 = _dafny.Map.Empty.slice().updateUnsafe(_307_v82,_212_v1);
        let _index51 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
        let _index52 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
        let _rhs36 = (_332_v102).dtor_cf2;
        let _rhs37 = !(_333_v103).contains(_module.__default.fm7(_219_globalState));
        let _rhs38 = _212_v1;
        let _lhs30 = _219_globalState;
        let _lhs31 = _327_v97;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
        let _lhs33 = _329_v99;
        let _lhs34 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
        _lhs30.f6 = _rhs36;
        _lhs31[_lhs32] = _rhs37;
        _lhs33[_lhs34] = _rhs38;
        let _334_v104;
        let _nw62 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _334_v104 = _nw62;
        let _335_v105;
        let _nw63 = Array((new BigNumber(4)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = _334_v104;
        _nw63[(_dafny.ONE).toNumber()] = _334_v104;
        _nw63[(new BigNumber(2)).toNumber()] = _334_v104;
        _nw63[(new BigNumber(3)).toNumber()] = _334_v104;
        _335_v105 = _nw63;
        let _336_v106;
        _336_v106 = _dafny.Map.Empty.slice().updateUnsafe((_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))],_334_v104);
        let _337_v107;
        _337_v107 = _dafny.MultiSet.fromElements(false, (_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]);
        let _index53 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_335_v105).length));
        (_335_v105)[_index53] = (((_336_v106).contains(new BigNumber((_337_v107).cardinality()))) ? ((_336_v106).get(new BigNumber((_337_v107).cardinality()))) : (_334_v104));
        let _index54 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_335_v105).length));
        let _rhs39 = (_330_v100)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_330_v100).length))];
        let _rhs40 = _334_v104;
        let _lhs35 = _335_v105;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_335_v105).length));
        _215_v4 = _rhs39;
        _lhs35[_lhs36] = _rhs40;
        let _338_v108;
        let _nw64 = new _module.C0();
        _nw64.__ctor((_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))], _212_v1, _215_v4);
        _338_v108 = _nw64;
        let _339_v109;
        _339_v109 = _module.D1.create_DC4();
        let _340_v110;
        _340_v110 = _dafny.Set.fromElements(_339_v109);
        let _index55 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
        let _rhs41 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_338_v108).f9, new BigNumber((_340_v110).length)), (_338_v108).f9);
        let _rhs42 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_306_v81, _307_v82))).IsDisjointFrom(_217_v6);
        let _lhs37 = _329_v99;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
        let _lhs39 = _327_v97;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
        _lhs37[_lhs38] = _rhs41;
        _lhs39[_lhs40] = _rhs42;
        if ((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]) {
          let _341_v111;
          _341_v111 = _dafny.Map.Empty.slice().updateUnsafe(false,(_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]);
          let _342_v112;
          _342_v112 = _dafny.Map.Empty.slice().updateUnsafe(_341_v111,_213_v2);
          _342_v112 = (_342_v112).update(_dafny.Map.Empty.slice().updateUnsafe((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))],_215_v4), _213_v2);
          let _343_v113;
          let _nw65 = new _module.C0();
          _nw65.__ctor((_338_v108).f9, (_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))], _215_v4);
          _343_v113 = _nw65;
          let _344_v114;
          _344_v114 = _dafny.Map.Empty.slice().updateUnsafe(_343_v113,(_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]);
          let _index57 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          (_327_v97)[_index57] = (((_344_v114).contains(_343_v113)) ? ((_344_v114).get(_343_v113)) : (true));
          let _345_v115;
          let _out7;
          _out7 = _module.__default.m0(new BigNumber(907), _dafny.MultiSet.fromElements(_215_v4, (_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]), _330_v100, _219_globalState);
          _345_v115 = _out7;
          let _rhs43 = (_338_v108).f8;
          let _rhs44 = _329_v99;
          let _rhs45 = _341_v111;
          let _rhs46 = (_338_v108).f9;
          let _rhs47 = _338_v108;
          let _lhs41 = _219_globalState;
          let _lhs42 = _219_globalState;
          _lhs41.f6 = _rhs43;
          _329_v99 = _rhs44;
          _341_v111 = _rhs45;
          _lhs42.f3 = _rhs46;
          _338_v108 = _rhs47;
          let _index58 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
          (_329_v99)[_index58] = (_338_v108).f8;
        } else {
          let _index59 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
          let _rhs48 = _215_v4;
          let _rhs49 = (_338_v108).f9;
          let _lhs43 = _327_v97;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          let _lhs45 = _329_v99;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length));
          _lhs43[_lhs44] = _rhs48;
          _lhs45[_lhs46] = _rhs49;
          let _346_v116;
          _346_v116 = _dafny.Map.Empty.slice().updateUnsafe((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))],_215_v4);
          (_219_globalState).f3 = (_module.__default.safeModuloInt(new BigNumber((_346_v116).length), (_307_v82)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_215_v4, _215_v4)).cardinality()), new BigNumber((_307_v82).length))])).multipliedBy(_module.__default.fm0((_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))], (_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))], _dafny.Seq.of((_338_v108).f9, (_338_v108).f8, (_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))]), _213_v2, _219_globalState));
          let _347_v117;
          _347_v117 = _dafny.Seq.of(_211_v0);
          let _index61 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          let _index62 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          let _rhs50 = ((_338_v108).f8).isLessThan((_338_v108).fm5(_module.__default.fm8((_338_v108).f9, (_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))], new BigNumber(356), (_329_v99)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_329_v99).length))], _219_globalState), _213_v2, _212_v1, _215_v4, _219_globalState));
          let _rhs51 = ((!_dafny.areEqual((_347_v117)[_module.__default.safeIndex((_338_v108).f9, new BigNumber((_347_v117).length))], _211_v0)) ? ((((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]) ? (_308_v83) : (_308_v83))) : (_module.__default.fm9(_219_globalState)));
          let _rhs52 = !((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]) || (_215_v4);
          let _lhs47 = _327_v97;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          let _lhs49 = _327_v97;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length));
          _lhs47[_lhs48] = _rhs50;
          _308_v83 = _rhs51;
          _lhs49[_lhs50] = _rhs52;
          _346_v116 = (_346_v116).update((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))], (_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]);
          let _348_v118;
          let _out8;
          _out8 = _module.__default.m0((_338_v108).f8, _dafny.MultiSet.fromElements((_327_v97)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_327_v97).length))]), _330_v100, _219_globalState);
          _348_v118 = _out8;
        }
      } else {
        let _349_v119;
        _349_v119 = _dafny.Seq.of(_215_v4);
        let _350_v120;
        _350_v120 = _dafny.MultiSet.fromElements(_349_v119, _349_v119);
        let _351_v121;
        _351_v121 = _module.D0.create_DC1(_350_v120, _212_v1);
        let _352_v122;
        _352_v122 = _dafny.MultiSet.fromElements(_module.D0.create_DC1(_350_v120, new BigNumber((_349_v119).length)), _351_v121);
        _215_v4 = !((_dafny.MultiSet.FromArray(_dafny.Seq.of(_351_v121))).IsProperSubsetOf(_352_v122));
        let _353_v123;
        _353_v123 = _dafny.MultiSet.fromElements(_215_v4, _215_v4);
        let _354_v124;
        let _out9;
        _out9 = _module.__default.m0(new BigNumber((_211_v0).length), _353_v123, _dafny.Seq.update(_349_v119, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_215_v4)).length), new BigNumber((_349_v119).length)), _215_v4), _219_globalState);
        _354_v124 = _out9;
        let _355_v125;
        _355_v125 = _dafny.Map.Empty.slice().updateUnsafe(_212_v1,false);
        _215_v4 = !((_module.__default.fm3(_module.__default.fm0(_212_v1, _215_v4, _306_v81, _213_v2, _219_globalState), _219_globalState)) && ((((_355_v125).contains(_212_v1)) ? ((_355_v125).get(_212_v1)) : (_215_v4))));
        let _356_v126;
        let _out10;
        _out10 = _module.__default.m0(_212_v1, (_353_v123).Difference(_353_v123), _module.__default.fm6(_212_v1, _213_v2, _215_v4, _215_v4, _219_globalState), _219_globalState);
        _356_v126 = _out10;
        let _357_v127;
        _357_v127 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,(_dafny.ZERO).minus(_212_v1));
        let _358_v128;
        _358_v128 = _dafny.Set.fromElements((((_357_v127).contains(!(_215_v4))) ? ((_357_v127).get(!(_215_v4))) : (_212_v1)));
        _358_v128 = _358_v128;
      }
      let _359_v129;
      _359_v129 = _dafny.MultiSet.fromElements(_215_v4, _215_v4, _215_v4, _215_v4, _module.__default.fm3(_212_v1, _219_globalState));
      let _360_v130;
      _360_v130 = _dafny.Map.Empty.slice().updateUnsafe(_213_v2,_359_v129);
      _360_v130 = (_360_v130).update(_213_v2, _359_v129);
      _215_v4 = _215_v4;
      _215_v4 = (_326_v96).dtor_cf0;
      let _361_v132;
      let _nw66 = new _module.C0();
      _nw66.__ctor(_212_v1, new BigNumber((function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(416), new BigNumber(754))) {
          let _362_v131 = _compr_11;
          if (((new BigNumber(416)).isLessThanOrEqualTo(_362_v131)) && ((_362_v131).isLessThan(new BigNumber(754)))) {
            _coll11.add((_362_v131).multipliedBy(_212_v1));
          }
        }
        return _coll11;
      }()).length), _215_v4);
      _361_v132 = _nw66;
      let _363_v133;
      _363_v133 = _module.D2.create_DC6(_361_v132);
      let _364_v134;
      let _nw67 = Array((new BigNumber(29)).toNumber());
      _nw67[(_dafny.ZERO).toNumber()] = _361_v132;
      _nw67[(_dafny.ONE).toNumber()] = _361_v132;
      _nw67[(new BigNumber(2)).toNumber()] = (_363_v133).dtor_cf6;
      _nw67[(new BigNumber(3)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(4)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(5)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(6)).toNumber()] = ((false) ? (_361_v132) : (_361_v132));
      _nw67[(new BigNumber(7)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(8)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(9)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(10)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(11)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(12)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(13)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(14)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(15)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(16)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(17)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(18)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(19)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(20)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(21)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(22)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(23)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(24)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(25)).toNumber()] = ((_215_v4) ? (_361_v132) : (_361_v132));
      _nw67[(new BigNumber(26)).toNumber()] = ((_215_v4) ? (_361_v132) : (_361_v132));
      _nw67[(new BigNumber(27)).toNumber()] = _361_v132;
      _nw67[(new BigNumber(28)).toNumber()] = _361_v132;
      _364_v134 = _nw67;
      let _index63 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length));
      (_364_v134)[_index63] = _361_v132;
      let _365_v135;
      let _init9 = ((_366_v4, _367_v2, _368_v132) => function (_369_i8) {
        return _module.D0.create_DC1(_dafny.MultiSet.fromElements(_dafny.Seq.of(_366_v4, _366_v4), (_module.D3.create_DC9(_dafny.Seq.of(_366_v4, _366_v4))).dtor_cf10, _dafny.Seq.of((_module.D3.create_DC10(_367_v2, _366_v4)).dtor_cf12)), (_368_v132).f9);
      })(_215_v4, _213_v2, _361_v132);
      let _nw68 = Array((new BigNumber(4)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw68.length); _i0_9++) {
        _nw68[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _365_v135 = _nw68;
      let _370_v136;
      _370_v136 = _dafny.Seq.of(_215_v4);
      let _371_v137;
      _371_v137 = _dafny.MultiSet.fromElements(_370_v136);
      let _372_v138;
      _372_v138 = _module.D0.create_DC1(_371_v137, (_361_v132).f9);
      let _index64 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_365_v135).length));
      (_365_v135)[_index64] = _372_v138;
      let _373_v139;
      let _nw69 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _373_v139 = _nw69;
      let _374_v140;
      _374_v140 = _dafny.Seq.of(_373_v139);
      let _375_v141;
      _375_v141 = _dafny.MultiSet.fromElements(_306_v81);
      let _index65 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length));
      let _index66 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_365_v135).length));
      let _rhs53 = new BigNumber((_dafny.Seq.Concat(_374_v140, _dafny.Seq.of(_373_v139))).length);
      let _rhs54 = _361_v132;
      let _rhs55 = _372_v138;
      let _rhs56 = _module.__default.fm10(new BigNumber((_370_v136).length), !(((_361_v132).f8).isLessThanOrEqualTo((_361_v132).f8)), _212_v1, (_375_v141).Difference(_375_v141), _219_globalState);
      let _lhs51 = _219_globalState;
      let _lhs52 = _364_v134;
      let _lhs53 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length));
      let _lhs54 = _365_v135;
      let _lhs55 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_365_v135).length));
      _lhs51.f6 = _rhs53;
      _lhs52[_lhs53] = _rhs54;
      _lhs54[_lhs55] = _rhs55;
      _211_v0 = _rhs56;
      let _376_v142;
      _376_v142 = _module.D3.create_DC10(_213_v2, _215_v4);
      let _source3 = _376_v142;
      if (_source3.is_DC10) {
        let _377___mcc_h2 = (_source3).cf11;
        let _378___mcc_h3 = (_source3).cf12;
        let _379_cf12 = _378___mcc_h3;
        let _380_cf11 = _377___mcc_h2;
        let _381_v143;
        _381_v143 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_215_v4,(_361_v132).f8));
        _380_cf11 = _module.__default.fm11(_381_v143, (new BigNumber((_217_v6).cardinality())).plus(_212_v1), _211_v0, _211_v0, _219_globalState);
        _215_v4 = _215_v4;
        let _382_v144;
        _382_v144 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,_212_v1);
        let _383_v145;
        _383_v145 = _dafny.Map.Empty.slice().updateUnsafe(_373_v139,_382_v144);
        let _384_v146;
        let _nw70 = Array((new BigNumber(8)).toNumber());
        _nw70[(_dafny.ZERO).toNumber()] = _213_v2;
        _nw70[(_dafny.ONE).toNumber()] = _module.__default.fm11(_dafny.MultiSet.fromElements(_382_v144, (((_383_v145).contains(_373_v139)) ? ((_383_v145).get(_373_v139)) : (_382_v144))), _212_v1, _dafny.Seq.UnicodeFromString("unvm"), _211_v0, _219_globalState);
        _nw70[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw70[(new BigNumber(3)).toNumber()] = _380_cf11;
        _nw70[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw70[(new BigNumber(5)).toNumber()] = _380_cf11;
        _nw70[(new BigNumber(6)).toNumber()] = _380_cf11;
        _nw70[(new BigNumber(7)).toNumber()] = _380_cf11;
        _384_v146 = _nw70;
        let _index67 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_384_v146).length));
        (_384_v146)[_index67] = _213_v2;
        let _index68 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_384_v146).length));
        (_384_v146)[_index68] = _380_cf11;
        let _source4 = _326_v96;
        if (_source4.is_DC1) {
          let _385___mcc_h9 = (_source4).cf1;
          let _386___mcc_h10 = (_source4).cf2;
          let _387_cf2 = _386___mcc_h10;
          let _388_cf1 = _385___mcc_h9;
          (_219_globalState).f3 = (_361_v132).f9;
          _387_cf2 = ((_361_v132).f9).plus((_dafny.ZERO).minus(_212_v1));
          let _389_v147;
          _389_v147 = _dafny.Map.Empty.slice().updateUnsafe(_257_v40,_306_v81);
          let _390_v148;
          _390_v148 = _dafny.Set.fromElements(new BigNumber((_307_v82).length));
          _389_v147 = (_389_v147).update((_257_v40).update((_361_v132).f9, new BigNumber((_390_v148).length)), _307_v82);
          (_219_globalState).f6 = (_361_v132).f8;
        } else if (_source4.is_DC0) {
          let _391___mcc_h11 = (_source4).cf0;
          let _392_cf0 = _391___mcc_h11;
          let _393_v149;
          let _out11;
          _out11 = _module.__default.m0(_module.__default.safeModuloInt((_361_v132).f9, _212_v1), _359_v129, _370_v136, _219_globalState);
          _393_v149 = _out11;
          (_219_globalState).f6 = _module.__default.safeModuloInt(_212_v1, (_361_v132).f9);
          let _394_v150;
          let _out12;
          _out12 = _module.__default.m0((_361_v132).f9, _359_v129, _370_v136, _219_globalState);
          _394_v150 = _out12;
          (_219_globalState).f6 = (_361_v132).f8;
        } else {
          let _395___mcc_h12 = (_source4).cf3;
          let _396_cf3 = _395___mcc_h12;
          _380_cf11 = _380_cf11;
          let _397_v151;
          _397_v151 = _dafny.Map.Empty.slice().updateUnsafe((_361_v132).f8,_dafny.Map.Empty.slice().updateUnsafe(_373_v139,_212_v1));
          let _398_v152;
          _398_v152 = _dafny.Map.Empty.slice().updateUnsafe(_373_v139,_212_v1);
          _397_v151 = (_397_v151).update((_361_v132).f9, _398_v152);
          let _399_v153;
          let _out13;
          _out13 = _module.__default.m0((_361_v132).f9, _module.__default.fm12((_361_v132).f8, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(546)), ((_400_v146) => function (_401_i9) {
            return (_400_v146)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_400_v146).length))];
          })(_384_v146))).length), _module.__default.fm13((_361_v132).f8, _219_globalState), _219_globalState), _370_v136, _219_globalState);
          _399_v153 = _out13;
          (_219_globalState).f3 = new BigNumber((_399_v153).length);
        }
      } else if (_source3.is_DC11) {
        let _402___mcc_h4 = (_source3).cf13;
        let _403___mcc_h5 = (_source3).cf14;
        let _404___mcc_h6 = (_source3).cf15;
        let _405___mcc_h7 = (_source3).cf16;
        let _406_cf16 = _405___mcc_h7;
        let _407_cf15 = _404___mcc_h6;
        let _408_cf14 = _403___mcc_h5;
        let _409_cf13 = _402___mcc_h4;
        let _410_v154;
        let _init10 = ((_411_cf16) => function (_412_i10) {
          return (_411_cf16) && (_411_cf16);
        })(_406_cf16);
        let _nw71 = Array((_dafny.ONE).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw71.length); _i0_10++) {
          _nw71[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _410_v154 = _nw71;
        let _index69 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length));
        (_410_v154)[_index69] = _215_v4;
        let _index70 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length));
        (_410_v154)[_index70] = !((((_361_v132).f8).multipliedBy(_212_v1)).isLessThanOrEqualTo(new BigNumber(390)));
        let _413_v155;
        _413_v155 = _module.D4.create_DC12(_373_v139);
        let _414_v156;
        let _nw72 = Array((new BigNumber(12)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _373_v139;
        _nw72[(_dafny.ONE).toNumber()] = _373_v139;
        _nw72[(new BigNumber(2)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(3)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(4)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(5)).toNumber()] = (_413_v155).dtor_cf17;
        _nw72[(new BigNumber(6)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(7)).toNumber()] = ((_215_v4) ? (_373_v139) : (_373_v139));
        _nw72[(new BigNumber(8)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(9)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(10)).toNumber()] = _373_v139;
        _nw72[(new BigNumber(11)).toNumber()] = _373_v139;
        _414_v156 = _nw72;
        let _index71 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_414_v156).length));
        (_414_v156)[_index71] = _373_v139;
        let _index72 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_414_v156).length));
        let _rhs57 = _373_v139;
        let _rhs58 = (_361_v132).f8;
        let _lhs56 = _414_v156;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_414_v156).length));
        _lhs56[_lhs57] = _rhs57;
        _409_cf13 = _rhs58;
        let _415_v157;
        _415_v157 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_module.__default.fm3(_212_v1, _219_globalState));
        let _416_v158;
        _416_v158 = _dafny.Set.fromElements(_257_v40, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_216_v5).length),new BigNumber((_415_v157).length)), _257_v40);
        if (!((_416_v158).IsDisjointFrom(_416_v158))) {
          (_219_globalState).f3 = ((_406_cf16) ? ((_361_v132).f9) : ((_361_v132).f8));
          let _417_v159;
          _417_v159 = _dafny.Map.Empty.slice().updateUnsafe(_410_v154,(_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]);
          let _index73 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length));
          let _nw73 = new _module.C0();
          _nw73.__ctor(_212_v1, new BigNumber((_417_v159).length), _407_cf15);
          (_364_v134)[_index73] = _nw73;
          _211_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_211_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-203)), function (_418_i11) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uw"), _211_v0));
          let _419_v160;
          let _nw74 = Array((new BigNumber(14)).toNumber());
          _419_v160 = _nw74;
          let _420_v161;
          let _nw75 = new _module.C0();
          _nw75.__ctor((_361_v132).f8, new BigNumber(-634), (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]);
          _420_v161 = _nw75;
          let _index74 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_419_v160).length));
          (_419_v160)[_index74] = _420_v161;
          let _index75 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_419_v160).length));
          (_419_v160)[_index75] = _420_v161;
          _211_v0 = _211_v0;
        } else {
          let _421_v162;
          let _nw76 = Array((_dafny.ONE).toNumber()).fill(_module.D1.Default());
          _421_v162 = _nw76;
          let _index76 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_421_v162).length));
          (_421_v162)[_index76] = _308_v83;
          let _index77 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_421_v162).length));
          (_421_v162)[_index77] = _module.D1.create_DC3(_307_v82);
          (_219_globalState).f6 = (((_361_v132).f8).plus((_361_v132).f9)).minus(_module.__default.safeModuloInt((_361_v132).f9, _409_cf13));
          let _422_v163;
          _422_v163 = _dafny.Map.Empty.slice().updateUnsafe((_361_v132).f9,(_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]);
          _422_v163 = (_422_v163).update((_361_v132).f9, _407_cf15);
          _407_cf15 = _407_cf15;
          let _index78 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length));
          (_410_v154)[_index78] = (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))];
        }
        let _423_v164;
        _423_v164 = _dafny.Set.fromElements((_361_v132).f9, (_dafny.ZERO).minus((_361_v132).f8));
        if ((function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of (_423_v164).Elements) {
            let _424_v165 = _compr_12;
            if ((_423_v164).contains(_424_v165)) {
              _coll12.add((_424_v165).plus(new BigNumber((_dafny.Seq.of(true, true)).length)));
            }
          }
          return _coll12;
        }()).equals((_423_v164).Difference(_423_v164))) {
          let _425_v166;
          let _nw77 = new _module.C0();
          _nw77.__ctor(_212_v1, (_361_v132).f8, (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]);
          _425_v166 = _nw77;
          _212_v1 = ((new BigNumber(620)).multipliedBy((_361_v132).f8)).plus((_372_v138).dtor_cf2);
          let _426_v167;
          _426_v167 = _module.D3.create_DC11(new BigNumber((_415_v157).length), new BigNumber(642), _module.__default.fm3((_425_v166).fm4(_219_globalState), _219_globalState), (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]);
          let _427_v168;
          _427_v168 = _dafny.Seq.of(_426_v167);
          _427_v168 = _427_v168;
          _407_cf15 = (!(((_406_cf16) ? ((_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))]) : (_407_cf15)))) && (true);
          let _428_v169;
          _428_v169 = _dafny.Map.Empty.slice().updateUnsafe(false,(_425_v166).f9);
          let _index79 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length));
          (_410_v154)[_index79] = (((_370_v136)[_module.__default.safeIndex(_408_cf14, new BigNumber((_370_v136).length))]) ? ((new BigNumber((_428_v169).length)).isEqualTo((_425_v166).f9)) : (_407_cf15));
        } else {
          let _429_v170;
          _429_v170 = _dafny.Map.Empty.slice().updateUnsafe(_409_cf13,_dafny.Seq.update(_211_v0, _module.__default.safeIndex(_408_cf14, new BigNumber((_211_v0).length)), _213_v2));
          let _430_v171;
          let _out14;
          _out14 = _module.__default.m0((_361_v132).f9, _359_v129, _module.__default.fm6(new BigNumber(((((_429_v170).contains((_361_v132).f8)) ? ((_429_v170).get((_361_v132).f8)) : (_211_v0))).length), _213_v2, !(false), _module.__default.fm3(_409_cf13, _219_globalState), _219_globalState), _219_globalState);
          _430_v171 = _out14;
          let _rhs59 = _359_v129;
          let _rhs60 = (_212_v1).plus(new BigNumber((_dafny.Seq.of(true, _215_v4, (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))])).length));
          let _lhs58 = _219_globalState;
          _359_v129 = _rhs59;
          _lhs58.f3 = _rhs60;
          let _index80 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length));
          (_410_v154)[_index80] = _407_cf15;
          let _431_v172;
          let _nw78 = Array((new BigNumber(12)).toNumber()).fill([]);
          _431_v172 = _nw78;
          let _index81 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_431_v172).length));
          (_431_v172)[_index81] = _410_v154;
          let _index82 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_431_v172).length));
          let _rhs61 = _408_cf14;
          let _rhs62 = _410_v154;
          let _rhs63 = _359_v129;
          let _lhs59 = _219_globalState;
          let _lhs60 = _431_v172;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_431_v172).length));
          _lhs59.f3 = _rhs61;
          _lhs60[_lhs61] = _rhs62;
          _359_v129 = _rhs63;
          let _432_v173;
          let _out15;
          _out15 = _module.__default.m0((_361_v132).fm5(_415_v157, _213_v2, new BigNumber(213), (_410_v154)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_410_v154).length))], _219_globalState), _359_v129, _dafny.Seq.of(_215_v4), _219_globalState);
          _432_v173 = _out15;
        }
      } else {
        let _433___mcc_h8 = (_source3).cf10;
        let _434_cf10 = _433___mcc_h8;
        if (!(!(!((_215_v4) === (_215_v4))))) {
          let _435_v174;
          let _nw79 = Array((new BigNumber(20)).toNumber()).fill(false);
          _435_v174 = _nw79;
          let _index83 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_435_v174).length));
          (_435_v174)[_index83] = (!(_215_v4)) === (_215_v4);
          let _index84 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_435_v174).length));
          (_435_v174)[_index84] = !(new BigNumber((_257_v40).length)).isEqualTo((_361_v132).f8);
          (_219_globalState).f6 = new BigNumber((_dafny.Seq.Concat(_211_v0, _211_v0)).length);
          let _436_v175;
          let _nw80 = Array((_dafny.ONE).toNumber());
          _436_v175 = _nw80;
          let _437_v176;
          let _nw81 = new _module.C0();
          _nw81.__ctor((_dafny.ZERO).minus((_361_v132).f8), (_361_v132).f8, _215_v4);
          _437_v176 = _nw81;
          let _index85 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_436_v175).length));
          (_436_v175)[_index85] = _437_v176;
          let _index86 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_436_v175).length));
          (_436_v175)[_index86] = _437_v176;
          let _438_v177;
          let _out16;
          _out16 = _module.__default.m0((_361_v132).f8, _359_v129, _434_cf10, _219_globalState);
          _438_v177 = _out16;
          let _index87 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_435_v174).length));
          (_435_v174)[_index87] = _module.__default.fm3(((_361_v132).f8).multipliedBy((_361_v132).f9), _219_globalState);
        } else {
          let _439_v178;
          _439_v178 = _dafny.Seq.of(_361_v132, _361_v132);
          let _440_v179;
          let _nw82 = new _module.C0();
          _nw82.__ctor(new BigNumber((_211_v0).length), new BigNumber((_439_v178).length), _215_v4);
          _440_v179 = _nw82;
          let _441_v180;
          _441_v180 = _dafny.Seq.of(_440_v179);
          let _442_v181;
          _442_v181 = _dafny.MultiSet.fromElements(_440_v179);
          let _443_v182;
          let _nw83 = Array((new BigNumber(23)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _440_v179;
          _nw83[(_dafny.ONE).toNumber()] = _440_v179;
          _nw83[(new BigNumber(2)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(3)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(4)).toNumber()] = (_441_v180)[_module.__default.safeIndex(new BigNumber(((_442_v181).update(_440_v179, _module.__default.abs((_361_v132).f8))).cardinality()), new BigNumber((_441_v180).length))];
          _nw83[(new BigNumber(5)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(6)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(7)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(8)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(9)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(10)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(11)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(12)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(13)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(14)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(15)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(16)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(17)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(18)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(19)).toNumber()] = ((true) ? (_440_v179) : (_440_v179));
          _nw83[(new BigNumber(20)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(21)).toNumber()] = _440_v179;
          _nw83[(new BigNumber(22)).toNumber()] = _440_v179;
          _443_v182 = _nw83;
          let _444_v183;
          _444_v183 = _440_v179;
          let _445_v184;
          _445_v184 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(793),_440_v179);
          let _nw84 = Array((new BigNumber(20)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _440_v179;
          _nw84[(_dafny.ONE).toNumber()] = _440_v179;
          _nw84[(new BigNumber(2)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(3)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(4)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(5)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(6)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(7)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(8)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(9)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(10)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(11)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(12)).toNumber()] = (_444_v183);
          _nw84[(new BigNumber(13)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(14)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(15)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(16)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(17)).toNumber()] = (((_445_v184).contains((_361_v132).f8)) ? ((_445_v184).get((_361_v132).f8)) : (_440_v179));
          _nw84[(new BigNumber(18)).toNumber()] = _440_v179;
          _nw84[(new BigNumber(19)).toNumber()] = _440_v179;
          _443_v182 = _nw84;
          let _446_v185;
          let _out17;
          _out17 = _module.__default.m0(_212_v1, _359_v129, _370_v136, _219_globalState);
          _446_v185 = _out17;
          let _447_v186;
          let _init11 = function (_448_i12) {
            return true;
          };
          let _nw85 = Array((new BigNumber(26)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw85.length); _i0_11++) {
            _nw85[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _447_v186 = _nw85;
          let _449_v187;
          _449_v187 = _dafny.Map.Empty.slice().updateUnsafe(_361_v132,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(144),_440_v179)).length));
          let _rhs64 = _447_v186;
          let _rhs65 = _440_v179.f7;
          let _rhs66 = (_212_v1).multipliedBy(((_361_v132).f9).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_449_v187)).length))));
          let _rhs67 = _373_v139;
          let _rhs68 = ((_361_v132).f8).multipliedBy((_dafny.ZERO).minus((_361_v132).f8));
          let _lhs62 = _219_globalState;
          _447_v186 = _rhs64;
          _215_v4 = _rhs65;
          _lhs62.f6 = _rhs66;
          _373_v139 = _rhs67;
          _212_v1 = _rhs68;
          let _450_v188;
          _450_v188 = _dafny.MultiSet.fromElements(_447_v186);
          let _451_v189;
          _451_v189 = _module.D3.create_DC11((_361_v132).f9, (_440_v179).fm4(_219_globalState), (_450_v188).IsSubsetOf(_450_v188), true);
          _451_v189 = _451_v189;
          _434_cf10 = _dafny.Seq.Concat(_434_cf10, _370_v136);
        }
        _215_v4 = !(!(_215_v4)) || (((_361_v132).f9).isEqualTo((_361_v132).f9));
        let _452_v190;
        let _out18;
        _out18 = _module.__default.m0((_361_v132).f9, _359_v129, _434_cf10, _219_globalState);
        _452_v190 = _out18;
        let _rhs69 = _215_v4;
        let _rhs70 = (_361_v132).f8;
        let _lhs63 = _219_globalState;
        _215_v4 = _rhs69;
        _lhs63.f6 = _rhs70;
      }
      _215_v4 = _215_v4;
      let _rhs71 = (_370_v136)[_module.__default.safeIndex(((_215_v4) ? ((_361_v132).f8) : ((_361_v132).f9)), new BigNumber((_370_v136).length))];
      let _rhs72 = true;
      let _rhs73 = (_dafny.ZERO).minus((_212_v1).plus(_212_v1));
      let _rhs74 = (_361_v132).f9;
      let _lhs64 = _219_globalState;
      let _lhs65 = _219_globalState;
      _215_v4 = _rhs71;
      _215_v4 = _rhs72;
      _lhs64.f3 = _rhs73;
      _lhs65.f3 = _rhs74;
      if (!(_215_v4) || (true)) {
        let _453_v191;
        let _out19;
        _out19 = _module.__default.m0(_212_v1, _dafny.MultiSet.FromArray(_370_v136), _dafny.Seq.Concat(_370_v136, _dafny.Seq.of(_215_v4, _215_v4, _215_v4)), _219_globalState);
        _453_v191 = _out19;
        _215_v4 = true;
        _215_v4 = !_dafny.areEqual(_211_v0, _453_v191);
        let _454_v192;
        let _nw86 = Array((new BigNumber(18)).toNumber()).fill(false);
        _454_v192 = _nw86;
        let _455_v193;
        _455_v193 = _module.D1.create_DC3(_dafny.Seq.update(_307_v82, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_454_v192,_212_v1)).length), new BigNumber((_307_v82).length)), (_361_v132).f8));
        let _456_v194;
        _456_v194 = _module.D1.create_DC5(_455_v193);
        let _source5 = _456_v194;
        if (_source5.is_DC4) {
          let _457_v195;
          let _nw87 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _457_v195 = _nw87;
          let _index88 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_457_v195).length));
          (_457_v195)[_index88] = _306_v81;
          let _index89 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_457_v195).length));
          (_457_v195)[_index89] = _306_v81;
          let _458_v196;
          _458_v196 = _dafny.Set.fromElements(new BigNumber(-939), (_361_v132).f8);
          let _459_v197;
          _459_v197 = _module.D6.create_DC16(_458_v196);
          (_219_globalState).f3 = (new BigNumber(((_459_v197).dtor_cf24).length)).plus(_212_v1);
          let _460_v198;
          _460_v198 = _dafny.Map.Empty.slice().updateUnsafe((_361_v132).f8,false);
          _460_v198 = (_460_v198).update(_212_v1, false);
          let _461_v199;
          _461_v199 = _dafny.Set.fromElements(_215_v4);
          let _462_v200;
          _462_v200 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(_370_v136, _module.__default.safeIndex((_361_v132).f8, new BigNumber((_370_v136).length)), _215_v4)).length), new BigNumber((_211_v0).length)),_461_v199);
          _462_v200 = (_462_v200).update(_module.__default.safeModuloInt((_dafny.ZERO).minus((_307_v82)[_module.__default.safeIndex(_212_v1, new BigNumber((_307_v82).length))]), new BigNumber(-792)), (_461_v199).Difference(_461_v199));
        } else if (_source5.is_DC3) {
          let _463___mcc_h13 = (_source5).cf4;
          let _464_cf4 = _463___mcc_h13;
          _215_v4 = !(_215_v4);
          let _465_v201;
          let _nw88 = new _module.C0();
          _nw88.__ctor(new BigNumber(965), (_361_v132).f8, _module.__default.fm3(_212_v1, _219_globalState));
          _465_v201 = _nw88;
          let _466_v202;
          _466_v202 = _dafny.Set.fromElements(_465_v201.f7, _465_v201.f7, _465_v201.f7, false, _465_v201.f7);
          _465_v201 = ((!((_361_v132).f9).isEqualTo(_module.__default.fm0(new BigNumber((_466_v202).length), true, _dafny.Seq.of((_361_v132).f9, _212_v1, (_361_v132).f9), _213_v2, _219_globalState))) ? (_465_v201) : (_465_v201));
          let _467_v203;
          _467_v203 = _dafny.Map.Empty.slice().updateUnsafe(_465_v201.f7,(_dafny.ZERO).minus(_212_v1));
          let _468_v204;
          let _nw89 = new _module.C0();
          _nw89.__ctor(new BigNumber(((_467_v203).Merge(_467_v203)).length), (_361_v132).f8, false);
          _468_v204 = _nw89;
          (_219_globalState).f3 = (((_468_v204).fm4(_219_globalState)).minus((_361_v132).f8)).plus((_361_v132).f9);
        } else {
          let _469___mcc_h14 = (_source5).cf5;
          let _470_cf5 = _469___mcc_h14;
          let _index90 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length));
          (_364_v134)[_index90] = ((((_361_v132).f9).isEqualTo((_361_v132).f8)) ? (_361_v132) : (_361_v132));
          let _471_v205;
          _471_v205 = _dafny.Map.Empty.slice().updateUnsafe((_361_v132).f9,_454_v192);
          _471_v205 = (_471_v205).update(((_215_v4) ? ((_361_v132).f9) : ((_361_v132).f9)), _454_v192);
          _215_v4 = (_module.__default.fm14(_219_globalState)).IsSubsetOf(_254_v37);
          let _index91 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_373_v139).length));
          (_373_v139)[_index91] = (_361_v132).f8;
          let _index92 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_373_v139).length));
          (_373_v139)[_index92] = (_361_v132).f9;
        }
        _215_v4 = _215_v4;
      } else {
        (_219_globalState).f3 = _212_v1;
        let _472_v206;
        _472_v206 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(!(_215_v4)),_373_v139);
        let _473_v207;
        _473_v207 = _dafny.Set.fromElements(_215_v4);
        _472_v206 = (_472_v206).update(_473_v207, _373_v139);
        (_219_globalState).f3 = (_361_v132).f8;
        if (_215_v4) {
          let _474_v208;
          let _nw90 = new _module.C0();
          _nw90.__ctor(((_361_v132).f8).multipliedBy((_361_v132).f9), _212_v1, _215_v4);
          _474_v208 = _nw90;
          let _475_v209;
          _475_v209 = _dafny.MultiSet.fromElements(_361_v132, _361_v132);
          let _476_v210;
          _476_v210 = _dafny.Map.Empty.slice().updateUnsafe(_475_v209,new BigNumber((_dafny.Set.fromElements(true, _215_v4, _215_v4)).length));
          _215_v4 = !(!(_476_v210).contains((_475_v209).Intersect(_475_v209)));
          (_219_globalState).f6 = new BigNumber(65);
          let _477_v212;
          _477_v212 = _dafny.Set.fromElements(new BigNumber((_211_v0).length), _module.__default.safeModuloInt(new BigNumber((function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of (_306_v81).Elements) {
              let _478_v211 = _compr_13;
              if (_dafny.Seq.contains(_306_v81, _478_v211)) {
                _coll13.add((_478_v211).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new BigNumber(931))).length)));
              }
            }
            return _coll13;
          }()).length), (_361_v132).f8));
          _477_v212 = ((_477_v212).Intersect(_477_v212)).Intersect(_477_v212);
          _213_v2 = _213_v2;
        } else {
          _211_v0 = _dafny.Seq.Concat(_module.__default.fm10((_361_v132).f8, _215_v4, new BigNumber(((_dafny.MultiSet.fromElements(_215_v4, _215_v4)).update(_module.__default.fm3((_361_v132).f8, _219_globalState), _module.__default.abs((_361_v132).f8))).cardinality()), _375_v141, _219_globalState), _211_v0);
          _361_v132 = (_364_v134)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_364_v134).length))];
          let _index93 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_373_v139).length));
          (_373_v139)[_index93] = (_dafny.ZERO).minus((_361_v132).f8);
          let _index94 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_373_v139).length));
          (_373_v139)[_index94] = (_212_v1).minus((_361_v132).f9);
          let _479_v213;
          _479_v213 = _dafny.Map.Empty.slice().updateUnsafe(!(_215_v4),_dafny.Set.fromElements(_215_v4));
          let _480_v214;
          let _out20;
          _out20 = _module.__default.m0(new BigNumber(((((_479_v213).contains(_215_v4)) ? ((_479_v213).get(_215_v4)) : (_473_v207))).length), _359_v129, _370_v136, _219_globalState);
          _480_v214 = _out20;
          let _481_v215;
          let _nw91 = new _module.C0();
          _nw91.__ctor((_361_v132).f8, _module.__default.safeModuloInt((_361_v132).f9, (_361_v132).f8), ((_361_v132).f8).isLessThanOrEqualTo((_373_v139)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_373_v139).length))]));
          _481_v215 = _nw91;
        }
        let _482_v216;
        _482_v216 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,_215_v4);
        let _483_v217;
        _483_v217 = _dafny.Map.Empty.slice().updateUnsafe(_215_v4,!(!((_370_v136)[_module.__default.safeIndex(new BigNumber((_482_v216).length), new BigNumber((_370_v136).length))])));
        let _484_v218;
        _484_v218 = _module.D3.create_DC9(_dafny.Seq.update(_dafny.Seq.update(_370_v136, _module.__default.safeIndex((_361_v132).f9, new BigNumber((_370_v136).length)), false), _module.__default.safeIndex((_361_v132).fm5(_482_v216, _213_v2, new BigNumber(472), _215_v4, _219_globalState), new BigNumber((_dafny.Seq.update(_370_v136, _module.__default.safeIndex((_361_v132).f9, new BigNumber((_370_v136).length)), false)).length)), _215_v4));
        let _485_v219;
        let _out21;
        _out21 = _module.__default.m0(_module.__default.safeDivisionInt((_361_v132).f8, _212_v1), _dafny.MultiSet.fromElements(_215_v4, _215_v4), (_484_v218).dtor_cf10, _219_globalState);
        _485_v219 = _out21;
      }
      _212_v1 = (_212_v1).minus(new BigNumber((_375_v141).cardinality()));
      process.stdout.write((_211_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_212_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_213_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(779),new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_215_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_217_v6).equals(_dafny.MultiSet.fromElements(new BigNumber(-166)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_218_v7).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_219_globalState).f1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_219_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_globalState.f4).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_219_globalState).f5).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_219_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v37).equals(_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v38).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_256_v39, _dafny.Seq.of(_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v40).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(778),new BigNumber(778)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_306_v81, _dafny.Seq.of(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_307_v82, _dafny.Seq.of(new BigNumber(2), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_308_v83).dtor_cf4, _dafny.Seq.of(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_309_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_326_v96).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_359_v129).equals(_dafny.MultiSet.fromElements(true, true, true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_360_v130).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_dafny.MultiSet.fromElements(true, true, true, true, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_361_v132).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_361_v132).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_363_v133).dtor_cf6).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_363_v133).dtor_cf6).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_363_v133).dtor_cf6.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[_dafny.ZERO]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[_dafny.ZERO]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[_dafny.ZERO].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[_dafny.ONE]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[_dafny.ONE]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[_dafny.ONE].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(2)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(2)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(2)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(3)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(3)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(3)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(4)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(4)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(4)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(5)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(5)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(5)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(6)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(6)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(6)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(7)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(7)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(7)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(8)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(8)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(8)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(9)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(9)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(9)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(10)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(10)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(10)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(11)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(11)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(11)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(12)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(12)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(12)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(13)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(13)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(13)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(14)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(14)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(14)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(15)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(15)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(15)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(16)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(16)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(16)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(17)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(17)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(17)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(18)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(18)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(18)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(19)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(19)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(19)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(20)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(20)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(20)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(21)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(21)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(21)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(22)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(22)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(22)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(23)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(23)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(23)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(24)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(24)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(24)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(25)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(25)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(25)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(26)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(26)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(26)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(27)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(27)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(27)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(28)]).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_364_v134)[new BigNumber(28)]).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_364_v134)[new BigNumber(28)].f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_365_v135)[_dafny.ZERO]).dtor_cf1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_365_v135)[_dafny.ZERO]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_365_v135)[_dafny.ONE]).dtor_cf1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false), _dafny.Seq.of(false, false), _dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_365_v135)[_dafny.ONE]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_365_v135)[new BigNumber(2)]).dtor_cf1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false), _dafny.Seq.of(false, false), _dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_365_v135)[new BigNumber(2)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_365_v135)[new BigNumber(3)]).dtor_cf1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false), _dafny.Seq.of(false, false), _dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_365_v135)[new BigNumber(3)]).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_370_v136, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_371_v137).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_372_v138).dtor_cf1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_372_v138).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_373_v139)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_374_v140).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_375_v141).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_376_v142).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_376_v142).dtor_cf12));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D0(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.MultiSet.Empty, _dafny.ZERO);
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
    static create_DC4() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
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
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4();
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
    static create_DC7(cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D2(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, null);
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
    static create_DC11(cf13, cf14, cf15, cf16) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
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
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC13(cf18, cf19, cf20) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC14(cf21, cf22) {
      let $dt = new D4(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf17 === other.cf17;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC15(cf23) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D6(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC16(cf24) {
      let $dt = new D6(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, false, false);
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
    static create_DC20(cf30, cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf29) {
      let $dt = new D7(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D7(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && this.cf31 === other.cf31 && this.cf32 === other.cf32;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(false, false, null);
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
    static create_DC23(cf35, cf36, cf37) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC24(cf38, cf39, cf40) {
      let $dt = new D8(1);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC25(cf41) {
      let $dt = new D8(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC22(cf34) {
      let $dt = new D8(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC26(cf42) {
      let $dt = new D8(4);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC22() { return this.$tag === 3; }
    get is_DC26() { return this.$tag === 4; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && this.cf36 === other.cf36 && this.cf37 === other.cf37;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf38 === other.cf38 && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(false, false, null);
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
    static create_DC28(cf44, cf45, cf46) {
      let $dt = new D9(0);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC29(cf47, cf48, cf49) {
      let $dt = new D9(1);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC30(cf50, cf51, cf52, cf53) {
      let $dt = new D9(2);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC27(cf43) {
      let $dt = new D9(3);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC27() { return this.$tag === 3; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC29" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC30" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48 && this.cf49 === other.cf49;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51) && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC28(false, _dafny.ZERO, []);
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
    static create_DC31(cf54) {
      let $dt = new D10(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC31" + "(" + this.cf54.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC33(cf56) {
      let $dt = new D11(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf55) {
      let $dt = new D11(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC34(cf57) {
      let $dt = new D11(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC33(false);
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
    static create_DC36(cf59, cf60, cf61, cf62, cf63) {
      let $dt = new D12(0);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC37(cf64, cf65, cf66) {
      let $dt = new D12(1);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC35(cf58) {
      let $dt = new D12(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf58) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64) && this.cf65 === other.cf65 && this.cf66 === other.cf66;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf58, other.cf58);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC36(_dafny.ZERO, _dafny.ZERO, false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = _dafny.ZERO;
      this.f4 = _dafny.Map.Empty;
      this.f6 = _dafny.ZERO;
      this._f0 = false;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f2 = false;
      this._f5 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
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
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f7 = false;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    set f7(value) {
      let _this = this;
      _this._f7 = value;
    };
    __ctor(f8, f9, f7) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f7 = f7;
      return;
    }
    fm4(globalState) {
      let _this = this;
      return (new BigNumber(-635)).minus((_this).f9);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this.f7) ? (_this.f7) : (_this.f7)),_this.f7)).length);
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
